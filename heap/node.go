package heap

type node struct {
	key   interface{}
	score int
}

// Create empty node
func NewNode() *node {
	return &node{}
}

// Initialize new node with values
func InitNode(key interface{}, score int) node {
	return node{key: key, score: score}
}

// Get key-field of node
func (n *node) GetKey() interface{} {
	return n.key
}

// Get score-field of node
func (n *node) GetScore() interface{} {
	return n.score
}

// Set key-field of node
func (n *node) SetKey(key interface{}) {
	n.key = key
}

// Set score-field of node
func (n *node) SetScore(score int) {
	n.score = score
}

// Initialize new node slice from interface{} array and score func
func InitNodes(interSl []interface{}, scoreFunc func(interface{}) int) []node {
	nodeSl := make([]node, len(interSl))
	for i, v := range interSl {
		nodeSl[i] = InitNode(v, scoreFunc(v))
	}
	return nodeSl
}

// Dump out all node.keys in a slice
func DumpKeys(nodeSl []node) []interface{} {
	interSl := make([]interface{}, len(nodeSl))
	for i, v := range nodeSl {
		interSl[i] = v.key
	}
	return interSl
}

// Dump out all node.scores in a slice
func DumpScores(nodeSl []node) []int {
	intSl := make([]int, len(nodeSl))
	for i, v := range nodeSl {
		intSl[i] = v.score
	}
	return intSl
}
