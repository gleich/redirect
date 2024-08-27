# syntax=docker/dockerfile:1
FROM golang:1.23 AS build

WORKDIR /src
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/redirect ./main.go

FROM alpine:3.20.2

RUN apk update && apk add --no-cache ca-certificates=20240705-r0

WORKDIR /src
COPY --from=build /bin/redirect /bin/redirect

CMD ["/bin/redirect"]
