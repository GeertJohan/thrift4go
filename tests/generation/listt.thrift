namespace java thrift4go.generated.listt

enum TestEnum {
	AAA,
	BBB,
	CCC,
	DDD,
	EEE,
}

struct WrappedLists {
	1: list<i32>      listInt32,
	2: list<string>   listString,
	3: list<TestEnum> listTestEnum,
}

service ListsTestService {
	WrappedLists echoWrappedLists(1: WrappedLists input);
	list<i32> echoInt32List(1: list<i32> input);
	list<string> echoStringList(1: list<string> input);
	list<TestEnum> echoTestEnumList(1: list<TestEnum> input);
}