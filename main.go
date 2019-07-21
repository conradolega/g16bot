package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/bwmarrin/discordgo"
)

func main() {
	discord, err := discordgo.New("Bot " + "DISCORD_API_KEY")
	if err != nil {
		fmt.Println("Error creating Discord session")
	}

	discord.AddHandler(ready)
	discord.AddHandler(messageCreate)

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

func messageCreate(s *discordgo.Session, event *discordgo.MessageCreate) {
	if event.Author.ID == s.State.User.ID {
		return
	}

	if event.Content == "!g16 hello" {
		s.ChannelMessageSend(event.ChannelID, "Hi "+event.Author.Mention())
	}

	if event.Content == "!g16 meme" {
		sendRandomMeme(s, event)
	}
}

func sendRandomMeme(s *discordgo.Session, event *discordgo.MessageCreate) {
	memeDir := "/where/your/memes/are/stored/"
	memes, err := ioutil.ReadDir(memeDir)
	if err != nil {
		fmt.Println("Failed to open memes folder")
	}

	rand.Seed(time.Now().Unix())
	meme := memes[rand.Intn(len(memes))]
	image, err := os.Open(memeDir + meme.Name())
	if err != nil {
		fmt.Println("Failed to open image")
	}

	s.ChannelFileSend(event.ChannelID, meme.Name(), image)
}
