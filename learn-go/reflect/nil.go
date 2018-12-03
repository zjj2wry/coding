// +build OMIT

package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := (*interface{})(nil)
	fmt.Println(reflect.TypeOf(a), reflect.ValueOf(a))
	var b interface{} = (*interface{})(nil)
	fmt.Println(reflect.TypeOf(b), reflect.ValueOf(b))
	fmt.Println(a == nil, b == nil)
}

/*
*interface {} <nil>
*interface {} <nil>
true false

Under the covers, interfaces are implemented as two elements, a type and a value.
The value, called the interface's dynamic value, is an arbitrary concrete value and the type
is that of the value. For the int value 3, an interface value contains, schematically, (int, 3).
An interface value is nil only if the inner value and type are both unset, (nil, nil). In particular,
a nil interface will always hold a nil type. If we store a nil pointer of type *int inside an
interface value, the inner type will be *int regardless of the value of the pointer: (*int, nil).
Such an interface value will therefore be non-nil even when the pointer inside is nil.


a := (*interface{})(nil) is equal with var a *interface{} = nil.
but var b interface{} = (*interface{})(nil) , mean b is type interface{}, and interface{} variable only
nil when it's type and value are both nil, obviously type *interface{} is not nil
*/
