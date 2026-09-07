module moul.io/graphman/lambda

go 1.26

require (
	github.com/aws/aws-lambda-go v1.55.0
	github.com/pkg/errors v0.9.1
	gopkg.in/yaml.v3 v3.0.1
	moul.io/graphman v1.6.0
	moul.io/graphman/viz v1.5.0
)

require github.com/awalterschulze/gographviz v0.0.0-20190522210029-fa59802746ab // indirect

replace (
	moul.io/graphman => ../
	moul.io/graphman/viz => ../viz/
)
