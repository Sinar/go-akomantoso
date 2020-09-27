package main

import "testing"

func TestSayItCmd_Run(t *testing.T) {
	type fields struct {
		ID            int
		DebateRawFile string
		DebateType    string
		OutputFile    string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &SayItCmd{
				ID:            tt.fields.ID,
				DebateRawFile: tt.fields.DebateRawFile,
				DebateType:    tt.fields.DebateType,
				OutputFile:    tt.fields.OutputFile,
			}
			if err := m.Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
