FROM golang:1.23-alpine
WORKDIR /usr/src/app
ENV GOPROXY=https://goproxy.cn,direct
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o myapp .
CMD ["./myapp"]
