# On-Call Bot
### 📒 This is a proof-of-concept 📒

This bot is used to map on-call users from PagerDuty to associated slack groups to make it easier to find and contact on-call teammates.

## Config
The bot requires PD-Escalation Policy --to-- Slack Group mappings to be configured in `./config.yml`.

It also requires both a `PAGERDUTY_TOKEN` and `SLACK_TOKEN` env vars to be set with properly scoped tokens.

## Running the Bot
`go run .` (should probably containerize this)

## Needs
[] Tests 

[] Dockerizing

[] Scheduled execution (ECS? SQS?) - We should run this on an hourly cadence so that it keeps slack groups closely aligned with on-call users.


