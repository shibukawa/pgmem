package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ShutdownXLOG(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[1])) = v6
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ShutdownXLOG[2])))
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v12 = int32(15)
	goto L3
L2:
	;
	v12 = int32(18)
	goto L3
L3:
	;
	v14 = F_errstart(m, v12, int32(0))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v14 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_errmsg(m, int32(_a_F_ShutdownXLOG_0), int32(0))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v25 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v25 < v27 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_errfinish(m, int32(_a_F_ShutdownXLOG_1), int32(_a_F_ShutdownXLOG_2), int32(_a_F_ShutdownXLOG_3))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v31 = v25
	goto L14
L12:
	;
	goto L13
L13:
	;
	v66 = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v68 <= v66 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[4]))
	v37 = v34 + v31*int32(96)
	v39 = v37 + int32(88)
	v42 = base.AtomicRmwXchg32(m, v37, int32(164), int32(1))
	if v42 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	F_s_lock(m, v37+int32(164), int32(_a_F_ShutdownXLOG_4), int32(3805), int32(_a_F_ShutdownXLOG_5))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v51 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v39)+76)), uint32(v51))
	if v50 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v56 = F_SendProcSignal(m, v50, int32(3), int32(-1))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v59 = v31 + int32(1)
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v59 < v61 {
		v31 = v59
		goto L14
	} else {
		goto L24
	}
L23:
	;
	goto L22
L24:
	;
	goto L15
L25:
	;
	v119 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ShutdownXLOG[5])))
	if v119 != int32(1) {
		goto L41
	} else {
		goto L42
	}
L26:
	;
	v72 = v66
	goto L27
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[4]))
	v78 = v75 + v72*int32(96)
	v79 = int32(164)
	v80 = v78 + v79
	v83 = base.AtomicRmwXchg32(m, v78, v79, int32(1))
	if v83 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L25
L29:
	;
	F_s_lock(m, v80, int32(_a_F_ShutdownXLOG_4), int32(3833), int32(_a_F_ShutdownXLOG_6))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v90 = v78 + int32(88)
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)))
	if v91 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L31
L33:
	;
	F_pg_usleep(m, int32(_a_F_ShutdownXLOG_7))
	mBase = m.M
	v110 = int32(0)
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v110 < v112 {
		v72 = v110
		goto L27
	} else {
		goto L40
	}
L34:
	;
	v104 = v72 + int32(1)
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v104 < v106 {
		v72 = v104
		goto L27
	} else {
		goto L39
	}
L35:
	;
	v94 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v80))), uint32(v94))
	goto L34
L36:
	;
	goto L37
L37:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v98 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v90)+76)), uint32(v98))
	if v97 != int32(4) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	goto L25
L40:
	;
	goto L28
L41:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[6]))
	if int32(0) < v136 {
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[7]))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+316))
	v126 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_ShutdownXLOG[5])) = uint8(base.B2i32(v125 != v126))
	if v125 == v126 {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v132 = F_CreateRestartPoint(m, int32(5))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	return
L45:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v146 = F_CreateCheckPoint(m, int32(5))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	v143 = F_XLogInsert(m, int32(0), int32(64))
	mBase = m.M
	v144 = m.ExcPending
	if v144 != 0 {
		goto L4
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	return
}
