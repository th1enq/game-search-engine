package elastic

import (
	"log"

	"github.com/elastic/go-elasticsearch"
	"github.com/one-piece-search-engine/config"
)

type ElasticSearch struct {
	*elasticsearch.Client
}

func Load(config *config.Config) *ElasticSearch {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("failed to loading elastic search: %v", err)
	}
	return &ElasticSearch{client}
}
