// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
    "flag"

    "fmt"
    "log"
    "net/http"

    "github.com/go-redis/redis"

    "strconv"

    "github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8081", "http service address")

var upgrader = websocket.Upgrader{} // use default options

var client = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})

func echo(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true }
    c, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Print("upgrade:", err)
        return
    }
    defer c.Close()
    for {
        mt, _, err := c.ReadMessage()
        if err != nil {
            log.Println("read:", err)
            break
        }
        // log.Printf("recv: %s", message)

        key := "counter"
        myval, _ := client.Get(key).Result()
        intv, _ := strconv.Atoi(string(myval))
        fmt.Println(intv)
        intstr := strconv.Itoa(intv)

        err = c.WriteMessage(mt, []byte(intstr))

        log.Printf("send: %s", intstr)

        if err != nil {
            log.Println("write:", err)
            break
        }
    }
}

func main() {

    // // Read the value back out of the store.

    flag.Parse()
    log.SetFlags(0)
    http.HandleFunc("/echo", echo)
    // http.HandleFunc("/", home)
    log.Fatal(http.ListenAndServe(*addr, nil))
}
