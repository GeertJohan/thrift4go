namespace java thrift4go.generated.listt

enum SimpleEnum {
	Zero,
	One,
	Two,
	Three,
}

struct StructWithLists {
	1: list<i32>      listInt32,
	2: list<string>   listString,
	3: list<SimpleEnum> listSimpleEnum,
}

service ListsTestService {
	list<i32> echoListInt32(1: list<i32> input);
	list<string> echoListString(1: list<string> input);
	list<SimpleEnum> echoListSimpleEnum(1: list<SimpleEnum> input);
	StructWithLists echoStructWithLists(1: StructWithLists input);
}