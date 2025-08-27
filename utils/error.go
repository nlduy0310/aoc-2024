package utils

import "log"

func PanicIf(err error) {

	if err != nil {
		panic(err)
	}
}

func ExitIf(err error) {

	if err != nil {
		log.Fatal(err)
	}
}

func ExitOn(cond bool, msg string) {

	if cond {
		log.Fatal(msg)
	}
}

func Assert(cond bool, msg string) {

	if !cond {
		panic(msg)
	}
}

func SafeCall(f func(), onPanic func(any)) {
	defer func() {
		if r := recover(); r != nil {
			onPanic(r)
		}
	}()
	f()
}
