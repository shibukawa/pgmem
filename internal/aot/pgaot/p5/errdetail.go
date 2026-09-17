package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_errdetail_busy_db(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v10 = base.B2i32(l0 <= v3)
	if v10|base.B2i32(l1 <= v3) == v3 {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
		F_errdetail(m, int32(_a_F_errdetail_busy_db_0), v7)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			m.G0 = v7 + int32(48)
			return
		}
	} else {
		if v10 == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
			F_errdetail_plural(m, int32(_a_F_errdetail_busy_db_1), int32(_a_F_errdetail_busy_db_2), l0, v7+int32(16))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				m.G0 = v7 + int32(48)
				return
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = l1
			F_errdetail_plural(m, int32(_a_F_errdetail_busy_db_3), int32(_a_F_errdetail_busy_db_4), l1, v7+int32(32))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				m.G0 = v7 + int32(48)
				return
			}
		}
	}
}
func F_errdetail_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(_a_F_errdetail_plural_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[0])) = v15 + int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[1]))
	if int32(0) <= v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = int32(_a_F_errdetail_plural_1)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[2]))
	v27 = v20 * int32(100)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errdetail_plural[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[2])) = v30
	v33 = v11 + int32(16)
	F_initStringInfo(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L24
	}
L4:
	;
	return
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errdetail_plural[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[5])) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	if l2 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v42 = l0
	goto L8
L7:
	;
	v42 = l1
	goto L8
L8:
	;
	v43 = F_appendStringInfoVA(m, v33, v42, l3)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = v43
	goto L13
L11:
	;
	goto L12
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errdetail_plural[6])))
	if v71 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v54 = v11 + int32(16)
	F_enlargeStringInfo(m, v54, v47)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errdetail_plural[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[5])) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	v61 = F_appendStringInfoVA(m, v54, v42, l3)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v61 != 0 {
		v47 = v61
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	F_pfree(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v75 = F_pstrdup(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errdetail_plural[6]))) = v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_pfree(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[2])) = v24
	v83 = int32(_a_F_errdetail_plural_0)
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errdetail_plural[0])) = v85 - int32(1)
	m.G0 = v11 + int32(32)
	return
L24:
	;
	F_errmsg_internal(m, int32(_a_F_errdetail_plural_2), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_errdetail_plural_3), int32(1303), int32(_a_F_errdetail_plural_4))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
