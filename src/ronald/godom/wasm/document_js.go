package wasm

// This file implements Document interface
// https://developer.mozilla.org/en-US/docs/Web/API/Document

// Properties

// Returns the currently focused element
func (v Value) ActiveElement() Value {
   return Value{v.Get("activeElement")}
}

func Head () Value {
   return Value{Document.Get("head")}
}

func Body () Value {
   return Value{Document.Get("body")}
}


// Methods

func CreateElement(tag string) Value {
   return Value{Document.Call("createElement", tag)}
}

func (v Value) CreateTextNode(textContent string) Value {
   return Value{Document.Call("createTextNode", textContent)}
}

func GetElementById (id string) Value {
   elem := Document.Call("getElementById", id)
   if elem.IsNull() {
      panic ("GetElementById, element not found:" + id)
   }
   return Value{elem}
}

func GetElementsByTag (tagName string) Value {
   elems := Document.Call ("getElementsByTagName", tagName)
   if elems.IsNull() {
      panic ("WASM.GetElementsByTag, elements not found: tagName = " + tagName)
   }
   return Value{elems}
}

func (v Value) Write(markup string) {
   Document.Call("write", markup)
}
