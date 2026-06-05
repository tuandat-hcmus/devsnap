CMD_DIR = cmd/devsnap
CMD_MAIN = $(CMD_DIR)/main.go
.PHONY: run
run: 
	go run $(CMD_MAIN)