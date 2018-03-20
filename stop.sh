pid=`ps -ef | grep gozh_server | grep -v grep | awk '{print $2}'`
kill -2 $pid