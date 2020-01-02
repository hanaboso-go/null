package null

import (
	"reflect"
	"testing"
)

func TestBool_MarshalJSON(t *testing.T) {
	type fields struct {
		Bool  bool
		Valid bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid true",
			fields: fields{
				Bool:  true,
				Valid: true,
			},
			want: []byte(`true`),
		},
		{
			name: "valid false",
			fields: fields{
				Bool:  true,
				Valid: true,
			},
			want: []byte(`true`),
		}, {
			name: "invalid",
			fields: fields{
				Bool:  true,
				Valid: false,
			},
			want: []byte(`null`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := Bool{
				Bool:  tt.fields.Bool,
				Valid: tt.fields.Valid,
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

func TestBool_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Bool
		wantErr bool
	}{
		{
			name: "valid true",
			args: args{data: []byte(`true`)},
			want: Bool{
				Bool:  true,
				Valid: true,
			},
		},
		{
			name: "valid false",
			args: args{data: []byte(`false`)},
			want: Bool{
				Bool:  false,
				Valid: true,
			},
		}, {
			name: "valid null",
			args: args{data: []byte(`null`)},
			want: Bool{
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
			var got Bool
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
