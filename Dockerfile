FROM golang:1.21-alpine

WORKDIR /usr/src/app

RUN apk add tzdata
ENV TZ Asia/Bangkok

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . .

RUN go build -o ./output/go-backend
EXPOSE 3000
CMD [ "./output/go-backend"]
