package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_poly_below(m *base.Module, l0 int32) int64 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+32))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+16))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.F64_gt(v14, v15))
						}
					} else {
						return base.I64_extend_i32_u(base.F64_gt(v14, v15))
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.F64_gt(v14, v15))
					}
				} else {
					return base.I64_extend_i32_u(base.F64_gt(v14, v15))
				}
			}
		}
	}
}
func F_poly_box(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v9 = F_palloc(m, int32(32))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int64(0)
		} else {
			v11 = *(*int64)(unsafe.Add(mBase, uint32(v4)+32))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v11
			v13 = *(*int64)(unsafe.Add(mBase, uint32(v4)+24))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v13
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v4)+16))
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v15
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v4)+8))
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v17
			return base.I64_extend_i32_u(v9)
		}
	}
}
func F_poly_contain_pt(m *base.Module, l0 int32) int64 {
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
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v12 = F_point_inside(m, v8, v9, v4+int32(40))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.B2i32(v12 != int32(0)))
		}
	}
}
func F_poly_contained(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 float64
	_ = v22
	var v25 float64
	_ = v25
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int64
	_ = v88
	var v90 int64
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v114 int64
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	v9 = int64(0)
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v15 = F_pg_detoast_datum(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v20 = F_pg_detoast_datum(m, v19)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	if v115 != v15 {
		goto L17
	} else {
		goto L18
	}
L4:
	;
	v22 = *(*float64)(unsafe.Add(mBase, uint32(v20)+8))
	v25 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
	if base.F64_ge(base.F64_add(v22, float64(1e-06)), v25) == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L5
	}
L5:
	;
	v29 = *(*float64)(unsafe.Add(mBase, uint32(v20)+24))
	v30 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
	if base.F64_le(v29, base.F64_add(v30, float64(1e-06))) == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v36 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v20)+16))
	if base.F64_le(v36, base.F64_add(v37, float64(1e-06))) == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L7
	}
L7:
	;
	v43 = *(*float64)(unsafe.Add(mBase, uint32(v20)+32))
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v15)+32))
	if base.F64_le(v43, base.F64_add(v44, float64(1e-06))) == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L8
	}
L8:
	;
	v51 = v15 + int32(40)
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	v57 = v51 + v52<<(uint(int32(4))%32) - int32(16)
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v57)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v58
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v57)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v60
	if int32(0) < v52 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v65 = v12 + int32(16)
	v72 = int32(0)
	goto L12
L10:
	;
	goto L11
L11:
	;
	v114 = int64(1)
	goto L3
L12:
	;
	v78 = v51 + v72<<(uint(int32(4))%32)
	v79 = *(*int64)(unsafe.Add(mBase, uint32(v78)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v65)+8)) = v79
	v81 = *(*int64)(unsafe.Add(mBase, uint32(v78)))
	*(*int64)(unsafe.Add(mBase, uint32(v65))) = v81
	v84 = F_lseg_inside_poly(m, v12, v65, v20, int32(0))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L1
	} else {
		goto L14
	}
L13:
	;
	goto L11
L14:
	;
	if v84 == int32(0) {
		v114 = v9
		goto L3
	} else {
		goto L15
	}
L15:
	;
	v88 = *(*int64)(unsafe.Add(mBase, uint32(v65)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v88
	v90 = *(*int64)(unsafe.Add(mBase, uint32(v65)))
	*(*int64)(unsafe.Add(mBase, uint32(v12))) = v90
	v93 = v72 + int32(1)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v15)+4))
	if v93 < v94 {
		v72 = v93
		goto L12
	} else {
		goto L16
	}
L16:
	;
	goto L13
L17:
	;
	F_pfree(m, v15)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L20
	}
L18:
	;
	goto L19
L19:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v119 != v20 {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	goto L19
L21:
	;
	F_pfree(m, v20)
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	m.G0 = v12 + int32(32)
	return v114
L24:
	;
	goto L23
}
func F_poly_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int64
	_ = v14
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
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
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
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
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
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v146 float64
	_ = v146
	var v147 float64
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v159 float64
	_ = v159
	var v160 float64
	_ = v160
	var v161 float64
	_ = v161
	var v162 float64
	_ = v162
	var v168 int32
	_ = v168
	var v169 float64
	_ = v169
	var v174 int32
	_ = v174
	var v178 float64
	_ = v178
	var v184 float64
	_ = v184
	var v185 float64
	_ = v185
	var v186 float64
	_ = v186
	var v191 int32
	_ = v191
	var v195 float64
	_ = v195
	var v201 float64
	_ = v201
	var v202 float64
	_ = v202
	var v203 float64
	_ = v203
	var v205 float64
	_ = v205
	var v211 float64
	_ = v211
	var v212 float64
	_ = v212
	var v214 float64
	_ = v214
	var v220 float64
	_ = v220
	var v222 int32
	_ = v222
	var v231 float64
	_ = v231
	var v232 float64
	_ = v232
	var v233 float64
	_ = v233
	var v234 float64
	_ = v234
	var v256 int64
	_ = v256
	v2 = int32(0)
	v14 = int64(0)
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
		v32 = v28
		v33 = v2
		for {
			v43 = int32(1)
			v47 = int32(44)
			v48 = F___strchrnul(m, v32+v43, v47)
			mBase = m.M
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
			if v50 == v47 {
				v54 = v48
			} else {
				v54 = int32(0)
			}
			if v54 != 0 {
				v32 = v54
				v33 = v33 + v43
				continue
			} else {
				break
			}
			break
		}
		v58 = int32(1)
		if v33&v58 != 0 {
			v62 = int32(-1)
		} else {
			v62 = (v33 + int32(2)) >> (uint(v58) % 32)
		}
		if int32(0) < v62 {
			v100 = v62 << (uint(int32(4)) % 32)
			v101 = base.I32_div_s(v100, v62)
			if base.B2i32(v101 == int32(16))&base.B2i32(v100 <= int32(2147483607)) == int32(0) {
				v109 = F_errsave_start(m, v19)
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int64(0)
				} else {
					if v109 == int32(0) {
						v256 = v14
						m.G0 = v17 + int32(16)
						return v256
					} else {
						F_errcode(m, int32(261))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_poly_in_0), int32(0))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int64(0)
							} else {
								F_errsave_finish(m, v19, int32(_a_F_poly_in_1), int32(3522), int32(_a_F_poly_in_2))
								mBase = m.M
								v124 = m.ExcPending
								if v124 != 0 {
									return int64(0)
								} else {
									v256 = v14
									m.G0 = v17 + int32(16)
									return v256
								}
							}
						}
					}
				}
			} else {
				v126 = v100 + int32(40)
				v127 = F_palloc0(m, v126)
				mBase = m.M
				v128 = m.ExcPending
				if v128 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v127)+4)) = v62
					*(*int32)(unsafe.Add(mBase, uint32(v127))) = v126 << (uint(int32(2)) % 32)
					v133 = int32(0)
					v135 = v127 + int32(40)
					v140 = F_path_decode(m, v20, v133, v62, v135, v17+int32(15), v133, int32(_a_F_poly_in_3), v20, v19)
					mBase = m.M
					v141 = m.ExcPending
					if v141 != 0 {
						return int64(0)
					} else {
						if v140 == int32(0) {
							v144 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v144)
							v256 = v14
						} else {
							v146 = *(*float64)(unsafe.Add(mBase, uint32(v127)+48))
							v147 = *(*float64)(unsafe.Add(mBase, uint32(v127)+40))
							v148 = *(*int32)(unsafe.Add(mBase, uint32(v127)+4))
							if v148 < int32(2) {
								v231 = v146
								v232 = v147
								v233 = v147
								v234 = v146
							} else {
								v155 = int32(1)
								v159 = v146
								v160 = v147
								v161 = v147
								v162 = v146
								for {
									v168 = v135 + v155<<(uint(int32(4))%32)
									v169 = *(*float64)(unsafe.Add(mBase, uint32(v168)))
									v174 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v169)&int64(9223372036854775807)))
									if v174 == int32(0) {
										if base.F64_gt(v161, v169) != 0 {
											v178 = v169
										} else {
											v178 = v161
										}
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v161)&int64(9223372036854775807)) {
											v184 = v169
										} else {
											v184 = v178
										}
										v185 = v184
									} else {
										v185 = v161
									}
									v186 = *(*float64)(unsafe.Add(mBase, uint32(v168)+8))
									v191 = base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v186)&int64(9223372036854775807)))
									if v191 == int32(0) {
										if base.F64_gt(v162, v186) != 0 {
											v195 = v186
										} else {
											v195 = v162
										}
										if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v162)&int64(9223372036854775807)) {
											v201 = v186
										} else {
											v201 = v195
										}
										v202 = v201
									} else {
										v202 = v162
									}
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v169)&int64(9223372036854775807)) {
										v203 = v169
									} else {
										v203 = v160
									}
									if base.F64_lt(v160, v169) != 0 {
										v205 = v169
									} else {
										v205 = v203
									}
									if base.Ui64(base.I64_reinterpret_f64(v160)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v211 = v205
									} else {
										v211 = v160
									}
									if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v186)&int64(9223372036854775807)) {
										v212 = v186
									} else {
										v212 = v159
									}
									if base.F64_lt(v159, v186) != 0 {
										v214 = v186
									} else {
										v214 = v212
									}
									if base.Ui64(base.I64_reinterpret_f64(v159)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
										v220 = v214
									} else {
										v220 = v159
									}
									v222 = v155 + int32(1)
									if v222 != v148 {
										v155 = v222
										v159 = v220
										v160 = v211
										v161 = v185
										v162 = v202
										continue
									} else {
										break
									}
									break
								}
								v231 = v220
								v232 = v211
								v233 = v185
								v234 = v202
							}
							*(*float64)(unsafe.Add(mBase, uint32(v127)+32)) = v234
							*(*float64)(unsafe.Add(mBase, uint32(v127)+8)) = v232
							*(*float64)(unsafe.Add(mBase, uint32(v127)+24)) = v233
							*(*float64)(unsafe.Add(mBase, uint32(v127)+16)) = v231
							v256 = base.I64_extend_i32_u(v127)
						}
						m.G0 = v17 + int32(16)
						return v256
					}
				}
			}
		} else {
			v79 = F_errsave_start(m, v19)
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int64(0)
			} else {
				if v79 == int32(0) {
					v256 = v14
					m.G0 = v17 + int32(16)
					return v256
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v87 = m.ExcPending
					if v87 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v20
						*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(_a_F_poly_in_3)
						F_errmsg(m, int32(_a_F_poly_in_4), v17)
						mBase = m.M
						v93 = m.ExcPending
						if v93 != 0 {
							return int64(0)
						} else {
							F_errsave_finish(m, v19, int32(_a_F_poly_in_1), int32(3513), int32(_a_F_poly_in_2))
							mBase = m.M
							v98 = m.ExcPending
							if v98 != 0 {
								return int64(0)
							} else {
								v256 = v14
								m.G0 = v17 + int32(16)
								return v256
							}
						}
					}
				}
			}
		}
	} else {
		v79 = F_errsave_start(m, v19)
		mBase = m.M
		v82 = m.ExcPending
		if v82 != 0 {
			return int64(0)
		} else {
			if v79 == int32(0) {
				v256 = v14
				m.G0 = v17 + int32(16)
				return v256
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v20
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(_a_F_poly_in_3)
					F_errmsg(m, int32(_a_F_poly_in_4), v17)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return int64(0)
					} else {
						F_errsave_finish(m, v19, int32(_a_F_poly_in_1), int32(3513), int32(_a_F_poly_in_2))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return int64(0)
						} else {
							v256 = v14
							m.G0 = v17 + int32(16)
							return v256
						}
					}
				}
			}
		}
	}
}
func F_poly_overleft(m *base.Module, l0 int32) int64 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = F_pg_detoast_datum(m, v11)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
			v15 = *(*float64)(unsafe.Add(mBase, uint32(v7)+8))
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			if v16 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int64(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
					if v20 != v12 {
						F_pfree(m, v12)
						mBase = m.M
						v23 = m.ExcPending
						if v23 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.F64_ge(v14, v15))
						}
					} else {
						return base.I64_extend_i32_u(base.F64_ge(v14, v15))
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
				if v20 != v12 {
					F_pfree(m, v12)
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.F64_ge(v14, v15))
					}
				} else {
					return base.I64_extend_i32_u(base.F64_ge(v14, v15))
				}
			}
		}
	}
}
