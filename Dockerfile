FROM golang:latest
LABEL maintainer="Derandi Hermanda"
LABEL version="1.0"
LABEL description="Golang Services."
WORKDIR /golang
COPY . .
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/handlers
RUN go get github.com/rs/cors
RUN go get github.com/mervick/aes-everywhere/go/aes256
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/sendgrid/sendgrid-go
RUN go build -o main .

ENTRYPOINT [ "./main" ]
EXPOSE 9000