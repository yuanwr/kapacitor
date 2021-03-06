package pipeline

// An HTTPOutNode caches the most recent data for each group it has received.
//
// The cached data is available at the given endpoint.
// The endpoint is the relative path from the API endpoint of the running task.
// For example if the task endpoint is at "/task/<task_name>" and endpoint is
// "top10", then the data can be requested from "/task/<task_name>/top10".
//
// Example:
//    stream
//        .window()
//            .period(10s)
//            .every(5s)
//        .mapReduce(influxql.top('value', 10))
//        //Publish the top 10 results over the last 10s updated every 5s.
//        .httpOut('top10')
//
type HTTPOutNode struct {
	node

	// The relative path where the cached data is exposed
	// tick:ignore
	Endpoint string
}

func newHTTPOutNode(wants EdgeType, endpoint string) *HTTPOutNode {
	return &HTTPOutNode{
		node: node{
			desc:     "http_out",
			wants:    wants,
			provides: NoEdge,
		},
		Endpoint: endpoint,
	}
}
