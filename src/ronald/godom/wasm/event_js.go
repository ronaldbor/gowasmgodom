package wasm

import (
   "syscall/js"
)


/* See https://ian-says.com/articles/golang-in-the-browser-with-web-assembly/
Event listeners:
*/

type Event struct {
   js.Value
}

func (e Event) Type () string {
   return e.Get("type").String()
}

func (e Event) Target () Value {
   return Value{e.Get("target")}
}

func (e Event) TimeStamp () int {
   return e.Get("timeStamp").Int()
}


type eventTarget struct {
   Obj         Value
   Id          string
   TagName     string
   Value       string
   TextContent string
}


type EventData struct {
   Obj       Event
   Type      string
   TimeStamp int
   Target    eventTarget
}



func (event Event) GetEventData () EventData {
   targetObj := event.Target()

   return EventData {
      Obj       : event,
      Type      : event.Type(),
      TimeStamp : event.TimeStamp(),
      Target    : eventTarget {
         Obj         : targetObj,
         Id          : targetObj.GetId(),
         TagName     : targetObj.GetTagName(),
         Value       : targetObj.GetValue(),
         TextContent : targetObj.GetTextContent(),
      },
   }

}









type EventHandlerFunc func(Event)


func (fn EventHandlerFunc) wrapGoFunction () func(js.Value, []js.Value) interface {} {
   return func(_ js.Value, args []js.Value) interface {} {
      fn(Event{args[0]})
      return nil
   }
}

func (v Value) AddEventListener (eventStr string, listenerFunc EventHandlerFunc) {
   // Log ("Add Event Listener for event: %s", eventStr)
   wrapper := listenerFunc.wrapGoFunction ()
   v.Call ("addEventListener", eventStr, js.FuncOf(wrapper) )
   // Log ("... Added")
}
 
