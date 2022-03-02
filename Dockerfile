FROM golang:1.16-alpine as builder
WORKDIR /app
COPY ./ ./
ENV CGO_ENABLED=0
RUN go build -o main main.go 

FROM ubuntu:latest as runner
WORKDIR /app
# Install ffmpeg
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y --no-install-recommends ffmpeg
COPY --from=builder /app/main ./
COPY ./assets/ ./assets/
RUN mkdir output
EXPOSE 8080
ENTRYPOINT ["/bin/sh"]
CMD ["-c", "./main"]