package discord

import (
	"github.com/bwmarrin/discordgo"

	du "github.com/opoccomaxao/wblitz-watcher/pkg/utils/discordutils"
)

type CommandFullName struct {
	Name    string
	SubName string
}

type CommandHandler func(*discordgo.InteractionCreate, *CommandData) (*Response, error)

type CommandData struct {
	Name    []string
	Options map[string]any
}

func (d *CommandData) ID() CommandFullName {
	res := CommandFullName{}

	if len(d.Name) > 0 {
		res.Name = d.Name[0]
	}

	if len(d.Name) > 1 {
		res.SubName = d.Name[1]
	}

	return res
}

func (d *CommandData) String(name string) string {
	return du.GetFromAnyMap[string](d.Options, name)
}

func (d *CommandData) Int(name string) int64 {
	return du.GetFromAnyMap[int64](d.Options, name)
}
