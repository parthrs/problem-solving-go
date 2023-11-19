CC=go

test: templates pkg

templates:
	go test ./dsandalgo/templates/... -count=1 -cover

pkg:
	go test ./dsandalgo/pkg/... -count=1 -cover