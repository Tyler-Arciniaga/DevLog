build:
	go build -o devlog
	@echo "Most up-to-date DevLog built"

install:
	go build -o devlog
	sudo mv devlog /usr/local/bin
	@echo "Installed DevLog to /usr/local/bin"

clean:
	rm -rf devlog
	@echo "Removed devlog binary file from working directory"