package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_LockViewRecurse(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
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
	var v45 int32
	_ = v45
	v3 = l2
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = F_table_open(m, l0, int32(0))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = F_get_view_query(m, v13)
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return
		} else {
			*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)) = uint8(v3)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = l1
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+180))
			if v19 == int32(0) {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
				v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
				v29 = v28
			} else {
				v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+5)))
				if v22 == int32(0) {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)+48))
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+80))
					v29 = v28
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, _c_F_LockViewRecurse[0]))
					v29 = v26
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = l0
			*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v29
			v32 = F_lappend_oid(m, l3, l0)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v32
				v37 = F_LockViewRecurse_walker(m, v15, v10+int32(12))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v10)+28))
					v40 = F_list_delete_last(m, v39)
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v40
						F_relation_close(m, v13, int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return
						} else {
							m.G0 = v10 + int32(32)
							return
						}
					}
				}
			}
		}
	}
}
func F_LookupOperName(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = F_OpernameGetOprid(m, l0, l1, l2)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if l3|v10 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				if l2 == int32(0) {
					F_errcode(m, int32(16801924))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_LookupOperName_0), int32(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_parser_errposition(m, int32(0), int32(-1))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_LookupOperName_1), int32(115), int32(_a_F_LookupOperName_2))
								mBase = m.M
								v60 = m.ExcPending
								if v60 != 0 {
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
					F_errcode(m, int32(52461700))
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_op_signature_string(m, l0, l1, l2)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = v26
							F_errmsg(m, int32(_a_F_LookupOperName_3), v8)
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								F_parser_errposition(m, int32(0), int32(-1))
								mBase = m.M
								v35 = m.ExcPending
								if v35 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_LookupOperName_1), int32(121), int32(_a_F_LookupOperName_2))
									mBase = m.M
									v40 = m.ExcPending
									if v40 != 0 {
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
		} else {
			m.G0 = v8 + int32(16)
			return v10
		}
	}
}
func F___lshrti3(m *base.Module, l0 int32, l1 int64, l2 int64, l3 int32) {
	mBase := m.M
	_ = mBase
	var v19 int64
	_ = v19
	var v23 int64
	_ = v23
	var v24 int64
	_ = v24
	if l3&int32(64) != 0 {
		v23 = int64(base.Ui64(l2) >> (uint(base.I64_extend_i32_u(l3+int32(-64))) % 64))
		v24 = int64(0)
	} else {
		if l3 == int32(0) {
			v23 = l1
			v24 = l2
		} else {
			v19 = base.I64_extend_i32_u(l3)
			v23 = l2<<(uint(base.I64_extend_i32_u(int32(64)-l3))%64) | int64(base.Ui64(l1)>>(uint(v19)%64))
			v24 = int64(base.Ui64(l2) >> (uint(v19) % 64))
		}
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v23
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v24
	return
}
func F_latch_sigurg_handler(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	v2 = m.G0
	v4 = v2 - int32(16)
	m.G0 = v4
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_latch_sigurg_handler[0]))
	if v7 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v4 + int32(16)
	return
L2:
	;
	v10 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v4)+15)) = uint8(v10)
	goto L3
L3:
	;
	v14 = *(*int32)(unsafe.Add(mBase, _c_F_latch_sigurg_handler[1]))
	v18 = F_write(m, v14, v4+int32(15), int32(1))
	mBase = m.M
	if int32(0) <= v18 {
		goto L1
	} else {
		goto L5
	}
L4:
	;
	goto L1
L5:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_latch_sigurg_handler[2]))
	if v22 == int32(27) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	goto L4
}
func F_lca_inner(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v153 int32
	_ = v153
	var v162 int32
	_ = v162
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v316 int32
	_ = v316
	var v333 int32
	_ = v333
	v3 = int32(0)
	if l1 <= v3 {
		v333 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v333
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	if v15 == int32(0) {
		v333 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = int32(1)
	v19 = v15 - v18
	if l1 != v18 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v30 = v19
	v33 = l0 + int32(4)
	goto L7
L5:
	;
	v162 = v19
	goto L6
L6:
	;
	if v162 <= int32(0) {
		goto L40
	} else {
		goto L41
	}
L7:
	;
	v37 = int32(0)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v39 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	switch v39 {
	case 0:
		v333 = v37
		goto L1
	case 1:
		v145 = v37
		goto L9
	default:
		goto L10
	}
L8:
	;
	v162 = v145
	goto L6
L9:
	;
	v153 = v33 + int32(4)
	if (v153-l0)>>(uint(int32(2))%32) < l1 {
		v30 = v145
		v33 = v153
		goto L7
	} else {
		goto L38
	}
L10:
	;
	v41 = v39 - int32(1)
	if v30 < v41 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v43 = v30
	goto L13
L12:
	;
	v43 = v41
	goto L13
L13:
	;
	if v43 <= int32(0) {
		v145 = v37
		goto L9
	} else {
		goto L14
	}
L14:
	;
	v50 = v14 + int32(8)
	v51 = v38 + int32(8)
	v52 = v37
	goto L15
L15:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51))))
	if v59 != v60 {
		v145 = v52
		goto L9
	} else {
		goto L17
	}
L16:
	;
	v145 = v43
	goto L9
L17:
	;
	v62 = int32(2)
	v63 = v50 + v62
	v65 = v51 + v62
	if base.Ui32(int32(4)) <= base.Ui32(v59) {
		goto L21
	} else {
		goto L22
	}
L18:
	;
	if v127 != 0 {
		v145 = v52
		goto L9
	} else {
		goto L36
	}
L19:
	;
	v127 = int32(0)
	goto L18
L20:
	;
	v101 = v96
	v102 = v97
	v103 = v98
	goto L30
L21:
	;
	if (v63|v65)&int32(3) != 0 {
		v96 = v63
		v97 = v65
		v98 = v59
		goto L20
	} else {
		goto L24
	}
L22:
	;
	v89 = v63
	v90 = v65
	v91 = v59
	goto L23
L23:
	;
	if v91 == int32(0) {
		goto L19
	} else {
		goto L29
	}
L24:
	;
	v73 = v63
	v74 = v65
	v75 = v59
	goto L25
L25:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	if v78 != v79 {
		v96 = v73
		v97 = v74
		v98 = v75
		goto L20
	} else {
		goto L27
	}
L26:
	;
	v89 = v84
	v90 = v82
	v91 = v86
	goto L23
L27:
	;
	v81 = int32(4)
	v82 = v74 + v81
	v84 = v73 + v81
	v86 = v75 - v81
	if base.Ui32(int32(3)) < base.Ui32(v86) {
		v73 = v84
		v74 = v82
		v75 = v86
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v96 = v89
	v97 = v90
	v98 = v91
	goto L20
L30:
	;
	v106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102))))
	if v106 == v107 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v127 = v106 - v107
	goto L18
L32:
	;
	v109 = int32(1)
	v114 = v103 - v109
	if v114 != 0 {
		v101 = v101 + v109
		v102 = v102 + v109
		v103 = v114
		goto L30
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	goto L19
L36:
	;
	v128 = int32(9)
	v130 = int32(_a_F_lca_inner_0)
	v139 = v52 + int32(1)
	if v139 != v43 {
		v50 = v50 + (v59+v128)&v130
		v51 = v51 + (v60+v128)&v130
		v52 = v139
		goto L15
	} else {
		goto L37
	}
L37:
	;
	goto L16
L38:
	;
	goto L8
L39:
	;
	v270 = F_palloc0(m, v261)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L53
	} else {
		goto L54
	}
L40:
	;
	v261 = int32(8)
	goto L39
L41:
	;
	goto L42
L42:
	;
	v173 = v162 & int32(3)
	v174 = int32(8)
	v176 = v14 + v174
	if base.Ui32(int32(4)) <= base.Ui32(v162) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v184 = v174
	v185 = v176
	v187 = int32(0)
	goto L46
L44:
	;
	v228 = v174
	v229 = v176
	goto L45
L45:
	;
	v240 = v228
	v241 = v229
	v243 = int32(0)
	goto L50
L46:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185))))
	v194 = int32(9)
	v196 = int32(_a_F_lca_inner_0)
	v197 = (v193 + v194) & v196
	v198 = v185 + v197
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198))))
	v203 = (v199 + v194) & v196
	v204 = v198 + v203
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204))))
	v209 = (v205 + v194) & v196
	v210 = v204 + v209
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v210))))
	v215 = (v211 + v194) & v196
	v216 = v210 + v215
	v220 = v215 + (v209 + (v203 + (v184 + v197)))
	v222 = v187 + int32(4)
	if v222 != v162&int32(2147483644) {
		v184 = v220
		v185 = v216
		v187 = v222
		goto L46
	} else {
		goto L48
	}
L47:
	;
	if v173 == int32(0) {
		v261 = v220
		goto L39
	} else {
		goto L49
	}
L48:
	;
	goto L47
L49:
	;
	v228 = v220
	v229 = v216
	goto L45
L50:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v241))))
	v253 = (v249 + int32(9)) & int32(_a_F_lca_inner_0)
	v255 = v253 + v240
	v257 = v243 + int32(1)
	if v257 != v173 {
		v240 = v255
		v241 = v241 + v253
		v243 = v257
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v261 = v255
	goto L39
L52:
	;
	goto L51
L53:
	;
	return int32(0)
L54:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v270)+4)) = uint16(v162)
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v261 << (uint(int32(2)) % 32)
	if int32(0) < v162 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v280 = int32(8)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v288 = v270 + v280
	v289 = v282 + v280
	v291 = int32(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	v333 = v270
	goto L1
L58:
	;
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289))))
	v301 = (v297 + int32(9)) & int32(_a_F_lca_inner_0)
	if v301 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	goto L57
L60:
	;
	base.MemoryCopy(m, v288, v289, v301)
	goto L62
L61:
	;
	goto L62
L62:
	;
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288))))
	v304 = int32(9)
	v306 = int32(_a_F_lca_inner_0)
	v309 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289))))
	v316 = v291 + int32(1)
	if v316 != v162 {
		v288 = v288 + (v303+v304)&v306
		v289 = v289 + (v309+v304)&v306
		v291 = v316
		goto L58
	} else {
		goto L63
	}
L63:
	;
	goto L59
}
func F_lcons(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13940(m, l0, l1, int64(4294967297))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_lcons_int(m *base.Module, l0 int32, l1 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13940(m, l0, l1, int64(4294967767))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_leftmostvalue_bit(m *base.Module) int32 {
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v2 = int32(0)
	v6 = F_DirectFunctionCall3Coll(m, int32(490), v2, int32(_a_F_leftmostvalue_bit_0), v2, int32(-1))
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		return v6
	}
}
func F_leftmostvalue_interval(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_palloc(m, int32(16))
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(v3)+8)) = int64(-9223372034707292160)
		*(*int64)(unsafe.Add(mBase, uint32(v3))) = int64(-9223372036854775807 - 1)
		return v3
	}
}
func F_like_escape(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
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
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
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
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v329 int32
	_ = v329
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v412 int32
	_ = v412
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v444 int32
	_ = v444
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v516 int32
	_ = v516
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = *(*int32)(unsafe.Add(mBase, _c_F_like_escape[0]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23*int32(28))+uint32(_c_F_like_escape[1])))
	goto L4
L4:
	;
	if v28 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v31 = F_SB_do_like_escape(m, v14, v19)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v34 = int32(1)
	v35 = v14 + v34
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	v38 = v36 & v34
	if v36 == v34 {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	return v31
L9:
	;
	v68 = int32(1)
	v69 = v19 + v68
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v72 = v70 & v68
	if v70 == v68 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v44 == int32(18) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v55 = int32(1)
	if v38 != 0 {
		v65 = int32(base.Ui32(v36)>>(uint(v55)%32)) - v55
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v47 = int32(16)
	goto L15
L14:
	;
	v47 = int32(0)
	goto L15
L15:
	;
	if base.Ui32((v44-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v54 = int32(4)
	goto L18
L17:
	;
	v54 = v47
	goto L18
L18:
	;
	v65 = v54
	goto L9
L19:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v65 = int32(base.Ui32(v59)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L20:
	;
	if v38 != 0 {
		goto L31
	} else {
		goto L32
	}
L21:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v78 == int32(18) {
		goto L24
	} else {
		goto L25
	}
L22:
	;
	goto L23
L23:
	;
	v89 = int32(1)
	if v72 != 0 {
		v99 = int32(base.Ui32(v70)>>(uint(v89)%32)) - v89
		goto L20
	} else {
		goto L30
	}
L24:
	;
	v81 = int32(16)
	goto L26
L25:
	;
	v81 = int32(0)
	goto L26
L26:
	;
	if base.Ui32((v78-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v88 = int32(4)
	goto L29
L28:
	;
	v88 = v81
	goto L29
L29:
	;
	v99 = v88
	goto L20
L30:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v99 = int32(base.Ui32(v93)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L31:
	;
	v100 = v35
	goto L33
L32:
	;
	v100 = v14 + int32(4)
	goto L33
L33:
	;
	v105 = F_palloc(m, v65<<(uint(int32(1))%32)+int32(4))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	v108 = v105 + int32(4)
	if v99 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L35:
	;
	return v105
L36:
	;
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v507 == int32(1) {
		goto L128
	} else {
		goto L129
	}
L37:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L122
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = (v472 - v105) << (uint(int32(2)) % 32)
	goto L35
L39:
	;
	if v65 <= int32(0) {
		v472 = v108
		goto L38
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	v229 = v19 + int32(4)
	if v72 != 0 {
		goto L62
	} else {
		goto L63
	}
L42:
	;
	v113 = v100
	v114 = v108
	v119 = v65
	goto L43
L43:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v125 == int32(92) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v472 = v215
	goto L38
L45:
	;
	v128 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v128)
	v132 = v114 + int32(1)
	goto L47
L46:
	;
	v132 = v114
	goto L47
L47:
	;
	v133 = F_pg_mblen_with_len(m, v113, v119)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v135 = v119 - v133
	if v133 <= int32(0) {
		v214 = v113
		v215 = v132
		goto L49
	} else {
		goto L50
	}
L49:
	;
	if int32(0) < v135 {
		v113 = v214
		v114 = v215
		v119 = v135
		goto L43
	} else {
		goto L61
	}
L50:
	;
	v140 = v133 & int32(7)
	if v140 != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v141 = v113
	v142 = v132
	v143 = v133
	v144 = int32(0)
	goto L54
L52:
	;
	v164 = v113
	v165 = v132
	v166 = v133
	goto L53
L53:
	;
	if base.Ui32(v133) < base.Ui32(int32(8)) {
		v214 = v164
		v215 = v165
		goto L49
	} else {
		goto L57
	}
L54:
	;
	v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v141))))
	*(*uint8)(unsafe.Add(mBase, uint32(v142))) = uint8(v153)
	v155 = int32(1)
	v156 = v142 + v155
	v158 = v141 + v155
	v160 = v143 - v155
	v162 = v144 + v155
	if v162 != v140 {
		v141 = v158
		v142 = v156
		v143 = v160
		v144 = v162
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v164 = v158
	v165 = v156
	v166 = v160
	goto L53
L56:
	;
	goto L55
L57:
	;
	v178 = v164
	v179 = v165
	v180 = v166
	goto L58
L58:
	;
	v190 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178))))
	*(*uint8)(unsafe.Add(mBase, uint32(v179))) = uint8(v190)
	v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+1)) = uint8(v192)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+2)) = uint8(v194)
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+3)) = uint8(v196)
	v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+4)) = uint8(v198)
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+5)) = uint8(v200)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+6)) = uint8(v202)
	v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v178)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v179)+7)) = uint8(v204)
	v206 = int32(8)
	v207 = v179 + v206
	v209 = v178 + v206
	if v206 < v180 {
		v178 = v209
		v179 = v207
		v180 = v180 - v206
		goto L58
	} else {
		goto L60
	}
L59:
	;
	v214 = v209
	v215 = v207
	goto L49
L60:
	;
	goto L59
L61:
	;
	goto L44
L62:
	;
	v230 = v69
	goto L64
L63:
	;
	v230 = v229
	goto L64
L64:
	;
	v231 = F_pg_mblen_with_len(m, v230, v99)
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	if v231 != v99 {
		goto L37
	} else {
		goto L66
	}
L66:
	;
	v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v236 = v234 & int32(1)
	if v236 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v237 = v69
	goto L69
L68:
	;
	v237 = v229
	goto L69
L69:
	;
	if v234 == int32(1) {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v265 == int32(92) {
		goto L36
	} else {
		goto L81
	}
L71:
	;
	v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69))))
	if v243 == int32(18) {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	v254 = int32(1)
	if v236 != 0 {
		v264 = int32(base.Ui32(v234)>>(uint(v254)%32)) - v254
		goto L70
	} else {
		goto L80
	}
L74:
	;
	v246 = int32(16)
	goto L76
L75:
	;
	v246 = int32(0)
	goto L76
L76:
	;
	if base.Ui32((v243-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v253 = int32(4)
	goto L79
L78:
	;
	v253 = v246
	goto L79
L79:
	;
	v264 = v253
	goto L70
L80:
	;
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v264 = int32(base.Ui32(v258)>>(uint(int32(2))%32)) - int32(4)
	goto L70
L81:
	;
	v268 = int32(0)
	if v65 <= v268 {
		v472 = v108
		goto L38
	} else {
		goto L82
	}
L82:
	;
	v271 = v100
	v272 = v108
	v275 = v268
	v277 = v65
	goto L83
L83:
	;
	v283 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v237))))
	if v283 == v284 {
		goto L88
	} else {
		goto L89
	}
L84:
	;
	v472 = v457
	goto L38
L85:
	;
	v468 = v277 - v461
	if int32(0) < v468 {
		v271 = v456
		v272 = v457
		v275 = v460
		v277 = v468
		goto L83
	} else {
		goto L121
	}
L86:
	;
	v453 = F_pg_mblen_with_len(m, v271, v277)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L1
	} else {
		goto L120
	}
L87:
	;
	v436 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v272)+1)) = uint8(v436)
	v444 = v349
	v452 = v272 + int32(2)
	goto L86
L88:
	;
	v286 = F_pg_mblen_with_len(m, v271, v277)
	mBase = m.M
	v287 = m.ExcPending
	if v287 != 0 {
		goto L1
	} else {
		goto L92
	}
L89:
	;
	v342 = v283
	goto L90
L90:
	;
	if v342&int32(255) == int32(92) {
		goto L102
	} else {
		goto L103
	}
L91:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v271))))
	v342 = v329
	goto L90
L92:
	;
	v288 = F_pg_mblen_with_len(m, v237, v264)
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L1
	} else {
		goto L93
	}
L93:
	;
	if v286 != v288 {
		goto L91
	} else {
		goto L94
	}
L94:
	;
	v293 = v286
	v294 = v271
	v296 = v237
	goto L95
L95:
	;
	if v293 != 0 {
		goto L97
	} else {
		goto L98
	}
L96:
	;
	if v275 != 0 {
		goto L91
	} else {
		goto L101
	}
L97:
	;
	v303 = int32(1)
	v305 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v296))))
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v305 == v306 {
		v293 = v293 - v303
		v294 = v294 + v303
		v296 = v296 + v303
		goto L95
	} else {
		goto L100
	}
L98:
	;
	goto L99
L99:
	;
	goto L96
L100:
	;
	goto L91
L101:
	;
	v312 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v312)
	v314 = int32(1)
	v444 = v314
	v452 = v272 + v314
	goto L86
L102:
	;
	v347 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v272))) = uint8(v347)
	v349 = int32(0)
	if v275 == v349 {
		goto L87
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v354 = F_pg_mblen_with_len(m, v271, v277)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L1
	} else {
		goto L106
	}
L105:
	;
	v444 = v349
	v452 = v272 + int32(1)
	goto L86
L106:
	;
	if v354 <= int32(0) {
		goto L107
	} else {
		goto L108
	}
L107:
	;
	v456 = v271
	v457 = v272
	v460 = int32(0)
	v461 = v354
	goto L85
L108:
	;
	goto L109
L109:
	;
	v361 = v354 & int32(7)
	if v361 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v362 = v271
	v363 = v272
	v364 = v354
	v365 = int32(0)
	goto L113
L111:
	;
	v385 = v271
	v386 = v272
	v387 = v354
	goto L112
L112:
	;
	v397 = int32(0)
	if base.Ui32(v354) < base.Ui32(int32(8)) {
		v456 = v385
		v457 = v386
		v460 = v397
		v461 = v354
		goto L85
	} else {
		goto L116
	}
L113:
	;
	v374 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v362))))
	*(*uint8)(unsafe.Add(mBase, uint32(v363))) = uint8(v374)
	v376 = int32(1)
	v377 = v363 + v376
	v379 = v362 + v376
	v381 = v364 - v376
	v383 = v365 + v376
	if v383 != v361 {
		v362 = v379
		v363 = v377
		v364 = v381
		v365 = v383
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v385 = v379
	v386 = v377
	v387 = v381
	goto L112
L115:
	;
	goto L114
L116:
	;
	v400 = v385
	v401 = v386
	v402 = v387
	goto L117
L117:
	;
	v412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400))))
	*(*uint8)(unsafe.Add(mBase, uint32(v401))) = uint8(v412)
	v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+1)) = uint8(v414)
	v416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+2)) = uint8(v416)
	v418 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+3)) = uint8(v418)
	v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+4)) = uint8(v420)
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+5)) = uint8(v422)
	v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+6)) = uint8(v424)
	v426 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v400)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v401)+7)) = uint8(v426)
	v428 = int32(8)
	v429 = v401 + v428
	v431 = v400 + v428
	if v428 < v402 {
		v400 = v431
		v401 = v429
		v402 = v402 - v428
		goto L117
	} else {
		goto L119
	}
L118:
	;
	v456 = v431
	v457 = v429
	v460 = v397
	v461 = v354
	goto L85
L119:
	;
	goto L118
L120:
	;
	v456 = v271 + v453
	v457 = v452
	v460 = v444
	v461 = v453
	goto L85
L121:
	;
	goto L84
L122:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v493 = m.ExcPending
	if v493 != 0 {
		goto L1
	} else {
		goto L123
	}
L123:
	;
	F_errmsg(m, int32(_a_F_like_escape_0), int32(0))
	mBase = m.M
	v497 = m.ExcPending
	if v497 != 0 {
		goto L1
	} else {
		goto L124
	}
L124:
	;
	F_errhint(m, int32(_a_F_like_escape_1), int32(0))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L125
	}
L125:
	;
	F_errfinish(m, int32(_a_F_like_escape_2), int32(438), int32(_a_F_like_escape_3))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L1
	} else {
		goto L126
	}
L126:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L127:
	;
	if v532 == int32(0) {
		goto L35
	} else {
		goto L138
	}
L128:
	;
	v511 = int32(18)
	v513 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v513 == v511 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L130
L130:
	;
	v524 = int32(1)
	if v507&v524 != 0 {
		v532 = int32(base.Ui32(v507) >> (uint(v524) % 32))
		goto L127
	} else {
		goto L137
	}
L131:
	;
	v516 = v511
	goto L133
L132:
	;
	v516 = int32(2)
	goto L133
L133:
	;
	if base.Ui32((v513-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v523 = int32(6)
	goto L136
L135:
	;
	v523 = v516
	goto L136
L136:
	;
	v532 = v523
	goto L127
L137:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v532 = int32(base.Ui32(v528) >> (uint(int32(2)) % 32))
	goto L127
L138:
	;
	base.MemoryCopy(m, v105, v14, v532)
	return v105
}
func F_like_escape_bytea(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum_packed(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum_packed(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_SB_do_like_escape(m, v3, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_lo_import_internal(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v121 int32
	_ = v121
	var v126 int32
	_ = v126
	v7 = m.G0
	v9 = v7 - int32(_a_F_lo_import_internal_0)
	m.G0 = v9
	F_PreventCommandIfReadOnly(m, int32(_a_F_lo_import_internal_1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = v9 + int32(48)
	F_text_to_cstring_buffer(m, l0, v17, int32(1024))
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v22 = F_OpenTransientFile(m, v17, int32(0))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L31
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L27
	}
L6:
	;
	if int32(0) <= v22 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_lo_import_internal[0])) = uint8(v27)
	v29 = F_inv_create(m, l1)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L1
	} else {
		goto L23
	}
L10:
	;
	v33 = *(*int32)(unsafe.Add(mBase, _c_F_lo_import_internal[1]))
	v34 = F_inv_open(m, v29, int32(_a_F_lo_import_internal_2), v33)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v39 = F_read(m, v22, v9+int32(1072), int32(_a_F_lo_import_internal_3))
	mBase = m.M
	if int32(0) < v39 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = v39
	goto L15
L13:
	;
	v56 = v39
	goto L14
L14:
	;
	if v56 < int32(0) {
		goto L5
	} else {
		goto L19
	}
L15:
	;
	v49 = v9 + int32(1072)
	v50 = F_inv_write(m, v34, v49, v42)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v56 = v53
	goto L14
L17:
	;
	v53 = F_read(m, v22, v49, int32(_a_F_lo_import_internal_3))
	mBase = m.M
	if int32(0) < v53 {
		v42 = v53
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_pfree(m, v34)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v66 = F_CloseTransientFile(m, v22)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v66 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	m.G0 = v9 + int32(_a_F_lo_import_internal_0)
	return v29
L23:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9))) = v9 + int32(48)
	F_errmsg(m, int32(_a_F_lo_import_internal_4), v9)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(_a_F_lo_import_internal_5), int32(445), int32(_a_F_lo_import_internal_6))
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L27:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v9 + int32(48)
	F_errmsg(m, int32(_a_F_lo_import_internal_7), v9+int32(16))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(_a_F_lo_import_internal_5), int32(468), int32(_a_F_lo_import_internal_6))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L31:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v9 + int32(48)
	F_errmsg(m, int32(_a_F_lo_import_internal_8), v9+int32(32))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(_a_F_lo_import_internal_5), int32(476), int32(_a_F_lo_import_internal_6))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L34
	}
L34:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_load_ident(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v1 int32
	_ = v1
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	v1 = int32(0)
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v1
	v17 = *(*int32)(unsafe.Add(mBase, _c_F_load_ident[0]))
	v21 = F_open_auth_file(m, v17, int32(15), v1, v1)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v119
L2:
	;
	return int32(0)
L3:
	;
	if v21 == int32(0) {
		v119 = v1
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_load_ident[0]))
	F_tokenize_auth_file(m, v28, v21, v12+int32(12), int32(15), int32(0))
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L2
	} else {
		goto L5
	}
L5:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_load_ident[1]))
	v41 = F_AllocSetContextCreateInternal(m, v36, int32(_a_F_load_ident_0), int32(0), int32(1024), int32(_a_F_load_ident_1))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v43 = int32(_a_F_load_ident_2)
	v44 = *(*int32)(unsafe.Add(mBase, _c_F_load_ident[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_load_ident[2])) = v41
	v47 = int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v48 == int32(0) {
		v88 = v47
		v92 = v1
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v94 = F_FreeFile(m, v21)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L2
	} else {
		goto L22
	}
L8:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v51 <= int32(0) {
		v88 = v47
		v92 = v1
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v54 = v1
	v57 = v47
	v61 = v1
	goto L10
L10:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v48)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v63+v54<<(uint(int32(2))%32))))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v67)+16))
	if v68 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v88 = v78
	v92 = v80
	goto L7
L12:
	;
	v82 = v54 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v82 < v83 {
		v54 = v82
		v57 = v78
		v61 = v80
		goto L10
	} else {
		goto L21
	}
L13:
	;
	v78 = int32(0)
	v80 = v61
	goto L12
L14:
	;
	goto L15
L15:
	;
	v71 = F_parse_ident_line(m, v67, int32(15))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L2
	} else {
		goto L16
	}
L16:
	;
	if v71 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v78 = int32(0)
	v80 = v61
	goto L12
L18:
	;
	goto L19
L19:
	;
	v76 = F_lappend(m, v61, v71)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v78 = v57
	v80 = v76
	goto L12
L21:
	;
	goto L11
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _c_F_load_ident[3]))
	F_MemoryContextDelete(m, v97)
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L2
	} else {
		goto L23
	}
L23:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_load_ident[2])) = v44
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, _c_F_load_ident[3])) = v103
	if v88 == v103 {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	F_MemoryContextDelete(m, v41)
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L2
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v111 = *(*int32)(unsafe.Add(mBase, _c_F_load_ident[4]))
	if v111 != 0 {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	v119 = int32(0)
	goto L1
L28:
	;
	F_MemoryContextDelete(m, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L2
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_load_ident[5])) = v92
	*(*int32)(unsafe.Add(mBase, _c_F_load_ident[4])) = v41
	v119 = int32(1)
	goto L1
L31:
	;
	goto L30
}
func F_log1p(m *base.Module, l0 float64) float64 {
	var v2 float64
	_ = v2
	var v8 int64
	_ = v8
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v31 int64
	_ = v31
	var v36 int32
	_ = v36
	var v49 float64
	_ = v49
	var v54 float64
	_ = v54
	var v68 float64
	_ = v68
	var v72 float64
	_ = v72
	var v73 float64
	_ = v73
	var v74 float64
	_ = v74
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v119 float64
	_ = v119
	v2 = float64(0)
	v8 = base.I64_reinterpret_f64(l0)
	if v8 <= int64(4601133429810003967) {
		if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v8) {
			if base.F64_eq(l0, float64(-1)) != 0 {
				v119 = math.Float64frombits(uint64(0xfff0000000000000))
				return v119
			} else {
				return base.F64_div(base.F64_sub(l0, l0), float64(0))
			}
		} else {
			if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v8)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
				return l0
			} else {
				if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v8) {
					v29 = float64(1)
					v30 = base.F64_add(l0, v29)
					v31 = base.I64_reinterpret_f64(v30)
					v36 = base.I32_wrap_i64(int64(base.Ui64(v31)>>(uint(int64(32))%64))) + int32(_a_F_log1p_0)
					if base.Ui32(int32(1074790399)) < base.Ui32(v36) {
						v49 = base.F64_add(base.F64_sub(l0, v30), v29)
					} else {
						v49 = base.F64_sub(l0, base.F64_add(v30, float64(-1)))
					}
					if base.Ui32(v36) <= base.Ui32(int32(1129316351)) {
						v54 = base.F64_div(v49, v30)
					} else {
						v54 = float64(0)
					}
					v68 = base.F64_convert_i32_s(int32(base.Ui32(v36)>>(uint(int32(20))%32)) - int32(1023))
					v72 = base.F64_add(base.F64_reinterpret_i64(v31&int64(4294967295)|base.I64_extend_i32_u(v36&int32(_a_F_log1p_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
					v73 = v68
					v74 = base.F64_add(base.F64_mul(v68, float64(1.9082149292705877e-10)), v54)
				} else {
					v72 = l0
					v73 = v2
					v74 = v2
				}
				v81 = base.F64_div(v72, base.F64_add(v72, float64(2)))
				v84 = base.F64_mul(v72, base.F64_mul(v72, float64(0.5)))
				v85 = base.F64_mul(v81, v81)
				v86 = base.F64_mul(v85, v85)
				v119 = base.F64_add(base.F64_mul(v73, float64(0.6931471803691238)), base.F64_add(v72, base.F64_sub(base.F64_add(base.F64_mul(v81, base.F64_add(v84, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v85, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v74), v84)))
				return v119
			}
		}
	} else {
		if base.Ui64(int64(9218868437227405311)) < base.Ui64(v8) {
			return l0
		} else {
			v29 = float64(1)
			v30 = base.F64_add(l0, v29)
			v31 = base.I64_reinterpret_f64(v30)
			v36 = base.I32_wrap_i64(int64(base.Ui64(v31)>>(uint(int64(32))%64))) + int32(_a_F_log1p_0)
			if base.Ui32(int32(1074790399)) < base.Ui32(v36) {
				v49 = base.F64_add(base.F64_sub(l0, v30), v29)
			} else {
				v49 = base.F64_sub(l0, base.F64_add(v30, float64(-1)))
			}
			if base.Ui32(v36) <= base.Ui32(int32(1129316351)) {
				v54 = base.F64_div(v49, v30)
			} else {
				v54 = float64(0)
			}
			v68 = base.F64_convert_i32_s(int32(base.Ui32(v36)>>(uint(int32(20))%32)) - int32(1023))
			v72 = base.F64_add(base.F64_reinterpret_i64(v31&int64(4294967295)|base.I64_extend_i32_u(v36&int32(_a_F_log1p_1)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
			v73 = v68
			v74 = base.F64_add(base.F64_mul(v68, float64(1.9082149292705877e-10)), v54)
			v81 = base.F64_div(v72, base.F64_add(v72, float64(2)))
			v84 = base.F64_mul(v72, base.F64_mul(v72, float64(0.5)))
			v85 = base.F64_mul(v81, v81)
			v86 = base.F64_mul(v85, v85)
			v119 = base.F64_add(base.F64_mul(v73, float64(0.6931471803691238)), base.F64_add(v72, base.F64_sub(base.F64_add(base.F64_mul(v81, base.F64_add(v84, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v85, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), v74), v84)))
			return v119
		}
	}
}
func F_logfile_open(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_logfile_open[0]))
	v16 = F_umask(m, (v11^int32(-1))&int32(383))
	mBase = m.M
	v17 = F_fopen(m, l0, l1)
	mBase = m.M
	v18 = F_umask(m, v16)
	mBase = m.M
	if v17 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = int32(-1)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v17)+48))
		if v21 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = int32(10)
		} else {
		}
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
		*(*int32)(unsafe.Add(mBase, uint32(v17))) = v24 | int32(64)
		m.G0 = v8 + int32(16)
		return v17
	} else {
		v29 = *(*int32)(unsafe.Add(mBase, _c_F_logfile_open[1]))
		if l2 != 0 {
			v32 = int32(15)
		} else {
			v32 = int32(22)
		}
		v34 = F_errstart(m, v32, int32(0))
		mBase = m.M
		v37 = m.ExcPending
		if v37 != 0 {
			return int32(0)
		} else {
			if v34 != 0 {
				F_errcode_for_file_access(m)
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg(m, int32(_a_F_logfile_open_0), v8)
					mBase = m.M
					v43 = m.ExcPending
					if v43 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_logfile_open_1), int32(1248), int32(_a_F_logfile_open_2))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_logfile_open[1])) = v29
							m.G0 = v8 + int32(16)
							return v17
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _c_F_logfile_open[1])) = v29
				m.G0 = v8 + int32(16)
				return v17
			}
		}
	}
}
func F_lpad(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v14 = F_pg_detoast_datum_packed(m, v13)
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
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = F_pg_detoast_datum_packed(m, v21)
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if int32(0) < v18 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = v18
	goto L6
L5:
	;
	v25 = int32(0)
	goto L6
L6:
	;
	v26 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v26 == int32(1) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v56 = int32(0)
	if v56 < v55 {
		goto L18
	} else {
		goto L19
	}
L8:
	;
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v32 == int32(18) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v43 = int32(1)
	if v26&v43 != 0 {
		v55 = int32(base.Ui32(v26)>>(uint(v43)%32)) - v43
		goto L7
	} else {
		goto L17
	}
L11:
	;
	v35 = int32(16)
	goto L13
L12:
	;
	v35 = int32(0)
	goto L13
L13:
	;
	if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v42 = int32(4)
	goto L16
L15:
	;
	v42 = v35
	goto L16
L16:
	;
	v55 = v42
	goto L7
L17:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
	goto L7
L18:
	;
	v59 = v55
	goto L20
L19:
	;
	v59 = v56
	goto L20
L20:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v60 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v90 = int32(1)
	v91 = v14 + v90
	v93 = v14 + int32(4)
	if v26&v90 != 0 {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v66 == int32(18) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v77 = int32(1)
	if v60&v77 != 0 {
		v89 = int32(base.Ui32(v60)>>(uint(v77)%32)) - v77
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v69 = int32(16)
	goto L27
L26:
	;
	v69 = int32(0)
	goto L27
L27:
	;
	if base.Ui32((v66-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v76 = int32(4)
	goto L30
L29:
	;
	v76 = v69
	goto L30
L30:
	;
	v89 = v76
	goto L21
L31:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v89 = int32(base.Ui32(v83)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v96 = v91
	goto L34
L33:
	;
	v96 = v93
	goto L34
L34:
	;
	v97 = F_pg_mbstrlen_with_len(m, v96, v59)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_lpad[0]))
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v100)+4))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v101*int32(28))+uint32(_c_F_lpad[1])))
	goto L37
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L1
	} else {
		goto L79
	}
L37:
	;
	if v97 < v25 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v109 = v97
	goto L40
L39:
	;
	v109 = v25
	goto L40
L40:
	;
	if v89 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v112 = v109
	goto L43
L42:
	;
	v112 = v25
	goto L43
L43:
	;
	v114 = base.I64_extend_i32_s(v106) * base.I64_extend_i32_s(v112)
	v118 = base.I32_wrap_i64(v114)
	if base.I32_wrap_i64(int64(base.Ui64(v114)>>(uint(int64(32))%64))) != v118>>(uint(int32(31))%32) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v123 = v118 + int32(4)
	if base.B2i32(v123 < v118)|base.B2i32(base.Ui32(int32(1073741824)) <= base.Ui32(v123)) != 0 {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	v128 = F_palloc(m, v123)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	v131 = v128 + int32(4)
	v132 = v112 - v109
	if v132 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v133 = int32(1)
	v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v135&v133 != 0 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	v166 = v131
	goto L49
L49:
	;
	if v109 != 0 {
		goto L66
	} else {
		goto L67
	}
L50:
	;
	v138 = v133
	goto L52
L51:
	;
	v138 = int32(4)
	goto L52
L52:
	;
	v139 = v22 + v138
	v140 = int32(0)
	if v140 < v89 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v143 = v89
	goto L55
L54:
	;
	v143 = v140
	goto L55
L55:
	;
	v144 = v139 + v143
	v145 = v131
	v146 = v139
	v150 = v132
	goto L56
L56:
	;
	v157 = F_pg_mblen_range(m, v146, v144)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L58
	}
L57:
	;
	v166 = v163
	goto L49
L58:
	;
	if v157 != 0 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	base.MemoryCopy(m, v145, v146, v157)
	goto L61
L60:
	;
	goto L61
L61:
	;
	v160 = v146 + v157
	if v160 == v144 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v162 = v139
	goto L64
L63:
	;
	v162 = v160
	goto L64
L64:
	;
	v163 = v145 + v157
	v165 = v150 - int32(1)
	if v165 != 0 {
		v145 = v163
		v146 = v162
		v150 = v165
		goto L56
	} else {
		goto L65
	}
L65:
	;
	goto L57
L66:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v178&int32(1) != 0 {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v201 = v166
	goto L68
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v128))) = (v201 - v128) << (uint(int32(2)) % 32)
	return v128
L69:
	;
	v181 = v91
	goto L71
L70:
	;
	v181 = v93
	goto L71
L71:
	;
	v182 = v166
	v183 = v181
	v189 = v109
	goto L72
L72:
	;
	v194 = F_pg_mblen_unbounded(m, v183)
	mBase = m.M
	v195 = m.ExcPending
	if v195 != 0 {
		goto L1
	} else {
		goto L74
	}
L73:
	;
	v201 = v198
	goto L68
L74:
	;
	if v194 != 0 {
		goto L75
	} else {
		goto L76
	}
L75:
	;
	base.MemoryCopy(m, v182, v183, v194)
	goto L77
L76:
	;
	goto L77
L77:
	;
	v198 = v182 + v194
	v200 = v189 - int32(1)
	if v200 != 0 {
		v182 = v198
		v183 = v183 + v194
		v189 = v200
		goto L72
	} else {
		goto L78
	}
L78:
	;
	goto L73
L79:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errmsg(m, int32(_a_F_lpad_0), int32(0))
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	F_errfinish(m, int32(_a_F_lpad_1), int32(206), int32(_a_F_lpad_2))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_lquery_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = F_deparse_lquery(m, v3)
		mBase = m.M
		v8 = m.ExcPending
		if v8 != 0 {
			return int32(0)
		} else {
			return v7
		}
	}
}
func F_lquery_recv(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pq_getmsgint(m, v8, int32(1))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 != int32(1) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v19 = m.ExcPending
			if v19 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v6))) = v10
				F_errmsg_internal(m, int32(_a_F_lquery_recv_0), v6)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_lquery_recv_1), int32(808), int32(_a_F_lquery_recv_2))
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v34 = F_pq_getmsgtext(m, v8, v29-v30, v6+int32(12))
			mBase = m.M
			v35 = m.ExcPending
			if v35 != 0 {
				return int32(0)
			} else {
				v37 = F_parse_lquery(m, v34, int32(0))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v34)
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(16)
						return v37
					}
				}
			}
		}
	}
}
func F_ltq_regex(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v11 = F_pg_detoast_datum(m, v10)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v19 = F_checkCond(m, v11+int32(16), v15, v6+int32(8), v18)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v21 != v6 {
					F_pfree(m, v6)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						if v25 != v11 {
							F_pfree(m, v11)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								return v19
							}
						} else {
							return v19
						}
					}
				} else {
					v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					if v25 != v11 {
						F_pfree(m, v11)
						mBase = m.M
						v28 = m.ExcPending
						if v28 != 0 {
							return int32(0)
						} else {
							return v19
						}
					} else {
						return v19
					}
				}
			}
		}
	}
}
