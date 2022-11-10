package slacker

import (
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestPt(t *testing.T) {
	text := Pt("test")
	assert.Equal(t,
		slack.NewTextBlockObject(slack.PlainTextType, "test", false, false),
		text,
	)
}

func TestMt(t *testing.T) {
	text := Mt("test")
	assert.Equal(t,
		slack.NewTextBlockObject(slack.MarkdownType, "test", false, false),
		text,
	)
}
