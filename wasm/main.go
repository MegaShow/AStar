package main

import (
	"syscall/js"
)

var nextStep = AStarNextStep()

func registerCallbacks() {
	global := js.Global()
	onSum := js.NewCallback(func(args []js.Value) {
		info := nextStep()
		global.Set("currentNode", js.ValueOf(info.CurrentNode.ToString()))
		global.Set("currentValue", js.ValueOf(info.CurrentNode.Value))
		global.Set("searchedNode", js.ValueOf(info.SearchedNode))
		global.Set("isSuccess", js.ValueOf(info.isSuccess))
		global.Set("isEnd", js.ValueOf(info.isEnd))
		global.Set("opened", js.ValueOf(info.Opened))
	})
	global.Set("next", onSum)
}

func main() {
	c := make(chan struct{}, 0)
	println("hello wasm")
	registerCallbacks()
	<-c
}
