FROM golang:alpine
WORKDIR /src
COPY . .
RUN go get
EXPOSE 8080
CMD go run main.go