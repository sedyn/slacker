# Slacker

Simple Slack block builder

## Usage

```go
package main

import . "github.com/kylepark/slacker"

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
