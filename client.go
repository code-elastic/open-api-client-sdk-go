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
	Header        model.Header
	InterfaceInfo model.InterfaceInfo
}

func NewClient(body, info []byte) (*Client, error) {
	var invokeInterface model.InvokeInterface
	var interfaceInfo model.InterfaceInfo
	err := json.Unmarshal(body, &invokeInterface)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(info, &interfaceInfo)
	if err != nil {
		return nil, err
	}

	fmt.Println("new client url: ", interfaceInfo.Url)

	return &Client{
		Header: model.Header{
			AccessKey: invokeInterface.AccessKey,
			Nonce:     rand.Int(),
			Timestamp: strconv.FormatInt(time.Now().Unix(), 10),
			Sign:      utils.GenSign(invokeInterface.Body, invokeInterface.SecretKey),
			Body:      invokeInterface.Body,
		},
		InterfaceInfo: interfaceInfo,
	}, nil
}
