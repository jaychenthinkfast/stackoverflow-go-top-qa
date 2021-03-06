package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	preUrl       = "https://stackoverflow.com"
	preDir       = "../contents/"
	txt          string
	baseUrl      = "https://stackoverflow.com/questions/tagged/go?tab=votes&page=%s&pagesize=50"
	templateFile = "template.md"
	top100File   = "top100.md"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	startUrls := getStartUrls(baseUrl)
	logrus.Debug(startUrls)
	getListToMd(startUrls)
}

func getStartUrls(baseUrl string) []string {
	return []string{
		fmt.Sprintf(baseUrl, "1"),
		fmt.Sprintf(baseUrl, "2"),
	}
}

func getListToMd(urls []string) {
	//读取模板内容，后面用于组装
	template, err := ioutil.ReadFile(templateFile)
	if err != nil {
		logrus.Fatal(err)
	}
	for _, url := range urls {
		logrus.Debug(url)
		urlRes := httpGet(url)
		logrus.Debug(urlRes)
		dom, err := goquery.NewDocumentFromReader(strings.NewReader(urlRes))
		if err != nil {
			logrus.Fatal(err)
		}
		dom.Find("#mainbar .question-hyperlink").Each(func(i int, selection *goquery.Selection) {
			href := selection.AttrOr("href", "")
			title := selection.Text()
			txt += "* [" + title + "](" + preUrl + href + ")\n"
			logrus.Debug(preUrl+href, title)
			hrefSlice := strings.Split(href, "/")
			fileName := hrefSlice[len(hrefSlice)-1]
			content := string(template) + preUrl + href
			ioutil.WriteFile(preDir+fileName+".md", []byte(content), 777)
		})
	}
	ioutil.WriteFile(top100File, []byte(txt), 777)
}

func httpGet(url string) string {
	var content *http.Response
	var err error
	//最多循环10次
	for i := 0; i < 10; i++ {
		content, err = http.Get(url)
		if err != nil {
			logrus.Error(err)
		} else {
			break
		}
	}
	if err != nil {
		logrus.Fatal(err)
	}
	body, err := ioutil.ReadAll(content.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	return string(body)
}
