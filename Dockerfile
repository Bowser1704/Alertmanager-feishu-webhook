FROM golang:1.12.13 
ENV GO111MODULE "on"
ENV GOPROXY "https://mirrors.aliyun.com/goproxy/"
WORKDIR /src/Alertmanager-feishu-webhook
COPY . /src/Alertmanager-feishu-webhook
RUN make
EXPOSE 8080
CMD ["./main", "-c", "conf/config.yaml"]
