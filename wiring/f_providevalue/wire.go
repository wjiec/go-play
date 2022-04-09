//go:build wireinject

package main

import (
	"bytes"
	"io"

	"github.com/google/wire"
)

func NewAlerter() Alerter {
	panic(wire.Build(
		NewSimpleAlerter,
		wire.Bind(new(Alerter), new(*SimpleAlerter)),
		// 这里我们直接提供一个命名类型的值当作依赖，这个值将被注入到
		// 其他需要这个类型的值的服务中
		wire.Value(Message("alert message")),
		// 对于接口类型的依赖，我们需要使用 wire.InterfaceValue 进行绑定和注入
		wire.InterfaceValue(new(io.Writer), &bytes.Buffer{}),
	))
}
