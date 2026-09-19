package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_NameOfDatum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v3 != 0 {
		v9 = v3
		return v9
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v5 = F_NameListToString(m, v4)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			v9 = v5
			return v9
		}
	}
}
func F_NonFiniteTimestampTzPart(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) float64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v71 float64
	_ = v71
	var v72 float64
	_ = v72
	v6 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	if base.B2i32(l0 == v6)|base.B2i32(l0 == int32(17)) == v6 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return float64(0)
		} else {
			F_errcode(m, int32(50856066))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return float64(0)
			} else {
				if l4 != 0 {
					v30 = int32(1184)
				} else {
					v30 = int32(1114)
				}
				v31 = F_format_type_be(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return float64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v31
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = l2
					F_errmsg(m, int32(_a_F_NonFiniteTimestampTzPart_0), v10)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return float64(0)
					} else {
						F_errfinish(m, int32(_a_F_NonFiniteTimestampTzPart_1), int32(_a_F_NonFiniteTimestampTzPart_2), int32(_a_F_NonFiniteTimestampTzPart_3))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return float64(0)
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
		switch l1 - int32(4) {
		case 0, 14, 15, 16, 17, 18, 19, 20, 25, 26, 28, 29, 30, 31, 33:
			v72 = float64(0)
			m.G0 = v10 + int32(32)
			return v72
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return float64(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v51 = m.ExcPending
				if v51 != 0 {
					return float64(0)
				} else {
					if l4 != 0 {
						v54 = int32(1184)
					} else {
						v54 = int32(1114)
					}
					v55 = F_format_type_be(m, v54)
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return float64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v55
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l2
						F_errmsg(m, int32(_a_F_NonFiniteTimestampTzPart_4), v10+int32(16))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return float64(0)
						} else {
							F_errfinish(m, int32(_a_F_NonFiniteTimestampTzPart_1), int32(_a_F_NonFiniteTimestampTzPart_5), int32(_a_F_NonFiniteTimestampTzPart_3))
							mBase = m.M
							v68 = m.ExcPending
							if v68 != 0 {
								return float64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			}
		case 7, 21, 22, 23, 24, 27, 32:
			if l3 != 0 {
				v71 = math.Float64frombits(uint64(0xfff0000000000000))
			} else {
				v71 = math.Float64frombits(uint64(0x7ff0000000000000))
			}
			v72 = v71
			m.G0 = v10 + int32(32)
			return v72
		}
	}
}
func F_NotifyMyFrontEnd(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_NotifyMyFrontEnd[0]))
	if v12 == int32(2) {
		v16 = v9 + int32(16)
		F_pq_beginmessage(m, v16, int32(65))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return
		} else {
			F_enlargeStringInfo(m, v16, int32(4))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v9)+20))
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v9)+16))
				v28 = int32(16711935)
				*(*int32)(unsafe.Add(mBase, uint32(v23+v24))) = base.I32_rotr(l2, int32(24))&v28 | base.I32_rotr(l2&v28, int32(8))
				*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v23 + int32(4)
				F_pq_sendstring(m, v16, l0)
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					F_pq_sendstring(m, v16, l1)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return
					} else {
						F_pq_endmessage(m, v16)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return
						} else {
							m.G0 = v9 + int32(32)
							return
						}
					}
				}
			}
		}
	} else {
		v47 = F_errstart(m, int32(17), int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return
		} else {
			if v47 == int32(0) {
				m.G0 = v9 + int32(32)
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = l1
				*(*int32)(unsafe.Add(mBase, uint32(v9))) = l0
				F_errmsg_internal(m, int32(_a_F_NotifyMyFrontEnd_0), v9)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return
				} else {
					F_errfinish(m, int32(_a_F_NotifyMyFrontEnd_1), int32(2377), int32(_a_F_NotifyMyFrontEnd_2))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return
					} else {
						m.G0 = v9 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_namefastcmp_c(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	goto L3
L1:
	;
	return v41 - v42
L3:
	;
	goto L4
L4:
	;
	v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v10 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v11 = l0
	v12 = l1
	v13 = int32(64)
	v14 = v10
	goto L9
L6:
	;
	v37 = l1
	v41 = int32(0)
	goto L7
L7:
	;
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37))))
	goto L1
L8:
	;
	v37 = v32
	v41 = v34
	goto L7
L9:
	;
	v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
	if base.B2i32(v14 != v16)|base.B2i32(v16 == int32(0)) != 0 {
		v32 = v12
		v34 = v14
		goto L8
	} else {
		goto L11
	}
L10:
	;
	v32 = v26
	v34 = int32(0)
	goto L8
L11:
	;
	v22 = v13 - int32(1)
	if v22 == int32(0) {
		v32 = v12
		v34 = v14
		goto L8
	} else {
		goto L12
	}
L12:
	;
	v25 = int32(1)
	v26 = v12 + v25
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+1)))
	if v27 != 0 {
		v11 = v11 + v25
		v12 = v26
		v13 = v22
		v14 = v27
		goto L9
	} else {
		goto L13
	}
L13:
	;
	goto L10
}
func F_nameletext(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(1542), v3, v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return base.B2i32(v6 <= int32(0))
	}
}
func F_namelt(m *base.Module, l0 int32) int32 {
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
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v6 == int32(950) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	return int32(base.Ui32(v61) >> (uint(int32(31)) % 32))
L2:
	;
	goto L7
L3:
	;
	goto L4
L4:
	;
	v55 = F_strlen(m, v5)
	mBase = m.M
	v56 = F_strlen(m, v4)
	mBase = m.M
	v57 = F_varstr_cmp(m, v5, v55, v4, v56, v6)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L19
	}
L5:
	;
	v61 = v46 - v47
	goto L1
L7:
	;
	goto L8
L8:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v15 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v16 = v5
	v17 = v4
	v18 = int32(64)
	v19 = v15
	goto L13
L10:
	;
	v42 = v4
	v46 = int32(0)
	goto L11
L11:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	goto L5
L12:
	;
	v42 = v37
	v46 = v39
	goto L11
L13:
	;
	v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	if base.B2i32(v19 != v21)|base.B2i32(v21 == int32(0)) != 0 {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L15
	}
L14:
	;
	v37 = v31
	v39 = int32(0)
	goto L12
L15:
	;
	v27 = v18 - int32(1)
	if v27 == int32(0) {
		v37 = v17
		v39 = v19
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v30 = int32(1)
	v31 = v17 + v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+1)))
	if v32 != 0 {
		v16 = v16 + v30
		v17 = v31
		v18 = v27
		v19 = v32
		goto L13
	} else {
		goto L17
	}
L17:
	;
	goto L14
L18:
	;
	return int32(0)
L19:
	;
	v61 = v57
	goto L1
}
func F_namerecv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v9)+12))
	v15 = F_pq_getmsgtext(m, v9, v10-v11, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if v19 < int32(64) {
			v23 = F_palloc0(m, int32(64))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				if v25 != 0 {
					base.MemoryCopy(m, v23, v15, v25)
				} else {
				}
				F_pfree(m, v15)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					m.G0 = v7 + int32(16)
					return v23
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34103428))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_namerecv_0), int32(0))
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(64)
						F_errdetail(m, int32(_a_F_namerecv_1), v7)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_namerecv_2), int32(95), int32(_a_F_namerecv_3))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
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
func F_networkjoinsel_semi(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v20 float64
	_ = v20
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 float32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 float64
	_ = v59
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v95 float64
	_ = v95
	var v102 int32
	_ = v102
	var v103 float32
	_ = v103
	var v106 float32
	_ = v106
	var v109 float32
	_ = v109
	var v112 float32
	_ = v112
	var v114 float64
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v143 float64
	_ = v143
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v170 float64
	_ = v170
	var v175 int32
	_ = v175
	var v179 float32
	_ = v179
	var v181 float64
	_ = v181
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v189 int64
	_ = v189
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v229 float64
	_ = v229
	var v231 float64
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 float32
	_ = v238
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 float64
	_ = v259
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v293 float64
	_ = v293
	var v303 int32
	_ = v303
	var v304 float32
	_ = v304
	var v307 float32
	_ = v307
	var v310 float32
	_ = v310
	var v313 float32
	_ = v313
	var v315 float64
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v341 float64
	_ = v341
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v368 float64
	_ = v368
	var v376 int32
	_ = v376
	var v380 float32
	_ = v380
	var v382 float64
	_ = v382
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int64
	_ = v390
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v428 float64
	_ = v428
	var v434 float64
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v452 float64
	_ = v452
	var v455 float64
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v479 int32
	_ = v479
	var v494 float64
	_ = v494
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v510 float32
	_ = v510
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v587 float64
	_ = v587
	var v588 int32
	_ = v588
	var v593 float64
	_ = v593
	var v594 float64
	_ = v594
	var v597 float64
	_ = v597
	var v627 float64
	_ = v627
	var v629 float64
	_ = v629
	var v631 int32
	_ = v631
	var v653 float64
	_ = v653
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v667 float64
	_ = v667
	var v677 int32
	_ = v677
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v701 float64
	_ = v701
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v759 int32
	_ = v759
	var v791 float64
	_ = v791
	var v792 int32
	_ = v792
	var v797 float64
	_ = v797
	var v798 float64
	_ = v798
	var v801 float64
	_ = v801
	var v831 float64
	_ = v831
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v834 float64
	_ = v834
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v867 float64
	_ = v867
	var v878 float64
	_ = v878
	var v879 float64
	_ = v879
	var v889 float64
	_ = v889
	var v893 int32
	_ = v893
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	v5 = int32(0)
	v20 = float64(0)
	v28 = m.G0
	v30 = v28 - int32(192)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	if v32 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v234 != 0 {
		goto L25
	} else {
		goto L26
	}
L2:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+22)))
	v36 = *(*float32)(unsafe.Add(mBase, uint32(v33+v34)+8))
	v42 = F_get_attstatsslot(m, v30+int32(128), v32, int32(1), int32(0), int32(3))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v187 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+160)) = v187
	v189 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+152)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v30)+144)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v30)+136)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v30)+128)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v30)+48)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v30)+56)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v30)+64)) = v189
	*(*int64)(unsafe.Add(mBase, uint32(v30)+72)) = v189
	*(*int32)(unsafe.Add(mBase, uint32(v30)+80)) = v187
	v218 = v5
	v219 = v5
	v223 = v5
	v229 = v20
	v231 = v20
	goto L1
L5:
	;
	return float64(0)
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v52 = F_get_attstatsslot(m, v30+int32(48), v48, int32(2), int32(0), int32(1))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v54 = int32(1024)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v30)+144))
	if v54 <= v55 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v58 = v54
	goto L10
L9:
	;
	v58 = v55
	goto L10
L10:
	;
	v59 = base.F64_promote_f32(v36)
	if v42 == int32(0) {
		v218 = v58
		v219 = v5
		v223 = v52
		v229 = v20
		v231 = v59
		goto L1
	} else {
		goto L11
	}
L11:
	;
	if v55 <= int32(0) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v218 = v58
	v219 = int32(1)
	v223 = v52
	v229 = v20
	v231 = v59
	goto L1
L13:
	;
	goto L14
L14:
	;
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v30)+148))
	v67 = v58 & int32(3)
	v68 = int32(0)
	if v55 < int32(4) {
		v123 = v68
		v143 = v20
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v150 = v123
	v158 = v5
	v170 = v143
	goto L21
L16:
	;
	v75 = v68
	v78 = v5
	v95 = v20
	goto L17
L17:
	;
	v102 = v65 + v75<<(uint(int32(2))%32)
	v103 = *(*float32)(unsafe.Add(mBase, uint32(v102)))
	v106 = *(*float32)(unsafe.Add(mBase, uint32(v102)+4))
	v109 = *(*float32)(unsafe.Add(mBase, uint32(v102)+8))
	v112 = *(*float32)(unsafe.Add(mBase, uint32(v102)+12))
	v114 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v95, base.F64_promote_f32(v103)), base.F64_promote_f32(v106)), base.F64_promote_f32(v109)), base.F64_promote_f32(v112))
	v115 = int32(4)
	v116 = v75 + v115
	v118 = v78 + v115
	if v118 != v58&int32(2044) {
		v75 = v116
		v78 = v118
		v95 = v114
		goto L17
	} else {
		goto L19
	}
L18:
	;
	if v67 != 0 {
		v123 = v116
		v143 = v114
		goto L15
	} else {
		goto L20
	}
L19:
	;
	goto L18
L20:
	;
	v218 = v58
	v219 = int32(1)
	v223 = v52
	v229 = v114
	v231 = v59
	goto L1
L21:
	;
	v175 = int32(1)
	v179 = *(*float32)(unsafe.Add(mBase, uint32(v65+v150<<(uint(int32(2))%32))))
	v181 = base.F64_add(v170, base.F64_promote_f32(v179))
	v185 = v158 + v175
	if v185 != v67 {
		v150 = v150 + v175
		v158 = v185
		v170 = v181
		goto L21
	} else {
		goto L23
	}
L22:
	;
	v218 = v58
	v219 = v175
	v223 = v52
	v229 = v181
	v231 = v59
	goto L1
L23:
	;
	goto L22
L24:
	;
	v436 = F_get_opcode(m, l0)
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L5
	} else {
		goto L46
	}
L25:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v235)+22)))
	v238 = *(*float32)(unsafe.Add(mBase, uint32(v235+v236)+8))
	v244 = F_get_attstatsslot(m, v30+int32(88), v234, int32(1), int32(0), int32(3))
	mBase = m.M
	v245 = m.ExcPending
	if v245 != 0 {
		goto L5
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	v388 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+120)) = v388
	v390 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v30)+112)) = v390
	*(*int64)(unsafe.Add(mBase, uint32(v30)+104)) = v390
	*(*int64)(unsafe.Add(mBase, uint32(v30)+96)) = v390
	*(*int64)(unsafe.Add(mBase, uint32(v30)+88)) = v390
	*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v390
	*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v390
	*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v390
	*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v390
	*(*int32)(unsafe.Add(mBase, uint32(v30)+40)) = v388
	v415 = v388
	v417 = v5
	v422 = v5
	v428 = v20
	v434 = v20
	goto L24
L28:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	v252 = F_get_attstatsslot(m, v30+int32(8), v248, int32(2), int32(0), int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L5
	} else {
		goto L29
	}
L29:
	;
	v254 = int32(1024)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v30)+104))
	if v254 <= v255 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v258 = v254
	goto L32
L31:
	;
	v258 = v255
	goto L32
L32:
	;
	v259 = base.F64_promote_f32(v238)
	if v244 == int32(0) {
		v415 = v258
		v417 = v5
		v422 = v252
		v428 = v20
		v434 = v259
		goto L24
	} else {
		goto L33
	}
L33:
	;
	if v255 <= int32(0) {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v415 = v258
	v417 = int32(1)
	v422 = v252
	v428 = v20
	v434 = v259
	goto L24
L35:
	;
	goto L36
L36:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v30)+108))
	v267 = v258 & int32(3)
	v268 = int32(0)
	if v255 < int32(4) {
		v324 = v268
		v341 = v20
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v351 = v324
	v359 = v268
	v368 = v341
	goto L43
L38:
	;
	v276 = v268
	v282 = v5
	v293 = v20
	goto L39
L39:
	;
	v303 = v265 + v276<<(uint(int32(2))%32)
	v304 = *(*float32)(unsafe.Add(mBase, uint32(v303)))
	v307 = *(*float32)(unsafe.Add(mBase, uint32(v303)+4))
	v310 = *(*float32)(unsafe.Add(mBase, uint32(v303)+8))
	v313 = *(*float32)(unsafe.Add(mBase, uint32(v303)+12))
	v315 = base.F64_add(base.F64_add(base.F64_add(base.F64_add(v293, base.F64_promote_f32(v304)), base.F64_promote_f32(v307)), base.F64_promote_f32(v310)), base.F64_promote_f32(v313))
	v316 = int32(4)
	v317 = v276 + v316
	v319 = v282 + v316
	if v319 != v258&int32(2044) {
		v276 = v317
		v282 = v319
		v293 = v315
		goto L39
	} else {
		goto L41
	}
L40:
	;
	if v267 != 0 {
		v324 = v317
		v341 = v315
		goto L37
	} else {
		goto L42
	}
L41:
	;
	goto L40
L42:
	;
	v415 = v258
	v417 = int32(1)
	v422 = v252
	v428 = v315
	v434 = v259
	goto L24
L43:
	;
	v376 = int32(1)
	v380 = *(*float32)(unsafe.Add(mBase, uint32(v265+v351<<(uint(int32(2))%32))))
	v382 = base.F64_add(v368, base.F64_promote_f32(v380))
	v386 = v359 + v376
	if v386 != v267 {
		v351 = v351 + v376
		v359 = v386
		v368 = v382
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v415 = v258
	v417 = v376
	v422 = v252
	v428 = v382
	v434 = v259
	goto L24
L45:
	;
	goto L44
L46:
	;
	F_fmgr_info(m, v436, v30+int32(164))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L5
	} else {
		goto L47
	}
L47:
	;
	if v422 == int32(0) {
		v455 = float64(0)
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v456 = v417 | v422
	v458 = int32(0)
	if base.B2i32(v219&v456 == v458)|base.B2i32(v218 <= v458) == v458 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v446 = *(*int32)(unsafe.Add(mBase, uint32(l3)+4))
	if v446 == int32(0) {
		v455 = float64(0)
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v452 = *(*float64)(unsafe.Add(mBase, uint32(v446)+16))
	v455 = base.F64_mul(base.F64_sub(base.F64_sub(float64(1), v434), v428), v452)
	goto L48
L51:
	;
	v465 = int32(0)
	v479 = v465
	v494 = v20
	goto L54
L52:
	;
	v653 = v20
	goto L53
L53:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
	if v456&(v223&base.B2i32(int32(2) < v660)) != 0 {
		goto L74
	} else {
		goto L75
	}
L54:
	;
	v502 = v479 << (uint(int32(2)) % 32)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v30)+140))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v502+v503)))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v507 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v30)+148))
	v510 = *(*float32)(unsafe.Add(mBase, uint32(v508+v502)))
	if v417&base.B2i32(v465 < v415) == int32(0) {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	v653 = v629
	goto L53
L56:
	;
	v629 = base.F64_add(base.F64_mul(base.F64_promote_f32(v510), v627), v494)
	v631 = v479 + int32(1)
	if v631 != v218 {
		v479 = v631
		v494 = v629
		goto L54
	} else {
		goto L73
	}
L57:
	;
	if v422&base.F64_gt(v455, float64(0)) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L58:
	;
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v30)+100))
	v518 = int32(0)
	goto L59
L59:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v514+v518<<(uint(int32(2))%32))))
	v550 = F_FunctionCall2Coll(m, v30+int32(164), int32(0), v505, v549)
	mBase = m.M
	v551 = m.ExcPending
	if v551 != 0 {
		goto L5
	} else {
		goto L61
	}
L60:
	;
	v627 = float64(1)
	goto L56
L61:
	;
	if v550 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v555 = v518 + int32(1)
	if v415 != v555 {
		v518 = v555
		goto L59
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	goto L60
L65:
	;
	goto L57
L66:
	;
	v627 = float64(0)
	goto L56
L67:
	;
	v587 = F_inet_hist_value_sel(m, v507, v506, v505, v465-l1)
	mBase = m.M
	v588 = m.ExcPending
	if v588 != 0 {
		goto L5
	} else {
		goto L68
	}
L68:
	;
	if base.F64_gt(v587, float64(0)) == int32(0) {
		goto L66
	} else {
		goto L69
	}
L69:
	;
	v593 = float64(1)
	v594 = base.F64_mul(v455, v587)
	if base.F64_gt(v594, v593) != 0 {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v597 = v593
	goto L72
L71:
	;
	v597 = v594
	goto L72
L72:
	;
	v627 = v597
	goto L56
L73:
	;
	goto L55
L74:
	;
	v665 = int32(0)
	v667 = float64(0)
	v677 = int32(1)
	v687 = v677
	v690 = v665
	v701 = v667
	goto L77
L75:
	;
	v867 = v653
	goto L76
L76:
	;
	if l0 == int32(3552) {
		goto L97
	} else {
		goto L98
	}
L77:
	;
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v30)+60))
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v709+v687<<(uint(int32(2))%32))))
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v30)+24))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v30)+20))
	if v417&base.B2i32(v665 < v415) == int32(0) {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	v867 = base.F64_add(v653, base.F64_div(base.F64_mul(base.F64_sub(base.F64_sub(float64(1), v231), v229), v834), base.F64_convert_i32_s(v833)))
	goto L76
L79:
	;
	v832 = int32(1)
	v833 = v690 + v832
	v834 = base.F64_add(v701, v831)
	v835 = v687 + (int32(base.Ui32(v660-int32(3))>>(uint(int32(10))%32)) + v677)
	v836 = *(*int32)(unsafe.Add(mBase, uint32(v30)+64))
	if v835 < v836-v832 {
		v687 = v835
		v690 = v833
		v701 = v834
		goto L77
	} else {
		goto L96
	}
L80:
	;
	if v422&base.F64_gt(v455, v667) == int32(0) {
		goto L89
	} else {
		goto L90
	}
L81:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v30)+100))
	v722 = int32(0)
	goto L82
L82:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v718+v722<<(uint(int32(2))%32))))
	v754 = F_FunctionCall2Coll(m, v30+int32(164), int32(0), v713, v753)
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L5
	} else {
		goto L84
	}
L83:
	;
	v831 = float64(1)
	goto L79
L84:
	;
	if v754 == int32(0) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v759 = v722 + int32(1)
	if v415 != v759 {
		v722 = v759
		goto L82
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	goto L83
L88:
	;
	goto L80
L89:
	;
	v831 = float64(0)
	goto L79
L90:
	;
	v791 = F_inet_hist_value_sel(m, v715, v714, v713, v665-l1)
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L5
	} else {
		goto L91
	}
L91:
	;
	if base.F64_gt(v791, float64(0)) == int32(0) {
		goto L89
	} else {
		goto L92
	}
L92:
	;
	v797 = float64(1)
	v798 = base.F64_mul(v455, v791)
	if base.F64_gt(v798, v797) != 0 {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	v801 = v797
	goto L95
L94:
	;
	v801 = v798
	goto L95
L95:
	;
	v831 = v801
	goto L79
L96:
	;
	goto L78
L97:
	;
	v878 = float64(0.01)
	goto L99
L98:
	;
	v878 = float64(0.005)
	goto L99
L99:
	;
	v879 = float64(1)
	if (v219|v223)&v456&int32(1) != 0 {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v889 = v867
	goto L102
L101:
	;
	v889 = base.F64_mul(v878, base.F64_mul(base.F64_sub(v879, v231), base.F64_sub(v879, v434)))
	goto L102
L102:
	;
	F_free_attstatsslot(m, v30+int32(128))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L5
	} else {
		goto L103
	}
L103:
	;
	F_free_attstatsslot(m, v30+int32(88))
	mBase = m.M
	v897 = m.ExcPending
	if v897 != 0 {
		goto L5
	} else {
		goto L104
	}
L104:
	;
	F_free_attstatsslot(m, v30+int32(48))
	mBase = m.M
	v901 = m.ExcPending
	if v901 != 0 {
		goto L5
	} else {
		goto L105
	}
L105:
	;
	F_free_attstatsslot(m, v30+int32(8))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L5
	} else {
		goto L106
	}
L106:
	;
	m.G0 = v30 + int32(192)
	return v889
}
func F_newnfa(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v24 int64
	_ = v24
	var v41 int32
	_ = v41
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
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
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
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v231 int32
	_ = v231
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v284 int32
	_ = v284
	var v290 int32
	_ = v290
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v316 int32
	_ = v316
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v341 int32
	_ = v341
	var v352 int32
	_ = v352
	v8 = F_palloc_extended(m, int32(84), int32(2))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	if v8 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v16 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = int32(0)
	v24 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v8)+32)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v24
	*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v24
	*(*int32)(unsafe.Add(mBase, uint32(v8)+76)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+52)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v8)+72)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v8)+56)) = int64(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = l2
	*(*int64)(unsafe.Add(mBase, uint32(v8)+64)) = int64(-4294967296)
	v41 = F_newstate(m, v8)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L9
	}
L6:
	;
	v18 = v16
	goto L8
L7:
	;
	v18 = int32(12)
	goto L8
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v18
	return int32(0)
L9:
	;
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v43 = int32(64)
	*(*uint8)(unsafe.Add(mBase, uint32(v41)+4)) = uint8(v43)
	goto L12
L11:
	;
	goto L12
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v41
	v46 = F_newstate(m, v8)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	if v46 != 0 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v48 = int32(62)
	*(*uint8)(unsafe.Add(mBase, uint32(v46)+4)) = uint8(v48)
	goto L16
L15:
	;
	goto L16
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v46
	v51 = F_newstate(m, v8)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v51
	v54 = F_newstate(m, v8)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v57 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v8)+36))
	if v299 != 0 {
		goto L132
	} else {
		goto L133
	}
L20:
	;
	goto L19
L21:
	;
	goto L22
L22:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	F_rainbow(m, v8, v58, int32(-1), v60, v61)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v67 = *(*int32)(unsafe.Add(mBase, _c_F_newnfa[0]))
	if v67 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v65)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v64)+8))
	if v70 <= v71 {
		goto L30
	} else {
		goto L31
	}
L27:
	;
	goto L26
L28:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v124 = *(*int32)(unsafe.Add(mBase, _c_F_newnfa[0]))
	if v124 != 0 {
		goto L50
	} else {
		goto L51
	}
L29:
	;
	F_createarc(m, v8, int32(94), int32(1), v65, v64)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L49
	}
L30:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v65)+20))
	if v73 == int32(0) {
		goto L29
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v64)+16))
	if v90 == int32(0) {
		goto L29
	} else {
		goto L41
	}
L33:
	;
	v78 = v73
	goto L34
L34:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
	if v81 != v64 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	goto L29
L36:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v78)+16))
	if v89 != 0 {
		v78 = v89
		goto L34
	} else {
		goto L40
	}
L37:
	;
	v83 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v78)+4)))
	if v83 != int32(1) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v78)))
	if v86 == int32(94) {
		goto L28
	} else {
		goto L39
	}
L39:
	;
	goto L36
L40:
	;
	goto L35
L41:
	;
	v95 = v90
	goto L42
L42:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v95)+8))
	if v98 != v65 {
		goto L44
	} else {
		goto L45
	}
L43:
	;
	goto L29
L44:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v95)+24))
	if v106 != 0 {
		v95 = v106
		goto L42
	} else {
		goto L48
	}
L45:
	;
	v100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v95)+4)))
	if v100 != int32(1) {
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v95)))
	if v103 == int32(94) {
		goto L28
	} else {
		goto L47
	}
L47:
	;
	goto L44
L48:
	;
	goto L43
L49:
	;
	goto L28
L50:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	if v127 <= v128 {
		goto L56
	} else {
		goto L57
	}
L53:
	;
	goto L52
L54:
	;
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v8)+52))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	F_rainbow(m, v8, v174, int32(-1), v176, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L1
	} else {
		goto L76
	}
L55:
	;
	F_createarc(m, v8, int32(94), int32(0), v122, v121)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L75
	}
L56:
	;
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v122)+20))
	if v130 == int32(0) {
		goto L55
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v121)+16))
	if v145 == int32(0) {
		goto L55
	} else {
		goto L67
	}
L59:
	;
	v135 = v130
	goto L60
L60:
	;
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v135)+12))
	if v138 != v121 {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	goto L55
L62:
	;
	v144 = *(*int32)(unsafe.Add(mBase, uint32(v135)+16))
	if v144 != 0 {
		v135 = v144
		goto L60
	} else {
		goto L66
	}
L63:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+4)))
	if v140 != 0 {
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v135)))
	if v141 == int32(94) {
		goto L54
	} else {
		goto L65
	}
L65:
	;
	goto L62
L66:
	;
	goto L61
L67:
	;
	v150 = v145
	goto L68
L68:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v150)+8))
	if v153 != v122 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L55
L70:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v150)+24))
	if v159 != 0 {
		v150 = v159
		goto L68
	} else {
		goto L74
	}
L71:
	;
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v150)+4)))
	if v155 != 0 {
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	if v156 == int32(94) {
		goto L54
	} else {
		goto L73
	}
L73:
	;
	goto L70
L74:
	;
	goto L69
L75:
	;
	goto L54
L76:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v181 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v183 = *(*int32)(unsafe.Add(mBase, _c_F_newnfa[0]))
	if v183 != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v180)+8))
	if v186 <= v187 {
		goto L83
	} else {
		goto L84
	}
L80:
	;
	goto L79
L81:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v8)+8))
	v240 = *(*int32)(unsafe.Add(mBase, _c_F_newnfa[0]))
	if v240 != 0 {
		goto L103
	} else {
		goto L104
	}
L82:
	;
	F_createarc(m, v8, int32(36), int32(1), v181, v180)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L1
	} else {
		goto L102
	}
L83:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
	if v189 == int32(0) {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v180)+16))
	if v206 == int32(0) {
		goto L82
	} else {
		goto L94
	}
L86:
	;
	v194 = v189
	goto L87
L87:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v194)+12))
	if v197 != v180 {
		goto L89
	} else {
		goto L90
	}
L88:
	;
	goto L82
L89:
	;
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	if v205 != 0 {
		v194 = v205
		goto L87
	} else {
		goto L93
	}
L90:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+4)))
	if v199 != int32(1) {
		goto L89
	} else {
		goto L91
	}
L91:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
	if v202 == int32(36) {
		goto L81
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	goto L88
L94:
	;
	v211 = v206
	goto L95
L95:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v211)+8))
	if v214 != v181 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	goto L82
L97:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v211)+24))
	if v222 != 0 {
		v211 = v222
		goto L95
	} else {
		goto L101
	}
L98:
	;
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v211)+4)))
	if v216 != int32(1) {
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
	if v219 == int32(36) {
		goto L81
	} else {
		goto L100
	}
L100:
	;
	goto L97
L101:
	;
	goto L96
L102:
	;
	goto L81
L103:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v237)+8))
	if v243 <= v244 {
		goto L109
	} else {
		goto L110
	}
L106:
	;
	goto L105
L107:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v290 == int32(0) {
		goto L129
	} else {
		goto L130
	}
L108:
	;
	F_createarc(m, v8, int32(36), int32(0), v238, v237)
	mBase = m.M
	v284 = m.ExcPending
	if v284 != 0 {
		goto L1
	} else {
		goto L128
	}
L109:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v238)+20))
	if v246 == int32(0) {
		goto L108
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v237)+16))
	if v261 == int32(0) {
		goto L108
	} else {
		goto L120
	}
L112:
	;
	v251 = v246
	goto L113
L113:
	;
	v254 = *(*int32)(unsafe.Add(mBase, uint32(v251)+12))
	if v254 != v237 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	goto L108
L115:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v251)+16))
	if v260 != 0 {
		v251 = v260
		goto L113
	} else {
		goto L119
	}
L116:
	;
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v251)+4)))
	if v256 != 0 {
		goto L115
	} else {
		goto L117
	}
L117:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v251)))
	if v257 == int32(36) {
		goto L107
	} else {
		goto L118
	}
L118:
	;
	goto L115
L119:
	;
	goto L114
L120:
	;
	v266 = v261
	goto L121
L121:
	;
	v269 = *(*int32)(unsafe.Add(mBase, uint32(v266)+8))
	if v269 != v238 {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L108
L123:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v266)+24))
	if v275 != 0 {
		v266 = v275
		goto L121
	} else {
		goto L127
	}
L124:
	;
	v271 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v266)+4)))
	if v271 != 0 {
		goto L123
	} else {
		goto L125
	}
L125:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v266)))
	if v272 == int32(36) {
		goto L107
	} else {
		goto L126
	}
L126:
	;
	goto L123
L127:
	;
	goto L122
L128:
	;
	goto L107
L129:
	;
	return v8
L130:
	;
	goto L131
L131:
	;
	goto L19
L132:
	;
	v302 = v299
	goto L135
L133:
	;
	goto L134
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = int32(0)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v8)+40))
	if v324 != 0 {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v302)))
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+136))
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v302)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v306)+136)) = v307 + v308*int32(-36) - int32(8)
	F_pfree(m, v302)
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L1
	} else {
		goto L137
	}
L136:
	;
	goto L134
L137:
	;
	if v305 != 0 {
		v302 = v305
		goto L135
	} else {
		goto L138
	}
L138:
	;
	goto L136
L139:
	;
	v327 = v324
	goto L142
L140:
	;
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = int32(0)
	F_pfree(m, v8)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L1
	} else {
		goto L146
	}
L142:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v8)+76))
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v331)+136))
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v331)+136)) = v332 + v333*int32(-40) - int32(8)
	F_pfree(m, v327)
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L1
	} else {
		goto L144
	}
L143:
	;
	goto L141
L144:
	;
	if v330 != 0 {
		v327 = v330
		goto L142
	} else {
		goto L145
	}
L145:
	;
	goto L143
L146:
	;
	return int32(0)
}
func F_nextval_oid(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_nextval_internal(m, v2, int32(1))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = F_Int64GetDatum(m, v4)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_nlikesel(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13991(m, l0, int32(0))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_nonword(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
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
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_nonword[0]))
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	if l1 == int32(97) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	return
L5:
	;
	goto L3
L6:
	;
	v17 = int32(36)
	goto L8
L7:
	;
	v17 = int32(94)
	goto L8
L8:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v18 <= v19 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v76 = *(*int32)(unsafe.Add(mBase, _c_F_nonword[0]))
	if v76 != 0 {
		goto L31
	} else {
		goto L32
	}
L10:
	;
	F_createarc(m, v8, v17, int32(1), l2, l3)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L4
	} else {
		goto L30
	}
L11:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v21 == int32(0) {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v39 == int32(0) {
		goto L10
	} else {
		goto L22
	}
L14:
	;
	v28 = v21
	goto L15
L15:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	if v31 != l3 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L10
L17:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v28)+16))
	if v38 != 0 {
		v28 = v38
		goto L15
	} else {
		goto L21
	}
L18:
	;
	v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	if v33 != int32(1) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v36 == v17 {
		goto L9
	} else {
		goto L20
	}
L20:
	;
	goto L17
L21:
	;
	goto L16
L22:
	;
	v46 = v39
	goto L23
L23:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v46)+8))
	if v49 != l2 {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	goto L10
L25:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v46)+24))
	if v56 != 0 {
		v46 = v56
		goto L23
	} else {
		goto L29
	}
L26:
	;
	v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)))
	if v51 != int32(1) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v46)))
	if v54 == v17 {
		goto L9
	} else {
		goto L28
	}
L28:
	;
	goto L25
L29:
	;
	goto L24
L30:
	;
	goto L9
L31:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l3)+8))
	if v79 <= v80 {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	goto L33
L35:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+92))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+100))
	F_colorcomplement(m, v131, v132, l1, v133, l2, l3)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L4
	} else {
		goto L57
	}
L36:
	;
	F_createarc(m, v74, v17, int32(0), l2, l3)
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L4
	} else {
		goto L56
	}
L37:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	if v82 == int32(0) {
		goto L36
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v98 == int32(0) {
		goto L36
	} else {
		goto L48
	}
L40:
	;
	v89 = v82
	goto L41
L41:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	if v92 != l3 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L36
L43:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	if v97 != 0 {
		v89 = v97
		goto L41
	} else {
		goto L47
	}
L44:
	;
	v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v89)+4)))
	if v94 != 0 {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	if v95 == v17 {
		goto L35
	} else {
		goto L46
	}
L46:
	;
	goto L43
L47:
	;
	goto L42
L48:
	;
	v105 = v98
	goto L49
L49:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v105)+8))
	if v108 != l2 {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L36
L51:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v105)+24))
	if v113 != 0 {
		v105 = v113
		goto L49
	} else {
		goto L55
	}
L52:
	;
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+4)))
	if v110 != 0 {
		goto L51
	} else {
		goto L53
	}
L53:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
	if v111 == v17 {
		goto L35
	} else {
		goto L54
	}
L54:
	;
	goto L51
L55:
	;
	goto L50
L56:
	;
	goto L35
L57:
	;
	return
}
func F_notification_match(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v81 int32
	_ = v81
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v6 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5))))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7))))
	if v6 != v8 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(1)
L2:
	;
	v10 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+2)))
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+2)))
	if v10 != v11 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v13 = int32(4)
	v14 = v5 + v13
	v16 = v7 + v13
	v19 = v6 + v10 + int32(2)
	if base.Ui32(v13) <= base.Ui32(v19) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if v81 != 0 {
		goto L1
	} else {
		goto L22
	}
L5:
	;
	v81 = int32(0)
	goto L4
L6:
	;
	v55 = v50
	v56 = v51
	v57 = v52
	goto L16
L7:
	;
	if (v14|v16)&int32(3) != 0 {
		v50 = v14
		v51 = v16
		v52 = v19
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v43 = v14
	v44 = v16
	v45 = v19
	goto L9
L9:
	;
	if v45 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L10:
	;
	v27 = v14
	v28 = v16
	v29 = v19
	goto L11
L11:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	if v32 != v33 {
		v50 = v27
		v51 = v28
		v52 = v29
		goto L6
	} else {
		goto L13
	}
L12:
	;
	v43 = v38
	v44 = v36
	v45 = v40
	goto L9
L13:
	;
	v35 = int32(4)
	v36 = v28 + v35
	v38 = v27 + v35
	v40 = v29 - v35
	if base.Ui32(int32(3)) < base.Ui32(v40) {
		v27 = v38
		v28 = v36
		v29 = v40
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v50 = v43
	v51 = v44
	v52 = v45
	goto L6
L16:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 == v61 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v81 = v60 - v61
	goto L4
L18:
	;
	v63 = int32(1)
	v68 = v57 - v63
	if v68 != 0 {
		v55 = v55 + v63
		v56 = v56 + v63
		v57 = v68
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	goto L5
L22:
	;
	return int32(0)
}
