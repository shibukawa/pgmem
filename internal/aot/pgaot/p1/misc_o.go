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
	var v241 int32
	_ = v241
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
		goto L56
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
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_OpernameGetCandidates[0]))
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_OpernameGetCandidates[1]))
	v67 = v41
	v77 = v4
	v78 = v4
	goto L17
L17:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v47+int32(48)+v77<<(uint(int32(2))%32))))
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
	v249 = v77 + int32(1)
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v47)+40))
	if v249 < v250 {
		v67 = v230
		v77 = v249
		v78 = v241
		goto L17
	} else {
		goto L55
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
	v241 = v78
	goto L19
L23:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v207)+4))
	if v225 < v114 {
		goto L52
	} else {
		goto L53
	}
L24:
	;
	v190 = v78 + v54
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
	v241 = v78 + int32(40)
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
		v230 = v67
		v241 = v78
		goto L19
	} else {
		goto L29
	}
L28:
	;
	v230 = v67
	v241 = v78
	goto L19
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v101 <= int32(0) {
		v230 = v67
		v241 = v78
		goto L19
	} else {
		goto L30
	}
L30:
	;
	v104 = int32(0)
	if v104 < v101 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v107 = v101
	goto L33
L32:
	;
	v107 = v104
	goto L33
L33:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v91)+68))
	v114 = int32(0)
	goto L35
L34:
	;
	if v67 == int32(0) {
		v174 = v114
		goto L24
	} else {
		goto L39
	}
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v108+v114<<(uint(int32(2))%32))))
	if base.B2i32(v110 != v62)&base.B2i32(v133 == v110) != 0 {
		goto L34
	} else {
		goto L37
	}
L36:
	;
	v230 = v67
	v241 = v78
	goto L19
L37:
	;
	v137 = v114 + int32(1)
	if v137 != v107 {
		v114 = v137
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v91)+80))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+37)))
	if v142 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v67)+32))
	if v141 != v143 {
		v174 = v114
		goto L24
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v148 = v67
	goto L45
L43:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v91)+84))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v67)+36))
	if v145 != v146 {
		v174 = v114
		goto L24
	} else {
		goto L44
	}
L44:
	;
	v207 = v67
	goto L23
L45:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
	if v166 == v141 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v174 = v114
	goto L24
L47:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v91)+84))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v148)+36))
	if v168 == v169 {
		v207 = v148
		goto L23
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	if v171 != 0 {
		v148 = v171
		goto L45
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	goto L46
L52:
	;
	v230 = v67
	v241 = v78
	goto L19
L53:
	;
	goto L54
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v207)+4)) = v114
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	*(*int32)(unsafe.Add(mBase, uint32(v207)+8)) = v228
	v230 = v67
	v241 = v78
	goto L19
L55:
	;
	goto L18
L56:
	;
	v272 = v252
	goto L3
}
func F_oauth_get_mechanisms(m *base.Module, l0 int32, l1 int32) {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	F_appendStringInfoString(m, l1, int32(_a_F_oauth_get_mechanisms_0))
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
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v200 int32
	_ = v200
	v3 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = v3
	goto L3
L1:
	;
	m.G0 = v14 + int32(16)
	return v200
L2:
	;
	if v26 == int32(0) {
		v200 = v16
		goto L1
	} else {
		goto L9
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21*int32(40))+uint32(_c_F_obtain_object_name_namespace[0])))
	v26 = base.B2i32(v25 == v17)
	if v26 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L2
L5:
	;
	v30 = v21 + int32(1)
	if v30 != int32(37) {
		v21 = v30
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
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v38 = F_table_open(m, v36, int32(1))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	F_relation_close(m, v38, int32(1))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L12
	} else {
		goto L61
	}
L11:
	;
	F_relation_close(m, v38, int32(1))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L12
	} else {
		goto L60
	}
L12:
	;
	return int32(0)
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = F_get_object_attnum_oid(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = F_get_catalog_object_by_oid(m, v38, v43, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	if v46 == int32(0) {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = F_get_object_attnum_namespace(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L12
	} else {
		goto L18
	}
L17:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = m.G0
	v83 = v81 - int32(16)
	m.G0 = v83
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_obtain_object_name_namespace[1]))
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L18:
	;
	if v51 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v58 = F_heap_getattr_2(m, v46, v51, v55, v14+int32(15))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v60 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_obtain_object_name_namespace[2]))
	goto L23
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)) = uint8(v74)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v75
	goto L17
L23:
	;
	if base.B2i32(v63 != int32(0))&base.B2i32(v58 == v63) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v74 = int32(1)
	v75 = int32(_a_F_obtain_object_name_namespace_0)
	goto L22
L25:
	;
	goto L26
L26:
	;
	v70 = F_isAnyTempNamespace(m, v58)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v70 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v72 = F_get_namespace_name(m, v58)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v74 = v3
	v75 = v72
	goto L22
L30:
	;
	if v139 == int32(0) {
		goto L11
	} else {
		goto L53
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L12
	} else {
		goto L50
	}
L32:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+36)))
	m.G0 = v83 + int32(16)
	goto L30
L33:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v87 == v80 {
		v131 = v86
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v97 = v3
	goto L41
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_obtain_object_name_namespace[1])) = v126
	v131 = v126
	goto L32
L38:
	;
	v126 = v102 + int32(_a_F_obtain_object_name_namespace_1)
	goto L37
L39:
	;
	v126 = v102 + int32(_a_F_obtain_object_name_namespace_2)
	goto L37
L40:
	;
	v126 = v102 + int32(_a_F_obtain_object_name_namespace_3)
	goto L37
L41:
	;
	v102 = v97 * int32(40)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_obtain_object_name_namespace[0])))
	if v80 != v103 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v126 = v102 + int32(_a_F_obtain_object_name_namespace_4)
	goto L37
L43:
	;
	if v97 == int32(36) {
		goto L31
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_obtain_object_name_namespace[3])))
	if v109 == v80 {
		goto L38
	} else {
		goto L47
	}
L47:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_obtain_object_name_namespace[4])))
	if v111 == v80 {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v102)+uint32(_c_F_obtain_object_name_namespace[5])))
	if v113 == v80 {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	v97 = v97 + int32(4)
	goto L41
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v80
	F_errmsg_internal(m, int32(_a_F_obtain_object_name_namespace_5), v83)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_obtain_object_name_namespace_6), int32(2777), int32(_a_F_obtain_object_name_namespace_7))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v158 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v160 = F_get_object_attnum_name(m, v159)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	if v160 == int32(0) {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v38)+52))
	v167 = F_heap_getattr_2(m, v46, v160, v164, v14+int32(15))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	if v169 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	v170 = F_pstrdup(m, v167)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v170
	goto L11
L60:
	;
	v200 = v16
	goto L1
L61:
	;
	v200 = int32(0)
	goto L1
}
func F_oidvectoreqfast(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_DirectFunctionCall2Coll(m, int32(1571), int32(0), l0, l1)
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	v2 = int32(0)
	v5 = F_open(m, l0, int32(_a_F_opendir_0), v2)
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
				F___memset(m, v30, int32(0), v18)
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
