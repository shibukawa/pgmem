package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_ReadyForQuery(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(2)
	if base.Ui32(l0-v9) <= base.Ui32(v9) {
		F_pq_beginmessage(m, v7, int32(90))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = m.G0
			v18 = v16 - int32(16)
			m.G0 = v18
			v21 = *(*int32)(unsafe.Add(mBase, _consts[25]))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
			if base.Ui32(int32(20)) <= base.Ui32(v22) {
				F_errstart_cold(m, int32(22), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
					if base.Ui32(v29) <= base.Ui32(int32(19)) {
						v38 = *(*int32)(unsafe.Add(mBase, uint32(v29<<(uint(int32(2))%32))+uint32(_consts[1205])))
						v39 = v38
					} else {
						v39 = int32(556295)
					}
					*(*int32)(unsafe.Add(mBase, uint32(v18))) = v39
					F_errmsg_internal(m, int32(207134), v18)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errfinish(m, int32(504485), int32(5036), int32(423212))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v51 = int32(*(*int8)(unsafe.Add(mBase, uint32(v22)+uint32(_consts[1206]))))
				m.G0 = v18 + int32(16)
				F_enlargeStringInfo(m, v7, int32(1))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v59 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					*(*uint8)(unsafe.Add(mBase, uint32(v58+v59))) = uint8(v51)
					*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v58 + int32(1)
					F_pq_endmessage(m, v7)
					mBase = m.M
					v66 = m.ExcPending
					if v66 != 0 {
						return
					} else {
						v68 = *(*int32)(unsafe.Add(mBase, _consts[323]))
						v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
						v70 = m.T0[v69].(func(*base.Module) int32)(m)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return
						} else {
							m.G0 = v7 + int32(16)
							return
						}
					}
				}
			}
		}
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_RecordKnownAssignedTransactionIds(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v10 = F_errstart(m, int32(11), int32(0))
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
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v6))) = l0
	v14 = *(*int32)(unsafe.Add(mBase, _consts[1164]))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v14
	F_errmsg_internal(m, int32(55875), v6)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v25 = *(*int32)(unsafe.Add(mBase, _consts[1164]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v25))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(l0)) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L6:
	;
	F_errfinish(m, int32(503161), int32(4410), int32(177803))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	m.G0 = v6 + int32(16)
	return
L9:
	;
	if v37 == int32(0) {
		goto L8
	} else {
		goto L13
	}
L10:
	;
	v37 = base.B2i32(base.Ui32(v25) < base.Ui32(l0))
	goto L9
L11:
	;
	goto L12
L12:
	;
	v37 = base.B2i32(int32(0) < l0-v25)
	goto L9
L13:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[1164]))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v41)) == int32(0) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if v53 != 0 {
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v53 = base.B2i32(base.Ui32(v41) < base.Ui32(l0))
	goto L14
L16:
	;
	goto L17
L17:
	;
	v53 = int32(base.Ui32(v41-l0) >> (uint(int32(31)) % 32))
	goto L14
L18:
	;
	v55 = v41
	goto L21
L19:
	;
	goto L20
L20:
	;
	v81 = *(*int32)(unsafe.Add(mBase, _consts[266]))
	if base.Ui32(v81) <= base.Ui32(int32(1)) {
		goto L32
	} else {
		goto L33
	}
L21:
	;
	v57 = int32(3)
	v59 = v55 + int32(1)
	if base.Ui32(v59) <= base.Ui32(v57) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L20
L23:
	;
	v62 = v57
	goto L25
L24:
	;
	v62 = v59
	goto L25
L25:
	;
	F_ExtendSUBTRANS(m, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v62)) == int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v76 != 0 {
		v55 = v62
		goto L21
	} else {
		goto L31
	}
L28:
	;
	v76 = base.B2i32(base.Ui32(v62) < base.Ui32(l0))
	goto L27
L29:
	;
	goto L30
L30:
	;
	v76 = int32(base.Ui32(v62-l0) >> (uint(int32(31)) % 32))
	goto L27
L31:
	;
	goto L22
L32:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1164])) = l0
	goto L8
L33:
	;
	goto L34
L34:
	;
	v86 = int32(3)
	v88 = *(*int32)(unsafe.Add(mBase, _consts[1164]))
	v90 = v88 + int32(1)
	if base.Ui32(v90) <= base.Ui32(v86) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v93 = v86
	goto L37
L36:
	;
	v93 = v90
	goto L37
L37:
	;
	F_KnownAssignedXidsAdd(m, v93, l0, int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1164])) = l0
	F_AdvanceNextFullTransactionIdPastXid(m, l0)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	goto L8
}
func F_RegisterDynamicBackgroundWorker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	v3 = int32(0)
	v11 = int32(*(*uint8)(unsafe.Add(mBase, _consts[427])))
	if v11 != int32(1) {
		v134 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v134
L2:
	;
	v15 = F_SanityCheckBackgroundWorker(m, l0, int32(21))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	if v15 == int32(0) {
		v134 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v23 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	v27 = F_LWLockAcquire(m, v23+int32(4224), int32(0))
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v30 = *(*int32)(unsafe.Add(mBase, _consts[732]))
	v32 = v21 & int32(16)
	if v32 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
	if int32(0) < v49 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _consts[733]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v30)+8))
	if base.Ui32(v37-v38) < base.Ui32(v36) {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v42 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v42+int32(4224))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	return int32(0)
L11:
	;
	v58 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v124+int32(4224))
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L3
	} else {
		goto L31
	}
L14:
	;
	v66 = v30 + int32(16) + v58*int32(1480)
	v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v67 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L13
L16:
	;
	goto L20
L17:
	;
	goto L18
L18:
	;
	v111 = v58 + int32(1)
	if v111 != v49 {
		v58 = v111
		goto L14
	} else {
		goto L30
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v66)+4)) = int32(-1)
	v77 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v66)+1)) = uint8(v77)
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v66)+8))
	v81 = v79 + int64(1)
	*(*int64)(unsafe.Add(mBase, uint32(v66)+8)) = v81
	if v32 != 0 {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v73 = F__emscripten_memcpy_bulkmem(m, v66+int32(16), l0, int32(1460))
	mBase = m.M
	goto L22
L22:
	;
	goto L19
L23:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v30)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v83 + int32(1)
	goto L25
L24:
	;
	goto L25
L25:
	;
	v87 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v66))) = uint8(v87)
	v91 = *(*int32)(unsafe.Add(mBase, _consts[86]))
	F_LWLockRelease(m, v91+int32(4224))
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	F_SendPostmasterSignal(m, int32(6))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	if l1 == int32(0) {
		v134 = v87
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v102 = F_palloc(m, int32(16))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v102))) = v58
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int64)(unsafe.Add(mBase, uint32(v106)+8)) = v81
	return int32(1)
L30:
	;
	goto L15
L31:
	;
	v134 = int32(0)
	goto L1
}
func F_ReleaseExternalFD(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v3 int32
	_ = v3
	v1 = int32(4452268)
	v3 = *(*int32)(unsafe.Add(mBase, _consts[272]))
	*(*int32)(unsafe.Add(mBase, _consts[272])) = v3 - int32(1)
	return
}
func F_RelfilenumberMapInvalidateCallback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v12 = *(*int32)(unsafe.Add(mBase, _consts[1383]))
	F_hash_seq_init(m, v7+int32(12), v12)
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
	v17 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L19
	}
L4:
	;
	if v17 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v21 = v17
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v7 + int32(32)
	return
L8:
	;
	if l1 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L7
L10:
	;
	v41 = F_hash_seq_search(m, v7+int32(12))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L17
	}
L11:
	;
	v31 = *(*int32)(unsafe.Add(mBase, _consts[1383]))
	v34 = F_hash_search(m, v31, v21, int32(2), int32(0))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L15
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v25 == int32(0) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if l1 != v25 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	if v34 == int32(0) {
		goto L3
	} else {
		goto L16
	}
L16:
	;
	goto L10
L17:
	;
	if v41 != 0 {
		v21 = v41
		goto L8
	} else {
		goto L18
	}
L18:
	;
	goto L9
L19:
	;
	F_errmsg_internal(m, int32(454824), int32(0))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	F_errfinish(m, int32(507561), int32(76), int32(325864))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RememberConstraintForRebuilding(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
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
	var v70 int32
	_ = v70
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
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	v3 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v11 == v3 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L18
	} else {
		goto L44
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L18
	} else {
		goto L41
	}
L3:
	;
	m.G0 = v9 + int32(32)
	return
L4:
	;
	if v50 != 0 {
		goto L3
	} else {
		goto L17
	}
L5:
	;
	v50 = int32(0)
	goto L4
L6:
	;
	goto L7
L7:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v18 <= int32(0) {
		v43 = v3
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v50 = v43
	goto L4
L9:
	;
	v21 = int32(0)
	if v21 < v18 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v24 = v18
	goto L12
L11:
	;
	v24 = v21
	goto L12
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v27 = int32(0)
	goto L13
L13:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25+v27<<(uint(int32(2))%32))))
	v36 = base.B2i32(v35 == l0)
	if v35 == l0 {
		v43 = v36
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v43 = v36
	goto L8
L15:
	;
	v38 = v27 + int32(1)
	if v38 != v24 {
		v27 = v38
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v52 = int32(0)
	v54 = F_pg_get_constraintdef_worker(m, l0, int32(1), v52, v52)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return
L19:
	;
	v56 = F_get_constraint_type(m, l0)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+108))
	if v56 == int32(110) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+112)) = v73
	v75 = F_get_constraint_index(m, l0)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L18
	} else {
		goto L29
	}
L22:
	;
	v61 = F_lcons_oid(m, l0, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L18
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v67 = F_lappend_oid(m, v58, l0)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v61
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v65 = F_lcons(m, v54, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v73 = v65
	goto L21
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+108)) = v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	v71 = F_lappend(m, v70, v54)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v73 = v71
	goto L21
L29:
	;
	if v75 == int32(0) {
		goto L3
	} else {
		goto L30
	}
L30:
	;
	v79 = F_get_index_isreplident(m, v75)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L31
	}
L31:
	;
	if v79 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l1)+124))
	if v81 != 0 {
		goto L2
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v85 = F_get_index_isclustered(m, v75)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L18
	} else {
		goto L37
	}
L35:
	;
	v82 = F_get_rel_name(m, v75)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+124)) = v82
	goto L34
L37:
	;
	if v85 == int32(0) {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+128))
	if v89 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v90 = F_get_rel_name(m, v75)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L18
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+128)) = v90
	goto L3
L41:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v104
	F_errmsg_internal(m, int32(10533), v9+int32(16))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L18
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(505917), int32(15275), int32(344158))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v120
	F_errmsg_internal(m, int32(161141), v9)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(505917), int32(15290), int32(344229))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L18
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_RemoveInheritance(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
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
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v357 int32
	_ = v357
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v465 int32
	_ = v465
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v529 int32
	_ = v529
	var v534 int32
	_ = v534
	v4 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(208)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+119)))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v25 = F_DeleteInheritsTuple(m, v20, v21, l2, v22+int32(4))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v517 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v519 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v518 + v519
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v517 + v519
	F_errmsg(m, int32(721871), v16+int32(32))
	mBase = m.M
	v529 = m.ExcPending
	if v529 != 0 {
		goto L2
	} else {
		goto L143
	}
L2:
	;
	return
L3:
	;
	if v25 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L2
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v58 = F_table_open(m, int32(1249), int32(3))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L2
	} else {
		goto L12
	}
L7:
	;
	F_errcode(m, int32(16908420))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	if v19 == int32(112) {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v40 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v39 + v40
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v38 + v40
	F_errmsg(m, int32(721748), v16+int32(48))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	F_errfinish(m, int32(505917), int32(17982), int32(427263))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L12:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v16-int32(-64), int32(1), int32(3), int32(184), v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v69 = int32(1)
	v74 = F_systable_beginscan(m, v58, int32(2659), v69, int32(0), v69, v16-int32(-64))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L2
	} else {
		goto L14
	}
L14:
	;
	v76 = F_systable_getnext(m, v74)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L15
	}
L15:
	;
	if v76 != 0 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v80 = v76
	goto L19
L17:
	;
	goto L18
L18:
	;
	F_systable_endscan(m, v74)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L2
	} else {
		goto L34
	}
L19:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v91)+22)))
	v93 = v91 + v92
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v93)+91)))
	if v94 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L18
L21:
	;
	v129 = F_systable_getnext(m, v74)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L32
	}
L22:
	;
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v93)+94)))
	if v95 <= int32(0) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	v101 = F_SearchSysCacheExistsAttName(m, v98, v93+int32(4))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L2
	} else {
		goto L24
	}
L24:
	;
	if v101 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v105 = F_heap_copytuple(m, v80)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v105)+16))
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+22)))
	v109 = v107 + v108
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+94)))
	v112 = v110 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v109)+94)) = uint16(v112)
	if v112&int32(65535) == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v118 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v109)+92)) = uint8(v118)
	goto L29
L28:
	;
	goto L29
L29:
	;
	F_CatalogTupleUpdate(m, v58, v105+int32(4), v105)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	F_pfree(m, v105)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	goto L21
L32:
	;
	if v129 != 0 {
		v80 = v129
		goto L19
	} else {
		goto L33
	}
L33:
	;
	goto L20
L34:
	;
	F_sequence_close(m, v58, int32(3))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L2
	} else {
		goto L35
	}
L35:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v152 = F_build_attrmap_by_name(m, v149, v150, int32(0))
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L36
	}
L36:
	;
	v156 = F_table_open(m, int32(2606), int32(3))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_ScanKeyInit(m, v16-int32(-64), int32(9), int32(3), int32(184), v163)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L38
	}
L38:
	;
	v167 = int32(1)
	v172 = F_systable_beginscan(m, v156, int32(2665), v167, int32(0), v167, v16-int32(-64))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v174 = F_systable_getnext(m, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L2
	} else {
		goto L40
	}
L40:
	;
	if v174 != 0 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v178 = v174
	v184 = v4
	v185 = v4
	goto L44
L42:
	;
	v234 = v4
	v235 = v4
	goto L43
L43:
	;
	F_systable_endscan(m, v172)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L2
	} else {
		goto L58
	}
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v178)+16))
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189)+22)))
	v191 = v189 + v190
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+106)))
	if v192 != 0 {
		v222 = v184
		v223 = v185
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v234 = v222
	v235 = v223
	goto L43
L46:
	;
	v224 = F_systable_getnext(m, v172)
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L2
	} else {
		goto L56
	}
L47:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+72)))
	if v193 == int32(99) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v198 = F_pstrdup(m, v191+int32(4))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L2
	} else {
		goto L51
	}
L49:
	;
	v203 = v184
	v204 = v193
	goto L50
L50:
	;
	if v204&int32(255) != int32(110) {
		v222 = v203
		v223 = v185
		goto L46
	} else {
		goto L53
	}
L51:
	;
	v200 = F_lappend(m, v184, v198)
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L52
	}
L52:
	;
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191)+72)))
	v203 = v200
	v204 = v202
	goto L50
L53:
	;
	v209 = F_extractNotNullColumn(m, v178)
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	v217 = int32(*(*int16)(unsafe.Add(mBase, uint32(v211+v209<<(uint(int32(1))%32)-int32(2)))))
	v218 = F_lappend_int(m, v185, v217)
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	v222 = v203
	v223 = v218
	goto L46
L56:
	;
	if v224 != 0 {
		v178 = v224
		v184 = v222
		v185 = v223
		goto L44
	} else {
		goto L57
	}
L57:
	;
	goto L45
L58:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_ScanKeyInit(m, v16-int32(-64), int32(9), int32(3), int32(184), v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L2
	} else {
		goto L59
	}
L59:
	;
	v250 = int32(1)
	v255 = F_systable_beginscan(m, v156, int32(2665), v250, int32(0), v250, v16-int32(-64))
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L2
	} else {
		goto L61
	}
L60:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L2
	} else {
		goto L140
	}
L61:
	;
	v257 = F_systable_getnext(m, v255)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	if v257 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v265 = v257
	v267 = v234
	v268 = v235
	goto L66
L64:
	;
	v437 = v234
	v438 = v235
	goto L65
L65:
	;
	if v437|v438 != 0 {
		goto L118
	} else {
		goto L119
	}
L66:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v265)+16))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272)+22)))
	v274 = v272 + v273
	v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+72)))
	switch v275 - int32(99) {
	case 0:
		goto L71
	default:
		v422 = v267
		v423 = v268
		goto L68
	case 11:
		goto L70
	}
L67:
	;
	v437 = v422
	v438 = v423
	goto L65
L68:
	;
	v427 = F_systable_getnext(m, v255)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L2
	} else {
		goto L116
	}
L69:
	;
	v391 = F_heap_copytuple(m, v265)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L2
	} else {
		goto L109
	}
L70:
	;
	v341 = F_extractNotNullColumn(m, v265)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L2
	} else {
		goto L94
	}
L71:
	;
	if v267 == int32(0) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v422 = int32(0)
	v423 = v268
	goto L68
L73:
	;
	goto L74
L74:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	if v281 <= int32(0) {
		v422 = v267
		v423 = v268
		goto L68
	} else {
		goto L75
	}
L75:
	;
	v285 = v274 + int32(4)
	v286 = int32(0)
	if v286 < v281 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v289 = v281
	goto L78
L77:
	;
	v289 = v286
	goto L78
L78:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v267)+12))
	v294 = int32(0)
	goto L79
L79:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v290+v294<<(uint(int32(2))%32))))
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v308))))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285))))
	if v312 == int32(0) {
		v331 = v311
		v332 = v312
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v422 = v267
	v423 = v268
	goto L68
L81:
	;
	if v332-v331 == int32(0) {
		goto L89
	} else {
		goto L90
	}
L82:
	;
	goto L81
L83:
	;
	if v311 != v312 {
		v331 = v311
		v332 = v312
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v316 = v285
	v317 = v308
	goto L85
L85:
	;
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v317)+1)))
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v316)+1)))
	if v321 == int32(0) {
		v331 = v320
		v332 = v321
		goto L82
	} else {
		goto L87
	}
L86:
	;
	v331 = v320
	v332 = v321
	goto L82
L87:
	;
	v324 = int32(1)
	if v320 == v321 {
		v316 = v316 + v324
		v317 = v317 + v324
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	v336 = F_list_delete_nth_cell(m, v267, v294)
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L2
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v339 = v294 + int32(1)
	if v339 != v289 {
		v294 = v339
		goto L79
	} else {
		goto L93
	}
L92:
	;
	v386 = v336
	v387 = v268
	goto L69
L93:
	;
	goto L80
L94:
	;
	if v268 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L95:
	;
	v422 = v267
	v423 = int32(0)
	goto L68
L96:
	;
	goto L97
L97:
	;
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v268)+4))
	if v346 <= int32(0) {
		v422 = v267
		v423 = v268
		goto L68
	} else {
		goto L98
	}
L98:
	;
	v349 = int32(0)
	if v349 < v346 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v352 = v346
	goto L101
L100:
	;
	v352 = v349
	goto L101
L101:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v268)+12))
	v357 = int32(0)
	goto L102
L102:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v353+v357<<(uint(int32(2))%32))))
	if v341 == v371 {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	v422 = v267
	v423 = v268
	goto L68
L104:
	;
	v373 = F_list_delete_nth_cell(m, v268, v357)
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L2
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v376 = v357 + int32(1)
	if v376 != v352 {
		v357 = v376
		goto L102
	} else {
		goto L108
	}
L107:
	;
	v386 = v267
	v387 = v373
	goto L69
L108:
	;
	goto L103
L109:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)+16))
	v394 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v393)+22)))
	v395 = v393 + v394
	v396 = int32(*(*int16)(unsafe.Add(mBase, uint32(v395)+104)))
	if v396 <= int32(0) {
		goto L60
	} else {
		goto L110
	}
L110:
	;
	v400 = v396 - int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v395)+104)) = uint16(v400)
	if v400&int32(65535) == int32(0) {
		goto L111
	} else {
		goto L112
	}
L111:
	;
	v406 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v395)+103)) = uint8(v406)
	goto L113
L112:
	;
	goto L113
L113:
	;
	F_CatalogTupleUpdate(m, v156, v391+int32(4), v391)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L2
	} else {
		goto L114
	}
L114:
	;
	F_pfree(m, v391)
	mBase = m.M
	v413 = m.ExcPending
	if v413 != 0 {
		goto L2
	} else {
		goto L115
	}
L115:
	;
	v422 = v386
	v423 = v387
	goto L68
L116:
	;
	if v427 != 0 {
		v265 = v427
		v267 = v422
		v268 = v423
		goto L66
	} else {
		goto L117
	}
L117:
	;
	goto L67
L118:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L2
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	F_systable_endscan(m, v255)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L2
	} else {
		goto L130
	}
L121:
	;
	if v437 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	v449 = v448
	goto L124
L123:
	;
	v449 = int32(0)
	goto L124
L124:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	if v438 != 0 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v438)+4))
	v454 = v452
	goto L127
L126:
	;
	v454 = int32(0)
	goto L127
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v454 + v449
	v457 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = v451 + v457
	*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v450 + v457
	F_errmsg_internal(m, int32(716478), v16)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L2
	} else {
		goto L128
	}
L128:
	;
	F_errfinish(m, int32(505917), int32(18136), int32(427263))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L2
	} else {
		goto L129
	}
L129:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L130:
	;
	F_sequence_close(m, v156, int32(3))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L2
	} else {
		goto L131
	}
L131:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	if v19 == int32(112) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v483 = int32(97)
	goto L134
L133:
	;
	v483 = int32(110)
	goto L134
L134:
	;
	F_drop_parent_dependency(m, v476, int32(1259), v478, v483)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L2
	} else {
		goto L135
	}
L135:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _consts[389]))
	if v487 != 0 {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v490 = int32(0)
	v491 = *(*int32)(unsafe.Add(mBase, uint32(l1)+56))
	F_RunObjectPostAlterHook(m, int32(2611), v489, v490, v491, v490)
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L2
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	m.G0 = v16 + int32(208)
	return
L139:
	;
	goto L138
L140:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v395 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v502
	F_errmsg_internal(m, int32(713185), v16+int32(16))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L2
	} else {
		goto L141
	}
L141:
	;
	F_errfinish(m, int32(505917), int32(18121), int32(427263))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L2
	} else {
		goto L142
	}
L142:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L143:
	;
	F_errfinish(m, int32(505917), int32(17976), int32(427263))
	mBase = m.M
	v534 = m.ExcPending
	if v534 != 0 {
		goto L2
	} else {
		goto L144
	}
L144:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ReportApplyConflict(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v56 int64
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
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
	var v272 int32
	_ = v272
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v418 int64
	_ = v418
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v428 int64
	_ = v428
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	v18 = m.G0
	v20 = v18 - int32(320)
	m.G0 = v20
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	F_initStringInfo(m, v20+int32(268))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l6 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v415 = int32(0)
	v417 = *(*int32)(unsafe.Add(mBase, _consts[832]))
	v418 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v417))))
	v420 = F_pgstat_prep_pending_entry(m, int32(5), v415, v418, v415)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L1
	} else {
		goto L125
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v29 <= int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v47 = int32(0)
	goto L6
L6:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v51+v47<<(uint(int32(2))%32))))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v55)+16))
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v55)+12)))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	F_initStringInfo(m, v20+int32(288))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	goto L3
L8:
	;
	switch l3 {
	case 0, 2, 6:
		goto L17
	case 1:
		goto L16
	case 3:
		goto L15
	case 4:
		goto L14
	case 5:
		goto L13
	default:
		goto L10
	}
L9:
	;
	if v60 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L10:
	;
	v236 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v236)+52))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v236)+56))
	F_initStringInfo(m, v20+int32(304))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L65
	}
L11:
	;
	v222 = F_timestamptz_to_str(m, v56)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L1
	} else {
		goto L63
	}
L12:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+52))
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v215)+56))
	F_initStringInfo(m, v20+int32(304))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L62
	}
L13:
	;
	F_appendStringInfoString(m, v20+int32(288), int32(661780))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L1
	} else {
		goto L61
	}
L14:
	;
	v167 = v57 & int32(65535)
	if v167 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L15:
	;
	F_appendStringInfoString(m, v20+int32(288), int32(661947))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L47
	}
L16:
	;
	v119 = v57 & int32(65535)
	if v119 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L17:
	;
	if v56 != int64(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v68 = v57 & int32(65535)
	if v68 == int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	v107 = F_get_rel_name(m, v59)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L32
	}
L21:
	;
	v71 = F_get_rel_name(m, v59)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v87 = F_replorigin_by_oid(m, v68, v20+int32(284))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L27
	}
L24:
	;
	v73 = F_timestamptz_to_str(m, v56)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+136)) = v73
	*(*int32)(unsafe.Add(mBase, uint32(v20)+132)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v71
	F_appendStringInfo(m, v20+int32(288), int32(616844), v20+int32(128))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	goto L10
L27:
	;
	v89 = F_get_rel_name(m, v59)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v87 == int32(0) {
		goto L11
	} else {
		goto L29
	}
L29:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v20)+284))
	v94 = F_timestamptz_to_str(m, v56)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+156)) = v94
	*(*int32)(unsafe.Add(mBase, uint32(v20)+152)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v20)+148)) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v20)+144)) = v89
	F_appendStringInfo(m, v20+int32(288), int32(617197), v20+int32(144))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	goto L10
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+116)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v20)+112)) = v107
	F_appendStringInfo(m, v20+int32(288), int32(591045), v20+int32(112))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L10
L34:
	;
	v122 = F_timestamptz_to_str(m, v56)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v135 = F_replorigin_by_oid(m, v119, v20+int32(284))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+180)) = v122
	*(*int32)(unsafe.Add(mBase, uint32(v20)+176)) = v58
	F_appendStringInfo(m, v20+int32(288), int32(616776), v20+int32(176))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	goto L12
L39:
	;
	if v135 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v20)+284))
	v138 = F_timestamptz_to_str(m, v56)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v150 = F_timestamptz_to_str(m, v56)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+200)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(v20)+196)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v20)+192)) = v137
	F_appendStringInfo(m, v20+int32(288), int32(617374), v20+int32(192))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	goto L12
L45:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+212)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v20)+208)) = v58
	F_appendStringInfo(m, v20+int32(288), int32(617012), v20+int32(208))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	goto L12
L47:
	;
	goto L12
L48:
	;
	v170 = F_timestamptz_to_str(m, v56)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v183 = F_replorigin_by_oid(m, v167, v20+int32(284))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+228)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v20)+224)) = v58
	F_appendStringInfo(m, v20+int32(288), int32(616708), v20+int32(224))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	goto L12
L53:
	;
	if v183 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v20)+284))
	v186 = F_timestamptz_to_str(m, v56)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v198 = F_timestamptz_to_str(m, v56)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+248)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v20)+244)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v20)+240)) = v185
	F_appendStringInfo(m, v20+int32(288), int32(617287), v20+int32(240))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	goto L12
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+260)) = v198
	*(*int32)(unsafe.Add(mBase, uint32(v20)+256)) = v58
	F_appendStringInfo(m, v20+int32(288), int32(616927), v20+int32(256))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	goto L12
L61:
	;
	goto L12
L62:
	;
	v261 = v215
	v262 = v216
	v263 = v217
	goto L9
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+168)) = v222
	*(*int32)(unsafe.Add(mBase, uint32(v20)+164)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v20)+160)) = v89
	F_appendStringInfo(m, v20+int32(288), int32(617097), v20+int32(160))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	goto L10
L65:
	;
	if base.Ui32(int32(6)) < base.Ui32(l3) {
		v261 = v236
		v262 = v237
		v263 = v238
		goto L9
	} else {
		goto L66
	}
L66:
	;
	if int32(1)<<(uint(l3)%32)&int32(69) == int32(0) {
		v261 = v236
		v262 = v237
		v263 = v238
		goto L9
	} else {
		goto L67
	}
L67:
	;
	v249 = F_build_index_value_desc(m, l0, v236, v60, v59)
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L1
	} else {
		goto L68
	}
L68:
	;
	if v249 == int32(0) {
		v261 = v236
		v262 = v237
		v263 = v238
		goto L9
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+96)) = v249
	F_appendStringInfo(m, v20+int32(304), int32(181917), v20+int32(96))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v261 = v236
	v262 = v237
	v263 = v238
	goto L9
L71:
	;
	if l5 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L72:
	;
	v268 = F_ExecBuildSlotValueDescription(m, v263, v60, v262, int32(0))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	if v268 == int32(0) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v20)+308))
	if v272 <= int32(0) {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	v282 = int32(182130)
	goto L77
L76:
	;
	F_appendStringInfoString(m, v20+int32(304), int32(759725))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+80)) = v268
	F_appendStringInfo(m, v20+int32(304), v282, v20+int32(80))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L79
	}
L78:
	;
	v282 = int32(182108)
	goto L77
L79:
	;
	goto L71
L80:
	;
	if l4 == int32(0) {
		goto L92
	} else {
		goto L93
	}
L81:
	;
	v294 = F_ExecGetInsertedCols(m, l1, l0)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v296 = F_ExecGetUpdatedCols(m, l1, l0)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v298 = F_bms_union(m, v294, v296)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	v300 = F_ExecBuildSlotValueDescription(m, v263, l5, v262, v298)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	if v300 == int32(0) {
		goto L80
	} else {
		goto L86
	}
L86:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v20)+308))
	if v304 <= int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v314 = int32(182166)
	goto L89
L88:
	;
	F_appendStringInfoString(m, v20+int32(304), int32(759725))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L90
	}
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+64)) = v300
	F_appendStringInfo(m, v20+int32(304), v314, v20-int32(-64))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L91
	}
L90:
	;
	v314 = int32(182152)
	goto L89
L91:
	;
	goto L80
L92:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v20)+308))
	if v360 == int32(0) {
		goto L114
	} else {
		goto L115
	}
L93:
	;
	v326 = F_GetRelationIdentityOrPK(m, v261)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	if v333 == int32(0) {
		goto L92
	} else {
		goto L101
	}
L95:
	;
	if v326 != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v328 = F_build_index_value_desc(m, l0, v261, l4, v326)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v331 = F_ExecBuildSlotValueDescription(m, v263, l4, v262, int32(0))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L100
	}
L99:
	;
	v333 = v328
	goto L94
L100:
	;
	v333 = v331
	goto L94
L101:
	;
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v20)+308))
	if int32(0) < v336 {
		goto L103
	} else {
		goto L104
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+48)) = v333
	F_appendStringInfo(m, v20+int32(304), v350, v20+int32(48))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L113
	}
L103:
	;
	F_appendStringInfoString(m, v20+int32(304), int32(759725))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	if v326 != 0 {
		goto L110
	} else {
		goto L111
	}
L106:
	;
	if v326 != 0 {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v346 = int32(181420)
	goto L109
L108:
	;
	v346 = int32(189842)
	goto L109
L109:
	;
	v350 = v346
	goto L102
L110:
	;
	v349 = int32(181440)
	goto L112
L111:
	;
	v349 = int32(189867)
	goto L112
L112:
	;
	v350 = v349
	goto L102
L113:
	;
	goto L92
L114:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v20)+272))
	if int32(0) < v380 {
		goto L119
	} else {
		goto L120
	}
L115:
	;
	F_appendStringInfoChar(m, v20+int32(304), int32(46))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v20)+304))
	if v368 == int32(0) {
		goto L114
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v368
	F_appendStringInfo(m, v20+int32(288), int32(210226), v20+int32(32))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	F_appendStringInfoChar(m, v20+int32(268), int32(10))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L1
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v20)+288))
	F_appendStringInfoString(m, v20+int32(268), v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L123
	}
L122:
	;
	goto L121
L123:
	;
	v394 = v47 + int32(1)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l6)+4))
	if v394 < v395 {
		v47 = v394
		goto L6
	} else {
		goto L124
	}
L124:
	;
	goto L7
L125:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v420)+12))
	v427 = v422 + l3<<(uint(int32(3))%32) + int32(16)
	v428 = *(*int64)(unsafe.Add(mBase, uint32(v427)))
	*(*int64)(unsafe.Add(mBase, uint32(v427))) = v428 + int64(1)
	v433 = F_errstart(m, l2, int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	if v433 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if base.Ui32(l3) <= base.Ui32(int32(6)) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	m.G0 = v20 + int32(320)
	return
L130:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_consts[833])))
	F_errcode(m, v441)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L1
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32))+uint32(_consts[834])))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v449)+68))
	v451 = F_get_namespace_name(m, v450)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L134
	}
L133:
	;
	goto L132
L134:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v448
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v451
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v453 + int32(4)
	F_errmsg(m, int32(180228), v20+int32(16))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v20)+268))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v464
	F_errdetail_internal(m, int32(210227), v20)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(504426), int32(130), int32(112148))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	goto L129
}
func F_ReservedPLKeywords_hash_func(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v13 int32
	_ = v13
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
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
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	v3 = int32(0)
	if l1 == v3 {
		v79 = int32(1)
		v87 = int32(0)
	} else {
		v13 = int32(1)
		if l1 == v13 {
			v58 = l0
			v59 = l1
			v60 = v13
			v61 = int32(8191)
			v66 = int32(0)
		} else {
			v23 = l0
			v24 = int32(0)
			v25 = v13
			v27 = v3
			for {
				v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
				v32 = int32(32)
				v33 = v31 | v32
				v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23))))
				v36 = v34 | v32
				v37 = int32(8191)
				v42 = v33 + (v36+v25*v37)*v37
				v43 = int32(257)
				v48 = (v24*v43+v36)*v43 + v33
				v49 = int32(2)
				v50 = v23 + v49
				v52 = v27 + v49
				if v52 != l1&int32(-2) {
					v23 = v50
					v24 = v48
					v25 = v42
					v27 = v52
					continue
				} else {
					break
				}
				break
			}
			v58 = v50
			v59 = v48
			v60 = v42
			v61 = v42 * int32(8191)
			v66 = v48 * int32(257)
		}
		if l1&v13 != 0 {
			v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58))))
			v69 = v67 | int32(32)
			v73 = v69 + v66
			v74 = v61 + v69
		} else {
			v73 = v59
			v74 = v60
		}
		v75 = int32(49)
		v76 = base.I32_rem_u_s(v74, v75)
		v78 = base.I32_rem_u_s(v73, v75)
		v79 = v76
		v87 = v78
	}
	v90 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79)+uint32(_consts[1471]))))
	v92 = int32(*(*int8)(unsafe.Add(mBase, uint32(v87)+uint32(_consts[1471]))))
	return v90 + v92
}
func F__readBitmapset(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v50 int64
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v11 = F_pg_strtok(m, v7+int32(44))
	mBase = m.M
	if v11 != 0 {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L23
	} else {
		goto L41
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L23
	} else {
		goto L38
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L23
	} else {
		goto L35
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L23
	} else {
		goto L32
	}
L5:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v12 != int32(1) {
		goto L4
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L23
	} else {
		goto L29
	}
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v15 != int32(40) {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	v20 = F_pg_strtok(m, v7+int32(44))
	mBase = m.M
	if v20 == int32(0) {
		goto L3
	} else {
		goto L10
	}
L10:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v23 != int32(1) {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v26 != int32(98) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v31 = F_pg_strtok(m, v7+int32(44))
	mBase = m.M
	if v31 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v33 = v31
	v34 = int32(0)
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L23
	} else {
		goto L26
	}
L16:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v36 != int32(1) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v50 = F_strtox_2(m, v33, v7+int32(40), int32(10), int64(2147483648))
	mBase = m.M
	goto L21
L19:
	;
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v39 != int32(41) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	m.G0 = v7 + int32(48)
	return v34
L21:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)+40))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	if v52 != v33+v53 {
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v56 = F_bms_add_member(m, v34, base.I32_wrap_i64(v50))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	return int32(0)
L24:
	;
	v62 = F_pg_strtok(m, v7+int32(44))
	mBase = m.M
	if v62 != 0 {
		v33 = v62
		v34 = v56
		goto L16
	} else {
		goto L25
	}
L25:
	;
	goto L17
L26:
	;
	F_errmsg_internal(m, int32(371020), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L23
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(506354), int32(234), int32(107702))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_errmsg_internal(m, int32(370989), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L23
	} else {
		goto L30
	}
L30:
	;
	F_errfinish(m, int32(506354), int32(217), int32(107702))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L23
	} else {
		goto L31
	}
L31:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v11
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v98
	F_errmsg_internal(m, int32(704385), v7+int32(32))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L23
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(506354), int32(219), int32(107702))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L23
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L35:
	;
	F_errmsg_internal(m, int32(370989), int32(0))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L23
	} else {
		goto L36
	}
L36:
	;
	F_errfinish(m, int32(506354), int32(223), int32(107702))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L23
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
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v33
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v128
	F_errmsg_internal(m, int32(704314), v7)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L23
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(506354), int32(239), int32(107702))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L23
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v20
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v7)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v143
	F_errmsg_internal(m, int32(704385), v7+int32(16))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L23
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(506354), int32(225), int32(107702))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L23
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_r_SUFFIX_KAN_OK(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	return base.B2i32(v3&int32(-2) != int32(2))
}
func F_r_VI_1(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v4 <= v5 {
		v72 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7+v4-int32(1)))))
		if v11 != int32(105) {
			v72 = v2
		} else {
			v15 = v4 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v15 <= v26 {
				v69 = int32(-1)
			} else {
				v38 = int32(1)
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39+v15-v38))))
				if int32(246) < v43 {
					v65 = v38
				} else {
					v45 = v43 - int32(97)
					if v45 < int32(0) {
						v65 = v38
					} else {
						v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v45)>>(uint(int32(3))%32)))+uint32(_consts[1473]))))
						if int32(base.Ui32(v51)>>(uint(v45&int32(7))%32))&int32(1) == int32(0) {
							v65 = v38
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v15 - int32(1)
							v65 = int32(0)
						}
					}
				}
				v69 = v65
			}
			v72 = base.B2i32(v69 == int32(0))
		}
	}
	return v72
}
func F_r_fix_chdz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v4
	v7 = v4 - int32(1)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v7 <= v8 {
		v45 = v2
		return v45
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10+v7))))
		if base.B2i32(v12 != int32(190))&base.B2i32(v12 != int32(141)) != 0 {
			v45 = v2
			return v45
		} else {
			v20 = F_find_among_b(m, l0, int32(4330928), int32(2))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				if v20 == int32(0) {
					v45 = v2
					return v45
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v26
					switch v20 - int32(1) {
					case 0:
						v32 = F_slice_from_s(m, l0, int32(1), int32(2226476))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v32 {
								v45 = int32(1)
							} else {
								v45 = v32
							}
							return v45
						}
					case 1:
						v38 = F_slice_from_s(m, l0, int32(1), int32(2226477))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							if v38 < int32(0) {
								v45 = v38
							} else {
								v45 = int32(1)
							}
							return v45
						}
					default:
						v45 = int32(1)
						return v45
					}
				}
			}
		}
	}
}
func F_r_fix_ending(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	v2 = int32(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v7-int32(4))))
	if v15 == v2 {
		v87 = int32(0)
	} else {
		v20 = v15 & int32(3)
		if base.Ui32(v15) < base.Ui32(int32(4)) {
			v54 = v7
			v55 = int32(0)
		} else {
			v27 = v7
			v28 = int32(0)
			v31 = v2
			for {
				v33 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27))))
				v34 = int32(-65)
				v37 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+1)))
				v41 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+2)))
				v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v27)+3)))
				v48 = v28 + base.B2i32(v34 < v33) + base.B2i32(v34 < v37) + base.B2i32(v34 < v41) + base.B2i32(v34 < v45)
				v49 = int32(4)
				v50 = v27 + v49
				v52 = v31 + v49
				if v52 != v15&int32(-4) {
					v27 = v50
					v28 = v48
					v31 = v52
					continue
				} else {
					break
				}
				break
			}
			v54 = v50
			v55 = v48
		}
		if v20 != 0 {
			v60 = v54
			v61 = v55
			v63 = v2
			for {
				v66 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60))))
				v69 = v61 + base.B2i32(int32(-65) < v66)
				v70 = int32(1)
				v73 = v63 + v70
				if v73 != v20 {
					v60 = v60 + v70
					v61 = v69
					v63 = v73
					continue
				} else {
					break
				}
				break
			}
			v76 = v69
		} else {
			v76 = v55
		}
		v87 = v76
	}
	if v87 < int32(4) {
		v346 = v2
		return v346
	} else {
		v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v90
		v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v92
		*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v92
		v97 = F_find_among_b(m, l0, int32(4397488), int32(17))
		mBase = m.M
		v100 = m.ExcPending
		if v100 != 0 {
			return int32(0)
		} else {
			if v97 == int32(0) {
				v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
				*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
				v226 = int32(3)
				v228 = int32(0)
				v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v223-v231 < v226 {
					v241 = v228
				} else {
					v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v237 = F_memcmp(m, v234+v223-v226, int32(2244581), v226)
					mBase = m.M
					if v237 != 0 {
						v241 = v228
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223 - v226
						v241 = int32(1)
					}
				}
				if v241 == int32(0) {
					v346 = v2
					return v346
				} else {
					v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v248 = F_find_among_b(m, l0, int32(4398128), int32(6))
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return int32(0)
					} else {
						v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						if v248 != 0 {
							v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v252 = v250 - v251
							v253 = int32(3)
							v255 = int32(0)
							v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v251-v258 < v253 {
								v268 = v255
							} else {
								v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v264 = F_memcmp(m, v261+v251-v253, int32(2244584), v253)
								mBase = m.M
								if v264 != 0 {
									v268 = v255
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v251 - v253
									v268 = int32(1)
								}
							}
							if v268 == int32(0) {
								v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v280 = v271 - v252
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
								v282 = v280
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
								v284 = F_slice_del(m, l0)
								mBase = m.M
								v285 = m.ExcPending
								if v285 != 0 {
									return int32(0)
								} else {
									if int32(0) <= v284 {
										v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
										v346 = int32(1)
									} else {
										v346 = v284
									}
									return v346
								}
							} else {
								v275 = F_find_among_b(m, l0, int32(4398256), int32(6))
								mBase = m.M
								v276 = m.ExcPending
								if v276 != 0 {
									return int32(0)
								} else {
									if v275 != 0 {
										v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v282 = v277
									} else {
										v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v280 = v278 - v252
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
										v282 = v280
									}
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
									v284 = F_slice_del(m, l0)
									mBase = m.M
									v285 = m.ExcPending
									if v285 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v284 {
											v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
											v346 = int32(1)
										} else {
											v346 = v284
										}
										return v346
									}
								}
							}
						} else {
							v288 = v245 - v244
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v250 - v288
							v293 = F_find_among_b(m, l0, int32(4398384), int32(11))
							mBase = m.M
							v294 = m.ExcPending
							if v294 != 0 {
								return int32(0)
							} else {
								if v293 == int32(0) {
									v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
									v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
									mBase = m.M
									v327 = m.ExcPending
									if v327 != 0 {
										return int32(0)
									} else {
										if v326 == int32(0) {
											v346 = v2
											return v346
										} else {
											v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v331 = v330 - v288
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
											v334 = F_slice_del(m, l0)
											mBase = m.M
											v335 = m.ExcPending
											if v335 != 0 {
												return int32(0)
											} else {
												if v334 < int32(0) {
													v346 = v334
												} else {
													v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
													v346 = int32(1)
												}
												return v346
											}
										}
									}
								} else {
									v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v297
									v299 = int32(3)
									v301 = int32(0)
									v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v297-v304 < v299 {
										v314 = v301
									} else {
										v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v310 = F_memcmp(m, v307+v297-v299, int32(2244587), v299)
										mBase = m.M
										if v310 != 0 {
											v314 = v301
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v297 - v299
											v314 = int32(1)
										}
									}
									if v314 == int32(0) {
										v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
										v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
										mBase = m.M
										v327 = m.ExcPending
										if v327 != 0 {
											return int32(0)
										} else {
											if v326 == int32(0) {
												v346 = v2
												return v346
											} else {
												v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v331 = v330 - v288
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
												v334 = F_slice_del(m, l0)
												mBase = m.M
												v335 = m.ExcPending
												if v335 != 0 {
													return int32(0)
												} else {
													if v334 < int32(0) {
														v346 = v334
													} else {
														v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
														v346 = int32(1)
													}
													return v346
												}
											}
										}
									} else {
										v317 = F_slice_del(m, l0)
										mBase = m.M
										v318 = m.ExcPending
										if v318 != 0 {
											return int32(0)
										} else {
											if int32(0) <= v317 {
												v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
												v346 = int32(1)
											} else {
												v346 = v317
											}
											return v346
										}
									}
								}
							}
						}
					}
				}
			} else {
				v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v103
				switch v97 - int32(1) {
				case 0:
					v107 = F_slice_del(m, l0)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v107 {
							v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
							v346 = int32(1)
						} else {
							v346 = v107
						}
						return v346
					}
				case 1:
					v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v114 = F_find_among_b(m, l0, int32(4397840), int32(3))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						if v114 == int32(0) {
							v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
							v226 = int32(3)
							v228 = int32(0)
							v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v223-v231 < v226 {
								v241 = v228
							} else {
								v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v237 = F_memcmp(m, v234+v223-v226, int32(2244581), v226)
								mBase = m.M
								if v237 != 0 {
									v241 = v228
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223 - v226
									v241 = int32(1)
								}
							}
							if v241 == int32(0) {
								v346 = v2
								return v346
							} else {
								v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v248 = F_find_among_b(m, l0, int32(4398128), int32(6))
								mBase = m.M
								v249 = m.ExcPending
								if v249 != 0 {
									return int32(0)
								} else {
									v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v248 != 0 {
										v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v252 = v250 - v251
										v253 = int32(3)
										v255 = int32(0)
										v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v251-v258 < v253 {
											v268 = v255
										} else {
											v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v264 = F_memcmp(m, v261+v251-v253, int32(2244584), v253)
											mBase = m.M
											if v264 != 0 {
												v268 = v255
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v251 - v253
												v268 = int32(1)
											}
										}
										if v268 == int32(0) {
											v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v280 = v271 - v252
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
											v282 = v280
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
											v284 = F_slice_del(m, l0)
											mBase = m.M
											v285 = m.ExcPending
											if v285 != 0 {
												return int32(0)
											} else {
												if int32(0) <= v284 {
													v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
													v346 = int32(1)
												} else {
													v346 = v284
												}
												return v346
											}
										} else {
											v275 = F_find_among_b(m, l0, int32(4398256), int32(6))
											mBase = m.M
											v276 = m.ExcPending
											if v276 != 0 {
												return int32(0)
											} else {
												if v275 != 0 {
													v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v282 = v277
												} else {
													v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v280 = v278 - v252
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
													v282 = v280
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
												v284 = F_slice_del(m, l0)
												mBase = m.M
												v285 = m.ExcPending
												if v285 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v284 {
														v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
														v346 = int32(1)
													} else {
														v346 = v284
													}
													return v346
												}
											}
										}
									} else {
										v288 = v245 - v244
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v250 - v288
										v293 = F_find_among_b(m, l0, int32(4398384), int32(11))
										mBase = m.M
										v294 = m.ExcPending
										if v294 != 0 {
											return int32(0)
										} else {
											if v293 == int32(0) {
												v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
												v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
												mBase = m.M
												v327 = m.ExcPending
												if v327 != 0 {
													return int32(0)
												} else {
													if v326 == int32(0) {
														v346 = v2
														return v346
													} else {
														v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v331 = v330 - v288
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
														v334 = F_slice_del(m, l0)
														mBase = m.M
														v335 = m.ExcPending
														if v335 != 0 {
															return int32(0)
														} else {
															if v334 < int32(0) {
																v346 = v334
															} else {
																v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
																v346 = int32(1)
															}
															return v346
														}
													}
												}
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v297
												v299 = int32(3)
												v301 = int32(0)
												v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v297-v304 < v299 {
													v314 = v301
												} else {
													v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v310 = F_memcmp(m, v307+v297-v299, int32(2244587), v299)
													mBase = m.M
													if v310 != 0 {
														v314 = v301
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v297 - v299
														v314 = int32(1)
													}
												}
												if v314 == int32(0) {
													v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
													v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
													mBase = m.M
													v327 = m.ExcPending
													if v327 != 0 {
														return int32(0)
													} else {
														if v326 == int32(0) {
															v346 = v2
															return v346
														} else {
															v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v331 = v330 - v288
															*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
															v334 = F_slice_del(m, l0)
															mBase = m.M
															v335 = m.ExcPending
															if v335 != 0 {
																return int32(0)
															} else {
																if v334 < int32(0) {
																	v346 = v334
																} else {
																	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
																	v346 = int32(1)
																}
																return v346
															}
														}
													}
												} else {
													v317 = F_slice_del(m, l0)
													mBase = m.M
													v318 = m.ExcPending
													if v318 != 0 {
														return int32(0)
													} else {
														if int32(0) <= v317 {
															v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
															v346 = int32(1)
														} else {
															v346 = v317
														}
														return v346
													}
												}
											}
										}
									}
								}
							}
						} else {
							v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v118 + (v103 - v111)
							v122 = F_slice_del(m, l0)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v122 {
									v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
									v346 = int32(1)
								} else {
									v346 = v122
								}
								return v346
							}
						}
					}
				case 2:
					v128 = F_slice_from_s(m, l0, int32(6), int32(2244545))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v128 {
							v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
							v346 = int32(1)
						} else {
							v346 = v128
						}
						return v346
					}
				case 3:
					v134 = F_slice_from_s(m, l0, int32(6), int32(2244551))
					mBase = m.M
					v135 = m.ExcPending
					if v135 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v134 {
							v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
							v346 = int32(1)
						} else {
							v346 = v134
						}
						return v346
					}
				case 4:
					v140 = F_slice_from_s(m, l0, int32(6), int32(2244557))
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v140 {
							v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
							v346 = int32(1)
						} else {
							v346 = v140
						}
						return v346
					}
				case 5:
					v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
					if v145 == int32(0) {
						v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
						*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
						v226 = int32(3)
						v228 = int32(0)
						v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v223-v231 < v226 {
							v241 = v228
						} else {
							v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v237 = F_memcmp(m, v234+v223-v226, int32(2244581), v226)
							mBase = m.M
							if v237 != 0 {
								v241 = v228
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223 - v226
								v241 = int32(1)
							}
						}
						if v241 == int32(0) {
							v346 = v2
							return v346
						} else {
							v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v248 = F_find_among_b(m, l0, int32(4398128), int32(6))
							mBase = m.M
							v249 = m.ExcPending
							if v249 != 0 {
								return int32(0)
							} else {
								v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								if v248 != 0 {
									v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									v252 = v250 - v251
									v253 = int32(3)
									v255 = int32(0)
									v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									if v251-v258 < v253 {
										v268 = v255
									} else {
										v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
										v264 = F_memcmp(m, v261+v251-v253, int32(2244584), v253)
										mBase = m.M
										if v264 != 0 {
											v268 = v255
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v251 - v253
											v268 = int32(1)
										}
									}
									if v268 == int32(0) {
										v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										v280 = v271 - v252
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
										v282 = v280
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
										v284 = F_slice_del(m, l0)
										mBase = m.M
										v285 = m.ExcPending
										if v285 != 0 {
											return int32(0)
										} else {
											if int32(0) <= v284 {
												v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
												v346 = int32(1)
											} else {
												v346 = v284
											}
											return v346
										}
									} else {
										v275 = F_find_among_b(m, l0, int32(4398256), int32(6))
										mBase = m.M
										v276 = m.ExcPending
										if v276 != 0 {
											return int32(0)
										} else {
											if v275 != 0 {
												v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												v282 = v277
											} else {
												v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												v280 = v278 - v252
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
												v282 = v280
											}
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
											v284 = F_slice_del(m, l0)
											mBase = m.M
											v285 = m.ExcPending
											if v285 != 0 {
												return int32(0)
											} else {
												if int32(0) <= v284 {
													v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
													v346 = int32(1)
												} else {
													v346 = v284
												}
												return v346
											}
										}
									}
								} else {
									v288 = v245 - v244
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v250 - v288
									v293 = F_find_among_b(m, l0, int32(4398384), int32(11))
									mBase = m.M
									v294 = m.ExcPending
									if v294 != 0 {
										return int32(0)
									} else {
										if v293 == int32(0) {
											v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
											v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
											mBase = m.M
											v327 = m.ExcPending
											if v327 != 0 {
												return int32(0)
											} else {
												if v326 == int32(0) {
													v346 = v2
													return v346
												} else {
													v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v331 = v330 - v288
													*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
													v334 = F_slice_del(m, l0)
													mBase = m.M
													v335 = m.ExcPending
													if v335 != 0 {
														return int32(0)
													} else {
														if v334 < int32(0) {
															v346 = v334
														} else {
															v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
															v346 = int32(1)
														}
														return v346
													}
												}
											}
										} else {
											v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v297
											v299 = int32(3)
											v301 = int32(0)
											v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											if v297-v304 < v299 {
												v314 = v301
											} else {
												v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
												v310 = F_memcmp(m, v307+v297-v299, int32(2244587), v299)
												mBase = m.M
												if v310 != 0 {
													v314 = v301
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v297 - v299
													v314 = int32(1)
												}
											}
											if v314 == int32(0) {
												v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
												v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
												mBase = m.M
												v327 = m.ExcPending
												if v327 != 0 {
													return int32(0)
												} else {
													if v326 == int32(0) {
														v346 = v2
														return v346
													} else {
														v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v331 = v330 - v288
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
														v334 = F_slice_del(m, l0)
														mBase = m.M
														v335 = m.ExcPending
														if v335 != 0 {
															return int32(0)
														} else {
															if v334 < int32(0) {
																v346 = v334
															} else {
																v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
																v346 = int32(1)
															}
															return v346
														}
													}
												}
											} else {
												v317 = F_slice_del(m, l0)
												mBase = m.M
												v318 = m.ExcPending
												if v318 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v317 {
														v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
														v346 = int32(1)
													} else {
														v346 = v317
													}
													return v346
												}
											}
										}
									}
								}
							}
						}
					} else {
						v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						v149 = int32(3)
						v151 = int32(0)
						v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						if v153-v154 < v149 {
							v164 = v151
						} else {
							v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v160 = F_memcmp(m, v157+v153-v149, int32(2244563), v149)
							mBase = m.M
							if v160 != 0 {
								v164 = v151
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v153 - v149
								v164 = int32(1)
							}
						}
						if v164 != 0 {
							v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
							v226 = int32(3)
							v228 = int32(0)
							v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v223-v231 < v226 {
								v241 = v228
							} else {
								v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v237 = F_memcmp(m, v234+v223-v226, int32(2244581), v226)
								mBase = m.M
								if v237 != 0 {
									v241 = v228
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223 - v226
									v241 = int32(1)
								}
							}
							if v241 == int32(0) {
								v346 = v2
								return v346
							} else {
								v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v248 = F_find_among_b(m, l0, int32(4398128), int32(6))
								mBase = m.M
								v249 = m.ExcPending
								if v249 != 0 {
									return int32(0)
								} else {
									v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v248 != 0 {
										v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v252 = v250 - v251
										v253 = int32(3)
										v255 = int32(0)
										v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v251-v258 < v253 {
											v268 = v255
										} else {
											v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v264 = F_memcmp(m, v261+v251-v253, int32(2244584), v253)
											mBase = m.M
											if v264 != 0 {
												v268 = v255
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v251 - v253
												v268 = int32(1)
											}
										}
										if v268 == int32(0) {
											v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v280 = v271 - v252
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
											v282 = v280
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
											v284 = F_slice_del(m, l0)
											mBase = m.M
											v285 = m.ExcPending
											if v285 != 0 {
												return int32(0)
											} else {
												if int32(0) <= v284 {
													v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
													v346 = int32(1)
												} else {
													v346 = v284
												}
												return v346
											}
										} else {
											v275 = F_find_among_b(m, l0, int32(4398256), int32(6))
											mBase = m.M
											v276 = m.ExcPending
											if v276 != 0 {
												return int32(0)
											} else {
												if v275 != 0 {
													v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v282 = v277
												} else {
													v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v280 = v278 - v252
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
													v282 = v280
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
												v284 = F_slice_del(m, l0)
												mBase = m.M
												v285 = m.ExcPending
												if v285 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v284 {
														v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
														v346 = int32(1)
													} else {
														v346 = v284
													}
													return v346
												}
											}
										}
									} else {
										v288 = v245 - v244
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v250 - v288
										v293 = F_find_among_b(m, l0, int32(4398384), int32(11))
										mBase = m.M
										v294 = m.ExcPending
										if v294 != 0 {
											return int32(0)
										} else {
											if v293 == int32(0) {
												v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
												v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
												mBase = m.M
												v327 = m.ExcPending
												if v327 != 0 {
													return int32(0)
												} else {
													if v326 == int32(0) {
														v346 = v2
														return v346
													} else {
														v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v331 = v330 - v288
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
														v334 = F_slice_del(m, l0)
														mBase = m.M
														v335 = m.ExcPending
														if v335 != 0 {
															return int32(0)
														} else {
															if v334 < int32(0) {
																v346 = v334
															} else {
																v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
																v346 = int32(1)
															}
															return v346
														}
													}
												}
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v297
												v299 = int32(3)
												v301 = int32(0)
												v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v297-v304 < v299 {
													v314 = v301
												} else {
													v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v310 = F_memcmp(m, v307+v297-v299, int32(2244587), v299)
													mBase = m.M
													if v310 != 0 {
														v314 = v301
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v297 - v299
														v314 = int32(1)
													}
												}
												if v314 == int32(0) {
													v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
													v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
													mBase = m.M
													v327 = m.ExcPending
													if v327 != 0 {
														return int32(0)
													} else {
														if v326 == int32(0) {
															v346 = v2
															return v346
														} else {
															v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v331 = v330 - v288
															*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
															v334 = F_slice_del(m, l0)
															mBase = m.M
															v335 = m.ExcPending
															if v335 != 0 {
																return int32(0)
															} else {
																if v334 < int32(0) {
																	v346 = v334
																} else {
																	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
																	v346 = int32(1)
																}
																return v346
															}
														}
													}
												} else {
													v317 = F_slice_del(m, l0)
													mBase = m.M
													v318 = m.ExcPending
													if v318 != 0 {
														return int32(0)
													} else {
														if int32(0) <= v317 {
															v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
															v346 = int32(1)
														} else {
															v346 = v317
														}
														return v346
													}
												}
											}
										}
									}
								}
							}
						} else {
							v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v165 + (v103 - v148)
							v171 = F_slice_from_s(m, l0, int32(6), int32(2244566))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v171 {
									v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
									v346 = int32(1)
								} else {
									v346 = v171
								}
								return v346
							}
						}
					}
				case 6:
					v177 = F_slice_from_s(m, l0, int32(3), int32(2244572))
					mBase = m.M
					v178 = m.ExcPending
					if v178 != 0 {
						return int32(0)
					} else {
						if int32(0) <= v177 {
							v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
							v346 = int32(1)
						} else {
							v346 = v177
						}
						return v346
					}
				case 7:
					v181 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v184 = F_find_among_b(m, l0, int32(4397904), int32(8))
					mBase = m.M
					v185 = m.ExcPending
					if v185 != 0 {
						return int32(0)
					} else {
						if v184 != 0 {
							v223 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v223
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223
							v226 = int32(3)
							v228 = int32(0)
							v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v223-v231 < v226 {
								v241 = v228
							} else {
								v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v237 = F_memcmp(m, v234+v223-v226, int32(2244581), v226)
								mBase = m.M
								if v237 != 0 {
									v241 = v228
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v223 - v226
									v241 = int32(1)
								}
							}
							if v241 == int32(0) {
								v346 = v2
								return v346
							} else {
								v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								v248 = F_find_among_b(m, l0, int32(4398128), int32(6))
								mBase = m.M
								v249 = m.ExcPending
								if v249 != 0 {
									return int32(0)
								} else {
									v250 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									if v248 != 0 {
										v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
										v252 = v250 - v251
										v253 = int32(3)
										v255 = int32(0)
										v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
										if v251-v258 < v253 {
											v268 = v255
										} else {
											v261 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
											v264 = F_memcmp(m, v261+v251-v253, int32(2244584), v253)
											mBase = m.M
											if v264 != 0 {
												v268 = v255
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v251 - v253
												v268 = int32(1)
											}
										}
										if v268 == int32(0) {
											v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											v280 = v271 - v252
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
											v282 = v280
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
											v284 = F_slice_del(m, l0)
											mBase = m.M
											v285 = m.ExcPending
											if v285 != 0 {
												return int32(0)
											} else {
												if int32(0) <= v284 {
													v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
													v346 = int32(1)
												} else {
													v346 = v284
												}
												return v346
											}
										} else {
											v275 = F_find_among_b(m, l0, int32(4398256), int32(6))
											mBase = m.M
											v276 = m.ExcPending
											if v276 != 0 {
												return int32(0)
											} else {
												if v275 != 0 {
													v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
													v282 = v277
												} else {
													v278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													v280 = v278 - v252
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
													v282 = v280
												}
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v282
												v284 = F_slice_del(m, l0)
												mBase = m.M
												v285 = m.ExcPending
												if v285 != 0 {
													return int32(0)
												} else {
													if int32(0) <= v284 {
														v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
														v346 = int32(1)
													} else {
														v346 = v284
													}
													return v346
												}
											}
										}
									} else {
										v288 = v245 - v244
										*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v250 - v288
										v293 = F_find_among_b(m, l0, int32(4398384), int32(11))
										mBase = m.M
										v294 = m.ExcPending
										if v294 != 0 {
											return int32(0)
										} else {
											if v293 == int32(0) {
												v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
												v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
												mBase = m.M
												v327 = m.ExcPending
												if v327 != 0 {
													return int32(0)
												} else {
													if v326 == int32(0) {
														v346 = v2
														return v346
													} else {
														v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
														v331 = v330 - v288
														*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
														v334 = F_slice_del(m, l0)
														mBase = m.M
														v335 = m.ExcPending
														if v335 != 0 {
															return int32(0)
														} else {
															if v334 < int32(0) {
																v346 = v334
															} else {
																v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
																v346 = int32(1)
															}
															return v346
														}
													}
												}
											} else {
												v297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v297
												v299 = int32(3)
												v301 = int32(0)
												v304 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
												if v297-v304 < v299 {
													v314 = v301
												} else {
													v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
													v310 = F_memcmp(m, v307+v297-v299, int32(2244587), v299)
													mBase = m.M
													if v310 != 0 {
														v314 = v301
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v297 - v299
														v314 = int32(1)
													}
												}
												if v314 == int32(0) {
													v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
													*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v321 - v288
													v326 = F_find_among_b(m, l0, int32(4398608), int32(9))
													mBase = m.M
													v327 = m.ExcPending
													if v327 != 0 {
														return int32(0)
													} else {
														if v326 == int32(0) {
															v346 = v2
															return v346
														} else {
															v330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
															v331 = v330 - v288
															*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v331
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v331
															v334 = F_slice_del(m, l0)
															mBase = m.M
															v335 = m.ExcPending
															if v335 != 0 {
																return int32(0)
															} else {
																if v334 < int32(0) {
																	v346 = v334
																} else {
																	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
																	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
																	v346 = int32(1)
																}
																return v346
															}
														}
													}
												} else {
													v317 = F_slice_del(m, l0)
													mBase = m.M
													v318 = m.ExcPending
													if v318 != 0 {
														return int32(0)
													} else {
														if int32(0) <= v317 {
															v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
															*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
															v346 = int32(1)
														} else {
															v346 = v317
														}
														return v346
													}
												}
											}
										}
									}
								}
							}
						} else {
							v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v186 + (v103 - v181)
							v190 = F_slice_del(m, l0)
							mBase = m.M
							v191 = m.ExcPending
							if v191 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v190 {
									v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
									v346 = int32(1)
								} else {
									v346 = v190
								}
								return v346
							}
						}
					}
				case 8:
					v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v103-int32(2) <= v194 {
						v217 = F_slice_from_s(m, l0, int32(6), int32(2244575))
						mBase = m.M
						v218 = m.ExcPending
						if v218 != 0 {
							return int32(0)
						} else {
							if int32(0) <= v217 {
								v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
								v346 = int32(1)
							} else {
								v346 = v217
							}
							return v346
						}
					} else {
						v198 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v103-int32(1)))))
						switch v202 - int32(136) {
						case 0, 5:
							v207 = F_find_among_b(m, l0, int32(4398064), int32(3))
							mBase = m.M
							v208 = m.ExcPending
							if v208 != 0 {
								return int32(0)
							} else {
								switch v207 - int32(1) {
								case 0:
									v211 = F_slice_del(m, l0)
									mBase = m.M
									v212 = m.ExcPending
									if v212 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v211 {
											v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
											v346 = int32(1)
										} else {
											v346 = v211
										}
										return v346
									}
								case 1:
									v217 = F_slice_from_s(m, l0, int32(6), int32(2244575))
									mBase = m.M
									v218 = m.ExcPending
									if v218 != 0 {
										return int32(0)
									} else {
										if int32(0) <= v217 {
											v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
											*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
											v346 = int32(1)
										} else {
											v346 = v217
										}
										return v346
									}
								default:
									v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
									v346 = int32(1)
									return v346
								}
							}
						default:
							v217 = F_slice_from_s(m, l0, int32(6), int32(2244575))
							mBase = m.M
							v218 = m.ExcPending
							if v218 != 0 {
								return int32(0)
							} else {
								if int32(0) <= v217 {
									v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
									v346 = int32(1)
								} else {
									v346 = v217
								}
								return v346
							}
						}
					}
				default:
					v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v343
					v346 = int32(1)
					return v346
				}
			}
		}
	}
}
func F_radians(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v7 float64
	_ = v7
	var v9 float64
	_ = v9
	var v17 float64
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v7 = base.F64_mul(v5, float64(0.017453292519943295))
	v9 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_eq(base.F64_abs(v7), v9)&base.F64_ne(base.F64_abs(v5), v9) == int32(0) {
		v17 = float64(0)
		if base.F64_eq(v7, v17)&base.F64_ne(v5, v17) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v22 = F_Float8GetDatum(m, v7)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				return v22
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_raise(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	v4 = m.G0
	v6 = v4 - int32(128)
	m.G0 = v6
	v11 = l0 - int32(1)
	if base.Ui32(v11) <= base.Ui32(int32(63)) {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(int32(base.Ui32(v11)>>(uint(int32(3))%32))&int32(536870908))+uint32(_consts[1611])))
		v23 = int32(base.Ui32(v19)>>(uint(v11)%32)) & int32(1)
	} else {
		v23 = int32(0)
	}
	if v23 != 0 {
		v27 = l0 - int32(1)
		if base.B2i32(base.Ui32(v27) <= base.Ui32(int32(63)))&base.B2i32(base.Ui32(int32(2)) < base.Ui32(l0-int32(32))) == int32(0) {
			*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(28)
		} else {
			v43 = int32(base.Ui32(v27)>>(uint(int32(3))%32)) & int32(536870908)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[1612])))
			*(*int32)(unsafe.Add(mBase, uint32(v43)+uint32(_consts[1612]))) = v45 | int32(1)<<(uint(v27)%32)
		}
		m.G0 = v6 + int32(128)
		return int32(0)
	} else {
		v52 = l0 * int32(140)
		v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1613]))))
		if v55&int32(4) != 0 {
			v61 = F__emscripten_memset_bulkmem(m, v6, base.I32_extend8_s(int32(0)), int32(128))
			mBase = m.M
			v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1614])))
			m.T0[v63].(func(*base.Module, int32, int32, int32))(m, l0, v61, int32(0))
			mBase = m.M
			v67 = m.ExcPending
			if v67 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(128)
				return int32(0)
			}
		} else {
			v68 = *(*int32)(unsafe.Add(mBase, uint32(v52)+uint32(_consts[1614])))
			switch v68 + int32(2) {
			case 0:
				m.G0 = v6 + int32(128)
				return int32(0)
			default:
				m.Env.X__call_sighandler(m, v68, l0)
				mBase = m.M
				m.G0 = v6 + int32(128)
				return int32(0)
			case 2:
				v75 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[1615])))
				if v75 == int32(0) {
					m.G0 = v6 + int32(128)
					return int32(0)
				} else {
					m.T0[v75].(func(*base.Module, int32))(m, l0)
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(128)
						return int32(0)
					}
				}
			}
		}
	}
}
func F_rangeTableEntry_used_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
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
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	v3 = int32(0)
	if l0 == v3 {
		v63 = v3
		return v63
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v7 - int32(58) {
		case 0:
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if v28 != 0 {
				v63 = v3
				return v63
			} else {
				v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if v29 != v30 {
					v63 = v3
					return v63
				} else {
					return int32(1)
				}
			}
		case 1, 2, 3, 4, 7, 8:
			v60 = F_expression_tree_walker_impl(m, l0, int32(1051), l1)
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v63 = v60
				return v63
			}
		case 5:
			v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v34 != v35 {
				v63 = v3
				return v63
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v37 != 0 {
					v63 = v3
					return v63
				} else {
					return int32(1)
				}
			}
		case 6:
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			if v40 != v41 {
				v60 = F_expression_tree_walker_impl(m, l0, int32(1051), l1)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v63 = v60
					return v63
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v43 != 0 {
					v60 = F_expression_tree_walker_impl(m, l0, int32(1051), l1)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						v63 = v60
						return v63
					}
				} else {
					return int32(1)
				}
			}
		case 9:
			v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v46 + int32(1)
			v52 = F_query_tree_walker_impl(m, l0, int32(1051), l1, int32(0))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v54 - int32(1)
				return v52
			}
		default:
			if v7 != int32(6) {
				v60 = F_expression_tree_walker_impl(m, l0, int32(1051), l1)
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v63 = v60
					return v63
				}
			} else {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				if v12 == v13 {
					v15 = int32(1)
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v16 == v17 {
						v63 = v15
						return v63
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v20 = F_bms_is_member(m, v16, v19)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							if v20 != 0 {
								v63 = v15
								return v63
							} else {
								return int32(0)
							}
						}
					}
				} else {
					return int32(0)
				}
			}
		}
	}
}
func F_rbt_left_right_iterator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v44
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v12 = v9
	goto L5
L3:
	;
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
	if v18 != int32(4141896) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if v15 != int32(4141896) {
		v12 = v15
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v44 = v12
	goto L1
L7:
	;
	goto L6
L8:
	;
	v23 = v18
	goto L11
L9:
	;
	goto L10
L10:
	;
	v32 = v5
	goto L14
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v26 != int32(4141896) {
		v23 = v26
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v44 = v23
	goto L1
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33
	if v33 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v44 = v33
	goto L1
L16:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v37)
	return int32(0)
L17:
	;
	goto L18
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v41 != v32 {
		v32 = v33
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
}
func F_rbt_right_left_iterator(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v5 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v44
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v12 = v9
	goto L5
L3:
	;
	goto L4
L4:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v18 != int32(4141896) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v12
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	if v15 != int32(4141896) {
		v12 = v15
		goto L5
	} else {
		goto L7
	}
L6:
	;
	v44 = v12
	goto L1
L7:
	;
	goto L6
L8:
	;
	v23 = v18
	goto L11
L9:
	;
	goto L10
L10:
	;
	v32 = v5
	goto L14
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v23
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	if v26 != int32(4141896) {
		v23 = v26
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v44 = v23
	goto L1
L14:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v33
	if v33 == int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v44 = v33
	goto L1
L16:
	;
	v37 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)) = uint8(v37)
	return int32(0)
L17:
	;
	goto L18
L18:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v33)+8))
	if v41 != v32 {
		v32 = v33
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
}
func F_read(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = l1
	v16 = m.Wasi_snapshot_preview1.Fd_read(m, l0, v7+int32(8), int32(1), v7+int32(4))
	mBase = m.M
	if v16 == int32(0) {
		v23 = int32(0)
	} else {
		*(*int32)(unsafe.Add(mBase, _consts[140])) = v16
		v23 = int32(-1)
	}
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	m.G0 = v7 + int32(16)
	if v23 != 0 {
		v29 = int32(-1)
	} else {
		v29 = v24
	}
	return v29
}
func F_readlink(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	if l2 != 0 {
		v12 = l1
	} else {
		v12 = v7 + int32(15)
	}
	v13 = int32(1)
	if base.Ui32(l2) <= base.Ui32(v13) {
		v16 = v13
	} else {
		v16 = l2
	}
	v17 = m.Env.X__syscall_readlinkat(m, int32(-100), l0, v12, v16)
	mBase = m.M
	if v12 == v7+int32(15) {
		v24 = v17 >> (uint(int32(31)) % 32) & v17
	} else {
		v24 = v17
	}
	if base.Ui32(int32(-4095)) <= base.Ui32(v24) {
		*(*int32)(unsafe.Add(mBase, _consts[140])) = int32(0) - v24
		v32 = int32(-1)
	} else {
		v32 = v24
	}
	m.G0 = v7 + int32(16)
	return v32
}
func F_readtup_cluster(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
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
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l3
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v16 = F_tuplesort_readtup_alloc(m, l0, l3+int32(14))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v16))) = l3 - int32(10)
		*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v16 + int32(24)
		v27 = F_LogicalTapeRead(m, l2, v16+int32(4), int32(6))
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return
		} else {
			if v27 == int32(6) {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = int32(0)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
				v34 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
				v35 = F_LogicalTapeRead(m, l2, v33, v34)
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
					if v35 != v37 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v80 = m.ExcPending
						if v80 != 0 {
							return
						} else {
							F_errmsg_internal(m, int32(517654), int32(0))
							mBase = m.M
							v84 = m.ExcPending
							if v84 != 0 {
								return
							} else {
								F_errfinish(m, int32(504643), int32(1522), int32(219560))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
						if v39&int32(1) != 0 {
							v45 = F_LogicalTapeRead(m, l2, v10+int32(12), int32(4))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return
							} else {
								if v45 != int32(4) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v93 = m.ExcPending
									if v93 != 0 {
										return
									} else {
										F_errmsg_internal(m, int32(517654), int32(0))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return
										} else {
											F_errfinish(m, int32(504643), int32(1524), int32(219560))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
									v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
									if v50 == int32(1) {
										v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
										v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+12)))
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
										v58 = F_heap_getattr_1(m, v16, v54, v55, l1+int32(8))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v58
											m.G0 = v10 + int32(16)
											return
										}
									} else {
										m.G0 = v10 + int32(16)
										return
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v16
							v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
							if v50 == int32(1) {
								v53 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
								v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53)+12)))
								v55 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v58 = F_heap_getattr_1(m, v16, v54, v55, l1+int32(8))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v58
									m.G0 = v10 + int32(16)
									return
								}
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v67 = m.ExcPending
				if v67 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(517654), int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return
					} else {
						F_errfinish(m, int32(504643), int32(1518), int32(219560))
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
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
}
func F_readtup_heap_1(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = l3 + int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v13
	v15 = F_tuplesort_readtup_alloc(m, l0, v13)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v15))) = v13
		v21 = l3 - int32(4)
		v22 = F_LogicalTapeRead(m, l2, v15+int32(10), v21)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			if v22 == v21 {
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+52)))
				if v25&int32(1) != 0 {
					v31 = F_LogicalTapeRead(m, l2, v10+int32(28), int32(4))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						if v31 != int32(4) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v72 = m.ExcPending
							if v72 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(517654), int32(0))
								mBase = m.M
								v76 = m.ExcPending
								if v76 != 0 {
									return
								} else {
									F_errfinish(m, int32(504643), int32(1327), int32(243877))
									mBase = m.M
									v81 = m.ExcPending
									if v81 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
							v37 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v15 - v37
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v36 + v37
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
							v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+10)))
							v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
							v50 = F_heap_getattr_1(m, v10+v37, v46, v47, l1+v37)
							mBase = m.M
							v51 = m.ExcPending
							if v51 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50
								m.G0 = v10 + int32(32)
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1))) = v15
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v37 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v15 - v37
					*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v36 + v37
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
					v46 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45)+10)))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
					v50 = F_heap_getattr_1(m, v10+v37, v46, v47, l1+v37)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v50
						m.G0 = v10 + int32(32)
						return
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(517654), int32(0))
					mBase = m.M
					v63 = m.ExcPending
					if v63 != 0 {
						return
					} else {
						F_errfinish(m, int32(504643), int32(1325), int32(243877))
						mBase = m.M
						v68 = m.ExcPending
						if v68 != 0 {
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
}
func F_reapply_stacked_values(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l2 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
	F_reapply_stacked_values(m, l0, l1, v12, v13, v14, v15, v16)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+112))
	if l3 != v47 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	return
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	switch v20 {
	case 0:
		v35 = int32(2)
		goto L8
	case 1:
		goto L11
	case 2:
		goto L10
	case 3:
		goto L9
	default:
		goto L7
	}
L7:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v43 == v11 {
		goto L1
	} else {
		goto L14
	}
L8:
	;
	v36 = int32(0)
	v40 = F_set_config_with_handle(m, v10, v36, l3, l4, l5, l6, v35, int32(1), int32(19), v36)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L5
	} else {
		goto L13
	}
L9:
	;
	v23 = int32(1)
	v24 = int32(0)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v33 = F_set_config_with_handle(m, v10, v24, v25, v26, int32(13), v28, v24, v23, int32(19), v24)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	v35 = int32(1)
	goto L8
L11:
	;
	v35 = int32(0)
	goto L8
L12:
	;
	v35 = v23
	goto L8
L13:
	;
	goto L7
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v43)+4)) = v45
	return
L15:
	;
	v55 = int32(0)
	v60 = F_set_config_with_handle(m, v10, v55, l3, l4, l5, l6, v55, int32(1), int32(19), v55)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L5
	} else {
		goto L20
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	if l4 != v49 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if l5 != v51 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if l6 == v53 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v62 == int32(0) {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v69 = int32(4534316)
	goto L24
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+56)) = int32(0)
	goto L1
L23:
	;
	goto L22
L24:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v72 == int32(0) {
		goto L23
	} else {
		goto L26
	}
L25:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	*(*int32)(unsafe.Add(mBase, uint32(v69))) = v76
	goto L23
L26:
	;
	if v72 != l0+int32(72) {
		v69 = v72
		goto L24
	} else {
		goto L27
	}
L27:
	;
	goto L25
}
func F_record_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v377 int32
	_ = v377
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v424 int32
	_ = v424
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v451 int32
	_ = v451
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v470 int32
	_ = v470
	var v475 int32
	_ = v475
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	v26 = m.G0
	v28 = v26 - int32(112)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v31 = F_pg_detoast_datum(m, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v36 = F_pg_detoast_datum(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v42 = F_lookup_rowtype_tupdesc(m, v40, v41)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	v47 = F_lookup_rowtype_tupdesc(m, v45, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+108)) = v31
	v52 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+104)) = v52
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+100)) = uint16(v52)
	v56 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+96)) = v56
	v58 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+92)) = int32(base.Ui32(v50) >> (uint(v58) % 32))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+88)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v28)+84)) = v52
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+80)) = uint16(v52)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+76)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = int32(base.Ui32(v61) >> (uint(v58) % 32))
	if v49 < v44 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v73 = v44
	goto L9
L8:
	;
	v73 = v49
	goto L9
L9:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+16))
	if v75 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	if v98 != v40 {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	v86 = F_MemoryContextAlloc(m, v81, v73<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	if v78 < v73 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	v97 = v75
	v98 = v80
	goto L10
L14:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v88)+16)) = v86
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+16))
	v92 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v91)+4)) = v92
	*(*int32)(unsafe.Add(mBase, uint32(v91))) = v73
	*(*int64)(unsafe.Add(mBase, uint32(v91)+12)) = v92
	v97 = v91
	v98 = int32(0)
	goto L10
L15:
	;
	v149 = F_palloc(m, v44<<(uint(int32(2))%32))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L31
	}
L16:
	;
	v107 = v73 << (uint(int32(2)) % 32)
	v109 = v97 + int32(20)
	if v109&int32(3) != 0 {
		goto L22
	} else {
		goto L23
	}
L17:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v97)+8))
	if v100 != v41 {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)+12))
	if v102 != v45 {
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v97)+16))
	if v104 == v46 {
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v97)+16)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v97)+12)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v97)+8)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v97)+4)) = v40
	goto L15
L22:
	;
	v135 = F__emscripten_memset_bulkmem(m, v109, base.I32_extend8_s(int32(0)), v107)
	mBase = m.M
	goto L30
L23:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v107) {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if base.Ui32(v109+v107) <= base.Ui32(v109) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v119 = v97 + v107 + int32(20)
	v121 = v97 + int32(24)
	if base.Ui32(v121) < base.Ui32(v119) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v123 = v119
	goto L28
L27:
	;
	v123 = v121
	goto L28
L28:
	;
	v132 = F__emscripten_memset_bulkmem(m, v109, base.I32_extend8_s(int32(0)), (v123-v97-int32(21))&int32(-4)+int32(4))
	mBase = m.M
	goto L29
L29:
	;
	goto L21
L30:
	;
	goto L21
L31:
	;
	v151 = F_palloc(m, v44)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	F_heap_deform_tuple(m, v28+int32(92), v42, v149, v151)
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v159 = F_palloc(m, v49<<(uint(int32(2))%32))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v161 = F_palloc(m, v49)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_heap_deform_tuple(m, v28+int32(72), v47, v159, v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v165 = int32(0)
	v168 = int32(20)
	v174 = int32(111)
	v179 = base.B2i32(v165 < v44)
	if v165 < v44 {
		goto L43
	} else {
		goto L44
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L1
	} else {
		goto L115
	}
L38:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v460 = m.ExcPending
	if v460 != 0 {
		goto L1
	} else {
		goto L110
	}
L39:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L104
	}
L40:
	;
	F_pfree(m, v149)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L1
	} else {
		goto L84
	}
L41:
	;
	if v350 != v44 {
		goto L37
	} else {
		goto L82
	}
L42:
	;
	v189 = v186
	v191 = base.B2i32(v165 < v49)
	v192 = v165
	v193 = v179
	v195 = v187
	goto L47
L43:
	;
	v180 = int32(0)
	v186 = v180
	v187 = v180
	goto L42
L44:
	;
	goto L45
L45:
	;
	v182 = int32(0)
	if v49 <= v182 {
		v347 = v182
		v350 = v165
		goto L41
	} else {
		goto L46
	}
L46:
	;
	v186 = v182
	v187 = v182
	goto L42
L47:
	;
	if v193&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v347 = v335
	v350 = v337
	goto L41
L49:
	;
	v343 = base.B2i32(v335 < v49)
	v344 = base.B2i32(v337 < v44)
	if v337 < v44 {
		v189 = v335
		v191 = v343
		v192 = v337
		v193 = v344
		v195 = v339
		goto L47
	} else {
		goto L80
	}
L50:
	;
	if v191&int32(1) == int32(0) {
		v347 = v189
		v350 = v192
		goto L41
	} else {
		goto L53
	}
L51:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42+v174+v217<<(uint(int32(4))%32)+v192*int32(100)))))
	if v224 != int32(1) {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v335 = v189
	v337 = v192 + int32(1)
	v339 = v195
	goto L49
L53:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	v235 = v233 << (uint(int32(4)) % 32)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47+v174+v235+v189*int32(100)))))
	if v240 == int32(1) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v335 = v189 + int32(1)
	v337 = v192
	v339 = v195
	goto L49
L55:
	;
	goto L56
L56:
	;
	if v193&int32(1) == int32(0) {
		v347 = v189
		v350 = v192
		goto L41
	} else {
		goto L57
	}
L57:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
	v253 = int32(100)
	v255 = v42 + v168 + v249<<(uint(int32(4))%32) + v192*v253
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+68))
	v260 = v235 + (v47 + v168) + v189*v253
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v260)+68))
	if v256 != v261 {
		goto L39
	} else {
		goto L58
	}
L58:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+96))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v255)+96))
	v267 = v97 + v168 + v195<<(uint(int32(2))%32)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v268 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v280 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192+v151))))
	if v280 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L60:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v268)))
	if v269 == v256 {
		v278 = v268
		goto L59
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	v272 = F_lookup_type_cache(m, v256, int32(64))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v272)+108))
	if v274 == int32(0) {
		goto L38
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v267))) = v272
	v278 = v272
	goto L59
L66:
	;
	v329 = int32(1)
	v335 = v189 + v329
	v337 = v192 + v329
	v339 = v195 + v329
	goto L49
L67:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v161))))
	if v284 != 0 {
		goto L66
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189+v161))))
	if v287 != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v377 = int32(1)
	goto L40
L71:
	;
	v377 = int32(-1)
	goto L40
L72:
	;
	goto L73
L73:
	;
	v289 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+52)) = uint8(v289)
	if v264 == v263 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v293 = v264
	goto L76
L75:
	;
	v293 = v289
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v293
	*(*int64)(unsafe.Add(mBase, uint32(v28)+40)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+36)) = v278 + int32(104)
	v300 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v28)+54)) = uint16(v300)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v149+v192<<(uint(v300)%32))))
	v306 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+60)) = uint8(v306)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+56)) = v305
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v159+v189<<(uint(v300)%32))))
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+68)) = uint8(v306)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+64)) = v312
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v278)+104))
	v320 = m.T0[v319].(func(*base.Module, int32) int32)(m, v28+int32(36))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v320 < int32(0) {
		v377 = int32(-1)
		goto L40
	} else {
		goto L78
	}
L78:
	;
	if v320 == int32(0) {
		goto L66
	} else {
		goto L79
	}
L79:
	;
	v377 = int32(1)
	goto L40
L80:
	;
	if v335 < v49 {
		v189 = v335
		v191 = v343
		v192 = v337
		v193 = v344
		v195 = v339
		goto L47
	} else {
		goto L81
	}
L81:
	;
	goto L48
L82:
	;
	if v347 != v49 {
		goto L37
	} else {
		goto L83
	}
L83:
	;
	v377 = int32(0)
	goto L40
L84:
	;
	F_pfree(m, v151)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	F_pfree(m, v159)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_pfree(m, v161)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	if int32(0) <= v407 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	F_DecrTupleDescRefCount(m, v42)
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v47)+12))
	if int32(0) <= v412 {
		goto L92
	} else {
		goto L93
	}
L91:
	;
	goto L90
L92:
	;
	F_DecrTupleDescRefCount(m, v47)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v417 != v31 {
		goto L96
	} else {
		goto L97
	}
L95:
	;
	goto L94
L96:
	;
	F_pfree(m, v31)
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v421 != v36 {
		goto L100
	} else {
		goto L101
	}
L99:
	;
	goto L98
L100:
	;
	F_pfree(m, v36)
	mBase = m.M
	v424 = m.ExcPending
	if v424 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	m.G0 = v28 + int32(112)
	return v377
L103:
	;
	goto L102
L104:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v255)+68))
	v437 = F_format_type_be(m, v436)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L1
	} else {
		goto L106
	}
L106:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v260)+68))
	v440 = F_format_type_be(m, v439)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28)+24)) = v195 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+20)) = v440
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v437
	F_errmsg(m, int32(483304), v28+int32(16))
	mBase = m.M
	v451 = m.ExcPending
	if v451 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(505559), int32(952), int32(241135))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v465 = F_format_type_be(m, v464)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L1
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v465
	F_errmsg(m, int32(193672), v28)
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(505559), int32(975), int32(241135))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L1
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(151323), int32(0))
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(505559), int32(1040), int32(241135))
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_record_image_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	v2 = int32(0)
	v24 = m.G0
	v26 = v24 + int32(-64)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v29 = F_pg_detoast_datum(m, v28)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v34 = F_pg_detoast_datum(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+8))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v38 = F_lookup_rowtype_tupdesc(m, v36, v37)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v34)+8))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v34)+4))
	v43 = F_lookup_rowtype_tupdesc(m, v41, v42)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v29
	v48 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+56)) = v48
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+52)) = uint16(v48)
	v52 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v52
	v54 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = int32(base.Ui32(v46) >> (uint(v54) % 32))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v48
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+32)) = uint16(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v26)+24)) = int32(base.Ui32(v57) >> (uint(v54) % 32))
	if v45 < v40 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v69 = v40
	goto L8
L7:
	;
	v69 = v45
	goto L8
L8:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)+16))
	if v71 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v94 != v36 {
		goto L15
	} else {
		goto L16
	}
L10:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v70)+20))
	v82 = F_MemoryContextAlloc(m, v77, v69<<(uint(int32(2))%32)+int32(20))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v71)))
	if v74 < v69 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v93 = v71
	v94 = v76
	goto L9
L13:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = v82
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+16))
	v88 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v87)+4)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v69
	*(*int64)(unsafe.Add(mBase, uint32(v87)+12)) = v88
	v93 = v87
	v94 = v2
	goto L9
L14:
	;
	v145 = F_palloc(m, v40<<(uint(int32(2))%32))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L30
	}
L15:
	;
	v103 = v69 << (uint(int32(2)) % 32)
	v105 = v93 + int32(20)
	if v105&int32(3) != 0 {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+8))
	if v96 != v37 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	if v98 != v41 {
		goto L15
	} else {
		goto L18
	}
L18:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v93)+16))
	if v100 == v42 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v93)+16)) = v42
	*(*int32)(unsafe.Add(mBase, uint32(v93)+12)) = v41
	*(*int32)(unsafe.Add(mBase, uint32(v93)+8)) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v36
	goto L14
L21:
	;
	v131 = F__emscripten_memset_bulkmem(m, v105, base.I32_extend8_s(int32(0)), v103)
	mBase = m.M
	goto L29
L22:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v103) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	if base.Ui32(v103+v105) <= base.Ui32(v105) {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v115 = v93 + v103 + int32(20)
	v117 = v93 + int32(24)
	if base.Ui32(v117) < base.Ui32(v115) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v119 = v115
	goto L27
L26:
	;
	v119 = v117
	goto L27
L27:
	;
	v128 = F__emscripten_memset_bulkmem(m, v105, base.I32_extend8_s(int32(0)), (v119-v93-int32(21))&int32(-4)+int32(4))
	mBase = m.M
	goto L28
L28:
	;
	goto L20
L29:
	;
	goto L20
L30:
	;
	v147 = F_palloc(m, v40)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	F_heap_deform_tuple(m, v24+int32(-20), v38, v145, v147)
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v155 = F_palloc(m, v45<<(uint(int32(2))%32))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v157 = F_palloc(m, v45)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	F_heap_deform_tuple(m, v24+int32(-40), v43, v155, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v161 = int32(0)
	v164 = int32(20)
	v168 = int32(111)
	v173 = base.B2i32(v161 < v40)
	if v161 < v40 {
		goto L41
	} else {
		goto L42
	}
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L95
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L89
	}
L38:
	;
	F_pfree(m, v145)
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L69
	}
L39:
	;
	if v296 != v40 {
		goto L36
	} else {
		goto L67
	}
L40:
	;
	v183 = v180
	v184 = v161
	v186 = base.B2i32(v161 < v45)
	v187 = v173
	v191 = v181
	goto L45
L41:
	;
	v174 = int32(0)
	v180 = v174
	v181 = v174
	goto L40
L42:
	;
	goto L43
L43:
	;
	v176 = int32(0)
	if v45 <= v176 {
		v295 = v176
		v296 = v161
		goto L39
	} else {
		goto L44
	}
L44:
	;
	v180 = v176
	v181 = v176
	goto L40
L45:
	;
	if v187&int32(1) == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v295 = v285
	v296 = v286
	goto L39
L47:
	;
	v291 = base.B2i32(v285 < v45)
	v292 = base.B2i32(v286 < v40)
	if v286 < v40 {
		v183 = v285
		v184 = v286
		v186 = v291
		v187 = v292
		v191 = v289
		goto L45
	} else {
		goto L65
	}
L48:
	;
	if v186&int32(1) == int32(0) {
		v295 = v183
		v296 = v184
		goto L39
	} else {
		goto L51
	}
L49:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38+v168+v209<<(uint(int32(4))%32)+v184*int32(100)))))
	if v216 != int32(1) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v285 = v183
	v286 = v184 + int32(1)
	v289 = v191
	goto L47
L51:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	v227 = v225 << (uint(int32(4)) % 32)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43+v168+v227+v183*int32(100)))))
	if v232 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v285 = v183 + int32(1)
	v286 = v184
	v289 = v191
	goto L47
L53:
	;
	goto L54
L54:
	;
	if v187&int32(1) == int32(0) {
		v295 = v183
		v296 = v184
		goto L39
	} else {
		goto L55
	}
L55:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v245 = int32(100)
	v247 = v38 + v164 + v241<<(uint(int32(4))%32) + v184*v245
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+68))
	v252 = v227 + (v43 + v164) + v183*v245
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v252)+68))
	if v248 != v253 {
		goto L37
	} else {
		goto L56
	}
L56:
	;
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v157))))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184+v147))))
	if v258 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v279 = int32(1)
	v285 = v183 + v279
	v286 = v184 + v279
	v289 = v191 + v279
	goto L47
L58:
	;
	if v256&int32(1) != 0 {
		goto L57
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v256&int32(1) != 0 {
		v336 = v2
		goto L38
	} else {
		goto L62
	}
L61:
	;
	v336 = v2
	goto L38
L62:
	;
	v265 = int32(2)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v145+v184<<(uint(v265)%32))))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v155+v183<<(uint(v265)%32))))
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+82)))
	v274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v252)+72)))
	v275 = F_datum_image_eq(m, v268, v272, v273, v274)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if v275 == int32(0) {
		v336 = v2
		goto L38
	} else {
		goto L64
	}
L64:
	;
	goto L57
L65:
	;
	if v285 < v45 {
		v183 = v285
		v184 = v286
		v186 = v291
		v187 = v292
		v191 = v289
		goto L45
	} else {
		goto L66
	}
L66:
	;
	goto L46
L67:
	;
	if v295 != v45 {
		goto L36
	} else {
		goto L68
	}
L68:
	;
	v336 = int32(1)
	goto L38
L69:
	;
	F_pfree(m, v147)
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_pfree(m, v155)
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_pfree(m, v157)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	if int32(0) <= v351 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	F_DecrTupleDescRefCount(m, v38)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	if int32(0) <= v356 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	goto L75
L77:
	;
	F_DecrTupleDescRefCount(m, v43)
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v361 != v29 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	goto L79
L81:
	;
	F_pfree(m, v29)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v365 != v34 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L83
L85:
	;
	F_pfree(m, v34)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	m.G0 = v26 - int32(-64)
	return v336
L88:
	;
	goto L87
L89:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L1
	} else {
		goto L90
	}
L90:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v247)+68))
	v381 = F_format_type_be(m, v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v252)+68))
	v384 = F_format_type_be(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L1
	} else {
		goto L92
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+8)) = v191 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v384
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v381
	F_errmsg(m, int32(483304), v26)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(505559), int32(1720), int32(236178))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errmsg(m, int32(151323), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(505559), int32(1753), int32(236178))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
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
func F_record_image_ge(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_record_image_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_record_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v296 int32
	_ = v296
	var v313 int32
	_ = v313
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v429 int32
	_ = v429
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v24 = F_pg_detoast_datum(m, v23)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v32 = F_lookup_rowtype_tupdesc(m, v30, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+44)) = v24
	v37 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v37
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)) = uint16(v37)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = int32(base.Ui32(v35) >> (uint(int32(2)) % 32))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
	if v48 == v37 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	if v67 == v30 {
		goto L11
	} else {
		goto L12
	}
L6:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v47)+20))
	v59 = F_MemoryContextAlloc(m, v54, v34*int32(44)+int32(12))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+8))
	if v51 != v34 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v48)))
	v67 = v53
	v68 = v48
	goto L5
L9:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v59
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v63)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v64))) = int64(0)
	v67 = v37
	v68 = v64
	goto L5
L10:
	;
	v114 = F_palloc(m, v34<<(uint(int32(2))%32))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L25
	}
L11:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v71 == v31 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v74 = v34 * int32(44)
	v76 = v74 + int32(12)
	if v68&int32(3) != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	goto L13
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v68)+8)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v30
	goto L10
L16:
	;
	v102 = F__emscripten_memset_bulkmem(m, v68, base.I32_extend8_s(int32(0)), v76)
	mBase = m.M
	goto L24
L17:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v76) {
		goto L16
	} else {
		goto L18
	}
L18:
	;
	if base.Ui32(v76+v68) <= base.Ui32(v68) {
		goto L15
	} else {
		goto L19
	}
L19:
	;
	v88 = v68 + v74 + int32(12)
	v90 = v68 + int32(4)
	if base.Ui32(v90) < base.Ui32(v88) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v92 = v88
	goto L22
L21:
	;
	v92 = v90
	goto L22
L22:
	;
	v99 = F__emscripten_memset_bulkmem(m, v68, base.I32_extend8_s(int32(0)), (v68^int32(-1)+v92)&int32(-4)+int32(4))
	mBase = m.M
	goto L23
L23:
	;
	goto L15
L24:
	;
	goto L15
L25:
	;
	v116 = F_palloc(m, v34)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	F_heap_deform_tuple(m, v21+int32(28), v32, v114, v116)
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	F_pq_begintypsend(m, v21+int32(12))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v125 = base.B2i32(v34 <= int32(0))
	if v34 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	F_enlargeStringInfo(m, v21+int32(12), int32(4))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L43
	}
L30:
	;
	v245 = int32(0)
	goto L29
L31:
	;
	goto L32
L32:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v128 = int32(4)
	v130 = v32 + v127<<(uint(v128)%32)
	v132 = v130 + int32(111)
	v133 = int32(0)
	if base.Ui32(v128) <= base.Ui32(v34) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v147 = v133
	v148 = v133
	v156 = int32(0)
	goto L36
L34:
	;
	v192 = v133
	v193 = v133
	goto L35
L35:
	;
	v210 = v34 & int32(3)
	if v210 == int32(0) {
		v245 = v193
		goto L29
	} else {
		goto L39
	}
L36:
	;
	v165 = v147 * int32(100)
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v165))))
	v168 = int32(1)
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+(v130+int32(211))))))
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+(v130+int32(311))))))
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165+(v130+int32(411))))))
	v185 = v148 + (v167 ^ v168) + (v172 ^ v168) + (v177 ^ v168) + (v182 ^ v168)
	v186 = int32(4)
	v187 = v147 + v186
	v189 = v156 + v186
	if v189 != v34&int32(2147483644) {
		v147 = v187
		v148 = v185
		v156 = v189
		goto L36
	} else {
		goto L38
	}
L37:
	;
	v192 = v187
	v193 = v185
	goto L35
L38:
	;
	goto L37
L39:
	;
	v214 = v192
	v215 = v193
	v219 = v133
	goto L40
L40:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v132+v214*int32(100)))))
	v235 = int32(1)
	v237 = v215 + (v234 ^ v235)
	v241 = v219 + v235
	if v241 != v210 {
		v214 = v214 + v235
		v215 = v237
		v219 = v241
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v245 = v237
	goto L29
L42:
	;
	goto L41
L43:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v269 = int32(24)
	v271 = int32(65280)
	v273 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v266+v267))) = v245<<(uint(v269)%32) | v245&v271<<(uint(v273)%32) | (int32(base.Ui32(v245)>>(uint(v273)%32))&v271 | int32(base.Ui32(v245)>>(uint(v269)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v266 + int32(4)
	if v125 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v296 = int32(0)
	goto L47
L45:
	;
	goto L46
L46:
	;
	F_pfree(m, v114)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L1
	} else {
		goto L65
	}
L47:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v32)))
	v319 = v32 + int32(20) + v313<<(uint(int32(4))%32) + v296*int32(100)
	v320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v319)+91)))
	if v320 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	goto L46
L49:
	;
	v440 = v296 + int32(1)
	if v440 != v34 {
		v296 = v440
		goto L47
	} else {
		goto L64
	}
L50:
	;
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v319)+68))
	F_enlargeStringInfo(m, v21+int32(12), int32(4))
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v330 = int32(24)
	v332 = int32(65280)
	v334 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v327+v328))) = v321<<(uint(v330)%32) | v321&v332<<(uint(v334)%32) | (int32(base.Ui32(v321)>>(uint(v334)%32))&v332 | int32(base.Ui32(v321)>>(uint(v330)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v327 + int32(4)
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296+v116))))
	if v350 == int32(1) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_enlargeStringInfo(m, v21+int32(12), int32(4))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L1
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v368 = v68 + int32(12) + v296*int32(44)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v368)))
	if v321 != v369 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v358+v359))) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v358 + int32(4)
	goto L49
L56:
	;
	F_getTypeBinaryOutputInfo(m, v321, v368+int32(4), v368+int32(12))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L1
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v114+v296<<(uint(int32(2))%32))))
	v391 = F_SendFunctionCall(m, v368+int32(16), v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v368)+4))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v380)+20))
	F_fmgr_info_cxt(m, v377, v368+int32(16), v381)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v368))) = v321
	goto L58
L61:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	F_enlargeStringInfo(m, v21+int32(12), int32(4))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v402 = int32(2)
	v404 = int32(4)
	v405 = int32(base.Ui32(v393)>>(uint(v402)%32)) - v404
	v406 = int32(24)
	v408 = int32(65280)
	v410 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v399+v400))) = v405<<(uint(v406)%32) | v405&v408<<(uint(v410)%32) | (int32(base.Ui32(v405)>>(uint(v410)%32))&v408 | int32(base.Ui32(v405)>>(uint(v406)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v399 + v404
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v391)))
	F_pq_sendbytes(m, v21+int32(12), v391+v404, int32(base.Ui32(v429)>>(uint(v402)%32))-v404)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	goto L49
L64:
	;
	goto L48
L65:
	;
	F_pfree(m, v116)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if int32(0) <= v464 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	F_DecrTupleDescRefCount(m, v32)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L1
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v470 = v21 + int32(12)
	v472 = *(*int32)(unsafe.Add(mBase, uint32(v470)))
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v470)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v472))) = v473 << (uint(int32(2)) % 32)
	goto L71
L70:
	;
	goto L69
L71:
	;
	m.G0 = v21 + int32(48)
	return v472
}
func F_record_smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v4 = F_record_cmp(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 < int32(0) {
			v10 = int32(20)
		} else {
			v10 = int32(28)
		}
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0+v10)))
		return v12
	}
}
func F_recurse_push_qual(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = l0
	goto L1
L1:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v19 != int32(142) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L11
	}
L3:
	;
	goto L2
L4:
	;
	if v19 != int32(63) {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	F_recurse_push_qual(m, v39, l1, l2, l3, l4)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L8
	} else {
		goto L10
	}
L7:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v25+v26<<(uint(int32(2))%32)-int32(4))))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+36))
	F_subquery_push_qual(m, v33, l2, l3, l4)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return
L9:
	;
	m.G0 = v10 + int32(16)
	return
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	v12 = v42
	goto L1
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v47
	F_errmsg_internal(m, int32(496237), v10)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(505500), int32(4096), int32(316294))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_recv(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v5 = int32(0)
	v7 = F_recvfrom(m, l0, l1, int32(8192), int32(64), v5, v5)
	return v7
}
func F_reduce_outer_joins_pass2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
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
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v330 int64
	_ = v330
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v410 int32
	_ = v410
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v481 int32
	_ = v481
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v506 int32
	_ = v506
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v579 int32
	_ = v579
	var v587 int32
	_ = v587
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v617 int32
	_ = v617
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	if l0 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v621 = m.ExcPending
	if v621 != 0 {
		goto L10
	} else {
		goto L197
	}
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	switch v19 - int32(63) {
	case 0:
		goto L9
	case 1:
		goto L7
	case 2:
		goto L8
	default:
		goto L1
	}
L3:
	;
	goto L4
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v608 = m.ExcPending
	if v608 != 0 {
		goto L10
	} else {
		goto L194
	}
L5:
	;
	m.G0 = v17 + int32(32)
	return
L6:
	;
	F_bms_free(m, v579)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L10
	} else {
		goto L193
	}
L7:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v99 {
	case 0, 4, 5:
		v496 = v99
		v498 = v98
		v499 = v96
		goto L31
	case 1:
		goto L38
	case 2:
		goto L36
	case 3:
		goto L37
	default:
		goto L35
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v36 = F_find_nonnullable_rels(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L10
	} else {
		goto L14
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	F_errmsg_internal(m, int32(314033), int32(0))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	F_errfinish(m, int32(511399), int32(3267), int32(573771))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L14:
	;
	v38 = F_bms_add_members(m, v36, l4)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L10
	} else {
		goto L15
	}
L15:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v41 = F_find_forced_null_vars(m, v40)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	v43 = F_mbms_add_members(m, v41, l5)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v48 = int32(0)
	goto L18
L18:
	;
	v62 = int32(0)
	if v46 == v62 {
		v72 = v62
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v45 == int32(0) {
		v579 = v38
		goto L6
	} else {
		goto L23
	}
L21:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v46)+4))
	if v66 <= v48 {
		v72 = int32(0)
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v46)+12))
	v72 = v68 + v48<<(uint(int32(2))%32)
	goto L20
L23:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v45)+4))
	if v75 <= v48 {
		v579 = v38
		goto L6
	} else {
		goto L24
	}
L24:
	;
	if v72 == int32(0) {
		v579 = v38
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v45)+12))
	v82 = v79 + v48<<(uint(int32(2))%32)
	if v82 == int32(0) {
		v579 = v38
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v82)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+4)))
	if v86 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v72)))
	F_reduce_outer_joins_pass2(m, v89, v85, l2, l3, v38, v43)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L10
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v48 = v48 + int32(1)
	goto L18
L30:
	;
	goto L29
L31:
	;
	if v97 == int32(0) {
		goto L159
	} else {
		goto L160
	}
L32:
	;
	v343 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v345 = F_find_nonnullable_vars_walker(m, v343, int32(1))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L10
	} else {
		goto L111
	}
L33:
	;
	v330 = *(*int64)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+12)) = base.I64_rotl(v330, int64(32))
	v334 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v334)+12))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v335)+4))
	v337 = *(*int32)(unsafe.Add(mBase, uint32(v335)))
	v340 = v336
	v341 = v337
	goto L32
L34:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v319 = F_palloc(m, int32(8))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L10
	} else {
		goto L109
	}
L35:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L10
	} else {
		goto L106
	}
L36:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v199 = int32(0)
	if l4 == v199 {
		v240 = v199
		goto L70
	} else {
		goto L71
	}
L37:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v150 = int32(0)
	if l4 == v150 {
		v191 = v150
		goto L55
	} else {
		goto L56
	}
L38:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v101 = int32(0)
	if l4 == v101 {
		v142 = v101
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v142 == int32(0) {
		v340 = v98
		v341 = v96
		goto L32
	} else {
		goto L53
	}
L40:
	;
	goto L39
L41:
	;
	if v100 == int32(0) {
		v142 = v101
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	if v110 < v111 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v113 = v110
	goto L45
L44:
	;
	v113 = v111
	goto L45
L45:
	;
	if v113 <= int32(1) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v116 = int32(1)
	goto L48
L47:
	;
	v116 = v113
	goto L48
L48:
	;
	v117 = int32(8)
	v122 = int32(0)
	goto L49
L49:
	;
	v129 = v122 << (uint(int32(2)) % 32)
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v100+v117+v129)))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v129+(l4+v117))))
	v134 = v131 & v133
	v136 = base.B2i32(v134 != int32(0))
	if v134 != 0 {
		v142 = v136
		goto L40
	} else {
		goto L51
	}
L50:
	;
	v142 = v136
	goto L40
L51:
	;
	v138 = v122 + int32(1)
	if v138 != v116 {
		v122 = v138
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v496 = int32(0)
	v498 = v98
	v499 = v96
	goto L31
L54:
	;
	if v191 == int32(0) {
		goto L33
	} else {
		goto L68
	}
L55:
	;
	goto L54
L56:
	;
	if v149 == int32(0) {
		v191 = v150
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	if v159 < v160 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v162 = v159
	goto L60
L59:
	;
	v162 = v160
	goto L60
L60:
	;
	if v162 <= int32(1) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v165 = int32(1)
	goto L63
L62:
	;
	v165 = v162
	goto L63
L63:
	;
	v166 = int32(8)
	v171 = int32(0)
	goto L64
L64:
	;
	v178 = v171 << (uint(int32(2)) % 32)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v149+v166+v178)))
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+(l4+v166))))
	v183 = v180 & v182
	v185 = base.B2i32(v183 != int32(0))
	if v183 != 0 {
		v191 = v185
		goto L55
	} else {
		goto L66
	}
L65:
	;
	v191 = v185
	goto L55
L66:
	;
	v187 = v171 + int32(1)
	if v187 != v165 {
		v171 = v187
		goto L64
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	v496 = int32(0)
	v498 = v98
	v499 = v96
	goto L31
L69:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v245 = int32(0)
	if l4 == v245 {
		v286 = v245
		goto L84
	} else {
		goto L85
	}
L70:
	;
	goto L69
L71:
	;
	if v198 == int32(0) {
		v240 = v199
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v198)+4))
	if v208 < v209 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v211 = v208
	goto L75
L74:
	;
	v211 = v209
	goto L75
L75:
	;
	if v211 <= int32(1) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v214 = int32(1)
	goto L78
L77:
	;
	v214 = v211
	goto L78
L78:
	;
	v215 = int32(8)
	v220 = int32(0)
	goto L79
L79:
	;
	v227 = v220 << (uint(int32(2)) % 32)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v198+v215+v227)))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+(l4+v215))))
	v232 = v229 & v231
	v234 = base.B2i32(v232 != int32(0))
	if v232 != 0 {
		v240 = v234
		goto L70
	} else {
		goto L81
	}
L80:
	;
	v240 = v234
	goto L70
L81:
	;
	v236 = v220 + int32(1)
	if v236 != v214 {
		v220 = v236
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	if v240 != 0 {
		goto L97
	} else {
		goto L98
	}
L84:
	;
	goto L83
L85:
	;
	if v244 == int32(0) {
		v286 = v245
		goto L84
	} else {
		goto L86
	}
L86:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v244)+4))
	if v254 < v255 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v257 = v254
	goto L89
L88:
	;
	v257 = v255
	goto L89
L89:
	;
	if v257 <= int32(1) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v260 = int32(1)
	goto L92
L91:
	;
	v260 = v257
	goto L92
L92:
	;
	v261 = int32(8)
	v266 = int32(0)
	goto L93
L93:
	;
	v273 = v266 << (uint(int32(2)) % 32)
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v244+v261+v273)))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273+(l4+v261))))
	v278 = v275 & v277
	v280 = base.B2i32(v278 != int32(0))
	if v278 != 0 {
		v286 = v280
		goto L84
	} else {
		goto L95
	}
L94:
	;
	v286 = v280
	goto L84
L95:
	;
	v282 = v266 + int32(1)
	if v282 != v260 {
		v266 = v282
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	if v286 != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	if v286 != 0 {
		goto L34
	} else {
		goto L105
	}
L100:
	;
	v496 = int32(0)
	v498 = v98
	v499 = v96
	goto L31
L101:
	;
	goto L102
L102:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v96)))
	v293 = F_palloc(m, int32(8))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L10
	} else {
		goto L103
	}
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v293)+4)) = v291
	*(*int32)(unsafe.Add(mBase, uint32(v293))) = v97
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v298 = F_lappend(m, v297, v293)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L10
	} else {
		goto L104
	}
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v298
	v340 = v98
	v341 = v96
	goto L32
L105:
	;
	v496 = int32(2)
	v498 = v98
	v499 = v96
	goto L31
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v99
	F_errmsg_internal(m, int32(495432), v17+int32(16))
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L10
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(511399), int32(3355), int32(573771))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L10
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v319)+4)) = v317
	*(*int32)(unsafe.Add(mBase, uint32(v319))) = v97
	v323 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v324 = F_lappend(m, v323, v319)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L10
	} else {
		goto L110
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v324
	goto L33
L111:
	;
	v347 = int32(0)
	v350 = v347
	v355 = v347
	goto L112
L112:
	;
	v363 = int32(0)
	if v345 == v363 {
		v373 = v363
		goto L114
	} else {
		goto L115
	}
L113:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v444 = int32(0)
	if v439 == v444 {
		v485 = v444
		goto L143
	} else {
		goto L144
	}
L114:
	;
	if l5 == int32(0) {
		goto L118
	} else {
		goto L119
	}
L115:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v345)+4))
	if v367 <= v350 {
		v373 = int32(0)
		goto L114
	} else {
		goto L116
	}
L116:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v345)+12))
	v373 = v369 + v350<<(uint(int32(2))%32)
	goto L114
L117:
	;
	goto L113
L118:
	;
	v439 = int32(0)
	goto L117
L119:
	;
	goto L120
L120:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v377 <= v350 {
		v439 = v355
		goto L117
	} else {
		goto L121
	}
L121:
	;
	if v373 == int32(0) {
		v439 = v355
		goto L117
	} else {
		goto L122
	}
L122:
	;
	v381 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	v384 = v381 + v350<<(uint(int32(2))%32)
	if v384 == int32(0) {
		v439 = v355
		goto L117
	} else {
		goto L123
	}
L123:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384)))
	v389 = int32(0)
	if v387 == v389 {
		v430 = v389
		goto L125
	} else {
		goto L126
	}
L124:
	;
	if v430 != 0 {
		goto L138
	} else {
		goto L139
	}
L125:
	;
	goto L124
L126:
	;
	if v388 == int32(0) {
		v430 = v389
		goto L125
	} else {
		goto L127
	}
L127:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v387)+4))
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v388)+4))
	if v398 < v399 {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	v401 = v398
	goto L130
L129:
	;
	v401 = v399
	goto L130
L130:
	;
	if v401 <= int32(1) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v404 = int32(1)
	goto L133
L132:
	;
	v404 = v401
	goto L133
L133:
	;
	v405 = int32(8)
	v410 = int32(0)
	goto L134
L134:
	;
	v417 = v410 << (uint(int32(2)) % 32)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v388+v405+v417)))
	v421 = *(*int32)(unsafe.Add(mBase, uint32(v417+(v387+v405))))
	v422 = v419 & v421
	v424 = base.B2i32(v422 != int32(0))
	if v422 != 0 {
		v430 = v424
		goto L125
	} else {
		goto L136
	}
L135:
	;
	v430 = v424
	goto L125
L136:
	;
	v426 = v410 + int32(1)
	if v426 != v404 {
		v410 = v426
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	v434 = F_bms_add_member(m, v355, v350)
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L10
	} else {
		goto L141
	}
L139:
	;
	v436 = v355
	goto L140
L140:
	;
	v350 = v350 + int32(1)
	v355 = v436
	goto L112
L141:
	;
	v436 = v434
	goto L140
L142:
	;
	if v485 != 0 {
		goto L156
	} else {
		goto L157
	}
L143:
	;
	goto L142
L144:
	;
	if v443 == int32(0) {
		v485 = v444
		goto L143
	} else {
		goto L145
	}
L145:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v439)+4))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v443)+4))
	if v453 < v454 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v456 = v453
	goto L148
L147:
	;
	v456 = v454
	goto L148
L148:
	;
	if v456 <= int32(1) {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v459 = int32(1)
	goto L151
L150:
	;
	v459 = v456
	goto L151
L151:
	;
	v460 = int32(8)
	v465 = int32(0)
	goto L152
L152:
	;
	v472 = v465 << (uint(int32(2)) % 32)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v443+v460+v472)))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v472+(v439+v460))))
	v477 = v474 & v476
	v479 = base.B2i32(v477 != int32(0))
	if v477 != 0 {
		v485 = v479
		goto L143
	} else {
		goto L154
	}
L153:
	;
	v485 = v479
	goto L143
L154:
	;
	v481 = v465 + int32(1)
	if v481 != v459 {
		v465 = v481
		goto L152
	} else {
		goto L155
	}
L155:
	;
	goto L153
L156:
	;
	v489 = int32(5)
	goto L158
L157:
	;
	v489 = int32(1)
	goto L158
L158:
	;
	v496 = v489
	v498 = v340
	v499 = v341
	goto L31
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v496
	v523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+4)))
	if v523 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L160:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v496 == v506 {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v508)+52))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(v509)+12))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(v510+v97<<(uint(int32(2))%32)-int32(4))))
	*(*int32)(unsafe.Add(mBase, uint32(v516)+44)) = v496
	if v496 != 0 {
		goto L159
	} else {
		goto L162
	}
L162:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v519 = F_bms_add_member(m, v518, v97)
	mBase = m.M
	v520 = m.ExcPending
	if v520 != 0 {
		goto L10
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v519
	goto L159
L164:
	;
	v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+4)))
	if v526 != int32(1) {
		goto L5
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v529 = int32(0)
	if v496 == int32(2) {
		v546 = v529
		v547 = v529
		goto L168
	} else {
		goto L169
	}
L167:
	;
	goto L166
L168:
	;
	v548 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v498)+4)))
	if v548 == int32(1) {
		goto L175
	} else {
		goto L176
	}
L169:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v534 = F_find_nonnullable_rels(m, v533)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L10
	} else {
		goto L170
	}
L170:
	;
	v536 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v537 = F_find_forced_null_vars(m, v536)
	mBase = m.M
	v538 = m.ExcPending
	if v538 != 0 {
		goto L10
	} else {
		goto L171
	}
L171:
	;
	if v496&int32(-5) != 0 {
		v546 = v534
		v547 = v537
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v541 = F_bms_add_members(m, v534, l4)
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L10
	} else {
		goto L173
	}
L173:
	;
	v543 = F_mbms_add_members(m, v537, l5)
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L10
	} else {
		goto L174
	}
L174:
	;
	v546 = v541
	v547 = v543
	goto L168
L175:
	;
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v554 = base.B2i32(v496 == int32(2))
	if v496 == int32(2) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	goto L177
L177:
	;
	v566 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v499)+4)))
	if v566 != int32(1) {
		v579 = v546
		goto L6
	} else {
		goto L191
	}
L178:
	;
	v555 = int32(0)
	goto L180
L179:
	;
	v555 = l4
	goto L180
L180:
	;
	v557 = v496 & int32(-5)
	if v557 != 0 {
		goto L181
	} else {
		goto L182
	}
L181:
	;
	v558 = v555
	goto L183
L182:
	;
	v558 = v546
	goto L183
L183:
	;
	if v496 == int32(2) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v560 = int32(0)
	goto L186
L185:
	;
	v560 = l5
	goto L186
L186:
	;
	if v557 != 0 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v561 = v560
	goto L189
L188:
	;
	v561 = v547
	goto L189
L189:
	;
	F_reduce_outer_joins_pass2(m, v551, v498, l2, l3, v558, v561)
	mBase = m.M
	v563 = m.ExcPending
	if v563 != 0 {
		goto L10
	} else {
		goto L190
	}
L190:
	;
	goto L177
L191:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	F_reduce_outer_joins_pass2(m, v569, v499, l2, l3, v546, v547)
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L10
	} else {
		goto L192
	}
L192:
	;
	v579 = v546
	goto L6
L193:
	;
	goto L5
L194:
	;
	F_errmsg_internal(m, int32(418866), int32(0))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L10
	} else {
		goto L195
	}
L195:
	;
	F_errfinish(m, int32(511399), int32(3265), int32(573771))
	mBase = m.M
	v617 = m.ExcPending
	if v617 != 0 {
		goto L10
	} else {
		goto L196
	}
L196:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L197:
	;
	v622 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v622
	F_errmsg_internal(m, int32(496237), v17)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L10
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(511399), int32(3522), int32(573771))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L10
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regcollationin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L31
	} else {
		goto L47
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v155
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v155 = int32(0)
	goto L2
L8:
	;
	v23 = int32(564831)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[956])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v11)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L30
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[957])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v11
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v11
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 == int32(0) {
		v98 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v98 - v11
	goto L9
L23:
	;
	v77 = v11
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v98 = v94
	goto L22
L26:
	;
	v98 = v77
	goto L22
L27:
	;
	goto L28
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v155 = v117
	goto L2
L33:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v122 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v155 = int32(0)
	goto L2
L36:
	;
	goto L37
L37:
	;
	v130 = F_get_collation_oid(m, v122, int32(1))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L31
	} else {
		goto L38
	}
L38:
	;
	if v130 != 0 {
		v155 = v130
		goto L2
	} else {
		goto L39
	}
L39:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L31
	} else {
		goto L40
	}
L40:
	;
	if v133 == int32(0) {
		v155 = v132
		goto L2
	} else {
		goto L41
	}
L41:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L31
	} else {
		goto L42
	}
L42:
	;
	v140 = F_NameListToString(m, v122)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	v143 = *(*int32)(unsafe.Add(mBase, _consts[495]))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	goto L44
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v144
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v140
	F_errmsg(m, int32(73497), v8)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	F_errsave_finish(m, v10, int32(512134), int32(1057), int32(281714))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L31
	} else {
		goto L46
	}
L46:
	;
	v155 = v132
	goto L2
L47:
	;
	F_errmsg_internal(m, int32(420867), int32(0))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(512134), int32(1041), int32(281714))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L31
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regconfigout(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v10 == int32(0) {
		v14 = F_pstrdup(m, int32(683984))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v45 = v14
			m.G0 = v8 + int32(16)
			return v45
		}
	} else {
		v19 = F_SearchSysCache1(m, int32(74), v10)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			if v19 != 0 {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(v19)+16))
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+22)))
				v23 = v21 + v22
				v26 = F_TSConfigIsVisible(m, v10)
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					if v26 != 0 {
						v32 = int32(0)
						v33 = F_quote_qualified_identifier(m, v32, v23+int32(4))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_ReleaseCatCache(m, v19)
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v45 = v33
								m.G0 = v8 + int32(16)
								return v45
							}
						}
					} else {
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v23)+68))
						v30 = F_get_namespace_name(m, v29)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							v32 = v30
							v33 = F_quote_qualified_identifier(m, v32, v23+int32(4))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_ReleaseCatCache(m, v19)
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
									return int32(0)
								} else {
									v45 = v33
									m.G0 = v8 + int32(16)
									return v45
								}
							}
						}
					}
				}
			} else {
				v38 = F_palloc(m, int32(64))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					v43 = F_pg_snprintf(m, v38, int32(64), int32(60472), v8)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						v45 = v38
						m.G0 = v8 + int32(16)
						return v45
					}
				}
			}
		}
	}
}
func F_regnamespacein(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
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
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
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
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v12 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L31
	} else {
		goto L53
	}
L2:
	;
	m.G0 = v8 + int32(16)
	return v175
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L4:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v15 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v12-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v175 = int32(0)
	goto L2
L8:
	;
	v23 = int32(564831)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[956])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v11)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L30
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[957])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v11
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v11
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v73 == int32(0) {
		v98 = v11
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v98 - v11
	goto L9
L23:
	;
	v77 = v11
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v98 = v94
	goto L22
L26:
	;
	v98 = v77
	goto L22
L27:
	;
	goto L28
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v11, int32(-1), v10, v8+int32(12))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v175 = v117
	goto L2
L33:
	;
	v122 = F_stringToQualifiedNameList(m, v11, v10)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v122 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v126 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v126)
	v175 = int32(0)
	goto L2
L36:
	;
	goto L37
L37:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	if v129 != int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = int32(0)
	v133 = F_errsave_start(m, v10)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L31
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)+4))
	v153 = F_get_namespace_oid(m, v151, int32(1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L31
	} else {
		goto L46
	}
L41:
	;
	if v133 == int32(0) {
		v175 = v132
		goto L2
	} else {
		goto L42
	}
L42:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L31
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(30241), int32(0))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L31
	} else {
		goto L44
	}
L44:
	;
	F_errsave_finish(m, v10, int32(512134), int32(1681), int32(282641))
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L31
	} else {
		goto L45
	}
L45:
	;
	v175 = v132
	goto L2
L46:
	;
	if v153 != 0 {
		v175 = v153
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v155 = int32(0)
	v156 = F_errsave_start(m, v10)
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L31
	} else {
		goto L48
	}
L48:
	;
	if v156 == int32(0) {
		v175 = v155
		goto L2
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L31
	} else {
		goto L50
	}
L50:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v165
	F_errmsg(m, int32(74155), v8)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L31
	} else {
		goto L51
	}
L51:
	;
	F_errsave_finish(m, v10, int32(512134), int32(1689), int32(282641))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L31
	} else {
		goto L52
	}
L52:
	;
	v175 = v155
	goto L2
L53:
	;
	F_errmsg_internal(m, int32(421109), int32(0))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L31
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(512134), int32(1671), int32(282641))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L31
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regprocedurein(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v2 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(432)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v13 == int32(45) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L31
	} else {
		goto L70
	}
L2:
	;
	m.G0 = v9 + int32(432)
	return v248
L3:
	;
	v119 = *(*int32)(unsafe.Add(mBase, _consts[123]))
	if v119 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L4:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+1)))
	if v16 != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if base.Ui32(int32(9)) < base.Ui32((v13-int32(48))&int32(255)) {
		goto L3
	} else {
		goto L8
	}
L7:
	;
	v248 = v2
	goto L2
L8:
	;
	v23 = int32(564831)
	v27 = m.G0
	v29 = v27 - int32(32)
	v30 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v29)+24)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+16)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29)+8)) = v30
	*(*int64)(unsafe.Add(mBase, uint32(v29))) = v30
	v38 = int32(*(*uint8)(unsafe.Add(mBase, _consts[956])))
	if v38 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v107 = F_strlen(m, v12)
	mBase = m.M
	if v106 != v107 {
		goto L3
	} else {
		goto L30
	}
L10:
	;
	v106 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, _consts[957])))
	if v42 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v46 = v12
	goto L16
L14:
	;
	goto L15
L15:
	;
	v56 = v23
	v57 = v38
	goto L19
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
	if v52 == v38 {
		v46 = v46 + int32(1)
		goto L16
	} else {
		goto L18
	}
L17:
	;
	v106 = v46 - v12
	goto L9
L18:
	;
	goto L17
L19:
	;
	v64 = v29 + int32(base.Ui32(v57)>>(uint(int32(3))%32))&int32(28)
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)))
	v66 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v65 | v66<<(uint(v57)%32)
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	if v70 != 0 {
		v56 = v56 + v66
		v57 = v70
		goto L19
	} else {
		goto L21
	}
L20:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if v73 == int32(0) {
		v98 = v12
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v106 = v98 - v12
	goto L9
L23:
	;
	v77 = v12
	v78 = v73
	goto L24
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v29+int32(base.Ui32(v78)>>(uint(int32(3))%32))&int32(28))))
	if int32(base.Ui32(v86)>>(uint(v78)%32))&int32(1) == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v98 = v94
	goto L22
L26:
	;
	v98 = v77
	goto L22
L27:
	;
	goto L28
L28:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+1)))
	v94 = v77 + int32(1)
	if v92 != 0 {
		v77 = v94
		v78 = v92
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v113 = F_DirectInputFunctionCallSafe(m, int32(547), v12, int32(-1), v11, v9+int32(16))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	return int32(0)
L32:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
	v248 = v117
	goto L2
L33:
	;
	v129 = F_parseNameAndArgTypes(m, v12, int32(0), v9+int32(428), v9+int32(424), v9+int32(16), v11)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	if v129 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v133 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v133)
	v248 = v2
	goto L2
L36:
	;
	goto L37
L37:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v9)+428))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v9)+424))
	v137 = int32(0)
	v142 = F_FuncnameGetCandidates(m, v135, v136, v137, v137, v137, v137, int32(1))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L31
	} else {
		goto L39
	}
L38:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v146)+8))
	v248 = v243
	goto L2
L39:
	;
	if v142 != 0 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v145 = v136 << (uint(int32(2)) % 32)
	v146 = v142
	goto L43
L41:
	;
	goto L42
L42:
	;
	v227 = F_errsave_start(m, v11)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L31
	} else {
		goto L65
	}
L43:
	;
	v153 = v146 + int32(32)
	v155 = v9 + int32(16)
	if base.Ui32(int32(4)) <= base.Ui32(v145) {
		goto L48
	} else {
		goto L49
	}
L44:
	;
	goto L42
L45:
	;
	if v217 == int32(0) {
		goto L38
	} else {
		goto L63
	}
L46:
	;
	v217 = int32(0)
	goto L45
L47:
	;
	v191 = v186
	v192 = v187
	v193 = v188
	goto L57
L48:
	;
	if (v153|v155)&int32(3) != 0 {
		v186 = v153
		v187 = v155
		v188 = v145
		goto L47
	} else {
		goto L51
	}
L49:
	;
	v179 = v153
	v180 = v155
	v181 = v145
	goto L50
L50:
	;
	if v181 == int32(0) {
		goto L46
	} else {
		goto L56
	}
L51:
	;
	v163 = v153
	v164 = v155
	v165 = v145
	goto L52
L52:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v168 != v169 {
		v186 = v163
		v187 = v164
		v188 = v165
		goto L47
	} else {
		goto L54
	}
L53:
	;
	v179 = v174
	v180 = v172
	v181 = v176
	goto L50
L54:
	;
	v171 = int32(4)
	v172 = v164 + v171
	v174 = v163 + v171
	v176 = v165 - v171
	if base.Ui32(int32(3)) < base.Ui32(v176) {
		v163 = v174
		v164 = v172
		v165 = v176
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	v186 = v179
	v187 = v180
	v188 = v181
	goto L47
L57:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v196 == v197 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v217 = v196 - v197
	goto L45
L59:
	;
	v199 = int32(1)
	v204 = v193 - v199
	if v204 != 0 {
		v191 = v191 + v199
		v192 = v192 + v199
		v193 = v204
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	goto L46
L63:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v146)))
	if v220 != 0 {
		v146 = v220
		goto L43
	} else {
		goto L64
	}
L64:
	;
	goto L44
L65:
	;
	if v227 == int32(0) {
		v248 = v2
		goto L2
	} else {
		goto L66
	}
L66:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L31
	} else {
		goto L67
	}
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v12
	F_errmsg(m, int32(72732), v9)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L31
	} else {
		goto L68
	}
L68:
	;
	F_errsave_finish(m, v11, int32(512134), int32(265), int32(282599))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L31
	} else {
		goto L69
	}
L69:
	;
	v248 = v2
	goto L2
L70:
	;
	F_errmsg_internal(m, int32(420966), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L31
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(512134), int32(240), int32(282599))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L31
	} else {
		goto L72
	}
L72:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_regprocsend(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_date_send(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_removeabbrev_datum(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	v4 = int32(0)
	if l2 <= v4 {
	} else {
		v11 = l2 & int32(3)
		v12 = int32(0)
		if base.Ui32(int32(4)) <= base.Ui32(l2) {
			v17 = v12
			v22 = v4
			for {
				v24 = int32(4)
				v26 = l1 + v17<<(uint(v24)%32)
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+4)) = v27
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v29
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+52)) = v33
				v36 = v17 + v24
				v38 = v22 + v24
				if v38 != l2&int32(2147483644) {
					v17 = v36
					v22 = v38
					continue
				} else {
					break
				}
				break
			}
			v40 = v36
		} else {
			v40 = v12
		}
		if v11 == int32(0) {
		} else {
			v49 = v40
			v53 = v4
			for {
				v58 = l1 + v49<<(uint(int32(4))%32)
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
				*(*int32)(unsafe.Add(mBase, uint32(v58)+4)) = v59
				v61 = int32(1)
				v64 = v53 + v61
				if v64 != v11 {
					v49 = v49 + v61
					v53 = v64
					continue
				} else {
					break
				}
				break
			}
		}
	}
	return
}
func F_removetraverse(m *base.Module, l0 int32, l1 int32) {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v94 int32
	_ = v94
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v167 int64
	_ = v167
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = m.T0[v10].(func(*base.Module) int32)(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v11 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = int32(101)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	if v17 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v21 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v19 = v17
	goto L8
L7:
	;
	v19 = int32(19)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v19
	return
L9:
	;
	return
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = l1
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v23 == int32(0) {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v31 = v23
	goto L12
L12:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v33 != 0 {
		goto L9
	} else {
		goto L14
	}
L13:
	;
	goto L9
L14:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	F_removetraverse(m, l0, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	if v38 != 0 {
		goto L9
	} else {
		goto L16
	}
L16:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	switch v40 - int32(76) {
	case 0, 18, 21, 38:
		goto L19
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 19, 20, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 35, 37, 39, 40, 41, 42, 43:
		goto L18
	case 34, 36, 44:
		goto L17
	default:
		goto L20
	}
L17:
	;
	if v39 != 0 {
		v31 = v39
		goto L12
	} else {
		goto L78
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+24)) = int32(101)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)+12))
	if v179 != 0 {
		goto L75
	} else {
		goto L76
	}
L19:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v47 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	if v47 != 0 {
		goto L22
	} else {
		goto L23
	}
L20:
	;
	if v40 != int32(36) {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v45)+8))
	if v50 <= v51 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L24
L26:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
	v108 = int32(*(*int16)(unsafe.Add(mBase, uint32(v31)+4)))
	if v108 < int32(0) {
		goto L49
	} else {
		goto L50
	}
L27:
	;
	F_createarc(m, l0, int32(110), int32(0), l1, v45)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L47
	}
L28:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v53 == int32(0) {
		goto L27
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v45)+16))
	if v69 == int32(0) {
		goto L27
	} else {
		goto L39
	}
L31:
	;
	v58 = v53
	goto L32
L32:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58)+12))
	if v62 != v45 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L27
L34:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+16))
	if v68 != 0 {
		v58 = v68
		goto L32
	} else {
		goto L38
	}
L35:
	;
	v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+4)))
	if v64 != 0 {
		goto L34
	} else {
		goto L36
	}
L36:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
	if v65 == int32(110) {
		goto L26
	} else {
		goto L37
	}
L37:
	;
	goto L34
L38:
	;
	goto L33
L39:
	;
	v74 = v69
	goto L40
L40:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
	if v78 != l1 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L27
L42:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v74)+24))
	if v84 != 0 {
		v74 = v84
		goto L40
	} else {
		goto L46
	}
L43:
	;
	v80 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v74)+4)))
	if v80 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v81 == int32(110) {
		goto L26
	} else {
		goto L45
	}
L45:
	;
	goto L42
L46:
	;
	goto L41
L47:
	;
	goto L26
L48:
	;
	goto L17
L49:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v31)+20))
	if v142 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L50:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
	v113 = v111 - int32(97)
	if base.Ui32(int32(17)) < base.Ui32(v113) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	if int32(1)<<(uint(v113)%32)&int32(163841) == int32(0) {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	if v122 != 0 {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v31)+36))
	if v123 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	if v135 != 0 {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v126)+20))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v127+v108*int32(24))+12)) = v131
	v135 = v131
	goto L54
L56:
	;
	goto L57
L57:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v31)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+32)) = v133
	v135 = v133
	goto L54
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135)+36)) = v123
	goto L60
L59:
	;
	goto L60
L60:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+32)) = int64(0)
	goto L49
L61:
	;
	if v141 != 0 {
		goto L65
	} else {
		goto L66
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v107)+20)) = v141
	goto L61
L63:
	;
	goto L64
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v142)+16)) = v141
	goto L61
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v141)+20)) = v142
	goto L67
L66:
	;
	goto L67
L67:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v107)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v107)+12)) = v148 - int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v31)+24))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v31)+28))
	if v153 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v159 = v31 + int32(8)
	if v152 != 0 {
		goto L72
	} else {
		goto L73
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+16)) = v152
	goto L68
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153)+24)) = v152
	goto L68
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v152)+28)) = v153
	goto L74
L73:
	;
	goto L74
L74:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v106)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+8)) = v161 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v31))) = int32(0)
	v167 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v159)+16)) = v167
	*(*int64)(unsafe.Add(mBase, uint32(v159)+8)) = v167
	*(*int64)(unsafe.Add(mBase, uint32(v159))) = v167
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+16)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v31
	goto L48
L75:
	;
	v181 = v179
	goto L77
L76:
	;
	v181 = int32(15)
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v178)+12)) = v181
	goto L17
L78:
	;
	goto L13
}
func F_renameatt_check(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+119)))
	if l2 == int32(0) {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+76))
		if v13 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
					return
				} else {
					F_errmsg(m, int32(402939), int32(0))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return
					} else {
						F_errfinish(m, int32(505917), int32(3802), int32(324794))
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			switch v10 - int32(73) {
			case 0, 26, 29, 32, 36, 39, 41, 45:
				v39 = *(*int32)(unsafe.Add(mBase, _consts[168]))
				v40 = F_object_ownercheck(m, int32(1259), l0, v39)
				mBase = m.M
				v41 = m.ExcPending
				if v41 != 0 {
					return
				} else {
					if v40 == int32(0) {
						v45 = F_get_rel_relkind(m, l0)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return
						} else {
							switch v45 - int32(73) {
							case 0, 32:
								v58 = int32(20)
							default:
								v56 = int32(41)
								v58 = v56
							case 10:
								v58 = int32(37)
							case 29:
								v56 = int32(18)
								v58 = v56
							case 36:
								v58 = int32(23)
							case 45:
								v58 = int32(51)
							}
							F_aclcheck_error(m, int32(2), v58, l1+int32(4))
							mBase = m.M
							v62 = m.ExcPending
							if v62 != 0 {
								return
							} else {
								v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[459])))
								if v64 == int32(0) {
									v68 = int32(1)
									if base.Ui32(l0) < base.Ui32(int32(12000)) {
										v76 = v68
									} else {
										v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
										if v71 == int32(99) {
											v76 = v68
										} else {
											v74 = F_isTempToastNamespace(m, v71)
											mBase = m.M
											v76 = v74
										}
									}
									if v76 != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v99 = m.ExcPending
										if v99 != 0 {
											return
										} else {
											F_errcode(m, int32(16797828))
											mBase = m.M
											v102 = m.ExcPending
											if v102 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1 + int32(4)
												F_errmsg(m, int32(334761), v8+int32(16))
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return
												} else {
													F_errfinish(m, int32(505917), int32(3835), int32(324794))
													mBase = m.M
													v115 = m.ExcPending
													if v115 != 0 {
														return
													} else {
														base.Wasm_trap_unreachable()
														for {
														}
													}
												}
											}
										}
									} else {
										m.G0 = v8 + int32(32)
										return
									}
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							}
						}
					} else {
						v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[459])))
						if v64 == int32(0) {
							v68 = int32(1)
							if base.Ui32(l0) < base.Ui32(int32(12000)) {
								v76 = v68
							} else {
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
								if v71 == int32(99) {
									v76 = v68
								} else {
									v74 = F_isTempToastNamespace(m, v71)
									mBase = m.M
									v76 = v74
								}
							}
							if v76 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v99 = m.ExcPending
								if v99 != 0 {
									return
								} else {
									F_errcode(m, int32(16797828))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1 + int32(4)
										F_errmsg(m, int32(334761), v8+int32(16))
										mBase = m.M
										v110 = m.ExcPending
										if v110 != 0 {
											return
										} else {
											F_errfinish(m, int32(505917), int32(3835), int32(324794))
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								}
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					}
				}
			default:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errcode(m, int32(151027844))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
						F_errmsg(m, int32(721795), v8)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return
						} else {
							F_errdetail_relkind_not_supported(m, base.I32_extend8_s(v10))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return
							} else {
								F_errfinish(m, int32(505917), int32(3823), int32(324794))
								mBase = m.M
								v36 = m.ExcPending
								if v36 != 0 {
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
		}
	} else {
		switch v10 - int32(73) {
		case 0, 26, 29, 32, 36, 39, 41, 45:
			v39 = *(*int32)(unsafe.Add(mBase, _consts[168]))
			v40 = F_object_ownercheck(m, int32(1259), l0, v39)
			mBase = m.M
			v41 = m.ExcPending
			if v41 != 0 {
				return
			} else {
				if v40 == int32(0) {
					v45 = F_get_rel_relkind(m, l0)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						switch v45 - int32(73) {
						case 0, 32:
							v58 = int32(20)
						default:
							v56 = int32(41)
							v58 = v56
						case 10:
							v58 = int32(37)
						case 29:
							v56 = int32(18)
							v58 = v56
						case 36:
							v58 = int32(23)
						case 45:
							v58 = int32(51)
						}
						F_aclcheck_error(m, int32(2), v58, l1+int32(4))
						mBase = m.M
						v62 = m.ExcPending
						if v62 != 0 {
							return
						} else {
							v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[459])))
							if v64 == int32(0) {
								v68 = int32(1)
								if base.Ui32(l0) < base.Ui32(int32(12000)) {
									v76 = v68
								} else {
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
									if v71 == int32(99) {
										v76 = v68
									} else {
										v74 = F_isTempToastNamespace(m, v71)
										mBase = m.M
										v76 = v74
									}
								}
								if v76 != 0 {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v99 = m.ExcPending
									if v99 != 0 {
										return
									} else {
										F_errcode(m, int32(16797828))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1 + int32(4)
											F_errmsg(m, int32(334761), v8+int32(16))
											mBase = m.M
											v110 = m.ExcPending
											if v110 != 0 {
												return
											} else {
												F_errfinish(m, int32(505917), int32(3835), int32(324794))
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								} else {
									m.G0 = v8 + int32(32)
									return
								}
							} else {
								m.G0 = v8 + int32(32)
								return
							}
						}
					}
				} else {
					v64 = int32(*(*uint8)(unsafe.Add(mBase, _consts[459])))
					if v64 == int32(0) {
						v68 = int32(1)
						if base.Ui32(l0) < base.Ui32(int32(12000)) {
							v76 = v68
						} else {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
							if v71 == int32(99) {
								v76 = v68
							} else {
								v74 = F_isTempToastNamespace(m, v71)
								mBase = m.M
								v76 = v74
							}
						}
						if v76 != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return
							} else {
								F_errcode(m, int32(16797828))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1 + int32(4)
									F_errmsg(m, int32(334761), v8+int32(16))
									mBase = m.M
									v110 = m.ExcPending
									if v110 != 0 {
										return
									} else {
										F_errfinish(m, int32(505917), int32(3835), int32(324794))
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							m.G0 = v8 + int32(32)
							return
						}
					} else {
						m.G0 = v8 + int32(32)
						return
					}
				}
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				F_errcode(m, int32(151027844))
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l1 + int32(4)
					F_errmsg(m, int32(721795), v8)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return
					} else {
						F_errdetail_relkind_not_supported(m, base.I32_extend8_s(v10))
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return
						} else {
							F_errfinish(m, int32(505917), int32(3823), int32(324794))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
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
	}
}
func F_replace_s(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
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
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v16 = F_palloc(m, int32(10))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v31 = v12
	goto L3
L3:
	;
	v33 = l1 - l2 + l3
	if v33 == int32(0) {
		v224 = v31
		goto L9
	} else {
		goto L10
	}
L4:
	;
	return int32(0)
L5:
	;
	if v16 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return int32(-1)
L7:
	;
	goto L8
L8:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1)
	v29 = v16 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v29
	v31 = v29
	goto L3
L9:
	;
	if l3 != 0 {
		goto L70
	} else {
		goto L71
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31-int32(4))))
	v39 = v38 + v33
	v41 = v31 - int32(8)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	if v42 < v39 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v46 = F_repalloc(m, v41, v39+int32(29))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	v62 = v31
	goto L13
L13:
	;
	v63 = l2 + v62
	v64 = v63 + v33
	v65 = v38 - l2
	if v64 == v63 {
		goto L20
	} else {
		goto L21
	}
L14:
	;
	if v46 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	F_pfree(m, v41)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v46))) = v39 + int32(20)
	v60 = v46 + int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v60
	v62 = v60
	goto L13
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	return int32(-1)
L19:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v210-int32(4)))) = v39
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v214 + v33
	v217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l2 <= v217 {
		goto L66
	} else {
		goto L67
	}
L20:
	;
	goto L19
L21:
	;
	v69 = v64 + v65
	if base.Ui32(v63-v69) <= base.Ui32(int32(0)-v65<<(uint(int32(1))%32)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v76 = F___memcpy(m, v64, v63, v65)
	mBase = m.M
	goto L19
L23:
	;
	goto L24
L24:
	;
	v79 = (v64 ^ v63) & int32(3)
	if base.Ui32(v64) < base.Ui32(v63) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	if v181 == int32(0) {
		goto L20
	} else {
		goto L61
	}
L26:
	;
	if base.Ui32(v159) <= base.Ui32(int32(3)) {
		v180 = v158
		v181 = v159
		v182 = v160
		goto L25
	} else {
		goto L57
	}
L27:
	;
	if v79 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	goto L29
L29:
	;
	if v79 != 0 {
		v141 = v65
		goto L40
	} else {
		goto L41
	}
L30:
	;
	v180 = v63
	v181 = v65
	v182 = v64
	goto L25
L31:
	;
	goto L32
L32:
	;
	if v64&int32(3) == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v158 = v63
	v159 = v65
	v160 = v64
	goto L26
L34:
	;
	goto L35
L35:
	;
	v86 = v63
	v87 = v65
	v88 = v64
	goto L36
L36:
	;
	if v87 == int32(0) {
		goto L20
	} else {
		goto L38
	}
L37:
	;
	v158 = v95
	v159 = v97
	v160 = v99
	goto L26
L38:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86))))
	*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v92)
	v94 = int32(1)
	v95 = v86 + v94
	v97 = v87 - v94
	v99 = v88 + v94
	if v99&int32(3) != 0 {
		v86 = v95
		v87 = v97
		v88 = v99
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	if v141 == int32(0) {
		goto L20
	} else {
		goto L53
	}
L41:
	;
	if v69&int32(3) != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v106 = v65
	goto L45
L43:
	;
	v121 = v65
	goto L44
L44:
	;
	if base.Ui32(v121) <= base.Ui32(int32(3)) {
		v141 = v121
		goto L40
	} else {
		goto L49
	}
L45:
	;
	if v106 == int32(0) {
		goto L20
	} else {
		goto L47
	}
L46:
	;
	v121 = v112
	goto L44
L47:
	;
	v112 = v106 - int32(1)
	v113 = v64 + v112
	v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v112))))
	*(*uint8)(unsafe.Add(mBase, uint32(v113))) = uint8(v115)
	if v113&int32(3) != 0 {
		v106 = v112
		goto L45
	} else {
		goto L48
	}
L48:
	;
	goto L46
L49:
	;
	v128 = v121
	goto L50
L50:
	;
	v132 = v128 - int32(4)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v63+v132)))
	*(*int32)(unsafe.Add(mBase, uint32(v64+v132))) = v135
	if base.Ui32(int32(3)) < base.Ui32(v132) {
		v128 = v132
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v141 = v132
	goto L40
L52:
	;
	goto L51
L53:
	;
	v148 = v141
	goto L54
L54:
	;
	v152 = v148 - int32(1)
	v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63+v152))))
	*(*uint8)(unsafe.Add(mBase, uint32(v64+v152))) = uint8(v155)
	if v152 != 0 {
		v148 = v152
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L20
L56:
	;
	goto L55
L57:
	;
	v165 = v158
	v166 = v159
	v167 = v160
	goto L58
L58:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165)))
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v169
	v171 = int32(4)
	v172 = v165 + v171
	v174 = v167 + v171
	v176 = v166 - v171
	if base.Ui32(int32(3)) < base.Ui32(v176) {
		v165 = v172
		v166 = v176
		v167 = v174
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v180 = v172
	v181 = v176
	v182 = v174
	goto L25
L60:
	;
	goto L59
L61:
	;
	v187 = v180
	v188 = v181
	v189 = v182
	goto L62
L62:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v187))))
	*(*uint8)(unsafe.Add(mBase, uint32(v189))) = uint8(v191)
	v193 = int32(1)
	v198 = v188 - v193
	if v198 != 0 {
		v187 = v187 + v193
		v188 = v198
		v189 = v189 + v193
		goto L62
	} else {
		goto L64
	}
L63:
	;
	goto L20
L64:
	;
	goto L63
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v221
	v224 = v210
	goto L9
L66:
	;
	v221 = v217 + v33
	goto L65
L67:
	;
	goto L68
L68:
	;
	if v217 <= l1 {
		v224 = v210
		goto L9
	} else {
		goto L69
	}
L69:
	;
	v221 = l1
	goto L65
L70:
	;
	v228 = l1 + v224
	if v228 == l4 {
		goto L74
	} else {
		goto L75
	}
L71:
	;
	goto L72
L72:
	;
	if l5 != 0 {
		goto L119
	} else {
		goto L120
	}
L73:
	;
	goto L72
L74:
	;
	goto L73
L75:
	;
	v232 = v228 + l3
	if base.Ui32(l4-v232) <= base.Ui32(int32(0)-l3<<(uint(int32(1))%32)) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v239 = F___memcpy(m, v228, l4, l3)
	mBase = m.M
	goto L73
L77:
	;
	goto L78
L78:
	;
	v242 = (v228 ^ l4) & int32(3)
	if base.Ui32(v228) < base.Ui32(l4) {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	if v344 == int32(0) {
		goto L74
	} else {
		goto L115
	}
L80:
	;
	if base.Ui32(v322) <= base.Ui32(int32(3)) {
		v343 = v321
		v344 = v322
		v345 = v323
		goto L79
	} else {
		goto L111
	}
L81:
	;
	if v242 != 0 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	if v242 != 0 {
		v304 = l3
		goto L94
	} else {
		goto L95
	}
L84:
	;
	v343 = l4
	v344 = l3
	v345 = v228
	goto L79
L85:
	;
	goto L86
L86:
	;
	if v228&int32(3) == int32(0) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v321 = l4
	v322 = l3
	v323 = v228
	goto L80
L88:
	;
	goto L89
L89:
	;
	v249 = l4
	v250 = l3
	v251 = v228
	goto L90
L90:
	;
	if v250 == int32(0) {
		goto L74
	} else {
		goto L92
	}
L91:
	;
	v321 = v258
	v322 = v260
	v323 = v262
	goto L80
L92:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
	*(*uint8)(unsafe.Add(mBase, uint32(v251))) = uint8(v255)
	v257 = int32(1)
	v258 = v249 + v257
	v260 = v250 - v257
	v262 = v251 + v257
	if v262&int32(3) != 0 {
		v249 = v258
		v250 = v260
		v251 = v262
		goto L90
	} else {
		goto L93
	}
L93:
	;
	goto L91
L94:
	;
	if v304 == int32(0) {
		goto L74
	} else {
		goto L107
	}
L95:
	;
	if v232&int32(3) != 0 {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	v269 = l3
	goto L99
L97:
	;
	v284 = l3
	goto L98
L98:
	;
	if base.Ui32(v284) <= base.Ui32(int32(3)) {
		v304 = v284
		goto L94
	} else {
		goto L103
	}
L99:
	;
	if v269 == int32(0) {
		goto L74
	} else {
		goto L101
	}
L100:
	;
	v284 = v275
	goto L98
L101:
	;
	v275 = v269 - int32(1)
	v276 = v228 + v275
	v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v275))))
	*(*uint8)(unsafe.Add(mBase, uint32(v276))) = uint8(v278)
	if v276&int32(3) != 0 {
		v269 = v275
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	v291 = v284
	goto L104
L104:
	;
	v295 = v291 - int32(4)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l4+v295)))
	*(*int32)(unsafe.Add(mBase, uint32(v228+v295))) = v298
	if base.Ui32(int32(3)) < base.Ui32(v295) {
		v291 = v295
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v304 = v295
	goto L94
L106:
	;
	goto L105
L107:
	;
	v311 = v304
	goto L108
L108:
	;
	v315 = v311 - int32(1)
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v315))))
	*(*uint8)(unsafe.Add(mBase, uint32(v228+v315))) = uint8(v318)
	if v315 != 0 {
		v311 = v315
		goto L108
	} else {
		goto L110
	}
L109:
	;
	goto L74
L110:
	;
	goto L109
L111:
	;
	v328 = v321
	v329 = v322
	v330 = v323
	goto L112
L112:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v328)))
	*(*int32)(unsafe.Add(mBase, uint32(v330))) = v332
	v334 = int32(4)
	v335 = v328 + v334
	v337 = v330 + v334
	v339 = v329 - v334
	if base.Ui32(int32(3)) < base.Ui32(v339) {
		v328 = v335
		v329 = v339
		v330 = v337
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v343 = v335
	v344 = v339
	v345 = v337
	goto L79
L114:
	;
	goto L113
L115:
	;
	v350 = v343
	v351 = v344
	v352 = v345
	goto L116
L116:
	;
	v354 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	*(*uint8)(unsafe.Add(mBase, uint32(v352))) = uint8(v354)
	v356 = int32(1)
	v361 = v351 - v356
	if v361 != 0 {
		v350 = v350 + v356
		v351 = v361
		v352 = v352 + v356
		goto L116
	} else {
		goto L118
	}
L117:
	;
	goto L74
L118:
	;
	goto L117
L119:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v33
	goto L121
L120:
	;
	goto L121
L121:
	;
	return int32(0)
}
func F_reservoir_get_next_S(m *base.Module, l0 int32, l1 float64, l2 int32) float64 {
	mBase := m.M
	_ = mBase
	var v4 float64
	_ = v4
	var v16 float64
	_ = v16
	var v21 int32
	_ = v21
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v62 float64
	_ = v62
	var v66 float64
	_ = v66
	var v68 float64
	_ = v68
	var v75 float64
	_ = v75
	var v76 float64
	_ = v76
	var v79 float64
	_ = v79
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v90 float64
	_ = v90
	var v93 float64
	_ = v93
	var v95 float64
	_ = v95
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v103 float64
	_ = v103
	var v107 float64
	_ = v107
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v144 float64
	_ = v144
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v151 float64
	_ = v151
	var v155 float64
	_ = v155
	var v157 float64
	_ = v157
	var v159 float64
	_ = v159
	var v162 float64
	_ = v162
	var v165 float64
	_ = v165
	var v171 float64
	_ = v171
	var v172 int32
	_ = v172
	var v173 float64
	_ = v173
	var v176 float64
	_ = v176
	var v180 float64
	_ = v180
	var v181 float64
	_ = v181
	var v182 float64
	_ = v182
	var v193 float64
	_ = v193
	var v194 float64
	_ = v194
	var v197 float64
	_ = v197
	var v204 float64
	_ = v204
	var v231 int64
	_ = v231
	var v232 int64
	_ = v232
	var v233 int64
	_ = v233
	var v254 float64
	_ = v254
	var v257 float64
	_ = v257
	var v259 float64
	_ = v259
	var v260 float64
	_ = v260
	var v263 float64
	_ = v263
	var v271 float64
	_ = v271
	var v291 float64
	_ = v291
	v4 = float64(0)
	v16 = base.F64_convert_i32_s(l2)
	if base.F64_ge(base.F64_mul(v16, float64(22)), l1) != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return v291
L2:
	;
	v21 = l0 + int32(8)
	goto L5
L3:
	;
	goto L4
L4:
	;
	v95 = float64(1)
	v96 = base.F64_add(l1, v95)
	v97 = base.F64_sub(l1, v16)
	v99 = base.F64_add(v97, v95)
	v100 = base.F64_div(v96, v99)
	v102 = l0 + int32(8)
	v103 = *(*float64)(unsafe.Add(mBase, uint32(l0)))
	v107 = v103
	goto L13
L5:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v41 = v39 ^ v40
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = base.I64_rotl(v41, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v41<<(uint(int64(16))%64) ^ base.I64_rotl(v39, int64(24)) ^ v41
	v62 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v39*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L7
L6:
	;
	v66 = base.F64_add(l1, float64(1))
	v68 = base.F64_div(base.F64_sub(v66, v16), v66)
	if base.F64_gt(v68, v62) == int32(0) {
		v291 = v4
		goto L1
	} else {
		goto L9
	}
L7:
	;
	if base.F64_eq(v62, float64(0)) != 0 {
		goto L5
	} else {
		goto L8
	}
L8:
	;
	goto L6
L9:
	;
	v75 = v66
	v76 = v68
	v79 = v4
	goto L10
L10:
	;
	v87 = float64(1)
	v88 = base.F64_add(v79, v87)
	v90 = base.F64_add(v75, v87)
	v93 = base.F64_mul(v76, base.F64_div(base.F64_sub(v90, v16), v90))
	if base.F64_gt(v93, v62) != 0 {
		v75 = v90
		v76 = v93
		v79 = v88
		goto L10
	} else {
		goto L12
	}
L11:
	;
	v291 = v88
	goto L1
L12:
	;
	goto L11
L13:
	;
	v121 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
	v122 = *(*int64)(unsafe.Add(mBase, uint32(v102)+8))
	v123 = v121 ^ v122
	*(*int64)(unsafe.Add(mBase, uint32(v102)+8)) = base.I64_rotl(v123, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v102))) = v123<<(uint(int64(16))%64) ^ base.I64_rotl(v121, int64(24)) ^ v123
	v144 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v121*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L15
L14:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l0))) = v271
	v291 = v150
	goto L1
L15:
	;
	if base.F64_eq(v144, float64(0)) != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v149 = base.F64_mul(l1, base.F64_add(v107, float64(-1)))
	v150 = base.F64_floor(v149)
	v151 = base.F64_add(v99, v150)
	v155 = base.F64_add(l1, v149)
	v157 = F_log(m, base.F64_div(base.F64_mul(v151, base.F64_mul(v100, base.F64_mul(v100, v144))), v155))
	mBase = m.M
	v159 = F_exp(m, base.F64_div(v157, v16))
	mBase = m.M
	v162 = base.F64_div(base.F64_mul(v99, base.F64_div(v155, v151)), l1)
	if base.F64_le(v159, v162) != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L14
L18:
	;
	v271 = base.F64_div(v162, v159)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v165 = base.F64_add(l1, v150)
	v171 = base.F64_div(base.F64_mul(base.F64_add(v165, float64(1)), base.F64_div(base.F64_mul(v96, v144), v99)), v155)
	v172 = base.F64_lt(v16, v150)
	if v172 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v173 = v151
	goto L23
L22:
	;
	v173 = v96
	goto L23
L23:
	;
	if base.F64_le(v173, v165) != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	if v172 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v204 = v171
	goto L26
L26:
	;
	goto L33
L27:
	;
	v176 = l1
	goto L29
L28:
	;
	v176 = base.F64_add(v97, v150)
	goto L29
L29:
	;
	v180 = v165
	v181 = v176
	v182 = v171
	goto L30
L30:
	;
	v193 = base.F64_mul(v182, base.F64_div(v180, v181))
	v194 = float64(-1)
	v197 = base.F64_add(v180, v194)
	if base.F64_ge(v197, v173) != 0 {
		v180 = v197
		v181 = base.F64_add(v181, v194)
		v182 = v193
		goto L30
	} else {
		goto L32
	}
L31:
	;
	v204 = v193
	goto L26
L32:
	;
	goto L31
L33:
	;
	v231 = *(*int64)(unsafe.Add(mBase, uint32(v102)))
	v232 = *(*int64)(unsafe.Add(mBase, uint32(v102)+8))
	v233 = v231 ^ v232
	*(*int64)(unsafe.Add(mBase, uint32(v102)+8)) = base.I64_rotl(v233, int64(37))
	*(*int64)(unsafe.Add(mBase, uint32(v102))) = v233<<(uint(int64(16))%64) ^ base.I64_rotl(v231, int64(24)) ^ v233
	v254 = F_ldexp(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v231*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
	mBase = m.M
	goto L35
L34:
	;
	v257 = F_log(m, v204)
	mBase = m.M
	v259 = F_exp(m, base.F64_div(v257, v16))
	mBase = m.M
	v260 = F_log(m, v254)
	mBase = m.M
	v263 = F_exp(m, base.F64_div(base.F64_neg(v260), v16))
	mBase = m.M
	if base.F64_le(v259, base.F64_div(v155, l1)) == int32(0) {
		v107 = v263
		goto L13
	} else {
		goto L37
	}
L35:
	;
	if base.F64_eq(v254, float64(0)) != 0 {
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v271 = v263
	goto L17
}
func F_resolve_anyelement_from_others(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
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
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	v5 = m.G0
	v7 = v5 + int32(-64)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9 != 0 {
		v10 = F_getBaseType(m, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v12 = F_get_element_type(m, v10)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				if v12 != 0 {
					v109 = v12
					*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
					m.G0 = v7 - int32(-64)
					return
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return
					} else {
						F_errcode(m, int32(67141764))
						mBase = m.M
						v20 = m.ExcPending
						if v20 != 0 {
							return
						} else {
							v21 = F_format_type_be(m, v10)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7)+52)) = v21
								*(*int32)(unsafe.Add(mBase, uint32(v7)+48)) = int32(24207)
								F_errmsg(m, int32(192478), v5+int32(-16))
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return
								} else {
									F_errfinish(m, int32(509710), int32(602), int32(137226))
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
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
			}
		}
	} else {
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v36 != 0 {
			v37 = F_getBaseType(m, v36)
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return
			} else {
				v39 = F_get_range_subtype(m, v37)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					if v39 != 0 {
						v109 = v39
						*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
						m.G0 = v7 - int32(-64)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errcode(m, int32(67141764))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								v48 = F_format_type_be(m, v37)
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+36)) = v48
									*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = int32(409441)
									F_errmsg(m, int32(192585), v5+int32(-32))
									mBase = m.M
									v57 = m.ExcPending
									if v57 != 0 {
										return
									} else {
										F_errfinish(m, int32(509710), int32(616), int32(137226))
										mBase = m.M
										v62 = m.ExcPending
										if v62 != 0 {
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
				}
			}
		} else {
			v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			if v63 != 0 {
				v64 = F_getBaseType(m, v63)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return
				} else {
					v66 = F_get_multirange_range(m, v64)
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						if v66 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return
							} else {
								F_errcode(m, int32(67141764))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return
								} else {
									v122 = F_format_type_be(m, v64)
									mBase = m.M
									v123 = m.ExcPending
									if v123 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v122
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(409450)
										F_errmsg(m, int32(192527), v7)
										mBase = m.M
										v129 = m.ExcPending
										if v129 != 0 {
											return
										} else {
											F_errfinish(m, int32(509710), int32(634), int32(137226))
											mBase = m.M
											v134 = m.ExcPending
											if v134 != 0 {
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
						} else {
							v70 = F_getBaseType(m, v66)
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return
							} else {
								v72 = F_get_range_subtype(m, v70)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return
								} else {
									if v72 != 0 {
										v109 = v72
										*(*int32)(unsafe.Add(mBase, uint32(l0))) = v109
										m.G0 = v7 - int32(-64)
										return
									} else {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v77 = m.ExcPending
										if v77 != 0 {
											return
										} else {
											F_errcode(m, int32(67141764))
											mBase = m.M
											v80 = m.ExcPending
											if v80 != 0 {
												return
											} else {
												v81 = F_format_type_be(m, v70)
												mBase = m.M
												v82 = m.ExcPending
												if v82 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = v81
													*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = int32(409450)
													F_errmsg(m, int32(192638), v5+int32(-48))
													mBase = m.M
													v90 = m.ExcPending
													if v90 != 0 {
														return
													} else {
														F_errfinish(m, int32(509710), int32(644), int32(137226))
														mBase = m.M
														v95 = m.ExcPending
														if v95 != 0 {
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
								}
							}
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v99 = m.ExcPending
				if v99 != 0 {
					return
				} else {
					F_errmsg_internal(m, int32(379043), int32(0))
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return
					} else {
						F_errfinish(m, int32(509710), int32(648), int32(137226))
						mBase = m.M
						v108 = m.ExcPending
						if v108 != 0 {
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
}
func F_rfree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v8 != int32(65239) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(0)
	if v13 == v11 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v18 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+72)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v18
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v13)+92))
	if v22 != v13+int32(180) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	F_pfree(m, v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	if v28 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return
L9:
	;
	goto L7
L10:
	;
	F_pfree(m, v28)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v13)+160))
	if v31 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L12
L14:
	;
	F_pfree(m, v31)
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L8
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v13)+164))
	if v34 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	goto L16
L18:
	;
	F_pfree(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L8
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v37 != 0 {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	F_freesubre(m, int32(0), v37)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L8
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v13)+424))
	if v41 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L24
L26:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v13)+428))
	v44 = v42 - int32(1)
	if int32(0) < v44 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	if v84 != 0 {
		goto L42
	} else {
		goto L43
	}
L29:
	;
	v47 = v41
	v49 = v44
	goto L32
L30:
	;
	goto L31
L31:
	;
	F_pfree(m, v41)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L41
	}
L32:
	;
	v53 = v47 + int32(124)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v54 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v47)+152))
	F_pfree(m, v55)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L8
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v68 = int32(1)
	if v68 < v49 {
		v47 = v47 + int32(88)
		v49 = v49 - v68
		goto L32
	} else {
		goto L40
	}
L37:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v47)+156))
	F_pfree(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v47)+160))
	F_pfree(m, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L8
	} else {
		goto L39
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = int32(0)
	goto L36
L40:
	;
	goto L33
L41:
	;
	goto L28
L42:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
	F_pfree(m, v85)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L8
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	F_pfree(m, v13)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L8
	} else {
		goto L48
	}
L45:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	F_pfree(m, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v13)+56))
	F_pfree(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = int32(0)
	goto L44
L48:
	;
	goto L1
}
