package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_skip_utf8(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	if l3 < int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	goto L3
L3:
	;
	if l3 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v11 = l1
	v13 = l3
	goto L7
L5:
	;
	v48 = l1
	goto L6
L6:
	;
	return v48
L7:
	;
	if l2 <= v11 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v48 = v42
	goto L6
L9:
	;
	return int32(-1)
L10:
	;
	goto L11
L11:
	;
	v19 = v11 + int32(1)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v11))))
	if base.Ui32(v21) < base.Ui32(int32(192)) {
		v42 = v19
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = int32(1)
	if v43 < v13 {
		v11 = v42
		v13 = v13 - v43
		goto L7
	} else {
		goto L19
	}
L13:
	;
	if l2 <= v19 {
		v42 = v19
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v28 = v19
	goto L15
L15:
	;
	v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0+v28))))
	if int32(-65) < v31 {
		v42 = v28
		goto L12
	} else {
		goto L17
	}
L16:
	;
	v42 = l2
	goto L12
L17:
	;
	v35 = v28 + int32(1)
	if v35 != l2 {
		v28 = v35
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	goto L8
}
func F_utf8_to_big5(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(6), int32(36))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_UtfToLocal(m, v6, v10, v5, int32(_a_F_utf8_to_big5_0), v18, v18, v18, int32(36), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_utf8_to_euc_jis_2004(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(6), int32(5))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(0)
		v27 = F_UtfToLocal(m, v6, v10, v5, int32(_a_F_utf8_to_euc_jis_2004_0), int32(_a_F_utf8_to_euc_jis_2004_1), int32(25), v17, int32(5), base.B2i32(v7 != v17))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_utf8_to_euc_kr(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(6), int32(3))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_UtfToLocal(m, v6, v10, v5, int32(_a_F_utf8_to_euc_kr_0), v18, v18, v18, int32(3), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_utf8_to_iso8859_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v12, v13, v14, int32(6), int32(8))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v14 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v160 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v156))) = uint8(v160)
	return v152 - v11
L4:
	;
	v152 = v11
	v156 = v10
	goto L3
L5:
	;
	goto L6
L6:
	;
	v23 = v11
	v24 = v14
	v27 = v10
	goto L7
L7:
	;
	v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	if v31 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v152 = v148
	v156 = v147
	goto L3
L9:
	;
	if v9 != 0 {
		v152 = v23
		v156 = v27
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v31 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	F_report_invalid_encoding(m, int32(6), v23, v24)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	if int32(0) <= v41 {
		goto L19
	} else {
		goto L20
	}
L15:
	;
	v142 = int32(-1)
	v143 = int32(1)
	v144 = v31
	goto L16
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v144)
	v147 = v27 + int32(1)
	v148 = v23 + v143
	v149 = v24 + v142
	if int32(0) < v149 {
		v23 = v148
		v24 = v149
		v27 = v147
		goto L7
	} else {
		goto L68
	}
L17:
	;
	if v65 != int32(2) {
		goto L58
	} else {
		goto L59
	}
L18:
	;
	if v65 <= v24 {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	v65 = int32(1)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v46 = v41 & int32(255)
	if v46&int32(224) == int32(192) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v65 = int32(2)
	goto L18
L23:
	;
	goto L24
L24:
	;
	if v46&int32(240) == int32(224) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v65 = int32(3)
	goto L18
L26:
	;
	goto L27
L27:
	;
	if v46&int32(248) == int32(240) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v63 = int32(4)
	goto L30
L29:
	;
	v63 = int32(1)
	goto L30
L30:
	;
	v65 = v63
	goto L18
L31:
	;
	v67 = int32(0)
	switch v65 - int32(1) {
	case 0:
		goto L38
	case 1:
		goto L39
	case 2:
		goto L40
	case 3:
		goto L41
	default:
		v116 = v67
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	if v9 != 0 {
		v152 = v23
		v156 = v27
		goto L3
	} else {
		goto L56
	}
L34:
	;
	if v116 != 0 {
		goto L17
	} else {
		goto L55
	}
L35:
	;
	goto L34
L36:
	;
	v116 = base.B2i32(base.Ui32(v108&int32(255)) < base.Ui32(int32(245)))
	goto L35
L37:
	;
	if base.I32_extend8_s(v103) < int32(-62) {
		v116 = v67
		goto L35
	} else {
		goto L54
	}
L38:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v103 = v102
	goto L37
L39:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+1)))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	switch v77 - int32(224) {
	case 0:
		goto L48
	default:
		goto L44
	case 13:
		goto L47
	case 16:
		goto L46
	case 20:
		goto L45
	}
L40:
	;
	v73 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+2)))
	if int32(-65) < v73 {
		v116 = v67
		goto L35
	} else {
		goto L43
	}
L41:
	;
	v70 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23)+3)))
	if int32(-65) < v70 {
		v116 = v67
		goto L35
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	goto L39
L44:
	;
	if v76 <= int32(-65) {
		v103 = v77
		goto L37
	} else {
		goto L53
	}
L45:
	;
	if int32(-113) < v76 {
		v116 = v67
		goto L35
	} else {
		goto L52
	}
L46:
	;
	if base.Ui32((v76-int32(-64))&int32(255)) < base.Ui32(int32(208)) {
		v116 = v67
		goto L35
	} else {
		goto L51
	}
L47:
	;
	if int32(-97) < v76 {
		v116 = v67
		goto L35
	} else {
		goto L50
	}
L48:
	;
	v80 = int32(224)
	if base.Ui32(v80) <= base.Ui32((v76-int32(-64))&int32(255)) {
		v108 = v80
		goto L36
	} else {
		goto L49
	}
L49:
	;
	v116 = v67
	goto L35
L50:
	;
	v108 = int32(237)
	goto L36
L51:
	;
	v108 = int32(240)
	goto L36
L52:
	;
	v108 = int32(244)
	goto L36
L53:
	;
	v116 = v67
	goto L35
L54:
	;
	v108 = v103
	goto L36
L55:
	;
	goto L33
L56:
	;
	F_report_invalid_encoding(m, int32(6), v23, v24)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L58:
	;
	if v9 != 0 {
		v152 = v23
		v156 = v27
		goto L3
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v31&int32(30) != int32(2) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	F_report_untranslatable_char(m, int32(6), int32(8), v23, v24)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L63:
	;
	if v9 != 0 {
		v152 = v23
		v156 = v27
		goto L3
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	v142 = int32(-2)
	v143 = int32(2)
	v144 = v136&int32(63) | v31<<(uint(int32(6))%32)
	goto L16
L66:
	;
	F_report_untranslatable_char(m, int32(6), int32(8), v23, v24)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	goto L8
}
func F_utf8_to_koi8r(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(6), int32(22))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_UtfToLocal(m, v6, v10, v5, int32(_a_F_utf8_to_koi8r_0), v18, v18, v18, int32(22), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
