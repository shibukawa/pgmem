package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckForSerializableConflictIn(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_CheckForSerializableConflictIn[0]))
	if v12 == int32(0) {
		m.G0 = v9 + int32(16)
		return
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		if base.Ui32(v15) < base.Ui32(int32(_a_F_CheckForSerializableConflictIn_0)) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+118)))
			if v19 == int32(116) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+108)))
				if v22&int32(8) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return
					} else {
						F_errcode(m, int32(16777220))
						mBase = m.M
						v77 = m.ExcPending
						if v77 != 0 {
							return
						} else {
							F_errmsg(m, int32(_a_F_CheckForSerializableConflictIn_1), int32(0))
							mBase = m.M
							v81 = m.ExcPending
							if v81 != 0 {
								return
							} else {
								F_errdetail_internal(m, int32(_a_F_CheckForSerializableConflictIn_2), int32(0))
								mBase = m.M
								v85 = m.ExcPending
								if v85 != 0 {
									return
								} else {
									F_errhint(m, int32(_a_F_CheckForSerializableConflictIn_3), int32(0))
									mBase = m.M
									v89 = m.ExcPending
									if v89 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_CheckForSerializableConflictIn_4), int32(_a_F_CheckForSerializableConflictIn_5), int32(_a_F_CheckForSerializableConflictIn_6))
										mBase = m.M
										v94 = m.ExcPending
										if v94 != 0 {
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
				} else {
					v26 = int32(1)
					*(*uint8)(unsafe.Add(mBase, _c_F_CheckForSerializableConflictIn[1])) = uint8(v26)
					if l1 != 0 {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v15
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v28
						v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
						v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v31 | v32<<(uint(int32(16))%32)
						v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v37
						F_CheckTargetForConflictsIn(m, v9)
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return
						} else {
							v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							v42 = v41
							if l2 != int32(-1) {
								v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v42
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
								F_CheckTargetForConflictsIn(m, v9)
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return
								} else {
									v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									v56 = v54
									v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
									*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
									F_CheckTargetForConflictsIn(m, v9)
									mBase = m.M
									v63 = m.ExcPending
									if v63 != 0 {
										return
									} else {
										m.G0 = v9 + int32(16)
										return
									}
								}
							} else {
								v56 = v42
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
								F_CheckTargetForConflictsIn(m, v9)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v9 + int32(16)
									return
								}
							}
						}
					} else {
						v42 = v15
						if l2 != int32(-1) {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = l2
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v42
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v46
							F_CheckTargetForConflictsIn(m, v9)
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
								v56 = v54
								v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
								F_CheckTargetForConflictsIn(m, v9)
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return
								} else {
									m.G0 = v9 + int32(16)
									return
								}
							}
						} else {
							v56 = v42
							v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(4294967295)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v56
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v57
							F_CheckTargetForConflictsIn(m, v9)
							mBase = m.M
							v63 = m.ExcPending
							if v63 != 0 {
								return
							} else {
								m.G0 = v9 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_CheckForSerializableConflictOutNeeded(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	v3 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_CheckForSerializableConflictOutNeeded[0]))
	if v6 == v3 {
		v56 = v3
		return v56
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		if v9 != 0 {
			v56 = v3
			return v56
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+108))
			if v10&int32(128) != 0 {
				F_ReleasePredicateLocks(m, int32(0), int32(1))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					v56 = v3
					return v56
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				if base.Ui32(v19) < base.Ui32(int32(_a_F_CheckForSerializableConflictOutNeeded_0)) {
					v56 = v3
					return v56
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+118)))
					if v23 == int32(116) {
						v56 = v3
						return v56
					} else {
						if v10&int32(8) == int32(0) {
							v56 = int32(1)
							return v56
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v34 = m.ExcPending
							if v34 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(16777220))
								mBase = m.M
								v37 = m.ExcPending
								if v37 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_CheckForSerializableConflictOutNeeded_1), int32(0))
									mBase = m.M
									v41 = m.ExcPending
									if v41 != 0 {
										return int32(0)
									} else {
										F_errdetail_internal(m, int32(_a_F_CheckForSerializableConflictOutNeeded_2), int32(0))
										mBase = m.M
										v45 = m.ExcPending
										if v45 != 0 {
											return int32(0)
										} else {
											F_errhint(m, int32(_a_F_CheckForSerializableConflictOutNeeded_3), int32(0))
											mBase = m.M
											v49 = m.ExcPending
											if v49 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_CheckForSerializableConflictOutNeeded_4), int32(3932), int32(_a_F_CheckForSerializableConflictOutNeeded_5))
												mBase = m.M
												v54 = m.ExcPending
												if v54 != 0 {
													return int32(0)
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
	}
}
func F_WaitForLSN(m *base.Module, l0 int32, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v108 int64
	_ = v108
	var v110 int32
	_ = v110
	var v113 int64
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int64
	_ = v204
	var v206 int32
	_ = v206
	var v207 int64
	_ = v207
	var v212 int64
	_ = v212
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v225 int64
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v283 int64
	_ = v283
	var v285 int32
	_ = v285
	var v286 int64
	_ = v286
	var v289 int64
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v314 int64
	_ = v314
	var v320 int64
	_ = v320
	var v329 int64
	_ = v329
	var v332 int32
	_ = v332
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int64
	_ = v398
	var v400 int32
	_ = v400
	var v401 int64
	_ = v401
	var v406 int64
	_ = v406
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[0]))
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLSN[2])))
	if v19 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_on_shmem_exit(m, int32(445), int64(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l2 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return int32(0)
L5:
	;
	v29 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLSN[2])) = uint8(v29)
	goto L3
L6:
	;
	v59 = int64(0)
	v60 = int32(17)
	goto L8
L7:
	;
	v39 = m.G0
	v40 = int32(16)
	v41 = v39 - v40
	m.G0 = v41
	F_gettimeofday(m, v41)
	mBase = m.M
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v41)))
	v45 = int64(*(*int32)(unsafe.Add(mBase, uint32(v41)+8)))
	m.G0 = v41 + v40
	goto L9
L8:
	;
	v63 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v65 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[0]))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[3]))
	v71 = F_LWLockAcquire(m, v67+int32(_a_F_WaitForLSN_0), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L10
	}
L9:
	;
	v59 = v45 + v44*int64(1000000) - int64(946684800000000) + base.I64_extend_i32_u(l2)*int64(1000)
	v60 = int32(25)
	goto L8
L10:
	;
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[0]))
	v77 = v63 + v65<<(uint(int32(5))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v77)+88)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v77)+80)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v77)+92)) = v74
	v82 = l0 * int32(12)
	v84 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	F_pairingheap_add(m, v82+v84+int32(32), v77+int32(100))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	v92 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v77)+96)) = uint8(v92)
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v96 = v95 + v82
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+40))
	if v97 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v96+int32(32))+8))
	v105 = *(*int64)(unsafe.Add(mBase, uint32(v102-int32(20))))
	v107 = v99
	v108 = v105
	goto L14
L13:
	;
	v107 = v95
	v108 = int64(-1)
	goto L14
L14:
	;
	v110 = l0 << (uint(int32(3)) % 32)
	v113 = base.AtomicRmwXchg64(m, v107+v110, int32(0), v108)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[3]))
	F_LWLockRelease(m, v115+int32(_a_F_WaitForLSN_0))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	goto L17
L16:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L4
	} else {
		goto L84
	}
L17:
	;
	v141 = F_GetCurrentLSNForWaitType(m, l0)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L4
	} else {
		goto L19
	}
L18:
	;
	v353 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v355 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[0]))
	v358 = v353 + v355<<(uint(int32(5))%32)
	v359 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v358)+96)))
	if v359 == int32(1) {
		goto L69
	} else {
		goto L70
	}
L19:
	;
	if base.Ui32(int32(2)) < base.Ui32(l0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if base.Ui64(l1) <= base.Ui64(v141) {
		goto L41
	} else {
		goto L42
	}
L21:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLSN[4])))
	if v147 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if v157 != 0 {
		goto L20
	} else {
		goto L26
	}
L23:
	;
	v152 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[5]))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+308))
	v155 = base.B2i32(v153 != int32(2))
	*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLSN[4])) = uint8(v155)
	v157 = v155
	goto L25
L24:
	;
	v157 = int32(0)
	goto L25
L25:
	;
	goto L22
L26:
	;
	v159 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v161 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[0]))
	v164 = v159 + v161<<(uint(int32(5))%32)
	v165 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164)+96)))
	if v165 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[3]))
	v173 = F_LWLockAcquire(m, v169+int32(_a_F_WaitForLSN_0), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v225 = F_GetCurrentLSNForWaitType(m, l0)
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L4
	} else {
		goto L39
	}
L30:
	;
	v176 = v164 + int32(80)
	v177 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v176)+16)))
	if v177 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v181 = l0 * int32(12)
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	F_pairingheap_remove(m, v181+v183+int32(32), v164+int32(100))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v217 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[3]))
	F_LWLockRelease(m, v217+int32(_a_F_WaitForLSN_0))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L4
	} else {
		goto L38
	}
L34:
	;
	v191 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v176)+16)) = uint8(v191)
	v194 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v195 = v181 + v194
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+40))
	if v196 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v198 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v195+int32(32))+8))
	v204 = *(*int64)(unsafe.Add(mBase, uint32(v201-int32(20))))
	v206 = v198
	v207 = v204
	goto L37
L36:
	;
	v206 = v194
	v207 = int64(-1)
	goto L37
L37:
	;
	v212 = base.AtomicRmwXchg64(m, v206+l0<<(uint(int32(3))%32), int32(0), v207)
	goto L33
L38:
	;
	goto L29
L39:
	;
	v227 = F_PromoteIsTriggered(m)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	return v227 ^ int32(1) | base.B2i32(base.Ui64(v225) < base.Ui64(l1))
L41:
	;
	goto L18
L42:
	;
	v236 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[6]))
	if v236 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15<<(uint(int32(5))%32)+v17+int32(96)))))
	if v239 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	goto L45
L47:
	;
	v243 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v245 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[0]))
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[3]))
	v251 = F_LWLockAcquire(m, v247+int32(_a_F_WaitForLSN_0), int32(0))
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L4
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	if l2 <= int32(0) {
		goto L57
	} else {
		goto L58
	}
L50:
	;
	v254 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[0]))
	v257 = v243 + v245<<(uint(int32(5))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v257)+88)) = l0
	*(*int64)(unsafe.Add(mBase, uint32(v257)+80)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v257)+92)) = v254
	v262 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	F_pairingheap_add(m, v262+v82+int32(32), v257+int32(100))
	mBase = m.M
	v269 = m.ExcPending
	if v269 != 0 {
		goto L4
	} else {
		goto L51
	}
L51:
	;
	v270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v257)+96)) = uint8(v270)
	v273 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v274 = v273 + v82
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+40))
	if v275 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v277 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v274+int32(32))+8))
	v283 = *(*int64)(unsafe.Add(mBase, uint32(v280-int32(20))))
	v285 = v277
	v286 = v283
	goto L54
L53:
	;
	v285 = v273
	v286 = int64(-1)
	goto L54
L54:
	;
	v289 = base.AtomicRmwXchg64(m, v285+v110, int32(0), v286)
	v291 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[3]))
	F_LWLockRelease(m, v291+int32(_a_F_WaitForLSN_0))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	goto L17
L56:
	;
	v337 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[7]))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_c_F_WaitForLSN[8])))
	v339 = F_WaitLatch(m, v337, v60, v335, v338)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L4
	} else {
		goto L66
	}
L57:
	;
	v335 = int32(-1)
	goto L56
L58:
	;
	goto L59
L59:
	;
	v300 = m.G0
	v301 = int32(16)
	v302 = v300 - v301
	m.G0 = v302
	F_gettimeofday(m, v302)
	mBase = m.M
	v305 = *(*int64)(unsafe.Add(mBase, uint32(v302)))
	v306 = int64(*(*int32)(unsafe.Add(mBase, uint32(v302)+8)))
	m.G0 = v302 + v301
	v314 = v306 + v305*int64(1000000) - int64(946684800000000)
	goto L60
L60:
	;
	if v59 <= v314 {
		v332 = int32(0)
		goto L62
	} else {
		goto L63
	}
L61:
	;
	if v332 <= int32(0) {
		goto L41
	} else {
		goto L65
	}
L62:
	;
	goto L61
L63:
	;
	v320 = v59 - v314
	if base.B2i32(int64(0) < v314)^base.B2i32(v320 < v59)|base.B2i32(int64(2147483646000) < v320) != 0 {
		v332 = int32(2147483647)
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v329 = base.I64_div_s(v320+int64(999), int64(1000))
	v332 = base.I32_wrap_i64(v329)
	goto L62
L65:
	;
	v335 = v332
	goto L56
L66:
	;
	if v339&int32(16) != 0 {
		goto L16
	} else {
		goto L67
	}
L67:
	;
	v344 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[7]))
	v345 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v344))) = v345
	v350 = base.AtomicRmwOr32(m, v345, int32(_a_F_WaitForLSN_1), v345)
	goto L68
L68:
	;
	goto L17
L69:
	;
	v363 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[3]))
	v367 = F_LWLockAcquire(m, v363+int32(_a_F_WaitForLSN_0), int32(0))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L4
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if base.Ui64(v141) < base.Ui64(l1) {
		goto L81
	} else {
		goto L82
	}
L72:
	;
	v370 = v358 + int32(80)
	v371 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v370)+16)))
	if v371 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v375 = l0 * int32(12)
	v377 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	F_pairingheap_remove(m, v375+v377+int32(32), v358+int32(100))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L4
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	v411 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[3]))
	F_LWLockRelease(m, v411+int32(_a_F_WaitForLSN_0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L4
	} else {
		goto L80
	}
L76:
	;
	v385 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v370)+16)) = uint8(v385)
	v388 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v389 = v375 + v388
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)+40))
	if v390 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v392 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLSN[1]))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v389+int32(32))+8))
	v398 = *(*int64)(unsafe.Add(mBase, uint32(v395-int32(20))))
	v400 = v392
	v401 = v398
	goto L79
L78:
	;
	v400 = v388
	v401 = int64(-1)
	goto L79
L79:
	;
	v406 = base.AtomicRmwXchg64(m, v400+l0<<(uint(int32(3))%32), int32(0), v401)
	goto L75
L80:
	;
	goto L71
L81:
	;
	v422 = int32(2)
	goto L83
L82:
	;
	v422 = int32(0)
	goto L83
L83:
	;
	return v422
L84:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(_a_F_WaitForLSN_2), int32(0))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	F_set_errcontext_domain(m, int32(0))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	F_errcontext_msg(m, int32(_a_F_WaitForLSN_3), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L4
	} else {
		goto L88
	}
L88:
	;
	F_errfinish(m, int32(_a_F_WaitForLSN_4), int32(565), int32(_a_F_WaitForLSN_5))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L4
	} else {
		goto L89
	}
L89:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WaitForLockersMultiple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
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
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v58 int64
	_ = v58
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v199 int64
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v251 int64
	_ = v251
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v291 int64
	_ = v291
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v444 int64
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v484 int32
	_ = v484
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	if l0 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v14 <= int32(0) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	m.G0 = v12 - int32(-64)
	return
L4:
	;
	v54 = v4
	v58 = int64(0)
	goto L6
L5:
	;
	if l2 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	if l2 != 0 {
		goto L19
	} else {
		goto L20
	}
L7:
	;
	v21 = v10 + int32(-48)
	goto L9
L8:
	;
	v21 = int32(0)
	goto L9
L9:
	;
	v26 = v4
	v27 = v4
	v28 = v4
	goto L10
L10:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v26<<(uint(int32(2))%32))))
	v36 = F_GetLockConflicts(m, v35, l1, v21)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v54 = v38
	v58 = base.I64_extend_i32_s(v43)
	goto L6
L12:
	;
	return
L13:
	;
	v38 = F_lappend(m, v27, v36)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if l2 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v42 = v40
	goto L17
L16:
	;
	v42 = int32(0)
	goto L17
L17:
	;
	v43 = v42 + v28
	v45 = v26 + int32(1)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v45 < v46 {
		v26 = v45
		v27 = v38
		v28 = v43
		goto L10
	} else {
		goto L18
	}
L18:
	;
	goto L11
L19:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[0]))
	if v62 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L21
L21:
	;
	if v54 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	goto L22
L24:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[1])))
	if v66&int32(1) == int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v71 = int32(_a_F_WaitForLockersMultiple_0)
	v73 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	v74 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v73 + v74
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v77 + v74
	v81 = int32(0)
	v83 = int32(_a_F_WaitForLockersMultiple_1)
	v84 = base.AtomicRmwOr32(m, v81, v83, v81)
	*(*int64)(unsafe.Add(mBase, uint32(v62+int32(24))+232)) = v58
	v92 = base.AtomicRmwOr32(m, v81, v83, v81)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	*(*int32)(unsafe.Add(mBase, uint32(v62))) = v93 + v74
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v99 - v74
	goto L23
L26:
	;
	if l2 != 0 {
		goto L62
	} else {
		goto L63
	}
L27:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v105 <= int32(0) {
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v108 = int32(0)
	v111 = v108
	v116 = v108
	goto L29
L29:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v54)+12))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v119+v111<<(uint(int32(2))%32))))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v123)+4))
	if v124 != 0 {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L26
L31:
	;
	v129 = v123
	v131 = v116
	goto L34
L32:
	;
	v269 = v116
	goto L33
L33:
	;
	v273 = v111 + int32(1)
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v54)+4))
	if v273 < v274 {
		v111 = v273
		v116 = v269
		goto L29
	} else {
		goto L61
	}
L34:
	;
	if l2 != 0 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v269 = v259
	goto L33
L36:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v129)+12))
	if v260 != 0 {
		v129 = v129 + int32(8)
		v131 = v259
		goto L34
	} else {
		goto L60
	}
L37:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v135 = int32(0)
	if v134 < v135 {
		v153 = v135
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L39
L39:
	;
	v251 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v251
	v256 = F_VirtualXactLock(m, v10+int32(-56), int32(1))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L12
	} else {
		goto L59
	}
L40:
	;
	if v153 != 0 {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v141 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[3]))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v141)+16))
	if base.Ui32(v142) <= base.Ui32(v134) {
		v153 = int32(0)
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v147 = v144 + v134*int32(768)
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v147)+12))
	if v149 != 0 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v150 = v147
	goto L46
L45:
	;
	v150 = int32(0)
	goto L46
L46:
	;
	v153 = v150
	goto L41
L47:
	;
	v155 = int64(*(*int32)(unsafe.Add(mBase, uint32(v153)+12)))
	v158 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[0]))
	if v158 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L48:
	;
	goto L49
L49:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(v129)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v199
	v202 = F_VirtualXactLock(m, v12, int32(1))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L12
	} else {
		goto L54
	}
L50:
	;
	goto L49
L51:
	;
	goto L50
L52:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[1])))
	if v162&int32(1) == int32(0) {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v167 = int32(_a_F_WaitForLockersMultiple_0)
	v169 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	v170 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v169 + v170
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v173 + v170
	v177 = int32(0)
	v179 = int32(_a_F_WaitForLockersMultiple_1)
	v180 = base.AtomicRmwOr32(m, v177, v179, v177)
	*(*int64)(unsafe.Add(mBase, uint32(v158+int32(40))+232)) = v155
	v188 = base.AtomicRmwOr32(m, v177, v179, v177)
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	*(*int32)(unsafe.Add(mBase, uint32(v158))) = v189 + v170
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v195 - v170
	goto L51
L54:
	;
	v206 = v131 + int32(1)
	v210 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[0]))
	if v210 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	v259 = v206
	goto L36
L56:
	;
	goto L55
L57:
	;
	v214 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[1])))
	if v214&int32(1) == int32(0) {
		goto L56
	} else {
		goto L58
	}
L58:
	;
	v219 = int32(_a_F_WaitForLockersMultiple_0)
	v221 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	v222 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v221 + v222
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v225 + v222
	v229 = int32(0)
	v231 = int32(_a_F_WaitForLockersMultiple_1)
	v232 = base.AtomicRmwOr32(m, v229, v231, v229)
	*(*int64)(unsafe.Add(mBase, uint32(v210+int32(32))+232)) = base.I64_extend_i32_s(v206)
	v240 = base.AtomicRmwOr32(m, v229, v231, v229)
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v210)))
	*(*int32)(unsafe.Add(mBase, uint32(v210))) = v241 + v222
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v247 - v222
	goto L56
L59:
	;
	v259 = v131
	goto L36
L60:
	;
	goto L35
L61:
	;
	goto L30
L62:
	;
	v286 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[4]))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v286
	v289 = *(*int64)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+48)) = v289
	v291 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v291
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v291
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v291
	goto L67
L63:
	;
	goto L64
L64:
	;
	F_list_free_deep(m, v54)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L12
	} else {
		goto L82
	}
L65:
	;
	goto L64
L66:
	;
	goto L65
L67:
	;
	v311 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[0]))
	if v311 == int32(0) {
		goto L66
	} else {
		goto L68
	}
L68:
	;
	v315 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[1])))
	if v315&int32(1) == int32(0) {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v320 = int32(_a_F_WaitForLockersMultiple_0)
	v322 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	v323 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v322 + v323
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v326 + v323
	v330 = int32(0)
	v333 = base.AtomicRmwOr32(m, v330, int32(_a_F_WaitForLockersMultiple_1), v330)
	goto L71
L70:
	;
	v460 = int32(0)
	v463 = base.AtomicRmwOr32(m, v460, int32(_a_F_WaitForLockersMultiple_1), v460)
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v311)))
	v465 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v311))) = v464 + v465
	v468 = int32(_a_F_WaitForLockersMultiple_0)
	v470 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_WaitForLockersMultiple[2])) = v470 - v465
	goto L66
L71:
	;
	goto L73
L73:
	;
	goto L74
L74:
	;
	v425 = int32(0)
	v428 = int32(0)
	goto L79
L79:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(-16)+v428<<(uint(int32(2))%32))))
	v438 = int32(3)
	v444 = *(*int64)(unsafe.Add(mBase, uint32(v10+int32(-48)+v428<<(uint(v438)%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v311+int32(232)+v437<<(uint(v438)%32)))) = v444
	v446 = int32(1)
	v449 = v425 + v446
	if v449 != int32(3) {
		v425 = v449
		v428 = v428 + v446
		goto L79
	} else {
		goto L81
	}
L80:
	;
	goto L70
L81:
	;
	goto L80
L82:
	;
	goto L3
}
func F_WaitForParallelWorkersToExit(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v2 < v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(22), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L15
	}
L2:
	;
	v11 = v6
	v12 = v2
	goto L5
L3:
	;
	goto L4
L4:
	;
	return
L5:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	if v14 == int32(0) {
		v38 = v11
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L4
L7:
	;
	v41 = v12 + int32(1)
	if v41 < v38 {
		v11 = v38
		v12 = v41
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v18 = v12 << (uint(int32(3)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v14+v18)))
	if v20 == int32(0) {
		v38 = v11
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v23 = F_WaitForBackgroundWorkerShutdown(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	if v23 == int32(3) {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+v18)))
	F_pfree(m, v29)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v32+v18))) = int32(0)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v38 = v36
	goto L7
L14:
	;
	goto L6
L15:
	;
	F_errcode(m, int32(16908741))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_WaitForParallelWorkersToExit_0), int32(0))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L10
	} else {
		goto L17
	}
L17:
	;
	F_errfinish(m, int32(_a_F_WaitForParallelWorkersToExit_1), int32(942), int32(_a_F_WaitForParallelWorkersToExit_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_WaitForProcSignalBarrier(m *base.Module, l0 int64) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int64
	_ = v41
	var v44 int64
	_ = v44
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v82 int64
	_ = v82
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v12 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v12 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = l0
	F_errmsg_internal(m, int32(_a_F_WaitForProcSignalBarrier_0), v8+int32(32))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForProcSignalBarrier[0]))
	v28 = v26 + int32(37)
	if int32(0) <= v28 {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	F_errfinish(m, int32(_a_F_WaitForProcSignalBarrier_1), int32(442), int32(_a_F_WaitForProcSignalBarrier_2))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	goto L5
L8:
	;
	v34 = v28
	goto L11
L9:
	;
	goto L10
L10:
	;
	v102 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L28
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _c_F_WaitForProcSignalBarrier[1]))
	v40 = v37 + v34*int32(112)
	v41 = int64(0)
	v44 = base.AtomicRmwCmpxchg64(m, v40, int32(96), v41, v41)
	if base.Ui64(v44) < base.Ui64(l0) {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	goto L10
L13:
	;
	v47 = v40 + int32(8)
	goto L16
L14:
	;
	goto L15
L15:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L26
	}
L16:
	;
	v57 = F_ConditionVariableTimedSleep(m, v40+int32(108), int32(_a_F_WaitForProcSignalBarrier_3), int32(134217770))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L15
L18:
	;
	v79 = int64(0)
	v82 = base.AtomicRmwCmpxchg64(m, v47, int32(88), v79, v79)
	if base.Ui64(v82) < base.Ui64(l0) {
		goto L16
	} else {
		goto L25
	}
L19:
	;
	if v57 == int32(0) {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v63 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v63 == int32(0) {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v47)))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v67
	F_errmsg(m, int32(_a_F_WaitForProcSignalBarrier_4), v8+int32(16))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F_errfinish(m, int32(_a_F_WaitForProcSignalBarrier_1), int32(463), int32(_a_F_WaitForProcSignalBarrier_2))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	goto L18
L25:
	;
	goto L17
L26:
	;
	if int32(0) < v34 {
		v34 = v34 - int32(1)
		goto L11
	} else {
		goto L27
	}
L27:
	;
	goto L12
L28:
	;
	if v102 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = l0
	F_errmsg_internal(m, int32(_a_F_WaitForProcSignalBarrier_5), v8)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v113 = int32(0)
	v116 = base.AtomicRmwOr32(m, v113, int32(_a_F_WaitForProcSignalBarrier_6), v113)
	m.G0 = v8 + int32(48)
	return
L32:
	;
	F_errfinish(m, int32(_a_F_WaitForProcSignalBarrier_1), int32(472), int32(_a_F_WaitForProcSignalBarrier_2))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	goto L31
}
