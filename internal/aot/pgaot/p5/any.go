package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_isAnyTempNamespace(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	v3 = F_get_namespace_name(m, l0)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v3 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v11 = int32(_a_F_isAnyTempNamespace_0)
	goto L8
L6:
	;
	if v49-v50 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L8:
	;
	goto L9
L9:
	;
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v18 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v19 = v3
	v20 = v11
	v21 = int32(8)
	v22 = v18
	goto L14
L11:
	;
	v45 = v11
	v49 = int32(0)
	goto L12
L12:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	goto L6
L13:
	;
	v45 = v40
	v49 = v42
	goto L12
L14:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if base.B2i32(v22 != v24)|base.B2i32(v24 == int32(0)) != 0 {
		v40 = v20
		v42 = v22
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v40 = v34
	v42 = int32(0)
	goto L13
L16:
	;
	v30 = v21 - int32(1)
	if v30 == int32(0) {
		v40 = v20
		v42 = v22
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v33 = int32(1)
	v34 = v20 + v33
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v35 != 0 {
		v19 = v19 + v33
		v20 = v34
		v21 = v30
		v22 = v35
		goto L14
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	F_pfree(m, v3)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	v64 = int32(_a_F_isAnyTempNamespace_1)
	goto L25
L22:
	;
	return int32(1)
L23:
	;
	F_pfree(m, v3)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L36
	}
L25:
	;
	goto L26
L26:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v71 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v72 = v3
	v73 = v64
	v74 = int32(14)
	v75 = v71
	goto L31
L28:
	;
	v98 = v64
	v102 = int32(0)
	goto L29
L29:
	;
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	goto L23
L30:
	;
	v98 = v93
	v102 = v95
	goto L29
L31:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	if base.B2i32(v75 != v77)|base.B2i32(v77 == int32(0)) != 0 {
		v93 = v73
		v95 = v75
		goto L30
	} else {
		goto L33
	}
L32:
	;
	v93 = v87
	v95 = int32(0)
	goto L30
L33:
	;
	v83 = v74 - int32(1)
	if v83 == int32(0) {
		v93 = v73
		v95 = v75
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v86 = int32(1)
	v87 = v73 + v86
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72)+1)))
	if v88 != 0 {
		v72 = v72 + v86
		v73 = v87
		v74 = v83
		v75 = v88
		goto L31
	} else {
		goto L35
	}
L35:
	;
	goto L32
L36:
	;
	return base.B2i32(v102-v103 == int32(0))
}
func F_read_any_attr(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v11 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10))))
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L18
	} else {
		goto L30
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L18
	} else {
		goto L24
	}
L3:
	;
	if base.Ui32(int32(26)) <= base.Ui32((v11&int32(223)-int32(65))&int32(255)) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L6:
	;
	if l1 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v11)
	goto L9
L8:
	;
	goto L9
L9:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v21 != int32(61) {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v25 = v10 + int32(2)
	v27 = v25
	goto L12
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42
	m.G0 = v8 + int32(32)
	return v25
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if v31 == int32(0) {
		v42 = v27
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v27))) = uint8(v38)
	v42 = v27 + int32(1)
	goto L11
L14:
	;
	if v31 != int32(44) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v27 = v27 + int32(1)
	goto L12
L16:
	;
	goto L17
L17:
	;
	goto L13
L18:
	;
	return int32(0)
L19:
	;
	F_errcode(m, int32(16908800))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	F_errmsg(m, int32(_a_F_read_any_attr_0), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errdetail(m, int32(_a_F_read_any_attr_1), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(_a_F_read_any_attr_2), int32(865), int32(_a_F_read_any_attr_3))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
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
	F_errcode(m, int32(16908800))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	F_errmsg(m, int32(_a_F_read_any_attr_0), int32(0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	F_sanitize_char_2(m, v11)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(_a_F_read_any_attr_4)
	F_errdetail(m, int32(_a_F_read_any_attr_5), v8+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(_a_F_read_any_attr_2), int32(879), int32(_a_F_read_any_attr_3))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L18
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
	F_errcode(m, int32(16908800))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	F_errmsg(m, int32(_a_F_read_any_attr_0), int32(0))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L18
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11
	F_errdetail(m, int32(_a_F_read_any_attr_6), v8)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_read_any_attr_2), int32(888), int32(_a_F_read_any_attr_3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
