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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v5 != v6 {
		if base.Ui32(v5) < base.Ui32(v6) {
			v11 = int32(-1)
		} else {
			v11 = int32(1)
		}
		return v11
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
		if v13 != v14 {
			if base.Ui32(v13) < base.Ui32(v14) {
				v19 = int32(-1)
			} else {
				v19 = int32(1)
			}
			return v19
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+2)))
			if v22 != v23 {
				if base.Ui32(v22) < base.Ui32(v23) {
					v28 = int32(-1)
				} else {
					v28 = int32(1)
				}
				v29 = v28
			} else {
				v29 = int32(0)
			}
			return v29
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
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v8 == int32(0) {
		v27 = v7
		v28 = v8
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v28 - v27
L2:
	;
	goto L1
L3:
	;
	if v7 != v8 {
		v27 = v7
		v28 = v8
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v12 = v3
	v13 = v4
	goto L5
L5:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v17 == int32(0) {
		v27 = v16
		v28 = v17
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v27 = v16
	v28 = v17
	goto L2
L7:
	;
	v20 = int32(1)
	if v16 == v17 {
		v12 = v12 + v20
		v13 = v13 + v20
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 == int32(2) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v44
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v7 == v8 {
		v44 = int32(0)
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
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
	if v20 == int32(0) {
		v39 = v19
		v40 = v20
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
	v44 = v40 - v39
	goto L1
L10:
	;
	goto L9
L11:
	;
	if v19 != v20 {
		v39 = v19
		v40 = v20
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v24 = v15
	v25 = v16
	goto L13
L13:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v28
		v40 = v29
		goto L10
	} else {
		goto L15
	}
L14:
	;
	v39 = v28
	v40 = v29
	goto L10
L15:
	;
	v32 = int32(1)
	if v28 == v29 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
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
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
	if v10 == int32(0) {
		v29 = v9
		v30 = v10
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v30 - v29
L2:
	;
	goto L1
L3:
	;
	if v9 != v10 {
		v29 = v9
		v30 = v10
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = v4
	v15 = v6
	goto L5
L5:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v19 == int32(0) {
		v29 = v18
		v30 = v19
		goto L2
	} else {
		goto L7
	}
L6:
	;
	v29 = v18
	v30 = v19
	goto L2
L7:
	;
	v22 = int32(1)
	if v18 == v19 {
		v14 = v14 + v22
		v15 = v15 + v22
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
}
