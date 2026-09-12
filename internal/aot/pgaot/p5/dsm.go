package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsm_pin_mapping(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v3 != 0 {
		F_ResourceOwnerForget(m, v3, l0, int32(1629852))
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
			return
		}
	} else {
		return
	}
}
func F_dsm_unpin_segment(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v17 = F_LWLockAcquire(m, v13+int32(4352), v2)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[656]))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v21 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v63 = v27 * int32(24)
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20+v63)+32)))
	if v65 != 0 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v27 = v2
	goto L7
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L14
	}
L7:
	;
	v33 = v20 + int32(12) + v27*int32(24)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if base.Ui32(int32(2)) <= base.Ui32(v34) {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L6
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	if v37 == l0 {
		goto L3
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v40 = v27 + int32(1)
	if v40 != v21 {
		v27 = v40
		goto L7
	} else {
		goto L13
	}
L12:
	;
	goto L11
L13:
	;
	goto L8
L14:
	;
	F_errmsg_internal(m, int32(391238), int32(0))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(498265), int32(1016), int32(95390))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L17:
	;
	v67 = l0 & int32(1)
	if v67 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L36
	}
L20:
	;
	v70 = v20
	goto L22
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[656]))
	v70 = v69
	goto L22
L22:
	;
	v71 = v70 + v63
	v72 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+32)) = uint8(v72)
	v75 = v71 + int32(16)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v78 = v76 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = v78
	v81 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v81+int32(4352))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v78 != int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	m.G0 = v10 + int32(16)
	return
L25:
	;
	v88 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v88
	if v67 == v88 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _consts[656]))
	*(*int32)(unsafe.Add(mBase, uint32(v136+v63)+16)) = int32(0)
	v141 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v141+int32(4352))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L35
	}
L27:
	;
	v105 = F_dsm_impl_op(m, int32(3), l0, int32(0), v10+int32(12), v10+int32(8), v10+int32(4), int32(19))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v121 = F_LWLockAcquire(m, v117+int32(4352), int32(0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L33
	}
L30:
	;
	if v105 == int32(0) {
		goto L24
	} else {
		goto L31
	}
L31:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v114 = F_LWLockAcquire(m, v110+int32(4352), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	goto L26
L33:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[657]))
	v126 = *(*int32)(unsafe.Add(mBase, _consts[656]))
	v129 = v126 + v27*int32(24)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+20))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v129)+24))
	F_FreePageManagerPut(m, v124, v130, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	goto L26
L35:
	;
	goto L24
L36:
	;
	F_errmsg_internal(m, int32(453543), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(498265), int32(1018), int32(95390))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
