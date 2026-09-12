package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_dsm_postmaster_shutdown(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v3
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v3
	v18 = *(*int32)(unsafe.Add(mBase, _consts[766]))
	if base.Ui32(v18) < base.Ui32(int32(12)) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v9 + int32(48)
	return
L2:
	;
	v130 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L15
	} else {
		goto L31
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	if v23 != int32(-1706017486) {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if base.Ui64(base.I64_extend_i32_u(v18)) < base.Ui64(base.I64_extend_i32_u(v27)*int64(24)+int64(12)) {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if base.Ui32(v27) < base.Ui32(v34) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v40 = int32(0)
	v41 = v22
	goto L10
L8:
	;
	goto L9
L9:
	;
	v93 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L15
	} else {
		goto L24
	}
L10:
	;
	v45 = v41 + v40*int32(24)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v46 == int32(0) {
		v81 = v41
		goto L12
	} else {
		goto L13
	}
L11:
	;
	goto L9
L12:
	;
	v83 = v40 + int32(1)
	if v83 != v34 {
		v40 = v83
		v41 = v81
		goto L10
	} else {
		goto L23
	}
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	if v49&int32(1) != 0 {
		v81 = v41
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v54 = F_errstart(m, int32(13), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return
L16:
	;
	if v54 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v49
	F_errmsg_internal(m, int32(56762), v9+int32(16))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v76 = F_dsm_impl_op(m, int32(3), v49, int32(0), v9+int32(36), v9+int32(40), v9+int32(32), int32(15))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L15
	} else {
		goto L22
	}
L20:
	;
	F_errfinish(m, int32(496919), int32(398), int32(243558))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L15
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	v81 = v79
	goto L12
L23:
	;
	goto L11
L24:
	;
	if v93 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _consts[768]))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v96
	F_errmsg_internal(m, int32(56816), v9)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L15
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[767]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v107
	v111 = *(*int32)(unsafe.Add(mBase, _consts[768]))
	v118 = F_dsm_impl_op(m, int32(3), v111, int32(0), int32(4431304), v9+int32(44), int32(4431308), int32(15))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L15
	} else {
		goto L30
	}
L28:
	;
	F_errfinish(m, int32(496919), int32(408), int32(243558))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L15
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v9)+44))
	*(*int32)(unsafe.Add(mBase, _consts[767])) = v121
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = int32(0)
	goto L1
L31:
	;
	if v130 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_errmsg(m, int32(83387), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L15
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(496919), int32(379), int32(243558))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L15
	} else {
		goto L34
	}
L34:
	;
	goto L1
}
