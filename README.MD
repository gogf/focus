# Golang社区系统
```
├── app
│ ├── dao
│ ├── model
│ ├── shared
│ └── system
│     ├── admin
│     │ └── internal
│     │     ├── api
│     │     ├── define
│     │     └── service
│     └── index
│         └── internal
│             ├── api
│             ├── define
│             └── service
├── config
├── document
├── library
├── packed
├── public
├── template
├── upload
├── Dockerfile
├── go.mod
└── main.go
```

|目录/文件名称   | 说明 | 描述
|---|---|---
|`app`           | 业务逻辑层 | 所有的业务逻辑存放目录。
| - `dao`        | 数据访问   | 数据库的访问操作，仅包含最基础的数据库CURD方法。
| - `model`      | 数据模型   | 存放数据相关的实体结构定义。
| - `system`     | 系统模块   | 内部可能包含多个子系统，不同子系统之间资源相互隔离。
| &nbsp; &nbsp; - `admin`    | 管理后台 | 子系统，管理后台。
| &nbsp; &nbsp; &nbsp; &nbsp; - `internal`       | 内部模块 | 系统内部模块，仅供当前内部系统调用，无法在系统间共享。
| &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - `api`     | 业务接口 | 系统内部接收/解析用户输入参数的入口/接口层
| &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - `define`  | 结构定义 | 系统内部的输入、输出数据结构定义。
| &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; - `service` | 逻辑封装 | 系统内部业务逻辑封装，实现特定的业务需求。
|`config`        | 配置管理   | 所有的配置文件存放目录。
|`docker`        | 镜像文件   | Docker镜像相关依赖文件，脚本文件等等。
|`document`      | 项目文档   | Documentation项目文档，如: 设计文档、帮助文档等等。
|`library`       | 公共库包   | 公共的功能封装包，往往不包含业务需求实现。
|`packed`        | 打包目录   | 将资源文件打包的`Go`文件存放在这里，`boot`包初始化时会自动调用。
|`public`        | 静态目录   | 仅有该目录下的文件才能对外提供静态服务访问。
|`template`      | 模板文件   | MVC模板文件存放的目录。
|`Dockerfile`    | 镜像描述 | 云原生时代用于编译生成Docker镜像的描述文件。
|`go.mod`        | 依赖管理   | 使用`Go Module`包管理的依赖描述文件。
|`main.go`       | 入口文件   | 程序入口文件。