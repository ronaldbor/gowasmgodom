package wasm

import (
   "fmt"
   "syscall/js"
)

type Value struct {
   js.Value
}

type WindowSpec struct {
   js.Value
   Location  LocationSpec
   Navigator NavigatorSpec
   History   HistorySpec
}


type LocationSpec struct {
   js.Value
}

type NavigatorSpec struct {
   js.Value
}

type HistorySpec struct {
   js.Value
}


// Window, Document and Console
var Window    = WindowSpec{
                   js.Global(),
                   LocationSpec{js.Global().Get("location")},
                   NavigatorSpec{js.Global().Get("navigator")},
                   HistorySpec{js.Global().Get("history")},
                }
var Document  = Value{js.Global().Get("document")}
var Console   = Value{js.Global().Get("console")}

// var Location  = LocationSpec{js.Global().Get("location")}
// var Navigator = NavigatorSpec{js.Global().Get("navigator")}
// var History   = HistorySpec{js.Global().Get("history")}



func Alert (format string, v ...interface{}) {
   msg := fmt.Sprintf (format, v...)
   Window.Call ("alert", msg)
}

func Log (format string, v ...interface{}) {
   msg := fmt.Sprintf (format, v...)
   Console.Call ("log", msg)
}



func (w WindowSpec) PageXOffset() float64 {
   return w.Get("pageXOffset").Float()
}

func (w WindowSpec) PageYOffset() float64 {
   return w.Get("pageYOffset").Float()
}

func (w WindowSpec) ScrollX() float64 {
   return w.Get("scrollX").Float()
}

func (w WindowSpec) ScrollY() float64 {
   return w.Get("scrollY").Float()
}


    
