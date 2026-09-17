package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ChooseRelationName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	v15 = m.G0
	v17 = v15 - int32(256)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+120)) = int32(4)
	v23 = F_table_open(m, int32(1259), int32(1))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v28 = v17 + int32(192)
	goto L6
L3:
	;
	v163 = int32(0)
	goto L34
L4:
	;
	v145 = F_strlen(m, v134)
	mBase = m.M
	goto L3
L6:
	;
	goto L7
L7:
	;
	v35 = int32(63)
	if (v28^l2)&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v138 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v138)
	goto L4
L9:
	;
	v119 = v114
	v120 = v115
	v121 = v116
	goto L30
L10:
	;
	if v109 == int32(0) {
		v134 = v107
		v135 = v108
		goto L8
	} else {
		goto L29
	}
L11:
	;
	v107 = l2
	v108 = v28
	v109 = v35
	goto L10
L12:
	;
	goto L13
L13:
	;
	v39 = int32(0)
	if base.B2i32(l2&int32(3) == v39)|int32(0) == v39 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v75 == int32(0) {
		v134 = v72
		v135 = v73
		goto L8
	} else {
		goto L23
	}
L15:
	;
	v51 = l2
	v52 = v28
	v53 = v35
	goto L18
L16:
	;
	goto L17
L17:
	;
	v72 = l2
	v73 = v28
	v74 = v35
	v75 = int32(1)
	goto L14
L18:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	*(*uint8)(unsafe.Add(mBase, uint32(v52))) = uint8(v55)
	if v55 == int32(0) {
		v114 = v51
		v115 = v52
		v116 = v53
		goto L9
	} else {
		goto L20
	}
L19:
	;
	v72 = v66
	v73 = v60
	v74 = v62
	v75 = v64
	goto L14
L20:
	;
	v59 = int32(1)
	v60 = v52 + v59
	v62 = v53 - v59
	v63 = int32(0)
	v64 = base.B2i32(v62 != v63)
	v66 = v51 + v59
	if v66&int32(3) == v63 {
		v72 = v66
		v73 = v60
		v74 = v62
		v75 = v64
		goto L14
	} else {
		goto L21
	}
L21:
	;
	if v62 != 0 {
		v51 = v66
		v52 = v60
		v53 = v62
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if base.B2i32(v78 == int32(0))|base.B2i32(base.Ui32(v74) < base.Ui32(int32(4))) != 0 {
		v107 = v72
		v108 = v73
		v109 = v74
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v85 = v72
	v86 = v73
	v87 = v74
	goto L25
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v93 = int32(-2139062144)
	if (int32(16843008)-v90|v90)&v93 != v93 {
		v114 = v85
		v115 = v86
		v116 = v87
		goto L9
	} else {
		goto L27
	}
L26:
	;
	v107 = v101
	v108 = v99
	v109 = v103
	goto L10
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = v90
	v98 = int32(4)
	v99 = v86 + v98
	v101 = v85 + v98
	v103 = v87 - v98
	if base.Ui32(int32(3)) < base.Ui32(v103) {
		v85 = v101
		v86 = v99
		v87 = v103
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v114 = v107
	v115 = v108
	v116 = v109
	goto L9
L30:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119))))
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v123)
	if v123 == int32(0) {
		v134 = v119
		v135 = v120
		goto L8
	} else {
		goto L32
	}
L31:
	;
	v134 = v130
	v135 = v128
	goto L8
L32:
	;
	v127 = int32(1)
	v128 = v120 + v127
	v130 = v119 + v127
	v132 = v121 - v127
	if v132 != 0 {
		v119 = v130
		v120 = v128
		v121 = v132
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v165 = v17 + int32(16)
	v171 = F_makeObjectName(m, l0, l1, v17+int32(192))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	F_relation_close(m, v23, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L1
	} else {
		goto L57
	}
L36:
	;
	goto L35
L37:
	;
	F_ScanKeyInit(m, v165, int32(2), int32(3), int32(62), v171)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	v175 = int32(3)
	F_ScanKeyInit(m, v17-int32(-64), v175, v175, int32(184), l3)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v185 = F_systable_beginscan(m, v23, int32(2663), int32(1), v17+int32(120), int32(2), v165)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	v187 = F_systable_getnext(m, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	F_systable_endscan(m, v185)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	if v187 == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	if l4 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_pfree(m, v171)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L55
	}
L46:
	;
	v195 = m.G0
	v197 = v195 - int32(96)
	m.G0 = v197
	v201 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	F_ScanKeyInit(m, v197, int32(2), int32(3), int32(62), v171)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v210 = int32(3)
	F_ScanKeyInit(m, v197+int32(48), v210, v210, int32(184), l3)
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v219 = F_systable_beginscan(m, v201, int32(2664), int32(1), int32(0), int32(2), v197)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v221 = F_systable_getnext(m, v219)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_systable_endscan(m, v219)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_relation_close(m, v201, int32(1))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	m.G0 = v197 + int32(96)
	if v221 == int32(0) {
		goto L36
	} else {
		goto L54
	}
L54:
	;
	goto L45
L55:
	;
	v239 = v163 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v239
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l2
	v246 = F_pg_snprintf(m, v17+int32(192), int32(64), int32(_a_F_ChooseRelationName_0), v17)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v163 = v239
	goto L34
L57:
	;
	m.G0 = v17 + int32(256)
	return v171
}
func F_DropRelationFiles(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int64
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int64
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	v9 = m.G0
	v11 = v9 - int32(80)
	m.G0 = v11
	v15 = F_palloc(m, l1<<(uint(int32(2))%32))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if int32(0) < l1 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	F_pfree(m, v15)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L24
	}
L4:
	;
	v25 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_smgrdounlinkall(m, v15, l1, l2)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L23
	}
L7:
	;
	v29 = l0 + v25*int32(12)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v32
	v37 = F_smgropen(m, v11-int32(-64), int32(-1))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	F_smgrdounlinkall(m, v15, l1, l2)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L18
	}
L9:
	;
	if l2 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+56)) = v39
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+48)) = v41
	F_XLogDropRelation(m, v11+int32(48), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15+v25<<(uint(int32(2))%32)))) = v37
	v78 = v25 + int32(1)
	if v78 != l1 {
		v25 = v78
		goto L7
	} else {
		goto L17
	}
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v48
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v50
	F_XLogDropRelation(m, v11+int32(32), int32(1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v59
	F_XLogDropRelation(m, v11+int32(16), int32(2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v66
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v68
	F_XLogDropRelation(m, v11, int32(3))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	goto L12
L17:
	;
	goto L8
L18:
	;
	v83 = int32(0)
	goto L19
L19:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v15+v83<<(uint(int32(2))%32))))
	F_smgrclose(m, v94)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L3
L21:
	;
	v98 = v83 + int32(1)
	if v98 != l1 {
		v83 = v98
		goto L19
	} else {
		goto L22
	}
L22:
	;
	goto L20
L23:
	;
	goto L3
L24:
	;
	m.G0 = v11 + int32(80)
	return
}
func F_LockRelationId(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(32)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v5)+24)) = int64(72057594037927936)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = v9
	v21 = F_LockAcquireExtended(m, v5+int32(16), int32(1), v2, v2, v5+int32(12), v2)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return
	} else {
		if v21 != int32(3) {
			F_ReceiveSharedInvalidMessages(m)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v27)+53)) = uint8(v28)
				m.G0 = v5 + int32(32)
				return
			}
		} else {
			m.G0 = v5 + int32(32)
			return
		}
	}
}
func F_RelationCacheInitFilePreInvalidate(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	v3 = m.G0
	v5 = v3 - int32(2080)
	m.G0 = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[0]))
	if v8 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+20)) = int32(_a_F_RelationCacheInitFilePreInvalidate_0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = v8
		v18 = F_pg_snprintf(m, v5+int32(1056), int32(1024), int32(_a_F_RelationCacheInitFilePreInvalidate_1), v5+int32(16))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_RelationCacheInitFilePreInvalidate_0)
			v26 = F_pg_snprintf(m, v5+int32(32), int32(1024), int32(_a_F_RelationCacheInitFilePreInvalidate_2), v5)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[1]))
				v33 = F_LWLockAcquire(m, v29+int32(2048), int32(0))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[0]))
					if v36 != 0 {
						F_unlink_initfile(m, v5+int32(1056), int32(21))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return
						} else {
							F_unlink_initfile(m, v5+int32(32), int32(21))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								m.G0 = v5 + int32(2080)
								return
							}
						}
					} else {
						F_unlink_initfile(m, v5+int32(32), int32(21))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v5 + int32(2080)
							return
						}
					}
				}
			}
		}
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = int32(_a_F_RelationCacheInitFilePreInvalidate_0)
		v26 = F_pg_snprintf(m, v5+int32(32), int32(1024), int32(_a_F_RelationCacheInitFilePreInvalidate_2), v5)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[1]))
			v33 = F_LWLockAcquire(m, v29+int32(2048), int32(0))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_RelationCacheInitFilePreInvalidate[0]))
				if v36 != 0 {
					F_unlink_initfile(m, v5+int32(1056), int32(21))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						F_unlink_initfile(m, v5+int32(32), int32(21))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							m.G0 = v5 + int32(2080)
							return
						}
					}
				} else {
					F_unlink_initfile(m, v5+int32(32), int32(21))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						m.G0 = v5 + int32(2080)
						return
					}
				}
			}
		}
	}
}
func F_RelationClose(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v5 = v3 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v5
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClose[0]))
	if v8 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, _c_F_RelationClose[1]))
		F_ResourceOwnerForget(m, v10, l0, int32(_a_F_RelationClose_0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v15 = v14
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
				if v16 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v24 == int32(0) {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						if v27 == int32(0) {
							return
						} else {
							F_MemoryContextDeleteChildren(m, v24)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
					if v19 == int32(0) {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if v24 == int32(0) {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
							if v27 == int32(0) {
								return
							} else {
								F_MemoryContextDeleteChildren(m, v24)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									return
								}
							}
						}
					} else {
						F_MemoryContextDeleteChildren(m, v16)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
							if v24 == int32(0) {
								return
							} else {
								v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
								if v27 == int32(0) {
									return
								} else {
									F_MemoryContextDeleteChildren(m, v24)
									mBase = m.M
									v31 = m.ExcPending
									if v31 != 0 {
										return
									} else {
										return
									}
								}
							}
						}
					}
				}
			}
		}
	} else {
		v15 = v5
		if v15 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
			if v16 == int32(0) {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
				if v24 == int32(0) {
					return
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
					if v27 == int32(0) {
						return
					} else {
						F_MemoryContextDeleteChildren(m, v24)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
				if v19 == int32(0) {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
					if v24 == int32(0) {
						return
					} else {
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
						if v27 == int32(0) {
							return
						} else {
							F_MemoryContextDeleteChildren(m, v24)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								return
							}
						}
					}
				} else {
					F_MemoryContextDeleteChildren(m, v16)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
						if v24 == int32(0) {
							return
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
							if v27 == int32(0) {
								return
							} else {
								F_MemoryContextDeleteChildren(m, v24)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_RelationGetIndexAttrBitmap(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v277 int32
	_ = v277
	var v279 int32
	_ = v279
	var v290 int32
	_ = v290
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v421 int32
	_ = v421
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v565 int32
	_ = v565
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v579 int32
	_ = v579
	v3 = int32(0)
	v20 = m.G0
	v22 = v20 - int32(32)
	m.G0 = v22
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)))
	if v24 == int32(1) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v22 + int32(32)
	return v579
L2:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	v572 = F_bms_copy(m, v571)
	mBase = m.M
	v573 = m.ExcPending
	if v573 != 0 {
		goto L11
	} else {
		goto L130
	}
L3:
	;
	switch l1 {
	case 0:
		goto L2
	case 1:
		goto L10
	case 2:
		goto L9
	case 3:
		goto L8
	case 4:
		goto L7
	default:
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+116)))
	if v55 != int32(1) {
		v579 = v3
		goto L1
	} else {
		goto L19
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L11
	} else {
		goto L16
	}
L7:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	v39 = F_bms_copy(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L15
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v36 = F_bms_copy(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L11
	} else {
		goto L14
	}
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	v33 = F_bms_copy(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L11
	} else {
		goto L13
	}
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v28 = F_bms_copy(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v579 = v28
	goto L1
L13:
	;
	v579 = v33
	goto L1
L14:
	;
	v579 = v36
	goto L1
L15:
	;
	v579 = v39
	goto L1
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = l1
	F_errmsg_internal(m, int32(_a_F_RelationGetIndexAttrBitmap_0), v22)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L11
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_RelationGetIndexAttrBitmap_1), int32(_a_F_RelationGetIndexAttrBitmap_2), int32(_a_F_RelationGetIndexAttrBitmap_3))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	v58 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	if v58 == int32(0) {
		v579 = v3
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v71 = v58
	goto L27
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L11
	} else {
		goto L127
	}
L23:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v579 = v555
	goto L1
L24:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v579 = v554
	goto L1
L25:
	;
	v579 = v460
	goto L1
L26:
	;
	v579 = v461
	goto L1
L27:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v83 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v83
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v83 < v90 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v579 = int32(0)
	goto L1
L29:
	;
	v99 = v83
	v106 = v83
	v107 = v83
	v108 = int32(0)
	goto L32
L30:
	;
	v453 = v83
	v460 = v83
	v461 = v83
	goto L31
L31:
	;
	v467 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L11
	} else {
		goto L101
	}
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v71)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v108<<(uint(int32(2))%32))))
	v119 = F_index_open(m, v117, int32(1))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L11
	} else {
		goto L34
	}
L33:
	;
	v453 = v421
	v460 = v428
	v461 = v429
	goto L31
L34:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v119)+196))
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[0]))
	if v123 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v127 = int32(_a_F_RelationGetIndexAttrBitmap_4)
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1]))
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v131
	v134 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	v184 = v123
	goto L37
L37:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+18)))
	if base.Ui32(v202&int32(2044)) <= base.Ui32(int32(19)) {
		goto L44
	} else {
		goto L45
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v134)+4)) = int64(-4294965047)
	v141 = int32(0)
	goto L39
L39:
	;
	v157 = int32(100)
	v158 = v141 * v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
	base.MemoryCopy(m, v158+(v134+v159<<(uint(int32(4))%32))+int32(20), v158+int32(_a_F_RelationGetIndexAttrBitmap_5), v157)
	F_populate_compact_attribute(m, v134, v141)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L11
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v134)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[0])) = v134
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v128
	v184 = v134
	goto L37
L41:
	;
	v173 = v141 + int32(1)
	if v173 != int32(21) {
		v141 = v173
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v218 = int32(0)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+23)))
	if v219 == v218 {
		goto L49
	} else {
		goto L50
	}
L44:
	;
	v210 = F_getmissingattr(m, v184, int32(20), v22+int32(23))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L11
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v215 = F_fastgetattr_3(m, v121, int32(20), v184, v22+int32(23))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L11
	} else {
		goto L48
	}
L47:
	;
	v217 = v210
	goto L43
L48:
	;
	v217 = v215
	goto L43
L49:
	;
	v222 = F_text_to_cstring(m, v217)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L11
	} else {
		goto L52
	}
L50:
	;
	v226 = v218
	goto L51
L51:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v119)+196))
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[0]))
	if v229 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v224 = F_stringToNode(m, v222)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L11
	} else {
		goto L53
	}
L53:
	;
	v226 = v224
	goto L51
L54:
	;
	v233 = int32(_a_F_RelationGetIndexAttrBitmap_4)
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1]))
	v237 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v237
	v240 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L11
	} else {
		goto L57
	}
L55:
	;
	v290 = v229
	goto L56
L56:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v227)+16))
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+18)))
	if base.Ui32(v308&int32(2047)) <= base.Ui32(int32(20)) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v240)+4)) = int64(-4294965047)
	v247 = int32(0)
	goto L58
L58:
	;
	v263 = int32(100)
	v264 = v247 * v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v240)))
	base.MemoryCopy(m, v264+(v240+v265<<(uint(int32(4))%32))+int32(20), v264+int32(_a_F_RelationGetIndexAttrBitmap_5), v263)
	F_populate_compact_attribute(m, v240, v247)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L11
	} else {
		goto L60
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[0])) = v240
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v234
	v290 = v240
	goto L56
L60:
	;
	v279 = v247 + int32(1)
	if v279 != int32(21) {
		v247 = v279
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
L62:
	;
	v324 = int32(0)
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+23)))
	if v326 == v324 {
		goto L68
	} else {
		goto L69
	}
L63:
	;
	v316 = F_getmissingattr(m, v290, int32(21), v22+int32(23))
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L11
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v321 = F_fastgetattr_3(m, v227, int32(21), v290, v22+int32(23))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L11
	} else {
		goto L67
	}
L66:
	;
	v323 = v316
	goto L62
L67:
	;
	v323 = v321
	goto L62
L68:
	;
	v329 = F_text_to_cstring(m, v323)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L11
	} else {
		goto L71
	}
L69:
	;
	v333 = v324
	goto L70
L70:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v119)+204))
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v338)+28)))
	if v339 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v331 = F_stringToNode(m, v329)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L11
	} else {
		goto L72
	}
L72:
	;
	v333 = v331
	goto L70
L73:
	;
	v340 = v22 + int32(24)
	goto L75
L74:
	;
	v340 = v22 + int32(28)
	goto L75
L75:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	v342 = int32(*(*int16)(unsafe.Add(mBase, uint32(v341)+8)))
	if int32(0) < v342 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v345 = int32(0)
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+12)))
	v354 = v324
	v355 = v341
	v357 = v99
	v364 = v106
	v365 = v107
	goto L79
L77:
	;
	v421 = v99
	v428 = v106
	v429 = v107
	goto L78
L78:
	;
	F_pull_varattnos(m, v226, int32(1), v340)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L11
	} else {
		goto L96
	}
L79:
	;
	v374 = int32(*(*int16)(unsafe.Add(mBase, uint32(v355+v354<<(uint(int32(1))%32))+48)))
	if v374 == int32(0) {
		v407 = v355
		v408 = v357
		v410 = v364
		v411 = v365
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v421 = v408
	v428 = v410
	v429 = v411
	goto L78
L81:
	;
	v413 = v354 + int32(1)
	v414 = int32(*(*int16)(unsafe.Add(mBase, uint32(v407)+8)))
	if v413 < v414 {
		v354 = v413
		v355 = v407
		v357 = v408
		v364 = v410
		v365 = v411
		goto L79
	} else {
		goto L95
	}
L82:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v379 = v374 + int32(7)
	v380 = F_bms_add_member(m, v377, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v340))) = v380
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	if base.B2i32(v333 == v345)&(v347&base.B2i32(v226 == v345)) == int32(0) {
		v391 = v383
		v392 = v357
		goto L84
	} else {
		goto L85
	}
L84:
	;
	if v82 != v117 {
		v399 = v391
		v400 = v365
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v386 = int32(*(*int16)(unsafe.Add(mBase, uint32(v383)+10)))
	if v386 <= v354 {
		v391 = v383
		v392 = v357
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v388 = F_bms_add_member(m, v357, v379)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	v391 = v390
	v392 = v388
	goto L84
L88:
	;
	if v81 != v117 {
		v407 = v399
		v408 = v392
		v410 = v364
		v411 = v400
		goto L81
	} else {
		goto L92
	}
L89:
	;
	v394 = int32(*(*int16)(unsafe.Add(mBase, uint32(v391)+10)))
	if v394 <= v354 {
		v399 = v391
		v400 = v365
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v396 = F_bms_add_member(m, v365, v379)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	v399 = v398
	v400 = v396
	goto L88
L92:
	;
	v402 = int32(*(*int16)(unsafe.Add(mBase, uint32(v399)+10)))
	if v402 <= v354 {
		v407 = v399
		v408 = v392
		v410 = v364
		v411 = v400
		goto L81
	} else {
		goto L93
	}
L93:
	;
	v404 = F_bms_add_member(m, v364, v379)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v119)+192))
	v407 = v406
	v408 = v392
	v410 = v404
	v411 = v400
	goto L81
L95:
	;
	goto L80
L96:
	;
	F_pull_varattnos(m, v333, int32(1), v340)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	F_relation_close(m, v119, int32(1))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L11
	} else {
		goto L98
	}
L98:
	;
	v445 = v108 + int32(1)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v445 < v446 {
		v99 = v421
		v106 = v428
		v107 = v429
		v108 = v445
		goto L32
	} else {
		goto L99
	}
L99:
	;
	goto L33
L100:
	;
	F_list_free(m, v467)
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L11
	} else {
		goto L118
	}
L101:
	;
	v469 = F_equal(m, v71, v467)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	if v469 == int32(0) {
		goto L100
	} else {
		goto L103
	}
L103:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v82 != v473 {
		goto L100
	} else {
		goto L104
	}
L104:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v81 != v475 {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	F_list_free(m, v467)
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	F_list_free(m, v71)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L11
	} else {
		goto L107
	}
L107:
	;
	v481 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v481)
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+156))
	F_bms_free(m, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L11
	} else {
		goto L108
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = int32(0)
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	F_bms_free(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L11
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = int32(0)
	v493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+164))
	F_bms_free(m, v493)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = int32(0)
	v498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	F_bms_free(m, v498)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = int32(0)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+172))
	F_bms_free(m, v503)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L11
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = int32(0)
	v508 = int32(_a_F_RelationGetIndexAttrBitmap_4)
	v509 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1]))
	v512 = *(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v512
	v514 = F_bms_copy(m, v453)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L11
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+156)) = v514
	v517 = F_bms_copy(m, v461)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L11
	} else {
		goto L114
	}
L114:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+160)) = v517
	v520 = F_bms_copy(m, v460)
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L11
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+164)) = v520
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	v524 = F_bms_copy(m, v523)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+168)) = v524
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	v528 = F_bms_copy(m, v527)
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	v530 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+152)) = uint8(v530)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+172)) = v528
	*(*int32)(unsafe.Add(mBase, _c_F_RelationGetIndexAttrBitmap[1])) = v509
	switch l1 {
	case 0:
		v579 = v453
		goto L1
	case 1:
		goto L26
	case 2:
		goto L25
	case 3:
		goto L24
	case 4:
		goto L23
	default:
		goto L22
	}
L118:
	;
	F_list_free(m, v71)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L11
	} else {
		goto L119
	}
L119:
	;
	F_bms_free(m, v453)
	mBase = m.M
	v540 = m.ExcPending
	if v540 != 0 {
		goto L11
	} else {
		goto L120
	}
L120:
	;
	F_bms_free(m, v461)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L11
	} else {
		goto L121
	}
L121:
	;
	F_bms_free(m, v460)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L11
	} else {
		goto L122
	}
L122:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_bms_free(m, v545)
	mBase = m.M
	v547 = m.ExcPending
	if v547 != 0 {
		goto L11
	} else {
		goto L123
	}
L123:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
	F_bms_free(m, v548)
	mBase = m.M
	v550 = m.ExcPending
	if v550 != 0 {
		goto L11
	} else {
		goto L124
	}
L124:
	;
	v551 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L11
	} else {
		goto L125
	}
L125:
	;
	if v551 != 0 {
		v71 = v551
		goto L27
	} else {
		goto L126
	}
L126:
	;
	goto L28
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = l1
	F_errmsg_internal(m, int32(_a_F_RelationGetIndexAttrBitmap_0), v22+int32(16))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L11
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(_a_F_RelationGetIndexAttrBitmap_1), int32(_a_F_RelationGetIndexAttrBitmap_6), int32(_a_F_RelationGetIndexAttrBitmap_3))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L11
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	v579 = v572
	goto L1
}
func F_RelationInitIndexAccessInfo(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v325 int32
	_ = v325
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v535 int32
	_ = v535
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v662 int32
	_ = v662
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int64
	_ = v693
	var v695 int32
	_ = v695
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v798 int32
	_ = v798
	v21 = m.G0
	v23 = v21 - int32(288)
	m.G0 = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v27 = F_SearchSysCache1(m, int32(34), v26)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L6
	} else {
		goto L138
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L6
	} else {
		goto L135
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L6
	} else {
		goto L132
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L6
	} else {
		goto L129
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L6
	} else {
		goto L126
	}
L6:
	;
	return
L7:
	;
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v29 = int32(_a_F_RelationInitIndexAccessInfo_0)
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0]))
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v33
	v35 = F_heap_copytuple(m, v27)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v707 = m.ExcPending
	if v707 != 0 {
		goto L6
	} else {
		goto L123
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v35
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+16))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+22)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v38 + v39
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v30
	F_ReleaseCatCache(m, v27)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L12
	}
L12:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+84))
	v49 = F_SearchSysCache1(m, int32(2), v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L6
	} else {
		goto L13
	}
L13:
	;
	if v49 == int32(0) {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v53+v54)+68))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+184)) = v56
	F_ReleaseCatCache(m, v49)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v61 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+120)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62)+8)))
	if v61 != v63 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v62)+10)))
	v66 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	v73 = F_AllocSetContextCreateInternal(m, v68, int32(_a_F_RelationInitIndexAccessInfo_1), v66, int32(1024), int32(_a_F_RelationInitIndexAccessInfo_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+200)) = v73
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v79 = F_MemoryContextStrdup(m, v73, v76+int32(4))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73)+36)) = v79
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v83 = F_GetIndexAmRoutine(m, v82)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L6
	} else {
		goto L19
	}
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
	v87 = F_MemoryContextAlloc(m, v85, int32(140))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L6
	} else {
		goto L20
	}
L20:
	;
	base.MemoryCopy(m, v87, v83, int32(140))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+204)) = v87
	F_pfree(m, v83)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	v95 = v65 << (uint(int32(2)) % 32)
	v96 = F_MemoryContextAllocZero(m, v73, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+208)) = v96
	v99 = F_MemoryContextAllocZero(m, v73, v95)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L6
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+212)) = v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+6)))
	if v103 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+220)) = v117
	v119 = F_MemoryContextAllocZero(m, v73, v95)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L6
	} else {
		goto L30
	}
L25:
	;
	v105 = v103 * base.I32_extend16_s(v61)
	v108 = F_MemoryContextAllocZero(m, v73, v105<<(uint(int32(2))%32))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L6
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = int32(0)
	v117 = v66
	goto L24
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+216)) = v108
	v113 = F_MemoryContextAllocZero(m, v73, v105*int32(28))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L6
	} else {
		goto L29
	}
L29:
	;
	v117 = v113
	goto L24
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+248)) = v119
	v123 = v65 << (uint(int32(1)) % 32)
	v124 = F_MemoryContextAllocZero(m, v73, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+224)) = v124
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v128 = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2]))
	if v130 == v128 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v133 = int32(_a_F_RelationInitIndexAccessInfo_0)
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0]))
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v137
	v140 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L6
	} else {
		goto L35
	}
L33:
	;
	v190 = v130
	goto L34
L34:
	;
	v212 = F_fastgetattr_3(m, v127, int32(17), v190, v23+int32(79))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L6
	} else {
		goto L40
	}
L35:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v140)+4)) = int64(-4294965047)
	v146 = v128
	goto L36
L36:
	;
	v164 = int32(100)
	v165 = v146 * v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	base.MemoryCopy(m, v165+(v140+v166<<(uint(int32(4))%32))+int32(20), v165+int32(_a_F_RelationInitIndexAccessInfo_3), v164)
	F_populate_compact_attribute(m, v140, v146)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L38
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v140)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2])) = v140
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v134
	v190 = v140
	goto L34
L38:
	;
	v180 = v146 + int32(1)
	if v180 != int32(21) {
		v146 = v180
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	if v95 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	base.MemoryCopy(m, v214, v212+int32(24), v95)
	goto L43
L42:
	;
	goto L43
L43:
	;
	v218 = int32(0)
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2]))
	if v221 == v218 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v224 = int32(_a_F_RelationInitIndexAccessInfo_0)
	v225 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0]))
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v228
	v231 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L47
	}
L45:
	;
	v281 = v221
	goto L46
L46:
	;
	v303 = F_fastgetattr_3(m, v219, int32(18), v281, v23+int32(79))
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L6
	} else {
		goto L52
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v231)+4)) = int64(-4294965047)
	v237 = v218
	goto L48
L48:
	;
	v255 = int32(100)
	v256 = v237 * v255
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v231)))
	base.MemoryCopy(m, v256+(v231+v257<<(uint(int32(4))%32))+int32(20), v256+int32(_a_F_RelationInitIndexAccessInfo_3), v255)
	F_populate_compact_attribute(m, v231, v237)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L6
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v231)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2])) = v231
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v225
	v281 = v231
	goto L46
L50:
	;
	v271 = v237 + int32(1)
	if v271 != int32(21) {
		v237 = v271
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	if int32(0) < v65 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+212))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+208))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+216))
	v311 = v103 << (uint(int32(2)) % 32)
	v325 = int32(0)
	goto L56
L54:
	;
	goto L55
L55:
	;
	v599 = int32(0)
	v600 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v602 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2]))
	if v602 == v599 {
		goto L110
	} else {
		goto L111
	}
L56:
	;
	v340 = v325 << (uint(int32(2)) % 32)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v303+int32(24)+v340)))
	if v342 == int32(0) {
		goto L3
	} else {
		goto L58
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+284)) = v342
	v347 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[3]))
	if v347 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v351 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	if v351 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v367 = v347
	goto L61
L61:
	;
	v373 = F_hash_search(m, v367, v23+int32(284), int32(1), v23+int32(80))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L6
	} else {
		goto L67
	}
L62:
	;
	F_CreateCacheMemoryContext(m)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L6
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = int64(85899345924)
	v364 = F_hash_create(m, int32(_a_F_RelationInitIndexAccessInfo_4), int32(64), v23+int32(80), int32(40))
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L6
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[3])) = v364
	v367 = v364
	goto L61
L67:
	;
	v375 = int32(0)
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+80)))
	if v377 == v375 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v340+v308))) = v558
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v340+v307))) = v561
	v563 = int32(0)
	if base.B2i32(v103 == v563)|base.B2i32(v311 == v563) == v563 {
		goto L106
	} else {
		goto L107
	}
L69:
	;
	v391 = int32(0)
	if base.B2i32(v103 == v375)|base.B2i32(v390 == v391) == v391 {
		goto L74
	} else {
		goto L75
	}
L70:
	;
	v380 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v373)+16)) = v380
	*(*uint16)(unsafe.Add(mBase, uint32(v373)+6)) = uint16(v103)
	*(*uint8)(unsafe.Add(mBase, uint32(v373)+4)) = uint8(v380)
	v390 = int32(1)
	goto L69
L71:
	;
	goto L72
L72:
	;
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+4)))
	if v386 != 0 {
		goto L68
	} else {
		goto L73
	}
L73:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	v390 = base.B2i32(v387 == int32(0))
	goto L69
L74:
	;
	v397 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	v398 = F_MemoryContextAllocZero(m, v397, v311)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L6
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	v410 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[4])))
	if v410 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373)+16)) = v398
	goto L76
L78:
	;
	v411 = int32(1)
	goto L80
L79:
	;
	v411 = base.B2i32(v402 != int32(1981)) & base.B2i32(v402 != int32(1979))
	goto L80
L80:
	;
	v413 = v23 + int32(128)
	F_ScanKeyInit(m, v413, int32(1), int32(3), int32(184), v402)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L6
	} else {
		goto L81
	}
L81:
	;
	v421 = F_table_open(m, int32(2616), int32(1))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v426 = F_systable_beginscan(m, v421, int32(2687), v411, int32(0), int32(1), v413)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L6
	} else {
		goto L83
	}
L83:
	;
	v428 = F_systable_getnext(m, v426)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L6
	} else {
		goto L84
	}
L84:
	;
	if v428 == int32(0) {
		goto L2
	} else {
		goto L85
	}
L85:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v428)+16))
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v432)+22)))
	v434 = v432 + v433
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v434)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v373)+8)) = v435
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v434)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v373)+12)) = v437
	F_systable_endscan(m, v426)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L6
	} else {
		goto L86
	}
L86:
	;
	F_relation_close(m, v421, int32(1))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L6
	} else {
		goto L87
	}
L87:
	;
	if v103 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	F_ScanKeyInit(m, v413, int32(2), int32(3), int32(184), v447)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L6
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v535 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v373)+4)) = uint8(v535)
	goto L68
L91:
	;
	v450 = int32(3)
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	F_ScanKeyInit(m, v23+int32(176), v450, v450, int32(184), v453)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L6
	} else {
		goto L92
	}
L92:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v373)+12))
	F_ScanKeyInit(m, v23+int32(224), int32(4), int32(3), int32(184), v459)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L6
	} else {
		goto L93
	}
L93:
	;
	v464 = F_table_open(m, int32(2603), int32(1))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L6
	} else {
		goto L94
	}
L94:
	;
	v469 = F_systable_beginscan(m, v464, int32(2655), v411, int32(0), int32(3), v413)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L6
	} else {
		goto L95
	}
L95:
	;
	goto L96
L96:
	;
	v491 = F_systable_getnext(m, v469)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L6
	} else {
		goto L98
	}
L97:
	;
	F_systable_endscan(m, v469)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L6
	} else {
		goto L104
	}
L98:
	;
	if v491 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v491)+16))
	v494 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+22)))
	v495 = v493 + v494
	v496 = int32(*(*int16)(unsafe.Add(mBase, uint32(v495)+16)))
	if v496 <= int32(0) {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	goto L97
L102:
	;
	v500 = v496 & int32(_a_F_RelationInitIndexAccessInfo_5)
	if base.Ui32(v103) < base.Ui32(v500) {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v495)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v502+v500<<(uint(int32(2))%32)-int32(4)))) = v508
	goto L96
L104:
	;
	F_relation_close(m, v464, int32(1))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L6
	} else {
		goto L105
	}
L105:
	;
	goto L90
L106:
	;
	v574 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	base.MemoryCopy(m, v309+v103*v325<<(uint(int32(2))%32), v574, v311)
	goto L108
L107:
	;
	goto L108
L108:
	;
	v577 = v325 + int32(1)
	if v577 != v65 {
		v325 = v577
		goto L56
	} else {
		goto L109
	}
L109:
	;
	goto L57
L110:
	;
	v605 = int32(_a_F_RelationInitIndexAccessInfo_0)
	v606 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0]))
	v609 = *(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v609
	v612 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L6
	} else {
		goto L113
	}
L111:
	;
	v662 = v602
	goto L112
L112:
	;
	v684 = F_fastgetattr_3(m, v600, int32(19), v662, v23+int32(79))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L6
	} else {
		goto L118
	}
L113:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v612)+4)) = int64(-4294965047)
	v618 = v599
	goto L114
L114:
	;
	v636 = int32(100)
	v637 = v618 * v636
	v638 = *(*int32)(unsafe.Add(mBase, uint32(v612)))
	base.MemoryCopy(m, v637+(v612+v638<<(uint(int32(4))%32))+int32(20), v637+int32(_a_F_RelationInitIndexAccessInfo_3), v636)
	F_populate_compact_attribute(m, v612, v618)
	mBase = m.M
	v650 = m.ExcPending
	if v650 != 0 {
		goto L6
	} else {
		goto L116
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v612)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[2])) = v612
	*(*int32)(unsafe.Add(mBase, _c_F_RelationInitIndexAccessInfo[0])) = v606
	v662 = v612
	goto L112
L116:
	;
	v652 = v618 + int32(1)
	if v652 != int32(21) {
		v618 = v652
		goto L114
	} else {
		goto L117
	}
L117:
	;
	goto L115
L118:
	;
	if v123 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	base.MemoryCopy(m, v686, v684+int32(24), v123)
	goto L121
L120:
	;
	goto L121
L121:
	;
	v691 = F_RelationGetIndexAttOptions(m, l0, int32(0))
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L6
	} else {
		goto L122
	}
L122:
	;
	v693 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+228)) = v693
	v695 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v695
	*(*int64)(unsafe.Add(mBase, uint32(l0)+236)) = v693
	*(*int32)(unsafe.Add(mBase, uint32(l0)+244)) = v695
	m.G0 = v23 + int32(288)
	return
L123:
	;
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v708
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_6), v23)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L6
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_7), int32(1471), int32(_a_F_RelationInitIndexAccessInfo_8))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L6
	} else {
		goto L125
	}
L125:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L126:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v722)+84))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v723
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_9), v23+int32(16))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L6
	} else {
		goto L127
	}
L127:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_7), int32(1485), int32(_a_F_RelationInitIndexAccessInfo_8))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L6
	} else {
		goto L128
	}
L128:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L129:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v739
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_10), v23-int32(-64))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L6
	} else {
		goto L130
	}
L130:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_7), int32(1493), int32(_a_F_RelationInitIndexAccessInfo_8))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L6
	} else {
		goto L131
	}
L131:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L132:
	;
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_11), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L6
	} else {
		goto L133
	}
L133:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_7), int32(1630), int32(_a_F_RelationInitIndexAccessInfo_12))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L6
	} else {
		goto L134
	}
L134:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L135:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v768
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_13), v23+int32(32))
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L6
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_7), int32(1766), int32(_a_F_RelationInitIndexAccessInfo_14))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L6
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	v785 = int32(*(*int16)(unsafe.Add(mBase, uint32(v495)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v785
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v23)+284))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v787
	F_errmsg_internal(m, int32(_a_F_RelationInitIndexAccessInfo_15), v23+int32(48))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L6
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_RelationInitIndexAccessInfo_7), int32(1800), int32(_a_F_RelationInitIndexAccessInfo_14))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L6
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RelationMapFilenumberToOid(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v3 = int32(0)
	if l1 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v105
L2:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v105 = v100
	goto L1
L3:
	;
	v8 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapFilenumberToOid[0]))
	if v8 < v10 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapFilenumberToOid[1]))
	if v51 < v53 {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	v14 = v8
	goto L9
L7:
	;
	goto L8
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapFilenumberToOid[2]))
	if v33 <= int32(0) {
		v105 = v3
		goto L1
	} else {
		goto L15
	}
L9:
	;
	v19 = v14 << (uint(int32(3)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_RelationMapFilenumberToOid[3])))
	if v20 == l0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v99 = v19 + int32(_a_F_RelationMapFilenumberToOid_0)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v25 = v14 + int32(1)
	if v25 != v10 {
		v14 = v25
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v38 = int32(0)
	goto L16
L16:
	;
	v43 = v38 << (uint(int32(3)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_RelationMapFilenumberToOid[4])))
	if v44 != l0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v99 = v43 + int32(_a_F_RelationMapFilenumberToOid_1)
	goto L2
L18:
	;
	v47 = v38 + int32(1)
	if v33 != v47 {
		v38 = v47
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v105 = v3
	goto L1
L22:
	;
	v81 = int32(0)
	goto L32
L23:
	;
	v99 = v62 + int32(_a_F_RelationMapFilenumberToOid_2)
	goto L2
L24:
	;
	v57 = v51
	goto L27
L25:
	;
	goto L26
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapFilenumberToOid[5]))
	if v74 <= int32(0) {
		v105 = v3
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v62 = v57 << (uint(int32(3)) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_c_F_RelationMapFilenumberToOid[6])))
	if l0 == v63 {
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v66 = v57 + int32(1)
	if v66 != v53 {
		v57 = v66
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L22
L32:
	;
	v86 = v81 << (uint(int32(3)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+uint32(_c_F_RelationMapFilenumberToOid[7])))
	if v87 != l0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v99 = v86 + int32(_a_F_RelationMapFilenumberToOid_3)
	goto L2
L34:
	;
	v90 = v81 + int32(1)
	if v74 != v90 {
		v81 = v90
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v105 = v3
	goto L1
}
func F_RelationMapOidToFilenumber(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	v3 = int32(0)
	if l1 == v3 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v105
L2:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	v105 = v100
	goto L1
L3:
	;
	v8 = int32(0)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapOidToFilenumber[0]))
	if v8 < v10 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v51 = int32(0)
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapOidToFilenumber[1]))
	if v51 < v53 {
		goto L24
	} else {
		goto L25
	}
L6:
	;
	v14 = v8
	goto L9
L7:
	;
	goto L8
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapOidToFilenumber[2]))
	if v33 <= int32(0) {
		v105 = v3
		goto L1
	} else {
		goto L15
	}
L9:
	;
	v19 = v14 << (uint(int32(3)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+uint32(_c_F_RelationMapOidToFilenumber[3])))
	if v20 == l0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v99 = v19 + int32(_a_F_RelationMapOidToFilenumber_0)
	goto L2
L12:
	;
	goto L13
L13:
	;
	v25 = v14 + int32(1)
	if v25 != v10 {
		v14 = v25
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	v38 = int32(0)
	goto L16
L16:
	;
	v43 = v38 << (uint(int32(3)) % 32)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_c_F_RelationMapOidToFilenumber[4])))
	if v44 != l0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v99 = v43 + int32(_a_F_RelationMapOidToFilenumber_1)
	goto L2
L18:
	;
	v47 = v38 + int32(1)
	if v33 != v47 {
		v38 = v47
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v105 = v3
	goto L1
L22:
	;
	v81 = int32(0)
	goto L32
L23:
	;
	v99 = v62 + int32(_a_F_RelationMapOidToFilenumber_2)
	goto L2
L24:
	;
	v57 = v51
	goto L27
L25:
	;
	goto L26
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_RelationMapOidToFilenumber[5]))
	if v74 <= int32(0) {
		v105 = v3
		goto L1
	} else {
		goto L31
	}
L27:
	;
	v62 = v57 << (uint(int32(3)) % 32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+uint32(_c_F_RelationMapOidToFilenumber[6])))
	if l0 == v63 {
		goto L23
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	v66 = v57 + int32(1)
	if v66 != v53 {
		v57 = v66
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L22
L32:
	;
	v86 = v81 << (uint(int32(3)) % 32)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+uint32(_c_F_RelationMapOidToFilenumber[7])))
	if v87 != l0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v99 = v86 + int32(_a_F_RelationMapOidToFilenumber_3)
	goto L2
L34:
	;
	v90 = v81 + int32(1)
	if v74 != v90 {
		v81 = v90
		goto L32
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L33
L37:
	;
	v105 = v3
	goto L1
}
func F_RelationMapOidToFilenumberForDatabase(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(528)
	m.G0 = v9
	F_read_relmap_file(m, v9+int32(4), l0, v3, int32(21))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v19 <= int32(0) {
		v42 = v3
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v9 + int32(528)
	return v42
L4:
	;
	v26 = v3
	goto L6
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v42 = v39
	goto L3
L6:
	;
	v32 = v9 + int32(12) + v26<<(uint(int32(3))%32)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	if l1 == v33 {
		goto L5
	} else {
		goto L8
	}
L7:
	;
	v42 = int32(0)
	goto L3
L8:
	;
	v36 = v26 + int32(1)
	if v36 != v19 {
		v26 = v36
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
}
func F_RelationTruncateIndexes(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v104 int32
	_ = v104
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v265 int32
	_ = v265
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v314 int32
	_ = v314
	var v332 int32
	_ = v332
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	v17 = F_RelationGetIndexList(m, l0)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	return
L3:
	;
	if v17 == int32(0) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v21 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v37 = int32(0)
	goto L6
L6:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v40+v37<<(uint(int32(2))%32))))
	v46 = F_index_open(m, v44, int32(8))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L2
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	F_RelationTruncate(m, v46, int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L2
	} else {
		goto L63
	}
L9:
	;
	v48 = int32(0)
	v50 = m.G0
	v52 = v50 - int32(16)
	m.G0 = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)+192))
	v55 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+8)))
	if base.Ui32(int32(_a_F_RelationTruncateIndexes_0)) < base.Ui32((v55-int32(33))&int32(_a_F_RelationTruncateIndexes_1)) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v63 = v54 + int32(48)
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+10)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+84))
	v67 = int32(0)
	v68 = m.G0
	v70 = v68 - int32(16)
	m.G0 = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v46)+196))
	if v72 == v67 {
		v228 = v67
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L2
	} else {
		goto L60
	}
L13:
	;
	m.G0 = v70 + int32(16)
	v243 = int32(0)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+12)))
	v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+13)))
	v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+20)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v46)+204))
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v248)+28)))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+15)))
	v252 = F_makeIndexInfo(m, v55, v64, v66, v228, v243, v244, v245, v246, v243, v249, v244&v250)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L2
	} else {
		goto L48
	}
L14:
	;
	v77 = F_heap_attisnull(m, v72, int32(20), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v77 != 0 {
		v228 = v67
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v79 = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v46)+196))
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[0]))
	if v82 == v79 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v85 = int32(_a_F_RelationTruncateIndexes_2)
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[1]))
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[1])) = v89
	v92 = F_CreateTemplateTupleDesc(m, int32(21))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L2
	} else {
		goto L20
	}
L18:
	;
	v139 = v82
	goto L19
L19:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+18)))
	if base.Ui32(v154&int32(2044)) <= base.Ui32(int32(19)) {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v92)+4)) = int64(-4294965047)
	v104 = v79
	goto L21
L21:
	;
	v112 = int32(100)
	v113 = v104 * v112
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	base.MemoryCopy(m, v113+(v92+v114<<(uint(int32(4))%32))+int32(20), v113+int32(_a_F_RelationTruncateIndexes_3), v112)
	F_populate_compact_attribute(m, v92, v104)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L2
	} else {
		goto L23
	}
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v92)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[0])) = v92
	*(*int32)(unsafe.Add(mBase, _c_F_RelationTruncateIndexes[1])) = v86
	v139 = v92
	goto L19
L23:
	;
	v128 = v104 + int32(1)
	if v128 != int32(21) {
		v104 = v128
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v170 = F_text_to_cstring(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L2
	} else {
		goto L31
	}
L26:
	;
	v162 = F_getmissingattr(m, v139, int32(20), v70+int32(15))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L2
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v167 = F_fastgetattr_3(m, v80, int32(20), v139, v70+int32(15))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L2
	} else {
		goto L30
	}
L29:
	;
	v169 = v162
	goto L25
L30:
	;
	v169 = v167
	goto L25
L31:
	;
	v172 = F_stringToNode(m, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L32
	}
L32:
	;
	F_pfree(m, v170)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	if v172 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v228 = int32(0)
	goto L13
L35:
	;
	goto L36
L36:
	;
	v179 = int32(0)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v180 <= v179 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v228 = int32(0)
	goto L13
L38:
	;
	goto L39
L39:
	;
	v189 = int32(0)
	v193 = v179
	goto L40
L40:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v172)+12))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v201+v193<<(uint(int32(2))%32))))
	v206 = F_exprType(m, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L42
	}
L41:
	;
	v228 = v218
	goto L13
L42:
	;
	v208 = F_exprTypmod(m, v205)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L2
	} else {
		goto L43
	}
L43:
	;
	v210 = F_exprCollation(m, v205)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L2
	} else {
		goto L44
	}
L44:
	;
	v212 = int32(1)
	v216 = F_makeConst(m, v206, v208, v210, v212, int32(0), v212, v212)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v218 = F_lappend(m, v189, v216)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v221 = v193 + int32(1)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v172)+4))
	if v221 < v222 {
		v189 = v218
		v193 = v221
		goto L40
	} else {
		goto L47
	}
L47:
	;
	goto L41
L48:
	;
	v255 = v252 + int32(12)
	v256 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v55-int32(1)) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	m.G0 = v52 + int32(16)
	goto L8
L50:
	;
	v265 = v256
	v277 = v48
	goto L53
L51:
	;
	v314 = v256
	goto L52
L52:
	;
	v332 = v314
	v345 = v48
	goto L57
L53:
	;
	v280 = v265 << (uint(int32(1)) % 32)
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v280+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v255+v280))) = uint16(v283)
	v286 = v280 | int32(2)
	v289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v286+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v255+v286))) = uint16(v289)
	v291 = int32(4)
	v292 = v280 | v291
	v295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v292+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v255+v292))) = uint16(v295)
	v298 = v280 | int32(6)
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v298+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v255+v298))) = uint16(v301)
	v304 = v265 + v291
	v306 = v277 + v291
	if v306 != v55&int32(60) {
		v265 = v304
		v277 = v306
		goto L53
	} else {
		goto L55
	}
L54:
	;
	if v55&int32(3) == int32(0) {
		goto L49
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v314 = v304
	goto L52
L57:
	;
	v346 = int32(1)
	v347 = v332 << (uint(v346) % 32)
	v350 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v347+v63))))
	*(*uint16)(unsafe.Add(mBase, uint32(v255+v347))) = uint16(v350)
	v355 = v345 + v346
	if v355 != v55&int32(3) {
		v332 = v332 + v346
		v345 = v355
		goto L57
	} else {
		goto L59
	}
L58:
	;
	goto L49
L59:
	;
	goto L58
L60:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v46)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v52)+4)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v55
	F_errmsg_internal(m, int32(_a_F_RelationTruncateIndexes_4), v52)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L2
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(_a_F_RelationTruncateIndexes_5), int32(2499), int32(_a_F_RelationTruncateIndexes_6))
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	F_index_build(m, l0, v46, v252, int32(1), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L2
	} else {
		goto L64
	}
L64:
	;
	F_relation_close(m, v46, int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L2
	} else {
		goto L65
	}
L65:
	;
	v402 = v37 + int32(1)
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v402 < v403 {
		v37 = v402
		goto L6
	} else {
		goto L66
	}
L66:
	;
	goto L7
}
func F_SetRelationTableSpace(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
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
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v16 = F_table_open(m, int32(1259), int32(3))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v19 = F_SearchSysCacheLockedCopy1(m, int32(57), v13)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			if v19 != 0 {
				v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(v11)+12)) = uint16(v21)
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v23
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v27 = v25 + v26
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_SetRelationTableSpace[0]))
				if l1 != v30 {
					v32 = l1
				} else {
					v32 = int32(0)
				}
				*(*int32)(unsafe.Add(mBase, uint32(v27)+92)) = v32
				if l2 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v27)+88)) = l2
				} else {
				}
				v36 = v11 + int32(8)
				F_CatalogTupleUpdate(m, v16, v36, v19)
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_UnlockTuple(m, v16, v36, int32(7))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+119)))
						switch v43 - int32(83) {
						case 0, 22, 26, 31, 33:
							F_pfree(m, v19)
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
								return
							} else {
								F_relation_close(m, v16, int32(3))
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									m.G0 = v11 + int32(16)
									return
								}
							}
						default:
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
							v49 = F_table_open(m, int32(1214), int32(3))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								v51 = int32(0)
								if base.B2i32(v46 == v51)|base.B2i32(v46 == int32(1663)) == v51 {
									F_shdepChangeDep(m, v49, int32(1259), v13, int32(1213), v46, int32(116))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return
									} else {
										F_relation_close(m, v49, int32(3))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											F_pfree(m, v19)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_relation_close(m, v16, int32(3))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								} else {
									v64 = int32(0)
									F_shdepDropDependency(m, v49, int32(1259), v13, v64, int32(1), v64, v64, v64)
									mBase = m.M
									v70 = m.ExcPending
									if v70 != 0 {
										return
									} else {
										F_relation_close(m, v49, int32(3))
										mBase = m.M
										v73 = m.ExcPending
										if v73 != 0 {
											return
										} else {
											F_pfree(m, v19)
											mBase = m.M
											v77 = m.ExcPending
											if v77 != 0 {
												return
											} else {
												F_relation_close(m, v16, int32(3))
												mBase = m.M
												v80 = m.ExcPending
												if v80 != 0 {
													return
												} else {
													m.G0 = v11 + int32(16)
													return
												}
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v13
					F_errmsg_internal(m, int32(_a_F_SetRelationTableSpace_0), v11)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_SetRelationTableSpace_1), int32(3767), int32(_a_F_SetRelationTableSpace_2))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	}
}
func F_generate_relation_name(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v122 int32
	_ = v122
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	v12 = m.G0
	v14 = v12 - int32(48)
	m.G0 = v14
	v17 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v197 = F_quote_identifier(m, v25)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L3
	} else {
		goto L45
	}
L2:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
	v168 = F_get_namespace_name_or_temp(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L3
	} else {
		goto L40
	}
L3:
	;
	return int32(0)
L4:
	;
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
	v23 = v21 + v22
	v25 = v23 + int32(4)
	if l1 == int32(0) {
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
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L37
	}
L8:
	;
	v135 = F_RelationIsVisible(m, l0)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L3
	} else {
		goto L34
	}
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v31 = int32(0)
	if v31 < v28 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = v28
	goto L13
L12:
	;
	v34 = v31
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v42 = int32(0)
	goto L14
L14:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v35+v42<<(uint(int32(2))%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
	if v51 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L8
L16:
	;
	v122 = v42 + int32(1)
	if v122 != v34 {
		v42 = v122
		goto L14
	} else {
		goto L33
	}
L17:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)+4))
	if v54 <= int32(0) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v57 = int32(0)
	if v57 < v54 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v60 = v54
	goto L21
L20:
	;
	v60 = v57
	goto L21
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v51)+12))
	v64 = int32(0)
	goto L22
L22:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v61+v64<<(uint(int32(2))%32))))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if base.B2i32(v81 == int32(0))|base.B2i32(v81 != v84) != 0 {
		v102 = v81
		v103 = v84
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L16
L24:
	;
	if v102-v103 == int32(0) {
		goto L2
	} else {
		goto L31
	}
L25:
	;
	goto L24
L26:
	;
	v87 = v78
	v88 = v25
	goto L27
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88)+1)))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	if v92 == int32(0) {
		v102 = v92
		v103 = v91
		goto L25
	} else {
		goto L29
	}
L28:
	;
	v102 = v92
	v103 = v91
	goto L25
L29:
	;
	v95 = int32(1)
	if v92 == v91 {
		v87 = v87 + v95
		v88 = v88 + v95
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v108 = v64 + int32(1)
	if v108 != v60 {
		v64 = v108
		goto L22
	} else {
		goto L32
	}
L32:
	;
	goto L23
L33:
	;
	goto L15
L34:
	;
	if v135 == int32(0) {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	F_initStringInfo(m, v14+int32(32))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L3
	} else {
		goto L36
	}
L36:
	;
	goto L1
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = l0
	F_errmsg_internal(m, int32(_a_F_generate_relation_name_0), v14)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_generate_relation_name_1), int32(_a_F_generate_relation_name_2), int32(_a_F_generate_relation_name_3))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	v171 = v14 + int32(32)
	F_initStringInfo(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	if v168 == int32(0) {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v176 = F_quote_identifier(m, v168)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L3
	} else {
		goto L43
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v176
	F_appendStringInfo(m, v171, int32(_a_F_generate_relation_name_4), v14+int32(16))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L3
	} else {
		goto L44
	}
L44:
	;
	goto L1
L45:
	;
	F_appendStringInfoString(m, v14+int32(32), v197)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L3
	} else {
		goto L46
	}
L46:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
	F_ReleaseCatCache(m, v17)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	m.G0 = v14 + int32(48)
	return v201
}
func F_getRelationDescription(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			if l2 != 0 {
				m.G0 = v9 + int32(32)
				return
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					F_errmsg_internal(m, int32(_a_F_getRelationDescription_0), v9)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_getRelationDescription_1), int32(_a_F_getRelationDescription_2), int32(_a_F_getRelationDescription_3))
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+22)))
			v31 = v29 + v30
			v32 = F_RelationIsVisible(m, l1)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				if v32 != 0 {
					v38 = int32(0)
					v41 = F_quote_qualified_identifier(m, v38, v31+int32(4))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
						switch v44 - int32(73) {
						case 0, 32:
							v55 = int32(_a_F_getRelationDescription_4)
						default:
							v55 = int32(_a_F_getRelationDescription_5)
						case 10:
							v55 = int32(_a_F_getRelationDescription_6)
						case 26:
							v55 = int32(_a_F_getRelationDescription_7)
						case 29:
							v55 = int32(_a_F_getRelationDescription_8)
						case 36:
							v55 = int32(_a_F_getRelationDescription_9)
						case 39, 41:
							v55 = int32(_a_F_getRelationDescription_10)
						case 43:
							v55 = int32(_a_F_getRelationDescription_11)
						case 45:
							v55 = int32(_a_F_getRelationDescription_12)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v41
						F_appendStringInfo(m, l0, v55, v9+int32(16))
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return
						} else {
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				} else {
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v31)+68))
					v36 = F_get_namespace_name(m, v35)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						v38 = v36
						v41 = F_quote_qualified_identifier(m, v38, v31+int32(4))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+119)))
							switch v44 - int32(73) {
							case 0, 32:
								v55 = int32(_a_F_getRelationDescription_4)
							default:
								v55 = int32(_a_F_getRelationDescription_5)
							case 10:
								v55 = int32(_a_F_getRelationDescription_6)
							case 26:
								v55 = int32(_a_F_getRelationDescription_7)
							case 29:
								v55 = int32(_a_F_getRelationDescription_8)
							case 36:
								v55 = int32(_a_F_getRelationDescription_9)
							case 39, 41:
								v55 = int32(_a_F_getRelationDescription_10)
							case 43:
								v55 = int32(_a_F_getRelationDescription_11)
							case 45:
								v55 = int32(_a_F_getRelationDescription_12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v41
							F_appendStringInfo(m, l0, v55, v9+int32(16))
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return
							} else {
								F_ReleaseCatCache(m, v12)
								mBase = m.M
								v62 = m.ExcPending
								if v62 != 0 {
									return
								} else {
									m.G0 = v9 + int32(32)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_getRelationIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = F_SearchSysCache1(m, int32(57), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		if v12 == int32(0) {
			if l3 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = l1
					F_errmsg_internal(m, int32(_a_F_getRelationIdentity_0), v9)
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_getRelationIdentity_1), int32(_a_F_getRelationIdentity_2), int32(_a_F_getRelationIdentity_3))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				if l2 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(0)
				}
				m.G0 = v9 + int32(32)
				return
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)))
			v24 = v22 + v23
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
			v26 = F_get_namespace_name_or_temp(m, v25)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return
			} else {
				v29 = v24 + int32(4)
				v30 = F_quote_qualified_identifier(m, v26, v29)
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					F_appendStringInfoString(m, l0, v30)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						if l2 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v26
							v35 = F_pstrdup(m, v29)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v35
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v35
								v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v39
								v45 = F_list_make2_impl(m, v9+int32(20), v9+int32(16))
								mBase = m.M
								v46 = m.ExcPending
								if v46 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l2))) = v45
									F_ReleaseCatCache(m, v12)
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return
									} else {
										m.G0 = v9 + int32(32)
										return
									}
								}
							}
						} else {
							F_ReleaseCatCache(m, v12)
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_get_relation_constraint_oid(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v9 = m.G0
	v11 = v9 - int32(160)
	m.G0 = v11
	v15 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v20 = v11 + int32(16)
		F_ScanKeyInit(m, v20, int32(9), int32(3), int32(184), l0)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			F_ScanKeyInit(m, v11-int32(-64), int32(10), int32(3), int32(184), int32(0))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				F_ScanKeyInit(m, v11+int32(112), int32(2), int32(3), int32(62), l1)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int32(0)
				} else {
					v45 = F_systable_beginscan(m, v15, int32(2665), int32(1), int32(0), int32(3), v20)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						v47 = F_systable_getnext(m, v45)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							if v47 != 0 {
								v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+22)))
								v52 = *(*int32)(unsafe.Add(mBase, uint32(v49+v50)))
								v53 = v52
							} else {
								v53 = int32(0)
							}
							F_systable_endscan(m, v45)
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								if l2|v53 == int32(0) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v62 = m.ExcPending
									if v62 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(67137668))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return int32(0)
										} else {
											v66 = F_get_rel_name(m, l0)
											mBase = m.M
											v67 = m.ExcPending
											if v67 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v66
												*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
												F_errmsg(m, int32(_a_F_get_relation_constraint_oid_0), v11)
												mBase = m.M
												v72 = m.ExcPending
												if v72 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_get_relation_constraint_oid_1), int32(1235), int32(_a_F_get_relation_constraint_oid_2))
													mBase = m.M
													v77 = m.ExcPending
													if v77 != 0 {
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
								} else {
									F_relation_close(m, v15, int32(1))
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return int32(0)
									} else {
										m.G0 = v11 + int32(160)
										return v53
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_relation_excluded_by_constraints(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
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
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v382 int32
	_ = v382
	v4 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	if v13 == v4 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if int32(0) < v18 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v382
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v25 = v4
	goto L8
L6:
	;
	goto L7
L7:
	;
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_relation_excluded_by_constraints[0]))
	switch v67 {
	case 0:
		v382 = int32(0)
		goto L4
	case 1:
		goto L17
	case 2:
		goto L18
	default:
		v76 = v4
		goto L16
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21+v25<<(uint(int32(2))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v51 = v25 + int32(1)
	if v51 != v18 {
		v25 = v51
		goto L8
	} else {
		goto L15
	}
L11:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v41 != int32(7) {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v44 = int32(1)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+24)))
	if v45 != 0 {
		v382 = v44
		goto L4
	} else {
		goto L13
	}
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	if v46 == int32(0) {
		v382 = v44
		goto L4
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	goto L9
L16:
	;
	v78 = int32(0)
	if v78 < v18 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v76 = base.B2i32(v73 == int32(0))
	goto L16
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v68 == int32(2) {
		v76 = v4
		goto L16
	} else {
		goto L19
	}
L19:
	;
	return int32(0)
L20:
	;
	v85 = int32(0)
	v86 = v78
	goto L23
L21:
	;
	v118 = v78
	goto L22
L22:
	;
	v127 = F_predicate_refuted_by(m, v118, v118, int32(1))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L25
	} else {
		goto L32
	}
L23:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v85<<(uint(int32(2))%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v100 = F_contain_mutable_functions(m, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v118 = v109
	goto L22
L25:
	;
	return int32(0)
L26:
	;
	if v100 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v107 = F_lappend(m, v86, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L25
	} else {
		goto L30
	}
L28:
	;
	v109 = v86
	goto L29
L29:
	;
	v111 = v85 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v111 < v112 {
		v85 = v111
		v86 = v109
		goto L23
	} else {
		goto L31
	}
L30:
	;
	v109 = v107
	goto L29
L31:
	;
	goto L24
L32:
	;
	if v127 != 0 {
		v382 = int32(1)
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v130 != 0 {
		v382 = int32(0)
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v131 = int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+20)))
	if v132 == v131 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+21)))
	v138 = base.B2i32(v135 == int32(112))
	goto L37
L36:
	;
	v138 = v131
	goto L37
L37:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v140 = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v143 = F_table_open(m, v141, v140)
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L25
	} else {
		goto L39
	}
L38:
	;
	if v76 == int32(0) {
		v312 = v277
		goto L71
	} else {
		goto L72
	}
L39:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)+16))
	if v146 == int32(0) {
		v277 = v140
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+14)))
	if v149 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v154 = int32(0)
	v155 = v140
	goto L44
L42:
	;
	v198 = v140
	goto L43
L43:
	;
	if v138 == int32(0) {
		v277 = v198
		goto L38
	} else {
		goto L59
	}
L44:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v166 = v163 + v154*int32(12)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+9)))
	if v167 != int32(1) {
		v189 = v155
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v198 = v189
	goto L43
L46:
	;
	v192 = v154 + int32(1)
	if v192 != v149 {
		v154 = v192
		v155 = v189
		goto L44
	} else {
		goto L58
	}
L47:
	;
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v166)+10)))
	if v170&v132 != 0 {
		v189 = v155
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v173 = F_stringToNode(m, v172)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L25
	} else {
		goto L49
	}
L49:
	;
	v175 = F_eval_const_expressions(m, l0, v173)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L25
	} else {
		goto L50
	}
L50:
	;
	v178 = F_canonicalize_qual(m, v175, int32(1))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L25
	} else {
		goto L51
	}
L51:
	;
	if v139 != int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_ChangeVarNodes(m, v178, int32(1), v139)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L25
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v185 = F_make_ands_implicit(m, v178)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L25
	} else {
		goto L56
	}
L55:
	;
	goto L54
L56:
	;
	v187 = F_list_concat(m, v155, v185)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L25
	} else {
		goto L57
	}
L57:
	;
	v189 = v187
	goto L46
L58:
	;
	goto L45
L59:
	;
	v208 = int32(1)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146)+16)))
	if v209 != v208 {
		v277 = v198
		goto L38
	} else {
		goto L60
	}
L60:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	if v213 <= int32(0) {
		v277 = v198
		goto L38
	} else {
		goto L61
	}
L61:
	;
	v219 = v208
	v220 = v198
	goto L62
L62:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v143)+52))
	v230 = v219 - int32(1)
	v233 = v228 + v230<<(uint(int32(4))%32)
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+31)))
	if v234 != int32(118) {
		v267 = v220
		goto L64
	} else {
		goto L65
	}
L63:
	;
	v277 = v267
	goto L38
L64:
	;
	v271 = v219 + int32(1)
	if v271 <= v213 {
		v219 = v271
		v220 = v267
		goto L62
	} else {
		goto L70
	}
L65:
	;
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233)+29)))
	if v237 != 0 {
		v267 = v220
		goto L64
	} else {
		goto L66
	}
L66:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v228)))
	v240 = F_palloc0(m, int32(20))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L25
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240))) = int32(52)
	v250 = v228 + v238<<(uint(int32(4))%32) + v230*int32(100)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v250)+88))
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v250)+96))
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v250)+116))
	v255 = F_makeVar(m, v139, base.I32_extend16_s(v219), v251, v252, v253, int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L25
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v240)+16)) = int32(-1)
	v259 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v240)+12)) = uint8(v259)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v240)+4)) = v255
	v264 = F_lappend(m, v220, v240)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L25
	} else {
		goto L69
	}
L69:
	;
	v267 = v264
	goto L64
L70:
	;
	goto L63
L71:
	;
	if v312 == int32(0) {
		goto L87
	} else {
		goto L88
	}
L72:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v143)+48))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+131)))
	if v288 != int32(1) {
		v312 = v277
		goto L71
	} else {
		goto L73
	}
L73:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	if v291 != 0 {
		v307 = v291
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v308 = F_list_concat(m, v277, v307)
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L25
	} else {
		goto L85
	}
L75:
	;
	v292 = F_RelationGetPartitionQual(m, v143)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L25
	} else {
		goto L76
	}
L76:
	;
	if v292 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l1)+248))
	v307 = v296
	goto L74
L78:
	;
	goto L79
L79:
	;
	v297 = F_expression_planner(m, v292)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L25
	} else {
		goto L80
	}
L80:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	if v299 != int32(1) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	F_ChangeVarNodes(m, v297, int32(1), v299)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L25
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+248)) = v297
	v307 = v297
	goto L74
L84:
	;
	goto L83
L85:
	;
	v312 = v308
	goto L71
L86:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l1)+184))
	v373 = F_predicate_refuted_by(m, v359, v371, int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L25
	} else {
		goto L103
	}
L87:
	;
	v315 = int32(0)
	F_relation_close(m, v143, v315)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L25
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v319 = int32(0)
	v320 = F_expand_generated_columns_in_expr(m, v312, v143, v139)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L25
	} else {
		goto L91
	}
L90:
	;
	v359 = v315
	goto L86
L91:
	;
	F_relation_close(m, v143, int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L25
	} else {
		goto L92
	}
L92:
	;
	if v320 == int32(0) {
		v359 = v319
		goto L86
	} else {
		goto L93
	}
L93:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v327 <= int32(0) {
		v359 = v319
		goto L86
	} else {
		goto L94
	}
L94:
	;
	v331 = v319
	v334 = int32(0)
	goto L95
L95:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v320)+12))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v343+v334<<(uint(int32(2))%32))))
	v348 = F_contain_mutable_functions(m, v347)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L25
	} else {
		goto L97
	}
L96:
	;
	v359 = v354
	goto L86
L97:
	;
	if v348 == int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v352 = F_lappend(m, v331, v347)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L25
	} else {
		goto L101
	}
L99:
	;
	v354 = v331
	goto L100
L100:
	;
	v356 = v334 + int32(1)
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v320)+4))
	if v356 < v357 {
		v331 = v354
		v334 = v356
		goto L95
	} else {
		goto L102
	}
L101:
	;
	v354 = v352
	goto L100
L102:
	;
	goto L96
L103:
	;
	v382 = v373
	goto L4
}
func F_relation_is_updatable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v391 int32
	_ = v391
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v415 int32
	_ = v415
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v525 int32
	_ = v525
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	F_check_stack_depth(m)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = F_try_relation_open(m, l0, int32(1))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v17 + int32(16)
	return v542
L4:
	;
	if v24 == int32(0) {
		v542 = v5
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v29 = int32(0)
	if l1 == v29 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v67 = int32(0)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v35 <= int32(0) {
		v61 = v29
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v67 = v61
	goto L6
L11:
	;
	v38 = int32(0)
	if v38 < v35 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = v35
	goto L14
L13:
	;
	v41 = v38
	goto L14
L14:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v44 = int32(0)
	goto L15
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v42+v44<<(uint(int32(2))%32))))
	v53 = base.B2i32(v52 == v28)
	if v52 == v28 {
		v61 = v53
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v61 = v53
	goto L10
L17:
	;
	v55 = v44 + int32(1)
	if v55 != v41 {
		v44 = v55
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+119)))
	switch v72 - int32(112) {
	case 0, 2:
		goto L24
	default:
		goto L23
	}
L22:
	;
	v542 = v5
	goto L3
L23:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	if v79 == int32(0) {
		v194 = v5
		goto L26
	} else {
		goto L27
	}
L24:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v542 = int32(28)
	goto L3
L26:
	;
	if l2 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L27:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	if v82 <= int32(0) {
		v194 = v5
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v82 != int32(1) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v182 = int32(28)
	if v174 != v182 {
		v194 = v174
		goto L26
	} else {
		goto L45
	}
L30:
	;
	v96 = v5
	v98 = v5
	v100 = v5
	goto L33
L31:
	;
	v144 = v5
	v146 = v5
	goto L32
L32:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v85+v144<<(uint(int32(2))%32))))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+17)))
	if v158 != int32(1) {
		v174 = v146
		goto L29
	} else {
		goto L43
	}
L33:
	;
	v108 = v85 + v96<<(uint(int32(2))%32)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+17)))
	if v110 != int32(1) {
		v120 = v98
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v82&int32(1) == int32(0) {
		v174 = v132
		goto L29
	} else {
		goto L42
	}
L35:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+17)))
	if v122 != int32(1) {
		v132 = v120
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	if v113 != 0 {
		v120 = v98
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+4))
	v120 = int32(1)<<(uint(v115)%32)&int32(28) | v98
	goto L35
L38:
	;
	v133 = int32(2)
	v134 = v96 + v133
	v136 = v100 + v133
	if v136 != v82&int32(2147483646) {
		v96 = v134
		v98 = v132
		v100 = v136
		goto L33
	} else {
		goto L41
	}
L39:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v125 != 0 {
		v132 = v120
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)+4))
	v132 = int32(1)<<(uint(v127)%32)&int32(28) | v120
	goto L38
L41:
	;
	goto L34
L42:
	;
	v144 = v134
	v146 = v132
	goto L32
L43:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
	if v161 != 0 {
		v174 = v146
		goto L29
	} else {
		goto L44
	}
L44:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v174 = int32(1)<<(uint(v163)%32)&int32(28) | v146
	goto L29
L45:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v542 = v182
	goto L3
L47:
	;
	v231 = v72 - int32(102)
	if v231 != 0 {
		goto L55
	} else {
		goto L56
	}
L48:
	;
	v229 = v194
	goto L47
L49:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v24)+76))
	if v204 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+15)))
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+10)))
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+20)))
	v218 = int32(28)
	v220 = v194 | (v207<<(uint(int32(2))%32)|v210<<(uint(int32(3))%32)|v214<<(uint(int32(4))%32))&v218
	if v220 != v218 {
		v229 = v220
		goto L47
	} else {
		goto L51
	}
L51:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v542 = int32(28)
	goto L3
L53:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L169
	}
L54:
	;
	v258 = F_get_view_query(m, v24)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L77
	}
L55:
	;
	if v231 == int32(16) {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	goto L57
L57:
	;
	v235 = F_GetFdwRoutineForRelation(m, v24, int32(0))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L62
	}
L58:
	;
	goto L54
L59:
	;
	v525 = v229
	goto L53
L61:
	;
	F_relation_close(m, v24, int32(1))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L76
	}
L62:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v235)+84))
	if v237 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v238 = m.T0[v237].(func(*base.Module, int32) int32)(m, v24)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v235)+52))
	if v243 != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v254 = v238 | v229
	goto L61
L67:
	;
	v244 = v229 | int32(8)
	goto L69
L68:
	;
	v244 = v229
	goto L69
L69:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v235)+64))
	if v247 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v248 = v244 | int32(4)
	goto L72
L71:
	;
	v248 = v244
	goto L72
L72:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v235)+68))
	if v251 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v252 = v248 | int32(16)
	goto L75
L74:
	;
	v252 = v248
	goto L75
L75:
	;
	v254 = v252
	goto L61
L76:
	;
	v542 = v254
	goto L3
L77:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v258)+120))
	if v266 != 0 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	if v402 != 0 {
		v525 = v229
		goto L53
	} else {
		goto L135
	}
L79:
	;
	v402 = int32(_a_F_relation_is_updatable_0)
	goto L78
L80:
	;
	goto L81
L81:
	;
	v268 = int32(_a_F_relation_is_updatable_1)
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v258)+100))
	if v269 != 0 {
		v391 = v268
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v402 = v391
	goto L78
L83:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v258)+108))
	if v270 != 0 {
		v391 = v268
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v258)+112))
	if v271 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v402 = int32(_a_F_relation_is_updatable_2)
	goto L78
L86:
	;
	goto L87
L87:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v258)+144))
	if v273 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v402 = int32(_a_F_relation_is_updatable_3)
	goto L78
L89:
	;
	goto L90
L90:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v258)+48))
	if v275 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v402 = int32(_a_F_relation_is_updatable_4)
	goto L78
L92:
	;
	goto L93
L93:
	;
	v277 = int32(_a_F_relation_is_updatable_5)
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v258)+128))
	if v278 != 0 {
		v391 = v277
		goto L82
	} else {
		goto L94
	}
L94:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v258)+132))
	if v279 != 0 {
		v391 = v277
		goto L82
	} else {
		goto L95
	}
L95:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+36)))
	if v280 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v402 = int32(_a_F_relation_is_updatable_6)
	goto L78
L97:
	;
	goto L98
L98:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+37)))
	if v282 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v402 = int32(_a_F_relation_is_updatable_7)
	goto L78
L100:
	;
	goto L101
L101:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258)+38)))
	if v284 != 0 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v402 = int32(_a_F_relation_is_updatable_8)
	goto L78
L103:
	;
	goto L104
L104:
	;
	v286 = int32(_a_F_relation_is_updatable_9)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v258)+60))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+4))
	if v288 == int32(0) {
		v391 = v286
		goto L82
	} else {
		goto L105
	}
L105:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v291 != int32(1) {
		v391 = v286
		goto L82
	} else {
		goto L106
	}
L106:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v288)+12))
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v294)))
	v296 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v296 != int32(63) {
		v391 = v286
		goto L82
	} else {
		goto L107
	}
L107:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v258)+52))
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v299)+12))
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v300+v301<<(uint(int32(2))%32)-int32(4))))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+12))
	if v308 != 0 {
		v391 = v286
		goto L82
	} else {
		goto L108
	}
L108:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307)+21)))
	v311 = v309 - int32(102)
	v316 = int32(1)
	v320 = (v311<<(uint(int32(7))%32) | int32(base.Ui32(v311&int32(254))>>(uint(v316)%32))) & int32(255)
	if base.B2i32(base.Ui32(int32(8)) < base.Ui32(v320))|base.B2i32(v316<<(uint(v320)%32)&int32(353) == int32(0)) != 0 {
		v391 = v286
		goto L82
	} else {
		goto L109
	}
L109:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v307)+32))
	if v332 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v333 = int32(_a_F_relation_is_updatable_10)
	goto L112
L111:
	;
	v333 = int32(0)
	goto L112
L112:
	;
	if int32(1)|v332 != 0 {
		v391 = v333
		goto L82
	} else {
		goto L113
	}
L113:
	;
	v337 = int32(_a_F_relation_is_updatable_11)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v258)+76))
	if v338 == int32(0) {
		v391 = v337
		goto L82
	} else {
		goto L114
	}
L114:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	if v341 <= int32(0) {
		v391 = v337
		goto L82
	} else {
		goto L115
	}
L115:
	;
	v344 = int32(0)
	if v344 < v341 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v348 = v341
	goto L118
L117:
	;
	v348 = v344
	goto L118
L118:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v338)+12))
	v350 = v344
	goto L119
L119:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v349+v350<<(uint(int32(2))%32))))
	v362 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v361)+26)))
	if v362 != 0 {
		v383 = int32(_a_F_relation_is_updatable_12)
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v391 = int32(0)
	goto L82
L121:
	;
	if v383 != 0 {
		goto L131
	} else {
		goto L132
	}
L122:
	;
	v363 = int32(_a_F_relation_is_updatable_13)
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v361)+4))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v364)))
	if v365 != int32(6) {
		v380 = v363
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v383 = v380
	goto L121
L124:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364)+4))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v295)+4))
	if v368 != v369 {
		v380 = v363
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v364)+28))
	if v371 != 0 {
		v380 = v363
		goto L123
	} else {
		goto L126
	}
L126:
	;
	v373 = int32(*(*int16)(unsafe.Add(mBase, uint32(v364)+8)))
	if v373 < int32(0) {
		v383 = int32(_a_F_relation_is_updatable_14)
		goto L121
	} else {
		goto L127
	}
L127:
	;
	if v373 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v378 = int32(0)
	goto L130
L129:
	;
	v378 = int32(_a_F_relation_is_updatable_15)
	goto L130
L130:
	;
	v380 = v378
	goto L123
L131:
	;
	v385 = v350 + int32(1)
	if v348 != v385 {
		v350 = v385
		goto L119
	} else {
		goto L134
	}
L132:
	;
	goto L133
L133:
	;
	goto L120
L134:
	;
	v391 = v337
	goto L82
L135:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v258)+60))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v404)+4))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v405)+12))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v406)))
	v409 = v17 + int32(12)
	if v409 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409))) = int32(0)
	goto L138
L137:
	;
	goto L138
L138:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v258)+76))
	if v412 == int32(0) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	if l3 != 0 {
		goto L156
	} else {
		goto L157
	}
L140:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v415 <= int32(0) {
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v425 = int32(7)
	v428 = int32(0)
	goto L142
L142:
	;
	v434 = v425 + int32(1)
	v435 = *(*int32)(unsafe.Add(mBase, uint32(v412)+12))
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v435+v428<<(uint(int32(2))%32))))
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v439)+26)))
	if v440 != 0 {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	goto L139
L144:
	;
	v466 = v428 + int32(1)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v412)+4))
	if v466 < v467 {
		v425 = v434
		v428 = v466
		goto L142
	} else {
		goto L155
	}
L145:
	;
	v462 = F_bms_is_member(m, base.I32_extend16_s(v434), int32(0))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L153
	}
L146:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v441)))
	if v442 != int32(6) {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v441)+4))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v407)+4))
	if v445 != v446 {
		goto L145
	} else {
		goto L148
	}
L148:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v441)+28))
	if v448 != 0 {
		goto L145
	} else {
		goto L149
	}
L149:
	;
	v449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v441)+8)))
	if v449 <= int32(0) {
		goto L145
	} else {
		goto L150
	}
L150:
	;
	if v409 == int32(0) {
		goto L144
	} else {
		goto L151
	}
L151:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v409)))
	v456 = F_bms_add_member(m, v454, base.I32_extend16_s(v434))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v409))) = v456
	goto L144
L153:
	;
	if v462 != 0 {
		goto L139
	} else {
		goto L154
	}
L154:
	;
	goto L144
L155:
	;
	goto L143
L156:
	;
	v484 = F_bms_int_members(m, v483, l3)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L1
	} else {
		goto L159
	}
L157:
	;
	v486 = v483
	goto L158
L158:
	;
	if v486 != 0 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	v486 = v484
	goto L158
L160:
	;
	v489 = int32(28)
	goto L162
L161:
	;
	v489 = int32(16)
	goto L162
L162:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v258)+52))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v490)+12))
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v258)+60))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+12))
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v494)))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v495)+4))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v491+v496<<(uint(int32(2))%32)-int32(4))))
	v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v502)+21)))
	switch v503 - int32(112) {
	case 0, 2:
		v519 = v489
		goto L163
	default:
		goto L164
	}
L163:
	;
	v525 = v519 | v229
	goto L53
L164:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v502)+16))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v24)+56))
	v508 = F_lappend_oid(m, l1, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L1
	} else {
		goto L165
	}
L165:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v258)+76))
	v511 = F_adjust_view_column_set(m, v486, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L1
	} else {
		goto L166
	}
L166:
	;
	v513 = F_relation_is_updatable(m, v506, v508, l2, v511)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v515 = F_list_delete_last(m, v508)
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L168
	}
L168:
	;
	v519 = v513 & v489
	goto L163
L169:
	;
	v542 = v525
	goto L3
}
func F_relation_statistics_update(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int64
	_ = v21
	var v38 int32
	_ = v38
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 float32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 float32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 float32
	_ = v153
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 + int32(-64)
	m.G0 = v19
	v21 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v19)+40)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v19)+32)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v19)+28)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = v2
	F_stats_check_required_arg(m, l0, int32(_a_F_relation_statistics_update_0), v2)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return int32(0)
	} else {
		F_stats_check_required_arg(m, l0, int32(_a_F_relation_statistics_update_0), int32(1))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v44 = F_text_to_cstring(m, v43)
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int32(0)
			} else {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v47 = F_text_to_cstring(m, v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_relation_statistics_update[0])))
					if v51 == int32(1) {
						v56 = *(*int32)(unsafe.Add(mBase, _c_F_relation_statistics_update[1]))
						v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+316))
						v59 = base.B2i32(v57 != int32(2))
						*(*uint8)(unsafe.Add(mBase, _c_F_relation_statistics_update[0])) = uint8(v59)
						v61 = v59
					} else {
						v61 = int32(0)
					}
					if v61 == int32(0) {
						v65 = F_makeRangeVar(m, v44, v47, int32(-1))
						mBase = m.M
						v66 = m.ExcPending
						if v66 != 0 {
							return int32(0)
						} else {
							v72 = F_RangeVarGetRelidExtended(m, v65, int32(4), int32(0), int32(1060), v17+int32(-40))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+40)))
								if v74 == int32(0) {
									v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									v78 = v77
								} else {
									v78 = v2
								}
								v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+48)))
								if v80 != 0 {
									v108 = v2
									v109 = float32(0)
									v110 = int32(1)
									v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
									if v111 == int32(0) {
										v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
										v115 = v114
									} else {
										v115 = v2
									}
									v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
									if v116 == int32(0) {
										v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
										v120 = v119
									} else {
										v120 = v2
									}
									v123 = F_table_open(m, int32(1259), int32(3))
									mBase = m.M
									v124 = m.ExcPending
									if v124 != 0 {
										return int32(0)
									} else {
										v126 = F_SearchSysCache1(m, int32(57), v72)
										mBase = m.M
										v127 = m.ExcPending
										if v127 != 0 {
											return int32(0)
										} else {
											if v126 == int32(0) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v248 = m.ExcPending
												if v248 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v19))) = v72
													F_errmsg_internal(m, int32(_a_F_relation_statistics_update_1), v19)
													mBase = m.M
													v252 = m.ExcPending
													if v252 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_relation_statistics_update_2), int32(144), int32(_a_F_relation_statistics_update_3))
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return int32(0)
														} else {
															base.Wasm_trap_unreachable()
															for {
															}
														}
													}
												}
											} else {
												v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
												v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+22)))
												v132 = v130 + v131
												v133 = int32(0)
												v135 = v17 + int32(-16)
												v137 = v17 + int32(-32)
												if v74 != 0 {
													v148 = v133
													v149 = v135
													v150 = v137
												} else {
													v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+96))
													if v78 == v138 {
														v148 = v133
														v149 = v135
														v150 = v137
													} else {
														v140 = int32(4)
														*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78
														*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(10)
														v148 = int32(1)
														v149 = v135 | v140
														v150 = v137 | v140
													}
												}
												if v108 == int32(0) {
													v160 = v148
												} else {
													v153 = *(*float32)(unsafe.Add(mBase, uint32(v132)+100))
													if base.F32_eq(v109, v153) != 0 {
														v160 = v148
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(11)
														*(*float32)(unsafe.Add(mBase, uint32(v150))) = v109
														v160 = v148 + int32(1)
													}
												}
												if v111 != 0 {
													v176 = v160
												} else {
													v161 = *(*int32)(unsafe.Add(mBase, uint32(v132)+104))
													if v115 == v161 {
														v176 = v160
													} else {
														v164 = v160 << (uint(int32(2)) % 32)
														*(*int32)(unsafe.Add(mBase, uint32(v164|(v17+int32(-32))))) = v115
														*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)|v164))) = int32(12)
														v176 = v160 + int32(1)
													}
												}
												if v116 != 0 {
													if v176 == int32(0) {
														F_ReleaseCatCache(m, v126)
														mBase = m.M
														v215 = m.ExcPending
														if v215 != 0 {
															return int32(0)
														} else {
															F_relation_close(m, v123, int32(3))
															mBase = m.M
															v218 = m.ExcPending
															if v218 != 0 {
																return int32(0)
															} else {
																F_CommandCounterIncrement(m)
																mBase = m.M
																v220 = m.ExcPending
																if v220 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v19 - int32(-64)
																	return v110
																}
															}
														}
													} else {
														v195 = v176
														v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
														v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v204)
																mBase = m.M
																v211 = m.ExcPending
																if v211 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v126)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v123, int32(3))
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v220 = m.ExcPending
																			if v220 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v19 - int32(-64)
																				return v110
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													v178 = *(*int32)(unsafe.Add(mBase, uint32(v132)+108))
													if v120 == v178 {
														if v176 == int32(0) {
															F_ReleaseCatCache(m, v126)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v123, int32(3))
																mBase = m.M
																v218 = m.ExcPending
																if v218 != 0 {
																	return int32(0)
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v220 = m.ExcPending
																	if v220 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v19 - int32(-64)
																		return v110
																	}
																}
															}
														} else {
															v195 = v176
															v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
															v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return int32(0)
															} else {
																F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v204)
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v126)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v123, int32(3))
																			mBase = m.M
																			v218 = m.ExcPending
																			if v218 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v220 = m.ExcPending
																				if v220 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v19 - int32(-64)
																					return v110
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v181 = v176 << (uint(int32(2)) % 32)
														*(*int32)(unsafe.Add(mBase, uint32(v181+(v17+int32(-32))))) = v120
														*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)+v181))) = int32(13)
														v195 = v176 + int32(1)
														v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
														v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return int32(0)
														} else {
															F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int32(0)
															} else {
																F_pfree(m, v204)
																mBase = m.M
																v211 = m.ExcPending
																if v211 != 0 {
																	return int32(0)
																} else {
																	F_ReleaseCatCache(m, v126)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v123, int32(3))
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v220 = m.ExcPending
																			if v220 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v19 - int32(-64)
																				return v110
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								} else {
									v81 = *(*float32)(unsafe.Add(mBase, uint32(l0)+44))
									if base.F32_lt(v81, float32(-1)) == int32(0) {
										v86 = int32(1)
										v108 = v86
										v109 = v81
										v110 = v86
										v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
										if v111 == int32(0) {
											v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
											v115 = v114
										} else {
											v115 = v2
										}
										v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
										if v116 == int32(0) {
											v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
											v120 = v119
										} else {
											v120 = v2
										}
										v123 = F_table_open(m, int32(1259), int32(3))
										mBase = m.M
										v124 = m.ExcPending
										if v124 != 0 {
											return int32(0)
										} else {
											v126 = F_SearchSysCache1(m, int32(57), v72)
											mBase = m.M
											v127 = m.ExcPending
											if v127 != 0 {
												return int32(0)
											} else {
												if v126 == int32(0) {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v248 = m.ExcPending
													if v248 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v19))) = v72
														F_errmsg_internal(m, int32(_a_F_relation_statistics_update_1), v19)
														mBase = m.M
														v252 = m.ExcPending
														if v252 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_relation_statistics_update_2), int32(144), int32(_a_F_relation_statistics_update_3))
															mBase = m.M
															v257 = m.ExcPending
															if v257 != 0 {
																return int32(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												} else {
													v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
													v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+22)))
													v132 = v130 + v131
													v133 = int32(0)
													v135 = v17 + int32(-16)
													v137 = v17 + int32(-32)
													if v74 != 0 {
														v148 = v133
														v149 = v135
														v150 = v137
													} else {
														v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+96))
														if v78 == v138 {
															v148 = v133
															v149 = v135
															v150 = v137
														} else {
															v140 = int32(4)
															*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78
															*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(10)
															v148 = int32(1)
															v149 = v135 | v140
															v150 = v137 | v140
														}
													}
													if v108 == int32(0) {
														v160 = v148
													} else {
														v153 = *(*float32)(unsafe.Add(mBase, uint32(v132)+100))
														if base.F32_eq(v109, v153) != 0 {
															v160 = v148
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(11)
															*(*float32)(unsafe.Add(mBase, uint32(v150))) = v109
															v160 = v148 + int32(1)
														}
													}
													if v111 != 0 {
														v176 = v160
													} else {
														v161 = *(*int32)(unsafe.Add(mBase, uint32(v132)+104))
														if v115 == v161 {
															v176 = v160
														} else {
															v164 = v160 << (uint(int32(2)) % 32)
															*(*int32)(unsafe.Add(mBase, uint32(v164|(v17+int32(-32))))) = v115
															*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)|v164))) = int32(12)
															v176 = v160 + int32(1)
														}
													}
													if v116 != 0 {
														if v176 == int32(0) {
															F_ReleaseCatCache(m, v126)
															mBase = m.M
															v215 = m.ExcPending
															if v215 != 0 {
																return int32(0)
															} else {
																F_relation_close(m, v123, int32(3))
																mBase = m.M
																v218 = m.ExcPending
																if v218 != 0 {
																	return int32(0)
																} else {
																	F_CommandCounterIncrement(m)
																	mBase = m.M
																	v220 = m.ExcPending
																	if v220 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v19 - int32(-64)
																		return v110
																	}
																}
															}
														} else {
															v195 = v176
															v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
															v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return int32(0)
															} else {
																F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v204)
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v126)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v123, int32(3))
																			mBase = m.M
																			v218 = m.ExcPending
																			if v218 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v220 = m.ExcPending
																				if v220 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v19 - int32(-64)
																					return v110
																				}
																			}
																		}
																	}
																}
															}
														}
													} else {
														v178 = *(*int32)(unsafe.Add(mBase, uint32(v132)+108))
														if v120 == v178 {
															if v176 == int32(0) {
																F_ReleaseCatCache(m, v126)
																mBase = m.M
																v215 = m.ExcPending
																if v215 != 0 {
																	return int32(0)
																} else {
																	F_relation_close(m, v123, int32(3))
																	mBase = m.M
																	v218 = m.ExcPending
																	if v218 != 0 {
																		return int32(0)
																	} else {
																		F_CommandCounterIncrement(m)
																		mBase = m.M
																		v220 = m.ExcPending
																		if v220 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v19 - int32(-64)
																			return v110
																		}
																	}
																}
															} else {
																v195 = v176
																v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
																v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																mBase = m.M
																v205 = m.ExcPending
																if v205 != 0 {
																	return int32(0)
																} else {
																	F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																	mBase = m.M
																	v209 = m.ExcPending
																	if v209 != 0 {
																		return int32(0)
																	} else {
																		F_pfree(m, v204)
																		mBase = m.M
																		v211 = m.ExcPending
																		if v211 != 0 {
																			return int32(0)
																		} else {
																			F_ReleaseCatCache(m, v126)
																			mBase = m.M
																			v215 = m.ExcPending
																			if v215 != 0 {
																				return int32(0)
																			} else {
																				F_relation_close(m, v123, int32(3))
																				mBase = m.M
																				v218 = m.ExcPending
																				if v218 != 0 {
																					return int32(0)
																				} else {
																					F_CommandCounterIncrement(m)
																					mBase = m.M
																					v220 = m.ExcPending
																					if v220 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v19 - int32(-64)
																						return v110
																					}
																				}
																			}
																		}
																	}
																}
															}
														} else {
															v181 = v176 << (uint(int32(2)) % 32)
															*(*int32)(unsafe.Add(mBase, uint32(v181+(v17+int32(-32))))) = v120
															*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)+v181))) = int32(13)
															v195 = v176 + int32(1)
															v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
															v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return int32(0)
															} else {
																F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int32(0)
																} else {
																	F_pfree(m, v204)
																	mBase = m.M
																	v211 = m.ExcPending
																	if v211 != 0 {
																		return int32(0)
																	} else {
																		F_ReleaseCatCache(m, v126)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v123, int32(3))
																			mBase = m.M
																			v218 = m.ExcPending
																			if v218 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v220 = m.ExcPending
																				if v220 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v19 - int32(-64)
																					return v110
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									} else {
										v90 = F_errstart(m, int32(19), int32(0))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
											return int32(0)
										} else {
											if v90 != 0 {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v94 = m.ExcPending
												if v94 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(_a_F_relation_statistics_update_4)
													F_errmsg(m, int32(_a_F_relation_statistics_update_5), v17+int32(-48))
													mBase = m.M
													v101 = m.ExcPending
													if v101 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(_a_F_relation_statistics_update_2), int32(117), int32(_a_F_relation_statistics_update_3))
														mBase = m.M
														v106 = m.ExcPending
														if v106 != 0 {
															return int32(0)
														} else {
															v108 = v2
															v109 = v81
															v110 = int32(0)
															v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
															if v111 == int32(0) {
																v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
																v115 = v114
															} else {
																v115 = v2
															}
															v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
															if v116 == int32(0) {
																v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
																v120 = v119
															} else {
																v120 = v2
															}
															v123 = F_table_open(m, int32(1259), int32(3))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int32(0)
															} else {
																v126 = F_SearchSysCache1(m, int32(57), v72)
																mBase = m.M
																v127 = m.ExcPending
																if v127 != 0 {
																	return int32(0)
																} else {
																	if v126 == int32(0) {
																		F_errstart_cold(m, int32(21), int32(0))
																		mBase = m.M
																		v248 = m.ExcPending
																		if v248 != 0 {
																			return int32(0)
																		} else {
																			*(*int32)(unsafe.Add(mBase, uint32(v19))) = v72
																			F_errmsg_internal(m, int32(_a_F_relation_statistics_update_1), v19)
																			mBase = m.M
																			v252 = m.ExcPending
																			if v252 != 0 {
																				return int32(0)
																			} else {
																				F_errfinish(m, int32(_a_F_relation_statistics_update_2), int32(144), int32(_a_F_relation_statistics_update_3))
																				mBase = m.M
																				v257 = m.ExcPending
																				if v257 != 0 {
																					return int32(0)
																				} else {
																					base.Wasm_trap_unreachable()
																					for {
																					}
																				}
																			}
																		}
																	} else {
																		v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
																		v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+22)))
																		v132 = v130 + v131
																		v133 = int32(0)
																		v135 = v17 + int32(-16)
																		v137 = v17 + int32(-32)
																		if v74 != 0 {
																			v148 = v133
																			v149 = v135
																			v150 = v137
																		} else {
																			v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+96))
																			if v78 == v138 {
																				v148 = v133
																				v149 = v135
																				v150 = v137
																			} else {
																				v140 = int32(4)
																				*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78
																				*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(10)
																				v148 = int32(1)
																				v149 = v135 | v140
																				v150 = v137 | v140
																			}
																		}
																		if v108 == int32(0) {
																			v160 = v148
																		} else {
																			v153 = *(*float32)(unsafe.Add(mBase, uint32(v132)+100))
																			if base.F32_eq(v109, v153) != 0 {
																				v160 = v148
																			} else {
																				*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(11)
																				*(*float32)(unsafe.Add(mBase, uint32(v150))) = v109
																				v160 = v148 + int32(1)
																			}
																		}
																		if v111 != 0 {
																			v176 = v160
																		} else {
																			v161 = *(*int32)(unsafe.Add(mBase, uint32(v132)+104))
																			if v115 == v161 {
																				v176 = v160
																			} else {
																				v164 = v160 << (uint(int32(2)) % 32)
																				*(*int32)(unsafe.Add(mBase, uint32(v164|(v17+int32(-32))))) = v115
																				*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)|v164))) = int32(12)
																				v176 = v160 + int32(1)
																			}
																		}
																		if v116 != 0 {
																			if v176 == int32(0) {
																				F_ReleaseCatCache(m, v126)
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v123, int32(3))
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v220 = m.ExcPending
																						if v220 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v19 - int32(-64)
																							return v110
																						}
																					}
																				}
																			} else {
																				v195 = v176
																				v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
																				v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return int32(0)
																				} else {
																					F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																					mBase = m.M
																					v209 = m.ExcPending
																					if v209 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v204)
																						mBase = m.M
																						v211 = m.ExcPending
																						if v211 != 0 {
																							return int32(0)
																						} else {
																							F_ReleaseCatCache(m, v126)
																							mBase = m.M
																							v215 = m.ExcPending
																							if v215 != 0 {
																								return int32(0)
																							} else {
																								F_relation_close(m, v123, int32(3))
																								mBase = m.M
																								v218 = m.ExcPending
																								if v218 != 0 {
																									return int32(0)
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v220 = m.ExcPending
																									if v220 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v19 - int32(-64)
																										return v110
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v178 = *(*int32)(unsafe.Add(mBase, uint32(v132)+108))
																			if v120 == v178 {
																				if v176 == int32(0) {
																					F_ReleaseCatCache(m, v126)
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v123, int32(3))
																						mBase = m.M
																						v218 = m.ExcPending
																						if v218 != 0 {
																							return int32(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v220 = m.ExcPending
																							if v220 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v19 - int32(-64)
																								return v110
																							}
																						}
																					}
																				} else {
																					v195 = v176
																					v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
																					v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																					mBase = m.M
																					v205 = m.ExcPending
																					if v205 != 0 {
																						return int32(0)
																					} else {
																						F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																						mBase = m.M
																						v209 = m.ExcPending
																						if v209 != 0 {
																							return int32(0)
																						} else {
																							F_pfree(m, v204)
																							mBase = m.M
																							v211 = m.ExcPending
																							if v211 != 0 {
																								return int32(0)
																							} else {
																								F_ReleaseCatCache(m, v126)
																								mBase = m.M
																								v215 = m.ExcPending
																								if v215 != 0 {
																									return int32(0)
																								} else {
																									F_relation_close(m, v123, int32(3))
																									mBase = m.M
																									v218 = m.ExcPending
																									if v218 != 0 {
																										return int32(0)
																									} else {
																										F_CommandCounterIncrement(m)
																										mBase = m.M
																										v220 = m.ExcPending
																										if v220 != 0 {
																											return int32(0)
																										} else {
																											m.G0 = v19 - int32(-64)
																											return v110
																										}
																									}
																								}
																							}
																						}
																					}
																				}
																			} else {
																				v181 = v176 << (uint(int32(2)) % 32)
																				*(*int32)(unsafe.Add(mBase, uint32(v181+(v17+int32(-32))))) = v120
																				*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)+v181))) = int32(13)
																				v195 = v176 + int32(1)
																				v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
																				v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																				mBase = m.M
																				v205 = m.ExcPending
																				if v205 != 0 {
																					return int32(0)
																				} else {
																					F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																					mBase = m.M
																					v209 = m.ExcPending
																					if v209 != 0 {
																						return int32(0)
																					} else {
																						F_pfree(m, v204)
																						mBase = m.M
																						v211 = m.ExcPending
																						if v211 != 0 {
																							return int32(0)
																						} else {
																							F_ReleaseCatCache(m, v126)
																							mBase = m.M
																							v215 = m.ExcPending
																							if v215 != 0 {
																								return int32(0)
																							} else {
																								F_relation_close(m, v123, int32(3))
																								mBase = m.M
																								v218 = m.ExcPending
																								if v218 != 0 {
																									return int32(0)
																								} else {
																									F_CommandCounterIncrement(m)
																									mBase = m.M
																									v220 = m.ExcPending
																									if v220 != 0 {
																										return int32(0)
																									} else {
																										m.G0 = v19 - int32(-64)
																										return v110
																									}
																								}
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											} else {
												v108 = v2
												v109 = v81
												v110 = int32(0)
												v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+56)))
												if v111 == int32(0) {
													v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
													v115 = v114
												} else {
													v115 = v2
												}
												v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+64)))
												if v116 == int32(0) {
													v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
													v120 = v119
												} else {
													v120 = v2
												}
												v123 = F_table_open(m, int32(1259), int32(3))
												mBase = m.M
												v124 = m.ExcPending
												if v124 != 0 {
													return int32(0)
												} else {
													v126 = F_SearchSysCache1(m, int32(57), v72)
													mBase = m.M
													v127 = m.ExcPending
													if v127 != 0 {
														return int32(0)
													} else {
														if v126 == int32(0) {
															F_errstart_cold(m, int32(21), int32(0))
															mBase = m.M
															v248 = m.ExcPending
															if v248 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v19))) = v72
																F_errmsg_internal(m, int32(_a_F_relation_statistics_update_1), v19)
																mBase = m.M
																v252 = m.ExcPending
																if v252 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_relation_statistics_update_2), int32(144), int32(_a_F_relation_statistics_update_3))
																	mBase = m.M
																	v257 = m.ExcPending
																	if v257 != 0 {
																		return int32(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														} else {
															v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
															v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+22)))
															v132 = v130 + v131
															v133 = int32(0)
															v135 = v17 + int32(-16)
															v137 = v17 + int32(-32)
															if v74 != 0 {
																v148 = v133
																v149 = v135
																v150 = v137
															} else {
																v138 = *(*int32)(unsafe.Add(mBase, uint32(v132)+96))
																if v78 == v138 {
																	v148 = v133
																	v149 = v135
																	v150 = v137
																} else {
																	v140 = int32(4)
																	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78
																	*(*int32)(unsafe.Add(mBase, uint32(v19)+48)) = int32(10)
																	v148 = int32(1)
																	v149 = v135 | v140
																	v150 = v137 | v140
																}
															}
															if v108 == int32(0) {
																v160 = v148
															} else {
																v153 = *(*float32)(unsafe.Add(mBase, uint32(v132)+100))
																if base.F32_eq(v109, v153) != 0 {
																	v160 = v148
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v149))) = int32(11)
																	*(*float32)(unsafe.Add(mBase, uint32(v150))) = v109
																	v160 = v148 + int32(1)
																}
															}
															if v111 != 0 {
																v176 = v160
															} else {
																v161 = *(*int32)(unsafe.Add(mBase, uint32(v132)+104))
																if v115 == v161 {
																	v176 = v160
																} else {
																	v164 = v160 << (uint(int32(2)) % 32)
																	*(*int32)(unsafe.Add(mBase, uint32(v164|(v17+int32(-32))))) = v115
																	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)|v164))) = int32(12)
																	v176 = v160 + int32(1)
																}
															}
															if v116 != 0 {
																if v176 == int32(0) {
																	F_ReleaseCatCache(m, v126)
																	mBase = m.M
																	v215 = m.ExcPending
																	if v215 != 0 {
																		return int32(0)
																	} else {
																		F_relation_close(m, v123, int32(3))
																		mBase = m.M
																		v218 = m.ExcPending
																		if v218 != 0 {
																			return int32(0)
																		} else {
																			F_CommandCounterIncrement(m)
																			mBase = m.M
																			v220 = m.ExcPending
																			if v220 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v19 - int32(-64)
																				return v110
																			}
																		}
																	}
																} else {
																	v195 = v176
																	v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
																	v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return int32(0)
																	} else {
																		F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																		mBase = m.M
																		v209 = m.ExcPending
																		if v209 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v204)
																			mBase = m.M
																			v211 = m.ExcPending
																			if v211 != 0 {
																				return int32(0)
																			} else {
																				F_ReleaseCatCache(m, v126)
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v123, int32(3))
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v220 = m.ExcPending
																						if v220 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v19 - int32(-64)
																							return v110
																						}
																					}
																				}
																			}
																		}
																	}
																}
															} else {
																v178 = *(*int32)(unsafe.Add(mBase, uint32(v132)+108))
																if v120 == v178 {
																	if v176 == int32(0) {
																		F_ReleaseCatCache(m, v126)
																		mBase = m.M
																		v215 = m.ExcPending
																		if v215 != 0 {
																			return int32(0)
																		} else {
																			F_relation_close(m, v123, int32(3))
																			mBase = m.M
																			v218 = m.ExcPending
																			if v218 != 0 {
																				return int32(0)
																			} else {
																				F_CommandCounterIncrement(m)
																				mBase = m.M
																				v220 = m.ExcPending
																				if v220 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v19 - int32(-64)
																					return v110
																				}
																			}
																		}
																	} else {
																		v195 = v176
																		v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
																		v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																		mBase = m.M
																		v205 = m.ExcPending
																		if v205 != 0 {
																			return int32(0)
																		} else {
																			F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																			mBase = m.M
																			v209 = m.ExcPending
																			if v209 != 0 {
																				return int32(0)
																			} else {
																				F_pfree(m, v204)
																				mBase = m.M
																				v211 = m.ExcPending
																				if v211 != 0 {
																					return int32(0)
																				} else {
																					F_ReleaseCatCache(m, v126)
																					mBase = m.M
																					v215 = m.ExcPending
																					if v215 != 0 {
																						return int32(0)
																					} else {
																						F_relation_close(m, v123, int32(3))
																						mBase = m.M
																						v218 = m.ExcPending
																						if v218 != 0 {
																							return int32(0)
																						} else {
																							F_CommandCounterIncrement(m)
																							mBase = m.M
																							v220 = m.ExcPending
																							if v220 != 0 {
																								return int32(0)
																							} else {
																								m.G0 = v19 - int32(-64)
																								return v110
																							}
																						}
																					}
																				}
																			}
																		}
																	}
																} else {
																	v181 = v176 << (uint(int32(2)) % 32)
																	*(*int32)(unsafe.Add(mBase, uint32(v181+(v17+int32(-32))))) = v120
																	*(*int32)(unsafe.Add(mBase, uint32(v17+int32(-16)+v181))) = int32(13)
																	v195 = v176 + int32(1)
																	v197 = *(*int32)(unsafe.Add(mBase, uint32(v123)+52))
																	v204 = F_heap_modify_tuple_by_cols(m, v126, v197, v195, v17+int32(-16), v17+int32(-32), v17+int32(-36))
																	mBase = m.M
																	v205 = m.ExcPending
																	if v205 != 0 {
																		return int32(0)
																	} else {
																		F_CatalogTupleUpdate(m, v123, v204+int32(4), v204)
																		mBase = m.M
																		v209 = m.ExcPending
																		if v209 != 0 {
																			return int32(0)
																		} else {
																			F_pfree(m, v204)
																			mBase = m.M
																			v211 = m.ExcPending
																			if v211 != 0 {
																				return int32(0)
																			} else {
																				F_ReleaseCatCache(m, v126)
																				mBase = m.M
																				v215 = m.ExcPending
																				if v215 != 0 {
																					return int32(0)
																				} else {
																					F_relation_close(m, v123, int32(3))
																					mBase = m.M
																					v218 = m.ExcPending
																					if v218 != 0 {
																						return int32(0)
																					} else {
																						F_CommandCounterIncrement(m)
																						mBase = m.M
																						v220 = m.ExcPending
																						if v220 != 0 {
																							return int32(0)
																						} else {
																							m.G0 = v19 - int32(-64)
																							return v110
																						}
																					}
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v228 = m.ExcPending
						if v228 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v231 = m.ExcPending
							if v231 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_relation_statistics_update_6), int32(0))
								mBase = m.M
								v235 = m.ExcPending
								if v235 != 0 {
									return int32(0)
								} else {
									F_errhint(m, int32(_a_F_relation_statistics_update_7), int32(0))
									mBase = m.M
									v239 = m.ExcPending
									if v239 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_relation_statistics_update_2), int32(98), int32(_a_F_relation_statistics_update_3))
										mBase = m.M
										v244 = m.ExcPending
										if v244 != 0 {
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
					}
				}
			}
		}
	}
}
