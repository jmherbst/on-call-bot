package main

import (
	"fmt"

	"github.com/cypress-io/on-call-bot/internal/config"
	"github.com/cypress-io/on-call-bot/internal/pagerduty"
	"github.com/cypress-io/on-call-bot/internal/slack"
)

func main() {
	cfg, cfgErr := config.GetConfig()
	if cfgErr != nil {
		fmt.Printf("%s", cfgErr)
		return
	}

	for escalationPolicy, slackUserGroupHandle := range cfg.Mappings {
		// Get on-call user's email address from PagerDuty
		userEmail := pagerduty.GetOncallUserEmail(cfg, escalationPolicy)

		// Update the Slack user group with on-call user
		slack.UpdateSlackUserGroup(cfg, slackUserGroupHandle, userEmail)

		fmt.Printf("Updated Slack user group(%s) with person on-call for PagerDuty policy('%s')\n", slackUserGroupHandle, escalationPolicy)
	}
}
