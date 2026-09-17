package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_analyze_mcv_list(m *base.Module, l0 int32, l1 int32, l2 float64, l3 float64, l4 int32, l5 float64) int32 {
	mBase := m.M
	_ = mBase
	var v7 float64
	_ = v7
	var v13 int32
	_ = v13
	var v19 float64
	_ = v19
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v45 int32
	_ = v45
	var v48 float64
	_ = v48
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 float64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v85 int32
	_ = v85
	var v88 float64
	_ = v88
	var v103 int32
	_ = v103
	var v106 float64
	_ = v106
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v135 float64
	_ = v135
	var v146 float64
	_ = v146
	var v153 int32
	_ = v153
	var v159 float64
	_ = v159
	var v170 float64
	_ = v170
	var v171 float64
	_ = v171
	var v175 float64
	_ = v175
	var v178 float64
	_ = v178
	var v181 float64
	_ = v181
	var v183 int32
	_ = v183
	var v185 float64
	_ = v185
	var v189 float64
	_ = v189
	var v194 int32
	_ = v194
	var v195 float64
	_ = v195
	var v197 float64
	_ = v197
	var v203 float64
	_ = v203
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v7 = float64(0)
	v13 = int32(0)
	v19 = base.F64_convert_i32_s(l4)
	if base.F64_eq(l5, v19)|base.F64_le(l5, float64(1)) != 0 {
		v219 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v219
L2:
	;
	if base.Ui32(l1) < base.Ui32(int32(2)) {
		v135 = v7
		goto L3
	} else {
		goto L4
	}
L3:
	;
	if base.F64_lt(l2, float64(0)) != 0 {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	v31 = l1 - int32(1)
	v32 = int32(3)
	v33 = v31 & v32
	v34 = int32(0)
	if base.Ui32(v32) <= base.Ui32(l1-int32(2)) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v45 = v34
	v48 = v7
	v57 = v13
	goto L8
L6:
	;
	v85 = v34
	v88 = v7
	goto L7
L7:
	;
	v103 = v85
	v106 = v88
	v116 = v13
	goto L12
L8:
	;
	v61 = l0 + v45<<(uint(int32(2))%32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v73 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v48, base.F64_convert_i32_s(v62)), base.F64_convert_i32_s(v65)), base.F64_convert_i32_s(v68)), base.F64_convert_i32_s(v71))
	v74 = int32(4)
	v75 = v45 + v74
	v77 = v57 + v74
	if v77 != v31&int32(-4) {
		v45 = v75
		v48 = v73
		v57 = v77
		goto L8
	} else {
		goto L10
	}
L9:
	;
	if v33 == int32(0) {
		v135 = v73
		goto L3
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	v85 = v75
	v88 = v73
	goto L7
L12:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0+v103<<(uint(int32(2))%32))))
	v122 = base.F64_add(v106, base.F64_convert_i32_s(v120))
	v123 = int32(1)
	v126 = v116 + v123
	if v126 != v33 {
		v103 = v103 + v123
		v106 = v122
		v116 = v126
		goto L12
	} else {
		goto L14
	}
L13:
	;
	v135 = v122
	goto L3
L14:
	;
	goto L13
L15:
	;
	v146 = base.F64_mul(l5, base.F64_neg(l2))
	goto L17
L16:
	;
	v146 = l2
	goto L17
L17:
	;
	v153 = l1
	v159 = v135
	goto L18
L18:
	;
	v170 = float64(1)
	v171 = float64(0)
	v175 = base.F64_sub(base.F64_sub(v170, base.F64_div(v159, v19)), l3)
	if base.F64_lt(v175, v171) != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v219 = int32(0)
	goto L1
L20:
	;
	v178 = v171
	goto L22
L21:
	;
	v178 = v175
	goto L22
L22:
	;
	if base.F64_gt(v178, float64(1)) != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v181 = v170
	goto L25
L24:
	;
	v181 = v178
	goto L25
L25:
	;
	v183 = v153 - int32(1)
	v185 = base.F64_sub(v146, base.F64_convert_i32_u(v183))
	if base.F64_gt(v185, float64(1)) != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v189 = base.F64_div(v181, v185)
	goto L28
L27:
	;
	v189 = v181
	goto L28
L28:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0+v183<<(uint(int32(2))%32))))
	v195 = base.F64_convert_i32_s(v194)
	v197 = base.F64_div(base.F64_mul(l5, v195), v19)
	v203 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l5, v19), base.F64_mul(base.F64_mul(v197, v19), base.F64_sub(l5, v197))), base.F64_mul(base.F64_mul(l5, l5), base.F64_add(l5, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v189, v19), base.F64_add(v203, v203)), float64(0.5)), v195) != 0 {
		v219 = v153
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v183 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0+v153<<(uint(int32(2))%32)-int32(8))))
	v153 = v183
	v159 = base.F64_sub(v159, base.F64_convert_i32_s(v214))
	goto L18
L31:
	;
	goto L32
L32:
	;
	goto L19
}
