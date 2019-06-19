#!/usr/bin/env bash
#author:zhangjinming
#date: 2018-07-06

BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)


echo "$PROJECTPATH start."
cd $PROJECTPATH
export HREGSERVER="192.168.1.109:2379";export CONFIGPATH="/com.vixtel.insight.topology/192.168.1.161/usr.local.vixtel.insight.datanorma.ni";
$PROJECTPATH/bin/insight-dn-ni 2>>logs/info.log 1>>logs/info.log  &
echo "*********************************"
echo "insight-dn-ni process"
ps -fe | grep $PROJECTPATH/bin/insight-dn-ni | grep -v grep
echo "*********************************"
sleep 1
