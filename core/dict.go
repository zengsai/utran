package core

// 代表一个词
type Word struct {
    Drec   int          // 翻译方向
    Key  string         // 被查询的词
    Prons    []Pron     // 音标
    Defs  []Def         // 定义
    Sents []SentPair    // 例句
}

type Pron struct {
    Ps string           // 音标
    Url string          // 音频本地地址
    WebUrl string       // 音频网络地址
}

type Def struct {
    Pos string      // 词性
    Str string      // 释义
}

type SentPair struct {
    Orig  string
    PronUrl string
    PronWebUrl string
    Str  string
}
