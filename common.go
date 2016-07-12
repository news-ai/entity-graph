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

type CompanyResult struct {
	Company `json:"company"`
}

type Company struct {
	Name         string    `json:"name"`
	Competitors  []Company `json:"competitors,omitempty"`
	Stakeholders []Person  `json:"stakeholders,omitempty"`
	Keywords     []Keyword `json:"keywords,omitempty"`
}

type Person struct {
	Job  string   `json:"job"`
	Role []string `json:"role"`
	Name string   `json:"name"`
}

type Keyword struct {
	Name string `json:"name"`
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
