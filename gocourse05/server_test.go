package main

import (
	"testing"
	"time"

	"gocourse05/camera"
)

func TestSaveProcessedData(t *testing.T) {
	var testMemory Memory
	type args struct {
		p camera.Processor
		m *Memory
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
				m: &testMemory,
			},
			wantErr: false,
		},
		{
			name: "Fail saveProcessedData",
			args: args{
				p: camera.NewDayCamera(1, "test", []camera.Data{}),
				m: &testMemory,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := saveProcessedData(tt.args.p, tt.args.m); (err != nil) != tt.wantErr {
				t.Errorf("saveProcessedData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSendProcessedData(t *testing.T) {
	type args struct {
		m   Memory
		url string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success sendProcessedData",
			args: args{
				m:   []camera.ProcessedData{camera.NewProcessedData(time.Now(), "Медвідь, побіг вліво")},
				url: "test",
			},
			wantErr: false,
		},
		{
			name: "Fail sendProcessedData",
			args: args{
				m:   []camera.ProcessedData{camera.NewProcessedData(time.Now(), "")},
				url: "test",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := sendProcessedData(tt.args.m, tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("sendProcessedData() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
