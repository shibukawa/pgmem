package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_WaitForParallelWorkersToAttach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L10
	} else {
		goto L35
	}
L2:
	;
	v13 = v11
	goto L5
L3:
	;
	goto L4
L4:
	;
	m.G0 = v9 + int32(16)
	return
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v19 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v23 = v13
	goto L9
L9:
	;
	v24 = int32(0)
	if v24 < v23 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	return
L11:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v23 = v22
	goto L9
L12:
	;
	v29 = v24
	goto L15
L13:
	;
	v108 = v23
	goto L14
L14:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v113 < v108 {
		v13 = v108
		goto L5
	} else {
		goto L34
	}
L15:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v34 = v33 + v29
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	if v35 != 0 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v108 = v105
	goto L14
L17:
	;
	v104 = v29 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v104 < v105 {
		v29 = v104
		goto L15
	} else {
		goto L33
	}
L18:
	;
	v37 = v29 << (uint(int32(3)) % 32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v39 = v37 + v38
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	if v40 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v43 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v34))) = uint8(v43)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v45 + v43
	goto L17
L20:
	;
	goto L21
L21:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v52 = F_GetBackgroundWorkerPid(m, v49, v9+int32(12))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L10
	} else {
		goto L25
	}
L22:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _consts[99]))
	v91 = F_WaitLatch(m, v87, int32(33), int32(-1), int32(134217734))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L10
	} else {
		goto L30
	}
L23:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70+v37)+4))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	v74 = F_shm_mq_get_sender(m, v73)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L10
	} else {
		goto L28
	}
L24:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54+v37)+4))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
	v58 = F_shm_mq_get_sender(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L10
	} else {
		goto L26
	}
L25:
	;
	switch v52 {
	case 0:
		goto L24
	default:
		goto L22
	case 2:
		goto L23
	}
L26:
	;
	if v58 == int32(0) {
		goto L17
	} else {
		goto L27
	}
L27:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v64 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v62+v29))) = uint8(v64)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v66 + v64
	goto L17
L28:
	;
	if v74 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v80 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v78+v29))) = uint8(v80)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+60)) = v82 + v80
	goto L17
L30:
	;
	if v91&int32(1) == int32(0) {
		goto L17
	} else {
		goto L31
	}
L31:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _consts[99]))
	*(*int32)(unsafe.Add(mBase, uint32(v98))) = int32(0)
	goto L32
L32:
	;
	goto L17
L33:
	;
	goto L16
L34:
	;
	goto L6
L35:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L10
	} else {
		goto L36
	}
L36:
	;
	F_errmsg(m, int32(342455), int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	F_errhint(m, int32(627073), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L10
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(498636), int32(760), int32(327061))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L10
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
