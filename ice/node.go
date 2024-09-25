package ice

import "github.com/google/uuid"

var NodeId string

func GetSelfNodeId() string {
	return "0"
}

func GetMasterNodeId() string {
	return "0"
}

// GenerateNodeId 生成一个新的节点ID，节点ID是一个字符串形式的UUID
func GenerateNodeId() string {
	return uuid.NewString()
}
