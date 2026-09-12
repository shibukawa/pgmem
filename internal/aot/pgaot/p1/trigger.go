package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecGetTriggerNewSlot(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	if v5 == int32(0) {
		v8 = int32(4470400)
		v9 = *(*int32)(unsafe.Add(mBase, _consts[3]))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
		*(*int32)(unsafe.Add(mBase, _consts[3])) = v12
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
		v15 = F_table_slot_callbacks(m, v10)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			v19 = F_ExecInitExtraTupleSlot(m, l0, v14, v15)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+76)) = v19
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v9
				v24 = v19
				return v24
			}
		}
	} else {
		v24 = v5
		return v24
	}
}
func F_ExecGetTriggerResultRel(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	v4 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v8 == v4 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v149
L2:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v45 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v11 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v14 = int32(0)
	if v14 < v11 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v17 = v11
	goto L7
L6:
	;
	v17 = v14
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v24 = v4
	goto L8
L8:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v18+v24<<(uint(int32(2))%32))))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+56))
	if l1 == v31 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L2
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+200))
	if v33 == l2 {
		v149 = v29
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v36 = v24 + int32(1)
	if v36 != v17 {
		v24 = v36
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	goto L9
L15:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v83 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v48 <= int32(0) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v51 = int32(0)
	if v51 < v48 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v55 = v48
	goto L20
L19:
	;
	v55 = v51
	goto L20
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v62 = v51
	goto L21
L21:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v56+v62<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+8))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+56))
	if l1 == v69 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L15
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v67)+200))
	if v71 == l2 {
		v149 = v67
		goto L1
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v74 = v62 + int32(1)
	if v74 != v55 {
		v62 = v74
		goto L21
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	goto L22
L28:
	;
	v122 = F_table_open(m, l1, int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L41
	} else {
		goto L42
	}
L29:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v86 <= int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v89 = int32(0)
	if v89 < v86 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v93 = v86
	goto L33
L32:
	;
	v93 = v89
	goto L33
L33:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v100 = v89
	goto L34
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v94+v100<<(uint(int32(2))%32))))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+56))
	if l1 == v107 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L28
L36:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105)+200))
	if v109 == l2 {
		v149 = v105
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v112 = v100 + int32(1)
	if v112 != v93 {
		v100 = v112
		goto L34
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	goto L35
L41:
	;
	return int32(0)
L42:
	;
	v126 = int32(4470400)
	v127 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v129
	v132 = F_palloc0(m, int32(216))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = int32(388)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	F_InitResultRelInfo(m, v132, v122, int32(0), l2, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	v141 = F_lappend(m, v140, v132)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L41
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+84)) = v141
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v127
	v149 = v132
	goto L1
}
func F_FreeTriggerDesc(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v6 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v10 = v5
	v12 = int32(0)
	goto L7
L5:
	;
	v75 = v5
	goto L6
L6:
	;
	F_pfree(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L9
	} else {
		goto L36
	}
L7:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	F_pfree(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = v70
	goto L6
L9:
	;
	return
L10:
	;
	v16 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+36)))
	if int32(0) < v16 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
	F_pfree(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L9
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v22 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+34)))
	if int32(0) < v22 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	goto L13
L15:
	;
	v26 = v22 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+34)) = uint16(v26)
	v30 = v26
	goto L18
L16:
	;
	goto L17
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	if v55 != 0 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+v30&int32(65535)<<(uint(int32(2))%32))))
	F_pfree(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L9
	} else {
		goto L20
	}
L19:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v10)+44))
	F_pfree(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L9
	} else {
		goto L22
	}
L20:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+34)))
	v43 = v41 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+34)) = uint16(v43)
	v45 = base.I32_extend16_s(v43)
	if int32(0) <= v45 {
		v30 = v45
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	goto L17
L23:
	;
	F_pfree(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	if v58 != 0 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	goto L25
L27:
	;
	F_pfree(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	if v61 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	F_pfree(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L9
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v67 = v12 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v67 < v68 {
		v10 = v10 + int32(60)
		v12 = v67
		goto L7
	} else {
		goto L35
	}
L34:
	;
	goto L33
L35:
	;
	goto L8
L36:
	;
	F_pfree(m, l0)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L9
	} else {
		goto L37
	}
L37:
	;
	goto L3
}
func F_trigger_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		F_errcode(m, int32(1088))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(220933)
			F_errmsg(m, int32(189401), v4)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(485272), int32(366), int32(66084))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
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
