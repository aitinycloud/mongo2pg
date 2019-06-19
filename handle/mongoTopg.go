package handle

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"github.com/robfig/cron"
	"go.mongodb.org/mongo-driver/bson"
)

func TestExec() {
	MongoDBToPostgreHandle()
}

func MongoDBToCSV(skip uint, limit uint, csvName string) {
	tmphost := strings.Split(DMMongoDbAddress, ":")
	host := tmphost[0]
	port := tmphost[1]
	command := fmt.Sprintf(`export IP="%v";export PORT=%v;export FILENAME="%v";export LIMIT=%v;export SKIP=%v;./bin/mongoexportcsv.sh`,
		host, port, csvName, limit, skip)
	cmd := exec.Command("/bin/bash", "-c", command)
	output, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	//logging.Info("Execute Shell:%s finished with output:\n%s", command, string(output))
}

func GetDeviceInfoCount() uint {
	coll := DMGetCollection("insight", "DeviceInfo")
	count, err := coll.CountDocuments(context.Background(), bson.D{})
	if err != nil {
		fmt.Printf("coll.CountDocuments err : ", err)
		return 0
	}
	return uint(count)
}

func MongoDBToPostgreHandle() {
	var MAXCOUNT uint = 1000000
	const THERAD = 2
	totalCount := GetDeviceInfoCount()
	if totalCount == 0 {
		return
	}
	const MINMAXCOUNT = 500000
	count := 0
	csvNameArr := []string{}
	//clear csv.
	command := " rm -rf migratedata/* "
	cmd := exec.Command("/bin/bash", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
		return
	}
	for i := uint(0); i < totalCount; i += MAXCOUNT {
		skip := i
		limit := uint(MAXCOUNT)
		count = count + 1
		csvName := fmt.Sprintf("DeviceInfo_%v.csv", count)
		csvNameArr = append(csvNameArr, csvName)
		if totalCount < MAXCOUNT*THERAD {
			MongoDBToCSV(skip, limit, csvName)
		} else {
			if count%THERAD == 0 {
				MongoDBToCSV(skip, limit, csvName)
			} else {
				go func() {
					MongoDBToCSV(skip, limit, csvName)
				}()
			}
		}
	}
	// CSV To postgresql
	time.Sleep(1 * time.Second)
	CSVToPostgre(csvNameArr)
}

func CSVToPostgre(csvName []string) {
	db := pg.GetDB()
	if db == nil {
		fmt.Printf("db == nil. ")
		return
	}
	//TRUNCATE
	sqlstr := "TRUNCATE " + TABLENAME + ";"
	_, err := db.Exec(sqlstr)
	if err != nil {
		fmt.Printf("SQL fail, : ", sqlstr)
	}
	//import
	for i := 0; i < len(csvName); i++ {
		// shell psql import to postgresql.
		tmphost := strings.Split(setting.PgHost, ":")
		host := tmphost[0]
		port := tmphost[1]
		command := fmt.Sprintf(`export PGIP="%v";export PGPORT=%v;export PGUSER="%v";export PGPASSWORD="%v";export TABLENAME="%v";export FILENAME="%v";./bin/csvimportpg.sh`,
			host, port, setting.PgUser, setting.PgPassword, TABLENAME, csvName[i])

		cmd := exec.Command("/bin/bash", "-c", command)
		output, err := cmd.Output()
		if err != nil {
			fmt.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			return
		}
		fmt.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))
	}
	// Conv data.
	sqlstr = fmt.Sprintf("update %v set type='hgu' where type='3';update %v set type='ihgu' where type='6';update %v set type='stb' where type='7';",
		TABLENAME, TABLENAME, TABLENAME)
	_, err = db.Exec(sqlstr)
	if err != nil {
		fmt.Printf("SQL fail, : ", sqlstr)
	}

	// tx begin.
	tx, err := db.Begin()
	if err != nil {
		fmt.Printf(err)
	}
	// rename NEW -> NEW_1122
	var OldName = TABLENAME
	var NewName = NEWTANLENAME
	sqlstr = "alter table " + NewName + " rename to " + NewName + "_1122" + ";"
	_, err = tx.Exec(sqlstr)
	if err != nil {
		fmt.Printf("SQL INSTER fail, : ", sqlstr)
	}
	// rename OLD -> NEW
	sqlstr = "alter table " + OldName + " rename to " + NewName + ";"
	_, err = tx.Exec(sqlstr)
	if err != nil {
		fmt.Printf("SQL INSTER fail, : ", sqlstr)
	}
	// rename NEW_1122 -> OLD
	sqlstr = "alter table " + NewName + "_1122" + " rename to " + OldName + ";"
	_, err = tx.Exec(sqlstr)
	if err != nil {
		fmt.Printf("SQL INSTER fail, : ", sqlstr)
	}
	if err := tx.Commit(); err != nil {
		tx.Rollback()
		fmt.Printf(err)
		return
	}
}

func MongoToPgSetup() {
	c := cron.New()
	const MINTIME = 3
	Min := MongoExportPeriod % 60
	if Min < MINTIME {
		Min = MINTIME
	}
	spec := fmt.Sprintf("@every %vm0s", Min)
	c.AddFunc(spec, MongoDBToPostgreHandle)
	c.Start()
}
