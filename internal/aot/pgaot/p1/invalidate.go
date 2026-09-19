package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
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
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v16 & int32(-4194305)
	v20 = F_BufTableHashCode(m, v8)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[0]))
	v30 = v23 + v20&int32(127)<<(uint(int32(7))%32) + int32(_a_F_InvalidateBuffer_0)
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
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
	v37 = F_LWLockAcquire(m, v30, int32(0))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	if v69&int32(33554432) != 0 {
		goto L69
	} else {
		goto L70
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = int32(_a_F_InvalidateBuffer_1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(_a_F_InvalidateBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(_a_F_InvalidateBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
	v49 = int32(_a_F_InvalidateBuffer_4)
	v51 = base.AtomicRmwOr32(m, l0, int32(24), v49)
	if v51&v49 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	goto L11
L9:
	;
	v69 = v51
	goto L10
L10:
	;
	v76 = int32(_a_F_InvalidateBuffer_5)
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[1]))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v8+int32(24))+8))
	if v79 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L11:
	;
	F_perform_spin_delay(m, v8+int32(24))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L13
	}
L12:
	;
	v69 = v65
	goto L10
L13:
	;
	v63 = int32(_a_F_InvalidateBuffer_4)
	v65 = base.AtomicRmwOr32(m, l0, int32(24), v63)
	if v65&v63 != 0 {
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	if v96 != v97 {
		goto L27
	} else {
		goto L28
	}
L16:
	;
	goto L15
L17:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[1])) = v94
	goto L16
L18:
	;
	if int32(999) < v77 {
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	if v77 < int32(11) {
		goto L16
	} else {
		goto L25
	}
L21:
	;
	v84 = int32(900)
	if v84 <= v77 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v87 = v84
	goto L24
L23:
	;
	v87 = v77
	goto L24
L24:
	;
	v94 = v87 + int32(100)
	goto L17
L25:
	;
	v94 = v77 - int32(1)
	goto L17
L26:
	;
	if v69&int32(_a_F_InvalidateBuffer_6) != 0 {
		goto L34
	} else {
		goto L35
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v69 & int32(-4194305)
	F_LWLockRelease(m, v30)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L33
	}
L28:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	if v99 != v100 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	if v102 != v103 {
		goto L27
	} else {
		goto L30
	}
L30:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v105 != v106 {
		goto L27
	} else {
		goto L31
	}
L31:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	if v108 == v109 {
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v69 & int32(-4194305)
	F_LWLockRelease(m, v30)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
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
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v125 = v123 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v125
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[2]))
	if v125 == v128 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	F_WaitIO(m, l0)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L68
	}
L39:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+4))
	if int32(0) < v174 {
		goto L3
	} else {
		goto L67
	}
L40:
	;
	v173 = int32(_a_F_InvalidateBuffer_7)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[3]))
	if v125 == v132 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v173 = int32(_a_F_InvalidateBuffer_8)
	goto L39
L44:
	;
	goto L45
L45:
	;
	v136 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[4]))
	if v125 == v136 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v173 = int32(_a_F_InvalidateBuffer_9)
	goto L39
L47:
	;
	goto L48
L48:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[5]))
	if v125 == v140 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v173 = int32(_a_F_InvalidateBuffer_10)
	goto L39
L50:
	;
	goto L51
L51:
	;
	v144 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[6]))
	if v125 == v144 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v173 = int32(_a_F_InvalidateBuffer_11)
	goto L39
L53:
	;
	goto L54
L54:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[7]))
	if v125 == v148 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v173 = int32(_a_F_InvalidateBuffer_12)
	goto L39
L56:
	;
	goto L57
L57:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[8]))
	if v125 == v152 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v173 = int32(_a_F_InvalidateBuffer_13)
	goto L39
L59:
	;
	goto L60
L60:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[9]))
	if v125 == v156 {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v173 = int32(_a_F_InvalidateBuffer_14)
	goto L39
L62:
	;
	goto L63
L63:
	;
	v160 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[10]))
	if v160 == int32(0) {
		goto L38
	} else {
		goto L64
	}
L64:
	;
	v164 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[11]))
	v167 = int32(0)
	v169 = F_hash_search(m, v164, v8+int32(24), v167, v167)
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v169 == int32(0) {
		goto L38
	} else {
		goto L66
	}
L66:
	;
	v173 = v169
	goto L39
L67:
	;
	goto L38
L68:
	;
	goto L5
L69:
	;
	F_BufTableDelete(m, v8, v20)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	F_LWLockRelease(m, v30)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
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
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L4
L75:
	;
	F_errmsg_internal(m, int32(_a_F_InvalidateBuffer_15), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(_a_F_InvalidateBuffer_3), int32(2236), int32(_a_F_InvalidateBuffer_16))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
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
