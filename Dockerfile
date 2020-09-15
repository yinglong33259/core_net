FROM  golang:1.15.2 AS build

# 容器环境变量添加，会覆盖默认的变量值
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

# 设置工作区
WORKDIR /opt/corenet

# 把全部文件添加到/opt/src目录
ADD . .

# 编译：把cmd/main.go编译成可执行的二进制文件，命名为corenet
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o corenet kube/cmd/main.go

FROM scratch as prod
# 在build阶段复制时区到
COPY --from=build /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
# 在build阶段复制可执行的go二进制文件app
COPY --from=build /go/release/corenet /
# 在build阶段复制配置文件
#COPY --from=build /go/release/config ./config
COPY ./10-corenet.conf /opt/corenet/10-corenet.conf

COPY ./run.sh /opt/corenet/run.sh
RUN chmod 755 /opt/corenet/run.sh
# 启动服务
CMD ["/run.sh"]
