# fashion-pipeline-api

## Compile
```
docker run --rm -it \
    -v $(pwd):/go/src/app \
    -w /go/src/app \
    golang:1.15 go build
```

## Run
```
docker run --rm -it \
    -v $(pwd):/go/src/app \
    -w /go/src/app \
    -p 8000:8000 \
    -e PORT=8000 \
    golang:1.15 ./app
```

## Lint
```
docker run --rm \
    -v $(pwd):/app \
    -w /app \
    golangci/golangci-lint:v1.31.0 golangci-lint run -v
```
