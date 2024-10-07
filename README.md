# NMU Schedule Bot

## About
This is a Telegram bot that allows to view schedule
of the [Bogomolets National Medical University](https://en.wikipedia.org/wiki/Bogomolets_National_Medical_University)
in a simple format.

It uses undocumented API at https://vnz.nmuofficial.com/WebDk so no guarantees are provided.

## Deployment

The bot is written in Go. In order to build it, run:
```shell
go build
```
and then run the resulting binary.

It can also be deployed using docker:
```shell
docker compose up
```

Configuration is provided using environment variables or `.env` file. (Look at provided `.env.example`).
[Telegram bot token](https://core.telegram.org/bots/tutorial), username, password and API version have to be specified.
Update period can also be configured. \
API version can be found at https://vnz.nmuofficial.com/profile.
