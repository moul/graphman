package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v3"
	"moul.io/graphman"
	"moul.io/graphman/viz"
)

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	if request.Body == "" {
		return nil, fmt.Errorf("invalid POST request")
	}
	yamlFile := []byte(request.Body)

	var config graphman.PertConfig
	if err := yaml.Unmarshal(yamlFile, &config); err != nil {
		return nil, errors.Wrap(err, "failed to parse yaml file")
	}

	if request.QueryStringParameters["no-simplify"] == "true" {
		config.Opts.NoSimplify = true
	}
	graph := graphman.FromPertConfig(config)
	if request.QueryStringParameters["debug"] == "true" {
		log.Println(graph)
	}

	// compute and highlight the shortest path
	shortestPath, distance := graph.FindShortestPath("Start", "Finish")
	if request.QueryStringParameters["debug"] == "true" {
		log.Println("Shortest path:", shortestPath, "distance:", distance)
	}
	for _, edge := range shortestPath {
		edge.Dst().SetColor("red")
		edge.SetColor("red")
	}
	graph.GetVertex("Start").SetColor("blue")
	graph.GetVertex("Finish").SetColor("blue")

	if request.QueryStringParameters["vertical"] == "true" {
		graph.Attrs["rankdir"] = "TB"
	}

	s, err := viz.ToGraphviz(graph, &viz.Opts{
		CommentsInLabel: request.QueryStringParameters["with-details"] == "true",
	})
	if err != nil {
		return nil, err
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       s,
		Headers: map[string]string{
			"Content-Type": "text/plain; charset=utf-8", // FIXME: graphviz content-type
		},
	}, nil
}

func main() {
	lambda.Start(handler)
}
