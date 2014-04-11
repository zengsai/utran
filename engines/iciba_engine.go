package engines

import . "github.com/zengsai/utran/core"
import "encoding/xml"
import "net/http"
import "strings"
import "fmt"

// import "fmt"

// const baseUrl string = "http://dict-co.iciba.com/api/dictionary.php?w="
// http://dict-co.iciba.com/api/dictionary.php?w=hello&key=286BEA637E514C07A29382C9C559D964

/*
query:

<?xml version="1.0" encoding="UTF-8"?>
<dict num="219" id="219" name="219">
<key>hello</key>
<ps>hə'ləʊ</ps>
<pron>http://res-tts.iciba.com/5/d/4/5d41402abc4b2a76b9719d911017c592.mp3</pron>
<ps>hɛˈlo, hə-</ps>
<pron>http://res.iciba.com/resource/amp3/1/0/5d/41/5d41402abc4b2a76b9719d911017c592.mp3</pron>
<pos>int.</pos>
<acceptation>哈喽，喂；你好，您好；表示问候；打招呼；
</acceptation>
<pos>n.</pos>
<acceptation>“喂”的招呼声或问候声；
</acceptation>
<pos>vi.</pos>
<acceptation>喊“喂”；
</acceptation>
<sent><orig>
This document contains Hello application components of each document summary of the contents.
</orig>
<trans>
此文件包含组成Hello应用程序的每个文件的内容摘要。
</trans></sent>
<sent><orig>
In the following example, CL produces a combined source and machine-code listing called HELLO. COD.
</orig>
<trans>
在下面的例子中，CL将产生一个命名为HELLO.COD的源代码与机器代码组合的清单文件。
</trans></sent>
<sent><orig>
Hello! Hello! Hello! Hello! Hel-lo!
</orig>
<trans>
你好！你好！你好！你好！你好！
</trans></sent>
<sent><orig>
Hello! Hello! Hello! Hello! I'm glad to meet you.
</orig>
<trans>
你好！你好！你好！你好！见到你很高兴。
</trans></sent>
<sent><orig>
Hello Marie. Hello Berlioz. Hello Toulouse.
</orig>
<trans>
你好玛丽，你好柏里欧，你好图鲁兹。
</trans></sent>
</dict>

translate:

<?xml version="1.0" encoding="UTF-8"?>
<dict num="219" id="219" name="219">
<key>At night they go to bed early</key><fy>他们晚上早点睡觉
</fy>
<sent><orig>
At night they go to bed early, but they don't always go to sleep!
</orig>
<trans>
夜晚，他们很早就上床了，但是他们并不总是就睡着的！
</trans></sent>
</dict>
*/

type Result struct {
	XMLName xml.Name `xml:"dict"`
	Key     string   `xml:"key"`
	Ps      []string `xml:"ps"`
	Pron    []string `xml:"pron"`
	Pos     []string `xml:"pos"`
	Acc     []string `xml:"acceptation"`
}

type TransResult struct {
	XMLName xml.Name `xml:"dict"`
	Key     string   `xml:"key"`
	Fy      string   `xml:"fy"`
}

type iciba_engine struct {
	engine
	host string
	uri  string
	key  string
}

func (e *iciba_engine) SupportQuery() bool {
	return true
}

func (e *iciba_engine) Query(word string) Word {
	var r Result

	url := fmt.Sprintf("%s%s%s&key=%s", e.host, e.uri, strings.ToLower(word), e.key)

	resp, err := http.Get(url)
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
	for i, v := range r.Ps {
		porn = append(porn, Pron{v, "", r.Pron[i]})
	}

	def := make([]Def, 0, 0)
	for i, v := range r.Pos {
		def = append(def, Def{v, r.Acc[i]})
	}

	// fmt.Println(porn)
	// fmt.Println(def)
	return Word{1, r.Key, porn, def, nil}
}

func (e *iciba_engine) SupportTranslate() bool {
	return true
}

func (e *iciba_engine) Translate(stens string) SentPair {
	var t TransResult

	url := fmt.Sprintf("%s%s%s&key=%s", e.host, e.uri, strings.ToLower(stens), e.key)

	resp, err := http.Get(url)
	if err != nil {
		return SentPair{}
	}

	decoder := xml.NewDecoder(resp.Body)
	defer resp.Body.Close()

	err = decoder.Decode(&t)
	// fmt.Println(r)
	if err != nil {
		return SentPair{}
	}

	return SentPair{t.Key, "", "", t.Fy}
}

func new_iciba_engine() Engine {
	var e iciba_engine
	e.flag = 1
	e.name = "ICIBA"
	e.vendor = "JINSHAN"
	e.host = "http://dict-co.iciba.com"
	e.uri = "/api/dictionary.php?w="
	e.key = "286BEA637E514C07A29382C9C559D964"
	return &e
}
