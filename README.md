**Run the Service:**
code>docker-compose up</code>.

**Run the test:**
code>docker-compose run --rm web go test ./test/  -coverpkg=./... -coverprofile ./coverage.out</code>.