package ice

import (
	"github.com/Iceinu-Project/Iceinu/config"
	"github.com/google/uuid"
)

var SelfNodeId string
var MasterNodeId string

func GetSelfNodeId() string {
	return SelfNodeId
}

func GetMasterNodeId() string {
	if !config.IceConf.Node.IsEnableNode {
		return GetSelfNodeId()
	}
	if config.IceConf.Node.IsMaster {
		return GetSelfNodeId()
	}
	return MasterNodeId
}

// SetMasterNodeId 设置主节点ID
func SetMasterNodeId(masterNodeId string) {
	MasterNodeId = masterNodeId
}

// SetSelfNodeId 设置自身节点ID
func SetSelfNodeId(selfNodeId string) {
	SelfNodeId = selfNodeId
}

// GenerateNodeId 生成一个新的节点ID，节点ID是一个字符串形式的UUID
func GenerateNodeId() string {
	return uuid.NewString()
}
