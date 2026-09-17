package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BuildTupleFromCStrings(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v14 = F_palloc(m, v11<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v18 = F_palloc(m, v11)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < v11 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v24 = int32(0)
	goto L7
L5:
	;
	goto L6
L6:
	;
	v83 = F_heap_form_tuple(m, v10, v14, v18)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L18
	}
L7:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v24<<(uint(int32(4))%32))+29)))
	if v34 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L6
L9:
	;
	v72 = v24 + int32(1)
	if v72 != v11 {
		v24 = v72
		goto L7
	} else {
		goto L17
	}
L10:
	;
	v38 = v24 << (uint(int32(2)) % 32)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = l1 + v38
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v46+v38)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49+v38)))
	v52 = F_InputFunctionCall(m, v40+v24*int32(28), v45, v48, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v24<<(uint(int32(2))%32)))) = int32(0)
	v67 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24+v18))) = uint8(v67)
	goto L9
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14+v38))) = v52
	v55 = v24 + v18
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v56 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v57 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v57)
	goto L9
L15:
	;
	goto L16
L16:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v59)
	goto L9
L17:
	;
	goto L8
L18:
	;
	F_pfree(m, v14)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_pfree(m, v18)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	return v83
}
func F_CreateTupleDescCopyConstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
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
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_palloc(m, v10*int32(116)+int32(20))
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
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v10
	*(*int64)(unsafe.Add(mBase, uint32(v15)+12)) = int64(4294967295)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+4)) = int64(-4294965047)
	v25 = v10 * int32(100)
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = int32(4)
	v29 = int32(20)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	base.MemoryCopy(m, v15+v10<<(uint(v26)%32)+v29, l0+v31<<(uint(v26)%32)+v29, v25)
	goto L5
L4:
	;
	goto L5
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if int32(0) < v38 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v43 = int32(0)
	goto L9
L7:
	;
	goto L8
L8:
	;
	if v9 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	F_populate_compact_attribute(m, v15, v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	v53 = v43 << (uint(int32(4)) % 32)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v53)+31)))
	*(*uint8)(unsafe.Add(mBase, uint32(v15+v53)+31)) = uint8(v56)
	v59 = v43 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v59 < v60 {
		v43 = v59
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	v71 = F_palloc0(m, int32(20))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v15)+8)) = v270
	return v15
L16:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+16)))
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+16)) = uint8(v73)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+17)))
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+17)) = uint8(v75)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+18)))
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+18)) = uint8(v77)
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v71)+12)) = uint16(v79)
	if v79 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v126 == int32(0) {
		goto L28
	} else {
		goto L29
	}
L18:
	;
	v85 = F_palloc(m, v79<<(uint(int32(3))%32))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v85
	v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+12)))
	v90 = v88 << (uint(int32(3)) % 32)
	if v90 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	base.MemoryCopy(m, v85, v91, v90)
	goto L22
L21:
	;
	goto L22
L22:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+12)))
	if v93 == int32(0) {
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v97 = v93
	goto L24
L24:
	;
	v105 = v97 - int32(1)
	v107 = v105 << (uint(int32(3)) % 32)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107+v108)+4))
	v111 = F_pstrdup(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	goto L17
L26:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v113+v107)+4)) = v111
	if base.Ui32(int32(1)) < base.Ui32(v97) {
		v97 = v105
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(v71)+14)) = uint16(v186)
	if v186 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L29:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v132 = F_palloc(m, v129<<(uint(int32(3))%32))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v132
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v137 = v135 << (uint(int32(3)) % 32)
	if v137 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	base.MemoryCopy(m, v132, v138, v137)
	goto L33
L32:
	;
	goto L33
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = v140 - int32(1)
	if v142 < int32(0) {
		goto L28
	} else {
		goto L34
	}
L34:
	;
	v148 = v142
	goto L35
L35:
	;
	v156 = v148 << (uint(int32(3)) % 32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	v158 = v156 + v157
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	if v159 == int32(1) {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L28
L37:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v165 = l0 + int32(20) + v148<<(uint(int32(4))%32)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+6)))
	v167 = int32(*(*int16)(unsafe.Add(mBase, uint32(v165)+4)))
	v168 = F_datumCopy(m, v162, v166, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	if int32(0) < v148 {
		v148 = v148 - int32(1)
		goto L35
	} else {
		goto L41
	}
L40:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v170+v156)+4)) = v168
	goto L39
L41:
	;
	goto L36
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v71
	goto L15
L43:
	;
	v192 = F_palloc(m, v186*int32(12))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v192
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+14)))
	v197 = v195 * int32(12)
	if v197 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	base.MemoryCopy(m, v192, v198, v197)
	goto L47
L46:
	;
	goto L47
L47:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+14)))
	if v200 == int32(0) {
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v204 = v200
	goto L49
L49:
	;
	v212 = v204 - int32(1)
	v214 = v212 * int32(12)
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v214+v215)))
	v218 = F_pstrdup(m, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L1
	} else {
		goto L51
	}
L50:
	;
	goto L42
L51:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v220+v214))) = v218
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v223+v214)+4))
	v226 = F_pstrdup(m, v225)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v228+v214)+4)) = v226
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v233+v214)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v231+v214)+8)) = uint8(v235)
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239+v214)+9)))
	*(*uint8)(unsafe.Add(mBase, uint32(v237+v214)+9)) = uint8(v241)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v214)+10)))
	*(*uint8)(unsafe.Add(mBase, uint32(v243+v214)+10)) = uint8(v247)
	if base.Ui32(int32(1)) < base.Ui32(v204) {
		v204 = v212
		goto L49
	} else {
		goto L53
	}
L53:
	;
	goto L50
}
func F_ExecSetTupleBound(m *base.Module, l0 int64, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	v5 = l1
	goto L2
L1:
	;
	return
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	switch v7 - int32(394) {
	case 0:
		goto L8
	default:
		goto L1
	case 3:
		goto L11
	case 4:
		goto L12
	case 17:
		goto L7
	case 32:
		goto L10
	case 33:
		goto L9
	case 38, 39:
		goto L6
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+120)) = l0
	v206 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+117)) = uint8(v206)
	goto L1
L4:
	;
	goto L3
L5:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v5+v202)))
	v5 = v204
	goto L2
L6:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+112)) = l0
	v202 = int32(36)
	goto L5
L7:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v5)+32))
	if v197 == int32(0) {
		v202 = int32(116)
		goto L5
	} else {
		goto L80
	}
L8:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v5)+36))
	if v195 != 0 {
		v5 = v195
		goto L2
	} else {
		goto L79
	}
L9:
	;
	if l0 < int64(0) {
		goto L76
	} else {
		goto L77
	}
L10:
	;
	if int64(0) <= l0 {
		goto L4
	} else {
		goto L75
	}
L11:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v5)+108))
	if v97 <= int32(0) {
		goto L1
	} else {
		goto L44
	}
L12:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)+108))
	if v10 <= int32(0) {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v16 = int32(0)
	goto L14
L14:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+104))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v17+v16<<(uint(int32(2))%32))))
	v24 = v21
	goto L18
L15:
	;
	goto L1
L16:
	;
	v94 = v16 + int32(1)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v5)+108))
	if v94 < v95 {
		v16 = v94
		goto L14
	} else {
		goto L43
	}
L17:
	;
	goto L16
L18:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	switch v26 - int32(394) {
	case 0:
		goto L24
	default:
		goto L17
	case 3:
		goto L27
	case 4:
		goto L28
	case 17:
		goto L23
	case 32:
		goto L26
	case 33:
		goto L25
	case 38, 39:
		goto L22
	}
L19:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+120)) = l0
	v85 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+117)) = uint8(v85)
	goto L17
L20:
	;
	goto L19
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v24+v81)))
	v24 = v83
	goto L18
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+112)) = l0
	v81 = int32(36)
	goto L21
L23:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	if v76 == int32(0) {
		v81 = int32(116)
		goto L21
	} else {
		goto L42
	}
L24:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	if v74 != 0 {
		v24 = v74
		goto L18
	} else {
		goto L41
	}
L25:
	;
	if l0 < int64(0) {
		goto L38
	} else {
		goto L39
	}
L26:
	;
	if int64(0) <= l0 {
		goto L20
	} else {
		goto L37
	}
L27:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v24)+108))
	if v46 <= int32(0) {
		goto L17
	} else {
		goto L33
	}
L28:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+108))
	if v29 <= int32(0) {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	v35 = int32(0)
	goto L30
L30:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v35<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, l0, v40)
	mBase = m.M
	v43 = v35 + int32(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v24)+108))
	if v43 < v44 {
		v35 = v43
		goto L30
	} else {
		goto L32
	}
L31:
	;
	goto L17
L32:
	;
	goto L31
L33:
	;
	v52 = int32(0)
	goto L34
L34:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v24)+104))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v53+v52<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, l0, v57)
	mBase = m.M
	v60 = v52 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v24)+108))
	if v60 < v61 {
		v52 = v60
		goto L34
	} else {
		goto L36
	}
L35:
	;
	goto L17
L36:
	;
	goto L35
L37:
	;
	v65 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+117)) = uint8(v65)
	goto L16
L38:
	;
	v69 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+116)) = uint8(v69)
	goto L16
L39:
	;
	goto L40
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v24)+120)) = l0
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+116)) = uint8(v72)
	goto L16
L41:
	;
	goto L17
L42:
	;
	goto L17
L43:
	;
	goto L15
L44:
	;
	v103 = int32(0)
	goto L45
L45:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v5)+104))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v104+v103<<(uint(int32(2))%32))))
	v111 = v108
	goto L49
L46:
	;
	goto L1
L47:
	;
	v181 = v103 + int32(1)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v5)+108))
	if v181 < v182 {
		v103 = v181
		goto L45
	} else {
		goto L74
	}
L48:
	;
	goto L47
L49:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v111)))
	switch v113 - int32(394) {
	case 0:
		goto L55
	default:
		goto L48
	case 3:
		goto L58
	case 4:
		goto L59
	case 17:
		goto L54
	case 32:
		goto L57
	case 33:
		goto L56
	case 38, 39:
		goto L53
	}
L50:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v111)+120)) = l0
	v172 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+117)) = uint8(v172)
	goto L48
L51:
	;
	goto L50
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v111+v168)))
	v111 = v170
	goto L49
L53:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v111)+112)) = l0
	v168 = int32(36)
	goto L52
L54:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
	if v163 == int32(0) {
		v168 = int32(116)
		goto L52
	} else {
		goto L73
	}
L55:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v111)+36))
	if v161 != 0 {
		v111 = v161
		goto L49
	} else {
		goto L72
	}
L56:
	;
	if l0 < int64(0) {
		goto L69
	} else {
		goto L70
	}
L57:
	;
	if int64(0) <= l0 {
		goto L51
	} else {
		goto L68
	}
L58:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v111)+108))
	if v133 <= int32(0) {
		goto L48
	} else {
		goto L64
	}
L59:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v111)+108))
	if v116 <= int32(0) {
		goto L48
	} else {
		goto L60
	}
L60:
	;
	v122 = int32(0)
	goto L61
L61:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v111)+104))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123+v122<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, l0, v127)
	mBase = m.M
	v130 = v122 + int32(1)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v111)+108))
	if v130 < v131 {
		v122 = v130
		goto L61
	} else {
		goto L63
	}
L62:
	;
	goto L48
L63:
	;
	goto L62
L64:
	;
	v139 = int32(0)
	goto L65
L65:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v111)+104))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v139<<(uint(int32(2))%32))))
	F_ExecSetTupleBound(m, l0, v144)
	mBase = m.M
	v147 = v139 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v111)+108))
	if v147 < v148 {
		v139 = v147
		goto L65
	} else {
		goto L67
	}
L66:
	;
	goto L48
L67:
	;
	goto L66
L68:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+117)) = uint8(v152)
	goto L47
L69:
	;
	v156 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+116)) = uint8(v156)
	goto L47
L70:
	;
	goto L71
L71:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v111)+120)) = l0
	v159 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+116)) = uint8(v159)
	goto L47
L72:
	;
	goto L48
L73:
	;
	goto L48
L74:
	;
	goto L46
L75:
	;
	v186 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+117)) = uint8(v186)
	return
L76:
	;
	v190 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+116)) = uint8(v190)
	return
L77:
	;
	goto L78
L78:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v5)+120)) = l0
	v193 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+116)) = uint8(v193)
	return
L79:
	;
	goto L1
L80:
	;
	goto L1
}
func F_GetNextTuple(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
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
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v15 = F_tuplesort_gettupleslot(m, l0, int32(1), v6, l2, v6)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			v17 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
			if v17 <= int32(0) {
				F_slot_getsomeattrs_int(m, l2, int32(1))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
					v24 = v23
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
					*(*int32)(unsafe.Add(mBase, uint32(l4))) = v27
					if v24 <= int32(2) {
						F_slot_getsomeattrs_int(m, l2, int32(3))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
							v36 = v34
							v37 = v35
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
							*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v38)
							v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v40
							v46 = F_index_form_tuple(m, l1, v10+int32(12), v10+int32(11))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l3))) = v46
								v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
								if v49 <= int32(1) {
									F_slot_getsomeattrs_int(m, l2, int32(2))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return
									} else {
										v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
										v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
										v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
										*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v57)
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
										*(*int32)(unsafe.Add(mBase, uint32(v46))) = v59
										m.G0 = v10 + int32(16)
										return
									}
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
									v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
									*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v57)
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									*(*int32)(unsafe.Add(mBase, uint32(v46))) = v59
									m.G0 = v10 + int32(16)
									return
								}
							}
						}
					} else {
						v36 = v26
						v37 = v25
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v38)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v40
						v46 = F_index_form_tuple(m, l1, v10+int32(12), v10+int32(11))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v46
							v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
							if v49 <= int32(1) {
								F_slot_getsomeattrs_int(m, l2, int32(2))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
									v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
									*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v57)
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									*(*int32)(unsafe.Add(mBase, uint32(v46))) = v59
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
								v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v57)
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								*(*int32)(unsafe.Add(mBase, uint32(v46))) = v59
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			} else {
				v24 = v17
				v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = v27
				if v24 <= int32(2) {
					F_slot_getsomeattrs_int(m, l2, int32(3))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
						v36 = v34
						v37 = v35
						v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
						*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v38)
						v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v40
						v46 = F_index_form_tuple(m, l1, v10+int32(12), v10+int32(11))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v46
							v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
							if v49 <= int32(1) {
								F_slot_getsomeattrs_int(m, l2, int32(2))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
									v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
									v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
									*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v57)
									v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
									*(*int32)(unsafe.Add(mBase, uint32(v46))) = v59
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
								v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v57)
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								*(*int32)(unsafe.Add(mBase, uint32(v46))) = v59
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				} else {
					v36 = v26
					v37 = v25
					v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)) = uint8(v38)
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v40
					v46 = F_index_form_tuple(m, l1, v10+int32(12), v10+int32(11))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v46
						v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(l2)+6)))
						if v49 <= int32(1) {
							F_slot_getsomeattrs_int(m, l2, int32(2))
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
								v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v57)
								v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
								*(*int32)(unsafe.Add(mBase, uint32(v46))) = v59
								m.G0 = v10 + int32(16)
								return
							}
						} else {
							v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
							v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
							v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v56)+4)))
							*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v57)
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
							*(*int32)(unsafe.Add(mBase, uint32(v46))) = v59
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(-1)
			m.G0 = v10 + int32(16)
			return
		}
	}
}
func F_LockTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v4 = int32(0)
	v5 = m.G0
	v6 = int32(16)
	v7 = v5 - v6
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v11
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v13 | v14<<(uint(v6)%32)
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v20 = int32(260)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+14)) = uint16(v20)
	*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)) = uint16(v19)
	v25 = F_LockAcquire(m, v7, l2, v4, v4)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_ResetTupleHashTable(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
	v6 = v4 * int32(12)
	if v6 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)+20))
		base.MemoryFill(m, v7, int32(0), v6)
	} else {
	}
	*(*int32)(unsafe.Add(mBase, uint32(v3)+8)) = int32(0)
	return
}
func F_TupleQueueReaderNext(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l2 != 0 {
		v10 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v10)
	} else {
	}
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = F_shm_mq_receive(m, v12, v8+int32(12), v8+int32(8), l1)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		switch v17 - int32(1) {
		case 0:
			v28 = v4
		case 1:
			if l2 == int32(0) {
				v28 = v4
			} else {
				v25 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v25)
				v28 = v4
			}
		default:
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
			v28 = v27
		}
		m.G0 = v8 + int32(16)
		return v28
	}
}
func F_exec_init_tuple_store(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
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
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v6 == int32(0) {
		F_errstart_cold(m, int32(21), int32(_a_F_exec_init_tuple_store_0))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			F_errcode(m, int32(1088))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_exec_init_tuple_store_1), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_exec_init_tuple_store_2), int32(3681), int32(_a_F_exec_init_tuple_store_3))
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
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
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		if v9 != int32(383) {
			F_errstart_cold(m, int32(21), int32(_a_F_exec_init_tuple_store_0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_exec_init_tuple_store_1), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_exec_init_tuple_store_2), int32(3681), int32(_a_F_exec_init_tuple_store_3))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
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
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			if v12&int32(2) == int32(0) {
				F_errstart_cold(m, int32(21), int32(_a_F_exec_init_tuple_store_0))
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v68 = m.ExcPending
					if v68 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_exec_init_tuple_store_4), int32(0))
						mBase = m.M
						v72 = m.ExcPending
						if v72 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_exec_init_tuple_store_2), int32(3687), int32(_a_F_exec_init_tuple_store_3))
							mBase = m.M
							v77 = m.ExcPending
							if v77 != 0 {
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
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				if v17 == int32(0) {
					F_errstart_cold(m, int32(21), int32(_a_F_exec_init_tuple_store_0))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_exec_init_tuple_store_4), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_exec_init_tuple_store_2), int32(3687), int32(_a_F_exec_init_tuple_store_3))
								mBase = m.M
								v77 = m.ExcPending
								if v77 != 0 {
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
					v20 = int32(_a_F_exec_init_tuple_store_5)
					v21 = *(*int32)(unsafe.Add(mBase, _c_F_exec_init_tuple_store[0]))
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, _c_F_exec_init_tuple_store[0])) = v23
					v25 = int32(_a_F_exec_init_tuple_store_6)
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_exec_init_tuple_store[1]))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
					*(*int32)(unsafe.Add(mBase, _c_F_exec_init_tuple_store[1])) = v28
					v36 = *(*int32)(unsafe.Add(mBase, _c_F_exec_init_tuple_store[2]))
					v37 = F_tuplestore_begin_heap(m, int32(base.Ui32(v12&int32(4))>>(uint(int32(2))%32)), int32(0), v36)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v37
						*(*int32)(unsafe.Add(mBase, _c_F_exec_init_tuple_store[0])) = v21
						*(*int32)(unsafe.Add(mBase, _c_F_exec_init_tuple_store[1])) = v26
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+44)) = v44
						return
					}
				}
			}
		}
	}
}
func F_expandTupleDesc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
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
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	v12 = int32(0)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v20 = v18
	goto L3
L2:
	;
	v20 = int32(0)
	goto L3
L3:
	;
	if l3 < v20 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v26 = v22 + l3<<(uint(int32(2))%32)
	goto L6
L5:
	;
	v26 = v12
	goto L6
L6:
	;
	if int32(0) < l2 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v40 = v26
	v43 = v12
	goto L10
L8:
	;
	goto L9
L9:
	;
	return
L10:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = l0 + v45<<(uint(int32(4))%32) + v43*int32(100)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+111)))
	if v52 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L9
L12:
	;
	v146 = v43 + int32(1)
	if v146 != l2 {
		v40 = v144
		v43 = v146
		goto L10
	} else {
		goto L51
	}
L13:
	;
	if l8 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	v92 = v51 + int32(20)
	if l9 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L16:
	;
	v77 = int32(0)
	if v40 == v77 {
		v144 = v77
		goto L12
	} else {
		goto L28
	}
L17:
	;
	if l9 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	v59 = F_pstrdup(m, int32(_a_F_expandTupleDesc_0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	if l10 == int32(0) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	return
L22:
	;
	v61 = F_makeString(m, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v63 = F_lappend(m, v57, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L21
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v63
	goto L20
L25:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v72 = F_makeNullConst(m, int32(23), int32(-1), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L21
	} else {
		goto L26
	}
L26:
	;
	v74 = F_lappend(m, v68, v72)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v74
	goto L16
L28:
	;
	v81 = v40 + int32(4)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)+12))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if base.Ui32(v81) < base.Ui32(v84+v85<<(uint(int32(2))%32)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v90 = v81
	goto L31
L30:
	;
	v90 = int32(0)
	goto L31
L31:
	;
	v144 = v90
	goto L12
L32:
	;
	if l10 != 0 {
		goto L46
	} else {
		goto L47
	}
L33:
	;
	v122 = v40
	goto L32
L34:
	;
	goto L35
L35:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v40 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v115 = F_pstrdup(m, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L21
	} else {
		goto L43
	}
L37:
	;
	v97 = v40 + int32(4)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if base.Ui32(v97) < base.Ui32(v100+v101<<(uint(int32(2))%32)) {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	v112 = int32(0)
	v114 = v51 + int32(24)
	goto L36
L40:
	;
	v106 = v97
	goto L42
L41:
	;
	v106 = int32(0)
	goto L42
L42:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
	v112 = v106
	v114 = v108
	goto L36
L43:
	;
	v117 = F_makeString(m, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L21
	} else {
		goto L44
	}
L44:
	;
	v119 = F_lappend(m, v95, v117)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L21
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v119
	v122 = v112
	goto L32
L46:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v92)+68))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v92)+76))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v92)+96))
	v131 = F_makeVar(m, l4, base.I32_extend16_s(l3+v43+int32(1)), v128, v129, v130, l5)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L21
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v144 = v122
	goto L12
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131)+44)) = l7
	*(*int32)(unsafe.Add(mBase, uint32(v131)+32)) = l6
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	v136 = F_lappend(m, v135, v131)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L21
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = v136
	goto L48
L51:
	;
	goto L11
}
