package engines

import . "github.com/zengsai/utran/core"

type Engine interface {
    Name() string
    Vendor() string
    SupportQuery() bool
    Query(word string) Word
    SupportTranslate() bool
    Translate(stens string) SentPair
}

type engine struct {
    flag int
    name string
    vendor string
}

func (e *engine)Name() string {
    return e.name
}

func (e *engine)Vendor() string {
    return e.vendor
}

func (e *engine)SupportQuery() bool {
    return true
}

func (e *engine)SupportTranslate() bool {
    return false
}


// 新增的引擎到这里注册
func New(name string) Engine {
    if name == "iciba" {
        return new_iciba_engine()
    }
    return nil
}
