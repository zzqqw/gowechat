package requests

import (
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
)

var err error

type Requests struct {
	baseUrl string
	path    string
	req     []interface{}
	urls    *url.URL
	resty   *resty.Client
	resp    *resty.Response
}
type Media struct {
	filename string
	filesize int64
	stream   io.Reader
}

func NewRequests() *Requests {
	return &Requests{
		resty: resty.NewWithClient(&http.Client{}),
	}
}

// SetBaseURL 设置baseUrl
func (r *Requests) r() *resty.Request {
	return r.resty.R()
}

// SetBaseURL 设置baseUrl
func (r *Requests) SetBaseURL(url string) *Requests {
	r.baseUrl = url
	return r
}
func (r *Requests) SetPath(path string) *Requests {
	r.path = path
	return r
}

// SetGetReq 设置get请求参数
func (r *Requests) SetGetReq(req interface{}) *Requests {
	r.req = append(r.req, req)
	return r
}
func (r *Requests) buildQueryUrl() {
	u, err := url.Parse(r.baseUrl)
	if err != nil {
		panic(err)
	}
	u.Path = r.path
	var rawQuery string
	req := r.req
	for i := range req {
		values, _ := query.Values(req[i])
		rawQuery += values.Encode() + "&"
	}
	u.RawQuery = rawQuery
	r.urls = u
}

// GetForObject Get 发送get请求
func (r *Requests) GetForObject(object interface{}) error {
	r.buildQueryUrl()
	urls := r.urls.String()
	r.resp, err = r.get(urls)
	if err != nil {
		return err
	}
	bodyResp := r.resp.Body()
	logrus.Debug("requests.Get:" + urls + " response:" + string(bodyResp))
	err = json.Unmarshal(bodyResp, &object)
	return err
}

// PostJsonForObject GetForObject  发送post请求
func (r *Requests) PostJsonForObject(req interface{}, object interface{}) error {
	r.buildQueryUrl()
	urls := r.urls.String()
	r.resp, err = r.postJson(urls, req)
	if err != nil {
		return err
	}
	bodyResp := r.resp.Body()
	logrus.Debug("requests.PostJson:" + urls + " response:" + string(bodyResp))
	err = json.Unmarshal(bodyResp, &object)
	return err
}

// Get 发送get请求
func (r *Requests) get(url string) (*resty.Response, error) {
	return r.r().Get(url)
}

// PostJson   发送post请求
func (r *Requests) postJson(url string, req interface{}) (*resty.Response, error) {
	return r.r().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		Post(url)
}

// Upload 上传文件
func (r *Requests) Upload(url string, media Media) (*resty.Response, error) {
	return r.r().
		SetFileReader("media", media.filename, media.stream).
		Post(url)
}
