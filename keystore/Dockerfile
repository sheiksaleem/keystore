FROM golang:1.16
WORKDIR /app
COPY . .
RUN go mod download

EXPOSE 8000
# Install server application

RUN CGO_ENABLED=0 GOOS=linux go build -a -o main .

CMD ["/main"]
