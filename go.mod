module github.com/joepeak/golib-toolbox

go 1.22

require (
	github.com/joepeak/golib-conf v0.3.0
	github.com/joepeak/golib-mq v0.3.0
	github.com/joepeak/golib-toolkit v0.3.0
)

// 使用本地依赖，方便开发和测试
replace github.com/joepeak/golib-conf => ../golib-conf

replace github.com/joepeak/golib-mq => ../golib-mq

replace github.com/joepeak/golib-toolkit => ../golib-toolkit
