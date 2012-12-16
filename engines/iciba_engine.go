package engines

import . "github.com/zengsai/utran/core"

const (
    url := "http://dict-co.iciba.com/api/dictionary.php?w=%s"
)

type iciba_engine struct {
    engine
    host string
    uri string
}

func (e *iciba_engine)Query(word string) Word {
    return Word{1, word, nil, []Def{{"vt", "扫招呼"}, {"vi", "不打招呼"}}, nil}
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
