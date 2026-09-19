package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_CheckDateTokenTable(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = F_strlen(m, l1)
	mBase = m.M
	if base.Ui32(int32(10)) < base.Ui32(v13) {
		v89 = l1
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v120
L2:
	;
	v96 = int32(0)
	v99 = F_errstart(m, int32(15), v96)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L18
	} else {
		goto L24
	}
L3:
	;
	v16 = int32(1)
	v23 = v16
	v24 = v16
	goto L4
L4:
	;
	v28 = l1 + v24<<(uint(int32(4))%32)
	v29 = F_strlen(m, v28)
	mBase = m.M
	if base.Ui32(int32(11)) <= base.Ui32(v29) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v120 = v84
	goto L1
L6:
	;
	v89 = v28
	goto L2
L7:
	;
	goto L8
L8:
	;
	v33 = v28 - int32(16)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	if base.B2i32(v36 == int32(0))|base.B2i32(v36 != v39) != 0 {
		v57 = v36
		v58 = v39
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v86 = v24 + int32(1)
	if v86 != l2 {
		v23 = v84
		v24 = v86
		goto L4
	} else {
		goto L23
	}
L10:
	;
	if v57-v58 < int32(0) {
		v84 = v23
		goto L9
	} else {
		goto L17
	}
L11:
	;
	goto L10
L12:
	;
	v42 = v33
	v43 = v28
	goto L13
L13:
	;
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	if v47 == int32(0) {
		v57 = v47
		v58 = v46
		goto L11
	} else {
		goto L15
	}
L14:
	;
	v57 = v47
	v58 = v46
	goto L11
L15:
	;
	v50 = int32(1)
	if v47 == v46 {
		v42 = v42 + v50
		v43 = v43 + v50
		goto L13
	} else {
		goto L16
	}
L16:
	;
	goto L14
L17:
	;
	v62 = int32(0)
	v65 = F_errstart(m, int32(15), v62)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	return int32(0)
L19:
	;
	if v65 == int32(0) {
		v84 = v62
		goto L9
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	F_errmsg_internal(m, int32(_a_F_CheckDateTokenTable_0), v11+int32(16))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	F_errfinish(m, int32(_a_F_CheckDateTokenTable_1), int32(_a_F_CheckDateTokenTable_2), int32(_a_F_CheckDateTokenTable_3))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L22
	}
L22:
	;
	v84 = v62
	goto L9
L23:
	;
	goto L5
L24:
	;
	if v99 == int32(0) {
		v120 = v96
		goto L1
	} else {
		goto L25
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(_a_F_CheckDateTokenTable_4), v11)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(_a_F_CheckDateTokenTable_1), int32(_a_F_CheckDateTokenTable_5), int32(_a_F_CheckDateTokenTable_3))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v120 = v96
	goto L1
}
func F_date_dist(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = F_DirectFunctionCall2Coll(m, int32(2427), int32(0), v4, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v11 = v6 >> (uint(int32(31)) % 32)
		return v6 ^ v11 - v11
	}
}
func F_date_ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	return base.B2i32(v3 <= v2)
}
func F_date_gt_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(-2147483648) {
		v17 = int64(-9223372036854775807 - 1)
		v26 = base.B2i32(v4 < v17) - base.B2i32(v17 < v4)
	} else {
		if v6 == int32(2147483647) {
			v17 = int64(9223372036854775807)
			v26 = base.B2i32(v4 < v17) - base.B2i32(v17 < v4)
		} else {
			if int32(106751982) < v6 {
				if v4 == int64(9223372036854775807) {
					v25 = int32(-1)
				} else {
					v25 = int32(1)
				}
				v26 = v25
			} else {
				v17 = base.I64_extend_i32_s(v6) * int64(86400000000)
				v26 = base.B2i32(v4 < v17) - base.B2i32(v17 < v4)
			}
		}
	}
	return base.B2i32(int32(0) < v26)
}
func F_date_mi_interval(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = Fn13889(m, l0, int32(1268))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_date_ne_timestamp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v6 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) != int32(0))
	} else {
		if v6 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) != int32(0))
		} else {
			if int32(106751982) < v6 {
				v24 = int32(1)
			} else {
				v18 = base.I64_extend_i32_s(v6) * int64(86400000000)
				v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) != int32(0))
			}
		}
	}
	return v24
}
func F_date_ne_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(_a_F_date_ne_timestamptz_0), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_date_ne_timestamptz[0]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return base.B2i32(v65 != int32(0))
}
func F_date_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pq_getmsgint(m, v2, int32(4))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if base.B2i32(base.Ui32(v4-int32(2147483647)) < base.Ui32(int32(2)))|base.B2i32(base.Ui32(v4+int32(_a_F_date_recv_0)) < base.Ui32(int32(2147483494))) == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_date_recv_1), int32(0))
					mBase = m.M
					v29 = m.ExcPending
					if v29 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_date_recv_2), int32(223), int32(_a_F_date_recv_3))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
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
			return v4
		}
	}
}
func F_date_send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_enlargeStringInfo(m, v6, int32(4))
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v21 = int32(16711935)
			*(*int32)(unsafe.Add(mBase, uint32(v16+v17))) = base.I32_rotr(v8, int32(24))&v21 | base.I32_rotr(v8&v21, int32(8))
			v30 = v16 + int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v30
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v33))) = v30 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return v33
		}
	}
}
func F_make_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v151 int32
	_ = v151
	var v161 int32
	_ = v161
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v270 int32
	_ = v270
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 int64
	_ = v285
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v308 int64
	_ = v308
	var v316 int32
	_ = v316
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int64
	_ = v331
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	v6 = m.G0
	v8 = v6 - int32(112)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+84)) = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+80)) = v14
	if v10 < int32(0) {
		v18 = int32(0)
		v19 = v18 - v10
		if base.B2i32(v19 < v18) != base.B2i32(v18 < v10) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v261 = m.ExcPending
			if v261 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v264 = m.ExcPending
				if v264 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					F_errmsg(m, int32(_a_F_make_date_0), v8)
					mBase = m.M
					v270 = m.ExcPending
					if v270 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_make_date_1), int32(267), int32(_a_F_make_date_2))
						mBase = m.M
						v275 = m.ExcPending
						if v275 != 0 {
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
			*(*int32)(unsafe.Add(mBase, uint32(v8)+88)) = v19
			v33 = v8 + int32(68)
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
			if int32(base.Ui32(v10)>>(uint(int32(31))%32)) != 0 {
				if int32(0) < v39 {
					*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = int32(1) - v39
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
					if base.Ui32(int32(-12)) <= base.Ui32(v151-int32(13)) {
						v161 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						if base.Ui32(int32(-31)) <= base.Ui32(v161-int32(32)) {
							v171 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
							v173 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
							if v173&int32(3) != 0 {
								v183 = int32(0)
							} else {
								v178 = base.I32_rem_s(v173, int32(100))
								if v178 != 0 {
									v183 = int32(1)
								} else {
									v180 = base.I32_rem_s(v173, int32(400))
									v183 = base.B2i32(v180 == int32(0))
								}
							}
							v186 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
							v192 = *(*int32)(unsafe.Add(mBase, uint32(v183*int32(52)+v186<<(uint(int32(2))%32))+uint32(_c_F_make_date[0])))
							if v171 <= v192 {
								v201 = int32(0)
							} else {
								v201 = int32(-2)
							}
						} else {
							v201 = int32(-3)
						}
					} else {
						v201 = int32(-3)
					}
				} else {
					v201 = int32(-2)
				}
			} else {
				if int32(0) < v39 {
					v151 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
					if base.Ui32(int32(-12)) <= base.Ui32(v151-int32(13)) {
						v161 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						if base.Ui32(int32(-31)) <= base.Ui32(v161-int32(32)) {
							v171 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
							v173 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
							if v173&int32(3) != 0 {
								v183 = int32(0)
							} else {
								v178 = base.I32_rem_s(v173, int32(100))
								if v178 != 0 {
									v183 = int32(1)
								} else {
									v180 = base.I32_rem_s(v173, int32(400))
									v183 = base.B2i32(v180 == int32(0))
								}
							}
							v186 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
							v192 = *(*int32)(unsafe.Add(mBase, uint32(v183*int32(52)+v186<<(uint(int32(2))%32))+uint32(_c_F_make_date[0])))
							if v171 <= v192 {
								v201 = int32(0)
							} else {
								v201 = int32(-2)
							}
						} else {
							v201 = int32(-3)
						}
					} else {
						v201 = int32(-3)
					}
				} else {
					v201 = int32(-2)
				}
			}
			if v201 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v279 = m.ExcPending
				if v279 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v282 = m.ExcPending
					if v282 != 0 {
						return int32(0)
					} else {
						v283 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v283
						v285 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v8)+52)) = base.I64_rotl(v285, int64(32))
						F_errmsg(m, int32(_a_F_make_date_0), v8+int32(48))
						mBase = m.M
						v293 = m.ExcPending
						if v293 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_make_date_1), int32(277), int32(_a_F_make_date_2))
							mBase = m.M
							v298 = m.ExcPending
							if v298 != 0 {
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
				v202 = *(*int32)(unsafe.Add(mBase, uint32(v8)+84))
				v203 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
				if v203 <= int32(-4713) {
					if v203 != int32(-4713) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v328 = m.ExcPending
							if v328 != 0 {
								return int32(0)
							} else {
								v329 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v329
								v331 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v331, int64(32))
								F_errmsg(m, int32(_a_F_make_date_3), v8+int32(32))
								mBase = m.M
								v339 = m.ExcPending
								if v339 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_make_date_1), int32(284), int32(_a_F_make_date_2))
									mBase = m.M
									v344 = m.ExcPending
									if v344 != 0 {
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
						if int32(10) < v202 {
							v217 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							v222 = base.B2i32(int32(2) < v202)
							if int32(2) < v202 {
								v223 = int32(_a_F_make_date_4)
							} else {
								v223 = int32(_a_F_make_date_5)
							}
							v224 = v223 + v203
							v229 = base.I32_div_s(v224, int32(4))
							v232 = base.I32_div_s(v224, int32(-100))
							v235 = base.I32_div_s(v224, int32(400))
							if int32(2) < v202 {
								v239 = int32(1)
							} else {
								v239 = int32(13)
							}
							v244 = base.I32_div_s((v239+v202)*int32(_a_F_make_date_6), int32(256))
							v247 = v217 + v224*int32(365) + v229 + v232 + v235 + v244 - int32(_a_F_make_date_7)
							if base.Ui32(int32(2147483494)) <= base.Ui32(v247) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v302 = m.ExcPending
								if v302 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v305 = m.ExcPending
									if v305 != 0 {
										return int32(0)
									} else {
										v306 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v306
										v308 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v308, int64(32))
										F_errmsg(m, int32(_a_F_make_date_3), v8+int32(16))
										mBase = m.M
										v316 = m.ExcPending
										if v316 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_date_1), int32(293), int32(_a_F_make_date_2))
											mBase = m.M
											v321 = m.ExcPending
											if v321 != 0 {
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
								m.G0 = v8 + int32(112)
								return v247 - int32(_a_F_make_date_8)
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v325 = m.ExcPending
							if v325 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v328 = m.ExcPending
								if v328 != 0 {
									return int32(0)
								} else {
									v329 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v329
									v331 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v331, int64(32))
									F_errmsg(m, int32(_a_F_make_date_3), v8+int32(32))
									mBase = m.M
									v339 = m.ExcPending
									if v339 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_date_1), int32(284), int32(_a_F_make_date_2))
										mBase = m.M
										v344 = m.ExcPending
										if v344 != 0 {
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
					if v203 < int32(_a_F_make_date_9) {
						v217 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						v222 = base.B2i32(int32(2) < v202)
						if int32(2) < v202 {
							v223 = int32(_a_F_make_date_4)
						} else {
							v223 = int32(_a_F_make_date_5)
						}
						v224 = v223 + v203
						v229 = base.I32_div_s(v224, int32(4))
						v232 = base.I32_div_s(v224, int32(-100))
						v235 = base.I32_div_s(v224, int32(400))
						if int32(2) < v202 {
							v239 = int32(1)
						} else {
							v239 = int32(13)
						}
						v244 = base.I32_div_s((v239+v202)*int32(_a_F_make_date_6), int32(256))
						v247 = v217 + v224*int32(365) + v229 + v232 + v235 + v244 - int32(_a_F_make_date_7)
						if base.Ui32(int32(2147483494)) <= base.Ui32(v247) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v302 = m.ExcPending
							if v302 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v305 = m.ExcPending
								if v305 != 0 {
									return int32(0)
								} else {
									v306 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v306
									v308 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v308, int64(32))
									F_errmsg(m, int32(_a_F_make_date_3), v8+int32(16))
									mBase = m.M
									v316 = m.ExcPending
									if v316 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_date_1), int32(293), int32(_a_F_make_date_2))
										mBase = m.M
										v321 = m.ExcPending
										if v321 != 0 {
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
							m.G0 = v8 + int32(112)
							return v247 - int32(_a_F_make_date_8)
						}
					} else {
						if base.B2i32(v203 != int32(_a_F_make_date_9))|base.B2i32(int32(6) <= v202) != 0 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v325 = m.ExcPending
							if v325 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v328 = m.ExcPending
								if v328 != 0 {
									return int32(0)
								} else {
									v329 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v329
									v331 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v331, int64(32))
									F_errmsg(m, int32(_a_F_make_date_3), v8+int32(32))
									mBase = m.M
									v339 = m.ExcPending
									if v339 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_date_1), int32(284), int32(_a_F_make_date_2))
										mBase = m.M
										v344 = m.ExcPending
										if v344 != 0 {
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
							v217 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							v222 = base.B2i32(int32(2) < v202)
							if int32(2) < v202 {
								v223 = int32(_a_F_make_date_4)
							} else {
								v223 = int32(_a_F_make_date_5)
							}
							v224 = v223 + v203
							v229 = base.I32_div_s(v224, int32(4))
							v232 = base.I32_div_s(v224, int32(-100))
							v235 = base.I32_div_s(v224, int32(400))
							if int32(2) < v202 {
								v239 = int32(1)
							} else {
								v239 = int32(13)
							}
							v244 = base.I32_div_s((v239+v202)*int32(_a_F_make_date_6), int32(256))
							v247 = v217 + v224*int32(365) + v229 + v232 + v235 + v244 - int32(_a_F_make_date_7)
							if base.Ui32(int32(2147483494)) <= base.Ui32(v247) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v302 = m.ExcPending
								if v302 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v305 = m.ExcPending
									if v305 != 0 {
										return int32(0)
									} else {
										v306 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v306
										v308 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v308, int64(32))
										F_errmsg(m, int32(_a_F_make_date_3), v8+int32(16))
										mBase = m.M
										v316 = m.ExcPending
										if v316 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_make_date_1), int32(293), int32(_a_F_make_date_2))
											mBase = m.M
											v321 = m.ExcPending
											if v321 != 0 {
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
								m.G0 = v8 + int32(112)
								return v247 - int32(_a_F_make_date_8)
							}
						}
					}
				}
			}
		}
	} else {
		v33 = v8 + int32(68)
		v39 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
		if int32(base.Ui32(v10)>>(uint(int32(31))%32)) != 0 {
			if int32(0) < v39 {
				*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = int32(1) - v39
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
				if base.Ui32(int32(-12)) <= base.Ui32(v151-int32(13)) {
					v161 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					if base.Ui32(int32(-31)) <= base.Ui32(v161-int32(32)) {
						v171 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						v173 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
						if v173&int32(3) != 0 {
							v183 = int32(0)
						} else {
							v178 = base.I32_rem_s(v173, int32(100))
							if v178 != 0 {
								v183 = int32(1)
							} else {
								v180 = base.I32_rem_s(v173, int32(400))
								v183 = base.B2i32(v180 == int32(0))
							}
						}
						v186 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
						v192 = *(*int32)(unsafe.Add(mBase, uint32(v183*int32(52)+v186<<(uint(int32(2))%32))+uint32(_c_F_make_date[0])))
						if v171 <= v192 {
							v201 = int32(0)
						} else {
							v201 = int32(-2)
						}
					} else {
						v201 = int32(-3)
					}
				} else {
					v201 = int32(-3)
				}
			} else {
				v201 = int32(-2)
			}
		} else {
			if int32(0) < v39 {
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
				if base.Ui32(int32(-12)) <= base.Ui32(v151-int32(13)) {
					v161 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					if base.Ui32(int32(-31)) <= base.Ui32(v161-int32(32)) {
						v171 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						v173 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
						if v173&int32(3) != 0 {
							v183 = int32(0)
						} else {
							v178 = base.I32_rem_s(v173, int32(100))
							if v178 != 0 {
								v183 = int32(1)
							} else {
								v180 = base.I32_rem_s(v173, int32(400))
								v183 = base.B2i32(v180 == int32(0))
							}
						}
						v186 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
						v192 = *(*int32)(unsafe.Add(mBase, uint32(v183*int32(52)+v186<<(uint(int32(2))%32))+uint32(_c_F_make_date[0])))
						if v171 <= v192 {
							v201 = int32(0)
						} else {
							v201 = int32(-2)
						}
					} else {
						v201 = int32(-3)
					}
				} else {
					v201 = int32(-3)
				}
			} else {
				v201 = int32(-2)
			}
		}
		if v201 != 0 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v279 = m.ExcPending
			if v279 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v282 = m.ExcPending
				if v282 != 0 {
					return int32(0)
				} else {
					v283 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v283
					v285 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+52)) = base.I64_rotl(v285, int64(32))
					F_errmsg(m, int32(_a_F_make_date_0), v8+int32(48))
					mBase = m.M
					v293 = m.ExcPending
					if v293 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_make_date_1), int32(277), int32(_a_F_make_date_2))
						mBase = m.M
						v298 = m.ExcPending
						if v298 != 0 {
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
			v202 = *(*int32)(unsafe.Add(mBase, uint32(v8)+84))
			v203 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
			if v203 <= int32(-4713) {
				if v203 != int32(-4713) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v325 = m.ExcPending
					if v325 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v328 = m.ExcPending
						if v328 != 0 {
							return int32(0)
						} else {
							v329 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v329
							v331 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v331, int64(32))
							F_errmsg(m, int32(_a_F_make_date_3), v8+int32(32))
							mBase = m.M
							v339 = m.ExcPending
							if v339 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_make_date_1), int32(284), int32(_a_F_make_date_2))
								mBase = m.M
								v344 = m.ExcPending
								if v344 != 0 {
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
					if int32(10) < v202 {
						v217 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						v222 = base.B2i32(int32(2) < v202)
						if int32(2) < v202 {
							v223 = int32(_a_F_make_date_4)
						} else {
							v223 = int32(_a_F_make_date_5)
						}
						v224 = v223 + v203
						v229 = base.I32_div_s(v224, int32(4))
						v232 = base.I32_div_s(v224, int32(-100))
						v235 = base.I32_div_s(v224, int32(400))
						if int32(2) < v202 {
							v239 = int32(1)
						} else {
							v239 = int32(13)
						}
						v244 = base.I32_div_s((v239+v202)*int32(_a_F_make_date_6), int32(256))
						v247 = v217 + v224*int32(365) + v229 + v232 + v235 + v244 - int32(_a_F_make_date_7)
						if base.Ui32(int32(2147483494)) <= base.Ui32(v247) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v302 = m.ExcPending
							if v302 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v305 = m.ExcPending
								if v305 != 0 {
									return int32(0)
								} else {
									v306 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v306
									v308 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v308, int64(32))
									F_errmsg(m, int32(_a_F_make_date_3), v8+int32(16))
									mBase = m.M
									v316 = m.ExcPending
									if v316 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_date_1), int32(293), int32(_a_F_make_date_2))
										mBase = m.M
										v321 = m.ExcPending
										if v321 != 0 {
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
							m.G0 = v8 + int32(112)
							return v247 - int32(_a_F_make_date_8)
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v328 = m.ExcPending
							if v328 != 0 {
								return int32(0)
							} else {
								v329 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v329
								v331 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v331, int64(32))
								F_errmsg(m, int32(_a_F_make_date_3), v8+int32(32))
								mBase = m.M
								v339 = m.ExcPending
								if v339 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_make_date_1), int32(284), int32(_a_F_make_date_2))
									mBase = m.M
									v344 = m.ExcPending
									if v344 != 0 {
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
				if v203 < int32(_a_F_make_date_9) {
					v217 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					v222 = base.B2i32(int32(2) < v202)
					if int32(2) < v202 {
						v223 = int32(_a_F_make_date_4)
					} else {
						v223 = int32(_a_F_make_date_5)
					}
					v224 = v223 + v203
					v229 = base.I32_div_s(v224, int32(4))
					v232 = base.I32_div_s(v224, int32(-100))
					v235 = base.I32_div_s(v224, int32(400))
					if int32(2) < v202 {
						v239 = int32(1)
					} else {
						v239 = int32(13)
					}
					v244 = base.I32_div_s((v239+v202)*int32(_a_F_make_date_6), int32(256))
					v247 = v217 + v224*int32(365) + v229 + v232 + v235 + v244 - int32(_a_F_make_date_7)
					if base.Ui32(int32(2147483494)) <= base.Ui32(v247) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v302 = m.ExcPending
						if v302 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v305 = m.ExcPending
							if v305 != 0 {
								return int32(0)
							} else {
								v306 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v306
								v308 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v308, int64(32))
								F_errmsg(m, int32(_a_F_make_date_3), v8+int32(16))
								mBase = m.M
								v316 = m.ExcPending
								if v316 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_make_date_1), int32(293), int32(_a_F_make_date_2))
									mBase = m.M
									v321 = m.ExcPending
									if v321 != 0 {
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
						m.G0 = v8 + int32(112)
						return v247 - int32(_a_F_make_date_8)
					}
				} else {
					if base.B2i32(v203 != int32(_a_F_make_date_9))|base.B2i32(int32(6) <= v202) != 0 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v325 = m.ExcPending
						if v325 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v328 = m.ExcPending
							if v328 != 0 {
								return int32(0)
							} else {
								v329 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v329
								v331 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v331, int64(32))
								F_errmsg(m, int32(_a_F_make_date_3), v8+int32(32))
								mBase = m.M
								v339 = m.ExcPending
								if v339 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_make_date_1), int32(284), int32(_a_F_make_date_2))
									mBase = m.M
									v344 = m.ExcPending
									if v344 != 0 {
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
						v217 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						v222 = base.B2i32(int32(2) < v202)
						if int32(2) < v202 {
							v223 = int32(_a_F_make_date_4)
						} else {
							v223 = int32(_a_F_make_date_5)
						}
						v224 = v223 + v203
						v229 = base.I32_div_s(v224, int32(4))
						v232 = base.I32_div_s(v224, int32(-100))
						v235 = base.I32_div_s(v224, int32(400))
						if int32(2) < v202 {
							v239 = int32(1)
						} else {
							v239 = int32(13)
						}
						v244 = base.I32_div_s((v239+v202)*int32(_a_F_make_date_6), int32(256))
						v247 = v217 + v224*int32(365) + v229 + v232 + v235 + v244 - int32(_a_F_make_date_7)
						if base.Ui32(int32(2147483494)) <= base.Ui32(v247) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v302 = m.ExcPending
							if v302 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v305 = m.ExcPending
								if v305 != 0 {
									return int32(0)
								} else {
									v306 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v306
									v308 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v308, int64(32))
									F_errmsg(m, int32(_a_F_make_date_3), v8+int32(16))
									mBase = m.M
									v316 = m.ExcPending
									if v316 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_make_date_1), int32(293), int32(_a_F_make_date_2))
										mBase = m.M
										v321 = m.ExcPending
										if v321 != 0 {
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
							m.G0 = v8 + int32(112)
							return v247 - int32(_a_F_make_date_8)
						}
					}
				}
			}
		}
	}
}
