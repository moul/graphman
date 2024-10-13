module moul.io/cmd/pertify

go 1.13

require (
	github.com/pkg/errors v0.9.1
	gopkg.in/urfave/cli.v2 v2.27.5
	gopkg.in/yaml.v3 v3.0.1
	moul.io/graphman v1.6.0
	moul.io/graphman/viz v0.0.0
)

replace moul.io/graphman => ../../

replace moul.io/graphman/viz => ../../viz/
