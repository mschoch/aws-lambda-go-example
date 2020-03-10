build:
	mkdir -p functions
	go get ./...
	go build -o functions/hello-lambda ./...
	mkdir -p functions/data
	cp data.js functions/data/
	echo "raw data" > functions/data/data.txt
