package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_try_partial_hashjoin_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v35 float64
	_ = v35
	var v36 int32
	_ = v36
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v45 float64
	_ = v45
	var v48 float64
	_ = v48
	var v53 float64
	_ = v53
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	var v58 int32
	_ = v58
	var v64 float64
	_ = v64
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v96 float64
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 float64
	_ = v107
	var v109 float64
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v123 float64
	_ = v123
	var v129 float64
	_ = v129
	var v130 float64
	_ = v130
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 float64
	_ = v153
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 float64
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 float64
	_ = v247
	var v254 float64
	_ = v254
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	v11 = m.G0
	v13 = v11 - int32(96)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v15 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(96)
	return
L2:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v16 != 0 {
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v28 = m.G0
	v30 = v28 - int32(16)
	m.G0 = v30
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v35 = *(*float64)(unsafe.Add(mBase, _consts[384]))
	if l4 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L4
L6:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v153 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
	v154 = int32(0)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v161 == v154 {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v39 = base.F64_convert_i32_s(v36)
	goto L9
L8:
	;
	v39 = float64(0)
	goto L9
L9:
	;
	v40 = base.F64_mul(v35, v39)
	v42 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v43 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v45 = float64(0)
	v48 = *(*float64)(unsafe.Add(mBase, _consts[380]))
	v53 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	if l7 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v56 = base.F64_convert_i32_s(v55)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _consts[381])))
	if v58 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v74 = v32
	goto L12
L12:
	;
	v75 = base.F64_add(base.F64_mul(v40, v33), base.F64_add(base.F64_sub(v42, v43), v45))
	v76 = base.F64_add(base.F64_mul(base.F64_add(v40, v48), v32), base.F64_add(base.F64_add(v43, v45), v53))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _consts[391])))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+32))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_ExecChooseHashTableSize(m, v74, v82, int32(1), l7, v84, v30, v30+int32(12), v30+int32(8), v30+int32(4))
	mBase = m.M
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if int32(2) <= v92 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v64 = base.F64_add(base.F64_mul(v56, float64(-0.3)), float64(1))
	if base.F64_gt(v64, float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v71 = v56
	goto L15
L15:
	;
	v74 = base.F64_mul(v32, v71)
	goto L12
L16:
	;
	v68 = v64
	goto L18
L17:
	;
	v68 = math.Float64frombits(uint64(0x8000000000000000))
	goto L18
L18:
	;
	v71 = base.F64_add(v68, v56)
	goto L15
L19:
	;
	v96 = *(*float64)(unsafe.Add(mBase, _consts[386]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+32))
	v99 = int32(7)
	v101 = int32(-8)
	v103 = int32(24)
	v107 = float64(0.0001220703125)
	v109 = base.F64_ceil(base.F64_mul(base.F64_mul(v33, base.F64_convert_i32_u((v98+v99)&v101+v103)), v107))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+32))
	v123 = base.F64_ceil(base.F64_mul(base.F64_mul(v32, base.F64_convert_i32_u((v112+v99)&v101+v103)), v107))
	v129 = base.F64_add(base.F64_mul(v96, base.F64_add(base.F64_add(v109, v109), v123)), v75)
	v130 = base.F64_add(base.F64_mul(v96, v123), v76)
	goto L21
L20:
	;
	v129 = v75
	v130 = v76
	goto L21
L21:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v129
	*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v130
	*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = base.F64_add(v129, v130)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v80 + (v78^int32(1))&int32(255) + v79
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v30)+12))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+88)) = v74
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v145
	m.G0 = v30 + int32(16)
	goto L6
L22:
	;
	if v299 == int32(0) {
		goto L1
	} else {
		goto L63
	}
L23:
	;
	v299 = v284
	goto L22
L24:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v218 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v164 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v173 = v154
	goto L27
L27:
	;
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179+v173<<(uint(int32(2))%32))))
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+64))
	v185 = F_compare_pathkeys(m, v154, v184)
	mBase = m.M
	if v185 == int32(3) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L24
L29:
	;
	v205 = v173 + int32(1)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	if v205 < v206 {
		v173 = v205
		goto L27
	} else {
		goto L34
	}
L30:
	;
	v188 = *(*float64)(unsafe.Add(mBase, uint32(v183)+56))
	v192 = int32(0)
	v196 = base.B2i32(base.F64_gt(v153, base.F64_mul(v188, float64(1.01))) == v192) | base.B2i32(v185 == int32(1))
	if v196 == v192 {
		v284 = v196
		goto L23
	} else {
		goto L31
	}
L31:
	;
	if v185 == int32(2) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if base.F64_lt(base.F64_mul(v153, float64(1.01)), v188) != 0 {
		v284 = v196
		goto L23
	} else {
		goto L33
	}
L33:
	;
	goto L29
L34:
	;
	goto L28
L35:
	;
	v299 = int32(1)
	goto L22
L36:
	;
	goto L37
L37:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v222 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v299 = int32(1)
	goto L22
L39:
	;
	goto L40
L40:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	v234 = int32(0)
	goto L41
L41:
	;
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v218)+12))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238+v234<<(uint(int32(2))%32))))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+40))
	if v152 != v243 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v284 = v274
	goto L23
L43:
	;
	if v227 != 0 {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v243 <= v152 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v247 = *(*float64)(unsafe.Add(mBase, uint32(v242)+56))
	if base.F64_le(v153, base.F64_mul(v247, float64(1.01))) == int32(0) {
		goto L43
	} else {
		goto L48
	}
L47:
	;
	v299 = int32(1)
	goto L22
L48:
	;
	v299 = int32(1)
	goto L22
L49:
	;
	v274 = int32(1)
	v276 = v234 + v274
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	if v276 < v277 {
		v234 = v276
		goto L41
	} else {
		goto L62
	}
L50:
	;
	v254 = *(*float64)(unsafe.Add(mBase, uint32(v242)+48))
	if base.F64_gt(v153, base.F64_mul(v254, float64(1.01))) == int32(0) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v260 = int32(0)
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	if v261 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v263 = v260
	goto L56
L55:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v242)+64))
	v263 = v262
	goto L56
L56:
	;
	v264 = F_compare_pathkeys(m, v154, v263)
	mBase = m.M
	if v264&int32(-3) != 0 {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v242)+16))
	if v268 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	v270 = v269
	goto L60
L59:
	;
	v270 = v260
	goto L60
L60:
	;
	v271 = F_bms_equal(m, int32(0), v270)
	mBase = m.M
	if v271 != 0 {
		v284 = v260
		goto L23
	} else {
		goto L61
	}
L61:
	;
	goto L49
L62:
	;
	goto L42
L63:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v304 = F_create_hashjoin_path(m, l0, l1, l5, v13, l6, l2, l3, l7, v302, int32(0), l4)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	return
L65:
	;
	F_add_partial_path(m, l1, v304)
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	goto L1
}
