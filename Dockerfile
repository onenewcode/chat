# 继承的基础镜像
# FROM golang:1.22
FROM alpine:3.19.1
# 指定接下来的工作路径
WORKDIR /chat
# 定义一个windos系统里的环境变量
#ENV VERSION=2.0.0	# optional
COPY . /chat

EXPOSE 8082

# 设置容器运行的命令
CMD ["/chat/cmd.sh"]
# CMD ["sleep", "10000000000000"]