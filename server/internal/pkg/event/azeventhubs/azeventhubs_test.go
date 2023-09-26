package azeventhubs

import (
	"context"
	"testing"

	"faceto-ai/internal/pkg/event"
)

func TestAZProducer_Send(t *testing.T) {
	t.Skip("run this test manually")
	type args struct {
		ctx context.Context
		msg event.Event
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
				msg: &event.IndexMessage{
					Message: *event.NewMessage("01GZX7K2N5APKQ4VFWEYSZZENZ", []byte("{\"uuid\": \"01GZK6H14KAY24E1R2SZC7AQ2Q\", \"status\": 2}")),
					Index: event.IndexValue{
						UUID:   "01GZK6H14KAY24E1R2SZC7AQ2Q",
						Status: 2,
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := NewAZProducer(newConf().Event, newLog())
			if err := p.Send(tt.args.ctx, tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Send() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAZConsumer_Receive(t *testing.T) {
	t.Skip("run this test manually")

	type args struct {
		ctx     context.Context
		handler event.Handler
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test",
			args: args{
				ctx: context.Background(),
				handler: func(context.Context, event.Event) error {
					return nil
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewAZConsumer(newConf().Event, newLog())
			if err := s.Receive(tt.args.ctx, tt.args.handler); (err != nil) != tt.wantErr {
				t.Errorf("Receive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
