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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v45 int32
	_ = v45
	var v47 float64
	_ = v47
	var v56 int32
	_ = v56
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
	var v83 int32
	_ = v83
	var v85 float64
	_ = v85
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v122 float64
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v134 float64
	_ = v134
	var v146 float64
	_ = v146
	var v155 int32
	_ = v155
	var v160 float64
	_ = v160
	var v172 float64
	_ = v172
	var v173 float64
	_ = v173
	var v177 float64
	_ = v177
	var v180 float64
	_ = v180
	var v183 float64
	_ = v183
	var v185 int32
	_ = v185
	var v187 float64
	_ = v187
	var v191 float64
	_ = v191
	var v196 int32
	_ = v196
	var v197 float64
	_ = v197
	var v199 float64
	_ = v199
	var v205 float64
	_ = v205
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	v7 = float64(0)
	v13 = int32(0)
	v19 = base.F64_convert_i32_s(l4)
	if base.F64_eq(l5, v19) != 0 {
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
	if base.F64_le(l5, float64(1)) != 0 {
		v219 = l1
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if base.Ui32(l1) < base.Ui32(int32(2)) {
		v134 = v7
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if base.F64_lt(l2, float64(0)) != 0 {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	v30 = l1 - int32(1)
	v31 = int32(3)
	v32 = v30 & v31
	if base.Ui32(l1-int32(2)) < base.Ui32(v31) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if v32 == int32(0) {
		v134 = v85
		goto L4
	} else {
		goto L13
	}
L7:
	;
	v83 = int32(0)
	v85 = v7
	goto L6
L8:
	;
	goto L9
L9:
	;
	v45 = int32(0)
	v47 = v7
	v56 = v13
	goto L10
L10:
	;
	v61 = l0 + v45<<(uint(int32(2))%32)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)+12))
	v73 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v47, base.F64_convert_i32_s(v62)), base.F64_convert_i32_s(v65)), base.F64_convert_i32_s(v68)), base.F64_convert_i32_s(v71))
	v74 = int32(4)
	v75 = v45 + v74
	v77 = v56 + v74
	if v77 != v30&int32(-4) {
		v45 = v75
		v47 = v73
		v56 = v77
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v83 = v75
	v85 = v73
	goto L6
L12:
	;
	goto L11
L13:
	;
	v103 = v83
	v105 = v85
	v113 = v13
	goto L14
L14:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0+v103<<(uint(int32(2))%32))))
	v122 = base.F64_add(v105, base.F64_convert_i32_s(v120))
	v123 = int32(1)
	v126 = v113 + v123
	if v126 != v32 {
		v103 = v103 + v123
		v105 = v122
		v113 = v126
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v134 = v122
	goto L4
L16:
	;
	goto L15
L17:
	;
	v146 = base.F64_mul(l5, base.F64_neg(l2))
	goto L19
L18:
	;
	v146 = l2
	goto L19
L19:
	;
	v155 = l1
	v160 = v134
	goto L20
L20:
	;
	v172 = float64(1)
	v173 = float64(0)
	v177 = base.F64_sub(base.F64_sub(v172, base.F64_div(v160, v19)), l3)
	if base.F64_lt(v177, v173) != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v219 = int32(0)
	goto L1
L22:
	;
	v180 = v173
	goto L24
L23:
	;
	v180 = v177
	goto L24
L24:
	;
	if base.F64_gt(v180, float64(1)) != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v183 = v172
	goto L27
L26:
	;
	v183 = v180
	goto L27
L27:
	;
	v185 = v155 - int32(1)
	v187 = base.F64_sub(v146, base.F64_convert_i32_u(v185))
	if base.F64_gt(v187, float64(1)) != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v191 = base.F64_div(v183, v187)
	goto L30
L29:
	;
	v191 = v183
	goto L30
L30:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0+v185<<(uint(int32(2))%32))))
	v197 = base.F64_convert_i32_s(v196)
	v199 = base.F64_div(base.F64_mul(l5, v197), v19)
	v205 = base.F64_sqrt(base.F64_div(base.F64_mul(base.F64_sub(l5, v19), base.F64_mul(base.F64_mul(v199, v19), base.F64_sub(l5, v199))), base.F64_mul(base.F64_mul(l5, l5), base.F64_add(l5, float64(-1)))))
	if base.F64_lt(base.F64_add(base.F64_add(base.F64_mul(v191, v19), base.F64_add(v205, v205)), float64(0.5)), v197) != 0 {
		v219 = v155
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v185 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0-int32(8)+v155<<(uint(int32(2))%32))))
	v155 = v185
	v160 = base.F64_sub(v160, base.F64_convert_i32_s(v214))
	goto L20
L33:
	;
	goto L34
L34:
	;
	goto L21
}
