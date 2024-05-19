
.PHONY: run_debug
run_debug:
	# git pull origin main;
	DEBUG_ALL=true k6 run script.js > output.txt 2>&1;

.PHONY: run
run:
	# git pull origin main;
	k6 run script.js;

.PHONY: runLoadTest
runLoadTest:
	# git pull origin main;
	LOAD_TEST=true k6 run script.js;

build_pb:
	protoc --go_out=entity/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=entity/pb \
		--go-grpc_opt=paths=source_relative *.proto
run_be:
	go run main.go