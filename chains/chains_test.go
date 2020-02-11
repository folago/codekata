package chains

import (
	"reflect"
	"testing"
)

func TestDistance(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "distance zero",
			args: args{
				a: "dog",
				b: "dog",
			},
			want: 0,
		},
		{
			name: "distance one",
			args: args{
				a: "dog",
				b: "cog",
			},
			want: 1,
		},
		{
			name: "different lenght",
			args: args{
				a: "dog",
				b: "dogg",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Distance(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Distance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBuildGraph(t *testing.T) {
	type args struct {
		dictionary []string
	}
	tests := []struct {
		name string
		args args
		want map[string][]string
	}{
		{
			name: "simple graph: dog, cog, cat",
			args: args{
				dictionary: []string{"dog", "cog", "cat"},
			},
			want: map[string][]string{"cog": []string{"dog"}, "dog": []string{"cog"}, "cat": []string{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BuildGraph(tt.args.dictionary); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildGraph() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPath(t *testing.T) {
	type args struct {
		graph map[string][]string
		start string
		stop  string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "three node graph, path exists",
			args: args{
				graph: map[string][]string{"cog": []string{"dog"}, "dog": []string{"cog"}, "cat": []string{}},
				start: "cog",
				stop:  "dog",
			},
			want: []string{"cog", "dog"},
		},
		{
			name: "two node graph, path does not exist",
			args: args{
				graph: map[string][]string{"cog": []string{"dog"}, "dog": []string{"cog"}},
				start: "cog",
				stop:  "cat",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Path(tt.args.graph, tt.args.start, tt.args.stop); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Path() = %v, want %v", got, tt.want)
			}
		})
	}
}
