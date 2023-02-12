package ahelper

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)
type Http_an struct{}

func newHttp() Http_an{
	return Http_an{}
}
	

func (h Http_an) Post(url string, method string, postdata string, refer string, headers map[string]string) []byte {
	//这里请注意，使用 InsecureSkipVerify: true 来跳过证书验证` 普通  client := &http.Client{}

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}

	//client := &http.Client{}
	req, err := http.NewRequest(strings.ToUpper(method), url, strings.NewReader(postdata))
	if err != nil {
		// handle error
		fmt.Println(err.Error())
		return nil
	}

	req.Header.Set("Referer", refer)
	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
			//fmt.Println(k, v)
		}

	} else {
		req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.0.0 Safari/537.36 Edg/108.0.1462.42")
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := client.Do(req)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
		fmt.Println(err.Error())
	}
	//fmt.Println(string(body))
	//clipboard.ReadAll()
	//clipboard.WriteAll(string(body))
	defer client.CloseIdleConnections() // 如果我们重复发送 HTTP 请求时没有调用这个方法来关闭空闲连接，会导致内存泄漏？
	return body
}

func (h Http_an) Post_str(url string, method string, postdata string, refer string, headers map[string]string) string {

	return string(newHttp().Post(url, method, postdata, refer, headers))

}
