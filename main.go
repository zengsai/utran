package main

import (
    "fmt"
    "github.com/zengsai/utran/engines"
    . "github.com/zengsai/utran/core"
)

func main() {
    engine := engines.New("iciba")
    if engine == nil {
        fmt.Println("no engine")
        return
    }

    if engine.SupportQuery() {
        //fmt.Println(engine.Name(), "support translate")
        word := engine.Query("hello")
        printWord(word)
    }

    if engine.SupportTranslate() {
        fmt.Println(engine.Name(), "support tanslate")
    }

    return
}

func printWord(w Word) {
        fmt.Println(w.Key, "\n")
        for _, v := range w.Prons {
            fmt.Print("[", v.Ps, "]\t\t")
        }
        fmt.Println("\n")
        for _, v := range w.Defs {
            fmt.Print(v.Pos, "\t", v.Str)
        }
}
