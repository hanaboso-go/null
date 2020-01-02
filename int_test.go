package null

import (
	"reflect"
	"testing"
)

func TestInt64_MarshalJSON(t *testing.T) {
	type fields struct {
		Int64 int64
		Valid bool
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
				Int64: 42.0,
				Valid: true,
			},
			want: []byte(`42`),
		}, {
			name: "invalid",
			fields: fields{
				Int64: 42,
				Valid: false,
			},
			want: []byte(`null`),
		}, {
			name: "empty valid",
			fields: fields{
				Int64: 0,
				Valid: true,
			},
			want: []byte(`0`),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := Int64{
				Int64: tt.fields.Int64,
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

func TestInt64_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Int64
		wantErr bool
	}{
		{
			name: "valid value",
			args: args{data: []byte(`42`)},
			want: Int64{
				Int64: 42,
				Valid: true,
			},
		}, {
			name: "valid null",
			args: args{data: []byte(`null`)},
			want: Int64{
				Valid: false,
			},
		}, {
			name:    "invalid string",
			args:    args{data: []byte(`"42"`)},
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
			var got Int64
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

func TestInt32_MarshalJSON(t *testing.T) {
	type fields struct {
		Int32 int32
		Valid bool
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
				Int32: 42,
				Valid: true,
			},
			want: []byte(`42`),
		}, {
			name: "invalid",
			fields: fields{
				Int32: 42,
				Valid: false,
			},
			want: []byte(`null`),
		}, {
			name: "empty valid",
			fields: fields{
				Int32: 0,
				Valid: true,
			},
			want: []byte(`0`),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := Int32{
				Int32: tt.fields.Int32,
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

func TestInt32_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Int32
		wantErr bool
	}{
		{
			name: "valid value",
			args: args{data: []byte(`42`)},
			want: Int32{
				Int32: 42,
				Valid: true,
			},
		}, {
			name: "valid null",
			args: args{data: []byte(`null`)},
			want: Int32{
				Valid: false,
			},
		}, {
			name:    "invalid string",
			args:    args{data: []byte(`"42"`)},
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
			var got Int32
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
