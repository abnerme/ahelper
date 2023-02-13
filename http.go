package ahelper

import (
	"crypto/tls"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

type Http_an struct{}

func newHttp() Http_an {
	return Http_an{}
}

type Option struct {
	f func(*options)
}

type proxy struct {
	ip   string
	port int
}

type options struct {
	method   string
	postdata string
	headers  map[string]string
	proxy    proxy
}

type Resp_an struct {
	Body      []byte
	Str       string
	Cookies   string
	ReloadUrl string
}

func WithPostData(post string) Option {
	return Option{func(op *options) {
		op.postdata = post
	}}
}

func WithHeaders(headers string) Option {
	return Option{func(op *options) {
		before := headers
		reg := regexp.MustCompile(`( )+|(\n)+|(\t)+`)
		headers := reg.ReplaceAllString(before, "$1$2") //只保留一个空格一个换行 去除tab
		fmt.Println(headers)
		headers_m := make(map[string]string)
		a := strings.Split(headers, "\n")
		ak := make([]string, len(a[:]))
		av := make([]string, len(a[:]))
		//要用copy复制值；若用等号仅表示指针，会造成修改ak也就是修改了av
		copy(ak, a[:])
		copy(av, a[:])
		//fmt.Println(ak[0], av[0])
		for k, v := range ak {
			i := strings.Index(v, ":")
			j := i + 1
			if i > 0 {
				ak[k] = v[:i]
				av[k] = v[j:]
			}

			//设置Header
			if len(av[k]) > 0 {
				headers_m[ak[k]] = av[k]
			}

		}

		op.headers = headers_m
	}}
}
func WithProxy(proxy proxy) Option {
	return Option{func(op *options) {
		op.proxy.ip = proxy.ip
		op.proxy.port = proxy.port
	}}
}

func (h Http_an) Request(url string, ops ...Option) (*Resp_an, error) {
	//初始化默认请求参数
	opt := &options{}
	opt.method = "GET"
	headers := make(map[string]string)
	headers["User-Agent"] = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/105.0.0.0 Safari/537.36"
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	opt.headers = headers

	//fmt.Println(opt)
	for _, do := range ops { //动态修改
		do.f(opt)
	}
	//这里请注意，使用 InsecureSkipVerify: true 来跳过证书验证` 普通  client := &http.Client{}

	client := &http.Client{Transport: &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}}
	req, err := http.NewRequest(strings.ToUpper(opt.method), url, strings.NewReader(opt.postdata))
	if err != nil {
		fmt.Println(err.Error())
	}
	for k, v := range opt.headers { //设置请求头
		req.Header.Set(k, v)
	}
	resp, err := client.Do(req) //发送请求
	if err != nil {
		fmt.Println(err.Error())
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}

	cookies := ""
	for i := range resp.Cookies() { //拼接返回cookie
		k := resp.Cookies()[i].Name
		v := resp.Cookies()[i].Value
		//fmt.Println(k, v)
		cookies += k + "=" + v + ";"
	}
	//返回常用
	resp_an := &Resp_an{}
	resp_an.Body = body
	resp_an.Cookies = cookies
	resp_an.Str = string(body)
	resp_an.ReloadUrl = resp.Header.Get("Location")

	defer resp.Body.Close()
	defer client.CloseIdleConnections() // 如果我们重复发送 HTTP 请求时没有调用这个方法来关闭空闲连接，会导致内存泄漏？
	return resp_an, err

}
