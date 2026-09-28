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
	var v24 int32
	_ = v24
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
			v21 = F_palloc_mul(m, int32(4), v16)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v21
				v27 = F_palloc0_mul(m, int32(4), v16)
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
	var v37 int32
	_ = v37
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
							v37 = F_errdetail(m, int32(_a_F_populate_array_check_dimension_1), int32(0))
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
func F_populate_composite(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) int64 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v178 int64
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v190 int64
	_ = v190
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v207 int32
	_ = v207
	var v216 int64
	_ = v216
	v15 = m.G0
	v17 = v15 + int32(-64)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v19 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	if v56 != 0 {
		v190 = int64(0)
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v32 = F_lookup_rowtype_tupdesc(m, v29, v31)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v29 = v22
	goto L2
L4:
	;
	goto L5
L5:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v23 != v24 {
		v29 = v24
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v26 == v27 {
		v54 = v19
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v29 = v23
	goto L2
L8:
	;
	return int64(0)
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	F_FreeTupleDesc(m, v36)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v39 = int32(_a_F_populate_composite_0)
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_populate_composite[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_populate_composite[0])) = l2
	v43 = F_CreateTupleDescCopy(m, v32)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v43
	*(*int32)(unsafe.Add(mBase, _c_F_populate_composite[0])) = v40
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v32)+12))
	if v48 < int32(0) {
		v54 = v40
		goto L1
	} else {
		goto L15
	}
L15:
	;
	F_DecrTupleDescRefCount(m, v32)
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v54 = v40
	goto L1
L17:
	;
	m.G0 = v17 - int32(-64)
	return v216
L18:
	;
	v207 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v207)
	v216 = int64(0)
	goto L17
L19:
	;
	if l1 == int32(2249) {
		v216 = v190
		goto L17
	} else {
		goto L69
	}
L20:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	*(*uint8)(unsafe.Add(mBase, uint32(v17)+8)) = uint8(v57)
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
		v134 = v54
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
	*(*int64)(unsafe.Add(mBase, uint32(v17)+24)) = int64(309237645376)
	v70 = *(*int32)(unsafe.Add(mBase, _c_F_populate_composite[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v70
	v77 = F_hash_create(m, int32(_a_F_populate_composite_1), int64(100), v15+int32(-48), int32(1048))
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
	*(*int32)(unsafe.Add(mBase, uint32(v83)+36)) = int32(1495)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+12)) = int32(1496)
	*(*int32)(unsafe.Add(mBase, uint32(v83))) = v80
	*(*int32)(unsafe.Add(mBase, uint32(v83)+24)) = int32(1497)
	*(*int32)(unsafe.Add(mBase, uint32(v83)+20)) = int32(1498)
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
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v114
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
	v134 = v54
	goto L40
L42:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v59)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)))
	if v122&int32(536870912) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+12)) = v121
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
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = int32(_a_F_populate_composite_2)
	F_errmsg(m, v133, v17)
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
	v162 = F_populate_record(m, v158, l0, l3, l2, v15+int32(-56), int32(0))
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
	if v164 == int32(453) {
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
	v171 = F_populate_record(m, v168, l0, l3, l2, v15+int32(-56), l6)
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
	if v173 != int32(453) {
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
		v190 = v178
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
	v190 = v178
	goto L19
L69:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if l1 == v193 {
		v216 = v190
		goto L17
	} else {
		goto L70
	}
L70:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5))))
	v198 = F_domain_check_safe(m, v190, v195, l1, l0+int32(16), l2, l6)
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L8
	} else {
		goto L71
	}
L71:
	;
	if v198 != 0 {
		v216 = v190
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int64
	_ = v41
	var v42 int64
	_ = v42
	var v43 int64
	_ = v43
	var v44 int64
	_ = v44
	var v45 int64
	_ = v45
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v48 int64
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v97 int64
	_ = v97
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
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v137 int32
	_ = v137
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v235 int64
	_ = v235
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v278 int32
	_ = v278
	var v294 int32
	_ = v294
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v336 int32
	_ = v336
	var v359 int32
	_ = v359
	var v374 int32
	_ = v374
	var v375 int64
	_ = v375
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v408 int32
	_ = v408
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v450 int64
	_ = v450
	var v451 int64
	_ = v451
	var v455 int64
	_ = v455
	var v456 int32
	_ = v456
	var v463 int32
	_ = v463
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	v7 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if l2 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v19 + int32(32)
	return v490
L2:
	;
	if v22 != 0 {
		goto L15
	} else {
		goto L16
	}
L3:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	if v26 == int32(1) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	v31 = *(*int64)(unsafe.Add(mBase, uint32(v30)+8))
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v30)+808))
	if v32 != int64(0) {
		goto L8
	} else {
		goto L9
	}
L5:
	;
	goto L6
L6:
	;
	if v25 == int32(0) {
		v490 = l2
		goto L1
	} else {
		goto L12
	}
L7:
	;
	if v97 != int64(0) {
		goto L2
	} else {
		goto L11
	}
L8:
	;
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v30)+752))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v30)+728))
	v37 = *(*int64)(unsafe.Add(mBase, uint32(v30)+704))
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v30)+680))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v30)+656))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v30)+632))
	v41 = *(*int64)(unsafe.Add(mBase, uint32(v30)+608))
	v42 = *(*int64)(unsafe.Add(mBase, uint32(v30)+584))
	v43 = *(*int64)(unsafe.Add(mBase, uint32(v30)+560))
	v44 = *(*int64)(unsafe.Add(mBase, uint32(v30)+536))
	v45 = *(*int64)(unsafe.Add(mBase, uint32(v30)+512))
	v46 = *(*int64)(unsafe.Add(mBase, uint32(v30)+488))
	v47 = *(*int64)(unsafe.Add(mBase, uint32(v30)+464))
	v48 = *(*int64)(unsafe.Add(mBase, uint32(v30)+440))
	v49 = *(*int64)(unsafe.Add(mBase, uint32(v30)+416))
	v50 = *(*int64)(unsafe.Add(mBase, uint32(v30)+392))
	v51 = *(*int64)(unsafe.Add(mBase, uint32(v30)+368))
	v52 = *(*int64)(unsafe.Add(mBase, uint32(v30)+344))
	v53 = *(*int64)(unsafe.Add(mBase, uint32(v30)+320))
	v54 = *(*int64)(unsafe.Add(mBase, uint32(v30)+296))
	v55 = *(*int64)(unsafe.Add(mBase, uint32(v30)+272))
	v56 = *(*int64)(unsafe.Add(mBase, uint32(v30)+248))
	v57 = *(*int64)(unsafe.Add(mBase, uint32(v30)+224))
	v58 = *(*int64)(unsafe.Add(mBase, uint32(v30)+200))
	v59 = *(*int64)(unsafe.Add(mBase, uint32(v30)+176))
	v60 = *(*int64)(unsafe.Add(mBase, uint32(v30)+152))
	v61 = *(*int64)(unsafe.Add(mBase, uint32(v30)+128))
	v62 = *(*int64)(unsafe.Add(mBase, uint32(v30)+104))
	v63 = *(*int64)(unsafe.Add(mBase, uint32(v30)+80))
	v64 = *(*int64)(unsafe.Add(mBase, uint32(v30)+56))
	v65 = *(*int64)(unsafe.Add(mBase, uint32(v30)+32))
	v97 = v35 + (v36 + (v37 + (v38 + (v39 + (v40 + (v41 + (v42 + (v43 + (v44 + (v45 + (v46 + (v47 + (v48 + (v49 + (v50 + (v51 + (v52 + (v53 + (v54 + (v55 + (v56 + (v57 + (v58 + (v59 + (v60 + (v61 + (v62 + (v63 + (v64 + (v65 + v31))))))))))))))))))))))))))))))
	goto L10
L9:
	;
	v97 = v31
	goto L10
L10:
	;
	goto L7
L11:
	;
	v490 = l2
	goto L1
L12:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
	if v102&int32(268435455) == int32(0) {
		v490 = l2
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L2
L14:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v151 == v152 {
		goto L31
	} else {
		goto L32
	}
L15:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v108 == v21 {
		v148 = v22
		goto L14
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	v111 = v21 << (uint(int32(6)) % 32)
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
	*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v21
	*(*int64)(unsafe.Add(mBase, uint32(v114))) = int64(0)
	if base.Ui32(v111) <= base.Ui32(int32(1024)) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v114
	v148 = v114
	goto L14
L22:
	;
	if v111 == int32(0) {
		goto L21
	} else {
		goto L25
	}
L23:
	;
	v137 = v111
	goto L24
L24:
	;
	if v137 == int32(0) {
		goto L21
	} else {
		goto L29
	}
L25:
	;
	v125 = v114 + v113
	v127 = v114 + int32(16)
	if base.Ui32(v127) < base.Ui32(v125) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v129 = v125
	goto L28
L27:
	;
	v129 = v127
	goto L28
L28:
	;
	v137 = (v129-v114-int32(13))&int32(-4) + int32(4)
	goto L24
L29:
	;
	base.MemoryFill(m, v114+int32(12), int32(0), v137)
	goto L21
L30:
	;
	v201 = F_palloc(m, v21<<(uint(int32(3))%32))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L19
	} else {
		goto L44
	}
L31:
	;
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v148)+4))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v154 == v155 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v160 = v21 << (uint(int32(6)) % 32)
	v162 = v160 | int32(12)
	if v148&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v162)) == int32(0) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L33
L35:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v148))) = v192
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v148)+8)) = v21
	*(*int32)(unsafe.Add(mBase, uint32(v148)+4)) = v194
	goto L30
L36:
	;
	v172 = v148 + v160 + int32(12)
	v174 = v148 + int32(4)
	if base.Ui32(v174) < base.Ui32(v172) {
		goto L39
	} else {
		goto L40
	}
L37:
	;
	goto L38
L38:
	;
	if v162 == int32(0) {
		goto L35
	} else {
		goto L43
	}
L39:
	;
	v176 = v172
	goto L41
L40:
	;
	v176 = v174
	goto L41
L41:
	;
	v181 = (v148^int32(-1)+v176)&int32(-4) + int32(4)
	if v181 == int32(0) {
		goto L35
	} else {
		goto L42
	}
L42:
	;
	base.MemoryFill(m, v148, int32(0), v181)
	goto L35
L43:
	;
	base.MemoryFill(m, v148, int32(0), v162)
	goto L35
L44:
	;
	v203 = F_palloc(m, v21)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
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
	v481 = F_heap_form_tuple(m, l0, v201, v203)
	mBase = m.M
	v482 = m.ExcPending
	if v482 != 0 {
		goto L19
	} else {
		goto L96
	}
L47:
	;
	v359 = int32(0)
	goto L64
L48:
	;
	if v21 <= int32(0) {
		goto L46
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+24)) = l2
	v324 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v324
	*(*uint16)(unsafe.Add(mBase, uint32(v19)+16)) = uint16(v324)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+8)) = int32(base.Ui32(v322) >> (uint(int32(2)) % 32))
	F_heap_deform_tuple(m, v19+int32(8), l0, v201, v203)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L19
	} else {
		goto L62
	}
L51:
	;
	v210 = v21 & int32(3)
	v211 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v21) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v217 = v211
	v229 = v7
	goto L55
L53:
	;
	v278 = v211
	goto L54
L54:
	;
	v294 = v278
	v308 = v7
	goto L59
L55:
	;
	v232 = int32(3)
	v235 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v201+v217<<(uint(v232)%32)))) = v235
	v238 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v217+v203))) = uint8(v238)
	v241 = v217 | v238
	*(*int64)(unsafe.Add(mBase, uint32(v201+v241<<(uint(v232)%32)))) = v235
	*(*uint8)(unsafe.Add(mBase, uint32(v203+v241))) = uint8(v238)
	v251 = v217 | int32(2)
	*(*int64)(unsafe.Add(mBase, uint32(v201+v251<<(uint(v232)%32)))) = v235
	*(*uint8)(unsafe.Add(mBase, uint32(v203+v251))) = uint8(v238)
	v261 = v217 | v232
	*(*int64)(unsafe.Add(mBase, uint32(v201+v261<<(uint(v232)%32)))) = v235
	*(*uint8)(unsafe.Add(mBase, uint32(v203+v261))) = uint8(v238)
	v270 = int32(4)
	v271 = v217 + v270
	v273 = v229 + v270
	if v273 != v21&int32(2147483644) {
		v217 = v271
		v229 = v273
		goto L55
	} else {
		goto L57
	}
L56:
	;
	if v210 == int32(0) {
		goto L47
	} else {
		goto L58
	}
L57:
	;
	goto L56
L58:
	;
	v278 = v271
	goto L54
L59:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v201+v294<<(uint(int32(3))%32)))) = int64(0)
	v315 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v294+v203))) = uint8(v315)
	v320 = v308 + v315
	if v320 != v210 {
		v294 = v294 + v315
		v308 = v320
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
	if v21 <= int32(0) {
		goto L46
	} else {
		goto L63
	}
L63:
	;
	goto L47
L64:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v375 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+16)) = v375
	*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v375
	v384 = l0 + v374<<(uint(int32(3))%32) + v359*int32(100)
	v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v384)+119)))
	if v385 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L46
L66:
	;
	v463 = v359 + int32(1)
	if v463 != v21 {
		v359 = v463
		goto L64
	} else {
		goto L95
	}
L67:
	;
	v389 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v359+v203))) = uint8(v389)
	goto L66
L68:
	;
	goto L69
L69:
	;
	v392 = v384 + int32(32)
	v393 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l4))))
	*(*uint8)(unsafe.Add(mBase, uint32(v19)+8)) = uint8(v393)
	v395 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v393 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	if v432 != 0 {
		goto L87
	} else {
		goto L88
	}
L71:
	;
	v399 = int32(0)
	v401 = F_hash_search(m, v395, v392, v399, v399)
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L19
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if v395 != 0 {
		goto L83
	} else {
		goto L84
	}
L74:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v415
	if v415 != 0 {
		goto L80
	} else {
		goto L81
	}
L75:
	;
	if v401 == int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = int32(11)
	v415 = int32(0)
	goto L74
L77:
	;
	goto L78
L78:
	;
	v408 = *(*int32)(unsafe.Add(mBase, uint32(v401)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v408
	if v408 == int32(11) {
		v415 = int32(0)
		goto L74
	} else {
		goto L79
	}
L79:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v401)+64))
	v415 = v413
	goto L74
L80:
	;
	v419 = int32(-1)
	goto L82
L81:
	;
	v419 = int32(0)
	goto L82
L82:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v419
	v432 = base.B2i32(v401 != int32(0))
	goto L70
L83:
	;
	v423 = F_strlen(m, v392)
	mBase = m.M
	v425 = F_getKeyJsonValueFromContainer(m, v395, v392, v423, int32(0))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L19
	} else {
		goto L86
	}
L84:
	;
	v428 = int32(0)
	goto L85
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+12)) = v428
	v432 = v428
	goto L70
L86:
	;
	v428 = v425
	goto L85
L87:
	;
	v433 = int32(0)
	goto L89
L88:
	;
	v433 = l2
	goto L89
L89:
	;
	if v433 != 0 {
		goto L66
	} else {
		goto L90
	}
L90:
	;
	v435 = v384 + int32(28)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v435)+76))
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v435)+68))
	v444 = v359 + v203
	v445 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v444))))
	if v445 != 0 {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v451 = int64(0)
	goto L93
L92:
	;
	v450 = *(*int64)(unsafe.Add(mBase, uint32(v201+v359<<(uint(int32(3))%32))))
	v451 = v450
	goto L93
L93:
	;
	v455 = F_populate_record_field(m, v148+int32(12)+v359<<(uint(int32(6))%32), v443, v436, v392, l3, v451, v19+int32(8), v444, l5, int32(0))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L19
	} else {
		goto L94
	}
L94:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v201+v359<<(uint(int32(3))%32)))) = v455
	goto L66
L95:
	;
	goto L65
L96:
	;
	F_pfree(m, v201)
	mBase = m.M
	v484 = m.ExcPending
	if v484 != 0 {
		goto L19
	} else {
		goto L97
	}
L97:
	;
	F_pfree(m, v203)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L19
	} else {
		goto L98
	}
L98:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v481)+16))
	v490 = v487
	goto L1
}
