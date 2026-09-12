package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_AtEOSubXact_HashTables(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	v15 = v13 - int32(1)
	if int32(0) <= v15 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = v15
	goto L4
L2:
	;
	goto L3
L3:
	;
	m.G0 = v10 + int32(16)
	return
L4:
	;
	v26 = v21 << (uint(int32(2)) % 32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1189])))
	if l1 <= v29 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	if l0 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	if int32(0) < v21 {
		v21 = v21 - int32(1)
		goto L4
	} else {
		goto L16
	}
L9:
	;
	v53 = int32(4514064)
	v54 = *(*int32)(unsafe.Add(mBase, _consts[1187]))
	v56 = v54 - int32(1)
	v58 = v56 << (uint(int32(2)) % 32)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v58)+uint32(_consts[1188])))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1188]))) = v61
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)+uint32(_consts[1189])))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1189]))) = v65
	*(*int32)(unsafe.Add(mBase, _consts[1187])) = v56
	goto L8
L10:
	;
	v35 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	if v35 == int32(0) {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v26)+uint32(_consts[1188])))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v41
	F_errmsg_internal(m, int32(239686), v10)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	F_errfinish(m, int32(499076), int32(1956), int32(167588))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	goto L5
}
func F_AtEOSubXact_Namespace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, _consts[319]))
	if l1 == v5 {
		if l0 != 0 {
			*(*int32)(unsafe.Add(mBase, _consts[319])) = l2
			return
		} else {
			v10 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[320])) = uint8(v10)
			v13 = int32(0)
			*(*int32)(unsafe.Add(mBase, _consts[316])) = v13
			*(*int32)(unsafe.Add(mBase, _consts[319])) = v13
			*(*int32)(unsafe.Add(mBase, _consts[318])) = v13
			*(*uint8)(unsafe.Add(mBase, _consts[321])) = uint8(v13)
			v25 = *(*int32)(unsafe.Add(mBase, _consts[124]))
			*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = v13
			return
		}
	} else {
		return
	}
}
func F_AtEOSubXact_PgStat(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v68 int64
	_ = v68
	var v71 int32
	_ = v71
	var v72 int64
	_ = v72
	var v73 int64
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v93 int64
	_ = v93
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v109 int64
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v254 int32
	_ = v254
	v3 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, _consts[204]))
	if v12 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v15 < l1 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, _consts[204])) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	if v20 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = l1 - int32(1)
	v29 = v20
	goto L7
L5:
	;
	goto L6
L6:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v132 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L7:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+68))
	if l0 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L6
L9:
	;
	if v34 != 0 {
		v29 = v34
		goto L7
	} else {
		goto L31
	}
L10:
	;
	F_pfree(m, v29)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L24
	} else {
		goto L30
	}
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	if v35 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+24)))
	if v85 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L14:
	;
	v79 = F_pgstat_get_xact_stack_level(m, v22)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L24
	} else {
		goto L25
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+56))
	if v38 != v22 {
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+24)))
	if v40 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v77
	goto L10
L18:
	;
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)))
	if v43 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v62 + v63
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v66)+8)) = v67 + v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v72 = *(*int64)(unsafe.Add(mBase, uint32(v71)+16))
	v73 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v71)+16)) = v72 + v73
	goto L17
L21:
	;
	v53 = v35
	goto L23
L22:
	;
	v44 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+24)) = uint8(v44)
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v35)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+48)) = v46
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v35)))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+32)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v35)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v35)+40)) = v50
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v53 = v52
	goto L23
L23:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v53))) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v56)+8)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v59)+16)) = v60
	goto L17
L24:
	;
	return
L25:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v79)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(v79)+20)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v22
	goto L9
L26:
	;
	v96 = *(*int64)(unsafe.Add(mBase, uint32(v33)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+40)) = v96 + v95
	v99 = *(*int64)(unsafe.Add(mBase, uint32(v33)+48))
	v100 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+48)) = v99 + v100
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v33)+56))
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v29)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+56)) = v103 + v104
	v107 = *(*int64)(unsafe.Add(mBase, uint32(v33)+96))
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v29)+8))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v33)+96)) = v107 + (v108 + v109)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v29)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v33)+8)) = v113
	goto L10
L27:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	v95 = v88
	goto L26
L28:
	;
	goto L29
L29:
	;
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v29)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v89
	v91 = *(*int64)(unsafe.Add(mBase, uint32(v29)+40))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v91
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v29)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v93
	v95 = v89
	goto L26
L30:
	;
	goto L9
L31:
	;
	goto L8
L32:
	;
	F_pfree(m, v12)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L24
	} else {
		goto L59
	}
L33:
	;
	v136 = l1 - int32(1)
	v138 = *(*int32)(unsafe.Add(mBase, _consts[204]))
	if v138 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v162 == int32(0) {
		goto L32
	} else {
		goto L40
	}
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v139 == v136 {
		v161 = v138
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	v144 = F_MemoryContextAlloc(m, v142, int32(24))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L24
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v146 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v144)+16)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v136
	v150 = v144 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v144)+8)) = v150
	v153 = int32(4503504)
	v154 = *(*int32)(unsafe.Add(mBase, _consts[204]))
	*(*int32)(unsafe.Add(mBase, uint32(v144)+20)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(v144)+4)) = v154
	*(*int32)(unsafe.Add(mBase, _consts[204])) = v144
	v161 = v144
	goto L34
L40:
	;
	v166 = v12 + int32(8)
	if v162 == v166 {
		goto L32
	} else {
		goto L41
	}
L41:
	;
	v169 = v161 + int32(8)
	v172 = v162
	v178 = v3
	goto L42
L42:
	;
	v182 = *(*int64)(unsafe.Add(mBase, uint32(v172-int32(12))))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v172)))
	*(*int32)(unsafe.Add(mBase, uint32(v184))) = v186
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v188 - int32(1)
	if l0 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	if v232 <= int32(0) {
		goto L32
	} else {
		goto L57
	}
L44:
	;
	if v184 != v166 {
		v172 = v184
		v178 = v232
		goto L42
	} else {
		goto L56
	}
L45:
	;
	F_pfree(m, v195)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L24
	} else {
		goto L55
	}
L46:
	;
	v195 = v172 - int32(20)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v172-int32(4)))))
	if v198 != int32(1) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	if v212 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v172-int32(16))))
	v205 = F_pgstat_drop_entry(m, v201, v204, v182)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L24
	} else {
		goto L50
	}
L50:
	;
	F_pfree(m, v195)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L24
	} else {
		goto L51
	}
L51:
	;
	v232 = v178 + (v205 ^ int32(1))
	goto L44
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v161)+12)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v161)+8)) = v169
	goto L54
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v172)+4)) = v169
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v172))) = v220
	*(*int32)(unsafe.Add(mBase, uint32(v220)+4)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(v161)+8)) = v172
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v161)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+16)) = v224 + int32(1)
	v232 = v178
	goto L44
L55:
	;
	v232 = v178
	goto L44
L56:
	;
	goto L43
L57:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[938]))
	v239 = *(*int64)(unsafe.Add(mBase, uint32(v238)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v238)+16)) = v239 + int64(1)
	goto L58
L58:
	;
	goto L32
L59:
	;
	goto L1
}
