module moul.io/cmd/pertify

go 1.12

require moul.io/graphman v0.0.0

require (
	github.com/pkg/errors v0.8.1
	gopkg.in/urfave/cli.v2 v2.0.0-20180128182452-d3ae77c26ac8
	gopkg.in/yaml.v3 v3.0.0-20190705120443-117fdf03f45f
	moul.io/graphman/viz v0.0.0
)

replace moul.io/graphman => ../../

replace moul.io/graphman/viz => ../../viz/
