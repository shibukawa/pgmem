package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_addTargetToGroupList(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v20 = F_exprType(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v20 == int32(705) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = int32(25)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v30 = int32(-1)
	v34 = F_coerce_type(m, l0, v27, int32(705), v26, v30, int32(0), int32(2), v30)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	v37 = v20
	goto L5
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v38 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v34
	v37 = v26
	goto L5
L7:
	;
	m.G0 = v17 + int32(32)
	return v274
L8:
	;
	v85 = F_palloc0(m, int32(20))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L16
	}
L9:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v43 <= int32(0) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v52 = int32(0)
	goto L12
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v46+v52<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v65 == v38 {
		v274 = l2
		goto L7
	} else {
		goto L14
	}
L13:
	;
	goto L8
L14:
	;
	v68 = v52 + int32(1)
	if v68 != v43 {
		v52 = v68
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85))) = int32(106)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v17
	v94 = int32(4541672)
	v95 = *(*int32)(unsafe.Add(mBase, _consts[74]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v95
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v17 + int32(8)
	goto L17
L17:
	;
	v101 = int32(0)
	F_get_sort_group_operators(m, v37, v101, int32(1), v101, v17+int32(28), v17+int32(24), v101, v17+int32(23))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	*(*int32)(unsafe.Add(mBase, _consts[74])) = v114
	goto L19
L19:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v116 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if l3 == int32(0) {
		v245 = int32(1)
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v252 = v116
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+4)) = v252
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v85)+8)) = v262
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v265 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v85)+16)) = uint16(v265)
	*(*int32)(unsafe.Add(mBase, uint32(v85)+12)) = v264
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+23)))
	*(*uint8)(unsafe.Add(mBase, uint32(v85)+18)) = uint8(v268)
	v270 = F_lappend(m, l2, v85)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L53
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v245
	v252 = v245
	goto L22
L24:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v123 <= int32(0) {
		v245 = int32(1)
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v127 = v123 & int32(3)
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v129 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v123) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v141 = int32(0)
	v142 = v129
	v145 = v129
	goto L29
L27:
	;
	v180 = v129
	v183 = v129
	goto L28
L28:
	;
	if v127 != 0 {
		goto L44
	} else {
		goto L45
	}
L29:
	;
	v153 = v128 + v145<<(uint(int32(2))%32)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+16))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v153)+8))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+16))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v153)+4))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+16))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)+16))
	if base.Ui32(v142) < base.Ui32(v161) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v180 = v169
	v183 = v171
	goto L28
L31:
	;
	v163 = v161
	goto L33
L32:
	;
	v163 = v142
	goto L33
L33:
	;
	if base.Ui32(v163) < base.Ui32(v159) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v165 = v159
	goto L36
L35:
	;
	v165 = v163
	goto L36
L36:
	;
	if base.Ui32(v165) < base.Ui32(v157) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v167 = v157
	goto L39
L38:
	;
	v167 = v165
	goto L39
L39:
	;
	if base.Ui32(v167) < base.Ui32(v155) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v169 = v155
	goto L42
L41:
	;
	v169 = v167
	goto L42
L42:
	;
	v170 = int32(4)
	v171 = v145 + v170
	v173 = v141 + v170
	if v173 != v123&int32(2147483644) {
		v141 = v173
		v142 = v169
		v145 = v171
		goto L29
	} else {
		goto L43
	}
L43:
	;
	goto L30
L44:
	;
	v189 = v129
	v194 = v180
	v197 = v183
	goto L47
L45:
	;
	v220 = v180
	goto L46
L46:
	;
	v245 = v220 + int32(1)
	goto L23
L47:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v128+v197<<(uint(int32(2))%32))))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v206)+16))
	if base.Ui32(v194) < base.Ui32(v207) {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v220 = v209
	goto L46
L49:
	;
	v209 = v207
	goto L51
L50:
	;
	v209 = v194
	goto L51
L51:
	;
	v210 = int32(1)
	v213 = v189 + v210
	if v213 != v127 {
		v189 = v213
		v194 = v209
		v197 = v197 + v210
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	v274 = v270
	goto L7
}
func F_transformUpdateTargetList(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
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
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	v12 = m.G0
	v13 = int32(16)
	v14 = v12 - v13
	m.G0 = v14
	v17 = F_transformTargetList(m, l0, l1, v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+120)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	if v24 <= v23 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v23 + int32(1)
	goto L5
L4:
	;
	goto L5
L5:
	;
	if l1 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v31 = v29
	goto L8
L7:
	;
	v31 = int32(0)
	goto L8
L8:
	;
	if v17 == int32(0) {
		v234 = v31
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L68
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L65
	}
L11:
	;
	if v234 != 0 {
		goto L9
	} else {
		goto L64
	}
L12:
	;
	v34 = int32(0)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v35 <= v34 {
		v234 = v31
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v44 = v34
	v46 = v31
	goto L14
L14:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v44<<(uint(int32(2))%32))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+26)))
	if v56 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v234 = v221
	goto L11
L16:
	;
	v225 = v44 + int32(1)
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v225 < v226 {
		v44 = v225
		v46 = v221
		goto L14
	} else {
		goto L63
	}
L17:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v59 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v55)+8)) = uint16(v59)
	v221 = v46
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v46 == int32(0) {
		goto L10
	} else {
		goto L20
	}
L20:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v72 = int32(0)
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v68)+48))
	v75 = int32(*(*int16)(unsafe.Add(mBase, uint32(v74)+120)))
	if v72 < v75 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	if v129 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L22:
	;
	v129 = v122
	goto L21
L23:
	;
	v122 = v81 + int32(1)
	goto L22
L24:
	;
	v81 = v72
	goto L27
L25:
	;
	goto L26
L26:
	;
	goto L35
L27:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v68)+52))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v90 = v83 + v84<<(uint(int32(4))%32) + v81*int32(100)
	v93 = F_namestrcmp(m, v90+int32(24), v70)
	mBase = m.M
	if v93 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+111)))
	if v96 != int32(1) {
		goto L23
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v100 = v81 + int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v68)+48))
	v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v101)+120)))
	if v100 < v102 {
		v81 = v100
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L28
L34:
	;
	v129 = int32(0)
	goto L21
L35:
	;
	v111 = F_SystemAttributeByName(m, v70)
	mBase = m.M
	if v111 == int32(0) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v114 = int32(*(*int16)(unsafe.Add(mBase, uint32(v111)+74)))
	if v114 != 0 {
		v122 = v114
		goto L22
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v198 = F_transformAssignedExpr(m, l0, v193, int32(17), v195, v129, v196, v197)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L58
	}
L41:
	;
	F_errcode(m, int32(50360452))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+48))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v140 + int32(4)
	F_errmsg(m, int32(76931), v14)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v149 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	F_parser_errposition(m, l0, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L56
	}
L45:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)+4))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155))))
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v152))))
	if v159 == int32(0) {
		v178 = v158
		v179 = v159
		goto L47
	} else {
		goto L48
	}
L46:
	;
	if v179-v178 != 0 {
		goto L44
	} else {
		goto L54
	}
L47:
	;
	goto L46
L48:
	;
	if v158 != v159 {
		v178 = v158
		v179 = v159
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v163 = v152
	v164 = v155
	goto L50
L50:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	if v168 == int32(0) {
		v178 = v167
		v179 = v168
		goto L47
	} else {
		goto L52
	}
L51:
	;
	v178 = v167
	v179 = v168
	goto L47
L52:
	;
	v171 = int32(1)
	if v167 == v168 {
		v163 = v163 + v171
		v164 = v164 + v171
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	F_errhint(m, int32(658979), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L55
	}
L55:
	;
	goto L44
L56:
	;
	F_errfinish(m, int32(518444), int32(2581), int32(81417))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55)+12)) = v195
	*(*uint16)(unsafe.Add(mBase, uint32(v55)+8)) = uint16(v129)
	*(*int32)(unsafe.Add(mBase, uint32(v55)+4)) = v198
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v39)+36))
	v206 = F_bms_add_member(m, v203, v129+int32(7))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v206
	v210 = v46 + int32(4)
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(v210) < base.Ui32(v212+v213<<(uint(int32(2))%32)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v218 = v210
	goto L62
L61:
	;
	v218 = int32(0)
	goto L62
L62:
	;
	v221 = v218
	goto L16
L63:
	;
	goto L15
L64:
	;
	m.G0 = v14 + int32(16)
	return v17
L65:
	;
	F_errmsg_internal(m, int32(221642), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(518444), int32(2567), int32(81417))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errmsg_internal(m, int32(221642), int32(0))
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(518444), int32(2595), int32(81417))
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
