build:
	pwd
	mkdir -p functions
	go version
	go env
	cd funcsrc/hello-lambda; go build -o ../../functions/hello-lambda