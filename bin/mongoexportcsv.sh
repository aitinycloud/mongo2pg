#!/usr/bin/env bash
#author:zhangjinming
#date: 2019-04-24
BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)

cd $PROJECTPATH
IP=`echo $IP`
PORT=`echo $PORT`
FILENAME=`echo $FILENAME`
LIMIT=`echo $LIMIT`
SKIP=`echo $SKIP`

./bin/mongoexport -h $IP --port $PORT -d insight -c DeviceInfo --type=csv -f hardwarearch,manufacturer,manufacturerOUI,productClass,serialNumber,hardwareVersion,softwareVersion,probeId,province,city,district,town,village,accessNE,aggregationNE,coreNE,account,userName,macAddr,ipv4Addr,ipv6Addr,hubId,type,name,description,vendor,version,interfaceVersion -o migratedata/$FILENAME --sort={_id:1} --limit=$LIMIT --skip=$SKIP
