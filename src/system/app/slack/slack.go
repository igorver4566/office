package slack

import (
	"errors"

	"github.com/nlopes/slack"
)

var Slack struct {
	API *slack.Client
}

var key = ""

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

func PostMessage(channel, message, username string) (string, error) {
	channelID, _, err := Slack.API.PostMessage(channel, slack.MsgOptionText(message, false), slack.MsgOptionUserName(username))
	if err != nil {
		return "", errors.New("Ошибка при отправке сообщения")
	}
	return channelID, nil
}

func ChatHistory(name string) ([]slack.Message, error) {

	details := slack.HistoryParameters{
		Count:     1000,
		Inclusive: false}
	history, err := Slack.API.GetGroupHistory(name, details)
	msg := history.Messages
	if err != nil {
		return nil, errors.New("Ошибка при получении сообщений")
	}
	return msg, nil
}
