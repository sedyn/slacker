package slacker

func OrBlock(condition bool, trueBlock, falseBlock SlackBlockBuilderProvider) SlackBlockBuilderProvider {
	if condition {
		return trueBlock
	} else {
		return falseBlock
	}
}

func OrFuncBlock(condition func() bool, trueBlock, falseBlock SlackBlockBuilderProvider) SlackBlockBuilderProvider {
	if condition() {
		return trueBlock
	} else {
		return falseBlock
	}
}
