echo "# GeekBasicGo" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/Yincaibing/GeekBasicGo.git
git push -u origin main

1、启动 mysql：
docker run --name ycbmysql -p 3306:3306 -e MYSQL_ROOT_PASSWORD=ycbpassword -d --network=mynetwork mysql:tag  这个network是 Docker 网络。例如运行 docker network create mynetwork，其中 "mynetwork" 是你自定义的网络名称。

2、链接mysql:
docker run -it --network some-network --rm mysql mysql -hsome-mysql -uYCBmysql-user -p

3、进入 mysql容器命令
docker exec -it mysql bash

4、查看 docker 容器 mysql 日志
docker logs YCB-mysql




# 下面是将 golang 应用打包成 Docker 镜像并部署到 Kubernetes 的整个过程
# 1、如何将应用打包成 docker镜像:https://docs.docker.com/build/guide/intro/
1、cd webook
2、首先在本地完成编译，生成一个可在 Linux 平台上执行的 webook 可执行文件：GOOS=linux GOARCH=arm go build -o webook .
3、docker build -t  webook .   这里要注意 dockerfile和 main.go和 go.sum要在同一个路径下
4、docker run -d -p 8080:8080 --network=mynetwork webook 运行镜像
5、docker tag local-image:latest  gin_practice:v1.0 给镜像打标签
6、    就可以用镜像标签推送：  docker push gin_practice:v1.0

# 2、k8s上部署我的应用
前提：在启动 service之前，需要确认 docker和 minikube本地docker集群是否启动，命令：
docker version  
minikube start
minikube status

1、创建一个 deployment.yaml文件
运行命令：kubectl apply -f deployment.yaml
2、如果你还需要对外暴露这个应用，你可以创建一个 service
运行命令：kubectl apply -f service.yaml
3、minikube service webook
