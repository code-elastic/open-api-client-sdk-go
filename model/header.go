package model

type Header struct {
	AccessKey string // AccessKey 是用于访问API的认证密钥
	Nonce     int    // Nonce 是一个随机数，用于防止重放攻击
	Timestamp string // Timestamp 是请求的时间戳，通常用于防止请求过时
	Sign      string // Sign 是请求的签名，用于验证请求的完整性和来源
	Body      string // Body 是请求的正文，通常包含具体的请求数据
}
