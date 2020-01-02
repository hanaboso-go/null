package null

import (
	"reflect"
	"testing"
)

func TestString_MarshalJSON(t *testing.T) {
	type fields struct {
		String string
		Valid  bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid",
			fields: fields{
				String: "value",
				Valid:  true,
			},
			want: []byte(`"value"`),
		}, {
			name: "invalid",
			fields: fields{
				String: "value",
				Valid:  false,
			},
			want: []byte(`null`),
		}, {
			name: "empty valid",
			fields: fields{
				String: "",
				Valid:  true,
			},
			want: []byte(`""`),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := String{
				String: tt.fields.String,
				Valid:  tt.fields.Valid,
			}
			got, err := ns.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MarshalJSON() got = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestString_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    String
		wantErr bool
	}{
		{
			name: "valid value",
			args: args{data: []byte(`"value"`)},
			want: String{
				String: "value",
				Valid:  true,
			},
		}, {
			name: "valid null",
			args: args{data: []byte(`null`)},
			want: String{
				Valid: false,
			},
		}, {
			name:    "invalid number",
			args:    args{data: []byte(`42`)},
			wantErr: true,
		}, {
			name:    "invalid object",
			args:    args{data: []byte(`{"key":"value"}`)},
			wantErr: true,
		}, {
			name:    "invalid JSON",
			args:    args{data: []byte(`423345}dfsf`)},
			wantErr: true,
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got String
			if err := got.UnmarshalJSON(tt.args.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnmarshalJSON() got = %v, want %v", got, tt.want)
			}
		})
	}
}
