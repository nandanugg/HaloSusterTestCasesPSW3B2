
.PHONY: run_debug
run_debug:
	git pull origin main;
	DEBUG_ALL=true k6 run script.js > output.txt 2>&1

.PHONY: run
run:
	git pull origin main;
	k6 run script.js

.PHONY: runLoadTest
runLoadTest:
	git pull origin main;
	LOAD_TEST=true k6 run script.js

start_db:
	docker run --name k6_postgres \
		-e POSTGRES_DB=k6 \
		-e POSTGRES_USER=postgres \
		-e POSTGRES_PASSWORD=postgres \
		-d -p 5432:5432 postgres:16.3-alpine3.19
clean_db:
	docker stop k6_postgres
	docker rm k6_postgres