package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_g_int_same(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v18 int32
	_ = v18
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
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v87 int32
	_ = v87
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v21 = F_ArrayGetNItemsSafe(m, v18, v10+int32(16))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v23 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L36
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L32
	}
L7:
	;
	v24 = F_array_contains_nulls(m, v10)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v26 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	if v24 != 0 {
		goto L6
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v27 = F_array_contains_nulls(m, v15)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v31 = base.I32_wrap_i64(v17)
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v35 = F_ArrayGetNItemsSafe(m, v32, v15+int32(16))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L18
	}
L15:
	;
	if v27 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	return v17 & int64(4294967295)
L18:
	;
	if v35 == v21 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v38)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
	if v40 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v87 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31))) = uint8(v87)
	goto L17
L22:
	;
	v48 = v40
	goto L24
L23:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	v48 = (v41<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L24
L24:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	if v50 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v58 = v50
	goto L27
L26:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v58 = (v51<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L27
L27:
	;
	v60 = v21
	v61 = v48 + v10
	v62 = v58 + v15
	goto L28
L28:
	;
	if v60 == int32(0) {
		goto L17
	} else {
		goto L30
	}
L29:
	;
	goto L21
L30:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
	v74 = int32(4)
	if v72 == v73 {
		v60 = v60 - int32(1)
		v61 = v61 + v74
		v62 = v62 + v74
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
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errmsg(m, int32(_a_F_g_int_same_0), int32(0))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_errfinish(m, int32(_a_F_g_int_same_1), int32(396), int32(_a_F_g_int_same_2))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
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
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errmsg(m, int32(_a_F_g_int_same_0), int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F_g_int_same_1), int32(397), int32(_a_F_g_int_same_2))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
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
