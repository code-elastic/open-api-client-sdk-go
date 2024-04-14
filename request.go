package open_api_client_sdk_go

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func (c *Client) DoRequest() []byte {
	// 创建一个http客户端
	client := &http.Client{} // 实例化HTTP客户端

	fmt.Println("client url: ", c.InterfaceInfo.Url) // 打印请求的URL，用于调试

	// 创建HTTP请求
	newRequest, err := http.NewRequest(c.InterfaceInfo.Method, c.InterfaceInfo.Url, nil) // 根据方法、URL创建新的请求
	if err != nil {
		fmt.Println("[request err] NewRequest ", err.Error()) // 打印创建请求的错误
		return nil
	}

	// 设置HTTP请求头
	newRequest.Header.Set("accessKey", c.Header.AccessKey)   // 设置AccessKey
	newRequest.Header.Set("sign", c.Header.Sign)             // 设置签名
	newRequest.Header.Set("timestamp", c.Header.Timestamp)   // 设置时间戳
	newRequest.Header.Set("nonce", strconv.Itoa(rand.Int())) // 设置随机数Nonce
	newRequest.Header.Set("body", c.Header.Body)             // 设置请求体内容

	// 执行HTTP请求
	response, err := client.Do(newRequest) // 发送请求并获取响应
	if err != nil {
		fmt.Println("[do request err] client.Do ", err.Error()) // 打印执行请求的错误
		return nil
	}
	defer func(Body io.ReadCloser) { // 确保响应体关闭
		err := Body.Close()
		if err != nil {
			fmt.Println("Body 关闭失败！", err.Error())
		}
	}(response.Body)

	content, err := io.ReadAll(response.Body) // 读取响应体内容
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	log.Println("Status: ", response.Status) // 打印响应状态
	log.Println("Body: ", string(content))   // 打印响应体

	return content // 返回响应内容
}
