package utility

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Godoc HTTPPostJSON
// ref: https://blog.logrocket.com/making-http-requests-in-go/
func HTTPPostJSON(url string, postBody interface{}) (string, error) {
	var result string
	contentType := "application/json"
	postBodyAsBytes, err := json.Marshal(postBody)
	if err != nil {
		return result, err
	}
	postBodyAsBuff := bytes.NewBuffer(postBodyAsBytes)

	//* It important for set timeout
	//? ref: https://medium.com/@nate510/don-t-use-go-s-default-http-client-4804cb19f779
	httpClient := &http.Client{
		Timeout: 30 * time.Second,
	}
	resp, err := httpClient.Post(url, contentType, postBodyAsBuff)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	result = string(body)
	return result, err
}
