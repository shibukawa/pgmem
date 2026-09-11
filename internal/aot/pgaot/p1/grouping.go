package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extract_grouping_cols(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	v3 = int32(0)
	if l0 == v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v99
L2:
	;
	v13 = F_palloc(m, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = F_palloc(m, v17<<(uint(int32(1))%32))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L5
	} else {
		goto L7
	}
L5:
	;
	return int32(0)
L6:
	;
	v99 = v13
	goto L1
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v22 <= int32(0) {
		v99 = v20
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v27 = v3
	goto L9
L9:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L5
	} else {
		goto L21
	}
L11:
	;
	goto L10
L12:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v36 <= int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39+v27<<(uint(int32(2))%32))))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v51 = int32(0)
	goto L14
L14:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v45+v51<<(uint(int32(2))%32))))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+16))
	if v44 != v60 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v65 = int32(1)
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v59)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20+v27<<(uint(v65)%32)))) = uint16(v68)
	v71 = v27 + v65
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v71 < v72 {
		v27 = v71
		goto L9
	} else {
		goto L20
	}
L16:
	;
	v63 = v51 + int32(1)
	if v63 != v36 {
		v51 = v63
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	goto L15
L19:
	;
	goto L11
L20:
	;
	v99 = v20
	goto L1
L21:
	;
	F_errmsg_internal(m, int32(69285), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(462424), int32(366), int32(359558))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_transformGroupingSet(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v20 == v8 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v210 = F_makeGroupingSet(m, v201, v202, v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L22
	} else {
		goto L45
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v201 = v23
	v202 = v8
	goto L1
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v24 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v166 != int32(3) {
		v201 = v166
		v202 = v159
		goto L1
	} else {
		goto L37
	}
L6:
	;
	v159 = v8
	goto L5
L7:
	;
	goto L8
L8:
	;
	v35 = v8
	v39 = v8
	goto L9
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v39<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v47 != int32(107) {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	v159 = v145
	goto L5
L11:
	;
	v145 = F_lappend(m, v35, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L22
	} else {
		goto L35
	}
L12:
	;
	v114 = int32(0)
	v116 = F_transformGroupClauseExpr(m, l0, v114, l1, v46, l3, l4, l5, l6, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L22
	} else {
		goto L32
	}
L13:
	;
	if v47 != int32(1) {
		goto L12
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v112 = F_transformGroupingSet(m, l0, l1, v46, l3, l4, l5, l6)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L22
	} else {
		goto L31
	}
L16:
	;
	v52 = int32(0)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v52 < v55 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v65 = v52
	v66 = v52
	v71 = v52
	goto L20
L18:
	;
	v101 = v52
	goto L19
L19:
	;
	v109 = F_exprLocation(m, v46)
	mBase = m.M
	v110 = F_makeGroupingSet(m, int32(1), v101, v109)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L22
	} else {
		goto L30
	}
L20:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v65<<(uint(int32(2))%32))))
	v79 = F_transformGroupClauseExpr(m, l0, v71, l1, v77, l3, l4, l5, l6, int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v101 = v87
	goto L19
L22:
	;
	return int32(0)
L23:
	;
	if v79 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v83 = F_bms_add_member(m, v71, v79)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L22
	} else {
		goto L27
	}
L25:
	;
	v87 = v66
	v88 = v71
	goto L26
L26:
	;
	v90 = v65 + int32(1)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v90 < v91 {
		v65 = v90
		v66 = v87
		v71 = v88
		goto L20
	} else {
		goto L29
	}
L27:
	;
	v85 = F_lappend_int(m, v66, v79)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v87 = v85
	v88 = v83
	goto L26
L29:
	;
	goto L21
L30:
	;
	v144 = v110
	goto L11
L31:
	;
	v144 = v112
	goto L11
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v116
	v124 = F_list_make1_impl(m, int32(471), v18+int32(8))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L22
	} else {
		goto L33
	}
L33:
	;
	v126 = F_exprLocation(m, v46)
	mBase = m.M
	v127 = F_makeGroupingSet(m, int32(1), v124, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L22
	} else {
		goto L34
	}
L34:
	;
	v144 = v127
	goto L11
L35:
	;
	v148 = v39 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v148 < v149 {
		v35 = v145
		v39 = v148
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L10
L37:
	;
	if v159 == int32(0) {
		v201 = v166
		v202 = v159
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
	if v172 < int32(13) {
		v201 = int32(3)
		v202 = v159
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L22
	} else {
		goto L40
	}
L40:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L22
	} else {
		goto L41
	}
L41:
	;
	F_errmsg(m, int32(113952), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L22
	} else {
		goto L42
	}
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_parser_errposition(m, l1, v186)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L22
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(468017), int32(2587), int32(99766))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L22
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L45:
	;
	m.G0 = v18 + int32(16)
	return v210
}
