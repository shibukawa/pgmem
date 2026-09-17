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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	v2 = int32(0)
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
	if v12 <= v2 {
		v16 = v10 + int32(1)
		if base.Ui32(int32(2147483646)) < base.Ui32(v10) {
			F_populate_array_report_expected_array(m, v11, v16)
			mBase = m.M
			v126 = m.ExcPending
			if v126 != 0 {
				return int32(0)
			} else {
				return int32(23)
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v11)+28)) = v16
			v21 = v16 << (uint(int32(2)) % 32)
			v22 = F_palloc(m, v21)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v22
				v27 = F_palloc0(m, v21)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v11)+24)) = v27
					v30 = int32(3)
					v31 = v16 & v30
					v32 = int32(0)
					if base.Ui32(v30) <= base.Ui32(v10) {
						v37 = v32
						v43 = v2
						for {
							v46 = v37 << (uint(int32(2)) % 32)
							v47 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							v49 = int32(-1)
							*(*int32)(unsafe.Add(mBase, uint32(v46+v47))) = v49
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v51+v46)+4)) = v49
							v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v55+v46)+8)) = v49
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v59+v46)+12)) = v49
							v63 = int32(4)
							v64 = v37 + v63
							v66 = v43 + v63
							if v66 != v16&int32(-4) {
								v37 = v64
								v43 = v66
								continue
							} else {
								break
							}
							break
						}
						if v31 == int32(0) {
						} else {
							v70 = v64
							v78 = v70
							v85 = v2
							for {
								v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v86+v78<<(uint(int32(2))%32)))) = int32(-1)
								v92 = int32(1)
								v95 = v85 + v92
								if v95 != v31 {
									v78 = v78 + v92
									v85 = v95
									continue
								} else {
									break
								}
								break
							}
						}
					} else {
						v70 = v32
						v78 = v70
						v85 = v2
						for {
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v86+v78<<(uint(int32(2))%32)))) = int32(-1)
							v92 = int32(1)
							v95 = v85 + v92
							if v95 != v31 {
								v78 = v78 + v92
								v85 = v95
								continue
							} else {
								break
							}
							break
						}
					}
					v105 = *(*int32)(unsafe.Add(mBase, uint32(v11)+28))
					v114 = v105
					if v10 < v114 {
						v117 = F_populate_array_check_dimension(m, v11, v10)
						mBase = m.M
						v118 = m.ExcPending
						if v118 != 0 {
							return int32(0)
						} else {
							if v117 == int32(0) {
								v123 = int32(23)
							} else {
								v123 = int32(0)
							}
							return v123
						}
					} else {
						v123 = int32(0)
						return v123
					}
				}
			}
		}
	} else {
		v114 = v12
		if v10 < v114 {
			v117 = F_populate_array_check_dimension(m, v11, v10)
			mBase = m.M
			v118 = m.ExcPending
			if v118 != 0 {
				return int32(0)
			} else {
				if v117 == int32(0) {
					v123 = int32(23)
				} else {
					v123 = int32(0)
				}
				return v123
			}
		} else {
			v123 = int32(0)
			return v123
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
		if v11 == v14 {
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
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
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
		v185 = int32(0)
		goto L19
	} else {
		goto L20
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
		v53 = v18
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
		v53 = v39
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
	v53 = v39
	goto L1
L17:
	;
	m.G0 = v16 - int32(-64)
	return v209
L18:
	;
	v205 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v205)
	v209 = int32(0)
	goto L17
L19:
	;
	if l1 == int32(2249) {
		v209 = v185
		goto L17
	} else {
		goto L69
	}
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)) = uint8(v57)
	v59 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v57 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	if l6 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L22:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	if v62 < int32(0) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
	switch v120 {
	case 0, 1, 2, 3, 32:
		v133 = int32(_a_F_populate_composite_3)
		v134 = v53
		goto L40
	default:
		goto L41
	case 18:
		goto L42
	}
L25:
	;
	v65 = F_strlen(m, v59)
	mBase = m.M
	v66 = v65
	goto L27
L26:
	;
	v66 = v62
	goto L27
L27:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+32)) = int64(309237645376)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_populate_composite[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v70
	v77 = F_hash_create(m, int32(_a_F_populate_composite_1), int32(100), v14+int32(-48), int32(1048))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L28
	}
L28:
	;
	v80 = F_palloc0(m, int32(24))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L8
	} else {
		goto L29
	}
L29:
	;
	v83 = F_palloc0(m, int32(40))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L8
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+8)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v80)+4)) = int32(_a_F_populate_composite_2)
	v90 = *(*int32)(unsafe.Add(mBase, _c_F_populate_composite[1]))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	goto L31
L31:
	;
	v93 = F_makeJsonLexContextCstringLen(m, int32(0), v59, v66, v91, int32(1))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v93
	*(*int32)(unsafe.Add(mBase, uint32(v83)+36)) = int32(1355)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = int32(1356)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = int32(1357)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = int32(1358)
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v106 = F_pg_parse_json(m, v105, v83)
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L8
	} else {
		goto L33
	}
L33:
	;
	if v106 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	F_json_errsave_error(m, v106, v105, l6)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L8
	} else {
		goto L37
	}
L35:
	;
	v114 = v77
	goto L36
L36:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	F_freeJsonLexContext(m, v115)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L39
	}
L37:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v80)+8))
	F_hash_destroy(m, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L8
	} else {
		goto L38
	}
L38:
	;
	v114 = int32(0)
	goto L36
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v114
	v153 = v114
	goto L21
L40:
	;
	v135 = F_errsave_start(m, l6)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L8
	} else {
		goto L49
	}
L41:
	;
	v133 = int32(_a_F_populate_composite_6)
	v134 = v53
	goto L40
L42:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v122&int32(536870912) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v121
	v153 = v121
	goto L21
L44:
	;
	goto L45
L45:
	;
	if v122&int32(268435456) != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v130 = int32(_a_F_populate_composite_3)
	goto L48
L47:
	;
	v130 = int32(_a_F_populate_composite_6)
	goto L48
L48:
	;
	v133 = v130
	v134 = v121
	goto L40
L49:
	;
	if v135 == int32(0) {
		v153 = v134
		goto L21
	} else {
		goto L50
	}
L50:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v141 = m.ExcPending
	if v141 != 0 {
		goto L8
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = int32(_a_F_populate_composite_2)
	F_errmsg(m, v133, v16)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L8
	} else {
		goto L52
	}
L52:
	;
	F_errsave_finish(m, l6, int32(_a_F_populate_composite_4), int32(3020), int32(_a_F_populate_composite_5))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L8
	} else {
		goto L53
	}
L53:
	;
	v153 = v134
	goto L21
L54:
	;
	v178 = F_HeapTupleHeaderGetDatum(m, v177)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L8
	} else {
		goto L66
	}
L55:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v162 = F_populate_record(m, v158, l0, l3, l2, v14+int32(-56), int32(0))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L8
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v164 == int32(447) {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v177 = v162
	goto L54
L59:
	;
	v167 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v167 != 0 {
		goto L18
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v171 = F_populate_record(m, v168, l0, l3, l2, v14+int32(-56), l6)
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L8
	} else {
		goto L63
	}
L62:
	;
	goto L61
L63:
	;
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v173 != int32(447) {
		v177 = v171
		goto L54
	} else {
		goto L64
	}
L64:
	;
	v176 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+4)))
	if v176 != 0 {
		goto L18
	} else {
		goto L65
	}
L65:
	;
	v177 = v171
	goto L54
L66:
	;
	if v57 == int32(0) {
		v185 = v178
		goto L19
	} else {
		goto L67
	}
L67:
	;
	F_hash_destroy(m, v153)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L8
	} else {
		goto L68
	}
L68:
	;
	v185 = v178
	goto L19
L69:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 == v192 {
		v209 = v185
		goto L17
	} else {
		goto L70
	}
L70:
	;
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	v197 = F_domain_check_safe(m, v185, v194, l1, l0+int32(16), l2, l6)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	if v197 != 0 {
		v209 = v185
		goto L17
	} else {
		goto L72
	}
L72:
	;
	goto L18
}
func F_populate_record(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v132 int32
	_ = v132
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v220 int32
	_ = v220
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v245 int32
	_ = v245
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v272 int32
	_ = v272
	var v288 int32
	_ = v288
	var v296 int32
	_ = v296
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v329 int32
	_ = v329
	var v351 int32
	_ = v351
	var v365 int32
	_ = v365
	var v366 int64
	_ = v366
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v454 int32
	_ = v454
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	v7 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(32)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v18 + int32(32)
	return v480
L2:
	;
	if v21 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v25 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v29)+412))
	if v31 != 0 {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	if v24 == int32(0) {
		v480 = l2
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v94 != 0 {
		goto L2
	} else {
		goto L11
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v29)+376))
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v29)+364))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v29)+352))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v29)+340))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v29)+328))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v29)+316))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v29)+304))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v29)+292))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v29)+280))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v29)+268))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v29)+256))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v29)+244))
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v29)+232))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v29)+220))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v29)+208))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v29)+196))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v29)+184))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v29)+172))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v29)+160))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v29)+148))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v29)+136))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v29)+124))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v29)+112))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v29)+100))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v29)+88))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v29)+76))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v29)+64))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v29)+52))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v29)+40))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v29)+28))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v29)+16))
	v94 = v32 + (v33 + (v34 + (v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + (v57 + (v58 + (v59 + (v60 + (v61 + (v62 + v30))))))))))))))))))))))))))))))
	goto L10
L9:
	;
	v94 = v30
	goto L10
L10:
	;
	goto L7
L11:
	;
	v480 = l2
	goto L1
L12:
	;
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v97&int32(268435455) == int32(0) {
		v480 = l2
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L2
L14:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v146 == v147 {
		goto L31
	} else {
		goto L32
	}
L15:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	if v103 == v20 {
		v143 = v21
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v106 = v20 << (uint(int32(6)) % 32)
	v108 = v106 | int32(12)
	v109 = F_MemoryContextAlloc(m, l3, v108)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v20
	*(*int64)(unsafe.Add(mBase, uint32(v109))) = int64(0)
	if base.Ui32(v106) <= base.Ui32(int32(1024)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v109
	v143 = v109
	goto L14
L22:
	;
	if v106 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	v132 = v106
	goto L24
L24:
	;
	if v132 == int32(0) {
		goto L21
	} else {
		goto L29
	}
L25:
	;
	v120 = v109 + v108
	v122 = v109 + int32(16)
	if base.Ui32(v122) < base.Ui32(v120) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v124 = v120
	goto L28
L27:
	;
	v124 = v122
	goto L28
L28:
	;
	v132 = (v124-v109-int32(13))&int32(-4) + int32(4)
	goto L24
L29:
	;
	base.MemoryFill(m, v109+int32(12), int32(0), v132)
	goto L21
L30:
	;
	v196 = F_palloc(m, v20<<(uint(int32(2))%32))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L19
	} else {
		goto L44
	}
L31:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v143)+4))
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v149 == v150 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v155 = v20 << (uint(int32(6)) % 32)
	v157 = v155 | int32(12)
	if v143&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v157)) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L33
L35:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v187
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v143)+8)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v143)+4)) = v189
	goto L30
L36:
	;
	v167 = v143 + v155 + int32(12)
	v169 = v143 + int32(4)
	if base.Ui32(v169) < base.Ui32(v167) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v157 == int32(0) {
		goto L35
	} else {
		goto L43
	}
L39:
	;
	v171 = v167
	goto L41
L40:
	;
	v171 = v169
	goto L41
L41:
	;
	v176 = (v143^int32(-1)+v171)&int32(-4) + int32(4)
	if v176 == int32(0) {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	base.MemoryFill(m, v143, int32(0), v176)
	goto L35
L43:
	;
	base.MemoryFill(m, v143, int32(0), v157)
	goto L35
L44:
	;
	v198 = F_palloc(m, v20)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L19
	} else {
		goto L45
	}
L45:
	;
	if l2 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	v471 = F_heap_form_tuple(m, l0, v196, v198)
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L19
	} else {
		goto L96
	}
L47:
	;
	v351 = int32(0)
	goto L64
L48:
	;
	if v20 <= int32(0) {
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = l2
	v317 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v317
	*(*uint16)(unsafe.Add(mBase, uint32(v18)+16)) = uint16(v317)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(base.Ui32(v315) >> (uint(int32(2)) % 32))
	F_heap_deform_tuple(m, v18+int32(8), l0, v196, v198)
	mBase = m.M
	v329 = m.ExcPending
	if v329 != 0 {
		goto L19
	} else {
		goto L62
	}
L51:
	;
	v205 = v20 & int32(3)
	v206 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v20) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v212 = v206
	v220 = v7
	goto L55
L53:
	;
	v272 = v206
	goto L54
L54:
	;
	v288 = v272
	v296 = int32(0)
	goto L59
L55:
	;
	v226 = int32(2)
	v229 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v196+v212<<(uint(v226)%32)))) = v229
	v232 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v212+v198))) = uint8(v232)
	v235 = v212 | v232
	*(*int32)(unsafe.Add(mBase, uint32(v196+v235<<(uint(v226)%32)))) = v229
	*(*uint8)(unsafe.Add(mBase, uint32(v198+v235))) = uint8(v232)
	v245 = v212 | v226
	*(*int32)(unsafe.Add(mBase, uint32(v196+v245<<(uint(v226)%32)))) = v229
	*(*uint8)(unsafe.Add(mBase, uint32(v198+v245))) = uint8(v232)
	v255 = v212 | int32(3)
	*(*int32)(unsafe.Add(mBase, uint32(v196+v255<<(uint(v226)%32)))) = v229
	*(*uint8)(unsafe.Add(mBase, uint32(v198+v255))) = uint8(v232)
	v264 = int32(4)
	v265 = v212 + v264
	v267 = v220 + v264
	if v267 != v20&int32(2147483644) {
		v212 = v265
		v220 = v267
		goto L55
	} else {
		goto L57
	}
L56:
	;
	if v205 == int32(0) {
		goto L47
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v272 = v265
	goto L54
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196+v288<<(uint(int32(2))%32)))) = int32(0)
	v308 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v288+v198))) = uint8(v308)
	v313 = v296 + v308
	if v313 != v205 {
		v288 = v288 + v308
		v296 = v313
		goto L59
	} else {
		goto L61
	}
L60:
	;
	goto L47
L61:
	;
	goto L60
L62:
	;
	if v20 <= int32(0) {
		goto L46
	} else {
		goto L63
	}
L63:
	;
	goto L47
L64:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v366 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v366
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v366
	v375 = l0 + v365<<(uint(int32(4))%32) + v351*int32(100)
	v376 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v375)+111)))
	if v376 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L46
L66:
	;
	v454 = v351 + int32(1)
	if v454 != v20 {
		v351 = v454
		goto L64
	} else {
		goto L95
	}
L67:
	;
	v380 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v351+v198))) = uint8(v380)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v383 = v375 + int32(24)
	v384 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	*(*uint8)(unsafe.Add(mBase, uint32(v18)+8)) = uint8(v384)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v384 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v423 != 0 {
		goto L87
	} else {
		goto L88
	}
L71:
	;
	v390 = int32(0)
	v392 = F_hash_search(m, v386, v383, v390, v390)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L19
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v386 != 0 {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v406
	if v406 != 0 {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	if v392 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = int32(11)
	v406 = int32(0)
	goto L74
L77:
	;
	goto L78
L78:
	;
	v399 = *(*int32)(unsafe.Add(mBase, uint32(v392)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v399
	if v399 == int32(11) {
		v406 = int32(0)
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v392)+64))
	v406 = v404
	goto L74
L80:
	;
	v410 = int32(-1)
	goto L82
L81:
	;
	v410 = int32(0)
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v410
	v423 = base.B2i32(v392 != int32(0))
	goto L70
L83:
	;
	v414 = F_strlen(m, v383)
	mBase = m.M
	v416 = F_getKeyJsonValueFromContainer(m, v386, v383, v414, int32(0))
	mBase = m.M
	v417 = m.ExcPending
	if v417 != 0 {
		goto L19
	} else {
		goto L86
	}
L84:
	;
	v419 = int32(0)
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v419
	v423 = v419
	goto L70
L86:
	;
	v419 = v416
	goto L85
L87:
	;
	v424 = int32(0)
	goto L89
L88:
	;
	v424 = l2
	goto L89
L89:
	;
	if v424 != 0 {
		goto L66
	} else {
		goto L90
	}
L90:
	;
	v426 = v375 + int32(20)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v426)+76))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v426)+68))
	v435 = v351 + v198
	v436 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v435))))
	if v436 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v442 = int32(0)
	goto L93
L92:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(v196+v351<<(uint(int32(2))%32))))
	v442 = v441
	goto L93
L93:
	;
	v446 = F_populate_record_field(m, v143+int32(12)+v351<<(uint(int32(6))%32), v434, v427, v383, l3, v442, v18+int32(8), v435, l5, int32(0))
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L19
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v196+v351<<(uint(int32(2))%32)))) = v446
	goto L66
L95:
	;
	goto L65
L96:
	;
	F_pfree(m, v196)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L19
	} else {
		goto L97
	}
L97:
	;
	F_pfree(m, v198)
	mBase = m.M
	v476 = m.ExcPending
	if v476 != 0 {
		goto L19
	} else {
		goto L98
	}
L98:
	;
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v471)+16))
	v480 = v477
	goto L1
}
