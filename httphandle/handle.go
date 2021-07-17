package httphandle

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

type IHttpHandle interface {
	HTTPGet(ip string, values map[string][]string) []byte
	HTTPGetRequest(url string, value []byte) []byte
	HTTPPostRequest(ip string, values map[string][]string) []byte
	PostData(r *http.Request) map[string]interface{}
	HTTPPostRawRequest(url string, value []byte) ([]byte, error)
	Option(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type HttpHandle struct {
	client             *http.Client
	getRequestWaitTime time.Time
}

func NewHttpHandle() *HttpHandle {
	HttpHandle := &HttpHandle{
		getRequestWaitTime: time.Time{},
	}
	httptr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},

		MaxIdleConns:        50,
		MaxIdleConnsPerHost: 50,
	}
	HttpHandle.client = &http.Client{
		Transport: httptr,
	}
	return HttpHandle
}

var Instans IHttpHandle
