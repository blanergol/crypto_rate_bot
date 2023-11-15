package helpers

import (
	"fmt"
	"time"

	"gopkg.in/telebot.v3"
)

func SendTelegramMessage(c telebot.Context, respList []string) error {
	var countWords int
	var resp string

	if len(respList) == 0 {
		return nil
	}

	a := fmt.Sprintf("<strong>Datetime</strong>: %s \n\n", time.Now().Format("Mon, 02 Jan 2006 15:04:05 MST"))
	respList = append([]string{a}, respList...)

	for i, r := range respList {
		countWords = countWords + len(r)
		if countWords > 400 {
			err := c.Send(resp, &telebot.SendOptions{ParseMode: telebot.ModeHTML, DisableWebPagePreview: true})
			if err != nil {
				return c.Send(err)
			}
			resp = ""
			countWords = 0
		} else {
			resp = resp + r
		}

		if countWords < 400 && len(respList) == i+1 {
			err := c.Send(resp, &telebot.SendOptions{ParseMode: telebot.ModeHTML, DisableWebPagePreview: true})
			if err != nil {
				return c.Send(err)
			}
		}
	}
	return nil
}
