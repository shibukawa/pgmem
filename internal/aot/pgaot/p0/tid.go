package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_TidNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
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
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	if v20 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_TidListEval(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v28 = v20
	goto L3
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v19 == int32(-1) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v28 = v27
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v46
	if base.B2i32(v46 < int32(0))|base.B2i32(v30 <= v46) != 0 {
		goto L16
	} else {
		goto L17
	}
L7:
	;
	if v29 < int32(0) {
		v46 = v30 - int32(1)
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v29 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v46 = v29 - int32(1)
	goto L6
L11:
	;
	v43 = int32(-1)
	goto L13
L12:
	;
	v43 = v29
	goto L13
L13:
	;
	v46 = v43 + int32(1)
	goto L6
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L40
	}
L15:
	;
	m.G0 = v13 + int32(16)
	return v15
L16:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	m.T0[v118].(func(*base.Module, int32))(m, v15)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L39
	}
L17:
	;
	v52 = int32(-1)
	if v19 == v52 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v56 = v52
	goto L20
L19:
	;
	v56 = int32(1)
	goto L20
L20:
	;
	v58 = v46
	goto L21
L21:
	;
	v69 = v28 + v58*int32(6)
	v70 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v70)
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v72
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v74 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L16
L23:
	;
	F_table_tuple_get_latest_tid(m, v31, v13+int32(8))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_TidNext[0]))
	if v82 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	v84 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TidNext[1])))
	if v84&int32(1) == int32(0) {
		goto L14
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v91)+60))
	v93 = m.T0[v92].(func(*base.Module, int32, int32, int32, int32) int32)(m, v16, v13+int32(8), v18, v15)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	if v93 != 0 {
		goto L15
	} else {
		goto L32
	}
L32:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v96 = v95 + v56
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v96
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_TidNext[2]))
	if v99 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L36
	}
L34:
	;
	v103 = v96
	goto L35
L35:
	;
	if v103 < int32(0) {
		goto L16
	} else {
		goto L37
	}
L36:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v103 = v102
	goto L35
L37:
	;
	if v103 < v30 {
		v58 = v103
		goto L21
	} else {
		goto L38
	}
L38:
	;
	goto L22
L39:
	;
	goto L15
L40:
	;
	F_errmsg_internal(m, int32(_a_F_TidNext_0), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_TidNext_1), int32(1264), int32(_a_F_TidNext_2))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_TidRecheck(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v3 != 0 {
		v22 = int32(1)
		return v22
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v7 != 0 {
			v13 = v7
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
			v17 = F_bsearch(m, l1+int32(28), v13, v14, int32(6), int32(770))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				v22 = base.B2i32(v17 != int32(0))
				return v22
			}
		} else {
			F_TidListEval(m, l0)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v13 = v12
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
				v17 = F_bsearch(m, l1+int32(28), v13, v14, int32(6), int32(770))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v22 = base.B2i32(v17 != int32(0))
					return v22
				}
			}
		}
	}
}
func F_TidStoreGetDSA(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	return v2
}
