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
	var v20 int32
	_ = v20
	if l0 < int32(0) {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetBlockNumber[0]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5+(l0^int32(-1))<<(uint(int32(6))%32))+16))
		return v11
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetBlockNumber[1]))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+l0<<(uint(int32(6))%32)+int32(-64))+16))
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
		v22 = v9 + (l0^int32(-1))<<(uint(int32(6))%32)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _c_F_BufferGetTag[1]))
		v22 = v16 + l0<<(uint(int32(6))%32) + int32(-64)
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
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			F_pfree(m, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v17 = v14 & int32(-5)
				v18 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v18
				*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
				v24 = v17 & int32(_a_F_ExecStoreBufferHeapTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v24)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v26
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v28)
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
				if v30 == l2 {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
					return
				} else {
					if v30 != 0 {
						F_ReleaseBuffer(m, v30)
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l2
							if l2 == int32(0) {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
								return
							} else {
								F_IncrBufferRefCount(m, l2)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
									*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
									return
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l2
						if l2 == int32(0) {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
							return
						} else {
							F_IncrBufferRefCount(m, l2)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
								return
							}
						}
					}
				}
			}
		} else {
			v17 = v8
			v18 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v18
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = l0
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v18)
			v24 = v17 & int32(_a_F_ExecStoreBufferHeapTuple_1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v24)
			v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v26
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+8)))
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+32)) = uint16(v28)
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+68))
			if v30 == l2 {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
				return
			} else {
				if v30 != 0 {
					F_ReleaseBuffer(m, v30)
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l2
						if l2 == int32(0) {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
							return
						} else {
							F_IncrBufferRefCount(m, l2)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
								return
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l2
					if l2 == int32(0) {
						v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
						return
					} else {
						F_IncrBufferRefCount(m, l2)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return
						} else {
							v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
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
				F_errfinish(m, int32(_a_F_ExecStoreBufferHeapTuple_3), int32(1594), int32(_a_F_ExecStoreBufferHeapTuple_4))
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
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int64
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
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
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v171 int32
	_ = v171
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v261 int64
	_ = v261
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int64
	_ = v292
	var v296 int64
	_ = v296
	var v303 int32
	_ = v303
	var v304 int64
	_ = v304
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int64
	_ = v315
	var v330 int32
	_ = v330
	var v334 int64
	_ = v334
	var v338 int64
	_ = v338
	var v340 int32
	_ = v340
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v21
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v27 == int32(116) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if v26 != 0 {
		goto L61
	} else {
		goto L62
	}
L2:
	;
	v32 = F_LocalBufferAlloc(m, v25, v24, l2, v19+int32(24))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v48 = F_IOContextForStrategy(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L5
	} else {
		goto L8
	}
L5:
	;
	return int32(0)
L6:
	;
	v36 = int32(3)
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	if v37 != int32(1) {
		v271 = v32
		v273 = v5
		v274 = v21
		v278 = v36
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v40 = int32(_a_F_StartReadBuffer_0)
	v42 = *(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[0]))
	*(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[0])) = v42 + int64(1)
	v271 = v32
	v273 = int32(1)
	v274 = v21
	v278 = v36
	goto L1
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[1]))
	F_ResourceOwnerEnlarge(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v56
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v58
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v60
	v65 = v19 + int32(4)
	v66 = F_BufTableHashCode(m, v65)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[2]))
	v76 = v69 + v66&int32(127)<<(uint(int32(7))%32) + int32(_a_F_StartReadBuffer_1)
	v78 = F_LWLockAcquire(m, v76, int32(1))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v80 = F_BufTableLookup(m, v65, v66)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	v259 = int32(_a_F_StartReadBuffer_2)
	v261 = *(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[3])) = v261 + int64(1)
	v271 = v252
	v273 = int32(1)
	v274 = int32(0)
	v278 = v48
	goto L1
L14:
	;
	if int32(0) <= v80 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[4]))
	v88 = v85 + v80<<(uint(int32(6))%32)
	v89 = F_PinBuffer(m, v88, v47)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_LWLockRelease(m, v76)
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L21
	}
L18:
	;
	F_LWLockRelease(m, v76)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v89)
	if v89 != 0 {
		v252 = v88
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v271 = v88
	v273 = v5
	v274 = int32(0)
	v278 = v48
	goto L1
L21:
	;
	v97 = F_GetVictimBuffer(m, v47, v48)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[4]))
	v102 = F_LWLockAcquire(m, v76, int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v106 = v100 + v97<<(uint(int32(6))%32)
	v108 = v106 + int32(-64)
	v112 = v106 - int32(44)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v114 = F_BufTableInsert(m, v19+int32(4), v66, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	if v114 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = int32(_a_F_StartReadBuffer_3)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = int32(_a_F_StartReadBuffer_4)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = int32(_a_F_StartReadBuffer_5)
	v124 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v124
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = int64(0)
	v130 = v106 - int32(40)
	v131 = int32(_a_F_StartReadBuffer_6)
	v133 = base.AtomicRmwOr32(m, v130, v124, v131)
	if v133&v131 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v230 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[1]))
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	F_ResourceOwnerForget(m, v230, v231+int32(1), int32(_a_F_StartReadBuffer_7))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L53
	}
L28:
	;
	goto L31
L29:
	;
	v171 = v133
	goto L30
L30:
	;
	v180 = int32(_a_F_StartReadBuffer_8)
	v181 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[5]))
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(24))+8))
	if v183 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	F_perform_spin_delay(m, v19+int32(24))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L5
	} else {
		goto L33
	}
L32:
	;
	v171 = v158
	goto L30
L33:
	;
	v156 = int32(_a_F_StartReadBuffer_6)
	v158 = base.AtomicRmwOr32(m, v130, int32(0), v156)
	if v158&v156 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = v200
	v202 = *(*int64)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v202
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = v204
	v206 = int32(0)
	v209 = base.AtomicRmwOr32(m, v206, int32(_a_F_StartReadBuffer_9), v206)
	v212 = int32(-2113667072)
	if v24 == int32(3) {
		goto L46
	} else {
		goto L47
	}
L36:
	;
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[5])) = v198
	goto L36
L38:
	;
	if int32(999) < v181 {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v181 < int32(11) {
		goto L36
	} else {
		goto L45
	}
L41:
	;
	v188 = int32(900)
	if v188 <= v181 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v191 = v188
	goto L44
L43:
	;
	v191 = v181
	goto L44
L44:
	;
	v198 = v191 + int32(100)
	goto L37
L45:
	;
	v198 = v181 - int32(1)
	goto L37
L46:
	;
	v217 = v212
	goto L48
L47:
	;
	v217 = int32(33816576)
	goto L48
L48:
	;
	if v27 == int32(112) {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v220 = v212
	goto L51
L50:
	;
	v220 = v217
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v171&int32(-38010881) | v220
	F_LWLockRelease(m, v76)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	v225 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v225)
	v271 = v108
	v273 = v124
	v274 = v225
	v278 = v48
	goto L1
L53:
	;
	F_UnpinBufferNoOwner(m, v108)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_StrategyFreeBuffer(m, v108)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v242 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[4]))
	v245 = v242 + v114<<(uint(int32(6))%32)
	v246 = F_PinBuffer(m, v245, v47)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_LWLockRelease(m, v76)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v246)
	if v246 != 0 {
		v252 = v245
		goto L13
	} else {
		goto L58
	}
L58:
	;
	v271 = v245
	v273 = int32(0)
	v274 = int32(0)
	v278 = v48
	goto L1
L59:
	;
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v271)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v363 + int32(1)
	v367 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	if v367 != 0 {
		goto L83
	} else {
		goto L84
	}
L60:
	;
	v330 = v274*int32(320) + v278<<(uint(int32(6))%32)
	v334 = *(*int64)(unsafe.Add(mBase, uint32(v330)+uint32(_c_F_StartReadBuffer[6])))
	*(*int64)(unsafe.Add(mBase, uint32(v330)+uint32(_c_F_StartReadBuffer[6]))) = v334 + int64(1)
	v338 = *(*int64)(unsafe.Add(mBase, uint32(v330)+uint32(_c_F_StartReadBuffer[7])))
	*(*int64)(unsafe.Add(mBase, uint32(v330)+uint32(_c_F_StartReadBuffer[7]))) = v338
	v340 = int32(1)
	F_pgstat_count_backend_io_op(m, v274, v278, int32(2), v340, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_StartReadBuffer[8])) = uint8(v340)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartReadBuffer[9])) = uint8(v340)
	goto L80
L61:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	if v282 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	goto L63
L63:
	;
	if v273 == int32(0) {
		goto L59
	} else {
		goto L79
	}
L64:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	if v303 != 0 {
		goto L73
	} else {
		goto L74
	}
L65:
	;
	if v273 == int32(0) {
		goto L59
	} else {
		goto L72
	}
L66:
	;
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+268)))
	if v285 != int32(1) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v282)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v282)+112)) = v296 + int64(1)
	goto L65
L69:
	;
	F_pgstat_assoc_relation(m, v26)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	v290 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	v292 = *(*int64)(unsafe.Add(mBase, uint32(v291)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v291)+112)) = v292 + int64(1)
	if v290 != 0 {
		goto L64
	} else {
		goto L71
	}
L71:
	;
	goto L59
L72:
	;
	goto L64
L73:
	;
	v304 = *(*int64)(unsafe.Add(mBase, uint32(v303)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v303)+120)) = v304 + int64(1)
	goto L60
L74:
	;
	goto L75
L75:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+268)))
	if v308 != int32(1) {
		goto L60
	} else {
		goto L76
	}
L76:
	;
	F_pgstat_assoc_relation(m, v26)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v313 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v314)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v314)+120)) = v315 + int64(1)
	if v313 != 0 {
		goto L60
	} else {
		goto L78
	}
L78:
	;
	goto L59
L79:
	;
	goto L60
L80:
	;
	v350 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartReadBuffer[10])))
	if v350 != int32(1) {
		goto L59
	} else {
		goto L81
	}
L81:
	;
	v353 = int32(_a_F_StartReadBuffer_10)
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[11]))
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[11])) = v355 + v357
	goto L59
L82:
	;
	m.G0 = v19 + int32(48)
	return v401
L83:
	;
	v401 = int32(0)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v369 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v369
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v369
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
	goto L86
L86:
	;
	v381 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[13]))
	if v381 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v382 = F_AsyncReadBuffers(m, l0, v19)
	mBase = m.M
	v383 = m.ExcPending
	if v383 != 0 {
		goto L5
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v386 | int32(8)
	if l3&int32(2) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v384)
	v401 = v382
	goto L82
L91:
	;
	v401 = int32(1)
	goto L82
L92:
	;
	goto L93
L93:
	;
	v395 = int32(1)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v399 = F_smgrprefetch(m, v396, v397, l2, v395)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v401 = v395
	goto L82
}
func F_UnpinBuffer(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_UnpinBuffer[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerForget(m, v3, v4+int32(1), int32(_a_F_UnpinBuffer_0))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_UnpinBufferNoOwner(m, l0)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			return
		}
	}
}
