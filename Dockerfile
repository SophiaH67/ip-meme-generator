FROM golang:1.16-alpine as builder
WORKDIR /app
COPY ./ ./ 
RUN go build -o main main.go 

FROM jrottenberg/ffmpeg:ubuntu as runner
WORKDIR /app
COPY --from=builder /app/main ./
EXPOSE 8080
CMD ["/bin/sh"]