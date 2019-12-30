module moul.io/cmd/pertify

go 1.13

require (
	github.com/pkg/errors v0.8.1
	gopkg.in/urfave/cli.v2 v2.1.1
	gopkg.in/yaml.v3 v3.0.0-20190709130402-674ba3eaed22
	moul.io/graphman v1.5.0
	moul.io/graphman/viz v0.0.0
)

replace moul.io/graphman => ../../

replace moul.io/graphman/viz => ../../viz/
