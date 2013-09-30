package neo4j

import (
	"log"
	"strconv"
	"testing"
	"time"
)

func TestSendCypherCreateWithLabel(t *testing.T) {
	neo4jConnection := Connect("")

	props := make([]map[string]interface{}, 0)
	p1 := map[string]interface{}{"testid": 1, "created": strconv.FormatInt(time.Now().UTC().UnixNano(), 10)}
	p2 := map[string]interface{}{"testid": 2, "created": strconv.FormatInt(time.Now().UTC().UnixNano(), 10)}
	p3 := map[string]interface{}{"testid": 3, "created": strconv.FormatInt(time.Now().UTC().UnixNano(), 10)}
	props = append(props, p1, p2, p3)

	cypher := &Cypher{
		Query: map[string]interface{}{
			"query": `
			CREATE (k:TestLabel {props})
			RETURN k`,
			"params": map[string]interface{}{
				"props": props,
			},
		},
	}

	batch := neo4jConnection.NewBatch()
	batch.Create(cypher)
	batch.Execute()

	if cypher.Payload == nil {
		t.Error("No cypher results")
	}
	if len(cypher.Payload) != 3 {
		t.Error("Wrong number of results")
	}

	for i, v := range cypher.Payload {
		id := int(v.Data["testid"].(float64))
		if id != (i + 1) {
			t.Error("Wrong results returned")
			log.Println(v.Data)
		}
	}
}

func TestSendCypherQuery(t *testing.T) {
	neo4jConnection := Connect("")

	cypher := &Cypher{
		Query: map[string]interface{}{
			"query": `
			MATCH (k:TestLabel)
			RETURN k
			ORDER BY k.created DESC
			LIMIT 3`,
		},
	}

	batch := neo4jConnection.NewBatch()
	batch.Create(cypher)
	batch.Execute()

	if cypher.Payload == nil {
		t.Error("No cypher results")
	}
	if len(cypher.Payload) != 3 {
		t.Error("Wrong number of results")
	}

	for i, v := range cypher.Payload {
		id := int(v.Data["testid"].(float64))
		// ids should be in DESC order 3, 2, 1
		if id != (3 - i) {
			t.Error("Wrong results returned")
			log.Println(v.Data)
		}
	}
}
