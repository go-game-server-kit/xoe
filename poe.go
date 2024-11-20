package xoe

func Poe(err error) bool {
	if err != nil {
		panic(err)
	}
	return false
}

func Poe1[T1 any](t1 T1, err error) T1 {
	if err != nil {
		panic(err)
	}
	return t1
}

func Poe2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		panic(err)
	}
	return t1, t2
}

func Poe3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3
}

func Poe4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4
}

func Poe5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4, t5
}

func Poe6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	if err != nil {
		panic(err)
	}
	return t1, t2, t3, t4, t5, t6
}
