/**
 * UbuntuTranslator
 * 作者：红猎人
 *
 * 该软件通过调用Google翻译API，实现简单的翻译功能。
 * Google语言API格式：
 *	    http://ajax.googleapis.com/ajax/services/language/translate?v=1.0&q=words&langpair=from|to
 * 返回数据格式（json）：
 *      return string: {"responseData": {"translatedText":"你好世界"}, "responseDetails": null, "responseStatus": 200}
 *
 *
 * 当前版本： 1.0.3
 */
package main

import (
	"http"
	"json"
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
)

const (
	e2c      = "en|zh"
	c2e      = "zh|en"
	baseUrl  = "http://ajax.googleapis.com/ajax/services/language/translate?v=1.0&q="
	langpair = "&langpair="
	version  = "v 1.0.3"
	year	 = "2010"
	author   = "红猎人"
	email    = "zengsai@gmail.com"
)

// 储存返回的数据存储，json解析中使用
type Response struct {
	ResponseData    Data
	ResponseDetails string
	ResponseStatus  int
}

type Data struct {
	TranslatedText string
}

// 三个布尔命令行选项的标识，默认为false。
var ec = flag.Bool("e2c", false, "translate from english to chinese")
var ce = flag.Bool("c2e", false, "translate from chinese to english")
var v = flag.Bool("v", false, "show version")

func main() {
	// 解析命令行参数，给三个命令行赋值
	// 如果命令行中有 -e2c，-c2e，-v等选项，相应选项标识被设置为 true。
	flag.Parse()

	if *v {
		println("UbuntuTranslator", year)
		println("版本：", version)
		println("作者：", author)
		println("邮件：", email)
		return
	}

	if flag.NArg() < 1 {
		println("Error: need more args")
		Usage()
		return
	}

	if *ce && *ec {
		println("Error: what's you mean?")
		Usage()
		return
	}

	var drection string
	switch {
	case *ce:
		drection = c2e
	case *ec:
		drection = e2c
	default:
		println("misson drection")
		Usage()
		return
	}

	// 获得用户要查询的词或句子
	var words string = ""
	for i := 0; i < flag.NArg(); i++ {
		if i > 0 && *ec {
			words += " "
		}
		words += flag.Arg(i)
	}

	words = strings.TrimSpace(words)
	// 对字符串进行编码，转换成URL格式，如把空格转换成 20%
	queryWords := http.URLEscape(words)

	url := baseUrl + queryWords + langpair + drection
	/*println(url)*/
	// 查询
	r, _, e := http.Get(url)
	if e != nil {
		println(e.String())
		return
	}
	// 获得返回的json数据
	b, e := ioutil.ReadAll(r.Body)
	r.Body.Close()
	if e != nil {
		println(e.String())
		return
	}

	// 解析json数据，并把结果存储在res中。
	var res Response
	// json.Unmarshal(string(b), &res)
	json.Unmarshal(b, &res)

	if text := res.ResponseData.TranslatedText; len(text) > 0 {
		fmt.Printf("\n%s\n%s\n\n", words, text)
		return
	}

	println("Unkown error")
}

func Usage() {
	println(`
Usage:

UbuntuTranslator [ option ] <words>

option:

    -e2c translate from english to chinese.
    -c2e translate from chinese to english.
    -v   show version
    `)
}
