#!/usr/bin/env bash
#

BINPATH=$(cd `dirname $0`; pwd)
PROJECTPATH=$(cd $BINPATH/..; pwd)
proc_list=$PROJECTPATH/bin/insight-dn-ni

#echo "insight-dn-ni daemon start."
daemonstart() {
    while test true
    do 
        $PROJECTPATH/bin/check.sh >> /dev/null
        res=$?
        if [ $res -ne 0 ]
        then
            #echo "restart insight-dn-ni."
            $PROJECTPATH/bin/restart.sh
        fi
        sleep 10
    done
}

daemonstart &
