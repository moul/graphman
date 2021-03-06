package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/pkg/errors"
	cli "gopkg.in/urfave/cli.v2"
	yaml "gopkg.in/yaml.v3"
	"moul.io/graphman"
	"moul.io/graphman/viz"
)

func main() {
	app := &cli.App{
		Name: os.Args[0],
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "file", Aliases: []string{"f"}, Value: "-", Usage: `path to the graph file ("-" for stdin)`},
			&cli.BoolFlag{Name: "dot", Usage: "print 'dot' compatible output"},
			&cli.BoolFlag{Name: "vertical", Usage: "displaying steps from top to bottom"},
			&cli.BoolFlag{Name: "with-details", Usage: "Show pert numbers"},
			&cli.BoolFlag{Name: "no-simplify", Usage: "Don't simplify the graph"},
			&cli.BoolFlag{Name: "standard-pert", Usage: "Use edges for actions instead of vertices"},
			&cli.BoolFlag{Name: "debug", Aliases: []string{"D"}, Usage: "verbose mode"},
		},
		Action: graph,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalf("error: %v", err)
	}
}

func graph(c *cli.Context) error {
	yamlFile, err := ioutil.ReadFile(c.String("file"))
	if err != nil {
		return errors.Wrap(err, "failed to open '--file'")
	}

	var config graphman.PertConfig
	if err = yaml.Unmarshal(yamlFile, &config); err != nil {
		return errors.Wrap(err, "failed to parse yaml file")
	}

	config.Opts.NoSimplify = c.Bool("no-simplify")
	config.Opts.StandardPert = c.Bool("standard-pert")

	graph := graphman.FromPertConfig(config)
	if c.Bool("debug") {
		log.Println(graph)
	}

	// compute and highlight the shortest path
	shortestPath, distance := graph.FindShortestPath("Start", "Finish")
	if c.Bool("debug") {
		log.Println("Shortest path:", shortestPath, "distance:", distance)
	}
	for _, edge := range shortestPath {
		edge.Dst().SetColor("red")
		edge.SetColor("red")
	}
	graph.GetVertex("Start").SetColor("blue")
	graph.GetVertex("Finish").SetColor("blue")

	if c.Bool("vertical") {
		graph.Attrs["rankdir"] = "TB"
	}

	s, err := viz.ToGraphviz(graph, &viz.Opts{
		CommentsInLabel: c.Bool("with-details"),
	})
	if err != nil {
		return err
	}
	fmt.Println(s)
	return nil
}
