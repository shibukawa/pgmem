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
			v18 = l0 + (v12 | v13<<(uint(int32(16))%32))
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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[619]))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+192)))
	if v6&int32(2) != 0 {
		v9 = int32(0)
		F_InitPostgres(m, v9, l0, v9, l1, l2<<(uint(int32(1))%32)&int32(6), v9)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, _consts[298]))
			if v19 != int32(1) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v44 = m.ExcPending
				if v44 != 0 {
					return
				} else {
					F_errmsg(m, int32(232382), int32(0))
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						F_errfinish(m, int32(519875), int32(913), int32(458668))
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
				*(*int32)(unsafe.Add(mBase, _consts[298])) = int32(2)
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
				F_errmsg(m, int32(272207), int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return
				} else {
					F_errfinish(m, int32(519875), int32(903), int32(458668))
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
		F_s_lock(m, l0, int32(519902), int32(215), int32(84185))
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
	v6 = *(*int32)(unsafe.Add(mBase, _consts[753]))
	if base.Ui32(int32(2)) <= base.Ui32(v6) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v10 = *(*int32)(unsafe.Add(mBase, _consts[752]))
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
	*(*uint8)(unsafe.Add(mBase, _consts[754])) = uint8(v43)
	v46 = *(*int32)(unsafe.Add(mBase, _consts[755]))
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
	v30 = *(*int32)(unsafe.Add(mBase, _consts[753]))
	v32 = *(*int32)(unsafe.Add(mBase, _consts[752]))
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
	v54 = *(*int32)(unsafe.Add(mBase, _consts[756]))
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
	v58 = *(*int32)(unsafe.Add(mBase, _consts[755]))
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
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if base.Ui32(v3) < base.Ui32(v4) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v9 = base.B2i32(v6 < v7)
	} else {
		v9 = int32(0)
	}
	return v9
}
func F___bswap_32(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = int32(24)
	v4 = int32(65280)
	v6 = int32(8)
	return l0<<(uint(v2)%32) | l0&v4<<(uint(v6)%32) | (int32(base.Ui32(l0)>>(uint(v6)%32))&v4 | int32(base.Ui32(l0)>>(uint(v2)%32)))
}
func F_before_shmem_exit(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v10 int32
	_ = v10
	var v22 int32
	_ = v22
	var v30 int32
	_ = v30
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	v6 = *(*int32)(unsafe.Add(mBase, _consts[775]))
	if v6 < int32(20) {
		v10 = v6 << (uint(int32(3)) % 32)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[776]))) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v10)+uint32(_consts[777]))) = l0
		*(*int32)(unsafe.Add(mBase, _consts[775])) = v6 + int32(1)
		v22 = int32(*(*uint8)(unsafe.Add(mBase, _consts[772])))
		if v22 == int32(0) {
			v30 = *(*int32)(unsafe.Add(mBase, _consts[773]))
			if v30 <= int32(31) {
				*(*int32)(unsafe.Add(mBase, _consts[773])) = v30 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v30<<(uint(int32(2))%32))+uint32(_consts[774]))) = int32(1103)
			} else {
			}
			v45 = int32(1)
			*(*uint8)(unsafe.Add(mBase, _consts[772])) = uint8(v45)
		} else {
		}
		return
	} else {
		F_errstart_cold(m, int32(22), int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return
		} else {
			F_errcode(m, int32(261))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return
			} else {
				F_errmsg_internal(m, int32(126554), int32(0))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					F_errfinish(m, int32(525084), int32(349), int32(106137))
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
	var v61 int64
	_ = v61
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	v9 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	if int32(0) <= v9 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _consts[401]))
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
	v155 = m.ExcPending
	if v155 != 0 {
		goto L13
	} else {
		goto L36
	}
L4:
	;
	v80 = *(*int32)(unsafe.Add(mBase, _consts[402]))
	v82 = *(*int32)(unsafe.Add(mBase, _consts[400]))
	v87 = v80 + v82*int32(20) + int32(16)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 == int32(0) {
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
	*(*int32)(unsafe.Add(mBase, _consts[401])) = v41
	*(*int32)(unsafe.Add(mBase, _consts[402])) = v43
	if v41 <= v13 {
		goto L4
	} else {
		goto L19
	}
L7:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _consts[190]))
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
	v32 = *(*int32)(unsafe.Add(mBase, _consts[402]))
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
	v57 = *(*int32)(unsafe.Add(mBase, _consts[402]))
	v60 = v57 + v52*int32(20)
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60))) = v61
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v60)+8)) = v61
	v68 = v52 + int32(1)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[401]))
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
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+9)))
	v149 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v143)+9)) = uint8(v149)
	return v148
L24:
	;
	v124 = int32(4562080)
	v125 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v128 = *(*int32)(unsafe.Add(mBase, _consts[141]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v128
	v131 = F_palloc0(m, int32(36))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L13
	} else {
		goto L34
	}
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)+4))
	if v91 <= int32(0) {
		goto L24
	} else {
		goto L26
	}
L26:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v88)+12))
	v99 = int32(0)
	goto L27
L27:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v94+v99<<(uint(int32(2))%32))))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
	if v107 != l0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L24
L29:
	;
	v115 = v99 + int32(1)
	if v91 != v115 {
		v99 = v115
		goto L27
	} else {
		goto L33
	}
L30:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
	if v109 != l1 {
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+8)))
	if v111 != int32(1) {
		v143 = v106
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
	*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = l0
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	v136 = F_lappend(m, v135, v131)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v136
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v125
	v143 = v131
	goto L23
L36:
	;
	F_errmsg_internal(m, int32(16975), int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L13
	} else {
		goto L37
	}
L37:
	;
	F_errfinish(m, int32(519958), int32(6591), int32(472268))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(290295)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v10
	v14 = *(*int64)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(993)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v14
	v18 = int32(4554984)
	v19 = *(*int32)(unsafe.Add(mBase, _consts[337]))
	*(*int32)(unsafe.Add(mBase, _consts[337])) = v8 + int32(4)
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
		*(*int32)(unsafe.Add(mBase, _consts[337])) = v40
		m.G0 = v8 + int32(32)
		return
	}
}
func F_big5_to_utf8(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_check_encoding_conversion_args(m, v8, v9, v10, int32(36), int32(6))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v18 = int32(0)
		v24 = F_LocalToUtf(m, v6, v10, v5, int32(4431188), v18, v18, v18, int32(36), base.B2i32(v7 != v18))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			return v24
		}
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
	var v30 int32
	_ = v30
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
	v30 = int32(0)
	goto L4
L4:
	;
	v36 = int32(1)
	v37 = v30 << (uint(v36) % 32)
	v39 = v37 | v36
	v41 = v37 + int32(2)
	if v41 < v28 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21+v30<<(uint(int32(2))%32)))) = v25
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
	*(*int32)(unsafe.Add(mBase, uint32(v21+v30<<(uint(int32(2))%32)))) = v77
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v28 = v79
	v30 = v61
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
	var v16 int32
	_ = v16
	var v32 int32
	_ = v32
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v95 int32
	_ = v95
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
	var v136 int32
	_ = v136
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v388 int32
	_ = v388
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v413 int32
	_ = v413
	v16 = l1 + int32(16)
	goto L1
L1:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v32 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v408
L3:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v38 == int32(1) {
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
	if v403 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L9:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v46 = v44 + int32(28)
	v48 = F_LWLockAcquire(m, v46, int32(0))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L6
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v37)+8))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+28))
	if v259 <= v257 {
		v318 = v259
		v321 = v257
		goto L52
	} else {
		goto L53
	}
L12:
	;
	v403 = v250
	goto L8
L13:
	;
	if v43 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v53 = v43 + int32(4)
	goto L16
L15:
	;
	v53 = int32(0)
	goto L16
L16:
	;
	if v42 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v57 = v42 + int32(4)
	goto L19
L18:
	;
	v57 = int32(0)
	goto L19
L19:
	;
	if v41 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v61 = v41 + int32(4)
	goto L22
L21:
	;
	v61 = int32(0)
	goto L22
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v44)+48))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v44)+12))
	if v63 <= v62 {
		v136 = v62
		goto L26
	} else {
		goto L27
	}
L23:
	;
	if v195 < v199 {
		goto L46
	} else {
		goto L47
	}
L24:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v53+v136<<(uint(int32(2))%32))))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v61+v165*int32(48))))
	v170 = v161 + v169
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v44)+44))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	if v171 < v172 {
		goto L41
	} else {
		goto L42
	}
L25:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v44)+44))
	v195 = v160
	v199 = v159
	goto L23
L26:
	;
	if v136 < v63 {
		goto L24
	} else {
		goto L40
	}
L27:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	v70 = v65
	v72 = v62
	goto L28
L28:
	;
	v80 = int32(256)
	if v70 <= v80 {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L25
L30:
	;
	v83 = v80
	goto L32
L31:
	;
	v83 = v70
	goto L32
L32:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v53+v72<<(uint(int32(2))%32))))
	v95 = v70
	goto L34
L33:
	;
	v123 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = v123
	v127 = v72 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v44)+48)) = v127
	if v127 != v63 {
		v70 = v123
		v72 = v127
		goto L28
	} else {
		goto L39
	}
L34:
	;
	if v95 == v83 {
		goto L33
	} else {
		goto L36
	}
L35:
	;
	if int32(255) < v95 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	v106 = int32(1)
	v109 = base.I32_div_s(v95, int32(32))
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v61+v87*int32(48)+v109<<(uint(int32(2))%32))+8))
	if int32(base.Ui32(v113)>>(uint(v95)%32))&v106 == int32(0) {
		v95 = v95 + v106
		goto L34
	} else {
		goto L37
	}
L37:
	;
	goto L35
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = v95
	v136 = v72
	goto L26
L39:
	;
	goto L29
L40:
	;
	goto L25
L41:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v57+v171<<(uint(int32(2))%32))))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v61+v177*int32(48))))
	if base.Ui32(v181) <= base.Ui32(v170) {
		v195 = v171
		v199 = v172
		goto L23
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v185 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v185)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v170
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v44)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+52)) = v188 + int32(1)
	F_LWLockRelease(m, v46)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L6
	} else {
		goto L45
	}
L44:
	;
	goto L43
L45:
	;
	v250 = int32(1)
	goto L12
L46:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(v57+v195<<(uint(int32(2))%32))))
	v216 = v61 + v213*int32(48)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v216
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	v219 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v219)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v218
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v222)
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v44)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v44)+44)) = v224 + int32(1)
	F_LWLockRelease(m, v46)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L6
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_LWLockRelease(m, v46)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L6
	} else {
		goto L50
	}
L49:
	;
	v250 = int32(1)
	goto L12
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	v250 = int32(0)
	goto L12
L51:
	;
	v403 = v388
	goto L8
L52:
	;
	if v318 <= v321 {
		goto L67
	} else {
		goto L68
	}
L53:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v266 = v261
	v267 = v257
	goto L54
L54:
	;
	v270 = int32(256)
	if v266 <= v270 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v318 = v314
	v321 = v312
	goto L52
L56:
	;
	v273 = v270
	goto L58
L57:
	;
	v273 = v266
	goto L58
L58:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v258)+92))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v274+v267<<(uint(int32(2))%32))))
	v285 = v266
	goto L60
L59:
	;
	v308 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v308
	v312 = v267 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+8)) = v312
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v258)+28))
	if v312 < v314 {
		v266 = v308
		v267 = v312
		goto L54
	} else {
		goto L65
	}
L60:
	;
	if v285 == v273 {
		goto L59
	} else {
		goto L62
	}
L61:
	;
	if int32(255) < v285 {
		goto L59
	} else {
		goto L64
	}
L62:
	;
	v290 = int32(1)
	v293 = base.I32_div_s(v285, int32(32))
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v278+int32(8)+v293<<(uint(int32(2))%32))))
	if int32(base.Ui32(v297)>>(uint(v285)%32))&v290 == int32(0) {
		v285 = v285 + v290
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v285
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v258)+28))
	v318 = v306
	v321 = v267
	goto L52
L65:
	;
	goto L55
L66:
	;
	if v355 < v356 {
		goto L74
	} else {
		goto L75
	}
L67:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v258)+24))
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v355 = v326
	v356 = v325
	goto L66
L68:
	;
	goto L69
L69:
	;
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v258)+92))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v328+v321<<(uint(int32(2))%32))))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v332)))
	v334 = v327 + v333
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v258)+24))
	if v335 < v336 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v258)+88))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v338+v335<<(uint(int32(2))%32))))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)))
	if base.Ui32(v343) <= base.Ui32(v334) {
		v355 = v335
		v356 = v336
		goto L66
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = int32(0)
	v347 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)) = uint16(v347)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v334
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v37)+12))
	v351 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+12)) = v350 + v351
	v388 = v351
	goto L51
L73:
	;
	goto L72
L74:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v258)+8))
	if v359 == int32(1) {
		goto L77
	} else {
		goto L78
	}
L75:
	;
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = int32(-1)
	v388 = int32(0)
	goto L51
L77:
	;
	v369 = v258 + int32(40)
	goto L79
L78:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v258)+88))
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v364+v355<<(uint(int32(2))%32))))
	v369 = v368
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+8)) = v369
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v369)))
	v372 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)) = uint8(v372)
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v371
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+5)) = uint8(v375)
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	v378 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v37)+4)) = v377 + v378
	v388 = v378
	goto L51
L80:
	;
	return int32(-1)
L81:
	;
	goto L82
L82:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v410 = *(*int32)(unsafe.Add(mBase, _consts[67]))
	if v410 != int32(3) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if base.Ui32(v413) <= base.Ui32(v408) {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	goto L2
L86:
	;
	goto L85
}
func F_bitshiftright(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
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
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v16 < int32(0) {
			v20 = int32(0)
			v22 = int32(-2147483640)
			if base.Ui32(v16) <= base.Ui32(v22) {
				v25 = v22
			} else {
				v25 = v16
			}
			v27 = F_DirectFunctionCall2Coll(m, int32(1553), v20, v12, v20-v25)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				return v27
			}
		} else {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v33 = F_palloc(m, int32(base.Ui32(v30)>>(uint(int32(2))%32)))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int32(0)
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v37 = v35 & int32(-4)
				*(*int32)(unsafe.Add(mBase, uint32(v33))) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				*(*int32)(unsafe.Add(mBase, uint32(v33)+4)) = v39
				v42 = v33 + int32(8)
				v44 = v42 & int32(3)
				if v39 <= v16 {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v48 = int32(base.Ui32(v46) >> (uint(int32(2)) % 32))
					v50 = v48 - int32(8)
					if v44 != 0 {
						v74 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), v50)
						mBase = m.M
						return v33
					} else {
						if v46&int32(12) != 0 {
							v74 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), v50)
							mBase = m.M
							return v33
						} else {
							if base.Ui32(int32(1024)) < base.Ui32(v50) {
								v74 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), v50)
								mBase = m.M
								return v33
							} else {
								v55 = v33 + v48
								if base.Ui32(v55) <= base.Ui32(v42) {
									return v33
								} else {
									v59 = v33 + int32(12)
									if base.Ui32(v59) < base.Ui32(v55) {
										v61 = v55
									} else {
										v61 = v59
									}
									v70 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), (v61-v33-int32(9))&int32(-4)+int32(4))
									mBase = m.M
									return v33
								}
							}
						}
					}
				} else {
					v77 = v16 & int32(7)
					v79 = int32(base.Ui32(v16) >> (uint(int32(3)) % 32))
					if v44 != 0 {
						v105 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), v79)
						mBase = m.M
					} else {
						if base.Ui32(int32(8199)) < base.Ui32(v16) {
							v105 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), v79)
							mBase = m.M
						} else {
							if v16&int32(24) != 0 {
								v105 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), v79)
								mBase = m.M
							} else {
								if base.Ui32(v42+v79) <= base.Ui32(v42) {
								} else {
									v89 = v33 + v79 + int32(8)
									v91 = v33 + int32(12)
									if base.Ui32(v91) < base.Ui32(v89) {
										v93 = v89
									} else {
										v93 = v91
									}
									v102 = F__emscripten_memset_bulkmem(m, v42, base.I32_extend8_s(int32(0)), (v93-v33-int32(9))&int32(-4)+int32(4))
									mBase = m.M
								}
							}
						}
					}
					v109 = v12 + int32(8)
					v110 = v42 + v79
					if v77 == int32(0) {
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v115 = int32(base.Ui32(v113) >> (uint(int32(2)) % 32))
						v118 = v115 - v79 - int32(8)
						if v118 != 0 {
							v119 = F__emscripten_memcpy_bulkmem(m, v110, v109, v118)
							mBase = m.M
						} else {
						}
						v122 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
						v165 = v33 + v115
						v171 = v122
					} else {
						if base.Ui32(v33+int32(base.Ui32(v35)>>(uint(int32(2))%32))) <= base.Ui32(v110) {
							v165 = v110
							v171 = v37
						} else {
							v127 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v110))) = uint8(v127)
							v131 = v110
							v134 = v109
							for {
								v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131))))
								v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
								v144 = v141 | int32(base.Ui32(v142)>>(uint(v77)%32))
								*(*uint8)(unsafe.Add(mBase, uint32(v131))) = uint8(v144)
								v147 = v131 + int32(1)
								v148 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
								v150 = int32(base.Ui32(v148) >> (uint(int32(2)) % 32))
								if base.Ui32(v147) < base.Ui32(v33+v150) {
									v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
									v154 = v153 << (uint(int32(8)-v77) % 32)
									*(*uint8)(unsafe.Add(mBase, uint32(v147))) = uint8(v154)
									v156 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
									v159 = int32(base.Ui32(v156) >> (uint(int32(2)) % 32))
									v160 = v156
								} else {
									v159 = v150
									v160 = v148
								}
								if base.Ui32(v147) < base.Ui32(v33+v159) {
									v131 = v147
									v134 = v134 + int32(1)
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
					v179 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
					v182 = v171<<(uint(int32(1))%32)&int32(-8) - v179 + int32(-64)
					if v182 <= int32(0) {
					} else {
						v186 = v165 - int32(1)
						v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v186))))
						v190 = v187 & (int32(255) << (uint(v182) % 32))
						*(*uint8)(unsafe.Add(mBase, uint32(v186))) = uint8(v190)
					}
					return v33
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
	var v133 int32
	_ = v133
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
	var v192 int32
	_ = v192
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
	v133 = int32(1)
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
	v146 = F_ReadBufferExtended(m, v143, v144, v133, v144, v94)
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
	v154 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v154+(v146^int32(-1))<<(uint(int32(2))%32))))
	v168 = v160
	goto L33
L35:
	;
	goto L36
L36:
	;
	v162 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	v192 = int32(1)
	v200 = v142
	goto L40
L40:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v19)+1168))
	v207 = v168 + int32(24) + v201*(v192&int32(65535)-int32(1))
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
	v287 = v192 + int32(1)
	if base.Ui32(v287&int32(65535)) <= base.Ui32(v177) {
		v192 = v287
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
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v234+(v207+int32(6))))))
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
	v312 = *(*int32)(unsafe.Add(mBase, _consts[48]))
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
	v316 = v133 + int32(1)
	if v316 != v98 {
		v133 = v316
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
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
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
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
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
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
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
	v17 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v22 = F_AllocSetContextCreateInternal(m, v17, int32(65190), int32(0), int32(8192), int32(8388608))
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
	v26 = int32(4562080)
	v27 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v22
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
		goto L92
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v27
	F_MemoryContextDelete(m, v22)
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L1
	} else {
		goto L91
	}
L9:
	;
	F_LockBuffer(m, v35, int32(2))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L41
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
	v43 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v43+(v35^int32(-1))<<(uint(int32(2))%32))))
	v57 = v49
	goto L10
L12:
	;
	goto L13
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
		goto L40
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
	v105 = int32(8160) - v104
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
	F_PageInit(m, v76, int32(8192), int32(8))
	mBase = m.M
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v76)+16)))
	v91 = v76 + v90
	v92 = int32(65411)
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
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v108 = int32(24)
	v110 = F___memcpy(m, v76+v104+v108, v32, v100)
	mBase = m.M
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102))))
	v113 = v111 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v102))) = uint16(v113)
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v118 = v115*v113 + v108
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+12)) = uint16(v118)
	goto L31
L30:
	;
	goto L31
L31:
	;
	goto L28
L32:
	;
	F_GenericXLogFinish(m, v73)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_pfree(m, v73)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L38
	}
L35:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_ReleaseBuffer(m, v35)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	goto L8
L38:
	;
	F_UnlockReleaseBuffer(m, v68)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v136 = v64
	goto L9
L40:
	;
	v136 = int32(-1)
	goto L9
L41:
	;
	v143 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+28)))
	v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+30)))
	if base.Ui32(v143) < base.Ui32(v144) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	F_UnlockReleaseBuffer(m, v35)
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L1
	} else {
		goto L90
	}
L43:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v57+v143<<(uint(int32(2))%32))+168))
	v152 = v143 + base.B2i32(v136 == v149)
	goto L45
L44:
	;
	v152 = v143
	goto L45
L45:
	;
	v154 = v152 & int32(65535)
	v155 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v158 = F_GenericXLogRegisterBuffer(m, v155, v35, int32(0))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v160 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+30)))
	if base.Ui32(v154) < base.Ui32(v160) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v163 = v154
	v164 = v155
	v168 = v158
	goto L51
L49:
	;
	v249 = v155
	v253 = v158
	goto L50
L50:
	;
	v258 = F_BloomNewBuffer(m, l0)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L1
	} else {
		goto L76
	}
L51:
	;
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v168+v163<<(uint(int32(2))%32))+168))
	v177 = F_ReadBuffer(m, l0, v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L53
	}
L52:
	;
	v249 = v240
	v253 = v243
	goto L50
L53:
	;
	F_LockBuffer(m, v177, int32(2))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	v183 = F_GenericXLogRegisterBuffer(m, v164, v177, int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+16)))
	v209 = v183 + v208
	v210 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209))))
	v211 = v207 * v210
	v212 = int32(8160) - v211
	if base.Ui32(v207) <= base.Ui32(v212) {
		goto L63
	} else {
		goto L64
	}
L56:
	;
	v185 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+14)))
	if v185 != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+16)))
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183+v186)+2)))
	if v188&int32(2) == int32(0) {
		goto L55
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	v193 = int32(0)
	F_PageInit(m, v183, int32(8192), int32(8))
	mBase = m.M
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+16)))
	v198 = v183 + v197
	v199 = int32(65411)
	*(*uint16)(unsafe.Add(mBase, uint32(v198)+6)) = uint16(v199)
	*(*uint16)(unsafe.Add(mBase, uint32(v198)+2)) = uint16(v193)
	goto L61
L60:
	;
	goto L59
L61:
	;
	goto L55
L62:
	;
	if base.Ui32(v207) <= base.Ui32(v212) {
		goto L66
	} else {
		goto L67
	}
L63:
	;
	v215 = int32(24)
	v217 = F___memcpy(m, v183+v211+v215, v32, v207)
	mBase = m.M
	v218 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v209))))
	v220 = v218 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v209))) = uint16(v220)
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v225 = v222*v220 + v215
	*(*uint16)(unsafe.Add(mBase, uint32(v183)+12)) = uint16(v225)
	goto L65
L64:
	;
	goto L65
L65:
	;
	goto L62
L66:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v168)+28)) = uint16(v163)
	F_GenericXLogFinish(m, v164)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_pfree(m, v164)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	F_UnlockReleaseBuffer(m, v177)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	goto L42
L71:
	;
	F_UnlockReleaseBuffer(m, v177)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v239 = v163 + int32(1)
	v240 = F_GenericXLogStart(m, l0)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v243 = F_GenericXLogRegisterBuffer(m, v240, v35, int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v243)+30)))
	if base.Ui32(v239) < base.Ui32(v245) {
		v163 = v239
		v164 = v240
		v168 = v243
		goto L51
	} else {
		goto L75
	}
L75:
	;
	goto L52
L76:
	;
	v261 = F_GenericXLogRegisterBuffer(m, v249, v258, int32(1))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v263 = int32(0)
	F_PageInit(m, v261, int32(8192), int32(8))
	mBase = m.M
	v267 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+16)))
	v268 = v261 + v267
	v269 = int32(65411)
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+6)) = uint16(v269)
	*(*uint16)(unsafe.Add(mBase, uint32(v268)+2)) = uint16(v263)
	goto L78
L78:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v261)+16)))
	v279 = v261 + v278
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279))))
	v281 = v277 * v280
	v282 = int32(8160) - v281
	if base.Ui32(v277) <= base.Ui32(v282) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if base.B2i32(base.Ui32(v277) <= base.Ui32(v282)) == int32(0) {
		goto L7
	} else {
		goto L83
	}
L80:
	;
	v285 = int32(24)
	v287 = F___memcpy(m, v261+v281+v285, v32, v277)
	mBase = m.M
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v279))))
	v290 = v288 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v279))) = uint16(v290)
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v14)+1164))
	v295 = v292*v290 + v285
	*(*uint16)(unsafe.Add(mBase, uint32(v261)+12)) = uint16(v295)
	goto L82
L81:
	;
	goto L82
L82:
	;
	goto L79
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+28)) = int32(65536)
	if v258 < int32(0) {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+168)) = v321
	F_GenericXLogFinish(m, v249)
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L88
	}
L85:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v306+(v258^int32(-1))<<(uint(int32(6))%32))+16))
	v321 = v312
	goto L84
L86:
	;
	goto L87
L87:
	;
	v314 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v314+v258<<(uint(int32(6))%32)+int32(-64))+16))
	v321 = v320
	goto L84
L88:
	;
	F_UnlockReleaseBuffer(m, v258)
	mBase = m.M
	v326 = m.ExcPending
	if v326 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	goto L42
L90:
	;
	goto L8
L91:
	;
	m.G0 = v14 + int32(1168)
	return int32(0)
L92:
	;
	F_errmsg_internal(m, int32(426720), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(516660), int32(324), int32(87799))
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L94
	}
L94:
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
	v52 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v52+(v44^int32(-1))<<(uint(int32(2))%32))))
	v66 = v58
	goto L20
L22:
	;
	goto L23
L23:
	;
	v60 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 != 0 {
		v8 = int32(1)
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = base.B2i32(v5 != int32(0))
	}
	return v8
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
		F_errmsg_internal(m, int32(499319), v6)
		mBase = m.M
		v23 = m.ExcPending
		if v23 != 0 {
			return
		} else {
			F_errfinish(m, int32(329894), int32(137), int32(221302))
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
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
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	v22 = int32(1)
	v23 = v21 & v22
	if v21 == v22 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v23 != 0 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v28&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v41 = int32(1)
	if v23 != 0 {
		v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v37 = v26
	goto L10
L9:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L10
L10:
	;
	if v28 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v40 = v26
	goto L13
L12:
	;
	v40 = v37
	goto L13
L13:
	;
	v51 = v40
	goto L4
L14:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v51 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v54 = v17
	goto L17
L16:
	;
	v54 = v12 + int32(4)
	goto L17
L17:
	;
	v59 = v51
	goto L18
L18:
	;
	if v59 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v78 = int32(1)
	v79 = v19 + v78
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v82 = v80 & v78
	if v80 == v78 {
		goto L26
	} else {
		goto L27
	}
L20:
	;
	goto L19
L21:
	;
	v77 = v51 >> (uint(int32(31)) % 32) & v51
	goto L20
L22:
	;
	goto L23
L23:
	;
	v71 = v59 - int32(1)
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v71))))
	if v73 == int32(32) {
		v59 = v71
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v77 = v59
	goto L20
L25:
	;
	if v82 != 0 {
		goto L36
	} else {
		goto L37
	}
L26:
	;
	v85 = int32(4)
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79))))
	if v87&int32(254) == int32(2) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	v100 = int32(1)
	if v82 != 0 {
		v110 = int32(base.Ui32(v80)>>(uint(v100)%32)) - v100
		goto L25
	} else {
		goto L35
	}
L29:
	;
	v96 = v85
	goto L31
L30:
	;
	v96 = base.B2i32(v87 == int32(18)) << (uint(v85) % 32)
	goto L31
L31:
	;
	if v87 == int32(1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v99 = v85
	goto L34
L33:
	;
	v99 = v96
	goto L34
L34:
	;
	v110 = v99
	goto L25
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v110 = int32(base.Ui32(v104)>>(uint(int32(2))%32)) - int32(4)
	goto L25
L36:
	;
	v113 = v79
	goto L38
L37:
	;
	v113 = v19 + int32(4)
	goto L38
L38:
	;
	v118 = v110
	goto L39
L39:
	;
	if v118 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v137 = int32(1)
	if v21&v137 != 0 {
		goto L46
	} else {
		goto L47
	}
L41:
	;
	goto L40
L42:
	;
	v136 = v110 >> (uint(int32(31)) % 32) & v110
	goto L41
L43:
	;
	goto L44
L44:
	;
	v130 = v118 - int32(1)
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113+v130))))
	if v132 == int32(32) {
		v118 = v130
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v136 = v118
	goto L41
L46:
	;
	v141 = v137
	goto L48
L47:
	;
	v141 = int32(4)
	goto L48
L48:
	;
	v143 = int32(1)
	if v80&v143 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v147 = v143
	goto L51
L50:
	;
	v147 = int32(4)
	goto L51
L51:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v150 = F_varstr_cmp(m, v12+v141, v77, v19+v147, v136, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v152 != v12 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	F_pfree(m, v12)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v156 != v19 {
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
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	return base.B2i32(v150 <= int32(0))
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
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
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
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
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
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
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
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
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v23 = v21 & v19
	if v21 == v19 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L1
	} else {
		goto L94
	}
L7:
	;
	if v23 != 0 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v26 = int32(4)
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v28&int32(254) == int32(2) {
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
	if v23 != 0 {
		v51 = int32(base.Ui32(v21)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v37 = v26
	goto L13
L12:
	;
	v37 = base.B2i32(v28 == int32(18)) << (uint(v26) % 32)
	goto L13
L13:
	;
	if v28 == int32(1) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = v26
	goto L16
L15:
	;
	v40 = v37
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
	v54 = v20
	goto L20
L19:
	;
	v54 = v11 + int32(4)
	goto L20
L20:
	;
	v59 = v51
	goto L21
L21:
	;
	if v59 <= int32(0) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	v77 = int32(1)
	v78 = v16 + v77
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v81 = v79 & v77
	if v79 == v77 {
		goto L29
	} else {
		goto L30
	}
L23:
	;
	goto L22
L24:
	;
	v76 = v51 >> (uint(int32(31)) % 32) & v51
	goto L23
L25:
	;
	goto L26
L26:
	;
	v70 = v59 - int32(1)
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54+v70))))
	if v72 == int32(32) {
		v59 = v70
		goto L21
	} else {
		goto L27
	}
L27:
	;
	v76 = v59
	goto L23
L28:
	;
	if v81 != 0 {
		goto L39
	} else {
		goto L40
	}
L29:
	;
	v84 = int32(4)
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v78))))
	if v86&int32(254) == int32(2) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v99 = int32(1)
	if v81 != 0 {
		v109 = int32(base.Ui32(v79)>>(uint(v99)%32)) - v99
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v95 = v84
	goto L34
L33:
	;
	v95 = base.B2i32(v86 == int32(18)) << (uint(v84) % 32)
	goto L34
L34:
	;
	if v86 == int32(1) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v98 = v84
	goto L37
L36:
	;
	v98 = v95
	goto L37
L37:
	;
	v109 = v98
	goto L28
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v109 = int32(base.Ui32(v103)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	v112 = v78
	goto L41
L40:
	;
	v112 = v16 + int32(4)
	goto L41
L41:
	;
	v117 = v109
	goto L42
L42:
	;
	if v117 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v136 = F_pg_newlocale_from_collation(m, v18)
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L50
	}
L44:
	;
	goto L43
L45:
	;
	v134 = v109 >> (uint(int32(31)) % 32) & v109
	goto L44
L46:
	;
	goto L47
L47:
	;
	v128 = v117 - int32(1)
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v112+v128))))
	if v130 == int32(32) {
		v117 = v128
		goto L42
	} else {
		goto L48
	}
L48:
	;
	v134 = v117
	goto L44
L49:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v239 != v11 {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+1)))
	if v138 == int32(1) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v134 != v76 {
		v238 = int32(1)
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v220 = int32(1)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v222&v220 != 0 {
		goto L79
	} else {
		goto L80
	}
L54:
	;
	v142 = int32(1)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	if v144&v142 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v147 = v142
	goto L57
L56:
	;
	v147 = int32(4)
	goto L57
L57:
	;
	v148 = v11 + v147
	v149 = int32(1)
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v151&v149 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v154 = v149
	goto L60
L59:
	;
	v154 = int32(4)
	goto L60
L60:
	;
	v155 = v16 + v154
	if base.Ui32(int32(4)) <= base.Ui32(v76) {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v238 = base.B2i32(v217 != int32(0))
	goto L49
L62:
	;
	v217 = int32(0)
	goto L61
L63:
	;
	v191 = v186
	v192 = v187
	v193 = v188
	goto L73
L64:
	;
	if (v148|v155)&int32(3) != 0 {
		v186 = v148
		v187 = v155
		v188 = v76
		goto L63
	} else {
		goto L67
	}
L65:
	;
	v179 = v148
	v180 = v155
	v181 = v76
	goto L66
L66:
	;
	if v181 == int32(0) {
		goto L62
	} else {
		goto L72
	}
L67:
	;
	v163 = v148
	v164 = v155
	v165 = v76
	goto L68
L68:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
	if v168 != v169 {
		v186 = v163
		v187 = v164
		v188 = v165
		goto L63
	} else {
		goto L70
	}
L69:
	;
	v179 = v174
	v180 = v172
	v181 = v176
	goto L66
L70:
	;
	v171 = int32(4)
	v172 = v164 + v171
	v174 = v163 + v171
	v176 = v165 - v171
	if base.Ui32(int32(3)) < base.Ui32(v176) {
		v163 = v174
		v164 = v172
		v165 = v176
		goto L68
	} else {
		goto L71
	}
L71:
	;
	goto L69
L72:
	;
	v186 = v179
	v187 = v180
	v188 = v181
	goto L63
L73:
	;
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v191))))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192))))
	if v196 == v197 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v217 = v196 - v197
	goto L61
L75:
	;
	v199 = int32(1)
	v204 = v193 - v199
	if v204 != 0 {
		v191 = v191 + v199
		v192 = v192 + v199
		v193 = v204
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
	v225 = v220
	goto L81
L80:
	;
	v225 = int32(4)
	goto L81
L81:
	;
	v227 = int32(1)
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if v229&v227 != 0 {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v232 = v227
	goto L84
L83:
	;
	v232 = int32(4)
	goto L84
L84:
	;
	v234 = F_varstr_cmp(m, v11+v225, v76, v16+v232, v134, v18)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v238 = base.B2i32(v234 != int32(0))
	goto L49
L86:
	;
	F_pfree(m, v11)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v243 != v16 {
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
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	return v238
L93:
	;
	goto L92
L94:
	;
	F_errcode(m, int32(34209924))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L95
	}
L95:
	;
	F_errmsg(m, int32(258179), int32(0))
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	F_errhint(m, int32(602565), int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	F_errfinish(m, int32(520042), int32(738), int32(113268))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
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
	var v31 int32
	_ = v31
	var v34 int64
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
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int64
	_ = v150
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int64
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
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
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
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
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
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
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v480 int32
	_ = v480
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v495 float64
	_ = v495
	var v497 float64
	_ = v497
	var v502 int32
	_ = v502
	var v503 float64
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
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
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v558 int32
	_ = v558
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
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
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v646 int32
	_ = v646
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
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
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
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
	var v698 int32
	_ = v698
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v726 float64
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
	var v737 int32
	_ = v737
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 float64
	_ = v744
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v768 float64
	_ = v768
	var v771 float64
	_ = v771
	var v772 int32
	_ = v772
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
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
	v771 = *(*float64)(unsafe.Add(mBase, uint32(v148)+8))
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v148)+40))
	F_brinRevmapTerminate(m, v772)
	mBase = m.M
	v774 = m.ExcPending
	if v774 != 0 {
		goto L3
	} else {
		goto L176
	}
L2:
	;
	v717 = int32(0)
	v724 = *(*int32)(unsafe.Add(mBase, uint32(l0)+188))
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v724)+140))
	v726 = m.T0[v725].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, l0, l1, l2, v717, v717, int32(1), v717, int32(-1), int32(14), v148, v717)
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L3
	} else {
		goto L171
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
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v31
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = l1
	v34 = *(*int64)(unsafe.Add(mBase, uint32(v20)+28))
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v34
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
	v702 = m.ExcPending
	if v702 != 0 {
		goto L3
	} else {
		goto L168
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
	v44 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(v39^int32(-1))<<(uint(int32(2))%32))))
	v58 = v50
	goto L8
L11:
	;
	goto L12
L12:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	F_PageInit(m, v58, int32(8192), int32(8))
	mBase = m.M
	v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
	v69 = int32(61585)
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
	v86 = *(*int32)(unsafe.Add(mBase, _consts[2]))
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
	v116 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v116+(v39^int32(-1))<<(uint(int32(2))%32))))
	v130 = v122
	goto L32
L34:
	;
	goto L35
L35:
	;
	v124 = *(*int32)(unsafe.Add(mBase, _consts[1]))
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
	v143 = int32(0)
	v145 = F_RelationGetNumberOfBlocksInFork(m, l0, v143)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L3
	} else {
		goto L38
	}
L38:
	;
	v148 = F_palloc(m, int32(80))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L3
	} else {
		goto L39
	}
L39:
	;
	v150 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v148)+8)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v148)+40)) = v140
	v154 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v148)+32)) = v154
	*(*int32)(unsafe.Add(mBase, uint32(v148)+28)) = v142
	*(*int64)(unsafe.Add(mBase, uint32(v148)+16)) = v150
	*(*int32)(unsafe.Add(mBase, uint32(v148)+24)) = v154
	v161 = F_brin_build_desc(m, l1)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L3
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+44)) = v161
	v164 = F_brin_new_memtuple(m, v161)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L3
	} else {
		goto L41
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+72)) = int32(0)
	v168 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v148)+64)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v148)+48)) = v164
	v172 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int64)(unsafe.Add(mBase, uint32(v148)+52)) = v168
	*(*int32)(unsafe.Add(mBase, uint32(v148)+60)) = v172
	if v145 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v177 = v145 - int32(1)
	v178 = base.I32_rem_u_s(v177, v142)
	v181 = v177 - v178
	goto L44
L43:
	;
	v181 = v143
	goto L44
L44:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v148)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+36)) = v181 + v182
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+128))
	if v185 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v148)+64))
	if v437 == int32(0) {
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
	v196 = *(*int32)(unsafe.Add(mBase, _consts[4]))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v196)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v196)+72)) = v197 + int32(1)
	goto L48
L48:
	;
	v204 = F_CreateParallelContext(m, int32(171754), int32(291620), v185)
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
	v212 = int32(4216240)
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
	v277 = *(*int32)(unsafe.Add(mBase, _consts[5]))
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
	v307 = *(*int32)(unsafe.Add(mBase, _consts[4]))
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
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v148)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v313)+12)) = v321
	v325 = *(*int32)(unsafe.Add(mBase, _consts[6]))
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
	*(*int64)(unsafe.Add(mBase, uint32(v313-int32(-64)))) = v338
	F_table_parallelscan_initialize(m, l0, v313+int32(96), v212)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L3
	} else {
		goto L87
	}
L87:
	;
	v350 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v351 = F_shm_toc_allocate(m, v350, v226)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L3
	} else {
		goto L88
	}
L88:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v204)+44))
	F_tuplesort_initialize_shared(m, v351, v189, v353)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L3
	} else {
		goto L89
	}
L89:
	;
	v356 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v356, int64(-5764607523034234879), v313)
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
	F_shm_toc_insert(m, v360, int64(-5764607523034234878), v351)
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L3
	} else {
		goto L91
	}
L91:
	;
	v365 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v365 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v367 = F_shm_toc_allocate(m, v366, v294)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L3
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v381 = F_mul_size(m, int32(32), v380)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L3
	} else {
		goto L101
	}
L95:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	if v294 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v373, int64(-5764607523034234877), v372)
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L3
	} else {
		goto L100
	}
L97:
	;
	v371 = F__emscripten_memcpy_bulkmem(m, v367, v370, v294)
	mBase = m.M
	v372 = v371
	goto L99
L98:
	;
	v372 = v367
	goto L99
L99:
	;
	goto L96
L100:
	;
	goto L94
L101:
	;
	v383 = F_shm_toc_allocate(m, v378, v381)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L3
	} else {
		goto L102
	}
L102:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v385, int64(-5764607523034234876), v383)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L3
	} else {
		goto L103
	}
L103:
	;
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v204)+12))
	v392 = F_mul_size(m, int32(128), v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L3
	} else {
		goto L104
	}
L104:
	;
	v394 = F_shm_toc_allocate(m, v389, v392)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L3
	} else {
		goto L105
	}
L105:
	;
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v204)+52))
	F_shm_toc_insert(m, v396, int64(-5764607523034234875), v394)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L3
	} else {
		goto L106
	}
L106:
	;
	F_LaunchParallelWorkers(m, v204)
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L3
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v192))) = v204
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v192)+24)) = v394
	*(*int32)(unsafe.Add(mBase, uint32(v192)+20)) = v383
	*(*int32)(unsafe.Add(mBase, uint32(v192)+16)) = v212
	*(*int32)(unsafe.Add(mBase, uint32(v192)+12)) = v351
	*(*int32)(unsafe.Add(mBase, uint32(v192)+8)) = v313
	*(*int32)(unsafe.Add(mBase, uint32(v192)+4)) = v403 + int32(1)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v204)+20))
	if v412 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L108:
	;
	F__brin_end_parallel(m, v192)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L3
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+64)) = v192
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v192)+8))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v192)+12))
	v421 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v192)+4))
	v423 = base.I32_div_s(v421, v422)
	F__brin_parallel_scan_and_build(m, v148, v418, v419, l0, l1, v423)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
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
	v427 = m.ExcPending
	if v427 != 0 {
		goto L3
	} else {
		goto L113
	}
L113:
	;
	goto L45
L114:
	;
	v441 = F_palloc0(m, int32(12))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L3
	} else {
		goto L115
	}
L115:
	;
	v443 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v441))) = uint8(v443)
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v148)+64))
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v445)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v441)+4)) = v446
	v448 = *(*int32)(unsafe.Add(mBase, uint32(v148)+64))
	v449 = *(*int32)(unsafe.Add(mBase, uint32(v448)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v441)+8)) = v449
	v452 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v453 = F_tuplesort_begin_index_brin(m, v452, v441)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L3
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+72)) = v453
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v148)+64))
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v456)+8))
	v461 = v457 + int32(44)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	goto L117
L117:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v461)))
	*(*int32)(unsafe.Add(mBase, uint32(v461))) = int32(1)
	if v480 != 0 {
		goto L119
	} else {
		goto L120
	}
L118:
	;
	v495 = *(*float64)(unsafe.Add(mBase, uint32(v457)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v148)+16)) = v495
	v497 = *(*float64)(unsafe.Add(mBase, uint32(v457)+64))
	*(*float64)(unsafe.Add(mBase, uint32(v148)+8)) = v497
	*(*int32)(unsafe.Add(mBase, uint32(v457)+44)) = int32(0)
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L3
	} else {
		goto L127
	}
L119:
	;
	F_s_lock(m, v461, int32(521025), int32(2587), int32(298000))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L3
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v457)+48))
	if v462 != v488 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L121
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v461))) = int32(0)
	F_ConditionVariableSleep(m, v457+int32(32), int32(134217767))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
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
	v503 = *(*float64)(unsafe.Add(mBase, uint32(v148)+16))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	F_tuplesort_performsort(m, v504)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L3
	} else {
		goto L128
	}
L128:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v148)+44))
	v508 = F_brin_new_memtuple(m, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L3
	} else {
		goto L129
	}
L129:
	;
	v511 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v516 = F_AllocSetContextCreateInternal(m, v511, int32(286277), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v517 = m.ExcPending
	if v517 != 0 {
		goto L3
	} else {
		goto L130
	}
L130:
	;
	v518 = int32(4562080)
	v519 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v516
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	v525 = F_tuplesort_getbrintuple(m, v522, v20+int32(20))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L3
	} else {
		goto L132
	}
L131:
	;
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v148)+36))
	F_brin_fill_empty_ranges(m, v148, v688, v689)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L3
	} else {
		goto L165
	}
L132:
	;
	if v525 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	F_tuplesort_end(m, v529)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L3
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v534 = v148 + int32(24)
	v537 = v508
	v539 = v525
	v543 = int32(-1)
	goto L138
L136:
	;
	v688 = int32(-1)
	goto L131
L137:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v148)+44))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v640)+4))
	v660 = F_brin_form_tuple(m, v656, v657, v640, v20+int32(44))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L3
	} else {
		goto L162
	}
L138:
	;
	if v543 != int32(-1) {
		goto L141
	} else {
		goto L142
	}
L139:
	;
	v633 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	F_tuplesort_end(m, v633)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L3
	} else {
		goto L160
	}
L140:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	F_brin_fill_empty_ranges(m, v148, v543, v624)
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L3
	} else {
		goto L157
	}
L141:
	;
	v558 = v539
	goto L145
L142:
	;
	goto L143
L143:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v148)+44))
	v604 = F_brin_deform_tuple(m, v603, v539, v537)
	mBase = m.M
	v605 = m.ExcPending
	if v605 != 0 {
		goto L3
	} else {
		goto L156
	}
L144:
	;
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v148)+44))
	v589 = F_brin_form_tuple(m, v586, v572, v537, v20+int32(44))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L3
	} else {
		goto L152
	}
L145:
	;
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v537)+4))
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v558)))
	if v572 != v573 {
		goto L144
	} else {
		goto L147
	}
L146:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	F_tuplesort_end(m, v583)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L3
	} else {
		goto L151
	}
L147:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v148)+44))
	F_union_tuples(m, v575, v537, v558)
	mBase = m.M
	v577 = m.ExcPending
	if v577 != 0 {
		goto L3
	} else {
		goto L148
	}
L148:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	v581 = F_tuplesort_getbrintuple(m, v578, v20+int32(20))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L3
	} else {
		goto L149
	}
L149:
	;
	if v581 != 0 {
		v558 = v581
		goto L145
	} else {
		goto L150
	}
L150:
	;
	goto L146
L151:
	;
	v640 = v537
	v646 = v543
	goto L137
L152:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v148)+28))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v148)+40))
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v589)))
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v596 = F_brin_doinsert(m, v591, v592, v593, v534, v594, v589, v595)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L3
	} else {
		goto L153
	}
L153:
	;
	F_MemoryContextReset(m, v516)
	mBase = m.M
	v599 = m.ExcPending
	if v599 != 0 {
		goto L3
	} else {
		goto L154
	}
L154:
	;
	v600 = *(*int32)(unsafe.Add(mBase, uint32(v148)+44))
	v601 = F_brin_deform_tuple(m, v600, v558, v537)
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L3
	} else {
		goto L155
	}
L155:
	;
	v609 = v558
	v623 = v601
	goto L140
L156:
	;
	v609 = v539
	v623 = v604
	goto L140
L157:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v609)))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v148)+72))
	v631 = F_tuplesort_getbrintuple(m, v628, v20+int32(20))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L3
	} else {
		goto L158
	}
L158:
	;
	if v631 != 0 {
		v537 = v623
		v539 = v631
		v543 = v627
		goto L138
	} else {
		goto L159
	}
L159:
	;
	goto L139
L160:
	;
	v636 = int32(-1)
	if v627 == v636 {
		v688 = v636
		goto L131
	} else {
		goto L161
	}
L161:
	;
	v640 = v623
	v646 = v627
	goto L137
L162:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v148)+28))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v148)+40))
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v660)))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v20)+44))
	v667 = F_brin_doinsert(m, v662, v663, v664, v534, v665, v660, v666)
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L3
	} else {
		goto L163
	}
L163:
	;
	F_pfree(m, v660)
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L3
	} else {
		goto L164
	}
L164:
	;
	v688 = v646
	goto L131
L165:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v519
	F_MemoryContextDelete(m, v516)
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L3
	} else {
		goto L166
	}
L166:
	;
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v148)+64))
	F__brin_end_parallel(m, v696)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L3
	} else {
		goto L167
	}
L167:
	;
	v768 = v503
	goto L1
L168:
	;
	v703 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v703 + int32(4)
	F_errmsg_internal(m, int32(530312), v20+int32(16))
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L3
	} else {
		goto L169
	}
L169:
	;
	F_errfinish(m, int32(521025), int32(1120), int32(451353))
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L3
	} else {
		goto L170
	}
L170:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L171:
	;
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v148)+44))
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
	v730 = *(*int32)(unsafe.Add(mBase, uint32(v148)+48))
	v733 = F_brin_form_tuple(m, v728, v729, v730, v20+int32(20))
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L3
	} else {
		goto L172
	}
L172:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v148)+28))
	v737 = *(*int32)(unsafe.Add(mBase, uint32(v148)+40))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v20)+20))
	v742 = F_brin_doinsert(m, v735, v736, v737, v148+int32(24), v740, v733, v741)
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L3
	} else {
		goto L173
	}
L173:
	;
	v744 = *(*float64)(unsafe.Add(mBase, uint32(v148)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v148)+8)) = base.F64_add(v744, float64(1))
	F_pfree(m, v733)
	mBase = m.M
	v749 = m.ExcPending
	if v749 != 0 {
		goto L3
	} else {
		goto L174
	}
L174:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v148)+32))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v148)+36))
	F_brin_fill_empty_ranges(m, v148, v750, v751)
	mBase = m.M
	v753 = m.ExcPending
	if v753 != 0 {
		goto L3
	} else {
		goto L175
	}
L175:
	;
	v768 = v726
	goto L1
L176:
	;
	F_terminate_brin_buildstate(m, v148)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L3
	} else {
		goto L177
	}
L177:
	;
	v778 = F_palloc(m, int32(16))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L3
	} else {
		goto L178
	}
L178:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v778)+8)) = v771
	*(*float64)(unsafe.Add(mBase, uint32(v778))) = v768
	m.G0 = v20 + int32(48)
	return v778
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
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v36+(v7+v24))))
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = int32(1)
		v11 = v6 + v10
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		v16 = v14 & v10
		if v16 != 0 {
			v17 = v11
		} else {
			v17 = v6 + int32(4)
		}
		if v14 == int32(1) {
			v20 = int32(4)
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
			if v22&int32(254) == int32(2) {
				v31 = v20
			} else {
				v31 = base.B2i32(v22 == int32(18)) << (uint(v20) % 32)
			}
			if v22 == int32(1) {
				v34 = v20
			} else {
				v34 = v31
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v16 != 0 {
				v45 = int32(base.Ui32(v14)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v47 = int32(1)
		v50 = F_dotrim(m, v17, v45, int32(780780), v47, v47, v47)
		mBase = m.M
		v51 = m.ExcPending
		if v51 != 0 {
			return int32(0)
		} else {
			return v50
		}
	}
}
func F_bttext_pattern_cmp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum_packed(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
			if v20 == int32(1) {
				v23 = int32(4)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+1)))
				if v25&int32(254) == int32(2) {
					v34 = v23
				} else {
					v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
				}
				if v25 == int32(1) {
					v37 = v23
				} else {
					v37 = v34
				}
				v50 = v37
			} else {
				v38 = int32(1)
				if v20&v38 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v38)%32)) - v38
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v51 == int32(1) {
				v54 = int32(4)
				v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v56&int32(254) == int32(2) {
					v65 = v54
				} else {
					v65 = base.B2i32(v56 == int32(18)) << (uint(v54) % 32)
				}
				if v56 == int32(1) {
					v68 = v54
				} else {
					v68 = v65
				}
				v81 = v68
			} else {
				v69 = int32(1)
				if v51&v69 != 0 {
					v81 = int32(base.Ui32(v51)>>(uint(v69)%32)) - v69
				} else {
					v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v81 = int32(base.Ui32(v75)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v82 = int32(1)
			if v20&v82 != 0 {
				v86 = v82
			} else {
				v86 = int32(4)
			}
			v88 = int32(1)
			if v51&v88 != 0 {
				v92 = v88
			} else {
				v92 = int32(4)
			}
			v94 = base.B2i32(v50 < v81)
			if v50 < v81 {
				v95 = v50
			} else {
				v95 = v81
			}
			v96 = F_memcmp(m, v7+v86, v14+v92, v95)
			mBase = m.M
			if v96 != 0 {
				v99 = v96
			} else {
				if v50 < v81 {
					v99 = int32(-1)
				} else {
					v99 = base.B2i32(v81 < v50)
				}
			}
			v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v100 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v104 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							return v99
						}
					} else {
						return v99
					}
				}
			} else {
				v104 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v104 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v107 = m.ExcPending
					if v107 != 0 {
						return int32(0)
					} else {
						return v99
					}
				} else {
					return v99
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
	v6 = (l0 - v3) & int32(65535)
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
	var v50 int32
	_ = v50
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
	var v83 int32
	_ = v83
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
	v50 = v34
	goto L7
L5:
	;
	v83 = v34
	goto L6
L6:
	;
	v86 = F_makeFuncExpr(m, l5, l3, v83, l4, int32(0))
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
	v83 = v69
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
	v69 = F_lappend(m, v50, v57)
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
		v50 = v69
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
	var v25 int32
	_ = v25
	var v27 int64
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
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
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
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
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
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
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
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
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v257 int32
	_ = v257
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
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int64
	_ = v284
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v301 float64
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v326 float64
	_ = v326
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 float64
	_ = v331
	var v332 float64
	_ = v332
	var v335 int32
	_ = v335
	var v336 float64
	_ = v336
	var v337 float64
	_ = v337
	var v339 float64
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v360 float64
	_ = v360
	var v366 int32
	_ = v366
	var v367 float64
	_ = v367
	var v368 float64
	_ = v368
	var v371 int32
	_ = v371
	var v379 int32
	_ = v379
	var v384 float64
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v396 int32
	_ = v396
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 float64
	_ = v409
	var v412 float64
	_ = v412
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v440 int32
	_ = v440
	var v446 int32
	_ = v446
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 float64
	_ = v459
	var v462 float64
	_ = v462
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v491 float64
	_ = v491
	var v492 float64
	_ = v492
	var v497 float64
	_ = v497
	var v498 int32
	_ = v498
	var v505 int32
	_ = v505
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v590 int32
	_ = v590
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v604 int32
	_ = v604
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 float64
	_ = v616
	var v617 float64
	_ = v617
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
	goto L4
L3:
	;
	v27 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+328)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v25)+72)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+20)) = v27
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = l0
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v34 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = F_copyObjectImpl(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L7
	}
L4:
	;
	v25 = F__emscripten_memcpy_bulkmem(m, v20, l0, int32(384))
	mBase = m.M
	goto L6
L6:
	;
	goto L3
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+4)) = v39
	v42 = int32(1)
	F_IncrementVarSublevelsUp(m, v39, v42, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v47 = F_copyObjectImpl(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+128)) = v47
	v50 = int32(1)
	F_IncrementVarSublevelsUp(m, v47, v50, v50)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v55 = F_copyObjectImpl(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v59 = F_pstrdup(m, int32(115427))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v62 = F_makeTargetEntry(m, v55, int32(1), v59, int32(0))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v62
	v69 = F_list_make1_impl(m, int32(1), v17+int32(4))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+76)) = v69
	*(*int32)(unsafe.Add(mBase, uint32(v25)+264)) = v69
	v73 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+112)) = v73
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+318)) = uint8(v73)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+40)) = uint8(v73)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+120)) = v73
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+36)) = uint8(v73)
	v84 = F_palloc0(m, int32(20))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+8)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v84))) = int32(52)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v91 = F_copyObjectImpl(m, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v84)+16)) = int32(-1)
	v95 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v84)+12)) = uint8(v95)
	*(*int32)(unsafe.Add(mBase, uint32(v84)+4)) = v91
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v39)+60))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+8))
	v100 = F_list_member(m, v99, v84)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	if v100 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v39)+60))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+8))
	v106 = F_lcons(m, v84, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v112 = F_palloc0(m, int32(20))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L22
	}
L21:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v39)+60))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+8)) = v106
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v112))) = int32(106)
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v25)+264))
	v117 = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v62)+16))
	if v126 == v117 {
		goto L24
	} else {
		goto L25
	}
L23:
	;
	v257 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+18)) = uint8(v257)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+17)) = uint8(v6)
	*(*uint8)(unsafe.Add(mBase, uint32(v112)+16)) = uint8(v5)
	*(*int32)(unsafe.Add(mBase, uint32(v112)+12)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v112)+4)) = v248
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v112
	*(*int32)(unsafe.Add(mBase, uint32(v17)+8)) = v112
	v267 = F_list_make1_impl(m, int32(1), v17)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L60
	}
L24:
	;
	if v116 == int32(0) {
		v244 = int32(1)
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v248 = v126
	goto L26
L26:
	;
	goto L23
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v62)+16)) = v244
	v248 = v244
	goto L26
L28:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if v133 <= int32(0) {
		v244 = int32(1)
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v136 = int32(0)
	if v136 < v133 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v139 = v133
	goto L32
L31:
	;
	v139 = v136
	goto L32
L32:
	;
	v141 = v139 & int32(3)
	v142 = int32(0)
	if int32(4) <= v133 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v151 = v142
	v152 = v117
	v153 = int32(0)
	goto L36
L34:
	;
	v186 = v142
	v187 = v117
	goto L35
L35:
	;
	if v141 != 0 {
		goto L51
	} else {
		goto L52
	}
L36:
	;
	v162 = v147 + v152<<(uint(int32(2))%32)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v162)+12))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)+16))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v162)+8))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v162)+4))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v169)+16))
	if base.Ui32(v151) < base.Ui32(v170) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	v186 = v178
	v187 = v180
	goto L35
L38:
	;
	v172 = v170
	goto L40
L39:
	;
	v172 = v151
	goto L40
L40:
	;
	if base.Ui32(v172) < base.Ui32(v168) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v174 = v168
	goto L43
L42:
	;
	v174 = v172
	goto L43
L43:
	;
	if base.Ui32(v174) < base.Ui32(v166) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v176 = v166
	goto L46
L45:
	;
	v176 = v174
	goto L46
L46:
	;
	if base.Ui32(v176) < base.Ui32(v164) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v178 = v164
	goto L49
L48:
	;
	v178 = v176
	goto L49
L49:
	;
	v179 = int32(4)
	v180 = v152 + v179
	v182 = v153 + v179
	if v182 != v139&int32(2147483644) {
		v151 = v178
		v152 = v180
		v153 = v182
		goto L36
	} else {
		goto L50
	}
L50:
	;
	goto L37
L51:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	v197 = int32(0)
	v199 = v186
	v200 = v187
	goto L54
L52:
	;
	v222 = v186
	goto L53
L53:
	;
	v244 = v222 + int32(1)
	goto L27
L54:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v195+v200<<(uint(int32(2))%32))))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v211)+16))
	if base.Ui32(v199) < base.Ui32(v212) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v222 = v214
	goto L53
L56:
	;
	v214 = v212
	goto L58
L57:
	;
	v214 = v199
	goto L58
L58:
	;
	v215 = int32(1)
	v218 = v197 + v215
	if v218 != v141 {
		v197 = v218
		v199 = v214
		v200 = v200 + v215
		goto L54
	} else {
		goto L59
	}
L59:
	;
	goto L55
L60:
	;
	v269 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+128)) = v269
	*(*int32)(unsafe.Add(mBase, uint32(v39)+124)) = v267
	v277 = F_Int64GetDatum(m, int64(1))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v279 = int32(0)
	v281 = F_makeConst(m, int32(20), int32(-1), v269, int32(8), v277, v279, v279)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+132)) = v281
	v284 = int64(4607182418800017408)
	*(*int64)(unsafe.Add(mBase, uint32(v25)+304)) = v284
	*(*int64)(unsafe.Add(mBase, uint32(v25)+296)) = v284
	v290 = F_query_planner(m, v25, int32(831), int32(0))
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	F_SS_identify_outer_params(m, v25)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v294 = int32(0)
	v301 = float64(0)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v25)+72))
	if v302 == v294 {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v290)+32))
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v25)+156))
	v491 = float64(1)
	v492 = *(*float64)(unsafe.Add(mBase, uint32(v290)+16))
	if base.F64_gt(v492, v491) != 0 {
		goto L97
	} else {
		goto L98
	}
L66:
	;
	goto L65
L67:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	if v305 <= int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v290)+32))
	if v385 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L69:
	;
	v379 = v294
	v384 = v301
	goto L68
L70:
	;
	goto L71
L71:
	;
	v308 = int32(1)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v302)+12))
	if v305 == v308 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	if v305&v308 == int32(0) {
		v379 = v355
		v384 = v360
		goto L68
	} else {
		goto L79
	}
L73:
	;
	v351 = int32(0)
	v355 = v294
	v360 = v301
	goto L72
L74:
	;
	goto L75
L75:
	;
	v317 = int32(0)
	v321 = v294
	v322 = v294
	v326 = v301
	goto L76
L76:
	;
	v327 = int32(2)
	v329 = v310 + v317<<(uint(v327)%32)
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v329)))
	v331 = *(*float64)(unsafe.Add(mBase, uint32(v330)+56))
	v332 = *(*float64)(unsafe.Add(mBase, uint32(v330)+64))
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v329)+4))
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v335)+56))
	v337 = *(*float64)(unsafe.Add(mBase, uint32(v335)+64))
	v339 = base.F64_add(base.F64_add(v326, base.F64_add(v331, v332)), base.F64_add(v336, v337))
	v340 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335)+38)))
	v341 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v330)+38)))
	v345 = v340&v341 ^ int32(1) | v321
	v347 = v317 + v327
	v349 = v322 + v327
	if v349 != v305&int32(2147483646) {
		v317 = v347
		v321 = v345
		v322 = v349
		v326 = v339
		goto L76
	} else {
		goto L78
	}
L77:
	;
	v351 = v347
	v355 = v345
	v360 = v339
	goto L72
L78:
	;
	goto L77
L79:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v310+v351<<(uint(int32(2))%32))))
	v367 = *(*float64)(unsafe.Add(mBase, uint32(v366)+56))
	v368 = *(*float64)(unsafe.Add(mBase, uint32(v366)+64))
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v366)+38)))
	v379 = v371 ^ int32(1) | v355
	v384 = base.F64_add(v360, base.F64_add(v367, v368))
	goto L68
L80:
	;
	if v379&int32(1) != 0 {
		goto L89
	} else {
		goto L90
	}
L81:
	;
	v388 = int32(0)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v389 <= v388 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v396 = v388
	goto L83
L83:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v385)+12))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v404+v396<<(uint(int32(2))%32))))
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v408)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v408)+48)) = base.F64_add(v384, v409)
	v412 = *(*float64)(unsafe.Add(mBase, uint32(v408)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v408)+56)) = base.F64_add(v384, v412)
	if v379&int32(1) != 0 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	goto L80
L85:
	;
	v415 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+21)) = uint8(v415)
	goto L87
L86:
	;
	goto L87
L87:
	;
	v418 = v396 + int32(1)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v385)+4))
	if v418 < v419 {
		v396 = v418
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L84
L89:
	;
	v433 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v290)+26)) = uint8(v433)
	*(*int32)(unsafe.Add(mBase, uint32(v290)+40)) = v433
	goto L65
L90:
	;
	goto L91
L91:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v290)+40))
	if v437 == int32(0) {
		goto L66
	} else {
		goto L92
	}
L92:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if v440 <= int32(0) {
		goto L66
	} else {
		goto L93
	}
L93:
	;
	v446 = int32(0)
	goto L94
L94:
	;
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v437)+12))
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v454+v446<<(uint(int32(2))%32))))
	v459 = *(*float64)(unsafe.Add(mBase, uint32(v458)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v458)+48)) = base.F64_add(v384, v459)
	v462 = *(*float64)(unsafe.Add(mBase, uint32(v458)+56))
	*(*float64)(unsafe.Add(mBase, uint32(v458)+56)) = base.F64_add(v384, v462)
	v466 = v446 + int32(1)
	v467 = *(*int32)(unsafe.Add(mBase, uint32(v437)+4))
	if v466 < v467 {
		v446 = v466
		goto L94
	} else {
		goto L96
	}
L95:
	;
	goto L66
L96:
	;
	goto L95
L97:
	;
	v497 = base.F64_div(v491, v492)
	goto L99
L98:
	;
	v497 = v491
	goto L99
L99:
	;
	v498 = int32(0)
	if v489 != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	if v604 != 0 {
		goto L139
	} else {
		goto L140
	}
L101:
	;
	goto L100
L102:
	;
	v513 = v498
	v516 = v498
	goto L107
L103:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	if int32(0) < v505 {
		goto L102
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v604 = v498
	goto L101
L106:
	;
	goto L105
L107:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v489)+12))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v518+v516<<(uint(int32(2))%32))))
	if v513 != 0 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	v604 = v590
	goto L101
L109:
	;
	v596 = v516 + int32(1)
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v489)+4))
	if v596 < v597 {
		v513 = v590
		v516 = v596
		goto L107
	} else {
		goto L138
	}
L110:
	;
	v523 = F_compare_fractional_path_costs(m, v513, v522, v497)
	mBase = m.M
	if v523 <= int32(0) {
		v590 = v513
		goto L109
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v522)+64))
	if v490 == v526 {
		goto L114
	} else {
		goto L115
	}
L113:
	;
	goto L112
L114:
	;
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v522)+16))
	if v578 != 0 {
		goto L132
	} else {
		goto L133
	}
L115:
	;
	v532 = int32(0)
	goto L116
L116:
	;
	v539 = int32(0)
	if v490 == v539 {
		v549 = v539
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v549 != 0 {
		v590 = v513
		goto L109
	} else {
		goto L131
	}
L118:
	;
	if v526 != 0 {
		goto L122
	} else {
		goto L123
	}
L119:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v490)+4))
	if v543 <= v532 {
		v549 = int32(0)
		goto L118
	} else {
		goto L120
	}
L120:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v490)+12))
	v549 = v545 + v532<<(uint(int32(2))%32)
	goto L118
L121:
	;
	if v549 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L122:
	;
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v526)+4))
	if v532 < v550 {
		goto L121
	} else {
		goto L125
	}
L123:
	;
	goto L124
L124:
	;
	if v549 == int32(0) {
		goto L114
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	v590 = v513
	goto L109
L127:
	;
	goto L117
L128:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v526)+12))
	v559 = v556 + v532<<(uint(int32(2))%32)
	if v559 == int32(0) {
		goto L127
	} else {
		goto L129
	}
L129:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v549)))
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v559)))
	if v564 == v565 {
		v532 = v532 + int32(1)
		goto L116
	} else {
		goto L130
	}
L130:
	;
	v590 = v513
	goto L109
L131:
	;
	goto L114
L132:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v578)+4))
	v581 = v579
	goto L134
L133:
	;
	v581 = int32(0)
	goto L134
L134:
	;
	v583 = F_bms_is_subset(m, v581, int32(0))
	mBase = m.M
	if v583 != 0 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v584 = v522
	goto L137
L136:
	;
	v584 = v513
	goto L137
L137:
	;
	v590 = v584
	goto L109
L138:
	;
	goto L108
L139:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v25)+264))
	v610 = F_make_pathtarget_from_tlist(m, v609)
	mBase = m.M
	v611 = m.ExcPending
	if v611 != 0 {
		goto L1
	} else {
		goto L142
	}
L140:
	;
	goto L141
L141:
	;
	m.G0 = v17 + int32(16)
	return base.B2i32(v604 != int32(0))
L142:
	;
	v612 = F_set_pathtarget_cost_width(m, v25, v610)
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v614 = F_apply_projection_to_path(m, v25, v290, v604, v612)
	mBase = m.M
	v615 = m.ExcPending
	if v615 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v616 = *(*float64)(unsafe.Add(mBase, uint32(v614)+56))
	v617 = *(*float64)(unsafe.Add(mBase, uint32(v614)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v614
	*(*int32)(unsafe.Add(mBase, uint32(l1)+16)) = v25
	*(*float64)(unsafe.Add(mBase, uint32(l1)+24)) = base.F64_add(v617, base.F64_mul(v497, base.F64_sub(v616, v617)))
	goto L141
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
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
	v133 = m.ExcPending
	if v133 != 0 {
		goto L34
	} else {
		goto L41
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L34
	} else {
		goto L37
	}
L3:
	;
	v100 = F_builtin_locale_encoding(m, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L34
	} else {
		goto L35
	}
L4:
	;
	v15 = int32(583358)
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1118])))
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20 == int32(0) {
		v39 = v19
		v40 = v20
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
	v99 = int32(571401)
	goto L3
L7:
	;
	if v40-v39 == int32(0) {
		v99 = v15
		goto L3
	} else {
		goto L15
	}
L8:
	;
	goto L7
L9:
	;
	if v19 != v20 {
		v39 = v19
		v40 = v20
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v24 = l1
	v25 = v15
	goto L11
L11:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v28
		v40 = v29
		goto L8
	} else {
		goto L13
	}
L12:
	;
	v39 = v28
	v40 = v29
	goto L8
L13:
	;
	v32 = int32(1)
	if v28 == v29 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v44 = int32(582958)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1119])))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v48 == int32(0) {
		v67 = v47
		v68 = v48
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v68-v67 == int32(0) {
		v99 = v15
		goto L3
	} else {
		goto L24
	}
L17:
	;
	goto L16
L18:
	;
	if v47 != v48 {
		v67 = v47
		v68 = v48
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v52 = l1
	v53 = v44
	goto L20
L20:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+1)))
	if v57 == int32(0) {
		v67 = v56
		v68 = v57
		goto L17
	} else {
		goto L22
	}
L21:
	;
	v67 = v56
	v68 = v57
	goto L17
L22:
	;
	v60 = int32(1)
	if v56 == v57 {
		v52 = v52 + v60
		v53 = v53 + v60
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	v72 = int32(543271)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1120])))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v77 == int32(0) {
		v96 = v76
		v97 = v77
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v97-v96 != 0 {
		goto L2
	} else {
		goto L33
	}
L26:
	;
	goto L25
L27:
	;
	if v76 != v77 {
		v96 = v76
		v97 = v77
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v81 = l1
	v82 = v72
	goto L29
L29:
	;
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82)+1)))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81)+1)))
	if v86 == int32(0) {
		v96 = v85
		v97 = v86
		goto L26
	} else {
		goto L31
	}
L30:
	;
	v96 = v85
	v97 = v86
	goto L26
L31:
	;
	v89 = int32(1)
	if v85 == v86 {
		v81 = v81 + v89
		v82 = v82 + v89
		goto L29
	} else {
		goto L32
	}
L32:
	;
	goto L30
L33:
	;
	v99 = v72
	goto L3
L34:
	;
	return int32(0)
L35:
	;
	if base.B2i32(int32(0) <= v100)&base.B2i32(l0 != v100) != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	m.G0 = v8 + int32(32)
	return v99
L37:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L34
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l1
	F_errmsg(m, int32(239215), v8+int32(16))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L34
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(524117), int32(1526), int32(416803))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L34
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
	F_errcode(m, int32(151027844))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L34
	} else {
		goto L42
	}
L42:
	;
	if base.Ui32(l0) <= base.Ui32(int32(41)) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v146
	F_errmsg(m, int32(755757), v8)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L34
	} else {
		goto L47
	}
L44:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(3))%32))+uint32(_consts[363])))
	v146 = v145
	goto L46
L45:
	;
	v146 = int32(791891)
	goto L46
L46:
	;
	goto L43
L47:
	;
	F_errfinish(m, int32(524117), int32(1533), int32(416803))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L34
	} else {
		goto L48
	}
L48:
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
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
				v17 = int32(4)
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
				if v19&int32(254) == int32(2) {
					v28 = v17
				} else {
					v28 = base.B2i32(v19 == int32(18)) << (uint(v17) % 32)
				}
				if v19 == int32(1) {
					v31 = v17
				} else {
					v31 = v28
				}
				v32 = F_bytea_overlay(m, v6, v11, v13, v31)
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					return v32
				}
			} else {
				if v14&int32(1) != 0 {
					v37 = int32(1)
					v41 = F_bytea_overlay(m, v6, v11, v13, int32(base.Ui32(v14)>>(uint(v37)%32))-v37)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						return v41
					}
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
					v49 = F_bytea_overlay(m, v6, v11, v13, int32(base.Ui32(v44)>>(uint(int32(2))%32))-int32(4))
					mBase = m.M
					v50 = m.ExcPending
					if v50 != 0 {
						return int32(0)
					} else {
						return v49
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
