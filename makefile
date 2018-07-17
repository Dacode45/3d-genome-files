build/fetch_genome_mac: main.go download.go const.go
	env GOOS=darwin GOARCH=amd64 go build -o $@ $^

build/fetch_genome_windows.exe: main.go download.go const.go
	env GOOS=windows GOARCH=amd64 go build -o $@ $^

build/fetch_genome_linux: main.go download.go const.go
	env GOOS=linux GOARCH=amd64 go build -o $@ $^

.PHONY: all
all: build/fetch_genome_linux build/fetch_genome_windows.exe build/fetch_genome_mac

.PHONY: clean
clean:
	rm -rf build