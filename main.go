package main

import (
	"log"
	"os"

	libgiphy "github.com/sanzaru/go-giphy"
	tele "gopkg.in/telebot.v3"
)

var (
	apiKey = os.Getenv("API_KEY")
	token  = os.Getenv("TELEGRAM_TOKEN")

	menu = &tele.ReplyMarkup{ResizeKeyboard: true}

	btnCats = menu.Text("😸 Gifs")
)

func main() {
	pref := tele.Settings{
		Token: token,
	}

	b, err := tele.NewBot(pref)
	if err != nil {
		log.Fatal(err)
		return
	}

	menu.Reply(
		menu.Row(btnCats),
	)

	b.Handle("/start", func(c tele.Context) error {
		return c.Send("Hi there, I'm a gif bot. Press the 😸 button to enjoy cat gifs!", menu)
	})

	b.Handle(&btnCats, func(c tele.Context) error {
		c.Notify("typing")
		gif := &tele.Animation{File: tele.FromURL(getCats())}
		return c.Send(gif.FileURL)
	})

	b.Start()
}

func getCats() string {
	giphy := libgiphy.NewGiphy(apiKey)

	gif, err := giphy.GetRandom("cats")
	if err != nil {
		log.Printf("failed to retrive gif, reason=%s\n", err)
	}

	return gif.Data.Bitly_gif_url
}
