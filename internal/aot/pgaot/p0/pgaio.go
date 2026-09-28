package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgaio_io_perform_synchronously(m *base.Module, l0 int32) {
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
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int64
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v10 = int32(_a_F_pgaio_io_perform_synchronously_0)
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1])) = v12 + int32(1)
	v18 = v8 + v9<<(uint(int32(3))%32)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	switch v20 {
	case 0:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v49 = m.ExcPending
		if v49 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_pgaio_io_perform_synchronously_1), int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_pgaio_io_perform_synchronously_2), int32(141), int32(_a_F_pgaio_io_perform_synchronously_3))
				mBase = m.M
				v58 = m.ExcPending
				if v58 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	case 1:
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[2]))
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = int32(167772183)
		v25 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+92)))
		if v27 != int32(1) {
			v59 = F_preadv(m, v26, v18, v27, v25)
			mBase = m.M
			v63 = v59
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v32 = F_pread(m, v26, v30, v31, v25)
			mBase = m.M
			v63 = v32
		}
		v65 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[2]))
		v66 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66
		if v66 <= v63 {
			v74 = v63
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[3]))
			v74 = int32(0) - v72
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74
		F_pgaio_io_process_completion(m, l0, v74)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			v80 = int32(_a_F_pgaio_io_perform_synchronously_0)
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1])) = v82 - int32(1)
			return
		}
	case 2:
		v34 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[2]))
		*(*int32)(unsafe.Add(mBase, uint32(v34))) = int32(167772186)
		v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+92)))
		if v39 == int32(1) {
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
			v43 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
			v44 = F_pwrite(m, v38, v42, v43, v37)
			mBase = m.M
			v63 = v44
		} else {
			v45 = F_pwritev(m, v38, v18, v39, v37)
			mBase = m.M
			v63 = v45
		}
		v65 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[2]))
		v66 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v65))) = v66
		if v66 <= v63 {
			v74 = v63
		} else {
			v72 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[3]))
			v74 = int32(0) - v72
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74
		F_pgaio_io_process_completion(m, l0, v74)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			v80 = int32(_a_F_pgaio_io_perform_synchronously_0)
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1])) = v82 - int32(1)
			return
		}
	default:
		v74 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v74
		F_pgaio_io_process_completion(m, l0, v74)
		mBase = m.M
		v79 = m.ExcPending
		if v79 != 0 {
			return
		} else {
			v80 = int32(_a_F_pgaio_io_perform_synchronously_0)
			v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgaio_io_perform_synchronously[1])) = v82 - int32(1)
			return
		}
	}
}
func F_pgaio_worker_needs_synchronous_execution(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(1)
	v5 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_pgaio_worker_needs_synchronous_execution[0])))
	if v5 != v3 {
		v18 = v3
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)))
		if v8&int32(2) != 0 {
			v18 = v3
		} else {
			v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11<<(uint(int32(2))%32))+uint32(_c_F_pgaio_worker_needs_synchronous_execution[1])))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
			v18 = base.B2i32(v15 == int32(0))
		}
	}
	return v18
}
func F_pgaio_worker_pm_test_grow_signal_sent(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	v1 = int32(0)
	v5 = base.AtomicRmwOr32(m, v1, int32(_a_F_pgaio_worker_pm_test_grow_signal_sent_0), v1)
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_pm_test_grow_signal_sent[0]))
	if v7 != 0 {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
		v10 = v8
	} else {
		v10 = int32(0)
	}
	return v10 & int32(1)
}
func F_pgaio_worker_submit(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int64
	_ = v152
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v162 int64
	_ = v162
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v237 int32
	_ = v237
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	v3 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l0 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v13 + int32(16)
	return l0
L2:
	;
	v258 = int32(0)
	goto L58
L3:
	;
	v151 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[0]))
	v152 = *(*int64)(unsafe.Add(mBase, uint32(v151)+8))
	if v152 == int64(0) {
		goto L37
	} else {
		goto L38
	}
L4:
	;
	v18 = v3
	goto L7
L5:
	;
	goto L6
L6:
	;
	v132 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[1]))
	v136 = F_LWLockConditionalAcquire(m, v132+int32(_a_F_pgaio_worker_submit_0), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L9
	} else {
		goto L34
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1+v18<<(uint(int32(2))%32))))
	F_pgaio_io_update_state(m, v28, int32(4))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[1]))
	v64 = F_LWLockConditionalAcquire(m, v60+int32(_a_F_pgaio_worker_submit_0), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L9
	} else {
		goto L15
	}
L9:
	;
	return int32(0)
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[2]))
	v37 = v35 + int32(152)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+156))
	if v38 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+160)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+156)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v35)+152)) = v37
	goto L13
L12:
	;
	goto L13
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+28)) = v37
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v35)+152))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v46
	v49 = v28 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+4)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v35)+152)) = v49
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v35)+160))
	v53 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+160)) = v52 + v53
	v57 = v18 + v53
	if v57 != l0 {
		v18 = v57
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	if v64 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v246 = l1
	v248 = l0
	goto L2
L17:
	;
	goto L18
L18:
	;
	v72 = v3
	goto L19
L19:
	;
	v80 = l1 + v72<<(uint(int32(2))%32)
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[3]))
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v84 = int32(1)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v89 = (v83 - v84) & (v86 + v84)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v82)+8))
	if v89 == v90 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v129 = int32(0)
	v142 = v129
	v144 = v129
	goto L3
L21:
	;
	v94 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L9
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[4]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+24))
	goto L32
L24:
	;
	if v94 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	F_errhidestmt(m)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L9
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v142 = v80
	v144 = l0 - v72
	goto L3
L28:
	;
	F_errhidecontext(m)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L9
	} else {
		goto L29
	}
L29:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[3]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v102
	F_errmsg_internal(m, int32(_a_F_pgaio_worker_submit_1), v13)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L9
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(_a_F_pgaio_worker_submit_2), int32(425), int32(_a_F_pgaio_worker_submit_3))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L9
	} else {
		goto L31
	}
L31:
	;
	goto L27
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v82+v120<<(uint(int32(2))%32))+12)) = (v113 - v116) >> (uint(int32(7)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v82)+4)) = v89
	v127 = v72 + int32(1)
	if v127 != l0 {
		v72 = v127
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L20
L34:
	;
	if v136 == int32(0) {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v142 = v3
	v144 = v3
	goto L3
L36:
	;
	if v144 <= int32(0) {
		goto L1
	} else {
		goto L57
	}
L37:
	;
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[1]))
	F_LWLockRelease(m, v156+int32(_a_F_pgaio_worker_submit_0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L9
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v162 = base.I64_ctz(v152)
	*(*int64)(unsafe.Add(mBase, uint32(v151)+8)) = v152 & base.I64_rotl(int64(-2), v162)
	v167 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[1]))
	F_LWLockRelease(m, v167+int32(_a_F_pgaio_worker_submit_0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L9
	} else {
		goto L41
	}
L40:
	;
	goto L36
L41:
	;
	v173 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[0]))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v173+base.I32_wrap_i64(v162)<<(uint(int32(2))%32))+28))
	if v178 == int32(-1) {
		goto L36
	} else {
		goto L42
	}
L42:
	;
	v182 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[5]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
	v188 = v183 + v178*int32(768) + int32(316)
	v189 = int32(0)
	v192 = base.AtomicRmwOr32(m, v189, int32(_a_F_pgaio_worker_submit_4), v189)
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v188)))
	if v193 != 0 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L36
L44:
	;
	goto L43
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v188))) = int32(1)
	v196 = int32(0)
	v199 = base.AtomicRmwOr32(m, v196, int32(_a_F_pgaio_worker_submit_4), v196)
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v188)+4))
	if v200 == v196 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v188)+12))
	if v203 == int32(0) {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[6]))
	if v207 == v203 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v209 = m.G0
	v211 = v209 - int32(16)
	m.G0 = v211
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[7]))
	if v214 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L50
L50:
	;
	v237 = F_pgmem_kill(m, v203, int32(23))
	mBase = m.M
	goto L44
L51:
	;
	m.G0 = v211 + int32(16)
	goto L43
L52:
	;
	v217 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v211)+15)) = uint8(v217)
	goto L53
L53:
	;
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[8]))
	v225 = F_write(m, v221, v211+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v225 {
		goto L51
	} else {
		goto L55
	}
L54:
	;
	goto L51
L55:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_submit[9]))
	if v229 == int32(27) {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	goto L54
L57:
	;
	v246 = v142
	v248 = v144
	goto L2
L58:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v246+v258<<(uint(int32(2))%32))))
	F_pgaio_io_perform_synchronously(m, v268)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L9
	} else {
		goto L60
	}
L59:
	;
	goto L1
L60:
	;
	v272 = v258 + int32(1)
	if v272 != v248 {
		v258 = v272
		goto L58
	} else {
		goto L61
	}
L61:
	;
	goto L59
}
func F_pgaio_workers_enabled(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_workers_enabled[0]))
	return base.B2i32(v2 == int32(1))
}
