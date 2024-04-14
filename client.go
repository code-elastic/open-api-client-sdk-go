package open_api_client_sdk_go

import (
	"encoding/json"
	"fmt"
	"github.com/code-elastic/open-api-client-sdk-go/model"
	"github.com/code-elastic/open-api-client-sdk-go/utils"
	"math/rand"
	"strconv"
	"time"
)

type Client struct {
	Header        model.Header        // Header 封装了用于API请求的头信息，如访问密钥、签名等
	InterfaceInfo model.InterfaceInfo // InterfaceInfo 包含特定API接口的相关信息，如URL、方法等
}

func NewClient(body, info []byte) (*Client, error) {
	var invokeInterface model.InvokeInterface
	var interfaceInfo model.InterfaceInfo
	err := json.Unmarshal(body, &invokeInterface) // 从body解析认证信息
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(info, &interfaceInfo) // 从info解析接口信息
	if err != nil {
		return nil, err
	}

	fmt.Println("new client url: ", interfaceInfo.Url) // 输出解析后的接口URL，用于调试

	return &Client{
		Header: model.Header{
			AccessKey: invokeInterface.AccessKey,                                      // 设置AccessKey
			Nonce:     rand.Int(),                                                     // 生成一个随机的Nonce值
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),                       // 设置当前时间的时间戳
			Sign:      utils.GenSign(invokeInterface.Body, invokeInterface.SecretKey), // 生成签名
			Body:      invokeInterface.Body,                                           // 设置请求体
		},
		InterfaceInfo: interfaceInfo, // 设置接口信息
	}, nil
}
