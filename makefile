
build:
	docker build --no-cache --build-arg sci_key=${SCIKEY} -t simple-ci .
run:
	docker run -it -v ${shell pwd}/src:/go/src/app/ --rm -p 4545:4545 simple-ci go run main.go
