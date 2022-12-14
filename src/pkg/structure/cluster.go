package structure

type Cluster struct {
	Name      string    `json:"name,omitempty" bson:"name,omitempty"`
	Resources Resources `json:"resources" bson:"resources,omitempty"`
}

type Resources struct {
	Cpu    int `json:"cpu" bson:"cpu"`
	Memory int `json:"memory" bson:"memory"`
}
