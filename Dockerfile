FROM golang:1.15-alpine
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o myapp .
CMD ["/app/myapp"]