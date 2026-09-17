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
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int64
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
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
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(1)
	if v21 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+32))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	v56 = int32(1)
	F_LockApplyTransactionForSession(m, v53, v55, v56, v56)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
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
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	if v31 == v29 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L7
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[1]))
	v39 = F_WaitLatch(m, v35, int32(41), int32(10), int32(134217758))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
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
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[1]))
	*(*int32)(unsafe.Add(mBase, uint32(v42))) = int32(0)
	goto L13
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[2]))
	if v46 == int32(0) {
		goto L3
	} else {
		goto L14
	}
L14:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	goto L3
L16:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[0]))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+32))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	v65 = int32(1)
	F_UnlockApplyTransactionForSession(m, v62, v64, v65, v65)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(1)
	if v70 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_s_lock(m, v69, int32(_a_F_pa_xact_finish_0), int32(1330), int32(_a_F_pa_xact_finish_1))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = int32(0)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if v80 == int32(2) {
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
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v86 = *(*int64)(unsafe.Add(mBase, uint32(v85)+24))
	F_store_flush_position(m, l1, v86)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[3]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v96 = F_hash_search(m, v90, v91+int32(4), int32(2), int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	if v96 == int32(0) {
		goto L22
	} else {
		goto L31
	}
L31:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v100 == int32(0) {
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
	v104 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[4]))
	if v104 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = int32(1)
	if v115 != 0 {
		goto L40
	} else {
		goto L41
	}
L36:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v107 = v105
	goto L38
L37:
	;
	v107 = int32(0)
	goto L38
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[5]))
	v111 = base.I32_div_s(v109, int32(2))
	if v107 <= v111 {
		goto L32
	} else {
		goto L39
	}
L39:
	;
	goto L35
L40:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_s_lock(m, v118, int32(_a_F_pa_xact_finish_2), int32(649), int32(_a_F_pa_xact_finish_3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = int32(0)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v124)+16))
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+12)))
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
	v139 = F_LWLockAcquire(m, v135+int32(_a_F_pa_xact_finish_4), int32(1))
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
	v145 = v142 + v127*int32(112)
	v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+34)))
	if v128 != v146 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_pa_xact_finish[6]))
	F_LWLockRelease(m, v158+int32(_a_F_pa_xact_finish_4))
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
	F_errmsg(m, int32(_a_F_pa_xact_finish_5), int32(0))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	F_errfinish(m, int32(_a_F_pa_xact_finish_0), int32(1307), int32(_a_F_pa_xact_finish_6))
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
	F_errmsg_internal(m, int32(_a_F_pa_xact_finish_7), int32(0))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_pa_xact_finish_0), int32(563), int32(_a_F_pa_xact_finish_8))
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
