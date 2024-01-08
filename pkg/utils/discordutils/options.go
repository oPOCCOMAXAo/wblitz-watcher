package du

import (
	"github.com/bwmarrin/discordgo"
)

type OptionDecoder func(*discordgo.ApplicationCommandInteractionDataOption) error

type DecodersMap map[string]OptionDecoder

func ParseOptions(
	options []*discordgo.ApplicationCommandInteractionDataOption,
	ref map[string]OptionDecoder,
) error {
	for _, option := range options {
		decoder, ok := ref[option.Name]
		if !ok {
			continue
		}

		err := decoder(option)
		if err != nil {
			return err
		}
	}

	return nil
}

func DecoderInt[T ~int64](ref *T) OptionDecoder {
	return func(option *discordgo.ApplicationCommandInteractionDataOption) error {
		*ref = T(option.IntValue())

		return nil
	}
}

func DecoderString[T ~string](ref *T) OptionDecoder {
	return func(option *discordgo.ApplicationCommandInteractionDataOption) error {
		*ref = T(option.StringValue())

		return nil
	}
}

func DecoderChannelID[T ~string](ref *T) OptionDecoder {
	return func(option *discordgo.ApplicationCommandInteractionDataOption) error {
		ch := option.ChannelValue(nil)

		*ref = T(ch.ID)

		return nil
	}
}

func GetFromAnyMap[R any](m map[string]any, key string) R {
	var res R

	v, ok := m[key]
	if !ok {
		return res
	}

	res, _ = v.(R)

	return res
}
