package main

import (
	"math/rand"
	"syscall/js"
	"time"
)

const (
	width  = 600
	height = 600
)

func getRandomNum() float32 {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	n := float32(rand.Intn(10000))
	return n / 10000.0
}

func render() {
	var canvas js.Value = js.
		Global().
		Get("document").
		Call("getElementById", "canvas")
	var context js.Value = canvas.Call("getContext", "2d")
	// reset
	canvas.Set("height", height)
	canvas.Set("width", width)
	context.Call("clearRect", 0, 0, width, height)
	context.Call("beginPath")
	for i := 0; i < 50; i++ {
		context.Call("moveTo", getRandomNum()*width, getRandomNum()*height)
		context.Call("lineTo", getRandomNum()*width, getRandomNum()*height)
	}
	context.Call("stroke")
}

func addEventListener() {
	// see https://tip.golang.org/pkg/syscall/js/?GOOS=js&GOARCH=wasm#NewCallback
	done := make(chan struct{})
	var cb = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		go func() {
			println("Pressed ", this.String())
			render()
		}()
		return nil
	})
	js.Global().
		Get("document").
		Call("getElementById", "canvas").
		Call("addEventListener", "click", cb)
	<-done
	cb.Release()
}

func main() {
	alert := js.Global().Get("alert")
	alert.Invoke("Start webasm application")
	println("wasm app works")
	render()
	addEventListener() //wait callbak
}
