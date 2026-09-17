package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecSeqScanWithQualProject(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 float64
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+68))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	F_MemoryContextReset(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v13|v14 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v11 + int32(16)
	return v109
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[0]))
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	goto L12
L7:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v28 = F_SeqNext(m, l0)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	goto L9
L11:
	;
	v109 = v28
	goto L3
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[0]))
	if v39 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v42 = F_SeqNext(m, l0)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v42
	if v14 != 0 {
		goto L27
	} else {
		goto L28
	}
L19:
	;
	if v42 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+4)))
	if v44&int32(2) == int32(0) {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	if v13 == int32(0) {
		v109 = v42
		goto L3
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	m.T0[v53].(func(*base.Module, int32))(m, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v109 = v51
	goto L3
L26:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v100 != 0 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v57 = int32(_a_F_ExecSeqScanWithQualProject_0)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[1]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[1])) = v60
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
	v65 = m.T0[v64].(func(*base.Module, int32, int32, int32) int32)(m, v14, v15, v11+int32(15))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	if v13 == int32(0) {
		v109 = v42
		goto L3
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[1])) = v58
	if v65 == int32(0) {
		goto L26
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)+8))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	m.T0[v77].(func(*base.Module, int32))(m, v75)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v80 = int32(_a_F_ExecSeqScanWithQualProject_0)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[1]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[1])) = v83
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v89 = m.T0[v88].(func(*base.Module, int32, int32, int32) int32)(m, v13+int32(4), v74, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_ExecSeqScanWithQualProject[1])) = v81
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)))
	v95 = v93 & int32(_a_F_ExecSeqScanWithQualProject_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+4)) = uint16(v95)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+6)) = uint16(v98)
	v109 = v75
	goto L3
L35:
	;
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v100)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v100)+240)) = base.F64_add(v101, float64(1))
	goto L37
L36:
	;
	goto L37
L37:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	F_MemoryContextReset(m, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L12
}
