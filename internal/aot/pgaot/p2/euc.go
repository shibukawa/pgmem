package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_euc_jp_to_mic(m *base.Module, l0 int32) int32 {
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
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v12, v13, v14, int32(1), int32(7))
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
	v93 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v86))) = uint8(v93)
	return v85 - v11
L4:
	;
	v85 = v11
	v86 = v10
	goto L3
L5:
	;
	goto L6
L6:
	;
	v23 = v11
	v24 = v10
	v25 = v14
	goto L7
L7:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
	v32 = base.I32_extend8_s(v31)
	if int32(0) <= v32 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v85 = v82
	v86 = v79
	goto L3
L9:
	;
	if int32(0) < v80 {
		v23 = v82
		v24 = v79
		v25 = v80
		goto L7
	} else {
		goto L29
	}
L10:
	;
	if v32 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v48 = F_pg_encoding_verifymbchar(m, int32(1), v23, v25)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L18
	}
L13:
	;
	if v9 != 0 {
		v85 = v23
		v86 = v24
		goto L3
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v32)
	v41 = int32(1)
	v79 = v24 + v41
	v80 = v25 - v41
	v82 = v23 + v41
	goto L9
L16:
	;
	F_report_invalid_encoding(m, int32(1), v23, v25)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
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
	if v48 < int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	if v9 != 0 {
		v85 = v23
		v86 = v24
		goto L3
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	switch v31 - int32(142) {
	case 0:
		goto L28
	case 1:
		goto L27
	default:
		goto L26
	}
L22:
	;
	F_report_invalid_encoding(m, int32(1), v23, v25)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
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
	v79 = v76
	v80 = v25 - v48
	v82 = v23 + v48
	goto L9
L25:
	;
	v76 = v24 + int32(3)
	goto L24
L26:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v32)
	v70 = int32(146)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v70)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)) = uint8(v72)
	goto L25
L27:
	;
	v63 = int32(148)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v63)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v65)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)) = uint8(v67)
	goto L25
L28:
	;
	v57 = int32(137)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v57)
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v59)
	v76 = v24 + int32(2)
	goto L24
L29:
	;
	goto L8
}
func F_euc_jp_to_utf8(m *base.Module, l0 int32) int32 {
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
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(1), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4432180), v18, v18, v18, int32(1), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
