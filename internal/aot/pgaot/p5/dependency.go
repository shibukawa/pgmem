package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_deleteDependencyRecordsFor(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v14 = F_table_open(m, int32(2608), int32(3))
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
	F_ScanKeyInit(m, v10, int32(1), int32(3), int32(184), l0)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	F_ScanKeyInit(m, v10+int32(48), int32(2), int32(3), int32(184), l1)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v34 = F_systable_beginscan(m, v14, int32(2673), int32(1), int32(0), int32(2), v10)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v41 = int32(0)
	goto L6
L6:
	;
	v43 = F_systable_getnext(m, v34)
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	F_systable_endscan(m, v34)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L17
	}
L8:
	;
	if v43 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	if l2 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	goto L7
L12:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45)+22)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v45+v46)+24)))
	if v48 == int32(101) {
		goto L6
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	F_CatalogTupleDelete(m, v14, v43+int32(4))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L16
	}
L15:
	;
	goto L14
L16:
	;
	v41 = v41 + int32(1)
	goto L6
L17:
	;
	F_sequence_close(m, v14, int32(3))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	m.G0 = v10 + int32(96)
	return v41
}
func F_dependency_is_compatible_clause(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
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
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v207 int32
	_ = v207
	v4 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v12 == int32(318) {
		goto L5
	} else {
		goto L6
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v207
L2:
	;
	if v185 == int32(27) {
		goto L64
	} else {
		goto L65
	}
L3:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	v184 = v180
	v185 = v183
	goto L2
L4:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v155)+20)))
	if v157 != int32(1) {
		v207 = v4
		goto L1
	} else {
		goto L57
	}
L5:
	;
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+10)))
	if v15 != 0 {
		v207 = v4
		goto L1
	} else {
		goto L8
	}
L6:
	;
	v71 = l0
	v72 = v12
	goto L7
L7:
	;
	switch v72 - int32(17) {
	case 0:
		goto L30
	default:
		v184 = v71
		v185 = v72
		goto L2
	case 3:
		v155 = v71
		goto L4
	case 4:
		goto L29
	}
L8:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v16 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	if v63 != int32(1) {
		v207 = v4
		goto L1
	} else {
		goto L25
	}
L10:
	;
	v63 = int32(0)
	goto L9
L11:
	;
	goto L12
L12:
	;
	v25 = int32(1)
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v26 <= v25 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v29 = v25
	goto L15
L14:
	;
	v29 = v26
	goto L15
L15:
	;
	v32 = int32(0)
	v34 = v32
	v35 = v32
	goto L16
L16:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(8)+v34<<(uint(int32(2))%32))))
	if v43 != 0 {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v63 = v56
	goto L9
L18:
	;
	goto L17
L19:
	;
	v44 = int32(2)
	if v35 != 0 {
		v56 = v44
		goto L18
	} else {
		goto L22
	}
L20:
	;
	v49 = v35
	goto L21
L21:
	;
	v52 = v34 + int32(1)
	if v52 != v29 {
		v34 = v52
		v35 = v49
		goto L16
	} else {
		goto L24
	}
L22:
	;
	v45 = int32(1)
	if base.Ui32(v45) < base.Ui32(base.I32_popcnt(v43)) {
		v56 = v44
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v49 = v45
	goto L21
L24:
	;
	v56 = v49
	goto L18
L25:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v66 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v155 = int32(0)
	goto L4
L27:
	;
	goto L28
L28:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	v71 = v66
	v72 = v70
	goto L7
L29:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	switch v107 - int32(1) {
	case 0:
		goto L43
	case 1:
		goto L42
	default:
		v180 = v71
		goto L3
	}
L30:
	;
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
	if v75 == int32(0) {
		v207 = v4
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+4))
	if v78 != int32(2) {
		v207 = v4
		goto L1
	} else {
		goto L32
	}
L32:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v83 = F_is_pseudo_constant_clause(m, v82)
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	return int32(0)
L34:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	if v83 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v88)))
	v92 = F_is_pseudo_constant_clause(m, v91)
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L33
	} else {
		goto L38
	}
L36:
	;
	v100 = v88
	goto L37
L37:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v71)+4))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v103 = F_get_oprrest(m, v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L33
	} else {
		goto L40
	}
L38:
	;
	if v92 == int32(0) {
		v207 = v4
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v71)+28))
	v97 = *(*int32)(unsafe.Add(mBase, uint32(v96)+12))
	v100 = v97 + int32(4)
	goto L37
L40:
	;
	if v103 == int32(101) {
		v180 = v102
		goto L3
	} else {
		goto L41
	}
L41:
	;
	v207 = v4
	goto L1
L42:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+12))
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v180 = v154
	goto L3
L43:
	;
	v110 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v110)
	v113 = int32(1)
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v71)+8))
	if v114 == v110 {
		v207 = v113
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v117 <= int32(0) {
		v207 = v113
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v123 = v110
	goto L46
L46:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v127+v123<<(uint(int32(2))%32))))
	v134 = F_dependency_is_compatible_clause(m, v131, l1, v10+int32(14))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L33
	} else {
		goto L49
	}
L47:
	;
	v207 = int32(0)
	goto L1
L48:
	;
	goto L47
L49:
	;
	if v134 == int32(0) {
		goto L48
	} else {
		goto L50
	}
L50:
	;
	v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+14)))
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2))))
	if v139 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v144 = int32(1)
	v146 = v123 + v144
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v114)+4))
	if v146 < v147 {
		v123 = v146
		goto L46
	} else {
		goto L56
	}
L52:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v138)
	goto L51
L53:
	;
	goto L54
L54:
	;
	if v138 != v139 {
		goto L48
	} else {
		goto L55
	}
L55:
	;
	goto L51
L56:
	;
	v207 = v144
	goto L1
L57:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	if v160 == int32(0) {
		v207 = v4
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
	if v163 != int32(2) {
		v207 = v4
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v166)+4))
	v168 = F_is_pseudo_constant_clause(m, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L33
	} else {
		goto L60
	}
L60:
	;
	if v168 == int32(0) {
		v207 = v4
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v155)+4))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+12))
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v174)))
	v176 = F_get_oprrest(m, v172)
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L33
	} else {
		goto L62
	}
L62:
	;
	if v176 != int32(101) {
		v207 = v4
		goto L1
	} else {
		goto L63
	}
L63:
	;
	v180 = v175
	goto L3
L64:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v191 = v189
	v192 = v190
	goto L66
L65:
	;
	v191 = v184
	v192 = v185
	goto L66
L66:
	;
	if v192 != int32(6) {
		v207 = v4
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v191)+4))
	if v195 != l1 {
		v207 = v4
		goto L1
	} else {
		goto L68
	}
L68:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v191)+28))
	if v197 != 0 {
		v207 = v4
		goto L1
	} else {
		goto L69
	}
L69:
	;
	v198 = int32(*(*int16)(unsafe.Add(mBase, uint32(v191)+8)))
	if v198 <= int32(0) {
		v207 = v4
		goto L1
	} else {
		goto L70
	}
L70:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(l2))) = uint16(v198)
	v207 = int32(1)
	goto L1
}
func F_recordDependencyOn(m *base.Module, l0 int32, l1 int32, l2 int32) {
	var v6 int32
	_ = v6
	F_recordMultipleDependencies(m, l0, l1, int32(1), l2)
	v6 = m.ExcPending
	if v6 != 0 {
		return
	} else {
		return
	}
}
func F_recordDependencyOnCurrentExtension(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v83 int32
	_ = v83
	v4 = m.G0
	v6 = v4 - int32(48)
	m.G0 = v6
	v9 = int32(*(*uint8)(unsafe.Add(mBase, _consts[256])))
	if v9 != int32(1) {
		m.G0 = v6 + int32(48)
		return
	} else {
		if l1 != 0 {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = F_getExtensionOfObject(m, v12, v13)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				if v14 != 0 {
					v17 = *(*int32)(unsafe.Add(mBase, _consts[222]))
					if v14 == v17 {
						m.G0 = v6 + int32(48)
						return
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v22 = m.ExcPending
						if v22 != 0 {
							return
						} else {
							F_errcode(m, int32(325))
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return
							} else {
								v27 = F_getObjectDescription(m, l0, int32(0))
								mBase = m.M
								v28 = m.ExcPending
								if v28 != 0 {
									return
								} else {
									v29 = F_get_extension_name(m, v14)
									mBase = m.M
									v30 = m.ExcPending
									if v30 != 0 {
										return
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v6)+20)) = v29
										*(*int32)(unsafe.Add(mBase, uint32(v6)+16)) = v27
										F_errmsg(m, int32(741405), v6+int32(16))
										mBase = m.M
										v37 = m.ExcPending
										if v37 != 0 {
											return
										} else {
											F_errfinish(m, int32(522500), int32(227), int32(283754))
											mBase = m.M
											v42 = m.ExcPending
											if v42 != 0 {
												return
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
					}
				} else {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						F_errcode(m, int32(325))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return
						} else {
							v51 = F_getObjectDescription(m, l0, int32(0))
							mBase = m.M
							v52 = m.ExcPending
							if v52 != 0 {
								return
							} else {
								v54 = *(*int32)(unsafe.Add(mBase, _consts[222]))
								v55 = F_get_extension_name(m, v54)
								mBase = m.M
								v56 = m.ExcPending
								if v56 != 0 {
									return
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v6)+4)) = v55
									*(*int32)(unsafe.Add(mBase, uint32(v6))) = v51
									F_errmsg(m, int32(741446), v6)
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										F_errdetail(m, int32(638952), int32(0))
										mBase = m.M
										v65 = m.ExcPending
										if v65 != 0 {
											return
										} else {
											F_errfinish(m, int32(522500), int32(235), int32(283754))
											mBase = m.M
											v70 = m.ExcPending
											if v70 != 0 {
												return
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
					}
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6)+36)) = int32(3079)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+44)) = int32(0)
			v76 = *(*int32)(unsafe.Add(mBase, _consts[222]))
			*(*int32)(unsafe.Add(mBase, uint32(v6)+40)) = v76
			F_recordMultipleDependencies(m, l0, v6+int32(36), int32(1), int32(101))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return
			} else {
				m.G0 = v6 + int32(48)
				return
			}
		}
	}
}
func F_recordDependencyOnNewAcl(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l3 != 0 {
		v10 = int32(0)
		v15 = F_aclmembers(m, l3, v8+int32(12))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			F_updateAclDependencies(m, l0, l1, v10, l2, v10, v10, v15, v17)
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	} else {
		m.G0 = v8 + int32(16)
		return
	}
}
