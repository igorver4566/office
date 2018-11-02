package slack

import (
	"errors"

	"github.com/nlopes/slack"
)

var Slack struct {
	API *slack.Client
}

var key = "xoxp-433530174866-432806158864-469824086896-e3129c3b97cc478da8e2469b8b159963"

func init() {
	Slack.API = slack.New(key)
}

func NewGroup(name, purpose string) (string, error) {
	group, err := Slack.API.CreateGroup(name)
	if err != nil {
		return "", err
	}
	_, ok, err := Slack.API.InviteUserToGroup(group.ID, "UCST50Q0P")

	if err != nil || ok {
		return "", errors.New("Ошибка добавления пользователя")
	}

	_, ok, err = Slack.API.InviteUserToGroup(group.ID, "UCUDVJZFX")

	if err != nil || ok {
		return "", errors.New("Ошибка добавления бота")
	}

	_, err = Slack.API.SetGroupPurpose(group.ID, purpose)

	if err != nil {
		return "", errors.New("Ошибка создания описания")
	}

	return group.ID, nil
}
