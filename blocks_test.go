package slacker

import (
	"testing"

	"github.com/slack-go/slack"
	"github.com/stretchr/testify/assert"
)

func TestBlocks(t *testing.T) {
	assert.Equal(t,
		slack.Blocks{
			BlockSet: []slack.Block{
				slack.NewSectionBlock(
					Pt("test"),
					nil,
					nil,
				),
				slack.NewSectionBlock(
					Pt("yea"),
					nil,
					nil,
				),
				slack.NewSectionBlock(
					Pt("nay"),
					nil,
					nil,
				),
			},
		},
		Blocks(
			Section(WithTextObj(Pt("test"))),
			OrBlock(
				true,
				Section(WithTextObj(Pt("yea"))),
				Section(WithTextObj(Pt("nay"))),
			),
			OrFuncBlock(
				func() bool {
					return false
				},
				Section(WithTextObj(Pt("yea"))),
				Section(WithTextObj(Pt("nay"))),
			),
		),
	)
}

func TestModalViewRequest(t *testing.T) {
	assert.Equal(t,
		slack.ModalViewRequest{
			Type:  slack.VTModal,
			Title: Pt("title"),
			Blocks: slack.Blocks{
				BlockSet: []slack.Block{
					slack.NewSectionBlock(Pt("Hello, world!"), nil, nil),
					slack.NewInputBlock(
						"blockId",
						Pt("Some options..."),
						nil,
						slack.NewOptionsSelectBlockElement(
							slack.OptTypeStatic,
							Pt("alphabet..."),
							"actionId",
							slack.NewOptionBlockObject("A", Pt("A"), nil),
							slack.NewOptionBlockObject("B", Pt("B"), nil),
							slack.NewOptionBlockObject("C", Pt("C"), nil),
						),
					),
				},
			},
			Close:           Pt("close"),
			Submit:          Pt("submit"),
			PrivateMetadata: "{\"item\":123}",
		},
		ModalViewRequest(
			Pt("title"), Pt("submit"), Pt("close"),
			Blocks(
				Section(WithTextObj(Pt("Hello, world!"))),
				OptionInput(
					Input("blockId", "actionId", WithLabel("Some options...")),
					WithOptionPlaceholder("alphabet..."),
					WithOptionStrings(
						"A", "B", "C",
					),
				),
			),
			WithJsonPrivateMetadata(map[string]int{
				"item": 123,
			}),
		),
	)
}
