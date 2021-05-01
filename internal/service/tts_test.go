package service

import (
	"testing"
)

func TestImportDeck(t *testing.T) {
	type args struct {
		id     string
		lang   string
		sleeve string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{"CotA", args{"ba1a3559-f377-40cb-9e8d-e61416b9c6b3", "pt", "red"}, 1, false},
		{"DT", args{"2e1219ed-4740-4b91-b7b0-26d3ff44499a", "pt", "black"}, 2, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ImportDeck(tt.args.id, tt.args.lang, tt.args.sleeve)
			if (err != nil) != tt.wantErr {
				t.Errorf("ImportDeck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil || len(got.ObjectStates) != tt.want {
				t.Errorf("ImportDeck() got = %v, want %v", got, tt.want)
			}
		})
	}
}
