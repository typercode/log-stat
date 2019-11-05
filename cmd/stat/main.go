package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type appAuthId struct {
	AppId        string `json:"appId"`
	DeviceId     string `json:"deviceId"`
	ApiVersionId string `json:"api-version-id"`
}

func isIn(l []string, o string) (result bool) {
	for _, v := range l {
		if o == v {
			return true
		}
	}
	return false
}

func (a *app) genMailContent() (mailContent string) {
	for k, v := range a.m {
		mailContent += "appId: " + k + "\n"
		mailContent += "api-version-id: "
		for _, v1 := range v {
			mailContent += v1 + ","
		}
		mailContent += "\n\n"
	}
	fmt.Println("result:")
	fmt.Println(mailContent)
	return
}

func (a *app) sendMail(content string) {
	// 空实现
}

type app struct {
	mailList string
	filePath string
	m        map[string][]string
}

func (a *app) parseFlags() {

	//文件路径，如果没有，默认取当前路径
	p := flag.String("p", "", "set the log file path(eg: /path/to/file/logs/app-name)")

	//文件名称
	f := flag.String("f", "", "set the log file name(eg: Log.log.2019-11-03)")

	//文件名前缀，因为一般的日志都已一个相同的文件前缀，当指定文件名时忽略
	fp := flag.String("fp", "MSSM-Auth.log.", "set the log file prefix,Ignore if f is set. (eg: MSSM-Auth.log.)")

	mailList := flag.String("ma", "", "set the mail address (eg: xxx@xxx.com)")

	//一定要有Parse，不然上面的flag.xx()都不起作用
	flag.Parse()

	a.mailList = *mailList

	path := *p

	if *p == "." || *p == "" {
		path, _ = os.Getwd()
	}

	var filePath = ""

	if *f == "" {
		format := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
		filePath = path + "/" + *fp + format
	} else {
		filePath = path + "/" + *f
	}

	fmt.Println("filePath: " + filePath)
	a.filePath = filePath
}

func (a *app) parseLog() {
	inputFile, inputError := os.Open(a.filePath)
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}

	//核心数据结构
	var m = make(map[string][]string)

	inputReader := bufio.NewReader(inputFile)

	lineNumber := 1

	lineStart := "AuthenticationController.customerVerify(..)="

	lineEnd := "^_^"

	for {
		inputString, readerError := inputReader.ReadString('\n')
		if strings.Contains(inputString, lineStart) {
			index := strings.Index(inputString, lineStart)
			content := inputString[index+len(lineStart):]
			content = content[:strings.Index(content, lineEnd)]
			//fmt.Println(content)
			appAuthIds := make([]appAuthId, 0)
			if err := json.Unmarshal([]byte(content), &appAuthIds); err == nil {
				//如果存在则添加，否则先新建，再添加
				for _, v := range appAuthIds {
					if !isIn(m[v.AppId], v.ApiVersionId) {
						s1 := append(m[v.AppId], v.ApiVersionId)
						m[v.AppId] = s1
					}
				}
			} else {
				fmt.Println(err)
			}
			lineNumber = lineNumber + 1
		}
		if readerError == io.EOF {
			break
		}
	}
	a.m = m
}

func main() {
	ap := app{}
	ap.parseFlags()
	ap.parseLog()
	ap.sendMail(ap.genMailContent())
}
