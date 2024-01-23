FROM golang:1.21.1-alpine

WORKDIR /app
COPY . .
RUN go build -o app
EXPOSE 2020
CMD ./app start -p 2020
