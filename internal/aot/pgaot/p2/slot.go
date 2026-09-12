package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_ExecBuildSlotValueDescription(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v185 int32
	_ = v185
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v220 int32
	_ = v220
	v5 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v21 = F_check_enable_rls(m, l0, v5, int32(1))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(48)
	return v220
L2:
	;
	return int32(0)
L3:
	;
	if v21 == int32(2) {
		v220 = v5
		goto L1
	} else {
		goto L4
	}
L4:
	;
	F_initStringInfo(m, v17+int32(32))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	F_appendStringInfoChar(m, v17+int32(32), int32(40))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v37 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v39 = F_pg_class_aclcheck(m, l0, v37, int64(2))
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L2
	} else {
		goto L7
	}
L7:
	;
	if v39 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	F_initStringInfo(m, v17+int32(16))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L2
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v52 < v51 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	F_appendStringInfoChar(m, v17+int32(16), int32(40))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	goto L10
L13:
	;
	F_slot_getsomeattrs_int(m, l1, v51)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L2
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v56 = int32(0)
	v57 = base.B2i32(v39 == v56)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v56 < v58 {
		goto L17
	} else {
		goto L18
	}
L16:
	;
	goto L15
L17:
	;
	v69 = v58
	v70 = int32(0)
	v71 = v5
	v73 = v57
	v77 = v5
	goto L20
L18:
	;
	v185 = v57
	goto L19
L19:
	;
	if v185 == int32(0) {
		v220 = v5
		goto L1
	} else {
		goto L56
	}
L20:
	;
	v83 = l2 + int32(20) + v69<<(uint(int32(4))%32) + v70*int32(100)
	v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+91)))
	if v84 != 0 {
		v168 = v71
		v169 = v73
		v171 = v77
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v185 = v169
	goto L19
L22:
	;
	v173 = v70 + int32(1)
	v174 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v173 < v174 {
		v69 = v174
		v70 = v173
		v71 = v168
		v73 = v169
		v77 = v171
		goto L20
	} else {
		goto L55
	}
L23:
	;
	if v39 != 0 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v85 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+74)))
	v87 = *(*int32)(unsafe.Add(mBase, _consts[279]))
	v89 = F_pg_attribute_aclcheck(m, l0, v85, v87, int64(2))
	mBase = m.M
	v90 = m.ExcPending
	if v90 != 0 {
		goto L2
	} else {
		goto L27
	}
L25:
	;
	v113 = v73
	v115 = v77
	goto L26
L26:
	;
	v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83)+90)))
	if v117 == int32(118) {
		v139 = int32(310511)
		goto L38
	} else {
		goto L39
	}
L27:
	;
	v91 = int32(*(*int16)(unsafe.Add(mBase, uint32(v83)+74)))
	v94 = F_bms_is_member(m, v91+int32(7), l3)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L28
	}
L28:
	;
	if v89 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v97 = v94
	goto L31
L30:
	;
	v97 = int32(1)
	goto L31
L31:
	;
	if v97 == int32(0) {
		v168 = v71
		v169 = v73
		v171 = v77
		goto L22
	} else {
		goto L32
	}
L32:
	;
	if v77 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	F_appendStringInfoString(m, v17+int32(16), int32(747599))
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L2
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	F_appendStringInfoString(m, v17+int32(16), v83+int32(4))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L2
	} else {
		goto L37
	}
L36:
	;
	goto L35
L37:
	;
	v111 = int32(1)
	v113 = v111
	v115 = v111
	goto L26
L38:
	;
	if v71&int32(1) != 0 {
		goto L43
	} else {
		goto L44
	}
L39:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121+v70))))
	if v123 != 0 {
		v139 = int32(304161)
		goto L38
	} else {
		goto L40
	}
L40:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v83)+68))
	F_getTypeOutputInfo(m, v124, v17+int32(12), v17+int32(11))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v132+v70<<(uint(int32(2))%32))))
	v137 = F_OidOutputFunctionCall(m, v131, v136)
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L2
	} else {
		goto L42
	}
L42:
	;
	v139 = v137
	goto L38
L43:
	;
	F_appendStringInfoString(m, v17+int32(32), int32(747599))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L2
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v147 = F_strlen(m, v139)
	mBase = m.M
	if v147 <= int32(64) {
		goto L48
	} else {
		goto L49
	}
L46:
	;
	goto L45
L47:
	;
	v168 = int32(1)
	v169 = v113
	v171 = v115
	goto L22
L48:
	;
	F_appendBinaryStringInfo(m, v17+int32(32), v139, v147)
	mBase = m.M
	v153 = m.ExcPending
	if v153 != 0 {
		goto L2
	} else {
		goto L51
	}
L49:
	;
	goto L50
L50:
	;
	v157 = F_pg_mbcliplen(m, v139, v147, int32(64))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L2
	} else {
		goto L52
	}
L51:
	;
	goto L47
L52:
	;
	F_appendBinaryStringInfo(m, v17+int32(32), v139, v157)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L2
	} else {
		goto L53
	}
L53:
	;
	F_appendStringInfoString(m, v17+int32(32), int32(661888))
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L2
	} else {
		goto L54
	}
L54:
	;
	goto L47
L55:
	;
	goto L21
L56:
	;
	F_appendStringInfoChar(m, v17+int32(32), int32(41))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L2
	} else {
		goto L57
	}
L57:
	;
	if v39 != 0 {
		goto L58
	} else {
		goto L59
	}
L58:
	;
	F_appendStringInfoString(m, v17+int32(16), int32(746868))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L2
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v220 = v209
	goto L1
L61:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v17)+32))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v17)+36))
	F_appendBinaryStringInfo(m, v17+int32(16), v204, v205)
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L2
	} else {
		goto L62
	}
L62:
	;
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v17)+16))
	v220 = v208
	goto L1
}
func F_ExecComputeSlotInfo(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
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
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v169 int32
	_ = v169
	v3 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v3)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v13 != 0 {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
		*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(base.B2i32(v14 != int32(0)))
		v144 = v13
		v145 = v14
		v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
		if v146 != int32(1) {
			v153 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
			*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
			v169 = int32(1)
		} else {
			if v144 == int32(0) {
				v153 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
				v169 = int32(1)
			} else {
				if v145 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v145
					v158 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v158)
					*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v144
					if v145 != int32(1619316) {
						v169 = int32(1)
					} else {
						v169 = int32(0)
					}
				} else {
					v153 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
					*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
					v169 = int32(1)
				}
			}
		}
	} else {
		if v10 == int32(0) {
			v153 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
			*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
			v169 = int32(1)
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			switch v20 - int32(2) {
			case 0:
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v10)+40))
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+102)))
				if v24 != int32(1) {
					if v23 == int32(0) {
						v153 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v169 = int32(1)
					} else {
						v40 = v8 + int32(15)
						v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+103)))
						if v42 == int32(1) {
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
							if v45 != 0 {
								if v40 == int32(0) {
									v75 = v45
									v78 = v75
								} else {
									v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+99)))
									*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v48)
									v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
									v78 = v50
								}
							} else {
								if v40 == int32(0) {
								} else {
									v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+99)))
									v66 = v53
									*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v66)
								}
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
								if v69 == int32(0) {
									v78 = int32(1619316)
								} else {
									v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
									v75 = v73
									v78 = v75
								}
							}
						} else {
							if v40 == int32(0) {
							} else {
								v56 = int32(0)
								v57 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
								if v57 == v56 {
									v66 = v56
								} else {
									v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)))
									v66 = int32(base.Ui32(v60)>>(uint(int32(4))%32)) & int32(1)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v66)
							}
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
							if v69 == int32(0) {
								v78 = int32(1619316)
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
								v75 = v73
								v78 = v75
							}
						}
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
						v144 = v79
						v145 = v78
						v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						if v146 != int32(1) {
							v153 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
							v169 = int32(1)
						} else {
							if v144 == int32(0) {
								v153 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v169 = int32(1)
							} else {
								if v145 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v145
									v158 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v158)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v144
									if v145 != int32(1619316) {
										v169 = int32(1)
									} else {
										v169 = int32(0)
									}
								} else {
									v153 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v169 = int32(1)
								}
							}
						}
					}
				} else {
					v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+98)))
					if v27 != int32(1) {
						v153 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v169 = int32(1)
					} else {
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+88))
						if v30 == int32(0) {
							if v23 == int32(0) {
								v153 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v169 = int32(1)
							} else {
								v40 = v8 + int32(15)
								v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+103)))
								if v42 == int32(1) {
									v45 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
									if v45 != 0 {
										if v40 == int32(0) {
											v75 = v45
											v78 = v75
										} else {
											v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+99)))
											*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v48)
											v50 = *(*int32)(unsafe.Add(mBase, uint32(v23)+92))
											v78 = v50
										}
									} else {
										if v40 == int32(0) {
										} else {
											v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+99)))
											v66 = v53
											*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v66)
										}
										v69 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
										if v69 == int32(0) {
											v78 = int32(1619316)
										} else {
											v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
											v75 = v73
											v78 = v75
										}
									}
								} else {
									if v40 == int32(0) {
									} else {
										v56 = int32(0)
										v57 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
										if v57 == v56 {
											v66 = v56
										} else {
											v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v57)+4)))
											v66 = int32(base.Ui32(v60)>>(uint(int32(4))%32)) & int32(1)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(v40))) = uint8(v66)
									}
									v69 = *(*int32)(unsafe.Add(mBase, uint32(v23)+60))
									if v69 == int32(0) {
										v78 = int32(1619316)
									} else {
										v73 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
										v75 = v73
										v78 = v75
									}
								}
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
								v144 = v79
								v145 = v78
								v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
								if v146 != int32(1) {
									v153 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v169 = int32(1)
								} else {
									if v144 == int32(0) {
										v153 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
										*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
										v169 = int32(1)
									} else {
										if v145 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v145
											v158 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v158)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v144
											if v145 != int32(1619316) {
												v169 = int32(1)
											} else {
												v169 = int32(0)
											}
										} else {
											v153 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
											*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
											v169 = int32(1)
										}
									}
								}
							}
						} else {
							v33 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v33)
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v23)+56))
							v144 = v35
							v145 = v30
							v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							if v146 != int32(1) {
								v153 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v169 = int32(1)
							} else {
								if v144 == int32(0) {
									v153 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v169 = int32(1)
								} else {
									if v145 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v145
										v158 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v158)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v144
										if v145 != int32(1619316) {
											v169 = int32(1)
										} else {
											v169 = int32(0)
										}
									} else {
										v153 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
										*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
										v169 = int32(1)
									}
								}
							}
						}
					}
				}
			case 1:
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v10)+36))
				v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+101)))
				if v81 != int32(1) {
					if v80 == int32(0) {
						v153 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v169 = int32(1)
					} else {
						v97 = v8 + int32(15)
						v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+103)))
						if v99 == int32(1) {
							v102 = *(*int32)(unsafe.Add(mBase, uint32(v80)+92))
							if v102 != 0 {
								if v97 == int32(0) {
									v132 = v102
									v135 = v132
								} else {
									v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+99)))
									*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v105)
									v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+92))
									v135 = v107
								}
							} else {
								if v97 == int32(0) {
								} else {
									v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+99)))
									v123 = v110
									*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v123)
								}
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v80)+60))
								if v126 == int32(0) {
									v135 = int32(1619316)
								} else {
									v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
									v132 = v130
									v135 = v132
								}
							}
						} else {
							if v97 == int32(0) {
							} else {
								v113 = int32(0)
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+60))
								if v114 == v113 {
									v123 = v113
								} else {
									v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
									v123 = int32(base.Ui32(v117)>>(uint(int32(4))%32)) & int32(1)
								}
								*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v123)
							}
							v126 = *(*int32)(unsafe.Add(mBase, uint32(v80)+60))
							if v126 == int32(0) {
								v135 = int32(1619316)
							} else {
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
								v132 = v130
								v135 = v132
							}
						}
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v80)+56))
						v144 = v136
						v145 = v135
						v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
						if v146 != int32(1) {
							v153 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
							v169 = int32(1)
						} else {
							if v144 == int32(0) {
								v153 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v169 = int32(1)
							} else {
								if v145 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v145
									v158 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v158)
									*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v144
									if v145 != int32(1619316) {
										v169 = int32(1)
									} else {
										v169 = int32(0)
									}
								} else {
									v153 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v169 = int32(1)
								}
							}
						}
					}
				} else {
					v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+97)))
					if v84 != int32(1) {
						v153 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v169 = int32(1)
					} else {
						v87 = *(*int32)(unsafe.Add(mBase, uint32(v10)+84))
						if v87 == int32(0) {
							if v80 == int32(0) {
								v153 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v169 = int32(1)
							} else {
								v97 = v8 + int32(15)
								v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+103)))
								if v99 == int32(1) {
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v80)+92))
									if v102 != 0 {
										if v97 == int32(0) {
											v132 = v102
											v135 = v132
										} else {
											v105 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+99)))
											*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v105)
											v107 = *(*int32)(unsafe.Add(mBase, uint32(v80)+92))
											v135 = v107
										}
									} else {
										if v97 == int32(0) {
										} else {
											v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+99)))
											v123 = v110
											*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v123)
										}
										v126 = *(*int32)(unsafe.Add(mBase, uint32(v80)+60))
										if v126 == int32(0) {
											v135 = int32(1619316)
										} else {
											v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
											v132 = v130
											v135 = v132
										}
									}
								} else {
									if v97 == int32(0) {
									} else {
										v113 = int32(0)
										v114 = *(*int32)(unsafe.Add(mBase, uint32(v80)+60))
										if v114 == v113 {
											v123 = v113
										} else {
											v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v114)+4)))
											v123 = int32(base.Ui32(v117)>>(uint(int32(4))%32)) & int32(1)
										}
										*(*uint8)(unsafe.Add(mBase, uint32(v97))) = uint8(v123)
									}
									v126 = *(*int32)(unsafe.Add(mBase, uint32(v80)+60))
									if v126 == int32(0) {
										v135 = int32(1619316)
									} else {
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v126)+8))
										v132 = v130
										v135 = v132
									}
								}
								v136 = *(*int32)(unsafe.Add(mBase, uint32(v80)+56))
								v144 = v136
								v145 = v135
								v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
								if v146 != int32(1) {
									v153 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v169 = int32(1)
								} else {
									if v144 == int32(0) {
										v153 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
										*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
										v169 = int32(1)
									} else {
										if v145 != 0 {
											*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v145
											v158 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v158)
											*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v144
											if v145 != int32(1619316) {
												v169 = int32(1)
											} else {
												v169 = int32(0)
											}
										} else {
											v153 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
											*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
											v169 = int32(1)
										}
									}
								}
							}
						} else {
							v90 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v90)
							v92 = *(*int32)(unsafe.Add(mBase, uint32(v80)+56))
							v144 = v92
							v145 = v87
							v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
							if v146 != int32(1) {
								v153 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
								*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
								v169 = int32(1)
							} else {
								if v144 == int32(0) {
									v153 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
									*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
									v169 = int32(1)
								} else {
									if v145 != 0 {
										*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v145
										v158 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v158)
										*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v144
										if v145 != int32(1619316) {
											v169 = int32(1)
										} else {
											v169 = int32(0)
										}
									} else {
										v153 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
										*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
										v169 = int32(1)
									}
								}
							}
						}
					}
				}
			case 2, 3, 4:
				v137 = *(*int32)(unsafe.Add(mBase, uint32(v10)+80))
				v138 = *(*int32)(unsafe.Add(mBase, uint32(v10)+76))
				v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+100)))
				if v139 != int32(1) {
					v144 = v138
					v145 = v137
				} else {
					v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+96)))
					*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)) = uint8(v142)
					v144 = v138
					v145 = v137
				}
				v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+15)))
				if v146 != int32(1) {
					v153 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
					*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
					v169 = int32(1)
				} else {
					if v144 == int32(0) {
						v153 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
						*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
						v169 = int32(1)
					} else {
						if v145 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v145
							v158 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v158)
							*(*int32)(unsafe.Add(mBase, uint32(l1)+24)) = v144
							if v145 != int32(1619316) {
								v169 = int32(1)
							} else {
								v169 = int32(0)
							}
						} else {
							v153 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
							*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
							v169 = int32(1)
						}
					}
				}
			default:
				v153 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+20)) = uint8(v153)
				*(*int64)(unsafe.Add(mBase, uint32(l1)+24)) = int64(0)
				v169 = int32(1)
			}
		}
	}
	m.G0 = v8 + int32(16)
	return v169
}
func F_ExecCreateScanSlotFromOuterPlan(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+56))
	F_ExecInitScanTupleSlot(m, l0, l1, v5, l2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		return
	}
}
func F_ExecSetSlotDescriptor(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(v4)+12))
	m.T0[v5].(func(*base.Module, int32))(m, l0)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v8 == int32(0) {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			if v16 != 0 {
				F_pfree(m, v16)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v19 != 0 {
						F_pfree(m, v19)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if int32(0) <= v23 {
								F_IncrTupleDescRefCount(m, l1)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						if int32(0) <= v23 {
							F_IncrTupleDescRefCount(m, l1)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v37 = F_MemoryContextAlloc(m, v35, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
									return
								}
							}
						}
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v19 != 0 {
					F_pfree(m, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						if int32(0) <= v23 {
							F_IncrTupleDescRefCount(m, l1)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v37 = F_MemoryContextAlloc(m, v35, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
									return
								}
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
					v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					if int32(0) <= v23 {
						F_IncrTupleDescRefCount(m, l1)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v37 = F_MemoryContextAlloc(m, v35, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
									return
								}
							}
						}
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
							v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v37 = F_MemoryContextAlloc(m, v35, v36)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
								return
							}
						}
					}
				}
			}
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			if v11 < int32(0) {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				if v16 != 0 {
					F_pfree(m, v16)
					mBase = m.M
					v18 = m.ExcPending
					if v18 != 0 {
						return
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v19 != 0 {
							F_pfree(m, v19)
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if int32(0) <= v23 {
									F_IncrTupleDescRefCount(m, l1)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
											v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v37 = F_MemoryContextAlloc(m, v35, v36)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
												return
											}
										}
									}
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if int32(0) <= v23 {
								F_IncrTupleDescRefCount(m, l1)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						}
					}
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v19 != 0 {
						F_pfree(m, v19)
						mBase = m.M
						v21 = m.ExcPending
						if v21 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if int32(0) <= v23 {
								F_IncrTupleDescRefCount(m, l1)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
						v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
						if int32(0) <= v23 {
							F_IncrTupleDescRefCount(m, l1)
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
							mBase = m.M
							v33 = m.ExcPending
							if v33 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
								v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v37 = F_MemoryContextAlloc(m, v35, v36)
								mBase = m.M
								v38 = m.ExcPending
								if v38 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
									return
								}
							}
						}
					}
				}
			} else {
				F_DecrTupleDescRefCount(m, v8)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return
				} else {
					v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
					if v16 != 0 {
						F_pfree(m, v16)
						mBase = m.M
						v18 = m.ExcPending
						if v18 != 0 {
							return
						} else {
							v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
							if v19 != 0 {
								F_pfree(m, v19)
								mBase = m.M
								v21 = m.ExcPending
								if v21 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
									v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
									if int32(0) <= v23 {
										F_IncrTupleDescRefCount(m, l1)
										mBase = m.M
										v27 = m.ExcPending
										if v27 != 0 {
											return
										} else {
											v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
											mBase = m.M
											v33 = m.ExcPending
											if v33 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
												v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
												v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
												v37 = F_MemoryContextAlloc(m, v35, v36)
												mBase = m.M
												v38 = m.ExcPending
												if v38 != 0 {
													return
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
													return
												}
											}
										}
									} else {
										v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
											v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v37 = F_MemoryContextAlloc(m, v35, v36)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
												return
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if int32(0) <= v23 {
									F_IncrTupleDescRefCount(m, l1)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
											v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v37 = F_MemoryContextAlloc(m, v35, v36)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
												return
											}
										}
									}
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							}
						}
					} else {
						v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
						if v19 != 0 {
							F_pfree(m, v19)
							mBase = m.M
							v21 = m.ExcPending
							if v21 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
								v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
								if int32(0) <= v23 {
									F_IncrTupleDescRefCount(m, l1)
									mBase = m.M
									v27 = m.ExcPending
									if v27 != 0 {
										return
									} else {
										v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
										mBase = m.M
										v33 = m.ExcPending
										if v33 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
											v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
											v37 = F_MemoryContextAlloc(m, v35, v36)
											mBase = m.M
											v38 = m.ExcPending
											if v38 != 0 {
												return
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
												return
											}
										}
									}
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = l1
							v23 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if int32(0) <= v23 {
								F_IncrTupleDescRefCount(m, l1)
								mBase = m.M
								v27 = m.ExcPending
								if v27 != 0 {
									return
								} else {
									v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
										v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
										v37 = F_MemoryContextAlloc(m, v35, v36)
										mBase = m.M
										v38 = m.ExcPending
										if v38 != 0 {
											return
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
											return
										}
									}
								}
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								v29 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
								v32 = F_MemoryContextAlloc(m, v28, v29<<(uint(int32(2))%32))
								mBase = m.M
								v33 = m.ExcPending
								if v33 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v32
									v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
									v37 = F_MemoryContextAlloc(m, v35, v36)
									mBase = m.M
									v38 = m.ExcPending
									if v38 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v37
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_GetSlotInvalidationCauseName(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	v3 = int32(373553)
	if base.Ui32(int32(8)) < base.Ui32(l0) {
		v18 = v3
	} else {
		if int32(base.Ui32(int32(279))>>(uint(l0)%32))&int32(1) == int32(0) {
			v18 = v3
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l0<<(uint(int32(2))%32))+uint32(_consts[682])))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
			v18 = v17
		}
	}
	return v18
}
func F_ReportSlotInvalidation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int64, l5 int64, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	v5 = l4
	v10 = m.G0
	v12 = v10 - int32(176)
	m.G0 = v12
	F_initStringInfo(m, v12+int32(160))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		F_initStringInfo(m, v12+int32(144))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return
		} else {
			switch l0 - int32(1) {
			case 0:
				*(*uint32)(unsafe.Add(mBase, uint32(v12)+84)) = uint32(v5)
				v26 = int64(base.Ui64(v5) >> (uint(int64(32)) % 64))
				*(*uint32)(unsafe.Add(mBase, uint32(v12)+80)) = uint32(v26)
				v28 = l5 - v5
				*(*int64)(unsafe.Add(mBase, uint32(v12)+88)) = v28
				if v28 == int64(1) {
					v36 = int32(630645)
				} else {
					v36 = int32(598204)
				}
				F_appendStringInfo(m, v12+int32(160), v36, v12+int32(80))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = int32(341478)
					F_appendStringInfo(m, v12+int32(144), int32(668123), v12-int32(-64))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return
					} else {
						v88 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							if v88 != 0 {
								if l1 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
									F_errmsg(m, int32(699643), v12+int32(32))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
										F_errdetail_internal(m, int32(206576), v12+int32(16))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
											if v110 != 0 {
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
												F_errhint(m, int32(206576), v12)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													F_errfinish(m, int32(494185), int32(1705), int32(265463))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return
													} else {
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v121)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v124)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											} else {
												F_errfinish(m, int32(494185), int32(1705), int32(265463))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v124)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
									F_errmsg(m, int32(699599), v12+int32(48))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
										F_errdetail_internal(m, int32(206576), v12+int32(16))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
											if v110 != 0 {
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
												F_errhint(m, int32(206576), v12)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													F_errfinish(m, int32(494185), int32(1705), int32(265463))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return
													} else {
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v121)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v124)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											} else {
												F_errfinish(m, int32(494185), int32(1705), int32(265463))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v124)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
								F_pfree(m, v121)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
									F_pfree(m, v124)
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return
									} else {
										m.G0 = v12 + int32(176)
										return
									}
								}
							}
						}
					}
				}
			case 1:
				*(*int32)(unsafe.Add(mBase, uint32(v12)+96)) = l6
				F_appendStringInfo(m, v12+int32(160), int32(578852), v12+int32(96))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return
				} else {
					v88 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						if v88 != 0 {
							if l1 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
								F_errmsg(m, int32(699643), v12+int32(32))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
									F_errdetail_internal(m, int32(206576), v12+int32(16))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
										if v110 != 0 {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
											F_errhint(m, int32(206576), v12)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												F_errfinish(m, int32(494185), int32(1705), int32(265463))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v124)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										} else {
											F_errfinish(m, int32(494185), int32(1705), int32(265463))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
												F_pfree(m, v121)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													F_pfree(m, v124)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														m.G0 = v12 + int32(176)
														return
													}
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
								F_errmsg(m, int32(699599), v12+int32(48))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
									F_errdetail_internal(m, int32(206576), v12+int32(16))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
										if v110 != 0 {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
											F_errhint(m, int32(206576), v12)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												F_errfinish(m, int32(494185), int32(1705), int32(265463))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v124)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										} else {
											F_errfinish(m, int32(494185), int32(1705), int32(265463))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
												F_pfree(m, v121)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													F_pfree(m, v124)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														m.G0 = v12 + int32(176)
														return
													}
												}
											}
										}
									}
								}
							}
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
							F_pfree(m, v121)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return
							} else {
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
								F_pfree(m, v124)
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return
								} else {
									m.G0 = v12 + int32(176)
									return
								}
							}
						}
					}
				}
			default:
				v88 = F_errstart(m, int32(15), int32(0))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
					return
				} else {
					if v88 != 0 {
						if l1 != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
							*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
							F_errmsg(m, int32(699643), v12+int32(32))
							mBase = m.M
							v96 = m.ExcPending
							if v96 != 0 {
								return
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
								F_errdetail_internal(m, int32(206576), v12+int32(16))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
									if v110 != 0 {
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
										F_errhint(m, int32(206576), v12)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											F_errfinish(m, int32(494185), int32(1705), int32(265463))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
												F_pfree(m, v121)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													F_pfree(m, v124)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														m.G0 = v12 + int32(176)
														return
													}
												}
											}
										}
									} else {
										F_errfinish(m, int32(494185), int32(1705), int32(265463))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
											F_pfree(m, v121)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												F_pfree(m, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													m.G0 = v12 + int32(176)
													return
												}
											}
										}
									}
								}
							}
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
							F_errmsg(m, int32(699599), v12+int32(48))
							mBase = m.M
							v102 = m.ExcPending
							if v102 != 0 {
								return
							} else {
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
								*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
								F_errdetail_internal(m, int32(206576), v12+int32(16))
								mBase = m.M
								v109 = m.ExcPending
								if v109 != 0 {
									return
								} else {
									v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
									if v110 != 0 {
										v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
										*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
										F_errhint(m, int32(206576), v12)
										mBase = m.M
										v115 = m.ExcPending
										if v115 != 0 {
											return
										} else {
											F_errfinish(m, int32(494185), int32(1705), int32(265463))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
												F_pfree(m, v121)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													F_pfree(m, v124)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														m.G0 = v12 + int32(176)
														return
													}
												}
											}
										}
									} else {
										F_errfinish(m, int32(494185), int32(1705), int32(265463))
										mBase = m.M
										v120 = m.ExcPending
										if v120 != 0 {
											return
										} else {
											v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
											F_pfree(m, v121)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return
											} else {
												v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												F_pfree(m, v124)
												mBase = m.M
												v126 = m.ExcPending
												if v126 != 0 {
													return
												} else {
													m.G0 = v12 + int32(176)
													return
												}
											}
										}
									}
								}
							}
						}
					} else {
						v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
						F_pfree(m, v121)
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return
						} else {
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
							F_pfree(m, v124)
							mBase = m.M
							v126 = m.ExcPending
							if v126 != 0 {
								return
							} else {
								m.G0 = v12 + int32(176)
								return
							}
						}
					}
				}
			case 3:
				F_appendStringInfoString(m, v12+int32(160), int32(609308))
				mBase = m.M
				v62 = m.ExcPending
				if v62 != 0 {
					return
				} else {
					v88 = F_errstart(m, int32(15), int32(0))
					mBase = m.M
					v89 = m.ExcPending
					if v89 != 0 {
						return
					} else {
						if v88 != 0 {
							if l1 != 0 {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
								*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
								F_errmsg(m, int32(699643), v12+int32(32))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
									F_errdetail_internal(m, int32(206576), v12+int32(16))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
										if v110 != 0 {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
											F_errhint(m, int32(206576), v12)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												F_errfinish(m, int32(494185), int32(1705), int32(265463))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v124)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										} else {
											F_errfinish(m, int32(494185), int32(1705), int32(265463))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
												F_pfree(m, v121)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													F_pfree(m, v124)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														m.G0 = v12 + int32(176)
														return
													}
												}
											}
										}
									}
								}
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
								F_errmsg(m, int32(699599), v12+int32(48))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
									*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
									F_errdetail_internal(m, int32(206576), v12+int32(16))
									mBase = m.M
									v109 = m.ExcPending
									if v109 != 0 {
										return
									} else {
										v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
										if v110 != 0 {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
											*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
											F_errhint(m, int32(206576), v12)
											mBase = m.M
											v115 = m.ExcPending
											if v115 != 0 {
												return
											} else {
												F_errfinish(m, int32(494185), int32(1705), int32(265463))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v124)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										} else {
											F_errfinish(m, int32(494185), int32(1705), int32(265463))
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return
											} else {
												v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
												F_pfree(m, v121)
												mBase = m.M
												v123 = m.ExcPending
												if v123 != 0 {
													return
												} else {
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
													F_pfree(m, v124)
													mBase = m.M
													v126 = m.ExcPending
													if v126 != 0 {
														return
													} else {
														m.G0 = v12 + int32(176)
														return
													}
												}
											}
										}
									}
								}
							}
						} else {
							v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
							F_pfree(m, v121)
							mBase = m.M
							v123 = m.ExcPending
							if v123 != 0 {
								return
							} else {
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
								F_pfree(m, v124)
								mBase = m.M
								v126 = m.ExcPending
								if v126 != 0 {
									return
								} else {
									m.G0 = v12 + int32(176)
									return
								}
							}
						}
					}
				}
			case 7:
				*(*int32)(unsafe.Add(mBase, uint32(v12)+128)) = l7
				*(*int32)(unsafe.Add(mBase, uint32(v12)+132)) = int32(65745)
				v67 = *(*int32)(unsafe.Add(mBase, _consts[683]))
				*(*int32)(unsafe.Add(mBase, uint32(v12)+136)) = v67
				F_appendStringInfo(m, v12+int32(160), int32(603982), v12+int32(128))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+112)) = int32(65745)
					F_appendStringInfo(m, v12+int32(144), int32(668123), v12+int32(112))
					mBase = m.M
					v84 = m.ExcPending
					if v84 != 0 {
						return
					} else {
						v88 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return
						} else {
							if v88 != 0 {
								if l1 != 0 {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = l3
									*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = l2
									F_errmsg(m, int32(699643), v12+int32(32))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
										F_errdetail_internal(m, int32(206576), v12+int32(16))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
											if v110 != 0 {
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
												F_errhint(m, int32(206576), v12)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													F_errfinish(m, int32(494185), int32(1705), int32(265463))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return
													} else {
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v121)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v124)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											} else {
												F_errfinish(m, int32(494185), int32(1705), int32(265463))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v124)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										}
									}
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = l3
									F_errmsg(m, int32(699599), v12+int32(48))
									mBase = m.M
									v102 = m.ExcPending
									if v102 != 0 {
										return
									} else {
										v103 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v103
										F_errdetail_internal(m, int32(206576), v12+int32(16))
										mBase = m.M
										v109 = m.ExcPending
										if v109 != 0 {
											return
										} else {
											v110 = *(*int32)(unsafe.Add(mBase, uint32(v12)+148))
											if v110 != 0 {
												v111 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
												*(*int32)(unsafe.Add(mBase, uint32(v12))) = v111
												F_errhint(m, int32(206576), v12)
												mBase = m.M
												v115 = m.ExcPending
												if v115 != 0 {
													return
												} else {
													F_errfinish(m, int32(494185), int32(1705), int32(265463))
													mBase = m.M
													v120 = m.ExcPending
													if v120 != 0 {
														return
													} else {
														v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
														F_pfree(m, v121)
														mBase = m.M
														v123 = m.ExcPending
														if v123 != 0 {
															return
														} else {
															v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
															F_pfree(m, v124)
															mBase = m.M
															v126 = m.ExcPending
															if v126 != 0 {
																return
															} else {
																m.G0 = v12 + int32(176)
																return
															}
														}
													}
												}
											} else {
												F_errfinish(m, int32(494185), int32(1705), int32(265463))
												mBase = m.M
												v120 = m.ExcPending
												if v120 != 0 {
													return
												} else {
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
													F_pfree(m, v121)
													mBase = m.M
													v123 = m.ExcPending
													if v123 != 0 {
														return
													} else {
														v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
														F_pfree(m, v124)
														mBase = m.M
														v126 = m.ExcPending
														if v126 != 0 {
															return
														} else {
															m.G0 = v12 + int32(176)
															return
														}
													}
												}
											}
										}
									}
								}
							} else {
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v12)+160))
								F_pfree(m, v121)
								mBase = m.M
								v123 = m.ExcPending
								if v123 != 0 {
									return
								} else {
									v124 = *(*int32)(unsafe.Add(mBase, uint32(v12)+144))
									F_pfree(m, v124)
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return
									} else {
										m.G0 = v12 + int32(176)
										return
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
