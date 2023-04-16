FROM golang:alpine as builder
LABEL "owner"="sobhankazemi80@gmail.com"
RUN mkdir app
WORKDIR /app
COPY . ./
RUN go build -o main ./...

FROM alpine 
RUN mkdir app
WORKDIR /app
COPY --from=builder /app/* .
CMD ["./main"]
