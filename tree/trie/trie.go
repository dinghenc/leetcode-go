package trie

type Node struct {
	children [26]*Node
	end      bool
}

func New() *Node {
	return new(Node)
}

func (n *Node) BuildDictionary(words []string) {
	for _, w := range words {
		n.Add(w)
	}
}

func (n *Node) Add(word string) {
	node := n
	for i := 0; i < len(word); i++ {
		if node.children[word[i]-'a'] == nil {
			node.children[word[i]-'a'] = New()
		}
		node = node.children[word[i]-'a']
	}
	node.end = true
}

func (n *Node) Search(word string) bool {
	node := n.searchEndNode(word)
	return node != nil && node.end
}

func (n *Node) StartWith(prefix string) bool {
	return n.searchEndNode(prefix) != nil
}

func (n *Node) searchEndNode(word string) *Node {
	node := n
	for i := 0; i < len(word); i++ {
		if node.children[word[i]-'a'] == nil {
			return nil
		}
		node = node.children[word[i]-'a']
	}
	return node
}
