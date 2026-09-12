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
		v5 = *(*int32)(unsafe.Add(mBase, _consts[8]))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v5+(l0^int32(-1))<<(uint(int32(6))%32))+16))
		return v11
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
		v9 = *(*int32)(unsafe.Add(mBase, _consts[8]))
		v22 = v9 + (l0^int32(-1))<<(uint(int32(6))%32)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, _consts[9]))
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
	if v5 == int32(1583088) {
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
				v24 = v17 & int32(65533)
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
			v24 = v17 & int32(65533)
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
			F_errmsg_internal(m, int32(82622), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return
			} else {
				F_errfinish(m, int32(477412), int32(1594), int32(370893))
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v171 int32
	_ = v171
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v203 int64
	_ = v203
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v257 int32
	_ = v257
	var v259 int64
	_ = v259
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
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
	var v296 int64
	_ = v296
	var v304 int32
	_ = v304
	var v305 int64
	_ = v305
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int64
	_ = v316
	var v334 int32
	_ = v334
	var v340 int64
	_ = v340
	var v346 int64
	_ = v346
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v394 int32
	_ = v394
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
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
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v47 = F_IOContextForStrategy(m, v46)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
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
		v270 = v21
		v273 = v37
		v276 = v36
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v40 = int32(4353512)
	v42 = *(*int64)(unsafe.Add(mBase, _consts[104]))
	*(*int64)(unsafe.Add(mBase, _consts[104])) = v42 + int64(1)
	v269 = v32
	v270 = v21
	v273 = v37
	v276 = v36
	goto L1
L8:
	;
	v50 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	F_ResourceOwnerEnlarge(m, v50)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L9
	}
L9:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v55
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v59
	v65 = F_BufTableHashCode(m, v19+int32(4))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	v68 = *(*int32)(unsafe.Add(mBase, _consts[2]))
	v75 = v68 + v65&int32(127)<<(uint(int32(7))%32) + int32(6912)
	v77 = F_LWLockAcquire(m, v75, int32(1))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v81 = F_BufTableLookup(m, v19+int32(4), v65)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L5
	} else {
		goto L14
	}
L13:
	;
	v257 = int32(4353480)
	v259 = *(*int64)(unsafe.Add(mBase, _consts[100]))
	*(*int64)(unsafe.Add(mBase, _consts[100])) = v259 + int64(1)
	v269 = v251
	v270 = int32(0)
	v273 = int32(1)
	v276 = v47
	goto L1
L14:
	;
	if int32(0) <= v81 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v86 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v89 = v86 + v81<<(uint(int32(6))%32)
	v90 = F_PinBuffer(m, v89, v46)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L5
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	F_LWLockRelease(m, v75)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L5
	} else {
		goto L21
	}
L18:
	;
	F_LWLockRelease(m, v75)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L5
	} else {
		goto L19
	}
L19:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v90)
	if v90 != 0 {
		v251 = v89
		goto L13
	} else {
		goto L20
	}
L20:
	;
	v269 = v89
	v270 = int32(0)
	v273 = int32(0)
	v276 = v47
	goto L1
L21:
	;
	v98 = F_GetVictimBuffer(m, v46, v47)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L5
	} else {
		goto L22
	}
L22:
	;
	v101 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v103 = F_LWLockAcquire(m, v75, int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L5
	} else {
		goto L23
	}
L23:
	;
	v107 = v101 + v98<<(uint(int32(6))%32)
	v109 = v107 + int32(-64)
	v113 = v107 - int32(44)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	v115 = F_BufTableInsert(m, v19+int32(4), v65, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L5
	} else {
		goto L24
	}
L24:
	;
	if v115 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = int32(220922)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = int32(6259)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = int32(478364)
	v125 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v125
	*(*int64)(unsafe.Add(mBase, uint32(v19)+24)) = int64(0)
	v131 = v107 - int32(40)
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v133 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v132 | v133
	if v132&v133 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	goto L27
L27:
	;
	v229 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v113)))
	F_ResourceOwnerForget(m, v229, v230+int32(1), int32(1593376))
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L5
	} else {
		goto L53
	}
L28:
	;
	goto L31
L29:
	;
	v171 = v132
	goto L30
L30:
	;
	v183 = int32(4074876)
	v184 = *(*int32)(unsafe.Add(mBase, _consts[736]))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v19+int32(24))+8))
	if v186 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L31:
	;
	F_perform_spin_delay(m, v19+int32(24))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
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
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
	v159 = int32(4194304)
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v158 | v159
	if v158&v159 != 0 {
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v203 = *(*int64)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int64)(unsafe.Add(mBase, uint32(v109))) = v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v109)+16)) = v205
	v207 = *(*int64)(unsafe.Add(mBase, uint32(v19)+12))
	*(*int64)(unsafe.Add(mBase, uint32(v109)+8)) = v207
	v211 = int32(-2113667072)
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
	*(*int32)(unsafe.Add(mBase, _consts[736])) = v201
	goto L36
L38:
	;
	if int32(999) < v184 {
		goto L36
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	if v184 < int32(11) {
		goto L36
	} else {
		goto L45
	}
L41:
	;
	v191 = int32(900)
	if v191 <= v184 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v194 = v191
	goto L44
L43:
	;
	v194 = v184
	goto L44
L44:
	;
	v201 = v194 + int32(100)
	goto L37
L45:
	;
	v201 = v184 - int32(1)
	goto L37
L46:
	;
	v216 = v211
	goto L48
L47:
	;
	v216 = int32(33816576)
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
	v219 = v211
	goto L51
L50:
	;
	v219 = v216
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v131))) = v171&int32(-38010881) | v219
	F_LWLockRelease(m, v75)
	mBase = m.M
	v223 = m.ExcPending
	if v223 != 0 {
		goto L5
	} else {
		goto L52
	}
L52:
	;
	v224 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v224)
	v269 = v109
	v270 = v224
	v273 = v125
	v276 = v47
	goto L1
L53:
	;
	F_UnpinBufferNoOwner(m, v109)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L5
	} else {
		goto L54
	}
L54:
	;
	F_StrategyFreeBuffer(m, v109)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L5
	} else {
		goto L55
	}
L55:
	;
	v241 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v244 = v241 + v115<<(uint(int32(6))%32)
	v245 = F_PinBuffer(m, v244, v46)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	F_LWLockRelease(m, v75)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L5
	} else {
		goto L57
	}
L57:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)) = uint8(v245)
	if v245 != 0 {
		v251 = v244
		goto L13
	} else {
		goto L58
	}
L58:
	;
	v269 = v244
	v270 = int32(0)
	v273 = int32(0)
	v276 = v47
	goto L1
L59:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(v269)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v371 + int32(1)
	v375 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	if v375 != 0 {
		goto L83
	} else {
		goto L84
	}
L60:
	;
	v334 = v270*int32(320) + v276<<(uint(int32(6))%32)
	v340 = *(*int64)(unsafe.Add(mBase, uint32(v334)+uint32(_consts[737])))
	*(*int64)(unsafe.Add(mBase, uint32(v334)+uint32(_consts[737]))) = v340 + int64(1)
	v346 = *(*int64)(unsafe.Add(mBase, uint32(v334)+uint32(_consts[738])))
	*(*int64)(unsafe.Add(mBase, uint32(v334)+uint32(_consts[738]))) = v346
	v348 = int32(1)
	F_pgstat_count_backend_io_op(m, v270, v276, int32(2), v348, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _consts[241])) = uint8(v348)
	*(*uint8)(unsafe.Add(mBase, _consts[242])) = uint8(v348)
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
	if v273 == int32(0) {
		goto L59
	} else {
		goto L79
	}
L64:
	;
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	if v304 != 0 {
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
	v296 = *(*int64)(unsafe.Add(mBase, uint32(v280)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v280)+112)) = v296 + int64(1)
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
	if v288&int32(1) != 0 {
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
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v304)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v304)+120)) = v305 + int64(1)
	goto L60
L74:
	;
	goto L75
L75:
	;
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+268)))
	if v309 != int32(1) {
		goto L60
	} else {
		goto L76
	}
L76:
	;
	F_pgstat_assoc_relation(m, v26)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L5
	} else {
		goto L77
	}
L77:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+24)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v26)+272))
	v316 = *(*int64)(unsafe.Add(mBase, uint32(v315)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v315)+120)) = v316 + int64(1)
	if v314&int32(1) != 0 {
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
	v358 = int32(*(*uint8)(unsafe.Add(mBase, _consts[739])))
	if v358 != int32(1) {
		goto L59
	} else {
		goto L81
	}
L81:
	;
	v361 = int32(4449936)
	v363 = *(*int32)(unsafe.Add(mBase, _consts[740]))
	v365 = *(*int32)(unsafe.Add(mBase, _consts[741]))
	*(*int32)(unsafe.Add(mBase, _consts[740])) = v363 + v365
	goto L59
L82:
	;
	m.G0 = v19 + int32(48)
	return v409
L83:
	;
	v409 = int32(0)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v377 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v377
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v377
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(l0+int32(36)))) = int32(-1)
	goto L86
L86:
	;
	v389 = *(*int32)(unsafe.Add(mBase, _consts[734]))
	if v389 != 0 {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v390 = F_AsyncReadBuffers(m, l0, v19)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L5
	} else {
		goto L90
	}
L88:
	;
	goto L89
L89:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v394 | int32(8)
	if l3&int32(2) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L90:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+32)) = uint16(v392)
	v409 = v390
	goto L82
L91:
	;
	v409 = int32(1)
	goto L82
L92:
	;
	goto L93
L93:
	;
	v403 = int32(1)
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v407 = F_smgrprefetch(m, v404, v405, l2, v403)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L5
	} else {
		goto L94
	}
L94:
	;
	v409 = v403
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[258]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_ResourceOwnerForget(m, v3, v4+int32(1), int32(1593376))
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
