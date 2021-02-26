package wasm

// This file implements Element interface
// https://developer.mozilla.org/en-US/docs/Web/API/Element

// Properties

func (v Value) ClassList() DOMTokenList {
   return DOMTokenList{v.Get("classList")}
}

// Methods

func (v Value) GetAttribute(attributeName string) string {
   return v.Call("getAttribute", attributeName).String()
}

func (v Value) GetBoundingClientRect() DOMRect {
   return DOMRect{v.Call("getBoundingClientRect")}
}

func (v Value) QuerySelector(selectors string) Value {
   return Value{v.Call("querySelector", selectors)}
}

func (v Value) QuerySelectorAll(selectors string) []Value {
   nodeList := v.Call("querySelectorAll", selectors)
   length := nodeList.Get("length").Int()
   var nodes []Value
   for i := 0; i < length; i++ {
      nodes = append(nodes, Value{nodeList.Call("item", i)})
   }
   return nodes
}



func (v Value) SetInnerHTML (html string) {
   v.Set("innerHTML", html)
}

func (v Value) GetInnerHTML () string {
   return v.Get("innerHTML").String()
}


func (v Value) SetOuterHTML (html string) {
   v.Set("outerHTML", html)
}

func (v Value) GetOuterHTML () string {
   return v.Get("outerHTML").String()
}


func (v Value) GetTagName() string {
   return v.Get("tagName").String()
}

func (v Value) SetTextContent (txt string) {
   v.Set("textContent", txt)
}

func (v Value) GetTextContent () string {
   return v.Get("textContent").String()
}


func (v Value) SetId (txt string) {
   v.Set("id", txt)
}

func (v Value) GetId () string {
   return v.Get("id").String()
}


func (v Value) SetType (txt string) {
   v.Set("type", txt)
}

func (v Value) GetType () string {
   return v.Get("type").String()
}


func (v Value) SetValue (txt string) {
   v.Set("value", txt)
}

func (v Value) GetValue () string {
   return v.Get("value").String()
}









func (v Value) AppendChild (elem Value) {
   v.Call ("appendChild", elem)
}

func (v Value) AppendAfter (elem Value) {
   v.Call ("appendAfter", elem)
}

func (v Value) HasChildNodes() bool {
   return v.Call ("hasChildNodes").Bool()
}

func (v Value) RemoveLastChild () {
   lastChild := v.Get ("lastChild")
   v.Call ("removeChild", lastChild)
}


func (v Value) Focus() {
   v.Call ("focus")
}


