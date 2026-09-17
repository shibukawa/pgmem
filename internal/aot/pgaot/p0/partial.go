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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v33 float64
	_ = v33
	var v34 int32
	_ = v34
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v47 float64
	_ = v47
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v54 float64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v62 int32
	_ = v62
	var v68 float64
	_ = v68
	var v72 float64
	_ = v72
	var v75 float64
	_ = v75
	var v77 float64
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v94 float64
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v105 float64
	_ = v105
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v121 float64
	_ = v121
	var v127 float64
	_ = v127
	var v129 float64
	_ = v129
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 float64
	_ = v149
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v184 float64
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 float64
	_ = v243
	var v250 float64
	_ = v250
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v282 int32
	_ = v282
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
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
	v26 = m.G0
	v28 = v26 - int32(16)
	m.G0 = v28
	v30 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v31 = *(*float64)(unsafe.Add(mBase, uint32(l2)+32))
	v33 = *(*float64)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[0]))
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
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v149 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
	v150 = int32(0)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v157 == v150 {
		goto L24
	} else {
		goto L25
	}
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v37 = base.F64_convert_i32_s(v34)
	goto L9
L8:
	;
	v37 = float64(0)
	goto L9
L9:
	;
	v38 = base.F64_mul(v33, v37)
	v41 = *(*float64)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[1]))
	v44 = *(*float64)(unsafe.Add(mBase, uint32(l2)+56))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(l2)+48))
	v47 = float64(0)
	v49 = base.F64_add(base.F64_mul(v38, v31), base.F64_add(base.F64_sub(v44, v45), v47))
	v52 = *(*float64)(unsafe.Add(mBase, uint32(l3)+56))
	v54 = base.F64_add(base.F64_mul(base.F64_add(v38, v41), v30), base.F64_add(base.F64_add(v45, v47), v52))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[2])))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l3)+40))
	if l7 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v60 = base.F64_convert_i32_s(v59)
	v62 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[3])))
	if v62 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v77 = v30
	goto L12
L12:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v79)+32))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_ExecChooseHashTableSize(m, v77, v80, int32(1), l7, v82, v28, v28+int32(12), v28+int32(8), v28+int32(4))
	mBase = m.M
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	if int32(2) <= v90 {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v68 = base.F64_add(base.F64_mul(v60, float64(-0.3)), float64(1))
	if base.F64_gt(v68, float64(0)) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v75 = v60
	goto L15
L15:
	;
	v77 = base.F64_mul(v30, v75)
	goto L12
L16:
	;
	v72 = v68
	goto L18
L17:
	;
	v72 = math.Float64frombits(uint64(0x8000000000000000))
	goto L18
L18:
	;
	v75 = base.F64_add(v72, v60)
	goto L15
L19:
	;
	v94 = *(*float64)(unsafe.Add(mBase, _c_F_try_partial_hashjoin_path[4]))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
	v97 = int32(7)
	v99 = int32(-8)
	v101 = int32(24)
	v105 = float64(0.0001220703125)
	v107 = base.F64_ceil(base.F64_mul(base.F64_mul(v31, base.F64_convert_i32_u((v96+v97)&v99+v101)), v105))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+32))
	v121 = base.F64_ceil(base.F64_mul(base.F64_mul(v30, base.F64_convert_i32_u((v110+v97)&v99+v101)), v105))
	v127 = base.F64_add(base.F64_mul(v94, v121), v54)
	v129 = base.F64_add(base.F64_mul(v94, base.F64_add(base.F64_add(v107, v107), v121)), v49)
	goto L21
L20:
	;
	v127 = v54
	v129 = v49
	goto L21
L21:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v13)+24)) = v129
	*(*float64)(unsafe.Add(mBase, uint32(v13)+8)) = v127
	*(*float64)(unsafe.Add(mBase, uint32(v13)+16)) = base.F64_add(v129, v127)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v58 + (v56 ^ int32(1)) + v57
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	*(*float64)(unsafe.Add(mBase, uint32(v13)+88)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v13)+84)) = v90
	*(*int32)(unsafe.Add(mBase, uint32(v13)+80)) = v141
	m.G0 = v28 + int32(16)
	goto L6
L22:
	;
	if v297 == int32(0) {
		goto L1
	} else {
		goto L63
	}
L23:
	;
	v297 = v282
	goto L22
L24:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v214 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v160 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v171 = v150
	goto L27
L27:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v157)+12))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v171<<(uint(int32(2))%32))))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+64))
	v181 = F_compare_pathkeys(m, v150, v180)
	mBase = m.M
	if v181 == int32(3) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L24
L29:
	;
	v201 = v171 + int32(1)
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	if v201 < v202 {
		v171 = v201
		goto L27
	} else {
		goto L34
	}
L30:
	;
	v184 = *(*float64)(unsafe.Add(mBase, uint32(v179)+56))
	v188 = int32(0)
	v192 = base.B2i32(base.F64_gt(v149, base.F64_mul(v184, float64(1.01))) == v188) | base.B2i32(v181 == int32(1))
	if v192 == v188 {
		v282 = v192
		goto L23
	} else {
		goto L31
	}
L31:
	;
	if v181 == int32(2) {
		goto L29
	} else {
		goto L32
	}
L32:
	;
	if base.F64_lt(base.F64_mul(v149, float64(1.01)), v184) != 0 {
		v282 = v192
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
	v297 = int32(1)
	goto L22
L36:
	;
	goto L37
L37:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v218 <= int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v297 = int32(1)
	goto L22
L39:
	;
	goto L40
L40:
	;
	v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
	v230 = int32(0)
	goto L41
L41:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v234+v230<<(uint(int32(2))%32))))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+40))
	if v148 != v239 {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	v282 = v272
	goto L23
L43:
	;
	if v223 != 0 {
		goto L50
	} else {
		goto L51
	}
L44:
	;
	if v239 <= v148 {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v243 = *(*float64)(unsafe.Add(mBase, uint32(v238)+56))
	if base.F64_le(v149, base.F64_mul(v243, float64(1.01))) == int32(0) {
		goto L43
	} else {
		goto L48
	}
L47:
	;
	v297 = int32(1)
	goto L22
L48:
	;
	v297 = int32(1)
	goto L22
L49:
	;
	v272 = int32(1)
	v274 = v230 + v272
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v274 < v275 {
		v230 = v274
		goto L41
	} else {
		goto L62
	}
L50:
	;
	v250 = *(*float64)(unsafe.Add(mBase, uint32(v238)+48))
	if base.F64_gt(v149, base.F64_mul(v250, float64(1.01))) == int32(0) {
		goto L49
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	if v257 != 0 {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	goto L52
L54:
	;
	v260 = int32(0)
	goto L56
L55:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v238)+64))
	v260 = v259
	goto L56
L56:
	;
	v261 = F_compare_pathkeys(m, v150, v260)
	mBase = m.M
	if v261&int32(-3) != 0 {
		goto L49
	} else {
		goto L57
	}
L57:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	if v265 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+4))
	v268 = v266
	goto L60
L59:
	;
	v268 = int32(0)
	goto L60
L60:
	;
	v269 = F_bms_equal(m, int32(0), v268)
	mBase = m.M
	if v269 != 0 {
		v282 = int32(0)
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
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v302 = F_create_hashjoin_path(m, l0, l1, l5, v13, l6, l2, l3, l7, v300, int32(0), l4)
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	return
L65:
	;
	F_add_partial_path(m, l1, v302)
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	goto L1
}
