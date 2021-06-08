build:
	go build -o bin/main main.go proposer.go util.go acceptor.go learner.go malicious.go

mal:
	go build -o bin/main main.go proposer.go util.go acceptor.go malicious.go learner.go

run: build
	./bin/main &
	./bin/main -p 5000 &
	./bin/main -p 8000 &

.PHONY: run

testlive: build
	./bin/main &
	./bin/main -p 5000 -l true &
	./bin/main -p 8000 &

.PHONY: testlive

testc: build
	./bin/main -w true &
	./bin/main -p 5000 -w true &
	./bin/main -p 8000 -w true &

.PHONY: testc

clean:
	rm -rf bin/*