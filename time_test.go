package null

import (
	"reflect"
	"testing"
	"time"
)

func TestTime_MarshalJSON(t *testing.T) {
	type fields struct {
		Time  time.Time
		Valid bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "valid date",
			fields: fields{
				Time:  time.Date(1993, 07, 30, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
			want: []byte(`"1993-07-30T00:00:00Z"`),
		}, {
			name: "invalid",
			fields: fields{
				Time:  time.Time{},
				Valid: false,
			},
			want: []byte(`null`),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := Time{
				Time:  tt.fields.Time,
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

func TestTime_UnmarshalJSON(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name    string
		args    args
		want    Time
		wantErr bool
	}{
		{
			name: "valid date",
			args: args{data: []byte(`"1993-07-30T00:00:00Z"`)},
			want: Time{
				Time:  time.Date(1993, 07, 30, 0, 0, 0, 0, time.UTC),
				Valid: true,
			},
		}, {
			name: "valid null",
			args: args{data: []byte(`null`)},
			want: Time{
				Valid: false,
			},
		}, {
			name:    "invalid string",
			args:    args{data: []byte(`"value"`)},
			wantErr: true,
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
			var got Time
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
