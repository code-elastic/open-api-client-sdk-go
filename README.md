# open-api-client-sdk-go
open API invoke client for Go

 客户端 SDK - Go 实现

## quickstart

1. install

```shell
go get -u https://github.com/code-elastic/open-api-client-sdk-go
```

2. use case

```go
func useCase() {
	var client *open_api_client_sdk_go.OpenClient
	client = open_api_client_sdk_go.NewOpenClient("qwert123456", "asdfgqwertzxcvb")
	client.GetNameByPost("huang")
}
```

