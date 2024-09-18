package filter

import (
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	type args struct {
		cleanLevel CleanLevel
	}
	tests := []struct {
		name string
		args args
		want Filter
	}{
		{
			name: "Low pollution",
			args: args{cleanLevel: Low},
			want: Filter{cleanLevel: Low, absorber: "sand", waterImprover: "t2w"},
		},
		{
			name: "Middle pollution",
			args: args{cleanLevel: Middle},
			want: Filter{cleanLevel: Middle, absorber: "coal", waterImprover: "cn2"},
		},
		{
			name: "High pollution",
			args: args{cleanLevel: High},
			want: Filter{cleanLevel: High, absorber: "vibranium", waterImprover: "yy78"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Create(tt.args.cleanLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create(): got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestSelect(t *testing.T) {
	type args struct {
		pollutionLevel int
	}
	tests := []struct {
		name string
		args args
		want Filter
	}{
		{
			name: "Low pollution",
			args: args{pollutionLevel: 101},
			want: Filter{cleanLevel: Low, absorber: "sand", waterImprover: "t2w"},
		},
		{
			name: "Middle pollution",
			args: args{pollutionLevel: 600},
			want: Filter{cleanLevel: Middle, absorber: "coal", waterImprover: "cn2"},
		},
		{
			name: "High pollution",
			args: args{pollutionLevel: 1200},
			want: Filter{cleanLevel: High, absorber: "vibranium", waterImprover: "yy78"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Select(tt.args.pollutionLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Select(): got = %v, want = %v", got, tt.want)
			}
		})
	}
}
