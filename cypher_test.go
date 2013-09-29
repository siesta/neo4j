package neo4j

import (
	"log"

	"testing"
)

func TestSendCypherCreateWithLabel(t *testing.T) {
	neo4jConnection := Connect("")

	cypher := &Cypher{
		Query: map[string]interface{}{
			"query": `
			CREATE (k:Labeled {id: {id}})
			return k`,
			"params": map[string]interface{}{
				"id": "testid123",
			},
		},
	}

	batch := neo4jConnection.NewBatch()
	batch.Create(cypher)
	batch.Execute()

	if cypher.Payload == nil {
		t.Error("No cypher results")
	}

	log.Println(cypher.Payload[0])
}

func TestSendCypherQuery(t *testing.T) {
	neo4jConnection := Connect("")

	cypher := &Cypher{
		Query: map[string]interface{}{
			"query": `
			MATCH (k:Labeled)
			WHERE k.id = {id}
			return k`,
			"params": map[string]interface{}{
				"id": "testid123",
			},
		},
	}

	batch := neo4jConnection.NewBatch()
	batch.Create(cypher)
	batch.Execute()

	if cypher.Payload == nil {
		t.Error("No cypher results")
	}

	log.Println(cypher.Payload[0])
}
