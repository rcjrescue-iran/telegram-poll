NOW := $(shell date +%s)

all: 	
	@go build -ldflags "-X main.buildVersion=$(NOW)" -i -o rcjrescue_server
run: all
	@./rcjrescue_server