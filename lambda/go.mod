module moul.io/graphman/lambda

go 1.13

require (
	github.com/aws/aws-lambda-go v1.13.2
	github.com/pkg/errors v0.8.1
	gopkg.in/yaml.v3 v3.0.1
	moul.io/graphman v1.6.0
	moul.io/graphman/viz v1.5.0
)

replace (
	moul.io/graphman => ../
	moul.io/graphman/viz => ../viz/
)
