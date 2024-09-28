package ice

var PluginVerifier string

// GetPluginVerifier 获取插件的校验值，用于确认各个实例之间的插件版本是否一致
//
// 校验是可选的，在氷犬的消息事件系统中实例间的插件是否相同实际上并不影响组网，但是如果插件不同可能会导致一些意想不到的问题
func GetPluginVerifier(isVerify bool) string {
	if !isVerify {
		PluginVerifier = "pass"
	} else {
		VerifyPlugin()
	}
	return PluginVerifier
}

// VerifyPlugin 校验插件是否一致，暂时是空函数
func VerifyPlugin() {

}
