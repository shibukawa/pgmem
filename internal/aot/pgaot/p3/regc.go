package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_regc_wc_isalpha(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_regc_wc_isalpha[0]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
	if v6 == int32(1) {
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v23 = int32(0)
			return v23
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_regc_wc_isalpha[1]))))
			v12 = int32(1)
			return int32(base.Ui32(v11)>>(uint(v12)%32)) & v12
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+20))
		v19 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, l0, v5)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = v19
			return v23
		}
	}
}
func F_regc_wc_ispunct(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_regc_wc_ispunct[0]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
	if v6 == int32(1) {
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v23 = int32(0)
			return v23
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_regc_wc_ispunct[1]))))
			return int32(base.Ui32(v11)>>(uint(int32(6))%32)) & int32(1)
		}
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+44))
		v19 = m.T0[v18].(func(*base.Module, int32, int32) int32)(m, l0, v5)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = v19
			return v23
		}
	}
}
func F_regc_wc_isspace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_regc_wc_isspace[0]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+2)))
	if v6 == int32(1) {
		if base.Ui32(int32(127)) < base.Ui32(l0) {
			v21 = int32(0)
			return v21
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_regc_wc_isspace[1]))))
			return int32(base.Ui32(v11) >> (uint(int32(7)) % 32))
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
		v17 = m.T0[v16].(func(*base.Module, int32, int32) int32)(m, l0, v5)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v21 = v17
			return v21
		}
	}
}
func F_regc_wc_isword(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	if l0 == int32(95) {
		v27 = int32(1)
		return v27
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_regc_wc_isword[0]))
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+2)))
		if v10 == int32(1) {
			if base.Ui32(int32(127)) < base.Ui32(l0) {
				v27 = int32(0)
				return v27
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_regc_wc_isword[1]))))
				return base.B2i32(v15&int32(3) != int32(0))
			}
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
			v23 = m.T0[v22].(func(*base.Module, int32, int32) int32)(m, l0, v9)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				v27 = v23
				return v27
			}
		}
	}
}
