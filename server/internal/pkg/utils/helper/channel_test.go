package helper

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"testing"
)

func TestRecoverPanic(t *testing.T) {
	ctx := context.Background()

	type args struct {
		ctx    context.Context
		log    *log.Helper
		format string
		a      []interface{}
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "recover panic case1",
			args: args{
				ctx:    ctx,
				log:    logger(ctx),
				format: "this is a test, %s",
				a:      []interface{}{"hello world"},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			RecoverPanic(tt.args.ctx, tt.args.log, tt.args.format, tt.args.a...)
		})
	}
}
