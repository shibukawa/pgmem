package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HeapTupleGetUpdateXid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	v2 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = F_GetMultiXactIdMembers(m, v10, v8+int32(12), v2)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if int32(0) < v14 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v22 = v2
	goto L8
L4:
	;
	v42 = v2
	goto L5
L5:
	;
	m.G0 = v8 + int32(16)
	return v42
L6:
	;
	F_pfree(m, v20)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L12
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v38 = v36
	goto L6
L8:
	;
	v28 = v20 + v22<<(uint(int32(3))%32)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v29) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	v38 = int32(0)
	goto L6
L10:
	;
	v33 = v22 + int32(1)
	if v33 != v14 {
		v22 = v33
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v42 = v38
	goto L5
}
func F_HeapTupleHeaderGetCmax(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v4&int32(32) != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderGetCmax[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+v3<<(uint(int32(3))%32))+4))
		v13 = v12
	} else {
		v13 = v3
	}
	return v13
}
func F_HeapTupleHeaderGetCmin(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v4&int32(32) != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_HeapTupleHeaderGetCmin[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+v3<<(uint(int32(3))%32))))
		v13 = v12
	} else {
		v13 = v3
	}
	return v13
}
func F_heap_getattr_7(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+18)))
	if base.Ui32(v13&int32(2047)) < base.Ui32(l1) {
		v17 = F_getmissingattr(m, l2, l1, l3)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int32(0)
		} else {
			v82 = v17
			m.G0 = v10 + int32(16)
			return v82
		}
	} else {
		v21 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v21)
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+20)))
		if v24&int32(1) == v21 {
			v29 = int32(4)
			v33 = l2 + l1<<(uint(v29)%32) + v29
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
			if v34 < int32(0) {
				v75 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					v82 = v75
					m.G0 = v10 + int32(16)
					return v82
				}
			} else {
				v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+22)))
				v39 = v23 + v37 + v34
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+6)))
				if v40 != int32(1) {
					v82 = v39
					m.G0 = v10 + int32(16)
					return v82
				} else {
					v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+4)))
					switch v43&int32(_a_F_heap_getattr_7_0) - int32(1) {
					case 0:
						v48 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39))))
						v82 = v48
						m.G0 = v10 + int32(16)
						return v82
					case 1:
						v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39))))
						v82 = v49
						m.G0 = v10 + int32(16)
						return v82
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v43
							F_errmsg_internal(m, int32(_a_F_heap_getattr_7_1), v10)
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_heap_getattr_7_2), int32(70), int32(_a_F_heap_getattr_7_3))
								mBase = m.M
								v63 = m.ExcPending
								if v63 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 3:
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
						v82 = v50
						m.G0 = v10 + int32(16)
						return v82
					}
				}
			}
		} else {
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+23)))
			v65 = int32(1)
			if int32(base.Ui32(v64)>>(uint(l1-v65)%32))&v65 != 0 {
				v75 = F_nocachegetattr(m, l0, l1, l2)
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					v82 = v75
					m.G0 = v10 + int32(16)
					return v82
				}
			} else {
				v70 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v70)
				v82 = int32(0)
				m.G0 = v10 + int32(16)
				return v82
			}
		}
	}
}
func F_heap_getnextslot(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
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
	var v39 int64
	_ = v39
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+29)))
	if v9&int32(1) != 0 {
		F_heapgettup_pagemode(m, l0, l1, v8, v7)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			if v18 == int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
				m.T0[v22].(func(*base.Module, int32))(m, l2)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v18 != int32(0))
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+272))
				if v28 == int32(0) {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+268)))
					if v31 != int32(1) {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						F_ExecStoreBufferHeapTuple(m, l0-int32(-64), l2, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v18 != int32(0))
						}
					} else {
						F_pgstat_assoc_relation(m, v27)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+272))
							v38 = v37
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39 + int64(1)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							F_ExecStoreBufferHeapTuple(m, l0-int32(-64), l2, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v18 != int32(0))
							}
						}
					}
				} else {
					v38 = v28
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39 + int64(1)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					F_ExecStoreBufferHeapTuple(m, l0-int32(-64), l2, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v18 != int32(0))
					}
				}
			}
		}
	} else {
		F_heapgettup(m, l0, l1, v8, v7)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
			if v18 == int32(0) {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
				m.T0[v22].(func(*base.Module, int32))(m, l2)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return base.B2i32(v18 != int32(0))
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+272))
				if v28 == int32(0) {
					v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27)+268)))
					if v31 != int32(1) {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
						F_ExecStoreBufferHeapTuple(m, l0-int32(-64), l2, v44)
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v18 != int32(0))
						}
					} else {
						F_pgstat_assoc_relation(m, v27)
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+272))
							v38 = v37
							v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
							*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39 + int64(1)
							v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
							F_ExecStoreBufferHeapTuple(m, l0-int32(-64), l2, v44)
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int32(0)
							} else {
								return base.B2i32(v18 != int32(0))
							}
						}
					}
				} else {
					v38 = v28
					v39 = *(*int64)(unsafe.Add(mBase, uint32(v38)+24))
					*(*int64)(unsafe.Add(mBase, uint32(v38)+24)) = v39 + int64(1)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
					F_ExecStoreBufferHeapTuple(m, l0-int32(-64), l2, v44)
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v18 != int32(0))
					}
				}
			}
		}
	}
}
func F_heap_multi_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
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
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v121 int32
	_ = v121
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
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
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v365 int32
	_ = v365
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v491 int32
	_ = v491
	var v502 int32
	_ = v502
	var v507 int32
	_ = v507
	var v516 int32
	_ = v516
	var v525 int32
	_ = v525
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v570 int32
	_ = v570
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v653 int32
	_ = v653
	var v661 int32
	_ = v661
	var v685 int32
	_ = v685
	var v691 int32
	_ = v691
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v723 int32
	_ = v723
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v790 int64
	_ = v790
	var v791 int32
	_ = v791
	var v827 int32
	_ = v827
	var v829 int32
	_ = v829
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v921 int32
	_ = v921
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v955 int32
	_ = v955
	var v991 int32
	_ = v991
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1064 int32
	_ = v1064
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1134 int32
	_ = v1134
	v7 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(_a_F_heap_multi_insert_0)
	m.G0 = v35
	v37 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = int32(0)
	v42 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[0]))
	if v42 < int32(2) {
		v81 = v7
		v82 = v7
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+118)))
	if v84 != int32(112) {
		v97 = v7
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+118)))
	if v46 != int32(112) {
		v61 = v7
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v63 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+118)))
	if v63 != int32(112) {
		v81 = v7
		v82 = v61
		goto L3
	} else {
		goto L10
	}
L6:
	;
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+119)))
	if v49 == int32(102) {
		v61 = v7
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L8
L8:
	;
	v56 = base.B2i32(base.Ui32(v52) < base.Ui32(int32(_a_F_heap_multi_insert_1))) ^ int32(1)
	v58 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[0]))
	if v58 < int32(2) {
		v81 = v7
		v82 = v56
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v61 = v56
	goto L5
L10:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L11
L11:
	;
	if base.Ui32(v67) < base.Ui32(int32(_a_F_heap_multi_insert_1)) {
		v81 = int32(1)
		v82 = v61
		goto L3
	} else {
		goto L12
	}
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v70 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v81 = int32(0)
	v82 = v61
	goto L3
L14:
	;
	goto L15
L15:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+119)))
	switch v76 - int32(109) {
	case 0, 5:
		goto L16
	default:
		v81 = int32(0)
		v82 = v61
		goto L3
	}
L16:
	;
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70)+104)))
	v81 = v79
	v82 = v61
	goto L3
L17:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v98 != 0 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	v89 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[0]))
	if int32(0) < v89 {
		v97 = int32(1)
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v93 != 0 {
		v97 = int32(0)
		goto L17
	} else {
		goto L20
	}
L20:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v97 = base.B2i32(v94 == int32(0))
	goto L17
L21:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v105 = base.I32_div_s(int32(_a_F_heap_multi_insert_2)-v100<<(uint(int32(13))%32), int32(100))
	v107 = v105
	goto L23
L22:
	;
	v107 = int32(0)
	goto L23
L23:
	;
	v110 = F_palloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if int32(0) < l2 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v897 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	if v897 != 0 {
		goto L165
	} else {
		goto L166
	}
L26:
	;
	v121 = int32(0)
	goto L29
L27:
	;
	goto L28
L28:
	;
	F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L164
	}
L29:
	;
	v148 = v121 << (uint(int32(2)) % 32)
	v149 = l1 + v148
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v153 = F_ExecFetchSlotHeapTuple(m, v150, int32(1), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L34
	}
L31:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+36)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v158)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v153)+12)) = v159
	v162 = F_heap_prepare_insert(m, l0, v153, v37, l3, l4)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148+v110))) = v162
	v166 = v121 + int32(1)
	if v166 != l2 {
		v121 = v166
		goto L29
	} else {
		goto L33
	}
L33:
	;
	goto L30
L34:
	;
	if v82 != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v174 = int32(24)
	goto L37
L36:
	;
	v174 = int32(8)
	goto L37
L37:
	;
	v175 = v81 & v97
	v177 = int32(_a_F_heap_multi_insert_3) - v107
	v178 = int32(4)
	v179 = l4 & v178
	v185 = v35 + int32(16) | v178
	v195 = v7
	v199 = int32(0)
	v203 = v7
	v212 = v7
	goto L38
L38:
	;
	v220 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[1]))
	if v220 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	goto L25
L40:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v223 = int32(0)
	if v203&base.B2i32(v195 != v223) == v223 {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L42
L44:
	;
	v388 = v110 + v195<<(uint(int32(2))%32)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v389)))
	v391 = int32(0)
	v396 = F_RelationGetBufferForTuple(m, l0, v390, v391, l4, l5, v35+int32(12), v391, v365-v385)
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L1
	} else {
		goto L63
	}
L45:
	;
	v228 = int32(1)
	if v195+v228 == l2 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	goto L47
L47:
	;
	v365 = v199
	v385 = v212 + int32(1)
	goto L44
L48:
	;
	v340 = *(*int32)(unsafe.Add(mBase, uint32(v110+v308<<(uint(int32(2))%32))))
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v340)))
	v365 = v317 + base.B2i32(base.Ui32(v311) < base.Ui32((v341+int32(7))&int32(-8)|int32(4)))
	v385 = int32(0)
	goto L44
L49:
	;
	v308 = v195
	v311 = v177
	v317 = v228
	goto L48
L50:
	;
	goto L51
L51:
	;
	v232 = l2 - v195
	v241 = v195
	v244 = v177
	v245 = int32(0)
	v250 = v228
	goto L52
L52:
	;
	v272 = v110 + v241<<(uint(int32(2))%32)
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v272)))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v280 = (v274+int32(7))&int32(-8) | int32(4)
	v281 = base.B2i32(base.Ui32(v244) < base.Ui32(v280))
	if base.Ui32(v244) < base.Ui32(v280) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	v302 = int32(0)
	if v232&int32(1) == v302 {
		v365 = v296
		v385 = v302
		goto L44
	} else {
		goto L61
	}
L54:
	;
	v282 = v177
	goto L56
L55:
	;
	v282 = v244
	goto L56
L56:
	;
	v283 = v282 - v280
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v272)+4))
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v284)))
	v291 = (v285+int32(7))&int32(-8) | int32(4)
	v292 = base.B2i32(base.Ui32(v283) < base.Ui32(v291))
	if base.Ui32(v283) < base.Ui32(v291) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v293 = v177
	goto L59
L58:
	;
	v293 = v283
	goto L59
L59:
	;
	v294 = v293 - v291
	v296 = v250 + v281 + v292
	v297 = int32(2)
	v298 = v241 + v297
	v300 = v245 + v297
	if v300 != v232&int32(-2) {
		v241 = v298
		v244 = v294
		v245 = v300
		v250 = v296
		goto L52
	} else {
		goto L60
	}
L60:
	;
	goto L53
L61:
	;
	v308 = v298
	v311 = v294
	v317 = v296
	goto L48
L62:
	;
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+12)))
	v417 = int32(_a_F_heap_multi_insert_4)
	v419 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[2])) = v419 + int32(1)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	F_RelationPutHeapTuple(m, v396, v423, int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L1
	} else {
		goto L67
	}
L63:
	;
	if v396 < int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v401 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[3]))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v401+(v396^int32(-1))<<(uint(int32(2))%32))))
	v415 = v407
	goto L62
L65:
	;
	goto L66
L66:
	;
	v409 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[4]))
	v415 = v409 + v396<<(uint(int32(13))%32) + int32(-8192)
	goto L62
L67:
	;
	if v175 != 0 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	F_log_heap_new_cid(m, l0, v429)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L1
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	v438 = base.B2i32((v416+int32(_a_F_heap_multi_insert_5))&int32(_a_F_heap_multi_insert_6) == int32(0)) | base.B2i32(base.Ui32(v416) < base.Ui32(int32(25)))
	v439 = int32(1)
	v441 = v195 + v439
	if l2 <= v441 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	goto L70
L72:
	;
	v596 = v438 & int32(base.Ui32(v179)>>(uint(int32(2))%32))
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+10)))
	v604 = base.B2i32(v179 == int32(0)) & int32(base.Ui32(v599&int32(4))>>(uint(int32(2))%32))
	if v604 != 0 {
		goto L105
	} else {
		goto L106
	}
L73:
	;
	v570 = v439
	v595 = v441
	goto L72
L74:
	;
	goto L75
L75:
	;
	v443 = l2 - v195
	v447 = v441
	v451 = v439
	goto L76
L76:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v110+v447<<(uint(int32(2))%32))))
	v483 = int32(4)
	v484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+14)))
	v485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415)+12)))
	v486 = v484 - v485
	if v486 <= v483 {
		goto L79
	} else {
		goto L80
	}
L77:
	;
	v570 = v443
	v595 = l2
	goto L72
L78:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v479)))
	if base.Ui32(v546) < base.Ui32((v547+int32(7))&int32(-8)+v107) {
		v570 = v451
		v595 = v447
		goto L72
	} else {
		goto L97
	}
L79:
	;
	v489 = v483
	goto L81
L80:
	;
	v489 = v486
	goto L81
L81:
	;
	v491 = v489 - int32(4)
	if v491 == int32(0) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v546 = int32(0)
	goto L78
L83:
	;
	goto L84
L84:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v485) {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v546 = v491
	goto L78
L86:
	;
	v502 = int32(base.Ui32(v485+int32(_a_F_heap_multi_insert_5)) >> (uint(int32(2)) % 32))
	goto L88
L87:
	;
	v502 = int32(0)
	goto L88
L88:
	;
	if base.Ui32(v502&int32(_a_F_heap_multi_insert_7)) < base.Ui32(int32(291)) {
		goto L85
	} else {
		goto L89
	}
L89:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415)+10)))
	if v507&int32(1) == int32(0) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v546 = int32(0)
	goto L78
L91:
	;
	goto L92
L92:
	;
	v516 = int32(1)
	goto L93
L93:
	;
	v525 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v415+int32(20)+v516&int32(_a_F_heap_multi_insert_7)<<(uint(int32(2))%32))+1)))
	if v525&int32(384) == int32(0) {
		goto L85
	} else {
		goto L95
	}
L94:
	;
	v546 = int32(0)
	goto L78
L95:
	;
	v531 = v516 + int32(1)
	v532 = int32(_a_F_heap_multi_insert_7)
	if base.Ui32(v531&v532) <= base.Ui32(v502&v532) {
		v516 = v531
		goto L93
	} else {
		goto L96
	}
L96:
	;
	goto L94
L97:
	;
	F_RelationPutHeapTuple(m, v396, v479, int32(0))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	if v175 != 0 {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	F_log_heap_new_cid(m, l0, v479)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L1
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	v560 = v451 + int32(1)
	if v443 != v560 {
		v447 = v560 + v195
		v451 = v560
		goto L76
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	goto L77
L104:
	;
	F_MarkBufferDirty(m, v396)
	mBase = m.M
	v637 = m.ExcPending
	if v637 != 0 {
		goto L1
	} else {
		goto L114
	}
L105:
	;
	v606 = v599 & int32(_a_F_heap_multi_insert_8)
	*(*uint16)(unsafe.Add(mBase, uint32(v415)+10)) = uint16(v606)
	if v396 < int32(0) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L107
L107:
	;
	if v596 == int32(0) {
		goto L104
	} else {
		goto L113
	}
L108:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v629 = F_visibilitymap_clear(m, v626, v627, int32(3))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L1
	} else {
		goto L112
	}
L109:
	;
	v611 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[5]))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v611+(v396^int32(-1))<<(uint(int32(6))%32))+16))
	v626 = v617
	goto L108
L110:
	;
	goto L111
L111:
	;
	v619 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[6]))
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v619+v396<<(uint(int32(6))%32)+int32(-64))+16))
	v626 = v625
	goto L108
L112:
	;
	goto L104
L113:
	;
	v634 = v599 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v415)+10)) = uint16(v634)
	goto L104
L114:
	;
	if v97 != 0 {
		goto L115
	} else {
		goto L116
	}
L115:
	;
	if v596 != 0 {
		goto L118
	} else {
		goto L119
	}
L116:
	;
	goto L117
L117:
	;
	v827 = int32(_a_F_heap_multi_insert_4)
	v829 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[2])) = v829 - int32(1)
	if v596 != 0 {
		goto L154
	} else {
		goto L155
	}
L118:
	;
	v639 = int32(32)
	goto L120
L119:
	;
	v639 = v604
	goto L120
L120:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)) = uint8(v639)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+18)) = uint16(v570)
	v642 = int32(0)
	if v438 != 0 {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v646 = v642
	goto L123
L122:
	;
	v646 = v570 << (uint(int32(1)) % 32)
	goto L123
L123:
	;
	v647 = v185 + v646
	if int32(0) < v570 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v653 = v647
	v661 = v642
	goto L127
L125:
	;
	v723 = v647
	goto L126
L126:
	;
	if v82 != 0 {
		goto L136
	} else {
		goto L137
	}
L127:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v388+v661<<(uint(int32(2))%32))))
	if v438 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	v723 = v716
	goto L126
L129:
	;
	v691 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v685)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v185+v661<<(uint(int32(1))%32)))) = uint16(v691)
	goto L131
L130:
	;
	goto L131
L131:
	;
	v696 = (v653 + int32(1)) & int32(-2)
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v685)+16))
	v698 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v697)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v696)+2)) = uint16(v698)
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v685)+16))
	v701 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v700)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v696)+4)) = uint16(v701)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v685)+16))
	v704 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v703)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v696)+6)) = uint8(v704)
	v707 = v696 + int32(7)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v685)))
	v710 = v708 - int32(23)
	if v710 != 0 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v711 = *(*int32)(unsafe.Add(mBase, uint32(v685)+16))
	base.MemoryCopy(m, v707, v711+int32(23), v710)
	goto L134
L133:
	;
	goto L134
L134:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v696))) = uint16(v710)
	v716 = v710 + v707
	v718 = v661 + int32(1)
	if v718 != v570 {
		v653 = v716
		v661 = v718
		goto L127
	} else {
		goto L135
	}
L135:
	;
	goto L128
L136:
	;
	v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)))
	v754 = v752 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)) = uint8(v754)
	goto L138
L137:
	;
	goto L138
L138:
	;
	if l2 == v595 {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)))
	v759 = v757 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)) = uint8(v759)
	goto L141
L140:
	;
	goto L141
L141:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	F_XLogRegisterData(m, v35+int32(16), v646+int32(4))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v770 = int32(0)
	if v438 != 0 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v773 = int32(6)
	goto L146
L145:
	;
	v773 = v770
	goto L146
L146:
	;
	F_XLogRegisterBuffer(m, v770, v396, v773|v174)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_XLogRegisterBufData(m, int32(0), v647, v723-v647)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v781 = int32(_a_F_heap_multi_insert_9)
	v783 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_multi_insert[7])))
	v784 = v783 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_multi_insert[7])) = uint8(v784)
	goto L149
L149:
	;
	if v438 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v789 = int32(208)
	goto L152
L151:
	;
	v789 = int32(80)
	goto L152
L152:
	;
	v790 = F_XLogInsert(m, int32(9), v789)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v415))) = base.I64_rotr(v790, int64(32))
	goto L117
L154:
	;
	if v396 < int32(0) {
		goto L158
	} else {
		goto L159
	}
L155:
	;
	goto L156
L156:
	;
	F_UnlockReleaseBuffer(m, v396)
	mBase = m.M
	v859 = m.ExcPending
	if v859 != 0 {
		goto L1
	} else {
		goto L162
	}
L157:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v856 = F_visibilitymap_set(m, l0, v851, v396, int64(0), v853, int32(0), int32(3))
	mBase = m.M
	v857 = m.ExcPending
	if v857 != 0 {
		goto L1
	} else {
		goto L161
	}
L158:
	;
	v836 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[5]))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v836+(v396^int32(-1))<<(uint(int32(6))%32))+16))
	v851 = v842
	goto L157
L159:
	;
	goto L160
L160:
	;
	v844 = *(*int32)(unsafe.Add(mBase, _c_F_heap_multi_insert[6]))
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v844+v396<<(uint(int32(6))%32)+int32(-64))+16))
	v851 = v850
	goto L157
L161:
	;
	goto L156
L162:
	;
	if v595 < l2 {
		v195 = v595
		v199 = v365
		v203 = v438
		v212 = v385
		goto L38
	} else {
		goto L163
	}
L163:
	;
	goto L39
L164:
	;
	goto L25
L165:
	;
	F_ReleaseBuffer(m, v897)
	mBase = m.M
	v899 = m.ExcPending
	if v899 != 0 {
		goto L1
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v900 = int32(0)
	F_CheckForSerializableConflictIn(m, l0, v900, int32(-1))
	mBase = m.M
	v904 = m.ExcPending
	if v904 != 0 {
		goto L1
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	v905 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L172
L170:
	;
	F_pgstat_count_heap_insert(m, l0, base.I64_extend_i32_s(l2))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L1
	} else {
		goto L188
	}
L171:
	;
	v991 = int32(0)
	if l2 != int32(1) {
		goto L181
	} else {
		goto L182
	}
L172:
	;
	v908 = int32(0)
	if base.B2i32(base.B2i32(base.Ui32(v905) < base.Ui32(int32(_a_F_heap_multi_insert_1))) == v908)|base.B2i32(l2 <= v908) == v908 {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	v921 = v900
	goto L176
L174:
	;
	goto L175
L175:
	;
	if l2 <= int32(0) {
		goto L170
	} else {
		goto L180
	}
L176:
	;
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v110+v921<<(uint(int32(2))%32))))
	F_CacheInvalidateHeapTuple(m, l0, v950, int32(0))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L178
	}
L177:
	;
	goto L171
L178:
	;
	v955 = v921 + int32(1)
	if v955 != l2 {
		v921 = v955
		goto L176
	} else {
		goto L179
	}
L179:
	;
	goto L177
L180:
	;
	goto L171
L181:
	;
	v1005 = v991
	v1006 = int32(0)
	goto L184
L182:
	;
	v1064 = v991
	goto L183
L183:
	;
	v1091 = v1064 << (uint(int32(2)) % 32)
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(l1+v1091)))
	v1095 = *(*int32)(unsafe.Add(mBase, uint32(v1091+v110)))
	v1096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1095)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1093)+32)) = uint16(v1096)
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1093)+28)) = v1098
	goto L170
L184:
	;
	v1031 = int32(2)
	v1032 = v1005 << (uint(v1031) % 32)
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l1+v1032)))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v1032+v110)))
	v1037 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1036)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1034)+32)) = uint16(v1037)
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v1036)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1034)+28)) = v1039
	v1042 = v1032 | int32(4)
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(l1+v1042)))
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(v1042+v110)))
	v1047 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1046)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1044)+32)) = uint16(v1047)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(v1046)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1044)+28)) = v1049
	v1052 = v1005 + v1031
	v1054 = v1006 + v1031
	if v1054 != l2&int32(-2) {
		v1005 = v1052
		v1006 = v1054
		goto L184
	} else {
		goto L186
	}
L185:
	;
	if l2&int32(1) == int32(0) {
		goto L170
	} else {
		goto L187
	}
L186:
	;
	goto L185
L187:
	;
	v1064 = v1052
	goto L183
L188:
	;
	m.G0 = v35 + int32(_a_F_heap_multi_insert_0)
	return
}
func F_heap_prepare_insert(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_heap_prepare_insert[0]))
	if v8 < int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)))
		v14 = v12 & int32(15)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v14)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+18)))
		v19 = v17 & int32(_a_F_heap_prepare_insert_0)
		*(*uint16)(unsafe.Add(mBase, uint32(v16)+18)) = uint16(v19)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)))
		v24 = v22 | int32(2048)
		*(*uint16)(unsafe.Add(mBase, uint32(v21)+20)) = uint16(v24)
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v26))) = l2
		if l4&int32(4) != 0 {
			v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30)+20)))
			v33 = v31 | int32(768)
			*(*uint16)(unsafe.Add(mBase, uint32(v30)+20)) = uint16(v33)
		} else {
		}
		v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = l3
		v38 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+20)))
		v40 = v38 & int32(_a_F_heap_prepare_insert_1)
		*(*uint16)(unsafe.Add(mBase, uint32(v36)+20)) = uint16(v40)
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v42)+4)) = int32(0)
		v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v45
		v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
		v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+119)))
		switch v48 - int32(109) {
		case 0, 5:
			v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
			v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+20)))
			if v52&int32(4) == int32(0) {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
				if base.Ui32(v57) < base.Ui32(int32(2033)) {
					v65 = l1
					return v65
				} else {
					v61 = F_heap_toast_insert_or_update(m, l0, l1, int32(0), l4)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = v61
						return v65
					}
				}
			} else {
				v61 = F_heap_toast_insert_or_update(m, l0, l1, int32(0), l4)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v65 = v61
					return v65
				}
			}
		default:
			v65 = l1
			return v65
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v70 = m.ExcPending
		if v70 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(322))
			mBase = m.M
			v73 = m.ExcPending
			if v73 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(_a_F_heap_prepare_insert_2), int32(0))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_heap_prepare_insert_3), int32(2283), int32(_a_F_heap_prepare_insert_4))
					mBase = m.M
					v82 = m.ExcPending
					if v82 != 0 {
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
func F_heap_prune_record_unchanged_lp_normal(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v12 int32
	_ = v12
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v77 int32
	_ = v77
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
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
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	v3 = l2
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = l1 + v3
	v12 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[0]))) = uint8(v12)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[1]))) = uint8(v12)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0+v3<<(uint(int32(2))%32))+20))
	v22 = l0 + v19&int32(_a_F_heap_prune_record_unchanged_lp_normal_0)
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[2]))))
	switch v23 - v12 {
	case 0:
		v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[3])))
		v27 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[3]))) = v26 + v27
		v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))))
		if v30 != v27 {
		} else {
			v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)))
			if v33&int32(256) == int32(0) {
				v38 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))) = uint8(v38)
			} else {
				v41 = int32(768)
				if v33&v41 != v41 {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
					v46 = v45
				} else {
					v46 = int32(2)
				}
				v47 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+8))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v48))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v46)) == int32(0) {
					v60 = base.B2i32(base.Ui32(v46) < base.Ui32(v48))
				} else {
					v60 = int32(base.Ui32(v46-v48) >> (uint(int32(31)) % 32))
				}
				if v60 == int32(0) {
					v63 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))) = uint8(v63)
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[5])))
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v65))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v46)) == int32(0) {
						v77 = base.B2i32(base.Ui32(v65) < base.Ui32(v46))
					} else {
						v77 = base.B2i32(int32(0) < v46-v65)
					}
					if base.B2i32(v77 == int32(0))|base.B2i32(base.Ui32(v46) < base.Ui32(int32(3))) != 0 {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[5]))) = v46
					}
				}
			}
		}
		v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
		if v166 != int32(1) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v173 = l1 + int32(2364)
			v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			v180 = F_heap_prepare_freeze_tuple(m, v22, v169, l1+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v173+v174*int32(12), v9+int32(15))
			mBase = m.M
			v181 = m.ExcPending
			if v181 != 0 {
				return
			} else {
				if v180 != 0 {
					v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v182 + int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v173+v182*int32(12))+10)) = uint16(v3)
				} else {
				}
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v191 != 0 {
				} else {
					v192 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v192)
				}
				m.G0 = v9 + int32(16)
				return
			}
		}
	case 1:
		v84 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))) = uint8(v84)
		v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[7])))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[7]))) = v86 + int32(1)
		v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)))
		if v90&int32(_a_F_heap_prune_record_unchanged_lp_normal_2) == int32(_a_F_heap_prune_record_unchanged_lp_normal_3) {
			v95 = F_HeapTupleGetUpdateXid(m, v22)
			mBase = m.M
			v96 = m.ExcPending
			if v96 != 0 {
				return
			} else {
				v98 = v95
				v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if v99 != 0 {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v99))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v98)) == int32(0) {
						v111 = base.B2i32(base.Ui32(v98) < base.Ui32(v99))
					} else {
						v111 = int32(base.Ui32(v98-v99) >> (uint(int32(31)) % 32))
					}
					if v111 == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v98
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v98
				}
				v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				if v166 != int32(1) {
					m.G0 = v9 + int32(16)
					return
				} else {
					v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v173 = l1 + int32(2364)
					v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					v180 = F_heap_prepare_freeze_tuple(m, v22, v169, l1+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v173+v174*int32(12), v9+int32(15))
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return
					} else {
						if v180 != 0 {
							v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v182 + int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v173+v182*int32(12))+10)) = uint16(v3)
						} else {
						}
						v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v191 != 0 {
						} else {
							v192 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v192)
						}
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		} else {
			v97 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			v98 = v97
			v99 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v99 != 0 {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v99))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v98)) == int32(0) {
					v111 = base.B2i32(base.Ui32(v98) < base.Ui32(v99))
				} else {
					v111 = int32(base.Ui32(v98-v99) >> (uint(int32(31)) % 32))
				}
				if v111 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v98
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v98
			}
			v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			if v166 != int32(1) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v173 = l1 + int32(2364)
				v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				v180 = F_heap_prepare_freeze_tuple(m, v22, v169, l1+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v173+v174*int32(12), v9+int32(15))
				mBase = m.M
				v181 = m.ExcPending
				if v181 != 0 {
					return
				} else {
					if v180 != 0 {
						v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v182 + int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v173+v182*int32(12))+10)) = uint16(v3)
					} else {
					}
					v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v191 != 0 {
					} else {
						v192 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v192)
					}
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	case 2:
		v162 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))) = uint8(v162)
		v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
		if v166 != int32(1) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v173 = l1 + int32(2364)
			v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			v180 = F_heap_prepare_freeze_tuple(m, v22, v169, l1+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v173+v174*int32(12), v9+int32(15))
			mBase = m.M
			v181 = m.ExcPending
			if v181 != 0 {
				return
			} else {
				if v180 != 0 {
					v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v182 + int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v173+v182*int32(12))+10)) = uint16(v3)
				} else {
				}
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v191 != 0 {
				} else {
					v192 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v192)
				}
				m.G0 = v9 + int32(16)
				return
			}
		}
	case 3:
		v115 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[4]))) = uint8(v115)
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[3])))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[3]))) = v117 + int32(1)
		v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v22)+20)))
		if v121&int32(_a_F_heap_prune_record_unchanged_lp_normal_2) == int32(_a_F_heap_prune_record_unchanged_lp_normal_3) {
			v126 = F_HeapTupleGetUpdateXid(m, v22)
			mBase = m.M
			v127 = m.ExcPending
			if v127 != 0 {
				return
			} else {
				v129 = v126
				v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if v130 != 0 {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v130))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v129)) == int32(0) {
						v142 = base.B2i32(base.Ui32(v129) < base.Ui32(v130))
					} else {
						v142 = int32(base.Ui32(v129-v130) >> (uint(int32(31)) % 32))
					}
					if v142 == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v129
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v129
				}
				v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				if v166 != int32(1) {
					m.G0 = v9 + int32(16)
					return
				} else {
					v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v173 = l1 + int32(2364)
					v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					v180 = F_heap_prepare_freeze_tuple(m, v22, v169, l1+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v173+v174*int32(12), v9+int32(15))
					mBase = m.M
					v181 = m.ExcPending
					if v181 != 0 {
						return
					} else {
						if v180 != 0 {
							v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v182 + int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v173+v182*int32(12))+10)) = uint16(v3)
						} else {
						}
						v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v191 != 0 {
						} else {
							v192 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v192)
						}
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		} else {
			v128 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
			v129 = v128
			v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v130 != 0 {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v130))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v129)) == int32(0) {
					v142 = base.B2i32(base.Ui32(v129) < base.Ui32(v130))
				} else {
					v142 = int32(base.Ui32(v129-v130) >> (uint(int32(31)) % 32))
				}
				if v142 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v129
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v129
			}
			v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			if v166 != int32(1) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v169 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v173 = l1 + int32(2364)
				v174 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				v180 = F_heap_prepare_freeze_tuple(m, v22, v169, l1+int32(_a_F_heap_prune_record_unchanged_lp_normal_1), v173+v174*int32(12), v9+int32(15))
				mBase = m.M
				v181 = m.ExcPending
				if v181 != 0 {
					return
				} else {
					if v180 != 0 {
						v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v182 + int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v173+v182*int32(12))+10)) = uint16(v3)
					} else {
					}
					v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v191 != 0 {
					} else {
						v192 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[6]))) = uint8(v192)
					}
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v149 = m.ExcPending
		if v149 != 0 {
			return
		} else {
			v152 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+uint32(_c_F_heap_prune_record_unchanged_lp_normal[2]))))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v152
			F_errmsg_internal(m, int32(_a_F_heap_prune_record_unchanged_lp_normal_4), v9)
			mBase = m.M
			v156 = m.ExcPending
			if v156 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_heap_prune_record_unchanged_lp_normal_5), int32(1474), int32(_a_F_heap_prune_record_unchanged_lp_normal_6))
				mBase = m.M
				v161 = m.ExcPending
				if v161 != 0 {
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
func F_heap_scan_stream_read_next_parallel(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v13 int64
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v86 int32
	_ = v86
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)))
	if v10 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v13
	v19 = int32(4095)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	if base.Ui32(v20) <= base.Ui32(v19) {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v85 = F_table_block_parallelscan_nextpage(m, v9, v8, v7)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L19
	} else {
		goto L29
	}
L4:
	;
	v23 = v19
	goto L6
L5:
	;
	v23 = v20
	goto L6
L6:
	;
	v25 = int32(base.Ui32(v23) >> (uint(int32(11)) % 32))
	if v25&(v25-int32(1)) != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v33 = int32(2) << (uint(base.I32_clz(v25)^int32(31)) % 32)
	goto L9
L8:
	;
	v33 = v25
	goto L9
L9:
	;
	if base.Ui32(int32(_a_F_heap_scan_stream_read_next_parallel_0)) <= base.Ui32(v33) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v36 = int32(_a_F_heap_scan_stream_read_next_parallel_0)
	goto L12
L11:
	;
	v36 = v33
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v36
	v39 = v7 + int32(24)
	v44 = int32(-1)
	goto L14
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(0)
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l1)+96))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	v79 = F_table_block_parallelscan_nextpage(m, v76, v77, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L19
	} else {
		goto L28
	}
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	*(*int32)(unsafe.Add(mBase, uint32(v39))) = int32(1)
	if v47 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v71
	goto L13
L16:
	;
	F_s_lock(m, v39, int32(_a_F_heap_scan_stream_read_next_parallel_1), int32(455), int32(_a_F_heap_scan_stream_read_next_parallel_2))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	goto L18
L18:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v7)+28))
	if v57 != int32(-1) {
		goto L13
	} else {
		goto L21
	}
L19:
	;
	return int32(0)
L20:
	;
	goto L18
L21:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+12)))
	if v60 != int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L15
L23:
	;
	v71 = int32(0)
	goto L22
L24:
	;
	goto L25
L25:
	;
	if v44 != int32(-1) {
		v71 = v44
		goto L22
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = int32(0)
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v7)+20))
	v69 = F_ss_get_location(m, v9, v68)
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	v44 = v69
	goto L14
L28:
	;
	v81 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+48)) = uint8(v81)
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v79
	return v79
L29:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = v85
	return v85
}
func F_heap_toast_delete(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	v8 = m.G0
	v10 = v8 - int32(_a_F_heap_toast_delete_0)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v14 = v10 + int32(1600)
	F_heap_deform_tuple(m, l1, v12, v14, v10)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v17 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v17 < v19 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v23 = v17
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v10 + int32(_a_F_heap_toast_delete_0)
	return
L6:
	;
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18+v23<<(uint(int32(4))%32))+24)))
	if v32 != int32(_a_F_heap_toast_delete_1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v51 = v23 + int32(1)
	if v51 != v19 {
		v23 = v51
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v10))))
	if v36 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v14+v23<<(uint(int32(2))%32))))
	v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40))))
	if v41 != int32(1) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v40)+1)))
	if v44 != int32(18) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_toast_delete_datum(m, v40, l2)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L8
L14:
	;
	goto L7
}
func F_heap_truncate_find_FKs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v436 int32
	_ = v436
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v449 int32
	_ = v449
	v13 = m.G0
	v15 = v13 - int32(48)
	m.G0 = v15
	v17 = F_list_copy(m, l0)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = F_table_open(m, int32(2606), int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v31 = v17
	v32 = int32(0)
	goto L4
L4:
	;
	v37 = int32(0)
	v42 = F_systable_beginscan(m, v23, v37, v37, v37, v37, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L10
	}
L5:
	;
	F_relation_close(m, v23, int32(1))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L102
	}
L6:
	;
	goto L5
L7:
	;
	F_list_free(m, v323)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L100
	}
L8:
	;
	v230 = v196
	v232 = v205
	v233 = v31
	v237 = v205
	goto L72
L9:
	;
	F_list_free(m, int32(0))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L71
	}
L10:
	;
	v44 = F_systable_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v49 = int32(0)
	v52 = v44
	v54 = v32
	goto L15
L13:
	;
	goto L14
L14:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L70
	}
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	v61 = v59 + v60
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+72)))
	if v62 != int32(102) {
		v196 = v49
		v198 = v54
		goto L17
	} else {
		goto L18
	}
L16:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L67
	}
L17:
	;
	v199 = F_systable_getnext(m, v42)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L65
	}
L18:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+96))
	v66 = int32(0)
	if v31 == v66 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	if v104 == int32(0) {
		v196 = v49
		v198 = v54
		goto L17
	} else {
		goto L32
	}
L20:
	;
	v104 = int32(0)
	goto L19
L21:
	;
	goto L22
L22:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v72 <= int32(0) {
		v98 = v66
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v104 = v98
	goto L19
L24:
	;
	v75 = int32(0)
	if v75 < v72 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v78 = v72
	goto L27
L26:
	;
	v78 = v75
	goto L27
L27:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v81 = int32(0)
	goto L28
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v79+v81<<(uint(int32(2))%32))))
	v90 = base.B2i32(v89 == v65)
	if v89 == v65 {
		v98 = v90
		goto L23
	} else {
		goto L30
	}
L29:
	;
	v98 = v90
	goto L23
L30:
	;
	v92 = v81 + int32(1)
	if v92 != v78 {
		v81 = v92
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v61)+92))
	if v107 == int32(0) {
		v152 = v49
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v61)+80))
	v154 = int32(0)
	if l0 == v154 {
		goto L51
	} else {
		goto L52
	}
L34:
	;
	v110 = int32(0)
	if v49 == v110 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v148 != 0 {
		v152 = v49
		goto L33
	} else {
		goto L48
	}
L36:
	;
	v148 = int32(0)
	goto L35
L37:
	;
	goto L38
L38:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v116 <= int32(0) {
		v142 = v110
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v148 = v142
	goto L35
L40:
	;
	v119 = int32(0)
	if v119 < v116 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v122 = v116
	goto L43
L42:
	;
	v122 = v119
	goto L43
L43:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v125 = int32(0)
	goto L44
L44:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v123+v125<<(uint(int32(2))%32))))
	v134 = base.B2i32(v133 == v107)
	if v133 == v107 {
		v142 = v134
		goto L39
	} else {
		goto L46
	}
L45:
	;
	v142 = v134
	goto L39
L46:
	;
	v136 = v125 + int32(1)
	if v136 != v122 {
		v125 = v136
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v61)+92))
	v150 = F_lappend_oid(m, v49, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v152 = v150
	goto L33
L50:
	;
	if v192 != 0 {
		v196 = v152
		v198 = v54
		goto L17
	} else {
		goto L63
	}
L51:
	;
	v192 = int32(0)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v160 <= int32(0) {
		v186 = v154
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v192 = v186
	goto L50
L55:
	;
	v163 = int32(0)
	if v163 < v160 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v166 = v160
	goto L58
L57:
	;
	v166 = v163
	goto L58
L58:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v169 = int32(0)
	goto L59
L59:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v167+v169<<(uint(int32(2))%32))))
	v178 = base.B2i32(v177 == v153)
	if v177 == v153 {
		v186 = v178
		goto L54
	} else {
		goto L61
	}
L60:
	;
	v186 = v178
	goto L54
L61:
	;
	v180 = v169 + int32(1)
	if v180 != v166 {
		v169 = v180
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v61)+80))
	v194 = F_lappend_oid(m, v54, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L64
	}
L64:
	;
	v196 = v152
	v198 = v194
	goto L17
L65:
	;
	if v199 != 0 {
		v49 = v196
		v52 = v199
		v54 = v198
		goto L15
	} else {
		goto L66
	}
L66:
	;
	goto L16
L67:
	;
	if v196 == int32(0) {
		v219 = v198
		goto L9
	} else {
		goto L68
	}
L68:
	;
	v205 = int32(0)
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v205 < v207 {
		goto L8
	} else {
		goto L69
	}
L69:
	;
	v323 = v196
	v326 = v31
	v330 = v205
	goto L7
L70:
	;
	v219 = v32
	goto L9
L71:
	;
	v340 = v31
	v341 = v219
	goto L6
L72:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v242+v232<<(uint(int32(2))%32))))
	F_ScanKeyInit(m, v15, int32(1), int32(3), int32(184), v246)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	v323 = v310
	v326 = v311
	v330 = v312
	goto L7
L74:
	;
	v250 = int32(1)
	v253 = F_systable_beginscan(m, v23, int32(2667), v250, int32(0), v250, v15)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	F_systable_endscan(m, v253)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L98
	}
L76:
	;
	v255 = F_systable_getnext(m, v253)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	if v255 == int32(0) {
		v310 = v230
		v311 = v233
		v312 = v237
		goto L75
	} else {
		goto L78
	}
L78:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+16))
	v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v259)+22)))
	v261 = v259 + v260
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v261)+92))
	if v262 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v263 = F_list_append_unique_oid(m, v230, v262)
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v261)+96))
	v266 = int32(0)
	if v233 == v266 {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	v310 = v263
	v311 = v233
	v312 = v237
	goto L75
L83:
	;
	if v304 != 0 {
		v310 = v230
		v311 = v233
		v312 = v237
		goto L75
	} else {
		goto L96
	}
L84:
	;
	v304 = int32(0)
	goto L83
L85:
	;
	goto L86
L86:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v233)+4))
	if v272 <= int32(0) {
		v298 = v266
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v304 = v298
	goto L83
L88:
	;
	v275 = int32(0)
	if v275 < v272 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	v278 = v272
	goto L91
L90:
	;
	v278 = v275
	goto L91
L91:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v233)+12))
	v281 = int32(0)
	goto L92
L92:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v279+v281<<(uint(int32(2))%32))))
	v290 = base.B2i32(v289 == v265)
	if v289 == v265 {
		v298 = v290
		goto L87
	} else {
		goto L94
	}
L93:
	;
	v298 = v290
	goto L87
L94:
	;
	v292 = v281 + int32(1)
	if v292 != v278 {
		v281 = v292
		goto L92
	} else {
		goto L95
	}
L95:
	;
	goto L93
L96:
	;
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v261)+96))
	v307 = F_lappend_oid(m, v233, v306)
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L1
	} else {
		goto L97
	}
L97:
	;
	v310 = v230
	v311 = v307
	v312 = int32(1)
	goto L75
L98:
	;
	v317 = v232 + int32(1)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v317 < v318 {
		v230 = v310
		v232 = v317
		v233 = v311
		v237 = v312
		goto L72
	} else {
		goto L99
	}
L99:
	;
	goto L73
L100:
	;
	if v330 != 0 {
		v31 = v326
		v32 = v198
		goto L4
	} else {
		goto L101
	}
L101:
	;
	v340 = v326
	v341 = v198
	goto L6
L102:
	;
	F_list_free(m, v340)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L103
	}
L103:
	;
	F_list_sort(m, v341, int32(467))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	v354 = int32(0)
	if v341 == v354 {
		goto L106
	} else {
		goto L107
	}
L105:
	;
	m.G0 = v15 + int32(48)
	return v341
L106:
	;
	goto L105
L107:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v341)+4))
	if v364 < int32(2) {
		goto L106
	} else {
		goto L108
	}
L108:
	;
	v367 = int32(1)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v341)+12))
	if v364 != int32(2) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v341)+4)) = v449 + int32(1)
	goto L106
L110:
	;
	v371 = int32(1)
	v372 = v364 - v371
	v377 = int32(0)
	v380 = v377
	v382 = v367
	v383 = v377
	goto L113
L111:
	;
	v425 = v354
	v427 = v367
	goto L112
L112:
	;
	v433 = int32(2)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v368+v427<<(uint(v433)%32))))
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v368+v425<<(uint(v433)%32))))
	if v436 == v440 {
		v449 = v425
		goto L109
	} else {
		goto L123
	}
L113:
	;
	v388 = int32(2)
	v390 = v368 + v382<<(uint(v388)%32)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)))
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v368+v380<<(uint(v388)%32))))
	if v391 != v395 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	if v372&v371 == int32(0) {
		v449 = v416
		goto L109
	} else {
		goto L122
	}
L115:
	;
	v398 = v380 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v368+v398<<(uint(int32(2))%32)))) = v391
	v403 = v398
	goto L117
L116:
	;
	v403 = v380
	goto L117
L117:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v368+v403<<(uint(int32(2))%32))))
	if v404 != v408 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v411 = v403 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v368+v411<<(uint(int32(2))%32)))) = v404
	v416 = v411
	goto L120
L119:
	;
	v416 = v403
	goto L120
L120:
	;
	v417 = int32(2)
	v418 = v382 + v417
	v420 = v383 + v417
	if v420 != v372&int32(-2) {
		v380 = v416
		v382 = v418
		v383 = v420
		goto L113
	} else {
		goto L121
	}
L121:
	;
	goto L114
L122:
	;
	v425 = v416
	v427 = v418
	goto L112
L123:
	;
	v443 = v425 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v368+v443<<(uint(int32(2))%32)))) = v436
	v449 = v443
	goto L109
}
func F_heap_update(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
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
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
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
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v392 int32
	_ = v392
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v431 int32
	_ = v431
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	var v520 int32
	_ = v520
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v619 int32
	_ = v619
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v739 int32
	_ = v739
	var v744 int32
	_ = v744
	var v747 int32
	_ = v747
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v781 int32
	_ = v781
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v827 int32
	_ = v827
	var v835 int32
	_ = v835
	var v844 int32
	_ = v844
	var v849 int32
	_ = v849
	var v852 int32
	_ = v852
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v905 int32
	_ = v905
	var v915 int32
	_ = v915
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1002 int32
	_ = v1002
	var v1021 int32
	_ = v1021
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1048 int32
	_ = v1048
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1056 int32
	_ = v1056
	var v1067 int32
	_ = v1067
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1107 int32
	_ = v1107
	var v1118 int32
	_ = v1118
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1161 int32
	_ = v1161
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1181 int32
	_ = v1181
	var v1188 int32
	_ = v1188
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1216 int32
	_ = v1216
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1276 int32
	_ = v1276
	var v1280 int32
	_ = v1280
	var v1283 int32
	_ = v1283
	var v1287 int32
	_ = v1287
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1315 int32
	_ = v1315
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1352 int32
	_ = v1352
	var v1357 int32
	_ = v1357
	var v1362 int32
	_ = v1362
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1370 int32
	_ = v1370
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1398 int32
	_ = v1398
	var v1400 int32
	_ = v1400
	var v1407 int32
	_ = v1407
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1411 int32
	_ = v1411
	var v1414 int32
	_ = v1414
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1428 int32
	_ = v1428
	var v1430 int32
	_ = v1430
	var v1435 int32
	_ = v1435
	var v1440 int32
	_ = v1440
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1473 int32
	_ = v1473
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1485 int32
	_ = v1485
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1524 int32
	_ = v1524
	var v1534 int32
	_ = v1534
	var v1537 int32
	_ = v1537
	var v1541 int32
	_ = v1541
	var v1576 int32
	_ = v1576
	var v1580 int32
	_ = v1580
	var v1584 int32
	_ = v1584
	var v1590 int32
	_ = v1590
	var v1615 int32
	_ = v1615
	var v1616 int32
	_ = v1616
	var v1618 int32
	_ = v1618
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1632 int32
	_ = v1632
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1636 int32
	_ = v1636
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1690 int32
	_ = v1690
	var v1695 int32
	_ = v1695
	var v1704 int32
	_ = v1704
	var v1713 int32
	_ = v1713
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1739 int32
	_ = v1739
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1749 int32
	_ = v1749
	var v1758 int32
	_ = v1758
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1766 int32
	_ = v1766
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1777 int32
	_ = v1777
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1787 int32
	_ = v1787
	var v1788 int32
	_ = v1788
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1795 int32
	_ = v1795
	var v1802 int32
	_ = v1802
	var v1804 int32
	_ = v1804
	var v1805 int32
	_ = v1805
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1812 int32
	_ = v1812
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1830 int32
	_ = v1830
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1850 int32
	_ = v1850
	var v1865 int32
	_ = v1865
	var v1871 int32
	_ = v1871
	var v1874 int64
	_ = v1874
	var v1875 int32
	_ = v1875
	var v1881 int32
	_ = v1881
	var v1883 int32
	_ = v1883
	var v1889 int32
	_ = v1889
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1900 int32
	_ = v1900
	var v1901 int32
	_ = v1901
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1948 int32
	_ = v1948
	var v1951 int32
	_ = v1951
	var v1955 int32
	_ = v1955
	var v1956 int32
	_ = v1956
	var v1957 int32
	_ = v1957
	var v1958 int32
	_ = v1958
	var v1961 int32
	_ = v1961
	var v1963 int32
	_ = v1963
	var v1974 int32
	_ = v1974
	var v1979 int32
	_ = v1979
	var v1988 int32
	_ = v1988
	var v1997 int32
	_ = v1997
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2018 int32
	_ = v2018
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2026 int32
	_ = v2026
	var v2031 int32
	_ = v2031
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2042 int32
	_ = v2042
	var v2047 int32
	_ = v2047
	var v2084 int32
	_ = v2084
	var v2085 int32
	_ = v2085
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2130 int32
	_ = v2130
	var v2134 int32
	_ = v2134
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2155 int32
	_ = v2155
	var v2167 int32
	_ = v2167
	var v2168 int32
	_ = v2168
	var v2170 int32
	_ = v2170
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2179 int32
	_ = v2179
	var v2186 int32
	_ = v2186
	var v2188 int32
	_ = v2188
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2222 int32
	_ = v2222
	var v2227 int32
	_ = v2227
	var v2234 int32
	_ = v2234
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2250 int32
	_ = v2250
	var v2251 int32
	_ = v2251
	var v2253 int32
	_ = v2253
	var v2255 int32
	_ = v2255
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2263 int32
	_ = v2263
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2276 int32
	_ = v2276
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2285 int32
	_ = v2285
	var v2292 int32
	_ = v2292
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2299 int32
	_ = v2299
	var v2301 int32
	_ = v2301
	var v2308 int32
	_ = v2308
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2318 int32
	_ = v2318
	var v2322 int32
	_ = v2322
	var v2334 int32
	_ = v2334
	var v2338 int32
	_ = v2338
	var v2339 int32
	_ = v2339
	var v2341 int32
	_ = v2341
	var v2343 int32
	_ = v2343
	var v2344 int32
	_ = v2344
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2348 int32
	_ = v2348
	var v2349 int32
	_ = v2349
	var v2351 int32
	_ = v2351
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2361 int32
	_ = v2361
	var v2362 int32
	_ = v2362
	var v2364 int32
	_ = v2364
	var v2367 int32
	_ = v2367
	var v2370 int32
	_ = v2370
	var v2371 int32
	_ = v2371
	var v2372 int32
	_ = v2372
	var v2374 int32
	_ = v2374
	var v2376 int32
	_ = v2376
	var v2378 int32
	_ = v2378
	var v2380 int32
	_ = v2380
	var v2381 int32
	_ = v2381
	var v2383 int32
	_ = v2383
	var v2384 int32
	_ = v2384
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2388 int32
	_ = v2388
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2392 int32
	_ = v2392
	var v2393 int32
	_ = v2393
	var v2394 int32
	_ = v2394
	var v2396 int32
	_ = v2396
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2413 int32
	_ = v2413
	var v2419 int32
	_ = v2419
	var v2420 int32
	_ = v2420
	var v2424 int32
	_ = v2424
	var v2427 int32
	_ = v2427
	var v2430 int32
	_ = v2430
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2445 int32
	_ = v2445
	var v2451 int32
	_ = v2451
	var v2453 int32
	_ = v2453
	var v2459 int32
	_ = v2459
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2463 int32
	_ = v2463
	var v2464 int32
	_ = v2464
	var v2468 int32
	_ = v2468
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2476 int32
	_ = v2476
	var v2482 int32
	_ = v2482
	var v2483 int32
	_ = v2483
	var v2487 int32
	_ = v2487
	var v2490 int32
	_ = v2490
	var v2493 int32
	_ = v2493
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2503 int32
	_ = v2503
	var v2508 int32
	_ = v2508
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2524 int32
	_ = v2524
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2536 int32
	_ = v2536
	var v2538 int32
	_ = v2538
	var v2539 int32
	_ = v2539
	var v2540 int32
	_ = v2540
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2548 int32
	_ = v2548
	var v2553 int32
	_ = v2553
	var v2558 int32
	_ = v2558
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2565 int32
	_ = v2565
	var v2572 int32
	_ = v2572
	var v2574 int32
	_ = v2574
	var v2576 int32
	_ = v2576
	var v2583 int32
	_ = v2583
	var v2589 int32
	_ = v2589
	var v2591 int32
	_ = v2591
	var v2597 int32
	_ = v2597
	var v2600 int32
	_ = v2600
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2609 int32
	_ = v2609
	var v2612 int32
	_ = v2612
	var v2618 int32
	_ = v2618
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2626 int32
	_ = v2626
	var v2628 int32
	_ = v2628
	var v2638 int32
	_ = v2638
	var v2644 int32
	_ = v2644
	var v2646 int32
	_ = v2646
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2654 int32
	_ = v2654
	var v2657 int64
	_ = v2657
	var v2658 int64
	_ = v2658
	var v2659 int64
	_ = v2659
	var v2665 int32
	_ = v2665
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2671 int32
	_ = v2671
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2691 int32
	_ = v2691
	var v2692 int32
	_ = v2692
	var v2724 int32
	_ = v2724
	var v2726 int32
	_ = v2726
	var v2729 int32
	_ = v2729
	var v2732 int32
	_ = v2732
	var v2735 int32
	_ = v2735
	var v2776 int32
	_ = v2776
	var v2784 int32
	_ = v2784
	var v2815 int32
	_ = v2815
	var v2819 int32
	_ = v2819
	var v2824 int32
	_ = v2824
	var v2830 int32
	_ = v2830
	var v2841 int32
	_ = v2841
	var v2863 int32
	_ = v2863
	var v2865 int32
	_ = v2865
	var v2867 int32
	_ = v2867
	var v2870 int32
	_ = v2870
	var v2873 int32
	_ = v2873
	var v2875 int32
	_ = v2875
	var v2917 int32
	_ = v2917
	var v2924 int32
	_ = v2924
	var v2958 int32
	_ = v2958
	var v2963 int32
	_ = v2963
	var v2968 int32
	_ = v2968
	var v2969 int32
	_ = v2969
	var v2973 int32
	_ = v2973
	var v2978 int32
	_ = v2978
	var v2979 int32
	_ = v2979
	var v3010 int32
	_ = v3010
	var v3011 int32
	_ = v3011
	var v3017 int32
	_ = v3017
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3027 int32
	_ = v3027
	var v3028 int32
	_ = v3028
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3039 int32
	_ = v3039
	var v3048 int32
	_ = v3048
	var v3049 int32
	_ = v3049
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3054 int32
	_ = v3054
	var v3056 int32
	_ = v3056
	var v3058 int32
	_ = v3058
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3062 int32
	_ = v3062
	var v3066 int32
	_ = v3066
	var v3068 int32
	_ = v3068
	var v3072 int32
	_ = v3072
	var v3087 int32
	_ = v3087
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3095 int32
	_ = v3095
	var v3097 int32
	_ = v3097
	var v3103 int32
	_ = v3103
	var v3108 int32
	_ = v3108
	var v3114 int32
	_ = v3114
	var v3116 int32
	_ = v3116
	var v3132 int32
	_ = v3132
	var v3140 int32
	_ = v3140
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3148 int32
	_ = v3148
	var v3150 int32
	_ = v3150
	var v3152 int32
	_ = v3152
	var v3159 int32
	_ = v3159
	var v3160 int32
	_ = v3160
	var v3161 int32
	_ = v3161
	var v3165 int32
	_ = v3165
	var v3167 int32
	_ = v3167
	var v3168 int32
	_ = v3168
	var v3173 int32
	_ = v3173
	var v3174 int32
	_ = v3174
	var v3176 int32
	_ = v3176
	var v3181 int32
	_ = v3181
	var v3182 int32
	_ = v3182
	var v3183 int32
	_ = v3183
	var v3184 int32
	_ = v3184
	var v3186 int32
	_ = v3186
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3192 int32
	_ = v3192
	var v3194 int32
	_ = v3194
	var v3195 int32
	_ = v3195
	var v3199 int32
	_ = v3199
	var v3203 int32
	_ = v3203
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3213 int32
	_ = v3213
	var v3215 int32
	_ = v3215
	var v3221 int32
	_ = v3221
	var v3222 int32
	_ = v3222
	var v3223 int32
	_ = v3223
	var v3225 int32
	_ = v3225
	var v3229 int32
	_ = v3229
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3235 int32
	_ = v3235
	var v3240 int64
	_ = v3240
	var v3241 int32
	_ = v3241
	var v3247 int32
	_ = v3247
	var v3253 int32
	_ = v3253
	var v3255 int32
	_ = v3255
	var v3261 int32
	_ = v3261
	var v3268 int32
	_ = v3268
	var v3274 int32
	_ = v3274
	var v3276 int32
	_ = v3276
	var v3282 int32
	_ = v3282
	var v3322 int32
	_ = v3322
	var v3324 int32
	_ = v3324
	var v3332 int32
	_ = v3332
	var v3335 int32
	_ = v3335
	var v3339 int32
	_ = v3339
	var v3341 int32
	_ = v3341
	var v3344 int32
	_ = v3344
	var v3348 int32
	_ = v3348
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3353 int32
	_ = v3353
	var v3354 int32
	_ = v3354
	var v3356 int32
	_ = v3356
	var v3359 int32
	_ = v3359
	var v3362 int32
	_ = v3362
	var v3364 int32
	_ = v3364
	var v3365 int32
	_ = v3365
	var v3368 int32
	_ = v3368
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3378 int32
	_ = v3378
	var v3379 int32
	_ = v3379
	var v3381 int32
	_ = v3381
	var v3382 int32
	_ = v3382
	var v3384 int32
	_ = v3384
	var v3386 int32
	_ = v3386
	var v3387 int32
	_ = v3387
	var v3389 int32
	_ = v3389
	var v3392 int32
	_ = v3392
	var v3396 int32
	_ = v3396
	var v3399 int64
	_ = v3399
	var v3408 int32
	_ = v3408
	var v3409 int32
	_ = v3409
	var v3410 int64
	_ = v3410
	var v3420 int32
	_ = v3420
	var v3421 int32
	_ = v3421
	var v3423 int32
	_ = v3423
	var v3426 int32
	_ = v3426
	var v3429 int32
	_ = v3429
	var v3433 int32
	_ = v3433
	var v3439 int32
	_ = v3439
	var v3441 int32
	_ = v3441
	var v3444 int32
	_ = v3444
	var v3480 int32
	_ = v3480
	var v3482 int32
	_ = v3482
	var v3504 int32
	_ = v3504
	var v3505 int32
	_ = v3505
	var v3507 int32
	_ = v3507
	var v3516 int32
	_ = v3516
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3522 int32
	_ = v3522
	var v3524 int32
	_ = v3524
	v10 = int32(0)
	v37 = m.G0
	v39 = v37 - int32(96)
	m.G0 = v39
	*(*int32)(unsafe.Add(mBase, uint32(v39)+52)) = l3
	v42 = F_GetCurrentTransactionId(m)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v46 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+31)) = uint8(v46)
	*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = v46
	*(*int32)(unsafe.Add(mBase, uint32(v39)+20)) = v46
	v54 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[0]))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+72))
	if v55 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_bms_free(m, v3505)
	mBase = m.M
	v3516 = m.ExcPending
	if v3516 != 0 {
		goto L1
	} else {
		goto L714
	}
L4:
	;
	v3480 = v3444
	v3482 = v70
	v3504 = v420
	v3505 = v67
	v3507 = v73
	goto L3
L5:
	;
	if v89 < int32(0) {
		goto L419
	} else {
		goto L420
	}
L6:
	;
	v2084 = *(*int32)(unsafe.Add(mBase, uint32(v1900)))
	v2085 = int32(0)
	v2092 = F_RelationGetBufferForTuple(m, l0, v2084, v89, v2085, v2085, v39+int32(20), v39+int32(24), v2085)
	mBase = m.M
	v2093 = m.ExcPending
	if v2093 != 0 {
		goto L1
	} else {
		goto L417
	}
L7:
	;
	if v58&int32(1) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v58 = int32(1)
	goto L10
L9:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+76)))
	v58 = v57
	goto L10
L10:
	;
	goto L7
L11:
	;
	v64 = F_RelationGetIndexAttrBitmap(m, l0, int32(3))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2035 = m.ExcPending
	if v2035 != 0 {
		goto L1
	} else {
		goto L413
	}
L14:
	;
	v67 = F_RelationGetIndexAttrBitmap(m, l0, int32(4))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v70 = F_RelationGetIndexAttrBitmap(m, l0, int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v73 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v76 = F_bms_add_members(m, int32(0), v64)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v78 = F_bms_add_members(m, v76, v67)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v80 = F_bms_add_members(m, v78, v70)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v82 = F_bms_add_members(m, v80, v73)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v88 = v84 | v85<<(uint(int32(16))%32)
	v89 = F_ReadBuffer(m, l0, v88)
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+10)))
	if v109&int32(4) != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	if v89 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v94+(v89^int32(-1))<<(uint(int32(2))%32))))
	v108 = v100
	goto L22
L25:
	;
	goto L26
L26:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v108 = v102 + v89<<(uint(int32(13))%32) + int32(-8192)
	goto L22
L27:
	;
	F_visibilitymap_pin(m, l0, v88, v39+int32(24))
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_LockBuffer(m, v89, int32(2))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v120 = v108 + int32(20)
	v121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v124 = v120 + v121<<(uint(int32(2))%32)
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	if v125&int32(_a_F_heap_update_0) != int32(_a_F_heap_update_1) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v89)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v144
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v108 + v146&int32(_a_F_heap_update_2)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = int32(base.Ui32(v151) >> (uint(int32(17)) % 32))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v155
	v157 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+40)) = uint16(v157)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v144
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v82 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v132 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_ReleaseBuffer(m, v132)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v135 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)) = uint16(v135)
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v137
	*(*int64)(unsafe.Add(mBase, uint32(l6)+8)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = int32(0)
	v3480 = int32(4)
	v3482 = v67
	v3504 = v73
	v3505 = v64
	v3507 = v70
	goto L3
L39:
	;
	goto L38
L40:
	;
	if int32(0) <= v217 {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v217 = base.I32_ctz(v203) | v204<<(uint(int32(5))%32)
	goto L40
L42:
	;
	v217 = int32(-2)
	goto L40
L43:
	;
	v170 = base.I32_div_s(int32(0), int32(32))
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v171 <= v170 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v174 = v82 + int32(8)
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174+v170<<(uint(int32(2))%32))))
	v181 = v178 & int32(-1)
	if v181 != 0 {
		v203 = v181
		v204 = v170
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v183 = v170 + int32(1)
	if v183 == v171 {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v186 = v183
	goto L47
L47:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v174+v186<<(uint(int32(2))%32))))
	if v193 != 0 {
		v203 = v193
		v204 = v186
		goto L41
	} else {
		goto L49
	}
L48:
	;
	goto L42
L49:
	;
	v195 = v186 + int32(1)
	if v195 != v171 {
		v186 = v195
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v223 = v217
	v244 = v10
	v247 = v10
	goto L54
L52:
	;
	v417 = v10
	v420 = v10
	goto L53
L53:
	;
	v431 = int32(0)
	if base.B2i32(v420 == v431)|base.B2i32(v70 == v431) != 0 {
		v476 = v431
		goto L92
	} else {
		goto L93
	}
L54:
	;
	v261 = v223<<(uint(int32(16))%32) - int32(_a_F_heap_update_3)
	if v261 == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v417 = v334
	v420 = v336
	goto L53
L56:
	;
	if v82 == int32(0) {
		goto L80
	} else {
		goto L81
	}
L57:
	;
	if v312&int32(1)|base.B2i32(v267 < int32(0)) != 0 {
		v334 = v244
		v336 = v247
		goto L56
	} else {
		goto L74
	}
L58:
	;
	v310 = F_bms_add_member(m, v247, v223)
	mBase = m.M
	v311 = m.ExcPending
	if v311 != 0 {
		goto L1
	} else {
		goto L73
	}
L59:
	;
	v267 = v261 >> (uint(int32(16)) % 32)
	if base.B2i32(v261 != int32(-393216))&base.B2i32(v267 < int32(0)) != 0 {
		goto L58
	} else {
		goto L60
	}
L60:
	;
	v275 = F_heap_getattr_1(m, v39+int32(32), v267, v160, v39+int32(80))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v279 = F_heap_getattr_1(m, l2, v267, v160, v39+int32(72))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+72)))
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+80)))
	if (v281|v282)&int32(1) == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if v267 <= int32(0) {
		goto L66
	} else {
		goto L67
	}
L64:
	;
	goto L65
L65:
	;
	if v282&int32(255) == v281 {
		v312 = v282
		goto L57
	} else {
		goto L72
	}
L66:
	;
	if v275 == v279 {
		v312 = int32(0)
		goto L57
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v294 = v160 + int32(4) + v267<<(uint(int32(4))%32)
	v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294)+6)))
	v296 = int32(*(*int16)(unsafe.Add(mBase, uint32(v294)+4)))
	v297 = F_datumIsEqual(m, v275, v279, v295, v296)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L70
	}
L69:
	;
	goto L58
L70:
	;
	if v297 == int32(0) {
		goto L58
	} else {
		goto L71
	}
L71:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+80)))
	v312 = v301
	goto L57
L72:
	;
	goto L58
L73:
	;
	v334 = v244
	v336 = v310
	goto L56
L74:
	;
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v160+v267<<(uint(int32(4))%32))+8)))
	if v321 != int32(_a_F_heap_update_4) {
		v334 = v244
		v336 = v247
		goto L56
	} else {
		goto L75
	}
L75:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v275))))
	if v324 != int32(1) {
		v334 = v244
		v336 = v247
		goto L56
	} else {
		goto L76
	}
L76:
	;
	v327 = F_bms_is_member(m, v223, v73)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v334 = v327 | v244
	v336 = v247
	goto L56
L78:
	;
	if int32(0) <= v392 {
		v223 = v392
		v244 = v334
		v247 = v336
		goto L54
	} else {
		goto L89
	}
L79:
	;
	v392 = base.I32_ctz(v378) | v379<<(uint(int32(5))%32)
	goto L78
L80:
	;
	v392 = int32(-2)
	goto L78
L81:
	;
	v343 = v223 + int32(1)
	v345 = base.I32_div_s(v343, int32(32))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	if v346 <= v345 {
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v349 = v82 + int32(8)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v349+v345<<(uint(int32(2))%32))))
	v356 = v353 & (int32(-1) << (uint(v343) % 32))
	if v356 != 0 {
		v378 = v356
		v379 = v345
		goto L79
	} else {
		goto L83
	}
L83:
	;
	v358 = v345 + int32(1)
	if v358 == v346 {
		goto L80
	} else {
		goto L84
	}
L84:
	;
	v361 = v358
	goto L85
L85:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v349+v361<<(uint(int32(2))%32))))
	if v368 != 0 {
		v378 = v368
		v379 = v361
		goto L79
	} else {
		goto L87
	}
L86:
	;
	goto L80
L87:
	;
	v370 = v361 + int32(1)
	if v370 != v346 {
		v361 = v370
		goto L85
	} else {
		goto L88
	}
L88:
	;
	goto L86
L89:
	;
	goto L55
L90:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v491 = F_HeapTupleSatisfiesUpdate(m, v39+int32(32), v490, v89)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L109
	}
L91:
	;
	if v476 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L92:
	;
	goto L91
L93:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	if v441 < v442 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v444 = v441
	goto L96
L95:
	;
	v444 = v442
	goto L96
L96:
	;
	if v444 <= int32(1) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v447 = int32(1)
	goto L99
L98:
	;
	v447 = v444
	goto L99
L99:
	;
	v448 = int32(8)
	v453 = int32(0)
	goto L100
L100:
	;
	v460 = v453 << (uint(int32(2)) % 32)
	v462 = *(*int32)(unsafe.Add(mBase, uint32(v70+v448+v460)))
	v464 = *(*int32)(unsafe.Add(mBase, uint32(v420+v448+v460)))
	v465 = v462 & v464
	v467 = base.B2i32(v465 != int32(0))
	if v465 != 0 {
		v476 = v467
		goto L92
	} else {
		goto L102
	}
L101:
	;
	v476 = v467
	goto L92
L102:
	;
	v469 = v453 + int32(1)
	if v469 != v447 {
		v453 = v469
		goto L100
	} else {
		goto L103
	}
L103:
	;
	goto L101
L104:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(2)
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(3)
	v487 = int32(5)
	goto L90
L107:
	;
	v487 = int32(4)
	goto L90
L108:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v1293)+4))
	v1295 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1293)+20)))
	v1296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1293)+18)))
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	F_compute_new_xmax_infomask(m, v1294, v1295, v1296, v42, v1297, int32(1), v39+int32(12), v39+int32(10), v39+int32(8))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L1
	} else {
		goto L270
	}
L109:
	;
	if v491 != int32(1) {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v496 = v39 + int32(36)
	v501 = v491
	v520 = int32(0)
	goto L113
L111:
	;
	goto L112
L112:
	;
	F_UnlockReleaseBuffer(m, v89)
	mBase = m.M
	v1276 = m.ExcPending
	if v1276 != 0 {
		goto L1
	} else {
		goto L265
	}
L113:
	;
	v536 = int32(1)
	if base.B2i32(l5 == int32(0))|base.B2i32(v501 != int32(5)) != 0 {
		v954 = v501
		v967 = v536
		v973 = v520
		goto L118
	} else {
		goto L119
	}
L114:
	;
	goto L112
L115:
	;
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v1235 = F_HeapTupleSatisfiesUpdate(m, v39+int32(32), v1234, v89)
	mBase = m.M
	v1236 = m.ExcPending
	if v1236 != 0 {
		goto L1
	} else {
		goto L263
	}
L116:
	;
	v1180 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v1180 != 0 {
		goto L108
	} else {
		goto L258
	}
L117:
	;
	v1037 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1038 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1037)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)) = uint16(v1038)
	v1040 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1040
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v1037)+4))
	v1043 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1037)+20)))
	if v1043&int32(_a_F_heap_update_5) != int32(_a_F_heap_update_6) {
		goto L229
	} else {
		goto L230
	}
L118:
	;
	v989 = int32(0)
	if base.B2i32(l4 == v989)|v954 == v989 {
		goto L222
	} else {
		goto L223
	}
L119:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)+4))
	v544 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v542)+20)))
	if v544&int32(_a_F_heap_update_6) != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v934 = v905 + int32(12)
	v935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v496)+2)))
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v496))))
	v937 = int32(16)
	v940 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+2)))
	v941 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934))))
	if v935|v936<<(uint(v937)%32) == v940|v941<<(uint(v937)%32) {
		goto L215
	} else {
		goto L216
	}
L121:
	;
	v547 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+66)) = uint8(v547)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v552 = F_DoesMultiXactIdConflict(m, v543, v544, v549, v39+int32(66))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L125
	}
L122:
	;
	goto L123
L123:
	;
	v715 = int32(0)
	if base.Ui32(v543) < base.Ui32(int32(3)) {
		goto L155
	} else {
		goto L156
	}
L124:
	;
	if v597&int32(128)|base.B2i32(v597&int32(_a_F_heap_update_7) == int32(64)) != 0 {
		goto L140
	} else {
		goto L141
	}
L125:
	;
	if v552 == int32(0) {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v557 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v556)+20)))
	v597 = v557
	v598 = v556
	v599 = v520
	v600 = int32(0)
	goto L124
L127:
	;
	goto L128
L128:
	;
	F_LockBuffer(m, v89, int32(0))
	mBase = m.M
	v561 = m.ExcPending
	if v561 != 0 {
		goto L1
	} else {
		goto L129
	}
L129:
	;
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+66)))
	if (v562|v520)&int32(1) != 0 {
		goto L131
	} else {
		goto L132
	}
L130:
	;
	v577 = int32(0)
	v582 = F_Do_MultiXactIdWait(m, v543, v487, v544, v577, l0, v496, int32(1), v39+int32(72), v577)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L1
	} else {
		goto L135
	}
L131:
	;
	v576 = v562 ^ int32(1) | v520
	goto L130
L132:
	;
	goto L133
L133:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v569*int32(12))+uint32(_c_F_heap_update[3])))
	F_LockTuple(m, l0, v496, v572)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	v576 = int32(1)
	goto L130
L135:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	F_LockBuffer(m, v89, int32(2))
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v589 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v588)+20)))
	if (v589^v544)&int32(_a_F_heap_update_8) != 0 {
		v1216 = v576
		goto L115
	} else {
		goto L137
	}
L137:
	;
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v588)+4))
	if v593 != v543 {
		v1216 = v576
		goto L115
	} else {
		goto L138
	}
L138:
	;
	v597 = v589
	v598 = v588
	v599 = v576
	v600 = base.B2i32(v584 != int32(0))
	goto L124
L139:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v905 = v714
	v915 = v599
	goto L120
L140:
	;
	v954 = int32(0)
	v967 = v552 ^ int32(1) | v600
	v973 = v599
	goto L118
L141:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v598)+4))
	v612 = F_GetMultiXactIdMembers(m, v608, v39+int32(80), int32(0))
	mBase = m.M
	v613 = m.ExcPending
	if v613 != 0 {
		goto L1
	} else {
		goto L142
	}
L142:
	;
	if v612 <= int32(0) {
		goto L140
	} else {
		goto L143
	}
L143:
	;
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v619 = int32(0)
	goto L145
L144:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(v656)))
	F_pfree(m, v617)
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L150
	}
L145:
	;
	v656 = v617 + v619<<(uint(int32(3))%32)
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v656)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v657) {
		goto L144
	} else {
		goto L147
	}
L146:
	;
	F_pfree(m, v617)
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L149
	}
L147:
	;
	v661 = v619 + int32(1)
	if v661 != v612 {
		v619 = v661
		goto L145
	} else {
		goto L148
	}
L148:
	;
	goto L146
L149:
	;
	goto L140
L150:
	;
	if v665 == int32(0) {
		goto L140
	} else {
		goto L151
	}
L151:
	;
	v670 = F_TransactionIdDidAbort(m, v665)
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	if v670 == int32(0) {
		goto L139
	} else {
		goto L153
	}
L153:
	;
	goto L140
L154:
	;
	if v835|(v476^int32(-1))&base.B2i32(v544&int32(80) == int32(16)) != 0 {
		v954 = v715
		v967 = v536
		v973 = v520
		goto L118
	} else {
		goto L194
	}
L155:
	;
	v835 = int32(0)
	goto L154
L156:
	;
	goto L157
L157:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[4]))
	if v726 == v543 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v835 = int32(1)
	goto L154
L159:
	;
	goto L160
L160:
	;
	v730 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[5]))
	if v730 <= int32(0) {
		goto L162
	} else {
		goto L163
	}
L161:
	;
	v835 = v827
	goto L154
L162:
	;
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[0]))
	if v734 == int32(0) {
		v827 = v715
		goto L161
	} else {
		goto L165
	}
L163:
	;
	goto L164
L164:
	;
	v796 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[6]))
	v798 = int32(0)
	v800 = v730 - int32(1)
	goto L184
L165:
	;
	v739 = v734
	goto L166
L166:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v739)+20))
	if v744 == int32(4) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	v827 = int32(0)
	goto L161
L168:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v739)+80))
	if v791 != 0 {
		v739 = v791
		goto L166
	} else {
		goto L183
	}
L169:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v739)))
	if v747 == int32(0) {
		goto L168
	} else {
		goto L170
	}
L170:
	;
	v750 = int32(1)
	if v543 == v747 {
		v827 = v750
		goto L161
	} else {
		goto L171
	}
L171:
	;
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v739)+52))
	v754 = v752 - int32(1)
	if v754 < int32(0) {
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v759 = int32(0)
	v761 = v754
	goto L173
L173:
	;
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v739)+48))
	v767 = int32(2)
	v768 = base.I32_div_s(v761-v759, v767)
	v769 = v768 + v759
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v765+v769<<(uint(v767)%32))))
	if v773 == v543 {
		v827 = v750
		goto L161
	} else {
		goto L175
	}
L174:
	;
	goto L168
L175:
	;
	v777 = F_TransactionIdPrecedes(m, v773, v543)
	mBase = m.M
	if v777 != 0 {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v778 = v769 + int32(1)
	goto L178
L177:
	;
	v778 = v759
	goto L178
L178:
	;
	if v777 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v781 = v761
	goto L181
L180:
	;
	v781 = v769 - int32(1)
	goto L181
L181:
	;
	if v778 <= v781 {
		v759 = v778
		v761 = v781
		goto L173
	} else {
		goto L182
	}
L182:
	;
	goto L174
L183:
	;
	goto L167
L184:
	;
	v805 = int32(2)
	v806 = base.I32_div_s(v800-v798, v805)
	v807 = v806 + v798
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v796+v807<<(uint(v805)%32))))
	v812 = base.B2i32(v811 == v543)
	if v811 == v543 {
		v827 = v812
		goto L161
	} else {
		goto L186
	}
L185:
	;
	v827 = v812
	goto L161
L186:
	;
	v815 = base.B2i32(base.Ui32(v811) < base.Ui32(v543))
	if base.Ui32(v811) < base.Ui32(v543) {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v816 = v807 + int32(1)
	goto L189
L188:
	;
	v816 = v798
	goto L189
L189:
	;
	if base.Ui32(v811) < base.Ui32(v543) {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v819 = v800
	goto L192
L191:
	;
	v819 = v807 - int32(1)
	goto L192
L192:
	;
	if v816 <= v819 {
		v798 = v816
		v800 = v819
		goto L184
	} else {
		goto L193
	}
L193:
	;
	goto L185
L194:
	;
	F_LockBuffer(m, v89, int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	if v520&int32(1) == int32(0) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v849 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v849*int32(12))+uint32(_c_F_heap_update[3])))
	F_LockTuple(m, l0, v496, v852)
	mBase = m.M
	v854 = m.ExcPending
	if v854 != 0 {
		goto L1
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v855 = int32(1)
	F_XactLockTableWait(m, v543, l0, v496, v855)
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L1
	} else {
		goto L200
	}
L199:
	;
	goto L198
L200:
	;
	F_LockBuffer(m, v89, int32(2))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	v862 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v863 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v862)+20)))
	if (v544^v863)&int32(_a_F_heap_update_8) != 0 {
		v1216 = v855
		goto L115
	} else {
		goto L202
	}
L202:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v862)+4))
	if v543 != v867 {
		v1216 = v855
		goto L115
	} else {
		goto L203
	}
L203:
	;
	if v863&int32(3072) != 0 {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	v889 = int32(0)
	v891 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v892 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v891)+21)))
	if v892&int32(8) != 0 {
		v954 = v889
		v967 = v889
		v973 = v855
		goto L118
	} else {
		goto L212
	}
L205:
	;
	if v863&int32(128)|base.B2i32(v863&int32(_a_F_heap_update_7) == int32(64)) != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	F_HeapTupleSetHintBits(m, v862, v89, int32(2048), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L1
	} else {
		goto L211
	}
L207:
	;
	v878 = F_TransactionIdDidCommit(m, v543)
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L208
	}
L208:
	;
	if v878 == int32(0) {
		goto L206
	} else {
		goto L209
	}
L209:
	;
	F_HeapTupleSetHintBits(m, v862, v89, int32(1024), v543)
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	goto L204
L211:
	;
	goto L204
L212:
	;
	v905 = v891
	v915 = v855
	goto L120
L213:
	;
	if v951 != 0 {
		goto L219
	} else {
		goto L220
	}
L214:
	;
	goto L213
L215:
	;
	v947 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v496)+4)))
	v948 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v934)+4)))
	if v947 == v948 {
		v951 = int32(1)
		goto L214
	} else {
		goto L218
	}
L216:
	;
	goto L217
L217:
	;
	v951 = int32(0)
	goto L214
L218:
	;
	goto L217
L219:
	;
	v952 = int32(4)
	goto L221
L220:
	;
	v952 = int32(3)
	goto L221
L221:
	;
	v1002 = v952
	v1021 = v915
	goto L117
L222:
	;
	v996 = F_HeapTupleSatisfiesVisibility(m, v39+int32(32), l4, v89)
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	if v954 == int32(0) {
		goto L116
	} else {
		goto L227
	}
L225:
	;
	if v996 != 0 {
		goto L116
	} else {
		goto L226
	}
L226:
	;
	v1002 = int32(3)
	v1021 = v973
	goto L117
L227:
	;
	v1002 = v954
	v1021 = v973
	goto L117
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v1118
	if v1002 == int32(2) {
		goto L241
	} else {
		goto L242
	}
L229:
	;
	v1118 = v1042
	goto L228
L230:
	;
	goto L231
L231:
	;
	v1048 = int32(0)
	v1052 = F_GetMultiXactIdMembers(m, v1042, v39+int32(80), v1048)
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	if v1052 <= int32(0) {
		v1118 = v1048
		goto L228
	} else {
		goto L233
	}
L233:
	;
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v1067 = v1048
	goto L236
L234:
	;
	F_pfree(m, v1056)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L1
	} else {
		goto L240
	}
L235:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1095)))
	v1105 = v1103
	goto L234
L236:
	;
	v1095 = v1056 + v1067<<(uint(int32(3))%32)
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1095)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1096) {
		goto L235
	} else {
		goto L238
	}
L237:
	;
	v1105 = int32(0)
	goto L234
L238:
	;
	v1100 = v1067 + int32(1)
	if v1100 != v1052 {
		v1067 = v1100
		goto L236
	} else {
		goto L239
	}
L239:
	;
	goto L237
L240:
	;
	v1118 = v1105
	goto L228
L241:
	;
	v1147 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1147)+8))
	v1150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1147)+20)))
	if v1150&int32(32) != 0 {
		goto L245
	} else {
		goto L246
	}
L242:
	;
	v1161 = int32(-1)
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v1161
	F_UnlockReleaseBuffer(m, v89)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L1
	} else {
		goto L248
	}
L244:
	;
	v1161 = v1159
	goto L243
L245:
	;
	v1154 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[7]))
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v1154+v1149<<(uint(int32(3))%32))+4))
	v1159 = v1158
	goto L247
L246:
	;
	v1159 = v1149
	goto L247
L247:
	;
	goto L244
L248:
	;
	if v1021&int32(1) != 0 {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v1167 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1167*int32(12))+uint32(_c_F_heap_update[3])))
	F_UnlockTuple(m, l0, v496, v1170)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L252
	}
L250:
	;
	goto L251
L251:
	;
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v1173 != 0 {
		goto L253
	} else {
		goto L254
	}
L252:
	;
	goto L251
L253:
	;
	F_ReleaseBuffer(m, v1173)
	mBase = m.M
	v1175 = m.ExcPending
	if v1175 != 0 {
		goto L1
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = int32(0)
	F_bms_free(m, v64)
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L1
	} else {
		goto L257
	}
L256:
	;
	goto L255
L257:
	;
	v3444 = v1002
	goto L4
L258:
	;
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+10)))
	if v1181&int32(4) == int32(0) {
		goto L108
	} else {
		goto L259
	}
L259:
	;
	F_LockBuffer(m, v89, int32(0))
	mBase = m.M
	v1188 = m.ExcPending
	if v1188 != 0 {
		goto L1
	} else {
		goto L260
	}
L260:
	;
	F_visibilitymap_pin(m, l0, v88, v39+int32(24))
	mBase = m.M
	v1192 = m.ExcPending
	if v1192 != 0 {
		goto L1
	} else {
		goto L261
	}
L261:
	;
	F_LockBuffer(m, v89, int32(2))
	mBase = m.M
	v1195 = m.ExcPending
	if v1195 != 0 {
		goto L1
	} else {
		goto L262
	}
L262:
	;
	v1216 = v973
	goto L115
L263:
	;
	if v1235 != int32(1) {
		v501 = v1235
		v520 = v1216
		goto L113
	} else {
		goto L264
	}
L264:
	;
	goto L114
L265:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1280 = m.ExcPending
	if v1280 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_errmsg(m, int32(_a_F_heap_update_9), int32(0))
	mBase = m.M
	v1287 = m.ExcPending
	if v1287 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	F_errfinish(m, int32(_a_F_heap_update_10), int32(3480), int32(_a_F_heap_update_11))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L1
	} else {
		goto L269
	}
L269:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L270:
	;
	v1308 = int32(_a_F_heap_update_12)
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1309)+20)))
	if v1310&int32(2048) != 0 {
		goto L272
	} else {
		goto L273
	}
L271:
	;
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1616 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1615)+20)))
	v1618 = v1616 & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v1615)+20)) = uint16(v1618)
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1620)+18)))
	v1623 = v1621 & int32(_a_F_heap_update_13)
	*(*uint16)(unsafe.Add(mBase, uint32(v1620)+18)) = uint16(v1623)
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1625))) = v42
	v1627 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v1627)+8)) = v1628
	v1630 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1627)+20)))
	v1632 = v1630 & int32(_a_F_heap_update_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v1627)+20)) = uint16(v1632)
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1634)+20)))
	v1636 = v1635 | v1590
	*(*uint16)(unsafe.Add(mBase, uint32(v1634)+20)) = uint16(v1636)
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1639 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1638)+18)))
	v1640 = v1639 | v1580
	*(*uint16)(unsafe.Add(mBase, uint32(v1638)+18)) = uint16(v1640)
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1642)+4)) = v1584
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	F_HeapTupleHeaderAdjustCmax(m, v1644, v39+int32(52), v39+int32(19))
	mBase = m.M
	v1650 = m.ExcPending
	if v1650 != 0 {
		goto L1
	} else {
		goto L326
	}
L272:
	;
	v1313 = int32(0)
	v1580 = v1313
	v1584 = v1313
	v1590 = v1308
	goto L271
L273:
	;
	goto L274
L274:
	;
	v1315 = int32(0)
	if v967&base.B2i32(v1310&int32(_a_F_heap_update_8) != int32(_a_F_heap_update_15)) == v1315 {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1580 = int32(0)
	v1584 = v1315
	v1590 = v1308
	goto L271
L276:
	;
	goto L277
L277:
	;
	v1324 = int32(0)
	v1325 = *(*int32)(unsafe.Add(mBase, uint32(v1309)+4))
	if v1325 == v1324 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1580 = v1324
	v1584 = int32(0)
	v1590 = v1308
	goto L271
L279:
	;
	goto L280
L280:
	;
	if v1310&int32(_a_F_heap_update_6) == int32(0) {
		v1580 = v1324
		v1584 = v1325
		v1590 = int32(_a_F_heap_update_16)
		goto L271
	} else {
		goto L281
	}
L281:
	;
	v1338 = F_GetMultiXactIdMembers(m, v1325, v39+int32(80), int32(0))
	mBase = m.M
	v1339 = m.ExcPending
	if v1339 != 0 {
		goto L1
	} else {
		goto L283
	}
L282:
	;
	v1580 = v1541
	v1584 = v1325
	v1590 = v1576 | int32(_a_F_heap_update_17)
	goto L271
L283:
	;
	if v1338 <= int32(0) {
		v1541 = v1324
		v1576 = int32(_a_F_heap_update_18)
		goto L282
	} else {
		goto L284
	}
L284:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	if v1338 == int32(1) {
		goto L287
	} else {
		goto L288
	}
L285:
	;
	F_pfree(m, v1342)
	mBase = m.M
	v1524 = m.ExcPending
	if v1524 != 0 {
		goto L1
	} else {
		goto L314
	}
L286:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1342+v1445<<(uint(int32(3))%32))+4))
	v1476 = *(*int32)(unsafe.Add(mBase, uint32(v1473<<(uint(int32(2))%32))+uint32(_c_F_heap_update[8])))
	if base.Ui32(v1444) < base.Ui32(v1476) {
		goto L308
	} else {
		goto L309
	}
L287:
	;
	v1345 = int32(0)
	v1435 = v1324
	v1440 = v1345
	v1444 = v1345
	v1445 = v1345
	goto L286
L288:
	;
	goto L289
L289:
	;
	v1352 = int32(0)
	v1357 = v1324
	v1362 = v1352
	v1366 = v1352
	v1367 = v1352
	v1370 = v1352
	goto L290
L290:
	;
	v1394 = v1342 + v1367<<(uint(int32(3))%32)
	v1395 = *(*int32)(unsafe.Add(mBase, uint32(v1394)+4))
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1395<<(uint(int32(2))%32))+uint32(_c_F_heap_update[8])))
	if base.Ui32(v1366) < base.Ui32(v1398) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	if v1338&int32(1) == int32(0) {
		v1488 = v1424
		v1493 = v1425
		v1497 = v1426
		goto L285
	} else {
		goto L307
	}
L292:
	;
	v1400 = v1398
	goto L294
L293:
	;
	v1400 = v1366
	goto L294
L294:
	;
	switch v1395 - int32(3) {
	case 0:
		goto L298
	case 1:
		v1407 = v1357
		goto L296
	case 2:
		goto L297
	default:
		v1409 = v1357
		v1410 = v1362
		goto L295
	}
L295:
	;
	v1411 = *(*int32)(unsafe.Add(mBase, uint32(v1394)+12))
	v1414 = *(*int32)(unsafe.Add(mBase, uint32(v1411<<(uint(int32(2))%32))+uint32(_c_F_heap_update[8])))
	switch v1411 - int32(3) {
	case 0:
		goto L302
	case 1:
		v1422 = v1409
		goto L300
	case 2:
		goto L301
	default:
		v1424 = v1409
		v1425 = v1410
		goto L299
	}
L296:
	;
	v1409 = v1407
	v1410 = int32(1)
	goto L295
L297:
	;
	v1407 = v1357 | int32(_a_F_heap_update_17)
	goto L296
L298:
	;
	v1409 = v1357 | int32(_a_F_heap_update_17)
	v1410 = v1362
	goto L295
L299:
	;
	if base.Ui32(v1400) < base.Ui32(v1414) {
		goto L303
	} else {
		goto L304
	}
L300:
	;
	v1424 = v1422
	v1425 = int32(1)
	goto L299
L301:
	;
	v1422 = v1409 | int32(_a_F_heap_update_17)
	goto L300
L302:
	;
	v1424 = v1409 | int32(_a_F_heap_update_17)
	v1425 = v1410
	goto L299
L303:
	;
	v1426 = v1414
	goto L305
L304:
	;
	v1426 = v1400
	goto L305
L305:
	;
	v1427 = int32(2)
	v1428 = v1367 + v1427
	v1430 = v1370 + v1427
	if v1430 != v1338&int32(2147483646) {
		v1357 = v1424
		v1362 = v1425
		v1366 = v1426
		v1367 = v1428
		v1370 = v1430
		goto L290
	} else {
		goto L306
	}
L306:
	;
	goto L291
L307:
	;
	v1435 = v1424
	v1440 = v1425
	v1444 = v1426
	v1445 = v1428
	goto L286
L308:
	;
	v1478 = v1476
	goto L310
L309:
	;
	v1478 = v1444
	goto L310
L310:
	;
	switch v1473 - int32(3) {
	case 0:
		goto L313
	case 1:
		v1485 = v1435
		goto L311
	case 2:
		goto L312
	default:
		v1488 = v1435
		v1493 = v1440
		v1497 = v1478
		goto L285
	}
L311:
	;
	v1488 = v1485
	v1493 = int32(1)
	v1497 = v1478
	goto L285
L312:
	;
	v1485 = v1435 | int32(_a_F_heap_update_17)
	goto L311
L313:
	;
	v1488 = v1435 | int32(_a_F_heap_update_17)
	v1493 = v1440
	v1497 = v1478
	goto L285
L314:
	;
	if v1497&int32(-2) == int32(2) {
		goto L315
	} else {
		goto L316
	}
L315:
	;
	if v1493 != 0 {
		v1541 = v1488
		v1576 = int32(_a_F_heap_update_19)
		goto L282
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	if v1497 != 0 {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	v1541 = v1488
	v1576 = int32(_a_F_heap_update_20)
	goto L282
L319:
	;
	v1534 = int32(_a_F_heap_update_6)
	goto L321
L320:
	;
	v1534 = int32(_a_F_heap_update_21)
	goto L321
L321:
	;
	if v1497 == int32(1) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1537 = int32(_a_F_heap_update_7)
	goto L324
L323:
	;
	v1537 = v1534
	goto L324
L324:
	;
	if v1493 != 0 {
		v1541 = v1488
		v1576 = v1537
		goto L282
	} else {
		goto L325
	}
L325:
	;
	v1541 = v1488
	v1576 = v1537 | int32(128)
	goto L282
L326:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1652 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1651)+119)))
	switch v1652 - int32(109) {
	case 0, 5:
		goto L328
	default:
		v1667 = int32(0)
		goto L327
	}
L327:
	;
	v1671 = int32(4)
	v1672 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+14)))
	v1673 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+12)))
	v1674 = v1672 - v1673
	if v1674 <= v1671 {
		goto L332
	} else {
		goto L333
	}
L328:
	;
	v1655 = int32(1)
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1657 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1656)+20)))
	if v1657&int32(4) != 0 {
		v1667 = v1655
		goto L327
	} else {
		goto L329
	}
L329:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1661 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1660)+20)))
	if v1661&int32(4) != 0 {
		v1667 = v1655
		goto L327
	} else {
		goto L330
	}
L330:
	;
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1667 = base.B2i32(base.Ui32(int32(2032)) < base.Ui32(v1664))
	goto L327
L331:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1739 = (v1735 + int32(7)) & int32(-8)
	if v1667 == int32(0) {
		goto L350
	} else {
		goto L351
	}
L332:
	;
	v1677 = v1671
	goto L334
L333:
	;
	v1677 = v1674
	goto L334
L334:
	;
	v1679 = v1677 - int32(4)
	if v1679 == int32(0) {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v1734 = int32(0)
	goto L331
L336:
	;
	goto L337
L337:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1673) {
		goto L339
	} else {
		goto L340
	}
L338:
	;
	v1734 = v1679
	goto L331
L339:
	;
	v1690 = int32(base.Ui32(v1673+int32(_a_F_heap_update_22)) >> (uint(int32(2)) % 32))
	goto L341
L340:
	;
	v1690 = int32(0)
	goto L341
L341:
	;
	if base.Ui32(v1690&int32(_a_F_heap_update_4)) < base.Ui32(int32(291)) {
		goto L338
	} else {
		goto L342
	}
L342:
	;
	v1695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+10)))
	if v1695&int32(1) == int32(0) {
		goto L343
	} else {
		goto L344
	}
L343:
	;
	v1734 = int32(0)
	goto L331
L344:
	;
	goto L345
L345:
	;
	v1704 = int32(1)
	goto L346
L346:
	;
	v1713 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108+int32(20)+v1704&int32(_a_F_heap_update_4)<<(uint(int32(2))%32))+1)))
	if v1713&int32(384) == int32(0) {
		goto L338
	} else {
		goto L348
	}
L347:
	;
	v1734 = int32(0)
	goto L331
L348:
	;
	v1719 = v1704 + int32(1)
	v1720 = int32(_a_F_heap_update_4)
	if base.Ui32(v1719&v1720) <= base.Ui32(v1690&v1720) {
		v1704 = v1719
		goto L346
	} else {
		goto L349
	}
L349:
	;
	goto L347
L350:
	;
	if base.Ui32(v1739) <= base.Ui32(v1734) {
		v2095 = l2
		v2130 = v89
		goto L5
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	v1744 = int32(0)
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v1745)+4))
	v1747 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1745)+20)))
	v1748 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1745)+18)))
	v1749 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	F_compute_new_xmax_infomask(m, v1746, v1747, v1748, v42, v1749, v1744, v39+int32(72), v39+int32(66), v39+int32(62))
	mBase = m.M
	v1758 = m.ExcPending
	if v1758 != 0 {
		goto L1
	} else {
		goto L354
	}
L353:
	;
	goto L352
L354:
	;
	v1759 = int32(_a_F_heap_update_23)
	v1761 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_update[9])) = v1761 + int32(1)
	v1765 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1766 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1765)+20)))
	v1768 = v1766 & int32(_a_F_heap_update_24)
	*(*uint16)(unsafe.Add(mBase, uint32(v1765)+20)) = uint16(v1768)
	v1770 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1765)+18)))
	v1772 = v1770 & int32(_a_F_heap_update_25)
	*(*uint16)(unsafe.Add(mBase, uint32(v1765)+18)) = uint16(v1772)
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1775 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1774)+18)))
	v1777 = v1775 & int32(_a_F_heap_update_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v1774)+18)) = uint16(v1777)
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1780 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1779)+4)) = v1780
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1783 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1782)+20)))
	v1784 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+66)))
	v1785 = v1783 | v1784
	*(*uint16)(unsafe.Add(mBase, uint32(v1782)+20)) = uint16(v1785)
	v1787 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1782)+18)))
	v1788 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+62)))
	v1789 = v1787 | v1788
	*(*uint16)(unsafe.Add(mBase, uint32(v1782)+18)) = uint16(v1789)
	v1791 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+19)))
	v1792 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1793 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v1792)+8)) = v1793
	v1795 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1792)+20)))
	v1802 = v1795&int32(_a_F_heap_update_14) | v1791<<(uint(int32(5))%32)&int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v1792)+20)) = uint16(v1802)
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1805 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v496)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1804)+16)) = uint16(v1805)
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v496)))
	*(*int32)(unsafe.Add(mBase, uint32(v1804)+12)) = v1807
	v1809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+10)))
	if v1809&int32(4) != 0 {
		goto L355
	} else {
		goto L356
	}
L355:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v1814 = F_visibilitymap_clear(m, v88, v1812, int32(2))
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L1
	} else {
		goto L358
	}
L356:
	;
	v1816 = v1744
	goto L357
L357:
	;
	F_MarkBufferDirty(m, v89)
	mBase = m.M
	v1818 = m.ExcPending
	if v1818 != 0 {
		goto L1
	} else {
		goto L359
	}
L358:
	;
	v1816 = v1814
	goto L357
L359:
	;
	v1819 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1819)+118)))
	if v1820 != int32(112) {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1881 = int32(_a_F_heap_update_23)
	v1883 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_update[9])) = v1883 - int32(1)
	F_LockBuffer(m, v89, int32(0))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L1
	} else {
		goto L371
	}
L361:
	;
	v1824 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[10]))
	if v1824 <= int32(0) {
		goto L362
	} else {
		goto L363
	}
L362:
	;
	v1827 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1827 != 0 {
		goto L360
	} else {
		goto L365
	}
L363:
	;
	goto L364
L364:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v1830 = m.ExcPending
	if v1830 != 0 {
		goto L1
	} else {
		goto L367
	}
L365:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1828 != 0 {
		goto L360
	} else {
		goto L366
	}
L366:
	;
	goto L364
L367:
	;
	F_XLogRegisterBuffer(m, int32(0), v89, int32(8))
	mBase = m.M
	v1834 = m.ExcPending
	if v1834 != 0 {
		goto L1
	} else {
		goto L368
	}
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = v1780
	v1836 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+40)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+84)) = uint16(v1836)
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1839 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1838)+18)))
	v1840 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1838)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+87)) = uint8(v1816)
	v1846 = int32(1)
	v1848 = int32(8)
	v1850 = int32(4)
	v1865 = int32(base.Ui32(v1839)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v1840)>>(uint(v1846)%32))&v1848 | (int32(base.Ui32(v1840)>>(uint(v1850)%32))&v1850 | (int32(base.Ui32(v1840)>>(uint(int32(12))%32))&v1846 | int32(base.Ui32(v1840)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+86)) = uint8(v1865)
	F_XLogRegisterData(m, v39+int32(80), v1848)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L1
	} else {
		goto L369
	}
L369:
	;
	v1874 = F_XLogInsert(m, int32(10), int32(96))
	mBase = m.M
	v1875 = m.ExcPending
	if v1875 != 0 {
		goto L1
	} else {
		goto L370
	}
L370:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v108))) = base.I64_rotr(v1874, int64(32))
	goto L360
L371:
	;
	if v1667 != 0 {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v1893 = F_heap_toast_insert_or_update(m, l0, l2, v39+int32(32), int32(0))
	mBase = m.M
	v1894 = m.ExcPending
	if v1894 != 0 {
		goto L1
	} else {
		goto L375
	}
L373:
	;
	v1900 = l2
	v1901 = v1739
	goto L374
L374:
	;
	if base.Ui32(v1734) < base.Ui32(v1901) {
		goto L6
	} else {
		goto L376
	}
L375:
	;
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1893)))
	v1900 = v1893
	v1901 = (v1895 + int32(7)) & int32(-8)
	goto L374
L376:
	;
	goto L377
L377:
	;
	v1939 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v1939 != 0 {
		goto L379
	} else {
		goto L380
	}
L379:
	;
	F_LockBuffer(m, v89, int32(2))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L1
	} else {
		goto L383
	}
L380:
	;
	v1940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+10)))
	if v1940&int32(4) == int32(0) {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	F_visibilitymap_pin(m, l0, v88, v39+int32(24))
	mBase = m.M
	v1948 = m.ExcPending
	if v1948 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	goto L379
L383:
	;
	v1955 = int32(4)
	v1956 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+14)))
	v1957 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+12)))
	v1958 = v1956 - v1957
	if v1958 <= v1955 {
		goto L385
	} else {
		goto L386
	}
L384:
	;
	if base.Ui32(v2018) < base.Ui32(v1901) {
		goto L403
	} else {
		goto L404
	}
L385:
	;
	v1961 = v1955
	goto L387
L386:
	;
	v1961 = v1958
	goto L387
L387:
	;
	v1963 = v1961 - int32(4)
	if v1963 == int32(0) {
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v2018 = int32(0)
	goto L384
L389:
	;
	goto L390
L390:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1957) {
		goto L392
	} else {
		goto L393
	}
L391:
	;
	v2018 = v1963
	goto L384
L392:
	;
	v1974 = int32(base.Ui32(v1957+int32(_a_F_heap_update_22)) >> (uint(int32(2)) % 32))
	goto L394
L393:
	;
	v1974 = int32(0)
	goto L394
L394:
	;
	if base.Ui32(v1974&int32(_a_F_heap_update_4)) < base.Ui32(int32(291)) {
		goto L391
	} else {
		goto L395
	}
L395:
	;
	v1979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+10)))
	if v1979&int32(1) == int32(0) {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v2018 = int32(0)
	goto L384
L397:
	;
	goto L398
L398:
	;
	v1988 = int32(1)
	goto L399
L399:
	;
	v1997 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108+int32(20)+v1988&int32(_a_F_heap_update_4)<<(uint(int32(2))%32))+1)))
	if v1997&int32(384) == int32(0) {
		goto L391
	} else {
		goto L401
	}
L400:
	;
	v2018 = int32(0)
	goto L384
L401:
	;
	v2003 = v1988 + int32(1)
	v2004 = int32(_a_F_heap_update_4)
	if base.Ui32(v2003&v2004) <= base.Ui32(v1974&v2004) {
		v1988 = v2003
		goto L399
	} else {
		goto L402
	}
L402:
	;
	goto L400
L403:
	;
	F_LockBuffer(m, v89, int32(0))
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L1
	} else {
		goto L406
	}
L404:
	;
	goto L405
L405:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v2023 == int32(0) {
		goto L408
	} else {
		goto L409
	}
L406:
	;
	goto L6
L407:
	;
	F_LockBuffer(m, v89, int32(0))
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L1
	} else {
		goto L412
	}
L408:
	;
	v2026 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+10)))
	if v2026&int32(4) != 0 {
		goto L407
	} else {
		goto L411
	}
L409:
	;
	goto L410
L410:
	;
	v2095 = v1900
	v2130 = v89
	goto L5
L411:
	;
	goto L410
L412:
	;
	goto L377
L413:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v2038 = m.ExcPending
	if v2038 != 0 {
		goto L1
	} else {
		goto L414
	}
L414:
	;
	F_errmsg(m, int32(_a_F_heap_update_27), int32(0))
	mBase = m.M
	v2042 = m.ExcPending
	if v2042 != 0 {
		goto L1
	} else {
		goto L415
	}
L415:
	;
	F_errfinish(m, int32(_a_F_heap_update_10), int32(3303), int32(_a_F_heap_update_11))
	mBase = m.M
	v2047 = m.ExcPending
	if v2047 != 0 {
		goto L1
	} else {
		goto L416
	}
L416:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L417:
	;
	v2095 = v1900
	v2130 = v2092
	goto L5
L418:
	;
	F_CheckForSerializableConflictIn(m, l0, v496, v2149)
	mBase = m.M
	v2151 = m.ExcPending
	if v2151 != 0 {
		goto L1
	} else {
		goto L422
	}
L419:
	;
	v2134 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[11]))
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v2134+(v89^int32(-1))<<(uint(int32(6))%32))+16))
	v2149 = v2140
	goto L418
L420:
	;
	goto L421
L421:
	;
	v2142 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[12]))
	v2148 = *(*int32)(unsafe.Add(mBase, uint32(v2142+v89<<(uint(int32(6))%32)+int32(-64))+16))
	v2149 = v2148
	goto L418
L422:
	;
	v2152 = base.B2i32(v2130 != v89)
	if v2152 == int32(0) {
		goto L424
	} else {
		goto L425
	}
L423:
	;
	v2263 = int32(0)
	if base.B2i32(v420 == v2263)|base.B2i32(v73 == v2263) != 0 {
		v2308 = v2263
		goto L458
	} else {
		goto L459
	}
L424:
	;
	v2155 = int32(0)
	if base.B2i32(v420 == v2155)|base.B2i32(v64 == v2155) != 0 {
		v2202 = v2155
		goto L428
	} else {
		goto L429
	}
L425:
	;
	goto L426
L426:
	;
	v2253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+10)))
	v2255 = v2253 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v108)+10)) = uint16(v2255)
	v2257 = int32(0)
	v2259 = v2257
	v2260 = v2257
	goto L423
L427:
	;
	if v2202 != 0 {
		v2259 = v2155
		v2260 = v2155
		goto L423
	} else {
		goto L440
	}
L428:
	;
	goto L427
L429:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v2167 < v2168 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v2170 = v2167
	goto L432
L431:
	;
	v2170 = v2168
	goto L432
L432:
	;
	if v2170 <= int32(1) {
		goto L433
	} else {
		goto L434
	}
L433:
	;
	v2173 = int32(1)
	goto L435
L434:
	;
	v2173 = v2170
	goto L435
L435:
	;
	v2174 = int32(8)
	v2179 = int32(0)
	goto L436
L436:
	;
	v2186 = v2179 << (uint(int32(2)) % 32)
	v2188 = *(*int32)(unsafe.Add(mBase, uint32(v64+v2174+v2186)))
	v2190 = *(*int32)(unsafe.Add(mBase, uint32(v420+v2174+v2186)))
	v2191 = v2188 & v2190
	v2193 = base.B2i32(v2191 != int32(0))
	if v2191 != 0 {
		v2202 = v2193
		goto L428
	} else {
		goto L438
	}
L437:
	;
	v2202 = v2193
	goto L428
L438:
	;
	v2195 = v2179 + int32(1)
	if v2195 != v2173 {
		v2179 = v2195
		goto L436
	} else {
		goto L439
	}
L439:
	;
	goto L437
L440:
	;
	v2204 = int32(0)
	if base.B2i32(v420 == v2204)|base.B2i32(v67 == v2204) != 0 {
		v2250 = v2204
		goto L442
	} else {
		goto L443
	}
L441:
	;
	if v2250 != 0 {
		goto L454
	} else {
		goto L455
	}
L442:
	;
	goto L441
L443:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(v67)+4))
	if v2215 < v2216 {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v2218 = v2215
	goto L446
L445:
	;
	v2218 = v2216
	goto L446
L446:
	;
	if v2218 <= int32(1) {
		goto L447
	} else {
		goto L448
	}
L447:
	;
	v2221 = int32(1)
	goto L449
L448:
	;
	v2221 = v2218
	goto L449
L449:
	;
	v2222 = int32(8)
	v2227 = int32(0)
	goto L450
L450:
	;
	v2234 = v2227 << (uint(int32(2)) % 32)
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v67+v2222+v2234)))
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v420+v2222+v2234)))
	v2239 = v2236 & v2238
	v2241 = base.B2i32(v2239 != int32(0))
	if v2239 != 0 {
		v2250 = v2241
		goto L442
	} else {
		goto L452
	}
L451:
	;
	v2250 = v2241
	goto L442
L452:
	;
	v2243 = v2227 + int32(1)
	if v2243 != v2221 {
		v2227 = v2243
		goto L450
	} else {
		goto L453
	}
L453:
	;
	goto L451
L454:
	;
	v2251 = int32(2)
	goto L456
L455:
	;
	v2251 = v2204
	goto L456
L456:
	;
	v2259 = v2251
	v2260 = int32(1)
	goto L423
L457:
	;
	v2314 = F_ExtractReplicaIdentity(m, l0, v39+int32(32), (v2308|v417)&int32(1), v39+int32(31))
	mBase = m.M
	v2315 = m.ExcPending
	if v2315 != 0 {
		goto L1
	} else {
		goto L470
	}
L458:
	;
	goto L457
L459:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v420)+4))
	v2274 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	if v2273 < v2274 {
		goto L460
	} else {
		goto L461
	}
L460:
	;
	v2276 = v2273
	goto L462
L461:
	;
	v2276 = v2274
	goto L462
L462:
	;
	if v2276 <= int32(1) {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v2279 = int32(1)
	goto L465
L464:
	;
	v2279 = v2276
	goto L465
L465:
	;
	v2280 = int32(8)
	v2285 = int32(0)
	goto L466
L466:
	;
	v2292 = v2285 << (uint(int32(2)) % 32)
	v2294 = *(*int32)(unsafe.Add(mBase, uint32(v73+v2280+v2292)))
	v2296 = *(*int32)(unsafe.Add(mBase, uint32(v420+v2280+v2292)))
	v2297 = v2294 & v2296
	v2299 = base.B2i32(v2297 != int32(0))
	if v2297 != 0 {
		v2308 = v2299
		goto L458
	} else {
		goto L468
	}
L467:
	;
	v2308 = v2299
	goto L458
L468:
	;
	v2301 = v2285 + int32(1)
	if v2301 != v2279 {
		v2285 = v2301
		goto L466
	} else {
		goto L469
	}
L469:
	;
	goto L467
L470:
	;
	v2316 = int32(_a_F_heap_update_23)
	v2318 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_update[9])) = v2318 + int32(1)
	v2322 = *(*int32)(unsafe.Add(mBase, uint32(v120)))
	if v2322 != 0 {
		goto L472
	} else {
		goto L473
	}
L471:
	;
	v2338 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2339 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2338)+18)))
	if v2260 != 0 {
		goto L481
	} else {
		goto L482
	}
L472:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2322))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v42)) == int32(0) {
		goto L476
	} else {
		goto L477
	}
L473:
	;
	goto L474
L474:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120))) = v42
	goto L471
L475:
	;
	if v2334 == int32(0) {
		goto L471
	} else {
		goto L479
	}
L476:
	;
	v2334 = base.B2i32(base.Ui32(v42) < base.Ui32(v2322))
	goto L475
L477:
	;
	goto L478
L478:
	;
	v2334 = int32(base.Ui32(v42-v2322) >> (uint(int32(31)) % 32))
	goto L475
L479:
	;
	goto L474
L480:
	;
	v2367 = int32(0)
	F_RelationPutHeapTuple(m, v2130, v2095, v2367)
	mBase = m.M
	v2370 = m.ExcPending
	if v2370 != 0 {
		goto L1
	} else {
		goto L484
	}
L481:
	;
	v2341 = v2339 | int32(_a_F_heap_update_28)
	*(*uint16)(unsafe.Add(mBase, uint32(v2338)+18)) = uint16(v2341)
	v2343 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+16))
	v2344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2343)+18)))
	v2345 = int32(_a_F_heap_update_1)
	v2346 = v2344 | v2345
	*(*uint16)(unsafe.Add(mBase, uint32(v2343)+18)) = uint16(v2346)
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2348)+18)))
	v2351 = v2349 | v2345
	*(*uint16)(unsafe.Add(mBase, uint32(v2348)+18)) = uint16(v2351)
	goto L480
L482:
	;
	goto L483
L483:
	;
	v2354 = v2339 & int32(_a_F_heap_update_26)
	*(*uint16)(unsafe.Add(mBase, uint32(v2338)+18)) = uint16(v2354)
	v2356 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+16))
	v2357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2356)+18)))
	v2358 = int32(_a_F_heap_update_2)
	v2359 = v2357 & v2358
	*(*uint16)(unsafe.Add(mBase, uint32(v2356)+18)) = uint16(v2359)
	v2361 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2362 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2361)+18)))
	v2364 = v2362 & v2358
	*(*uint16)(unsafe.Add(mBase, uint32(v2361)+18)) = uint16(v2364)
	goto L480
L484:
	;
	v2371 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2371)+20)))
	v2374 = v2372 & int32(_a_F_heap_update_24)
	*(*uint16)(unsafe.Add(mBase, uint32(v2371)+20)) = uint16(v2374)
	v2376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2371)+18)))
	v2378 = v2376 & int32(_a_F_heap_update_25)
	*(*uint16)(unsafe.Add(mBase, uint32(v2371)+18)) = uint16(v2378)
	v2380 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2380)+4)) = v2381
	v2383 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2383)+20)))
	v2385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+10)))
	v2386 = v2384 | v2385
	*(*uint16)(unsafe.Add(mBase, uint32(v2383)+20)) = uint16(v2386)
	v2388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2383)+18)))
	v2389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)))
	v2390 = v2388 | v2389
	*(*uint16)(unsafe.Add(mBase, uint32(v2383)+18)) = uint16(v2390)
	v2392 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+19)))
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2394 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v2393)+8)) = v2394
	v2396 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)))
	v2403 = v2396&int32(_a_F_heap_update_14) | v2392<<(uint(int32(5))%32)&int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v2393)+20)) = uint16(v2403)
	v2405 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2406 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2095)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2405)+16)) = uint16(v2406)
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2405)+12)) = v2408
	if v89 < int32(0) {
		goto L487
	} else {
		goto L488
	}
L485:
	;
	v2469 = int32(0)
	v2470 = base.B2i32(v2130 == v89)
	if v2470 == v2469 {
		goto L497
	} else {
		goto L498
	}
L486:
	;
	v2440 = v2438 & int32(_a_F_heap_update_29)
	*(*uint16)(unsafe.Add(mBase, uint32(v2437)+10)) = uint16(v2440)
	if v89 < int32(0) {
		goto L493
	} else {
		goto L494
	}
L487:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v2419 = *(*int32)(unsafe.Add(mBase, uint32(v2413+(v89^int32(-1))<<(uint(int32(2))%32))))
	v2420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2419)+10)))
	if v2420&int32(4) != 0 {
		v2437 = v2419
		v2438 = v2420
		goto L486
	} else {
		goto L490
	}
L488:
	;
	goto L489
L489:
	;
	v2424 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v2427 = v2424 + v89<<(uint(int32(13))%32)
	v2430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2427-int32(_a_F_heap_update_30)))))
	if v2430&int32(4) == int32(0) {
		v2468 = v2367
		goto L485
	} else {
		goto L491
	}
L490:
	;
	v2468 = v2367
	goto L485
L491:
	;
	v2437 = v2427 + int32(-8192)
	v2438 = v2430
	goto L486
L492:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v2463 = F_visibilitymap_clear(m, v2460, v2461, int32(3))
	mBase = m.M
	v2464 = m.ExcPending
	if v2464 != 0 {
		goto L1
	} else {
		goto L496
	}
L493:
	;
	v2445 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[11]))
	v2451 = *(*int32)(unsafe.Add(mBase, uint32(v2445+(v89^int32(-1))<<(uint(int32(6))%32))+16))
	v2460 = v2451
	goto L492
L494:
	;
	goto L495
L495:
	;
	v2453 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[12]))
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2453+v89<<(uint(int32(6))%32)+int32(-64))+16))
	v2460 = v2459
	goto L492
L496:
	;
	v2468 = int32(1)
	goto L485
L497:
	;
	if v2130 < int32(0) {
		goto L502
	} else {
		goto L503
	}
L498:
	;
	v2536 = v2469
	goto L499
L499:
	;
	F_MarkBufferDirty(m, v89)
	mBase = m.M
	v2538 = m.ExcPending
	if v2538 != 0 {
		goto L1
	} else {
		goto L513
	}
L500:
	;
	F_MarkBufferDirty(m, v2130)
	mBase = m.M
	v2533 = m.ExcPending
	if v2533 != 0 {
		goto L1
	} else {
		goto L512
	}
L501:
	;
	v2503 = v2501 & int32(_a_F_heap_update_29)
	*(*uint16)(unsafe.Add(mBase, uint32(v2500)+10)) = uint16(v2503)
	if v2130 < int32(0) {
		goto L508
	} else {
		goto L509
	}
L502:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v2482 = *(*int32)(unsafe.Add(mBase, uint32(v2476+(v2130^int32(-1))<<(uint(int32(2))%32))))
	v2483 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2482)+10)))
	if v2483&int32(4) != 0 {
		v2500 = v2482
		v2501 = v2483
		goto L501
	} else {
		goto L505
	}
L503:
	;
	goto L504
L504:
	;
	v2487 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v2490 = v2487 + v2130<<(uint(int32(13))%32)
	v2493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2490-int32(_a_F_heap_update_30)))))
	if v2493&int32(4) == int32(0) {
		v2531 = v2469
		goto L500
	} else {
		goto L506
	}
L505:
	;
	v2531 = v2469
	goto L500
L506:
	;
	v2500 = v2490 + int32(-8192)
	v2501 = v2493
	goto L501
L507:
	;
	v2524 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v2526 = F_visibilitymap_clear(m, v2523, v2524, int32(3))
	mBase = m.M
	v2527 = m.ExcPending
	if v2527 != 0 {
		goto L1
	} else {
		goto L511
	}
L508:
	;
	v2508 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[11]))
	v2514 = *(*int32)(unsafe.Add(mBase, uint32(v2508+(v2130^int32(-1))<<(uint(int32(6))%32))+16))
	v2523 = v2514
	goto L507
L509:
	;
	goto L510
L510:
	;
	v2516 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[12]))
	v2522 = *(*int32)(unsafe.Add(mBase, uint32(v2516+v2130<<(uint(int32(6))%32)+int32(-64))+16))
	v2523 = v2522
	goto L507
L511:
	;
	v2531 = int32(1)
	goto L500
L512:
	;
	v2536 = v2531
	goto L499
L513:
	;
	v2539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2539)+118)))
	if v2540 != int32(112) {
		goto L514
	} else {
		goto L515
	}
L514:
	;
	v3322 = int32(_a_F_heap_update_23)
	v3324 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[9]))
	*(*int32)(unsafe.Add(mBase, _c_F_heap_update[9])) = v3324 - int32(1)
	if v2470 == int32(0) {
		goto L662
	} else {
		goto L663
	}
L515:
	;
	v2544 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[10]))
	if v2544 <= int32(0) {
		goto L517
	} else {
		goto L518
	}
L516:
	;
	v2576 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v2576)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)) = uint16(v2576)
	if v2130 < v2576 {
		goto L533
	} else {
		goto L534
	}
L517:
	;
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2547 != 0 {
		goto L514
	} else {
		goto L520
	}
L518:
	;
	goto L519
L519:
	;
	if v2544 == int32(1) {
		goto L516
	} else {
		goto L522
	}
L520:
	;
	v2548 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2548 == int32(0) {
		goto L516
	} else {
		goto L521
	}
L521:
	;
	goto L514
L522:
	;
	v2553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L523
L523:
	;
	if base.B2i32(base.Ui32(v2553) < base.Ui32(int32(_a_F_heap_update_31))) == int32(0) {
		goto L524
	} else {
		goto L525
	}
L524:
	;
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v2558 == int32(0) {
		goto L516
	} else {
		goto L527
	}
L525:
	;
	goto L526
L526:
	;
	F_log_heap_new_cid(m, l0, v39+int32(32))
	mBase = m.M
	v2572 = m.ExcPending
	if v2572 != 0 {
		goto L1
	} else {
		goto L530
	}
L527:
	;
	v2561 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2561)+119)))
	switch v2562 - int32(109) {
	case 0, 5:
		goto L528
	default:
		goto L516
	}
L528:
	;
	v2565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2558)+104)))
	if v2565 != int32(1) {
		goto L516
	} else {
		goto L529
	}
L529:
	;
	goto L526
L530:
	;
	F_log_heap_new_cid(m, l0, v2095)
	mBase = m.M
	v2574 = m.ExcPending
	if v2574 != 0 {
		goto L1
	} else {
		goto L531
	}
L531:
	;
	goto L516
L532:
	;
	v2600 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[10]))
	if v2600 < int32(2) {
		v2618 = int32(0)
		goto L536
	} else {
		goto L537
	}
L533:
	;
	v2583 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v2589 = *(*int32)(unsafe.Add(mBase, uint32(v2583+(v2130^int32(-1))<<(uint(int32(2))%32))))
	v2597 = v2589
	goto L532
L534:
	;
	goto L535
L535:
	;
	v2591 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v2597 = v2591 + v2130<<(uint(int32(13))%32) + int32(-8192)
	goto L532
L536:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L1
	} else {
		goto L541
	}
L537:
	;
	v2604 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2605 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2604)+118)))
	if v2605 != int32(112) {
		v2618 = int32(0)
		goto L536
	} else {
		goto L538
	}
L538:
	;
	v2609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2604)+119)))
	if v2609 == int32(102) {
		v2618 = int32(0)
		goto L536
	} else {
		goto L539
	}
L539:
	;
	v2612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L540
L540:
	;
	v2618 = base.B2i32(base.Ui32(v2612) < base.Ui32(int32(_a_F_heap_update_31))) ^ int32(1)
	goto L536
L541:
	;
	v2621 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+16))
	v2622 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2621)+18)))
	if v2618|v2152 != 0 {
		goto L543
	} else {
		goto L544
	}
L542:
	;
	v3011 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+87)) = uint8(v3010)
	if v2618 == v3011 {
		goto L593
	} else {
		goto L594
	}
L543:
	;
	v2969 = int32(0)
	if v2536 != 0 {
		goto L590
	} else {
		goto L591
	}
L544:
	;
	v2626 = m.G0
	v2628 = v2626 - int32(16)
	m.G0 = v2628
	F_GetFullPageWriteInfo(m, v2628+int32(8), v2628+int32(7))
	mBase = m.M
	if v89 < int32(0) {
		goto L547
	} else {
		goto L548
	}
L545:
	;
	if v2665 != 0 {
		goto L543
	} else {
		goto L555
	}
L546:
	;
	v2653 = int32(1)
	v2654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2628)+7)))
	if v2654 == v2653 {
		goto L551
	} else {
		goto L552
	}
L547:
	;
	v2638 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v2644 = *(*int32)(unsafe.Add(mBase, uint32(v2638+(v89^int32(-1))<<(uint(int32(2))%32))))
	v2652 = v2644
	goto L546
L548:
	;
	goto L549
L549:
	;
	v2646 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v2652 = v2646 + v89<<(uint(int32(13))%32) + int32(-8192)
	goto L546
L550:
	;
	m.G0 = v2628 + int32(16)
	goto L545
L551:
	;
	v2657 = *(*int64)(unsafe.Add(mBase, uint32(v2628)+8))
	v2658 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2652)+4)))
	v2659 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2652))))
	if base.Ui64(v2658|v2659<<(uint(int64(32))%64)) <= base.Ui64(v2657) {
		v2665 = v2653
		goto L550
	} else {
		goto L554
	}
L552:
	;
	goto L553
L553:
	;
	v2665 = int32(0)
	goto L550
L554:
	;
	goto L553
L555:
	;
	v2669 = *(*int32)(unsafe.Add(mBase, uint32(v2095)))
	v2670 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+16))
	v2671 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2670)+22)))
	v2672 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2672)+22)))
	v2674 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v2674)
	v2677 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v2678 = v2677 - v2673
	v2679 = v2669 - v2671
	if v2678 < v2679 {
		goto L557
	} else {
		goto L558
	}
L556:
	;
	v2815 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)) = uint16(v2815)
	v2819 = v2681 - v2784&int32(_a_F_heap_update_4)
	if v2815 < v2819 {
		goto L571
	} else {
		goto L572
	}
L557:
	;
	v2681 = v2678
	goto L559
L558:
	;
	v2681 = v2679
	goto L559
L559:
	;
	if int32(0) < v2681 {
		goto L560
	} else {
		goto L561
	}
L560:
	;
	v2691 = int32(0)
	v2692 = v2674
	goto L563
L561:
	;
	goto L562
L562:
	;
	v2776 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v2776)
	v2784 = v2776
	goto L556
L563:
	;
	v2724 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2691+(v2671+v2670)))))
	v2726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2691+(v2673+v2672)))))
	if v2724 == v2726 {
		goto L565
	} else {
		goto L566
	}
L564:
	;
	if base.Ui32(int32(2)) < base.Ui32(v2735&int32(_a_F_heap_update_4)) {
		v2784 = v2735
		goto L556
	} else {
		goto L569
	}
L565:
	;
	v2729 = v2692 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v2729)
	v2732 = v2729 & int32(_a_F_heap_update_4)
	if base.Ui32(v2732) < base.Ui32(v2681) {
		v2691 = v2732
		v2692 = v2729
		goto L563
	} else {
		goto L568
	}
L566:
	;
	v2735 = v2692
	goto L567
L567:
	;
	goto L564
L568:
	;
	v2735 = v2729
	goto L567
L569:
	;
	goto L562
L570:
	;
	if v2536 != 0 {
		goto L581
	} else {
		goto L582
	}
L571:
	;
	v2824 = int32(0)
	v2830 = v2824
	v2841 = v2824
	goto L574
L572:
	;
	goto L573
L573:
	;
	v2917 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)) = uint16(v2917)
	v2924 = v2917
	goto L570
L574:
	;
	v2863 = v2841 ^ int32(-1)
	v2865 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2669+v2670+v2863))))
	v2867 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2863+(v2672+v2677)))))
	if v2865 == v2867 {
		goto L576
	} else {
		goto L577
	}
L575:
	;
	if base.Ui32(int32(2)) < base.Ui32(v2875&int32(_a_F_heap_update_4)) {
		v2924 = v2875
		goto L570
	} else {
		goto L580
	}
L576:
	;
	v2870 = v2830 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)) = uint16(v2870)
	v2873 = v2870 & int32(_a_F_heap_update_4)
	if base.Ui32(v2873) < base.Ui32(v2819) {
		v2830 = v2870
		v2841 = v2873
		goto L574
	} else {
		goto L579
	}
L577:
	;
	v2875 = v2830
	goto L578
L578:
	;
	goto L575
L579:
	;
	v2875 = v2870
	goto L578
L580:
	;
	goto L573
L581:
	;
	v2958 = v2468 | int32(2)
	goto L583
L582:
	;
	v2958 = v2468
	goto L583
L583:
	;
	if v2784&int32(_a_F_heap_update_4) != 0 {
		goto L584
	} else {
		goto L585
	}
L584:
	;
	v2963 = v2958 | int32(32)
	goto L586
L585:
	;
	v2963 = v2958
	goto L586
L586:
	;
	if v2924&int32(_a_F_heap_update_4) != 0 {
		goto L587
	} else {
		goto L588
	}
L587:
	;
	v2968 = v2963 | int32(64)
	goto L589
L588:
	;
	v2968 = v2963
	goto L589
L589:
	;
	v2978 = v2924
	v2979 = v2784
	v3010 = v2968
	goto L542
L590:
	;
	v2973 = v2468 | int32(2)
	goto L592
L591:
	;
	v2973 = v2468
	goto L592
L592:
	;
	v2978 = v2969
	v2979 = v2969
	v3010 = v2973
	goto L542
L593:
	;
	if v2622 < v3011 {
		goto L599
	} else {
		goto L600
	}
L594:
	;
	v3017 = v3010 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+87)) = uint8(v3017)
	if v2314 == int32(0) {
		goto L593
	} else {
		goto L595
	}
L595:
	;
	v3023 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3023)+130)))
	if v3024 == int32(102) {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	v3027 = int32(20)
	goto L598
L597:
	;
	v3027 = int32(24)
	goto L598
L598:
	;
	v3028 = v3027 | v3010
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+87)) = uint8(v3028)
	goto L593
L599:
	;
	v3032 = int32(64)
	goto L601
L600:
	;
	v3032 = int32(32)
	goto L601
L601:
	;
	v3033 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2095)+8)))
	if v3033 != int32(1) {
		goto L603
	} else {
		goto L604
	}
L602:
	;
	v3056 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+40)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+84)) = uint16(v3056)
	v3058 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v3059 = *(*int32)(unsafe.Add(mBase, uint32(v3058)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = v3059
	v3061 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3058)+18)))
	v3062 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3058)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+92)) = uint16(v3033)
	v3066 = int32(16)
	v3068 = int32(1)
	v3072 = int32(4)
	v3087 = int32(base.Ui32(v3061)>>(uint(int32(9))%32))&v3066 | (int32(base.Ui32(v3062)>>(uint(v3068)%32))&int32(8) | (int32(base.Ui32(v3062)>>(uint(v3072)%32))&v3072 | (int32(base.Ui32(v3062)>>(uint(int32(12))%32))&v3068 | int32(base.Ui32(v3062)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+86)) = uint8(v3087)
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+16))
	v3090 = *(*int32)(unsafe.Add(mBase, uint32(v3089)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+88)) = v3090
	if v2618 != 0 {
		goto L612
	} else {
		goto L613
	}
L603:
	;
	v3053 = v3032
	v3054 = int32(8)
	goto L602
L604:
	;
	goto L605
L605:
	;
	v3039 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2597)+12)))
	v3048 = base.B2i32(base.Ui32(int32(24)) < base.Ui32(v3039)) & base.B2i32((v3039+int32(_a_F_heap_update_22))&int32(_a_F_heap_update_32) == int32(4))
	if v3048 != 0 {
		goto L606
	} else {
		goto L607
	}
L606:
	;
	v3049 = int32(14)
	goto L608
L607:
	;
	v3049 = int32(8)
	goto L608
L608:
	;
	if v3048 != 0 {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	v3052 = v3032 | int32(-128)
	goto L611
L610:
	;
	v3052 = v3032
	goto L611
L611:
	;
	v3053 = v3052
	v3054 = v3049
	goto L602
L612:
	;
	v3095 = v3054 | v3066
	goto L614
L613:
	;
	v3095 = v3054
	goto L614
L614:
	;
	F_XLogRegisterBuffer(m, int32(0), v2130, v3095)
	mBase = m.M
	v3097 = m.ExcPending
	if v3097 != 0 {
		goto L1
	} else {
		goto L615
	}
L615:
	;
	if v2470 == int32(0) {
		goto L616
	} else {
		goto L617
	}
L616:
	;
	F_XLogRegisterBuffer(m, int32(1), v89, int32(8))
	mBase = m.M
	v3103 = m.ExcPending
	if v3103 != 0 {
		goto L1
	} else {
		goto L619
	}
L617:
	;
	goto L618
L618:
	;
	F_XLogRegisterData(m, v39+int32(80), int32(14))
	mBase = m.M
	v3108 = m.ExcPending
	if v3108 != 0 {
		goto L1
	} else {
		goto L620
	}
L619:
	;
	goto L618
L620:
	;
	if (v2978|v2979)&int32(_a_F_heap_update_4) == int32(0) {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	v3147 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+16))
	v3148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3147)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+72)) = uint16(v3148)
	v3150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3147)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+74)) = uint16(v3150)
	v3152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3147)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+76)) = uint8(v3152)
	F_XLogRegisterBufData(m, int32(0), v39+int32(72), int32(5))
	mBase = m.M
	v3159 = m.ExcPending
	if v3159 != 0 {
		goto L1
	} else {
		goto L632
	}
L622:
	;
	v3114 = int32(_a_F_heap_update_4)
	v3116 = int32(0)
	if base.B2i32(v2978&v3114 == v3116)|base.B2i32(v2979&v3114 == v3116) == v3116 {
		goto L623
	} else {
		goto L624
	}
L623:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+64)) = uint16(v2978)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+62)) = uint16(v2979)
	F_XLogRegisterBufData(m, int32(0), v39+int32(62), int32(4))
	mBase = m.M
	v3132 = m.ExcPending
	if v3132 != 0 {
		goto L1
	} else {
		goto L626
	}
L624:
	;
	goto L625
L625:
	;
	if v2979&int32(_a_F_heap_update_4) != 0 {
		goto L627
	} else {
		goto L628
	}
L626:
	;
	goto L621
L627:
	;
	F_XLogRegisterBufData(m, int32(0), v39+int32(60), int32(2))
	mBase = m.M
	v3140 = m.ExcPending
	if v3140 != 0 {
		goto L1
	} else {
		goto L630
	}
L628:
	;
	goto L629
L629:
	;
	F_XLogRegisterBufData(m, int32(0), v39+int32(58), int32(2))
	mBase = m.M
	v3146 = m.ExcPending
	if v3146 != 0 {
		goto L1
	} else {
		goto L631
	}
L630:
	;
	goto L621
L631:
	;
	goto L621
L632:
	;
	v3160 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+16))
	v3161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)))
	if v3161 == int32(0) {
		goto L634
	} else {
		goto L635
	}
L633:
	;
	v3203 = int32(0)
	if base.B2i32(v2314 == v3203)|(v2618^int32(1)) == v3203 {
		goto L643
	} else {
		goto L644
	}
L634:
	;
	v3165 = int32(23)
	v3167 = *(*int32)(unsafe.Add(mBase, uint32(v2095)))
	v3168 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)))
	F_XLogRegisterBufData(m, int32(0), v3160+v3165, v3167-v3168-v3165)
	mBase = m.M
	v3173 = m.ExcPending
	if v3173 != 0 {
		goto L1
	} else {
		goto L637
	}
L635:
	;
	goto L636
L636:
	;
	v3174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3160)+22)))
	v3176 = v3174 - int32(23)
	if v3176 != 0 {
		goto L638
	} else {
		goto L639
	}
L637:
	;
	goto L633
L638:
	;
	F_XLogRegisterBufData(m, int32(0), v3160+int32(23), v3176)
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		goto L1
	} else {
		goto L641
	}
L639:
	;
	v3186 = v3161
	v3187 = v3160
	v3188 = int32(23)
	goto L640
L640:
	;
	v3192 = v3188 + v3186&int32(_a_F_heap_update_4)
	v3194 = *(*int32)(unsafe.Add(mBase, uint32(v2095)))
	v3195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)))
	F_XLogRegisterBufData(m, int32(0), v3192+v3187, v3194-(v3192+v3195))
	mBase = m.M
	v3199 = m.ExcPending
	if v3199 != 0 {
		goto L1
	} else {
		goto L642
	}
L641:
	;
	v3182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)))
	v3183 = *(*int32)(unsafe.Add(mBase, uint32(v2095)+16))
	v3184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3183)+22)))
	v3186 = v3182
	v3187 = v3183
	v3188 = v3184
	goto L640
L642:
	;
	goto L633
L643:
	;
	v3210 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+16))
	v3211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3210)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+66)) = uint16(v3211)
	v3213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3210)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+68)) = uint16(v3213)
	v3215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3210)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+70)) = uint8(v3215)
	F_XLogRegisterData(m, v39+int32(66), int32(5))
	mBase = m.M
	v3221 = m.ExcPending
	if v3221 != 0 {
		goto L1
	} else {
		goto L646
	}
L644:
	;
	goto L645
L645:
	;
	v3232 = int32(_a_F_heap_update_33)
	v3234 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_heap_update[13])))
	v3235 = v3234 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_heap_update[13])) = uint8(v3235)
	goto L648
L646:
	;
	v3222 = *(*int32)(unsafe.Add(mBase, uint32(v2314)+16))
	v3223 = int32(23)
	v3225 = *(*int32)(unsafe.Add(mBase, uint32(v2314)))
	F_XLogRegisterData(m, v3222+v3223, v3225-v3223)
	mBase = m.M
	v3229 = m.ExcPending
	if v3229 != 0 {
		goto L1
	} else {
		goto L647
	}
L647:
	;
	goto L645
L648:
	;
	v3240 = F_XLogInsert(m, int32(10), v3053&int32(255))
	mBase = m.M
	v3241 = m.ExcPending
	if v3241 != 0 {
		goto L1
	} else {
		goto L649
	}
L649:
	;
	if v2470 == int32(0) {
		goto L650
	} else {
		goto L651
	}
L650:
	;
	if v2130 < int32(0) {
		goto L654
	} else {
		goto L655
	}
L651:
	;
	goto L652
L652:
	;
	if v89 < int32(0) {
		goto L658
	} else {
		goto L659
	}
L653:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3261))) = base.I64_rotr(v3240, int64(32))
	goto L652
L654:
	;
	v3247 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v3253 = *(*int32)(unsafe.Add(mBase, uint32(v3247+(v2130^int32(-1))<<(uint(int32(2))%32))))
	v3261 = v3253
	goto L653
L655:
	;
	goto L656
L656:
	;
	v3255 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v3261 = v3255 + v2130<<(uint(int32(13))%32) + int32(-8192)
	goto L653
L657:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3282))) = base.I64_rotr(v3240, int64(32))
	goto L514
L658:
	;
	v3268 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[1]))
	v3274 = *(*int32)(unsafe.Add(mBase, uint32(v3268+(v89^int32(-1))<<(uint(int32(2))%32))))
	v3282 = v3274
	goto L657
L659:
	;
	goto L660
L660:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[2]))
	v3282 = v3276 + v89<<(uint(int32(13))%32) + int32(-8192)
	goto L657
L661:
	;
	F_ReleaseBuffer(m, v89)
	mBase = m.M
	v3350 = m.ExcPending
	if v3350 != 0 {
		goto L1
	} else {
		goto L671
	}
L662:
	;
	F_LockBuffer(m, v2130, int32(0))
	mBase = m.M
	v3332 = m.ExcPending
	if v3332 != 0 {
		goto L1
	} else {
		goto L665
	}
L663:
	;
	goto L664
L664:
	;
	F_LockBuffer(m, v89, int32(0))
	mBase = m.M
	v3344 = m.ExcPending
	if v3344 != 0 {
		goto L1
	} else {
		goto L669
	}
L665:
	;
	F_LockBuffer(m, v89, int32(0))
	mBase = m.M
	v3335 = m.ExcPending
	if v3335 != 0 {
		goto L1
	} else {
		goto L666
	}
L666:
	;
	F_CacheInvalidateHeapTuple(m, l0, v39+int32(32), v2095)
	mBase = m.M
	v3339 = m.ExcPending
	if v3339 != 0 {
		goto L1
	} else {
		goto L667
	}
L667:
	;
	F_ReleaseBuffer(m, v2130)
	mBase = m.M
	v3341 = m.ExcPending
	if v3341 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	goto L661
L669:
	;
	F_CacheInvalidateHeapTuple(m, l0, v39+int32(32), v2095)
	mBase = m.M
	v3348 = m.ExcPending
	if v3348 != 0 {
		goto L1
	} else {
		goto L670
	}
L670:
	;
	goto L661
L671:
	;
	v3351 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if v3351 != 0 {
		goto L672
	} else {
		goto L673
	}
L672:
	;
	F_ReleaseBuffer(m, v3351)
	mBase = m.M
	v3353 = m.ExcPending
	if v3353 != 0 {
		goto L1
	} else {
		goto L675
	}
L673:
	;
	goto L674
L674:
	;
	v3354 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v3354 != 0 {
		goto L676
	} else {
		goto L677
	}
L675:
	;
	goto L674
L676:
	;
	F_ReleaseBuffer(m, v3354)
	mBase = m.M
	v3356 = m.ExcPending
	if v3356 != 0 {
		goto L1
	} else {
		goto L679
	}
L677:
	;
	goto L678
L678:
	;
	if v973&int32(1) != 0 {
		goto L680
	} else {
		goto L681
	}
L679:
	;
	goto L678
L680:
	;
	v3359 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v3362 = *(*int32)(unsafe.Add(mBase, uint32(v3359*int32(12))+uint32(_c_F_heap_update[3])))
	F_UnlockTuple(m, l0, v496, v3362)
	mBase = m.M
	v3364 = m.ExcPending
	if v3364 != 0 {
		goto L1
	} else {
		goto L683
	}
L681:
	;
	goto L682
L682:
	;
	v3365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v3365 == int32(0) {
		goto L685
	} else {
		goto L686
	}
L683:
	;
	goto L682
L684:
	;
	if v2095 != l2 {
		goto L702
	} else {
		goto L703
	}
L685:
	;
	v3368 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v3368 != int32(1) {
		goto L684
	} else {
		goto L688
	}
L686:
	;
	v3374 = v3365
	goto L687
L687:
	;
	v3376 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[0]))
	v3377 = *(*int32)(unsafe.Add(mBase, uint32(v3376)+28))
	goto L690
L688:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v3372 = m.ExcPending
	if v3372 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v3374 = v3373
	goto L687
L690:
	;
	v3378 = *(*int32)(unsafe.Add(mBase, uint32(v3374)+8))
	if v3378 != 0 {
		goto L692
	} else {
		goto L693
	}
L691:
	;
	v3399 = *(*int64)(unsafe.Add(mBase, uint32(v3396)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3396)+8)) = v3399 + int64(1)
	if v2260|v2152 == int32(0) {
		goto L684
	} else {
		goto L698
	}
L692:
	;
	v3379 = *(*int32)(unsafe.Add(mBase, uint32(v3378)+56))
	if v3379 == v3377 {
		v3396 = v3378
		goto L691
	} else {
		goto L695
	}
L693:
	;
	goto L694
L694:
	;
	v3381 = F_pgstat_get_xact_stack_level(m, v3377)
	mBase = m.M
	v3382 = m.ExcPending
	if v3382 != 0 {
		goto L1
	} else {
		goto L696
	}
L695:
	;
	goto L694
L696:
	;
	v3384 = *(*int32)(unsafe.Add(mBase, _c_F_heap_update[14]))
	v3386 = F_MemoryContextAllocZero(m, v3384, int32(72))
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L1
	} else {
		goto L697
	}
L697:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3386)+56)) = v3377
	v3389 = *(*int32)(unsafe.Add(mBase, uint32(v3374)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3386)+64)) = v3374
	*(*int32)(unsafe.Add(mBase, uint32(v3386)+60)) = v3389
	v3392 = *(*int32)(unsafe.Add(mBase, uint32(v3381)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3386)+68)) = v3392
	*(*int32)(unsafe.Add(mBase, uint32(v3381)+20)) = v3386
	*(*int32)(unsafe.Add(mBase, uint32(v3374)+8)) = v3386
	v3396 = v3386
	goto L691
L698:
	;
	if v2260 != 0 {
		goto L699
	} else {
		goto L700
	}
L699:
	;
	v3408 = int32(64)
	goto L701
L700:
	;
	v3408 = int32(72)
	goto L701
L701:
	;
	v3409 = v3374 + v3408
	v3410 = *(*int64)(unsafe.Add(mBase, uint32(v3409)))
	*(*int64)(unsafe.Add(mBase, uint32(v3409))) = v3410 + int64(1)
	goto L684
L702:
	;
	v3420 = v2095 + int32(4)
	v3421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3420)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)) = uint16(v3421)
	v3423 = *(*int32)(unsafe.Add(mBase, uint32(v3420)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v3423
	F_pfree(m, v2095)
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L1
	} else {
		goto L705
	}
L703:
	;
	goto L704
L704:
	;
	if v2260 != 0 {
		goto L706
	} else {
		goto L707
	}
L705:
	;
	goto L704
L706:
	;
	v3429 = v2259
	goto L708
L707:
	;
	v3429 = int32(1)
	goto L708
L708:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v3429
	if v2314 == int32(0) {
		goto L709
	} else {
		goto L710
	}
L709:
	;
	F_bms_free(m, v64)
	mBase = m.M
	v3441 = m.ExcPending
	if v3441 != 0 {
		goto L1
	} else {
		goto L713
	}
L710:
	;
	v3433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+31)))
	if v3433&int32(1) == int32(0) {
		goto L709
	} else {
		goto L711
	}
L711:
	;
	F_pfree(m, v2314)
	mBase = m.M
	v3439 = m.ExcPending
	if v3439 != 0 {
		goto L1
	} else {
		goto L712
	}
L712:
	;
	goto L709
L713:
	;
	v3444 = int32(0)
	goto L4
L714:
	;
	F_bms_free(m, v3482)
	mBase = m.M
	v3518 = m.ExcPending
	if v3518 != 0 {
		goto L1
	} else {
		goto L715
	}
L715:
	;
	F_bms_free(m, v3507)
	mBase = m.M
	v3520 = m.ExcPending
	if v3520 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	F_bms_free(m, v3504)
	mBase = m.M
	v3522 = m.ExcPending
	if v3522 != 0 {
		goto L1
	} else {
		goto L717
	}
L717:
	;
	F_bms_free(m, v82)
	mBase = m.M
	v3524 = m.ExcPending
	if v3524 != 0 {
		goto L1
	} else {
		goto L718
	}
L718:
	;
	m.G0 = v39 + int32(96)
	return v3480
}
func F_rewrite_heap_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
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
	var v70 int32
	_ = v70
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
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
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
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
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
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
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v304 int32
	_ = v304
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v315 int64
	_ = v315
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int64
	_ = v324
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
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
	var v406 int32
	_ = v406
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v441 int32
	_ = v441
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = int32(_a_F_rewrite_heap_tuple_0)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_rewrite_heap_tuple[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, _c_F_rewrite_heap_tuple[0])) = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v27
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
	*(*int64)(unsafe.Add(mBase, uint32(v25))) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31)+20)))
	v34 = v32 & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v31)+20)) = uint16(v34)
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v36)+18)))
	v39 = v37 & int32(_a_F_rewrite_heap_tuple_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v36)+18)) = uint16(v39)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v42 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+20)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+20)))
	v47 = v42 | v44&int32(_a_F_rewrite_heap_tuple_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+20)) = uint16(v47)
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+136))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v51)+140))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v56 = m.G0
	v58 = v56 + int32(-64)
	m.G0 = v58
	*(*int32)(unsafe.Add(mBase, uint32(v58)+44)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v58)+40)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v58)+36)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v58)+32)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v58)+28)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v58)+24)) = v52
	*(*int32)(unsafe.Add(mBase, uint32(v58)+20)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v58)+16)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v58)+12)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v58)+8)) = v54
	v70 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v58)+4)) = uint8(v70)
	v80 = F_heap_prepare_freeze_tuple(m, v49, v56+int32(-40), v56+int32(-60), v56+int32(-12), v56+int32(-13))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v80 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v58)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v82
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v58)+60)))
	if v84&int32(2) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v58 - int32(-64)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v102 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v101)+16)) = uint16(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v101)+12)) = int32(-1)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+21)))
	if v107&int32(8) != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(2)
	goto L8
L7:
	;
	goto L8
L8:
	;
	if v84&int32(4) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v49)+8)) = int32(0)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+58)))
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+20)) = uint16(v93)
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+18)) = uint16(v95)
	goto L5
L12:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_rewrite_heap_tuple[0])) = v21
	m.G0 = v18 + int32(80)
	return
L13:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)) = uint16(v202)
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v204
	v209 = v18 + int32(8) | int32(4)
	v211 = v18 + int32(62)
	v213 = v18 + int32(44)
	v215 = v18 + int32(56)
	v219 = l2
	v220 = int32(0)
	goto L41
L14:
	;
	v110 = F_HeapTupleHeaderIsOnlyLocked(m, v106)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v110 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+16)))
	if v113 == int32(_a_F_rewrite_heap_tuple_3) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+12)))
	v117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+14)))
	if v116&v117 == int32(_a_F_rewrite_heap_tuple_4) {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v122 = l1 + int32(4)
	v124 = v112 + int32(12)
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+2)))
	v126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	v127 = int32(16)
	v130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+2)))
	v131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v125|v126<<(uint(v127)%32) == v130|v131<<(uint(v127)%32) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	if v141 != 0 {
		goto L13
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)))
	if v137 == v138 {
		v141 = int32(1)
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v141 = int32(0)
	goto L22
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = int64(0)
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v146)+20)))
	if v147&int32(_a_F_rewrite_heap_tuple_5) == int32(_a_F_rewrite_heap_tuple_6) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v157
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v156)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v159
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)) = uint16(v161)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v165 = v18 + int32(8)
	v166 = int32(0)
	v168 = F_hash_search(m, v163, v165, v166, v166)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L33
	}
L29:
	;
	v152 = F_HeapTupleGetUpdateXid(m, v146)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	v156 = v146
	v157 = v155
	goto L28
L32:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v156 = v154
	v157 = v152
	goto L28
L33:
	;
	if v168 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v176 = F_hash_search(m, v172, v165, int32(1), v18+int32(7))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v168)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v185)+16)) = uint16(v186)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v168)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v185)+12)) = v188
	v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v196 = F_hash_search(m, v190, v18+int32(8), int32(2), v18+int32(7))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v176)+16)) = uint16(v178)
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v122)))
	*(*int32)(unsafe.Add(mBase, uint32(v176)+12)) = v180
	v182 = F_heap_copytuple(m, l2)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v176)+20)) = v182
	goto L12
L39:
	;
	goto L13
L40:
	;
	if v220&int32(1) == int32(0) {
		goto L12
	} else {
		goto L97
	}
L41:
	;
	F_raw_heap_insert(m, l0, v219)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v428 = F_hash_search(m, v422, v18+int32(8), int32(1), v18+int32(7))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L1
	} else {
		goto L96
	}
L43:
	;
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)) = uint16(v234)
	v236 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v236
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+76)) = uint16(v238)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v219)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = v240
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v242 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v348)+20)))
	if v349&int32(_a_F_rewrite_heap_tuple_7) == int32(0) {
		goto L40
	} else {
		goto L75
	}
L45:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v246)+20)))
	v248 = int32(768)
	if v247&v248 != v248 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v246)))
	v253 = v252
	goto L48
L47:
	;
	v253 = int32(2)
	goto L48
L48:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v247&int32(_a_F_rewrite_heap_tuple_5) == int32(_a_F_rewrite_heap_tuple_6) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v263 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v253) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v259 = F_HeapTupleGetUpdateXid(m, v246)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	v262 = v261
	goto L49
L53:
	;
	v262 = v259
	goto L49
L54:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v254))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v253)) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v281 = v263
	goto L56
L56:
	;
	if base.Ui32(v262) < base.Ui32(int32(3)) {
		v307 = v263
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v281 = v278 ^ int32(1)
	goto L56
L58:
	;
	v278 = base.B2i32(base.Ui32(v253) < base.Ui32(v254))
	goto L57
L59:
	;
	goto L60
L60:
	;
	v278 = int32(base.Ui32(v253-v254) >> (uint(int32(31)) % 32))
	goto L57
L61:
	;
	if v307|v281 != int32(1) {
		goto L44
	} else {
		goto L68
	}
L62:
	;
	v284 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+20)))
	if v285&int32(128)|base.B2i32(v285&int32(_a_F_rewrite_heap_tuple_8) == int32(64)) != 0 {
		v307 = v263
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v254))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v262)) == int32(0) {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v307 = v304 ^ int32(1)
	goto L61
L65:
	;
	v304 = base.B2i32(base.Ui32(v262) < base.Ui32(v254))
	goto L64
L66:
	;
	goto L67
L67:
	;
	v304 = int32(base.Ui32(v262-v254) >> (uint(int32(31)) % 32))
	goto L64
L68:
	;
	v312 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v312)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v313
	v315 = *(*int64)(unsafe.Add(mBase, uint32(v312)))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v315
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v215)+4)) = uint16(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v215))) = v319
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v321)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v213)+8)) = v322
	v324 = *(*int64)(unsafe.Add(mBase, uint32(v321)))
	*(*int64)(unsafe.Add(mBase, uint32(v213))) = v324
	v326 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+76)))
	*(*uint16)(unsafe.Add(mBase, uint32(v211)+4)) = uint16(v326)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v211))) = v328
	if v281 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	F_logical_rewrite_log_mapping(m, l0, v253, v18+int32(32))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	if v307&base.B2i32(v253 != v262) == int32(0) {
		goto L44
	} else {
		goto L73
	}
L72:
	;
	goto L71
L73:
	;
	F_logical_rewrite_log_mapping(m, l0, v262, v18+int32(32))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L74
	}
L74:
	;
	goto L44
L75:
	;
	v354 = int32(768)
	if v349&v354 != v354 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v360 = v358
	goto L78
L77:
	;
	v360 = int32(2)
	goto L78
L78:
	;
	v361 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v361))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v360)) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	if v373 != 0 {
		goto L40
	} else {
		goto L83
	}
L80:
	;
	v373 = base.B2i32(base.Ui32(v360) < base.Ui32(v361))
	goto L79
L81:
	;
	goto L82
L82:
	;
	v373 = int32(base.Ui32(v360-v361) >> (uint(int32(31)) % 32))
	goto L79
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = int64(0)
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v219)+16))
	v379 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v378)+20)))
	v380 = int32(768)
	if v379&v380 != v380 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v378)))
	v386 = v384
	goto L86
L85:
	;
	v386 = int32(2)
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v386
	v388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v209)+4)) = uint16(v388)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v209))) = v390
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v395 = int32(0)
	v397 = F_hash_search(m, v392, v18+int32(8), v395, v395)
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	if v397 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v220&int32(1) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	goto L42
L91:
	;
	F_pfree(m, v219)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L1
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v397)+20))
	v404 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v397)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+28)) = uint16(v404)
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v397)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v406
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v403)+16))
	v409 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v408)+16)) = uint16(v409)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v408)+12)) = v411
	v413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v419 = F_hash_search(m, v413, v18+int32(8), int32(2), v18+int32(7))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L95
	}
L94:
	;
	goto L93
L95:
	;
	v219 = v403
	v220 = int32(1)
	goto L41
L96:
	;
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v18)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v428)+16)) = uint16(v430)
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	*(*int32)(unsafe.Add(mBase, uint32(v428)+12)) = v432
	goto L40
L97:
	;
	F_pfree(m, v219)
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L1
	} else {
		goto L98
	}
L98:
	;
	goto L12
}
