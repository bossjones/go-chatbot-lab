package robots

import (
	"reflect"
	"testing"
)

func TestNewRobot(t *testing.T) {
	type args struct {
		adapterpath *string
		adapter     *string
		httpd       *bool
		name        *string
		alias       *string
	}
	tests := []struct {
		name string
		args args
		want *Robot
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRobot(tt.args.adapterpath, tt.args.adapter, tt.args.httpd, tt.args.name, tt.args.alias); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRobot() = %v, want %v", got, tt.want)
			}
		})
	}
}
