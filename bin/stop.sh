#!/usr/bin/env bash
#author:zhangjinming
#date: 2018-07-06

BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)

SERVER_NAME=$PROJECTPATH/bin/insight-dn-ni

echo "insight-dn-ni stop."
runNum=`ps -ef | tr -s ' '| sort -n | grep $SERVER_NAME | grep -v "grep" | awk '{print $2}' | wc -l`
if [ $runNum -eq 1 ] ; then
    PID=`ps -ef | tr -s ' ' | sort -n | grep $SERVER_NAME | grep -v "grep" | awk '{print $2}'`
    if [ $PID -gt 1 ] ; then
        echo "$SERVER_NAME PID $PID. "
        sleep 1
        kill $PID
        sleep 1
        if [ $? -eq 0 ] ; then
            echo "停止服务 $SERVER_NAME 成功!!"
            exit 0 
        else
            echo "停止服务发生异常,请手工停止并检查原因."
        fi
    else
        echo "获取PID失败,获取PID为 $PID .请手工停止服务 $SERVER_NAME ."
    fi
else
    if [ $runNum -eq 0 ] ; then
        echo "当前正在运行服务数量 $runNum 个.无正在运行的服务."
        exit 0
    else
        echo "当前正在运行服务数量 $runNum 个,请手动选择停止!!"
    fi
fi
