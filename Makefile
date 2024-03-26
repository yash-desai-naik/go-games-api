# Variables
IMAGE_NAME = go-games-api
DOCKER_RUN_OPTIONS = --rm -p 8080:8080
KUBECTL = kubectl

build:
	docker build -t go-games-api .
run:
	docker run --rm -p 8080:8080 go-games-api
test:
	docker run --rm go-games-api go test -v ./...
# Deploy to Kubernetes cluster
deploy:
	$(KUBECTL) apply -f deployment.yaml
	$(KUBECTL) apply -f service.yaml

# Delete from Kubernetes cluster
undeploy:
	$(KUBECTL) delete -f deployment.yaml
	$(KUBECTL) delete -f service.yaml