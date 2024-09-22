# 🧊Iceinu 氷犬

![Go Badge](https://img.shields.io/badge/Go-1.22%2B-cyan?logo=go)
[![workflow](https://github.com/Iceinu-Project/iceinu/actions/workflows/go.yml/badge.svg)](https://github.com/Iceinu-Project/iceinu/actions)
[![goreportcard](https://img.shields.io/badge/go%20report-A+-brightgreen.svg?style=flat)](https://goreportcard.com/report/github.com/Iceinu-Project/iceinu)
[![QQGroup Badge](https://img.shields.io/badge/QQ群-970801565-blue?)](https://qm.qq.com/q/93crfU39ny)

氷犬Iceinu 是一个多用途的Go语言聊天机器人框架，可以将其作为开发套件来进行二次开发，亦或者作为库按需引入来快速编写自己的聊天机器人（暂时没有实现）。

🚧暂时还在~~画饼~~施工中，晚点再来探索吧~

## 开发进度

Refact分支正在进行重构，目前进度还在推进，过几天再来探索吧~

## 特性

~~哪里是特性，完全是画饼，一半都还没实现完~~

- 基于Go开发,性能表现良好
- 基于统一事件驱动的消息推送机制
- 以Satori作为基础实现了统一消息标准
- 可直接发送Satori标准的XHTML消息
- 模块化适配器设计
- 动态/静态集群，跨平台集群
- 完整的动态权限管理系统
- 在插件间共享数据库连接池
- 类Alconna的命令解析器
- 可配置自动向指定用户/频道发送日志
- 主动信息推送/订阅机制
- 从HTML+CSS模板渲染图片（基于wkhtmltoimage集成，未来可能会实现）

## 直接部署

（目前仍然属于开发前期，部署了也暂时没什么用处）

访问[Github Action](https://github.com/Iceinu-Project/iceinu/actions)就可以获取Iceinu的自动构建二进制文件，Iceinu默认集成了`LagrangeGo`所以无需再配置onebot
协议连接，第一次启动时会自动检测并生成配置文件，完成配置之后在命令行中输入回车即可开始运行。

你可以参照Iceinu数据库配置指南来配置Iceinu使用的PostgreSQL数据库。

Iceinu在设计上支持集群部署，且支持动态组网式集群（需要各个Bot实例之间可以相互访问）和静态总控式集群（需要一个Bot实例作为总控，这个实例本身不能下线）

## 二次开发

（文档还没写，Release之前会开始编写文档）

Go语言的静态特性让它非常不怎么适合传统意义上的模块化加载，所以Iceinu并没有也不会实现从外部进行的插件加载。

不过Iceinu通过接口定义了内部插件的实现，直接拉取代码跟着插件文档进行二次开发即可扩展更多的插件功能。

```shell
git clone git@github.com:Iceinu-Project/iceinu.git
```

## 鸣谢

- [Lagrange.Core](https://github.com/LagrangeDev/Lagrange.Core) NTQQ通信协议的C#实现
- [LagrangeGo](https://github.com/LagrangeDev/LagrangeGo) Lagrange.Core的Go实现
- [LagrangeGo-Teamplate](https://github.com/ExquisiteCore/LagrangeGo-Template) LagrangeGo的模板示例项目
- [go-cqhttp](https://github.com/Mrs4s/go-cqhttp) 基于 Mirai 以及 MiraiGo 的 OneBot Golang 原生实现
- [ZeroBot](https://github.com/wdvxdr1123/ZeroBot) 基于onebot协议的Golang聊天机器人开发框架
- [Logrus](https://github.com/sirupsen/logrus) 强大好用的Go日志库