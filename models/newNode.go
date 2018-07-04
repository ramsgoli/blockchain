package models

type newNodes struct {
	Nodes []string `json:"nodes"`
}

type newNodesResponse struct {
	message string
	len     int
}
