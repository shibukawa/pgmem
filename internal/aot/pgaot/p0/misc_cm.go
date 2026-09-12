package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CMPTRGM_CHOOSE(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_CMPTRGM_CHOOSE[0]))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+256)))
	if v9 != 0 {
		v10 = int32(_a_F_CMPTRGM_CHOOSE_0)
	} else {
		v10 = int32(_a_F_CMPTRGM_CHOOSE_1)
	}
	*(*int32)(unsafe.Add(mBase, _c_F_CMPTRGM_CHOOSE[1])) = v10
	v12 = m.T0[v10].(func(*base.Module, int32, int32) int32)(m, l0, l1)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		return v12
	}
}
func F_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v67 int32
	_ = v67
	v5 = l2 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v5) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v67
L2:
	;
	v67 = int32(0)
	goto L1
L3:
	;
	v41 = v36
	v42 = v37
	v43 = v38
	goto L13
L4:
	;
	if (l0|l1)&int32(3) != 0 {
		v36 = l0
		v37 = l1
		v38 = v5
		goto L3
	} else {
		goto L7
	}
L5:
	;
	v29 = l0
	v30 = l1
	v31 = v5
	goto L6
L6:
	;
	if v31 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v13 = l0
	v14 = l1
	v15 = v5
	goto L8
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v18 != v19 {
		v36 = v13
		v37 = v14
		v38 = v15
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v29 = v24
	v30 = v22
	v31 = v26
	goto L6
L10:
	;
	v21 = int32(4)
	v22 = v14 + v21
	v24 = v13 + v21
	v26 = v15 - v21
	if base.Ui32(int32(3)) < base.Ui32(v26) {
		v13 = v24
		v14 = v22
		v15 = v26
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v36 = v29
	v37 = v30
	v38 = v31
	goto L3
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	if v46 == v47 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v67 = v46 - v47
	goto L1
L15:
	;
	v49 = int32(1)
	v54 = v43 - v49
	if v54 != 0 {
		v41 = v41 + v49
		v42 = v42 + v49
		v43 = v54
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L2
}
func F_cmpEntries(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v6 = int32(1)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	if v8 == v6 {
		if v7&int32(1) != 0 {
			v25 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v25)
			v28 = int32(0)
		} else {
			v28 = v6
		}
		return v28
	} else {
		if v7&int32(1) != 0 {
			v28 = int32(-1)
			return v28
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v20 = F_FunctionCall2Coll(m, v16, v17, v18, v19)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v20 != 0 {
					v28 = v20
				} else {
					v25 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)) = uint8(v25)
					v28 = int32(0)
				}
				return v28
			}
		}
	}
}
func F_cmp_lbestatus(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+408))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+408))
	return v3 - v4
}
