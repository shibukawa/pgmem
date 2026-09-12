package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bms_make_singleton(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	if l0 < int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(440676), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494441), int32(223), int32(245851))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v23 = int32(base.Ui32(l0) >> (uint(int32(5)) % 32))
		v25 = v23 + int32(1)
		v30 = F_palloc0(m, v25<<(uint(int32(2))%32)+int32(8))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v25
			*(*int32)(unsafe.Add(mBase, uint32(v30))) = int32(445)
			*(*int32)(unsafe.Add(mBase, uint32(v30+v23<<(uint(int32(2))%32))+8)) = int32(1) << (uint(l0) % 32)
			return v30
		}
	}
}
func F_bms_nonempty_difference(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v11 = int32(1)
	if l1 == int32(0) {
		v49 = v11
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return v49
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v15 < v14 {
		v49 = v11
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v17 = int32(1)
	if v14 <= v17 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v20 = v17
	goto L9
L8:
	;
	v20 = v14
	goto L9
L9:
	;
	v21 = int32(8)
	v26 = int32(0)
	goto L10
L10:
	;
	v33 = v26 << (uint(int32(2)) % 32)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0+v21+v33)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33+(l1+v21))))
	v40 = v35 & (v37 ^ int32(-1))
	v42 = base.B2i32(v40 != int32(0))
	if v40 != 0 {
		v49 = v42
		goto L4
	} else {
		goto L12
	}
L11:
	;
	v49 = v42
	goto L4
L12:
	;
	v44 = v26 + int32(1)
	if v44 != v20 {
		v26 = v44
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
}
func F_bms_overlap_list(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v3 = int32(0)
	if l0 == v3 {
		v53 = v3
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L15
	}
L2:
	;
	return v53
L3:
	;
	if l1 == int32(0) {
		v53 = v3
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v13 <= int32(0) {
		v53 = v3
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v21 = int32(0)
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v18+v21<<(uint(int32(2))%32))))
	if v31 < int32(0) {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v53 = int32(0)
	goto L2
L8:
	;
	v35 = int32(base.Ui32(v31) >> (uint(int32(5)) % 32))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v35 < v36 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(1)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8)+v35<<(uint(int32(2))%32))))
	if int32(base.Ui32(v42)>>(uint(v31)%32))&v38 != 0 {
		v53 = v38
		goto L2
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v48 = v21 + int32(1)
	if v48 != v13 {
		v21 = v48
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	goto L7
L14:
	;
	return int32(0)
L15:
	;
	F_errmsg_internal(m, int32(440676), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(494441), int32(624), int32(74224))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
