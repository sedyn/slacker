package slacker

import "github.com/slack-go/slack"

type Viewer struct {
	view *slack.View
}

func NewViewer(view *slack.View) Viewer {
	return Viewer{view: view}
}

func (v *Viewer) Value(key string) {
	//value, ok := v.view.State.Values[key]
}
