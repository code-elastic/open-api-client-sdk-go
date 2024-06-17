package open_api_client_sdk_go

import "testing"

func TestOpenClient_GetNameByGet(t *testing.T) {
	type fields struct {
		appKey    string
		appSecret string
		Nonce     int
		Timestamp string
		Sign      string
		Parameter string
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
		{name: "time1", fields: fields{}, args: args{"hhh"}},
		{name: "time2", fields: fields{}, args: args{"zzz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oc := OpenClient{
				appKey:    tt.fields.appKey,
				appSecret: tt.fields.appSecret,
				Nonce:     tt.fields.Nonce,
				Timestamp: tt.fields.Timestamp,
				Sign:      tt.fields.Sign,
				Parameter: tt.fields.Parameter,
			}
			oc.GetNameByGet(tt.args.name)
		})
	}
}

func TestOpenClient_GetNameByPost(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "time1", args: args{"hhh"}},
		{name: "time2", args: args{"zzz"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oc := NewOpenClient("qwert123456", "asdfgqwertzxcvb")
			oc.GetNameByPost(tt.args.name)
		})
	}
}
