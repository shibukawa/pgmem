package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpernameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	v4 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	F_DeconstructQualifiedName(m, l0, v21+int32(12), v21+int32(8))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v31 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v21 + int32(16)
	return v272
L4:
	;
	v41 = int32(0)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v47 = F_SearchSysCacheList(m, int32(39), int32(1), v44, v41, v41)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L13
	}
L5:
	;
	v32 = F_LookupExplicitNamespace(m, v31, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L11
	}
L8:
	;
	if l2 == int32(0) {
		v40 = v32
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v32 != 0 {
		v40 = v32
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v272 = int32(0)
	goto L3
L11:
	;
	v40 = v4
	goto L4
L12:
	;
	F_ReleaseCatCacheList(m, v47)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L60
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	if v49 <= int32(0) {
		v252 = v41
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v54 = F_palloc(m, v49*int32(40))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	if v56 <= int32(0) {
		v252 = v41
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	v64 = *(*int32)(unsafe.Add(mBase, _consts[317]))
	v67 = v41
	v77 = v4
	v78 = v4
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(48)+v78<<(uint(int32(2))%32))))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+56))
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+22)))
	v91 = v89 + v90
	if l1 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v252 = v230
	goto L12
L19:
	;
	v249 = v78 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	if v249 < v250 {
		v67 = v230
		v77 = v240
		v78 = v249
		goto L17
	} else {
		goto L59
	}
L20:
	;
	if v40 != 0 {
		goto L25
	} else {
		goto L26
	}
L21:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+76)))
	if v94 == l1&int32(255) {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v230 = v67
	v240 = v77
	goto L19
L23:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if v225 < v114 {
		goto L56
	} else {
		goto L57
	}
L24:
	;
	v190 = v77 + v54
	*(*int32)(unsafe.Add(mBase, uint32(v190)+4)) = v174
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+28)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v190)+20)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v190)+12)) = int64(8589934594)
	*(*int32)(unsafe.Add(mBase, uint32(v190)+8)) = v192
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v190)+32)) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v91)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v190)+36)) = v202
	v230 = v190
	v240 = v77 + int32(40)
	goto L19
L25:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	if v97 == v40 {
		v174 = int32(0)
		goto L24
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v64 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v230 = v67
	v240 = v77
	goto L19
L29:
	;
	v230 = v67
	v240 = v77
	goto L19
L30:
	;
	goto L31
L31:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v101 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v230 = v67
	v240 = v77
	goto L19
L33:
	;
	goto L34
L34:
	;
	v104 = int32(0)
	if v104 < v101 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v107 = v101
	goto L37
L36:
	;
	v107 = v104
	goto L37
L37:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	v114 = int32(0)
	goto L39
L38:
	;
	if v67 == int32(0) {
		v174 = v114
		goto L24
	} else {
		goto L43
	}
L39:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v108+v114<<(uint(int32(2))%32))))
	if base.B2i32(v110 != v62)&base.B2i32(v133 == v110) != 0 {
		goto L38
	} else {
		goto L41
	}
L40:
	;
	v230 = v67
	v240 = v77
	goto L19
L41:
	;
	v137 = v114 + int32(1)
	if v137 != v107 {
		v114 = v137
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+37)))
	if v142 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
	if v141 != v143 {
		v174 = v114
		goto L24
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v148 = v67
	goto L49
L47:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v91)+84))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v67)+36))
	if v145 != v146 {
		v174 = v114
		goto L24
	} else {
		goto L48
	}
L48:
	;
	v207 = v67
	goto L23
L49:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
	if v166 == v141 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v174 = v114
	goto L24
L51:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v91)+84))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v148)+36))
	if v168 == v169 {
		v207 = v148
		goto L23
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v171 != 0 {
		v148 = v171
		goto L49
	} else {
		goto L55
	}
L54:
	;
	goto L53
L55:
	;
	goto L50
L56:
	;
	v230 = v67
	v240 = v77
	goto L19
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v114
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v228
	v230 = v67
	v240 = v77
	goto L19
L59:
	;
	goto L18
L60:
	;
	v272 = v252
	goto L3
}
func F_oauth_get_mechanisms(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	F_appendStringInfoString(m, l1, int32(549475))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		F_appendStringInfoChar(m, l1, int32(0))
		v8 = m.ExcPending
		if v8 != 0 {
			return
		} else {
			return
		}
	}
}
func F_obtain_object_name_namespace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	v15 = int32(1)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v20 = v3
	goto L3
L1:
	;
	m.G0 = v13 + int32(16)
	return v216
L2:
	;
	if v27 == int32(0) {
		v216 = v15
		goto L1
	} else {
		goto L9
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20*int32(40))+uint32(_consts[326])))
	v27 = base.B2i32(v26 == v16)
	if v27 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L2
L5:
	;
	v31 = v20 + int32(1)
	if v31 != int32(37) {
		v20 = v31
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	goto L4
L8:
	;
	goto L7
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = F_table_open(m, v37, int32(1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_sequence_close(m, v39, int32(1))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L12
	} else {
		goto L64
	}
L11:
	;
	F_sequence_close(m, v39, int32(1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L63
	}
L12:
	;
	return int32(0)
L13:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v44 = F_get_object_attnum_oid(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v47 = F_get_catalog_object_by_oid(m, v39, v44, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if v47 == int32(0) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = F_get_object_attnum_namespace(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v83 = m.G0
	v85 = v83 - int32(16)
	m.G0 = v85
	v88 = *(*int32)(unsafe.Add(mBase, _consts[325]))
	if v88 != 0 {
		goto L33
	} else {
		goto L34
	}
L18:
	;
	if v52 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v59 = F_heap_getattr_2(m, v47, v52, v56, v13+int32(15))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v61 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _consts[316]))
	goto L23
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)) = uint8(v75)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v76
	goto L17
L23:
	;
	if base.B2i32(v64 != int32(0))&base.B2i32(v59 == v64) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v75 = int32(1)
	v76 = int32(246608)
	goto L22
L25:
	;
	goto L26
L26:
	;
	v71 = F_isAnyTempNamespace(m, v59)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v71 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v73 = F_get_namespace_name(m, v59)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v75 = v3
	v76 = v73
	goto L22
L30:
	;
	if v157 == int32(0) {
		goto L11
	} else {
		goto L56
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L12
	} else {
		goto L53
	}
L32:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+36)))
	m.G0 = v85 + int32(16)
	goto L30
L33:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v89 == v81 {
		v149 = v88
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v95 = int32(0)
	goto L38
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _consts[325])) = v145
	v149 = v145
	goto L32
L38:
	;
	v103 = v95 * int32(40)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v103)+uint32(_consts[326])))
	if v81 != v106 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v145 = v103 + int32(798656)
	goto L37
L40:
	;
	if v95 == int32(36) {
		goto L31
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	v113 = (v95 | int32(1)) * int32(40)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v113)+uint32(_consts[326])))
	if v81 == v116 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v145 = v113 + int32(798656)
	goto L37
L45:
	;
	goto L46
L46:
	;
	v123 = (v95 | int32(2)) * int32(40)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v123)+uint32(_consts[326])))
	if v81 == v126 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v145 = v123 + int32(798656)
	goto L37
L48:
	;
	goto L49
L49:
	;
	v133 = (v95 | int32(3)) * int32(40)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v133)+uint32(_consts[326])))
	if v81 == v136 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v145 = v133 + int32(798656)
	goto L37
L51:
	;
	v95 = v95 + int32(4)
	goto L38
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = v81
	F_errmsg_internal(m, int32(64596), v85)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(515466), int32(2777), int32(527606))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L56:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v176 != 0 {
		goto L11
	} else {
		goto L57
	}
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v178 = F_get_object_attnum_name(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L12
	} else {
		goto L58
	}
L58:
	;
	if v178 == int32(0) {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v185 = F_heap_getattr_2(m, v47, v178, v182, v13+int32(15))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L12
	} else {
		goto L60
	}
L60:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+15)))
	if v187 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	v188 = F_pstrdup(m, v185)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v188
	goto L11
L63:
	;
	v216 = v15
	goto L1
L64:
	;
	v216 = int32(0)
	goto L1
}
func F_oidvectoreqfast(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_DirectFunctionCall2Coll(m, int32(1587), int32(0), l0, l1)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v5 != int32(0))
	}
}
func F_oidvectorge(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_oidvectorlt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2) >> (uint(int32(31)) % 32))
	}
}
func F_opendir(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v2 = int32(0)
	v5 = F_open(m, l0, int32(589824), v2)
	mBase = m.M
	if v2 <= v5 {
		v18 = base.I32_wrap_i64(base.I64_extend_i32_u(int32(1)) * base.I64_extend_i32_u(int32(2072)))
		v30 = F_emscripten_builtin_malloc(m, v18)
		mBase = m.M
		if v30 == int32(0) {
		} else {
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30-int32(4)))))
			if v35&int32(3) == int32(0) {
			} else {
				v41 = F___memset(m, v30, int32(0), v18)
				mBase = m.M
			}
		}
		if v30 == int32(0) {
			v44 = m.Wasi_snapshot_preview1.Fd_close(m, v5)
			mBase = m.M
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v30)+8)) = v5
			v48 = v30
			return v48
		}
	} else {
		v48 = v2
		return v48
	}
}
func F_ordered_set_shutdown(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 != 0 {
		F_tuplesort_end(m, v3)
		mBase = m.M
		v5 = m.ExcPending
		if v5 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
			if v9 != 0 {
				v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
				v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
				m.T0[v11].(func(*base.Module, int32))(m, v9)
				mBase = m.M
				v13 = m.ExcPending
				if v13 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+20))
		if v9 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
			m.T0[v11].(func(*base.Module, int32))(m, v9)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
