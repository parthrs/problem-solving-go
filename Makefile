
test: twopointer pkg arrays stacks misc

misc:
	go test ./misc -cover -v

twopointer:
	go test ./leet-code/twopointer -cover -v

pkg:
	go test ./pkg/... -cover -v

arrays:
	go test ./leet-code/arrays -cover -v

stacks:
	go test ./leet-code/stacks -cover -v
