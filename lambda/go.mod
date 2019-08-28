module moul.io/graphman/lambda

go 1.12

require (
	github.com/aws/aws-lambda-go v1.11.1
	github.com/pkg/errors v0.8.1
	gopkg.in/yaml.v3 v3.0.0-20190709130402-674ba3eaed22
	moul.io/graphman v1.3.0
	moul.io/graphman/viz v1.2.0
)

replace moul.io/graphman => ../

replace moul.io/graphman/viz => ../viz/
