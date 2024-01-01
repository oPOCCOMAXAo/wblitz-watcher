package discord

import (
	"log"

	"github.com/bwmarrin/discordgo"
)

func (s *Service) onReady(
	_ *discordgo.Session,
	event *discordgo.Ready,
) {
	log.Printf("%+v\n", event)

	bulkCmds := []*discordgo.ApplicationCommand{}
	for _, desc := range s.commands {
		bulkCmds = append(bulkCmds, desc.Command)
	}

	for _, guild := range event.Guilds {
		cmds, err := s.session.ApplicationCommands(s.config.ApplicationID, guild.ID)
		if err != nil {
			log.Printf("%+v\n", err)
		}

		found := map[string]bool{}

		for _, cmd := range cmds {
			_, ok := s.commands[cmd.Name]
			if !ok {
				err = s.session.ApplicationCommandDelete(
					s.config.ApplicationID,
					guild.ID,
					cmd.ID,
				)
				if err != nil {
					log.Printf("%+v\n", err)
				}
			}

			found[cmd.Name] = true
		}

		_, err = s.session.ApplicationCommandBulkOverwrite(s.config.ApplicationID, guild.ID, bulkCmds)
		if err != nil {
			log.Printf("%+v\n", err)
		}
	}
}

func (s *Service) onInteractionCreate(
	session *discordgo.Session,
	event *discordgo.InteractionCreate,
) {
	data := event.ApplicationCommandData()

	log.Printf("%s\n%+v\n\n", data.Name, event)

	cmd, ok := s.commands[data.Name]
	if !ok {
		return
	}

	if cmd.Handler == nil {
		return
	}

	resp, err := cmd.Handler(event)
	if err != nil {
		log.Printf("%+v\n", err)
	}

	if resp == nil {
		return
	}

	err = session.InteractionRespond(
		event.Interaction,
		resp,
	)
	if err != nil {
		log.Printf("%+v\n", err)
	}
}
