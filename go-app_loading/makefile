.PHONY: build run dev

SVC_EXEC=loading_svc

build:
	@GOARCH=wasm GOOS=js go build -C cmd/loading_pwa -o ../../web/app.wasm
	@go build -C cmd/loading_svc -o ../../$(SVC_EXEC)

start_svc: build
	@./$(SVC_EXEC)

stop_svc:
	@killall -q -HUP $(SVC_EXEC) || echo

dev: stop_svc start_svc

