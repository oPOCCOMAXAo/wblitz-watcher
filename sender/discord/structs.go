package discord

import (
	"time"
)

type Message struct {
	Content string  `json:"content"`
	Embeds  []Embed `json:"embeds,omitempty"`
}

type Embed struct {
	Type        EmbedType `json:"type,omitempty"`
	Description string    `json:"description,omitempty"`
	Timestamp   time.Time `json:"timestamp,omitempty"`
	Color       Color     `json:"color,omitempty"`
	Author      *Author   `json:"author,omitempty"`
	Fields      []Field   `json:"fields,omitempty"`
	Thumbnail   *Image    `json:"thumbnail,omitempty"`
	Image       *Image    `json:"image,omitempty"` // could be only image.
}

type Author struct {
	Name    string `json:"name"`
	IconURL string `json:"icon_url"`
}

type Field struct {
	Inline bool   `json:"inline"`
	Name   string `json:"name"`
	Value  string `json:"value"`
}

type Image struct {
	URL string `json:"url"`
}
