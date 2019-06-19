package handle

import (
	"encoding/json"
	"testing"
	"time"

	"../models/pg"
)

func Test_HandleInit(t *testing.T) {
	Setup()
	t.Log("OK")
}

func Test_WorkVirtualDev(t *testing.T) {
	time.Sleep(1 * time.Second)
	msgmap := make(map[string]interface{})
	msgmap["probeId"] = "A4CA4EF23EE65CB2F8BCD222EC879F9D"
	msgmap["hubId"] = "A4CA4EF23EE65CB2F8BCD222EC879F9D"
	msgmap["type"] = "HGU"
	msgmap["name"] = "com.chinatelecom.all.smartgateway.xrobot_sd5116"
	msgmap["description"] = "description"
	msgmap["vendor"] = "testVixtel"
	msgmap["version"] = "testversion"
	msgmap["versionInfo"] = "versionInfo"
	msgmap["versionTarget"] = "versionTarget"
	msgmap["interfaceVersion"] = "ttt1.0"
	msgmap["signUpTime"] = 1542552399911154
	msgmap["upgradeTime"] = "0"
	msgmap["config"] = "config"
	msgmap["configInfo"] = "configInfo"
	msgmap["configTarget"] = "configTarget"
	msgmap["status"] = 1
	msgmap["lastOffTime"] = 1542552399911154
	msgmap["signInTime"] = 1542552399911154
	msgmap["signInCount"] = "2"
	msgmap["lifeTime"] = 1542552399911154
	msgmap["updateTime"] = 1542552399911154
	msgstr, err := json.Marshal(msgmap)
	if err != nil {
		t.Error(msgstr, err)
		return
	}
	//MQSendMsg(SUBJECT_VIRTUALDEV, string(msgstr))
	t.Log("OK")
	//time.Sleep(1 * time.Second)
}

func Test_WorkHostDevices(t *testing.T) {
	//Setup()
	time.Sleep(1 * time.Second)
	msgmap := make(map[string]interface{})
	msgmap["probeId"] = "A4CA4EF23EE65CB2F8BCD222EC879F9D"
	msgmap["hardwareArch"] = "hardwareArch"
	msgmap["manufacturer"] = "youhuatest"
	msgmap["manufacturerOUI"] = "manufacturerOUI"
	msgmap["productClass"] = "PT924G"
	msgmap["serialNumber"] = "serialNumber"
	msgmap["hardwareVersion"] = "V4.0.00"
	msgmap["softwareVersion"] = "V10.00.924R001"
	msgmap["macAddr"] = "2C431A9BFE70"
	msgmap["ipv4Addr"] = "114.242.25.188"
	msgmap["ipv6Addr"] = "ipv6Addr"
	msgmap["hostName"] = "hostNametest"
	msgmap["timeZone"] = "timeZone"
	msgmap["netDevCfg"] = "netDevCfgtest"
	msgmap["memSize"] = 40
	msgmap["cpuFreq"] = 16.9
	msgmap["diskSize"] = 12
	msgmap["osVersion"] = "osVersion"
	msgstr, err := json.Marshal(msgmap)
	if err != nil {
		t.Error(msgstr, err)
		return
	}
	//MQSendMsg(SUBJECT_HOSTDEVICE, string(msgstr))
	t.Log("OK")
	//for {
	//time.Sleep(3 * time.Second)
	//}
}

func Test_SetNeRelationStatusAndTimeByMacS(t *testing.T) {
	var macaddrs = []string{"100000000000", "100000000001", "100000000002", "100000000003", "100000000004"}
	var status int = 11
	var lastModifyTime int64 = time.Now().UnixNano() / 1000000
	SetNeRelationStatusAndTimeByMacS(macaddrs, status, lastModifyTime)
	t.Log("OK")
}

func Test_BulkUpdate(t *testing.T) {
	dbname := "insight_cfg"
	tablename := "virtualdev"
	condstr := "probeid"
	results := []map[string]string{}
	Item := map[string]string{}
	Item["probeid"] = "69C25DA08050DCC82C515309BCA82F11"
	Item["type"] = "IHGU"
	Item["name"] = "testname"
	Item["description"] = "testdescription"
	results = append(results, Item)
	Item["probeid"] = "69C25DA08050DCC82C515309BCA82F22"
	results = append(results, Item)
	Item["probeid"] = "69C25DA08050DCC82C515309BCA82F33"
	results = append(results, Item)
	pg.BulkUpdate(dbname, tablename, condstr, results)
	t.Log("OK")
}
