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
	var v132 int32
	_ = v132
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v257 int32
	_ = v257
	var v259 int64
	_ = v259
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v294 int64
	_ = v294
	var v301 int32
	_ = v301
	var v302 int64
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v328 int32
	_ = v328
	var v332 int64
	_ = v332
	var v336 int64
	_ = v336
	var v338 int32
	_ = v338
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
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
		v269 = v32
		v271 = v5
		v272 = v21
		v276 = v36
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v40 = int32(_a_F_StartReadBuffer_0)
	v42 = *(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[0]))
	*(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[0])) = v42 + int64(1)
	v269 = v32
	v271 = int32(1)
	v272 = v21
	v276 = v36
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
	v257 = int32(_a_F_StartReadBuffer_2)
	v259 = *(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[3]))
	*(*int64)(unsafe.Add(mBase, _c_F_StartReadBuffer[3])) = v259 + int64(1)
	v269 = v250
	v271 = int32(1)
	v272 = int32(0)
	v276 = v48
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
		v250 = v88
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v269 = v88
	v271 = v5
	v272 = int32(0)
	v276 = v48
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
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v132 = int32(_a_F_StartReadBuffer_6)
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 | v132
	if v131&v132 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v228 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[1]))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	F_ResourceOwnerForget(m, v228, v229+int32(1), int32(_a_F_StartReadBuffer_7))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L5
	} else {
		goto L53
	}
L28:
	;
	goto L31
L29:
	;
	v173 = v131
	goto L30
L30:
	;
	v182 = int32(_a_F_StartReadBuffer_8)
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[5]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(24))+8))
	if v185 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	F_perform_spin_delay(m, v19+int32(24))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L5
	} else {
		goto L33
	}
L32:
	;
	v173 = v157
	goto L30
L33:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v158 = int32(_a_F_StartReadBuffer_6)
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v157 | v158
	if v157&v158 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v108)+16)) = v202
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v108)+8)) = v204
	v206 = *(*int64)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = v206
	v210 = int32(-2113667072)
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
	*(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[5])) = v200
	goto L36
L38:
	;
	if int32(999) < v183 {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v183 < int32(11) {
		goto L36
	} else {
		goto L45
	}
L41:
	;
	v190 = int32(900)
	if v190 <= v183 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v193 = v190
	goto L44
L43:
	;
	v193 = v183
	goto L44
L44:
	;
	v200 = v193 + int32(100)
	goto L37
L45:
	;
	v200 = v183 - int32(1)
	goto L37
L46:
	;
	v215 = v210
	goto L48
L47:
	;
	v215 = int32(33816576)
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
	v218 = v210
	goto L51
L50:
	;
	v218 = v215
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v173&int32(-38010881) | v218
	F_LWLockRelease(m, v76)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	v223 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v223)
	v269 = v108
	v271 = v124
	v272 = v223
	v276 = v48
	goto L1
L53:
	;
	F_UnpinBufferNoOwner(m, v108)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_StrategyFreeBuffer(m, v108)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[4]))
	v243 = v240 + v114<<(uint(int32(6))%32)
	v244 = F_PinBuffer(m, v243, v47)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_LWLockRelease(m, v76)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v244)
	if v244 != 0 {
		v250 = v243
		goto L13
	} else {
		goto L58
	}
L58:
	;
	v269 = v243
	v271 = int32(0)
	v272 = int32(0)
	v276 = v48
	goto L1
L59:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v269)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v361 + int32(1)
	v365 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	if v365 != 0 {
		goto L83
	} else {
		goto L84
	}
L60:
	;
	v328 = v272*int32(320) + v276<<(uint(int32(6))%32)
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v328)+uint32(_c_F_StartReadBuffer[6])))
	*(*int64)(unsafe.Add(mBase, uint32(v328)+uint32(_c_F_StartReadBuffer[6]))) = v332 + int64(1)
	v336 = *(*int64)(unsafe.Add(mBase, uint32(v328)+uint32(_c_F_StartReadBuffer[7])))
	*(*int64)(unsafe.Add(mBase, uint32(v328)+uint32(_c_F_StartReadBuffer[7]))) = v336
	v338 = int32(1)
	F_pgstat_count_backend_io_op(m, v272, v276, int32(2), v338, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_StartReadBuffer[8])) = uint8(v338)
	*(*uint8)(unsafe.Add(mBase, _c_F_StartReadBuffer[9])) = uint8(v338)
	goto L80
L61:
	;
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	if v280 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	goto L63
L63:
	;
	if v271 == int32(0) {
		goto L59
	} else {
		goto L79
	}
L64:
	;
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	if v301 != 0 {
		goto L73
	} else {
		goto L74
	}
L65:
	;
	if v271 == int32(0) {
		goto L59
	} else {
		goto L72
	}
L66:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+268)))
	if v283 != int32(1) {
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(v280)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v280)+112)) = v294 + int64(1)
	goto L65
L69:
	;
	F_pgstat_assoc_relation(m, v26)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L5
	} else {
		goto L70
	}
L70:
	;
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	v290 = *(*int64)(unsafe.Add(mBase, uint32(v289)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v289)+112)) = v290 + int64(1)
	if v288 != 0 {
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
	v302 = *(*int64)(unsafe.Add(mBase, uint32(v301)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v301)+120)) = v302 + int64(1)
	goto L60
L74:
	;
	goto L75
L75:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+268)))
	if v306 != int32(1) {
		goto L60
	} else {
		goto L76
	}
L76:
	;
	F_pgstat_assoc_relation(m, v26)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	v312 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	v313 = *(*int64)(unsafe.Add(mBase, uint32(v312)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v312)+120)) = v313 + int64(1)
	if v311 != 0 {
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
	v348 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_StartReadBuffer[10])))
	if v348 != int32(1) {
		goto L59
	} else {
		goto L81
	}
L81:
	;
	v351 = int32(_a_F_StartReadBuffer_9)
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[11]))
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[12]))
	*(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[11])) = v353 + v355
	goto L59
L82:
	;
	m.G0 = v19 + int32(48)
	return v399
L83:
	;
	v399 = int32(0)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v367 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v367
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v367
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
	goto L86
L86:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _c_F_StartReadBuffer[13]))
	if v379 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v380 = F_AsyncReadBuffers(m, l0, v19)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L5
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v384 | int32(8)
	if l3&int32(2) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v382)
	v399 = v380
	goto L82
L91:
	;
	v399 = int32(1)
	goto L82
L92:
	;
	goto L93
L93:
	;
	v393 = int32(1)
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v397 = F_smgrprefetch(m, v394, v395, l2, v393)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v399 = v393
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
