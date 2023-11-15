package helpers

import "gopkg.in/telebot.v3"

func SendTelegramWith400Words(c telebot.Context, respList []string) error {
	var countWords int
	var resp string
	for i, r := range respList {
		countWords = countWords + len(r)
		if countWords > 400 {
			err := c.Send(resp)
			if err != nil {
				return c.Send(err)
			}
			resp = ""
			countWords = 0
		} else {
			resp = resp + r
		}

		if countWords < 400 && len(respList) == i+1 {
			err := c.Send(resp)
			if err != nil {
				return c.Send(err)
			}
		}
	}
	return nil
}
