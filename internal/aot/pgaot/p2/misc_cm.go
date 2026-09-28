package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CMPTRGM_UNSIGNED(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != v6 {
		v21 = v5
		v22 = v6
		if base.Ui32(v21) < base.Ui32(v22) {
			v26 = int32(-1)
		} else {
			v26 = int32(1)
		}
		return v26
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v8 != v9 {
			v21 = v8
			v22 = v9
			if base.Ui32(v21) < base.Ui32(v22) {
				v26 = int32(-1)
			} else {
				v26 = int32(1)
			}
			return v26
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			if v11 != v12 {
				if base.Ui32(v11) < base.Ui32(v12) {
					v17 = int32(-1)
				} else {
					v17 = int32(1)
				}
				v19 = v17
			} else {
				v19 = int32(0)
			}
			return v19
		}
	}
}
func F_cmpaliases(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if base.B2i32(v7 == int32(0))|base.B2i32(v7 != v10) != 0 {
		v28 = v7
		v29 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v28 - v29
L2:
	;
	goto L1
L3:
	;
	v13 = v3
	v14 = v4
	goto L4
L4:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v18 == int32(0) {
		v28 = v18
		v29 = v17
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v28 = v18
	v29 = v17
	goto L2
L6:
	;
	v21 = int32(1)
	if v18 == v17 {
		v13 = v13 + v21
		v14 = v14 + v21
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
func F_cmpcmdflag(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 == int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v45
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v7 == v8 {
		v45 = int32(0)
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if base.B2i32(v19 == int32(0))|base.B2i32(v19 != v22) != 0 {
		v40 = v19
		v41 = v22
		goto L10
	} else {
		goto L11
	}
L5:
	;
	if base.Ui32(v8) < base.Ui32(v7) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v13 = int32(1)
	goto L8
L7:
	;
	v13 = int32(-1)
	goto L8
L8:
	;
	return v13
L9:
	;
	v45 = v40 - v41
	goto L1
L10:
	;
	goto L9
L11:
	;
	v25 = v15
	v26 = v16
	goto L12
L12:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v30 == int32(0) {
		v40 = v30
		v41 = v29
		goto L10
	} else {
		goto L14
	}
L13:
	;
	v40 = v30
	v41 = v29
	goto L10
L14:
	;
	v33 = int32(1)
	if v30 == v29 {
		v25 = v25 + v33
		v26 = v26 + v33
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
func F_cmpspellaffix(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	if base.B2i32(v9 == int32(0))|base.B2i32(v9 != v12) != 0 {
		v30 = v9
		v31 = v12
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v30 - v31
L2:
	;
	goto L1
L3:
	;
	v15 = v4
	v16 = v6
	goto L4
L4:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	if v20 == int32(0) {
		v30 = v20
		v31 = v19
		goto L2
	} else {
		goto L6
	}
L5:
	;
	v30 = v20
	v31 = v19
	goto L2
L6:
	;
	v23 = int32(1)
	if v20 == v19 {
		v15 = v15 + v23
		v16 = v16 + v23
		goto L4
	} else {
		goto L7
	}
L7:
	;
	goto L5
}
