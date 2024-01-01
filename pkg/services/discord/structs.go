package discord

import "github.com/bwmarrin/discordgo"

type InteractionDescription struct {
	Handler CommandHandler
	Command *discordgo.ApplicationCommand
}

type CommandHandler func(*discordgo.InteractionCreate) (*discordgo.InteractionResponse, error)
