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
	request := ModalViewRequest(
		Pt("title"), nil, Pt("close"),
		Blocks(
			Section(WithTextObj(Pt("text"))),
		),
		WithJsonPrivateMetadata(metadata),
	)
}
```

## Helper

### Pt

Same as `slack.NewTextBlockObject(slack.PlainTextType, text, false, false)`

### Mt

Same as `slack.NewTextBlockObject(slack.MarkdownType, text, false, false)`

## DSL

### Blocks

Run slack block providers and merge blocks into `slack.Blocks`

```go
Blocks(
	Section(WithTextObj(Pt("text"))),
	Section(WithTextObj(Pt("text")))
)
```

### OrBlock

Select block based on condition

```go
// Pt("true") will return
OrBlock(true, Pt("true"), Pt("false"))
```

### OrFuncBlock

Same as above but condition is function.

```go
OrFuncBlock(func() {
	return random.Intn(100) & 1 == 0
}, Pt("Even"), Pt("Odd"))
```

### ModalViewRequest

Build slack.ModalViewRequest
