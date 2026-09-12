package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_euc_tw_to_mic(m *base.Module, l0 int32) int32 {
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v12, v13, v14, int32(4), int32(7))
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
	v107 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v100))) = uint8(v107)
	return v99 - v11
L4:
	;
	v99 = v11
	v100 = v10
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
	v31 = int32(*(*int8)(unsafe.Add(mBase, uint32(v23))))
	if v31 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v99 = v96
	v100 = v91
	goto L3
L9:
	;
	if int32(0) < v92 {
		v23 = v96
		v24 = v91
		v25 = v92
		goto L7
	} else {
		goto L32
	}
L10:
	;
	v35 = F_pg_encoding_verifymbchar(m, int32(4), v23, v25)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if v31 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L13:
	;
	if v35 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	if v9 != 0 {
		v99 = v23
		v100 = v24
		goto L3
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	if v31 == int32(-114) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	F_report_invalid_encoding(m, int32(4), v23, v25)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v91 = v76
	v92 = v25 - v35
	v96 = v23 + v35
	goto L9
L20:
	;
	v45 = v24 + int32(1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	switch v46 - int32(161) {
	case 0:
		goto L26
	case 1:
		goto L25
	default:
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v31)
	v68 = int32(149)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v68)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+2)) = uint8(v70)
	v76 = v24 + int32(3)
	goto L19
L23:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60))) = uint8(v61)
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)) = uint8(v63)
	v76 = v60 + int32(2)
	goto L19
L24:
	;
	v53 = int32(157)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v53)
	v56 = v46 + int32(83)
	*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)) = uint8(v56)
	v60 = v24 + int32(2)
	goto L23
L25:
	;
	v51 = int32(150)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v51)
	v60 = v45
	goto L23
L26:
	;
	v49 = int32(149)
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v49)
	v60 = v45
	goto L23
L27:
	;
	if v9 != 0 {
		v99 = v23
		v100 = v24
		goto L3
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v31)
	v85 = int32(1)
	v91 = v24 + v85
	v92 = v25 - v85
	v96 = v23 + v85
	goto L9
L30:
	;
	F_report_invalid_encoding(m, int32(4), v23, v25)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
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
func F_euc_tw_to_utf8(m *base.Module, l0 int32) int32 {
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
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(4), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4389780), v18, v18, v18, int32(4), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
	}
}
