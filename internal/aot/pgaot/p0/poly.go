package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_above(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+32))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.F64_lt(v14, v15)
						}
					} else {
						return base.F64_lt(v14, v15)
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return base.F64_lt(v14, v15)
					}
				} else {
					return base.F64_lt(v14, v15)
				}
			}
		}
	}
}
func F_poly_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
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
	var v51 float64
	_ = v51
	var v60 int32
	_ = v60
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 float64
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 float64
	_ = v81
	var v85 int32
	_ = v85
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v92 int32
	_ = v92
	var v95 float64
	_ = v95
	var v99 float64
	_ = v99
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v108 float64
	_ = v108
	var v109 int32
	_ = v109
	var v125 float64
	_ = v125
	var v127 int32
	_ = v127
	var v136 int32
	_ = v136
	var v142 float64
	_ = v142
	var v144 float64
	_ = v144
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v150 float64
	_ = v150
	var v154 int32
	_ = v154
	var v157 float64
	_ = v157
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v165 float64
	_ = v165
	var v170 float64
	_ = v170
	var v171 int32
	_ = v171
	var v178 float64
	_ = v178
	var v184 float64
	_ = v184
	var v185 float64
	_ = v185
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v202 float64
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v244 float64
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v17 - int32(-64)
	return v262
L2:
	;
	v246 = F_Float8GetDatum(m, v244)
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L3
	} else {
		goto L50
	}
L3:
	;
	return int32(0)
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v25 = F_pg_detoast_datum(m, v24)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v27 = F_poly_overlap_internal(m, v20, v25)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	if v27 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if int32(0) < v31 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v230 = F_Float8GetDatum(m, float64(0))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L3
	} else {
		goto L49
	}
L10:
	;
	v34 = int32(40)
	v35 = v25 + v34
	v37 = v20 + v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	v41 = v38
	v42 = v31
	v44 = v2
	v45 = v2
	v51 = float64(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	v226 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v226)
	v262 = int32(0)
	goto L1
L13:
	;
	v60 = v44
	goto L15
L14:
	;
	if v45 != 0 {
		v244 = v51
		goto L2
	} else {
		goto L48
	}
L15:
	;
	if base.B2i32(v41 <= int32(0)) == int32(0) {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L14
L17:
	;
	if v60 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v210 = v60 + int32(1)
	if v210 < v42 {
		v60 = v210
		goto L15
	} else {
		goto L47
	}
L20:
	;
	v71 = v60
	goto L22
L21:
	;
	v71 = v42
	goto L22
L22:
	;
	v72 = int32(4)
	v74 = v37 + v71<<(uint(v72)%32)
	v75 = int32(16)
	v76 = v74 - v75
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v76)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+32)) = v77
	v79 = int32(8)
	v80 = v74 - v79
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v80)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+40)) = v81
	v85 = v37 + v60<<(uint(v72)%32)
	v86 = *(*float64)(unsafe.Add(mBase, uint32(v85)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+48)) = v86
	v88 = *(*float64)(unsafe.Add(mBase, uint32(v85)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+56)) = v88
	v92 = v35 + v41<<(uint(v72)%32)
	v95 = *(*float64)(unsafe.Add(mBase, uint32(v92-v75)))
	*(*float64)(unsafe.Add(mBase, uint32(v17))) = v95
	v99 = *(*float64)(unsafe.Add(mBase, uint32(v92-v79)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+8)) = v99
	v101 = *(*float64)(unsafe.Add(mBase, uint32(v25)+40))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+16)) = v101
	v103 = *(*float64)(unsafe.Add(mBase, uint32(v25)+48))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+24)) = v103
	v108 = F_lseg_closept_lseg(m, int32(0), v15+int32(-32), v17)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L3
	} else {
		goto L23
	}
L23:
	;
	if v45 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if int32(2) <= v127 {
		goto L30
	} else {
		goto L31
	}
L25:
	;
	v125 = v108
	goto L24
L26:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v108)&int64(9223372036854775807)) {
		v125 = v51
		goto L24
	} else {
		goto L27
	}
L27:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v51)&int64(9223372036854775807)) {
		goto L25
	} else {
		goto L28
	}
L28:
	;
	if base.F64_gt(v51, v108) == int32(0) {
		v125 = v51
		goto L24
	} else {
		goto L29
	}
L29:
	;
	goto L25
L30:
	;
	v136 = int32(1)
	v142 = v125
	goto L33
L31:
	;
	v192 = v127
	v202 = v125
	goto L32
L32:
	;
	v204 = int32(1)
	v206 = v60 + v204
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
	if v206 < v207 {
		v41 = v192
		v42 = v207
		v44 = v206
		v45 = v204
		v51 = v202
		goto L13
	} else {
		goto L46
	}
L33:
	;
	v144 = *(*float64)(unsafe.Add(mBase, uint32(v76)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+32)) = v144
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v80)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+40)) = v146
	v148 = *(*float64)(unsafe.Add(mBase, uint32(v85)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+48)) = v148
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v85)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+56)) = v150
	v154 = v35 + v136<<(uint(int32(4))%32)
	v157 = *(*float64)(unsafe.Add(mBase, uint32(v154-int32(16))))
	*(*float64)(unsafe.Add(mBase, uint32(v17))) = v157
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v154-int32(8))))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+8)) = v161
	v163 = *(*float64)(unsafe.Add(mBase, uint32(v154)))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+16)) = v163
	v165 = *(*float64)(unsafe.Add(mBase, uint32(v154)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v17)+24)) = v165
	v170 = F_lseg_closept_lseg(m, int32(0), v15+int32(-32), v17)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L3
	} else {
		goto L35
	}
L34:
	;
	v192 = v188
	v202 = v185
	goto L32
L35:
	;
	if base.Ui64(base.I64_reinterpret_f64(v170)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	if base.F64_gt(v142, v170) != 0 {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	v185 = v142
	goto L38
L38:
	;
	v187 = v136 + int32(1)
	v188 = *(*int32)(unsafe.Add(mBase, uint32(v25)+4))
	if v187 < v188 {
		v136 = v187
		v142 = v185
		goto L33
	} else {
		goto L45
	}
L39:
	;
	v178 = v170
	goto L41
L40:
	;
	v178 = v142
	goto L41
L41:
	;
	if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v142)&int64(9223372036854775807)) {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v184 = v170
	goto L44
L43:
	;
	v184 = v178
	goto L44
L44:
	;
	v185 = v184
	goto L38
L45:
	;
	goto L34
L46:
	;
	v244 = v202
	goto L2
L47:
	;
	goto L16
L48:
	;
	goto L12
L49:
	;
	v262 = v230
	goto L1
L50:
	;
	v262 = v246
	goto L1
}
func F_poly_left(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.F64_gt(v14, v15)
						}
					} else {
						return base.F64_gt(v14, v15)
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return base.F64_gt(v14, v15)
					}
				} else {
					return base.F64_gt(v14, v15)
				}
			}
		}
	}
}
func F_poly_overabove(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+32))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+32))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.F64_le(v14, v15)
						}
					} else {
						return base.F64_le(v14, v15)
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return base.F64_le(v14, v15)
					}
				} else {
					return base.F64_le(v14, v15)
				}
			}
		}
	}
}
func F_poly_overbelow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v15 float64
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							return base.F64_ge(v14, v15)
						}
					} else {
						return base.F64_ge(v14, v15)
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						return base.F64_ge(v14, v15)
					}
				} else {
					return base.F64_ge(v14, v15)
				}
			}
		}
	}
}
func F_poly_overlap_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v25 float64
	_ = v25
	var v26 float64
	_ = v26
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int64
	_ = v54
	var v56 int64
	_ = v56
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
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v85 int32
	_ = v85
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v106 int64
	_ = v106
	var v108 int64
	_ = v108
	var v110 int64
	_ = v110
	var v119 int32
	_ = v119
	var v130 int32
	_ = v130
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v144 int64
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v152 int32
	_ = v152
	var v153 int64
	_ = v153
	var v155 int64
	_ = v155
	var v158 int32
	_ = v158
	var v166 int64
	_ = v166
	var v168 int64
	_ = v168
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	v3 = int32(0)
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(l1)+8))
	if base.F64_le(v18, base.F64_add(v19, float64(1e-06))) == v3 {
		v199 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v16 - int32(-64)
	return v199
L2:
	;
	v25 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
	v26 = *(*float64)(unsafe.Add(mBase, uint32(l0)+8))
	if base.F64_le(v25, base.F64_add(v26, float64(1e-06))) == int32(0) {
		v199 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v32 = *(*float64)(unsafe.Add(mBase, uint32(l0)+32))
	v33 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_le(v32, base.F64_add(v33, float64(1e-06))) == int32(0) {
		v199 = v3
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v39 = *(*float64)(unsafe.Add(mBase, uint32(l1)+32))
	v40 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	if base.F64_le(v39, base.F64_add(v40, float64(1e-06))) == int32(0) {
		v199 = v3
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v47 = l0 + int32(40)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v53 = v47 + v48<<(uint(int32(4))%32) - int32(16)
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v53)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v54
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v53)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v56
	if v48 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v187 = l1 + int32(40)
	v188 = F_point_inside(m, v47, v176, v187)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L20
	} else {
		goto L27
	}
L7:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v176 = v60
	goto L6
L8:
	;
	goto L9
L9:
	;
	v62 = l1 + int32(40)
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v65 = v14 + int32(-48)
	v67 = v14 + int32(-16)
	v71 = v63
	v72 = v48
	v76 = v3
	goto L10
L10:
	;
	v85 = v62 + v71<<(uint(int32(4))%32) - int32(16)
	v96 = v76
	goto L12
L11:
	;
	v176 = v71
	goto L6
L12:
	;
	v103 = v47 + v96<<(uint(int32(4))%32)
	v104 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v67)+8)) = v104
	v106 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = v106
	v108 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v108
	v110 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v110
	if base.B2i32(v71 <= int32(0)) == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	v119 = int32(0)
	goto L17
L15:
	;
	goto L16
L16:
	;
	v166 = *(*int64)(unsafe.Add(mBase, uint32(v103)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v166
	v168 = *(*int64)(unsafe.Add(mBase, uint32(v103)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v168
	v171 = v96 + int32(1)
	if v171 < v72 {
		v96 = v171
		goto L12
	} else {
		goto L26
	}
L17:
	;
	v130 = v62 + v119<<(uint(int32(4))%32)
	v131 = *(*int64)(unsafe.Add(mBase, uint32(v130)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v130)))
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v133
	v138 = F_lseg_interpt_lseg(m, int32(0), v14+int32(-32), v16)
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v153 = *(*int64)(unsafe.Add(mBase, uint32(v67)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v153
	v155 = *(*int64)(unsafe.Add(mBase, uint32(v67)))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = v155
	v158 = v96 + int32(1)
	if base.B2i32(v152 <= v158)|v138 == int32(0) {
		v71 = v148
		v72 = v152
		v76 = v158
		goto L10
	} else {
		goto L24
	}
L19:
	;
	goto L18
L20:
	;
	return int32(0)
L21:
	;
	v142 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v142
	v144 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v144
	v147 = v119 + int32(1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v148 <= v147 {
		goto L19
	} else {
		goto L22
	}
L22:
	;
	if v138 == int32(0) {
		v119 = v147
		goto L17
	} else {
		goto L23
	}
L23:
	;
	goto L19
L24:
	;
	if v138 == int32(0) {
		v176 = v148
		goto L6
	} else {
		goto L25
	}
L25:
	;
	v199 = int32(1)
	goto L1
L26:
	;
	goto L13
L27:
	;
	if v188 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v199 = int32(1)
	goto L1
L29:
	;
	goto L30
L30:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v192 = F_point_inside(m, v187, v191, v47)
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L20
	} else {
		goto L31
	}
L31:
	;
	v199 = base.B2i32(v192 != int32(0))
	goto L1
}
