package wasm

import (
   "fmt"
   "syscall/js"
)


func (v Value) RemoveAllChildren () {
   for v.HasChildNodes () {
      // Log ("RemoveAllChildren: Removing a child from div")
      v.RemoveLastChild ()
   }
}


//***************************************************************************
//*
//* Exposed function to the Javascript domain
//*
//***************************************************************************

type ExportedFunc func (this js.Value, args []js.Value) interface{}

func ExposeFunction (funcNameInJS string, funcInWASM ExportedFunc) {
   Window.Set (funcNameInJS, js.FuncOf(funcInWASM) )
}


//***************************************************************************
//*
//* Exposed function to the Javascript domain
//*
//***************************************************************************

/*
func GetElementByIdInDiv (divName string, tagName string) Value {
   div := Document
   if (divName != "") {
      div = GetElementById (divName)
   }

   elem := div.Call ("getElementsById", tagName)
   if elem.IsNull() {
      panic ("WASM.GetElementByIdInDiv, element not found: divName = " + divName + ", tagName = " + tagName)
   }
   return Value{elem}
}
*/


func GetElementsByTagInDiv (divName string, tagName string) Value {
   div := Document
   if (divName != "") {
      div = GetElementById (divName)
   }

   elems := div.Call ("getElementsByTagName", tagName)
   if elems.IsNull() {
      panic ("WASM.GetElementsByTag, elements not found: divName = " + divName + ", tagName = " + tagName)
   }
   return Value{elems}
}


func OpenWebSiteInNewTab (format string, v ...interface{}) {
   url := fmt.Sprintf (format, v...)
   win := Window.Call ("open", url, "_blank");
   win.Call ("focus")   // Go to the tab
}


func (v Value) IsFocused() bool {
   return v.IsEqualNode (Document.ActiveElement())
}

func (v Value) IsEqualNode (n Value) bool {
   return v.Call("isEqualNode", n).Bool()
}



type Options map[string]string

func (el *Value) parseOptions (data Options) Value {
   for key, value := range data {
      switch key {
         case "id"         : el.SetId (value)
         case "type"       : el.SetType (value)
         case "value"      : el.SetValue (value)
         case "textContent": el.SetTextContent (value)
         default: Log ("parseOptions, Unsupported option: '%s'", key)
      }
   }
   return *el
}

func createElementWithOptions (tagName string, idStr string, data Options) Value {
   Log ("Create %s with id '%s'", tagName, idStr)
   if (idStr == "") {
      panic ("Empty id for element is not allowed")
   }
   el := CreateElement (tagName)
   el.SetId (idStr)
   el.parseOptions (data)
   return el
}


func CreateDIV (idStr string, options Options) Value {
   return createElementWithOptions ("DIV", idStr, options)
}

func CreateINPUT (idStr string, options Options) Value {
   return createElementWithOptions ("INPUT", idStr, options)
}

func CreateBR (idStr string, options Options) Value {
   return createElementWithOptions ("BR", idStr, options)
}

func CreateP (idStr string, options Options) Value {
   return createElementWithOptions ("P", idStr, options)
}

func CreateHTML (htmlStr string) Value {
   el := CreateElement ("div")
   el.SetInnerHTML (htmlStr)
   return el
}



