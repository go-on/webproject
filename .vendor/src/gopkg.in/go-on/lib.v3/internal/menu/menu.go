package menu

/*
Menu provides methods to construct flexible menu trees that are independent from the
representation.

A menu tree consists of a root node with edges (children) that have menu items as leafs
and may have further edges.
*/

import (
	"encoding/json"

	"io"
)

// helper for simple json based loading and saving
type Json struct {
	Text string `json:",omitempty"`
	Path string `json:",omitempty"`
	Subs []Json `json:",omitempty"`
}

func (j Json) toNode() *Node {
	n := &Node{Leaf: Item(j.Text, j.Path)}

	for _, sub := range j.Subs {
		n.Edges = append(n.Edges, sub.toNode())
	}
	return n
}

// Node is a node of a menu tree
type Node struct {
	// Edges is a collection of menu nodes that are children the current node
	Edges []*Node

	// Leaf contains the item of the current node.
	// The root node of a menu tree has no leaf (<nil>).
	Leaf Leaf
}

func (n *Node) toJson() Json {
	j := Json{}
	if n.Leaf != nil {
		j.Text = n.Leaf.String()
		j.Path = n.Leaf.Path()
	}
	for _, edge := range n.Edges {
		j.Subs = append(j.Subs, edge.toJson())
	}
	return j
}

func (n Node) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.toJson())
}

// UnmarshalJSON([]byte) error
func (n *Node) UnmarshalJSON(data []byte) (err error) {
	//println("UnmarshalJSON called")
	j := &Json{}
	err = json.Unmarshal(data, &j)
	if err != nil {
		return
	}
	nn := j.toNode()
	n.Leaf = nn.Leaf
	n.Edges = nn.Edges
	return nil
}

// FindByPath returns the first node that has the given path
// by starting with the leaf of m and then going through its egdes.
// If no node is found, <nil> is returned
func (n *Node) FindByPath(path string) *Node {
	if n.Leaf != nil {
		if n.Leaf.Path() == path {
			return n
		}
	}

	for _, edge := range n.Edges {
		found := edge.FindByPath(path)
		if found != nil {
			return found
		}
	}
	return nil
}

// FindByText returns the first node that has the given text
// by starting with the leaf and then going through its egdes.
// If no node is found, <nil> is returned
func (n *Node) FindByText(text string) *Node {
	if n.Leaf != nil {
		if n.Leaf.String() == text {
			return n
		}
	}

	for _, edge := range n.Edges {
		found := edge.FindByText(text)
		if found != nil {
			return found
		}
	}
	return nil
}

// RootAt looks for the node with the given path and returns its
// parent node at nesting level depth.
// It can be used to get a sub menu for the request path within a http handler
// to show it in the layout at a different place than other parts of the menu tree.
// The returned node is a copy of the original node (without the leaf) and may
// not be used to manipulate the original node.
// For the later, FindByText() or FindByPath() can be used.
// If a node was found but it has no edges, <nil> is returned.
func (n *Node) RootAt(depth int, path string) *Node {
	root, found := n.rootAt(depth, path, 0)
	if !found || root == nil {
		return nil
	}
	// a root without edges equals to nil
	if len(root.Edges) == 0 {
		return nil
	}

	// make a copy
	return &Node{Edges: root.Edges, Leaf: root.Leaf}
}

// rootAt is the internal method for the recursive part of RootAt
// The returned node is the root node of the sub menu and pathFound indicates,
// that the given path is found inside the tree of n.
func (n *Node) rootAt(depth int, path string, currentDepth int) (root *Node, pathFound bool) {

	// Check if the n has the given path
	if n.Leaf != nil && path == n.Leaf.Path() {

		// We found the root node
		if currentDepth == depth {
			return n, true
		}

		// We have not found the root node, but the node matching the path
		// - no need to look further, return to the parent
		return nil, true
	}

	// We did not find the node matching the path yet, so look into the edges
	for _, edge := range n.Edges {
		root, pathFound = edge.rootAt(depth, path, currentDepth+1)

		// The path is found in the tree of edge and therefor in the tree of n
		if pathFound {

			// n is at the requested depth, therefor n is the root node we want
			if currentDepth == depth {
				return n, true
			}

			// We have found the path but we are already to deep in the tree, so return to the parent
			// with the info, that the path has been found in our tree, leave it to the parent
			// to find the right root node
			if currentDepth > depth {
				return nil, true
			}

			// We have found the path but we are not deep enough inside the tree.
			// So lets return what the edge has found as root node.
			// It should not be possible for root to be <nil> here.
			return root, true
		}
	}

	// We did not find a node with the given path inside our tree and therefor no root node.
	return
}

// Leaf is the content of a Node
// It can be used for the html representation of the menu
type Leaf interface {

	// Path returns the URL path for the leaf.
	// If it returns an empty string or a string that begins with ~ or $, the leaf
	// is not treated as a link.
	Path() string

	// String returns the string that is used to represent the leaf.
	// If the leaf is treated as a link (see Path()), String() would return the
	// linked text.
	String() string
}

// WriterTo provides a WriterTo that writes a representation of a menu tree
type WriterTo interface {

	// WriterTo returns an io.WriterTo that writes the menu tree for the given root node.
	// The leaf matching the given path may be treated in a special way.
	// The given path may be seen as "active" or "selected" path.
	// The tree should only be visited up to the given depth.
	// WriterTo must handle the case when the root is <nil> (empty menu).
	WriterTo(root *Node, depth int, path string) io.WriterTo
}

type item struct{ text, path string }

func (s *item) String() string { return s.text }
func (s *item) Path() string   { return s.path }

// Item creates a menu item, based on the given text and path.
// The returned item can be used as leaf in a menu tree.
func Item(text, path string) Leaf { return &item{text, path} }

type WriterToFunc func(root *Node, depth int, path string) io.WriterTo

func (wf WriterToFunc) WriterTo(root *Node, depth int, path string) io.WriterTo {
	return wf(root, depth, path)
}
