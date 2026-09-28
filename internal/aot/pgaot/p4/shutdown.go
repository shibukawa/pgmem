package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_ShutdownXLOG(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
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
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[1])) = v7
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ShutdownXLOG[2])))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int32(15)
	goto L3
L2:
	;
	v13 = int32(18)
	goto L3
L3:
	;
	v15 = F_errstart(m, v13, int32(0))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return
L5:
	;
	if v15 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	F_errmsg(m, int32(_a_F_ShutdownXLOG_0), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if int32(0) < v27 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_errfinish(m, int32(_a_F_ShutdownXLOG_1), int32(_a_F_ShutdownXLOG_2), int32(_a_F_ShutdownXLOG_3))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	v32 = int32(0)
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
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[4]))
	v38 = v35 + v32*int32(96)
	v40 = v38 + int32(88)
	v43 = base.AtomicRmwXchg32(m, v38, int32(164), int32(1))
	if v43 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	F_s_lock(m, v38+int32(164), int32(_a_F_ShutdownXLOG_4))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v50 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v40)+76)), uint32(v50))
	if v49 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v55 = F_SendProcSignal(m, v49, int32(3), int32(-1))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v58 = v32 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v58 < v60 {
		v32 = v58
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
	v73 = v66
	goto L27
L27:
	;
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[4]))
	v79 = v76 + v73*int32(96)
	v80 = int32(164)
	v81 = v79 + v80
	v84 = base.AtomicRmwXchg32(m, v79, v80, int32(1))
	if v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L25
L29:
	;
	F_s_lock(m, v81, int32(_a_F_ShutdownXLOG_4))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v89 = v79 + int32(88)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v90 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	goto L31
L33:
	;
	F_pg_usleep(m, int32(_a_F_ShutdownXLOG_5))
	mBase = m.M
	v109 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v109 < v111 {
		v73 = v109
		goto L27
	} else {
		goto L40
	}
L34:
	;
	v103 = v73 + int32(1)
	v105 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v103 < v105 {
		v73 = v103
		goto L27
	} else {
		goto L39
	}
L35:
	;
	v93 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v81))), uint32(v93))
	goto L34
L36:
	;
	goto L37
L37:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v97 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v89)+76)), uint32(v97))
	if v96 != int32(4) {
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
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)+308))
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
