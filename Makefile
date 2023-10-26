include .envrc
run:
	echo ${OPEN_API_KEY}
	go build -o ~/go/bin/pa ./cmd/cli