all: esper

again: clean all

esper: esper.go cmd/esper/main.go
	go build -o esper cmd/esper/main.go

clean:
	rm -f esper

test:
	go test -cover

push:
	got send
	git push github

fmt:
	gofmt -s -w *.go cmd/*/main.go

cover:
	go test -coverprofile=cover.out
	go tool cover -html cover.out

README.md: README.gmi
	sisyphus -f markdown <README.gmi >README.md

doc: README.md

release: push
	git push github --tags
	got send -T
