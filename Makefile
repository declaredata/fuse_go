.PHONY: test
test:
	go test ./...

.PHONY: download-fuse-proto
download-proto:
	curl -o fuse.proto -L https://raw.githubusercontent.com/declaredata/fuse_python/refs/heads/main/proto/sds.proto

.PHONY: fuse-proto
gen-proto:
	buf generate
	
