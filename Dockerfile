FROM golang:latest
LABEL maintainer="Derandi Hermanda"
LABEL version="1.0"
LABEL description="Golang Services."
WORKDIR /golang
COPY . .
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/handlers
RUN go build -o main .

ENTRYPOINT [ "./main" ]
EXPOSE 9000