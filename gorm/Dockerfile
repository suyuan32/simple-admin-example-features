FROM golang:1.20.2-alpine3.17 as builder

# Define the project name | 定义项目名称
ARG PROJECT=example

WORKDIR /home
COPY . .

RUN go env -w GO111MODULE=on \
    && go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w CGO_ENABLED=0 \
    && go env \
    && go mod tidy \
    && go build -ldflags="-s -w" -o /home/${PROJECT}_rpc ${PROJECT}.go

FROM alpine:latest

# Define the project name | 定义项目名称
ARG PROJECT=example
# Define the config file name | 定义配置文件名
ARG CONFIG_FILE=example.yaml
# Define the author | 定义作者
ARG AUTHOR=RyanSU@yuansu.china.work@gmail.com

LABEL MAINTAINER=${AUTHOR}

WORKDIR /home
ENV PROJECT=${PROJECT}
ENV CONFIG_FILE=${CONFIG_FILE}

COPY --from=builder /home/${PROJECT}_rpc ./
COPY --from=builder /home/etc/${CONFIG_FILE} ./etc/

EXPOSE 9100
ENTRYPOINT ./${PROJECT}_rpc -f etc/${CONFIG_FILE}