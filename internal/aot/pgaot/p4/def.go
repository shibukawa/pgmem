package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_buildDefItem(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
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
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v135
L2:
	;
	v125 = F_pstrdup(m, l0)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L8
	} else {
		goto L44
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_buildDefItem[0])) = int32(0)
	v18 = F_strtol(m, l1, v7+int32(12), int32(10))
	mBase = m.M
	goto L4
L4:
	;
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_buildDefItem[0]))
	if v20 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_buildDefItem[0])) = int32(0)
	v37 = F_strtod(m, l1, v7+int32(12))
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L8
	} else {
		goto L12
	}
L6:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v22 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v23 = F_pstrdup(m, l0)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v27 = F_makeInteger(m, v18)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v30 = F_makeDefElem(m, v23, v27, int32(-1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L8
	} else {
		goto L11
	}
L11:
	;
	v135 = v30
	goto L1
L12:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _c_F_buildDefItem[0]))
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = int32(_a_F_buildDefItem_0)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_buildDefItem[1])))
	if base.B2i32(v55 == int32(0))|base.B2i32(v55 != v58) != 0 {
		v76 = v55
		v77 = v58
		goto L21
	} else {
		goto L22
	}
L14:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v41))))
	if v42 != 0 {
		goto L13
	} else {
		goto L15
	}
L15:
	;
	v43 = F_pstrdup(m, l0)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v45 = F_pstrdup(m, l1)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L8
	} else {
		goto L17
	}
L17:
	;
	v47 = F_makeFloat(m, v45)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L8
	} else {
		goto L18
	}
L18:
	;
	v50 = F_makeDefElem(m, v43, v47, int32(-1))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L8
	} else {
		goto L19
	}
L19:
	;
	v135 = v50
	goto L1
L20:
	;
	if v76-v77 == int32(0) {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	goto L20
L22:
	;
	v61 = l1
	v62 = v52
	goto L23
L23:
	;
	v65 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
	if v66 == int32(0) {
		v76 = v66
		v77 = v65
		goto L21
	} else {
		goto L25
	}
L24:
	;
	v76 = v66
	v77 = v65
	goto L21
L25:
	;
	v69 = int32(1)
	if v66 == v65 {
		v61 = v61 + v69
		v62 = v62 + v69
		goto L23
	} else {
		goto L26
	}
L26:
	;
	goto L24
L27:
	;
	v81 = F_pstrdup(m, l0)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v89 = int32(_a_F_buildDefItem_1)
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	v95 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_buildDefItem[2])))
	if base.B2i32(v92 == int32(0))|base.B2i32(v92 != v95) != 0 {
		v113 = v92
		v114 = v95
		goto L34
	} else {
		goto L35
	}
L30:
	;
	v84 = F_makeBoolean(m, int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L8
	} else {
		goto L31
	}
L31:
	;
	v87 = F_makeDefElem(m, v81, v84, int32(-1))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v135 = v87
	goto L1
L33:
	;
	if v113-v114 != 0 {
		goto L2
	} else {
		goto L40
	}
L34:
	;
	goto L33
L35:
	;
	v98 = l1
	v99 = v89
	goto L36
L36:
	;
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+1)))
	v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v98)+1)))
	if v103 == int32(0) {
		v113 = v103
		v114 = v102
		goto L34
	} else {
		goto L38
	}
L37:
	;
	v113 = v103
	v114 = v102
	goto L34
L38:
	;
	v106 = int32(1)
	if v103 == v102 {
		v98 = v98 + v106
		v99 = v99 + v106
		goto L36
	} else {
		goto L39
	}
L39:
	;
	goto L37
L40:
	;
	v116 = F_pstrdup(m, l0)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L41
	}
L41:
	;
	v119 = F_makeBoolean(m, int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L8
	} else {
		goto L42
	}
L42:
	;
	v122 = F_makeDefElem(m, v116, v119, int32(-1))
	mBase = m.M
	v123 = m.ExcPending
	if v123 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v135 = v122
	goto L1
L44:
	;
	v127 = F_pstrdup(m, l1)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L8
	} else {
		goto L45
	}
L45:
	;
	v129 = F_makeString(m, v127)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L8
	} else {
		goto L46
	}
L46:
	;
	v132 = F_makeDefElem(m, v125, v129, int32(-1))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L8
	} else {
		goto L47
	}
L47:
	;
	v135 = v132
	goto L1
}
func F_defGetNumeric(m *base.Module, l0 int32) float64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 float64
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 float64
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		switch v10 - int32(465) {
		case 0:
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v39 = base.F64_convert_i32_s(v37)
			m.G0 = v7 + int32(32)
			return v39
		case 1:
			v13 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v14 = F_atof(m, v13)
			mBase = m.M
			v17 = m.ExcPending
			if v17 != 0 {
				return float64(0)
			} else {
				v39 = v14
				m.G0 = v7 + int32(32)
				return v39
			}
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return float64(0)
			} else {
				F_errcode(m, int32(16801924))
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return float64(0)
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v25
					F_errmsg(m, int32(_a_F_defGetNumeric_0), v7+int32(16))
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return float64(0)
					} else {
						F_errfinish(m, int32(_a_F_defGetNumeric_1), int32(85), int32(_a_F_defGetNumeric_2))
						mBase = m.M
						v36 = m.ExcPending
						if v36 != 0 {
							return float64(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return float64(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v50 = m.ExcPending
			if v50 != 0 {
				return float64(0)
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v51
				F_errmsg(m, int32(_a_F_defGetNumeric_0), v7)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return float64(0)
				} else {
					F_errfinish(m, int32(_a_F_defGetNumeric_1), int32(74), int32(_a_F_defGetNumeric_2))
					mBase = m.M
					v60 = m.ExcPending
					if v60 != 0 {
						return float64(0)
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
func F_defGetQualifiedName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		if v10 == int32(1) {
			v46 = v9
			m.G0 = v7 + int32(32)
			return v46
		} else {
			if v10 != int32(468) {
				if v10 == int32(68) {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
					v46 = v17
					m.G0 = v7 + int32(32)
					return v46
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v23 = m.ExcPending
					if v23 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v26 = m.ExcPending
						if v26 != 0 {
							return int32(0)
						} else {
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v27
							F_errmsg(m, int32(_a_F_defGetQualifiedName_0), v7+int32(16))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_defGetQualifiedName_1), int32(259), int32(_a_F_defGetQualifiedName_2))
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
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
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v7)+24)) = v9
				*(*int32)(unsafe.Add(mBase, uint32(v7)+28)) = v9
				v44 = F_list_make1_impl(m, int32(1), v7+int32(24))
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return int32(0)
				} else {
					v46 = v44
					m.G0 = v7 + int32(32)
					return v46
				}
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v58
				F_errmsg(m, int32(_a_F_defGetQualifiedName_3), v7)
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_defGetQualifiedName_1), int32(245), int32(_a_F_defGetQualifiedName_2))
					mBase = m.M
					v67 = m.ExcPending
					if v67 != 0 {
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
func F_defGetStreamingMode(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v107 int32
	_ = v107
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v128 int32
	_ = v128
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v173 int32
	_ = v173
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v217 int32
	_ = v217
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v7 + int32(16)
	return v265
L2:
	;
	v265 = int32(116)
	goto L1
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
	if v12 == int32(465) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L8
	} else {
		goto L80
	}
L5:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
	switch v16 {
	case 0:
		v265 = int32(102)
		goto L1
	case 1:
		goto L2
	default:
		goto L4
	}
L6:
	;
	goto L7
L7:
	;
	v17 = int32(102)
	v18 = F_defGetString(m, l0)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v25 = v18
	v26 = int32(_a_F_defGetStreamingMode_0)
	goto L11
L10:
	;
	if v63 == int32(0) {
		v265 = v17
		goto L1
	} else {
		goto L23
	}
L11:
	;
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	v30 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v29 == v30 {
		v52 = v29
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v63 = int32(0)
	goto L10
L13:
	;
	v54 = int32(1)
	if v52 != 0 {
		v25 = v25 + v54
		v26 = v26 + v54
		goto L11
	} else {
		goto L22
	}
L14:
	;
	if base.Ui32((v29-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v40 = v29 | int32(32)
	goto L17
L16:
	;
	v40 = v29
	goto L17
L17:
	;
	if base.Ui32((v30-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v49 = v30 | int32(32)
	goto L20
L19:
	;
	v49 = v30
	goto L20
L20:
	;
	if v40 == v49 {
		v52 = v40
		goto L13
	} else {
		goto L21
	}
L21:
	;
	v63 = v40 - v49
	goto L10
L22:
	;
	goto L12
L23:
	;
	v69 = v18
	v70 = int32(_a_F_defGetStreamingMode_1)
	goto L25
L24:
	;
	if v107 == int32(0) {
		v265 = v17
		goto L1
	} else {
		goto L37
	}
L25:
	;
	v73 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v73 == v74 {
		v96 = v73
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v107 = int32(0)
	goto L24
L27:
	;
	v98 = int32(1)
	if v96 != 0 {
		v69 = v69 + v98
		v70 = v70 + v98
		goto L25
	} else {
		goto L36
	}
L28:
	;
	if base.Ui32((v73-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v84 = v73 | int32(32)
	goto L31
L30:
	;
	v84 = v73
	goto L31
L31:
	;
	if base.Ui32((v74-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v93 = v74 | int32(32)
	goto L34
L33:
	;
	v93 = v74
	goto L34
L34:
	;
	if v84 == v93 {
		v96 = v84
		goto L27
	} else {
		goto L35
	}
L35:
	;
	v107 = v84 - v93
	goto L24
L36:
	;
	goto L26
L37:
	;
	v113 = v18
	v114 = int32(_a_F_defGetStreamingMode_2)
	goto L39
L38:
	;
	if v151 == int32(0) {
		goto L2
	} else {
		goto L51
	}
L39:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114))))
	if v117 == v118 {
		v140 = v117
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v151 = int32(0)
	goto L38
L41:
	;
	v142 = int32(1)
	if v140 != 0 {
		v113 = v113 + v142
		v114 = v114 + v142
		goto L39
	} else {
		goto L50
	}
L42:
	;
	if base.Ui32((v117-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v128 = v117 | int32(32)
	goto L45
L44:
	;
	v128 = v117
	goto L45
L45:
	;
	if base.Ui32((v118-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v137 = v118 | int32(32)
	goto L48
L47:
	;
	v137 = v118
	goto L48
L48:
	;
	if v128 == v137 {
		v140 = v128
		goto L41
	} else {
		goto L49
	}
L49:
	;
	v151 = v128 - v137
	goto L38
L50:
	;
	goto L40
L51:
	;
	v158 = v18
	v159 = int32(_a_F_defGetStreamingMode_3)
	goto L53
L52:
	;
	if v196 == int32(0) {
		v265 = int32(116)
		goto L1
	} else {
		goto L65
	}
L53:
	;
	v162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v158))))
	v163 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v162 == v163 {
		v185 = v162
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v196 = int32(0)
	goto L52
L55:
	;
	v187 = int32(1)
	if v185 != 0 {
		v158 = v158 + v187
		v159 = v159 + v187
		goto L53
	} else {
		goto L64
	}
L56:
	;
	if base.Ui32((v162-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v173 = v162 | int32(32)
	goto L59
L58:
	;
	v173 = v162
	goto L59
L59:
	;
	if base.Ui32((v163-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v182 = v163 | int32(32)
	goto L62
L61:
	;
	v182 = v163
	goto L62
L62:
	;
	if v173 == v182 {
		v185 = v173
		goto L55
	} else {
		goto L63
	}
L63:
	;
	v196 = v173 - v182
	goto L52
L64:
	;
	goto L54
L65:
	;
	v202 = v18
	v203 = int32(_a_F_defGetStreamingMode_4)
	goto L67
L66:
	;
	if v240 != 0 {
		goto L4
	} else {
		goto L79
	}
L67:
	;
	v206 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if v206 == v207 {
		v229 = v206
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v240 = int32(0)
	goto L66
L69:
	;
	v231 = int32(1)
	if v229 != 0 {
		v202 = v202 + v231
		v203 = v203 + v231
		goto L67
	} else {
		goto L78
	}
L70:
	;
	if base.Ui32((v206-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v217 = v206 | int32(32)
	goto L73
L72:
	;
	v217 = v206
	goto L73
L73:
	;
	if base.Ui32((v207-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v226 = v207 | int32(32)
	goto L76
L75:
	;
	v226 = v207
	goto L76
L76:
	;
	if v217 == v226 {
		v229 = v217
		goto L69
	} else {
		goto L77
	}
L77:
	;
	v240 = v217 - v226
	goto L66
L78:
	;
	goto L68
L79:
	;
	v265 = int32(112)
	goto L1
L80:
	;
	F_errcode(m, int32(16801924))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L81
	}
L81:
	;
	v251 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v7))) = v251
	F_errmsg(m, int32(_a_F_defGetStreamingMode_5), v7)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(_a_F_defGetStreamingMode_6), int32(2509), int32(_a_F_defGetStreamingMode_7))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L8
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_defGetString(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v5 = m.G0
	v7 = v5 - int32(48)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v9 != 0 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
		switch v10 - int32(465) {
		case 0:
			v47 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v7)+32)) = v47
			v52 = F_psprintf(m, int32(_a_F_defGetString_0), v7+int32(32))
			mBase = m.M
			v53 = m.ExcPending
			if v53 != 0 {
				return int32(0)
			} else {
				v73 = v52
				m.G0 = v7 + int32(48)
				return v73
			}
		case 1:
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v73 = v15
			m.G0 = v7 + int32(48)
			return v73
		case 2:
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+4)))
			if v18 != 0 {
				v19 = int32(_a_F_defGetString_1)
			} else {
				v19 = int32(_a_F_defGetString_2)
			}
			v73 = v19
			m.G0 = v7 + int32(48)
			return v73
		case 3:
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v9)+4))
			v73 = v20
			m.G0 = v7 + int32(48)
			return v73
		default:
			switch v10 - int32(68) {
			case 0:
				v21 = F_TypeNameToString(m, v9)
				mBase = m.M
				v24 = m.ExcPending
				if v24 != 0 {
					return int32(0)
				} else {
					v73 = v21
					m.G0 = v7 + int32(48)
					return v73
				}
			case 1, 2, 3, 4, 5, 6, 7, 8:
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
					*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v35
					F_errmsg_internal(m, int32(_a_F_defGetString_3), v7+int32(16))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_defGetString_4), int32(59), int32(_a_F_defGetString_5))
						mBase = m.M
						v46 = m.ExcPending
						if v46 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			case 9:
				v26 = F_pstrdup(m, int32(_a_F_defGetString_6))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					v73 = v26
					m.G0 = v7 + int32(48)
					return v73
				}
			default:
				if v10 == int32(1) {
					v71 = F_NameListToString(m, v9)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return int32(0)
					} else {
						v73 = v71
						m.G0 = v7 + int32(48)
						return v73
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
						v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
						*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v35
						F_errmsg_internal(m, int32(_a_F_defGetString_3), v7+int32(16))
						mBase = m.M
						v41 = m.ExcPending
						if v41 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_defGetString_4), int32(59), int32(_a_F_defGetString_5))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
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
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(16801924))
			mBase = m.M
			v60 = m.ExcPending
			if v60 != 0 {
				return int32(0)
			} else {
				v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v7))) = v61
				F_errmsg(m, int32(_a_F_defGetString_7), v7)
				mBase = m.M
				v65 = m.ExcPending
				if v65 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_defGetString_4), int32(41), int32(_a_F_defGetString_5))
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
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
