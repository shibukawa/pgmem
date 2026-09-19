package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
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
	var v20 int32
	_ = v20
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
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v65 int64
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v135 int32
	_ = v135
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
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
	v122 = base.AtomicRmwXchg32(m, v11, int32(0), int32(1))
	if v122 != 0 {
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
	v20 = base.AtomicRmwXchg32(m, v11, int32(0), int32(1))
	if v20 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	F_s_lock(m, v22, int32(_a_F_LogicalConfirmReceivedLocation_0), int32(1834), int32(_a_F_LogicalConfirmReceivedLocation_1))
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
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
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
	v73 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29))), uint32(v73))
	F_ReplicationSlotMarkDirty(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L10
	} else {
		goto L24
	}
L16:
	;
	v63 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+248)) = v63
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v29)+256))
	*(*int64)(unsafe.Add(mBase, uint32(v29)+256)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v29)+104)) = v65
	v71 = v62
	v72 = int32(1)
	goto L15
L17:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v29)+240)) = int64(0)
	v53 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v29)+236)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v29)+100)) = v37
	v56 = int32(1)
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v29)+248))
	if base.Ui64(l0) <= base.Ui64(v58-int64(1)) {
		v71 = v56
		v72 = v53
		goto L15
	} else {
		goto L23
	}
L18:
	;
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v29)+248))
	if base.Ui64(v44-int64(1)) < base.Ui64(l0) {
		v62 = int32(0)
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
	v48 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v29))), uint32(v48))
	goto L1
L23:
	;
	v62 = v56
	goto L16
L24:
	;
	F_ReplicationSlotSave(m)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L10
	} else {
		goto L25
	}
L25:
	;
	v82 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L10
	} else {
		goto L26
	}
L26:
	;
	if v82 != 0 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v72
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v71
	F_errmsg_internal(m, int32(_a_F_LogicalConfirmReceivedLocation_2), v8)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L10
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v71 == int32(0) {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	F_errfinish(m, int32(_a_F_LogicalConfirmReceivedLocation_0), int32(1906), int32(_a_F_LogicalConfirmReceivedLocation_1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	v100 = base.AtomicRmwXchg32(m, v97, int32(0), int32(1))
	if v100 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	F_s_lock(m, v102, int32(_a_F_LogicalConfirmReceivedLocation_0), int32(1917), int32(_a_F_LogicalConfirmReceivedLocation_1))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L10
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+100))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v110
	v112 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v109))), uint32(v112))
	F_ReplicationSlotsComputeRequiredXmin(m, v112)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
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
	v119 = m.ExcPending
	if v119 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	goto L1
L39:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	F_s_lock(m, v124, int32(_a_F_LogicalConfirmReceivedLocation_0), int32(1927), int32(_a_F_LogicalConfirmReceivedLocation_1))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L10
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _c_F_LogicalConfirmReceivedLocation[0]))
	v132 = *(*int64)(unsafe.Add(mBase, uint32(v131)+120))
	if base.Ui64(v132) < base.Ui64(l0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L41
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v131)+120)) = l0
	goto L45
L44:
	;
	goto L45
L45:
	;
	v135 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v131))), uint32(v135))
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
