package slacker

import "github.com/slack-go/slack"

type plainTextInputBuilder struct {
	inputBuilder
	placeholder *slack.TextBlockObject
}

func (b *plainTextInputBuilder) Build() slack.Block {
	element := slack.NewPlainTextInputBlockElement(b.placeholder, b.actionId)
	return slack.NewInputBlock(b.blockId, b.label, b.hint, element)
}

type PlainTextInputBuilderOption = func(*plainTextInputBuilder)

func WithPlainTextPlaceholder(placeholder string) PlainTextInputBuilderOption {
	return func(builder *plainTextInputBuilder) {
		builder.placeholder = Pt(placeholder)
	}
}

func PlainTextInput(inputOption InputBuilderOption, options ...PlainTextInputBuilderOption) SlackBlockBuilderProvider {
	return func() SlackBlockBuilder {
		builder := &plainTextInputBuilder{}
		inputOption(&builder.inputBuilder)
		return applyOptions(builder, options)
	}
}
