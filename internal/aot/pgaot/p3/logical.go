package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogicalConfirmReceivedLocation(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int64
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)+240))
	if v12 == int64(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v8 + int32(16)
	return
L2:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
	if v117 != 0 {
		goto L39
	} else {
		goto L40
	}
L3:
	;
	v15 = *(*int64)(unsafe.Add(mBase, uint32(v11)+248))
	if v15 == int64(0) {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(1)
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	F_s_lock(m, v22, int32(475923), int32(1834), int32(253684))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	v30 = *(*int64)(unsafe.Add(mBase, uint32(v29)+120))
	if base.Ui64(v30) < base.Ui64(l0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	goto L9
L12:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+120)) = l0
	goto L14
L13:
	;
	goto L14
L14:
	;
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v29)+240))
	if base.Ui64(l0) <= base.Ui64(v33-int64(1)) {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L24
	}
L16:
	;
	v62 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+248)) = v62
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v29)+256))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+256)) = v62
	*(*int64)(unsafe.Add(mBase, uint32(v29)+104)) = v64
	v70 = v61
	v71 = int32(1)
	goto L15
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = int64(0)
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v37
	v55 = int32(1)
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v29)+248))
	if base.Ui64(l0) <= base.Ui64(v57-int64(1)) {
		v70 = v55
		v71 = v52
		goto L15
	} else {
		goto L23
	}
L18:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v29)+248))
	if base.Ui64(v44-int64(1)) < base.Ui64(l0) {
		v61 = int32(0)
		goto L16
	} else {
		goto L22
	}
L19:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+236))
	if v37 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+100))
	if v40 != v37 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = int32(0)
	goto L1
L23:
	;
	v61 = v55
	goto L16
L24:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v80 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	if v80 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v71
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v70
	F_errmsg_internal(m, int32(55027), v8)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L10
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v70 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	F_errfinish(m, int32(475923), int32(1906), int32(253684))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = int32(1)
	if v96 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	F_s_lock(m, v100, int32(475923), int32(1917), int32(253684))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L10
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	v108 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = v108
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v107)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = v110
	F_ReplicationSlotsComputeRequiredXmin(m, v108)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L10
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	F_ReplicationSlotsComputeRequiredLSN(m)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	goto L1
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	F_s_lock(m, v121, int32(475923), int32(1927), int32(253684))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _consts[664]))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v128)+120))
	if base.Ui64(v129) < base.Ui64(l0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v128)+120)) = l0
	goto L45
L44:
	;
	goto L45
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = int32(0)
	goto L1
}
func F_LogicalTapeSetBlocks(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	return v2 - v3
}
