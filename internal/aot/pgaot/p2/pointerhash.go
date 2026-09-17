package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pointerhash_insert_hash(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = F_pointerhash_insert_hash_internal(m, l0, l1, l2, l3)
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_pointerhash_insert_hash_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v109 int64
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v176 int64
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	goto L1
L1:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v26) <= base.Ui32(v25) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L4:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v212 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v211 + v212
	*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)) = uint8(v212)
	*(*int32)(unsafe.Add(mBase, uint32(v206))) = l1
	v218 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v218)
	return v206
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L10
	} else {
		goto L45
	}
L6:
	;
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v28 == int64(4294967296) {
		goto L5
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v38 = int32(0)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v41 = v40 & l2
	v44 = v39 + v41<<(uint(int32(3))%32)
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+4)))
	if v45 == v38 {
		v206 = v44
		goto L4
	} else {
		goto L12
	}
L9:
	;
	F_pointerhash_grow(m, l0, v28<<(uint(int64(1))%64))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return int32(0)
L11:
	;
	goto L8
L12:
	;
	v53 = v38
	v54 = v41
	v55 = v44
	goto L13
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	if l1 == v60 {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v206 = v184
	goto L4
L15:
	;
	v62 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v62)
	return v55
L16:
	;
	goto L17
L17:
	;
	v66 = v54 + int32(1)
	v67 = int32(16)
	v71 = (int32(base.Ui32(v60)>>(uint(v67)%32)) ^ v60) * int32(-2048144789)
	v76 = (int32(base.Ui32(v71)>>(uint(int32(13))%32)) ^ v71) * int32(-1028477387)
	v80 = v40 & (int32(base.Ui32(v76)>>(uint(v67)%32)) ^ v76)
	if base.Ui32(v54) < base.Ui32(v80) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v84 = v54 + v82
	goto L20
L19:
	;
	v84 = v54
	goto L20
L20:
	;
	if base.Ui32(v84-v80) < base.Ui32(v53) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v88 = v66 & v40
	v91 = v39 + v88<<(uint(int32(3))%32)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+4)))
	if v92 != 0 {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v171 = v53 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v171) {
		goto L40
	} else {
		goto L41
	}
L24:
	;
	v97 = int32(0)
	v98 = v88
	goto L27
L25:
	;
	v128 = v88
	v129 = v91
	goto L26
L26:
	;
	if v128 != v54 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	if base.Ui32(int32(150)) <= base.Ui32(v97) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v128 = v118
	v129 = v121
	goto L26
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v109 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v107), base.F64_convert_i64_u(v109)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v114 = int32(1)
	v118 = (v98 + v114) & v40
	v121 = v39 + v118<<(uint(int32(3))%32)
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+4)))
	if v122 != 0 {
		v97 = v97 + v114
		v98 = v118
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L28
L34:
	;
	v141 = v128
	v142 = v129
	goto L37
L35:
	;
	goto L36
L36:
	;
	v206 = v55
	goto L4
L37:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v151 = v148 & (v141 - int32(1))
	v154 = v39 + v151<<(uint(int32(3))%32)
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v154)))
	*(*int64)(unsafe.Add(mBase, uint32(v142))) = v155
	if v151 != v54 {
		v141 = v151
		v142 = v154
		goto L37
	} else {
		goto L39
	}
L38:
	;
	goto L36
L39:
	;
	goto L38
L40:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v176 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v174), base.F64_convert_i64_u(v176)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v181 = v66 & v40
	v184 = v39 + v181<<(uint(int32(3))%32)
	v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+4)))
	if v185 != 0 {
		v53 = v171
		v54 = v181
		v55 = v184
		goto L13
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	goto L14
L45:
	;
	F_errmsg_internal(m, int32(_a_F_pointerhash_insert_hash_internal_0), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L10
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_pointerhash_insert_hash_internal_1), int32(630), int32(_a_F_pointerhash_insert_hash_internal_2))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L10
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
