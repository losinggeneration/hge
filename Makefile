GOOPTS?=-tags "opengl 1_4"
DIRS=hge legacy helpers/animation helpers/color helpers/distortionmesh helpers/font helpers/gui helpers/guictrls helpers/particle helpers/rect helpers/sprite helpers/strings helpers/vector

all:
	for i in $(DIRS); do (cd $$i; go build $(GOOPTS)); done

tutorials: install
	for i in tutorials/tutorial*; do (cd $$i; go build $(GOOPTS)); done

fmt:
	for i in $(DIRS); do (cd $$i; go fmt); done

clean:
	find . -name "*~" -delete
	for i in tutorials/tutorial*; do rm -f $$i/$$(basename $$i); done

install:
	for i in $(DIRS); do (cd $$i; go install $(GOOPTS)); done

.PHONY: all tutorials fmt clean install
