package open_api_client_sdk_go

import (
	"encoding/json"
	"fmt"
)

// SerializeToJSON 将任何数据类型转换为 JSON 格式的字符串。
func SerializeToJSON(data interface{}) (string, error) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return "", fmt.Errorf("error marshaling JSON: %w", err)
	}
	return string(jsonData), nil
}
