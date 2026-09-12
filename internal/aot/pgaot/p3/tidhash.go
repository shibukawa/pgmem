package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_tidhash_delete_item(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var __phi27 int32
	_ = __phi27
	var v29 int32
	_ = v29
	var __phi29 int32
	_ = __phi29
	var v30 int32
	_ = v30
	var __phi30 int32
	_ = __phi30
	var v31 int32
	_ = v31
	var __phi31 int32
	_ = __phi31
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v47 int64
	_ = v47
	var v54 int64
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v8 - v9
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v15 = int32(3)
	v19 = v13 & ((l1-v12)>>(uint(v15)%32) + v9)
	v22 = v12 + v19<<(uint(v15)%32)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+6)))
	if v23 != v9 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v74 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v69)+6)) = uint8(v74)
	return
L2:
	;
	v69 = l1
	goto L1
L3:
	;
	goto L4
L4:
	;
	__phi27 = l1
	__phi29 = v19
	__phi30 = v13
	__phi31 = v22
	v27 = __phi27
	v29 = __phi29
	v30 = __phi30
	v31 = __phi31
	goto L5
L5:
	;
	v33 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)))
	v35 = v33 << (uint(int64(32)) % 64)
	v36 = int64(33)
	v38 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31))))
	v42 = (int64(base.Ui64(v35)>>(uint(v36)%64)) ^ (v38 | v35)) * int64(-49064778989728563)
	v47 = (int64(base.Ui64(v42)>>(uint(v36)%64)) ^ v42) * int64(-4265267296055464877)
	if v29 == v30&base.I32_wrap_i64(int64(base.Ui64(v47)>>(uint(v36)%64))^v47) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v69 = v31
	goto L1
L7:
	;
	v69 = v27
	goto L1
L8:
	;
	goto L9
L9:
	;
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v58 = int32(1)
	v60 = v57 & (v29 + v58)
	v63 = v56 + v60<<(uint(int32(3))%32)
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+6)))
	if v64 == v58 {
		__phi27 = v31
		__phi29 = v60
		__phi30 = v57
		__phi31 = v63
		v27 = __phi27
		v29 = __phi29
		v30 = __phi30
		v31 = __phi31
		goto L5
	} else {
		goto L10
	}
L10:
	;
	goto L6
}
func F_tidhash_grow(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int64
	_ = v98
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v103 int64
	_ = v103
	var v107 int64
	_ = v107
	var v112 int64
	_ = v112
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int64
	_ = v144
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v153 int64
	_ = v153
	var v158 int64
	_ = v158
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v4 = int32(0)
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
	v237 = m.ExcPending
	if v237 != 0 {
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
	v224 = m.ExcPending
	if v224 != 0 {
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
	v84 = v4
	goto L28
L25:
	;
	goto L26
L26:
	;
	F_pfree(m, v32)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L11
	} else {
		goto L45
	}
L27:
	;
	v130 = v125
	v134 = v4
	goto L33
L28:
	;
	v94 = v32 + v84<<(uint(int32(3))%32)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+6)))
	if v95 != int32(1) {
		v125 = v84
		goto L27
	} else {
		goto L30
	}
L29:
	;
	v125 = int32(0)
	goto L27
L30:
	;
	v98 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v94)+4)))
	v100 = v98 << (uint(int64(32)) % 64)
	v101 = int64(33)
	v103 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v94))))
	v107 = (int64(base.Ui64(v100)>>(uint(v101)%64)) ^ (v103 | v100)) * int64(-49064778989728563)
	v112 = (int64(base.Ui64(v107)>>(uint(v101)%64)) ^ v107) * int64(-4265267296055464877)
	if v60&base.I32_wrap_i64(int64(base.Ui64(v112)>>(uint(v101)%64))^v112) == v84 {
		v125 = v84
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v120 = v84 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v120)) < base.Ui64(v33) {
		v84 = v120
		goto L28
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	v140 = v32 + v130<<(uint(int32(3))%32)
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v140)+6)))
	if v141 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L26
L35:
	;
	v144 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v140)+4)))
	v146 = v144 << (uint(int64(32)) % 64)
	v147 = int64(33)
	v149 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v140))))
	v153 = (int64(base.Ui64(v146)>>(uint(v147)%64)) ^ (v149 | v146)) * int64(-49064778989728563)
	v158 = (int64(base.Ui64(v153)>>(uint(v147)%64)) ^ v153) * int64(-4265267296055464877)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v167 = base.I32_wrap_i64(int64(base.Ui64(v158)>>(uint(v147)%64)) ^ v158)
	goto L38
L36:
	;
	goto L37
L37:
	;
	v198 = v130 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v198)) < base.Ui64(v33) {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	v176 = v167 & v163
	v181 = v39 + v176<<(uint(int32(3))%32)
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181)+6)))
	if v182 != 0 {
		v167 = v176 + int32(1)
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v183 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
	*(*int64)(unsafe.Add(mBase, uint32(v181))) = v183
	goto L37
L40:
	;
	goto L39
L41:
	;
	v202 = v198
	goto L43
L42:
	;
	v202 = int32(0)
	goto L43
L43:
	;
	v204 = v134 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v204)) < base.Ui64(v33) {
		v130 = v202
		v134 = v204
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
	F_errmsg_internal(m, int32(421049), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(344283), int32(327), int32(359628))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
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
	F_errmsg_internal(m, int32(421049), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L11
	} else {
		goto L50
	}
L50:
	;
	F_errfinish(m, int32(344283), int32(327), int32(359628))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
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
func F_tidhash_insert(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	var v29 int64
	_ = v29
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = l1 + int32(4)
	v14 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	v15 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1))))
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v16)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v18
	v23 = v14 << (uint(int64(32)) % 64)
	v24 = int64(33)
	v29 = (int64(base.Ui64(v23)>>(uint(v24)%64)) ^ (v23 | v15)) * int64(-49064778989728563)
	v34 = (int64(base.Ui64(v29)>>(uint(v24)%64)) ^ v29) * int64(-4265267296055464877)
	v39 = F_tidhash_insert_hash_internal(m, l0, v10+int32(8), base.I32_wrap_i64(int64(base.Ui64(v34)>>(uint(v24)%64))^v34), l2)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		return int32(0)
	} else {
		m.G0 = v10 + int32(16)
		return v39
	}
}
func F_tidhash_reset(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v8 = F__emscripten_memset_bulkmem(m, v2, base.I32_extend8_s(int32(0)), v4<<(uint(int32(3))%32))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	return
}
