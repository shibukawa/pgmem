package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datumEstimateSpace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	if l1 != 0 {
		v8 = int32(4)
	} else {
		v8 = int32(8)
	}
	if l1|l2 != 0 {
		v33 = v8
		return v33
	} else {
		if l3 != int32(-1) {
			v29 = F_datumGetSize(m, l0, int32(0), l3)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v33 = v29 + int32(4)
				return v33
			}
		} else {
			v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v12 != int32(1) {
				v29 = F_datumGetSize(m, l0, int32(0), l3)
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v33 = v29 + int32(4)
					return v33
				}
			} else {
				v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
				if v15&int32(254) != int32(2) {
					v29 = F_datumGetSize(m, l0, int32(0), l3)
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						v33 = v29 + int32(4)
						return v33
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+2))
					v21 = F_EOH_get_flat_size(m, v20)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v21 + int32(4)
					}
				}
			}
		}
	}
}
func F_datumIsEqual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	if l2 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return base.B2i32(l0 == l1)
L2:
	;
	goto L3
L3:
	;
	v8 = F_datumGetSize(m, l0, int32(0), l3)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v13 = F_datumGetSize(m, l1, int32(0), l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	if v8 == v13 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v8) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v79 = int32(1)
	goto L9
L9:
	;
	return base.B2i32(v79 == int32(0))
L10:
	;
	v79 = v77
	goto L9
L11:
	;
	v77 = int32(0)
	goto L10
L12:
	;
	v51 = v46
	v52 = v47
	v53 = v48
	goto L22
L13:
	;
	if (l0|l1)&int32(3) != 0 {
		v46 = l0
		v47 = l1
		v48 = v8
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v39 = l0
	v40 = l1
	v41 = v8
	goto L15
L15:
	;
	if v41 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L16:
	;
	v23 = l0
	v24 = l1
	v25 = v8
	goto L17
L17:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v28 != v29 {
		v46 = v23
		v47 = v24
		v48 = v25
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v39 = v34
	v40 = v32
	v41 = v36
	goto L15
L19:
	;
	v31 = int32(4)
	v32 = v24 + v31
	v34 = v23 + v31
	v36 = v25 - v31
	if base.Ui32(int32(3)) < base.Ui32(v36) {
		v23 = v34
		v24 = v32
		v25 = v36
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v46 = v39
	v47 = v40
	v48 = v41
	goto L12
L22:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 == v57 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v77 = v56 - v57
	goto L10
L24:
	;
	v59 = int32(1)
	v64 = v53 - v59
	if v64 != 0 {
		v51 = v51 + v59
		v52 = v52 + v59
		v53 = v64
		goto L22
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	goto L23
L27:
	;
	goto L11
}
