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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	v3 = int32(0)
	if l0 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = F_palloc(m, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = F_palloc(m, v18<<(uint(int32(1))%32))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return int32(0)
L5:
	;
	return v13
L6:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v23 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L4
	} else {
		goto L22
	}
L8:
	;
	v28 = v3
	goto L11
L9:
	;
	goto L10
L10:
	;
	return v21
L11:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v37 <= int32(0) {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v28<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v51 = int32(0)
	goto L15
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v46+v51<<(uint(int32(2))%32))))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if v45 != v61 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v66 = int32(1)
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v21+v28<<(uint(v66)%32)))) = uint16(v69)
	v72 = v28 + v66
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v72 < v73 {
		v28 = v72
		goto L11
	} else {
		goto L21
	}
L17:
	;
	v64 = v51 + int32(1)
	if v64 != v37 {
		v51 = v64
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L7
L21:
	;
	goto L12
L22:
	;
	F_errmsg_internal(m, int32(_a_F_extract_grouping_cols_0), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_extract_grouping_cols_1), int32(366), int32(_a_F_extract_grouping_cols_2))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
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
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
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
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
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
	var v158 int32
	_ = v158
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
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
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v211 = F_makeGroupingSet(m, v203, v202, v210)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L16
	} else {
		goto L43
	}
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v202 = v8
	v203 = v23
	goto L1
L3:
	;
	goto L4
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(0) < v24 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = v8
	v40 = v8
	goto L8
L6:
	;
	v158 = v8
	goto L7
L7:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.B2i32(v158 == int32(0))|base.B2i32(v168 != int32(3)) != 0 {
		v202 = v158
		v203 = v168
		goto L1
	} else {
		goto L36
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v42+v40<<(uint(int32(2))%32))))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v47 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	v158 = v145
	goto L7
L10:
	;
	v145 = F_lappend(m, v34, v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L16
	} else {
		goto L34
	}
L11:
	;
	v114 = int32(0)
	v116 = F_transformGroupClauseExpr(m, l0, v114, l1, v46, l3, l4, l5, l6, v114)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L16
	} else {
		goto L31
	}
L12:
	;
	if v47 != int32(107) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v56 = int32(0)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v56 < v59 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v52 = F_transformGroupingSet(m, l0, l1, v46, l3, l4, l5, l6)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	return int32(0)
L17:
	;
	v144 = v52
	goto L10
L18:
	;
	v69 = v56
	v70 = v56
	v74 = v56
	goto L21
L19:
	;
	v102 = v56
	goto L20
L20:
	;
	v111 = F_exprLocation(m, v46)
	mBase = m.M
	v112 = F_makeGroupingSet(m, int32(1), v102, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L16
	} else {
		goto L30
	}
L21:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v77+v70<<(uint(int32(2))%32))))
	v83 = F_transformGroupClauseExpr(m, l0, v74, l1, v81, l3, l4, l5, l6, int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v102 = v89
	goto L20
L23:
	;
	if v83 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v85 = F_bms_add_member(m, v74, v83)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L16
	} else {
		goto L27
	}
L25:
	;
	v89 = v69
	v90 = v74
	goto L26
L26:
	;
	v92 = v70 + int32(1)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v92 < v93 {
		v69 = v89
		v70 = v92
		v74 = v90
		goto L21
	} else {
		goto L29
	}
L27:
	;
	v87 = F_lappend_int(m, v69, v83)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L16
	} else {
		goto L28
	}
L28:
	;
	v89 = v87
	v90 = v85
	goto L26
L29:
	;
	goto L22
L30:
	;
	v144 = v112
	goto L10
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v116
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v116
	v124 = F_list_make1_impl(m, int32(471), v18+int32(8))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L16
	} else {
		goto L32
	}
L32:
	;
	v126 = F_exprLocation(m, v46)
	mBase = m.M
	v127 = F_makeGroupingSet(m, int32(1), v124, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L16
	} else {
		goto L33
	}
L33:
	;
	v144 = v127
	goto L10
L34:
	;
	v148 = v40 + int32(1)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v148 < v149 {
		v34 = v145
		v40 = v148
		goto L8
	} else {
		goto L35
	}
L35:
	;
	goto L9
L36:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	if v173 < int32(13) {
		v202 = v158
		v203 = int32(3)
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L16
	} else {
		goto L38
	}
L38:
	;
	F_errcode(m, int32(17039621))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L16
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(_a_F_transformGroupingSet_0), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	F_parser_errposition(m, l1, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L16
	} else {
		goto L41
	}
L41:
	;
	F_errfinish(m, int32(_a_F_transformGroupingSet_1), int32(2587), int32(_a_F_transformGroupingSet_2))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L16
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
	m.G0 = v18 + int32(16)
	return v211
}
