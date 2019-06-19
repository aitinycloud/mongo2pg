package handle

import "time"

//MONGODB_NAME
const MONGODB_NAME = "testdb"

//mongo timeout hours
const MONGOMAXTIMEOUTHOURS = 6

const TABLENAME = "devices_tmptable"
const NEWTANLENAME = "devices"

var DMMongoDbAddress = "mongo://127.0.0.1:27017"

var pgUser = "postgre"
var pgPwd = "postgre"
var pgIp = "127.0.0.1"
var pgPort = 5432
var pgDbName = "testdb"

var MongoExportPeriod = 120 * time.Second
