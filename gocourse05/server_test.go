package main

import (
	"gocourse05/camera"
	"testing"
	"time"
)

func TestSaveProcessedData(t *testing.T) {
	server := Server{Memory: make([]camera.ProcessedData, 0)}
	type args struct {
		p Processor
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success saveProcessedData",
			args: args{
				p: camera.NewDayCamera(1, "test", []camera.Data{{1, "Коза", "Побігла вліво"}}),
			},
			wantErr: false,
		},
		{
			name: "Fail saveProcessedData",
			args: args{
				p: camera.NewDayCamera(1, "test", []camera.Data{}),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server.setProcessor(tt.args.p)
			if err := server.saveProcessedData(); (err != nil) != tt.wantErr {
				t.Errorf("saveProcessedData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendProcessedData(t *testing.T) {
	type args struct {
		url    string
		server Server
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success sendProcessedData",
			args: args{
				url:    "example.com",
				server: Server{Memory: []camera.ProcessedData{{time.Now(), "Медвідь пішов"}}},
			},
			wantErr: false,
		},
		{
			name: "Fail sendProcessedData",
			args: args{
				url:    "example.com",
				server: Server{Memory: []camera.ProcessedData{}},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.args.server.sendProcessedData(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("sendProcessedData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
