namespace java thrift4go.generated.listt

enum UndefinedValues {
	One,
	Two,
	Three,
}

struct StructWithLists {
	1: list<i32>      listInt32,
	2: list<string>   listString,
	3: list<UndefinedValues> listUndefinedValues,
}

service ListsTestService {
	list<i32> echoInt32List(1: list<i32> input);
	list<string> echoStringList(1: list<string> input);
	list<UndefinedValues> echoUndefinedValuesList(1: list<UndefinedValues> input);
	StructWithLists echoStructWithLists(1: StructWithLists input);
}