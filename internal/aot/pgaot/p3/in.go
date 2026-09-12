package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetupLockInTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int64
	_ = v209
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v17 = *(*int32)(unsafe.Add(mBase, _consts[43]))
	v21 = F_hash_search_with_hash_value(m, v17, l2, l3, int32(3), v14+int32(23))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		if v21 == int32(0) {
			v224 = int32(0)
			m.G0 = v14 + int32(32)
			return v224
		} else {
			v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+23)))
			if v27 == int32(0) {
				v30 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+128)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v30
				*(*int32)(unsafe.Add(mBase, uint32(v21)+40)) = v30
				*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = int64(0)
				v39 = v21 + int32(32)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(v21)+32)) = v39
				v43 = v21 + int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v43
				*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v43
				v46 = int32(40)
				v49 = v21 + int32(44)
				if v49&int32(3) == v30 {
					v55 = v21 + int32(84)
					v57 = v21 + int32(48)
					if base.Ui32(v57) < base.Ui32(v55) {
						v59 = v55
					} else {
						v59 = v57
					}
					v67 = (v59-v21-int32(45))&int32(-4) + int32(4)
				} else {
					v67 = v46
				}
				v71 = F__emscripten_memset_bulkmem(m, v49, base.I32_extend8_s(int32(0)), v67)
				mBase = m.M
				v73 = v21 + int32(88)
				if v73&int32(3) == int32(0) {
					v79 = v21 + int32(128)
					v81 = v21 + int32(92)
					if base.Ui32(v81) < base.Ui32(v79) {
						v83 = v79
					} else {
						v83 = v81
					}
					v91 = (v83-v21-int32(89))&int32(-4) + int32(4)
				} else {
					v91 = v46
				}
				v95 = F__emscripten_memset_bulkmem(m, v73, base.I32_extend8_s(int32(0)), v91)
				mBase = m.M
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = v21
			*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = l1
			v103 = *(*int32)(unsafe.Add(mBase, _consts[856]))
			v112 = F_hash_search_with_hash_value(m, v103, v14+int32(24), l1<<(uint(int32(4))%32)^l3, int32(3), v14+int32(23))
			mBase = m.M
			v113 = m.ExcPending
			if v113 != 0 {
				return int32(0)
			} else {
				if v112 == int32(0) {
					v116 = int32(0)
					v117 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
					if v117 != 0 {
						v224 = v116
						m.G0 = v14 + int32(32)
						return v224
					} else {
						v119 = *(*int32)(unsafe.Add(mBase, _consts[43]))
						v122 = F_hash_search_with_hash_value(m, v119, v21, l3, int32(2), int32(0))
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							if v122 != 0 {
								v224 = v116
								m.G0 = v14 + int32(32)
								return v224
							} else {
								F_errstart_cold(m, int32(23), int32(0))
								mBase = m.M
								v127 = m.ExcPending
								if v127 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(444782), int32(0))
									mBase = m.M
									v131 = m.ExcPending
									if v131 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(497495), int32(1358), int32(397034))
										mBase = m.M
										v136 = m.ExcPending
										if v136 != 0 {
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
					v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+23)))
					if v137 == int32(0) {
						v140 = *(*int32)(unsafe.Add(mBase, uint32(l1)+616))
						*(*int64)(unsafe.Add(mBase, uint32(v112)+12)) = int64(0)
						if v140 != 0 {
							v143 = v140
						} else {
							v143 = l1
						}
						*(*int32)(unsafe.Add(mBase, uint32(v112)+8)) = v143
						v146 = v112 + int32(20)
						v148 = v21 + int32(24)
						v151 = *(*int32)(unsafe.Add(mBase, uint32(v21)+28))
						if v151 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v21)+28)) = v148
							*(*int32)(unsafe.Add(mBase, uint32(v21)+24)) = v148
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v112)+24)) = v148
						v157 = *(*int32)(unsafe.Add(mBase, uint32(v148)))
						*(*int32)(unsafe.Add(mBase, uint32(v112)+20)) = v157
						*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = v146
						*(*int32)(unsafe.Add(mBase, uint32(v148))) = v146
						v162 = v112 + int32(28)
						v165 = l1 + l3&int32(15)<<(uint(int32(3))%32)
						v167 = v165 + int32(148)
						v168 = *(*int32)(unsafe.Add(mBase, uint32(v165)+152))
						if v168 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v167)+4)) = v167
							*(*int32)(unsafe.Add(mBase, uint32(v167))) = v167
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v112)+32)) = v167
						v174 = *(*int32)(unsafe.Add(mBase, uint32(v167)))
						*(*int32)(unsafe.Add(mBase, uint32(v112)+28)) = v174
						*(*int32)(unsafe.Add(mBase, uint32(v174)+4)) = v162
						*(*int32)(unsafe.Add(mBase, uint32(v167))) = v162
					} else {
					}
					v183 = *(*int32)(unsafe.Add(mBase, uint32(v21)+84))
					v184 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v21)+84)) = v183 + v184
					v188 = l4 << (uint(int32(2)) % 32)
					v191 = v21 + v188 + int32(44)
					v192 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
					*(*int32)(unsafe.Add(mBase, uint32(v191))) = v192 + v184
					v196 = *(*int32)(unsafe.Add(mBase, uint32(v112)+12))
					if int32(base.Ui32(v196)>>(uint(l4)%32))&v184 == int32(0) {
						v224 = v112
						m.G0 = v14 + int32(32)
						return v224
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v205 = m.ExcPending
						if v205 != 0 {
							return int32(0)
						} else {
							v206 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v208 = *(*int32)(unsafe.Add(mBase, uint32(v206+v188)))
							v209 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
							v210 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v14)+12)) = v210
							*(*int64)(unsafe.Add(mBase, uint32(v14)+4)) = v209
							*(*int32)(unsafe.Add(mBase, uint32(v14))) = v208
							F_errmsg_internal(m, int32(431756), v14)
							mBase = m.M
							v216 = m.ExcPending
							if v216 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(497495), int32(1449), int32(397034))
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
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
		}
	}
}
func F_find_in_path(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v223 int32
	_ = v223
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v13 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v251 = m.ExcPending
	if v251 != 0 {
		goto L22
	} else {
		goto L76
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L22
	} else {
		goto L72
	}
L3:
	;
	m.G0 = v11 - int32(-64)
	return v223
L4:
	;
	v14 = F_strlen(m, l0)
	mBase = m.M
	v18 = l1
	goto L7
L5:
	;
	goto L6
L6:
	;
	v223 = int32(0)
	goto L3
L7:
	;
	v26 = v18
	goto L10
L8:
	;
	goto L6
L9:
	;
	if v36 == v18 {
		goto L2
	} else {
		goto L17
	}
L10:
	;
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if v28 == int32(0) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L9
L12:
	;
	goto L11
L13:
	;
	v36 = int32(0)
	goto L12
L14:
	;
	goto L15
L15:
	;
	if v28 == int32(58) {
		v36 = v26
		goto L12
	} else {
		goto L16
	}
L16:
	;
	v26 = v26 + int32(1)
	goto L10
L17:
	;
	if v36 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v44 = v42 + int32(1)
	v45 = F_palloc(m, v44)
	mBase = m.M
	v48 = m.ExcPending
	if v48 != 0 {
		goto L22
	} else {
		goto L23
	}
L19:
	;
	v40 = F_strlen(m, v18)
	mBase = m.M
	v42 = v40
	goto L18
L20:
	;
	goto L21
L21:
	;
	v42 = v36 - v18
	goto L18
L22:
	;
	return int32(0)
L23:
	;
	if v44 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v165 = F_substitute_path_macro(m, v45, int32(213184), int32(4512128))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L22
	} else {
		goto L56
	}
L25:
	;
	v160 = F_strlen(m, v156)
	mBase = m.M
	goto L24
L26:
	;
	v156 = v18
	goto L25
L27:
	;
	goto L28
L28:
	;
	v54 = v44 - int32(1)
	if (v45^v18)&int32(3) != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	v153 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v150))) = uint8(v153)
	v156 = v149
	goto L25
L30:
	;
	v134 = v129
	v135 = v130
	v136 = v131
	goto L52
L31:
	;
	if v124 == int32(0) {
		v149 = v122
		v150 = v123
		goto L29
	} else {
		goto L51
	}
L32:
	;
	v122 = v18
	v123 = v45
	v124 = v54
	goto L31
L33:
	;
	goto L34
L34:
	;
	v58 = int32(0)
	if v18&int32(3) == v58 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v91 == int32(0) {
		v149 = v88
		v150 = v89
		goto L29
	} else {
		goto L44
	}
L36:
	;
	v88 = v18
	v89 = v45
	v90 = v54
	v91 = base.B2i32(v54 != v58)
	goto L35
L37:
	;
	if v54 == int32(0) {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v67 = v18
	v68 = v45
	v69 = v54
	goto L39
L39:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	*(*uint8)(unsafe.Add(mBase, uint32(v68))) = uint8(v71)
	if v71 == int32(0) {
		v129 = v67
		v130 = v68
		v131 = v69
		goto L30
	} else {
		goto L41
	}
L40:
	;
	v88 = v82
	v89 = v76
	v90 = v78
	v91 = v80
	goto L35
L41:
	;
	v75 = int32(1)
	v76 = v68 + v75
	v78 = v69 - v75
	v79 = int32(0)
	v80 = base.B2i32(v78 != v79)
	v82 = v67 + v75
	if v82&int32(3) == v79 {
		v88 = v82
		v89 = v76
		v90 = v78
		v91 = v80
		goto L35
	} else {
		goto L42
	}
L42:
	;
	if v78 != 0 {
		v67 = v82
		v68 = v76
		v69 = v78
		goto L39
	} else {
		goto L43
	}
L43:
	;
	goto L40
L44:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v88))))
	if v94 == int32(0) {
		v122 = v88
		v123 = v89
		v124 = v90
		goto L31
	} else {
		goto L45
	}
L45:
	;
	if base.Ui32(v90) < base.Ui32(int32(4)) {
		v122 = v88
		v123 = v89
		v124 = v90
		goto L31
	} else {
		goto L46
	}
L46:
	;
	v100 = v88
	v101 = v89
	v102 = v90
	goto L47
L47:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v100)))
	v108 = int32(-2139062144)
	if (int32(16843008)-v105|v105)&v108 != v108 {
		v129 = v100
		v130 = v101
		v131 = v102
		goto L30
	} else {
		goto L49
	}
L48:
	;
	v122 = v116
	v123 = v114
	v124 = v118
	goto L31
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v101))) = v105
	v113 = int32(4)
	v114 = v101 + v113
	v116 = v100 + v113
	v118 = v102 - v113
	if base.Ui32(int32(3)) < base.Ui32(v118) {
		v100 = v116
		v101 = v114
		v102 = v118
		goto L47
	} else {
		goto L50
	}
L50:
	;
	goto L48
L51:
	;
	v129 = v122
	v130 = v123
	v131 = v124
	goto L30
L52:
	;
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v134))))
	*(*uint8)(unsafe.Add(mBase, uint32(v135))) = uint8(v138)
	if v138 == int32(0) {
		v149 = v134
		v150 = v135
		goto L29
	} else {
		goto L54
	}
L53:
	;
	v149 = v145
	v150 = v143
	goto L29
L54:
	;
	v142 = int32(1)
	v143 = v135 + v142
	v145 = v134 + v142
	v147 = v136 - v142
	if v147 != 0 {
		v134 = v145
		v135 = v143
		v136 = v147
		goto L52
	} else {
		goto L55
	}
L55:
	;
	goto L53
L56:
	;
	F_pfree(m, v45)
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L22
	} else {
		goto L57
	}
L57:
	;
	F_canonicalize_path_enc(m, v165)
	mBase = m.M
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165))))
	if v170 != int32(47) {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	v173 = F_strlen(m, v165)
	mBase = m.M
	v175 = F_palloc(m, v173+(v14+int32(2)))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L22
	} else {
		goto L59
	}
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v165
	v182 = F_pg_sprintf(m, v175, int32(176970), v9+int32(-32))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L22
	} else {
		goto L60
	}
L60:
	;
	F_pfree(m, v165)
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L22
	} else {
		goto L61
	}
L61:
	;
	v188 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L22
	} else {
		goto L62
	}
L62:
	;
	if v188 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v175
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(321394)
	F_errmsg_internal(m, int32(712352), v9+int32(-48))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L22
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v203 = F_pg_file_exists(m, v175)
	mBase = m.M
	v204 = m.ExcPending
	if v204 != 0 {
		goto L22
	} else {
		goto L68
	}
L66:
	;
	F_errfinish(m, int32(494981), int32(631), int32(321394))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L22
	} else {
		goto L67
	}
L67:
	;
	goto L65
L68:
	;
	if v203 != 0 {
		v223 = v175
		goto L3
	} else {
		goto L69
	}
L69:
	;
	F_pfree(m, v175)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L22
	} else {
		goto L70
	}
L70:
	;
	v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v42))))
	if v209 != 0 {
		v18 = v18 + v44
		goto L7
	} else {
		goto L71
	}
L71:
	;
	goto L8
L72:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L22
	} else {
		goto L73
	}
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(321272)
	F_errmsg(m, int32(700610), v11)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L22
	} else {
		goto L74
	}
L74:
	;
	F_errfinish(m, int32(494981), int32(606), int32(321394))
	mBase = m.M
	v247 = m.ExcPending
	if v247 != 0 {
		goto L22
	} else {
		goto L75
	}
L75:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L76:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L22
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(321272)
	F_errmsg(m, int32(321661), v9+int32(-16))
	mBase = m.M
	v261 = m.ExcPending
	if v261 != 0 {
		goto L22
	} else {
		goto L78
	}
L78:
	;
	F_errfinish(m, int32(494981), int32(625), int32(321394))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L22
	} else {
		goto L79
	}
L79:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_in_grouping_U(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v23 int32
	_ = v23
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v75 int32
	_ = v75
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v23 = v13
	goto L2
L1:
	;
	return v113
L2:
	;
	if v14 <= v23 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	v113 = int32(0)
	goto L1
L4:
	;
	return int32(-1)
L5:
	;
	goto L6
L6:
	;
	v31 = int32(1)
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23+v15))))
	if base.Ui32(v33) < base.Ui32(int32(192)) {
		v90 = v33
		v91 = v31
		goto L7
	} else {
		goto L8
	}
L7:
	;
	if l3 < v90 {
		v113 = v91
		goto L1
	} else {
		goto L20
	}
L8:
	;
	v37 = v23 + int32(1)
	if v37 == v14 {
		v90 = v33
		v91 = v31
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v37+v15))))
	v42 = v40 & int32(63)
	if base.Ui32(int32(224)) <= base.Ui32(v33) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v46+v15))))
	v58 = v56 & int32(63)
	if base.Ui32(int32(240)) <= base.Ui32(v33) {
		goto L16
	} else {
		goto L17
	}
L11:
	;
	v46 = v23 + int32(2)
	if v46 != v14 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v90 = v33<<(uint(int32(6))%32)&int32(1984) | v42
	v91 = int32(2)
	goto L7
L14:
	;
	goto L13
L15:
	;
	v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15+v62))))
	v90 = v75&int32(63) | (v33<<(uint(int32(18))%32)&int32(1835008) | v42<<(uint(int32(12))%32) | v58<<(uint(int32(6))%32))
	v91 = int32(4)
	goto L7
L16:
	;
	v62 = v23 + int32(3)
	if v62 != v14 {
		goto L15
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v90 = v33<<(uint(int32(12))%32)&int32(61440) | v42<<(uint(int32(6))%32) | v58
	v91 = int32(3)
	goto L7
L19:
	;
	goto L18
L20:
	;
	v95 = v90 - l2
	if v95 < int32(0) {
		v113 = v91
		goto L1
	} else {
		goto L21
	}
L21:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+int32(base.Ui32(v95)>>(uint(int32(3))%32))))))
	if int32(base.Ui32(v101)>>(uint(v95&int32(7))%32))&int32(1) == int32(0) {
		v113 = v91
		goto L1
	} else {
		goto L22
	}
L22:
	;
	v109 = v91 + v23
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v109
	if l4 != 0 {
		v23 = v109
		goto L2
	} else {
		goto L23
	}
L23:
	;
	goto L3
}
func F_in_range_int8_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int64
	_ = v29
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if int64(0) <= v7 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v13 != 0 {
			v14 = int64(0) - v7
		} else {
			v14 = v7
		}
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
		v19 = v18 + v14
		if base.B2i32(v14 < int64(0)) != base.B2i32(v19 < v18) {
			v22 = int32(0)
			return base.B2i32(v13 != v22) ^ base.B2i32(v10 != v22)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v29 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
			if v10 != 0 {
				return base.B2i32(v29 <= v19)
			} else {
				return base.B2i32(v19 <= v29)
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(251364), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(500014), int32(413), int32(554473))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
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
func F_in_range_numeric_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v91 int32
	_ = v91
	var v108 int32
	_ = v108
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v171 int32
	_ = v171
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v222 int32
	_ = v222
	var v232 int32
	_ = v232
	var v244 int32
	_ = v244
	var v245 int64
	_ = v245
	var v252 int32
	_ = v252
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v356 int32
	_ = v356
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v378 int32
	_ = v378
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int32(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v26 = F_pg_detoast_datum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v29 = F_pg_detoast_datum(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+4)))
				if v31 == int32(49152) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v366 = m.ExcPending
					if v366 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50593922))
						mBase = m.M
						v369 = m.ExcPending
						if v369 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(251364), int32(0))
							mBase = m.M
							v373 = m.ExcPending
							if v373 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(499844), int32(2699), int32(490620))
								mBase = m.M
								v378 = m.ExcPending
								if v378 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					if v31 == int32(61440) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v366 = m.ExcPending
						if v366 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50593922))
							mBase = m.M
							v369 = m.ExcPending
							if v369 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(251364), int32(0))
								mBase = m.M
								v373 = m.ExcPending
								if v373 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(499844), int32(2699), int32(490620))
									mBase = m.M
									v378 = m.ExcPending
									if v378 != 0 {
										return int32(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v36 = int32(49152)
						v37 = v31 & v36
						if v37 != v36 {
							if v37 != int32(32768) {
								v48 = v37
							} else {
								v48 = v31 << (uint(int32(1)) % 32) & int32(16384)
							}
						} else {
							v48 = v31 & int32(61440)
						}
						if v48 == int32(16384) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v366 = m.ExcPending
							if v366 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50593922))
								mBase = m.M
								v369 = m.ExcPending
								if v369 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(251364), int32(0))
									mBase = m.M
									v373 = m.ExcPending
									if v373 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(499844), int32(2699), int32(490620))
										mBase = m.M
										v378 = m.ExcPending
										if v378 != 0 {
											return int32(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
							v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
							v53 = base.I32_extend16_s(v52)
							v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
							if v59 == int32(49152) {
								v344 = base.B2i32(v51 == int32(0)) | base.B2i32(v53 == int32(-16384))
								v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
								if v345 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v348 = m.ExcPending
									if v348 != 0 {
										return int32(0)
									} else {
										v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v349 != v26 {
											F_pfree(m, v26)
											mBase = m.M
											v352 = m.ExcPending
											if v352 != 0 {
												return int32(0)
											} else {
												v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v353 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v356 = m.ExcPending
													if v356 != 0 {
														return int32(0)
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											}
										} else {
											v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v353 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return int32(0)
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											} else {
												m.G0 = v18 + int32(96)
												return v344
											}
										}
									}
								} else {
									v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
									if v349 != v26 {
										F_pfree(m, v26)
										mBase = m.M
										v352 = m.ExcPending
										if v352 != 0 {
											return int32(0)
										} else {
											v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v353 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return int32(0)
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											} else {
												m.G0 = v18 + int32(96)
												return v344
											}
										}
									} else {
										v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
										if v353 != v29 {
											F_pfree(m, v29)
											mBase = m.M
											v356 = m.ExcPending
											if v356 != 0 {
												return int32(0)
											} else {
												m.G0 = v18 + int32(96)
												return v344
											}
										} else {
											m.G0 = v18 + int32(96)
											return v344
										}
									}
								}
							} else {
								if v53 == int32(-16384) {
									v344 = base.B2i32(v51 != int32(0))
									v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
									if v345 != v21 {
										F_pfree(m, v21)
										mBase = m.M
										v348 = m.ExcPending
										if v348 != 0 {
											return int32(0)
										} else {
											v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if v349 != v26 {
												F_pfree(m, v26)
												mBase = m.M
												v352 = m.ExcPending
												if v352 != 0 {
													return int32(0)
												} else {
													v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v353 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v356 = m.ExcPending
														if v356 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												}
											} else {
												v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v353 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v356 = m.ExcPending
													if v356 != 0 {
														return int32(0)
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											}
										}
									} else {
										v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
										if v349 != v26 {
											F_pfree(m, v26)
											mBase = m.M
											v352 = m.ExcPending
											if v352 != 0 {
												return int32(0)
											} else {
												v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v353 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v356 = m.ExcPending
													if v356 != 0 {
														return int32(0)
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											}
										} else {
											v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
											if v353 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return int32(0)
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											} else {
												m.G0 = v18 + int32(96)
												return v344
											}
										}
									}
								} else {
									v66 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
									v67 = base.I32_extend16_s(v59)
									v68 = base.I32_extend16_s(v31)
									if base.Ui32(int32(-16384)) <= base.Ui32(v68) {
										if v66 != 0 {
											if v53 != int32(-12288) {
												v344 = base.B2i32(v51 == int32(0)) | base.B2i32(v67 == int32(-4096))
											} else {
												v344 = int32(1)
											}
										} else {
											if v53 != int32(-4096) {
												v344 = base.B2i32(v67 == int32(-12288)) | base.B2i32(v51 != int32(0))
											} else {
												v344 = int32(1)
											}
										}
										v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v345 != v21 {
											F_pfree(m, v21)
											mBase = m.M
											v348 = m.ExcPending
											if v348 != 0 {
												return int32(0)
											} else {
												v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												if v349 != v26 {
													F_pfree(m, v26)
													mBase = m.M
													v352 = m.ExcPending
													if v352 != 0 {
														return int32(0)
													} else {
														v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v353 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v356 = m.ExcPending
															if v356 != 0 {
																return int32(0)
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													}
												} else {
													v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v353 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v356 = m.ExcPending
														if v356 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												}
											}
										} else {
											v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
											if v349 != v26 {
												F_pfree(m, v26)
												mBase = m.M
												v352 = m.ExcPending
												if v352 != 0 {
													return int32(0)
												} else {
													v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v353 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v356 = m.ExcPending
														if v356 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												}
											} else {
												v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
												if v353 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v356 = m.ExcPending
													if v356 != 0 {
														return int32(0)
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												} else {
													m.G0 = v18 + int32(96)
													return v344
												}
											}
										}
									} else {
										if base.Ui32(int32(-16384)) <= base.Ui32(v67) {
											v91 = int32(-12288)
											if v67 == v91 {
												v344 = base.B2i32(v51 == int32(0)) | base.B2i32(v53 == v91)
											} else {
												v344 = base.B2i32(v53 == int32(-4096)) | base.B2i32(v51 != int32(0))
											}
											v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
											if v345 != v21 {
												F_pfree(m, v21)
												mBase = m.M
												v348 = m.ExcPending
												if v348 != 0 {
													return int32(0)
												} else {
													v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													if v349 != v26 {
														F_pfree(m, v26)
														mBase = m.M
														v352 = m.ExcPending
														if v352 != 0 {
															return int32(0)
														} else {
															v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															if v353 != v29 {
																F_pfree(m, v29)
																mBase = m.M
																v356 = m.ExcPending
																if v356 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v18 + int32(96)
																	return v344
																}
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														}
													} else {
														v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v353 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v356 = m.ExcPending
															if v356 != 0 {
																return int32(0)
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													}
												}
											} else {
												v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
												if v349 != v26 {
													F_pfree(m, v26)
													mBase = m.M
													v352 = m.ExcPending
													if v352 != 0 {
														return int32(0)
													} else {
														v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v353 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v356 = m.ExcPending
															if v356 != 0 {
																return int32(0)
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													}
												} else {
													v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
													if v353 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v356 = m.ExcPending
														if v356 != 0 {
															return int32(0)
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													} else {
														m.G0 = v18 + int32(96)
														return v344
													}
												}
											}
										} else {
											if base.Ui32(int32(-16384)) <= base.Ui32(v53) {
												v344 = base.B2i32(v53 == int32(-4096)) ^ base.B2i32(v51 != int32(0))
												v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
												if v345 != v21 {
													F_pfree(m, v21)
													mBase = m.M
													v348 = m.ExcPending
													if v348 != 0 {
														return int32(0)
													} else {
														v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
														if v349 != v26 {
															F_pfree(m, v26)
															mBase = m.M
															v352 = m.ExcPending
															if v352 != 0 {
																return int32(0)
															} else {
																v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																if v353 != v29 {
																	F_pfree(m, v29)
																	mBase = m.M
																	v356 = m.ExcPending
																	if v356 != 0 {
																		return int32(0)
																	} else {
																		m.G0 = v18 + int32(96)
																		return v344
																	}
																} else {
																	m.G0 = v18 + int32(96)
																	return v344
																}
															}
														} else {
															v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															if v353 != v29 {
																F_pfree(m, v29)
																mBase = m.M
																v356 = m.ExcPending
																if v356 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v18 + int32(96)
																	return v344
																}
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														}
													}
												} else {
													v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
													if v349 != v26 {
														F_pfree(m, v26)
														mBase = m.M
														v352 = m.ExcPending
														if v352 != 0 {
															return int32(0)
														} else {
															v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
															if v353 != v29 {
																F_pfree(m, v29)
																mBase = m.M
																v356 = m.ExcPending
																if v356 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v18 + int32(96)
																	return v344
																}
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														}
													} else {
														v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
														if v353 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v356 = m.ExcPending
															if v356 != 0 {
																return int32(0)
															} else {
																m.G0 = v18 + int32(96)
																return v344
															}
														} else {
															m.G0 = v18 + int32(96)
															return v344
														}
													}
												}
											} else {
												v108 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
												v114 = base.B2i32(int32(0) <= v67)
												if int32(0) <= v67 {
													v115 = int32(-8)
												} else {
													v115 = int32(-6)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(base.Ui32(int32(base.Ui32(v108)>>(uint(int32(2))%32))+v115) >> (uint(int32(1)) % 32))
												if int32(0) <= v67 {
													v120 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
													v130 = v120
												} else {
													v130 = v59<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v59&int32(63)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v130
												v132 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v132
												v141 = base.B2i32(v67 < v132)
												if v67 < v132 {
													v142 = int32(base.Ui32(v59)>>(uint(int32(7))%32)) & int32(63)
												} else {
													v142 = v59 & int32(16383)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v142
												v149 = v59 & int32(49152)
												if v149 == int32(32768) {
													v152 = v59 << (uint(int32(1)) % 32) & int32(16384)
												} else {
													v152 = v149
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v152
												if v67 < v132 {
													v156 = int32(6)
												} else {
													v156 = int32(8)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v21 + v156
												v159 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
												v165 = base.B2i32(int32(0) <= v53)
												if int32(0) <= v53 {
													v166 = int32(-8)
												} else {
													v166 = int32(-6)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = int32(base.Ui32(int32(base.Ui32(v159)>>(uint(int32(2))%32))+v166) >> (uint(int32(1)) % 32))
												if int32(0) <= v53 {
													v171 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+6)))
													v181 = v171
												} else {
													v181 = v52<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v52&int32(63)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v181
												v183 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v183
												v192 = base.B2i32(v53 < v183)
												if v53 < v183 {
													v193 = int32(base.Ui32(v52)>>(uint(int32(7))%32)) & int32(63)
												} else {
													v193 = v52 & int32(16383)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v193
												v200 = v52 & int32(49152)
												if v200 == int32(32768) {
													v203 = v52 << (uint(int32(1)) % 32) & int32(16384)
												} else {
													v203 = v200
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v203
												if v53 < v183 {
													v207 = int32(6)
												} else {
													v207 = int32(8)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v26 + v207
												v210 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
												v216 = base.B2i32(int32(0) <= v68)
												if int32(0) <= v68 {
													v217 = int32(-8)
												} else {
													v217 = int32(-6)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = int32(base.Ui32(int32(base.Ui32(v210)>>(uint(int32(2))%32))+v217) >> (uint(int32(1)) % 32))
												if int32(0) <= v68 {
													v222 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+6)))
													v232 = v222
												} else {
													v232 = v31<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v31&int32(63)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v232
												if v37 != int32(49152) {
													if v37 != int32(32768) {
														v244 = v37
													} else {
														v244 = v31 << (uint(int32(1)) % 32) & int32(16384)
													}
												} else {
													v244 = v31 & int32(61440)
												}
												v245 = int64(0)
												*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v245
												*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v245
												*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v244
												*(*int64)(unsafe.Add(mBase, uint32(v18))) = v245
												v252 = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v252
												v261 = base.B2i32(v68 < v252)
												if v68 < v252 {
													v262 = int32(base.Ui32(v31)>>(uint(int32(7))%32)) & int32(63)
												} else {
													v262 = v31 & int32(16383)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v262
												if v68 < v252 {
													v266 = int32(6)
												} else {
													v266 = int32(8)
												}
												*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v29 + v266
												if v66 != 0 {
													F_sub_var(m, v18+int32(48), v18+int32(24), v18)
													mBase = m.M
													v274 = m.ExcPending
													if v274 != 0 {
														return int32(0)
													} else {
														v282 = v18 + int32(72)
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
														v290 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
														if v291 == int32(0) {
															if v290 == int32(0) {
																v327 = int32(0)
															} else {
																if v289 == int32(16384) {
																	v301 = int32(1)
																} else {
																	v301 = int32(-1)
																}
																v327 = v301
															}
														} else {
															v302 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
															if v290 == int32(0) {
																if v302 != 0 {
																	v307 = int32(-1)
																} else {
																	v307 = int32(1)
																}
																v327 = v307
															} else {
																v308 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
																v309 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
																v310 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
																v311 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
																if v302 == int32(0) {
																	if v289 == int32(16384) {
																		v327 = int32(1)
																	} else {
																		v317 = F_cmp_abs_common(m, v311, v291, v310, v309, v290, v308)
																		mBase = m.M
																		v327 = v317
																	}
																} else {
																	if v289 == int32(0) {
																		v327 = int32(-1)
																	} else {
																		v321 = F_cmp_abs_common(m, v309, v290, v308, v311, v291, v310)
																		mBase = m.M
																		v327 = v321
																	}
																}
															}
														}
														v328 = int32(0)
														v332 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
														if v332 != 0 {
															F_pfree(m, v332)
															mBase = m.M
															v334 = m.ExcPending
															if v334 != 0 {
																return int32(0)
															} else {
																if v51 != 0 {
																	v335 = base.B2i32(v327 <= v328)
																} else {
																	v335 = base.B2i32(v328 <= v327)
																}
																v344 = v335
																v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																if v345 != v21 {
																	F_pfree(m, v21)
																	mBase = m.M
																	v348 = m.ExcPending
																	if v348 != 0 {
																		return int32(0)
																	} else {
																		v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																		if v349 != v26 {
																			F_pfree(m, v26)
																			mBase = m.M
																			v352 = m.ExcPending
																			if v352 != 0 {
																				return int32(0)
																			} else {
																				v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																				if v353 != v29 {
																					F_pfree(m, v29)
																					mBase = m.M
																					v356 = m.ExcPending
																					if v356 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v18 + int32(96)
																						return v344
																					}
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			}
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	}
																} else {
																	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	if v349 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v352 = m.ExcPending
																		if v352 != 0 {
																			return int32(0)
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																}
															}
														} else {
															if v51 != 0 {
																v335 = base.B2i32(v327 <= v328)
															} else {
																v335 = base.B2i32(v328 <= v327)
															}
															v344 = v335
															v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															if v345 != v21 {
																F_pfree(m, v21)
																mBase = m.M
																v348 = m.ExcPending
																if v348 != 0 {
																	return int32(0)
																} else {
																	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	if v349 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v352 = m.ExcPending
																		if v352 != 0 {
																			return int32(0)
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																}
															} else {
																v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																if v349 != v26 {
																	F_pfree(m, v26)
																	mBase = m.M
																	v352 = m.ExcPending
																	if v352 != 0 {
																		return int32(0)
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																} else {
																	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																	if v353 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v356 = m.ExcPending
																		if v356 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return v344
																	}
																}
															}
														}
													}
												} else {
													F_add_var(m, v18+int32(48), v18+int32(24), v18)
													mBase = m.M
													v280 = m.ExcPending
													if v280 != 0 {
														return int32(0)
													} else {
														v282 = v18 + int32(72)
														v289 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
														v290 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
														v291 = *(*int32)(unsafe.Add(mBase, uint32(v282)))
														if v291 == int32(0) {
															if v290 == int32(0) {
																v327 = int32(0)
															} else {
																if v289 == int32(16384) {
																	v301 = int32(1)
																} else {
																	v301 = int32(-1)
																}
																v327 = v301
															}
														} else {
															v302 = *(*int32)(unsafe.Add(mBase, uint32(v282)+8))
															if v290 == int32(0) {
																if v302 != 0 {
																	v307 = int32(-1)
																} else {
																	v307 = int32(1)
																}
																v327 = v307
															} else {
																v308 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
																v309 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
																v310 = *(*int32)(unsafe.Add(mBase, uint32(v282)+4))
																v311 = *(*int32)(unsafe.Add(mBase, uint32(v282)+20))
																if v302 == int32(0) {
																	if v289 == int32(16384) {
																		v327 = int32(1)
																	} else {
																		v317 = F_cmp_abs_common(m, v311, v291, v310, v309, v290, v308)
																		mBase = m.M
																		v327 = v317
																	}
																} else {
																	if v289 == int32(0) {
																		v327 = int32(-1)
																	} else {
																		v321 = F_cmp_abs_common(m, v309, v290, v308, v311, v291, v310)
																		mBase = m.M
																		v327 = v321
																	}
																}
															}
														}
														v328 = int32(0)
														v332 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
														if v332 != 0 {
															F_pfree(m, v332)
															mBase = m.M
															v334 = m.ExcPending
															if v334 != 0 {
																return int32(0)
															} else {
																if v51 != 0 {
																	v335 = base.B2i32(v327 <= v328)
																} else {
																	v335 = base.B2i32(v328 <= v327)
																}
																v344 = v335
																v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
																if v345 != v21 {
																	F_pfree(m, v21)
																	mBase = m.M
																	v348 = m.ExcPending
																	if v348 != 0 {
																		return int32(0)
																	} else {
																		v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																		if v349 != v26 {
																			F_pfree(m, v26)
																			mBase = m.M
																			v352 = m.ExcPending
																			if v352 != 0 {
																				return int32(0)
																			} else {
																				v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																				if v353 != v29 {
																					F_pfree(m, v29)
																					mBase = m.M
																					v356 = m.ExcPending
																					if v356 != 0 {
																						return int32(0)
																					} else {
																						m.G0 = v18 + int32(96)
																						return v344
																					}
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			}
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	}
																} else {
																	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	if v349 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v352 = m.ExcPending
																		if v352 != 0 {
																			return int32(0)
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																}
															}
														} else {
															if v51 != 0 {
																v335 = base.B2i32(v327 <= v328)
															} else {
																v335 = base.B2i32(v328 <= v327)
															}
															v344 = v335
															v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
															if v345 != v21 {
																F_pfree(m, v21)
																mBase = m.M
																v348 = m.ExcPending
																if v348 != 0 {
																	return int32(0)
																} else {
																	v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																	if v349 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v352 = m.ExcPending
																		if v352 != 0 {
																			return int32(0)
																		} else {
																			v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																			if v353 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v356 = m.ExcPending
																				if v356 != 0 {
																					return int32(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return v344
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		}
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																}
															} else {
																v349 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
																if v349 != v26 {
																	F_pfree(m, v26)
																	mBase = m.M
																	v352 = m.ExcPending
																	if v352 != 0 {
																		return int32(0)
																	} else {
																		v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																		if v353 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v356 = m.ExcPending
																			if v356 != 0 {
																				return int32(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return v344
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	}
																} else {
																	v353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
																	if v353 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v356 = m.ExcPending
																		if v356 != 0 {
																			return int32(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return v344
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return v344
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
							}
						}
					}
				}
			}
		}
	}
}
func F_in_range_time_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if int64(0) <= v7 {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v15 != 0 {
			v45 = v12 - v7
			if v10 != 0 {
				return base.B2i32(v14 <= v45)
			} else {
				return base.B2i32(v45 <= v14)
			}
		} else {
			v19 = v7 + v12
			if base.B2i32(v7 < int64(0))^base.B2i32(v19 < v12) == int32(0) {
				v45 = v19
				if v10 != 0 {
					return base.B2i32(v14 <= v45)
				} else {
					return base.B2i32(v45 <= v14)
				}
			} else {
				return base.B2i32(v10 != int32(0))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(251364), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498487), int32(2180), int32(309304))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
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
func F_in_range_timetz_interval(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v48 int64
	_ = v48
	var v50 int64
	_ = v50
	var v51 int64
	_ = v51
	var v52 int32
	_ = v52
	var v56 int64
	_ = v56
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	if int64(0) <= v8 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		if v15 != 0 {
			v45 = v14 - v8
			v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v48 = int64(1000000)
			v50 = base.I64_extend_i32_s(v46)*v48 + v45
			v51 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			v56 = v51 + base.I64_extend_i32_s(v52)*v48
			if v11 != 0 {
				if v50 < v56 {
					return int32(0)
				} else {
					return base.B2i32(v52 <= v46) | base.B2i32(v56 < v50)
				}
			} else {
				if v50 < v56 {
					return int32(1)
				} else {
					return base.B2i32(v46 <= v52) & base.B2i32(v50 <= v56)
				}
			}
		} else {
			v19 = v8 + v14
			if base.B2i32(v8 < int64(0))^base.B2i32(v19 < v14) == int32(0) {
				v45 = v19
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				v48 = int64(1000000)
				v50 = base.I64_extend_i32_s(v46)*v48 + v45
				v51 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v56 = v51 + base.I64_extend_i32_s(v52)*v48
				if v11 != 0 {
					if v50 < v56 {
						return int32(0)
					} else {
						return base.B2i32(v52 <= v46) | base.B2i32(v56 < v50)
					}
				} else {
					if v50 < v56 {
						return int32(1)
					} else {
						return base.B2i32(v46 <= v52) & base.B2i32(v50 <= v56)
					}
				}
			} else {
				return base.B2i32(v11 != int32(0))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				F_errmsg(m, int32(251364), int32(0))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(498487), int32(2732), int32(308694))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
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
