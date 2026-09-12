package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_euc_cn_to_mic(m *base.Module, l0 int32) int32 {
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
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, int32(2), int32(7))
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
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v73)
	return v66 - v10
L4:
	;
	v66 = v10
	v67 = v9
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
	v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22))))
	if v29 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v66 = v62
	v67 = v61
	goto L3
L9:
	;
	v62 = v22 + v60
	v63 = v24 + v59
	if int32(0) < v63 {
		v22 = v62
		v23 = v61
		v24 = v63
		goto L7
	} else {
		goto L25
	}
L10:
	;
	if v24 != int32(1) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	goto L12
L12:
	;
	if v29 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L13:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v29)
	v41 = int32(145)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v41)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)) = uint8(v43)
	v59 = int32(-2)
	v60 = int32(2)
	v61 = v23 + int32(3)
	goto L9
L14:
	;
	v34 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v34 < int32(0) {
		goto L13
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v8 != 0 {
		v66 = v22
		v67 = v23
		goto L3
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	F_report_invalid_encoding(m, int32(2), v22, v24)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	if v8 != 0 {
		v66 = v22
		v67 = v23
		goto L3
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v29)
	v55 = int32(1)
	v59 = int32(-1)
	v60 = v55
	v61 = v23 + v55
	goto L9
L23:
	;
	F_report_invalid_encoding(m, int32(2), v22, v24)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
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
	goto L8
}
func F_euc_cn_to_utf8(m *base.Module, l0 int32) int32 {
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
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(2), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4356596), v18, v18, v18, int32(2), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
func F_euc_kr_to_mic(m *base.Module, l0 int32) int32 {
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
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v11, v12, v13, int32(3), int32(7))
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
	v73 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v67))) = uint8(v73)
	return v66 - v10
L4:
	;
	v66 = v10
	v67 = v9
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
	v29 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22))))
	if v29 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v66 = v62
	v67 = v61
	goto L3
L9:
	;
	v62 = v22 + v60
	v63 = v24 + v59
	if int32(0) < v63 {
		v22 = v62
		v23 = v61
		v24 = v63
		goto L7
	} else {
		goto L24
	}
L10:
	;
	v33 = F_pg_encoding_verifymbchar(m, int32(3), v22, v24)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v29 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	if v33 != int32(2) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v8 != 0 {
		v66 = v22
		v67 = v23
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)) = uint8(v29)
	v41 = int32(147)
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v41)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)) = uint8(v43)
	v59 = int32(-2)
	v60 = int32(2)
	v61 = v23 + int32(3)
	goto L9
L17:
	;
	F_report_invalid_encoding(m, int32(3), v22, v24)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L19:
	;
	if v8 != 0 {
		v66 = v22
		v67 = v23
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v23))) = uint8(v29)
	v55 = int32(1)
	v59 = int32(-1)
	v60 = v55
	v61 = v23 + v55
	goto L9
L22:
	;
	F_report_invalid_encoding(m, int32(3), v22, v24)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
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
	goto L8
}
func F_euc_kr_to_utf8(m *base.Module, l0 int32) int32 {
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
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(3), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4356916), v18, v18, v18, int32(3), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
