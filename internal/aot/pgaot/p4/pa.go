package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
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
		v12 = *(*int32)(unsafe.Add(mBase, _consts[816]))
		if v12 == int32(0) {
			v32 = v2
			m.G0 = v6 + int32(16)
			return v32
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _consts[817]))
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[818]))
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
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
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
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	v8 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+32))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
	F_UnlockApplyTransactionForSession(m, v9, v11, int32(0), int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
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
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(1)
	if v23 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+32))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
	v58 = int32(1)
	F_LockApplyTransactionForSession(m, v55, v57, v58, v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L16
	}
L5:
	;
	F_s_lock(m, v22, int32(472255), int32(1330), int32(334459))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v31 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v33 == v31 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	v41 = F_WaitLatch(m, v37, int32(41), int32(10), int32(134217758))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
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
	v44 = *(*int32)(unsafe.Add(mBase, _consts[496]))
	*(*int32)(unsafe.Add(mBase, uint32(v44))) = int32(0)
	goto L13
L13:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v48 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _consts[818]))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+32))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	v67 = int32(1)
	F_UnlockApplyTransactionForSession(m, v64, v66, v67, v67)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(1)
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_s_lock(m, v71, int32(472255), int32(1330), int32(334459))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = int32(0)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v82 == int32(2) {
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
	v190 = m.ExcPending
	if v190 != 0 {
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
	v174 = m.ExcPending
	if v174 != 0 {
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
	v92 = *(*int32)(unsafe.Add(mBase, _consts[816]))
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
	v169 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v169)
	return
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _consts[819]))
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
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = int32(1)
	if v117 != 0 {
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
	v111 = *(*int32)(unsafe.Add(mBase, _consts[820]))
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
	F_s_lock(m, v120, int32(472331), int32(649), int32(222478))
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
	*(*int32)(unsafe.Add(mBase, uint32(v126))) = int32(0)
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v126)+16))
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+12)))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v131 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L42
L44:
	;
	F_shm_mq_detach(m, v131)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v137 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v141 = F_LWLockAcquire(m, v137+int32(5504), int32(1))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
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
	v144 = *(*int32)(unsafe.Add(mBase, _consts[821]))
	v147 = v144 + v129*int32(112)
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v147)+34)))
	if v148 != v130&int32(65535) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v162+int32(5504))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L53
	}
L50:
	;
	v153 = v147 + int32(16)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)+20))
	if v154 == int32(0) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	F_logicalrep_worker_stop_internal(m, v153, int32(12))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
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
	v168 = m.ExcPending
	if v168 != 0 {
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
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	F_errmsg(m, int32(209843), int32(0))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(472255), int32(1307), int32(306589))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
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
	F_errmsg_internal(m, int32(424111), int32(0))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(472255), int32(563), int32(209612))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
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
