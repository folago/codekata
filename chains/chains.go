package chains

import (
	"bufio"
	"os"
	"strings"
)

// Distance returns the number of letters that differs from string a to b.
// The strings must be of the same length, if not -1 is returned.
// Important: here we use bytes as characters, so ASCII strings only
func Distance(a, b string) int {
	if len(a) != len(b) {
		return -1
	}
	dist := 0
	for i := range a {
		if a[i] != b[i] {
			dist++
		}
	}
	return dist
}

// BuildGraph bulda a graph as an adjacency list from the list of word that is the dictionary. The graph
// is a map that connect each word with all the words at distance 1 from it
// This is a naive implementation, maybe there is a better one.
func BuildGraph(dictionary []string) map[string][]string {
	ret := make(map[string][]string)
	for _, str := range dictionary {
		ret[str] = []string{}
		for _, ss := range dictionary {
			if Distance(str, ss) == 1 {
				ret[str] = append(ret[str], ss)
			}
		}
	}
	return ret
}

// Path implements a breadth first serach, termination: no more new nodes to visit or we found the stop/end of the path.
// Each iteration we create a new level of the tree, check for the end of the path and add the nodes to the visited set.
// In case of multiple paths with the same legth Path return the first one found.
func Path(graph map[string][]string, start, stop string) []string {
	queue := []*revlist{&revlist{nil, start}} //the beginning
	visited := make(map[string]bool)
	visited[start] = true
	nextLevel := []*revlist{}
	count := 0

	for {
		//new level
		nextLevel = []*revlist{}
		for _, node := range queue {
			if node.val == stop { //we have a winner
				return revpath(node)
			}
			children := explore(graph, node, visited)
			nextLevel = append(nextLevel, children...)
		}
		//check for termination
		if len(nextLevel) == 0 { //no more nodes not visited -> no path
			return nil
		}
		queue = nextLevel
		count++
	}
}

//explore returns all the nodes reachable from on node in a reverse list
func explore(graph map[string][]string, parent *revlist, visited map[string]bool) []*revlist {
	ret := []*revlist{}

	// new level
	list := graph[parent.val]
	for _, node := range list {
		if !visited[node] {
			ret = append(ret, &revlist{parent, node})
			visited[node] = true
		}
	}
	return ret
}

// TODO: figure out if this is a huge memory leak
type revlist struct {
	parent *revlist
	val    string
}

//we need to reconstruct the reverse path and then reverse it
// we have a backward linked list with the first element with a nil parent
func revpath(end *revlist) []string {
	ret := []string{}
	for end.parent != nil {
		ret = append(ret, end.val)
		end = end.parent
	}
	//we misssed the first one since its parent == nil
	ret = append(ret, end.val)

	//thanks slice tricks wiki!
	for left, right := 0, len(ret)-1; left < right; left, right = left+1, right-1 {
		ret[left], ret[right] = ret[right], ret[left]
	}
	return ret
}

// ReadWords reads all the words of a lenght from a file.
func ReadWords(fname string, length int) ([]string, error) {
	fin, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer fin.Close()
	var (
		ret     = []string{}
		scanner = bufio.NewScanner(fin)
		word    string
	)
	for scanner.Scan() {
		word = strings.TrimSpace(scanner.Text())
		if len(word) == length {
			ret = append(ret, word)
		}
	}
	if err = scanner.Err(); err != nil {
		return nil, err
	}
	return ret, nil
}
