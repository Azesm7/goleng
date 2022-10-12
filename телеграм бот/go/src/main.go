package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"

	tgbotapi "gopkg.in/telegram-bot-api.v4"
)

// go get gopkg.in/telegram-bot-api.v4
// регистрация бота BotFather
// ngrok heroku
const (
	botToken   = "5679460063:AAGQOf7gsFwIczxlNy89_30I2BMLPqUBuG0" //  токен бот (регистрация бота)
	WebhookURL = "https://2d54-89-250-30-73.eu.ngrok.io"          // выложить бот в интернет
)

var rss = map[string]string{ // список rss
	"Harb": "https://habrahabr.ru/rss/best/",
}

type RSS struct { // создание структуры xml
	Items []Item `xml:"channel>item"`
}
type Item struct { //создание структуры xml
	URL   string `xml:"guid"`
	Title string `xml:"title"`
}

func main() {
	bot, err := tgbotapi.NewBotAPI(botToken) //создание объекта бота в скобках указывем токен бота
	if err != nil {
		panic(err)
	}

	fmt.Printf("Authorized on account %s\n", bot.Self.UserName) //инцелизируется

	_, err = bot.SetWebhook(tgbotapi.NewWebhook(WebhookURL)) //начинаем слушать сервер ,ставим http сервер внутри который будет обрабатывать по этому url все сообщения от телеграма
	if err != nil {
		panic(err)
	}

	updates := bot.ListenForWebhook("/") // создаём канал  который будет всё обрабатывать и возращать (канал в который приходят сообщение)

	go http.ListenAndServe(":8080", nil) //создание веб сервера который будет обрабатывать все http запросы
	fmt.Println("starting server at :8080")
	// получаем все обнавления из канала updates
	for updete := range updates {
		if url, ok := rss[updete.Message.Text]; ok { //если нам пришло сообщение и в нём есть наш rss url то
			rss, err := getNews(url) // мы получим его новасти
			if err != nil {          // если всё плохо мы скажем извените у нас ошибка
				bot.Send(tgbotapi.NewMessage(
					updete.Message.Chat.ID,
					"sorry, error happend",
				))

			}
			for _, item := range rss.Items { //если всё хорошо то выведим эти сообщения
				bot.Send(tgbotapi.NewMessage( //вывод используется через bot.Send(объект нового соощения(мой индификатор каму отправить, url и заголовок этой новасти))
					updete.Message.Chat.ID,
					item.URL+"\n"+item.Title,
				))
			}
		} else {
			bot.Send(tgbotapi.NewMessage( //если таго резафибра нет то будет отпрвлена сообщение об ошибки
				updete.Message.Chat.ID,
				`there is only Harb feed availible`,
			))
		}
	}
}
func getNews(url string) (*RSS, error) {
	resp, err := http.Get(url) // преходит по Get запросу по этому url
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body) //вычитывам данные
	rss := new(RSS)
	err = xml.Unmarshal(body, rss) //распоковыем в xml
	if err != nil {
		return nil, err
	}
	return rss, nil
}
