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
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v232 int32
	_ = v232
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	if l1&int32(3) == int32(0) {
		v36 = l1
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v11 + int32(32)
	return v232
L2:
	;
	v207 = int32(0)
	v210 = F_errstart(m, int32(15), v207)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L53
	} else {
		goto L59
	}
L3:
	;
	if base.Ui32(int32(10)) < base.Ui32(v69) {
		v200 = l1
		goto L2
	} else {
		goto L20
	}
L4:
	;
	v69 = v61 - l1
	goto L3
L5:
	;
	v40 = v36
	goto L14
L6:
	;
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v20 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v69 = int32(0)
	goto L3
L8:
	;
	goto L9
L9:
	;
	v25 = l1
	goto L10
L10:
	;
	v29 = v25 + int32(1)
	if v29&int32(3) == int32(0) {
		v36 = v29
		goto L5
	} else {
		goto L12
	}
L11:
	;
	v61 = v29
	goto L4
L12:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v29))))
	if v34 != 0 {
		v25 = v29
		goto L10
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v49 = int32(-2139062144)
	if (int32(16843008)-v46|v46)&v49 == v49 {
		v40 = v40 + int32(4)
		goto L14
	} else {
		goto L16
	}
L15:
	;
	v55 = v40
	goto L17
L16:
	;
	goto L15
L17:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55))))
	if v59 != 0 {
		v55 = v55 + int32(1)
		goto L17
	} else {
		goto L19
	}
L18:
	;
	v61 = v55
	goto L4
L19:
	;
	goto L18
L20:
	;
	v72 = int32(1)
	v79 = v72
	v80 = v72
	goto L21
L21:
	;
	v84 = l1 + v79<<(uint(int32(4))%32)
	if v84&int32(3) == int32(0) {
		v108 = v84
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v232 = v195
	goto L1
L23:
	;
	if base.Ui32(int32(11)) <= base.Ui32(v141) {
		goto L40
	} else {
		goto L41
	}
L24:
	;
	v141 = v133 - v84
	goto L23
L25:
	;
	v112 = v108
	goto L34
L26:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	if v92 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v141 = int32(0)
	goto L23
L28:
	;
	goto L29
L29:
	;
	v97 = v84
	goto L30
L30:
	;
	v101 = v97 + int32(1)
	if v101&int32(3) == int32(0) {
		v108 = v101
		goto L25
	} else {
		goto L32
	}
L31:
	;
	v133 = v101
	goto L24
L32:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	if v106 != 0 {
		v97 = v101
		goto L30
	} else {
		goto L33
	}
L33:
	;
	goto L31
L34:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	v121 = int32(-2139062144)
	if (int32(16843008)-v118|v118)&v121 == v121 {
		v112 = v112 + int32(4)
		goto L34
	} else {
		goto L36
	}
L35:
	;
	v127 = v112
	goto L37
L36:
	;
	goto L35
L37:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v131 != 0 {
		v127 = v127 + int32(1)
		goto L37
	} else {
		goto L39
	}
L38:
	;
	v133 = v127
	goto L24
L39:
	;
	goto L38
L40:
	;
	v200 = v84
	goto L2
L41:
	;
	goto L42
L42:
	;
	v145 = v84 - int32(16)
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v149 == int32(0) {
		v168 = v148
		v169 = v149
		goto L45
	} else {
		goto L46
	}
L43:
	;
	v197 = v79 + int32(1)
	if v197 != l2 {
		v79 = v197
		v80 = v195
		goto L21
	} else {
		goto L58
	}
L44:
	;
	if v169-v168 < int32(0) {
		v195 = v80
		goto L43
	} else {
		goto L52
	}
L45:
	;
	goto L44
L46:
	;
	if v148 != v149 {
		v168 = v148
		v169 = v149
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v153 = v145
	v154 = v84
	goto L48
L48:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154)+1)))
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v153)+1)))
	if v158 == int32(0) {
		v168 = v157
		v169 = v158
		goto L45
	} else {
		goto L50
	}
L49:
	;
	v168 = v157
	v169 = v158
	goto L45
L50:
	;
	v161 = int32(1)
	if v157 == v158 {
		v153 = v153 + v161
		v154 = v154 + v161
		goto L48
	} else {
		goto L51
	}
L51:
	;
	goto L49
L52:
	;
	v173 = int32(0)
	v176 = F_errstart(m, int32(15), v173)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	return int32(0)
L54:
	;
	if v176 == int32(0) {
		v195 = v173
		goto L43
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v84
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v145
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = l0
	F_errmsg_internal(m, int32(698991), v11+int32(16))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L53
	} else {
		goto L56
	}
L56:
	;
	F_errfinish(m, int32(490329), int32(4925), int32(390628))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L53
	} else {
		goto L57
	}
L57:
	;
	v195 = v173
	goto L43
L58:
	;
	goto L22
L59:
	;
	if v210 == int32(0) {
		v232 = v207
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v200
	*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = int32(11)
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = l0
	F_errmsg_internal(m, int32(664934), v11)
	mBase = m.M
	v220 = m.ExcPending
	if v220 != 0 {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	F_errfinish(m, int32(490329), int32(4914), int32(390628))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L53
	} else {
		goto L62
	}
L62:
	;
	v232 = v207
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
	v6 = F_DirectFunctionCall2Coll(m, int32(2442), int32(0), v4, v5)
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
			v23 = F_DirectFunctionCall2Coll(m, int32(1283), int32(0), v19, v3)
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
				v23 = F_DirectFunctionCall2Coll(m, int32(1283), int32(0), v19, v3)
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
						F_errmsg(m, int32(233532), int32(0))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490092), int32(658), int32(30663))
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
					v23 = F_DirectFunctionCall2Coll(m, int32(1283), int32(0), v19, v3)
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
						F_errmsg(m, int32(395881), int32(0))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490092), int32(223), int32(35974))
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
					F_errmsg(m, int32(457980), v8)
					mBase = m.M
					v269 = m.ExcPending
					if v269 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490092), int32(267), int32(351095))
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
						F_errmsg(m, int32(457980), v8+int32(48))
						mBase = m.M
						v292 = m.ExcPending
						if v292 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490092), int32(277), int32(351095))
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
								F_errmsg(m, int32(458024), v8+int32(32))
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490092), int32(284), int32(351095))
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
										F_errmsg(m, int32(458024), v8+int32(16))
										mBase = m.M
										v315 = m.ExcPending
										if v315 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490092), int32(293), int32(351095))
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
									F_errmsg(m, int32(458024), v8+int32(32))
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490092), int32(284), int32(351095))
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
									F_errmsg(m, int32(458024), v8+int32(16))
									mBase = m.M
									v315 = m.ExcPending
									if v315 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490092), int32(293), int32(351095))
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
									F_errmsg(m, int32(458024), v8+int32(32))
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490092), int32(284), int32(351095))
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
										F_errmsg(m, int32(458024), v8+int32(32))
										mBase = m.M
										v338 = m.ExcPending
										if v338 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490092), int32(284), int32(351095))
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
											F_errmsg(m, int32(458024), v8+int32(16))
											mBase = m.M
											v315 = m.ExcPending
											if v315 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(490092), int32(293), int32(351095))
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
					F_errmsg(m, int32(457980), v8+int32(48))
					mBase = m.M
					v292 = m.ExcPending
					if v292 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490092), int32(277), int32(351095))
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
							F_errmsg(m, int32(458024), v8+int32(32))
							mBase = m.M
							v338 = m.ExcPending
							if v338 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490092), int32(284), int32(351095))
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
									F_errmsg(m, int32(458024), v8+int32(16))
									mBase = m.M
									v315 = m.ExcPending
									if v315 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490092), int32(293), int32(351095))
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
								F_errmsg(m, int32(458024), v8+int32(32))
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490092), int32(284), int32(351095))
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
								F_errmsg(m, int32(458024), v8+int32(16))
								mBase = m.M
								v315 = m.ExcPending
								if v315 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490092), int32(293), int32(351095))
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
								F_errmsg(m, int32(458024), v8+int32(32))
								mBase = m.M
								v338 = m.ExcPending
								if v338 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490092), int32(284), int32(351095))
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
									F_errmsg(m, int32(458024), v8+int32(32))
									mBase = m.M
									v338 = m.ExcPending
									if v338 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490092), int32(284), int32(351095))
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
										F_errmsg(m, int32(458024), v8+int32(16))
										mBase = m.M
										v315 = m.ExcPending
										if v315 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490092), int32(293), int32(351095))
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
