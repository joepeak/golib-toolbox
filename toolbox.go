package toolbox

import "fmt"

// Golib Toolbox 元包
//
// 这个包作为统一入口，整合了 golib-conf、golib-mq 和 golib-toolkit 三个组件库
//
// 使用方式：
//   import "faithtech.top/golib/golib-toolbox"
//
//   // 然后可以使用子包：
//   // import "faithtech.top/golib/golib-conf"
//   // import "faithtech.top/golib/golib-mq/kafkamq"
//   // import "faithtech.top/golib/golib-toolkit/database"
//
// 注意：这个包本身不提供具体功能，只是作为依赖的统一入口

// Version 版本信息
const Version = "1.0.0"

// 组件版本信息
const (
	ConfVersion    = "v0.3.1" // golib-conf 版本
	MQVersion      = "v0.3.1" // golib-mq 版本
	ToolkitVersion = "v0.3.0" // golib-toolkit 版本
)

// GetVersion 获取完整版本信息
func GetVersion() string {
	return fmt.Sprintf("golib-toolbox %s (conf:%s, mq:%s, toolkit:%s)",
		Version, ConfVersion, MQVersion, ToolkitVersion)
}
