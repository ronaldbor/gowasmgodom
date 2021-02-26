package wasm

// This file implements Location interface
// https://developer.mozilla.org/en-US/docs/Web/API/Location


// Properties

func (l *LocationSpec) Host() string {
   return l.Get("host").String()
}

func (l *LocationSpec) Hostname() string {
   return l.Get("hostname").String()
}

func (l *LocationSpec) Href() string {
   return l.Get("href").String()
}

func (l *LocationSpec) Origin() string {
   return l.Get("origin").String()
}

func (l *LocationSpec) Pathname() string {
   return l.Get("pathname").String()
}

func (l *LocationSpec) Port() string {
   return l.Get("port").String()
}

func (l *LocationSpec) Protocol() string {
   return l.Get("protocol").String()
}

func (l *LocationSpec) Search() string {
   return l.Get("search").String()
}
