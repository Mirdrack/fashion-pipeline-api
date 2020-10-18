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
    golang:1.15 ./app
```
