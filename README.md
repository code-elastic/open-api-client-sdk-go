# open-api-client-sdk-go
open API invoke client for Go

## quickstart

1. install

```shell
go get -u https://github.com/code-elastic/open-api-client-sdk-go
```

2. use case

```go
func TestInvoke(c *gin.Context) {
	//
	body, err := c.GetRawData()
	
	var interfaceInfo model.InterfaceInfo
	// TODO 给 interfaceInfo 赋值
	info, err := json.Marshal(&interfaceInfo)
	
	client, err := client.NewClient(body, info)
	if err != nil {
		c.JSON(http.StatusForbidden, utils.ResponseError(utils.ParamsError, "调用失败"))
		return
	}
	response := client.DoRequest()
	var res utils.Response
	err = json.Unmarshal(response, &res)
	if err != nil {
		fmt.Println("Unmarshal Error", err.Error())
		return
	}
	if res.Code != 0 {
		c.JSON(http.StatusForbidden, res)
		return
	}
	c.JSON(http.StatusOK, utils.ResponseOK(res.Data))
}
```

