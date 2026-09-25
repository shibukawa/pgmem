package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v14 int64
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v227 int32
	_ = v227
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v12
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v20 = base.AtomicRmwOr32(m, v2, int32(_a_F_InvalidateBuffer_0), v2)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v16 & int32(-4194305)
	v24 = F_BufTableHashCode(m, v8)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[0]))
	v34 = v27 + v24&int32(127)<<(uint(int32(7))%32) + int32(_a_F_InvalidateBuffer_1)
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L1
	} else {
		goto L75
	}
L4:
	;
	m.G0 = v8 + int32(48)
	return
L5:
	;
	v41 = F_LWLockAcquire(m, v34, int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v198 = int32(0)
	v201 = base.AtomicRmwOr32(m, v198, int32(_a_F_InvalidateBuffer_0), v198)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v198
	if v73&int32(33554432) != 0 {
		goto L69
	} else {
		goto L70
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(_a_F_InvalidateBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_InvalidateBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_InvalidateBuffer_4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
	v53 = int32(_a_F_InvalidateBuffer_5)
	v55 = base.AtomicRmwOr32(m, l0, int32(24), v53)
	if v55&v53 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L11
L9:
	;
	v73 = v55
	goto L10
L10:
	;
	v80 = int32(_a_F_InvalidateBuffer_6)
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[1]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(24))+8))
	if v83 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	F_perform_spin_delay(m, v8+int32(24))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v73 = v69
	goto L10
L13:
	;
	v67 = int32(_a_F_InvalidateBuffer_5)
	v69 = base.AtomicRmwOr32(m, l0, int32(24), v67)
	if v69&v67 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v100 != v101 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[1])) = v98
	goto L16
L18:
	;
	if int32(999) < v81 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v81 < int32(11) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	v88 = int32(900)
	if v88 <= v81 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v91 = v88
	goto L24
L23:
	;
	v91 = v81
	goto L24
L24:
	;
	v98 = v91 + int32(100)
	goto L17
L25:
	;
	v98 = v81 - int32(1)
	goto L17
L26:
	;
	if v73&int32(_a_F_InvalidateBuffer_7) != 0 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	v115 = int32(0)
	v118 = base.AtomicRmwOr32(m, v115, int32(_a_F_InvalidateBuffer_0), v115)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v73 & int32(-4194305)
	F_LWLockRelease(m, v34)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L33
	}
L28:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v103 != v104 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v106 != v107 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v109 != v110 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v112 == v113 {
		goto L26
	} else {
		goto L32
	}
L32:
	;
	goto L27
L33:
	;
	goto L4
L34:
	;
	v126 = int32(0)
	v129 = base.AtomicRmwOr32(m, v126, int32(_a_F_InvalidateBuffer_0), v126)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v73 & int32(-4194305)
	F_LWLockRelease(m, v34)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	goto L6
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v137 = v135 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v137
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[2]))
	if v137 == v140 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	F_WaitIO(m, l0)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L68
	}
L39:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	if int32(0) < v186 {
		goto L3
	} else {
		goto L67
	}
L40:
	;
	v185 = int32(_a_F_InvalidateBuffer_8)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[3]))
	if v137 == v144 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v185 = int32(_a_F_InvalidateBuffer_9)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[4]))
	if v137 == v148 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v185 = int32(_a_F_InvalidateBuffer_10)
	goto L39
L47:
	;
	goto L48
L48:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[5]))
	if v137 == v152 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v185 = int32(_a_F_InvalidateBuffer_11)
	goto L39
L50:
	;
	goto L51
L51:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[6]))
	if v137 == v156 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v185 = int32(_a_F_InvalidateBuffer_12)
	goto L39
L53:
	;
	goto L54
L54:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[7]))
	if v137 == v160 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v185 = int32(_a_F_InvalidateBuffer_13)
	goto L39
L56:
	;
	goto L57
L57:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[8]))
	if v137 == v164 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v185 = int32(_a_F_InvalidateBuffer_14)
	goto L39
L59:
	;
	goto L60
L60:
	;
	v168 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[9]))
	if v137 == v168 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v185 = int32(_a_F_InvalidateBuffer_15)
	goto L39
L62:
	;
	goto L63
L63:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[10]))
	if v172 == int32(0) {
		goto L38
	} else {
		goto L64
	}
L64:
	;
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[11]))
	v179 = int32(0)
	v181 = F_hash_search(m, v176, v8+int32(24), v179, v179)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v181 == int32(0) {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	v185 = v181
	goto L39
L67:
	;
	goto L38
L68:
	;
	goto L5
L69:
	;
	F_BufTableDelete(m, v8, v24)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_LWLockRelease(m, v34)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_StrategyFreeBuffer(m, l0)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L4
L75:
	;
	F_errmsg_internal(m, int32(_a_F_InvalidateBuffer_16), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_InvalidateBuffer_4), int32(2236), int32(_a_F_InvalidateBuffer_17))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_InvalidateTSCacheCallBack(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	v4 = m.G0
	v6 = v4 - int32(32)
	m.G0 = v6
	v9 = v6 + int32(12)
	F_hash_seq_init(m, v9, l0)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v12 = F_hash_seq_search(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v16 = v12
	goto L7
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateTSCacheCallBack[0]))
	if l0 == v27 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v17 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+4)) = uint8(v17)
	v21 = F_hash_seq_search(m, v6+int32(12))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if v21 != 0 {
		v16 = v21
		goto L7
	} else {
		goto L10
	}
L10:
	;
	goto L8
L11:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateTSCacheCallBack[1])) = int32(0)
	goto L13
L12:
	;
	goto L13
L13:
	;
	m.G0 = v6 + int32(32)
	return
}
