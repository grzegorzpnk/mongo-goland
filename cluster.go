package main

type Cluster struct {
	Name      string    `json:"name"`
	Resources Resources `json:"resources"`
}

type Resources struct {
	Cpu    int `json:"cpu"`
	Memory int `json:"memory"`
}
