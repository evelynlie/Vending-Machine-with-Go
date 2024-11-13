# Define the main binary name
BINARY := vm

# Find all .go files
GO_FILES := $(wildcard *.go)

# Targets
.PHONY: all build clean

# Build the final binary
build: $(BINARY)

$(BINARY): $(GO_FILES)
	go build -o $(BINARY) $(GO_FILES)

# Clean up the binary
clean:
	rm -f $(BINARY)