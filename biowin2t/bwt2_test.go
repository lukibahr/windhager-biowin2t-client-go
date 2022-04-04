package biowin2t

import (
	"context"
	"reflect"
	"testing"
)

func TestNewClient(t *testing.T) {
	url := "http://localhost/api/1.0/lookup"
	username := "usr"
	password := "passwd"
	c := NewWindhagerClient(url,
		username, password)
	if c.MesEndpoint != url {
		t.Errorf("NewWindhagerClient MesEndpoint = %v, want %v", c.MesEndpoint, url)
	}
	if c.MesUsername != username {
		t.Errorf("NewWindhagerClient MesUsername = %v, want %v", c.MesUsername, username)
	}
	if c.MesPassword != password {
		t.Errorf("NewWindhagerClient MesPassword = %v, want %v", c.MesPassword, password)
	}
}

func TestWindhagerClient_GetTimeUntilNextMajorMaintenanceInHours(t *testing.T) {

	url := "http://localhost/api/1.0/lookup"
	username := "usr"
	password := "passwd"
	type fields struct {
		MesEndpoint string
		MesUsername string
		MesPassword string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *successResponse
		wantErr bool
	}{
		{name: "dummy_test", fields: fields{url, username, password}, args: args{ctx: context.Background()}, want: &successResponse{StatusCode: 200}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &WindhagerClient{
				MesEndpoint: tt.fields.MesEndpoint,
				MesUsername: tt.fields.MesUsername,
				MesPassword: tt.fields.MesPassword,
			}
			got, err := c.GetTimeUntilNextMajorMaintenanceInHours(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("WindhagerClient.GetTimeUntilNextMajorMaintenanceInHours() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("WindhagerClient.GetTimeUntilNextMajorMaintenanceInHours() = %v, want %v", got, tt.want)
			}
		})
	}
}
