// Webserver that serves gzipped content

package main

import "flag"
import "log"
import "fmt"
import "net/http"


// Command flags
// Usage example:
//    go run server.go -listerPort=8084

var listenPort = flag.Int  ("listenPort",  8083, "Listening port")


func myHandler (w http.ResponseWriter, r *http.Request) {
   fileName := r.URL.Path[1:]
   log.Printf ("Requested '%s'\n", fileName)
   http.ServeFile (w, r, fileName)
}


func main() {
   flag.Parse ()
   http.HandleFunc ("/", myHandler)

   listenAddress := fmt.Sprintf (":%d", *listenPort)
   log.Printf ("listening on %s ...", listenAddress)
   err := http.ListenAndServe (listenAddress, nil)
   if err != nil {
     log.Fatal ("ERROR starting http server : ", err)
     return
   }
}

