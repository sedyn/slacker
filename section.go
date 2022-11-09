package slacker

import "github.com/slack-go/slack"

type sectionBlockBuilder struct {
	textObj   *slack.TextBlockObject
	fields    []*slack.TextBlockObject
	accessory *slack.Accessory
	options   []slack.SectionBlockOption
}

func (b *sectionBlockBuilder) Build() slack.Block {
	return slack.NewSectionBlock(b.textObj, b.fields, b.accessory, b.options...)
}

type SectionBlockBuilderOption = func(*sectionBlockBuilder)

func WithTextObj(textObj *slack.TextBlockObject) SectionBlockBuilderOption {
	return func(builder *sectionBlockBuilder) {
		builder.textObj = textObj
	}
}

func WithFields(fields ...*slack.TextBlockObject) SectionBlockBuilderOption {
	return func(builder *sectionBlockBuilder) {
		builder.fields = fields
	}
}

func WithAccessory(accessory *slack.Accessory) SectionBlockBuilderOption {
	return func(builder *sectionBlockBuilder) {
		builder.accessory = accessory
	}
}

func WithSectionOptions(options ...slack.SectionBlockOption) SectionBlockBuilderOption {
	return func(builder *sectionBlockBuilder) {
		builder.options = options
	}
}

func Section(options ...SectionBlockBuilderOption) SlackBlockBuilderProvider {
	return func() SlackBlockBuilder {
		return applyOptions(&sectionBlockBuilder{}, options)
	}
}
