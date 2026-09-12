package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_vac_tid_reaped(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int64
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v277 int64
	_ = v277
	var v279 int32
	_ = v279
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v343 int32
	_ = v343
	v3 = int32(0)
	v11 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
	v12 = int64(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	v15 = v11 | v12<<(uint(int64(16))%64)
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	return v343
L2:
	;
	v343 = v329
	goto L1
L3:
	;
	v307 = int32(*(*int8)(unsafe.Add(mBase, uint32(v297)+1)))
	if v307 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L4:
	;
	if v288 != 0 {
		v297 = v288
		goto L3
	} else {
		goto L71
	}
L5:
	;
	v20 = *(*int64)(unsafe.Add(mBase, uint32(v18)+32))
	if base.Ui64(v20) < base.Ui64(v15) {
		v329 = v3
		goto L2
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v157 = *(*int64)(unsafe.Add(mBase, uint32(v18)+8))
	if base.Ui64(v157) < base.Ui64(v15) {
		v329 = v3
		goto L2
	} else {
		goto L41
	}
L8:
	;
	v22 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+48)))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	v26 = v23
	v32 = v22
	goto L9
L9:
	;
	v35 = base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(v32) % 64)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v37 = F_dsa_get_address(m, v36, v26)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L10:
	;
	if v147&int32(1) != 0 {
		v297 = v135
		goto L3
	} else {
		goto L39
	}
L11:
	;
	v145 = int64(8)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if base.B2i32(v32 < v145) == int32(0) {
		v26 = v147
		v32 = v32 - v145
		goto L9
	} else {
		goto L38
	}
L12:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v37+int32(base.Ui32(v35)>>(uint(int32(3))%32))&int32(28))+4))
	if int32(base.Ui32(v120)>>(uint(v35)%32))&int32(1) == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L36
	}
L13:
	;
	v104 = int32(255)
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v35&v104)+12)))
	if v107 == v104 {
		v329 = v3
		goto L2
	} else {
		goto L34
	}
L14:
	;
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
	if v74 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L26
	}
L15:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37)+2)))
	if v44 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L18
	}
L16:
	;
	return int32(0)
L17:
	;
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	switch v41 - int32(1) {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	default:
		goto L15
	}
L18:
	;
	v52 = int32(0)
	goto L19
L19:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52+(v37+int32(3))))))
	if v35&int32(255) == v63 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v343 = int32(0)
	goto L1
L21:
	;
	v69 = v37 + v52<<(uint(int32(2))%32) + int32(8)
	if v69 != 0 {
		v135 = v69
		goto L11
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v71 = v52 + int32(1)
	if v71 != v44 {
		v52 = v71
		goto L19
	} else {
		goto L25
	}
L24:
	;
	v329 = v3
	goto L2
L25:
	;
	goto L20
L26:
	;
	v80 = int32(0)
	goto L27
L27:
	;
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+(v37+int32(3))))))
	if v91 == v35&int32(255) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v343 = int32(0)
	goto L1
L29:
	;
	v99 = v37 + v80<<(uint(int32(2))%32) + int32(36)
	if v99 != 0 {
		v135 = v99
		goto L11
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v101 = v80 + int32(1)
	if v101 != v74 {
		v80 = v101
		goto L27
	} else {
		goto L33
	}
L32:
	;
	v329 = v3
	goto L2
L33:
	;
	goto L28
L34:
	;
	v114 = v37 + v107<<(uint(int32(2))%32) + int32(268)
	if v114 != 0 {
		v135 = v114
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v329 = v3
	goto L2
L36:
	;
	v132 = v37 + v35&int32(255)<<(uint(int32(2))%32) + int32(36)
	if v132 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v135 = v132
	goto L11
L38:
	;
	goto L10
L39:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v155 = F_dsa_get_address(m, v154, v147)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L16
	} else {
		goto L40
	}
L40:
	;
	v288 = v155
	goto L4
L41:
	;
	v159 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v18)+24)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v163 = v160
	v169 = v159
	goto L42
L42:
	;
	v172 = base.I32_wrap_i64(int64(base.Ui64(v15) >> (uint(v169) % 64)))
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163))))
	switch v173 - int32(1) {
	case 0:
		goto L47
	case 1:
		goto L46
	case 2:
		goto L45
	default:
		goto L48
	}
L43:
	;
	if v279&int32(1) != 0 {
		v297 = v267
		goto L3
	} else {
		goto L70
	}
L44:
	;
	v277 = int64(8)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if base.B2i32(v169 < v277) == int32(0) {
		v163 = v279
		v169 = v169 - v277
		goto L42
	} else {
		goto L69
	}
L45:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v163+int32(base.Ui32(v172)>>(uint(int32(3))%32))&int32(28))+4))
	if int32(base.Ui32(v252)>>(uint(v172)%32))&int32(1) == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L67
	}
L46:
	;
	v236 = int32(255)
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v172&v236)+12)))
	if v239 == v236 {
		v329 = v3
		goto L2
	} else {
		goto L65
	}
L47:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+2)))
	if v206 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L57
	}
L48:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+2)))
	if v176 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L49
	}
L49:
	;
	v184 = int32(0)
	goto L50
L50:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+(v163+int32(3))))))
	if v172&int32(255) == v195 {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v343 = int32(0)
	goto L1
L52:
	;
	v201 = v163 + v184<<(uint(int32(2))%32) + int32(8)
	if v201 != 0 {
		v267 = v201
		goto L44
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v203 = v184 + int32(1)
	if v203 != v176 {
		v184 = v203
		goto L50
	} else {
		goto L56
	}
L55:
	;
	v329 = v3
	goto L2
L56:
	;
	goto L51
L57:
	;
	v212 = int32(0)
	goto L58
L58:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v212+(v163+int32(3))))))
	if v223 == v172&int32(255) {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v343 = int32(0)
	goto L1
L60:
	;
	v231 = v163 + v212<<(uint(int32(2))%32) + int32(36)
	if v231 != 0 {
		v267 = v231
		goto L44
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v233 = v212 + int32(1)
	if v233 != v206 {
		v212 = v233
		goto L58
	} else {
		goto L64
	}
L63:
	;
	v329 = v3
	goto L2
L64:
	;
	goto L59
L65:
	;
	v246 = v163 + v239<<(uint(int32(2))%32) + int32(268)
	if v246 != 0 {
		v267 = v246
		goto L44
	} else {
		goto L66
	}
L66:
	;
	v329 = v3
	goto L2
L67:
	;
	v264 = v163 + v172&int32(255)<<(uint(int32(2))%32) + int32(36)
	if v264 == int32(0) {
		v329 = v3
		goto L2
	} else {
		goto L68
	}
L68:
	;
	v267 = v264
	goto L44
L69:
	;
	goto L43
L70:
	;
	v288 = v279
	goto L4
L71:
	;
	v343 = int32(0)
	goto L1
L72:
	;
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+2)))
	v343 = base.B2i32(v310 == v16)
	goto L1
L73:
	;
	goto L74
L74:
	;
	v314 = int32(base.Ui32(v16) >> (uint(int32(5)) % 32))
	if v307 <= v314 {
		v343 = int32(0)
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v297+v314<<(uint(int32(2))%32))+4))
	v329 = int32(base.Ui32(v319)>>(uint(v16)%32)) & int32(1)
	goto L2
}
func F_validate_option_array_item(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	v4 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(16)
	m.G0 = v15
	if l1 != 0 {
		v176 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v184 = int32(1)
	v189 = F_find_option(m, l0, v184, (l2|v176)&v184, int32(21))
	mBase = m.M
	v192 = m.ExcPending
	if v192 != 0 {
		goto L44
	} else {
		goto L45
	}
L2:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v17 == int32(0) {
		v176 = v4
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v17
	v27 = int32(1)
	v29 = l0
	v32 = v4
	goto L4
L4:
	;
	v34 = v24 & int32(255)
	if v34 == int32(46) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v176 = base.B2i32(v24&int32(255) != int32(46)) & v161
	goto L1
L6:
	;
	v165 = v29 + int32(1)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v166 != 0 {
		v24 = v166
		v27 = base.B2i32(v34 == int32(46))
		v29 = v165
		v32 = v161
		goto L4
	} else {
		goto L42
	}
L7:
	;
	if v27 == int32(0) {
		v161 = int32(1)
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v41 = int32(_a_F_validate_option_array_item_0)
	v42 = base.I32_extend8_s(v24)
	v43 = int32(54)
	goto L14
L10:
	;
	v176 = int32(0)
	goto L1
L11:
	;
	if v42 < int32(0) {
		v161 = v32
		goto L6
	} else {
		goto L37
	}
L12:
	;
	v146 = int32(0)
	goto L11
L13:
	;
	v124 = v117
	v126 = v119
	goto L31
L14:
	;
	goto L22
L22:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_validate_option_array_item[0])))
	if v80 == v42&int32(255) {
		v110 = v41
		v112 = v43
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v112 == int32(0) {
		goto L12
	} else {
		goto L30
	}
L24:
	;
	goto L25
L25:
	;
	v90 = v41
	v92 = v43
	goto L26
L26:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	v97 = v96 ^ v42&int32(255)*int32(16843009)
	v100 = int32(-2139062144)
	if (int32(16843008)-v97|v97)&v100 != v100 {
		v117 = v90
		v119 = v92
		goto L13
	} else {
		goto L28
	}
L27:
	;
	v110 = v105
	v112 = v107
	goto L23
L28:
	;
	v104 = int32(4)
	v105 = v90 + v104
	v107 = v92 - v104
	if base.Ui32(int32(3)) < base.Ui32(v107) {
		v90 = v105
		v92 = v107
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v117 = v110
	v119 = v112
	goto L13
L31:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	if v42&int32(255) == v129 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L12
L33:
	;
	v146 = v124
	goto L11
L34:
	;
	goto L35
L35:
	;
	v131 = int32(1)
	v134 = v126 - v131
	if v134 != 0 {
		v124 = v124 + v131
		v126 = v134
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	if v146 != 0 {
		v161 = v32
		goto L6
	} else {
		goto L38
	}
L38:
	;
	v149 = int32(0)
	if v27 != 0 {
		v176 = v149
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(63)) < base.Ui32(v34) {
		v176 = v149
		goto L1
	} else {
		goto L40
	}
L40:
	;
	if int64(1)<<(uint(base.I64_extend_i32_u(v42))%64)&int64(287948969894477825) == int64(0) {
		v176 = v149
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v161 = v32
	goto L6
L42:
	;
	goto L5
L43:
	;
	m.G0 = v15 + int32(16)
	return v261
L44:
	;
	return int32(0)
L45:
	;
	if v189|v176 == int32(0) {
		v261 = v4
		goto L43
	} else {
		goto L46
	}
L46:
	;
	if v189 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v189)+4))
	switch v229 - int32(5) {
	case 0:
		goto L63
	case 1:
		goto L61
	default:
		goto L62
	}
L48:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+21)))
	if v196&int32(2) == int32(0) {
		goto L47
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v202 = F_superuser(m)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L44
	} else {
		goto L52
	}
L51:
	;
	goto L50
L52:
	;
	if v202 != 0 {
		v261 = int32(1)
		goto L43
	} else {
		goto L53
	}
L53:
	;
	v205 = *(*int32)(unsafe.Add(mBase, _c_F_validate_option_array_item[1]))
	v207 = F_pg_parameter_aclcheck(m, l0, v205, int64(4096))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L44
	} else {
		goto L54
	}
L54:
	;
	v210 = base.B2i32(v207 == int32(0))
	if l2 != 0 {
		v261 = v210
		goto L43
	} else {
		goto L55
	}
L55:
	;
	if v207 == int32(0) {
		v261 = v210
		goto L43
	} else {
		goto L56
	}
L56:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L44
	} else {
		goto L57
	}
L57:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L44
	} else {
		goto L58
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = l0
	F_errmsg(m, int32(_a_F_validate_option_array_item_1), v15)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L44
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_validate_option_array_item_2), int32(_a_F_validate_option_array_item_3), int32(_a_F_validate_option_array_item_4))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L44
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v247 = F_superuser(m)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L44
	} else {
		goto L70
	}
L62:
	;
	if l2 != 0 {
		v261 = v4
		goto L43
	} else {
		goto L69
	}
L63:
	;
	v232 = F_superuser(m)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L44
	} else {
		goto L64
	}
L64:
	;
	if v232 != 0 {
		goto L61
	} else {
		goto L65
	}
L65:
	;
	v235 = *(*int32)(unsafe.Add(mBase, _c_F_validate_option_array_item[1]))
	v237 = F_pg_parameter_aclcheck(m, l0, v235, int64(4096))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L44
	} else {
		goto L66
	}
L66:
	;
	if l2 == int32(0) {
		goto L61
	} else {
		goto L67
	}
L67:
	;
	if v237 == int32(0) {
		goto L61
	} else {
		goto L68
	}
L68:
	;
	v261 = v4
	goto L43
L69:
	;
	goto L61
L70:
	;
	if v247 != 0 {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v249 = int32(5)
	goto L73
L72:
	;
	v249 = int32(6)
	goto L73
L73:
	;
	v252 = *(*int32)(unsafe.Add(mBase, _c_F_validate_option_array_item[1]))
	v253 = int32(0)
	v257 = F_set_config_with_handle(m, l0, int32(0), l1, v249, int32(12), v252, v253, v253, v253, v253)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L44
	} else {
		goto L74
	}
L74:
	;
	v261 = int32(1)
	goto L43
}
func F_validate_remote_info(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = int64(68719476752)
	F_initStringInfo(m, v6+int32(-24))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[0]))
	v18 = F_quote_literal_cstr(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v18
	F_appendStringInfo(m, v6+int32(-24), int32(_a_F_validate_remote_info_0), v6+int32(-32))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[1]))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+20))
	v32 = base.B2i32(v30 == int32(2))
	goto L5
L5:
	;
	if v32 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[2]))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+60))
	v44 = m.T0[v43].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v37, int32(2), v6+int32(-8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L8
L10:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	F_pfree(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	if v49 == int32(2) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L1
	} else {
		goto L56
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L53
	}
L15:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	v54 = F_MakeSingleTupleTableSlot(m, v52, int32(_a_F_validate_remote_info_1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L49
	}
L18:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	v59 = F_tuplestore_gettupleslot(m, v56, int32(1), int32(0), v54)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	if v59 == int32(0) {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v63 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+6)))
	if v63 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	F_slot_getsomeattrs_int(m, v54, int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v70 != 0 {
		goto L13
	} else {
		goto L25
	}
L24:
	;
	goto L23
L25:
	;
	v71 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+6)))
	if v71 <= int32(1) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	F_slot_getsomeattrs_int(m, v54, int32(2))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v78 = v69
	goto L28
L28:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+4))
	if v79 == int32(0) {
		goto L12
	} else {
		goto L30
	}
L29:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v54)+16))
	v78 = v77
	goto L28
L30:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v54)+8))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)+12))
	m.T0[v83].(func(*base.Module, int32))(m, v54)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v86 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_pfree(m, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v89 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L34
L36:
	;
	F_tuplestore_end(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
	if v92 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L38
L40:
	;
	F_FreeTupleDesc(m, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	F_pfree(m, v44)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	if v32 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	m.G0 = v8 - int32(-64)
	return
L48:
	;
	goto L47
L49:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v110
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v108
	F_errmsg(m, int32(_a_F_validate_remote_info_8), v6+int32(-48))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errhint(m, int32(_a_F_validate_remote_info_9), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(989), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
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
	F_errmsg_internal(m, int32(_a_F_validate_remote_info_2), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(994), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
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
	F_errcode(m, int32(1088))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errmsg(m, int32(_a_F_validate_remote_info_5), int32(0))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(1009), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L60:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(_a_F_validate_remote_info_6)
	v166 = *(*int32)(unsafe.Add(mBase, _c_F_validate_remote_info[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v166
	F_errmsg(m, int32(_a_F_validate_remote_info_7), v8)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(_a_F_validate_remote_info_3), int32(1019), int32(_a_F_validate_remote_info_4))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_varbittypmodin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_anybit_typmodin(m, v3, int32(_a_F_varbittypmodin_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_visibilitymap_count(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v85 int64
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int64
	_ = v122
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v158 int64
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int64
	_ = v168
	var v169 int32
	_ = v169
	var v173 int64
	_ = v173
	var v174 int32
	_ = v174
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v188 int64
	_ = v188
	var v192 int64
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v204 int64
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v211 int64
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v247 int64
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int64
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int64
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v284 int64
	_ = v284
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int64
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v306 int64
	_ = v306
	var v307 int64
	_ = v307
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int64
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v330 int64
	_ = v330
	var v331 int32
	_ = v331
	var v335 int64
	_ = v335
	var v336 int32
	_ = v336
	var v340 int64
	_ = v340
	var v341 int32
	_ = v341
	var v345 int64
	_ = v345
	var v346 int32
	_ = v346
	var v350 int64
	_ = v350
	var v354 int64
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v366 int64
	_ = v366
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	v4 = int32(0)
	v12 = F_vm_readbuf(m, l0, v4, v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v17 = v4
	v18 = v12
	v19 = v4
	v20 = v4
	goto L6
L4:
	;
	v382 = v4
	v384 = v4
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v384
	if l2 != 0 {
		goto L64
	} else {
		goto L65
	}
L6:
	;
	if v18 < int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v382 = v369
	v384 = v371
	goto L5
L8:
	;
	v42 = v40 + int32(24)
	v43 = int32(85)
	v44 = int32(0)
	v49 = int64(0)
	v50 = int32(_a_F_visibilitymap_count_0)
	if (v40+int32(27))&int32(-4) == v42 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_count[0]))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26+(v18^int32(-1))<<(uint(int32(2))%32))))
	v40 = v32
	goto L8
L10:
	;
	goto L11
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_visibilitymap_count[1]))
	v40 = v34 + v18<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L12:
	;
	if l2 != 0 {
		goto L35
	} else {
		goto L36
	}
L13:
	;
	v57 = int32(1431655765)
	v58 = v42
	v61 = v44
	v62 = v50
	v65 = v49
	goto L16
L14:
	;
	v115 = v42
	v119 = v50
	v122 = v49
	goto L15
L15:
	;
	if v119 == int32(0) {
		v204 = v122
		goto L22
	} else {
		goto L23
	}
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	v85 = base.I64_extend_i32_u(base.I32_popcnt(v66&v57)) + (base.I64_extend_i32_u(base.I32_popcnt(v70&v57)) + (base.I64_extend_i32_u(base.I32_popcnt(v74&v57)) + (v65 + base.I64_extend_i32_u(base.I32_popcnt(v78&v57)))))
	v86 = int32(16)
	v87 = v62 - v86
	v89 = v58 + v86
	v91 = v61 + int32(4)
	if v91 != int32(2040) {
		v58 = v89
		v61 = v91
		v62 = v87
		v65 = v85
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v97 = v89
	v98 = v87
	v99 = v44
	v101 = v85
	goto L19
L18:
	;
	goto L17
L19:
	;
	v102 = int32(4)
	v103 = v98 - v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v108 = v101 + base.I64_extend_i32_u(base.I32_popcnt(v104&v57))
	v110 = v97 + v102
	v112 = v99 + int32(1)
	if v112 != int32(2) {
		v97 = v110
		v98 = v103
		v99 = v112
		v101 = v108
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v115 = v110
	v119 = v103
	v122 = v108
	goto L15
L21:
	;
	goto L20
L22:
	;
	goto L12
L23:
	;
	v126 = v119 & int32(3)
	if v126 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if base.Ui32(v119) < base.Ui32(int32(4)) {
		v204 = v158
		goto L22
	} else {
		goto L31
	}
L25:
	;
	v151 = v115
	v153 = v119
	v158 = v122
	goto L24
L26:
	;
	goto L27
L27:
	;
	v132 = v119
	v133 = v115
	v135 = int32(0)
	v137 = v122
	goto L28
L28:
	;
	v138 = int32(1)
	v139 = v132 - v138
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133))))
	v144 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v140&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v145 = v137 + v144
	v147 = v133 + v138
	v149 = v135 + v138
	if v149 != v126 {
		v132 = v139
		v133 = v147
		v135 = v149
		v137 = v145
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v151 = v147
	v153 = v139
	v158 = v145
	goto L24
L30:
	;
	goto L29
L31:
	;
	v161 = v151
	v163 = v153
	v168 = v158
	goto L32
L32:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+3)))
	v173 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v169&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+2)))
	v178 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v174&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161)+1)))
	v183 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v179&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	v188 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v184&v43)+uint32(_c_F_visibilitymap_count[2]))))
	v192 = v173 + (v178 + (v183 + (v168 + v188)))
	v193 = int32(4)
	v196 = v163 - v193
	if v196 != 0 {
		v161 = v161 + v193
		v163 = v196
		v168 = v192
		goto L32
	} else {
		goto L34
	}
L33:
	;
	v204 = v192
	goto L22
L34:
	;
	goto L33
L35:
	;
	v205 = int32(170)
	v206 = int32(0)
	v211 = int64(0)
	v212 = int32(_a_F_visibilitymap_count_0)
	if (v40+int32(27))&int32(-4) == v42 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v369 = v17
	goto L37
L37:
	;
	v371 = v19 + base.I32_wrap_i64(v204)
	F_ReleaseBuffer(m, v18)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L1
	} else {
		goto L61
	}
L38:
	;
	v369 = v17 + base.I32_wrap_i64(v366)
	goto L37
L39:
	;
	v219 = int32(-1431655766)
	v220 = v42
	v223 = v206
	v224 = v212
	v227 = v211
	goto L42
L40:
	;
	v277 = v42
	v281 = v212
	v284 = v211
	goto L41
L41:
	;
	if v281 == int32(0) {
		v366 = v284
		goto L48
	} else {
		goto L49
	}
L42:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v220)+12))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v220)+8))
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v220)+4))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	v247 = base.I64_extend_i32_u(base.I32_popcnt(v228&v219)) + (base.I64_extend_i32_u(base.I32_popcnt(v232&v219)) + (base.I64_extend_i32_u(base.I32_popcnt(v236&v219)) + (v227 + base.I64_extend_i32_u(base.I32_popcnt(v240&v219)))))
	v248 = int32(16)
	v249 = v224 - v248
	v251 = v220 + v248
	v253 = v223 + int32(4)
	if v253 != int32(2040) {
		v220 = v251
		v223 = v253
		v224 = v249
		v227 = v247
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v259 = v251
	v260 = v249
	v261 = v206
	v263 = v247
	goto L45
L44:
	;
	goto L43
L45:
	;
	v264 = int32(4)
	v265 = v260 - v264
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v259)))
	v270 = v263 + base.I64_extend_i32_u(base.I32_popcnt(v266&v219))
	v272 = v259 + v264
	v274 = v261 + int32(1)
	if v274 != int32(2) {
		v259 = v272
		v260 = v265
		v261 = v274
		v263 = v270
		goto L45
	} else {
		goto L47
	}
L46:
	;
	v277 = v272
	v281 = v265
	v284 = v270
	goto L41
L47:
	;
	goto L46
L48:
	;
	goto L38
L49:
	;
	v288 = v281 & int32(3)
	if v288 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	if base.Ui32(v281) < base.Ui32(int32(4)) {
		v366 = v320
		goto L48
	} else {
		goto L57
	}
L51:
	;
	v313 = v277
	v315 = v281
	v320 = v284
	goto L50
L52:
	;
	goto L53
L53:
	;
	v294 = v281
	v295 = v277
	v297 = int32(0)
	v299 = v284
	goto L54
L54:
	;
	v300 = int32(1)
	v301 = v294 - v300
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295))))
	v306 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v302&v205)+uint32(_c_F_visibilitymap_count[2]))))
	v307 = v299 + v306
	v309 = v295 + v300
	v311 = v297 + v300
	if v311 != v288 {
		v294 = v301
		v295 = v309
		v297 = v311
		v299 = v307
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v313 = v309
	v315 = v301
	v320 = v307
	goto L50
L56:
	;
	goto L55
L57:
	;
	v323 = v313
	v325 = v315
	v330 = v320
	goto L58
L58:
	;
	v331 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+3)))
	v335 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v331&v205)+uint32(_c_F_visibilitymap_count[2]))))
	v336 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+2)))
	v340 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v336&v205)+uint32(_c_F_visibilitymap_count[2]))))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323)+1)))
	v345 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v341&v205)+uint32(_c_F_visibilitymap_count[2]))))
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v323))))
	v350 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v346&v205)+uint32(_c_F_visibilitymap_count[2]))))
	v354 = v335 + (v340 + (v345 + (v330 + v350)))
	v355 = int32(4)
	v358 = v325 - v355
	if v358 != 0 {
		v323 = v323 + v355
		v325 = v358
		v330 = v354
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v366 = v354
	goto L48
L60:
	;
	goto L59
L61:
	;
	v375 = v20 + int32(1)
	v377 = F_vm_readbuf(m, l0, v375, int32(0))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	if v377 != 0 {
		v17 = v369
		v18 = v377
		v19 = v371
		v20 = v375
		goto L6
	} else {
		goto L63
	}
L63:
	;
	goto L7
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v382
	goto L66
L65:
	;
	goto L66
L66:
	;
	return
}
