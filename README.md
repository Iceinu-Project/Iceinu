# <img src="./static/img/icedog.png" alt="logo" style="height: 6vh; position: relative; top: 1.2vh; max-width: 100%;"> Iceinu 氷犬

![Go Badge](https://img.shields.io/badge/Go-1.22%2B-cyan?logo=go)

Iceinu是使用Go编写的分布式QQ聊天机器人，可以直接拉取进行部署，也可以作为开发框架在其基础上自行增减功能和开发插件。

🚧暂时还在施工中，晚点再来探索吧~

## 直接部署

访问Github Action就可以获取Iceinu的自动构建二进制文件，Iceinu默认集成了`Lagrange Go SDK`所以无需再配置onebot协议连接，第一次启动时会自动检测并生成配置文件，完成配置之后在命令行中输入回车即可开始运行。

你可以参照Iceinu数据库配置指南来配置Iceinu使用的PostgreSQL数据库。

Iceinu可以以分布式模式进行部署，通过在启动时附加`--node`参数即可启动Iceinu的子节点模式，在子节点的配置文件中设置主节点的ip即可进行连接。

还可以通过`--multi`参数进行本地多节点模式部署，实现单主机同时连接多个账号。

## 二次开发

Go语言的静态特性让它非常不怎么适合传统意义上的模块化加载，所以Iceinu并没有也不会实现从外部进行的插件加载。

不过Iceinu通过接口定义了内部插件的实现，直接拉取代码跟着插件文档进行二次开发即可扩展更多的插件功能。

```shell
git clone git@github.com:Iceinu-Project/iceinu.git
```

## 鸣谢

- [ZeroBot](https://github.com/wdvxdr1123/ZeroBot) | 基于onebot协议的Golang聊天机器人开发框架