package slacker

import (
	"github.com/slack-go/slack"
)

type SlackBlockBuilder interface {
	Build() slack.Block
}

type SlackBlockBuilderProvider func() SlackBlockBuilder

// Blocks executes SlackBlockBuilderProvider and merges blocks into slack.Blocks
func Blocks(providers ...SlackBlockBuilderProvider) slack.Blocks {
	blockSet := make([]slack.Block, len(providers))
	for i, provider := range providers {
		blockSet[i] = provider().Build()
	}

	return slack.Blocks{BlockSet: blockSet}
}

func applyOptions[T SlackBlockBuilder](builder T, options []func(T)) T {
	for _, opt := range options {
		opt(builder)
	}
	return builder
}
