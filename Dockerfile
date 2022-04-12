# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

ENV PAGERDUTY_TOKEN=$PAGERDUTY_TOKEN
ENV SLACK_TOKEN=$SLACK_TOKEN

WORKDIR /app

COPY . /app

RUN go build -buildvcs=false -o on-call-bot


CMD [ "/app/on-call-bot" ]
