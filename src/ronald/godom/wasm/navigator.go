package wasm

// This file implements Navigator interface
// https://developer.mozilla.org/en-US/docs/Web/API/Navigator


// Properties

func (n *NavigatorSpec) Language() string {
   return n.Get("language").String()
}

func (n *NavigatorSpec) Languages() string {
   return n.Get("languages").String()
}
