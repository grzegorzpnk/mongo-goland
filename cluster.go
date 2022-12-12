package main

type Cluster struct {
	name      string    `json:"name"`
	resources Resources `json:"resources"`
}

type Resources struct {
	cpu    int `json:"cpu"`
	memory int `json:"memory"`
}
