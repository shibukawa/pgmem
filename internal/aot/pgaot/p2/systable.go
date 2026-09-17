package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_systable_beginscan_ordered(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v109 int64
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[0]))
	if v21 != v19 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[1]))
	v25 = F_list_member_ptr(m, v24, v19)
	mBase = m.M
	v27 = v25
	goto L4
L3:
	;
	v27 = int32(1)
	goto L4
L4:
	;
	goto L1
L5:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[2])))
	if v31 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L10
	} else {
		goto L49
	}
L8:
	;
	v57 = F_palloc(m, int32(24))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L15
	}
L9:
	;
	v36 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	if v36 == int32(0) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v42 + int32(4)
	F_errmsg_internal(m, int32(_a_F_systable_beginscan_ordered_0), v17+int32(16))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	F_errfinish(m, int32(_a_F_systable_beginscan_ordered_1), int32(668), int32(_a_F_systable_beginscan_ordered_2))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v57))) = l0
	v62 = F_table_slot_create(m, l0, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v62
	if l2 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v68 = F_GetCatalogSnapshot(m, v67)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	v72 = l2
	v73 = v6
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v73
	v77 = F_palloc(m, l3*int32(48))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L10
	} else {
		goto L22
	}
L20:
	;
	v70 = F_RegisterSnapshot(m, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v72 = v70
	v73 = v70
	goto L19
L22:
	;
	if l3 <= int32(0) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v206 = int32(0)
	v208 = F_index_beginscan(m, l0, l1, v72, v206, l3, v206)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L10
	} else {
		goto L43
	}
L24:
	;
	v93 = v6
	goto L25
L25:
	;
	v96 = v93 * int32(48)
	v97 = v77 + v96
	v98 = l4 + v96
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v98)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+40)) = v99
	v101 = *(*int64)(unsafe.Add(mBase, uint32(v98)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+32)) = v101
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v98)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+24)) = v103
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v98)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+16)) = v105
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v98)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v97)+8)) = v107
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v98)))
	*(*int64)(unsafe.Add(mBase, uint32(v97))) = v109
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v112 = int32(*(*int16)(unsafe.Add(mBase, uint32(v111)+8)))
	if v112 <= int32(0) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L10
	} else {
		goto L40
	}
L27:
	;
	goto L26
L28:
	;
	if v152 == v158 {
		goto L27
	} else {
		goto L38
	}
L29:
	;
	v152 = int32(0)
	v158 = v112
	goto L28
L30:
	;
	goto L31
L31:
	;
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v98)+4)))
	v125 = int32(0)
	goto L32
L32:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111+int32(48)+v125<<(uint(int32(1))%32)))))
	if v137 == v119 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L27
L34:
	;
	v140 = v125 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v97)+4)) = uint16(v140)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v143 = int32(*(*int16)(unsafe.Add(mBase, uint32(v142)+8)))
	v152 = v125
	v158 = v143
	goto L28
L35:
	;
	goto L36
L36:
	;
	v145 = v125 + int32(1)
	if v145 != v112 {
		v125 = v145
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	v163 = v93 + int32(1)
	if l3 != v163 {
		v93 = v163
		goto L25
	} else {
		goto L39
	}
L39:
	;
	goto L23
L40:
	;
	F_errmsg_internal(m, int32(_a_F_systable_beginscan_ordered_3), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_systable_beginscan_ordered_1), int32(707), int32(_a_F_systable_beginscan_ordered_2))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L10
	} else {
		goto L42
	}
L42:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v208
	v211 = int32(0)
	F_index_rescan(m, v208, v77, l3, v211, v211)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = int32(0)
	F_pfree(m, v77)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L10
	} else {
		goto L45
	}
L45:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[3]))
	if v220 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v222 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_systable_beginscan_ordered[4])) = uint8(v222)
	goto L48
L47:
	;
	goto L48
L48:
	;
	m.G0 = v17 + int32(32)
	return v57
L49:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L10
	} else {
		goto L50
	}
L50:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v235 + int32(4)
	F_errmsg(m, int32(_a_F_systable_beginscan_ordered_4), v17)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L10
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_systable_beginscan_ordered_1), int32(664), int32(_a_F_systable_beginscan_ordered_2))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L10
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
