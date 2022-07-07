## server项目结构

```
├── api
├── config
├── core
├── global
├── initialize
├── middleware
├── model
├── packfile
├── router
├── service
└── utils
```

| 文件夹       | 说明                    | 描述                                     |
| ------------ | ----------------------- |----------------------------------------|
| `api`        | api层                   | api层                                   |
| `config`     | 配置包                  | config.yaml对应的配置结构体                    |
| `core`       | 核心文件                | 核心组件(zap, viper, server)的初始化           |
| `global`     | 全局对象                | 全局对象                                   |
| `initialize` | 初始化 | router,redis,gorm,validator, timer的初始化 |
| `middleware` | 中间件层 | 用于存放 `gin` 中间件代码                       |
| `model`      | 模型层                  | 模型对应数据表以及数据库操作                         |
| `packfile`   | 静态文件打包            | 静态文件打包                                 |
| `router`     | 路由层                  | 路由层                                    |
| `service`    | service层               | 存放业务逻辑问题                               |
| `source` | source层 | 存放初始化数据的函数                             |
| `utils`      | 工具包                  | 工具函数封装                                 |

