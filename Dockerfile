FROM golang:alpine as builder
ENV GO111MODULE=on
#RUN mkdir /app
#ADD . /app
WORKDIR /app
COPY go.mod ./
RUN go mod download
RUN go clean --modcache
COPY . .
RUN go build -o main

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/config/.env .
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]



#FROM golang:1.16-alpine AS builder
#
#RUN mkdir /app
#ADD . /app
#WORKDIR /app
#RUN go clean --modcache
#RUN go build -o main
## EXPOSE 8080
## CMD ["/app/main"]
#
## stage 2
#FROM alpine:3.14
#WORKDIR /root/
#COPY --from=builder /app/config/.env /config/
#COPY --from=builder /app/main .
#EXPOSE 8000
#CMD ["./main"]