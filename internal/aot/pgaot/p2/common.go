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
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v315 int32
	_ = v315
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v15 = int32(_a_F_compute_common_attribute_0)
	v18 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[0])))
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v19 == int32(0) {
		v38 = v18
		v39 = v19
		goto L6
	} else {
		goto L7
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L67
	} else {
		goto L120
	}
L2:
	;
	F_errorConflictingDefElem(m, l2, l0)
	mBase = m.M
	v295 = m.ExcPending
	if v295 != 0 {
		goto L67
	} else {
		goto L119
	}
L3:
	;
	return v291
L4:
	;
	v291 = int32(1)
	goto L3
L5:
	;
	if v39-v38 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L6:
	;
	goto L5
L7:
	;
	if v18 != v19 {
		v38 = v18
		v39 = v19
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v23 = v14
	v24 = v15
	goto L9
L9:
	;
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+1)))
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+1)))
	if v28 == int32(0) {
		v38 = v27
		v39 = v28
		goto L6
	} else {
		goto L11
	}
L10:
	;
	v38 = v27
	v39 = v28
	goto L6
L11:
	;
	v31 = int32(1)
	if v27 == v28 {
		v23 = v23 + v31
		v24 = v24 + v31
		goto L9
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v45 = int32(_a_F_compute_common_attribute_1)
	v48 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[1])))
	v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v49 == int32(0) {
		v68 = v48
		v69 = v49
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v43 != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = l2
	goto L4
L18:
	;
	if v69-v68 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L19:
	;
	goto L18
L20:
	;
	if v48 != v49 {
		v68 = v48
		v69 = v49
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v53 = v14
	v54 = v45
	goto L22
L22:
	;
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+1)))
	v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53)+1)))
	if v58 == int32(0) {
		v68 = v57
		v69 = v58
		goto L19
	} else {
		goto L24
	}
L23:
	;
	v68 = v57
	v69 = v58
	goto L19
L24:
	;
	v61 = int32(1)
	if v57 == v58 {
		v53 = v53 + v61
		v54 = v54 + v61
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v75 = int32(_a_F_compute_common_attribute_2)
	v78 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[2])))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v79 == int32(0) {
		v98 = v78
		v99 = v79
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if v73 != 0 {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = l2
	goto L4
L31:
	;
	if v99-v98 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L32:
	;
	goto L31
L33:
	;
	if v78 != v79 {
		v98 = v78
		v99 = v79
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v83 = v14
	v84 = v75
	goto L35
L35:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v84)+1)))
	v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+1)))
	if v88 == int32(0) {
		v98 = v87
		v99 = v88
		goto L32
	} else {
		goto L37
	}
L36:
	;
	v98 = v87
	v99 = v88
	goto L32
L37:
	;
	v91 = int32(1)
	if v87 == v88 {
		v83 = v83 + v91
		v84 = v84 + v91
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	if v103 != 0 {
		goto L2
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v105 = int32(_a_F_compute_common_attribute_3)
	v108 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[3])))
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v109 == int32(0) {
		v128 = v108
		v129 = v109
		goto L44
	} else {
		goto L45
	}
L42:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = l2
	goto L4
L43:
	;
	if v129-v128 == int32(0) {
		goto L51
	} else {
		goto L52
	}
L44:
	;
	goto L43
L45:
	;
	if v108 != v109 {
		v128 = v108
		v129 = v109
		goto L44
	} else {
		goto L46
	}
L46:
	;
	v113 = v14
	v114 = v105
	goto L47
L47:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+1)))
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113)+1)))
	if v118 == int32(0) {
		v128 = v117
		v129 = v118
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v128 = v117
	v129 = v118
	goto L44
L49:
	;
	v121 = int32(1)
	if v117 == v118 {
		v113 = v113 + v121
		v114 = v114 + v121
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L54
	}
L52:
	;
	goto L53
L53:
	;
	v135 = int32(_a_F_compute_common_attribute_4)
	v138 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[4])))
	v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v139 == int32(0) {
		v158 = v138
		v159 = v139
		goto L57
	} else {
		goto L58
	}
L54:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v133 != 0 {
		goto L2
	} else {
		goto L55
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l6))) = l2
	goto L4
L56:
	;
	if v159-v158 == int32(0) {
		goto L64
	} else {
		goto L65
	}
L57:
	;
	goto L56
L58:
	;
	if v138 != v139 {
		v158 = v138
		v159 = v139
		goto L57
	} else {
		goto L59
	}
L59:
	;
	v143 = v14
	v144 = v135
	goto L60
L60:
	;
	v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144)+1)))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+1)))
	if v148 == int32(0) {
		v158 = v147
		v159 = v148
		goto L57
	} else {
		goto L62
	}
L61:
	;
	v158 = v147
	v159 = v148
	goto L57
L62:
	;
	v151 = int32(1)
	if v147 == v148 {
		v143 = v143 + v151
		v144 = v144 + v151
		goto L60
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(l7)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v165 = F_lappend(m, v163, v164)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v170 = int32(_a_F_compute_common_attribute_5)
	v173 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[5])))
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v174 == int32(0) {
		v193 = v173
		v194 = v174
		goto L70
	} else {
		goto L71
	}
L67:
	;
	return int32(0)
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v165
	goto L4
L69:
	;
	if v194-v193 == int32(0) {
		goto L77
	} else {
		goto L78
	}
L70:
	;
	goto L69
L71:
	;
	if v173 != v174 {
		v193 = v173
		v194 = v174
		goto L70
	} else {
		goto L72
	}
L72:
	;
	v178 = v14
	v179 = v170
	goto L73
L73:
	;
	v182 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)))
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	if v183 == int32(0) {
		v193 = v182
		v194 = v183
		goto L70
	} else {
		goto L75
	}
L74:
	;
	v193 = v182
	v194 = v183
	goto L70
L75:
	;
	v186 = int32(1)
	if v182 == v183 {
		v178 = v178 + v186
		v179 = v179 + v186
		goto L73
	} else {
		goto L76
	}
L76:
	;
	goto L74
L77:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	v200 = int32(_a_F_compute_common_attribute_6)
	v203 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[6])))
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v204 == int32(0) {
		v223 = v203
		v224 = v204
		goto L83
	} else {
		goto L84
	}
L80:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(l8)))
	if v198 != 0 {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l8))) = l2
	goto L4
L82:
	;
	if v224-v223 == int32(0) {
		goto L90
	} else {
		goto L91
	}
L83:
	;
	goto L82
L84:
	;
	if v203 != v204 {
		v223 = v203
		v224 = v204
		goto L83
	} else {
		goto L85
	}
L85:
	;
	v208 = v14
	v209 = v200
	goto L86
L86:
	;
	v212 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
	v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
	if v213 == int32(0) {
		v223 = v212
		v224 = v213
		goto L83
	} else {
		goto L88
	}
L87:
	;
	v223 = v212
	v224 = v213
	goto L83
L88:
	;
	v216 = int32(1)
	if v212 == v213 {
		v208 = v208 + v216
		v209 = v209 + v216
		goto L86
	} else {
		goto L89
	}
L89:
	;
	goto L87
L90:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L93
	}
L91:
	;
	goto L92
L92:
	;
	v230 = int32(_a_F_compute_common_attribute_7)
	v233 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[7])))
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v234 == int32(0) {
		v253 = v233
		v254 = v234
		goto L96
	} else {
		goto L97
	}
L93:
	;
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l9)))
	if v228 != 0 {
		goto L2
	} else {
		goto L94
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l9))) = l2
	goto L4
L95:
	;
	if v254-v253 == int32(0) {
		goto L103
	} else {
		goto L104
	}
L96:
	;
	goto L95
L97:
	;
	if v233 != v234 {
		v253 = v233
		v254 = v234
		goto L96
	} else {
		goto L98
	}
L98:
	;
	v238 = v14
	v239 = v230
	goto L99
L99:
	;
	v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239)+1)))
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v238)+1)))
	if v243 == int32(0) {
		v253 = v242
		v254 = v243
		goto L96
	} else {
		goto L101
	}
L100:
	;
	v253 = v242
	v254 = v243
	goto L96
L101:
	;
	v246 = int32(1)
	if v242 == v243 {
		v238 = v238 + v246
		v239 = v239 + v246
		goto L99
	} else {
		goto L102
	}
L102:
	;
	goto L100
L103:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L106
	}
L104:
	;
	goto L105
L105:
	;
	v260 = int32(0)
	v261 = int32(_a_F_compute_common_attribute_8)
	v264 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_compute_common_attribute[8])))
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v265 == v260 {
		v284 = v264
		v285 = v265
		goto L109
	} else {
		goto L110
	}
L106:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l10)))
	if v258 != 0 {
		goto L2
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l10))) = l2
	goto L4
L108:
	;
	if v285-v284 != 0 {
		v291 = v260
		goto L3
	} else {
		goto L116
	}
L109:
	;
	goto L108
L110:
	;
	if v264 != v265 {
		v284 = v264
		v285 = v265
		goto L109
	} else {
		goto L111
	}
L111:
	;
	v269 = v14
	v270 = v261
	goto L112
L112:
	;
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v270)+1)))
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v269)+1)))
	if v274 == int32(0) {
		v284 = v273
		v285 = v274
		goto L109
	} else {
		goto L114
	}
L113:
	;
	v284 = v273
	v285 = v274
	goto L109
L114:
	;
	v277 = int32(1)
	if v273 == v274 {
		v269 = v269 + v277
		v270 = v270 + v277
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	if l1 != 0 {
		goto L1
	} else {
		goto L117
	}
L117:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(l11)))
	if v287 != 0 {
		goto L2
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l11))) = l2
	goto L4
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	F_errcode(m, int32(50724996))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L67
	} else {
		goto L121
	}
L121:
	;
	F_errmsg(m, int32(_a_F_compute_common_attribute_9), int32(0))
	mBase = m.M
	v307 = m.ExcPending
	if v307 != 0 {
		goto L67
	} else {
		goto L122
	}
L122:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
	F_parser_errposition(m, l0, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L67
	} else {
		goto L123
	}
L123:
	;
	F_errfinish(m, int32(_a_F_compute_common_attribute_10), int32(612), int32(_a_F_compute_common_attribute_11))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L67
	} else {
		goto L124
	}
L124:
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
	var v27 int32
	_ = v27
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
	v27 = v20
	goto L11
L11:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v30+v27<<(uint(int32(2))%32))))
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
	v42 = v27 + int32(1)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v42 < v43 {
		v27 = v42
		goto L11
	} else {
		goto L17
	}
L17:
	;
	goto L12
}
