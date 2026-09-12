package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_uuid_extract_version(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(*(*int8)(unsafe.Add(mBase, uint32(v3)+8)))
	if int32(-64) <= v4 {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v7)
		return int32(0)
	} else {
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+6)))
		return int32(base.Ui32(v11) >> (uint(int32(4)) % 32))
	}
}
func F_uuid_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(16)
	goto L4
L1:
	;
	return int32(base.Ui32(v66^int32(-1)) >> (uint(int32(31)) % 32))
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (v2|v3)&int32(3) != 0 {
		v35 = v2
		v36 = v3
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = v2
	v13 = v3
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
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
func F_uuid_lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = int32(16)
	goto L4
L1:
	;
	return int32(base.Ui32(v66) >> (uint(int32(31)) % 32))
L2:
	;
	v66 = int32(0)
	goto L1
L3:
	;
	v40 = v35
	v41 = v36
	v42 = v37
	goto L13
L4:
	;
	if (v2|v3)&int32(3) != 0 {
		v35 = v2
		v36 = v3
		v37 = v4
		goto L3
	} else {
		goto L7
	}
L6:
	;
	if v25 == int32(0) {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	v12 = v2
	v13 = v3
	v14 = v4
	goto L8
L8:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v17 != v18 {
		v35 = v12
		v36 = v13
		v37 = v14
		goto L3
	} else {
		goto L10
	}
L9:
	;
	goto L6
L10:
	;
	v20 = int32(4)
	v21 = v13 + v20
	v23 = v12 + v20
	v25 = v14 - v20
	if base.Ui32(int32(3)) < base.Ui32(v25) {
		v12 = v23
		v13 = v21
		v14 = v25
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v35 = v23
	v36 = v21
	v37 = v25
	goto L3
L13:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v45 == v46 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v66 = v45 - v46
	goto L1
L15:
	;
	v48 = int32(1)
	v53 = v42 - v48
	if v53 != 0 {
		v40 = v40 + v48
		v41 = v41 + v48
		v42 = v53
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
func F_uuid_ns_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int64
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v3 = m.G0
	v5 = v3 - int32(48)
	m.G0 = v5
	v7 = int32(577506)
	v8 = *(*int64)(unsafe.Add(mBase, _consts[1127]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+29)) = v8
	v10 = *(*int64)(unsafe.Add(mBase, _consts[1128]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = v10
	v12 = *(*int64)(unsafe.Add(mBase, _consts[1129]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+16)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, _consts[1130]))
	*(*int64)(unsafe.Add(mBase, uint32(v5))) = v14
	v16 = *(*int64)(unsafe.Add(mBase, _consts[1131]))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+8)) = v16
	v20 = F_DirectFunctionCall1Coll(m, int32(3392), int32(0), v5)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(48)
		return v20
	}
}
func F_uuid_skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v17 int64
	_ = v17
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_palloc(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = F_palloc(m, int32(16))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = v13
			*(*int64)(unsafe.Add(mBase, uint32(v6))) = v13
			v17 = int64(-1)
			*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v17
			*(*int64)(unsafe.Add(mBase, uint32(v11))) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(1546)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(1547)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v11
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v6
			return int32(0)
		}
	}
}
func F_uuid_unparse(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+6)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+9)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+11)))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+14)))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+15)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v21)+48)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v32
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v23
	v57 = F_snprintf(m, l1, int32(37), int32(30560), v21)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		return
	} else {
		m.G0 = v21 - int32(-64)
		return
	}
}
