# STAGE 1
FROM golang:alpine as build
RUN apk update && apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy
RUN go build -o medioker-bank
# yang diatas butuh golang untuk build nya aja. tapi sizenya besar
# jadi di stage 2 ini kita buat alpine supaya sizenya lebih kecil
# STAGE 2
FROM alpine
WORKDIR /app
# karena diatas as buil maka fromnya from build, kalau diatas as tokonyadia maka dibawah from tokonyadia
COPY --from=build /app/medioker-bank /app/medioker-bank
ENTRYPOINT [ "/app/medioker-bank" ]
