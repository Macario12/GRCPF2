FROM golang:1.18
WORKDIR /serverContainer
COPY . .
RUN go mod download

EXPOSE 50051
CMD ["go","run","Server/main.go"]