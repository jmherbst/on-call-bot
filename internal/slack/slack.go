package slack

import (
	"fmt"

	"github.com/cypress-io/on-call-bot/internal/config"
	"github.com/slack-go/slack"
)

func getClient(config config.Config) *slack.Client {
	return slack.New(config.SlackToken)
}

// TODO Fix panic errs

func getUserID(client *slack.Client, email string) string {

	slackUser, err := client.GetUserByEmail(email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User %s\n", slackUser.Name)

	return slackUser.ID

}

func getGroupID(client *slack.Client, groupHandle string) string {
	groups, err := client.GetUserGroups(slack.GetUserGroupsOptionIncludeCount(true))
	if err != nil {
		panic(err)
	}

	var groupID string
	for _, group := range groups {
		if group.Handle == groupHandle {
			fmt.Printf("ID: %s, Name: %s\n", group.ID, group.Name)
			groupID = group.ID
		}
	}

	return groupID
}

func updateUserGroup(client *slack.Client, groupID string, userID string) {
	_, updateGroupErr := client.UpdateUserGroupMembers(groupID, userID)
	if updateGroupErr != nil {
		fmt.Printf("ERR %s\n", updateGroupErr)

		panic(updateGroupErr)
	}
}

func UpdateSlackUserGroup(config config.Config, groupHandle string, email string) {
	client := getClient(config)

	// Get the slack user ID matching the on-call user's email address
	slackUserID := getUserID(client, email)

	// Get the slack user group ID matching the group handle
	slackGroupID := getGroupID(client, groupHandle)

	// Update the slack user group with on-call user
	updateUserGroup(client, slackGroupID, slackUserID)
}
