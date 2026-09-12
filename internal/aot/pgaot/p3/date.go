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
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = F_strlen(m, l1)
	mBase = m.M
	if base.Ui32(int32(10)) < base.Ui32(v13) {
		v88 = l1
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
	v95 = int32(0)
	v98 = F_errstart(m, int32(15), v95)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L19
	} else {
		goto L25
	}
L3:
	;
	v16 = int32(1)
	v23 = v16
	v24 = v16
	goto L4
L4:
	;
	v28 = l1 + v23<<(uint(int32(4))%32)
	v29 = F_strlen(m, v28)
	mBase = m.M
	if base.Ui32(int32(11)) <= base.Ui32(v29) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v120 = v83
	goto L1
L6:
	;
	v88 = v28
	goto L2
L7:
	;
	goto L8
L8:
	;
	v33 = v28 - int32(16)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28))))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33))))
	if v37 == int32(0) {
		v56 = v36
		v57 = v37
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v85 = v23 + int32(1)
	if v85 != l2 {
		v23 = v85
		v24 = v83
		goto L4
	} else {
		goto L24
	}
L10:
	;
	if v57-v56 < int32(0) {
		v83 = v24
		goto L9
	} else {
		goto L18
	}
L11:
	;
	goto L10
L12:
	;
	if v36 != v37 {
		v56 = v36
		v57 = v37
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v41 = v33
	v42 = v28
	goto L14
L14:
	;
	v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+1)))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41)+1)))
	if v46 == int32(0) {
		v56 = v45
		v57 = v46
		goto L11
	} else {
		goto L16
	}
L15:
	;
	v56 = v45
	v57 = v46
	goto L11
L16:
	;
	v49 = int32(1)
	if v45 == v46 {
		v41 = v41 + v49
		v42 = v42 + v49
		goto L14
	} else {
		goto L17
	}
L17:
	;
	goto L15
L18:
	;
	v61 = int32(0)
	v64 = F_errstart(m, int32(15), v61)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	return int32(0)
L20:
	;
	if v64 == int32(0) {
		v83 = v61
		goto L9
	} else {
		goto L21
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v28
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v33
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	F_errmsg_internal(m, int32(705985), v11+int32(16))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	F_errfinish(m, int32(495687), int32(4925), int32(394798))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L19
	} else {
		goto L23
	}
L23:
	;
	v83 = v61
	goto L9
L24:
	;
	goto L5
L25:
	;
	if v98 == int32(0) {
		v120 = v95
		goto L1
	} else {
		goto L26
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v88
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(671928), v11)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L19
	} else {
		goto L27
	}
L27:
	;
	F_errfinish(m, int32(495687), int32(4914), int32(394798))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L19
	} else {
		goto L28
	}
L28:
	;
	v120 = v95
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2443), int32(0), v4, v5)
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
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v7 int32
	_ = v7
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v7 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v19 = F_Int64GetDatum(m, v18)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			v23 = F_DirectFunctionCall2Coll(m, int32(1284), int32(0), v19, v3)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				return v23
			}
		}
	} else {
		if v7 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v19 = F_Int64GetDatum(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = F_DirectFunctionCall2Coll(m, int32(1284), int32(0), v19, v3)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					return v23
				}
			}
		} else {
			if int32(106751983) <= v7 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(236026), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495450), int32(658), int32(31099))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
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
				v18 = base.I64_extend_i32_s(v7) * int64(86400000000)
				v19 = F_Int64GetDatum(m, v18)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = F_DirectFunctionCall2Coll(m, int32(1284), int32(0), v19, v3)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						return v23
					}
				}
			}
		}
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
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[515]))
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
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pq_getmsgint(m, v2, int32(4))
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if base.Ui32(v4-int32(2147483647)) < base.Ui32(int32(2)) {
			return v4
		} else {
			if base.Ui32(v4+int32(2451545)) < base.Ui32(int32(2147483494)) {
				return v4
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v22 = m.ExcPending
					if v22 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(400118), int32(0))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495450), int32(223), int32(36438))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
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
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
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
			v19 = int32(24)
			v21 = int32(65280)
			v23 = int32(8)
			*(*int32)(unsafe.Add(mBase, uint32(v16+v17))) = v8<<(uint(v19)%32) | v8&v21<<(uint(v23)%32) | (int32(base.Ui32(v8)>>(uint(v23)%32))&v21 | int32(base.Ui32(v8)>>(uint(v19)%32)))
			v36 = v16 + int32(4)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v36
			v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			*(*int32)(unsafe.Add(mBase, uint32(v39))) = v36 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return v39
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
	var v38 int32
	_ = v38
	var v150 int32
	_ = v150
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v191 int32
	_ = v191
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v284 int64
	_ = v284
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int64
	_ = v307
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int64
	_ = v330
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
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
			v260 = m.ExcPending
			if v260 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v263 = m.ExcPending
				if v263 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v14
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v12
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = v10
					F_errmsg(m, int32(462981), v8)
					mBase = m.M
					v269 = m.ExcPending
					if v269 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495450), int32(267), int32(354951))
						mBase = m.M
						v274 = m.ExcPending
						if v274 != 0 {
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
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
			if int32(base.Ui32(v10)>>(uint(int32(31))%32)) != 0 {
				if int32(0) < v38 {
					*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = int32(1) - v38
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
					if base.Ui32(int32(-12)) <= base.Ui32(v150-int32(13)) {
						v160 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						if base.Ui32(int32(-31)) <= base.Ui32(v160-int32(32)) {
							v170 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
							v172 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
							if v172&int32(3) != 0 {
								v182 = int32(0)
							} else {
								v177 = base.I32_rem_s(v172, int32(100))
								if v177 != 0 {
									v182 = int32(1)
								} else {
									v179 = base.I32_rem_s(v172, int32(400))
									v182 = base.B2i32(v179 == int32(0))
								}
							}
							v185 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
							v191 = *(*int32)(unsafe.Add(mBase, uint32(v182*int32(52)+v185<<(uint(int32(2))%32))+uint32(_consts[940])))
							if v170 <= v191 {
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
				if int32(0) < v38 {
					v150 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
					if base.Ui32(int32(-12)) <= base.Ui32(v150-int32(13)) {
						v160 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						if base.Ui32(int32(-31)) <= base.Ui32(v160-int32(32)) {
							v170 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
							v172 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
							if v172&int32(3) != 0 {
								v182 = int32(0)
							} else {
								v177 = base.I32_rem_s(v172, int32(100))
								if v177 != 0 {
									v182 = int32(1)
								} else {
									v179 = base.I32_rem_s(v172, int32(400))
									v182 = base.B2i32(v179 == int32(0))
								}
							}
							v185 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
							v191 = *(*int32)(unsafe.Add(mBase, uint32(v182*int32(52)+v185<<(uint(int32(2))%32))+uint32(_consts[940])))
							if v170 <= v191 {
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
				v278 = m.ExcPending
				if v278 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(134217858))
					mBase = m.M
					v281 = m.ExcPending
					if v281 != 0 {
						return int32(0)
					} else {
						v282 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
						*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v282
						v284 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v8)+52)) = base.I64_rotl(v284, int64(32))
						F_errmsg(m, int32(462981), v8+int32(48))
						mBase = m.M
						v292 = m.ExcPending
						if v292 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(495450), int32(277), int32(354951))
							mBase = m.M
							v297 = m.ExcPending
							if v297 != 0 {
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
						v324 = m.ExcPending
						if v324 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v327 = m.ExcPending
							if v327 != 0 {
								return int32(0)
							} else {
								v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v328
								v330 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v330, int64(32))
								F_errmsg(m, int32(463025), v8+int32(32))
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495450), int32(284), int32(354951))
									mBase = m.M
									v343 = m.ExcPending
									if v343 != 0 {
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
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							v221 = base.B2i32(int32(2) < v202)
							if int32(2) < v202 {
								v222 = int32(4800)
							} else {
								v222 = int32(4799)
							}
							v223 = v222 + v203
							v228 = base.I32_div_s(v223, int32(4))
							v231 = base.I32_div_s(v223, int32(-100))
							v234 = base.I32_div_s(v223, int32(400))
							if int32(2) < v202 {
								v238 = int32(1)
							} else {
								v238 = int32(13)
							}
							v243 = base.I32_div_s((v238+v202)*int32(7834), int32(256))
							v246 = v216 + v223*int32(365) + v228 + v231 + v234 + v243 - int32(32167)
							if base.Ui32(int32(2147483494)) <= base.Ui32(v246) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v301 = m.ExcPending
								if v301 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v304 = m.ExcPending
									if v304 != 0 {
										return int32(0)
									} else {
										v305 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v305
										v307 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v307, int64(32))
										F_errmsg(m, int32(463025), v8+int32(16))
										mBase = m.M
										v315 = m.ExcPending
										if v315 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(495450), int32(293), int32(354951))
											mBase = m.M
											v320 = m.ExcPending
											if v320 != 0 {
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
								return v246 - int32(2451545)
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v324 = m.ExcPending
							if v324 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v327 = m.ExcPending
								if v327 != 0 {
									return int32(0)
								} else {
									v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v328
									v330 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v330, int64(32))
									F_errmsg(m, int32(463025), v8+int32(32))
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495450), int32(284), int32(354951))
										mBase = m.M
										v343 = m.ExcPending
										if v343 != 0 {
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
					if v203 < int32(5874898) {
						v216 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						v221 = base.B2i32(int32(2) < v202)
						if int32(2) < v202 {
							v222 = int32(4800)
						} else {
							v222 = int32(4799)
						}
						v223 = v222 + v203
						v228 = base.I32_div_s(v223, int32(4))
						v231 = base.I32_div_s(v223, int32(-100))
						v234 = base.I32_div_s(v223, int32(400))
						if int32(2) < v202 {
							v238 = int32(1)
						} else {
							v238 = int32(13)
						}
						v243 = base.I32_div_s((v238+v202)*int32(7834), int32(256))
						v246 = v216 + v223*int32(365) + v228 + v231 + v234 + v243 - int32(32167)
						if base.Ui32(int32(2147483494)) <= base.Ui32(v246) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v301 = m.ExcPending
							if v301 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v304 = m.ExcPending
								if v304 != 0 {
									return int32(0)
								} else {
									v305 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v305
									v307 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v307, int64(32))
									F_errmsg(m, int32(463025), v8+int32(16))
									mBase = m.M
									v315 = m.ExcPending
									if v315 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495450), int32(293), int32(354951))
										mBase = m.M
										v320 = m.ExcPending
										if v320 != 0 {
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
							return v246 - int32(2451545)
						}
					} else {
						if v203 != int32(5874898) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v324 = m.ExcPending
							if v324 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v327 = m.ExcPending
								if v327 != 0 {
									return int32(0)
								} else {
									v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v328
									v330 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v330, int64(32))
									F_errmsg(m, int32(463025), v8+int32(32))
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495450), int32(284), int32(354951))
										mBase = m.M
										v343 = m.ExcPending
										if v343 != 0 {
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
							if int32(6) <= v202 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v324 = m.ExcPending
								if v324 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v327 = m.ExcPending
									if v327 != 0 {
										return int32(0)
									} else {
										v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v328
										v330 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v330, int64(32))
										F_errmsg(m, int32(463025), v8+int32(32))
										mBase = m.M
										v338 = m.ExcPending
										if v338 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(495450), int32(284), int32(354951))
											mBase = m.M
											v343 = m.ExcPending
											if v343 != 0 {
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
								v216 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
								v221 = base.B2i32(int32(2) < v202)
								if int32(2) < v202 {
									v222 = int32(4800)
								} else {
									v222 = int32(4799)
								}
								v223 = v222 + v203
								v228 = base.I32_div_s(v223, int32(4))
								v231 = base.I32_div_s(v223, int32(-100))
								v234 = base.I32_div_s(v223, int32(400))
								if int32(2) < v202 {
									v238 = int32(1)
								} else {
									v238 = int32(13)
								}
								v243 = base.I32_div_s((v238+v202)*int32(7834), int32(256))
								v246 = v216 + v223*int32(365) + v228 + v231 + v234 + v243 - int32(32167)
								if base.Ui32(int32(2147483494)) <= base.Ui32(v246) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v301 = m.ExcPending
									if v301 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v304 = m.ExcPending
										if v304 != 0 {
											return int32(0)
										} else {
											v305 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
											*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v305
											v307 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
											*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v307, int64(32))
											F_errmsg(m, int32(463025), v8+int32(16))
											mBase = m.M
											v315 = m.ExcPending
											if v315 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(495450), int32(293), int32(354951))
												mBase = m.M
												v320 = m.ExcPending
												if v320 != 0 {
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
									return v246 - int32(2451545)
								}
							}
						}
					}
				}
			}
		}
	} else {
		v33 = v8 + int32(68)
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
		if int32(base.Ui32(v10)>>(uint(int32(31))%32)) != 0 {
			if int32(0) < v38 {
				*(*int32)(unsafe.Add(mBase, uint32(v33)+20)) = int32(1) - v38
				v150 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
				if base.Ui32(int32(-12)) <= base.Ui32(v150-int32(13)) {
					v160 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					if base.Ui32(int32(-31)) <= base.Ui32(v160-int32(32)) {
						v170 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						v172 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
						if v172&int32(3) != 0 {
							v182 = int32(0)
						} else {
							v177 = base.I32_rem_s(v172, int32(100))
							if v177 != 0 {
								v182 = int32(1)
							} else {
								v179 = base.I32_rem_s(v172, int32(400))
								v182 = base.B2i32(v179 == int32(0))
							}
						}
						v185 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
						v191 = *(*int32)(unsafe.Add(mBase, uint32(v182*int32(52)+v185<<(uint(int32(2))%32))+uint32(_consts[940])))
						if v170 <= v191 {
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
			if int32(0) < v38 {
				v150 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
				if base.Ui32(int32(-12)) <= base.Ui32(v150-int32(13)) {
					v160 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
					if base.Ui32(int32(-31)) <= base.Ui32(v160-int32(32)) {
						v170 = *(*int32)(unsafe.Add(mBase, uint32(v33)+12))
						v172 = *(*int32)(unsafe.Add(mBase, uint32(v33)+20))
						if v172&int32(3) != 0 {
							v182 = int32(0)
						} else {
							v177 = base.I32_rem_s(v172, int32(100))
							if v177 != 0 {
								v182 = int32(1)
							} else {
								v179 = base.I32_rem_s(v172, int32(400))
								v182 = base.B2i32(v179 == int32(0))
							}
						}
						v185 = *(*int32)(unsafe.Add(mBase, uint32(v33)+16))
						v191 = *(*int32)(unsafe.Add(mBase, uint32(v182*int32(52)+v185<<(uint(int32(2))%32))+uint32(_consts[940])))
						if v170 <= v191 {
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
			v278 = m.ExcPending
			if v278 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v281 = m.ExcPending
				if v281 != 0 {
					return int32(0)
				} else {
					v282 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v282
					v284 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
					*(*int64)(unsafe.Add(mBase, uint32(v8)+52)) = base.I64_rotl(v284, int64(32))
					F_errmsg(m, int32(462981), v8+int32(48))
					mBase = m.M
					v292 = m.ExcPending
					if v292 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(495450), int32(277), int32(354951))
						mBase = m.M
						v297 = m.ExcPending
						if v297 != 0 {
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
					v324 = m.ExcPending
					if v324 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v327 = m.ExcPending
						if v327 != 0 {
							return int32(0)
						} else {
							v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v328
							v330 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v330, int64(32))
							F_errmsg(m, int32(463025), v8+int32(32))
							mBase = m.M
							v338 = m.ExcPending
							if v338 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(495450), int32(284), int32(354951))
								mBase = m.M
								v343 = m.ExcPending
								if v343 != 0 {
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
						v216 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
						v221 = base.B2i32(int32(2) < v202)
						if int32(2) < v202 {
							v222 = int32(4800)
						} else {
							v222 = int32(4799)
						}
						v223 = v222 + v203
						v228 = base.I32_div_s(v223, int32(4))
						v231 = base.I32_div_s(v223, int32(-100))
						v234 = base.I32_div_s(v223, int32(400))
						if int32(2) < v202 {
							v238 = int32(1)
						} else {
							v238 = int32(13)
						}
						v243 = base.I32_div_s((v238+v202)*int32(7834), int32(256))
						v246 = v216 + v223*int32(365) + v228 + v231 + v234 + v243 - int32(32167)
						if base.Ui32(int32(2147483494)) <= base.Ui32(v246) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v301 = m.ExcPending
							if v301 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v304 = m.ExcPending
								if v304 != 0 {
									return int32(0)
								} else {
									v305 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v305
									v307 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v307, int64(32))
									F_errmsg(m, int32(463025), v8+int32(16))
									mBase = m.M
									v315 = m.ExcPending
									if v315 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495450), int32(293), int32(354951))
										mBase = m.M
										v320 = m.ExcPending
										if v320 != 0 {
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
							return v246 - int32(2451545)
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v324 = m.ExcPending
						if v324 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v327 = m.ExcPending
							if v327 != 0 {
								return int32(0)
							} else {
								v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v328
								v330 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v330, int64(32))
								F_errmsg(m, int32(463025), v8+int32(32))
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495450), int32(284), int32(354951))
									mBase = m.M
									v343 = m.ExcPending
									if v343 != 0 {
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
				if v203 < int32(5874898) {
					v216 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
					v221 = base.B2i32(int32(2) < v202)
					if int32(2) < v202 {
						v222 = int32(4800)
					} else {
						v222 = int32(4799)
					}
					v223 = v222 + v203
					v228 = base.I32_div_s(v223, int32(4))
					v231 = base.I32_div_s(v223, int32(-100))
					v234 = base.I32_div_s(v223, int32(400))
					if int32(2) < v202 {
						v238 = int32(1)
					} else {
						v238 = int32(13)
					}
					v243 = base.I32_div_s((v238+v202)*int32(7834), int32(256))
					v246 = v216 + v223*int32(365) + v228 + v231 + v234 + v243 - int32(32167)
					if base.Ui32(int32(2147483494)) <= base.Ui32(v246) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v301 = m.ExcPending
						if v301 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v304 = m.ExcPending
							if v304 != 0 {
								return int32(0)
							} else {
								v305 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v305
								v307 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v307, int64(32))
								F_errmsg(m, int32(463025), v8+int32(16))
								mBase = m.M
								v315 = m.ExcPending
								if v315 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495450), int32(293), int32(354951))
									mBase = m.M
									v320 = m.ExcPending
									if v320 != 0 {
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
						return v246 - int32(2451545)
					}
				} else {
					if v203 != int32(5874898) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v324 = m.ExcPending
						if v324 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v327 = m.ExcPending
							if v327 != 0 {
								return int32(0)
							} else {
								v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
								*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v328
								v330 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v330, int64(32))
								F_errmsg(m, int32(463025), v8+int32(32))
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(495450), int32(284), int32(354951))
									mBase = m.M
									v343 = m.ExcPending
									if v343 != 0 {
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
						if int32(6) <= v202 {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v324 = m.ExcPending
							if v324 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v327 = m.ExcPending
								if v327 != 0 {
									return int32(0)
								} else {
									v328 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
									*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v328
									v330 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
									*(*int64)(unsafe.Add(mBase, uint32(v8)+36)) = base.I64_rotl(v330, int64(32))
									F_errmsg(m, int32(463025), v8+int32(32))
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(495450), int32(284), int32(354951))
										mBase = m.M
										v343 = m.ExcPending
										if v343 != 0 {
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
							v216 = *(*int32)(unsafe.Add(mBase, uint32(v8)+80))
							v221 = base.B2i32(int32(2) < v202)
							if int32(2) < v202 {
								v222 = int32(4800)
							} else {
								v222 = int32(4799)
							}
							v223 = v222 + v203
							v228 = base.I32_div_s(v223, int32(4))
							v231 = base.I32_div_s(v223, int32(-100))
							v234 = base.I32_div_s(v223, int32(400))
							if int32(2) < v202 {
								v238 = int32(1)
							} else {
								v238 = int32(13)
							}
							v243 = base.I32_div_s((v238+v202)*int32(7834), int32(256))
							v246 = v216 + v223*int32(365) + v228 + v231 + v234 + v243 - int32(32167)
							if base.Ui32(int32(2147483494)) <= base.Ui32(v246) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v301 = m.ExcPending
								if v301 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v304 = m.ExcPending
									if v304 != 0 {
										return int32(0)
									} else {
										v305 = *(*int32)(unsafe.Add(mBase, uint32(v8)+88))
										*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v305
										v307 = *(*int64)(unsafe.Add(mBase, uint32(v8)+80))
										*(*int64)(unsafe.Add(mBase, uint32(v8)+20)) = base.I64_rotl(v307, int64(32))
										F_errmsg(m, int32(463025), v8+int32(16))
										mBase = m.M
										v315 = m.ExcPending
										if v315 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(495450), int32(293), int32(354951))
											mBase = m.M
											v320 = m.ExcPending
											if v320 != 0 {
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
								return v246 - int32(2451545)
							}
						}
					}
				}
			}
		}
	}
}
