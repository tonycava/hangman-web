FROM golang:1.19-alpine as builder

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o back.exe .

FROM alpine

WORKDIR /app

COPY --from=builder /app .

EXPOSE 8000

CMD [ "./back.exe" ]