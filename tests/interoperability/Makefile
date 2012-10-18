export GOPATH  = $(CURDIR)/../../lib/go

TEST_ARTIFACTS = \
	gen-go \
	gen-java \
	src/main/java/thrift4go/generated/* \
	target \
	test-compile-stamp \
	test-exercise-stamp \
	test-generation-stamp \
	test-stamp \
	test-start-servers-stamp

THRIFT = $(CURDIR)/../merge/thrift-$(THRIFT_VERSION)/compiler/cpp/thrift

all: test

test: test-stamp

test-stamp: test-exercise-stamp
	touch $@

test-exercise-stamp: test-start-servers-stamp exercise_enumm_test.go
	cp exercise_enumm_test.go gen-go/enumm
	cd gen-go/enumm && go test -v -x .
	touch $@

test-start-servers-stamp: test-compile-stamp
	java -jar target/interoperability-unversioned-jar-with-dependencies.jar &
	sleep 10
	touch $@

test-compile-stamp: test-generation-stamp
	mvn compile
	mvn assembly:assembly
	mvn assembly:single
	touch $@

test-generation-stamp:
	"$(THRIFT)" --gen java "$(CURDIR)/../generation/enumm.thrift"
	"$(THRIFT)" --gen go "$(CURDIR)/../generation/enumm.thrift"
	"$(THRIFT)" --gen java "$(CURDIR)/../generation/listt.thrift"
	"$(THRIFT)" --gen go "$(CURDIR)/../generation/listt.thrift"
	mkdir -vp "src/main/java/thrift4go/generated/enumm"
	mv gen-java/thrift4go/generated/enumm/*.java "src/main/java/thrift4go/generated/enumm"
	mkdir -vp "src/main/java/thrift4go/generated/listt"
	mv gen-java/thrift4go/generated/listt/*.java "src/main/java/thrift4go/generated/listt"
	touch $@

clean:
	rm -rf $(TEST_ARTIFACTS)

	# kill servers for enumm
	-pkill -f 8080 # binary
	-pkill -f 8081 # json
	-pkill -f 8082 # compact

	# kill servers for listt
	-pkill -f 8180 # binary
	-pkill -f 8181 # json
	-pkill -f 8182 # compact

.PHONY: test
