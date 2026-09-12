package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_DecodeNumber(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int64
	_ = v64
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v204 float64
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 float64
	_ = v210
	var v213 float64
	_ = v213
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v242 int32
	_ = v242
	var v257 int32
	_ = v257
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v309 int32
	_ = v309
	v9 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v9
	*(*int32)(unsafe.Add(mBase, _consts[40])) = v9
	v27 = F_strtol(m, l1, v17+int32(8), int32(10))
	mBase = m.M
	goto L1
L1:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v30 == int32(68) {
		v309 = int32(-2)
		goto L2
	} else {
		goto L3
	}
L2:
	;
	m.G0 = v17 + int32(16)
	return v309
L3:
	;
	v33 = int32(-1)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
	if v34 == l1 {
		v309 = v33
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v36 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v224 = l3 & int32(14)
	if l0 != int32(3) {
		goto L61
	} else {
		goto L62
	}
L6:
	;
	if v36 != int32(46) {
		v309 = v33
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if int32(3) <= v34-l1 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v46 = F_DecodeNumberField(m, l0, l1, l3|int32(14), l4, l5, l6, l7)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v34
	v55 = v34 + int32(1)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v56 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	return int32(0)
L12:
	;
	v309 = v46 >> (uint(int32(31)) % 32) & v46
	goto L2
L13:
	;
	v57 = int32(529931)
	v61 = m.G0
	v63 = v61 - int32(32)
	v64 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v63)+24)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v63)+16)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v63)+8)) = v64
	*(*int64)(unsafe.Add(mBase, uint32(v63))) = v64
	v72 = int32(*(*uint8)(unsafe.Add(mBase, _consts[255])))
	if v72 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v210 = float64(0)
	goto L15
L15:
	;
	v213 = base.F64_nearest(base.F64_mul(v210, float64(1e+06)))
	if base.F64_lt(base.F64_abs(v213), float64(2.147483648e+09)) != 0 {
		goto L58
	} else {
		goto L59
	}
L16:
	;
	if v55&int32(3) == int32(0) {
		v164 = v55
		goto L39
	} else {
		goto L40
	}
L17:
	;
	v140 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _consts[256])))
	if v76 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v80 = v55
	goto L23
L21:
	;
	goto L22
L22:
	;
	v90 = v57
	v91 = v72
	goto L26
L23:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80))))
	if v86 == v72 {
		v80 = v80 + int32(1)
		goto L23
	} else {
		goto L25
	}
L24:
	;
	v140 = v80 - v55
	goto L16
L25:
	;
	goto L24
L26:
	;
	v98 = v63 + int32(base.Ui32(v91)>>(uint(int32(3))%32))&int32(28)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v100 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = v99 | v100<<(uint(v91)%32)
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v90)+1)))
	if v104 != 0 {
		v90 = v90 + v100
		v91 = v104
		goto L26
	} else {
		goto L28
	}
L27:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v107 == int32(0) {
		v132 = v55
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L27
L29:
	;
	v140 = v132 - v55
	goto L16
L30:
	;
	v111 = v55
	v112 = v107
	goto L31
L31:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v63+int32(base.Ui32(v112)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v120)>>(uint(v112)%32))&int32(1) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v132 = v128
	goto L29
L33:
	;
	v132 = v111
	goto L29
L34:
	;
	goto L35
L35:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v111)+1)))
	v128 = v111 + int32(1)
	if v126 != 0 {
		v111 = v128
		v112 = v126
		goto L31
	} else {
		goto L36
	}
L36:
	;
	goto L32
L37:
	;
	if v140 != v197 {
		v309 = v33
		goto L2
	} else {
		goto L54
	}
L38:
	;
	v197 = v189 - v55
	goto L37
L39:
	;
	v168 = v164
	goto L48
L40:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v148 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v197 = int32(0)
	goto L37
L42:
	;
	goto L43
L43:
	;
	v153 = v55
	goto L44
L44:
	;
	v157 = v153 + int32(1)
	if v157&int32(3) == int32(0) {
		v164 = v157
		goto L39
	} else {
		goto L46
	}
L45:
	;
	v189 = v157
	goto L38
L46:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157))))
	if v162 != 0 {
		v153 = v157
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v168)))
	v177 = int32(-2139062144)
	if (int32(16843008)-v174|v174)&v177 == v177 {
		v168 = v168 + int32(4)
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v183 = v168
	goto L51
L50:
	;
	goto L49
L51:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v187 != 0 {
		v183 = v183 + int32(1)
		goto L51
	} else {
		goto L53
	}
L52:
	;
	v189 = v183
	goto L38
L53:
	;
	goto L52
L54:
	;
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v204 = F_strtod(m, v34, v17+int32(12))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L11
	} else {
		goto L55
	}
L55:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
	if v207 != 0 {
		v309 = v33
		goto L2
	} else {
		goto L56
	}
L56:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v209 != 0 {
		v309 = v33
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v210 = v204
	goto L15
L58:
	;
	v217 = base.I32_trunc_f64_s(v213)
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v217
	goto L5
L59:
	;
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = int32(-2147483648)
	goto L5
L61:
	;
	switch v224 - int32(1) {
	case 0, 2, 4, 6, 8, 10, 12:
		goto L68
	case 1:
		goto L73
	case 3:
		goto L74
	case 5:
		goto L72
	case 7:
		goto L71
	case 9:
		goto L70
	case 11:
		v309 = v33
		goto L2
	case 13:
		goto L69
	default:
		goto L75
	}
L62:
	;
	if v224 != int32(4) {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if base.Ui32(int32(365)) < base.Ui32(v27-int32(1)) {
		goto L61
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(32778)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+28)) = v27
	v309 = int32(0)
	goto L2
L65:
	;
	v300 = int32(0)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v301 != int32(4) {
		v309 = v300
		goto L2
	} else {
		goto L97
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5)+16)) = v27
	goto L65
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
	goto L66
L68:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L69:
	;
	v290 = F_DecodeNumberField(m, l0, l1, l3, l4, l5, l6, l7)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L11
	} else {
		goto L96
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v27
	goto L65
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
	goto L66
L72:
	;
	if l2 != 0 {
		goto L90
	} else {
		goto L91
	}
L73:
	;
	if l2 != 0 {
		goto L82
	} else {
		goto L83
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(2)
	goto L66
L75:
	;
	if l0 <= int32(2) {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	if v242 != int32(1) {
		goto L67
	} else {
		goto L81
	}
L77:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	if v242 != 0 {
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v27
	goto L65
L80:
	;
	goto L79
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L65
L82:
	;
	if l0 <= int32(2) {
		goto L86
	} else {
		goto L87
	}
L83:
	;
	goto L84
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L65
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L65
L86:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[514]))
	if v257 != 0 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v27
	goto L65
L89:
	;
	goto L88
L90:
	;
	if l0 < int32(3) {
		goto L93
	} else {
		goto L94
	}
L91:
	;
	goto L92
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L65
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v27
	goto L65
L94:
	;
	v269 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l7))))
	if v269 != int32(1) {
		goto L93
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(8)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l5)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l5)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(l5)+12)) = v274
	v277 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(v277)
	goto L65
L96:
	;
	v309 = v290 >> (uint(int32(31)) % 32) & v290
	goto L2
L97:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l7))) = uint8(base.B2i32(l0 < int32(3)))
	v309 = v300
	goto L2
}
func F_DeconstructQualifiedName(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if l0 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v106
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v108
	m.G0 = v10 + int32(32)
	return
L2:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
	v106 = v102
	v108 = v104
	goto L1
L3:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v100 = v98
	v102 = int32(0)
	goto L2
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L8
	} else {
		goto L24
	}
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v14 - int32(1) {
	case 0:
		goto L3
	case 1:
		goto L7
	case 2:
		goto L6
	default:
		goto L4
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v30 = *(*int32)(unsafe.Add(mBase, _consts[119]))
	v31 = F_get_database_name(m, v30)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	v100 = v17 + int32(4)
	v102 = v21
	goto L2
L8:
	;
	return
L9:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31))))
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v36 == int32(0) {
		v55 = v35
		v56 = v36
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v56-v55 == int32(0) {
		v106 = v26
		v108 = v24
		goto L1
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v35 != v36 {
		v55 = v35
		v56 = v36
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v40 = v28
	v41 = v31
	goto L14
L14:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v45 == int32(0) {
		v55 = v44
		v56 = v45
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v55 = v44
	v56 = v45
	goto L11
L16:
	;
	v48 = int32(1)
	if v44 == v45 {
		v40 = v40 + v48
		v41 = v41 + v48
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L8
	} else {
		goto L20
	}
L20:
	;
	v67 = F_NameListToString(m, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L8
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v67
	F_errmsg(m, int32(196182), v10+int32(16))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L8
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(482615), int32(3333), int32(368395))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
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
	F_errcode(m, int32(16801924))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L8
	} else {
		goto L25
	}
L25:
	;
	v87 = F_NameListToString(m, l0)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v87
	F_errmsg(m, int32(197127), v10)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L8
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(482615), int32(3339), int32(368395))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_DeescapeQuotedString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v177 int32
	_ = v177
	v2 = int32(0)
	if l0&int32(3) == v2 {
		v34 = l0
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v69 = v67 - int32(1)
	v70 = F_palloc(m, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L18
	} else {
		goto L19
	}
L2:
	;
	v67 = v59 - l0
	goto L1
L3:
	;
	v38 = v34
	goto L12
L4:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v18 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v67 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v23 = l0
	goto L8
L8:
	;
	v27 = v23 + int32(1)
	if v27&int32(3) == int32(0) {
		v34 = v27
		goto L3
	} else {
		goto L10
	}
L9:
	;
	v59 = v27
	goto L2
L10:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v32 != 0 {
		v23 = v27
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v47 = int32(-2139062144)
	if (int32(16843008)-v44|v44)&v47 == v47 {
		v38 = v38 + int32(4)
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v53 = v38
	goto L15
L14:
	;
	goto L13
L15:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
	if v57 != 0 {
		v53 = v53 + int32(1)
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v59 = v53
	goto L2
L17:
	;
	goto L16
L18:
	;
	return int32(0)
L19:
	;
	if int32(0) < v69 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v77 = l0 + int32(1)
	v79 = int32(0)
	v83 = v2
	goto L23
L21:
	;
	v168 = v2
	goto L22
L22:
	;
	v177 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v168+v70-int32(1)))) = uint8(v177)
	return v70
L23:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79+v77))))
	if v90 != int32(39) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v168 = v160
	goto L22
L25:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v83+v70))) = uint8(v152)
	v159 = int32(1)
	v160 = v83 + v159
	v162 = v154 + v159
	if v162 < v69 {
		v79 = v162
		v83 = v160
		goto L23
	} else {
		goto L44
	}
L26:
	;
	v152 = v90
	v154 = v79
	goto L25
L27:
	;
	if v90 != int32(92) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v143 = int32(39)
	v145 = v79 + int32(1)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77+v145))))
	if v147 == v143 {
		v152 = v143
		v154 = v145
		goto L25
	} else {
		goto L43
	}
L30:
	;
	v97 = v79 + int32(1)
	v98 = v77 + v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v101 = v99 - int32(48)
	switch v101 {
	case 0, 1, 2, 3, 4, 5, 6, 7:
		goto L32
	default:
		goto L31
	case 50:
		v152 = int32(8)
		v154 = v97
		goto L25
	case 54:
		goto L36
	case 62:
		goto L35
	case 66:
		goto L34
	case 68:
		goto L33
	}
L31:
	;
	v152 = v99
	v154 = v97
	goto L25
L32:
	;
	v106 = int32(0)
	if base.Ui32(int32(55)) < base.Ui32(v99) {
		v139 = v106
		v141 = v106
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v152 = int32(9)
	v154 = v97
	goto L25
L34:
	;
	v152 = int32(13)
	v154 = v97
	goto L25
L35:
	;
	v152 = int32(10)
	v154 = v97
	goto L25
L36:
	;
	v152 = int32(12)
	v154 = v97
	goto L25
L37:
	;
	v152 = v139
	v154 = v79 + v141
	goto L25
L38:
	;
	v111 = v101 & int32(255)
	v112 = int32(1)
	v113 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v113 < int32(48) {
		v139 = v111
		v141 = v112
		goto L37
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(55)) < base.Ui32(v113) {
		v139 = v111
		v141 = v112
		goto L37
	} else {
		goto L40
	}
L40:
	;
	v118 = int32(48)
	v124 = (v113-v118)&int32(255) + v111<<(uint(int32(3))%32)
	v125 = int32(2)
	v126 = int32(*(*int8)(unsafe.Add(mBase, uint32(v98)+2)))
	if v126 < v118 {
		v139 = v124
		v141 = v125
		goto L37
	} else {
		goto L41
	}
L41:
	;
	if base.Ui32(int32(55)) < base.Ui32(v126) {
		v139 = v124
		v141 = v125
		goto L37
	} else {
		goto L42
	}
L42:
	;
	v131 = int32(3)
	v139 = (v126-int32(48))&int32(255) + v124<<(uint(v131)%32)
	v141 = v131
	goto L37
L43:
	;
	goto L26
L44:
	;
	goto L24
}
func F_DelRoleMems(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int64
	_ = v193
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	v9 = int32(0)
	v20 = m.G0
	v22 = v20 + int32(-64)
	m.G0 = v22
	v25 = F_check_role_grantor(m, l0, l2, l5, v9)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = F_table_open(m, int32(1261), int32(3))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	F_LockSharedObject(m, int32(1260), l2, int32(4))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v38 = int32(0)
	v40 = F_SearchSysCacheList(m, int32(9), int32(1), l2, v38, v38)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L6
	}
L5:
	;
	v101 = v40 + int32(48)
	v111 = v9
	goto L13
L6:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v42 == int32(0) {
		v93 = v9
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v47 = F_palloc(m, v42<<(uint(int32(2))%32))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v49 <= int32(0) {
		v93 = v47
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v55 = int32(0)
	goto L10
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47+v55<<(uint(int32(2))%32)))) = int32(0)
	v78 = v55 + int32(1)
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v78 < v79 {
		v55 = v78
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v93 = v47
	goto L5
L12:
	;
	goto L11
L13:
	;
	v121 = int32(0)
	if l3 == v121 {
		v130 = v121
		goto L15
	} else {
		goto L16
	}
L15:
	;
	if l4 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v124 <= v111 {
		v130 = v121
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v130 = v126 + v111<<(uint(int32(2))%32)
	goto L15
L18:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if int32(0) < v282 {
		goto L49
	} else {
		goto L50
	}
L19:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if int32(0) < v142 {
		goto L24
	} else {
		goto L25
	}
L20:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v133 <= v111 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	if v130 == int32(0) {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v140 = v137 + v111<<(uint(int32(2))%32)
	if v140 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	v152 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	F_ReleaseCatCacheList(m, v40)
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L1
	} else {
		goto L47
	}
L27:
	;
	v170 = v152 << (uint(int32(2)) % 32)
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v93+v170)))
	if v172 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L26
L29:
	;
	v251 = v152 + int32(1)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v251 < v252 {
		v152 = v251
		goto L27
	} else {
		goto L46
	}
L30:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v101+v170)))
	if v172 == int32(4) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v176)+56))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180)+22)))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180+v181)))
	F_deleteSharedDependencyRecordsFor(m, int32(1261), v183, int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v191 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20+int32(-8)))) = v191
	v193 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20+int32(-16)))) = v193
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v193
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v193
	*(*int32)(unsafe.Add(mBase, uint32(v22)+27)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v22)+19)) = v191
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v191
	switch v172 - int32(1) {
	case 0:
		goto L37
	case 1:
		goto L40
	case 2:
		goto L39
	default:
		goto L38
	}
L34:
	;
	F_CatalogTupleDelete(m, v29, v176+int32(44))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	goto L29
L36:
	;
	v242 = F_heap_modify_tuple(m, v176+int32(40), v31, v20+int32(-32), v20+int32(-40), v20+int32(-48))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L1
	} else {
		goto L44
	}
L37:
	;
	v230 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+20)) = uint8(v230)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+48)) = int32(0)
	goto L36
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v213 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+22)) = uint8(v213)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+56)) = int32(0)
	goto L36
L40:
	;
	v209 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+21)) = uint8(v209)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+52)) = int32(0)
	goto L36
L41:
	;
	F_errmsg_internal(m, int32(249311), int32(0))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(478477), int32(2089), int32(143796))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	F_CatalogTupleUpdate(m, v29, v242+int32(4), v242)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	goto L29
L46:
	;
	goto L28
L47:
	;
	F_sequence_close(m, v29, int32(0))
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	m.G0 = v22 - int32(-64)
	return
L49:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
	v289 = int32(0)
	goto L52
L50:
	;
	goto L51
L51:
	;
	v362 = F_errstart(m, int32(19), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L65
	}
L52:
	;
	v307 = v289 << (uint(int32(2)) % 32)
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v101+v307)))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+56))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+22)))
	v312 = v310 + v311
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	if v313 != v285 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L51
L54:
	;
	v339 = v289 + int32(1)
	if v339 != v282 {
		v289 = v339
		goto L52
	} else {
		goto L64
	}
L55:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v312)+12))
	if v315 != v25 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v317&int32(2) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93+v307))) = int32(2)
	v111 = v111 + int32(1)
	goto L13
L58:
	;
	goto L59
L59:
	;
	if v317&int32(4) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93+v307))) = int32(3)
	v111 = v111 + int32(1)
	goto L13
L61:
	;
	goto L62
L62:
	;
	F_plan_recursive_revoke(m, v40, v93, v289, v317&int32(1), l7)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v111 = v111 + int32(1)
	goto L13
L64:
	;
	goto L53
L65:
	;
	if v362 != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v364 = F_get_rolespec_name(m, v281)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v111 = v111 + int32(1)
	goto L13
L69:
	;
	v367 = F_GetUserNameFromId(m, v25, int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v364
	F_errmsg(m, int32(680514), v22)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(478477), int32(2027), int32(143796))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L68
}
func F_DropPreparedStatement(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	v3 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = *(*int32)(unsafe.Add(mBase, _consts[476]))
	if v11 != 0 {
		v12 = int32(0)
		v14 = F_hash_search(m, v11, l0, v12, v12)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = v14
			if v16 == int32(0) {
				v19 = l1
			} else {
				v19 = v3
			}
			if v19 == int32(0) {
				if v16 != 0 {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
					F_DropCachedPlan(m, v22)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, _consts[476]))
						v29 = F_hash_search(m, v26, v16, int32(2), int32(0))
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				} else {
					m.G0 = v7 + int32(16)
					return
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errcode(m, int32(386))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
						F_errmsg(m, int32(68622), v7)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errfinish(m, int32(481881), int32(454), int32(91683))
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
	} else {
		v16 = v3
		if v16 == int32(0) {
			v19 = l1
		} else {
			v19 = v3
		}
		if v19 == int32(0) {
			if v16 != 0 {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+64))
				F_DropCachedPlan(m, v22)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _consts[476]))
					v29 = F_hash_search(m, v26, v16, int32(2), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						m.G0 = v7 + int32(16)
						return
					}
				}
			} else {
				m.G0 = v7 + int32(16)
				return
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_errcode(m, int32(386))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v7))) = l0
					F_errmsg(m, int32(68622), v7)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errfinish(m, int32(481881), int32(454), int32(91683))
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
}
func F_danish_ISO_8859_1_close_env(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_SN_close_env(m, l0, int32(1))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F_datadir_fsync_fname(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v21 = *(*int32)(unsafe.Add(mBase, _consts[293]))
	if v21 != 0 {
		v22 = F_GetCurrentTimestamp(m)
		mBase = m.M
		v24 = *(*int64)(unsafe.Add(mBase, _consts[294]))
		F_TimestampDifference(m, v24, v22, v18+int32(12), v18+int32(8))
		mBase = m.M
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v7+int32(28)))) = v30
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v7+int32(24)))) = v32
		*(*int32)(unsafe.Add(mBase, _consts[293])) = int32(0)
	} else {
	}
	m.G0 = v18 + int32(16)
	if base.B2i32(v21 != int32(0)) == int32(0) {
		v67 = F_fsync_fname_ext(m, l0, l1, int32(1), l2)
		mBase = m.M
		v68 = m.ExcPending
		if v68 != 0 {
			return
		} else {
			m.G0 = v7 + int32(32)
			return
		}
	} else {
		v47 = F_errstart(m, int32(15), int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			if v47 == int32(0) {
				v67 = F_fsync_fname_ext(m, l0, l1, int32(1), l2)
				mBase = m.M
				v68 = m.ExcPending
				if v68 != 0 {
					return
				} else {
					m.G0 = v7 + int32(32)
					return
				}
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v51
				*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l0
				v54 = *(*int32)(unsafe.Add(mBase, uint32(v7)+24))
				v56 = base.I32_div_s(v54, int32(10000))
				*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v56
				F_errmsg(m, int32(195039), v7)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					F_errfinish(m, int32(482819), int32(3832), int32(364110))
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v67 = F_fsync_fname_ext(m, l0, l1, int32(1), l2)
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
							return
						} else {
							m.G0 = v7 + int32(32)
							return
						}
					}
				}
			}
		}
	}
}
func F_date2timestamp_no_overflow(m *base.Module, l0 int32) float64 {
	if l0 == int32(-2147483648) {
		return float64(-1.7976931348623157e+308)
	} else {
		if l0 == int32(2147483647) {
			return float64(1.7976931348623157e+308)
		} else {
			return base.F64_mul(base.F64_convert_i32_s(l0), float64(8.64e+10))
		}
	}
}
func F_dcbrt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v13 int32
	_ = v13
	var v20 float64
	_ = v20
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v65 float64
	_ = v65
	var v67 float64
	_ = v67
	var v76 float64
	_ = v76
	var v80 float64
	_ = v80
	var v82 float64
	_ = v82
	var v90 float64
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v13 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v6))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(int32(2146435072)) <= base.Ui32(v13) {
		v80 = base.F64_add(v6, v6)
	} else {
		if base.Ui32(int32(1048575)) < base.Ui32(v13) {
			v30 = v13
			v31 = v6
			v32 = int32(715094163)
			v34 = base.I32_div_u_s(v30, int32(3))
			v40 = base.F64_copysign(base.F64_reinterpret_i64(base.I64_extend_i32_u(v32+v34)<<(uint(int64(32))%64)), v31)
			v43 = base.F64_mul(base.F64_mul(v40, v40), base.F64_div(v40, v6))
			v65 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_mul(v40, base.F64_add(base.F64_mul(base.F64_mul(v43, base.F64_mul(v43, v43)), base.F64_add(base.F64_mul(v43, float64(0.14599619288661245)), float64(-0.758397934778766))), base.F64_add(base.F64_mul(v43, base.F64_add(base.F64_mul(v43, float64(1.6214297201053545)), float64(-1.8849797954337717))), float64(1.87595182427177)))))&int64(-1073741824) + int64(2147483648))
			v67 = base.F64_div(v6, base.F64_mul(v65, v65))
			v76 = base.F64_add(base.F64_mul(v65, base.F64_div(base.F64_sub(v67, v65), base.F64_add(base.F64_add(v65, v65), v67))), v65)
		} else {
			v20 = base.F64_mul(v6, float64(1.8014398509481984e+16))
			v26 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v20))>>(uint(int64(32))%64))) & int32(2147483647)
			if v26 == int32(0) {
				v76 = v6
			} else {
				v30 = v26
				v31 = v20
				v32 = int32(696219795)
				v34 = base.I32_div_u_s(v30, int32(3))
				v40 = base.F64_copysign(base.F64_reinterpret_i64(base.I64_extend_i32_u(v32+v34)<<(uint(int64(32))%64)), v31)
				v43 = base.F64_mul(base.F64_mul(v40, v40), base.F64_div(v40, v6))
				v65 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_mul(v40, base.F64_add(base.F64_mul(base.F64_mul(v43, base.F64_mul(v43, v43)), base.F64_add(base.F64_mul(v43, float64(0.14599619288661245)), float64(-0.758397934778766))), base.F64_add(base.F64_mul(v43, base.F64_add(base.F64_mul(v43, float64(1.6214297201053545)), float64(-1.8849797954337717))), float64(1.87595182427177)))))&int64(-1073741824) + int64(2147483648))
				v67 = base.F64_div(v6, base.F64_mul(v65, v65))
				v76 = base.F64_add(base.F64_mul(v65, base.F64_div(base.F64_sub(v67, v65), base.F64_add(base.F64_add(v65, v65), v67))), v65)
			}
		}
		v80 = v76
	}
	v82 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v80), v82)&base.F64_ne(base.F64_abs(v6), v82) == int32(0) {
		v90 = float64(0)
		if base.F64_eq(v80, v90)&base.F64_ne(v6, v90) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v95 = F_Float8GetDatum(m, v80)
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return int32(0)
			} else {
				return v95
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v101 = m.ExcPending
		if v101 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_dcosd(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	var v21 int32
	_ = v21
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v41 float64
	_ = v41
	var v44 int64
	_ = v44
	var v51 float64
	_ = v51
	var v54 int32
	_ = v54
	var v56 int64
	_ = v56
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v72 int32
	_ = v72
	var v86 int32
	_ = v86
	var v91 int64
	_ = v91
	var v95 int32
	_ = v95
	var v97 int64
	_ = v97
	var v101 int64
	_ = v101
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int64
	_ = v119
	var v123 int64
	_ = v123
	var v130 int64
	_ = v130
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v140 int32
	_ = v140
	var v144 int64
	_ = v144
	var v146 int32
	_ = v146
	var v149 int64
	_ = v149
	var v163 int64
	_ = v163
	var v174 float64
	_ = v174
	var v178 float64
	_ = v178
	var v182 float64
	_ = v182
	var v186 float64
	_ = v186
	var v191 float64
	_ = v191
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v211 float64
	_ = v211
	var v215 int32
	_ = v215
	var v216 float64
	_ = v216
	var v217 float64
	_ = v217
	var v222 float64
	_ = v222
	var v224 float64
	_ = v224
	var v226 float64
	_ = v226
	var v229 float64
	_ = v229
	var v233 float64
	_ = v233
	var v237 float64
	_ = v237
	var v241 float64
	_ = v241
	var v250 float64
	_ = v250
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v270 float64
	_ = v270
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v276 float64
	_ = v276
	var v282 float64
	_ = v282
	var v283 float64
	_ = v283
	var v285 float64
	_ = v285
	var v287 float64
	_ = v287
	var v289 float64
	_ = v289
	var v298 float64
	_ = v298
	var v302 float64
	_ = v302
	var v309 float64
	_ = v309
	var v312 float64
	_ = v312
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui64(base.I64_reinterpret_f64(v11)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L92
	} else {
		goto L98
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L92
	} else {
		goto L94
	}
L3:
	;
	if base.F64_eq(base.F64_abs(v11), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	v312 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L5
L5:
	;
	v313 = F_Float8GetDatum(m, v312)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L92
	} else {
		goto L93
	}
L6:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _consts[958])))
	if v21 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_init_degree_constants(m)
	mBase = m.M
	goto L9
L8:
	;
	goto L9
L9:
	;
	v32 = base.I64_reinterpret_f64(v11)
	v36 = int32(2047)
	v37 = base.I32_wrap_i64(int64(base.Ui64(v32)>>(uint(int64(52))%64))) & v36
	if v37 != v36 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if base.F64_eq(base.F64_abs(v302), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L88
	}
L11:
	;
	if base.F64_lt(v174, float64(0)) != 0 {
		goto L52
	} else {
		goto L53
	}
L12:
	;
	v44 = v32 << (uint(int64(1)) % 64)
	if base.Ui64(v44) <= base.Ui64(int64(-9156662467374350336)) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v41 = base.F64_mul(v11, float64(360))
	v174 = base.F64_div(v41, v41)
	goto L11
L14:
	;
	if v44 == int64(-9156662467374350336) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	goto L16
L16:
	;
	if v37 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L17:
	;
	v51 = base.F64_mul(v11, float64(0))
	goto L19
L18:
	;
	v51 = v11
	goto L19
L19:
	;
	v174 = v51
	goto L11
L20:
	;
	if int32(1031) < v86 {
		goto L30
	} else {
		goto L31
	}
L21:
	;
	v54 = int32(0)
	v56 = v32 << (uint(int64(12)) % 64)
	if int64(0) <= v56 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v86 = v37
	v91 = v32&int64(4503599627370495) | int64(4503599627370496)
	goto L20
L24:
	;
	v60 = v54
	v62 = v56
	goto L27
L25:
	;
	v72 = v54
	goto L26
L26:
	;
	v86 = v72
	v91 = v32 << (uint(base.I64_extend_i32_u(int32(1)-v72)) % 64)
	goto L20
L27:
	;
	v66 = v60 - int32(1)
	v68 = v62 << (uint(int64(1)) % 64)
	if int64(0) <= v68 {
		v60 = v66
		v62 = v68
		goto L27
	} else {
		goto L29
	}
L28:
	;
	v72 = v66
	goto L26
L29:
	;
	goto L28
L30:
	;
	v95 = v86
	v97 = v91
	goto L33
L31:
	;
	v117 = v86
	v119 = v91
	goto L32
L32:
	;
	v123 = v119 - int64(6333186975989760)
	if v123 < int64(0) {
		v130 = v119
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v101 = v97 - int64(6333186975989760)
	if v101 < int64(0) {
		v108 = v97
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v117 = int32(1031)
	v119 = v110
	goto L32
L35:
	;
	v110 = v108 << (uint(int64(1)) % 64)
	v112 = v95 - int32(1)
	if int32(1031) < v112 {
		v95 = v112
		v97 = v110
		goto L33
	} else {
		goto L38
	}
L36:
	;
	if v101 != int64(0) {
		v108 = v101
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v174 = base.F64_mul(v11, float64(0))
	goto L11
L38:
	;
	goto L34
L39:
	;
	if base.Ui64(int64(4503599627370495)) < base.Ui64(v130) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	if v123 != int64(0) {
		v130 = v123
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v174 = base.F64_mul(v11, float64(0))
	goto L11
L42:
	;
	if int32(0) < v146 {
		goto L49
	} else {
		goto L50
	}
L43:
	;
	v146 = v117
	v149 = v130
	goto L42
L44:
	;
	goto L45
L45:
	;
	v134 = v117
	v136 = v130
	goto L46
L46:
	;
	v140 = v134 - int32(1)
	v144 = v136 << (uint(int64(1)) % 64)
	if base.Ui64(v136) < base.Ui64(int64(2251799813685248)) {
		v134 = v140
		v136 = v144
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v146 = v140
	v149 = v144
	goto L42
L48:
	;
	goto L47
L49:
	;
	v163 = v149 - int64(4503599627370496) | base.I64_extend_i32_u(v146)<<(uint(int64(52))%64)
	goto L51
L50:
	;
	v163 = int64(base.Ui64(v149) >> (uint(base.I64_extend_i32_u(int32(1)-v146)) % 64))
	goto L51
L51:
	;
	v174 = base.F64_reinterpret_i64(v163 | v32&int64(-9223372036854775807-1))
	goto L11
L52:
	;
	v178 = base.F64_neg(v174)
	goto L54
L53:
	;
	v178 = v174
	goto L54
L54:
	;
	if base.F64_gt(v178, float64(180)) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v182 = base.F64_sub(float64(360), v178)
	goto L57
L56:
	;
	v182 = v178
	goto L57
L57:
	;
	if base.F64_gt(v182, float64(90)) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v186 = base.F64_sub(float64(180), v182)
	goto L60
L59:
	;
	v186 = v182
	goto L60
L60:
	;
	if base.F64_le(v186, float64(60)) != 0 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v191 = base.F64_mul(v186, float64(0.017453292519943295))
	v195 = m.G0
	v197 = v195 - int32(16)
	m.G0 = v197
	v204 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v191))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v204) <= base.Ui32(int32(1072243195)) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	goto L63
L63:
	;
	v250 = base.F64_mul(base.F64_sub(float64(90), v186), float64(0.017453292519943295))
	v254 = m.G0
	v256 = v254 - int32(16)
	m.G0 = v256
	v263 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v250))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v263) <= base.Ui32(int32(1072243195)) {
		goto L77
	} else {
		goto L78
	}
L64:
	;
	v237 = base.F64_sub(float64(1), v233)
	*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v237
	v241 = *(*float64)(unsafe.Add(mBase, _consts[955]))
	v302 = base.F64_add(base.F64_mul(base.F64_div(v237, v241), float64(-0.5)), float64(1))
	goto L10
L65:
	;
	m.G0 = v197 + int32(16)
	goto L64
L66:
	;
	if base.Ui32(v204) < base.Ui32(int32(1044816030)) {
		v233 = float64(1)
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v204) {
		v233 = base.F64_sub(v191, v191)
		goto L65
	} else {
		goto L70
	}
L69:
	;
	v211 = F___cos(m, v191, float64(0))
	mBase = m.M
	v233 = v211
	goto L65
L70:
	;
	v215 = F___rem_pio2(m, v191, v197)
	mBase = m.M
	v216 = *(*float64)(unsafe.Add(mBase, uint32(v197)+8))
	v217 = *(*float64)(unsafe.Add(mBase, uint32(v197)))
	switch v215&int32(3) - int32(1) {
	case 0:
		goto L73
	case 1:
		goto L72
	case 2:
		goto L71
	default:
		goto L74
	}
L71:
	;
	v229 = F___sin(m, v217, v216, int32(1))
	mBase = m.M
	v233 = v229
	goto L65
L72:
	;
	v226 = F___cos(m, v217, v216)
	mBase = m.M
	v233 = base.F64_neg(v226)
	goto L65
L73:
	;
	v224 = F___sin(m, v217, v216, int32(1))
	mBase = m.M
	v233 = base.F64_neg(v224)
	goto L65
L74:
	;
	v222 = F___cos(m, v217, v216)
	mBase = m.M
	v233 = v222
	goto L65
L75:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = v289
	v298 = *(*float64)(unsafe.Add(mBase, _consts[953]))
	v302 = base.F64_mul(base.F64_div(v289, v298), float64(0.5))
	goto L10
L76:
	;
	m.G0 = v256 + int32(16)
	goto L75
L77:
	;
	if base.Ui32(v263) < base.Ui32(int32(1045430272)) {
		v289 = v250
		goto L76
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v263) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v270 = F___sin(m, v250, float64(0), int32(0))
	mBase = m.M
	v289 = v270
	goto L76
L81:
	;
	v289 = base.F64_sub(v250, v250)
	goto L76
L82:
	;
	goto L83
L83:
	;
	v274 = F___rem_pio2(m, v250, v256)
	mBase = m.M
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v256)+8))
	v276 = *(*float64)(unsafe.Add(mBase, uint32(v256)))
	switch v274&int32(3) - int32(1) {
	case 0:
		goto L86
	case 1:
		goto L85
	case 2:
		goto L84
	default:
		goto L87
	}
L84:
	;
	v287 = F___cos(m, v276, v275)
	mBase = m.M
	v289 = base.F64_neg(v287)
	goto L76
L85:
	;
	v285 = F___sin(m, v276, v275, int32(1))
	mBase = m.M
	v289 = base.F64_neg(v285)
	goto L76
L86:
	;
	v283 = F___cos(m, v276, v275)
	mBase = m.M
	v289 = v283
	goto L76
L87:
	;
	v282 = F___sin(m, v276, v275, int32(1))
	mBase = m.M
	v289 = v282
	goto L76
L88:
	;
	if base.F64_gt(v182, float64(90)) != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v309 = base.F64_neg(v302)
	goto L91
L90:
	;
	v309 = v302
	goto L91
L91:
	;
	v312 = v309
	goto L5
L92:
	;
	return int32(0)
L93:
	;
	m.G0 = v7 + int32(16)
	return v313
L94:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L92
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(388004), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L92
	} else {
		goto L96
	}
L96:
	;
	F_errfinish(m, int32(476769), int32(2334), int32(405075))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L92
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
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dcosh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v9 float64
	_ = v9
	var v10 int64
	_ = v10
	var v22 int64
	_ = v22
	var v27 int32
	_ = v27
	var v62 float64
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 float64
	_ = v69
	var v76 float64
	_ = v76
	var v77 int32
	_ = v77
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v85 float64
	_ = v85
	var v87 float64
	_ = v87
	var v88 int32
	_ = v88
	var v90 float64
	_ = v90
	var v91 float64
	_ = v91
	var v107 float64
	_ = v107
	var v110 float64
	_ = v110
	var v116 float64
	_ = v116
	var v125 float64
	_ = v125
	var v140 float64
	_ = v140
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v161 float64
	_ = v161
	var v164 float64
	_ = v164
	var v170 float64
	_ = v170
	var v180 float64
	_ = v180
	var v182 float64
	_ = v182
	var v194 float64
	_ = v194
	var v196 float64
	_ = v196
	var v197 float64
	_ = v197
	var v204 float64
	_ = v204
	var v211 float64
	_ = v211
	var v215 float64
	_ = v215
	var v220 float64
	_ = v220
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
	v9 = base.F64_abs(v5)
	v10 = base.I64_reinterpret_f64(v9)
	if base.Ui64(v10) <= base.Ui64(int64(4604418530035630079)) {
		if base.Ui64(v10) < base.Ui64(int64(4490088828488384512)) {
			v220 = float64(1)
		} else {
			v22 = base.I64_reinterpret_f64(v9)
			v27 = base.I32_wrap_i64(int64(base.Ui64(v22)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1078159482)) <= base.Ui32(v27) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v9)&int64(9223372036854775807)) {
					v182 = v9
					v194 = v182
				} else {
					if v22 < int64(0) {
						v194 = float64(-1)
					} else {
						if base.F64_gt(v9, float64(709.782712893384)) == int32(0) {
							v62 = base.F64_add(base.F64_mul(v9, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v9))
							if base.F64_lt(base.F64_abs(v62), float64(2.147483648e+09)) != 0 {
								v66 = base.I32_trunc_f64_s(v62)
								v68 = v66
							} else {
								v68 = int32(-2147483648)
							}
							v69 = base.F64_convert_i32_s(v68)
							v76 = base.F64_mul(v69, float64(1.9082149292705877e-10))
							v77 = v68
							v78 = base.F64_add(v9, base.F64_mul(v69, float64(-0.6931471803691238)))
							v79 = base.F64_sub(v78, v76)
							v85 = v79
							v87 = base.F64_sub(base.F64_sub(v78, v79), v76)
							v88 = v77
							v90 = base.F64_mul(v85, float64(0.5))
							v91 = base.F64_mul(v85, v90)
							v107 = base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v110 = base.F64_sub(float64(3), base.F64_mul(v107, v90))
							v116 = base.F64_mul(v91, base.F64_div(base.F64_sub(v107, v110), base.F64_sub(float64(6), base.F64_mul(v85, v110))))
							if v88 == int32(0) {
								v194 = base.F64_sub(v85, base.F64_sub(base.F64_mul(v85, v116), v91))
							} else {
								v125 = base.F64_sub(base.F64_sub(base.F64_mul(v85, base.F64_sub(v116, v87)), v87), v91)
								switch v88 + int32(1) {
								case 0:
									v194 = base.F64_add(base.F64_mul(base.F64_sub(v85, v125), float64(0.5)), float64(-0.5))
								default:
									v149 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v88+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v88) {
										v154 = base.F64_add(base.F64_sub(v85, v125), float64(1))
										if v88 == int32(1024) {
											v161 = base.F64_mul(base.F64_add(v154, v154), float64(8.98846567431158e+307))
										} else {
											v161 = base.F64_mul(v154, v149)
										}
										v194 = base.F64_add(v161, float64(-1))
									} else {
										v164 = float64(1)
										v170 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v88) << (uint(int64(52)) % 64))
										if base.Ui32(v88) <= base.Ui32(int32(19)) {
											v180 = base.F64_add(base.F64_sub(v164, v170), base.F64_sub(v85, v125))
										} else {
											v180 = base.F64_add(base.F64_sub(v85, base.F64_add(v125, v170)), v164)
										}
										v182 = base.F64_mul(v180, v149)
										v194 = v182
									}
								case 2:
									if base.F64_lt(v85, float64(-0.25)) != 0 {
										v194 = base.F64_mul(base.F64_sub(v125, base.F64_add(v85, float64(0.5))), float64(-2))
									} else {
										v140 = base.F64_sub(v85, v125)
										v194 = base.F64_add(base.F64_add(v140, v140), float64(1))
									}
								}
							}
						} else {
							v194 = base.F64_mul(v9, float64(8.98846567431158e+307))
						}
					}
				}
			} else {
				if base.Ui32(v27) < base.Ui32(int32(1071001155)) {
					if base.Ui32(v27) < base.Ui32(int32(1016070144)) {
						v182 = v9
						v194 = v182
					} else {
						v85 = v9
						v87 = float64(0)
						v88 = int32(0)
						v90 = base.F64_mul(v85, float64(0.5))
						v91 = base.F64_mul(v85, v90)
						v107 = base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v110 = base.F64_sub(float64(3), base.F64_mul(v107, v90))
						v116 = base.F64_mul(v91, base.F64_div(base.F64_sub(v107, v110), base.F64_sub(float64(6), base.F64_mul(v85, v110))))
						if v88 == int32(0) {
							v194 = base.F64_sub(v85, base.F64_sub(base.F64_mul(v85, v116), v91))
						} else {
							v125 = base.F64_sub(base.F64_sub(base.F64_mul(v85, base.F64_sub(v116, v87)), v87), v91)
							switch v88 + int32(1) {
							case 0:
								v194 = base.F64_add(base.F64_mul(base.F64_sub(v85, v125), float64(0.5)), float64(-0.5))
							default:
								v149 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v88+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v88) {
									v154 = base.F64_add(base.F64_sub(v85, v125), float64(1))
									if v88 == int32(1024) {
										v161 = base.F64_mul(base.F64_add(v154, v154), float64(8.98846567431158e+307))
									} else {
										v161 = base.F64_mul(v154, v149)
									}
									v194 = base.F64_add(v161, float64(-1))
								} else {
									v164 = float64(1)
									v170 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v88) << (uint(int64(52)) % 64))
									if base.Ui32(v88) <= base.Ui32(int32(19)) {
										v180 = base.F64_add(base.F64_sub(v164, v170), base.F64_sub(v85, v125))
									} else {
										v180 = base.F64_add(base.F64_sub(v85, base.F64_add(v125, v170)), v164)
									}
									v182 = base.F64_mul(v180, v149)
									v194 = v182
								}
							case 2:
								if base.F64_lt(v85, float64(-0.25)) != 0 {
									v194 = base.F64_mul(base.F64_sub(v125, base.F64_add(v85, float64(0.5))), float64(-2))
								} else {
									v140 = base.F64_sub(v85, v125)
									v194 = base.F64_add(base.F64_add(v140, v140), float64(1))
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(1072734897)) < base.Ui32(v27) {
						v62 = base.F64_add(base.F64_mul(v9, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v9))
						if base.F64_lt(base.F64_abs(v62), float64(2.147483648e+09)) != 0 {
							v66 = base.I32_trunc_f64_s(v62)
							v68 = v66
						} else {
							v68 = int32(-2147483648)
						}
						v69 = base.F64_convert_i32_s(v68)
						v76 = base.F64_mul(v69, float64(1.9082149292705877e-10))
						v77 = v68
						v78 = base.F64_add(v9, base.F64_mul(v69, float64(-0.6931471803691238)))
					} else {
						if int64(0) <= v22 {
							v76 = float64(1.9082149292705877e-10)
							v77 = int32(1)
							v78 = base.F64_add(v9, float64(-0.6931471803691238))
						} else {
							v76 = float64(-1.9082149292705877e-10)
							v77 = int32(-1)
							v78 = base.F64_add(v9, float64(0.6931471803691238))
						}
					}
					v79 = base.F64_sub(v78, v76)
					v85 = v79
					v87 = base.F64_sub(base.F64_sub(v78, v79), v76)
					v88 = v77
					v90 = base.F64_mul(v85, float64(0.5))
					v91 = base.F64_mul(v85, v90)
					v107 = base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, base.F64_add(base.F64_mul(v91, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v110 = base.F64_sub(float64(3), base.F64_mul(v107, v90))
					v116 = base.F64_mul(v91, base.F64_div(base.F64_sub(v107, v110), base.F64_sub(float64(6), base.F64_mul(v85, v110))))
					if v88 == int32(0) {
						v194 = base.F64_sub(v85, base.F64_sub(base.F64_mul(v85, v116), v91))
					} else {
						v125 = base.F64_sub(base.F64_sub(base.F64_mul(v85, base.F64_sub(v116, v87)), v87), v91)
						switch v88 + int32(1) {
						case 0:
							v194 = base.F64_add(base.F64_mul(base.F64_sub(v85, v125), float64(0.5)), float64(-0.5))
						default:
							v149 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v88+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v88) {
								v154 = base.F64_add(base.F64_sub(v85, v125), float64(1))
								if v88 == int32(1024) {
									v161 = base.F64_mul(base.F64_add(v154, v154), float64(8.98846567431158e+307))
								} else {
									v161 = base.F64_mul(v154, v149)
								}
								v194 = base.F64_add(v161, float64(-1))
							} else {
								v164 = float64(1)
								v170 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v88) << (uint(int64(52)) % 64))
								if base.Ui32(v88) <= base.Ui32(int32(19)) {
									v180 = base.F64_add(base.F64_sub(v164, v170), base.F64_sub(v85, v125))
								} else {
									v180 = base.F64_add(base.F64_sub(v85, base.F64_add(v125, v170)), v164)
								}
								v182 = base.F64_mul(v180, v149)
								v194 = v182
							}
						case 2:
							if base.F64_lt(v85, float64(-0.25)) != 0 {
								v194 = base.F64_mul(base.F64_sub(v125, base.F64_add(v85, float64(0.5))), float64(-2))
							} else {
								v140 = base.F64_sub(v85, v125)
								v194 = base.F64_add(base.F64_add(v140, v140), float64(1))
							}
						}
					}
				}
			}
			v196 = float64(1)
			v197 = base.F64_add(v194, v196)
			v220 = base.F64_add(base.F64_div(base.F64_mul(v194, v194), base.F64_add(v197, v197)), v196)
		}
	} else {
		if base.Ui64(v10) <= base.Ui64(int64(4649454526309335039)) {
			v204 = F_exp(m, v9)
			mBase = m.M
			v220 = base.F64_mul(base.F64_add(v204, base.F64_div(float64(1), v204)), float64(0.5))
		} else {
			v211 = float64(2.247116418577895e+307)
			v215 = F_exp(m, base.F64_add(v9, float64(-1416.0996898839683)))
			mBase = m.M
			v220 = base.F64_mul(base.F64_mul(base.F64_mul(float64(1), v211), v215), v211)
		}
	}
	if base.F64_eq(v220, float64(0)) != 0 {
		F_float_underflow_error(m)
		mBase = m.M
		v226 = m.ExcPending
		if v226 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v227 = F_Float8GetDatum(m, v220)
		mBase = m.M
		v228 = m.ExcPending
		if v228 != 0 {
			return int32(0)
		} else {
			return v227
		}
	}
}
func F_debackslash(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v8 = F_palloc(m, l1+int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if int32(0) < l1 {
			v14 = l0
			v15 = l1
			v16 = v8
			for {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				v22 = int32(1)
				v24 = base.B2i32(v19 == int32(92)) & base.B2i32(v15 != v22)
				v25 = v14 + v24
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
				*(*uint8)(unsafe.Add(mBase, uint32(v16))) = uint8(v26)
				v31 = v16 + v22
				v34 = v24 ^ int32(-1) + v15
				if int32(0) < v34 {
					v14 = v25 + v22
					v15 = v34
					v16 = v31
					continue
				} else {
					break
				}
				break
			}
			v39 = v31
		} else {
			v39 = v8
		}
		v42 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v39))) = uint8(v42)
		return v8
	}
}
func F_decrypt_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
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
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v271 int32
	_ = v271
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v297 int32
	_ = v297
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v448 int32
	_ = v448
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v511 int64
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v554 int32
	_ = v554
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	v11 = m.G0
	v13 = v11 - int32(208)
	m.G0 = v13
	*(*int32)(unsafe.Add(mBase, uint32(v13)+196)) = int32(0)
	F_init_work(m, v13+int32(196), l1, l5, v13+int32(152))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v25 = int32(1)
	v26 = l2 + v25
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	v31 = v29 & v25
	if v31 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v32 = v26
	goto L5
L4:
	;
	v32 = l2 + int32(4)
	goto L5
L5:
	;
	if v29 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v61 = F_mbuf_create_from_data(m, v32, v60)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v35 = int32(4)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v37&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v50 = int32(1)
	if v31 != 0 {
		v60 = int32(base.Ui32(v29)>>(uint(v50)%32)) - v50
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v46 = v35
	goto L12
L11:
	;
	v46 = base.B2i32(v37 == int32(18)) << (uint(v35) % 32)
	goto L12
L12:
	;
	if v37 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v49 = v35
	goto L15
L14:
	;
	v49 = v46
	goto L15
L15:
	;
	v60 = v49
	goto L6
L16:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v60 = int32(base.Ui32(v54)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2))))
	if v63 == int32(1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v93 = F_mbuf_create(m, v90+int32(2048))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L32
	}
L19:
	;
	v66 = int32(6)
	v68 = int32(18)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v70 == v68 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v82 = int32(1)
	if v63&v82 != 0 {
		v90 = int32(base.Ui32(v63) >> (uint(v82) % 32))
		goto L18
	} else {
		goto L31
	}
L22:
	;
	v73 = v68
	goto L24
L23:
	;
	v73 = int32(2)
	goto L24
L24:
	;
	if v70&int32(254) == int32(2) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = v66
	goto L27
L26:
	;
	v78 = v73
	goto L27
L27:
	;
	if v70 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v81 = v66
	goto L30
L29:
	;
	v81 = v78
	goto L30
L30:
	;
	v90 = v81
	goto L18
L31:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v90 = int32(base.Ui32(v86) >> (uint(int32(2)) % 32))
	goto L18
L32:
	;
	v98 = F_mbuf_append(m, v93, v13+int32(204), int32(4))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if l0 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	v241 = int32(0)
	if v241 <= v236 {
		goto L92
	} else {
		goto L93
	}
L35:
	;
	if l4 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	goto L37
L37:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v189 = int32(1)
	v190 = l3 + v189
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v195 = v193 & v189
	if v195 != 0 {
		goto L74
	} else {
		goto L75
	}
L38:
	;
	v144 = int32(1)
	v145 = l3 + v144
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l3))))
	v150 = v148 & v144
	if v150 != 0 {
		goto L57
	} else {
		goto L58
	}
L39:
	;
	v102 = int32(0)
	v139 = v102
	v143 = v102
	goto L38
L40:
	;
	goto L41
L41:
	;
	v105 = l4 + int32(4)
	v106 = int32(1)
	v107 = l4 + v106
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	v110 = v108 & v106
	if v108 == v106 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	if v110 != 0 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L44
L44:
	;
	if v110 != 0 {
		goto L54
	} else {
		goto L55
	}
L45:
	;
	v113 = v107
	goto L47
L46:
	;
	v113 = v105
	goto L47
L47:
	;
	v114 = int32(4)
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v116&int32(254) == int32(2) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v125 = v114
	goto L50
L49:
	;
	v125 = base.B2i32(v116 == int32(18)) << (uint(v114) % 32)
	goto L50
L50:
	;
	if v116 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v128 = v114
	goto L53
L52:
	;
	v128 = v125
	goto L53
L53:
	;
	v139 = v113
	v143 = v128
	goto L38
L54:
	;
	v129 = int32(1)
	v139 = v107
	v143 = int32(base.Ui32(v108)>>(uint(v129)%32)) - v129
	goto L38
L55:
	;
	goto L56
L56:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v139 = v105
	v143 = int32(base.Ui32(v133)>>(uint(int32(2))%32)) - int32(4)
	goto L38
L57:
	;
	v151 = v145
	goto L59
L58:
	;
	v151 = l3 + int32(4)
	goto L59
L59:
	;
	if v148 == int32(1) {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v180 = F_mbuf_create_from_data(m, v151, v179)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L71
	}
L61:
	;
	v154 = int32(4)
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v156&int32(254) == int32(2) {
		goto L64
	} else {
		goto L65
	}
L62:
	;
	goto L63
L63:
	;
	v169 = int32(1)
	if v150 != 0 {
		v179 = int32(base.Ui32(v148)>>(uint(v169)%32)) - v169
		goto L60
	} else {
		goto L70
	}
L64:
	;
	v165 = v154
	goto L66
L65:
	;
	v165 = base.B2i32(v156 == int32(18)) << (uint(v154) % 32)
	goto L66
L66:
	;
	if v156 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v168 = v154
	goto L69
L68:
	;
	v168 = v165
	goto L69
L69:
	;
	v179 = v168
	goto L60
L70:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v179 = int32(base.Ui32(v173)>>(uint(int32(2))%32)) - int32(4)
	goto L60
L71:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v184 = F_pgp_set_pubkey(m, v182, v180, v139, v143, int32(1))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v186 = F_mbuf_free(m, v180)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v236 = v184
	goto L34
L74:
	;
	v196 = v190
	goto L76
L75:
	;
	v196 = l3 + int32(4)
	goto L76
L76:
	;
	if v193 == int32(1) {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v226 = int32(-13)
	if v196 == int32(0) {
		v234 = v226
		goto L89
	} else {
		goto L90
	}
L78:
	;
	v199 = int32(4)
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v201&int32(254) == int32(2) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	goto L80
L80:
	;
	v214 = int32(1)
	if v195 != 0 {
		v224 = int32(base.Ui32(v193)>>(uint(v214)%32)) - v214
		goto L77
	} else {
		goto L87
	}
L81:
	;
	v210 = v199
	goto L83
L82:
	;
	v210 = base.B2i32(v201 == int32(18)) << (uint(v199) % 32)
	goto L83
L83:
	;
	if v201 == int32(1) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v213 = v199
	goto L86
L85:
	;
	v213 = v210
	goto L86
L86:
	;
	v224 = v213
	goto L77
L87:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v224 = int32(base.Ui32(v218)>>(uint(int32(2))%32)) - int32(4)
	goto L77
L88:
	;
	v236 = v234
	goto L34
L89:
	;
	goto L88
L90:
	;
	if v224 <= int32(0) {
		v234 = v226
		goto L89
	} else {
		goto L91
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188)+128)) = v224
	*(*int32)(unsafe.Add(mBase, uint32(v188)+124)) = v196
	v234 = int32(0)
	goto L89
L92:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v245 = F_pgp_decrypt(m, v244, v61, v93)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	v493 = v236
	v494 = v241
	goto L94
L94:
	;
	v496 = F_mbuf_free(m, v61)
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L161
	}
L95:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v13)+156))
	if v247 == int32(0) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v489)+88))
	v493 = v245
	v494 = base.B2i32(v490 != int32(0))
	goto L94
L97:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v13)+160))
	if v251 < int32(0) {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v13)+164))
	if v277 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L99:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v250)+60))
	if v251 == v254 {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v258 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L101
	}
L101:
	;
	if v258 == int32(0) {
		goto L98
	} else {
		goto L102
	}
L102:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v250)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+136)) = v262
	*(*int32)(unsafe.Add(mBase, uint32(v13)+132)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v13)+128)) = int32(232163)
	F_errmsg(m, int32(451964), v13+int32(128))
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(480439), int32(149), int32(104335))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	goto L98
L105:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v13)+168))
	if v303 < int32(0) {
		goto L112
	} else {
		goto L113
	}
L106:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v250)+44))
	if v277 == v280 {
		goto L105
	} else {
		goto L107
	}
L107:
	;
	v284 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	if v284 == int32(0) {
		goto L105
	} else {
		goto L109
	}
L109:
	;
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v250)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+120)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v13)+116)) = v277
	*(*int32)(unsafe.Add(mBase, uint32(v13)+112)) = int32(397158)
	F_errmsg(m, int32(451964), v13+int32(112))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L110
	}
L110:
	;
	F_errfinish(m, int32(480439), int32(150), int32(104335))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	goto L105
L112:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v13)+176))
	if v329 < int32(0) {
		goto L119
	} else {
		goto L120
	}
L113:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v250)+48))
	if v303 == v306 {
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v310 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L115
	}
L115:
	;
	if v310 == int32(0) {
		goto L112
	} else {
		goto L116
	}
L116:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v250)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+104)) = v314
	*(*int32)(unsafe.Add(mBase, uint32(v13)+100)) = v303
	*(*int32)(unsafe.Add(mBase, uint32(v13)+96)) = int32(84535)
	F_errmsg(m, int32(451964), v13+int32(96))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(480439), int32(151), int32(104335))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	goto L112
L119:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v13)+184))
	if v355 < int32(0) {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v250)+52))
	if v329 == v332 {
		goto L119
	} else {
		goto L121
	}
L121:
	;
	v336 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	if v336 == int32(0) {
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v250)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+88)) = v340
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v329
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = int32(232093)
	F_errmsg(m, int32(451964), v13+int32(80))
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errfinish(m, int32(480439), int32(152), int32(104335))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	goto L119
L126:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v250)+76))
	if v381 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L127:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v250)+76))
	if v355 == v358 {
		goto L126
	} else {
		goto L128
	}
L128:
	;
	v362 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	if v362 == int32(0) {
		goto L126
	} else {
		goto L130
	}
L130:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v250)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v13)+68)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v13)+64)) = int32(19976)
	F_errmsg(m, int32(451964), v13-int32(-64))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(480439), int32(153), int32(104335))
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L1
	} else {
		goto L132
	}
L132:
	;
	goto L126
L133:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v13)+188))
	if v411 < int32(0) {
		goto L141
	} else {
		goto L142
	}
L134:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v13)+172))
	if v384 < int32(0) {
		goto L133
	} else {
		goto L135
	}
L135:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v250)+56))
	if v384 == v387 {
		goto L133
	} else {
		goto L136
	}
L136:
	;
	v391 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	if v391 == int32(0) {
		goto L133
	} else {
		goto L138
	}
L138:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v250)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v395
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = int32(232159)
	F_errmsg(m, int32(451964), v13+int32(48))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(480439), int32(155), int32(104335))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	goto L133
L141:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v13)+180))
	if v437 < int32(0) {
		goto L148
	} else {
		goto L149
	}
L142:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v250)+72))
	if v411 == v414 {
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v418 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	if v418 == int32(0) {
		goto L141
	} else {
		goto L145
	}
L145:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v250)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v422
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v411
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(475208)
	F_errmsg(m, int32(451964), v13+int32(32))
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(480439), int32(156), int32(104335))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	goto L141
L148:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v13)+192))
	if v463 < int32(0) {
		goto L96
	} else {
		goto L155
	}
L149:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v250)+64))
	if v437 == v440 {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v444 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	if v444 == int32(0) {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v250)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(232117)
	F_errmsg(m, int32(451964), v13+int32(16))
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	F_errfinish(m, int32(480439), int32(157), int32(104335))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	goto L148
L155:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v250)+88))
	if v463 == v466 {
		goto L96
	} else {
		goto L156
	}
L156:
	;
	v470 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L157
	}
L157:
	;
	if v470 == int32(0) {
		goto L96
	} else {
		goto L158
	}
L158:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v250)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v474
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v463
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(397239)
	F_errmsg(m, int32(451964), v13)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L159
	}
L159:
	;
	F_errfinish(m, int32(480439), int32(158), int32(104335))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L160
	}
L160:
	;
	goto L96
L161:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v13)+196))
	v499 = F_pgp_free(m, v498)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L162
	}
L162:
	;
	if v493 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v506 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v93)+16)) = uint16(v506)
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	*(*int32)(unsafe.Add(mBase, uint32(v13+int32(200)))) = v509
	v511 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v93)+8)) = v511
	*(*int64)(unsafe.Add(mBase, uint32(v93))) = v511
	goto L166
L164:
	;
	goto L165
L165:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1305])) = int32(0)
	goto L213
L166:
	;
	v516 = F_mbuf_free(m, v93)
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L1
	} else {
		goto L167
	}
L167:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v13)+200))
	*(*int32)(unsafe.Add(mBase, uint32(v518))) = (v508 - v509) << (uint(int32(2)) % 32)
	v522 = int32(0)
	if v494&base.B2i32(l1 != v522) == v522 {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1305])) = int32(0)
	goto L212
L169:
	;
	v609 = v518
	goto L168
L170:
	;
	goto L171
L171:
	;
	v528 = *(*int32)(unsafe.Add(mBase, _consts[451]))
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v528)+4))
	goto L172
L172:
	;
	v530 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	if v530 == int32(1) {
		goto L174
	} else {
		goto L175
	}
L173:
	;
	v561 = int32(1)
	if v530&v561 != 0 {
		goto L184
	} else {
		goto L185
	}
L174:
	;
	v533 = int32(4)
	v535 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+1)))
	if v535&int32(254) == int32(2) {
		goto L177
	} else {
		goto L178
	}
L175:
	;
	goto L176
L176:
	;
	v548 = int32(1)
	if v530&v548 != 0 {
		v560 = int32(base.Ui32(v530)>>(uint(v548)%32)) - v548
		goto L173
	} else {
		goto L183
	}
L177:
	;
	v544 = v533
	goto L179
L178:
	;
	v544 = base.B2i32(v535 == int32(18)) << (uint(v533) % 32)
	goto L179
L179:
	;
	if v535 == int32(1) {
		goto L180
	} else {
		goto L181
	}
L180:
	;
	v547 = v533
	goto L182
L181:
	;
	v547 = v544
	goto L182
L182:
	;
	v560 = v547
	goto L173
L183:
	;
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v560 = int32(base.Ui32(v554)>>(uint(int32(2))%32)) - int32(4)
	goto L173
L184:
	;
	v565 = v561
	goto L186
L185:
	;
	v565 = int32(4)
	goto L186
L186:
	;
	v566 = v518 + v565
	v568 = F_pg_do_encoding_conversion(m, v566, v560, int32(6), v529)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L1
	} else {
		goto L187
	}
L187:
	;
	if v568 == v566 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v609 = v518
	goto L168
L189:
	;
	goto L190
L190:
	;
	v571 = F_cstring_to_text(m, v568)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	F_pfree(m, v568)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L192
	}
L192:
	;
	if v518 == v571 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v609 = v518
	goto L168
L194:
	;
	goto L195
L195:
	;
	v577 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518))))
	if v577 == int32(1) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v605 = F___memset(m, v518, int32(0), v604)
	mBase = m.M
	goto L210
L197:
	;
	v580 = int32(6)
	v582 = int32(18)
	v584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518)+1)))
	if v584 == v582 {
		goto L200
	} else {
		goto L201
	}
L198:
	;
	goto L199
L199:
	;
	v596 = int32(1)
	if v577&v596 != 0 {
		v604 = int32(base.Ui32(v577) >> (uint(v596) % 32))
		goto L196
	} else {
		goto L209
	}
L200:
	;
	v587 = v582
	goto L202
L201:
	;
	v587 = int32(2)
	goto L202
L202:
	;
	if v584&int32(254) == int32(2) {
		goto L203
	} else {
		goto L204
	}
L203:
	;
	v592 = v580
	goto L205
L204:
	;
	v592 = v587
	goto L205
L205:
	;
	if v584 == int32(1) {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v595 = v580
	goto L208
L207:
	;
	v595 = v592
	goto L208
L208:
	;
	v604 = v595
	goto L196
L209:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v518)))
	v604 = int32(base.Ui32(v600) >> (uint(int32(2)) % 32))
	goto L196
L210:
	;
	F_pfree(m, v518)
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v609 = v571
	goto L168
L212:
	;
	m.G0 = v13 + int32(208)
	return v609
L213:
	;
	v621 = F_mbuf_free(m, v93)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	F_px_THROW_ERROR(m, v493)
	mBase = m.M
	v624 = m.ExcPending
	if v624 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_deparse_expression(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int64
	_ = v19
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	v3 = l2
	v6 = m.G0
	v8 = v6 + int32(-64)
	m.G0 = v8
	F_initStringInfo(m, v6+int32(-16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+40)) = uint8(v3)
		v17 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v17
		v19 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v17
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+43)) = uint8(v17)
		v26 = int32(1)
		*(*uint16)(unsafe.Add(mBase, uint32(v8)+41)) = uint16(v26)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v17
		*(*int64)(unsafe.Add(mBase, uint32(v8)+28)) = v19
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v6 + int32(-16)
		F_get_rule_expr(m, l0, v6+int32(-56), l3)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v8)+48))
			m.G0 = v8 - int32(-64)
			return v39
		}
	}
}
func F_deparse_lquery(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
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
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v337 int32
	_ = v337
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v15 = l0 + int32(16)
	v16 = int32(1)
	v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = v15
	v20 = v16
	v22 = int32(0)
	goto L4
L2:
	;
	v55 = v16
	goto L3
L3:
	;
	v62 = F_palloc(m, v55)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L15
	}
L4:
	;
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+4)))
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v55 = v44
	goto L3
L6:
	;
	v51 = v22 + int32(1)
	if v51 != v17 {
		v19 = v19 + (v43+int32(7))&int32(131064)
		v20 = v44
		v22 = v51
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v28 = int32(2)
	v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19))))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+2)))
	if v35&int32(32) != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19))))
	v43 = v40
	v44 = v20 + int32(27)
	goto L6
L10:
	;
	v38 = int32(27)
	goto L12
L11:
	;
	v38 = v28
	goto L12
L12:
	;
	v43 = v31
	v44 = v20 + v27<<(uint(v28)%32) + v31 + v38
	goto L6
L13:
	;
	goto L5
L14:
	;
	return int32(0)
L15:
	;
	v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if v66 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v69 = v62
	v71 = v15
	v75 = int32(0)
	goto L19
L17:
	;
	v329 = v62
	goto L18
L18:
	;
	v337 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v329))) = uint8(v337)
	m.G0 = v12 - int32(-64)
	return v62
L19:
	;
	if v75 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v329 = v315
	goto L18
L21:
	;
	v77 = int32(46)
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v77)
	v81 = v69 + int32(1)
	goto L23
L22:
	;
	v81 = v69
	goto L23
L23:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v82 != 0 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
	if v210&int32(32) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L25:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+2)))
	if v83&int32(16) != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v197 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v197)
	v202 = v81 + int32(1)
	goto L24
L28:
	;
	v86 = int32(33)
	*(*uint8)(unsafe.Add(mBase, uint32(v81))) = uint8(v86)
	v89 = v81 + int32(1)
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v90 == int32(0) {
		v202 = v89
		goto L24
	} else {
		goto L31
	}
L29:
	;
	v93 = v81
	goto L30
L30:
	;
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+20)))
	if v96 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v93 = v89
	goto L30
L32:
	;
	v99 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+20)))
	v100 = v98 + v99
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	if v101&int32(4) != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v97 = F__emscripten_memcpy_bulkmem(m, v93, v71+int32(23), v96)
	mBase = m.M
	v98 = v97
	goto L35
L34:
	;
	v98 = v93
	goto L35
L35:
	;
	goto L32
L36:
	;
	v104 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v104)
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	v109 = v100 + int32(1)
	v110 = v106
	goto L38
L37:
	;
	v109 = v100
	v110 = v101
	goto L38
L38:
	;
	if v110&int32(2) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v113 = int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v109))) = uint8(v113)
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+22)))
	v118 = v109 + int32(1)
	v119 = v115
	goto L41
L40:
	;
	v118 = v109
	v119 = v110
	goto L41
L41:
	;
	if v119&int32(1) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v122 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v118))) = uint8(v122)
	v126 = v118 + int32(1)
	goto L44
L43:
	;
	v126 = v118
	goto L44
L44:
	;
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
	if base.Ui32(v127) < base.Ui32(int32(2)) {
		v202 = v126
		goto L24
	} else {
		goto L45
	}
L45:
	;
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+20)))
	v140 = v126
	v141 = v71 + int32(16) + (v132+int32(7))&int32(131064)
	v145 = int32(1)
	goto L46
L46:
	;
	v148 = int32(124)
	*(*uint8)(unsafe.Add(mBase, uint32(v140))) = uint8(v148)
	v151 = v140 + int32(1)
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+12)))
	if v154 != 0 {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v202 = v184
	goto L24
L48:
	;
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+12)))
	v158 = v156 + v157
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+14)))
	if v159&int32(4) != 0 {
		goto L52
	} else {
		goto L53
	}
L49:
	;
	v155 = F__emscripten_memcpy_bulkmem(m, v151, v141+int32(15), v154)
	mBase = m.M
	v156 = v155
	goto L51
L50:
	;
	v156 = v151
	goto L51
L51:
	;
	goto L48
L52:
	;
	v162 = int32(37)
	*(*uint8)(unsafe.Add(mBase, uint32(v158))) = uint8(v162)
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+14)))
	v167 = v158 + int32(1)
	v168 = v164
	goto L54
L53:
	;
	v167 = v158
	v168 = v159
	goto L54
L54:
	;
	if v168&int32(2) != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v171 = int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v167))) = uint8(v171)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+14)))
	v176 = v167 + int32(1)
	v177 = v173
	goto L57
L56:
	;
	v176 = v167
	v177 = v168
	goto L57
L57:
	;
	if v177&int32(1) != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v180 = int32(42)
	*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v180)
	v184 = v176 + int32(1)
	goto L60
L59:
	;
	v184 = v176
	goto L60
L60:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+12)))
	v194 = v145 + int32(1)
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
	if base.Ui32(v194) < base.Ui32(v195) {
		v140 = v184
		v141 = v141 + int32(8) + (v187+int32(7))&int32(131064)
		v145 = v194
		goto L46
	} else {
		goto L61
	}
L61:
	;
	goto L47
L62:
	;
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71))))
	v325 = v75 + int32(1)
	v326 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	if base.Ui32(v325) < base.Ui32(v326) {
		v69 = v315
		v71 = v71 + (v318+int32(7))&int32(131064)
		v75 = v325
		goto L19
	} else {
		goto L105
	}
L63:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v215 != 0 {
		v315 = v202
		goto L62
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+6)))
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+8)))
	if v216 == v217 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L65
L67:
	;
	if v202&int32(3) == int32(0) {
		v280 = v202
		goto L90
	} else {
		goto L91
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v216
	v221 = F_pg_sprintf(m, v202, int32(6667), v12)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L14
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	if v216 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	goto L67
L72:
	;
	if v217 == int32(65535) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	if v217 == int32(65535) {
		goto L83
	} else {
		goto L84
	}
L75:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
	if v227 == int32(0) {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v217
	v240 = F_pg_sprintf(m, v202, int32(6672), v10+int32(-48))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L14
	} else {
		goto L82
	}
L78:
	;
	v230 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v202))) = uint8(v230)
	goto L67
L79:
	;
	goto L80
L80:
	;
	v234 = F_pg_sprintf(m, v202, int32(6877), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L14
	} else {
		goto L81
	}
L81:
	;
	goto L67
L82:
	;
	goto L67
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v216
	v248 = F_pg_sprintf(m, v202, int32(6881), v10+int32(-32))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L14
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v217
	*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v216
	v255 = F_pg_sprintf(m, v202, int32(6678), v10+int32(-16))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L14
	} else {
		goto L87
	}
L86:
	;
	goto L67
L87:
	;
	goto L67
L88:
	;
	v315 = v313 + v202
	goto L62
L89:
	;
	v313 = v305 - v202
	goto L88
L90:
	;
	v284 = v280
	goto L99
L91:
	;
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	if v264 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v313 = int32(0)
	goto L88
L93:
	;
	goto L94
L94:
	;
	v269 = v202
	goto L95
L95:
	;
	v273 = v269 + int32(1)
	if v273&int32(3) == int32(0) {
		v280 = v273
		goto L90
	} else {
		goto L97
	}
L96:
	;
	v305 = v273
	goto L89
L97:
	;
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273))))
	if v278 != 0 {
		v269 = v273
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v293 = int32(-2139062144)
	if (int32(16843008)-v290|v290)&v293 == v293 {
		v284 = v284 + int32(4)
		goto L99
	} else {
		goto L101
	}
L100:
	;
	v299 = v284
	goto L102
L101:
	;
	goto L100
L102:
	;
	v303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	if v303 != 0 {
		v299 = v299 + int32(1)
		goto L102
	} else {
		goto L104
	}
L103:
	;
	v305 = v299
	goto L89
L104:
	;
	goto L103
L105:
	;
	goto L20
}
func F_dexp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v13 float64
	_ = v13
	var v16 float64
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 float64
	_ = v25
	var v31 float64
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		if base.F64_eq(base.F64_abs(v4), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			v13 = float64(0)
			if base.F64_gt(v4, v13) != 0 {
				v16 = v4
			} else {
				v16 = v13
			}
			v17 = F_Float8GetDatum(m, v16)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				return v17
			}
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
			v25 = F_exp(m, v4)
			mBase = m.M
			if base.F64_eq(base.F64_abs(v25), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				if base.F64_eq(v25, float64(0)) != 0 {
					F_float_underflow_error(m)
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v31 = v25
					v32 = F_Float8GetDatum(m, v31)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						return v32
					}
				}
			}
		}
	} else {
		v31 = v4
		v32 = F_Float8GetDatum(m, v31)
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int32(0)
		} else {
			return v32
		}
	}
}
func F_dispell_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v39 int32
	_ = v39
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
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
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
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int64
	_ = v414
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v428 int32
	_ = v428
	var v445 int32
	_ = v445
	v2 = int32(0)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v12 <= v2 {
		v445 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v445
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = v15 + int32(8)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v20 = F_str_tolower(m, v18, v12, int32(100))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	v24 = int32(0)
	v25 = int32(1)
	v27 = F_NormalizeSubWord(m, v17, v20, v24)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	if v27 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v29 == int32(0) {
		v71 = v24
		v72 = v2
		v78 = v25
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v84 = v24
	v85 = v2
	v91 = v25
	goto L8
L8:
	;
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+44)))
	if v95 != int32(1) {
		v379 = v84
		goto L23
	} else {
		goto L24
	}
L9:
	;
	F_pfree(m, v27)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L3
	} else {
		goto L22
	}
L10:
	;
	v32 = v24
	v33 = v2
	v35 = v29
	v37 = v27
	v39 = v25
	goto L11
L11:
	;
	if v32 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v71 = v48
	v72 = v62
	v78 = v65
	goto L9
L13:
	;
	v46 = F_palloc(m, int32(8192))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L3
	} else {
		goto L16
	}
L14:
	;
	v48 = v32
	v49 = v33
	goto L15
L15:
	;
	v50 = v49 - v48
	if v50 <= int32(8183) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v48 = v46
	v49 = v46
	goto L15
L17:
	;
	v53 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+2)) = uint16(v53)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v35
	*(*uint16)(unsafe.Add(mBase, uint32(v49))) = uint16(v39)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+12)) = v53
	v60 = v49 + int32(8)
	v62 = v60
	v63 = v60 - v48
	goto L19
L18:
	;
	v62 = v49
	v63 = v50
	goto L19
L19:
	;
	v65 = v39 + int32(1)
	if int32(8191) < v63 {
		v71 = v48
		v72 = v62
		v78 = v65
		goto L9
	} else {
		goto L20
	}
L20:
	;
	v69 = v37 + int32(4)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v70 != 0 {
		v32 = v48
		v33 = v62
		v35 = v70
		v37 = v69
		v39 = v65
		goto L11
	} else {
		goto L21
	}
L21:
	;
	goto L12
L22:
	;
	v84 = v71
	v85 = v72
	v91 = v78
	goto L8
L23:
	;
	if v379 == int32(0) {
		v445 = v2
		goto L1
	} else {
		goto L95
	}
L24:
	;
	v98 = int32(0)
	if v20&int32(3) == v98 {
		v123 = v20
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v159 = F_SplitToVariants(m, v17, v98, v98, v20, v156, int32(0), int32(-1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L3
	} else {
		goto L42
	}
L26:
	;
	v156 = v148 - v20
	goto L25
L27:
	;
	v127 = v123
	goto L36
L28:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v107 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v156 = int32(0)
	goto L25
L30:
	;
	goto L31
L31:
	;
	v112 = v20
	goto L32
L32:
	;
	v116 = v112 + int32(1)
	if v116&int32(3) == int32(0) {
		v123 = v116
		goto L27
	} else {
		goto L34
	}
L33:
	;
	v148 = v116
	goto L26
L34:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116))))
	if v121 != 0 {
		v112 = v116
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127)))
	v136 = int32(-2139062144)
	if (int32(16843008)-v133|v133)&v136 == v136 {
		v127 = v127 + int32(4)
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v142 = v127
	goto L39
L38:
	;
	goto L37
L39:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v142))))
	if v146 != 0 {
		v142 = v142 + int32(1)
		goto L39
	} else {
		goto L41
	}
L40:
	;
	v148 = v142
	goto L26
L41:
	;
	goto L40
L42:
	;
	if v159 == int32(0) {
		v379 = v84
		goto L23
	} else {
		goto L43
	}
L43:
	;
	v163 = v84
	v164 = v85
	v165 = v159
	v170 = v91
	goto L44
L44:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if int32(2) <= v174 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v379 = v324
	goto L23
L46:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v177+v174<<(uint(int32(2))%32)-int32(4))))
	v185 = F_NormalizeSubWord(m, v17, v183, int32(8))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L3
	} else {
		goto L49
	}
L47:
	;
	v324 = v163
	v325 = v164
	v329 = v174
	v331 = v170
	goto L48
L48:
	;
	v335 = int32(0)
	if v329 <= v335 {
		goto L85
	} else {
		goto L86
	}
L49:
	;
	if v185 != 0 {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	if v187 != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v312 = v163
	v313 = v164
	v319 = v170
	goto L52
L52:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v324 = v312
	v325 = v313
	v329 = v323
	v331 = v319
	goto L48
L53:
	;
	v188 = v163
	v189 = v164
	v190 = v187
	v194 = v185
	v195 = v170
	goto L56
L54:
	;
	v286 = v163
	v287 = v164
	v293 = v170
	goto L55
L55:
	;
	F_pfree(m, v185)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L3
	} else {
		goto L83
	}
L56:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if int32(0) < v199-int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v286 = v267
	v287 = v280
	v293 = v282
	goto L55
L58:
	;
	v205 = v188
	v206 = v189
	v207 = int32(0)
	goto L61
L59:
	;
	v251 = v188
	v252 = v189
	v253 = v190
	goto L60
L60:
	;
	if v251 == int32(0) {
		goto L75
	} else {
		goto L76
	}
L61:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v216+v207<<(uint(int32(2))%32))))
	if v185 != v194 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	v251 = v230
	v252 = v243
	v253 = v250
	goto L60
L63:
	;
	v222 = F_pstrdup(m, v220)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L3
	} else {
		goto L66
	}
L64:
	;
	v224 = v220
	goto L65
L65:
	;
	if v205 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v224 = v222
	goto L65
L67:
	;
	v228 = F_palloc(m, int32(8192))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L3
	} else {
		goto L70
	}
L68:
	;
	v230 = v205
	v231 = v206
	goto L69
L69:
	;
	if v231-v230 <= int32(8183) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v230 = v228
	v231 = v228
	goto L69
L71:
	;
	v235 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v231)+2)) = uint16(v235)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+4)) = v224
	*(*uint16)(unsafe.Add(mBase, uint32(v231))) = uint16(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v231)+12)) = v235
	v243 = v231 + int32(8)
	goto L73
L72:
	;
	v243 = v231
	goto L73
L73:
	;
	v244 = int32(1)
	v245 = v207 + v244
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v245 < v246-v244 {
		v205 = v230
		v206 = v243
		v207 = v245
		goto L61
	} else {
		goto L74
	}
L74:
	;
	goto L62
L75:
	;
	v265 = F_palloc(m, int32(8192))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L3
	} else {
		goto L78
	}
L76:
	;
	v267 = v251
	v268 = v252
	goto L77
L77:
	;
	if v268-v267 <= int32(8183) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v267 = v265
	v268 = v265
	goto L77
L79:
	;
	v272 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+2)) = uint16(v272)
	*(*int32)(unsafe.Add(mBase, uint32(v268)+4)) = v253
	*(*uint16)(unsafe.Add(mBase, uint32(v268))) = uint16(v195)
	*(*int32)(unsafe.Add(mBase, uint32(v268)+12)) = v272
	v280 = v268 + int32(8)
	goto L81
L80:
	;
	v280 = v268
	goto L81
L81:
	;
	v282 = v195 + int32(1)
	v284 = v194 + int32(4)
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	if v285 != 0 {
		v188 = v267
		v189 = v280
		v190 = v285
		v194 = v284
		v195 = v282
		goto L56
	} else {
		goto L82
	}
L82:
	;
	goto L57
L83:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v299))) = int32(0)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v302+v303<<(uint(int32(2))%32)-int32(4))))
	F_pfree(m, v309)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L3
	} else {
		goto L84
	}
L84:
	;
	v312 = v286
	v313 = v287
	v319 = v293
	goto L52
L85:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	F_pfree(m, v374)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L3
	} else {
		goto L92
	}
L86:
	;
	v340 = v335
	goto L87
L87:
	;
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v165)+8))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v349+v340<<(uint(int32(2))%32))))
	if v353 == int32(0) {
		goto L85
	} else {
		goto L89
	}
L88:
	;
	goto L85
L89:
	;
	F_pfree(m, v353)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L3
	} else {
		goto L90
	}
L90:
	;
	v359 = v340 + int32(1)
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	if v359 < v360 {
		v340 = v359
		goto L87
	} else {
		goto L91
	}
L91:
	;
	goto L88
L92:
	;
	F_pfree(m, v165)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L3
	} else {
		goto L93
	}
L93:
	;
	if v373 != 0 {
		v163 = v324
		v164 = v325
		v165 = v373
		v170 = v331
		goto L44
	} else {
		goto L94
	}
L94:
	;
	goto L45
L95:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v379)+4))
	if v392 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v395 = v379
	v398 = v392
	v399 = v379
	v403 = v379 + int32(4)
	goto L99
L97:
	;
	v428 = v379
	goto L98
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v428)+4)) = int32(0)
	v445 = v379
	goto L1
L99:
	;
	v406 = F_searchstoplist(m, v15, v398)
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L3
	} else {
		goto L102
	}
L100:
	;
	v428 = v418
	goto L98
L101:
	;
	v420 = v395 + int32(12)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	if v423 != 0 {
		v395 = v395 + int32(8)
		v398 = v423
		v399 = v418
		v403 = v420
		goto L99
	} else {
		goto L110
	}
L102:
	;
	if v406 != 0 {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v403)))
	F_pfree(m, v408)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L3
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v395 != v399 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v403))) = int32(0)
	v418 = v399
	goto L101
L107:
	;
	v414 = *(*int64)(unsafe.Add(mBase, uint32(v395)))
	*(*int64)(unsafe.Add(mBase, uint32(v399))) = v414
	goto L109
L108:
	;
	goto L109
L109:
	;
	v418 = v399 + int32(8)
	goto L101
L110:
	;
	goto L100
}
func F_div_var_int(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
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
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int64
	_ = v71
	var v74 int32
	_ = v74
	var v88 int64
	_ = v88
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v103 int64
	_ = v103
	var v104 int64
	_ = v104
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
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
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v383 int32
	_ = v383
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v19 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L10
	} else {
		goto L91
	}
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = int32(1)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = v35 - l2
	v40 = base.I32_div_s(l4+int32(3), int32(4))
	v43 = v36 + v40 + v34
	if v43 <= v34 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	F_pfree(m, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(0)
	v28 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = v28
	*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v28
	return
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	v46 = v34
	goto L14
L13:
	;
	v46 = v43
	goto L14
L14:
	;
	v47 = v46 + l5
	v52 = F_palloc(m, v47<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v54 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v52))) = uint16(v54)
	v57 = v52 + int32(2)
	v64 = l1 >> (uint(int32(31)) % 32)
	v66 = l1 ^ v64 - v64
	if base.Ui32(int32(429497)) <= base.Ui32(v66) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if base.B2i32(v33 == v54)^base.B2i32(v54 < l1) != 0 {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	if v47 <= int32(0) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	if v47 <= int32(0) {
		goto L16
	} else {
		goto L27
	}
L20:
	;
	v71 = base.I64_extend_i32_u(v66)
	v74 = int32(0)
	v88 = int64(0)
	goto L21
L21:
	;
	if v74 < v19 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L16
L23:
	;
	v99 = int64(*(*int16)(unsafe.Add(mBase, uint32(v32+v74<<(uint(int32(1))%32)))))
	v100 = v99
	goto L25
L24:
	;
	v100 = int64(0)
	goto L25
L25:
	;
	v103 = v100 + v88*int64(10000)
	v104 = base.I64_div_u_s(v103, v71)
	*(*uint16)(unsafe.Add(mBase, uint32(v57+v74<<(uint(int32(1))%32)))) = uint16(v104)
	v109 = v74 + int32(1)
	if v109 != v47 {
		v74 = v109
		v88 = v103 - v104*v71
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v115 = int32(0)
	v122 = int32(0)
	goto L28
L28:
	;
	if v115 < v19 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L16
L30:
	;
	v140 = int32(*(*int16)(unsafe.Add(mBase, uint32(v32+v115<<(uint(int32(1))%32)))))
	v141 = v140
	goto L32
L31:
	;
	v141 = int32(0)
	goto L32
L32:
	;
	v144 = v141 + v122*int32(10000)
	v145 = base.I32_div_u_s(v144, v66)
	*(*uint16)(unsafe.Add(mBase, uint32(v57+v115<<(uint(int32(1))%32)))) = uint16(v145)
	v150 = v115 + int32(1)
	if v150 != v47 {
		v115 = v150
		v122 = v144 - v66*v145
		goto L28
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	v172 = int32(16384)
	goto L36
L35:
	;
	v172 = int32(0)
	goto L36
L36:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v173 != 0 {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	F_pfree(m, v173)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L10
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(l3)+16)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v47
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v172
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v36
	if l5 != 0 {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v430
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v429
	return
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+4)) = int64(0)
	v429 = v408
	v430 = int32(0)
	goto L41
L43:
	;
	if int32(0) < v344 {
		goto L78
	} else {
		goto L79
	}
L44:
	;
	v341 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v343 = v341
	v344 = v342
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = l4
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v190 = l4 + v187<<(uint(int32(2))%32)
	if v190+int32(4) < int32(0) {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L47
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = l4
	v309 = l4 + v36<<(uint(int32(2))%32)
	if v309+int32(4) <= int32(0) {
		v408 = v57
		goto L42
	} else {
		goto L75
	}
L48:
	;
	goto L44
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l3))) = int64(0)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v201 = l4 & int32(3)
	v205 = base.I32_div_s(v190+int32(7), int32(4))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v206 <= v205 {
		goto L56
	} else {
		goto L57
	}
L52:
	;
	goto L48
L53:
	;
	if int32(0) <= v272 {
		goto L52
	} else {
		goto L74
	}
L54:
	;
	v252 = v246
	goto L68
L55:
	;
	v219 = int32(1)
	v220 = v205 - v219
	v223 = v199 + v220<<(uint(v219)%32)
	v224 = int32(*(*int16)(unsafe.Add(mBase, uint32(v223))))
	v225 = int32(2)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v201<<(uint(v225)%32))+uint32(_consts[975])))
	v230 = base.I32_rem_s(v224, v229)
	v231 = v224 - v230
	*(*uint16)(unsafe.Add(mBase, uint32(v223))) = uint16(v231)
	v234 = base.I32_div_s(v229, v225)
	if v230 < v234 {
		v272 = v220
		goto L53
	} else {
		goto L63
	}
L56:
	;
	if v201 == int32(0) {
		goto L52
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v205
	if v201 != 0 {
		goto L55
	} else {
		goto L61
	}
L59:
	;
	if v205 != v206 {
		goto L52
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v205
	goto L55
L61:
	;
	v216 = int32(*(*int16)(unsafe.Add(mBase, uint32(v199+v205<<(uint(int32(1))%32)))))
	if v216 <= int32(4999) {
		v272 = v205
		goto L53
	} else {
		goto L62
	}
L62:
	;
	v246 = v205
	goto L54
L63:
	;
	v237 = v229 + base.I32_extend16_s(v231)
	if int32(9999) < v237 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v242 = v237 + int32(55536)
	goto L66
L65:
	;
	v242 = v237
	goto L66
L66:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v223))) = uint16(v242)
	if v237 < int32(10000) {
		v272 = v220
		goto L53
	} else {
		goto L67
	}
L67:
	;
	v246 = v220
	goto L54
L68:
	;
	v258 = int32(1)
	v259 = v252 - v258
	v262 = v199 + v259<<(uint(v258)%32)
	v265 = int32(*(*int16)(unsafe.Add(mBase, uint32(v262))))
	v267 = base.B2i32(int32(9998) < v265)
	if int32(9998) < v265 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v272 = v259
	goto L53
L70:
	;
	v268 = int32(-9999)
	goto L72
L71:
	;
	v268 = v258
	goto L72
L72:
	;
	v269 = v268 + v265
	*(*uint16)(unsafe.Add(mBase, uint32(v262))) = uint16(v269)
	if int32(9998) < v265 {
		v252 = v259
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+20)) = v280 - int32(2)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v285 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v284 + v285
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v288 + v285
	goto L52
L75:
	;
	v317 = base.I32_div_s(v309+int32(7), int32(4))
	if v46 < v317 {
		goto L44
	} else {
		goto L76
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v317
	v321 = l4 & int32(3)
	if v321 == int32(0) {
		v343 = v57
		v344 = v317
		goto L43
	} else {
		goto L77
	}
L77:
	;
	v327 = int32(2)
	v328 = v57 + v317<<(uint(int32(1))%32) - v327
	v329 = int32(*(*int16)(unsafe.Add(mBase, uint32(v328))))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v321<<(uint(v327)%32))+uint32(_consts[975])))
	v335 = base.I32_rem_s(v329, v334)
	v336 = v329 - v335
	*(*uint16)(unsafe.Add(mBase, uint32(v328))) = uint16(v336)
	goto L44
L78:
	;
	v351 = v343
	v352 = v344
	goto L82
L79:
	;
	goto L80
L80:
	;
	if v344 != 0 {
		v429 = v343
		v430 = v344
		goto L41
	} else {
		goto L90
	}
L81:
	;
	v383 = v352
	goto L86
L82:
	;
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351))))
	if v369 != 0 {
		goto L81
	} else {
		goto L84
	}
L83:
	;
	v408 = v343 + v344<<(uint(int32(1))%32)
	goto L42
L84:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v371 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v370 - v371
	if v371 < v352 {
		v351 = v351 + int32(2)
		v352 = v352 - v371
		goto L82
	} else {
		goto L85
	}
L85:
	;
	goto L83
L86:
	;
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351-int32(2)+v383<<(uint(int32(1))%32)))))
	if v403 != 0 {
		v429 = v351
		v430 = v383
		goto L41
	} else {
		goto L88
	}
L87:
	;
	v408 = v351
	goto L42
L88:
	;
	v404 = int32(1)
	if v404 < v383 {
		v383 = v383 - v404
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v408 = v343
	goto L42
L91:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L10
	} else {
		goto L92
	}
L92:
	;
	F_errmsg(m, int32(230644), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(483087), int32(9924), int32(87840))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_do_setval(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v213 int32
	_ = v213
	var v218 int32
	_ = v218
	v3 = l2
	v9 = m.G0
	v11 = v9 - int32(112)
	m.G0 = v11
	F_init_sequence(m, l0, v11+int32(108), v11+int32(104))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+108))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v22 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v24 = F_pg_class_aclcheck(m, v20, v22, int64(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L1
	} else {
		goto L58
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L55
	}
L5:
	;
	if v24 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = F_SearchSysCache1(m, int32(61), l0)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L51
	}
L9:
	;
	if v29 == int32(0) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v35 = v33 + v34
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v35)+32))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v35)+24))
	F_ReleaseCatCache(m, v29)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+24)))
	if v41 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_PreventCommandIfReadOnly(m, int32(649492))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_PreventCommandIfParallelMode(m, int32(649492))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v54 = F_read_seq_tuple(m, v40, v11+int32(100), v11+int32(80))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if l1 < v36 {
		goto L3
	} else {
		goto L18
	}
L18:
	;
	if v37 < l1 {
		goto L3
	} else {
		goto L19
	}
L19:
	;
	if v3 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+118)))
	if v67 != int32(112) {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
	v64 = v60
	goto L20
L22:
	;
	goto L23
L23:
	;
	v61 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+12)) = uint8(v61)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = l1
	v64 = l1
	goto L20
L24:
	;
	v78 = int32(4449876)
	v80 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v80 + int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v54)+16)) = uint8(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v54))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v54)+8)) = int64(0)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+100))
	F_MarkBufferDirty(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L32
	}
L25:
	;
	v71 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v71 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	if v74 != 0 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v76 = F_GetTopTransactionId(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v75 != 0 {
		goto L24
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	goto L24
L32:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+118)))
	if v92 != int32(112) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v146 = int32(4449876)
	v148 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	*(*int32)(unsafe.Add(mBase, _consts[14])) = v148 - int32(1)
	F_UnlockReleaseBuffer(m, v88)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L49
	}
L34:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	if v96 <= int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v40)+32))
	if v99 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v88 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v40)+40))
	if v100 != 0 {
		goto L33
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L44
	}
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v104+(v88^int32(-1))<<(uint(int32(2))%32))))
	v118 = v110
	goto L40
L42:
	;
	goto L43
L43:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v118 = v112 + v88<<(uint(int32(13))%32) + int32(-8192)
	goto L40
L44:
	;
	F_XLogRegisterBuffer(m, int32(0), v88, int32(6))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v40)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+72)) = v125
	v127 = *(*int64)(unsafe.Add(mBase, uint32(v40)))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+64)) = v127
	F_XLogRegisterData(m, v11-int32(-64), int32(12))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+96))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v11)+80))
	F_XLogRegisterData(m, v134, v135)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v140 = F_XLogInsert(m, int32(15), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v118))) = base.I64_rotr(v140, int64(32))
	goto L33
L49:
	;
	F_sequence_close(m, v40, int32(0))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	m.G0 = v11 + int32(112)
	return
L51:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v11)+104))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = v168 + int32(4)
	F_errmsg(m, int32(188298), v11+int32(48))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	F_errfinish(m, int32(482601), int32(964), int32(297627))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(51191), v11)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(482601), int32(968), int32(297627))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
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
	F_errcode(m, int32(50331778))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v40)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+40)) = v37
	*(*int64)(unsafe.Add(mBase, uint32(v11)+32)) = v36
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v202 + int32(4)
	F_errmsg(m, int32(642621), v11+int32(16))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(482601), int32(993), int32(297627))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dpi(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_Float8GetDatum(m, float64(3.141592653589793))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_dprintf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l1
	v10 = m.G0
	v11 = int32(144)
	v12 = v10 - v11
	m.G0 = v12
	v17 = F__emscripten_memset_bulkmem(m, v12, base.I32_extend8_s(int32(0)), v11)
	mBase = m.M
	v18 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = l0
	v21 = int32(167965)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+44)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = int32(5874)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v18
	v28 = F_vfprintf(m, v17, v21, l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return
	} else {
		m.G0 = v17 + int32(144)
		m.G0 = v7 + int32(16)
		return
	}
}
func F_drandom(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int64
	_ = v59
	var v62 int64
	_ = v62
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v90 int64
	_ = v90
	var v93 int64
	_ = v93
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v113 int64
	_ = v113
	var v118 int64
	_ = v118
	var v123 int64
	_ = v123
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v164 float64
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	v3 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1023])))
	if v3 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = int32(16)
	v8 = int32(0)
	v12 = m.G0
	v14 = v12 - v7
	m.G0 = v14
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v8
	v20 = F_open(m, int32(277917), v8, v14)
	mBase = m.M
	if v20 != int32(-1) {
		goto L6
	} else {
		goto L7
	}
L2:
	;
	goto L3
L3:
	;
	v138 = int32(4440064)
	v141 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
	v142 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
	v143 = v141 ^ v142
	*(*int64)(unsafe.Add(mBase, _consts[1021])) = base.I64_rotl(v143, int64(37))
	*(*int64)(unsafe.Add(mBase, _consts[1022])) = v143<<(uint(int64(16))%64) ^ base.I64_rotl(v141, int64(24)) ^ v143
	v164 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v141*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L31
L4:
	;
	v136 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1023])) = uint8(v136)
	goto L3
L5:
	;
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	goto L10
L7:
	;
	v53 = v8
	goto L8
L8:
	;
	m.G0 = v14 + int32(16)
	goto L5
L9:
	;
	v48 = F_close(m, v20)
	mBase = m.M
	v53 = v46
	goto L8
L10:
	;
	v26 = int32(4440064)
	v27 = v7
	goto L11
L11:
	;
	v32 = F_read(m, v20, v26, v27)
	mBase = m.M
	if v32 <= int32(0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v46 = int32(1)
	goto L9
L13:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[40]))
	if v36 == int32(27) {
		goto L11
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v41 = v27 - v32
	if v41 != 0 {
		v26 = v26 + v32
		v27 = v41
		goto L11
	} else {
		goto L17
	}
L16:
	;
	v46 = int32(0)
	goto L9
L17:
	;
	goto L12
L18:
	;
	v58 = int32(4440064)
	v59 = *(*int64)(unsafe.Add(mBase, _consts[1022]))
	if v59 != int64(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	goto L20
L20:
	;
	v70 = int32(4440064)
	v74 = m.G0
	v75 = int32(16)
	v76 = v74 - v75
	m.G0 = v76
	F___gettimeofday(m, v76)
	mBase = m.M
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
	v80 = int64(*(*int32)(unsafe.Add(mBase, uint32(v76)+8)))
	m.G0 = v76 + v75
	goto L26
L21:
	;
	goto L4
L22:
	;
	goto L21
L23:
	;
	v62 = *(*int64)(unsafe.Add(mBase, _consts[1021]))
	if v62 != int64(0) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1021])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _consts[1022])) = int64(6364136223846793005)
	goto L22
L26:
	;
	v90 = int64(*(*uint32)(unsafe.Add(mBase, _consts[129])))
	v93 = v80 + v79*int64(1000000) - int64(946684800000000) ^ v90<<(uint(int64(32))%64)
	v97 = v93 + int64(4354685564936845354)
	v98 = int64(30)
	v101 = int64(-4658895280553007687)
	v102 = (int64(base.Ui64(v97)>>(uint(v98)%64)) ^ v97) * v101
	v103 = int64(27)
	v106 = int64(-7723592293110705685)
	v107 = (int64(base.Ui64(v102)>>(uint(v103)%64)) ^ v102) * v106
	v108 = int64(31)
	*(*int64)(unsafe.Add(mBase, _consts[1021])) = int64(base.Ui64(v107)>>(uint(v108)%64)) ^ v107
	v113 = v93 - int64(7046029254386353131)
	v118 = (int64(base.Ui64(v113)>>(uint(v98)%64)) ^ v113) * v101
	v123 = (int64(base.Ui64(v118)>>(uint(v103)%64)) ^ v118) * v106
	*(*int64)(unsafe.Add(mBase, _consts[1022])) = int64(base.Ui64(v123)>>(uint(v108)%64)) ^ v123
	if v113|v97 == int64(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	goto L4
L28:
	;
	*(*int64)(unsafe.Add(mBase, _consts[1021])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _consts[1022])) = int64(6364136223846793005)
	goto L30
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	v165 = F_Float8GetDatum(m, v164)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	return int32(0)
L33:
	;
	return v165
}
func F_dsimple_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = F_palloc0(m, int32(12))
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
	v19 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v19)
	if v13 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L46
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L42
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L38
	}
L6:
	;
	m.G0 = v11 + int32(16)
	return v15
L7:
	;
	v23 = int32(0)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v24 <= v23 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v27 = v23
	v32 = v2
	v33 = v2
	goto L9
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v27<<(uint(int32(2))%32))))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v41 = int32(164640)
	v44 = int32(*(*uint8)(unsafe.Add(mBase, _consts[893])))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v45 == int32(0) {
		v64 = v44
		v65 = v45
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L6
L11:
	;
	v108 = v27 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	if v108 < v109 {
		v27 = v108
		v32 = v105
		v33 = v106
		goto L9
	} else {
		goto L37
	}
L12:
	;
	if v65-v64 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	goto L12
L14:
	;
	if v44 != v45 {
		v64 = v44
		v65 = v45
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v49 = v40
	v50 = v41
	goto L16
L16:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50)+1)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+1)))
	if v54 == int32(0) {
		v64 = v53
		v65 = v54
		goto L13
	} else {
		goto L18
	}
L17:
	;
	v64 = v53
	v65 = v54
	goto L13
L18:
	;
	v57 = int32(1)
	if v53 == v54 {
		v49 = v49 + v57
		v50 = v50 + v57
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	if v33 != 0 {
		goto L5
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v75 = int32(81087)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _consts[894])))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v79 == int32(0) {
		v98 = v78
		v99 = v79
		goto L27
	} else {
		goto L28
	}
L23:
	;
	v69 = F_defGetString(m, v39)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_readstoplist(m, v69, v15, int32(1162))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v105 = v32
	v106 = int32(1)
	goto L11
L26:
	;
	if v99-v98 != 0 {
		goto L3
	} else {
		goto L34
	}
L27:
	;
	goto L26
L28:
	;
	if v78 != v79 {
		v98 = v78
		v99 = v79
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v83 = v40
	v84 = v75
	goto L30
L30:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v87
		v99 = v88
		goto L27
	} else {
		goto L32
	}
L31:
	;
	v98 = v87
	v99 = v88
	goto L27
L32:
	;
	v91 = int32(1)
	if v87 == v88 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	if v32 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v101 = F_defGetBoolean(m, v39)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v15)+8)) = uint8(v101)
	v105 = int32(1)
	v106 = v33
	goto L11
L37:
	;
	goto L10
L38:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errmsg(m, int32(125578), int32(0))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(482043), int32(50), int32(95855))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(125551), int32(0))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_errfinish(m, int32(482043), int32(59), int32(95855))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L46:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v162
	F_errmsg(m, int32(690700), v11)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(482043), int32(68), int32(95855))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dsynonym_lexize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	if v10 <= v2 {
		v54 = v2
		m.G0 = v8 + int32(16)
		return v54
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		if v14 <= int32(0) {
			v54 = v2
			m.G0 = v8 + int32(16)
			return v54
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+8)))
			if v18 == int32(1) {
				v21 = F_pnstrdup(m, v17, v10)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v28 = v21
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v36 = F_bsearch(m, v8, v32, v33, int32(16), int32(1163))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						F_pfree(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v36 == int32(0) {
								v54 = v2
								m.G0 = v8 + int32(16)
								return v54
							} else {
								v44 = F_palloc0(m, int32(16))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
									v48 = F_pnstrdup(m, v46, v47)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v48
										v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
										*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)) = uint16(v51)
										v54 = v44
										m.G0 = v8 + int32(16)
										return v54
									}
								}
							}
						}
					}
				}
			} else {
				v26 = F_str_tolower(m, v17, v10, int32(100))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v28 = v26
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v28
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v36 = F_bsearch(m, v8, v32, v33, int32(16), int32(1163))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						F_pfree(m, v38)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							if v36 == int32(0) {
								v54 = v2
								m.G0 = v8 + int32(16)
								return v54
							} else {
								v44 = F_palloc0(m, int32(16))
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
									v47 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
									v48 = F_pnstrdup(m, v46, v47)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v44)+4)) = v48
										v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+12)))
										*(*uint16)(unsafe.Add(mBase, uint32(v44)+2)) = uint16(v51)
										v54 = v44
										m.G0 = v8 + int32(16)
										return v54
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
func F_dt2time(m *base.Module, l0 int64, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int64
	_ = v8
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	v8 = base.I64_div_s(l0, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l1))) = uint32(v8)
	v13 = base.I64_extend32_s(v8)*int64(-3600000000) + l0
	v15 = base.I64_div_s(v13, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l2))) = uint32(v15)
	v20 = base.I64_extend32_s(v15)*int64(-60000000) + v13
	v22 = base.I64_div_s(v20, int64(1000000))
	*(*uint32)(unsafe.Add(mBase, uint32(l3))) = uint32(v22)
	v26 = v22*int64(4293967296) + v20
	*(*uint32)(unsafe.Add(mBase, uint32(l4))) = uint32(v26)
	return
}
func F_dtan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 float64
	_ = v34
	var v38 int32
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	if base.Ui64(base.I64_reinterpret_f64(v4)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0)
		if base.F64_eq(base.F64_abs(v4), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50331778))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(388004), int32(0))
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(476769), int32(1980), int32(272320))
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
		} else {
			v18 = m.G0
			v20 = v18 - int32(16)
			m.G0 = v20
			v27 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v4))>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(v27) <= base.Ui32(int32(1072243195)) {
				if base.Ui32(v27) < base.Ui32(int32(1044381696)) {
					v44 = v4
				} else {
					v34 = F___tan(m, v4, float64(0), int32(0))
					mBase = m.M
					v44 = v34
				}
			} else {
				if base.Ui32(int32(2146435072)) <= base.Ui32(v27) {
					v44 = base.F64_sub(v4, v4)
				} else {
					v38 = F___rem_pio2(m, v4, v20)
					mBase = m.M
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v20)))
					v40 = *(*float64)(unsafe.Add(mBase, uint32(v20)+8))
					v43 = F___tan(m, v39, v40, v38&int32(1))
					mBase = m.M
					v44 = v43
				}
			}
			m.G0 = v20 + int32(16)
			v50 = v44
			v51 = F_Float8GetDatum(m, v50)
			mBase = m.M
			v54 = m.ExcPending
			if v54 != 0 {
				return int32(0)
			} else {
				return v51
			}
		}
	} else {
		v50 = math.Float64frombits(uint64(0x7ff8000000000000))
		v51 = F_Float8GetDatum(m, v50)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			return v51
		}
	}
}
func F_dtand(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 float64
	_ = v13
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v36 int64
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 float64
	_ = v45
	var v48 int64
	_ = v48
	var v55 float64
	_ = v55
	var v58 int32
	_ = v58
	var v60 int64
	_ = v60
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v70 int32
	_ = v70
	var v72 int64
	_ = v72
	var v76 int32
	_ = v76
	var v90 int32
	_ = v90
	var v95 int64
	_ = v95
	var v99 int32
	_ = v99
	var v101 int64
	_ = v101
	var v105 int64
	_ = v105
	var v112 int64
	_ = v112
	var v114 int64
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v123 int64
	_ = v123
	var v127 int64
	_ = v127
	var v134 int64
	_ = v134
	var v138 int32
	_ = v138
	var v140 int64
	_ = v140
	var v144 int32
	_ = v144
	var v148 int64
	_ = v148
	var v150 int32
	_ = v150
	var v153 int64
	_ = v153
	var v167 int64
	_ = v167
	var v178 float64
	_ = v178
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 float64
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 float64
	_ = v192
	var v195 int32
	_ = v195
	var v196 float64
	_ = v196
	var v200 float64
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v213 int32
	_ = v213
	var v220 float64
	_ = v220
	var v224 int32
	_ = v224
	var v225 float64
	_ = v225
	var v226 float64
	_ = v226
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v235 float64
	_ = v235
	var v237 float64
	_ = v237
	var v239 float64
	_ = v239
	var v248 float64
	_ = v248
	var v256 float64
	_ = v256
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v276 float64
	_ = v276
	var v280 int32
	_ = v280
	var v281 float64
	_ = v281
	var v282 float64
	_ = v282
	var v287 float64
	_ = v287
	var v289 float64
	_ = v289
	var v291 float64
	_ = v291
	var v294 float64
	_ = v294
	var v298 float64
	_ = v298
	var v302 float64
	_ = v302
	var v306 float64
	_ = v306
	var v312 float64
	_ = v312
	var v313 int32
	_ = v313
	var v318 float64
	_ = v318
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v338 float64
	_ = v338
	var v342 int32
	_ = v342
	var v343 float64
	_ = v343
	var v344 float64
	_ = v344
	var v349 float64
	_ = v349
	var v351 float64
	_ = v351
	var v353 float64
	_ = v353
	var v356 float64
	_ = v356
	var v360 float64
	_ = v360
	var v364 float64
	_ = v364
	var v368 float64
	_ = v368
	var v377 float64
	_ = v377
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v397 float64
	_ = v397
	var v401 int32
	_ = v401
	var v402 float64
	_ = v402
	var v403 float64
	_ = v403
	var v409 float64
	_ = v409
	var v410 float64
	_ = v410
	var v412 float64
	_ = v412
	var v414 float64
	_ = v414
	var v416 float64
	_ = v416
	var v425 float64
	_ = v425
	var v429 float64
	_ = v429
	var v430 float64
	_ = v430
	var v432 float64
	_ = v432
	var v435 float64
	_ = v435
	var v438 float64
	_ = v438
	var v441 float64
	_ = v441
	var v447 float64
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	if base.Ui64(base.I64_reinterpret_f64(v13)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L127
	} else {
		goto L129
	}
L2:
	;
	if base.F64_eq(base.F64_abs(v13), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	v447 = math.Float64frombits(uint64(0x7ff8000000000000))
	goto L4
L4:
	;
	v448 = F_Float8GetDatum(m, v447)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L127
	} else {
		goto L128
	}
L5:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[958])))
	if v23 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_init_degree_constants(m)
	mBase = m.M
	goto L8
L7:
	;
	goto L8
L8:
	;
	v27 = int32(0)
	v36 = base.I64_reinterpret_f64(v13)
	v40 = int32(2047)
	v41 = base.I32_wrap_i64(int64(base.Ui64(v36)>>(uint(int64(52))%64))) & v40
	if v41 != v40 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v180 = base.F64_lt(v178, float64(0))
	if v180 != 0 {
		goto L50
	} else {
		goto L51
	}
L10:
	;
	v48 = v36 << (uint(int64(1)) % 64)
	if base.Ui64(v48) <= base.Ui64(int64(-9156662467374350336)) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v45 = base.F64_mul(v13, float64(360))
	v178 = base.F64_div(v45, v45)
	goto L9
L12:
	;
	if v48 == int64(-9156662467374350336) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	goto L14
L14:
	;
	if v41 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v55 = base.F64_mul(v13, float64(0))
	goto L17
L16:
	;
	v55 = v13
	goto L17
L17:
	;
	v178 = v55
	goto L9
L18:
	;
	if int32(1031) < v90 {
		goto L28
	} else {
		goto L29
	}
L19:
	;
	v58 = int32(0)
	v60 = v36 << (uint(int64(12)) % 64)
	if int64(0) <= v60 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v90 = v41
	v95 = v36&int64(4503599627370495) | int64(4503599627370496)
	goto L18
L22:
	;
	v64 = v58
	v66 = v60
	goto L25
L23:
	;
	v76 = v58
	goto L24
L24:
	;
	v90 = v76
	v95 = v36 << (uint(base.I64_extend_i32_u(int32(1)-v76)) % 64)
	goto L18
L25:
	;
	v70 = v64 - int32(1)
	v72 = v66 << (uint(int64(1)) % 64)
	if int64(0) <= v72 {
		v64 = v70
		v66 = v72
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v76 = v70
	goto L24
L27:
	;
	goto L26
L28:
	;
	v99 = v90
	v101 = v95
	goto L31
L29:
	;
	v121 = v90
	v123 = v95
	goto L30
L30:
	;
	v127 = v123 - int64(6333186975989760)
	if v127 < int64(0) {
		v134 = v123
		goto L37
	} else {
		goto L38
	}
L31:
	;
	v105 = v101 - int64(6333186975989760)
	if v105 < int64(0) {
		v112 = v101
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v121 = int32(1031)
	v123 = v114
	goto L30
L33:
	;
	v114 = v112 << (uint(int64(1)) % 64)
	v116 = v99 - int32(1)
	if int32(1031) < v116 {
		v99 = v116
		v101 = v114
		goto L31
	} else {
		goto L36
	}
L34:
	;
	if v105 != int64(0) {
		v112 = v105
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v178 = base.F64_mul(v13, float64(0))
	goto L9
L36:
	;
	goto L32
L37:
	;
	if base.Ui64(int64(4503599627370495)) < base.Ui64(v134) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	if v127 != int64(0) {
		v134 = v127
		goto L37
	} else {
		goto L39
	}
L39:
	;
	v178 = base.F64_mul(v13, float64(0))
	goto L9
L40:
	;
	if int32(0) < v150 {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	v150 = v121
	v153 = v134
	goto L40
L42:
	;
	goto L43
L43:
	;
	v138 = v121
	v140 = v134
	goto L44
L44:
	;
	v144 = v138 - int32(1)
	v148 = v140 << (uint(int64(1)) % 64)
	if base.Ui64(v140) < base.Ui64(int64(2251799813685248)) {
		v138 = v144
		v140 = v148
		goto L44
	} else {
		goto L46
	}
L45:
	;
	v150 = v144
	v153 = v148
	goto L40
L46:
	;
	goto L45
L47:
	;
	v167 = v153 - int64(4503599627370496) | base.I64_extend_i32_u(v150)<<(uint(int64(52))%64)
	goto L49
L48:
	;
	v167 = int64(base.Ui64(v153) >> (uint(base.I64_extend_i32_u(int32(1)-v150)) % 64))
	goto L49
L49:
	;
	v178 = base.F64_reinterpret_i64(v167 | v36&int64(-9223372036854775807-1))
	goto L9
L50:
	;
	v181 = int32(-1)
	goto L52
L51:
	;
	v181 = int32(1)
	goto L52
L52:
	;
	if v180 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v184 = base.F64_neg(v178)
	goto L55
L54:
	;
	v184 = v178
	goto L55
L55:
	;
	v186 = base.F64_gt(v184, float64(180))
	if v186 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v187 = v27 - v181
	goto L58
L57:
	;
	v187 = v181
	goto L58
L58:
	;
	if v186 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v195 != 0 {
		goto L93
	} else {
		goto L94
	}
L60:
	;
	v192 = base.F64_sub(float64(360), v184)
	goto L62
L61:
	;
	v192 = v184
	goto L62
L62:
	;
	v195 = base.F64_gt(v192, float64(90))
	if v195 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v196 = base.F64_sub(float64(180), v192)
	goto L65
L64:
	;
	v196 = v192
	goto L65
L65:
	;
	if base.F64_le(v196, float64(30)) != 0 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v200 = base.F64_mul(v196, float64(0.017453292519943295))
	v204 = m.G0
	v206 = v204 - int32(16)
	m.G0 = v206
	v213 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v200))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v213) <= base.Ui32(int32(1072243195)) {
		goto L71
	} else {
		goto L72
	}
L67:
	;
	goto L68
L68:
	;
	v256 = base.F64_mul(base.F64_sub(float64(90), v196), float64(0.017453292519943295))
	v260 = m.G0
	v262 = v260 - int32(16)
	m.G0 = v262
	v269 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v256))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v269) <= base.Ui32(int32(1072243195)) {
		goto L84
	} else {
		goto L85
	}
L69:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v239
	v248 = *(*float64)(unsafe.Add(mBase, _consts[953]))
	v312 = base.F64_mul(base.F64_div(v239, v248), float64(0.5))
	goto L59
L70:
	;
	m.G0 = v206 + int32(16)
	goto L69
L71:
	;
	if base.Ui32(v213) < base.Ui32(int32(1045430272)) {
		v239 = v200
		goto L70
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v213) {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v220 = F___sin(m, v200, float64(0), int32(0))
	mBase = m.M
	v239 = v220
	goto L70
L75:
	;
	v239 = base.F64_sub(v200, v200)
	goto L70
L76:
	;
	goto L77
L77:
	;
	v224 = F___rem_pio2(m, v200, v206)
	mBase = m.M
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v206)+8))
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v206)))
	switch v224&int32(3) - int32(1) {
	case 0:
		goto L80
	case 1:
		goto L79
	case 2:
		goto L78
	default:
		goto L81
	}
L78:
	;
	v237 = F___cos(m, v226, v225)
	mBase = m.M
	v239 = base.F64_neg(v237)
	goto L70
L79:
	;
	v235 = F___sin(m, v226, v225, int32(1))
	mBase = m.M
	v239 = base.F64_neg(v235)
	goto L70
L80:
	;
	v233 = F___cos(m, v226, v225)
	mBase = m.M
	v239 = v233
	goto L70
L81:
	;
	v232 = F___sin(m, v226, v225, int32(1))
	mBase = m.M
	v239 = v232
	goto L70
L82:
	;
	v302 = base.F64_sub(float64(1), v298)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v302
	v306 = *(*float64)(unsafe.Add(mBase, _consts[955]))
	v312 = base.F64_add(base.F64_mul(base.F64_div(v302, v306), float64(-0.5)), float64(1))
	goto L59
L83:
	;
	m.G0 = v262 + int32(16)
	goto L82
L84:
	;
	if base.Ui32(v269) < base.Ui32(int32(1044816030)) {
		v298 = float64(1)
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v269) {
		v298 = base.F64_sub(v256, v256)
		goto L83
	} else {
		goto L88
	}
L87:
	;
	v276 = F___cos(m, v256, float64(0))
	mBase = m.M
	v298 = v276
	goto L83
L88:
	;
	v280 = F___rem_pio2(m, v256, v262)
	mBase = m.M
	v281 = *(*float64)(unsafe.Add(mBase, uint32(v262)+8))
	v282 = *(*float64)(unsafe.Add(mBase, uint32(v262)))
	switch v280&int32(3) - int32(1) {
	case 0:
		goto L91
	case 1:
		goto L90
	case 2:
		goto L89
	default:
		goto L92
	}
L89:
	;
	v294 = F___sin(m, v282, v281, int32(1))
	mBase = m.M
	v298 = v294
	goto L83
L90:
	;
	v291 = F___cos(m, v282, v281)
	mBase = m.M
	v298 = base.F64_neg(v291)
	goto L83
L91:
	;
	v289 = F___sin(m, v282, v281, int32(1))
	mBase = m.M
	v298 = base.F64_neg(v289)
	goto L83
L92:
	;
	v287 = F___cos(m, v282, v281)
	mBase = m.M
	v298 = v287
	goto L83
L93:
	;
	v313 = v27 - v187
	goto L95
L94:
	;
	v313 = v187
	goto L95
L95:
	;
	if base.F64_le(v196, float64(60)) != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v430 = base.F64_div(v312, v429)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v430
	v432 = float64(0)
	v435 = *(*float64)(unsafe.Add(mBase, _consts[957]))
	v438 = base.F64_mul(base.F64_div(v430, v435), base.F64_convert_i32_s(v313))
	if base.F64_eq(v438, v432) != 0 {
		goto L124
	} else {
		goto L125
	}
L97:
	;
	v318 = base.F64_mul(v196, float64(0.017453292519943295))
	v322 = m.G0
	v324 = v322 - int32(16)
	m.G0 = v324
	v331 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v318))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v331) <= base.Ui32(int32(1072243195)) {
		goto L102
	} else {
		goto L103
	}
L98:
	;
	goto L99
L99:
	;
	v377 = base.F64_mul(base.F64_sub(float64(90), v196), float64(0.017453292519943295))
	v381 = m.G0
	v383 = v381 - int32(16)
	m.G0 = v383
	v390 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(v377))>>(uint(int64(32))%64))) & int32(2147483647)
	if base.Ui32(v390) <= base.Ui32(int32(1072243195)) {
		goto L113
	} else {
		goto L114
	}
L100:
	;
	v364 = base.F64_sub(float64(1), v360)
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v364
	v368 = *(*float64)(unsafe.Add(mBase, _consts[955]))
	v429 = base.F64_add(base.F64_mul(base.F64_div(v364, v368), float64(-0.5)), float64(1))
	goto L96
L101:
	;
	m.G0 = v324 + int32(16)
	goto L100
L102:
	;
	if base.Ui32(v331) < base.Ui32(int32(1044816030)) {
		v360 = float64(1)
		goto L101
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v331) {
		v360 = base.F64_sub(v318, v318)
		goto L101
	} else {
		goto L106
	}
L105:
	;
	v338 = F___cos(m, v318, float64(0))
	mBase = m.M
	v360 = v338
	goto L101
L106:
	;
	v342 = F___rem_pio2(m, v318, v324)
	mBase = m.M
	v343 = *(*float64)(unsafe.Add(mBase, uint32(v324)+8))
	v344 = *(*float64)(unsafe.Add(mBase, uint32(v324)))
	switch v342&int32(3) - int32(1) {
	case 0:
		goto L109
	case 1:
		goto L108
	case 2:
		goto L107
	default:
		goto L110
	}
L107:
	;
	v356 = F___sin(m, v344, v343, int32(1))
	mBase = m.M
	v360 = v356
	goto L101
L108:
	;
	v353 = F___cos(m, v344, v343)
	mBase = m.M
	v360 = base.F64_neg(v353)
	goto L101
L109:
	;
	v351 = F___sin(m, v344, v343, int32(1))
	mBase = m.M
	v360 = base.F64_neg(v351)
	goto L101
L110:
	;
	v349 = F___cos(m, v344, v343)
	mBase = m.M
	v360 = v349
	goto L101
L111:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v9)+8)) = v416
	v425 = *(*float64)(unsafe.Add(mBase, _consts[953]))
	v429 = base.F64_mul(base.F64_div(v416, v425), float64(0.5))
	goto L96
L112:
	;
	m.G0 = v383 + int32(16)
	goto L111
L113:
	;
	if base.Ui32(v390) < base.Ui32(int32(1045430272)) {
		v416 = v377
		goto L112
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if base.Ui32(int32(2146435072)) <= base.Ui32(v390) {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	v397 = F___sin(m, v377, float64(0), int32(0))
	mBase = m.M
	v416 = v397
	goto L112
L117:
	;
	v416 = base.F64_sub(v377, v377)
	goto L112
L118:
	;
	goto L119
L119:
	;
	v401 = F___rem_pio2(m, v377, v383)
	mBase = m.M
	v402 = *(*float64)(unsafe.Add(mBase, uint32(v383)+8))
	v403 = *(*float64)(unsafe.Add(mBase, uint32(v383)))
	switch v401&int32(3) - int32(1) {
	case 0:
		goto L122
	case 1:
		goto L121
	case 2:
		goto L120
	default:
		goto L123
	}
L120:
	;
	v414 = F___cos(m, v403, v402)
	mBase = m.M
	v416 = base.F64_neg(v414)
	goto L112
L121:
	;
	v412 = F___sin(m, v403, v402, int32(1))
	mBase = m.M
	v416 = base.F64_neg(v412)
	goto L112
L122:
	;
	v410 = F___cos(m, v403, v402)
	mBase = m.M
	v416 = v410
	goto L112
L123:
	;
	v409 = F___sin(m, v403, v402, int32(1))
	mBase = m.M
	v416 = v409
	goto L112
L124:
	;
	v441 = v432
	goto L126
L125:
	;
	v441 = v438
	goto L126
L126:
	;
	v447 = v441
	goto L4
L127:
	;
	return int32(0)
L128:
	;
	m.G0 = v9 + int32(16)
	return v448
L129:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v462 = m.ExcPending
	if v462 != 0 {
		goto L127
	} else {
		goto L130
	}
L130:
	;
	F_errmsg(m, int32(388004), int32(0))
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L127
	} else {
		goto L131
	}
L131:
	;
	F_errfinish(m, int32(476769), int32(2512), int32(412471))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L127
	} else {
		goto L132
	}
L132:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dtanh(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float64
	_ = v7
	var v8 int64
	_ = v8
	var v19 float64
	_ = v19
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v66 float64
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v80 float64
	_ = v80
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v83 float64
	_ = v83
	var v89 float64
	_ = v89
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v94 float64
	_ = v94
	var v95 float64
	_ = v95
	var v111 float64
	_ = v111
	var v114 float64
	_ = v114
	var v120 float64
	_ = v120
	var v129 float64
	_ = v129
	var v144 float64
	_ = v144
	var v153 float64
	_ = v153
	var v158 float64
	_ = v158
	var v165 float64
	_ = v165
	var v168 float64
	_ = v168
	var v174 float64
	_ = v174
	var v184 float64
	_ = v184
	var v186 float64
	_ = v186
	var v198 float64
	_ = v198
	var v205 float64
	_ = v205
	var v212 int64
	_ = v212
	var v217 int32
	_ = v217
	var v252 float64
	_ = v252
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 float64
	_ = v259
	var v266 float64
	_ = v266
	var v267 int32
	_ = v267
	var v268 float64
	_ = v268
	var v269 float64
	_ = v269
	var v275 float64
	_ = v275
	var v277 float64
	_ = v277
	var v278 int32
	_ = v278
	var v280 float64
	_ = v280
	var v281 float64
	_ = v281
	var v297 float64
	_ = v297
	var v300 float64
	_ = v300
	var v306 float64
	_ = v306
	var v315 float64
	_ = v315
	var v330 float64
	_ = v330
	var v339 float64
	_ = v339
	var v344 float64
	_ = v344
	var v351 float64
	_ = v351
	var v354 float64
	_ = v354
	var v360 float64
	_ = v360
	var v370 float64
	_ = v370
	var v372 float64
	_ = v372
	var v384 float64
	_ = v384
	var v391 float64
	_ = v391
	var v398 int64
	_ = v398
	var v403 int32
	_ = v403
	var v438 float64
	_ = v438
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v445 float64
	_ = v445
	var v452 float64
	_ = v452
	var v453 int32
	_ = v453
	var v454 float64
	_ = v454
	var v455 float64
	_ = v455
	var v461 float64
	_ = v461
	var v463 float64
	_ = v463
	var v464 int32
	_ = v464
	var v466 float64
	_ = v466
	var v467 float64
	_ = v467
	var v483 float64
	_ = v483
	var v486 float64
	_ = v486
	var v492 float64
	_ = v492
	var v501 float64
	_ = v501
	var v516 float64
	_ = v516
	var v525 float64
	_ = v525
	var v530 float64
	_ = v530
	var v537 float64
	_ = v537
	var v540 float64
	_ = v540
	var v546 float64
	_ = v546
	var v556 float64
	_ = v556
	var v558 float64
	_ = v558
	var v570 float64
	_ = v570
	var v575 float64
	_ = v575
	var v580 float64
	_ = v580
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = base.F64_abs(v6)
	v8 = base.I64_reinterpret_f64(v7)
	if base.Ui64(int64(4603122931675955200)) <= base.Ui64(v8) {
		if base.Ui64(int64(4626322721511309312)) <= base.Ui64(v8) {
			v575 = base.F64_add(base.F64_div(math.Float64frombits(uint64(0x8000000000000000)), v7), float64(1))
		} else {
			v19 = base.F64_add(v7, v7)
			v26 = base.I64_reinterpret_f64(v19)
			v31 = base.I32_wrap_i64(int64(base.Ui64(v26)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1078159482)) <= base.Ui32(v31) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v19)&int64(9223372036854775807)) {
					v186 = v19
					v198 = v186
				} else {
					if v26 < int64(0) {
						v198 = float64(-1)
					} else {
						if base.F64_gt(v19, float64(709.782712893384)) == int32(0) {
							v66 = base.F64_add(base.F64_mul(v19, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v19))
							if base.F64_lt(base.F64_abs(v66), float64(2.147483648e+09)) != 0 {
								v70 = base.I32_trunc_f64_s(v66)
								v72 = v70
							} else {
								v72 = int32(-2147483648)
							}
							v73 = base.F64_convert_i32_s(v72)
							v80 = base.F64_mul(v73, float64(1.9082149292705877e-10))
							v81 = v72
							v82 = base.F64_add(v19, base.F64_mul(v73, float64(-0.6931471803691238)))
							v83 = base.F64_sub(v82, v80)
							v89 = v83
							v91 = base.F64_sub(base.F64_sub(v82, v83), v80)
							v92 = v81
							v94 = base.F64_mul(v89, float64(0.5))
							v95 = base.F64_mul(v89, v94)
							v111 = base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v114 = base.F64_sub(float64(3), base.F64_mul(v111, v94))
							v120 = base.F64_mul(v95, base.F64_div(base.F64_sub(v111, v114), base.F64_sub(float64(6), base.F64_mul(v89, v114))))
							if v92 == int32(0) {
								v198 = base.F64_sub(v89, base.F64_sub(base.F64_mul(v89, v120), v95))
							} else {
								v129 = base.F64_sub(base.F64_sub(base.F64_mul(v89, base.F64_sub(v120, v91)), v91), v95)
								switch v92 + int32(1) {
								case 0:
									v198 = base.F64_add(base.F64_mul(base.F64_sub(v89, v129), float64(0.5)), float64(-0.5))
								default:
									v153 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v92+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v92) {
										v158 = base.F64_add(base.F64_sub(v89, v129), float64(1))
										if v92 == int32(1024) {
											v165 = base.F64_mul(base.F64_add(v158, v158), float64(8.98846567431158e+307))
										} else {
											v165 = base.F64_mul(v158, v153)
										}
										v198 = base.F64_add(v165, float64(-1))
									} else {
										v168 = float64(1)
										v174 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v92) << (uint(int64(52)) % 64))
										if base.Ui32(v92) <= base.Ui32(int32(19)) {
											v184 = base.F64_add(base.F64_sub(v168, v174), base.F64_sub(v89, v129))
										} else {
											v184 = base.F64_add(base.F64_sub(v89, base.F64_add(v129, v174)), v168)
										}
										v186 = base.F64_mul(v184, v153)
										v198 = v186
									}
								case 2:
									if base.F64_lt(v89, float64(-0.25)) != 0 {
										v198 = base.F64_mul(base.F64_sub(v129, base.F64_add(v89, float64(0.5))), float64(-2))
									} else {
										v144 = base.F64_sub(v89, v129)
										v198 = base.F64_add(base.F64_add(v144, v144), float64(1))
									}
								}
							}
						} else {
							v198 = base.F64_mul(v19, float64(8.98846567431158e+307))
						}
					}
				}
			} else {
				if base.Ui32(v31) < base.Ui32(int32(1071001155)) {
					if base.Ui32(v31) < base.Ui32(int32(1016070144)) {
						v186 = v19
						v198 = v186
					} else {
						v89 = v19
						v91 = float64(0)
						v92 = int32(0)
						v94 = base.F64_mul(v89, float64(0.5))
						v95 = base.F64_mul(v89, v94)
						v111 = base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v114 = base.F64_sub(float64(3), base.F64_mul(v111, v94))
						v120 = base.F64_mul(v95, base.F64_div(base.F64_sub(v111, v114), base.F64_sub(float64(6), base.F64_mul(v89, v114))))
						if v92 == int32(0) {
							v198 = base.F64_sub(v89, base.F64_sub(base.F64_mul(v89, v120), v95))
						} else {
							v129 = base.F64_sub(base.F64_sub(base.F64_mul(v89, base.F64_sub(v120, v91)), v91), v95)
							switch v92 + int32(1) {
							case 0:
								v198 = base.F64_add(base.F64_mul(base.F64_sub(v89, v129), float64(0.5)), float64(-0.5))
							default:
								v153 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v92+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v92) {
									v158 = base.F64_add(base.F64_sub(v89, v129), float64(1))
									if v92 == int32(1024) {
										v165 = base.F64_mul(base.F64_add(v158, v158), float64(8.98846567431158e+307))
									} else {
										v165 = base.F64_mul(v158, v153)
									}
									v198 = base.F64_add(v165, float64(-1))
								} else {
									v168 = float64(1)
									v174 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v92) << (uint(int64(52)) % 64))
									if base.Ui32(v92) <= base.Ui32(int32(19)) {
										v184 = base.F64_add(base.F64_sub(v168, v174), base.F64_sub(v89, v129))
									} else {
										v184 = base.F64_add(base.F64_sub(v89, base.F64_add(v129, v174)), v168)
									}
									v186 = base.F64_mul(v184, v153)
									v198 = v186
								}
							case 2:
								if base.F64_lt(v89, float64(-0.25)) != 0 {
									v198 = base.F64_mul(base.F64_sub(v129, base.F64_add(v89, float64(0.5))), float64(-2))
								} else {
									v144 = base.F64_sub(v89, v129)
									v198 = base.F64_add(base.F64_add(v144, v144), float64(1))
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(1072734897)) < base.Ui32(v31) {
						v66 = base.F64_add(base.F64_mul(v19, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v19))
						if base.F64_lt(base.F64_abs(v66), float64(2.147483648e+09)) != 0 {
							v70 = base.I32_trunc_f64_s(v66)
							v72 = v70
						} else {
							v72 = int32(-2147483648)
						}
						v73 = base.F64_convert_i32_s(v72)
						v80 = base.F64_mul(v73, float64(1.9082149292705877e-10))
						v81 = v72
						v82 = base.F64_add(v19, base.F64_mul(v73, float64(-0.6931471803691238)))
					} else {
						if int64(0) <= v26 {
							v80 = float64(1.9082149292705877e-10)
							v81 = int32(1)
							v82 = base.F64_add(v19, float64(-0.6931471803691238))
						} else {
							v80 = float64(-1.9082149292705877e-10)
							v81 = int32(-1)
							v82 = base.F64_add(v19, float64(0.6931471803691238))
						}
					}
					v83 = base.F64_sub(v82, v80)
					v89 = v83
					v91 = base.F64_sub(base.F64_sub(v82, v83), v80)
					v92 = v81
					v94 = base.F64_mul(v89, float64(0.5))
					v95 = base.F64_mul(v89, v94)
					v111 = base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, base.F64_add(base.F64_mul(v95, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v114 = base.F64_sub(float64(3), base.F64_mul(v111, v94))
					v120 = base.F64_mul(v95, base.F64_div(base.F64_sub(v111, v114), base.F64_sub(float64(6), base.F64_mul(v89, v114))))
					if v92 == int32(0) {
						v198 = base.F64_sub(v89, base.F64_sub(base.F64_mul(v89, v120), v95))
					} else {
						v129 = base.F64_sub(base.F64_sub(base.F64_mul(v89, base.F64_sub(v120, v91)), v91), v95)
						switch v92 + int32(1) {
						case 0:
							v198 = base.F64_add(base.F64_mul(base.F64_sub(v89, v129), float64(0.5)), float64(-0.5))
						default:
							v153 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v92+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v92) {
								v158 = base.F64_add(base.F64_sub(v89, v129), float64(1))
								if v92 == int32(1024) {
									v165 = base.F64_mul(base.F64_add(v158, v158), float64(8.98846567431158e+307))
								} else {
									v165 = base.F64_mul(v158, v153)
								}
								v198 = base.F64_add(v165, float64(-1))
							} else {
								v168 = float64(1)
								v174 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v92) << (uint(int64(52)) % 64))
								if base.Ui32(v92) <= base.Ui32(int32(19)) {
									v184 = base.F64_add(base.F64_sub(v168, v174), base.F64_sub(v89, v129))
								} else {
									v184 = base.F64_add(base.F64_sub(v89, base.F64_add(v129, v174)), v168)
								}
								v186 = base.F64_mul(v184, v153)
								v198 = v186
							}
						case 2:
							if base.F64_lt(v89, float64(-0.25)) != 0 {
								v198 = base.F64_mul(base.F64_sub(v129, base.F64_add(v89, float64(0.5))), float64(-2))
							} else {
								v144 = base.F64_sub(v89, v129)
								v198 = base.F64_add(base.F64_add(v144, v144), float64(1))
							}
						}
					}
				}
			}
			v575 = base.F64_sub(float64(1), base.F64_div(float64(2), base.F64_add(v198, float64(2))))
		}
	} else {
		if base.Ui64(int64(4598272728187797504)) <= base.Ui64(v8) {
			v205 = base.F64_add(v7, v7)
			v212 = base.I64_reinterpret_f64(v205)
			v217 = base.I32_wrap_i64(int64(base.Ui64(v212)>>(uint(int64(32))%64))) & int32(2147483647)
			if base.Ui32(int32(1078159482)) <= base.Ui32(v217) {
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v205)&int64(9223372036854775807)) {
					v372 = v205
					v384 = v372
				} else {
					if v212 < int64(0) {
						v384 = float64(-1)
					} else {
						if base.F64_gt(v205, float64(709.782712893384)) == int32(0) {
							v252 = base.F64_add(base.F64_mul(v205, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v205))
							if base.F64_lt(base.F64_abs(v252), float64(2.147483648e+09)) != 0 {
								v256 = base.I32_trunc_f64_s(v252)
								v258 = v256
							} else {
								v258 = int32(-2147483648)
							}
							v259 = base.F64_convert_i32_s(v258)
							v266 = base.F64_mul(v259, float64(1.9082149292705877e-10))
							v267 = v258
							v268 = base.F64_add(v205, base.F64_mul(v259, float64(-0.6931471803691238)))
							v269 = base.F64_sub(v268, v266)
							v275 = v269
							v277 = base.F64_sub(base.F64_sub(v268, v269), v266)
							v278 = v267
							v280 = base.F64_mul(v275, float64(0.5))
							v281 = base.F64_mul(v275, v280)
							v297 = base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v300 = base.F64_sub(float64(3), base.F64_mul(v297, v280))
							v306 = base.F64_mul(v281, base.F64_div(base.F64_sub(v297, v300), base.F64_sub(float64(6), base.F64_mul(v275, v300))))
							if v278 == int32(0) {
								v384 = base.F64_sub(v275, base.F64_sub(base.F64_mul(v275, v306), v281))
							} else {
								v315 = base.F64_sub(base.F64_sub(base.F64_mul(v275, base.F64_sub(v306, v277)), v277), v281)
								switch v278 + int32(1) {
								case 0:
									v384 = base.F64_add(base.F64_mul(base.F64_sub(v275, v315), float64(0.5)), float64(-0.5))
								default:
									v339 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v278+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v278) {
										v344 = base.F64_add(base.F64_sub(v275, v315), float64(1))
										if v278 == int32(1024) {
											v351 = base.F64_mul(base.F64_add(v344, v344), float64(8.98846567431158e+307))
										} else {
											v351 = base.F64_mul(v344, v339)
										}
										v384 = base.F64_add(v351, float64(-1))
									} else {
										v354 = float64(1)
										v360 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v278) << (uint(int64(52)) % 64))
										if base.Ui32(v278) <= base.Ui32(int32(19)) {
											v370 = base.F64_add(base.F64_sub(v354, v360), base.F64_sub(v275, v315))
										} else {
											v370 = base.F64_add(base.F64_sub(v275, base.F64_add(v315, v360)), v354)
										}
										v372 = base.F64_mul(v370, v339)
										v384 = v372
									}
								case 2:
									if base.F64_lt(v275, float64(-0.25)) != 0 {
										v384 = base.F64_mul(base.F64_sub(v315, base.F64_add(v275, float64(0.5))), float64(-2))
									} else {
										v330 = base.F64_sub(v275, v315)
										v384 = base.F64_add(base.F64_add(v330, v330), float64(1))
									}
								}
							}
						} else {
							v384 = base.F64_mul(v205, float64(8.98846567431158e+307))
						}
					}
				}
			} else {
				if base.Ui32(v217) < base.Ui32(int32(1071001155)) {
					if base.Ui32(v217) < base.Ui32(int32(1016070144)) {
						v372 = v205
						v384 = v372
					} else {
						v275 = v205
						v277 = float64(0)
						v278 = int32(0)
						v280 = base.F64_mul(v275, float64(0.5))
						v281 = base.F64_mul(v275, v280)
						v297 = base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v300 = base.F64_sub(float64(3), base.F64_mul(v297, v280))
						v306 = base.F64_mul(v281, base.F64_div(base.F64_sub(v297, v300), base.F64_sub(float64(6), base.F64_mul(v275, v300))))
						if v278 == int32(0) {
							v384 = base.F64_sub(v275, base.F64_sub(base.F64_mul(v275, v306), v281))
						} else {
							v315 = base.F64_sub(base.F64_sub(base.F64_mul(v275, base.F64_sub(v306, v277)), v277), v281)
							switch v278 + int32(1) {
							case 0:
								v384 = base.F64_add(base.F64_mul(base.F64_sub(v275, v315), float64(0.5)), float64(-0.5))
							default:
								v339 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v278+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v278) {
									v344 = base.F64_add(base.F64_sub(v275, v315), float64(1))
									if v278 == int32(1024) {
										v351 = base.F64_mul(base.F64_add(v344, v344), float64(8.98846567431158e+307))
									} else {
										v351 = base.F64_mul(v344, v339)
									}
									v384 = base.F64_add(v351, float64(-1))
								} else {
									v354 = float64(1)
									v360 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v278) << (uint(int64(52)) % 64))
									if base.Ui32(v278) <= base.Ui32(int32(19)) {
										v370 = base.F64_add(base.F64_sub(v354, v360), base.F64_sub(v275, v315))
									} else {
										v370 = base.F64_add(base.F64_sub(v275, base.F64_add(v315, v360)), v354)
									}
									v372 = base.F64_mul(v370, v339)
									v384 = v372
								}
							case 2:
								if base.F64_lt(v275, float64(-0.25)) != 0 {
									v384 = base.F64_mul(base.F64_sub(v315, base.F64_add(v275, float64(0.5))), float64(-2))
								} else {
									v330 = base.F64_sub(v275, v315)
									v384 = base.F64_add(base.F64_add(v330, v330), float64(1))
								}
							}
						}
					}
				} else {
					if base.Ui32(int32(1072734897)) < base.Ui32(v217) {
						v252 = base.F64_add(base.F64_mul(v205, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v205))
						if base.F64_lt(base.F64_abs(v252), float64(2.147483648e+09)) != 0 {
							v256 = base.I32_trunc_f64_s(v252)
							v258 = v256
						} else {
							v258 = int32(-2147483648)
						}
						v259 = base.F64_convert_i32_s(v258)
						v266 = base.F64_mul(v259, float64(1.9082149292705877e-10))
						v267 = v258
						v268 = base.F64_add(v205, base.F64_mul(v259, float64(-0.6931471803691238)))
					} else {
						if int64(0) <= v212 {
							v266 = float64(1.9082149292705877e-10)
							v267 = int32(1)
							v268 = base.F64_add(v205, float64(-0.6931471803691238))
						} else {
							v266 = float64(-1.9082149292705877e-10)
							v267 = int32(-1)
							v268 = base.F64_add(v205, float64(0.6931471803691238))
						}
					}
					v269 = base.F64_sub(v268, v266)
					v275 = v269
					v277 = base.F64_sub(base.F64_sub(v268, v269), v266)
					v278 = v267
					v280 = base.F64_mul(v275, float64(0.5))
					v281 = base.F64_mul(v275, v280)
					v297 = base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, base.F64_add(base.F64_mul(v281, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
					v300 = base.F64_sub(float64(3), base.F64_mul(v297, v280))
					v306 = base.F64_mul(v281, base.F64_div(base.F64_sub(v297, v300), base.F64_sub(float64(6), base.F64_mul(v275, v300))))
					if v278 == int32(0) {
						v384 = base.F64_sub(v275, base.F64_sub(base.F64_mul(v275, v306), v281))
					} else {
						v315 = base.F64_sub(base.F64_sub(base.F64_mul(v275, base.F64_sub(v306, v277)), v277), v281)
						switch v278 + int32(1) {
						case 0:
							v384 = base.F64_add(base.F64_mul(base.F64_sub(v275, v315), float64(0.5)), float64(-0.5))
						default:
							v339 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v278+int32(1023)) << (uint(int64(52)) % 64))
							if base.Ui32(int32(57)) <= base.Ui32(v278) {
								v344 = base.F64_add(base.F64_sub(v275, v315), float64(1))
								if v278 == int32(1024) {
									v351 = base.F64_mul(base.F64_add(v344, v344), float64(8.98846567431158e+307))
								} else {
									v351 = base.F64_mul(v344, v339)
								}
								v384 = base.F64_add(v351, float64(-1))
							} else {
								v354 = float64(1)
								v360 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v278) << (uint(int64(52)) % 64))
								if base.Ui32(v278) <= base.Ui32(int32(19)) {
									v370 = base.F64_add(base.F64_sub(v354, v360), base.F64_sub(v275, v315))
								} else {
									v370 = base.F64_add(base.F64_sub(v275, base.F64_add(v315, v360)), v354)
								}
								v372 = base.F64_mul(v370, v339)
								v384 = v372
							}
						case 2:
							if base.F64_lt(v275, float64(-0.25)) != 0 {
								v384 = base.F64_mul(base.F64_sub(v315, base.F64_add(v275, float64(0.5))), float64(-2))
							} else {
								v330 = base.F64_sub(v275, v315)
								v384 = base.F64_add(base.F64_add(v330, v330), float64(1))
							}
						}
					}
				}
			}
			v575 = base.F64_div(v384, base.F64_add(v384, float64(2)))
		} else {
			if base.Ui64(v8) < base.Ui64(int64(4503599627370496)) {
				v575 = v7
			} else {
				v391 = base.F64_mul(v7, float64(-2))
				v398 = base.I64_reinterpret_f64(v391)
				v403 = base.I32_wrap_i64(int64(base.Ui64(v398)>>(uint(int64(32))%64))) & int32(2147483647)
				if base.Ui32(int32(1078159482)) <= base.Ui32(v403) {
					if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v391)&int64(9223372036854775807)) {
						v558 = v391
						v570 = v558
					} else {
						if v398 < int64(0) {
							v570 = float64(-1)
						} else {
							if base.F64_gt(v391, float64(709.782712893384)) == int32(0) {
								v438 = base.F64_add(base.F64_mul(v391, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v391))
								if base.F64_lt(base.F64_abs(v438), float64(2.147483648e+09)) != 0 {
									v442 = base.I32_trunc_f64_s(v438)
									v444 = v442
								} else {
									v444 = int32(-2147483648)
								}
								v445 = base.F64_convert_i32_s(v444)
								v452 = base.F64_mul(v445, float64(1.9082149292705877e-10))
								v453 = v444
								v454 = base.F64_add(v391, base.F64_mul(v445, float64(-0.6931471803691238)))
								v455 = base.F64_sub(v454, v452)
								v461 = v455
								v463 = base.F64_sub(base.F64_sub(v454, v455), v452)
								v464 = v453
								v466 = base.F64_mul(v461, float64(0.5))
								v467 = base.F64_mul(v461, v466)
								v483 = base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
								v486 = base.F64_sub(float64(3), base.F64_mul(v483, v466))
								v492 = base.F64_mul(v467, base.F64_div(base.F64_sub(v483, v486), base.F64_sub(float64(6), base.F64_mul(v461, v486))))
								if v464 == int32(0) {
									v570 = base.F64_sub(v461, base.F64_sub(base.F64_mul(v461, v492), v467))
								} else {
									v501 = base.F64_sub(base.F64_sub(base.F64_mul(v461, base.F64_sub(v492, v463)), v463), v467)
									switch v464 + int32(1) {
									case 0:
										v570 = base.F64_add(base.F64_mul(base.F64_sub(v461, v501), float64(0.5)), float64(-0.5))
									default:
										v525 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v464+int32(1023)) << (uint(int64(52)) % 64))
										if base.Ui32(int32(57)) <= base.Ui32(v464) {
											v530 = base.F64_add(base.F64_sub(v461, v501), float64(1))
											if v464 == int32(1024) {
												v537 = base.F64_mul(base.F64_add(v530, v530), float64(8.98846567431158e+307))
											} else {
												v537 = base.F64_mul(v530, v525)
											}
											v570 = base.F64_add(v537, float64(-1))
										} else {
											v540 = float64(1)
											v546 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v464) << (uint(int64(52)) % 64))
											if base.Ui32(v464) <= base.Ui32(int32(19)) {
												v556 = base.F64_add(base.F64_sub(v540, v546), base.F64_sub(v461, v501))
											} else {
												v556 = base.F64_add(base.F64_sub(v461, base.F64_add(v501, v546)), v540)
											}
											v558 = base.F64_mul(v556, v525)
											v570 = v558
										}
									case 2:
										if base.F64_lt(v461, float64(-0.25)) != 0 {
											v570 = base.F64_mul(base.F64_sub(v501, base.F64_add(v461, float64(0.5))), float64(-2))
										} else {
											v516 = base.F64_sub(v461, v501)
											v570 = base.F64_add(base.F64_add(v516, v516), float64(1))
										}
									}
								}
							} else {
								v570 = base.F64_mul(v391, float64(8.98846567431158e+307))
							}
						}
					}
				} else {
					if base.Ui32(v403) < base.Ui32(int32(1071001155)) {
						if base.Ui32(v403) < base.Ui32(int32(1016070144)) {
							v558 = v391
							v570 = v558
						} else {
							v461 = v391
							v463 = float64(0)
							v464 = int32(0)
							v466 = base.F64_mul(v461, float64(0.5))
							v467 = base.F64_mul(v461, v466)
							v483 = base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
							v486 = base.F64_sub(float64(3), base.F64_mul(v483, v466))
							v492 = base.F64_mul(v467, base.F64_div(base.F64_sub(v483, v486), base.F64_sub(float64(6), base.F64_mul(v461, v486))))
							if v464 == int32(0) {
								v570 = base.F64_sub(v461, base.F64_sub(base.F64_mul(v461, v492), v467))
							} else {
								v501 = base.F64_sub(base.F64_sub(base.F64_mul(v461, base.F64_sub(v492, v463)), v463), v467)
								switch v464 + int32(1) {
								case 0:
									v570 = base.F64_add(base.F64_mul(base.F64_sub(v461, v501), float64(0.5)), float64(-0.5))
								default:
									v525 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v464+int32(1023)) << (uint(int64(52)) % 64))
									if base.Ui32(int32(57)) <= base.Ui32(v464) {
										v530 = base.F64_add(base.F64_sub(v461, v501), float64(1))
										if v464 == int32(1024) {
											v537 = base.F64_mul(base.F64_add(v530, v530), float64(8.98846567431158e+307))
										} else {
											v537 = base.F64_mul(v530, v525)
										}
										v570 = base.F64_add(v537, float64(-1))
									} else {
										v540 = float64(1)
										v546 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v464) << (uint(int64(52)) % 64))
										if base.Ui32(v464) <= base.Ui32(int32(19)) {
											v556 = base.F64_add(base.F64_sub(v540, v546), base.F64_sub(v461, v501))
										} else {
											v556 = base.F64_add(base.F64_sub(v461, base.F64_add(v501, v546)), v540)
										}
										v558 = base.F64_mul(v556, v525)
										v570 = v558
									}
								case 2:
									if base.F64_lt(v461, float64(-0.25)) != 0 {
										v570 = base.F64_mul(base.F64_sub(v501, base.F64_add(v461, float64(0.5))), float64(-2))
									} else {
										v516 = base.F64_sub(v461, v501)
										v570 = base.F64_add(base.F64_add(v516, v516), float64(1))
									}
								}
							}
						}
					} else {
						if base.Ui32(int32(1072734897)) < base.Ui32(v403) {
							v438 = base.F64_add(base.F64_mul(v391, float64(1.4426950408889634)), base.F64_copysign(float64(0.5), v391))
							if base.F64_lt(base.F64_abs(v438), float64(2.147483648e+09)) != 0 {
								v442 = base.I32_trunc_f64_s(v438)
								v444 = v442
							} else {
								v444 = int32(-2147483648)
							}
							v445 = base.F64_convert_i32_s(v444)
							v452 = base.F64_mul(v445, float64(1.9082149292705877e-10))
							v453 = v444
							v454 = base.F64_add(v391, base.F64_mul(v445, float64(-0.6931471803691238)))
						} else {
							if int64(0) <= v398 {
								v452 = float64(1.9082149292705877e-10)
								v453 = int32(1)
								v454 = base.F64_add(v391, float64(-0.6931471803691238))
							} else {
								v452 = float64(-1.9082149292705877e-10)
								v453 = int32(-1)
								v454 = base.F64_add(v391, float64(0.6931471803691238))
							}
						}
						v455 = base.F64_sub(v454, v452)
						v461 = v455
						v463 = base.F64_sub(base.F64_sub(v454, v455), v452)
						v464 = v453
						v466 = base.F64_mul(v461, float64(0.5))
						v467 = base.F64_mul(v461, v466)
						v483 = base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, base.F64_add(base.F64_mul(v467, float64(-2.0109921818362437e-07)), float64(4.008217827329362e-06))), float64(-7.93650757867488e-05))), float64(0.0015873015872548146))), float64(-0.03333333333333313))), float64(1))
						v486 = base.F64_sub(float64(3), base.F64_mul(v483, v466))
						v492 = base.F64_mul(v467, base.F64_div(base.F64_sub(v483, v486), base.F64_sub(float64(6), base.F64_mul(v461, v486))))
						if v464 == int32(0) {
							v570 = base.F64_sub(v461, base.F64_sub(base.F64_mul(v461, v492), v467))
						} else {
							v501 = base.F64_sub(base.F64_sub(base.F64_mul(v461, base.F64_sub(v492, v463)), v463), v467)
							switch v464 + int32(1) {
							case 0:
								v570 = base.F64_add(base.F64_mul(base.F64_sub(v461, v501), float64(0.5)), float64(-0.5))
							default:
								v525 = base.F64_reinterpret_i64(base.I64_extend_i32_u(v464+int32(1023)) << (uint(int64(52)) % 64))
								if base.Ui32(int32(57)) <= base.Ui32(v464) {
									v530 = base.F64_add(base.F64_sub(v461, v501), float64(1))
									if v464 == int32(1024) {
										v537 = base.F64_mul(base.F64_add(v530, v530), float64(8.98846567431158e+307))
									} else {
										v537 = base.F64_mul(v530, v525)
									}
									v570 = base.F64_add(v537, float64(-1))
								} else {
									v540 = float64(1)
									v546 = base.F64_reinterpret_i64(base.I64_extend_i32_u(int32(1023)-v464) << (uint(int64(52)) % 64))
									if base.Ui32(v464) <= base.Ui32(int32(19)) {
										v556 = base.F64_add(base.F64_sub(v540, v546), base.F64_sub(v461, v501))
									} else {
										v556 = base.F64_add(base.F64_sub(v461, base.F64_add(v501, v546)), v540)
									}
									v558 = base.F64_mul(v556, v525)
									v570 = v558
								}
							case 2:
								if base.F64_lt(v461, float64(-0.25)) != 0 {
									v570 = base.F64_mul(base.F64_sub(v501, base.F64_add(v461, float64(0.5))), float64(-2))
								} else {
									v516 = base.F64_sub(v461, v501)
									v570 = base.F64_add(base.F64_add(v516, v516), float64(1))
								}
							}
						}
					}
				}
				v575 = base.F64_div(base.F64_neg(v570), base.F64_add(v570, float64(2)))
			}
		}
	}
	if base.I64_reinterpret_f64(v6) < int64(0) {
		v580 = base.F64_neg(v575)
	} else {
		v580 = v575
	}
	if base.F64_eq(base.F64_abs(v580), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		F_float_overflow_error(m)
		mBase = m.M
		v587 = m.ExcPending
		if v587 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v588 = F_Float8GetDatum(m, v580)
		mBase = m.M
		v589 = m.ExcPending
		if v589 != 0 {
			return int32(0)
		} else {
			return v588
		}
	}
}
func F_dtof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 float32
	_ = v6
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = base.F32_demote_f64(v5)
	if base.F32_eq(base.F32_abs(v6), math.Float32frombits(uint32(0x7f800000)))&base.F64_ne(base.F64_abs(v5), math.Float64frombits(uint64(0x7ff0000000000000))) == int32(0) {
		if base.F32_eq(v6, float32(0))&base.F64_ne(v5, float64(0)) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.I32_reinterpret_f32(v6)
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_dump_cursor_direction(m *base.Module, l0 int32) {
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
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = int32(4548456)
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1256]))
	*(*int32)(unsafe.Add(mBase, _consts[1256])) = v10 + int32(2)
	if int32(-1) <= v10 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v18 = int32(0)
	goto L4
L2:
	;
	goto L3
L3:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	switch v37 {
	case 0:
		goto L14
	case 1:
		goto L13
	case 2:
		goto L12
	case 3:
		goto L11
	default:
		goto L10
	}
L4:
	;
	F_pg_printf(m, int32(711394), int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	v28 = v18 + int32(1)
	v31 = *(*int32)(unsafe.Add(mBase, _consts[1256]))
	if v28 < v31 {
		v18 = v28
		goto L4
	} else {
		goto L8
	}
L8:
	;
	goto L5
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v60 != 0 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = v37
	F_pg_printf(m, int32(456733), v5+int32(-16))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L6
	} else {
		goto L19
	}
L11:
	;
	F_pg_printf(m, int32(709898), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L6
	} else {
		goto L18
	}
L12:
	;
	F_pg_printf(m, int32(709984), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L17
	}
L13:
	;
	F_pg_printf(m, int32(710138), int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L16
	}
L14:
	;
	F_pg_printf(m, int32(710125), int32(0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L9
L16:
	;
	goto L9
L17:
	;
	goto L9
L18:
	;
	goto L9
L19:
	;
	goto L9
L20:
	;
	v100 = int32(4548456)
	v101 = *(*int32)(unsafe.Add(mBase, _consts[1256]))
	*(*int32)(unsafe.Add(mBase, _consts[1256])) = v101 - int32(2)
	m.G0 = v7 - int32(-64)
	return
L21:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v60)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v61
	F_pg_printf(m, int32(652483), v5+int32(-32))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L6
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v92
	F_pg_printf(m, int32(715197), v7)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L6
	} else {
		goto L33
	}
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+16))
	if int32(0) <= v68 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+20)))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v68
	if v71 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	F_pg_printf(m, int32(722305), int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L6
	} else {
		goto L32
	}
L28:
	;
	v78 = int32(641616)
	goto L30
L29:
	;
	v78 = int32(722455)
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v78
	F_pg_printf(m, int32(168443), v5+int32(-48))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L6
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	goto L20
L33:
	;
	goto L20
}
func F_dumptuples(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
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
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int64
	_ = v214
	var v217 int64
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v14 <= v13 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L21
	} else {
		goto L58
	}
L2:
	;
	m.G0 = v11 - int32(-64)
	return
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	if v13 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	if v16 < int64(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	if l1 != 0 {
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	if l1 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L8:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+145)))
	if v19&int32(1) == int32(0) {
		goto L3
	} else {
		goto L9
	}
L9:
	;
	goto L2
L10:
	;
	goto L3
L11:
	;
	v76 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v73))) = v72 + v76
	v80 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1185])))
	if v80 != v76 {
		goto L23
	} else {
		goto L24
	}
L12:
	;
	if int32(0) < v26 {
		goto L2
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v26 == int32(2147483647) {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v72 = v26
	v73 = l0 + int32(168)
	goto L11
L16:
	;
	v36 = l0 + int32(168)
	if v26 <= int32(0) {
		v72 = v26
		v73 = v36
		goto L11
	} else {
		goto L17
	}
L17:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	if v39 < v40 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v43 = F_LogicalTapeCreate(m, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v63 = base.I32_rem_s(v62, v39)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v61+v63<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v62 + int32(1)
	v72 = v26
	v73 = v36
	goto L11
L21:
	;
	return
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+196)) = v43
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	*(*int32)(unsafe.Add(mBase, uint32(v46+v47<<(uint(int32(2))%32)))) = v43
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v53 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+188)) = v52 + v53
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+192)) = v56 + v53
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v72 = v60
	v73 = v36
	goto L11
L23:
	;
	F_tuplesort_sort_memtuples(m, l0)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L21
	} else {
		goto L30
	}
L24:
	;
	v85 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	if v85 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v93 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L21
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+40)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v89
	F_errmsg_internal(m, int32(196638), v9+int32(-32))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L21
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(476118), int32(2351), int32(156846))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L21
	} else {
		goto L29
	}
L29:
	;
	goto L23
L30:
	;
	v113 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1185])))
	if v113 != int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v143 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L32:
	;
	v118 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L21
	} else {
		goto L33
	}
L33:
	;
	if v118 == int32(0) {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v126 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L21
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v126
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v123
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v122
	F_errmsg_internal(m, int32(196681), v9+int32(-48))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L21
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(476118), int32(2362), int32(156846))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L21
	} else {
		goto L37
	}
L37:
	;
	goto L31
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_MemoryContextReset(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L21
	} else {
		goto L50
	}
L39:
	;
	v146 = int32(1)
	v148 = int32(0)
	if v143 != v146 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v155 = v148
	v157 = int32(0)
	goto L43
L41:
	;
	v184 = v148
	goto L42
L42:
	;
	if v143&v146 == int32(0) {
		goto L38
	} else {
		goto L48
	}
L43:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v164 = v155 << (uint(int32(4)) % 32)
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	m.T0[v167].(func(*base.Module, int32, int32, int32))(m, l0, v162, v164+v165)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L21
	} else {
		goto L45
	}
L44:
	;
	v184 = v179
	goto L42
L45:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	m.T0[v175].(func(*base.Module, int32, int32, int32))(m, l0, v170, v171+v164+int32(16))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L21
	} else {
		goto L46
	}
L46:
	;
	v178 = int32(2)
	v179 = v155 + v178
	v181 = v157 + v178
	if v181 != v143&int32(-2) {
		v155 = v179
		v157 = v181
		goto L43
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	m.T0[v198].(func(*base.Module, int32, int32, int32))(m, l0, v193, v194+v184<<(uint(int32(4))%32))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L21
	} else {
		goto L49
	}
L49:
	;
	goto L38
L50:
	;
	v214 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(0)
	v217 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v214 + v217
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+196))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+60)) = int32(0)
	F_LogicalTapeWrite(m, v220, v9+int32(-4), int32(4))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L21
	} else {
		goto L51
	}
L51:
	;
	v229 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1185])))
	if v229 != int32(1) {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v234 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	if v234 == int32(0) {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+232))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v243 = F_pg_rusage_show(m, l0+int32(256))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L21
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v243
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v240
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v239
	v248 = int32(1)
	v250 = base.I32_rem_s(v240-v248, v238)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v250 + v248
	F_errmsg_internal(m, int32(196862), v11)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L21
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(476118), int32(2395), int32(156846))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L21
	} else {
		goto L57
	}
L57:
	;
	goto L2
L58:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L21
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(2147483647)
	F_errmsg(m, int32(76230), v9+int32(-16))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L21
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(476118), int32(2341), int32(156846))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L21
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_dup2(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	if l0 != l1 {
		for {
			v9 = m.Env.X__syscall_dup3(m, l0, l1, int32(0))
			mBase = m.M
			if v9 == int32(-10) {
				continue
			} else {
				break
			}
			break
		}
		v30 = v9
		if base.Ui32(int32(-4095)) <= base.Ui32(v30) {
			*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v30
			v38 = int32(-1)
		} else {
			v38 = v30
		}
		v39 = v38
	} else {
		v12 = m.G0
		v14 = v12 - int32(32)
		m.G0 = v14
		v18 = m.Wasi_snapshot_preview1.Fd_fdstat_get(m, l0, v14+int32(8))
		mBase = m.M
		if v18 != 0 {
			*(*int32)(unsafe.Add(mBase, _consts[40])) = v18
			v23 = int32(0)
		} else {
			v23 = int32(1)
		}
		m.G0 = v14 + int32(32)
		if v23 != 0 {
			v39 = l0
		} else {
			v30 = int32(-8)
			if base.Ui32(int32(-4095)) <= base.Ui32(v30) {
				*(*int32)(unsafe.Add(mBase, _consts[40])) = int32(0) - v30
				v38 = int32(-1)
			} else {
				v38 = v30
			}
			v39 = v38
		}
	}
	return v39
}
func F_duptraverse(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v26 == int32(0) {
		goto L9
	} else {
		goto L17
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l2
	goto L11
L13:
	;
	goto L14
L14:
	;
	v20 = F_newstate(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v20
	if v20 == int32(0) {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	v31 = v26
	goto L18
L18:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v33 != 0 {
		goto L9
	} else {
		goto L20
	}
L19:
	;
	goto L9
L20:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	F_duptraverse(m, l0, v34, int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if v39 != 0 {
		goto L9
	} else {
		goto L22
	}
L22:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+24))
	F_cparc(m, l0, v31, v40, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v45 != 0 {
		v31 = v45
		goto L18
	} else {
		goto L24
	}
L24:
	;
	goto L19
}
func F_dutch_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v236 int32
	_ = v236
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v274 int32
	_ = v274
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
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v369 int32
	_ = v369
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v407 int32
	_ = v407
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v592 int32
	_ = v592
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v614 int32
	_ = v614
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v630 int32
	_ = v630
	var v643 int32
	_ = v643
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v669 int32
	_ = v669
	var v676 int32
	_ = v676
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v714 int32
	_ = v714
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v727 int32
	_ = v727
	var v730 int32
	_ = v730
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v765 int32
	_ = v765
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v785 int32
	_ = v785
	var v791 int32
	_ = v791
	var v799 int32
	_ = v799
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v819 int32
	_ = v819
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 int32
	_ = v834
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v864 int32
	_ = v864
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v880 int32
	_ = v880
	var v893 int32
	_ = v893
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v913 int32
	_ = v913
	var v919 int32
	_ = v919
	var v926 int32
	_ = v926
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v964 int32
	_ = v964
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1002 int32
	_ = v1002
	var v1015 int32
	_ = v1015
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1049 int32
	_ = v1049
	var v1060 int32
	_ = v1060
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1134 int32
	_ = v1134
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1158 int32
	_ = v1158
	var v1160 int32
	_ = v1160
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1207 int32
	_ = v1207
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1220 int32
	_ = v1220
	var v1226 int32
	_ = v1226
	var v1243 int32
	_ = v1243
	var v1250 int32
	_ = v1250
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1275 int32
	_ = v1275
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1286 int32
	_ = v1286
	var v1289 int32
	_ = v1289
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1298 int32
	_ = v1298
	var v1299 int32
	_ = v1299
	var v1301 int32
	_ = v1301
	var v1303 int32
	_ = v1303
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1311 int32
	_ = v1311
	var v1314 int32
	_ = v1314
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1321 int32
	_ = v1321
	var v1324 int32
	_ = v1324
	var v1327 int32
	_ = v1327
	var v1331 int32
	_ = v1331
	var v1334 int32
	_ = v1334
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1345 int32
	_ = v1345
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1357 int32
	_ = v1357
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1363 int32
	_ = v1363
	var v1365 int32
	_ = v1365
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1390 int32
	_ = v1390
	var v1393 int32
	_ = v1393
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1404 int32
	_ = v1404
	var v1407 int32
	_ = v1407
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1429 int32
	_ = v1429
	var v1433 int32
	_ = v1433
	var v1437 int32
	_ = v1437
	var v1438 int32
	_ = v1438
	var v1441 int32
	_ = v1441
	var v1442 int32
	_ = v1442
	var v1444 int32
	_ = v1444
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1460 int32
	_ = v1460
	var v1461 int32
	_ = v1461
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1490 int32
	_ = v1490
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1529 int32
	_ = v1529
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1549 int32
	_ = v1549
	var v1551 int32
	_ = v1551
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1598 int32
	_ = v1598
	var v1606 int32
	_ = v1606
	var v1607 int32
	_ = v1607
	var v1611 int32
	_ = v1611
	var v1617 int32
	_ = v1617
	var v1634 int32
	_ = v1634
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1645 int32
	_ = v1645
	var v1647 int32
	_ = v1647
	var v1649 int32
	_ = v1649
	var v1660 int32
	_ = v1660
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1680 int32
	_ = v1680
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1701 int32
	_ = v1701
	var v1702 int32
	_ = v1702
	var v1704 int32
	_ = v1704
	var v1706 int32
	_ = v1706
	var v1712 int32
	_ = v1712
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1731 int32
	_ = v1731
	var v1733 int32
	_ = v1733
	var v1735 int32
	_ = v1735
	var v1753 int32
	_ = v1753
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1772 int32
	_ = v1772
	var v1789 int32
	_ = v1789
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1824 int32
	_ = v1824
	var v1829 int32
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1836 int32
	_ = v1836
	var v1840 int32
	_ = v1840
	var v1854 int32
	_ = v1854
	var v1859 int32
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1868 int32
	_ = v1868
	var v1871 int32
	_ = v1871
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1886 int32
	_ = v1886
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1892 int32
	_ = v1892
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1908 int32
	_ = v1908
	var v1915 int32
	_ = v1915
	var v1917 int32
	_ = v1917
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1931 int32
	_ = v1931
	var v1934 int32
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1945 int32
	_ = v1945
	var v1946 int32
	_ = v1946
	var v1960 int32
	_ = v1960
	var v1966 int32
	_ = v1966
	var v1970 int32
	_ = v1970
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v7
	goto L2
L1:
	;
	return v1970
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v9
	v16 = v9 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 <= v16 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v7 == v138 {
		v157 = v7
		goto L51
	} else {
		goto L52
	}
L4:
	;
	goto L3
L5:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v9 = v135
	goto L2
L6:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L32
L7:
	;
	v34 = F_find_among(m, l0, int32(4212000), int32(11))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v9
	v73 = v9
	v74 = v17
	goto L6
L9:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v16))))
	if v21&int32(224) != int32(160) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if int32(1)<<(uint(v21)%32)&int32(340306450) != 0 {
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	return int32(0)
L13:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v38
	switch v34 - int32(1) {
	case 0:
		goto L19
	case 1:
		goto L18
	case 2:
		goto L17
	case 3:
		goto L16
	case 4:
		goto L15
	case 5:
		goto L14
	default:
		goto L5
	}
L14:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v73 = v38
	v74 = v72
	goto L6
L15:
	;
	v68 = F_slice_from_s(m, l0, int32(1), int32(2148661))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L28
	}
L16:
	;
	v62 = F_slice_from_s(m, l0, int32(1), int32(2148660))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L12
	} else {
		goto L26
	}
L17:
	;
	v56 = F_slice_from_s(m, l0, int32(1), int32(2148659))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L12
	} else {
		goto L24
	}
L18:
	;
	v50 = F_slice_from_s(m, l0, int32(1), int32(2148658))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L22
	}
L19:
	;
	v44 = F_slice_from_s(m, l0, int32(1), int32(2148657))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	if int32(0) <= v44 {
		goto L5
	} else {
		goto L21
	}
L21:
	;
	v1970 = v44
	goto L1
L22:
	;
	if int32(0) <= v50 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v1970 = v50
	goto L1
L24:
	;
	if int32(0) <= v56 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	v1970 = v56
	goto L1
L26:
	;
	if int32(0) <= v62 {
		goto L5
	} else {
		goto L27
	}
L27:
	;
	v1970 = v62
	goto L1
L28:
	;
	if int32(0) <= v68 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v1970 = v68
	goto L1
L30:
	;
	if v128 < int32(0) {
		goto L4
	} else {
		goto L50
	}
L32:
	;
	goto L33
L33:
	;
	goto L34
L34:
	;
	v83 = v73
	v85 = int32(1)
	goto L37
L36:
	;
	v128 = v113
	goto L30
L37:
	;
	if v74 <= v83 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	goto L36
L39:
	;
	v128 = int32(-1)
	goto L30
L40:
	;
	goto L41
L41:
	;
	v90 = v83 + int32(1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v83))))
	if base.Ui32(v92) < base.Ui32(int32(192)) {
		v113 = v90
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v114 = int32(1)
	if v114 < v85 {
		v83 = v113
		v85 = v85 - v114
		goto L37
	} else {
		goto L49
	}
L43:
	;
	if v74 <= v90 {
		v113 = v90
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v99 = v90
	goto L45
L45:
	;
	v102 = int32(*(*int8)(unsafe.Add(mBase, uint32(v76+v99))))
	if int32(-65) < v102 {
		v113 = v99
		goto L42
	} else {
		goto L47
	}
L46:
	;
	v113 = v74
	goto L42
L47:
	;
	v106 = v99 + int32(1)
	if v106 != v74 {
		v99 = v106
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	goto L38
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v128
	goto L5
L51:
	;
	v163 = v157
	goto L56
L52:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140+v7))))
	if v142 != int32(121) {
		v157 = v7
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v145 = int32(1)
	v146 = v7 + v145
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v146
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v146
	v151 = F_slice_from_s(m, l0, v145, int32(2148662))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L12
	} else {
		goto L54
	}
L54:
	;
	if v151 < int32(0) {
		v1970 = v151
		goto L1
	} else {
		goto L55
	}
L55:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v157 = v155
	goto L51
L56:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L61
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v507)+4)) = v508
	*(*int32)(unsafe.Add(mBase, uint32(v507)+8)) = v508
	v511 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	goto L153
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v163
	goto L129
L59:
	;
	if v281 != 0 {
		goto L83
	} else {
		goto L84
	}
L60:
	;
	v281 = v274
	goto L59
L61:
	;
	if v176 <= v175 {
		goto L63
	} else {
		goto L64
	}
L62:
	;
	v274 = int32(0)
	goto L60
L63:
	;
	v281 = int32(-1)
	goto L59
L64:
	;
	goto L65
L65:
	;
	v192 = int32(1)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175+v177))))
	if base.Ui32(v194) < base.Ui32(int32(192)) {
		v251 = v194
		v252 = v192
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if int32(232) < v251 {
		v274 = v252
		goto L60
	} else {
		goto L79
	}
L67:
	;
	v198 = v175 + int32(1)
	if v198 == v176 {
		v251 = v194
		v252 = v192
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v177))))
	v203 = v201 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v194) {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207+v177))))
	v219 = v217 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v194) {
		goto L75
	} else {
		goto L76
	}
L70:
	;
	v207 = v175 + int32(2)
	if v207 != v176 {
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v251 = v194<<(uint(int32(6))%32)&int32(1984) | v203
	v252 = int32(2)
	goto L66
L73:
	;
	goto L72
L74:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v223))))
	v251 = v236&int32(63) | (v194<<(uint(int32(18))%32)&int32(1835008) | v203<<(uint(int32(12))%32) | v219<<(uint(int32(6))%32))
	v252 = int32(4)
	goto L66
L75:
	;
	v223 = v175 + int32(3)
	if v223 != v176 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v251 = v194<<(uint(int32(12))%32)&int32(61440) | v203<<(uint(int32(6))%32) | v219
	v252 = int32(3)
	goto L66
L78:
	;
	goto L77
L79:
	;
	v256 = v251 - int32(97)
	if v256 < int32(0) {
		v274 = v252
		goto L60
	} else {
		goto L80
	}
L80:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v256)>>(uint(int32(3))%32)))+uint32(_consts[1274]))))
	if int32(base.Ui32(v262)>>(uint(v256&int32(7))%32))&int32(1) == int32(0) {
		v274 = v252
		goto L60
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v252 + v175
	goto L82
L82:
	;
	goto L62
L83:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v283 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v448 = v283
	v449 = v282
	goto L58
L84:
	;
	goto L85
L85:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v284
	v286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v287 == v284 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v448 = v286
	v449 = v284
	goto L58
L87:
	;
	goto L88
L88:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v286))))
	if v290 == int32(105) {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v163
	goto L56
L90:
	;
	v439 = F_slice_from_s(m, l0, int32(1), int32(2148689))
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L12
	} else {
		goto L125
	}
L91:
	;
	v294 = v284 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v294
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v294
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L96
L92:
	;
	v419 = v286
	v420 = v287
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v284
	if v284 == v420 {
		goto L119
	} else {
		goto L120
	}
L94:
	;
	if v414 == int32(0) {
		goto L90
	} else {
		goto L118
	}
L95:
	;
	v414 = v407
	goto L94
L96:
	;
	if v309 <= v294 {
		goto L98
	} else {
		goto L99
	}
L97:
	;
	v407 = int32(0)
	goto L95
L98:
	;
	v414 = int32(-1)
	goto L94
L99:
	;
	goto L100
L100:
	;
	v325 = int32(1)
	v327 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294+v310))))
	if base.Ui32(v327) < base.Ui32(int32(192)) {
		v384 = v327
		v385 = v325
		goto L101
	} else {
		goto L102
	}
L101:
	;
	if int32(232) < v384 {
		v407 = v385
		goto L95
	} else {
		goto L114
	}
L102:
	;
	v331 = v284 + int32(2)
	if v331 == v309 {
		v384 = v327
		v385 = v325
		goto L101
	} else {
		goto L103
	}
L103:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331+v310))))
	v336 = v334 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v327) {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v340+v310))))
	v352 = v350 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v327) {
		goto L110
	} else {
		goto L111
	}
L105:
	;
	v340 = v284 + int32(3)
	if v340 != v309 {
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v384 = v327<<(uint(int32(6))%32)&int32(1984) | v336
	v385 = int32(2)
	goto L101
L108:
	;
	goto L107
L109:
	;
	v369 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310+v356))))
	v384 = v369&int32(63) | (v327<<(uint(int32(18))%32)&int32(1835008) | v336<<(uint(int32(12))%32) | v352<<(uint(int32(6))%32))
	v385 = int32(4)
	goto L101
L110:
	;
	v356 = v284 + int32(4)
	if v356 != v309 {
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v384 = v327<<(uint(int32(12))%32)&int32(61440) | v336<<(uint(int32(6))%32) | v352
	v385 = int32(3)
	goto L101
L113:
	;
	goto L112
L114:
	;
	v389 = v384 - int32(97)
	if v389 < int32(0) {
		v407 = v385
		goto L95
	} else {
		goto L115
	}
L115:
	;
	v395 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v389)>>(uint(int32(3))%32)))+uint32(_consts[1274]))))
	if int32(base.Ui32(v395)>>(uint(v389&int32(7))%32))&int32(1) == int32(0) {
		v407 = v385
		goto L95
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v385 + v294
	goto L117
L117:
	;
	goto L97
L118:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v419 = v417
	v420 = v418
	goto L93
L119:
	;
	v448 = v419
	v449 = v284
	goto L58
L120:
	;
	goto L121
L121:
	;
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v419))))
	if v424 != int32(121) {
		v448 = v419
		v449 = v420
		goto L58
	} else {
		goto L122
	}
L122:
	;
	v427 = int32(1)
	v428 = v284 + v427
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v428
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v428
	v433 = F_slice_from_s(m, l0, v427, int32(2148690))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L12
	} else {
		goto L123
	}
L123:
	;
	if int32(0) <= v433 {
		goto L89
	} else {
		goto L124
	}
L124:
	;
	v1970 = v433
	goto L1
L125:
	;
	if v439 < int32(0) {
		v1970 = v439
		goto L1
	} else {
		goto L126
	}
L126:
	;
	goto L89
L127:
	;
	if int32(0) <= v502 {
		goto L147
	} else {
		goto L148
	}
L129:
	;
	goto L130
L130:
	;
	goto L131
L131:
	;
	v457 = v163
	v459 = int32(1)
	goto L134
L133:
	;
	v502 = v487
	goto L127
L134:
	;
	if v449 <= v457 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L133
L136:
	;
	v502 = int32(-1)
	goto L127
L137:
	;
	goto L138
L138:
	;
	v464 = v457 + int32(1)
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v448+v457))))
	if base.Ui32(v466) < base.Ui32(int32(192)) {
		v487 = v464
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v488 = int32(1)
	if v488 < v459 {
		v457 = v487
		v459 = v459 - v488
		goto L134
	} else {
		goto L146
	}
L140:
	;
	if v449 <= v464 {
		v487 = v464
		goto L139
	} else {
		goto L141
	}
L141:
	;
	v473 = v464
	goto L142
L142:
	;
	v476 = int32(*(*int8)(unsafe.Add(mBase, uint32(v448+v473))))
	if int32(-65) < v476 {
		v487 = v473
		goto L139
	} else {
		goto L144
	}
L143:
	;
	v487 = v449
	goto L139
L144:
	;
	v480 = v473 + int32(1)
	if v480 != v449 {
		v473 = v480
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	goto L135
L147:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v502
	v163 = v502
	goto L56
L148:
	;
	goto L149
L149:
	;
	goto L57
L150:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1070
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1070
	if v1070 <= v7 {
		v1257 = v449
		goto L279
	} else {
		goto L280
	}
L151:
	;
	if v565 < int32(0) {
		goto L150
	} else {
		goto L171
	}
L153:
	;
	goto L154
L154:
	;
	goto L155
L155:
	;
	v520 = v512
	v522 = int32(3)
	goto L158
L157:
	;
	v565 = v550
	goto L151
L158:
	;
	if v513 <= v520 {
		goto L160
	} else {
		goto L161
	}
L159:
	;
	goto L157
L160:
	;
	v565 = int32(-1)
	goto L151
L161:
	;
	goto L162
L162:
	;
	v527 = v520 + int32(1)
	v529 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v511+v520))))
	if base.Ui32(v529) < base.Ui32(int32(192)) {
		v550 = v527
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v551 = int32(1)
	if v551 < v522 {
		v520 = v550
		v522 = v522 - v551
		goto L158
	} else {
		goto L170
	}
L164:
	;
	if v513 <= v527 {
		v550 = v527
		goto L163
	} else {
		goto L165
	}
L165:
	;
	v536 = v527
	goto L166
L166:
	;
	v539 = int32(*(*int8)(unsafe.Add(mBase, uint32(v511+v536))))
	if int32(-65) < v539 {
		v550 = v536
		goto L163
	} else {
		goto L168
	}
L167:
	;
	v550 = v513
	goto L163
L168:
	;
	v543 = v536 + int32(1)
	if v543 != v513 {
		v536 = v543
		goto L166
	} else {
		goto L169
	}
L169:
	;
	goto L167
L170:
	;
	goto L159
L171:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v568))) = v565
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v512
	v583 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v592 = v512
	goto L174
L172:
	;
	if v687 < int32(0) {
		goto L150
	} else {
		goto L197
	}
L173:
	;
	v687 = v659
	goto L172
L174:
	;
	if v583 <= v592 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v687 = int32(-1)
	goto L172
L177:
	;
	goto L178
L178:
	;
	v599 = int32(1)
	v601 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v592+v584))))
	if base.Ui32(v601) < base.Ui32(int32(192)) {
		v658 = v601
		v659 = v599
		goto L179
	} else {
		goto L180
	}
L179:
	;
	if int32(232) < v658 {
		goto L192
	} else {
		goto L193
	}
L180:
	;
	v605 = v592 + int32(1)
	if v605 == v583 {
		v658 = v601
		v659 = v599
		goto L179
	} else {
		goto L181
	}
L181:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v605+v584))))
	v610 = v608 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v601) {
		goto L183
	} else {
		goto L184
	}
L182:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v614+v584))))
	v626 = v624 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v601) {
		goto L188
	} else {
		goto L189
	}
L183:
	;
	v614 = v592 + int32(2)
	if v614 != v583 {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v658 = v601<<(uint(int32(6))%32)&int32(1984) | v610
	v659 = int32(2)
	goto L179
L186:
	;
	goto L185
L187:
	;
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584+v630))))
	v658 = v643&int32(63) | (v601<<(uint(int32(18))%32)&int32(1835008) | v610<<(uint(int32(12))%32) | v626<<(uint(int32(6))%32))
	v659 = int32(4)
	goto L179
L188:
	;
	v630 = v592 + int32(3)
	if v630 != v583 {
		goto L187
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v658 = v601<<(uint(int32(12))%32)&int32(61440) | v610<<(uint(int32(6))%32) | v626
	v659 = int32(3)
	goto L179
L191:
	;
	goto L190
L192:
	;
	v676 = v659 + v592
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v676
	v592 = v676
	goto L174
L193:
	;
	v663 = v658 - int32(97)
	if v663 < int32(0) {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v663)>>(uint(int32(3))%32)))+uint32(_consts[1274]))))
	if int32(base.Ui32(v669)>>(uint(v663&int32(7))%32))&int32(1) != 0 {
		goto L173
	} else {
		goto L195
	}
L195:
	;
	goto L192
L197:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v691 = v690 + v687
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v691
	v705 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v706 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v714 = v691
	goto L200
L198:
	;
	if v810 < int32(0) {
		goto L150
	} else {
		goto L222
	}
L199:
	;
	v810 = v781
	goto L198
L200:
	;
	if v705 <= v714 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v810 = int32(-1)
	goto L198
L203:
	;
	goto L204
L204:
	;
	v721 = int32(1)
	v723 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v714+v706))))
	if base.Ui32(v723) < base.Ui32(int32(192)) {
		v780 = v723
		v781 = v721
		goto L205
	} else {
		goto L206
	}
L205:
	;
	if int32(232) < v780 {
		goto L199
	} else {
		goto L218
	}
L206:
	;
	v727 = v714 + int32(1)
	if v727 == v705 {
		v780 = v723
		v781 = v721
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v727+v706))))
	v732 = v730 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v723) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v736+v706))))
	v748 = v746 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v723) {
		goto L214
	} else {
		goto L215
	}
L209:
	;
	v736 = v714 + int32(2)
	if v736 != v705 {
		goto L208
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v780 = v723<<(uint(int32(6))%32)&int32(1984) | v732
	v781 = int32(2)
	goto L205
L212:
	;
	goto L211
L213:
	;
	v765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706+v752))))
	v780 = v765&int32(63) | (v723<<(uint(int32(18))%32)&int32(1835008) | v732<<(uint(int32(12))%32) | v748<<(uint(int32(6))%32))
	v781 = int32(4)
	goto L205
L214:
	;
	v752 = v714 + int32(3)
	if v752 != v705 {
		goto L213
	} else {
		goto L217
	}
L215:
	;
	goto L216
L216:
	;
	v780 = v723<<(uint(int32(12))%32)&int32(61440) | v732<<(uint(int32(6))%32) | v748
	v781 = int32(3)
	goto L205
L217:
	;
	goto L216
L218:
	;
	v785 = v780 - int32(97)
	if v785 < int32(0) {
		goto L199
	} else {
		goto L219
	}
L219:
	;
	v791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v785)>>(uint(int32(3))%32)))+uint32(_consts[1274]))))
	if int32(base.Ui32(v791)>>(uint(v785&int32(7))%32))&int32(1) == int32(0) {
		goto L199
	} else {
		goto L220
	}
L220:
	;
	v799 = v781 + v714
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v799
	v714 = v799
	goto L200
L222:
	;
	v813 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v814 = v813 + v810
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v814
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v817 = *(*int32)(unsafe.Add(mBase, uint32(v816)))
	if v817 < v814 {
		goto L223
	} else {
		goto L224
	}
L223:
	;
	v819 = v814
	goto L225
L224:
	;
	v819 = v817
	goto L225
L225:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v816)+8)) = v819
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v842 = v832
	goto L228
L226:
	;
	if v937 < int32(0) {
		goto L150
	} else {
		goto L251
	}
L227:
	;
	v937 = v909
	goto L226
L228:
	;
	if v833 <= v842 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v937 = int32(-1)
	goto L226
L231:
	;
	goto L232
L232:
	;
	v849 = int32(1)
	v851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v842+v834))))
	if base.Ui32(v851) < base.Ui32(int32(192)) {
		v908 = v851
		v909 = v849
		goto L233
	} else {
		goto L234
	}
L233:
	;
	if int32(232) < v908 {
		goto L246
	} else {
		goto L247
	}
L234:
	;
	v855 = v842 + int32(1)
	if v855 == v833 {
		v908 = v851
		v909 = v849
		goto L233
	} else {
		goto L235
	}
L235:
	;
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v855+v834))))
	v860 = v858 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v851) {
		goto L237
	} else {
		goto L238
	}
L236:
	;
	v874 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v864+v834))))
	v876 = v874 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v851) {
		goto L242
	} else {
		goto L243
	}
L237:
	;
	v864 = v842 + int32(2)
	if v864 != v833 {
		goto L236
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v908 = v851<<(uint(int32(6))%32)&int32(1984) | v860
	v909 = int32(2)
	goto L233
L240:
	;
	goto L239
L241:
	;
	v893 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834+v880))))
	v908 = v893&int32(63) | (v851<<(uint(int32(18))%32)&int32(1835008) | v860<<(uint(int32(12))%32) | v876<<(uint(int32(6))%32))
	v909 = int32(4)
	goto L233
L242:
	;
	v880 = v842 + int32(3)
	if v880 != v833 {
		goto L241
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v908 = v851<<(uint(int32(12))%32)&int32(61440) | v860<<(uint(int32(6))%32) | v876
	v909 = int32(3)
	goto L233
L245:
	;
	goto L244
L246:
	;
	v926 = v909 + v842
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v926
	v842 = v926
	goto L228
L247:
	;
	v913 = v908 - int32(97)
	if v913 < int32(0) {
		goto L246
	} else {
		goto L248
	}
L248:
	;
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v913)>>(uint(int32(3))%32)))+uint32(_consts[1274]))))
	if int32(base.Ui32(v919)>>(uint(v913&int32(7))%32))&int32(1) != 0 {
		goto L227
	} else {
		goto L249
	}
L249:
	;
	goto L246
L251:
	;
	v940 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v941 = v940 + v937
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v941
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v964 = v941
	goto L254
L252:
	;
	if v1060 < int32(0) {
		goto L150
	} else {
		goto L276
	}
L253:
	;
	v1060 = v1031
	goto L252
L254:
	;
	if v955 <= v964 {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v1060 = int32(-1)
	goto L252
L257:
	;
	goto L258
L258:
	;
	v971 = int32(1)
	v973 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v964+v956))))
	if base.Ui32(v973) < base.Ui32(int32(192)) {
		v1030 = v973
		v1031 = v971
		goto L259
	} else {
		goto L260
	}
L259:
	;
	if int32(232) < v1030 {
		goto L253
	} else {
		goto L272
	}
L260:
	;
	v977 = v964 + int32(1)
	if v977 == v955 {
		v1030 = v973
		v1031 = v971
		goto L259
	} else {
		goto L261
	}
L261:
	;
	v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977+v956))))
	v982 = v980 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v973) {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	v996 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v986+v956))))
	v998 = v996 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v973) {
		goto L268
	} else {
		goto L269
	}
L263:
	;
	v986 = v964 + int32(2)
	if v986 != v955 {
		goto L262
	} else {
		goto L266
	}
L264:
	;
	goto L265
L265:
	;
	v1030 = v973<<(uint(int32(6))%32)&int32(1984) | v982
	v1031 = int32(2)
	goto L259
L266:
	;
	goto L265
L267:
	;
	v1015 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v956+v1002))))
	v1030 = v1015&int32(63) | (v973<<(uint(int32(18))%32)&int32(1835008) | v982<<(uint(int32(12))%32) | v998<<(uint(int32(6))%32))
	v1031 = int32(4)
	goto L259
L268:
	;
	v1002 = v964 + int32(3)
	if v1002 != v955 {
		goto L267
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	v1030 = v973<<(uint(int32(12))%32)&int32(61440) | v982<<(uint(int32(6))%32) | v998
	v1031 = int32(3)
	goto L259
L271:
	;
	goto L270
L272:
	;
	v1035 = v1030 - int32(97)
	if v1035 < int32(0) {
		goto L253
	} else {
		goto L273
	}
L273:
	;
	v1041 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1035)>>(uint(int32(3))%32)))+uint32(_consts[1274]))))
	if int32(base.Ui32(v1041)>>(uint(v1035&int32(7))%32))&int32(1) == int32(0) {
		goto L253
	} else {
		goto L274
	}
L274:
	;
	v1049 = v1031 + v964
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1049
	v964 = v1049
	goto L254
L276:
	;
	v1063 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1063)+4)) = v1064 + v1060
	goto L150
L277:
	;
	v1269 = F_r_e_ending_2(m, l0)
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L12
	} else {
		goto L325
	}
L278:
	;
	if v1110 < int32(0) {
		goto L322
	} else {
		goto L323
	}
L279:
	;
	v1258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1258
	v1266 = v1258
	v1267 = v1257
	v1268 = v1258
	goto L277
L280:
	;
	v1074 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1078 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1074+v1070-int32(1)))))
	if v1078&int32(224) != int32(96) {
		v1257 = v449
		goto L279
	} else {
		goto L281
	}
L281:
	;
	if int32(1)<<(uint(v1078)%32)&int32(540704) == int32(0) {
		v1257 = v449
		goto L279
	} else {
		goto L282
	}
L282:
	;
	v1091 = F_find_among_b(m, l0, int32(4212224), int32(5))
	mBase = m.M
	v1092 = m.ExcPending
	if v1092 != 0 {
		goto L12
	} else {
		goto L283
	}
L283:
	;
	if v1091 == int32(0) {
		v1257 = v449
		goto L279
	} else {
		goto L284
	}
L284:
	;
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1095
	switch v1091 - int32(1) {
	case 0:
		goto L287
	case 1:
		goto L286
	case 2:
		goto L285
	default:
		v1257 = v449
		goto L279
	}
L285:
	;
	v1116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(v1116)+8))
	if v1095 < v1117 {
		goto L297
	} else {
		goto L298
	}
L286:
	;
	v1110 = F_r_en_ending_2(m, l0)
	mBase = m.M
	v1111 = m.ExcPending
	if v1111 != 0 {
		goto L12
	} else {
		goto L293
	}
L287:
	;
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1099)+8))
	if v1095 < v1100 {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	v1257 = int32(0)
	goto L279
L289:
	;
	goto L290
L290:
	;
	v1105 = F_slice_from_s(m, l0, int32(4), int32(2148711))
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L12
	} else {
		goto L291
	}
L291:
	;
	if v1105 < int32(0) {
		v1970 = v1105
		goto L1
	} else {
		goto L292
	}
L292:
	;
	v1257 = int32(1)
	goto L279
L293:
	;
	if v1110 != 0 {
		goto L294
	} else {
		goto L295
	}
L294:
	;
	v1115 = int32(base.Ui32(v1110) >> (uint(int32(31)) % 32))
	goto L296
L295:
	;
	v1115 = int32(2)
	goto L296
L296:
	;
	switch v1115 {
	case 0, 2:
		v1257 = v1110
		goto L279
	default:
		goto L278
	}
L297:
	;
	v1257 = int32(0)
	goto L279
L298:
	;
	goto L299
L299:
	;
	v1120 = int32(1)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L302
L300:
	;
	if v1250 != 0 {
		v1257 = v1120
		goto L279
	} else {
		goto L319
	}
L301:
	;
	v1250 = v1243
	goto L300
L302:
	;
	if v1137 <= v1138 {
		v1243 = int32(-1)
		goto L301
	} else {
		goto L304
	}
L303:
	;
	v1243 = int32(0)
	goto L301
L304:
	;
	v1155 = int32(1)
	v1156 = v1137 - v1155
	v1158 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1134+v1156))))
	v1160 = v1158 & int32(255)
	if v1156 == v1138 {
		v1215 = v1160
		v1216 = v1155
		goto L305
	} else {
		goto L306
	}
L305:
	;
	if int32(232) < v1215 {
		goto L314
	} else {
		goto L315
	}
L306:
	;
	if int32(0) <= v1158 {
		v1215 = v1160
		v1216 = v1155
		goto L305
	} else {
		goto L307
	}
L307:
	;
	v1166 = v1160 & int32(63)
	v1168 = v1137 - int32(2)
	v1170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1134+v1168))))
	v1172 = v1170 << (uint(int32(6)) % 32)
	if base.B2i32(v1168 != v1138)&base.B2i32(base.Ui32(v1170) < base.Ui32(int32(192))) == int32(0) {
		goto L308
	} else {
		goto L309
	}
L308:
	;
	v1215 = v1172&int32(1984) | v1166
	v1216 = int32(2)
	goto L305
L309:
	;
	goto L310
L310:
	;
	v1185 = v1172&int32(4032) | v1166
	v1187 = v1137 - int32(3)
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1134+v1187))))
	if base.B2i32(v1187 != v1138)&base.B2i32(base.Ui32(v1189) < base.Ui32(int32(224))) == int32(0) {
		goto L311
	} else {
		goto L312
	}
L311:
	;
	v1215 = v1189<<(uint(int32(12))%32)&int32(61440) | v1185
	v1216 = int32(3)
	goto L305
L312:
	;
	goto L313
L313:
	;
	v1207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1137+(v1134-int32(4))))))
	v1215 = v1189<<(uint(int32(12))%32)&int32(258048) | v1207&int32(7)<<(uint(int32(18))%32) | v1185
	v1216 = int32(4)
	goto L305
L314:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1137 - v1216
	goto L318
L315:
	;
	v1220 = v1215 - int32(97)
	if v1220 < int32(0) {
		goto L314
	} else {
		goto L316
	}
L316:
	;
	v1226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1220)>>(uint(int32(3))%32)))+uint32(_consts[1275]))))
	if int32(base.Ui32(v1226)>>(uint(v1220&int32(7))%32))&int32(1) == int32(0) {
		goto L314
	} else {
		goto L317
	}
L317:
	;
	v1250 = v1216
	goto L300
L318:
	;
	goto L303
L319:
	;
	v1251 = F_slice_del(m, l0)
	mBase = m.M
	v1252 = m.ExcPending
	if v1252 != 0 {
		goto L12
	} else {
		goto L320
	}
L320:
	;
	if v1251 < int32(0) {
		v1970 = v1251
		goto L1
	} else {
		goto L321
	}
L321:
	;
	v1257 = v1120
	goto L279
L322:
	;
	return v1110
L323:
	;
	goto L324
L324:
	;
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1266 = v1263
	v1267 = v1110
	v1268 = v1264
	goto L277
L325:
	;
	if v1269 < int32(0) {
		v1970 = v1269
		goto L1
	} else {
		goto L326
	}
L326:
	;
	v1273 = v1266 - v1268
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1275 = v1273 + v1274
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1275
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1275
	v1278 = int32(4)
	v1280 = int32(0)
	v1283 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1275-v1283 < v1278 {
		v1293 = v1280
		goto L334
	} else {
		goto L335
	}
L327:
	;
	v1868 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1868
	v1871 = v1868
	goto L488
L328:
	;
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1529 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L422
L329:
	;
	if v1505 < int32(0) {
		v1970 = v1505
		goto L1
	} else {
		goto L419
	}
L330:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1355
	v1360 = v1355 - int32(1)
	v1361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1360 <= v1361 {
		goto L362
	} else {
		goto L363
	}
L331:
	;
	if v1336 < int32(0) {
		v1505 = v1341
		goto L329
	} else {
		goto L360
	}
L332:
	;
	v1349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1350 = v1349 + v1273
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1350
	v1355 = v1350
	v1356 = v1347
	v1357 = v1349
	goto L330
L333:
	;
	if v1293 == int32(0) {
		v1347 = v1267
		goto L332
	} else {
		goto L337
	}
L334:
	;
	goto L333
L335:
	;
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1289 = F_memcmp(m, v1286+v1275-v1278, int32(2148737), v1278)
	mBase = m.M
	if v1289 != 0 {
		v1293 = v1280
		goto L334
	} else {
		goto L336
	}
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1275 - v1278
	v1293 = int32(1)
	goto L334
L337:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1296
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1299 = *(*int32)(unsafe.Add(mBase, uint32(v1298)+4))
	if v1296 < v1299 {
		v1347 = v1267
		goto L332
	} else {
		goto L338
	}
L338:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1301 < v1296 {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1307 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1303+v1296-int32(1)))))
	if v1307 == int32(99) {
		v1347 = v1267
		goto L332
	} else {
		goto L342
	}
L340:
	;
	goto L341
L341:
	;
	v1310 = F_slice_del(m, l0)
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L12
	} else {
		goto L343
	}
L342:
	;
	goto L341
L343:
	;
	if v1310 < int32(0) {
		v1970 = v1310
		goto L1
	} else {
		goto L344
	}
L344:
	;
	v1314 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1314
	v1316 = int32(2)
	v1318 = int32(0)
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1314-v1321 < v1316 {
		v1331 = v1318
		goto L346
	} else {
		goto L347
	}
L345:
	;
	if v1331 == int32(0) {
		v1347 = v1267
		goto L332
	} else {
		goto L349
	}
L346:
	;
	goto L345
L347:
	;
	v1324 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1327 = F_memcmp(m, v1324+v1314-v1316, int32(2148741), v1316)
	mBase = m.M
	if v1327 != 0 {
		v1331 = v1318
		goto L346
	} else {
		goto L348
	}
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1314 - v1316
	v1331 = int32(1)
	goto L346
L349:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1334
	v1336 = F_r_en_ending_2(m, l0)
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L12
	} else {
		goto L350
	}
L350:
	;
	v1339 = base.B2i32(v1336 < int32(0))
	if v1336 < int32(0) {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1340 = v1336
	goto L353
L352:
	;
	v1340 = v1267
	goto L353
L353:
	;
	if v1336 != 0 {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1341 = v1340
	goto L356
L355:
	;
	v1341 = v1267
	goto L356
L356:
	;
	if v1336 != 0 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1345 = int32(base.Ui32(v1336) >> (uint(int32(31)) % 32))
	goto L359
L358:
	;
	v1345 = int32(4)
	goto L359
L359:
	;
	switch v1345 {
	case 0, 4:
		v1347 = v1341
		goto L332
	default:
		goto L331
	}
L360:
	;
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1355 = v1352
	v1356 = v1341
	v1357 = v1353
	goto L330
L361:
	;
	v1494 = int32(0)
	v1495 = base.B2i32(v1464 < v1494)
	if v1495 == v1494 {
		goto L328
	} else {
		goto L412
	}
L362:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1490 + (v1355 - v1357)
	goto L328
L363:
	;
	v1363 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1363+v1360))))
	if v1365&int32(224) != int32(96) {
		goto L362
	} else {
		goto L364
	}
L364:
	;
	if int32(1)<<(uint(v1365)%32)&int32(264336) == int32(0) {
		goto L362
	} else {
		goto L365
	}
L365:
	;
	v1378 = F_find_among_b(m, l0, int32(4212336), int32(6))
	mBase = m.M
	v1379 = m.ExcPending
	if v1379 != 0 {
		goto L12
	} else {
		goto L366
	}
L366:
	;
	if v1378 == int32(0) {
		goto L362
	} else {
		goto L367
	}
L367:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1382
	switch v1378 - int32(1) {
	case 0:
		goto L372
	case 1:
		goto L371
	case 2:
		goto L370
	case 3:
		goto L369
	case 4:
		goto L368
	default:
		goto L362
	}
L368:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+4))
	if v1382 < v1478 {
		goto L362
	} else {
		goto L408
	}
L369:
	;
	v1470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v1470)+4))
	if v1382 < v1471 {
		goto L362
	} else {
		goto L405
	}
L370:
	;
	v1457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1458 = *(*int32)(unsafe.Add(mBase, uint32(v1457)+4))
	if v1382 < v1458 {
		goto L362
	} else {
		goto L398
	}
L371:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1441)+4))
	if v1382 < v1442 {
		goto L362
	} else {
		goto L391
	}
L372:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1386)+4))
	if v1382 < v1387 {
		goto L362
	} else {
		goto L373
	}
L373:
	;
	v1389 = F_slice_del(m, l0)
	mBase = m.M
	v1390 = m.ExcPending
	if v1390 != 0 {
		goto L12
	} else {
		goto L374
	}
L374:
	;
	if v1389 < int32(0) {
		v1970 = v1389
		goto L1
	} else {
		goto L375
	}
L375:
	;
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1393
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1396 = int32(2)
	v1398 = int32(0)
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1393-v1401 < v1396 {
		v1411 = v1398
		goto L378
	} else {
		goto L379
	}
L376:
	;
	v1433 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1433 + (v1393 - v1395)
	v1437 = F_r_undouble(m, l0)
	mBase = m.M
	v1438 = m.ExcPending
	if v1438 != 0 {
		goto L12
	} else {
		goto L389
	}
L377:
	;
	if v1411 == int32(0) {
		goto L376
	} else {
		goto L381
	}
L378:
	;
	goto L377
L379:
	;
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1407 = F_memcmp(m, v1404+v1393-v1396, int32(2148743), v1396)
	mBase = m.M
	if v1407 != 0 {
		v1411 = v1398
		goto L378
	} else {
		goto L380
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1393 - v1396
	v1411 = int32(1)
	goto L378
L381:
	;
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1414
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1416)+4))
	if v1414 < v1417 {
		goto L376
	} else {
		goto L382
	}
L382:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1419 < v1414 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1421+v1414-int32(1)))))
	if v1425 == int32(101) {
		goto L376
	} else {
		goto L386
	}
L384:
	;
	goto L385
L385:
	;
	v1428 = F_slice_del(m, l0)
	mBase = m.M
	v1429 = m.ExcPending
	if v1429 != 0 {
		goto L12
	} else {
		goto L387
	}
L386:
	;
	goto L385
L387:
	;
	if int32(0) <= v1428 {
		goto L362
	} else {
		goto L388
	}
L388:
	;
	v1970 = v1428
	goto L1
L389:
	;
	if int32(0) <= v1437 {
		goto L362
	} else {
		goto L390
	}
L390:
	;
	v1970 = v1437
	goto L1
L391:
	;
	v1444 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1444 < v1382 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v1446 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446+v1382-int32(1)))))
	if v1450 == int32(101) {
		goto L362
	} else {
		goto L395
	}
L393:
	;
	goto L394
L394:
	;
	v1453 = F_slice_del(m, l0)
	mBase = m.M
	v1454 = m.ExcPending
	if v1454 != 0 {
		goto L12
	} else {
		goto L396
	}
L395:
	;
	goto L394
L396:
	;
	if int32(0) <= v1453 {
		goto L362
	} else {
		goto L397
	}
L397:
	;
	v1970 = v1453
	goto L1
L398:
	;
	v1460 = F_slice_del(m, l0)
	mBase = m.M
	v1461 = m.ExcPending
	if v1461 != 0 {
		goto L12
	} else {
		goto L399
	}
L399:
	;
	if v1460 < int32(0) {
		v1970 = v1460
		goto L1
	} else {
		goto L400
	}
L400:
	;
	v1464 = F_r_e_ending_2(m, l0)
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L12
	} else {
		goto L401
	}
L401:
	;
	if v1464 != 0 {
		goto L402
	} else {
		goto L403
	}
L402:
	;
	v1469 = int32(base.Ui32(v1464) >> (uint(int32(31)) % 32))
	goto L404
L403:
	;
	v1469 = int32(6)
	goto L404
L404:
	;
	switch v1469 {
	case 0, 6:
		goto L362
	default:
		goto L361
	}
L405:
	;
	v1473 = F_slice_del(m, l0)
	mBase = m.M
	v1474 = m.ExcPending
	if v1474 != 0 {
		goto L12
	} else {
		goto L406
	}
L406:
	;
	if int32(0) <= v1473 {
		goto L362
	} else {
		goto L407
	}
L407:
	;
	v1970 = v1473
	goto L1
L408:
	;
	v1480 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+12))
	if v1480 == int32(0) {
		goto L362
	} else {
		goto L409
	}
L409:
	;
	v1483 = F_slice_del(m, l0)
	mBase = m.M
	v1484 = m.ExcPending
	if v1484 != 0 {
		goto L12
	} else {
		goto L410
	}
L410:
	;
	if v1483 < int32(0) {
		v1970 = v1483
		goto L1
	} else {
		goto L411
	}
L411:
	;
	goto L362
L412:
	;
	if v1464 < v1494 {
		goto L413
	} else {
		goto L414
	}
L413:
	;
	v1498 = v1464
	goto L415
L414:
	;
	v1498 = v1356
	goto L415
L415:
	;
	if v1464 != 0 {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1499 = v1498
	goto L418
L417:
	;
	v1499 = v1356
	goto L418
L418:
	;
	v1505 = v1499
	goto L329
L419:
	;
	goto L327
L420:
	;
	if v1641 != 0 {
		goto L327
	} else {
		goto L439
	}
L421:
	;
	v1641 = v1634
	goto L420
L422:
	;
	if v1528 <= v1529 {
		v1634 = int32(-1)
		goto L421
	} else {
		goto L424
	}
L423:
	;
	v1634 = int32(0)
	goto L421
L424:
	;
	v1546 = int32(1)
	v1547 = v1528 - v1546
	v1549 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1525+v1547))))
	v1551 = v1549 & int32(255)
	if v1547 == v1529 {
		v1606 = v1551
		v1607 = v1546
		goto L425
	} else {
		goto L426
	}
L425:
	;
	if int32(232) < v1606 {
		goto L434
	} else {
		goto L435
	}
L426:
	;
	if int32(0) <= v1549 {
		v1606 = v1551
		v1607 = v1546
		goto L425
	} else {
		goto L427
	}
L427:
	;
	v1557 = v1551 & int32(63)
	v1559 = v1528 - int32(2)
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1525+v1559))))
	v1563 = v1561 << (uint(int32(6)) % 32)
	if base.B2i32(v1559 != v1529)&base.B2i32(base.Ui32(v1561) < base.Ui32(int32(192))) == int32(0) {
		goto L428
	} else {
		goto L429
	}
L428:
	;
	v1606 = v1563&int32(1984) | v1557
	v1607 = int32(2)
	goto L425
L429:
	;
	goto L430
L430:
	;
	v1576 = v1563&int32(4032) | v1557
	v1578 = v1528 - int32(3)
	v1580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1525+v1578))))
	if base.B2i32(v1578 != v1529)&base.B2i32(base.Ui32(v1580) < base.Ui32(int32(224))) == int32(0) {
		goto L431
	} else {
		goto L432
	}
L431:
	;
	v1606 = v1580<<(uint(int32(12))%32)&int32(61440) | v1576
	v1607 = int32(3)
	goto L425
L432:
	;
	goto L433
L433:
	;
	v1598 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1528+(v1525-int32(4))))))
	v1606 = v1580<<(uint(int32(12))%32)&int32(258048) | v1598&int32(7)<<(uint(int32(18))%32) | v1576
	v1607 = int32(4)
	goto L425
L434:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1528 - v1607
	goto L438
L435:
	;
	v1611 = v1606 - int32(73)
	if v1611 < int32(0) {
		goto L434
	} else {
		goto L436
	}
L436:
	;
	v1617 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1611)>>(uint(int32(3))%32)))+uint32(_consts[1276]))))
	if int32(base.Ui32(v1617)>>(uint(v1611&int32(7))%32))&int32(1) == int32(0) {
		goto L434
	} else {
		goto L437
	}
L437:
	;
	v1641 = v1607
	goto L420
L438:
	;
	goto L423
L439:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1644 = v1642 - int32(1)
	v1645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1644 <= v1645 {
		goto L327
	} else {
		goto L440
	}
L440:
	;
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1647+v1644))))
	if v1649&int32(224) != int32(96) {
		goto L327
	} else {
		goto L441
	}
L441:
	;
	if int32(1)<<(uint(v1649)%32)&int32(2129954) == int32(0) {
		goto L327
	} else {
		goto L442
	}
L442:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1663 = F_find_among_b(m, l0, int32(4212464), int32(4))
	mBase = m.M
	v1664 = m.ExcPending
	if v1664 != 0 {
		goto L12
	} else {
		goto L443
	}
L443:
	;
	if v1663 == int32(0) {
		goto L327
	} else {
		goto L444
	}
L444:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1683 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L447
L445:
	;
	if v1796 != 0 {
		goto L327
	} else {
		goto L464
	}
L446:
	;
	v1796 = v1789
	goto L445
L447:
	;
	if v1683 <= v1684 {
		v1789 = int32(-1)
		goto L446
	} else {
		goto L449
	}
L448:
	;
	v1789 = int32(0)
	goto L446
L449:
	;
	v1701 = int32(1)
	v1702 = v1683 - v1701
	v1704 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1680+v1702))))
	v1706 = v1704 & int32(255)
	if v1702 == v1684 {
		v1761 = v1706
		v1762 = v1701
		goto L450
	} else {
		goto L451
	}
L450:
	;
	if int32(232) < v1761 {
		goto L459
	} else {
		goto L460
	}
L451:
	;
	if int32(0) <= v1704 {
		v1761 = v1706
		v1762 = v1701
		goto L450
	} else {
		goto L452
	}
L452:
	;
	v1712 = v1706 & int32(63)
	v1714 = v1683 - int32(2)
	v1716 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1680+v1714))))
	v1718 = v1716 << (uint(int32(6)) % 32)
	if base.B2i32(v1714 != v1684)&base.B2i32(base.Ui32(v1716) < base.Ui32(int32(192))) == int32(0) {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	v1761 = v1718&int32(1984) | v1712
	v1762 = int32(2)
	goto L450
L454:
	;
	goto L455
L455:
	;
	v1731 = v1718&int32(4032) | v1712
	v1733 = v1683 - int32(3)
	v1735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1680+v1733))))
	if base.B2i32(v1733 != v1684)&base.B2i32(base.Ui32(v1735) < base.Ui32(int32(224))) == int32(0) {
		goto L456
	} else {
		goto L457
	}
L456:
	;
	v1761 = v1735<<(uint(int32(12))%32)&int32(61440) | v1731
	v1762 = int32(3)
	goto L450
L457:
	;
	goto L458
L458:
	;
	v1753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1683+(v1680-int32(4))))))
	v1761 = v1735<<(uint(int32(12))%32)&int32(258048) | v1753&int32(7)<<(uint(int32(18))%32) | v1731
	v1762 = int32(4)
	goto L450
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1683 - v1762
	goto L463
L460:
	;
	v1766 = v1761 - int32(97)
	if v1766 < int32(0) {
		goto L459
	} else {
		goto L461
	}
L461:
	;
	v1772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1766)>>(uint(int32(3))%32)))+uint32(_consts[1274]))))
	if int32(base.Ui32(v1772)>>(uint(v1766&int32(7))%32))&int32(1) == int32(0) {
		goto L459
	} else {
		goto L462
	}
L462:
	;
	v1796 = v1762
	goto L445
L463:
	;
	goto L448
L464:
	;
	v1797 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1799 = v1797 + (v1642 - v1660)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1799
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1799
	v1802 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L467
L465:
	;
	if v1854 < int32(0) {
		goto L327
	} else {
		goto L485
	}
L467:
	;
	goto L468
L468:
	;
	goto L469
L469:
	;
	v1810 = v1799
	v1812 = int32(1)
	goto L472
L471:
	;
	v1854 = v1836
	goto L465
L472:
	;
	if v1810 <= v1803 {
		goto L474
	} else {
		goto L475
	}
L473:
	;
	goto L471
L474:
	;
	v1854 = int32(-1)
	goto L465
L475:
	;
	goto L476
L476:
	;
	v1817 = v1810 - int32(1)
	v1819 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1802+v1817))))
	if int32(0) <= v1819 {
		v1836 = v1817
		goto L477
	} else {
		goto L478
	}
L477:
	;
	v1840 = int32(1)
	if v1840 < v1812 {
		v1810 = v1836
		v1812 = v1812 - v1840
		goto L472
	} else {
		goto L484
	}
L478:
	;
	if v1817 <= v1803 {
		v1836 = v1817
		goto L477
	} else {
		goto L479
	}
L479:
	;
	v1824 = v1817
	goto L480
L480:
	;
	v1829 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1802+v1824))))
	if base.Ui32(int32(191)) < base.Ui32(v1829) {
		v1836 = v1824
		goto L477
	} else {
		goto L482
	}
L481:
	;
	v1836 = v1803
	goto L477
L482:
	;
	v1833 = v1824 - int32(1)
	if v1803 < v1833 {
		v1824 = v1833
		goto L480
	} else {
		goto L483
	}
L483:
	;
	goto L481
L484:
	;
	goto L473
L485:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1854
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1854
	v1859 = F_slice_del(m, l0)
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L12
	} else {
		goto L486
	}
L486:
	;
	if v1859 < int32(0) {
		v1970 = v1859
		goto L1
	} else {
		goto L487
	}
L487:
	;
	goto L327
L488:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1871
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v1877 <= v1871 {
		goto L493
	} else {
		goto L494
	}
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1868
	v1970 = int32(1)
	goto L1
L490:
	;
	goto L489
L491:
	;
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1871 = v1966
	goto L488
L492:
	;
	v1908 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L506
L493:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1871
	v1906 = v1871
	v1907 = v1877
	goto L492
L494:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1879+v1871))))
	switch v1881 - int32(73) {
	case 0, 16:
		goto L495
	default:
		goto L493
	}
L495:
	;
	v1886 = F_find_among(m, l0, int32(4212608), int32(3))
	mBase = m.M
	v1887 = m.ExcPending
	if v1887 != 0 {
		goto L12
	} else {
		goto L496
	}
L496:
	;
	v1888 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1888
	switch v1886 - int32(1) {
	case 0:
		goto L498
	case 1:
		goto L497
	case 2:
		goto L499
	default:
		goto L491
	}
L497:
	;
	v1901 = F_slice_from_s(m, l0, int32(1), int32(2148822))
	mBase = m.M
	v1902 = m.ExcPending
	if v1902 != 0 {
		goto L12
	} else {
		goto L502
	}
L498:
	;
	v1895 = F_slice_from_s(m, l0, int32(1), int32(2148821))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L12
	} else {
		goto L500
	}
L499:
	;
	v1892 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1906 = v1888
	v1907 = v1892
	goto L492
L500:
	;
	if int32(0) <= v1895 {
		goto L491
	} else {
		goto L501
	}
L501:
	;
	v1970 = v1895
	goto L1
L502:
	;
	if int32(0) <= v1901 {
		goto L491
	} else {
		goto L503
	}
L503:
	;
	v1970 = v1901
	goto L1
L504:
	;
	if v1960 < int32(0) {
		goto L490
	} else {
		goto L524
	}
L506:
	;
	goto L507
L507:
	;
	goto L508
L508:
	;
	v1915 = v1906
	v1917 = int32(1)
	goto L511
L510:
	;
	v1960 = v1945
	goto L504
L511:
	;
	if v1907 <= v1915 {
		goto L513
	} else {
		goto L514
	}
L512:
	;
	goto L510
L513:
	;
	v1960 = int32(-1)
	goto L504
L514:
	;
	goto L515
L515:
	;
	v1922 = v1915 + int32(1)
	v1924 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1908+v1915))))
	if base.Ui32(v1924) < base.Ui32(int32(192)) {
		v1945 = v1922
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v1946 = int32(1)
	if v1946 < v1917 {
		v1915 = v1945
		v1917 = v1917 - v1946
		goto L511
	} else {
		goto L523
	}
L517:
	;
	if v1907 <= v1922 {
		v1945 = v1922
		goto L516
	} else {
		goto L518
	}
L518:
	;
	v1931 = v1922
	goto L519
L519:
	;
	v1934 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1908+v1931))))
	if int32(-65) < v1934 {
		v1945 = v1931
		goto L516
	} else {
		goto L521
	}
L520:
	;
	v1945 = v1907
	goto L516
L521:
	;
	v1938 = v1931 + int32(1)
	if v1938 != v1907 {
		v1931 = v1938
		goto L519
	} else {
		goto L522
	}
L522:
	;
	goto L520
L523:
	;
	goto L512
L524:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1960
	goto L491
}
