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
		{
			name: "two node graph, path does not exist but node does",
			args: args{
				graph: map[string][]string{"cog": []string{"dog"}, "dog": []string{"cog"}, "cat": []string{}},
				start: "cog",
				stop:  "cat",
			},
			want: nil,
		},
		{
			name: "6 nodes graph, path exists",
			args: args{
				graph: map[string][]string{"A": {"B", "C"}, "B": {"D", "A"}, "C": {"E", "F", "A"}, "E": {"D", "E"}, "F": {"C"}},
				start: "A",
				stop:  "D",
			},
			want: []string{"A", "B", "D"},
		},
		{
			name: "6 nods graph, no path exists",
			args: args{
				graph: map[string][]string{"A": {"B", "C"}, "B": {"D", "A"}, "C": {"E", "F", "A"}, "E": {"D", "E"}, "F": {"C"}, "G": {}},
				start: "A",
				stop:  "G",
			},
			want: nil,
		},
		{
			name: "10 nodes graph, more pats exist",
			args: args{
				graph: map[string][]string{"A": {"B", "C"}, "B": {"D", "A"}, "C": {"E", "F", "A"},
					"D": {"G", "X", "B"}, "E": {"C", "H"}, "F": {"C", "I"}, "G": {"D", "X"},
					"H": {"E", "X"}, "I": {"F", "X"}, "X": {"D", "G", "H", "I"},
				},
				//      A
				//     / \
				//    B   C
				//   /    /\
				//  D    E  F
				//  |\   |  |
				//  | G  H  I
				//  | /  |  |
				//  |/   /  /
				//  X---'--'
				start: "A",
				stop:  "X",
			},
			want: []string{"A", "B", "D", "X"},
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

func TestPathLong(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
	//TODO: add more workds and make this test more significant
	words, err := ReadWords("testdata/word3.txt", 3)
	if err != nil {
		t.Error(err)
	}
	graph3 := BuildGraph(words)

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
			name: "63 node graph, path exists",
			args: args{
				graph: graph3,
				start: "cog",
				stop:  "dog",
			},
			want: []string{"cog", "dog"},
		},
		{
			name: "63 node graph, path exists",
			args: args{
				graph: graph3,
				start: "dog",
				stop:  "ham",
			},
			want: []string{"dog", "dig", "did", "dad", "had", "ham"},
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

func BenchmarkPathDogHam(b *testing.B) {
	words, err := ReadWords("testdata/wordlist.txt", 3)
	if err != nil {
		b.Error(err)
	}
	graph3 := BuildGraph(words)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Path(graph3, "dog", "ham")
	}
}
func BenchmarkPathHamDog(b *testing.B) {
	words, err := ReadWords("testdata/wordlist.txt", 3)
	if err != nil {
		b.Error(err)
	}
	graph3 := BuildGraph(words)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Path(graph3, "ham", "dog")
	}
}
func BenchmarkPathRubyCode(b *testing.B) {
	words, err := ReadWords("testdata/wordlist.txt", 4)
	if err != nil {
		b.Error(err)
	}
	graph3 := BuildGraph(words)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Path(graph3, "ruby", "code")
	}
}

func BenchmarkPathCodeRuby(b *testing.B) {
	words, err := ReadWords("testdata/wordlist.txt", 4)
	if err != nil {
		b.Error(err)
	}
	graph3 := BuildGraph(words)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Path(graph3, "code", "ruby")
	}
}

func TestReadWords(t *testing.T) {
	type args struct {
		fname  string
		length int
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "read words lenght 5 list",
			args: args{
				fname:  "testdata/words_3_5.txt",
				length: 3,
			},
			want: []string{"hae", "hag", "hah", "haj", "ham"},
		},
		{
			name: "read words lenght 5 list",
			args: args{
				fname:  "testdata/words_3_5.txt",
				length: 5,
			},
			want: []string{"aahed", "aalii", "aargh", "abaca"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ReadWords(tt.args.fname, tt.args.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadWords() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
