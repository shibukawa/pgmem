package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BTreeTupleGetPointsToTID(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	v2 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+7)))
	if v2&int32(32) == int32(0) {
		v18 = l0
	} else {
		v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
		if v7&int32(32) == int32(0) {
			v18 = l0
		} else {
			v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
			v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
			v18 = v12 + (l0 + v13<<(uint(int32(16))%32))
		}
	}
	return v18
}
func F_BackgroundWorkerInitializeConnectionByOid(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerInitializeConnectionByOid[0]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+192)))
	if v6&int32(2) != 0 {
		v9 = int32(0)
		F_InitPostgres(m, v9, l0, v9, l1, l2<<(uint(int32(1))%32)&int32(6), v9)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerInitializeConnectionByOid[1]))
			if v19 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_BackgroundWorkerInitializeConnectionByOid_0), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_BackgroundWorkerInitializeConnectionByOid_1), int32(913), int32(_a_F_BackgroundWorkerInitializeConnectionByOid_2))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_BackgroundWorkerInitializeConnectionByOid[1])) = int32(2)
				return
			}
		}
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_BackgroundWorkerInitializeConnectionByOid_3), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_BackgroundWorkerInitializeConnectionByOid_1), int32(903), int32(_a_F_BackgroundWorkerInitializeConnectionByOid_2))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		}
	}
}
func F_BarrierArriveAndDetachExceptLast(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v3 != 0 {
		F_s_lock(m, l0, int32(_a_F_BarrierArriveAndDetachExceptLast_0), int32(215), int32(_a_F_BarrierArriveAndDetachExceptLast_1))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if int32(2) <= v13 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 - int32(1)
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19 + int32(1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
			return base.B2i32(v13 < int32(2))
		}
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if int32(2) <= v13 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v13 - int32(1)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19 + int32(1)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
		return base.B2i32(v13 < int32(2))
	}
}
func F_BarrierInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v11 int32
	_ = v11
	v2 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2
	v4 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v4
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = v4
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v2)
	v11 = l0 + int32(24)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = int64(-4294967296)
	return
}
func F_BeforeShmemExit_Files(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[0]))
	if base.Ui32(int32(2)) <= base.Ui32(v6) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[1]))
	v12 = int32(1)
	v13 = v6
	v14 = v10
	goto L4
L2:
	;
	goto L3
L3:
	;
	v43 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[2])) = uint8(v43)
	v46 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[3]))
	if v43 < v46 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v18 = v14 + v12*int32(48)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
	if v19&int32(3) == int32(0) {
		v33 = v13
		v34 = v14
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v36 = v12 + int32(1)
	if base.Ui32(v36) < base.Ui32(v33) {
		v12 = v36
		v13 = v33
		v14 = v34
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v18)+32))
	if v24 == int32(0) {
		v33 = v13
		v34 = v14
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_FileClose(m, v12)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[0]))
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[1]))
	v33 = v30
	v34 = v32
	goto L6
L11:
	;
	goto L5
L12:
	;
	goto L15
L13:
	;
	goto L14
L14:
	;
	return
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[4]))
	v55 = F_FreeDesc(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L9
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[3]))
	if int32(0) < v58 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
}
func F_BlockSampler_HasMore(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v2) < base.Ui32(v3) {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = base.B2i32(v5 < v6)
	} else {
		v9 = int32(0)
	}
	return v9
}
func F___bswap_32(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	v4 = int32(16711935)
	return base.I32_rotr(l0, int32(24))&v4 | base.I32_rotr(l0&v4, int32(8))
}
func F_before_shmem_exit(m *base.Module, l0 int32, l1 int32) {
	var v10 int32
	_ = v10
	Fn13847(m, l0, l1, int32(_a_F_before_shmem_exit_0), int32(349), int32(_a_F_before_shmem_exit_1), int32(_a_F_before_shmem_exit_2), int32(_a_F_before_shmem_exit_3), int32(_a_F_before_shmem_exit_4))
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		return
	}
}
func F_before_stmt_triggers_fired(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[0]))
	if int32(0) <= v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[1]))
	if v9 < v13 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L13
	} else {
		goto L36
	}
L4:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[2]))
	v82 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[0]))
	v85 = v80 + v82*int32(20)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	if v86 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L5:
	;
	v16 = v9 + int32(1)
	if v13 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[1])) = v41
	*(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[2])) = v43
	if v41 <= v13 {
		goto L4
	} else {
		goto L19
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[3]))
	v21 = int32(8)
	if v16 <= v21 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[2]))
	v34 = v13 << (uint(int32(1)) % 32)
	if v34 < v16 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v24 = v21
	goto L12
L11:
	;
	v24 = v16
	goto L12
L12:
	;
	v27 = F_MemoryContextAlloc(m, v20, v24*int32(20))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v41 = v24
	v43 = v27
	goto L6
L15:
	;
	v36 = v16
	goto L17
L16:
	;
	v36 = v34
	goto L17
L17:
	;
	v39 = F_repalloc(m, v32, v36*int32(20))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v41 = v36
	v43 = v39
	goto L6
L19:
	;
	v52 = v13
	goto L20
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[2]))
	v60 = v57 + v52*int32(20)
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = int32(0)
	v63 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v63
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = v63
	v68 = v52 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[1]))
	if v68 < v70 {
		v52 = v68
		goto L20
	} else {
		goto L22
	}
L21:
	;
	goto L4
L22:
	;
	goto L21
L23:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141)+9)))
	v147 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v141)+9)) = uint8(v147)
	return v146
L24:
	;
	v122 = int32(_a_F_before_stmt_triggers_fired_0)
	v123 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[4]))
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[5]))
	*(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[4])) = v126
	v129 = F_palloc0(m, int32(36))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L13
	} else {
		goto L34
	}
L25:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+4))
	if v89 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v86)+12))
	v97 = int32(0)
	goto L27
L27:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v92+v97<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
	if v105 != l0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L24
L29:
	;
	v113 = v97 + int32(1)
	if v89 != v113 {
		v97 = v113
		goto L27
	} else {
		goto L33
	}
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	if v107 != l1 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+8)))
	if v109 != int32(1) {
		v141 = v104
		goto L23
	} else {
		goto L32
	}
L32:
	;
	goto L29
L33:
	;
	goto L28
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = l0
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v85)+16))
	v134 = F_lappend(m, v133, v129)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v85)+16)) = v134
	*(*int32)(unsafe.Add(mBase, _c_F_before_stmt_triggers_fired[4])) = v123
	v141 = v129
	goto L23
L36:
	;
	F_errmsg_internal(m, int32(_a_F_before_stmt_triggers_fired_1), int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(_a_F_before_stmt_triggers_fired_2), int32(_a_F_before_stmt_triggers_fired_3), int32(_a_F_before_stmt_triggers_fired_4))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_begin_cb_wrapper(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int64
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_begin_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(993)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v14
	v18 = int32(_a_F_begin_cb_wrapper_1)
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_begin_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_begin_cb_wrapper[0])) = v8 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v8 + int32(16)
	v28 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+147)) = uint8(v28)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+160)) = v30
	v32 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v10)+164)) = uint8(v3)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+152)) = v32
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
	m.T0[v36].(func(*base.Module, int32, int32))(m, v10, l1)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		return
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		*(*int32)(unsafe.Add(mBase, _c_F_begin_cb_wrapper[0])) = v40
		m.G0 = v8 + int32(32)
		return
	}
}
func F_big5_to_utf8(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn13848(m, l0, int32(36), v3, v3, v3, int32(_a_F_big5_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_binaryheap_allocate(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	v9 = F_palloc(m, l0<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l2
		*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l0
		v16 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v9)+8)) = uint8(v16)
		*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(0)
		return v9
	}
}
func F_binaryheap_remove_first(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
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
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return v10
L2:
	;
	goto L3
L3:
	;
	v18 = v11 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v18
	v21 = l0 + int32(20)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21+v18<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v25
	v28 = v18
	v32 = int32(0)
	goto L4
L4:
	;
	v36 = int32(1)
	v37 = v32 << (uint(v36) % 32)
	v39 = v37 | v36
	v41 = v37 + int32(2)
	if v41 < v28 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+v32<<(uint(int32(2))%32)))) = v25
	return v10
L6:
	;
	goto L5
L7:
	;
	v43 = int32(2)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v21+v39<<(uint(v43)%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v21+v41<<(uint(v43)%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v53 = m.T0[v52].(func(*base.Module, int32, int32, int32) int32)(m, v46, v50, v51)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v61 = v39
	v62 = v28
	goto L9
L9:
	;
	if v62 <= v39 {
		goto L6
	} else {
		goto L15
	}
L10:
	;
	return int32(0)
L11:
	;
	if v53 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v59 = v41
	goto L14
L13:
	;
	v59 = v39
	goto L14
L14:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v61 = v59
	v62 = v60
	goto L9
L15:
	;
	v66 = v21 + v61<<(uint(int32(2))%32)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = m.T0[v69].(func(*base.Module, int32, int32, int32) int32)(m, v25, v67, v68)
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	if int32(0) <= v70 {
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v21+v32<<(uint(int32(2))%32)))) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = v79
	v32 = v61
	goto L4
}
func F_bitcmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
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
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
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
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int32(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v18 = v15 + int32(8)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v20 = int32(2)
	v21 = int32(base.Ui32(v19) >> (uint(v20) % 32))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	v24 = int32(base.Ui32(v22) >> (uint(v20) % 32))
	if base.Ui32(v21) < base.Ui32(v24) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = v21
	goto L7
L6:
	;
	v26 = v24
	goto L7
L7:
	;
	v28 = v26 - int32(8)
	if base.Ui32(int32(4)) <= base.Ui32(v28) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if v90 != 0 {
		v99 = v90
		goto L1
	} else {
		goto L26
	}
L9:
	;
	v90 = int32(0)
	goto L8
L10:
	;
	v64 = v59
	v65 = v60
	v66 = v61
	goto L20
L11:
	;
	if (v13|v18)&int32(3) != 0 {
		v59 = v13
		v60 = v18
		v61 = v28
		goto L10
	} else {
		goto L14
	}
L12:
	;
	v52 = v13
	v53 = v18
	v54 = v28
	goto L13
L13:
	;
	if v54 == int32(0) {
		goto L9
	} else {
		goto L19
	}
L14:
	;
	v36 = v13
	v37 = v18
	v38 = v28
	goto L15
L15:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	if v41 != v42 {
		v59 = v36
		v60 = v37
		v61 = v38
		goto L10
	} else {
		goto L17
	}
L16:
	;
	v52 = v47
	v53 = v45
	v54 = v49
	goto L13
L17:
	;
	v44 = int32(4)
	v45 = v37 + v44
	v47 = v36 + v44
	v49 = v38 - v44
	if base.Ui32(int32(3)) < base.Ui32(v49) {
		v36 = v47
		v37 = v45
		v38 = v49
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	v59 = v52
	v60 = v53
	v61 = v54
	goto L10
L20:
	;
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64))))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	if v69 == v70 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v90 = v69 - v70
	goto L8
L22:
	;
	v72 = int32(1)
	v77 = v66 - v72
	if v77 != 0 {
		v64 = v64 + v72
		v65 = v65 + v72
		v66 = v77
		goto L20
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	goto L21
L25:
	;
	goto L9
L26:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v92 == v93 {
		v99 = int32(0)
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v92 < v93 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v98 = int32(-1)
	goto L30
L29:
	;
	v98 = int32(1)
	goto L30
L30:
	;
	v99 = v98
	goto L1
L31:
	;
	F_pfree(m, v8)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L2
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v106 != v15 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_pfree(m, v15)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	return v99
L38:
	;
	goto L37
}
func F_bitmapheap_stream_read_next(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v135 int32
	_ = v135
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v404 int32
	_ = v404
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	v17 = l1 + int32(16)
	goto L1
L1:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_bitmapheap_stream_read_next[0]))
	if v34 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v425
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v40 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	return int32(0)
L7:
	;
	goto L5
L8:
	;
	if v420 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L9:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v48 = v46 + int32(28)
	v50 = F_LWLockAcquire(m, v48, int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v39)+8))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+28))
	if v265 <= v263 {
		goto L46
	} else {
		goto L47
	}
L12:
	;
	v420 = v256
	goto L8
L13:
	;
	v52 = int32(4)
	v53 = v45 + v52
	v55 = v44 + v52
	v57 = v43 + v52
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46)+48))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	if v59 <= v58 {
		v135 = v58
		goto L17
	} else {
		goto L18
	}
L14:
	;
	F_LWLockRelease(m, v48)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L6
	} else {
		goto L42
	}
L15:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v57+v198<<(uint(int32(2))%32))))
	if v45 != 0 {
		goto L38
	} else {
		goto L39
	}
L16:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v196 <= v195 {
		goto L14
	} else {
		goto L37
	}
L17:
	;
	if v59 <= v135 {
		goto L16
	} else {
		goto L31
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	v62 = v61
	v67 = v58
	goto L19
L19:
	;
	v77 = int32(256)
	if v62 <= v77 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L16
L21:
	;
	v80 = v77
	goto L23
L22:
	;
	v80 = v62
	goto L23
L23:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v55+v67<<(uint(int32(2))%32))))
	v90 = v62
	goto L25
L24:
	;
	v123 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+52)) = v123
	v127 = v67 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+48)) = v127
	if v127 != v59 {
		v62 = v123
		v67 = v127
		goto L19
	} else {
		goto L30
	}
L25:
	;
	if v90 == v80 {
		goto L24
	} else {
		goto L27
	}
L26:
	;
	if int32(255) < v90 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v106 = int32(1)
	v109 = base.I32_div_s(v90, int32(32))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v53+v84*int32(48)+int32(8)+v109<<(uint(int32(2))%32))))
	if int32(base.Ui32(v113)>>(uint(v90)%32))&v106 == int32(0) {
		v90 = v90 + v106
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+52)) = v90
	v135 = v67
	goto L17
L30:
	;
	goto L20
L31:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v55+v135<<(uint(int32(2))%32))))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v53+v150*int32(48))))
	v155 = v146 + v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v156 < v157 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v57+v156<<(uint(int32(2))%32))))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v53+v162*int32(48))))
	if base.Ui32(v166) <= base.Ui32(v155) {
		v198 = v156
		goto L15
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v170 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v170)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v155
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v46)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+52)) = v173 + int32(1)
	F_LWLockRelease(m, v48)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L6
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	v256 = int32(1)
	goto L12
L37:
	;
	v198 = v195
	goto L15
L38:
	;
	v220 = v53
	goto L40
L39:
	;
	v220 = int32(0)
	goto L40
L40:
	;
	v221 = v216*int32(48) + v220
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v221
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v221)))
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v224)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v223
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v221)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v227)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v46)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+44)) = v229 + int32(1)
	F_LWLockRelease(m, v48)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L6
	} else {
		goto L41
	}
L41:
	;
	v256 = int32(1)
	goto L12
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	v256 = int32(0)
	goto L12
L43:
	;
	v420 = v404
	goto L8
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	v404 = int32(0)
	goto L43
L45:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v264)+8))
	if v370 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L46:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v264)+24))
	if v360 <= v359 {
		goto L44
	} else {
		goto L65
	}
L47:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v270 = v267
	v273 = v263
	goto L48
L48:
	;
	v276 = int32(256)
	if v270 <= v276 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L46
L50:
	;
	v279 = v276
	goto L52
L51:
	;
	v279 = v270
	goto L52
L52:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v264)+92))
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v280+v273<<(uint(int32(2))%32))))
	v289 = v270
	goto L54
L53:
	;
	v343 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v343
	v347 = v273 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+8)) = v347
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v264)+28))
	if v347 < v349 {
		v270 = v343
		v273 = v347
		goto L48
	} else {
		goto L64
	}
L54:
	;
	if v289 == v279 {
		goto L53
	} else {
		goto L56
	}
L55:
	;
	if int32(255) < v289 {
		goto L53
	} else {
		goto L58
	}
L56:
	;
	v296 = int32(1)
	v299 = base.I32_div_s(v289, int32(32))
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v284+int32(8)+v299<<(uint(int32(2))%32))))
	if int32(base.Ui32(v303)>>(uint(v289)%32))&v296 == int32(0) {
		v289 = v289 + v296
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v289
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v264)+28))
	if v312 <= v273 {
		goto L46
	} else {
		goto L59
	}
L59:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v264)+92))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v315+v273<<(uint(int32(2))%32))))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v319)))
	v321 = v314 + v320
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v264)+24))
	if v322 < v323 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v264)+88))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v325+v322<<(uint(int32(2))%32))))
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	if base.Ui32(v330) <= base.Ui32(v321) {
		v364 = v322
		goto L45
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v334 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v334)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v321
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	v338 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v337 + v338
	v404 = v338
	goto L43
L63:
	;
	goto L62
L64:
	;
	goto L49
L65:
	;
	v364 = v359
	goto L45
L66:
	;
	v380 = v264 + int32(40)
	goto L68
L67:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v264)+88))
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v375+v364<<(uint(int32(2))%32))))
	v380 = v379
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v380
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v380)))
	v383 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v383)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v382
	v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v380)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v386)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
	v389 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+4)) = v388 + v389
	v404 = v389
	goto L43
L69:
	;
	return int32(-1)
L70:
	;
	goto L71
L71:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v427 = *(*int32)(unsafe.Add(mBase, _c_F_bitmapheap_stream_read_next[1]))
	if v427 != int32(3) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if base.Ui32(v430) <= base.Ui32(v425) {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	goto L2
L75:
	;
	goto L74
}
func F_bitshiftright(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
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
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v201 int32
	_ = v201
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v15 < int32(0) {
			v19 = int32(0)
			v21 = int32(-2147483640)
			if base.Ui32(v15) <= base.Ui32(v21) {
				v24 = v21
			} else {
				v24 = v15
			}
			v26 = F_DirectFunctionCall2Coll(m, int32(1534), v19, v11, v19-v24)
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				return v26
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v32 = F_palloc(m, int32(base.Ui32(v29)>>(uint(int32(2))%32)))
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v36 = v34 & int32(-4)
				*(*int32)(unsafe.Add(mBase, uint32(v32))) = v36
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v32)+4)) = v38
				v41 = v32 + int32(8)
				if v38 <= v15 {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v45 = int32(base.Ui32(v43) >> (uint(int32(2)) % 32))
					v47 = v45 - int32(8)
					if v43&int32(12)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v47)) == int32(0) {
						v55 = v32 + v45
						if base.Ui32(v55) <= base.Ui32(v41) {
							return v32
						} else {
							v58 = v32 + int32(12)
							if base.Ui32(v58) < base.Ui32(v55) {
								v60 = v55
							} else {
								v60 = v58
							}
							v67 = (v60-v32-int32(9))&int32(-4) + int32(4)
							if v67 == int32(0) {
								return v32
							} else {
								v201 = v67
								base.MemoryFill(m, v41, int32(0), v201)
								return v32
							}
						}
					} else {
						if v47 == int32(0) {
							return v32
						} else {
							v201 = v47
							base.MemoryFill(m, v41, int32(0), v201)
							return v32
						}
					}
				} else {
					v73 = v15 & int32(7)
					v75 = int32(base.Ui32(v15) >> (uint(int32(3)) % 32))
					if v15&int32(24)|base.B2i32(base.Ui32(int32(_a_F_bitshiftright_0)) < base.Ui32(v15)) == int32(0) {
						if v75 == int32(0) {
						} else {
							v87 = v32 + v75 + int32(8)
							v89 = v32 + int32(12)
							if base.Ui32(v89) < base.Ui32(v87) {
								v91 = v87
							} else {
								v91 = v89
							}
							v98 = (v91-v32-int32(9))&int32(-4) + int32(4)
							if v98 == int32(0) {
							} else {
								base.MemoryFill(m, v41, int32(0), v98)
							}
						}
					} else {
						if v75 == int32(0) {
						} else {
							base.MemoryFill(m, v41, int32(0), v75)
						}
					}
					v110 = v11 + int32(8)
					v111 = v75 + v41
					if v73 == int32(0) {
						v114 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						v116 = int32(base.Ui32(v114) >> (uint(int32(2)) % 32))
						v119 = v116 - v75 - int32(8)
						if v119 != 0 {
							base.MemoryCopy(m, v111, v110, v119)
						} else {
						}
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
						v165 = v116 + v32
						v171 = v122
					} else {
						if base.Ui32(int32(base.Ui32(v34)>>(uint(int32(2))%32))) <= base.Ui32(v75+int32(8)) {
							v165 = v111
							v171 = v36
						} else {
							v128 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v111))) = uint8(v128)
							v132 = v111
							v136 = v110
							for {
								v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
								v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
								v144 = v141 | int32(base.Ui32(v142)>>(uint(v73)%32))
								*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v144)
								v147 = v132 + int32(1)
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
								v150 = int32(base.Ui32(v148) >> (uint(int32(2)) % 32))
								if base.Ui32(v147) < base.Ui32(v32+v150) {
									v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
									v154 = v153 << (uint(int32(8)-v73) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v154)
									v156 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
									v159 = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
									v160 = v156
								} else {
									v159 = v150
									v160 = v148
								}
								if base.Ui32(v147) < base.Ui32(v159+v32) {
									v132 = v147
									v136 = v136 + int32(1)
									continue
								} else {
									break
								}
								break
							}
							v165 = v147
							v171 = v160
						}
					}
					v178 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
					v181 = v171<<(uint(int32(1))%32)&int32(-8) - v178 + int32(-64)
					if v181 <= int32(0) {
					} else {
						v185 = v165 - int32(1)
						v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185))))
						v189 = v186 & (int32(255) << (uint(v181) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(v185))) = uint8(v189)
					}
					return v32
				}
			}
		}
	}
}
func F_blgetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v117 int32
	_ = v117
	var v118 int64
	_ = v118
	var v131 int32
	_ = v131
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v200 int64
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v285 int64
	_ = v285
	var v287 int32
	_ = v287
	var v308 int64
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v335 int64
	_ = v335
	var v337 int32
	_ = v337
	v18 = int64(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v94 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L14
	}
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1032))
	v25 = F_palloc0(m, v22<<(uint(int32(1))%32))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int64(0)
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v25
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v30 <= int32(0) {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v37 = v21
	v40 = int32(0)
	goto L6
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	if v54&int32(1) != 0 {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L1
L8:
	;
	F_pfree(m, v53)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L3
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v37)+44))
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v37)+4)))
	F_signValue(m, v19+int32(4), v53, v63, v64-int32(1))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L3
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(0)
	return int64(0)
L12:
	;
	v72 = v40 + int32(1)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v72 < v73 {
		v37 = v37 + int32(48)
		v40 = v72
		goto L6
	} else {
		goto L13
	}
L13:
	;
	goto L7
L14:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v98 = F_RelationGetNumberOfBlocksInFork(m, v96, int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+272))
	if v101 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v117 != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+268)))
	if v104 != int32(1) {
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v111 = v101
	goto L19
L19:
	;
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v111)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v111)+16)) = v112 + int64(1)
	goto L16
L20:
	;
	F_pgstat_assoc_relation(m, v100)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L21
	}
L21:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+272))
	v111 = v110
	goto L19
L22:
	;
	v118 = *(*int64)(unsafe.Add(mBase, uint32(v117)))
	*(*int64)(unsafe.Add(mBase, uint32(v117))) = v118 + int64(1)
	goto L24
L23:
	;
	goto L24
L24:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v98) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v131 = int32(1)
	v142 = v18
	goto L28
L26:
	;
	v335 = v18
	goto L27
L27:
	;
	F_bms_free(m, v94)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L3
	} else {
		goto L58
	}
L28:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v144 = int32(0)
	v146 = F_ReadBufferExtended(m, v143, v144, v131, v144, v94)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L3
	} else {
		goto L30
	}
L29:
	;
	v335 = v308
	goto L27
L30:
	;
	F_LockBuffer(m, v146, int32(1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	if v146 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v146)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L3
	} else {
		goto L52
	}
L33:
	;
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+14)))
	if v169 == int32(0) {
		v308 = v142
		goto L32
	} else {
		goto L37
	}
L34:
	;
	v154 = *(*int32)(unsafe.Add(mBase, _c_F_blgetbitmap[0]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154+(v146^int32(-1))<<(uint(int32(2))%32))))
	v168 = v160
	goto L33
L35:
	;
	goto L36
L36:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _c_F_blgetbitmap[1]))
	v168 = v162 + v146<<(uint(int32(13))%32) + int32(-8192)
	goto L33
L37:
	;
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+16)))
	v173 = v168 + v172
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v173)+2)))
	if v174&int32(2) != 0 {
		v308 = v142
		goto L32
	} else {
		goto L38
	}
L38:
	;
	v177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173))))
	if v177 == int32(0) {
		v308 = v142
		goto L32
	} else {
		goto L39
	}
L39:
	;
	v188 = int32(1)
	v200 = v142
	goto L40
L40:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1168))
	v207 = v168 + int32(24) + v201*(v188&int32(_a_F_blgetbitmap_0)-int32(1))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1032))
	if int32(0) < v208 {
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v308 = v285
	goto L32
L42:
	;
	v287 = v188 + int32(1)
	if base.Ui32(v287&int32(_a_F_blgetbitmap_0)) <= base.Ui32(v177) {
		v188 = v287
		v200 = v285
		goto L40
	} else {
		goto L51
	}
L43:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v217 = int32(0)
	goto L46
L44:
	;
	goto L45
L45:
	;
	v262 = int32(1)
	F_tbm_add_tuples(m, l1, v207, v262, v262)
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L3
	} else {
		goto L50
	}
L46:
	;
	v234 = v217 << (uint(int32(1)) % 32)
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v213+v234))))
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207+int32(6)+v234))))
	if v236&v238 != v236 {
		v285 = v200
		goto L42
	} else {
		goto L48
	}
L47:
	;
	goto L45
L48:
	;
	v242 = v217 + int32(1)
	if v242 != v208 {
		v217 = v242
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v285 = v200 + int64(1)
	goto L42
L51:
	;
	goto L41
L52:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_blgetbitmap[2]))
	if v312 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v316 = v131 + int32(1)
	if v316 != v98 {
		v131 = v316
		v142 = v308
		goto L28
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	goto L29
L58:
	;
	return v335
}
func F_blinsert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v339 int32
	_ = v339
	var v354 int32
	_ = v354
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v372 int32
	_ = v372
	v12 = m.G0
	v14 = v12 - int32(1168)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_blinsert[0]))
	v22 = F_AllocSetContextCreateInternal(m, v17, int32(_a_F_blinsert_0), int32(0), int32(_a_F_blinsert_1), int32(_a_F_blinsert_2))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = int32(_a_F_blinsert_3)
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_blinsert[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_blinsert[0])) = v22
	F_initBloomState(m, v14, l0)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = F_BloomFormTuple(m, v14, l3, l1, l2)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v35 = F_ReadBuffer(m, l0, int32(0))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	F_LockBuffer(m, v35, int32(1))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	if v35 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L1
	} else {
		goto L101
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_blinsert[0])) = v27
	F_MemoryContextDelete(m, v22)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L100
	}
L9:
	;
	F_LockBuffer(m, v35, int32(2))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L44
	}
L10:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+28)))
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+30)))
	if base.Ui32(v58) < base.Ui32(v59) {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_blinsert[1]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v35^int32(-1))<<(uint(int32(2))%32))))
	v57 = v49
	goto L10
L12:
	;
	goto L13
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_blinsert[2]))
	v57 = v51 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L10
L14:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v57+v58<<(uint(int32(2))%32))+168))
	F_LockBuffer(m, v35, int32(0))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_LockBuffer(m, v35, int32(0))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L43
	}
L17:
	;
	v68 = F_ReadBuffer(m, l0, v64)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_LockBuffer(m, v68, int32(2))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v73 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)))
	v102 = v76 + v101
	v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102))))
	v104 = v100 * v103
	v105 = int32(_a_F_blinsert_4) - v104
	if base.Ui32(v100) <= base.Ui32(v105) {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v76 = F_GenericXLogRegisterBuffer(m, v73, v68, int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+14)))
	if v78 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76+v79)+2)))
	if v81&int32(2) == int32(0) {
		goto L20
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v86 = int32(0)
	F_PageInit(m, v76, int32(_a_F_blinsert_1), int32(8))
	mBase = m.M
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)))
	v91 = v76 + v90
	v92 = int32(_a_F_blinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+6)) = uint16(v92)
	*(*uint16)(unsafe.Add(mBase, uint32(v91)+2)) = uint16(v86)
	goto L27
L26:
	;
	goto L25
L27:
	;
	goto L20
L28:
	;
	if base.Ui32(v100) <= base.Ui32(v105) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if v100 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	base.MemoryCopy(m, v76+v104+int32(24), v32, v100)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102))))
	v113 = v111 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v102))) = uint16(v113)
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+1164)))
	v118 = v113*v115 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+12)) = uint16(v118)
	goto L31
L35:
	;
	F_GenericXLogFinish(m, v73)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_pfree(m, v73)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L41
	}
L38:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_ReleaseBuffer(m, v35)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L8
L41:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v136 = v64
	goto L9
L43:
	;
	v136 = int32(-1)
	goto L9
L44:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+28)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+30)))
	if base.Ui32(v143) < base.Ui32(v144) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	F_UnlockReleaseBuffer(m, v35)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L99
	}
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v57+v143<<(uint(int32(2))%32))+168))
	v152 = v143 + base.B2i32(v136 == v149)
	goto L48
L47:
	;
	v152 = v143
	goto L48
L48:
	;
	v154 = v152 & int32(_a_F_blinsert_6)
	v155 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v158 = F_GenericXLogRegisterBuffer(m, v155, v35, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+30)))
	if base.Ui32(v154) < base.Ui32(v160) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v163 = v154
	v164 = v155
	v168 = v158
	goto L54
L52:
	;
	v249 = v155
	v253 = v158
	goto L53
L53:
	;
	v258 = F_BloomNewBuffer(m, l0)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L82
	}
L54:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v168+v163<<(uint(int32(2))%32))+168))
	v177 = F_ReadBuffer(m, l0, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v249 = v240
	v253 = v243
	goto L53
L56:
	;
	F_LockBuffer(m, v177, int32(2))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v183 = F_GenericXLogRegisterBuffer(m, v164, v177, int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+16)))
	v209 = v183 + v208
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209))))
	v211 = v207 * v210
	v212 = int32(_a_F_blinsert_4) - v211
	if base.Ui32(v207) <= base.Ui32(v212) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
	if v185 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+16)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v186)+2)))
	if v188&int32(2) == int32(0) {
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v193 = int32(0)
	F_PageInit(m, v183, int32(_a_F_blinsert_1), int32(8))
	mBase = m.M
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+16)))
	v198 = v183 + v197
	v199 = int32(_a_F_blinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v198)+6)) = uint16(v199)
	*(*uint16)(unsafe.Add(mBase, uint32(v198)+2)) = uint16(v193)
	goto L64
L63:
	;
	goto L62
L64:
	;
	goto L58
L65:
	;
	if base.Ui32(v207) <= base.Ui32(v212) {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	if v207 != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	goto L68
L68:
	;
	goto L65
L69:
	;
	base.MemoryCopy(m, v183+v211+int32(24), v32, v207)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209))))
	v220 = v218 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v209))) = uint16(v220)
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+1164)))
	v225 = v220*v222 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v183)+12)) = uint16(v225)
	goto L68
L72:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v168)+28)) = uint16(v163)
	F_GenericXLogFinish(m, v164)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_pfree(m, v164)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	F_UnlockReleaseBuffer(m, v177)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L45
L77:
	;
	F_UnlockReleaseBuffer(m, v177)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v239 = v163 + int32(1)
	v240 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v243 = F_GenericXLogRegisterBuffer(m, v240, v35, int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243)+30)))
	if base.Ui32(v239) < base.Ui32(v245) {
		v163 = v239
		v164 = v240
		v168 = v243
		goto L54
	} else {
		goto L81
	}
L81:
	;
	goto L55
L82:
	;
	v261 = F_GenericXLogRegisterBuffer(m, v249, v258, int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v263 = int32(0)
	F_PageInit(m, v261, int32(_a_F_blinsert_1), int32(8))
	mBase = m.M
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+16)))
	v268 = v261 + v267
	v269 = int32(_a_F_blinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+6)) = uint16(v269)
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+2)) = uint16(v263)
	goto L84
L84:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+16)))
	v279 = v261 + v278
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279))))
	v281 = v277 * v280
	v282 = int32(_a_F_blinsert_4) - v281
	if base.Ui32(v277) <= base.Ui32(v282) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if base.B2i32(base.Ui32(v277) <= base.Ui32(v282)) == int32(0) {
		goto L7
	} else {
		goto L92
	}
L86:
	;
	if v277 != 0 {
		goto L89
	} else {
		goto L90
	}
L87:
	;
	goto L88
L88:
	;
	goto L85
L89:
	;
	base.MemoryCopy(m, v261+v281+int32(24), v32, v277)
	goto L91
L90:
	;
	goto L91
L91:
	;
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279))))
	v290 = v288 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v279))) = uint16(v290)
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+1164)))
	v295 = v290*v292 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v261)+12)) = uint16(v295)
	goto L88
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+28)) = int32(_a_F_blinsert_7)
	if v258 < int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+168)) = v321
	F_GenericXLogFinish(m, v249)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L97
	}
L94:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_blinsert[3]))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306+(v258^int32(-1))<<(uint(int32(6))%32))+16))
	v321 = v312
	goto L93
L95:
	;
	goto L96
L96:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _c_F_blinsert[4]))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v314+v258<<(uint(int32(6))%32)+int32(-64))+16))
	v321 = v320
	goto L93
L97:
	;
	F_UnlockReleaseBuffer(m, v258)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	goto L45
L99:
	;
	goto L8
L100:
	;
	m.G0 = v14 + int32(1168)
	return int32(0)
L101:
	;
	F_errmsg_internal(m, int32(_a_F_blinsert_8), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_blinsert_9), int32(324), int32(_a_F_blinsert_10))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_blvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v8 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v104 = l1
	goto L3
L3:
	;
	return v104
L4:
	;
	v15 = F_palloc0(m, int32(40))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v19 = l1
	goto L6
L6:
	;
	v21 = F_RelationGetNumberOfBlocksInFork(m, v11, int32(0))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v19 = v15
	goto L6
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = int64(0)
	if base.Ui32(int32(2)) <= base.Ui32(v21) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v33 = int32(1)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_FreeSpaceMapVacuum(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L7
	} else {
		goto L31
	}
L13:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L7
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v41 = int32(0)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v44 = F_ReadBufferExtended(m, v11, v41, v33, v41, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L16
	}
L16:
	;
	F_LockBuffer(m, v44, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if v44 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	F_UnlockReleaseBuffer(m, v44)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L7
	} else {
		goto L29
	}
L19:
	;
	v82 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69))))
	*(*float64)(unsafe.Add(mBase, uint32(v19)+8)) = base.F64_add(v82, base.F64_convert_i32_u(v83))
	goto L18
L20:
	;
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+14)))
	if v67 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_blvacuumcleanup[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52+(v44^int32(-1))<<(uint(int32(2))%32))))
	v66 = v58
	goto L20
L22:
	;
	goto L23
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _c_F_blvacuumcleanup[1]))
	v66 = v60 + v44<<(uint(int32(13))%32) + int32(-8192)
	goto L20
L24:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+16)))
	v69 = v66 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+2)))
	if v70&int32(2) == int32(0) {
		goto L19
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	F_RecordFreeIndexPage(m, v11, v33)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L7
	} else {
		goto L28
	}
L27:
	;
	goto L26
L28:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v78 + int32(1)
	goto L18
L29:
	;
	v91 = v33 + int32(1)
	if v91 != v21 {
		v33 = v91
		goto L13
	} else {
		goto L30
	}
L30:
	;
	goto L14
L31:
	;
	v104 = v19
	goto L3
}
func F_boolor_statefunc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 != 0 {
		v7 = int32(1)
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v7 = base.B2i32(v4 != int32(0))
	}
	return v7
}
func F_boot_yyerror(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12+v13<<(uint(int32(2))%32))))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
		*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v6))) = l1
		F_errmsg_internal(m, int32(_a_F_boot_yyerror_0), v6)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errfinish(m, int32(_a_F_boot_yyerror_1), int32(137), int32(_a_F_boot_yyerror_2))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_bpcharle(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = v12 + int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v25 = v23 & int32(1)
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = v17
	goto L6
L5:
	;
	v26 = v12 + int32(4)
	goto L6
L6:
	;
	if v23 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v58 = v53
	goto L18
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v32 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v35 = int32(16)
	goto L13
L12:
	;
	v35 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = int32(4)
	goto L16
L15:
	;
	v42 = v35
	goto L16
L16:
	;
	v53 = v42
	goto L7
L17:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v58 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v77 = int32(1)
	v78 = v19 + v77
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v83 = v81 & v77
	if v83 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v76 = v53 & (v53 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v70 = v58 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v70))))
	if v72 == int32(32) {
		v58 = v70
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v76 = v58
	goto L20
L25:
	;
	v84 = v78
	goto L27
L26:
	;
	v84 = v19 + int32(4)
	goto L27
L27:
	;
	if v81 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v116 = v111
	goto L39
L29:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v90 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v101 = int32(1)
	if v83 != 0 {
		v111 = int32(base.Ui32(v81)>>(uint(v101)%32)) - v101
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v93 = int32(16)
	goto L34
L33:
	;
	v93 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v90-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v100 = int32(4)
	goto L37
L36:
	;
	v100 = v93
	goto L37
L37:
	;
	v111 = v100
	goto L28
L38:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v111 = int32(base.Ui32(v105)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v116 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v135 = int32(1)
	if v23&v135 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v134 = v111 & (v111 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v128 = v116 - int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v128))))
	if v130 == int32(32) {
		v116 = v128
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v134 = v116
	goto L41
L46:
	;
	v139 = v135
	goto L48
L47:
	;
	v139 = int32(4)
	goto L48
L48:
	;
	v141 = int32(1)
	if v81&v141 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v145 = v141
	goto L51
L50:
	;
	v145 = int32(4)
	goto L51
L51:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v148 = F_varstr_cmp(m, v12+v139, v76, v19+v145, v134, v147)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v150 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v154 != v19 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v19)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return base.B2i32(v148 <= int32(0))
L60:
	;
	goto L59
}
func F_bpcharne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
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
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum_packed(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v16 = F_pg_detoast_datum_packed(m, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v18 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v19 = int32(1)
	v20 = v11 + v19
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v25 = v23 & v19
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L94
	}
L7:
	;
	v26 = v20
	goto L9
L8:
	;
	v26 = v11 + int32(4)
	goto L9
L9:
	;
	if v23 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v58 = v53
	goto L21
L11:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v32 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v43 = int32(1)
	if v25 != 0 {
		v53 = int32(base.Ui32(v23)>>(uint(v43)%32)) - v43
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v35 = int32(16)
	goto L16
L15:
	;
	v35 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v42 = int32(4)
	goto L19
L18:
	;
	v42 = v35
	goto L19
L19:
	;
	v53 = v42
	goto L10
L20:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v53 = int32(base.Ui32(v47)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	if v58 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v76 = int32(1)
	v77 = v16 + v76
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v82 = v80 & v76
	if v82 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v75 = v53 & (v53 >> (uint(int32(31)) % 32))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v69 = v58 - int32(1)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v69))))
	if v71 == int32(32) {
		v58 = v69
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v75 = v58
	goto L23
L28:
	;
	v83 = v77
	goto L30
L29:
	;
	v83 = v16 + int32(4)
	goto L30
L30:
	;
	if v80 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v115 = v110
	goto L42
L32:
	;
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	if v89 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v100 = int32(1)
	if v82 != 0 {
		v110 = int32(base.Ui32(v80)>>(uint(v100)%32)) - v100
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v92 = int32(16)
	goto L37
L36:
	;
	v92 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v89-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v99 = int32(4)
	goto L40
L39:
	;
	v99 = v92
	goto L40
L40:
	;
	v110 = v99
	goto L31
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	if v115 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v134 = F_pg_newlocale_from_collation(m, v18)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	goto L43
L45:
	;
	v132 = v110 & (v110 >> (uint(int32(31)) % 32))
	goto L44
L46:
	;
	goto L47
L47:
	;
	v126 = v115 - int32(1)
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83+v126))))
	if v128 == int32(32) {
		v115 = v126
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v132 = v115
	goto L44
L49:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v237 != v11 {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	if v136 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v132 != v75 {
		v236 = int32(1)
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v218 = int32(1)
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v220&v218 != 0 {
		goto L79
	} else {
		goto L80
	}
L54:
	;
	v140 = int32(1)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v142&v140 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v145 = v140
	goto L57
L56:
	;
	v145 = int32(4)
	goto L57
L57:
	;
	v146 = v11 + v145
	v147 = int32(1)
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v149&v147 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v152 = v147
	goto L60
L59:
	;
	v152 = int32(4)
	goto L60
L60:
	;
	v153 = v16 + v152
	if base.Ui32(int32(4)) <= base.Ui32(v75) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v236 = base.B2i32(v215 != int32(0))
	goto L49
L62:
	;
	v215 = int32(0)
	goto L61
L63:
	;
	v189 = v184
	v190 = v185
	v191 = v186
	goto L73
L64:
	;
	if (v146|v153)&int32(3) != 0 {
		v184 = v146
		v185 = v153
		v186 = v75
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v177 = v146
	v178 = v153
	v179 = v75
	goto L66
L66:
	;
	if v179 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v161 = v146
	v162 = v153
	v163 = v75
	goto L68
L68:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v166 != v167 {
		v184 = v161
		v185 = v162
		v186 = v163
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v177 = v172
	v178 = v170
	v179 = v174
	goto L66
L70:
	;
	v169 = int32(4)
	v170 = v162 + v169
	v172 = v161 + v169
	v174 = v163 - v169
	if base.Ui32(int32(3)) < base.Ui32(v174) {
		v161 = v172
		v162 = v170
		v163 = v174
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v184 = v177
	v185 = v178
	v186 = v179
	goto L63
L73:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190))))
	if v194 == v195 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v215 = v194 - v195
	goto L61
L75:
	;
	v197 = int32(1)
	v202 = v191 - v197
	if v202 != 0 {
		v189 = v189 + v197
		v190 = v190 + v197
		v191 = v202
		goto L73
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	goto L74
L78:
	;
	goto L62
L79:
	;
	v223 = v218
	goto L81
L80:
	;
	v223 = int32(4)
	goto L81
L81:
	;
	v225 = int32(1)
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v227&v225 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v230 = v225
	goto L84
L83:
	;
	v230 = int32(4)
	goto L84
L84:
	;
	v232 = F_varstr_cmp(m, v11+v223, v75, v16+v230, v132, v18)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v236 = base.B2i32(v232 != int32(0))
	goto L49
L86:
	;
	F_pfree(m, v11)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v241 != v16 {
		goto L90
	} else {
		goto L91
	}
L89:
	;
	goto L88
L90:
	;
	F_pfree(m, v16)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	return v236
L93:
	;
	goto L92
L94:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_bpcharne_0), int32(0))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errhint(m, int32(_a_F_bpcharne_1), int32(0))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_bpcharne_2), int32(738), int32(_a_F_bpcharne_3))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_brinbuild(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int64
	_ = v149
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int64
	_ = v167
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int64
	_ = v329
	var v330 int64
	_ = v330
	var v333 int32
	_ = v333
	var v338 int64
	_ = v338
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v477 int32
	_ = v477
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v491 int32
	_ = v491
	var v492 float64
	_ = v492
	var v494 float64
	_ = v494
	var v499 int32
	_ = v499
	var v500 float64
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v552 int32
	_ = v552
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 float64
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 float64
	_ = v737
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v761 float64
	_ = v761
	var v764 float64
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	v18 = m.G0
	v20 = v18 - int32(48)
	m.G0 = v20
	v23 = F_RelationGetNumberOfBlocksInFork(m, l1, int32(0))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v764 = *(*float64)(unsafe.Add(mBase, uint32(v147)+8))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	F_brinRevmapTerminate(m, v765)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L3
	} else {
		goto L173
	}
L2:
	;
	v710 = int32(0)
	v717 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v717)+140))
	v719 = m.T0[v718].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v710, v710, int32(1), v710, int32(-1), int32(14), v147, v710)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L3
	} else {
		goto L168
	}
L3:
	;
	return int32(0)
L4:
	;
	if v23 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = l1
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v20)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v32
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v34
	v36 = int32(0)
	v39 = F_ExtendBufferedRel(m, v20, v36, v36, int32(9))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L3
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L3
	} else {
		goto L165
	}
L8:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v59 != 0 {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	if v39 < int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[0]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(v39^int32(-1))<<(uint(int32(2))%32))))
	v58 = v50
	goto L8
L11:
	;
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[1]))
	v58 = v52 + v39<<(uint(int32(13))%32) + int32(-8192)
	goto L8
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
	v62 = v60
	goto L15
L14:
	;
	v62 = int32(128)
	goto L15
L15:
	;
	F_PageInit(m, v58, int32(_a_F_brinbuild_0), int32(8))
	mBase = m.M
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
	v69 = int32(_a_F_brinbuild_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v58+v67)+6)) = uint16(v69)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+36)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+32)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = int32(-1475306246)
	v77 = int32(40)
	*(*uint16)(unsafe.Add(mBase, uint32(v58)+12)) = uint16(v77)
	goto L16
L16:
	;
	F_MarkBufferDirty(m, v39)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L3
	} else {
		goto L17
	}
L17:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+118)))
	if v82 != int32(112) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	F_UnlockReleaseBuffer(m, v39)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L3
	} else {
		goto L36
	}
L19:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[2]))
	if v86 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v89 != 0 {
		goto L18
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v91 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+24)) = uint16(v91)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+180))
	if v93 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v90 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	v96 = v94
	goto L27
L26:
	;
	v96 = int32(128)
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v96
	F_XLogBeginInsert(m)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	F_XLogRegisterData(m, v20+int32(20), int32(6))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	F_XLogRegisterBuffer(m, int32(0), v39, int32(14))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v111 = F_XLogInsert(m, int32(17), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L3
	} else {
		goto L31
	}
L31:
	;
	if v39 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v130))) = base.I64_rotr(v111, int64(32))
	goto L18
L33:
	;
	v116 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[0]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+(v39^int32(-1))<<(uint(int32(2))%32))))
	v130 = v122
	goto L32
L34:
	;
	goto L35
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[1]))
	v130 = v124 + v39<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L36:
	;
	v140 = F_brinRevmapInitialize(m, l1, v20+int32(40))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L3
	} else {
		goto L37
	}
L37:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v20)+40))
	v144 = F_RelationGetNumberOfBlocksInFork(m, l0, int32(0))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v147 = F_palloc(m, int32(80))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v149 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v147)+8)) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v147))) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v147)+16)) = v149
	v154 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v147)+40)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(v147)+32)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v147)+28)) = v142
	v160 = F_brin_build_desc(m, l1)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+44)) = v160
	v163 = F_brin_new_memtuple(m, v160)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+72)) = int32(0)
	v167 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v147)+64)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v147)+48)) = v163
	v171 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v147)+52)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v147)+60)) = v171
	if v144 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v176 = v144 - int32(1)
	v177 = base.I32_rem_u_s(v176, v142)
	v181 = v176 - v177
	goto L44
L43:
	;
	v181 = int32(0)
	goto L44
L44:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v147)+36)) = v181 + v182
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v185 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	if v434 == int32(0) {
		goto L2
	} else {
		goto L113
	}
L46:
	;
	v189 = v185 + int32(1)
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+121)))
	v192 = F_palloc0(m, int32(28))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L3
	} else {
		goto L47
	}
L47:
	;
	v196 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[4]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+72)) = v197 + int32(1)
	goto L48
L48:
	;
	v204 = F_CreateParallelContext(m, int32(_a_F_brinbuild_2), int32(_a_F_brinbuild_3), v185)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L3
	} else {
		goto L49
	}
L49:
	;
	if v190 == int32(1) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v208 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L3
	} else {
		goto L53
	}
L51:
	;
	v212 = int32(_a_F_brinbuild_4)
	goto L52
L52:
	;
	v214 = F_table_parallelscan_estimate(m, l0, v212)
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L3
	} else {
		goto L55
	}
L53:
	;
	v210 = F_RegisterSnapshot(m, v208)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v212 = v210
	goto L52
L55:
	;
	v216 = F_add_size(m, int32(96), v214)
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L3
	} else {
		goto L56
	}
L56:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
	v223 = F_add_size(m, v218, (v216+int32(31))&int32(-32))
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L3
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+36)) = v223
	v226 = F_tuplesort_estimate_shared(m, v189)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L3
	} else {
		goto L58
	}
L58:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
	v233 = F_add_size(m, v228, (v226+int32(31))&int32(-32))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L3
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+36)) = v233
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v204)+40))
	v238 = F_add_size(m, v236, int32(2))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L3
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+40)) = v238
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v244 = F_mul_size(m, int32(32), v243)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L3
	} else {
		goto L61
	}
L61:
	;
	v250 = F_add_size(m, v241, (v244+int32(31))&int32(-32))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L3
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+36)) = v250
	v253 = int32(1)
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v204)+40))
	v256 = F_add_size(m, v254, v253)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L3
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+40)) = v256
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v262 = F_mul_size(m, int32(128), v261)
	mBase = m.M
	v263 = m.ExcPending
	if v263 != 0 {
		goto L3
	} else {
		goto L64
	}
L64:
	;
	v268 = F_add_size(m, v259, (v262+int32(31))&int32(-32))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L3
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+36)) = v268
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v204)+40))
	v273 = F_add_size(m, v271, int32(1))
	mBase = m.M
	v274 = m.ExcPending
	if v274 != 0 {
		goto L3
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+40)) = v273
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[5]))
	if v277 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v204)+36))
	v279 = F_strlen(m, v277)
	mBase = m.M
	v284 = F_add_size(m, v278, v279&int32(-32)+int32(32))
	mBase = m.M
	v285 = m.ExcPending
	if v285 != 0 {
		goto L3
	} else {
		goto L70
	}
L68:
	;
	v294 = v253
	goto L69
L69:
	;
	F_InitializeParallelDSM(m, v204)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L3
	} else {
		goto L72
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+36)) = v284
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v204)+40))
	v289 = F_add_size(m, v287, int32(1))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L3
	} else {
		goto L71
	}
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v204)+40)) = v289
	v294 = v279 + int32(1)
	goto L69
L72:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v204)+44))
	if v297 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v212)))
	switch v300 {
	case 0, 5:
		goto L77
	default:
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v313 = F_shm_toc_allocate(m, v312, v216)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L3
	} else {
		goto L81
	}
L76:
	;
	F_DestroyParallelContext(m, v204)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L3
	} else {
		goto L79
	}
L77:
	;
	F_UnregisterSnapshot(m, v212)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L3
	} else {
		goto L78
	}
L78:
	;
	goto L76
L79:
	;
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[4]))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v307)+72)) = v308 - int32(1)
	goto L80
L80:
	;
	goto L45
L81:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v313))) = v315
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v313)+16)) = v189
	*(*uint8)(unsafe.Add(mBase, uint32(v313)+8)) = uint8(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v313)+4)) = v317
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v313)+12)) = v321
	v325 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[6]))
	if v325 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v313)+24)) = v330
	v333 = v313 + int32(32)
	*(*int32)(unsafe.Add(mBase, uint32(v333)+8)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v333))) = int64(-4294967296)
	goto L86
L83:
	;
	v330 = int64(0)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v329 = *(*int64)(unsafe.Add(mBase, uint32(v325)+392))
	v330 = v329
	goto L82
L86:
	;
	v338 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v313)+56)) = v338
	*(*int64)(unsafe.Add(mBase, uint32(v313)+44)) = v338
	*(*int64)(unsafe.Add(mBase, uint32(v313)+64)) = v338
	F_table_parallelscan_initialize(m, l0, v313+int32(96), v212)
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L3
	} else {
		goto L87
	}
L87:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v349 = F_shm_toc_allocate(m, v348, v226)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v204)+44))
	F_tuplesort_initialize_shared(m, v349, v189, v351)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v354, int64(-5764607523034234879), v313)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L3
	} else {
		goto L90
	}
L90:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v358, int64(-5764607523034234878), v349)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[5]))
	if v363 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v365 = F_shm_toc_allocate(m, v364, v294)
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L3
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v378 = F_mul_size(m, int32(32), v377)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L3
	} else {
		goto L100
	}
L95:
	;
	if v294 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[5]))
	base.MemoryCopy(m, v365, v368, v294)
	goto L98
L97:
	;
	goto L98
L98:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v370, int64(-5764607523034234877), v365)
	mBase = m.M
	v373 = m.ExcPending
	if v373 != 0 {
		goto L3
	} else {
		goto L99
	}
L99:
	;
	goto L94
L100:
	;
	v380 = F_shm_toc_allocate(m, v375, v378)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L3
	} else {
		goto L101
	}
L101:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v382, int64(-5764607523034234876), v380)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v389 = F_mul_size(m, int32(128), v388)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L3
	} else {
		goto L103
	}
L103:
	;
	v391 = F_shm_toc_allocate(m, v386, v389)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v393, int64(-5764607523034234875), v391)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	F_LaunchParallelWorkers(m, v204)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v204
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = v391
	*(*int32)(unsafe.Add(mBase, uint32(v192)+20)) = v380
	*(*int32)(unsafe.Add(mBase, uint32(v192)+16)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v192)+12)) = v349
	*(*int32)(unsafe.Add(mBase, uint32(v192)+8)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v400 + int32(1)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	if v409 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	F__brin_end_parallel(m, v192)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L3
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+64)) = v192
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v192)+12))
	v418 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[7]))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v420 = base.I32_div_s(v418, v419)
	F__brin_parallel_scan_and_build(m, v147, v415, v416, l0, l1, v420)
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L3
	} else {
		goto L111
	}
L110:
	;
	goto L45
L111:
	;
	F_WaitForParallelWorkersToAttach(m, v204)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L3
	} else {
		goto L112
	}
L112:
	;
	goto L45
L113:
	;
	v438 = F_palloc0(m, int32(12))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L3
	} else {
		goto L114
	}
L114:
	;
	v440 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v438))) = uint8(v440)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v442)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+4)) = v443
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+8)) = v446
	v449 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[7]))
	v450 = F_tuplesort_begin_index_brin(m, v449, v438)
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+72)) = v450
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v453)+8))
	v458 = v454 + int32(44)
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v453)+4))
	goto L116
L116:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v458)))
	*(*int32)(unsafe.Add(mBase, uint32(v458))) = int32(1)
	if v477 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	v492 = *(*float64)(unsafe.Add(mBase, uint32(v454)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v147)+16)) = v492
	v494 = *(*float64)(unsafe.Add(mBase, uint32(v454)+64))
	*(*float64)(unsafe.Add(mBase, uint32(v147)+8)) = v494
	*(*int32)(unsafe.Add(mBase, uint32(v454)+44)) = int32(0)
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L3
	} else {
		goto L126
	}
L118:
	;
	F_s_lock(m, v458, int32(_a_F_brinbuild_5), int32(2587), int32(_a_F_brinbuild_6))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L3
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v454)+48))
	if v459 != v485 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v458))) = int32(0)
	F_ConditionVariableSleep(m, v454+int32(32), int32(134217767))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L3
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	goto L117
L125:
	;
	goto L116
L126:
	;
	v500 = *(*float64)(unsafe.Add(mBase, uint32(v147)+16))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	F_tuplesort_performsort(m, v501)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v505 = F_brin_new_memtuple(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[3]))
	v513 = F_AllocSetContextCreateInternal(m, v508, int32(_a_F_brinbuild_7), int32(0), int32(_a_F_brinbuild_0), int32(_a_F_brinbuild_8))
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	v515 = int32(_a_F_brinbuild_9)
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_brinbuild[3])) = v513
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v522 = F_tuplesort_getbrintuple(m, v519, v20+int32(20))
	mBase = m.M
	v523 = m.ExcPending
	if v523 != 0 {
		goto L3
	} else {
		goto L131
	}
L130:
	;
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v147)+36))
	F_brin_fill_empty_ranges(m, v147, v681, v682)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L3
	} else {
		goto L162
	}
L131:
	;
	if v522 == int32(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	F_tuplesort_end(m, v526)
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L3
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	v531 = v147 + int32(24)
	v533 = v522
	v534 = v505
	v535 = int32(-1)
	goto L137
L135:
	;
	v681 = int32(-1)
	goto L130
L136:
	;
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v633)+4))
	v653 = F_brin_form_tuple(m, v649, v650, v633, v20+int32(44))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L3
	} else {
		goto L159
	}
L137:
	;
	if v535 != int32(-1) {
		goto L139
	} else {
		goto L140
	}
L138:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	F_tuplesort_end(m, v626)
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L3
	} else {
		goto L157
	}
L139:
	;
	v552 = v533
	goto L143
L140:
	;
	v597 = v533
	goto L141
L141:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v615 = F_brin_deform_tuple(m, v614, v597, v534)
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L3
	} else {
		goto L153
	}
L142:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v586 = F_brin_form_tuple(m, v583, v569, v534, v20+int32(44))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L3
	} else {
		goto L150
	}
L143:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v534)+4))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	if v569 != v570 {
		goto L142
	} else {
		goto L145
	}
L144:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	F_tuplesort_end(m, v580)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L3
	} else {
		goto L149
	}
L145:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	F_union_tuples(m, v572, v534, v552)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L3
	} else {
		goto L146
	}
L146:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v578 = F_tuplesort_getbrintuple(m, v575, v20+int32(20))
	mBase = m.M
	v579 = m.ExcPending
	if v579 != 0 {
		goto L3
	} else {
		goto L147
	}
L147:
	;
	if v578 != 0 {
		v552 = v578
		goto L143
	} else {
		goto L148
	}
L148:
	;
	goto L144
L149:
	;
	v633 = v534
	v634 = v535
	goto L136
L150:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v586)))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v593 = F_brin_doinsert(m, v588, v589, v590, v531, v591, v586, v592)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L3
	} else {
		goto L151
	}
L151:
	;
	F_MemoryContextReset(m, v513)
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L3
	} else {
		goto L152
	}
L152:
	;
	v597 = v552
	goto L141
L153:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	F_brin_fill_empty_ranges(m, v147, v535, v617)
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L3
	} else {
		goto L154
	}
L154:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v597)))
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v624 = F_tuplesort_getbrintuple(m, v621, v20+int32(20))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L3
	} else {
		goto L155
	}
L155:
	;
	if v624 != 0 {
		v533 = v624
		v534 = v615
		v535 = v620
		goto L137
	} else {
		goto L156
	}
L156:
	;
	goto L138
L157:
	;
	v629 = int32(-1)
	if v620 == v629 {
		v681 = v629
		goto L130
	} else {
		goto L158
	}
L158:
	;
	v633 = v615
	v634 = v620
	goto L136
L159:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v653)))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v660 = F_brin_doinsert(m, v655, v656, v657, v531, v658, v653, v659)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L3
	} else {
		goto L160
	}
L160:
	;
	F_pfree(m, v653)
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L3
	} else {
		goto L161
	}
L161:
	;
	v681 = v634
	goto L130
L162:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brinbuild[3])) = v516
	F_MemoryContextDelete(m, v513)
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	F__brin_end_parallel(m, v689)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	v761 = v500
	goto L1
L165:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v696 + int32(4)
	F_errmsg_internal(m, int32(_a_F_brinbuild_10), v20+int32(16))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L3
	} else {
		goto L166
	}
L166:
	;
	F_errfinish(m, int32(_a_F_brinbuild_5), int32(1120), int32(_a_F_brinbuild_11))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L168:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v147)+32))
	v723 = *(*int32)(unsafe.Add(mBase, uint32(v147)+48))
	v726 = F_brin_form_tuple(m, v721, v722, v723, v20+int32(20))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	v733 = *(*int32)(unsafe.Add(mBase, uint32(v147)+32))
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v735 = F_brin_doinsert(m, v728, v729, v730, v147+int32(24), v733, v726, v734)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L3
	} else {
		goto L170
	}
L170:
	;
	v737 = *(*float64)(unsafe.Add(mBase, uint32(v147)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v147)+8)) = base.F64_add(v737, float64(1))
	F_pfree(m, v726)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L3
	} else {
		goto L171
	}
L171:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v147)+32))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v147)+36))
	F_brin_fill_empty_ranges(m, v147, v743, v744)
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L3
	} else {
		goto L172
	}
L172:
	;
	v761 = v719
	goto L1
L173:
	;
	F_terminate_brin_buildstate(m, v147)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L3
	} else {
		goto L174
	}
L174:
	;
	v771 = F_palloc(m, int32(16))
	mBase = m.M
	v772 = m.ExcPending
	if v772 != 0 {
		goto L3
	} else {
		goto L175
	}
L175:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v771)+8)) = v764
	*(*float64)(unsafe.Add(mBase, uint32(v771))) = v761
	m.G0 = v20 + int32(48)
	return v771
}
func F_btarraycmp(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_btint2cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return v2 - v3
}
func F_btint2skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(195)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(196)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(140737488322560)
	return int32(0)
}
func F_btint42cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+28)))
	return base.B2i32(v4 < v3) - base.B2i32(v3 < v4)
}
func F_btint4cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v4 < v3) - base.B2i32(v3 < v4)
}
func F_btint4skipsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+12)) = int32(198)
	*(*int32)(unsafe.Add(mBase, uint32(v2)+8)) = int32(199)
	*(*int64)(unsafe.Add(mBase, uint32(v2))) = int64(9223372034707292160)
	return int32(0)
}
func F_btoidsortsupport(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(203)
	return int32(0)
}
func F_btoidvectorcmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_check_valid_oidvector(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_check_valid_oidvector(m, v7)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
	if v15 == v16 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v24 = int32(24)
	v30 = int32(0)
	goto L10
L5:
	;
	if int32(0) < v15 {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	return v15 - v16
L8:
	;
	return int32(0)
L9:
	;
	if base.Ui32(v40) < base.Ui32(v38) {
		goto L14
	} else {
		goto L15
	}
L10:
	;
	v36 = v30 << (uint(int32(2)) % 32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v8+v24+v36)))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v7+v24+v36)))
	if v38 != v40 {
		goto L9
	} else {
		goto L12
	}
L11:
	;
	return int32(0)
L12:
	;
	v43 = v30 + int32(1)
	if v43 != v15 {
		v30 = v43
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v50 = int32(1)
	goto L16
L15:
	;
	v50 = int32(-1)
	goto L16
L16:
	;
	return v50
}
func F_btrim1(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = Fn13856(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_bttext_pattern_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
			if v17 == int32(1) {
				v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
				if v23 == int32(18) {
					v26 = int32(16)
				} else {
					v26 = int32(0)
				}
				if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v33 = int32(4)
				} else {
					v33 = v26
				}
				v46 = v33
			} else {
				v34 = int32(1)
				if v17&v34 != 0 {
					v46 = int32(base.Ui32(v17)>>(uint(v34)%32)) - v34
				} else {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v47 == int32(1) {
				v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v53 == int32(18) {
					v56 = int32(16)
				} else {
					v56 = int32(0)
				}
				if base.Ui32((v53-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v63 = int32(4)
				} else {
					v63 = v56
				}
				v76 = v63
			} else {
				v64 = int32(1)
				if v47&v64 != 0 {
					v76 = int32(base.Ui32(v47)>>(uint(v64)%32)) - v64
				} else {
					v70 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v76 = int32(base.Ui32(v70)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v77 = int32(1)
			if v17&v77 != 0 {
				v81 = v77
			} else {
				v81 = int32(4)
			}
			v83 = int32(1)
			if v47&v83 != 0 {
				v87 = v83
			} else {
				v87 = int32(4)
			}
			v89 = base.B2i32(v46 < v76)
			if v46 < v76 {
				v90 = v46
			} else {
				v90 = v76
			}
			v91 = F_memcmp(m, v6+v81, v11+v87, v90)
			mBase = m.M
			if v91 != 0 {
				v94 = v91
			} else {
				if v46 < v76 {
					v94 = int32(-1)
				} else {
					v94 = base.B2i32(v76 < v46)
				}
			}
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v95 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v99 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int32(0)
						} else {
							return v94
						}
					} else {
						return v94
					}
				}
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v99 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int32(0)
					} else {
						return v94
					}
				} else {
					return v94
				}
			}
		}
	}
}
func F_bttranslatestrategy(m *base.Module, l0 int32, l1 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	v3 = int32(1)
	v6 = (l0 - v3) & int32(_a_F_bttranslatestrategy_0)
	if base.Ui32(v6) < base.Ui32(int32(5)) {
		v12 = v6 + v3
	} else {
		v12 = int32(0)
	}
	return v12
}
func F_build_aggregate_finalfn_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v17 = F_palloc0(m, int32(28))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v19 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v17
	*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v17
	v34 = F_list_make1_impl(m, int32(1), v14+int32(8))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = l1 - int32(1)
	if int32(0) < v37 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v43 = int32(0)
	v48 = v34
	goto L7
L5:
	;
	v81 = v34
	goto L6
L6:
	;
	v86 = F_makeFuncExpr(m, l5, l3, v81, l4, int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0+v43<<(uint(int32(2))%32))))
	v57 = F_palloc0(m, int32(28))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v81 = v69
	goto L6
L9:
	;
	v59 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v57)+16)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v57)+12)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v59
	*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(4294967304)
	v69 = F_lappend(m, v48, v57)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v72 = v43 + int32(1)
	if v72 != v37 {
		v43 = v72
		v48 = v69
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v86
	m.G0 = v14 + int32(16)
	return
}
func F_build_aggregate_transfn_expr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	v4 = l3
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = F_palloc0(m, int32(28))
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
	v22 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = l4
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = int64(4294967304)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v20
	v37 = F_list_make1_impl(m, int32(1), v17+int32(8))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l2 < l1 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v42 = l2
	v52 = v37
	goto L7
L5:
	;
	v88 = v37
	goto L6
L6:
	;
	v90 = int32(0)
	v92 = F_makeFuncExpr(m, l6, l4, v88, l5, v90)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0+v42<<(uint(int32(2))%32))))
	v59 = F_palloc0(m, int32(28))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L9
	}
L8:
	;
	v88 = v71
	goto L6
L9:
	;
	v61 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v59)+24)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v59)+20)) = l5
	*(*int32)(unsafe.Add(mBase, uint32(v59)+16)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v59)+12)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(v59))) = int64(4294967304)
	v71 = F_lappend(m, v52, v59)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v74 = v42 + int32(1)
	if v74 != l1 {
		v42 = v74
		v52 = v71
		goto L7
	} else {
		goto L11
	}
L11:
	;
	goto L8
L12:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v92)+13)) = uint8(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v92
	if l9 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	if l7 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L15
L15:
	;
	m.G0 = v17 + int32(16)
	return
L16:
	;
	v97 = F_makeFuncExpr(m, l7, l4, v88, l5, int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	v100 = v90
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = v100
	goto L15
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v97)+13)) = uint8(v4)
	v100 = v97
	goto L18
}
func F_build_minmax_path(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v245 int32
	_ = v245
	var v249 int32
	_ = v249
	var v258 int32
	_ = v258
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int64
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v302 float64
	_ = v302
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 float64
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 float64
	_ = v332
	var v333 float64
	_ = v333
	var v336 int32
	_ = v336
	var v337 float64
	_ = v337
	var v338 float64
	_ = v338
	var v340 float64
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v363 float64
	_ = v363
	var v367 int32
	_ = v367
	var v368 float64
	_ = v368
	var v369 float64
	_ = v369
	var v372 int32
	_ = v372
	var v379 int32
	_ = v379
	var v385 float64
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 float64
	_ = v410
	var v413 float64
	_ = v413
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v455 int32
	_ = v455
	var v459 int32
	_ = v459
	var v460 float64
	_ = v460
	var v463 float64
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 float64
	_ = v492
	var v493 float64
	_ = v493
	var v498 float64
	_ = v498
	var v499 int32
	_ = v499
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v538 int32
	_ = v538
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v563 int32
	_ = v563
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v598 int32
	_ = v598
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 float64
	_ = v638
	var v639 float64
	_ = v639
	v5 = l4
	v6 = l5
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = F_palloc(m, int32(384))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	base.MemoryCopy(m, v20, l0, int32(384))
	v26 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+328)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+20)) = v26
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l0
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v33 + int32(1)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v38 = F_copyObjectImpl(m, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v38
	v41 = int32(1)
	F_IncrementVarSublevelsUp(m, v38, v41, v41)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v46 = F_copyObjectImpl(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v46
	v49 = int32(1)
	F_IncrementVarSublevelsUp(m, v46, v49, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v54 = F_copyObjectImpl(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v58 = F_pstrdup(m, int32(_a_F_build_minmax_path_0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v61 = F_makeTargetEntry(m, v54, int32(1), v58, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v61
	v68 = F_list_make1_impl(m, int32(1), v17+int32(4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+76)) = v68
	*(*int32)(unsafe.Add(mBase, uint32(v20)+264)) = v68
	v72 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+112)) = v72
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+318)) = uint8(v72)
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+40)) = uint8(v72)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+120)) = v72
	*(*uint8)(unsafe.Add(mBase, uint32(v38)+36)) = uint8(v72)
	v83 = F_palloc0(m, int32(20))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = int32(52)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v90 = F_copyObjectImpl(m, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v83)+16)) = int32(-1)
	v94 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v83)+12)) = uint8(v94)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+4)) = v90
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	v99 = F_list_member(m, v98, v83)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v99 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v103)+8))
	v105 = F_lcons(m, v83, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v111 = F_palloc0(m, int32(20))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L18
	}
L17:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v38)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+8)) = v105
	goto L16
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v111))) = int32(106)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v116 = int32(0)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v61)+16))
	if v125 == v116 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v258 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+18)) = uint8(v258)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+17)) = uint8(v6)
	*(*uint8)(unsafe.Add(mBase, uint32(v111)+16)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v111)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v111)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v111)+4)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v111
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v111
	v268 = F_list_make1_impl(m, int32(1), v17)
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L55
	}
L20:
	;
	if v115 == int32(0) {
		v245 = int32(1)
		goto L23
	} else {
		goto L24
	}
L21:
	;
	v249 = v125
	goto L22
L22:
	;
	goto L19
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v245
	v249 = v245
	goto L22
L24:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	if v132 <= int32(0) {
		v245 = int32(1)
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v135 = int32(0)
	if v135 < v132 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v138 = v132
	goto L28
L27:
	;
	v138 = v135
	goto L28
L28:
	;
	v140 = v138 & int32(3)
	v141 = int32(0)
	if int32(4) <= v132 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v245 = v223 + int32(1)
	goto L23
L30:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v150 = v141
	v151 = int32(0)
	v152 = v116
	goto L33
L31:
	;
	v187 = v141
	v189 = v116
	goto L32
L32:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v198 = int32(0)
	v200 = v187
	v202 = v189
	goto L49
L33:
	;
	v161 = v146 + v152<<(uint(int32(2))%32)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)+12))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+16))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+16))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+16))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v168)+16))
	if base.Ui32(v150) < base.Ui32(v169) {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v140 == int32(0) {
		v223 = v177
		goto L29
	} else {
		goto L48
	}
L35:
	;
	v171 = v169
	goto L37
L36:
	;
	v171 = v150
	goto L37
L37:
	;
	if base.Ui32(v171) < base.Ui32(v167) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v173 = v167
	goto L40
L39:
	;
	v173 = v171
	goto L40
L40:
	;
	if base.Ui32(v173) < base.Ui32(v165) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v175 = v165
	goto L43
L42:
	;
	v175 = v173
	goto L43
L43:
	;
	if base.Ui32(v175) < base.Ui32(v163) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v177 = v163
	goto L46
L45:
	;
	v177 = v175
	goto L46
L46:
	;
	v178 = int32(4)
	v179 = v152 + v178
	v181 = v151 + v178
	if v181 != v138&int32(2147483644) {
		v150 = v177
		v151 = v181
		v152 = v179
		goto L33
	} else {
		goto L47
	}
L47:
	;
	goto L34
L48:
	;
	v187 = v177
	v189 = v179
	goto L32
L49:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v196+v202<<(uint(int32(2))%32))))
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v212)+16))
	if base.Ui32(v200) < base.Ui32(v213) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	v223 = v215
	goto L29
L51:
	;
	v215 = v213
	goto L53
L52:
	;
	v215 = v200
	goto L53
L53:
	;
	v216 = int32(1)
	v219 = v198 + v216
	if v219 != v140 {
		v198 = v219
		v200 = v215
		v202 = v202 + v216
		goto L49
	} else {
		goto L54
	}
L54:
	;
	goto L50
L55:
	;
	v270 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v38)+128)) = v270
	*(*int32)(unsafe.Add(mBase, uint32(v38)+124)) = v268
	v278 = F_Int64GetDatum(m, int64(1))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L56
	}
L56:
	;
	v280 = int32(0)
	v282 = F_makeConst(m, int32(20), int32(-1), v270, int32(8), v278, v280, v280)
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v38)+132)) = v282
	v285 = int64(4607182418800017408)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+304)) = v285
	*(*int64)(unsafe.Add(mBase, uint32(v20)+296)) = v285
	v291 = F_query_planner(m, v20, int32(831), int32(0))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_SS_identify_outer_params(m, v20)
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v295 = int32(0)
	v302 = float64(0)
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	if v303 == v295 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v291)+32))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v20)+156))
	v492 = float64(1)
	v493 = *(*float64)(unsafe.Add(mBase, uint32(v291)+16))
	if base.F64_gt(v493, v492) != 0 {
		goto L90
	} else {
		goto L91
	}
L61:
	;
	goto L60
L62:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v303)+4))
	if v306 <= int32(0) {
		v379 = v295
		v385 = v302
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v291)+32))
	if v386 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L64:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v303)+12))
	if v306 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v309+v354<<(uint(int32(2))%32))))
	v368 = *(*float64)(unsafe.Add(mBase, uint32(v367)+56))
	v369 = *(*float64)(unsafe.Add(mBase, uint32(v367)+64))
	v372 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v367)+38)))
	v379 = v372 ^ int32(1) | v357
	v385 = base.F64_add(v363, base.F64_add(v368, v369))
	goto L63
L66:
	;
	v354 = int32(0)
	v357 = v295
	v363 = v302
	goto L65
L67:
	;
	goto L68
L68:
	;
	v318 = int32(0)
	v321 = v295
	v326 = v295
	v327 = v302
	goto L69
L69:
	;
	v328 = int32(2)
	v330 = v309 + v318<<(uint(v328)%32)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v330)))
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v331)+56))
	v333 = *(*float64)(unsafe.Add(mBase, uint32(v331)+64))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v330)+4))
	v337 = *(*float64)(unsafe.Add(mBase, uint32(v336)+56))
	v338 = *(*float64)(unsafe.Add(mBase, uint32(v336)+64))
	v340 = base.F64_add(base.F64_add(v327, base.F64_add(v332, v333)), base.F64_add(v337, v338))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v336)+38)))
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v331)+38)))
	v346 = base.B2i32(v341&v342 == int32(0)) | v321
	v348 = v318 + v328
	v350 = v326 + v328
	if v350 != v306&int32(2147483646) {
		v318 = v348
		v321 = v346
		v326 = v350
		v327 = v340
		goto L69
	} else {
		goto L71
	}
L70:
	;
	if v306&int32(1) == int32(0) {
		v379 = v346
		v385 = v340
		goto L63
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v354 = v348
	v357 = v346
	v363 = v340
	goto L65
L73:
	;
	if v379&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L74:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v389 <= int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v395 = int32(0)
	goto L76
L76:
	;
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v386)+12))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v405+v395<<(uint(int32(2))%32))))
	v410 = *(*float64)(unsafe.Add(mBase, uint32(v409)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v409)+48)) = base.F64_add(v385, v410)
	v413 = *(*float64)(unsafe.Add(mBase, uint32(v409)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v409)+56)) = base.F64_add(v385, v413)
	if v379&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L73
L78:
	;
	v416 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v409)+21)) = uint8(v416)
	goto L80
L79:
	;
	goto L80
L80:
	;
	v419 = v395 + int32(1)
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v419 < v420 {
		v395 = v419
		goto L76
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	v434 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v291)+26)) = uint8(v434)
	*(*int32)(unsafe.Add(mBase, uint32(v291)+40)) = v434
	goto L60
L83:
	;
	goto L84
L84:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v291)+40))
	if v438 == int32(0) {
		goto L61
	} else {
		goto L85
	}
L85:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v441 <= int32(0) {
		goto L61
	} else {
		goto L86
	}
L86:
	;
	v445 = int32(0)
	goto L87
L87:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v438)+12))
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v455+v445<<(uint(int32(2))%32))))
	v460 = *(*float64)(unsafe.Add(mBase, uint32(v459)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v459)+48)) = base.F64_add(v385, v460)
	v463 = *(*float64)(unsafe.Add(mBase, uint32(v459)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v459)+56)) = base.F64_add(v385, v463)
	v467 = v445 + int32(1)
	v468 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	if v467 < v468 {
		v445 = v467
		goto L87
	} else {
		goto L89
	}
L88:
	;
	goto L61
L89:
	;
	goto L88
L90:
	;
	v498 = base.F64_div(v492, v493)
	goto L92
L91:
	;
	v498 = v492
	goto L92
L92:
	;
	v499 = int32(0)
	if v490 == v499 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v630 != 0 {
		goto L132
	} else {
		goto L133
	}
L94:
	;
	v630 = int32(0)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v490)+4))
	if int32(0) < v510 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v518 = v499
	v521 = v499
	goto L100
L98:
	;
	v613 = v499
	goto L99
L99:
	;
	v630 = v613
	goto L93
L100:
	;
	v524 = *(*int32)(unsafe.Add(mBase, uint32(v490)+12))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v524+v521<<(uint(int32(2))%32))))
	if v518 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v613 = v598
	goto L99
L102:
	;
	v605 = v521 + int32(1)
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v490)+4))
	if v605 < v606 {
		v518 = v598
		v521 = v605
		goto L100
	} else {
		goto L131
	}
L103:
	;
	v529 = F_compare_fractional_path_costs(m, v518, v528, v498)
	mBase = m.M
	if v529 <= int32(0) {
		v598 = v518
		goto L102
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528)+64))
	if v491 == v532 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L105
L107:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v528)+16))
	if v586 != 0 {
		goto L125
	} else {
		goto L126
	}
L108:
	;
	v538 = int32(0)
	goto L109
L109:
	;
	v546 = int32(0)
	if v491 == v546 {
		v556 = v546
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v556 != 0 {
		v598 = v518
		goto L102
	} else {
		goto L124
	}
L111:
	;
	if v532 != 0 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v491)+4))
	if v550 <= v538 {
		v556 = int32(0)
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v491)+12))
	v556 = v552 + v538<<(uint(int32(2))%32)
	goto L111
L114:
	;
	if v556 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L115:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v532)+4))
	if v538 < v557 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if v556 == int32(0) {
		goto L107
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	v598 = v518
	goto L102
L120:
	;
	goto L110
L121:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v532)+12))
	if v563 == int32(0) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v556)))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v563+v538<<(uint(int32(2))%32))))
	if v570 == v572 {
		v538 = v538 + int32(1)
		goto L109
	} else {
		goto L123
	}
L123:
	;
	v598 = v518
	goto L102
L124:
	;
	goto L107
L125:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v586)+4))
	v589 = v587
	goto L127
L126:
	;
	v589 = int32(0)
	goto L127
L127:
	;
	v591 = F_bms_is_subset(m, v589, int32(0))
	mBase = m.M
	if v591 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v592 = v528
	goto L130
L129:
	;
	v592 = v518
	goto L130
L130:
	;
	v598 = v592
	goto L102
L131:
	;
	goto L101
L132:
	;
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v20)+264))
	v632 = F_make_pathtarget_from_tlist(m, v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	m.G0 = v17 + int32(16)
	return base.B2i32(v630 != int32(0))
L135:
	;
	v634 = F_set_pathtarget_cost_width(m, v20, v632)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v636 = F_apply_projection_to_path(m, v20, v291, v630, v634)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v638 = *(*float64)(unsafe.Add(mBase, uint32(v636)+56))
	v639 = *(*float64)(unsafe.Add(mBase, uint32(v636)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v636
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v20
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = base.F64_add(v639, base.F64_mul(v498, base.F64_sub(v638, v639)))
	goto L134
}
func F_builtin_validate_locale(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v10 != int32(67) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L31
	} else {
		goto L38
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L31
	} else {
		goto L34
	}
L3:
	;
	v103 = F_builtin_locale_encoding(m, v102)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L31
	} else {
		goto L32
	}
L4:
	;
	v15 = int32(_a_F_builtin_validate_locale_0)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_builtin_validate_locale[0])))
	if base.B2i32(v19 == int32(0))|base.B2i32(v19 != v22) != 0 {
		v40 = v19
		v41 = v22
		goto L8
	} else {
		goto L9
	}
L5:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+1)))
	if v13 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v102 = int32(_a_F_builtin_validate_locale_1)
	goto L3
L7:
	;
	if v40-v41 == int32(0) {
		v102 = v15
		goto L3
	} else {
		goto L14
	}
L8:
	;
	goto L7
L9:
	;
	v25 = l1
	v26 = v15
	goto L10
L10:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	if v30 == int32(0) {
		v40 = v30
		v41 = v29
		goto L8
	} else {
		goto L12
	}
L11:
	;
	v40 = v30
	v41 = v29
	goto L8
L12:
	;
	v33 = int32(1)
	if v30 == v29 {
		v25 = v25 + v33
		v26 = v26 + v33
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v45 = int32(_a_F_builtin_validate_locale_2)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v51 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_builtin_validate_locale[1])))
	if base.B2i32(v48 == int32(0))|base.B2i32(v48 != v51) != 0 {
		v69 = v48
		v70 = v51
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v69-v70 == int32(0) {
		v102 = v15
		goto L3
	} else {
		goto L22
	}
L16:
	;
	goto L15
L17:
	;
	v54 = l1
	v55 = v45
	goto L18
L18:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	if v59 == int32(0) {
		v69 = v59
		v70 = v58
		goto L16
	} else {
		goto L20
	}
L19:
	;
	v69 = v59
	v70 = v58
	goto L16
L20:
	;
	v62 = int32(1)
	if v59 == v58 {
		v54 = v54 + v62
		v55 = v55 + v62
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v74 = int32(_a_F_builtin_validate_locale_3)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v81 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_builtin_validate_locale[2])))
	if base.B2i32(v78 == int32(0))|base.B2i32(v78 != v81) != 0 {
		v99 = v78
		v100 = v81
		goto L24
	} else {
		goto L25
	}
L23:
	;
	if v99-v100 != 0 {
		goto L2
	} else {
		goto L30
	}
L24:
	;
	goto L23
L25:
	;
	v84 = l1
	v85 = v74
	goto L26
L26:
	;
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+1)))
	v89 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	if v89 == int32(0) {
		v99 = v89
		v100 = v88
		goto L24
	} else {
		goto L28
	}
L27:
	;
	v99 = v89
	v100 = v88
	goto L24
L28:
	;
	v92 = int32(1)
	if v89 == v88 {
		v84 = v84 + v92
		v85 = v85 + v92
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v102 = v74
	goto L3
L31:
	;
	return int32(0)
L32:
	;
	if base.B2i32(int32(0) <= v103)&base.B2i32(l0 != v103) != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	m.G0 = v8 + int32(32)
	return v102
L34:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L31
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
	F_errmsg(m, int32(_a_F_builtin_validate_locale_4), v8+int32(16))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(_a_F_builtin_validate_locale_5), int32(1526), int32(_a_F_builtin_validate_locale_6))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L31
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L31
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v146
	F_errmsg(m, int32(_a_F_builtin_validate_locale_7), v8)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L31
	} else {
		goto L44
	}
L41:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_builtin_validate_locale[3])))
	v146 = v144
	goto L43
L42:
	;
	v146 = int32(_a_F_builtin_validate_locale_8)
	goto L43
L43:
	;
	goto L40
L44:
	;
	F_errfinish(m, int32(_a_F_builtin_validate_locale_5), int32(1533), int32(_a_F_builtin_validate_locale_6))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_bytealtrim(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v12 = F_dobyteatrim(m, v3, v8, int32(1), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_byteaoverlay_no_len(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v14 == int32(1) {
				v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v20 == int32(18) {
					v23 = int32(16)
				} else {
					v23 = int32(0)
				}
				if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v30 = int32(4)
				} else {
					v30 = v23
				}
				v31 = F_bytea_overlay(m, v6, v11, v13, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					return v31
				}
			} else {
				if v14&int32(1) != 0 {
					v36 = int32(1)
					v40 = F_bytea_overlay(m, v6, v11, v13, int32(base.Ui32(v14)>>(uint(v36)%32))-v36)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						return v40
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v48 = F_bytea_overlay(m, v6, v11, v13, int32(base.Ui32(v43)>>(uint(int32(2))%32))-int32(4))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						return v48
					}
				}
			}
		}
	}
}
func F_bytearecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v8 = v6 - v7
	v10 = v8 + int32(4)
	v11 = F_palloc(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v10 << (uint(int32(2)) % 32)
		F_pq_copymsgbytes(m, v5, v11+int32(4), v8)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_byteatrim(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
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
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = int32(1)
			v12 = F_dobyteatrim(m, v3, v8, v10, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
