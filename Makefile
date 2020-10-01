build:
    pwd
	mkdir -p functions
	cd funcsrc/hello-lambda
	pwd
	ls -l
	go version
	go env
	go build -o ../../functions/hello-lambda