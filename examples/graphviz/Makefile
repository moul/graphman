.PHONY: run
run:
	set -eo pipefail; go run . | tee /tmp/graphman.dot
	set -eo pipefail; cat /tmp/graphman.dot | dot -Tsvg > /tmp/graphman.svg
	open /tmp/graphman.svg
