module moul.io/graphman/examples

go 1.13

require (
	moul.io/graphman v1.5.0
	moul.io/graphman/viz v1.5.0
)

replace (
	moul.io/graphman => ../
	moul.io/graphman/viz => ../viz/
)
