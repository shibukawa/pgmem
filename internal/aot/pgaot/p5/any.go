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
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
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
	v11 = int32(508713)
	goto L8
L6:
	;
	if v48-v49 == int32(0) {
		goto L20
	} else {
		goto L21
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
	v44 = v11
	v48 = int32(0)
	goto L12
L12:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	goto L6
L13:
	;
	v44 = v39
	v48 = v41
	goto L12
L14:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v22 != v24 {
		v39 = v20
		v41 = v22
		goto L13
	} else {
		goto L16
	}
L15:
	;
	v39 = v33
	v41 = int32(0)
	goto L13
L16:
	;
	if v24 == int32(0) {
		v39 = v20
		v41 = v22
		goto L13
	} else {
		goto L17
	}
L17:
	;
	v29 = v21 - int32(1)
	if v29 == int32(0) {
		v39 = v20
		v41 = v22
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v32 = int32(1)
	v33 = v20 + v32
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+1)))
	if v34 != 0 {
		v19 = v19 + v32
		v20 = v33
		v21 = v29
		v22 = v34
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	F_pfree(m, v3)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v63 = int32(508698)
	goto L26
L23:
	;
	return int32(1)
L24:
	;
	F_pfree(m, v3)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L1
	} else {
		goto L38
	}
L26:
	;
	goto L27
L27:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3))))
	if v70 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v71 = v3
	v72 = v63
	v73 = int32(14)
	v74 = v70
	goto L32
L29:
	;
	v96 = v63
	v100 = int32(0)
	goto L30
L30:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	goto L24
L31:
	;
	v96 = v91
	v100 = v93
	goto L30
L32:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v72))))
	if v74 != v76 {
		v91 = v72
		v93 = v74
		goto L31
	} else {
		goto L34
	}
L33:
	;
	v91 = v85
	v93 = int32(0)
	goto L31
L34:
	;
	if v76 == int32(0) {
		v91 = v72
		v93 = v74
		goto L31
	} else {
		goto L35
	}
L35:
	;
	v81 = v73 - int32(1)
	if v81 == int32(0) {
		v91 = v72
		v93 = v74
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v84 = int32(1)
	v85 = v72 + v84
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+1)))
	if v86 != 0 {
		v71 = v71 + v84
		v72 = v85
		v73 = v81
		v74 = v86
		goto L32
	} else {
		goto L37
	}
L37:
	;
	goto L33
L38:
	;
	return base.B2i32(v100-v101 == int32(0))
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
	var v28 int32
	_ = v28
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
	v28 = v25
	goto L12
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v42
	m.G0 = v8 + int32(32)
	return v25
L12:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if v31 == int32(0) {
		v42 = v28
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v38 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28))) = uint8(v38)
	v42 = v28 + int32(1)
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
	v28 = v28 + int32(1)
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
	F_errmsg(m, int32(406069), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errdetail(m, int32(627919), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(498384), int32(865), int32(206783))
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
	F_errmsg(m, int32(406069), int32(0))
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(4419024)
	F_errdetail(m, int32(667086), v8+int32(16))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	F_errfinish(m, int32(498384), int32(879), int32(206783))
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
	F_errmsg(m, int32(406069), int32(0))
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
	F_errdetail(m, int32(670517), v8)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(498384), int32(888), int32(206783))
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
