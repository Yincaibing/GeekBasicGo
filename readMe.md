echo "# GeekBasicGo" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/Yincaibing/GeekBasicGo.git
git push -u origin main

1、启动 mysql：
docker run --name ycbmysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=ycbpassword -d mysql:latest

2、链接mysql:
docker run -it --network some-network --rm mysql mysql -hsome-mysql -uYCBmysql-user -p

3、进入 mysql容器命令
docker exec -it mysql bash

4、查看 docker 容器 mysql 日志
docker logs YCB-mysql





# 将应用打包成 docker镜像
1、cd webook
2、首先在本地完成编译，生成一个可在 Linux 平台上执行的 webook 可执行文件：GOOS=linux GOARCH=arm go build -o webook .
3、docker build -t flycash/webook:v0.0.1 .

