package slacker

import (
	"testing"
)

func TestBlocks(t *testing.T) {
	Blocks(
		Section(WithFields()),
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
	)

	ModalViewRequest(
		Pt("title"), Pt("submit"), Pt("cancel"),
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
	)
}
