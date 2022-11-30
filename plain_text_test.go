package slacker

import (
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestPlainTextInput(t *testing.T) {
	plainTextInput := build(PlainTextInput(
		Input("blockId", "actionId", WithLabel("label")),
		WithPlainTextPlaceholder("placeholder"),
	))

	assert.Equal(t,
		slack.NewInputBlock(
			"blockId",
			Pt("label"),
			nil,
			slack.NewPlainTextInputBlockElement(Pt("placeholder"), "actionId"),
		),
		plainTextInput,
	)
}
