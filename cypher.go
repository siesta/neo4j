package neo4j

import "encoding/json"

type Cypher struct {
	Query   map[string]interface{}
	Payload []*NodeResponse
}

type CypherResponse struct {
	Columns map[string]interface{}   `json:"columns"`
	Data    []map[string]interface{} `json:"data"`
}

func (c *Cypher) mapBatchResponse(neo4j *Neo4j, data interface{}) (bool, error) {
	encodedData, err := jsonEncode(data)
	payload, err := c.decodeResponse(encodedData)
	if err != nil {
		return false, err
	}

	c.Payload = payload

	return true, nil
}

func (c *Cypher) getBatchQuery(operation string) (map[string]interface{}, error) {
	return map[string]interface{}{
		"method": "POST",
		"to":     "/cypher",
		"body":   c.Query,
	}, nil
}

func (c *Cypher) decodeResponse(data string) ([]*NodeResponse, error) {
	resp := map[string]interface{}{}

	err := json.Unmarshal([]byte(data), &resp)
	if err != nil {
		return nil, err
	}

	nodeResponses := make([]*NodeResponse, 0)
	// Iterate through top level "data" array and append each NodeResponse to nodeResponses slice
	for _, v := range resp["data"].([]interface{}) {
		columnData := v.([]interface{})[0]
		jsonColumnData, err := json.Marshal(columnData)
		if err != nil {
			return nil, err
		}

		n := &NodeResponse{}
		err = json.Unmarshal(jsonColumnData, &n)
		if err != nil {
			return nil, err
		}

		nodeResponses = append(nodeResponses, n)
	}

	return nodeResponses, nil
}
