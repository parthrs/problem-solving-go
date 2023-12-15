
test: twopointer pkg arrays stacks misc

misc:
	go test ./misc -cover -count=1 -v

twopointer:
	go test ./leet-code/twopointer -count=1 -cover -v

pkg:
	go test ./pkg -count=1 -cover -v

arrays:
	go test ./leet-code/arrays -count=1 -cover -v

stacks:
	go test ./leet-code/stacks -count=1 -cover -v
