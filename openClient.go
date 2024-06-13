package open_api_client_sdk_go

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type OpenClient struct {
	appKey    string
	appSecret string
	Nonce     int
	Timestamp string
	Sign      string
	Parameter string
}

type GetNameByGetResponseData struct {
	Name string `json:"name"`
}

func (oc OpenClient) GetNameByGet(name string) {
	// 创建 URL 并添加 Query 参数
	baseURL, err := url.Parse("http://127.0.0.1:9002/api/user/get")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	params := url.Values{}
	params.Add("name", name)
	baseURL.RawQuery = params.Encode()

	// 发送 GET 请求
	resp, err := http.Get(baseURL.String())
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 解析 JSON 响应数据
	var data GetNameByGetResponseData
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 输出响应数据
	fmt.Printf("Name: %s\n", data)

}

type GetNameByPostRequestData struct {
	Name string `json:"name"`
}

type GetNameByPostResponseData struct {
	Name string `json:"name"`
}

func (oc OpenClient) GetNameByPost(name string) {
	// 创建要发送的 JSON 数据
	data := GetNameByPostRequestData{Name: name}
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 创建 HTTP POST 请求
	req, err := http.NewRequest("POST", "http://127.0.0.1:9002/api/user/post", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// 使用 http.Client 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// 解析 JSON 响应数据
	var responseData GetNameByPostResponseData
	err = json.NewDecoder(resp.Body).Decode(&responseData)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// 输出解析后的数据
	fmt.Printf("Name: %s\n", responseData)
}
