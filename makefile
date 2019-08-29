build:
	CGO_ENABLE=0 go build -o add_curr_days main.go

run:
	go run main.go 0

copy:
	cp add_curr_days ~/.local/bin/
