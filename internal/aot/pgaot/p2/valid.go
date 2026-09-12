package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_check_valid_internal_signature(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	if l0 != int32(2281) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	goto L3
L3:
	;
	v8 = int32(0)
	if l2 <= v8 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = F_pstrdup(m, int32(622789))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L12
	} else {
		goto L13
	}
L5:
	;
	v11 = v8
	goto L6
L6:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1+v11<<(uint(int32(2))%32))))
	if v17 != int32(2281) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v21 = v11 + int32(1)
	if l2 != v21 {
		v11 = v21
		goto L6
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	goto L7
L11:
	;
	goto L4
L12:
	;
	return int32(0)
L13:
	;
	return v29
}
func F_check_valid_version_name(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	v8 = F_strlen(m, l0)
	mBase = m.M
	if v8 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L19
	} else {
		goto L35
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L30
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L19
	} else {
		goto L25
	}
L4:
	;
	v10 = F_strstr(m, l0, int32(670019))
	mBase = m.M
	if v10 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v11 == int32(45) {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v8-int32(1)))))
	if v17 == int32(45) {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v21 = l0
	goto L11
L10:
	;
	if v31 != 0 {
		goto L1
	} else {
		goto L18
	}
L11:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v23 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L10
L13:
	;
	goto L12
L14:
	;
	v31 = int32(0)
	goto L13
L15:
	;
	goto L16
L16:
	;
	if v23 == int32(47) {
		v31 = v21
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v21 = v21 + int32(1)
	goto L11
L18:
	;
	m.G0 = v6 - int32(-64)
	return
L19:
	;
	return
L20:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	F_errmsg(m, int32(727079), v6)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errdetail(m, int32(570991), int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(496595), int32(419), int32(379133))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L19
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L19
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+48)) = l0
	F_errmsg(m, int32(727079), v4+int32(-16))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	F_errdetail(m, int32(669682), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(496595), int32(428), int32(379133))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L19
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
	F_errcode(m, int32(50856066))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L19
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = l0
	F_errmsg(m, int32(727079), v4+int32(-48))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L19
	} else {
		goto L32
	}
L32:
	;
	F_errdetail(m, int32(669758), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L19
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(496595), int32(437), int32(379133))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L19
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L19
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6)+32)) = l0
	F_errmsg(m, int32(727079), v4+int32(-32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L19
	} else {
		goto L37
	}
L37:
	;
	F_errdetail(m, int32(588052), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(496595), int32(447), int32(379133))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L19
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
