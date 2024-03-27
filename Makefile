# Makefile for Go Games API

# Variables
IMAGE_NAME = go-games-api
DOCKERHUB_USERNAME = "<YOUR-Dockerhub-USERNAME>"

# Build Docker image
build:
	docker build -t $(DOCKERHUB_USERNAME)/$(IMAGE_NAME):latest .

# Push Docker image to Docker Hub
push:
	docker login 
	docker push $(DOCKERHUB_USERNAME)/$(IMAGE_NAME):latest

# Run Docker container
run:
	docker run -d -p 8080:8080 $(DOCKERHUB_USERNAME)/$(IMAGE_NAME):latest

# Test the application
test:
	docker run --rm -p 8081:8080 $(DOCKERHUB_USERNAME)/$(IMAGE_NAME):latest go test -v ./...

# Deploy to Kubernetes cluster
deploy:
	@kubectl get namespace go-games >/dev/null 2>&1 || kubectl create namespace go-games
	kubectl apply -f deployments.yaml

# Delete from Kubernetes cluster
undeploy:
	kubectl delete -f deployments.yaml

# List number of pods running
pods:
	kubectl get pods --namespace=go-games
