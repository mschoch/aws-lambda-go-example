build:
	mkdir -p functions
	go get ./...
	go build -o functions/hello-lambda ./...
	mkdir -p functions/data
	cp functions/hello-lambda functions/data/data
	echo "raw data" > functions/data/data.txt
	zip -rj functions/data.zip functions/data
	unzip -l functions/data.zip
	rm -r functions/data
