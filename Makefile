run:
	@read -p "Enter the Go file name (without .go): " filename; \
	dir=$$(find . -type f -name "$${filename}.go" -exec dirname {} \; | head -n 1); \
	if [ -n "$$dir" ]; then \
		echo "Found directory: $$dir"; \
		cd "$$dir" && time GO111MODULE=off go run "$${filename}.go"; \
	else \
		echo "Error: Go file '$${filename}.go' not found."; \
	fi
