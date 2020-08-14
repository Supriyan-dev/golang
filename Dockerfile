FROM golang:latest
LABEL maintainer="Derandi Hermanda"
WORKDIR /golang
COPY . .
RUN go get github.com/gin-gonic/gin
RUN go get gopkg.in/mgo.v2
RUN go get gopkg.in/mgo.v2/bson
RUN go build -o main ../golang
ENTRYPOINT [ "./main" ]
EXPOSE 3000