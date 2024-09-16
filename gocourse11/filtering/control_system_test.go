package filtering

import (
	"reflect"
	"testing"
)

func TestCreateFilter(t *testing.T) {
	type args struct {
		cleanerLevel CleanerLevel
	}
	tests := []struct {
		name string
		args args
		want Filter
	}{
		{
			name: "Low pollution",
			args: args{cleanerLevel: Low},
			want: Filter{cleanLevel: Low, absorber: "sand", waterImprover: "t2w"},
		},
		{
			name: "Middle pollution",
			args: args{cleanerLevel: Middle},
			want: Filter{cleanLevel: Middle, absorber: "coal", waterImprover: "cn2"},
		},
		{
			name: "High pollution",
			args: args{cleanerLevel: High},
			want: Filter{cleanLevel: High, absorber: "vibranium", waterImprover: "yy78"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateFilter(tt.args.cleanerLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateFilter(): got = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestSelectFilter(t *testing.T) {
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
			if got := SelectFilter(tt.args.pollutionLevel); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectFilter(): got = %v, want = %v", got, tt.want)
			}
		})
	}
}
