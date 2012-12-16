package main

import (
    "fmt"
    "github.com/zengsai/utran/engines"
)

func main() {
    engine := engines.New("iciba")
    if engine == nil {
        fmt.Println("no engine")
        return
    }

    if engine.SupportQuery() {
        fmt.Println(engine.Name(), "support translate")
        word := engine.Query("hello")
        fmt.Println("\t", word.Key, " -> ", word.Defs)
    }

    if engine.SupportTranslate() {
        fmt.Println(engine.Name(), "support tanslate")
    }

    return
}
