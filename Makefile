deploy:
	flyctl deploy

launch:
	flyctl launch

docker_init:
	docker build -t everybody8/quiz-api:v$(version) .

docker_push:
	docker push everybody8/quiz-api:v$(version)

swag:
	swag init