package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_OpernameGetCandidates(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v235 int32
	_ = v235
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(16)
	m.G0 = v23
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	F_DeconstructQualifiedName(m, l0, v23+int32(12), v23+int32(8))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v23 + int32(16)
	return v306
L4:
	;
	v53 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23)+8)))
	v54 = int64(0)
	v56 = F_SearchSysCacheList(m, int32(39), int32(1), v53, v54, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L12
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v36 | int32(1)
	v40 = F_LookupExplicitNamespace(m, v35, l2)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	if v40 == int32(0) {
		v306 = v5
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v44 | int32(2)
	v50 = v40
	goto L4
L10:
	;
	v50 = v5
	goto L4
L11:
	;
	F_ReleaseCatCacheList(m, v56)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L55
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v56)+56))
	if v58 <= int32(0) {
		v284 = v5
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v63 = F_palloc(m, v58*int32(40))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v56)+56))
	if v65 <= int32(0) {
		v284 = v5
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _c_F_OpernameGetCandidates[0]))
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_OpernameGetCandidates[1]))
	v80 = v5
	v87 = v5
	v88 = v5
	goto L16
L16:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v56-int32(-64)+v87<<(uint(int32(2))%32))))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+72))
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+22)))
	v102 = v100 + v101
	if l1 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v284 = v260
	goto L11
L18:
	;
	v277 = v87 + int32(1)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v56)+56))
	if v277 < v278 {
		v80 = v260
		v87 = v277
		v88 = v268
		goto L16
	} else {
		goto L54
	}
L19:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v107 | int32(4)
	if v50 != 0 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+76)))
	if v105 == l1&int32(255) {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v260 = v80
	v268 = v88
	goto L18
L22:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v251 < v127 {
		goto L51
	} else {
		goto L52
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v107 | int32(12)
	v214 = v88 + v63
	*(*int32)(unsafe.Add(mBase, uint32(v214)+4)) = v191
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+28)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v214)+20)) = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v214)+12)) = int64(8589934594)
	*(*int32)(unsafe.Add(mBase, uint32(v214)+8)) = v216
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v102)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v214)+32)) = v224
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v102)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v214)+36)) = v226
	v260 = v214
	v268 = v88 + int32(40)
	goto L18
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v102)+68))
	if v112 == v50 {
		v191 = int32(0)
		goto L23
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	if v73 == int32(0) {
		v260 = v80
		v268 = v88
		goto L18
	} else {
		goto L28
	}
L27:
	;
	v260 = v80
	v268 = v88
	goto L18
L28:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v116 <= int32(0) {
		v260 = v80
		v268 = v88
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v119 = int32(0)
	if v119 < v116 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v122 = v116
	goto L32
L31:
	;
	v122 = v119
	goto L32
L32:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v73)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v102)+68))
	v127 = int32(0)
	goto L34
L33:
	;
	if v80 == int32(0) {
		v191 = v127
		goto L23
	} else {
		goto L38
	}
L34:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v123+v127<<(uint(int32(2))%32))))
	if base.B2i32(v125 != v71)&base.B2i32(v150 == v125) != 0 {
		goto L33
	} else {
		goto L36
	}
L35:
	;
	v260 = v80
	v268 = v88
	goto L18
L36:
	;
	v154 = v127 + int32(1)
	if v154 != v122 {
		v127 = v154
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v102)+80))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+53)))
	if v159 != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v80)+32))
	if v158 != v160 {
		v191 = v127
		goto L23
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v169 = v80
	goto L44
L42:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v102)+84))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v80)+36))
	if v162 != v163 {
		v191 = v127
		goto L23
	} else {
		goto L43
	}
L43:
	;
	v235 = v80
	goto L22
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v169)+32))
	if v185 == v158 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v191 = v127
	goto L23
L46:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v102)+84))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v169)+36))
	if v187 == v188 {
		v235 = v169
		goto L22
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	if v190 != 0 {
		v169 = v190
		goto L44
	} else {
		goto L50
	}
L49:
	;
	goto L48
L50:
	;
	goto L45
L51:
	;
	v260 = v80
	v268 = v88
	goto L18
L52:
	;
	goto L53
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v235)+4)) = v127
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	*(*int32)(unsafe.Add(mBase, uint32(v235)+8)) = v254
	v260 = v80
	v268 = v88
	goto L18
L54:
	;
	goto L17
L55:
	;
	v306 = v284
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
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
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
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
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
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v172 int64
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v207 int32
	_ = v207
	v3 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	v17 = int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = v3
	goto L3
L1:
	;
	m.G0 = v15 + int32(16)
	return v207
L2:
	;
	if v27 == int32(0) {
		v207 = v17
		goto L1
	} else {
		goto L9
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22*int32(40))+uint32(_c_F_obtain_object_name_namespace[0])))
	v27 = base.B2i32(v26 == v18)
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
	v31 = v22 + int32(1)
	if v31 != int32(37) {
		v22 = v31
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
	F_relation_close(m, v39, int32(1))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L12
	} else {
		goto L61
	}
L11:
	;
	F_relation_close(m, v39, int32(1))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L12
	} else {
		goto L60
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
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = m.G0
	v86 = v84 - int32(16)
	m.G0 = v86
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_obtain_object_name_namespace[1]))
	if v89 != 0 {
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
	v59 = F_heap_getattr_2(m, v47, v52, v56, v15+int32(15))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v61 != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	v62 = base.I32_wrap_i64(v59)
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_obtain_object_name_namespace[2]))
	goto L23
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+38)) = uint8(v76)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v77
	goto L17
L23:
	;
	if base.B2i32(v65 != int32(0))&base.B2i32(v62 == v65) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v76 = int32(1)
	v77 = int32(_a_F_obtain_object_name_namespace_0)
	goto L22
L25:
	;
	goto L26
L26:
	;
	v72 = F_isAnyTempNamespace(m, v62)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L27
	}
L27:
	;
	if v72 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	v74 = F_get_namespace_name(m, v62)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L12
	} else {
		goto L29
	}
L29:
	;
	v76 = v3
	v77 = v74
	goto L22
L30:
	;
	if v144 == int32(0) {
		goto L11
	} else {
		goto L53
	}
L31:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L12
	} else {
		goto L50
	}
L32:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+36)))
	m.G0 = v86 + int32(16)
	goto L30
L33:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v90 == v83 {
		v135 = v89
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v100 = v3
	goto L41
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_obtain_object_name_namespace[1])) = v130
	v135 = v130
	goto L32
L38:
	;
	v130 = v106 + int32(_a_F_obtain_object_name_namespace_1)
	goto L37
L39:
	;
	v130 = v106 + int32(_a_F_obtain_object_name_namespace_2)
	goto L37
L40:
	;
	v130 = v106 + int32(_a_F_obtain_object_name_namespace_3)
	goto L37
L41:
	;
	v106 = v100 * int32(40)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+uint32(_c_F_obtain_object_name_namespace[0])))
	if v83 != v107 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v130 = v106 + int32(_a_F_obtain_object_name_namespace_4)
	goto L37
L43:
	;
	if v100 == int32(36) {
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
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v106)+uint32(_c_F_obtain_object_name_namespace[3])))
	if v113 == v83 {
		goto L38
	} else {
		goto L47
	}
L47:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v106)+uint32(_c_F_obtain_object_name_namespace[4])))
	if v115 == v83 {
		goto L39
	} else {
		goto L48
	}
L48:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v106)+uint32(_c_F_obtain_object_name_namespace[5])))
	if v117 == v83 {
		goto L40
	} else {
		goto L49
	}
L49:
	;
	v100 = v100 + int32(4)
	goto L41
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v83
	F_errmsg_internal(m, int32(_a_F_obtain_object_name_namespace_5), v86)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L12
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_obtain_object_name_namespace_6), int32(2827), int32(_a_F_obtain_object_name_namespace_7))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
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
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v163 != 0 {
		goto L11
	} else {
		goto L54
	}
L54:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v165 = F_get_object_attnum_name(m, v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L12
	} else {
		goto L55
	}
L55:
	;
	if v165 == int32(0) {
		goto L11
	} else {
		goto L56
	}
L56:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v172 = F_heap_getattr_2(m, v47, v165, v169, v15+int32(15))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L12
	} else {
		goto L57
	}
L57:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+15)))
	if v174 != 0 {
		goto L11
	} else {
		goto L58
	}
L58:
	;
	v176 = F_pstrdup(m, base.I32_wrap_i64(v172))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L12
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v176
	goto L11
L60:
	;
	v207 = v17
	goto L1
L61:
	;
	v207 = int32(0)
	goto L1
}
func F_oid8in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = F_uint64in_subr(m, v2, int32(_a_F_oid8in_0), v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
	}
}
func F_oidvectoreqfast(m *base.Module, l0 int64, l1 int64) int32 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_DirectFunctionCall2Coll(m, int32(1784), int32(0), l0, l1)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v5 != int64(0))
	}
}
func F_oidvectorge(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return int64(base.Ui64(v2^int64(-1))>>(uint(int64(31))%64)) & int64(1)
	}
}
func F_oidvectorlt(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_btoidvectorcmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return int64(base.Ui64(v2)>>(uint(int64(31))%64)) & int64(1)
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
func F_ordered_set_shutdown(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v4 = base.I32_wrap_i64(l0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
	if v5 != 0 {
		F_tuplesort_end(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(0)
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
			if v11 != 0 {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				m.T0[v13].(func(*base.Module, int32))(m, v11)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(0)
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		if v11 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
			m.T0[v13].(func(*base.Module, int32))(m, v11)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
