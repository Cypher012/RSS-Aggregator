BINARY=rss-aggregator
SRC=$(wildcard *.go)

.PHONY: run dev clean file

$(BINARY): $(SRC)
	go build -o $(BINARY)

run: $(BINARY)
	./$(BINARY)

dev: run

clean:
	rm -f $(BINARY)

%:
	touch $@ && zed ./$@
