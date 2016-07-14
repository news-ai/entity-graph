package main

import (
	"flag"
	"log"

	"github.com/news-ai/entitygraph"
	"github.com/news-ai/entitygraph/graph"

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

	company := "Y Combinator"
	companyResults := graph.QueryGraph(config, sess, company)
	log.Println(companyResults)
}
