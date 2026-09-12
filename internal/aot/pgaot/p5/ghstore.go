package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ghstore_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(360406)
			F_errmsg(m, int32(190934), v5)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(487768), int32(112), int32(66823))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
func F_ghstore_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v78 int32
	_ = v78
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 == v2 {
		v27 = v2
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v27&int32(1) != 0 {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	goto L1
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
	if v14 == int32(0) {
		v27 = v2
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	if v17 != int32(7) {
		v27 = v2
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v20 != int32(17) {
		v27 = v2
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+24)))
	v27 = v23 ^ int32(1)
	goto L2
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v31 = F_get_fn_opclass_options(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v36 = int32(16)
	goto L9
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v38 = int32(4)
	v39 = v37 & v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+4)))
	if v40&v38 != 0 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	return int32(0)
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v36 = v35
	goto L9
L12:
	;
	return v6
L13:
	;
	v78 = int32(base.Ui32(v39) >> (uint(int32(2)) % 32))
	goto L15
L14:
	;
	if v39 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v78)
	goto L12
L16:
	;
	v78 = int32(0)
	goto L15
L17:
	;
	v45 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v6))) = uint8(v45)
	v47 = int32(0)
	if v36 <= v47 {
		goto L12
	} else {
		goto L18
	}
L18:
	;
	v50 = int32(8)
	v54 = v47
	goto L19
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+(v8+v50)))))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+(v7+v50)))))
	if v60 != v62 {
		goto L16
	} else {
		goto L21
	}
L20:
	;
	goto L12
L21:
	;
	v65 = v54 + int32(1)
	if v36 != v65 {
		v54 = v65
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
}
