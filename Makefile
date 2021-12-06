.PHONY: start

EMAIL ?= methrilion@gmail.com

# usage:
# make start email=your_email@example.com password=your_password
start:
	EMAIL=$(email) PASSWORD=$(password) go run main.go