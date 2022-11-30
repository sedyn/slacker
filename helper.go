package slacker

import (
	"github.com/slack-go/slack"
)

type TextBlockObjectOption int

const (
	WithEmoji TextBlockObjectOption = iota
	WithVerbatim
)

func useTextBlockObjectOptions(options []TextBlockObjectOption) (useEmoji, useVerbatim bool) {
	for _, option := range options {
		switch option {
		case WithEmoji:
			useEmoji = true
		case WithVerbatim:
			useVerbatim = true
		}
	}

	return
}

// Pt is helper function for slack.PlainTextType
func Pt(text string, options ...TextBlockObjectOption) *slack.TextBlockObject {
	useEmoji, useVerbatim := useTextBlockObjectOptions(options)
	return slack.NewTextBlockObject(slack.PlainTextType, text, useEmoji, useVerbatim)
}

// Mt is helper function for slack.MarkdownType
func Mt(text string, options ...TextBlockObjectOption) *slack.TextBlockObject {
	useEmoji, useVerbatim := useTextBlockObjectOptions(options)
	return slack.NewTextBlockObject(slack.MarkdownType, text, useEmoji, useVerbatim)
}

// build returns single slack.Block
func build(provider SlackBlockBuilderProvider) slack.Block {
	return provider().Build()
}
