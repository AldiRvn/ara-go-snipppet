type HTTP struct{}

func (this HTTP) bodyBuilder(requestBody interface{}) io.Reader {
	requestBodyAsByte, err := json.Marshal(requestBody)
	if err != nil {
		return nil
	}

	return bytes.NewBuffer(requestBodyAsByte)
}
func (this HTTP) httpRequest(method, url string, requestBody interface{}, addHeader map[string]interface{}) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	
	var bodyToRequest io.Reader
	if requestBody != nil {
		bodyToRequest = this.bodyBuilder(requestBody)
	}
	req, err := http.NewRequest(method, url, this.bodyBuilder(requestBody))
	if err != nil {
		return nil, err
	}

	req.Header = http.Header{
		"Content-Type": []string{"application/json"},
	}
	for key, value := range addHeader {
		req.Header.Set(key, fmt.Sprint(value))
	}
	// log.Println(req.Header)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(res.Body)
}
func (this HTTP) POST(url string, requestBody interface{}, addHeader map[string]interface{}) ([]byte, error) {
	return this.httpRequest("POST", url, requestBody, addHeader)
}
func (this HTTP) GET(url string, addHeader map[string]interface{}) ([]byte, error) {
	return this.httpRequest("GET", url, map[string]interface{}{}, addHeader)
}

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
