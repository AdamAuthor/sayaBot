//    █████╗ ██████╗  █████╗ ███╗   ███╗ █████╗ ██╗   ██╗████████╗██╗  ██╗ ██████╗ ██████╗
//   ██╔══██╗██╔══██╗██╔══██╗████╗ ████║██╔══██╗██║   ██║╚══██╔══╝██║  ██║██╔═══██╗██╔══██╗
//   ███████║██║  ██║███████║██╔████╔██║███████║██║   ██║   ██║   ███████║██║   ██║██████╔╝
//   ██╔══██║██║  ██║██╔══██║██║╚██╔╝██║██╔══██║██║   ██║   ██║   ██╔══██║██║   ██║██╔══██╗
//   ██║  ██║██████╔╝██║  ██║██║ ╚═╝ ██║██║  ██║╚██████╔╝   ██║   ██║  ██║╚██████╔╝██║  ██║
//   ╚═╝  ╚═╝╚═════╝ ╚═╝  ╚═╝╚═╝     ╚═╝╚═╝  ╚═╝ ╚═════╝    ╚═╝   ╚═╝  ╚═╝ ╚═════╝ ╚═╝  ╚═╝

package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	_ "github.com/golang-migrate/migrate"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"os"
	"tgbot/pkg/repository"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error with reading configs: %s", err.Error())
	}

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repository.NewRepository(db)
	tgToken := os.Getenv("TOKEN")
	bot, err := tgbotapi.NewBotAPI(tgToken)
	if err != nil {
		log.Panic(err)
	}

	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60 * 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil { // If we got a message
			fmt.Println("Hello World!")
		}
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
