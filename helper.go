package slacker

import (
	"github.com/slack-go/slack"
)

// Pt is helper function for slack.PlainTextType
func Pt(text string) *slack.TextBlockObject {
	return slack.NewTextBlockObject(slack.PlainTextType, text, false, false)
}

// Mt is helper function for slack.MarkdownType
func Mt(text string) *slack.TextBlockObject {
	return slack.NewTextBlockObject(slack.MarkdownType, text, false, false)
}
