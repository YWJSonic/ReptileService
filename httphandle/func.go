package httphandle

import (
	"net/http"

	"github.com/YWJSonic/ReptileService/httphandle/httpDriver"
	"github.com/julienschmidt/httprouter"
)

// HTTPGet ...
func (self *HttpHandle) HTTPGet(ip string, values map[string][]string) []byte {
	return httpDriver.HTTPGet(ip, values)
}

// HTTPGetRequest Http Raw Request
func (self *HttpHandle) HTTPGetRequest(url string, value []byte) []byte {
	return httpDriver.HTTPGetRequest(self.client, url, value)
}

// HTTPPostRequest ...
func (self *HttpHandle) HTTPPostRequest(ip string, values map[string][]string) []byte {
	return httpDriver.HTTPPostRequest(ip, values)
}

// PostData get http post data
func (self *HttpHandle) PostData(r *http.Request) map[string]interface{} {
	return httpDriver.PostData(r)
}

// HTTPPostRawRequest Http Raw Request
func (self *HttpHandle) HTTPPostRawRequest(url string, value []byte) ([]byte, error) {
	return httpDriver.HTTPPostRawRequest(self.client, url, value)
}

// Option add header option
func (self *HttpHandle) Option(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.WriteHeader(http.StatusOK)
}
