FROM golang:1.12.13 
ENV GO111MODULE "on"
ENV GOPROXY "https://goproxy.cn"
WORKDIR /src/Alertmanager-feishu-webhook
COPY . /src/Alertmanager-feishu-webhook
RUN make
FROM ubuntu 
COPY --from=0 /src/Alertmanager-feishu-webhook .
EXPOSE 8080
CMD ["./main", "-c", "conf/config.yaml"]
