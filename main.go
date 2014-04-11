package main

import (
	"flag"
	"fmt"
	. "github.com/zengsai/utran/core"
	"github.com/zengsai/utran/engines"
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("need words")
		return
	}

	engine := engines.New("iciba")
	if engine == nil {
		fmt.Println("no engine")
		return
	}

	if flag.NArg() == 1 {
		query := flag.Arg(0)
		if engine.SupportQuery() {
			word := engine.Query(query)
			printWord(word)
			fmt.Println("===============================================")
			fmt.Println("Content Provided by", engine.Name()+"."+engine.Vendor())
		}
	} else {
		var orig string
		for _, w := range flag.Args() {
			orig += w + " "
		}

		if engine.SupportTranslate() {
			fmt.Println(orig)
			fmt.Println("===============================================")
			sp := engine.Translate(orig)
			fmt.Printf("%s", sp.Str)
			fmt.Println("===============================================")
			fmt.Println("Content Provided by", engine.Name()+"."+engine.Vendor())
		}
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
