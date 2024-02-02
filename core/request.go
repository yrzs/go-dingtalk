package core

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Request 发起请求
func Request(req *http.Request, v interface{}) error {
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	b, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("status = %d, msg = %s", response.StatusCode, string(b))
	}
	fmt.Println("请求响应体", string(b))
	if err := json.Unmarshal(b, v); err != nil {
		return err
	}

	return nil
}
