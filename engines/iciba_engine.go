package engines

import . "github.com/zengsai/utran/core"
import   "encoding/xml"
import "net/http"
import "strings"
// import "fmt"

// const baseUrl string = "http://dict-co.iciba.com/api/dictionary.php?w="

type Result struct {
    XMLName xml.Name `xml:"dict"`
    Key string  `xml:"key"`
    Ps  []string `xml:"ps"`
    Pron []string `xml:"pron"`
    Pos []string `xml:"pos"`
    Acc []string `xml:"acceptation"`
}

type iciba_engine struct {
    engine
    host string
    uri string
}

func (e *iciba_engine)Query(word string) Word {
    var r Result

    resp, err := http.Get(e.host + e.uri + strings.ToLower(word))
    if err != nil {
        return Word{}
    }

    decoder := xml.NewDecoder(resp.Body)
    defer resp.Body.Close()

    err = decoder.Decode(&r)
    // fmt.Println(r)
    if err != nil {
        return Word{}
    }

    porn := make([]Pron, 0, 0)
    for i,v := range r.Ps {
        porn = append(porn, Pron{v, "", r.Pron[i]})
    }

    def := make([]Def, 0, 0)
    for i,v := range r.Pos {
        def = append(def, Def{v, r.Acc[i]})
    }

    // fmt.Println(porn)
    // fmt.Println(def)
    return Word{1, r.Key, porn, def, nil}
}

func (e *iciba_engine)Translate(stens string) SentPair {
    return SentPair{"workd", "hello", "porn url", "str"}
}

func new_iciba_engine() Engine {
    var e iciba_engine
    e.flag = 1
    e.name = "ICIBA"
    e.vendor = "JINSHAN"
    e.host = "http://dict-co.iciba.com"
    e.uri = "/api/dictionary.php?w="
    return &e
}
