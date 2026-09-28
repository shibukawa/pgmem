package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"sync/atomic"
	"unsafe"
)

func F_BTreeTupleGetHeapTID(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	v4 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)))
	if v4&int32(_a_F_BTreeTupleGetHeapTID_0) == int32(0) {
		return l0
	} else {
		v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
		if v10&int32(_a_F_BTreeTupleGetHeapTID_0) == int32(0) {
			v15 = int32(0)
			if v10&int32(_a_F_BTreeTupleGetHeapTID_1) == v15 {
				v32 = v15
				return v32
			} else {
				return l0 + v4&int32(_a_F_BTreeTupleGetHeapTID_2) - int32(6)
			}
		} else {
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+2)))
			v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
			v32 = v26 + (l0 + v27<<(uint(int32(16))%32))
			return v32
		}
	}
}
func F_BTreeTupleGetHeapTIDCareful(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
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
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v13&int32(_a_F_BTreeTupleGetHeapTIDCareful_0) != 0 {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v18 = v16 & int32(_a_F_BTreeTupleGetHeapTIDCareful_0)
		if v18|base.B2i32(l2 == int32(0)) != 0 {
			if v18 != 0 {
				v49 = l2
			} else {
				v49 = int32(1)
			}
			if v49 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v115 = m.ExcPending
				if v115 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33557032))
					mBase = m.M
					v118 = m.ExcPending
					if v118 != 0 {
						return int32(0)
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+48))
						v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v121
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v120 + int32(4)
						F_errmsg_internal(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_1), v10+int32(16))
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3555), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				if v18 == int32(0) {
					if v16&int32(_a_F_BTreeTupleGetHeapTIDCareful_4) == int32(0) {
						v77 = int32(0)
						v80 = int32(1)
					} else {
						v71 = l1 + v13&int32(_a_F_BTreeTupleGetHeapTIDCareful_5) - int32(6)
						v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
						v77 = v71
						v80 = base.B2i32(v74 == int32(0))
					}
				} else {
					v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
					v66 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
					v71 = v65 + (l1 + v66<<(uint(int32(16))%32))
					v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
					v77 = v71
					v80 = base.B2i32(v74 == int32(0))
				}
				if v80 != 0 {
					v81 = l2
				} else {
					v81 = int32(0)
				}
				if v81 == int32(0) {
					m.G0 = v10 + int32(48)
					return v77
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(33557032))
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
							v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v97
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96 + int32(4)
							F_errmsg(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_6), v10)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3563), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)+48))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v33
					*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v32 + int32(4)
					F_errmsg_internal(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_7), v10+int32(32))
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3548), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
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
	} else {
		if l2 != 0 {
			v71 = l1
			v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v71)+4)))
			v77 = v71
			v80 = base.B2i32(v74 == int32(0))
			if v80 != 0 {
				v81 = l2
			} else {
				v81 = int32(0)
			}
			if v81 == int32(0) {
				m.G0 = v10 + int32(48)
				return v77
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v91 = m.ExcPending
				if v91 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(33557032))
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+48))
						v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v97
						*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v96 + int32(4)
						F_errmsg(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_6), v10)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3563), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(33557032))
				mBase = m.M
				v118 = m.ExcPending
				if v118 != 0 {
					return int32(0)
				} else {
					v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v120 = *(*int32)(unsafe.Add(mBase, uint32(v119)+48))
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v121
					*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v120 + int32(4)
					F_errmsg_internal(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_1), v10+int32(16))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_BTreeTupleGetHeapTIDCareful_2), int32(3555), int32(_a_F_BTreeTupleGetHeapTIDCareful_3))
						mBase = m.M
						v135 = m.ExcPending
						if v135 != 0 {
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
func F_BeginImplicitTransactionBlock(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_BeginImplicitTransactionBlock[0]))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	if v4 == int32(1) {
		*(*int32)(unsafe.Add(mBase, uint32(v3)+24)) = int32(4)
	} else {
	}
	return
}
func F_BitvecInit(m *base.Module) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, _c_F_BitvecInit[0])) = int32(_a_F_BitvecInit_0)
	*(*int32)(unsafe.Add(mBase, _c_F_BitvecInit[1])) = int32(_a_F_BitvecInit_1)
	return
}
func F_BlockSampler_Next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v28 int64
	_ = v28
	var v29 int64
	_ = v29
	var v30 int64
	_ = v30
	var v51 float64
	_ = v51
	var v54 int32
	_ = v54
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 float64
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 float64
	_ = v77
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v96 int32
	_ = v96
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v11 = v9 - v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v14 = v12 - v13
	if base.Ui32(v11) < base.Ui32(v14) {
		v17 = l0 + int32(16)
		for {
			v28 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v17)+8))
			v30 = v28 ^ v29
			*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = base.I64_rotl(v30, int64(37))
			*(*int64)(unsafe.Add(mBase, uint32(v17))) = v30<<(uint(int64(16))%64) ^ base.I64_rotl(v28, int64(24)) ^ v30
			v51 = F_scalbn(m, base.F64_convert_i64_u(int64(base.Ui64(base.I64_rotl(v28*int64(5), int64(7))*int64(9))>>(uint(int64(12))%64))), int32(-52))
			mBase = m.M
			if base.F64_eq(v51, float64(0)) != 0 {
				continue
			} else {
				break
			}
			break
		}
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		v56 = base.F64_convert_i32_s(v11)
		v59 = base.F64_sub(float64(1), base.F64_div(v56, base.F64_convert_i32_u(v14)))
		if base.F64_gt(v59, v51) != 0 {
			v62 = v54
			v64 = v14
			v67 = v59
			for {
				v69 = int32(1)
				v70 = v62 + v69
				v73 = v64 - v69
				v77 = base.F64_mul(v67, base.F64_sub(float64(1), base.F64_div(v56, base.F64_convert_i32_u(v73))))
				if base.F64_lt(v51, v77) != 0 {
					v62 = v70
					v64 = v73
					v67 = v77
					continue
				} else {
					break
				}
				break
			}
			v80 = v70
		} else {
			v80 = v54
		}
		v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v89 = v80
		v90 = v87
	} else {
		v89 = v13
		v90 = v10
	}
	v96 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v89 + v96
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v90 + v96
	return v89
}
func F_BogusGetChunkSpace(m *base.Module, l0 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14206(m, l0, int32(_a_F_BogusGetChunkSpace_0), int32(335), int32(_a_F_BogusGetChunkSpace_1))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_BuildIndex_2(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int64
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v94 int32
	_ = v94
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
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 float64
	_ = v158
	var v161 float64
	_ = v161
	var v165 float64
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
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
	var v205 int32
	_ = v205
	var v217 float64
	_ = v217
	var v220 int32
	_ = v220
	var v222 float64
	_ = v222
	var v225 float64
	_ = v225
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 float64
	_ = v233
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 float64
	_ = v243
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 float64
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v335 int32
	_ = v335
	var v340 int32
	_ = v340
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v436 int32
	_ = v436
	var v452 int32
	_ = v452
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
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
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v754 int32
	_ = v754
	var v757 int64
	_ = v757
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v803 int32
	_ = v803
	var v806 int32
	_ = v806
	var v808 int32
	_ = v808
	var v819 int32
	_ = v819
	var v820 int32
	_ = v820
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v831 int32
	_ = v831
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v851 int32
	_ = v851
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v869 int32
	_ = v869
	var v871 int32
	_ = v871
	var v883 int32
	_ = v883
	var v885 int32
	_ = v885
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v892 int32
	_ = v892
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v903 int32
	_ = v903
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v959 int32
	_ = v959
	var v960 float64
	_ = v960
	var v962 float64
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v979 float64
	_ = v979
	var v980 int32
	_ = v980
	var v998 float64
	_ = v998
	var v1000 int32
	_ = v1000
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1026 int32
	_ = v1026
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1035 int32
	_ = v1035
	var v1039 int32
	_ = v1039
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1072 int32
	_ = v1072
	var v1077 float64
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1085 int32
	_ = v1085
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1118 int32
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1136 int32
	_ = v1136
	var v1149 int64
	_ = v1149
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1177 int32
	_ = v1177
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1204 int64
	_ = v1204
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1216 int32
	_ = v1216
	var v1224 int32
	_ = v1224
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1245 int64
	_ = v1245
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1278 int32
	_ = v1278
	var v1279 int32
	_ = v1279
	var v1285 int32
	_ = v1285
	var v1289 int32
	_ = v1289
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1313 int64
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1326 int32
	_ = v1326
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1334 int32
	_ = v1334
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1337 int32
	_ = v1337
	var v1339 int32
	_ = v1339
	var v1340 int32
	_ = v1340
	var v1344 int64
	_ = v1344
	var v1350 int32
	_ = v1350
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1374 int32
	_ = v1374
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1389 int32
	_ = v1389
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1400 int32
	_ = v1400
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1415 int32
	_ = v1415
	var v1422 int32
	_ = v1422
	var v1426 int32
	_ = v1426
	var v1431 int32
	_ = v1431
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1442 int32
	_ = v1442
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1452 int32
	_ = v1452
	var v1460 int32
	_ = v1460
	var v1465 int32
	_ = v1465
	v19 = m.G0
	v21 = v19 + int32(-64)
	m.G0 = v21
	F_InitBuildState_2(m, l3, l0, l1, l2)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v29 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v70 = int32(1)
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v71 == int32(0) {
		v94 = v70
		goto L7
	} else {
		goto L8
	}
L4:
	;
	goto L3
L5:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v33&int32(1) == int32(0) {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v38 = int32(_a_F_BuildIndex_2_0)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v41 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v40 + v41
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v44 + v41
	v48 = int32(0)
	v50 = int32(_a_F_BuildIndex_2_1)
	v51 = base.AtomicRmwOr32(m, v48, v50, v48)
	*(*int64)(unsafe.Add(mBase, uint32(v29+int32(80))+232)) = int64(2)
	v59 = base.AtomicRmwOr32(m, v48, v50, v48)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
	*(*int32)(unsafe.Add(mBase, uint32(v29))) = v60 + v41
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v66 - v41
	goto L4
L7:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l3)+164))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	v104 = F_mul_size(m, v94, (v99+int32(7))&int32(-8))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L17
	}
L8:
	;
	v75 = F_RelationGetNumberOfBlocksInFork(m, v71, int32(0))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v75 == int32(0) {
		v94 = v70
		goto L7
	} else {
		goto L10
	}
L10:
	;
	v81 = base.I64_extend_i32_u(v75) * int64(291)
	v82 = int32(_a_F_BuildIndex_2_2)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v85 = v83 * int32(50)
	if v85 <= v82 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v88 = v82
	goto L13
L12:
	;
	v88 = v85
	goto L13
L13:
	;
	v89 = base.I64_extend_i32_u(v88)
	if base.Ui64(v81) < base.Ui64(v89) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v91 = v81
	goto L16
L15:
	;
	v91 = v89
	goto L16
L16:
	;
	v94 = base.I32_wrap_i64(v91)
	goto L7
L17:
	;
	v106 = F_add_size(m, int32(20), v104)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v108 = F_add_size(m, v97, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+164)) = v108
	F_IvfflatCheckMemoryUsage(m, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l3)+76))
	v115 = F_VectorArrayInit(m, v94, v113, v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+64)) = v115
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v118 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(l3)+164))
	F_IvfflatKmeans(m, v359, v360, v361, v362, v363)
	mBase = m.M
	v365 = m.ExcPending
	if v365 != 0 {
		goto L1
	} else {
		goto L73
	}
L23:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v115)+4))
	v123 = F_RelationGetNumberOfBlocksInFork(m, v118, int32(0))
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l3)+144)) = int64(-4616189618054758400)
	*(*int64)(unsafe.Add(mBase, uint32(l3)+136)) = int64(0)
	v130 = l3 + int32(80)
	v132 = Fn14349(m, int64(32))
	mBase = m.M
	goto L25
L25:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v130)+8)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+4)) = v121
	*(*int32)(unsafe.Add(mBase, uint32(v130))) = v123
	F_pg_prng_seed(m, l3+int32(96), base.I64_extend_i32_u(v132))
	mBase = m.M
	goto L27
L26:
	;
	v150 = l3 + int32(120)
	v151 = F_pg_prng_uint32(m)
	mBase = m.M
	F_pg_prng_seed(m, v150, base.I64_extend_i32_u(v151))
	mBase = m.M
	goto L31
L27:
	;
	goto L29
L29:
	;
	goto L26
L30:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if base.Ui32(v167) < base.Ui32(v168) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	v158 = F_pg_prng_double(m, v150)
	mBase = m.M
	if base.F64_eq(v158, float64(0)) != 0 {
		goto L31
	} else {
		goto L33
	}
L32:
	;
	v161 = F_log(m, v158)
	mBase = m.M
	v165 = F_exp(m, base.F64_div(base.F64_neg(v161), base.F64_convert_i32_s(v121)))
	mBase = m.M
	*(*float64)(unsafe.Add(mBase, uint32(l3+int32(112)))) = v165
	goto L30
L33:
	;
	goto L32
L34:
	;
	if v174 != 0 {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v174 = base.B2i32(v170 < v171)
	goto L37
L36:
	;
	v174 = int32(0)
	goto L37
L37:
	;
	goto L34
L38:
	;
	goto L41
L39:
	;
	goto L40
L40:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l3)+56))
	if v307 != 0 {
		goto L62
	} else {
		goto L63
	}
L41:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v202 = v200 - v201
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v205 = v203 - v204
	if base.Ui32(v202) < base.Ui32(v205) {
		goto L44
	} else {
		goto L45
	}
L42:
	;
	goto L40
L43:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v271 = int32(0)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v268)+188))
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v277)+140))
	v279 = m.T0[v278].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v268, v269, v270, v271, v271, v271, v255, int32(1), int32(_a_F_BuildIndex_2_3), l3, v271)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L56
	}
L44:
	;
	goto L47
L45:
	;
	v255 = v204
	v256 = v201
	goto L46
L46:
	;
	v262 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v130)+8)) = v255 + v262
	*(*int32)(unsafe.Add(mBase, uint32(v130)+12)) = v256 + v262
	goto L43
L47:
	;
	v217 = F_pg_prng_double(m, l3+int32(96))
	mBase = m.M
	if base.F64_eq(v217, float64(0)) != 0 {
		goto L47
	} else {
		goto L49
	}
L48:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v222 = base.F64_convert_i32_s(v202)
	v225 = base.F64_sub(float64(1), base.F64_div(v222, base.F64_convert_i32_u(v205)))
	if base.F64_gt(v225, v217) != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	goto L48
L50:
	;
	v228 = v220
	v230 = v205
	v233 = v225
	goto L53
L51:
	;
	v246 = v220
	goto L52
L52:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v255 = v246
	v256 = v253
	goto L46
L53:
	;
	v235 = int32(1)
	v236 = v228 + v235
	v239 = v230 - v235
	v243 = base.F64_mul(v233, base.F64_sub(float64(1), base.F64_div(v222, base.F64_convert_i32_u(v239))))
	if base.F64_lt(v217, v243) != 0 {
		v228 = v236
		v230 = v239
		v233 = v243
		goto L53
	} else {
		goto L55
	}
L54:
	;
	v246 = v236
	goto L52
L55:
	;
	goto L54
L56:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	if base.Ui32(v281) < base.Ui32(v282) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v288 != 0 {
		goto L41
	} else {
		goto L61
	}
L58:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v130)+12))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v288 = base.B2i32(v284 < v285)
	goto L60
L59:
	;
	v288 = int32(0)
	goto L60
L60:
	;
	goto L57
L61:
	;
	goto L42
L62:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v309 = *(*int32)(unsafe.Add(mBase, uint32(l3)+60))
	v310 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	F_IvfflatNormVectors(m, v308, v309, v310, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v314)))
	v316 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v316 <= v315 {
		goto L22
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	v320 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	if v320 == int32(0) {
		goto L22
	} else {
		goto L68
	}
L68:
	;
	F_errmsg(m, int32(_a_F_BuildIndex_2_4), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v330 = F_errdetail(m, int32(_a_F_BuildIndex_2_5), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errhint(m, int32(_a_F_BuildIndex_2_6), int32(0))
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_7), int32(471), int32(_a_F_BuildIndex_2_8))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	goto L22
L73:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(l3)+64))
	F_VectorArrayFree(m, v366)
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v370 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v371 = F_HnswNewBuffer(m, l1, l4)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v371
	v375 = v19 + int32(-4)
	v377 = v19 + int32(-8)
	v379 = v19 + int32(-12)
	F_IvfflatInitRegisterPage(m, l1, v375, v377, v379)
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	*(*uint16)(unsafe.Add(mBase, uint32(v382)+34)) = uint16(v370)
	*(*uint16)(unsafe.Add(mBase, uint32(v382)+32)) = uint16(v369)
	*(*int64)(unsafe.Add(mBase, uint32(v382)+24)) = int64(4316983719)
	v387 = int32(36)
	*(*uint16)(unsafe.Add(mBase, uint32(v382)+12)) = uint16(v387)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	F_IvfflatCommitBuffer(m, v389, v390)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)+12))
	v397 = F_add_size(m, int32(8), v396)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v402 = (v397 + int32(7)) & int32(-8)
	v403 = F_palloc0(m, v402)
	mBase = m.M
	v404 = m.ExcPending
	if v404 != 0 {
		goto L1
	} else {
		goto L79
	}
L79:
	;
	v405 = F_HnswNewBuffer(m, l1, l4)
	mBase = m.M
	v406 = m.ExcPending
	if v406 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = v405
	F_IvfflatInitRegisterPage(m, l1, v375, v377, v379)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	if int32(0) < v393 {
		goto L85
	} else {
		goto L86
	}
L82:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L1
	} else {
		goto L316
	}
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1435 = m.ExcPending
	if v1435 != 0 {
		goto L1
	} else {
		goto L313
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L1
	} else {
		goto L310
	}
L85:
	;
	v414 = v403 + v402
	v416 = v403 + int32(4)
	if base.Ui32(v416) < base.Ui32(v414) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v570 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	F_IvfflatCommitBuffer(m, v569, v570)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L132
	}
L88:
	;
	v418 = v414
	goto L90
L89:
	;
	v418 = v416
	goto L90
L90:
	;
	v425 = v403 & int32(3)
	if v425 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v426 = v402
	goto L93
L92:
	;
	v426 = (v403^int32(-1)+v418)&int32(-4) + int32(4)
	goto L93
L93:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v402) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v429 = v402
	goto L96
L95:
	;
	v429 = v426
	goto L96
L96:
	;
	v436 = int32(0)
	goto L97
L97:
	;
	v452 = int32(0)
	if base.B2i32(v425|v402 == v452)|base.B2i32(v429 == v452) == v452 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	goto L87
L99:
	;
	base.MemoryFill(m, v403, int32(0), v429)
	goto L101
L100:
	;
	goto L101
L101:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v403))) = int64(-1)
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v395)+4))
	if v463 <= v436 {
		goto L84
	} else {
		goto L102
	}
L102:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v395)+16))
	v466 = *(*int32)(unsafe.Add(mBase, uint32(v395)+12))
	v468 = v465 + v466*v436
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468))))
	if v469 == int32(1) {
		goto L104
	} else {
		goto L105
	}
L103:
	;
	if v494 != 0 {
		goto L114
	} else {
		goto L115
	}
L104:
	;
	v473 = int32(18)
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468)+1)))
	if v475 == v473 {
		goto L107
	} else {
		goto L108
	}
L105:
	;
	goto L106
L106:
	;
	v486 = int32(1)
	if v469&v486 != 0 {
		v494 = int32(base.Ui32(v469) >> (uint(v486) % 32))
		goto L103
	} else {
		goto L113
	}
L107:
	;
	v478 = v473
	goto L109
L108:
	;
	v478 = int32(2)
	goto L109
L109:
	;
	if base.Ui32((v475-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v485 = int32(6)
	goto L112
L111:
	;
	v485 = v478
	goto L112
L112:
	;
	v494 = v485
	goto L103
L113:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v468)))
	v494 = int32(base.Ui32(v490) >> (uint(int32(2)) % 32))
	goto L103
L114:
	;
	base.MemoryCopy(m, v403+int32(8), v468, v494)
	goto L116
L115:
	;
	goto L116
L116:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v497 = int32(4)
	v498 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v496)+14)))
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v496)+12)))
	v500 = v498 - v499
	if v500 <= v497 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if base.Ui32(v503-int32(4)) < base.Ui32(v402) {
		goto L121
	} else {
		goto L122
	}
L118:
	;
	v503 = v497
	goto L120
L119:
	;
	v503 = v500
	goto L120
L120:
	;
	goto L117
L121:
	;
	F_IvfflatAppendPage(m, l1, v19+int32(-4), v19+int32(-8), v19+int32(-12), l4)
	mBase = m.M
	v514 = m.ExcPending
	if v514 != 0 {
		goto L1
	} else {
		goto L124
	}
L122:
	;
	goto L123
L123:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v516 = int32(0)
	v518 = F_PageAddItemExtended(m, v515, v403, v402, v516, v516)
	mBase = m.M
	v519 = m.ExcPending
	if v519 != 0 {
		goto L1
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	if v518 == int32(0) {
		goto L83
	} else {
		goto L126
	}
L126:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v522 < int32(0) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v545 = v542 + v436<<(uint(int32(3))%32)
	*(*uint16)(unsafe.Add(mBase, uint32(v545)+4)) = uint16(v518)
	*(*int32)(unsafe.Add(mBase, uint32(v545))) = v541
	v549 = v436 + int32(1)
	if v549 != v393 {
		v436 = v549
		goto L97
	} else {
		goto L131
	}
L128:
	;
	v526 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[3]))
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v526+(v522^int32(-1))*int32(56))+16))
	v541 = v532
	goto L127
L129:
	;
	goto L130
L130:
	;
	v534 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[4]))
	v535 = int32(56)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v534+v522*v535-v535)+16))
	v541 = v540
	goto L127
L131:
	;
	goto L98
L132:
	;
	F_pfree(m, v403)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	v579 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v579 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L134:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v620 == int32(0) {
		goto L138
	} else {
		goto L139
	}
L135:
	;
	goto L134
L136:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v583&int32(1) == int32(0) {
		goto L135
	} else {
		goto L137
	}
L137:
	;
	v588 = int32(_a_F_BuildIndex_2_0)
	v590 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v591 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v590 + v591
	v594 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	*(*int32)(unsafe.Add(mBase, uint32(v579))) = v594 + v591
	v598 = int32(0)
	v600 = int32(_a_F_BuildIndex_2_1)
	v601 = base.AtomicRmwOr32(m, v598, v600, v598)
	*(*int64)(unsafe.Add(mBase, uint32(v579+int32(80))+232)) = int64(3)
	v609 = base.AtomicRmwOr32(m, v598, v600, v598)
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	*(*int32)(unsafe.Add(mBase, uint32(v579))) = v610 + v591
	v616 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v616 - v591
	goto L135
L138:
	;
	v883 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	if v883 != 0 {
		goto L212
	} else {
		goto L213
	}
L139:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v620)+56))
	v624 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v624)+56))
	v626 = F_plan_create_index_workers(m, v623, v625)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	if v626 <= int32(0) {
		goto L138
	} else {
		goto L141
	}
L141:
	;
	v631 = v626 + int32(1)
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v633 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v632)+121)))
	v635 = F_palloc0(m, int32(24))
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	v639 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v639)+72)) = v640 + int32(1)
	goto L143
L143:
	;
	v647 = F_CreateParallelContext(m, int32(_a_F_BuildIndex_2_9), int32(_a_F_BuildIndex_2_10), v626)
	mBase = m.M
	v648 = m.ExcPending
	if v648 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	if v633 == int32(1) {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	v651 = F_GetTransactionSnapshot(m)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L1
	} else {
		goto L148
	}
L146:
	;
	v655 = int32(_a_F_BuildIndex_2_11)
	goto L147
L147:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v658 = F_table_parallelscan_estimate(m, v657, v655)
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L1
	} else {
		goto L150
	}
L148:
	;
	v653 = F_RegisterSnapshot(m, v651)
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v655 = v653
	goto L147
L150:
	;
	v660 = F_add_size(m, int32(64), v658)
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v647)+36))
	v667 = F_add_size(m, v662, (v660+int32(31))&int32(-32))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v647)+36)) = v667
	v670 = F_tuplesort_estimate_shared(m, v631)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v647)+36))
	v677 = F_add_size(m, v672, (v670+int32(31))&int32(-32))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v647)+36)) = v677
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v680)+4))
	v682 = *(*int32)(unsafe.Add(mBase, uint32(v680)+12))
	v683 = v681 * v682
	v688 = F_add_size(m, v677, (v683+int32(31))&int32(-32))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v647)+36)) = v688
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v647)+40))
	v693 = F_add_size(m, v691, int32(3))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v647)+40)) = v693
	v697 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[6]))
	if v697 != 0 {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v698 = *(*int32)(unsafe.Add(mBase, uint32(v647)+36))
	v699 = F_strlen(m, v697)
	mBase = m.M
	v704 = F_add_size(m, v698, v699&int32(-32)+int32(32))
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L1
	} else {
		goto L160
	}
L158:
	;
	v716 = int32(1)
	goto L159
L159:
	;
	F_InitializeParallelDSM(m, v647)
	mBase = m.M
	v718 = m.ExcPending
	if v718 != 0 {
		goto L1
	} else {
		goto L162
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v647)+36)) = v704
	v707 = *(*int32)(unsafe.Add(mBase, uint32(v647)+40))
	v709 = F_add_size(m, v707, int32(1))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v647)+40)) = v709
	v716 = v699 + int32(1)
	goto L159
L162:
	;
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v647)+44))
	if v719 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v655)))
	if v722 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L165
L165:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v647)+52))
	v737 = F_shm_toc_allocate(m, v736, v660)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L1
	} else {
		goto L172
	}
L166:
	;
	F_UnregisterSnapshot(m, v655)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L1
	} else {
		goto L169
	}
L167:
	;
	goto L168
L168:
	;
	F_DestroyParallelContext(m, v647)
	mBase = m.M
	v728 = m.ExcPending
	if v728 != 0 {
		goto L1
	} else {
		goto L170
	}
L169:
	;
	goto L168
L170:
	;
	v731 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v731)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v731)+72)) = v732 - int32(1)
	goto L171
L171:
	;
	goto L138
L172:
	;
	v739 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v740 = *(*int32)(unsafe.Add(mBase, uint32(v739)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v737))) = v740
	v742 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v742)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v737)+12)) = v631
	*(*uint8)(unsafe.Add(mBase, uint32(v737)+8)) = uint8(v633)
	*(*int32)(unsafe.Add(mBase, uint32(v737)+4)) = v743
	v748 = v737 + int32(16)
	v749 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v748))), uint32(v749))
	*(*int64)(unsafe.Add(mBase, uint32(v748)+4)) = int64(-1)
	goto L173
L173:
	;
	v754 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v737)+28)), uint32(v754))
	v757 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v737)+40)) = v757
	*(*int32)(unsafe.Add(mBase, uint32(v737)+32)) = v754
	*(*int64)(unsafe.Add(mBase, uint32(v737)+48)) = v757
	v763 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_table_parallelscan_initialize(m, v763, v737-int32(-64), v655)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L174
	}
L174:
	;
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v647)+52))
	v769 = F_shm_toc_allocate(m, v768, v670)
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L175
	}
L175:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v647)+44))
	F_tuplesort_initialize_shared(m, v769, v631, v771)
	mBase = m.M
	v773 = m.ExcPending
	if v773 != 0 {
		goto L1
	} else {
		goto L176
	}
L176:
	;
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v647)+52))
	v775 = F_shm_toc_allocate(m, v774, v683)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L177
	}
L177:
	;
	if v683 != 0 {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v777 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v777)+16))
	base.MemoryCopy(m, v775, v778, v683)
	goto L180
L179:
	;
	goto L180
L180:
	;
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v647)+52))
	F_shm_toc_insert(m, v780, int64(-6917529027641081855), v737)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L1
	} else {
		goto L181
	}
L181:
	;
	v784 = *(*int32)(unsafe.Add(mBase, uint32(v647)+52))
	F_shm_toc_insert(m, v784, int64(-6917529027641081854), v769)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L182
	}
L182:
	;
	v788 = *(*int32)(unsafe.Add(mBase, uint32(v647)+52))
	F_shm_toc_insert(m, v788, int64(-6917529027641081853), v775)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L183
	}
L183:
	;
	v793 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[6]))
	if v793 != 0 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(v647)+52))
	v795 = F_shm_toc_allocate(m, v794, v716)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	F_LaunchParallelWorkers(m, v647)
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L1
	} else {
		goto L192
	}
L187:
	;
	if v716 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v798 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[6]))
	base.MemoryCopy(m, v795, v798, v716)
	goto L190
L189:
	;
	goto L190
L190:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(v647)+52))
	F_shm_toc_insert(m, v800, int64(-6917529027641081852), v795)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L1
	} else {
		goto L191
	}
L191:
	;
	goto L186
L192:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v635))) = v647
	v808 = *(*int32)(unsafe.Add(mBase, uint32(v647)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v635)+20)) = v775
	*(*int32)(unsafe.Add(mBase, uint32(v635)+16)) = v655
	*(*int32)(unsafe.Add(mBase, uint32(v635)+12)) = v769
	*(*int32)(unsafe.Add(mBase, uint32(v635)+8)) = v737
	*(*int32)(unsafe.Add(mBase, uint32(v635)+4)) = v808 + int32(1)
	if v808 == int32(0) {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	F_WaitForParallelWorkersToFinish(m, v647)
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L1
	} else {
		goto L196
	}
L194:
	;
	goto L195
L195:
	;
	v838 = F_errstart(m, int32(14), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L203
	}
L196:
	;
	v820 = *(*int32)(unsafe.Add(mBase, uint32(v635)+16))
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v820)))
	if v821 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L197:
	;
	F_UnregisterSnapshot(m, v820)
	mBase = m.M
	v825 = m.ExcPending
	if v825 != 0 {
		goto L1
	} else {
		goto L200
	}
L198:
	;
	goto L199
L199:
	;
	v826 = *(*int32)(unsafe.Add(mBase, uint32(v635)))
	F_DestroyParallelContext(m, v826)
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L1
	} else {
		goto L201
	}
L200:
	;
	goto L199
L201:
	;
	v831 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v832 = *(*int32)(unsafe.Add(mBase, uint32(v831)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v831)+72)) = v832 - int32(1)
	goto L202
L202:
	;
	goto L138
L203:
	;
	if v838 != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v840 = *(*int32)(unsafe.Add(mBase, uint32(v647)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v840
	F_errmsg(m, int32(_a_F_BuildIndex_2_12), v19+int32(-32))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+172)) = v635
	v854 = F_palloc0(m, int32(12))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L209
	}
L207:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_7), int32(955), int32(_a_F_BuildIndex_2_13))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	goto L206
L209:
	;
	v856 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v854)+4)) = v856
	v858 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v854)+8)) = v858
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v635)+8))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v635)+12))
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v635)+20))
	v864 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[7]))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v635)+4))
	v866 = base.I32_div_s(v864, v865)
	F_IvfflatParallelScanAndSort(m, v854, v860, v861, v862, v866, int32(1))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	F_WaitForParallelWorkersToAttach(m, v647)
	mBase = m.M
	v871 = m.ExcPending
	if v871 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	goto L138
L212:
	;
	v885 = F_palloc0(m, int32(12))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L215
	}
L213:
	;
	v895 = int32(0)
	goto L214
L214:
	;
	v897 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[7]))
	v898 = *(*int32)(unsafe.Add(mBase, uint32(l3)+156))
	v899 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+52)) = uint16(v899)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+60)) = int32(97)
	v903 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v903
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+48)) = uint8(v903)
	v917 = F_tuplesort_begin_heap(m, v898, v899, v19+int32(-12), v19+int32(-4), v19+int32(-8), v19+int32(-16), v897, v895, v903)
	mBase = m.M
	v918 = m.ExcPending
	if v918 != 0 {
		goto L1
	} else {
		goto L216
	}
L215:
	;
	v887 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v885))) = uint8(v887)
	v889 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	v890 = *(*int32)(unsafe.Add(mBase, uint32(v889)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v885)+4)) = v890
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v889)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v885)+8)) = v892
	v895 = v885
	goto L214
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3)+152)) = v917
	v920 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v920 != 0 {
		goto L217
	} else {
		goto L218
	}
L217:
	;
	v921 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	if v921 != 0 {
		goto L221
	} else {
		goto L222
	}
L218:
	;
	v1019 = v917
	goto L219
L219:
	;
	F_tuplesort_performsort(m, v1019)
	mBase = m.M
	v1021 = m.ExcPending
	if v1021 != 0 {
		goto L1
	} else {
		goto L236
	}
L220:
	;
	*(*float64)(unsafe.Add(mBase, uint32(l3)+40)) = v998
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	v1019 = v1000
	goto L219
L221:
	;
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v921)+8))
	v926 = v922 + int32(28)
	v927 = *(*int32)(unsafe.Add(mBase, uint32(v921)+4))
	goto L224
L222:
	;
	goto L223
L223:
	;
	v968 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v970 = int32(1)
	v971 = int32(0)
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v920)+188))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v977)+140))
	v979 = m.T0[v978].(func(*base.Module, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32, int32) float64)(m, v920, v968, v969, v970, v971, v970, v971, int32(-1), int32(_a_F_BuildIndex_2_14), l3, v971)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L1
	} else {
		goto L235
	}
L224:
	;
	v948 = base.AtomicRmwXchg32(m, v926, int32(0), int32(1))
	if v948 != 0 {
		goto L226
	} else {
		goto L227
	}
L225:
	;
	v960 = *(*float64)(unsafe.Add(mBase, uint32(v922)+48))
	*(*float64)(unsafe.Add(mBase, uint32(l3)+32)) = v960
	v962 = *(*float64)(unsafe.Add(mBase, uint32(v922)+40))
	v963 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v922)+28)), uint32(v963))
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L1
	} else {
		goto L234
	}
L226:
	;
	F_s_lock(m, v926, int32(_a_F_BuildIndex_2_15))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v952 = *(*int32)(unsafe.Add(mBase, uint32(v922)+32))
	if v927 != v952 {
		goto L230
	} else {
		goto L231
	}
L229:
	;
	goto L228
L230:
	;
	v954 = int32(0)
	atomic.StoreUint32((*uint32)(unsafe.Add(mBase, uint32(v926))), uint32(v954))
	F_ConditionVariableSleep(m, v922+int32(16), int32(134217767))
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	goto L225
L233:
	;
	goto L224
L234:
	;
	v998 = v962
	goto L220
L235:
	;
	v998 = v979
	goto L220
L236:
	;
	v1022 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	v1023 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v1023
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l3)+156))
	v1028 = F_MakeSingleTupleTableSlot(m, v1026, int32(_a_F_BuildIndex_2_16))
	mBase = m.M
	v1029 = m.ExcPending
	if v1029 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	v1030 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v1035 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v1035 == int32(0) {
		goto L239
	} else {
		goto L240
	}
L238:
	;
	v1077 = *(*float64)(unsafe.Add(mBase, uint32(l3)+32))
	v1081 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v1081 == int32(0) {
		goto L243
	} else {
		goto L244
	}
L239:
	;
	goto L238
L240:
	;
	v1039 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v1039&int32(1) == int32(0) {
		goto L239
	} else {
		goto L241
	}
L241:
	;
	v1044 = int32(_a_F_BuildIndex_2_0)
	v1046 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v1047 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1046 + v1047
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1035)))
	*(*int32)(unsafe.Add(mBase, uint32(v1035))) = v1050 + v1047
	v1054 = int32(0)
	v1056 = int32(_a_F_BuildIndex_2_1)
	v1057 = base.AtomicRmwOr32(m, v1054, v1056, v1054)
	*(*int64)(unsafe.Add(mBase, uint32(v1035+int32(80))+232)) = int64(4)
	v1065 = base.AtomicRmwOr32(m, v1054, v1056, v1054)
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1035)))
	*(*int32)(unsafe.Add(mBase, uint32(v1035))) = v1066 + v1047
	v1072 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1072 - v1047
	goto L239
L242:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	F_GetNextTuple(m, v1122, v1030, v1028, v19+int32(-8), v19+int32(-4))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L1
	} else {
		goto L246
	}
L243:
	;
	goto L242
L244:
	;
	v1085 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v1085&int32(1) == int32(0) {
		goto L243
	} else {
		goto L245
	}
L245:
	;
	v1090 = int32(_a_F_BuildIndex_2_0)
	v1092 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v1093 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1092 + v1093
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1081)))
	*(*int32)(unsafe.Add(mBase, uint32(v1081))) = v1096 + v1093
	v1100 = int32(0)
	v1102 = int32(_a_F_BuildIndex_2_1)
	v1103 = base.AtomicRmwOr32(m, v1100, v1102, v1100)
	*(*int64)(unsafe.Add(mBase, uint32(v1081+int32(88))+232)) = base.I64_trunc_sat_f64_s(v1077)
	v1111 = base.AtomicRmwOr32(m, v1100, v1102, v1100)
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1081)))
	*(*int32)(unsafe.Add(mBase, uint32(v1081))) = v1112 + v1093
	v1118 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1118 - v1093
	goto L243
L246:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v1130 = *(*int32)(unsafe.Add(mBase, uint32(v1129)))
	if int32(0) < v1130 {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	v1136 = v1023
	v1149 = int64(0)
	goto L250
L248:
	;
	goto L249
L249:
	;
	v1374 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	F_tuplesort_end(m, v1374)
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L1
	} else {
		goto L291
	}
L250:
	;
	v1153 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[8]))
	if v1153 != 0 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	goto L249
L252:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L1
	} else {
		goto L255
	}
L253:
	;
	goto L254
L254:
	;
	v1156 = F_HnswNewBuffer(m, v1022, l4)
	mBase = m.M
	v1157 = m.ExcPending
	if v1157 != 0 {
		goto L1
	} else {
		goto L256
	}
L255:
	;
	goto L254
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+52)) = v1156
	F_IvfflatInitRegisterPage(m, v1022, v19+int32(-12), v19+int32(-16), v19+int32(-20))
	mBase = m.M
	v1166 = m.ExcPending
	if v1166 != 0 {
		goto L1
	} else {
		goto L257
	}
L257:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v1167 < int32(0) {
		goto L259
	} else {
		goto L260
	}
L258:
	;
	v1187 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v1136 == v1187 {
		goto L262
	} else {
		goto L263
	}
L259:
	;
	v1171 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[3]))
	v1177 = *(*int32)(unsafe.Add(mBase, uint32(v1171+(v1167^int32(-1))*int32(56))+16))
	v1186 = v1177
	goto L258
L260:
	;
	goto L261
L261:
	;
	v1179 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[4]))
	v1180 = int32(56)
	v1185 = *(*int32)(unsafe.Add(mBase, uint32(v1179+v1167*v1180-v1180)+16))
	v1186 = v1185
	goto L258
L262:
	;
	v1204 = v1149
	goto L265
L263:
	;
	v1313 = v1149
	goto L264
L264:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	if v1316 < int32(0) {
		goto L285
	} else {
		goto L286
	}
L265:
	;
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v1208 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1207)+6)))
	v1209 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v1210 = int32(4)
	v1211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1209)+14)))
	v1212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1209)+12)))
	v1213 = v1211 - v1212
	if v1213 <= v1210 {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	v1313 = v1245
	goto L264
L267:
	;
	v1224 = (v1208&int32(_a_F_BuildIndex_2_17) + int32(7)) & int32(_a_F_BuildIndex_2_18)
	if base.Ui32(v1216-int32(4)) < base.Ui32(v1224) {
		goto L271
	} else {
		goto L272
	}
L268:
	;
	v1216 = v1210
	goto L270
L269:
	;
	v1216 = v1213
	goto L270
L270:
	;
	goto L267
L271:
	;
	F_IvfflatAppendPage(m, v1022, v19+int32(-12), v19+int32(-16), v19+int32(-20), l4)
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L1
	} else {
		goto L274
	}
L272:
	;
	goto L273
L273:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v21)+48))
	v1235 = int32(0)
	v1237 = F_PageAddItemExtended(m, v1234, v1207, v1224, v1235, v1235)
	mBase = m.M
	v1238 = m.ExcPending
	if v1238 != 0 {
		goto L1
	} else {
		goto L275
	}
L274:
	;
	goto L273
L275:
	;
	if v1237 == int32(0) {
		goto L82
	} else {
		goto L276
	}
L276:
	;
	F_pfree(m, v1207)
	mBase = m.M
	v1242 = m.ExcPending
	if v1242 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	v1245 = v1204 + int64(1)
	v1248 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[0]))
	if v1248 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L278:
	;
	v1289 = *(*int32)(unsafe.Add(mBase, uint32(l3)+152))
	F_GetNextTuple(m, v1289, v1030, v1028, v19+int32(-8), v19+int32(-4))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L1
	} else {
		goto L282
	}
L279:
	;
	goto L278
L280:
	;
	v1252 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_BuildIndex_2[1])))
	if v1252&int32(1) == int32(0) {
		goto L279
	} else {
		goto L281
	}
L281:
	;
	v1257 = int32(_a_F_BuildIndex_2_0)
	v1259 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	v1260 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1259 + v1260
	v1263 = *(*int32)(unsafe.Add(mBase, uint32(v1248)))
	*(*int32)(unsafe.Add(mBase, uint32(v1248))) = v1263 + v1260
	v1267 = int32(0)
	v1269 = int32(_a_F_BuildIndex_2_1)
	v1270 = base.AtomicRmwOr32(m, v1267, v1269, v1267)
	*(*int64)(unsafe.Add(mBase, uint32(v1248+int32(96))+232)) = v1245
	v1278 = base.AtomicRmwOr32(m, v1267, v1269, v1267)
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v1248)))
	*(*int32)(unsafe.Add(mBase, uint32(v1248))) = v1279 + v1260
	v1285 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[2])) = v1285 - v1260
	goto L279
L282:
	;
	v1296 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	if v1296 == v1136 {
		v1204 = v1245
		goto L265
	} else {
		goto L283
	}
L283:
	;
	goto L266
L284:
	;
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v21)+52))
	v1337 = *(*int32)(unsafe.Add(mBase, uint32(v21)+44))
	F_IvfflatCommitBuffer(m, v1336, v1337)
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L288
	}
L285:
	;
	v1320 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[3]))
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v1320+(v1316^int32(-1))*int32(56))+16))
	v1335 = v1326
	goto L284
L286:
	;
	goto L287
L287:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[4]))
	v1329 = int32(56)
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1328+v1316*v1329-v1329)+16))
	v1335 = v1334
	goto L284
L288:
	;
	v1340 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	v1344 = *(*int64)(unsafe.Add(mBase, uint32(v1340+v1136<<(uint(int32(3))%32))))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v1344
	F_IvfflatUpdateList(m, v1022, v19+int32(-40), v1335, int32(-1), v1186, l4)
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L1
	} else {
		goto L289
	}
L289:
	;
	v1352 = v1136 + int32(1)
	v1353 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	v1354 = *(*int32)(unsafe.Add(mBase, uint32(v1353)))
	if v1352 < v1354 {
		v1136 = v1352
		v1149 = v1313
		goto L250
	} else {
		goto L290
	}
L290:
	;
	goto L251
L291:
	;
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(l3)+172))
	if v1377 != 0 {
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1377)))
	F_WaitForParallelWorkersToFinish(m, v1378)
	mBase = m.M
	v1380 = m.ExcPending
	if v1380 != 0 {
		goto L1
	} else {
		goto L295
	}
L293:
	;
	goto L294
L294:
	;
	if l4 == int32(3) {
		goto L302
	} else {
		goto L303
	}
L295:
	;
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1377)+16))
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1381)))
	if v1382 == int32(0) {
		goto L296
	} else {
		goto L297
	}
L296:
	;
	F_UnregisterSnapshot(m, v1381)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L1
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1377)))
	F_DestroyParallelContext(m, v1387)
	mBase = m.M
	v1389 = m.ExcPending
	if v1389 != 0 {
		goto L1
	} else {
		goto L300
	}
L299:
	;
	goto L298
L300:
	;
	v1392 = *(*int32)(unsafe.Add(mBase, _c_F_BuildIndex_2[5]))
	v1393 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1392)+72)) = v1393 - int32(1)
	goto L301
L301:
	;
	goto L294
L302:
	;
	v1400 = int32(3)
	v1402 = F_RelationGetNumberOfBlocksInFork(m, l1, v1400)
	mBase = m.M
	v1403 = m.ExcPending
	if v1403 != 0 {
		goto L1
	} else {
		goto L305
	}
L303:
	;
	goto L304
L304:
	;
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(l3)+68))
	F_VectorArrayFree(m, v1407)
	mBase = m.M
	v1409 = m.ExcPending
	if v1409 != 0 {
		goto L1
	} else {
		goto L307
	}
L305:
	;
	F_log_newpage_range(m, l1, v1400, v1402, int32(1))
	mBase = m.M
	v1406 = m.ExcPending
	if v1406 != 0 {
		goto L1
	} else {
		goto L306
	}
L306:
	;
	goto L304
L307:
	;
	v1410 = *(*int32)(unsafe.Add(mBase, uint32(l3)+72))
	F_pfree(m, v1410)
	mBase = m.M
	v1412 = m.ExcPending
	if v1412 != 0 {
		goto L1
	} else {
		goto L308
	}
L308:
	;
	v1413 = *(*int32)(unsafe.Add(mBase, uint32(l3)+168))
	F_MemoryContextDelete(m, v1413)
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L1
	} else {
		goto L309
	}
L309:
	;
	m.G0 = v21 - int32(-64)
	return
L310:
	;
	F_errmsg_internal(m, int32(_a_F_BuildIndex_2_19), int32(0))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L1
	} else {
		goto L311
	}
L311:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_20), int32(326), int32(_a_F_BuildIndex_2_21))
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L1
	} else {
		goto L312
	}
L312:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L313:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v1436 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BuildIndex_2_22), v21)
	mBase = m.M
	v1442 = m.ExcPending
	if v1442 != 0 {
		goto L1
	} else {
		goto L314
	}
L314:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_7), int32(546), int32(_a_F_BuildIndex_2_23))
	mBase = m.M
	v1447 = m.ExcPending
	if v1447 != 0 {
		goto L1
	} else {
		goto L315
	}
L315:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L316:
	;
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1022)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v1452 + int32(4)
	F_errmsg_internal(m, int32(_a_F_BuildIndex_2_22), v19+int32(-48))
	mBase = m.M
	v1460 = m.ExcPending
	if v1460 != 0 {
		goto L1
	} else {
		goto L317
	}
L317:
	;
	F_errfinish(m, int32(_a_F_BuildIndex_2_7), int32(315), int32(_a_F_BuildIndex_2_24))
	mBase = m.M
	v1465 = m.ExcPending
	if v1465 != 0 {
		goto L1
	} else {
		goto L318
	}
L318:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_basque_ISO_8859_1_close_env(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	F_SN_delete_env(m, l0)
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		return
	}
}
func F_basque_ISO_8859_1_create_env(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_SN_new_env(m, int32(40))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		if v3 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+36)) = int32(0)
			*(*int64)(unsafe.Add(mBase, uint32(v3)+28)) = int64(0)
		} else {
		}
		return v3
	}
}
func F_bernoulli_beginsamplescan(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v7 float32
	_ = v7
	var v15 float64
	_ = v15
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l1)))
	if base.F32_lt(v7, float32(0))|base.F32_gt(v7, float32(100)) == int32(0) {
		v15 = base.F64_promote_f32(v7)
		if base.Ui64(base.I64_reinterpret_f64(v15)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
			v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
			v39 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v38)+12)) = uint16(v39)
			*(*int32)(unsafe.Add(mBase, uint32(v38)+8)) = l3
			*(*int64)(unsafe.Add(mBase, uint32(v38))) = base.I64_trunc_sat_f64_u(base.F64_nearest(base.F64_div(base.F64_mul(v15, float64(4.294967296e+09)), float64(100))))
			v50 = base.F32_ge(v7, float32(25))
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+133)) = uint8(v50)
			v52 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+132)) = uint8(v52)
			return
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return
			} else {
				F_errcode(m, int32(403177602))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_bernoulli_beginsamplescan_0), int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_bernoulli_beginsamplescan_1), int32(148), int32(_a_F_bernoulli_beginsamplescan_2))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			F_errcode(m, int32(403177602))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return
			} else {
				F_errmsg(m, int32(_a_F_bernoulli_beginsamplescan_0), int32(0))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_bernoulli_beginsamplescan_1), int32(148), int32(_a_F_bernoulli_beginsamplescan_2))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
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
func F_bitlt(m *base.Module, l0 int32) int64 {
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
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v102 != v8 {
		goto L31
	} else {
		goto L32
	}
L2:
	;
	return int64(0)
L3:
	;
	v13 = v8 + int32(8)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	return base.I64_extend_i32_u(int32(base.Ui32(v99) >> (uint(int32(31)) % 32)))
L38:
	;
	goto L37
}
func F_bitncmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	v8 = base.I32_div_s(l2, int32(8))
	if base.Ui32(int32(4)) <= base.Ui32(v8) {
		goto L6
	} else {
		goto L7
	}
L1:
	;
	if v158 != 0 {
		goto L36
	} else {
		goto L37
	}
L2:
	;
	return v156
L3:
	;
	if v70 != 0 {
		v156 = v70
		goto L2
	} else {
		goto L21
	}
L4:
	;
	v70 = int32(0)
	goto L3
L5:
	;
	v44 = v39
	v45 = v40
	v46 = v41
	goto L15
L6:
	;
	if (l0|l1)&int32(3) != 0 {
		v39 = l0
		v40 = l1
		v41 = v8
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v32 = l0
	v33 = l1
	v34 = v8
	goto L8
L8:
	;
	if v34 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v16 = l0
	v17 = l1
	v18 = v8
	goto L10
L10:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if v21 != v22 {
		v39 = v16
		v40 = v17
		v41 = v18
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v32 = v27
	v33 = v25
	v34 = v29
	goto L8
L12:
	;
	v24 = int32(4)
	v25 = v17 + v24
	v27 = v16 + v24
	v29 = v18 - v24
	if base.Ui32(int32(3)) < base.Ui32(v29) {
		v16 = v27
		v17 = v25
		v18 = v29
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v39 = v32
	v40 = v33
	v41 = v34
	goto L5
L15:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45))))
	if v49 == v50 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v70 = v49 - v50
	goto L3
L17:
	;
	v52 = int32(1)
	v57 = v46 - v52
	if v57 != 0 {
		v44 = v44 + v52
		v45 = v45 + v52
		v46 = v57
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	goto L16
L20:
	;
	goto L4
L21:
	;
	v71 = int32(0)
	v74 = l2 - v8<<(uint(int32(3))%32)
	if v74 <= v71 {
		v156 = v71
		goto L2
	} else {
		goto L22
	}
L22:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0+v8))))
	v79 = int32(128)
	v80 = v78 & v79
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v8))))
	if v80 != v82&v79 {
		v158 = v80
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v74 == int32(1) {
		v156 = v71
		goto L2
	} else {
		goto L24
	}
L24:
	;
	v88 = int32(1)
	v90 = int32(128)
	v91 = v78 << (uint(v88) % 32) & v90
	if v91 != v82<<(uint(v88)%32)&v90 {
		v158 = v91
		goto L1
	} else {
		goto L25
	}
L25:
	;
	if v74 < int32(3) {
		v156 = v71
		goto L2
	} else {
		goto L26
	}
L26:
	;
	v99 = int32(2)
	v101 = int32(128)
	v102 = v78 << (uint(v99) % 32) & v101
	if v102 != v82<<(uint(v99)%32)&v101 {
		v158 = v102
		goto L1
	} else {
		goto L27
	}
L27:
	;
	if v74 == int32(3) {
		v156 = v71
		goto L2
	} else {
		goto L28
	}
L28:
	;
	v110 = int32(3)
	v112 = int32(128)
	v113 = v78 << (uint(v110) % 32) & v112
	if v113 != v82<<(uint(v110)%32)&v112 {
		v158 = v113
		goto L1
	} else {
		goto L29
	}
L29:
	;
	if v74 < int32(5) {
		v156 = v71
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v121 = int32(4)
	v123 = int32(128)
	v124 = v78 << (uint(v121) % 32) & v123
	if v124 != v82<<(uint(v121)%32)&v123 {
		v158 = v124
		goto L1
	} else {
		goto L31
	}
L31:
	;
	if v74 == int32(5) {
		v156 = v71
		goto L2
	} else {
		goto L32
	}
L32:
	;
	v132 = int32(5)
	v134 = int32(128)
	v135 = v78 << (uint(v132) % 32) & v134
	if v135 != v82<<(uint(v132)%32)&v134 {
		v158 = v135
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if v74 < int32(7) {
		v156 = v71
		goto L2
	} else {
		goto L34
	}
L34:
	;
	v143 = int32(6)
	v145 = int32(128)
	v146 = v78 << (uint(v143) % 32) & v145
	if v146 != v82<<(uint(v143)%32)&v145 {
		v158 = v146
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v156 = v71
	goto L2
L36:
	;
	v161 = int32(1)
	goto L38
L37:
	;
	v161 = int32(-1)
	goto L38
L38:
	;
	return v161
}
func F_bittoint4(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v77 int64
	_ = v77
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
		if base.Ui32(int32(33)) <= base.Ui32(v13) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v17 = F_errsave_start(m, v16)
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int64(0)
			} else {
				if v17 == int32(0) {
					v77 = int64(0)
					return v77
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_bittoint4_0), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, v16, int32(_a_F_bittoint4_1), int32(1596), int32(_a_F_bittoint4_2))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
								return int64(0)
							} else {
								return int64(0)
							}
						}
					}
				}
			}
		} else {
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
			v37 = int32(base.Ui32(v35) >> (uint(int32(2)) % 32))
			if base.Ui32(int32(36)) <= base.Ui32(v35) {
				v43 = v9 + int32(8)
				v46 = v2
				for {
					v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43))))
					v53 = v50 | v46<<(uint(int32(8))%32)
					v55 = v43 + int32(1)
					if base.Ui32(v55) < base.Ui32(v9+v37) {
						v43 = v55
						v46 = v53
						continue
					} else {
						break
					}
					break
				}
				v60 = v53
			} else {
				v60 = v2
			}
			v77 = base.I64_extend_i32_s(int32(base.Ui32(v60) >> (uint(v37<<(uint(int32(3))%32)-v13+int32(-64)) % 32)))
			return v77
		}
	}
}
func F_bittoint8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v51 int64
	_ = v51
	var v53 int32
	_ = v53
	var v56 int64
	_ = v56
	var v69 int64
	_ = v69
	v2 = int64(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		if base.Ui32(int32(65)) <= base.Ui32(v12) {
			v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v16 = F_errsave_start(m, v15)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return int64(0)
			} else {
				if v16 == int32(0) {
					v69 = v2
					return v69
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_bittoint8_0), int32(0))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, v15, int32(_a_F_bittoint8_1), int32(1676), int32(_a_F_bittoint8_2))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int64(0)
							} else {
								return int64(0)
							}
						}
					}
				}
			}
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
			v36 = int32(base.Ui32(v34) >> (uint(int32(2)) % 32))
			if base.Ui32(int32(36)) <= base.Ui32(v34) {
				v42 = v8 + int32(8)
				v43 = v2
				for {
					v48 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
					v51 = v48 | v43<<(uint(int64(8))%64)
					v53 = v42 + int32(1)
					if base.Ui32(v53) < base.Ui32(v8+v36) {
						v42 = v53
						v43 = v51
						continue
					} else {
						break
					}
					break
				}
				v56 = v51
			} else {
				v56 = v2
			}
			v69 = int64(base.Ui64(v56) >> (uint(base.I64_extend_i32_u(v36<<(uint(int32(3))%32)-v12+int32(-64))) % 64))
			return v69
		}
	}
}
func F_bittypmodin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v8 = F_anybit_typmodin(m, v3, int32(_a_F_bittypmodin_0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(v8)
		}
	}
}
func F_booland_statefunc(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v7 int64
	_ = v7
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v2 == int64(0) {
		return int64(0)
	} else {
		v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		return base.I64_extend_i32_u(base.B2i32(v7 != int64(0)))
	}
}
func F_boolin(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int64
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int64
	_ = v156
	v2 = int32(0)
	v8 = int64(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = v13
	goto L1
L1:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	if base.B2i32(base.Ui32(int32(5)) <= base.Ui32(v22-int32(9)))&base.B2i32(v22 != int32(32)) == int32(0) {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	v34 = F_strlen(m, v16)
	mBase = m.M
	if v34 == int32(0) {
		v66 = v2
		goto L6
	} else {
		goto L7
	}
L3:
	;
	v16 = v16 + int32(1)
	goto L1
L4:
	;
	goto L5
L5:
	;
	goto L2
L6:
	;
	v69 = v11 + int32(15)
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	switch v71 - int32(48) {
	case 0:
		goto L22
	case 1:
		goto L23
	default:
		goto L20
	case 22, 54:
		goto L27
	case 30, 62:
		goto L25
	case 31, 63:
		goto L24
	case 36, 68:
		goto L28
	case 41, 73:
		goto L26
	}
L7:
	;
	v40 = v34
	goto L8
L8:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16+v40-int32(1)))))
	if base.B2i32(base.Ui32(v48-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v48 == int32(32)) == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v66 = v2
	goto L6
L10:
	;
	v66 = v40
	goto L6
L11:
	;
	goto L12
L12:
	;
	v59 = v40 - int32(1)
	if v59 != 0 {
		v40 = v59
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	m.G0 = v11 + int32(16)
	return v156
L15:
	;
	if v132 != 0 {
		goto L46
	} else {
		goto L47
	}
L16:
	;
	v132 = v126
	goto L15
L17:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v69))) = uint8(v122)
	v126 = v124
	goto L16
L18:
	;
	if v69 == int32(0) {
		v126 = v117
		goto L16
	} else {
		goto L45
	}
L19:
	;
	v117 = int32(1)
	goto L18
L20:
	;
	v113 = int32(0)
	if v69 != 0 {
		v122 = v113
		v124 = v113
		goto L17
	} else {
		goto L44
	}
L21:
	;
	v122 = int32(0)
	v124 = v108
	goto L17
L22:
	;
	v103 = int32(1)
	if v66 != v103 {
		goto L20
	} else {
		goto L42
	}
L23:
	;
	v100 = int32(1)
	if v66 != v100 {
		goto L20
	} else {
		goto L41
	}
L24:
	;
	v89 = int32(2)
	if base.Ui32(v66) <= base.Ui32(v89) {
		goto L35
	} else {
		goto L36
	}
L25:
	;
	v85 = F_pg_strncasecmp(m, v16, int32(_a_F_boolin_0), v66)
	mBase = m.M
	if v85 != 0 {
		goto L20
	} else {
		goto L33
	}
L26:
	;
	v81 = F_pg_strncasecmp(m, v16, int32(_a_F_boolin_1), v66)
	mBase = m.M
	if v81 == int32(0) {
		goto L19
	} else {
		goto L32
	}
L27:
	;
	v77 = F_pg_strncasecmp(m, v16, int32(_a_F_boolin_2), v66)
	mBase = m.M
	if v77 != 0 {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	v75 = F_pg_strncasecmp(m, v16, int32(_a_F_boolin_3), v66)
	mBase = m.M
	if v75 != 0 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	goto L19
L30:
	;
	if v69 != 0 {
		v108 = int32(1)
		goto L21
	} else {
		goto L31
	}
L31:
	;
	v132 = int32(1)
	goto L15
L32:
	;
	goto L20
L33:
	;
	if v69 != 0 {
		v108 = int32(1)
		goto L21
	} else {
		goto L34
	}
L34:
	;
	v132 = int32(1)
	goto L15
L35:
	;
	v92 = v89
	goto L37
L36:
	;
	v92 = v66
	goto L37
L37:
	;
	v93 = F_pg_strncasecmp(m, v16, int32(_a_F_boolin_4), v92)
	mBase = m.M
	if v93 == int32(0) {
		goto L19
	} else {
		goto L38
	}
L38:
	;
	v97 = F_pg_strncasecmp(m, v16, int32(_a_F_boolin_5), v92)
	mBase = m.M
	if v97 != 0 {
		goto L20
	} else {
		goto L39
	}
L39:
	;
	if v69 != 0 {
		v108 = int32(1)
		goto L21
	} else {
		goto L40
	}
L40:
	;
	v132 = int32(1)
	goto L15
L41:
	;
	v117 = v100
	goto L18
L42:
	;
	if v69 != 0 {
		v108 = v103
		goto L21
	} else {
		goto L43
	}
L43:
	;
	v132 = int32(1)
	goto L15
L44:
	;
	v126 = v113
	goto L16
L45:
	;
	v122 = v117
	v124 = int32(1)
	goto L17
L46:
	;
	v133 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
	v156 = v133
	goto L14
L47:
	;
	goto L48
L48:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v135 = F_errsave_start(m, v134)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	return int64(0)
L50:
	;
	if v135 == int32(0) {
		v156 = v8
		goto L14
	} else {
		goto L51
	}
L51:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L49
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v13
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_boolin_6)
	F_errmsg(m, int32(_a_F_boolin_7), v11)
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L49
	} else {
		goto L53
	}
L53:
	;
	F_errsave_finish(m, v134, int32(_a_F_boolin_8), int32(151), int32(_a_F_boolin_9))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L49
	} else {
		goto L54
	}
L54:
	;
	v156 = v8
	goto L14
}
func F_boolle(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(0)
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 == v3) | base.B2i32(v5 != v3))
}
func F_boolne(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int64
	_ = v2
	var v3 int64
	_ = v3
	var v5 int64
	_ = v5
	v2 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = int64(0)
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_u(base.B2i32(v2 != v3) ^ base.B2i32(v5 != v3))
}
func F_boolop(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
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
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum_copy(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v17 = F_pg_detoast_datum(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			if v19 != 0 {
				v20 = F_array_contains_nulls(m, v12)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					if v20 != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(67108994))
							mBase = m.M
							v91 = m.ExcPending
							if v91 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_boolop_0), int32(0))
								mBase = m.M
								v95 = m.ExcPending
								if v95 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_boolop_1), int32(424), int32(_a_F_boolop_2))
									mBase = m.M
									v100 = m.ExcPending
									if v100 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						v25 = F_ArrayGetNItemsSafe(m, v22, v12+int32(16))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int64(0)
						} else {
							v27 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)) = uint8(v27)
							v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
							if v29 != 0 {
								v37 = v29
							} else {
								v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
								v37 = (v30<<(uint(int32(3))%32) + int32(23)) & int32(-8)
							}
							F_isort(m, v37+v12, v25, v9+int32(7))
							mBase = m.M
							v42 = F__int_unique(m, v12)
							mBase = m.M
							v43 = m.ExcPending
							if v43 != 0 {
								return int64(0)
							} else {
								v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
								v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
								if v44 != 0 {
									v52 = v44
								} else {
									v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
								}
								v53 = v42 + v52
								*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v53
								v57 = F_ArrayGetNItemsSafe(m, v45, v42+int32(16))
								mBase = m.M
								v58 = m.ExcPending
								if v58 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v53 + v57<<(uint(int32(2))%32)
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
									v72 = F_execute(m, v17+v63<<(uint(int32(3))%32), v9+int32(8), int32(0), int32(1), int32(_a_F_boolop_3))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int64(0)
									} else {
										F_pfree(m, v42)
										mBase = m.M
										v75 = m.ExcPending
										if v75 != 0 {
											return int64(0)
										} else {
											v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v76 != v17 {
												F_pfree(m, v17)
												mBase = m.M
												v79 = m.ExcPending
												if v79 != 0 {
													return int64(0)
												} else {
													m.G0 = v9 + int32(16)
													return base.I64_extend_i32_u(v72)
												}
											} else {
												m.G0 = v9 + int32(16)
												return base.I64_extend_i32_u(v72)
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
				v25 = F_ArrayGetNItemsSafe(m, v22, v12+int32(16))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int64(0)
				} else {
					v27 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v9)+7)) = uint8(v27)
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
					if v29 != 0 {
						v37 = v29
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
						v37 = (v30<<(uint(int32(3))%32) + int32(23)) & int32(-8)
					}
					F_isort(m, v37+v12, v25, v9+int32(7))
					mBase = m.M
					v42 = F__int_unique(m, v12)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int64(0)
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
						if v44 != 0 {
							v52 = v44
						} else {
							v52 = (v45<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						}
						v53 = v42 + v52
						*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v53
						v57 = F_ArrayGetNItemsSafe(m, v45, v42+int32(16))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v53 + v57<<(uint(int32(2))%32)
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							v72 = F_execute(m, v17+v63<<(uint(int32(3))%32), v9+int32(8), int32(0), int32(1), int32(_a_F_boolop_3))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								F_pfree(m, v42)
								mBase = m.M
								v75 = m.ExcPending
								if v75 != 0 {
									return int64(0)
								} else {
									v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v76 != v17 {
										F_pfree(m, v17)
										mBase = m.M
										v79 = m.ExcPending
										if v79 != 0 {
											return int64(0)
										} else {
											m.G0 = v9 + int32(16)
											return base.I64_extend_i32_u(v72)
										}
									} else {
										m.G0 = v9 + int32(16)
										return base.I64_extend_i32_u(v72)
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
func F_booltext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	if v4 == int64(0) {
		v7 = int32(_a_F_booltext_0)
	} else {
		v7 = int32(_a_F_booltext_1)
	}
	v8 = F_cstring_to_text(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(v8)
	}
}
func F_boot_yylex_init(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int64
	_ = v29
	if l0 == int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_boot_yylex_init[0])) = int32(28)
		return int32(1)
	} else {
		v11 = F_palloc(m, int32(96))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0))) = v11
			if v11 == int32(0) {
				*(*int32)(unsafe.Add(mBase, _c_F_boot_yylex_init[0])) = int32(48)
				return int32(1)
			} else {
				v23 = int32(0)
				base.MemoryFill(m, v11, v23, int32(96))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(v26)+60)) = v23
				v29 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v26)+52)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v26)+44)) = v23
				*(*int64)(unsafe.Add(mBase, uint32(v26)+36)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v26)+4)) = v29
				*(*int64)(unsafe.Add(mBase, uint32(v26)+12)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v26)+20)) = v23
				return v23
			}
		}
	}
}
func F_bpchargt(m *base.Module, l0 int32) int64 {
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
	return int64(0)
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
	v18 = int32(1)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v22 = v20 & v18
	if v22 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v23 = v18
	goto L6
L5:
	;
	v23 = int32(4)
	goto L6
L6:
	;
	v24 = v23 + v11
	if v20 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v57 = v51
	goto L18
L8:
	;
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v30 == int32(18) {
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
	if v22 != 0 {
		v51 = int32(base.Ui32(v20)>>(uint(v41)%32)) - v41
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v33 = int32(16)
	goto L13
L12:
	;
	v33 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v30-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v40 = int32(4)
	goto L16
L15:
	;
	v40 = v33
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
	if v57 <= int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v74 = int32(1)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16))))
	v78 = v76 & v74
	if v78 != 0 {
		goto L25
	} else {
		goto L26
	}
L20:
	;
	goto L19
L21:
	;
	v73 = v51 & (v51 >> (uint(int32(31)) % 32))
	goto L20
L22:
	;
	goto L23
L23:
	;
	v67 = v57 - int32(1)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24+v67))))
	if v69 == int32(32) {
		v57 = v67
		goto L18
	} else {
		goto L24
	}
L24:
	;
	v73 = v57
	goto L20
L25:
	;
	v79 = v74
	goto L27
L26:
	;
	v79 = int32(4)
	goto L27
L27:
	;
	v80 = v79 + v16
	if v76 == int32(1) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v113 = v107
	goto L39
L29:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v86 == int32(18) {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	goto L31
L31:
	;
	v97 = int32(1)
	if v78 != 0 {
		v107 = int32(base.Ui32(v76)>>(uint(v97)%32)) - v97
		goto L28
	} else {
		goto L38
	}
L32:
	;
	v89 = int32(16)
	goto L34
L33:
	;
	v89 = int32(0)
	goto L34
L34:
	;
	if base.Ui32((v86-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v96 = int32(4)
	goto L37
L36:
	;
	v96 = v89
	goto L37
L37:
	;
	v107 = v96
	goto L28
L38:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
	goto L28
L39:
	;
	if v113 <= int32(0) {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v131 = F_varstr_cmp(m, v24, v73, v80, v128, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L46
	}
L41:
	;
	goto L40
L42:
	;
	v128 = v107 & (v107 >> (uint(int32(31)) % 32))
	goto L41
L43:
	;
	goto L44
L44:
	;
	v123 = v113 - int32(1)
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80+v123))))
	if v125 == int32(32) {
		v113 = v123
		goto L39
	} else {
		goto L45
	}
L45:
	;
	v128 = v113
	goto L41
L46:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v133 != v11 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	F_pfree(m, v11)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v137 != v16 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	F_pfree(m, v16)
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	return base.I64_extend_i32_u(base.B2i32(int32(0) < v131))
L54:
	;
	goto L53
}
func F_bpcharlen(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	v15 = v13 & int32(1)
	if v15 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v16 = int32(1)
	goto L5
L4:
	;
	v16 = int32(4)
	goto L5
L5:
	;
	if v13 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v48 = v44
	goto L17
L7:
	;
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v23 == int32(18) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v34 = int32(1)
	if v15 != 0 {
		v44 = int32(base.Ui32(v13)>>(uint(v34)%32)) - v34
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v26 = int32(16)
	goto L12
L11:
	;
	v26 = int32(0)
	goto L12
L12:
	;
	if base.Ui32((v23-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v33 = int32(4)
	goto L15
L14:
	;
	v33 = v26
	goto L15
L15:
	;
	v44 = v33
	goto L6
L16:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v44 = int32(base.Ui32(v38)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	if v48 <= int32(0) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_bpcharlen[0]))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65*int32(28))+uint32(_c_F_bpcharlen[1])))
	goto L24
L19:
	;
	goto L18
L20:
	;
	v62 = v44 & (v44 >> (uint(int32(31)) % 32))
	goto L19
L21:
	;
	goto L22
L22:
	;
	v56 = v48 - int32(1)
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9+v16+v56))))
	if v58 == int32(32) {
		v48 = v56
		goto L17
	} else {
		goto L23
	}
L23:
	;
	v62 = v48
	goto L19
L24:
	;
	if v70 != int32(1) {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v73 = int32(1)
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v75&v73 != 0 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v82 = v62
	goto L27
L27:
	;
	return base.I64_extend_i32_s(v82)
L28:
	;
	v78 = v73
	goto L30
L29:
	;
	v78 = int32(4)
	goto L30
L30:
	;
	v80 = F_pg_mbstrlen_with_len(m, v9+v78, v62)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v82 = v80
	goto L27
}
func F_bqarr_out(m *base.Module, l0 int32) int64 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int64
	_ = v37
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int64(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		if v15 != 0 {
			v16 = int32(32)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v8))) = v11 + v15<<(uint(int32(3))%32)
			v24 = F_palloc_mul(m, int32(1), v16)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v24
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v24
				v28 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v24))) = uint8(v28)
				F_infix_3(m, v8, int32(1))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v33 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int64(0)
						} else {
							v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)))
							m.G0 = v8 + int32(16)
							return v37
						}
					} else {
						v37 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v8)+4)))
						m.G0 = v8 + int32(16)
						return v37
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_bqarr_out_0), int32(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int64(0)
					} else {
						F_errfinish(m, int32(_a_F_bqarr_out_1), int32(696), int32(_a_F_bqarr_out_2))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int64(0)
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
func F_brinbeginscan(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	v4 = F_RelationGetIndexScan(m, l0, l1, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(12))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = F_brinRevmapInitialize(m, l0, v9)
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = v11
				v14 = F_brin_build_desc(m, l0)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v4)+36)) = v9
					return v4
				}
			}
		}
	}
}
func F_brinendscan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
	F_brinRevmapTerminate(m, v3)
	mBase = m.M
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(v2)+8))
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
		F_MemoryContextDelete(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			F_pfree(m, v2)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_brinhandler(m *base.Module, l0 int32) int64 {
	return int64(4096)
}
func F_brinoptions(m *base.Module, l0 int64, l1 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = F_build_reloptions(m, l0, l1, int32(1024), int32(12), int32(_a_F_brinoptions_0), int32(2))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_brtuple_disk_tupdesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
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
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v2 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v11 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(_a_F_brtuple_disk_tupdesc_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_brtuple_disk_tupdesc[0]))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, _c_F_brtuple_disk_tupdesc[0])) = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v20 = F_CreateTemplateTupleDesc(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v191 = v11
	goto L3
L3:
	;
	return v191
L4:
	;
	return int32(0)
L5:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if int32(0) < v25 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v33 = int32(1)
	v35 = v25
	v36 = v2
	goto L9
L7:
	;
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_brtuple_disk_tupdesc[0])) = v15
	v102 = int32(0)
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	if v102 < v111 {
		goto L20
	} else {
		goto L21
	}
L9:
	;
	v44 = l0 + int32(20) + v36<<(uint(int32(2))%32)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45))))
	if v46 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L8
L11:
	;
	v49 = v33
	v50 = int32(0)
	v53 = v45
	goto L14
L12:
	;
	v79 = v33
	v81 = v35
	goto L13
L13:
	;
	v88 = v36 + int32(1)
	if v88 < v81 {
		v33 = v79
		v35 = v81
		v36 = v88
		goto L9
	} else {
		goto L18
	}
L14:
	;
	v58 = int32(0)
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v53+v50<<(uint(int32(2))%32))+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	F_TupleDescInitEntry(m, v20, base.I32_extend16_s(v49), v58, v63, int32(-1), v58)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L4
	} else {
		goto L16
	}
L15:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v79 = v69
	v81 = v76
	goto L13
L16:
	;
	v68 = int32(1)
	v69 = v49 + v68
	v71 = v50 + v68
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v72))))
	if base.Ui32(v71) < base.Ui32(v73) {
		v49 = v69
		v50 = v71
		v53 = v72
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	goto L10
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v20
	v191 = v20
	goto L3
L20:
	;
	v115 = v20 + int32(28)
	v122 = v102
	v123 = v111
	v125 = v102
	goto L24
L21:
	;
	v179 = v102
	v186 = v111
	goto L22
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+20)) = v186
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v179
	goto L19
L23:
	;
	v179 = v173
	v186 = v152
	goto L22
L24:
	;
	v131 = v115 + v111<<(uint(int32(3))%32) + v122*int32(100)
	v134 = v115 + v122<<(uint(int32(3))%32)
	if v111 != v123 {
		v152 = v123
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v173 = v111
	goto L23
L26:
	;
	v153 = int32(*(*int16)(unsafe.Add(mBase, uint32(v134)+2)))
	if v153 <= int32(0) {
		v173 = v122
		goto L23
	} else {
		goto L34
	}
L27:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+7)))
	if v136 != int32(118) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v152 = v122
	goto L26
L29:
	;
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+4)))
	if v139 != int32(1) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+6)))
	if v142&int32(6) != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	v145 = int32(*(*int16)(unsafe.Add(mBase, uint32(v134)+2)))
	if v145 <= int32(0) {
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+90)))
	if v148 != int32(118) {
		v152 = v111
		goto L26
	} else {
		goto L33
	}
L33:
	;
	goto L28
L34:
	;
	v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v131)+90)))
	if v156 == int32(118) {
		v173 = v122
		goto L23
	} else {
		goto L35
	}
L35:
	;
	v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134)+5)))
	v165 = (v125 + v159 - int32(1)) & (int32(0) - v159)
	if int32(_a_F_brtuple_disk_tupdesc_1) < v165 {
		v173 = v122
		goto L23
	} else {
		goto L36
	}
L36:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v134))) = uint16(v165)
	v171 = v122 + int32(1)
	if v171 != v111 {
		v122 = v171
		v123 = v152
		v125 = v165 + v153
		goto L24
	} else {
		goto L37
	}
L37:
	;
	goto L25
}
func F_btbpchar_pattern_sortsupport(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn14237(m, l0, int32(1042))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_btbuildempty(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	v5 = F__bt_allequalimage(m, l0, int32(0))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		v8 = F_smgr_bulk_start_rel(m, l0, int32(3))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return
		} else {
			v10 = F_smgr_bulk_get_buf(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return
			} else {
				v12 = int32(0)
				F_PageInit(m, v10, int32(_a_F_btbuildempty_0), int32(16))
				mBase = m.M
				*(*uint8)(unsafe.Add(mBase, uint32(v10)+64)) = uint8(v5)
				*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = int64(-4616189618054758400)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v12
				*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v12
				*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(17180209506)
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+16)))
				v30 = int32(8)
				*(*uint16)(unsafe.Add(mBase, uint32(v10+v28)+12)) = uint16(v30)
				v32 = int32(72)
				*(*uint16)(unsafe.Add(mBase, uint32(v10)+12)) = uint16(v32)
				F_smgr_bulk_write(m, v8, int32(0), v10, int32(1))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_smgr_bulk_finish(m, v8)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						return
					}
				}
			}
		}
	}
}
func F_btbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int64
	_ = v65
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v93 int32
	_ = v93
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v140 int32
	_ = v140
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
	var v186 int32
	_ = v186
	var v188 int64
	_ = v188
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v259 int32
	_ = v259
	var v268 int32
	_ = v268
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v328 int64
	_ = v328
	var v334 int32
	_ = v334
	var v357 int32
	_ = v357
	var v361 int32
	_ = v361
	var v378 int32
	_ = v378
	var v387 int32
	_ = v387
	var v395 int32
	_ = v395
	var v416 int32
	_ = v416
	var v417 int64
	_ = v417
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int64
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	v5 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(208)
	m.G0 = v23
	v31 = v5
	v32 = v5
	v33 = v5
	v34 = v5
	v35 = v5
	v36 = v5
	v37 = int32(-1)
	v44 = int64(0)
	goto L1
L1:
	;
	goto L3
L2:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L3:
	;
	if v37 != int32(1) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L2
L5:
	;
	v416 = int32(m.ExcTag)
	v417 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v416 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v82 = v31
	v83 = v32
	v84 = v33
	v85 = v34
	v86 = v35
	v87 = v36
	v88 = v44
	goto L8
L8:
	;
	if v82 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L9:
	;
	v58 = v36
	v59 = l1
	goto L11
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v35
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v44
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v48
	v56 = F_palloc0(m, int32(40))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L5
	} else {
		goto L12
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v34
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v35
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v48
	v65 = base.I64_extend_i32_u(v48)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v65
	F_before_shmem_exit(m, int32(217), v65)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L5
	} else {
		goto L13
	}
L12:
	;
	v58 = v56
	v59 = v56
	goto L11
L13:
	;
	v72 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[0]))
	v74 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[1]))
	goto L14
L14:
	;
	v76 = v23 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v76))) = v23 + int32(12)
	goto L17
L15:
	;
	v82 = int32(0)
	v83 = v48
	v84 = v59
	v85 = v74
	v86 = v72
	v87 = v58
	v88 = v65
	goto L8
L17:
	;
	goto L15
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v86
	v93 = int32(16)
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[1])) = v23 + v93
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v83
	v101 = m.G0
	v103 = v101 - v93
	m.G0 = v103
	v106 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	v110 = F_LWLockAcquire(m, v106+int32(2560), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L5
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[0])) = v86
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[1])) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v83
	F_cancel_before_shmem_exit(m, int32(217), v88)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L5
	} else {
		goto L60
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v83
	F_btvacuumscan(m, l0, v84, l2, l3, v125&int32(_a_F_btbulkdelete_0))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L5
	} else {
		goto L48
	}
L22:
	;
	v113 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[3]))
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113))))
	v116 = v114 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v113))) = uint16(v116)
	if base.Ui32(int32(_a_F_btbulkdelete_1)) <= base.Ui32(v114) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v121 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[3]))
	v122 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v121))) = uint16(v122)
	v125 = v122
	goto L25
L24:
	;
	v125 = v116
	goto L25
L25:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[3]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if int32(0) < v128 {
		goto L28
	} else {
		goto L29
	}
L26:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	F_LWLockRelease(m, v234+int32(2560))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L5
	} else {
		goto L44
	}
L27:
	;
	v212 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	F_LWLockRelease(m, v212+int32(2560))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L5
	} else {
		goto L40
	}
L28:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v83)+60))
	v140 = int32(0)
	goto L31
L29:
	;
	goto L30
L30:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v127)+8))
	if v186 <= v128 {
		goto L26
	} else {
		goto L38
	}
L31:
	;
	v157 = v127 + int32(12) + v140*int32(12)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
	if v133 == v158 {
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L30
L33:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v83)+64))
	if v160 == v161 {
		goto L27
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v164 = v140 + int32(1)
	if v164 != v128 {
		v140 = v164
		goto L31
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	goto L32
L38:
	;
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v83)+60))
	v191 = v127 + v128*int32(12)
	*(*uint16)(unsafe.Add(mBase, uint32(v191)+20)) = uint16(v125)
	*(*int64)(unsafe.Add(mBase, uint32(v191)+12)) = v188
	v195 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[3]))
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v195)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v195)+4)) = v196 + int32(1)
	v201 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	F_LWLockRelease(m, v201+int32(2560))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L5
	} else {
		goto L39
	}
L39:
	;
	m.G0 = v103 + int32(16)
	goto L21
L40:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v83)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v103))) = v221 + int32(4)
	F_errmsg_internal(m, int32(_a_F_btbulkdelete_2), v103)
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L5
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F_btbulkdelete_3), int32(504), int32(_a_F_btbulkdelete_4))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L5
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
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L5
	} else {
		goto L45
	}
L45:
	;
	F_errmsg_internal(m, int32(_a_F_btbulkdelete_5), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L5
	} else {
		goto L46
	}
L46:
	;
	F_errfinish(m, int32(_a_F_btbulkdelete_3), int32(512), int32(_a_F_btbulkdelete_4))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v83
	F_cancel_before_shmem_exit(m, int32(217), v88)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L5
	} else {
		goto L49
	}
L49:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[0])) = v86
	*(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[1])) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v86
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v85
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v83
	v279 = int32(0)
	v281 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	v285 = F_LWLockAcquire(m, v281+int32(2560), v279)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L5
	} else {
		goto L50
	}
L50:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[3]))
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v288)+4))
	if v289 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_btbulkdelete[2]))
	F_LWLockRelease(m, v357+int32(2560))
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L5
	} else {
		goto L59
	}
L52:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v83)+60))
	v303 = v279
	goto L53
L53:
	;
	v317 = v288 + int32(12) + v303*int32(12)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	if v318 != v294 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	goto L51
L55:
	;
	v334 = v303 + int32(1)
	if v334 != v289 {
		v303 = v334
		goto L53
	} else {
		goto L58
	}
L56:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v317)+4))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v83)+64))
	if v320 != v321 {
		goto L55
	} else {
		goto L57
	}
L57:
	;
	v325 = v288 + v289*int32(12)
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v317)+8)) = v326
	v328 = *(*int64)(unsafe.Add(mBase, uint32(v325)))
	*(*int64)(unsafe.Add(mBase, uint32(v317))) = v328
	*(*int32)(unsafe.Add(mBase, uint32(v288)+4)) = v289 - int32(1)
	goto L51
L58:
	;
	goto L54
L59:
	;
	m.G0 = v23 + int32(208)
	return v84
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v83
	F__bt_end_vacuum_callback(m, int32(0), v88)
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L5
	} else {
		goto L61
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+180)) = v85
	*(*int32)(unsafe.Add(mBase, uint32(v23)+176)) = v86
	*(*int64)(unsafe.Add(mBase, uint32(v23)+184)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v23)+196)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v23)+200)) = v87
	*(*int32)(unsafe.Add(mBase, uint32(v23)+204)) = v83
	F_pg_re_throw(m)
	mBase = m.M
	v395 = m.ExcPending
	if v395 != 0 {
		goto L5
	} else {
		goto L62
	}
L62:
	;
	goto L4
L63:
	;
	v421 = int32(v417)
	m.G0 = v23
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v421)+4))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v421)))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	if v23+int32(12) == v427 {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	m.ExcPending = 1
	goto L72
L65:
	;
	if v431 != 0 {
		goto L69
	} else {
		goto L70
	}
L66:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v424)+4))
	v431 = v429
	goto L68
L67:
	;
	v431 = int32(0)
	goto L68
L68:
	;
	goto L65
L69:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v23)+204))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v23)+196))
	v435 = *(*int64)(unsafe.Add(mBase, uint32(v23)+184))
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v23)+180))
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v23)+176))
	v31 = v423
	v32 = v432
	v33 = v434
	v34 = v436
	v35 = v437
	v36 = v433
	v37 = v431
	v44 = v435
	goto L1
L70:
	;
	goto L71
L71:
	;
	F___wasm_longjmp(m, v424, v423)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	return int32(0)
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_btequalimage(m *base.Module, l0 int32) int64 {
	return int64(1)
}
func F_btgetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v19 int64
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int64
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int64
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v19 = int64(0)
	goto L2
L1:
	;
	return v62
L2:
	;
	v21 = F__bt_first(m, l0, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v70 == int32(0) {
		goto L1
	} else {
		goto L18
	}
L4:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	if v63 == int32(0) {
		goto L1
	} else {
		goto L16
	}
L5:
	;
	return int64(0)
L6:
	;
	if v21 == int32(0) {
		v62 = v19
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v29 = l0 + int32(60)
	v33 = v19
	goto L8
L8:
	;
	F_tbm_add_tuples(m, l1, v29, int32(1), int32(0))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L10
	}
L10:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v40 = v38 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+100)) = v40
	v43 = v33 + int64(1)
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+96))
	if v44 < v40 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = F__bt_next(m, l0, int32(1))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L14
	}
L12:
	;
	v52 = v40
	goto L13
L13:
	;
	v29 = v10 + int32(104) + v52*int32(10)
	v33 = v43
	goto L8
L14:
	;
	if v47 == int32(0) {
		v62 = v43
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v10)+100))
	v52 = v51
	goto L13
L16:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v67 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v66)+18)) = uint16(v67)
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+17)))
	if v69 != 0 {
		v19 = v62
		goto L2
	} else {
		goto L17
	}
L17:
	;
	goto L3
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+17)))
	if v74 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v70)+24))
	v76 = v70 + v75
	v78 = v76 + int32(12)
	v80 = F_LWLockAcquire(m, v78, int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L20
	}
L20:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	if v82 != int32(4) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v76)+8)) = int32(4)
	F_LWLockRelease(m, v78)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	F_LWLockRelease(m, v78)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L5
	} else {
		goto L26
	}
L24:
	;
	F_ConditionVariableBroadcast(m, v76+int32(28))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L5
	} else {
		goto L25
	}
L25:
	;
	return v62
L26:
	;
	goto L1
}
func F_btgettreeheight(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F__bt_getrootheight(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_btint82cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = int64(*(*int16)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_s(base.B2i32(v5 < v4) - base.B2i32(v4 < v5))
}
func F_btint84cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = int64(*(*int32)(unsafe.Add(mBase, uint32(l0)+40)))
	return base.I64_extend_i32_s(base.B2i32(v5 < v4) - base.B2i32(v4 < v5))
}
func F_btint8cmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int64
	_ = v5
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	return base.I64_extend_i32_s(base.B2i32(v5 < v4) - base.B2i32(v4 < v5))
}
func F_btint8skipsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = int32(203)
	*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = int32(204)
	*(*int64)(unsafe.Add(mBase, uint32(v6)+8)) = int64(9223372036854775807)
	*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(-9223372036854775807 - 1)
	return int64(0)
}
func F_btmarkpos(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btmarkpos[0])))
	if v4 != 0 {
		F_ReleaseBuffer(m, v4)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btmarkpos[0]))) = int32(0)
			v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
			if v9 != int32(-1) {
				v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+100))
				*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = v12
				return
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btmarkpos[0]))) = int64(-4294967296)
				*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(-1)
				return
			}
		}
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v3)+60))
		if v9 != int32(-1) {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)+100))
			*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = v12
			return
		} else {
			*(*int64)(unsafe.Add(mBase, uint32(v3)+uint32(_c_F_btmarkpos[0]))) = int64(-4294967296)
			*(*int32)(unsafe.Add(mBase, uint32(v3)+52)) = int32(-1)
			return
		}
	}
}
func F_btnamecmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
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
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	goto L6
L2:
	;
	goto L3
L3:
	;
	v57 = F_strlen(m, v5)
	mBase = m.M
	v58 = F_strlen(m, v4)
	mBase = m.M
	v59 = F_varstr_cmp(m, v5, v57, v4, v58, v6)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	return base.I64_extend_i32_s(v46 - v47)
L6:
	;
	goto L7
L7:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L12
L9:
	;
	v42 = v4
	v46 = int32(0)
	goto L10
L10:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L4
L11:
	;
	v42 = v37
	v46 = v39
	goto L10
L12:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(v19 != v21)|base.B2i32(v21 == int32(0)) != 0 {
		v37 = v17
		v39 = v19
		goto L11
	} else {
		goto L14
	}
L13:
	;
	v37 = v31
	v39 = int32(0)
	goto L11
L14:
	;
	v27 = v18 - int32(1)
	if v27 == int32(0) {
		v37 = v17
		v39 = v19
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v30 = int32(1)
	v31 = v17 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v32 != 0 {
		v16 = v16 + v30
		v17 = v31
		v18 = v27
		v19 = v32
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	return int64(0)
L18:
	;
	return base.I64_extend_i32_s(v59)
}
func F_btparallelrescan(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(v3)+24))
	v5 = v3 + v4
	v7 = v5 + int32(12)
	v9 = F_LWLockAcquire(m, v7, int32(0))
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(v5)+8)) = int32(0)
		*(*int64)(unsafe.Add(mBase, uint32(v5))) = int64(-1)
		F_LWLockRelease(m, v7)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			return
		}
	}
}
func F_btvacuumcleanup(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v178 int32
	_ = v178
	var v181 int64
	_ = v181
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v184 int32
	_ = v184
	var v186 int64
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 float64
	_ = v203
	var v204 float64
	_ = v204
	var v209 int32
	_ = v209
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v9 != 0 {
		v209 = l1
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v209
L2:
	;
	if l1 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v12 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = F_ReadBuffer(m, v13, v12)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v76 = l1
	goto L5
L5:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v76)+28))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v76)+32))
	v83 = v81 - v82
	v84 = m.G0
	v86 = v84 - int32(32)
	m.G0 = v86
	v89 = F_ReadBuffer(m, v80, int32(0))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L7
	} else {
		goto L29
	}
L6:
	;
	if v63 == int32(0) {
		v209 = v12
		goto L1
	} else {
		goto L26
	}
L7:
	;
	return int32(0)
L8:
	;
	F_LockBufferInternal(m, v15, int32(1))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	F__bt_checkpage(m, v13, v15)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L7
	} else {
		goto L10
	}
L10:
	;
	if v15 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+28))
	if base.Ui32(v42) <= base.Ui32(int32(2)) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v27 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[0]))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v27+(v15^int32(-1))<<(uint(int32(2))%32))))
	v41 = v33
	goto L11
L13:
	;
	goto L14
L14:
	;
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[1]))
	v41 = v35 + v15<<(uint(int32(13))%32) + int32(-8192)
	goto L11
L15:
	;
	F_UnlockReleaseBuffer(m, v15)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v41)+48))
	F_UnlockReleaseBuffer(m, v15)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L7
	} else {
		goto L19
	}
L18:
	;
	v63 = int32(1)
	goto L6
L19:
	;
	if v48 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v63 = v60
	goto L6
L21:
	;
	v53 = F_RelationGetNumberOfBlocksInFork(m, v13, int32(0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L7
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v60 = int32(0)
	goto L20
L24:
	;
	v56 = base.I32_div_u_s(v53, int32(20))
	if base.Ui32(v56) < base.Ui32(v48) {
		v60 = int32(1)
		goto L20
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	v67 = F_palloc0(m, int32(40))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L7
	} else {
		goto L27
	}
L27:
	;
	v69 = int32(0)
	F_btvacuumscan(m, l0, v67, v69, v69, v69)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L7
	} else {
		goto L28
	}
L28:
	;
	v74 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v67)+4)) = uint8(v74)
	v76 = v67
	goto L5
L29:
	;
	F_LockBufferInternal(m, v89, int32(1))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L7
	} else {
		goto L30
	}
L30:
	;
	F__bt_checkpage(m, v80, v89)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	if v89 < int32(0) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v89)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L7
	} else {
		goto L60
	}
L33:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v113)+28))
	if base.Ui32(int32(3)) <= base.Ui32(v114) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v99 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[0]))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v99+(v89^int32(-1))<<(uint(int32(2))%32))))
	v113 = v105
	goto L33
L35:
	;
	goto L36
L36:
	;
	v107 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[1]))
	v113 = v107 + v89<<(uint(int32(13))%32) + int32(-8192)
	goto L33
L37:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113)+48))
	if v117 == v83 {
		goto L32
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	F_UnlockBuffer(m, v89)
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L7
	} else {
		goto L41
	}
L40:
	;
	goto L39
L41:
	;
	F_LockBufferInternal(m, v89, int32(3))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	v124 = int32(_a_F_btvacuumcleanup_0)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[2])) = v126 + int32(1)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v113)+28))
	if base.Ui32(v130) <= base.Ui32(int32(2)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v133 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v113)+64)) = uint8(v133)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+28)) = int32(3)
	v137 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v113)+12)) = uint16(v137)
	goto L45
L44:
	;
	goto L45
L45:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v113)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v113)+48)) = v83
	F_MarkBufferDirty(m, v89)
	mBase = m.M
	v143 = m.ExcPending
	if v143 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v80)+48))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+118)))
	if v145 != int32(112) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v113))) = base.I64_rotl(v186, int64(32))
	v190 = int32(_a_F_btvacuumcleanup_0)
	v192 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[2])) = v192 - int32(1)
	goto L32
L48:
	;
	v183 = F_XLogGetFakeLSN(m, v80)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L7
	} else {
		goto L59
	}
L49:
	;
	v149 = *(*int32)(unsafe.Add(mBase, _c_F_btvacuumcleanup[3]))
	if v149 <= int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v80)+32))
	if v152 != 0 {
		goto L48
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L7
	} else {
		goto L55
	}
L53:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v80)+40))
	if v153 != 0 {
		goto L48
	} else {
		goto L54
	}
L54:
	;
	goto L52
L55:
	;
	F_XLogRegisterBuffer(m, int32(0), v89, int32(14))
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L56
	}
L56:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v113)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v113)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+8)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v113)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+12)) = v164
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v113)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+16)) = v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v113)+44))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+24)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v86)+20)) = v168
	v171 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+64)))
	*(*uint8)(unsafe.Add(mBase, uint32(v86)+28)) = uint8(v171)
	F_XLogRegisterBufData(m, int32(0), v86+int32(4), int32(28))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L7
	} else {
		goto L57
	}
L57:
	;
	v181 = F_XLogInsert(m, int32(11), int32(224))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L58
	}
L58:
	;
	v186 = v181
	goto L47
L59:
	;
	v186 = v183
	goto L47
L60:
	;
	m.G0 = v86 + int32(32)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v202 != 0 {
		v209 = v76
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v203 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v204 = *(*float64)(unsafe.Add(mBase, uint32(v76)+8))
	if base.F64_lt(v203, v204) == int32(0) {
		v209 = v76
		goto L1
	} else {
		goto L62
	}
L62:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v76)+8)) = v203
	v209 = v76
	goto L1
}
func F_build_attrmap_by_name_if_req(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	v11 = F_build_attrmap_by_name(m, l0, l1, l2)
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v15 != v16 {
		v83 = v11
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return v83
L4:
	;
	v18 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v19 <= v18 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_pfree(m, v66)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L19
	}
L6:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v66 = v22
	goto L5
L7:
	;
	goto L8
L8:
	;
	v23 = int32(28)
	v29 = v18
	goto L9
L9:
	;
	v38 = v29 << (uint(int32(3)) % 32)
	v39 = l0 + v23 + v38
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+6)))
	if v40&int32(2) != 0 {
		v83 = v11
		goto L3
	} else {
		goto L11
	}
L10:
	;
	v66 = v45
	goto L5
L11:
	;
	v43 = int32(1)
	v44 = v29 + v43
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45+v29<<(uint(v43)%32)))))
	if v44 != v49 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	if base.B2i32(v40&int32(4) == int32(0))|v49 != 0 {
		v83 = v11
		goto L3
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	if v19 != v44 {
		v29 = v44
		goto L9
	} else {
		goto L18
	}
L15:
	;
	v56 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+2)))
	v57 = l1 + v23 + v38
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+2)))
	if v56 != v58 {
		v83 = v11
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+5)))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+5)))
	if v60 != v61 {
		v83 = v11
		goto L3
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	goto L10
L19:
	;
	F_pfree(m, v11)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v83 = int32(0)
	goto L3
}
func F_build_colinfo_names_hash(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v11 < int32(32) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v9 + int32(48)
	return
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = int64(274877907008)
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_build_colinfo_names_hash[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v17
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v24 = F_hash_create(m, int32(_a_F_build_colinfo_names_hash_0), base.I64_extend_i32_s(v20+v11), v9, int32(1048))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v24
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v27 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v32 = int32(0)
	v33 = v27
	goto L8
L6:
	;
	goto L7
L7:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if int32(0) < v63 {
		goto L15
	} else {
		goto L16
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v37+v32<<(uint(int32(2))%32))))
	if v41 == int32(0) {
		v52 = v33
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L7
L10:
	;
	v55 = v32 + int32(1)
	if v55 < v52 {
		v32 = v55
		v33 = v52
		goto L8
	} else {
		goto L14
	}
L11:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v44 == int32(0) {
		v52 = v33
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v49 = F_hash_search(m, v44, v41, int32(1), int32(0))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v52 = v51
	goto L10
L14:
	;
	goto L9
L15:
	;
	v68 = int32(0)
	v69 = v63
	goto L18
L16:
	;
	goto L17
L17:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v99 == int32(0) {
		goto L1
	} else {
		goto L25
	}
L18:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v73+v68<<(uint(int32(2))%32))))
	if v77 == int32(0) {
		v88 = v69
		goto L20
	} else {
		goto L21
	}
L19:
	;
	goto L17
L20:
	;
	v91 = v68 + int32(1)
	if v91 < v88 {
		v68 = v91
		v69 = v88
		goto L18
	} else {
		goto L24
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v80 == int32(0) {
		v88 = v69
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v85 = F_hash_search(m, v80, v77, int32(1), int32(0))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v88 = v87
	goto L20
L24:
	;
	goto L19
L25:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v102 <= int32(0) {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v107 = int32(0)
	goto L27
L27:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v112 != 0 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	goto L1
L29:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v99)+12))
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v113+v107<<(uint(int32(2))%32))))
	v120 = F_hash_search(m, v112, v117, int32(1), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L3
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v123 = v107 + int32(1)
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v99)+4))
	if v123 < v124 {
		v107 = v123
		goto L27
	} else {
		goto L33
	}
L32:
	;
	goto L31
L33:
	;
	goto L28
}
func F_build_joinrel_tlist(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v43 int64
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v233 int32
	_ = v233
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v260 int32
	_ = v260
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v322 int32
	_ = v322
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int64
	_ = v382
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v392 int32
	_ = v392
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int64
	_ = v491
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v538 int32
	_ = v538
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v622 int64
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v641 int64
	_ = v641
	var v643 int64
	_ = v643
	var v646 int64
	_ = v646
	var v648 int32
	_ = v648
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v21 = int64(*(*int32)(unsafe.Add(mBase, uint32(v20)+32)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l2)+40))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v23 == v7 {
		v641 = v21
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L14
	} else {
		goto L159
	}
L2:
	;
	v643 = int64(1073741823)
	if v643 <= v641 {
		goto L156
	} else {
		goto L157
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v26 <= int32(0) {
		v641 = v21
		goto L2
	} else {
		goto L4
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v41 = v7
	v43 = v21
	goto L5
L5:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v41<<(uint(int32(2))%32))))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	if v50 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v641 = v622
	goto L2
L7:
	;
	v625 = v41 + int32(1)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v23)+4))
	if v625 < v626 {
		v41 = v625
		v43 = v622
		goto L5
	} else {
		goto L154
	}
L8:
	;
	if v50 == int32(321) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v398 == int32(-4) {
		goto L104
	} else {
		goto L105
	}
L11:
	;
	v55 = F_find_placeholder_info(m, l0, v49)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L14
	} else {
		goto L100
	}
L14:
	;
	return
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55)+20))
	if v57 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	if v112 == int32(0) {
		v622 = v43
		goto L7
	} else {
		goto L30
	}
L17:
	;
	v112 = int32(0)
	goto L16
L18:
	;
	goto L19
L19:
	;
	v65 = int32(1)
	if v29 == int32(0) {
		v102 = v65
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v112 = v102
	goto L16
L21:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v69 < v68 {
		v102 = v65
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v71 = int32(1)
	if v68 <= v71 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v74 = v71
	goto L25
L24:
	;
	v74 = v68
	goto L25
L25:
	;
	v75 = int32(8)
	v80 = int32(0)
	goto L26
L26:
	;
	v87 = v80 << (uint(int32(2)) % 32)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v57+v75+v87)))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v29+v75+v87)))
	v94 = v89 & (v91 ^ int32(-1))
	v96 = base.B2i32(v94 != int32(0))
	if v94 != 0 {
		v102 = v96
		goto L20
	} else {
		goto L28
	}
L27:
	;
	v102 = v96
	goto L20
L28:
	;
	v98 = v80 + int32(1)
	if v98 != v74 {
		v80 = v98
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	if l5 != 0 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v115 = F_copyObjectImpl(m, v49)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L14
	} else {
		goto L34
	}
L32:
	;
	v363 = v49
	goto L33
L33:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v377 = *(*int32)(unsafe.Add(mBase, uint32(v376)+4))
	v378 = F_lappend(m, v377, v363)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L14
	} else {
		goto L99
	}
L34:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v117 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	if l4 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L36:
	;
	v120 = F_bms_is_member(m, v117, v29)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L14
	} else {
		goto L37
	}
L37:
	;
	if v120 == int32(0) {
		goto L35
	} else {
		goto L38
	}
L38:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v126 = int32(0)
	if v124 == v126 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	if v179 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L40:
	;
	v179 = int32(1)
	goto L39
L41:
	;
	goto L42
L42:
	;
	if v125 == int32(0) {
		v172 = v126
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v179 = v172
	goto L39
L44:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v124)+4))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)+4))
	if v136 < v135 {
		v172 = v126
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v138 = int32(1)
	if v135 <= v138 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v141 = v138
	goto L48
L47:
	;
	v141 = v135
	goto L48
L48:
	;
	v142 = int32(8)
	v147 = int32(0)
	goto L49
L49:
	;
	v154 = v147 << (uint(int32(2)) % 32)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v124+v142+v154)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v125+v142+v154)))
	v161 = v156 & (v158 ^ int32(-1))
	v163 = base.B2i32(v161 == int32(0))
	if v161 != 0 {
		v172 = v163
		goto L43
	} else {
		goto L51
	}
L50:
	;
	v172 = v163
	goto L43
L51:
	;
	v165 = v147 + int32(1)
	if v165 != v141 {
		v147 = v165
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v182 != int32(2) {
		goto L35
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v245 = F_bms_add_member(m, v243, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L14
	} else {
		goto L72
	}
L56:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v187 = int32(0)
	if v185 == v187 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	if v240 == int32(0) {
		goto L35
	} else {
		goto L71
	}
L58:
	;
	v240 = int32(1)
	goto L57
L59:
	;
	goto L60
L60:
	;
	if v186 == int32(0) {
		v233 = v187
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v240 = v233
	goto L57
L62:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v197 < v196 {
		v233 = v187
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v199 = int32(1)
	if v196 <= v199 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v202 = v199
	goto L66
L65:
	;
	v202 = v196
	goto L66
L66:
	;
	v203 = int32(8)
	v208 = int32(0)
	goto L67
L67:
	;
	v215 = v208 << (uint(int32(2)) % 32)
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v185+v203+v215)))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v186+v203+v215)))
	v222 = v217 & (v219 ^ int32(-1))
	v224 = base.B2i32(v222 == int32(0))
	if v222 != 0 {
		v233 = v224
		goto L61
	} else {
		goto L69
	}
L68:
	;
	v233 = v224
	goto L61
L69:
	;
	v226 = v208 + int32(1)
	if v226 != v202 {
		v208 = v226
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	goto L55
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v245
	goto L35
L73:
	;
	v354 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v356 = F_bms_intersect(m, v355, v29)
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
		goto L14
	} else {
		goto L97
	}
L74:
	;
	v250 = int32(0)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v251 <= v250 {
		goto L73
	} else {
		goto L75
	}
L75:
	;
	v260 = v250
	goto L76
L76:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v115)+8))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v270+v260<<(uint(int32(2))%32))))
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v274)+16))
	v276 = int32(0)
	if v269 == v276 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	goto L73
L78:
	;
	if v329 != 0 {
		goto L92
	} else {
		goto L93
	}
L79:
	;
	v329 = int32(1)
	goto L78
L80:
	;
	goto L81
L81:
	;
	if v275 == int32(0) {
		v322 = v276
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v329 = v322
	goto L78
L83:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v269)+4))
	v286 = *(*int32)(unsafe.Add(mBase, uint32(v275)+4))
	if v286 < v285 {
		v322 = v276
		goto L82
	} else {
		goto L84
	}
L84:
	;
	v288 = int32(1)
	if v285 <= v288 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v291 = v288
	goto L87
L86:
	;
	v291 = v285
	goto L87
L87:
	;
	v292 = int32(8)
	v297 = int32(0)
	goto L88
L88:
	;
	v304 = v297 << (uint(int32(2)) % 32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v269+v292+v304)))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v275+v292+v304)))
	v311 = v306 & (v308 ^ int32(-1))
	v313 = base.B2i32(v311 == int32(0))
	if v311 != 0 {
		v322 = v313
		goto L82
	} else {
		goto L90
	}
L89:
	;
	v322 = v313
	goto L82
L90:
	;
	v315 = v297 + int32(1)
	if v315 != v291 {
		v297 = v315
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v274)+24))
	v332 = F_bms_add_member(m, v330, v331)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L14
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v336 = v260 + int32(1)
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v336 < v337 {
		v260 = v336
		goto L76
	} else {
		goto L96
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v332
	goto L94
L96:
	;
	goto L77
L97:
	;
	v358 = F_bms_join(m, v354, v356)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L14
	} else {
		goto L98
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v115)+12)) = v358
	v363 = v115
	goto L33
L99:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v380)+4)) = v378
	v382 = int64(*(*int32)(unsafe.Add(mBase, uint32(v55)+24)))
	v622 = v43 + v382
	goto L7
L100:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v49)))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v388
	F_errmsg_internal(m, int32(_a_F_build_joinrel_tlist_0), v18)
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L14
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(_a_F_build_joinrel_tlist_1), int32(1331), int32(_a_F_build_joinrel_tlist_2))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L14
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	v491 = int64(*(*int32)(unsafe.Add(mBase, uint32(v490))))
	if l5 == int32(0) {
		v589 = v49
		goto L124
	} else {
		goto L125
	}
L104:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v401)+12))
	v403 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+8)))
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v402+v403<<(uint(int32(2))%32)-int32(4))))
	v490 = v409 + int32(8)
	goto L103
L105:
	;
	goto L106
L106:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if base.Ui32(v412) <= base.Ui32(v398) {
		goto L1
	} else {
		goto L107
	}
L107:
	;
	v414 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v414+v398<<(uint(int32(2))%32))))
	if v418 == int32(0) {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v421 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+8)))
	v422 = int32(*(*int16)(unsafe.Add(mBase, uint32(v418)+88)))
	v425 = (v421 - v422) << (uint(int32(2)) % 32)
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v418)+92))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v425+v426)))
	if v428 == int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	if v483 == int32(0) {
		v622 = v43
		goto L7
	} else {
		goto L123
	}
L110:
	;
	v483 = int32(0)
	goto L109
L111:
	;
	goto L112
L112:
	;
	v436 = int32(1)
	if v29 == int32(0) {
		v473 = v436
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v483 = v473
	goto L109
L114:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v428)+4))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v440 < v439 {
		v473 = v436
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v442 = int32(1)
	if v439 <= v442 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v445 = v442
	goto L118
L117:
	;
	v445 = v439
	goto L118
L118:
	;
	v446 = int32(8)
	v451 = int32(0)
	goto L119
L119:
	;
	v458 = v451 << (uint(int32(2)) % 32)
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v428+v446+v458)))
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v29+v446+v458)))
	v465 = v460 & (v462 ^ int32(-1))
	v467 = base.B2i32(v465 != int32(0))
	if v465 != 0 {
		v473 = v467
		goto L113
	} else {
		goto L121
	}
L120:
	;
	v473 = v467
	goto L113
L121:
	;
	v469 = v451 + int32(1)
	if v469 != v445 {
		v451 = v469
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v418)+96))
	v490 = v486 + v425
	goto L103
L124:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	v605 = F_lappend(m, v604, v589)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L14
	} else {
		goto L153
	}
L125:
	;
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v494 == int32(-4) {
		v589 = v49
		goto L124
	} else {
		goto L126
	}
L126:
	;
	v497 = F_copyObjectImpl(m, v49)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L14
	} else {
		goto L127
	}
L127:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	if v499 == int32(0) {
		goto L128
	} else {
		goto L129
	}
L128:
	;
	if l4 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L129:
	;
	v502 = F_bms_is_member(m, v499, v29)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L14
	} else {
		goto L130
	}
L130:
	;
	if v502 == int32(0) {
		goto L128
	} else {
		goto L131
	}
L131:
	;
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v508 = F_bms_is_member(m, v506, v507)
	mBase = m.M
	v509 = m.ExcPending
	if v509 != 0 {
		goto L14
	} else {
		goto L132
	}
L132:
	;
	if v508 == int32(0) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	if v512 != int32(2) {
		goto L128
	} else {
		goto L136
	}
L134:
	;
	goto L135
L135:
	;
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v497)+24))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v523 = F_bms_add_member(m, v521, v522)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L14
	} else {
		goto L139
	}
L136:
	;
	v515 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v517 = F_bms_is_member(m, v515, v516)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L14
	} else {
		goto L137
	}
L137:
	;
	if v517 == int32(0) {
		goto L128
	} else {
		goto L138
	}
L138:
	;
	goto L135
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+24)) = v523
	goto L128
L140:
	;
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v497)+24))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l3)+32))
	v582 = F_bms_intersect(m, v581, v29)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L14
	} else {
		goto L151
	}
L141:
	;
	v528 = int32(0)
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v529 <= v528 {
		goto L140
	} else {
		goto L142
	}
L142:
	;
	v538 = v528
	goto L143
L143:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v497)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v548+v538<<(uint(int32(2))%32))))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v552)+16))
	v554 = F_bms_is_member(m, v547, v553)
	mBase = m.M
	v555 = m.ExcPending
	if v555 != 0 {
		goto L14
	} else {
		goto L145
	}
L144:
	;
	goto L140
L145:
	;
	if v554 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v497)+24))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v552)+24))
	v558 = F_bms_add_member(m, v556, v557)
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L14
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v562 = v538 + int32(1)
	v563 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v562 < v563 {
		v538 = v562
		goto L143
	} else {
		goto L150
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+24)) = v558
	goto L148
L150:
	;
	goto L144
L151:
	;
	v584 = F_bms_join(m, v580, v582)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L14
	} else {
		goto L152
	}
L152:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v497)+24)) = v584
	v589 = v497
	goto L124
L153:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v607)+4)) = v605
	v622 = v43 + v491
	goto L7
L154:
	;
	goto L6
L155:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v648)+32)) = base.I32_wrap_i64(v646)
	m.G0 = v18 + int32(32)
	return
L156:
	;
	v646 = v643
	goto L158
L157:
	;
	v646 = v641
	goto L158
L158:
	;
	goto L155
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v398
	F_errmsg_internal(m, int32(_a_F_build_joinrel_tlist_3), v18+int32(16))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L14
	} else {
		goto L160
	}
L160:
	;
	F_errfinish(m, int32(_a_F_build_joinrel_tlist_1), int32(556), int32(_a_F_build_joinrel_tlist_4))
	mBase = m.M
	v668 = m.ExcPending
	if v668 != 0 {
		goto L14
	} else {
		goto L161
	}
L161:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_build_paths_for_OR(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
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
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
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
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	v5 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(144)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+116))
	if v17 == v5 {
		v201 = v5
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v15 + int32(144)
	return v201
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 <= int32(0) {
		v201 = v5
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = v5
	v33 = v5
	v34 = v5
	goto L4
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v35+v32<<(uint(int32(2))%32))))
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+111)))
	if v40 == int32(0) {
		v184 = v33
		v185 = v34
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v201 = v185
	goto L1
L6:
	;
	v187 = v32 + int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v187 < v188 {
		v32 = v187
		v33 = v184
		v34 = v185
		goto L4
	} else {
		goto L36
	}
L7:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	if v45 == v43 {
		v68 = v33
		v69 = v43
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v72 = int32(0)
	base.MemoryFill(m, v15+int32(12), v72, int32(132))
	if l2 == v72 {
		v110 = v43
		goto L19
	} else {
		goto L20
	}
L9:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+100)))
	if v49 != 0 {
		v68 = v33
		v69 = int32(0)
		goto L8
	} else {
		goto L10
	}
L10:
	;
	if v33 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v55 = v33
	v56 = v45
	goto L13
L12:
	;
	v50 = F_list_concat_copy(m, l2, l3)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	v58 = F_predicate_implied_by(m, v56, v55, int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L14
	} else {
		goto L16
	}
L14:
	;
	return int32(0)
L15:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	v55 = v50
	v56 = v54
	goto L13
L16:
	;
	if v58 == int32(0) {
		v184 = v55
		v185 = v34
		goto L6
	} else {
		goto L17
	}
L17:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v39)+88))
	v64 = F_predicate_implied_by(m, v62, l3, int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L14
	} else {
		goto L18
	}
L18:
	;
	v68 = v55
	v69 = v64 ^ int32(1)
	goto L8
L19:
	;
	if (v110|v69)&int32(1) == int32(0) {
		v184 = v68
		v185 = v34
		goto L6
	} else {
		goto L26
	}
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v77 <= int32(0) {
		v110 = v43
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v84 = v43
	goto L22
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v84<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v96, v39, v15+int32(12))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L14
	} else {
		goto L24
	}
L23:
	;
	v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+12)))
	v110 = v105
	goto L19
L24:
	;
	v102 = v84 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v102 < v103 {
		v84 = v102
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if l3 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v170 = F_build_index_paths(m, l0, l1, v39, v15+int32(12), v69, int32(1), int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L14
	} else {
		goto L34
	}
L28:
	;
	v125 = int32(0)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v126 <= v125 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v133 = v125
	goto L30
L30:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v141+v133<<(uint(int32(2))%32))))
	F_match_clause_to_index(m, l0, v145, v39, v15+int32(12))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L14
	} else {
		goto L32
	}
L31:
	;
	goto L27
L32:
	;
	v151 = v133 + int32(1)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v151 < v152 {
		v133 = v151
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v172 = F_list_concat(m, v34, v170)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L14
	} else {
		goto L35
	}
L35:
	;
	v184 = v68
	v185 = v172
	goto L6
L36:
	;
	goto L5
}
func F_build_pgstattuple_type(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v14 float64
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int64
	_ = v33
	var v37 int64
	_ = v37
	var v39 float64
	_ = v39
	var v41 float64
	_ = v41
	var v43 int64
	_ = v43
	var v48 int64
	_ = v48
	var v54 float64
	_ = v54
	var v55 float64
	_ = v55
	var v56 float64
	_ = v56
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int64
	_ = v91
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int64
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int64
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int64
	_ = v137
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	v14 = float64(0)
	v17 = m.G0
	v19 = v17 - int32(3024)
	m.G0 = v19
	v24 = F_get_call_result_type(m, l1, int32(0), v19+int32(140))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		return int64(0)
	} else {
		if v24 == int32(1) {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+140))
			v31 = F_TupleDescGetAttInMetadata(m, v30)
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return int64(0)
			} else {
				v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
				if v33 == int64(0) {
					v54 = v14
					v55 = v14
					v56 = float64(0)
				} else {
					v37 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
					v39 = float64(100)
					v41 = base.F64_convert_i64_u(v33)
					v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
					v48 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
					v54 = base.F64_div(base.F64_mul(base.F64_convert_i64_u(v37), v39), v41)
					v55 = base.F64_div(base.F64_mul(base.F64_convert_i64_u(v43), v39), v41)
					v56 = base.F64_div(base.F64_mul(base.F64_convert_i64_u(v48), v39), v41)
				}
				*(*int64)(unsafe.Add(mBase, uint32(v19)+128)) = v33
				v59 = v19 + int32(2656)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+3008)) = v59
				v62 = v19 + int32(2342)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+3004)) = v62
				v65 = v19 + int32(2028)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+3000)) = v65
				v68 = v19 + int32(1714)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2996)) = v68
				v71 = v19 + int32(1400)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2992)) = v71
				v74 = v19 + int32(1086)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2988)) = v74
				v77 = v19 + int32(772)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2984)) = v77
				v80 = v19 + int32(458)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2980)) = v80
				v83 = v19 + int32(144)
				*(*int32)(unsafe.Add(mBase, uint32(v19)+2976)) = v83
				v89 = F_pg_snprintf(m, v83, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(128))
				mBase = m.M
				v90 = m.ExcPending
				if v90 != 0 {
					return int64(0)
				} else {
					v91 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v19)+112)) = v91
					v97 = F_pg_snprintf(m, v80, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(112))
					mBase = m.M
					v98 = m.ExcPending
					if v98 != 0 {
						return int64(0)
					} else {
						v99 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
						*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v99
						v105 = F_pg_snprintf(m, v77, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(96))
						mBase = m.M
						v106 = m.ExcPending
						if v106 != 0 {
							return int64(0)
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v19)+80)) = v56
							v112 = F_pg_snprintf(m, v74, int32(314), int32(_a_F_build_pgstattuple_type_1), v19+int32(80))
							mBase = m.M
							v113 = m.ExcPending
							if v113 != 0 {
								return int64(0)
							} else {
								v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v114
								v120 = F_pg_snprintf(m, v71, int32(314), int32(_a_F_build_pgstattuple_type_0), v19-int32(-64))
								mBase = m.M
								v121 = m.ExcPending
								if v121 != 0 {
									return int64(0)
								} else {
									v122 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
									*(*int64)(unsafe.Add(mBase, uint32(v19)+48)) = v122
									v128 = F_pg_snprintf(m, v68, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(48))
									mBase = m.M
									v129 = m.ExcPending
									if v129 != 0 {
										return int64(0)
									} else {
										*(*float64)(unsafe.Add(mBase, uint32(v19)+32)) = v55
										v135 = F_pg_snprintf(m, v65, int32(314), int32(_a_F_build_pgstattuple_type_1), v19+int32(32))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
											return int64(0)
										} else {
											v137 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
											*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v137
											v143 = F_pg_snprintf(m, v62, int32(314), int32(_a_F_build_pgstattuple_type_0), v19+int32(16))
											mBase = m.M
											v144 = m.ExcPending
											if v144 != 0 {
												return int64(0)
											} else {
												*(*float64)(unsafe.Add(mBase, uint32(v19))) = v54
												v148 = F_pg_snprintf(m, v59, int32(314), int32(_a_F_build_pgstattuple_type_1), v19)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int64(0)
												} else {
													v152 = F_BuildTupleFromCStrings(m, v31, v19+int32(2976))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int64(0)
													} else {
														v154 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
														v155 = F_HeapTupleHeaderGetDatum(m, v154)
														mBase = m.M
														v156 = m.ExcPending
														if v156 != 0 {
															return int64(0)
														} else {
															m.G0 = v19 + int32(3024)
															return v155
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
			v164 = m.ExcPending
			if v164 != 0 {
				return int64(0)
			} else {
				F_errmsg_internal(m, int32(_a_F_build_pgstattuple_type_2), int32(0))
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_build_pgstattuple_type_3), int32(109), int32(_a_F_build_pgstattuple_type_4))
					mBase = m.M
					v173 = m.ExcPending
					if v173 != 0 {
						return int64(0)
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
func F_build_sorted_items(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v111 int32
	_ = v111
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v156 int32
	_ = v156
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v185 int32
	_ = v185
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v221 int32
	_ = v221
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v266 int64
	_ = v266
	var v267 int32
	_ = v267
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int64
	_ = v284
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v324 int32
	_ = v324
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v358 int32
	_ = v358
	var v364 int32
	_ = v364
	v6 = int32(0)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = l3 * v19
	v30 = F_palloc0(m, (v19*int32(12)+int32(7))&int32(-8)+v26*int32(9))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if int32(0) < v34 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v43 = v30 + (v34*int32(12)+int32(7))&int32(-8)
	v53 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v100 = F_palloc_mul(m, int32(4), v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v68 = v30 + v53*int32(12)
	v69 = l3 * v53
	*(*int32)(unsafe.Add(mBase, uint32(v68)+4)) = v43 + v26<<(uint(int32(3))%32) + v69
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v43 + v69<<(uint(int32(3))%32)
	v77 = v53 + int32(1)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v77 < v78 {
		v53 = v77
		goto L6
	} else {
		goto L8
	}
L7:
	;
	goto L5
L8:
	;
	goto L7
L9:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if int32(0) < v102 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v111 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v156 <= int32(0) {
		goto L19
	} else {
		goto L20
	}
L13:
	;
	v125 = v111 << (uint(int32(2)) % 32)
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v127+v125)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
	v131 = F_get_typlen(m, v130)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v100+v125))) = v131
	v135 = v111 + int32(1)
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v135 < v136 {
		v111 = v135
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	F_qsort_interruptible(m, v30, v324, int32(12), int32(1140), l2)
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L46
	}
L18:
	;
	F_pfree(m, v30)
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L1
	} else {
		goto L45
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	goto L18
L20:
	;
	goto L21
L21:
	;
	v171 = v6
	v176 = v6
	goto L22
L22:
	;
	if base.B2i32(l3 <= int32(0)) == int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v324
	if v324 != 0 {
		goto L17
	} else {
		goto L44
	}
L24:
	;
	v335 = v176 + int32(1)
	v336 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v335 < v336 {
		v171 = v324
		v176 = v335
		goto L22
	} else {
		goto L43
	}
L25:
	;
	v185 = v30 + v171*int32(12)
	v198 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v324 = v171 + int32(1)
	goto L24
L28:
	;
	v207 = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v208 <= v207 {
		v247 = v207
		goto L30
	} else {
		goto L31
	}
L29:
	;
	goto L27
L30:
	;
	v261 = v247 << (uint(int32(2)) % 32)
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261+v262)))
	v266 = *(*int64)(unsafe.Add(mBase, uint32(v264+v176<<(uint(int32(3))%32))))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v267+v261)))
	v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269+v176))))
	if v271 != 0 {
		v284 = v266
		goto L36
	} else {
		goto L37
	}
L31:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4+v198<<(uint(int32(1))%32)))))
	v221 = v207
	goto L32
L32:
	;
	v237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211+v221<<(uint(int32(1))%32)))))
	if v215 == v237 {
		v247 = v221
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v247 = v208
	goto L30
L34:
	;
	v240 = v221 + int32(1)
	if v240 != v208 {
		v221 = v240
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v185)))
	*(*int64)(unsafe.Add(mBase, uint32(v285+v198<<(uint(int32(3))%32)))) = v284
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	*(*uint8)(unsafe.Add(mBase, uint32(v290+v198))) = uint8(v271)
	v294 = v198 + int32(1)
	if v294 != l3 {
		v198 = v294
		goto L28
	} else {
		goto L42
	}
L37:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v261+v100)))
	if v273 != int32(-1) {
		v284 = v266
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v276 = F_toast_raw_datum_size(m, v266)
	mBase = m.M
	v277 = m.ExcPending
	if v277 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v276) {
		v324 = v171
		goto L24
	} else {
		goto L40
	}
L40:
	;
	v281 = F_pg_detoast_datum(m, base.I32_wrap_i64(v266))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L41
	}
L41:
	;
	v284 = base.I64_extend_i32_u(v281)
	goto L36
L42:
	;
	goto L29
L43:
	;
	goto L23
L44:
	;
	goto L18
L45:
	;
	return int32(0)
L46:
	;
	return v30
}
func F_buildoidvector(m *base.Module, l0 int32, l1 int32) int32 {
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14241(m, l0, l1, int32(26), int32(2))
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_byteacmp(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
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
	var v162 int32
	_ = v162
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v76 = int32(1)
	if v16&v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = int32(1)
	if v46&v63 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = int32(16)
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(4)
	goto L24
L23:
	;
	v62 = v55
	goto L24
L24:
	;
	v75 = v62
	goto L15
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v80 = v76
	goto L28
L27:
	;
	v80 = int32(4)
	goto L28
L28:
	;
	v81 = v9 + v80
	v82 = int32(1)
	if v46&v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = int32(4)
	goto L31
L31:
	;
	v87 = v14 + v86
	v88 = base.B2i32(v45 < v75)
	if v45 < v75 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v45
	goto L34
L33:
	;
	v89 = v75
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v152 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v81|v87)&int32(3) != 0 {
		v120 = v81
		v121 = v87
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v81
	v114 = v87
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v81
	v98 = v87
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
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
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v156 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
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
	if v151 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	goto L59
L61:
	;
	v162 = v151
	goto L63
L62:
	;
	v162 = base.B2i32(v75 < v45) - v88
	goto L63
L63:
	;
	return base.I64_extend_i32_s(v162)
}
func F_byteaeq(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_toast_raw_datum_size(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return base.I64_extend_i32_u(v110)
L2:
	;
	return int64(0)
L3:
	;
	v13 = F_toast_raw_datum_size(m, v7)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L2
	} else {
		goto L4
	}
L4:
	;
	if v9 != v13 {
		v110 = int32(0)
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v17 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v8))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v20 = F_pg_detoast_datum_packed(m, base.I32_wrap_i64(v7))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	v22 = int32(1)
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if v24&v22 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v27 = v22
	goto L10
L9:
	;
	v27 = int32(4)
	goto L10
L10:
	;
	v28 = v17 + v27
	v29 = int32(1)
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20))))
	if v31&v29 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v34 = v29
	goto L13
L12:
	;
	v34 = int32(4)
	goto L13
L13:
	;
	v35 = v20 + v34
	v36 = int32(4)
	v37 = v9 - v36
	if base.Ui32(v36) <= base.Ui32(v37) {
		goto L17
	} else {
		goto L18
	}
L14:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v100 != v17 {
		goto L32
	} else {
		goto L33
	}
L15:
	;
	v99 = int32(0)
	goto L14
L16:
	;
	v73 = v68
	v74 = v69
	v75 = v70
	goto L26
L17:
	;
	if (v28|v35)&int32(3) != 0 {
		v68 = v28
		v69 = v35
		v70 = v37
		goto L16
	} else {
		goto L20
	}
L18:
	;
	v61 = v28
	v62 = v35
	v63 = v37
	goto L19
L19:
	;
	if v63 == int32(0) {
		goto L15
	} else {
		goto L25
	}
L20:
	;
	v45 = v28
	v46 = v35
	v47 = v37
	goto L21
L21:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v50 != v51 {
		v68 = v45
		v69 = v46
		v70 = v47
		goto L16
	} else {
		goto L23
	}
L22:
	;
	v61 = v56
	v62 = v54
	v63 = v58
	goto L19
L23:
	;
	v53 = int32(4)
	v54 = v46 + v53
	v56 = v45 + v53
	v58 = v47 - v53
	if base.Ui32(int32(3)) < base.Ui32(v58) {
		v45 = v56
		v46 = v54
		v47 = v58
		goto L21
	} else {
		goto L24
	}
L24:
	;
	goto L22
L25:
	;
	v68 = v61
	v69 = v62
	v70 = v63
	goto L16
L26:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 == v79 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v99 = v78 - v79
	goto L14
L28:
	;
	v81 = int32(1)
	v86 = v75 - v81
	if v86 != 0 {
		v73 = v73 + v81
		v74 = v74 + v81
		v75 = v86
		goto L26
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	goto L27
L31:
	;
	goto L15
L32:
	;
	F_pfree(m, v17)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L2
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v105 = base.B2i32(v99 == int32(0))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v20 == v106 {
		v110 = v105
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	F_pfree(m, v20)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L2
	} else {
		goto L37
	}
L37:
	;
	v110 = v105
	goto L1
}
func F_bytealt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
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
	var v160 int32
	_ = v160
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v76 = int32(1)
	if v16&v76 != 0 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v52 == int32(18) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v63 = int32(1)
	if v46&v63 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v63)%32)) - v63
		goto L15
	} else {
		goto L25
	}
L19:
	;
	v55 = int32(16)
	goto L21
L20:
	;
	v55 = int32(0)
	goto L21
L21:
	;
	if base.Ui32((v52-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v62 = int32(4)
	goto L24
L23:
	;
	v62 = v55
	goto L24
L24:
	;
	v75 = v62
	goto L15
L25:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v69)>>(uint(int32(2))%32)) - int32(4)
	goto L15
L26:
	;
	v80 = v76
	goto L28
L27:
	;
	v80 = int32(4)
	goto L28
L28:
	;
	v81 = v9 + v80
	v82 = int32(1)
	if v46&v82 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v86 = v82
	goto L31
L30:
	;
	v86 = int32(4)
	goto L31
L31:
	;
	v87 = v14 + v86
	v88 = base.B2i32(v45 < v75)
	if v45 < v75 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v89 = v45
	goto L34
L33:
	;
	v89 = v75
	goto L34
L34:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v89) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v152 != v9 {
		goto L53
	} else {
		goto L54
	}
L36:
	;
	v151 = int32(0)
	goto L35
L37:
	;
	v125 = v120
	v126 = v121
	v127 = v122
	goto L47
L38:
	;
	if (v81|v87)&int32(3) != 0 {
		v120 = v81
		v121 = v87
		v122 = v89
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v113 = v81
	v114 = v87
	v115 = v89
	goto L40
L40:
	;
	if v115 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v97 = v81
	v98 = v87
	v99 = v89
	goto L42
L42:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v97)))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	if v102 != v103 {
		v120 = v97
		v121 = v98
		v122 = v99
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v113 = v108
	v114 = v106
	v115 = v110
	goto L40
L44:
	;
	v105 = int32(4)
	v106 = v98 + v105
	v108 = v97 + v105
	v110 = v99 - v105
	if base.Ui32(int32(3)) < base.Ui32(v110) {
		v97 = v108
		v98 = v106
		v99 = v110
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v120 = v113
	v121 = v114
	v122 = v115
	goto L37
L47:
	;
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125))))
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	if v130 == v131 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v151 = v130 - v131
	goto L35
L49:
	;
	v133 = int32(1)
	v138 = v127 - v133
	if v138 != 0 {
		v125 = v125 + v133
		v126 = v126 + v133
		v127 = v138
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	F_pfree(m, v9)
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
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v156 != v14 {
		goto L57
	} else {
		goto L58
	}
L56:
	;
	goto L55
L57:
	;
	F_pfree(m, v14)
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
	v160 = int32(0)
	return base.I64_extend_i32_u(base.B2i32(v151 == v160)&v88 | base.B2i32(v151 < v160))
L60:
	;
	goto L59
}
func F_byteaout(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v172 int64
	_ = v172
	var v176 int32
	_ = v176
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v191 int32
	_ = v191
	var v198 int64
	_ = v198
	var v201 int64
	_ = v201
	var v202 int64
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v217 int64
	_ = v217
	var v221 int32
	_ = v221
	var v228 int64
	_ = v228
	var v231 int64
	_ = v231
	var v240 int64
	_ = v240
	var v250 int64
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, _c_F_byteaout[0]))
		switch v19 {
		case 0:
			v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			v120 = int32(1)
			v121 = v119 & v120
			if v119 == v120 {
				v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if base.Ui32((v125-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v151 = int32(4)
					if v121 != 0 {
						v154 = int32(1)
					} else {
						v154 = int32(4)
					}
					v155 = v14 + v154
					if v151 == int32(1) {
						v210 = v155
						v217 = int64(1)
						v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
						if base.Ui32((v221-int32(127))&int32(255)) < base.Ui32(int32(161)) {
							v228 = int64(4)
						} else {
							v228 = int64(1)
						}
						if v221 == int32(92) {
							v231 = int64(2)
						} else {
							v231 = v228
						}
						v240 = v217 + v231
					} else {
						v165 = v155
						v167 = int32(0)
						v172 = int64(1)
						for {
							v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
							if base.Ui32((v176-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v183 = int64(4)
							} else {
								v183 = int64(1)
							}
							if v176 == int32(92) {
								v186 = int64(2)
							} else {
								v186 = v183
							}
							v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
							if base.Ui32((v191-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v198 = int64(4)
							} else {
								v198 = int64(1)
							}
							if v191 == int32(92) {
								v201 = int64(2)
							} else {
								v201 = v198
							}
							v202 = v172 + v186 + v201
							v203 = int32(2)
							v204 = v165 + v203
							v206 = v167 + v203
							if v206 != v151&int32(-2) {
								v165 = v204
								v167 = v206
								v172 = v202
								continue
							} else {
								break
							}
							break
						}
						if v151&int32(1) == int32(0) {
							v240 = v202
						} else {
							v210 = v204
							v217 = v202
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
							if base.Ui32((v221-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v228 = int64(4)
							} else {
								v228 = int64(1)
							}
							if v221 == int32(92) {
								v231 = int64(2)
							} else {
								v231 = v228
							}
							v240 = v217 + v231
						}
					}
					if base.Ui64(int64(1073741824)) <= base.Ui64(v240) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v358 = m.ExcPending
						if v358 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v361 = m.ExcPending
							if v361 != 0 {
								return int64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_byteaout_0), int32(0))
								mBase = m.M
								v365 = m.ExcPending
								if v365 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_byteaout_1), int32(316), int32(_a_F_byteaout_2))
									mBase = m.M
									v370 = m.ExcPending
									if v370 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v250 = v240
						v252 = F_palloc(m, base.I32_wrap_i64(v250))
						mBase = m.M
						v253 = m.ExcPending
						if v253 != 0 {
							return int64(0)
						} else {
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
							v255 = int32(1)
							v256 = v254 & v255
							if v254 == v255 {
								v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
								if base.Ui32((v260-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v285 = int32(4)
									if v256 != 0 {
										v288 = int32(1)
									} else {
										v288 = int32(4)
									}
									v290 = v252
									v292 = v14 + v288
									v294 = v285
									for {
										v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
										if v298 == int32(92) {
											v301 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
											v335 = v290 + int32(2)
										} else {
											if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v311 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
												v313 = int32(7)
												v315 = int32(48)
												v316 = v298&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
												v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
												v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
												v335 = v290 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
												v335 = v290 + int32(1)
											}
										}
										v336 = int32(1)
										v339 = v294 - v336
										if v339 != 0 {
											v290 = v335
											v292 = v292 + v336
											v294 = v339
											continue
										} else {
											break
										}
										break
									}
									v340 = v335
									v343 = v252
								} else {
									if v260 == int32(18) {
										v271 = int32(16)
									} else {
										v271 = int32(0)
									}
									v283 = v271
									if v283 != 0 {
										v285 = v283
										if v256 != 0 {
											v288 = int32(1)
										} else {
											v288 = int32(4)
										}
										v290 = v252
										v292 = v14 + v288
										v294 = v285
										for {
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
											if v298 == int32(92) {
												v301 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
												v335 = v290 + int32(2)
											} else {
												if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v311 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
													v313 = int32(7)
													v315 = int32(48)
													v316 = v298&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
													v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
													v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
													v335 = v290 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
													v335 = v290 + int32(1)
												}
											}
											v336 = int32(1)
											v339 = v294 - v336
											if v339 != 0 {
												v290 = v335
												v292 = v292 + v336
												v294 = v339
												continue
											} else {
												break
											}
											break
										}
										v340 = v335
										v343 = v252
									} else {
										v340 = v252
										v343 = v252
									}
								}
							} else {
								v272 = int32(1)
								if v256 != 0 {
									v283 = int32(base.Ui32(v254)>>(uint(v272)%32)) - v272
								} else {
									v276 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
									v283 = int32(base.Ui32(v276)>>(uint(int32(2))%32)) - int32(4)
								}
								if v283 != 0 {
									v285 = v283
									if v256 != 0 {
										v288 = int32(1)
									} else {
										v288 = int32(4)
									}
									v290 = v252
									v292 = v14 + v288
									v294 = v285
									for {
										v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
										if v298 == int32(92) {
											v301 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
											v335 = v290 + int32(2)
										} else {
											if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v311 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
												v313 = int32(7)
												v315 = int32(48)
												v316 = v298&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
												v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
												v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
												v335 = v290 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
												v335 = v290 + int32(1)
											}
										}
										v336 = int32(1)
										v339 = v294 - v336
										if v339 != 0 {
											v290 = v335
											v292 = v292 + v336
											v294 = v339
											continue
										} else {
											break
										}
										break
									}
									v340 = v335
									v343 = v252
								} else {
									v340 = v252
									v343 = v252
								}
							}
							v348 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v340))) = uint8(v348)
							m.G0 = v11 + int32(16)
							return base.I64_extend_i32_u(v343)
						}
					}
				} else {
					if v125 == int32(18) {
						v136 = int32(16)
					} else {
						v136 = int32(0)
					}
					v148 = v136
					if v148 != 0 {
						v151 = v148
						if v121 != 0 {
							v154 = int32(1)
						} else {
							v154 = int32(4)
						}
						v155 = v14 + v154
						if v151 == int32(1) {
							v210 = v155
							v217 = int64(1)
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
							if base.Ui32((v221-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v228 = int64(4)
							} else {
								v228 = int64(1)
							}
							if v221 == int32(92) {
								v231 = int64(2)
							} else {
								v231 = v228
							}
							v240 = v217 + v231
						} else {
							v165 = v155
							v167 = int32(0)
							v172 = int64(1)
							for {
								v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
								if base.Ui32((v176-int32(127))&int32(255)) < base.Ui32(int32(161)) {
									v183 = int64(4)
								} else {
									v183 = int64(1)
								}
								if v176 == int32(92) {
									v186 = int64(2)
								} else {
									v186 = v183
								}
								v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
								if base.Ui32((v191-int32(127))&int32(255)) < base.Ui32(int32(161)) {
									v198 = int64(4)
								} else {
									v198 = int64(1)
								}
								if v191 == int32(92) {
									v201 = int64(2)
								} else {
									v201 = v198
								}
								v202 = v172 + v186 + v201
								v203 = int32(2)
								v204 = v165 + v203
								v206 = v167 + v203
								if v206 != v151&int32(-2) {
									v165 = v204
									v167 = v206
									v172 = v202
									continue
								} else {
									break
								}
								break
							}
							if v151&int32(1) == int32(0) {
								v240 = v202
							} else {
								v210 = v204
								v217 = v202
								v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
								if base.Ui32((v221-int32(127))&int32(255)) < base.Ui32(int32(161)) {
									v228 = int64(4)
								} else {
									v228 = int64(1)
								}
								if v221 == int32(92) {
									v231 = int64(2)
								} else {
									v231 = v228
								}
								v240 = v217 + v231
							}
						}
						if base.Ui64(int64(1073741824)) <= base.Ui64(v240) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v358 = m.ExcPending
							if v358 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(261))
								mBase = m.M
								v361 = m.ExcPending
								if v361 != 0 {
									return int64(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_byteaout_0), int32(0))
									mBase = m.M
									v365 = m.ExcPending
									if v365 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_byteaout_1), int32(316), int32(_a_F_byteaout_2))
										mBase = m.M
										v370 = m.ExcPending
										if v370 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v250 = v240
							v252 = F_palloc(m, base.I32_wrap_i64(v250))
							mBase = m.M
							v253 = m.ExcPending
							if v253 != 0 {
								return int64(0)
							} else {
								v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
								v255 = int32(1)
								v256 = v254 & v255
								if v254 == v255 {
									v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
									if base.Ui32((v260-int32(1))&int32(255)) < base.Ui32(int32(3)) {
										v285 = int32(4)
										if v256 != 0 {
											v288 = int32(1)
										} else {
											v288 = int32(4)
										}
										v290 = v252
										v292 = v14 + v288
										v294 = v285
										for {
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
											if v298 == int32(92) {
												v301 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
												v335 = v290 + int32(2)
											} else {
												if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v311 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
													v313 = int32(7)
													v315 = int32(48)
													v316 = v298&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
													v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
													v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
													v335 = v290 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
													v335 = v290 + int32(1)
												}
											}
											v336 = int32(1)
											v339 = v294 - v336
											if v339 != 0 {
												v290 = v335
												v292 = v292 + v336
												v294 = v339
												continue
											} else {
												break
											}
											break
										}
										v340 = v335
										v343 = v252
									} else {
										if v260 == int32(18) {
											v271 = int32(16)
										} else {
											v271 = int32(0)
										}
										v283 = v271
										if v283 != 0 {
											v285 = v283
											if v256 != 0 {
												v288 = int32(1)
											} else {
												v288 = int32(4)
											}
											v290 = v252
											v292 = v14 + v288
											v294 = v285
											for {
												v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
												if v298 == int32(92) {
													v301 = int32(_a_F_byteaout_3)
													*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
													v335 = v290 + int32(2)
												} else {
													if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
														v311 = int32(92)
														*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
														v313 = int32(7)
														v315 = int32(48)
														v316 = v298&v313 | v315
														*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
														v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
														*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
														v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
														*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
														v335 = v290 + int32(4)
													} else {
														*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
														v335 = v290 + int32(1)
													}
												}
												v336 = int32(1)
												v339 = v294 - v336
												if v339 != 0 {
													v290 = v335
													v292 = v292 + v336
													v294 = v339
													continue
												} else {
													break
												}
												break
											}
											v340 = v335
											v343 = v252
										} else {
											v340 = v252
											v343 = v252
										}
									}
								} else {
									v272 = int32(1)
									if v256 != 0 {
										v283 = int32(base.Ui32(v254)>>(uint(v272)%32)) - v272
									} else {
										v276 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
										v283 = int32(base.Ui32(v276)>>(uint(int32(2))%32)) - int32(4)
									}
									if v283 != 0 {
										v285 = v283
										if v256 != 0 {
											v288 = int32(1)
										} else {
											v288 = int32(4)
										}
										v290 = v252
										v292 = v14 + v288
										v294 = v285
										for {
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
											if v298 == int32(92) {
												v301 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
												v335 = v290 + int32(2)
											} else {
												if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v311 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
													v313 = int32(7)
													v315 = int32(48)
													v316 = v298&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
													v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
													v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
													v335 = v290 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
													v335 = v290 + int32(1)
												}
											}
											v336 = int32(1)
											v339 = v294 - v336
											if v339 != 0 {
												v290 = v335
												v292 = v292 + v336
												v294 = v339
												continue
											} else {
												break
											}
											break
										}
										v340 = v335
										v343 = v252
									} else {
										v340 = v252
										v343 = v252
									}
								}
								v348 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v340))) = uint8(v348)
								m.G0 = v11 + int32(16)
								return base.I64_extend_i32_u(v343)
							}
						}
					} else {
						v250 = int64(1)
						v252 = F_palloc(m, base.I32_wrap_i64(v250))
						mBase = m.M
						v253 = m.ExcPending
						if v253 != 0 {
							return int64(0)
						} else {
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
							v255 = int32(1)
							v256 = v254 & v255
							if v254 == v255 {
								v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
								if base.Ui32((v260-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v285 = int32(4)
									if v256 != 0 {
										v288 = int32(1)
									} else {
										v288 = int32(4)
									}
									v290 = v252
									v292 = v14 + v288
									v294 = v285
									for {
										v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
										if v298 == int32(92) {
											v301 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
											v335 = v290 + int32(2)
										} else {
											if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v311 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
												v313 = int32(7)
												v315 = int32(48)
												v316 = v298&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
												v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
												v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
												v335 = v290 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
												v335 = v290 + int32(1)
											}
										}
										v336 = int32(1)
										v339 = v294 - v336
										if v339 != 0 {
											v290 = v335
											v292 = v292 + v336
											v294 = v339
											continue
										} else {
											break
										}
										break
									}
									v340 = v335
									v343 = v252
								} else {
									if v260 == int32(18) {
										v271 = int32(16)
									} else {
										v271 = int32(0)
									}
									v283 = v271
									if v283 != 0 {
										v285 = v283
										if v256 != 0 {
											v288 = int32(1)
										} else {
											v288 = int32(4)
										}
										v290 = v252
										v292 = v14 + v288
										v294 = v285
										for {
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
											if v298 == int32(92) {
												v301 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
												v335 = v290 + int32(2)
											} else {
												if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v311 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
													v313 = int32(7)
													v315 = int32(48)
													v316 = v298&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
													v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
													v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
													v335 = v290 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
													v335 = v290 + int32(1)
												}
											}
											v336 = int32(1)
											v339 = v294 - v336
											if v339 != 0 {
												v290 = v335
												v292 = v292 + v336
												v294 = v339
												continue
											} else {
												break
											}
											break
										}
										v340 = v335
										v343 = v252
									} else {
										v340 = v252
										v343 = v252
									}
								}
							} else {
								v272 = int32(1)
								if v256 != 0 {
									v283 = int32(base.Ui32(v254)>>(uint(v272)%32)) - v272
								} else {
									v276 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
									v283 = int32(base.Ui32(v276)>>(uint(int32(2))%32)) - int32(4)
								}
								if v283 != 0 {
									v285 = v283
									if v256 != 0 {
										v288 = int32(1)
									} else {
										v288 = int32(4)
									}
									v290 = v252
									v292 = v14 + v288
									v294 = v285
									for {
										v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
										if v298 == int32(92) {
											v301 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
											v335 = v290 + int32(2)
										} else {
											if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v311 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
												v313 = int32(7)
												v315 = int32(48)
												v316 = v298&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
												v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
												v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
												v335 = v290 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
												v335 = v290 + int32(1)
											}
										}
										v336 = int32(1)
										v339 = v294 - v336
										if v339 != 0 {
											v290 = v335
											v292 = v292 + v336
											v294 = v339
											continue
										} else {
											break
										}
										break
									}
									v340 = v335
									v343 = v252
								} else {
									v340 = v252
									v343 = v252
								}
							}
							v348 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v340))) = uint8(v348)
							m.G0 = v11 + int32(16)
							return base.I64_extend_i32_u(v343)
						}
					}
				}
			} else {
				v137 = int32(1)
				if v121 != 0 {
					v148 = int32(base.Ui32(v119)>>(uint(v137)%32)) - v137
				} else {
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v148 = int32(base.Ui32(v141)>>(uint(int32(2))%32)) - int32(4)
				}
				if v148 != 0 {
					v151 = v148
					if v121 != 0 {
						v154 = int32(1)
					} else {
						v154 = int32(4)
					}
					v155 = v14 + v154
					if v151 == int32(1) {
						v210 = v155
						v217 = int64(1)
						v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
						if base.Ui32((v221-int32(127))&int32(255)) < base.Ui32(int32(161)) {
							v228 = int64(4)
						} else {
							v228 = int64(1)
						}
						if v221 == int32(92) {
							v231 = int64(2)
						} else {
							v231 = v228
						}
						v240 = v217 + v231
					} else {
						v165 = v155
						v167 = int32(0)
						v172 = int64(1)
						for {
							v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
							if base.Ui32((v176-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v183 = int64(4)
							} else {
								v183 = int64(1)
							}
							if v176 == int32(92) {
								v186 = int64(2)
							} else {
								v186 = v183
							}
							v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+1)))
							if base.Ui32((v191-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v198 = int64(4)
							} else {
								v198 = int64(1)
							}
							if v191 == int32(92) {
								v201 = int64(2)
							} else {
								v201 = v198
							}
							v202 = v172 + v186 + v201
							v203 = int32(2)
							v204 = v165 + v203
							v206 = v167 + v203
							if v206 != v151&int32(-2) {
								v165 = v204
								v167 = v206
								v172 = v202
								continue
							} else {
								break
							}
							break
						}
						if v151&int32(1) == int32(0) {
							v240 = v202
						} else {
							v210 = v204
							v217 = v202
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
							if base.Ui32((v221-int32(127))&int32(255)) < base.Ui32(int32(161)) {
								v228 = int64(4)
							} else {
								v228 = int64(1)
							}
							if v221 == int32(92) {
								v231 = int64(2)
							} else {
								v231 = v228
							}
							v240 = v217 + v231
						}
					}
					if base.Ui64(int64(1073741824)) <= base.Ui64(v240) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v358 = m.ExcPending
						if v358 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v361 = m.ExcPending
							if v361 != 0 {
								return int64(0)
							} else {
								F_errmsg_internal(m, int32(_a_F_byteaout_0), int32(0))
								mBase = m.M
								v365 = m.ExcPending
								if v365 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_byteaout_1), int32(316), int32(_a_F_byteaout_2))
									mBase = m.M
									v370 = m.ExcPending
									if v370 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v250 = v240
						v252 = F_palloc(m, base.I32_wrap_i64(v250))
						mBase = m.M
						v253 = m.ExcPending
						if v253 != 0 {
							return int64(0)
						} else {
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
							v255 = int32(1)
							v256 = v254 & v255
							if v254 == v255 {
								v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
								if base.Ui32((v260-int32(1))&int32(255)) < base.Ui32(int32(3)) {
									v285 = int32(4)
									if v256 != 0 {
										v288 = int32(1)
									} else {
										v288 = int32(4)
									}
									v290 = v252
									v292 = v14 + v288
									v294 = v285
									for {
										v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
										if v298 == int32(92) {
											v301 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
											v335 = v290 + int32(2)
										} else {
											if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v311 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
												v313 = int32(7)
												v315 = int32(48)
												v316 = v298&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
												v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
												v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
												v335 = v290 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
												v335 = v290 + int32(1)
											}
										}
										v336 = int32(1)
										v339 = v294 - v336
										if v339 != 0 {
											v290 = v335
											v292 = v292 + v336
											v294 = v339
											continue
										} else {
											break
										}
										break
									}
									v340 = v335
									v343 = v252
								} else {
									if v260 == int32(18) {
										v271 = int32(16)
									} else {
										v271 = int32(0)
									}
									v283 = v271
									if v283 != 0 {
										v285 = v283
										if v256 != 0 {
											v288 = int32(1)
										} else {
											v288 = int32(4)
										}
										v290 = v252
										v292 = v14 + v288
										v294 = v285
										for {
											v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
											if v298 == int32(92) {
												v301 = int32(_a_F_byteaout_3)
												*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
												v335 = v290 + int32(2)
											} else {
												if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
													v311 = int32(92)
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
													v313 = int32(7)
													v315 = int32(48)
													v316 = v298&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
													v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
													v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
													*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
													v335 = v290 + int32(4)
												} else {
													*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
													v335 = v290 + int32(1)
												}
											}
											v336 = int32(1)
											v339 = v294 - v336
											if v339 != 0 {
												v290 = v335
												v292 = v292 + v336
												v294 = v339
												continue
											} else {
												break
											}
											break
										}
										v340 = v335
										v343 = v252
									} else {
										v340 = v252
										v343 = v252
									}
								}
							} else {
								v272 = int32(1)
								if v256 != 0 {
									v283 = int32(base.Ui32(v254)>>(uint(v272)%32)) - v272
								} else {
									v276 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
									v283 = int32(base.Ui32(v276)>>(uint(int32(2))%32)) - int32(4)
								}
								if v283 != 0 {
									v285 = v283
									if v256 != 0 {
										v288 = int32(1)
									} else {
										v288 = int32(4)
									}
									v290 = v252
									v292 = v14 + v288
									v294 = v285
									for {
										v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
										if v298 == int32(92) {
											v301 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
											v335 = v290 + int32(2)
										} else {
											if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v311 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
												v313 = int32(7)
												v315 = int32(48)
												v316 = v298&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
												v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
												v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
												v335 = v290 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
												v335 = v290 + int32(1)
											}
										}
										v336 = int32(1)
										v339 = v294 - v336
										if v339 != 0 {
											v290 = v335
											v292 = v292 + v336
											v294 = v339
											continue
										} else {
											break
										}
										break
									}
									v340 = v335
									v343 = v252
								} else {
									v340 = v252
									v343 = v252
								}
							}
							v348 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v340))) = uint8(v348)
							m.G0 = v11 + int32(16)
							return base.I64_extend_i32_u(v343)
						}
					}
				} else {
					v250 = int64(1)
					v252 = F_palloc(m, base.I32_wrap_i64(v250))
					mBase = m.M
					v253 = m.ExcPending
					if v253 != 0 {
						return int64(0)
					} else {
						v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
						v255 = int32(1)
						v256 = v254 & v255
						if v254 == v255 {
							v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
							if base.Ui32((v260-int32(1))&int32(255)) < base.Ui32(int32(3)) {
								v285 = int32(4)
								if v256 != 0 {
									v288 = int32(1)
								} else {
									v288 = int32(4)
								}
								v290 = v252
								v292 = v14 + v288
								v294 = v285
								for {
									v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
									if v298 == int32(92) {
										v301 = int32(_a_F_byteaout_3)
										*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
										v335 = v290 + int32(2)
									} else {
										if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
											v311 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
											v313 = int32(7)
											v315 = int32(48)
											v316 = v298&v313 | v315
											*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
											v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
											*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
											v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
											*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
											v335 = v290 + int32(4)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
											v335 = v290 + int32(1)
										}
									}
									v336 = int32(1)
									v339 = v294 - v336
									if v339 != 0 {
										v290 = v335
										v292 = v292 + v336
										v294 = v339
										continue
									} else {
										break
									}
									break
								}
								v340 = v335
								v343 = v252
							} else {
								if v260 == int32(18) {
									v271 = int32(16)
								} else {
									v271 = int32(0)
								}
								v283 = v271
								if v283 != 0 {
									v285 = v283
									if v256 != 0 {
										v288 = int32(1)
									} else {
										v288 = int32(4)
									}
									v290 = v252
									v292 = v14 + v288
									v294 = v285
									for {
										v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
										if v298 == int32(92) {
											v301 = int32(_a_F_byteaout_3)
											*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
											v335 = v290 + int32(2)
										} else {
											if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
												v311 = int32(92)
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
												v313 = int32(7)
												v315 = int32(48)
												v316 = v298&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
												v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
												v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
												*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
												v335 = v290 + int32(4)
											} else {
												*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
												v335 = v290 + int32(1)
											}
										}
										v336 = int32(1)
										v339 = v294 - v336
										if v339 != 0 {
											v290 = v335
											v292 = v292 + v336
											v294 = v339
											continue
										} else {
											break
										}
										break
									}
									v340 = v335
									v343 = v252
								} else {
									v340 = v252
									v343 = v252
								}
							}
						} else {
							v272 = int32(1)
							if v256 != 0 {
								v283 = int32(base.Ui32(v254)>>(uint(v272)%32)) - v272
							} else {
								v276 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
								v283 = int32(base.Ui32(v276)>>(uint(int32(2))%32)) - int32(4)
							}
							if v283 != 0 {
								v285 = v283
								if v256 != 0 {
									v288 = int32(1)
								} else {
									v288 = int32(4)
								}
								v290 = v252
								v292 = v14 + v288
								v294 = v285
								for {
									v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v292))))
									if v298 == int32(92) {
										v301 = int32(_a_F_byteaout_3)
										*(*uint16)(unsafe.Add(mBase, uint32(v290))) = uint16(v301)
										v335 = v290 + int32(2)
									} else {
										if base.Ui32((v298-int32(127))&int32(255)) <= base.Ui32(int32(160)) {
											v311 = int32(92)
											*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v311)
											v313 = int32(7)
											v315 = int32(48)
											v316 = v298&v313 | v315
											*(*uint8)(unsafe.Add(mBase, uint32(v290)+3)) = uint8(v316)
											v321 = int32(base.Ui32(v298)>>(uint(int32(6))%32)) | v315
											*(*uint8)(unsafe.Add(mBase, uint32(v290)+1)) = uint8(v321)
											v328 = int32(base.Ui32(v298)>>(uint(int32(3))%32))&v313 | v315
											*(*uint8)(unsafe.Add(mBase, uint32(v290)+2)) = uint8(v328)
											v335 = v290 + int32(4)
										} else {
											*(*uint8)(unsafe.Add(mBase, uint32(v290))) = uint8(v298)
											v335 = v290 + int32(1)
										}
									}
									v336 = int32(1)
									v339 = v294 - v336
									if v339 != 0 {
										v290 = v335
										v292 = v292 + v336
										v294 = v339
										continue
									} else {
										break
									}
									break
								}
								v340 = v335
								v343 = v252
							} else {
								v340 = v252
								v343 = v252
							}
						}
						v348 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v340))) = uint8(v348)
						m.G0 = v11 + int32(16)
						return base.I64_extend_i32_u(v343)
					}
				}
			}
		case 1:
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
			if v20 == int32(1) {
				v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
				if v26 == int32(18) {
					v29 = int32(16)
				} else {
					v29 = int32(0)
				}
				if base.Ui32((v26-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v36 = int32(4)
				} else {
					v36 = v29
				}
				v49 = v36
			} else {
				v37 = int32(1)
				if v20&v37 != 0 {
					v49 = int32(base.Ui32(v20)>>(uint(v37)%32)) - v37
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					v49 = int32(base.Ui32(v43)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			v54 = F_palloc(m, v49<<(uint(int32(1))%32)+int32(3))
			mBase = m.M
			v55 = m.ExcPending
			if v55 != 0 {
				return int64(0)
			} else {
				v56 = int32(_a_F_byteaout_4)
				*(*uint16)(unsafe.Add(mBase, uint32(v54))) = uint16(v56)
				v59 = v54 + int32(2)
				v60 = int32(1)
				v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
				v64 = v62 & v60
				if v64 != 0 {
					v65 = v60
				} else {
					v65 = int32(4)
				}
				v66 = v14 + v65
				if v62 == int32(1) {
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
					if v72 == int32(18) {
						v75 = int32(16)
					} else {
						v75 = int32(0)
					}
					if base.Ui32((v72-int32(1))&int32(255)) < base.Ui32(int32(3)) {
						v82 = int32(4)
					} else {
						v82 = v75
					}
					v93 = v82
				} else {
					v83 = int32(1)
					if v64 != 0 {
						v93 = int32(base.Ui32(v62)>>(uint(v83)%32)) - v83
					} else {
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
						v93 = int32(base.Ui32(v87)>>(uint(int32(2))%32)) - int32(4)
					}
				}
				if v93 != 0 {
					v96 = v66
					v98 = v59
					for {
						v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
						v101 = int32(1)
						v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100<<(uint(v101)%32))+uint32(_c_F_byteaout[1]))))
						*(*uint16)(unsafe.Add(mBase, uint32(v98))) = uint16(v103)
						v108 = v96 + v101
						if base.Ui32(v108) < base.Ui32(v66+v93) {
							v96 = v108
							v98 = v98 + int32(2)
							continue
						} else {
							break
						}
						break
					}
				} else {
				}
				v340 = v59 + base.I32_wrap_i64(base.I64_extend_i32_u(v93)<<(uint(int64(1))%64))
				v343 = v54
				v348 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v340))) = uint8(v348)
				m.G0 = v11 + int32(16)
				return base.I64_extend_i32_u(v343)
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v374 = m.ExcPending
			if v374 != 0 {
				return int64(0)
			} else {
				v376 = *(*int32)(unsafe.Add(mBase, _c_F_byteaout[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v11))) = v376
				F_errmsg_internal(m, int32(_a_F_byteaout_5), v11)
				mBase = m.M
				v380 = m.ExcPending
				if v380 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_byteaout_1), int32(347), int32(_a_F_byteaout_2))
					mBase = m.M
					v385 = m.ExcPending
					if v385 != 0 {
						return int64(0)
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
func F_byteapos(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
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
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
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
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum_packed(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = F_pg_detoast_datum_packed(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9))))
	if v16 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v46 == int32(1) {
		goto L17
	} else {
		goto L18
	}
L5:
	;
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+1)))
	if v22 == int32(18) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v33 = int32(1)
	if v16&v33 != 0 {
		v45 = int32(base.Ui32(v16)>>(uint(v33)%32)) - v33
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v25 = int32(16)
	goto L10
L9:
	;
	v25 = int32(0)
	goto L10
L10:
	;
	if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v32 = int32(4)
	goto L13
L12:
	;
	v32 = v25
	goto L13
L13:
	;
	v45 = v32
	goto L4
L14:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	if v80 <= v45 {
		goto L26
	} else {
		goto L27
	}
L16:
	;
	if int32(0) < v75 {
		v80 = v75
		goto L15
	} else {
		goto L25
	}
L17:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if base.Ui32((v50-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		v80 = int32(4)
		goto L15
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v62 = int32(1)
	if v46&v62 != 0 {
		v75 = int32(base.Ui32(v46)>>(uint(v62)%32)) - v62
		goto L16
	} else {
		goto L24
	}
L20:
	;
	if v50 == int32(18) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v61 = int32(16)
	goto L23
L22:
	;
	v61 = int32(0)
	goto L23
L23:
	;
	v75 = v61
	goto L16
L24:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v75 = int32(base.Ui32(v68)>>(uint(int32(2))%32)) - int32(4)
	goto L16
L25:
	;
	return int64(1)
L26:
	;
	v84 = int32(1)
	if v16&v84 != 0 {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	goto L28
L28:
	;
	return int64(0)
L29:
	;
	v90 = v84
	goto L31
L30:
	;
	v90 = int32(4)
	goto L31
L31:
	;
	v93 = int32(1)
	if v46&v93 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v97 = v93
	goto L34
L33:
	;
	v97 = int32(4)
	goto L34
L34:
	;
	v98 = v14 + v97
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98))))
	v100 = int32(0)
	v101 = v9 + v90
	goto L35
L35:
	;
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v99 != v107 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	goto L28
L37:
	;
	v175 = int32(1)
	v178 = v100 + v175
	if v178 != v45-v80+v84 {
		v100 = v178
		v101 = v101 + v175
		goto L35
	} else {
		goto L58
	}
L38:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v80) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if v170 != 0 {
		goto L37
	} else {
		goto L57
	}
L40:
	;
	v170 = int32(0)
	goto L39
L41:
	;
	v144 = v139
	v145 = v140
	v146 = v141
	goto L51
L42:
	;
	if (v101|v98)&int32(3) != 0 {
		v139 = v101
		v140 = v98
		v141 = v80
		goto L41
	} else {
		goto L45
	}
L43:
	;
	v132 = v101
	v133 = v98
	v134 = v80
	goto L44
L44:
	;
	if v134 == int32(0) {
		goto L40
	} else {
		goto L50
	}
L45:
	;
	v116 = v101
	v117 = v98
	v118 = v80
	goto L46
L46:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v117)))
	if v121 != v122 {
		v139 = v116
		v140 = v117
		v141 = v118
		goto L41
	} else {
		goto L48
	}
L47:
	;
	v132 = v127
	v133 = v125
	v134 = v129
	goto L44
L48:
	;
	v124 = int32(4)
	v125 = v117 + v124
	v127 = v116 + v124
	v129 = v118 - v124
	if base.Ui32(int32(3)) < base.Ui32(v129) {
		v116 = v127
		v117 = v125
		v118 = v129
		goto L46
	} else {
		goto L49
	}
L49:
	;
	goto L47
L50:
	;
	v139 = v132
	v140 = v133
	v141 = v134
	goto L41
L51:
	;
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v149 == v150 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v170 = v149 - v150
	goto L39
L53:
	;
	v152 = int32(1)
	v157 = v146 - v152
	if v157 != 0 {
		v144 = v144 + v152
		v145 = v145 + v152
		v146 = v157
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	goto L40
L57:
	;
	return base.I64_extend_i32_u(v100 + int32(1))
L58:
	;
	goto L36
}
func F_byteartrim(m *base.Module, l0 int32) int64 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int64(0)
		} else {
			v12 = F_dobyteatrim(m, v3, v8, int32(0), int32(1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v12)
			}
		}
	}
}
