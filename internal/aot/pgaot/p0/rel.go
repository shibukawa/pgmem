package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_fillRelOptions(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 float64
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v283 int32
	_ = v283
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v352 int32
	_ = v352
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	v8 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	if v8 < l3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v25 = l1
	v36 = v8
	goto L4
L2:
	;
	v370 = l1
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v370 << (uint(int32(2)) % 32)
	m.G0 = v18 + int32(32)
	return
L4:
	;
	if l6 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v370 = v352
	goto L3
L6:
	;
	v367 = v36 + int32(1)
	if v367 != l3 {
		v25 = v352
		v36 = v367
		goto L4
	} else {
		goto L105
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(0)
	v352 = v25
	goto L6
L8:
	;
	v343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v343 != 0 {
		goto L102
	} else {
		goto L103
	}
L9:
	;
	if l4 == int32(0) {
		v352 = v25
		goto L6
	} else {
		goto L98
	}
L10:
	;
	v41 = l2 + v36<<(uint(int32(4))%32)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v53 = int32(0)
	goto L11
L11:
	;
	v62 = l5 + v53<<(uint(int32(4))%32)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63))))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
	if v67 == int32(0) {
		v86 = v66
		v87 = v67
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	v93 = v92 + l0
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v62)+12))
	if int32(0) < v94 {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	if v87-v86 != 0 {
		goto L21
	} else {
		goto L22
	}
L14:
	;
	goto L13
L15:
	;
	if v66 != v67 {
		v86 = v66
		v87 = v67
		goto L14
	} else {
		goto L16
	}
L16:
	;
	v71 = v43
	v72 = v63
	goto L17
L17:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v76 == int32(0) {
		v86 = v75
		v87 = v76
		goto L14
	} else {
		goto L19
	}
L18:
	;
	v86 = v75
	v87 = v76
	goto L14
L19:
	;
	v79 = int32(1)
	if v75 == v76 {
		v71 = v71 + v79
		v72 = v72 + v79
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v90 = v53 + int32(1)
	if l6 != v90 {
		v53 = v90
		goto L11
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	goto L12
L24:
	;
	goto L9
L25:
	;
	v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0+v94))) = uint8(v98)
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v101 = v100
	goto L27
L26:
	;
	v101 = v42
	goto L27
L27:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)+20))
	switch v102 {
	case 0:
		goto L8
	case 1:
		goto L32
	case 2:
		goto L31
	case 3:
		goto L30
	case 4:
		goto L29
	default:
		goto L28
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L52
	} else {
		goto L95
	}
L29:
	;
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v127 == int32(1) {
		goto L45
	} else {
		goto L46
	}
L30:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v123 != 0 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v115 != 0 {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)))
	if v107 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v108 = v41 + int32(8)
	goto L35
L34:
	;
	v108 = v101 + int32(24)
	goto L35
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v109
	v352 = v25
	goto L6
L36:
	;
	v116 = v41 + int32(8)
	goto L38
L37:
	;
	v116 = v101 + int32(24)
	goto L38
L38:
	;
	v117 = *(*float64)(unsafe.Add(mBase, uint32(v116)))
	*(*float64)(unsafe.Add(mBase, uint32(v93))) = v117
	v352 = v25
	goto L6
L39:
	;
	v124 = v41 + int32(8)
	goto L41
L40:
	;
	v124 = v101 + int32(28)
	goto L41
L41:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v125
	v352 = v25
	goto L6
L42:
	;
	v151 = l0 + v25
	if (v136^v151)&int32(3) != 0 {
		goto L60
	} else {
		goto L61
	}
L43:
	;
	v145 = m.T0[v143].(func(*base.Module, int32, int32) int32)(m, v142, l0+v25)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L52
	} else {
		goto L53
	}
L44:
	;
	v138 = int32(0)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v101)+36))
	if v139 == v138 {
		goto L7
	} else {
		goto L51
	}
L45:
	;
	v135 = v41 + int32(8)
	goto L47
L46:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+28)))
	if v132 != 0 {
		goto L44
	} else {
		goto L48
	}
L47:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v101)+36))
	if v137 != 0 {
		v142 = v136
		v143 = v137
		goto L43
	} else {
		goto L49
	}
L48:
	;
	v135 = v101 + int32(40)
	goto L47
L49:
	;
	if v136 != 0 {
		goto L42
	} else {
		goto L50
	}
L50:
	;
	goto L7
L51:
	;
	v142 = v138
	v143 = v139
	goto L43
L52:
	;
	return
L53:
	;
	if v145 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v25
	v352 = v25 + v145
	goto L6
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = int32(0)
	v352 = v25
	goto L6
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93))) = v25
	if v136&int32(3) == int32(0) {
		v250 = v136
		goto L80
	} else {
		goto L81
	}
L58:
	;
	goto L57
L59:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v206))) = uint8(v205)
	if v205&int32(255) == int32(0) {
		goto L58
	} else {
		goto L74
	}
L60:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	v204 = v136
	v205 = v157
	v206 = v151
	goto L59
L61:
	;
	goto L62
L62:
	;
	if v136&int32(3) != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v161 = v136
	v163 = v151
	goto L66
L64:
	;
	v175 = v136
	v177 = v151
	goto L65
L65:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175)))
	v182 = int32(-2139062144)
	if (int32(16843008)-v179|v179)&v182 != v182 {
		v204 = v175
		v205 = v179
		v206 = v177
		goto L59
	} else {
		goto L70
	}
L66:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v161))))
	*(*uint8)(unsafe.Add(mBase, uint32(v163))) = uint8(v164)
	if v164 == int32(0) {
		goto L58
	} else {
		goto L68
	}
L67:
	;
	v175 = v171
	v177 = v169
	goto L65
L68:
	;
	v168 = int32(1)
	v169 = v163 + v168
	v171 = v161 + v168
	if v171&int32(3) != 0 {
		v161 = v171
		v163 = v169
		goto L66
	} else {
		goto L69
	}
L69:
	;
	goto L67
L70:
	;
	v187 = v175
	v188 = v179
	v189 = v177
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v189))) = v188
	v191 = int32(4)
	v192 = v189 + v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v187)+4))
	v195 = v187 + v191
	v199 = int32(-2139062144)
	if (v193|(int32(16843008)-v193))&v199 == v199 {
		v187 = v195
		v188 = v193
		v189 = v192
		goto L71
	} else {
		goto L73
	}
L72:
	;
	v204 = v195
	v205 = v193
	v206 = v192
	goto L59
L73:
	;
	goto L72
L74:
	;
	v213 = v204
	v215 = v206
	goto L75
L75:
	;
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)) = uint8(v216)
	v218 = int32(1)
	if v216 != 0 {
		v213 = v213 + v218
		v215 = v215 + v218
		goto L75
	} else {
		goto L77
	}
L76:
	;
	goto L58
L77:
	;
	goto L76
L78:
	;
	v352 = v283 + v25 + int32(1)
	goto L6
L79:
	;
	v283 = v275 - v136
	goto L78
L80:
	;
	v254 = v250
	goto L89
L81:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v234 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v283 = int32(0)
	goto L78
L83:
	;
	goto L84
L84:
	;
	v239 = v136
	goto L85
L85:
	;
	v243 = v239 + int32(1)
	if v243&int32(3) == int32(0) {
		v250 = v243
		goto L80
	} else {
		goto L87
	}
L86:
	;
	v275 = v243
	goto L79
L87:
	;
	v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v243))))
	if v248 != 0 {
		v239 = v243
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v254)))
	v263 = int32(-2139062144)
	if (int32(16843008)-v260|v260)&v263 == v263 {
		v254 = v254 + int32(4)
		goto L89
	} else {
		goto L91
	}
L90:
	;
	v269 = v254
	goto L92
L91:
	;
	goto L90
L92:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269))))
	if v273 != 0 {
		v269 = v269 + int32(1)
		goto L92
	} else {
		goto L94
	}
L93:
	;
	v275 = v269
	goto L79
L94:
	;
	goto L93
L95:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v291)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v292
	F_errmsg_internal(m, int32(465619), v18)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L52
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(482714), int32(1850), int32(135219))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L52
	} else {
		goto L97
	}
L97:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L52
	} else {
		goto L99
	}
L99:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l2+v36<<(uint(int32(4))%32))))
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v326)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v327
	F_errmsg_internal(m, int32(385335), v18+int32(16))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L52
	} else {
		goto L100
	}
L100:
	;
	F_errfinish(m, int32(482714), int32(1859), int32(135219))
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L52
	} else {
		goto L101
	}
L101:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L102:
	;
	v344 = v41 + int32(8)
	goto L104
L103:
	;
	v344 = v101 + int32(24)
	goto L104
L104:
	;
	v345 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344))))
	*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v345)
	v352 = v25
	goto L6
L105:
	;
	goto L5
}
func F_get_rel_data_width(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v69 int32
	_ = v69
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	v7 = int64(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v9 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+120)))
	if int32(0) < v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = v8
	v17 = int32(1)
	v19 = v7
	goto L4
L2:
	;
	v77 = v7
	goto L3
L3:
	;
	v78 = int64(1073741823)
	if v78 <= v77 {
		goto L22
	} else {
		goto L23
	}
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v27 = v20 + v21<<(uint(int32(4))%32) + v17*int32(100)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+11)))
	if v28 != 0 {
		v63 = v15
		v66 = v19
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v77 = v66
	goto L3
L6:
	;
	v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v63)+120)))
	if v17 < v69 {
		v15 = v63
		v17 = v17 + int32(1)
		v19 = v66
		goto L4
	} else {
		goto L20
	}
L7:
	;
	if l1 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v42 = F_get_attavgwidth(m, v40, base.I32_extend16_s(v17))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1+v17<<(uint(int32(2))%32))))
	if v34 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v63 = v15
	v66 = v19 + base.I64_extend_i32_u(v34)
	goto L6
L11:
	;
	return int32(0)
L12:
	;
	if v42 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = v27 - int32(80)
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+68))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v49)+76))
	v52 = F_get_typavgwidth(m, v50, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	v54 = v42
	goto L15
L15:
	;
	if l1 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v54 = v52
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1+v17<<(uint(int32(2))%32)))) = v54
	goto L19
L18:
	;
	goto L19
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v63 = v62
	v66 = v19 + base.I64_extend_i32_s(v54)
	goto L6
L20:
	;
	goto L5
L21:
	;
	return base.I32_wrap_i64(v81)
L22:
	;
	v81 = v78
	goto L24
L23:
	;
	v81 = v77
	goto L24
L24:
	;
	goto L21
}
func F_get_rel_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	v4 = F_SearchSysCache1(m, int32(57), l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v4)+16))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+22)))
			v17 = F_pstrdup(m, v12+v13+int32(4))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				F_ReleaseCatCache(m, v4)
				mBase = m.M
				v20 = m.ExcPending
				if v20 != 0 {
					return int32(0)
				} else {
					return v17
				}
			}
		}
	}
}
func F_rel_supports_distinctness(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v69 int32
	_ = v69
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v5 != 0 {
		v69 = int32(0)
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v69
L2:
	;
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
	switch v6 {
	case 0:
		goto L5
	case 1:
		goto L4
	default:
		goto L3
	}
L3:
	;
	v69 = int32(0)
	goto L1
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39+v40<<(uint(int32(2))%32))))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)+36))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+120))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+38)))
	if v47 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v7 == int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v10 <= int32(0) {
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v13 = int32(0)
	if v13 < v10 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v17 = v10
	goto L10
L9:
	;
	v17 = v13
	goto L10
L10:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v20 = v13
	goto L11
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v18+v20<<(uint(int32(2))%32))))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+101)))
	if v27 != int32(1) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L3
L13:
	;
	v37 = v20 + int32(1)
	if v37 != v17 {
		v20 = v37
		goto L11
	} else {
		goto L17
	}
L14:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+103)))
	if v30 != int32(1) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+88))
	if v33 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	return int32(1)
L17:
	;
	goto L12
L18:
	;
	v50 = int32(1)
	if v46 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v46 == int32(0) {
		goto L3
	} else {
		goto L27
	}
L21:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+100))
	if v51 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v45)+108))
	if v52 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+36)))
	if v53 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L24
	}
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v45)+112))
	if v54 != 0 {
		v69 = v50
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v45)+144))
	if v55 == int32(0) {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v69 = v50
	goto L1
L27:
	;
	return int32(1)
}
func F_remove_rel_from_joinlist(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	v4 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	if l0 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v73
L2:
	;
	v19 = v4
	v22 = v4
	goto L7
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v13 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v73 = v4
	goto L1
L6:
	;
	goto L5
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24+v22<<(uint(int32(2))%32))))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v29 != int32(1) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v73 = v64
	goto L1
L9:
	;
	v67 = v22 + int32(1)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v67 < v68 {
		v19 = v64
		v22 = v67
		goto L7
	} else {
		goto L25
	}
L10:
	;
	v62 = F_lappend(m, v19, v28)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L17
	} else {
		goto L24
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L17
	} else {
		goto L21
	}
L12:
	;
	if v29 != int32(63) {
		goto L11
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v40 = F_remove_rel_from_joinlist(m, v28, l1, l2)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v34 != l1 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v36 + int32(1)
	v64 = v19
	goto L9
L17:
	;
	return int32(0)
L18:
	;
	if v40 == int32(0) {
		v64 = v19
		goto L9
	} else {
		goto L19
	}
L19:
	;
	v46 = F_lappend(m, v19, v40)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v64 = v46
	goto L9
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v52
	F_errmsg_internal(m, int32(474944), v11)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(482739), int32(822), int32(72038))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L17
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
	v64 = v62
	goto L9
L25:
	;
	goto L8
}
