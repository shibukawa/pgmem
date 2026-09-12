package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_add_partial_path(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v125 int32
	_ = v125
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	v3 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v17 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v187 = F_list_insert_nth(m, v177, v181, l1)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L4
	} else {
		goto L65
	}
L7:
	;
	v177 = v3
	v181 = v3
	goto L6
L8:
	;
	goto L9
L9:
	;
	v22 = v3
	v23 = v17
	v26 = v3
	goto L10
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v22 < v32 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v177 = v174
	v181 = v168
	goto L6
L12:
	;
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)+64))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+v22<<(uint(int32(2))%32))))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+64))
	if v35 == v41 {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	v168 = v26
	goto L14
L14:
	;
	goto L11
L15:
	;
	if v155 != 0 {
		v22 = v154 + int32(1)
		v23 = v155
		v26 = v157
		goto L10
	} else {
		goto L64
	}
L16:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v147 = F_list_delete_nth_cell(m, v146, v22)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L4
	} else {
		goto L62
	}
L17:
	;
	F_pfree(m, l1)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L61
	}
L18:
	;
	v130 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v131 = *(*float64)(unsafe.Add(mBase, uint32(v40)+56))
	if base.F64_ge(v130, v131) != 0 {
		goto L56
	} else {
		goto L57
	}
L19:
	;
	if v98 == int32(3) {
		v125 = v34
		goto L18
	} else {
		goto L41
	}
L20:
	;
	v98 = int32(0)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v49 = int32(0)
	goto L25
L23:
	;
	if v88 != 0 {
		goto L38
	} else {
		goto L39
	}
L24:
	;
	v82 = int32(0)
	v88 = base.B2i32(v62 == v82)
	v90 = base.B2i32(v71 != v82) << (uint(int32(1)) % 32)
	goto L23
L25:
	;
	v52 = int32(0)
	if v35 == v52 {
		v62 = v52
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v98 = int32(3)
	goto L19
L27:
	;
	if v41 != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v56 <= v49 {
		v62 = int32(0)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v62 = v58 + v49<<(uint(int32(2))%32)
	goto L27
L30:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v71 = v68 + v49<<(uint(int32(2))%32)
	if v62 == int32(0) {
		goto L24
	} else {
		goto L35
	}
L31:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v49 < v63 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v65 = int32(0)
	v88 = base.B2i32(v62 == v65)
	v90 = v65
	goto L23
L34:
	;
	goto L33
L35:
	;
	if v71 == int32(0) {
		goto L24
	} else {
		goto L36
	}
L36:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v78 == v79 {
		v49 = v49 + int32(1)
		goto L25
	} else {
		goto L37
	}
L37:
	;
	goto L26
L38:
	;
	v92 = v90
	goto L40
L39:
	;
	v92 = int32(1)
	goto L40
L40:
	;
	v98 = v92
	goto L19
L41:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v101 != v102 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v125 = int32(0)
	goto L18
L43:
	;
	if v101 <= v102 {
		goto L16
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v105 = *(*float64)(unsafe.Add(mBase, uint32(l1)+56))
	v106 = *(*float64)(unsafe.Add(mBase, uint32(v40)+56))
	if base.F64_gt(v105, base.F64_mul(v106, float64(1.01))) != 0 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L42
L47:
	;
	v125 = base.B2i32(v98 == int32(1))
	goto L18
L48:
	;
	goto L49
L49:
	;
	if base.F64_lt(base.F64_mul(v105, float64(1.01)), v106) != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	if v98 == int32(2) {
		v125 = v34
		goto L18
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	switch v98 - int32(1) {
	case 0:
		goto L16
	case 1:
		goto L17
	default:
		goto L54
	}
L53:
	;
	goto L16
L54:
	;
	if base.F64_gt(v106, base.F64_mul(v105, float64(1.0000000001))) != 0 {
		goto L16
	} else {
		goto L55
	}
L55:
	;
	goto L42
L56:
	;
	if v125 == int32(0) {
		goto L17
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	if v125 != 0 {
		v154 = v22
		v155 = v23
		v157 = v26
		goto L15
	} else {
		goto L60
	}
L59:
	;
	v154 = v22
	v155 = v23
	v157 = v22 + int32(1)
	goto L15
L60:
	;
	goto L17
L61:
	;
	return
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v147
	F_pfree(m, v40)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L4
	} else {
		goto L63
	}
L63:
	;
	v154 = v22 - int32(1)
	v155 = v147
	v157 = v26
	goto L15
L64:
	;
	v168 = v157
	goto L14
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v187
	return
}
func F_add_partial_path_precheck(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v94 float64
	_ = v94
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 float64
	_ = v156
	var v164 float64
	_ = v164
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v299 int32
	_ = v299
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v11 == v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v299
L2:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v124 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v14 <= int32(0) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v23 = v5
	goto L5
L5:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29+v23<<(uint(int32(2))%32))))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+64))
	if l3 == v34 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	goto L2
L7:
	;
	v111 = v23 + int32(1)
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v111 < v112 {
		v23 = v111
		goto L5
	} else {
		goto L34
	}
L8:
	;
	if v91 == int32(3) {
		goto L7
	} else {
		goto L30
	}
L9:
	;
	v91 = int32(0)
	goto L8
L10:
	;
	goto L11
L11:
	;
	v42 = int32(0)
	goto L14
L12:
	;
	if v81 != 0 {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	v75 = int32(0)
	v81 = base.B2i32(v55 == v75)
	v83 = base.B2i32(v64 != v75) << (uint(int32(1)) % 32)
	goto L12
L14:
	;
	v45 = int32(0)
	if l3 == v45 {
		v55 = v45
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v91 = int32(3)
	goto L8
L16:
	;
	if v34 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v49 <= v42 {
		v55 = int32(0)
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v55 = v51 + v42<<(uint(int32(2))%32)
	goto L16
L19:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v34)+12))
	v64 = v61 + v42<<(uint(int32(2))%32)
	if v55 == int32(0) {
		goto L13
	} else {
		goto L24
	}
L20:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	if v42 < v56 {
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v58 = int32(0)
	v81 = base.B2i32(v55 == v58)
	v83 = v58
	goto L12
L23:
	;
	goto L22
L24:
	;
	if v64 == int32(0) {
		goto L13
	} else {
		goto L25
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	if v71 == v72 {
		v42 = v42 + int32(1)
		goto L14
	} else {
		goto L26
	}
L26:
	;
	goto L15
L27:
	;
	v85 = v83
	goto L29
L28:
	;
	v85 = int32(1)
	goto L29
L29:
	;
	v91 = v85
	goto L8
L30:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(v33)+56))
	v98 = int32(0)
	v102 = base.B2i32(base.F64_gt(l2, base.F64_mul(v94, float64(1.01))) == v98) | base.B2i32(v91 == int32(1))
	if v102 == v98 {
		v299 = v102
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v91 == int32(2) {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	if base.F64_lt(base.F64_mul(l2, float64(1.01)), v94) != 0 {
		v299 = v102
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L7
L34:
	;
	goto L6
L35:
	;
	return int32(1)
L36:
	;
	goto L37
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v129 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	return int32(1)
L39:
	;
	goto L40
L40:
	;
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	v142 = int32(0)
	goto L41
L41:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v146+v142<<(uint(int32(2))%32))))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+40))
	if l1 != v151 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v299 = v289
	goto L1
L43:
	;
	if v135 != 0 {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v151 <= l1 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v150)+56))
	if base.F64_le(l2, base.F64_mul(v156, float64(1.01))) == int32(0) {
		goto L43
	} else {
		goto L48
	}
L47:
	;
	return int32(1)
L48:
	;
	return int32(1)
L49:
	;
	v289 = int32(1)
	v291 = v142 + v289
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	if v291 < v292 {
		v142 = v291
		goto L41
	} else {
		goto L96
	}
L50:
	;
	v164 = *(*float64)(unsafe.Add(mBase, uint32(v150)+48))
	if base.F64_gt(l2, base.F64_mul(v164, float64(1.01))) == int32(0) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v170 = int32(0)
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	if v171 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v173 = v170
	goto L56
L55:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v150)+64))
	v173 = v172
	goto L56
L56:
	;
	if l3 == v173 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v230&int32(-3) != 0 {
		goto L49
	} else {
		goto L79
	}
L58:
	;
	v230 = int32(0)
	goto L57
L59:
	;
	goto L60
L60:
	;
	v181 = int32(0)
	goto L63
L61:
	;
	if v220 != 0 {
		goto L76
	} else {
		goto L77
	}
L62:
	;
	v214 = int32(0)
	v220 = base.B2i32(v194 == v214)
	v222 = base.B2i32(v203 != v214) << (uint(int32(1)) % 32)
	goto L61
L63:
	;
	v184 = int32(0)
	if l3 == v184 {
		v194 = v184
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v230 = int32(3)
	goto L57
L65:
	;
	if v173 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v188 <= v181 {
		v194 = int32(0)
		goto L65
	} else {
		goto L67
	}
L67:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v194 = v190 + v181<<(uint(int32(2))%32)
	goto L65
L68:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v203 = v200 + v181<<(uint(int32(2))%32)
	if v194 == int32(0) {
		goto L62
	} else {
		goto L73
	}
L69:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if v181 < v195 {
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v197 = int32(0)
	v220 = base.B2i32(v194 == v197)
	v222 = v197
	goto L61
L72:
	;
	goto L71
L73:
	;
	if v203 == int32(0) {
		goto L62
	} else {
		goto L74
	}
L74:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v203)))
	if v210 == v211 {
		v181 = v181 + int32(1)
		goto L63
	} else {
		goto L75
	}
L75:
	;
	goto L64
L76:
	;
	v224 = v222
	goto L78
L77:
	;
	v224 = int32(1)
	goto L78
L78:
	;
	v230 = v224
	goto L57
L79:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v150)+16))
	if v234 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
	v236 = v235
	goto L82
L81:
	;
	v236 = v170
	goto L82
L82:
	;
	goto L84
L83:
	;
	if int32(0)|v236 == int32(0) {
		v299 = v170
		goto L1
	} else {
		goto L95
	}
L84:
	;
	goto L83
L95:
	;
	goto L49
L96:
	;
	goto L42
}
func F_findPartialMatch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
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
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
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
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v243 int32
	_ = v243
	var v264 int32
	_ = v264
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = v17 + int32(4)
	v27 = int32(-1)
	v28 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	if v28 == int64(0) {
		v50 = v27
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v59 = v17 + int32(4)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+8)))
	v66 = v63
	goto L12
L2:
	;
	v53 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+8)) = uint8(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v50
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v50
	goto L1
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v33 = int32(0)
	goto L4
L4:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31+v33*int32(12))+4))
	if v41 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v50 = v27
	goto L2
L6:
	;
	v50 = v33
	goto L2
L7:
	;
	goto L8
L8:
	;
	v45 = v33 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v45)) < base.Ui64(v28) {
		v33 = v45
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
L10:
	;
	m.G0 = v17 + int32(16)
	return v264
L11:
	;
	if v98 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L12:
	;
	if v66&int32(1) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v98 = v81
	goto L11
L14:
	;
	v98 = int32(0)
	goto L11
L15:
	;
	goto L16
L16:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+20))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	v77 = v73 & (v74 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v59))) = v77
	v81 = v72 + v74*int32(12)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v85 = v82 & (v83 ^ v77)
	if v85 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v88 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v59)+8)) = uint8(v88)
	goto L19
L18:
	;
	goto L19
L19:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	if v92 != int32(1) {
		v66 = base.B2i32(v85 == int32(0))
		goto L12
	} else {
		goto L20
	}
L20:
	;
	goto L13
L21:
	;
	v264 = int32(0)
	goto L10
L22:
	;
	goto L23
L23:
	;
	v103 = v20 - int32(1)
	v108 = v98
	goto L24
L24:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v119 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v135
	v264 = int32(1)
	goto L10
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v108)))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v127 = F_ExecStoreMinimalTuple(m, v124, v125, int32(0))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L29
	} else {
		goto L31
	}
L29:
	;
	return int32(0)
L30:
	;
	goto L28
L31:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	F_MemoryContextReset(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	v134 = int32(4536272)
	v135 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v131
	if v103 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	goto L25
L34:
	;
	v144 = v103
	goto L35
L35:
	;
	v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19+v144<<(uint(int32(1))%32)))))
	v158 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v158 < v157 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v135
	v202 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v204 = v17 + int32(4)
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v204)+8)))
	v211 = v208
	goto L53
L37:
	;
	F_slot_getsomeattrs_int(m, l1, v157)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L29
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v163 = v157 - int32(1)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v164))))
	if v166 != 0 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L39
L41:
	;
	goto L36
L42:
	;
	if v144 <= int32(0) {
		goto L33
	} else {
		goto L51
	}
L43:
	;
	v168 = v163 << (uint(int32(2)) % 32)
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v168+v169)))
	v172 = int32(*(*int16)(unsafe.Add(mBase, uint32(v130)+6)))
	if v172 < v157 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	F_slot_getsomeattrs_int(m, v130, v157)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L29
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v130)+20))
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176+v163))))
	if v178 != 0 {
		goto L42
	} else {
		goto L48
	}
L47:
	;
	goto L46
L48:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v129+v144<<(uint(int32(2))%32))))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v186+v168)))
	v189 = F_FunctionCall2Coll(m, l2+v144*int32(28), v185, v171, v188)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L29
	} else {
		goto L49
	}
L49:
	;
	if v189 == int32(0) {
		goto L41
	} else {
		goto L50
	}
L50:
	;
	goto L42
L51:
	;
	v144 = v144 - int32(1)
	goto L35
L52:
	;
	if v243 != 0 {
		v108 = v243
		goto L24
	} else {
		goto L62
	}
L53:
	;
	if v211&int32(1) != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v243 = v226
	goto L52
L55:
	;
	v243 = int32(0)
	goto L52
L56:
	;
	goto L57
L57:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v202)+20))
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v204)))
	v222 = v218 & (v219 - int32(1))
	*(*int32)(unsafe.Add(mBase, uint32(v204))) = v222
	v226 = v217 + v219*int32(12)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v202)+12))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v204)+4))
	v230 = v227 & (v228 ^ v222)
	if v230 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v233 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v204)+8)) = uint8(v233)
	goto L60
L59:
	;
	goto L60
L60:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v226)+4))
	if v237 != int32(1) {
		v211 = base.B2i32(v230 == int32(0))
		goto L53
	} else {
		goto L61
	}
L61:
	;
	goto L54
L62:
	;
	v264 = int32(0)
	goto L10
}
