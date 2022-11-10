package slacker

import "github.com/slack-go/slack"

type optionInputBuilder struct {
	inputBuilder
	placeholder *slack.TextBlockObject
	options     []*slack.OptionBlockObject
}

func (b *optionInputBuilder) Build() slack.Block {
	element := slack.NewOptionsSelectBlockElement(slack.OptTypeStatic, b.placeholder, b.actionId, b.options...)
	return slack.NewInputBlock(b.blockId, b.label, b.hint, element)
}

type OptionInputBuilderOption = func(*optionInputBuilder)

func WithOptionPlaceholder(placeholder string) OptionInputBuilderOption {
	return func(builder *optionInputBuilder) {
		builder.placeholder = Pt(placeholder)
	}
}

func WithOptionInputOptions(options []*slack.OptionBlockObject) OptionInputBuilderOption {
	return func(builder *optionInputBuilder) {
		builder.options = options
	}
}

func WithOptionStrings(options ...string) OptionInputBuilderOption {
	return func(builder *optionInputBuilder) {
		optionBlockObjects := make([]*slack.OptionBlockObject, len(options))
		for i, o := range options {
			optionBlockObjects[i] = slack.NewOptionBlockObject(o, Pt(o), nil)
		}
		builder.options = optionBlockObjects
	}
}

func OptionInput(inputOption InputBuilderOption, options ...OptionInputBuilderOption) SlackBlockBuilderProvider {
	return func() SlackBlockBuilder {
		builder := &optionInputBuilder{}
		inputOption(&builder.inputBuilder)
		return applyOptions(builder, options)
	}
}
