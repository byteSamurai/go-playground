FROM golang:1.16-alpine

RUN apk update && apk add git make uglify-js
  
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...
RUN make -C static

CMD ["go-webapp"]

EXPOSE 3000
