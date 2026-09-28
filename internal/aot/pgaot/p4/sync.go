package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_FinishSyncWorker(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[0]))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
	goto L1
L1:
	;
	if v13 == int32(2) {
		goto L2
	} else {
		goto L3
	}
L2:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v21 = int64(0)
	v23 = int32(_a_F_FinishSyncWorker_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[1]))
	v28 = base.AtomicRmwCmpxchg64(m, v24, int32(272), v21, v21)
	*(*int64)(unsafe.Add(mBase, _c_F_FinishSyncWorker[2])) = v28
	v30 = int32(0)
	v33 = base.AtomicRmwOr32(m, v30, int32(_a_F_FinishSyncWorker_1), v30)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[1]))
	v40 = base.AtomicRmwCmpxchg64(m, v36, int32(264), v21, v21)
	*(*int64)(unsafe.Add(mBase, _c_F_FinishSyncWorker[3])) = v40
	goto L8
L5:
	;
	return
L6:
	;
	v19 = F_pgstat_report_stat(m, int32(1))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	goto L4
L8:
	;
	F_XLogFlush(m, v40)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[4]))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+16)))
	if v46 != int32(1) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_proc_exit(m, int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L5
	} else {
		goto L44
	}
L11:
	;
	F_StartTransactionCommand(m)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L5
	} else {
		goto L34
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	if v49 != int32(2) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v54 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	if v54 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[5]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v58
	F_errmsg(m, int32(_a_F_FinishSyncWorker_2), v9)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[6]))
	v74 = F_LWLockAcquire(m, v70+int32(_a_F_FinishSyncWorker_3), int32(1))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L5
	} else {
		goto L20
	}
L18:
	;
	F_errfinish(m, int32(_a_F_FinishSyncWorker_4), int32(71), int32(_a_F_FinishSyncWorker_5))
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[7]))
	if v77 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[6]))
	F_LWLockRelease(m, v122+int32(_a_F_FinishSyncWorker_3))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L5
	} else {
		goto L33
	}
L22:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[4]))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+32))
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[8]))
	v88 = int32(0)
	goto L23
L23:
	;
	v95 = v84 + int32(16) + v88<<(uint(int32(7))%32)
	v96 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+16)))
	if v96 != int32(1) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v95)+120)) = int64(0)
	goto L21
L25:
	;
	goto L24
L26:
	;
	v111 = v88 + int32(1)
	if v111 != v77 {
		v88 = v111
		goto L23
	} else {
		goto L32
	}
L27:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v99 == int32(4) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v95)+32))
	if base.B2i32(v102 != v82)|base.B2i32(v99 != int32(3)) != 0 {
		goto L26
	} else {
		goto L29
	}
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v95)+36))
	if v107 != 0 {
		goto L26
	} else {
		goto L30
	}
L30:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v95)+20))
	if v108 != 0 {
		goto L25
	} else {
		goto L31
	}
L31:
	;
	goto L26
L32:
	;
	goto L21
L33:
	;
	goto L10
L34:
	;
	v131 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	if v131 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v134 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[5]))
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v134)+24))
	v137 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[4]))
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+36))
	v139 = F_get_rel_name(m, v138)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L5
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	F_CommitTransactionCommand(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L42
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v139
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v135
	F_errmsg(m, int32(_a_F_FinishSyncWorker_6), v9+int32(16))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L5
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_FinishSyncWorker_4), int32(85), int32(_a_F_FinishSyncWorker_5))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	goto L38
L42:
	;
	v157 = *(*int32)(unsafe.Add(mBase, _c_F_FinishSyncWorker[4]))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+32))
	F_logicalrep_worker_wakeup(m, v158)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	goto L10
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_SyncScanShmemInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	v4 = int32(_a_F_SyncScanShmemInit_0)
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_SyncScanShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v5 + int32(8)
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_SyncScanShmemInit[0]))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v10 + int32(464)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v10 + int32(32)
	v24 = int32(1)
	for {
		v28 = v24 * int32(24)
		v29 = int32(_a_F_SyncScanShmemInit_0)
		v30 = *(*int32)(unsafe.Add(mBase, _c_F_SyncScanShmemInit[0]))
		v31 = v28 + v30
		*(*int64)(unsafe.Add(mBase, uint32(v31)+24)) = int64(-4294967296)
		*(*int64)(unsafe.Add(mBase, uint32(v31)+16)) = int64(0)
		v37 = *(*int32)(unsafe.Add(mBase, _c_F_SyncScanShmemInit[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v37 + v28 - int32(16)
		v43 = *(*int32)(unsafe.Add(mBase, _c_F_SyncScanShmemInit[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v31)+12)) = v28 + v43 + int32(32)
		v49 = v24 + int32(1)
		if v49 != int32(19) {
			v24 = v49
			continue
		} else {
			break
		}
		break
	}
	v53 = v49 * int32(24)
	v54 = int32(_a_F_SyncScanShmemInit_0)
	v55 = *(*int32)(unsafe.Add(mBase, _c_F_SyncScanShmemInit[0]))
	v56 = v53 + v55
	*(*int64)(unsafe.Add(mBase, uint32(v56)+24)) = int64(-4294967296)
	*(*int64)(unsafe.Add(mBase, uint32(v56)+16)) = int64(0)
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_SyncScanShmemInit[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v56)+12)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v56)+8)) = v53 + v62 - int32(16)
	return
}
