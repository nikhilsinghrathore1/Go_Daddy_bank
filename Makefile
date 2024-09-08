build: 
	@go build -o bin/Go_Daddy_bank

run: build 
	@./bin/Go_Daddy_bank

test : 
	@go test -v ./... 