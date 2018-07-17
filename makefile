build/fetch_genome_mac: main.go download.go const.go
	env GOOS=darwin GOARCH=amd64 go build -o $@ $^

build/fetch_genome_windows: main.go download.go const.go
	env GOOS=darwin GOARCH=amd64 go build -o $@ $^

build/fetch_genome: main.go download.go const.go
	go build -o $@ $^

.PHONY: all
all: build/fetch_genome build/fetch_genome_windows build/fetch_genome_mac

.PHONY: clean
clean:
	rm -rf build