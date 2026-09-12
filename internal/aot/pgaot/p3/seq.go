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
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 float64
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
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
	return v110
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
	v110 = v28
	goto L3
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _consts[8]))
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
		v110 = v42
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
	v110 = v51
	goto L3
L26:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v101 != 0 {
		goto L35
	} else {
		goto L36
	}
L27:
	;
	v57 = int32(4470560)
	v58 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v60
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
		v110 = v42
		goto L3
	} else {
		goto L32
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v58
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
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)+72))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	m.T0[v78].(func(*base.Module, int32))(m, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v81 = int32(4470560)
	v82 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v84
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v90 = m.T0[v89].(func(*base.Module, int32, int32, int32) int32)(m, v13+int32(4), v75, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v82
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)))
	v96 = v94 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+4)) = uint16(v96)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v76)+12))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+6)) = uint16(v99)
	v110 = v76
	goto L3
L35:
	;
	v102 = *(*float64)(unsafe.Add(mBase, uint32(v101)+240))
	*(*float64)(unsafe.Add(mBase, uint32(v101)+240)) = base.F64_add(v102, float64(1))
	goto L37
L36:
	;
	goto L37
L37:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v15)+20))
	F_MemoryContextReset(m, v106)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L12
}
