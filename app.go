package main

import (
   "syscall/js"
   . "ronald/godom/wasm"
)




func eventHandler (event Event) {
   Log ("buttonHandler called")

   eventData := event.GetEventData ()

   Log ("Event:")
   Log ("   Type         = '%s'", eventData.Type)
   Log ("   TimeStamp    = '%d'", eventData.TimeStamp)
   Log ("   Target:")
   Log ("      Id          = '%s'", eventData.Target.Id)
   Log ("      TagName     = '%s'", eventData.Target.TagName)
   Log ("      Value       = '%s'", eventData.Target.Value)
   Log ("      TextContent = '%s'", eventData.Target.TextContent)
}



func functionInGo (this js.Value, args []js.Value) interface{} {
   Log ("functionInGo:")
   Log ("   Length args  = %d", len(args))
   for i,arg := range args {
      argType := arg.Type()
      Log ("%d:   Type = %s", i, argType.String() )
      switch (argType) {
         case js.TypeObject:
            id := arg.Get("id")
            Log ("     ObjId = %s", id.String() )
         case js.TypeString:
            Log ("     Value = %s", arg.String() )
         case js.TypeNumber:
            Log ("     Value = %d", arg.Int() )
         case js.TypeBoolean:
            Log ("     Value = %t", arg.Bool() )
         default:
            Log ("     Value = [other]")
      }
   }
   return nil
}


func main () {
   main5 ()
   select {}   // Prevent exiting
}


func main5 () {
   Log ("LOGGING: Starting main5")
   Alert ("Hi there, press OK to start")

   // Get Document.body
   body := Body()

   maindiv := CreateDIV ("maindiv", Options{"value": "bello"})
   maindiv.AddEventListener ("click", eventHandler)   // Use event delegation; look for the id in the handler

   input1 := CreateINPUT ("ip1", Options{"type": "text", "value": "click me at input1"})

   br := CreateBR ("br", Options{})

   input2 := CreateINPUT ("ip2", Options{"type": "text", "value": "click me at input2"})

   Log ("Create INPUT input3")
   ExposeFunction ("myhandler", functionInGo)
   input3 := CreateHTML (`<input id="ip3" type="text" value="click me at input3" onclick="myhandler(this, 'Hi', 4, false)">`)

   p := CreateP ("wodan", Options{"textContent": "click me at p"})
   // p2 := CreateP ("", Options{"textContent": "click me at p2"})

   Log ("Build page")
   body.AppendChild (maindiv)
   maindiv.AppendChild (input3)
   maindiv.AppendChild (br)
   maindiv.AppendChild (input2)
   maindiv.AppendChild (p)
   // maindiv.AppendChild (p2)
   maindiv.AppendChild (input1)

   Log ("Get head and body contents")
   allHTMLHead := Head().GetInnerHTML ()
   Log ("HTML Head = %s", allHTMLHead)
   allHTMLBody := Body().GetInnerHTML ()
   Log ("HTML Body = %s", allHTMLBody)
}

