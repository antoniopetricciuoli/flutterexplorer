package handlers

import (
	"flutterexplorerbot/messages"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
)

func HandleStart(client gobotapi.Client, update types.Message) {
	_, err := client.Invoke(&methods.SendMessage{
		ChatID:    update.Chat.ID,
		Text:      messages.StartMessage,
		ParseMode: "HTML",
		ReplyMarkup: &types.InlineKeyboardMarkup{
			InlineKeyboard: [][]types.InlineKeyboardButton{
				{
					{
						Text:              messages.TryIt,
						SwitchInlineQuery: "!w",
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
