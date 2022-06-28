.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		rpc/*.proto
.PHONY: clean
clean:
	rm -rf ./rpc/*.pb.go
.PHONY: bundle #打包证书
bundle:
	fyne bundle --package=config config/cert > config/bundled.go