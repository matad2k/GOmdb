package goomdb

import (
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	type args struct {
		api string
	}
	tests := []struct {
		name    string
		args    args
		want    *client
		wantErr bool
	}{
		{
			name:    "",
			args:    args{},
			want:    &client{},
			wantErr: true,
		},
		{
			name:    "EWE",
			args:    args{"EWE"},
			want:    &client{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.args.api)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewClient() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewClient() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_generateQueryString(t *testing.T) {
	type fields struct {
		apikey string
	}
	type args struct {
		query string
		mode  uint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "movieByTitle",
			fields: fields{"testApi"},
			args:   args{"TestMovie", movieByTitle},
			want:   "http://www.omdbapi.com/?apikey=testApi&t=TestMovie",
		},
		{
			name:   "moviebyID",
			fields: fields{"testApi"},
			args:   args{"TestMovie", movieByID},
			want:   "http://www.omdbapi.com/?apikey=testApi&i=TestMovie",
		},
		{
			name:   "moviebyID",
			fields: fields{"testApi"},
			args:   args{"TestMovie", serachMovie},
			want:   "http://www.omdbapi.com/?apikey=testApi&s=TestMovie",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &client{
				apikey: tt.fields.apikey,
			}
			if got := c.generateQueryString(tt.args.query, tt.args.mode); got != tt.want {
				t.Errorf("generateQueryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
