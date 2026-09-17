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
	var v3 int32
	_ = v3
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v26 int64
	_ = v26
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int64
	_ = v41
	var v51 int64
	_ = v51
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v98 int64
	_ = v98
	var v103 int64
	_ = v103
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int64
	_ = v134
	var v136 int64
	_ = v136
	var v137 int64
	_ = v137
	var v139 int64
	_ = v139
	var v143 int64
	_ = v143
	var v148 int64
	_ = v148
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int64
	_ = v172
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v207 int32
	_ = v207
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	v3 = int32(0)
	v12 = int64(2)
	if base.Ui64(l1) <= base.Ui64(v12) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v15 = v12
	goto L3
L2:
	;
	v15 = l1
	goto L3
L3:
	;
	v16 = int64(1)
	if v15&(v15-v16) == int64(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v15
	goto L6
L5:
	;
	v26 = v16 << (uint(int64(64)-base.I64_clz(v15)) % 64)
	goto L6
L6:
	;
	if base.Ui64(v26<<(uint(int64(3))%64)) < base.Ui64(int64(2147483647)) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v38 = F_MemoryContextAllocExtended(m, v33, base.I32_wrap_i64(v26)<<(uint(int32(3))%32), int32(5))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L11
	} else {
		goto L42
	}
L10:
	;
	goto L9
L11:
	;
	return
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v38
	v41 = int64(1)
	if v26&(v26-v41) == int64(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v51 = v26
	goto L15
L14:
	;
	v51 = v41 << (uint(int64(64)-base.I64_clz(v26)) % 64)
	goto L15
L15:
	;
	if base.Ui64(int64(2147483647)) <= base.Ui64(v51<<(uint(int64(3))%64)) {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v51
	v59 = base.I32_wrap_i64(v51) - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v59
	if v51 == int64(4294967296) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v68 = int32(-85899346)
	goto L19
L18:
	;
	v68 = base.I32_trunc_sat_f64_u(base.F64_mul(base.F64_convert_i64_u(v51), float64(0.9)))
	goto L19
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v68
	if v32 != int64(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v74 = v3
	goto L24
L21:
	;
	goto L22
L22:
	;
	F_pfree(m, v31)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L11
	} else {
		goto L41
	}
L23:
	;
	v119 = v116
	v126 = v3
	goto L29
L24:
	;
	v85 = v31 + v74<<(uint(int32(3))%32)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+6)))
	if v86 != int32(1) {
		v116 = v74
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v116 = int32(0)
	goto L23
L26:
	;
	v89 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	v91 = v89 << (uint(int64(32)) % 64)
	v92 = int64(33)
	v94 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v85))))
	v98 = (int64(base.Ui64(v91)>>(uint(v92)%64)) ^ (v94 | v91)) * int64(-49064778989728563)
	v103 = (int64(base.Ui64(v98)>>(uint(v92)%64)) ^ v98) * int64(-4265267296055464877)
	if v59&base.I32_wrap_i64(int64(base.Ui64(v103)>>(uint(v92)%64))^v103) == v74 {
		v116 = v74
		goto L23
	} else {
		goto L27
	}
L27:
	;
	v111 = v74 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v111)) < base.Ui64(v32) {
		v74 = v111
		goto L24
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	v130 = v31 + v119<<(uint(int32(3))%32)
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+6)))
	if v131 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L22
L31:
	;
	v134 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v130)+4)))
	v136 = v134 << (uint(int64(32)) % 64)
	v137 = int64(33)
	v139 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v130))))
	v143 = (int64(base.Ui64(v136)>>(uint(v137)%64)) ^ (v139 | v136)) * int64(-49064778989728563)
	v148 = (int64(base.Ui64(v143)>>(uint(v137)%64)) ^ v143) * int64(-4265267296055464877)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v158 = base.I32_wrap_i64(int64(base.Ui64(v148)>>(uint(v137)%64)) ^ v148)
	goto L34
L32:
	;
	goto L33
L33:
	;
	v186 = v119 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v186)) < base.Ui64(v32) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v165 = v158 & v153
	v170 = v38 + v165<<(uint(int32(3))%32)
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v170)+6)))
	if v171 != 0 {
		v158 = v165 + int32(1)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v172 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	*(*int64)(unsafe.Add(mBase, uint32(v170))) = v172
	goto L33
L36:
	;
	goto L35
L37:
	;
	v190 = v186
	goto L39
L38:
	;
	v190 = int32(0)
	goto L39
L39:
	;
	v192 = v126 + int32(1)
	if base.Ui64(base.I64_extend_i32_u(v192)) < base.Ui64(v32) {
		v119 = v190
		v126 = v192
		goto L29
	} else {
		goto L40
	}
L40:
	;
	goto L30
L41:
	;
	return
L42:
	;
	F_errmsg_internal(m, int32(_a_F_tidhash_grow_0), int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L11
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_tidhash_grow_1), int32(327), int32(_a_F_tidhash_grow_2))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L11
	} else {
		goto L44
	}
L44:
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
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = v3 << (uint(int32(3)) % 32)
	if v5 != 0 {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		base.MemoryFill(m, v6, int32(0), v5)
	} else {
	}
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	return
}
