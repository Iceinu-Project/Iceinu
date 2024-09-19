# 🧊Iceinu 氷犬

![Go Badge](https://img.shields.io/badge/Go-1.22%2B-cyan?logo=go)

氷犬Iceinu 是一个多用途的Go语言聊天机器人框架(主要为NTQQ设计)，可以将其作为开发套件来进行二次开发，亦或者作为库按需引入来快速编写自己的聊天机器人（暂时没有实现）。

🚧暂时还在~~画饼~~施工中，晚点再来探索吧~

## 开发进度

目前距离第一个Release版本的发布还有很多需要进行的工作：
- [x] 基础的事件总线系统
- [x] 内置LagrangeGo适配器的部分事件发送
- [x] 类Satori的消息解析
- [x] 配置读取/生成/补全
- [ ] 数据库连接，统一数据库接口设计
- [ ] 消息发送/统一Client设计
- [ ] 插件系统设计和实现
- [ ] 完善内置LagrangeGo适配器的事件接收和处理
- [ ] 集群模式设计和实现（集群的动态总控和静态总控模式）
- [ ] 排障和性能优化
- [ ] 自动化测试
- [ ] 确定各项基础程序设计，编写使用文档

## 特性

- 基于Go开发,性能表现良好
- 基于事件驱动的消息推送机制
- 类Satori的事件和消息系统，可以直接构建Satori标准的XHTML作为消息内容
- 模块化设计，适配多平台，可自由开发插件

## 直接部署

（目前仍然属于开发前期，部署了也暂时没什么用处）

访问Github Action就可以获取Iceinu的自动构建二进制文件，Iceinu默认集成了`LagrangeGo`所以无需再配置onebot协议连接，第一次启动时会自动检测并生成配置文件，完成配置之后在命令行中输入回车即可开始运行。

你可以参照Iceinu数据库配置指南来配置Iceinu使用的PostgreSQL数据库。

Iceinu可以以分布式模式进行部署，通过在启动时附加`--node`参数即可启动Iceinu的子节点模式，在子节点的配置文件中设置主节点的ip即可进行连接。

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