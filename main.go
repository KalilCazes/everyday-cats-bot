package main

import (
	"fmt"
	"log"
	"os"

	libgiphy "github.com/sanzaru/go-giphy"
	tele "gopkg.in/telebot.v3"
)

var (
	apiKey = os.Getenv("API_KEY")
	token  = os.Getenv("TELEGRAM_TOKEN")
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

	b.Handle("/cats", func(c tele.Context) error {
		gif := &tele.Animation{File: tele.FromURL(getCats())}
		return c.Send(gif.FileURL)
	})

	b.Start()

}

func getCats() string {
	giphy := libgiphy.NewGiphy(apiKey)

	gif, err := giphy.GetRandom("cats")
	if err != nil {
		fmt.Println("error:", err)
	}

	return gif.Data.Bitly_gif_url
}
