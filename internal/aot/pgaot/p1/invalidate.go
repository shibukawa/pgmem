package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InvalidateBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v48 int64
	_ = v48
	var v59 int64
	_ = v59
	var v70 int64
	_ = v70
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v92 int64
	_ = v92
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v142 int64
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int64
	_ = v162
	var v169 int64
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v205 int64
	_ = v205
	var v209 int64
	_ = v209
	var v218 int64
	_ = v218
	var v223 int64
	_ = v223
	var v239 int32
	_ = v239
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v14
	v16 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v18
	v22 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
	v23 = F_BufTableHashCode(m, v12)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[0]))
	v33 = v26 + v23&int32(127)<<(uint(int32(7))%32) + int32(_a_F_InvalidateBuffer_0)
	goto L5
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L62
	}
L4:
	;
	F_LWLockRelease(m, v33)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L61
	}
L5:
	;
	v44 = F_LWLockAcquire(m, v33, int32(0))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v205 = v142 | int64(4194304)
	v209 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v205, v142&int64(-17179869184))
	if v209 != v205 {
		goto L53
	} else {
		goto L54
	}
L7:
	;
	v46 = int64(4194304)
	v48 = base.AtomicRmwOr64(m, l0, int32(24), v46)
	if v48&v46 != int64(0) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v59 = v48
	goto L11
L9:
	;
	v142 = v48
	goto L10
L10:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v145 != v146 {
		goto L33
	} else {
		goto L34
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = int32(_a_F_InvalidateBuffer_1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = int32(_a_F_InvalidateBuffer_2)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = int32(_a_F_InvalidateBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(0)
	v70 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v70
	if v59&int64(4194304) != v70 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v142 = v131
	goto L10
L13:
	;
	goto L16
L14:
	;
	goto L15
L15:
	;
	v109 = int32(_a_F_InvalidateBuffer_4)
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[1]))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(24))+8))
	if v112 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L16:
	;
	F_perform_spin_delay(m, v12+int32(24))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	goto L15
L18:
	;
	v89 = int64(0)
	v92 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v89, v89)
	if v92&int64(4194304) != v89 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v129 = int64(4194304)
	v131 = base.AtomicRmwOr64(m, l0, int32(24), v129)
	if v131&v129 != int64(0) {
		v59 = v131
		goto L11
	} else {
		goto L31
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[1])) = v127
	goto L21
L23:
	;
	if int32(999) < v110 {
		goto L21
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	if v110 < int32(11) {
		goto L21
	} else {
		goto L30
	}
L26:
	;
	v117 = int32(900)
	if v117 <= v110 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v120 = v117
	goto L29
L28:
	;
	v120 = v110
	goto L29
L29:
	;
	v127 = v120 + int32(100)
	goto L22
L30:
	;
	v127 = v110 - int32(1)
	goto L22
L31:
	;
	goto L12
L32:
	;
	if v142&int64(262143) != int64(0) {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	v162 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
	goto L4
L34:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v148 != v149 {
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v151 != v152 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v154 != v155 {
		goto L33
	} else {
		goto L37
	}
L37:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v157 == v158 {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	goto L33
L39:
	;
	v169 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
	F_LWLockRelease(m, v33)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	goto L6
L42:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v174 = v172 + int32(1)
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateBuffer[2]))
	if v176 != int32(-1) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	F_WaitIO(m, l0)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L52
	}
L44:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)+8))
	if int32(0) < v192 {
		goto L3
	} else {
		goto L51
	}
L45:
	;
	v180 = v176 << (uint(int32(4)) % 32)
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180)+uint32(_c_F_InvalidateBuffer[3])))
	if v183 == v174 {
		v191 = v180 + int32(_a_F_InvalidateBuffer_5)
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v187 = F_GetPrivateRefCountEntrySlow(m, v174, int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	if v187 == int32(0) {
		goto L43
	} else {
		goto L50
	}
L50:
	;
	v191 = v187
	goto L44
L51:
	;
	goto L43
L52:
	;
	goto L5
L53:
	;
	v218 = v209
	goto L56
L54:
	;
	goto L55
L55:
	;
	if v142&int64(33554432) == int64(0) {
		goto L4
	} else {
		goto L59
	}
L56:
	;
	v223 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v218, v218&int64(-17179607041))
	if v218 != v223 {
		v218 = v223
		goto L56
	} else {
		goto L58
	}
L57:
	;
	goto L55
L58:
	;
	goto L57
L59:
	;
	F_BufTableDelete(m, v12, v23)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L4
L61:
	;
	m.G0 = v12 + int32(48)
	return
L62:
	;
	F_errmsg_internal(m, int32(_a_F_InvalidateBuffer_6), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_errfinish(m, int32(_a_F_InvalidateBuffer_3), int32(2426), int32(_a_F_InvalidateBuffer_7))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_InvalidateTSCacheCallBack(m *base.Module, l0 int64, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v10 = v7 + int32(12)
	v11 = base.I32_wrap_i64(l0)
	F_hash_seq_init(m, v10, v11)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v14 = F_hash_seq_search(m, v10)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if v14 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v18 = v14
	goto L7
L5:
	;
	goto L6
L6:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_InvalidateTSCacheCallBack[0]))
	if v11 == v31 {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	v20 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)) = uint8(v20)
	v24 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	goto L6
L9:
	;
	if v24 != 0 {
		v18 = v24
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
	m.G0 = v7 + int32(32)
	return
}
