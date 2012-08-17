DIRS=hge legacy helpers/animation helpers/color helpers/distortionmesh helpers/font helpers/gui helpers/guictrls helpers/particle helpers/rect helpers/sprite helpers/strings helpers/vector

all:
	for i in $(DIRS); do (cd $$i; go build); done

tutorials: install
	for i in tutorials/tutorial*; do (cd $$i; go build); done

fmt:
	go fmt

clean:
	rm *~
	for i in tutorials/tutorial*; do rm $$i/$$i; done

install:
	for i in $(DIRS); do (cd $$i; go install); done
