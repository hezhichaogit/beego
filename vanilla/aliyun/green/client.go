package green

import (
	"net/http"
	"encoding/json"
	"net/url"
	"strings"
	"io/ioutil"
	"net"
	"time"
)

type DefaultClient struct {
	Profile Profile;
}

func (defaultClient DefaultClient) GetResponse(path string, clinetInfo ClinetInfo, bizData BizData) string{
	clientInfoJson, _ := json.Marshal(clinetInfo)
	bizDataJson, _ := json.Marshal(bizData)
	
	var netTransport = &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}
	var client = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}
	//client := &http.Client{}
	req, err := http.NewRequest(method, host + path + "?clientInfo=" + url.QueryEscape(string(clientInfoJson)), strings.NewReader(string(bizDataJson)))

	if err != nil {
		// handle error
		return ErrorResult(err)
	} else {
		addRequestHeader(string(bizDataJson), req, string(clientInfoJson), path, defaultClient.Profile.AccessKeyId, defaultClient.Profile.AccessKeySecret)

		response, _ := client.Do(req)

		defer response.Body.Close()

		body, err := ioutil.ReadAll(response.Body)
		if(err != nil) {
			// handle error
			return ErrorResult(err)
		} else {
			return string(body)
		}
	}
}

type IAliYunClient interface {
	GetResponse(path string, clinetInfo ClinetInfo, bizData BizData) string
}
