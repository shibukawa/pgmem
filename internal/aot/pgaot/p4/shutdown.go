package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ShutdownXLOG(m *base.Module, l0 int32, l1 int32) {
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
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
	v26 = int32(0)
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v26 < v28 {
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
	v32 = v26
	goto L14
L12:
	;
	goto L13
L13:
	;
	v68 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v70 <= v68 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[4]))
	v39 = v36 + v32*int32(96)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+164)) = int32(1)
	if v40 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	F_s_lock(m, v39+int32(164), int32(_a_F_ShutdownXLOG_4), int32(3805), int32(_a_F_ShutdownXLOG_5))
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
	v51 = v39 + int32(88)
	*(*int32)(unsafe.Add(mBase, uint32(v51)+76)) = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51)))
	if v54 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L18
L20:
	;
	v57 = F_SendProcSignal(m, v54, int32(3), int32(-1))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L4
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v60 = v32 + int32(1)
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v60 < v62 {
		v32 = v60
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
	v122 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ShutdownXLOG[5])))
	if v122 != int32(1) {
		goto L42
	} else {
		goto L43
	}
L26:
	;
	v74 = v68
	goto L27
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[4]))
	v81 = v78 + v74*int32(96)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+164))
	*(*int32)(unsafe.Add(mBase, uint32(v81)+164)) = int32(1)
	v86 = v81 + int32(164)
	if v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L25
L29:
	;
	F_s_lock(m, v86, int32(_a_F_ShutdownXLOG_4), int32(3833), int32(_a_F_ShutdownXLOG_6))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v93 = v81 + int32(88)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)))
	if v94 == int32(0) {
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
	v111 = m.ExcPending
	if v111 != 0 {
		goto L4
	} else {
		goto L40
	}
L34:
	;
	v105 = v74 + int32(1)
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v105 < v107 {
		v74 = v105
		goto L27
	} else {
		goto L39
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86))) = int32(0)
	goto L34
L36:
	;
	goto L37
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+76)) = int32(0)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v101 != int32(4) {
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
	v112 = int32(0)
	v114 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[3]))
	if v112 < v114 {
		v74 = v112
		goto L27
	} else {
		goto L41
	}
L41:
	;
	goto L28
L42:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[6]))
	if int32(0) < v139 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_ShutdownXLOG[7]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+316))
	v129 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _c_F_ShutdownXLOG[5])) = uint8(base.B2i32(v128 != v129))
	if v128 == v129 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v135 = F_CreateRestartPoint(m, int32(5))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	return
L46:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	v149 = F_CreateCheckPoint(m, int32(5))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L51
	}
L49:
	;
	v146 = F_XLogInsert(m, int32(0), int32(64))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	return
}
