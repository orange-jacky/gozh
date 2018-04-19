# gozh
golang 中文活跃社区


在gozh目录下操作

1.创建image
	docker-compose build

2.运行gozh服务
    docker-compose up 
    后台运行
    docker-compose up -d

3.停止服务
	docker-compose down

4.修改代码或者配置后,使修改生效
	先重新创建image
	docker-compose build
	再重启服务
	docker-compose up 


