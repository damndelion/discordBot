package applicator

import (
	"discord-go-bot/config"
	"discord-go-bot/internal/handler"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// Run sets up the application
func Run(cfg *config.Config) {
	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatal(err)
	}
	handler.NewHandler(session, cfg)

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	log.Println("Bot is running...")

	shutdownSignal := make(chan os.Signal, 1)
	signal.Notify(shutdownSignal, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-shutdownSignal

}
