package main

import (
	"flag"
	"log"

	"github.com/news-ai/entitygraph"

	"github.com/jmoiron/sqlx"
	"github.com/jprobinson/go-utils/utils"
	_ "gopkg.in/cq.v1"
	_ "gopkg.in/cq.v1/types"
)

const logPath = "/var/log/entitygraph/graphd.log"

var (
	logArg  = flag.String("log", logPath, "log path")
	reparse = flag.Bool("r", false, "reparse all alerts and events")
)

func main() {
	flag.Parse()

	if *logArg != "stderr" {
		logSetup := utils.NewDefaultLogSetup(*logArg)
		logSetup.SetupLogging()
		go utils.ListenForLogSignal(logSetup)
	} else {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
	}

	config := entitygraph.NewConfig()
	sess, err := config.NeoSession()
	if err != nil {
		log.Fatal(err)
	}

	runGraph(config, sess)
}

func runGraph(config *entitygraph.Config, db *sqlx.DB) {
	query := "The Matrix"
	cypher := `MATCH (movie:Movie) 
				 WHERE movie.title =~ {0} 
				 RETURN movie.title as title, movie.tagline as tagline, movie.released as released`
	movies := []entitygraph.Movie{}
	param := "(?i).*" + query + ".*"
	err := db.Select(&movies, cypher, param)
	if err != nil {
		log.Println("error querying search:", err)
	}

	movieResults := []entitygraph.MovieResult{}
	for _, x := range movies {
		movieResults = append(movieResults, entitygraph.MovieResult{x})
	}

	log.Println(movieResults)
}
