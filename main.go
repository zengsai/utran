package main

import (
    "fmt"
    "flag"
    "github.com/zengsai/utran/engines"
    . "github.com/zengsai/utran/core"
)

func main() {
    flag.Parse()

    if flag.NArg() < 1 {
        fmt.Println("need words")
        return
    }

    query := flag.Arg(0)

    engine := engines.New("iciba")
    if engine == nil {
        fmt.Println("no engine")
        return
    }

    if engine.SupportQuery() {
        word := engine.Query(query)
        printWord(word)
        fmt.Println("===============================================")
        fmt.Println("Content Provided by", engine.Name() + "." + engine.Vendor())
    }

    if engine.SupportTranslate() {
        fmt.Println(engine.Name(), "support tanslate")
    }

    return
}

func printWord(w Word) {
        fmt.Print(w.Key, "\t")
        for _, v := range w.Prons {
            fmt.Print("[", v.Ps, "]")
            break
        }
        fmt.Println("\n===============================================")
        for _, v := range w.Defs {
            fmt.Print(v.Pos, "\t", v.Str)
        }
}
