package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type args struct {
	frequency int //秒
	url       string
}

var (
	loger *log.Logger
)

func Get(url string, params map[string]string, headers map[string]string) {
	//new request
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		loger.Println(err)
		return
	}
	//add params
	q := req.URL.Query()
	if params != nil {
		for key, val := range params {
			q.Add(key, val)
		}
		req.URL.RawQuery = q.Encode()
	}
	//add headers
	if headers != nil {
		for key, val := range headers {
			req.Header.Add(key, val)
		}
	}
	//http client
	client := &http.Client{}

	resp, err := client.Do(req)
	if resp != nil {
		all, _ := ioutil.ReadAll(resp.Body)
		if all != nil {
			strBody := string(all)
			loger.Printf("Go %s URL : %s ,Body : %s\n", http.MethodGet, req.URL.String(), strBody)
		} else {
			loger.Printf("Go %s URL : %s ,result is null\n", http.MethodGet, req.URL.String())
		}
	}
}

func (a *args) parseFlags() {

	//请求频率 实际请求频率还需要考虑接口的rt，如果rt大雨f，则只能按照rt的时间频率来执行，因为目前是单线程的。
	f := flag.Int("f", 10, "frequency eg: 10")

	//请求地址
	url := flag.String("url", "http://192.168.118.71:11001/mocksuccess/v1/200", "url,eg: http://192.168.118.71:11001/mocksuccess/v1/200")

	//一定要有Parse，不然上面的flag.xx()都不起作用
	flag.Parse()

	a.frequency = *f
	a.url = *url
	fmt.Printf("url : %v , frequency : %v\n", a.url, a.frequency)
}

func main() {
	args := args{}
	args.parseFlags()
	file := "./req.log"
	logFile, _ := os.OpenFile(file, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	loger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
	loger.Printf("[Begin] url : %v , frequency : %v\n", args.url, args.frequency)
	for {
		time.Sleep(time.Second * time.Duration(args.frequency))
		Get(args.url, nil, nil)
	}
}
