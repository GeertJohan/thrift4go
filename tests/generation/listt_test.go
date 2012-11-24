package listt

import (
	"math"
	"testing"
	"thrift"
)

func TestStructWithListsNew(t *testing.T) {
	emission := NewStructWithLists()

	if emission == nil {
		t.Errorf("NewStructWithLists emitted nil, not the struct.")
	}
}

func TestStructWithListsPushAndIter(t *testing.T) {
	sl := NewStructWithLists()
	// ISSUE: The second parameter (int:12345) is not being used by the function. See tlist.go:35
	// ISSUE: Why do I have to create a new TList this way? Should imo be included in NewStructWithLists.
	sl.ListInt32 = thrift.NewTList(thrift.I32, 12345)
	sl.ListString = thrift.NewTList(thrift.STRING, 12345)
	// ISSUE: Using thrift.NewTListDefault() for the ListSimpleEnum fails. 
	//			The elemType for the TList is never properly set as TypeFromValue(anEnumValue) returns ttype STOP. (tlist.go:88)
	//			Therefore, every pushed enum is changed to `nil` before appended to the list by elemType.CoerceData. (tlist.go:90)
	sl.ListSimpleEnum = thrift.NewTList(thrift.ENUM, 12345)

	//## basetype int32, ListInt32
	i32TestValues := []int32{math.MinInt32, (math.MinInt32 + 1), -100, -42, -4, -2, -1, 0, 1, 2, 4, 42, 100, (math.MaxInt32 - 1), math.MaxInt32}
	for _, insertValue := range i32TestValues {
		sl.ListInt32.Push(insertValue)
	}
	i32Iter := sl.ListInt32.Iter()
	for idx, expectedValue := range i32TestValues {
		read, chanOpen := <-i32Iter
		if !chanOpen {
			t.Errorf("Iter channel for ListInt32 is closed. It is expected to be open as more values are expected. At #%d", idx)
			break
		}
		// ISSUE:  Imho type assertion should not be done in user-code, the generated code should have a list of int's (directly) for use here so run-time type assertion can be omited and compile optimizations can take place.
		//			This should also be the case with user types.. This shouldn't be just for basic types (int32, float, string) but also user-defined eidxs/structs..
		//			Example: thrift IDL `list<MyStructDefinition>` should generate go code `type TListMyStructDefinition`
		readValue, ok := read.(int32)
		if !ok {
			t.Errorf("A value read from the ListInt32 could not be type-asserted to be an int32. This should not occur.")
		}
		if expectedValue != readValue {
			t.Errorf("i32 push/iter failed on #%d. Expected: %d. Received: %d.", idx, expectedValue, readValue)
		}
	}
	if _, chanOpen := <-i32Iter; chanOpen {
		t.Error("Iter channel for ListInt32 is open. It is expected to be close as expected values have been read already.")
	}

	//## basetype string, ListString
	stringTestValues := []string{
		"Time is an illusion. Lunchtime doubly so.",
		"The ships hung in the sky in much the same way that bricks don't.",
		"You've got to know where your towel is.",
		"DON'T PANIC!",
		"Special chars: \n ' \" \\ ",
	}
	for _, insertValue := range stringTestValues {
		sl.ListString.Push(insertValue)
	}
	stringIter := sl.ListString.Iter()
	for idx, expectedValue := range stringTestValues {
		read, chanOpen := <-stringIter
		if !chanOpen {
			t.Errorf("Iter channel for ListString is closed. It is expected to be open as more values are expected. At #%d", idx)
			break
		}
		readValue, ok := read.(string)
		if !ok {
			t.Errorf("A value read from the ListString could not be type-asserted to be an string. This should not occur.")
		}
		if expectedValue != readValue {
			t.Errorf("string push/iter failed on #%d. Expected: '%s'. Received: '%s'.", idx, expectedValue, readValue)
		}
	}
	if _, chanOpen := <-stringIter; chanOpen {
		t.Error("Iter channel for ListString is open. It is expected to be close as expected values have been read already.")
	}

	//## enum SimpleEnum, ListSimpleEnum
	simpleEnumTestValues := []SimpleEnum{
		SimpleEnum_Zero,
		SimpleEnum_One,
		SimpleEnum_Two,
		SimpleEnum_Three,
	}
	for _, insertValue := range simpleEnumTestValues {
		sl.ListSimpleEnum.Push(insertValue)
	}
	simpleEnumIter := sl.ListSimpleEnum.Iter()
	for idx, expectedValue := range simpleEnumTestValues {
		read, chanOpen := <-simpleEnumIter
		if !chanOpen {
			t.Errorf("Iter channel for ListSimpleEnum is closed. It is expected to be open as more values are expected. At #%d", idx)
			break
		}
		// The curent ttype implementation handles enum as int64.
		readInt64, ok := read.(int64)
		if !ok {
			t.Errorf("A value read from the ListSimpleEnum could not be type-asserted to be an SimpleEnum. This should not occur. Raw value: %v. Go value: %#v. Go type: %T.", read, read, read)
		} else {
			// convert readInt64 to our enum type again..
			readValue := SimpleEnum(readInt64)
			if expectedValue != readValue {
				t.Errorf("SimpleEnum push/iter failed on #%d. Expected: '%s'. Received: '%s'.", idx, expectedValue, readValue)
			}
		}
	}
	if _, chanOpen := <-simpleEnumIter; chanOpen {
		t.Error("Iter channel for ListSimpleEnum is open. It is expected to be close as expected values have been read already.")
	}
}
