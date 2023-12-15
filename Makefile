test: misc pkg lc

misc:
	go test ./misc -cover -count=1 -v

pkg:
	go test ./pkg -count=1 -cover -v

lc:
	go test ./leet-code/... -count=1 -cover -v
