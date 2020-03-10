build:
	mkdir -p functions
	go get ./...
	go build -o functions/hello-lambda ./...
	echo "test" > functions/foo.txt
	mkdir -p functions/subdir
	echo "test" > functions/subdir/bar.txt
