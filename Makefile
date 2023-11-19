CC=go

test: templates pkg

templates:
	go test ./dsandalgo/templates/... -count=1 -cover -v

pkg:
	go test ./dsandalgo/pkg/... -count=1 -cover -v