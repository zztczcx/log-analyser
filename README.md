## Task
parse a file and make a statistic

## Assumptions

1. The file to process could be very big 
2. When the url or ip have the same visits, we get the top 3 by alphabetic order.
3. Assume the memory is enough.

## Design

This program used a lot goroutines to speed up the processing which is good for processing a big log file.
Basically we treat this process as mapReduce.

producer ---(fan out)--> parser --> analyse ---(fan in)--> reducer --> report


## How to run it

```
git clone git@github.com:zztczcx/log-analyser.git 

cd log-analyser

docker build -t log-analyser .
docker run -v ./testdata/:/app/testdata --name=log-analyser --rm log-analyser  -input="./testdata/http.log"
```

## Testing

```
go test
```
