build-image:
	docker build -t igorgalindop/finance .

run-app:
	docker-compose -f .devops/app.yml up -d

down-app:
	docker-compose -f .devops/app.yml down

lint:
	golint ./...
	go fmt -n ./...

unit_test:
	go test ./...