module moul.io/graphman/lambda

go 1.13

require (
	github.com/aws/aws-lambda-go v1.13.2
	github.com/pkg/errors v0.9.1
	gopkg.in/yaml.v3 v3.0.0-20190924164351-c8b7dadae555
	moul.io/graphman v1.6.0
	moul.io/graphman/viz v1.5.0
)

replace (
	moul.io/graphman => ../
	moul.io/graphman/viz => ../viz/
)
