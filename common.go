package entitygraph

import (
	"log"

	"github.com/jmoiron/sqlx"
)

const (
	configFile = "/opt/entitygraph/etc/config.json"

	ServerLog = "/var/log/entitygraph/server.log"
	AccessLog = "/var/log/entitygraph/access.log"

	neo4jURL      = "localhost:7474"
	neo4jUser     = "neo4j"
	neo4jPassword = "e(U6V4mHzuRX[9P"
)

type MovieResult struct {
	Movie `json:"movie"`
}

type Movie struct {
	Released int      `json:"released"`
	Title    string   `json:"title,omitempty"`
	Tagline  string   `json:"tagline,omitempty"`
	Cast     []Person `json:"cast,omitempty"`
}

type Person struct {
	Job  string   `json:"job"`
	Role []string `json:"role"`
	Name string   `json:"name"`
}

type D3Response struct {
	Nodes []Node `json:"nodes"`
	Links []Link `json:"links"`
}

type Node struct {
	Title string `json:"title"`
	Label string `json:"label"`
}

type Link struct {
	Source int `json:"source"`
	Target int `json:"target"`
}

type Config struct {
}

func (c *Config) NeoSession() (*sqlx.DB, error) {
	// make conn pass it to data
	constructedURL := "http://" + neo4jUser + ":" + neo4jPassword + "@" + neo4jURL
	db, err := sqlx.Connect("neo4j-cypher", constructedURL)
	if err != nil {
		log.Println("error connecting to neo4j:", err)
	}

	return db, nil
}

func NewConfig() *Config {
	config := Config{}
	return &config
}
