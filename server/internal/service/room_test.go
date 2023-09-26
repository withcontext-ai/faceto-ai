package service

import (
	"testing"
)

func Test_getJoinToken(t *testing.T) {
	type args struct {
		apiKey    string
		apiSecret string
		room      string
		identity  string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "case",
			args: args{
				//   api_key: "${LIVEKIT_API_KEY:API4WvidXeg2FP8}"
				//  secret_key: "${LIVEKIT_SECRET_KEY:bWixV5U4Ink9ZVgdsbhrxRaQqcbRk2PM6grXH79voBA}"
				apiKey:    "API4WvidXeg2FP8",
				apiSecret: "bWixV5U4Ink9ZVgdsbhrxRaQqcbRk2PM6grXH79voBA",
				room:      "5p65-pqvp",
				identity:  "111",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getJoinToken(tt.args.apiKey, tt.args.apiSecret, tt.args.room, tt.args.identity)
			if err != nil {
				t.Errorf("getJoinToken() error = %v", err)
				return
			}
			// {"level":"DEBUG","ts":"2023-05-30T11:32:21.596+0800","caller":"service/room_test.go:41","msg":"at.ToJWT:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODU0MjExNDEsImlzcyI6IkFQSWZ1dG0zU2F6V0FlNiIsIm5iZiI6MTY4NTQxNzU0MSwic3ViIjoiMTExIiwidmlkZW8iOnsicm9vbSI6IjVwNjUtcHF2cCIsInJvb21Kb2luIjp0cnVlfX0.Cinv5iQoWYEd37tJCybW4jij9tUNMXfgHI6pSRuWXPo"}
			// {"level":"DEBUG","ts":"2023-05-30T11:43:11.434+0800","caller":"service/room_test.go:41","msg":"at.ToJWT:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODU0MjE3OTEsImlzcyI6IkFQSWZ1dG0zU2F6V0FlNiIsIm5iZiI6MTY4NTQxODE5MSwic3ViIjoiMTExIiwidmlkZW8iOnsicm9vbSI6IjVwNjUtcHF2cCIsInJvb21Kb2luIjp0cnVlfX0.wfFuA4at8kDwJL2o_SgKYnRBoxvAAgMP2Yi4tWaTEMA"}
			// {"level":"DEBUG","ts":"2023-05-30T11:45:42.064+0800","caller":"service/room_test.go:41","msg":"at.ToJWT:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODU0MjE5NDIsImlzcyI6IkFQSWZ1dG0zU2F6V0FlNiIsIm5iZiI6MTY4NTQxODM0Miwic3ViIjoiMTExIiwidmlkZW8iOnsicm9vbSI6IjVwNjUtcHF2cCIsInJvb21Kb2luIjp0cnVlLCJyb29tTGlzdCI6dHJ1ZX19.fSMB3lgWF8xYLlrZrsSEHgctwe0bWTJU3U3IfndLz1w"}
			// {"level":"DEBUG","ts":"2023-05-30T11:50:25.020+0800","caller":"service/room_test.go:41","msg":"at.ToJWT:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODU0MjIyMjUsImlzcyI6IkFQSWZ1dG0zU2F6V0FlNiIsIm5iZiI6MTY4NTQxODYyNSwic3ViIjoiMTExIiwidmlkZW8iOnsicm9vbSI6IjVwNjUtcHF2cCIsInJvb21BZG1pbiI6dHJ1ZSwicm9vbUNyZWF0ZSI6dHJ1ZSwicm9vbUpvaW4iOnRydWUsInJvb21MaXN0Ijp0cnVlLCJyb29tUmVjb3JkIjp0cnVlfX0.SQoxxqGZOcWhwGT9HyjB0GUxfShTf9U0PUm_sv3U_x0"}
			// {"level":"DEBUG","ts":"2023-05-30T11:55:35.386+0800","caller":"service/room_test.go:43","msg":"at.ToJWT:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2ODU0MjI1MzUsImlzcyI6IkFQSTRXdmlkWGVnMkZQOCIsIm5iZiI6MTY4NTQxODkzNSwic3ViIjoiMTExIiwidmlkZW8iOnsicm9vbSI6IjVwNjUtcHF2cCIsInJvb21BZG1pbiI6dHJ1ZSwicm9vbUNyZWF0ZSI6dHJ1ZSwicm9vbUpvaW4iOnRydWUsInJvb21MaXN0Ijp0cnVlLCJyb29tUmVjb3JkIjp0cnVlfX0.XnYlcXdbXQgTyKoMX1iRXWi9g93wnO4jdOZCNbquXcI"}
			t.Logf("got->%s", got)
		})
	}
}
