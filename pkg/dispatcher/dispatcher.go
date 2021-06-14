package dispatcher

// Dispatcher - handles the telegram commands and holds telebot.Bot
type Dispatcher struct{}

func New() *Dispatcher {
	return &Dispatcher{}
}
