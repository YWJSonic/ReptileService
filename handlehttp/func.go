package handlehttp

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// ConnectPool HttpClient http get http request connect pool
func ConnectPool() *http.Client {
	if clientConnect == nil {
		clientConnect = new(httpClient)
		httptr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},

			MaxIdleConns:        50,
			MaxIdleConnsPerHost: 50,
		}
		clientConnect.Client = &http.Client{
			Transport: httptr,
		}
	}
	return clientConnect.Client
}

// HTTPGet ...
func HTTPGet(ip string, values map[string][]string) []byte {
	res, err := http.Get(ip)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return result
}

// HTTPGetRequest Http Raw Request
func HTTPGetRequest(client *http.Client, url string, value []byte) []byte {
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(value))
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTPPostRawRequest Resp", resp, " ---- ", err)
		return nil
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

// HTTPPostRequest ...
func HTTPPostRequest(ip string, values map[string][]string) []byte {
	// res, err := http.Post(ip, "application/x-www-form-urlencoded", strings.NewReader("name=cjb"))
	res, err := http.PostForm(ip, values)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	result, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return result

}

// PostData get http post data
func PostData(r *http.Request) map[string]interface{} {
	data := map[string]interface{}{}
	contentType := r.Header.Get("Content-type")

	if contentType == "application/x-www-form-urlencoded" {
		err := r.ParseForm()
		if err != nil {
			panic(err)
		}
		v := r.Form
		postdata := v.Get("POST")
		if err := json.Unmarshal([]byte(postdata), &data); err != nil {
			panic(err)
		}

	} else {
		d := json.NewDecoder(r.Body)
		err := d.Decode(&data)
		if err != nil {
			panic(err)
		}
	}

	return data
}

// HTTPPostRawRequest Http Raw Request
func HTTPPostRawRequest(client *http.Client, url string, value []byte) ([]byte, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(value))
	req.Header.Set("Content-Type", "application/json")
	fmt.Println("HTTPPostRawRequest", req)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("HTTPPostRawRequest Resp", resp)
		fmt.Println("HTTPPostRawRequest Error", err)
		resp.Body.Close()
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body, nil
}

// Option add header option
func Option(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	headers := w.Header()
	headers.Add("Access-Control-Allow-Origin", "*")
	headers.Add("Vary", "Origin")
	headers.Add("Vary", "Access-Control-Request-Method")
	headers.Add("Vary", "Access-Control-Request-Headers")
	headers.Add("Access-Control-Allow-Headers", "Content-Type, Origin, Accept, token")
	headers.Add("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.WriteHeader(http.StatusOK)
}
