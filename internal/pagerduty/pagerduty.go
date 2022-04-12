package pagerduty

import (
	"fmt"

	"github.com/PagerDuty/go-pagerduty"
	"github.com/cypress-io/on-call-bot/internal/config"
)

func getClient(config config.Config) *pagerduty.Client {
	return pagerduty.NewClient(config.PagerDutyToken)
}

// TODO Fix panic errs

func getPolicyID(client *pagerduty.Client, escalationPolicy string) string {
	var opts pagerduty.ListEscalationPoliciesOptions
	opts.Query = escalationPolicy
	eps, err := client.ListEscalationPolicies(opts)
	if err != nil {
		panic(err)
	}

	// ensure eps is only 1 item
	if len(eps.EscalationPolicies) != 1 {
		panic("More than 1 escalation policy matches config defined policy " + fmt.Sprintf("%s -- Found: #%d", escalationPolicy, len(eps.EscalationPolicies)))
	}

	return eps.EscalationPolicies[0].ID
}

func getOncallUserID(client *pagerduty.Client, policyID string) string {
	// Get the oncalls for escalation policy
	var onCallOpts pagerduty.ListOnCallOptions
	onCallOpts.EscalationPolicyIDs = []string{policyID}
	onCalls, err := client.ListOnCalls(onCallOpts)
	if err != nil {
		panic(err)
	}

	// Get the primary escalation oncall object
	var userID string
	for _, oncall := range onCalls.OnCalls {
		if oncall.EscalationLevel == 1 {
			userID = oncall.User.ID
			fmt.Printf("Primary escalation User ID found: %s\n", oncall.User.ID)
		}
	}

	return userID
}

func getUserEmailFromID(client *pagerduty.Client, userID string) string {
	var userOpts pagerduty.GetUserOptions
	pdUser, err := client.GetUser(userID, userOpts)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", pdUser.Email)

	return pdUser.Email
}

func GetOncallUserEmail(config config.Config, escalationPolicy string) string {
	client := getClient(config)

	// Get the policy ID matching the escalation policy
	policyID := getPolicyID(client, escalationPolicy)

	// Get primary on-call user ID
	userID := getOncallUserID(client, policyID)

	return getUserEmailFromID(client, userID)
}
