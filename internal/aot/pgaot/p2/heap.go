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
		v8 = *(*int32)(unsafe.Add(mBase, _consts[66]))
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
		v8 = *(*int32)(unsafe.Add(mBase, _consts[66]))
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
		v23 = int32(1)
		v24 = l1 - v23
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
		v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+20)))
		if v26&v23 == v21 {
			v35 = l2 + v24<<(uint(int32(4))%32) + int32(20)
			v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)))
			if v36 < int32(0) {
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
				v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+22)))
				v41 = v25 + v39 + v36
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+6)))
				if v42 != int32(1) {
					v82 = v41
					m.G0 = v10 + int32(16)
					return v82
				} else {
					v45 = int32(*(*int16)(unsafe.Add(mBase, uint32(v35)+4)))
					switch v45&int32(65535) - int32(1) {
					case 0:
						v50 = int32(*(*int8)(unsafe.Add(mBase, uint32(v41))))
						v82 = v50
						m.G0 = v10 + int32(16)
						return v82
					case 1:
						v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41))))
						v82 = v51
						m.G0 = v10 + int32(16)
						return v82
					default:
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v45
							F_errmsg_internal(m, int32(477953), v10)
							mBase = m.M
							v60 = m.ExcPending
							if v60 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(323177), int32(70), int32(67251))
								mBase = m.M
								v65 = m.ExcPending
								if v65 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					case 3:
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
						v82 = v52
						m.G0 = v10 + int32(16)
						return v82
					}
				}
			}
		} else {
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+23)))
			if int32(base.Ui32(v66)>>(uint(v24)%32))&int32(1) != 0 {
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v127 int32
	_ = v127
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
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
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v342 int32
	_ = v342
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v512 int32
	_ = v512
	var v517 int32
	_ = v517
	var v526 int32
	_ = v526
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v586 int32
	_ = v586
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v631 int32
	_ = v631
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v744 int32
	_ = v744
	var v769 int32
	_ = v769
	var v771 int32
	_ = v771
	var v775 int32
	_ = v775
	var v777 int32
	_ = v777
	var v780 int32
	_ = v780
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v811 int32
	_ = v811
	var v812 int64
	_ = v812
	var v813 int32
	_ = v813
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v881 int32
	_ = v881
	var v915 int32
	_ = v915
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v936 int32
	_ = v936
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v1006 int32
	_ = v1006
	var v1008 int32
	_ = v1008
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1077 int32
	_ = v1077
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1149 int32
	_ = v1149
	v7 = int32(0)
	v33 = m.G0
	v35 = v33 - int32(8208)
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
	v41 = int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v43 < int32(2) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+118)))
	if v87 != int32(112) {
		v100 = v7
		goto L21
	} else {
		goto L22
	}
L4:
	;
	v84 = v41
	v85 = v7
	goto L3
L5:
	;
	goto L6
L6:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+118)))
	if v47 != int32(112) {
		v62 = v7
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+118)))
	if v64 != int32(112) {
		goto L12
	} else {
		goto L13
	}
L8:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+119)))
	if v50 == int32(102) {
		v62 = v7
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L10
L10:
	;
	v57 = base.B2i32(base.Ui32(v53) < base.Ui32(int32(12000))) ^ int32(1)
	v59 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v59 < int32(2) {
		v84 = v41
		v85 = v57
		goto L3
	} else {
		goto L11
	}
L11:
	;
	v62 = v57
	goto L7
L12:
	;
	v84 = v41
	v85 = v62
	goto L3
L13:
	;
	goto L14
L14:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L15
L15:
	;
	if base.Ui32(v68) < base.Ui32(int32(12000)) {
		v84 = int32(0)
		v85 = v62
		goto L3
	} else {
		goto L16
	}
L16:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v71 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v84 = int32(1)
	v85 = v62
	goto L3
L18:
	;
	goto L19
L19:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v76)+119)))
	switch v77 - int32(109) {
	case 0, 5:
		goto L20
	default:
		v84 = int32(1)
		v85 = v62
		goto L3
	}
L20:
	;
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+104)))
	v84 = base.B2i32(v80 == int32(0))
	v85 = v62
	goto L3
L21:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v101 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v92 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if int32(0) < v92 {
		v100 = int32(1)
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v96 != 0 {
		v100 = int32(0)
		goto L21
	} else {
		goto L24
	}
L24:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v100 = base.B2i32(v97 == int32(0))
	goto L21
L25:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v101)+4))
	v108 = base.I32_div_s(int32(819200)-v103<<(uint(int32(13))%32), int32(100))
	v109 = v108
	goto L27
L26:
	;
	v109 = v7
	goto L27
L27:
	;
	v112 = F_palloc(m, l2<<(uint(int32(2))%32))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if l2 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v915 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	if v915 != 0 {
		goto L169
	} else {
		goto L170
	}
L30:
	;
	F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v127 = int32(0)
	goto L34
L33:
	;
	goto L29
L34:
	;
	v154 = v127 << (uint(int32(2)) % 32)
	v155 = l1 + v154
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v159 = F_ExecFetchSlotHeapTuple(m, v156, int32(1), int32(0))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	F_CheckForSerializableConflictIn(m, l0, int32(0), int32(-1))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L39
	}
L36:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v161)+36)) = v162
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v159)+12)) = v165
	v168 = F_heap_prepare_insert(m, l0, v159, v37, l3, l4)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v154+v112))) = v168
	v172 = v127 + int32(1)
	if v172 != l2 {
		v127 = v172
		goto L34
	} else {
		goto L38
	}
L38:
	;
	goto L35
L39:
	;
	v178 = int32(4)
	v181 = int32(8168) - v109
	v182 = int32(0)
	v184 = v84 | base.B2i32(v100 == v182)
	v186 = l4 & v178
	v192 = v35 + int32(16) | v178
	v206 = v182
	v208 = v7
	v209 = v7
	v216 = v7
	goto L40
L40:
	;
	v227 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v227 != 0 {
		goto L42
	} else {
		goto L43
	}
L41:
	;
	goto L29
L42:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v230 = int32(0)
	if v209&base.B2i32(v208 != v230) == v230 {
		goto L47
	} else {
		goto L48
	}
L45:
	;
	goto L44
L46:
	;
	v394 = v112 + v208<<(uint(int32(2))%32)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v395)))
	v397 = int32(0)
	v402 = F_RelationGetBufferForTuple(m, l0, v396, v397, l4, l5, v35+int32(12), v397, v372-v382)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L1
	} else {
		goto L64
	}
L47:
	;
	v235 = int32(1)
	v236 = l2 - v208
	if v208+v235 != l2 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	goto L49
L49:
	;
	v372 = v206
	v382 = v216 + int32(1)
	goto L46
L50:
	;
	v251 = v181
	v252 = v208
	v254 = int32(0)
	v257 = v235
	goto L53
L51:
	;
	v316 = v181
	v317 = v208
	v322 = v235
	goto L52
L52:
	;
	v342 = int32(0)
	if v236&v235 == v342 {
		v372 = v322
		v382 = v342
		goto L46
	} else {
		goto L62
	}
L53:
	;
	v278 = v252 << (uint(int32(2)) % 32)
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v112+v278)))
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v280)))
	v287 = (v281+int32(7))&int32(-8) | int32(4)
	v288 = base.B2i32(base.Ui32(v251) < base.Ui32(v287))
	if base.Ui32(v251) < base.Ui32(v287) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v316 = v302
	v317 = v306
	v322 = v304
	goto L52
L55:
	;
	v289 = v181
	goto L57
L56:
	;
	v289 = v251
	goto L57
L57:
	;
	v290 = v289 - v287
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v278+(v112+v178))))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v292)))
	v299 = (v293+int32(7))&int32(-8) | int32(4)
	v300 = base.B2i32(base.Ui32(v290) < base.Ui32(v299))
	if base.Ui32(v290) < base.Ui32(v299) {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	v301 = v181
	goto L60
L59:
	;
	v301 = v290
	goto L60
L60:
	;
	v302 = v301 - v299
	v304 = v288 + v257 + v300
	v305 = int32(2)
	v306 = v252 + v305
	v308 = v254 + v305
	if v308 != v236&int32(-2) {
		v251 = v302
		v252 = v306
		v254 = v308
		v257 = v304
		goto L53
	} else {
		goto L61
	}
L61:
	;
	goto L54
L62:
	;
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v112+v317<<(uint(int32(2))%32))))
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	v372 = v322 + base.B2i32(base.Ui32(v316) < base.Ui32((v349+int32(7))&int32(-8)|int32(4)))
	v382 = v342
	goto L46
L63:
	;
	v422 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421)+12)))
	v423 = int32(4470804)
	v425 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v425 + int32(1)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	F_RelationPutHeapTuple(m, v402, v429, int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L1
	} else {
		goto L68
	}
L64:
	;
	if v402 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v407 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v407+(v402^int32(-1))<<(uint(int32(2))%32))))
	v421 = v413
	goto L63
L66:
	;
	goto L67
L67:
	;
	v415 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v421 = v415 + v402<<(uint(int32(13))%32) + int32(-8192)
	goto L63
L68:
	;
	v434 = v422 & int32(65535)
	v441 = int32(0)
	if v184 == v441 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v394)))
	F_log_heap_new_cid(m, l0, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L1
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v448 = base.B2i32((v434+int32(262120))&int32(262140) == v441) | base.B2i32(base.Ui32(v434) < base.Ui32(int32(25)))
	v449 = int32(1)
	v451 = v208 + v449
	if l2 <= v451 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L71
L73:
	;
	v610 = v448 & int32(base.Ui32(v186)>>(uint(int32(2))%32))
	v613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421)+10)))
	v618 = base.B2i32(v186 == int32(0)) & int32(base.Ui32(v613&int32(4))>>(uint(int32(2))%32))
	if v618 != 0 {
		goto L106
	} else {
		goto L107
	}
L74:
	;
	v586 = v449
	v609 = v451
	goto L73
L75:
	;
	goto L76
L76:
	;
	v453 = l2 - v208
	v461 = v451
	v463 = v449
	goto L77
L77:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v112+v461<<(uint(int32(2))%32))))
	v493 = int32(4)
	v494 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421)+14)))
	v495 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v421)+12)))
	v496 = v494 - v495
	if v496 <= v493 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v586 = v453
	v609 = l2
	goto L73
L79:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v489)))
	if base.Ui32(v558) < base.Ui32((v559+int32(7))&int32(-8)+v109) {
		v586 = v463
		v609 = v461
		goto L73
	} else {
		goto L98
	}
L80:
	;
	v499 = v493
	goto L82
L81:
	;
	v499 = v496
	goto L82
L82:
	;
	v501 = v499 - int32(4)
	if v501 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v558 = int32(0)
	goto L79
L84:
	;
	goto L85
L85:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v495) {
		goto L87
	} else {
		goto L88
	}
L86:
	;
	v558 = v501
	goto L79
L87:
	;
	v512 = int32(base.Ui32(v495+int32(262120)) >> (uint(int32(2)) % 32))
	goto L89
L88:
	;
	v512 = int32(0)
	goto L89
L89:
	;
	if base.Ui32(v512&int32(65535)) < base.Ui32(int32(291)) {
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v421)+10)))
	if v517&int32(1) == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v558 = int32(0)
	goto L79
L92:
	;
	goto L93
L93:
	;
	v526 = int32(1)
	goto L94
L94:
	;
	v537 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v526&int32(65535)<<(uint(int32(2))%32)+(v421+int32(24))-int32(3)))))
	if v537&int32(384) == int32(0) {
		goto L86
	} else {
		goto L96
	}
L95:
	;
	v558 = int32(0)
	goto L79
L96:
	;
	v543 = v526 + int32(1)
	v544 = int32(65535)
	if base.Ui32(v543&v544) <= base.Ui32(v512&v544) {
		v526 = v543
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	F_RelationPutHeapTuple(m, v402, v489, int32(0))
	mBase = m.M
	v568 = m.ExcPending
	if v568 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	if v184 == int32(0) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	F_log_heap_new_cid(m, l0, v489)
	mBase = m.M
	v572 = m.ExcPending
	if v572 != 0 {
		goto L1
	} else {
		goto L103
	}
L101:
	;
	goto L102
L102:
	;
	v574 = v463 + int32(1)
	if v574 != v453 {
		v461 = v574 + v208
		v463 = v574
		goto L77
	} else {
		goto L104
	}
L103:
	;
	goto L102
L104:
	;
	goto L78
L105:
	;
	F_MarkBufferDirty(m, v402)
	mBase = m.M
	v653 = m.ExcPending
	if v653 != 0 {
		goto L1
	} else {
		goto L115
	}
L106:
	;
	v620 = v613 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v421)+10)) = uint16(v620)
	if v402 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	goto L108
L108:
	;
	if v610 == int32(0) {
		goto L105
	} else {
		goto L114
	}
L109:
	;
	v641 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v643 = F_visibilitymap_clear(m, v640, v641, int32(3))
	mBase = m.M
	v644 = m.ExcPending
	if v644 != 0 {
		goto L1
	} else {
		goto L113
	}
L110:
	;
	v625 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v625+(v402^int32(-1))<<(uint(int32(6))%32))+16))
	v640 = v631
	goto L109
L111:
	;
	goto L112
L112:
	;
	v633 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v633+v402<<(uint(int32(6))%32)+int32(-64))+16))
	v640 = v639
	goto L109
L113:
	;
	goto L105
L114:
	;
	v648 = v613 | int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v421)+10)) = uint16(v648)
	goto L105
L115:
	;
	if v100 != 0 {
		goto L116
	} else {
		goto L117
	}
L116:
	;
	if v610 != 0 {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	v849 = int32(4470804)
	v851 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v851 - int32(1)
	if v610 != 0 {
		goto L159
	} else {
		goto L160
	}
L119:
	;
	v655 = int32(32)
	goto L121
L120:
	;
	v655 = v618
	goto L121
L121:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)) = uint8(v655)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+18)) = uint16(v586)
	v658 = int32(0)
	if v448 != 0 {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v662 = v658
	goto L124
L123:
	;
	v662 = v586 << (uint(int32(1)) % 32)
	goto L124
L124:
	;
	v663 = v192 + v662
	if int32(0) < v586 {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v673 = v663
	v676 = v658
	goto L128
L126:
	;
	v744 = v663
	goto L127
L127:
	;
	if v85 != 0 {
		goto L138
	} else {
		goto L139
	}
L128:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v394+v676<<(uint(int32(2))%32))))
	if v448 == int32(0) {
		goto L130
	} else {
		goto L131
	}
L129:
	;
	v744 = v733
	goto L127
L130:
	;
	v707 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v701)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v192+v676<<(uint(int32(1))%32)))) = uint16(v707)
	goto L132
L131:
	;
	goto L132
L132:
	;
	v712 = (v673 + int32(1)) & int32(-2)
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v701)+16))
	v714 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v713)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v712)+2)) = uint16(v714)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v701)+16))
	v717 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v716)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v712)+4)) = uint16(v717)
	v719 = *(*int32)(unsafe.Add(mBase, uint32(v701)+16))
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v719)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v712)+6)) = uint8(v720)
	v723 = v712 + int32(7)
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v701)+16))
	v725 = int32(23)
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	v729 = v727 - v725
	if v729 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v712))) = uint16(v729)
	v733 = v731 + v729
	v735 = v676 + int32(1)
	if v735 != v586 {
		v673 = v733
		v676 = v735
		goto L128
	} else {
		goto L137
	}
L134:
	;
	v730 = F__emscripten_memcpy_bulkmem(m, v723, v724+v725, v729)
	mBase = m.M
	v731 = v730
	goto L136
L135:
	;
	v731 = v723
	goto L136
L136:
	;
	goto L133
L137:
	;
	goto L129
L138:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)))
	v771 = v769 | int32(8)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)) = uint8(v771)
	goto L140
L139:
	;
	goto L140
L140:
	;
	if l2 == v609 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)))
	v777 = v775 | int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+16)) = uint8(v777)
	goto L143
L142:
	;
	goto L143
L143:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L1
	} else {
		goto L144
	}
L144:
	;
	v782 = v35 + int32(16)
	F_XLogRegisterData(m, v782, v663-v782)
	mBase = m.M
	v787 = m.ExcPending
	if v787 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v788 = int32(0)
	if v448 != 0 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v791 = int32(6)
	goto L148
L147:
	;
	v791 = v788
	goto L148
L148:
	;
	if v85 != 0 {
		goto L149
	} else {
		goto L150
	}
L149:
	;
	v794 = v791 | int32(16)
	goto L151
L150:
	;
	v794 = v791
	goto L151
L151:
	;
	F_XLogRegisterBuffer(m, v788, v402, v794|int32(8))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L1
	} else {
		goto L152
	}
L152:
	;
	F_XLogRegisterBufData(m, int32(0), v663, v744-v663)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v803 = int32(4371988)
	v805 = int32(*(*uint8)(unsafe.Add(mBase, _consts[60])))
	v806 = v805 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[60])) = uint8(v806)
	goto L154
L154:
	;
	if v448 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v811 = int32(208)
	goto L157
L156:
	;
	v811 = int32(80)
	goto L157
L157:
	;
	v812 = F_XLogInsert(m, int32(9), v811)
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L158
	}
L158:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v421))) = base.I64_rotr(v812, int64(32))
	goto L118
L159:
	;
	if v402 < int32(0) {
		goto L163
	} else {
		goto L164
	}
L160:
	;
	goto L161
L161:
	;
	F_UnlockReleaseBuffer(m, v402)
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L1
	} else {
		goto L167
	}
L162:
	;
	v875 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v878 = F_visibilitymap_set(m, l0, v873, v402, int64(0), v875, int32(0), int32(3))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L1
	} else {
		goto L166
	}
L163:
	;
	v858 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v858+(v402^int32(-1))<<(uint(int32(6))%32))+16))
	v873 = v864
	goto L162
L164:
	;
	goto L165
L165:
	;
	v866 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v866+v402<<(uint(int32(6))%32)+int32(-64))+16))
	v873 = v872
	goto L162
L166:
	;
	goto L161
L167:
	;
	if v609 < l2 {
		v206 = v372
		v208 = v609
		v209 = v448
		v216 = v382
		goto L40
	} else {
		goto L168
	}
L168:
	;
	goto L41
L169:
	;
	F_ReleaseBuffer(m, v915)
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L172
	}
L170:
	;
	goto L171
L171:
	;
	v918 = int32(0)
	F_CheckForSerializableConflictIn(m, l0, v918, int32(-1))
	mBase = m.M
	v922 = m.ExcPending
	if v922 != 0 {
		goto L1
	} else {
		goto L173
	}
L172:
	;
	goto L171
L173:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L175
L174:
	;
	if l2 <= int32(0) {
		goto L182
	} else {
		goto L183
	}
L175:
	;
	if base.B2i32(base.Ui32(v923) < base.Ui32(int32(12000))) == int32(0) {
		goto L174
	} else {
		goto L176
	}
L176:
	;
	if l2 <= int32(0) {
		goto L174
	} else {
		goto L177
	}
L177:
	;
	v936 = v918
	goto L178
L178:
	;
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v112+v936<<(uint(int32(2))%32))))
	F_CacheInvalidateHeapTuple(m, l0, v965, int32(0))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L1
	} else {
		goto L180
	}
L179:
	;
	goto L174
L180:
	;
	v970 = v936 + int32(1)
	if v970 != l2 {
		v936 = v970
		goto L178
	} else {
		goto L181
	}
L181:
	;
	goto L179
L182:
	;
	F_pgstat_count_heap_insert(m, l0, base.I64_extend_i32_s(l2))
	mBase = m.M
	v1149 = m.ExcPending
	if v1149 != 0 {
		goto L1
	} else {
		goto L191
	}
L183:
	;
	v1006 = int32(1)
	v1008 = int32(0)
	if l2 != v1006 {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v1020 = v1008
	v1023 = int32(0)
	goto L187
L185:
	;
	v1077 = v1008
	goto L186
L186:
	;
	if l2&v1006 == int32(0) {
		goto L182
	} else {
		goto L190
	}
L187:
	;
	v1046 = int32(2)
	v1047 = v1020 << (uint(v1046) % 32)
	v1049 = *(*int32)(unsafe.Add(mBase, uint32(l1+v1047)))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1047+v112)))
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v1051)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1049)+28)) = v1052
	v1054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1051)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1049)+32)) = uint16(v1054)
	v1057 = v1047 | int32(4)
	v1059 = *(*int32)(unsafe.Add(mBase, uint32(l1+v1057)))
	v1061 = *(*int32)(unsafe.Add(mBase, uint32(v1057+v112)))
	v1062 = *(*int32)(unsafe.Add(mBase, uint32(v1061)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1059)+28)) = v1062
	v1064 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1061)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1059)+32)) = uint16(v1064)
	v1067 = v1020 + v1046
	v1069 = v1023 + v1046
	if v1069 != l2&int32(2147483646) {
		v1020 = v1067
		v1023 = v1069
		goto L187
	} else {
		goto L189
	}
L188:
	;
	v1077 = v1067
	goto L186
L189:
	;
	goto L188
L190:
	;
	v1106 = v1077 << (uint(int32(2)) % 32)
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(l1+v1106)))
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1106+v112)))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+28)) = v1111
	v1113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1110)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1108)+32)) = uint16(v1113)
	goto L182
L191:
	;
	m.G0 = v35 + int32(8208)
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
	v8 = *(*int32)(unsafe.Add(mBase, _consts[46]))
	if v8 < int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)))
		v14 = v12 & int32(15)
		*(*uint16)(unsafe.Add(mBase, uint32(v11)+20)) = uint16(v14)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+18)))
		v19 = v17 & int32(8191)
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
		v40 = v38 & int32(65503)
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
				F_errmsg(m, int32(218895), int32(0))
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(491798), int32(2283), int32(81189))
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
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	v3 = l2
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = l1 + v3
	v14 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[72]))) = uint8(v14)
	*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[73]))) = uint8(v14)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v3<<(uint(int32(2))%32)+l0)+20))
	v24 = l0 + v21&int32(32767)
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[74]))))
	switch v27 - v14 {
	case 0:
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[75])))
		v31 = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[75]))) = v30 + v31
		v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[76]))))
		if v34 != v31 {
		} else {
			v37 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+20)))
			if v37&int32(256) == int32(0) {
				v42 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[76]))) = uint8(v42)
			} else {
				v45 = int32(768)
				if v37&v45 != v45 {
					v49 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
					v50 = v49
				} else {
					v50 = int32(2)
				}
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v51)+8))
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v52))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == int32(0) {
					v64 = base.B2i32(base.Ui32(v50) < base.Ui32(v52))
				} else {
					v64 = int32(base.Ui32(v50-v52) >> (uint(int32(31)) % 32))
				}
				if v64 == int32(0) {
					v67 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[76]))) = uint8(v67)
				} else {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[77])))
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v69))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v50)) == int32(0) {
						v81 = base.B2i32(base.Ui32(v69) < base.Ui32(v50))
					} else {
						v81 = base.B2i32(int32(0) < v50-v69)
					}
					if base.Ui32(v50) < base.Ui32(int32(3)) {
					} else {
						if v81 == int32(0) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[77]))) = v50
						}
					}
				}
			}
		}
		v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
		if v167 != int32(1) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v174 = l1 + int32(2364)
			v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			v181 = F_heap_prepare_freeze_tuple(m, v24, v170, l1+int32(7616), v174+v175*int32(12), v9+int32(15))
			mBase = m.M
			v182 = m.ExcPending
			if v182 != 0 {
				return
			} else {
				if v181 != 0 {
					v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v183 + int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v174+v183*int32(12))+10)) = uint16(v3)
				} else {
				}
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v192 != 0 {
				} else {
					v193 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[78]))) = uint8(v193)
				}
				m.G0 = v9 + int32(16)
				return
			}
		}
	case 1:
		v87 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[76]))) = uint8(v87)
		v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[79])))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[79]))) = v89 + int32(1)
		v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+20)))
		if v93&int32(6272) == int32(4096) {
			v98 = F_HeapTupleGetUpdateXid(m, v24)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return
			} else {
				v101 = v98
				v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if v102 != 0 {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v102))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v101)) == int32(0) {
						v114 = base.B2i32(base.Ui32(v101) < base.Ui32(v102))
					} else {
						v114 = int32(base.Ui32(v101-v102) >> (uint(int32(31)) % 32))
					}
					if v114 == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v101
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v101
				}
				v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				if v167 != int32(1) {
					m.G0 = v9 + int32(16)
					return
				} else {
					v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v174 = l1 + int32(2364)
					v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					v181 = F_heap_prepare_freeze_tuple(m, v24, v170, l1+int32(7616), v174+v175*int32(12), v9+int32(15))
					mBase = m.M
					v182 = m.ExcPending
					if v182 != 0 {
						return
					} else {
						if v181 != 0 {
							v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v183 + int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v174+v183*int32(12))+10)) = uint16(v3)
						} else {
						}
						v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v192 != 0 {
						} else {
							v193 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[78]))) = uint8(v193)
						}
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		} else {
			v100 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v101 = v100
			v102 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v102 != 0 {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v102))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v101)) == int32(0) {
					v114 = base.B2i32(base.Ui32(v101) < base.Ui32(v102))
				} else {
					v114 = int32(base.Ui32(v101-v102) >> (uint(int32(31)) % 32))
				}
				if v114 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v101
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v101
			}
			v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			if v167 != int32(1) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v174 = l1 + int32(2364)
				v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				v181 = F_heap_prepare_freeze_tuple(m, v24, v170, l1+int32(7616), v174+v175*int32(12), v9+int32(15))
				mBase = m.M
				v182 = m.ExcPending
				if v182 != 0 {
					return
				} else {
					if v181 != 0 {
						v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v183 + int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v174+v183*int32(12))+10)) = uint16(v3)
					} else {
					}
					v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v192 != 0 {
					} else {
						v193 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[78]))) = uint8(v193)
					}
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	case 2:
		v163 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[76]))) = uint8(v163)
		v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
		if v167 != int32(1) {
			m.G0 = v9 + int32(16)
			return
		} else {
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
			v174 = l1 + int32(2364)
			v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			v181 = F_heap_prepare_freeze_tuple(m, v24, v170, l1+int32(7616), v174+v175*int32(12), v9+int32(15))
			mBase = m.M
			v182 = m.ExcPending
			if v182 != 0 {
				return
			} else {
				if v181 != 0 {
					v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v183 + int32(1)
					*(*uint16)(unsafe.Add(mBase, uint32(v174+v183*int32(12))+10)) = uint16(v3)
				} else {
				}
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
				if v192 != 0 {
				} else {
					v193 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[78]))) = uint8(v193)
				}
				m.G0 = v9 + int32(16)
				return
			}
		}
	case 3:
		v118 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[76]))) = uint8(v118)
		v120 = *(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[75])))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[75]))) = v120 + int32(1)
		v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+20)))
		if v124&int32(6272) == int32(4096) {
			v129 = F_HeapTupleGetUpdateXid(m, v24)
			mBase = m.M
			v130 = m.ExcPending
			if v130 != 0 {
				return
			} else {
				v132 = v129
				v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				if v133 != 0 {
					if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v133))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v132)) == int32(0) {
						v145 = base.B2i32(base.Ui32(v132) < base.Ui32(v133))
					} else {
						v145 = int32(base.Ui32(v132-v133) >> (uint(int32(31)) % 32))
					}
					if v145 == int32(0) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v132
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v132
				}
				v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
				if v167 != int32(1) {
					m.G0 = v9 + int32(16)
					return
				} else {
					v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					v174 = l1 + int32(2364)
					v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
					v181 = F_heap_prepare_freeze_tuple(m, v24, v170, l1+int32(7616), v174+v175*int32(12), v9+int32(15))
					mBase = m.M
					v182 = m.ExcPending
					if v182 != 0 {
						return
					} else {
						if v181 != 0 {
							v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v183 + int32(1)
							*(*uint16)(unsafe.Add(mBase, uint32(v174+v183*int32(12))+10)) = uint16(v3)
						} else {
						}
						v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
						if v192 != 0 {
						} else {
							v193 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[78]))) = uint8(v193)
						}
						m.G0 = v9 + int32(16)
						return
					}
				}
			}
		} else {
			v131 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
			v132 = v131
			v133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			if v133 != 0 {
				if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v133))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v132)) == int32(0) {
					v145 = base.B2i32(base.Ui32(v132) < base.Ui32(v133))
				} else {
					v145 = int32(base.Ui32(v132-v133) >> (uint(int32(31)) % 32))
				}
				if v145 == int32(0) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v132
				}
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = v132
			}
			v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+5)))
			if v167 != int32(1) {
				m.G0 = v9 + int32(16)
				return
			} else {
				v170 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
				v174 = l1 + int32(2364)
				v175 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
				v181 = F_heap_prepare_freeze_tuple(m, v24, v170, l1+int32(7616), v174+v175*int32(12), v9+int32(15))
				mBase = m.M
				v182 = m.ExcPending
				if v182 != 0 {
					return
				} else {
					if v181 != 0 {
						v183 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = v183 + int32(1)
						*(*uint16)(unsafe.Add(mBase, uint32(v174+v183*int32(12))+10)) = uint16(v3)
					} else {
					}
					v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+15)))
					if v192 != 0 {
					} else {
						v193 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+uint32(_consts[78]))) = uint8(v193)
					}
					m.G0 = v9 + int32(16)
					return
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v152 = m.ExcPending
		if v152 != 0 {
			return
		} else {
			v153 = int32(*(*int8)(unsafe.Add(mBase, uint32(v11)+uint32(_consts[74]))))
			*(*int32)(unsafe.Add(mBase, uint32(v9))) = v153
			F_errmsg_internal(m, int32(463241), v9)
			mBase = m.M
			v157 = m.ExcPending
			if v157 != 0 {
				return
			} else {
				F_errfinish(m, int32(490606), int32(1474), int32(310536))
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
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
	*(*int64)(unsafe.Add(mBase, uint32(v8))) = v13
	*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v13
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
	if base.Ui32(int32(8192)) <= base.Ui32(v33) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v36 = int32(8192)
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
	F_s_lock(m, v39, int32(491815), int32(455), int32(99252))
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	v8 = m.G0
	v10 = v8 - int32(8000)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	F_heap_deform_tuple(m, l1, v12, v10+int32(1600), v10)
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
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	if int32(0) < v18 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v26 = int32(0)
	goto L6
L4:
	;
	goto L5
L5:
	;
	m.G0 = v10 + int32(8000)
	return
L6:
	;
	v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v17+int32(24)+v26<<(uint(int32(4))%32)))))
	if v36 != int32(65535) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	goto L5
L8:
	;
	v55 = v26 + int32(1)
	if v55 != v18 {
		v26 = v55
		goto L6
	} else {
		goto L14
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26+v10))))
	if v40 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v10+int32(1600)+v26<<(uint(int32(2))%32))))
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44))))
	if v45 != int32(1) {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v44)+1)))
	if v48 != int32(18) {
		goto L8
	} else {
		goto L12
	}
L12:
	;
	F_toast_delete_datum(m, v44, l2)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
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
	var v97 int32
	_ = v97
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
	var v141 int32
	_ = v141
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
	var v185 int32
	_ = v185
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
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v353 int32
	_ = v353
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v455 int32
	_ = v455
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
	v29 = v17
	v30 = int32(0)
	goto L6
L4:
	;
	F_sequence_close(m, v23, int32(1))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L103
	}
L5:
	;
	F_list_free(m, int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L102
	}
L6:
	;
	v37 = int32(0)
	v42 = F_systable_beginscan(m, v23, v37, v37, v37, v37, v37)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L101
	}
L8:
	;
	v44 = F_systable_getnext(m, v42)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v44 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v48 = v44
	v49 = int32(0)
	v52 = v30
	goto L13
L11:
	;
	goto L12
L12:
	;
	goto L7
L13:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+22)))
	v61 = v59 + v60
	v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+72)))
	if v62 != int32(102) {
		v196 = v49
		v198 = v52
		goto L15
	} else {
		goto L16
	}
L14:
	;
	F_systable_endscan(m, v42)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L1
	} else {
		goto L65
	}
L15:
	;
	v199 = F_systable_getnext(m, v42)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L1
	} else {
		goto L63
	}
L16:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v61)+96))
	v66 = int32(0)
	if v29 == v66 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if v104 == int32(0) {
		v196 = v49
		v198 = v52
		goto L15
	} else {
		goto L30
	}
L18:
	;
	v104 = int32(0)
	goto L17
L19:
	;
	goto L20
L20:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	if v72 <= int32(0) {
		v97 = v66
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v104 = v97
	goto L17
L22:
	;
	v75 = int32(0)
	if v75 < v72 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v78 = v72
	goto L25
L24:
	;
	v78 = v75
	goto L25
L25:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v29)+12))
	v81 = int32(0)
	goto L26
L26:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v79+v81<<(uint(int32(2))%32))))
	v90 = base.B2i32(v89 == v65)
	if v89 == v65 {
		v97 = v90
		goto L21
	} else {
		goto L28
	}
L27:
	;
	v97 = v90
	goto L21
L28:
	;
	v92 = v81 + int32(1)
	if v92 != v78 {
		v81 = v92
		goto L26
	} else {
		goto L29
	}
L29:
	;
	goto L27
L30:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v61)+92))
	if v107 == int32(0) {
		v152 = v49
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v61)+80))
	v154 = int32(0)
	if l0 == v154 {
		goto L49
	} else {
		goto L50
	}
L32:
	;
	v110 = int32(0)
	if v49 == v110 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v148 != 0 {
		v152 = v49
		goto L31
	} else {
		goto L46
	}
L34:
	;
	v148 = int32(0)
	goto L33
L35:
	;
	goto L36
L36:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v49)+4))
	if v116 <= int32(0) {
		v141 = v110
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v148 = v141
	goto L33
L38:
	;
	v119 = int32(0)
	if v119 < v116 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v122 = v116
	goto L41
L40:
	;
	v122 = v119
	goto L41
L41:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v49)+12))
	v125 = int32(0)
	goto L42
L42:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v123+v125<<(uint(int32(2))%32))))
	v134 = base.B2i32(v133 == v107)
	if v133 == v107 {
		v141 = v134
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v141 = v134
	goto L37
L44:
	;
	v136 = v125 + int32(1)
	if v136 != v122 {
		v125 = v136
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v61)+92))
	v150 = F_lappend_oid(m, v49, v149)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v152 = v150
	goto L31
L48:
	;
	if v192 != 0 {
		v196 = v152
		v198 = v52
		goto L15
	} else {
		goto L61
	}
L49:
	;
	v192 = int32(0)
	goto L48
L50:
	;
	goto L51
L51:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v160 <= int32(0) {
		v185 = v154
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v192 = v185
	goto L48
L53:
	;
	v163 = int32(0)
	if v163 < v160 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v166 = v160
	goto L56
L55:
	;
	v166 = v163
	goto L56
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v169 = int32(0)
	goto L57
L57:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v167+v169<<(uint(int32(2))%32))))
	v178 = base.B2i32(v177 == v153)
	if v177 == v153 {
		v185 = v178
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v185 = v178
	goto L52
L59:
	;
	v180 = v169 + int32(1)
	if v180 != v166 {
		v169 = v180
		goto L57
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v61)+80))
	v194 = F_lappend_oid(m, v52, v193)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v196 = v152
	v198 = v194
	goto L15
L63:
	;
	if v199 != 0 {
		v48 = v199
		v49 = v196
		v52 = v198
		goto L13
	} else {
		goto L64
	}
L64:
	;
	goto L14
L65:
	;
	if v196 == int32(0) {
		v324 = v198
		goto L5
	} else {
		goto L66
	}
L66:
	;
	v205 = int32(0)
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v205 < v206 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	F_list_free(m, v310)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L99
	}
L68:
	;
	v211 = int32(0)
	v214 = v29
	v217 = v196
	v220 = v205
	goto L71
L69:
	;
	goto L70
L70:
	;
	v307 = v29
	v310 = v196
	v313 = v205
	goto L67
L71:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v196)+12))
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v225+v211<<(uint(int32(2))%32))))
	F_ScanKeyInit(m, v15, int32(1), int32(3), int32(184), v229)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L73
	}
L72:
	;
	v307 = v293
	v310 = v294
	v313 = v295
	goto L67
L73:
	;
	v233 = int32(1)
	v236 = F_systable_beginscan(m, v23, int32(2667), v233, int32(0), v233, v15)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L75
	}
L74:
	;
	F_systable_endscan(m, v236)
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L1
	} else {
		goto L97
	}
L75:
	;
	v238 = F_systable_getnext(m, v236)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	if v238 == int32(0) {
		v293 = v214
		v294 = v217
		v295 = v220
		goto L74
	} else {
		goto L77
	}
L77:
	;
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v238)+16))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v242)+22)))
	v244 = v242 + v243
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+92))
	if v245 != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v246 = F_list_append_unique_oid(m, v217, v245)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L1
	} else {
		goto L81
	}
L79:
	;
	goto L80
L80:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v244)+96))
	v249 = int32(0)
	if v214 == v249 {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	v293 = v214
	v294 = v246
	v295 = v220
	goto L74
L82:
	;
	if v287 != 0 {
		v293 = v214
		v294 = v217
		v295 = v220
		goto L74
	} else {
		goto L95
	}
L83:
	;
	v287 = int32(0)
	goto L82
L84:
	;
	goto L85
L85:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v255 <= int32(0) {
		v280 = v249
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v287 = v280
	goto L82
L87:
	;
	v258 = int32(0)
	if v258 < v255 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v261 = v255
	goto L90
L89:
	;
	v261 = v258
	goto L90
L90:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v214)+12))
	v264 = int32(0)
	goto L91
L91:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v262+v264<<(uint(int32(2))%32))))
	v273 = base.B2i32(v272 == v248)
	if v272 == v248 {
		v280 = v273
		goto L86
	} else {
		goto L93
	}
L92:
	;
	v280 = v273
	goto L86
L93:
	;
	v275 = v264 + int32(1)
	if v275 != v261 {
		v264 = v275
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	v289 = *(*int32)(unsafe.Add(mBase, uint32(v244)+96))
	v290 = F_lappend_oid(m, v214, v289)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L1
	} else {
		goto L96
	}
L96:
	;
	v293 = v290
	v294 = v217
	v295 = int32(1)
	goto L74
L97:
	;
	v300 = v211 + int32(1)
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v196)+4))
	if v300 < v301 {
		v211 = v300
		v214 = v293
		v217 = v294
		v220 = v295
		goto L71
	} else {
		goto L98
	}
L98:
	;
	goto L72
L99:
	;
	if v313 != 0 {
		v29 = v307
		v30 = v198
		goto L6
	} else {
		goto L100
	}
L100:
	;
	v338 = v307
	v339 = v198
	goto L4
L101:
	;
	v324 = v30
	goto L5
L102:
	;
	v338 = v29
	v339 = v324
	goto L4
L103:
	;
	F_list_free(m, v338)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L104
	}
L104:
	;
	F_list_sort(m, v339, int32(467))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L1
	} else {
		goto L105
	}
L105:
	;
	if v339 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	m.G0 = v15 + int32(48)
	return v339
L107:
	;
	goto L106
L108:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v339)+4))
	if v365 < int32(2) {
		goto L107
	} else {
		goto L109
	}
L109:
	;
	v368 = int32(1)
	v370 = v365 - v368
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v339)+12))
	if v365 == int32(2) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v370&v368 == int32(0) {
		v455 = v429
		goto L123
	} else {
		goto L124
	}
L111:
	;
	v429 = int32(0)
	v431 = v368
	goto L110
L112:
	;
	goto L113
L113:
	;
	v381 = int32(0)
	v384 = v381
	v386 = v368
	v387 = v381
	goto L114
L114:
	;
	v393 = int32(2)
	v394 = v386 << (uint(v393) % 32)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v373+v394)))
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v373+v384<<(uint(v393)%32))))
	if v396 != v400 {
		goto L116
	} else {
		goto L117
	}
L115:
	;
	v429 = v422
	v431 = v424
	goto L110
L116:
	;
	v403 = v384 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v373+v403<<(uint(int32(2))%32)))) = v396
	v408 = v403
	goto L118
L117:
	;
	v408 = v384
	goto L118
L118:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(v394+(v373+int32(4)))))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v373+v408<<(uint(int32(2))%32))))
	if v410 != v414 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v417 = v408 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v373+v417<<(uint(int32(2))%32)))) = v410
	v422 = v417
	goto L121
L120:
	;
	v422 = v408
	goto L121
L121:
	;
	v423 = int32(2)
	v424 = v386 + v423
	v426 = v387 + v423
	if v426 != v370&int32(-2) {
		v384 = v422
		v386 = v424
		v387 = v426
		goto L114
	} else {
		goto L122
	}
L122:
	;
	goto L115
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v339)+4)) = v455 + int32(1)
	goto L107
L124:
	;
	v440 = int32(2)
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v373+v431<<(uint(v440)%32))))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v373+v429<<(uint(v440)%32))))
	if v443 == v447 {
		v455 = v429
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v450 = v429 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v373+v450<<(uint(int32(2))%32)))) = v443
	v455 = v450
	goto L123
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v218 int32
	_ = v218
	var v228 int32
	_ = v228
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
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
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v431 int32
	_ = v431
	var v440 int32
	_ = v440
	var v441 int32
	_ = v441
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v452 int32
	_ = v452
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v481 int32
	_ = v481
	var v486 int32
	_ = v486
	var v488 int32
	_ = v488
	var v510 int32
	_ = v510
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v535 int32
	_ = v535
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v682 int32
	_ = v682
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v729 int32
	_ = v729
	var v740 int32
	_ = v740
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v753 int32
	_ = v753
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v766 int32
	_ = v766
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
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
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v841 int32
	_ = v841
	var v849 int32
	_ = v849
	var v861 int32
	_ = v861
	var v866 int32
	_ = v866
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v886 int32
	_ = v886
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v902 int32
	_ = v902
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v928 int32
	_ = v928
	var v934 int32
	_ = v934
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v992 int32
	_ = v992
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1020 int32
	_ = v1020
	var v1037 int32
	_ = v1037
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1116 int32
	_ = v1116
	var v1119 int32
	_ = v1119
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1210 int32
	_ = v1210
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1236 int32
	_ = v1236
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1263 int32
	_ = v1263
	var v1265 int32
	_ = v1265
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1323 int32
	_ = v1323
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1339 int32
	_ = v1339
	var v1347 int32
	_ = v1347
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1358 int32
	_ = v1358
	var v1360 int32
	_ = v1360
	var v1361 int32
	_ = v1361
	var v1372 int32
	_ = v1372
	var v1398 int32
	_ = v1398
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1412 int32
	_ = v1412
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1443 int32
	_ = v1443
	var v1476 int32
	_ = v1476
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1493 int32
	_ = v1493
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1523 int32
	_ = v1523
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1535 int32
	_ = v1535
	var v1536 int32
	_ = v1536
	var v1538 int32
	_ = v1538
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1552 int32
	_ = v1552
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1567 int32
	_ = v1567
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1573 int32
	_ = v1573
	var v1574 int32
	_ = v1574
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1590 int32
	_ = v1590
	var v1595 int32
	_ = v1595
	var v1604 int32
	_ = v1604
	var v1615 int32
	_ = v1615
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1646 int32
	_ = v1646
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1663 int32
	_ = v1663
	var v1667 int32
	_ = v1667
	var v1668 int32
	_ = v1668
	var v1670 int32
	_ = v1670
	var v1672 int32
	_ = v1672
	var v1674 int32
	_ = v1674
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1679 int32
	_ = v1679
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1719 int32
	_ = v1719
	var v1720 int32
	_ = v1720
	var v1721 int32
	_ = v1721
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1731 int32
	_ = v1731
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1747 int32
	_ = v1747
	var v1749 int32
	_ = v1749
	var v1751 int32
	_ = v1751
	var v1766 int32
	_ = v1766
	var v1772 int32
	_ = v1772
	var v1775 int64
	_ = v1775
	var v1776 int32
	_ = v1776
	var v1782 int32
	_ = v1782
	var v1784 int32
	_ = v1784
	var v1790 int32
	_ = v1790
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1859 int32
	_ = v1859
	var v1862 int32
	_ = v1862
	var v1864 int32
	_ = v1864
	var v1875 int32
	_ = v1875
	var v1880 int32
	_ = v1880
	var v1889 int32
	_ = v1889
	var v1900 int32
	_ = v1900
	var v1906 int32
	_ = v1906
	var v1907 int32
	_ = v1907
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1926 int32
	_ = v1926
	var v1929 int32
	_ = v1929
	var v1934 int32
	_ = v1934
	var v1938 int32
	_ = v1938
	var v1941 int32
	_ = v1941
	var v1945 int32
	_ = v1945
	var v1950 int32
	_ = v1950
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2000 int32
	_ = v2000
	var v2033 int32
	_ = v2033
	var v2037 int32
	_ = v2037
	var v2043 int32
	_ = v2043
	var v2045 int32
	_ = v2045
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2054 int32
	_ = v2054
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2081 int32
	_ = v2081
	var v2088 int32
	_ = v2088
	var v2090 int32
	_ = v2090
	var v2092 int32
	_ = v2092
	var v2093 int32
	_ = v2093
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2101 int32
	_ = v2101
	var v2106 int32
	_ = v2106
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2122 int32
	_ = v2122
	var v2123 int32
	_ = v2123
	var v2128 int32
	_ = v2128
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2152 int32
	_ = v2152
	var v2154 int32
	_ = v2154
	var v2156 int32
	_ = v2156
	var v2158 int32
	_ = v2158
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2164 int32
	_ = v2164
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2176 int32
	_ = v2176
	var v2179 int32
	_ = v2179
	var v2180 int32
	_ = v2180
	var v2185 int32
	_ = v2185
	var v2192 int32
	_ = v2192
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2197 int32
	_ = v2197
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2205 int32
	_ = v2205
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2216 int32
	_ = v2216
	var v2218 int32
	_ = v2218
	var v2222 int32
	_ = v2222
	var v2234 int32
	_ = v2234
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2254 int32
	_ = v2254
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2264 int32
	_ = v2264
	var v2267 int32
	_ = v2267
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2274 int32
	_ = v2274
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2281 int32
	_ = v2281
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2286 int32
	_ = v2286
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2290 int32
	_ = v2290
	var v2292 int32
	_ = v2292
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2296 int32
	_ = v2296
	var v2301 int32
	_ = v2301
	var v2302 int32
	_ = v2302
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2307 int32
	_ = v2307
	var v2312 int32
	_ = v2312
	var v2318 int32
	_ = v2318
	var v2319 int32
	_ = v2319
	var v2323 int32
	_ = v2323
	var v2326 int32
	_ = v2326
	var v2329 int32
	_ = v2329
	var v2336 int32
	_ = v2336
	var v2337 int32
	_ = v2337
	var v2339 int32
	_ = v2339
	var v2344 int32
	_ = v2344
	var v2350 int32
	_ = v2350
	var v2352 int32
	_ = v2352
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2360 int32
	_ = v2360
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2367 int32
	_ = v2367
	var v2368 int32
	_ = v2368
	var v2369 int32
	_ = v2369
	var v2375 int32
	_ = v2375
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2392 int32
	_ = v2392
	var v2399 int32
	_ = v2399
	var v2400 int32
	_ = v2400
	var v2402 int32
	_ = v2402
	var v2407 int32
	_ = v2407
	var v2413 int32
	_ = v2413
	var v2415 int32
	_ = v2415
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2423 int32
	_ = v2423
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2428 int32
	_ = v2428
	var v2432 int32
	_ = v2432
	var v2433 int32
	_ = v2433
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2439 int32
	_ = v2439
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2447 int32
	_ = v2447
	var v2452 int32
	_ = v2452
	var v2457 int32
	_ = v2457
	var v2460 int32
	_ = v2460
	var v2461 int32
	_ = v2461
	var v2464 int32
	_ = v2464
	var v2471 int32
	_ = v2471
	var v2473 int32
	_ = v2473
	var v2475 int32
	_ = v2475
	var v2482 int32
	_ = v2482
	var v2488 int32
	_ = v2488
	var v2490 int32
	_ = v2490
	var v2496 int32
	_ = v2496
	var v2499 int32
	_ = v2499
	var v2503 int32
	_ = v2503
	var v2504 int32
	_ = v2504
	var v2508 int32
	_ = v2508
	var v2511 int32
	_ = v2511
	var v2517 int32
	_ = v2517
	var v2519 int32
	_ = v2519
	var v2520 int32
	_ = v2520
	var v2521 int32
	_ = v2521
	var v2524 int32
	_ = v2524
	var v2526 int32
	_ = v2526
	var v2536 int32
	_ = v2536
	var v2542 int32
	_ = v2542
	var v2544 int32
	_ = v2544
	var v2550 int32
	_ = v2550
	var v2551 int32
	_ = v2551
	var v2552 int32
	_ = v2552
	var v2555 int64
	_ = v2555
	var v2556 int64
	_ = v2556
	var v2557 int64
	_ = v2557
	var v2563 int32
	_ = v2563
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2569 int32
	_ = v2569
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2572 int32
	_ = v2572
	var v2575 int32
	_ = v2575
	var v2576 int32
	_ = v2576
	var v2577 int32
	_ = v2577
	var v2579 int32
	_ = v2579
	var v2595 int32
	_ = v2595
	var v2596 int32
	_ = v2596
	var v2622 int32
	_ = v2622
	var v2624 int32
	_ = v2624
	var v2627 int32
	_ = v2627
	var v2630 int32
	_ = v2630
	var v2632 int32
	_ = v2632
	var v2674 int32
	_ = v2674
	var v2687 int32
	_ = v2687
	var v2713 int32
	_ = v2713
	var v2717 int32
	_ = v2717
	var v2722 int32
	_ = v2722
	var v2735 int32
	_ = v2735
	var v2739 int32
	_ = v2739
	var v2761 int32
	_ = v2761
	var v2763 int32
	_ = v2763
	var v2765 int32
	_ = v2765
	var v2768 int32
	_ = v2768
	var v2771 int32
	_ = v2771
	var v2773 int32
	_ = v2773
	var v2815 int32
	_ = v2815
	var v2829 int32
	_ = v2829
	var v2856 int32
	_ = v2856
	var v2861 int32
	_ = v2861
	var v2866 int32
	_ = v2866
	var v2867 int32
	_ = v2867
	var v2871 int32
	_ = v2871
	var v2882 int32
	_ = v2882
	var v2883 int32
	_ = v2883
	var v2908 int32
	_ = v2908
	var v2909 int32
	_ = v2909
	var v2915 int32
	_ = v2915
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2925 int32
	_ = v2925
	var v2926 int32
	_ = v2926
	var v2930 int32
	_ = v2930
	var v2931 int32
	_ = v2931
	var v2937 int32
	_ = v2937
	var v2946 int32
	_ = v2946
	var v2947 int32
	_ = v2947
	var v2950 int32
	_ = v2950
	var v2951 int32
	_ = v2951
	var v2953 int32
	_ = v2953
	var v2954 int32
	_ = v2954
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2959 int32
	_ = v2959
	var v2960 int32
	_ = v2960
	var v2964 int32
	_ = v2964
	var v2966 int32
	_ = v2966
	var v2970 int32
	_ = v2970
	var v2985 int32
	_ = v2985
	var v2987 int32
	_ = v2987
	var v2988 int32
	_ = v2988
	var v2993 int32
	_ = v2993
	var v2995 int32
	_ = v2995
	var v3001 int32
	_ = v3001
	var v3006 int32
	_ = v3006
	var v3027 int32
	_ = v3027
	var v3035 int32
	_ = v3035
	var v3041 int32
	_ = v3041
	var v3042 int32
	_ = v3042
	var v3043 int32
	_ = v3043
	var v3045 int32
	_ = v3045
	var v3047 int32
	_ = v3047
	var v3054 int32
	_ = v3054
	var v3055 int32
	_ = v3055
	var v3056 int32
	_ = v3056
	var v3060 int32
	_ = v3060
	var v3062 int32
	_ = v3062
	var v3063 int32
	_ = v3063
	var v3068 int32
	_ = v3068
	var v3069 int32
	_ = v3069
	var v3071 int32
	_ = v3071
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3078 int32
	_ = v3078
	var v3079 int32
	_ = v3079
	var v3081 int32
	_ = v3081
	var v3082 int32
	_ = v3082
	var v3083 int32
	_ = v3083
	var v3087 int32
	_ = v3087
	var v3089 int32
	_ = v3089
	var v3090 int32
	_ = v3090
	var v3094 int32
	_ = v3094
	var v3098 int32
	_ = v3098
	var v3105 int32
	_ = v3105
	var v3106 int32
	_ = v3106
	var v3108 int32
	_ = v3108
	var v3110 int32
	_ = v3110
	var v3116 int32
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3118 int32
	_ = v3118
	var v3120 int32
	_ = v3120
	var v3124 int32
	_ = v3124
	var v3127 int32
	_ = v3127
	var v3129 int32
	_ = v3129
	var v3130 int32
	_ = v3130
	var v3135 int64
	_ = v3135
	var v3136 int32
	_ = v3136
	var v3142 int32
	_ = v3142
	var v3148 int32
	_ = v3148
	var v3150 int32
	_ = v3150
	var v3156 int32
	_ = v3156
	var v3163 int32
	_ = v3163
	var v3169 int32
	_ = v3169
	var v3171 int32
	_ = v3171
	var v3177 int32
	_ = v3177
	var v3217 int32
	_ = v3217
	var v3219 int32
	_ = v3219
	var v3227 int32
	_ = v3227
	var v3230 int32
	_ = v3230
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3239 int32
	_ = v3239
	var v3243 int32
	_ = v3243
	var v3245 int32
	_ = v3245
	var v3246 int32
	_ = v3246
	var v3248 int32
	_ = v3248
	var v3249 int32
	_ = v3249
	var v3251 int32
	_ = v3251
	var v3254 int32
	_ = v3254
	var v3259 int32
	_ = v3259
	var v3261 int32
	_ = v3261
	var v3262 int32
	_ = v3262
	var v3265 int32
	_ = v3265
	var v3269 int32
	_ = v3269
	var v3270 int32
	_ = v3270
	var v3271 int32
	_ = v3271
	var v3273 int32
	_ = v3273
	var v3274 int32
	_ = v3274
	var v3275 int32
	_ = v3275
	var v3276 int32
	_ = v3276
	var v3278 int32
	_ = v3278
	var v3279 int32
	_ = v3279
	var v3281 int32
	_ = v3281
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3286 int32
	_ = v3286
	var v3289 int32
	_ = v3289
	var v3293 int32
	_ = v3293
	var v3296 int64
	_ = v3296
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3307 int64
	_ = v3307
	var v3317 int32
	_ = v3317
	var v3318 int32
	_ = v3318
	var v3320 int32
	_ = v3320
	var v3323 int32
	_ = v3323
	var v3326 int32
	_ = v3326
	var v3330 int32
	_ = v3330
	var v3334 int32
	_ = v3334
	var v3336 int32
	_ = v3336
	var v3341 int32
	_ = v3341
	var v3377 int32
	_ = v3377
	var v3397 int32
	_ = v3397
	var v3398 int32
	_ = v3398
	var v3400 int32
	_ = v3400
	var v3401 int32
	_ = v3401
	var v3411 int32
	_ = v3411
	var v3413 int32
	_ = v3413
	var v3415 int32
	_ = v3415
	var v3417 int32
	_ = v3417
	var v3419 int32
	_ = v3419
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
	v56 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)+72))
	if v57 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	F_bms_free(m, v3398)
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		goto L1
	} else {
		goto L733
	}
L4:
	;
	v3377 = v3341
	v3397 = v418
	v3398 = v68
	v3400 = v71
	v3401 = v74
	goto L3
L5:
	;
	if v90 < int32(0) {
		goto L431
	} else {
		goto L432
	}
L6:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1801)))
	v1988 = int32(0)
	v1995 = F_RelationGetBufferForTuple(m, l0, v1987, v90, v1988, v1988, v39+int32(20), v39+int32(24), v1988)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L1
	} else {
		goto L429
	}
L7:
	;
	if v59&int32(1) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v59 = int32(1)
	goto L10
L9:
	;
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+76)))
	v59 = v58
	goto L10
L10:
	;
	goto L7
L11:
	;
	v65 = F_RelationGetIndexAttrBitmap(m, l0, int32(3))
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
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
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L425
	}
L14:
	;
	v68 = F_RelationGetIndexAttrBitmap(m, l0, int32(4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	v71 = F_RelationGetIndexAttrBitmap(m, l0, int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v74 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v77 = F_bms_add_members(m, int32(0), v65)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v79 = F_bms_add_members(m, v77, v68)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v81 = F_bms_add_members(m, v79, v71)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v83 = F_bms_add_members(m, v81, v74)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v89 = v85 | v86<<(uint(int32(16))%32)
	v90 = F_ReadBuffer(m, l0, v89)
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+10)))
	if v110&int32(4) != 0 {
		goto L27
	} else {
		goto L28
	}
L23:
	;
	if v90 < int32(0) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v95+(v90^int32(-1))<<(uint(int32(2))%32))))
	v109 = v101
	goto L22
L25:
	;
	goto L26
L26:
	;
	v103 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v109 = v103 + v90<<(uint(int32(13))%32) + int32(-8192)
	goto L22
L27:
	;
	F_visibilitymap_pin(m, l0, v89, v39+int32(24))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	F_LockBuffer(m, v90, int32(2))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	v120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	v125 = v120<<(uint(int32(2))%32) + v109 + int32(20)
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	if v126&int32(98304) != int32(32768) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_UnlockReleaseBuffer(m, v90)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+44)) = v145
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+48)) = v109 + v147&int32(32767)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+32)) = int32(base.Ui32(v152) >> (uint(int32(17)) % 32))
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+40)) = uint16(v156)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+36)) = v158
	*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v145
	v161 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v83 == int32(0) {
		goto L42
	} else {
		goto L43
	}
L35:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v133 != 0 {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	F_ReleaseBuffer(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v136
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)) = uint16(v139)
	*(*int64)(unsafe.Add(mBase, uint32(l6)+8)) = int64(-4294967296)
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = int32(0)
	v3377 = int32(4)
	v3397 = v74
	v3398 = v65
	v3400 = v68
	v3401 = v71
	goto L3
L39:
	;
	goto L38
L40:
	;
	if int32(0) <= v218 {
		goto L51
	} else {
		goto L52
	}
L41:
	;
	v218 = base.I32_ctz(v204) | v205<<(uint(int32(5))%32)
	goto L40
L42:
	;
	v218 = int32(-2)
	goto L40
L43:
	;
	v171 = base.I32_div_s(int32(0), int32(32))
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v172 <= v171 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v175 = v83 + int32(8)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+v171<<(uint(int32(2))%32))))
	v182 = v179 & int32(-1)
	if v182 != 0 {
		v204 = v182
		v205 = v171
		goto L41
	} else {
		goto L45
	}
L45:
	;
	v184 = v171 + int32(1)
	if v184 == v172 {
		goto L42
	} else {
		goto L46
	}
L46:
	;
	v187 = v184
	goto L47
L47:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v175+v187<<(uint(int32(2))%32))))
	if v194 != 0 {
		v204 = v194
		v205 = v187
		goto L41
	} else {
		goto L49
	}
L48:
	;
	goto L42
L49:
	;
	v196 = v187 + int32(1)
	if v196 != v172 {
		v187 = v196
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v228 = v218
	v246 = v10
	v248 = v10
	goto L54
L52:
	;
	v416 = v10
	v418 = v10
	goto L53
L53:
	;
	v431 = int32(0)
	if v418 == v431 {
		v472 = v431
		goto L95
	} else {
		goto L96
	}
L54:
	;
	v264 = v228<<(uint(int32(16))%32) - int32(458752)
	if v264 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v416 = v335
	v418 = v336
	goto L53
L56:
	;
	if v83 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L57:
	;
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161+int32(8)+int32(base.Ui32(v264)>>(uint(int32(12))%32))))))
	if v321 != int32(65535) {
		v335 = v246
		v336 = v248
		goto L56
	} else {
		goto L78
	}
L58:
	;
	if v270 < int32(0) {
		v335 = v246
		v336 = v248
		goto L56
	} else {
		goto L76
	}
L59:
	;
	v313 = F_bms_add_member(m, v248, v228)
	mBase = m.M
	v314 = m.ExcPending
	if v314 != 0 {
		goto L1
	} else {
		goto L75
	}
L60:
	;
	v270 = v264 >> (uint(int32(16)) % 32)
	if base.B2i32(v264 != int32(-393216))&base.B2i32(v270 < int32(0)) != 0 {
		goto L59
	} else {
		goto L61
	}
L61:
	;
	v278 = F_heap_getattr_1(m, v39+int32(32), v270, v161, v39+int32(80))
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	v282 = F_heap_getattr_1(m, l2, v270, v161, v39+int32(72))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+72)))
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+80)))
	if v285 != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if v285 == v284&int32(1) {
		goto L58
	} else {
		goto L74
	}
L65:
	;
	if v284&int32(1) != 0 {
		goto L64
	} else {
		goto L66
	}
L66:
	;
	if v270 <= int32(0) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	if v278 != v282 {
		goto L59
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v295 = v270<<(uint(int32(4))%32) + (v161 + int32(20)) - int32(16)
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v295)+6)))
	v297 = int32(*(*int16)(unsafe.Add(mBase, uint32(v295)+4)))
	v298 = F_datumIsEqual(m, v278, v282, v296, v297)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L1
	} else {
		goto L71
	}
L70:
	;
	goto L58
L71:
	;
	if v298 == int32(0) {
		goto L59
	} else {
		goto L72
	}
L72:
	;
	v302 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+80)))
	if v302&int32(1) != 0 {
		v335 = v246
		v336 = v248
		goto L56
	} else {
		goto L73
	}
L73:
	;
	goto L57
L74:
	;
	goto L59
L75:
	;
	v335 = v246
	v336 = v313
	goto L56
L76:
	;
	if v285 != 0 {
		v335 = v246
		v336 = v248
		goto L56
	} else {
		goto L77
	}
L77:
	;
	goto L57
L78:
	;
	v324 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278))))
	if v324 != int32(1) {
		v335 = v246
		v336 = v248
		goto L56
	} else {
		goto L79
	}
L79:
	;
	v327 = F_bms_is_member(m, v228, v74)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v335 = v327 | v246
	v336 = v248
	goto L56
L81:
	;
	if int32(0) <= v392 {
		v228 = v392
		v246 = v335
		v248 = v336
		goto L54
	} else {
		goto L92
	}
L82:
	;
	v392 = base.I32_ctz(v378) | v379<<(uint(int32(5))%32)
	goto L81
L83:
	;
	v392 = int32(-2)
	goto L81
L84:
	;
	v343 = v228 + int32(1)
	v345 = base.I32_div_s(v343, int32(32))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v83)+4))
	if v346 <= v345 {
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v349 = v83 + int32(8)
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v349+v345<<(uint(int32(2))%32))))
	v356 = v353 & (int32(-1) << (uint(v343) % 32))
	if v356 != 0 {
		v378 = v356
		v379 = v345
		goto L82
	} else {
		goto L86
	}
L86:
	;
	v358 = v345 + int32(1)
	if v358 == v346 {
		goto L83
	} else {
		goto L87
	}
L87:
	;
	v361 = v358
	goto L88
L88:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v349+v361<<(uint(int32(2))%32))))
	if v368 != 0 {
		v378 = v368
		v379 = v361
		goto L82
	} else {
		goto L90
	}
L89:
	;
	goto L83
L90:
	;
	v370 = v361 + int32(1)
	if v370 != v346 {
		v361 = v370
		goto L88
	} else {
		goto L91
	}
L91:
	;
	goto L89
L92:
	;
	goto L55
L93:
	;
	v488 = v39 + int32(36)
	v510 = int32(0)
	goto L112
L94:
	;
	if v472 == int32(0) {
		goto L108
	} else {
		goto L109
	}
L95:
	;
	goto L94
L96:
	;
	if v71 == int32(0) {
		v472 = v431
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v440 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	if v440 < v441 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v443 = v440
	goto L100
L99:
	;
	v443 = v441
	goto L100
L100:
	;
	if v443 <= int32(1) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v446 = int32(1)
	goto L103
L102:
	;
	v446 = v443
	goto L103
L103:
	;
	v447 = int32(8)
	v452 = int32(0)
	goto L104
L104:
	;
	v459 = v452 << (uint(int32(2)) % 32)
	v461 = *(*int32)(unsafe.Add(mBase, uint32(v71+v447+v459)))
	v463 = *(*int32)(unsafe.Add(mBase, uint32(v459+(v418+v447))))
	v464 = v461 & v463
	v466 = base.B2i32(v464 != int32(0))
	if v464 != 0 {
		v472 = v466
		goto L95
	} else {
		goto L106
	}
L105:
	;
	v472 = v466
	goto L95
L106:
	;
	v468 = v452 + int32(1)
	if v468 != v446 {
		v452 = v468
		goto L104
	} else {
		goto L107
	}
L107:
	;
	goto L105
L108:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(2)
	F_MultiXactIdSetOldestMember(m)
	mBase = m.M
	v481 = m.ExcPending
	if v481 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(3)
	v486 = int32(5)
	goto L93
L111:
	;
	v486 = int32(4)
	goto L93
L112:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	v530 = F_HeapTupleSatisfiesUpdate(m, v39+int32(32), v529, v90)
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L1
	} else {
		goto L119
	}
L113:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+4))
	v1216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1214)+20)))
	v1217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1214)+18)))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	F_compute_new_xmax_infomask(m, v1215, v1216, v1217, v42, v1218, int32(1), v39+int32(12), v39+int32(10), v39+int32(8))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L1
	} else {
		goto L279
	}
L114:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v1198 != 0 {
		goto L273
	} else {
		goto L274
	}
L115:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = v1054
	v1056 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1053)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)) = uint16(v1056)
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1053)+4))
	v1059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1053)+20)))
	if v1059&int32(6272) != int32(4096) {
		goto L244
	} else {
		goto L245
	}
L116:
	;
	if l4 == int32(0) {
		goto L237
	} else {
		goto L238
	}
L117:
	;
	if l5 == int32(0) {
		goto L125
	} else {
		goto L126
	}
L118:
	;
	F_UnlockReleaseBuffer(m, v90)
	mBase = m.M
	v535 = m.ExcPending
	if v535 != 0 {
		goto L1
	} else {
		goto L120
	}
L119:
	;
	switch v530 - int32(1) {
	case 0:
		goto L118
	default:
		v973 = int32(1)
		v975 = v530
		v992 = v510
		goto L116
	case 4:
		goto L117
	}
L120:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v539 = m.ExcPending
	if v539 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	F_errmsg(m, int32(380337), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(491798), int32(3480), int32(352024))
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L125:
	;
	v1020 = int32(5)
	v1037 = v510
	goto L115
L126:
	;
	goto L127
L127:
	;
	v555 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v555)+4))
	v557 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v555)+20)))
	if v557&int32(4096) != 0 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v953 = v928 + int32(12)
	v954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488)+2)))
	v955 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488))))
	v956 = int32(16)
	v959 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v953)+2)))
	v960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v953))))
	if v954|v955<<(uint(v956)%32) == v959|v960<<(uint(v956)%32) {
		goto L230
	} else {
		goto L231
	}
L129:
	;
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v928 = v913
	v934 = v615
	goto L128
L130:
	;
	v560 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+66)) = uint8(v560)
	v562 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v565 = F_DoesMultiXactIdConflict(m, v556, v557, v562, v39+int32(66))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L1
	} else {
		goto L134
	}
L131:
	;
	goto L132
L132:
	;
	v729 = int32(0)
	if base.Ui32(v556) < base.Ui32(int32(3)) {
		goto L164
	} else {
		goto L165
	}
L133:
	;
	if v612&int32(128) != 0 {
		goto L148
	} else {
		goto L149
	}
L134:
	;
	if v565 == int32(0) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v570 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v569)+20)))
	v612 = v570
	v613 = v569
	v615 = v510
	v616 = int32(0)
	goto L133
L136:
	;
	goto L137
L137:
	;
	F_LockBuffer(m, v90, int32(0))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	v575 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+66)))
	if (v575|v510)&int32(1) != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v592 = int32(0)
	v597 = F_Do_MultiXactIdWait(m, v556, v486, v557, v592, l0, v488, int32(1), v39+int32(72), v592)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L1
	} else {
		goto L144
	}
L140:
	;
	v591 = v575 ^ int32(1) | v510
	goto L139
L141:
	;
	goto L142
L142:
	;
	v582 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v582*int32(12))+uint32(_consts[62])))
	F_LockTuple(m, l0, v488, v587)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L1
	} else {
		goto L143
	}
L143:
	;
	v591 = int32(1)
	goto L139
L144:
	;
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	F_LockBuffer(m, v90, int32(2))
	mBase = m.M
	v602 = m.ExcPending
	if v602 != 0 {
		goto L1
	} else {
		goto L145
	}
L145:
	;
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v603)+20)))
	if (v604^v557)&int32(4304) != 0 {
		v510 = v591
		goto L112
	} else {
		goto L146
	}
L146:
	;
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v603)+4))
	if v608 != v556 {
		v510 = v591
		goto L112
	} else {
		goto L147
	}
L147:
	;
	v612 = v604
	v613 = v603
	v615 = v591
	v616 = base.B2i32(v599 != int32(0))
	goto L133
L148:
	;
	v973 = v565 ^ int32(1) | v616
	v975 = int32(0)
	v992 = v615
	goto L116
L149:
	;
	if v612&int32(4176) == int32(64) {
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v623 = *(*int32)(unsafe.Add(mBase, uint32(v613)+4))
	v627 = F_GetMultiXactIdMembers(m, v623, v39+int32(80), int32(0))
	mBase = m.M
	v628 = m.ExcPending
	if v628 != 0 {
		goto L1
	} else {
		goto L151
	}
L151:
	;
	if v627 <= int32(0) {
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v636 = int32(0)
	goto L154
L153:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v671)))
	F_pfree(m, v632)
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L159
	}
L154:
	;
	v671 = v632 + v636<<(uint(int32(3))%32)
	v672 = *(*int32)(unsafe.Add(mBase, uint32(v671)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v672) {
		goto L153
	} else {
		goto L156
	}
L155:
	;
	F_pfree(m, v632)
	mBase = m.M
	v679 = m.ExcPending
	if v679 != 0 {
		goto L1
	} else {
		goto L158
	}
L156:
	;
	v676 = v636 + int32(1)
	if v676 != v627 {
		v636 = v676
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	goto L148
L159:
	;
	if v680 == int32(0) {
		goto L148
	} else {
		goto L160
	}
L160:
	;
	v685 = F_TransactionIdDidAbort(m, v680)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L161
	}
L161:
	;
	if v685 == int32(0) {
		goto L129
	} else {
		goto L162
	}
L162:
	;
	goto L148
L163:
	;
	if v849 != 0 {
		goto L203
	} else {
		goto L204
	}
L164:
	;
	v849 = int32(0)
	goto L163
L165:
	;
	goto L166
L166:
	;
	v740 = *(*int32)(unsafe.Add(mBase, _consts[63]))
	if v740 == v556 {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v849 = int32(1)
	goto L163
L168:
	;
	goto L169
L169:
	;
	v744 = *(*int32)(unsafe.Add(mBase, _consts[64]))
	if v744 <= int32(0) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v849 = v841
	goto L163
L171:
	;
	v748 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	if v748 == int32(0) {
		v841 = v729
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v810 = *(*int32)(unsafe.Add(mBase, _consts[65]))
	v812 = int32(0)
	v814 = v744 - int32(1)
	goto L193
L174:
	;
	v753 = v748
	goto L175
L175:
	;
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v753)+20))
	if v758 == int32(4) {
		goto L177
	} else {
		goto L178
	}
L176:
	;
	v841 = int32(0)
	goto L170
L177:
	;
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v753)+80))
	if v805 != 0 {
		v753 = v805
		goto L175
	} else {
		goto L192
	}
L178:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v753)))
	if v761 == int32(0) {
		goto L177
	} else {
		goto L179
	}
L179:
	;
	v764 = int32(1)
	if v556 == v761 {
		v841 = v764
		goto L170
	} else {
		goto L180
	}
L180:
	;
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v753)+52))
	v768 = v766 - int32(1)
	if v768 < int32(0) {
		goto L177
	} else {
		goto L181
	}
L181:
	;
	v773 = int32(0)
	v775 = v768
	goto L182
L182:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v753)+48))
	v781 = int32(2)
	v782 = base.I32_div_s(v775-v773, v781)
	v783 = v782 + v773
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v779+v783<<(uint(v781)%32))))
	if v787 == v556 {
		v841 = v764
		goto L170
	} else {
		goto L184
	}
L183:
	;
	goto L177
L184:
	;
	v791 = F_TransactionIdPrecedes(m, v787, v556)
	mBase = m.M
	if v791 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v792 = v783 + int32(1)
	goto L187
L186:
	;
	v792 = v773
	goto L187
L187:
	;
	if v791 != 0 {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v795 = v775
	goto L190
L189:
	;
	v795 = v783 - int32(1)
	goto L190
L190:
	;
	if v792 <= v795 {
		v773 = v792
		v775 = v795
		goto L182
	} else {
		goto L191
	}
L191:
	;
	goto L183
L192:
	;
	goto L176
L193:
	;
	v819 = int32(2)
	v820 = base.I32_div_s(v814-v812, v819)
	v821 = v820 + v812
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v810+v821<<(uint(v819)%32))))
	v826 = base.B2i32(v825 == v556)
	if v825 == v556 {
		v841 = v826
		goto L170
	} else {
		goto L195
	}
L194:
	;
	v841 = v826
	goto L170
L195:
	;
	v829 = base.B2i32(base.Ui32(v825) < base.Ui32(v556))
	if base.Ui32(v825) < base.Ui32(v556) {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v830 = v821 + int32(1)
	goto L198
L197:
	;
	v830 = v812
	goto L198
L198:
	;
	if base.Ui32(v825) < base.Ui32(v556) {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v833 = v814
	goto L201
L200:
	;
	v833 = v821 - int32(1)
	goto L201
L201:
	;
	if v830 <= v833 {
		v812 = v830
		v814 = v833
		goto L193
	} else {
		goto L202
	}
L202:
	;
	goto L194
L203:
	;
	v973 = int32(1)
	v975 = v729
	v992 = v510
	goto L116
L204:
	;
	goto L205
L205:
	;
	if (v472^int32(-1))&base.B2i32(v557&int32(80) == int32(16)) != 0 {
		goto L206
	} else {
		goto L207
	}
L206:
	;
	v973 = int32(1)
	v975 = v729
	v992 = v510
	goto L116
L207:
	;
	goto L208
L208:
	;
	F_LockBuffer(m, v90, int32(0))
	mBase = m.M
	v861 = m.ExcPending
	if v861 != 0 {
		goto L1
	} else {
		goto L209
	}
L209:
	;
	if v510&int32(1) == int32(0) {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v866 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v866*int32(12))+uint32(_consts[62])))
	F_LockTuple(m, l0, v488, v871)
	mBase = m.M
	v873 = m.ExcPending
	if v873 != 0 {
		goto L1
	} else {
		goto L213
	}
L211:
	;
	goto L212
L212:
	;
	v874 = int32(1)
	F_XactLockTableWait(m, v556, l0, v488, v874)
	mBase = m.M
	v877 = m.ExcPending
	if v877 != 0 {
		goto L1
	} else {
		goto L214
	}
L213:
	;
	goto L212
L214:
	;
	F_LockBuffer(m, v90, int32(2))
	mBase = m.M
	v880 = m.ExcPending
	if v880 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v882 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v881)+20)))
	if (v882^v557)&int32(4304) != 0 {
		v510 = v874
		goto L112
	} else {
		goto L216
	}
L216:
	;
	v886 = *(*int32)(unsafe.Add(mBase, uint32(v881)+4))
	if v556 != v886 {
		v510 = v874
		goto L112
	} else {
		goto L217
	}
L217:
	;
	if v882&int32(3072) != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v907 = int32(0)
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v910 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v909)+21)))
	if v910&int32(8) != 0 {
		v973 = v907
		v975 = v907
		v992 = v874
		goto L116
	} else {
		goto L227
	}
L219:
	;
	if v882&int32(128) != 0 {
		goto L220
	} else {
		goto L221
	}
L220:
	;
	F_HeapTupleSetHintBits(m, v881, v90, int32(2048), int32(0))
	mBase = m.M
	v906 = m.ExcPending
	if v906 != 0 {
		goto L1
	} else {
		goto L226
	}
L221:
	;
	if v882&int32(4176) == int32(64) {
		goto L220
	} else {
		goto L222
	}
L222:
	;
	v896 = F_TransactionIdDidCommit(m, v556)
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	if v896 == int32(0) {
		goto L220
	} else {
		goto L224
	}
L224:
	;
	F_HeapTupleSetHintBits(m, v881, v90, int32(1024), v556)
	mBase = m.M
	v902 = m.ExcPending
	if v902 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	goto L218
L226:
	;
	goto L218
L227:
	;
	v928 = v909
	v934 = v874
	goto L128
L228:
	;
	if v970 != 0 {
		goto L234
	} else {
		goto L235
	}
L229:
	;
	goto L228
L230:
	;
	v966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488)+4)))
	v967 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v953)+4)))
	if v966 == v967 {
		v970 = int32(1)
		goto L229
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v970 = int32(0)
	goto L229
L233:
	;
	goto L232
L234:
	;
	v971 = int32(4)
	goto L236
L235:
	;
	v971 = int32(3)
	goto L236
L236:
	;
	v1020 = v971
	v1037 = v934
	goto L115
L237:
	;
	if v975 == int32(0) {
		goto L114
	} else {
		goto L242
	}
L238:
	;
	if v975 != 0 {
		goto L237
	} else {
		goto L239
	}
L239:
	;
	v1012 = F_HeapTupleSatisfiesVisibility(m, v39+int32(32), l4, v90)
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L1
	} else {
		goto L240
	}
L240:
	;
	if v1012 != 0 {
		goto L114
	} else {
		goto L241
	}
L241:
	;
	v1020 = int32(3)
	v1037 = v992
	goto L115
L242:
	;
	v1020 = v975
	v1037 = v992
	goto L115
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+8)) = v1125
	if v1020 == int32(2) {
		goto L256
	} else {
		goto L257
	}
L244:
	;
	v1125 = v1058
	goto L243
L245:
	;
	goto L246
L246:
	;
	v1064 = int32(0)
	v1068 = F_GetMultiXactIdMembers(m, v1058, v39+int32(80), v1064)
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L1
	} else {
		goto L247
	}
L247:
	;
	if v1068 <= int32(0) {
		v1125 = v1064
		goto L243
	} else {
		goto L248
	}
L248:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	v1074 = v1064
	goto L251
L249:
	;
	F_pfree(m, v1072)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L255
	}
L250:
	;
	v1119 = *(*int32)(unsafe.Add(mBase, uint32(v1111)))
	v1121 = v1119
	goto L249
L251:
	;
	v1111 = v1072 + v1074<<(uint(int32(3))%32)
	v1112 = *(*int32)(unsafe.Add(mBase, uint32(v1111)+4))
	if base.Ui32(int32(4)) <= base.Ui32(v1112) {
		goto L250
	} else {
		goto L253
	}
L252:
	;
	v1121 = int32(0)
	goto L249
L253:
	;
	v1116 = v1074 + int32(1)
	if v1116 != v1068 {
		v1074 = v1116
		goto L251
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1125 = v1121
	goto L243
L256:
	;
	v1163 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1165 = *(*int32)(unsafe.Add(mBase, uint32(v1163)+8))
	v1166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1163)+20)))
	if v1166&int32(32) != 0 {
		goto L260
	} else {
		goto L261
	}
L257:
	;
	v1177 = int32(-1)
	goto L258
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6)+12)) = v1177
	F_UnlockReleaseBuffer(m, v90)
	mBase = m.M
	v1180 = m.ExcPending
	if v1180 != 0 {
		goto L1
	} else {
		goto L263
	}
L259:
	;
	v1177 = v1175
	goto L258
L260:
	;
	v1170 = *(*int32)(unsafe.Add(mBase, _consts[66]))
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1170+v1165<<(uint(int32(3))%32))+4))
	v1175 = v1174
	goto L262
L261:
	;
	v1175 = v1165
	goto L262
L262:
	;
	goto L259
L263:
	;
	if v1037&int32(1) != 0 {
		goto L264
	} else {
		goto L265
	}
L264:
	;
	v1183 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1183*int32(12))+uint32(_consts[62])))
	F_UnlockTuple(m, l0, v488, v1188)
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L1
	} else {
		goto L267
	}
L265:
	;
	goto L266
L266:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v1191 != 0 {
		goto L268
	} else {
		goto L269
	}
L267:
	;
	goto L266
L268:
	;
	F_ReleaseBuffer(m, v1191)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L1
	} else {
		goto L271
	}
L269:
	;
	goto L270
L270:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = int32(0)
	F_bms_free(m, v65)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L1
	} else {
		goto L272
	}
L271:
	;
	goto L270
L272:
	;
	v3341 = v1020
	goto L4
L273:
	;
	goto L113
L274:
	;
	v1199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+10)))
	if v1199&int32(4) == int32(0) {
		goto L273
	} else {
		goto L275
	}
L275:
	;
	F_LockBuffer(m, v90, int32(0))
	mBase = m.M
	v1206 = m.ExcPending
	if v1206 != 0 {
		goto L1
	} else {
		goto L276
	}
L276:
	;
	F_visibilitymap_pin(m, l0, v89, v39+int32(24))
	mBase = m.M
	v1210 = m.ExcPending
	if v1210 != 0 {
		goto L1
	} else {
		goto L277
	}
L277:
	;
	F_LockBuffer(m, v90, int32(2))
	mBase = m.M
	v1213 = m.ExcPending
	if v1213 != 0 {
		goto L1
	} else {
		goto L278
	}
L278:
	;
	v510 = v992
	goto L112
L279:
	;
	v1229 = int32(10240)
	v1230 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1231 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1230)+20)))
	if v1231&int32(2048) != 0 {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1516 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1515)+20)))
	v1518 = v1516 & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v1515)+20)) = uint16(v1518)
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1520)+18)))
	v1523 = v1521 & int32(8191)
	*(*uint16)(unsafe.Add(mBase, uint32(v1520)+18)) = uint16(v1523)
	v1525 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1525))) = v42
	v1527 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v1527)+8)) = v1528
	v1530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1527)+20)))
	v1532 = v1530 & int32(65503)
	*(*uint16)(unsafe.Add(mBase, uint32(v1527)+20)) = uint16(v1532)
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1535 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1534)+20)))
	v1536 = v1535 | v1483
	*(*uint16)(unsafe.Add(mBase, uint32(v1534)+20)) = uint16(v1536)
	v1538 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1538)+18)))
	v1540 = v1539 | v1482
	*(*uint16)(unsafe.Add(mBase, uint32(v1538)+18)) = uint16(v1540)
	v1542 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v1542)+4)) = v1493
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	F_HeapTupleHeaderAdjustCmax(m, v1544, v39+int32(52), v39+int32(19))
	mBase = m.M
	v1550 = m.ExcPending
	if v1550 != 0 {
		goto L1
	} else {
		goto L335
	}
L281:
	;
	v1234 = int32(0)
	v1482 = v1234
	v1483 = v1229
	v1493 = v1234
	goto L280
L282:
	;
	goto L283
L283:
	;
	v1236 = int32(0)
	if v973&base.B2i32(v1231&int32(4304) != int32(4224)) == v1236 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1482 = int32(0)
	v1483 = v1229
	v1493 = v1236
	goto L280
L285:
	;
	goto L286
L286:
	;
	v1245 = int32(0)
	v1246 = *(*int32)(unsafe.Add(mBase, uint32(v1230)+4))
	if v1246 == v1245 {
		goto L287
	} else {
		goto L288
	}
L287:
	;
	v1482 = v1245
	v1483 = v1229
	v1493 = int32(0)
	goto L280
L288:
	;
	goto L289
L289:
	;
	if v1231&int32(4096) == int32(0) {
		v1482 = v1245
		v1483 = int32(8336)
		v1493 = v1246
		goto L280
	} else {
		goto L290
	}
L290:
	;
	v1259 = F_GetMultiXactIdMembers(m, v1246, v39+int32(80), int32(0))
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L1
	} else {
		goto L292
	}
L291:
	;
	v1482 = v1443
	v1483 = v1476 | int32(8192)
	v1493 = v1246
	goto L280
L292:
	;
	if v1259 <= int32(0) {
		v1443 = v1245
		v1476 = int32(4240)
		goto L291
	} else {
		goto L293
	}
L293:
	;
	v1263 = int32(1)
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v39)+80))
	if v1259 == v1263 {
		goto L295
	} else {
		goto L296
	}
L294:
	;
	if v1259&v1263 == int32(0) {
		v1414 = v1358
		v1415 = v1360
		v1418 = v1372
		goto L315
	} else {
		goto L316
	}
L295:
	;
	v1268 = int32(0)
	v1358 = v1268
	v1360 = v1245
	v1361 = v1268
	v1372 = v1268
	goto L294
L296:
	;
	goto L297
L297:
	;
	v1273 = int32(0)
	v1278 = v1273
	v1280 = v1245
	v1281 = v1273
	v1290 = v1273
	v1292 = v1273
	goto L298
L298:
	;
	v1315 = v1265 + v1281<<(uint(int32(3))%32)
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+4))
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v1316<<(uint(int32(2))%32))+uint32(_consts[67])))
	if base.Ui32(v1278) < base.Ui32(v1321) {
		goto L300
	} else {
		goto L301
	}
L299:
	;
	v1358 = v1351
	v1360 = v1349
	v1361 = v1353
	v1372 = v1350
	goto L294
L300:
	;
	v1323 = v1321
	goto L302
L301:
	;
	v1323 = v1278
	goto L302
L302:
	;
	switch v1316 - int32(3) {
	case 0:
		goto L306
	case 1:
		v1330 = v1280
		goto L304
	case 2:
		goto L305
	default:
		v1332 = v1280
		v1333 = v1292
		goto L303
	}
L303:
	;
	v1334 = *(*int32)(unsafe.Add(mBase, uint32(v1315)+12))
	v1339 = *(*int32)(unsafe.Add(mBase, uint32(v1334<<(uint(int32(2))%32))+uint32(_consts[67])))
	switch v1334 - int32(3) {
	case 0:
		goto L310
	case 1:
		v1347 = v1332
		goto L308
	case 2:
		goto L309
	default:
		v1349 = v1332
		v1350 = v1333
		goto L307
	}
L304:
	;
	v1332 = v1330
	v1333 = int32(1)
	goto L303
L305:
	;
	v1330 = v1280 | int32(8192)
	goto L304
L306:
	;
	v1332 = v1280 | int32(8192)
	v1333 = v1292
	goto L303
L307:
	;
	if base.Ui32(v1323) < base.Ui32(v1339) {
		goto L311
	} else {
		goto L312
	}
L308:
	;
	v1349 = v1347
	v1350 = int32(1)
	goto L307
L309:
	;
	v1347 = v1332 | int32(8192)
	goto L308
L310:
	;
	v1349 = v1332 | int32(8192)
	v1350 = v1333
	goto L307
L311:
	;
	v1351 = v1339
	goto L313
L312:
	;
	v1351 = v1323
	goto L313
L313:
	;
	v1352 = int32(2)
	v1353 = v1281 + v1352
	v1355 = v1290 + v1352
	if v1355 != v1259&int32(2147483646) {
		v1278 = v1351
		v1280 = v1349
		v1281 = v1353
		v1290 = v1355
		v1292 = v1350
		goto L298
	} else {
		goto L314
	}
L314:
	;
	goto L299
L315:
	;
	F_pfree(m, v1265)
	mBase = m.M
	v1420 = m.ExcPending
	if v1420 != 0 {
		goto L1
	} else {
		goto L323
	}
L316:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1265+v1361<<(uint(int32(3))%32))+4))
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v1398<<(uint(int32(2))%32))+uint32(_consts[67])))
	if base.Ui32(v1358) < base.Ui32(v1403) {
		goto L317
	} else {
		goto L318
	}
L317:
	;
	v1405 = v1403
	goto L319
L318:
	;
	v1405 = v1358
	goto L319
L319:
	;
	switch v1398 - int32(3) {
	case 0:
		goto L322
	case 1:
		v1412 = v1360
		goto L320
	case 2:
		goto L321
	default:
		v1414 = v1405
		v1415 = v1360
		v1418 = v1372
		goto L315
	}
L320:
	;
	v1414 = v1405
	v1415 = v1412
	v1418 = int32(1)
	goto L315
L321:
	;
	v1412 = v1360 | int32(8192)
	goto L320
L322:
	;
	v1414 = v1405
	v1415 = v1360 | int32(8192)
	v1418 = v1372
	goto L315
L323:
	;
	if v1414&int32(-2) == int32(2) {
		goto L324
	} else {
		goto L325
	}
L324:
	;
	if v1418&int32(1) != 0 {
		v1443 = v1415
		v1476 = int32(4160)
		goto L291
	} else {
		goto L327
	}
L325:
	;
	goto L326
L326:
	;
	if v1414 != 0 {
		goto L328
	} else {
		goto L329
	}
L327:
	;
	v1443 = v1415
	v1476 = int32(4288)
	goto L291
L328:
	;
	v1432 = int32(4096)
	goto L330
L329:
	;
	v1432 = int32(4112)
	goto L330
L330:
	;
	if v1414 == int32(1) {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1435 = int32(4176)
	goto L333
L332:
	;
	v1435 = v1432
	goto L333
L333:
	;
	if v1418&int32(1) != 0 {
		v1443 = v1415
		v1476 = v1435
		goto L291
	} else {
		goto L334
	}
L334:
	;
	v1443 = v1415
	v1476 = v1435 | int32(128)
	goto L291
L335:
	;
	v1551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1551)+119)))
	switch v1552 - int32(109) {
	case 0, 5:
		goto L337
	default:
		v1567 = int32(0)
		goto L336
	}
L336:
	;
	v1571 = int32(4)
	v1572 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+14)))
	v1573 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+12)))
	v1574 = v1572 - v1573
	if v1574 <= v1571 {
		goto L341
	} else {
		goto L342
	}
L337:
	;
	v1555 = int32(1)
	v1556 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1556)+20)))
	if v1557&int32(4) != 0 {
		v1567 = v1555
		goto L336
	} else {
		goto L338
	}
L338:
	;
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v1561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1560)+20)))
	if v1561&int32(4) != 0 {
		v1567 = v1555
		goto L336
	} else {
		goto L339
	}
L339:
	;
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1567 = base.B2i32(base.Ui32(int32(2032)) < base.Ui32(v1564))
	goto L336
L340:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1641 = (v1637 + int32(7)) & int32(-8)
	if v1567 == int32(0) {
		goto L359
	} else {
		goto L360
	}
L341:
	;
	v1577 = v1571
	goto L343
L342:
	;
	v1577 = v1574
	goto L343
L343:
	;
	v1579 = v1577 - int32(4)
	if v1579 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v1636 = int32(0)
	goto L340
L345:
	;
	goto L346
L346:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1573) {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	v1636 = v1579
	goto L340
L348:
	;
	v1590 = int32(base.Ui32(v1573+int32(262120)) >> (uint(int32(2)) % 32))
	goto L350
L349:
	;
	v1590 = int32(0)
	goto L350
L350:
	;
	if base.Ui32(v1590&int32(65535)) < base.Ui32(int32(291)) {
		goto L347
	} else {
		goto L351
	}
L351:
	;
	v1595 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+10)))
	if v1595&int32(1) == int32(0) {
		goto L352
	} else {
		goto L353
	}
L352:
	;
	v1636 = int32(0)
	goto L340
L353:
	;
	goto L354
L354:
	;
	v1604 = int32(1)
	goto L355
L355:
	;
	v1615 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1604&int32(65535)<<(uint(int32(2))%32)+(v109+int32(24))-int32(3)))))
	if v1615&int32(384) == int32(0) {
		goto L347
	} else {
		goto L357
	}
L356:
	;
	v1636 = int32(0)
	goto L340
L357:
	;
	v1621 = v1604 + int32(1)
	v1622 = int32(65535)
	if base.Ui32(v1621&v1622) <= base.Ui32(v1590&v1622) {
		v1604 = v1621
		goto L355
	} else {
		goto L358
	}
L358:
	;
	goto L356
L359:
	;
	if base.Ui32(v1641) <= base.Ui32(v1636) {
		v2000 = l2
		v2033 = v90
		goto L5
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	v1646 = int32(0)
	v1647 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1648 = *(*int32)(unsafe.Add(mBase, uint32(v1647)+4))
	v1649 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1647)+20)))
	v1650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1647)+18)))
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	F_compute_new_xmax_infomask(m, v1648, v1649, v1650, v42, v1651, v1646, v39+int32(72), v39+int32(66), v39+int32(62))
	mBase = m.M
	v1660 = m.ExcPending
	if v1660 != 0 {
		goto L1
	} else {
		goto L363
	}
L362:
	;
	goto L361
L363:
	;
	v1661 = int32(4470804)
	v1663 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1663 + int32(1)
	v1667 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1668 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1667)+20)))
	v1670 = v1668 & int32(9007)
	*(*uint16)(unsafe.Add(mBase, uint32(v1667)+20)) = uint16(v1670)
	v1672 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1667)+18)))
	v1674 = v1672 & int32(57343)
	*(*uint16)(unsafe.Add(mBase, uint32(v1667)+18)) = uint16(v1674)
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1677 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1676)+18)))
	v1679 = v1677 & int32(49151)
	*(*uint16)(unsafe.Add(mBase, uint32(v1676)+18)) = uint16(v1679)
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1682 = *(*int32)(unsafe.Add(mBase, uint32(v39)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v1681)+4)) = v1682
	v1684 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1685 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1684)+20)))
	v1686 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+66)))
	v1687 = v1685 | v1686
	*(*uint16)(unsafe.Add(mBase, uint32(v1684)+20)) = uint16(v1687)
	v1689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1684)+18)))
	v1690 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+62)))
	v1691 = v1689 | v1690
	*(*uint16)(unsafe.Add(mBase, uint32(v1684)+18)) = uint16(v1691)
	v1693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+19)))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v1694)+8)) = v1695
	v1697 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1694)+20)))
	if v1693 != 0 {
		goto L364
	} else {
		goto L365
	}
L364:
	;
	v1702 = int32(32)
	goto L366
L365:
	;
	v1702 = int32(0)
	goto L366
L366:
	;
	v1703 = v1697&int32(65503) | v1702
	*(*uint16)(unsafe.Add(mBase, uint32(v1694)+20)) = uint16(v1703)
	v1705 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v488)))
	*(*int32)(unsafe.Add(mBase, uint32(v1705)+12)) = v1706
	v1708 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v488)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1705)+16)) = uint16(v1708)
	v1710 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+10)))
	if v1710&int32(4) != 0 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v1715 = F_visibilitymap_clear(m, v89, v1713, int32(2))
	mBase = m.M
	v1716 = m.ExcPending
	if v1716 != 0 {
		goto L1
	} else {
		goto L370
	}
L368:
	;
	v1717 = v1646
	goto L369
L369:
	;
	F_MarkBufferDirty(m, v90)
	mBase = m.M
	v1719 = m.ExcPending
	if v1719 != 0 {
		goto L1
	} else {
		goto L371
	}
L370:
	;
	v1717 = v1715
	goto L369
L371:
	;
	v1720 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v1721 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1720)+118)))
	if v1721 != int32(112) {
		goto L372
	} else {
		goto L373
	}
L372:
	;
	v1782 = int32(4470804)
	v1784 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v1784 - int32(1)
	F_LockBuffer(m, v90, int32(0))
	mBase = m.M
	v1790 = m.ExcPending
	if v1790 != 0 {
		goto L1
	} else {
		goto L383
	}
L373:
	;
	v1725 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v1725 <= int32(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1728 != 0 {
		goto L372
	} else {
		goto L377
	}
L375:
	;
	goto L376
L376:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L1
	} else {
		goto L379
	}
L377:
	;
	v1729 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v1729 != 0 {
		goto L372
	} else {
		goto L378
	}
L378:
	;
	goto L376
L379:
	;
	F_XLogRegisterBuffer(m, int32(0), v90, int32(8))
	mBase = m.M
	v1735 = m.ExcPending
	if v1735 != 0 {
		goto L1
	} else {
		goto L380
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = v1682
	v1737 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+40)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+84)) = uint16(v1737)
	v1739 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v1740 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1739)+18)))
	v1741 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1739)+20)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+87)) = uint8(v1717)
	v1747 = int32(1)
	v1749 = int32(8)
	v1751 = int32(4)
	v1766 = int32(base.Ui32(v1740)>>(uint(int32(9))%32))&int32(16) | (int32(base.Ui32(v1741)>>(uint(v1747)%32))&v1749 | (int32(base.Ui32(v1741)>>(uint(v1751)%32))&v1751 | (int32(base.Ui32(v1741)>>(uint(int32(12))%32))&v1747 | int32(base.Ui32(v1741)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+86)) = uint8(v1766)
	F_XLogRegisterData(m, v39+int32(80), v1749)
	mBase = m.M
	v1772 = m.ExcPending
	if v1772 != 0 {
		goto L1
	} else {
		goto L381
	}
L381:
	;
	v1775 = F_XLogInsert(m, int32(10), int32(96))
	mBase = m.M
	v1776 = m.ExcPending
	if v1776 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v109))) = base.I64_rotr(v1775, int64(32))
	goto L372
L383:
	;
	if v1567 != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v1794 = F_heap_toast_insert_or_update(m, l0, l2, v39+int32(32), int32(0))
	mBase = m.M
	v1795 = m.ExcPending
	if v1795 != 0 {
		goto L1
	} else {
		goto L387
	}
L385:
	;
	v1801 = l2
	v1802 = v1641
	goto L386
L386:
	;
	if base.Ui32(v1636) < base.Ui32(v1802) {
		goto L6
	} else {
		goto L388
	}
L387:
	;
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(v1794)))
	v1801 = v1794
	v1802 = (v1796 + int32(7)) & int32(-8)
	goto L386
L388:
	;
	goto L389
L389:
	;
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v1840 != 0 {
		goto L391
	} else {
		goto L392
	}
L391:
	;
	F_LockBuffer(m, v90, int32(2))
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L1
	} else {
		goto L395
	}
L392:
	;
	v1841 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+10)))
	if v1841&int32(4) == int32(0) {
		goto L391
	} else {
		goto L393
	}
L393:
	;
	F_visibilitymap_pin(m, l0, v89, v39+int32(24))
	mBase = m.M
	v1849 = m.ExcPending
	if v1849 != 0 {
		goto L1
	} else {
		goto L394
	}
L394:
	;
	goto L391
L395:
	;
	v1856 = int32(4)
	v1857 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+14)))
	v1858 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+12)))
	v1859 = v1857 - v1858
	if v1859 <= v1856 {
		goto L397
	} else {
		goto L398
	}
L396:
	;
	if base.Ui32(v1921) < base.Ui32(v1802) {
		goto L415
	} else {
		goto L416
	}
L397:
	;
	v1862 = v1856
	goto L399
L398:
	;
	v1862 = v1859
	goto L399
L399:
	;
	v1864 = v1862 - int32(4)
	if v1864 == int32(0) {
		goto L400
	} else {
		goto L401
	}
L400:
	;
	v1921 = int32(0)
	goto L396
L401:
	;
	goto L402
L402:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1858) {
		goto L404
	} else {
		goto L405
	}
L403:
	;
	v1921 = v1864
	goto L396
L404:
	;
	v1875 = int32(base.Ui32(v1858+int32(262120)) >> (uint(int32(2)) % 32))
	goto L406
L405:
	;
	v1875 = int32(0)
	goto L406
L406:
	;
	if base.Ui32(v1875&int32(65535)) < base.Ui32(int32(291)) {
		goto L403
	} else {
		goto L407
	}
L407:
	;
	v1880 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+10)))
	if v1880&int32(1) == int32(0) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v1921 = int32(0)
	goto L396
L409:
	;
	goto L410
L410:
	;
	v1889 = int32(1)
	goto L411
L411:
	;
	v1900 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1889&int32(65535)<<(uint(int32(2))%32)+(v109+int32(24))-int32(3)))))
	if v1900&int32(384) == int32(0) {
		goto L403
	} else {
		goto L413
	}
L412:
	;
	v1921 = int32(0)
	goto L396
L413:
	;
	v1906 = v1889 + int32(1)
	v1907 = int32(65535)
	if base.Ui32(v1906&v1907) <= base.Ui32(v1875&v1907) {
		v1889 = v1906
		goto L411
	} else {
		goto L414
	}
L414:
	;
	goto L412
L415:
	;
	F_LockBuffer(m, v90, int32(0))
	mBase = m.M
	v1925 = m.ExcPending
	if v1925 != 0 {
		goto L1
	} else {
		goto L418
	}
L416:
	;
	goto L417
L417:
	;
	v1926 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v1926 == int32(0) {
		goto L420
	} else {
		goto L421
	}
L418:
	;
	goto L6
L419:
	;
	F_LockBuffer(m, v90, int32(0))
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L1
	} else {
		goto L424
	}
L420:
	;
	v1929 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v109)+10)))
	if v1929&int32(4) != 0 {
		goto L419
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	v2000 = v1801
	v2033 = v90
	goto L5
L423:
	;
	goto L422
L424:
	;
	goto L389
L425:
	;
	F_errcode(m, int32(322))
	mBase = m.M
	v1941 = m.ExcPending
	if v1941 != 0 {
		goto L1
	} else {
		goto L426
	}
L426:
	;
	F_errmsg(m, int32(257727), int32(0))
	mBase = m.M
	v1945 = m.ExcPending
	if v1945 != 0 {
		goto L1
	} else {
		goto L427
	}
L427:
	;
	F_errfinish(m, int32(491798), int32(3303), int32(352024))
	mBase = m.M
	v1950 = m.ExcPending
	if v1950 != 0 {
		goto L1
	} else {
		goto L428
	}
L428:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L429:
	;
	v2000 = v1801
	v2033 = v1995
	goto L5
L430:
	;
	F_CheckForSerializableConflictIn(m, l0, v488, v2052)
	mBase = m.M
	v2054 = m.ExcPending
	if v2054 != 0 {
		goto L1
	} else {
		goto L434
	}
L431:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v2043 = *(*int32)(unsafe.Add(mBase, uint32(v2037+(v90^int32(-1))<<(uint(int32(6))%32))+16))
	v2052 = v2043
	goto L430
L432:
	;
	goto L433
L433:
	;
	v2045 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v2045+v90<<(uint(int32(6))%32)+int32(-64))+16))
	v2052 = v2051
	goto L430
L434:
	;
	v2055 = base.B2i32(v2033 != v90)
	if v2055 == int32(0) {
		goto L436
	} else {
		goto L437
	}
L435:
	;
	v2164 = int32(0)
	if v418 == v2164 {
		v2205 = v2164
		goto L472
	} else {
		goto L473
	}
L436:
	;
	v2058 = int32(0)
	if v418 == v2058 {
		v2101 = v2058
		goto L440
	} else {
		goto L441
	}
L437:
	;
	goto L438
L438:
	;
	v2154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v109)+10)))
	v2156 = v2154 | int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v109)+10)) = uint16(v2156)
	v2158 = int32(0)
	v2160 = v2158
	v2161 = v2158
	goto L435
L439:
	;
	if v2101 != 0 {
		v2160 = v2058
		v2161 = v2058
		goto L435
	} else {
		goto L453
	}
L440:
	;
	goto L439
L441:
	;
	if v65 == int32(0) {
		v2101 = v2058
		goto L440
	} else {
		goto L442
	}
L442:
	;
	v2069 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v65)+4))
	if v2069 < v2070 {
		goto L443
	} else {
		goto L444
	}
L443:
	;
	v2072 = v2069
	goto L445
L444:
	;
	v2072 = v2070
	goto L445
L445:
	;
	if v2072 <= int32(1) {
		goto L446
	} else {
		goto L447
	}
L446:
	;
	v2075 = int32(1)
	goto L448
L447:
	;
	v2075 = v2072
	goto L448
L448:
	;
	v2076 = int32(8)
	v2081 = int32(0)
	goto L449
L449:
	;
	v2088 = v2081 << (uint(int32(2)) % 32)
	v2090 = *(*int32)(unsafe.Add(mBase, uint32(v65+v2076+v2088)))
	v2092 = *(*int32)(unsafe.Add(mBase, uint32(v2088+(v418+v2076))))
	v2093 = v2090 & v2092
	v2095 = base.B2i32(v2093 != int32(0))
	if v2093 != 0 {
		v2101 = v2095
		goto L440
	} else {
		goto L451
	}
L450:
	;
	v2101 = v2095
	goto L440
L451:
	;
	v2097 = v2081 + int32(1)
	if v2097 != v2075 {
		v2081 = v2097
		goto L449
	} else {
		goto L452
	}
L452:
	;
	goto L450
L453:
	;
	v2106 = int32(0)
	if v418 == v2106 {
		v2148 = v2106
		goto L455
	} else {
		goto L456
	}
L454:
	;
	if v2148 != 0 {
		goto L468
	} else {
		goto L469
	}
L455:
	;
	goto L454
L456:
	;
	if v68 == int32(0) {
		v2148 = v2106
		goto L455
	} else {
		goto L457
	}
L457:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v2116 < v2117 {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	v2119 = v2116
	goto L460
L459:
	;
	v2119 = v2117
	goto L460
L460:
	;
	if v2119 <= int32(1) {
		goto L461
	} else {
		goto L462
	}
L461:
	;
	v2122 = int32(1)
	goto L463
L462:
	;
	v2122 = v2119
	goto L463
L463:
	;
	v2123 = int32(8)
	v2128 = int32(0)
	goto L464
L464:
	;
	v2135 = v2128 << (uint(int32(2)) % 32)
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v68+v2123+v2135)))
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v2135+(v418+v2123))))
	v2140 = v2137 & v2139
	v2142 = base.B2i32(v2140 != int32(0))
	if v2140 != 0 {
		v2148 = v2142
		goto L455
	} else {
		goto L466
	}
L465:
	;
	v2148 = v2142
	goto L455
L466:
	;
	v2144 = v2128 + int32(1)
	if v2144 != v2122 {
		v2128 = v2144
		goto L464
	} else {
		goto L467
	}
L467:
	;
	goto L465
L468:
	;
	v2152 = int32(2)
	goto L470
L469:
	;
	v2152 = v2106
	goto L470
L470:
	;
	v2160 = v2152
	v2161 = int32(1)
	goto L435
L471:
	;
	v2214 = F_ExtractReplicaIdentity(m, l0, v39+int32(32), (v2205|v416)&int32(1), v39+int32(31))
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L1
	} else {
		goto L485
	}
L472:
	;
	goto L471
L473:
	;
	if v74 == int32(0) {
		v2205 = v2164
		goto L472
	} else {
		goto L474
	}
L474:
	;
	v2173 = *(*int32)(unsafe.Add(mBase, uint32(v418)+4))
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
	if v2173 < v2174 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v2176 = v2173
	goto L477
L476:
	;
	v2176 = v2174
	goto L477
L477:
	;
	if v2176 <= int32(1) {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v2179 = int32(1)
	goto L480
L479:
	;
	v2179 = v2176
	goto L480
L480:
	;
	v2180 = int32(8)
	v2185 = int32(0)
	goto L481
L481:
	;
	v2192 = v2185 << (uint(int32(2)) % 32)
	v2194 = *(*int32)(unsafe.Add(mBase, uint32(v74+v2180+v2192)))
	v2196 = *(*int32)(unsafe.Add(mBase, uint32(v2192+(v418+v2180))))
	v2197 = v2194 & v2196
	v2199 = base.B2i32(v2197 != int32(0))
	if v2197 != 0 {
		v2205 = v2199
		goto L472
	} else {
		goto L483
	}
L482:
	;
	v2205 = v2199
	goto L472
L483:
	;
	v2201 = v2185 + int32(1)
	if v2201 != v2179 {
		v2185 = v2201
		goto L481
	} else {
		goto L484
	}
L484:
	;
	goto L482
L485:
	;
	v2216 = int32(4470804)
	v2218 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v2218 + int32(1)
	v2222 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	if v2222 != 0 {
		goto L487
	} else {
		goto L488
	}
L486:
	;
	v2238 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2238)+18)))
	if v2161 != 0 {
		goto L496
	} else {
		goto L497
	}
L487:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v2222))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v42)) == int32(0) {
		goto L491
	} else {
		goto L492
	}
L488:
	;
	goto L489
L489:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v42
	goto L486
L490:
	;
	if v2234 == int32(0) {
		goto L486
	} else {
		goto L494
	}
L491:
	;
	v2234 = base.B2i32(base.Ui32(v42) < base.Ui32(v2222))
	goto L490
L492:
	;
	goto L493
L493:
	;
	v2234 = int32(base.Ui32(v42-v2222) >> (uint(int32(31)) % 32))
	goto L490
L494:
	;
	goto L489
L495:
	;
	v2267 = int32(0)
	F_RelationPutHeapTuple(m, v2033, v2000, v2267)
	mBase = m.M
	v2270 = m.ExcPending
	if v2270 != 0 {
		goto L1
	} else {
		goto L499
	}
L496:
	;
	v2241 = v2239 | int32(16384)
	*(*uint16)(unsafe.Add(mBase, uint32(v2238)+18)) = uint16(v2241)
	v2243 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+16))
	v2244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2243)+18)))
	v2245 = int32(32768)
	v2246 = v2244 | v2245
	*(*uint16)(unsafe.Add(mBase, uint32(v2243)+18)) = uint16(v2246)
	v2248 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2248)+18)))
	v2251 = v2249 | v2245
	*(*uint16)(unsafe.Add(mBase, uint32(v2248)+18)) = uint16(v2251)
	goto L495
L497:
	;
	goto L498
L498:
	;
	v2254 = v2239 & int32(49151)
	*(*uint16)(unsafe.Add(mBase, uint32(v2238)+18)) = uint16(v2254)
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+16))
	v2257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2256)+18)))
	v2258 = int32(32767)
	v2259 = v2257 & v2258
	*(*uint16)(unsafe.Add(mBase, uint32(v2256)+18)) = uint16(v2259)
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v2262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2261)+18)))
	v2264 = v2262 & v2258
	*(*uint16)(unsafe.Add(mBase, uint32(v2261)+18)) = uint16(v2264)
	goto L495
L499:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2271)+20)))
	v2274 = v2272 & int32(9007)
	*(*uint16)(unsafe.Add(mBase, uint32(v2271)+20)) = uint16(v2274)
	v2276 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2271)+18)))
	v2278 = v2276 & int32(57343)
	*(*uint16)(unsafe.Add(mBase, uint32(v2271)+18)) = uint16(v2278)
	v2280 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2281 = *(*int32)(unsafe.Add(mBase, uint32(v39)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2280)+4)) = v2281
	v2283 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2283)+20)))
	v2285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+10)))
	v2286 = v2284 | v2285
	*(*uint16)(unsafe.Add(mBase, uint32(v2283)+20)) = uint16(v2286)
	v2288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2283)+18)))
	v2289 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+8)))
	v2290 = v2288 | v2289
	*(*uint16)(unsafe.Add(mBase, uint32(v2283)+18)) = uint16(v2290)
	v2292 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+19)))
	v2293 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2294 = *(*int32)(unsafe.Add(mBase, uint32(v39)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v2293)+8)) = v2294
	v2296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2293)+20)))
	if v2292 != 0 {
		goto L500
	} else {
		goto L501
	}
L500:
	;
	v2301 = int32(32)
	goto L502
L501:
	;
	v2301 = int32(0)
	goto L502
L502:
	;
	v2302 = v2296&int32(65503) | v2301
	*(*uint16)(unsafe.Add(mBase, uint32(v2293)+20)) = uint16(v2302)
	v2304 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2305 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2304)+12)) = v2305
	v2307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2000)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2304)+16)) = uint16(v2307)
	if v90 < int32(0) {
		goto L505
	} else {
		goto L506
	}
L503:
	;
	v2368 = int32(0)
	v2369 = base.B2i32(v2033 == v90)
	if v2369 == v2368 {
		goto L515
	} else {
		goto L516
	}
L504:
	;
	v2339 = v2336 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v2337)+10)) = uint16(v2339)
	if v90 < int32(0) {
		goto L511
	} else {
		goto L512
	}
L505:
	;
	v2312 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v2318 = *(*int32)(unsafe.Add(mBase, uint32(v2312+(v90^int32(-1))<<(uint(int32(2))%32))))
	v2319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2318)+10)))
	if v2319&int32(4) != 0 {
		v2336 = v2319
		v2337 = v2318
		goto L504
	} else {
		goto L508
	}
L506:
	;
	goto L507
L507:
	;
	v2323 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v2326 = v2323 + v90<<(uint(int32(13))%32)
	v2329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2326-int32(8182)))))
	if v2329&int32(4) == int32(0) {
		v2367 = v2267
		goto L503
	} else {
		goto L509
	}
L508:
	;
	v2367 = v2267
	goto L503
L509:
	;
	v2336 = v2329
	v2337 = v2326 + int32(-8192)
	goto L504
L510:
	;
	v2360 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	v2362 = F_visibilitymap_clear(m, v2359, v2360, int32(3))
	mBase = m.M
	v2363 = m.ExcPending
	if v2363 != 0 {
		goto L1
	} else {
		goto L514
	}
L511:
	;
	v2344 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v2350 = *(*int32)(unsafe.Add(mBase, uint32(v2344+(v90^int32(-1))<<(uint(int32(6))%32))+16))
	v2359 = v2350
	goto L510
L512:
	;
	goto L513
L513:
	;
	v2352 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2352+v90<<(uint(int32(6))%32)+int32(-64))+16))
	v2359 = v2358
	goto L510
L514:
	;
	v2367 = int32(1)
	goto L503
L515:
	;
	if v2033 < int32(0) {
		goto L520
	} else {
		goto L521
	}
L516:
	;
	v2433 = v2368
	goto L517
L517:
	;
	F_MarkBufferDirty(m, v90)
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		goto L1
	} else {
		goto L531
	}
L518:
	;
	F_MarkBufferDirty(m, v2033)
	mBase = m.M
	v2432 = m.ExcPending
	if v2432 != 0 {
		goto L1
	} else {
		goto L530
	}
L519:
	;
	v2402 = v2399 & int32(65531)
	*(*uint16)(unsafe.Add(mBase, uint32(v2400)+10)) = uint16(v2402)
	if v2033 < int32(0) {
		goto L526
	} else {
		goto L527
	}
L520:
	;
	v2375 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v2381 = *(*int32)(unsafe.Add(mBase, uint32(v2375+(v2033^int32(-1))<<(uint(int32(2))%32))))
	v2382 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2381)+10)))
	if v2382&int32(4) != 0 {
		v2399 = v2382
		v2400 = v2381
		goto L519
	} else {
		goto L523
	}
L521:
	;
	goto L522
L522:
	;
	v2386 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v2389 = v2386 + v2033<<(uint(int32(13))%32)
	v2392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2389-int32(8182)))))
	if v2392&int32(4) == int32(0) {
		v2428 = v2368
		goto L518
	} else {
		goto L524
	}
L523:
	;
	v2428 = v2368
	goto L518
L524:
	;
	v2399 = v2392
	v2400 = v2389 + int32(-8192)
	goto L519
L525:
	;
	v2423 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	v2425 = F_visibilitymap_clear(m, v2422, v2423, int32(3))
	mBase = m.M
	v2426 = m.ExcPending
	if v2426 != 0 {
		goto L1
	} else {
		goto L529
	}
L526:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2407+(v2033^int32(-1))<<(uint(int32(6))%32))+16))
	v2422 = v2413
	goto L525
L527:
	;
	goto L528
L528:
	;
	v2415 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v2415+v2033<<(uint(int32(6))%32)+int32(-64))+16))
	v2422 = v2421
	goto L525
L529:
	;
	v2428 = int32(1)
	goto L518
L530:
	;
	v2433 = v2428
	goto L517
L531:
	;
	v2438 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2439 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2438)+118)))
	if v2439 != int32(112) {
		goto L532
	} else {
		goto L533
	}
L532:
	;
	v3217 = int32(4470804)
	v3219 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v3219 - int32(1)
	if v2369 == int32(0) {
		goto L681
	} else {
		goto L682
	}
L533:
	;
	v2443 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v2443 <= int32(0) {
		goto L535
	} else {
		goto L536
	}
L534:
	;
	v2475 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v2475)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)) = uint16(v2475)
	if v2033 < v2475 {
		goto L551
	} else {
		goto L552
	}
L535:
	;
	v2446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2446 != 0 {
		goto L532
	} else {
		goto L538
	}
L536:
	;
	goto L537
L537:
	;
	if v2443 == int32(1) {
		goto L534
	} else {
		goto L540
	}
L538:
	;
	v2447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2447 == int32(0) {
		goto L534
	} else {
		goto L539
	}
L539:
	;
	goto L532
L540:
	;
	v2452 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L541
L541:
	;
	if base.B2i32(base.Ui32(v2452) < base.Ui32(int32(12000))) == int32(0) {
		goto L542
	} else {
		goto L543
	}
L542:
	;
	v2457 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v2457 == int32(0) {
		goto L534
	} else {
		goto L545
	}
L543:
	;
	goto L544
L544:
	;
	F_log_heap_new_cid(m, l0, v39+int32(32))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L1
	} else {
		goto L548
	}
L545:
	;
	v2460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2460)+119)))
	switch v2461 - int32(109) {
	case 0, 5:
		goto L546
	default:
		goto L534
	}
L546:
	;
	v2464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2457)+104)))
	if v2464 != int32(1) {
		goto L534
	} else {
		goto L547
	}
L547:
	;
	goto L544
L548:
	;
	F_log_heap_new_cid(m, l0, v2000)
	mBase = m.M
	v2473 = m.ExcPending
	if v2473 != 0 {
		goto L1
	} else {
		goto L549
	}
L549:
	;
	goto L534
L550:
	;
	v2499 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v2499 < int32(2) {
		v2517 = int32(0)
		goto L554
	} else {
		goto L555
	}
L551:
	;
	v2482 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v2488 = *(*int32)(unsafe.Add(mBase, uint32(v2482+(v2033^int32(-1))<<(uint(int32(2))%32))))
	v2496 = v2488
	goto L550
L552:
	;
	goto L553
L553:
	;
	v2490 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v2496 = v2490 + v2033<<(uint(int32(13))%32) + int32(-8192)
	goto L550
L554:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v2519 = m.ExcPending
	if v2519 != 0 {
		goto L1
	} else {
		goto L559
	}
L555:
	;
	v2503 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2503)+118)))
	if v2504 != int32(112) {
		v2517 = int32(0)
		goto L554
	} else {
		goto L556
	}
L556:
	;
	v2508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2503)+119)))
	if v2508 == int32(102) {
		v2517 = int32(0)
		goto L554
	} else {
		goto L557
	}
L557:
	;
	v2511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L558
L558:
	;
	v2517 = base.B2i32(base.Ui32(v2511) < base.Ui32(int32(12000))) ^ int32(1)
	goto L554
L559:
	;
	v2520 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+16))
	v2521 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2520)+18)))
	if v2033 != v90 {
		goto L561
	} else {
		goto L562
	}
L560:
	;
	v2909 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+87)) = uint8(v2908)
	if v2517 == v2909 {
		goto L612
	} else {
		goto L613
	}
L561:
	;
	v2867 = int32(0)
	if v2433 != 0 {
		goto L609
	} else {
		goto L610
	}
L562:
	;
	if v2517 != 0 {
		goto L561
	} else {
		goto L563
	}
L563:
	;
	v2524 = m.G0
	v2526 = v2524 - int32(16)
	m.G0 = v2526
	F_GetFullPageWriteInfo(m, v2526+int32(8), v2526+int32(7))
	mBase = m.M
	if v90 < int32(0) {
		goto L566
	} else {
		goto L567
	}
L564:
	;
	if v2563 != 0 {
		goto L561
	} else {
		goto L574
	}
L565:
	;
	v2551 = int32(1)
	v2552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2526)+7)))
	if v2552 == v2551 {
		goto L570
	} else {
		goto L571
	}
L566:
	;
	v2536 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v2542 = *(*int32)(unsafe.Add(mBase, uint32(v2536+(v90^int32(-1))<<(uint(int32(2))%32))))
	v2550 = v2542
	goto L565
L567:
	;
	goto L568
L568:
	;
	v2544 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v2550 = v2544 + v90<<(uint(int32(13))%32) + int32(-8192)
	goto L565
L569:
	;
	m.G0 = v2526 + int32(16)
	goto L564
L570:
	;
	v2555 = *(*int64)(unsafe.Add(mBase, uint32(v2526)+8))
	v2556 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2550)+4)))
	v2557 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v2550))))
	if base.Ui64(v2556|v2557<<(uint(int64(32))%64)) <= base.Ui64(v2555) {
		v2563 = v2551
		goto L569
	} else {
		goto L573
	}
L571:
	;
	goto L572
L572:
	;
	v2563 = int32(0)
	goto L569
L573:
	;
	goto L572
L574:
	;
	v2567 = *(*int32)(unsafe.Add(mBase, uint32(v2000)))
	v2568 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+16))
	v2569 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2568)+22)))
	v2570 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2570)+22)))
	v2572 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v2572)
	v2575 = *(*int32)(unsafe.Add(mBase, uint32(v39)+32))
	v2576 = v2575 - v2571
	v2577 = v2567 - v2569
	if v2576 < v2577 {
		goto L576
	} else {
		goto L577
	}
L575:
	;
	v2713 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)) = uint16(v2713)
	v2717 = v2579 - v2687&int32(65535)
	if v2713 < v2717 {
		goto L590
	} else {
		goto L591
	}
L576:
	;
	v2579 = v2576
	goto L578
L577:
	;
	v2579 = v2577
	goto L578
L578:
	;
	if int32(0) < v2579 {
		goto L579
	} else {
		goto L580
	}
L579:
	;
	v2595 = int32(0)
	v2596 = v2572
	goto L582
L580:
	;
	goto L581
L581:
	;
	v2674 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v2674)
	v2687 = v2674
	goto L575
L582:
	;
	v2622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2596+(v2569+v2568)))))
	v2624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2596+(v2571+v2570)))))
	if v2622 == v2624 {
		goto L584
	} else {
		goto L585
	}
L583:
	;
	if base.Ui32(int32(2)) < base.Ui32(v2632&int32(65535)) {
		v2687 = v2632
		goto L575
	} else {
		goto L588
	}
L584:
	;
	v2627 = v2595 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)) = uint16(v2627)
	v2630 = v2627 & int32(65535)
	if base.Ui32(v2630) < base.Ui32(v2579) {
		v2595 = v2627
		v2596 = v2630
		goto L582
	} else {
		goto L587
	}
L585:
	;
	v2632 = v2595
	goto L586
L586:
	;
	goto L583
L587:
	;
	v2632 = v2627
	goto L586
L588:
	;
	goto L581
L589:
	;
	if v2433 != 0 {
		goto L600
	} else {
		goto L601
	}
L590:
	;
	v2722 = int32(0)
	v2735 = v2722
	v2739 = v2722
	goto L593
L591:
	;
	goto L592
L592:
	;
	v2815 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)) = uint16(v2815)
	v2829 = v2815
	goto L589
L593:
	;
	v2761 = v2739 ^ int32(-1)
	v2763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2567+v2568+v2761))))
	v2765 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2761+(v2570+v2575)))))
	if v2763 == v2765 {
		goto L595
	} else {
		goto L596
	}
L594:
	;
	if base.Ui32(int32(2)) < base.Ui32(v2773&int32(65535)) {
		v2829 = v2773
		goto L589
	} else {
		goto L599
	}
L595:
	;
	v2768 = v2735 + int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)) = uint16(v2768)
	v2771 = v2768 & int32(65535)
	if base.Ui32(v2771) < base.Ui32(v2717) {
		v2735 = v2768
		v2739 = v2771
		goto L593
	} else {
		goto L598
	}
L596:
	;
	v2773 = v2735
	goto L597
L597:
	;
	goto L594
L598:
	;
	v2773 = v2768
	goto L597
L599:
	;
	goto L592
L600:
	;
	v2856 = v2367 | int32(2)
	goto L602
L601:
	;
	v2856 = v2367
	goto L602
L602:
	;
	if v2687&int32(65535) != 0 {
		goto L603
	} else {
		goto L604
	}
L603:
	;
	v2861 = v2856 | int32(32)
	goto L605
L604:
	;
	v2861 = v2856
	goto L605
L605:
	;
	if v2829&int32(65535) != 0 {
		goto L606
	} else {
		goto L607
	}
L606:
	;
	v2866 = v2861 | int32(64)
	goto L608
L607:
	;
	v2866 = v2861
	goto L608
L608:
	;
	v2882 = v2687
	v2883 = v2829
	v2908 = v2866
	goto L560
L609:
	;
	v2871 = v2367 | int32(2)
	goto L611
L610:
	;
	v2871 = v2367
	goto L611
L611:
	;
	v2882 = v2867
	v2883 = v2867
	v2908 = v2871
	goto L560
L612:
	;
	if v2521 < v2909 {
		goto L618
	} else {
		goto L619
	}
L613:
	;
	v2915 = v2908 | int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+87)) = uint8(v2915)
	if v2214 == int32(0) {
		goto L612
	} else {
		goto L614
	}
L614:
	;
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2921)+130)))
	if v2922 == int32(102) {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v2925 = int32(20)
	goto L617
L616:
	;
	v2925 = int32(24)
	goto L617
L617:
	;
	v2926 = v2925 | v2908
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+87)) = uint8(v2926)
	goto L612
L618:
	;
	v2930 = int32(64)
	goto L620
L619:
	;
	v2930 = int32(32)
	goto L620
L620:
	;
	v2931 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2000)+8)))
	if v2931 != int32(1) {
		goto L622
	} else {
		goto L623
	}
L621:
	;
	v2954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+40)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+84)) = uint16(v2954)
	v2956 = *(*int32)(unsafe.Add(mBase, uint32(v39)+48))
	v2957 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+80)) = v2957
	v2959 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2956)+18)))
	v2960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2956)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+92)) = uint16(v2931)
	v2964 = int32(16)
	v2966 = int32(1)
	v2970 = int32(4)
	v2985 = int32(base.Ui32(v2959)>>(uint(int32(9))%32))&v2964 | (int32(base.Ui32(v2960)>>(uint(v2966)%32))&int32(8) | (int32(base.Ui32(v2960)>>(uint(v2970)%32))&v2970 | (int32(base.Ui32(v2960)>>(uint(int32(12))%32))&v2966 | int32(base.Ui32(v2960)>>(uint(int32(6))%32))&int32(2))))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+86)) = uint8(v2985)
	v2987 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+16))
	v2988 = *(*int32)(unsafe.Add(mBase, uint32(v2987)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v39)+88)) = v2988
	if v2517 != 0 {
		goto L631
	} else {
		goto L632
	}
L622:
	;
	v2951 = v2930
	v2953 = int32(8)
	goto L621
L623:
	;
	goto L624
L624:
	;
	v2937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2496)+12)))
	v2946 = base.B2i32(base.Ui32(int32(24)) < base.Ui32(v2937)) & base.B2i32((v2937+int32(262120))&int32(262140) == int32(4))
	if v2946 != 0 {
		goto L625
	} else {
		goto L626
	}
L625:
	;
	v2947 = int32(14)
	goto L627
L626:
	;
	v2947 = int32(8)
	goto L627
L627:
	;
	if v2946 != 0 {
		goto L628
	} else {
		goto L629
	}
L628:
	;
	v2950 = v2930 | int32(-128)
	goto L630
L629:
	;
	v2950 = v2930
	goto L630
L630:
	;
	v2951 = v2950
	v2953 = v2947
	goto L621
L631:
	;
	v2993 = v2953 | v2964
	goto L633
L632:
	;
	v2993 = v2953
	goto L633
L633:
	;
	F_XLogRegisterBuffer(m, int32(0), v2033, v2993)
	mBase = m.M
	v2995 = m.ExcPending
	if v2995 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	if v2369 == int32(0) {
		goto L635
	} else {
		goto L636
	}
L635:
	;
	F_XLogRegisterBuffer(m, int32(1), v90, int32(8))
	mBase = m.M
	v3001 = m.ExcPending
	if v3001 != 0 {
		goto L1
	} else {
		goto L638
	}
L636:
	;
	goto L637
L637:
	;
	F_XLogRegisterData(m, v39+int32(80), int32(14))
	mBase = m.M
	v3006 = m.ExcPending
	if v3006 != 0 {
		goto L1
	} else {
		goto L639
	}
L638:
	;
	goto L637
L639:
	;
	if (v2882|v2883)&int32(65535) == int32(0) {
		goto L640
	} else {
		goto L641
	}
L640:
	;
	v3042 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+16))
	v3043 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3042)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+72)) = uint16(v3043)
	v3045 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3042)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+74)) = uint16(v3045)
	v3047 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3042)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+76)) = uint8(v3047)
	F_XLogRegisterBufData(m, int32(0), v39+int32(72), int32(5))
	mBase = m.M
	v3054 = m.ExcPending
	if v3054 != 0 {
		goto L1
	} else {
		goto L651
	}
L641:
	;
	if v2883&int32(65535) == int32(0) {
		goto L642
	} else {
		goto L643
	}
L642:
	;
	if v2882&int32(65535) != 0 {
		goto L646
	} else {
		goto L647
	}
L643:
	;
	if v2882&int32(65535) == int32(0) {
		goto L642
	} else {
		goto L644
	}
L644:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+64)) = uint16(v2883)
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+62)) = uint16(v2882)
	F_XLogRegisterBufData(m, int32(0), v39+int32(62), int32(4))
	mBase = m.M
	v3027 = m.ExcPending
	if v3027 != 0 {
		goto L1
	} else {
		goto L645
	}
L645:
	;
	goto L640
L646:
	;
	F_XLogRegisterBufData(m, int32(0), v39+int32(60), int32(2))
	mBase = m.M
	v3035 = m.ExcPending
	if v3035 != 0 {
		goto L1
	} else {
		goto L649
	}
L647:
	;
	goto L648
L648:
	;
	F_XLogRegisterBufData(m, int32(0), v39+int32(58), int32(2))
	mBase = m.M
	v3041 = m.ExcPending
	if v3041 != 0 {
		goto L1
	} else {
		goto L650
	}
L649:
	;
	goto L640
L650:
	;
	goto L640
L651:
	;
	v3055 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+16))
	v3056 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)))
	if v3056 == int32(0) {
		goto L653
	} else {
		goto L654
	}
L652:
	;
	v3098 = int32(0)
	if base.B2i32(v2214 == v3098)|(v2517^int32(1)) == v3098 {
		goto L662
	} else {
		goto L663
	}
L653:
	;
	v3060 = int32(23)
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(v2000)))
	v3063 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)))
	F_XLogRegisterBufData(m, int32(0), v3055+v3060, v3062-v3063-v3060)
	mBase = m.M
	v3068 = m.ExcPending
	if v3068 != 0 {
		goto L1
	} else {
		goto L656
	}
L654:
	;
	goto L655
L655:
	;
	v3069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3055)+22)))
	v3071 = v3069 - int32(23)
	if v3071 != 0 {
		goto L657
	} else {
		goto L658
	}
L656:
	;
	goto L652
L657:
	;
	F_XLogRegisterBufData(m, int32(0), v3055+int32(23), v3071)
	mBase = m.M
	v3076 = m.ExcPending
	if v3076 != 0 {
		goto L1
	} else {
		goto L660
	}
L658:
	;
	v3081 = v3055
	v3082 = v3056
	v3083 = int32(23)
	goto L659
L659:
	;
	v3087 = v3083 + v3082&int32(65535)
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(v2000)))
	v3090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+58)))
	F_XLogRegisterBufData(m, int32(0), v3087+v3081, v3089-(v3087+v3090))
	mBase = m.M
	v3094 = m.ExcPending
	if v3094 != 0 {
		goto L1
	} else {
		goto L661
	}
L660:
	;
	v3077 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+60)))
	v3078 = *(*int32)(unsafe.Add(mBase, uint32(v2000)+16))
	v3079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3078)+22)))
	v3081 = v3078
	v3082 = v3077
	v3083 = v3079
	goto L659
L661:
	;
	goto L652
L662:
	;
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(v2214)+16))
	v3106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3105)+18)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+66)) = uint16(v3106)
	v3108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3105)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v39)+68)) = uint16(v3108)
	v3110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3105)+22)))
	*(*uint8)(unsafe.Add(mBase, uint32(v39)+70)) = uint8(v3110)
	F_XLogRegisterData(m, v39+int32(66), int32(5))
	mBase = m.M
	v3116 = m.ExcPending
	if v3116 != 0 {
		goto L1
	} else {
		goto L665
	}
L663:
	;
	goto L664
L664:
	;
	v3127 = int32(4371988)
	v3129 = int32(*(*uint8)(unsafe.Add(mBase, _consts[60])))
	v3130 = v3129 | int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[60])) = uint8(v3130)
	goto L667
L665:
	;
	v3117 = *(*int32)(unsafe.Add(mBase, uint32(v2214)+16))
	v3118 = int32(23)
	v3120 = *(*int32)(unsafe.Add(mBase, uint32(v2214)))
	F_XLogRegisterData(m, v3117+v3118, v3120-v3118)
	mBase = m.M
	v3124 = m.ExcPending
	if v3124 != 0 {
		goto L1
	} else {
		goto L666
	}
L666:
	;
	goto L664
L667:
	;
	v3135 = F_XLogInsert(m, int32(10), v2951&int32(255))
	mBase = m.M
	v3136 = m.ExcPending
	if v3136 != 0 {
		goto L1
	} else {
		goto L668
	}
L668:
	;
	if v2369 == int32(0) {
		goto L669
	} else {
		goto L670
	}
L669:
	;
	if v2033 < int32(0) {
		goto L673
	} else {
		goto L674
	}
L670:
	;
	goto L671
L671:
	;
	if v90 < int32(0) {
		goto L677
	} else {
		goto L678
	}
L672:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3156))) = base.I64_rotr(v3135, int64(32))
	goto L671
L673:
	;
	v3142 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v3148 = *(*int32)(unsafe.Add(mBase, uint32(v3142+(v2033^int32(-1))<<(uint(int32(2))%32))))
	v3156 = v3148
	goto L672
L674:
	;
	goto L675
L675:
	;
	v3150 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v3156 = v3150 + v2033<<(uint(int32(13))%32) + int32(-8192)
	goto L672
L676:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3177))) = base.I64_rotr(v3135, int64(32))
	goto L532
L677:
	;
	v3163 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v3169 = *(*int32)(unsafe.Add(mBase, uint32(v3163+(v90^int32(-1))<<(uint(int32(2))%32))))
	v3177 = v3169
	goto L676
L678:
	;
	goto L679
L679:
	;
	v3171 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v3177 = v3171 + v90<<(uint(int32(13))%32) + int32(-8192)
	goto L676
L680:
	;
	F_ReleaseBuffer(m, v90)
	mBase = m.M
	v3245 = m.ExcPending
	if v3245 != 0 {
		goto L1
	} else {
		goto L690
	}
L681:
	;
	F_LockBuffer(m, v2033, int32(0))
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L1
	} else {
		goto L684
	}
L682:
	;
	goto L683
L683:
	;
	F_LockBuffer(m, v90, int32(0))
	mBase = m.M
	v3239 = m.ExcPending
	if v3239 != 0 {
		goto L1
	} else {
		goto L688
	}
L684:
	;
	F_LockBuffer(m, v90, int32(0))
	mBase = m.M
	v3230 = m.ExcPending
	if v3230 != 0 {
		goto L1
	} else {
		goto L685
	}
L685:
	;
	F_CacheInvalidateHeapTuple(m, l0, v39+int32(32), v2000)
	mBase = m.M
	v3234 = m.ExcPending
	if v3234 != 0 {
		goto L1
	} else {
		goto L686
	}
L686:
	;
	F_ReleaseBuffer(m, v2033)
	mBase = m.M
	v3236 = m.ExcPending
	if v3236 != 0 {
		goto L1
	} else {
		goto L687
	}
L687:
	;
	goto L680
L688:
	;
	F_CacheInvalidateHeapTuple(m, l0, v39+int32(32), v2000)
	mBase = m.M
	v3243 = m.ExcPending
	if v3243 != 0 {
		goto L1
	} else {
		goto L689
	}
L689:
	;
	goto L680
L690:
	;
	v3246 = *(*int32)(unsafe.Add(mBase, uint32(v39)+20))
	if v3246 != 0 {
		goto L691
	} else {
		goto L692
	}
L691:
	;
	F_ReleaseBuffer(m, v3246)
	mBase = m.M
	v3248 = m.ExcPending
	if v3248 != 0 {
		goto L1
	} else {
		goto L694
	}
L692:
	;
	goto L693
L693:
	;
	v3249 = *(*int32)(unsafe.Add(mBase, uint32(v39)+24))
	if v3249 != 0 {
		goto L695
	} else {
		goto L696
	}
L694:
	;
	goto L693
L695:
	;
	F_ReleaseBuffer(m, v3249)
	mBase = m.M
	v3251 = m.ExcPending
	if v3251 != 0 {
		goto L1
	} else {
		goto L698
	}
L696:
	;
	goto L697
L697:
	;
	if v992&int32(1) != 0 {
		goto L699
	} else {
		goto L700
	}
L698:
	;
	goto L697
L699:
	;
	v3254 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v3259 = *(*int32)(unsafe.Add(mBase, uint32(v3254*int32(12))+uint32(_consts[62])))
	F_UnlockTuple(m, l0, v488, v3259)
	mBase = m.M
	v3261 = m.ExcPending
	if v3261 != 0 {
		goto L1
	} else {
		goto L702
	}
L700:
	;
	goto L701
L701:
	;
	v3262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	if v3262 == int32(0) {
		goto L704
	} else {
		goto L705
	}
L702:
	;
	goto L701
L703:
	;
	if l2 != v2000 {
		goto L721
	} else {
		goto L722
	}
L704:
	;
	v3265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+268)))
	if v3265 != int32(1) {
		goto L703
	} else {
		goto L707
	}
L705:
	;
	v3271 = v3262
	goto L706
L706:
	;
	v3273 = *(*int32)(unsafe.Add(mBase, _consts[61]))
	v3274 = *(*int32)(unsafe.Add(mBase, uint32(v3273)+28))
	goto L709
L707:
	;
	F_pgstat_assoc_relation(m, l0)
	mBase = m.M
	v3269 = m.ExcPending
	if v3269 != 0 {
		goto L1
	} else {
		goto L708
	}
L708:
	;
	v3270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+272))
	v3271 = v3270
	goto L706
L709:
	;
	v3275 = *(*int32)(unsafe.Add(mBase, uint32(v3271)+8))
	if v3275 != 0 {
		goto L711
	} else {
		goto L712
	}
L710:
	;
	v3296 = *(*int64)(unsafe.Add(mBase, uint32(v3293)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v3293)+8)) = v3296 + int64(1)
	if v2161|v2055 == int32(0) {
		goto L703
	} else {
		goto L717
	}
L711:
	;
	v3276 = *(*int32)(unsafe.Add(mBase, uint32(v3275)+56))
	if v3276 == v3274 {
		v3293 = v3275
		goto L710
	} else {
		goto L714
	}
L712:
	;
	goto L713
L713:
	;
	v3278 = F_pgstat_get_xact_stack_level(m, v3274)
	mBase = m.M
	v3279 = m.ExcPending
	if v3279 != 0 {
		goto L1
	} else {
		goto L715
	}
L714:
	;
	goto L713
L715:
	;
	v3281 = *(*int32)(unsafe.Add(mBase, _consts[68]))
	v3283 = F_MemoryContextAllocZero(m, v3281, int32(72))
	mBase = m.M
	v3284 = m.ExcPending
	if v3284 != 0 {
		goto L1
	} else {
		goto L716
	}
L716:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3283)+56)) = v3274
	v3286 = *(*int32)(unsafe.Add(mBase, uint32(v3271)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v3283)+64)) = v3271
	*(*int32)(unsafe.Add(mBase, uint32(v3283)+60)) = v3286
	v3289 = *(*int32)(unsafe.Add(mBase, uint32(v3278)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v3283)+68)) = v3289
	*(*int32)(unsafe.Add(mBase, uint32(v3278)+20)) = v3283
	*(*int32)(unsafe.Add(mBase, uint32(v3271)+8)) = v3283
	v3293 = v3283
	goto L710
L717:
	;
	if v2161 != 0 {
		goto L718
	} else {
		goto L719
	}
L718:
	;
	v3305 = int32(64)
	goto L720
L719:
	;
	v3305 = int32(72)
	goto L720
L720:
	;
	v3306 = v3271 + v3305
	v3307 = *(*int64)(unsafe.Add(mBase, uint32(v3306)))
	*(*int64)(unsafe.Add(mBase, uint32(v3306))) = v3307 + int64(1)
	goto L703
L721:
	;
	v3317 = v2000 + int32(4)
	v3318 = *(*int32)(unsafe.Add(mBase, uint32(v3317)))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v3318
	v3320 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3317)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(l2)+8)) = uint16(v3320)
	F_pfree(m, v2000)
	mBase = m.M
	v3323 = m.ExcPending
	if v3323 != 0 {
		goto L1
	} else {
		goto L724
	}
L722:
	;
	goto L723
L723:
	;
	if v2161 != 0 {
		goto L725
	} else {
		goto L726
	}
L724:
	;
	goto L723
L725:
	;
	v3326 = v2160
	goto L727
L726:
	;
	v3326 = int32(1)
	goto L727
L727:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = v3326
	if v2214 == int32(0) {
		goto L728
	} else {
		goto L729
	}
L728:
	;
	F_bms_free(m, v65)
	mBase = m.M
	v3336 = m.ExcPending
	if v3336 != 0 {
		goto L1
	} else {
		goto L732
	}
L729:
	;
	v3330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+31)))
	if v3330 != int32(1) {
		goto L728
	} else {
		goto L730
	}
L730:
	;
	F_pfree(m, v2214)
	mBase = m.M
	v3334 = m.ExcPending
	if v3334 != 0 {
		goto L1
	} else {
		goto L731
	}
L731:
	;
	goto L728
L732:
	;
	v3341 = int32(0)
	goto L4
L733:
	;
	F_bms_free(m, v3400)
	mBase = m.M
	v3413 = m.ExcPending
	if v3413 != 0 {
		goto L1
	} else {
		goto L734
	}
L734:
	;
	F_bms_free(m, v3401)
	mBase = m.M
	v3415 = m.ExcPending
	if v3415 != 0 {
		goto L1
	} else {
		goto L735
	}
L735:
	;
	F_bms_free(m, v3397)
	mBase = m.M
	v3417 = m.ExcPending
	if v3417 != 0 {
		goto L1
	} else {
		goto L736
	}
L736:
	;
	F_bms_free(m, v83)
	mBase = m.M
	v3419 = m.ExcPending
	if v3419 != 0 {
		goto L1
	} else {
		goto L737
	}
L737:
	;
	m.G0 = v39 + int32(96)
	return v3377
}
func F_rewrite_heap_tuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v72 int32
	_ = v72
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v325 int64
	_ = v325
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int64
	_ = v332
	var v334 int32
	_ = v334
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v452 int32
	_ = v452
	v18 = m.G0
	v20 = v18 - int32(80)
	m.G0 = v20
	v22 = int32(4476144)
	v23 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v25
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
	*(*int64)(unsafe.Add(mBase, uint32(v27))) = v29
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v27)+8)) = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+20)))
	v36 = v34 & int32(15)
	*(*uint16)(unsafe.Add(mBase, uint32(v33)+20)) = uint16(v36)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+18)))
	v41 = v39 & int32(8191)
	*(*uint16)(unsafe.Add(mBase, uint32(v38)+18)) = uint16(v41)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v43)+20)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v45)+20)))
	v49 = v44 | v46&int32(65520)
	*(*uint16)(unsafe.Add(mBase, uint32(v43)+20)) = uint16(v49)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+48))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)+136))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v53)+140))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v58 = m.G0
	v60 = v58 + int32(-64)
	m.G0 = v60
	*(*int32)(unsafe.Add(mBase, uint32(v60)+44)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v60)+40)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v60)+36)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v60)+32)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v60)+28)) = v55
	*(*int32)(unsafe.Add(mBase, uint32(v60)+24)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v60)+20)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v60)+16)) = v56
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v56
	v72 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v72)
	v82 = F_heap_prepare_freeze_tuple(m, v51, v58+int32(-40), v58+int32(-60), v58+int32(-12), v58+int32(-13))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if v82 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v60)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v84
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+60)))
	if v86&int32(2) != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	m.G0 = v60 - int32(-64)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v104 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v103)+16)) = uint16(v104)
	*(*int32)(unsafe.Add(mBase, uint32(v103)+12)) = int32(-1)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+21)))
	if v109&int32(8) != 0 {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = int32(2)
	goto L8
L7:
	;
	goto L8
L8:
	;
	if v86&int32(4) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v51)+8)) = int32(0)
	goto L11
L10:
	;
	goto L11
L11:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+58)))
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+20)) = uint16(v95)
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v60)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v51)+18)) = uint16(v97)
	goto L5
L12:
	;
	*(*int32)(unsafe.Add(mBase, _consts[0])) = v23
	m.G0 = v20 + int32(80)
	return
L13:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v205
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)) = uint16(v207)
	v212 = v20 + int32(8) | int32(4)
	v214 = v20 + int32(62)
	v216 = v20 + int32(44)
	v218 = v20 + int32(56)
	v222 = l2
	v223 = int32(0)
	goto L41
L14:
	;
	v112 = F_HeapTupleHeaderIsOnlyLocked(m, v108)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	if v112 != 0 {
		goto L13
	} else {
		goto L16
	}
L16:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v115 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+16)))
	if v115 == int32(65533) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+12)))
	v119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114)+14)))
	if v118&v119 == int32(65535) {
		goto L13
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v124 = l1 + int32(4)
	v126 = v114 + int32(12)
	v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+2)))
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	v129 = int32(16)
	v132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+2)))
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126))))
	if v127|v128<<(uint(v129)%32) == v132|v133<<(uint(v129)%32) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	goto L19
L21:
	;
	if v143 != 0 {
		goto L13
	} else {
		goto L27
	}
L22:
	;
	goto L21
L23:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v126)+4)))
	if v139 == v140 {
		v143 = int32(1)
		goto L22
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v143 = int32(0)
	goto L22
L26:
	;
	goto L25
L27:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = int64(0)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v148)+20)))
	if v149&int32(6272) == int32(4096) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v159
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v158)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v161)
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v158)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v163
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v168 = int32(0)
	v170 = F_hash_search(m, v165, v20+int32(8), v168, v168)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L33
	}
L29:
	;
	v154 = F_HeapTupleGetUpdateXid(m, v148)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L32
	}
L30:
	;
	goto L31
L31:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v158 = v148
	v159 = v157
	goto L28
L32:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v158 = v156
	v159 = v154
	goto L28
L33:
	;
	if v170 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v180 = F_hash_search(m, v174, v20+int32(8), int32(1), v20+int32(7))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L1
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v189)+12)) = v190
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v189)+16)) = uint16(v192)
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v200 = F_hash_search(m, v194, v20+int32(8), int32(2), v20+int32(7))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L39
	}
L37:
	;
	v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v180)+16)) = uint16(v182)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v124)))
	*(*int32)(unsafe.Add(mBase, uint32(v180)+12)) = v184
	v186 = F_heap_copytuple(m, l2)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = v186
	goto L12
L39:
	;
	goto L13
L40:
	;
	if v223&int32(1) == int32(0) {
		goto L12
	} else {
		goto L98
	}
L41:
	;
	F_raw_heap_insert(m, l0, v222)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L43
	}
L42:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v438 = F_hash_search(m, v432, v20+int32(8), int32(1), v20+int32(7))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L1
	} else {
		goto L97
	}
L43:
	;
	v240 = v20 + int32(4)
	v242 = v222 + int32(8)
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242))))
	*(*uint16)(unsafe.Add(mBase, uint32(v240))) = uint16(v243)
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v245
	v248 = v20 + int32(76)
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v242))))
	*(*uint16)(unsafe.Add(mBase, uint32(v248))) = uint16(v249)
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v222)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v251
	v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
	if v253 != int32(1) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v358)+20)))
	if v359&int32(8192) == int32(0) {
		goto L40
	} else {
		goto L76
	}
L45:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v257)+20)))
	v259 = int32(768)
	if v258&v259 != v259 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
	v264 = v263
	goto L48
L47:
	;
	v264 = int32(2)
	goto L48
L48:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v258&int32(6272) == int32(4096) {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v274 = int32(0)
	if base.Ui32(int32(3)) <= base.Ui32(v264) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	v270 = F_HeapTupleGetUpdateXid(m, v257)
	mBase = m.M
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
	v273 = v272
	goto L49
L53:
	;
	v273 = v270
	goto L49
L54:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v265))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v264)) == int32(0) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v292 = v274
	goto L56
L56:
	;
	if base.Ui32(v273) < base.Ui32(int32(3)) {
		v317 = v274
		goto L61
	} else {
		goto L62
	}
L57:
	;
	v292 = v289 ^ int32(1)
	goto L56
L58:
	;
	v289 = base.B2i32(base.Ui32(v264) < base.Ui32(v265))
	goto L57
L59:
	;
	goto L60
L60:
	;
	v289 = int32(base.Ui32(v264-v265) >> (uint(int32(31)) % 32))
	goto L57
L61:
	;
	if v292|v317 != int32(1) {
		goto L44
	} else {
		goto L69
	}
L62:
	;
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+20)))
	if v296&int32(128) != 0 {
		v317 = v274
		goto L61
	} else {
		goto L63
	}
L63:
	;
	if v296&int32(4176) == int32(64) {
		v317 = v274
		goto L61
	} else {
		goto L64
	}
L64:
	;
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v265))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v273)) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L65:
	;
	v317 = v314 ^ int32(1)
	goto L61
L66:
	;
	v314 = base.B2i32(base.Ui32(v273) < base.Ui32(v265))
	goto L65
L67:
	;
	goto L68
L68:
	;
	v314 = int32(base.Ui32(v273-v265) >> (uint(int32(31)) % 32))
	goto L65
L69:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v322)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v323
	v325 = *(*int64)(unsafe.Add(mBase, uint32(v322)))
	*(*int64)(unsafe.Add(mBase, uint32(v20)+32)) = v325
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v218)+4)) = uint16(v327)
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v218))) = v329
	v331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v332 = *(*int64)(unsafe.Add(mBase, uint32(v331)))
	*(*int64)(unsafe.Add(mBase, uint32(v216))) = v332
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v331)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v216)+8)) = v334
	v336 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248))))
	*(*uint16)(unsafe.Add(mBase, uint32(v214)+4)) = uint16(v336)
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v20)+72))
	*(*int32)(unsafe.Add(mBase, uint32(v214))) = v338
	if v292 != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	F_logical_rewrite_log_mapping(m, l0, v264, v20+int32(32))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if v317&base.B2i32(v264 != v273) == int32(0) {
		goto L44
	} else {
		goto L74
	}
L73:
	;
	goto L72
L74:
	;
	F_logical_rewrite_log_mapping(m, l0, v273, v20+int32(32))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	goto L44
L76:
	;
	v364 = int32(768)
	if v359&v364 != v364 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v358)))
	v370 = v368
	goto L79
L78:
	;
	v370 = int32(2)
	goto L79
L79:
	;
	v371 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if base.B2i32(base.Ui32(int32(2)) < base.Ui32(v371))&base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v370)) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v383 != 0 {
		goto L40
	} else {
		goto L84
	}
L81:
	;
	v383 = base.B2i32(base.Ui32(v370) < base.Ui32(v371))
	goto L80
L82:
	;
	goto L83
L83:
	;
	v383 = int32(base.Ui32(v370-v371) >> (uint(int32(31)) % 32))
	goto L80
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = int64(0)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v222)+16))
	v389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v388)+20)))
	v390 = int32(768)
	if v389&v390 != v390 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v388)))
	v396 = v394
	goto L87
L86:
	;
	v396 = int32(2)
	goto L87
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v396
	v398 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)))
	*(*uint16)(unsafe.Add(mBase, uint32(v212)+4)) = uint16(v398)
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v20)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v212))) = v400
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v405 = int32(0)
	v407 = F_hash_search(m, v402, v20+int32(8), v405, v405)
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L88
	}
L88:
	;
	if v407 != 0 {
		goto L89
	} else {
		goto L90
	}
L89:
	;
	if v223&int32(1) != 0 {
		goto L92
	} else {
		goto L93
	}
L90:
	;
	goto L91
L91:
	;
	goto L42
L92:
	;
	F_pfree(m, v222)
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L1
	} else {
		goto L95
	}
L93:
	;
	goto L94
L94:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v407)+20))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v407)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v414
	v416 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v407)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+28)) = uint16(v416)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v413)+16))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v418)+12)) = v419
	v421 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240))))
	*(*uint16)(unsafe.Add(mBase, uint32(v418)+16)) = uint16(v421)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v429 = F_hash_search(m, v423, v20+int32(8), int32(2), v20+int32(7))
	mBase = m.M
	v430 = m.ExcPending
	if v430 != 0 {
		goto L1
	} else {
		goto L96
	}
L95:
	;
	goto L94
L96:
	;
	v222 = v413
	v223 = int32(1)
	goto L41
L97:
	;
	v440 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v438)+16)) = uint16(v440)
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	*(*int32)(unsafe.Add(mBase, uint32(v438)+12)) = v442
	goto L40
L98:
	;
	F_pfree(m, v222)
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L1
	} else {
		goto L99
	}
L99:
	;
	goto L12
}
