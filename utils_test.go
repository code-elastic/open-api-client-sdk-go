package open_api_client_sdk_go

import "testing"

func TestSerializeToJSON(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		result  string
		wantErr bool
	}{
		{name: "t1", args: args{data: map[string]int{"a": 1, "b": 2}}, result: `{"a":1,"b":2}`, wantErr: false},
		{name: "t2", args: args{data: []int{1, 2, 3}}, result: "[1,2,3]", wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SerializeToJSON(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("SerializeToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.result {
				t.Errorf("SerializeToJSON() got = %v, result %v", got, tt.result)
			}
		})
	}
}
