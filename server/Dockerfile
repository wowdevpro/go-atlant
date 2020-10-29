FROM golang:alpine
WORKDIR /app
COPY /server .
COPY /cert /app/cert
#RUN cd /app && ls -la
RUN apk add git

RUN go get -v go.mongodb.org/mongo-driver/bson
RUN go get -v go.mongodb.org/mongo-driver/mongo
RUN go get -v google.golang.org/grpc/grpclog
RUN go get -v github.com/wowdevpro/go-atlant/proto

RUN go build -o main .

EXPOSE 5300

CMD ["./main"]