package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	flags "github.com/jessevdk/go-flags"
)

var opts struct {
	Proxy      string `short:"p" long:"proxy" description:"Proxy server address"`
	BotID      string `short:"b" long:"bot_id" required:"yes" description:"Telegram bot token"`
	ChatID     int64  `short:"c" long:"chat_id" required:"yes" description:"Chat ID message should be sent to"`
	ParseMode  string `short:"m" long:"parse_mode" description:"Parse mode. Could be 'Markdown' or 'HTML'" choice:"Markdown" choice:"HTML" default:"Markdown"`
	Verbose    bool   `short:"v" long:"verbose" description:"Show verbose debug information"`
	Positional struct {
		Message []string
	} `positional-args:"yes" positional-arg-name:"message" required:"yes"`
}

func main() {
	_, err := flags.Parse(&opts)

	if err != nil {
		return
	}

	var bot *tgbotapi.BotAPI

	if len(opts.Proxy) > 0 {
		bot, err = tgbotapi.NewBotAPIWithClient(opts.BotID, newHTTPClient(opts.Proxy))
		if err != nil {
			log.Panic(err)
		}
	} else {
		bot, err = tgbotapi.NewBotAPI(opts.BotID)
		if err != nil {
			log.Panic(err)
		}
	}

	if opts.Verbose {
		bot.Debug = true
	}

	messageStr := strings.Replace(strings.Join(opts.Positional.Message, " "), "\\n", "\n", -1)

	msg := tgbotapi.NewMessage(opts.ChatID, messageStr)
	msg.ParseMode = opts.ParseMode

	bot.Send(msg)
}

func newHTTPClient(proxy string) *http.Client {
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		log.Println(err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	return &http.Client{Transport: transport}
}
