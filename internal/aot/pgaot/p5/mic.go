package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_mic_to_euc_cn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v10, v11, v12, int32(7), int32(2))
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
	if v12 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v79)
	return v73 - v9
L4:
	;
	v73 = v9
	v74 = v8
	goto L3
L5:
	;
	goto L6
L6:
	;
	v21 = v9
	v22 = v8
	v24 = v12
	goto L7
L7:
	;
	v27 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21))))
	if v27 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v73 = v69
	v74 = v67
	goto L3
L9:
	;
	v70 = v68 + v24
	if int32(0) < v70 {
		v21 = v69
		v22 = v67
		v24 = v70
		goto L7
	} else {
		goto L30
	}
L10:
	;
	if v27 != int32(-111) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	if v27 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L13:
	;
	if v7 != 0 {
		v73 = v21
		v74 = v22
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	if base.Ui32(v24) < base.Ui32(int32(3)) {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	F_report_untranslatable_char(m, int32(7), int32(2), v21, v24)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v38)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)) = uint8(v49)
	v67 = v22 + int32(2)
	v68 = int32(-3)
	v69 = v21 + int32(3)
	goto L9
L19:
	;
	if v7 != 0 {
		v73 = v21
		v74 = v22
		goto L3
	} else {
		goto L23
	}
L20:
	;
	v38 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21)+1)))
	if int32(0) <= v38 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v21)+2)))
	if v41 < int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	F_report_invalid_encoding(m, int32(7), v21, v24)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L25:
	;
	if v7 != 0 {
		v73 = v21
		v74 = v22
		goto L3
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22))) = uint8(v27)
	v62 = int32(1)
	v67 = v22 + v62
	v68 = int32(-1)
	v69 = v21 + v62
	goto L9
L28:
	;
	F_report_invalid_encoding(m, int32(7), v21, v24)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L30:
	;
	goto L8
}
func F_mic_to_euc_kr(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, int32(7), int32(3))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v13 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v74))) = uint8(v79)
	return v72 - v10
L4:
	;
	v72 = v10
	v74 = v9
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = v10
	v23 = v13
	v24 = v9
	goto L7
L7:
	;
	v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22))))
	if int32(0) <= v29 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v72 = v66
	v74 = v68
	goto L3
L9:
	;
	if int32(0) < v67 {
		v22 = v66
		v23 = v67
		v24 = v68
		goto L7
	} else {
		goto L29
	}
L10:
	;
	if v29 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v45 = F_pg_encoding_verifymbchar(m, int32(7), v22, v23)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v8 != 0 {
		v72 = v22
		v74 = v24
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v29)
	v38 = int32(1)
	v66 = v22 + v38
	v67 = v23 - v38
	v68 = v24 + v38
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(7), v22, v23)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v45 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v8 != 0 {
		v72 = v22
		v74 = v24
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v29 == int32(-109) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	F_report_invalid_encoding(m, int32(7), v22, v23)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
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
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v54)
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v56)
	v66 = v22 + v45
	v67 = v23 - v45
	v68 = v24 + int32(2)
	goto L9
L25:
	;
	goto L26
L26:
	;
	if v8 != 0 {
		v72 = v22
		v74 = v24
		goto L3
	} else {
		goto L27
	}
L27:
	;
	F_report_untranslatable_char(m, int32(7), int32(3), v22, v23)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	goto L8
}
func F_mic_to_iso(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(7), int32(25))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_mic2latin_with_table(m, v6, v5, v10, int32(139), int32(25), int32(2200880), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
func F_mic_to_koi8r(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(7), int32(22))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_mic2latin(m, v6, v5, v10, int32(139), int32(22), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
func F_mic_to_latin3(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(7), int32(10))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v21 = F_mic2latin(m, v6, v5, v10, int32(131), int32(10), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			return v21
		}
	}
}
func F_mic_to_sjis(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v222 int32
	_ = v222
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v14, v15, v16, int32(7), int32(35))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v16 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v222 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v213))) = uint8(v222)
	return v212 - v13
L4:
	;
	v212 = v13
	v213 = v12
	goto L3
L5:
	;
	goto L6
L6:
	;
	v25 = v13
	v26 = v12
	v28 = v16
	goto L7
L7:
	;
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v36 = base.I32_extend8_s(v35)
	if int32(0) <= v36 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v212 = v209
	v213 = v200
	goto L3
L9:
	;
	if int32(0) < v202 {
		v25 = v209
		v26 = v200
		v28 = v202
		goto L7
	} else {
		goto L62
	}
L10:
	;
	if v36 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v52 = F_pg_encoding_verifymbchar(m, int32(7), v25, v28)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v11 != 0 {
		v212 = v25
		v213 = v26
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v36)
	v45 = int32(1)
	v200 = v26 + v45
	v202 = v28 - v45
	v209 = v25 + v45
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(7), v25, v28)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	if v52 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v11 != 0 {
		v212 = v25
		v213 = v26
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	switch v35 - int32(137) {
	case 0:
		goto L25
	default:
		goto L26
	case 9:
		goto L28
	case 11:
		goto L27
	}
L22:
	;
	F_report_invalid_encoding(m, int32(7), v25, v28)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
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
	v200 = v196
	v202 = v28 - v52
	v209 = v25 + v52
	goto L9
L25:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v182)
	v196 = v26 + int32(1)
	goto L24
L26:
	;
	if v11 != 0 {
		v212 = v25
		v213 = v26
		goto L3
	} else {
		goto L60
	}
L27:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v106 = v102 | v103<<(uint(int32(8))%32)
	if base.Ui32(v106) <= base.Ui32(int32(62880)) {
		goto L44
	} else {
		goto L45
	}
L28:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+2)))
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if base.Ui32(int32(62881)) <= base.Ui32(v61|v62<<(uint(int32(8))%32)) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v86)
	if base.Ui32(v61) < base.Ui32(int32(224)) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v86 = (v62-int32(245))>>(uint(int32(1))%32) + int32(240)
	v87 = v62 - int32(84)
	goto L29
L31:
	;
	goto L32
L32:
	;
	if base.Ui32(v62) < base.Ui32(int32(223)) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v84 = int32(129)
	goto L35
L34:
	;
	v84 = int32(193)
	goto L35
L35:
	;
	v86 = int32(base.Ui32(v62+int32(351))>>(uint(int32(1))%32)) + v84
	v87 = v62
	goto L29
L36:
	;
	v93 = int32(-97)
	goto L38
L37:
	;
	v93 = int32(-96)
	goto L38
L38:
	;
	if v87&int32(1) != 0 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v97 = v93
	goto L41
L40:
	;
	v97 = int32(-2)
	goto L41
L41:
	;
	v98 = v97 + v61
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)) = uint8(v98)
	v196 = v26 + int32(2)
	goto L24
L42:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+2)))
	v170 = int32(8)
	v174 = v169<<(uint(v170)%32) | int32(base.Ui32(v169)>>(uint(v170)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v174)
	v196 = v26 + int32(2)
	goto L24
L43:
	;
	v137 = int32(0)
	goto L54
L44:
	;
	if v106 != int32(62451) {
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v120 = int32(base.Ui32(v103+int32(267))>>(uint(int32(1))%32)) - int32(11)
	*(*uint8)(unsafe.Add(mBase, uint32(v26))) = uint8(v120)
	if base.Ui32(v102) < base.Ui32(int32(224)) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v165 = int32(2202464)
	goto L42
L48:
	;
	v126 = int32(-97)
	goto L50
L49:
	;
	v126 = int32(-96)
	goto L50
L50:
	;
	if v103&int32(1) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v130 = v126
	goto L53
L52:
	;
	v130 = int32(-2)
	goto L53
L53:
	;
	v131 = v130 + v102
	*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)) = uint8(v131)
	v196 = v26 + int32(2)
	goto L24
L54:
	;
	v146 = v137 + int32(1)
	v148 = v146 << (uint(int32(3)) % 32)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+uint32(_consts[1450]))))
	if v151 == int32(65535) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v165 = v148 + int32(2202464)
	goto L42
L56:
	;
	v154 = int32(44161)
	*(*uint16)(unsafe.Add(mBase, uint32(v26))) = uint16(v154)
	v196 = v26 + int32(2)
	goto L24
L57:
	;
	goto L58
L58:
	;
	if v151 != v106 {
		v137 = v146
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	F_report_untranslatable_char(m, int32(7), int32(35), v25, v28)
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L62:
	;
	goto L8
}
