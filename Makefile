build:
	mkdir -p functions
	cd funcsrc/hello-lambda
	go build -o ../../functions/hello-lambda