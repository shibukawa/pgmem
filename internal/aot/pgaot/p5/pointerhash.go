package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pointerhash_grow(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int64
	_ = v13
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v27 int64
	_ = v27
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v52 int64
	_ = v52
	var v60 int32
	_ = v60
	var v67 float64
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int64
	_ = v173
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	v3 = int32(0)
	v13 = int64(2)
	if base.Ui64(l1) <= base.Ui64(v13) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L11
	} else {
		goto L49
	}
L2:
	;
	v16 = v13
	goto L4
L3:
	;
	v16 = l1
	goto L4
L4:
	;
	v17 = int64(1)
	if v16&(v16-v17) == int64(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = v16
	goto L7
L6:
	;
	v27 = v17 << (uint(int64(64)-base.I64_clz(v16)) % 64)
	goto L7
L7:
	;
	if base.Ui64(v27<<(uint(int64(3))%64)) < base.Ui64(int64(2147483647)) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v39 = F_MemoryContextAllocExtended(m, v34, base.I32_wrap_i64(v27)<<(uint(int32(3))%32), int32(5))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L11
	} else {
		goto L46
	}
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v39
	v42 = int64(1)
	if v27&(v27-v42) == int64(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = v27
	goto L15
L14:
	;
	v52 = v42 << (uint(int64(64)-base.I64_clz(v27)) % 64)
	goto L15
L15:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v52<<(uint(int64(3))%64)) {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v52
	v60 = base.I32_wrap_i64(v52) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v60
	v67 = base.F64_mul(base.F64_convert_i64_u(v52), float64(0.9))
	if base.F64_lt(v67, float64(4.294967296e+09))&base.F64_ge(v67, float64(0)) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v52 == int64(4294967296) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v73 = base.I32_trunc_f64_u(v67)
	v75 = v73
	goto L17
L19:
	;
	goto L20
L20:
	;
	v75 = int32(0)
	goto L17
L21:
	;
	v76 = int32(-85899346)
	goto L23
L22:
	;
	v76 = v75
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v76
	if v33 != int64(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v83 = v3
	goto L28
L25:
	;
	goto L26
L26:
	;
	F_pfree(m, v32)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L11
	} else {
		goto L45
	}
L27:
	;
	v124 = v120
	v127 = v3
	goto L33
L28:
	;
	v94 = v32 + v83<<(uint(int32(3))%32)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	if v95 != int32(1) {
		v120 = v83
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v120 = int32(0)
	goto L27
L30:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v99 = int32(16)
	v103 = (int32(base.Ui32(v98)>>(uint(v99)%32)) ^ v98) * int32(-2048144789)
	v108 = (int32(base.Ui32(v103)>>(uint(int32(13))%32)) ^ v103) * int32(-1028477387)
	if (int32(base.Ui32(v108)>>(uint(v99)%32))^v108)&v60 == v83 {
		v120 = v83
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v115 = v83 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v115)) < base.Ui64(v33) {
		v83 = v115
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v135 = v32 + v124<<(uint(int32(3))%32)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+4)))
	if v136 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L26
L35:
	;
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	v140 = int32(16)
	v144 = (int32(base.Ui32(v139)>>(uint(v140)%32)) ^ v139) * int32(-2048144789)
	v149 = (int32(base.Ui32(v144)>>(uint(int32(13))%32)) ^ v144) * int32(-1028477387)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v156 = int32(base.Ui32(v149)>>(uint(v140)%32)) ^ v149
	goto L38
L36:
	;
	goto L37
L37:
	;
	v188 = v124 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v188)) < base.Ui64(v33) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v166 = v156 & v153
	v171 = v39 + v166<<(uint(int32(3))%32)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v171)+4)))
	if v172 != 0 {
		v156 = v166 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v173 = *(*int64)(unsafe.Add(mBase, uint32(v135)))
	*(*int64)(unsafe.Add(mBase, uint32(v171))) = v173
	goto L37
L40:
	;
	goto L39
L41:
	;
	v192 = v188
	goto L43
L42:
	;
	v192 = int32(0)
	goto L43
L43:
	;
	v194 = v127 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v194)) < base.Ui64(v33) {
		v124 = v192
		v127 = v194
		goto L33
	} else {
		goto L44
	}
L44:
	;
	goto L34
L45:
	;
	return
L46:
	;
	F_errmsg_internal(m, int32(409411), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(333945), int32(327), int32(348830))
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L49:
	;
	F_errmsg_internal(m, int32(409411), int32(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(333945), int32(327), int32(348830))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
