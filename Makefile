CC=go

test: twopointer pkg arrays stringrelated

twopointer:
	go test ./dsandalgo/twopointer -count=1 -cover -v

pkg:
	go test ./dsandalgo/pkg/... -count=1 -cover -v

arrays:
	go test ./dsandalgo/arrays -count=1 -cover -v

stringrelated:
	go test ./dsandalgo/stringrelated -count=1 -cover -v