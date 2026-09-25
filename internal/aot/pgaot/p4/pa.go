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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
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
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+32))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v61 = int32(1)
	F_LockApplyTransactionForSession(m, v58, v60, v61, v61)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L16
	}
L5:
	;
	F_s_lock(m, v20, int32(_a_F_pa_xact_finish_0), int32(1330), int32(_a_F_pa_xact_finish_1))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v30 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v20))), uint32(v30))
	if v29 == v30 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[1]))
	v40 = F_WaitLatch(m, v36, int32(41), int32(10), int32(134217758))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[1]))
	v44 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = v44
	v49 = base.AtomicRmwOr32(m, v44, int32(_a_F_pa_xact_finish_2), v44)
	goto L13
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[2]))
	if v51 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+32))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	v70 = int32(1)
	F_UnlockApplyTransactionForSession(m, v67, v69, v70, v70)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v77 = base.AtomicRmwXchg32(m, v74, int32(0), int32(1))
	if v77 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_s_lock(m, v74, int32(_a_F_pa_xact_finish_0), int32(1330), int32(_a_F_pa_xact_finish_1))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	v84 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v74))), uint32(v84))
	if v83 == int32(2) {
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
	v193 = m.ExcPending
	if v193 != 0 {
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
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L55
	}
L26:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v92 = *(*int64)(unsafe.Add(mBase, uint32(v91)+24))
	F_store_flush_position(m, l1, v92)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v96 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[3]))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v102 = F_hash_search(m, v96, v97+int32(4), int32(2), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v102 == int32(0) {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v106 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	v172 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v172)
	return
L33:
	;
	v110 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[4]))
	if v110 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v123 = base.AtomicRmwXchg32(m, v120, int32(0), int32(1))
	if v123 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v110)+4))
	v113 = v111
	goto L38
L37:
	;
	v113 = int32(0)
	goto L38
L38:
	;
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[5]))
	v117 = base.I32_div_s(v115, int32(2))
	if v113 <= v117 {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_s_lock(m, v124, int32(_a_F_pa_xact_finish_3), int32(649), int32(_a_F_pa_xact_finish_4))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+16))
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v130)+12)))
	v133 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v130))), uint32(v133))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v136 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_shm_mq_detach(m, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v142 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[6]))
	v146 = F_LWLockAcquire(m, v142+int32(_a_F_pa_xact_finish_5), int32(1))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
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
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[7]))
	v152 = v149 + v131*int32(112)
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152)+34)))
	if v153 != v132 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v165 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[6]))
	F_LWLockRelease(m, v165+int32(_a_F_pa_xact_finish_5))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v156 = v152 + int32(16)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+20))
	if v157 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	F_logicalrep_worker_stop_internal(m, v156, int32(12))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
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
	v171 = m.ExcPending
	if v171 != 0 {
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
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_pa_xact_finish_6), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_pa_xact_finish_0), int32(1307), int32(_a_F_pa_xact_finish_7))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_pa_xact_finish_8), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_pa_xact_finish_0), int32(563), int32(_a_F_pa_xact_finish_9))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
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
