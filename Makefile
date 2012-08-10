all:
	go build

tutorials: install
	for i in tutorial*; do (cd $$i; go build); done

fmt:
	go fmt

clean:
	rm *~

install:
	go install
