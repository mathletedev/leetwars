dev:
	air -build.cmd "make build app" -build.bin "./bin/app" -- -dev

build app:
	go build -o ./bin/app ./cmd/app

clean:
	go clean -modcache

db:
	psql leetwars < ./leetwars.sql
