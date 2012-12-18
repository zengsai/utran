package engines

import . "github.com/zengsai/utran/core"
import   "encoding/xml"
// import "fmt"

const (
    url string = "http://dict-co.iciba.com/api/dictionary.php?w=%s"
    content string = `
<?xml version="1.0" encoding="UTF-8"?>
<dict num="219" id="219" name="219">
<key>word</key>
<ps>wə:d</ps>
<pron>http://res.iciba.com/resource/amp3/0/0/c4/7d/c47d187067c6cf953245f128b5fde62a.mp3</pron>
<ps>wɚd</ps>
<pron>http://res.iciba.com/resource/amp3/1/0/c4/7d/c47d187067c6cf953245f128b5fde62a.mp3</pron>
<pos>n.</pos>
<acceptation>单词，歌词，台词；（说的）话；诺言；命令；
</acceptation>
<pos>vt.</pos>
<acceptation>措辞，用词；用言语表达；
</acceptation>
<pos>vi.</pos>
<acceptation>讲话；
</acceptation>
</dict>
`
)

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
    err := xml.Unmarshal([]byte(content), &r)
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
    e.name = "iciba"
    e.vendor = "jinshan"
    e.host = "www.iciba.com"
    e.uri = "query.php?word=%s"
    return &e
}
