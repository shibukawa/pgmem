package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecGetCommonChildSlotOps(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
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
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+103)))
	if v7 == int32(1) {
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+99)))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v6)+92))
		if v11 == int32(0) {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
			if v27 != 0 {
				if v10&int32(1) != 0 {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v27)+8))
					v39 = v30
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+103)))
					if v41 == int32(1) {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+99)))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v5)+92))
						if v45 != 0 {
							v61 = v45
							v62 = v44
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
							if v46 != 0 {
								v58 = v46
								v59 = v44
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
								v61 = v60
								v62 = v59
							} else {
								v61 = int32(_a_F_ExecGetCommonChildSlotOps_0)
								v62 = v44
							}
						}
						if v61 == v39 {
							v65 = v39
						} else {
							v65 = int32(0)
						}
						if v62&int32(1) != 0 {
							v69 = v65
						} else {
							v69 = int32(0)
						}
						v71 = v69
						return v71
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
						if v48 == int32(0) {
							return int32(0)
						} else {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
							v58 = v48
							v59 = int32(base.Ui32(v53&int32(16)) >> (uint(int32(4)) % 32))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
							v61 = v60
							v62 = v59
							if v61 == v39 {
								v65 = v39
							} else {
								v65 = int32(0)
							}
							if v62&int32(1) != 0 {
								v69 = v65
							} else {
								v69 = int32(0)
							}
							v71 = v69
							return v71
						}
					}
				} else {
					return int32(0)
				}
			} else {
				if v10&int32(1) != 0 {
					v39 = int32(_a_F_ExecGetCommonChildSlotOps_0)
					v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+103)))
					if v41 == int32(1) {
						v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+99)))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v5)+92))
						if v45 != 0 {
							v61 = v45
							v62 = v44
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
							if v46 != 0 {
								v58 = v46
								v59 = v44
								v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
								v61 = v60
								v62 = v59
							} else {
								v61 = int32(_a_F_ExecGetCommonChildSlotOps_0)
								v62 = v44
							}
						}
						if v61 == v39 {
							v65 = v39
						} else {
							v65 = int32(0)
						}
						if v62&int32(1) != 0 {
							v69 = v65
						} else {
							v69 = int32(0)
						}
						v71 = v69
						return v71
					} else {
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
						if v48 == int32(0) {
							return int32(0)
						} else {
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
							v58 = v48
							v59 = int32(base.Ui32(v53&int32(16)) >> (uint(int32(4)) % 32))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
							v61 = v60
							v62 = v59
							if v61 == v39 {
								v65 = v39
							} else {
								v65 = int32(0)
							}
							if v62&int32(1) != 0 {
								v69 = v65
							} else {
								v69 = int32(0)
							}
							v71 = v69
							return v71
						}
					}
				} else {
					return int32(0)
				}
			}
		} else {
			if v10&int32(1) != 0 {
				v39 = v11
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+103)))
				if v41 == int32(1) {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+99)))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v5)+92))
					if v45 != 0 {
						v61 = v45
						v62 = v44
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
						if v46 != 0 {
							v58 = v46
							v59 = v44
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
							v61 = v60
							v62 = v59
						} else {
							v61 = int32(_a_F_ExecGetCommonChildSlotOps_0)
							v62 = v44
						}
					}
					if v61 == v39 {
						v65 = v39
					} else {
						v65 = int32(0)
					}
					if v62&int32(1) != 0 {
						v69 = v65
					} else {
						v69 = int32(0)
					}
					v71 = v69
					return v71
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
					if v48 == int32(0) {
						return int32(0)
					} else {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
						v58 = v48
						v59 = int32(base.Ui32(v53&int32(16)) >> (uint(int32(4)) % 32))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
						v61 = v60
						v62 = v59
						if v61 == v39 {
							v65 = v39
						} else {
							v65 = int32(0)
						}
						if v62&int32(1) != 0 {
							v69 = v65
						} else {
							v69 = int32(0)
						}
						v71 = v69
						return v71
					}
				}
			} else {
				return int32(0)
			}
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
		if v18 == int32(0) {
			v71 = v2
			return v71
		} else {
			v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+4)))
			if v21&int32(16) == int32(0) {
				v71 = v2
				return v71
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
				v39 = v26
				v41 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+103)))
				if v41 == int32(1) {
					v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+99)))
					v45 = *(*int32)(unsafe.Add(mBase, uint32(v5)+92))
					if v45 != 0 {
						v61 = v45
						v62 = v44
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
						if v46 != 0 {
							v58 = v46
							v59 = v44
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
							v61 = v60
							v62 = v59
						} else {
							v61 = int32(_a_F_ExecGetCommonChildSlotOps_0)
							v62 = v44
						}
					}
					if v61 == v39 {
						v65 = v39
					} else {
						v65 = int32(0)
					}
					if v62&int32(1) != 0 {
						v69 = v65
					} else {
						v69 = int32(0)
					}
					v71 = v69
					return v71
				} else {
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
					if v48 == int32(0) {
						return int32(0)
					} else {
						v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+4)))
						v58 = v48
						v59 = int32(base.Ui32(v53&int32(16)) >> (uint(int32(4)) % 32))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
						v61 = v60
						v62 = v59
						if v61 == v39 {
							v65 = v39
						} else {
							v65 = int32(0)
						}
						if v62&int32(1) != 0 {
							v69 = v65
						} else {
							v69 = int32(0)
						}
						v71 = v69
						return v71
					}
				}
			}
		}
	}
}
func F_compute_common_attribute(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32, l11 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v15 = int32(_a_F_compute_common_attribute_0)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v21 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[0])))
	if base.B2i32(v18 == int32(0))|base.B2i32(v18 != v21) != 0 {
		v39 = v18
		v40 = v21
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L62
	} else {
		goto L111
	}
L2:
	;
	F_errorConflictingDefElem(m, l2, l0)
	mBase = m.M
	v304 = m.ExcPending
	if v304 != 0 {
		goto L62
	} else {
		goto L110
	}
L3:
	;
	return v300
L4:
	;
	v300 = int32(1)
	goto L3
L5:
	;
	if v39-v40 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L6:
	;
	goto L5
L7:
	;
	v24 = v14
	v25 = v15
	goto L8
L8:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25)+1)))
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	if v29 == int32(0) {
		v39 = v29
		v40 = v28
		goto L6
	} else {
		goto L10
	}
L9:
	;
	v39 = v29
	v40 = v28
	goto L6
L10:
	;
	v32 = int32(1)
	if v29 == v28 {
		v24 = v24 + v32
		v25 = v25 + v32
		goto L8
	} else {
		goto L11
	}
L11:
	;
	goto L9
L12:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v46 = int32(_a_F_compute_common_attribute_1)
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[1])))
	if base.B2i32(v49 == int32(0))|base.B2i32(v49 != v52) != 0 {
		v70 = v49
		v71 = v52
		goto L18
	} else {
		goto L19
	}
L15:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v44 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = l2
	goto L4
L17:
	;
	if v70-v71 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	goto L17
L19:
	;
	v55 = v14
	v56 = v46
	goto L20
L20:
	;
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+1)))
	if v60 == int32(0) {
		v70 = v60
		v71 = v59
		goto L18
	} else {
		goto L22
	}
L21:
	;
	v70 = v60
	v71 = v59
	goto L18
L22:
	;
	v63 = int32(1)
	if v60 == v59 {
		v55 = v55 + v63
		v56 = v56 + v63
		goto L20
	} else {
		goto L23
	}
L23:
	;
	goto L21
L24:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v77 = int32(_a_F_compute_common_attribute_2)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v83 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[2])))
	if base.B2i32(v80 == int32(0))|base.B2i32(v80 != v83) != 0 {
		v101 = v80
		v102 = v83
		goto L30
	} else {
		goto L31
	}
L27:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v75 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = l2
	goto L4
L29:
	;
	if v101-v102 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	goto L29
L31:
	;
	v86 = v14
	v87 = v77
	goto L32
L32:
	;
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+1)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v86)+1)))
	if v91 == int32(0) {
		v101 = v91
		v102 = v90
		goto L30
	} else {
		goto L34
	}
L33:
	;
	v101 = v91
	v102 = v90
	goto L30
L34:
	;
	v94 = int32(1)
	if v91 == v90 {
		v86 = v86 + v94
		v87 = v87 + v94
		goto L32
	} else {
		goto L35
	}
L35:
	;
	goto L33
L36:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v106 != 0 {
		goto L2
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v108 = int32(_a_F_compute_common_attribute_3)
	v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v114 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[3])))
	if base.B2i32(v111 == int32(0))|base.B2i32(v111 != v114) != 0 {
		v132 = v111
		v133 = v114
		goto L41
	} else {
		goto L42
	}
L39:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l2
	goto L4
L40:
	;
	if v132-v133 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	goto L40
L42:
	;
	v117 = v14
	v118 = v108
	goto L43
L43:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v118)+1)))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v117)+1)))
	if v122 == int32(0) {
		v132 = v122
		v133 = v121
		goto L41
	} else {
		goto L45
	}
L44:
	;
	v132 = v122
	v133 = v121
	goto L41
L45:
	;
	v125 = int32(1)
	if v122 == v121 {
		v117 = v117 + v125
		v118 = v118 + v125
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L50
	}
L48:
	;
	goto L49
L49:
	;
	v139 = int32(_a_F_compute_common_attribute_4)
	v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v145 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[4])))
	if base.B2i32(v142 == int32(0))|base.B2i32(v142 != v145) != 0 {
		v163 = v142
		v164 = v145
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v137 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v137 != 0 {
		goto L2
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = l2
	goto L4
L52:
	;
	if v163-v164 == int32(0) {
		goto L59
	} else {
		goto L60
	}
L53:
	;
	goto L52
L54:
	;
	v148 = v14
	v149 = v139
	goto L55
L55:
	;
	v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v149)+1)))
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v148)+1)))
	if v153 == int32(0) {
		v163 = v153
		v164 = v152
		goto L53
	} else {
		goto L57
	}
L56:
	;
	v163 = v153
	v164 = v152
	goto L53
L57:
	;
	v156 = int32(1)
	if v153 == v152 {
		v148 = v148 + v156
		v149 = v149 + v156
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v170 = F_lappend(m, v168, v169)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	v175 = int32(_a_F_compute_common_attribute_5)
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v181 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[5])))
	if base.B2i32(v178 == int32(0))|base.B2i32(v178 != v181) != 0 {
		v199 = v178
		v200 = v181
		goto L65
	} else {
		goto L66
	}
L62:
	;
	return int32(0)
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v170
	goto L4
L64:
	;
	if v199-v200 == int32(0) {
		goto L71
	} else {
		goto L72
	}
L65:
	;
	goto L64
L66:
	;
	v184 = v14
	v185 = v175
	goto L67
L67:
	;
	v188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+1)))
	v189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+1)))
	if v189 == int32(0) {
		v199 = v189
		v200 = v188
		goto L65
	} else {
		goto L69
	}
L68:
	;
	v199 = v189
	v200 = v188
	goto L65
L69:
	;
	v192 = int32(1)
	if v189 == v188 {
		v184 = v184 + v192
		v185 = v185 + v192
		goto L67
	} else {
		goto L70
	}
L70:
	;
	goto L68
L71:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	v206 = int32(_a_F_compute_common_attribute_6)
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v212 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[6])))
	if base.B2i32(v209 == int32(0))|base.B2i32(v209 != v212) != 0 {
		v230 = v209
		v231 = v212
		goto L77
	} else {
		goto L78
	}
L74:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	if v204 != 0 {
		goto L2
	} else {
		goto L75
	}
L75:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = l2
	goto L4
L76:
	;
	if v230-v231 == int32(0) {
		goto L83
	} else {
		goto L84
	}
L77:
	;
	goto L76
L78:
	;
	v215 = v14
	v216 = v206
	goto L79
L79:
	;
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v216)+1)))
	v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v220 == int32(0) {
		v230 = v220
		v231 = v219
		goto L77
	} else {
		goto L81
	}
L80:
	;
	v230 = v220
	v231 = v219
	goto L77
L81:
	;
	v223 = int32(1)
	if v220 == v219 {
		v215 = v215 + v223
		v216 = v216 + v223
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v237 = int32(_a_F_compute_common_attribute_7)
	v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[7])))
	if base.B2i32(v240 == int32(0))|base.B2i32(v240 != v243) != 0 {
		v261 = v240
		v262 = v243
		goto L89
	} else {
		goto L90
	}
L86:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v235 != 0 {
		goto L2
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = l2
	goto L4
L88:
	;
	if v261-v262 == int32(0) {
		goto L95
	} else {
		goto L96
	}
L89:
	;
	goto L88
L90:
	;
	v246 = v14
	v247 = v237
	goto L91
L91:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v247)+1)))
	v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v246)+1)))
	if v251 == int32(0) {
		v261 = v251
		v262 = v250
		goto L89
	} else {
		goto L93
	}
L92:
	;
	v261 = v251
	v262 = v250
	goto L89
L93:
	;
	v254 = int32(1)
	if v251 == v250 {
		v246 = v246 + v254
		v247 = v247 + v254
		goto L91
	} else {
		goto L94
	}
L94:
	;
	goto L92
L95:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L98
	}
L96:
	;
	goto L97
L97:
	;
	v268 = int32(0)
	v269 = int32(_a_F_compute_common_attribute_8)
	v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v275 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[8])))
	if base.B2i32(v272 == v268)|base.B2i32(v272 != v275) != 0 {
		v293 = v272
		v294 = v275
		goto L101
	} else {
		goto L102
	}
L98:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	if v266 != 0 {
		goto L2
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = l2
	goto L4
L100:
	;
	if v293-v294 != 0 {
		v300 = v268
		goto L3
	} else {
		goto L107
	}
L101:
	;
	goto L100
L102:
	;
	v278 = v14
	v279 = v269
	goto L103
L103:
	;
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279)+1)))
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v278)+1)))
	if v283 == int32(0) {
		v293 = v283
		v294 = v282
		goto L101
	} else {
		goto L105
	}
L104:
	;
	v293 = v283
	v294 = v282
	goto L101
L105:
	;
	v286 = int32(1)
	if v283 == v282 {
		v278 = v278 + v286
		v279 = v279 + v286
		goto L103
	} else {
		goto L106
	}
L106:
	;
	goto L104
L107:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L108
	}
L108:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	if v296 != 0 {
		goto L2
	} else {
		goto L109
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = l2
	goto L4
L110:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L111:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L62
	} else {
		goto L112
	}
L112:
	;
	F_errmsg(m, int32(_a_F_compute_common_attribute_9), int32(0))
	mBase = m.M
	v316 = m.ExcPending
	if v316 != 0 {
		goto L62
	} else {
		goto L113
	}
L113:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	F_parser_errposition(m, l0, v317)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L62
	} else {
		goto L114
	}
L114:
	;
	F_errfinish(m, int32(_a_F_compute_common_attribute_10), int32(612), int32(_a_F_compute_common_attribute_11))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L62
	} else {
		goto L115
	}
L115:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_select_common_typmod(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(-1)
L2:
	;
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v8 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
	v13 = F_exprType(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	if v13 != l1 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v18 = F_exprTypmod(m, v12)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L4
	} else {
		goto L7
	}
L7:
	;
	v20 = int32(1)
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v21 <= v20 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return v18
L9:
	;
	goto L10
L10:
	;
	v28 = v20
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v28<<(uint(int32(2))%32))))
	v35 = F_exprType(m, v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L4
	} else {
		goto L13
	}
L12:
	;
	return v18
L13:
	;
	if v35 != l1 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	v38 = F_exprTypmod(m, v34)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	if v38 != v18 {
		goto L1
	} else {
		goto L16
	}
L16:
	;
	v42 = v28 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v42 < v43 {
		v28 = v42
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L12
}
