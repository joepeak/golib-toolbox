package main

import (
    "fmt"

    // 元包信息
    toolbox "github.com/joepeak/golib-toolbox"
)

func main() {
    fmt.Println("=== Golib Toolbox 简单示例 ===")
    
    // 显示版本信息
    fmt.Println("版本信息:", toolbox.GetVersion())
    
    fmt.Println("✓ 元包导入成功")
    fmt.Println("\n=== 示例完成 ===")
}
