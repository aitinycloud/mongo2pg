#!/usr/bin/env bash
#author:zhangjinming
#date: 2018-07-06

BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)

res=`ps -fe | grep $PROJECTPATH/bin/insight-dn-ni | grep -v grep`
if [ $? -ne 0 ]
then
	echo "stop"
	exit 1
else
	echo "running"
	exit 0
fi
