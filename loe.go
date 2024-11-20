package xoe

func Loe(err error) bool {
	if err != nil {
		logger.Println("Loe:", err)
		return true
	}
	return false
}

func Loe1[T1 any](t1 T1, err error) T1 {
	if err != nil {
		logger.Println("Loe1:", err)
	}
	return t1
}

func Loe2[T1 any, T2 any](t1 T1, t2 T2, err error) (T1, T2) {
	if err != nil {
		logger.Println("Loe2:", err)
	}
	return t1, t2
}

func Loe3[T1 any, T2 any, T3 any](t1 T1, t2 T2, t3 T3, err error) (T1, T2, T3) {
	if err != nil {
		logger.Println("Loe3:", err)
	}
	return t1, t2, t3
}

func Loe4[T1 any, T2 any, T3 any, T4 any](t1 T1, t2 T2, t3 T3, t4 T4, err error) (T1, T2, T3, T4) {
	if err != nil {
		logger.Println("Loe4:", err)
	}
	return t1, t2, t3, t4
}

func Loe5[T1 any, T2 any, T3 any, T4 any, T5 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, err error) (T1, T2, T3, T4, T5) {
	if err != nil {
		logger.Println("Loe5:", err)
	}
	return t1, t2, t3, t4, t5
}

func Loe6[T1 any, T2 any, T3 any, T4 any, T5 any, T6 any](t1 T1, t2 T2, t3 T3, t4 T4, t5 T5, t6 T6, err error) (T1, T2, T3, T4, T5, T6) {
	if err != nil {
		logger.Println("Loe6:", err)
	}
	return t1, t2, t3, t4, t5, t6
}
