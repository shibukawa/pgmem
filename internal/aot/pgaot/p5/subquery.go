package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SubqueryNext(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+116))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+52))
	if v3 != 0 {
		F_ExecReScan(m, v2)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return int32(0)
		} else {
			v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
			v9 = m.T0[v8].(func(*base.Module, int32) int32)(m, v2)
			mBase = m.M
			v10 = m.ExcPending
			if v10 != 0 {
				return int32(0)
			} else {
				return v9
			}
		}
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(v2)+12))
		v9 = m.T0[v8].(func(*base.Module, int32) int32)(m, v2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			return v9
		}
	}
}
func F_find_subquery_safe_quals(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v8 + int32(32)
	return
L2:
	;
	v12 = l0
	goto L6
L3:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v102 = F_lappend(m, v101, v96)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L17
	} else {
		goto L34
	}
L4:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_find_subquery_safe_quals(m, v87, l1)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L17
	} else {
		goto L31
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L17
	} else {
		goto L28
	}
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v17 != int32(64) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L17
	} else {
		goto L25
	}
L8:
	;
	switch v17 - int32(63) {
	case 0:
		goto L1
	default:
		goto L5
	case 2:
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	switch v52 {
	case 0:
		goto L4
	case 1, 4, 5:
		goto L23
	case 2:
		goto L1
	case 3:
		v54 = int32(16)
		goto L22
	default:
		goto L21
	}
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v22 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v50 != 0 {
		v96 = v50
		goto L3
	} else {
		goto L20
	}
L13:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v25 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v31 = int32(0)
	goto L15
L15:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v34+v31<<(uint(int32(2))%32))))
	F_find_subquery_safe_quals(m, v38, l1)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L12
L17:
	;
	return
L18:
	;
	v42 = v31 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v42 < v43 {
		v31 = v42
		goto L15
	} else {
		goto L19
	}
L19:
	;
	goto L16
L20:
	;
	goto L1
L21:
	;
	goto L7
L22:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v12+v54)))
	if v56 != 0 {
		v12 = v56
		goto L6
	} else {
		goto L24
	}
L23:
	;
	v54 = int32(12)
	goto L22
L24:
	;
	goto L1
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v61
	F_errmsg_internal(m, int32(_a_F_find_subquery_safe_quals_0), v8+int32(16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L17
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_find_subquery_safe_quals_1), int32(2345), int32(_a_F_find_subquery_safe_quals_2))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v77
	F_errmsg_internal(m, int32(_a_F_find_subquery_safe_quals_3), v8)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L17
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_find_subquery_safe_quals_1), int32(2351), int32(_a_F_find_subquery_safe_quals_2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L17
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	F_find_subquery_safe_quals(m, v90, l1)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L17
	} else {
		goto L32
	}
L32:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
	if v93 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v96 = v93
	goto L3
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v102
	goto L1
}
