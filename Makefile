#application
run:
	go run cmd/api/main.go

#docker
docker-compose:
	docker-compose up -d 

docker-build:
	docker build -t bundles:1.0 .

docker-images:
	docker images

docker-run:
	docker run -p 8080:8080 bundles:1.0

