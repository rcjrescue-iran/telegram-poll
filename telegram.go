package main

import (
	"log"
	"strconv"
	"time"

	"github.com/tucnak/telebot"
)

const (
	//token is telegram robot token
	token = "328416454:AAHs2sRT4gBjT44U23wa5nQRsRdbrQ4tzlQ" // will be revoke

	//messageID from telegram channel - if equal to zero - telebot will post new result and fetch new messageID
	messageID = 0

	//defaultMessage is a text message that will send to users if new text message received
	defaultMessage = `این سرویس پس از مسابقات غیر فعال شده است`

	adminID = 83919508
)

var (
	bot *telebot.Bot

	oldMsg string

	telegramChat = telebot.Chat{Type: "channel", Username: "rcjrescue"}

	sendOption = telebot.SendOptions{
		ReplyMarkup: telebot.ReplyMarkup{
			InlineKeyboard: [][]telebot.KeyboardButton{
				[]telebot.KeyboardButton{
					telebot.KeyboardButton{
						URL:  "https://rcjrescue.zibaei.net/",
						Text: "شرکت در نظر سنجی",
					},
				},
			},
			ResizeKeyboard: true,
		},
	}
)

func initTelegram() {
	var err error
	bot, err = telebot.NewBot(token)
	if err != nil {
		log.Println(err)
		return
	}
}

func listenToMessages() {
	messages := make(chan telebot.Message)
	bot.Listen(messages, 1*time.Second)

	for message := range messages {
		bot.SendMessage(message.Sender, defaultMessage, nil)
	}
}

func autoUpdate() {
	text := "نتایج نظر سنجی" + "\n\n"
	surveys := GetTotalSurveys()
	length := len(surveys)

	result := make([]Survey, 4)

	for _, s := range surveys {
		// sorry for bad coding :| , yes I know , I can make it better but I haven't many times.
		if s.Level == 1 {
			result[1].Level++
		} else if s.Level == 2 {
			result[2].Level++
		} else if s.Level == 3 {
			result[3].Level++
		} else {
			result[0].Level++
		}

		if s.Refree == 1 {
			result[1].Refree++
		} else if s.Refree == 2 {
			result[2].Refree++
		} else if s.Refree == 3 {
			result[3].Refree++
		} else {
			result[0].Refree++
		}

		if s.Quality == 1 {
			result[1].Quality++
		} else if s.Quality == 2 {
			result[2].Quality++
		} else if s.Quality == 3 {
			result[3].Quality++
		} else {
			result[0].Quality++
		}

		if s.Timing == 1 {
			result[1].Timing++
		} else if s.Timing == 2 {
			result[2].Timing++
		} else if s.Timing == 3 {
			result[3].Timing++
		} else {
			result[0].Timing++
		}

		if s.Partition == 1 {
			result[1].Partition++
		} else if s.Partition == 2 {
			result[2].Partition++
		} else if s.Partition == 3 {
			result[3].Partition++
		} else {
			result[0].Partition++
		}

		if s.Proportionality == 1 {
			result[1].Proportionality++
		} else if s.Proportionality == 2 {
			result[2].Proportionality++
		} else if s.Proportionality == 3 {
			result[3].Proportionality++
		} else {
			result[0].Proportionality++
		}

		if s.Idea == 1 {
			result[1].Idea++
		} else if s.Idea == 2 {
			result[2].Idea++
		} else if s.Idea == 3 {
			result[3].Idea++
		} else {
			result[0].Idea++
		}

		if s.Morality == 1 {
			result[1].Morality++
		} else if s.Morality == 2 {
			result[2].Morality++
		} else if s.Morality == 3 {
			result[3].Morality++
		} else {
			result[0].Morality++
		}

		if s.Broadcast == 1 {
			result[1].Broadcast++
		} else if s.Broadcast == 2 {
			result[2].Broadcast++
		} else if s.Broadcast == 3 {
			result[3].Broadcast++
		} else {
			result[0].Broadcast++
		}

		if s.Points == 1 {
			result[1].Points++
		} else if s.Points == 2 {
			result[2].Points++
		} else if s.Points == 3 {
			result[3].Points++
		} else {
			result[0].Points++
		}
	}

	text += "⭐️سطح برگزاری مسابقات : \n"
	text += "خوب : " + getResult(result[1].Level, length) + "\n"
	text += "متوسط : " + getResult(result[2].Level, length) + "\n"
	text += "بد : " + getResult(result[3].Level, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Level, length) + "\n"

	text += "\n"

	text += "⭐️کیفیت داوری : \n"
	text += "خوب : " + getResult(result[1].Refree, length) + "\n"
	text += "متوسط : " + getResult(result[2].Refree, length) + "\n"
	text += "بد : " + getResult(result[3].Refree, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Refree, length) + "\n"

	text += "\n"

	text += "⭐️تناسب سختی مراحل : \n"
	text += "خوب : " + getResult(result[1].Proportionality, length) + "\n"
	text += "متوسط : " + getResult(result[2].Proportionality, length) + "\n"
	text += "بد : " + getResult(result[3].Proportionality, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Proportionality, length) + "\n"

	text += "\n"

	text += "⭐️زمان بندی : \n"
	text += "خوب : " + getResult(result[1].Timing, length) + "\n"
	text += "متوسط : " + getResult(result[2].Timing, length) + "\n"
	text += "بد : " + getResult(result[3].Timing, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Timing, length) + "\n"

	text += "\n"

	text += "⭐️برخورد کمیته با تیم ها : \n"
	text += "خوب : " + getResult(result[1].Morality, length) + "\n"
	text += "متوسط : " + getResult(result[2].Morality, length) + "\n"
	text += "بد : " + getResult(result[3].Morality, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Morality, length) + "\n"

	text += "\n"

	text += "⭐️ایده و برگزاری سوپرتیم : \n"
	text += "خوب : " + getResult(result[1].Idea, length) + "\n"
	text += "متوسط : " + getResult(result[2].Idea, length) + "\n"
	text += "بد : " + getResult(result[3].Idea, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Idea, length) + "\n"

	text += "\n"

	text += "⭐️طرح و کیفیت مسابقات : \n"
	text += "خوب : " + getResult(result[1].Quality, length) + "\n"
	text += "متوسط : " + getResult(result[2].Quality, length) + "\n"
	text += "بد : " + getResult(result[3].Quality, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Quality, length) + "\n"

	text += "\n"

	text += "⭐️تاثیر کوتاهی پارتیشن ها : \n"
	text += "خوب : " + getResult(result[1].Partition, length) + "\n"
	text += "متوسط : " + getResult(result[2].Partition, length) + "\n"
	text += "بد : " + getResult(result[3].Partition, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Partition, length) + "\n"

	text += "\n"

	text += "⭐️اطلاع رسانی : \n"
	text += "خوب : " + getResult(result[1].Broadcast, length) + "\n"
	text += "متوسط : " + getResult(result[2].Broadcast, length) + "\n"
	text += "بد : " + getResult(result[3].Broadcast, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Broadcast, length) + "\n"

	text += "\n"

	text += "⭐️نحوه امتیاز بندی : \n"
	text += "خوب : " + getResult(result[1].Points, length) + "\n"
	text += "متوسط : " + getResult(result[2].Points, length) + "\n"
	text += "بد : " + getResult(result[3].Points, length) + "\n"
	text += "پر نشده : " + getResult(result[0].Points, length) + "\n"

	text += "\n."

	var err error
	msg := telebot.Message{}

	if messageID == 0 {
		msg, err = bot.SendMessage(telegramChat, text, &sendOption)
		if err != nil {
			log.Println(err)
		}
	} else {
		msg.ID = messageID
	}

	msg.ID = 4

	if oldMsg != text {

		_, err = bot.EditMessageText(telegramChat, msg, text, &sendOption)
		if err != nil {
			log.Println(err)
		}
	}
	oldMsg = text

	time.AfterFunc(10*time.Second, autoUpdate)
}

func getResult(input, total int) string {
	return strconv.FormatFloat(float64(float32(input)/float32(total))*100, 'f', 2, 32)
}
