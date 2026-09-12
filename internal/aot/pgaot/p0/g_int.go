package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_same(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v126 int32
	_ = v126
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v19 = F_ArrayGetNItems(m, v16, v8+int32(16))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v21 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L36
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L32
	}
L7:
	;
	v22 = F_array_contains_nulls(m, v8)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v24 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v22 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v25 = F_array_contains_nulls(m, v13)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v30 = F_ArrayGetNItems(m, v27, v13+int32(16))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	if v25 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	return v15
L18:
	;
	if v30 == v19 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v33)
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v35 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v78 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v78)
	goto L17
L22:
	;
	v43 = v35
	goto L24
L23:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v43 = (v36<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L24
L24:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	if v45 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v53 = v45
	goto L27
L26:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v53 = (v46<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L27
L27:
	;
	v55 = v19
	v56 = v43 + v8
	v57 = v53 + v13
	goto L28
L28:
	;
	if v55 == int32(0) {
		goto L17
	} else {
		goto L30
	}
L29:
	;
	goto L21
L30:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v67 = int32(4)
	if v65 == v66 {
		v55 = v55 - int32(1)
		v56 = v56 + v67
		v57 = v57 + v67
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(152677), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(493034), int32(395), int32(376881))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L36:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(152677), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(493034), int32(396), int32(376881))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
