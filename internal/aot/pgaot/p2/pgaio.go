package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_pgaio_io_get_op_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)))
	if base.Ui32(v2) <= base.Ui32(int32(2)) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_get_op_name[0])))
		v9 = v7
	} else {
		v9 = int32(0)
	}
	return v9
}
func F_pgaio_io_get_state_name(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.Ui32(v2) <= base.Ui32(int32(7)) {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v2<<(uint(int32(2))%32))+uint32(_c_F_pgaio_io_get_state_name[0])))
		v9 = v7
	} else {
		v9 = int32(0)
	}
	return v9
}
func F_pgaio_submit_staged(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_submit_staged[0]))
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+22)))
	if v10 == int32(0) {
		m.G0 = v6 + int32(16)
		return
	} else {
		v13 = int32(_a_F_pgaio_submit_staged_0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_submit_staged[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_pgaio_submit_staged[1])) = v15 + int32(1)
		v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_submit_staged[2]))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+32))
		v24 = m.T0[v23].(func(*base.Module, int32, int32) int32)(m, v10, v9+int32(24))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			v26 = int32(_a_F_pgaio_submit_staged_0)
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_submit_staged[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_pgaio_submit_staged[1])) = v28 - int32(1)
			v33 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_submit_staged[0]))
			v34 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v33)+22)) = uint16(v34)
			v38 = F_errstart(m, int32(11), v34)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				if v38 == int32(0) {
					m.G0 = v6 + int32(16)
					return
				} else {
					F_errhidestmt(m)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errhidecontext(m)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6))) = v24
							F_errmsg_internal(m, int32(_a_F_pgaio_submit_staged_1), v6)
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_pgaio_submit_staged_2), int32(1157), int32(_a_F_pgaio_submit_staged_3))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									m.G0 = v6 + int32(16)
									return
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_pgaio_worker_die(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v17 int64
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v49 int64
	_ = v49
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int64
	_ = v66
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[0]))
	v10 = F_LWLockAcquire(m, v6+int32(_a_F_pgaio_worker_die_0), int32(0))
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
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[1]))
	v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
	v17 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[2])))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v14 & base.I64_rotl(int64(-2), v17)
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[0]))
	F_LWLockRelease(m, v22+int32(_a_F_pgaio_worker_die_0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[0]))
	v32 = F_LWLockAcquire(m, v28+int32(_a_F_pgaio_worker_die_1), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = int32(_a_F_pgaio_worker_die_2)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[1]))
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[2]))
	*(*int32)(unsafe.Add(mBase, uint32(v35+v37<<(uint(int32(2))%32))+28)) = int32(-1)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[1]))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v44)+16))
	v49 = v45 & base.I64_rotl(int64(-2), base.I64_extend_i32_u(v37))
	*(*int64)(unsafe.Add(mBase, uint32(v44)+16)) = v49
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+24)) = v51 - int32(1)
	v56 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[0]))
	F_LWLockRelease(m, v56+int32(_a_F_pgaio_worker_die_1))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v49 != int64(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[1]))
	v65 = v64
	v66 = v49
	goto L9
L7:
	;
	goto L8
L8:
	;
	return
L9:
	;
	v70 = base.I64_ctz(v66)
	v72 = v66 & base.I64_rotl(int64(-2), v70)
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v65+base.I32_wrap_i64(v70)<<(uint(int32(2))%32))+28))
	if v77 != int32(-1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[3]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)))
	v87 = v82 + v77*int32(768) + int32(316)
	v88 = int32(0)
	v91 = base.AtomicRmwOr32(m, v88, int32(_a_F_pgaio_worker_die_3), v88)
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v92 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v141 = v65
	goto L13
L13:
	;
	if v72 != int64(0) {
		v65 = v141
		v66 = v72
		goto L9
	} else {
		goto L28
	}
L14:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[1]))
	v141 = v140
	goto L13
L15:
	;
	goto L14
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = int32(1)
	v95 = int32(0)
	v98 = base.AtomicRmwOr32(m, v95, int32(_a_F_pgaio_worker_die_3), v95)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v87)+4))
	if v99 == v95 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	if v102 == int32(0) {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[4]))
	if v106 == v102 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v108 = m.G0
	v110 = v108 - int32(16)
	m.G0 = v110
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[5]))
	if v113 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	v136 = F_pgmem_kill(m, v102, int32(23))
	mBase = m.M
	goto L15
L22:
	;
	m.G0 = v110 + int32(16)
	goto L14
L23:
	;
	v116 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v110)+15)) = uint8(v116)
	goto L24
L24:
	;
	v120 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[6]))
	v124 = F_write(m, v120, v110+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v124 {
		goto L22
	} else {
		goto L26
	}
L25:
	;
	goto L22
L26:
	;
	v128 = *(*int32)(unsafe.Add(mBase, _c_F_pgaio_worker_die[7]))
	if v128 == int32(27) {
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
L28:
	;
	goto L10
}
func F_pgaio_worker_shmem_request(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	v2 = m.G0
	v4 = v2 - int32(32)
	m.G0 = v4
	*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(_a_F_pgaio_worker_shmem_request_0)
	*(*int64)(unsafe.Add(mBase, uint32(v4)+20)) = int64(268)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(_a_F_pgaio_worker_shmem_request_1)
	F_ShmemRequestStructWithOpts(m, v4+int32(16))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(_a_F_pgaio_worker_shmem_request_2)
		*(*int64)(unsafe.Add(mBase, uint32(v4)+4)) = int64(156)
		*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_pgaio_worker_shmem_request_3)
		F_ShmemRequestStructWithOpts(m, v4)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			m.G0 = v4 + int32(32)
			return
		}
	}
}
func F_pgaio_wref_valid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return base.B2i32(v2 != int32(-1))
}
