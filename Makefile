build:
	pwd
	mkdir -p functions
	go version
	go env
	go build -o ../../functions/hello-lambda ./funcsrc/hello-lambda