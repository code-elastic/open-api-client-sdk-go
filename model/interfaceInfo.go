package model

type InvokeInterface struct {
	AccessKey string `json:"access_key" binding:"required"` // AccessKey 是用于API认证的公钥，必填字段
	SecretKey string `json:"secret_key" binding:"required"` // SecretKey 是用于API认证的私钥，必填字段
	Body      string `json:"body"`                          // Body 是API请求的正文，可选字段
}
type InterfaceInfo struct {
	UserId         uint   `json:"user_id"`         // UserId 是使用此接口的用户ID
	Name           string `json:"name"`            // Name 是接口的名称
	Description    string `json:"description"`     // Description 是对接口功能的详细描述
	Url            string `json:"url"`             // Url 是接口的请求URL
	Method         string `json:"method"`          // Method 是请求使用的方法（如GET、POST等）
	RequestHeader  string `json:"request_header"`  // RequestHeader 是请求时使用的头部信息
	ResponseHeader string `json:"response_header"` // ResponseHeader 是响应时的头部信息
	RequestParams  string `json:"request_params"`  // RequestParams 是请求参数的详细描述
	Status         bool   `json:"status"`          // Status 表示接口的启用状态
}
