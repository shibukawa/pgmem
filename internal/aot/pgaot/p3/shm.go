package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_shm_mq_attach(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v11 int32
	_ = v11
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	v5 = F_palloc(m, int32(44))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5)+12)) = v9
		v11 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = v11
		*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v5))) = l0
		*(*int64)(unsafe.Add(mBase, uint32(v5)+20)) = v9
		*(*int64)(unsafe.Add(mBase, uint32(v5)+28)) = v9
		*(*uint16)(unsafe.Add(mBase, uint32(v5)+36)) = uint16(v11)
		v22 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v5)+40)) = v22
		if l1 != 0 {
			F_on_dsm_detach(m, l1, int32(1104), l0)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				return v5
			}
		} else {
			return v5
		}
	}
}
func F_shm_mq_detach(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int64
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var __phi46 int32
	_ = __phi46
	var v47 int32
	_ = v47
	var __phi47 int32
	_ = __phi47
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(v8)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v9 + base.I64_extend_i32_u(v7)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(0)
	goto L3
L2:
	;
	goto L3
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(1)
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	F_s_lock(m, v16, int32(496776), int32(886), int32(312814))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v27 = *(*int32)(unsafe.Add(mBase, _consts[137]))
	if v25 == v27 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return
L8:
	;
	goto L6
L9:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v30 = v29
	goto L11
L10:
	;
	v30 = v25
	goto L11
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(0)
	v33 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+36)) = uint8(v33)
	if v30 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	F_SetLatch(m, v30+int32(20))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v39 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	goto L14
L16:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	if v41 != 0 {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	goto L18
L18:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v81 != 0 {
		goto L30
	} else {
		goto L31
	}
L19:
	;
	goto L18
L20:
	;
	__phi46 = v41
	__phi47 = v39 + int32(32)
	v46 = __phi46
	v47 = __phi47
	goto L23
L21:
	;
	goto L22
L22:
	;
	goto L19
L23:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	v52 = v46 - int32(8)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 != int32(1104) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L22
L25:
	;
	if v50 != 0 {
		__phi46 = v50
		__phi47 = v46
		v46 = __phi46
		v47 = __phi47
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v46-int32(4))))
	if v58 != v40 {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v50
	F_pfree(m, v52)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	goto L19
L29:
	;
	goto L24
L30:
	;
	F_pfree(m, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_pfree(m, l0)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L7
	} else {
		goto L34
	}
L33:
	;
	goto L32
L34:
	;
	return
}
func F_shm_mq_detach_callback(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(1)
	if v3 != 0 {
		F_s_lock(m, l1, int32(496776), int32(886), int32(312814))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v13 = *(*int32)(unsafe.Add(mBase, _consts[137]))
			if v11 == v13 {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
				v16 = v15
			} else {
				v16 = v11
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
			v19 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v19)
			if v16 != 0 {
				F_SetLatch(m, v16+int32(20))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return
				} else {
					return
				}
			} else {
				return
			}
		}
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
		v13 = *(*int32)(unsafe.Add(mBase, _consts[137]))
		if v11 == v13 {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			v16 = v15
		} else {
			v16 = v11
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
		v19 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+36)) = uint8(v19)
		if v16 != 0 {
			F_SetLatch(m, v16+int32(20))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				return
			}
		} else {
			return
		}
	}
}
func F_shm_mq_receive_bytes(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v18 int64
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int64
	_ = v25
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v37 int32
	_ = v37
	var v40 int64
	_ = v40
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v53 int64
	_ = v53
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int64
	_ = v85
	var v87 int64
	_ = v87
	var v89 int32
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v92 int64
	_ = v92
	var v93 int64
	_ = v93
	var v105 int64
	_ = v105
	var v107 int64
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v16
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = base.I64_extend_i32_u(v20)
	v22 = v18 + v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v15)+32))
	v24 = base.I64_extend_i32_u(v23)
	v25 = base.I64_rem_u_s(v22, v24)
	v26 = v16 - v22
	v27 = base.I64_extend_i32_u(l1)
	if base.Ui64(v27) <= base.Ui64(v26) {
		v105 = v26
		v107 = v25
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v112 = base.I32_wrap_i64(v107)
	v113 = v23 - v112
	if base.Ui64(v105) < base.Ui64(base.I64_extend_i32_u(v113)) {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	if base.Ui64(v24) <= base.Ui64(v26+v25) {
		v105 = v26
		v107 = v25
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v37 = v20
	v40 = v16
	v43 = v21
	goto L4
L4:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+36)))
	if v45 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v105 = v93
	v107 = v92
	goto L1
L6:
	;
	v85 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v85
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v90 = base.I64_extend_i32_u(v89)
	v91 = v87 + v90
	v92 = base.I64_rem_u_s(v91, v24)
	v93 = v85 - v91
	if base.Ui64(v27) <= base.Ui64(v93) {
		v105 = v93
		v107 = v92
		goto L1
	} else {
		goto L23
	}
L7:
	;
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v15)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v48
	if v48 != v40 {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v37 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	return int32(2)
L11:
	;
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v15)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v53 + v43
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
	F_SetLatch(m, v56+int32(20))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	if l2 != 0 {
		goto L16
	} else {
		goto L17
	}
L14:
	;
	return int32(0)
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = int32(0)
	goto L13
L16:
	;
	return int32(1)
L17:
	;
	goto L18
L18:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v72 = F_WaitLatch(m, v68, int32(33), int32(0), int32(134217763))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(0)
	goto L20
L20:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v79 == int32(0) {
		goto L6
	} else {
		goto L21
	}
L21:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L14
	} else {
		goto L22
	}
L22:
	;
	goto L6
L23:
	;
	if base.Ui64(v93+v92) < base.Ui64(v24) {
		v37 = v89
		v40 = v85
		v43 = v90
		goto L4
	} else {
		goto L24
	}
L24:
	;
	goto L5
L25:
	;
	v116 = base.I32_wrap_i64(v105)
	goto L27
L26:
	;
	v116 = v113
	goto L27
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v116
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+37)))
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v15 + (v118 + v112) + int32(38)
	return int32(0)
}
func F_shm_mq_wait_internal(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	goto L2
L1:
	;
	m.G0 = v9 + int32(16)
	return v62
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1)
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v62 = (v30 ^ int32(-1)) & base.B2i32(v27 != int32(0))
	goto L1
L4:
	;
	F_s_lock(m, l0, int32(496776), int32(1228), int32(311742))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(0)
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	goto L6
L9:
	;
	goto L3
L10:
	;
	if v27 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	if l2 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	v45 = F_WaitLatch(m, v41, int32(33), int32(0), int32(134217761))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L16
	}
L13:
	;
	v35 = F_GetBackgroundWorkerPid(m, l2, v9+int32(12))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L7
	} else {
		goto L14
	}
L14:
	;
	if base.Ui32(v35) <= base.Ui32(int32(1)) {
		goto L12
	} else {
		goto L15
	}
L15:
	;
	v62 = int32(0)
	goto L1
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[311]))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = int32(0)
	goto L17
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	if v52 == int32(0) {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	goto L2
}
func F_shm_toc_attach(m *base.Module, l0 int64, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	if v4 == l0 {
		v6 = l1
	} else {
		v6 = int32(0)
	}
	return v6
}
func F_shm_toc_insert(m *base.Module, l0 int32, l1 int64, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(1)
	v11 = l0 + int32(8)
	if v7 != 0 {
		F_s_lock(m, v11, int32(501011), int32(184), int32(82145))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v22 = v18 + v19<<(uint(int32(4))%32)
			if base.Ui32(v17) < base.Ui32(v22+int32(40)) {
				v32 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v32
				F_errstart_cold(m, int32(21), v32)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errcode(m, int32(8389))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errmsg(m, int32(14090), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errfinish(m, int32(501011), int32(200), int32(82145))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
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
				if v19 == int32(-1) {
					v32 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v32
					F_errstart_cold(m, int32(21), v32)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_errcode(m, int32(8389))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errmsg(m, int32(14090), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errfinish(m, int32(501011), int32(200), int32(82145))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
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
					if base.Ui32(v22+int32(24)) < base.Ui32(int32(-16)) {
						v52 = l0 + v19<<(uint(int32(4))%32)
						*(*int64)(unsafe.Add(mBase, uint32(v52)+24)) = l1
						*(*int32)(unsafe.Add(mBase, uint32(v52)+32)) = l2 - l0
						*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
						v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58 + int32(1)
						return
					} else {
						v32 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v32
						F_errstart_cold(m, int32(21), v32)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return
						} else {
							F_errcode(m, int32(8389))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
								return
							} else {
								F_errmsg(m, int32(14090), int32(0))
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return
								} else {
									F_errfinish(m, int32(501011), int32(200), int32(82145))
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
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
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v22 = v18 + v19<<(uint(int32(4))%32)
		if base.Ui32(v17) < base.Ui32(v22+int32(40)) {
			v32 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v11))) = v32
			F_errstart_cold(m, int32(21), v32)
			mBase = m.M
			v37 = m.ExcPending
			if v37 != 0 {
				return
			} else {
				F_errcode(m, int32(8389))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_errmsg(m, int32(14090), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						F_errfinish(m, int32(501011), int32(200), int32(82145))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
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
			if v19 == int32(-1) {
				v32 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v32
				F_errstart_cold(m, int32(21), v32)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_errcode(m, int32(8389))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_errmsg(m, int32(14090), int32(0))
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							F_errfinish(m, int32(501011), int32(200), int32(82145))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
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
				if base.Ui32(v22+int32(24)) < base.Ui32(int32(-16)) {
					v52 = l0 + v19<<(uint(int32(4))%32)
					*(*int64)(unsafe.Add(mBase, uint32(v52)+24)) = l1
					*(*int32)(unsafe.Add(mBase, uint32(v52)+32)) = l2 - l0
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
					v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v58 + int32(1)
					return
				} else {
					v32 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v32
					F_errstart_cold(m, int32(21), v32)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_errcode(m, int32(8389))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							F_errmsg(m, int32(14090), int32(0))
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return
							} else {
								F_errfinish(m, int32(501011), int32(200), int32(82145))
								mBase = m.M
								v49 = m.ExcPending
								if v49 != 0 {
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
