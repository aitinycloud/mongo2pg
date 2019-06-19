#!/usr/bin/env bash
#author:zhangjinming
#date: 2019-04-24
BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)

cd $PROJECTPATH
PGUSER=`echo $PGUSER`
PGPASSWORD=`echo $PGPASSWORD`
PGIP=`echo $PGIP`
PGPORT=`echo $PGPORT`
DATABASE="insight"
TABLENAME=`echo $TABLENAME`
FILENAME=`echo $FILENAME`

$PROJECTPATH/bin/opt/PostgreSQL/10/bin/psql -h $PGIP -p $PGPORT -d $DATABASE -U $PGUSER -c '\copy '$TABLENAME' (hardwarearch,manufacturer,manufactureroui,productclass,serialnumber,hardwareversion,softwareversion,probeid,province,city,district,town,village,accessne,aggregationne,corene,account,username,macaddr,ipv4addr,ipv6addr,hubid,type,name,description,vendor,version,interfaceversion) from migratedata/'$FILENAME' with csv header'
exit 0
