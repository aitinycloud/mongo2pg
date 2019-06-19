#!/usr/bin/env bash
#author:zhangjinming
#date: 2018-07-06

BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)


echo "$PROJECTPATH start."
cd $PROJECTPATH
$PROJECTPATH/bin/insight-dn-ni
echo "*********************************"
echo "insight-dn-ni process"
ps -fe | grep $PROJECTPATH/bin/insight-dn-ni | grep -v grep
echo "*********************************"
sleep 1
