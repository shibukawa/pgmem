package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ER_flatten_into(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v5&int32(17) == int32(1) {
		if l2 != 0 {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
			base.MemoryCopy(m, l1, v11, l2)
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(l1))) = l2 << (uint(int32(2)) % 32)
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v16
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v18
		return
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v20 == int32(0) {
			v23 = F_expanded_record_fetch_tupdesc(m, l0)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				v25 = v23
				if l2 != 0 {
					base.MemoryFill(m, l1, int32(0), l2)
				} else {
				}
				*(*int32)(unsafe.Add(mBase, uint32(l1))) = l2 << (uint(int32(2)) % 32)
				v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v31
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				v34 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v34)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v33
				v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
				v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)))
				v43 = v39 | v40&int32(_a_F_ER_flatten_into_0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v43)
				v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+22)) = uint8(v45)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
				v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
				if v57 != 0 {
					v58 = l1 + int32(23)
				} else {
					v58 = v34
				}
				F_heap_fill_tuple(m, v25, v48, v49, l1+v50, l1+int32(20), v58)
				mBase = m.M
				v60 = m.ExcPending
				if v60 != 0 {
					return
				} else {
					return
				}
			}
		} else {
			v25 = v20
			if l2 != 0 {
				base.MemoryFill(m, l1, int32(0), l2)
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(l1))) = l2 << (uint(int32(2)) % 32)
			v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v31
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
			v34 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)) = uint16(v34)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v33
			v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25))))
			v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)))
			v43 = v39 | v40&int32(_a_F_ER_flatten_into_0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+18)) = uint16(v43)
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+22)) = uint8(v45)
			v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
			v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+80)))
			if v57 != 0 {
				v58 = l1 + int32(23)
			} else {
				v58 = v34
			}
			F_heap_fill_tuple(m, v25, v48, v49, l1+v50, l1+int32(20), v58)
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_EstimateSubplanHashTableSpace(m *base.Module, l0 float64, l1 int32, l2 int32) int32 {
	var v9 int32
	_ = v9
	var v11 float64
	_ = v11
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v29 int64
	_ = v29
	var v31 int64
	_ = v31
	var v51 float64
	_ = v51
	var v57 int32
	_ = v57
	var v62 float64
	_ = v62
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v72 int32
	_ = v72
	var v74 float64
	_ = v74
	var v77 int64
	_ = v77
	var v78 int64
	_ = v78
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v92 int64
	_ = v92
	var v94 int64
	_ = v94
	var v114 float64
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	v9 = int32(-1)
	v11 = base.F64_div(l0, float64(0.9))
	if base.F64_ge(v11, float64(4.294967296e+09)) != 0 {
		v57 = v9
	} else {
		v14 = int64(2)
		v15 = base.I64_trunc_sat_f64_u(v11)
		if base.Ui64(v15) <= base.Ui64(v14) {
			v18 = v14
		} else {
			v18 = v15
		}
		v19 = int64(1)
		if v18&(v18-v19) == int64(0) {
			v29 = v18
		} else {
			v29 = v19 << (uint(int64(64)-base.I64_clz(v18)) % 64)
		}
		v31 = v29 * int64(12)
		if base.Ui64(int64(2147483646)) < base.Ui64(v31) {
			v57 = v9
		} else {
			v51 = base.F64_add(base.F64_mul(l0, base.F64_convert_i32_u((l1+int32(7))&int32(-8)+int32(16))), base.F64_convert_i32_u(base.I32_wrap_i64(v31)+int32(32)))
			if base.F64_ge(v51, float64(4.294967295e+09)) != 0 {
				v57 = v9
			} else {
				v57 = base.I32_trunc_sat_f64_u(v51)
			}
		}
	}
	if l2|base.B2i32(v57 == int32(-1)) != 0 {
		v126 = v57
	} else {
		v62 = float64(1)
		v64 = base.F64_mul(l0, float64(0.0625))
		if base.F64_lt(v64, v62) != 0 {
			v67 = v62
		} else {
			v67 = v64
		}
		v72 = int32(-1)
		v74 = base.F64_div(v67, float64(0.9))
		if base.F64_ge(v74, float64(4.294967296e+09)) != 0 {
			v120 = v72
		} else {
			v77 = int64(2)
			v78 = base.I64_trunc_sat_f64_u(v74)
			if base.Ui64(v78) <= base.Ui64(v77) {
				v81 = v77
			} else {
				v81 = v78
			}
			v82 = int64(1)
			if v81&(v81-v82) == int64(0) {
				v92 = v81
			} else {
				v92 = v82 << (uint(int64(64)-base.I64_clz(v81)) % 64)
			}
			v94 = v92 * int64(12)
			if base.Ui64(int64(2147483646)) < base.Ui64(v94) {
				v120 = v72
			} else {
				v114 = base.F64_add(base.F64_mul(v67, base.F64_convert_i32_u((l1+int32(7))&int32(-8)+int32(16))), base.F64_convert_i32_u(base.I32_wrap_i64(v94)+int32(32)))
				if base.F64_ge(v114, float64(4.294967295e+09)) != 0 {
					v120 = v72
				} else {
					v120 = base.I32_trunc_sat_f64_u(v114)
				}
			}
		}
		v121 = v120 + v57
		if base.Ui32(v121) < base.Ui32(v57) {
			v123 = int32(-1)
		} else {
			v123 = v121
		}
		v126 = v123
	}
	return v126
}
func F_EvictUnpinnedBufferInternal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v5 int32
	_ = v5
	var v13 int64
	_ = v13
	var v22 int64
	_ = v22
	var v24 int32
	_ = v24
	var v25 int64
	_ = v25
	var v31 int64
	_ = v31
	var v35 int64
	_ = v35
	var v45 int64
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v125 int64
	_ = v125
	v3 = int64(0)
	v5 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v5)
	v13 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v3, v3)
	if v13&int64(16777216) == v3 {
		v125 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
		return int32(0)
	} else {
		if v13&int64(262143) != int64(0) {
			v125 = base.AtomicRmwSub64(m, l0, int32(24), int64(4194304))
			return int32(0)
		} else {
			v22 = int64(0)
			v24 = int32(24)
			v25 = base.AtomicRmwCmpxchg64(m, l0, v24, v22, v22)
			v31 = base.AtomicRmwCmpxchg64(m, l0, v24, v25, v25&int64(-4194305)+int64(1))
			if v25 != v31 {
				v35 = v31
				for {
					v45 = base.AtomicRmwCmpxchg64(m, l0, int32(24), v35, v35&int64(-4194305)+int64(1))
					if v35 != v45 {
						v35 = v45
						continue
					} else {
						break
					}
					break
				}
			} else {
			}
			v54 = int32(_a_F_EvictUnpinnedBufferInternal_0)
			v55 = *(*int32)(unsafe.Add(mBase, _c_F_EvictUnpinnedBufferInternal[0]))
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v61 = int32(1)
			v62 = v60 + v61
			*(*int32)(unsafe.Add(mBase, uint32(v55<<(uint(int32(2))%32))+uint32(_c_F_EvictUnpinnedBufferInternal[1]))) = v62
			v65 = v55 << (uint(int32(4)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_c_F_EvictUnpinnedBufferInternal[2]))) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_c_F_EvictUnpinnedBufferInternal[3]))) = v62
			*(*int32)(unsafe.Add(mBase, _c_F_EvictUnpinnedBufferInternal[0])) = int32(-1)
			*(*int32)(unsafe.Add(mBase, _c_F_EvictUnpinnedBufferInternal[4])) = v55
			*(*int32)(unsafe.Add(mBase, uint32(v65)+uint32(_c_F_EvictUnpinnedBufferInternal[5]))) = v61
			v83 = *(*int32)(unsafe.Add(mBase, _c_F_EvictUnpinnedBufferInternal[6]))
			F_ResourceOwnerRemember(m, v83, base.I64_extend_i32_s(v62), int32(_a_F_EvictUnpinnedBufferInternal_1))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				if v13&int64(8388608) != int64(0) {
					v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v96 = v94 + int32(1)
					F_BufferLockAcquire(m, v96, l0, int32(2))
					mBase = m.M
					v99 = m.ExcPending
					if v99 != 0 {
						return int32(0)
					} else {
						F_FlushBuffer(m, l0, int32(0), int32(3))
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							F_BufferLockUnlock(m, v96, l0)
							mBase = m.M
							v105 = m.ExcPending
							if v105 != 0 {
								return int32(0)
							} else {
								v106 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l1))) = uint8(v106)
								v109 = F_InvalidateVictimBuffer(m, l0)
								mBase = m.M
								v110 = m.ExcPending
								if v110 != 0 {
									return int32(0)
								} else {
									v112 = *(*int32)(unsafe.Add(mBase, _c_F_EvictUnpinnedBufferInternal[6]))
									v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									F_ResourceOwnerForget(m, v112, base.I64_extend_i32_s(v113+int32(1)), int32(_a_F_EvictUnpinnedBufferInternal_1))
									mBase = m.M
									v119 = m.ExcPending
									if v119 != 0 {
										return int32(0)
									} else {
										F_UnpinBufferNoOwner(m, l0)
										mBase = m.M
										v121 = m.ExcPending
										if v121 != 0 {
											return int32(0)
										} else {
											return v109
										}
									}
								}
							}
						}
					}
				} else {
					v109 = F_InvalidateVictimBuffer(m, l0)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						v112 = *(*int32)(unsafe.Add(mBase, _c_F_EvictUnpinnedBufferInternal[6]))
						v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						F_ResourceOwnerForget(m, v112, base.I64_extend_i32_s(v113+int32(1)), int32(_a_F_EvictUnpinnedBufferInternal_1))
						mBase = m.M
						v119 = m.ExcPending
						if v119 != 0 {
							return int32(0)
						} else {
							F_UnpinBufferNoOwner(m, l0)
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								return v109
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecASInsertTriggers(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+52))
	if v5 == int32(0) {
		return
	} else {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+12)))
		if v8 != int32(1) {
			return
		} else {
			v11 = int32(0)
			F_AfterTriggerSaveEvent(m, l0, l1, v11, v11, v11, v11, v11, v11, v11, v11, l2, v11)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return
			} else {
				return
			}
		}
	}
}
func F_ExecAppend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
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
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
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
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v96 int64
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v118 int32
	_ = v118
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int64
	_ = v132
	var v135 int64
	_ = v135
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v270 int32
	_ = v270
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v359 int32
	_ = v359
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v388 int32
	_ = v388
	var v391 int32
	_ = v391
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v439 int32
	_ = v439
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v545 int32
	_ = v545
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v564 int32
	_ = v564
	var v566 int32
	_ = v566
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v588 int32
	_ = v588
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v622 int32
	_ = v622
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v629 int32
	_ = v629
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)))
	if v4 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v658 = v655 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v658
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v660+v658<<(uint(int32(2))%32))))
	return v664
L2:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v648)+8))
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v649)+12))
	m.T0[v650].(func(*base.Module, int32))(m, v648)
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L13
	} else {
		goto L197
	}
L3:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v7 == int32(0) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	goto L84
L6:
	;
	goto L2
L7:
	;
	goto L8
L8:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	if v10 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v276 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v277 = m.T0[v276].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L13
	} else {
		goto L80
	}
L10:
	;
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)))
	if v13 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v91 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)) = uint8(base.B2i32(v90 == v91))
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	v96 = int64(0)
	if v94 == v91 {
		goto L37
	} else {
		goto L38
	}
L12:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+168))
	v15 = int32(0)
	v17 = F_ExecFindMatchingSubPlans(m, v14, v15, v15)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	return int32(0)
L14:
	;
	v21 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+172)) = uint8(v21)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v17
	if v17 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	goto L11
L16:
	;
	goto L17
L17:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v29 = int32(0)
	if base.B2i32(v17 == v29)|base.B2i32(v28 == v29) != 0 {
		v74 = v29
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v74 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L19:
	;
	goto L18
L20:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v39 < v40 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v42 = v39
	goto L23
L22:
	;
	v42 = v40
	goto L23
L23:
	;
	if v42 <= int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v45 = int32(1)
	goto L26
L25:
	;
	v45 = v42
	goto L26
L26:
	;
	v46 = int32(8)
	v51 = int32(0)
	goto L27
L27:
	;
	v58 = v51 << (uint(int32(2)) % 32)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v28+v46+v58)))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v17+v46+v58)))
	v63 = v60 & v62
	v65 = base.B2i32(v63 != int32(0))
	if v63 != 0 {
		v74 = v65
		goto L19
	} else {
		goto L29
	}
L28:
	;
	v74 = v65
	goto L19
L29:
	;
	v67 = v51 + int32(1)
	if v67 != v45 {
		v51 = v67
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = int32(0)
	goto L11
L32:
	;
	goto L33
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v81 = F_bms_intersect(m, v79, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L34
	}
L34:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+176))
	v84 = F_bms_del_members(m, v83, v81)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L13
	} else {
		goto L35
	}
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+180)) = v81
	*(*int32)(unsafe.Add(mBase, uint32(l0)+176)) = v84
	goto L11
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+144)) = v140
	if v140 == int32(0) {
		goto L9
	} else {
		goto L51
	}
L37:
	;
	v140 = int32(0)
	goto L36
L38:
	;
	goto L39
L39:
	;
	v101 = v94 + int32(8)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
	if v102 == int32(1) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v101)))
	v140 = base.I32_popcnt(v105)
	goto L36
L41:
	;
	goto L42
L42:
	;
	v108 = v102 << (uint(int32(2)) % 32)
	if v108 <= int32(7) {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	v140 = base.I32_wrap_i64(v135)
	goto L36
L44:
	;
	if v108 == int32(0) {
		v135 = v96
		goto L43
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	v132 = F_pg_popcount_optimized(m, v101, v108)
	mBase = m.M
	v135 = v132
	goto L43
L47:
	;
	v113 = v108
	v114 = v101
	v115 = v96
	goto L48
L48:
	;
	v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+3)))
	v117 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v116)+uint32(_c_F_ExecAppend[0]))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+2)))
	v119 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v118)+uint32(_c_F_ExecAppend[0]))))
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	v121 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v120)+uint32(_c_F_ExecAppend[0]))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	v123 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v122)+uint32(_c_F_ExecAppend[0]))))
	v127 = v117 + (v119 + (v121 + (v115 + v123)))
	v128 = int32(4)
	v131 = v113 - v128
	if v131 != 0 {
		v113 = v131
		v114 = v114 + v128
		v115 = v127
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v135 = v127
	goto L43
L50:
	;
	goto L49
L51:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v144 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	if v201 < int32(0) {
		goto L9
	} else {
		goto L63
	}
L53:
	;
	v201 = base.I32_ctz(v187) | v188<<(uint(int32(5))%32)
	goto L52
L54:
	;
	v201 = int32(-2)
	goto L52
L55:
	;
	v152 = int32(0)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v144)+4))
	if v155 <= v152 {
		goto L54
	} else {
		goto L56
	}
L56:
	;
	v158 = v144 + int32(8)
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
	v165 = v162 & int32(-1)
	if v165 != 0 {
		v187 = v165
		v188 = v152
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v166 = int32(1)
	if v166 == v155 {
		goto L54
	} else {
		goto L58
	}
L58:
	;
	v170 = v166
	goto L59
L59:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v158+v170<<(uint(int32(2))%32))))
	if v177 != 0 {
		v187 = v177
		v188 = v170
		goto L53
	} else {
		goto L61
	}
L60:
	;
	goto L54
L61:
	;
	v179 = v170 + int32(1)
	if v179 != v155 {
		v170 = v179
		goto L59
	} else {
		goto L62
	}
L62:
	;
	goto L60
L63:
	;
	v205 = v201
	goto L64
L64:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207+v205<<(uint(int32(2))%32))))
	F_ExecAsyncRequest(m, v211)
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L13
	} else {
		goto L66
	}
L65:
	;
	goto L9
L66:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v214 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	if int32(0) <= v270 {
		v205 = v270
		goto L64
	} else {
		goto L78
	}
L68:
	;
	v270 = base.I32_ctz(v256) | v257<<(uint(int32(5))%32)
	goto L67
L69:
	;
	v270 = int32(-2)
	goto L67
L70:
	;
	v221 = v205 + int32(1)
	v223 = int32(base.Ui32(v221) >> (uint(int32(5)) % 32))
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v214)+4))
	if v224 <= v223 {
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v227 = v214 + int32(8)
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227+v223<<(uint(int32(2))%32))))
	v234 = v231 & (int32(-1) << (uint(v221) % 32))
	if v234 != 0 {
		v256 = v234
		v257 = v223
		goto L68
	} else {
		goto L72
	}
L72:
	;
	v236 = v223 + int32(1)
	if v236 == v224 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v239 = v236
	goto L74
L74:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v227+v239<<(uint(int32(2))%32))))
	if v246 != 0 {
		v256 = v246
		v257 = v239
		goto L68
	} else {
		goto L76
	}
L75:
	;
	goto L69
L76:
	;
	v248 = v239 + int32(1)
	if v248 != v224 {
		v239 = v248
		goto L74
	} else {
		goto L77
	}
L77:
	;
	goto L75
L78:
	;
	goto L65
L79:
	;
	v280 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+116)) = uint8(v280)
	goto L5
L80:
	;
	if v277 != 0 {
		goto L79
	} else {
		goto L81
	}
L81:
	;
	v279 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v279 != 0 {
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L2
L83:
	;
	return v643
L84:
	;
	v289 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAppend[1]))
	if v289 != 0 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v639 = *(*int32)(unsafe.Add(mBase, uint32(v638)+8))
	v640 = *(*int32)(unsafe.Add(mBase, uint32(v639)+12))
	m.T0[v640].(func(*base.Module, int32))(m, v638)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L13
	} else {
		goto L196
	}
L86:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L13
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	v293 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v293 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L89:
	;
	goto L88
L90:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+112))
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v612+v613<<(uint(int32(2))%32))))
	v618 = *(*int32)(unsafe.Add(mBase, uint32(v617)+52))
	if v618 != 0 {
		goto L180
	} else {
		goto L181
	}
L91:
	;
	goto L132
L92:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v298 <= int32(0) {
		goto L98
	} else {
		goto L99
	}
L93:
	;
	if v292 != 0 {
		goto L92
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	if v292 == int32(0) {
		goto L91
	} else {
		goto L97
	}
L96:
	;
	goto L90
L97:
	;
	goto L92
L98:
	;
	v301 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v301
	if v292 == v301 {
		goto L103
	} else {
		goto L104
	}
L99:
	;
	v439 = v298
	goto L100
L100:
	;
	v655 = v439
	goto L1
L101:
	;
	if int32(0) <= v359 {
		goto L112
	} else {
		goto L113
	}
L102:
	;
	v359 = base.I32_ctz(v345) | v346<<(uint(int32(5))%32)
	goto L101
L103:
	;
	v359 = int32(-2)
	goto L101
L104:
	;
	v310 = int32(0)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v313 <= v310 {
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v316 = v292 + int32(8)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v316)))
	v323 = v320 & int32(-1)
	if v323 != 0 {
		v345 = v323
		v346 = v310
		goto L102
	} else {
		goto L106
	}
L106:
	;
	v324 = int32(1)
	if v324 == v313 {
		goto L103
	} else {
		goto L107
	}
L107:
	;
	v328 = v324
	goto L108
L108:
	;
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v316+v328<<(uint(int32(2))%32))))
	if v335 != 0 {
		v345 = v335
		v346 = v328
		goto L102
	} else {
		goto L110
	}
L109:
	;
	goto L103
L110:
	;
	v337 = v328 + int32(1)
	if v337 != v313 {
		v328 = v337
		goto L108
	} else {
		goto L111
	}
L111:
	;
	goto L109
L112:
	;
	v363 = v359
	goto L115
L113:
	;
	goto L114
L114:
	;
	F_bms_free(m, v292)
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L13
	} else {
		goto L130
	}
L115:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365+v363<<(uint(int32(2))%32))))
	F_ExecAsyncRequest(m, v369)
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L13
	} else {
		goto L117
	}
L116:
	;
	goto L114
L117:
	;
	if v292 == int32(0) {
		goto L120
	} else {
		goto L121
	}
L118:
	;
	if int32(0) <= v427 {
		v363 = v427
		goto L115
	} else {
		goto L129
	}
L119:
	;
	v427 = base.I32_ctz(v413) | v414<<(uint(int32(5))%32)
	goto L118
L120:
	;
	v427 = int32(-2)
	goto L118
L121:
	;
	v378 = v363 + int32(1)
	v380 = int32(base.Ui32(v378) >> (uint(int32(5)) % 32))
	v381 = *(*int32)(unsafe.Add(mBase, uint32(v292)+4))
	if v381 <= v380 {
		goto L120
	} else {
		goto L122
	}
L122:
	;
	v384 = v292 + int32(8)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v384+v380<<(uint(int32(2))%32))))
	v391 = v388 & (int32(-1) << (uint(v378) % 32))
	if v391 != 0 {
		v413 = v391
		v414 = v380
		goto L119
	} else {
		goto L123
	}
L123:
	;
	v393 = v380 + int32(1)
	if v393 == v381 {
		goto L120
	} else {
		goto L124
	}
L124:
	;
	v396 = v393
	goto L125
L125:
	;
	v403 = *(*int32)(unsafe.Add(mBase, uint32(v384+v396<<(uint(int32(2))%32))))
	if v403 != 0 {
		v413 = v403
		v414 = v396
		goto L119
	} else {
		goto L127
	}
L126:
	;
	goto L120
L127:
	;
	v405 = v396 + int32(1)
	if v405 != v381 {
		v396 = v405
		goto L125
	} else {
		goto L128
	}
L128:
	;
	goto L126
L129:
	;
	goto L116
L130:
	;
	v435 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v435 <= int32(0) {
		goto L91
	} else {
		goto L131
	}
L131:
	;
	v439 = v435
	goto L100
L132:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if int32(0) < v447 {
		goto L135
	} else {
		goto L136
	}
L133:
	;
	goto L90
L134:
	;
	v608 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v608 != 0 {
		goto L132
	} else {
		goto L179
	}
L135:
	;
	v451 = *(*int32)(unsafe.Add(mBase, _c_F_ExecAppend[1]))
	if v451 != 0 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	goto L137
L137:
	;
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+140)))
	if v602 != int32(1) {
		goto L90
	} else {
		goto L178
	}
L138:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L13
	} else {
		goto L141
	}
L139:
	;
	goto L140
L140:
	;
	F_ExecAppendAsyncEventWait(m, l0)
	mBase = m.M
	v455 = m.ExcPending
	if v455 != 0 {
		goto L13
	} else {
		goto L142
	}
L141:
	;
	goto L140
L142:
	;
	v456 = *(*int32)(unsafe.Add(mBase, uint32(l0)+148))
	if v456 == int32(0) {
		goto L134
	} else {
		goto L143
	}
L143:
	;
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v459 <= int32(0) {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v462 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+148)) = v462
	if v456 == v462 {
		goto L149
	} else {
		goto L150
	}
L145:
	;
	v600 = v459
	goto L146
L146:
	;
	v655 = v600
	goto L1
L147:
	;
	if int32(0) <= v520 {
		goto L158
	} else {
		goto L159
	}
L148:
	;
	v520 = base.I32_ctz(v506) | v507<<(uint(int32(5))%32)
	goto L147
L149:
	;
	v520 = int32(-2)
	goto L147
L150:
	;
	v471 = int32(0)
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	if v474 <= v471 {
		goto L149
	} else {
		goto L151
	}
L151:
	;
	v477 = v456 + int32(8)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v477)))
	v484 = v481 & int32(-1)
	if v484 != 0 {
		v506 = v484
		v507 = v471
		goto L148
	} else {
		goto L152
	}
L152:
	;
	v485 = int32(1)
	if v485 == v474 {
		goto L149
	} else {
		goto L153
	}
L153:
	;
	v489 = v485
	goto L154
L154:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v477+v489<<(uint(int32(2))%32))))
	if v496 != 0 {
		v506 = v496
		v507 = v489
		goto L148
	} else {
		goto L156
	}
L155:
	;
	goto L149
L156:
	;
	v498 = v489 + int32(1)
	if v498 != v474 {
		v489 = v498
		goto L154
	} else {
		goto L157
	}
L157:
	;
	goto L155
L158:
	;
	v524 = v520
	goto L161
L159:
	;
	goto L160
L160:
	;
	F_bms_free(m, v456)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L13
	} else {
		goto L176
	}
L161:
	;
	v526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v526+v524<<(uint(int32(2))%32))))
	F_ExecAsyncRequest(m, v530)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L13
	} else {
		goto L163
	}
L162:
	;
	goto L160
L163:
	;
	if v456 == int32(0) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	if int32(0) <= v588 {
		v524 = v588
		goto L161
	} else {
		goto L175
	}
L165:
	;
	v588 = base.I32_ctz(v574) | v575<<(uint(int32(5))%32)
	goto L164
L166:
	;
	v588 = int32(-2)
	goto L164
L167:
	;
	v539 = v524 + int32(1)
	v541 = int32(base.Ui32(v539) >> (uint(int32(5)) % 32))
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v456)+4))
	if v542 <= v541 {
		goto L166
	} else {
		goto L168
	}
L168:
	;
	v545 = v456 + int32(8)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v545+v541<<(uint(int32(2))%32))))
	v552 = v549 & (int32(-1) << (uint(v539) % 32))
	if v552 != 0 {
		v574 = v552
		v575 = v541
		goto L165
	} else {
		goto L169
	}
L169:
	;
	v554 = v541 + int32(1)
	if v554 == v542 {
		goto L166
	} else {
		goto L170
	}
L170:
	;
	v557 = v554
	goto L171
L171:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(v545+v557<<(uint(int32(2))%32))))
	if v564 != 0 {
		v574 = v564
		v575 = v557
		goto L165
	} else {
		goto L173
	}
L172:
	;
	goto L166
L173:
	;
	v566 = v557 + int32(1)
	if v566 != v542 {
		v557 = v566
		goto L171
	} else {
		goto L174
	}
L174:
	;
	goto L172
L175:
	;
	goto L162
L176:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v596 <= int32(0) {
		goto L134
	} else {
		goto L177
	}
L177:
	;
	v600 = v596
	goto L146
L178:
	;
	goto L2
L179:
	;
	goto L133
L180:
	;
	F_ExecReScan(m, v617)
	mBase = m.M
	v620 = m.ExcPending
	if v620 != 0 {
		goto L13
	} else {
		goto L183
	}
L181:
	;
	goto L182
L182:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v617)+12))
	v622 = m.T0[v621].(func(*base.Module, int32) int32)(m, v617)
	mBase = m.M
	v623 = m.ExcPending
	if v623 != 0 {
		goto L13
	} else {
		goto L184
	}
L183:
	;
	goto L182
L184:
	;
	if v622 != 0 {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v622)+4)))
	if v624&int32(2) == int32(0) {
		v643 = v622
		goto L83
	} else {
		goto L188
	}
L186:
	;
	goto L187
L187:
	;
	v629 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if int32(0) < v629 {
		goto L189
	} else {
		goto L190
	}
L188:
	;
	goto L187
L189:
	;
	F_ExecAppendAsyncEventWait(m, l0)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L13
	} else {
		goto L192
	}
L190:
	;
	goto L191
L191:
	;
	v634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+184))
	v635 = m.T0[v634].(func(*base.Module, int32) int32)(m, l0)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L13
	} else {
		goto L193
	}
L192:
	;
	goto L191
L193:
	;
	if v635 != 0 {
		goto L84
	} else {
		goto L194
	}
L194:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+144))
	if v637 != 0 {
		goto L84
	} else {
		goto L195
	}
L195:
	;
	goto L85
L196:
	;
	v643 = v638
	goto L83
L197:
	;
	return v648
}
func F_ExecCheck(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l0 != 0 {
		v11 = int32(_a_F_ExecCheck_0)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheck[0]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		*(*int32)(unsafe.Add(mBase, _c_F_ExecCheck[0])) = v14
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v19 = m.T0[v18].(func(*base.Module, int32, int32, int32) int64)(m, l0, l1, v8+int32(15))
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_ExecCheck[0])) = v12
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
			v29 = v25 | base.B2i32(v19 != int64(0))
			m.G0 = v8 + int32(16)
			return v29 & int32(1)
		}
	} else {
		v29 = int32(1)
		m.G0 = v8 + int32(16)
		return v29 & int32(1)
	}
}
func F_ExecCheckOneRelPerms(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int64
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v183 int32
	_ = v183
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	v7 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v9 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F_ExecCheckOneRelPerms[0]))
	v14 = v13
	goto L3
L2:
	;
	v14 = v9
	goto L3
L3:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = F_pg_class_aclmask(m, v15, v14, v7)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	return v214
L5:
	;
	return int32(0)
L6:
	;
	v22 = v7 & (v16 ^ int64(-1))
	if v22 == int64(0) {
		v214 = int32(1)
		goto L4
	} else {
		goto L7
	}
L7:
	;
	if base.Ui64(int64(7)) < base.Ui64(v22) {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	goto L10
L10:
	;
	if v22&int64(2) == int64(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v22&int64(1) == int64(0) {
		goto L54
	} else {
		goto L55
	}
L12:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v33 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v41 = v33
	goto L15
L14:
	;
	v36 = F_pg_attribute_aclcheck_all(m, v15, v14, int64(2), int32(1))
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L5
	} else {
		goto L16
	}
L15:
	;
	if v41 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L16:
	;
	if v36 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	return int32(0)
L18:
	;
	goto L19
L19:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v41 = v40
	goto L15
L20:
	;
	if v98 < int32(0) {
		goto L11
	} else {
		goto L31
	}
L21:
	;
	v98 = base.I32_ctz(v84) | v85<<(uint(int32(5))%32)
	goto L20
L22:
	;
	v98 = int32(-2)
	goto L20
L23:
	;
	v49 = int32(0)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	if v52 <= v49 {
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v55 = v41 + int32(8)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v62 = v59 & int32(-1)
	if v62 != 0 {
		v84 = v62
		v85 = v49
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v63 = int32(1)
	if v63 == v52 {
		goto L22
	} else {
		goto L26
	}
L26:
	;
	v67 = v63
	goto L27
L27:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v55+v67<<(uint(int32(2))%32))))
	if v74 != 0 {
		v84 = v74
		v85 = v67
		goto L21
	} else {
		goto L29
	}
L28:
	;
	goto L22
L29:
	;
	v76 = v67 + int32(1)
	if v76 != v52 {
		v67 = v76
		goto L27
	} else {
		goto L30
	}
L30:
	;
	goto L28
L31:
	;
	v105 = v98
	goto L32
L32:
	;
	v108 = v105 - int32(7)
	if v108&int32(_a_F_ExecCheckOneRelPerms_0) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L11
L34:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v127 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L35:
	;
	v113 = int32(0)
	v116 = F_pg_attribute_aclcheck_all(m, v15, v14, int64(2), v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L5
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v120 = F_pg_attribute_aclcheck(m, v15, base.I32_extend16_s(v108), v14, int64(2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L5
	} else {
		goto L40
	}
L38:
	;
	if v116 != 0 {
		v214 = v113
		goto L4
	} else {
		goto L39
	}
L39:
	;
	goto L34
L40:
	;
	if v120 == int32(0) {
		goto L34
	} else {
		goto L41
	}
L41:
	;
	return int32(0)
L42:
	;
	if int32(0) <= v183 {
		v105 = v183
		goto L32
	} else {
		goto L53
	}
L43:
	;
	v183 = base.I32_ctz(v169) | v170<<(uint(int32(5))%32)
	goto L42
L44:
	;
	v183 = int32(-2)
	goto L42
L45:
	;
	v134 = v105 + int32(1)
	v136 = int32(base.Ui32(v134) >> (uint(int32(5)) % 32))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
	if v137 <= v136 {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v140 = v127 + int32(8)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(int32(2))%32))))
	v147 = v144 & (int32(-1) << (uint(v134) % 32))
	if v147 != 0 {
		v169 = v147
		v170 = v136
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v149 = v136 + int32(1)
	if v149 == v137 {
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v152 = v149
	goto L49
L49:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v140+v152<<(uint(int32(2))%32))))
	if v159 != 0 {
		v169 = v159
		v170 = v152
		goto L43
	} else {
		goto L51
	}
L50:
	;
	goto L44
L51:
	;
	v161 = v152 + int32(1)
	if v161 != v137 {
		v152 = v161
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L33
L54:
	;
	if base.Ui64(int64(4)) <= base.Ui64(v22) {
		goto L58
	} else {
		goto L59
	}
L55:
	;
	v196 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v198 = F_ExecCheckPermissionsModified(m, v15, v14, v196, int64(1))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	if v198 != 0 {
		goto L54
	} else {
		goto L57
	}
L57:
	;
	return int32(0)
L58:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v207 = F_ExecCheckPermissionsModified(m, v15, v14, v205, int64(4))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L5
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v214 = int32(1)
	goto L4
L61:
	;
	if v207 == int32(0) {
		v214 = int32(0)
		goto L4
	} else {
		goto L62
	}
L62:
	;
	goto L60
}
func F_ExecCleanTargetListLength(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
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
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	v2 = int32(0)
	if l0 == v2 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v13 <= int32(0) {
			return int32(0)
		} else {
			if v13 != int32(1) {
				v20 = int32(0)
				if v20 < v13 {
					v23 = v13
				} else {
					v23 = v20
				}
				v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v29 = int32(0)
				v32 = v29
				v33 = v29
				v34 = v2
				for {
					v39 = int32(2)
					v41 = v28 + v33<<(uint(v39)%32)
					v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+26)))
					v44 = int32(1)
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
					v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v47)+26)))
					v51 = v34 + (v43 ^ v44) + (v48 ^ v44)
					v53 = v33 + v39
					v55 = v32 + v39
					if v55 != v23&int32(2147483646) {
						v32 = v55
						v33 = v53
						v34 = v51
						continue
					} else {
						break
					}
					break
				}
				if v23&int32(1) == int32(0) {
					v79 = v51
				} else {
					v61 = v53
					v62 = v51
					v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v61<<(uint(int32(2))%32))))
					v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+26)))
					v79 = v62 + (v72 ^ int32(1))
				}
			} else {
				v61 = v2
				v62 = v2
				v67 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v67+v61<<(uint(int32(2))%32))))
				v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71)+26)))
				v79 = v62 + (v72 ^ int32(1))
			}
			return v79
		}
	}
}
func F_ExecConditionalAssignProjectionInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+44))
	if v12 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v15 = v13
	goto L3
L2:
	;
	v15 = int32(0)
	goto L3
L3:
	;
	if int32(0) < v10 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0+v120))) = v126
	return
L5:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	if v96 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L6:
	;
	v26 = v15
	v27 = int32(1)
	goto L9
L7:
	;
	v72 = v15
	goto L8
L8:
	;
	if v72 != 0 {
		goto L5
	} else {
		goto L26
	}
L9:
	;
	if v26 == int32(0) {
		goto L5
	} else {
		goto L11
	}
L10:
	;
	v72 = v66
	goto L8
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v36 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v39 != int32(6) {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	v42 = int32(*(*int16)(unsafe.Add(mBase, uint32(v36)+8)))
	if v27 != v42 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	v46 = l1 + v10<<(uint(int32(3))%32) - int32(72) + v27*int32(100)
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+91)))
	if v47 != 0 {
		goto L5
	} else {
		goto L15
	}
L15:
	;
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+88)))
	if v48 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v46)+68))
	if v49 != v50 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v36)+16))
	if v52 != int32(-1) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v46)+76))
	if v52 != v55 {
		goto L5
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v58 = v26 + int32(4)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
	if base.Ui32(v58) < base.Ui32(v60+v61<<(uint(int32(2))%32)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	goto L20
L22:
	;
	v66 = v58
	goto L24
L23:
	;
	v66 = int32(0)
	goto L24
L24:
	;
	if v27 != v10 {
		v26 = v66
		v27 = v27 + int32(1)
		goto L9
	} else {
		goto L25
	}
L25:
	;
	goto L10
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+68)) = int32(0)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+100)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)) = uint8(v81)
	v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+96)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)) = uint8(v83)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v120 = int32(92)
	v126 = v86
	goto L4
L27:
	;
	F_ExecInitResultSlot(m, l0, int32(_a_F_ExecConditionalAssignProjectionInfo_0))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	v111 = v96
	v112 = v12
	goto L29
L29:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	v115 = F_ExecBuildProjectionInfo(m, v112, v114, v111, l0, l1)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L30
	} else {
		goto L32
	}
L30:
	;
	return
L31:
	;
	v102 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+103)) = uint8(v102)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+99)) = uint8(v102)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+92)) = int32(_a_F_ExecConditionalAssignProjectionInfo_0)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v109)+44))
	v111 = v108
	v112 = v110
	goto L29
L32:
	;
	v120 = int32(68)
	v126 = v115
	goto L4
}
func F_ExecForceStoreHeapTuple(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v6 == int32(_a_F_ExecForceStoreHeapTuple_0) {
		v9 = F_ExecStoreHeapTuple(m, l0, l1, l2)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return
		} else {
			return
		}
	} else {
		if v6 == int32(_a_F_ExecForceStoreHeapTuple_1) {
			v13 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			if v13&int32(4) != 0 {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
				F_pfree(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					v21 = v19 & int32(-5)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v21)
					v23 = v21
					v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
					if v24 != 0 {
						F_ReleaseBuffer(m, v24)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return
						} else {
							v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
							v28 = v27
							v29 = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v29
							*(*int64)(unsafe.Add(mBase, uint32(l1)+44)) = int64(0)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v29)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
							v40 = v28 & int32(_a_F_ExecForceStoreHeapTuple_2)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v40)
							v42 = int32(_a_F_ExecForceStoreHeapTuple_3)
							v43 = *(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0]))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
							*(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0])) = v45
							v47 = F_heap_copytuple(m, l0)
							mBase = m.M
							v48 = m.ExcPending
							if v48 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v47
								v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
								v52 = v50 | int32(4)
								*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
								*(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0])) = v43
								if l2 != 0 {
									F_pfree(m, l0)
									mBase = m.M
									v80 = m.ExcPending
									if v80 != 0 {
										return
									} else {
										return
									}
								} else {
									return
								}
							}
						}
					} else {
						v28 = v23
						v29 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v29
						*(*int64)(unsafe.Add(mBase, uint32(l1)+44)) = int64(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v29)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
						v40 = v28 & int32(_a_F_ExecForceStoreHeapTuple_2)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v40)
						v42 = int32(_a_F_ExecForceStoreHeapTuple_3)
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0]))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0])) = v45
						v47 = F_heap_copytuple(m, l0)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v47
							v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
							v52 = v50 | int32(4)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
							*(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0])) = v43
							if l2 != 0 {
								F_pfree(m, l0)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									return
								}
							} else {
								return
							}
						}
					}
				}
			} else {
				v23 = v13
				v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
				if v24 != 0 {
					F_ReleaseBuffer(m, v24)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
						v28 = v27
						v29 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v29
						*(*int64)(unsafe.Add(mBase, uint32(l1)+44)) = int64(0)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v29)
						*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
						v40 = v28 & int32(_a_F_ExecForceStoreHeapTuple_2)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v40)
						v42 = int32(_a_F_ExecForceStoreHeapTuple_3)
						v43 = *(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0]))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
						*(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0])) = v45
						v47 = F_heap_copytuple(m, l0)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v47
							v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
							v52 = v50 | int32(4)
							*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
							*(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0])) = v43
							if l2 != 0 {
								F_pfree(m, l0)
								mBase = m.M
								v80 = m.ExcPending
								if v80 != 0 {
									return
								} else {
									return
								}
							} else {
								return
							}
						}
					}
				} else {
					v28 = v23
					v29 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v29
					*(*int64)(unsafe.Add(mBase, uint32(l1)+44)) = int64(0)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v29)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v29)
					v40 = v28 & int32(_a_F_ExecForceStoreHeapTuple_2)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v40)
					v42 = int32(_a_F_ExecForceStoreHeapTuple_3)
					v43 = *(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0]))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					*(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0])) = v45
					v47 = F_heap_copytuple(m, l0)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l1)+44)) = v47
						v50 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
						v52 = v50 | int32(4)
						*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v52)
						*(*int32)(unsafe.Add(mBase, _c_F_ExecForceStoreHeapTuple[0])) = v43
						if l2 != 0 {
							F_pfree(m, l0)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								return
							}
						} else {
							return
						}
					}
				}
			}
		} else {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
			m.T0[v56].(func(*base.Module, int32))(m, l1)
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
				F_heap_deform_tuple(m, l0, v59, v60, v61)
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return
				} else {
					v64 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					v66 = v64 & int32(_a_F_ExecForceStoreHeapTuple_2)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v66)
					v68 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v69)
					if l2 == int32(0) {
						return
					} else {
						v73 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
						v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+28))
						m.T0[v74].(func(*base.Module, int32))(m, l1)
						mBase = m.M
						v76 = m.ExcPending
						if v76 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v80 = m.ExcPending
							if v80 != 0 {
								return
							} else {
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_ExecGrant_Language_check(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
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
	var v19 int32
	_ = v19
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+22)))
	v9 = v7 + v8
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+73)))
	if v10 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			F_errcode(m, int32(151027844))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v5))) = v9 + int32(4)
				F_errmsg(m, int32(_a_F_ExecGrant_Language_check_0), v5)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v28 = F_errdetail(m, int32(_a_F_ExecGrant_Language_check_1), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return
					} else {
						F_errfinish(m, int32(_a_F_ExecGrant_Language_check_2), int32(2252), int32(_a_F_ExecGrant_Language_check_3))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
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
		m.G0 = v5 + int32(16)
		return
	}
}
func F_ExecGrant_Type_check(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+22)))
	v5 = v3 + v4
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)+92))
	if v6 != 0 {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)+88))
		if v7 == int32(_a_F_ExecGrant_Type_check_0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errcode(m, int32(16910080))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_ExecGrant_Type_check_1), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_ExecGrant_Type_check_2), int32(0))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecGrant_Type_check_3), int32(2403), int32(_a_F_ExecGrant_Type_check_4))
							mBase = m.M
							v32 = m.ExcPending
							if v32 != 0 {
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
			v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+79)))
			if v10 == int32(109) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					F_errcode(m, int32(16910080))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						F_errmsg(m, int32(_a_F_ExecGrant_Type_check_5), int32(0))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							F_errhint(m, int32(_a_F_ExecGrant_Type_check_6), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_ExecGrant_Type_check_3), int32(2408), int32(_a_F_ExecGrant_Type_check_4))
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
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
				return
			}
		}
	} else {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+79)))
		if v10 == int32(109) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return
			} else {
				F_errcode(m, int32(16910080))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_ExecGrant_Type_check_5), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return
					} else {
						F_errhint(m, int32(_a_F_ExecGrant_Type_check_6), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return
						} else {
							F_errfinish(m, int32(_a_F_ExecGrant_Type_check_3), int32(2408), int32(_a_F_ExecGrant_Type_check_4))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
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
			return
		}
	}
}
func F_ExecGrant_common(m *base.Module, l0 int32, l1 int32, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
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
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int64
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v116 int32
	_ = v116
	var v117 int64
	_ = v117
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int64
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int64
	_ = v129
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v246 int32
	_ = v246
	v19 = m.G0
	v21 = v19 - int32(48)
	m.G0 = v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	if v23 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v30 = F_get_object_catcache_oid(m, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	if v26 != int64(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = l2
	goto L1
L4:
	;
	return
L5:
	;
	v33 = F_table_open(m, l1, int32(3))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L4
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v35 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L4
	} else {
		goto L58
	}
L8:
	;
	F_relation_close(m, v33, int32(3))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L4
	} else {
		goto L57
	}
L9:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v38 <= int32(0) {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v54 = int32(0)
	goto L11
L11:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v35)+12))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v59+v54<<(uint(int32(2))%32))))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = F_palloc0_mul(m, int32(8), v66)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	goto L8
L13:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = F_palloc0_mul(m, int32(1), v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v75)))
	v77 = F_palloc0_mul(m, int32(1), v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v80 = F_SearchSysCacheLocked1(m, v30, base.I64_extend_i32_u(v63))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v80 == int32(0) {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	if l3 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	m.T0[l3].(func(*base.Module, int32, int32))(m, l0, v80)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v86 = F_get_object_attnum_owner(m, l1)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	v88 = F_SysCacheGetAttrNotNull(m, v30, v80, v86)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	v90 = base.I32_wrap_i64(v88)
	v91 = F_get_object_attnum_acl(m, l1)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v95 = F_SysCacheGetAttr(m, v30, v80, v91, v21+int32(47))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+47)))
	if v97 == int32(1) {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v117 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	F_select_best_grantor(m, v116, v117, v114, v90, v21+int32(28), v21+int32(32))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L34
	}
L27:
	;
	v101 = F_get_object_type(m, l1, v63)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v108 = F_pg_detoast_datum_copy(m, base.I32_wrap_i64(v95))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L4
	} else {
		goto L32
	}
L30:
	;
	v103 = F_acldefault(m, v101, v90)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L4
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = int32(0)
	v114 = v103
	v115 = int32(0)
	goto L26
L32:
	;
	v112 = F_aclmembers(m, v108, v21+int32(24))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L33
	}
L33:
	;
	v114 = v108
	v115 = v112
	goto L26
L34:
	;
	v124 = F_get_object_attnum_name(m, l1)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L4
	} else {
		goto L35
	}
L35:
	;
	v126 = F_SysCacheGetAttrNotNull(m, v30, v80, v124)
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L36
	}
L36:
	;
	v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v129 = *(*int64)(unsafe.Add(mBase, uint32(v21)+32))
	v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+12)))
	v131 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v133 = F_get_object_type(m, l1, v63)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L4
	} else {
		goto L37
	}
L37:
	;
	v136 = int32(0)
	v138 = F_restrict_and_check_grant(m, v128, v129, v130, v131, v63, v132, v133, base.I32_wrap_i64(v126), v136, v136)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L4
	} else {
		goto L38
	}
L38:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v142 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
	v145 = F_merge_acl_with_grant(m, v114, v140, v141, v142, v143, v138, v144, v90)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L4
	} else {
		goto L39
	}
L39:
	;
	v149 = F_aclmembers(m, v145, v21+int32(20))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	v151 = F_get_object_attnum_acl(m, l1)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L4
	} else {
		goto L41
	}
L41:
	;
	v154 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v151+v77-v154))) = uint8(v154)
	v158 = F_get_object_attnum_acl(m, l1)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67+v158<<(uint(int32(3))%32)-int32(8)))) = base.I64_extend_i32_u(v145)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v33)+52))
	v168 = F_heap_modify_tuple(m, v80, v167, v67, v72, v77)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	F_CatalogTupleUpdate(m, v33, v168+int32(4), v168)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L4
	} else {
		goto L44
	}
L44:
	;
	F_UnlockTuple(m, v33, v80+int32(4), int32(7))
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L4
	} else {
		goto L45
	}
L45:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrant_common[0])))
	if v180 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(v21)+24))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	F_updateAclDependencies(m, l1, v63, int32(0), v90, v115, v193, v149, v194)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L4
	} else {
		goto L52
	}
L47:
	;
	v184 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecGrant_common[1])))
	if v184&int32(1) == int32(0) {
		goto L46
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	F_recordExtensionInitPrivWorker(m, v63, l1, int32(0), v145)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L4
	} else {
		goto L51
	}
L50:
	;
	goto L49
L51:
	;
	goto L46
L52:
	;
	F_ReleaseCatCache(m, v80)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L4
	} else {
		goto L53
	}
L53:
	;
	F_pfree(m, v145)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	F_CommandCounterIncrement(m)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L55
	}
L55:
	;
	v204 = v54 + int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v35)+4))
	if v204 < v205 {
		v54 = v204
		goto L11
	} else {
		goto L56
	}
L56:
	;
	goto L12
L57:
	;
	m.G0 = v21 + int32(48)
	return
L58:
	;
	v235 = F_get_object_class_descr(m, l1)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v63
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v235
	F_errmsg_internal(m, int32(_a_F_ExecGrant_common_0), v21)
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	F_errfinish(m, int32(_a_F_ExecGrant_common_1), int32(2142), int32(_a_F_ExecGrant_common_2))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L4
	} else {
		goto L61
	}
L61:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_ExecMergeAppend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int64
	_ = v217
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v234 int32
	_ = v234
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_ExecMergeAppend[0]))
	if v7 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)))
	if v12 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L4:
	;
	return int32(0)
L5:
	;
	goto L3
L6:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v241)+8))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
	m.T0[v243].(func(*base.Module, int32))(m, v241)
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L4
	} else {
		goto L69
	}
L7:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)))
	if v225 == int32(0) {
		goto L66
	} else {
		goto L67
	}
L8:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+108))
	if v15 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v188 = *(*int64)(unsafe.Add(mBase, uint32(v187)+24))
	v191 = base.I32_wrap_i64(v188) << (uint(int32(2)) % 32)
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v191+v192)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+52))
	if v195 != 0 {
		goto L56
	} else {
		goto L57
	}
L11:
	;
	goto L6
L12:
	;
	goto L13
L13:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v18 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+132))
	v22 = int32(0)
	v24 = F_ExecFindMatchingSubPlans(m, v21, v22, v22)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L17
	}
L15:
	;
	v27 = v18
	goto L16
L16:
	;
	if v27 == int32(0) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+136)) = v24
	v27 = v24
	goto L16
L18:
	;
	if int32(0) <= v84 {
		goto L29
	} else {
		goto L30
	}
L19:
	;
	v84 = base.I32_ctz(v70) | v71<<(uint(int32(5))%32)
	goto L18
L20:
	;
	v84 = int32(-2)
	goto L18
L21:
	;
	v35 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	if v38 <= v35 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v41 = v27 + int32(8)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)))
	v48 = v45 & int32(-1)
	if v48 != 0 {
		v70 = v48
		v71 = v35
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v49 = int32(1)
	if v49 == v38 {
		goto L20
	} else {
		goto L24
	}
L24:
	;
	v53 = v49
	goto L25
L25:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v41+v53<<(uint(int32(2))%32))))
	if v60 != 0 {
		v70 = v60
		v71 = v53
		goto L19
	} else {
		goto L27
	}
L26:
	;
	goto L20
L27:
	;
	v62 = v53 + int32(1)
	if v62 != v38 {
		v53 = v62
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v88 = v84
	goto L32
L30:
	;
	goto L31
L31:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_binaryheap_build(m, v182)
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L4
	} else {
		goto L55
	}
L32:
	;
	v93 = v88 << (uint(int32(2)) % 32)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+104))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93+v94)))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+52))
	if v97 != 0 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	goto L31
L34:
	;
	F_ExecReScan(m, v96)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	goto L36
L36:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v101 = m.T0[v100].(func(*base.Module, int32) int32)(m, v96)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L38
	}
L37:
	;
	goto L36
L38:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v103+v93))) = v101
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v106+v93)))
	if v108 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+136))
	if v118 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L40:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108)+4)))
	if v111&int32(2) != 0 {
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_binaryheap_add_unordered(m, v114, base.I64_extend_i32_u(v88))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L4
	} else {
		goto L42
	}
L42:
	;
	goto L39
L43:
	;
	if int32(0) <= v174 {
		v88 = v174
		goto L32
	} else {
		goto L54
	}
L44:
	;
	v174 = base.I32_ctz(v160) | v161<<(uint(int32(5))%32)
	goto L43
L45:
	;
	v174 = int32(-2)
	goto L43
L46:
	;
	v125 = v88 + int32(1)
	v127 = int32(base.Ui32(v125) >> (uint(int32(5)) % 32))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v118)+4))
	if v128 <= v127 {
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v131 = v118 + int32(8)
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v131+v127<<(uint(int32(2))%32))))
	v138 = v135 & (int32(-1) << (uint(v125) % 32))
	if v138 != 0 {
		v160 = v138
		v161 = v127
		goto L44
	} else {
		goto L48
	}
L48:
	;
	v140 = v127 + int32(1)
	if v140 == v128 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v143 = v140
	goto L50
L50:
	;
	v150 = *(*int32)(unsafe.Add(mBase, uint32(v131+v143<<(uint(int32(2))%32))))
	if v150 != 0 {
		v160 = v150
		v161 = v143
		goto L44
	} else {
		goto L52
	}
L51:
	;
	goto L45
L52:
	;
	v152 = v143 + int32(1)
	if v152 != v128 {
		v143 = v152
		goto L50
	} else {
		goto L53
	}
L53:
	;
	goto L51
L54:
	;
	goto L33
L55:
	;
	v185 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+128)) = uint8(v185)
	goto L7
L56:
	;
	F_ExecReScan(m, v194)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	v199 = m.T0[v198].(func(*base.Module, int32) int32)(m, v194)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L4
	} else {
		goto L60
	}
L59:
	;
	goto L58
L60:
	;
	v201 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	*(*int32)(unsafe.Add(mBase, uint32(v201+v191))) = v199
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v204+v191)))
	if v206 == int32(0) {
		goto L61
	} else {
		goto L62
	}
L61:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	v217 = F_binaryheap_remove_first(m, v216)
	mBase = m.M
	v218 = m.ExcPending
	if v218 != 0 {
		goto L4
	} else {
		goto L65
	}
L62:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
	if v209&int32(2) != 0 {
		goto L61
	} else {
		goto L63
	}
L63:
	;
	v212 = *(*int32)(unsafe.Add(mBase, uint32(l0)+124))
	F_binaryheap_replace_first(m, v212, base.I64_extend32_s(v188))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	goto L7
L65:
	;
	goto L7
L66:
	;
	goto L6
L67:
	;
	goto L68
L68:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
	v229 = *(*int64)(unsafe.Add(mBase, uint32(v224)+24))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v228+base.I32_wrap_i64(v229)<<(uint(int32(2))%32))))
	return v234
L69:
	;
	return v241
}
func F_ExecReadyInterpretedExpr(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v262 int32
	_ = v262
	var v267 int32
	_ = v267
	v2 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
	if v8 == v2 {
		v12 = int32(0)
		v15 = F_ExecInterpExpr(m, v12, v12, v12)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = base.I32_wrap_i64(v15)
			*(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0])) = v17
			v20 = v2
			for {
				v25 = int32(2)
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v17+v20<<(uint(v25)%32))))
				v29 = int32(3)
				v30 = v20 << (uint(v29) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_ExecReadyInterpretedExpr[1]))) = v20
				*(*int32)(unsafe.Add(mBase, uint32(v30)+uint32(_c_F_ExecReadyInterpretedExpr[2]))) = v28
				v34 = v20 | int32(1)
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v17+v34<<(uint(v25)%32))))
				v40 = v34 << (uint(v29) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_ExecReadyInterpretedExpr[1]))) = v34
				*(*int32)(unsafe.Add(mBase, uint32(v40)+uint32(_c_F_ExecReadyInterpretedExpr[2]))) = v38
				v44 = v20 | v25
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v17+v44<<(uint(v25)%32))))
				v50 = v44 << (uint(v29) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_ExecReadyInterpretedExpr[1]))) = v44
				*(*int32)(unsafe.Add(mBase, uint32(v50)+uint32(_c_F_ExecReadyInterpretedExpr[2]))) = v48
				v54 = v20 | v29
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v17+v54<<(uint(v25)%32))))
				v60 = v54 << (uint(v29) % 32)
				*(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_ExecReadyInterpretedExpr[1]))) = v54
				*(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_ExecReadyInterpretedExpr[2]))) = v58
				v64 = v20 + int32(4)
				if v64 != int32(120) {
					v20 = v64
					continue
				} else {
					break
				}
				break
			}
			F_pg_qsort(m, int32(_a_F_ExecReadyInterpretedExpr_0), int32(120), int32(8), int32(634))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return
			} else {
				v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
				if v79&int32(32) == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(635)
					v87 = v79 | int32(32)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v87)
					v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					switch v89 - int32(2) {
					case 0:
						v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
						switch v217 - int32(7) {
						case 0:
							v267 = int32(656)
						case 1:
							v267 = int32(650)
						case 2:
							v267 = int32(651)
						default:
							v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
							v235 = int32(0)
							for {
								v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v243 = v240 + v235*int32(40)
								v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
								v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
								v251 = v235 + int32(1)
								v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v251 < v252 {
									v235 = v251
									continue
								} else {
									break
								}
								break
							}
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
							v256 = v254
							v262 = v256 | int32(64)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
							v267 = int32(655)
						case 11:
							v267 = int32(652)
						case 12:
							v267 = int32(653)
						case 13:
							v267 = int32(654)
						case 18:
							v267 = int32(649)
						}
					case 1:
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+40))
						v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
						if base.B2i32(v142 != int32(2))|base.B2i32(v141 != int32(7)) == int32(0) {
							v267 = int32(640)
						} else {
							if base.B2i32(v142 != int32(3))|base.B2i32(v141 != int32(8)) == int32(0) {
								v267 = int32(641)
							} else {
								if base.B2i32(v142 != int32(4))|base.B2i32(v141 != int32(9)) == int32(0) {
									v267 = int32(642)
								} else {
									if base.B2i32(v142 != int32(2))|base.B2i32(v141 != int32(18)) == int32(0) {
										v267 = int32(643)
									} else {
										if base.B2i32(v142 != int32(3))|base.B2i32(v141 != int32(19)) == int32(0) {
											v267 = int32(644)
										} else {
											if base.B2i32(v142 != int32(4))|base.B2i32(v141 != int32(20)) == int32(0) {
												v267 = int32(645)
											} else {
												if base.B2i32(v142 != int32(56))|base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v141-int32(27))) == int32(0) {
													v267 = int32(646)
												} else {
													if base.B2i32(v142 != int32(7))|base.B2i32(v141 != int32(86)) == int32(0) {
														v267 = int32(647)
													} else {
														if base.B2i32(v142 != int32(8))|base.B2i32(v141 != int32(86)) != 0 {
															v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
															v235 = int32(0)
															for {
																v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																v243 = v240 + v235*int32(40)
																v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
																v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
																*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
																v251 = v235 + int32(1)
																v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																if v251 < v252 {
																	v235 = v251
																	continue
																} else {
																	break
																}
																break
															}
															v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
															v256 = v254
															v262 = v256 | int32(64)
															*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
															v267 = int32(655)
														} else {
															v267 = int32(648)
														}
													}
												}
											}
										}
									}
								}
							}
						}
					case 2:
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+80))
						v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
						v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)+40))
						v114 = base.B2i32(v108 == int32(3)) & base.B2i32(v111 == int32(8))
						v115 = int32(0)
						if base.B2i32(v114 == v115)|base.B2i32(v107 != int32(86)) == v115 {
							v267 = int32(637)
						} else {
							if base.B2i32(v108 != int32(2))|base.B2i32(v111 != int32(7))|base.B2i32(v107 != int32(86)) == int32(0) {
								v267 = int32(638)
							} else {
								if v114&base.B2i32(v107 == int32(87)) == int32(0) {
									v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
									v235 = int32(0)
									for {
										v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v243 = v240 + v235*int32(40)
										v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
										v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
										v251 = v235 + int32(1)
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v251 < v252 {
											v235 = v251
											continue
										} else {
											break
										}
										break
									}
									v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
									v256 = v254
									v262 = v256 | int32(64)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
									v267 = int32(655)
								} else {
									v267 = int32(639)
								}
							}
						}
					case 3:
						v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
						if v93 != int32(2) {
							v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
							v235 = int32(0)
							for {
								v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v243 = v240 + v235*int32(40)
								v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
								v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
								v251 = v235 + int32(1)
								v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v251 < v252 {
									v235 = v251
									continue
								} else {
									break
								}
								break
							}
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
							v256 = v254
							v262 = v256 | int32(64)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
							v267 = int32(655)
						} else {
							v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+40))
							if v96 != int32(85) {
								v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
								v235 = int32(0)
								for {
									v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v243 = v240 + v235*int32(40)
									v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
									v251 = v235 + int32(1)
									v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v251 < v252 {
										v235 = v251
										continue
									} else {
										break
									}
									break
								}
								v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
								v256 = v254
								v262 = v256 | int32(64)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
								v267 = int32(655)
							} else {
								v99 = *(*int32)(unsafe.Add(mBase, uint32(v92)+80))
								if v99 != int32(7) {
									v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
									v235 = int32(0)
									for {
										v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										v243 = v240 + v235*int32(40)
										v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
										v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
										*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
										v251 = v235 + int32(1)
										v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v251 < v252 {
											v235 = v251
											continue
										} else {
											break
										}
										break
									}
									v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
									v256 = v254
									v262 = v256 | int32(64)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
									v267 = int32(655)
								} else {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v92)+120))
									if v102 != int32(88) {
										v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
										v235 = int32(0)
										for {
											v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											v243 = v240 + v235*int32(40)
											v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
											v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
											*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
											v251 = v235 + int32(1)
											v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v251 < v252 {
												v235 = v251
												continue
											} else {
												break
											}
											break
										}
										v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
										v256 = v254
										v262 = v256 | int32(64)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
										v267 = int32(655)
									} else {
										v267 = int32(636)
									}
								}
							}
						}
					default:
						if v89 <= int32(0) {
							v256 = v87
						} else {
							v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
							v235 = int32(0)
							for {
								v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v243 = v240 + v235*int32(40)
								v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
								v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
								v251 = v235 + int32(1)
								v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v251 < v252 {
									v235 = v251
									continue
								} else {
									break
								}
								break
							}
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
							v256 = v254
						}
						v262 = v256 | int32(64)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
						v267 = int32(655)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v267
				} else {
				}
				return
			}
		}
	} else {
		v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
		if v79&int32(32) == int32(0) {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(635)
			v87 = v79 | int32(32)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v87)
			v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			switch v89 - int32(2) {
			case 0:
				v216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
				switch v217 - int32(7) {
				case 0:
					v267 = int32(656)
				case 1:
					v267 = int32(650)
				case 2:
					v267 = int32(651)
				default:
					v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
					v235 = int32(0)
					for {
						v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v243 = v240 + v235*int32(40)
						v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
						v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
						v251 = v235 + int32(1)
						v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v251 < v252 {
							v235 = v251
							continue
						} else {
							break
						}
						break
					}
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
					v256 = v254
					v262 = v256 | int32(64)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
					v267 = int32(655)
				case 11:
					v267 = int32(652)
				case 12:
					v267 = int32(653)
				case 13:
					v267 = int32(654)
				case 18:
					v267 = int32(649)
				}
			case 1:
				v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+40))
				v142 = *(*int32)(unsafe.Add(mBase, uint32(v140)))
				if base.B2i32(v142 != int32(2))|base.B2i32(v141 != int32(7)) == int32(0) {
					v267 = int32(640)
				} else {
					if base.B2i32(v142 != int32(3))|base.B2i32(v141 != int32(8)) == int32(0) {
						v267 = int32(641)
					} else {
						if base.B2i32(v142 != int32(4))|base.B2i32(v141 != int32(9)) == int32(0) {
							v267 = int32(642)
						} else {
							if base.B2i32(v142 != int32(2))|base.B2i32(v141 != int32(18)) == int32(0) {
								v267 = int32(643)
							} else {
								if base.B2i32(v142 != int32(3))|base.B2i32(v141 != int32(19)) == int32(0) {
									v267 = int32(644)
								} else {
									if base.B2i32(v142 != int32(4))|base.B2i32(v141 != int32(20)) == int32(0) {
										v267 = int32(645)
									} else {
										if base.B2i32(v142 != int32(56))|base.B2i32(base.Ui32(int32(3)) <= base.Ui32(v141-int32(27))) == int32(0) {
											v267 = int32(646)
										} else {
											if base.B2i32(v142 != int32(7))|base.B2i32(v141 != int32(86)) == int32(0) {
												v267 = int32(647)
											} else {
												if base.B2i32(v142 != int32(8))|base.B2i32(v141 != int32(86)) != 0 {
													v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
													v235 = int32(0)
													for {
														v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
														v243 = v240 + v235*int32(40)
														v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
														v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
														*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
														v251 = v235 + int32(1)
														v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v251 < v252 {
															v235 = v251
															continue
														} else {
															break
														}
														break
													}
													v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
													v256 = v254
													v262 = v256 | int32(64)
													*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
													v267 = int32(655)
												} else {
													v267 = int32(648)
												}
											}
										}
									}
								}
							}
						}
					}
				}
			case 2:
				v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v107 = *(*int32)(unsafe.Add(mBase, uint32(v106)+80))
				v108 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v111 = *(*int32)(unsafe.Add(mBase, uint32(v106)+40))
				v114 = base.B2i32(v108 == int32(3)) & base.B2i32(v111 == int32(8))
				v115 = int32(0)
				if base.B2i32(v114 == v115)|base.B2i32(v107 != int32(86)) == v115 {
					v267 = int32(637)
				} else {
					if base.B2i32(v108 != int32(2))|base.B2i32(v111 != int32(7))|base.B2i32(v107 != int32(86)) == int32(0) {
						v267 = int32(638)
					} else {
						if v114&base.B2i32(v107 == int32(87)) == int32(0) {
							v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
							v235 = int32(0)
							for {
								v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v243 = v240 + v235*int32(40)
								v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
								v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
								v251 = v235 + int32(1)
								v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v251 < v252 {
									v235 = v251
									continue
								} else {
									break
								}
								break
							}
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
							v256 = v254
							v262 = v256 | int32(64)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
							v267 = int32(655)
						} else {
							v267 = int32(639)
						}
					}
				}
			case 3:
				v92 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
				if v93 != int32(2) {
					v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
					v235 = int32(0)
					for {
						v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v243 = v240 + v235*int32(40)
						v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
						v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
						v251 = v235 + int32(1)
						v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v251 < v252 {
							v235 = v251
							continue
						} else {
							break
						}
						break
					}
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
					v256 = v254
					v262 = v256 | int32(64)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
					v267 = int32(655)
				} else {
					v96 = *(*int32)(unsafe.Add(mBase, uint32(v92)+40))
					if v96 != int32(85) {
						v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
						v235 = int32(0)
						for {
							v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							v243 = v240 + v235*int32(40)
							v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
							v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
							*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
							v251 = v235 + int32(1)
							v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
							if v251 < v252 {
								v235 = v251
								continue
							} else {
								break
							}
							break
						}
						v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
						v256 = v254
						v262 = v256 | int32(64)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
						v267 = int32(655)
					} else {
						v99 = *(*int32)(unsafe.Add(mBase, uint32(v92)+80))
						if v99 != int32(7) {
							v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
							v235 = int32(0)
							for {
								v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								v243 = v240 + v235*int32(40)
								v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
								v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
								v251 = v235 + int32(1)
								v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
								if v251 < v252 {
									v235 = v251
									continue
								} else {
									break
								}
								break
							}
							v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
							v256 = v254
							v262 = v256 | int32(64)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
							v267 = int32(655)
						} else {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v92)+120))
							if v102 != int32(88) {
								v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
								v235 = int32(0)
								for {
									v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									v243 = v240 + v235*int32(40)
									v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
									v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
									*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
									v251 = v235 + int32(1)
									v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
									if v251 < v252 {
										v235 = v251
										continue
									} else {
										break
									}
									break
								}
								v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
								v256 = v254
								v262 = v256 | int32(64)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
								v267 = int32(655)
							} else {
								v267 = int32(636)
							}
						}
					}
				}
			default:
				if v89 <= int32(0) {
					v256 = v87
				} else {
					v233 = *(*int32)(unsafe.Add(mBase, _c_F_ExecReadyInterpretedExpr[0]))
					v235 = int32(0)
					for {
						v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						v243 = v240 + v235*int32(40)
						v244 = *(*int32)(unsafe.Add(mBase, uint32(v243)))
						v248 = *(*int32)(unsafe.Add(mBase, uint32(v233+v244<<(uint(int32(2))%32))))
						*(*int32)(unsafe.Add(mBase, uint32(v243))) = v248
						v251 = v235 + int32(1)
						v252 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						if v251 < v252 {
							v235 = v251
							continue
						} else {
							break
						}
						break
					}
					v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)))
					v256 = v254
				}
				v262 = v256 | int32(64)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v262)
				v267 = int32(655)
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v267
		} else {
		}
		return
	}
}
func F_ExecRenameStmt(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int64
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
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
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v333 int64
	_ = v333
	var v334 int64
	_ = v334
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v389 int64
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v421 int64
	_ = v421
	var v422 int64
	_ = v422
	var v423 int64
	_ = v423
	var v424 int64
	_ = v424
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v442 int64
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v452 int64
	_ = v452
	var v453 int32
	_ = v453
	var v457 int64
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int64
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v528 int32
	_ = v528
	var v531 int32
	_ = v531
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v603 int32
	_ = v603
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v625 int32
	_ = v625
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v634 int32
	_ = v634
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v711 int32
	_ = v711
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v776 int32
	_ = v776
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v801 int32
	_ = v801
	var v806 int32
	_ = v806
	var v810 int32
	_ = v810
	var v813 int32
	_ = v813
	var v819 int32
	_ = v819
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v840 int32
	_ = v840
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v885 int32
	_ = v885
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v925 int32
	_ = v925
	var v927 int32
	_ = v927
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v941 int32
	_ = v941
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v973 int32
	_ = v973
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v995 int32
	_ = v995
	var v1000 int32
	_ = v1000
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1021 int32
	_ = v1021
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
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1067 int32
	_ = v1067
	var v1069 int32
	_ = v1069
	var v1072 int64
	_ = v1072
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1082 int32
	_ = v1082
	var v1093 int32
	_ = v1093
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1125 int64
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1135 int32
	_ = v1135
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1148 int32
	_ = v1148
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1166 int32
	_ = v1166
	var v1167 int32
	_ = v1167
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1175 int64
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1185 int32
	_ = v1185
	var v1188 int32
	_ = v1188
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1199 int32
	_ = v1199
	var v1206 int32
	_ = v1206
	var v1207 int32
	_ = v1207
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1233 int32
	_ = v1233
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1243 int32
	_ = v1243
	var v1250 int32
	_ = v1250
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1261 int32
	_ = v1261
	var v1266 int32
	_ = v1266
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1274 int32
	_ = v1274
	var v1283 int32
	_ = v1283
	var v1288 int32
	_ = v1288
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1299 int32
	_ = v1299
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1322 int32
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1338 int32
	_ = v1338
	var v1340 int32
	_ = v1340
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1348 int32
	_ = v1348
	var v1349 int32
	_ = v1349
	var v1352 int32
	_ = v1352
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1361 int32
	_ = v1361
	var v1367 int64
	_ = v1367
	var v1369 int32
	_ = v1369
	var v1374 int32
	_ = v1374
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1380 int32
	_ = v1380
	var v1381 int32
	_ = v1381
	var v1382 int32
	_ = v1382
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1392 int32
	_ = v1392
	var v1393 int32
	_ = v1393
	var v1394 int32
	_ = v1394
	var v1400 int32
	_ = v1400
	var v1412 int32
	_ = v1412
	var v1416 int32
	_ = v1416
	var v1417 int32
	_ = v1417
	var v1418 int32
	_ = v1418
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1423 int32
	_ = v1423
	var v1424 int32
	_ = v1424
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1494 int32
	_ = v1494
	var v1496 int32
	_ = v1496
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1505 int32
	_ = v1505
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1515 int32
	_ = v1515
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int64
	_ = v1533
	var v1535 int32
	_ = v1535
	var v1537 int32
	_ = v1537
	var v1541 int64
	_ = v1541
	var v1543 int32
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1551 int32
	_ = v1551
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1564 int64
	_ = v1564
	var v1566 int32
	_ = v1566
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
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1600 int32
	_ = v1600
	var v1604 int32
	_ = v1604
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1616 int32
	_ = v1616
	var v1619 int32
	_ = v1619
	var v1626 int32
	_ = v1626
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1631 int32
	_ = v1631
	var v1640 int32
	_ = v1640
	var v1645 int32
	_ = v1645
	var v1649 int32
	_ = v1649
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1654 int32
	_ = v1654
	var v1661 int32
	_ = v1661
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1671 int32
	_ = v1671
	var v1673 int32
	_ = v1673
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1676 int32
	_ = v1676
	var v1677 int32
	_ = v1677
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1692 int32
	_ = v1692
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1709 int32
	_ = v1709
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1714 int32
	_ = v1714
	var v1720 int32
	_ = v1720
	var v1725 int32
	_ = v1725
	var v1728 int32
	_ = v1728
	var v1729 int32
	_ = v1729
	var v1730 int32
	_ = v1730
	var v1733 int32
	_ = v1733
	var v1734 int32
	_ = v1734
	var v1737 int32
	_ = v1737
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1747 int32
	_ = v1747
	var v1755 int32
	_ = v1755
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1771 int32
	_ = v1771
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1786 int32
	_ = v1786
	var v1793 int32
	_ = v1793
	var v1798 int32
	_ = v1798
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1813 int32
	_ = v1813
	var v1814 int32
	_ = v1814
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1822 int32
	_ = v1822
	var v1827 int32
	_ = v1827
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1837 int32
	_ = v1837
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1846 int32
	_ = v1846
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1856 int32
	_ = v1856
	var v1858 int32
	_ = v1858
	var v1859 int64
	_ = v1859
	var v1860 int32
	_ = v1860
	var v1863 int32
	_ = v1863
	var v1864 int64
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1868 int32
	_ = v1868
	var v1869 int32
	_ = v1869
	var v1874 int32
	_ = v1874
	var v1877 int64
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1887 int32
	_ = v1887
	var v1888 int32
	_ = v1888
	var v1891 int32
	_ = v1891
	var v1896 int32
	_ = v1896
	var v1898 int32
	_ = v1898
	var v1899 int32
	_ = v1899
	var v1903 int32
	_ = v1903
	var v1904 int32
	_ = v1904
	var v1906 int32
	_ = v1906
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1926 int32
	_ = v1926
	var v1927 int32
	_ = v1927
	var v1930 int32
	_ = v1930
	var v1931 int32
	_ = v1931
	var v1932 int32
	_ = v1932
	var v1934 int32
	_ = v1934
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1938 int32
	_ = v1938
	var v1939 int32
	_ = v1939
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1947 int32
	_ = v1947
	var v1951 int32
	_ = v1951
	var v1955 int32
	_ = v1955
	var v1960 int32
	_ = v1960
	var v1977 int64
	_ = v1977
	var v1981 int64
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1994 int32
	_ = v1994
	var v1996 int32
	_ = v1996
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2002 int32
	_ = v2002
	var v2003 int32
	_ = v2003
	var v2005 int32
	_ = v2005
	var v2011 int64
	_ = v2011
	var v2013 int64
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2016 int32
	_ = v2016
	var v2018 int32
	_ = v2018
	var v2019 int64
	_ = v2019
	var v2022 int32
	_ = v2022
	var v2023 int32
	_ = v2023
	var v2044 int32
	_ = v2044
	var v2048 int32
	_ = v2048
	var v2051 int32
	_ = v2051
	var v2056 int32
	_ = v2056
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2081 int32
	_ = v2081
	var v2086 int32
	_ = v2086
	var v2087 int32
	_ = v2087
	var v2089 int32
	_ = v2089
	var v2091 int32
	_ = v2091
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2111 int32
	_ = v2111
	var v2113 int32
	_ = v2113
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
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2129 int32
	_ = v2129
	var v2135 int32
	_ = v2135
	var v2140 int32
	_ = v2140
	var v2142 int32
	_ = v2142
	var v2145 int64
	_ = v2145
	var v2150 int32
	_ = v2150
	var v2154 int32
	_ = v2154
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2173 int32
	_ = v2173
	var v2184 int32
	_ = v2184
	var v2185 int32
	_ = v2185
	var v2194 int32
	_ = v2194
	var v2199 int32
	_ = v2199
	var v2203 int32
	_ = v2203
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2214 int32
	_ = v2214
	var v2219 int32
	_ = v2219
	var v2223 int32
	_ = v2223
	var v2226 int32
	_ = v2226
	var v2232 int32
	_ = v2232
	var v2237 int32
	_ = v2237
	var v2242 int32
	_ = v2242
	var v2248 int32
	_ = v2248
	var v2253 int32
	_ = v2253
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2262 int32
	_ = v2262
	var v2267 int32
	_ = v2267
	var v2268 int32
	_ = v2268
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2274 int32
	_ = v2274
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2277 int32
	_ = v2277
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2280 int32
	_ = v2280
	var v2282 int32
	_ = v2282
	var v2284 int32
	_ = v2284
	var v2285 int32
	_ = v2285
	var v2287 int32
	_ = v2287
	var v2288 int32
	_ = v2288
	var v2299 int32
	_ = v2299
	var v2300 int32
	_ = v2300
	var v2301 int32
	_ = v2301
	var v2303 int32
	_ = v2303
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2310 int32
	_ = v2310
	var v2313 int32
	_ = v2313
	var v2314 int32
	_ = v2314
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2320 int32
	_ = v2320
	var v2322 int32
	_ = v2322
	var v2323 int32
	_ = v2323
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2330 int32
	_ = v2330
	var v2337 int32
	_ = v2337
	var v2348 int32
	_ = v2348
	var v2352 int32
	_ = v2352
	var v2354 int32
	_ = v2354
	var v2356 int32
	_ = v2356
	var v2357 int32
	_ = v2357
	var v2402 int32
	_ = v2402
	var v2404 int32
	_ = v2404
	var v2406 int32
	_ = v2406
	var v2408 int32
	_ = v2408
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	v3 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(160)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	switch v19 - int32(1) {
	case 0, 6, 7, 13, 15, 16, 18, 20, 23, 25, 28, 29, 34, 38, 39, 45, 46, 47, 48:
		goto L10
	default:
		goto L9
	case 3, 5:
		goto L15
	case 8:
		goto L20
	case 11, 49:
		goto L11
	case 12, 40:
		goto L8
	case 17, 19, 22, 37, 41, 51:
		goto L16
	case 27:
		goto L12
	case 33:
		goto L19
	case 35:
		goto L14
	case 36:
		goto L18
	case 42:
		goto L17
	case 44:
		goto L13
	}
L1:
	;
	m.G0 = v17 + int32(160)
	return
L2:
	;
	v2267 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+48))
	v2268 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2267)+120)))
	v2271 = F_palloc0(m, v2268<<(uint(int32(3))%32))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L35
	} else {
		goto L695
	}
L3:
	;
	v2256 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+16))
	v2257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2256)+22)))
	v2258 = v2256 + v2257
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v2258)+4))
	v2260 = *(*int32)(unsafe.Add(mBase, uint32(v2258)+72))
	F_IsThereOpClassInNamespace(m, v1838, v2259, v2260)
	mBase = m.M
	v2262 = m.ExcPending
	if v2262 != 0 {
		goto L35
	} else {
		goto L694
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L35
	} else {
		goto L691
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2223 = m.ExcPending
	if v2223 != 0 {
		goto L35
	} else {
		goto L687
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2203 = m.ExcPending
	if v2203 != 0 {
		goto L35
	} else {
		goto L682
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2184 = m.ExcPending
	if v2184 != 0 {
		goto L35
	} else {
		goto L679
	}
L8:
	;
	v2087 = m.G0
	v2089 = v2087 - int32(32)
	m.G0 = v2089
	v2091 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v2091 == int32(13) {
		goto L653
	} else {
		goto L654
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2076 = m.ExcPending
	if v2076 != 0 {
		goto L35
	} else {
		goto L647
	}
L10:
	;
	v1828 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1829 = int32(0)
	F_get_object_address(m, l0, v19, v1828, v1829, int32(8), v1829)
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L35
	} else {
		goto L549
	}
L11:
	;
	v1667 = m.G0
	v1669 = v1667 - int32(96)
	m.G0 = v1669
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v1674 = F_makeTypeNameFromNameList(m, v1673)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L35
	} else {
		goto L494
	}
L12:
	;
	v1511 = m.G0
	v1513 = v1511 - int32(144)
	m.G0 = v1513
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1519 = F_RangeVarGetRelidExtended(m, v1515, int32(8), int32(0), int32(608), l1)
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L35
	} else {
		goto L457
	}
L13:
	;
	v1328 = m.G0
	v1330 = v1328 - int32(160)
	m.G0 = v1330
	v1332 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1334 = int32(0)
	v1337 = F_RangeVarGetRelidExtended(m, v1332, int32(8), v1334, int32(625), v1334)
	mBase = m.M
	v1338 = m.ExcPending
	if v1338 != 0 {
		goto L35
	} else {
		goto L417
	}
L14:
	;
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1146 = m.G0
	v1148 = v1146 - int32(48)
	m.G0 = v1148
	v1151 = int32(0)
	v1154 = F_RangeVarGetRelidExtended(m, v1143, int32(8), v1151, int32(1118), v1151)
	mBase = m.M
	v1155 = m.ExcPending
	if v1155 != 0 {
		goto L35
	} else {
		goto L369
	}
L15:
	;
	v1093 = m.G0
	v1095 = v1093 - int32(16)
	m.G0 = v1095
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v1102 = F_RangeVarGetRelidExtended(m, v1097, int32(8), v1099, int32(620), int32(0))
	mBase = m.M
	v1103 = m.ExcPending
	if v1103 != 0 {
		goto L35
	} else {
		goto L358
	}
L16:
	;
	v1019 = m.G0
	v1021 = v1019 - int32(16)
	m.G0 = v1021
	v1023 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v1028 = base.B2i32(v1026 == int32(20))
	if v1026 == int32(20) {
		goto L337
	} else {
		goto L338
	}
L17:
	;
	v829 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v830 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v831 = m.G0
	v833 = v831 - int32(128)
	m.G0 = v833
	v837 = F_table_open(m, int32(1213), int32(3))
	mBase = m.M
	v838 = m.ExcPending
	if v838 != 0 {
		goto L35
	} else {
		goto L273
	}
L18:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v681 = m.G0
	v683 = v681 - int32(48)
	m.G0 = v683
	v687 = F_table_open(m, int32(2615), int32(3))
	mBase = m.M
	v688 = m.ExcPending
	if v688 != 0 {
		goto L35
	} else {
		goto L223
	}
L19:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v275 = m.G0
	v277 = v275 - int32(256)
	m.G0 = v277
	v280 = F_strcspn(m, v274, int32(_a_F_ExecRenameStmt_0))
	mBase = m.M
	v281 = v280 + v274
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v281))))
	if v283 != 0 {
		goto L104
	} else {
		goto L105
	}
L20:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v24 = m.G0
	v26 = v24 - int32(96)
	m.G0 = v26
	v29 = F_strcspn(m, v23, int32(_a_F_ExecRenameStmt_0))
	mBase = m.M
	v30 = v29 + v23
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30))))
	if v32 != 0 {
		goto L29
	} else {
		goto L30
	}
L21:
	;
	goto L1
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L35
	} else {
		goto L92
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L35
	} else {
		goto L87
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L35
	} else {
		goto L83
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L35
	} else {
		goto L79
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L35
	} else {
		goto L75
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L35
	} else {
		goto L71
	}
L28:
	;
	if v33 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v33 = v30
	goto L31
L30:
	;
	v33 = int32(0)
	goto L31
L31:
	;
	goto L28
L32:
	;
	v38 = F_table_open(m, int32(1262), int32(3))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L35
	} else {
		goto L67
	}
L35:
	;
	return
L36:
	;
	v43 = int32(0)
	v57 = F_get_db_info(m, v22, int32(8), v26+int32(92), v43, v43, v43, v43, v43, v43, v43, v43, v43, v43, v43, v43, v43, v43)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	if v57 == int32(0) {
		goto L27
	} else {
		goto L38
	}
L38:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v26)+92))
	v64 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v65 = F_object_ownercheck(m, int32(1262), v62, v64)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	if v65 == int32(0) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	F_aclcheck_error(m, int32(2), int32(9), v22)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L35
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	v73 = F_superuser(m)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L35
	} else {
		goto L44
	}
L43:
	;
	goto L42
L44:
	;
	if v73 == int32(0) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v79 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0])))
	v80 = F_SearchSysCache1(m, int32(11), v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L35
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	v95 = F_get_database_oid(m, v23, int32(1))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L35
	} else {
		goto L52
	}
L48:
	;
	if v80 == int32(0) {
		goto L26
	} else {
		goto L49
	}
L49:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80)+16))
	v85 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+22)))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84+v85)+71)))
	F_ReleaseCatCache(m, v80)
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
		goto L35
	} else {
		goto L50
	}
L50:
	;
	if v87 == int32(0) {
		goto L26
	} else {
		goto L51
	}
L51:
	;
	goto L47
L52:
	;
	if v95 != 0 {
		goto L25
	} else {
		goto L53
	}
L53:
	;
	v98 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[1]))
	if v62 == v98 {
		goto L24
	} else {
		goto L54
	}
L54:
	;
	v104 = F_CountOtherDBBackends(m, v62, v26+int32(80), v26+int32(76))
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L35
	} else {
		goto L55
	}
L55:
	;
	if v104 != 0 {
		goto L23
	} else {
		goto L56
	}
L56:
	;
	v108 = F_SearchSysCacheLockedCopy1(m, int32(21), base.I64_extend_i32_u(v62))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L35
	} else {
		goto L57
	}
L57:
	;
	if v108 == int32(0) {
		goto L22
	} else {
		goto L58
	}
L58:
	;
	v112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v108)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+88)) = uint16(v112)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v26)+84)) = v114
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+22)))
	v122 = F_strncpy(m, v116+v117+int32(4), v23, int32(64))
	mBase = m.M
	v123 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v122)+63)) = uint8(v123)
	goto L59
L59:
	;
	v126 = v26 + int32(84)
	F_CatalogTupleUpdate(m, v38, v126, v108)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L35
	} else {
		goto L60
	}
L60:
	;
	F_UnlockTuple(m, v38, v126, int32(7))
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L35
	} else {
		goto L61
	}
L61:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[2]))
	if v133 != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v135 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1262), v62, v135, v135, v135)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L35
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v140 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v140
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v62
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1262)
	F_relation_close(m, v38, v140)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L35
	} else {
		goto L66
	}
L65:
	;
	goto L64
L66:
	;
	m.G0 = v26 + int32(96)
	goto L21
L67:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L35
	} else {
		goto L68
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v23
	F_errmsg(m, int32(_a_F_ExecRenameStmt_1), v26-int32(-64))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L35
	} else {
		goto L69
	}
L69:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_2), int32(1928), int32(_a_F_ExecRenameStmt_3))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L35
	} else {
		goto L70
	}
L70:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L71:
	;
	F_errcode(m, int32(1283))
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L35
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v22
	F_errmsg(m, int32(_a_F_ExecRenameStmt_4), v26+int32(48))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L35
	} else {
		goto L73
	}
L73:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_2), int32(1940), int32(_a_F_ExecRenameStmt_3))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L35
	} else {
		goto L74
	}
L74:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L75:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L35
	} else {
		goto L76
	}
L76:
	;
	F_errmsg(m, int32(_a_F_ExecRenameStmt_5), int32(0))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L35
	} else {
		goto L77
	}
L77:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_2), int32(1951), int32(_a_F_ExecRenameStmt_3))
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L35
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	F_errcode(m, int32(67240068))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L35
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v23
	F_errmsg(m, int32(_a_F_ExecRenameStmt_6), v26+int32(32))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L35
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_2), int32(1969), int32(_a_F_ExecRenameStmt_3))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L35
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L83:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L35
	} else {
		goto L84
	}
L84:
	;
	F_errmsg(m, int32(_a_F_ExecRenameStmt_7), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L35
	} else {
		goto L85
	}
L85:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_2), int32(1980), int32(_a_F_ExecRenameStmt_3))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L35
	} else {
		goto L86
	}
L86:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L87:
	;
	F_errcode(m, int32(100663621))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L35
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26))) = v22
	F_errmsg(m, int32(_a_F_ExecRenameStmt_8), v26)
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L35
	} else {
		goto L89
	}
L89:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v26)+80))
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v26)+76))
	F_errdetail_busy_db(m, v249, v250)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L35
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_2), int32(1993), int32(_a_F_ExecRenameStmt_3))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L35
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+16)) = v62
	F_errmsg_internal(m, int32(_a_F_ExecRenameStmt_9), v26+int32(16))
	mBase = m.M
	v267 = m.ExcPending
	if v267 != 0 {
		goto L35
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_2), int32(1998), int32(_a_F_ExecRenameStmt_3))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L35
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	goto L1
L96:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L35
	} else {
		goto L218
	}
L97:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L35
	} else {
		goto L214
	}
L98:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v616 = m.ExcPending
	if v616 != 0 {
		goto L35
	} else {
		goto L209
	}
L99:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L35
	} else {
		goto L204
	}
L100:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L35
	} else {
		goto L200
	}
L101:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L35
	} else {
		goto L196
	}
L102:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L35
	} else {
		goto L192
	}
L103:
	;
	if v284 == int32(0) {
		goto L107
	} else {
		goto L108
	}
L104:
	;
	v284 = v281
	goto L106
L105:
	;
	v284 = int32(0)
	goto L106
L106:
	;
	goto L103
L107:
	;
	v289 = F_table_open(m, int32(1260), int32(3))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L35
	} else {
		goto L110
	}
L108:
	;
	goto L109
L109:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v528 = m.ExcPending
	if v528 != 0 {
		goto L35
	} else {
		goto L188
	}
L110:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v289)+52))
	v294 = F_SearchSysCache1(m, int32(10), base.I64_extend_i32_u(v273))
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L35
	} else {
		goto L111
	}
L111:
	;
	if v294 == int32(0) {
		goto L102
	} else {
		goto L112
	}
L112:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	v299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v298)+22)))
	v300 = v298 + v299
	v301 = *(*int32)(unsafe.Add(mBase, uint32(v300)))
	v303 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[3]))
	if v301 == v303 {
		goto L101
	} else {
		goto L113
	}
L113:
	;
	v306 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[4]))
	if v301 == v306 {
		goto L100
	} else {
		goto L114
	}
L114:
	;
	v309 = v300 + int32(4)
	v310 = int32(0)
	v311 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309))))
	if v311 != int32(112) {
		v320 = v310
		goto L116
	} else {
		goto L117
	}
L115:
	;
	if v320 != 0 {
		goto L99
	} else {
		goto L119
	}
L116:
	;
	goto L115
L117:
	;
	v314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+1)))
	if v314 != int32(103) {
		v320 = v310
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+2)))
	v320 = base.B2i32(v317 == int32(95))
	goto L116
L119:
	;
	v321 = int32(0)
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	if v322 != int32(112) {
		v331 = v321
		goto L121
	} else {
		goto L122
	}
L120:
	;
	if v331 != 0 {
		goto L98
	} else {
		goto L124
	}
L121:
	;
	goto L120
L122:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+1)))
	if v325 != int32(103) {
		v331 = v321
		goto L121
	} else {
		goto L123
	}
L123:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274)+2)))
	v331 = base.B2i32(v328 == int32(95))
	goto L121
L124:
	;
	v333 = base.I64_extend_i32_u(v274)
	v334 = int64(0)
	v337 = F_SearchSysCacheExists(m, int32(10), v333, v334, v334, v334)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L35
	} else {
		goto L125
	}
L125:
	;
	if v337 != 0 {
		goto L97
	} else {
		goto L126
	}
L126:
	;
	v339 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v300)+68)))
	if v339 == int32(1) {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v277)+120)) = int64(0)
	v383 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+128)) = v383
	v385 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+121)) = uint8(v385)
	v389 = F_DirectFunctionCall1Coll(m, int32(534), v383, v333)
	mBase = m.M
	v390 = m.ExcPending
	if v390 != 0 {
		goto L35
	} else {
		goto L142
	}
L128:
	;
	v342 = F_superuser(m)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L35
	} else {
		goto L131
	}
L129:
	;
	goto L130
L130:
	;
	v370 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v371 = F_has_createrole_privilege(m, v370)
	mBase = m.M
	v372 = m.ExcPending
	if v372 != 0 {
		goto L35
	} else {
		goto L138
	}
L131:
	;
	if v342 != 0 {
		goto L127
	} else {
		goto L132
	}
L132:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L35
	} else {
		goto L133
	}
L133:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L35
	} else {
		goto L134
	}
L134:
	;
	F_errmsg(m, int32(_a_F_ExecRenameStmt_10), int32(0))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L35
	} else {
		goto L135
	}
L135:
	;
	v355 = int32(_a_F_ExecRenameStmt_11)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+84)) = v355
	*(*int32)(unsafe.Add(mBase, uint32(v277)+80)) = v355
	v362 = F_errdetail(m, int32(_a_F_ExecRenameStmt_12), v277+int32(80))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L35
	} else {
		goto L136
	}
L136:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1436), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v368 = m.ExcPending
	if v368 != 0 {
		goto L35
	} else {
		goto L137
	}
L137:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L138:
	;
	if v371 == int32(0) {
		goto L96
	} else {
		goto L139
	}
L139:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v377 = F_is_admin_of_role(m, v376, v301)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L35
	} else {
		goto L140
	}
L140:
	;
	if v377 == int32(0) {
		goto L96
	} else {
		goto L141
	}
L141:
	;
	goto L127
L142:
	;
	v391 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+133)) = uint8(v391)
	*(*int64)(unsafe.Add(mBase, uint32(v277)+152)) = v389
	v394 = *(*int32)(unsafe.Add(mBase, uint32(v294)+16))
	v395 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v394)+18)))
	if base.Ui32(int32(11)) <= base.Ui32(v395&int32(2047)) {
		goto L145
	} else {
		goto L146
	}
L143:
	;
	v500 = F_heap_modify_tuple(m, v294, v291, v277+int32(144), v277+int32(132), v277+int32(120))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L35
	} else {
		goto L180
	}
L144:
	;
	v464 = F_text_to_cstring(m, base.I32_wrap_i64(v462))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L35
	} else {
		goto L173
	}
L145:
	;
	v400 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+255)) = uint8(v400)
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+20)))
	if v402&int32(1) == v400 {
		goto L148
	} else {
		goto L149
	}
L146:
	;
	goto L147
L147:
	;
	v457 = F_getmissingattr(m, v291, int32(11), v277+int32(255))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L35
	} else {
		goto L171
	}
L148:
	;
	v407 = int32(*(*int16)(unsafe.Add(mBase, uint32(v291)+108)))
	if int32(0) <= v407 {
		goto L151
	} else {
		goto L152
	}
L149:
	;
	goto L150
L150:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+24)))
	if v444&int32(4) == int32(0) {
		goto L167
	} else {
		goto L168
	}
L151:
	;
	v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v394)+22)))
	v412 = v394 + v410 + v407
	v413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v291)+112)))
	if v413 == int32(1) {
		goto L154
	} else {
		goto L155
	}
L152:
	;
	goto L153
L153:
	;
	v442 = F_nocachegetattr(m, v294, int32(11), v291)
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L35
	} else {
		goto L166
	}
L154:
	;
	v416 = int32(*(*int16)(unsafe.Add(mBase, uint32(v291)+110)))
	if base.I32_popcnt(v416) != int32(1) {
		goto L157
	} else {
		goto L158
	}
L155:
	;
	goto L156
L156:
	;
	v462 = base.I64_extend_i32_u(v412)
	goto L144
L157:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L35
	} else {
		goto L163
	}
L158:
	;
	switch base.I32_ctz(v416) {
	case 0:
		goto L162
	case 1:
		goto L161
	case 2:
		goto L160
	case 3:
		goto L159
	default:
		goto L157
	}
L159:
	;
	v424 = *(*int64)(unsafe.Add(mBase, uint32(v412)))
	v462 = v424
	goto L144
L160:
	;
	v423 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v412))))
	v462 = v423
	goto L144
L161:
	;
	v422 = int64(*(*int16)(unsafe.Add(mBase, uint32(v412))))
	v462 = v422
	goto L144
L162:
	;
	v421 = int64(*(*int8)(unsafe.Add(mBase, uint32(v412))))
	v462 = v421
	goto L144
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+64)) = v416
	F_errmsg_internal(m, int32(_a_F_ExecRenameStmt_15), v277-int32(-64))
	mBase = m.M
	v434 = m.ExcPending
	if v434 != 0 {
		goto L35
	} else {
		goto L164
	}
L164:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_16), int32(123), int32(_a_F_ExecRenameStmt_17))
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L35
	} else {
		goto L165
	}
L165:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L166:
	;
	v462 = v442
	goto L144
L167:
	;
	v449 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+255)) = uint8(v449)
	goto L143
L168:
	;
	goto L169
L169:
	;
	v452 = F_nocachegetattr(m, v294, int32(11), v291)
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L35
	} else {
		goto L170
	}
L170:
	;
	v462 = v452
	goto L144
L171:
	;
	v459 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v277)+255)))
	if v459 != 0 {
		goto L143
	} else {
		goto L172
	}
L172:
	;
	v462 = v457
	goto L144
L173:
	;
	v466 = F_get_password_type(m, v464)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L35
	} else {
		goto L174
	}
L174:
	;
	if v466 != int32(1) {
		goto L143
	} else {
		goto L175
	}
L175:
	;
	v470 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+142)) = uint8(v470)
	*(*uint8)(unsafe.Add(mBase, uint32(v277)+130)) = uint8(v470)
	v476 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L35
	} else {
		goto L176
	}
L176:
	;
	if v476 == int32(0) {
		goto L143
	} else {
		goto L177
	}
L177:
	;
	F_errmsg(m, int32(_a_F_ExecRenameStmt_18), int32(0))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L35
	} else {
		goto L178
	}
L178:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1467), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v488 = m.ExcPending
	if v488 != 0 {
		goto L35
	} else {
		goto L179
	}
L179:
	;
	goto L143
L180:
	;
	F_CatalogTupleUpdate(m, v289, v294+int32(4), v500)
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L35
	} else {
		goto L181
	}
L181:
	;
	v505 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[2]))
	if v505 != 0 {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v507 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1260), v301, v507, v507, v507)
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L35
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v301
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1260)
	F_ReleaseCatCache(m, v294)
	mBase = m.M
	v518 = m.ExcPending
	if v518 != 0 {
		goto L35
	} else {
		goto L186
	}
L185:
	;
	goto L184
L186:
	;
	F_relation_close(m, v289, int32(0))
	mBase = m.M
	v521 = m.ExcPending
	if v521 != 0 {
		goto L35
	} else {
		goto L187
	}
L187:
	;
	m.G0 = v277 + int32(256)
	goto L95
L188:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v531 = m.ExcPending
	if v531 != 0 {
		goto L35
	} else {
		goto L189
	}
L189:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+112)) = v274
	F_errmsg(m, int32(_a_F_ExecRenameStmt_19), v277+int32(112))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L35
	} else {
		goto L190
	}
L190:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1361), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v542 = m.ExcPending
	if v542 != 0 {
		goto L35
	} else {
		goto L191
	}
L191:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L192:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L35
	} else {
		goto L193
	}
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277))) = v273
	F_errmsg(m, int32(_a_F_ExecRenameStmt_20), v277)
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L35
	} else {
		goto L194
	}
L194:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1370), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L35
	} else {
		goto L195
	}
L195:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L196:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L35
	} else {
		goto L197
	}
L197:
	;
	F_errmsg(m, int32(_a_F_ExecRenameStmt_21), int32(0))
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L35
	} else {
		goto L198
	}
L198:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1386), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L35
	} else {
		goto L199
	}
L199:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L200:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v581 = m.ExcPending
	if v581 != 0 {
		goto L35
	} else {
		goto L201
	}
L201:
	;
	F_errmsg(m, int32(_a_F_ExecRenameStmt_22), int32(0))
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L35
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1390), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v590 = m.ExcPending
	if v590 != 0 {
		goto L35
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	F_errcode(m, int32(151818372))
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L35
	} else {
		goto L205
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+16)) = v309
	F_errmsg(m, int32(_a_F_ExecRenameStmt_23), v277+int32(16))
	mBase = m.M
	v603 = m.ExcPending
	if v603 != 0 {
		goto L35
	} else {
		goto L206
	}
L206:
	;
	v606 = F_errdetail(m, int32(_a_F_ExecRenameStmt_24), int32(0))
	mBase = m.M
	v607 = m.ExcPending
	if v607 != 0 {
		goto L35
	} else {
		goto L207
	}
L207:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1401), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L35
	} else {
		goto L208
	}
L208:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L209:
	;
	F_errcode(m, int32(151818372))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L35
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+32)) = v274
	F_errmsg(m, int32(_a_F_ExecRenameStmt_23), v277+int32(32))
	mBase = m.M
	v625 = m.ExcPending
	if v625 != 0 {
		goto L35
	} else {
		goto L211
	}
L211:
	;
	v628 = F_errdetail(m, int32(_a_F_ExecRenameStmt_24), int32(0))
	mBase = m.M
	v629 = m.ExcPending
	if v629 != 0 {
		goto L35
	} else {
		goto L212
	}
L212:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1408), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L35
	} else {
		goto L213
	}
L213:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L214:
	;
	F_errcode(m, int32(_a_F_ExecRenameStmt_25))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L35
	} else {
		goto L215
	}
L215:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+48)) = v274
	F_errmsg(m, int32(_a_F_ExecRenameStmt_26), v277+int32(48))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L35
	} else {
		goto L216
	}
L216:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1423), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L35
	} else {
		goto L217
	}
L217:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L218:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v659 = m.ExcPending
	if v659 != 0 {
		goto L35
	} else {
		goto L219
	}
L219:
	;
	F_errmsg(m, int32(_a_F_ExecRenameStmt_10), int32(0))
	mBase = m.M
	v663 = m.ExcPending
	if v663 != 0 {
		goto L35
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v277)+104)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v277)+100)) = int32(_a_F_ExecRenameStmt_27)
	*(*int32)(unsafe.Add(mBase, uint32(v277)+96)) = int32(_a_F_ExecRenameStmt_28)
	v672 = F_errdetail(m, int32(_a_F_ExecRenameStmt_29), v277+int32(96))
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L35
	} else {
		goto L221
	}
L221:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_13), int32(1446), int32(_a_F_ExecRenameStmt_14))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L35
	} else {
		goto L222
	}
L222:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L223:
	;
	v692 = F_SearchSysCacheCopy(m, int32(37), base.I64_extend_i32_u(v679), int64(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L35
	} else {
		goto L227
	}
L224:
	;
	goto L1
L225:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L35
	} else {
		goto L268
	}
L226:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L35
	} else {
		goto L264
	}
L227:
	;
	if v692 != 0 {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v692)+16))
	v695 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694)+22)))
	v696 = v694 + v695
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v696)))
	v699 = F_get_namespace_oid(m, v680, int32(1))
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L35
	} else {
		goto L231
	}
L229:
	;
	goto L230
L230:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L35
	} else {
		goto L260
	}
L231:
	;
	if v699 != 0 {
		goto L226
	} else {
		goto L232
	}
L232:
	;
	v703 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v704 = F_object_ownercheck(m, int32(2615), v697, v703)
	mBase = m.M
	v705 = m.ExcPending
	if v705 != 0 {
		goto L35
	} else {
		goto L233
	}
L233:
	;
	if v704 == int32(0) {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	F_aclcheck_error(m, int32(2), int32(37), v679)
	mBase = m.M
	v711 = m.ExcPending
	if v711 != 0 {
		goto L35
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v714 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[1]))
	v716 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v718 = F_object_aclcheck(m, int32(1262), v714, v716, int64(512))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L35
	} else {
		goto L238
	}
L237:
	;
	goto L236
L238:
	;
	if v718 != 0 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v722 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[1]))
	v723 = F_get_database_name(m, v722)
	mBase = m.M
	v724 = m.ExcPending
	if v724 != 0 {
		goto L35
	} else {
		goto L242
	}
L240:
	;
	goto L241
L241:
	;
	v728 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecRenameStmt[5])))
	if v728 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	F_aclcheck_error(m, v718, int32(9), v723)
	mBase = m.M
	v726 = m.ExcPending
	if v726 != 0 {
		goto L35
	} else {
		goto L243
	}
L243:
	;
	goto L241
L244:
	;
	v731 = int32(0)
	v732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680))))
	if v732 != int32(112) {
		v741 = v731
		goto L248
	} else {
		goto L249
	}
L245:
	;
	goto L246
L246:
	;
	v745 = F_strncpy(m, v696+int32(4), v680, int32(64))
	mBase = m.M
	v746 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v745)+63)) = uint8(v746)
	goto L252
L247:
	;
	if v741 != 0 {
		goto L225
	} else {
		goto L251
	}
L248:
	;
	goto L247
L249:
	;
	v735 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+1)))
	if v735 != int32(103) {
		v741 = v731
		goto L248
	} else {
		goto L250
	}
L250:
	;
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v680)+2)))
	v741 = base.B2i32(v738 == int32(95))
	goto L248
L251:
	;
	goto L246
L252:
	;
	F_CatalogTupleUpdate(m, v687, v692+int32(4), v692)
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L35
	} else {
		goto L253
	}
L253:
	;
	v753 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[2]))
	if v753 != 0 {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v755 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2615), v697, v755, v755, v755)
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L35
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v760 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v760
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v697
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2615)
	F_relation_close(m, v687, v760)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L35
	} else {
		goto L258
	}
L257:
	;
	goto L256
L258:
	;
	F_pfree(m, v692)
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L35
	} else {
		goto L259
	}
L259:
	;
	m.G0 = v683 + int32(48)
	goto L224
L260:
	;
	F_errcode(m, int32(1411))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L35
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v683))) = v679
	F_errmsg(m, int32(_a_F_ExecRenameStmt_30), v683)
	mBase = m.M
	v783 = m.ExcPending
	if v783 != 0 {
		goto L35
	} else {
		goto L262
	}
L262:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_31), int32(266), int32(_a_F_ExecRenameStmt_32))
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L35
	} else {
		goto L263
	}
L263:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L264:
	;
	F_errcode(m, int32(100794500))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L35
	} else {
		goto L265
	}
L265:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v683)+32)) = v680
	F_errmsg(m, int32(_a_F_ExecRenameStmt_33), v683+int32(32))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L35
	} else {
		goto L266
	}
L266:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_31), int32(275), int32(_a_F_ExecRenameStmt_32))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L35
	} else {
		goto L267
	}
L267:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L268:
	;
	F_errcode(m, int32(151818372))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L35
	} else {
		goto L269
	}
L269:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v683)+16)) = v680
	F_errmsg(m, int32(_a_F_ExecRenameStmt_34), v683+int32(16))
	mBase = m.M
	v819 = m.ExcPending
	if v819 != 0 {
		goto L35
	} else {
		goto L270
	}
L270:
	;
	v822 = F_errdetail(m, int32(_a_F_ExecRenameStmt_35), int32(0))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L35
	} else {
		goto L271
	}
L271:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_31), int32(292), int32(_a_F_ExecRenameStmt_32))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L35
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L273:
	;
	v840 = v833 - int32(-64)
	F_ScanKeyInit(m, v840, int32(2), int32(3), int32(62), base.I64_extend_i32_u(v829))
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L35
	} else {
		goto L274
	}
L274:
	;
	v848 = F_table_beginscan_catalog(m, v837, int32(1), v840)
	mBase = m.M
	v849 = m.ExcPending
	if v849 != 0 {
		goto L35
	} else {
		goto L279
	}
L275:
	;
	goto L1
L276:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1004 = m.ExcPending
	if v1004 != 0 {
		goto L35
	} else {
		goto L330
	}
L277:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v986 = m.ExcPending
	if v986 != 0 {
		goto L35
	} else {
		goto L326
	}
L278:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L35
	} else {
		goto L321
	}
L279:
	;
	v850 = F_heap_getnext(m, v848)
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L35
	} else {
		goto L280
	}
L280:
	;
	if v850 != 0 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v852 = F_heap_copytuple(m, v850)
	mBase = m.M
	v853 = m.ExcPending
	if v853 != 0 {
		goto L35
	} else {
		goto L284
	}
L282:
	;
	goto L283
L283:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v948 = m.ExcPending
	if v948 != 0 {
		goto L35
	} else {
		goto L317
	}
L284:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v852)+16))
	v855 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v854)+22)))
	v856 = v854 + v855
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v856)))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v848)))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v858)+188))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v859)+12))
	m.T0[v860].(func(*base.Module, int32))(m, v848)
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L35
	} else {
		goto L285
	}
L285:
	;
	F_shdepLockAndCheckObject(m, int32(1213), v857)
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L35
	} else {
		goto L286
	}
L286:
	;
	v868 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v869 = F_object_ownercheck(m, int32(1213), v857, v868)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L35
	} else {
		goto L287
	}
L287:
	;
	if v869 == int32(0) {
		goto L288
	} else {
		goto L289
	}
L288:
	;
	F_aclcheck_error(m, int32(1), int32(43), v829)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L35
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v878 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecRenameStmt[5])))
	if v878 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L291:
	;
	goto L290
L292:
	;
	v881 = int32(0)
	v882 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830))))
	if v882 != int32(112) {
		v891 = v881
		goto L296
	} else {
		goto L297
	}
L293:
	;
	goto L294
L294:
	;
	v893 = F_strcspn(m, v830, int32(_a_F_ExecRenameStmt_0))
	mBase = m.M
	v894 = v893 + v830
	v896 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v894))))
	if v896 != 0 {
		goto L301
	} else {
		goto L302
	}
L295:
	;
	if v891 != 0 {
		goto L278
	} else {
		goto L299
	}
L296:
	;
	goto L295
L297:
	;
	v885 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830)+1)))
	if v885 != int32(103) {
		v891 = v881
		goto L296
	} else {
		goto L298
	}
L298:
	;
	v888 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v830)+2)))
	v891 = base.B2i32(v888 == int32(95))
	goto L296
L299:
	;
	goto L294
L300:
	;
	if v897 != 0 {
		goto L277
	} else {
		goto L304
	}
L301:
	;
	v897 = v894
	goto L303
L302:
	;
	v897 = int32(0)
	goto L303
L303:
	;
	goto L300
L304:
	;
	v899 = v833 - int32(-64)
	F_ScanKeyInit(m, v899, int32(2), int32(3), int32(62), base.I64_extend_i32_u(v830))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L35
	} else {
		goto L305
	}
L305:
	;
	v907 = F_table_beginscan_catalog(m, v837, int32(1), v899)
	mBase = m.M
	v908 = m.ExcPending
	if v908 != 0 {
		goto L35
	} else {
		goto L306
	}
L306:
	;
	v909 = F_heap_getnext(m, v907)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L35
	} else {
		goto L307
	}
L307:
	;
	if v909 != 0 {
		goto L276
	} else {
		goto L308
	}
L308:
	;
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v907)))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v911)+188))
	v913 = *(*int32)(unsafe.Add(mBase, uint32(v912)+12))
	m.T0[v913].(func(*base.Module, int32))(m, v907)
	mBase = m.M
	v915 = m.ExcPending
	if v915 != 0 {
		goto L35
	} else {
		goto L309
	}
L309:
	;
	v919 = F_strncpy(m, v856+int32(4), v830, int32(64))
	mBase = m.M
	v920 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v919)+63)) = uint8(v920)
	goto L310
L310:
	;
	F_CatalogTupleUpdate(m, v837, v852+int32(4), v852)
	mBase = m.M
	v925 = m.ExcPending
	if v925 != 0 {
		goto L35
	} else {
		goto L311
	}
L311:
	;
	v927 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[2]))
	if v927 != 0 {
		goto L312
	} else {
		goto L313
	}
L312:
	;
	v929 = int32(0)
	F_RunObjectPostAlterHook(m, int32(1213), v857, v929, v929, v929)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L35
	} else {
		goto L315
	}
L313:
	;
	goto L314
L314:
	;
	v934 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v934
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1213)
	F_relation_close(m, v837, v934)
	mBase = m.M
	v941 = m.ExcPending
	if v941 != 0 {
		goto L35
	} else {
		goto L316
	}
L315:
	;
	goto L314
L316:
	;
	m.G0 = v833 + int32(128)
	goto L275
L317:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v951 = m.ExcPending
	if v951 != 0 {
		goto L35
	} else {
		goto L318
	}
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v833))) = v829
	F_errmsg(m, int32(_a_F_ExecRenameStmt_36), v833)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L35
	} else {
		goto L319
	}
L319:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_37), int32(967), int32(_a_F_ExecRenameStmt_38))
	mBase = m.M
	v960 = m.ExcPending
	if v960 != 0 {
		goto L35
	} else {
		goto L320
	}
L320:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L321:
	;
	F_errcode(m, int32(151818372))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L35
	} else {
		goto L322
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v833)+48)) = v830
	F_errmsg(m, int32(_a_F_ExecRenameStmt_39), v833+int32(48))
	mBase = m.M
	v973 = m.ExcPending
	if v973 != 0 {
		goto L35
	} else {
		goto L323
	}
L323:
	;
	v976 = F_errdetail(m, int32(_a_F_ExecRenameStmt_40), int32(0))
	mBase = m.M
	v977 = m.ExcPending
	if v977 != 0 {
		goto L35
	} else {
		goto L324
	}
L324:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_37), int32(987), int32(_a_F_ExecRenameStmt_38))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L35
	} else {
		goto L325
	}
L325:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L326:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v989 = m.ExcPending
	if v989 != 0 {
		goto L35
	} else {
		goto L327
	}
L327:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v833)+32)) = v830
	F_errmsg(m, int32(_a_F_ExecRenameStmt_41), v833+int32(32))
	mBase = m.M
	v995 = m.ExcPending
	if v995 != 0 {
		goto L35
	} else {
		goto L328
	}
L328:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_37), int32(993), int32(_a_F_ExecRenameStmt_38))
	mBase = m.M
	v1000 = m.ExcPending
	if v1000 != 0 {
		goto L35
	} else {
		goto L329
	}
L329:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L330:
	;
	F_errcode(m, int32(_a_F_ExecRenameStmt_25))
	mBase = m.M
	v1007 = m.ExcPending
	if v1007 != 0 {
		goto L35
	} else {
		goto L331
	}
L331:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v833)+16)) = v830
	F_errmsg(m, int32(_a_F_ExecRenameStmt_42), v833+int32(16))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L35
	} else {
		goto L332
	}
L332:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_37), int32(1015), int32(_a_F_ExecRenameStmt_38))
	mBase = m.M
	v1018 = m.ExcPending
	if v1018 != 0 {
		goto L35
	} else {
		goto L333
	}
L333:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L334:
	;
	m.G0 = v1021 + int32(16)
	goto L1
L335:
	;
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	F_RenameRelationInternal(m, v1077, v1079, int32(0), v1078)
	mBase = m.M
	v1082 = m.ExcPending
	if v1082 != 0 {
		goto L35
	} else {
		goto L356
	}
L336:
	;
	v1074 = F_get_rel_relkind(m, v1050)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L35
	} else {
		goto L355
	}
L337:
	;
	v1029 = int32(4)
	goto L339
L338:
	;
	v1029 = int32(8)
	goto L339
L339:
	;
	v1030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v1032 = F_RangeVarGetRelidExtended(m, v1023, v1029, v1030, int32(621), l1)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L35
	} else {
		goto L340
	}
L340:
	;
	if v1032 != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v1034 = F_get_rel_relkind(m, v1032)
	mBase = m.M
	v1035 = m.ExcPending
	if v1035 != 0 {
		goto L35
	} else {
		goto L344
	}
L342:
	;
	goto L343
L343:
	;
	v1055 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1056 = m.ExcPending
	if v1056 != 0 {
		goto L35
	} else {
		goto L349
	}
L344:
	;
	if base.B2i32(v1034&int32(223) == int32(73))|base.B2i32(v1026 != int32(20)) != 0 {
		v1077 = v1032
		v1078 = v1028
		goto L335
	} else {
		goto L345
	}
L345:
	;
	F_UnlockRelationOid(m, v1032, int32(4))
	mBase = m.M
	v1045 = m.ExcPending
	if v1045 != 0 {
		goto L35
	} else {
		goto L346
	}
L346:
	;
	v1046 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1048 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v1050 = F_RangeVarGetRelidExtended(m, v1046, int32(8), v1048, int32(621), l1)
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L35
	} else {
		goto L347
	}
L347:
	;
	if v1050 != 0 {
		goto L336
	} else {
		goto L348
	}
L348:
	;
	goto L343
L349:
	;
	if v1055 != 0 {
		goto L350
	} else {
		goto L351
	}
L350:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v1057)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1021))) = v1058
	F_errmsg(m, int32(_a_F_ExecRenameStmt_43), v1021)
	mBase = m.M
	v1062 = m.ExcPending
	if v1062 != 0 {
		goto L35
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[6]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1069
	v1072 = *(*int64)(unsafe.Add(mBase, _c_F_ExecRenameStmt[7]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1072
	goto L334
L353:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_44), int32(_a_F_ExecRenameStmt_45), int32(_a_F_ExecRenameStmt_46))
	mBase = m.M
	v1067 = m.ExcPending
	if v1067 != 0 {
		goto L35
	} else {
		goto L354
	}
L354:
	;
	goto L352
L355:
	;
	v1077 = v1050
	v1078 = int32(0)
	goto L335
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1077
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	goto L334
L357:
	;
	m.G0 = v1095 + int32(16)
	goto L1
L358:
	;
	if v1102 == int32(0) {
		goto L359
	} else {
		goto L360
	}
L359:
	;
	v1108 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L35
	} else {
		goto L362
	}
L360:
	;
	goto L361
L361:
	;
	v1127 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1129)+16)))
	v1131 = int32(0)
	v1133 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v1134 = F_renameatt_internal(m, v1102, v1127, v1128, v1130, v1131, v1131, v1133)
	mBase = m.M
	v1135 = m.ExcPending
	if v1135 != 0 {
		goto L35
	} else {
		goto L368
	}
L362:
	;
	if v1108 != 0 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1110)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v1095))) = v1111
	F_errmsg(m, int32(_a_F_ExecRenameStmt_43), v1095)
	mBase = m.M
	v1115 = m.ExcPending
	if v1115 != 0 {
		goto L35
	} else {
		goto L366
	}
L364:
	;
	goto L365
L365:
	;
	v1122 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[6]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1122
	v1125 = *(*int64)(unsafe.Add(mBase, _c_F_ExecRenameStmt[7]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v1125
	goto L357
L366:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_44), int32(4058), int32(_a_F_ExecRenameStmt_47))
	mBase = m.M
	v1120 = m.ExcPending
	if v1120 != 0 {
		goto L35
	} else {
		goto L367
	}
L367:
	;
	goto L365
L368:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1134
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1102
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1259)
	goto L357
L369:
	;
	v1157 = F_relation_open(m, v1154, int32(0))
	mBase = m.M
	v1158 = m.ExcPending
	if v1158 != 0 {
		goto L35
	} else {
		goto L370
	}
L370:
	;
	v1161 = F_table_open(m, int32(2618), int32(3))
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L35
	} else {
		goto L371
	}
L371:
	;
	v1166 = F_SearchSysCacheCopy(m, int32(60), base.I64_extend_i32_u(v1154), base.I64_extend_i32_u(v1144))
	mBase = m.M
	v1167 = m.ExcPending
	if v1167 != 0 {
		goto L35
	} else {
		goto L376
	}
L372:
	;
	goto L1
L373:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1308 = m.ExcPending
	if v1308 != 0 {
		goto L35
	} else {
		goto L413
	}
L374:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1292 = m.ExcPending
	if v1292 != 0 {
		goto L35
	} else {
		goto L409
	}
L375:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1270 = m.ExcPending
	if v1270 != 0 {
		goto L35
	} else {
		goto L405
	}
L376:
	;
	if v1166 != 0 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1168 = *(*int32)(unsafe.Add(mBase, uint32(v1166)+16))
	v1169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1168)+22)))
	v1170 = v1168 + v1169
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v1170)))
	v1175 = int64(0)
	v1177 = F_SearchSysCacheExists(m, int32(60), base.I64_extend_i32_u(v1154), base.I64_extend_i32_u(v1145), v1175, v1175)
	mBase = m.M
	v1178 = m.ExcPending
	if v1178 != 0 {
		goto L35
	} else {
		goto L380
	}
L378:
	;
	goto L379
L379:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1250 = m.ExcPending
	if v1250 != 0 {
		goto L35
	} else {
		goto L401
	}
L380:
	;
	if v1177 != 0 {
		goto L375
	} else {
		goto L381
	}
L381:
	;
	v1179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1170)+72)))
	if v1179 == int32(49) {
		goto L374
	} else {
		goto L382
	}
L382:
	;
	v1182 = int32(_a_F_ExecRenameStmt_48)
	v1185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145))))
	v1188 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExecRenameStmt[8])))
	if base.B2i32(v1185 == int32(0))|base.B2i32(v1185 != v1188) != 0 {
		v1206 = v1185
		v1207 = v1188
		goto L384
	} else {
		goto L385
	}
L383:
	;
	if v1206-v1207 == int32(0) {
		goto L373
	} else {
		goto L390
	}
L384:
	;
	goto L383
L385:
	;
	v1191 = v1145
	v1192 = v1182
	goto L386
L386:
	;
	v1195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1192)+1)))
	v1196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1191)+1)))
	if v1196 == int32(0) {
		v1206 = v1196
		v1207 = v1195
		goto L384
	} else {
		goto L388
	}
L387:
	;
	v1206 = v1196
	v1207 = v1195
	goto L384
L388:
	;
	v1199 = int32(1)
	if v1196 == v1195 {
		v1191 = v1191 + v1199
		v1192 = v1192 + v1199
		goto L386
	} else {
		goto L389
	}
L389:
	;
	goto L387
L390:
	;
	v1214 = F_strncpy(m, v1170+int32(4), v1145, int32(64))
	mBase = m.M
	v1215 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1214)+63)) = uint8(v1215)
	goto L391
L391:
	;
	F_CatalogTupleUpdate(m, v1161, v1166+int32(4), v1166)
	mBase = m.M
	v1220 = m.ExcPending
	if v1220 != 0 {
		goto L35
	} else {
		goto L392
	}
L392:
	;
	v1222 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[2]))
	if v1222 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v1224 = int32(0)
	F_RunObjectPostAlterHook(m, int32(2618), v1171, v1224, v1224, v1224)
	mBase = m.M
	v1228 = m.ExcPending
	if v1228 != 0 {
		goto L35
	} else {
		goto L396
	}
L394:
	;
	goto L395
L395:
	;
	F_pfree(m, v1166)
	mBase = m.M
	v1230 = m.ExcPending
	if v1230 != 0 {
		goto L35
	} else {
		goto L397
	}
L396:
	;
	goto L395
L397:
	;
	F_relation_close(m, v1161, int32(3))
	mBase = m.M
	v1233 = m.ExcPending
	if v1233 != 0 {
		goto L35
	} else {
		goto L398
	}
L398:
	;
	F_CacheInvalidateRelcache(m, v1157)
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L35
	} else {
		goto L399
	}
L399:
	;
	v1236 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v1236
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1171
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2618)
	F_relation_close(m, v1157, v1236)
	mBase = m.M
	v1243 = m.ExcPending
	if v1243 != 0 {
		goto L35
	} else {
		goto L400
	}
L400:
	;
	m.G0 = v1148 + int32(48)
	goto L372
L401:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1253 = m.ExcPending
	if v1253 != 0 {
		goto L35
	} else {
		goto L402
	}
L402:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1148))) = v1144
	*(*int32)(unsafe.Add(mBase, uint32(v1148)+4)) = v1254 + int32(4)
	F_errmsg(m, int32(_a_F_ExecRenameStmt_49), v1148)
	mBase = m.M
	v1261 = m.ExcPending
	if v1261 != 0 {
		goto L35
	} else {
		goto L403
	}
L403:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_50), int32(813), int32(_a_F_ExecRenameStmt_51))
	mBase = m.M
	v1266 = m.ExcPending
	if v1266 != 0 {
		goto L35
	} else {
		goto L404
	}
L404:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L405:
	;
	F_errcode(m, int32(_a_F_ExecRenameStmt_25))
	mBase = m.M
	v1273 = m.ExcPending
	if v1273 != 0 {
		goto L35
	} else {
		goto L406
	}
L406:
	;
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1148)+16)) = v1145
	*(*int32)(unsafe.Add(mBase, uint32(v1148)+20)) = v1274 + int32(4)
	F_errmsg(m, int32(_a_F_ExecRenameStmt_52), v1148+int32(16))
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L35
	} else {
		goto L407
	}
L407:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_50), int32(822), int32(_a_F_ExecRenameStmt_51))
	mBase = m.M
	v1288 = m.ExcPending
	if v1288 != 0 {
		goto L35
	} else {
		goto L408
	}
L408:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L409:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L35
	} else {
		goto L410
	}
L410:
	;
	F_errmsg(m, int32(_a_F_ExecRenameStmt_53), int32(0))
	mBase = m.M
	v1299 = m.ExcPending
	if v1299 != 0 {
		goto L35
	} else {
		goto L411
	}
L411:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_50), int32(831), int32(_a_F_ExecRenameStmt_51))
	mBase = m.M
	v1304 = m.ExcPending
	if v1304 != 0 {
		goto L35
	} else {
		goto L412
	}
L412:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L413:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v1311 = m.ExcPending
	if v1311 != 0 {
		goto L35
	} else {
		goto L414
	}
L414:
	;
	v1312 = *(*int32)(unsafe.Add(mBase, uint32(v1157)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v1148)+36)) = int32(_a_F_ExecRenameStmt_48)
	*(*int32)(unsafe.Add(mBase, uint32(v1148)+32)) = v1312 + int32(4)
	F_errmsg(m, int32(_a_F_ExecRenameStmt_54), v1148+int32(32))
	mBase = m.M
	v1322 = m.ExcPending
	if v1322 != 0 {
		goto L35
	} else {
		goto L415
	}
L415:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_50), int32(842), int32(_a_F_ExecRenameStmt_51))
	mBase = m.M
	v1327 = m.ExcPending
	if v1327 != 0 {
		goto L35
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
	v1340 = F_relation_open(m, v1337, int32(0))
	mBase = m.M
	v1341 = m.ExcPending
	if v1341 != 0 {
		goto L35
	} else {
		goto L418
	}
L418:
	;
	v1342 = *(*int32)(unsafe.Add(mBase, uint32(v1340)+48))
	v1343 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1342)+119)))
	if v1343 == int32(112) {
		goto L419
	} else {
		goto L420
	}
L419:
	;
	v1348 = F_find_all_inheritors(m, v1337, int32(8), int32(0))
	mBase = m.M
	v1349 = m.ExcPending
	if v1349 != 0 {
		goto L35
	} else {
		goto L422
	}
L420:
	;
	goto L421
L421:
	;
	v1352 = F_table_open(m, int32(2620), int32(3))
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L35
	} else {
		goto L423
	}
L422:
	;
	goto L421
L423:
	;
	v1355 = v1330 + int32(48)
	F_ScanKeyInit(m, v1355, int32(2), int32(3), int32(184), base.I64_extend_i32_u(v1337))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L35
	} else {
		goto L424
	}
L424:
	;
	v1367 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+20)))
	F_ScanKeyInit(m, v1330+int32(104), int32(4), int32(3), int32(62), v1367)
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L35
	} else {
		goto L425
	}
L425:
	;
	v1374 = F_systable_beginscan(m, v1352, int32(2701), int32(1), int32(0), int32(2), v1355)
	mBase = m.M
	v1375 = m.ExcPending
	if v1375 != 0 {
		goto L35
	} else {
		goto L428
	}
L426:
	;
	goto L1
L427:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1480 = m.ExcPending
	if v1480 != 0 {
		goto L35
	} else {
		goto L450
	}
L428:
	;
	v1376 = F_systable_getnext(m, v1374)
	mBase = m.M
	v1377 = m.ExcPending
	if v1377 != 0 {
		goto L35
	} else {
		goto L429
	}
L429:
	;
	if v1376 != 0 {
		goto L430
	} else {
		goto L431
	}
L430:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v1376)+16))
	v1379 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1378)+22)))
	v1380 = v1378 + v1379
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+8))
	if v1381 != 0 {
		goto L427
	} else {
		goto L433
	}
L431:
	;
	goto L432
L432:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1459 = m.ExcPending
	if v1459 != 0 {
		goto L35
	} else {
		goto L446
	}
L433:
	;
	v1382 = *(*int32)(unsafe.Add(mBase, uint32(v1380)))
	v1383 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1384 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_renametrig_internal(m, v1352, v1340, v1376, v1383, v1384)
	mBase = m.M
	v1386 = m.ExcPending
	if v1386 != 0 {
		goto L35
	} else {
		goto L434
	}
L434:
	;
	v1387 = *(*int32)(unsafe.Add(mBase, uint32(v1340)+48))
	v1388 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1387)+119)))
	if v1388 != int32(112) {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1382
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(2620)
	F_systable_endscan(m, v1374)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L35
	} else {
		goto L443
	}
L436:
	;
	v1392 = F_RelationGetPartitionDesc(m, v1340, int32(1))
	mBase = m.M
	v1393 = m.ExcPending
	if v1393 != 0 {
		goto L35
	} else {
		goto L437
	}
L437:
	;
	v1394 = *(*int32)(unsafe.Add(mBase, uint32(v1392)))
	if v1394 <= int32(0) {
		goto L435
	} else {
		goto L438
	}
L438:
	;
	v1400 = int32(0)
	goto L439
L439:
	;
	v1412 = *(*int32)(unsafe.Add(mBase, uint32(v1392)+8))
	v1416 = *(*int32)(unsafe.Add(mBase, uint32(v1412+v1400<<(uint(int32(2))%32))))
	v1417 = *(*int32)(unsafe.Add(mBase, uint32(v1380)))
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	F_renametrig_partition(m, v1352, v1416, v1417, v1418, v1419)
	mBase = m.M
	v1421 = m.ExcPending
	if v1421 != 0 {
		goto L35
	} else {
		goto L441
	}
L440:
	;
	goto L435
L441:
	;
	v1423 = v1400 + int32(1)
	v1424 = *(*int32)(unsafe.Add(mBase, uint32(v1392)))
	if v1423 < v1424 {
		v1400 = v1423
		goto L439
	} else {
		goto L442
	}
L442:
	;
	goto L440
L443:
	;
	F_relation_close(m, v1352, int32(3))
	mBase = m.M
	v1449 = m.ExcPending
	if v1449 != 0 {
		goto L35
	} else {
		goto L444
	}
L444:
	;
	F_relation_close(m, v1340, int32(0))
	mBase = m.M
	v1452 = m.ExcPending
	if v1452 != 0 {
		goto L35
	} else {
		goto L445
	}
L445:
	;
	m.G0 = v1330 + int32(160)
	goto L426
L446:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L35
	} else {
		goto L447
	}
L447:
	;
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1340)+48))
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1330))) = v1464
	*(*int32)(unsafe.Add(mBase, uint32(v1330)+4)) = v1463 + int32(4)
	F_errmsg(m, int32(_a_F_ExecRenameStmt_55), v1330)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L35
	} else {
		goto L448
	}
L448:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_56), int32(1573), int32(_a_F_ExecRenameStmt_57))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L35
	} else {
		goto L449
	}
L449:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L450:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v1483 = m.ExcPending
	if v1483 != 0 {
		goto L35
	} else {
		goto L451
	}
L451:
	;
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(v1340)+48))
	v1485 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1330)+32)) = v1485
	*(*int32)(unsafe.Add(mBase, uint32(v1330)+36)) = v1484 + int32(4)
	F_errmsg(m, int32(_a_F_ExecRenameStmt_58), v1330+int32(32))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L35
	} else {
		goto L452
	}
L452:
	;
	v1496 = F_get_partition_parent(m, v1337, int32(0))
	mBase = m.M
	v1497 = m.ExcPending
	if v1497 != 0 {
		goto L35
	} else {
		goto L453
	}
L453:
	;
	v1498 = F_get_rel_name(m, v1496)
	mBase = m.M
	v1499 = m.ExcPending
	if v1499 != 0 {
		goto L35
	} else {
		goto L454
	}
L454:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1330)+16)) = v1498
	F_errhint(m, int32(_a_F_ExecRenameStmt_59), v1330+int32(16))
	mBase = m.M
	v1505 = m.ExcPending
	if v1505 != 0 {
		goto L35
	} else {
		goto L455
	}
L455:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_56), int32(1547), int32(_a_F_ExecRenameStmt_57))
	mBase = m.M
	v1510 = m.ExcPending
	if v1510 != 0 {
		goto L35
	} else {
		goto L456
	}
L456:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L457:
	;
	v1522 = F_relation_open(m, v1519, int32(0))
	mBase = m.M
	v1523 = m.ExcPending
	if v1523 != 0 {
		goto L35
	} else {
		goto L458
	}
L458:
	;
	v1526 = F_table_open(m, int32(3256), int32(3))
	mBase = m.M
	v1527 = m.ExcPending
	if v1527 != 0 {
		goto L35
	} else {
		goto L459
	}
L459:
	;
	v1529 = v1513 + int32(32)
	v1530 = int32(3)
	v1533 = base.I64_extend_i32_u(v1519)
	F_ScanKeyInit(m, v1529, v1530, v1530, int32(184), v1533)
	mBase = m.M
	v1535 = m.ExcPending
	if v1535 != 0 {
		goto L35
	} else {
		goto L460
	}
L460:
	;
	v1537 = v1513 + int32(88)
	v1541 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+24)))
	F_ScanKeyInit(m, v1537, int32(2), int32(3), int32(62), v1541)
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L35
	} else {
		goto L461
	}
L461:
	;
	v1548 = F_systable_beginscan(m, v1526, int32(3258), int32(1), int32(0), int32(2), v1529)
	mBase = m.M
	v1549 = m.ExcPending
	if v1549 != 0 {
		goto L35
	} else {
		goto L464
	}
L462:
	;
	goto L1
L463:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1649 = m.ExcPending
	if v1649 != 0 {
		goto L35
	} else {
		goto L490
	}
L464:
	;
	v1550 = F_systable_getnext(m, v1548)
	mBase = m.M
	v1551 = m.ExcPending
	if v1551 != 0 {
		goto L35
	} else {
		goto L465
	}
L465:
	;
	if v1550 == int32(0) {
		goto L466
	} else {
		goto L467
	}
L466:
	;
	F_systable_endscan(m, v1548)
	mBase = m.M
	v1555 = m.ExcPending
	if v1555 != 0 {
		goto L35
	} else {
		goto L469
	}
L467:
	;
	goto L468
L468:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1626 = m.ExcPending
	if v1626 != 0 {
		goto L35
	} else {
		goto L486
	}
L469:
	;
	v1556 = int32(3)
	F_ScanKeyInit(m, v1529, v1556, v1556, int32(184), v1533)
	mBase = m.M
	v1560 = m.ExcPending
	if v1560 != 0 {
		goto L35
	} else {
		goto L470
	}
L470:
	;
	v1564 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l1)+20)))
	F_ScanKeyInit(m, v1537, int32(2), int32(3), int32(62), v1564)
	mBase = m.M
	v1566 = m.ExcPending
	if v1566 != 0 {
		goto L35
	} else {
		goto L471
	}
L471:
	;
	v1571 = F_systable_beginscan(m, v1526, int32(3258), int32(1), int32(0), int32(2), v1529)
	mBase = m.M
	v1572 = m.ExcPending
	if v1572 != 0 {
		goto L35
	} else {
		goto L472
	}
L472:
	;
	v1573 = F_systable_getnext(m, v1571)
	mBase = m.M
	v1574 = m.ExcPending
	if v1574 != 0 {
		goto L35
	} else {
		goto L473
	}
L473:
	;
	if v1573 == int32(0) {
		goto L463
	} else {
		goto L474
	}
L474:
	;
	v1577 = *(*int32)(unsafe.Add(mBase, uint32(v1573)+16))
	v1578 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1577)+22)))
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1577+v1578)))
	v1581 = F_heap_copytuple(m, v1573)
	mBase = m.M
	v1582 = m.ExcPending
	if v1582 != 0 {
		goto L35
	} else {
		goto L475
	}
L475:
	;
	v1583 = *(*int32)(unsafe.Add(mBase, uint32(v1581)+16))
	v1584 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1583)+22)))
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1590 = F_strncpy(m, v1583+v1584+int32(4), v1588, int32(64))
	mBase = m.M
	v1591 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1590)+63)) = uint8(v1591)
	goto L476
L476:
	;
	F_CatalogTupleUpdate(m, v1526, v1581+int32(4), v1581)
	mBase = m.M
	v1596 = m.ExcPending
	if v1596 != 0 {
		goto L35
	} else {
		goto L477
	}
L477:
	;
	v1598 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[2]))
	if v1598 != 0 {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v1600 = int32(0)
	F_RunObjectPostAlterHook(m, int32(3256), v1580, v1600, v1600, v1600)
	mBase = m.M
	v1604 = m.ExcPending
	if v1604 != 0 {
		goto L35
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1580
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(3256)
	F_CacheInvalidateRelcache(m, v1522)
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L35
	} else {
		goto L482
	}
L481:
	;
	goto L480
L482:
	;
	F_systable_endscan(m, v1571)
	mBase = m.M
	v1613 = m.ExcPending
	if v1613 != 0 {
		goto L35
	} else {
		goto L483
	}
L483:
	;
	F_relation_close(m, v1526, int32(3))
	mBase = m.M
	v1616 = m.ExcPending
	if v1616 != 0 {
		goto L35
	} else {
		goto L484
	}
L484:
	;
	F_relation_close(m, v1522, int32(0))
	mBase = m.M
	v1619 = m.ExcPending
	if v1619 != 0 {
		goto L35
	} else {
		goto L485
	}
L485:
	;
	m.G0 = v1513 + int32(144)
	goto L462
L486:
	;
	F_errcode(m, int32(_a_F_ExecRenameStmt_25))
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L35
	} else {
		goto L487
	}
L487:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+48))
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1513)+16)) = v1631
	*(*int32)(unsafe.Add(mBase, uint32(v1513)+20)) = v1630 + int32(4)
	F_errmsg(m, int32(_a_F_ExecRenameStmt_60), v1513+int32(16))
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L35
	} else {
		goto L488
	}
L488:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_61), int32(1147), int32(_a_F_ExecRenameStmt_62))
	mBase = m.M
	v1645 = m.ExcPending
	if v1645 != 0 {
		goto L35
	} else {
		goto L489
	}
L489:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L490:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v1652 = m.ExcPending
	if v1652 != 0 {
		goto L35
	} else {
		goto L491
	}
L491:
	;
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1522)+48))
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v1513))) = v1654
	*(*int32)(unsafe.Add(mBase, uint32(v1513)+4)) = v1653 + int32(4)
	F_errmsg(m, int32(_a_F_ExecRenameStmt_63), v1513)
	mBase = m.M
	v1661 = m.ExcPending
	if v1661 != 0 {
		goto L35
	} else {
		goto L492
	}
L492:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_61), int32(1175), int32(_a_F_ExecRenameStmt_62))
	mBase = m.M
	v1666 = m.ExcPending
	if v1666 != 0 {
		goto L35
	} else {
		goto L493
	}
L493:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L494:
	;
	v1676 = F_typenameTypeId(m, int32(0), v1674)
	mBase = m.M
	v1677 = m.ExcPending
	if v1677 != 0 {
		goto L35
	} else {
		goto L495
	}
L495:
	;
	v1680 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v1681 = m.ExcPending
	if v1681 != 0 {
		goto L35
	} else {
		goto L496
	}
L496:
	;
	v1685 = F_SearchSysCacheCopy(m, int32(82), base.I64_extend_i32_u(v1676), int64(0))
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L35
	} else {
		goto L500
	}
L497:
	;
	goto L1
L498:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L35
	} else {
		goto L542
	}
L499:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1775 = m.ExcPending
	if v1775 != 0 {
		goto L35
	} else {
		goto L536
	}
L500:
	;
	if v1685 != 0 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(v1685)+16))
	v1688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1687)+22)))
	v1689 = v1687 + v1688
	v1692 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v1693 = F_object_ownercheck(m, int32(1247), v1676, v1692)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L35
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1762 = m.ExcPending
	if v1762 != 0 {
		goto L35
	} else {
		goto L533
	}
L504:
	;
	if v1693 == int32(0) {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	F_aclcheck_error_type(m, int32(2), v1676)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L35
	} else {
		goto L508
	}
L506:
	;
	goto L507
L507:
	;
	v1700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1689)+79)))
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v1701 == int32(12) {
		goto L510
	} else {
		goto L511
	}
L508:
	;
	goto L507
L509:
	;
	v1733 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+92))
	if v1733 != 0 {
		goto L522
	} else {
		goto L523
	}
L510:
	;
	if v1700 == int32(100) {
		goto L509
	} else {
		goto L513
	}
L511:
	;
	goto L512
L512:
	;
	if v1700 != int32(99) {
		goto L509
	} else {
		goto L519
	}
L513:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L35
	} else {
		goto L514
	}
L514:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1712 = m.ExcPending
	if v1712 != 0 {
		goto L35
	} else {
		goto L515
	}
L515:
	;
	v1713 = F_format_type_be(m, v1676)
	mBase = m.M
	v1714 = m.ExcPending
	if v1714 != 0 {
		goto L35
	} else {
		goto L516
	}
L516:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+48)) = v1713
	F_errmsg(m, int32(_a_F_ExecRenameStmt_64), v1669+int32(48))
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
		goto L35
	} else {
		goto L517
	}
L517:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_65), int32(3840), int32(_a_F_ExecRenameStmt_66))
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L35
	} else {
		goto L518
	}
L518:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L519:
	;
	v1728 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+84))
	v1729 = F_get_rel_relkind(m, v1728)
	mBase = m.M
	v1730 = m.ExcPending
	if v1730 != 0 {
		goto L35
	} else {
		goto L520
	}
L520:
	;
	if v1729 != int32(99) {
		goto L499
	} else {
		goto L521
	}
L521:
	;
	goto L509
L522:
	;
	v1734 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+88))
	if v1734 == int32(_a_F_ExecRenameStmt_67) {
		goto L498
	} else {
		goto L525
	}
L523:
	;
	goto L524
L524:
	;
	v1737 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1689)+79)))
	if v1737 == int32(99) {
		goto L527
	} else {
		goto L528
	}
L525:
	;
	goto L524
L526:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1676
	*(*int32)(unsafe.Add(mBase, uint32(l0))) = int32(1247)
	F_relation_close(m, v1680, int32(3))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L35
	} else {
		goto L532
	}
L527:
	;
	v1740 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+84))
	v1741 = int32(0)
	F_RenameRelationInternal(m, v1740, v1671, v1741, v1741)
	mBase = m.M
	v1744 = m.ExcPending
	if v1744 != 0 {
		goto L35
	} else {
		goto L530
	}
L528:
	;
	goto L529
L529:
	;
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+68))
	F_RenameTypeInternal(m, v1676, v1671, v1745)
	mBase = m.M
	v1747 = m.ExcPending
	if v1747 != 0 {
		goto L35
	} else {
		goto L531
	}
L530:
	;
	goto L526
L531:
	;
	goto L526
L532:
	;
	m.G0 = v1669 + int32(96)
	goto L497
L533:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669))) = v1676
	F_errmsg_internal(m, int32(_a_F_ExecRenameStmt_68), v1669)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L35
	} else {
		goto L534
	}
L534:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_65), int32(3828), int32(_a_F_ExecRenameStmt_66))
	mBase = m.M
	v1771 = m.ExcPending
	if v1771 != 0 {
		goto L35
	} else {
		goto L535
	}
L535:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L536:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1778 = m.ExcPending
	if v1778 != 0 {
		goto L35
	} else {
		goto L537
	}
L537:
	;
	v1779 = F_format_type_be(m, v1676)
	mBase = m.M
	v1780 = m.ExcPending
	if v1780 != 0 {
		goto L35
	} else {
		goto L538
	}
L538:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+80)) = v1779
	F_errmsg(m, int32(_a_F_ExecRenameStmt_69), v1669+int32(80))
	mBase = m.M
	v1786 = m.ExcPending
	if v1786 != 0 {
		goto L35
	} else {
		goto L539
	}
L539:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+64)) = int32(_a_F_ExecRenameStmt_70)
	F_errhint(m, int32(_a_F_ExecRenameStmt_71), v1669-int32(-64))
	mBase = m.M
	v1793 = m.ExcPending
	if v1793 != 0 {
		goto L35
	} else {
		goto L540
	}
L540:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_65), int32(3855), int32(_a_F_ExecRenameStmt_66))
	mBase = m.M
	v1798 = m.ExcPending
	if v1798 != 0 {
		goto L35
	} else {
		goto L541
	}
L541:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L542:
	;
	F_errcode(m, int32(151027844))
	mBase = m.M
	v1805 = m.ExcPending
	if v1805 != 0 {
		goto L35
	} else {
		goto L543
	}
L543:
	;
	v1806 = F_format_type_be(m, v1676)
	mBase = m.M
	v1807 = m.ExcPending
	if v1807 != 0 {
		goto L35
	} else {
		goto L544
	}
L544:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+32)) = v1806
	F_errmsg(m, int32(_a_F_ExecRenameStmt_72), v1669+int32(32))
	mBase = m.M
	v1813 = m.ExcPending
	if v1813 != 0 {
		goto L35
	} else {
		goto L545
	}
L545:
	;
	v1814 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+92))
	v1815 = F_format_type_be(m, v1814)
	mBase = m.M
	v1816 = m.ExcPending
	if v1816 != 0 {
		goto L35
	} else {
		goto L546
	}
L546:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1669)+16)) = v1815
	F_errhint(m, int32(_a_F_ExecRenameStmt_73), v1669+int32(16))
	mBase = m.M
	v1822 = m.ExcPending
	if v1822 != 0 {
		goto L35
	} else {
		goto L547
	}
L547:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_65), int32(3864), int32(_a_F_ExecRenameStmt_66))
	mBase = m.M
	v1827 = m.ExcPending
	if v1827 != 0 {
		goto L35
	} else {
		goto L548
	}
L548:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L549:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1836 = F_table_open(m, v1834, int32(3))
	mBase = m.M
	v1837 = m.ExcPending
	if v1837 != 0 {
		goto L35
	} else {
		goto L550
	}
L550:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v1839 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+56))
	v1841 = F_get_object_catcache_oid(m, v1840)
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L35
	} else {
		goto L551
	}
L551:
	;
	v1843 = F_get_object_catcache_name(m, v1840)
	mBase = m.M
	v1844 = m.ExcPending
	if v1844 != 0 {
		goto L35
	} else {
		goto L552
	}
L552:
	;
	v1845 = F_get_object_attnum_name(m, v1840)
	mBase = m.M
	v1846 = m.ExcPending
	if v1846 != 0 {
		goto L35
	} else {
		goto L553
	}
L553:
	;
	v1847 = F_get_object_attnum_namespace(m, v1840)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L35
	} else {
		goto L554
	}
L554:
	;
	v1849 = F_get_object_attnum_owner(m, v1840)
	mBase = m.M
	v1850 = m.ExcPending
	if v1850 != 0 {
		goto L35
	} else {
		goto L555
	}
L555:
	;
	v1852 = F_SearchSysCache1(m, v1841, base.I64_extend_i32_u(v1839))
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L35
	} else {
		goto L556
	}
L556:
	;
	if v1852 == int32(0) {
		goto L7
	} else {
		goto L557
	}
L557:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+52))
	v1858 = v17 + int32(159)
	v1859 = F_heap_getattr_2(m, v1852, v1845, v1856, v1858)
	mBase = m.M
	v1860 = m.ExcPending
	if v1860 != 0 {
		goto L35
	} else {
		goto L558
	}
L558:
	;
	if int32(0) < v1847 {
		goto L559
	} else {
		goto L560
	}
L559:
	;
	v1863 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+52))
	v1864 = F_heap_getattr_2(m, v1852, v1847, v1863, v1858)
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L35
	} else {
		goto L562
	}
L560:
	;
	v1867 = v3
	goto L561
L561:
	;
	v1868 = F_superuser(m)
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L35
	} else {
		goto L569
	}
L562:
	;
	v1867 = base.I32_wrap_i64(v1864)
	goto L561
L563:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+16))
	v2065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2064)+22)))
	v2066 = v2064 + v2065
	v2067 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2066)+104)))
	v2070 = *(*int32)(unsafe.Add(mBase, uint32(v2066)+68))
	F_IsThereFunctionInNamespace(m, v1838, v2067, v2066+int32(112), v2070)
	mBase = m.M
	v2072 = m.ExcPending
	if v2072 != 0 {
		goto L35
	} else {
		goto L646
	}
L564:
	;
	v2019 = int64(0)
	v2022 = F_SearchSysCacheExists(m, v1843, v1977, v2019, v2019, v2019)
	mBase = m.M
	v2023 = m.ExcPending
	if v2023 != 0 {
		goto L35
	} else {
		goto L627
	}
L565:
	;
	v2011 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[1])))
	v2013 = int64(0)
	v2015 = F_SearchSysCacheExists(m, int32(66), v2011, base.I64_extend_i32_u(v1838), v2013, v2013)
	mBase = m.M
	v2016 = m.ExcPending
	if v2016 != 0 {
		goto L35
	} else {
		goto L624
	}
L566:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+16))
	v2000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1999)+22)))
	v2001 = v1999 + v2000
	v2002 = *(*int32)(unsafe.Add(mBase, uint32(v2001)+4))
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v2001)+72))
	F_IsThereOpFamilyInNamespace(m, v1838, v2002, v2003)
	mBase = m.M
	v2005 = m.ExcPending
	if v2005 != 0 {
		goto L35
	} else {
		goto L623
	}
L567:
	;
	v1991 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+16))
	v1992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1991)+22)))
	v1994 = *(*int32)(unsafe.Add(mBase, uint32(v1991+v1992)+68))
	F_IsThereCollationInNamespace(m, v1838, v1994)
	mBase = m.M
	v1996 = m.ExcPending
	if v1996 != 0 {
		goto L35
	} else {
		goto L622
	}
L568:
	;
	if v1843 < int32(0) {
		goto L2
	} else {
		goto L617
	}
L569:
	;
	if v1868 == int32(0) {
		goto L570
	} else {
		goto L571
	}
L570:
	;
	if v1849 <= int32(0) {
		goto L6
	} else {
		goto L573
	}
L571:
	;
	goto L572
L572:
	;
	if v1840 <= int32(2752) {
		goto L609
	} else {
		goto L610
	}
L573:
	;
	v1874 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+52))
	v1877 = F_heap_getattr_2(m, v1852, v1849, v1874, v17+int32(159))
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L35
	} else {
		goto L574
	}
L574:
	;
	v1880 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v1882 = F_has_privs_of_role(m, v1880, base.I32_wrap_i64(v1877))
	mBase = m.M
	v1883 = m.ExcPending
	if v1883 != 0 {
		goto L35
	} else {
		goto L575
	}
L575:
	;
	if v1882 == int32(0) {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	v1887 = F_get_object_type(m, v1840, v1839)
	mBase = m.M
	v1888 = m.ExcPending
	if v1888 != 0 {
		goto L35
	} else {
		goto L579
	}
L577:
	;
	goto L578
L578:
	;
	if v1867 == int32(0) {
		goto L581
	} else {
		goto L582
	}
L579:
	;
	F_aclcheck_error(m, int32(2), v1887, base.I32_wrap_i64(v1859))
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L35
	} else {
		goto L580
	}
L580:
	;
	goto L578
L581:
	;
	if v1840 <= int32(2752) {
		goto L587
	} else {
		goto L588
	}
L582:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v1898 = F_object_aclcheck(m, int32(2615), v1867, v1896, int64(512))
	mBase = m.M
	v1899 = m.ExcPending
	if v1899 != 0 {
		goto L35
	} else {
		goto L583
	}
L583:
	;
	if v1898 == int32(0) {
		goto L581
	} else {
		goto L584
	}
L584:
	;
	v1903 = F_get_namespace_name(m, v1867)
	mBase = m.M
	v1904 = m.ExcPending
	if v1904 != 0 {
		goto L35
	} else {
		goto L585
	}
L585:
	;
	F_aclcheck_error(m, v1898, int32(37), v1903)
	mBase = m.M
	v1906 = m.ExcPending
	if v1906 != 0 {
		goto L35
	} else {
		goto L586
	}
L586:
	;
	goto L581
L587:
	;
	if v1840 == int32(1255) {
		goto L563
	} else {
		goto L590
	}
L588:
	;
	goto L589
L589:
	;
	if v1840 == int32(2753) {
		goto L566
	} else {
		goto L592
	}
L590:
	;
	if v1840 != int32(2616) {
		goto L568
	} else {
		goto L591
	}
L591:
	;
	goto L3
L592:
	;
	if v1840 == int32(3456) {
		goto L567
	} else {
		goto L593
	}
L593:
	;
	if v1840 != int32(_a_F_ExecRenameStmt_74) {
		goto L568
	} else {
		goto L594
	}
L594:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[1]))
	v1924 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[0]))
	v1926 = F_object_aclcheck(m, int32(1262), v1922, v1924, int64(512))
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L35
	} else {
		goto L595
	}
L595:
	;
	if v1926 != 0 {
		goto L596
	} else {
		goto L597
	}
L596:
	;
	v1930 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[1]))
	v1931 = F_get_database_name(m, v1930)
	mBase = m.M
	v1932 = m.ExcPending
	if v1932 != 0 {
		goto L35
	} else {
		goto L599
	}
L597:
	;
	goto L598
L598:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+16))
	v1936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1935)+22)))
	v1938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1935+v1936)+89)))
	if v1938 != 0 {
		goto L565
	} else {
		goto L601
	}
L599:
	;
	F_aclcheck_error(m, v1926, int32(9), v1931)
	mBase = m.M
	v1934 = m.ExcPending
	if v1934 != 0 {
		goto L35
	} else {
		goto L600
	}
L600:
	;
	goto L598
L601:
	;
	v1939 = F_superuser(m)
	mBase = m.M
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L35
	} else {
		goto L602
	}
L602:
	;
	if v1939 != 0 {
		goto L565
	} else {
		goto L603
	}
L603:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L35
	} else {
		goto L604
	}
L604:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L35
	} else {
		goto L605
	}
L605:
	;
	F_errmsg(m, int32(_a_F_ExecRenameStmt_75), int32(0))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L35
	} else {
		goto L606
	}
L606:
	;
	F_errhint(m, int32(_a_F_ExecRenameStmt_76), int32(0))
	mBase = m.M
	v1955 = m.ExcPending
	if v1955 != 0 {
		goto L35
	} else {
		goto L607
	}
L607:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_77), int32(251), int32(_a_F_ExecRenameStmt_78))
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L35
	} else {
		goto L608
	}
L608:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L609:
	;
	if v1840 == int32(1255) {
		goto L563
	} else {
		goto L612
	}
L610:
	;
	goto L611
L611:
	;
	if v1840 == int32(2753) {
		goto L566
	} else {
		goto L614
	}
L612:
	;
	if v1840 == int32(2616) {
		goto L3
	} else {
		goto L613
	}
L613:
	;
	goto L568
L614:
	;
	if v1840 == int32(3456) {
		goto L567
	} else {
		goto L615
	}
L615:
	;
	if v1840 == int32(_a_F_ExecRenameStmt_74) {
		goto L565
	} else {
		goto L616
	}
L616:
	;
	goto L568
L617:
	;
	v1977 = base.I64_extend_i32_u(v1838)
	if v1867 == int32(0) {
		goto L564
	} else {
		goto L618
	}
L618:
	;
	v1981 = int64(0)
	v1983 = F_SearchSysCacheExists(m, v1843, v1977, base.I64_extend_i32_u(v1867), v1981, v1981)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L35
	} else {
		goto L619
	}
L619:
	;
	if v1983 == int32(0) {
		goto L2
	} else {
		goto L620
	}
L620:
	;
	F_report_namespace_conflict(m, v1840, v1838, v1867)
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L35
	} else {
		goto L621
	}
L621:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L622:
	;
	goto L2
L623:
	;
	goto L2
L624:
	;
	if v2015 != 0 {
		goto L5
	} else {
		goto L625
	}
L625:
	;
	F_LogicalRepWorkersWakeupAtCommit(m, v1839)
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L35
	} else {
		goto L626
	}
L626:
	;
	goto L2
L627:
	;
	if v2022 == int32(0) {
		goto L2
	} else {
		goto L628
	}
L628:
	;
	if v1840 <= int32(3465) {
		goto L635
	} else {
		goto L636
	}
L629:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2048 = m.ExcPending
	if v2048 != 0 {
		goto L35
	} else {
		goto L642
	}
L630:
	;
	if v1840 != int32(3466) {
		goto L4
	} else {
		goto L641
	}
L631:
	;
	v2044 = int32(_a_F_ExecRenameStmt_79)
	goto L629
L632:
	;
	v2044 = int32(_a_F_ExecRenameStmt_80)
	goto L629
L633:
	;
	v2044 = int32(_a_F_ExecRenameStmt_81)
	goto L629
L634:
	;
	v2044 = int32(_a_F_ExecRenameStmt_82)
	goto L629
L635:
	;
	if v1840 == int32(1417) {
		goto L633
	} else {
		goto L638
	}
L636:
	;
	goto L637
L637:
	;
	switch v1840 - int32(_a_F_ExecRenameStmt_74) {
	case 0:
		goto L631
	case 1, 2, 3:
		goto L4
	case 4:
		goto L632
	default:
		goto L630
	}
L638:
	;
	if v1840 == int32(2328) {
		goto L634
	} else {
		goto L639
	}
L639:
	;
	if v1840 != int32(2612) {
		goto L4
	} else {
		goto L640
	}
L640:
	;
	v2044 = int32(_a_F_ExecRenameStmt_83)
	goto L629
L641:
	;
	v2044 = int32(_a_F_ExecRenameStmt_84)
	goto L629
L642:
	;
	F_errcode(m, int32(_a_F_ExecRenameStmt_25))
	mBase = m.M
	v2051 = m.ExcPending
	if v2051 != 0 {
		goto L35
	} else {
		goto L643
	}
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v1838
	F_errmsg(m, v2044, v17+int32(48))
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L35
	} else {
		goto L644
	}
L644:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_77), int32(107), int32(_a_F_ExecRenameStmt_85))
	mBase = m.M
	v2061 = m.ExcPending
	if v2061 != 0 {
		goto L35
	} else {
		goto L645
	}
L645:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L646:
	;
	goto L2
L647:
	;
	v2077 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v2077
	F_errmsg_internal(m, int32(_a_F_ExecRenameStmt_86), v17)
	mBase = m.M
	v2081 = m.ExcPending
	if v2081 != 0 {
		goto L35
	} else {
		goto L648
	}
L648:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_77), int32(452), int32(_a_F_ExecRenameStmt_87))
	mBase = m.M
	v2086 = m.ExcPending
	if v2086 != 0 {
		goto L35
	} else {
		goto L649
	}
L649:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L650:
	;
	m.G0 = v2089 + int32(32)
	goto L1
L651:
	;
	v2163 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v2164 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v2165 != 0 {
		goto L675
	} else {
		goto L676
	}
L652:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2150 = m.ExcPending
	if v2150 != 0 {
		goto L35
	} else {
		goto L672
	}
L653:
	;
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v2096 = F_makeTypeNameFromNameList(m, v2095)
	mBase = m.M
	v2097 = m.ExcPending
	if v2097 != 0 {
		goto L35
	} else {
		goto L656
	}
L654:
	;
	goto L655
L655:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+32)))
	v2122 = F_RangeVarGetRelidExtended(m, v2117, int32(8), v2119, int32(620), int32(0))
	mBase = m.M
	v2123 = m.ExcPending
	if v2123 != 0 {
		goto L35
	} else {
		goto L664
	}
L656:
	;
	v2098 = F_typenameTypeId(m, int32(0), v2096)
	mBase = m.M
	v2099 = m.ExcPending
	if v2099 != 0 {
		goto L35
	} else {
		goto L657
	}
L657:
	;
	v2102 = F_table_open(m, int32(1247), int32(3))
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L35
	} else {
		goto L658
	}
L658:
	;
	v2106 = F_SearchSysCache1(m, int32(82), base.I64_extend_i32_u(v2098))
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L35
	} else {
		goto L659
	}
L659:
	;
	if v2106 == int32(0) {
		goto L652
	} else {
		goto L660
	}
L660:
	;
	F_checkDomainOwner(m, v2106)
	mBase = m.M
	v2111 = m.ExcPending
	if v2111 != 0 {
		goto L35
	} else {
		goto L661
	}
L661:
	;
	F_ReleaseCatCache(m, v2106)
	mBase = m.M
	v2113 = m.ExcPending
	if v2113 != 0 {
		goto L35
	} else {
		goto L662
	}
L662:
	;
	F_relation_close(m, v2102, int32(0))
	mBase = m.M
	v2116 = m.ExcPending
	if v2116 != 0 {
		goto L35
	} else {
		goto L663
	}
L663:
	;
	v2160 = v2098
	v2161 = v3
	goto L651
L664:
	;
	if v2122 != 0 {
		v2160 = v3
		v2161 = v2122
		goto L651
	} else {
		goto L665
	}
L665:
	;
	v2126 = F_errstart(m, int32(18), int32(0))
	mBase = m.M
	v2127 = m.ExcPending
	if v2127 != 0 {
		goto L35
	} else {
		goto L666
	}
L666:
	;
	if v2126 != 0 {
		goto L667
	} else {
		goto L668
	}
L667:
	;
	v2128 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v2129 = *(*int32)(unsafe.Add(mBase, uint32(v2128)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v2089)+16)) = v2129
	F_errmsg(m, int32(_a_F_ExecRenameStmt_43), v2089+int32(16))
	mBase = m.M
	v2135 = m.ExcPending
	if v2135 != 0 {
		goto L35
	} else {
		goto L670
	}
L668:
	;
	goto L669
L669:
	;
	v2142 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[6]))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v2142
	v2145 = *(*int64)(unsafe.Add(mBase, _c_F_ExecRenameStmt[7]))
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v2145
	goto L650
L670:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_44), int32(_a_F_ExecRenameStmt_88), int32(_a_F_ExecRenameStmt_89))
	mBase = m.M
	v2140 = m.ExcPending
	if v2140 != 0 {
		goto L35
	} else {
		goto L671
	}
L671:
	;
	goto L669
L672:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2089))) = v2098
	F_errmsg_internal(m, int32(_a_F_ExecRenameStmt_68), v2089)
	mBase = m.M
	v2154 = m.ExcPending
	if v2154 != 0 {
		goto L35
	} else {
		goto L673
	}
L673:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_44), int32(_a_F_ExecRenameStmt_90), int32(_a_F_ExecRenameStmt_89))
	mBase = m.M
	v2159 = m.ExcPending
	if v2159 != 0 {
		goto L35
	} else {
		goto L674
	}
L674:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L675:
	;
	v2166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2165)+16)))
	v2168 = v2166
	goto L677
L676:
	;
	v2168 = int32(0)
	goto L677
L677:
	;
	F_rename_constraint_internal(m, l0, v2161, v2160, v2163, v2164, v2168&int32(1), int32(0))
	mBase = m.M
	v2173 = m.ExcPending
	if v2173 != 0 {
		goto L35
	} else {
		goto L678
	}
L678:
	;
	goto L650
L679:
	;
	v2185 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v1839
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v2185 + int32(4)
	F_errmsg_internal(m, int32(_a_F_ExecRenameStmt_91), v17+int32(16))
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L35
	} else {
		goto L680
	}
L680:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_77), int32(183), int32(_a_F_ExecRenameStmt_78))
	mBase = m.M
	v2199 = m.ExcPending
	if v2199 != 0 {
		goto L35
	} else {
		goto L681
	}
L681:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L682:
	;
	F_errcode(m, int32(16797828))
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L35
	} else {
		goto L683
	}
L683:
	;
	v2207 = F_getObjectDescriptionOids(m, v1840, v1839)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L35
	} else {
		goto L684
	}
L684:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v2207
	F_errmsg(m, int32(_a_F_ExecRenameStmt_92), v17+int32(80))
	mBase = m.M
	v2214 = m.ExcPending
	if v2214 != 0 {
		goto L35
	} else {
		goto L685
	}
L685:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_77), int32(209), int32(_a_F_ExecRenameStmt_78))
	mBase = m.M
	v2219 = m.ExcPending
	if v2219 != 0 {
		goto L35
	} else {
		goto L686
	}
L686:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L687:
	;
	F_errcode(m, int32(_a_F_ExecRenameStmt_25))
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L35
	} else {
		goto L688
	}
L688:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v1838
	F_errmsg(m, int32(_a_F_ExecRenameStmt_79), v17-int32(-64))
	mBase = m.M
	v2232 = m.ExcPending
	if v2232 != 0 {
		goto L35
	} else {
		goto L689
	}
L689:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_77), int32(107), int32(_a_F_ExecRenameStmt_85))
	mBase = m.M
	v2237 = m.ExcPending
	if v2237 != 0 {
		goto L35
	} else {
		goto L690
	}
L690:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L691:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v1840
	F_errmsg_internal(m, int32(_a_F_ExecRenameStmt_93), v17+int32(32))
	mBase = m.M
	v2248 = m.ExcPending
	if v2248 != 0 {
		goto L35
	} else {
		goto L692
	}
L692:
	;
	F_errfinish(m, int32(_a_F_ExecRenameStmt_77), int32(101), int32(_a_F_ExecRenameStmt_85))
	mBase = m.M
	v2253 = m.ExcPending
	if v2253 != 0 {
		goto L35
	} else {
		goto L693
	}
L693:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L694:
	;
	goto L2
L695:
	;
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+48))
	v2274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2273)+120)))
	v2275 = F_palloc0(m, v2274)
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L35
	} else {
		goto L696
	}
L696:
	;
	v2277 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+48))
	v2278 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2277)+120)))
	v2279 = F_palloc0(m, v2278)
	mBase = m.M
	v2280 = m.ExcPending
	if v2280 != 0 {
		goto L35
	} else {
		goto L697
	}
L697:
	;
	v2282 = v17 + int32(95)
	v2284 = F_strncpy(m, v2282, v1838, int32(64))
	mBase = m.M
	v2285 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2284)+63)) = uint8(v2285)
	goto L698
L698:
	;
	v2287 = int32(1)
	v2288 = v1845 - v2287
	*(*int64)(unsafe.Add(mBase, uint32(v2271+v2288<<(uint(int32(3))%32)))) = base.I64_extend_i32_u(v2282)
	*(*uint8)(unsafe.Add(mBase, uint32(v2288+v2279))) = uint8(v2287)
	v2299 = *(*int32)(unsafe.Add(mBase, uint32(v1836)+52))
	v2300 = F_heap_modify_tuple(m, v1852, v2299, v2271, v2275, v2279)
	mBase = m.M
	v2301 = m.ExcPending
	if v2301 != 0 {
		goto L35
	} else {
		goto L699
	}
L699:
	;
	F_CatalogTupleUpdate(m, v1836, v1852+int32(4), v2300)
	mBase = m.M
	v2303 = m.ExcPending
	if v2303 != 0 {
		goto L35
	} else {
		goto L700
	}
L700:
	;
	v2305 = *(*int32)(unsafe.Add(mBase, _c_F_ExecRenameStmt[2]))
	if v2305 != 0 {
		goto L701
	} else {
		goto L702
	}
L701:
	;
	v2306 = int32(0)
	F_RunObjectPostAlterHook(m, v1840, v1839, v2306, v2306, v2306)
	mBase = m.M
	v2310 = m.ExcPending
	if v2310 != 0 {
		goto L35
	} else {
		goto L704
	}
L702:
	;
	goto L703
L703:
	;
	if v1840 == int32(_a_F_ExecRenameStmt_94) {
		goto L705
	} else {
		goto L706
	}
L704:
	;
	goto L703
L705:
	;
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v1852)+16))
	v2314 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2313)+22)))
	v2315 = v2313 + v2314
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(v2315)))
	v2317 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2315)+72)))
	if v2317 != 0 {
		goto L709
	} else {
		goto L710
	}
L706:
	;
	goto L707
L707:
	;
	F_pfree(m, v2271)
	mBase = m.M
	v2402 = m.ExcPending
	if v2402 != 0 {
		goto L35
	} else {
		goto L723
	}
L708:
	;
	goto L707
L709:
	;
	F_CacheInvalidateRelSync(m, int32(0))
	mBase = m.M
	v2320 = m.ExcPending
	if v2320 != 0 {
		goto L35
	} else {
		goto L712
	}
L710:
	;
	goto L711
L711:
	;
	v2322 = F_GetIncludedPublicationRelations(m, v2316, int32(2))
	mBase = m.M
	v2323 = m.ExcPending
	if v2323 != 0 {
		goto L35
	} else {
		goto L714
	}
L712:
	;
	goto L708
L713:
	;
	goto L708
L714:
	;
	v2324 = F_GetAllSchemaPublicationRelations(m, v2316)
	mBase = m.M
	v2325 = m.ExcPending
	if v2325 != 0 {
		goto L35
	} else {
		goto L715
	}
L715:
	;
	v2326 = F_list_concat_unique_oid(m, v2322, v2324)
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L35
	} else {
		goto L716
	}
L716:
	;
	if v2326 == int32(0) {
		goto L713
	} else {
		goto L717
	}
L717:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v2326)+4))
	if v2330 <= int32(0) {
		goto L713
	} else {
		goto L718
	}
L718:
	;
	v2337 = int32(0)
	goto L719
L719:
	;
	v2348 = *(*int32)(unsafe.Add(mBase, uint32(v2326)+12))
	v2352 = *(*int32)(unsafe.Add(mBase, uint32(v2348+v2337<<(uint(int32(2))%32))))
	F_CacheInvalidateRelSync(m, v2352)
	mBase = m.M
	v2354 = m.ExcPending
	if v2354 != 0 {
		goto L35
	} else {
		goto L721
	}
L720:
	;
	goto L713
L721:
	;
	v2356 = v2337 + int32(1)
	v2357 = *(*int32)(unsafe.Add(mBase, uint32(v2326)+4))
	if v2356 < v2357 {
		v2337 = v2356
		goto L719
	} else {
		goto L722
	}
L722:
	;
	goto L720
L723:
	;
	F_pfree(m, v2275)
	mBase = m.M
	v2404 = m.ExcPending
	if v2404 != 0 {
		goto L35
	} else {
		goto L724
	}
L724:
	;
	F_pfree(m, v2279)
	mBase = m.M
	v2406 = m.ExcPending
	if v2406 != 0 {
		goto L35
	} else {
		goto L725
	}
L725:
	;
	F_pfree(m, v2300)
	mBase = m.M
	v2408 = m.ExcPending
	if v2408 != 0 {
		goto L35
	} else {
		goto L726
	}
L726:
	;
	F_ReleaseCatCache(m, v1852)
	mBase = m.M
	v2410 = m.ExcPending
	if v2410 != 0 {
		goto L35
	} else {
		goto L727
	}
L727:
	;
	F_relation_close(m, v1836, int32(3))
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L35
	} else {
		goto L728
	}
L728:
	;
	goto L1
}
func F_ExecStoreMinimalTuple(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v6 == int32(_a_F_ExecStoreMinimalTuple_0) {
		v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		if v9&int32(4) != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
			F_pfree(m, v12)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
				v20 = v17 & int32(-5)
				v21 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v21)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v21
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v21)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = l0
				v31 = v20 & int32(_a_F_ExecStoreMinimalTuple_1)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v31)
				v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v34 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l0 - v34
				*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v33 + v34
				if l2 != 0 {
					v41 = v31 | int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v41)
				} else {
				}
				return l1
			}
		} else {
			v20 = v9
			v21 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+36)) = uint16(v21)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+32)) = int32(-1)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+72)) = v21
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)) = uint16(v21)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = l0
			v31 = v20 & int32(_a_F_ExecStoreMinimalTuple_1)
			*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v31)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v34 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(l1)+68)) = l0 - v34
			*(*int32)(unsafe.Add(mBase, uint32(l1)+52)) = v33 + v34
			if l2 != 0 {
				v41 = v31 | int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)) = uint16(v41)
			} else {
			}
			return l1
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_ExecStoreMinimalTuple_2), int32(0))
			mBase = m.M
			v51 = m.ExcPending
			if v51 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_ExecStoreMinimalTuple_3), int32(1740), int32(_a_F_ExecStoreMinimalTuple_4))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
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
func F_ExecStoreVirtualTuple(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	v2 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v4 = v2 & int32(_a_F_ExecStoreVirtualTuple_0)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)) = uint16(v4)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+6)) = uint16(v7)
	return l0
}
func F_ExecSupportsBackwardScan(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	v2 = int32(0)
	if l0 == v2 {
		v60 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v60
L2:
	;
	v6 = l0
	goto L3
L3:
	;
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+36)))
	if v9 != 0 {
		v60 = v2
		goto L1
	} else {
		goto L5
	}
L4:
	;
	v60 = int32(1)
	goto L1
L5:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
	switch v10 - int32(335) {
	case 0:
		goto L14
	default:
		v60 = v2
		goto L1
	case 3:
		goto L13
	case 8, 14, 15, 17, 18, 20, 29, 31:
		goto L6
	case 10:
		goto L12
	case 11:
		goto L11
	case 16:
		goto L10
	case 24:
		goto L8
	case 41, 42:
		goto L7
	}
L6:
	;
	goto L4
L7:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
	if v57 != 0 {
		v6 = v57
		goto L3
	} else {
		goto L28
	}
L8:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+80)))
	return v53 & int32(1)
L9:
	;
	if v52 != 0 {
		v6 = v52
		goto L3
	} else {
		goto L27
	}
L10:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	v52 = v51
	goto L9
L11:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	v48 = F_IndexSupportsBackwardScan(m, v47)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L21
	} else {
		goto L26
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	v44 = F_IndexSupportsBackwardScan(m, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L21
	} else {
		goto L25
	}
L13:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
	if int32(0) < v14 {
		v60 = v2
		goto L1
	} else {
		goto L16
	}
L14:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
	if v13 != 0 {
		v52 = v13
		goto L9
	} else {
		goto L15
	}
L15:
	;
	v60 = v2
	goto L1
L16:
	;
	v17 = int32(1)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
	if v18 == int32(0) {
		v60 = v17
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v21 = int32(0)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v22 <= v21 {
		v60 = v17
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v25 = v21
	goto L19
L19:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v28+v25<<(uint(int32(2))%32))))
	v33 = F_ExecSupportsBackwardScan(m, v32)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v60 = v33
	goto L1
L21:
	;
	return int32(0)
L22:
	;
	if v33 == int32(0) {
		v60 = v33
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v40 = v25 + int32(1)
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v40 < v41 {
		v25 = v40
		goto L19
	} else {
		goto L24
	}
L24:
	;
	goto L20
L25:
	;
	return v44
L26:
	;
	return v48
L27:
	;
	v60 = v2
	goto L1
L28:
	;
	v60 = v2
	goto L1
}
func F_ExtendBufferedRelTo(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int64
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
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
	var v117 int64
	_ = v117
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
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int64
	_ = v146
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
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
	var v171 int32
	_ = v171
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int64
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int64
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v277 int32
	_ = v277
	var v279 int64
	_ = v279
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v340 int32
	_ = v340
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int64
	_ = v375
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v402 int32
	_ = v402
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v449 int32
	_ = v449
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v467 int32
	_ = v467
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int64
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int64
	_ = v499
	var v501 int64
	_ = v501
	var v504 int32
	_ = v504
	var v505 int64
	_ = v505
	var v512 int64
	_ = v512
	var v515 int64
	_ = v515
	var v520 int64
	_ = v520
	var v537 int64
	_ = v537
	var v544 int64
	_ = v544
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v577 int64
	_ = v577
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v624 int32
	_ = v624
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v640 int32
	_ = v640
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v652 int32
	_ = v652
	var v654 int32
	_ = v654
	var v657 int32
	_ = v657
	var v659 int64
	_ = v659
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v681 int64
	_ = v681
	var v685 int64
	_ = v685
	var v687 int32
	_ = v687
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int64
	_ = v719
	var v725 int32
	_ = v725
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v751 int64
	_ = v751
	var v757 int32
	_ = v757
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v785 int64
	_ = v785
	var v787 int32
	_ = v787
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v805 int32
	_ = v805
	v6 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(528)
	m.G0 = v21
	*(*int32)(unsafe.Add(mBase, uint32(v21)+404)) = v6
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	if v26 == v6 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+48))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29)+118)))
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)) = uint8(v30)
	goto L3
L2:
	;
	goto L3
L3:
	;
	if l2&int32(4) == int32(0) {
		v182 = v25
		goto L4
	} else {
		goto L5
	}
L4:
	;
	if l2&int32(16) != 0 {
		goto L55
	} else {
		goto L56
	}
L5:
	;
	if v25 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v107 = F_smgrexists(m, v106, l1)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L13
	} else {
		goto L27
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v98+l1<<(uint(int32(2))%32))+20))
	switch v102 + int32(1) {
	case 0:
		goto L26
	case 1:
		v106 = v98
		goto L6
	default:
		v182 = v25
		goto L4
	}
L10:
	;
	v64 = v36
	goto L12
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+136)) = v38
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+128)) = v40
	v44 = F_smgropen(m, v21+int32(128), v37)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v64+l1<<(uint(int32(2))%32))+20))
	if base.Ui32(v68-int32(1)) < base.Ui32(int32(-2)) {
		v182 = v25
		goto L4
	} else {
		goto L19
	}
L13:
	;
	return int32(0)
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v44
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	if v50 != 0 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v64 = v62
	goto L12
L16:
	;
	v58 = v50
	goto L18
L17:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v44)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v51)+4)) = v52
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v44)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v52))) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v44)+72))
	v58 = v56
	goto L18
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v44)+72)) = v58 + int32(1)
	goto L15
L19:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	if v73 != 0 {
		v106 = v73
		goto L6
	} else {
		goto L20
	}
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v25)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+120)) = v75
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v25)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+112)) = v77
	v81 = F_smgropen(m, v21+int32(112), v74)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L13
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+12)) = v81
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v81)+72))
	if v85 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v25)+12))
	v106 = v97
	goto L6
L23:
	;
	v93 = v85
	goto L25
L24:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v81)+76))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v81)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v86)+4)) = v87
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v81)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v87))) = v89
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v81)+72))
	v93 = v91
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v81)+72)) = v93 + int32(1)
	goto L22
L26:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v106 = v105
	goto L6
L27:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v107 != 0 {
		v182 = v109
		goto L4
	} else {
		goto L28
	}
L28:
	;
	F_LockRelationForExtension(m, v109, int32(7))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L13
	} else {
		goto L29
	}
L29:
	;
	if v109 != 0 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	F_UnlockRelationForExtension(m, v109, int32(7))
	mBase = m.M
	v181 = m.ExcPending
	if v181 != 0 {
		goto L13
	} else {
		goto L54
	}
L31:
	;
	F_smgrcreate(m, v171, l1, int32(base.Ui32(l2&int32(2))>>(uint(int32(1))%32)))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L13
	} else {
		goto L53
	}
L32:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	if v113 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v168 = F_smgrexists(m, v167, l1)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L13
	} else {
		goto L51
	}
L35:
	;
	v139 = v113
	goto L37
L36:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+104)) = v115
	v117 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+96)) = v117
	v121 = F_smgropen(m, v21+int32(96), v114)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L13
	} else {
		goto L38
	}
L37:
	;
	v140 = F_smgrexists(m, v139, l1)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L13
	} else {
		goto L43
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v121
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v121)+72))
	if v125 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v139 = v137
	goto L37
L40:
	;
	v133 = v125
	goto L42
L41:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(v121)+76))
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v121)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v127
	v129 = *(*int32)(unsafe.Add(mBase, uint32(v121)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v127))) = v129
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v121)+72))
	v133 = v131
	goto L42
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+72)) = v133 + int32(1)
	goto L39
L43:
	;
	if v140 != 0 {
		goto L30
	} else {
		goto L44
	}
L44:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	if v142 != 0 {
		v171 = v142
		goto L31
	} else {
		goto L45
	}
L45:
	;
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v109)+20))
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v109)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+88)) = v144
	v146 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+80)) = v146
	v150 = F_smgropen(m, v21+int32(80), v143)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L13
	} else {
		goto L46
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v109)+12)) = v150
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v150)+72))
	if v154 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v109)+12))
	v171 = v166
	goto L31
L48:
	;
	v162 = v154
	goto L50
L49:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v155)+4)) = v156
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v156))) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v150)+72))
	v162 = v160
	goto L50
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+72)) = v162 + int32(1)
	goto L47
L51:
	;
	if v168 != 0 {
		goto L30
	} else {
		goto L52
	}
L52:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v171 = v170
	goto L31
L53:
	;
	goto L30
L54:
	;
	v182 = v109
	goto L4
L55:
	;
	if v182 != 0 {
		goto L59
	} else {
		goto L60
	}
L56:
	;
	goto L57
L57:
	;
	if v182 != 0 {
		goto L69
	} else {
		goto L70
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v212+l1<<(uint(int32(2))%32))+20)) = int32(-1)
	goto L57
L59:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	if v186 != 0 {
		v212 = v186
		goto L58
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v212 = v211
	goto L58
L62:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+72)) = v188
	v190 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+64)) = v190
	v194 = F_smgropen(m, v21-int32(-64), v187)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L13
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = v194
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v194)+72))
	if v198 != 0 {
		goto L65
	} else {
		goto L66
	}
L64:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v212 = v210
	goto L58
L65:
	;
	v206 = v198
	goto L67
L66:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v194)+76))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v194)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v199)+4)) = v200
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v194)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v200))) = v202
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v194)+72))
	v206 = v204
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v194)+72)) = v206 + int32(1)
	goto L64
L68:
	;
	v247 = l4 - int32(1)
	v248 = F_smgrnblocks(m, v245, l1)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L13
	} else {
		goto L79
	}
L69:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	if v219 != 0 {
		v245 = v219
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v245 = v244
	goto L68
L72:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v221
	v223 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+48)) = v223
	v227 = F_smgropen(m, v21+int32(48), v220)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L13
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = v227
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
	if v231 != 0 {
		goto L75
	} else {
		goto L76
	}
L74:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v245 = v243
	goto L68
L75:
	;
	v239 = v231
	goto L77
L76:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v227)+76))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v227)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v232)+4)) = v233
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v227)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v233))) = v235
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v227)+72))
	v239 = v237
	goto L77
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v227)+72)) = v239 + int32(1)
	goto L74
L78:
	;
	m.G0 = v21 + int32(528)
	return v805
L79:
	;
	if base.Ui32(v248) < base.Ui32(l3) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if base.Ui32(v247) < base.Ui32(int32(2)) {
		goto L83
	} else {
		goto L84
	}
L81:
	;
	goto L82
L82:
	;
	if v182 != 0 {
		goto L112
	} else {
		goto L113
	}
L83:
	;
	v255 = l2 | int32(32)
	goto L85
L84:
	;
	v255 = l2
	goto L85
L85:
	;
	v267 = v6
	v268 = v248
	goto L86
L86:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v277
	v279 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+32)) = v279
	if base.Ui64(base.I64_extend_i32_u(v268)-int64(-64)) <= base.Ui64(base.I64_extend_i32_u(l3)) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	if v340 != 0 {
		v805 = v340
		goto L78
	} else {
		goto L104
	}
L88:
	;
	v290 = int32(64)
	goto L90
L89:
	;
	v290 = l3 - v268
	goto L90
L90:
	;
	v295 = F_ExtendBufferedRelCommon(m, v21+int32(32), l1, int32(0), v255, v290, l3, v21+int32(144), v21+int32(404))
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L13
	} else {
		goto L91
	}
L91:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v21)+404))
	v298 = v295 + v297
	if v297 != 0 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	v302 = int32(0)
	v308 = v267
	goto L95
L93:
	;
	v340 = v267
	goto L94
L94:
	;
	if base.Ui32(v298) < base.Ui32(l3) {
		v267 = v340
		v268 = v298
		goto L86
	} else {
		goto L103
	}
L95:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v21+int32(144)+v302<<(uint(int32(2))%32))))
	if l3-int32(1) == v302+v295 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	v340 = v328
	goto L94
L97:
	;
	v330 = v302 + int32(1)
	if v330 != v297 {
		v302 = v330
		v308 = v328
		goto L95
	} else {
		goto L102
	}
L98:
	;
	v328 = v323
	goto L97
L99:
	;
	goto L100
L100:
	;
	F_ReleaseBuffer(m, v323)
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L13
	} else {
		goto L101
	}
L101:
	;
	v328 = v308
	goto L97
L102:
	;
	goto L96
L103:
	;
	goto L87
L104:
	;
	goto L82
L105:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+416)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+412)) = v182
	v785 = *(*int64)(unsafe.Add(mBase, uint32(v21)+412))
	*(*int64)(unsafe.Add(mBase, uint32(v21))) = v785
	v787 = *(*int32)(unsafe.Add(mBase, uint32(v21)+420))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = v787
	if base.Ui32(v247) < base.Ui32(int32(2)) {
		goto L203
	} else {
		goto L204
	}
L106:
	;
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v757)+20))
	v775 = v773 + int32(1)
	v776 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+411)))
	F_ZeroAndLockBuffer(m, v775, l4, v776)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L13
	} else {
		goto L202
	}
L107:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v182)+272))
	if v741 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L108:
	;
	v677 = v670*int32(320) + v667<<(uint(int32(6))%32)
	v681 = *(*int64)(unsafe.Add(mBase, uint32(v677)+uint32(_c_F_ExtendBufferedRelTo[0])))
	*(*int64)(unsafe.Add(mBase, uint32(v677)+uint32(_c_F_ExtendBufferedRelTo[0]))) = v681 + int64(1)
	v685 = *(*int64)(unsafe.Add(mBase, uint32(v677)+uint32(_c_F_ExtendBufferedRelTo[1])))
	*(*int64)(unsafe.Add(mBase, uint32(v677)+uint32(_c_F_ExtendBufferedRelTo[1]))) = v685
	v687 = int32(1)
	F_pgstat_count_backend_io_op(m, v670, v667, int32(2), v687, int64(0))
	mBase = m.M
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[2])) = uint8(v687)
	*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[3])) = uint8(v687)
	goto L187
L109:
	;
	v657 = int32(_a_F_ExtendBufferedRelTo_0)
	v659 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[4]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[4])) = v659 + int64(1)
	v666 = v654
	v667 = v421
	v670 = int32(0)
	goto L108
L110:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L13
	} else {
		goto L183
	}
L111:
	;
	v412 = int32(1)
	v413 = l3 - v412
	if base.Ui32(v247) <= base.Ui32(v412) {
		goto L129
	} else {
		goto L130
	}
L112:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	if v369 == int32(0) {
		goto L115
	} else {
		goto L116
	}
L113:
	;
	goto L114
L114:
	;
	if l3 == int32(0) {
		goto L105
	} else {
		goto L128
	}
L115:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v182)+20))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v182)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v373
	v375 = *(*int64)(unsafe.Add(mBase, uint32(v182)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v375
	v379 = F_smgropen(m, v21+int32(16), v372)
	mBase = m.M
	v380 = m.ExcPending
	if v380 != 0 {
		goto L13
	} else {
		goto L118
	}
L116:
	;
	v397 = v369
	goto L117
L117:
	;
	v398 = *(*int32)(unsafe.Add(mBase, uint32(v182)+48))
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398)+118)))
	if v399 == int32(116) {
		goto L123
	} else {
		goto L124
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v182)+12)) = v379
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v379)+72))
	if v383 != 0 {
		goto L120
	} else {
		goto L121
	}
L119:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v182)+12))
	v397 = v395
	goto L117
L120:
	;
	v391 = v383
	goto L122
L121:
	;
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v379)+76))
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v379)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v384)+4)) = v385
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v379)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = v387
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v379)+72))
	v391 = v389
	goto L122
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v379)+72)) = v391 + int32(1)
	goto L119
L123:
	;
	v402 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+24)))
	if v402 == int32(0) {
		goto L110
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	if l3 != 0 {
		v410 = v397
		v411 = v399
		goto L111
	} else {
		goto L127
	}
L126:
	;
	goto L125
L127:
	;
	goto L105
L128:
	;
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+8)))
	v408 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v410 = v408
	v411 = v407
	goto L111
L129:
	;
	if v411&int32(255) != int32(116) {
		goto L134
	} else {
		goto L135
	}
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+444)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+440)) = l1
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+436)) = uint8(v411)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+428)) = v182
	*(*int32)(unsafe.Add(mBase, uint32(v21)+432)) = v410
	v624 = v21 + int32(428)
	if l4 == int32(3) {
		goto L175
	} else {
		goto L176
	}
L132:
	;
	if v182 == int32(0) {
		v757 = v599
		goto L106
	} else {
		goto L174
	}
L133:
	;
	F_UnpinBuffer(m, v485)
	mBase = m.M
	v584 = m.ExcPending
	if v584 != 0 {
		goto L13
	} else {
		goto L170
	}
L134:
	;
	v421 = F_IOContextForStrategy(m, int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L13
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v570 = F_LocalBufferAlloc(m, v410, l1, v413, v21+int32(411))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L13
	} else {
		goto L168
	}
L137:
	;
	v424 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[5]))
	F_ResourceOwnerEnlarge(m, v424)
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L13
	} else {
		goto L138
	}
L138:
	;
	F_ReservePrivateRefCountEntry(m)
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L13
	} else {
		goto L139
	}
L139:
	;
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v410)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+508)) = v429
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v410)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+512)) = v431
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v410)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+524)) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v21)+520)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v21)+516)) = v433
	v438 = v21 + int32(508)
	v439 = F_BufTableHashCode(m, v438)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L13
	} else {
		goto L140
	}
L140:
	;
	v442 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[6]))
	v449 = v442 + v439&int32(127)<<(uint(int32(7))%32) + int32(_a_F_ExtendBufferedRelTo_1)
	v451 = F_LWLockAcquire(m, v449, int32(1))
	mBase = m.M
	v452 = m.ExcPending
	if v452 != 0 {
		goto L13
	} else {
		goto L141
	}
L141:
	;
	v453 = F_BufTableLookup(m, v438, v439)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L13
	} else {
		goto L142
	}
L142:
	;
	if int32(0) <= v453 {
		goto L143
	} else {
		goto L144
	}
L143:
	;
	v458 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[7]))
	v461 = v458 + v453*int32(56)
	v462 = int32(0)
	v464 = F_PinBuffer(m, v461, v462, v462)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L13
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	F_LWLockRelease(m, v449)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L13
	} else {
		goto L149
	}
L146:
	;
	F_LWLockRelease(m, v449)
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L13
	} else {
		goto L147
	}
L147:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+411)) = uint8(v464)
	if v464 == int32(0) {
		v599 = v461
		goto L132
	} else {
		goto L148
	}
L148:
	;
	v654 = v461
	goto L109
L149:
	;
	v474 = F_GetVictimBuffer(m, int32(0), v421)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L13
	} else {
		goto L150
	}
L150:
	;
	v477 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[7]))
	v479 = F_LWLockAcquire(m, v449, int32(0))
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L13
	} else {
		goto L151
	}
L151:
	;
	v481 = int32(56)
	v483 = v477 + v474*v481
	v485 = v483 - v481
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v483-int32(36))))
	v491 = F_BufTableInsert(m, v21+int32(508), v439, v490)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L13
	} else {
		goto L152
	}
L152:
	;
	if int32(0) <= v491 {
		goto L133
	} else {
		goto L153
	}
L153:
	;
	v495 = F_LockBufHdr(m, v485)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L13
	} else {
		goto L154
	}
L154:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v21)+524))
	*(*int32)(unsafe.Add(mBase, uint32(v485)+16)) = v497
	v499 = *(*int64)(unsafe.Add(mBase, uint32(v21)+516))
	*(*int64)(unsafe.Add(mBase, uint32(v485)+8)) = v499
	v501 = *(*int64)(unsafe.Add(mBase, uint32(v21)+508))
	*(*int64)(unsafe.Add(mBase, uint32(v485))) = v501
	v504 = v483 - int32(32)
	v505 = int64(2181300224)
	if v411&int32(255) == int32(112) {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	v512 = v505
	goto L157
L156:
	;
	v512 = int64(33816576)
	goto L157
L157:
	;
	if l1 == int32(3) {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v515 = v505
	goto L160
L159:
	;
	v515 = v512
	goto L160
L160:
	;
	v520 = base.AtomicRmwCmpxchg64(m, v504, int32(0), v495, v515|v495&int64(-38010881))
	if v495 != v520 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v537 = v520
	goto L164
L162:
	;
	goto L163
L163:
	;
	F_LWLockRelease(m, v449)
	mBase = m.M
	v565 = m.ExcPending
	if v565 != 0 {
		goto L13
	} else {
		goto L167
	}
L164:
	;
	v544 = base.AtomicRmwCmpxchg64(m, v504, int32(0), v537, v537&int64(-38010881)|v515)
	if v537 != v544 {
		v537 = v544
		goto L164
	} else {
		goto L166
	}
L165:
	;
	goto L163
L166:
	;
	goto L165
L167:
	;
	v566 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+411)) = uint8(v566)
	v599 = v485
	goto L132
L168:
	;
	v572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+411)))
	if v572 == int32(0) {
		v599 = v570
		goto L132
	} else {
		goto L169
	}
L169:
	;
	v575 = int32(_a_F_ExtendBufferedRelTo_2)
	v577 = *(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[8]))
	*(*int64)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[8])) = v577 + int64(1)
	v666 = v570
	v667 = int32(3)
	v670 = int32(1)
	goto L108
L170:
	;
	v586 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[7]))
	v589 = v586 + v491*int32(56)
	v590 = int32(0)
	v592 = F_PinBuffer(m, v589, v590, v590)
	mBase = m.M
	v593 = m.ExcPending
	if v593 != 0 {
		goto L13
	} else {
		goto L171
	}
L171:
	;
	F_LWLockRelease(m, v449)
	mBase = m.M
	v595 = m.ExcPending
	if v595 != 0 {
		goto L13
	} else {
		goto L172
	}
L172:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+411)) = uint8(v592)
	if v592 != 0 {
		v654 = v589
		goto L109
	} else {
		goto L173
	}
L173:
	;
	v599 = v589
	goto L132
L174:
	;
	v725 = v599
	goto L107
L175:
	;
	v631 = int32(9)
	goto L177
L176:
	;
	v631 = int32(8)
	goto L177
L177:
	;
	v632 = F_StartReadBuffer(m, v624, v21+int32(424), v413, v631)
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L13
	} else {
		goto L178
	}
L178:
	;
	if v632 != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v634 = F_WaitReadBuffers(m, v624)
	mBase = m.M
	v635 = m.ExcPending
	if v635 != 0 {
		goto L13
	} else {
		goto L182
	}
L180:
	;
	goto L181
L181:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v21)+424))
	v805 = v636
	goto L78
L182:
	;
	goto L181
L183:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L13
	} else {
		goto L184
	}
L184:
	;
	F_errmsg(m, int32(_a_F_ExtendBufferedRelTo_3), int32(0))
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L13
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(_a_F_ExtendBufferedRelTo_4), int32(1296), int32(_a_F_ExtendBufferedRelTo_5))
	mBase = m.M
	v652 = m.ExcPending
	if v652 != 0 {
		goto L13
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	v697 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[9])))
	if v697 == int32(1) {
		goto L188
	} else {
		goto L189
	}
L188:
	;
	v700 = int32(_a_F_ExtendBufferedRelTo_6)
	v702 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[10]))
	v704 = *(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[11]))
	*(*int32)(unsafe.Add(mBase, _c_F_ExtendBufferedRelTo[10])) = v702 + v704
	goto L190
L189:
	;
	goto L190
L190:
	;
	if v182 == int32(0) {
		v757 = v666
		goto L106
	} else {
		goto L191
	}
L191:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v182)+272))
	if v709 == int32(0) {
		goto L192
	} else {
		goto L193
	}
L192:
	;
	v712 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+268)))
	if v712 != int32(1) {
		v725 = v666
		goto L107
	} else {
		goto L195
	}
L193:
	;
	v718 = v709
	goto L194
L194:
	;
	v719 = *(*int64)(unsafe.Add(mBase, uint32(v718)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v718)+120)) = v719 + int64(1)
	v725 = v666
	goto L107
L195:
	;
	F_pgstat_assoc_relation(m, v182)
	mBase = m.M
	v716 = m.ExcPending
	if v716 != 0 {
		goto L13
	} else {
		goto L196
	}
L196:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v182)+272))
	v718 = v717
	goto L194
L197:
	;
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+268)))
	if v744 != int32(1) {
		v757 = v725
		goto L106
	} else {
		goto L200
	}
L198:
	;
	v750 = v741
	goto L199
L199:
	;
	v751 = *(*int64)(unsafe.Add(mBase, uint32(v750)+112))
	*(*int64)(unsafe.Add(mBase, uint32(v750)+112)) = v751 + int64(1)
	v757 = v725
	goto L106
L200:
	;
	F_pgstat_assoc_relation(m, v182)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L13
	} else {
		goto L201
	}
L201:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v182)+272))
	v750 = v749
	goto L199
L202:
	;
	v805 = v775
	goto L78
L203:
	;
	v794 = int32(9)
	goto L205
L204:
	;
	v794 = int32(1)
	goto L205
L205:
	;
	v795 = F_ExtendBufferedRel(m, v21, l1, int32(0), v794)
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L13
	} else {
		goto L206
	}
L206:
	;
	v805 = v795
	goto L78
}
func F_ExtractReplicaIdentity(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	v5 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(_a_F_ExtractReplicaIdentity_0)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+130)))
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v21 = *(*int32)(unsafe.Add(mBase, _c_F_ExtractReplicaIdentity[0]))
	if v21 <= int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v13 + int32(_a_F_ExtractReplicaIdentity_0)
	return v132
L2:
	;
	v25 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_ExtractReplicaIdentity[1])))
	if v25&int32(1) == int32(0) {
		v132 = v5
		goto L1
	} else {
		goto L5
	}
L3:
	;
	goto L4
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+118)))
	if v31 != int32(112) {
		v132 = v5
		goto L1
	} else {
		goto L6
	}
L5:
	;
	goto L4
L6:
	;
	if v21 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v36 != 0 {
		v132 = v5
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+119)))
	if v38 == int32(102) {
		v132 = v5
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v37 != 0 {
		v132 = v5
		goto L1
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	goto L13
L13:
	;
	if base.Ui32(v41) < base.Ui32(int32(_a_F_ExtractReplicaIdentity_1)) {
		v132 = v5
		goto L1
	} else {
		goto L14
	}
L14:
	;
	switch v17 - int32(102) {
	case 0:
		goto L16
	default:
		goto L15
	case 8:
		v132 = v5
		goto L1
	}
L15:
	;
	if l2 == int32(0) {
		v132 = v5
		goto L1
	} else {
		goto L22
	}
L16:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46)+20)))
	if v47&int32(4) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v132 = l1
	goto L1
L18:
	;
	goto L19
L19:
	;
	v52 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v52)
	v54 = F_toast_flatten_tuple(m, l1, v15)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	v132 = v54
	goto L1
L22:
	;
	v61 = F_RelationGetIndexAttrBitmap(m, l0, int32(2))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L20
	} else {
		goto L23
	}
L23:
	;
	if v61 == int32(0) {
		v132 = v5
		goto L1
	} else {
		goto L24
	}
L24:
	;
	F_heap_deform_tuple(m, l1, v15, v13, v13+int32(_a_F_ExtractReplicaIdentity_2))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L20
	} else {
		goto L25
	}
L25:
	;
	v69 = int32(0)
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v69 < v70 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v73 = v69
	goto L29
L27:
	;
	goto L28
L28:
	;
	v110 = F_heap_form_tuple(m, v15, v13, v13+int32(_a_F_ExtractReplicaIdentity_2))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L20
	} else {
		goto L36
	}
L29:
	;
	v85 = F_bms_is_member(m, v73+int32(8), v61)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L20
	} else {
		goto L31
	}
L30:
	;
	goto L28
L31:
	;
	if v85 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v92 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v13+int32(_a_F_ExtractReplicaIdentity_2)+v73))) = uint8(v92)
	goto L34
L33:
	;
	goto L34
L34:
	;
	v95 = v73 + int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
	if v95 < v96 {
		v73 = v95
		goto L29
	} else {
		goto L35
	}
L35:
	;
	goto L30
L36:
	;
	v112 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v112)
	F_bms_free(m, v61)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L20
	} else {
		goto L37
	}
L37:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v110)+16))
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v116)+20)))
	if v117&int32(4) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v132 = v110
	goto L1
L39:
	;
	goto L40
L40:
	;
	v122 = F_toast_flatten_tuple(m, v110, v15)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L20
	} else {
		goto L41
	}
L41:
	;
	F_pfree(m, v110)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L20
	} else {
		goto L42
	}
L42:
	;
	v132 = v122
	goto L1
}
func F___env_rm_add(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	v9 = *(*int32)(unsafe.Add(mBase, _c_F___env_rm_add[0]))
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v11 = *(*int32)(unsafe.Add(mBase, _c_F___env_rm_add[1]))
	v13 = l1
	v15 = int32(0)
	goto L4
L2:
	;
	v38 = l1
	goto L3
L3:
	;
	if v38 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L4:
	;
	v21 = v11 + v15<<(uint(int32(2))%32)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if l0 == v22 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v38 = v33
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v13
	F_emscripten_builtin_free(m, l0)
	mBase = m.M
	return
L7:
	;
	goto L8
L8:
	;
	v26 = int32(0)
	if v22|base.B2i32(v13 == v26) == v26 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v13
	v33 = int32(0)
	goto L11
L10:
	;
	v33 = v13
	goto L11
L11:
	;
	v35 = v15 + int32(1)
	if v35 != v9 {
		v13 = v33
		v15 = v35
		goto L4
	} else {
		goto L12
	}
L12:
	;
	goto L5
L13:
	;
	return
L14:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F___env_rm_add[1]))
	v52 = F_emscripten_builtin_realloc(m, v47, v9<<(uint(int32(2))%32)+int32(4))
	mBase = m.M
	if v52 == int32(0) {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, _c_F___env_rm_add[1])) = v52
	v57 = int32(_a_F___env_rm_add_0)
	v59 = *(*int32)(unsafe.Add(mBase, _c_F___env_rm_add[0]))
	*(*int32)(unsafe.Add(mBase, _c_F___env_rm_add[0])) = v59 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v52+v59<<(uint(int32(2))%32)))) = v38
	goto L13
}
func F__exit(m *base.Module, l0 int32) {
	F__Exit(m, l0)
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_each(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_each(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_each_object_field_start(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+32))
	if v5 == int32(1) {
		v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)))
		if v8 != int32(1) {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
			*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v18
			return int32(0)
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v4)+28))
			if v11 != int32(1) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
				*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v18
				return int32(0)
			} else {
				v14 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+21)) = uint8(v14)
				return int32(0)
			}
		}
	} else {
		return int32(0)
	}
}
func F_eclass_member_iterator_next(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	v2 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	if v6 == v2 {
		v110 = v2
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v110
L2:
	;
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v97 = v91 + int32(4)
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v92)+4))
	if base.Ui32(v97) < base.Ui32(v93+v99<<(uint(int32(2))%32)) {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
	v91 = v9
	v92 = v6
	v93 = v10
	goto L3
L5:
	;
	goto L6
L6:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = v11
	goto L7
L7:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v17 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v84)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v88
	v91 = v88
	v92 = v84
	v93 = v88
	goto L3
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v73
	if v73 <= int32(0) {
		v110 = v2
		goto L1
	} else {
		goto L20
	}
L10:
	;
	v73 = base.I32_ctz(v59) | v60<<(uint(int32(5))%32)
	goto L9
L11:
	;
	v73 = int32(-2)
	goto L9
L12:
	;
	v24 = v13 + int32(1)
	v26 = int32(base.Ui32(v24) >> (uint(int32(5)) % 32))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v27 <= v26 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v30 = v17 + int32(8)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v26<<(uint(int32(2))%32))))
	v37 = v34 & (int32(-1) << (uint(v24) % 32))
	if v37 != 0 {
		v59 = v37
		v60 = v26
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v39 = v26 + int32(1)
	if v39 == v27 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	v42 = v39
	goto L16
L16:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v30+v42<<(uint(int32(2))%32))))
	if v49 != 0 {
		v59 = v49
		v60 = v42
		goto L10
	} else {
		goto L18
	}
L17:
	;
	goto L11
L18:
	;
	v51 = v42 + int32(1)
	if v51 != v27 {
		v42 = v51
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)+12))
	if v78 <= v73 {
		v110 = v2
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77)+20))
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v80+v73<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v84
	if v84 == int32(0) {
		v13 = v73
		goto L7
	} else {
		goto L22
	}
L22:
	;
	goto L8
L23:
	;
	v104 = v97
	goto L25
L24:
	;
	v104 = int32(0)
	goto L25
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v104
	v110 = v95
	goto L1
}
func F_elem_contained_by_range(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
		if v19 != 0 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
			if v20 == v17 {
				v30 = v19
				v31 = F_range_contains_elem_internal(m, v30, v13, v11)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					m.G0 = v9 + int32(16)
					return base.I64_extend_i32_u(v31)
				}
			} else {
				v23 = F_lookup_type_cache(m, v17, int32(2048))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
					if v25 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int64(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
							F_errmsg_internal(m, int32(_a_F_elem_contained_by_range_0), v9)
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_elem_contained_by_range_1), int32(1946), int32(_a_F_elem_contained_by_range_2))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
						v30 = v23
						v31 = F_range_contains_elem_internal(m, v30, v13, v11)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int64(0)
						} else {
							m.G0 = v9 + int32(16)
							return base.I64_extend_i32_u(v31)
						}
					}
				}
			}
		} else {
			v23 = F_lookup_type_cache(m, v17, int32(2048))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v23)+200))
				if v25 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v17
						F_errmsg_internal(m, int32(_a_F_elem_contained_by_range_0), v9)
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_elem_contained_by_range_1), int32(1946), int32(_a_F_elem_contained_by_range_2))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v23
					v30 = v23
					v31 = F_range_contains_elem_internal(m, v30, v13, v11)
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int64(0)
					} else {
						m.G0 = v9 + int32(16)
						return base.I64_extend_i32_u(v31)
					}
				}
			}
		}
	}
}
func F_english_UTF_8_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v197 int32
	_ = v197
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v274 int32
	_ = v274
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v455 int32
	_ = v455
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v473 int32
	_ = v473
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v493 int32
	_ = v493
	var v506 int32
	_ = v506
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v577 int32
	_ = v577
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v599 int32
	_ = v599
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v628 int32
	_ = v628
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v662 int32
	_ = v662
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v679 int32
	_ = v679
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v702 int32
	_ = v702
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v724 int32
	_ = v724
	var v734 int32
	_ = v734
	var v736 int32
	_ = v736
	var v740 int32
	_ = v740
	var v753 int32
	_ = v753
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v773 int32
	_ = v773
	var v779 int32
	_ = v779
	var v786 int32
	_ = v786
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v824 int32
	_ = v824
	var v831 int32
	_ = v831
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v856 int32
	_ = v856
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v875 int32
	_ = v875
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v901 int32
	_ = v901
	var v909 int32
	_ = v909
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v962 int32
	_ = v962
	var v966 int32
	_ = v966
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v983 int32
	_ = v983
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v993 int32
	_ = v993
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v1002 int32
	_ = v1002
	var v1004 int32
	_ = v1004
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1017 int32
	_ = v1017
	var v1022 int32
	_ = v1022
	var v1026 int32
	_ = v1026
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1047 int32
	_ = v1047
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1057 int32
	_ = v1057
	var v1063 int32
	_ = v1063
	var v1064 int32
	_ = v1064
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1075 int32
	_ = v1075
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1084 int32
	_ = v1084
	var v1090 int32
	_ = v1090
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1120 int32
	_ = v1120
	var v1137 int32
	_ = v1137
	var v1138 int32
	_ = v1138
	var v1147 int32
	_ = v1147
	var v1154 int32
	_ = v1154
	var v1155 int32
	_ = v1155
	var v1157 int32
	_ = v1157
	var v1159 int32
	_ = v1159
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1185 int32
	_ = v1185
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1207 int32
	_ = v1207
	var v1209 int32
	_ = v1209
	var v1217 int32
	_ = v1217
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1229 int32
	_ = v1229
	var v1238 int32
	_ = v1238
	var v1253 int32
	_ = v1253
	var v1256 int32
	_ = v1256
	var v1259 int32
	_ = v1259
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1290 int32
	_ = v1290
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1311 int32
	_ = v1311
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1334 int32
	_ = v1334
	var v1349 int32
	_ = v1349
	var v1350 int32
	_ = v1350
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1366 int32
	_ = v1366
	var v1367 int32
	_ = v1367
	var v1383 int32
	_ = v1383
	var v1384 int32
	_ = v1384
	var v1386 int32
	_ = v1386
	var v1388 int32
	_ = v1388
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1418 int32
	_ = v1418
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1446 int32
	_ = v1446
	var v1450 int32
	_ = v1450
	var v1452 int32
	_ = v1452
	var v1458 int32
	_ = v1458
	var v1475 int32
	_ = v1475
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1529 int32
	_ = v1529
	var v1536 int32
	_ = v1536
	var v1537 int32
	_ = v1537
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1548 int32
	_ = v1548
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1554 int32
	_ = v1554
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1571 int32
	_ = v1571
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1599 int32
	_ = v1599
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1611 int32
	_ = v1611
	var v1620 int32
	_ = v1620
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1644 int32
	_ = v1644
	var v1648 int32
	_ = v1648
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1653 int32
	_ = v1653
	var v1665 int32
	_ = v1665
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1707 int32
	_ = v1707
	var v1709 int32
	_ = v1709
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1739 int32
	_ = v1739
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1779 int32
	_ = v1779
	var v1795 int32
	_ = v1795
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1815 int32
	_ = v1815
	var v1821 int32
	_ = v1821
	var v1826 int32
	_ = v1826
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1834 int32
	_ = v1834
	var v1835 int32
	_ = v1835
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1851 int32
	_ = v1851
	var v1856 int32
	_ = v1856
	var v1861 int32
	_ = v1861
	var v1864 int32
	_ = v1864
	var v1871 int32
	_ = v1871
	var v1873 int32
	_ = v1873
	var v1874 int32
	_ = v1874
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1884 int32
	_ = v1884
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1898 int32
	_ = v1898
	var v1900 int32
	_ = v1900
	var v1906 int32
	_ = v1906
	var v1911 int32
	_ = v1911
	var v1915 int32
	_ = v1915
	var v1918 int32
	_ = v1918
	var v1922 int32
	_ = v1922
	var v1936 int32
	_ = v1936
	var v1941 int32
	_ = v1941
	var v1949 int32
	_ = v1949
	var v1952 int32
	_ = v1952
	var v1954 int32
	_ = v1954
	var v1958 int32
	_ = v1958
	var v1964 int32
	_ = v1964
	var v1980 int32
	_ = v1980
	var v1981 int32
	_ = v1981
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2028 int32
	_ = v2028
	var v2030 int32
	_ = v2030
	var v2032 int32
	_ = v2032
	var v2050 int32
	_ = v2050
	var v2052 int32
	_ = v2052
	var v2060 int32
	_ = v2060
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2072 int32
	_ = v2072
	var v2089 int32
	_ = v2089
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2098 int32
	_ = v2098
	var v2102 int32
	_ = v2102
	var v2103 int32
	_ = v2103
	var v2107 int32
	_ = v2107
	var v2109 int32
	_ = v2109
	var v2113 int32
	_ = v2113
	var v2114 int32
	_ = v2114
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2157 int32
	_ = v2157
	var v2158 int32
	_ = v2158
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2169 int32
	_ = v2169
	var v2170 int32
	_ = v2170
	var v2175 int32
	_ = v2175
	var v2176 int32
	_ = v2176
	var v2181 int32
	_ = v2181
	var v2182 int32
	_ = v2182
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2223 int32
	_ = v2223
	var v2227 int32
	_ = v2227
	var v2235 int32
	_ = v2235
	var v2236 int32
	_ = v2236
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2257 int32
	_ = v2257
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2275 int32
	_ = v2275
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2280 int32
	_ = v2280
	var v2287 int32
	_ = v2287
	var v2289 int32
	_ = v2289
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2310 int32
	_ = v2310
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2338 int32
	_ = v2338
	var v2342 int32
	_ = v2342
	var v2344 int32
	_ = v2344
	var v2350 int32
	_ = v2350
	var v2366 int32
	_ = v2366
	var v2373 int32
	_ = v2373
	var v2374 int32
	_ = v2374
	var v2379 int32
	_ = v2379
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2389 int32
	_ = v2389
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2397 int32
	_ = v2397
	var v2412 int32
	_ = v2412
	var v2413 int32
	_ = v2413
	var v2416 int32
	_ = v2416
	var v2418 int32
	_ = v2418
	var v2424 int32
	_ = v2424
	var v2425 int32
	_ = v2425
	var v2430 int32
	_ = v2430
	var v2431 int32
	_ = v2431
	var v2436 int32
	_ = v2436
	var v2437 int32
	_ = v2437
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2446 int32
	_ = v2446
	var v2449 int32
	_ = v2449
	var v2451 int32
	_ = v2451
	var v2456 int32
	_ = v2456
	var v2461 int32
	_ = v2461
	var v2463 int32
	_ = v2463
	var v2467 int32
	_ = v2467
	var v2468 int32
	_ = v2468
	var v2470 int32
	_ = v2470
	var v2472 int32
	_ = v2472
	var v2487 int32
	_ = v2487
	var v2488 int32
	_ = v2488
	var v2491 int32
	_ = v2491
	var v2493 int32
	_ = v2493
	var v2497 int32
	_ = v2497
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2504 int32
	_ = v2504
	var v2506 int32
	_ = v2506
	var v2516 int32
	_ = v2516
	var v2521 int32
	_ = v2521
	var v2526 int32
	_ = v2526
	var v2528 int32
	_ = v2528
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2537 int32
	_ = v2537
	var v2543 int32
	_ = v2543
	var v2544 int32
	_ = v2544
	var v2547 int32
	_ = v2547
	var v2551 int32
	_ = v2551
	var v2553 int32
	_ = v2553
	var v2555 int32
	_ = v2555
	var v2558 int32
	_ = v2558
	var v2564 int32
	_ = v2564
	var v2569 int32
	_ = v2569
	var v2573 int32
	_ = v2573
	var v2574 int32
	_ = v2574
	var v2577 int32
	_ = v2577
	var v2578 int32
	_ = v2578
	var v2585 int32
	_ = v2585
	var v2590 int32
	_ = v2590
	var v2591 int32
	_ = v2591
	var v2592 int32
	_ = v2592
	var v2594 int32
	_ = v2594
	var v2599 int32
	_ = v2599
	var v2604 int32
	_ = v2604
	var v2605 int32
	_ = v2605
	var v2610 int32
	_ = v2610
	var v2613 int32
	_ = v2613
	var v2615 int32
	_ = v2615
	var v2617 int32
	_ = v2617
	var v2621 int32
	_ = v2621
	var v2627 int32
	_ = v2627
	var v2634 int32
	_ = v2634
	var v2638 int32
	_ = v2638
	var v2640 int32
	_ = v2640
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2659 int32
	_ = v2659
	var v2668 int32
	_ = v2668
	var v2670 int32
	_ = v2670
	var v2675 int32
	_ = v2675
	var v2677 int32
	_ = v2677
	var v2684 int32
	_ = v2684
	var v2687 int32
	_ = v2687
	var v2691 int32
	_ = v2691
	var v2698 int32
	_ = v2698
	var v2699 int32
	_ = v2699
	var v2713 int32
	_ = v2713
	var v2718 int32
	_ = v2718
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2729 int32
	_ = v2729
	var v2737 int32
	_ = v2737
	var v2749 int32
	_ = v2749
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	v10 = v7 + int32(2)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v11 <= v10 {
		v90 = v11
		goto L3
	} else {
		goto L4
	}
L1:
	;
	return v2749
L2:
	;
	v2749 = int32(1)
	goto L1
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L36
L4:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13+v10))))
	if base.B2i32(v15&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v15)%32)&int32(42750482) == int32(0)) != 0 {
		v90 = v11
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = F_find_among(m, l0, int32(_a_F_english_UTF_8_stem_0), int32(15), int32(0))
	mBase = m.M
	v33 = m.ExcPending
	if v33 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	return int32(0)
L7:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v30 == int32(0) {
		v90 = v34
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
	if v37 < v34 {
		v90 = v34
		goto L3
	} else {
		goto L9
	}
L9:
	;
	switch v30 - int32(1) {
	case 0:
		goto L17
	case 1:
		goto L16
	case 2:
		goto L15
	case 3:
		goto L14
	case 4:
		goto L13
	case 5:
		goto L12
	case 6:
		goto L11
	case 7:
		goto L10
	default:
		goto L2
	}
L10:
	;
	v86 = F_slice_from_s(m, l0, int32(5), int32(_a_F_english_UTF_8_stem_1))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L6
	} else {
		goto L32
	}
L11:
	;
	v80 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_2))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L6
	} else {
		goto L30
	}
L12:
	;
	v74 = F_slice_from_s(m, l0, int32(5), int32(_a_F_english_UTF_8_stem_3))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L6
	} else {
		goto L28
	}
L13:
	;
	v68 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_4))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L6
	} else {
		goto L26
	}
L14:
	;
	v62 = F_slice_from_s(m, l0, int32(5), int32(_a_F_english_UTF_8_stem_5))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L6
	} else {
		goto L24
	}
L15:
	;
	v56 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_6))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L6
	} else {
		goto L22
	}
L16:
	;
	v50 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_7))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L6
	} else {
		goto L20
	}
L17:
	;
	v44 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_8))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	if int32(0) <= v44 {
		goto L2
	} else {
		goto L19
	}
L19:
	;
	v2749 = v44
	goto L1
L20:
	;
	if int32(0) <= v50 {
		goto L2
	} else {
		goto L21
	}
L21:
	;
	v2749 = v50
	goto L1
L22:
	;
	if int32(0) <= v56 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	v2749 = v56
	goto L1
L24:
	;
	if int32(0) <= v62 {
		goto L2
	} else {
		goto L25
	}
L25:
	;
	v2749 = v62
	goto L1
L26:
	;
	if int32(0) <= v68 {
		goto L2
	} else {
		goto L27
	}
L27:
	;
	v2749 = v68
	goto L1
L28:
	;
	if int32(0) <= v74 {
		goto L2
	} else {
		goto L29
	}
L29:
	;
	v2749 = v74
	goto L1
L30:
	;
	if int32(0) <= v80 {
		goto L2
	} else {
		goto L31
	}
L31:
	;
	v2749 = v80
	goto L1
L32:
	;
	if int32(0) <= v86 {
		goto L2
	} else {
		goto L33
	}
L33:
	;
	v2749 = v86
	goto L1
L34:
	;
	if int32(0) <= v146 {
		goto L54
	} else {
		goto L55
	}
L36:
	;
	goto L37
L37:
	;
	goto L38
L38:
	;
	v101 = v7
	v103 = int32(3)
	goto L41
L40:
	;
	v146 = v131
	goto L34
L41:
	;
	if v90 <= v101 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L40
L43:
	;
	v146 = int32(-1)
	goto L34
L44:
	;
	goto L45
L45:
	;
	v108 = v101 + int32(1)
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v94+v101))))
	if base.Ui32(v110) < base.Ui32(int32(192)) {
		v131 = v108
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v132 = int32(1)
	if v132 < v103 {
		v101 = v131
		v103 = v103 - v132
		goto L41
	} else {
		goto L53
	}
L47:
	;
	if v90 <= v108 {
		v131 = v108
		goto L46
	} else {
		goto L48
	}
L48:
	;
	v117 = v108
	goto L49
L49:
	;
	v120 = int32(*(*int8)(unsafe.Add(mBase, uint32(v94+v117))))
	if int32(-65) < v120 {
		v131 = v117
		goto L46
	} else {
		goto L51
	}
L50:
	;
	v131 = v90
	goto L46
L51:
	;
	v124 = v117 + int32(1)
	if v124 != v90 {
		v117 = v124
		goto L49
	} else {
		goto L52
	}
L52:
	;
	goto L50
L53:
	;
	goto L42
L54:
	;
	v149 = v146
	goto L56
L55:
	;
	v149 = v7
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v149
	if v146 < int32(0) {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	v153 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v153)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v7 == v157 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v197 = v7
	goto L68
L59:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v161 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159+v7))))
	if v161 == int32(39) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v165 = v7 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v165
	v168 = F_slice_del(m, l0)
	mBase = m.M
	if v168 < int32(0) {
		v2749 = v168
		goto L1
	} else {
		goto L63
	}
L61:
	;
	v177 = v159
	goto L62
L62:
	;
	v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v177+v7))))
	if v179 != int32(121) {
		goto L58
	} else {
		goto L65
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v7
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v7 == v173 {
		goto L58
	} else {
		goto L64
	}
L64:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v177 = v175
	goto L62
L65:
	;
	v182 = int32(1)
	v183 = v7 + v182
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v183
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v183
	v188 = F_slice_from_s(m, l0, v182, int32(_a_F_english_UTF_8_stem_9))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L6
	} else {
		goto L66
	}
L66:
	;
	if v188 < int32(0) {
		v2749 = v188
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v192 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v192)
	goto L58
L68:
	;
	v213 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L75
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v404
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v404
	v408 = v7 + int32(3)
	if v404 <= v408 {
		goto L129
	} else {
		goto L130
	}
L70:
	;
	goto L69
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v197
	v391 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v322 + v391
	v396 = F_slice_from_s(m, l0, v391, int32(_a_F_english_UTF_8_stem_10))
	mBase = m.M
	v397 = m.ExcPending
	if v397 != 0 {
		goto L6
	} else {
		goto L125
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v197
	goto L106
L73:
	;
	if v319 != 0 {
		goto L97
	} else {
		goto L98
	}
L74:
	;
	v319 = v312
	goto L73
L75:
	;
	if v214 <= v213 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v312 = int32(0)
	goto L74
L77:
	;
	v319 = int32(-1)
	goto L73
L78:
	;
	goto L79
L79:
	;
	v230 = int32(1)
	v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v213+v215))))
	if base.Ui32(v232) < base.Ui32(int32(192)) {
		v289 = v232
		v290 = v230
		goto L80
	} else {
		goto L81
	}
L80:
	;
	if int32(121) < v289 {
		v312 = v290
		goto L74
	} else {
		goto L93
	}
L81:
	;
	v236 = v213 + int32(1)
	if v236 == v214 {
		v289 = v232
		v290 = v230
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236+v215))))
	v241 = v239 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v232) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v255 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v245+v215))))
	v257 = v255 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v232) {
		goto L89
	} else {
		goto L90
	}
L84:
	;
	v245 = v213 + int32(2)
	if v245 != v214 {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v289 = v232<<(uint(int32(6))%32)&int32(1984) | v241
	v290 = int32(2)
	goto L80
L87:
	;
	goto L86
L88:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215+v261))))
	v289 = v274&int32(63) | (v232<<(uint(int32(18))%32)&int32(_a_F_english_UTF_8_stem_11) | v241<<(uint(int32(12))%32) | v257<<(uint(int32(6))%32))
	v290 = int32(4)
	goto L80
L89:
	;
	v261 = v213 + int32(3)
	if v261 != v214 {
		goto L88
	} else {
		goto L92
	}
L90:
	;
	goto L91
L91:
	;
	v289 = v232<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v241<<(uint(int32(6))%32) | v257
	v290 = int32(3)
	goto L80
L92:
	;
	goto L91
L93:
	;
	v294 = v289 - int32(97)
	if v294 < int32(0) {
		v312 = v290
		goto L74
	} else {
		goto L94
	}
L94:
	;
	v300 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v294)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[0]))))
	if int32(base.Ui32(v300)>>(uint(v294&int32(7))%32))&int32(1) == int32(0) {
		v312 = v290
		goto L74
	} else {
		goto L95
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v290 + v213
	goto L96
L96:
	;
	goto L76
L97:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v331 = v320
	v332 = v321
	goto L72
L98:
	;
	goto L99
L99:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v322
	v324 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v325 == v322 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v331 = v322
	v332 = v324
	goto L72
L101:
	;
	goto L102
L102:
	;
	v328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v324+v322))))
	if v328 == int32(121) {
		goto L71
	} else {
		goto L103
	}
L103:
	;
	v331 = v325
	v332 = v324
	goto L72
L104:
	;
	if v386 < int32(0) {
		goto L70
	} else {
		goto L124
	}
L106:
	;
	goto L107
L107:
	;
	goto L108
L108:
	;
	v341 = v197
	v343 = int32(1)
	goto L111
L110:
	;
	v386 = v371
	goto L104
L111:
	;
	if v331 <= v341 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	goto L110
L113:
	;
	v386 = int32(-1)
	goto L104
L114:
	;
	goto L115
L115:
	;
	v348 = v341 + int32(1)
	v350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v332+v341))))
	if base.Ui32(v350) < base.Ui32(int32(192)) {
		v371 = v348
		goto L116
	} else {
		goto L117
	}
L116:
	;
	v372 = int32(1)
	if v372 < v343 {
		v341 = v371
		v343 = v343 - v372
		goto L111
	} else {
		goto L123
	}
L117:
	;
	if v331 <= v348 {
		v371 = v348
		goto L116
	} else {
		goto L118
	}
L118:
	;
	v357 = v348
	goto L119
L119:
	;
	v360 = int32(*(*int8)(unsafe.Add(mBase, uint32(v332+v357))))
	if int32(-65) < v360 {
		v371 = v357
		goto L116
	} else {
		goto L121
	}
L120:
	;
	v371 = v331
	goto L116
L121:
	;
	v364 = v357 + int32(1)
	if v364 != v331 {
		v357 = v364
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	goto L112
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v386
	v197 = v386
	goto L68
L125:
	;
	if v396 < int32(0) {
		v2749 = v396
		goto L1
	} else {
		goto L126
	}
L126:
	;
	v400 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)) = uint8(v400)
	v402 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v197 = v402
	goto L68
L127:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v7
	v928 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v928
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v928
	if v928 <= v7 {
		v957 = v928
		goto L236
	} else {
		goto L237
	}
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v679
	v692 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v694 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v702 = v692
	goto L187
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v7
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v455 = v7
	goto L136
L130:
	;
	v410 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v410+v408))))
	if base.B2i32(v412&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v412)%32)&int32(_a_F_english_UTF_8_stem_13) == int32(0)) != 0 {
		goto L129
	} else {
		goto L131
	}
L131:
	;
	v427 = F_find_among(m, l0, int32(_a_F_english_UTF_8_stem_14), int32(9), int32(0))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L6
	} else {
		goto L132
	}
L132:
	;
	if v427 == int32(0) {
		goto L129
	} else {
		goto L133
	}
L133:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v679 = v431
	goto L128
L134:
	;
	if v550 < int32(0) {
		goto L127
	} else {
		goto L159
	}
L135:
	;
	v550 = v522
	goto L134
L136:
	;
	if v446 <= v455 {
		goto L138
	} else {
		goto L139
	}
L138:
	;
	v550 = int32(-1)
	goto L134
L139:
	;
	goto L140
L140:
	;
	v462 = int32(1)
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455+v447))))
	if base.Ui32(v464) < base.Ui32(int32(192)) {
		v521 = v464
		v522 = v462
		goto L141
	} else {
		goto L142
	}
L141:
	;
	if int32(121) < v521 {
		goto L154
	} else {
		goto L155
	}
L142:
	;
	v468 = v455 + int32(1)
	if v468 == v446 {
		v521 = v464
		v522 = v462
		goto L141
	} else {
		goto L143
	}
L143:
	;
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v468+v447))))
	v473 = v471 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v464) {
		goto L145
	} else {
		goto L146
	}
L144:
	;
	v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v477+v447))))
	v489 = v487 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v464) {
		goto L150
	} else {
		goto L151
	}
L145:
	;
	v477 = v455 + int32(2)
	if v477 != v446 {
		goto L144
	} else {
		goto L148
	}
L146:
	;
	goto L147
L147:
	;
	v521 = v464<<(uint(int32(6))%32)&int32(1984) | v473
	v522 = int32(2)
	goto L141
L148:
	;
	goto L147
L149:
	;
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447+v493))))
	v521 = v506&int32(63) | (v464<<(uint(int32(18))%32)&int32(_a_F_english_UTF_8_stem_11) | v473<<(uint(int32(12))%32) | v489<<(uint(int32(6))%32))
	v522 = int32(4)
	goto L141
L150:
	;
	v493 = v455 + int32(3)
	if v493 != v446 {
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v521 = v464<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v473<<(uint(int32(6))%32) | v489
	v522 = int32(3)
	goto L141
L153:
	;
	goto L152
L154:
	;
	v539 = v522 + v455
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v539
	v455 = v539
	goto L136
L155:
	;
	v526 = v521 - int32(97)
	if v526 < int32(0) {
		goto L154
	} else {
		goto L156
	}
L156:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v526)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[0]))))
	if int32(base.Ui32(v532)>>(uint(v526&int32(7))%32))&int32(1) != 0 {
		goto L135
	} else {
		goto L157
	}
L157:
	;
	goto L154
L159:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v554 = v553 + v550
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v554
	v568 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v577 = v554
	goto L162
L160:
	;
	if v673 < int32(0) {
		goto L127
	} else {
		goto L184
	}
L161:
	;
	v673 = v644
	goto L160
L162:
	;
	if v568 <= v577 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	v673 = int32(-1)
	goto L160
L165:
	;
	goto L166
L166:
	;
	v584 = int32(1)
	v586 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v577+v569))))
	if base.Ui32(v586) < base.Ui32(int32(192)) {
		v643 = v586
		v644 = v584
		goto L167
	} else {
		goto L168
	}
L167:
	;
	if int32(121) < v643 {
		goto L161
	} else {
		goto L180
	}
L168:
	;
	v590 = v577 + int32(1)
	if v590 == v568 {
		v643 = v586
		v644 = v584
		goto L167
	} else {
		goto L169
	}
L169:
	;
	v593 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590+v569))))
	v595 = v593 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v586) {
		goto L171
	} else {
		goto L172
	}
L170:
	;
	v609 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599+v569))))
	v611 = v609 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v586) {
		goto L176
	} else {
		goto L177
	}
L171:
	;
	v599 = v577 + int32(2)
	if v599 != v568 {
		goto L170
	} else {
		goto L174
	}
L172:
	;
	goto L173
L173:
	;
	v643 = v586<<(uint(int32(6))%32)&int32(1984) | v595
	v644 = int32(2)
	goto L167
L174:
	;
	goto L173
L175:
	;
	v628 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569+v615))))
	v643 = v628&int32(63) | (v586<<(uint(int32(18))%32)&int32(_a_F_english_UTF_8_stem_11) | v595<<(uint(int32(12))%32) | v611<<(uint(int32(6))%32))
	v644 = int32(4)
	goto L167
L176:
	;
	v615 = v577 + int32(3)
	if v615 != v568 {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v643 = v586<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v595<<(uint(int32(6))%32) | v611
	v644 = int32(3)
	goto L167
L179:
	;
	goto L178
L180:
	;
	v648 = v643 - int32(97)
	if v648 < int32(0) {
		goto L161
	} else {
		goto L181
	}
L181:
	;
	v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v648)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[0]))))
	if int32(base.Ui32(v654)>>(uint(v648&int32(7))%32))&int32(1) == int32(0) {
		goto L161
	} else {
		goto L182
	}
L182:
	;
	v662 = v644 + v577
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v662
	v577 = v662
	goto L162
L184:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v677 = v676 + v673
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v677
	v679 = v677
	goto L128
L185:
	;
	if v797 < int32(0) {
		goto L127
	} else {
		goto L210
	}
L186:
	;
	v797 = v769
	goto L185
L187:
	;
	if v693 <= v702 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v797 = int32(-1)
	goto L185
L190:
	;
	goto L191
L191:
	;
	v709 = int32(1)
	v711 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v702+v694))))
	if base.Ui32(v711) < base.Ui32(int32(192)) {
		v768 = v711
		v769 = v709
		goto L192
	} else {
		goto L193
	}
L192:
	;
	if int32(121) < v768 {
		goto L205
	} else {
		goto L206
	}
L193:
	;
	v715 = v702 + int32(1)
	if v715 == v693 {
		v768 = v711
		v769 = v709
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v718 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v715+v694))))
	v720 = v718 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v711) {
		goto L196
	} else {
		goto L197
	}
L195:
	;
	v734 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v724+v694))))
	v736 = v734 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v711) {
		goto L201
	} else {
		goto L202
	}
L196:
	;
	v724 = v702 + int32(2)
	if v724 != v693 {
		goto L195
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v768 = v711<<(uint(int32(6))%32)&int32(1984) | v720
	v769 = int32(2)
	goto L192
L199:
	;
	goto L198
L200:
	;
	v753 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v694+v740))))
	v768 = v753&int32(63) | (v711<<(uint(int32(18))%32)&int32(_a_F_english_UTF_8_stem_11) | v720<<(uint(int32(12))%32) | v736<<(uint(int32(6))%32))
	v769 = int32(4)
	goto L192
L201:
	;
	v740 = v702 + int32(3)
	if v740 != v693 {
		goto L200
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	v768 = v711<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v720<<(uint(int32(6))%32) | v736
	v769 = int32(3)
	goto L192
L204:
	;
	goto L203
L205:
	;
	v786 = v769 + v702
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v786
	v702 = v786
	goto L187
L206:
	;
	v773 = v768 - int32(97)
	if v773 < int32(0) {
		goto L205
	} else {
		goto L207
	}
L207:
	;
	v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v773)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[0]))))
	if int32(base.Ui32(v779)>>(uint(v773&int32(7))%32))&int32(1) != 0 {
		goto L186
	} else {
		goto L208
	}
L208:
	;
	goto L205
L210:
	;
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v801 = v800 + v797
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v801
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v816 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v824 = v801
	goto L213
L211:
	;
	if v920 < int32(0) {
		goto L127
	} else {
		goto L235
	}
L212:
	;
	v920 = v891
	goto L211
L213:
	;
	if v815 <= v824 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v920 = int32(-1)
	goto L211
L216:
	;
	goto L217
L217:
	;
	v831 = int32(1)
	v833 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824+v816))))
	if base.Ui32(v833) < base.Ui32(int32(192)) {
		v890 = v833
		v891 = v831
		goto L218
	} else {
		goto L219
	}
L218:
	;
	if int32(121) < v890 {
		goto L212
	} else {
		goto L231
	}
L219:
	;
	v837 = v824 + int32(1)
	if v837 == v815 {
		v890 = v833
		v891 = v831
		goto L218
	} else {
		goto L220
	}
L220:
	;
	v840 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v837+v816))))
	v842 = v840 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v833) {
		goto L222
	} else {
		goto L223
	}
L221:
	;
	v856 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v846+v816))))
	v858 = v856 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v833) {
		goto L227
	} else {
		goto L228
	}
L222:
	;
	v846 = v824 + int32(2)
	if v846 != v815 {
		goto L221
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v890 = v833<<(uint(int32(6))%32)&int32(1984) | v842
	v891 = int32(2)
	goto L218
L225:
	;
	goto L224
L226:
	;
	v875 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v816+v862))))
	v890 = v875&int32(63) | (v833<<(uint(int32(18))%32)&int32(_a_F_english_UTF_8_stem_11) | v842<<(uint(int32(12))%32) | v858<<(uint(int32(6))%32))
	v891 = int32(4)
	goto L218
L227:
	;
	v862 = v824 + int32(3)
	if v862 != v815 {
		goto L226
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v890 = v833<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v842<<(uint(int32(6))%32) | v858
	v891 = int32(3)
	goto L218
L230:
	;
	goto L229
L231:
	;
	v895 = v890 - int32(97)
	if v895 < int32(0) {
		goto L212
	} else {
		goto L232
	}
L232:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v895)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[0]))))
	if int32(base.Ui32(v901)>>(uint(v895&int32(7))%32))&int32(1) == int32(0) {
		goto L212
	} else {
		goto L233
	}
L233:
	;
	v909 = v891 + v824
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v909
	v824 = v909
	goto L213
L235:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v923 + v920
	goto L127
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v957
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v957 <= v960 {
		goto L244
	} else {
		goto L245
	}
L237:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v932+v928-int32(1)))))
	if base.B2i32(v936 != int32(115))&base.B2i32(v936 != int32(39)) != 0 {
		v957 = v928
		goto L236
	} else {
		goto L238
	}
L238:
	;
	v945 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_15), int32(3), int32(0))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L6
	} else {
		goto L239
	}
L239:
	;
	if v945 == int32(0) {
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v949
	v957 = v949
	goto L236
L241:
	;
	goto L242
L242:
	;
	v951 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v951
	v953 = F_slice_del(m, l0)
	mBase = m.M
	if v953 < int32(0) {
		v2749 = v953
		goto L1
	} else {
		goto L243
	}
L243:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v957 = v956
	goto L236
L244:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1265
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1265
	v1269 = v1265 - int32(1)
	v1270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1269 <= v1270 {
		goto L322
	} else {
		goto L323
	}
L245:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v962+v957-int32(1)))))
	v968 = v966 - int32(100)
	v969 = int32(0)
	if base.B2i32(v968 == v969)|base.B2i32(v968 == int32(15)) == v969 {
		goto L244
	} else {
		goto L246
	}
L246:
	;
	v979 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_16), int32(6), int32(0))
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L6
	} else {
		goto L247
	}
L247:
	;
	if v979 == int32(0) {
		goto L244
	} else {
		goto L248
	}
L248:
	;
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v983
	switch v979 - int32(1) {
	case 0:
		goto L251
	case 1:
		goto L250
	case 2:
		goto L249
	default:
		goto L244
	}
L249:
	;
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L282
L250:
	;
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v994 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v995 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L256
L251:
	;
	v989 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_UTF_8_stem_17))
	mBase = m.M
	v990 = m.ExcPending
	if v990 != 0 {
		goto L6
	} else {
		goto L252
	}
L252:
	;
	if int32(0) <= v989 {
		goto L244
	} else {
		goto L253
	}
L253:
	;
	v2749 = v989
	goto L1
L254:
	;
	if int32(0) <= v1047 {
		goto L273
	} else {
		goto L274
	}
L256:
	;
	goto L257
L257:
	;
	goto L258
L258:
	;
	v1002 = v983
	v1004 = int32(2)
	goto L261
L260:
	;
	v1047 = v1029
	goto L254
L261:
	;
	if v1002 <= v995 {
		goto L263
	} else {
		goto L264
	}
L262:
	;
	goto L260
L263:
	;
	v1047 = int32(-1)
	goto L254
L264:
	;
	goto L265
L265:
	;
	v1009 = v1002 - int32(1)
	v1011 = int32(*(*int8)(unsafe.Add(mBase, uint32(v994+v1009))))
	if base.B2i32(int32(0) <= v1011)|base.B2i32(v1009 <= v995) != 0 {
		v1029 = v1009
		goto L266
	} else {
		goto L267
	}
L266:
	;
	v1033 = int32(1)
	if v1033 < v1004 {
		v1002 = v1029
		v1004 = v1004 - v1033
		goto L261
	} else {
		goto L272
	}
L267:
	;
	v1017 = v1009
	goto L268
L268:
	;
	v1022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v994+v1017))))
	if base.Ui32(int32(191)) < base.Ui32(v1022) {
		v1029 = v1017
		goto L266
	} else {
		goto L270
	}
L269:
	;
	v1029 = v995
	goto L266
L270:
	;
	v1026 = v1017 - int32(1)
	if v995 < v1026 {
		v1017 = v1026
		goto L268
	} else {
		goto L271
	}
L271:
	;
	goto L269
L272:
	;
	goto L262
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1047
	v1053 = F_slice_from_s(m, l0, int32(1), int32(_a_F_english_UTF_8_stem_18))
	mBase = m.M
	v1054 = m.ExcPending
	if v1054 != 0 {
		goto L6
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1057 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1057 + (v983 - v993)
	v1063 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_UTF_8_stem_19))
	mBase = m.M
	v1064 = m.ExcPending
	if v1064 != 0 {
		goto L6
	} else {
		goto L278
	}
L276:
	;
	if int32(0) <= v1053 {
		goto L244
	} else {
		goto L277
	}
L277:
	;
	v2749 = v1053
	goto L1
L278:
	;
	if int32(0) <= v1063 {
		goto L244
	} else {
		goto L279
	}
L279:
	;
	v2749 = v1063
	goto L1
L280:
	;
	if v1120 < int32(0) {
		goto L244
	} else {
		goto L299
	}
L282:
	;
	goto L283
L283:
	;
	goto L284
L284:
	;
	v1075 = v983
	v1077 = int32(1)
	goto L287
L286:
	;
	v1120 = v1102
	goto L280
L287:
	;
	if v1075 <= v1068 {
		goto L289
	} else {
		goto L290
	}
L288:
	;
	goto L286
L289:
	;
	v1120 = int32(-1)
	goto L280
L290:
	;
	goto L291
L291:
	;
	v1082 = v1075 - int32(1)
	v1084 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1067+v1082))))
	if base.B2i32(int32(0) <= v1084)|base.B2i32(v1082 <= v1068) != 0 {
		v1102 = v1082
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v1106 = int32(1)
	if v1106 < v1077 {
		v1075 = v1102
		v1077 = v1077 - v1106
		goto L287
	} else {
		goto L298
	}
L293:
	;
	v1090 = v1082
	goto L294
L294:
	;
	v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067+v1090))))
	if base.Ui32(int32(191)) < base.Ui32(v1095) {
		v1102 = v1090
		goto L292
	} else {
		goto L296
	}
L295:
	;
	v1102 = v1068
	goto L292
L296:
	;
	v1099 = v1090 - int32(1)
	if v1068 < v1099 {
		v1090 = v1099
		goto L294
	} else {
		goto L297
	}
L297:
	;
	goto L295
L298:
	;
	goto L288
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1120
	v1137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1147 = v1120
	goto L302
L300:
	;
	if v1253 < int32(0) {
		goto L244
	} else {
		goto L318
	}
L301:
	;
	v1253 = int32(-1)
	goto L300
L302:
	;
	if v1147 <= v1137 {
		goto L301
	} else {
		goto L304
	}
L304:
	;
	v1154 = int32(1)
	v1155 = v1147 - v1154
	v1157 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1138+v1155))))
	v1159 = v1157 & int32(255)
	if base.B2i32(v1155 == v1137)|base.B2i32(int32(0) <= v1157) != 0 {
		v1217 = v1159
		v1221 = v1154
		goto L305
	} else {
		goto L306
	}
L305:
	;
	if int32(121) < v1217 {
		goto L313
	} else {
		goto L314
	}
L306:
	;
	v1166 = v1159 & int32(63)
	v1168 = v1147 - int32(2)
	v1170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138+v1168))))
	v1172 = v1170 << (uint(int32(6)) % 32)
	if base.B2i32(v1168 != v1137)&base.B2i32(base.Ui32(v1170) < base.Ui32(int32(192))) == int32(0) {
		goto L307
	} else {
		goto L308
	}
L307:
	;
	v1217 = v1172&int32(1984) | v1166
	v1221 = int32(2)
	goto L305
L308:
	;
	goto L309
L309:
	;
	v1185 = v1172&int32(4032) | v1166
	v1187 = v1147 - int32(3)
	v1189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1138+v1187))))
	if base.B2i32(v1187 != v1137)&base.B2i32(base.Ui32(v1189) < base.Ui32(int32(224))) == int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1217 = v1189<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v1185
	v1221 = int32(3)
	goto L305
L311:
	;
	goto L312
L312:
	;
	v1207 = int32(4)
	v1209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1147+v1138-v1207))))
	v1217 = v1189<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_20) | v1209&int32(7)<<(uint(int32(18))%32) | v1185
	v1221 = v1207
	goto L305
L313:
	;
	v1238 = v1147 - v1221
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1238
	v1147 = v1238
	goto L302
L314:
	;
	v1223 = v1217 - int32(97)
	if v1223 < int32(0) {
		goto L313
	} else {
		goto L315
	}
L315:
	;
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1223)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[0]))))
	if int32(base.Ui32(v1229)>>(uint(v1223&int32(7))%32))&int32(1) == int32(0) {
		goto L313
	} else {
		goto L316
	}
L316:
	;
	v1253 = v1221
	goto L300
L318:
	;
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1256 - v1253
	v1259 = F_slice_del(m, l0)
	mBase = m.M
	if v1259 < int32(0) {
		v2749 = v1259
		goto L1
	} else {
		goto L319
	}
L319:
	;
	goto L244
L320:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1949
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1949
	v1952 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1949 <= v1952 {
		goto L457
	} else {
		goto L458
	}
L321:
	;
	v1288 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_21), int32(7), int32(0))
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L6
	} else {
		goto L326
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1265
	goto L320
L323:
	;
	v1272 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1272+v1269))))
	if v1274&int32(224) != int32(96) {
		goto L322
	} else {
		goto L324
	}
L324:
	;
	if int32(1)<<(uint(v1274)%32)&int32(33554576) != 0 {
		goto L321
	} else {
		goto L325
	}
L325:
	;
	goto L322
L326:
	;
	v1290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1290
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	switch v1288 - int32(1) {
	case 0:
		goto L329
	case 1:
		goto L327
	case 2:
		goto L328
	default:
		goto L320
	}
L327:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1503 = v1292 - v1290
	v1504 = v1502 - v1503
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1504
	v1519 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1520 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1529 = v1504
	goto L369
L328:
	;
	v1328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1290 <= v1328 {
		goto L327
	} else {
		goto L339
	}
L329:
	;
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1290 < v1295 {
		goto L320
	} else {
		goto L330
	}
L330:
	;
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1290-int32(2) <= v1297 {
		goto L331
	} else {
		goto L332
	}
L331:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1318 + (v1290 - v1292)
	v1324 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_UTF_8_stem_22))
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L6
	} else {
		goto L337
	}
L332:
	;
	v1301 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1301+v1290-int32(1)))))
	if v1305 != int32(99) {
		goto L331
	} else {
		goto L333
	}
L333:
	;
	v1311 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_23), int32(3), int32(0))
	mBase = m.M
	v1312 = m.ExcPending
	if v1312 != 0 {
		goto L6
	} else {
		goto L334
	}
L334:
	;
	if v1311 == int32(0) {
		goto L331
	} else {
		goto L335
	}
L335:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1315 <= v1316 {
		goto L320
	} else {
		goto L336
	}
L336:
	;
	goto L331
L337:
	;
	if int32(0) <= v1324 {
		goto L320
	} else {
		goto L338
	}
L338:
	;
	v2749 = v1324
	goto L1
L339:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1332 = int32(1)
	v1334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1330+v1290-v1332))))
	if base.B2i32(v1334&int32(224) != int32(96))|base.B2i32(v1332<<(uint(v1334)%32)&int32(34881536) == int32(0)) != 0 {
		goto L327
	} else {
		goto L340
	}
L340:
	;
	v1349 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_24), int32(7), int32(0))
	mBase = m.M
	v1350 = m.ExcPending
	if v1350 != 0 {
		goto L6
	} else {
		goto L343
	}
L341:
	;
	v1497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1497 <= v1498 {
		goto L320
	} else {
		goto L366
	}
L342:
	;
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1366 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1367 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L346
L343:
	;
	switch v1349 {
	case 0:
		goto L327
	case 1:
		goto L342
	case 2:
		goto L341
	default:
		goto L320
	}
L344:
	;
	if v1482 != 0 {
		goto L327
	} else {
		goto L362
	}
L345:
	;
	v1482 = v1475
	goto L344
L346:
	;
	if v1352 <= v1366 {
		v1475 = int32(-1)
		goto L345
	} else {
		goto L348
	}
L347:
	;
	v1475 = int32(0)
	goto L345
L348:
	;
	v1383 = int32(1)
	v1384 = v1352 - v1383
	v1386 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1367+v1384))))
	v1388 = v1386 & int32(255)
	if base.B2i32(v1384 == v1366)|base.B2i32(int32(0) <= v1386) != 0 {
		v1446 = v1388
		v1450 = v1383
		goto L349
	} else {
		goto L350
	}
L349:
	;
	if int32(121) < v1446 {
		goto L357
	} else {
		goto L358
	}
L350:
	;
	v1395 = v1388 & int32(63)
	v1397 = v1352 - int32(2)
	v1399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1367+v1397))))
	v1401 = v1399 << (uint(int32(6)) % 32)
	if base.B2i32(v1397 != v1366)&base.B2i32(base.Ui32(v1399) < base.Ui32(int32(192))) == int32(0) {
		goto L351
	} else {
		goto L352
	}
L351:
	;
	v1446 = v1401&int32(1984) | v1395
	v1450 = int32(2)
	goto L349
L352:
	;
	goto L353
L353:
	;
	v1414 = v1401&int32(4032) | v1395
	v1416 = v1352 - int32(3)
	v1418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1367+v1416))))
	if base.B2i32(v1416 != v1366)&base.B2i32(base.Ui32(v1418) < base.Ui32(int32(224))) == int32(0) {
		goto L354
	} else {
		goto L355
	}
L354:
	;
	v1446 = v1418<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v1414
	v1450 = int32(3)
	goto L349
L355:
	;
	goto L356
L356:
	;
	v1436 = int32(4)
	v1438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1352+v1367-v1436))))
	v1446 = v1418<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_20) | v1438&int32(7)<<(uint(int32(18))%32) | v1414
	v1450 = v1436
	goto L349
L357:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1352 - v1450
	goto L361
L358:
	;
	v1452 = v1446 - int32(97)
	if v1452 < int32(0) {
		goto L357
	} else {
		goto L359
	}
L359:
	;
	v1458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1452)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[0]))))
	if int32(base.Ui32(v1458)>>(uint(v1452&int32(7))%32))&int32(1) == int32(0) {
		goto L357
	} else {
		goto L360
	}
L360:
	;
	v1482 = v1450
	goto L344
L361:
	;
	goto L347
L362:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1484 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1484 < v1483 {
		goto L327
	} else {
		goto L363
	}
L363:
	;
	v1486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1488 = v1486 + (v1352 - v1351)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1488
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1488
	v1493 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_UTF_8_stem_25))
	mBase = m.M
	v1494 = m.ExcPending
	if v1494 != 0 {
		goto L6
	} else {
		goto L364
	}
L364:
	;
	if int32(0) <= v1493 {
		goto L320
	} else {
		goto L365
	}
L365:
	;
	v2749 = v1493
	goto L1
L366:
	;
	goto L327
L367:
	;
	if v1635 < int32(0) {
		goto L320
	} else {
		goto L385
	}
L368:
	;
	v1635 = int32(-1)
	goto L367
L369:
	;
	if v1529 <= v1519 {
		goto L368
	} else {
		goto L371
	}
L371:
	;
	v1536 = int32(1)
	v1537 = v1529 - v1536
	v1539 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1520+v1537))))
	v1541 = v1539 & int32(255)
	if base.B2i32(v1537 == v1519)|base.B2i32(int32(0) <= v1539) != 0 {
		v1599 = v1541
		v1603 = v1536
		goto L372
	} else {
		goto L373
	}
L372:
	;
	if int32(121) < v1599 {
		goto L380
	} else {
		goto L381
	}
L373:
	;
	v1548 = v1541 & int32(63)
	v1550 = v1529 - int32(2)
	v1552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1520+v1550))))
	v1554 = v1552 << (uint(int32(6)) % 32)
	if base.B2i32(v1550 != v1519)&base.B2i32(base.Ui32(v1552) < base.Ui32(int32(192))) == int32(0) {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1599 = v1554&int32(1984) | v1548
	v1603 = int32(2)
	goto L372
L375:
	;
	goto L376
L376:
	;
	v1567 = v1554&int32(4032) | v1548
	v1569 = v1529 - int32(3)
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1520+v1569))))
	if base.B2i32(v1569 != v1519)&base.B2i32(base.Ui32(v1571) < base.Ui32(int32(224))) == int32(0) {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1599 = v1571<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v1567
	v1603 = int32(3)
	goto L372
L378:
	;
	goto L379
L379:
	;
	v1589 = int32(4)
	v1591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1529+v1520-v1589))))
	v1599 = v1571<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_20) | v1591&int32(7)<<(uint(int32(18))%32) | v1567
	v1603 = v1589
	goto L372
L380:
	;
	v1620 = v1529 - v1603
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1620
	v1529 = v1620
	goto L369
L381:
	;
	v1605 = v1599 - int32(97)
	if v1605 < int32(0) {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	v1611 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1605)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[0]))))
	if int32(base.Ui32(v1611)>>(uint(v1605&int32(7))%32))&int32(1) == int32(0) {
		goto L380
	} else {
		goto L383
	}
L383:
	;
	v1635 = v1603
	goto L367
L385:
	;
	v1638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1638 - v1503
	v1641 = F_slice_del(m, l0)
	mBase = m.M
	if v1641 < int32(0) {
		v2749 = v1641
		goto L1
	} else {
		goto L386
	}
L386:
	;
	v1644 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1644
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1644
	v1648 = v1644 - int32(1)
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1648 <= v1649 {
		v1807 = v1644
		goto L390
	} else {
		goto L391
	}
L387:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1881 = v1879 + (v1644 - v1665)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1881
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1881
	v1884 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L438
L388:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1878 = v1877
	goto L387
L389:
	;
	v1873 = F_slice_from_s(m, l0, int32(1), v1871)
	mBase = m.M
	v1874 = m.ExcPending
	if v1874 != 0 {
		goto L6
	} else {
		goto L434
	}
L390:
	;
	v1810 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1807 != v1810 {
		goto L320
	} else {
		goto L421
	}
L391:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1651+v1648))))
	if base.B2i32(v1653&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v1653)%32)&int32(68514004) == int32(0)) != 0 {
		v1807 = v1644
		goto L390
	} else {
		goto L392
	}
L392:
	;
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1670 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_26), int32(13), int32(0))
	mBase = m.M
	v1671 = m.ExcPending
	if v1671 != 0 {
		goto L6
	} else {
		goto L395
	}
L393:
	;
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1807 = v1806
	goto L390
L394:
	;
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L398
L395:
	;
	switch v1670 - int32(1) {
	case 0:
		v1871 = int32(_a_F_english_UTF_8_stem_27)
		goto L389
	case 1:
		goto L394
	case 2:
		goto L393
	default:
		goto L388
	}
L396:
	;
	v1803 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1802 != 0 {
		v1878 = v1803
		goto L387
	} else {
		goto L419
	}
L397:
	;
	v1802 = v1795
	goto L396
L398:
	;
	if v1686 <= v1687 {
		v1795 = int32(-1)
		goto L397
	} else {
		goto L400
	}
L399:
	;
	v1795 = int32(0)
	goto L397
L400:
	;
	v1704 = int32(1)
	v1705 = v1686 - v1704
	v1707 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1688+v1705))))
	v1709 = v1707 & int32(255)
	if base.B2i32(v1705 == v1687)|base.B2i32(int32(0) <= v1707) != 0 {
		v1767 = v1709
		v1771 = v1704
		goto L401
	} else {
		goto L402
	}
L401:
	;
	if int32(111) < v1767 {
		goto L409
	} else {
		goto L410
	}
L402:
	;
	v1716 = v1709 & int32(63)
	v1718 = v1686 - int32(2)
	v1720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1688+v1718))))
	v1722 = v1720 << (uint(int32(6)) % 32)
	if base.B2i32(v1718 != v1687)&base.B2i32(base.Ui32(v1720) < base.Ui32(int32(192))) == int32(0) {
		goto L403
	} else {
		goto L404
	}
L403:
	;
	v1767 = v1722&int32(1984) | v1716
	v1771 = int32(2)
	goto L401
L404:
	;
	goto L405
L405:
	;
	v1735 = v1722&int32(4032) | v1716
	v1737 = v1686 - int32(3)
	v1739 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1688+v1737))))
	if base.B2i32(v1737 != v1687)&base.B2i32(base.Ui32(v1739) < base.Ui32(int32(224))) == int32(0) {
		goto L406
	} else {
		goto L407
	}
L406:
	;
	v1767 = v1739<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v1735
	v1771 = int32(3)
	goto L401
L407:
	;
	goto L408
L408:
	;
	v1757 = int32(4)
	v1759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1686+v1688-v1757))))
	v1767 = v1739<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_20) | v1759&int32(7)<<(uint(int32(18))%32) | v1735
	v1771 = v1757
	goto L401
L409:
	;
	v1802 = v1771
	goto L396
L410:
	;
	goto L411
L411:
	;
	v1773 = v1767 - int32(97)
	if v1773 < int32(0) {
		goto L412
	} else {
		goto L413
	}
L412:
	;
	v1802 = v1771
	goto L396
L413:
	;
	goto L414
L414:
	;
	v1779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1773)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[1]))))
	if int32(base.Ui32(v1779)>>(uint(v1773&int32(7))%32))&int32(1) == int32(0) {
		goto L415
	} else {
		goto L416
	}
L415:
	;
	v1802 = v1771
	goto L396
L416:
	;
	goto L417
L417:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1686 - v1771
	goto L418
L418:
	;
	goto L399
L419:
	;
	v1804 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v1803 < v1804 {
		v1878 = v1803
		goto L387
	} else {
		goto L420
	}
L420:
	;
	goto L320
L421:
	;
	v1812 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1821 = F_out_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_28), int32(89), int32(121), int32(0))
	mBase = m.M
	if v1821 != 0 {
		goto L424
	} else {
		goto L425
	}
L422:
	;
	if v1861 == int32(0) {
		goto L320
	} else {
		goto L433
	}
L423:
	;
	v1861 = int32(1)
	goto L422
L424:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1835 = v1812 - v1815
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1834 - v1835
	v1842 = F_out_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_29), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1842 != 0 {
		goto L428
	} else {
		goto L429
	}
L425:
	;
	v1826 = F_in_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_29), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1826 != 0 {
		goto L424
	} else {
		goto L426
	}
L426:
	;
	v1830 = int32(0)
	v1831 = F_out_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_29), int32(97), int32(121), v1830)
	mBase = m.M
	if v1831 == v1830 {
		goto L423
	} else {
		goto L427
	}
L427:
	;
	goto L424
L428:
	;
	v1851 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1851 - v1835
	v1856 = F_eq_s_b(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_30))
	mBase = m.M
	if v1856 != 0 {
		goto L423
	} else {
		goto L432
	}
L429:
	;
	v1847 = F_in_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_29), int32(97), int32(121), int32(0))
	mBase = m.M
	if v1847 != 0 {
		goto L428
	} else {
		goto L430
	}
L430:
	;
	v1848 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1849 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1848 <= v1849 {
		goto L423
	} else {
		goto L431
	}
L431:
	;
	goto L428
L432:
	;
	v1861 = int32(0)
	goto L422
L433:
	;
	v1864 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1864 + (v1807 - v1812)
	v1871 = int32(_a_F_english_UTF_8_stem_31)
	goto L389
L434:
	;
	if int32(0) <= v1873 {
		goto L320
	} else {
		goto L435
	}
L435:
	;
	v2749 = v1873
	goto L1
L436:
	;
	if v1936 < int32(0) {
		goto L320
	} else {
		goto L455
	}
L438:
	;
	goto L439
L439:
	;
	goto L440
L440:
	;
	v1891 = v1881
	v1893 = int32(1)
	goto L443
L442:
	;
	v1936 = v1918
	goto L436
L443:
	;
	if v1891 <= v1878 {
		goto L445
	} else {
		goto L446
	}
L444:
	;
	goto L442
L445:
	;
	v1936 = int32(-1)
	goto L436
L446:
	;
	goto L447
L447:
	;
	v1898 = v1891 - int32(1)
	v1900 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1884+v1898))))
	if base.B2i32(int32(0) <= v1900)|base.B2i32(v1898 <= v1878) != 0 {
		v1918 = v1898
		goto L448
	} else {
		goto L449
	}
L448:
	;
	v1922 = int32(1)
	if v1922 < v1893 {
		v1891 = v1918
		v1893 = v1893 - v1922
		goto L443
	} else {
		goto L454
	}
L449:
	;
	v1906 = v1898
	goto L450
L450:
	;
	v1911 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1884+v1906))))
	if base.Ui32(int32(191)) < base.Ui32(v1911) {
		v1918 = v1906
		goto L448
	} else {
		goto L452
	}
L451:
	;
	v1918 = v1878
	goto L448
L452:
	;
	v1915 = v1906 - int32(1)
	if v1878 < v1915 {
		v1906 = v1915
		goto L450
	} else {
		goto L453
	}
L453:
	;
	goto L451
L454:
	;
	goto L444
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1936
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1936
	v1941 = F_slice_del(m, l0)
	mBase = m.M
	if v1941 < int32(0) {
		v2749 = v1941
		goto L1
	} else {
		goto L456
	}
L456:
	;
	goto L320
L457:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2107
	v2109 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2107
	v2113 = v2107 - int32(1)
	v2114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2113 <= v2114 {
		v2379 = v2109
		goto L482
	} else {
		goto L483
	}
L458:
	;
	v1954 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1954+v1949-int32(1)))))
	if v1958|int32(32) != int32(121) {
		goto L457
	} else {
		goto L459
	}
L459:
	;
	v1964 = v1949 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1964
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1964
	v1980 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L462
L460:
	;
	if v2096 != 0 {
		goto L457
	} else {
		goto L478
	}
L461:
	;
	v2096 = v2089
	goto L460
L462:
	;
	if v1964 <= v1980 {
		v2089 = int32(-1)
		goto L461
	} else {
		goto L464
	}
L463:
	;
	v2089 = int32(0)
	goto L461
L464:
	;
	v1997 = int32(1)
	v1998 = v1964 - v1997
	v2000 = int32(*(*int8)(unsafe.Add(mBase, uint32(v1981+v1998))))
	v2002 = v2000 & int32(255)
	if base.B2i32(v1998 == v1980)|base.B2i32(int32(0) <= v2000) != 0 {
		v2060 = v2002
		v2064 = v1997
		goto L465
	} else {
		goto L466
	}
L465:
	;
	if int32(121) < v2060 {
		goto L473
	} else {
		goto L474
	}
L466:
	;
	v2009 = v2002 & int32(63)
	v2011 = v1964 - int32(2)
	v2013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981+v2011))))
	v2015 = v2013 << (uint(int32(6)) % 32)
	if base.B2i32(v2011 != v1980)&base.B2i32(base.Ui32(v2013) < base.Ui32(int32(192))) == int32(0) {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v2060 = v2015&int32(1984) | v2009
	v2064 = int32(2)
	goto L465
L468:
	;
	goto L469
L469:
	;
	v2028 = v2015&int32(4032) | v2009
	v2030 = v1964 - int32(3)
	v2032 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1981+v2030))))
	if base.B2i32(v2030 != v1980)&base.B2i32(base.Ui32(v2032) < base.Ui32(int32(224))) == int32(0) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v2060 = v2032<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v2028
	v2064 = int32(3)
	goto L465
L471:
	;
	goto L472
L472:
	;
	v2050 = int32(4)
	v2052 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1964+v1981-v2050))))
	v2060 = v2032<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_20) | v2052&int32(7)<<(uint(int32(18))%32) | v2028
	v2064 = v2050
	goto L465
L473:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1964 - v2064
	goto L477
L474:
	;
	v2066 = v2060 - int32(97)
	if v2066 < int32(0) {
		goto L473
	} else {
		goto L475
	}
L475:
	;
	v2072 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2066)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[0]))))
	if int32(base.Ui32(v2072)>>(uint(v2066&int32(7))%32))&int32(1) == int32(0) {
		goto L473
	} else {
		goto L476
	}
L476:
	;
	v2096 = v2064
	goto L460
L477:
	;
	goto L463
L478:
	;
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2098 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2097 <= v2098 {
		goto L457
	} else {
		goto L479
	}
L479:
	;
	v2102 = F_slice_from_s(m, l0, int32(1), int32(_a_F_english_UTF_8_stem_32))
	mBase = m.M
	v2103 = m.ExcPending
	if v2103 != 0 {
		goto L6
	} else {
		goto L480
	}
L480:
	;
	if v2102 < int32(0) {
		v2749 = v2102
		goto L1
	} else {
		goto L481
	}
L481:
	;
	goto L457
L482:
	;
	if v2379 < int32(0) {
		v2749 = v2379
		goto L1
	} else {
		goto L562
	}
L483:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2116+v2113))))
	if base.B2i32(v2118&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2118)%32)&int32(_a_F_english_UTF_8_stem_33) == int32(0)) != 0 {
		v2379 = v2109
		goto L482
	} else {
		goto L484
	}
L484:
	;
	v2133 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_34), int32(25), int32(0))
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L6
	} else {
		goto L485
	}
L485:
	;
	if v2133 == int32(0) {
		v2379 = v2109
		goto L482
	} else {
		goto L486
	}
L486:
	;
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2137
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2137 < v2139 {
		v2379 = v2109
		goto L482
	} else {
		goto L487
	}
L487:
	;
	switch v2133 - int32(1) {
	case 0:
		goto L504
	case 1:
		goto L503
	case 2:
		goto L502
	case 3:
		goto L501
	case 4:
		goto L500
	case 5:
		goto L499
	case 6:
		goto L498
	case 7:
		goto L497
	case 8:
		goto L496
	case 9:
		goto L495
	case 10:
		goto L494
	case 11:
		goto L493
	case 12:
		goto L492
	case 13:
		goto L491
	case 14:
		goto L490
	case 15:
		goto L489
	default:
		goto L488
	}
L488:
	;
	v2379 = int32(1)
	goto L482
L489:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	goto L539
L490:
	;
	v2241 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_35))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L6
	} else {
		goto L535
	}
L491:
	;
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2137 <= v2221 {
		v2379 = v2109
		goto L482
	} else {
		goto L531
	}
L492:
	;
	v2217 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_UTF_8_stem_36))
	mBase = m.M
	v2218 = m.ExcPending
	if v2218 != 0 {
		goto L6
	} else {
		goto L529
	}
L493:
	;
	v2211 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_37))
	mBase = m.M
	v2212 = m.ExcPending
	if v2212 != 0 {
		goto L6
	} else {
		goto L527
	}
L494:
	;
	v2205 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_38))
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L6
	} else {
		goto L525
	}
L495:
	;
	v2199 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_39))
	mBase = m.M
	v2200 = m.ExcPending
	if v2200 != 0 {
		goto L6
	} else {
		goto L523
	}
L496:
	;
	v2193 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_40))
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L6
	} else {
		goto L521
	}
L497:
	;
	v2187 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_UTF_8_stem_41))
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L6
	} else {
		goto L519
	}
L498:
	;
	v2181 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_42))
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L6
	} else {
		goto L517
	}
L499:
	;
	v2175 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_43))
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L6
	} else {
		goto L515
	}
L500:
	;
	v2169 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_44))
	mBase = m.M
	v2170 = m.ExcPending
	if v2170 != 0 {
		goto L6
	} else {
		goto L513
	}
L501:
	;
	v2163 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_45))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L6
	} else {
		goto L511
	}
L502:
	;
	v2157 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_46))
	mBase = m.M
	v2158 = m.ExcPending
	if v2158 != 0 {
		goto L6
	} else {
		goto L509
	}
L503:
	;
	v2151 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_47))
	mBase = m.M
	v2152 = m.ExcPending
	if v2152 != 0 {
		goto L6
	} else {
		goto L507
	}
L504:
	;
	v2145 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_48))
	mBase = m.M
	v2146 = m.ExcPending
	if v2146 != 0 {
		goto L6
	} else {
		goto L505
	}
L505:
	;
	if int32(0) <= v2145 {
		goto L488
	} else {
		goto L506
	}
L506:
	;
	v2379 = v2145
	goto L482
L507:
	;
	if int32(0) <= v2151 {
		goto L488
	} else {
		goto L508
	}
L508:
	;
	v2379 = v2151
	goto L482
L509:
	;
	if int32(0) <= v2157 {
		goto L488
	} else {
		goto L510
	}
L510:
	;
	v2379 = v2157
	goto L482
L511:
	;
	if int32(0) <= v2163 {
		goto L488
	} else {
		goto L512
	}
L512:
	;
	v2379 = v2163
	goto L482
L513:
	;
	if int32(0) <= v2169 {
		goto L488
	} else {
		goto L514
	}
L514:
	;
	v2379 = v2169
	goto L482
L515:
	;
	if int32(0) <= v2175 {
		goto L488
	} else {
		goto L516
	}
L516:
	;
	v2379 = v2175
	goto L482
L517:
	;
	if int32(0) <= v2181 {
		goto L488
	} else {
		goto L518
	}
L518:
	;
	v2379 = v2181
	goto L482
L519:
	;
	if int32(0) <= v2187 {
		goto L488
	} else {
		goto L520
	}
L520:
	;
	v2379 = v2187
	goto L482
L521:
	;
	if int32(0) <= v2193 {
		goto L488
	} else {
		goto L522
	}
L522:
	;
	v2379 = v2193
	goto L482
L523:
	;
	if int32(0) <= v2199 {
		goto L488
	} else {
		goto L524
	}
L524:
	;
	v2379 = v2199
	goto L482
L525:
	;
	if int32(0) <= v2205 {
		goto L488
	} else {
		goto L526
	}
L526:
	;
	v2379 = v2205
	goto L482
L527:
	;
	if int32(0) <= v2211 {
		goto L488
	} else {
		goto L528
	}
L528:
	;
	v2379 = v2211
	goto L482
L529:
	;
	if int32(0) <= v2217 {
		goto L488
	} else {
		goto L530
	}
L530:
	;
	v2379 = v2217
	goto L482
L531:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2223+v2137-int32(1)))))
	if v2227 != int32(108) {
		v2379 = v2109
		goto L482
	} else {
		goto L532
	}
L532:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2137 - int32(1)
	v2235 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_UTF_8_stem_49))
	mBase = m.M
	v2236 = m.ExcPending
	if v2236 != 0 {
		goto L6
	} else {
		goto L533
	}
L533:
	;
	if int32(0) <= v2235 {
		goto L488
	} else {
		goto L534
	}
L534:
	;
	v2379 = v2235
	goto L482
L535:
	;
	if int32(0) <= v2241 {
		goto L488
	} else {
		goto L536
	}
L536:
	;
	v2379 = v2241
	goto L482
L537:
	;
	if v2373 != 0 {
		v2379 = v2109
		goto L482
	} else {
		goto L560
	}
L538:
	;
	v2373 = v2366
	goto L537
L539:
	;
	if v2257 <= v2258 {
		v2366 = int32(-1)
		goto L538
	} else {
		goto L541
	}
L540:
	;
	v2366 = int32(0)
	goto L538
L541:
	;
	v2275 = int32(1)
	v2276 = v2257 - v2275
	v2278 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2259+v2276))))
	v2280 = v2278 & int32(255)
	if base.B2i32(v2276 == v2258)|base.B2i32(int32(0) <= v2278) != 0 {
		v2338 = v2280
		v2342 = v2275
		goto L542
	} else {
		goto L543
	}
L542:
	;
	if int32(116) < v2338 {
		goto L550
	} else {
		goto L551
	}
L543:
	;
	v2287 = v2280 & int32(63)
	v2289 = v2257 - int32(2)
	v2291 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2259+v2289))))
	v2293 = v2291 << (uint(int32(6)) % 32)
	if base.B2i32(v2289 != v2258)&base.B2i32(base.Ui32(v2291) < base.Ui32(int32(192))) == int32(0) {
		goto L544
	} else {
		goto L545
	}
L544:
	;
	v2338 = v2293&int32(1984) | v2287
	v2342 = int32(2)
	goto L542
L545:
	;
	goto L546
L546:
	;
	v2306 = v2293&int32(4032) | v2287
	v2308 = v2257 - int32(3)
	v2310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2259+v2308))))
	if base.B2i32(v2308 != v2258)&base.B2i32(base.Ui32(v2310) < base.Ui32(int32(224))) == int32(0) {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v2338 = v2310<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_12) | v2306
	v2342 = int32(3)
	goto L542
L548:
	;
	goto L549
L549:
	;
	v2328 = int32(4)
	v2330 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2257+v2259-v2328))))
	v2338 = v2310<<(uint(int32(12))%32)&int32(_a_F_english_UTF_8_stem_20) | v2330&int32(7)<<(uint(int32(18))%32) | v2306
	v2342 = v2328
	goto L542
L550:
	;
	v2373 = v2342
	goto L537
L551:
	;
	goto L552
L552:
	;
	v2344 = v2338 - int32(99)
	if v2344 < int32(0) {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	v2373 = v2342
	goto L537
L554:
	;
	goto L555
L555:
	;
	v2350 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v2344)>>(uint(int32(3))%32)))+uint32(_c_F_english_UTF_8_stem[2]))))
	if int32(base.Ui32(v2350)>>(uint(v2344&int32(7))%32))&int32(1) == int32(0) {
		goto L556
	} else {
		goto L557
	}
L556:
	;
	v2373 = v2342
	goto L537
L557:
	;
	goto L558
L558:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2257 - v2342
	goto L559
L559:
	;
	goto L540
L560:
	;
	v2374 = F_slice_del(m, l0)
	mBase = m.M
	if v2374 < int32(0) {
		v2379 = v2374
		goto L482
	} else {
		goto L561
	}
L561:
	;
	goto L488
L562:
	;
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2384
	v2386 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2384
	v2389 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2384-int32(2) <= v2389 {
		v2456 = v2386
		goto L563
	} else {
		goto L564
	}
L563:
	;
	if v2456 < int32(0) {
		v2749 = v2456
		goto L1
	} else {
		goto L587
	}
L564:
	;
	v2393 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2395 = int32(1)
	v2397 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2393+v2384-v2395))))
	if base.B2i32(v2397&int32(224) != int32(96))|base.B2i32(v2395<<(uint(v2397)%32)&int32(_a_F_english_UTF_8_stem_50) == int32(0)) != 0 {
		v2456 = v2386
		goto L563
	} else {
		goto L565
	}
L565:
	;
	v2412 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_51), int32(9), int32(0))
	mBase = m.M
	v2413 = m.ExcPending
	if v2413 != 0 {
		goto L6
	} else {
		goto L566
	}
L566:
	;
	if v2412 == int32(0) {
		v2456 = v2386
		goto L563
	} else {
		goto L567
	}
L567:
	;
	v2416 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2416
	v2418 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2416 < v2418 {
		v2456 = v2386
		goto L563
	} else {
		goto L568
	}
L568:
	;
	switch v2412 - int32(1) {
	case 0:
		goto L575
	case 1:
		goto L574
	case 2:
		goto L573
	case 3:
		goto L572
	case 4:
		goto L571
	case 5:
		goto L570
	default:
		goto L569
	}
L569:
	;
	v2456 = int32(1)
	goto L563
L570:
	;
	v2449 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2416 < v2449 {
		v2456 = v2386
		goto L563
	} else {
		goto L585
	}
L571:
	;
	v2446 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2446 {
		goto L569
	} else {
		goto L584
	}
L572:
	;
	v2442 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_UTF_8_stem_52))
	mBase = m.M
	v2443 = m.ExcPending
	if v2443 != 0 {
		goto L6
	} else {
		goto L582
	}
L573:
	;
	v2436 = F_slice_from_s(m, l0, int32(2), int32(_a_F_english_UTF_8_stem_53))
	mBase = m.M
	v2437 = m.ExcPending
	if v2437 != 0 {
		goto L6
	} else {
		goto L580
	}
L574:
	;
	v2430 = F_slice_from_s(m, l0, int32(3), int32(_a_F_english_UTF_8_stem_54))
	mBase = m.M
	v2431 = m.ExcPending
	if v2431 != 0 {
		goto L6
	} else {
		goto L578
	}
L575:
	;
	v2424 = F_slice_from_s(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_55))
	mBase = m.M
	v2425 = m.ExcPending
	if v2425 != 0 {
		goto L6
	} else {
		goto L576
	}
L576:
	;
	if int32(0) <= v2424 {
		goto L569
	} else {
		goto L577
	}
L577:
	;
	v2456 = v2424
	goto L563
L578:
	;
	if int32(0) <= v2430 {
		goto L569
	} else {
		goto L579
	}
L579:
	;
	v2456 = v2430
	goto L563
L580:
	;
	if int32(0) <= v2436 {
		goto L569
	} else {
		goto L581
	}
L581:
	;
	v2456 = v2436
	goto L563
L582:
	;
	if int32(0) <= v2442 {
		goto L569
	} else {
		goto L583
	}
L583:
	;
	v2456 = v2442
	goto L563
L584:
	;
	v2456 = v2446
	goto L563
L585:
	;
	v2451 = F_slice_del(m, l0)
	mBase = m.M
	if v2451 < int32(0) {
		v2456 = v2451
		goto L563
	} else {
		goto L586
	}
L586:
	;
	goto L569
L587:
	;
	v2461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2461
	v2463 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2461
	v2467 = v2461 - int32(1)
	v2468 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2467 <= v2468 {
		v2521 = v2463
		goto L588
	} else {
		goto L589
	}
L588:
	;
	if v2521 < int32(0) {
		v2749 = v2521
		goto L1
	} else {
		goto L601
	}
L589:
	;
	v2470 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2470+v2467))))
	if base.B2i32(v2472&int32(224) != int32(96))|base.B2i32(int32(1)<<(uint(v2472)%32)&int32(_a_F_english_UTF_8_stem_56) == int32(0)) != 0 {
		v2521 = v2463
		goto L588
	} else {
		goto L590
	}
L590:
	;
	v2487 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_57), int32(18), int32(0))
	mBase = m.M
	v2488 = m.ExcPending
	if v2488 != 0 {
		goto L6
	} else {
		goto L591
	}
L591:
	;
	if v2487 == int32(0) {
		v2521 = v2463
		goto L588
	} else {
		goto L592
	}
L592:
	;
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2491
	v2493 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2491 < v2493 {
		v2521 = v2463
		goto L588
	} else {
		goto L593
	}
L593:
	;
	switch v2487 - int32(1) {
	case 0:
		goto L596
	case 1:
		goto L595
	default:
		goto L594
	}
L594:
	;
	v2521 = int32(1)
	goto L588
L595:
	;
	v2500 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2491 <= v2500 {
		v2521 = v2463
		goto L588
	} else {
		goto L598
	}
L596:
	;
	v2497 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2497 {
		goto L594
	} else {
		goto L597
	}
L597:
	;
	v2521 = v2497
	goto L588
L598:
	;
	v2502 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2504 = int32(1)
	v2506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2502+v2491-v2504))))
	if base.Ui32(v2504) < base.Ui32((v2506-int32(115))&int32(255)) {
		v2521 = v2463
		goto L588
	} else {
		goto L599
	}
L599:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2491 - int32(1)
	v2516 = F_slice_del(m, l0)
	mBase = m.M
	if v2516 < int32(0) {
		v2521 = v2516
		goto L588
	} else {
		goto L600
	}
L600:
	;
	goto L594
L601:
	;
	v2526 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2526
	v2528 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2526
	v2531 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2526 <= v2531 {
		v2634 = v2528
		goto L602
	} else {
		goto L603
	}
L602:
	;
	if v2634 < int32(0) {
		v2749 = v2634
		goto L1
	} else {
		goto L631
	}
L603:
	;
	v2533 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2537 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2533+v2526-int32(1)))))
	switch v2537 - int32(101) {
	case 0, 7:
		goto L604
	default:
		v2634 = v2528
		goto L602
	}
L604:
	;
	v2543 = F_find_among_b(m, l0, int32(_a_F_english_UTF_8_stem_58), int32(2), int32(0))
	mBase = m.M
	v2544 = m.ExcPending
	if v2544 != 0 {
		goto L6
	} else {
		goto L605
	}
L605:
	;
	if v2543 == int32(0) {
		v2634 = v2528
		goto L602
	} else {
		goto L606
	}
L606:
	;
	v2547 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2547
	switch v2543 - int32(1) {
	case 0:
		goto L609
	case 1:
		goto L608
	default:
		goto L607
	}
L607:
	;
	v2634 = int32(1)
	goto L602
L608:
	;
	v2613 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2547 < v2613 {
		v2634 = v2528
		goto L602
	} else {
		goto L627
	}
L609:
	;
	v2551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v2547 < v2551 {
		goto L610
	} else {
		goto L611
	}
L610:
	;
	v2553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2547 < v2553 {
		v2634 = v2528
		goto L602
	} else {
		goto L613
	}
L611:
	;
	goto L612
L612:
	;
	v2610 = F_slice_del(m, l0)
	mBase = m.M
	if int32(0) <= v2610 {
		goto L607
	} else {
		goto L626
	}
L613:
	;
	v2555 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2558 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2564 = F_out_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_28), int32(89), int32(121), int32(0))
	mBase = m.M
	if v2564 != 0 {
		goto L616
	} else {
		goto L617
	}
L614:
	;
	if v2604 != 0 {
		v2634 = v2528
		goto L602
	} else {
		goto L625
	}
L615:
	;
	v2604 = int32(1)
	goto L614
L616:
	;
	v2577 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v2578 = v2555 - v2558
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2577 - v2578
	v2585 = F_out_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_29), int32(97), int32(121), int32(0))
	mBase = m.M
	if v2585 != 0 {
		goto L620
	} else {
		goto L621
	}
L617:
	;
	v2569 = F_in_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_29), int32(97), int32(121), int32(0))
	mBase = m.M
	if v2569 != 0 {
		goto L616
	} else {
		goto L618
	}
L618:
	;
	v2573 = int32(0)
	v2574 = F_out_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_29), int32(97), int32(121), v2573)
	mBase = m.M
	if v2574 == v2573 {
		goto L615
	} else {
		goto L619
	}
L619:
	;
	goto L616
L620:
	;
	v2594 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2594 - v2578
	v2599 = F_eq_s_b(m, l0, int32(4), int32(_a_F_english_UTF_8_stem_30))
	mBase = m.M
	if v2599 != 0 {
		goto L615
	} else {
		goto L624
	}
L621:
	;
	v2590 = F_in_grouping_b_U(m, l0, int32(_a_F_english_UTF_8_stem_29), int32(97), int32(121), int32(0))
	mBase = m.M
	if v2590 != 0 {
		goto L620
	} else {
		goto L622
	}
L622:
	;
	v2591 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2592 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2591 <= v2592 {
		goto L615
	} else {
		goto L623
	}
L623:
	;
	goto L620
L624:
	;
	v2604 = int32(0)
	goto L614
L625:
	;
	v2605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2605 + (v2547 - v2555)
	goto L612
L626:
	;
	v2634 = v2610
	goto L602
L627:
	;
	v2615 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v2547 <= v2615 {
		v2634 = v2528
		goto L602
	} else {
		goto L628
	}
L628:
	;
	v2617 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2621 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2617+v2547-int32(1)))))
	if v2621 != int32(108) {
		v2634 = v2528
		goto L602
	} else {
		goto L629
	}
L629:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2547 - int32(1)
	v2627 = F_slice_del(m, l0)
	mBase = m.M
	if v2627 < int32(0) {
		v2634 = v2627
		goto L602
	} else {
		goto L630
	}
L630:
	;
	goto L607
L631:
	;
	v2638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2638
	v2640 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+36)))
	if v2640 != 0 {
		goto L632
	} else {
		goto L633
	}
L632:
	;
	goto L636
L633:
	;
	v2737 = int32(0)
	goto L634
L634:
	;
	if v2737 < int32(0) {
		v2749 = v2737
		goto L1
	} else {
		goto L669
	}
L635:
	;
	v2737 = v2729
	goto L634
L636:
	;
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v2649 = v2647
	goto L639
L637:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2647
	v2729 = int32(1)
	goto L635
L638:
	;
	goto L637
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v2649
	v2655 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v2656 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v2656 != v2649 {
		goto L642
	} else {
		goto L643
	}
L640:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2649
	v2718 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v2649 + v2718
	v2723 = F_slice_from_s(m, l0, v2718, int32(_a_F_english_UTF_8_stem_59))
	mBase = m.M
	v2724 = m.ExcPending
	if v2724 != 0 {
		goto L6
	} else {
		goto L667
	}
L641:
	;
	goto L640
L642:
	;
	v2659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2649+v2655))))
	if v2659 == int32(89) {
		goto L641
	} else {
		goto L645
	}
L643:
	;
	goto L644
L644:
	;
	goto L648
L645:
	;
	goto L644
L646:
	;
	if v2713 < int32(0) {
		goto L638
	} else {
		goto L666
	}
L648:
	;
	goto L649
L649:
	;
	goto L650
L650:
	;
	v2668 = v2649
	v2670 = int32(1)
	goto L653
L652:
	;
	v2713 = v2698
	goto L646
L653:
	;
	if v2656 <= v2668 {
		goto L655
	} else {
		goto L656
	}
L654:
	;
	goto L652
L655:
	;
	v2713 = int32(-1)
	goto L646
L656:
	;
	goto L657
L657:
	;
	v2675 = v2668 + int32(1)
	v2677 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2655+v2668))))
	if base.Ui32(v2677) < base.Ui32(int32(192)) {
		v2698 = v2675
		goto L658
	} else {
		goto L659
	}
L658:
	;
	v2699 = int32(1)
	if v2699 < v2670 {
		v2668 = v2698
		v2670 = v2670 - v2699
		goto L653
	} else {
		goto L665
	}
L659:
	;
	if v2656 <= v2675 {
		v2698 = v2675
		goto L658
	} else {
		goto L660
	}
L660:
	;
	v2684 = v2675
	goto L661
L661:
	;
	v2687 = int32(*(*int8)(unsafe.Add(mBase, uint32(v2655+v2684))))
	if int32(-65) < v2687 {
		v2698 = v2684
		goto L658
	} else {
		goto L663
	}
L662:
	;
	v2698 = v2656
	goto L658
L663:
	;
	v2691 = v2684 + int32(1)
	if v2691 != v2656 {
		v2684 = v2691
		goto L661
	} else {
		goto L664
	}
L664:
	;
	goto L662
L665:
	;
	goto L654
L666:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2713
	v2649 = v2713
	goto L639
L667:
	;
	if int32(0) <= v2723 {
		goto L636
	} else {
		goto L668
	}
L668:
	;
	v2729 = v2723
	goto L635
L669:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v2638
	goto L2
}
func F_eq_s_b(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
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
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	v4 = int32(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6-v7 < l1 {
		v78 = v4
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v78
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = v10 + v6 - l1
	if base.Ui32(int32(4)) <= base.Ui32(l1) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v74 != 0 {
		v78 = v4
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v74 = int32(0)
	goto L3
L5:
	;
	v48 = v43
	v49 = v44
	v50 = v45
	goto L15
L6:
	;
	if (v12|l2)&int32(3) != 0 {
		v43 = v12
		v44 = l2
		v45 = l1
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v36 = v12
	v37 = l2
	v38 = l1
	goto L8
L8:
	;
	if v38 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v20 = v12
	v21 = l2
	v22 = l1
	goto L10
L10:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v25 != v26 {
		v43 = v20
		v44 = v21
		v45 = v22
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v36 = v31
	v37 = v29
	v38 = v33
	goto L8
L12:
	;
	v28 = int32(4)
	v29 = v21 + v28
	v31 = v20 + v28
	v33 = v22 - v28
	if base.Ui32(int32(3)) < base.Ui32(v33) {
		v20 = v31
		v21 = v29
		v22 = v33
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v43 = v36
	v44 = v37
	v45 = v38
	goto L5
L15:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
	v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
	if v53 == v54 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v74 = v53 - v54
	goto L3
L17:
	;
	v56 = int32(1)
	v61 = v50 - v56
	if v61 != 0 {
		v48 = v48 + v56
		v49 = v49 + v56
		v50 = v61
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v6 - l1
	v78 = int32(1)
	goto L1
}
func F_eq_v(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
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
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	v3 = int32(0)
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1-int32(4))))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v9-v10 < v8 {
		v80 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v80
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v14 = v13 + v10
	if base.Ui32(int32(4)) <= base.Ui32(v8) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	if v76 != 0 {
		v80 = v3
		goto L1
	} else {
		goto L21
	}
L4:
	;
	v76 = int32(0)
	goto L3
L5:
	;
	v50 = v45
	v51 = v46
	v52 = v47
	goto L15
L6:
	;
	if (v14|l1)&int32(3) != 0 {
		v45 = v14
		v46 = l1
		v47 = v8
		goto L5
	} else {
		goto L9
	}
L7:
	;
	v38 = v14
	v39 = l1
	v40 = v8
	goto L8
L8:
	;
	if v40 == int32(0) {
		goto L4
	} else {
		goto L14
	}
L9:
	;
	v22 = v14
	v23 = l1
	v24 = v8
	goto L10
L10:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	if v27 != v28 {
		v45 = v22
		v46 = v23
		v47 = v24
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v38 = v33
	v39 = v31
	v40 = v35
	goto L8
L12:
	;
	v30 = int32(4)
	v31 = v23 + v30
	v33 = v22 + v30
	v35 = v24 - v30
	if base.Ui32(int32(3)) < base.Ui32(v35) {
		v22 = v33
		v23 = v31
		v24 = v35
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v45 = v38
	v46 = v39
	v47 = v40
	goto L5
L15:
	;
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v50))))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
	if v55 == v56 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v76 = v55 - v56
	goto L3
L17:
	;
	v58 = int32(1)
	v63 = v52 - v58
	if v63 != 0 {
		v50 = v50 + v58
		v51 = v51 + v58
		v52 = v63
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
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v8 + v10
	v80 = int32(1)
	goto L1
}
func F_eqjoinsel_semi(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 float64, l7 float64, l8 int32, l9 int32, l10 int32, l11 int32, l12 int32, l13 int32, l14 int32, l15 int32, l16 int32, l17 int32, l18 int32) float64 {
	mBase := m.M
	_ = mBase
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 float64
	_ = v32
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v40 float64
	_ = v40
	var v41 int32
	_ = v41
	var v42 float64
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v53 float32
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v86 int32
	_ = v86
	var v95 float64
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v120 float32
	_ = v120
	var v123 float64
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v134 float32
	_ = v134
	var v137 float64
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v152 float64
	_ = v152
	var v157 int32
	_ = v157
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v177 float32
	_ = v177
	var v187 float64
	_ = v187
	var v204 float64
	_ = v204
	var v236 float64
	_ = v236
	var v241 int32
	_ = v241
	var v242 float64
	_ = v242
	var v243 float64
	_ = v243
	var v244 float64
	_ = v244
	var v252 float64
	_ = v252
	var v253 float64
	_ = v253
	var v256 float64
	_ = v256
	var v264 float64
	_ = v264
	var v267 float32
	_ = v267
	var v270 float64
	_ = v270
	var v275 float64
	_ = v275
	var v310 float64
	_ = v310
	v25 = m.G0
	v27 = v25 - int32(16)
	m.G0 = v27
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	if v29 == int32(0) {
		v37 = l7
		v38 = l9
	} else {
		v32 = *(*float64)(unsafe.Add(mBase, uint32(v29)+16))
		if base.F64_ge(l7, v32) == int32(0) {
			v37 = l7
			v38 = l9
		} else {
			v37 = v32
			v38 = int32(0)
		}
	}
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l18)+16))
	v41 = base.F64_ge(v37, v40)
	if v41 != 0 {
		v42 = v40
	} else {
		v42 = v37
	}
	v43 = int32(0)
	v45 = v38 & base.B2i32(v41 == v43)
	if base.B2i32(l13 == v43)|base.B2i32(l14 == v43) == v43 {
		v53 = *(*float32)(unsafe.Add(mBase, uint32(l12)+8))
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l11)+16))
		v55 = base.F64_convert_i32_s(v54)
		if base.F64_lt(v55, v42) != 0 {
			v57 = v55
		} else {
			v57 = v42
		}
		v58 = base.I32_trunc_sat_f64_s(v57)
		if v54 != v58 {
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
			if v60 != 0 {
				base.MemoryFill(m, l15, int32(0), v60)
			} else {
			}
			if v58 != 0 {
				base.MemoryFill(m, l16, int32(0), v58)
			} else {
			}
			v65 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
			F_eqjoinsel_find_matches(m, l0, l1, l2, l3, l4, l10, l11, v65, v58, l15, l16, l17, v27+int32(8))
			mBase = m.M
			v71 = m.ExcPending
			if v71 != 0 {
				return float64(0)
			} else {
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
				if v74 <= int32(0) {
					v236 = float64(0)
				} else {
					if v74 == int32(1) {
						v152 = float64(0)
						v157 = int32(0)
						v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+l15))))
						if v170 != int32(1) {
							v187 = v152
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
							v177 = *(*float32)(unsafe.Add(mBase, uint32(v173+v157<<(uint(int32(2))%32))))
							v187 = base.F64_add(v152, base.F64_promote_f32(v177))
						}
					} else {
						v86 = int32(0)
						v95 = float64(0)
						v97 = v86
						v100 = v86
						for {
							v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+l15))))
							if v113 == int32(1) {
								v116 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
								v120 = *(*float32)(unsafe.Add(mBase, uint32(v116+v100<<(uint(int32(2))%32))))
								v123 = base.F64_add(v95, base.F64_promote_f32(v120))
							} else {
								v123 = v95
							}
							v124 = int32(1)
							v125 = v100 | v124
							v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l15+v125))))
							if v127 == v124 {
								v130 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
								v134 = *(*float32)(unsafe.Add(mBase, uint32(v130+v125<<(uint(int32(2))%32))))
								v137 = base.F64_add(v123, base.F64_promote_f32(v134))
							} else {
								v137 = v123
							}
							v138 = int32(2)
							v139 = v100 + v138
							v141 = v97 + v138
							if v141 != v74&int32(2147483646) {
								v95 = v137
								v97 = v141
								v100 = v139
								continue
							} else {
								break
							}
							break
						}
						if v74&int32(1) == int32(0) {
							v187 = v137
						} else {
							v152 = v137
							v157 = v139
							v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+l15))))
							if v170 != int32(1) {
								v187 = v152
							} else {
								v173 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
								v177 = *(*float32)(unsafe.Add(mBase, uint32(v173+v157<<(uint(int32(2))%32))))
								v187 = base.F64_add(v152, base.F64_promote_f32(v177))
							}
						}
					}
					v204 = float64(0)
					if base.F64_lt(v187, v204) != 0 {
						v236 = v204
					} else {
						if base.F64_gt(v187, float64(1)) == int32(0) {
							v236 = v187
						} else {
							v236 = float64(1)
						}
					}
				}
				if l8|v45 != 0 {
					v252 = float64(0.5)
				} else {
					v241 = *(*int32)(unsafe.Add(mBase, uint32(l17)))
					v242 = base.F64_convert_i32_s(v241)
					v243 = base.F64_sub(l6, v242)
					v244 = base.F64_sub(v42, v242)
					if base.F64_le(v243, v244)|base.F64_lt(v244, float64(0)) != 0 {
						v252 = float64(1)
					} else {
						v252 = base.F64_div(v244, v243)
					}
				}
				v253 = float64(0)
				v256 = base.F64_sub(base.F64_sub(float64(1), v236), base.F64_promote_f32(v53))
				if base.F64_lt(v256, v253) != 0 {
					v264 = v253
				} else {
					if base.F64_gt(v256, float64(1)) == int32(0) {
						v264 = v256
					} else {
						v264 = float64(1)
					}
				}
				v310 = base.F64_add(base.F64_mul(v252, v264), v236)
				m.G0 = v27 + int32(16)
				return v310
			}
		} else {
			v74 = *(*int32)(unsafe.Add(mBase, uint32(l10)+16))
			if v74 <= int32(0) {
				v236 = float64(0)
			} else {
				if v74 == int32(1) {
					v152 = float64(0)
					v157 = int32(0)
					v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+l15))))
					if v170 != int32(1) {
						v187 = v152
					} else {
						v173 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
						v177 = *(*float32)(unsafe.Add(mBase, uint32(v173+v157<<(uint(int32(2))%32))))
						v187 = base.F64_add(v152, base.F64_promote_f32(v177))
					}
				} else {
					v86 = int32(0)
					v95 = float64(0)
					v97 = v86
					v100 = v86
					for {
						v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100+l15))))
						if v113 == int32(1) {
							v116 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
							v120 = *(*float32)(unsafe.Add(mBase, uint32(v116+v100<<(uint(int32(2))%32))))
							v123 = base.F64_add(v95, base.F64_promote_f32(v120))
						} else {
							v123 = v95
						}
						v124 = int32(1)
						v125 = v100 | v124
						v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l15+v125))))
						if v127 == v124 {
							v130 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
							v134 = *(*float32)(unsafe.Add(mBase, uint32(v130+v125<<(uint(int32(2))%32))))
							v137 = base.F64_add(v123, base.F64_promote_f32(v134))
						} else {
							v137 = v123
						}
						v138 = int32(2)
						v139 = v100 + v138
						v141 = v97 + v138
						if v141 != v74&int32(2147483646) {
							v95 = v137
							v97 = v141
							v100 = v139
							continue
						} else {
							break
						}
						break
					}
					if v74&int32(1) == int32(0) {
						v187 = v137
					} else {
						v152 = v137
						v157 = v139
						v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157+l15))))
						if v170 != int32(1) {
							v187 = v152
						} else {
							v173 = *(*int32)(unsafe.Add(mBase, uint32(l10)+20))
							v177 = *(*float32)(unsafe.Add(mBase, uint32(v173+v157<<(uint(int32(2))%32))))
							v187 = base.F64_add(v152, base.F64_promote_f32(v177))
						}
					}
				}
				v204 = float64(0)
				if base.F64_lt(v187, v204) != 0 {
					v236 = v204
				} else {
					if base.F64_gt(v187, float64(1)) == int32(0) {
						v236 = v187
					} else {
						v236 = float64(1)
					}
				}
			}
			if l8|v45 != 0 {
				v252 = float64(0.5)
			} else {
				v241 = *(*int32)(unsafe.Add(mBase, uint32(l17)))
				v242 = base.F64_convert_i32_s(v241)
				v243 = base.F64_sub(l6, v242)
				v244 = base.F64_sub(v42, v242)
				if base.F64_le(v243, v244)|base.F64_lt(v244, float64(0)) != 0 {
					v252 = float64(1)
				} else {
					v252 = base.F64_div(v244, v243)
				}
			}
			v253 = float64(0)
			v256 = base.F64_sub(base.F64_sub(float64(1), v236), base.F64_promote_f32(v53))
			if base.F64_lt(v256, v253) != 0 {
				v264 = v253
			} else {
				if base.F64_gt(v256, float64(1)) == int32(0) {
					v264 = v256
				} else {
					v264 = float64(1)
				}
			}
			v310 = base.F64_add(base.F64_mul(v252, v264), v236)
			m.G0 = v27 + int32(16)
			return v310
		}
	} else {
		if l12 != 0 {
			v267 = *(*float32)(unsafe.Add(mBase, uint32(l12)+8))
			v270 = base.F64_promote_f32(v267)
		} else {
			v270 = float64(0)
		}
		if l8|v45 == int32(0) {
			v275 = base.F64_sub(float64(1), v270)
			if base.F64_lt(v42, float64(0))|base.F64_le(l6, v42) != 0 {
				v310 = v275
			} else {
				v310 = base.F64_mul(base.F64_div(v42, l6), v275)
			}
		} else {
			v310 = base.F64_mul(base.F64_sub(float64(1), v270), float64(0.5))
		}
		m.G0 = v27 + int32(16)
		return v310
	}
}
func F_errhidestmt(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_errhidestmt[0]))
	if v3 < int32(0) {
		*(*int32)(unsafe.Add(mBase, _c_F_errhidestmt[0])) = int32(-1)
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			F_errmsg_internal(m, int32(_a_F_errhidestmt_0), int32(0))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_errhidestmt_1), int32(1629), int32(_a_F_errhidestmt_2))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v24 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v3*int32(100))+uint32(_c_F_errhidestmt[1]))) = uint8(v24)
		return
	}
}
func F_errhint_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
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
	var v71 int32
	_ = v71
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
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(_a_F_errhint_plural_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[0])) = v15 + int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[1]))
	if int32(0) <= v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = int32(_a_F_errhint_plural_1)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[2]))
	v27 = v20 * int32(100)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errhint_plural[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[2])) = v30
	v33 = v11 + int32(16)
	F_initStringInfo(m, v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L4
	} else {
		goto L24
	}
L4:
	;
	return
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errhint_plural[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[5])) = v37
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	if l2 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v42 = l0
	goto L8
L7:
	;
	v42 = l1
	goto L8
L8:
	;
	v43 = F_appendStringInfoVA(m, v33, v42, l3)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L9
	}
L9:
	;
	if v43 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v47 = v43
	goto L13
L11:
	;
	goto L12
L12:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errhint_plural[6])))
	if v71 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v54 = v11 + int32(16)
	F_enlargeStringInfo(m, v54, v47)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errhint_plural[4])))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[5])) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	v61 = F_appendStringInfoVA(m, v54, v42, l3)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v61 != 0 {
		v47 = v61
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	F_pfree(m, v71)
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v75 = F_pstrdup(m, v74)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errhint_plural[6]))) = v75
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_pfree(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[2])) = v24
	v83 = int32(_a_F_errhint_plural_0)
	v85 = *(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errhint_plural[0])) = v85 - int32(1)
	m.G0 = v11 + int32(32)
	return
L24:
	;
	F_errmsg_internal(m, int32(_a_F_errhint_plural_2), int32(0))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_errhint_plural_3), int32(1559), int32(_a_F_errhint_plural_4))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_errmsg_plural(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = int32(_a_F_errmsg_plural_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[0])) = v15 + int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[1]))
	if int32(0) <= v20 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v23 = int32(_a_F_errmsg_plural_1)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[2]))
	v27 = v20 * int32(100)
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errmsg_plural[3])))
	*(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[2])) = v30
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errmsg_plural[4]))) = l0
	v34 = v11 + int32(16)
	F_initStringInfo(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[1])) = int32(-1)
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L4
	} else {
		goto L24
	}
L4:
	;
	return
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errmsg_plural[5])))
	*(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[6])) = v38
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	if l2 == int32(1) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v43 = l0
	goto L8
L7:
	;
	v43 = l1
	goto L8
L8:
	;
	v44 = F_appendStringInfoVA(m, v34, v43, l3)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L4
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
	v46 = v44
	goto L13
L11:
	;
	goto L12
L12:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errmsg_plural[7])))
	if v72 != 0 {
		goto L18
	} else {
		goto L19
	}
L13:
	;
	v55 = v11 + int32(16)
	F_enlargeStringInfo(m, v55, v46)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L4
	} else {
		goto L15
	}
L14:
	;
	goto L12
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errmsg_plural[5])))
	*(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[6])) = v59
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = l3
	v62 = F_appendStringInfoVA(m, v55, v43, l3)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	if v62 != 0 {
		v46 = v62
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	F_pfree(m, v72)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L4
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v76 = F_pstrdup(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L4
	} else {
		goto L22
	}
L21:
	;
	goto L20
L22:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+uint32(_c_F_errmsg_plural[7]))) = v76
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	F_pfree(m, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[2])) = v24
	v84 = int32(_a_F_errmsg_plural_0)
	v86 = *(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_errmsg_plural[0])) = v86 - int32(1)
	m.G0 = v11 + int32(32)
	return
L24:
	;
	F_errmsg_internal(m, int32(_a_F_errmsg_plural_2), int32(0))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_errmsg_plural_3), int32(1379), int32(_a_F_errmsg_plural_4))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L4
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_error_view_not_updatable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v228 int32
	_ = v228
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	v9 = m.G0
	v11 = v9 - int32(224)
	m.G0 = v11
	if l1 != int32(5) {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v11 + int32(224)
	return
L2:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v123 <= int32(0) {
		goto L1
	} else {
		goto L42
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L10
	} else {
		goto L39
	}
L4:
	;
	switch l1 - int32(2) {
	case 0:
		goto L8
	case 1:
		goto L9
	case 2:
		goto L7
	default:
		goto L3
	}
L5:
	;
	goto L6
L6:
	;
	if l2 != 0 {
		goto L2
	} else {
		goto L38
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L10
	} else {
		goto L29
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L10
	} else {
		goto L20
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	return
L11:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v24 + int32(4)
	F_errmsg(m, int32(_a_F_error_view_not_updatable_0), v11+int32(32))
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L10
	} else {
		goto L13
	}
L13:
	;
	if l3 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l3
	F_errdetail_internal(m, int32(_a_F_error_view_not_updatable_1), v11+int32(16))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L10
	} else {
		goto L17
	}
L15:
	;
	goto L16
L16:
	;
	F_errhint(m, int32(_a_F_error_view_not_updatable_2), int32(0))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L10
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	F_errfinish(m, int32(_a_F_error_view_not_updatable_3), int32(3215), int32(_a_F_error_view_not_updatable_4))
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L20:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L10
	} else {
		goto L21
	}
L21:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+64)) = v55 + int32(4)
	F_errmsg(m, int32(_a_F_error_view_not_updatable_5), v11-int32(-64))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L10
	} else {
		goto L22
	}
L22:
	;
	if l3 != 0 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = l3
	F_errdetail_internal(m, int32(_a_F_error_view_not_updatable_1), v11+int32(48))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L10
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	F_errhint(m, int32(_a_F_error_view_not_updatable_6), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L10
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	F_errfinish(m, int32(_a_F_error_view_not_updatable_3), int32(3223), int32(_a_F_error_view_not_updatable_4))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L10
	} else {
		goto L28
	}
L28:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L29:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L10
	} else {
		goto L30
	}
L30:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+96)) = v86 + int32(4)
	F_errmsg(m, int32(_a_F_error_view_not_updatable_7), v11+int32(96))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L10
	} else {
		goto L31
	}
L31:
	;
	if l3 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+80)) = l3
	F_errdetail_internal(m, int32(_a_F_error_view_not_updatable_1), v11+int32(80))
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L10
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_errhint(m, int32(_a_F_error_view_not_updatable_8), int32(0))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L10
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	F_errfinish(m, int32(_a_F_error_view_not_updatable_3), int32(3231), int32(_a_F_error_view_not_updatable_4))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L10
	} else {
		goto L37
	}
L37:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L38:
	;
	goto L1
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l1
	F_errmsg_internal(m, int32(_a_F_error_view_not_updatable_9), v11)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L10
	} else {
		goto L40
	}
L40:
	;
	F_errfinish(m, int32(_a_F_error_view_not_updatable_3), int32(3279), int32(_a_F_error_view_not_updatable_4))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L10
	} else {
		goto L41
	}
L41:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L42:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v127 = int32(0)
	if v127 < v123 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v131 = v123
	goto L45
L44:
	;
	v131 = v127
	goto L45
L45:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v134 = v127
	goto L46
L46:
	;
	v141 = int32(2)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v132+v134<<(uint(v141)%32))))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	switch v145 - v141 {
	case 0:
		goto L51
	case 1:
		goto L52
	case 2:
		goto L50
	default:
		goto L49
	case 5:
		goto L48
	}
L47:
	;
	goto L1
L48:
	;
	v261 = v134 + int32(1)
	if v261 != v131 {
		v134 = v261
		goto L46
	} else {
		goto L95
	}
L49:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L10
	} else {
		goto L92
	}
L50:
	;
	if v126 != 0 {
		goto L79
	} else {
		goto L80
	}
L51:
	;
	if v126 != 0 {
		goto L66
	} else {
		goto L67
	}
L52:
	;
	if v126 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+10)))
	if v148 != 0 {
		goto L48
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L10
	} else {
		goto L57
	}
L56:
	;
	goto L55
L57:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L10
	} else {
		goto L58
	}
L58:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+144)) = v156 + int32(4)
	F_errmsg(m, int32(_a_F_error_view_not_updatable_0), v11+int32(144))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L10
	} else {
		goto L59
	}
L59:
	;
	if l3 != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+128)) = l3
	F_errdetail_internal(m, int32(_a_F_error_view_not_updatable_1), v11+int32(128))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L10
	} else {
		goto L63
	}
L61:
	;
	goto L62
L62:
	;
	F_errhint(m, int32(_a_F_error_view_not_updatable_10), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L10
	} else {
		goto L64
	}
L63:
	;
	goto L62
L64:
	;
	F_errfinish(m, int32(_a_F_error_view_not_updatable_3), int32(3250), int32(_a_F_error_view_not_updatable_4))
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L10
	} else {
		goto L65
	}
L65:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L66:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+15)))
	if v180 != 0 {
		goto L48
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L10
	} else {
		goto L70
	}
L69:
	;
	goto L68
L70:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L10
	} else {
		goto L71
	}
L71:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+176)) = v188 + int32(4)
	F_errmsg(m, int32(_a_F_error_view_not_updatable_5), v11+int32(176))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L10
	} else {
		goto L72
	}
L72:
	;
	if l3 != 0 {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+160)) = l3
	F_errdetail_internal(m, int32(_a_F_error_view_not_updatable_1), v11+int32(160))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L10
	} else {
		goto L76
	}
L74:
	;
	goto L75
L75:
	;
	F_errhint(m, int32(_a_F_error_view_not_updatable_11), int32(0))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L10
	} else {
		goto L77
	}
L76:
	;
	goto L75
L77:
	;
	F_errfinish(m, int32(_a_F_error_view_not_updatable_3), int32(3259), int32(_a_F_error_view_not_updatable_4))
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L10
	} else {
		goto L78
	}
L78:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L79:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126)+20)))
	if v212 != 0 {
		goto L48
	} else {
		goto L82
	}
L80:
	;
	goto L81
L81:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L10
	} else {
		goto L83
	}
L82:
	;
	goto L81
L83:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L10
	} else {
		goto L84
	}
L84:
	;
	v220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+208)) = v220 + int32(4)
	F_errmsg(m, int32(_a_F_error_view_not_updatable_7), v11+int32(208))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L10
	} else {
		goto L85
	}
L85:
	;
	if l3 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+192)) = l3
	F_errdetail_internal(m, int32(_a_F_error_view_not_updatable_1), v11+int32(192))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L10
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	F_errhint(m, int32(_a_F_error_view_not_updatable_12), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L10
	} else {
		goto L90
	}
L89:
	;
	goto L88
L90:
	;
	F_errfinish(m, int32(_a_F_error_view_not_updatable_3), int32(3268), int32(_a_F_error_view_not_updatable_4))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
		goto L10
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v144)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+112)) = v248
	F_errmsg_internal(m, int32(_a_F_error_view_not_updatable_13), v11+int32(112))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L10
	} else {
		goto L93
	}
L93:
	;
	F_errfinish(m, int32(_a_F_error_view_not_updatable_3), int32(3273), int32(_a_F_error_view_not_updatable_4))
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L10
	} else {
		goto L94
	}
L94:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L95:
	;
	goto L47
}
func F_errsave_start(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	if l0 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		if v5 == int32(453) {
			v15 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+4)) = uint8(v15)
			v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+5)))
			if v17 == v15 {
				v20 = int32(_a_F_errsave_start_0)
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_errsave_start[0]))
				v23 = int32(1)
				*(*int32)(unsafe.Add(mBase, _c_F_errsave_start[0])) = v22 + v23
				v26 = int32(_a_F_errsave_start_1)
				v28 = *(*int32)(unsafe.Add(mBase, _c_F_errsave_start[1]))
				v30 = v28 + v23
				*(*int32)(unsafe.Add(mBase, _c_F_errsave_start[1])) = v30
				if int32(5) <= v30 {
					*(*int32)(unsafe.Add(mBase, _c_F_errsave_start[1])) = int32(-1)
					F_errstart_cold(m, int32(24), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						F_errmsg_internal(m, int32(_a_F_errsave_start_2), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_errsave_start_3), int32(783), int32(_a_F_errsave_start_4))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				} else {
					v34 = int32(100)
					v35 = v30 * v34
					base.MemoryFill(m, v35+int32(_a_F_errsave_start_5), int32(0), v34)
					v42 = *(*int32)(unsafe.Add(mBase, _c_F_errsave_start[2]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_errsave_start[3]))) = int32(2600)
					v49 = int32(_a_F_errsave_start_6)
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_errsave_start[4]))) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_errsave_start[5]))) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_errsave_start[6]))) = int32(15)
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_errsave_start[7]))) = v42
					v63 = *(*int32)(unsafe.Add(mBase, _c_F_errsave_start[8]))
					*(*int32)(unsafe.Add(mBase, uint32(v35)+uint32(_c_F_errsave_start[9]))) = v63
					*(*int32)(unsafe.Add(mBase, _c_F_errsave_start[0])) = v22
					v73 = int32(1)
					return v73
				}
			} else {
				v73 = int32(0)
				return v73
			}
		} else {
			v10 = F_errstart(m, int32(21), int32(0))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	} else {
		v10 = F_errstart(m, int32(21), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return v10
		}
	}
}
func F_escape_xml(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v66 int32
	_ = v66
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	F_initStringInfo(m, v7)
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
	v13 = l0
	goto L3
L3:
	;
	v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
	switch v17 - int32(38) {
	case 0:
		goto L10
	case 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 23:
		goto L6
	case 22:
		goto L9
	case 24:
		goto L8
	default:
		goto L11
	}
L4:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	m.G0 = v7 + int32(16)
	return v66
L5:
	;
	goto L4
L6:
	;
	v44 = base.I32_extend8_s(v17)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	if v45 <= v46+int32(1) {
		goto L18
	} else {
		goto L19
	}
L7:
	;
	F_appendStringInfoString(m, v7, int32(_a_F_escape_xml_0))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L17
	}
L8:
	;
	F_appendStringInfoString(m, v7, int32(_a_F_escape_xml_1))
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L16
	}
L9:
	;
	F_appendStringInfoString(m, v7, int32(_a_F_escape_xml_2))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L15
	}
L10:
	;
	F_appendStringInfoString(m, v7, int32(_a_F_escape_xml_3))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L14
	}
L11:
	;
	if v17 == int32(0) {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	if v17 == int32(13) {
		goto L7
	} else {
		goto L13
	}
L13:
	;
	goto L6
L14:
	;
	v13 = v13 + int32(1)
	goto L3
L15:
	;
	v13 = v13 + int32(1)
	goto L3
L16:
	;
	v13 = v13 + int32(1)
	goto L3
L17:
	;
	v13 = v13 + int32(1)
	goto L3
L18:
	;
	F_appendStringInfoChar(m, v7, v44)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v52+v46))) = uint8(v44)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v57 = v55 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v57
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v61 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v59+v57))) = uint8(v61)
	goto L20
L20:
	;
	v13 = v13 + int32(1)
	goto L3
L21:
	;
	goto L20
}
func F_esperanto_UTF_8_create_env(m *base.Module) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_SN_new_env(m, int32(28))
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_examine_indexcol_variable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
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
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+44))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v8+l2<<(uint(int32(2))%32))))
	if v12 != 0 {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v13 != 0 {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+76))
			v29 = v14
			v30 = v13 + v15<<(uint(int32(2))%32)
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+76))
			v29 = v22
			v30 = v21 + v23<<(uint(int32(2))%32) - int32(4)
		}
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
		v32 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v31)+16)))
		*(*int32)(unsafe.Add(mBase, uint32(l3)+4)) = v29
		v35 = *(*int32)(unsafe.Add(mBase, _c_F_examine_indexcol_variable[0]))
		if v35 == int32(0) {
			v63 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+20)))
			v64 = F_SearchSysCache3(m, int32(65), v32, base.I64_extend16_s(base.I64_extend_i32_u(v12)), v63)
			mBase = m.M
			v65 = m.ExcPending
			if v65 != 0 {
				return
			} else {
				v106 = v64
				*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1704)
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v106
				return
			}
		} else {
			v39 = m.T0[v35].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v31, base.I32_extend16_s(v12), l3)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return
			} else {
				if v39 == int32(0) {
					v63 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v31)+20)))
					v64 = F_SearchSysCache3(m, int32(65), v32, base.I64_extend16_s(base.I64_extend_i32_u(v12)), v63)
					mBase = m.M
					v65 = m.ExcPending
					if v65 != 0 {
						return
					} else {
						v106 = v64
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1704)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v106
						return
					}
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					if v43 == int32(0) {
						return
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						if v46 != 0 {
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_examine_indexcol_variable_0), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_examine_indexcol_variable_1), int32(_a_F_examine_indexcol_variable_2), int32(_a_F_examine_indexcol_variable_3))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
			}
		}
	} else {
		v68 = base.I32_extend16_s(l2 + int32(1))
		v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v71 = *(*int32)(unsafe.Add(mBase, _c_F_examine_indexcol_variable[1]))
		if v71 == int32(0) {
			v99 = F_SearchSysCache3(m, int32(65), base.I64_extend_i32_u(v69), base.I64_extend_i32_s(v68), int64(0))
			mBase = m.M
			v100 = m.ExcPending
			if v100 != 0 {
				return
			} else {
				v106 = v99
				*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1704)
				*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v106
				return
			}
		} else {
			v74 = m.T0[v71].(func(*base.Module, int32, int32, int32, int32) int32)(m, l0, v69, v68, l3)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return
			} else {
				if v74 == int32(0) {
					v99 = F_SearchSysCache3(m, int32(65), base.I64_extend_i32_u(v69), base.I64_extend_i32_s(v68), int64(0))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return
					} else {
						v106 = v99
						*(*int32)(unsafe.Add(mBase, uint32(l3)+12)) = int32(1704)
						*(*int32)(unsafe.Add(mBase, uint32(l3)+8)) = v106
						return
					}
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
					if v78 == int32(0) {
						return
					} else {
						v81 = *(*int32)(unsafe.Add(mBase, uint32(l3)+12))
						if v81 != 0 {
							return
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v85 = m.ExcPending
							if v85 != 0 {
								return
							} else {
								F_errmsg_internal(m, int32(_a_F_examine_indexcol_variable_0), int32(0))
								mBase = m.M
								v89 = m.ExcPending
								if v89 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_examine_indexcol_variable_1), int32(_a_F_examine_indexcol_variable_4), int32(_a_F_examine_indexcol_variable_3))
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
			}
		}
	}
}
func F_execTuplesHashPrepare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	v16 = F_palloc(m, l0<<(uint(int32(2))%32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v16
	v21 = F_palloc(m, l0*int32(28))
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v21
	if int32(0) < l0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L15
	}
L5:
	;
	v31 = int32(0)
	goto L8
L6:
	;
	goto L7
L7:
	;
	m.G0 = v12 + int32(16)
	return
L8:
	;
	v36 = v31 << (uint(int32(2)) % 32)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1+v36)))
	v39 = F_get_opcode(m, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L1
	} else {
		goto L10
	}
L9:
	;
	goto L7
L10:
	;
	v45 = F_get_op_hash_functions(m, v38, v12+int32(12), v12+int32(8))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v45 == int32(0) {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v49+v36))) = v39
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F_fmgr_info(m, v52, v53+v31*int32(28))
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	v60 = v31 + int32(1)
	if v60 != l0 {
		v31 = v60
		goto L8
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v38
	F_errmsg_internal(m, int32(_a_F_execTuplesHashPrepare_0), v12)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_execTuplesHashPrepare_1), int32(121), int32(_a_F_execTuplesHashPrepare_2))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_exec_move_row_from_fields(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
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
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int64
	_ = v190
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v219 int32
	_ = v219
	var v223 int64
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v272 int32
	_ = v272
	var v289 int32
	_ = v289
	var v297 int32
	_ = v297
	var v315 int32
	_ = v315
	var v325 int32
	_ = v325
	var v344 int32
	_ = v344
	var v346 int32
	_ = v346
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v388 int32
	_ = v388
	var v394 int32
	_ = v394
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v435 int32
	_ = v435
	var v438 int32
	_ = v438
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v465 int64
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int64
	_ = v483
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v495 int64
	_ = v495
	var v496 int32
	_ = v496
	var v498 int64
	_ = v498
	var v499 int32
	_ = v499
	var v505 int64
	_ = v505
	var v507 int32
	_ = v507
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v518 int64
	_ = v518
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v529 int32
	_ = v529
	var v531 int32
	_ = v531
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v592 int32
	_ = v592
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v649 int32
	_ = v649
	var v656 int32
	_ = v656
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v676 int32
	_ = v676
	var v688 int32
	_ = v688
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v711 int32
	_ = v711
	var v719 int32
	_ = v719
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int64
	_ = v738
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v773 int64
	_ = v773
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v797 int32
	_ = v797
	var v799 int64
	_ = v799
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v814 int32
	_ = v814
	var v832 int32
	_ = v832
	var v842 int32
	_ = v842
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v872 int32
	_ = v872
	var v876 int32
	_ = v876
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	v7 = int32(0)
	v23 = m.G0
	v25 = v23 - int32(656)
	m.G0 = v25
	if l5 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	switch v57 - int32(1) {
	case 0:
		goto L13
	case 1:
		goto L12
	default:
		goto L11
	}
L2:
	;
	v53 = v7
	v54 = int32(_a_F_exec_move_row_from_fields_0)
	v55 = v7
	v56 = int32(1)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_exec_move_row_from_fields[0])))
	if v33&int32(8) != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v53 = v31
	v54 = int32(_a_F_exec_move_row_from_fields_1)
	v55 = int32(21)
	v56 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_fields[1]))
	v53 = v31
	v54 = int32(_a_F_exec_move_row_from_fields_0)
	v55 = v41 << (uint(int32(28)) % 32) >> (uint(int32(31)) % 32) & int32(19)
	v56 = base.B2i32(v41&int32(8) == int32(0))
	goto L1
L8:
	;
	m.G0 = v25 + int32(656)
	return
L9:
	;
	if v56|base.B2i32(v53 <= v814) != 0 {
		goto L8
	} else {
		goto L160
	}
L10:
	;
	v649 = int32(0)
	v656 = v7
	goto L139
L11:
	;
	F_errstart_cold(m, int32(21), int32(_a_F_exec_move_row_from_fields_2))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L18
	} else {
		goto L136
	}
L12:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	if v64 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L13:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if int32(0) < v60 {
		goto L10
	} else {
		goto L14
	}
L14:
	;
	v814 = int32(0)
	goto L9
L15:
	;
	v67 = F_expanded_record_fetch_tupdesc(m, l2)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v69 = v64
	goto L17
L17:
	;
	if l5 == v69 {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	return
L19:
	;
	v69 = v67
	goto L17
L20:
	;
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	if v404&int32(4) == int32(0) {
		goto L70
	} else {
		goto L71
	}
L21:
	;
	v388 = l3
	v394 = l4
	goto L20
L22:
	;
	goto L23
L23:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if base.Ui32(v71) < base.Ui32(int32(65)) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v89 = int32(0)
	if v89 < v71 {
		goto L29
	} else {
		goto L30
	}
L25:
	;
	v87 = v25 + int32(144)
	v88 = v25 + int32(80)
	goto L24
L26:
	;
	goto L27
L27:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(l0)+128))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v78)+20))
	v82 = F_MemoryContextAlloc(m, v79, v71*int32(9))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L28
	}
L28:
	;
	v87 = v82
	v88 = v82 + v71<<(uint(int32(3))%32)
	goto L24
L29:
	;
	v98 = v89
	v105 = v7
	goto L32
L30:
	;
	v297 = v89
	goto L31
L31:
	;
	if v56|base.B2i32(v53 <= v297) != 0 {
		v388 = v87
		v394 = v88
		goto L20
	} else {
		goto L56
	}
L32:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v120 = v69 + v114<<(uint(int32(3))%32) + v105*int32(100)
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v120)+119)))
	if v121 == int32(0) {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	v297 = v272
	goto L31
L34:
	;
	v125 = v120 + int32(28)
	if v98 < v53 {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	v272 = v98
	goto L36
L36:
	;
	v289 = v105 + int32(1)
	if v289 != v71 {
		v98 = v272
		v105 = v289
		goto L32
	} else {
		goto L55
	}
L37:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v125)+68))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v125)+76))
	v260 = F_exec_cast_value(m, l0, v252, v25+int32(79), v240, v250, v258, v259)
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L18
	} else {
		goto L54
	}
L38:
	;
	v223 = *(*int64)(unsafe.Add(mBase, uint32(l3+v139<<(uint(int32(3))%32))))
	v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v139))))
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+79)) = uint8(v225)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v157)+76))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(v157)+68))
	v237 = v139 + int32(1)
	v240 = v230
	v250 = v229
	v252 = v223
	goto L37
L39:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v139 = v98
	goto L42
L40:
	;
	v170 = v98
	goto L41
L41:
	;
	v186 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v25)+79)) = uint8(v186)
	v188 = int32(-1)
	v189 = int32(705)
	v190 = int64(0)
	if v56 != 0 {
		v237 = v170
		v240 = v189
		v250 = v188
		v252 = v190
		goto L37
	} else {
		goto L46
	}
L42:
	;
	v157 = l5 + v127<<(uint(int32(3))%32) + int32(28) + v139*int32(100)
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v157)+91)))
	if v158 != int32(1) {
		goto L38
	} else {
		goto L44
	}
L43:
	;
	v170 = v53
	goto L41
L44:
	;
	v162 = v139 + int32(1)
	if v162 != v53 {
		v139 = v162
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v192 = F_errstart(m, v55, int32(_a_F_exec_move_row_from_fields_2))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L18
	} else {
		goto L47
	}
L47:
	;
	if v192 == int32(0) {
		v237 = v170
		v240 = v189
		v250 = v188
		v252 = v190
		goto L37
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_errmsg(m, int32(_a_F_exec_move_row_from_fields_3), int32(0))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+36)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v25)+32)) = int32(_a_F_exec_move_row_from_fields_4)
	v209 = F_errdetail(m, int32(_a_F_exec_move_row_from_fields_5), v25+int32(32))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	F_errhint(m, int32(_a_F_exec_move_row_from_fields_6), int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(_a_F_exec_move_row_from_fields_7), int32(_a_F_exec_move_row_from_fields_8), int32(_a_F_exec_move_row_from_fields_9))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L18
	} else {
		goto L53
	}
L53:
	;
	v237 = v170
	v240 = v189
	v250 = v188
	v252 = v190
	goto L37
L54:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v87+v105<<(uint(int32(3))%32)))) = v260
	v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+79)))
	*(*uint8)(unsafe.Add(mBase, uint32(v105+v88))) = uint8(v264)
	v272 = v237
	goto L36
L55:
	;
	goto L33
L56:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v325 = v297
	goto L57
L57:
	;
	v344 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v315<<(uint(int32(3))%32)+v325*int32(100))+119)))
	if v344 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v349 = F_errstart(m, v55, int32(_a_F_exec_move_row_from_fields_2))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L18
	} else {
		goto L63
	}
L59:
	;
	v346 = v325 + int32(1)
	if v53 != v346 {
		v325 = v346
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	v388 = v87
	v394 = v88
	goto L20
L63:
	;
	if v349 == int32(0) {
		v388 = v87
		v394 = v88
		goto L20
	} else {
		goto L64
	}
L64:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	F_errmsg(m, int32(_a_F_exec_move_row_from_fields_3), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L18
	} else {
		goto L66
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+20)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = int32(_a_F_exec_move_row_from_fields_4)
	v366 = F_errdetail(m, int32(_a_F_exec_move_row_from_fields_5), v25+int32(16))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L18
	} else {
		goto L67
	}
L67:
	;
	F_errhint(m, int32(_a_F_exec_move_row_from_fields_6), int32(0))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L18
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(_a_F_exec_move_row_from_fields_7), int32(_a_F_exec_move_row_from_fields_10), int32(_a_F_exec_move_row_from_fields_9))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L18
	} else {
		goto L69
	}
L69:
	;
	v388 = v87
	v394 = v88
	goto L20
L70:
	;
	F_deconstruct_expanded_record(m, l2)
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L18
	} else {
		goto L73
	}
L71:
	;
	v412 = v404
	goto L72
L72:
	;
	v413 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+68)) = v413
	v416 = v412 & int32(-2)
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v416
	v418 = *(*int32)(unsafe.Add(mBase, uint32(l2)+44))
	v419 = int32(_a_F_exec_move_row_from_fields_11)
	v420 = *(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_fields[2]))
	v422 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_fields[2])) = v422
	v424 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	if v413 < v424 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v412 = v411
	goto L72
L74:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(l2)+60))
	v428 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	v435 = int32(0)
	v438 = v424
	goto L77
L75:
	;
	v554 = v416
	goto L76
L76:
	;
	if v554&int32(64) != 0 {
		goto L105
	} else {
		goto L106
	}
L77:
	;
	v455 = v435 << (uint(int32(3)) % 32)
	v456 = v418 + int32(28) + v455
	v457 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+6)))
	if v457&int32(4) == int32(0) {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v531 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	v554 = v531
	goto L76
L79:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+v394))))
	v465 = *(*int64)(unsafe.Add(mBase, uint32(v455+v388)))
	v466 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+4)))
	if v466 != 0 {
		v518 = v465
		goto L82
	} else {
		goto L83
	}
L80:
	;
	v526 = v438
	goto L81
L81:
	;
	v529 = v435 + int32(1)
	if v529 < v526 {
		v435 = v529
		v438 = v526
		goto L77
	} else {
		goto L104
	}
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v455+v428))) = v518
	*(*uint8)(unsafe.Add(mBase, uint32(v435+v427))) = uint8(v463)
	v523 = *(*int32)(unsafe.Add(mBase, uint32(l2)+64))
	v526 = v523
	goto L81
L83:
	;
	if v463&int32(1) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v471 = int32(*(*int16)(unsafe.Add(mBase, uint32(v456)+2)))
	if v471 != int32(-1) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v505 = v465
	goto L86
L86:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435+v427))))
	if v507 != 0 {
		v518 = v505
		goto L82
	} else {
		goto L98
	}
L87:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v499 | int32(8)
	v505 = v498
	goto L86
L88:
	;
	v495 = F_datumCopy(m, v465, int32(0), v471)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L18
	} else {
		goto L97
	}
L89:
	;
	v474 = base.I32_wrap_i64(v465)
	v475 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
	if v475 != int32(1) {
		goto L88
	} else {
		goto L90
	}
L90:
	;
	if (v399^int32(-1))&int32(1) != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v478 = F_detoast_external_attr(m, v474)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L18
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v483 = F_datumCopy(m, v465, int32(0), int32(-1))
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L18
	} else {
		goto L95
	}
L94:
	;
	v498 = base.I64_extend_i32_u(v478)
	goto L87
L95:
	;
	v486 = int32(*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v483)))))
	if v486 != int32(1) {
		v498 = v483
		goto L87
	} else {
		goto L96
	}
L96:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
	*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v489 | int32(16)
	v498 = v483
	goto L87
L97:
	;
	v498 = v495
	goto L87
L98:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(v455+v428)))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l2)+88))
	if base.Ui32(v510) <= base.Ui32(v509) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v512 = *(*int32)(unsafe.Add(mBase, uint32(l2)+92))
	if base.Ui32(v509) < base.Ui32(v512) {
		v518 = v505
		goto L82
	} else {
		goto L102
	}
L100:
	;
	goto L101
L101:
	;
	F_pfree(m, v509)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L18
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v518 = v505
	goto L82
L104:
	;
	goto L78
L105:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	if v557 == int32(0) {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L107
L107:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_fields[2])) = v420
	v587 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v588 = *(*int32)(unsafe.Add(mBase, uint32(l0)+80))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v587)+16))
	if v592 != v588 {
		goto L116
	} else {
		goto L117
	}
L108:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_exec_move_row_from_fields[2])) = v571
	v578 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	v581 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	F_domain_check(m, base.I64_extend_i32_u(l2+int32(18)), int32(0), v578, l2+int32(104), v581)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L18
	} else {
		goto L114
	}
L109:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v565 = F_AllocSetContextCreateInternal(m, v560, int32(_a_F_exec_move_row_from_fields_12), int32(0), int32(1024), int32(_a_F_exec_move_row_from_fields_13))
	mBase = m.M
	v566 = m.ExcPending
	if v566 != 0 {
		goto L18
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	F_MemoryContextReset(m, v557)
	mBase = m.M
	v569 = m.ExcPending
	if v569 != 0 {
		goto L18
	} else {
		goto L113
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l2)+96)) = v565
	v571 = v565
	goto L108
L113:
	;
	v570 = *(*int32)(unsafe.Add(mBase, uint32(l2)+96))
	v571 = v570
	goto L108
L114:
	;
	goto L107
L115:
	;
	v621 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v621 != 0 {
		goto L132
	} else {
		goto L133
	}
L116:
	;
	if v592 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L117:
	;
	goto L118
L118:
	;
	goto L115
L119:
	;
	if v588 != 0 {
		goto L126
	} else {
		goto L127
	}
L120:
	;
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v587)+28))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v587)+24))
	if v597 != 0 {
		goto L122
	} else {
		goto L123
	}
L121:
	;
	if v596 == int32(0) {
		goto L119
	} else {
		goto L125
	}
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v597)+28)) = v596
	goto L121
L123:
	;
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v592)+20)) = v596
	goto L121
L125:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v587)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v596)+24)) = v602
	goto L119
L126:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v587)+24)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v587)+16)) = v588
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v588)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v587)+28)) = v609
	if v609 != 0 {
		goto L129
	} else {
		goto L130
	}
L127:
	;
	goto L128
L128:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v587)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v587)+16)) = int32(0)
	goto L118
L129:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v609)+24)) = v587
	goto L131
L130:
	;
	goto L131
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v588)+20)) = v587
	goto L115
L132:
	;
	F_DeleteExpandedObject(m, base.I64_extend_i32_u(v621+int32(12)))
	mBase = m.M
	v626 = m.ExcPending
	if v626 != 0 {
		goto L18
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = l2
	goto L8
L135:
	;
	goto L134
L136:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v25))) = v632
	F_errmsg_internal(m, int32(_a_F_exec_move_row_from_fields_14), v25)
	mBase = m.M
	v636 = m.ExcPending
	if v636 != 0 {
		goto L18
	} else {
		goto L137
	}
L137:
	;
	F_errfinish(m, int32(_a_F_exec_move_row_from_fields_7), int32(_a_F_exec_move_row_from_fields_15), int32(_a_F_exec_move_row_from_fields_9))
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L18
	} else {
		goto L138
	}
L138:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L139:
	;
	v665 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v666 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v667 = int32(2)
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v666+v656<<(uint(v667)%32))))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v665+v670<<(uint(v667)%32))))
	if v649 < v53 {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v814 = v784
	goto L9
L141:
	;
	F_exec_assign_value(m, l0, v674, v799, v780&int32(1), v787, v797)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L18
	} else {
		goto L158
	}
L142:
	;
	v769 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4+v688))))
	v773 = *(*int64)(unsafe.Add(mBase, uint32(l3+v688<<(uint(int32(3))%32))))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v706)+76))
	v775 = *(*int32)(unsafe.Add(mBase, uint32(v706)+68))
	v780 = v769
	v784 = v688 + int32(1)
	v787 = v775
	v797 = v774
	v799 = v773
	goto L141
L143:
	;
	v676 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v688 = v649
	goto L146
L144:
	;
	v719 = v649
	goto L145
L145:
	;
	v735 = int32(-1)
	v736 = int32(705)
	v737 = int32(1)
	v738 = int64(0)
	if v56 != 0 {
		v780 = v737
		v784 = v719
		v787 = v736
		v797 = v735
		v799 = v738
		goto L141
	} else {
		goto L150
	}
L146:
	;
	v706 = l5 + v676<<(uint(int32(3))%32) + int32(28) + v688*int32(100)
	v707 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v706)+91)))
	if v707 != int32(1) {
		goto L142
	} else {
		goto L148
	}
L147:
	;
	v719 = v53
	goto L145
L148:
	;
	v711 = v688 + int32(1)
	if v711 != v53 {
		v688 = v711
		goto L146
	} else {
		goto L149
	}
L149:
	;
	goto L147
L150:
	;
	v740 = F_errstart(m, v55, int32(_a_F_exec_move_row_from_fields_2))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L18
	} else {
		goto L151
	}
L151:
	;
	if v740 == int32(0) {
		v780 = v737
		v784 = v719
		v787 = v736
		v797 = v735
		v799 = v738
		goto L141
	} else {
		goto L152
	}
L152:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L18
	} else {
		goto L153
	}
L153:
	;
	F_errmsg(m, int32(_a_F_exec_move_row_from_fields_3), int32(0))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L18
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+68)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v25)+64)) = int32(_a_F_exec_move_row_from_fields_4)
	v757 = F_errdetail(m, int32(_a_F_exec_move_row_from_fields_5), v25-int32(-64))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L18
	} else {
		goto L155
	}
L155:
	;
	F_errhint(m, int32(_a_F_exec_move_row_from_fields_6), int32(0))
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L18
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_exec_move_row_from_fields_7), int32(_a_F_exec_move_row_from_fields_16), int32(_a_F_exec_move_row_from_fields_9))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L18
	} else {
		goto L157
	}
L157:
	;
	v780 = v737
	v784 = v719
	v787 = v736
	v797 = v735
	v799 = v738
	goto L141
L158:
	;
	v805 = v656 + int32(1)
	v806 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	if v805 < v806 {
		v649 = v784
		v656 = v805
		goto L139
	} else {
		goto L159
	}
L159:
	;
	goto L140
L160:
	;
	v832 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v842 = v814
	goto L161
L161:
	;
	v861 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5+v832<<(uint(int32(3))%32)+v842*int32(100))+119)))
	if v861 != 0 {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v866 = F_errstart(m, v55, int32(_a_F_exec_move_row_from_fields_2))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L18
	} else {
		goto L167
	}
L163:
	;
	v863 = v842 + int32(1)
	if v53 != v863 {
		v842 = v863
		goto L161
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	goto L162
L166:
	;
	goto L8
L167:
	;
	if v866 == int32(0) {
		goto L8
	} else {
		goto L168
	}
L168:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L18
	} else {
		goto L169
	}
L169:
	;
	F_errmsg(m, int32(_a_F_exec_move_row_from_fields_3), int32(0))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L18
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v25)+52)) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v25)+48)) = int32(_a_F_exec_move_row_from_fields_4)
	v883 = F_errdetail(m, int32(_a_F_exec_move_row_from_fields_5), v25+int32(48))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L18
	} else {
		goto L171
	}
L171:
	;
	F_errhint(m, int32(_a_F_exec_move_row_from_fields_6), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L18
	} else {
		goto L172
	}
L172:
	;
	F_errfinish(m, int32(_a_F_exec_move_row_from_fields_7), int32(_a_F_exec_move_row_from_fields_17), int32(_a_F_exec_move_row_from_fields_9))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L18
	} else {
		goto L173
	}
L173:
	;
	goto L8
}
func F_exp_var(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int64
	_ = v14
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v51 float64
	_ = v51
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int64
	_ = v64
	var v79 int32
	_ = v79
	var v82 float64
	_ = v82
	var v87 float64
	_ = v87
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v275 int32
	_ = v275
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v14 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v14
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v14
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v25 = F_palloc(m, v20<<(uint(int32(1))%32)+int32(2))
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v27 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v25))) = uint16(v27)
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v29 <= v27 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v41 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+32)) = v41
	v43 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v43
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v25
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v25 + int32(2)
	v51 = F_numericvar_to_double_no_overflow(m, v12+int32(24))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L8
	}
L4:
	;
	v33 = v29 << (uint(int32(1)) % 32)
	if v33 == int32(0) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	base.MemoryCopy(m, v25+int32(2), v38, v33)
	goto L3
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L81
	}
L7:
	;
	m.G0 = v12 + int32(48)
	return
L8:
	;
	v53 = base.F64_abs(v51)
	if base.F64_ge(v53, float64(6000)) != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if base.F64_gt(v51, float64(0)) != 0 {
		goto L6
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if base.F64_gt(v53, float64(0.01)) != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if v58 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	F_pfree(m, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	v64 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v64
	*(*int64)(unsafe.Add(mBase, uint32(l1)+16)) = v64
	goto L7
L16:
	;
	goto L15
L17:
	;
	v79 = int32(1)
	v82 = v51
	goto L20
L18:
	;
	v101 = int32(0)
	goto L19
L19:
	;
	v112 = v12 + int32(24)
	F_add_var(m, int32(_a_F_exp_var_0), v112, l1)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L24
	}
L20:
	;
	v87 = base.F64_mul(v82, float64(0.5))
	if base.F64_gt(base.F64_abs(v87), float64(0.01)) != 0 {
		v79 = v79 + int32(1)
		v82 = v87
		goto L20
	} else {
		goto L22
	}
L21:
	;
	v92 = v12 + int32(24)
	v93 = int32(1)
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v12)+36))
	F_div_var_int(m, v92, v93<<(uint(v79)%32), int32(0), v92, v96+v79, v93)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L23
	}
L22:
	;
	goto L21
L23:
	;
	v101 = v79
	goto L19
L24:
	;
	v122 = base.I32_trunc_sat_f64_s(base.F64_mul(base.F64_convert_i32_s(v101), float64(0.301029995663981))) + (l2 + base.I32_trunc_sat_f64_s(base.F64_mul(v51, float64(0.434294481903252)))) + int32(1)
	v123 = int32(0)
	if v123 < v122 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v126 = v122
	goto L27
L26:
	;
	v126 = v123
	goto L27
L27:
	;
	v128 = v126 + int32(7)
	F_mul_var(m, v112, v112, v12, v128)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v131 = int32(2)
	F_div_var_int(m, v12, v131, int32(0), v12, v128, int32(1))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v137 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v143 = v131
	goto L33
L31:
	;
	goto L32
L32:
	;
	if int32(0) < v101 {
		goto L39
	} else {
		goto L40
	}
L33:
	;
	F_add_var(m, l1, v12, l1)
	mBase = m.M
	v148 = m.ExcPending
	if v148 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	F_mul_var(m, v12, v12+int32(24), v12, v128)
	mBase = m.M
	v152 = m.ExcPending
	if v152 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v153 = int32(1)
	v154 = v143 + v153
	F_div_var_int(m, v12, v154, int32(0), v12, v128, v153)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	if v159 != 0 {
		v143 = v154
		goto L33
	} else {
		goto L38
	}
L38:
	;
	goto L34
L39:
	;
	v173 = v101
	goto L42
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+12)) = l2
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v214 = l2 + v211<<(uint(int32(2))%32)
	if v214+int32(4) < int32(0) {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v182 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v185 = v126 + int32(8) - v182<<(uint(int32(3))%32)
	v186 = int32(0)
	if v186 < v185 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L41
L44:
	;
	v189 = v185
	goto L46
L45:
	;
	v189 = v186
	goto L46
L46:
	;
	F_mul_var(m, l1, l1, l1, v189)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v192 = int32(1)
	if base.Ui32(v192) < base.Ui32(v173) {
		v173 = v173 - v192
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	if v329 != 0 {
		goto L75
	} else {
		goto L76
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = int64(0)
	goto L49
L51:
	;
	goto L52
L52:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v225 = l2 & int32(3)
	v229 = base.I32_div_s(v214+int32(7), int32(4))
	v230 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v230 <= v229 {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	goto L49
L54:
	;
	if int32(0) <= v295 {
		goto L53
	} else {
		goto L74
	}
L55:
	;
	v275 = v269
	goto L68
L56:
	;
	v244 = int32(1)
	v245 = v229 - v244
	v248 = v223 + v245<<(uint(v244)%32)
	v249 = int32(*(*int16)(unsafe.Add(mBase, uint32(v248))))
	v250 = int32(2)
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v225<<(uint(v250)%32))+uint32(_c_F_exp_var[0])))
	v253 = base.I32_rem_s(v249, v252)
	v254 = v249 - v253
	*(*uint16)(unsafe.Add(mBase, uint32(v248))) = uint16(v254)
	v257 = base.I32_div_s(v252, v250)
	if v253 < v257 {
		v295 = v245
		goto L54
	} else {
		goto L63
	}
L57:
	;
	if base.B2i32(v225 == int32(0))|base.B2i32(v229 != v230) != 0 {
		goto L53
	} else {
		goto L60
	}
L58:
	;
	goto L59
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v229
	if v225 != 0 {
		goto L56
	} else {
		goto L61
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v229
	goto L56
L61:
	;
	v241 = int32(*(*int16)(unsafe.Add(mBase, uint32(v223+v229<<(uint(int32(1))%32)))))
	if v241 <= int32(_a_F_exp_var_1) {
		v295 = v229
		goto L54
	} else {
		goto L62
	}
L62:
	;
	v269 = v229
	goto L55
L63:
	;
	v260 = v252 + base.I32_extend16_s(v254)
	if int32(_a_F_exp_var_2) < v260 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v265 = v260 + int32(_a_F_exp_var_3)
	goto L66
L65:
	;
	v265 = v260
	goto L66
L66:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v248))) = uint16(v265)
	if v260 < int32(_a_F_exp_var_4) {
		v295 = v245
		goto L54
	} else {
		goto L67
	}
L67:
	;
	v269 = v245
	goto L55
L68:
	;
	v281 = int32(1)
	v282 = v275 - v281
	v285 = v223 + v282<<(uint(v281)%32)
	v288 = int32(*(*int16)(unsafe.Add(mBase, uint32(v285))))
	v290 = base.B2i32(int32(_a_F_exp_var_5) < v288)
	if int32(_a_F_exp_var_5) < v288 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v295 = v282
	goto L54
L70:
	;
	v291 = int32(-9999)
	goto L72
L71:
	;
	v291 = v281
	goto L72
L72:
	;
	v292 = v291 + v288
	*(*uint16)(unsafe.Add(mBase, uint32(v285))) = uint16(v292)
	if int32(_a_F_exp_var_5) < v288 {
		v275 = v282
		goto L68
	} else {
		goto L73
	}
L73:
	;
	goto L69
L74:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+20)) = v303 - int32(2)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v308 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v307 + v308
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v311 + v308
	goto L53
L75:
	;
	F_pfree(m, v329)
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L1
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
	if v332 == int32(0) {
		goto L7
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	F_pfree(m, v332)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	goto L7
L81:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	F_errmsg(m, int32(_a_F_exp_var_6), int32(0))
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errfinish(m, int32(_a_F_exp_var_7), int32(_a_F_exp_var_8), int32(_a_F_exp_var_9))
	mBase = m.M
	v364 = m.ExcPending
	if v364 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_expand_groupingset_node(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
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
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v13 {
	case 0:
		goto L6
	case 1:
		goto L5
	case 2:
		goto L4
	case 3:
		goto L3
	case 4:
		goto L2
	default:
		v178 = v2
		goto L1
	}
L1:
	;
	m.G0 = v11 + int32(16)
	return v178
L2:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v148 == int32(0) {
		v178 = v2
		goto L1
	} else {
		goto L44
	}
L3:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v89 != 0 {
		goto L25
	} else {
		goto L26
	}
L4:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v31 == int32(0) {
		v80 = v2
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v23
	v29 = F_list_make1_impl(m, int32(1), v11+int32(4))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L7
	} else {
		goto L9
	}
L6:
	;
	v14 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v14
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v14
	v19 = F_list_make1_impl(m, int32(1), v11)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	return int32(0)
L8:
	;
	v178 = v19
	goto L1
L9:
	;
	v178 = v29
	goto L1
L10:
	;
	v87 = F_lappend(m, v80, int32(0))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L24
	}
L11:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v34 <= int32(0) {
		v80 = v2
		goto L10
	} else {
		goto L12
	}
L12:
	;
	v37 = v34
	v39 = v2
	goto L13
L13:
	;
	v45 = int32(0)
	v47 = v45
	v52 = v45
	v53 = v37
	goto L15
L14:
	;
	v80 = v74
	goto L10
L15:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v47 < v55 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	v74 = F_lappend(m, v39, v70)
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L7
	} else {
		goto L22
	}
L17:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v57+v47<<(uint(int32(2))%32))))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v61)+8))
	v63 = F_list_concat(m, v52, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L7
	} else {
		goto L20
	}
L18:
	;
	v70 = v52
	goto L19
L19:
	;
	goto L16
L20:
	;
	v65 = int32(1)
	v68 = v53 - v65
	if v68 != 0 {
		v47 = v47 + v65
		v52 = v63
		v53 = v68
		goto L15
	} else {
		goto L21
	}
L21:
	;
	v70 = v63
	goto L19
L22:
	;
	if int32(1) < v37 {
		v37 = v37 - int32(1)
		v39 = v74
		goto L13
	} else {
		goto L23
	}
L23:
	;
	goto L14
L24:
	;
	v178 = v87
	goto L1
L25:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	v91 = v90
	goto L27
L26:
	;
	v91 = v2
	goto L27
L27:
	;
	v94 = v2
	v98 = v2
	goto L28
L28:
	;
	if v89 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	v178 = v141
	goto L1
L30:
	;
	v141 = F_lappend(m, v94, v136)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L42
	}
L31:
	;
	v136 = int32(0)
	goto L30
L32:
	;
	goto L33
L33:
	;
	v104 = int32(0)
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v106 <= v104 {
		v136 = v104
		goto L30
	} else {
		goto L34
	}
L34:
	;
	v109 = int32(1)
	v112 = v104
	v114 = v104
	goto L35
L35:
	;
	if v109&v98 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	v136 = v126
	goto L30
L37:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v118+v114<<(uint(int32(2))%32))))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	v124 = F_list_concat(m, v112, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L40
	}
L38:
	;
	v126 = v112
	goto L39
L39:
	;
	v127 = int32(1)
	v130 = v114 + v127
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
	if v130 < v131 {
		v109 = v109 << (uint(v127) % 32)
		v112 = v126
		v114 = v130
		goto L35
	} else {
		goto L41
	}
L40:
	;
	v126 = v124
	goto L39
L41:
	;
	goto L36
L42:
	;
	v144 = v98 + int32(1)
	if int32(base.Ui32(v144)>>(uint(v91)%32)) == int32(0) {
		v94 = v141
		v98 = v144
		goto L28
	} else {
		goto L43
	}
L43:
	;
	goto L29
L44:
	;
	v151 = int32(0)
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v152 <= v151 {
		v178 = v2
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v155 = v151
	v157 = v2
	goto L46
L46:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v148)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v163+v155<<(uint(int32(2))%32))))
	v168 = F_expand_groupingset_node(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L7
	} else {
		goto L48
	}
L47:
	;
	v178 = v170
	goto L1
L48:
	;
	v170 = F_list_concat(m, v157, v168)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L49
	}
L49:
	;
	v173 = v155 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	if v173 < v174 {
		v155 = v173
		v157 = v170
		goto L46
	} else {
		goto L50
	}
L50:
	;
	goto L47
}
func F_explicit_bzero(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v71 int32
	_ = v71
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int64
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v3 = int32(0)
	if l1 == v3 {
	} else {
		*(*uint8)(unsafe.Add(mBase, uint32(l0))) = uint8(v3)
		v10 = l0 + l1
		*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(1)))) = uint8(v3)
		if base.Ui32(l1) < base.Ui32(int32(3)) {
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+2)) = uint8(v3)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+1)) = uint8(v3)
			*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(3)))) = uint8(v3)
			*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(2)))) = uint8(v3)
			if base.Ui32(l1) < base.Ui32(int32(7)) {
			} else {
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+3)) = uint8(v3)
				*(*uint8)(unsafe.Add(mBase, uint32(v10-int32(4)))) = uint8(v3)
				if base.Ui32(l1) < base.Ui32(int32(9)) {
				} else {
					v32 = int32(0)
					v35 = (v32 - l0) & int32(3)
					v36 = l0 + v35
					*(*int32)(unsafe.Add(mBase, uint32(v36))) = v32
					v44 = (l1 - v35) & int32(-4)
					v45 = v36 + v44
					*(*int32)(unsafe.Add(mBase, uint32(v45-int32(4)))) = v32
					if base.Ui32(v44) < base.Ui32(int32(9)) {
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v36)+4)) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v45-int32(8)))) = v32
						*(*int32)(unsafe.Add(mBase, uint32(v45-int32(12)))) = v32
						if base.Ui32(v44) < base.Ui32(int32(25)) {
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v36)+24)) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v36)+20)) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v36)+16)) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v36)+12)) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v45-int32(16)))) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v45-int32(20)))) = v32
							v71 = int32(24)
							*(*int32)(unsafe.Add(mBase, uint32(v45-v71))) = v32
							*(*int32)(unsafe.Add(mBase, uint32(v45-int32(28)))) = v32
							v80 = v36&int32(4) | v71
							v81 = v44 - v80
							if base.Ui32(v81) < base.Ui32(int32(32)) {
							} else {
								v86 = base.I64_extend_i32_u(v32) * int64(4294967297)
								v89 = v80 + v36
								v90 = v81
								for {
									*(*int64)(unsafe.Add(mBase, uint32(v89)+24)) = v86
									*(*int64)(unsafe.Add(mBase, uint32(v89)+16)) = v86
									*(*int64)(unsafe.Add(mBase, uint32(v89)+8)) = v86
									*(*int64)(unsafe.Add(mBase, uint32(v89))) = v86
									v98 = int32(32)
									v101 = v90 - v98
									if base.Ui32(int32(31)) < base.Ui32(v101) {
										v89 = v89 + v98
										v90 = v101
										continue
									} else {
										break
									}
									break
								}
							}
						}
					}
				}
			}
		}
	}
	return
}
func F_exprs_known_equal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
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
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v11 == v5 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v14 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v24 = v5
	goto L4
L4:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v27+v24<<(uint(int32(2))%32))))
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v31)+41)))
	if v32 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v133 = v24 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	if v133 < v134 {
		v24 = v133
		goto L4
	} else {
		goto L40
	}
L7:
	;
	if l3 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	v34 = int32(0)
	if v33 == v34 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	goto L10
L10:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v31)+16))
	if v75 == int32(0) {
		goto L6
	} else {
		goto L25
	}
L11:
	;
	if v72 == int32(0) {
		goto L6
	} else {
		goto L24
	}
L12:
	;
	v72 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v33)+4))
	if v40 <= int32(0) {
		v66 = v34
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v72 = v66
	goto L11
L16:
	;
	v43 = int32(0)
	if v43 < v40 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v46 = v40
	goto L19
L18:
	;
	v46 = v43
	goto L19
L19:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
	v49 = int32(0)
	goto L20
L20:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v47+v49<<(uint(int32(2))%32))))
	v58 = base.B2i32(v57 == l3)
	if v57 == l3 {
		v66 = v58
		goto L15
	} else {
		goto L22
	}
L21:
	;
	v66 = v58
	goto L15
L22:
	;
	v60 = v49 + int32(1)
	if v60 != v46 {
		v49 = v60
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	goto L10
L25:
	;
	v78 = int32(0)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v81 <= v78 {
		goto L6
	} else {
		goto L26
	}
L26:
	;
	v84 = v78
	v90 = v78
	v92 = v78
	goto L27
L27:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v94+v84<<(uint(int32(2))%32))))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v100 = F_equal(m, l1, v99)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	return int32(1)
L29:
	;
	if v109&int32(1)&v110 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	return int32(0)
L31:
	;
	if v100 != 0 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v109 = v90
	v110 = int32(1)
	goto L29
L33:
	;
	goto L34
L34:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v98)+4))
	v106 = F_equal(m, l2, v105)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L30
	} else {
		goto L35
	}
L35:
	;
	v109 = v106 | v90
	v110 = v92
	goto L29
L36:
	;
	v117 = v84 + int32(1)
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v118 <= v117 {
		goto L6
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	goto L28
L39:
	;
	v84 = v117
	v90 = v109
	v92 = v110
	goto L27
L40:
	;
	goto L5
}
