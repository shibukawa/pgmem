package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_mic_to_euc_jp(m *base.Module, l0 int32) int32 {
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
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, int32(7), int32(1))
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
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v87))) = uint8(v93)
	return v86 - v10
L4:
	;
	v86 = v10
	v87 = v9
	goto L3
L5:
	;
	goto L6
L6:
	;
	v22 = v10
	v23 = v9
	v24 = v13
	goto L7
L7:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	v30 = base.I32_extend8_s(v29)
	if int32(0) <= v30 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v86 = v83
	v87 = v80
	goto L3
L9:
	;
	if int32(0) < v81 {
		v22 = v83
		v23 = v80
		v24 = v81
		goto L7
	} else {
		goto L32
	}
L10:
	;
	if v30 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v46 = F_pg_encoding_verifymbchar(m, int32(7), v22, v24)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v8 != 0 {
		v86 = v22
		v87 = v23
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v30)
	v39 = int32(1)
	v80 = v23 + v39
	v81 = v24 - v39
	v83 = v22 + v39
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(7), v22, v24)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
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
	if v46 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v8 != 0 {
		v86 = v22
		v87 = v23
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	switch v29 - int32(137) {
	case 0:
		goto L26
	default:
		goto L27
	case 9:
		goto L28
	case 11:
		goto L29
	}
L22:
	;
	F_report_invalid_encoding(m, int32(7), v22, v24)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	v80 = v77
	v81 = v24 - v46
	v83 = v22 + v46
	goto L9
L25:
	;
	v77 = v23 + int32(2)
	goto L24
L26:
	;
	v71 = int32(142)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v71)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v73)
	goto L25
L27:
	;
	if v8 != 0 {
		v86 = v22
		v87 = v23
		goto L3
	} else {
		goto L30
	}
L28:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v63)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v65)
	goto L25
L29:
	;
	v55 = int32(143)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v55)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v57)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)) = uint8(v59)
	v77 = v23 + int32(3)
	goto L24
L30:
	;
	F_report_untranslatable_char(m, int32(7), int32(1), v22, v24)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	goto L8
}
func F_mic_to_win866(m *base.Module, l0 int32) int32 {
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
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(7), int32(20))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v22 = F_mic2latin_with_table(m, v6, v5, v10, int32(139), int32(20), int32(2234320), base.B2i32(v7 != int32(0)))
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return int32(0)
		} else {
			return v22
		}
	}
}
