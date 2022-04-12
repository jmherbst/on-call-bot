# On-Call Bot
## Overview
This bot is used to map on-call users from PagerDuty to associated slack groups to make it easier to find and contact on-call teammates.

## Automation
This bot runs every hour via a scheduled GitHub Action workflow named [`cron`](https://github.com/cypress-io/on-call-bot/actions/workflows/cron.yml).

## Adding Your Team to the Bot
If you'd like the bot to automate syncing your PagerDuty Escalation to your slack group, you can add those to the `./config.yml` mappings.

```yaml
  <pagerduty escalation policy name>: <team-name>-oncall
```

## Running the Bot Locally
`go run .` or via Docker with:
```bash
docker build -t on-call-bot .
docker run --env PAGERDUTY_TOKEN --env SLACK_TOKEN --rm on-call-bot:latest
```

## Needs
- Tests
