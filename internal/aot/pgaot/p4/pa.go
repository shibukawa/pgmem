package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_pa_find_worker(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	v2 = int32(0)
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = l0
	if l0 == v2 {
		v32 = v2
		m.G0 = v6 + int32(16)
		return v32
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_pa_find_worker[0]))
		if v12 == int32(0) {
			v32 = v2
			m.G0 = v6 + int32(16)
			return v32
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_pa_find_worker[1]))
			if v16 != 0 {
				v32 = v16
				m.G0 = v6 + int32(16)
				return v32
			} else {
				v22 = F_hash_search(m, v12, v6+int32(12), int32(0), v6+int32(11))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+11)))
					if v27 != int32(1) {
						v32 = int32(0)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
						v32 = v30
					}
					m.G0 = v6 + int32(16)
					return v32
				}
			}
		}
	}
}
func F_pa_lock_stream(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_pa_lock_stream[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+32))
	F_LockApplyTransactionForSession(m, v4, l0, int32(0), int32(8))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
func F_pa_xact_finish(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
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
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
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
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	F_UnlockApplyTransactionForSession(m, v8, v10, int32(0), int32(8))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	goto L3
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = base.AtomicRmwXchg32(m, v20, int32(0), int32(1))
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[0]))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+32))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v59 = int32(1)
	F_LockApplyTransactionForSession(m, v56, v58, v59, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L16
	}
L5:
	;
	F_s_lock(m, v20, int32(_a_F_pa_xact_finish_0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v28 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v20))), uint32(v28))
	if v27 == v28 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[1]))
	v38 = F_WaitLatch(m, v34, int32(41), int32(10), int32(134217758))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	goto L4
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[1]))
	v42 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v42
	v47 = base.AtomicRmwOr32(m, v42, int32(_a_F_pa_xact_finish_1), v42)
	goto L13
L13:
	;
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[2]))
	if v49 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[0]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+32))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v68 = int32(1)
	F_UnlockApplyTransactionForSession(m, v65, v67, v68, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v75 = base.AtomicRmwXchg32(m, v72, int32(0), int32(1))
	if v75 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_s_lock(m, v72, int32(_a_F_pa_xact_finish_0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v72)+8))
	v80 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v72))), uint32(v80))
	if v79 == int32(2) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L20
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L59
	}
L23:
	;
	if l1 != int64(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L1
	} else {
		goto L55
	}
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v87)+24))
	F_store_flush_position(m, l1, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[3]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v98 = F_hash_search(m, v92, v93+int32(4), int32(2), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v98 == int32(0) {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v102 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v165 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v165)
	return
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[4]))
	if v106 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v119 = base.AtomicRmwXchg32(m, v116, int32(0), int32(1))
	if v119 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	v109 = v107
	goto L38
L37:
	;
	v109 = int32(0)
	goto L38
L38:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[5]))
	v113 = base.I32_div_s(v111, int32(2))
	if v109 <= v113 {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	F_s_lock(m, v116, int32(_a_F_pa_xact_finish_0))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+16))
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+12)))
	v126 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v123))), uint32(v126))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v129 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_shm_mq_detach(m, v129)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[6]))
	v139 = F_LWLockAcquire(m, v135+int32(_a_F_pa_xact_finish_2), int32(1))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L48
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = int32(0)
	goto L46
L48:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[7]))
	v145 = v142 + v124<<(uint(int32(7))%32)
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+34)))
	if v146 != v125 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[6]))
	F_LWLockRelease(m, v158+int32(_a_F_pa_xact_finish_2))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v149 = v145 + int32(16)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)+20))
	if v150 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	F_logicalrep_worker_stop_internal(m, v149, int32(12))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L49
L53:
	;
	F_pa_free_worker_info(m, l0)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	return
L55:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_pa_xact_finish_3), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_pa_xact_finish_4), int32(1319), int32(_a_F_pa_xact_finish_5))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L59:
	;
	F_errmsg_internal(m, int32(_a_F_pa_xact_finish_6), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_pa_xact_finish_4), int32(567), int32(_a_F_pa_xact_finish_7))
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
