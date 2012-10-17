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
	test-start-binary-server-stamp \
	test-start-compact-server-stamp \
	test-start-json-server-stamp

THRIFT = $(CURDIR)/../merge/thrift-$(THRIFT_VERSION)/compiler/cpp/thrift
#IDL = $(CURDIR)/../generation/enumm.thrift
#GENERATED = gen-java/thrift4go/generated/*.java
#GENERATED_TARGET = src/main/java/thrift4go/generated

all: test

test: test-stamp

test-stamp: test-exercise-stamp
	touch $@

test-exercise-stamp: test-start-binary-server-stamp test-start-json-server-stamp test-start-compact-server-stamp exercise_enumm_test.go
	cp exercise_enumm_test.go gen-go/enumm
	cd gen-go/enumm && go test -v -x .
	touch $@

test-start-binary-server-stamp: test-compile-stamp
	java -jar target/interoperability-unversioned-jar-with-dependencies.jar BINARY 8080 &
	sleep 5
	touch $@

test-start-json-server-stamp: test-compile-stamp
	java -jar target/interoperability-unversioned-jar-with-dependencies.jar JSON 8081 &
	sleep 5
	touch $@

test-start-compact-server-stamp: test-compile-stamp
	java -jar target/interoperability-unversioned-jar-with-dependencies.jar COMPACT 8082 &
	sleep 5
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
	#mkdir -vp "$(GENERATED_TARGET)"
	#mv $(GENERATED) "$(GENERATED_TARGET)"
	mkdir -vp "src/main/java/thrift4go/generated/enumm"
	mv gen-java/thrift4go/generated/enumm/*.java "src/main/java/thrift4go/generated/enumm"
	mkdir -vp "src/main/java/thrift4go/generated/listt"
	mv gen-java/thrift4go/generated/listt/*.java "src/main/java/thrift4go/generated/listt"
	touch $@

clean:
	rm -rf $(TEST_ARTIFACTS)
	-pkill -f 8080
	-pkill -f 8081
	-pkill -f 8082
	-pkill -f 8083

.PHONY: test
