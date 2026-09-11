package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_populate_array_array_end(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v12 <= v2 {
		v16 = v10 + int32(1)
		if base.Ui32(int32(2147483647)) <= base.Ui32(v10) {
			F_populate_array_report_expected_array(m, v11, v16)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				return int32(23)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v16
			v27 = v16 << (uint(int32(2)) % 32)
			v28 = F_palloc(m, v27)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v28
				v31 = F_palloc0(m, v27)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v31
					v34 = int32(3)
					v35 = v16 & v34
					v36 = int32(0)
					if base.Ui32(v34) <= base.Ui32(v10) {
						v41 = v36
						v47 = v2
						for {
							v50 = v41 << (uint(int32(2)) % 32)
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							v53 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v50+v51))) = v53
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v55+v50)+4)) = v53
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v59+v50)+8)) = v53
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v63+v50)+12)) = v53
							v67 = int32(4)
							v68 = v41 + v67
							v70 = v47 + v67
							if v70 != v16&int32(-4) {
								v41 = v68
								v47 = v70
								continue
							} else {
								break
							}
							break
						}
						v72 = v68
					} else {
						v72 = v36
					}
					if v35 != 0 {
						v80 = v72
						v85 = v2
						for {
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v88+v80<<(uint(int32(2))%32)))) = int32(-1)
							v94 = int32(1)
							v97 = v85 + v94
							if v97 != v35 {
								v80 = v80 + v94
								v85 = v97
								continue
							} else {
								break
							}
							break
						}
					} else {
					}
					v107 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
					v116 = v107
					if v10 < v116 {
						v119 = F_populate_array_check_dimension(m, v11, v10)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int32(0)
						} else {
							if v119 == int32(0) {
								v125 = int32(23)
							} else {
								v125 = int32(0)
							}
							return v125
						}
					} else {
						v125 = int32(0)
						return v125
					}
				}
			}
		}
	} else {
		v116 = v12
		if v10 < v116 {
			v119 = F_populate_array_check_dimension(m, v11, v10)
			mBase = m.M
			v120 = m.ExcPending
			if v120 != 0 {
				return int32(0)
			} else {
				if v119 == int32(0) {
					v125 = int32(23)
				} else {
					v125 = int32(0)
				}
				return v125
			}
		} else {
			v125 = int32(0)
			return v125
		}
	}
}
func F_populate_array_check_dimension(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = l1 << (uint(int32(2)) % 32)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(v7+v9)))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = v12 + v9
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	if v14 == int32(-1) {
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = v11
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v44 = v18
		v46 = l1 << (uint(int32(2)) % 32)
		v48 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v44+v46))) = v48
		if l1 <= v48 {
			v64 = int32(1)
			return v64
		} else {
			v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v56 = v53 + v46 - int32(4)
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
			v58 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v56))) = v57 + v58
			return v58
		}
	} else {
		if v14 == v11 {
			v44 = v7
			v46 = l1 << (uint(int32(2)) % 32)
			v48 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v44+v46))) = v48
			if l1 <= v48 {
				v64 = int32(1)
				return v64
			} else {
				v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v56 = v53 + v46 - int32(4)
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
				v58 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v56))) = v57 + v58
				return v58
			}
		} else {
			v20 = int32(0)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
			v22 = F_errsave_start(m, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				if v22 == int32(0) {
					v64 = v20
					return v64
				} else {
					F_errcode(m, int32(33685634))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_populate_array_check_dimension_0), int32(0))
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errdetail(m, int32(_a_F_populate_array_check_dimension_1), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, v21, int32(_a_F_populate_array_check_dimension_2), int32(2601), int32(_a_F_populate_array_check_dimension_3))
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v64 = v20
									return v64
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_populate_array_report_expected_array(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
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
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
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
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(80)
	m.G0 = v8
	if l1 <= v3 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v8 + int32(80)
	return
L2:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v14 = F_errsave_start(m, v13)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	F_initStringInfo(m, v8-int32(-64))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L5
	} else {
		goto L19
	}
L5:
	;
	return
L6:
	;
	if v12 != 0 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	if v14 == int32(0) {
		goto L1
	} else {
		goto L15
	}
L10:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L5
	} else {
		goto L11
	}
L11:
	;
	F_errmsg(m, int32(_a_F_populate_array_report_expected_array_0), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L5
	} else {
		goto L12
	}
L12:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v25
	F_errhint(m, int32(_a_F_populate_array_report_expected_array_1), v8)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L5
	} else {
		goto L13
	}
L13:
	;
	F_errsave_finish(m, v13, int32(_a_F_populate_array_report_expected_array_2), int32(2518), int32(_a_F_populate_array_report_expected_array_3))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L5
	} else {
		goto L14
	}
L14:
	;
	goto L1
L15:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L5
	} else {
		goto L16
	}
L16:
	;
	F_errmsg(m, int32(_a_F_populate_array_report_expected_array_0), int32(0))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L5
	} else {
		goto L17
	}
L17:
	;
	F_errsave_finish(m, v13, int32(_a_F_populate_array_report_expected_array_2), int32(2522), int32(_a_F_populate_array_report_expected_array_3))
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L5
	} else {
		goto L18
	}
L18:
	;
	goto L1
L19:
	;
	v56 = v3
	goto L20
L20:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v58+v56<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+48)) = v62
	F_appendStringInfo(m, v8-int32(-64), int32(_a_F_populate_array_report_expected_array_4), v8+int32(48))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L5
	} else {
		goto L22
	}
L21:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v76 = F_errsave_start(m, v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L5
	} else {
		goto L24
	}
L22:
	;
	v72 = v56 + int32(1)
	if v72 != l1 {
		v56 = v72
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if v74 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	F_errsave_finish(m, v75, int32(_a_F_populate_array_report_expected_array_2), v116, int32(_a_F_populate_array_report_expected_array_3))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L5
	} else {
		goto L37
	}
L26:
	;
	if v76 == int32(0) {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if v76 == int32(0) {
		goto L1
	} else {
		goto L33
	}
L29:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L5
	} else {
		goto L30
	}
L30:
	;
	F_errmsg(m, int32(_a_F_populate_array_report_expected_array_0), int32(0))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L5
	} else {
		goto L31
	}
L31:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v89
	*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v88
	F_errhint(m, int32(_a_F_populate_array_report_expected_array_5), v8+int32(32))
	mBase = m.M
	v96 = m.ExcPending
	if v96 != 0 {
		goto L5
	} else {
		goto L32
	}
L32:
	;
	v116 = int32(2542)
	goto L25
L33:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L5
	} else {
		goto L34
	}
L34:
	;
	F_errmsg(m, int32(_a_F_populate_array_report_expected_array_0), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L5
	} else {
		goto L35
	}
L35:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v8)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v107
	F_errhint(m, int32(_a_F_populate_array_report_expected_array_6), v8+int32(16))
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L5
	} else {
		goto L36
	}
L36:
	;
	v116 = int32(2548)
	goto L25
L37:
	;
	goto L1
}
func F_populate_composite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
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
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v280 int32
	_ = v280
	v14 = m.G0
	v16 = v14 + int32(-64)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v18 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	if v56 != 0 {
		v253 = int32(0)
		goto L20
	} else {
		goto L21
	}
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v31 = F_lookup_rowtype_tupdesc(m, v28, v30)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v28 = v21
	goto L2
L4:
	;
	goto L5
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 != v23 {
		v28 = v23
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v25 == v26 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v28 = v22
	goto L2
L8:
	;
	return int32(0)
L9:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v35 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_FreeTupleDesc(m, v35)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v38 = int32(_a_F_populate_composite_0)
	v39 = *(*int32)(unsafe.Add(mBase, _c_F_populate_composite[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_populate_composite[0])) = l2
	v42 = F_CreateTupleDescCopy(m, v31)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v42
	*(*int32)(unsafe.Add(mBase, _c_F_populate_composite[0])) = v39
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	if v47 < int32(0) {
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_DecrTupleDescRefCount(m, v31)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	goto L1
L17:
	;
	m.G0 = v16 - int32(-64)
	return v280
L18:
	;
	v280 = int32(0)
	goto L17
L19:
	;
	v270 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v270)
	goto L18
L20:
	;
	if l1 == int32(2249) {
		v280 = v253
		goto L17
	} else {
		goto L90
	}
L21:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v57)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v57 == int32(1) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	if l6 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L23:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v62 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	goto L25
L25:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	if v175 == int32(18) {
		goto L61
	} else {
		goto L62
	}
L26:
	;
	if v59&int32(3) == int32(0) {
		v88 = v59
		goto L31
	} else {
		goto L32
	}
L27:
	;
	v122 = v62
	goto L28
L28:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = int64(309237645376)
	v126 = *(*int32)(unsafe.Add(mBase, _c_F_populate_composite[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v126
	v133 = F_hash_create(m, int32(_a_F_populate_composite_1), int32(100), v14+int32(-48), int32(1048))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L8
	} else {
		goto L46
	}
L29:
	;
	v122 = v121
	goto L28
L30:
	;
	v121 = v113 - v59
	goto L29
L31:
	;
	v92 = v88
	goto L40
L32:
	;
	v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59))))
	if v72 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v121 = int32(0)
	goto L29
L34:
	;
	goto L35
L35:
	;
	v77 = v59
	goto L36
L36:
	;
	v81 = v77 + int32(1)
	if v81&int32(3) == int32(0) {
		v88 = v81
		goto L31
	} else {
		goto L38
	}
L37:
	;
	v113 = v81
	goto L30
L38:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if v86 != 0 {
		v77 = v81
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	v101 = int32(-2139062144)
	if (int32(16843008)-v98|v98)&v101 == v101 {
		v92 = v92 + int32(4)
		goto L40
	} else {
		goto L42
	}
L41:
	;
	v107 = v92
	goto L43
L42:
	;
	goto L41
L43:
	;
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
	if v111 != 0 {
		v107 = v107 + int32(1)
		goto L43
	} else {
		goto L45
	}
L44:
	;
	v113 = v107
	goto L30
L45:
	;
	goto L44
L46:
	;
	v136 = F_palloc0(m, int32(24))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v139 = F_palloc0(m, int32(40))
	mBase = m.M
	v140 = m.ExcPending
	if v140 != 0 {
		goto L8
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136)+8)) = v133
	*(*int32)(unsafe.Add(mBase, uint32(v136)+4)) = int32(_a_F_populate_composite_2)
	v146 = *(*int32)(unsafe.Add(mBase, _c_F_populate_composite[1]))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v146)+4))
	goto L49
L49:
	;
	v149 = F_makeJsonLexContextCstringLen(m, int32(0), v59, v122, v147, int32(1))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v149
	*(*int32)(unsafe.Add(mBase, uint32(v139)+36)) = int32(1370)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+12)) = int32(1371)
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v139)+24)) = int32(1372)
	*(*int32)(unsafe.Add(mBase, uint32(v139)+20)) = int32(1373)
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	v162 = F_pg_parse_json(m, v161, v139)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	if v162 != 0 {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	F_json_errsave_error(m, v162, v161, l6)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L8
	} else {
		goto L55
	}
L53:
	;
	v170 = v133
	goto L54
L54:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v136)))
	F_freeJsonLexContext(m, v171)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L8
	} else {
		goto L57
	}
L55:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v136)+8))
	F_hash_destroy(m, v166)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L8
	} else {
		goto L56
	}
L56:
	;
	v170 = int32(0)
	goto L54
L57:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v170
	goto L22
L58:
	;
	v202 = F_errsave_start(m, l6)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L8
	} else {
		goto L70
	}
L59:
	;
	v200 = int32(_a_F_populate_composite_3)
	goto L58
L60:
	;
	if v193&int32(268435456) != 0 {
		goto L67
	} else {
		goto L68
	}
L61:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
	if v179&int32(536870912) == int32(0) {
		v193 = v179
		goto L60
	} else {
		goto L64
	}
L62:
	;
	goto L63
L63:
	;
	v185 = int32(_a_F_populate_composite_4)
	if base.Ui32(v175) < base.Ui32(int32(4)) {
		v200 = v185
		goto L58
	} else {
		goto L65
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v178
	goto L22
L65:
	;
	switch v175 - int32(18) {
	case 0:
		goto L66
	default:
		goto L59
	case 14:
		v200 = v185
		goto L58
	}
L66:
	;
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	v193 = v191
	goto L60
L67:
	;
	v198 = int32(_a_F_populate_composite_4)
	goto L69
L68:
	;
	v198 = int32(_a_F_populate_composite_3)
	goto L69
L69:
	;
	v200 = v198
	goto L58
L70:
	;
	if v202 == int32(0) {
		goto L22
	} else {
		goto L71
	}
L71:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L8
	} else {
		goto L72
	}
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(_a_F_populate_composite_2)
	F_errmsg(m, v200, v16)
	mBase = m.M
	v212 = m.ExcPending
	if v212 != 0 {
		goto L8
	} else {
		goto L73
	}
L73:
	;
	F_errsave_finish(m, l6, int32(_a_F_populate_composite_5), int32(3020), int32(_a_F_populate_composite_6))
	mBase = m.M
	v217 = m.ExcPending
	if v217 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	goto L22
L75:
	;
	v245 = F_HeapTupleHeaderGetDatum(m, v244)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L87
	}
L76:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v229 = F_populate_record(m, v225, l0, l3, l2, v14+int32(-56), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L8
	} else {
		goto L79
	}
L77:
	;
	goto L78
L78:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v231 == int32(447) {
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v244 = v229
	goto L75
L80:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v234 != 0 {
		goto L19
	} else {
		goto L83
	}
L81:
	;
	goto L82
L82:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v238 = F_populate_record(m, v235, l0, l3, l2, v14+int32(-56), l6)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L8
	} else {
		goto L84
	}
L83:
	;
	goto L82
L84:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v240 != int32(447) {
		v244 = v238
		goto L75
	} else {
		goto L85
	}
L85:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v243 != 0 {
		goto L19
	} else {
		goto L86
	}
L86:
	;
	v244 = v238
	goto L75
L87:
	;
	if v57 == int32(0) {
		v253 = v245
		goto L20
	} else {
		goto L88
	}
L88:
	;
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	F_hash_destroy(m, v249)
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L8
	} else {
		goto L89
	}
L89:
	;
	v253 = v245
	goto L20
L90:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 == v260 {
		v280 = v253
		goto L17
	} else {
		goto L91
	}
L91:
	;
	v262 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	v265 = F_domain_check_safe(m, v253, v262, l1, l0+int32(16), l2, l6)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L8
	} else {
		goto L92
	}
L92:
	;
	if v265 != 0 {
		v280 = v253
		goto L17
	} else {
		goto L93
	}
L93:
	;
	v267 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v267)
	goto L18
}
func F_populate_record(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v298 int32
	_ = v298
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v342 int32
	_ = v342
	var v371 int32
	_ = v371
	var v388 int32
	_ = v388
	var v389 int64
	_ = v389
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v433 int32
	_ = v433
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v470 int32
	_ = v470
	var v473 int32
	_ = v473
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v493 int32
	_ = v493
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v533 int32
	_ = v533
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	v7 = int32(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v21 + int32(32)
	return v562
L2:
	;
	if v24 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v28 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v32)+4))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v32)+412))
	if v34 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	if v27 == int32(0) {
		v562 = l2
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v99 != 0 {
		goto L2
	} else {
		goto L11
	}
L8:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v32)+376))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32)+364))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v32)+352))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32)+340))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v32)+328))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v32)+316))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v32)+304))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v32)+292))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v32)+280))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v32)+268))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v32)+256))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v32)+244))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v32)+232))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+220))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v32)+208))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v32)+196))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v32)+184))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v32)+172))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v32)+160))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v32)+148))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v32)+136))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v32)+124))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v32)+112))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v32)+100))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v32)+88))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v32)+76))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v32-int32(-64))))
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v32)+52))
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v32)+40))
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v32)+28))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v32)+16))
	v99 = v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + (v57 + (v58 + (v59 + (v60 + (v63 + (v64 + (v65 + (v66 + (v67 + v33))))))))))))))))))))))))))))))
	goto L10
L9:
	;
	v99 = v33
	goto L10
L10:
	;
	goto L7
L11:
	;
	v562 = l2
	goto L1
L12:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
	if v102&int32(268435455) == int32(0) {
		v562 = l2
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L2
L14:
	;
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v153 == v154 {
		goto L31
	} else {
		goto L32
	}
L15:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v108 == v23 {
		v149 = v24
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v111 = v23 << (uint(int32(6)) % 32)
	v113 = v111 | int32(12)
	v114 = F_MemoryContextAlloc(m, l3, v113)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L17
L19:
	;
	return int32(0)
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v23
	*(*int64)(unsafe.Add(mBase, uint32(v114))) = int64(0)
	v122 = v114 + int32(12)
	if base.Ui32(int32(1024)) < base.Ui32(v111) {
		v141 = v111
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v114
	v149 = v114
	goto L14
L22:
	;
	v145 = F__emscripten_memset_bulkmem(m, v122, base.I32_extend8_s(int32(0)), v141)
	mBase = m.M
	goto L29
L23:
	;
	if v122&int32(3) != 0 {
		v141 = v111
		goto L22
	} else {
		goto L24
	}
L24:
	;
	if base.Ui32(v111+v122) <= base.Ui32(v122) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	v129 = v114 + v113
	v131 = v114 + int32(16)
	if base.Ui32(v131) < base.Ui32(v129) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v133 = v129
	goto L28
L27:
	;
	v133 = v131
	goto L28
L28:
	;
	v141 = (v133-v114-int32(13))&int32(-4) + int32(4)
	goto L22
L29:
	;
	goto L21
L30:
	;
	v200 = F_palloc(m, v23<<(uint(int32(2))%32))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L19
	} else {
		goto L45
	}
L31:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v156 == v157 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v160 = v23 << (uint(int32(6)) % 32)
	v162 = v160 | int32(12)
	if v149&int32(3) != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L33
L35:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v149))) = v191
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v149)+8)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v149)+4)) = v193
	goto L30
L36:
	;
	v188 = F__emscripten_memset_bulkmem(m, v149, base.I32_extend8_s(int32(0)), v162)
	mBase = m.M
	goto L44
L37:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v162) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	if base.Ui32(v162+v149) <= base.Ui32(v149) {
		goto L35
	} else {
		goto L39
	}
L39:
	;
	v174 = v149 + v160 + int32(12)
	v176 = v149 + int32(4)
	if base.Ui32(v176) < base.Ui32(v174) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v178 = v174
	goto L42
L41:
	;
	v178 = v176
	goto L42
L42:
	;
	v185 = F__emscripten_memset_bulkmem(m, v149, base.I32_extend8_s(int32(0)), (v149^int32(-1)+v178)&int32(-4)+int32(4))
	mBase = m.M
	goto L43
L43:
	;
	goto L35
L44:
	;
	goto L35
L45:
	;
	v202 = F_palloc(m, v23)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L19
	} else {
		goto L46
	}
L46:
	;
	if l2 == int32(0) {
		goto L49
	} else {
		goto L50
	}
L47:
	;
	v553 = F_heap_form_tuple(m, l0, v200, v202)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L19
	} else {
		goto L114
	}
L48:
	;
	if v23 <= int32(0) {
		goto L47
	} else {
		goto L64
	}
L49:
	;
	if v23 <= int32(0) {
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = l2
	v330 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v330
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+16)) = uint16(v330)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v21)+8)) = int32(base.Ui32(v328) >> (uint(int32(2)) % 32))
	F_heap_deform_tuple(m, v21+int32(8), l0, v200, v202)
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L19
	} else {
		goto L63
	}
L52:
	;
	v209 = v23 & int32(3)
	v210 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v23) {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v217 = v210
	v226 = v7
	goto L56
L54:
	;
	v278 = v210
	goto L55
L55:
	;
	if v209 == int32(0) {
		goto L48
	} else {
		goto L59
	}
L56:
	;
	v234 = int32(2)
	v237 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v200+v217<<(uint(v234)%32)))) = v237
	v240 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v217+v202))) = uint8(v240)
	v243 = v217 | v240
	*(*int32)(unsafe.Add(mBase, uint32(v200+v243<<(uint(v234)%32)))) = v237
	*(*uint8)(unsafe.Add(mBase, uint32(v202+v243))) = uint8(v240)
	v253 = v217 | v234
	*(*int32)(unsafe.Add(mBase, uint32(v200+v253<<(uint(v234)%32)))) = v237
	*(*uint8)(unsafe.Add(mBase, uint32(v202+v253))) = uint8(v240)
	v263 = v217 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v200+v263<<(uint(v234)%32)))) = v237
	*(*uint8)(unsafe.Add(mBase, uint32(v202+v263))) = uint8(v240)
	v272 = int32(4)
	v273 = v217 + v272
	v275 = v226 + v272
	if v275 != v23&int32(2147483644) {
		v217 = v273
		v226 = v275
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v278 = v273
	goto L55
L58:
	;
	goto L57
L59:
	;
	v298 = v278
	v309 = v210
	goto L60
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200+v298<<(uint(int32(2))%32)))) = int32(0)
	v321 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v298+v202))) = uint8(v321)
	v326 = v309 + v321
	if v326 != v209 {
		v298 = v298 + v321
		v309 = v326
		goto L60
	} else {
		goto L62
	}
L61:
	;
	goto L48
L62:
	;
	goto L61
L63:
	;
	goto L48
L64:
	;
	v371 = int32(0)
	goto L65
L65:
	;
	v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v389 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v21+int32(16)))) = v389
	*(*int64)(unsafe.Add(mBase, uint32(v21)+8)) = v389
	v398 = l0 + int32(20) + v388<<(uint(int32(4))%32) + v371*int32(100)
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v398)+91)))
	if v399 == int32(1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L47
L67:
	;
	v533 = v371 + int32(1)
	if v533 != v23 {
		v371 = v533
		goto L65
	} else {
		goto L113
	}
L68:
	;
	v403 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v371+v202))) = uint8(v403)
	goto L67
L69:
	;
	goto L70
L70:
	;
	v406 = v398 + int32(4)
	v407 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+8)) = uint8(v407)
	v409 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v407 == int32(1) {
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v503 != 0 {
		goto L105
	} else {
		goto L106
	}
L72:
	;
	v413 = int32(0)
	v415 = F_hash_search(m, v409, v406, v413, v413)
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L19
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	if v409 != 0 {
		goto L84
	} else {
		goto L85
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v429
	if v429 != 0 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	if v415 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = int32(11)
	v429 = int32(0)
	goto L75
L78:
	;
	goto L79
L79:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v415)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+20)) = v422
	if v422 == int32(11) {
		v429 = int32(0)
		goto L75
	} else {
		goto L80
	}
L80:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v415)+64))
	v429 = v427
	goto L75
L81:
	;
	v433 = int32(-1)
	goto L83
L82:
	;
	v433 = int32(0)
	goto L83
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+16)) = v433
	v503 = base.B2i32(v415 != int32(0))
	goto L71
L84:
	;
	if v406&int32(3) == int32(0) {
		v460 = v406
		goto L89
	} else {
		goto L90
	}
L85:
	;
	v498 = int32(0)
	goto L86
L86:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v498
	v503 = v498
	goto L71
L87:
	;
	v495 = F_getKeyJsonValueFromContainer(m, v409, v406, v493, int32(0))
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L19
	} else {
		goto L104
	}
L88:
	;
	v493 = v485 - v406
	goto L87
L89:
	;
	v464 = v460
	goto L98
L90:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v406))))
	if v444 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v493 = int32(0)
	goto L87
L92:
	;
	goto L93
L93:
	;
	v449 = v406
	goto L94
L94:
	;
	v453 = v449 + int32(1)
	if v453&int32(3) == int32(0) {
		v460 = v453
		goto L89
	} else {
		goto L96
	}
L95:
	;
	v485 = v453
	goto L88
L96:
	;
	v458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v453))))
	if v458 != 0 {
		v449 = v453
		goto L94
	} else {
		goto L97
	}
L97:
	;
	goto L95
L98:
	;
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v464)))
	v473 = int32(-2139062144)
	if (int32(16843008)-v470|v470)&v473 == v473 {
		v464 = v464 + int32(4)
		goto L98
	} else {
		goto L100
	}
L99:
	;
	v479 = v464
	goto L101
L100:
	;
	goto L99
L101:
	;
	v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v479))))
	if v483 != 0 {
		v479 = v479 + int32(1)
		goto L101
	} else {
		goto L103
	}
L102:
	;
	v485 = v479
	goto L88
L103:
	;
	goto L102
L104:
	;
	v498 = v495
	goto L86
L105:
	;
	v504 = int32(0)
	goto L107
L106:
	;
	v504 = l2
	goto L107
L107:
	;
	if v504 != 0 {
		goto L67
	} else {
		goto L108
	}
L108:
	;
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v398)+76))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v398)+68))
	v514 = v371 + v202
	v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v514))))
	if v515 != 0 {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v520 = int32(0)
	goto L111
L110:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v200+v371<<(uint(int32(2))%32))))
	v520 = v519
	goto L111
L111:
	;
	v524 = F_populate_record_field(m, v149+int32(12)+v371<<(uint(int32(6))%32), v506, v505, v406, l3, v520, v21+int32(8), v514, l5, int32(0))
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L19
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v200+v371<<(uint(int32(2))%32)))) = v524
	goto L67
L113:
	;
	goto L66
L114:
	;
	F_pfree(m, v200)
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L19
	} else {
		goto L115
	}
L115:
	;
	F_pfree(m, v202)
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L19
	} else {
		goto L116
	}
L116:
	;
	v559 = *(*int32)(unsafe.Add(mBase, uint32(v553)+16))
	v562 = v559
	goto L1
}
