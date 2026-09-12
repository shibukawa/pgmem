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
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
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
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
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
	if v46 < int32(0) {
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
	v137 = m.ExcPending
	if v137 != 0 {
		goto L4
	} else {
		goto L41
	}
L15:
	;
	m.G0 = v13 + int32(16)
	return v15
L16:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	m.T0[v117].(func(*base.Module, int32))(m, v15)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L40
	}
L17:
	;
	if v30 <= v46 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v51 = int32(-1)
	if v19 == v51 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v55 = v51
	goto L21
L20:
	;
	v55 = int32(1)
	goto L21
L21:
	;
	v57 = v46
	goto L22
L22:
	;
	v68 = v28 + v57*int32(6)
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+12)) = uint16(v69)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v71
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v73 == int32(1) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L16
L24:
	;
	F_table_tuple_get_latest_tid(m, v31, v13+int32(8))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	if v81 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L26
L28:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _consts[26])))
	if v83&int32(1) == int32(0) {
		goto L14
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v16)+188))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+60))
	v92 = m.T0[v91].(func(*base.Module, int32, int32, int32, int32) int32)(m, v16, v13+int32(8), v18, v15)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	if v92 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v95 = v94 + v55
	*(*int32)(unsafe.Add(mBase, uint32(l0)+128)) = v95
	v98 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v98 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v102 = v95
	goto L36
L36:
	;
	if v102 < int32(0) {
		goto L16
	} else {
		goto L38
	}
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v102 = v101
	goto L36
L38:
	;
	if v102 < v30 {
		v57 = v102
		goto L22
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	goto L15
L41:
	;
	F_errmsg_internal(m, int32(324035), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(314654), int32(1264), int32(261766))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_TidRecheck(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
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
	var v23 int32
	_ = v23
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+120)))
	if v5 != 0 {
		v23 = int32(1)
		return v23
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
		if v8 != 0 {
			v14 = v8
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
			v18 = F_bsearch(m, l1+int32(28), v14, v15, int32(6), int32(769))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				v23 = base.B2i32(v18 != int32(0))
				return v23
			}
		} else {
			F_TidListEval(m, l0)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
				v14 = v13
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
				v18 = F_bsearch(m, l1+int32(28), v14, v15, int32(6), int32(769))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v23 = base.B2i32(v18 != int32(0))
					return v23
				}
			}
		}
	}
}
