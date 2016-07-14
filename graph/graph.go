package graph

import (
	"log"

	"github.com/news-ai/entitygraph"

	"github.com/jmoiron/sqlx"
	_ "gopkg.in/cq.v1"
	_ "gopkg.in/cq.v1/types"
)

func QueryGraph(config *entitygraph.Config, db *sqlx.DB, companyName string) []entitygraph.CompanyResult {
	cypher := `MATCH (company:Company) 
                 WHERE company.name =~ {0} 
                 RETURN company.name as name, company.description as description`
	companies := []entitygraph.Company{}
	param := "(?i).*" + companyName + ".*"
	err := db.Select(&companies, cypher, param)
	if err != nil {
		log.Println("error querying search:", err)
	}

	companyResults := []entitygraph.CompanyResult{}
	for _, x := range companies {
		companyResults = append(companyResults, entitygraph.CompanyResult{x})
	}

	return companyResults
}

func AddCompanyToGraph(config *entitygraph.Config, db *sqlx.DB, company *entitygraph.Company) {

}

func AddCompetitorToGraph(config *entitygraph.Config, db *sqlx.DB, company *entitygraph.Company, competitor *entitygraph.Company) {

}

func AddStakeholderToGraph(config *entitygraph.Config, db *sqlx.DB, company *entitygraph.Company, stakeholder *entitygraph.Stakeholders) {

}

func AddKeywordToGraph(config *entitygraph.Config, db *sqlx.DB, company *entitygraph.Company, keyword *entitygraph.Keyword) {

}
