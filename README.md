# Golib Toolbox

Golib Toolbox 是一个 Go 语言工具库的元包，整合了多个高质量的基础组件库。

## 包含的组件

- **golib-conf** - 配置管理库，提供统一的配置文件管理、日志配置等功能
- **golib-mq** - 消息队列库，支持 Kafka、Redis、NATS 等多种消息队列
- **golib-toolkit** - 工具集库，包含数据库、OSS、SMTP、i18n 等常用工具

## 快速开始

### 一键引入所有组件

```bash
go get github.com/joepeak/golib-toolbox
```

### 使用示例

```go
package main

import (
    "context"
    "log"
    
    // 配置管理
    "github.com/joepeak/golib-conf"
    
    // 消息队列
    "github.com/joepeak/golib-mq/kafkamq"
    "github.com/joepeak/golib-mq/redismq"
    
    // 工具集
    "github.com/joepeak/golib-toolkit/database"
    "github.com/joepeak/golib-toolkit/oss"
)

func main() {
    // 初始化配置
    conf.Init()
    
    // 使用消息队列
    kafkaProducer := kafkamq.NewProducer()
    redisProducer := redismq.NewProducer()
    
    // 使用工具集
    db := database.NewMySQL()
    ossClient := oss.NewAliyunOSS()
    
    // ... 你的业务逻辑
}
```

## 按需使用

如果你只需要特定功能，也可以单独引入各个子包：

```bash
# 只需要配置管理
go get github.com/joepeak/golib-conf

# 只需要消息队列
go get github.com/joepeak/golib-mq

# 只需要工具集
go get github.com/joepeak/golib-toolkit
```

## 架构设计

Golib Toolbox 采用"元包"设计模式：

```
golib-toolbox/          # 元包（统一入口）
├── golib-conf/          # 配置管理（独立维护）
├── golib-mq/            # 消息队列（独立维护）
└── golib-toolkit/       # 工具集（独立维护）
```

### 优势

- **模块化**: 每个组件职责单一，可独立使用和维护
- **灵活性**: 按需引入，避免不必要的依赖
- **便利性**: 提供统一入口，简化项目依赖管理
- **兼容性**: 保持向后兼容，不影响现有项目

## 版本管理

- golib-conf: v0.3.1
- golib-mq: v0.3.1  
- golib-toolkit: v0.3.0

## 许可证

本项目采用 MIT 许可证。

## 贡献

欢迎提交 Issue 和 Pull Request。在贡献代码前，请确保：

1. 代码符合 Go 语言规范
2. 添加适当的测试
3. 更新相关文档

## 联系方式

如有问题或建议，请通过以下方式联系：

- 提交 Issue: [GitHub Issues](https://github.com/joepeak/golib-toolbox/issues)
- 邮箱: [项目邮箱]
