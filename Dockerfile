FROM golang:1.21.1-alpine
WORKDIR /app
COPY . .
RUN go build -o app

FROM alpine:3.18
WORKDIR /app
COPY --from=0 /app/app .
COPY config.yaml ./
EXPOSE 2020
CMD ./app start -p 2020
