FROM golang:1.14
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["main"]
