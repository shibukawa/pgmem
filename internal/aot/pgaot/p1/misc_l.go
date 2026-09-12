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
					v26 = *(*int32)(unsafe.Add(mBase, _consts[31]))
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
						F_sequence_close(m, v13, int32(0))
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
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = F_OpernameGetOprid(m, l0, l1, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l3 != 0 {
			m.G0 = v9 + int32(16)
			return v11
		} else {
			if v11 != 0 {
				m.G0 = v9 + int32(16)
				return v11
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					if l2 == int32(0) {
						F_errcode(m, int32(16801924))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(460459), int32(0))
							mBase = m.M
							v49 = m.ExcPending
							if v49 != 0 {
								return int32(0)
							} else {
								F_parser_errposition(m, int32(0), int32(-1))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(514923), int32(115), int32(397213))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
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
						v23 = m.ExcPending
						if v23 != 0 {
							return int32(0)
						} else {
							v24 = F_op_signature_string(m, l0, l1, l2)
							mBase = m.M
							v25 = m.ExcPending
							if v25 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = v24
								F_errmsg(m, int32(208127), v9)
								mBase = m.M
								v29 = m.ExcPending
								if v29 != 0 {
									return int32(0)
								} else {
									F_parser_errposition(m, int32(0), int32(-1))
									mBase = m.M
									v33 = m.ExcPending
									if v33 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(514923), int32(121), int32(397213))
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
					}
				}
			}
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
	v7 = *(*int32)(unsafe.Add(mBase, _consts[799]))
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
	v14 = *(*int32)(unsafe.Add(mBase, _consts[800]))
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[137]))
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
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
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
	var v262 int32
	_ = v262
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
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v334 int32
	_ = v334
	v3 = int32(0)
	if l1 <= v3 {
		v334 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v334
L2:
	;
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
	if v15 == int32(0) {
		v334 = v3
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
		v334 = v37
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
	v50 = v38 + int32(8)
	v51 = v14 + int32(8)
	v52 = v37
	goto L15
L15:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v51))))
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v50))))
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
	v63 = v51 + v62
	v65 = v50 + v62
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
	v130 = int32(131064)
	v139 = v52 + int32(1)
	if v139 != v43 {
		v50 = v50 + (v60+v128)&v130
		v51 = v51 + (v59+v128)&v130
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
	v270 = F_palloc0(m, v262)
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L53
	} else {
		goto L54
	}
L40:
	;
	v262 = int32(8)
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
	v184 = v176
	v185 = v174
	v187 = int32(0)
	goto L46
L44:
	;
	v226 = v176
	v227 = v174
	goto L45
L45:
	;
	if v173 == int32(0) {
		v262 = v227
		goto L39
	} else {
		goto L49
	}
L46:
	;
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v184))))
	v194 = int32(9)
	v196 = int32(131064)
	v197 = (v193 + v194) & v196
	v198 = v184 + v197
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198))))
	v203 = (v199 + v194) & v196
	v204 = v198 + v203
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v204))))
	v209 = (v205 + v194) & v196
	v210 = v204 + v209
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v210))))
	v215 = (v211 + v194) & v196
	v216 = v210 + v215
	v220 = v215 + (v209 + (v203 + (v185 + v197)))
	v222 = v187 + int32(4)
	if v222 != v162&int32(2147483644) {
		v184 = v216
		v185 = v220
		v187 = v222
		goto L46
	} else {
		goto L48
	}
L47:
	;
	v226 = v216
	v227 = v220
	goto L45
L48:
	;
	goto L47
L49:
	;
	v240 = v226
	v241 = v227
	v243 = int32(0)
	goto L50
L50:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v240))))
	v253 = (v249 + int32(9)) & int32(131064)
	v255 = v241 + v253
	v257 = v243 + int32(1)
	if v257 != v173 {
		v240 = v240 + v253
		v241 = v255
		v243 = v257
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v262 = v255
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
	*(*int32)(unsafe.Add(mBase, uint32(v270))) = v262 << (uint(int32(2)) % 32)
	if int32(0) < v162 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v280 = int32(8)
	v282 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v288 = v282 + v280
	v289 = v270 + v280
	v291 = int32(0)
	goto L58
L56:
	;
	goto L57
L57:
	;
	v334 = v270
	goto L1
L58:
	;
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288))))
	v301 = (v297 + int32(9)) & int32(131064)
	if v301 != 0 {
		goto L61
	} else {
		goto L62
	}
L59:
	;
	goto L57
L60:
	;
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v303))))
	v305 = int32(9)
	v307 = int32(131064)
	v310 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v288))))
	v317 = v291 + int32(1)
	if v317 != v162 {
		v288 = v288 + (v310+v305)&v307
		v289 = v303 + (v304+v305)&v307
		v291 = v317
		goto L58
	} else {
		goto L64
	}
L61:
	;
	v302 = F__emscripten_memcpy_bulkmem(m, v289, v288, v301)
	mBase = m.M
	v303 = v302
	goto L63
L62:
	;
	v303 = v289
	goto L63
L63:
	;
	goto L60
L64:
	;
	goto L59
}
func F_lcons(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	if l1 == int32(0) {
		v7 = F_palloc(m, int32(32))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(4294967297)
			v16 = v7 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
			return v7
		}
	} else {
		F_new_head_cell(m, l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
			return l1
		}
	}
}
func F_lcons_int(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	if l1 == int32(0) {
		v7 = F_palloc(m, int32(32))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(4)
			*(*int64)(unsafe.Add(mBase, uint32(v7))) = int64(4294967767)
			v16 = v7 + int32(16)
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v16
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = l0
			return v7
		}
	} else {
		F_new_head_cell(m, l1)
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = l0
			return l1
		}
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
	v6 = F_DirectFunctionCall3Coll(m, int32(490), v2, int32(785340), v2, int32(-1))
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
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
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
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v261 int32
	_ = v261
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v334 int32
	_ = v334
	var v347 int32
	_ = v347
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v372 int32
	_ = v372
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v438 int32
	_ = v438
	var v445 int32
	_ = v445
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v469 int32
	_ = v469
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v481 int32
	_ = v481
	var v500 int32
	_ = v500
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
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
	v22 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v23*int32(28))+uint32(_consts[355])))
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
	v69 = int32(1)
	v70 = v19 + v69
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v73 = v71 & v69
	if v71 == v69 {
		goto L21
	} else {
		goto L22
	}
L10:
	;
	v41 = int32(4)
	v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v43&int32(254) == int32(2) {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	goto L12
L12:
	;
	v56 = int32(1)
	if v38 != 0 {
		v66 = int32(base.Ui32(v36)>>(uint(v56)%32)) - v56
		goto L9
	} else {
		goto L19
	}
L13:
	;
	v52 = v41
	goto L15
L14:
	;
	v52 = base.B2i32(v43 == int32(18)) << (uint(v41) % 32)
	goto L15
L15:
	;
	if v43 == int32(1) {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v55 = v41
	goto L18
L17:
	;
	v55 = v52
	goto L18
L18:
	;
	v66 = v55
	goto L9
L19:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v66 = int32(base.Ui32(v60)>>(uint(int32(2))%32)) - int32(4)
	goto L9
L20:
	;
	if v38 != 0 {
		goto L28
	} else {
		goto L29
	}
L21:
	;
	v76 = int32(4)
	v77 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if base.Ui32((v77-int32(1))&int32(255)) < base.Ui32(int32(3)) {
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
	if v73 != 0 {
		v99 = int32(base.Ui32(v71)>>(uint(v89)%32)) - v89
		goto L20
	} else {
		goto L27
	}
L24:
	;
	v88 = v76
	goto L26
L25:
	;
	v88 = base.B2i32(v77 == int32(18)) << (uint(v76) % 32)
	goto L26
L26:
	;
	v99 = v88
	goto L20
L27:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v99 = int32(base.Ui32(v93)>>(uint(int32(2))%32)) - int32(4)
	goto L20
L28:
	;
	v100 = v35
	goto L30
L29:
	;
	v100 = v14 + int32(4)
	goto L30
L30:
	;
	v105 = F_palloc(m, v66<<(uint(int32(1))%32)+int32(4))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v108 = v105 + int32(4)
	if v99 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L32:
	;
	v517 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v517 == int32(1) {
		goto L123
	} else {
		goto L124
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L118
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105))) = (v481 - v105) << (uint(int32(2)) % 32)
	return v105
L35:
	;
	if v66 <= int32(0) {
		v481 = v108
		goto L34
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v231 = v19 + int32(4)
	if v73 != 0 {
		goto L58
	} else {
		goto L59
	}
L38:
	;
	v113 = v100
	v114 = v108
	v119 = v66
	goto L39
L39:
	;
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v113))))
	if v125 == int32(92) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v481 = v217
	goto L34
L41:
	;
	v128 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v114))) = uint8(v128)
	v132 = v114 + int32(1)
	goto L43
L42:
	;
	v132 = v114
	goto L43
L43:
	;
	v133 = F_pg_mblen_with_len(m, v113, v119)
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	v135 = v119 - v133
	if v133 <= int32(0) {
		v216 = v113
		v217 = v132
		goto L45
	} else {
		goto L46
	}
L45:
	;
	if int32(0) < v135 {
		v113 = v216
		v114 = v217
		v119 = v135
		goto L39
	} else {
		goto L57
	}
L46:
	;
	v140 = v133 & int32(7)
	if v140 != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v141 = v113
	v142 = v132
	v143 = v133
	v144 = int32(0)
	goto L50
L48:
	;
	v164 = v113
	v165 = v132
	v166 = v133
	goto L49
L49:
	;
	if base.Ui32(v133) < base.Ui32(int32(8)) {
		v216 = v164
		v217 = v165
		goto L45
	} else {
		goto L53
	}
L50:
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
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v164 = v158
	v165 = v156
	v166 = v160
	goto L49
L52:
	;
	goto L51
L53:
	;
	v178 = v164
	v179 = v165
	v180 = v166
	goto L54
L54:
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
	if base.Ui32(v180-int32(9)) < base.Ui32(int32(-2)) {
		v178 = v209
		v179 = v207
		v180 = v180 - v206
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v216 = v209
	v217 = v207
	goto L45
L56:
	;
	goto L55
L57:
	;
	goto L40
L58:
	;
	v232 = v70
	goto L60
L59:
	;
	v232 = v231
	goto L60
L60:
	;
	v233 = F_pg_mblen_with_len(m, v232, v99)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	if v233 != v99 {
		goto L33
	} else {
		goto L62
	}
L62:
	;
	v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	v238 = v236 & int32(1)
	if v238 != 0 {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v239 = v70
	goto L65
L64:
	;
	v239 = v231
	goto L65
L65:
	;
	if v236 == int32(1) {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	if v268 == int32(92) {
		goto L32
	} else {
		goto L77
	}
L67:
	;
	v242 = int32(4)
	v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
	if v244&int32(254) == int32(2) {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v257 = int32(1)
	if v238 != 0 {
		v267 = int32(base.Ui32(v236)>>(uint(v257)%32)) - v257
		goto L66
	} else {
		goto L76
	}
L70:
	;
	v253 = v242
	goto L72
L71:
	;
	v253 = base.B2i32(v244 == int32(18)) << (uint(v242) % 32)
	goto L72
L72:
	;
	if v244 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v256 = v242
	goto L75
L74:
	;
	v256 = v253
	goto L75
L75:
	;
	v267 = v256
	goto L66
L76:
	;
	v261 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v267 = int32(base.Ui32(v261)>>(uint(int32(2))%32)) - int32(4)
	goto L66
L77:
	;
	v271 = int32(0)
	if v66 <= v271 {
		v481 = v108
		goto L34
	} else {
		goto L78
	}
L78:
	;
	v274 = v100
	v275 = v108
	v280 = v66
	v281 = v271
	goto L79
L79:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	v287 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v239))))
	if v286 == v287 {
		goto L84
	} else {
		goto L85
	}
L80:
	;
	v481 = v466
	goto L34
L81:
	;
	v477 = v280 - v469
	if int32(0) < v477 {
		v274 = v465
		v275 = v466
		v280 = v477
		v281 = v472
		goto L79
	} else {
		goto L117
	}
L82:
	;
	v462 = F_pg_mblen_with_len(m, v274, v280)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L116
	}
L83:
	;
	v445 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v275)+1)) = uint8(v445)
	v456 = v356
	v461 = v275 + int32(2)
	goto L82
L84:
	;
	v289 = F_pg_mblen_with_len(m, v274, v280)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L88
	}
L85:
	;
	v347 = v286
	goto L86
L86:
	;
	if v347&int32(255) == int32(92) {
		goto L98
	} else {
		goto L99
	}
L87:
	;
	v334 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v274))))
	v347 = v334
	goto L86
L88:
	;
	v291 = F_pg_mblen_with_len(m, v239, v267)
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L1
	} else {
		goto L89
	}
L89:
	;
	if v289 != v291 {
		goto L87
	} else {
		goto L90
	}
L90:
	;
	v296 = v289
	v297 = v274
	v299 = v239
	goto L91
L91:
	;
	if v296 != 0 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	if v281&int32(1) != 0 {
		goto L87
	} else {
		goto L97
	}
L93:
	;
	v306 = int32(1)
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v299))))
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v297))))
	if v308 == v309 {
		v296 = v296 - v306
		v297 = v297 + v306
		v299 = v299 + v306
		goto L91
	} else {
		goto L96
	}
L94:
	;
	goto L95
L95:
	;
	goto L92
L96:
	;
	goto L87
L97:
	;
	v317 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v275))) = uint8(v317)
	v319 = int32(1)
	v456 = v319
	v461 = v275 + v319
	goto L82
L98:
	;
	v352 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v275))) = uint8(v352)
	v356 = int32(0)
	if v281&int32(1) == v356 {
		goto L83
	} else {
		goto L101
	}
L99:
	;
	goto L100
L100:
	;
	v361 = F_pg_mblen_with_len(m, v274, v280)
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L1
	} else {
		goto L102
	}
L101:
	;
	v456 = v356
	v461 = v275 + int32(1)
	goto L82
L102:
	;
	if v361 <= int32(0) {
		goto L103
	} else {
		goto L104
	}
L103:
	;
	v465 = v274
	v466 = v275
	v469 = v361
	v472 = int32(0)
	goto L81
L104:
	;
	goto L105
L105:
	;
	v368 = v361 & int32(7)
	if v368 != 0 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v369 = v274
	v370 = v275
	v371 = v361
	v372 = int32(0)
	goto L109
L107:
	;
	v392 = v274
	v393 = v275
	v394 = v361
	goto L108
L108:
	;
	v404 = int32(0)
	if base.Ui32(v361) < base.Ui32(int32(8)) {
		v465 = v392
		v466 = v393
		v469 = v361
		v472 = v404
		goto L81
	} else {
		goto L112
	}
L109:
	;
	v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v369))))
	*(*uint8)(unsafe.Add(mBase, uint32(v370))) = uint8(v381)
	v383 = int32(1)
	v384 = v370 + v383
	v386 = v369 + v383
	v388 = v371 - v383
	v390 = v372 + v383
	if v390 != v368 {
		v369 = v386
		v370 = v384
		v371 = v388
		v372 = v390
		goto L109
	} else {
		goto L111
	}
L110:
	;
	v392 = v386
	v393 = v384
	v394 = v388
	goto L108
L111:
	;
	goto L110
L112:
	;
	v407 = v392
	v408 = v393
	v409 = v394
	goto L113
L113:
	;
	v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407))))
	*(*uint8)(unsafe.Add(mBase, uint32(v408))) = uint8(v419)
	v421 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+1)) = uint8(v421)
	v423 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+2)))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+2)) = uint8(v423)
	v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+3)))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+3)) = uint8(v425)
	v427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+4)) = uint8(v427)
	v429 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+5)))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+5)) = uint8(v429)
	v431 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+6)))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+6)) = uint8(v431)
	v433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v407)+7)))
	*(*uint8)(unsafe.Add(mBase, uint32(v408)+7)) = uint8(v433)
	v435 = int32(8)
	v436 = v408 + v435
	v438 = v407 + v435
	if base.Ui32(v409-int32(9)) < base.Ui32(int32(-2)) {
		v407 = v438
		v408 = v436
		v409 = v409 - v435
		goto L113
	} else {
		goto L115
	}
L114:
	;
	v465 = v438
	v466 = v436
	v469 = v361
	v472 = v404
	goto L81
L115:
	;
	goto L114
L116:
	;
	v465 = v274 + v462
	v466 = v461
	v469 = v462
	v472 = v456
	goto L81
L117:
	;
	goto L80
L118:
	;
	F_errcode(m, int32(84410498))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L1
	} else {
		goto L119
	}
L119:
	;
	F_errmsg(m, int32(344181), int32(0))
	mBase = m.M
	v507 = m.ExcPending
	if v507 != 0 {
		goto L1
	} else {
		goto L120
	}
L120:
	;
	F_errhint(m, int32(631920), int32(0))
	mBase = m.M
	v511 = m.ExcPending
	if v511 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	F_errfinish(m, int32(518052), int32(438), int32(387035))
	mBase = m.M
	v516 = m.ExcPending
	if v516 != 0 {
		goto L1
	} else {
		goto L122
	}
L122:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L123:
	;
	v520 = int32(6)
	v522 = int32(18)
	v524 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35))))
	if v524 == v522 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	if v517&int32(1) != 0 {
		goto L139
	} else {
		goto L140
	}
L126:
	;
	v527 = v522
	goto L128
L127:
	;
	v527 = int32(2)
	goto L128
L128:
	;
	if v524&int32(254) == int32(2) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v532 = v520
	goto L131
L130:
	;
	v532 = v527
	goto L131
L131:
	;
	if v524 == int32(1) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	v535 = v520
	goto L134
L133:
	;
	v535 = v532
	goto L134
L134:
	;
	if v535 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	return v537
L136:
	;
	v536 = F__emscripten_memcpy_bulkmem(m, v105, v14, v535)
	mBase = m.M
	v537 = v536
	goto L138
L137:
	;
	v537 = v105
	goto L138
L138:
	;
	goto L135
L139:
	;
	v542 = int32(base.Ui32(v517) >> (uint(int32(1)) % 32))
	if v542 != 0 {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	goto L141
L141:
	;
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v548 = int32(base.Ui32(v546) >> (uint(int32(2)) % 32))
	if v548 != 0 {
		goto L147
	} else {
		goto L148
	}
L142:
	;
	return v544
L143:
	;
	v543 = F__emscripten_memcpy_bulkmem(m, v105, v14, v542)
	mBase = m.M
	v544 = v543
	goto L145
L144:
	;
	v544 = v105
	goto L145
L145:
	;
	goto L142
L146:
	;
	return v550
L147:
	;
	v549 = F__emscripten_memcpy_bulkmem(m, v105, v14, v548)
	mBase = m.M
	v550 = v549
	goto L149
L148:
	;
	v550 = v105
	goto L149
L149:
	;
	goto L146
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
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
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
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	v6 = m.G0
	v8 = v6 - int32(9264)
	m.G0 = v8
	F_PreventCommandIfReadOnly(m, int32(708561))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	F_text_to_cstring_buffer(m, l0, v8+int32(48), int32(1024))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v23 = F_OpenTransientFile(m, v8+int32(48), int32(0))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L6
	}
L4:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L31
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		goto L1
	} else {
		goto L27
	}
L6:
	;
	if int32(0) <= v23 {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v28 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _consts[443])) = uint8(v28)
	v30 = F_inv_create(m, l1)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
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
	v76 = m.ExcPending
	if v76 != 0 {
		goto L1
	} else {
		goto L23
	}
L10:
	;
	v34 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v35 = F_inv_open(m, v30, int32(131072), v34)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	v40 = F_read(m, v23, v8+int32(1072), int32(8192))
	mBase = m.M
	if int32(0) < v40 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v43 = v40
	goto L15
L13:
	;
	v58 = v40
	goto L14
L14:
	;
	if v58 < int32(0) {
		goto L5
	} else {
		goto L19
	}
L15:
	;
	v50 = F_inv_write(m, v35, v8+int32(1072), v43)
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	v58 = v55
	goto L14
L17:
	;
	v55 = F_read(m, v23, v8+int32(1072), int32(8192))
	mBase = m.M
	if int32(0) < v55 {
		v43 = v55
		goto L15
	} else {
		goto L18
	}
L18:
	;
	goto L16
L19:
	;
	F_pfree(m, v35)
	mBase = m.M
	v66 = m.ExcPending
	if v66 != 0 {
		goto L1
	} else {
		goto L20
	}
L20:
	;
	v67 = F_CloseTransientFile(m, v23)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L21
	}
L21:
	;
	if v67 != 0 {
		goto L4
	} else {
		goto L22
	}
L22:
	;
	m.G0 = v8 + int32(9264)
	return v30
L23:
	;
	F_errcode_for_file_access(m)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8))) = v8 + int32(48)
	F_errmsg(m, int32(309802), v8)
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	F_errfinish(m, int32(514476), int32(445), int32(322982))
	mBase = m.M
	v89 = m.ExcPending
	if v89 != 0 {
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
	v95 = m.ExcPending
	if v95 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v8 + int32(48)
	F_errmsg(m, int32(309913), v8+int32(16))
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L29
	}
L29:
	;
	F_errfinish(m, int32(514476), int32(468), int32(322982))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
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
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v8 + int32(48)
	F_errmsg(m, int32(311009), v8+int32(32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	F_errfinish(m, int32(514476), int32(476), int32(322982))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
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
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
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
	v17 = *(*int32)(unsafe.Add(mBase, _consts[448]))
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
	v28 = *(*int32)(unsafe.Add(mBase, _consts[448]))
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
	v36 = *(*int32)(unsafe.Add(mBase, _consts[449]))
	v41 = F_AllocSetContextCreateInternal(m, v36, int32(65621), int32(0), int32(1024), int32(8192))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L2
	} else {
		goto L6
	}
L6:
	;
	v43 = int32(4548768)
	v44 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v41
	v47 = int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
	if v48 == int32(0) {
		v87 = v47
		v88 = v1
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
		v87 = v47
		v88 = v1
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v54 = v1
	v56 = v47
	v57 = v1
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
	v87 = v78
	v88 = v79
	goto L7
L12:
	;
	v82 = v54 + int32(1)
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v48)+4))
	if v82 < v83 {
		v54 = v82
		v56 = v78
		v57 = v79
		goto L10
	} else {
		goto L21
	}
L13:
	;
	v78 = int32(0)
	v79 = v57
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
	v79 = v57
	goto L12
L18:
	;
	goto L19
L19:
	;
	v76 = F_lappend(m, v57, v71)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L20
	}
L20:
	;
	v78 = v56
	v79 = v76
	goto L12
L21:
	;
	goto L11
L22:
	;
	v97 = *(*int32)(unsafe.Add(mBase, _consts[444]))
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
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v44
	v103 = int32(0)
	*(*int32)(unsafe.Add(mBase, _consts[444])) = v103
	if v87 == v103 {
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
	v111 = *(*int32)(unsafe.Add(mBase, _consts[450]))
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
	*(*int32)(unsafe.Add(mBase, _consts[451])) = v88
	*(*int32)(unsafe.Add(mBase, _consts[450])) = v41
	v119 = int32(1)
	goto L1
L31:
	;
	goto L30
}
func F_log1p(m *base.Module, l0 float64) float64 {
	var v2 float64
	_ = v2
	var v9 int64
	_ = v9
	var v32 float64
	_ = v32
	var v33 int64
	_ = v33
	var v38 int32
	_ = v38
	var v53 float64
	_ = v53
	var v55 float64
	_ = v55
	var v70 float64
	_ = v70
	var v72 float64
	_ = v72
	var v76 float64
	_ = v76
	var v81 float64
	_ = v81
	var v84 float64
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v124 float64
	_ = v124
	v2 = float64(0)
	v9 = base.I64_reinterpret_f64(l0)
	if v9 <= int64(4601133429810003967) {
		if base.Ui64(int64(-4616189618054758400)) <= base.Ui64(v9) {
			if base.F64_eq(l0, float64(-1)) != 0 {
				v124 = math.Float64frombits(uint64(0xfff0000000000000))
				return v124
			} else {
				return base.F64_div(base.F64_sub(l0, l0), float64(0))
			}
		} else {
			if base.Ui32(base.I32_wrap_i64(int64(base.Ui64(v9)>>(uint(int64(31))%64)))) < base.Ui32(int32(2034237440)) {
				return l0
			} else {
				if base.Ui64(int64(-4624424114038243328)) <= base.Ui64(v9) {
					v32 = base.F64_add(l0, float64(1))
					v33 = base.I64_reinterpret_f64(v32)
					v38 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(32))%64))) + int32(614242)
					if base.Ui32(v38) <= base.Ui32(int32(1129316351)) {
						if base.Ui32(int32(1074790399)) < base.Ui32(v38) {
							v53 = base.F64_add(base.F64_sub(l0, v32), float64(1))
						} else {
							v53 = base.F64_sub(l0, base.F64_add(v32, float64(-1)))
						}
						v55 = base.F64_div(v53, v32)
					} else {
						v55 = v2
					}
					v70 = base.F64_add(base.F64_reinterpret_i64(v33&int64(4294967295)|base.I64_extend_i32_u(v38&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
					v72 = v55
					v76 = base.F64_convert_i32_s(int32(base.Ui32(v38)>>(uint(int32(20))%32)) - int32(1023))
				} else {
					v70 = l0
					v72 = v2
					v76 = float64(0)
				}
				v81 = base.F64_div(v70, base.F64_add(v70, float64(2)))
				v84 = base.F64_mul(v70, base.F64_mul(v70, float64(0.5)))
				v85 = base.F64_mul(v81, v81)
				v86 = base.F64_mul(v85, v85)
				v124 = base.F64_add(base.F64_mul(v76, float64(0.6931471803691238)), base.F64_add(v70, base.F64_sub(base.F64_add(base.F64_mul(v81, base.F64_add(v84, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v85, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v76, float64(1.9082149292705877e-10)), v72)), v84)))
				return v124
			}
		}
	} else {
		if base.Ui64(int64(9218868437227405311)) < base.Ui64(v9) {
			return l0
		} else {
			v32 = base.F64_add(l0, float64(1))
			v33 = base.I64_reinterpret_f64(v32)
			v38 = base.I32_wrap_i64(int64(base.Ui64(v33)>>(uint(int64(32))%64))) + int32(614242)
			if base.Ui32(v38) <= base.Ui32(int32(1129316351)) {
				if base.Ui32(int32(1074790399)) < base.Ui32(v38) {
					v53 = base.F64_add(base.F64_sub(l0, v32), float64(1))
				} else {
					v53 = base.F64_sub(l0, base.F64_add(v32, float64(-1)))
				}
				v55 = base.F64_div(v53, v32)
			} else {
				v55 = v2
			}
			v70 = base.F64_add(base.F64_reinterpret_i64(v33&int64(4294967295)|base.I64_extend_i32_u(v38&int32(1048575)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
			v72 = v55
			v76 = base.F64_convert_i32_s(int32(base.Ui32(v38)>>(uint(int32(20))%32)) - int32(1023))
			v81 = base.F64_div(v70, base.F64_add(v70, float64(2)))
			v84 = base.F64_mul(v70, base.F64_mul(v70, float64(0.5)))
			v85 = base.F64_mul(v81, v81)
			v86 = base.F64_mul(v85, v85)
			v124 = base.F64_add(base.F64_mul(v76, float64(0.6931471803691238)), base.F64_add(v70, base.F64_sub(base.F64_add(base.F64_mul(v81, base.F64_add(v84, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v85, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, base.F64_add(base.F64_mul(v86, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_add(base.F64_mul(v76, float64(1.9082149292705877e-10)), v72)), v84)))
			return v124
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
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _consts[428]))
	v17 = int32(4442524)
	v18 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	*(*int32)(unsafe.Add(mBase, _consts[361])) = (v11 ^ int32(-1)) & int32(383)
	v21 = F___syscall_ret(m, v18)
	mBase = m.M
	v22 = F_fopen(m, l0, l1)
	mBase = m.M
	v24 = int32(4442524)
	v25 = *(*int32)(unsafe.Add(mBase, _consts[361]))
	*(*int32)(unsafe.Add(mBase, _consts[361])) = v21
	v28 = F___syscall_ret(m, v25)
	mBase = m.M
	if v22 != 0 {
		*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = int32(-1)
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
		if v31 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v22)+80)) = int32(10)
		} else {
		}
		v34 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
		*(*int32)(unsafe.Add(mBase, uint32(v22))) = v34 | int32(64)
		m.G0 = v8 + int32(16)
		return v22
	} else {
		v39 = *(*int32)(unsafe.Add(mBase, _consts[137]))
		if l2 != 0 {
			v42 = int32(15)
		} else {
			v42 = int32(22)
		}
		v44 = F_errstart(m, v42, int32(0))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			if v44 != 0 {
				F_errcode_for_file_access(m)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg(m, int32(310623), v8)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(515104), int32(1248), int32(292728))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[137])) = v39
							m.G0 = v8 + int32(16)
							return v22
						}
					}
				}
			} else {
				*(*int32)(unsafe.Add(mBase, _consts[137])) = v39
				m.G0 = v8 + int32(16)
				return v22
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v116 int64
	_ = v116
	var v120 int32
	_ = v120
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
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
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v24 == int32(1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if int32(0) < v18 {
		goto L15
	} else {
		goto L16
	}
L5:
	;
	v27 = int32(4)
	v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+1)))
	if v29&int32(254) == int32(2) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L7
L7:
	;
	v42 = int32(1)
	if v24&v42 != 0 {
		v54 = int32(base.Ui32(v24)>>(uint(v42)%32)) - v42
		goto L4
	} else {
		goto L14
	}
L8:
	;
	v38 = v27
	goto L10
L9:
	;
	v38 = base.B2i32(v29 == int32(18)) << (uint(v27) % 32)
	goto L10
L10:
	;
	if v29 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v41 = v27
	goto L13
L12:
	;
	v41 = v38
	goto L13
L13:
	;
	v54 = v41
	goto L4
L14:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
	v54 = int32(base.Ui32(v48)>>(uint(int32(2))%32)) - int32(4)
	goto L4
L15:
	;
	v56 = v18
	goto L17
L16:
	;
	v56 = int32(0)
	goto L17
L17:
	;
	v57 = int32(0)
	if v57 < v54 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = v54
	goto L20
L19:
	;
	v60 = v57
	goto L20
L20:
	;
	v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v61 == int32(1) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v92 = int32(1)
	v93 = v14 + v92
	v95 = v14 + int32(4)
	if v24&v92 != 0 {
		goto L32
	} else {
		goto L33
	}
L22:
	;
	v64 = int32(4)
	v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+1)))
	if v66&int32(254) == int32(2) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v79 = int32(1)
	if v61&v79 != 0 {
		v91 = int32(base.Ui32(v61)>>(uint(v79)%32)) - v79
		goto L21
	} else {
		goto L31
	}
L25:
	;
	v75 = v64
	goto L27
L26:
	;
	v75 = base.B2i32(v66 == int32(18)) << (uint(v64) % 32)
	goto L27
L27:
	;
	if v66 == int32(1) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	v78 = v64
	goto L30
L29:
	;
	v78 = v75
	goto L30
L30:
	;
	v91 = v78
	goto L21
L31:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v91 = int32(base.Ui32(v85)>>(uint(int32(2))%32)) - int32(4)
	goto L21
L32:
	;
	v98 = v93
	goto L34
L33:
	;
	v98 = v95
	goto L34
L34:
	;
	v99 = F_pg_mbstrlen_with_len(m, v98, v60)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	v102 = *(*int32)(unsafe.Add(mBase, _consts[356]))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v103*int32(28))+uint32(_consts[355])))
	goto L37
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L82
	}
L37:
	;
	if v99 < v56 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v111 = v99
	goto L40
L39:
	;
	v111 = v56
	goto L40
L40:
	;
	if v91 <= int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v114 = v111
	goto L43
L42:
	;
	v114 = v56
	goto L43
L43:
	;
	v116 = base.I64_extend_i32_s(v108) * base.I64_extend_i32_s(v114)
	v120 = base.I32_wrap_i64(v116)
	if base.I32_wrap_i64(int64(base.Ui64(v116)>>(uint(int64(32))%64))) != v120>>(uint(int32(31))%32) {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v125 = v120 + int32(4)
	if v125 < v120 {
		goto L36
	} else {
		goto L45
	}
L45:
	;
	if base.Ui32(int32(1073741824)) <= base.Ui32(v125) {
		goto L36
	} else {
		goto L46
	}
L46:
	;
	v129 = F_palloc(m, v125)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	v132 = v129 + int32(4)
	v133 = v114 - v111
	if v133 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v134 = int32(1)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22))))
	if v136&v134 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	v168 = v132
	goto L50
L50:
	;
	if v111 != 0 {
		goto L68
	} else {
		goto L69
	}
L51:
	;
	v139 = v134
	goto L53
L52:
	;
	v139 = int32(4)
	goto L53
L53:
	;
	v140 = v22 + v139
	v141 = int32(0)
	if v141 < v91 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v144 = v91
	goto L56
L55:
	;
	v144 = v141
	goto L56
L56:
	;
	v145 = v140 + v144
	v146 = v132
	v148 = v140
	v150 = v133
	goto L57
L57:
	;
	v158 = F_pg_mblen_range(m, v148, v145)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v168 = v165
	goto L50
L59:
	;
	if v158 != 0 {
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v162 = v158 + v148
	if v162 == v145 {
		goto L64
	} else {
		goto L65
	}
L61:
	;
	v160 = F__emscripten_memcpy_bulkmem(m, v146, v148, v158)
	mBase = m.M
	v161 = v160
	goto L63
L62:
	;
	v161 = v146
	goto L63
L63:
	;
	goto L60
L64:
	;
	v164 = v140
	goto L66
L65:
	;
	v164 = v162
	goto L66
L66:
	;
	v165 = v161 + v158
	v167 = v150 - int32(1)
	if v167 != 0 {
		v146 = v165
		v148 = v164
		v150 = v167
		goto L57
	} else {
		goto L67
	}
L67:
	;
	goto L58
L68:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
	if v180&int32(1) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v204 = v168
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v129))) = (v204 - v129) << (uint(int32(2)) % 32)
	return v129
L71:
	;
	v183 = v93
	goto L73
L72:
	;
	v183 = v95
	goto L73
L73:
	;
	v184 = v168
	v186 = v183
	v187 = v111
	goto L74
L74:
	;
	v196 = F_pg_mblen_unbounded(m, v186)
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L76
	}
L75:
	;
	v204 = v201
	goto L70
L76:
	;
	if v196 != 0 {
		goto L78
	} else {
		goto L79
	}
L77:
	;
	v201 = v199 + v196
	v203 = v187 - int32(1)
	if v203 != 0 {
		v184 = v201
		v186 = v196 + v186
		v187 = v203
		goto L74
	} else {
		goto L81
	}
L78:
	;
	v198 = F__emscripten_memcpy_bulkmem(m, v184, v186, v196)
	mBase = m.M
	v199 = v198
	goto L80
L79:
	;
	v199 = v184
	goto L80
L80:
	;
	goto L77
L81:
	;
	goto L75
L82:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	F_errmsg(m, int32(416350), int32(0))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	F_errfinish(m, int32(512580), int32(206), int32(480922))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
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
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
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
				F_errmsg_internal(m, int32(489520), v6)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(515787), int32(808), int32(37209))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v33 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
			v38 = F_pq_getmsgtext(m, v8, v33-v34, v6+int32(12))
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v41 = F_parse_lquery(m, v38, int32(0))
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F_pfree(m, v38)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						m.G0 = v6 + int32(16)
						return v41
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
