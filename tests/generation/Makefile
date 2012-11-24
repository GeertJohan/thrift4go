export GOPATH  = $(CURDIR)/../../lib/go

TEST_ARTIFACTS = \
	gen-go \
	test-compile-stamp \
	test-exercise-stamp \
	test-generation-stamp \
	test-stamp \
	test-validate-stamp

THRIFT=$(CURDIR)/../merge/thrift-$(THRIFT_VERSION)/compiler/cpp/thrift

all: test

test: test-stamp

test-stamp: test-exercise-stamp
	touch $@

test-exercise-stamp: test-compile-stamp enumm_test.go listt_test.go
	cp -f enumm_test.go gen-go/enumm
	cp -f listt_test.go gen-go/listt
	cd gen-go/enumm && go test -v -x .
	cd gen-go/listt && go test -v -x .
	touch $@

test-compile-stamp: test-validate-stamp
	cd gen-go/enumm && go build -v -x .
	cd gen-go/listt && go build -v -x .
	touch $@

test-validate-stamp: test-generation-stamp
	# TODO(GeertJohan): enable golden diff test 
	# diff -Nru golden/gen-go/ gen-go/
	touch $@

test-generation-stamp:
	$(THRIFT) --gen go enumm.thrift
	$(THRIFT) --gen go listt.thrift
	touch $@

clean:
	rm -rf $(TEST_ARTIFACTS)

.PHONY: test
