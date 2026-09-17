package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_init_var_from_num(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if int32(0) <= v10 {
		v13 = int32(-8)
	} else {
		v13 = int32(-6)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(base.Ui32(int32(base.Ui32(v5)>>(uint(int32(2))%32))+v13) >> (uint(int32(1)) % 32))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v18 < int32(0) {
		v22 = v18 & int32(_a_F_init_var_from_num_0)
		v34 = v22<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v22&int32(63)
	} else {
		v32 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
		v34 = v32
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v34
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v37 = int32(_a_F_init_var_from_num_1)
	v38 = v36 & v37
	if v38 != v37 {
		if v38 != int32(_a_F_init_var_from_num_2) {
			v49 = v38
		} else {
			v49 = v36 << (uint(int32(1)) % 32) & int32(_a_F_init_var_from_num_3)
		}
	} else {
		v49 = v36 & int32(_a_F_init_var_from_num_4)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v49
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v51 < int32(0) {
		v60 = int32(base.Ui32(v51)>>(uint(int32(7))%32)) & int32(63)
	} else {
		v60 = v51 & int32(_a_F_init_var_from_num_5)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v60
	v62 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+4)))
	v63 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v63
	if v62 < v63 {
		v69 = int32(6)
	} else {
		v69 = int32(8)
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = l0 + v69
	return
}
func F_markVarForSelectPriv(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v6 == int32(0) {
		v49 = l0
	} else {
		v10 = v6 & int32(7)
		if base.Ui32(int32(8)) <= base.Ui32(v6) {
			v16 = l0
			v18 = int32(0)
			for {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
				v30 = v18 + int32(8)
				if v30 != v6&int32(-8) {
					v16 = v28
					v18 = v30
					continue
				} else {
					break
				}
				break
			}
			if v10 == int32(0) {
				v49 = v28
			} else {
				v34 = v28
				v40 = v34
				v42 = int32(0)
				for {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					v47 = v42 + int32(1)
					if v47 != v10 {
						v40 = v45
						v42 = v47
						continue
					} else {
						break
					}
					break
				}
				v49 = v45
			}
		} else {
			v34 = l0
			v40 = v34
			v42 = int32(0)
			for {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
				v47 = v42 + int32(1)
				if v47 != v10 {
					v40 = v45
					v42 = v47
					continue
				} else {
					break
				}
				break
			}
			v49 = v45
		}
	}
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+8)))
	F_markRTEForSelectPriv(m, v49, v54, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		return
	} else {
		return
	}
}
