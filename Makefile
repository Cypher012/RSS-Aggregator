BINARY=rss-aggregator
DB_URL="postgres://cipher:cipher2017@localhost:5432/rss-db?sslmode=disable"
SRC=$(wildcard *.go)

.PHONY: run dev clean file migrate-up migrate-down

$(BINARY): $(SRC)
	go build -o $(BINARY)

run: $(BINARY)
	./$(BINARY)

dev: run

clean:
	rm -f $(BINARY)

%:
	touch $@ && zed ./$@

migrate-up:
	goose -dir sql/schema postgres $(DB_URL) up

migrate-down:
	goose -dir sql/schema postgres $(DB_URL) down
