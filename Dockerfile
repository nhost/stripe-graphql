FROM golang as builder

WORKDIR /build 
ADD . .

RUN go generate ./graph

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o main .


FROM scratch

WORKDIR /app
COPY --from=builder /build/main /app/

CMD ["./main"]

EXPOSE 8080