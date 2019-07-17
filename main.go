package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + "DISCORD_API_KEY")
	if err != nil {
		fmt.Println("Error creating Discord session")
	}

	discord.AddHandler(ready)

	err = discord.Open()
	if err != nil {
		fmt.Println("Error connecting to Discord")
	}

	fmt.Println("G16bot is ready to go!")
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-channel

	discord.Close()
}

func ready(s *discordgo.Session, event *discordgo.Ready) {
	s.UpdateStatus(0, "Hello G16!")
}
