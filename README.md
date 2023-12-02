## Task
parse a file and make a statistic

## Assumptions

1. The file to process could be very big 
2. When the url or ip have the same visits, we get the top 3 by alphabetic order.
3. Assume the memory is enough.

## Design

It uses channel and goroutines to decouple each part of the processing and made a pipeline.
we can treat this process as a MapReduce

producer ---(fan out)--> parser --> analyse ---(fan in)--> reducer --> report


## How to run it

### using docker

```
git clone git@github.com:zztczcx/log-analyser.git 

cd log-analyser

docker build -t log-analyser .
docker run -v ./testdata/:/app/testdata --name=log-analyser --rm log-analyser  -input="./testdata/http.log"
```

### command line

```
go run main.go -input="./testdata/http.log"
```

## Testing

```
go test ./...
```

## TODO

Currently the producer is reading file one by one, when the file is getting bigger, 
we can try to use multiple gorountes to read the file.
