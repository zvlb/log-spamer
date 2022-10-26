export MESSAGE_PER_SECOND=5
#export MESSAGE_SIZE=100
export MESSAGE_SIZE_FROM=10
export MESSAGE_SIZE_TO=40
	
# .PHONY: export_env
# export_env:
# 	@export MESSAGE_PER_SECOND=100
# 	@export MESSAGE_SIZE=100
# 	@export MESSAGES_SIZES_FROM=100
# 	@export MESSAGE_SIZE_TO=100

.PHONY: run
run:
	@go run ./cmd/app/main.go

.PHONY: test
test:
	@echo &MESSAGE_PER_SECOND &MESSAGE_SIZE &MESSAGES_SIZES_FROM &MESSAGE_SIZE_TO