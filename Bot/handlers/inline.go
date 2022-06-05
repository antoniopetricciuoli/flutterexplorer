package handlers

import (
	"flutterexplorerbot/messages"
	"flutterexplorerbot/utils"
	"fmt"
	"github.com/Squirrel-Network/gobotapi"
	"github.com/Squirrel-Network/gobotapi/methods"
	"github.com/Squirrel-Network/gobotapi/types"
	"html"
	"strings"
)

func HandleInline(client gobotapi.Client, update types.InlineQuery) {

	if update.Query != "" {
		if strings.HasPrefix(update.Query, "!w") {
			var inlineQueryResult []types.InlineQueryResult
			query := strings.Join(strings.Split(update.Query, " ")[1:], " ")
			var switchPmText string
			widgets, err := utils.GetWidgets(query)

			if len(widgets) > 0 {
				switchPmText = messages.Search
				for i, widget := range widgets {
					if i >= 50 {
						break
					}

					inlineQueryResult = append(inlineQueryResult, types.InlineQueryResultArticle{
						ID:          fmt.Sprintf("%d", i),
						Title:       widget.Name,
						Description: widget.Description,
						InputMessageContent: types.InputTextMessageContent{
							MessageText:           fmt.Sprintf(messages.WidgetResult, html.EscapeString(widget.Name), html.EscapeString(widget.Description), html.EscapeString(widget.Link)),
							DisableWebPagePreview: true,
							ParseMode:             "HTML",
						},
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
				}
			} else {
				switchPmText = fmt.Sprintf(messages.NoResult, query)
			}

			_, err = client.Invoke(&methods.AnswerInlineQuery{
				InlineQueryID:     update.ID,
				SwitchPmText:      switchPmText,
				SwitchPmParameter: "start",
				Results:           inlineQueryResult,
			})

			if err != nil {
				panic(err)
			}
		} else if strings.HasPrefix(update.Query, "!p") {
			var inlineQueryResult []types.InlineQueryResult
			var switchPmText string
			var link string
			query := strings.Join(strings.Split(update.Query, " ")[1:], " ")
			packages, err := utils.GetPackages(query)

			if len(packages) > 0 {
				switchPmText = messages.Search
				for i, p := range packages {
					if i >= 50 {
						break
					}

					if strings.HasPrefix(p.Link, "https") {
						link = p.Link
					} else {
						link = fmt.Sprintf("https://pub.dev%s", p.Link)
					}

					inlineQueryResult = append(inlineQueryResult, types.InlineQueryResultArticle{
						ID:          fmt.Sprintf("%d", i),
						Title:       p.Name,
						Description: p.Description,
						InputMessageContent: types.InputTextMessageContent{
							MessageText:           fmt.Sprintf(messages.PackageResult, html.EscapeString(p.Name), html.EscapeString(p.Description), html.EscapeString(p.Metadata.Version), p.Scores.PubPoints, p.Scores.Likes),
							DisableWebPagePreview: true,
							ParseMode:             "HTML",
						},
						ReplyMarkup: &types.InlineKeyboardMarkup{
							InlineKeyboard: [][]types.InlineKeyboardButton{
								{
									{
										Text: "🔗 Pub.dev",
										URL:  html.EscapeString(link),
									},
								},
								{
									{
										Text:              messages.TryIt,
										SwitchInlineQuery: "!p",
									},
								},
							},
						},
					})
				}
			} else {
				switchPmText = fmt.Sprintf(messages.NoResult, query)
			}

			_, err = client.Invoke(&methods.AnswerInlineQuery{
				InlineQueryID:     update.ID,
				SwitchPmText:      switchPmText,
				SwitchPmParameter: "start",
				Results:           inlineQueryResult,
			})

			if err != nil {
				panic(err)
			}
		}

	} else {
		_, err := client.Invoke(&methods.AnswerInlineQuery{
			InlineQueryID:     update.ID,
			SwitchPmText:      messages.Search,
			SwitchPmParameter: "start",
			CacheTime:         1,
			IsPersonal:        true,
			Results: []types.InlineQueryResult{
				types.InlineQueryResultArticle{
					ID:          "flutter",
					Title:       "Flutter",
					Description: messages.Flutter,
					ThumbURL:    "https://docs.flutter.dev/assets/images/docs/catalog-widget-placeholder.png",
					InputMessageContent: types.InputTextMessageContent{
						MessageText: messages.Flutter,
					},
					ReplyMarkup: &types.InlineKeyboardMarkup{
						InlineKeyboard: [][]types.InlineKeyboardButton{
							{
								{
									Text:              messages.TryIt,
									SwitchInlineQuery: "@FlutterExplorerBot",
								},
							},
						},
					},
				},
				types.InlineQueryResultArticle{
					ID:          "info",
					Title:       "Info sul bot",
					Description: "Informazioni sul bot",
					InputMessageContent: types.InputTextMessageContent{
						MessageText: messages.StartMessage,
						ParseMode:   "HTML",
					},
					ThumbURL: "https://i.imgur.com/zRglRz3.png",
					ReplyMarkup: &types.InlineKeyboardMarkup{
						InlineKeyboard: [][]types.InlineKeyboardButton{
							{
								{
									Text:              messages.TryIt,
									SwitchInlineQuery: "@FlutterExplorerBot",
								},
							},
						},
					},
				},
			},
		})

		if err != nil {
			panic(err)
		}
	}
}
