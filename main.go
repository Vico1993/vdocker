package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Vico1993/vdocker/tools"
	"github.com/yanzay/tbot"
)

func main() {
	conf, err := tools.LoadConfig()
	if err != nil {
		fmt.Println("ICI")
		log.Fatal(err.Error())
	}

	fmt.Println(conf.TokenTelegram)

	bot := tbot.New(conf.TokenTelegram)
	c := bot.Client()

	bot.HandleMessage("/docker", func(m *tbot.Message) {
		c.SendChatAction(m.Chat.ID, tbot.ActionTyping)
		time.Sleep(1 * time.Second)
		c.SendMessage(m.Chat.ID, "hello!")
	})

	err = bot.Start()
	if err != nil {
		log.Fatal(err)
	}
}
