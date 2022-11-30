package slacker

import (
	"encoding/json"

	"github.com/slack-go/slack"
)

type ModalOptions = func(*slack.ModalViewRequest)

func WithPrivateMetadata(privateMetadata string) ModalOptions {
	return func(request *slack.ModalViewRequest) {
		request.PrivateMetadata = privateMetadata
	}
}

func WithJsonPrivateMetadata(privateMetadata interface{}) ModalOptions {
	return func(request *slack.ModalViewRequest) {
		// TODO handle error
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
