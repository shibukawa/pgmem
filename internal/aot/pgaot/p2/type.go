package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecTypeFromExprList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v235 int32
	_ = v235
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = F_CreateTemplateTupleDesc(m, int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v103 = F_CreateTemplateTupleDesc(m, v102)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L25
	}
L4:
	;
	return int32(0)
L5:
	;
	v14 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	if v14 < v23 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	return v10
L7:
	;
	v27 = v10 + int32(28)
	v34 = v14
	v35 = v23
	v37 = v14
	goto L11
L8:
	;
	v91 = v14
	v98 = v23
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v91
	goto L6
L10:
	;
	v91 = v85
	v98 = v64
	goto L9
L11:
	;
	v43 = v27 + v23<<(uint(int32(3))%32) + v34*int32(100)
	v46 = v27 + v34<<(uint(int32(3))%32)
	if v23 != v35 {
		v64 = v35
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v85 = v23
	goto L10
L13:
	;
	v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+2)))
	if v65 <= int32(0) {
		v85 = v34
		goto L10
	} else {
		goto L21
	}
L14:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+7)))
	if v48 != int32(118) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v64 = v34
	goto L13
L16:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)))
	if v51 != int32(1) {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+6)))
	if v54&int32(6) != 0 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v46)+2)))
	if v57 <= int32(0) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+90)))
	if v60 != int32(118) {
		v64 = v23
		goto L13
	} else {
		goto L20
	}
L20:
	;
	goto L15
L21:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+90)))
	if v68 == int32(118) {
		v85 = v34
		goto L10
	} else {
		goto L22
	}
L22:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+5)))
	v77 = (v37 + v71 - int32(1)) & (int32(0) - v71)
	if int32(_a_F_ExecTypeFromExprList_0) < v77 {
		v85 = v34
		goto L10
	} else {
		goto L23
	}
L23:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v77)
	v83 = v34 + int32(1)
	if v83 != v23 {
		v34 = v83
		v35 = v64
		v37 = v77 + v65
		goto L11
	} else {
		goto L24
	}
L24:
	;
	goto L12
L25:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v106 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v111 = int32(1)
	v112 = v2
	goto L29
L27:
	;
	goto L28
L28:
	;
	v151 = int32(0)
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v151 < v160 {
		goto L38
	} else {
		goto L39
	}
L29:
	;
	v115 = base.I32_extend16_s(v111)
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v117+v112<<(uint(int32(2))%32))))
	v122 = F_exprType(m, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L31
	}
L30:
	;
	goto L28
L31:
	;
	v124 = F_exprTypmod(m, v121)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L32
	}
L32:
	;
	F_TupleDescInitEntry(m, v103, v115, int32(0), v122, v124, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v129 = F_exprCollation(m, v121)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v131<<(uint(int32(3))%32)+v115*int32(100))+24)) = v129
	goto L35
L35:
	;
	v139 = int32(1)
	v142 = v112 + v139
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v142 < v143 {
		v111 = v111 + v139
		v112 = v142
		goto L29
	} else {
		goto L36
	}
L36:
	;
	goto L30
L37:
	;
	return v103
L38:
	;
	v164 = v103 + int32(28)
	v171 = v151
	v172 = v160
	v174 = v151
	goto L42
L39:
	;
	v228 = v151
	v235 = v160
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+20)) = v235
	*(*int32)(unsafe.Add(mBase, uint32(v103)+16)) = v228
	goto L37
L41:
	;
	v228 = v222
	v235 = v201
	goto L40
L42:
	;
	v180 = v164 + v160<<(uint(int32(3))%32) + v171*int32(100)
	v183 = v164 + v171<<(uint(int32(3))%32)
	if v160 != v172 {
		v201 = v172
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v222 = v160
	goto L41
L44:
	;
	v202 = int32(*(*int16)(unsafe.Add(mBase, uint32(v183)+2)))
	if v202 <= int32(0) {
		v222 = v171
		goto L41
	} else {
		goto L52
	}
L45:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+7)))
	if v185 != int32(118) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v201 = v171
	goto L44
L47:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
	if v188 != int32(1) {
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+6)))
	if v191&int32(6) != 0 {
		goto L46
	} else {
		goto L49
	}
L49:
	;
	v194 = int32(*(*int16)(unsafe.Add(mBase, uint32(v183)+2)))
	if v194 <= int32(0) {
		goto L46
	} else {
		goto L50
	}
L50:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+90)))
	if v197 != int32(118) {
		v201 = v160
		goto L44
	} else {
		goto L51
	}
L51:
	;
	goto L46
L52:
	;
	v205 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+90)))
	if v205 == int32(118) {
		v222 = v171
		goto L41
	} else {
		goto L53
	}
L53:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+5)))
	v214 = (v174 + v208 - int32(1)) & (int32(0) - v208)
	if int32(_a_F_ExecTypeFromExprList_0) < v214 {
		v222 = v171
		goto L41
	} else {
		goto L54
	}
L54:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v183))) = uint16(v214)
	v220 = v171 + int32(1)
	if v220 != v160 {
		v171 = v220
		v172 = v201
		v174 = v214 + v202
		goto L42
	} else {
		goto L55
	}
L55:
	;
	goto L43
}
func F_ExecTypeFromTLInternal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
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
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v103 = F_CreateTemplateTupleDesc(m, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L27
	} else {
		goto L28
	}
L2:
	;
	v8 = int32(0)
	if l0 == v8 {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	if l0 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L5:
	;
	v102 = v96
	goto L1
L6:
	;
	v96 = int32(0)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v96 = int32(0)
	goto L5
L10:
	;
	goto L11
L11:
	;
	if v18 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v96 = v83
	goto L5
L13:
	;
	v24 = int32(0)
	if v24 < v18 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v65 = v8
	v66 = v8
	goto L15
L15:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71+v65<<(uint(int32(2))%32))))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+26)))
	v83 = v66 + (v76 ^ int32(1))
	goto L12
L16:
	;
	v27 = v18
	goto L18
L17:
	;
	v27 = v24
	goto L18
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v33 = int32(0)
	v36 = v33
	v37 = v33
	v38 = v8
	goto L19
L19:
	;
	v43 = int32(2)
	v45 = v32 + v37<<(uint(v43)%32)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+26)))
	v48 = int32(1)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+26)))
	v55 = v38 + (v47 ^ v48) + (v52 ^ v48)
	v57 = v37 + v43
	v59 = v36 + v43
	if v59 != v27&int32(2147483646) {
		v36 = v59
		v37 = v57
		v38 = v55
		goto L19
	} else {
		goto L21
	}
L20:
	;
	if v27&int32(1) == int32(0) {
		v83 = v55
		goto L12
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v65 = v57
	v66 = v55
	goto L15
L23:
	;
	v102 = v101
	goto L1
L24:
	;
	v101 = int32(0)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v101 = v100
	goto L23
L27:
	;
	return int32(0)
L28:
	;
	if l0 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v163 = int32(0)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	if v163 < v172 {
		goto L46
	} else {
		goto L47
	}
L30:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v110 <= int32(0) {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v117 = int32(1)
	v118 = int32(0)
	goto L32
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v120+v118<<(uint(int32(2))%32))))
	if l1 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L29
L34:
	;
	v153 = v118 + int32(1)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v153 < v154 {
		v117 = v150
		v118 = v153
		goto L32
	} else {
		goto L44
	}
L35:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124)+26)))
	if v125 != 0 {
		v150 = v117
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v126 = base.I32_extend16_s(v117)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v129 = F_exprType(m, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L27
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v132 = F_exprTypmod(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L27
	} else {
		goto L40
	}
L40:
	;
	F_TupleDescInitEntry(m, v103, v126, v127, v129, v132, int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L27
	} else {
		goto L41
	}
L41:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v138 = F_exprCollation(m, v137)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L27
	} else {
		goto L42
	}
L42:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v140<<(uint(int32(3))%32)+v126*int32(100))+24)) = v138
	goto L43
L43:
	;
	v150 = v117 + int32(1)
	goto L34
L44:
	;
	goto L33
L45:
	;
	return v103
L46:
	;
	v176 = v103 + int32(28)
	v183 = v163
	v184 = v172
	v186 = v163
	goto L50
L47:
	;
	v240 = v163
	v247 = v172
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v103)+20)) = v247
	*(*int32)(unsafe.Add(mBase, uint32(v103)+16)) = v240
	goto L45
L49:
	;
	v240 = v234
	v247 = v213
	goto L48
L50:
	;
	v192 = v176 + v172<<(uint(int32(3))%32) + v183*int32(100)
	v195 = v176 + v183<<(uint(int32(3))%32)
	if v172 != v184 {
		v213 = v184
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v234 = v172
	goto L49
L52:
	;
	v214 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195)+2)))
	if v214 <= int32(0) {
		v234 = v183
		goto L49
	} else {
		goto L60
	}
L53:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+7)))
	if v197 != int32(118) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v213 = v183
	goto L52
L55:
	;
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+4)))
	if v200 != int32(1) {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+6)))
	if v203&int32(6) != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	v206 = int32(*(*int16)(unsafe.Add(mBase, uint32(v195)+2)))
	if v206 <= int32(0) {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+90)))
	if v209 != int32(118) {
		v213 = v172
		goto L52
	} else {
		goto L59
	}
L59:
	;
	goto L54
L60:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192)+90)))
	if v217 == int32(118) {
		v234 = v183
		goto L49
	} else {
		goto L61
	}
L61:
	;
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195)+5)))
	v226 = (v186 + v220 - int32(1)) & (int32(0) - v220)
	if int32(_a_F_ExecTypeFromTLInternal_0) < v226 {
		v234 = v183
		goto L49
	} else {
		goto L62
	}
L62:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v195))) = uint16(v226)
	v232 = v183 + int32(1)
	if v232 != v172 {
		v183 = v232
		v184 = v213
		v186 = v226 + v214
		goto L50
	} else {
		goto L63
	}
L63:
	;
	goto L51
}
func F_TypeCategory(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	F_get_type_category_preferred(m, l0, v5+int32(15), v5+int32(14))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*int8)(unsafe.Add(mBase, uint32(v5)+15)))
		m.G0 = v5 + int32(16)
		return v15
	}
}
func F_TypeShellMake(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int64
	_ = v34
	var v38 int64
	_ = v38
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	v8 = m.G0
	v10 = v8 - int32(352)
	m.G0 = v10
	v14 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
		v17 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+88)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v10)+80)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v10)+64)) = v17
		v27 = int32(0)
		base.MemoryFill(m, v10+int32(96), v27, int32(256))
		v31 = F_strncpy(m, v10, l1, int32(64))
		mBase = m.M
		*(*uint8)(unsafe.Add(mBase, uint32(v31)+63)) = uint8(v27)
		v34 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+160)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+152)) = int64(80)
		v38 = int64(112)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+144)) = v38
		*(*int64)(unsafe.Add(mBase, uint32(v10)+136)) = int64(1)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+128)) = int64(4)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+120)) = base.I64_extend_i32_u(l3)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+112)) = base.I64_extend_i32_u(l2)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+104)) = base.I64_extend_i32_u(v10)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+168)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+184)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+176)) = int64(44)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+192)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+200)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+208)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+232)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+224)) = int64(2399)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+216)) = int64(2398)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+240)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+248)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+256)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+264)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+280)) = v38
		*(*int64)(unsafe.Add(mBase, uint32(v10)+272)) = int64(105)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+288)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+296)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+312)) = v34
		*(*int64)(unsafe.Add(mBase, uint32(v10)+304)) = int64(-1)
		*(*int64)(unsafe.Add(mBase, uint32(v10)+320)) = v34
		v90 = int32(257)
		*(*uint16)(unsafe.Add(mBase, uint32(v10)+93)) = uint16(v90)
		v92 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v10)+95)) = uint8(v92)
		v95 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_TypeShellMake[0])))
		if v95 == v92 {
			v99 = *(*int32)(unsafe.Add(mBase, _c_F_TypeShellMake[1]))
			if v99 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v154 = m.ExcPending
				if v154 != 0 {
					return
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v157 = m.ExcPending
					if v157 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_TypeShellMake_0), int32(0))
						mBase = m.M
						v161 = m.ExcPending
						if v161 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_TypeShellMake_1), int32(133), int32(_a_F_TypeShellMake_2))
							mBase = m.M
							v166 = m.ExcPending
							if v166 != 0 {
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
				*(*int32)(unsafe.Add(mBase, _c_F_TypeShellMake[1])) = int32(0)
				v109 = v99
				*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = base.I64_extend_i32_u(v109)
				v116 = F_heap_form_tuple(m, v16, v10+int32(96), v10-int32(-64))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return
				} else {
					F_CatalogTupleInsert(m, v14, v116)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						v121 = *(*int32)(unsafe.Add(mBase, _c_F_TypeShellMake[2]))
						if v121 != 0 {
							v122 = int32(0)
							F_GenerateTypeDependencies(m, v116, v14, v122, v122, v122, v122, v122, int32(1), v122)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								v132 = *(*int32)(unsafe.Add(mBase, _c_F_TypeShellMake[3]))
								if v132 != 0 {
									v134 = int32(0)
									F_RunObjectPostCreateHook(m, int32(1247), v109, v134, v134)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
										F_pfree(m, v116)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return
										} else {
											F_relation_close(m, v14, int32(3))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												m.G0 = v10 + int32(352)
												return
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
									F_pfree(m, v116)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return
									} else {
										F_relation_close(m, v14, int32(3))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return
										} else {
											m.G0 = v10 + int32(352)
											return
										}
									}
								}
							}
						} else {
							v132 = *(*int32)(unsafe.Add(mBase, _c_F_TypeShellMake[3]))
							if v132 != 0 {
								v134 = int32(0)
								F_RunObjectPostCreateHook(m, int32(1247), v109, v134, v134)
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
									F_pfree(m, v116)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return
									} else {
										F_relation_close(m, v14, int32(3))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return
										} else {
											m.G0 = v10 + int32(352)
											return
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
								F_pfree(m, v116)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return
								} else {
									F_relation_close(m, v14, int32(3))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return
									} else {
										m.G0 = v10 + int32(352)
										return
									}
								}
							}
						}
					}
				}
			}
		} else {
			v107 = F_GetNewOidWithIndex(m, v14, int32(2703), int32(1))
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return
			} else {
				v109 = v107
				*(*int64)(unsafe.Add(mBase, uint32(v10)+96)) = base.I64_extend_i32_u(v109)
				v116 = F_heap_form_tuple(m, v16, v10+int32(96), v10-int32(-64))
				mBase = m.M
				v117 = m.ExcPending
				if v117 != 0 {
					return
				} else {
					F_CatalogTupleInsert(m, v14, v116)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return
					} else {
						v121 = *(*int32)(unsafe.Add(mBase, _c_F_TypeShellMake[2]))
						if v121 != 0 {
							v122 = int32(0)
							F_GenerateTypeDependencies(m, v116, v14, v122, v122, v122, v122, v122, int32(1), v122)
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return
							} else {
								v132 = *(*int32)(unsafe.Add(mBase, _c_F_TypeShellMake[3]))
								if v132 != 0 {
									v134 = int32(0)
									F_RunObjectPostCreateHook(m, int32(1247), v109, v134, v134)
									mBase = m.M
									v137 = m.ExcPending
									if v137 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
										F_pfree(m, v116)
										mBase = m.M
										v144 = m.ExcPending
										if v144 != 0 {
											return
										} else {
											F_relation_close(m, v14, int32(3))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return
											} else {
												m.G0 = v10 + int32(352)
												return
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
									F_pfree(m, v116)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return
									} else {
										F_relation_close(m, v14, int32(3))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return
										} else {
											m.G0 = v10 + int32(352)
											return
										}
									}
								}
							}
						} else {
							v132 = *(*int32)(unsafe.Add(mBase, _c_F_TypeShellMake[3]))
							if v132 != 0 {
								v134 = int32(0)
								F_RunObjectPostCreateHook(m, int32(1247), v109, v134, v134)
								mBase = m.M
								v137 = m.ExcPending
								if v137 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
									*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
									F_pfree(m, v116)
									mBase = m.M
									v144 = m.ExcPending
									if v144 != 0 {
										return
									} else {
										F_relation_close(m, v14, int32(3))
										mBase = m.M
										v147 = m.ExcPending
										if v147 != 0 {
											return
										} else {
											m.G0 = v10 + int32(352)
											return
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
								*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
								F_pfree(m, v116)
								mBase = m.M
								v144 = m.ExcPending
								if v144 != 0 {
									return
								} else {
									F_relation_close(m, v14, int32(3))
									mBase = m.M
									v147 = m.ExcPending
									if v147 != 0 {
										return
									} else {
										m.G0 = v10 + int32(352)
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
}
func F_appendTypeNameToBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v6 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+13)))
	if v45 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v7 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = F_format_type_be(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L14
	} else {
		goto L18
	}
L5:
	;
	v14 = int32(0)
	goto L6
L6:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
	v18 = v16
	goto L10
L9:
	;
	v18 = int32(0)
	goto L10
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v22 = v19 + v14<<(uint(int32(2))%32)
	if v18 != v22 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	F_appendStringInfoChar(m, l1, int32(46))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	F_appendStringInfoString(m, l1, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return
L15:
	;
	goto L13
L16:
	;
	v32 = v14 + int32(1)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	if v32 < v33 {
		v14 = v32
		goto L6
	} else {
		goto L17
	}
L17:
	;
	goto L7
L18:
	;
	F_appendStringInfoString(m, l1, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L1
L20:
	;
	F_appendStringInfoString(m, l1, int32(_a_F_appendTypeNameToBuffer_0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L14
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v51 != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	goto L22
L24:
	;
	F_appendStringInfoString(m, l1, int32(_a_F_appendTypeNameToBuffer_1))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L14
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	return
L27:
	;
	goto L26
}
func F_assign_record_type_identifier(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l0 != int32(2249) {
		v12 = F_lookup_type_cache(m, l0, int32(256))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)+188))
			if v16 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						v55 = F_format_type_be(m, l0)
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v55
							F_errmsg(m, int32(_a_F_assign_record_type_identifier_0), v7)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_assign_record_type_identifier_1), int32(2175), int32(_a_F_assign_record_type_identifier_2))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int64(0)
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
				v19 = *(*int64)(unsafe.Add(mBase, uint32(v12)+192))
				v43 = v19
				m.G0 = v7 + int32(16)
				return v43
			}
		}
	} else {
		if l1 < int32(0) {
			v35 = int32(_a_F_assign_record_type_identifier_3)
			v37 = *(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_identifier[0]))
			v39 = v37 + int64(1)
			*(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_identifier[0])) = v39
			v43 = v39
		} else {
			v23 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_identifier[1]))
			if v23 <= l1 {
				v35 = int32(_a_F_assign_record_type_identifier_3)
				v37 = *(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_identifier[0]))
				v39 = v37 + int64(1)
				*(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_identifier[0])) = v39
				v43 = v39
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, _c_F_assign_record_type_identifier[2]))
				v29 = v26 + l1<<(uint(int32(4))%32)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
				if v30 == int32(0) {
					v35 = int32(_a_F_assign_record_type_identifier_3)
					v37 = *(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_identifier[0]))
					v39 = v37 + int64(1)
					*(*int64)(unsafe.Add(mBase, _c_F_assign_record_type_identifier[0])) = v39
					v43 = v39
				} else {
					v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)))
					v43 = v33
				}
			}
		}
		m.G0 = v7 + int32(16)
		return v43
	}
}
func F_findTypeSubscriptingFunction(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(2281)
	v10 = int32(1)
	v14 = F_LookupFuncName(m, l0, v10, v6+int32(44), v10)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int32(0)
	} else {
		if v14 != 0 {
			v18 = F_get_func_rettype(m, v14)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				if v18 != int32(2281) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(117833860))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							v57 = F_NameListToString(m, l0)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(_a_F_findTypeSubscriptingFunction_0)
								*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = v57
								F_errmsg(m, int32(_a_F_findTypeSubscriptingFunction_1), v6+int32(32))
								mBase = m.M
								v66 = m.ExcPending
								if v66 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_findTypeSubscriptingFunction_2), int32(2340), int32(_a_F_findTypeSubscriptingFunction_3))
									mBase = m.M
									v71 = m.ExcPending
									if v71 != 0 {
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
					if v14 == int32(_a_F_findTypeSubscriptingFunction_4) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(117833860))
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								v79 = F_NameListToString(m, l0)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v79
									F_errmsg(m, int32(_a_F_findTypeSubscriptingFunction_5), v6+int32(16))
									mBase = m.M
									v86 = m.ExcPending
									if v86 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_findTypeSubscriptingFunction_2), int32(2350), int32(_a_F_findTypeSubscriptingFunction_3))
										mBase = m.M
										v91 = m.ExcPending
										if v91 != 0 {
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
						m.G0 = v6 + int32(48)
						return v14
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(52461700))
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int32(0)
				} else {
					v39 = F_func_signature_string(m, l0, int32(1), int32(0), v6+int32(44))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6))) = v39
						F_errmsg(m, int32(_a_F_findTypeSubscriptingFunction_6), v6)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_findTypeSubscriptingFunction_2), int32(2334), int32(_a_F_findTypeSubscriptingFunction_3))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
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
func F_format_type_extended(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
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
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int64
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(80)
	return v284
L2:
	;
	v28 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L8
	} else {
		goto L10
	}
L3:
	;
	if l2&int32(8) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v284 = int32(0)
	goto L1
L5:
	;
	goto L6
L6:
	;
	if l2&int32(2) == int32(0) {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v22 = F_pstrdup(m, int32(_a_F_format_type_extended_13))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v284 = v22
	goto L1
L10:
	;
	if v28 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if l2&int32(8) != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
	v55 = v53 + v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+92))
	if v56 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v284 = int32(0)
	goto L1
L15:
	;
	goto L16
L16:
	;
	if l2&int32(2) != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v38 = F_pstrdup(m, int32(_a_F_format_type_extended_0))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L8
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L21
	}
L20:
	;
	v284 = v38
	goto L1
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = l0
	F_errmsg_internal(m, int32(_a_F_format_type_extended_1), v12)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_format_type_extended_2), int32(137), int32(_a_F_format_type_extended_3))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L23
	}
L23:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L24:
	;
	v106 = l2 & int32(1)
	v109 = base.B2i32(int32(0) <= l1) & l2
	if v101 <= int32(1082) {
		goto L68
	} else {
		goto L69
	}
L25:
	;
	v101 = l0
	v102 = v55
	v103 = v28
	v104 = v4
	goto L24
L26:
	;
	goto L27
L27:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+88))
	if v59 != int32(_a_F_format_type_extended_14) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v101 = l0
	v102 = v55
	v103 = v28
	v104 = v4
	goto L24
L29:
	;
	goto L30
L30:
	;
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+129)))
	if v62 == int32(112) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v101 = l0
	v102 = v55
	v103 = v28
	v104 = v4
	goto L24
L32:
	;
	goto L33
L33:
	;
	F_ReleaseCatCache(m, v28)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L34
	}
L34:
	;
	v69 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v56))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L8
	} else {
		goto L35
	}
L35:
	;
	if v69 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if l2&int32(8) != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96)+22)))
	v101 = v56
	v102 = v96 + v97
	v103 = v69
	v104 = int32(1)
	goto L24
L39:
	;
	v284 = int32(0)
	goto L1
L40:
	;
	goto L41
L41:
	;
	if l2&int32(2) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v79 = F_pstrdup(m, int32(_a_F_format_type_extended_15))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L46
	}
L45:
	;
	v284 = v79
	goto L1
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = l0
	F_errmsg_internal(m, int32(_a_F_format_type_extended_1), v12-int32(-64))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(_a_F_format_type_extended_2), int32(162), int32(_a_F_format_type_extended_3))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	if v104 != 0 {
		goto L157
	} else {
		goto L158
	}
L50:
	;
	if l2&int32(4) == int32(0) {
		goto L143
	} else {
		goto L144
	}
L51:
	;
	if v231 != 0 {
		v273 = v231
		goto L49
	} else {
		goto L141
	}
L52:
	;
	if v101 != int32(114) {
		goto L50
	} else {
		goto L139
	}
L53:
	;
	if v109 != 0 {
		goto L134
	} else {
		goto L135
	}
L54:
	;
	if v109 != 0 {
		goto L129
	} else {
		goto L130
	}
L55:
	;
	if v109 != 0 {
		goto L124
	} else {
		goto L125
	}
L56:
	;
	v203 = F_pstrdup(m, int32(_a_F_format_type_extended_16))
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L8
	} else {
		goto L123
	}
L57:
	;
	if v109 != 0 {
		goto L118
	} else {
		goto L119
	}
L58:
	;
	if v109 != 0 {
		goto L113
	} else {
		goto L114
	}
L59:
	;
	if v109 != 0 {
		goto L108
	} else {
		goto L109
	}
L60:
	;
	v179 = F_pstrdup(m, int32(_a_F_format_type_extended_17))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L8
	} else {
		goto L107
	}
L61:
	;
	v176 = F_pstrdup(m, int32(_a_F_format_type_extended_8))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L8
	} else {
		goto L106
	}
L62:
	;
	v173 = F_pstrdup(m, int32(_a_F_format_type_extended_10))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L8
	} else {
		goto L105
	}
L63:
	;
	v170 = F_pstrdup(m, int32(_a_F_format_type_extended_9))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L8
	} else {
		goto L104
	}
L64:
	;
	v167 = F_pstrdup(m, int32(_a_F_format_type_extended_18))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L103
	}
L65:
	;
	v164 = F_pstrdup(m, int32(_a_F_format_type_extended_12))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L102
	}
L66:
	;
	if v109 != 0 {
		goto L96
	} else {
		goto L97
	}
L67:
	;
	v154 = F_pstrdup(m, int32(_a_F_format_type_extended_4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L8
	} else {
		goto L95
	}
L68:
	;
	if v101 <= int32(699) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if v101 <= int32(1265) {
		goto L75
	} else {
		goto L76
	}
L71:
	;
	switch v101 - int32(16) {
	case 0:
		goto L67
	case 1, 2, 3, 6:
		goto L50
	case 4:
		goto L61
	case 5:
		goto L63
	case 7:
		goto L62
	default:
		goto L52
	}
L72:
	;
	goto L73
L73:
	;
	switch v101 - int32(700) {
	case 0:
		goto L65
	case 1:
		goto L64
	default:
		goto L74
	}
L74:
	;
	switch v101 - int32(1042) {
	case 0:
		goto L66
	case 1:
		goto L53
	default:
		goto L50
	}
L75:
	;
	switch v101 - int32(1184) {
	case 0:
		goto L55
	case 1:
		goto L50
	case 2:
		goto L59
	default:
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	switch v101 - int32(1560) {
	case 0:
		goto L83
	case 1:
		goto L50
	case 2:
		goto L54
	default:
		goto L84
	}
L78:
	;
	if v101 == int32(1083) {
		goto L58
	} else {
		goto L79
	}
L79:
	;
	if v101 != int32(1114) {
		goto L50
	} else {
		goto L80
	}
L80:
	;
	if v109 == int32(0) {
		goto L56
	} else {
		goto L81
	}
L81:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v132 = F_printTypmod(m, int32(_a_F_format_type_extended_19), l1, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	v231 = v132
	goto L51
L83:
	;
	if v109 != 0 {
		goto L89
	} else {
		goto L90
	}
L84:
	;
	if v101 == int32(1266) {
		goto L57
	} else {
		goto L85
	}
L85:
	;
	if v101 != int32(1700) {
		goto L50
	} else {
		goto L86
	}
L86:
	;
	if v109 == int32(0) {
		goto L60
	} else {
		goto L87
	}
L87:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v144 = F_printTypmod(m, int32(_a_F_format_type_extended_17), l1, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L88
	}
L88:
	;
	v231 = v144
	goto L51
L89:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v148 = F_printTypmod(m, int32(_a_F_format_type_extended_20), l1, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L8
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	if v106 != 0 {
		goto L50
	} else {
		goto L93
	}
L92:
	;
	v231 = v148
	goto L51
L93:
	;
	v151 = F_pstrdup(m, int32(_a_F_format_type_extended_20))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L8
	} else {
		goto L94
	}
L94:
	;
	v231 = v151
	goto L51
L95:
	;
	v231 = v154
	goto L51
L96:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v158 = F_printTypmod(m, int32(_a_F_format_type_extended_21), l1, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L8
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	if v106 != 0 {
		goto L50
	} else {
		goto L100
	}
L99:
	;
	v231 = v158
	goto L51
L100:
	;
	v161 = F_pstrdup(m, int32(_a_F_format_type_extended_21))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	v231 = v161
	goto L51
L102:
	;
	v231 = v164
	goto L51
L103:
	;
	v231 = v167
	goto L51
L104:
	;
	v231 = v170
	goto L51
L105:
	;
	v231 = v173
	goto L51
L106:
	;
	v231 = v176
	goto L51
L107:
	;
	v231 = v179
	goto L51
L108:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v183 = F_printTypmod(m, int32(_a_F_format_type_extended_22), l1, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L8
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v186 = F_pstrdup(m, int32(_a_F_format_type_extended_22))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L8
	} else {
		goto L112
	}
L111:
	;
	v231 = v183
	goto L51
L112:
	;
	v231 = v186
	goto L51
L113:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v190 = F_printTypmod(m, int32(_a_F_format_type_extended_23), l1, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L8
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	v193 = F_pstrdup(m, int32(_a_F_format_type_extended_24))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L8
	} else {
		goto L117
	}
L116:
	;
	v231 = v190
	goto L51
L117:
	;
	v231 = v193
	goto L51
L118:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v197 = F_printTypmod(m, int32(_a_F_format_type_extended_23), l1, v196)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v200 = F_pstrdup(m, int32(_a_F_format_type_extended_25))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L8
	} else {
		goto L122
	}
L121:
	;
	v231 = v197
	goto L51
L122:
	;
	v231 = v200
	goto L51
L123:
	;
	v231 = v203
	goto L51
L124:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v207 = F_printTypmod(m, int32(_a_F_format_type_extended_19), l1, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L8
	} else {
		goto L127
	}
L125:
	;
	goto L126
L126:
	;
	v210 = F_pstrdup(m, int32(_a_F_format_type_extended_26))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L8
	} else {
		goto L128
	}
L127:
	;
	v231 = v207
	goto L51
L128:
	;
	v231 = v210
	goto L51
L129:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v214 = F_printTypmod(m, int32(_a_F_format_type_extended_27), l1, v213)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L8
	} else {
		goto L132
	}
L130:
	;
	goto L131
L131:
	;
	v217 = F_pstrdup(m, int32(_a_F_format_type_extended_27))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L8
	} else {
		goto L133
	}
L132:
	;
	v231 = v214
	goto L51
L133:
	;
	v231 = v217
	goto L51
L134:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	v221 = F_printTypmod(m, int32(_a_F_format_type_extended_28), l1, v220)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L8
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v224 = F_pstrdup(m, int32(_a_F_format_type_extended_28))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L8
	} else {
		goto L138
	}
L137:
	;
	v231 = v221
	goto L51
L138:
	;
	v231 = v224
	goto L51
L139:
	;
	v229 = F_pstrdup(m, int32(_a_F_format_type_extended_11))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L8
	} else {
		goto L140
	}
L140:
	;
	v231 = v229
	goto L51
L141:
	;
	goto L50
L142:
	;
	v248 = F_quote_qualified_identifier(m, v245, v102+int32(4))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L8
	} else {
		goto L149
	}
L143:
	;
	v237 = int32(0)
	v239 = F_TypeIsVisibleExt(m, v101, v237)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L8
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v102)+68))
	v243 = F_get_namespace_name_or_temp(m, v242)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L8
	} else {
		goto L148
	}
L146:
	;
	if v239 != 0 {
		v245 = v237
		goto L142
	} else {
		goto L147
	}
L147:
	;
	goto L145
L148:
	;
	v245 = v243
	goto L142
L149:
	;
	if v109 == int32(0) {
		v273 = v248
		goto L49
	} else {
		goto L150
	}
L150:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v102)+120))
	if v252 == int32(0) {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v248
	v260 = F_psprintf(m, int32(_a_F_format_type_extended_6), v12+int32(32))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L8
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v264 = F_OidFunctionCall1Coll(m, v252, int32(0), base.I64_extend_i32_u(l1))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L8
	} else {
		goto L155
	}
L154:
	;
	v273 = v260
	goto L49
L155:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+52)) = uint32(v264)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v248
	v271 = F_psprintf(m, int32(_a_F_format_type_extended_7), v12+int32(48))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L8
	} else {
		goto L156
	}
L156:
	;
	v273 = v271
	goto L49
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v273
	v279 = F_psprintf(m, int32(_a_F_format_type_extended_5), v12+int32(16))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L8
	} else {
		goto L160
	}
L158:
	;
	v281 = v273
	goto L159
L159:
	;
	F_ReleaseCatCache(m, v103)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L8
	} else {
		goto L161
	}
L160:
	;
	v281 = v279
	goto L159
L161:
	;
	v284 = v281
	goto L1
}
func F_getTypeBinaryOutputInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v14 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		if v14 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+22)))
			v18 = v16 + v17
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+82)))
			if v19 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return
					} else {
						v58 = F_format_type_be(m, l0)
						mBase = m.M
						v59 = m.ExcPending
						if v59 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v58
							F_errmsg(m, int32(_a_F_getTypeBinaryOutputInfo_0), v10+int32(32))
							mBase = m.M
							v65 = m.ExcPending
							if v65 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_getTypeBinaryOutputInfo_1), int32(3301), int32(_a_F_getTypeBinaryOutputInfo_2))
								mBase = m.M
								v70 = m.ExcPending
								if v70 != 0 {
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
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+112))
				if v22 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							v78 = F_format_type_be(m, l0)
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v78
								F_errmsg(m, int32(_a_F_getTypeBinaryOutputInfo_3), v10+int32(16))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_getTypeBinaryOutputInfo_1), int32(3306), int32(_a_F_getTypeBinaryOutputInfo_2))
									mBase = m.M
									v90 = m.ExcPending
									if v90 != 0 {
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
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+78)))
					if v26 != 0 {
						v31 = int32(0)
					} else {
						v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+76)))
						v31 = base.B2i32(v28 == int32(_a_F_getTypeBinaryOutputInfo_4))
					}
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v31)
					F_ReleaseCatCache(m, v14)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return
					} else {
						m.G0 = v10 + int32(48)
						return
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10))) = l0
				F_errmsg_internal(m, int32(_a_F_getTypeBinaryOutputInfo_5), v10)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_getTypeBinaryOutputInfo_1), int32(3294), int32(_a_F_getTypeBinaryOutputInfo_2))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
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
func F_getTypeInputInfo(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
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
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v13 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(l0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v13 != 0 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+22)))
			v17 = v15 + v16
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+82)))
			if v18 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return
				} else {
					F_errcode(m, int32(67137668))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return
					} else {
						v57 = F_format_type_be(m, l0)
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v57
							F_errmsg(m, int32(_a_F_getTypeInputInfo_0), v9+int32(32))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_getTypeInputInfo_1), int32(3202), int32(_a_F_getTypeInputInfo_2))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
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
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+100))
				if v21 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return
					} else {
						F_errcode(m, int32(52461700))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							v77 = F_format_type_be(m, l0)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v77
								F_errmsg(m, int32(_a_F_getTypeInputInfo_3), v9+int32(16))
								mBase = m.M
								v84 = m.ExcPending
								if v84 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_getTypeInputInfo_1), int32(3207), int32(_a_F_getTypeInputInfo_2))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
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
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v21
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
					v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
					v27 = v25 + v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+92))
					if v28 != 0 {
						v30 = v28
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
						v30 = v29
					}
					*(*int32)(unsafe.Add(mBase, uint32(l2))) = v30
					F_ReleaseCatCache(m, v13)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_getTypeInputInfo_4), v9)
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_getTypeInputInfo_1), int32(3195), int32(_a_F_getTypeInputInfo_2))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
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
func F_get_record_type_from_argument(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = F_get_fn_expr_argtype(m, v9, int32(0))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l2))) = v11
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
		F_prepare_column_cache(m, l2+int32(4), v11, int32(-1), v17, int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
			if v21|int32(32) != int32(99) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					F_errcode(m, int32(67141764))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l1
						F_errmsg(m, int32(_a_F_get_record_type_from_argument_0), v7)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_get_record_type_from_argument_1), int32(3650), int32(_a_F_get_record_type_from_argument_2))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
				m.G0 = v7 + int32(16)
				return
			}
		}
	}
}
func F_parseTypeString(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
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
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = F_typeStringToTypeName(m, l0, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if v11 == int32(0) {
			v92 = int32(0)
			m.G0 = v9 + int32(48)
			return v92
		} else {
			if l3 != 0 {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
				v24 = base.B2i32(v20 == int32(453))
			} else {
				v24 = int32(0)
			}
			v25 = F_LookupTypeNameExtended(m, int32(0), v11, l2, int32(1), v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 == int32(0) {
					v29 = int32(0)
					v30 = F_errsave_start(m, l3)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int32(0)
					} else {
						if v30 == int32(0) {
							v92 = v29
							m.G0 = v9 + int32(48)
							return v92
						} else {
							F_errcode(m, int32(67137668))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v38 = v9 + int32(32)
								F_initStringInfo(m, v38)
								mBase = m.M
								v40 = m.ExcPending
								if v40 != 0 {
									return int32(0)
								} else {
									F_appendTypeNameToBuffer(m, v11, v38)
									mBase = m.M
									v42 = m.ExcPending
									if v42 != 0 {
										return int32(0)
									} else {
										v43 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
										*(*int32)(unsafe.Add(mBase, uint32(v9))) = v43
										F_errmsg(m, int32(_a_F_parseTypeString_0), v9)
										mBase = m.M
										v47 = m.ExcPending
										if v47 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, l3, int32(_a_F_parseTypeString_1), int32(802), int32(_a_F_parseTypeString_2))
											mBase = m.M
											v52 = m.ExcPending
											if v52 != 0 {
												return int32(0)
											} else {
												v92 = v29
												m.G0 = v9 + int32(48)
												return v92
											}
										}
									}
								}
							}
						}
					}
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v25)+16))
					v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+22)))
					v55 = v53 + v54
					v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+82)))
					if v56 == int32(0) {
						F_ReleaseCatCache(m, v25)
						mBase = m.M
						v60 = m.ExcPending
						if v60 != 0 {
							return int32(0)
						} else {
							v61 = int32(0)
							v62 = F_errsave_start(m, l3)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return int32(0)
							} else {
								if v62 == int32(0) {
									v92 = v61
									m.G0 = v9 + int32(48)
									return v92
								} else {
									F_errcode(m, int32(67137668))
									mBase = m.M
									v68 = m.ExcPending
									if v68 != 0 {
										return int32(0)
									} else {
										v70 = v9 + int32(32)
										F_initStringInfo(m, v70)
										mBase = m.M
										v72 = m.ExcPending
										if v72 != 0 {
											return int32(0)
										} else {
											F_appendTypeNameToBuffer(m, v11, v70)
											mBase = m.M
											v74 = m.ExcPending
											if v74 != 0 {
												return int32(0)
											} else {
												v75 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
												*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v75
												F_errmsg(m, int32(_a_F_parseTypeString_3), v9+int32(16))
												mBase = m.M
												v81 = m.ExcPending
												if v81 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, l3, int32(_a_F_parseTypeString_1), int32(814), int32(_a_F_parseTypeString_2))
													mBase = m.M
													v86 = m.ExcPending
													if v86 != 0 {
														return int32(0)
													} else {
														v92 = v61
														m.G0 = v9 + int32(48)
														return v92
													}
												}
											}
										}
									}
								}
							}
						}
					} else {
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
						*(*int32)(unsafe.Add(mBase, uint32(l1))) = v87
						F_ReleaseCatCache(m, v25)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							v92 = int32(1)
							m.G0 = v9 + int32(48)
							return v92
						}
					}
				}
			}
		}
	}
}
func F_typeDepNeeded(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v22 = int32(1)
	goto L2
L1:
	;
	m.G0 = v10 + int32(16)
	return v88
L2:
	;
	if base.B2i32(int32(0)|base.B2i32(base.Ui32(int32(_a_F_typeDepNeeded_0)) < base.Ui32(l0)) == v3)&((v22|base.B2i32(l0 != int32(2200)))&v22) != 0 {
		v88 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = int32(1)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v31 == v30 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_get_func_signature(m, v34, v10+int32(12), v10+int32(8))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	F_op_input_types(m, v73, v10+int32(12), v10+int32(8))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L7
	} else {
		goto L16
	}
L7:
	;
	return int32(0)
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v44 <= int32(0) {
		v67 = v30
		goto L9
	} else {
		goto L10
	}
L9:
	;
	F_pfree(m, v43)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L7
	} else {
		goto L15
	}
L10:
	;
	v49 = int32(0)
	goto L11
L11:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v43+v49<<(uint(int32(2))%32))))
	v59 = base.B2i32(l0 != v58)
	if l0 == v58 {
		v67 = v59
		goto L9
	} else {
		goto L13
	}
L12:
	;
	v67 = v59
	goto L9
L13:
	;
	v62 = v49 + int32(1)
	if v62 != v44 {
		v49 = v62
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v88 = v67
	goto L1
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	v88 = base.B2i32(l0 != v80) & base.B2i32(l0 != v82)
	goto L1
}
