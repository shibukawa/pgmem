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
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
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
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
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
	v53 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[0]))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+32))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v57 = int32(1)
	F_LockApplyTransactionForSession(m, v54, v56, v57, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v43))) = int32(0)
	goto L13
L13:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[2]))
	if v47 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[0]))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)+32))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v66 = int32(1)
	F_UnlockApplyTransactionForSession(m, v63, v65, v66, v66)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v73 = base.AtomicRmwXchg32(m, v70, int32(0), int32(1))
	if v73 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_s_lock(m, v70, int32(_a_F_pa_xact_finish_0), int32(1330), int32(_a_F_pa_xact_finish_1))
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	v80 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v70))), uint32(v80))
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
	v189 = m.ExcPending
	if v189 != 0 {
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
	v173 = m.ExcPending
	if v173 != 0 {
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
	v168 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v168)
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
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_s_lock(m, v120, int32(_a_F_pa_xact_finish_2), int32(649), int32(_a_F_pa_xact_finish_3))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+12)))
	v129 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v126))), uint32(v129))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v132 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_shm_mq_detach(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v138 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[6]))
	v142 = F_LWLockAcquire(m, v138+int32(_a_F_pa_xact_finish_4), int32(1))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
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
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[7]))
	v148 = v145 + v127*int32(112)
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+34)))
	if v149 != v128 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[6]))
	F_LWLockRelease(m, v161+int32(_a_F_pa_xact_finish_4))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v152 = v148 + int32(16)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+20))
	if v153 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	F_logicalrep_worker_stop_internal(m, v152, int32(12))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
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
	v167 = m.ExcPending
	if v167 != 0 {
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
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(_a_F_pa_xact_finish_5), int32(0))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_pa_xact_finish_0), int32(1307), int32(_a_F_pa_xact_finish_6))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
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
	F_errmsg_internal(m, int32(_a_F_pa_xact_finish_7), int32(0))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_pa_xact_finish_0), int32(563), int32(_a_F_pa_xact_finish_8))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
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
