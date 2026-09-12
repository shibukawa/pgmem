package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_InitSharedLatch(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = int64(0)
	v6 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v6)
	return
}
func F_SharedInvalBackendInit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	v1 = l0
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = *(*int32)(unsafe.Add(mBase, _consts[100]))
	if int32(0) <= v13 {
		v17 = *(*int32)(unsafe.Add(mBase, _consts[102]))
		if v17+int32(38) <= v13 {
			F_errstart_cold(m, int32(23), int32(0))
			mBase = m.M
			v90 = m.ExcPending
			if v90 != 0 {
				return
			} else {
				v92 = *(*int32)(unsafe.Add(mBase, _consts[100]))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v92
				v95 = *(*int32)(unsafe.Add(mBase, _consts[102]))
				*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v95 + int32(38)
				F_errmsg_internal(m, int32(661873), v10+int32(16))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return
				} else {
					F_errfinish(m, int32(490282), int32(282), int32(100015))
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
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, _consts[837]))
			v24 = *(*int32)(unsafe.Add(mBase, _consts[44]))
			v28 = F_LWLockAcquire(m, v24+int32(768), int32(0))
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return
			} else {
				v32 = v22 + v13<<(uint(int32(4))%32)
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[838])))
				if v35 != 0 {
					v110 = *(*int32)(unsafe.Add(mBase, _consts[44]))
					F_LWLockRelease(m, v110+int32(768))
					mBase = m.M
					v114 = m.ExcPending
					if v114 != 0 {
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v35
							v121 = *(*int32)(unsafe.Add(mBase, _consts[100]))
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v121
							F_errmsg_internal(m, int32(466393), v10)
							mBase = m.M
							v125 = m.ExcPending
							if v125 != 0 {
								return
							} else {
								F_errfinish(m, int32(490282), int32(297), int32(100015))
								mBase = m.M
								v130 = m.ExcPending
								if v130 != 0 {
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
					v37 = *(*int32)(unsafe.Add(mBase, _consts[100]))
					v39 = *(*int32)(unsafe.Add(mBase, _consts[837]))
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[839])))
					*(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[839]))) = v40 + int32(1)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v39)+uint32(_consts[840])))
					*(*int32)(unsafe.Add(mBase, uint32(v44+v40<<(uint(int32(2))%32)))) = v37
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[841])))
					*(*int32)(unsafe.Add(mBase, _consts[274])) = v50
					v53 = *(*int32)(unsafe.Add(mBase, _consts[129]))
					*(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[838]))) = v53
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[842]))) = uint8(v1)
					v57 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[843]))) = uint8(v57)
					*(*uint16)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[844]))) = uint16(v57)
					*(*int32)(unsafe.Add(mBase, uint32(v32)+uint32(_consts[845]))) = v55
					v63 = *(*int32)(unsafe.Add(mBase, _consts[44]))
					F_LWLockRelease(m, v63+int32(768))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
						return
					} else {
						F_on_shmem_exit(m, int32(1106), v22)
						mBase = m.M
						v70 = m.ExcPending
						if v70 != 0 {
							return
						} else {
							m.G0 = v10 + int32(32)
							return
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(105679), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return
			} else {
				F_errfinish(m, int32(490282), int32(279), int32(100015))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
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
func F_UnlockSharedObjectForSession(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v2 = int32(0)
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = int32(264)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+14)) = uint16(v7)
	*(*uint16)(unsafe.Add(mBase, uint32(v5)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v5)+4)) = int32(1262)
	*(*int32)(unsafe.Add(mBase, uint32(v5))) = v2
	v18 = F_LockRelease(m, v5, int32(8), int32(1))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return
	} else {
		m.G0 = v5 + int32(16)
		return
	}
}
func F_process_shared_preload_libraries(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v2 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[1111])) = uint8(v2)
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1112]))
	F_load_libraries(m, v5, int32(167411), int32(0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		v11 = int32(1)
		*(*uint8)(unsafe.Add(mBase, _consts[1113])) = uint8(v11)
		v14 = int32(0)
		*(*uint8)(unsafe.Add(mBase, _consts[1111])) = uint8(v14)
		return
	}
}
func F_shared_buffer_readv_complete_local(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v5&int32(448) == int32(64) {
		v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
		*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
		return
	} else {
		if v5&int32(33292288) == int32(0) {
			v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
			*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l1+int32(104))+4))
			F_pgstat_report_checksum_failures_in_db(m, v16, int32(base.Ui32(v5)>>(uint(int32(18))%32))&int32(127))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
				*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
				return
			}
		}
	}
}
func F_shared_record_typmod_registry_detach(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	v4 = *(*int32)(unsafe.Add(mBase, _consts[108]))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	if v5 != 0 {
		F_pfree(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			v9 = *(*int32)(unsafe.Add(mBase, _consts[108]))
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
			v12 = v9
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			if v13 != 0 {
				F_pfree(m, v13)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v17 = *(*int32)(unsafe.Add(mBase, _consts[108]))
					*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(0)
					v20 = v17
					*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(0)
					return
				}
			} else {
				v20 = v12
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(0)
				return
			}
		}
	} else {
		v12 = v4
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
		if v13 != 0 {
			F_pfree(m, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, _consts[108]))
				*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(0)
				v20 = v17
				*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(0)
				return
			}
		} else {
			v20 = v12
			*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = int32(0)
			return
		}
	}
}
func F_shared_ts_extend_down(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v8 int32
	_ = v8
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
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int32
	_ = v33
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
	var v44 int64
	_ = v44
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	v3 = l2
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v11 = F_dsa_allocate_extended(m, v8, int32(24), int32(0))
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = F_dsa_get_address(m, v15, v11)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(1024)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v11
	v22 = l3 - int32(8)
	if v22 <= int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v59 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+2)) = uint8(v59)
	*(*uint8)(unsafe.Add(mBase, uint32(v53)+3)) = uint8(v3)
	return v53 + int32(8)
L5:
	;
	v53 = v16
	goto L4
L6:
	;
	goto L7
L7:
	;
	v31 = v16
	v32 = base.I64_extend_i32_u(v22)
	goto L8
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v36 = F_dsa_allocate_extended(m, v33, int32(24), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	v53 = v39
	goto L4
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v39 = F_dsa_get_address(m, v38, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v39))) = int64(1024)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+8)) = v36
	v44 = int64(base.Ui64(v3) >> (uint(v32) % 64))
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+3)) = uint8(v44)
	v46 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+2)) = uint8(v46)
	v48 = int64(8)
	if base.Ui64(v48) < base.Ui64(v32) {
		v31 = v39
		v32 = v32 - v48
		goto L8
	} else {
		goto L12
	}
L12:
	;
	goto L9
}
