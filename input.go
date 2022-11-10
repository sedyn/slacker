package slacker

import "github.com/slack-go/slack"

type inputBuilder struct {
	label    *slack.TextBlockObject
	hint     *slack.TextBlockObject
	actionId string
	blockId  string
}

func (b *inputBuilder) Label(label *slack.TextBlockObject) {
	b.label = label
}

func (b *inputBuilder) Hint(hint *slack.TextBlockObject) {
	b.hint = hint
}

type InputBuilderOption func(*inputBuilder)

func WithLabel(text string) InputBuilderOption {
	return func(builder *inputBuilder) {
		builder.Label(Pt(text))
	}
}

func WithHint(text string) InputBuilderOption {
	return func(builder *inputBuilder) {
		builder.Hint(Pt(text))
	}
}

func Input(blockId, actionId string, options ...InputBuilderOption) InputBuilderOption {
	return func(builder *inputBuilder) {
		builder.blockId = blockId
		builder.actionId = actionId
		for _, opt := range options {
			opt(builder)
		}
	}
}
