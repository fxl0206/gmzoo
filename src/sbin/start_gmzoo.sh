for pid in `ps aux | grep "gmzoo" | grep -v grep | awk '{print $2}'`
do
kill -9 $pid
echo "$pid killed success!"
cd $GOPATH
echo "gmzoo conf/rgmzoo.ini  2>&1 | cronolog log/wss_%y%m%d.log &"
gmzoo conf/rgmzoo.ini  2>&1 | cronolog log/wss_%y%m%d.log &
done