package main

import (
	"gtihub.com/Iceinu-Project/iceinu/adapter"
	"gtihub.com/Iceinu-Project/iceinu/adapter/lagrange"
	"gtihub.com/Iceinu-Project/iceinu/ice"
	"gtihub.com/Iceinu-Project/iceinu/logger"
)

func main() {
	logger.Infof("正在启动Iceinu，请稍等...")
	// 自定义事件总线中间件
	ice.Bus.AddMiddleware(func(eventType string, payload interface{}) {
		logger.Debugf("事件推送：%s", eventType)
	})

	// 初始化适配器，默认使用LagrangeGo适配器
	var a adapter.IceAdapter
	a = &lagrange.AdapterLagrange{}
	a.Init()

	select {
	// 阻塞主线程
	}
}
