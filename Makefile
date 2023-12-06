
test: twopointer pkg arrays stringrelated stacks misc

misc:
	go test ./misc -cover -v

twopointer:
	go test ./dsandalgo/twopointer -cover -v

pkg:
	go test ./dsandalgo/pkg/... -cover -v

arrays:
	go test ./dsandalgo/arrays -cover -v

stringrelated:
	go test ./dsandalgo/stringrelated -cover -v

stacks:
	go test ./dsandalgo/stacks -cover -v
