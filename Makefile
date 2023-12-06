
test: twopointer pkg arrays stringrelated stacks misc

misc:
	go test ./misc -count=1 -cover -v

twopointer:
	go test ./dsandalgo/twopointer -count=1 -cover -v

pkg:
	go test ./dsandalgo/pkg/... -count=1 -cover -v

arrays:
	go test ./dsandalgo/arrays -count=1 -cover -v

stringrelated:
	go test ./dsandalgo/stringrelated -count=1 -cover -v

stacks:
	go test ./dsandalgo/stacks -count=1 -cover -v
