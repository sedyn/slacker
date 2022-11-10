package slacker

import "github.com/slack-go/slack"

type Viewer struct {
	view *slack.View
}

func NewViewer(view *slack.View) Viewer {
	return Viewer{view: view}
}

type BlockAction struct {
	slack.BlockAction
}

func (b *BlockAction) SelectedOption() string {
	return b.BlockAction.SelectedOption.Value
}

func (v *Viewer) Value(blockId, actionId string) *BlockAction {
	values, ok := v.view.State.Values[blockId]
	if !ok {
		return nil
	}

	value, ok := values[actionId]
	if !ok {
		return nil
	}

	return &BlockAction{value}
}
