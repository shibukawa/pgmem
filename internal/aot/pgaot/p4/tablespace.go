package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_get_tablespace_maintenance_io_concurrency(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = F_get_tablespace(m, l0)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
		if v6 != 0 {
			v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+28))
			if int32(0) <= v7 {
				v13 = v7
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, _consts[294]))
				v13 = v12
			}
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, _consts[294]))
			v13 = v12
		}
		return v13
	}
}
func F_sendTablespace(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v125 int64
	_ = v125
	v9 = m.G0
	v11 = v9 - int32(1152)
	m.G0 = v11
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = int32(520889)
	v22 = F_pg_snprintf(m, v11+int32(128), int32(1024), int32(165202), v11+int32(16))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v32 = F___fstatat(m, int32(-100), v11+int32(128), v11+int32(32), int32(256))
	mBase = m.M
	goto L4
L3:
	;
	m.G0 = v11 + int32(1152)
	return v125
L4:
	;
	if v32 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[140]))
	if v34 == int32(44) {
		v125 = int64(0)
		goto L3
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F__tarWriteHeader(m, l0, int32(520889), int32(0), v11+int32(32), l3)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v11 + int32(128)
	F_errmsg(m, int32(278218), v11)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(465325), int32(1160), int32(393888))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	if l1&int32(3) == int32(0) {
		v85 = l1
		goto L16
	} else {
		goto L17
	}
L14:
	;
	v121 = F_sendDir(m, l0, v11+int32(128), v118, l3, int32(0), int32(1), l4, l2, l5)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L31
	}
L15:
	;
	v118 = v110 - l1
	goto L14
L16:
	;
	v89 = v85
	goto L25
L17:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v69 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v118 = int32(0)
	goto L14
L19:
	;
	goto L20
L20:
	;
	v74 = l1
	goto L21
L21:
	;
	v78 = v74 + int32(1)
	if v78&int32(3) == int32(0) {
		v85 = v78
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v110 = v78
	goto L15
L23:
	;
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v83 != 0 {
		v74 = v78
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v98 = int32(-2139062144)
	if (int32(16843008)-v95|v95)&v98 == v98 {
		v89 = v89 + int32(4)
		goto L25
	} else {
		goto L27
	}
L26:
	;
	v104 = v89
	goto L28
L27:
	;
	goto L26
L28:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
	if v108 != 0 {
		v104 = v104 + int32(1)
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v110 = v104
	goto L15
L30:
	;
	goto L29
L31:
	;
	v125 = v121 + int64(512)
	goto L3
}
func F_tablespace_reloptions(m *base.Module, l0 int32, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(128), int32(32), int32(706976), int32(4))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
