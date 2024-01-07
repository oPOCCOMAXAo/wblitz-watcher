package discord

import "github.com/bwmarrin/discordgo"

type InteractionDescription struct {
	Handler CommandHandler
	Command *discordgo.ApplicationCommand
}

type CommandHandler func(*discordgo.InteractionCreate) (*Response, error)

type Response struct {
	Content         string
	Components      []discordgo.MessageComponent
	Embeds          []*discordgo.MessageEmbed
	Files           []*discordgo.File
	AllowedMentions discordgo.MessageAllowedMentions
}

func (r *Response) WebHookEdit() *discordgo.WebhookEdit {
	return &discordgo.WebhookEdit{
		Content:         &r.Content,
		Components:      &r.Components,
		Embeds:          &r.Embeds,
		Files:           r.Files,
		AllowedMentions: &r.AllowedMentions,
	}
}
