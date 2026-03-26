package main

import (
	"fmt"

	// 配置管理
	_ "github.com/joepeak/golib-conf"

	// 消息队列
	_ "github.com/joepeak/golib-mq/kafkamq"
	_ "github.com/joepeak/golib-mq/redismq"

	// 工具集
	"github.com/joepeak/golib-toolkit/convert"
	_ "github.com/joepeak/golib-toolkit/database"
	_ "github.com/joepeak/golib-toolkit/oss"

	// 元包信息
	toolbox "github.com/joepeak/golib-toolbox"
)

func main() {
	fmt.Println("=== Golib Toolbox 示例 ===")

	// 显示版本信息
	fmt.Println("版本信息:", toolbox.GetVersion())

	// 1. 配置管理示例
	fmt.Println("\n1. 配置管理:")
	fmt.Println("✓ 配置系统已通过 init() 初始化")

	// 2. 消息队列示例
	fmt.Println("\n2. 消息队列:")
	fmt.Println("✓ Kafka 模块已导入")
	fmt.Println("✓ Redis MQ 模块已导入")

	// 3. 工具集示例
	fmt.Println("\n3. 工具集:")
	fmt.Println("✓ 数据库模块已导入")
	fmt.Println("✓ OSS 模块已导入")
	fmt.Println("✓ 转换工具已导入")

	// 4. 转换工具使用示例
	fmt.Println("\n4. 工具使用示例:")
	str := "Hello, World!"
	reversedStr := convert.Reverse(str)
	fmt.Printf("原字符串: %s\n", str)
	fmt.Printf("反转后: %s\n", reversedStr)

	// 测试手机号脱敏
	mobile := "13812345678"
	maskedMobile := convert.ReplaceByStar(mobile)
	fmt.Printf("手机号脱敏: %s -> %s\n", mobile, maskedMobile)

	fmt.Println("\n=== 示例完成 ===")
}
