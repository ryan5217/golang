#!/bin/sh
# DateTime: 2017-01-18
# Author: zhaochenglan
# chkconfig:   - 84 16
# Source function library.
exit 0
source /etc/profile
. /etc/rc.d/init.d/functions
# Source networking configuration.
. /etc/sysconfig/network
# Check that networking is up.
[ "$NETWORKING" = "no" ] && exit 0
#port=6060
script_name=vigo.gaodun.com
lockfile='/var/lock/subsys/'${script_name}'rpc'
pid=`ps aux|grep /gaodun/domain/|grep ucenter|grep -E 'rpc$'|grep -v grep|awk  '{print $2}'`
script_pwd=`dirname $(readlink /proc/$$/fd/255)`
log='/gaodun/logs/'${script_name}'_rpc.log'
start() {
    if [ -f "${lockfile}" ]; then
	echo "server already start,pid:$pid"
        return 0
    fi
    nohup ${script_pwd}/${script_name}  rpc >> ${log} >> ${log} 2>&1 &
    retval=$?
    [ $retval -eq 0 ] && touch $lockfile
    echo "service start ok"
    return $retval
}
stop() {
    if [ -f "${lockfile}" ]; then
     rm -f $lockfile
    fi
    if [ -z "$pid" ]; then
      echo "not find program "
      return 0
    fi
    echo -n $"Stopping ...... "
    kill -USR1  $pid
    sleep 3
    pid=`ps aux|grep /gaodun/domain/|grep ucenter|grep -E 'rpc$'|grep -v grep|awk  '{print $2}'`
    while [ -n "$pid" ]
        do
            pid=`ps aux|grep /gaodun/domain/|grep ucenter|grep -E 'rpc$'|grep -v grep|awk  '{print $2}'`
            kill -9 $pid
            sleep 1
        done
    retval=$?
    [ $retval -eq 0 ] && rm -f $lockfile
    echo "kill program use signal 2,pid:$pid"
    return $retval
}

restart() {
    stop
    start
}

status() {
    if [ -z "$pid" ]; then
      echo "not find program"
   else
      echo "program is running,pid:$pid"
   fi
}
case "$1" in
    start)
        $1
        ;;
    stop)
        $1
        ;;
    restart)
        $1
        ;;
    status)
        $1
        ;;
    *)
        echo $"Usage: $0 {start|stop|status|restart}"
        exit 2
esac

