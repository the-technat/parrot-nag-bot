package main

import (
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	RESTER_PORT            = 6789
	CONFIG_FILE_TYPE       = "yaml"
	CONFIG_FILE_PATH       = "/etc/parrot-nag-bot/"
	CONFIG_FILE_NAME       = "parrot-nag-bot_config"
	TELEBOT_POLLER_TIMEOUT = 10 * time.Second
)

func init() {

	//-------------------------- Configuration --------------------------
	// create viper.Viper object and set settings
	config := viper.New()
	config.SetConfigName(CONFIG_FILE_NAME) // name of config file (without extension)
	config.SetConfigType(CONFIG_FILE_TYPE) // REQUIRED if the config file does not have the extension in the name
	config.AddConfigPath(CONFIG_FILE_PATH) // path to look for the config file in
	config.AddConfigPath(".")              // optionally look for config in the working directory

	// parse config file
	if err := config.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.WithFields(log.Fields{
				"pkg":   "main",
				"error": err,
			}).Fatalf("Configfile was not found in any of the following paths: [%s, %s]", ".", CONFIG_FILE_PATH)
		} else {
			log.WithFields(log.Fields{
				"pkg":   "main",
				"error": err,
			}).Fatal("Error while parsing configfile, can not proceed")
		}
	}

	// If successful set Defaults

	// Adopt environment variables
	config.BindEnv("telegram.token", "TELEGRAM_TOKEN")

	// Live reload the config
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file has changed: %s", e.Name)
	})

	// check if log level is in correct format, otherwise set default
	loglevel, err := log.ParseLevel(config.GetString("logging.level"))
	if err != nil {
		log.WithFields(log.Fields{
			"pkg":   "main",
			"error": err,
		}).Errorf("Logging level is not provided or of wrong type: %s, using default (warning)", config.GetString("logging.level"))
		loglevel = log.WarnLevel
		config.SetDefault("logging.level", "warning")
	}

	//-------------------------- Logging --------------------------
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Set loglevel from config (or default)
	log.SetLevel(loglevel)

	//-------------------------- Bot --------------------------
	// instance tb.Bot
	bot, err := tb.NewBot(tb.Settings{
		Token:  config.GetString("telegram.token"),
		Poller: &tb.LongPoller{Timeout: TELEBOT_POLLER_TIMEOUT},
	})

	if err != nil {
		log.WithFields(log.Fields{
			"pkg":   "main",
			"error": err,
		}).Fatal("Failure when creating bot, maybe TELEGRAM_TOKEN is wrong or expired?")
	}

	//-------------------------- Rest --------------------------
	// instance rest.API

	// And go on to running the bot / rest api
	main(bot)

}

func main(bot *tb.Bot) {

	//-------------------------- rester --------------------------
	// Take off with the Rest API, provide at Port
	// Gracefull shutdown even in container

	//-------------------------- polling --------------------------
	// Start bot polling
	bot.Start()
}
