package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_SetupLockInTable(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v35 int64
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v69 int32
	_ = v69
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v174 int64
	_ = v174
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	v11 = m.G0
	v13 = v11 - int32(32)
	m.G0 = v13
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_SetupLockInTable[0]))
	v20 = F_hash_search_with_hash_value(m, v16, l2, l3, int32(3), v13+int32(23))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		return int32(0)
	} else {
		if v20 == int32(0) {
			v189 = int32(0)
			m.G0 = v13 + int32(32)
			return v189
		} else {
			v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+23)))
			if v26 == int32(0) {
				v29 = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+128)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v29
				*(*int32)(unsafe.Add(mBase, uint32(v20)+40)) = v29
				v35 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v20)+16)) = v35
				v38 = v20 + int32(32)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+36)) = v38
				*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v38
				v42 = v20 + int32(24)
				*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v42
				*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v42
				*(*int64)(unsafe.Add(mBase, uint32(v20)+44)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v20)+52)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v20)+60)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v20)+68)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v20)+76)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v20)+88)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v20)+96)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v20)+104)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v20)+112)) = v35
				*(*int64)(unsafe.Add(mBase, uint32(v20)+120)) = v35
			} else {
			}
			*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v13)+28)) = l1
			v69 = *(*int32)(unsafe.Add(mBase, _c_F_SetupLockInTable[1]))
			v78 = F_hash_search_with_hash_value(m, v69, v13+int32(24), l1<<(uint(int32(4))%32)^l3, int32(3), v13+int32(23))
			mBase = m.M
			v79 = m.ExcPending
			if v79 != 0 {
				return int32(0)
			} else {
				if v78 == int32(0) {
					v82 = int32(0)
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
					if v83 != 0 {
						v189 = v82
						m.G0 = v13 + int32(32)
						return v189
					} else {
						v85 = *(*int32)(unsafe.Add(mBase, _c_F_SetupLockInTable[0]))
						v88 = F_hash_search_with_hash_value(m, v85, v20, l3, int32(2), int32(0))
						mBase = m.M
						v89 = m.ExcPending
						if v89 != 0 {
							return int32(0)
						} else {
							if v88 != 0 {
								v189 = v82
								m.G0 = v13 + int32(32)
								return v189
							} else {
								F_errstart_cold(m, int32(24), int32(0))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int32(0)
								} else {
									F_errmsg_internal(m, int32(_a_F_SetupLockInTable_0), int32(0))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_SetupLockInTable_1), int32(1399), int32(_a_F_SetupLockInTable_2))
										mBase = m.M
										v102 = m.ExcPending
										if v102 != 0 {
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
					v103 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+23)))
					if v103 == int32(0) {
						v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+364))
						*(*int64)(unsafe.Add(mBase, uint32(v78)+12)) = int64(0)
						if v106 != 0 {
							v109 = v106
						} else {
							v109 = l1
						}
						*(*int32)(unsafe.Add(mBase, uint32(v78)+8)) = v109
						v112 = v20 + int32(24)
						v113 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
						if v113 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v112
							*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v112
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v78)+24)) = v112
						v119 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
						*(*int32)(unsafe.Add(mBase, uint32(v78)+20)) = v119
						v122 = v78 + int32(20)
						*(*int32)(unsafe.Add(mBase, uint32(v119)+4)) = v122
						*(*int32)(unsafe.Add(mBase, uint32(v112))) = v122
						v129 = l1 + l3&int32(15)<<(uint(int32(3))%32)
						v131 = v129 + int32(420)
						v132 = *(*int32)(unsafe.Add(mBase, uint32(v129)+424))
						if v132 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(v131)+4)) = v131
							*(*int32)(unsafe.Add(mBase, uint32(v131))) = v131
						} else {
						}
						*(*int32)(unsafe.Add(mBase, uint32(v78)+32)) = v131
						v138 = *(*int32)(unsafe.Add(mBase, uint32(v131)))
						*(*int32)(unsafe.Add(mBase, uint32(v78)+28)) = v138
						v141 = v78 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(v138)+4)) = v141
						*(*int32)(unsafe.Add(mBase, uint32(v131))) = v141
					} else {
					}
					v148 = *(*int32)(unsafe.Add(mBase, uint32(v20)+84))
					v149 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v20)+84)) = v148 + v149
					v153 = l4 << (uint(int32(2)) % 32)
					v156 = v20 + v153 + int32(44)
					v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
					*(*int32)(unsafe.Add(mBase, uint32(v156))) = v157 + v149
					v161 = *(*int32)(unsafe.Add(mBase, uint32(v78)+12))
					if int32(base.Ui32(v161)>>(uint(l4)%32))&v149 == int32(0) {
						v189 = v78
						m.G0 = v13 + int32(32)
						return v189
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v170 = m.ExcPending
						if v170 != 0 {
							return int32(0)
						} else {
							v171 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							v173 = *(*int32)(unsafe.Add(mBase, uint32(v171+v153)))
							v174 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
							v175 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v175
							*(*int64)(unsafe.Add(mBase, uint32(v13)+4)) = v174
							*(*int32)(unsafe.Add(mBase, uint32(v13))) = v173
							F_errmsg_internal(m, int32(_a_F_SetupLockInTable_3), v13)
							mBase = m.M
							v181 = m.ExcPending
							if v181 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_SetupLockInTable_1), int32(1490), int32(_a_F_SetupLockInTable_2))
								mBase = m.M
								v186 = m.ExcPending
								if v186 != 0 {
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v215 int32
	_ = v215
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
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
	v245 = m.ExcPending
	if v245 != 0 {
		goto L15
	} else {
		goto L68
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L15
	} else {
		goto L64
	}
L3:
	;
	m.G0 = v11 - int32(-64)
	return v215
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
	v215 = int32(0)
	goto L3
L7:
	;
	v26 = Fn14265(m, v18, int32(58))
	mBase = m.M
	goto L9
L8:
	;
	goto L6
L9:
	;
	if v26 == v18 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	if v26 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L11:
	;
	v34 = v32 + int32(1)
	v35 = F_palloc(m, v34)
	mBase = m.M
	v38 = m.ExcPending
	if v38 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v30 = F_strlen(m, v18)
	mBase = m.M
	v32 = v30
	goto L11
L13:
	;
	goto L14
L14:
	;
	v32 = v26 - v18
	goto L11
L15:
	;
	return int32(0)
L16:
	;
	if v34 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v159 = F_substitute_path_macro(m, v35, int32(_a_F_find_in_path_0), int32(_a_F_find_in_path_1))
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L15
	} else {
		goto L48
	}
L18:
	;
	v154 = F_strlen(m, v150)
	mBase = m.M
	goto L17
L19:
	;
	v150 = v18
	goto L18
L20:
	;
	goto L21
L21:
	;
	v44 = v34 - int32(1)
	if (v35^v18)&int32(3) != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v147 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v144))) = uint8(v147)
	v150 = v143
	goto L18
L23:
	;
	v128 = v123
	v129 = v124
	v130 = v125
	goto L44
L24:
	;
	if v118 == int32(0) {
		v143 = v116
		v144 = v117
		goto L22
	} else {
		goto L43
	}
L25:
	;
	v116 = v18
	v117 = v35
	v118 = v44
	goto L24
L26:
	;
	goto L27
L27:
	;
	v48 = int32(0)
	if base.B2i32(v18&int32(3) == v48)|base.B2i32(v44 == v48) == v48 {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	if v84 == int32(0) {
		v143 = v81
		v144 = v82
		goto L22
	} else {
		goto L37
	}
L29:
	;
	v60 = v18
	v61 = v35
	v62 = v44
	goto L32
L30:
	;
	goto L31
L31:
	;
	v81 = v18
	v82 = v35
	v83 = v44
	v84 = base.B2i32(v44 != v48)
	goto L28
L32:
	;
	v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
	*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v64)
	if v64 == int32(0) {
		v123 = v60
		v124 = v61
		v125 = v62
		goto L23
	} else {
		goto L34
	}
L33:
	;
	v81 = v75
	v82 = v69
	v83 = v71
	v84 = v73
	goto L28
L34:
	;
	v68 = int32(1)
	v69 = v61 + v68
	v71 = v62 - v68
	v72 = int32(0)
	v73 = base.B2i32(v71 != v72)
	v75 = v60 + v68
	if v75&int32(3) == v72 {
		v81 = v75
		v82 = v69
		v83 = v71
		v84 = v73
		goto L28
	} else {
		goto L35
	}
L35:
	;
	if v71 != 0 {
		v60 = v75
		v61 = v69
		v62 = v71
		goto L32
	} else {
		goto L36
	}
L36:
	;
	goto L33
L37:
	;
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	if base.B2i32(v87 == int32(0))|base.B2i32(base.Ui32(v83) < base.Ui32(int32(4))) != 0 {
		v116 = v81
		v117 = v82
		v118 = v83
		goto L24
	} else {
		goto L38
	}
L38:
	;
	v94 = v81
	v95 = v82
	v96 = v83
	goto L39
L39:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)))
	v102 = int32(-2139062144)
	if (int32(16843008)-v99|v99)&v102 != v102 {
		v123 = v94
		v124 = v95
		v125 = v96
		goto L23
	} else {
		goto L41
	}
L40:
	;
	v116 = v110
	v117 = v108
	v118 = v112
	goto L24
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v95))) = v99
	v107 = int32(4)
	v108 = v95 + v107
	v110 = v94 + v107
	v112 = v96 - v107
	if base.Ui32(int32(3)) < base.Ui32(v112) {
		v94 = v110
		v95 = v108
		v96 = v112
		goto L39
	} else {
		goto L42
	}
L42:
	;
	goto L40
L43:
	;
	v123 = v116
	v124 = v117
	v125 = v118
	goto L23
L44:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128))))
	*(*uint8)(unsafe.Add(mBase, uint32(v129))) = uint8(v132)
	if v132 == int32(0) {
		v143 = v128
		v144 = v129
		goto L22
	} else {
		goto L46
	}
L45:
	;
	v143 = v139
	v144 = v137
	goto L22
L46:
	;
	v136 = int32(1)
	v137 = v129 + v136
	v139 = v128 + v136
	v141 = v130 - v136
	if v141 != 0 {
		v128 = v139
		v129 = v137
		v130 = v141
		goto L44
	} else {
		goto L47
	}
L47:
	;
	goto L45
L48:
	;
	F_pfree(m, v35)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L15
	} else {
		goto L49
	}
L49:
	;
	F_canonicalize_path_enc(m, v159)
	mBase = m.M
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v159))))
	if v164 != int32(47) {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v167 = F_strlen(m, v159)
	mBase = m.M
	v169 = F_palloc(m, v167+(v14+int32(2)))
	mBase = m.M
	v170 = m.ExcPending
	if v170 != 0 {
		goto L15
	} else {
		goto L51
	}
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+36)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = v159
	v176 = F_pg_sprintf(m, v169, int32(_a_F_find_in_path_2), v9+int32(-32))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L15
	} else {
		goto L52
	}
L52:
	;
	F_pfree(m, v159)
	mBase = m.M
	v179 = m.ExcPending
	if v179 != 0 {
		goto L15
	} else {
		goto L53
	}
L53:
	;
	v182 = F_errstart(m, int32(12), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L15
	} else {
		goto L54
	}
L54:
	;
	if v182 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v169
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = int32(_a_F_find_in_path_3)
	F_errmsg_internal(m, int32(_a_F_find_in_path_4), v9+int32(-48))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L15
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v197 = F_pg_file_exists(m, v169)
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L15
	} else {
		goto L60
	}
L58:
	;
	F_errfinish(m, int32(_a_F_find_in_path_5), int32(631), int32(_a_F_find_in_path_3))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L15
	} else {
		goto L59
	}
L59:
	;
	goto L57
L60:
	;
	if v197 != 0 {
		v215 = v169
		goto L3
	} else {
		goto L61
	}
L61:
	;
	F_pfree(m, v169)
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L15
	} else {
		goto L62
	}
L62:
	;
	v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18+v32))))
	if v203 != 0 {
		v18 = v18 + v34
		goto L7
	} else {
		goto L63
	}
L63:
	;
	goto L8
L64:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L15
	} else {
		goto L65
	}
L65:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(_a_F_find_in_path_6)
	F_errmsg(m, int32(_a_F_find_in_path_7), v11)
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L15
	} else {
		goto L66
	}
L66:
	;
	F_errfinish(m, int32(_a_F_find_in_path_5), int32(606), int32(_a_F_find_in_path_3))
	mBase = m.M
	v241 = m.ExcPending
	if v241 != 0 {
		goto L15
	} else {
		goto L67
	}
L67:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L68:
	;
	F_errcode(m, int32(33579140))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L15
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+48)) = int32(_a_F_find_in_path_6)
	F_errmsg(m, int32(_a_F_find_in_path_8), v9+int32(-16))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L15
	} else {
		goto L70
	}
L70:
	;
	F_errfinish(m, int32(_a_F_find_in_path_5), int32(625), int32(_a_F_find_in_path_3))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L15
	} else {
		goto L71
	}
L71:
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
	v90 = v75&int32(63) | (v33<<(uint(int32(18))%32)&int32(_a_F_in_grouping_U_0) | v42<<(uint(int32(12))%32) | v58<<(uint(int32(6))%32))
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
	v90 = v33<<(uint(int32(12))%32)&int32(_a_F_in_grouping_U_1) | v42<<(uint(int32(6))%32) | v58
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
func F_in_range_int8_int8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int64
	_ = v12
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v27 int64
	_ = v27
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	v6 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	if int64(0) <= v6 {
		v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v10 = int64(0)
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
		v14 = base.B2i32(v12 != v10)
		if v12 != v10 {
			v15 = v10 - v6
		} else {
			v15 = v6
		}
		v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v19 = v18 + v15
		if base.B2i32(v15 < int64(0)) != base.B2i32(v19 < v18) {
			return base.I64_extend_i32_u(v14 ^ base.B2i32(v9 != int64(0)))
		} else {
			v27 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
			if v9 != int64(0) {
				return base.I64_extend_i32_u(base.B2i32(v27 <= v19))
			} else {
				return base.I64_extend_i32_u(base.B2i32(v19 <= v27))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_int8_int8_0), int32(0))
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_int8_int8_1), int32(420), int32(_a_F_in_range_int8_int8_2))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int64(0)
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
func F_in_range_numeric_numeric(m *base.Module, l0 int32) int64 {
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
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v49 int32
	_ = v49
	var v52 int64
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v67 int64
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v94 int32
	_ = v94
	var v111 int32
	_ = v111
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v235 int32
	_ = v235
	var v247 int32
	_ = v247
	var v249 int64
	_ = v249
	var v255 int32
	_ = v255
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		return int64(0)
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v26 = F_pg_detoast_datum(m, v25)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int64(0)
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
			v29 = F_pg_detoast_datum(m, v28)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int64(0)
			} else {
				v31 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v29)+4)))
				if base.B2i32(v31 == int32(_a_F_in_range_numeric_numeric_0))|base.B2i32(v31 == int32(_a_F_in_range_numeric_numeric_1)) != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v374 = m.ExcPending
					if v374 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50593922))
						mBase = m.M
						v377 = m.ExcPending
						if v377 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_in_range_numeric_numeric_2), int32(0))
							mBase = m.M
							v381 = m.ExcPending
							if v381 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_in_range_numeric_numeric_3), int32(2604), int32(_a_F_in_range_numeric_numeric_4))
								mBase = m.M
								v386 = m.ExcPending
								if v386 != 0 {
									return int64(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					}
				} else {
					v37 = int32(_a_F_in_range_numeric_numeric_0)
					v38 = v31 & v37
					if v38 != v37 {
						if v38 != int32(_a_F_in_range_numeric_numeric_5) {
							v49 = v38
						} else {
							v49 = v31 << (uint(int32(1)) % 32) & int32(_a_F_in_range_numeric_numeric_6)
						}
					} else {
						v49 = v31 & int32(_a_F_in_range_numeric_numeric_1)
					}
					if v49 == int32(_a_F_in_range_numeric_numeric_6) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v374 = m.ExcPending
						if v374 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50593922))
							mBase = m.M
							v377 = m.ExcPending
							if v377 != 0 {
								return int64(0)
							} else {
								F_errmsg(m, int32(_a_F_in_range_numeric_numeric_2), int32(0))
								mBase = m.M
								v381 = m.ExcPending
								if v381 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_in_range_numeric_numeric_3), int32(2604), int32(_a_F_in_range_numeric_numeric_4))
									mBase = m.M
									v386 = m.ExcPending
									if v386 != 0 {
										return int64(0)
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						}
					} else {
						v52 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
						v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
						v56 = base.I32_extend16_s(v55)
						v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
						if v60 == int32(_a_F_in_range_numeric_numeric_0) {
							v351 = base.B2i32(v52 == int64(0)) | base.B2i32(v56 == int32(-16384))
							v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
							if v352 != v21 {
								F_pfree(m, v21)
								mBase = m.M
								v355 = m.ExcPending
								if v355 != 0 {
									return int64(0)
								} else {
									v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v356 != v26 {
										F_pfree(m, v26)
										mBase = m.M
										v359 = m.ExcPending
										if v359 != 0 {
											return int64(0)
										} else {
											v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v360 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v363 = m.ExcPending
												if v363 != 0 {
													return int64(0)
												} else {
													m.G0 = v18 + int32(96)
													return base.I64_extend_i32_u(v351)
												}
											} else {
												m.G0 = v18 + int32(96)
												return base.I64_extend_i32_u(v351)
											}
										}
									} else {
										v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v360 != v29 {
											F_pfree(m, v29)
											mBase = m.M
											v363 = m.ExcPending
											if v363 != 0 {
												return int64(0)
											} else {
												m.G0 = v18 + int32(96)
												return base.I64_extend_i32_u(v351)
											}
										} else {
											m.G0 = v18 + int32(96)
											return base.I64_extend_i32_u(v351)
										}
									}
								}
							} else {
								v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
								if v356 != v26 {
									F_pfree(m, v26)
									mBase = m.M
									v359 = m.ExcPending
									if v359 != 0 {
										return int64(0)
									} else {
										v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v360 != v29 {
											F_pfree(m, v29)
											mBase = m.M
											v363 = m.ExcPending
											if v363 != 0 {
												return int64(0)
											} else {
												m.G0 = v18 + int32(96)
												return base.I64_extend_i32_u(v351)
											}
										} else {
											m.G0 = v18 + int32(96)
											return base.I64_extend_i32_u(v351)
										}
									}
								} else {
									v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
									if v360 != v29 {
										F_pfree(m, v29)
										mBase = m.M
										v363 = m.ExcPending
										if v363 != 0 {
											return int64(0)
										} else {
											m.G0 = v18 + int32(96)
											return base.I64_extend_i32_u(v351)
										}
									} else {
										m.G0 = v18 + int32(96)
										return base.I64_extend_i32_u(v351)
									}
								}
							}
						} else {
							if v56 == int32(-16384) {
								v351 = base.B2i32(v52 != int64(0))
								v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
								if v352 != v21 {
									F_pfree(m, v21)
									mBase = m.M
									v355 = m.ExcPending
									if v355 != 0 {
										return int64(0)
									} else {
										v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v356 != v26 {
											F_pfree(m, v26)
											mBase = m.M
											v359 = m.ExcPending
											if v359 != 0 {
												return int64(0)
											} else {
												v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v360 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v363 = m.ExcPending
													if v363 != 0 {
														return int64(0)
													} else {
														m.G0 = v18 + int32(96)
														return base.I64_extend_i32_u(v351)
													}
												} else {
													m.G0 = v18 + int32(96)
													return base.I64_extend_i32_u(v351)
												}
											}
										} else {
											v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v360 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v363 = m.ExcPending
												if v363 != 0 {
													return int64(0)
												} else {
													m.G0 = v18 + int32(96)
													return base.I64_extend_i32_u(v351)
												}
											} else {
												m.G0 = v18 + int32(96)
												return base.I64_extend_i32_u(v351)
											}
										}
									}
								} else {
									v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
									if v356 != v26 {
										F_pfree(m, v26)
										mBase = m.M
										v359 = m.ExcPending
										if v359 != 0 {
											return int64(0)
										} else {
											v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v360 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v363 = m.ExcPending
												if v363 != 0 {
													return int64(0)
												} else {
													m.G0 = v18 + int32(96)
													return base.I64_extend_i32_u(v351)
												}
											} else {
												m.G0 = v18 + int32(96)
												return base.I64_extend_i32_u(v351)
											}
										}
									} else {
										v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
										if v360 != v29 {
											F_pfree(m, v29)
											mBase = m.M
											v363 = m.ExcPending
											if v363 != 0 {
												return int64(0)
											} else {
												m.G0 = v18 + int32(96)
												return base.I64_extend_i32_u(v351)
											}
										} else {
											m.G0 = v18 + int32(96)
											return base.I64_extend_i32_u(v351)
										}
									}
								}
							} else {
								v67 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
								v68 = base.I32_extend16_s(v60)
								v69 = base.I32_extend16_s(v31)
								if base.Ui32(int32(-16384)) <= base.Ui32(v69) {
									if v67 != int64(0) {
										if v56 != int32(-12288) {
											v351 = base.B2i32(v52 == int64(0)) | base.B2i32(v68 == int32(-4096))
										} else {
											v351 = int32(1)
										}
									} else {
										if v56 != int32(-4096) {
											v351 = base.B2i32(v68 == int32(-12288)) | base.B2i32(v52 != int64(0))
										} else {
											v351 = int32(1)
										}
									}
									v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
									if v352 != v21 {
										F_pfree(m, v21)
										mBase = m.M
										v355 = m.ExcPending
										if v355 != 0 {
											return int64(0)
										} else {
											v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v356 != v26 {
												F_pfree(m, v26)
												mBase = m.M
												v359 = m.ExcPending
												if v359 != 0 {
													return int64(0)
												} else {
													v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if v360 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v363 = m.ExcPending
														if v363 != 0 {
															return int64(0)
														} else {
															m.G0 = v18 + int32(96)
															return base.I64_extend_i32_u(v351)
														}
													} else {
														m.G0 = v18 + int32(96)
														return base.I64_extend_i32_u(v351)
													}
												}
											} else {
												v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v360 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v363 = m.ExcPending
													if v363 != 0 {
														return int64(0)
													} else {
														m.G0 = v18 + int32(96)
														return base.I64_extend_i32_u(v351)
													}
												} else {
													m.G0 = v18 + int32(96)
													return base.I64_extend_i32_u(v351)
												}
											}
										}
									} else {
										v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
										if v356 != v26 {
											F_pfree(m, v26)
											mBase = m.M
											v359 = m.ExcPending
											if v359 != 0 {
												return int64(0)
											} else {
												v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v360 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v363 = m.ExcPending
													if v363 != 0 {
														return int64(0)
													} else {
														m.G0 = v18 + int32(96)
														return base.I64_extend_i32_u(v351)
													}
												} else {
													m.G0 = v18 + int32(96)
													return base.I64_extend_i32_u(v351)
												}
											}
										} else {
											v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
											if v360 != v29 {
												F_pfree(m, v29)
												mBase = m.M
												v363 = m.ExcPending
												if v363 != 0 {
													return int64(0)
												} else {
													m.G0 = v18 + int32(96)
													return base.I64_extend_i32_u(v351)
												}
											} else {
												m.G0 = v18 + int32(96)
												return base.I64_extend_i32_u(v351)
											}
										}
									}
								} else {
									if base.Ui32(int32(-16384)) <= base.Ui32(v68) {
										v94 = int32(-12288)
										if v68 == v94 {
											v351 = base.B2i32(v52 == int64(0)) | base.B2i32(v56 == v94)
										} else {
											v351 = base.B2i32(v56 == int32(-4096)) | base.B2i32(v52 != int64(0))
										}
										v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v352 != v21 {
											F_pfree(m, v21)
											mBase = m.M
											v355 = m.ExcPending
											if v355 != 0 {
												return int64(0)
											} else {
												v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												if v356 != v26 {
													F_pfree(m, v26)
													mBase = m.M
													v359 = m.ExcPending
													if v359 != 0 {
														return int64(0)
													} else {
														v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
														if v360 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v363 = m.ExcPending
															if v363 != 0 {
																return int64(0)
															} else {
																m.G0 = v18 + int32(96)
																return base.I64_extend_i32_u(v351)
															}
														} else {
															m.G0 = v18 + int32(96)
															return base.I64_extend_i32_u(v351)
														}
													}
												} else {
													v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if v360 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v363 = m.ExcPending
														if v363 != 0 {
															return int64(0)
														} else {
															m.G0 = v18 + int32(96)
															return base.I64_extend_i32_u(v351)
														}
													} else {
														m.G0 = v18 + int32(96)
														return base.I64_extend_i32_u(v351)
													}
												}
											}
										} else {
											v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
											if v356 != v26 {
												F_pfree(m, v26)
												mBase = m.M
												v359 = m.ExcPending
												if v359 != 0 {
													return int64(0)
												} else {
													v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if v360 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v363 = m.ExcPending
														if v363 != 0 {
															return int64(0)
														} else {
															m.G0 = v18 + int32(96)
															return base.I64_extend_i32_u(v351)
														}
													} else {
														m.G0 = v18 + int32(96)
														return base.I64_extend_i32_u(v351)
													}
												}
											} else {
												v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
												if v360 != v29 {
													F_pfree(m, v29)
													mBase = m.M
													v363 = m.ExcPending
													if v363 != 0 {
														return int64(0)
													} else {
														m.G0 = v18 + int32(96)
														return base.I64_extend_i32_u(v351)
													}
												} else {
													m.G0 = v18 + int32(96)
													return base.I64_extend_i32_u(v351)
												}
											}
										}
									} else {
										if base.Ui32(int32(-16384)) <= base.Ui32(v56) {
											v351 = base.B2i32(v56 == int32(-4096)) ^ base.B2i32(v52 != int64(0))
											v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
											if v352 != v21 {
												F_pfree(m, v21)
												mBase = m.M
												v355 = m.ExcPending
												if v355 != 0 {
													return int64(0)
												} else {
													v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
													if v356 != v26 {
														F_pfree(m, v26)
														mBase = m.M
														v359 = m.ExcPending
														if v359 != 0 {
															return int64(0)
														} else {
															v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
															if v360 != v29 {
																F_pfree(m, v29)
																mBase = m.M
																v363 = m.ExcPending
																if v363 != 0 {
																	return int64(0)
																} else {
																	m.G0 = v18 + int32(96)
																	return base.I64_extend_i32_u(v351)
																}
															} else {
																m.G0 = v18 + int32(96)
																return base.I64_extend_i32_u(v351)
															}
														}
													} else {
														v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
														if v360 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v363 = m.ExcPending
															if v363 != 0 {
																return int64(0)
															} else {
																m.G0 = v18 + int32(96)
																return base.I64_extend_i32_u(v351)
															}
														} else {
															m.G0 = v18 + int32(96)
															return base.I64_extend_i32_u(v351)
														}
													}
												}
											} else {
												v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
												if v356 != v26 {
													F_pfree(m, v26)
													mBase = m.M
													v359 = m.ExcPending
													if v359 != 0 {
														return int64(0)
													} else {
														v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
														if v360 != v29 {
															F_pfree(m, v29)
															mBase = m.M
															v363 = m.ExcPending
															if v363 != 0 {
																return int64(0)
															} else {
																m.G0 = v18 + int32(96)
																return base.I64_extend_i32_u(v351)
															}
														} else {
															m.G0 = v18 + int32(96)
															return base.I64_extend_i32_u(v351)
														}
													}
												} else {
													v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
													if v360 != v29 {
														F_pfree(m, v29)
														mBase = m.M
														v363 = m.ExcPending
														if v363 != 0 {
															return int64(0)
														} else {
															m.G0 = v18 + int32(96)
															return base.I64_extend_i32_u(v351)
														}
													} else {
														m.G0 = v18 + int32(96)
														return base.I64_extend_i32_u(v351)
													}
												}
											}
										} else {
											v111 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
											v117 = base.B2i32(int32(0) <= v68)
											if int32(0) <= v68 {
												v118 = int32(-8)
											} else {
												v118 = int32(-6)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+72)) = int32(base.Ui32(int32(base.Ui32(v111)>>(uint(int32(2))%32))+v118) >> (uint(int32(1)) % 32))
											if int32(0) <= v68 {
												v123 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
												v133 = v123
											} else {
												v133 = v60<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v60&int32(63)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+76)) = v133
											v135 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v135
											v144 = base.B2i32(v68 < v135)
											if v68 < v135 {
												v145 = int32(base.Ui32(v60)>>(uint(int32(7))%32)) & int32(63)
											} else {
												v145 = v60 & int32(_a_F_in_range_numeric_numeric_7)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v145
											v152 = v60 & int32(_a_F_in_range_numeric_numeric_0)
											if v152 == int32(_a_F_in_range_numeric_numeric_5) {
												v155 = v60 << (uint(int32(1)) % 32) & int32(_a_F_in_range_numeric_numeric_6)
											} else {
												v155 = v152
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+80)) = v155
											if v68 < v135 {
												v159 = int32(6)
											} else {
												v159 = int32(8)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v21 + v159
											v162 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
											v168 = base.B2i32(int32(0) <= v56)
											if int32(0) <= v56 {
												v169 = int32(-8)
											} else {
												v169 = int32(-6)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = int32(base.Ui32(int32(base.Ui32(v162)>>(uint(int32(2))%32))+v169) >> (uint(int32(1)) % 32))
											if int32(0) <= v56 {
												v174 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+6)))
												v184 = v174
											} else {
												v184 = v55<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v55&int32(63)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v184
											v186 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v186
											v195 = base.B2i32(v56 < v186)
											if v56 < v186 {
												v196 = int32(base.Ui32(v55)>>(uint(int32(7))%32)) & int32(63)
											} else {
												v196 = v55 & int32(_a_F_in_range_numeric_numeric_7)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v196
											v203 = v55 & int32(_a_F_in_range_numeric_numeric_0)
											if v203 == int32(_a_F_in_range_numeric_numeric_5) {
												v206 = v55 << (uint(int32(1)) % 32) & int32(_a_F_in_range_numeric_numeric_6)
											} else {
												v206 = v203
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v206
											if v56 < v186 {
												v210 = int32(6)
											} else {
												v210 = int32(8)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v26 + v210
											v213 = *(*int32)(unsafe.Add(mBase, uint32(v29)))
											v219 = base.B2i32(int32(0) <= v69)
											if int32(0) <= v69 {
												v220 = int32(-8)
											} else {
												v220 = int32(-6)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = int32(base.Ui32(int32(base.Ui32(v213)>>(uint(int32(2))%32))+v220) >> (uint(int32(1)) % 32))
											if int32(0) <= v69 {
												v225 = int32(*(*int16)(unsafe.Add(mBase, uint32(v29)+6)))
												v235 = v225
											} else {
												v235 = v31<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v31&int32(63)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v235
											if v38 != int32(_a_F_in_range_numeric_numeric_0) {
												if v38 != int32(_a_F_in_range_numeric_numeric_5) {
													v247 = v38
												} else {
													v247 = v31 << (uint(int32(1)) % 32) & int32(_a_F_in_range_numeric_numeric_6)
												}
											} else {
												v247 = v31 & int32(_a_F_in_range_numeric_numeric_1)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v247
											v249 = int64(0)
											*(*int64)(unsafe.Add(mBase, uint32(v18))) = v249
											*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v249
											*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v249
											v255 = int32(0)
											*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v255
											v264 = base.B2i32(v69 < v255)
											if v69 < v255 {
												v265 = int32(base.Ui32(v31)>>(uint(int32(7))%32)) & int32(63)
											} else {
												v265 = v31 & int32(_a_F_in_range_numeric_numeric_7)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v265
											if v69 < v255 {
												v269 = int32(6)
											} else {
												v269 = int32(8)
											}
											*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v29 + v269
											if v67 != int64(0) {
												F_sub_var(m, v18+int32(48), v18+int32(24), v18)
												mBase = m.M
												v279 = m.ExcPending
												if v279 != 0 {
													return int64(0)
												} else {
													v289 = v18 + int32(72)
													v296 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
													v298 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
													if v298 == int32(0) {
														if v297 == int32(0) {
															v334 = int32(0)
														} else {
															if v296 == int32(_a_F_in_range_numeric_numeric_6) {
																v308 = int32(1)
															} else {
																v308 = int32(-1)
															}
															v334 = v308
														}
													} else {
														v309 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
														if v297 == int32(0) {
															if v309 != 0 {
																v314 = int32(-1)
															} else {
																v314 = int32(1)
															}
															v334 = v314
														} else {
															v315 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
															v316 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
															v317 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
															v318 = *(*int32)(unsafe.Add(mBase, uint32(v289)+20))
															if v309 == int32(0) {
																if v296 == int32(_a_F_in_range_numeric_numeric_6) {
																	v334 = int32(1)
																} else {
																	v324 = F_cmp_abs_common(m, v318, v298, v317, v316, v297, v315)
																	mBase = m.M
																	v334 = v324
																}
															} else {
																if v296 == int32(0) {
																	v334 = int32(-1)
																} else {
																	v328 = F_cmp_abs_common(m, v316, v297, v315, v318, v298, v317)
																	mBase = m.M
																	v334 = v328
																}
															}
														}
													}
													v337 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
													if v337 != 0 {
														F_pfree(m, v337)
														mBase = m.M
														v339 = m.ExcPending
														if v339 != 0 {
															return int64(0)
														} else {
															if v52 != int64(0) {
																v342 = base.B2i32(v334 <= int32(0))
															} else {
																v342 = base.B2i32(int32(0) <= v334)
															}
															v351 = v342
															v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
															if v352 != v21 {
																F_pfree(m, v21)
																mBase = m.M
																v355 = m.ExcPending
																if v355 != 0 {
																	return int64(0)
																} else {
																	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																	if v356 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v359 = m.ExcPending
																		if v359 != 0 {
																			return int64(0)
																		} else {
																			v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																			if v360 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v363 = m.ExcPending
																				if v363 != 0 {
																					return int64(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return base.I64_extend_i32_u(v351)
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return base.I64_extend_i32_u(v351)
																			}
																		}
																	} else {
																		v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																		if v360 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v363 = m.ExcPending
																			if v363 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return base.I64_extend_i32_u(v351)
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	}
																}
															} else {
																v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																if v356 != v26 {
																	F_pfree(m, v26)
																	mBase = m.M
																	v359 = m.ExcPending
																	if v359 != 0 {
																		return int64(0)
																	} else {
																		v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																		if v360 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v363 = m.ExcPending
																			if v363 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return base.I64_extend_i32_u(v351)
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	}
																} else {
																	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																	if v360 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v363 = m.ExcPending
																		if v363 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return base.I64_extend_i32_u(v351)
																	}
																}
															}
														}
													} else {
														if v52 != int64(0) {
															v342 = base.B2i32(v334 <= int32(0))
														} else {
															v342 = base.B2i32(int32(0) <= v334)
														}
														v351 = v342
														v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														if v352 != v21 {
															F_pfree(m, v21)
															mBase = m.M
															v355 = m.ExcPending
															if v355 != 0 {
																return int64(0)
															} else {
																v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																if v356 != v26 {
																	F_pfree(m, v26)
																	mBase = m.M
																	v359 = m.ExcPending
																	if v359 != 0 {
																		return int64(0)
																	} else {
																		v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																		if v360 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v363 = m.ExcPending
																			if v363 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return base.I64_extend_i32_u(v351)
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	}
																} else {
																	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																	if v360 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v363 = m.ExcPending
																		if v363 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return base.I64_extend_i32_u(v351)
																	}
																}
															}
														} else {
															v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															if v356 != v26 {
																F_pfree(m, v26)
																mBase = m.M
																v359 = m.ExcPending
																if v359 != 0 {
																	return int64(0)
																} else {
																	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																	if v360 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v363 = m.ExcPending
																		if v363 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return base.I64_extend_i32_u(v351)
																	}
																}
															} else {
																v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																if v360 != v29 {
																	F_pfree(m, v29)
																	mBase = m.M
																	v363 = m.ExcPending
																	if v363 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v18 + int32(96)
																		return base.I64_extend_i32_u(v351)
																	}
																} else {
																	m.G0 = v18 + int32(96)
																	return base.I64_extend_i32_u(v351)
																}
															}
														}
													}
												}
											} else {
												F_add_var(m, v18+int32(48), v18+int32(24), v18)
												mBase = m.M
												v285 = m.ExcPending
												if v285 != 0 {
													return int64(0)
												} else {
													v289 = v18 + int32(72)
													v296 = *(*int32)(unsafe.Add(mBase, uint32(v18)+8))
													v297 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
													v298 = *(*int32)(unsafe.Add(mBase, uint32(v289)))
													if v298 == int32(0) {
														if v297 == int32(0) {
															v334 = int32(0)
														} else {
															if v296 == int32(_a_F_in_range_numeric_numeric_6) {
																v308 = int32(1)
															} else {
																v308 = int32(-1)
															}
															v334 = v308
														}
													} else {
														v309 = *(*int32)(unsafe.Add(mBase, uint32(v289)+8))
														if v297 == int32(0) {
															if v309 != 0 {
																v314 = int32(-1)
															} else {
																v314 = int32(1)
															}
															v334 = v314
														} else {
															v315 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
															v316 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
															v317 = *(*int32)(unsafe.Add(mBase, uint32(v289)+4))
															v318 = *(*int32)(unsafe.Add(mBase, uint32(v289)+20))
															if v309 == int32(0) {
																if v296 == int32(_a_F_in_range_numeric_numeric_6) {
																	v334 = int32(1)
																} else {
																	v324 = F_cmp_abs_common(m, v318, v298, v317, v316, v297, v315)
																	mBase = m.M
																	v334 = v324
																}
															} else {
																if v296 == int32(0) {
																	v334 = int32(-1)
																} else {
																	v328 = F_cmp_abs_common(m, v316, v297, v315, v318, v298, v317)
																	mBase = m.M
																	v334 = v328
																}
															}
														}
													}
													v337 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
													if v337 != 0 {
														F_pfree(m, v337)
														mBase = m.M
														v339 = m.ExcPending
														if v339 != 0 {
															return int64(0)
														} else {
															if v52 != int64(0) {
																v342 = base.B2i32(v334 <= int32(0))
															} else {
																v342 = base.B2i32(int32(0) <= v334)
															}
															v351 = v342
															v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
															if v352 != v21 {
																F_pfree(m, v21)
																mBase = m.M
																v355 = m.ExcPending
																if v355 != 0 {
																	return int64(0)
																} else {
																	v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																	if v356 != v26 {
																		F_pfree(m, v26)
																		mBase = m.M
																		v359 = m.ExcPending
																		if v359 != 0 {
																			return int64(0)
																		} else {
																			v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																			if v360 != v29 {
																				F_pfree(m, v29)
																				mBase = m.M
																				v363 = m.ExcPending
																				if v363 != 0 {
																					return int64(0)
																				} else {
																					m.G0 = v18 + int32(96)
																					return base.I64_extend_i32_u(v351)
																				}
																			} else {
																				m.G0 = v18 + int32(96)
																				return base.I64_extend_i32_u(v351)
																			}
																		}
																	} else {
																		v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																		if v360 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v363 = m.ExcPending
																			if v363 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return base.I64_extend_i32_u(v351)
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	}
																}
															} else {
																v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																if v356 != v26 {
																	F_pfree(m, v26)
																	mBase = m.M
																	v359 = m.ExcPending
																	if v359 != 0 {
																		return int64(0)
																	} else {
																		v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																		if v360 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v363 = m.ExcPending
																			if v363 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return base.I64_extend_i32_u(v351)
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	}
																} else {
																	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																	if v360 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v363 = m.ExcPending
																		if v363 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return base.I64_extend_i32_u(v351)
																	}
																}
															}
														}
													} else {
														if v52 != int64(0) {
															v342 = base.B2i32(v334 <= int32(0))
														} else {
															v342 = base.B2i32(int32(0) <= v334)
														}
														v351 = v342
														v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
														if v352 != v21 {
															F_pfree(m, v21)
															mBase = m.M
															v355 = m.ExcPending
															if v355 != 0 {
																return int64(0)
															} else {
																v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
																if v356 != v26 {
																	F_pfree(m, v26)
																	mBase = m.M
																	v359 = m.ExcPending
																	if v359 != 0 {
																		return int64(0)
																	} else {
																		v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																		if v360 != v29 {
																			F_pfree(m, v29)
																			mBase = m.M
																			v363 = m.ExcPending
																			if v363 != 0 {
																				return int64(0)
																			} else {
																				m.G0 = v18 + int32(96)
																				return base.I64_extend_i32_u(v351)
																			}
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	}
																} else {
																	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																	if v360 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v363 = m.ExcPending
																		if v363 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return base.I64_extend_i32_u(v351)
																	}
																}
															}
														} else {
															v356 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
															if v356 != v26 {
																F_pfree(m, v26)
																mBase = m.M
																v359 = m.ExcPending
																if v359 != 0 {
																	return int64(0)
																} else {
																	v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																	if v360 != v29 {
																		F_pfree(m, v29)
																		mBase = m.M
																		v363 = m.ExcPending
																		if v363 != 0 {
																			return int64(0)
																		} else {
																			m.G0 = v18 + int32(96)
																			return base.I64_extend_i32_u(v351)
																		}
																	} else {
																		m.G0 = v18 + int32(96)
																		return base.I64_extend_i32_u(v351)
																	}
																}
															} else {
																v360 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
																if v360 != v29 {
																	F_pfree(m, v29)
																	mBase = m.M
																	v363 = m.ExcPending
																	if v363 != 0 {
																		return int64(0)
																	} else {
																		m.G0 = v18 + int32(96)
																		return base.I64_extend_i32_u(v351)
																	}
																} else {
																	m.G0 = v18 + int32(96)
																	return base.I64_extend_i32_u(v351)
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
func F_in_range_time_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int64
	_ = v7
	var v10 int64
	_ = v10
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v19 int64
	_ = v19
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v7 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	if int64(0) <= v7 {
		v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
		if v13 != int64(0) {
			v46 = v11 - v7
			if v10 != int64(0) {
				return base.I64_extend_i32_u(base.B2i32(v12 <= v46))
			} else {
				return base.I64_extend_i32_u(base.B2i32(v46 <= v12))
			}
		} else {
			v19 = v7 + v11
			if base.B2i32(v7 < int64(0))^base.B2i32(v19 < v11) == int32(0) {
				v46 = v19
				if v10 != int64(0) {
					return base.I64_extend_i32_u(base.B2i32(v12 <= v46))
				} else {
					return base.I64_extend_i32_u(base.B2i32(v46 <= v12))
				}
			} else {
				return base.I64_extend_i32_u(base.B2i32(v10 != int64(0)))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v33 = m.ExcPending
		if v33 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_time_interval_0), int32(0))
				mBase = m.M
				v40 = m.ExcPending
				if v40 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_time_interval_1), int32(2232), int32(_a_F_in_range_time_interval_2))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
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
func F_in_range_timetz_interval(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v21 int64
	_ = v21
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v51 int64
	_ = v51
	var v53 int64
	_ = v53
	var v54 int64
	_ = v54
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(v7)))
	if int64(0) <= v8 {
		v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v14 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(l0)+72))
		if v15 != int64(0) {
			v48 = v14 - v8
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
			v51 = int64(1000000)
			v53 = base.I64_extend_i32_s(v49)*v51 + v48
			v54 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
			v59 = v54 + base.I64_extend_i32_s(v55)*v51
			if v11 != int64(0) {
				if v53 < v59 {
					return int64(0)
				} else {
					return base.I64_extend_i32_u(base.B2i32(v55 <= v49) | base.B2i32(v59 < v53))
				}
			} else {
				if v53 < v59 {
					return int64(1)
				} else {
					return base.I64_extend_i32_u(base.B2i32(v49 <= v55) & base.B2i32(v53 <= v59))
				}
			}
		} else {
			v21 = v8 + v14
			if base.B2i32(v8 < int64(0))^base.B2i32(v21 < v14) == int32(0) {
				v48 = v21
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
				v51 = int64(1000000)
				v53 = base.I64_extend_i32_s(v49)*v51 + v48
				v54 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
				v59 = v54 + base.I64_extend_i32_s(v55)*v51
				if v11 != int64(0) {
					if v53 < v59 {
						return int64(0)
					} else {
						return base.I64_extend_i32_u(base.B2i32(v55 <= v49) | base.B2i32(v59 < v53))
					}
				} else {
					if v53 < v59 {
						return int64(1)
					} else {
						return base.I64_extend_i32_u(base.B2i32(v49 <= v55) & base.B2i32(v53 <= v59))
					}
				}
			} else {
				return base.I64_extend_i32_u(base.B2i32(v11 != int64(0)))
			}
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int64(0)
		} else {
			F_errcode(m, int32(50593922))
			mBase = m.M
			v38 = m.ExcPending
			if v38 != 0 {
				return int64(0)
			} else {
				F_errmsg(m, int32(_a_F_in_range_timetz_interval_0), int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_in_range_timetz_interval_1), int32(2790), int32(_a_F_in_range_timetz_interval_2))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int64(0)
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
