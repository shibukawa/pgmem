package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_QualifiedNameGetCreationNamespace(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_DeconstructQualifiedName(m, l0, v6+int32(12), l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	if v14 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v6 + int32(16)
	return v93
L4:
	;
	F_AccessTempTableNamespace(m, base.B2i32(v14 == int32(0)))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L30
	}
L5:
	;
	v15 = int32(224539)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _consts[429])))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v19 == int32(0) {
		v38 = v18
		v39 = v19
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L7
L7:
	;
	F_recomputeNamespacePath(m)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L23
	}
L8:
	;
	if v39-v38 == int32(0) {
		goto L4
	} else {
		goto L16
	}
L9:
	;
	goto L8
L10:
	;
	if v18 != v19 {
		v38 = v18
		v39 = v19
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v23 = v14
	v24 = v15
	goto L12
L12:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v27
		v39 = v28
		goto L9
	} else {
		goto L14
	}
L13:
	;
	v38 = v27
	v39 = v28
	goto L9
L14:
	;
	v31 = int32(1)
	if v27 == v28 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
L16:
	;
	v44 = int32(0)
	v47 = F_GetSysCacheOid(m, int32(37), v14, v44, v44, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v47 != 0 {
		v93 = v47
		goto L3
	} else {
		goto L18
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = v14
	F_errmsg(m, int32(69579), v6)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(477805), int32(3547), int32(415679))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L23:
	;
	v68 = int32(*(*uint8)(unsafe.Add(mBase, _consts[430])))
	if v68 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _consts[431]))
	if v70 != 0 {
		v93 = v70
		goto L3
	} else {
		goto L25
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_errmsg(m, int32(267151), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(477805), int32(3525), int32(399872))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[126]))
	v93 = v92
	goto L3
}
func F_quote_literal_cstr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	v2 = int32(0)
	if l0&int32(3) == v2 {
		v33 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v71 = F_palloc(m, v66<<(uint(int32(1))%32)+int32(4))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v66 = v58 - l0
	goto L1
L3:
	;
	v37 = v33
	goto L12
L4:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v66 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v22 = l0
	goto L8
L8:
	;
	v26 = v22 + int32(1)
	if v26&int32(3) == int32(0) {
		v33 = v26
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v58 = v26
	goto L2
L10:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v31 != 0 {
		v22 = v26
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v46 = int32(-2139062144)
	if (int32(16843008)-v43|v43)&v46 == v46 {
		v37 = v37 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v52 = v37
	goto L15
L14:
	;
	goto L13
L15:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v56 != 0 {
		v52 = v52 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v58 = v52
	goto L2
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	v75 = l0 + v66
	if base.Ui32(v75) <= base.Ui32(l0) {
		v97 = v71
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v105 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v105)
	v108 = v97 + int32(1)
	if v66 == int32(0) {
		v235 = v97
		v236 = v108
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v78 = l0
	goto L23
L22:
	;
	v92 = int32(69)
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v92)
	v97 = v71 + int32(1)
	goto L20
L23:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v86 == int32(92) {
		goto L22
	} else {
		goto L25
	}
L24:
	;
	v97 = v71
	goto L20
L25:
	;
	v90 = v78 + int32(1)
	if v90 != v75 {
		v78 = v90
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v243 = int32(39)
	*(*uint8)(unsafe.Add(mBase, uint32(v236))) = uint8(v243)
	v245 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v235)+2)) = uint8(v245)
	return v71
L28:
	;
	v112 = v66 & int32(3)
	if v112 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	if base.Ui32(v66) < base.Ui32(int32(4)) {
		v235 = v146
		v236 = v147
		goto L27
	} else {
		goto L40
	}
L30:
	;
	v145 = l0
	v146 = v97
	v147 = v108
	v151 = v66
	goto L29
L31:
	;
	goto L32
L32:
	;
	v115 = l0
	v116 = v97
	v117 = v108
	v120 = v2
	v121 = v66
	goto L33
L33:
	;
	v125 = v121 - int32(1)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	if v126 == int32(92) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v145 = v141
	v146 = v136
	v147 = v139
	v151 = v125
	goto L29
L35:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v136))) = uint8(v135)
	v138 = int32(1)
	v139 = v136 + v138
	v141 = v115 + v138
	v143 = v120 + v138
	if v143 != v112 {
		v115 = v141
		v116 = v136
		v117 = v139
		v120 = v143
		v121 = v125
		goto L33
	} else {
		goto L39
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v117))) = uint8(v126)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
	v135 = v132
	v136 = v116 + int32(2)
	goto L35
L37:
	;
	if v126 == int32(39) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v135 = v126
	v136 = v117
	goto L35
L39:
	;
	goto L34
L40:
	;
	v156 = v145
	v157 = v146
	v158 = v147
	v162 = v151
	goto L41
L41:
	;
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	if base.B2i32(v165 != int32(92))&base.B2i32(v165 != int32(39)) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v235 = v226
	v236 = v229
	goto L27
L43:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v165)
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v156))))
	v177 = v157 + int32(2)
	v178 = v174
	goto L45
L44:
	;
	v177 = v158
	v178 = v165
	goto L45
L45:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v177))) = uint8(v178)
	v181 = v156 + int32(1)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	if v182 == int32(92) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v194))) = uint8(v193)
	v197 = v156 + int32(2)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	if v198 == int32(92) {
		goto L51
	} else {
		goto L52
	}
L47:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v177)+1)) = uint8(v182)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181))))
	v193 = v190
	v194 = v177 + int32(2)
	goto L46
L48:
	;
	if v182 == int32(39) {
		goto L47
	} else {
		goto L49
	}
L49:
	;
	v193 = v182
	v194 = v177 + int32(1)
	goto L46
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v210))) = uint8(v209)
	v213 = v156 + int32(3)
	v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	if v214 == int32(92) {
		goto L55
	} else {
		goto L56
	}
L51:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)) = uint8(v198)
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197))))
	v209 = v206
	v210 = v194 + int32(2)
	goto L50
L52:
	;
	if v198 == int32(39) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v209 = v198
	v210 = v194 + int32(1)
	goto L50
L54:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v226))) = uint8(v225)
	v229 = v226 + int32(1)
	v230 = int32(4)
	v233 = v162 - v230
	if v233 != 0 {
		v156 = v156 + v230
		v157 = v226
		v158 = v229
		v162 = v233
		goto L41
	} else {
		goto L58
	}
L55:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v210)+1)) = uint8(v214)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213))))
	v225 = v222
	v226 = v210 + int32(2)
	goto L54
L56:
	;
	if v214 == int32(39) {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v225 = v214
	v226 = v210 + int32(1)
	goto L54
L58:
	;
	goto L42
}
