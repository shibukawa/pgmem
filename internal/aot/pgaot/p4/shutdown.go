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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
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
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[283]))
	*(*int32)(unsafe.Add(mBase, _consts[173])) = v7
	v12 = int32(*(*uint8)(unsafe.Add(mBase, _consts[284])))
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
	F_errmsg(m, int32(240916), int32(0))
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
	v28 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	if v26 < v28 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	F_errfinish(m, int32(490437), int32(6655), int32(527927))
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
	v70 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	if v70 <= v68 {
		goto L25
	} else {
		goto L26
	}
L14:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[285]))
	v39 = v36 + v32*int32(96)
	v41 = v39 + int32(164)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = int32(1)
	v46 = v39 + int32(88)
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
	F_s_lock(m, v41, int32(487926), int32(3805), int32(326154))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+76)) = int32(0)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
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
	v62 = *(*int32)(unsafe.Add(mBase, _consts[273]))
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
	v122 = int32(*(*uint8)(unsafe.Add(mBase, _consts[189])))
	if v122 != int32(1) {
		goto L39
	} else {
		goto L40
	}
L26:
	;
	v74 = v68
	goto L27
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _consts[285]))
	v81 = v78 + v74*int32(96)
	v83 = v81 + int32(164)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(1)
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
	F_s_lock(m, v83, int32(487926), int32(3833), int32(326173))
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
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L31
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(0)
	v113 = v74 + int32(1)
	v115 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	if v113 < v115 {
		v74 = v113
		goto L27
	} else {
		goto L38
	}
L34:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v97 == int32(4) {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v100 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v100
	F_pg_usleep(m, int32(10000))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _consts[273]))
	if int32(0) < v107 {
		v74 = v100
		goto L27
	} else {
		goto L37
	}
L37:
	;
	goto L25
L38:
	;
	goto L28
L39:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[286]))
	if int32(0) < v139 {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _consts[190]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+316))
	v129 = int32(2)
	*(*uint8)(unsafe.Add(mBase, _consts[189])) = uint8(base.B2i32(v128 != v129))
	if v128 == v129 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v135 = F_CreateRestartPoint(m, int32(5))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	return
L43:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v149 = F_CreateCheckPoint(m, int32(5))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L48
	}
L46:
	;
	v146 = F_XLogInsert(m, int32(0), int32(64))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L4
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	return
}
