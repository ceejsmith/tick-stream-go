package stream

import (
    "net/http"
    "time"

    "golang.org/x/net/websocket"
)

func Strings(strings []string) {
    tickStream := func(ws *websocket.Conn) {
        for _, s := range strings {
            websocket.Message.Send(ws, s)
            time.Sleep(1000 * time.Millisecond)
        }
        ws.Close()
    }

    http.Handle("/tick", websocket.Handler(tickStream))
    http.ListenAndServe(":12345", nil)
}
