package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_below(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 float64
	_ = v21
	var v24 float64
	_ = v24
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v42 float64
	_ = v42
	var v43 float64
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int64
	_ = v87
	var v89 int64
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum(m, v13)
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v104 != v14 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v21 = *(*float64)(unsafe.Add(mBase, uint32(v19)+8))
	v24 = *(*float64)(unsafe.Add(mBase, uint32(v14)+8))
	if base.F64_ge(base.F64_add(v21, float64(1e-06)), v24) == int32(0) {
		v100 = v2
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v28 = *(*float64)(unsafe.Add(mBase, uint32(v19)+24))
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
	if base.F64_le(v28, base.F64_add(v29, float64(1e-06))) == int32(0) {
		v100 = v2
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v35 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	v36 = *(*float64)(unsafe.Add(mBase, uint32(v19)+16))
	if base.F64_le(v35, base.F64_add(v36, float64(1e-06))) == int32(0) {
		v100 = v2
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v42 = *(*float64)(unsafe.Add(mBase, uint32(v19)+32))
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v14)+32))
	if base.F64_le(v42, base.F64_add(v43, float64(1e-06))) == int32(0) {
		v100 = v2
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v50 = v14 + int32(40)
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	v56 = v50 + v51<<(uint(int32(4))%32) - int32(16)
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v56)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v57
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v56)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v59
	if v51 <= int32(0) {
		v100 = int32(1)
		goto L3
	} else {
		goto L9
	}
L9:
	;
	v65 = v11 + int32(16)
	v72 = v2
	goto L10
L10:
	;
	v76 = v50 + v72<<(uint(int32(4))%32)
	v77 = *(*int64)(unsafe.Add(mBase, uint32(v76)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v77
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v79
	v82 = F_lseg_inside_poly(m, v11, v65, v19, int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L12
	}
L11:
	;
	v100 = v91
	goto L3
L12:
	;
	if v82 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v100 = int32(0)
	goto L3
L14:
	;
	goto L15
L15:
	;
	v87 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v11)+8)) = v87
	v89 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v11))) = v89
	v91 = int32(1)
	v93 = v72 + v91
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
	if v93 < v94 {
		v72 = v93
		goto L10
	} else {
		goto L16
	}
L16:
	;
	goto L11
L17:
	;
	F_pfree(m, v14)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v108 != v19 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_pfree(m, v19)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	m.G0 = v11 + int32(32)
	return v100
L24:
	;
	goto L23
}
func F_poly_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
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
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v156 int32
	_ = v156
	var v159 float64
	_ = v159
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v162 float64
	_ = v162
	var v167 int32
	_ = v167
	var v168 float64
	_ = v168
	var v173 int32
	_ = v173
	var v177 float64
	_ = v177
	var v183 float64
	_ = v183
	var v184 float64
	_ = v184
	var v185 float64
	_ = v185
	var v190 int32
	_ = v190
	var v194 float64
	_ = v194
	var v200 float64
	_ = v200
	var v201 float64
	_ = v201
	var v202 float64
	_ = v202
	var v204 float64
	_ = v204
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v213 float64
	_ = v213
	var v219 float64
	_ = v219
	var v221 int32
	_ = v221
	var v230 float64
	_ = v230
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v241 int32
	_ = v241
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v20 = int32(44)
	v21 = F___strchrnul(m, v19, v20)
	mBase = m.M
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v23 == v20 {
		v27 = v21
	} else {
		v27 = v2
	}
	if v27 != 0 {
		v29 = v2
		v32 = v27
		for {
			v41 = int32(1)
			v45 = int32(44)
			v46 = F___strchrnul(m, v32+v41, v45)
			mBase = m.M
			v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46))))
			if v48 == v45 {
				v52 = v46
			} else {
				v52 = int32(0)
			}
			if v52 != 0 {
				v29 = v29 + v41
				v32 = v52
				continue
			} else {
				break
			}
			break
		}
		v56 = int32(1)
		if v29&v56 != 0 {
			v60 = int32(-1)
		} else {
			v60 = (v29 + int32(2)) >> (uint(v56) % 32)
		}
		if int32(0) < v60 {
			v98 = v60 << (uint(int32(4)) % 32)
			v99 = base.I32_div_s(v98, v60)
			if base.B2i32(v99 == int32(16))&base.B2i32(v98 <= int32(2147483607)) == int32(0) {
				v107 = int32(0)
				v108 = F_errsave_start(m, v18)
				mBase = m.M
				v109 = m.ExcPending
				if v109 != 0 {
					return int32(0)
				} else {
					if v108 == int32(0) {
						v241 = v107
						m.G0 = v16 + int32(16)
						return v241
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_poly_in_0), int32(0))
							mBase = m.M
							v118 = m.ExcPending
							if v118 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v18, int32(_a_F_poly_in_1), int32(3438), int32(_a_F_poly_in_2))
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return int32(0)
								} else {
									v241 = v107
									m.G0 = v16 + int32(16)
									return v241
								}
							}
						}
					}
				}
			} else {
				v125 = v98 + int32(40)
				v126 = F_palloc0(m, v125)
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v126)+4)) = v60
					*(*int32)(unsafe.Add(mBase, uint32(v126))) = v125 << (uint(int32(2)) % 32)
					v132 = int32(0)
					v134 = v126 + int32(40)
					v139 = F_path_decode(m, v19, v132, v60, v134, v16+int32(15), v132, int32(_a_F_poly_in_3), v19, v18)
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int32(0)
					} else {
						if v139 == int32(0) {
							v143 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v143)
							v241 = int32(0)
						} else {
							v146 = *(*float64)(unsafe.Add(mBase, uint32(v126)+48))
							v147 = *(*float64)(unsafe.Add(mBase, uint32(v126)+40))
							v148 = *(*int32)(unsafe.Add(mBase, uint32(v126)+4))
							if v148 < int32(2) {
								v230 = v146
								v231 = v147
								v232 = v147
								v233 = v146
							} else {
								v156 = int32(1)
								v159 = v146
								v160 = v147
								v161 = v147
								v162 = v146
								for {
									v167 = v134 + v156<<(uint(int32(4))%32)
									v168 = *(*float64)(unsafe.Add(mBase, uint32(v167)))
									v173 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v168)&int64(9223372036854775807)))
									if v173 == int32(0) {
										if base.F64_gt(v161, v168) != 0 {
											v177 = v168
										} else {
											v177 = v161
										}
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v161)&int64(9223372036854775807)) {
											v183 = v168
										} else {
											v183 = v177
										}
										v184 = v183
									} else {
										v184 = v161
									}
									v185 = *(*float64)(unsafe.Add(mBase, uint32(v167)+8))
									v190 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v185)&int64(9223372036854775807)))
									if v190 == int32(0) {
										if base.F64_gt(v162, v185) != 0 {
											v194 = v185
										} else {
											v194 = v162
										}
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v162)&int64(9223372036854775807)) {
											v200 = v185
										} else {
											v200 = v194
										}
										v201 = v200
									} else {
										v201 = v162
									}
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v168)&int64(9223372036854775807)) {
										v202 = v168
									} else {
										v202 = v160
									}
									if base.F64_lt(v160, v168) != 0 {
										v204 = v168
									} else {
										v204 = v202
									}
									if base.Ui64(base.I64_reinterpret_f64(v160)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v210 = v204
									} else {
										v210 = v160
									}
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v185)&int64(9223372036854775807)) {
										v211 = v185
									} else {
										v211 = v159
									}
									if base.F64_lt(v159, v185) != 0 {
										v213 = v185
									} else {
										v213 = v211
									}
									if base.Ui64(base.I64_reinterpret_f64(v159)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v219 = v213
									} else {
										v219 = v159
									}
									v221 = v156 + int32(1)
									if v221 != v148 {
										v156 = v221
										v159 = v219
										v160 = v210
										v161 = v184
										v162 = v201
										continue
									} else {
										break
									}
									break
								}
								v230 = v219
								v231 = v210
								v232 = v184
								v233 = v201
							}
							*(*float64)(unsafe.Add(mBase, uint32(v126)+32)) = v233
							*(*float64)(unsafe.Add(mBase, uint32(v126)+8)) = v231
							*(*float64)(unsafe.Add(mBase, uint32(v126)+24)) = v232
							*(*float64)(unsafe.Add(mBase, uint32(v126)+16)) = v230
							v241 = v126
						}
						m.G0 = v16 + int32(16)
						return v241
					}
				}
			}
		} else {
			v76 = int32(0)
			v77 = F_errsave_start(m, v18)
			mBase = m.M
			v80 = m.ExcPending
			if v80 != 0 {
				return int32(0)
			} else {
				if v77 == int32(0) {
					v241 = v76
					m.G0 = v16 + int32(16)
					return v241
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v19
						*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(_a_F_poly_in_3)
						F_errmsg(m, int32(_a_F_poly_in_4), v16)
						mBase = m.M
						v91 = m.ExcPending
						if v91 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, v18, int32(_a_F_poly_in_1), int32(3429), int32(_a_F_poly_in_2))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return int32(0)
							} else {
								v241 = v76
								m.G0 = v16 + int32(16)
								return v241
							}
						}
					}
				}
			}
		}
	} else {
		v76 = int32(0)
		v77 = F_errsave_start(m, v18)
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return int32(0)
		} else {
			if v77 == int32(0) {
				v241 = v76
				m.G0 = v16 + int32(16)
				return v241
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v19
					*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(_a_F_poly_in_3)
					F_errmsg(m, int32(_a_F_poly_in_4), v16)
					mBase = m.M
					v91 = m.ExcPending
					if v91 != 0 {
						return int32(0)
					} else {
						F_errsave_finish(m, v18, int32(_a_F_poly_in_1), int32(3429), int32(_a_F_poly_in_2))
						mBase = m.M
						v96 = m.ExcPending
						if v96 != 0 {
							return int32(0)
						} else {
							v241 = v76
							m.G0 = v16 + int32(16)
							return v241
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
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
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
