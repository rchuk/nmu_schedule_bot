FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go build

FROM alpine:3.20.3 AS release

WORKDIR /app

COPY --from=build /app/nmu_schedule_bot nmu_schedule_bot

RUN addgroup -S group && adduser -S user -G group
USER user

ENTRYPOINT ["/nmu_schedule_bot"]
