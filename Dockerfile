# 继承的基础镜像
# FROM golang:1.22
FROM alpine:latest
# 指定接下来的工作路径
WORKDIR /
# 定义一个windos系统里的环境变量
#ENV VERSION=2.0.0	# optional

# 将容器 3000 端口暴露出来， 允许外部连接这个端口
EXPOSE 8082
# 拷贝当前文件夹下的所有文件到，工作目录
COPY . .

# 设置容器运行的命令
CMD ["/cmd.sh"]