FROM golang:1.12.9

WORKDIR /usr/src/minio-admin-platform
COPY . .

ENV GOPROXY      "https://goproxy.cn"
RUN go mod tidy
RUN go build -o minio-admin-platform main.go

FROM centos:7.4.1708
ENV GIN_MODE release
ARG ACCESSKEY
ARG SECRETKEY
ARG MINIO_ENDPOINT
ENV ACCESSKEY=$ACCESSKEY
ENV SECRETKEY=$SECRETKEY
ENV MINIO_ENDPOINT=$MINIO_ENDPOINT
WORKDIR /usr/src/app
COPY --from=0 /usr/src/minio-admin-platform/minio-admin-platform .
COPY --from=0 /usr/src/minio-admin-platform/conf ./conf
RUN chmod 777 ./minio-admin-platform

EXPOSE 9000
CMD ["./minio-admin-platform"]
