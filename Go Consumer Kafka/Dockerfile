FROM golang:1.18-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /consumer


CMD["/consumer"]
#FROM gcr.io/distroless/static
#ENV MONGO_ADDRESS="mongodb://root:root@172.17.0.2:27017"
#ENV REDIS_ADDRESS="redis://34.136.166.39:6379"
#COPY --from=build-env /go/bin/app /
#CMD ["/app"]