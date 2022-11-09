package slacker

import "github.com/slack-go/slack"

type InputBuilder struct {
	label    *slack.TextBlockObject
	hint     *slack.TextBlockObject
	actionId string
	blockId  string
}

func (b *InputBuilder) Label(label *slack.TextBlockObject) {
	b.label = label
}

func (b *InputBuilder) Hint(hint *slack.TextBlockObject) {
	b.hint = hint
}

type InputBuilderOption func(*InputBuilder)

func WithLabel(text string) InputBuilderOption {
	return func(builder *InputBuilder) {
		builder.Label(Pt(text))
	}
}

func WithHint(text string) InputBuilderOption {
	return func(builder *InputBuilder) {
		builder.Hint(Pt(text))
	}
}

func Input(blockId, actionId string, options ...InputBuilderOption) InputBuilderOption {
	return func(builder *InputBuilder) {
		builder.blockId = blockId
		builder.actionId = actionId
		for _, opt := range options {
			opt(builder)
		}
	}
}
