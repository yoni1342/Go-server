FROM golang:alpine

COPY ./ /server

WORKDIR /server

# ENV GIN_MODE=${GIN_MODE}

RUN go mod download

RUN go mod tidy

RUN go build -o server

RUN chmod +x /server/server

EXPOSE 8082

CMD ["/server/server"]