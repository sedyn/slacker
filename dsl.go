package slacker

import (
	"encoding/json"

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

type ModalOptions = func(*slack.ModalViewRequest)

func WithPrivateMetadata(privateMetadata string) ModalOptions {
	return func(request *slack.ModalViewRequest) {
		request.PrivateMetadata = privateMetadata
	}
}

func WithJsonPrivateMetadata(privateMetadata interface{}) ModalOptions {
	return func(request *slack.ModalViewRequest) {
		// TODO handle error?
		metadata, _ := json.Marshal(privateMetadata)
		request.PrivateMetadata = string(metadata)
	}
}

func ModalViewRequest(
	title, submit, close *slack.TextBlockObject,
	blocks slack.Blocks,
	options ...ModalOptions,
) slack.ModalViewRequest {
	request := slack.ModalViewRequest{
		Type:   slack.VTModal,
		Title:  title,
		Blocks: blocks,
		Close:  close,
		Submit: submit,
	}
	for _, opt := range options {
		opt(&request)
	}

	return request
}

func applyOptions[T SlackBlockBuilder](builder T, options []func(T)) T {
	for _, opt := range options {
		opt(builder)
	}
	return builder
}
