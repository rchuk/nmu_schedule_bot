FROM golang:1.23-alpine AS build

WORKDIR /app

COPY . .

RUN go build

FROM alpine:3.20.3 AS release

WORKDIR /app

COPY --from=build /app/nmu_schedule_bot nmu_schedule_bot

RUN apk add tzdata

ARG UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    appuser
USER appuser

ENTRYPOINT ["/app/nmu_schedule_bot"]
