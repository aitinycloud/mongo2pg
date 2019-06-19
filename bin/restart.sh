#!/usr/bin/env bash
#author:zhangjinming
#date: 2018-07-06

BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)

echo "insight-DataReprocessService restart."
cd $PROJECTPATH
$PROJECTPATH/bin/stop.sh
sleep 1
$PROJECTPATH/bin/start.sh
