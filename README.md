golang-monitor-zookeeper
===================
golang monitor zookeeper server status width  websocket

install steps
===============================
1. git clone https://github.com/fxl0206/gmzoo
2. cd gmzoo/src/sbin
3. ./install.sh
4. cd ../../sbin
5. ./build.sh
6. ./start_gmzoo.sh

引用项目
=====================================
https://github.com/samuel/go-zookeeper
https://code.google.com/p/go.net/websocket
https://github.com/widuu/conf

已实现功能
==================================================
1.Golang提供Http Web服务，生成简易监控页面

2.Golang提供WebSocket服务，实时推送zookeeper最新节点信息到监控页
待实现功能
==================================================
1.监控zookeeper节点值变化，并通过webSocket实时反馈给前台监控页

...未完待续....
