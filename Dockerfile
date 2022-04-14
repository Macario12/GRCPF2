FROM golang:1.18
WORKDIR /subscriberContainer
COPY . .
RUN go mod download

EXPOSE 50051
CMD ["go","run","Subscriber/subscriber.go"]