package main

import (
	"strconv"
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
		global.Set("currentDeep", js.ValueOf(info.CurrentNode.Depth))
		global.Set("isEnd", js.ValueOf(info.isEnd))
		global.Set("opened", js.ValueOf(info.Opened))
		if info.isSuccess {
			finishNode := info.CurrentNode
			global.Call("clearParent")
			for finishNode.Parent != nil {
				global.Call("addParent", js.ValueOf(finishNode.ToString()+":"+strconv.Itoa(finishNode.Value)+"="+strconv.Itoa(finishNode.Depth)+
					"+"+strconv.Itoa(finishNode.Value-finishNode.Depth)))
				finishNode = CopyNode(*finishNode.Parent)
			}
		}
	})
	global.Set("next", onSum)
}

func main() {
	c := make(chan struct{}, 0)
	println("hello wasm")
	registerCallbacks()
	<-c
}
