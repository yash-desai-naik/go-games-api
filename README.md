### Setup
1. create `.evn` file from `env.example`
2. mention you `DOCKERHUB_USERNAME`
3. run following commands:

```bash
go mod tidy

```

### Run
```bash
make run
```

### Test
```bash
make test
```
### Deploy

```bash
make build
make push
make deploy
```

### Undeploy
```bash
make undeploy
```

### List Pods
```bash
kubectl get pods --namespace=go-games
```