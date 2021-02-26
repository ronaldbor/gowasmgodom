package wasm

// This file implements History interface
// https://developer.mozilla.org/en-US/docs/Web/API/History


// Properties

func (h *HistorySpec) Length() int {
   return h.Get("length").Int()
}

// Methods

func (h *HistorySpec) Back() {
   h.Call("back")
}

func (h *HistorySpec) Forward() {
   h.Call("forward")
}

func (h *HistorySpec) Go(p int) {
   h.Call("go", p)
}

func (h *HistorySpec) PushState(stateObj interface{}, title, url string) {
   h.Call("pushState", stateObj, title, url)
}

func (h *HistorySpec) ReplaceState(stateObj interface{}, title, url string) {
   h.Call("replaceState", stateObj, title, url)
}
