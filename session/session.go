package session

import (
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

// Session HTTP 会话
type Session struct {
	client    *http.Client
	UserAgent string
}

// New 初始化 HTTP 会话
func New() *Session {
	s := new(Session)
	s.client = &http.Client{}
	s.UserAgent = "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.92 Safari/537.36"
	jar, _ := cookiejar.New(nil)
	s.client.Jar = jar
	return s
}

// request 请求
func (s *Session) request(method, url string, data url.Values) ([]byte, error) {
	req, err := http.NewRequest(method, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", s.UserAgent)
	if method == "GET" {
		req.Header.Set("Content-Type", "text/html; charset=utf-8")
	}
	if method == "POST" {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	res, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}

// SetProxy 设置代理
func (s *Session) SetProxy(urlstr string) {
	if urlstr == "" {
		s.client.Transport = &http.Transport{
			Proxy: nil,
		}
	} else {
		urlproxy, err := url.Parse(urlstr)
		if err != nil {
			panic(err)
		}
		s.client.Transport = &http.Transport{
			Proxy: http.ProxyURL(urlproxy),
		}
	}
}

// Get Get
func (s *Session) Get(url string) ([]byte, error) {
	return s.request("GET", url, nil)
}

// Post Post
func (s *Session) Post(url string, data url.Values) ([]byte, error) {
	return s.request("POST", url, data)
}
