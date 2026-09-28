package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_datumEstimateSpace(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	if l1 != 0 {
		v8 = int32(4)
	} else {
		v8 = int32(12)
	}
	if l1|l2 != 0 {
		v37 = v8
		return v37
	} else {
		if l3 != int32(-1) {
			v32 = F_datumGetSize(m, l0, int32(0), l3)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v37 = v32 + int32(4)
				return v37
			}
		} else {
			v12 = base.I32_wrap_i64(l0)
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v13 != int32(1) {
				v32 = F_datumGetSize(m, l0, int32(0), l3)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v37 = v32 + int32(4)
					return v37
				}
			} else {
				v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
				if v16&int32(254) != int32(2) {
					v32 = F_datumGetSize(m, l0, int32(0), l3)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v37 = v32 + int32(4)
						return v37
					}
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(l0))+2))
					v23 = F_EOH_get_flat_size(m, v22)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return v23 + int32(4)
					}
				}
			}
		}
	}
}
func F_datumIsEqual(m *base.Module, l0 int64, l1 int64, l2 int32, l3 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
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
	v16 = base.I32_wrap_i64(l0)
	v17 = base.I32_wrap_i64(l1)
	if base.Ui32(int32(4)) <= base.Ui32(v8) {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	v81 = int32(1)
	goto L9
L9:
	;
	return base.B2i32(v81 == int32(0))
L10:
	;
	v81 = v79
	goto L9
L11:
	;
	v79 = int32(0)
	goto L10
L12:
	;
	v53 = v48
	v54 = v49
	v55 = v50
	goto L22
L13:
	;
	if (v16|v17)&int32(3) != 0 {
		v48 = v16
		v49 = v17
		v50 = v8
		goto L12
	} else {
		goto L16
	}
L14:
	;
	v41 = v16
	v42 = v17
	v43 = v8
	goto L15
L15:
	;
	if v43 == int32(0) {
		goto L11
	} else {
		goto L21
	}
L16:
	;
	v25 = v16
	v26 = v17
	v27 = v8
	goto L17
L17:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v30 != v31 {
		v48 = v25
		v49 = v26
		v50 = v27
		goto L12
	} else {
		goto L19
	}
L18:
	;
	v41 = v36
	v42 = v34
	v43 = v38
	goto L15
L19:
	;
	v33 = int32(4)
	v34 = v26 + v33
	v36 = v25 + v33
	v38 = v27 - v33
	if base.Ui32(int32(3)) < base.Ui32(v38) {
		v25 = v36
		v26 = v34
		v27 = v38
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v48 = v41
	v49 = v42
	v50 = v43
	goto L12
L22:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54))))
	if v58 == v59 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v79 = v58 - v59
	goto L10
L24:
	;
	v61 = int32(1)
	v66 = v55 - v61
	if v66 != 0 {
		v53 = v53 + v61
		v54 = v54 + v61
		v55 = v66
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
