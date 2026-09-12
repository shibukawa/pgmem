package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_below(m *base.Module, l0 int32) int32 {
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
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v15)+32))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v8)+16))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v19 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v23 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.F64_gt(v17, v18)
						}
					} else {
						return base.F64_gt(v17, v18)
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v23 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return base.F64_gt(v17, v18)
					}
				} else {
					return base.F64_gt(v17, v18)
				}
			}
		}
	}
}
func F_poly_box(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int64
	_ = v11
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v9 = F_palloc(m, int32(32))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v4)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v11
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v4)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v13
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v4)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v15
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v17
			return v9
		}
	}
}
func F_poly_contain_pt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v12 = F_point_inside(m, v8, v9, v4+int32(40))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			return base.B2i32(v12 != int32(0))
		}
	}
}
func F_poly_contained(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v26 float64
	_ = v26
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v63 int64
	_ = v63
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v83 int64
	_ = v83
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v97 int64
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	v2 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v114 != v16 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v23 = *(*float64)(unsafe.Add(mBase, uint32(v21)+8))
	v26 = *(*float64)(unsafe.Add(mBase, uint32(v16)+8))
	if base.F64_ge(base.F64_add(v23, float64(1e-06)), v26) == int32(0) {
		v106 = v2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v30 = *(*float64)(unsafe.Add(mBase, uint32(v21)+24))
	v31 = *(*float64)(unsafe.Add(mBase, uint32(v16)+24))
	if base.F64_le(v30, base.F64_add(v31, float64(1e-06))) == int32(0) {
		v106 = v2
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v16)+16))
	v38 = *(*float64)(unsafe.Add(mBase, uint32(v21)+16))
	if base.F64_le(v37, base.F64_add(v38, float64(1e-06))) == int32(0) {
		v106 = v2
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v21)+32))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v16)+32))
	if base.F64_le(v44, base.F64_add(v45, float64(1e-06))) == int32(0) {
		v106 = v2
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v52 = v13 + int32(8)
	v54 = v16 + int32(40)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v60 = v54 + v55<<(uint(int32(4))%32) - int32(16)
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v60)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = v61
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v60)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v63
	if v55 <= int32(0) {
		v106 = int32(1)
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v69 = v13 + int32(16)
	v76 = v2
	goto L10
L10:
	;
	v82 = v54 + v76<<(uint(int32(4))%32)
	v83 = *(*int64)(unsafe.Add(mBase, uint32(v82)))
	*(*int64)(unsafe.Add(mBase, uint32(v69))) = v83
	v86 = v13 + int32(24)
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v82)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v86))) = v87
	v90 = F_lseg_inside_poly(m, v13, v69, v21, int32(0))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v106 = v99
	goto L3
L12:
	;
	if v90 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v106 = int32(0)
	goto L3
L14:
	;
	goto L15
L15:
	;
	v95 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
	*(*int64)(unsafe.Add(mBase, uint32(v52))) = v95
	v97 = *(*int64)(unsafe.Add(mBase, uint32(v69)))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v97
	v99 = int32(1)
	v101 = v76 + v99
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v101 < v102 {
		v76 = v101
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	F_pfree(m, v16)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v118 != v21 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_pfree(m, v21)
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	m.G0 = v13 + int32(32)
	return v106
L24:
	;
	goto L23
}
func F_poly_in(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v164 float64
	_ = v164
	var v165 float64
	_ = v165
	var v166 float64
	_ = v166
	var v167 float64
	_ = v167
	var v171 int32
	_ = v171
	var v172 float64
	_ = v172
	var v177 int32
	_ = v177
	var v181 float64
	_ = v181
	var v187 float64
	_ = v187
	var v188 float64
	_ = v188
	var v189 float64
	_ = v189
	var v191 float64
	_ = v191
	var v193 int64
	_ = v193
	var v197 float64
	_ = v197
	var v202 int32
	_ = v202
	var v206 float64
	_ = v206
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v214 float64
	_ = v214
	var v215 float64
	_ = v215
	var v217 float64
	_ = v217
	var v223 float64
	_ = v223
	var v225 int32
	_ = v225
	var v236 float64
	_ = v236
	var v237 float64
	_ = v237
	var v238 float64
	_ = v238
	var v239 float64
	_ = v239
	var v246 int32
	_ = v246
	v2 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = int32(44)
	v22 = F___strchrnul(m, v20, v21)
	mBase = m.M
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v24 == v21 {
		v28 = v22
	} else {
		v28 = v2
	}
	if v28 != 0 {
		v30 = v2
		v31 = v28
		for {
			v43 = int32(1)
			v47 = int32(44)
			v48 = F___strchrnul(m, v31+v43, v47)
			mBase = m.M
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
			if v50 == v47 {
				v54 = v48
			} else {
				v54 = int32(0)
			}
			if v54 != 0 {
				v30 = v30 + v43
				v31 = v54
				continue
			} else {
				break
			}
			break
		}
		v58 = int32(1)
		if v30&v58 != 0 {
			v62 = int32(-1)
		} else {
			v62 = (v30 + int32(2)) >> (uint(v58) % 32)
		}
		if int32(0) < v62 {
			v101 = v62 << (uint(int32(4)) % 32)
			v102 = base.I32_div_s(v101, v62)
			if base.B2i32(v102 == int32(16))&base.B2i32(v101 <= int32(2147483607)) == int32(0) {
				v110 = int32(0)
				v111 = F_errsave_start(m, v19)
				mBase = m.M
				v112 = m.ExcPending
				if v112 != 0 {
					return int32(0)
				} else {
					if v111 == int32(0) {
						v246 = v110
						m.G0 = v17 + int32(16)
						return v246
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v117 = m.ExcPending
						if v117 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(440753), int32(0))
							mBase = m.M
							v121 = m.ExcPending
							if v121 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v19, int32(493522), int32(3438), int32(278820))
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return int32(0)
								} else {
									v246 = v110
									m.G0 = v17 + int32(16)
									return v246
								}
							}
						}
					}
				}
			} else {
				v128 = v101 + int32(40)
				v129 = F_palloc0(m, v128)
				mBase = m.M
				v130 = m.ExcPending
				if v130 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v129)+4)) = v62
					*(*int32)(unsafe.Add(mBase, uint32(v129))) = v128 << (uint(int32(2)) % 32)
					v135 = int32(0)
					v137 = v129 + int32(40)
					v142 = F_path_decode(m, v20, v135, v62, v137, v17+int32(15), v135, int32(272921), v20, v19)
					mBase = m.M
					v143 = m.ExcPending
					if v143 != 0 {
						return int32(0)
					} else {
						if v142 == int32(0) {
							v146 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v146)
							v246 = int32(0)
						} else {
							v149 = *(*float64)(unsafe.Add(mBase, uint32(v129)+48))
							v150 = *(*float64)(unsafe.Add(mBase, uint32(v129)+40))
							v151 = *(*int32)(unsafe.Add(mBase, uint32(v129)+4))
							if v151 < int32(2) {
								v236 = v149
								v237 = v150
								v238 = v150
								v239 = v149
							} else {
								v157 = int32(1)
								v164 = v149
								v165 = v150
								v166 = v150
								v167 = v149
								for {
									v171 = v137 + v157<<(uint(int32(4))%32)
									v172 = *(*float64)(unsafe.Add(mBase, uint32(v171)))
									v177 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v172)&int64(9223372036854775807)))
									if v177 == int32(0) {
										if base.F64_lt(v172, v166) != 0 {
											v181 = v172
										} else {
											v181 = v166
										}
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v166)&int64(9223372036854775807)) {
											v187 = v172
										} else {
											v187 = v181
										}
										v188 = v187
									} else {
										v188 = v166
									}
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v172)&int64(9223372036854775807)) {
										v189 = v172
									} else {
										v189 = v165
									}
									if base.F64_gt(v172, v165) != 0 {
										v191 = v172
									} else {
										v191 = v189
									}
									v193 = int64(9223372036854775807)
									v197 = *(*float64)(unsafe.Add(mBase, uint32(v171)+8))
									v202 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v197)&v193))
									if v202 == int32(0) {
										if base.F64_lt(v197, v167) != 0 {
											v206 = v197
										} else {
											v206 = v167
										}
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v167)&int64(9223372036854775807)) {
											v212 = v197
										} else {
											v212 = v206
										}
										v213 = v212
									} else {
										v213 = v167
									}
									if base.Ui64(base.I64_reinterpret_f64(v165)&v193) < base.Ui64(int64(9218868437227405313)) {
										v214 = v191
									} else {
										v214 = v165
									}
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v197)&v193) {
										v215 = v197
									} else {
										v215 = v164
									}
									if base.F64_gt(v197, v164) != 0 {
										v217 = v197
									} else {
										v217 = v215
									}
									if base.Ui64(base.I64_reinterpret_f64(v164)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v223 = v217
									} else {
										v223 = v164
									}
									v225 = v157 + int32(1)
									if v225 != v151 {
										v157 = v225
										v164 = v223
										v165 = v214
										v166 = v188
										v167 = v213
										continue
									} else {
										break
									}
									break
								}
								v236 = v223
								v237 = v214
								v238 = v188
								v239 = v213
							}
							*(*float64)(unsafe.Add(mBase, uint32(v129)+32)) = v239
							*(*float64)(unsafe.Add(mBase, uint32(v129)+8)) = v237
							*(*float64)(unsafe.Add(mBase, uint32(v129)+24)) = v238
							*(*float64)(unsafe.Add(mBase, uint32(v129)+16)) = v236
							v246 = v129
						}
						m.G0 = v17 + int32(16)
						return v246
					}
				}
			}
		} else {
			v79 = int32(0)
			v80 = F_errsave_start(m, v19)
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int32(0)
			} else {
				if v80 == int32(0) {
					v246 = v79
					m.G0 = v17 + int32(16)
					return v246
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v88 = m.ExcPending
					if v88 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v20
						*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(272921)
						F_errmsg(m, int32(724727), v17)
						mBase = m.M
						v94 = m.ExcPending
						if v94 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v19, int32(493522), int32(3429), int32(278820))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v246 = v79
								m.G0 = v17 + int32(16)
								return v246
							}
						}
					}
				}
			}
		}
	} else {
		v79 = int32(0)
		v80 = F_errsave_start(m, v19)
		mBase = m.M
		v83 = m.ExcPending
		if v83 != 0 {
			return int32(0)
		} else {
			if v80 == int32(0) {
				v246 = v79
				m.G0 = v17 + int32(16)
				return v246
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(272921)
					F_errmsg(m, int32(724727), v17)
					mBase = m.M
					v94 = m.ExcPending
					if v94 != 0 {
						return int32(0)
					} else {
						F_errsave_finish(m, v19, int32(493522), int32(3429), int32(278820))
						mBase = m.M
						v99 = m.ExcPending
						if v99 != 0 {
							return int32(0)
						} else {
							v246 = v79
							m.G0 = v17 + int32(16)
							return v246
						}
					}
				}
			}
		}
	}
}
func F_poly_overleft(m *base.Module, l0 int32) int32 {
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
	var v17 float64
	_ = v17
	var v18 float64
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v13 = l0 + int32(28)
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = F_pg_detoast_datum(m, v14)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			v17 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v8)+8))
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v19 != v8 {
				F_pfree(m, v8)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int32(0)
				} else {
					v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					if v23 != v15 {
						F_pfree(m, v15)
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							return base.F64_ge(v17, v18)
						}
					} else {
						return base.F64_ge(v17, v18)
					}
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				if v23 != v15 {
					F_pfree(m, v15)
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						return base.F64_ge(v17, v18)
					}
				} else {
					return base.F64_ge(v17, v18)
				}
			}
		}
	}
}
