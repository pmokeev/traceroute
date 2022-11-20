package traceroute

import (
	"reflect"
	"testing"
)

func Test_sender_createPacket(t *testing.T) {
	type fields struct {
		config *Config
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "Success",
			fields: fields{
				config: &Config{
					PacketSize: 60,
					Protocol:   "UDP",
				},
			},
			want:    make([]byte, 60),
			wantErr: false,
		},
		{
			name: "Failed",
			fields: fields{
				config: &Config{
					Protocol: "TCP",
				},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sender{
				config: tt.fields.config,
			}
			got, err := s.createPacket()
			if (err != nil) != tt.wantErr {
				t.Errorf("sender.createPacket() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sender.createPacket() = %v, want %v", got, tt.want)
			}
		})
	}
}
