serve:
	cd website; yarn start

build:
	cd website; yarn build
	cd website && go generate
	go build -trimpath