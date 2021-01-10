FROM golang:alpine
WORKDIR /src
COPY . .
RUN go get
EXPOSE 9090
CMD go run main.go