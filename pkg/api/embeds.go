package api

import "github.com/bwmarrin/discordgo"

// Courtesy of https://gist.github.com/Necroforger/8b0b70b1a69fa7828b8ad6387ebb3835

//EmbedBuilder ...
type EmbedBuilder struct {
	*discordgo.MessageEmbed
}

// Constants for message embed character limits
const (
	EmbedLimitTitle       = 256
	EmbedLimitDescription = 2048
	EmbedLimitFieldValue  = 1024
	EmbedLimitFieldName   = 256
	EmbedLimitField       = 25
	EmbedLimitFooter      = 2048
	EmbedLimit            = 4000
)

//NewEmbed returns a new embed object
func NewEmbedBuilder() *EmbedBuilder {
	return &EmbedBuilder{&discordgo.MessageEmbed{}}
}

//SetTitle ...
func (e *EmbedBuilder) SetTitle(name string) *EmbedBuilder {
	e.Title = name
	return e
}

//SetDescription [desc]
func (e *EmbedBuilder) SetDescription(description string) *EmbedBuilder {
	if len(description) > 2048 {
		description = description[:2048]
	}
	e.Description = description
	return e
}

//AddField [name] [value]
func (e *EmbedBuilder) AddField(name, value string) *EmbedBuilder {
	if len(value) > 1024 {
		value = value[:1024]
	}

	if len(name) > 1024 {
		name = name[:1024]
	}

	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name:  name,
		Value: value,
	})

	return e

}

//AddInlineField [name] [value]
func (e *EmbedBuilder) AddInlineField(name, value string) *EmbedBuilder {
	if len(value) > 1024 {
		value = value[:1024]
	}

	if len(name) > 1024 {
		name = name[:1024]
	}

	e.Fields = append(e.Fields, &discordgo.MessageEmbedField{
		Name:   name,
		Value:  value,
		Inline: true,
	})

	return e

}

//SetFooter [Text] [iconURL]
func (e *EmbedBuilder) SetFooter(args ...string) *EmbedBuilder {
	iconURL := ""
	text := ""
	proxyURL := ""

	switch {
	case len(args) > 2:
		proxyURL = args[2]
		fallthrough
	case len(args) > 1:
		iconURL = args[1]
		fallthrough
	case len(args) > 0:
		text = args[0]
	case len(args) == 0:
		return e
	}

	e.Footer = &discordgo.MessageEmbedFooter{
		IconURL:      iconURL,
		Text:         text,
		ProxyIconURL: proxyURL,
	}

	return e
}

//SetImage ...
func (e *EmbedBuilder) SetImage(args ...string) *EmbedBuilder {
	var URL string
	var proxyURL string

	if len(args) == 0 {
		return e
	}
	if len(args) > 0 {
		URL = args[0]
	}
	if len(args) > 1 {
		proxyURL = args[1]
	}
	e.Image = &discordgo.MessageEmbedImage{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

//SetThumbnail ...
func (e *EmbedBuilder) SetThumbnail(args ...string) *EmbedBuilder {
	var URL string
	var proxyURL string

	if len(args) == 0 {
		return e
	}
	if len(args) > 0 {
		URL = args[0]
	}
	if len(args) > 1 {
		proxyURL = args[1]
	}
	e.Thumbnail = &discordgo.MessageEmbedThumbnail{
		URL:      URL,
		ProxyURL: proxyURL,
	}
	return e
}

//SetAuthor ...
func (e *EmbedBuilder) SetAuthor(args ...string) *EmbedBuilder {
	var (
		name     string
		iconURL  string
		URL      string
		proxyURL string
	)

	if len(args) == 0 {
		return e
	}
	if len(args) > 0 {
		name = args[0]
	}
	if len(args) > 1 {
		iconURL = args[1]
	}
	if len(args) > 2 {
		URL = args[2]
	}
	if len(args) > 3 {
		proxyURL = args[3]
	}

	e.Author = &discordgo.MessageEmbedAuthor{
		Name:         name,
		IconURL:      iconURL,
		URL:          URL,
		ProxyIconURL: proxyURL,
	}

	return e
}

//SetURL ...
func (e *EmbedBuilder) SetURL(URL string) *EmbedBuilder {
	e.URL = URL
	return e
}

//SetColor ...
func (e *EmbedBuilder) SetColor(clr int) *EmbedBuilder {
	e.Color = clr
	return e
}

// InlineAllFields sets all fields in the embed to be inline
func (e *EmbedBuilder) InlineAllFields() *EmbedBuilder {
	for _, v := range e.Fields {
		v.Inline = true
	}
	return e
}

// Truncate truncates any embed value over the character limit.
func (e *EmbedBuilder) Truncate() *EmbedBuilder {
	e.TruncateDescription()
	e.TruncateFields()
	e.TruncateFooter()
	e.TruncateTitle()
	return e
}

// TruncateFields truncates fields that are too long
func (e *EmbedBuilder) TruncateFields() *EmbedBuilder {
	if len(e.Fields) > 25 {
		e.Fields = e.Fields[:EmbedLimitField]
	}

	for _, v := range e.Fields {

		if len(v.Name) > EmbedLimitFieldName {
			v.Name = v.Name[:EmbedLimitFieldName]
		}

		if len(v.Value) > EmbedLimitFieldValue {
			v.Value = v.Value[:EmbedLimitFieldValue]
		}

	}
	return e
}

// TruncateDescription ...
func (e *EmbedBuilder) TruncateDescription() *EmbedBuilder {
	if len(e.Description) > EmbedLimitDescription {
		e.Description = e.Description[:EmbedLimitDescription]
	}
	return e
}

// TruncateTitle ...
func (e *EmbedBuilder) TruncateTitle() *EmbedBuilder {
	if len(e.Title) > EmbedLimitTitle {
		e.Title = e.Title[:EmbedLimitTitle]
	}
	return e
}

// TruncateFooter ...
func (e *EmbedBuilder) TruncateFooter() *EmbedBuilder {
	if e.Footer != nil && len(e.Footer.Text) > EmbedLimitFooter {
		e.Footer.Text = e.Footer.Text[:EmbedLimitFooter]
	}
	return e
}

// Build ...
func (e *EmbedBuilder) Build() *discordgo.MessageEmbed {
	return e.MessageEmbed
}