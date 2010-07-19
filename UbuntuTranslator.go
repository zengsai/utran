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
 * 改变内容：如果用户在 Dict 中没有查询到单词，尝试通过 Google 翻译返回结果。
 * 之前版本：1.1.0
 * 修改者：红猎人
 * 日期：2010-07-19
 *
 * 改变内容：如果用户查询的是单词，则调用词海的字典API，返回更详细的解析，只支持中英互译。
 *           如果用户查询的是短语或句子，则调用Google的翻译API，返回翻译结果。
 * 之前版本：1.0.3
 * 修改者：红猎人
 * 日期：2010-06-30
 *
 * 当前版本： 1.1.1
 */
package main

import (
	"http"
	"json"
	"fmt"
	"io/ioutil"
	"flag"
	"strings"
	"xml"
)

const (
	e2c           = "en|zh"
	c2e           = "zh|en"
	googleBaseUrl = "http://ajax.googleapis.com/ajax/services/language/translate?v=1.0&q="
	langpair      = "&langpair="
	dictBaseUrl   = "http://api.dict.cn/ws.php?utf8=true&q="
	version       = "v 1.1.1"
	year          = "2010"
	author        = "红猎人"
	email         = "zengsai@gmail.com"
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

// Dict Struct
type Sent struct {
	Orig  string
	Trans string
}
type DictResult struct {
	XMLName xml.Name "dict"
	Key     string
	Lang    string
	Audio   string
	Pron    string
	Def     string
	Sent    []Sent
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

	var isWord bool = true
	var url string = ""
	var words string = ""

	if flag.NArg() > 1 {
		isWord = false
	}

req:
	if isWord {
		// dict
		url = dictBaseUrl + flag.Arg(0)
	} else {
		// 获得用户要查询的词或句子
		for i := 0; i < flag.NArg(); i++ {
			if i > 0 && *ec {
				words += " "
			}
			words += flag.Arg(i)
		}

		words = strings.TrimSpace(words)
		// 对字符串进行编码，转换成URL格式，如把空格转换成 20%
		queryWords := http.URLEscape(words)

		url = googleBaseUrl + queryWords + langpair + drection
	}
	/*println(url)*/
	// 查询
	r, _, e := http.Get(url)
	if e != nil {
		println(e.String())
		return
	}

	if isWord {
		var res DictResult
		xml.Unmarshal(r.Body, &res)
		r.Body.Close()

		if res.Def != "Not Found" {
			fmt.Printf("\n原词：%s\t发音:[%s]\n\n解释\n***********************************\n%s\n",
				res.Key, res.Pron, res.Def)

			if len(res.Sent) > 0 {
				fmt.Printf("\n例句\n***********************************\n")
				for _, sent := range res.Sent {
					fmt.Printf("%s\n%s\n\n", sent.Orig, sent.Trans)
				}
			}

			return
		} else {
			fmt.Printf("dict 没有找到该词, 使用 Google 翻译\n")
            isWord = false
            goto req
		}
	} else {
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