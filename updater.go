package main

import (
    "fmt"
    // lediscfg "github.com/siddontang/ledisdb/config"
    // "github.com/siddontang/ledisdb/ledis"

    "github.com/go-redis/redis"

    "strconv"
    "time"
)

var client = redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "", // no password set
    DB:       0,  // use default DB
})

func doEvery(d time.Duration, f func(time.Time)) {
    for x := range time.Tick(d) {
        f(x)
    }
}

func incrementCounter(t time.Time) {
    key := "counter"
    myval, _ := client.Get(key).Result()
    intv, _ := strconv.Atoi(string(myval))
    fmt.Println(intv)
    intv++
    intstr := strconv.Itoa(intv)
    client.Set(key, intstr, 0)
    // fmt.Println(myval)
    //
    // intv++
    //
    // db.Set(key, []byte(intstr))
}

func main() {

    doEvery(1000*time.Millisecond, incrementCounter)

}
