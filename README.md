# gowasmgodom

Example of combination of golang, web assembly and a godom interface.

Golang is used as the programming language and compiled to the webassembly format, which runs in the webbrowser.
It uses the godomwasm library (see src/ronald/godom/wasm) for creating DOM elements from Golang and attaches some eventhandlers to it.
The godomwasm library uses the dom by invoking its functions via the syscall/js.


## Prepare environment

I am working on CentOS 8, so normally when I issue the command "go" for the first time, it will fail to find it and suggests to install it. Confirm.

If you are working with other Linux OSses, you can just download Golang at https://golang.org/dl/, and select the version for Linux.
For example: https://golang.org/dl/go1.14.6.linux-amd64.tar.gz

Unpack the tar.gz-file, and add the "\<unpacked-map\>/bin" to your PATH-variable to make "go" available.


## Compiling

1. The file ```server.go``` is the webserver and serves the code that runs in the browser client.
2. The file ```app.go``` is the application and creates the DOM elements that builds the HTML-page. 

Compile your sources:
1. Compiling ```server.go```: ```sh ./comp.sh linux server.go```, resulting in ```server.bin```
2. Compiling ```app.go```   : ```sh ./comp.sh wasm  app.go```   , resulting in ```app.wasm```, plus ```wasm_exec.js```


## Running the website

You need the following files:
1. ```server.bin```
2. ```index.html```
3. ```app.wasm```
4. ```gluecode.js```
5. ```wasm_exec.js```

The last one is copied by the comp.sh from ```<location-go-compiler>/misc/wasm/wasm_exec.js```

1. Start the webserver: ```./server.bin``` (or: ```go run server.go```)
2. Open a browser and go to ```http://localhost:8083```

If you want logging information, open a console in the browser.\
For Google Chrome: Menu -> More Tools -> Developer Tools\
For Firefox: Menu -> Web Developer -> Web Console


## License

Check if it does not infringe the rights of the original.
See https://github.com/siongui/godom.

## Credits

A huge credit to https://github.com/siongui/godom/wasm, which I used as a starting point.
This original code is clear and easy understandable; I assumed it was ment to stay close at the JavaScript interface for DOM elements.

I added the event-stuff and some functions to easily add elements.
1. ```events_js.go```
2. ```extra_js.go```
3. ```location.go```
4. ```navigation.go```
5. ```history.go```

I also tried to abstract a little bit from the strict JS dom-interface.



