package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BTreeShmemRequest(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	v4 = int32(12)
	Fn14205(m, l0, int32(_a_F_BTreeShmemRequest_0), int32(_a_F_BTreeShmemRequest_1), v4, int32(_a_F_BTreeShmemRequest_2), v4)
	v8 = m.ExcPending
	if v8 != 0 {
		return
	} else {
		return
	}
}
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
func F_BarrierArriveAndDetachExceptLast(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	v5 = base.AtomicRmwXchg32(m, l0, int32(0), int32(1))
	if v5 != 0 {
		F_s_lock(m, l0, int32(_a_F_BarrierArriveAndDetachExceptLast_0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if int32(2) <= v11 {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11 - int32(1)
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17 + int32(1)
			}
			v21 = int32(0)
			atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v21))
			return base.B2i32(v11 < int32(2))
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if int32(2) <= v11 {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v11 - int32(1)
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v17 + int32(1)
		}
		v21 = int32(0)
		atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v21))
		return base.B2i32(v11 < int32(2))
	}
}
func F_BarrierInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	v2 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(l0))), uint32(v2))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v2)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2
	v14 = l0 + int32(24)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v14))), uint32(v2))
	*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = int64(-1)
	return
}
func F_BecomeLockGroupMember(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_BecomeLockGroupMember[0]))
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_BecomeLockGroupMember[1]))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v12 = base.I32_div_s(l0-v9, int32(768))
	v14 = base.I32_rem_s(v12, int32(16))
	v19 = v6 + v14<<(uint(int32(7))%32) + int32(_a_F_BecomeLockGroupMember_0)
	v21 = F_LWLockAcquire(m, v19, v3)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v25 != l1 {
			v49 = v3
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+364))
			if v27 != l0 {
				v49 = v3
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_BecomeLockGroupMember[2]))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+364)) = l0
				v33 = l0 + int32(368)
				v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+372))
				if v34 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+372)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(l0)+368)) = v33
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(v30)+380)) = v33
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
				*(*int32)(unsafe.Add(mBase, uint32(v30)+376)) = v40
				v43 = v30 + int32(376)
				*(*int32)(unsafe.Add(mBase, uint32(v40)+4)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = v43
				v49 = int32(1)
			}
		}
		F_LWLockRelease(m, v19)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			return v49
		}
	}
}
func F_BeforeShmemExit_Files(m *base.Module, l0 int32, l1 int64) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[0]))
	if base.Ui32(int32(2)) <= base.Ui32(v7) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[1]))
	v13 = int32(1)
	v15 = v7
	v16 = v11
	goto L4
L2:
	;
	goto L3
L3:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[2])) = uint8(v46)
	v49 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[3]))
	if v46 < v49 {
		goto L12
	} else {
		goto L13
	}
L4:
	;
	v20 = v16 + v13*int32(48)
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+4)))
	if v21&int32(3) == int32(0) {
		v35 = v15
		v36 = v16
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	v38 = v13 + int32(1)
	if base.Ui32(v38) < base.Ui32(v35) {
		v13 = v38
		v15 = v35
		v16 = v36
		goto L4
	} else {
		goto L11
	}
L7:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v20)+32))
	if v26 == int32(0) {
		v35 = v15
		v36 = v16
		goto L6
	} else {
		goto L8
	}
L8:
	;
	F_FileClose(m, v13)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	return
L10:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[0]))
	v34 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[1]))
	v35 = v32
	v36 = v34
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
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[4]))
	v59 = F_FreeDesc(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L9
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_BeforeShmemExit_Files[3]))
	if int32(0) < v62 {
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
func F_before_shmem_exit(m *base.Module, l0 int32, l1 int64) {
	var v10 int32
	_ = v10
	Fn14230(m, l0, l1, int32(_a_F_before_shmem_exit_0), int32(349), int32(_a_F_before_shmem_exit_1), int32(_a_F_before_shmem_exit_2), int32(_a_F_before_shmem_exit_3), int32(_a_F_before_shmem_exit_4))
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(1058)
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
func F_big5_to_utf8(m *base.Module, l0 int32) int64 {
	var v3 int32
	_ = v3
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v3 = int32(0)
	v7 = Fn14231(m, l0, int32(36), v3, v3, v3, int32(_a_F_big5_to_utf8_0))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
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
	v9 = F_palloc(m, l0<<(uint(int32(3))%32)+int32(24))
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
func F_binaryheap_remove_first(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int64
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
	var v46 int64
	_ = v46
	var v50 int64
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
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v77 int64
	_ = v77
	var v79 int32
	_ = v79
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
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
	v21 = l0 + int32(24)
	v25 = *(*int64)(unsafe.Add(mBase, uint32(v21+v18<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = v25
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
	*(*int64)(unsafe.Add(mBase, uint32(v21+v32<<(uint(int32(3))%32)))) = v25
	return v10
L6:
	;
	goto L5
L7:
	;
	v43 = int32(3)
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v21+v39<<(uint(v43)%32))))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v21+v41<<(uint(v43)%32))))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v53 = m.T0[v52].(func(*base.Module, int64, int64, int32) int32)(m, v46, v50, v51)
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
	return int64(0)
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
	v66 = v21 + v61<<(uint(int32(3))%32)
	v67 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v70 = m.T0[v69].(func(*base.Module, int64, int64, int32) int32)(m, v25, v67, v68)
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
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	*(*int64)(unsafe.Add(mBase, uint32(v21+v32<<(uint(int32(3))%32)))) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = v79
	v32 = v61
	goto L4
}
func F_bitcmp(m *base.Module, l0 int32) int64 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int64(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_s(v99)
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
	v430 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
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
func F_bitshiftright(m *base.Module, l0 int32) int64 {
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int64
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v100 int32
	_ = v100
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
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
	var v204 int32
	_ = v204
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		if v15 < int32(0) {
			v19 = int32(0)
			v22 = int32(-2147483640)
			if base.Ui32(v15) <= base.Ui32(v22) {
				v25 = v22
			} else {
				v25 = v15
			}
			v28 = F_DirectFunctionCall2Coll(m, int32(1747), v19, base.I64_extend_i32_u(v11), base.I64_extend_i32_u(v19-v25))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int64(0)
			} else {
				return v28
			}
		} else {
			v31 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v34 = F_palloc(m, int32(base.Ui32(v31)>>(uint(int32(2))%32)))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int64(0)
			} else {
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v38 = v36 & int32(-4)
				*(*int32)(unsafe.Add(mBase, uint32(v34))) = v38
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v34)+4)) = v40
				v43 = v34 + int32(8)
				if v40 <= v15 {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v47 = int32(base.Ui32(v45) >> (uint(int32(2)) % 32))
					v49 = v47 - int32(8)
					if v45&int32(12)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v49)) == int32(0) {
						v57 = v34 + v47
						if base.Ui32(v57) <= base.Ui32(v43) {
							return base.I64_extend_i32_u(v34)
						} else {
							v60 = v34 + int32(12)
							if base.Ui32(v60) < base.Ui32(v57) {
								v62 = v57
							} else {
								v62 = v60
							}
							v69 = (v62-v34-int32(9))&int32(-4) + int32(4)
							if v69 == int32(0) {
								return base.I64_extend_i32_u(v34)
							} else {
								v204 = v69
								base.MemoryFill(m, v43, int32(0), v204)
								return base.I64_extend_i32_u(v34)
							}
						}
					} else {
						if v49 == int32(0) {
							return base.I64_extend_i32_u(v34)
						} else {
							v204 = v49
							base.MemoryFill(m, v43, int32(0), v204)
							return base.I64_extend_i32_u(v34)
						}
					}
				} else {
					v75 = v15 & int32(7)
					v77 = int32(base.Ui32(v15) >> (uint(int32(3)) % 32))
					if v15&int32(24)|base.B2i32(base.Ui32(int32(_a_F_bitshiftright_0)) < base.Ui32(v15)) == int32(0) {
						if v77 == int32(0) {
						} else {
							v89 = v34 + v77 + int32(8)
							v91 = v34 + int32(12)
							if base.Ui32(v91) < base.Ui32(v89) {
								v93 = v89
							} else {
								v93 = v91
							}
							v100 = (v93-v34-int32(9))&int32(-4) + int32(4)
							if v100 == int32(0) {
							} else {
								base.MemoryFill(m, v43, int32(0), v100)
							}
						}
					} else {
						if v77 == int32(0) {
						} else {
							base.MemoryFill(m, v43, int32(0), v77)
						}
					}
					v112 = v11 + int32(8)
					v113 = v77 + v43
					if v75 == int32(0) {
						v116 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
						v118 = int32(base.Ui32(v116) >> (uint(int32(2)) % 32))
						v121 = v118 - v77 - int32(8)
						if v121 != 0 {
							base.MemoryCopy(m, v113, v112, v121)
						} else {
						}
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						v167 = v118 + v34
						v173 = v124
					} else {
						if base.Ui32(int32(base.Ui32(v36)>>(uint(int32(2))%32))) <= base.Ui32(v77+int32(8)) {
							v167 = v113
							v173 = v38
						} else {
							v130 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v113))) = uint8(v130)
							v134 = v113
							v138 = v112
							for {
								v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
								v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
								v146 = v143 | int32(base.Ui32(v144)>>(uint(v75)%32))
								*(*uint8)(unsafe.Add(mBase, uint32(v134))) = uint8(v146)
								v149 = v134 + int32(1)
								v150 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
								v152 = int32(base.Ui32(v150) >> (uint(int32(2)) % 32))
								if base.Ui32(v149) < base.Ui32(v34+v152) {
									v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v138))))
									v156 = v155 << (uint(int32(8)-v75) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v149))) = uint8(v156)
									v158 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
									v161 = int32(base.Ui32(v158) >> (uint(int32(2)) % 32))
									v162 = v158
								} else {
									v161 = v152
									v162 = v150
								}
								if base.Ui32(v149) < base.Ui32(v161+v34) {
									v134 = v149
									v138 = v138 + int32(1)
									continue
								} else {
									break
								}
								break
							}
							v167 = v149
							v173 = v162
						}
					}
					v180 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
					v183 = v173<<(uint(int32(1))%32)&int32(-8) - v180 + int32(-64)
					if v183 <= int32(0) {
					} else {
						v187 = v167 - int32(1)
						v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
						v191 = v188 & (int32(255) << (uint(v183) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(v187))) = uint8(v191)
					}
					return base.I64_extend_i32_u(v34)
				}
			}
		}
	}
}
func F_blgetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int64
	_ = v124
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v151 int32
	_ = v151
	var v160 int64
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v202 int32
	_ = v202
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v305 int64
	_ = v305
	var v307 int32
	_ = v307
	var v329 int64
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v357 int64
	_ = v357
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v380 int64
	_ = v380
	v19 = int64(0)
	v20 = m.G0
	v22 = v20 - int32(16)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v25 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v22 + int32(16)
	return v380
L2:
	;
	v98 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L15
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v24)+1032))
	v29 = F_palloc0_mul(m, int32(2), v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int64(0)
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = v29
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v34 <= int32(0) {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v41 = v26
	v46 = int32(0)
	goto L7
L7:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v59&int32(1) != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	goto L2
L9:
	;
	F_pfree(m, v58)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v66 = *(*int64)(unsafe.Add(mBase, uint32(v41)+48))
	v67 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+4)))
	F_signValue(m, v24+int32(4), v58, v66, v67-int32(1))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v24))) = int32(0)
	v380 = v19
	goto L1
L13:
	;
	v75 = v46 + int32(1)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v75 < v76 {
		v41 = v41 + int32(56)
		v46 = v75
		goto L7
	} else {
		goto L14
	}
L14:
	;
	goto L8
L15:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v102 = F_RelationGetNumberOfBlocksInFork(m, v100, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+272))
	if v105 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v123 != 0 {
		goto L23
	} else {
		goto L24
	}
L18:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+268)))
	if v108 != int32(1) {
		v122 = v104
		goto L17
	} else {
		goto L21
	}
L19:
	;
	v115 = v105
	v116 = v104
	goto L20
L20:
	;
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v115)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v115)+16)) = v117 + int64(1)
	v122 = v116
	goto L17
L21:
	;
	F_pgstat_assoc_relation(m, v104)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+272))
	v115 = v114
	v116 = v113
	goto L20
L23:
	;
	v124 = *(*int64)(unsafe.Add(mBase, uint32(v123)))
	*(*int64)(unsafe.Add(mBase, uint32(v123))) = v124 + int64(1)
	goto L25
L24:
	;
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v102
	v132 = int32(0)
	v137 = F_read_stream_begin_relation(m, int32(12), v98, v122, v132, int32(3), v22+int32(8), v132)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v102) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v151 = int32(1)
	v160 = v19
	goto L30
L28:
	;
	v357 = v19
	goto L29
L29:
	;
	F_read_stream_end(m, v137)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L4
	} else {
		goto L60
	}
L30:
	;
	v162 = F_read_stream_next_buffer(m, v137, int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L32
	}
L31:
	;
	v357 = v329
	goto L29
L32:
	;
	F_LockBufferInternal(m, v162, int32(1))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	if v162 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	F_UnlockReleaseBuffer(m, v162)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L54
	}
L35:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184)+14)))
	if v185 == int32(0) {
		v329 = v160
		goto L34
	} else {
		goto L39
	}
L36:
	;
	v170 = *(*int32)(unsafe.Add(mBase, _c_F_blgetbitmap[0]))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v170+(v162^int32(-1))<<(uint(int32(2))%32))))
	v184 = v176
	goto L35
L37:
	;
	goto L38
L38:
	;
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_blgetbitmap[1]))
	v184 = v178 + v162<<(uint(int32(13))%32) + int32(-8192)
	goto L35
L39:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184)+16)))
	v189 = v184 + v188
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+2)))
	if v190&int32(2) != 0 {
		v329 = v160
		goto L34
	} else {
		goto L40
	}
L40:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v189))))
	if v193 == int32(0) {
		v329 = v160
		goto L34
	} else {
		goto L41
	}
L41:
	;
	v202 = int32(1)
	v217 = v160
	goto L42
L42:
	;
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v24)+1168))
	v224 = v184 + int32(24) + v218*(v202&int32(_a_F_blgetbitmap_0)-int32(1))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v24)+1032))
	if int32(0) < v225 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v329 = v305
	goto L34
L44:
	;
	v307 = v202 + int32(1)
	if base.Ui32(v307&int32(_a_F_blgetbitmap_0)) <= base.Ui32(v193) {
		v202 = v307
		v217 = v305
		goto L42
	} else {
		goto L53
	}
L45:
	;
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v234 = int32(0)
	goto L48
L46:
	;
	goto L47
L47:
	;
	v281 = int32(1)
	F_tbm_add_tuples(m, l1, v224, v281, v281)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L4
	} else {
		goto L52
	}
L48:
	;
	v252 = v234 << (uint(int32(1)) % 32)
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v230+v252))))
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v224+int32(6)+v252))))
	if v254&v256 != v254 {
		v305 = v217
		goto L44
	} else {
		goto L50
	}
L49:
	;
	goto L47
L50:
	;
	v260 = v234 + int32(1)
	if v260 != v225 {
		v234 = v260
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v305 = v217 + int64(1)
	goto L44
L53:
	;
	goto L43
L54:
	;
	v333 = *(*int32)(unsafe.Add(mBase, _c_F_blgetbitmap[2]))
	if v333 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L4
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v337 = v151 + int32(1)
	if v337 != v102 {
		v151 = v337
		v160 = v329
		goto L30
	} else {
		goto L59
	}
L58:
	;
	goto L57
L59:
	;
	goto L31
L60:
	;
	F_bms_free(m, v98)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	v380 = v357
	goto L1
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
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
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v337 int32
	_ = v337
	var v352 int32
	_ = v352
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
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
	F_LockBufferInternal(m, v35, int32(1))
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
	v361 = m.ExcPending
	if v361 != 0 {
		goto L1
	} else {
		goto L101
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_blinsert[0])) = v27
	F_MemoryContextDelete(m, v22)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L100
	}
L9:
	;
	F_LockBufferInternal(m, v35, int32(3))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
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
	F_UnlockBuffer(m, v35)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_UnlockBuffer(m, v35)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L43
	}
L17:
	;
	v67 = F_ReadBuffer(m, l0, v64)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	F_LockBufferInternal(m, v67, int32(3))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v72 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+16)))
	v101 = v75 + v100
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101))))
	v103 = v99 * v102
	v104 = int32(_a_F_blinsert_4) - v103
	if base.Ui32(v99) <= base.Ui32(v104) {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	v75 = F_GenericXLogRegisterBuffer(m, v72, v67, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+14)))
	if v77 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+16)))
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75+v78)+2)))
	if v80&int32(2) == int32(0) {
		goto L20
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v85 = int32(0)
	F_PageInit(m, v75, int32(_a_F_blinsert_1), int32(8))
	mBase = m.M
	v89 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v75)+16)))
	v90 = v75 + v89
	v91 = int32(_a_F_blinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+6)) = uint16(v91)
	*(*uint16)(unsafe.Add(mBase, uint32(v90)+2)) = uint16(v85)
	goto L27
L26:
	;
	goto L25
L27:
	;
	goto L20
L28:
	;
	if base.Ui32(v99) <= base.Ui32(v104) {
		goto L35
	} else {
		goto L36
	}
L29:
	;
	if v99 != 0 {
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
	base.MemoryCopy(m, v75+v103+int32(24), v32, v99)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101))))
	v112 = v110 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v101))) = uint16(v112)
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+1164)))
	v117 = v112*v114 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v75)+12)) = uint16(v117)
	goto L31
L35:
	;
	F_GenericXLogFinish(m, v72)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	F_pfree(m, v72)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L41
	}
L38:
	;
	F_UnlockReleaseBuffer(m, v67)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_ReleaseBuffer(m, v35)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	goto L8
L41:
	;
	F_UnlockReleaseBuffer(m, v67)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	v134 = v64
	goto L9
L43:
	;
	v134 = int32(-1)
	goto L9
L44:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+28)))
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+30)))
	if base.Ui32(v141) < base.Ui32(v142) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	F_UnlockReleaseBuffer(m, v35)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L1
	} else {
		goto L99
	}
L46:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v57+v141<<(uint(int32(2))%32))+168))
	v150 = v141 + base.B2i32(v134 == v147)
	goto L48
L47:
	;
	v150 = v141
	goto L48
L48:
	;
	v152 = v150 & int32(_a_F_blinsert_6)
	v153 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v156 = F_GenericXLogRegisterBuffer(m, v153, v35, int32(0))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+30)))
	if base.Ui32(v152) < base.Ui32(v158) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v161 = v152
	v162 = v153
	v166 = v156
	goto L54
L52:
	;
	v247 = v153
	v251 = v156
	goto L53
L53:
	;
	v256 = F_BloomNewBuffer(m, l0)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L82
	}
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v166+v161<<(uint(int32(2))%32))+168))
	v175 = F_ReadBuffer(m, l0, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v247 = v238
	v251 = v241
	goto L53
L56:
	;
	F_LockBufferInternal(m, v175, int32(3))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v181 = F_GenericXLogRegisterBuffer(m, v162, v175, int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v207 = v181 + v206
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207))))
	v209 = v205 * v208
	v210 = int32(_a_F_blinsert_4) - v209
	if base.Ui32(v205) <= base.Ui32(v210) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+14)))
	if v183 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v184)+2)))
	if v186&int32(2) == int32(0) {
		goto L58
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v191 = int32(0)
	F_PageInit(m, v181, int32(_a_F_blinsert_1), int32(8))
	mBase = m.M
	v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v196 = v181 + v195
	v197 = int32(_a_F_blinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+6)) = uint16(v197)
	*(*uint16)(unsafe.Add(mBase, uint32(v196)+2)) = uint16(v191)
	goto L64
L63:
	;
	goto L62
L64:
	;
	goto L58
L65:
	;
	if base.Ui32(v205) <= base.Ui32(v210) {
		goto L72
	} else {
		goto L73
	}
L66:
	;
	if v205 != 0 {
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
	base.MemoryCopy(m, v181+v209+int32(24), v32, v205)
	goto L71
L70:
	;
	goto L71
L71:
	;
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v207))))
	v218 = v216 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v207))) = uint16(v218)
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+1164)))
	v223 = v218*v220 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v181)+12)) = uint16(v223)
	goto L68
L72:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v166)+28)) = uint16(v161)
	F_GenericXLogFinish(m, v162)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L75
	}
L73:
	;
	goto L74
L74:
	;
	F_pfree(m, v162)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	F_UnlockReleaseBuffer(m, v175)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	goto L45
L77:
	;
	F_UnlockReleaseBuffer(m, v175)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v237 = v161 + int32(1)
	v238 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v241 = F_GenericXLogRegisterBuffer(m, v238, v35, int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241)+30)))
	if base.Ui32(v237) < base.Ui32(v243) {
		v161 = v237
		v162 = v238
		v166 = v241
		goto L54
	} else {
		goto L81
	}
L81:
	;
	goto L55
L82:
	;
	v259 = F_GenericXLogRegisterBuffer(m, v247, v256, int32(1))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v261 = int32(0)
	F_PageInit(m, v259, int32(_a_F_blinsert_1), int32(8))
	mBase = m.M
	v265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+16)))
	v266 = v259 + v265
	v267 = int32(_a_F_blinsert_5)
	*(*uint16)(unsafe.Add(mBase, uint32(v266)+6)) = uint16(v267)
	*(*uint16)(unsafe.Add(mBase, uint32(v266)+2)) = uint16(v261)
	goto L84
L84:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v259)+16)))
	v277 = v259 + v276
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277))))
	v279 = v275 * v278
	v280 = int32(_a_F_blinsert_4) - v279
	if base.Ui32(v275) <= base.Ui32(v280) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	if base.B2i32(base.Ui32(v275) <= base.Ui32(v280)) == int32(0) {
		goto L7
	} else {
		goto L92
	}
L86:
	;
	if v275 != 0 {
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
	base.MemoryCopy(m, v259+v279+int32(24), v32, v275)
	goto L91
L90:
	;
	goto L91
L91:
	;
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v277))))
	v288 = v286 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v277))) = uint16(v288)
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+1164)))
	v293 = v288*v290 + int32(24)
	*(*uint16)(unsafe.Add(mBase, uint32(v259)+12)) = uint16(v293)
	goto L88
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+28)) = int32(_a_F_blinsert_7)
	if v256 < int32(0) {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v251)+168)) = v319
	F_GenericXLogFinish(m, v247)
	mBase = m.M
	v322 = m.ExcPending
	if v322 != 0 {
		goto L1
	} else {
		goto L97
	}
L94:
	;
	v304 = *(*int32)(unsafe.Add(mBase, _c_F_blinsert[3]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v304+(v256^int32(-1))*int32(56))+16))
	v319 = v310
	goto L93
L95:
	;
	goto L96
L96:
	;
	v312 = *(*int32)(unsafe.Add(mBase, _c_F_blinsert[4]))
	v313 = int32(56)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v312+v256*v313-v313)+16))
	v319 = v318
	goto L93
L97:
	;
	F_UnlockReleaseBuffer(m, v256)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
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
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L102
	}
L102:
	;
	F_errfinish(m, int32(_a_F_blinsert_9), int32(324), int32(_a_F_blinsert_10))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
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
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 float64
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v14 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v126 = l1
	goto L3
L3:
	;
	m.G0 = v12 + int32(16)
	return v126
L4:
	;
	v21 = F_palloc0(m, int32(40))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v25 = l1
	goto L6
L6:
	;
	v27 = F_RelationGetNumberOfBlocksInFork(m, v17, int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L7
	} else {
		goto L9
	}
L7:
	;
	return int32(0)
L8:
	;
	v25 = v21
	goto L6
L9:
	;
	v29 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v29
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v27
	*(*int64)(unsafe.Add(mBase, uint32(v25)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v27
	v35 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v35
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v45 = F_read_stream_begin_relation(m, int32(13), v39, v17, v29, int32(3), v12+int32(8), v29)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v27) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v56 = v35
	goto L14
L12:
	;
	goto L13
L13:
	;
	F_read_stream_end(m, v45)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L7
	} else {
		goto L32
	}
L14:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	goto L13
L16:
	;
	v62 = F_read_stream_next_buffer(m, v45, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	F_LockBufferInternal(m, v62, int32(1))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	if v62 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	F_UnlockReleaseBuffer(m, v62)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L7
	} else {
		goto L30
	}
L20:
	;
	v100 = *(*float64)(unsafe.Add(mBase, uint32(v25)+8))
	v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87))))
	*(*float64)(unsafe.Add(mBase, uint32(v25)+8)) = base.F64_add(v100, base.F64_convert_i32_u(v101))
	goto L19
L21:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+14)))
	if v85 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_blvacuumcleanup[0]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v70+(v62^int32(-1))<<(uint(int32(2))%32))))
	v84 = v76
	goto L21
L23:
	;
	goto L24
L24:
	;
	v78 = *(*int32)(unsafe.Add(mBase, _c_F_blvacuumcleanup[1]))
	v84 = v78 + v62<<(uint(int32(13))%32) + int32(-8192)
	goto L21
L25:
	;
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v84)+16)))
	v87 = v84 + v86
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+2)))
	if v88&int32(2) == int32(0) {
		goto L20
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_RecordFreeIndexPage(m, v17, v56)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L29
	}
L28:
	;
	goto L27
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v25)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = v96 + int32(1)
	goto L19
L30:
	;
	v109 = v56 + int32(1)
	if v109 != v27 {
		v56 = v109
		goto L14
	} else {
		goto L31
	}
L31:
	;
	goto L15
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_FreeSpaceMapVacuum(m, v122)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	v126 = v25
	goto L3
}
func F_boolor_statefunc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v5 int64
	_ = v5
	var v10 int64
	_ = v10
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v2 == int64(0) {
		v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v10 = base.I64_extend_i32_u(base.B2i32(v5 != int64(0)))
	} else {
		v10 = int64(1)
	}
	return v10
}
func F_boot_get_role_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	v2 = int32(0)
	v4 = int32(_a_F_boot_get_role_oid_0)
	v7 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[0])))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v7 == v2)|base.B2i32(v7 != v10) != 0 {
		v28 = v7
		v29 = v10
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return v513
L2:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(v511)+4))
	v513 = v512
	goto L1
L3:
	;
	if v28-v29 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_1)
		goto L2
	} else {
		goto L10
	}
L4:
	;
	goto L3
L5:
	;
	v13 = v4
	v14 = l0
	goto L6
L6:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+1)))
	if v18 == int32(0) {
		v28 = v18
		v29 = v17
		goto L4
	} else {
		goto L8
	}
L7:
	;
	v28 = v18
	v29 = v17
	goto L4
L8:
	;
	v21 = int32(1)
	if v18 == v17 {
		v13 = v13 + v21
		v14 = v14 + v21
		goto L6
	} else {
		goto L9
	}
L9:
	;
	goto L7
L10:
	;
	v34 = int32(_a_F_boot_get_role_oid_2)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[1])))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v37 == int32(0))|base.B2i32(v37 != v40) != 0 {
		v58 = v37
		v59 = v40
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if v58-v59 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_3)
		goto L2
	} else {
		goto L18
	}
L12:
	;
	goto L11
L13:
	;
	v43 = v34
	v44 = l0
	goto L14
L14:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	if v48 == int32(0) {
		v58 = v48
		v59 = v47
		goto L12
	} else {
		goto L16
	}
L15:
	;
	v58 = v48
	v59 = v47
	goto L12
L16:
	;
	v51 = int32(1)
	if v48 == v47 {
		v43 = v43 + v51
		v44 = v44 + v51
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v64 = int32(_a_F_boot_get_role_oid_4)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[2])))
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v67 == int32(0))|base.B2i32(v67 != v70) != 0 {
		v88 = v67
		v89 = v70
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v88-v89 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_5)
		goto L2
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v73 = v64
	v74 = l0
	goto L22
L22:
	;
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+1)))
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+1)))
	if v78 == int32(0) {
		v88 = v78
		v89 = v77
		goto L20
	} else {
		goto L24
	}
L23:
	;
	v88 = v78
	v89 = v77
	goto L20
L24:
	;
	v81 = int32(1)
	if v78 == v77 {
		v73 = v73 + v81
		v74 = v74 + v81
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v94 = int32(_a_F_boot_get_role_oid_6)
	v97 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[3])))
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v97 == int32(0))|base.B2i32(v97 != v100) != 0 {
		v118 = v97
		v119 = v100
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v118-v119 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_7)
		goto L2
	} else {
		goto L34
	}
L28:
	;
	goto L27
L29:
	;
	v103 = v94
	v104 = l0
	goto L30
L30:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
	if v108 == int32(0) {
		v118 = v108
		v119 = v107
		goto L28
	} else {
		goto L32
	}
L31:
	;
	v118 = v108
	v119 = v107
	goto L28
L32:
	;
	v111 = int32(1)
	if v108 == v107 {
		v103 = v103 + v111
		v104 = v104 + v111
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v124 = int32(_a_F_boot_get_role_oid_8)
	v127 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[4])))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v127 == int32(0))|base.B2i32(v127 != v130) != 0 {
		v148 = v127
		v149 = v130
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v148-v149 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_9)
		goto L2
	} else {
		goto L42
	}
L36:
	;
	goto L35
L37:
	;
	v133 = v124
	v134 = l0
	goto L38
L38:
	;
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+1)))
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+1)))
	if v138 == int32(0) {
		v148 = v138
		v149 = v137
		goto L36
	} else {
		goto L40
	}
L39:
	;
	v148 = v138
	v149 = v137
	goto L36
L40:
	;
	v141 = int32(1)
	if v138 == v137 {
		v133 = v133 + v141
		v134 = v134 + v141
		goto L38
	} else {
		goto L41
	}
L41:
	;
	goto L39
L42:
	;
	v154 = int32(_a_F_boot_get_role_oid_10)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[5])))
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v157 == int32(0))|base.B2i32(v157 != v160) != 0 {
		v178 = v157
		v179 = v160
		goto L44
	} else {
		goto L45
	}
L43:
	;
	if v178-v179 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_11)
		goto L2
	} else {
		goto L50
	}
L44:
	;
	goto L43
L45:
	;
	v163 = v154
	v164 = l0
	goto L46
L46:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+1)))
	v168 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163)+1)))
	if v168 == int32(0) {
		v178 = v168
		v179 = v167
		goto L44
	} else {
		goto L48
	}
L47:
	;
	v178 = v168
	v179 = v167
	goto L44
L48:
	;
	v171 = int32(1)
	if v168 == v167 {
		v163 = v163 + v171
		v164 = v164 + v171
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v184 = int32(_a_F_boot_get_role_oid_12)
	v187 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[6])))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v187 == int32(0))|base.B2i32(v187 != v190) != 0 {
		v208 = v187
		v209 = v190
		goto L52
	} else {
		goto L53
	}
L51:
	;
	if v208-v209 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_13)
		goto L2
	} else {
		goto L58
	}
L52:
	;
	goto L51
L53:
	;
	v193 = v184
	v194 = l0
	goto L54
L54:
	;
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+1)))
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+1)))
	if v198 == int32(0) {
		v208 = v198
		v209 = v197
		goto L52
	} else {
		goto L56
	}
L55:
	;
	v208 = v198
	v209 = v197
	goto L52
L56:
	;
	v201 = int32(1)
	if v198 == v197 {
		v193 = v193 + v201
		v194 = v194 + v201
		goto L54
	} else {
		goto L57
	}
L57:
	;
	goto L55
L58:
	;
	v214 = int32(_a_F_boot_get_role_oid_14)
	v217 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[7])))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v217 == int32(0))|base.B2i32(v217 != v220) != 0 {
		v238 = v217
		v239 = v220
		goto L60
	} else {
		goto L61
	}
L59:
	;
	if v238-v239 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_15)
		goto L2
	} else {
		goto L66
	}
L60:
	;
	goto L59
L61:
	;
	v223 = v214
	v224 = l0
	goto L62
L62:
	;
	v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224)+1)))
	v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v223)+1)))
	if v228 == int32(0) {
		v238 = v228
		v239 = v227
		goto L60
	} else {
		goto L64
	}
L63:
	;
	v238 = v228
	v239 = v227
	goto L60
L64:
	;
	v231 = int32(1)
	if v228 == v227 {
		v223 = v223 + v231
		v224 = v224 + v231
		goto L62
	} else {
		goto L65
	}
L65:
	;
	goto L63
L66:
	;
	v244 = int32(_a_F_boot_get_role_oid_16)
	v247 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[8])))
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v247 == int32(0))|base.B2i32(v247 != v250) != 0 {
		v268 = v247
		v269 = v250
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v268-v269 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_17)
		goto L2
	} else {
		goto L74
	}
L68:
	;
	goto L67
L69:
	;
	v253 = v244
	v254 = l0
	goto L70
L70:
	;
	v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254)+1)))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v253)+1)))
	if v258 == int32(0) {
		v268 = v258
		v269 = v257
		goto L68
	} else {
		goto L72
	}
L71:
	;
	v268 = v258
	v269 = v257
	goto L68
L72:
	;
	v261 = int32(1)
	if v258 == v257 {
		v253 = v253 + v261
		v254 = v254 + v261
		goto L70
	} else {
		goto L73
	}
L73:
	;
	goto L71
L74:
	;
	v274 = int32(_a_F_boot_get_role_oid_18)
	v277 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[9])))
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v277 == int32(0))|base.B2i32(v277 != v280) != 0 {
		v298 = v277
		v299 = v280
		goto L76
	} else {
		goto L77
	}
L75:
	;
	if v298-v299 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_19)
		goto L2
	} else {
		goto L82
	}
L76:
	;
	goto L75
L77:
	;
	v283 = v274
	v284 = l0
	goto L78
L78:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284)+1)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+1)))
	if v288 == int32(0) {
		v298 = v288
		v299 = v287
		goto L76
	} else {
		goto L80
	}
L79:
	;
	v298 = v288
	v299 = v287
	goto L76
L80:
	;
	v291 = int32(1)
	if v288 == v287 {
		v283 = v283 + v291
		v284 = v284 + v291
		goto L78
	} else {
		goto L81
	}
L81:
	;
	goto L79
L82:
	;
	v304 = int32(_a_F_boot_get_role_oid_20)
	v307 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[10])))
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v307 == int32(0))|base.B2i32(v307 != v310) != 0 {
		v328 = v307
		v329 = v310
		goto L84
	} else {
		goto L85
	}
L83:
	;
	if v328-v329 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_21)
		goto L2
	} else {
		goto L90
	}
L84:
	;
	goto L83
L85:
	;
	v313 = v304
	v314 = l0
	goto L86
L86:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v314)+1)))
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+1)))
	if v318 == int32(0) {
		v328 = v318
		v329 = v317
		goto L84
	} else {
		goto L88
	}
L87:
	;
	v328 = v318
	v329 = v317
	goto L84
L88:
	;
	v321 = int32(1)
	if v318 == v317 {
		v313 = v313 + v321
		v314 = v314 + v321
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	v334 = int32(_a_F_boot_get_role_oid_22)
	v337 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[11])))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v337 == int32(0))|base.B2i32(v337 != v340) != 0 {
		v358 = v337
		v359 = v340
		goto L92
	} else {
		goto L93
	}
L91:
	;
	if v358-v359 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_23)
		goto L2
	} else {
		goto L98
	}
L92:
	;
	goto L91
L93:
	;
	v343 = v334
	v344 = l0
	goto L94
L94:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+1)))
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v343)+1)))
	if v348 == int32(0) {
		v358 = v348
		v359 = v347
		goto L92
	} else {
		goto L96
	}
L95:
	;
	v358 = v348
	v359 = v347
	goto L92
L96:
	;
	v351 = int32(1)
	if v348 == v347 {
		v343 = v343 + v351
		v344 = v344 + v351
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v364 = int32(_a_F_boot_get_role_oid_24)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[12])))
	v370 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v367 == int32(0))|base.B2i32(v367 != v370) != 0 {
		v388 = v367
		v389 = v370
		goto L100
	} else {
		goto L101
	}
L99:
	;
	if v388-v389 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_25)
		goto L2
	} else {
		goto L106
	}
L100:
	;
	goto L99
L101:
	;
	v373 = v364
	v374 = l0
	goto L102
L102:
	;
	v377 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v374)+1)))
	v378 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+1)))
	if v378 == int32(0) {
		v388 = v378
		v389 = v377
		goto L100
	} else {
		goto L104
	}
L103:
	;
	v388 = v378
	v389 = v377
	goto L100
L104:
	;
	v381 = int32(1)
	if v378 == v377 {
		v373 = v373 + v381
		v374 = v374 + v381
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v394 = int32(_a_F_boot_get_role_oid_26)
	v397 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[13])))
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v397 == int32(0))|base.B2i32(v397 != v400) != 0 {
		v418 = v397
		v419 = v400
		goto L108
	} else {
		goto L109
	}
L107:
	;
	if v418-v419 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_27)
		goto L2
	} else {
		goto L114
	}
L108:
	;
	goto L107
L109:
	;
	v403 = v394
	v404 = l0
	goto L110
L110:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v404)+1)))
	v408 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v403)+1)))
	if v408 == int32(0) {
		v418 = v408
		v419 = v407
		goto L108
	} else {
		goto L112
	}
L111:
	;
	v418 = v408
	v419 = v407
	goto L108
L112:
	;
	v411 = int32(1)
	if v408 == v407 {
		v403 = v403 + v411
		v404 = v404 + v411
		goto L110
	} else {
		goto L113
	}
L113:
	;
	goto L111
L114:
	;
	v424 = int32(_a_F_boot_get_role_oid_28)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[14])))
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v427 == int32(0))|base.B2i32(v427 != v430) != 0 {
		v448 = v427
		v449 = v430
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v448-v449 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_29)
		goto L2
	} else {
		goto L122
	}
L116:
	;
	goto L115
L117:
	;
	v433 = v424
	v434 = l0
	goto L118
L118:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v434)+1)))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433)+1)))
	if v438 == int32(0) {
		v448 = v438
		v449 = v437
		goto L116
	} else {
		goto L120
	}
L119:
	;
	v448 = v438
	v449 = v437
	goto L116
L120:
	;
	v441 = int32(1)
	if v438 == v437 {
		v433 = v433 + v441
		v434 = v434 + v441
		goto L118
	} else {
		goto L121
	}
L121:
	;
	goto L119
L122:
	;
	v454 = int32(_a_F_boot_get_role_oid_30)
	v457 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[15])))
	v460 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v457 == int32(0))|base.B2i32(v457 != v460) != 0 {
		v478 = v457
		v479 = v460
		goto L124
	} else {
		goto L125
	}
L123:
	;
	if v478-v479 == int32(0) {
		v511 = int32(_a_F_boot_get_role_oid_31)
		goto L2
	} else {
		goto L130
	}
L124:
	;
	goto L123
L125:
	;
	v463 = v454
	v464 = l0
	goto L126
L126:
	;
	v467 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v464)+1)))
	v468 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v463)+1)))
	if v468 == int32(0) {
		v478 = v468
		v479 = v467
		goto L124
	} else {
		goto L128
	}
L127:
	;
	v478 = v468
	v479 = v467
	goto L124
L128:
	;
	v471 = int32(1)
	if v468 == v467 {
		v463 = v463 + v471
		v464 = v464 + v471
		goto L126
	} else {
		goto L129
	}
L129:
	;
	goto L127
L130:
	;
	v483 = int32(_a_F_boot_get_role_oid_32)
	v486 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_boot_get_role_oid[16])))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if base.B2i32(v486 == int32(0))|base.B2i32(v486 != v489) != 0 {
		v507 = v486
		v508 = v489
		goto L132
	} else {
		goto L133
	}
L131:
	;
	if v507-v508 != 0 {
		v513 = v2
		goto L1
	} else {
		goto L138
	}
L132:
	;
	goto L131
L133:
	;
	v492 = v483
	v493 = l0
	goto L134
L134:
	;
	v496 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+1)))
	v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+1)))
	if v497 == int32(0) {
		v507 = v497
		v508 = v496
		goto L132
	} else {
		goto L136
	}
L135:
	;
	v507 = v497
	v508 = v496
	goto L132
L136:
	;
	v500 = int32(1)
	if v497 == v496 {
		v492 = v492 + v500
		v493 = v493 + v500
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v511 = int32(_a_F_boot_get_role_oid_33)
	goto L2
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
func F_bpcharle(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
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
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v18 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v22 = v20 & v18
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v18
	goto L6
L5:
	;
	v23 = int32(4)
	goto L6
L6:
	;
	v24 = v23 + v11
	if v20 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = v51
	goto L18
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v30 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v41 = int32(1)
	if v22 != 0 {
		v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v33 = int32(16)
	goto L13
L12:
	;
	v33 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = int32(4)
	goto L16
L15:
	;
	v40 = v33
	goto L16
L16:
	;
	v51 = v40
	goto L7
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	if v57 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v74 = int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v78 = v76 & v74
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v73 = v51 & (v51 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v67 = v57 - int32(1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v67))))
	if v69 == int32(32) {
		v57 = v67
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v73 = v57
	goto L20
L25:
	;
	v79 = v74
	goto L27
L26:
	;
	v79 = int32(4)
	goto L27
L27:
	;
	v80 = v79 + v16
	if v76 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v113 = v107
	goto L39
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v86 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v97 = int32(1)
	if v78 != 0 {
		v107 = int32(base.Ui32(v76)>>(uint(v97)%32)) - v97
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v89 = int32(16)
	goto L34
L33:
	;
	v89 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v96 = int32(4)
	goto L37
L36:
	;
	v96 = v89
	goto L37
L37:
	;
	v107 = v96
	goto L28
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v113 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = F_varstr_cmp(m, v24, v73, v80, v128, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L46
	}
L41:
	;
	goto L40
L42:
	;
	v128 = v107 & (v107 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v123 = v113 - int32(1)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v123))))
	if v125 == int32(32) {
		v113 = v123
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v128 = v113
	goto L41
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v133 != v11 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_pfree(m, v11)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v137 != v16 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	F_pfree(m, v16)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return base.I64_extend_i32_u(base.B2i32(v131 <= int32(0)))
L54:
	;
	goto L53
}
func F_bpcharne(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v264 int32
	_ = v264
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v23 = v21 & v19
	if v23 != 0 {
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
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L94
	}
L7:
	;
	v24 = v19
	goto L9
L8:
	;
	v24 = int32(4)
	goto L9
L9:
	;
	if v21 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v58 = v52
	goto L21
L11:
	;
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v31 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v42 = int32(1)
	if v23 != 0 {
		v52 = int32(base.Ui32(v21)>>(uint(v42)%32)) - v42
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v34 = int32(16)
	goto L16
L15:
	;
	v34 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v31-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v41 = int32(4)
	goto L19
L18:
	;
	v41 = v34
	goto L19
L19:
	;
	v52 = v41
	goto L10
L20:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v52 = int32(base.Ui32(v46)>>(uint(int32(2))%32)) - int32(4)
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
	v75 = int32(1)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v79 = v77 & v75
	if v79 != 0 {
		goto L28
	} else {
		goto L29
	}
L23:
	;
	goto L22
L24:
	;
	v74 = v52 & (v52 >> (uint(int32(31)) % 32))
	goto L23
L25:
	;
	goto L26
L26:
	;
	v68 = v58 - int32(1)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v11+v68))))
	if v70 == int32(32) {
		v58 = v68
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v74 = v58
	goto L23
L28:
	;
	v80 = v75
	goto L30
L29:
	;
	v80 = int32(4)
	goto L30
L30:
	;
	if v77 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v114 = v108
	goto L42
L32:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v87 == int32(18) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v98 = int32(1)
	if v79 != 0 {
		v108 = int32(base.Ui32(v77)>>(uint(v98)%32)) - v98
		goto L31
	} else {
		goto L41
	}
L35:
	;
	v90 = int32(16)
	goto L37
L36:
	;
	v90 = int32(0)
	goto L37
L37:
	;
	if base.Ui32((v87-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v97 = int32(4)
	goto L40
L39:
	;
	v97 = v90
	goto L40
L40:
	;
	v108 = v97
	goto L31
L41:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v108 = int32(base.Ui32(v102)>>(uint(int32(2))%32)) - int32(4)
	goto L31
L42:
	;
	if v114 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v132 = F_pg_newlocale_from_collation(m, v18)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	goto L43
L45:
	;
	v129 = v108 & (v108 >> (uint(int32(31)) % 32))
	goto L44
L46:
	;
	goto L47
L47:
	;
	v124 = v114 - int32(1)
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v16+v124))))
	if v126 == int32(32) {
		v114 = v124
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v129 = v114
	goto L44
L49:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v235 != v11 {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132))))
	if v134 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v129 != v74 {
		v234 = int32(1)
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v216 = int32(1)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v218&v216 != 0 {
		goto L79
	} else {
		goto L80
	}
L54:
	;
	v138 = int32(1)
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v140&v138 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v143 = v138
	goto L57
L56:
	;
	v143 = int32(4)
	goto L57
L57:
	;
	v144 = v11 + v143
	v145 = int32(1)
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v147&v145 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v150 = v145
	goto L60
L59:
	;
	v150 = int32(4)
	goto L60
L60:
	;
	v151 = v16 + v150
	if base.Ui32(int32(4)) <= base.Ui32(v74) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v234 = base.B2i32(v213 != int32(0))
	goto L49
L62:
	;
	v213 = int32(0)
	goto L61
L63:
	;
	v187 = v182
	v188 = v183
	v189 = v184
	goto L73
L64:
	;
	if (v144|v151)&int32(3) != 0 {
		v182 = v144
		v183 = v151
		v184 = v74
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v175 = v144
	v176 = v151
	v177 = v74
	goto L66
L66:
	;
	if v177 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v159 = v144
	v160 = v151
	v161 = v74
	goto L68
L68:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v164 != v165 {
		v182 = v159
		v183 = v160
		v184 = v161
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v175 = v170
	v176 = v168
	v177 = v172
	goto L66
L70:
	;
	v167 = int32(4)
	v168 = v160 + v167
	v170 = v159 + v167
	v172 = v161 - v167
	if base.Ui32(int32(3)) < base.Ui32(v172) {
		v159 = v170
		v160 = v168
		v161 = v172
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v182 = v175
	v183 = v176
	v184 = v177
	goto L63
L73:
	;
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	if v192 == v193 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v213 = v192 - v193
	goto L61
L75:
	;
	v195 = int32(1)
	v200 = v189 - v195
	if v200 != 0 {
		v187 = v187 + v195
		v188 = v188 + v195
		v189 = v200
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
	v221 = v216
	goto L81
L80:
	;
	v221 = int32(4)
	goto L81
L81:
	;
	v223 = int32(1)
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v225&v223 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v228 = v223
	goto L84
L83:
	;
	v228 = int32(4)
	goto L84
L84:
	;
	v230 = F_varstr_cmp(m, v11+v221, v74, v16+v228, v129, v18)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v234 = base.B2i32(v230 != int32(0))
	goto L49
L86:
	;
	F_pfree(m, v11)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v239 != v16 {
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
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	return base.I64_extend_i32_u(v234)
L93:
	;
	goto L92
L94:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(_a_F_bpcharne_0), int32(0))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errhint(m, int32(_a_F_bpcharne_1), int32(0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(_a_F_bpcharne_2), int32(741), int32(_a_F_bpcharne_3))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
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
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int64
	_ = v331
	var v332 int64
	_ = v332
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v344 int64
	_ = v344
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v496 int32
	_ = v496
	var v497 float64
	_ = v497
	var v499 float64
	_ = v499
	var v501 int32
	_ = v501
	var v505 int32
	_ = v505
	var v506 float64
	_ = v506
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v514 int32
	_ = v514
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v558 int32
	_ = v558
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 float64
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 float64
	_ = v743
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v767 float64
	_ = v767
	var v770 float64
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
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
	v770 = *(*float64)(unsafe.Add(mBase, uint32(v147)+8))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	F_brinRevmapTerminate(m, v771)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L3
	} else {
		goto L174
	}
L2:
	;
	v716 = int32(0)
	v718 = int32(1)
	v723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v723)+140))
	v725 = m.T0[v724].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v716, v716, v718, v716, int32(-1), v718, v147, v716)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L3
	} else {
		goto L169
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
	v701 = m.ExcPending
	if v701 != 0 {
		goto L3
	} else {
		goto L166
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
	*(*int64)(unsafe.Add(mBase, uint32(v130))) = base.I64_rotl(v111, int64(32))
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
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	if v440 == int32(0) {
		goto L2
	} else {
		goto L114
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
	v244 = F_mul_size(m, int32(40), v243)
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
	if v300 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v315 = F_shm_toc_allocate(m, v314, v216)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L3
	} else {
		goto L82
	}
L76:
	;
	F_UnregisterSnapshot(m, v212)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L3
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	F_DestroyParallelContext(m, v204)
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L3
	} else {
		goto L80
	}
L79:
	;
	goto L78
L80:
	;
	v309 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[4]))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v309)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v309)+72)) = v310 - int32(1)
	goto L81
L81:
	;
	goto L45
L82:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v315))) = v317
	v319 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+16)) = v189
	*(*uint8)(unsafe.Add(mBase, uint32(v315)+8)) = uint8(v190)
	*(*int32)(unsafe.Add(mBase, uint32(v315)+4)) = v319
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v315)+12)) = v323
	v327 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[6]))
	if v327 == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v315)+24)) = v332
	v335 = v315 + int32(32)
	v336 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v335))), uint32(v336))
	*(*int64)(unsafe.Add(mBase, uint32(v335)+4)) = int64(-1)
	goto L87
L84:
	;
	v332 = int64(0)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v331 = *(*int64)(unsafe.Add(mBase, uint32(v327)+392))
	v332 = v331
	goto L83
L87:
	;
	v341 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v315)+44)), uint32(v341))
	v344 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v315)+56)) = v344
	*(*int32)(unsafe.Add(mBase, uint32(v315)+48)) = v341
	*(*int64)(unsafe.Add(mBase, uint32(v315)+64)) = v344
	F_table_parallelscan_initialize(m, l0, v315+int32(96), v212)
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v355 = F_shm_toc_allocate(m, v354, v226)
	mBase = m.M
	v356 = m.ExcPending
	if v356 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v204)+44))
	F_tuplesort_initialize_shared(m, v355, v189, v357)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L3
	} else {
		goto L90
	}
L90:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v360, int64(-5764607523034234879), v315)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v364, int64(-5764607523034234878), v355)
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L3
	} else {
		goto L92
	}
L92:
	;
	v369 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[5]))
	if v369 != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v370 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v371 = F_shm_toc_allocate(m, v370, v294)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L3
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v384 = F_mul_size(m, int32(40), v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L3
	} else {
		goto L101
	}
L96:
	;
	if v294 != 0 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v374 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[5]))
	base.MemoryCopy(m, v371, v374, v294)
	goto L99
L98:
	;
	goto L99
L99:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v376, int64(-5764607523034234877), v371)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L3
	} else {
		goto L100
	}
L100:
	;
	goto L95
L101:
	;
	v386 = F_shm_toc_allocate(m, v381, v384)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v388, int64(-5764607523034234876), v386)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L3
	} else {
		goto L103
	}
L103:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v395 = F_mul_size(m, int32(128), v394)
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	v397 = F_shm_toc_allocate(m, v392, v395)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v399, int64(-5764607523034234875), v397)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	F_LaunchParallelWorkers(m, v204)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v204
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = v397
	*(*int32)(unsafe.Add(mBase, uint32(v192)+20)) = v386
	*(*int32)(unsafe.Add(mBase, uint32(v192)+16)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v192)+12)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v192)+8)) = v315
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v406 + int32(1)
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	if v415 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F__brin_end_parallel(m, v192)
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L3
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+64)) = v192
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v192)+12))
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[7]))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v426 = base.I32_div_s(v424, v425)
	F__brin_parallel_scan_and_build(m, v147, v421, v422, l0, l1, v426)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L3
	} else {
		goto L112
	}
L111:
	;
	goto L45
L112:
	;
	F_WaitForParallelWorkersToAttach(m, v204)
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	goto L45
L114:
	;
	v444 = F_palloc0(m, int32(12))
	mBase = m.M
	v445 = m.ExcPending
	if v445 != 0 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	v446 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v444))) = uint8(v446)
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v444)+4)) = v449
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v451)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v444)+8)) = v452
	v455 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[7]))
	v456 = F_tuplesort_begin_index_brin(m, v455, v444)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v147)+72)) = v456
	v459 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v459)+8))
	v464 = v460 + int32(44)
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v459)+4))
	goto L117
L117:
	;
	v485 = base.AtomicRmwXchg32(m, v464, int32(0), int32(1))
	if v485 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v497 = *(*float64)(unsafe.Add(mBase, uint32(v460)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v147)+16)) = v497
	v499 = *(*float64)(unsafe.Add(mBase, uint32(v460)+64))
	*(*float64)(unsafe.Add(mBase, uint32(v147)+8)) = v499
	v501 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v460)+44)), uint32(v501))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L3
	} else {
		goto L127
	}
L119:
	;
	F_s_lock(m, v464, int32(_a_F_brinbuild_5))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L3
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v460)+48))
	if v465 != v489 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L121
L123:
	;
	v491 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v464))), uint32(v491))
	F_ConditionVariableSleep(m, v460+int32(32), int32(134217767))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L3
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	goto L118
L126:
	;
	goto L117
L127:
	;
	v506 = *(*float64)(unsafe.Add(mBase, uint32(v147)+16))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	F_tuplesort_performsort(m, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v511 = F_brin_new_memtuple(m, v510)
	mBase = m.M
	v512 = m.ExcPending
	if v512 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	v514 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[3]))
	v519 = F_AllocSetContextCreateInternal(m, v514, int32(_a_F_brinbuild_6), int32(0), int32(_a_F_brinbuild_0), int32(_a_F_brinbuild_7))
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	v521 = int32(_a_F_brinbuild_8)
	v522 = *(*int32)(unsafe.Add(mBase, _c_F_brinbuild[3]))
	*(*int32)(unsafe.Add(mBase, _c_F_brinbuild[3])) = v519
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v528 = F_tuplesort_getbrintuple(m, v525, v20+int32(20))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L3
	} else {
		goto L132
	}
L131:
	;
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v147)+36))
	F_brin_fill_empty_ranges(m, v147, v687, v688)
	mBase = m.M
	v690 = m.ExcPending
	if v690 != 0 {
		goto L3
	} else {
		goto L163
	}
L132:
	;
	if v528 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	F_tuplesort_end(m, v532)
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L3
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v537 = v147 + int32(24)
	v539 = v528
	v540 = v511
	v541 = int32(-1)
	goto L138
L136:
	;
	v687 = int32(-1)
	goto L131
L137:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v639)+4))
	v659 = F_brin_form_tuple(m, v655, v656, v639, v20+int32(44))
	mBase = m.M
	v660 = m.ExcPending
	if v660 != 0 {
		goto L3
	} else {
		goto L160
	}
L138:
	;
	if v541 != int32(-1) {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	F_tuplesort_end(m, v632)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L3
	} else {
		goto L158
	}
L140:
	;
	v558 = v539
	goto L144
L141:
	;
	v603 = v539
	goto L142
L142:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v621 = F_brin_deform_tuple(m, v620, v603, v540)
	mBase = m.M
	v622 = m.ExcPending
	if v622 != 0 {
		goto L3
	} else {
		goto L154
	}
L143:
	;
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v592 = F_brin_form_tuple(m, v589, v575, v540, v20+int32(44))
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L3
	} else {
		goto L151
	}
L144:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	if v575 != v576 {
		goto L143
	} else {
		goto L146
	}
L145:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	F_tuplesort_end(m, v586)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L3
	} else {
		goto L150
	}
L146:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	F_union_tuples(m, v578, v540, v558)
	mBase = m.M
	v580 = m.ExcPending
	if v580 != 0 {
		goto L3
	} else {
		goto L147
	}
L147:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v584 = F_tuplesort_getbrintuple(m, v581, v20+int32(20))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	if v584 != 0 {
		v558 = v584
		goto L144
	} else {
		goto L149
	}
L149:
	;
	goto L145
L150:
	;
	v639 = v540
	v640 = v541
	goto L137
L151:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v592)))
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v599 = F_brin_doinsert(m, v594, v595, v596, v537, v597, v592, v598)
	mBase = m.M
	v600 = m.ExcPending
	if v600 != 0 {
		goto L3
	} else {
		goto L152
	}
L152:
	;
	F_MemoryContextReset(m, v519)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L3
	} else {
		goto L153
	}
L153:
	;
	v603 = v558
	goto L142
L154:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	F_brin_fill_empty_ranges(m, v147, v541, v623)
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L3
	} else {
		goto L155
	}
L155:
	;
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v603)))
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v147)+72))
	v630 = F_tuplesort_getbrintuple(m, v627, v20+int32(20))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L3
	} else {
		goto L156
	}
L156:
	;
	if v630 != 0 {
		v539 = v630
		v540 = v621
		v541 = v626
		goto L138
	} else {
		goto L157
	}
L157:
	;
	goto L139
L158:
	;
	v635 = int32(-1)
	if v626 == v635 {
		v687 = v635
		goto L131
	} else {
		goto L159
	}
L159:
	;
	v639 = v621
	v640 = v626
	goto L137
L160:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v659)))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v666 = F_brin_doinsert(m, v661, v662, v663, v537, v664, v659, v665)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L3
	} else {
		goto L161
	}
L161:
	;
	F_pfree(m, v659)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L3
	} else {
		goto L162
	}
L162:
	;
	v687 = v640
	goto L131
L163:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brinbuild[3])) = v522
	F_MemoryContextDelete(m, v519)
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v147)+64))
	F__brin_end_parallel(m, v695)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L3
	} else {
		goto L165
	}
L165:
	;
	v767 = v506
	goto L1
L166:
	;
	v702 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v702 + int32(4)
	F_errmsg_internal(m, int32(_a_F_brinbuild_9), v20+int32(16))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	F_errfinish(m, int32(_a_F_brinbuild_10), int32(1125), int32(_a_F_brinbuild_11))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L3
	} else {
		goto L168
	}
L168:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L169:
	;
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v147)+44))
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v147)+32))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v147)+48))
	v732 = F_brin_form_tuple(m, v727, v728, v729, v20+int32(20))
	mBase = m.M
	v733 = m.ExcPending
	if v733 != 0 {
		goto L3
	} else {
		goto L170
	}
L170:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(v147)))
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v147)+28))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v147)+40))
	v739 = *(*int32)(unsafe.Add(mBase, uint32(v147)+32))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v741 = F_brin_doinsert(m, v734, v735, v736, v147+int32(24), v739, v732, v740)
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L3
	} else {
		goto L171
	}
L171:
	;
	v743 = *(*float64)(unsafe.Add(mBase, uint32(v147)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v147)+8)) = base.F64_add(v743, float64(1))
	F_pfree(m, v732)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L3
	} else {
		goto L172
	}
L172:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v147)+32))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v147)+36))
	F_brin_fill_empty_ranges(m, v147, v749, v750)
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L3
	} else {
		goto L173
	}
L173:
	;
	v767 = v725
	goto L1
L174:
	;
	F_terminate_brin_buildstate(m, v147)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L3
	} else {
		goto L175
	}
L175:
	;
	v777 = F_palloc(m, int32(16))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L3
	} else {
		goto L176
	}
L176:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v777)+8)) = v770
	*(*float64)(unsafe.Add(mBase, uint32(v777))) = v767
	m.G0 = v20 + int32(48)
	return v777
}
func F_btarraycmp(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v2)
	}
}
func F_btint2cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	v2 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return v2 - v3
}
func F_btint2skipsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(197)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(198)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(32767)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(-32768)
	return int64(0)
}
func F_btint42cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_s(base.B2i32(v4 < v3) - base.B2i32(v3 < v4))
}
func F_btint4cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_s(base.B2i32(v4 < v3) - base.B2i32(v3 < v4))
}
func F_btint4skipsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(200)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(201)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(2147483647)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(-2147483648)
	return int64(0)
}
func F_btoidsortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v2)+16)) = int32(205)
	return int64(0)
}
func F_btoidvectorcmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	var v56 int64
	_ = v56
	v7 = int64(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	F_check_valid_oidvector(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	F_check_valid_oidvector(m, v8)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v16 == v17 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return v56
L5:
	;
	v30 = int32(0)
	goto L10
L6:
	;
	if v16 <= int32(0) {
		v56 = v7
		goto L4
	} else {
		goto L9
	}
L7:
	;
	goto L8
L8:
	;
	return base.I64_extend_i32_s(v16 - v17)
L9:
	;
	v21 = int32(24)
	goto L5
L10:
	;
	v37 = v30 << (uint(int32(2)) % 32)
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9+v21+v37)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v8+v21+v37)))
	if v39 == v41 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	if base.Ui32(v41) < base.Ui32(v39) {
		goto L16
	} else {
		goto L17
	}
L12:
	;
	v44 = v30 + int32(1)
	if v16 != v44 {
		v30 = v44
		goto L10
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	goto L11
L15:
	;
	v56 = v7
	goto L4
L16:
	;
	v49 = int64(1)
	goto L18
L17:
	;
	v49 = int64(-1)
	goto L18
L18:
	;
	v56 = v49
	goto L4
}
func F_btrim1(m *base.Module, l0 int32) int64 {
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = Fn14240(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return v4
	}
}
func F_bttext_pattern_cmp(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v11 = F_pg_detoast_datum_packed(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int64(0)
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
			v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v95 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int64(0)
				} else {
					v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v99 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v102 = m.ExcPending
						if v102 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_s(v94)
						}
					} else {
						return base.I64_extend_i32_s(v94)
					}
				}
			} else {
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v99 != v11 {
					F_pfree(m, v11)
					mBase = m.M
					v102 = m.ExcPending
					if v102 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_s(v94)
					}
				} else {
					return base.I64_extend_i32_s(v94)
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
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
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
	var v231 int32
	_ = v231
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v266 int32
	_ = v266
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v293 int64
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v310 float64
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 float64
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 float64
	_ = v340
	var v341 float64
	_ = v341
	var v344 int32
	_ = v344
	var v345 float64
	_ = v345
	var v346 float64
	_ = v346
	var v348 float64
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v371 float64
	_ = v371
	var v375 int32
	_ = v375
	var v376 float64
	_ = v376
	var v377 float64
	_ = v377
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v393 float64
	_ = v393
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v403 int32
	_ = v403
	var v413 int32
	_ = v413
	var v417 int32
	_ = v417
	var v418 float64
	_ = v418
	var v421 float64
	_ = v421
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 float64
	_ = v468
	var v471 float64
	_ = v471
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 float64
	_ = v500
	var v501 float64
	_ = v501
	var v506 float64
	_ = v506
	var v507 int32
	_ = v507
	var v518 int32
	_ = v518
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v540 int32
	_ = v540
	var v546 int32
	_ = v546
	var v554 int32
	_ = v554
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v580 int32
	_ = v580
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v621 int32
	_ = v621
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 float64
	_ = v646
	var v647 float64
	_ = v647
	v5 = l4
	v6 = l5
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v20 = F_palloc(m, int32(400))
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
	base.MemoryCopy(m, v20, l0, int32(400))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = l0
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v28 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v27 + v28
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v34 = F_choose_plan_name(m, v31, int32(_a_F_build_minmax_path_0), v28)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v34
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+344)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+28)) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v37
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v46 = F_copyObjectImpl(m, v45)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v46
	v49 = int32(1)
	F_IncrementVarSublevelsUp(m, v46, v49, v49)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	v54 = F_copyObjectImpl(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v54
	v57 = int32(1)
	F_IncrementVarSublevelsUp(m, v54, v57, v57)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v62 = F_copyObjectImpl(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v66 = F_pstrdup(m, int32(_a_F_build_minmax_path_1))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v69 = F_makeTargetEntry(m, v62, int32(1), v66, int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v69
	v76 = F_list_make1_impl(m, int32(1), v17+int32(4))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46)+76)) = v76
	*(*int32)(unsafe.Add(mBase, uint32(v20)+284)) = v76
	v80 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+112)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+334)) = uint8(v80)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+40)) = uint8(v80)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+120)) = v80
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+36)) = uint8(v80)
	v91 = F_palloc0(m, int32(20))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = int32(52)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v98 = F_copyObjectImpl(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v91)+16)) = int32(-1)
	v102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v91)+12)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v91)+4)) = v98
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v46)+60))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	v107 = F_list_member(m, v106, v91)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	if v107 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v46)+60))
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v111)+8))
	v113 = F_lcons(m, v91, v112)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v119 = F_palloc0(m, int32(20))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v46)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v115)+8)) = v113
	goto L17
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v119))) = int32(106)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v20)+284))
	v124 = int32(0)
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v69)+16))
	if v133 == v124 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v266 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+18)) = uint8(v266)
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+17)) = uint8(v6)
	*(*uint8)(unsafe.Add(mBase, uint32(v119)+16)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v119)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v119)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v257
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v119
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v119
	v276 = F_list_make1_impl(m, int32(1), v17)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L56
	}
L21:
	;
	if v123 == int32(0) {
		v253 = int32(1)
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v257 = v133
	goto L23
L23:
	;
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v69)+16)) = v253
	v257 = v253
	goto L23
L25:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v140 <= int32(0) {
		v253 = int32(1)
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v143 = int32(0)
	if v143 < v140 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v146 = v140
	goto L29
L28:
	;
	v146 = v143
	goto L29
L29:
	;
	v148 = v146 & int32(3)
	v149 = int32(0)
	if int32(4) <= v140 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	v253 = v231 + int32(1)
	goto L24
L31:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v158 = v149
	v159 = int32(0)
	v160 = v124
	goto L34
L32:
	;
	v195 = v149
	v197 = v124
	goto L33
L33:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v123)+12))
	v206 = int32(0)
	v208 = v195
	v210 = v197
	goto L50
L34:
	;
	v169 = v154 + v160<<(uint(int32(2))%32)
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v170)+16))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v169)+8))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v172)+16))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v169)+4))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)+16))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v169)))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	if base.Ui32(v158) < base.Ui32(v177) {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v148 == int32(0) {
		v231 = v185
		goto L30
	} else {
		goto L49
	}
L36:
	;
	v179 = v177
	goto L38
L37:
	;
	v179 = v158
	goto L38
L38:
	;
	if base.Ui32(v179) < base.Ui32(v175) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v181 = v175
	goto L41
L40:
	;
	v181 = v179
	goto L41
L41:
	;
	if base.Ui32(v181) < base.Ui32(v173) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v183 = v173
	goto L44
L43:
	;
	v183 = v181
	goto L44
L44:
	;
	if base.Ui32(v183) < base.Ui32(v171) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v185 = v171
	goto L47
L46:
	;
	v185 = v183
	goto L47
L47:
	;
	v186 = int32(4)
	v187 = v160 + v186
	v189 = v159 + v186
	if v189 != v146&int32(2147483644) {
		v158 = v185
		v159 = v189
		v160 = v187
		goto L34
	} else {
		goto L48
	}
L48:
	;
	goto L35
L49:
	;
	v195 = v185
	v197 = v187
	goto L33
L50:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v204+v210<<(uint(int32(2))%32))))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+16))
	if base.Ui32(v208) < base.Ui32(v221) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v231 = v223
	goto L30
L52:
	;
	v223 = v221
	goto L54
L53:
	;
	v223 = v208
	goto L54
L54:
	;
	v224 = int32(1)
	v227 = v206 + v224
	if v227 != v148 {
		v206 = v227
		v208 = v223
		v210 = v210 + v224
		goto L50
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	v278 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+128)) = v278
	*(*int32)(unsafe.Add(mBase, uint32(v46)+124)) = v276
	v288 = F_makeConst(m, int32(20), int32(-1), v278, int32(8), int64(1), v278, int32(1))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L57
	}
L57:
	;
	v290 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v46)+136)) = v290
	*(*int32)(unsafe.Add(mBase, uint32(v46)+132)) = v288
	v293 = int64(4607182418800017408)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+320)) = v293
	*(*int64)(unsafe.Add(mBase, uint32(v20)+312)) = v293
	v299 = F_query_planner(m, v20, int32(877), v290)
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_SS_identify_outer_params(m, v20)
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v303 = int32(0)
	v310 = float64(0)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v20)+80))
	if v311 == v303 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v498 = *(*int32)(unsafe.Add(mBase, uint32(v299)+44))
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v20)+176))
	v500 = float64(1)
	v501 = *(*float64)(unsafe.Add(mBase, uint32(v299)+16))
	if base.F64_gt(v501, v500) != 0 {
		goto L90
	} else {
		goto L91
	}
L61:
	;
	goto L60
L62:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v311)+4))
	if v314 <= int32(0) {
		v387 = v303
		v393 = v310
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v299)+44))
	if v394 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L64:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v311)+12))
	if v314 == int32(1) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v317+v362<<(uint(int32(2))%32))))
	v376 = *(*float64)(unsafe.Add(mBase, uint32(v375)+56))
	v377 = *(*float64)(unsafe.Add(mBase, uint32(v375)+64))
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+39)))
	v387 = v380 ^ int32(1) | v365
	v393 = base.F64_add(v371, base.F64_add(v376, v377))
	goto L63
L66:
	;
	v362 = int32(0)
	v365 = v303
	v371 = v310
	goto L65
L67:
	;
	goto L68
L68:
	;
	v326 = int32(0)
	v329 = v303
	v334 = v303
	v335 = v310
	goto L69
L69:
	;
	v336 = int32(2)
	v338 = v317 + v326<<(uint(v336)%32)
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
	v340 = *(*float64)(unsafe.Add(mBase, uint32(v339)+56))
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v339)+64))
	v344 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
	v345 = *(*float64)(unsafe.Add(mBase, uint32(v344)+56))
	v346 = *(*float64)(unsafe.Add(mBase, uint32(v344)+64))
	v348 = base.F64_add(base.F64_add(v335, base.F64_add(v340, v341)), base.F64_add(v345, v346))
	v349 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v344)+39)))
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v339)+39)))
	v354 = base.B2i32(v349&v350 == int32(0)) | v329
	v356 = v326 + v336
	v358 = v334 + v336
	if v358 != v314&int32(2147483646) {
		v326 = v356
		v329 = v354
		v334 = v358
		v335 = v348
		goto L69
	} else {
		goto L71
	}
L70:
	;
	if v314&int32(1) == int32(0) {
		v387 = v354
		v393 = v348
		goto L63
	} else {
		goto L72
	}
L71:
	;
	goto L70
L72:
	;
	v362 = v356
	v365 = v354
	v371 = v348
	goto L65
L73:
	;
	if v387&int32(1) != 0 {
		goto L82
	} else {
		goto L83
	}
L74:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v397 <= int32(0) {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v403 = int32(0)
	goto L76
L76:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v394)+12))
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v413+v403<<(uint(int32(2))%32))))
	v418 = *(*float64)(unsafe.Add(mBase, uint32(v417)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v417)+48)) = base.F64_add(v393, v418)
	v421 = *(*float64)(unsafe.Add(mBase, uint32(v417)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v417)+56)) = base.F64_add(v393, v421)
	if v387&int32(1) != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	goto L73
L78:
	;
	v424 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v417)+21)) = uint8(v424)
	goto L80
L79:
	;
	goto L80
L80:
	;
	v427 = v403 + int32(1)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v394)+4))
	if v427 < v428 {
		v403 = v427
		goto L76
	} else {
		goto L81
	}
L81:
	;
	goto L77
L82:
	;
	v442 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v299)+26)) = uint8(v442)
	*(*int32)(unsafe.Add(mBase, uint32(v299)+52)) = v442
	goto L60
L83:
	;
	goto L84
L84:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v299)+52))
	if v446 == int32(0) {
		goto L61
	} else {
		goto L85
	}
L85:
	;
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	if v449 <= int32(0) {
		goto L61
	} else {
		goto L86
	}
L86:
	;
	v453 = int32(0)
	goto L87
L87:
	;
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v446)+12))
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v463+v453<<(uint(int32(2))%32))))
	v468 = *(*float64)(unsafe.Add(mBase, uint32(v467)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v467)+48)) = base.F64_add(v393, v468)
	v471 = *(*float64)(unsafe.Add(mBase, uint32(v467)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v467)+56)) = base.F64_add(v393, v471)
	v475 = v453 + int32(1)
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v446)+4))
	if v475 < v476 {
		v453 = v475
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
	v506 = base.F64_div(v500, v501)
	goto L92
L91:
	;
	v506 = v500
	goto L92
L92:
	;
	v507 = int32(0)
	if v498 == v507 {
		goto L94
	} else {
		goto L95
	}
L93:
	;
	if v638 != 0 {
		goto L132
	} else {
		goto L133
	}
L94:
	;
	v638 = int32(0)
	goto L93
L95:
	;
	goto L96
L96:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if int32(0) < v518 {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v526 = v507
	v529 = v507
	goto L100
L98:
	;
	v621 = v507
	goto L99
L99:
	;
	v638 = v621
	goto L93
L100:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v498)+12))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v532+v529<<(uint(int32(2))%32))))
	if v526 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	v621 = v606
	goto L99
L102:
	;
	v613 = v529 + int32(1)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v498)+4))
	if v613 < v614 {
		v526 = v606
		v529 = v613
		goto L100
	} else {
		goto L131
	}
L103:
	;
	v537 = F_compare_fractional_path_costs(m, v526, v536, v506)
	mBase = m.M
	if v537 <= int32(0) {
		v606 = v526
		goto L102
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v536)+64))
	if v499 == v540 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	goto L105
L107:
	;
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v536)+16))
	if v594 != 0 {
		goto L125
	} else {
		goto L126
	}
L108:
	;
	v546 = int32(0)
	goto L109
L109:
	;
	v554 = int32(0)
	if v499 == v554 {
		v564 = v554
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v564 != 0 {
		v606 = v526
		goto L102
	} else {
		goto L124
	}
L111:
	;
	if v540 != 0 {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v499)+4))
	if v558 <= v546 {
		v564 = int32(0)
		goto L111
	} else {
		goto L113
	}
L113:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v499)+12))
	v564 = v560 + v546<<(uint(int32(2))%32)
	goto L111
L114:
	;
	if v564 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L115:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v540)+4))
	if v546 < v565 {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if v564 == int32(0) {
		goto L107
	} else {
		goto L119
	}
L118:
	;
	goto L117
L119:
	;
	v606 = v526
	goto L102
L120:
	;
	goto L110
L121:
	;
	v571 = *(*int32)(unsafe.Add(mBase, uint32(v540)+12))
	if v571 == int32(0) {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v564)))
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v571+v546<<(uint(int32(2))%32))))
	if v578 == v580 {
		v546 = v546 + int32(1)
		goto L109
	} else {
		goto L123
	}
L123:
	;
	v606 = v526
	goto L102
L124:
	;
	goto L107
L125:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
	v597 = v595
	goto L127
L126:
	;
	v597 = int32(0)
	goto L127
L127:
	;
	v599 = F_bms_is_subset(m, v597, int32(0))
	mBase = m.M
	if v599 != 0 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v600 = v536
	goto L130
L129:
	;
	v600 = v526
	goto L130
L130:
	;
	v606 = v600
	goto L102
L131:
	;
	goto L101
L132:
	;
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v20)+284))
	v640 = F_make_pathtarget_from_tlist(m, v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
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
	return base.B2i32(v638 != int32(0))
L135:
	;
	v642 = F_set_pathtarget_cost_width(m, v20, v640)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v644 = F_apply_projection_to_path(m, v20, v299, v638, v642)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v646 = *(*float64)(unsafe.Add(mBase, uint32(v644)+56))
	v647 = *(*float64)(unsafe.Add(mBase, uint32(v644)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v644
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v20
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = base.F64_add(v647, base.F64_mul(v506, base.F64_sub(v646, v647)))
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
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
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
	F_errfinish(m, int32(_a_F_builtin_validate_locale_5), int32(1820), int32(_a_F_builtin_validate_locale_6))
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
	if base.B2i32(l0 == int32(7))|base.B2i32(base.Ui32(int32(41)) < base.Ui32(l0)) != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v151
	F_errmsg(m, int32(_a_F_builtin_validate_locale_7), v8)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L31
	} else {
		goto L44
	}
L41:
	;
	v151 = int32(_a_F_builtin_validate_locale_8)
	goto L43
L42:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_c_F_builtin_validate_locale[3])))
	v151 = v150
	goto L43
L43:
	;
	goto L40
L44:
	;
	F_errfinish(m, int32(_a_F_builtin_validate_locale_5), int32(1827), int32(_a_F_builtin_validate_locale_6))
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
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
func F_byteacat(m *base.Module, l0 int32) int64 {
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
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = F_bytea_catenate(m, v3, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v10)
			}
		}
	}
}
func F_byteafastcmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v10 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(l1))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v19&int32(1) != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v22 = int32(1)
	goto L6
L5:
	;
	v22 = int32(4)
	goto L6
L6:
	;
	v23 = int32(1)
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
	v27 = v25 & v23
	if v27 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = v23
	goto L9
L8:
	;
	v28 = int32(4)
	goto L9
L9:
	;
	v29 = v28 + v10
	v30 = v17 + v22
	if v25 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v19 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L11:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
	if v36 == int32(18) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	v47 = int32(1)
	if v27 != 0 {
		v57 = int32(base.Ui32(v25)>>(uint(v47)%32)) - v47
		goto L10
	} else {
		goto L20
	}
L14:
	;
	v39 = int32(16)
	goto L16
L15:
	;
	v39 = int32(0)
	goto L16
L16:
	;
	if base.Ui32((v36-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = int32(4)
	goto L19
L18:
	;
	v46 = v39
	goto L19
L19:
	;
	v57 = v46
	goto L10
L20:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
	v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
	goto L10
L21:
	;
	v87 = base.B2i32(v57 < v86)
	if v57 < v86 {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
	if v63 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v74 = int32(1)
	if v19&v74 != 0 {
		v86 = int32(base.Ui32(v19)>>(uint(v74)%32)) - v74
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v66 = int32(16)
	goto L27
L26:
	;
	v66 = int32(0)
	goto L27
L27:
	;
	if base.Ui32((v63-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v73 = int32(4)
	goto L30
L29:
	;
	v73 = v66
	goto L30
L30:
	;
	v86 = v73
	goto L21
L31:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v86 = int32(base.Ui32(v80)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v88 = v57
	goto L34
L33:
	;
	v88 = v86
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v88) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if base.I64_extend_i32_u(v10) != l0 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v150 = int32(0)
	goto L35
L37:
	;
	v124 = v119
	v125 = v120
	v126 = v121
	goto L47
L38:
	;
	if (v29|v30)&int32(3) != 0 {
		v119 = v29
		v120 = v30
		v121 = v88
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v112 = v29
	v113 = v30
	v114 = v88
	goto L40
L40:
	;
	if v114 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v96 = v29
	v97 = v30
	v98 = v88
	goto L42
L42:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	if v101 != v102 {
		v119 = v96
		v120 = v97
		v121 = v98
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v112 = v107
	v113 = v105
	v114 = v109
	goto L40
L44:
	;
	v104 = int32(4)
	v105 = v97 + v104
	v107 = v96 + v104
	v109 = v98 - v104
	if base.Ui32(int32(3)) < base.Ui32(v109) {
		v96 = v107
		v97 = v105
		v98 = v109
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v119 = v112
	v120 = v113
	v121 = v114
	goto L37
L47:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v124))))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	if v129 == v130 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v150 = v129 - v130
	goto L35
L49:
	;
	v132 = int32(1)
	v137 = v126 - v132
	if v137 != 0 {
		v124 = v124 + v132
		v125 = v125 + v132
		v126 = v137
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v10)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	if base.I64_extend_i32_u(v17) != l1 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v17)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	if v150 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v161 = v150
	goto L63
L62:
	;
	v161 = base.B2i32(v86 < v57) - v87
	goto L63
L63:
	;
	return v161
}
func F_bytealtrim(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v12 = F_dobyteatrim(m, v3, v8, int32(1), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v12)
			}
		}
	}
}
func F_byteaoctetlen(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_toast_raw_datum_size(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_s(v3 - int32(4))
	}
}
func F_byteaoverlay_no_len(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = F_pg_detoast_datum_packed(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
			if v13 == int32(1) {
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
				if v19 == int32(18) {
					v22 = int32(16)
				} else {
					v22 = int32(0)
				}
				if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v29 = int32(4)
				} else {
					v29 = v22
				}
				v42 = v29
			} else {
				v30 = int32(1)
				if v13&v30 != 0 {
					v42 = int32(base.Ui32(v13)>>(uint(v30)%32)) - v30
				} else {
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v42 = int32(base.Ui32(v36)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v43 = F_bytea_overlay(m, v4, v9, base.I32_wrap_i64(v11), v42)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v43)
			}
		}
	}
}
func F_bytearecv(m *base.Module, l0 int32) int64 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
	v8 = v6 - v7
	v10 = v8 + int32(4)
	v11 = F_palloc(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = v10 << (uint(int32(2)) % 32)
		F_pq_copymsgbytes(m, v5, v11+int32(4), v8)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v11)
		}
	}
}
func F_byteatrim(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v10 = int32(1)
			v12 = F_dobyteatrim(m, v3, v8, v10, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v12)
			}
		}
	}
}
