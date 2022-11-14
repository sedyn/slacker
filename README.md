# Slacker

Simple Slack block builder

## Usage

```go
package main

import . "github.com/snpkx/slacker"

func main() {
	blocks := Blocks(
		Section(WithTextObj(Pt("Hello, world!"))),
		Section(WithFields(Pt("left"), Pt("right"))),
		OptionInput(
			Input("blockId", "actionId", WithLabel("Choose one")),
			WithOptionStrings("A", "B", "C"),
			WithOptionPlaceholder("alphabets..."),
		),
	)
}
```

```go
package main

import . "github.com/snpkx/slacker"

type Metadata struct {
	Stage int
}

func main() {
	metadata := Metadata{
		Stage: 1,
	}
	blocks := ModalViewRequest(
		Pt("title"), nil, Pt("close"),
		Blocks(
			Section(WithTextObj(Pt("text"))),
		),
		WithJsonPrivateMetadata(metadata),
	)
}
```
