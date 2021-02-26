// Webserver that serves gzipped content

package main

import "os"
import "os/exec"
import "flag"
import "log"
import "fmt"
import "net/http"
import "strings"


// Command flags
// Usage example:
//    go run server.go -listenPort=8084 -gzipEnabled=false

var (
   listenPort  = flag.Int  ("listenPort",  8083,  "Listening port")
   gzipEnabled = flag.Bool ("gzipEnabled", false, "Enable gzipped file servicing")
)


func fileSize (fileName string) int64 {
   f, err := os.Stat (fileName)
   if (err != nil) {
      return -1
   }
   return f.Size()
}


// 1. If gzip is not supported, supply the plain file
// 2. If the gzipped version doesnot exist, check if the file is large enough for gzipping.
// 3. If the gzipped file exists, supply the gzipped version
func myGzipHandler (w http.ResponseWriter, r *http.Request) {
   fileName     := r.URL.Path[1:]
   // fileNameGzip := fileName + ".gz"
   log.Printf ("Requested '%s'\n", fileName)

   if (*gzipEnabled == true) {
      contentType := r.Header.Get ("Accept-Encoding")
      if (strings.Contains (contentType, "gzip") == true) {
         // The browser accepts compressed files
         fileNameGzip := fileName + ".gz"
         if (fileSize (fileNameGzip) < 0) {
            // The compressed file is not available; check if gzipped is usefull
            size := fileSize (fileName)
            if (size > 2*1024*1024) {
               // Larger than 2 MB; compress it for future requests
               log.Printf ("   Compressing (size = %d)\n", size)
               cmd := exec.Command ("gzip", "-k", fileName)
               _, err2 := cmd.Output ()
               if err2 != nil {
                  log.Printf (err2.Error())
               }
            }
         }

         if (fileSize (fileNameGzip) > 0) {
            // The file is available
            fileName = fileNameGzip
            w.Header().Set ("Content-Type",     "gzip, deflate")
            w.Header().Set ("Content-Encoding", "gzip")
         }
      }
   }

   log.Printf ("Send '%s'\n", fileName)
   http.ServeFile (w, r, fileName)
}


func main() {
   flag.PrintDefaults ()

   flag.Parse ()
   http.HandleFunc ("/", myGzipHandler)

   listenAddress := fmt.Sprintf (":%d", *listenPort)
   log.Printf ("listening on %s ...", listenAddress)
   err := http.ListenAndServe (listenAddress, nil)
   if err != nil {
     log.Fatal ("ERROR starting http server : ", err)
     return
   }
}

