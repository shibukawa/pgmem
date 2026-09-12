package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_MarkAsPreparingGuts(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int64, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int64
	_ = v85
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	v24 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v29 = v25 + v26*int32(640)
	if v29&int32(3) == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = int64(0)
	v71 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+56))
	if v72 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	if base.Ui32(v29+int32(640)) <= base.Ui32(v29) {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v62 = F__emscripten_memset_bulkmem(m, v29+int32(8), base.I32_extend8_s(int32(0)), int32(632))
	mBase = m.M
	goto L10
L5:
	;
	v40 = int32(640)
	v41 = v26 * v40
	v42 = v41 + v25
	v44 = v42 + v40
	v46 = v42 + int32(4)
	if base.Ui32(v46) < base.Ui32(v44) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v48 = v44
	goto L8
L7:
	;
	v48 = v46
	goto L8
L8:
	;
	v56 = F__emscripten_memset_bulkmem(m, v29, base.I32_extend8_s(int32(0)), (v25^int32(-1)+v48-v41)&int32(-4)+int32(4))
	mBase = m.M
	goto L9
L9:
	;
	goto L1
L10:
	;
	goto L1
L11:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+124)) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+120)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v29)+36)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v29)+52)) = v78
	v85 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+112)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v29)+92)) = v85
	*(*uint16)(unsafe.Add(mBase, uint32(v29)+74)) = uint16(v79)
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+72)) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+68)) = v79
	*(*int32)(unsafe.Add(mBase, uint32(v29)+64)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v29)+60)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v29)+44)) = v79
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+277)) = uint8(v79)
	v102 = v29 + int32(268)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+272)) = v102
	v105 = v29 + int32(260)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+264)) = v105
	v108 = v29 + int32(252)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+256)) = v108
	v111 = v29 + int32(244)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+248)) = v111
	v114 = v29 + int32(236)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+240)) = v114
	v117 = v29 + int32(228)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+232)) = v117
	v120 = v29 + int32(220)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+224)) = v120
	v123 = v29 + int32(212)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+216)) = v123
	v126 = v29 + int32(204)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+208)) = v126
	v129 = v29 + int32(196)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+200)) = v129
	v132 = v29 + int32(188)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+192)) = v132
	v135 = v29 + int32(180)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+184)) = v135
	v138 = v29 + int32(172)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+176)) = v138
	v141 = v29 + int32(164)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+168)) = v141
	v144 = v29 + int32(156)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+160)) = v144
	v147 = v29 + int32(148)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+152)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v29)+148)) = v147
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v108))) = v108
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v114
	*(*int32)(unsafe.Add(mBase, uint32(v117))) = v117
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v120
	*(*int32)(unsafe.Add(mBase, uint32(v123))) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = v129
	*(*int32)(unsafe.Add(mBase, uint32(v132))) = v132
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v135
	*(*int32)(unsafe.Add(mBase, uint32(v138))) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v141))) = v141
	*(*int32)(unsafe.Add(mBase, uint32(v144))) = v144
	*(*uint8)(unsafe.Add(mBase, uint32(v29)+276)) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+36)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = l3
	v171 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+46)) = uint8(v79)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v79)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v171
	v178 = l0 + int32(47)
	if (l2^v178)&int32(3) != 0 {
		goto L18
	} else {
		goto L19
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = v72
	v75 = *(*int32)(unsafe.Add(mBase, _consts[157]))
	v78 = v75
	goto L11
L13:
	;
	goto L14
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29)+56)) = l1
	v78 = int32(-1)
	goto L11
L15:
	;
	*(*int32)(unsafe.Add(mBase, _consts[158])) = l0
	return
L16:
	;
	goto L15
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v233))) = uint8(v232)
	if v232&int32(255) == int32(0) {
		goto L16
	} else {
		goto L32
	}
L18:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v231 = l2
	v232 = v184
	v233 = v178
	goto L17
L19:
	;
	goto L20
L20:
	;
	if l2&int32(3) != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v188 = l2
	v190 = v178
	goto L24
L22:
	;
	v202 = l2
	v204 = v178
	goto L23
L23:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	v209 = int32(-2139062144)
	if (int32(16843008)-v206|v206)&v209 != v209 {
		v231 = v202
		v232 = v206
		v233 = v204
		goto L17
	} else {
		goto L28
	}
L24:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	*(*uint8)(unsafe.Add(mBase, uint32(v190))) = uint8(v191)
	if v191 == int32(0) {
		goto L16
	} else {
		goto L26
	}
L25:
	;
	v202 = v198
	v204 = v196
	goto L23
L26:
	;
	v195 = int32(1)
	v196 = v190 + v195
	v198 = v188 + v195
	if v198&int32(3) != 0 {
		v188 = v198
		v190 = v196
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	v214 = v202
	v215 = v206
	v216 = v204
	goto L29
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v216))) = v215
	v218 = int32(4)
	v219 = v216 + v218
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	v222 = v214 + v218
	v226 = int32(-2139062144)
	if (v220|(int32(16843008)-v220))&v226 == v226 {
		v214 = v222
		v215 = v220
		v216 = v219
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v231 = v222
	v232 = v220
	v233 = v219
	goto L17
L31:
	;
	goto L30
L32:
	;
	v240 = v231
	v242 = v233
	goto L33
L33:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v240)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v242)+1)) = uint8(v243)
	v245 = int32(1)
	if v243 != 0 {
		v240 = v240 + v245
		v242 = v242 + v245
		goto L33
	} else {
		goto L35
	}
L34:
	;
	goto L16
L35:
	;
	goto L34
}
func F___math_invalid(m *base.Module, l0 float64) float64 {
	var v2 float64
	_ = v2
	v2 = base.F64_sub(l0, l0)
	return base.F64_div(v2, v2)
}
func F___math_uflow(m *base.Module, l0 int32) float64 {
	var v2 float64
	_ = v2
	var v4 float64
	_ = v4
	var v5 float64
	_ = v5
	v2 = float64(1.2882297539194267e-231)
	if l0 != 0 {
		v4 = base.F64_neg(v2)
	} else {
		v4 = v2
	}
	v5 = F_fp_barrier_1(m, v4)
	return base.F64_mul(v2, v5)
}
func F__mdfd_openseg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	v9 = m.G0
	v11 = v9 - int32(176)
	m.G0 = v11
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	F_GetRelationPath(m, v11+int32(93), v15, v16, v17, v18, l1)
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
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	goto L30
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(93)
	v30 = F_pg_sprintf(m, v11+int32(10), int32(38675), v11)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v33 = v11 + int32(10)
	v35 = v11 + int32(93)
	if (v35^v33)&int32(3) != 0 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	goto L3
L8:
	;
	goto L3
L9:
	;
	goto L8
L10:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v90))) = uint8(v89)
	if v89&int32(255) == int32(0) {
		goto L9
	} else {
		goto L25
	}
L11:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	v88 = v35
	v89 = v41
	v90 = v33
	goto L10
L12:
	;
	goto L13
L13:
	;
	if v35&int32(3) != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v45 = v35
	v47 = v33
	goto L17
L15:
	;
	v59 = v35
	v61 = v33
	goto L16
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v66 = int32(-2139062144)
	if (int32(16843008)-v63|v63)&v66 != v66 {
		v88 = v59
		v89 = v63
		v90 = v61
		goto L10
	} else {
		goto L21
	}
L17:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	*(*uint8)(unsafe.Add(mBase, uint32(v47))) = uint8(v48)
	if v48 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L18:
	;
	v59 = v55
	v61 = v53
	goto L16
L19:
	;
	v52 = int32(1)
	v53 = v47 + v52
	v55 = v45 + v52
	if v55&int32(3) != 0 {
		v45 = v55
		v47 = v53
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v71 = v59
	v72 = v63
	v73 = v61
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v72
	v75 = int32(4)
	v76 = v73 + v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v79 = v71 + v75
	v83 = int32(-2139062144)
	if (v77|(int32(16843008)-v77))&v83 == v83 {
		v71 = v79
		v72 = v77
		v73 = v76
		goto L22
	} else {
		goto L24
	}
L23:
	;
	v88 = v79
	v89 = v77
	v90 = v76
	goto L10
L24:
	;
	goto L23
L25:
	;
	v97 = v88
	v99 = v90
	goto L26
L26:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)) = uint8(v100)
	v102 = int32(1)
	if v100 != 0 {
		v97 = v97 + v102
		v99 = v99 + v102
		goto L26
	} else {
		goto L28
	}
L27:
	;
	goto L9
L28:
	;
	goto L27
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	if v122&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	v115 = F__emscripten_memcpy_bulkmem(m, v11+int32(93), v11+int32(10), int32(83))
	mBase = m.M
	goto L32
L32:
	;
	goto L29
L33:
	;
	v125 = int32(16386)
	goto L35
L34:
	;
	v125 = int32(2)
	goto L35
L35:
	;
	v127 = F_PathNameOpenFile(m, v11+int32(93), v125|l3)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	if int32(0) <= v127 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v135 = l0 + l1<<(uint(int32(2))%32) + int32(40)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v138 = l2 + int32(1)
	if v138 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v188 = int32(0)
	goto L39
L39:
	;
	m.G0 = v11 + int32(176)
	return v188
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v138
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+56))
	v185 = v182 + l2<<(uint(int32(3))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v185)+4)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v185))) = v127
	v188 = v185
	goto L39
L41:
	;
	if v136 <= int32(0) {
		goto L40
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	if v136 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L44:
	;
	v147 = l0 + l1<<(uint(int32(2))%32) + int32(56)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	F_pfree(m, v148)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = int32(0)
	goto L40
L46:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _consts[1184]))
	v162 = F_MemoryContextAlloc(m, v159, v138<<(uint(int32(3))%32))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v138 <= v136 {
		goto L40
	} else {
		goto L50
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(int32(2))%32))+56)) = v162
	goto L40
L50:
	;
	v170 = l0 + l1<<(uint(int32(2))%32) + int32(56)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)))
	v174 = F_repalloc(m, v171, v138<<(uint(int32(3))%32))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v170))) = v174
	goto L40
}
func F_makeBoolean(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v1 = l0
	v4 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+4)) = uint8(v1)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(467)
		return v4
	}
}
func F_makeCompoundFlags(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v59 int32
	_ = v59
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(1056)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v9 == v3 {
		v59 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v7 + int32(1056)
	return v59
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12+l1<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v18 == int32(0) {
		v59 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(0)
	goto L4
L4:
	;
	F_getNextFlagFromString(m, l0, v7+int32(12), v7+int32(16))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v59 = v51 & int32(15)
	goto L1
L6:
	;
	return int32(0)
L7:
	;
	F_setCompoundAffixFlagValue(m, l0, v7+int32(1044), v7+int32(16), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v47 = F_bsearch(m, v7+int32(1044), v43, v44, int32(12), int32(1166))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	if v47 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
	v51 = v49 | v23
	goto L12
L11:
	;
	v51 = v23
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52))))
	if v53 != 0 {
		v23 = v51
		goto L4
	} else {
		goto L13
	}
L13:
	;
	goto L5
}
func F_makeWholeRowVar(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v80 int32
	_ = v80
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
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	v5 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = int32(2249)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	switch v15 {
	case 0:
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v17 = F_get_rel_type_id(m, v16)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v17 != 0 {
				v104 = v5
				v105 = v5
				v106 = v17
				v109 = F_palloc0(m, int32(48))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					v111 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
					*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
					v115 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
					*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
					*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
					*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
					m.G0 = v12 + int32(32)
					return v109
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return int32(0)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
						v29 = F_get_rel_name(m, v28)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v29
							F_errmsg(m, int32(363660), v12)
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(486443), int32(155), int32(226914))
								mBase = m.M
								v39 = m.ExcPending
								if v39 != 0 {
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
	case 1:
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		if v40 != 0 {
			v41 = F_get_rel_type_id(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				if v41 != 0 {
					v104 = v5
					v105 = v5
					v106 = v41
					v109 = F_palloc0(m, int32(48))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						v111 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
						*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
						v115 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
						*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
						*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
						*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
						m.G0 = v12 + int32(32)
						return v109
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(151027844))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
							v51 = F_get_rel_name(m, v50)
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v51
								F_errmsg(m, int32(363660), v12+int32(16))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(486443), int32(181), int32(226914))
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
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
		} else {
			v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if v64 == int32(0) {
				v101 = v14
				v104 = v5
				v105 = v5
				v106 = v101
				v109 = F_palloc0(m, int32(48))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					v111 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
					*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
					v115 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
					*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
					*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
					*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
					m.G0 = v12 + int32(32)
					return v109
				}
			} else {
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v64)+12))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)))
				v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
				v70 = F_exprType(m, v69)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					v73 = F_type_is_rowtype(m, v70)
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						if v73 != 0 {
							v75 = v70
						} else {
							v75 = int32(2249)
						}
						v101 = v75
						v104 = v5
						v105 = v5
						v106 = v101
						v109 = F_palloc0(m, int32(48))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							v111 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
							*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
							v115 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
							*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
							*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
							*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
							m.G0 = v12 + int32(32)
							return v109
						}
					}
				}
			}
		}
	default:
		v104 = v5
		v105 = v5
		v106 = v14
		v109 = F_palloc0(m, int32(48))
		mBase = m.M
		v110 = m.ExcPending
		if v110 != 0 {
			return int32(0)
		} else {
			v111 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
			*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
			*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
			v115 = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
			*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
			*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
			*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
			*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
			*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
			*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
			*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
			m.G0 = v12 + int32(32)
			return v109
		}
	case 3:
		v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+72)))
		if v76 != 0 {
			v104 = v5
			v105 = v5
			v106 = v14
			v109 = F_palloc0(m, int32(48))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return int32(0)
			} else {
				v111 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
				*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
				*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
				v115 = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
				*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
				*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
				*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
				*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
				*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
				*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
				m.G0 = v12 + int32(32)
				return v109
			}
		} else {
			v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
			if v77 == int32(0) {
				v101 = v14
				v104 = v5
				v105 = v5
				v106 = v101
				v109 = F_palloc0(m, int32(48))
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					v111 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
					*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
					v115 = int32(-1)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
					*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
					*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
					*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
					*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
					m.G0 = v12 + int32(32)
					return v109
				}
			} else {
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+4))
				if v80 != int32(1) {
					v104 = v5
					v105 = v5
					v106 = v14
					v109 = F_palloc0(m, int32(48))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						v111 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
						*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
						*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
						v115 = int32(-1)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
						*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
						*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
						*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
						m.G0 = v12 + int32(32)
						return v109
					}
				} else {
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v84)+4))
					v86 = F_exprType(m, v85)
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int32(0)
					} else {
						v89 = F_type_is_rowtype(m, v86)
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							if v89 != 0 {
								v91 = v86
							} else {
								v91 = int32(2249)
							}
							if v89 != 0 {
								v104 = int32(0)
								v105 = v5
								v106 = v91
								v109 = F_palloc0(m, int32(48))
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									v111 = int32(0)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
									*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
									v115 = int32(-1)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
									*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
									*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
									*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
									*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
									*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
									m.G0 = v12 + int32(32)
									return v109
								}
							} else {
								v93 = int32(0)
								if l3 == v93 {
									v104 = v93
									v105 = v5
									v106 = v91
									v109 = F_palloc0(m, int32(48))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return int32(0)
									} else {
										v111 = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
										*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
										*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
										v115 = int32(-1)
										*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
										*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
										*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
										*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
										*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
										*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
										*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
										*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
										m.G0 = v12 + int32(32)
										return v109
									}
								} else {
									v97 = F_exprCollation(m, v85)
									mBase = m.M
									v98 = m.ExcPending
									if v98 != 0 {
										return int32(0)
									} else {
										v104 = v97
										v105 = int32(1)
										v106 = v86
										v109 = F_palloc0(m, int32(48))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return int32(0)
										} else {
											v111 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v109)+32)) = v111
											*(*int32)(unsafe.Add(mBase, uint32(v109)+28)) = l2
											*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v104
											v115 = int32(-1)
											*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v115
											*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v106
											*(*uint16)(unsafe.Add(mBase, uint32(v109)+8)) = uint16(v105)
											*(*int32)(unsafe.Add(mBase, uint32(v109)+4)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v109))) = int32(6)
											*(*int32)(unsafe.Add(mBase, uint32(v109)+44)) = v115
											*(*uint16)(unsafe.Add(mBase, uint32(v109)+40)) = uint16(v105)
											*(*int32)(unsafe.Add(mBase, uint32(v109)+36)) = l1
											*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = v111
											m.G0 = v12 + int32(32)
											return v109
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
func F_make_absolute_path(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L20
	} else {
		goto L73
	}
L2:
	;
	F_emscripten_builtin_free(m, v18)
	mBase = m.M
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L20
	} else {
		goto L69
	}
L3:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != int32(47) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v214 = int32(0)
	goto L5
L5:
	;
	m.G0 = v8 + int32(16)
	return v214
L6:
	;
	F_canonicalize_path_enc(m, v208)
	mBase = m.M
	v214 = v208
	goto L5
L7:
	;
	v13 = int32(1024)
	v15 = F_emscripten_builtin_malloc(m, v13)
	mBase = m.M
	if v15 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	goto L9
L9:
	;
	v196 = F_strlen(m, l0)
	mBase = m.M
	v198 = v196 + int32(1)
	v199 = F_emscripten_builtin_malloc(m, v198)
	mBase = m.M
	if v199 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L10:
	;
	if v18&int32(3) == int32(0) {
		v91 = v18
		goto L30
	} else {
		goto L31
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v23
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L20
	} else {
		goto L25
	}
L12:
	;
	v17 = v13
	v18 = v15
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L20
	} else {
		goto L21
	}
L15:
	;
	v21 = F_getcwd(m, v18, v17)
	mBase = m.M
	if v21 != 0 {
		goto L10
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	F_emscripten_builtin_free(m, v18)
	mBase = m.M
	if v23 != int32(68) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	v28 = v17 << (uint(int32(1)) % 32)
	v29 = F_emscripten_builtin_malloc(m, v28)
	mBase = m.M
	if v29 != 0 {
		v17 = v28
		v18 = v29
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	return int32(0)
L21:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	F_errmsg(m, int32(13796), int32(0))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(489352), int32(829), int32(316285))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	F_errmsg_internal(m, int32(287283), int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(489352), int32(851), int32(316285))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L20
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	if l0&int32(3) == int32(0) {
		v148 = l0
		goto L47
	} else {
		goto L48
	}
L29:
	;
	v124 = v116 - v18
	goto L28
L30:
	;
	v95 = v91
	goto L39
L31:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	if v75 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v124 = int32(0)
	goto L28
L33:
	;
	goto L34
L34:
	;
	v80 = v18
	goto L35
L35:
	;
	v84 = v80 + int32(1)
	if v84&int32(3) == int32(0) {
		v91 = v84
		goto L30
	} else {
		goto L37
	}
L36:
	;
	v116 = v84
	goto L29
L37:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v89 != 0 {
		v80 = v84
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v104 = int32(-2139062144)
	if (int32(16843008)-v101|v101)&v104 == v104 {
		v95 = v95 + int32(4)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v110 = v95
	goto L42
L41:
	;
	goto L40
L42:
	;
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v110))))
	if v114 != 0 {
		v110 = v110 + int32(1)
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v116 = v110
	goto L29
L44:
	;
	goto L43
L45:
	;
	v185 = F_emscripten_builtin_malloc(m, v124+v181+int32(2))
	mBase = m.M
	if v185 == int32(0) {
		goto L2
	} else {
		goto L62
	}
L46:
	;
	v181 = v173 - l0
	goto L45
L47:
	;
	v152 = v148
	goto L56
L48:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v132 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v181 = int32(0)
	goto L45
L50:
	;
	goto L51
L51:
	;
	v137 = l0
	goto L52
L52:
	;
	v141 = v137 + int32(1)
	if v141&int32(3) == int32(0) {
		v148 = v141
		goto L47
	} else {
		goto L54
	}
L53:
	;
	v173 = v141
	goto L46
L54:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	if v146 != 0 {
		v137 = v141
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v161 = int32(-2139062144)
	if (int32(16843008)-v158|v158)&v161 == v161 {
		v152 = v152 + int32(4)
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v167 = v152
	goto L59
L58:
	;
	goto L57
L59:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167))))
	if v171 != 0 {
		v167 = v167 + int32(1)
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v173 = v167
	goto L46
L61:
	;
	goto L60
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v18
	v191 = F_pg_sprintf(m, v185, int32(174219), v8)
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L20
	} else {
		goto L63
	}
L63:
	;
	F_emscripten_builtin_free(m, v18)
	mBase = m.M
	v208 = v185
	goto L6
L64:
	;
	if v204 == int32(0) {
		goto L1
	} else {
		goto L68
	}
L65:
	;
	v204 = int32(0)
	goto L64
L66:
	;
	goto L67
L67:
	;
	v203 = F___memcpy(m, v199, l0, v198)
	mBase = m.M
	v204 = v203
	goto L64
L68:
	;
	v208 = v204
	goto L6
L69:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L20
	} else {
		goto L70
	}
L70:
	;
	F_errmsg(m, int32(13796), int32(0))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L20
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(489352), int32(866), int32(316285))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L20
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L73:
	;
	F_errcode(m, int32(8389))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L20
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(13796), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L20
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(489352), int32(883), int32(316285))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L20
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_make_attrmap(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v4 = F_palloc0(m, int32(8))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = l0
		v11 = F_palloc0(m, l0<<(uint(int32(1))%32))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4))) = v11
			return v4
		}
	}
}
func F_make_canonical_pathkey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	v5 = l4
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+92)))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = l1
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L22
	} else {
		goto L25
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
	if v20 != 0 {
		v12 = v20
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v21 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L5
L7:
	;
	return v93
L8:
	;
	v66 = int32(4470560)
	v67 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+280))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v69
	v72 = F_palloc0(m, int32(20))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L22
	} else {
		goto L23
	}
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v24 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v27 = int32(0)
	if v27 < v24 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v31 = v24
	goto L13
L12:
	;
	v31 = v27
	goto L13
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v34 = v27
	goto L14
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32+v34<<(uint(int32(2))%32))))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v12 != v46 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L8
L16:
	;
	v55 = v34 + int32(1)
	if v55 != v31 {
		v34 = v55
		goto L14
	} else {
		goto L21
	}
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if l2 != v48 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if l3 != v50 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+16)))
	if v52 == v5 {
		v93 = v45
		goto L7
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	goto L15
L22:
	;
	return int32(0)
L23:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v72)+16)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v72)+4)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v72))) = int32(275)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v83 = F_lappend(m, v82, v72)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v83
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v67
	v93 = v72
	goto L7
L25:
	;
	F_errmsg_internal(m, int32(110782), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(485125), int32(66), int32(20781))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L22
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_make_empty_range(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v5)+14)) = uint8(v7)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v2
	*(*int32)(unsafe.Add(mBase, uint32(v5)+3)) = v2
	v21 = F_make_range(m, l0, v5+int32(8), v5, v7, v2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		m.G0 = v5 + int32(16)
		return v21
	}
}
func F_make_oper_cache_key(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
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
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v253 int32
	_ = v253
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	F_DeconstructQualifiedName(m, l2, v11+int32(28), v11+int32(24))
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
	if l1&int32(3) != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)+24))
	goto L15
L4:
	;
	v41 = int32(136)
	goto L6
L5:
	;
	v26 = l1 + int32(136)
	if base.Ui32(v26) <= base.Ui32(l1) {
		goto L3
	} else {
		goto L7
	}
L6:
	;
	v43 = F__emscripten_memset_bulkmem(m, l1, base.I32_extend8_s(int32(0)), v41)
	mBase = m.M
	goto L11
L7:
	;
	v31 = l1 + int32(4)
	if base.Ui32(v31) < base.Ui32(v26) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = v26
	goto L10
L9:
	;
	v33 = v31
	goto L10
L10:
	;
	v41 = (l1^int32(-1)+v33)&int32(-4) + int32(4)
	goto L6
L11:
	;
	goto L3
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l1)+64)) = l3
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v164 != 0 {
		goto L46
	} else {
		goto L47
	}
L13:
	;
	v159 = F_strlen(m, v148)
	mBase = m.M
	goto L12
L15:
	;
	goto L16
L16:
	;
	v53 = int32(63)
	if (l1^v46)&int32(3) != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v152 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v152)
	goto L13
L18:
	;
	v133 = v128
	v134 = v129
	v135 = v130
	goto L40
L19:
	;
	if v123 == int32(0) {
		v148 = v121
		v149 = v122
		goto L17
	} else {
		goto L39
	}
L20:
	;
	v121 = v46
	v122 = l1
	v123 = v53
	goto L19
L21:
	;
	goto L22
L22:
	;
	if v46&int32(3) == int32(0) {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v90 == int32(0) {
		v148 = v87
		v149 = v88
		goto L17
	} else {
		goto L32
	}
L24:
	;
	v87 = v46
	v88 = l1
	v89 = v53
	v90 = int32(1)
	goto L23
L25:
	;
	goto L26
L26:
	;
	v66 = v46
	v67 = l1
	v68 = v53
	goto L27
L27:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v70)
	if v70 == int32(0) {
		v128 = v66
		v129 = v67
		v130 = v68
		goto L18
	} else {
		goto L29
	}
L28:
	;
	v87 = v81
	v88 = v75
	v89 = v77
	v90 = v79
	goto L23
L29:
	;
	v74 = int32(1)
	v75 = v67 + v74
	v77 = v68 - v74
	v78 = int32(0)
	v79 = base.B2i32(v77 != v78)
	v81 = v66 + v74
	if v81&int32(3) == v78 {
		v87 = v81
		v88 = v75
		v89 = v77
		v90 = v79
		goto L23
	} else {
		goto L30
	}
L30:
	;
	if v77 != 0 {
		v66 = v81
		v67 = v75
		v68 = v77
		goto L27
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v93 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87))))
	if v93 == int32(0) {
		v121 = v87
		v122 = v88
		v123 = v89
		goto L19
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(v89) < base.Ui32(int32(4)) {
		v121 = v87
		v122 = v88
		v123 = v89
		goto L19
	} else {
		goto L34
	}
L34:
	;
	v99 = v87
	v100 = v88
	v101 = v89
	goto L35
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	v107 = int32(-2139062144)
	if (int32(16843008)-v104|v104)&v107 != v107 {
		v128 = v99
		v129 = v100
		v130 = v101
		goto L18
	} else {
		goto L37
	}
L36:
	;
	v121 = v115
	v122 = v113
	v123 = v117
	goto L19
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100))) = v104
	v112 = int32(4)
	v113 = v100 + v112
	v115 = v99 + v112
	v117 = v101 - v112
	if base.Ui32(int32(3)) < base.Ui32(v117) {
		v99 = v115
		v100 = v113
		v101 = v117
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v128 = v121
	v129 = v122
	v130 = v123
	goto L18
L40:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v137)
	if v137 == int32(0) {
		v148 = v133
		v149 = v134
		goto L17
	} else {
		goto L42
	}
L41:
	;
	v148 = v144
	v149 = v142
	goto L17
L42:
	;
	v141 = int32(1)
	v142 = v134 + v141
	v144 = v133 + v141
	v146 = v135 - v141
	if v146 != 0 {
		v133 = v144
		v134 = v142
		v135 = v146
		goto L40
	} else {
		goto L43
	}
L43:
	;
	goto L41
L44:
	;
	m.G0 = v11 + int32(32)
	return v253
L45:
	;
	v253 = int32(1)
	goto L44
L46:
	;
	v166 = v11 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v166)+12)) = int32(489)
	*(*int32)(unsafe.Add(mBase, uint32(v166)+4)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v166)+16)) = v166
	v172 = int32(4463464)
	v173 = *(*int32)(unsafe.Add(mBase, _consts[385]))
	*(*int32)(unsafe.Add(mBase, uint32(v166)+8)) = v173
	*(*int32)(unsafe.Add(mBase, _consts[385])) = v11 + int32(12)
	goto L49
L47:
	;
	goto L48
L48:
	;
	v189 = int32(0)
	F_recomputeNamespacePath(m)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L52
	}
L49:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	v181 = F_LookupExplicitNamespace(m, v179, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v181
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v11+int32(4))+8))
	*(*int32)(unsafe.Add(mBase, _consts[385])) = v187
	goto L51
L51:
	;
	goto L45
L52:
	;
	v197 = *(*int32)(unsafe.Add(mBase, _consts[379]))
	if v197 == int32(0) {
		v232 = v189
		goto L53
	} else {
		goto L54
	}
L53:
	;
	if int32(16) < v232 {
		v253 = v189
		goto L44
	} else {
		goto L65
	}
L54:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v200 <= int32(0) {
		v232 = v189
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _consts[251]))
	v205 = v189
	v208 = v189
	goto L56
L56:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v197)+12))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v213+v208<<(uint(int32(2))%32))))
	if v204 != v217 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v232 = v227
	goto L53
L58:
	;
	if v205 < int32(16) {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	v227 = v205
	goto L60
L60:
	;
	v229 = v208 + int32(1)
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v197)+4))
	if v229 < v230 {
		v205 = v227
		v208 = v229
		goto L56
	} else {
		goto L64
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+int32(72)+v205<<(uint(int32(2))%32)))) = v217
	goto L63
L62:
	;
	goto L63
L63:
	;
	v227 = v205 + int32(1)
	goto L60
L64:
	;
	goto L57
L65:
	;
	goto L45
}
func F_make_placeholder_expr(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v6 = F_palloc0(m, int32(24))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = int32(319)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+68))
		v19 = v17 + int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+68)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v10
		*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v19
		return v6
	}
}
func F_make_trigrams(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
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
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var __phi237 int32
	_ = __phi237
	var v238 int32
	_ = v238
	var __phi238 int32
	_ = __phi238
	var v239 int32
	_ = v239
	var __phi239 int32
	_ = __phi239
	var v240 int32
	_ = v240
	var __phi240 int32
	_ = __phi240
	var v241 int32
	_ = v241
	var __phi241 int32
	_ = __phi241
	var v243 int32
	_ = v243
	var __phi243 int32
	_ = __phi243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	if int32(3) <= l2 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L10
	} else {
		goto L60
	}
L2:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = l2 + v18 - int32(2)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if base.Ui32(v21) <= base.Ui32(v22) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	return
L5:
	;
	v43 = v38 + v37*int32(3) + int32(5)
	v45 = *(*int32)(unsafe.Add(mBase, _consts[485]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	goto L13
L6:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = v18
	v38 = v24
	goto L5
L7:
	;
	goto L8
L8:
	;
	if base.Ui32(int32(357913942)) <= base.Ui32(v21) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v32 = F_repalloc(m, v27, v21*int32(3)+int32(5))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v32
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v37 = v36
	v38 = v32
	goto L5
L12:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v371 = base.I32_div_s(v354-v366-int32(5), int32(3))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v371
	goto L4
L13:
	;
	if base.Ui32(v46) <= base.Ui32(int32(41)) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v56 == int32(1) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46*int32(28))+uint32(_consts[1294])))
	v56 = v55
	goto L17
L16:
	;
	v56 = int32(1)
	goto L17
L17:
	;
	goto L14
L18:
	;
	v61 = l1 + l2 - int32(2)
	if base.Ui32(v61) <= base.Ui32(l1) {
		v354 = v43
		goto L12
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v161 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1))))
	if v161 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L21:
	;
	v63 = int32(3)
	v66 = l2 & v63
	if v66 != int32(2) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v71 = l1
	v72 = int32(0)
	v73 = v43
	goto L25
L23:
	;
	v101 = l1
	v103 = v43
	goto L24
L24:
	;
	if base.Ui32(l2-v63) < base.Ui32(int32(3)) {
		v354 = v103
		goto L12
	} else {
		goto L28
	}
L25:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71))))
	*(*uint8)(unsafe.Add(mBase, uint32(v73))) = uint8(v85)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)) = uint8(v87)
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v73)+2)) = uint8(v89)
	v92 = v73 + int32(3)
	v93 = int32(1)
	v94 = v71 + v93
	v96 = v72 + v93
	if v66^v96 != int32(2) {
		v71 = v94
		v72 = v96
		v73 = v92
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v101 = v94
	v103 = v92
	goto L24
L27:
	;
	goto L26
L28:
	;
	v118 = v101
	v120 = v103
	goto L29
L29:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118))))
	*(*uint8)(unsafe.Add(mBase, uint32(v120))) = uint8(v132)
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+1)) = uint8(v134)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+2)) = uint8(v136)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+3)) = uint8(v138)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+4)) = uint8(v140)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+5)) = uint8(v142)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+6)) = uint8(v144)
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+7)) = uint8(v146)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+8)) = uint8(v148)
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+9)) = uint8(v150)
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+10)) = uint8(v152)
	v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v120)+11)) = uint8(v154)
	v157 = v120 + int32(12)
	v159 = v118 + int32(4)
	if v159 != v61 {
		v118 = v159
		v120 = v157
		goto L29
	} else {
		goto L31
	}
L30:
	;
	v354 = v157
	goto L12
L31:
	;
	goto L30
L32:
	;
	v227 = F_pg_mblen_unbounded(m, v217)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L46
	}
L33:
	;
	v203 = F_pg_mblen_unbounded(m, l1)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L10
	} else {
		goto L42
	}
L34:
	;
	v164 = int32(*(*int8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v164 < int32(0) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v174 = v43
	v175 = l1
	goto L36
L36:
	;
	v186 = int32(*(*int8)(unsafe.Add(mBase, uint32(v175)+2)))
	if v186 < int32(0) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v354 = v199
	goto L12
L38:
	;
	v215 = v174
	v216 = v175
	v217 = v175 + int32(2)
	v218 = int32(1)
	v221 = int32(1)
	goto L32
L39:
	;
	goto L40
L40:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175))))
	*(*uint8)(unsafe.Add(mBase, uint32(v174))) = uint8(v192)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v174)+1)) = uint8(v194)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v174)+2)) = uint8(v196)
	v199 = v174 + int32(3)
	v201 = v175 + int32(1)
	if v201 != l1+l2-int32(2) {
		v174 = v199
		v175 = v201
		goto L36
	} else {
		goto L41
	}
L41:
	;
	goto L37
L42:
	;
	if l2 <= v203 {
		v354 = v43
		goto L12
	} else {
		goto L43
	}
L43:
	;
	v206 = l1 + v203
	v207 = F_pg_mblen_unbounded(m, v206)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L10
	} else {
		goto L44
	}
L44:
	;
	v209 = v207 + v206
	if base.Ui32(l1+l2) <= base.Ui32(v209) {
		v354 = v43
		goto L12
	} else {
		goto L45
	}
L45:
	;
	v215 = v43
	v216 = l1
	v217 = v209
	v218 = v203
	v221 = v207
	goto L32
L46:
	;
	v231 = v227 + (v216 + v218 + v221)
	v232 = l1 + l2
	if base.Ui32(v232) < base.Ui32(v231) {
		v354 = v215
		goto L12
	} else {
		goto L47
	}
L47:
	;
	__phi237 = v215
	__phi238 = v216
	__phi239 = v227
	__phi240 = v218
	__phi241 = v231
	__phi243 = v221
	v237 = __phi237
	v238 = __phi238
	v239 = __phi239
	v240 = __phi240
	v241 = __phi241
	v243 = __phi243
	goto L48
L48:
	;
	v249 = int32(255)
	v253 = v241 - v238
	switch v253 {
	case 0:
		v313 = v253
		goto L51
	default:
		goto L52
	case 3:
		goto L53
	}
L49:
	;
	v354 = v344
	goto L12
L50:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v237)+2)) = uint8(v341)
	v344 = v237 + int32(3)
	if v241 == v232 {
		v354 = v344
		goto L12
	} else {
		goto L57
	}
L51:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v237))) = uint16(v313)
	v341 = int32(base.Ui32(v313) >> (uint(int32(16)) % 32))
	goto L50
L52:
	;
	v260 = v238
	v264 = v253
	v267 = v249
	v270 = v249
	v271 = v249
	v272 = v249
	goto L54
L53:
	;
	v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238))))
	*(*uint8)(unsafe.Add(mBase, uint32(v237))) = uint8(v254)
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v237)+1)) = uint8(v256)
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+2)))
	v341 = v258
	goto L50
L54:
	;
	v274 = int32(8)
	v278 = int32(16)
	v283 = int32(24)
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	v294 = *(*int32)(unsafe.Add(mBase, uint32((v286^v267&int32(255))<<(uint(int32(2))%32))+uint32(_consts[1341])))
	v295 = v270<<(uint(v274)%32)&int32(65280) | v271<<(uint(v278)%32)&int32(16711680) | v272<<(uint(v283)%32) ^ v294
	v302 = int32(1)
	v305 = v264 - v302
	if v305 != 0 {
		v260 = v260 + v302
		v264 = v305
		v267 = int32(base.Ui32(v295) >> (uint(v283) % 32))
		v270 = v294
		v271 = int32(base.Ui32(v295) >> (uint(v274) % 32))
		v272 = int32(base.Ui32(v295) >> (uint(v278) % 32))
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v313 = v295 ^ int32(-1)
	goto L51
L56:
	;
	goto L55
L57:
	;
	v347 = F_pg_mblen_unbounded(m, v241)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	v349 = v241 + v347
	if base.Ui32(v349) <= base.Ui32(v232) {
		__phi237 = v344
		__phi238 = v238 + v240
		__phi239 = v347
		__phi240 = v243
		__phi241 = v349
		__phi243 = v239
		v237 = __phi237
		v238 = __phi238
		v239 = __phi239
		v240 = __phi240
		v241 = __phi241
		v243 = __phi243
		goto L48
	} else {
		goto L59
	}
L59:
	;
	goto L49
L60:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L10
	} else {
		goto L61
	}
L61:
	;
	F_errmsg(m, int32(13796), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L10
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(487330), int32(134), int32(24080))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L10
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_makesearch(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v162 int32
	_ = v162
	var v171 int32
	_ = v171
	var v186 int32
	_ = v186
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v344 int64
	_ = v344
	var v350 int32
	_ = v350
	var v361 int32
	_ = v361
	v3 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v10 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+56)))
	v18 = v10
	goto L4
L3:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v186 == int32(0) {
		goto L1
	} else {
		goto L67
	}
L4:
	;
	v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	if v13 != v24 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	F_rainbow(m, l1, v29, int32(-1), v9, v9)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1+int32(58)))))
	if v24 != v26 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v28 != 0 {
		v18 = v28
		goto L4
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	goto L3
L12:
	;
	return
L13:
	;
	v33 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+56)))
	v35 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v35 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v38 <= v39 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L16
L18:
	;
	v102 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+58)))
	v104 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v104 != 0 {
		goto L40
	} else {
		goto L41
	}
L19:
	;
	F_createarc(m, l1, int32(112), v33, v9, v9)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L12
	} else {
		goto L39
	}
L20:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v41 == int32(0) {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v62 == int32(0) {
		goto L19
	} else {
		goto L31
	}
L23:
	;
	v48 = v41
	goto L24
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	if v54 != v9 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L19
L26:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	if v61 != 0 {
		v48 = v61
		goto L24
	} else {
		goto L30
	}
L27:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)))
	if v56 != v33&int32(65535) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	if v58 == int32(112) {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	goto L26
L30:
	;
	goto L25
L31:
	;
	v69 = v62
	goto L32
L32:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v75 != v9 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L19
L34:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v69)+24))
	if v82 != 0 {
		v69 = v82
		goto L32
	} else {
		goto L38
	}
L35:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+4)))
	if v77 != v33&int32(65535) {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v79 == int32(112) {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	goto L33
L39:
	;
	goto L18
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L12
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v9)+8))
	if v107 <= v108 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	goto L42
L44:
	;
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+64)))
	if v171&int32(2) == int32(0) {
		goto L3
	} else {
		goto L66
	}
L45:
	;
	F_createarc(m, l1, int32(112), v102, v9, v9)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L12
	} else {
		goto L65
	}
L46:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
	if v110 == int32(0) {
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	if v131 == int32(0) {
		goto L45
	} else {
		goto L57
	}
L49:
	;
	v117 = v110
	goto L50
L50:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v117)+12))
	if v123 != v9 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	goto L45
L52:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v117)+16))
	if v130 != 0 {
		v117 = v130
		goto L50
	} else {
		goto L56
	}
L53:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v117)+4)))
	if v125 != v102&int32(65535) {
		goto L52
	} else {
		goto L54
	}
L54:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v127 == int32(112) {
		goto L44
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	goto L51
L57:
	;
	v138 = v131
	goto L58
L58:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	if v144 != v9 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L45
L60:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v138)+24))
	if v151 != 0 {
		v138 = v151
		goto L58
	} else {
		goto L64
	}
L61:
	;
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+4)))
	if v146 != v102&int32(65535) {
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	if v148 == int32(112) {
		goto L44
	} else {
		goto L63
	}
L63:
	;
	goto L60
L64:
	;
	goto L59
L65:
	;
	goto L44
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = int32(256)
	goto L3
L67:
	;
	v194 = v186
	v195 = v3
	goto L68
L68:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v197)+16))
	if v198 == int32(0) {
		v221 = v195
		goto L70
	} else {
		goto L71
	}
L69:
	;
	if v221 == int32(0) {
		goto L1
	} else {
		goto L83
	}
L70:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	if v223 != 0 {
		v194 = v223
		v195 = v221
		goto L68
	} else {
		goto L82
	}
L71:
	;
	v203 = v198
	goto L72
L72:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v203)+8))
	if v9 == v209 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v197)+24))
	if v212 != 0 {
		v221 = v195
		goto L70
	} else {
		goto L78
	}
L74:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v203)+24))
	if v211 != 0 {
		v203 = v211
		goto L72
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	goto L73
L77:
	;
	v221 = v195
	goto L70
L78:
	;
	if v195 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v213 = v195
	goto L81
L80:
	;
	v213 = v197
	goto L81
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v197)+24)) = v213
	v221 = v197
	goto L70
L82:
	;
	goto L69
L83:
	;
	v232 = v221
	goto L84
L84:
	;
	v234 = F_newstate(m, l1)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L12
	} else {
		goto L86
	}
L85:
	;
	goto L1
L86:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v236 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v234)+12))
	if v237 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v232)+16))
	if v264 != 0 {
		goto L96
	} else {
		goto L97
	}
L89:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v232)+20))
	if v238 == int32(0) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	v243 = v238
	goto L91
L91:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
	v250 = int32(*(*int16)(unsafe.Add(mBase, uint32(v243)+4)))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v243)+12))
	F_createarc(m, l1, v249, v250, v234, v251)
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L12
	} else {
		goto L93
	}
L92:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v255 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v243)+16))
	if v254 != 0 {
		v243 = v254
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	goto L88
L96:
	;
	v267 = v264
	goto L99
L97:
	;
	goto L98
L98:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v232)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+24)) = int32(0)
	if v361 == v232 {
		goto L1
	} else {
		goto L133
	}
L99:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v267)+24))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	if v9 != v274 {
		goto L101
	} else {
		goto L102
	}
L100:
	;
	goto L98
L101:
	;
	F_cparc(m, l1, v267, v274, v234)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L12
	} else {
		goto L104
	}
L102:
	;
	goto L103
L103:
	;
	if v273 != 0 {
		v267 = v273
		goto L99
	} else {
		goto L132
	}
L104:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v267)+8))
	v285 = int32(*(*int16)(unsafe.Add(mBase, uint32(v267)+4)))
	if v285 < int32(0) {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	goto L103
L106:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v267)+16))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v267)+20))
	if v319 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L107:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	v290 = v288 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v290) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	if int32(1)<<(uint(v290)%32)&int32(163841) == int32(0) {
		goto L106
	} else {
		goto L109
	}
L109:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v299 != 0 {
		goto L106
	} else {
		goto L110
	}
L110:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v267)+36))
	if v300 == int32(0) {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	if v312 != 0 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v303)+20))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v267)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v304+v285*int32(24))+12)) = v308
	v312 = v308
	goto L111
L113:
	;
	goto L114
L114:
	;
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v267)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v300)+32)) = v310
	v312 = v310
	goto L111
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v312)+36)) = v300
	goto L117
L116:
	;
	goto L117
L117:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v267)+32)) = int64(0)
	goto L106
L118:
	;
	if v318 != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v284)+20)) = v318
	goto L118
L120:
	;
	goto L121
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+16)) = v318
	goto L118
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v318)+20)) = v319
	goto L124
L123:
	;
	goto L124
L124:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v284)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v284)+12)) = v325 - int32(1)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v267)+24))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v267)+28))
	if v330 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L125:
	;
	v336 = v267 + int32(8)
	if v329 != 0 {
		goto L129
	} else {
		goto L130
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v283)+16)) = v329
	goto L125
L127:
	;
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v330)+24)) = v329
	goto L125
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v329)+28)) = v330
	goto L131
L130:
	;
	goto L131
L131:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v283)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v283)+8)) = v338 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = int32(0)
	v344 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v336)+16)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v336)+8)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v336))) = v344
	v350 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v267)+16)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v267
	goto L105
L132:
	;
	goto L100
L133:
	;
	if v361 != 0 {
		v232 = v361
		goto L84
	} else {
		goto L134
	}
L134:
	;
	goto L85
}
func F_markreachable(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v8 = m.T0[v7].(func(*base.Module) int32)(m)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(101)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v18 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v16 = v14
	goto L8
L7:
	;
	v16 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v16
	return
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l2
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v20 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v24 = v20
	goto L12
L12:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	F_markreachable(m, l0, v26, l2)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)+16))
	if v29 != 0 {
		v24 = v29
		goto L12
	} else {
		goto L15
	}
L15:
	;
	goto L13
}
func F_maybe_reread_subscription(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
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
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
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
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	v7 = m.G0
	v9 = v7 - int32(80)
	m.G0 = v9
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[844])))
	if v12 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L10
	} else {
		goto L137
	}
L2:
	;
	if v336 == int32(2) {
		goto L132
	} else {
		goto L133
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _consts[25]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
	v19 = base.B2i32(v17 == int32(2))
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v9 + int32(80)
	return
L6:
	;
	if v19 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v24 = int32(4470560)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[28]))
	v28 = *(*int32)(unsafe.Add(mBase, _consts[845]))
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v28
	v31 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	v34 = F_GetSubscription(m, v32, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L10
	} else {
		goto L12
	}
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	if v34 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v40 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L10
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+25)))
	if v65 != 0 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	if v40 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v44
	F_errmsg(m, int32(433074), v9)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L10
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if v56 == int32(2) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	F_errfinish(m, int32(486965), int32(3991), int32(243246))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
	F_ApplyLauncherForgetWorkerStartTime(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L10
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L10
	} else {
		goto L26
	}
L25:
	;
	goto L24
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v34)+36))
	v105 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+36))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
	if v110 == int32(0) {
		v129 = v109
		v130 = v110
		goto L44
	} else {
		goto L45
	}
L28:
	;
	v68 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L29
	}
L29:
	;
	if v68 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+64)) = v72
	F_errmsg(m, int32(447640), v9-int32(-64))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L10
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+16)))
	if base.B2i32(v87 == int32(1))&base.B2i32(v86 == int32(3)) != 0 {
		goto L27
	} else {
		goto L35
	}
L33:
	;
	F_errfinish(m, int32(486965), int32(4005), int32(243246))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	if v86 == int32(2) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v85)+32))
	F_ApplyLauncherForgetWorkerStartTime(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L10
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L10
	} else {
		goto L40
	}
L39:
	;
	goto L38
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+24)))
	if v290 != 0 {
		v344 = v289
		goto L102
	} else {
		goto L103
	}
L42:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+16)))
	if v231 != int32(1) {
		goto L87
	} else {
		goto L88
	}
L43:
	;
	if v130-v129 != 0 {
		goto L42
	} else {
		goto L51
	}
L44:
	;
	goto L43
L45:
	;
	if v109 != v110 {
		v129 = v109
		v130 = v110
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v114 = v103
	v115 = v106
	goto L47
L47:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
	v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	if v119 == int32(0) {
		v129 = v118
		v130 = v119
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v129 = v118
	v130 = v119
	goto L44
L49:
	;
	v122 = int32(1)
	if v118 == v119 {
		v114 = v114 + v122
		v115 = v115 + v122
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v137 == int32(0) {
		v156 = v136
		v157 = v137
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v157-v156 != 0 {
		goto L42
	} else {
		goto L60
	}
L53:
	;
	goto L52
L54:
	;
	if v136 != v137 {
		v156 = v136
		v157 = v137
		goto L53
	} else {
		goto L55
	}
L55:
	;
	v141 = v132
	v142 = v133
	goto L56
L56:
	;
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142)+1)))
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+1)))
	if v146 == int32(0) {
		v156 = v145
		v157 = v146
		goto L53
	} else {
		goto L58
	}
L57:
	;
	v156 = v145
	v157 = v146
	goto L53
L58:
	;
	v149 = int32(1)
	if v145 == v146 {
		v141 = v141 + v149
		v142 = v142 + v149
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v34)+40))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v105)+40))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v160))))
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v164 == int32(0) {
		v183 = v163
		v184 = v164
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v184-v183 != 0 {
		goto L42
	} else {
		goto L69
	}
L62:
	;
	goto L61
L63:
	;
	if v163 != v164 {
		v183 = v163
		v184 = v164
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v168 = v159
	v169 = v160
	goto L65
L65:
	;
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v169)+1)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v168)+1)))
	if v173 == int32(0) {
		v183 = v172
		v184 = v173
		goto L62
	} else {
		goto L67
	}
L66:
	;
	v183 = v172
	v184 = v173
	goto L62
L67:
	;
	v176 = int32(1)
	if v172 == v173 {
		v168 = v168 + v176
		v169 = v169 + v176
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+26)))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+26)))
	if v186 != v187 {
		goto L42
	} else {
		goto L70
	}
L70:
	;
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+27)))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+27)))
	if v189 != v190 {
		goto L42
	} else {
		goto L71
	}
L71:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34)+30)))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+30)))
	if v192 != v193 {
		goto L42
	} else {
		goto L72
	}
L72:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v34)+52))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v105)+52))
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v196))))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v200 == int32(0) {
		v219 = v199
		v220 = v200
		goto L74
	} else {
		goto L75
	}
L73:
	;
	if v220-v219 != 0 {
		goto L42
	} else {
		goto L81
	}
L74:
	;
	goto L73
L75:
	;
	if v199 != v200 {
		v219 = v199
		v220 = v200
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v204 = v195
	v205 = v196
	goto L77
L77:
	;
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+1)))
	if v209 == int32(0) {
		v219 = v208
		v220 = v209
		goto L74
	} else {
		goto L79
	}
L78:
	;
	v219 = v208
	v220 = v209
	goto L74
L79:
	;
	v212 = int32(1)
	if v208 == v209 {
		v204 = v204 + v212
		v205 = v205 + v212
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v34)+20))
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v105)+20))
	if v222 != v223 {
		goto L42
	} else {
		goto L82
	}
L82:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v34)+48))
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v105)+48))
	v227 = F_equal(m, v225, v226)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L83
	}
L83:
	;
	if v227 != 0 {
		goto L41
	} else {
		goto L84
	}
L84:
	;
	goto L42
L85:
	;
	v270 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+16)))
	if base.B2i32(v272 == int32(1))&base.B2i32(v271 == int32(3)) != 0 {
		goto L41
	} else {
		goto L96
	}
L86:
	;
	v256 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v257
	F_errmsg(m, v253, v9+int32(48))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L10
	} else {
		goto L94
	}
L87:
	;
	v247 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L10
	} else {
		goto L92
	}
L88:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v234 != int32(3) {
		goto L87
	} else {
		goto L89
	}
L89:
	;
	v239 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L10
	} else {
		goto L90
	}
L90:
	;
	if v239 == int32(0) {
		goto L85
	} else {
		goto L91
	}
L91:
	;
	v253 = int32(396394)
	v254 = int32(4036)
	goto L86
L92:
	;
	if v247 == int32(0) {
		goto L85
	} else {
		goto L93
	}
L93:
	;
	v253 = int32(396302)
	v254 = int32(4040)
	goto L86
L94:
	;
	F_errfinish(m, int32(486965), v254, int32(243246))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L10
	} else {
		goto L95
	}
L95:
	;
	goto L85
L96:
	;
	if v271 == int32(2) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v270)+32))
	F_ApplyLauncherForgetWorkerStartTime(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L10
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L10
	} else {
		goto L101
	}
L100:
	;
	goto L99
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v344)+4))
	if v346 != v347 {
		goto L1
	} else {
		goto L118
	}
L103:
	;
	v291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v289)+24)))
	if v291 != int32(1) {
		v344 = v289
		goto L102
	} else {
		goto L104
	}
L104:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+16)))
	if v296 != int32(1) {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	v335 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v337 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+16)))
	if v337 != int32(1) {
		goto L2
	} else {
		goto L116
	}
L106:
	;
	v321 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v322
	F_errmsg(m, v318, v9+int32(32))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L10
	} else {
		goto L114
	}
L107:
	;
	v312 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L10
	} else {
		goto L112
	}
L108:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v295)))
	if v299 != int32(3) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v304 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	if v304 == int32(0) {
		goto L105
	} else {
		goto L111
	}
L111:
	;
	v318 = int32(448598)
	v319 = int32(4054)
	goto L106
L112:
	;
	if v312 == int32(0) {
		goto L105
	} else {
		goto L113
	}
L113:
	;
	v318 = int32(448464)
	v319 = int32(4058)
	goto L106
L114:
	;
	F_errfinish(m, int32(486965), v319, int32(243246))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L10
	} else {
		goto L115
	}
L115:
	;
	goto L105
L116:
	;
	if v336 != int32(3) {
		goto L2
	} else {
		goto L117
	}
L117:
	;
	v343 = *(*int32)(unsafe.Add(mBase, _consts[822]))
	v344 = v343
	goto L102
L118:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v344)+16))
	F_pfree(m, v349)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L10
	} else {
		goto L119
	}
L119:
	;
	v352 = *(*int32)(unsafe.Add(mBase, uint32(v344)+36))
	F_pfree(m, v352)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L10
	} else {
		goto L120
	}
L120:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v344)+40))
	if v355 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	F_pfree(m, v355)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L10
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v344)+48))
	F_list_free_deep(m, v358)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L10
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	F_pfree(m, v344)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L10
	} else {
		goto L126
	}
L126:
	;
	*(*int32)(unsafe.Add(mBase, _consts[28])) = v25
	*(*int32)(unsafe.Add(mBase, _consts[822])) = v34
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v34)+44))
	F_SetConfigOption(m, int32(98975), v368, int32(4), int32(10))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L10
	} else {
		goto L127
	}
L127:
	;
	if v19 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L10
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v378 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[844])) = uint8(v378)
	goto L5
L131:
	;
	goto L130
L132:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v335)+32))
	F_ApplyLauncherForgetWorkerStartTime(m, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L10
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L10
	} else {
		goto L136
	}
L135:
	;
	goto L134
L136:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L137:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v402
	F_errmsg_internal(m, int32(20164), v9+int32(16))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L10
	} else {
		goto L138
	}
L138:
	;
	F_errfinish(m, int32(486965), int32(4067), int32(243246))
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L10
	} else {
		goto L139
	}
L139:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_md5_text(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = int32(0)
		v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
		if v16 == int32(1) {
			v19 = int32(4)
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
			if v21&int32(254) == int32(2) {
				v30 = v19
			} else {
				v30 = base.B2i32(v21 == int32(18)) << (uint(v19) % 32)
			}
			if v21 == int32(1) {
				v33 = v19
			} else {
				v33 = v30
			}
			v46 = v33
		} else {
			v34 = int32(1)
			if v16&v34 != 0 {
				v46 = int32(base.Ui32(v16)>>(uint(v34)%32)) - v34
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = int32(1)
		if v16&v47 != 0 {
			v51 = v47
		} else {
			v51 = int32(4)
		}
		v57 = F_pg_md5_hash(m, v10+v51, v46, v5+int32(-48), v5+int32(-52))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int32(0)
		} else {
			if v57 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(2600))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(541337)
						v70 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v70
						F_errmsg(m, int32(199700), v7)
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(486382), int32(49), int32(62552))
							mBase = m.M
							v79 = m.ExcPending
							if v79 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v82 = F_cstring_to_text(m, v5+int32(-48))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 - int32(-64)
					return v82
				}
			}
		}
	}
}
func F_md_readv_complete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v17 int32
	_ = v17
	var v27 int64
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v39 int64
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = l1 + int32(104)
	v13 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v13
	v17 = base.I32_wrap_i64(int64(base.Ui64(v13) >> (uint(int64(32)) % 64)))
	if v13 < int64(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(257) - v17<<(uint(int32(9))%32)
		v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
		*(*int64)(unsafe.Add(mBase, uint32(v9))) = v27
		F_pgaio_result_report(m, v9, v12, int32(16))
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return
		} else {
			m.G0 = v9 + int32(16)
			return
		}
	} else {
		v33 = int32(base.Ui32(v17) >> (uint(int32(13)) % 32))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33
		if base.Ui64(v13) <= base.Ui64(int64(35184372088831)) {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(257)
			v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v39
			F_pgaio_result_report(m, v9+int32(8), v12, int32(16))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		} else {
			v46 = base.I32_wrap_i64(v13)
			if v46&int32(448) == int32(256) {
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
				if base.Ui32(v51) <= base.Ui32(v33) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v46&int32(-512) | int32(129)
				}
			}
			m.G0 = v9 + int32(16)
			return
		}
	}
}
func F_mda_get_offset_values(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v109 int32
	_ = v109
	v5 = int32(0)
	v11 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1+l0<<(uint(v11)%32)-int32(4)))) = v5
	v19 = l0 - v11
	if v5 <= v19 {
		v26 = v19
		v28 = v5
		for {
			v33 = v26 << (uint(int32(2)) % 32)
			v34 = l1 + v33
			v36 = *(*int32)(unsafe.Add(mBase, uint32(l2+v33)))
			v37 = int32(1)
			v38 = v36 - v37
			*(*int32)(unsafe.Add(mBase, uint32(v34))) = v38
			v41 = v26 + v37
			if l0 <= v41 {
			} else {
				if v28&int32(1) == int32(0) {
					v47 = int32(2)
					v48 = v41 << (uint(v47) % 32)
					v50 = *(*int32)(unsafe.Add(mBase, uint32(l3+v48)))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l2+v48)))
					v56 = v38 - (v50-int32(1))*v54
					*(*int32)(unsafe.Add(mBase, uint32(v34))) = v56
					v60 = v26 + v47
					v61 = v56
				} else {
					v60 = v41
					v61 = v38
				}
				if v28 == int32(0) {
				} else {
					v68 = v60
					v69 = v61
					for {
						v74 = int32(2)
						v75 = v68 << (uint(v74) % 32)
						v77 = *(*int32)(unsafe.Add(mBase, uint32(l3+v75)))
						v78 = int32(1)
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l2+v75)))
						v83 = v69 - (v77-v78)*v81
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = v83
						v86 = v75 + int32(4)
						v88 = *(*int32)(unsafe.Add(mBase, uint32(l3+v86)))
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l2+v86)))
						v94 = v83 - (v88-v78)*v92
						*(*int32)(unsafe.Add(mBase, uint32(v34))) = v94
						v97 = v68 + v74
						if v97 != l0 {
							v68 = v97
							v69 = v94
							continue
						} else {
							break
						}
						break
					}
				}
			}
			v109 = int32(1)
			if int32(0) < v26 {
				v26 = v26 - v109
				v28 = v28 + v109
				continue
			} else {
				break
			}
			break
		}
	} else {
	}
	return
}
func F_mdcbuf_free(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+116))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+20))
	m.T0[v5].(func(*base.Module, int32))(m, v4)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v9 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+116)) = v9
		v13 = F___memset(m, l0, v9, int32(8240))
		mBase = m.M
		F_pfree(m, l0)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			return
		}
	}
}
func F_mdcbuf_init(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_palloc0(m, int32(8240))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(8192)
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = v5
		*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = int32(1)
		return int32(0)
	}
}
func F_mdcbuf_read(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
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
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v338 int64
	_ = v338
	var v339 int64
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v551 int32
	_ = v551
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v580 int32
	_ = v580
	var v587 int32
	_ = v587
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v611 int32
	_ = v611
	var v614 int32
	_ = v614
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v657 int32
	_ = v657
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v694 int32
	_ = v694
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(48)
	return v694
L2:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v685 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v685
	if l2 < v684 {
		goto L200
	} else {
		goto L201
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if l2 <= v16 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v16 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = l0 + int32(46)
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v181 = F_pullf_read(m, l1, v173-(v16+v174)+int32(22), v13+int32(12))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L54
	} else {
		goto L55
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v22 = l0 + int32(46)
	if v20 == v22 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	if v22 == v20 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L5
L9:
	;
	goto L8
L10:
	;
	v27 = v22 + v16
	if base.Ui32(v20-v27) <= base.Ui32(int32(0)-v16<<(uint(int32(1))%32)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = F___memcpy(m, v22, v20, v16)
	mBase = m.M
	goto L8
L12:
	;
	goto L13
L13:
	;
	v37 = (v22 ^ v20) & int32(3)
	if base.Ui32(v22) < base.Ui32(v20) {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	if v139 == int32(0) {
		goto L9
	} else {
		goto L50
	}
L15:
	;
	if base.Ui32(v117) <= base.Ui32(int32(3)) {
		v138 = v116
		v139 = v117
		v140 = v118
		goto L14
	} else {
		goto L46
	}
L16:
	;
	if v37 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	if v37 != 0 {
		v99 = v16
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v138 = v20
	v139 = v16
	v140 = v22
	goto L14
L20:
	;
	goto L21
L21:
	;
	if v22&int32(3) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v116 = v20
	v117 = v16
	v118 = v22
	goto L15
L23:
	;
	goto L24
L24:
	;
	v44 = v20
	v45 = v16
	v46 = v22
	goto L25
L25:
	;
	if v45 == int32(0) {
		goto L9
	} else {
		goto L27
	}
L26:
	;
	v116 = v53
	v117 = v55
	v118 = v57
	goto L15
L27:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	*(*uint8)(unsafe.Add(mBase, uint32(v46))) = uint8(v50)
	v52 = int32(1)
	v53 = v44 + v52
	v55 = v45 - v52
	v57 = v46 + v52
	if v57&int32(3) != 0 {
		v44 = v53
		v45 = v55
		v46 = v57
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	if v99 == int32(0) {
		goto L9
	} else {
		goto L42
	}
L30:
	;
	if v27&int32(3) != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v64 = v16
	goto L34
L32:
	;
	v79 = v16
	goto L33
L33:
	;
	if base.Ui32(v79) <= base.Ui32(int32(3)) {
		v99 = v79
		goto L29
	} else {
		goto L38
	}
L34:
	;
	if v64 == int32(0) {
		goto L9
	} else {
		goto L36
	}
L35:
	;
	v79 = v70
	goto L33
L36:
	;
	v70 = v64 - int32(1)
	v71 = v22 + v70
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v70))))
	*(*uint8)(unsafe.Add(mBase, uint32(v71))) = uint8(v73)
	if v71&int32(3) != 0 {
		v64 = v70
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	v86 = v79
	goto L39
L39:
	;
	v90 = v86 - int32(4)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v20+v90)))
	*(*int32)(unsafe.Add(mBase, uint32(v22+v90))) = v93
	if base.Ui32(int32(3)) < base.Ui32(v90) {
		v86 = v90
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v99 = v90
	goto L29
L41:
	;
	goto L40
L42:
	;
	v106 = v99
	goto L43
L43:
	;
	v110 = v106 - int32(1)
	v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v110))))
	*(*uint8)(unsafe.Add(mBase, uint32(v22+v110))) = uint8(v113)
	if v110 != 0 {
		v106 = v110
		goto L43
	} else {
		goto L45
	}
L44:
	;
	goto L9
L45:
	;
	goto L44
L46:
	;
	v123 = v116
	v124 = v117
	v125 = v118
	goto L47
L47:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v123)))
	*(*int32)(unsafe.Add(mBase, uint32(v125))) = v127
	v129 = int32(4)
	v130 = v123 + v129
	v132 = v125 + v129
	v134 = v124 - v129
	if base.Ui32(int32(3)) < base.Ui32(v134) {
		v123 = v130
		v124 = v134
		v125 = v132
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v138 = v130
	v139 = v134
	v140 = v132
	goto L14
L49:
	;
	goto L48
L50:
	;
	v145 = v138
	v146 = v139
	v147 = v140
	goto L51
L51:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v149)
	v151 = int32(1)
	v156 = v146 - v151
	if v156 != 0 {
		v145 = v145 + v151
		v146 = v156
		v147 = v147 + v151
		goto L51
	} else {
		goto L53
	}
L52:
	;
	goto L9
L53:
	;
	goto L52
L54:
	;
	return int32(0)
L55:
	;
	if v181 < int32(0) {
		v694 = v181
		goto L1
	} else {
		goto L56
	}
L56:
	;
	if v181 == int32(0) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(1)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v194 != int32(211) {
		v296 = int32(225353)
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L59
L59:
	;
	if base.Ui32(int32(22)) <= base.Ui32(v181) {
		goto L87
	} else {
		goto L88
	}
L60:
	;
	F_px_debug(m, v296, int32(0))
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L54
	} else {
		goto L85
	}
L61:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+25)))
	if v199 != int32(20) {
		v296 = int32(225353)
		goto L60
	} else {
		goto L62
	}
L62:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v202)+116))
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v203)+12))
	m.T0[v207].(func(*base.Module, int32, int32, int32))(m, v203, l0+int32(24), int32(2))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L54
	} else {
		goto L63
	}
L63:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v210)+116))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v211)+16))
	m.T0[v214].(func(*base.Module, int32, int32))(m, v211, v13+int32(16))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L54
	} else {
		goto L64
	}
L64:
	;
	v218 = v13 + int32(16)
	v220 = l0 + int32(26)
	v221 = int32(20)
	goto L68
L65:
	;
	v288 = F___memset(m, v13+int32(16), int32(0), int32(20))
	mBase = m.M
	goto L83
L66:
	;
	v283 = int32(0)
	goto L65
L67:
	;
	v257 = v252
	v258 = v253
	v259 = v254
	goto L77
L68:
	;
	if (v218|v220)&int32(3) != 0 {
		v252 = v218
		v253 = v220
		v254 = v221
		goto L67
	} else {
		goto L71
	}
L70:
	;
	if v242 == int32(0) {
		goto L66
	} else {
		goto L76
	}
L71:
	;
	v229 = v218
	v230 = v220
	v231 = v221
	goto L72
L72:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v229)))
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v234 != v235 {
		v252 = v229
		v253 = v230
		v254 = v231
		goto L67
	} else {
		goto L74
	}
L73:
	;
	goto L70
L74:
	;
	v237 = int32(4)
	v238 = v230 + v237
	v240 = v229 + v237
	v242 = v231 - v237
	if base.Ui32(int32(3)) < base.Ui32(v242) {
		v229 = v240
		v230 = v238
		v231 = v242
		goto L72
	} else {
		goto L75
	}
L75:
	;
	goto L73
L76:
	;
	v252 = v240
	v253 = v238
	v254 = v242
	goto L67
L77:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257))))
	v263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v258))))
	if v262 == v263 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v283 = v262 - v263
	goto L65
L79:
	;
	v265 = int32(1)
	v270 = v259 - v265
	if v270 != 0 {
		v257 = v257 + v265
		v258 = v258 + v265
		v259 = v270
		goto L77
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	goto L78
L82:
	;
	goto L66
L83:
	;
	if v283 == int32(0) {
		goto L2
	} else {
		goto L84
	}
L84:
	;
	v296 = int32(319915)
	goto L60
L85:
	;
	v694 = int32(-100)
	goto L1
L86:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v675 + v670
	goto L2
L87:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v307 = l0 + int32(24)
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v308 != 0 {
		goto L91
	} else {
		goto L92
	}
L88:
	;
	goto L89
L89:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v350 = v349 + v181
	if int32(23) <= v350 {
		goto L100
	} else {
		goto L101
	}
L90:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v311)+116))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	m.T0[v313].(func(*base.Module, int32, int32, int32))(m, v312, v307, v308)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L54
	} else {
		goto L94
	}
L91:
	;
	v309 = F__emscripten_memcpy_bulkmem(m, v303+v304, v307, v308)
	mBase = m.M
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L90
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v319 = v308 + v318
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v325 = v181 - int32(22)
	if v325 != 0 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+116))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)+12))
	m.T0[v330].(func(*base.Module, int32, int32, int32))(m, v329, v323, v325)
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L54
	} else {
		goto L99
	}
L96:
	;
	v326 = F__emscripten_memcpy_bulkmem(m, v321+v319, v323, v325)
	mBase = m.M
	goto L98
L97:
	;
	goto L98
L98:
	;
	goto L95
L99:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v333 + v325
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v337 = v325 + v336
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v337)))
	v339 = *(*int64)(unsafe.Add(mBase, uint32(v337)+8))
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v337)+16))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v342 = v307 + v341
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v337)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v342)+20)) = uint16(v343)
	*(*int32)(unsafe.Add(mBase, uint32(v342)+16)) = v340
	*(*int64)(unsafe.Add(mBase, uint32(v342)+8)) = v339
	*(*int64)(unsafe.Add(mBase, uint32(v342))) = v338
	v670 = int32(22)
	goto L86
L100:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v357 = l0 + int32(24)
	v359 = v350 - int32(22)
	if v359 != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	v518 = v349
	goto L102
L102:
	;
	v523 = l0 + v518 + int32(24)
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v523 == v524 {
		goto L155
	} else {
		goto L156
	}
L103:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+116))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v363)+12))
	m.T0[v364].(func(*base.Module, int32, int32, int32))(m, v363, v357, v359)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L54
	} else {
		goto L107
	}
L104:
	;
	v360 = F__emscripten_memcpy_bulkmem(m, v353+v354, v357, v359)
	mBase = m.M
	goto L106
L105:
	;
	goto L106
L106:
	;
	goto L103
L107:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v367 + v359
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v371 = v370 - v359
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v371
	v373 = v359 + v357
	if v357 == v373 {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v518 = v371
	goto L102
L109:
	;
	goto L108
L110:
	;
	v377 = v357 + v371
	if base.Ui32(v373-v377) <= base.Ui32(int32(0)-v371<<(uint(int32(1))%32)) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v384 = F___memcpy(m, v357, v373, v371)
	mBase = m.M
	goto L108
L112:
	;
	goto L113
L113:
	;
	v387 = (v357 ^ v373) & int32(3)
	if base.Ui32(v357) < base.Ui32(v373) {
		goto L116
	} else {
		goto L117
	}
L114:
	;
	if v489 == int32(0) {
		goto L109
	} else {
		goto L150
	}
L115:
	;
	if base.Ui32(v467) <= base.Ui32(int32(3)) {
		v488 = v466
		v489 = v467
		v490 = v468
		goto L114
	} else {
		goto L146
	}
L116:
	;
	if v387 != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	if v387 != 0 {
		v449 = v371
		goto L129
	} else {
		goto L130
	}
L119:
	;
	v488 = v373
	v489 = v371
	v490 = v357
	goto L114
L120:
	;
	goto L121
L121:
	;
	if v357&int32(3) == int32(0) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v466 = v373
	v467 = v371
	v468 = v357
	goto L115
L123:
	;
	goto L124
L124:
	;
	v394 = v373
	v395 = v371
	v396 = v357
	goto L125
L125:
	;
	if v395 == int32(0) {
		goto L109
	} else {
		goto L127
	}
L126:
	;
	v466 = v403
	v467 = v405
	v468 = v407
	goto L115
L127:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394))))
	*(*uint8)(unsafe.Add(mBase, uint32(v396))) = uint8(v400)
	v402 = int32(1)
	v403 = v394 + v402
	v405 = v395 - v402
	v407 = v396 + v402
	if v407&int32(3) != 0 {
		v394 = v403
		v395 = v405
		v396 = v407
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	if v449 == int32(0) {
		goto L109
	} else {
		goto L142
	}
L130:
	;
	if v377&int32(3) != 0 {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v414 = v371
	goto L134
L132:
	;
	v429 = v371
	goto L133
L133:
	;
	if base.Ui32(v429) <= base.Ui32(int32(3)) {
		v449 = v429
		goto L129
	} else {
		goto L138
	}
L134:
	;
	if v414 == int32(0) {
		goto L109
	} else {
		goto L136
	}
L135:
	;
	v429 = v420
	goto L133
L136:
	;
	v420 = v414 - int32(1)
	v421 = v357 + v420
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v420))))
	*(*uint8)(unsafe.Add(mBase, uint32(v421))) = uint8(v423)
	if v421&int32(3) != 0 {
		v414 = v420
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v436 = v429
	goto L139
L139:
	;
	v440 = v436 - int32(4)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v373+v440)))
	*(*int32)(unsafe.Add(mBase, uint32(v357+v440))) = v443
	if base.Ui32(int32(3)) < base.Ui32(v440) {
		v436 = v440
		goto L139
	} else {
		goto L141
	}
L140:
	;
	v449 = v440
	goto L129
L141:
	;
	goto L140
L142:
	;
	v456 = v449
	goto L143
L143:
	;
	v460 = v456 - int32(1)
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373+v460))))
	*(*uint8)(unsafe.Add(mBase, uint32(v357+v460))) = uint8(v463)
	if v460 != 0 {
		v456 = v460
		goto L143
	} else {
		goto L145
	}
L144:
	;
	goto L109
L145:
	;
	goto L144
L146:
	;
	v473 = v466
	v474 = v467
	v475 = v468
	goto L147
L147:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v473)))
	*(*int32)(unsafe.Add(mBase, uint32(v475))) = v477
	v479 = int32(4)
	v480 = v473 + v479
	v482 = v475 + v479
	v484 = v474 - v479
	if base.Ui32(int32(3)) < base.Ui32(v484) {
		v473 = v480
		v474 = v484
		v475 = v482
		goto L147
	} else {
		goto L149
	}
L148:
	;
	v488 = v480
	v489 = v484
	v490 = v482
	goto L114
L149:
	;
	goto L148
L150:
	;
	v495 = v488
	v496 = v489
	v497 = v490
	goto L151
L151:
	;
	v499 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v495))))
	*(*uint8)(unsafe.Add(mBase, uint32(v497))) = uint8(v499)
	v501 = int32(1)
	v506 = v496 - v501
	if v506 != 0 {
		v495 = v495 + v501
		v496 = v506
		v497 = v497 + v501
		goto L151
	} else {
		goto L153
	}
L152:
	;
	goto L109
L153:
	;
	goto L152
L154:
	;
	v670 = v181
	goto L86
L155:
	;
	goto L154
L156:
	;
	v528 = v523 + v181
	if base.Ui32(v524-v528) <= base.Ui32(int32(0)-v181<<(uint(int32(1))%32)) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v535 = F___memcpy(m, v523, v524, v181)
	mBase = m.M
	goto L154
L158:
	;
	goto L159
L159:
	;
	v538 = (v523 ^ v524) & int32(3)
	if base.Ui32(v523) < base.Ui32(v524) {
		goto L162
	} else {
		goto L163
	}
L160:
	;
	if v640 == int32(0) {
		goto L155
	} else {
		goto L196
	}
L161:
	;
	if base.Ui32(v618) <= base.Ui32(int32(3)) {
		v639 = v617
		v640 = v618
		v641 = v619
		goto L160
	} else {
		goto L192
	}
L162:
	;
	if v538 != 0 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	goto L164
L164:
	;
	if v538 != 0 {
		v600 = v181
		goto L175
	} else {
		goto L176
	}
L165:
	;
	v639 = v524
	v640 = v181
	v641 = v523
	goto L160
L166:
	;
	goto L167
L167:
	;
	if v523&int32(3) == int32(0) {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	v617 = v524
	v618 = v181
	v619 = v523
	goto L161
L169:
	;
	goto L170
L170:
	;
	v545 = v524
	v546 = v181
	v547 = v523
	goto L171
L171:
	;
	if v546 == int32(0) {
		goto L155
	} else {
		goto L173
	}
L172:
	;
	v617 = v554
	v618 = v556
	v619 = v558
	goto L161
L173:
	;
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v545))))
	*(*uint8)(unsafe.Add(mBase, uint32(v547))) = uint8(v551)
	v553 = int32(1)
	v554 = v545 + v553
	v556 = v546 - v553
	v558 = v547 + v553
	if v558&int32(3) != 0 {
		v545 = v554
		v546 = v556
		v547 = v558
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	if v600 == int32(0) {
		goto L155
	} else {
		goto L188
	}
L176:
	;
	if v528&int32(3) != 0 {
		goto L177
	} else {
		goto L178
	}
L177:
	;
	v565 = v181
	goto L180
L178:
	;
	v580 = v181
	goto L179
L179:
	;
	if base.Ui32(v580) <= base.Ui32(int32(3)) {
		v600 = v580
		goto L175
	} else {
		goto L184
	}
L180:
	;
	if v565 == int32(0) {
		goto L155
	} else {
		goto L182
	}
L181:
	;
	v580 = v571
	goto L179
L182:
	;
	v571 = v565 - int32(1)
	v572 = v523 + v571
	v574 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524+v571))))
	*(*uint8)(unsafe.Add(mBase, uint32(v572))) = uint8(v574)
	if v572&int32(3) != 0 {
		v565 = v571
		goto L180
	} else {
		goto L183
	}
L183:
	;
	goto L181
L184:
	;
	v587 = v580
	goto L185
L185:
	;
	v591 = v587 - int32(4)
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v524+v591)))
	*(*int32)(unsafe.Add(mBase, uint32(v523+v591))) = v594
	if base.Ui32(int32(3)) < base.Ui32(v591) {
		v587 = v591
		goto L185
	} else {
		goto L187
	}
L186:
	;
	v600 = v591
	goto L175
L187:
	;
	goto L186
L188:
	;
	v607 = v600
	goto L189
L189:
	;
	v611 = v607 - int32(1)
	v614 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v524+v611))))
	*(*uint8)(unsafe.Add(mBase, uint32(v523+v611))) = uint8(v614)
	if v611 != 0 {
		v607 = v611
		goto L189
	} else {
		goto L191
	}
L190:
	;
	goto L155
L191:
	;
	goto L190
L192:
	;
	v624 = v617
	v625 = v618
	v626 = v619
	goto L193
L193:
	;
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v624)))
	*(*int32)(unsafe.Add(mBase, uint32(v626))) = v628
	v630 = int32(4)
	v631 = v624 + v630
	v633 = v626 + v630
	v635 = v625 - v630
	if base.Ui32(int32(3)) < base.Ui32(v635) {
		v624 = v631
		v625 = v635
		v626 = v633
		goto L193
	} else {
		goto L195
	}
L194:
	;
	v639 = v631
	v640 = v635
	v641 = v633
	goto L160
L195:
	;
	goto L194
L196:
	;
	v646 = v639
	v647 = v640
	v648 = v641
	goto L197
L197:
	;
	v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v646))))
	*(*uint8)(unsafe.Add(mBase, uint32(v648))) = uint8(v650)
	v652 = int32(1)
	v657 = v647 - v652
	if v657 != 0 {
		v646 = v646 + v652
		v647 = v657
		v648 = v648 + v652
		goto L197
	} else {
		goto L199
	}
L198:
	;
	goto L155
L199:
	;
	goto L198
L200:
	;
	v688 = l2
	goto L202
L201:
	;
	v688 = v684
	goto L202
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v684 - v688
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v688 + v685
	v694 = v688
	goto L1
}
func F_mdsyncfiletag(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v375 int32
	_ = v375
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v401 int64
	_ = v401
	var v402 int64
	_ = v402
	var v406 int64
	_ = v406
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v422 int64
	_ = v422
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v434 int64
	_ = v434
	var v435 int64
	_ = v435
	var v439 int64
	_ = v439
	var v491 int32
	_ = v491
	var v492 int64
	_ = v492
	var v496 int32
	_ = v496
	var v514 int32
	_ = v514
	var v515 int64
	_ = v515
	var v519 int32
	_ = v519
	var v536 int32
	_ = v536
	var v537 int64
	_ = v537
	var v542 int32
	_ = v542
	var v543 int64
	_ = v543
	var v548 int32
	_ = v548
	var v559 int32
	_ = v559
	v9 = m.G0
	v11 = v9 - int32(192)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v13
	v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v15
	v20 = F_smgropen(m, v11+int32(8), int32(-1))
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
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v25 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+2)))
	v28 = v20 + v25<<(uint(int32(2))%32)
	v29 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+40)))
	if base.Ui64(v24) < base.Ui64(v29) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	m.G0 = v11 + int32(192)
	return v559
L4:
	;
	v392 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1185])))
	v395 = m.G0
	v397 = v395 - int32(16)
	m.G0 = v397
	if v392 != 0 {
		goto L110
	} else {
		goto L111
	}
L5:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+56))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v31+base.I32_wrap_i64(v24)<<(uint(int32(3))%32))))
	v38 = *(*int32)(unsafe.Add(mBase, _consts[943]))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38+v36*int32(48))+32))
	goto L8
L6:
	;
	goto L7
L7:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	F_GetRelationPath(m, v11+int32(109), v160, v161, v162, v163, v25)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L41
	}
L8:
	;
	goto L12
L9:
	;
	v390 = v36
	goto L4
L10:
	;
	v155 = F_strlen(m, v144)
	mBase = m.M
	goto L9
L12:
	;
	goto L13
L13:
	;
	v49 = int32(1023)
	if (l1^v42)&int32(3) != 0 {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v145))) = uint8(v148)
	goto L10
L15:
	;
	v129 = v124
	v130 = v125
	v131 = v126
	goto L37
L16:
	;
	if v119 == int32(0) {
		v144 = v117
		v145 = v118
		goto L14
	} else {
		goto L36
	}
L17:
	;
	v117 = v42
	v118 = l1
	v119 = v49
	goto L16
L18:
	;
	goto L19
L19:
	;
	if v42&int32(3) == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	if v86 == int32(0) {
		v144 = v83
		v145 = v84
		goto L14
	} else {
		goto L29
	}
L21:
	;
	v83 = v42
	v84 = l1
	v85 = v49
	v86 = int32(1)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v62 = v42
	v63 = l1
	v64 = v49
	goto L24
L24:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	*(*uint8)(unsafe.Add(mBase, uint32(v63))) = uint8(v66)
	if v66 == int32(0) {
		v124 = v62
		v125 = v63
		v126 = v64
		goto L15
	} else {
		goto L26
	}
L25:
	;
	v83 = v77
	v84 = v71
	v85 = v73
	v86 = v75
	goto L20
L26:
	;
	v70 = int32(1)
	v71 = v63 + v70
	v73 = v64 - v70
	v74 = int32(0)
	v75 = base.B2i32(v73 != v74)
	v77 = v62 + v70
	if v77&int32(3) == v74 {
		v83 = v77
		v84 = v71
		v85 = v73
		v86 = v75
		goto L20
	} else {
		goto L27
	}
L27:
	;
	if v73 != 0 {
		v62 = v77
		v63 = v71
		v64 = v73
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	if v89 == int32(0) {
		v117 = v83
		v118 = v84
		v119 = v85
		goto L16
	} else {
		goto L30
	}
L30:
	;
	if base.Ui32(v85) < base.Ui32(int32(4)) {
		v117 = v83
		v118 = v84
		v119 = v85
		goto L16
	} else {
		goto L31
	}
L31:
	;
	v95 = v83
	v96 = v84
	v97 = v85
	goto L32
L32:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v103 = int32(-2139062144)
	if (int32(16843008)-v100|v100)&v103 != v103 {
		v124 = v95
		v125 = v96
		v126 = v97
		goto L15
	} else {
		goto L34
	}
L33:
	;
	v117 = v111
	v118 = v109
	v119 = v113
	goto L16
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v100
	v108 = int32(4)
	v109 = v96 + v108
	v111 = v95 + v108
	v113 = v97 - v108
	if base.Ui32(int32(3)) < base.Ui32(v113) {
		v95 = v111
		v96 = v109
		v97 = v113
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v124 = v117
	v125 = v118
	v126 = v119
	goto L15
L37:
	;
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v129))))
	*(*uint8)(unsafe.Add(mBase, uint32(v130))) = uint8(v133)
	if v133 == int32(0) {
		v144 = v129
		v145 = v130
		goto L14
	} else {
		goto L39
	}
L38:
	;
	v144 = v140
	v145 = v138
	goto L14
L39:
	;
	v137 = int32(1)
	v138 = v130 + v137
	v140 = v129 + v137
	v142 = v131 - v137
	if v142 != 0 {
		v129 = v140
		v130 = v138
		v131 = v142
		goto L37
	} else {
		goto L40
	}
L40:
	;
	goto L38
L41:
	;
	v166 = base.I32_wrap_i64(v24)
	if v166 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L69
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v166
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(109)
	v174 = F_pg_sprintf(m, v11+int32(26), int32(38675), v11)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v177 = v11 + int32(26)
	v179 = v11 + int32(109)
	if (v179^v177)&int32(3) != 0 {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	goto L42
L47:
	;
	goto L42
L48:
	;
	goto L47
L49:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v234))) = uint8(v233)
	if v233&int32(255) == int32(0) {
		goto L48
	} else {
		goto L64
	}
L50:
	;
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179))))
	v232 = v179
	v233 = v185
	v234 = v177
	goto L49
L51:
	;
	goto L52
L52:
	;
	if v179&int32(3) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v189 = v179
	v191 = v177
	goto L56
L54:
	;
	v203 = v179
	v205 = v177
	goto L55
L55:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	v210 = int32(-2139062144)
	if (int32(16843008)-v207|v207)&v210 != v210 {
		v232 = v203
		v233 = v207
		v234 = v205
		goto L49
	} else {
		goto L60
	}
L56:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	*(*uint8)(unsafe.Add(mBase, uint32(v191))) = uint8(v192)
	if v192 == int32(0) {
		goto L48
	} else {
		goto L58
	}
L57:
	;
	v203 = v199
	v205 = v197
	goto L55
L58:
	;
	v196 = int32(1)
	v197 = v191 + v196
	v199 = v189 + v196
	if v199&int32(3) != 0 {
		v189 = v199
		v191 = v197
		goto L56
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	v215 = v203
	v216 = v207
	v217 = v205
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217))) = v216
	v219 = int32(4)
	v220 = v217 + v219
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v215)+4))
	v223 = v215 + v219
	v227 = int32(-2139062144)
	if (v221|(int32(16843008)-v221))&v227 == v227 {
		v215 = v223
		v216 = v221
		v217 = v220
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v232 = v223
	v233 = v221
	v234 = v220
	goto L49
L63:
	;
	goto L62
L64:
	;
	v241 = v232
	v243 = v234
	goto L65
L65:
	;
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v241)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v243)+1)) = uint8(v244)
	v246 = int32(1)
	if v244 != 0 {
		v241 = v241 + v246
		v243 = v243 + v246
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
	v262 = v11 + int32(109)
	goto L75
L69:
	;
	v259 = F__emscripten_memcpy_bulkmem(m, v11+int32(109), v11+int32(26), int32(83))
	mBase = m.M
	goto L71
L71:
	;
	goto L68
L72:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	if v381&int32(1) != 0 {
		goto L104
	} else {
		goto L105
	}
L73:
	;
	v375 = F_strlen(m, v364)
	mBase = m.M
	goto L72
L75:
	;
	goto L76
L76:
	;
	v269 = int32(81)
	if (l1^v262)&int32(3) != 0 {
		goto L80
	} else {
		goto L81
	}
L77:
	;
	v368 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v365))) = uint8(v368)
	goto L73
L78:
	;
	v349 = v344
	v350 = v345
	v351 = v346
	goto L100
L79:
	;
	if v339 == int32(0) {
		v364 = v337
		v365 = v338
		goto L77
	} else {
		goto L99
	}
L80:
	;
	v337 = v262
	v338 = l1
	v339 = v269
	goto L79
L81:
	;
	goto L82
L82:
	;
	if v262&int32(3) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v306 == int32(0) {
		v364 = v303
		v365 = v304
		goto L77
	} else {
		goto L92
	}
L84:
	;
	v303 = v262
	v304 = l1
	v305 = v269
	v306 = int32(1)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v282 = v262
	v283 = l1
	v284 = v269
	goto L87
L87:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v282))))
	*(*uint8)(unsafe.Add(mBase, uint32(v283))) = uint8(v286)
	if v286 == int32(0) {
		v344 = v282
		v345 = v283
		v346 = v284
		goto L78
	} else {
		goto L89
	}
L88:
	;
	v303 = v297
	v304 = v291
	v305 = v293
	v306 = v295
	goto L83
L89:
	;
	v290 = int32(1)
	v291 = v283 + v290
	v293 = v284 - v290
	v294 = int32(0)
	v295 = base.B2i32(v293 != v294)
	v297 = v282 + v290
	if v297&int32(3) == v294 {
		v303 = v297
		v304 = v291
		v305 = v293
		v306 = v295
		goto L83
	} else {
		goto L90
	}
L90:
	;
	if v293 != 0 {
		v282 = v297
		v283 = v291
		v284 = v293
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v303))))
	if v309 == int32(0) {
		v337 = v303
		v338 = v304
		v339 = v305
		goto L79
	} else {
		goto L93
	}
L93:
	;
	if base.Ui32(v305) < base.Ui32(int32(4)) {
		v337 = v303
		v338 = v304
		v339 = v305
		goto L79
	} else {
		goto L94
	}
L94:
	;
	v315 = v303
	v316 = v304
	v317 = v305
	goto L95
L95:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v315)))
	v323 = int32(-2139062144)
	if (int32(16843008)-v320|v320)&v323 != v323 {
		v344 = v315
		v345 = v316
		v346 = v317
		goto L78
	} else {
		goto L97
	}
L96:
	;
	v337 = v331
	v338 = v329
	v339 = v333
	goto L79
L97:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v316))) = v320
	v328 = int32(4)
	v329 = v316 + v328
	v331 = v315 + v328
	v333 = v317 - v328
	if base.Ui32(int32(3)) < base.Ui32(v333) {
		v315 = v331
		v316 = v329
		v317 = v333
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v344 = v337
	v345 = v338
	v346 = v339
	goto L78
L100:
	;
	v353 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v349))))
	*(*uint8)(unsafe.Add(mBase, uint32(v350))) = uint8(v353)
	if v353 == int32(0) {
		v364 = v349
		v365 = v350
		goto L77
	} else {
		goto L102
	}
L101:
	;
	v364 = v360
	v365 = v358
	goto L77
L102:
	;
	v357 = int32(1)
	v358 = v350 + v357
	v360 = v349 + v357
	v362 = v351 - v357
	if v362 != 0 {
		v349 = v360
		v350 = v358
		v351 = v362
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	v384 = int32(16386)
	goto L106
L105:
	;
	v384 = int32(2)
	goto L106
L106:
	;
	v385 = F_PathNameOpenFile(m, l1, v384)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	if int32(0) <= v385 {
		v390 = v385
		goto L4
	} else {
		goto L108
	}
L108:
	;
	v559 = int32(-1)
	goto L3
L109:
	;
	v411 = F_FileSync(m, v390, int32(167772182))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L113
	}
L110:
	;
	F___clock_gettime(m, int32(1), v397)
	mBase = m.M
	v401 = int64(*(*int32)(unsafe.Add(mBase, uint32(v397)+8)))
	v402 = *(*int64)(unsafe.Add(mBase, uint32(v397)))
	v406 = v401 + v402*int64(1000000000)
	goto L112
L111:
	;
	v406 = int64(0)
	goto L112
L112:
	;
	m.G0 = v397 + int32(16)
	goto L109
L113:
	;
	v414 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if base.Ui64(v29) <= base.Ui64(v24) {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	F_FileClose(m, v390)
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v420 = int32(1)
	v422 = int64(0)
	v426 = m.G0
	v428 = v426 - int32(16)
	m.G0 = v428
	if v406 != v422 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L116
L118:
	;
	*(*int32)(unsafe.Add(mBase, _consts[140])) = v414
	v559 = v411
	goto L3
L119:
	;
	F___clock_gettime(m, int32(1), v428)
	mBase = m.M
	v434 = int64(*(*int32)(unsafe.Add(mBase, uint32(v428)+8)))
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v428)))
	v439 = v434 + (v435*int64(1000000000) - v406)
	goto L123
L120:
	;
	goto L121
L121:
	;
	v536 = int32(4451504)
	v537 = *(*int64)(unsafe.Add(mBase, _consts[1186]))
	*(*int64)(unsafe.Add(mBase, _consts[1186])) = v537 + base.I64_extend_i32_u(v420)
	v542 = int32(4450544)
	v543 = *(*int64)(unsafe.Add(mBase, _consts[1187]))
	*(*int64)(unsafe.Add(mBase, _consts[1187])) = v543 + v422
	F_pgstat_count_backend_io_op(m, int32(0), int32(3), v420, v420, v422)
	mBase = m.M
	v548 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[145])) = uint8(v548)
	*(*uint8)(unsafe.Add(mBase, _consts[873])) = uint8(v548)
	m.G0 = v428 + int32(16)
	goto L118
L122:
	;
	v491 = int32(4452464)
	v492 = *(*int64)(unsafe.Add(mBase, _consts[1188]))
	*(*int64)(unsafe.Add(mBase, _consts[1188])) = v492 + v439
	v496 = *(*int32)(unsafe.Add(mBase, _consts[264]))
	if base.Ui32(int32(16)) < base.Ui32(v496) {
		goto L132
	} else {
		goto L133
	}
L123:
	;
	goto L125
L125:
	;
	goto L126
L126:
	;
	goto L122
L132:
	;
	goto L121
L133:
	;
	if int32(1)<<(uint(v496)%32)&int32(115186) == int32(0) {
		goto L132
	} else {
		goto L134
	}
L134:
	;
	v514 = int32(4449360)
	v515 = *(*int64)(unsafe.Add(mBase, _consts[1189]))
	*(*int64)(unsafe.Add(mBase, _consts[1189])) = v515 + v439
	v519 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[145])) = uint8(v519)
	*(*uint8)(unsafe.Add(mBase, _consts[876])) = uint8(v519)
	goto L132
}
func F_merge_clump(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	if l1 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v183 = F_lappend(m, v176, v177)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L14
	} else {
		goto L55
	}
L2:
	;
	v176 = int32(0)
	v177 = l2
	goto L1
L3:
	;
	goto L4
L4:
	;
	v13 = l1
	v14 = l2
	goto L5
L5:
	;
	v20 = int32(0)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v21 <= v20 {
		v139 = v21
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v143 == int32(1) {
		v176 = v13
		v177 = v14
		goto L1
	} else {
		goto L44
	}
L7:
	;
	goto L6
L8:
	;
	v29 = v20
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v29<<(uint(int32(2))%32))))
	if l3 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v139 = v133
	goto L7
L11:
	;
	v132 = v29 + int32(1)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v132 < v133 {
		v29 = v132
		goto L9
	} else {
		goto L43
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v51 = F_make_join_rel(m, l0, v49, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L14
	} else {
		goto L19
	}
L13:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v39 = F_have_relevant_joinclause(m, l0, v37, v38)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	return int32(0)
L15:
	;
	if v39 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v43 = F_have_join_order_restriction(m, l0, v37, v38)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L14
	} else {
		goto L17
	}
L17:
	;
	if v43 == int32(0) {
		goto L11
	} else {
		goto L18
	}
L18:
	;
	goto L12
L19:
	;
	if v51 == int32(0) {
		goto L11
	} else {
		goto L20
	}
L20:
	;
	F_generate_partitionwise_join_paths(m, l0, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L21
	}
L21:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v59 = int32(0)
	v66 = base.B2i32(v57|v58 == v59)
	if v57 == v59 {
		v105 = v66
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v105 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L23:
	;
	goto L22
L24:
	;
	if v58 == int32(0) {
		v105 = v66
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	if v72 != v73 {
		v105 = int32(0)
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v75 = int32(1)
	if v72 <= v75 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v78 = v75
	goto L29
L28:
	;
	v78 = v72
	goto L29
L29:
	;
	v79 = int32(8)
	v84 = int32(0)
	goto L30
L30:
	;
	v92 = v84 << (uint(int32(2)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v57+v79+v92)))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+(v58+v79))))
	v97 = base.B2i32(v94 == v96)
	if v96 != v94 {
		v105 = v97
		goto L23
	} else {
		goto L32
	}
L31:
	;
	v105 = v97
	goto L23
L32:
	;
	v100 = v84 + int32(1)
	if v100 != v78 {
		v84 = v100
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	F_generate_useful_gather_paths(m, l0, v51, int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L14
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	F_set_cheapest(m, v51)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L14
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36))) = v51
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v117 + v118
	F_pfree(m, v14)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L14
	} else {
		goto L39
	}
L39:
	;
	v123 = F_list_delete_nth_cell(m, v13, v29)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L14
	} else {
		goto L40
	}
L40:
	;
	if v123 != 0 {
		v13 = v123
		v14 = v36
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v126 = F_lappend(m, int32(0), v36)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L14
	} else {
		goto L42
	}
L42:
	;
	return v126
L43:
	;
	goto L10
L44:
	;
	if v139 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v149 = F_list_insert_nth(m, v13, int32(0), v14)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L14
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v159 = int32(0)
	goto L50
L48:
	;
	return v149
L49:
	;
	v172 = F_list_insert_nth(m, v13, v171, v14)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L14
	} else {
		goto L54
	}
L50:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v152+v159<<(uint(int32(2))%32))))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+4))
	if v166 < v143 {
		v171 = v159
		goto L49
	} else {
		goto L52
	}
L51:
	;
	v171 = v139
	goto L49
L52:
	;
	v169 = v159 + int32(1)
	if v169 != v139 {
		v159 = v169
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	return v172
L55:
	;
	return v183
}
func F_miss(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v46 int32
	_ = v46
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v572 int32
	_ = v572
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v599 int32
	_ = v599
	var v606 int32
	_ = v606
	var v614 int32
	_ = v614
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v660 int32
	_ = v660
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v707 int32
	_ = v707
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v761 int32
	_ = v761
	var v763 int32
	_ = v763
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v813 int32
	_ = v813
	var v828 int32
	_ = v828
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v864 int32
	_ = v864
	var v887 int32
	_ = v887
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v923 int32
	_ = v923
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v941 int32
	_ = v941
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v962 int32
	_ = v962
	var v965 int32
	_ = v965
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v992 int32
	_ = v992
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1044 int32
	_ = v1044
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1061 int32
	_ = v1061
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1133 int32
	_ = v1133
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1165 int32
	_ = v1165
	var v1167 int32
	_ = v1167
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1243 int32
	_ = v1243
	var v1248 int32
	_ = v1248
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1292 int32
	_ = v1292
	var v1298 int32
	_ = v1298
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1316 int32
	_ = v1316
	var v1322 int32
	_ = v1322
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1335 int32
	_ = v1335
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1358 int32
	_ = v1358
	var v1361 int32
	_ = v1361
	var v1368 int32
	_ = v1368
	var v1373 int32
	_ = v1373
	var v1376 int32
	_ = v1376
	var v1379 int32
	_ = v1379
	var v1385 int32
	_ = v1385
	var v1401 int32
	_ = v1401
	var v1408 int32
	_ = v1408
	var v1411 int32
	_ = v1411
	var v1412 int32
	_ = v1412
	var v1416 int32
	_ = v1416
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1438 int32
	_ = v1438
	var v1451 int32
	_ = v1451
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1473 int32
	_ = v1473
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1493 int32
	_ = v1493
	var v1495 int32
	_ = v1495
	var v1503 int32
	_ = v1503
	var v1505 int32
	_ = v1505
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1522 int64
	_ = v1522
	var v1524 int32
	_ = v1524
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1545 int32
	_ = v1545
	var v1547 int32
	_ = v1547
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1556 int64
	_ = v1556
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1572 int32
	_ = v1572
	var v1578 int32
	_ = v1578
	var v1583 int32
	_ = v1583
	var v1590 int32
	_ = v1590
	var v1602 int32
	_ = v1602
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1627 int32
	_ = v1627
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1658 int32
	_ = v1658
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1705 int32
	_ = v1705
	var v1708 int32
	_ = v1708
	var v1723 int32
	_ = v1723
	var v1737 int32
	_ = v1737
	var v1742 int32
	_ = v1742
	var v1746 int64
	_ = v1746
	var v1768 int32
	_ = v1768
	v7 = int32(0)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v23+l3<<(uint(int32(2))%32))))
	if v27 != 0 {
		v1768 = v27
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v1768
L2:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v30 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if int32(0) < v35 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L5
L8:
	;
	v46 = v7
	goto L11
L9:
	;
	goto L10
L10:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v92 <= int32(0) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v60+v46<<(uint(int32(2))%32)))) = int32(0)
	v67 = v46 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v67 < v68 {
		v46 = v67
		goto L11
	} else {
		goto L13
	}
L12:
	;
	goto L10
L13:
	;
	goto L12
L14:
	;
	return int32(0)
L15:
	;
	goto L16
L16:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+20))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+l3*int32(24))+20))
	v115 = v92
	v119 = v7
	v120 = v7
	v124 = int32(1)
	v125 = v7
	goto L17
L17:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v130+int32(base.Ui32(v120)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v136)>>(uint(v120)%32))&int32(1) == int32(0) {
		v228 = v115
		v232 = v119
		v237 = v124
		v238 = v125
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v246 = int32(0)
	if v232 == v246 {
		v1768 = v246
		goto L1
	} else {
		goto L35
	}
L19:
	;
	v244 = v120 + int32(1)
	if v244 < v228 {
		v115 = v228
		v119 = v232
		v120 = v244
		v124 = v237
		v125 = v238
		goto L17
	} else {
		goto L34
	}
L20:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v28)+32))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v142+v120<<(uint(int32(2))%32))))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146))))
	if v147 == int32(65535) {
		v228 = v115
		v232 = v119
		v237 = v124
		v238 = v125
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v156 = v147
	v158 = v146
	v161 = v119
	v166 = v124
	v167 = v125
	goto L22
L22:
	;
	v172 = int32(65535)
	v173 = v156 & v172
	if base.B2i32(v173 != int32(65534))|base.B2i32(v102&int32(2) != int32(0)) != 0 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v228 = v220
	v232 = v212
	v237 = v213
	v238 = v214
	goto L19
L24:
	;
	v181 = base.B2i32(v173 != l3&v172)
	goto L26
L25:
	;
	v181 = int32(0)
	goto L26
L26:
	;
	if v181 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v190 = v184 + int32(base.Ui32(v185)>>(uint(int32(3))%32))&int32(536870908)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v192 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v191 | v192<<(uint(v185)%32)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v197 == v198 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v212 = v161
	v213 = v166
	v214 = v167
	goto L29
L29:
	;
	v216 = v158 + int32(8)
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v216))))
	if v217 != int32(65535) {
		v156 = v217
		v158 = v216
		v161 = v212
		v166 = v213
		v167 = v214
		goto L22
	} else {
		goto L33
	}
L30:
	;
	v200 = v192
	goto L32
L31:
	;
	v200 = v167
	goto L32
L32:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v28)+28))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202+v197))))
	v205 = int32(1)
	v212 = v205
	v213 = (int32(0) - v204&v205) & v166
	v214 = v200
	goto L29
L33:
	;
	goto L23
L34:
	;
	goto L18
L35:
	;
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+8)))
	if v249&int32(1) != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v253 = l0
	v254 = l1
	v255 = l2
	v256 = l3
	v257 = l4
	v258 = l5
	v260 = v228
	v262 = int32(0)
	v263 = v28
	v269 = v237
	v270 = v238
	v271 = v246
	goto L39
L37:
	;
	v992 = l0
	v993 = l1
	v994 = l2
	v995 = l3
	v996 = l4
	v997 = l5
	v1008 = v237
	v1009 = v238
	v1010 = v246
	v1014 = int32(1)
	goto L38
L38:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v993)+28))
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v993)+16))
	if v1016 == int32(1) {
		goto L163
	} else {
		goto L164
	}
L39:
	;
	v275 = int32(0)
	if v260 <= v275 {
		v967 = v253
		v968 = v254
		v969 = v255
		v970 = v256
		v971 = v257
		v972 = v258
		v976 = v262
		v983 = v269
		v984 = v270
		v985 = v271
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v992 = v967
	v993 = v968
	v994 = v969
	v995 = v970
	v996 = v971
	v997 = v972
	v1008 = v983
	v1009 = v984
	v1010 = v985
	v1014 = base.B2i32(v976 == int32(0))
	goto L38
L41:
	;
	goto L40
L42:
	;
	v279 = v253
	v280 = v254
	v281 = v255
	v282 = v256
	v283 = v257
	v284 = v258
	v286 = v260
	v288 = v262
	v289 = v263
	v291 = v275
	v295 = v269
	v296 = v270
	v297 = v271
	v299 = v275
	goto L43
L43:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v280)+28))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v301+int32(base.Ui32(v291)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v307)>>(uint(v291)%32))&int32(1) == int32(0) {
		v949 = v286
		v951 = v288
		v958 = v295
		v959 = v296
		v962 = v299
		goto L45
	} else {
		goto L46
	}
L44:
	;
	if v962 != 0 {
		v253 = v279
		v254 = v280
		v255 = v281
		v256 = v282
		v257 = v283
		v258 = v284
		v260 = v949
		v262 = v951
		v263 = v289
		v269 = v958
		v270 = v959
		v271 = v297
		goto L39
	} else {
		goto L161
	}
L45:
	;
	v965 = v291 + int32(1)
	if v965 < v949 {
		v286 = v949
		v288 = v951
		v291 = v965
		v295 = v958
		v296 = v959
		v299 = v962
		goto L43
	} else {
		goto L160
	}
L46:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v289)+32))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v313+v291<<(uint(int32(2))%32))))
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v317))))
	if v318 == int32(65535) {
		v949 = v286
		v951 = v288
		v958 = v295
		v959 = v296
		v962 = v299
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v327 = v318
	v329 = v317
	v330 = v288
	v337 = v295
	v338 = v296
	v341 = v299
	goto L48
L48:
	;
	v343 = base.I32_extend16_s(v327)
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	if v343 < v344 {
		v923 = v330
		v930 = v337
		v931 = v338
		v934 = v341
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v941 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v949 = v941
	v951 = v923
	v958 = v930
	v959 = v931
	v962 = v934
	goto L45
L50:
	;
	v937 = v329 + int32(8)
	v938 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v937))))
	if v938 != int32(65535) {
		v327 = v938
		v329 = v937
		v330 = v923
		v337 = v930
		v338 = v931
		v341 = v934
		goto L48
	} else {
		goto L159
	}
L51:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v280)+28))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v346+int32(base.Ui32(v347)>>(uint(int32(3))%32))&int32(536870908))))
	if int32(base.Ui32(v353)>>(uint(v347)%32))&int32(1) != 0 {
		v923 = v330
		v930 = v337
		v931 = v338
		v934 = v341
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v279)))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)+28))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v358)+4))
	v360 = m.T0[v359].(func(*base.Module) int32)(m)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L6
	} else {
		goto L53
	}
L53:
	;
	if v360 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v279)+36))
	if v362 != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v279)+4))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)+424))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
	v371 = v343 - v370
	v373 = v371 << (uint(int32(2)) % 32)
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v279)+44))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v373+v374)))
	if v376 != 0 {
		v394 = v376
		goto L62
	} else {
		goto L63
	}
L57:
	;
	v364 = v362
	goto L59
L58:
	;
	v364 = int32(19)
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v279)+36)) = v364
	return int32(0)
L60:
	;
	if v857 != 0 {
		goto L153
	} else {
		goto L154
	}
L61:
	;
	if v864 == int32(0) {
		v923 = int32(1)
		v930 = v337
		v931 = v338
		v934 = v341
		goto L50
	} else {
		goto L152
	}
L62:
	;
	v398 = int32(2)
	v399 = v369 + v371*int32(88) + v398
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	if v400&v398 != 0 {
		goto L67
	} else {
		goto L68
	}
L63:
	;
	v385 = F_newdfa(m, v279, v369+v371*int32(88)+int32(36), v368+int32(72), int32(0))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L6
	} else {
		goto L64
	}
L64:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v279)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v387+v373))) = v385
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v279)+44))
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v390+v373)))
	if v392 != 0 {
		v394 = v392
		goto L62
	} else {
		goto L65
	}
L65:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v279)+36))
	v864 = v393
	goto L61
L66:
	;
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v279)+36))
	if v856 != 0 {
		goto L60
	} else {
		goto L151
	}
L67:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v279)+32))
	v404 = int32(0)
	v406 = F_shortest(m, v279, v394, v283, v283, v403, v404, v404)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L6
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v394)+40))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v416)+8)))
	if v417&int32(2) != 0 {
		goto L72
	} else {
		goto L73
	}
L70:
	;
	v408 = int32(0)
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	v856 = base.B2i32(v406 != v408) ^ base.B2i32(v410&int32(1) == v408)
	goto L66
L71:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v399))))
	v856 = v813 ^ (v828^int32(-1))&int32(1)
	goto L66
L72:
	;
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v416)+40))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v279)+24))
	v813 = base.B2i32(base.Ui32(v420) <= base.Ui32((v283-v421)>>(uint(int32(2))%32)))
	goto L71
L73:
	;
	goto L74
L74:
	;
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v279)+48))
	v427 = v426 + v373
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v394)+44))
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v279)+52))
	v430 = v429 + v373
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v430)))
	if base.Ui32(v431) <= base.Ui32(v283) {
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if base.Ui32(v283) <= base.Ui32(v681) {
		goto L124
	} else {
		goto L125
	}
L76:
	;
	v434 = v431
	goto L78
L77:
	;
	v434 = int32(0)
	goto L78
L78:
	;
	if v434 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v437 = int32(0)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v279)+24))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v437 < v447 {
		goto L85
	} else {
		goto L86
	}
L80:
	;
	goto L81
L81:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
	if v679 != 0 {
		v681 = v431
		v683 = v679
		goto L75
	} else {
		goto L122
	}
L82:
	;
	if v660 == int32(0) {
		v813 = v437
		goto L71
	} else {
		goto L119
	}
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v639)+20)) = v438
	*(*int64)(unsafe.Add(mBase, uint32(v394)+48)) = int64(0)
	v660 = v639
	goto L82
L84:
	;
	v614 = int32(0)
	goto L116
L85:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v394)+20))
	v451 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v450)+8)))
	if v451&int32(1) != 0 {
		v606 = v450
		goto L84
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v455 = F_getvacant(m, v279, v394, v438, v438)
	mBase = m.M
	if v455 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L87
L89:
	;
	v660 = int32(0)
	goto L82
L90:
	;
	goto L91
L91:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v394)+16))
	if int32(0) < v459 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v463 = int32(0)
	goto L95
L93:
	;
	goto L94
L94:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v394)+40))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v496)+12))
	v502 = v495 + int32(base.Ui32(v497)>>(uint(int32(3))%32))&int32(536870908)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v502)))
	v504 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v502))) = v503 | v504<<(uint(v497)%32)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v394)+16))
	if v509 == v504 {
		goto L99
	} else {
		goto L100
	}
L95:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v455)))
	*(*int32)(unsafe.Add(mBase, uint32(v474+v463<<(uint(int32(2))%32)))) = int32(0)
	v481 = v463 + int32(1)
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v394)+16))
	if v481 < v482 {
		v463 = v481
		goto L95
	} else {
		goto L97
	}
L96:
	;
	goto L94
L97:
	;
	goto L96
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v455)+8)) = int32(13)
	*(*int32)(unsafe.Add(mBase, uint32(v455)+4)) = v588
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v599 <= int32(0) {
		v639 = v455
		goto L83
	} else {
		goto L115
	}
L99:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v508)))
	v588 = v512
	goto L98
L100:
	;
	goto L101
L101:
	;
	if v509 <= int32(0) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v588 = int32(0)
	goto L98
L103:
	;
	goto L104
L104:
	;
	v517 = v509 & int32(3)
	v518 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v509) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v524 = v518
	v527 = v518
	v533 = v437
	goto L108
L106:
	;
	v551 = v518
	v554 = v518
	goto L107
L107:
	;
	if v517 == int32(0) {
		v588 = v554
		goto L98
	} else {
		goto L111
	}
L108:
	;
	v537 = v508 + v524<<(uint(int32(2))%32)
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v537)+12))
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v537)+8))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v537)))
	v545 = v538 ^ (v539 ^ (v540 ^ (v541 ^ v527)))
	v546 = int32(4)
	v547 = v524 + v546
	v549 = v533 + v546
	if v549 != v509&int32(2147483644) {
		v524 = v547
		v527 = v545
		v533 = v549
		goto L108
	} else {
		goto L110
	}
L109:
	;
	v551 = v547
	v554 = v545
	goto L107
L110:
	;
	goto L109
L111:
	;
	v564 = v551
	v567 = v554
	v572 = v437
	goto L112
L112:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v508+v564<<(uint(int32(2))%32))))
	v579 = v578 ^ v567
	v580 = int32(1)
	v583 = v572 + v580
	if v583 != v517 {
		v564 = v564 + v580
		v567 = v579
		v572 = v583
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v588 = v579
	goto L98
L114:
	;
	goto L113
L115:
	;
	v606 = v455
	goto L84
L116:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v394)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v625+v614<<(uint(int32(5))%32))+20)) = int32(0)
	v632 = v614 + int32(1)
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v632 < v633 {
		v614 = v632
		goto L116
	} else {
		goto L118
	}
L117:
	;
	v639 = v606
	goto L83
L118:
	;
	goto L117
L119:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v394)+40))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v279)+8))
	v667 = int32(1)
	v672 = int32(*(*int16)(unsafe.Add(mBase, uint32(v663+(v664^int32(-1))&v667<<(uint(v667)%32))+20)))
	v673 = *(*int32)(unsafe.Add(mBase, uint32(v279)+24))
	v674 = F_miss(m, v279, v394, v660, v672, v438, v673)
	mBase = m.M
	v675 = m.ExcPending
	if v675 != 0 {
		goto L6
	} else {
		goto L120
	}
L120:
	;
	if v674 == int32(0) {
		v813 = v437
		goto L71
	} else {
		goto L121
	}
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v674)+20)) = v438
	v681 = v438
	v683 = v674
	goto L75
L122:
	;
	v813 = int32(0)
	goto L71
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = v744
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = v743
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v279)+32))
	if base.Ui32(v743) < base.Ui32(v761) {
		goto L140
	} else {
		goto L141
	}
L124:
	;
	v743 = v681
	v744 = v683
	goto L123
L125:
	;
	goto L126
L126:
	;
	v691 = v681
	v694 = v683
	goto L127
L127:
	;
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v691)))
	if base.Ui32(v707) <= base.Ui32(int32(2047)) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v743 = v734
	v744 = v732
	goto L123
L129:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v694)+24))
	v718 = base.I32_extend16_s(v716)
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v717+v718<<(uint(int32(2))%32))))
	if v722 != 0 {
		v732 = v722
		goto L133
	} else {
		goto L134
	}
L130:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v428)+24))
	v714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v710+v707<<(uint(int32(1))%32)))))
	v716 = v714
	goto L129
L131:
	;
	goto L132
L132:
	;
	v715 = F_pg_reg_getcolor(m, v428, v707)
	mBase = m.M
	v716 = v715
	goto L129
L133:
	;
	v734 = v691 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v732)+20)) = v734
	if base.Ui32(v734) < base.Ui32(v283) {
		v691 = v734
		v694 = v732
		goto L127
	} else {
		goto L137
	}
L134:
	;
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v279)+24))
	v726 = F_miss(m, v279, v394, v694, v718, v691+int32(4), v725)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L6
	} else {
		goto L135
	}
L135:
	;
	if v726 != 0 {
		v732 = v726
		goto L133
	} else {
		goto L136
	}
L136:
	;
	v728 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v427))) = v728
	*(*int32)(unsafe.Add(mBase, uint32(v430))) = v691
	v813 = v728
	goto L71
L137:
	;
	goto L128
L138:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(v799)+8))
	v802 = int32(1)
	v813 = int32(base.Ui32(v801)>>(uint(v802)%32)) & v802
	goto L71
L139:
	;
	if v797 != 0 {
		v799 = v797
		goto L138
	} else {
		goto L150
	}
L140:
	;
	v763 = *(*int32)(unsafe.Add(mBase, uint32(v743)))
	if base.Ui32(v763) <= base.Ui32(int32(2047)) {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	goto L142
L142:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v394)+40))
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v279)+8))
	v791 = int32(*(*int16)(unsafe.Add(mBase, uint32(v784+(v785^int32(-1))&int32(2))+24)))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v279)+24))
	v793 = F_miss(m, v279, v394, v744, v791, v743, v792)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L6
	} else {
		goto L149
	}
L143:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v744)+24))
	v774 = base.I32_extend16_s(v772)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v773+v774<<(uint(int32(2))%32))))
	if v778 != 0 {
		v799 = v778
		goto L138
	} else {
		goto L147
	}
L144:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v428)+24))
	v770 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v766+v763<<(uint(int32(1))%32)))))
	v772 = v770
	goto L143
L145:
	;
	goto L146
L146:
	;
	v771 = F_pg_reg_getcolor(m, v428, v763)
	mBase = m.M
	v772 = v771
	goto L143
L147:
	;
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v279)+24))
	v782 = F_miss(m, v279, v394, v744, v774, v743+int32(4), v781)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L6
	} else {
		goto L148
	}
L148:
	;
	v797 = v782
	goto L139
L149:
	;
	v797 = v793
	goto L139
L150:
	;
	v813 = int32(0)
	goto L71
L151:
	;
	v864 = v857
	goto L61
L152:
	;
	return int32(0)
L153:
	;
	return int32(0)
L154:
	;
	goto L155
L155:
	;
	v887 = int32(1)
	v888 = *(*int32)(unsafe.Add(mBase, uint32(v280)+28))
	v889 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v894 = v888 + int32(base.Ui32(v889)>>(uint(int32(3))%32))&int32(536870908)
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v894)))
	*(*int32)(unsafe.Add(mBase, uint32(v894))) = v895 | v887<<(uint(v889)%32)
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v902 = *(*int32)(unsafe.Add(mBase, uint32(v289)+16))
	if v901 == v902 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v904 = v887
	goto L158
L157:
	;
	v904 = v338
	goto L158
L158:
	;
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v289)+28))
	v908 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v906+v901))))
	v909 = int32(1)
	v923 = v909
	v930 = (int32(0) - v908&v909) & v337
	v931 = v904
	v934 = v887
	goto L50
L159:
	;
	goto L49
L160:
	;
	goto L44
L161:
	;
	v967 = v279
	v968 = v280
	v969 = v281
	v970 = v282
	v971 = v283
	v972 = v284
	v976 = v951
	v983 = v958
	v984 = v959
	v985 = v297
	goto L41
L162:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v993)+20))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v993)+4))
	if int32(0) < v1150 {
		goto L181
	} else {
		goto L182
	}
L163:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v1015)))
	v1133 = v1019
	goto L162
L164:
	;
	goto L165
L165:
	;
	if v1016 <= int32(0) {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v1133 = int32(0)
	goto L162
L167:
	;
	goto L168
L168:
	;
	v1024 = v1016 & int32(3)
	v1025 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1016) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1039 = v1025
	v1041 = v1025
	v1044 = int32(0)
	goto L172
L170:
	;
	v1077 = v1025
	v1079 = v1025
	goto L171
L171:
	;
	if v1024 == int32(0) {
		v1133 = v1077
		goto L162
	} else {
		goto L175
	}
L172:
	;
	v1057 = v1015 + v1041<<(uint(int32(2))%32)
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+8))
	v1060 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+4))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1057)))
	v1065 = v1058 ^ (v1059 ^ (v1060 ^ (v1061 ^ v1039)))
	v1066 = int32(4)
	v1067 = v1041 + v1066
	v1069 = v1044 + v1066
	if v1069 != v1016&int32(2147483644) {
		v1039 = v1065
		v1041 = v1067
		v1044 = v1069
		goto L172
	} else {
		goto L174
	}
L173:
	;
	v1077 = v1065
	v1079 = v1067
	goto L171
L174:
	;
	goto L173
L175:
	;
	v1101 = v1077
	v1102 = v1025
	v1103 = v1079
	goto L176
L176:
	;
	v1120 = *(*int32)(unsafe.Add(mBase, uint32(v1015+v1103<<(uint(int32(2))%32))))
	v1121 = v1120 ^ v1101
	v1122 = int32(1)
	v1125 = v1102 + v1122
	if v1125 != v1024 {
		v1101 = v1121
		v1102 = v1125
		v1103 = v1103 + v1122
		goto L176
	} else {
		goto L178
	}
L177:
	;
	v1133 = v1121
	goto L162
L178:
	;
	goto L177
L179:
	;
	if v1014 != 0 {
		goto L299
	} else {
		goto L300
	}
L180:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v993)+4))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v993)))
	if v1280 < v1281 {
		goto L213
	} else {
		goto L214
	}
L181:
	;
	v1154 = v1016 << (uint(int32(2)) % 32)
	v1165 = v1149
	v1167 = v1150
	goto L184
L182:
	;
	goto L183
L183:
	;
	if v1150 != 0 {
		v1723 = v1149
		goto L179
	} else {
		goto L210
	}
L184:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, uint32(v1165)+4))
	if v1133 == v1179 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	goto L180
L186:
	;
	if v1016 == int32(1) {
		v1723 = v1165
		goto L179
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v1248 = int32(1)
	if v1248 < v1167 {
		v1165 = v1165 + int32(32)
		v1167 = v1167 - v1248
		goto L184
	} else {
		goto L209
	}
L189:
	;
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v1165)))
	if base.Ui32(int32(4)) <= base.Ui32(v1154) {
		goto L193
	} else {
		goto L194
	}
L190:
	;
	if v1243 == int32(0) {
		v1723 = v1165
		goto L179
	} else {
		goto L208
	}
L191:
	;
	v1243 = int32(0)
	goto L190
L192:
	;
	v1217 = v1212
	v1218 = v1213
	v1219 = v1214
	goto L202
L193:
	;
	if (v1015|v1181)&int32(3) != 0 {
		v1212 = v1015
		v1213 = v1181
		v1214 = v1154
		goto L192
	} else {
		goto L196
	}
L194:
	;
	v1205 = v1015
	v1206 = v1181
	v1207 = v1154
	goto L195
L195:
	;
	if v1207 == int32(0) {
		goto L191
	} else {
		goto L201
	}
L196:
	;
	v1189 = v1015
	v1190 = v1181
	v1191 = v1154
	goto L197
L197:
	;
	v1194 = *(*int32)(unsafe.Add(mBase, uint32(v1189)))
	v1195 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	if v1194 != v1195 {
		v1212 = v1189
		v1213 = v1190
		v1214 = v1191
		goto L192
	} else {
		goto L199
	}
L198:
	;
	v1205 = v1200
	v1206 = v1198
	v1207 = v1202
	goto L195
L199:
	;
	v1197 = int32(4)
	v1198 = v1190 + v1197
	v1200 = v1189 + v1197
	v1202 = v1191 - v1197
	if base.Ui32(int32(3)) < base.Ui32(v1202) {
		v1189 = v1200
		v1190 = v1198
		v1191 = v1202
		goto L197
	} else {
		goto L200
	}
L200:
	;
	goto L198
L201:
	;
	v1212 = v1205
	v1213 = v1206
	v1214 = v1207
	goto L192
L202:
	;
	v1222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1217))))
	v1223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1218))))
	if v1222 == v1223 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v1243 = v1222 - v1223
	goto L190
L204:
	;
	v1225 = int32(1)
	v1230 = v1219 - v1225
	if v1230 != 0 {
		v1217 = v1217 + v1225
		v1218 = v1218 + v1225
		v1219 = v1230
		goto L202
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	goto L203
L207:
	;
	goto L191
L208:
	;
	goto L188
L209:
	;
	goto L185
L210:
	;
	goto L180
L211:
	;
	if v1641 == int32(0) {
		v1768 = v1010
		goto L1
	} else {
		goto L291
	}
L212:
	;
	v1456 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+12))
	if v1456 != 0 {
		goto L252
	} else {
		goto L253
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v993)+4)) = v1280 + int32(1)
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v993)+24))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v993)+16))
	v1288 = int32(0)
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(v993)+20))
	v1292 = v1289 + v1280<<(uint(int32(5))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1292)+16)) = uint16(v1288)
	*(*int64)(unsafe.Add(mBase, uint32(v1292)+8)) = int64(0)
	v1298 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1292))) = v1286 + v1280*v1287<<(uint(v1298)%32)
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v993)+32))
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v993)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+24)) = v1302 + v1303*v1280<<(uint(v1298)%32)
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v993)+36))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(v993)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1292)+28)) = v1309 + v1310*v1280<<(uint(int32(3))%32)
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v993)+12))
	if v1316 <= v1288 {
		v1451 = v1292
		goto L212
	} else {
		goto L216
	}
L214:
	;
	goto L215
L215:
	;
	v1348 = base.I32_div_s(v1281<<(uint(int32(1))%32), int32(3))
	v1349 = int32(2)
	if v1348 < (v996-v997)>>(uint(v1349)%32) {
		goto L220
	} else {
		goto L221
	}
L216:
	;
	v1322 = v1288
	goto L217
L217:
	;
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+24))
	v1333 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1329+v1322<<(uint(int32(2))%32)))) = v1333
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1335+v1322<<(uint(int32(3))%32)))) = v1333
	v1342 = v1322 + int32(1)
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v993)+12))
	if v1342 < v1343 {
		v1322 = v1342
		goto L217
	} else {
		goto L219
	}
L218:
	;
	v1451 = v1292
	goto L212
L219:
	;
	goto L218
L220:
	;
	v1356 = v996 - v1348<<(uint(v1349)%32)
	goto L222
L221:
	;
	v1356 = v997
	goto L222
L222:
	;
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v993)+56))
	v1358 = *(*int32)(unsafe.Add(mBase, uint32(v993)+20))
	v1361 = v1358 + v1281<<(uint(int32(5))%32)
	if base.Ui32(v1357) < base.Ui32(v1361) {
		goto L224
	} else {
		goto L225
	}
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v993)+56)) = v1438 + int32(32)
	v1451 = v1438
	goto L212
L224:
	;
	v1368 = v1357
	goto L227
L225:
	;
	goto L226
L226:
	;
	if base.Ui32(v1358) < base.Ui32(v1357) {
		goto L237
	} else {
		goto L238
	}
L227:
	;
	v1373 = *(*int32)(unsafe.Add(mBase, uint32(v1368)+20))
	if base.Ui32(v1356) <= base.Ui32(v1373) {
		goto L229
	} else {
		goto L230
	}
L228:
	;
	goto L226
L229:
	;
	v1376 = v1373
	goto L231
L230:
	;
	v1376 = int32(0)
	goto L231
L231:
	;
	if v1376 == int32(0) {
		goto L232
	} else {
		goto L233
	}
L232:
	;
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1368)+8)))
	if v1379&int32(4) == int32(0) {
		v1438 = v1368
		goto L223
	} else {
		goto L235
	}
L233:
	;
	goto L234
L234:
	;
	v1385 = v1368 + int32(32)
	if base.Ui32(v1385) < base.Ui32(v1361) {
		v1368 = v1385
		goto L227
	} else {
		goto L236
	}
L235:
	;
	goto L234
L236:
	;
	goto L228
L237:
	;
	v1401 = v1358
	goto L240
L238:
	;
	goto L239
L239:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(v992)+36))
	if v1428 != 0 {
		goto L249
	} else {
		goto L250
	}
L240:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1401)+20))
	if base.Ui32(v1356) <= base.Ui32(v1408) {
		goto L243
	} else {
		goto L244
	}
L241:
	;
	goto L239
L242:
	;
	v1416 = v1401 + int32(32)
	if base.Ui32(v1416) < base.Ui32(v1357) {
		v1401 = v1416
		goto L240
	} else {
		goto L248
	}
L243:
	;
	v1411 = v1408
	goto L245
L244:
	;
	v1411 = int32(0)
	goto L245
L245:
	;
	if v1411 != 0 {
		goto L242
	} else {
		goto L246
	}
L246:
	;
	v1412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1401)+8)))
	if v1412&int32(4) != 0 {
		goto L242
	} else {
		goto L247
	}
L247:
	;
	v1438 = v1401
	goto L223
L248:
	;
	goto L241
L249:
	;
	v1430 = v1428
	goto L251
L250:
	;
	v1430 = int32(15)
	goto L251
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v992)+36)) = v1430
	v1641 = int32(0)
	goto L211
L252:
	;
	v1457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1451)+16)))
	v1461 = v1456
	v1462 = v1457
	goto L255
L253:
	;
	goto L254
L254:
	;
	v1493 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1451)+12)) = v1493
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v993)+12))
	if v1493 < v1495 {
		goto L258
	} else {
		goto L259
	}
L255:
	;
	v1468 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+24))
	v1469 = base.I32_extend16_s(v1462)
	v1473 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1468+v1469<<(uint(int32(2))%32)))) = v1473
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1461)+28))
	v1478 = v1475 + v1469<<(uint(int32(3))%32)
	v1479 = *(*int32)(unsafe.Add(mBase, uint32(v1478)))
	*(*int32)(unsafe.Add(mBase, uint32(v1478))) = v1473
	v1482 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478)+4)))
	if v1479 != 0 {
		v1461 = v1479
		v1462 = v1482
		goto L255
	} else {
		goto L257
	}
L256:
	;
	goto L254
L257:
	;
	goto L256
L258:
	;
	v1503 = v1495
	v1505 = int32(0)
	goto L261
L259:
	;
	goto L260
L260:
	;
	v1602 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+8))
	if v1602&int32(2) == int32(0) {
		v1616 = v1602
		goto L277
	} else {
		goto L278
	}
L261:
	;
	v1510 = v1505 << (uint(int32(2)) % 32)
	v1511 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+24))
	v1513 = *(*int32)(unsafe.Add(mBase, uint32(v1510+v1511)))
	if v1513 != 0 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	goto L260
L263:
	;
	v1514 = *(*int32)(unsafe.Add(mBase, uint32(v1513)+12))
	if v1514 != v1451 {
		goto L267
	} else {
		goto L268
	}
L264:
	;
	v1583 = v1503
	goto L265
L265:
	;
	v1590 = v1505 + int32(1)
	if v1590 < v1583 {
		v1503 = v1583
		v1505 = v1590
		goto L261
	} else {
		goto L276
	}
L266:
	;
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+24))
	v1570 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1568+v1510))) = v1570
	v1572 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1572+v1505<<(uint(int32(3))%32)))) = v1570
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v993)+12))
	v1583 = v1578
	goto L265
L267:
	;
	v1524 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1513)+16)))
	v1528 = v1524
	v1529 = v1514
	goto L271
L268:
	;
	v1516 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1513)+16)))
	if v1505 != v1516 {
		goto L267
	} else {
		goto L269
	}
L269:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+28))
	v1522 = *(*int64)(unsafe.Add(mBase, uint32(v1518+v1505<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1513)+12)) = v1522
	goto L266
L270:
	;
	v1549 = int32(3)
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+28))
	v1556 = *(*int64)(unsafe.Add(mBase, uint32(v1552+v1505<<(uint(v1549)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v1548+v1547<<(uint(v1549)%32)))) = v1556
	goto L266
L271:
	;
	v1535 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+28))
	v1538 = v1535 + v1528<<(uint(int32(3))%32)
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v1538)))
	if v1539 == int32(0) {
		v1547 = v1528
		v1548 = v1535
		goto L270
	} else {
		goto L273
	}
L272:
	;
	v1545 = *(*int32)(unsafe.Add(mBase, uint32(v1529)+28))
	v1547 = base.I32_extend16_s(v1528)
	v1548 = v1545
	goto L270
L273:
	;
	v1542 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1538)+4)))
	if v1539 != v1451 {
		v1528 = v1542
		v1529 = v1539
		goto L271
	} else {
		goto L274
	}
L274:
	;
	if v1542 != v1505 {
		v1528 = v1542
		v1529 = v1539
		goto L271
	} else {
		goto L275
	}
L275:
	;
	goto L272
L276:
	;
	goto L262
L277:
	;
	if v1616&int32(8) == int32(0) {
		goto L284
	} else {
		goto L285
	}
L278:
	;
	v1607 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+20))
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v993)+48))
	if v1607 == v1608 {
		v1616 = v1602
		goto L277
	} else {
		goto L279
	}
L279:
	;
	if base.Ui32(v1607) <= base.Ui32(v1608) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1612 = v1608
	goto L282
L281:
	;
	v1612 = int32(0)
	goto L282
L282:
	;
	if v1612 != 0 {
		v1616 = v1602
		goto L277
	} else {
		goto L283
	}
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v993)+48)) = v1607
	v1614 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+8))
	v1616 = v1614
	goto L277
L284:
	;
	v1641 = v1451
	goto L211
L285:
	;
	v1622 = *(*int32)(unsafe.Add(mBase, uint32(v1451)+20))
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v993)+52))
	if v1622 == v1623 {
		goto L284
	} else {
		goto L286
	}
L286:
	;
	if base.Ui32(v1622) <= base.Ui32(v1623) {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1627 = v1623
	goto L289
L288:
	;
	v1627 = int32(0)
	goto L289
L289:
	;
	if v1627 != 0 {
		goto L284
	} else {
		goto L290
	}
L290:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v993)+52)) = v1622
	goto L284
L291:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v993)+16))
	if int32(0) < v1644 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1658 = int32(0)
	goto L295
L293:
	;
	goto L294
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1641)+4)) = v1133
	v1705 = int32(0)
	v1708 = base.B2i32(v1009 != v1705) << (uint(int32(1)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v1641)+8)) = v1708
	if v1008 == v1705 {
		v1723 = v1641
		goto L179
	} else {
		goto L298
	}
L295:
	;
	v1671 = v1658 << (uint(int32(2)) % 32)
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v1641)))
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v993)+28))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1674+v1671)))
	*(*int32)(unsafe.Add(mBase, uint32(v1671+v1672))) = v1676
	v1679 = v1658 + int32(1)
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v993)+16))
	if v1679 < v1680 {
		v1658 = v1679
		goto L295
	} else {
		goto L297
	}
L296:
	;
	goto L294
L297:
	;
	goto L296
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1641)+8)) = v1708 | int32(8)
	v1723 = v1641
	goto L179
L299:
	;
	v1737 = *(*int32)(unsafe.Add(mBase, uint32(v994)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1737+v995<<(uint(int32(2))%32)))) = v1723
	v1742 = *(*int32)(unsafe.Add(mBase, uint32(v994)+28))
	v1746 = *(*int64)(unsafe.Add(mBase, uint32(v1723)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v1742+v995<<(uint(int32(3))%32)))) = v1746
	*(*uint16)(unsafe.Add(mBase, uint32(v1723)+16)) = uint16(v995)
	*(*int32)(unsafe.Add(mBase, uint32(v1723)+12)) = v994
	goto L301
L300:
	;
	goto L301
L301:
	;
	v1768 = v1723
	goto L1
}
func F_missing_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v5 != v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if v6 < v5 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v5) {
		goto L10
	} else {
		goto L11
	}
L4:
	;
	v11 = int32(1)
	goto L6
L5:
	;
	v11 = int32(-1)
	goto L6
L6:
	;
	return v11
L7:
	;
	return v76
L8:
	;
	v76 = int32(0)
	goto L7
L9:
	;
	v50 = v45
	v51 = v46
	v52 = v47
	goto L19
L10:
	;
	if (v13|v14)&int32(3) != 0 {
		v45 = v13
		v46 = v14
		v47 = v5
		goto L9
	} else {
		goto L13
	}
L11:
	;
	v38 = v13
	v39 = v14
	v40 = v5
	goto L12
L12:
	;
	if v40 == int32(0) {
		goto L8
	} else {
		goto L18
	}
L13:
	;
	v22 = v13
	v23 = v14
	v24 = v5
	goto L14
L14:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v27 != v28 {
		v45 = v22
		v46 = v23
		v47 = v24
		goto L9
	} else {
		goto L16
	}
L15:
	;
	v38 = v33
	v39 = v31
	v40 = v35
	goto L12
L16:
	;
	v30 = int32(4)
	v31 = v23 + v30
	v33 = v22 + v30
	v35 = v24 - v30
	if base.Ui32(int32(3)) < base.Ui32(v35) {
		v22 = v33
		v23 = v31
		v24 = v35
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v45 = v38
	v46 = v39
	v47 = v40
	goto L9
L19:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 == v56 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v76 = v55 - v56
	goto L7
L21:
	;
	v58 = int32(1)
	v63 = v52 - v58
	if v63 != 0 {
		v50 = v50 + v58
		v51 = v51 + v58
		v52 = v63
		goto L19
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L20
L24:
	;
	goto L8
}
func F_mix_decrypt_normal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l2 <= int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2 + v9
		return l2
	} else {
		v20 = l1
		v22 = l3
		v23 = v9
		for {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			*(*uint8)(unsafe.Add(mBase, uint32(v23+(l0+int32(84))))) = uint8(v28)
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+(l0+int32(52))))))
			v32 = v28 ^ v31
			*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v32)
			v34 = int32(1)
			v39 = v23 + v34
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v41 = v40 + l2
			if v39 < v41 {
				v20 = v20 + v34
				v22 = v22 + v34
				v23 = v39
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
		return l2
	}
}
func F_mix_encrypt_normal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l2 <= int32(0) {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = l2 + v9
		return l2
	} else {
		v20 = l1
		v22 = l3
		v23 = v9
		for {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
			v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+(l0+int32(52))))))
			v31 = v28 ^ v30
			*(*uint8)(unsafe.Add(mBase, uint32(v23+(l0+int32(84))))) = uint8(v31)
			*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v31)
			v34 = int32(1)
			v39 = v23 + v34
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			v41 = v40 + l2
			if v39 < v41 {
				v20 = v20 + v34
				v22 = v22 + v34
				v23 = v39
				continue
			} else {
				break
			}
			break
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v41
		return l2
	}
}
func F_mul_var(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var __phi84 int32
	_ = __phi84
	var v88 int32
	_ = v88
	var __phi88 int32
	_ = __phi88
	var v92 int32
	_ = v92
	var __phi92 int32
	_ = __phi92
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v771 int32
	_ = v771
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v826 int32
	_ = v826
	var v827 int32
	_ = v827
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v859 int32
	_ = v859
	var v892 int32
	_ = v892
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1002 int32
	_ = v1002
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1048 int32
	_ = v1048
	var v1051 int32
	_ = v1051
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1192 int32
	_ = v1192
	var v1194 int32
	_ = v1194
	var v1195 int32
	_ = v1195
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1238 int32
	_ = v1238
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1276 int32
	_ = v1276
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1365 int32
	_ = v1365
	var v1367 int32
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1370 int32
	_ = v1370
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 int32
	_ = v1381
	var v1383 int32
	_ = v1383
	var v1385 int32
	_ = v1385
	var v1389 int64
	_ = v1389
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1404 int32
	_ = v1404
	var v1405 int32
	_ = v1405
	var v1406 int32
	_ = v1406
	var v1408 int32
	_ = v1408
	var v1412 int32
	_ = v1412
	var v1418 int32
	_ = v1418
	var v1424 int32
	_ = v1424
	var v1449 int32
	_ = v1449
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1480 int32
	_ = v1480
	var v1510 int32
	_ = v1510
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1529 int32
	_ = v1529
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1556 int32
	_ = v1556
	var v1558 int32
	_ = v1558
	var v1561 int32
	_ = v1561
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1575 int64
	_ = v1575
	var v1577 int64
	_ = v1577
	var v1578 int32
	_ = v1578
	var v1581 int32
	_ = v1581
	var v1586 int64
	_ = v1586
	var v1588 int64
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1594 int32
	_ = v1594
	var v1596 int64
	_ = v1596
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1604 int32
	_ = v1604
	var v1605 int32
	_ = v1605
	var v1612 int32
	_ = v1612
	var v1615 int32
	_ = v1615
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1648 int64
	_ = v1648
	var v1652 int32
	_ = v1652
	var v1659 int64
	_ = v1659
	var v1663 int32
	_ = v1663
	var v1670 int64
	_ = v1670
	var v1674 int32
	_ = v1674
	var v1681 int64
	_ = v1681
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1687 int32
	_ = v1687
	var v1693 int32
	_ = v1693
	var v1724 int32
	_ = v1724
	var v1727 int32
	_ = v1727
	var v1760 int64
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1766 int32
	_ = v1766
	var v1800 int32
	_ = v1800
	var v1810 int32
	_ = v1810
	var v1822 int32
	_ = v1822
	var v1833 int32
	_ = v1833
	var v1845 int64
	_ = v1845
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1859 int64
	_ = v1859
	var v1860 int64
	_ = v1860
	var v1866 int64
	_ = v1866
	var v1873 int32
	_ = v1873
	var v1875 int32
	_ = v1875
	var v1898 int64
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1903 int64
	_ = v1903
	var v1904 int64
	_ = v1904
	var v1909 int64
	_ = v1909
	var v1913 int64
	_ = v1913
	var v1914 int64
	_ = v1914
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1922 int64
	_ = v1922
	var v1923 int64
	_ = v1923
	var v1927 int64
	_ = v1927
	var v1931 int64
	_ = v1931
	var v1932 int64
	_ = v1932
	var v1935 int32
	_ = v1935
	var v1941 int32
	_ = v1941
	var v1966 int64
	_ = v1966
	var v1972 int32
	_ = v1972
	var v1973 int64
	_ = v1973
	var v1974 int64
	_ = v1974
	var v1976 int64
	_ = v1976
	var v1979 int64
	_ = v1979
	var v2042 int64
	_ = v2042
	var v2046 int32
	_ = v2046
	var v2048 int32
	_ = v2048
	var v2053 int32
	_ = v2053
	var v2055 int32
	_ = v2055
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2060 int32
	_ = v2060
	var v2067 int32
	_ = v2067
	var v2070 int32
	_ = v2070
	var v2097 int32
	_ = v2097
	var v2099 int32
	_ = v2099
	var v2100 int64
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2104 int64
	_ = v2104
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2113 int64
	_ = v2113
	var v2117 int64
	_ = v2117
	var v2122 int32
	_ = v2122
	var v2124 int32
	_ = v2124
	var v2130 int32
	_ = v2130
	var v2161 int32
	_ = v2161
	var v2162 int64
	_ = v2162
	var v2166 int64
	_ = v2166
	var v2197 int64
	_ = v2197
	var v2201 int32
	_ = v2201
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2268 int32
	_ = v2268
	var v2290 int64
	_ = v2290
	var v2295 int32
	_ = v2295
	var v2299 int64
	_ = v2299
	var v2300 int64
	_ = v2300
	var v2304 int32
	_ = v2304
	var v2308 int64
	_ = v2308
	var v2312 int64
	_ = v2312
	var v2313 int64
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2320 int32
	_ = v2320
	var v2356 int32
	_ = v2356
	var v2368 int32
	_ = v2368
	var v2377 int32
	_ = v2377
	var v2379 int32
	_ = v2379
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2394 int32
	_ = v2394
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2401 int32
	_ = v2401
	var v2402 int32
	_ = v2402
	var v2403 int32
	_ = v2403
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2420 int32
	_ = v2420
	var v2424 int32
	_ = v2424
	var v2430 int32
	_ = v2430
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2440 int32
	_ = v2440
	var v2443 int32
	_ = v2443
	var v2445 int32
	_ = v2445
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2450 int32
	_ = v2450
	var v2458 int32
	_ = v2458
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2466 int32
	_ = v2466
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2495 int32
	_ = v2495
	var v2498 int32
	_ = v2498
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2539 int32
	_ = v2539
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2581 int32
	_ = v2581
	var v2612 int32
	_ = v2612
	var v2615 int32
	_ = v2615
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v32 < v33 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v35 = v32
	goto L3
L2:
	;
	v35 = v33
	goto L3
L3:
	;
	if v35 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v38 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v48 = base.B2i32(v33 < v32)
	if v33 < v32 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	F_pfree(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v44 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v44
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v44
	return
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	v49 = v32
	goto L14
L13:
	;
	v49 = v33
	goto L14
L14:
	;
	if v33 < v32 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v50 = l0
	goto L17
L16:
	;
	v50 = l1
	goto L17
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
	if v33 < v32 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v52 = l1
	goto L20
L19:
	;
	v52 = l0
	goto L20
L20:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+20))
	if int32(6) < v35 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v1337 = v33 + v32
	v1338 = int32(2)
	v1339 = base.I32_div_s(v1337, v1338)
	v1340 = int32(1)
	v1341 = v1339 + v1340
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v1356 = v1342 + (v1343 + ((v35^int32(-1))&v1340 - (v1337 + v49&v1340) + v1341<<(uint(v1340)%32)))
	v1357 = int32(3)
	v1360 = base.I32_div_s(l3+v1357, int32(4))
	v1365 = base.I32_div_s(v1356+v1360+v1357, v1338)
	v1367 = v1365 + v1340
	if v1341 < v1367 {
		goto L83
	} else {
		goto L84
	}
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	if l3 != v56+v57 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v64 = int32(0)
	v65 = v33 + v32
	v67 = v65 << (uint(int32(1)) % 32)
	v70 = F_palloc(m, v67+int32(2))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L24
	}
L24:
	;
	v72 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v70))) = uint16(v72)
	v76 = v70 + int32(2)
	switch v35 - int32(1) {
	case 0:
		goto L37
	case 1:
		goto L36
	case 2:
		goto L35
	case 3:
		goto L34
	case 4:
		goto L33
	case 5:
		goto L32
	default:
		v859 = v64
		goto L31
	}
L25:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v1170 != 0 {
		goto L64
	} else {
		goto L65
	}
L26:
	;
	v1126 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51))))
	v1127 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v1129 = v1126*v1127 + v1096
	v1130 = int32(10000)
	v1131 = base.I32_div_u_s(v1129, v1130)
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+2)) = uint16(v1131)
	v1135 = v1129 - v1131*v1130
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+4)) = uint16(v1135)
	goto L25
L27:
	;
	v1081 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+2)))
	v1082 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v1085 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51))))
	v1086 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v1088 = v1081*v1082 + v1051 + v1085*v1086
	v1089 = int32(10000)
	v1090 = base.I32_div_u_s(v1088, v1089)
	v1093 = v1088 - v1090*v1089
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+6)) = uint16(v1093)
	v1096 = v1090
	goto L26
L28:
	;
	v1032 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+4)))
	v1033 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v1036 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+2)))
	v1037 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v1040 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51))))
	v1041 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+4)))
	v1043 = v1032*v1033 + v1002 + v1036*v1037 + v1040*v1041
	v1044 = int32(10000)
	v1045 = base.I32_div_u_s(v1043, v1044)
	v1048 = v1043 - v1045*v1044
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+8)) = uint16(v1048)
	v1051 = v1045
	goto L27
L29:
	;
	v979 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+6)))
	v980 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v983 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+4)))
	v984 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v987 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+2)))
	v988 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+4)))
	v991 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51))))
	v992 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+6)))
	v994 = v979*v980 + v949 + v983*v984 + v987*v988 + v991*v992
	v995 = int32(10000)
	v996 = base.I32_div_u_s(v994, v995)
	v999 = v994 - v996*v995
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+10)) = uint16(v999)
	v1002 = v996
	goto L28
L30:
	;
	v922 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+8)))
	v923 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v926 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+6)))
	v927 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v930 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+4)))
	v931 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+4)))
	v934 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+2)))
	v935 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+6)))
	v938 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51))))
	v939 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+8)))
	v941 = v922*v923 + v892 + v926*v927 + v930*v931 + v934*v935 + v938*v939
	v942 = int32(10000)
	v943 = base.I32_div_u_s(v941, v942)
	v946 = v941 - v943*v942
	*(*uint16)(unsafe.Add(mBase, uint32(v70)+12)) = uint16(v946)
	v949 = v943
	goto L29
L31:
	;
	switch v35 - int32(2) {
	case 0:
		v1096 = v859
		goto L26
	case 1:
		v1051 = v859
		goto L27
	case 2:
		v1002 = v859
		goto L28
	case 3:
		v949 = v859
		goto L29
	case 4:
		v892 = v859
		goto L30
	default:
		goto L25
	}
L32:
	;
	v651 = int32(1)
	v652 = v49 - v651
	v655 = v51 + v652<<(uint(v651)%32)
	v656 = int32(*(*int16)(unsafe.Add(mBase, uint32(v655))))
	v657 = int32(10)
	v658 = v53 + v657
	v659 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
	v660 = v656 * v659
	v661 = int32(10000)
	v662 = base.I32_div_u_s(v660, v661)
	v665 = v660 - v662*v661
	*(*uint16)(unsafe.Add(mBase, uint32(v70+v67))) = uint16(v665)
	v667 = v76 + v67
	v668 = int32(4)
	v670 = int32(*(*int16)(unsafe.Add(mBase, uint32(v655))))
	v671 = int32(8)
	v672 = v53 + v671
	v673 = int32(*(*int16)(unsafe.Add(mBase, uint32(v672))))
	v678 = v51 + v49<<(uint(v651)%32)
	v680 = v678 - v668
	v681 = int32(*(*int16)(unsafe.Add(mBase, uint32(v680))))
	v682 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
	v684 = v670*v673 + v662 + v681*v682
	v686 = base.I32_div_u_s(v684, v661)
	v689 = v684 - v686*v661
	*(*uint16)(unsafe.Add(mBase, uint32(v667-v668))) = uint16(v689)
	v691 = int32(6)
	v693 = int32(*(*int16)(unsafe.Add(mBase, uint32(v655))))
	v695 = v53 + v691
	v696 = int32(*(*int16)(unsafe.Add(mBase, uint32(v695))))
	v699 = int32(*(*int16)(unsafe.Add(mBase, uint32(v680))))
	v700 = int32(*(*int16)(unsafe.Add(mBase, uint32(v672))))
	v704 = v678 - v691
	v705 = int32(*(*int16)(unsafe.Add(mBase, uint32(v704))))
	v706 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
	v708 = v693*v696 + v686 + v699*v700 + v705*v706
	v710 = base.I32_div_u_s(v708, v661)
	v713 = v708 - v710*v661
	*(*uint16)(unsafe.Add(mBase, uint32(v667-v691))) = uint16(v713)
	v717 = int32(*(*int16)(unsafe.Add(mBase, uint32(v655))))
	v719 = v53 + v668
	v720 = int32(*(*int16)(unsafe.Add(mBase, uint32(v719))))
	v723 = int32(*(*int16)(unsafe.Add(mBase, uint32(v680))))
	v724 = int32(*(*int16)(unsafe.Add(mBase, uint32(v695))))
	v727 = int32(*(*int16)(unsafe.Add(mBase, uint32(v704))))
	v728 = int32(*(*int16)(unsafe.Add(mBase, uint32(v672))))
	v732 = v678 - v671
	v733 = int32(*(*int16)(unsafe.Add(mBase, uint32(v732))))
	v734 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
	v736 = v717*v720 + v710 + v723*v724 + v727*v728 + v733*v734
	v738 = base.I32_div_u_s(v736, v661)
	v741 = v736 - v738*v661
	*(*uint16)(unsafe.Add(mBase, uint32(v667-v671))) = uint16(v741)
	v745 = int32(*(*int16)(unsafe.Add(mBase, uint32(v655))))
	v746 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v749 = int32(*(*int16)(unsafe.Add(mBase, uint32(v680))))
	v750 = int32(*(*int16)(unsafe.Add(mBase, uint32(v719))))
	v753 = int32(*(*int16)(unsafe.Add(mBase, uint32(v704))))
	v754 = int32(*(*int16)(unsafe.Add(mBase, uint32(v695))))
	v757 = int32(*(*int16)(unsafe.Add(mBase, uint32(v732))))
	v758 = int32(*(*int16)(unsafe.Add(mBase, uint32(v672))))
	v763 = int32(*(*int16)(unsafe.Add(mBase, uint32(v678-v657))))
	v764 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
	v766 = v745*v746 + v738 + v749*v750 + v753*v754 + v757*v758 + v763*v764
	v768 = base.I32_div_u_s(v766, v661)
	v771 = v766 - v768*v661
	*(*uint16)(unsafe.Add(mBase, uint32(v667-v657))) = uint16(v771)
	if v652 < int32(5) {
		v892 = v768
		goto L30
	} else {
		goto L60
	}
L33:
	;
	v481 = int32(1)
	v482 = v49 - v481
	v485 = v51 + v482<<(uint(v481)%32)
	v486 = int32(*(*int16)(unsafe.Add(mBase, uint32(v485))))
	v487 = int32(8)
	v488 = v53 + v487
	v489 = int32(*(*int16)(unsafe.Add(mBase, uint32(v488))))
	v490 = v486 * v489
	v491 = int32(10000)
	v492 = base.I32_div_u_s(v490, v491)
	v495 = v490 - v492*v491
	*(*uint16)(unsafe.Add(mBase, uint32(v70+v67))) = uint16(v495)
	v497 = v76 + v67
	v498 = int32(4)
	v500 = int32(*(*int16)(unsafe.Add(mBase, uint32(v485))))
	v501 = int32(6)
	v502 = v53 + v501
	v503 = int32(*(*int16)(unsafe.Add(mBase, uint32(v502))))
	v508 = v51 + v49<<(uint(v481)%32)
	v510 = v508 - v498
	v511 = int32(*(*int16)(unsafe.Add(mBase, uint32(v510))))
	v512 = int32(*(*int16)(unsafe.Add(mBase, uint32(v488))))
	v514 = v500*v503 + v492 + v511*v512
	v516 = base.I32_div_u_s(v514, v491)
	v519 = v514 - v516*v491
	*(*uint16)(unsafe.Add(mBase, uint32(v497-v498))) = uint16(v519)
	v523 = int32(*(*int16)(unsafe.Add(mBase, uint32(v485))))
	v525 = v53 + v498
	v526 = int32(*(*int16)(unsafe.Add(mBase, uint32(v525))))
	v529 = int32(*(*int16)(unsafe.Add(mBase, uint32(v510))))
	v530 = int32(*(*int16)(unsafe.Add(mBase, uint32(v502))))
	v534 = v508 - v501
	v535 = int32(*(*int16)(unsafe.Add(mBase, uint32(v534))))
	v536 = int32(*(*int16)(unsafe.Add(mBase, uint32(v488))))
	v538 = v523*v526 + v516 + v529*v530 + v535*v536
	v540 = base.I32_div_u_s(v538, v491)
	v543 = v538 - v540*v491
	*(*uint16)(unsafe.Add(mBase, uint32(v497-v501))) = uint16(v543)
	v547 = int32(*(*int16)(unsafe.Add(mBase, uint32(v485))))
	v548 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v551 = int32(*(*int16)(unsafe.Add(mBase, uint32(v510))))
	v552 = int32(*(*int16)(unsafe.Add(mBase, uint32(v525))))
	v555 = int32(*(*int16)(unsafe.Add(mBase, uint32(v534))))
	v556 = int32(*(*int16)(unsafe.Add(mBase, uint32(v502))))
	v561 = int32(*(*int16)(unsafe.Add(mBase, uint32(v508-v487))))
	v562 = int32(*(*int16)(unsafe.Add(mBase, uint32(v488))))
	v564 = v547*v548 + v540 + v551*v552 + v555*v556 + v561*v562
	v566 = base.I32_div_u_s(v564, v491)
	v569 = v564 - v566*v491
	*(*uint16)(unsafe.Add(mBase, uint32(v497-v487))) = uint16(v569)
	if v482 < v498 {
		v949 = v566
		goto L29
	} else {
		goto L56
	}
L34:
	;
	v345 = int32(1)
	v346 = v49 - v345
	v349 = v51 + v346<<(uint(v345)%32)
	v350 = int32(*(*int16)(unsafe.Add(mBase, uint32(v349))))
	v351 = int32(6)
	v352 = v53 + v351
	v353 = int32(*(*int16)(unsafe.Add(mBase, uint32(v352))))
	v354 = v350 * v353
	v355 = int32(10000)
	v356 = base.I32_div_u_s(v354, v355)
	v359 = v354 - v356*v355
	*(*uint16)(unsafe.Add(mBase, uint32(v70+v67))) = uint16(v359)
	v361 = v76 + v67
	v362 = int32(4)
	v364 = int32(*(*int16)(unsafe.Add(mBase, uint32(v349))))
	v366 = v53 + v362
	v367 = int32(*(*int16)(unsafe.Add(mBase, uint32(v366))))
	v372 = v51 + v49<<(uint(v345)%32)
	v374 = v372 - v362
	v375 = int32(*(*int16)(unsafe.Add(mBase, uint32(v374))))
	v376 = int32(*(*int16)(unsafe.Add(mBase, uint32(v352))))
	v378 = v364*v367 + v356 + v375*v376
	v380 = base.I32_div_u_s(v378, v355)
	v383 = v378 - v380*v355
	*(*uint16)(unsafe.Add(mBase, uint32(v361-v362))) = uint16(v383)
	v387 = int32(*(*int16)(unsafe.Add(mBase, uint32(v349))))
	v388 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v391 = int32(*(*int16)(unsafe.Add(mBase, uint32(v374))))
	v392 = int32(*(*int16)(unsafe.Add(mBase, uint32(v366))))
	v397 = int32(*(*int16)(unsafe.Add(mBase, uint32(v372-v351))))
	v398 = int32(*(*int16)(unsafe.Add(mBase, uint32(v352))))
	v400 = v387*v388 + v380 + v391*v392 + v397*v398
	v402 = base.I32_div_u_s(v400, v355)
	v405 = v400 - v402*v355
	*(*uint16)(unsafe.Add(mBase, uint32(v361-v351))) = uint16(v405)
	if v346 < int32(3) {
		v1002 = v402
		goto L28
	} else {
		goto L52
	}
L35:
	;
	v241 = int32(1)
	v242 = v49 - v241
	v245 = v51 + v242<<(uint(v241)%32)
	v246 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245))))
	v247 = int32(4)
	v248 = v53 + v247
	v249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248))))
	v250 = v246 * v249
	v251 = int32(10000)
	v252 = base.I32_div_u_s(v250, v251)
	v255 = v250 - v252*v251
	*(*uint16)(unsafe.Add(mBase, uint32(v70+v67))) = uint16(v255)
	v260 = int32(*(*int16)(unsafe.Add(mBase, uint32(v245))))
	v261 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v269 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51+v49<<(uint(v241)%32)-v247))))
	v270 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248))))
	v272 = v260*v261 + v252 + v269*v270
	v274 = base.I32_div_u_s(v272, v251)
	v277 = v272 - v274*v251
	*(*uint16)(unsafe.Add(mBase, uint32(v76+v67-v247))) = uint16(v277)
	if v242 < int32(2) {
		v1051 = v274
		goto L27
	} else {
		goto L48
	}
L36:
	;
	v165 = int32(1)
	v166 = v49 - v165
	v170 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51+v166<<(uint(v165)%32)))))
	v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v172 = v170 * v171
	v173 = int32(10000)
	v174 = base.I32_div_u_s(v172, v173)
	v177 = v172 - v174*v173
	*(*uint16)(unsafe.Add(mBase, uint32(v70+v67))) = uint16(v177)
	if v166 <= int32(0) {
		v1096 = v174
		goto L26
	} else {
		goto L44
	}
L37:
	;
	v80 = v49 - int32(1)
	if int32(0) <= v80 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	__phi84 = v64
	__phi88 = v80
	__phi92 = v49
	v84 = __phi84
	v88 = __phi88
	v92 = __phi92
	goto L41
L39:
	;
	v133 = v64
	goto L40
L40:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v76))) = uint16(v133)
	v859 = v133
	goto L31
L41:
	;
	v114 = int32(1)
	v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51+v88<<(uint(v114)%32)))))
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v123 = v120*v121 + v84
	v124 = int32(10000)
	v125 = base.I32_div_u_s(v123, v124)
	v128 = v123 - v125*v124
	*(*uint16)(unsafe.Add(mBase, uint32(v76+v92<<(uint(v114)%32)))) = uint16(v128)
	if v88 != 0 {
		__phi84 = v125
		__phi88 = v88 - v114
		__phi92 = v88
		v84 = __phi84
		v88 = __phi88
		v92 = __phi92
		goto L41
	} else {
		goto L43
	}
L42:
	;
	v133 = v125
	goto L40
L43:
	;
	goto L42
L44:
	;
	v185 = v166
	v186 = v174
	goto L45
L45:
	;
	v216 = int32(1)
	v217 = v185 << (uint(v216) % 32)
	v219 = v51 + v217
	v220 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219))))
	v221 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v226 = int32(*(*int16)(unsafe.Add(mBase, uint32(v219-int32(2)))))
	v227 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53+int32(2)))))
	v229 = v220*v221 + v186 + v226*v227
	v230 = int32(10000)
	v231 = base.I32_div_u_s(v229, v230)
	v234 = v229 - v231*v230
	*(*uint16)(unsafe.Add(mBase, uint32(v70+int32(4)+v217))) = uint16(v234)
	if base.Ui32(v216) < base.Ui32(v185) {
		v185 = v185 - v216
		v186 = v231
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v859 = v231
	goto L31
L47:
	;
	goto L46
L48:
	;
	v283 = v242
	v284 = v274
	goto L49
L49:
	;
	v314 = int32(1)
	v315 = v283 << (uint(v314) % 32)
	v317 = v51 + v315
	v318 = int32(*(*int16)(unsafe.Add(mBase, uint32(v317))))
	v319 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v322 = int32(2)
	v324 = int32(*(*int16)(unsafe.Add(mBase, uint32(v317-v322))))
	v325 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+2)))
	v330 = int32(*(*int16)(unsafe.Add(mBase, uint32(v317-int32(4)))))
	v331 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248))))
	v333 = v318*v319 + v284 + v324*v325 + v330*v331
	v334 = int32(10000)
	v335 = base.I32_div_u_s(v333, v334)
	v338 = v333 - v335*v334
	*(*uint16)(unsafe.Add(mBase, uint32(v70+int32(4)+v315))) = uint16(v338)
	if base.Ui32(v322) < base.Ui32(v283) {
		v283 = v283 - v314
		v284 = v335
		goto L49
	} else {
		goto L51
	}
L50:
	;
	v859 = v335
	goto L31
L51:
	;
	goto L50
L52:
	;
	v413 = v346
	v414 = v402
	goto L53
L53:
	;
	v444 = int32(1)
	v445 = v413 << (uint(v444) % 32)
	v447 = v51 + v445
	v448 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447))))
	v449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v454 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447-int32(2)))))
	v455 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53+int32(2)))))
	v460 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447-int32(4)))))
	v461 = int32(*(*int16)(unsafe.Add(mBase, uint32(v366))))
	v466 = int32(*(*int16)(unsafe.Add(mBase, uint32(v447-int32(6)))))
	v467 = int32(*(*int16)(unsafe.Add(mBase, uint32(v352))))
	v469 = v448*v449 + v414 + v454*v455 + v460*v461 + v466*v467
	v470 = int32(10000)
	v471 = base.I32_div_u_s(v469, v470)
	v474 = v469 - v471*v470
	*(*uint16)(unsafe.Add(mBase, uint32(v70+int32(4)+v445))) = uint16(v474)
	if base.Ui32(int32(3)) < base.Ui32(v413) {
		v413 = v413 - v444
		v414 = v471
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v859 = v471
	goto L31
L55:
	;
	goto L54
L56:
	;
	v578 = v566
	v582 = v482
	goto L57
L57:
	;
	v608 = int32(1)
	v609 = v582 << (uint(v608) % 32)
	v611 = v609 + v51
	v612 = int32(*(*int16)(unsafe.Add(mBase, uint32(v611))))
	v613 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v618 = int32(*(*int16)(unsafe.Add(mBase, uint32(v611-int32(2)))))
	v619 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53+int32(2)))))
	v622 = int32(4)
	v624 = int32(*(*int16)(unsafe.Add(mBase, uint32(v611-v622))))
	v625 = int32(*(*int16)(unsafe.Add(mBase, uint32(v525))))
	v630 = int32(*(*int16)(unsafe.Add(mBase, uint32(v611-int32(6)))))
	v631 = int32(*(*int16)(unsafe.Add(mBase, uint32(v502))))
	v636 = int32(*(*int16)(unsafe.Add(mBase, uint32(v611-int32(8)))))
	v637 = int32(*(*int16)(unsafe.Add(mBase, uint32(v488))))
	v639 = v612*v613 + v578 + v618*v619 + v624*v625 + v630*v631 + v636*v637
	v640 = int32(10000)
	v641 = base.I32_div_u_s(v639, v640)
	v644 = v639 - v641*v640
	*(*uint16)(unsafe.Add(mBase, uint32(v70+int32(4)+v609))) = uint16(v644)
	if base.Ui32(v622) < base.Ui32(v582) {
		v578 = v641
		v582 = v582 - v608
		goto L57
	} else {
		goto L59
	}
L58:
	;
	v859 = v641
	goto L31
L59:
	;
	goto L58
L60:
	;
	v780 = v768
	v784 = v652
	goto L61
L61:
	;
	v810 = int32(1)
	v811 = v784 << (uint(v810) % 32)
	v813 = v811 + v51
	v814 = int32(*(*int16)(unsafe.Add(mBase, uint32(v813))))
	v815 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
	v820 = int32(*(*int16)(unsafe.Add(mBase, uint32(v813-int32(2)))))
	v821 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53+int32(2)))))
	v826 = int32(*(*int16)(unsafe.Add(mBase, uint32(v813-int32(4)))))
	v827 = int32(*(*int16)(unsafe.Add(mBase, uint32(v719))))
	v832 = int32(*(*int16)(unsafe.Add(mBase, uint32(v813-int32(6)))))
	v833 = int32(*(*int16)(unsafe.Add(mBase, uint32(v695))))
	v838 = int32(*(*int16)(unsafe.Add(mBase, uint32(v813-int32(8)))))
	v839 = int32(*(*int16)(unsafe.Add(mBase, uint32(v672))))
	v844 = int32(*(*int16)(unsafe.Add(mBase, uint32(v813-int32(10)))))
	v845 = int32(*(*int16)(unsafe.Add(mBase, uint32(v658))))
	v847 = v814*v815 + v780 + v820*v821 + v826*v827 + v832*v833 + v838*v839 + v844*v845
	v848 = int32(10000)
	v849 = base.I32_div_u_s(v847, v848)
	v852 = v847 - v849*v848
	*(*uint16)(unsafe.Add(mBase, uint32(v70+int32(4)+v811))) = uint16(v852)
	if base.Ui32(int32(5)) < base.Ui32(v784) {
		v780 = v849
		v784 = v784 - v810
		goto L61
	} else {
		goto L63
	}
L62:
	;
	v859 = v849
	goto L31
L63:
	;
	goto L62
L64:
	;
	F_pfree(m, v1170)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L10
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v70
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v65
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = base.B2i32(v62 != v63) << (uint(int32(14)) % 32)
	v1179 = v61 + v60 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1179
	v1181 = *(*int32)(unsafe.Add(mBase, uint32(v50)+12))
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v1181 + v1182
	if int32(0) < v65 {
		goto L70
	} else {
		goto L71
	}
L67:
	;
	goto L66
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1311
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v1310
	return
L69:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v1310 = v1276
	v1311 = int32(0)
	goto L68
L70:
	;
	v1192 = v1179
	v1194 = v76
	v1195 = v65
	goto L74
L71:
	;
	goto L72
L72:
	;
	if v65 != 0 {
		v1310 = v76
		v1311 = v65
		goto L68
	} else {
		goto L82
	}
L73:
	;
	v1238 = v1195
	goto L78
L74:
	;
	v1219 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1194))))
	if v1219 != 0 {
		goto L73
	} else {
		goto L76
	}
L75:
	;
	v1276 = v76 + v67
	goto L69
L76:
	;
	v1220 = int32(1)
	v1221 = v1192 - v1220
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1221
	if v1220 < v1195 {
		v1192 = v1221
		v1194 = v1194 + int32(2)
		v1195 = v1195 - v1220
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	v1265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1194-int32(2)+v1238<<(uint(int32(1))%32)))))
	if v1265 != 0 {
		v1310 = v1194
		v1311 = v1238
		goto L68
	} else {
		goto L80
	}
L79:
	;
	v1276 = v1194
	goto L69
L80:
	;
	v1266 = int32(1)
	if v1266 < v1238 {
		v1238 = v1238 - v1266
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v1276 = v76
	goto L69
L83:
	;
	v1369 = v1341
	goto L85
L84:
	;
	v1369 = v1367
	goto L85
L85:
	;
	v1370 = int32(1)
	v1372 = int32(2)
	v1373 = base.I32_div_s(v49+v1370, v1372)
	v1377 = base.I32_div_s(v35+v1370, v1372)
	v1378 = v1373 + v1377
	v1379 = v1341 - v1378
	v1381 = v1379 + v1370
	if v1369 <= v1381 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v1383 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v1395 = v1369 - v1381
	v1398 = v1369 << (uint(int32(3)) % 32)
	if v1373 < v1395 {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	F_pfree(m, v1383)
	mBase = m.M
	v1385 = m.ExcPending
	if v1385 != 0 {
		goto L10
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v1389 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v1389
	*(*int64)(unsafe.Add(mBase, uint32(l2)+16)) = v1389
	return
L92:
	;
	goto L91
L93:
	;
	v1400 = v1373
	goto L95
L94:
	;
	v1400 = v1395
	goto L95
L95:
	;
	v1404 = F_palloc(m, v1398+v1400<<(uint(int32(2))%32))
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L10
	} else {
		goto L96
	}
L96:
	;
	v1406 = v1404 + v1398
	v1408 = v1400 - int32(1)
	if v1408 <= int32(0) {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	if v1377 < v1395 {
		goto L108
	} else {
		goto L109
	}
L98:
	;
	v1529 = int32(0)
	goto L97
L99:
	;
	goto L100
L100:
	;
	v1412 = int32(0)
	if v1400 != int32(2) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v1418 = int32(0)
	v1424 = v1412
	goto L104
L102:
	;
	v1480 = v1412
	goto L103
L103:
	;
	if v1408&int32(1) == int32(0) {
		v1529 = v1408
		goto L97
	} else {
		goto L107
	}
L104:
	;
	v1449 = int32(2)
	v1450 = v1424 << (uint(v1449) % 32)
	v1452 = v1450 + v51
	v1453 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1452))))
	v1454 = int32(10000)
	v1456 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1452)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1406+v1450))) = v1453*v1454 + v1456
	v1460 = v1450 | int32(4)
	v1462 = v1460 + v51
	v1463 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1462))))
	v1466 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1462)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1406+v1460))) = v1463*v1454 + v1466
	v1470 = v1424 + v1449
	v1472 = v1418 + v1449
	if v1472 != v1408&int32(2147483646) {
		v1418 = v1472
		v1424 = v1470
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v1480 = v1470
	goto L103
L106:
	;
	goto L105
L107:
	;
	v1510 = v1480 << (uint(int32(2)) % 32)
	v1512 = v51 + v1510
	v1513 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1512))))
	v1516 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1512)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v1406+v1510))) = v1513*int32(10000) + v1516
	v1529 = v1408
	goto L97
L108:
	;
	v1550 = v1377
	goto L110
L109:
	;
	v1550 = v1395
	goto L110
L110:
	;
	v1552 = v1529 << (uint(int32(2)) % 32)
	v1554 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51+v1552))))
	v1556 = v1554 * int32(10000)
	v1558 = int32(1)
	v1561 = v1529<<(uint(v1558)%32) | v1558
	if v1561 < v49 {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v1566 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51+v1561<<(uint(int32(1))%32)))))
	v1568 = v1556 + v1566
	goto L113
L112:
	;
	v1568 = v1556
	goto L113
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1552+v1406))) = v1568
	v1575 = int64(*(*int16)(unsafe.Add(mBase, uint32(v1550<<(uint(int32(2))%32)+v53-int32(4)))))
	v1577 = v1575 * int64(10000)
	v1578 = int32(1)
	v1581 = v1550<<(uint(v1578)%32) - v1578
	if v1581 < v35 {
		goto L114
	} else {
		goto L115
	}
L114:
	;
	v1586 = int64(*(*int16)(unsafe.Add(mBase, uint32(v53+v1581<<(uint(int32(1))%32)))))
	v1588 = v1577 + v1586
	goto L116
L115:
	;
	v1588 = v1577
	goto L116
L116:
	;
	v1590 = v1550 + v1379
	v1592 = v1590 << (uint(int32(3)) % 32)
	v1594 = F__emscripten_memset_bulkmem(m, v1404, base.I32_extend8_s(int32(0)), v1592)
	mBase = m.M
	goto L117
L117:
	;
	v1596 = v1588 & int64(4294967295)
	v1597 = v1369 - v1590
	if v1400 < v1597 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v1800 = v1550 - int32(2)
	if int32(0) <= v1800 {
		goto L133
	} else {
		goto L134
	}
L119:
	;
	v1599 = v1400
	goto L121
L120:
	;
	v1599 = v1597
	goto L121
L121:
	;
	if v1599 <= int32(0) {
		goto L118
	} else {
		goto L122
	}
L122:
	;
	v1602 = v1592 + v1594
	v1604 = v1599 & int32(3)
	v1605 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v1599) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v1612 = int32(0)
	v1615 = v1605
	goto L126
L124:
	;
	v1693 = v1605
	goto L125
L125:
	;
	if v1604 == int32(0) {
		goto L118
	} else {
		goto L129
	}
L126:
	;
	v1642 = int32(3)
	v1645 = int32(2)
	v1648 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1406+v1615<<(uint(v1645)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1602+v1615<<(uint(v1642)%32)))) = v1596 * v1648
	v1652 = v1615 | int32(1)
	v1659 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1406+v1652<<(uint(v1645)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1602+v1652<<(uint(v1642)%32)))) = v1596 * v1659
	v1663 = v1615 | v1645
	v1670 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1406+v1663<<(uint(v1645)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1602+v1663<<(uint(v1642)%32)))) = v1596 * v1670
	v1674 = v1615 | v1642
	v1681 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1406+v1674<<(uint(v1645)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1602+v1674<<(uint(v1642)%32)))) = v1596 * v1681
	v1684 = int32(4)
	v1685 = v1615 + v1684
	v1687 = v1612 + v1684
	if v1687 != v1599&int32(2147483644) {
		v1612 = v1687
		v1615 = v1685
		goto L126
	} else {
		goto L128
	}
L127:
	;
	v1693 = v1685
	goto L125
L128:
	;
	goto L127
L129:
	;
	v1724 = int32(0)
	v1727 = v1693
	goto L130
L130:
	;
	v1760 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1406+v1727<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v1602+v1727<<(uint(int32(3))%32)))) = v1596 * v1760
	v1763 = int32(1)
	v1766 = v1724 + v1763
	if v1766 != v1604 {
		v1724 = v1766
		v1727 = v1727 + v1763
		goto L130
	} else {
		goto L132
	}
L131:
	;
	goto L118
L132:
	;
	goto L131
L133:
	;
	v1810 = v1594 - int32(8)
	v1822 = v1800
	v1833 = v1369 + v1378 - (v1550 + v1339)
	v1845 = v1596
	goto L136
L134:
	;
	goto L135
L135:
	;
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	if v2241 != 0 {
		goto L180
	} else {
		goto L181
	}
L136:
	;
	v1851 = v53 + v1822<<(uint(int32(2))%32)
	v1852 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1851))))
	v1855 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1851)+2)))
	v1856 = v1852*int32(10000) + v1855
	if v1856 == int32(0) {
		v2197 = v1845
		goto L138
	} else {
		goto L139
	}
L137:
	;
	goto L135
L138:
	;
	v2201 = int32(1)
	if int32(0) < v1822 {
		v1822 = v1822 - v2201
		v1833 = v1833 + v2201
		v1845 = v2197
		goto L136
	} else {
		goto L179
	}
L139:
	;
	v1859 = base.I64_extend_i32_u(v1856)
	v1860 = v1845 + v1859
	if base.Ui64(int64(184467440738)) <= base.Ui64(v1860) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	if v1369 <= int32(0) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	v2042 = v1860
	goto L142
L142:
	;
	v2046 = v1369 + (v1379 ^ int32(-1)) - v1822
	if v1400 < v2046 {
		goto L162
	} else {
		goto L163
	}
L143:
	;
	v2042 = base.I64_extend_i32_u(v1856 + int32(1))
	goto L142
L144:
	;
	v1866 = int64(0)
	if v1369 != int32(1) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v1873 = v1369
	v1875 = int32(0)
	v1898 = v1866
	goto L148
L146:
	;
	v1941 = v1369
	v1966 = v1866
	goto L147
L147:
	;
	if v1369&int32(1) == int32(0) {
		goto L143
	} else {
		goto L158
	}
L148:
	;
	v1902 = v1810 + v1873<<(uint(int32(3))%32)
	v1903 = *(*int64)(unsafe.Add(mBase, uint32(v1902)))
	v1904 = v1903 + v1898
	if base.Ui64(v1904) < base.Ui64(int64(100000000)) {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	v1941 = v1918
	v1966 = v1932
	goto L147
L150:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1902))) = v1913
	v1918 = v1873 - int32(2)
	v1921 = v1594 + v1918<<(uint(int32(3))%32)
	v1922 = *(*int64)(unsafe.Add(mBase, uint32(v1921)))
	v1923 = v1922 + v1914
	if base.Ui64(int64(100000000)) <= base.Ui64(v1923) {
		goto L154
	} else {
		goto L155
	}
L151:
	;
	v1913 = v1904
	v1914 = int64(0)
	goto L150
L152:
	;
	goto L153
L153:
	;
	v1909 = base.I64_div_u_s(v1904, int64(100000000))
	v1913 = v1909*int64(-100000000) + v1904
	v1914 = v1909
	goto L150
L154:
	;
	v1927 = base.I64_div_u_s(v1923, int64(100000000))
	v1931 = v1927*int64(-100000000) + v1923
	v1932 = v1927
	goto L156
L155:
	;
	v1931 = v1923
	v1932 = int64(0)
	goto L156
L156:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1921))) = v1931
	v1935 = v1875 + int32(2)
	if v1935 != v1369&int32(2147483646) {
		v1873 = v1918
		v1875 = v1935
		v1898 = v1932
		goto L148
	} else {
		goto L157
	}
L157:
	;
	goto L149
L158:
	;
	v1972 = v1810 + v1941<<(uint(int32(3))%32)
	v1973 = *(*int64)(unsafe.Add(mBase, uint32(v1972)))
	v1974 = v1973 + v1966
	v1976 = base.I64_rem_u_s(v1974, int64(100000000))
	if base.Ui64(int64(99999999)) < base.Ui64(v1974) {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v1979 = v1976
	goto L161
L160:
	;
	v1979 = v1974
	goto L161
L161:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1972))) = v1979
	goto L143
L162:
	;
	v2048 = v1400
	goto L164
L163:
	;
	v2048 = v2046
	goto L164
L164:
	;
	if v2048 <= int32(0) {
		v2197 = v2042
		goto L138
	} else {
		goto L165
	}
L165:
	;
	v2053 = v1594 + v1381<<(uint(int32(3))%32) + v1822<<(uint(int32(3))%32)
	if v1373 < v1833 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v2055 = v1373
	goto L168
L167:
	;
	v2055 = v1833
	goto L168
L168:
	;
	if v2055 < v1395 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v2057 = v2055
	goto L171
L170:
	;
	v2057 = v1395
	goto L171
L171:
	;
	v2058 = int32(1)
	v2060 = int32(0)
	if v2057 != v2058 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v2067 = int32(0)
	v2070 = v2060
	goto L175
L173:
	;
	v2130 = v2060
	goto L174
L174:
	;
	if v2057&v2058 == int32(0) {
		v2197 = v2042
		goto L138
	} else {
		goto L178
	}
L175:
	;
	v2097 = int32(3)
	v2099 = v2053 + v2070<<(uint(v2097)%32)
	v2100 = *(*int64)(unsafe.Add(mBase, uint32(v2099)))
	v2101 = int32(2)
	v2104 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1406+v2070<<(uint(v2101)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v2099))) = v2100 + v2104*v1859
	v2109 = v2070 | int32(1)
	v2112 = v2053 + v2109<<(uint(v2097)%32)
	v2113 = *(*int64)(unsafe.Add(mBase, uint32(v2112)))
	v2117 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1406+v2109<<(uint(v2101)%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v2112))) = v2113 + v2117*v1859
	v2122 = v2070 + v2101
	v2124 = v2067 + v2101
	if v2124 != v2057&int32(-2) {
		v2067 = v2124
		v2070 = v2122
		goto L175
	} else {
		goto L177
	}
L176:
	;
	v2130 = v2122
	goto L174
L177:
	;
	goto L176
L178:
	;
	v2161 = v2053 + v2130<<(uint(int32(3))%32)
	v2162 = *(*int64)(unsafe.Add(mBase, uint32(v2161)))
	v2166 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1406+v2130<<(uint(int32(2))%32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v2161))) = v2162 + v2166*v1859
	v2197 = v2042
	goto L138
L179:
	;
	goto L137
L180:
	;
	F_pfree(m, v2241)
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L10
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v2246 = int32(2)
	v2250 = F_palloc(m, v1369<<(uint(v2246)%32)|v2246)
	mBase = m.M
	v2251 = m.ExcPending
	if v2251 != 0 {
		goto L10
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v2250
	v2253 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v2250))) = uint16(v2253)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v1369 << (uint(int32(1)) % 32)
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2258 = v2256 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2258
	if v2253 < v1369 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v2268 = v1369
	v2290 = int64(0)
	goto L188
L186:
	;
	goto L187
L187:
	;
	F_pfree(m, v1594)
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L10
	} else {
		goto L194
	}
L188:
	;
	v2295 = v2268 - int32(1)
	v2299 = *(*int64)(unsafe.Add(mBase, uint32(v1594+v2295<<(uint(int32(3))%32))))
	v2300 = v2299 + v2290
	v2304 = v2258 + v2295<<(uint(int32(2))%32)
	if base.Ui64(int64(100000000)) <= base.Ui64(v2300) {
		goto L190
	} else {
		goto L191
	}
L189:
	;
	goto L187
L190:
	;
	v2308 = base.I64_div_u_s(v2300, int64(100000000))
	v2312 = v2308
	v2313 = v2308*int64(-100000000) + v2300
	goto L192
L191:
	;
	v2312 = int64(0)
	v2313 = v2300
	goto L192
L192:
	;
	v2314 = base.I32_wrap_i64(v2313)
	v2315 = int32(10000)
	v2316 = base.I32_div_u_s(v2314, v2315)
	*(*uint16)(unsafe.Add(mBase, uint32(v2304))) = uint16(v2316)
	v2320 = v2314 - v2316*v2315
	*(*uint16)(unsafe.Add(mBase, uint32(v2304)+2)) = uint16(v2320)
	if base.Ui32(int32(1)) < base.Ui32(v2268) {
		v2268 = v2295
		v2290 = v2312
		goto L188
	} else {
		goto L193
	}
L193:
	;
	goto L189
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = base.B2i32(v1393 != v1394) << (uint(int32(14)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v1356
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = l3
	v2368 = l3 + v1356<<(uint(int32(2))%32)
	if v2368+int32(4) < int32(0) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v2484 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2485 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if int32(0) < v2485 {
		goto L224
	} else {
		goto L225
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
	goto L195
L197:
	;
	goto L198
L198:
	;
	v2377 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v2379 = l3 & int32(3)
	v2383 = base.I32_div_s(v2368+int32(7), int32(4))
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v2384 <= v2383 {
		goto L203
	} else {
		goto L204
	}
L199:
	;
	goto L195
L200:
	;
	if int32(0) <= v2450 {
		goto L199
	} else {
		goto L221
	}
L201:
	;
	v2430 = v2424
	goto L215
L202:
	;
	v2397 = int32(1)
	v2398 = v2383 - v2397
	v2401 = v2377 + v2398<<(uint(v2397)%32)
	v2402 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2401))))
	v2403 = int32(2)
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v2379<<(uint(v2403)%32))+uint32(_consts[1299])))
	v2408 = base.I32_rem_s(v2402, v2407)
	v2409 = v2402 - v2408
	*(*uint16)(unsafe.Add(mBase, uint32(v2401))) = uint16(v2409)
	v2412 = base.I32_div_s(v2407, v2403)
	if v2408 < v2412 {
		v2450 = v2398
		goto L200
	} else {
		goto L210
	}
L203:
	;
	if v2379 == int32(0) {
		goto L199
	} else {
		goto L206
	}
L204:
	;
	goto L205
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2383
	if v2379 != 0 {
		goto L202
	} else {
		goto L208
	}
L206:
	;
	if v2383 != v2384 {
		goto L199
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2383
	goto L202
L208:
	;
	v2394 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2377+v2383<<(uint(int32(1))%32)))))
	if v2394 <= int32(4999) {
		v2450 = v2383
		goto L200
	} else {
		goto L209
	}
L209:
	;
	v2424 = v2383
	goto L201
L210:
	;
	v2415 = v2407 + base.I32_extend16_s(v2409)
	if int32(9999) < v2415 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v2420 = v2415 + int32(55536)
	goto L213
L212:
	;
	v2420 = v2415
	goto L213
L213:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2401))) = uint16(v2420)
	if v2415 < int32(10000) {
		v2450 = v2398
		goto L200
	} else {
		goto L214
	}
L214:
	;
	v2424 = v2398
	goto L201
L215:
	;
	v2436 = int32(1)
	v2437 = v2430 - v2436
	v2440 = v2377 + v2437<<(uint(v2436)%32)
	v2443 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2440))))
	v2445 = base.B2i32(int32(9998) < v2443)
	if int32(9998) < v2443 {
		goto L217
	} else {
		goto L218
	}
L216:
	;
	v2450 = v2437
	goto L200
L217:
	;
	v2446 = int32(-9999)
	goto L219
L218:
	;
	v2446 = v2436
	goto L219
L219:
	;
	v2447 = v2446 + v2443
	*(*uint16)(unsafe.Add(mBase, uint32(v2440))) = uint16(v2447)
	if int32(9998) < v2443 {
		v2430 = v2437
		goto L215
	} else {
		goto L220
	}
L220:
	;
	goto L216
L221:
	;
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2458 - int32(2)
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v2463 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2462 + v2463
	v2466 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2466 + v2463
	goto L199
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v2612
	*(*int32)(unsafe.Add(mBase, uint32(l2)+20)) = v2615
	return
L223:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2)+4)) = int64(0)
	v2612 = int32(0)
	v2615 = v2581
	goto L222
L224:
	;
	v2495 = v2485
	v2498 = v2484
	goto L228
L225:
	;
	goto L226
L226:
	;
	if v2485 != 0 {
		v2612 = v2485
		v2615 = v2484
		goto L222
	} else {
		goto L236
	}
L227:
	;
	v2539 = v2495
	goto L232
L228:
	;
	v2522 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2498))))
	if v2522 != 0 {
		goto L227
	} else {
		goto L230
	}
L229:
	;
	v2581 = v2484 + v2485<<(uint(int32(1))%32)
	goto L223
L230:
	;
	v2523 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v2524 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v2523 - v2524
	if v2524 < v2495 {
		v2495 = v2495 - v2524
		v2498 = v2498 + int32(2)
		goto L228
	} else {
		goto L231
	}
L231:
	;
	goto L229
L232:
	;
	v2569 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2498-int32(2)+v2539<<(uint(int32(1))%32)))))
	if v2569 != 0 {
		v2612 = v2539
		v2615 = v2498
		goto L222
	} else {
		goto L234
	}
L233:
	;
	v2581 = v2498
	goto L223
L234:
	;
	v2570 = int32(1)
	if v2570 < v2539 {
		v2539 = v2539 - v2570
		goto L232
	} else {
		goto L235
	}
L235:
	;
	goto L233
L236:
	;
	v2581 = v2484
	goto L223
}
