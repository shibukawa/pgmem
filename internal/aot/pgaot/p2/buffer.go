package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BufferGetBlockNumber(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	if l0 < int32(0) {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetBlockNumber[0]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5+(l0^int32(-1))*int32(56))+16))
		return v11
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetBlockNumber[1]))
		v15 = int32(56)
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+l0*v15-v15)+16))
		return v20
	}
}
func F_BufferGetTag(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	if l0 < int32(0) {
		v9 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetTag[0]))
		v22 = v9 + (l0^int32(-1))*int32(56)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetTag[1]))
		v17 = int32(56)
		v22 = v16 + l0*v17 - v17
	}
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v23
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v29
	return
}
func F_BufferManagerShmemRequest(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v2 = m.G0
	v4 = v2 + int32(-64)
	m.G0 = v4
	*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = int32(_a_F_BufferManagerShmemRequest_0)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+60)) = int32(_a_F_BufferManagerShmemRequest_1)
	*(*int32)(unsafe.Add(mBase, uint32(v4)+56)) = int32(128)
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemRequest[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v4)+52)) = v13 * int32(56)
	F_ShmemRequestStructWithOpts(m, v2+int32(-16))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v4)+32)) = int32(_a_F_BufferManagerShmemRequest_2)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+44)) = int32(_a_F_BufferManagerShmemRequest_3)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+40)) = int32(_a_F_BufferManagerShmemRequest_4)
		v28 = *(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemRequest[0]))
		*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v28 << (uint(int32(13)) % 32)
		F_ShmemRequestStructWithOpts(m, v2+int32(-32))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+16)) = int32(_a_F_BufferManagerShmemRequest_5)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+28)) = int32(_a_F_BufferManagerShmemRequest_6)
			*(*int32)(unsafe.Add(mBase, uint32(v4)+24)) = int32(128)
			v43 = *(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemRequest[0]))
			*(*int32)(unsafe.Add(mBase, uint32(v4)+20)) = v43 << (uint(int32(4)) % 32)
			F_ShmemRequestStructWithOpts(m, v2+int32(-48))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v4))) = int32(_a_F_BufferManagerShmemRequest_7)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+12)) = int32(_a_F_BufferManagerShmemRequest_8)
				*(*int32)(unsafe.Add(mBase, uint32(v4)+8)) = int32(0)
				v58 = *(*int32)(unsafe.Add(mBase, _c_F_BufferManagerShmemRequest[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v4)+4)) = v58 * int32(20)
				F_ShmemRequestStructWithOpts(m, v4)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					m.G0 = v4 - int32(-64)
					return
				}
			}
		}
	}
}
func F_BufferSetHintBits16(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v29 int32
	_ = v29
	var v31 int64
	_ = v31
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
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	v2 = l1
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l2 < int32(0) {
		*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v2)
		v13 = int64(0)
		v15 = *(*int32)(unsafe.Add(mBase, _c_F_BufferSetHintBits16[0]))
		v20 = v15 + (l2^int32(-1))*int32(56)
		v24 = base.AtomicRmwCmpxchg64(m, v20, int32(24), v13, v13)
		if v24&int64(8388608) == v13 {
			v29 = int32(_a_F_BufferSetHintBits16_0)
			v31 = *(*int64)(unsafe.Add(mBase, _c_F_BufferSetHintBits16[1]))
			*(*int64)(unsafe.Add(mBase, _c_F_BufferSetHintBits16[1])) = v31 + int64(1)
		} else {
		}
		*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v24 | int64(8388608)
		m.G0 = v8 + int32(16)
		return
	} else {
		v39 = *(*int32)(unsafe.Add(mBase, _c_F_BufferSetHintBits16[2]))
		v40 = int32(56)
		v44 = v39 + l2*v40 - v40
		v47 = F_SharedBufferBeginSetHintBits(m, l2, v44, v8+int32(8))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			if v47 == int32(0) {
				m.G0 = v8 + int32(16)
				return
			} else {
				*(*uint16)(unsafe.Add(mBase, uint32(l0))) = uint16(v2)
				v52 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
				F_MarkSharedBufferDirtyHint(m, l2, v44, v52, int32(1))
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		}
	}
}
func F_ExecStoreBufferHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v5 == int32(_a_F_ExecStoreBufferHeapTuple_0) {
		v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v8&int32(4) != 0 {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
			F_pfree(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v17 = v14 & int32(-5)
				v18 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
				v24 = v17 & int32(_a_F_ExecStoreBufferHeapTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v24)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v26
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v28)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
				if v30 == l2 {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
					return
				} else {
					if v30 != 0 {
						F_ReleaseBuffer(m, v30)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = l2
							if l2 == int32(0) {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
								return
							} else {
								F_IncrBufferRefCount(m, l2)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
									return
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = l2
						if l2 == int32(0) {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
							return
						} else {
							F_IncrBufferRefCount(m, l2)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
								return
							}
						}
					}
				}
			}
		} else {
			v17 = v8
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = l0
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
			v24 = v17 & int32(_a_F_ExecStoreBufferHeapTuple_1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v26
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v28)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
			if v30 == l2 {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
				return
			} else {
				if v30 != 0 {
					F_ReleaseBuffer(m, v30)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = l2
						if l2 == int32(0) {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
							return
						} else {
							F_IncrBufferRefCount(m, l2)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = l2
					if l2 == int32(0) {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
						return
					} else {
						F_IncrBufferRefCount(m, l2)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v39
							return
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_ExecStoreBufferHeapTuple_2), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_ExecStoreBufferHeapTuple_3), int32(1687), int32(_a_F_ExecStoreBufferHeapTuple_4))
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
	}
}
func F_StartReadBuffer(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v151 int64
	_ = v151
	var v162 int64
	_ = v162
	var v190 int32
	_ = v190
	var v191 int64
	_ = v191
	var v194 int64
	_ = v194
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v241 int64
	_ = v241
	var v243 int64
	_ = v243
	var v264 int64
	_ = v264
	var v267 int32
	_ = v267
	var v269 int64
	_ = v269
	var v271 int64
	_ = v271
	var v274 int64
	_ = v274
	var v275 int64
	_ = v275
	var v280 int64
	_ = v280
	var v283 int64
	_ = v283
	var v288 int64
	_ = v288
	var v306 int64
	_ = v306
	var v313 int64
	_ = v313
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v347 int64
	_ = v347
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v377 int32
	_ = v377
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v415 int32
	_ = v415
	var v417 int64
	_ = v417
	var v421 int32
	_ = v421
	var v434 int32
	_ = v434
	var v438 int64
	_ = v438
	var v442 int64
	_ = v442
	var v444 int32
	_ = v444
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int64
	_ = v476
	var v484 int32
	_ = v484
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int64
	_ = v509
	var v517 int32
	_ = v517
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	v4 = l3
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	v24 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v27 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v517)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v532 + int32(1)
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	if v536 != 0 {
		goto L101
	} else {
		goto L102
	}
L2:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v51)+272))
	if v499 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L3:
	;
	v434 = v48*int32(320) + v47<<(uint(int32(6))%32)
	v438 = *(*int64)(unsafe.Add(mBase, uint32(v434)+uint32(_c_F_StartReadBuffer[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v434)+uint32(_c_F_StartReadBuffer[0]))) = v438 + int64(1)
	v442 = *(*int64)(unsafe.Add(mBase, uint32(v434)+uint32(_c_F_StartReadBuffer[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v434)+uint32(_c_F_StartReadBuffer[1]))) = v442
	v444 = int32(1)
	F_pgstat_count_backend_io_op(m, v48, v47, int32(2), v444, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_StartReadBuffer[2])) = uint8(v444)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartReadBuffer[3])) = uint8(v444)
	goto L85
L4:
	;
	v415 = int32(_a_F_StartReadBuffer_0)
	v417 = *(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[4])) = v417 + int64(1)
	v421 = v410
	goto L3
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L13
	} else {
		goto L81
	}
L6:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v38 != int32(116) {
		goto L10
	} else {
		goto L11
	}
L7:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+48))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+118)))
	if v31 != int32(116) {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+24)))
	if v34 == int32(0) {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	goto L6
L10:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v43 = F_IOContextForStrategy(m, v42)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v47 = int32(3)
	v48 = v24
	goto L12
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v52 != int32(116) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	return int32(0)
L14:
	;
	v47 = v43
	v48 = int32(0)
	goto L12
L15:
	;
	if v51 == int32(0) {
		v517 = v377
		goto L1
	} else {
		goto L80
	}
L16:
	;
	v352 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[5]))
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	F_ResourceOwnerForget(m, v352, base.I64_extend_i32_s(v353+int32(1)), int32(_a_F_StartReadBuffer_1))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L13
	} else {
		goto L75
	}
L17:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[5]))
	F_ResourceOwnerEnlarge(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v340 = F_LocalBufferAlloc(m, v50, v49, l2, v22+int32(24))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L13
	} else {
		goto L73
	}
L20:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+4)) = v62
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+8)) = v64
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v49
	*(*int32)(unsafe.Add(mBase, uint32(v22)+12)) = v66
	v71 = v22 + int32(4)
	v72 = F_BufTableHashCode(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L13
	} else {
		goto L22
	}
L22:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[6]))
	v82 = v75 + v72&int32(127)<<(uint(int32(7))%32) + int32(_a_F_StartReadBuffer_2)
	v84 = F_LWLockAcquire(m, v82, int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L23
	}
L23:
	;
	v86 = F_BufTableLookup(m, v71, v72)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L13
	} else {
		goto L24
	}
L24:
	;
	if int32(0) <= v86 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v91 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[7]))
	v94 = v91 + v86*int32(56)
	v96 = F_PinBuffer(m, v94, v55, int32(0))
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L13
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	F_LWLockRelease(m, v82)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L13
	} else {
		goto L31
	}
L28:
	;
	F_LWLockRelease(m, v82)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)) = uint8(v96)
	if v96 == int32(0) {
		v377 = v94
		goto L15
	} else {
		goto L30
	}
L30:
	;
	v410 = v94
	goto L4
L31:
	;
	v105 = F_GetVictimBuffer(m, v55, v47)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L13
	} else {
		goto L32
	}
L32:
	;
	v108 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[7]))
	v110 = F_LWLockAcquire(m, v82, int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	v112 = int32(56)
	v114 = v108 + v105*v112
	v116 = v114 - v112
	v120 = v114 - int32(36)
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	v122 = F_BufTableInsert(m, v22+int32(4), v72, v121)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	if int32(0) <= v122 {
		goto L16
	} else {
		goto L35
	}
L35:
	;
	v127 = v114 - int32(32)
	v128 = int64(4194304)
	v130 = base.AtomicRmwOr64(m, v127, int32(0), v128)
	if v130&v128 != int64(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v151 = v130
	goto L39
L37:
	;
	v264 = v130
	goto L38
L38:
	;
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v22)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v116)+16)) = v267
	v269 = *(*int64)(unsafe.Add(mBase, uint32(v22)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v116)+8)) = v269
	v271 = *(*int64)(unsafe.Add(mBase, uint32(v22)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v116))) = v271
	v274 = v264 | int64(4194304)
	v275 = int64(2181300224)
	if v49 == int32(3) {
		goto L60
	} else {
		goto L61
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+44)) = int32(_a_F_StartReadBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+40)) = int32(_a_F_StartReadBuffer_4)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+36)) = int32(_a_F_StartReadBuffer_5)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+32)) = int32(0)
	v162 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v22)+24)) = v162
	if v151&int64(4194304) != v162 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v264 = v243
	goto L38
L41:
	;
	goto L44
L42:
	;
	goto L43
L43:
	;
	v221 = int32(_a_F_StartReadBuffer_6)
	v222 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[8]))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v22+int32(24))+8))
	if v224 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	F_perform_spin_delay(m, v22+int32(24))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L13
	} else {
		goto L46
	}
L45:
	;
	goto L43
L46:
	;
	v191 = int64(0)
	v194 = base.AtomicRmwCmpxchg64(m, v127, int32(0), v191, v191)
	if v194&int64(4194304) != v191 {
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v241 = int64(4194304)
	v243 = base.AtomicRmwOr64(m, v127, int32(0), v241)
	if v243&v241 != int64(0) {
		v151 = v243
		goto L39
	} else {
		goto L59
	}
L49:
	;
	goto L48
L50:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[8])) = v239
	goto L49
L51:
	;
	if int32(999) < v222 {
		goto L49
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	if v222 < int32(11) {
		goto L49
	} else {
		goto L58
	}
L54:
	;
	v229 = int32(900)
	if v229 <= v222 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v232 = v229
	goto L57
L56:
	;
	v232 = v222
	goto L57
L57:
	;
	v239 = v232 + int32(100)
	goto L50
L58:
	;
	v239 = v222 - int32(1)
	goto L50
L59:
	;
	goto L40
L60:
	;
	v280 = v275
	goto L62
L61:
	;
	v280 = int64(33816576)
	goto L62
L62:
	;
	if v52 == int32(112) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v283 = v275
	goto L65
L64:
	;
	v283 = v280
	goto L65
L65:
	;
	v288 = base.AtomicRmwCmpxchg64(m, v127, int32(0), v274, v283|v264&int64(-38010881))
	if v288 != v274 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v306 = v288
	goto L69
L67:
	;
	goto L68
L68:
	;
	F_LWLockRelease(m, v82)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L13
	} else {
		goto L72
	}
L69:
	;
	v313 = base.AtomicRmwCmpxchg64(m, v127, int32(0), v306, v306&int64(-38010881)|v283)
	if v306 != v313 {
		v306 = v313
		goto L69
	} else {
		goto L71
	}
L70:
	;
	goto L68
L71:
	;
	goto L70
L72:
	;
	v336 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)) = uint8(v336)
	v377 = v116
	goto L15
L73:
	;
	v342 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)))
	if v342 != int32(1) {
		v377 = v340
		goto L15
	} else {
		goto L74
	}
L74:
	;
	v345 = int32(_a_F_StartReadBuffer_7)
	v347 = *(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[9]))
	*(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[9])) = v347 + int64(1)
	v421 = v340
	goto L3
L75:
	;
	F_UnpinBufferNoOwner(m, v116)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L13
	} else {
		goto L76
	}
L76:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[7]))
	v366 = v363 + v122*int32(56)
	v368 = F_PinBuffer(m, v366, v55, int32(0))
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L13
	} else {
		goto L77
	}
L77:
	;
	F_LWLockRelease(m, v82)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L13
	} else {
		goto L78
	}
L78:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v22)+24)) = uint8(v368)
	if v368 != 0 {
		v410 = v366
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v377 = v366
	goto L15
L80:
	;
	v484 = v377
	goto L2
L81:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L13
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_StartReadBuffer_8), int32(0))
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L13
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_StartReadBuffer_5), int32(1392), int32(_a_F_StartReadBuffer_9))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L13
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L85:
	;
	v454 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartReadBuffer[10])))
	if v454 == int32(1) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v457 = int32(_a_F_StartReadBuffer_10)
	v459 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[11]))
	v461 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[11])) = v459 + v461
	goto L88
L87:
	;
	goto L88
L88:
	;
	if v51 == int32(0) {
		v517 = v421
		goto L1
	} else {
		goto L89
	}
L89:
	;
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v51)+272))
	if v466 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+268)))
	if v469 != int32(1) {
		v484 = v421
		goto L2
	} else {
		goto L93
	}
L91:
	;
	v475 = v466
	goto L92
L92:
	;
	v476 = *(*int64)(unsafe.Add(mBase, uint32(v475)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v475)+120)) = v476 + int64(1)
	v484 = v421
	goto L2
L93:
	;
	F_pgstat_assoc_relation(m, v51)
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L13
	} else {
		goto L94
	}
L94:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v51)+272))
	v475 = v474
	goto L92
L95:
	;
	v502 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+268)))
	if v502 != int32(1) {
		v517 = v484
		goto L1
	} else {
		goto L98
	}
L96:
	;
	v508 = v499
	goto L97
L97:
	;
	v509 = *(*int64)(unsafe.Add(mBase, uint32(v508)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v508)+112)) = v509 + int64(1)
	v517 = v484
	goto L1
L98:
	;
	F_pgstat_assoc_relation(m, v51)
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L13
	} else {
		goto L99
	}
L99:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v51)+272))
	v508 = v507
	goto L97
L100:
	;
	m.G0 = v22 + int32(48)
	return v570
L101:
	;
	v570 = int32(0)
	goto L100
L102:
	;
	goto L103
L103:
	;
	v538 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v22))) = v538
	*(*int32)(unsafe.Add(mBase, uint32(l0)+30)) = v538
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v4)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
	goto L104
L104:
	;
	v550 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[13]))
	if v550 != 0 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v551 = F_AsyncReadBuffers(m, l0, v22)
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L13
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)))
	v557 = v555 | int32(8)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+28)) = uint16(v557)
	if v4&int32(2) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+30)) = uint16(v553)
	v570 = v551
	goto L100
L109:
	;
	v570 = int32(1)
	goto L100
L110:
	;
	goto L111
L111:
	;
	v564 = int32(1)
	v565 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v566 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v568 = F_smgrprefetch(m, v565, v566, l2, v564)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L13
	} else {
		goto L112
	}
L112:
	;
	v570 = v564
	goto L100
}
func F_UnlockBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	if int32(0) <= l0 {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_UnlockBuffer[0]))
		v6 = int32(56)
		F_BufferLockUnlock(m, l0, v5+l0*v6-v6)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_UnpinBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBuffer[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerForget(m, v3, base.I64_extend_i32_s(v4+int32(1)), int32(_a_F_UnpinBuffer_0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		F_UnpinBufferNoOwner(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			return
		}
	}
}
