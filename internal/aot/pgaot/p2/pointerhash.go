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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v238 int32
	_ = v238
	var v245 int32
	_ = v245
	goto L1
L1:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if base.Ui32(v28) <= base.Ui32(v27) {
		goto L7
	} else {
		goto L8
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(0)
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v238))) = l1
	v245 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v245)
	return v238
L5:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v226 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v225 + v226
	*(*uint8)(unsafe.Add(mBase, uint32(v219)+4)) = uint8(v226)
	v238 = v219
	goto L4
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L11
	} else {
		goto L46
	}
L7:
	;
	v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if v30 == int64(4294967296) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v40 = int32(0)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v43 = v42 & l2
	v46 = v41 + v43<<(uint(int32(3))%32)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)))
	if v47 == v40 {
		v219 = v46
		goto L5
	} else {
		goto L13
	}
L10:
	;
	F_pointerhash_grow(m, l0, v30<<(uint(int64(1))%64))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	goto L9
L13:
	;
	v55 = v40
	v56 = v43
	v57 = v46
	goto L14
L14:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	if l1 == v63 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v219 = v197
	goto L5
L16:
	;
	v65 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v65)
	return v57
L17:
	;
	goto L18
L18:
	;
	v69 = v56 + int32(1)
	v70 = int32(16)
	v74 = (int32(base.Ui32(v63)>>(uint(v70)%32)) ^ v63) * int32(-2048144789)
	v79 = (int32(base.Ui32(v74)>>(uint(int32(13))%32)) ^ v74) * int32(-1028477387)
	v83 = v42 & (int32(base.Ui32(v79)>>(uint(v70)%32)) ^ v79)
	if base.Ui32(v56) < base.Ui32(v83) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = v56 + v85
	goto L21
L20:
	;
	v87 = v56
	goto L21
L21:
	;
	if base.Ui32(v87-v83) < base.Ui32(v55) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v91 = v42 & v69
	v94 = v41 + v91<<(uint(int32(3))%32)
	v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94)+4)))
	if v95 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v184 = v55 + int32(1)
	if base.Ui32(int32(26)) <= base.Ui32(v184) {
		goto L41
	} else {
		goto L42
	}
L25:
	;
	v100 = int32(0)
	v101 = v91
	goto L28
L26:
	;
	v132 = v91
	v136 = v94
	goto L27
L27:
	;
	if v132 != v56 {
		goto L35
	} else {
		goto L36
	}
L28:
	;
	if base.Ui32(int32(150)) <= base.Ui32(v100) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v132 = v122
	v136 = v125
	goto L27
L30:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v111), base.F64_convert_i64_u(v113)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v118 = int32(1)
	v122 = (v101 + v118) & v42
	v125 = v41 + v122<<(uint(int32(3))%32)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125)+4)))
	if v126 != 0 {
		v100 = v100 + v118
		v101 = v122
		goto L28
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	goto L29
L35:
	;
	v146 = v132
	v150 = v136
	goto L38
L36:
	;
	goto L37
L37:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v178 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v177 + v178
	*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)) = uint8(v178)
	v238 = v57
	goto L4
L38:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v157 = v154 & (v146 - int32(1))
	v160 = v41 + v157<<(uint(int32(3))%32)
	v161 = *(*int64)(unsafe.Add(mBase, uint32(v160)))
	*(*int64)(unsafe.Add(mBase, uint32(v150))) = v161
	if v157 != v56 {
		v146 = v157
		v150 = v160
		goto L38
	} else {
		goto L40
	}
L39:
	;
	goto L37
L40:
	;
	goto L39
L41:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v189 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	if base.F64_ge(base.F64_div(base.F64_convert_i32_u(v187), base.F64_convert_i64_u(v189)), float64(0.1)) != 0 {
		goto L3
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v194 = v42 & v69
	v197 = v41 + v194<<(uint(int32(3))%32)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v197)+4)))
	if v198 != 0 {
		v55 = v184
		v56 = v194
		v57 = v197
		goto L14
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	goto L15
L46:
	;
	F_errmsg_internal(m, int32(485500), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	F_errfinish(m, int32(344283), int32(630), int32(327766))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
