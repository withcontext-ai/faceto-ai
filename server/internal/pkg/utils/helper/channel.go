package helper

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"runtime"
)

func RecoverPanic(ctx context.Context, log *log.Helper, format string, a ...interface{}) {
	if err := recover(); err != nil {
		stack := make([]byte, 4<<10)
		length := runtime.Stack(stack, true)
		format = "[PANIC RECOVER] stack:%s, goroutine run error, " + format
		a = append([]interface{}{string(stack[:length])}, a...)
		log.WithContext(ctx).Errorf(format, a...)
	}
}
