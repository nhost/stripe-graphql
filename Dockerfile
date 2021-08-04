FROM golang:1.17rc2-alpine3.14 as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN go generate ./graph 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .
FROM scratch
COPY --from=builder /build/main /app/
WORKDIR /app
CMD ["./main"]
EXPOSE 8080