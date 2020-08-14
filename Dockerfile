FROM golang:latest
LABEL maintainer="Derandi Hermanda"
WORKDIR /golang
COPY . .
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go build -o main .

ENTRYPOINT [ "./main" ]
EXPOSE 9000