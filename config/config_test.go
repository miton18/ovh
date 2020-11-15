package config

import (
	"os"
	"reflect"
	"testing"
)

func TestFromEnv(t *testing.T) {
	tests := []struct {
		name    string
		pre func()
		want    *Configuration
		wantErr bool
	}{{
		name: "main",
		pre: func() {
			os.Setenv("OVH_ENDPOINT", "ovh-eu")
			os.Setenv("OVH_APPLICATION_KEY", "key")
			os.Setenv("OVH_APPLICATION_SECRET", "secret")
			os.Setenv("OVH_CONSUMER_KEY", "ck")
		},
		want: &Configuration{
			Endpoint: OVH_EU,
			ApplicationKey: "key",
			ApplicationSecret: "secret",
			ConsumerKey: "ck",
		},
		wantErr: false,
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.pre()

			got, err := FromEnv()
			if (err != nil) != tt.wantErr {
				t.Errorf("FromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromEnv() got = %v, want %v", got, tt.want)
			}
		})
	}
}
