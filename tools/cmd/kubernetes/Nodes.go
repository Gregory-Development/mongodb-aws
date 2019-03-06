package kubernetes

type Node struct {
	Kind string `json:"kind"`
	ApiVersion string `json:"apiVersion"`
	Metadata nodeMetadata `json:"metadata"`
}

type nodeMetadata struct {
	Name string `json:"name"`
	Labels label `json:"labels"`
}

type label struct {
	Name string `json:"name"`
}

func (n *Node) NewNode() {

}

func (n *Node) ModifyNode() {

}

func (n *Node) DeleteNode() {

}
