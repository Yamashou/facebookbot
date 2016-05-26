.PHONY: all get clean test run

all:
	gofmt -w .
	go install -v -gcflags -N "github.com/ledyba/go-cabocha/..."

clean:
	go clean -i "github.com/ledyba/go-cabocha"

test:
	go test "github.com/ledyba/go-cabocha"
