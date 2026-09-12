package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__bt_binsrch_array_skey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v269 int32
	_ = v269
	v9 = int32(0)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v16 = v14 - int32(1)
	if l1 == v9 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(0)
	return v269
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v248
	return v254
L3:
	;
	if v130 < v132 {
		goto L52
	} else {
		goto L53
	}
L4:
	;
	v128 = int32(0)
	v130 = v9
	v131 = int32(-1)
	v132 = v16
	goto L3
L5:
	;
	goto L6
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l5)+12))
	if l2 == int32(1) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	v123 = int32(0)
	v125 = v20 - int32(2)
	if v125 < v123 {
		v248 = v76
		v254 = v123
		goto L2
	} else {
		goto L51
	}
L8:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v115
	return v118
L9:
	;
	v24 = v20 + int32(1)
	if v16 < v24 {
		goto L13
	} else {
		goto L14
	}
L10:
	;
	goto L11
L11:
	;
	v76 = int32(-1)
	v78 = v20 - int32(1)
	if v78 < int32(0) {
		v248 = v76
		v254 = v9
		goto L2
	} else {
		goto L34
	}
L12:
	;
	if v70 <= v16 {
		v128 = v69
		v130 = v70
		v131 = v71
		v132 = v16
		goto L3
	} else {
		goto L33
	}
L13:
	;
	v69 = int32(0)
	v70 = v24
	v71 = int32(-1)
	goto L12
L14:
	;
	goto L15
L15:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v30 = v28 & int32(1)
	if l4 != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	v69 = v65
	v70 = v20 + int32(2)
	v71 = v24
	goto L12
L17:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v24<<(uint(int32(2))%32))))
	v50 = F_FunctionCall2Coll(m, l0, v44, l3, v49)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(-1)
	return v24
L19:
	;
	if v30 != 0 {
		v269 = v24
		goto L1
	} else {
		goto L22
	}
L20:
	;
	goto L21
L21:
	;
	if v30 == int32(0) {
		goto L17
	} else {
		goto L24
	}
L22:
	;
	if v28&int32(33554432) != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	v65 = int32(1)
	goto L16
L24:
	;
	if v28&int32(33554432) == int32(0) {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v65 = int32(1)
	goto L16
L26:
	;
	return int32(0)
L27:
	;
	v54 = int32(1)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)))
	if v55&v54 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v50 < int32(0) {
		v65 = v54
		goto L16
	} else {
		goto L31
	}
L29:
	;
	v62 = v50
	goto L30
L30:
	;
	if v62 <= int32(0) {
		v115 = v62
		v118 = v24
		goto L8
	} else {
		goto L32
	}
L31:
	;
	v62 = int32(0) - v50
	goto L30
L32:
	;
	v65 = v62
	goto L16
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(1)
	return v16
L34:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v83 = v81 & int32(1)
	if l4 != 0 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	if v112 < int32(0) {
		v122 = v112
		goto L7
	} else {
		goto L50
	}
L36:
	;
	v112 = int32(0) - v97
	goto L35
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(1)
	return v78
L38:
	;
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v92 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v92+v78<<(uint(int32(2))%32))))
	v97 = F_FunctionCall2Coll(m, l0, v91, l3, v96)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L26
	} else {
		goto L47
	}
L39:
	;
	v122 = int32(-1)
	goto L7
L40:
	;
	if v83 != 0 {
		v269 = v78
		goto L1
	} else {
		goto L43
	}
L41:
	;
	goto L42
L42:
	;
	if v83 == int32(0) {
		goto L38
	} else {
		goto L45
	}
L43:
	;
	if v81&int32(33554432) != 0 {
		goto L39
	} else {
		goto L44
	}
L44:
	;
	goto L37
L45:
	;
	if v81&int32(33554432) != 0 {
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L39
L47:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)))
	if v99&int32(1) == int32(0) {
		v112 = v97
		goto L35
	} else {
		goto L48
	}
L48:
	;
	if int32(0) <= v97 {
		goto L36
	} else {
		goto L49
	}
L49:
	;
	goto L37
L50:
	;
	v115 = v112
	v118 = v78
	goto L8
L51:
	;
	v128 = v122
	v130 = v123
	v131 = v78
	v132 = v125
	goto L3
L52:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v138 = v135
	v144 = v130
	v146 = v132
	goto L55
L53:
	;
	v193 = v128
	v200 = v130
	v201 = v131
	goto L54
L54:
	;
	if v200 == v201 {
		goto L85
	} else {
		goto L86
	}
L55:
	;
	v149 = v138 & int32(1)
	v152 = base.I32_div_s(v146-v144, int32(2))
	v153 = v152 + v144
	if l4 != 0 {
		goto L58
	} else {
		goto L59
	}
L56:
	;
	v193 = v183
	v200 = v190
	v201 = v153
	goto L54
L57:
	;
	v186 = base.B2i32(int32(0) < v183)
	if int32(0) < v183 {
		goto L78
	} else {
		goto L79
	}
L58:
	;
	if v149 != 0 {
		v269 = v153
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	if v149 != 0 {
		goto L65
	} else {
		goto L66
	}
L61:
	;
	if v138&int32(33554432) != 0 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v158 = int32(-1)
	goto L64
L63:
	;
	v158 = int32(1)
	goto L64
L64:
	;
	v183 = v158
	v184 = v138
	goto L57
L65:
	;
	if v138&int32(33554432) != 0 {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	v164 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v165 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v165+v153<<(uint(int32(2))%32))))
	v170 = F_FunctionCall2Coll(m, l0, v164, l3, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L26
	} else {
		goto L71
	}
L68:
	;
	v163 = int32(1)
	goto L70
L69:
	;
	v163 = int32(-1)
	goto L70
L70:
	;
	v183 = v163
	v184 = v138
	goto L57
L71:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	if v172&int32(16777216) != 0 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v183 = int32(1)
	v184 = v172
	goto L57
L73:
	;
	if v170 < int32(0) {
		goto L72
	} else {
		goto L76
	}
L74:
	;
	v179 = v170
	goto L75
L75:
	;
	if v179 == int32(0) {
		v269 = v153
		goto L1
	} else {
		goto L77
	}
L76:
	;
	v179 = int32(0) - v170
	goto L75
L77:
	;
	v183 = v179
	v184 = v172
	goto L57
L78:
	;
	v187 = v146
	goto L80
L79:
	;
	v187 = v153
	goto L80
L80:
	;
	if int32(0) < v183 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v190 = v153 + int32(1)
	goto L83
L82:
	;
	v190 = v144
	goto L83
L83:
	;
	if v190 < v187 {
		v138 = v184
		v144 = v190
		v146 = v187
		goto L55
	} else {
		goto L84
	}
L84:
	;
	goto L56
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v193
	return v200
L86:
	;
	goto L87
L87:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v209 = v207 & int32(1)
	if l4 != 0 {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	if v209 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if v209 != 0 {
		goto L97
	} else {
		goto L98
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = int32(0)
	return v200
L92:
	;
	goto L93
L93:
	;
	if v207&int32(33554432) != 0 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v217 = int32(-1)
	goto L96
L95:
	;
	v217 = int32(1)
	goto L96
L96:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v217
	return v200
L97:
	;
	if v207&int32(33554432) != 0 {
		goto L100
	} else {
		goto L101
	}
L98:
	;
	goto L99
L99:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(l6)+12))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v228+v200<<(uint(int32(2))%32))))
	v233 = F_FunctionCall2Coll(m, l0, v227, l3, v232)
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L26
	} else {
		goto L103
	}
L100:
	;
	v224 = int32(1)
	goto L102
L101:
	;
	v224 = int32(-1)
	goto L102
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l7))) = v224
	return v200
L103:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l6)+3)))
	if v235&int32(1) == int32(0) {
		v248 = v233
		v254 = v200
		goto L2
	} else {
		goto L104
	}
L104:
	;
	v241 = int32(0)
	if v233 < v241 {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v245 = int32(1)
	goto L107
L106:
	;
	v245 = v241 - v233
	goto L107
L107:
	;
	v248 = v245
	v254 = v200
	goto L2
}
func F__bt_check_third_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v18 int32
	_ = v18
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+6)))
	v18 = (v12&int32(8191) + int32(7)) & int32(16376)
	if base.Ui32(v18) < base.Ui32(int32(2705)) {
		m.G0 = v10 + int32(48)
		return
	} else {
		if l2|base.B2i32(base.Ui32(int32(2712)) < base.Ui32(v18)) == int32(0) {
			m.G0 = v10 + int32(48)
			return
		} else {
			v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3)+16)))
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l3+v26)+12)))
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				if v28&int32(1) == int32(0) {
					v121 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v18
					*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v121 + int32(4)
					F_errmsg_internal(m, int32(675672), v10)
					mBase = m.M
					v128 = m.ExcPending
					if v128 != 0 {
						return
					} else {
						F_errfinish(m, int32(488920), int32(4229), int32(403270))
						mBase = m.M
						v133 = m.ExcPending
						if v133 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				} else {
					F_errcode(m, int32(261))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
						if l2 != 0 {
							v43 = int32(2704)
						} else {
							v43 = int32(2712)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v43
						if l2 != 0 {
							v47 = int32(4)
						} else {
							v47 = int32(3)
						}
						*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v47
						*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v18
						*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v40 + int32(4)
						F_errmsg(m, int32(672280), v10+int32(32))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+6)))
							if v58&int32(8192) == int32(0) {
								v86 = l4
							} else {
								v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+4)))
								if v63&int32(8192) == int32(0) {
									v68 = int32(0)
									if v63&int32(4096) == v68 {
										v84 = v68
										v86 = v84
									} else {
										v86 = l4 + v58&int32(8191) - int32(6)
									}
								} else {
									v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4)+2)))
									v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l4))))
									v84 = l4 + (v78 | v79<<(uint(int32(16))%32))
									v86 = v84
								}
							}
							v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+2)))
							v88 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86))))
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+48))
							v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v86)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v90
							*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v89 + int32(4)
							v95 = int32(16)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v87 | v88<<(uint(v95)%32)
							F_errdetail(m, int32(645634), v10+v95)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return
							} else {
								F_errhint(m, int32(605353), int32(0))
								mBase = m.M
								v107 = m.ExcPending
								if v107 != 0 {
									return
								} else {
									v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
									F_errtableconstraint(m, l1, v108+int32(4))
									mBase = m.M
									v112 = m.ExcPending
									if v112 != 0 {
										return
									} else {
										F_errfinish(m, int32(488920), int32(4245), int32(403270))
										mBase = m.M
										v117 = m.ExcPending
										if v117 != 0 {
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
	}
}
func F__bt_delitems_cmp(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	v3 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
	v4 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
	return v3 - v4
}
func F__bt_delitems_delete_check(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v166 int32
	_ = v166
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v294 int32
	_ = v294
	var v301 int32
	_ = v301
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v330 int32
	_ = v330
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v420 int32
	_ = v420
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v443 int32
	_ = v443
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v457 int32
	_ = v457
	var v463 int32
	_ = v463
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v530 int32
	_ = v530
	var v543 int32
	_ = v543
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v616 int32
	_ = v616
	var v617 int32
	_ = v617
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v654 int32
	_ = v654
	var v681 int32
	_ = v681
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v726 int32
	_ = v726
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v741 int32
	_ = v741
	var v746 int32
	_ = v746
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v781 int32
	_ = v781
	var v784 int32
	_ = v784
	var v787 int32
	_ = v787
	var v790 int32
	_ = v790
	var v794 int32
	_ = v794
	var v798 int32
	_ = v798
	var v803 int32
	_ = v803
	var v812 int32
	_ = v812
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v827 int64
	_ = v827
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v882 int32
	_ = v882
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	v5 = int32(0)
	v25 = m.G0
	v27 = v25 - int32(3296)
	m.G0 = v27
	if l1 < v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+188))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+76))
	v49 = m.T0[v48].(func(*base.Module, int32, int32) int32)(m, l2, l3)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v32+(l1^int32(-1))<<(uint(int32(2))%32))))
	v46 = v38
	goto L1
L3:
	;
	goto L4
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v46 = v40 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	return
L6:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if v52 < int32(2) {
		v73 = v5
		goto L7
	} else {
		goto L8
	}
L7:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	v76 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	F_pg_qsort(m, v76, v77, int32(8), int32(209))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L5
	} else {
		goto L14
	}
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v55)+118)))
	if v56 != int32(112) {
		v73 = v5
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(l2)+56))
	goto L10
L10:
	;
	if base.Ui32(v60) < base.Ui32(int32(12000)) {
		v73 = int32(1)
		goto L7
	} else {
		goto L11
	}
L11:
	;
	v63 = int32(0)
	v64 = *(*int32)(unsafe.Add(mBase, uint32(l2)+180))
	if v64 == v63 {
		v73 = v63
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(l2)+48))
	v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67)+119)))
	switch v68 - int32(109) {
	case 0, 5:
		goto L13
	default:
		v73 = v63
		goto L7
	}
L13:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v64)+104)))
	v73 = v71
	goto L7
L14:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v82 == int32(0) {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	m.G0 = v27 + int32(3296)
	return
L16:
	;
	if v82 <= int32(0) {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	if l1 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L18:
	;
	v400 = v5
	v403 = v5
	goto L17
L19:
	;
	goto L20
L20:
	;
	v97 = v5
	v98 = int32(0)
	v100 = v5
	v107 = v5
	goto L21
L21:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v119 = int32(*(*int16)(unsafe.Add(mBase, uint32(v115+v107<<(uint(int32(3))%32))+6)))
	v122 = v114 + v119*int32(6)
	v123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v122))))
	if v123 == v98&int32(65535) {
		v372 = v97
		v373 = v98
		v375 = v100
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v400 = v372
	v403 = v375
	goto L17
L23:
	;
	v390 = v107 + int32(1)
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v390 < v391 {
		v97 = v372
		v98 = v373
		v100 = v375
		v107 = v390
		goto L21
	} else {
		goto L63
	}
L24:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(v123<<(uint(int32(2))%32)+(v46+int32(24))-int32(4))))
	v135 = v46 + v132&int32(32767)
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135)+7)))
	if v136&int32(32) != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	v155 = v139 & int32(4095)
	if v155 == int32(0) {
		v348 = v97
		v351 = v100
		goto L31
	} else {
		goto L32
	}
L26:
	;
	v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+4)))
	if v139&int32(8192) != 0 {
		goto L25
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
	if v143 != int32(1) {
		v372 = v97
		v373 = v98
		v375 = v100
		goto L23
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v148 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(1648)+v100<<(uint(v148)%32)))) = uint16(v123)
	v372 = v97
	v373 = v98
	v375 = v100 + v148
	goto L23
L31:
	;
	v372 = v348
	v373 = v123
	v375 = v351
	goto L23
L32:
	;
	v162 = int32(0)
	v166 = v107
	v173 = v162
	v175 = v162
	goto L33
L33:
	;
	v188 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	if v188 <= v166 {
		v294 = v166
		v301 = v173
		goto L35
	} else {
		goto L36
	}
L34:
	;
	if v301 == int32(0) {
		v348 = v97
		v351 = v100
		goto L31
	} else {
		goto L58
	}
L35:
	;
	v317 = v175 + int32(1)
	if v317 != v155 {
		v166 = v294
		v173 = v301
		v175 = v317
		goto L33
	} else {
		goto L57
	}
L36:
	;
	v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135)+2)))
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v135))))
	v198 = v135 + (v190 | v191<<(uint(int32(16))%32)) + v175*int32(6)
	v202 = v166
	v206 = v188
	v214 = int32(-1)
	goto L37
L37:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l3)+24))
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l3)+20))
	v228 = v225 + v202<<(uint(int32(3))%32)
	v229 = int32(*(*int16)(unsafe.Add(mBase, uint32(v228)+6)))
	v232 = v224 + v229*int32(6)
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v232))))
	if v233 != v123 {
		v270 = v202
		v272 = v214
		goto L39
	} else {
		goto L40
	}
L38:
	;
	if v272 != 0 {
		v294 = v270
		v301 = v173
		goto L35
	} else {
		goto L51
	}
L39:
	;
	goto L38
L40:
	;
	v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v232)+2)))
	if v235 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+2)))
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228))))
	v243 = int32(16)
	v245 = v241 | v242<<(uint(v243)%32)
	v246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198)+2)))
	v247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198))))
	v250 = v246 | v247<<(uint(v243)%32)
	if base.Ui32(v245) < base.Ui32(v250) {
		v261 = int32(-1)
		goto L45
	} else {
		goto L46
	}
L42:
	;
	v265 = v206
	v266 = v214
	goto L43
L43:
	;
	v268 = v202 + int32(1)
	if v268 < v265 {
		v202 = v268
		v206 = v265
		v214 = v266
		goto L37
	} else {
		goto L50
	}
L44:
	;
	if int32(0) <= v261 {
		v270 = v202
		v272 = v261
		goto L39
	} else {
		goto L49
	}
L45:
	;
	goto L44
L46:
	;
	if base.Ui32(v250) < base.Ui32(v245) {
		v261 = int32(1)
		goto L45
	} else {
		goto L47
	}
L47:
	;
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+4)))
	v256 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198)+4)))
	if base.Ui32(v255) < base.Ui32(v256) {
		v261 = int32(-1)
		goto L45
	} else {
		goto L48
	}
L48:
	;
	v261 = base.B2i32(base.Ui32(v256) < base.Ui32(v255))
	goto L45
L49:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(l3)+16))
	v265 = v264
	v266 = v261
	goto L43
L50:
	;
	v270 = v268
	v272 = v266
	goto L39
L51:
	;
	if v173 != 0 {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	v283 = int32(1)
	v284 = v281 + v283
	*(*uint16)(unsafe.Add(mBase, uint32(v282)+6)) = uint16(v284)
	*(*uint16)(unsafe.Add(mBase, uint32(v282+v281&int32(65535)<<(uint(v283)%32))+8)) = uint16(v175)
	v294 = v270
	v301 = v282
	goto L35
L53:
	;
	v273 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v173)+6)))
	v281 = v273
	v282 = v173
	goto L52
L54:
	;
	goto L55
L55:
	;
	v275 = F_palloc(m, v155<<(uint(int32(1))%32)+int32(8))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L5
	} else {
		goto L56
	}
L56:
	;
	v277 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v275)+6)) = uint16(v277)
	*(*uint16)(unsafe.Add(mBase, uint32(v275)+4)) = uint16(v123)
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v135
	v281 = int32(0)
	v282 = v275
	goto L52
L57:
	;
	goto L34
L58:
	;
	v321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v301)+6)))
	if v155 == v321 {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(1648)+v100<<(uint(int32(1))%32)))) = uint16(v123)
	F_pfree(m, v301)
	mBase = m.M
	v330 = m.ExcPending
	if v330 != 0 {
		goto L5
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v97<<(uint(int32(2))%32)))) = v301
	v348 = v97 + int32(1)
	v351 = v100
	goto L31
L62:
	;
	v348 = v97
	v351 = v100 + int32(1)
	goto L31
L63:
	;
	goto L22
L64:
	;
	v435 = int32(0)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v438 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v437)+118)))
	if v438 != int32(112) {
		v451 = v435
		goto L68
	} else {
		goto L69
	}
L65:
	;
	v420 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v420+(l1^int32(-1))<<(uint(int32(2))%32))))
	v434 = v426
	goto L64
L66:
	;
	goto L67
L67:
	;
	v428 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v434 = v428 + l1<<(uint(int32(13))%32) + int32(-8192)
	goto L64
L68:
	;
	if int32(0) < v400 {
		goto L74
	} else {
		goto L75
	}
L69:
	;
	v443 = *(*int32)(unsafe.Add(mBase, _consts[10]))
	if int32(0) < v443 {
		v451 = int32(1)
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v447 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v447 != 0 {
		v451 = int32(0)
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v451 = base.B2i32(v448 == int32(0))
	goto L68
L72:
	;
	if int32(0) < v403 {
		goto L118
	} else {
		goto L119
	}
L73:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L5
	} else {
		goto L111
	}
L74:
	;
	v457 = int32(0)
	v463 = v435
	goto L77
L75:
	;
	goto L76
L76:
	;
	v703 = int32(0)
	v704 = int32(4474964)
	v706 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v706 + int32(1)
	v751 = v703
	v753 = v703
	goto L72
L77:
	;
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v457<<(uint(int32(2))%32))))
	F__bt_update_posting(m, v484)
	mBase = m.M
	v486 = m.ExcPending
	if v486 != 0 {
		goto L5
	} else {
		goto L79
	}
L78:
	;
	v503 = int32(0)
	if v451 != 0 {
		goto L81
	} else {
		goto L82
	}
L79:
	;
	v487 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v484)+6)))
	v490 = int32(1)
	v493 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v484)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(2480)+v457<<(uint(v490)%32)))) = uint16(v493)
	v499 = v463 + v487<<(uint(v490)%32) + int32(2)
	v501 = v457 + v490
	if v501 != v400 {
		v457 = v501
		v463 = v499
		goto L77
	} else {
		goto L80
	}
L80:
	;
	goto L78
L81:
	;
	v508 = int32(0)
	v509 = F_palloc(m, v499)
	mBase = m.M
	v510 = m.ExcPending
	if v510 != 0 {
		goto L5
	} else {
		goto L84
	}
L82:
	;
	v626 = v503
	v628 = v503
	goto L83
L83:
	;
	v646 = int32(4474964)
	v648 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v648 + int32(1)
	v654 = v503
	goto L106
L84:
	;
	if v400 != int32(1) {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v519 = v508
	v520 = v503
	v530 = int32(0)
	goto L88
L86:
	;
	v581 = v508
	v582 = v503
	goto L87
L87:
	;
	if v400&int32(1) != 0 {
		goto L99
	} else {
		goto L100
	}
L88:
	;
	v543 = int32(2)
	v545 = v27 + int32(16) + v519<<(uint(v543)%32)
	v546 = *(*int32)(unsafe.Add(mBase, uint32(v545)))
	v547 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v546)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v520+v509))) = uint16(v547)
	v550 = v520 + v543
	v555 = v547 << (uint(int32(1)) % 32)
	if v555 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v581 = v574
	v582 = v572
	goto L87
L90:
	;
	v558 = v550 + v555
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v545)+4))
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v560)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v509+v558))) = uint16(v561)
	v564 = v558 + int32(2)
	v569 = v561 << (uint(int32(1)) % 32)
	if v569 != 0 {
		goto L95
	} else {
		goto L96
	}
L91:
	;
	v556 = F__emscripten_memcpy_bulkmem(m, v509+v550, v546+int32(8), v555)
	mBase = m.M
	goto L93
L92:
	;
	goto L93
L93:
	;
	goto L90
L94:
	;
	v572 = v564 + v569
	v573 = int32(2)
	v574 = v519 + v573
	v576 = v530 + v573
	if v576 != v400&int32(2147483646) {
		v519 = v574
		v520 = v572
		v530 = v576
		goto L88
	} else {
		goto L98
	}
L95:
	;
	v570 = F__emscripten_memcpy_bulkmem(m, v509+v564, v560+int32(8), v569)
	mBase = m.M
	goto L97
L96:
	;
	goto L97
L97:
	;
	goto L94
L98:
	;
	goto L89
L99:
	;
	v602 = v582 + v509
	v605 = int32(2)
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v581<<(uint(v605)%32))))
	v609 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v608)+6)))
	*(*uint16)(unsafe.Add(mBase, uint32(v602))) = uint16(v609)
	v616 = v609 << (uint(int32(1)) % 32)
	if v616 != 0 {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L101
L101:
	;
	v626 = v499
	v628 = v509
	goto L83
L102:
	;
	goto L101
L103:
	;
	v617 = F__emscripten_memcpy_bulkmem(m, v602+v605, v608+int32(8), v616)
	mBase = m.M
	goto L105
L104:
	;
	goto L105
L105:
	;
	goto L102
L106:
	;
	v681 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27+int32(2480)+v654<<(uint(int32(1))%32)))))
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v654<<(uint(int32(2))%32))))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v687)))
	v689 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v688)+6)))
	v696 = F_PageIndexTupleOverwrite(m, v434, v681, v688, (v689&int32(8191)+int32(7))&int32(16376))
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L5
	} else {
		goto L108
	}
L107:
	;
	v751 = v626
	v753 = v628
	goto L72
L108:
	;
	if v696 == int32(0) {
		goto L73
	} else {
		goto L109
	}
L109:
	;
	v701 = v654 + int32(1)
	if v701 != v400 {
		v654 = v701
		goto L106
	} else {
		goto L110
	}
L110:
	;
	goto L107
L111:
	;
	if l1 < int32(0) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v734 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v27))) = v733
	*(*int32)(unsafe.Add(mBase, uint32(v27)+4)) = v734 + int32(4)
	F_errmsg_internal(m, int32(674726), v27)
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L5
	} else {
		goto L116
	}
L113:
	;
	v718 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v724 = *(*int32)(unsafe.Add(mBase, uint32(v718+(l1^int32(-1))<<(uint(int32(6))%32))+16))
	v733 = v724
	goto L112
L114:
	;
	goto L115
L115:
	;
	v726 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v732 = *(*int32)(unsafe.Add(mBase, uint32(v726+l1<<(uint(int32(6))%32)+int32(-64))+16))
	v733 = v732
	goto L112
L116:
	;
	F_errfinish(m, int32(494298), int32(1320), int32(347667))
	mBase = m.M
	v746 = m.ExcPending
	if v746 != 0 {
		goto L5
	} else {
		goto L117
	}
L117:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L118:
	;
	F_PageIndexMultiDelete(m, v434, v27+int32(1648), v403)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L5
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v777 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v434)+16)))
	v778 = v434 + v777
	v779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v778)+12)))
	v781 = v779 & int32(65471)
	*(*uint16)(unsafe.Add(mBase, uint32(v778)+12)) = uint16(v781)
	F_MarkBufferDirty(m, l1)
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L5
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	if v451 != 0 {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v27)+2476)) = uint8(v73)
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+2474)) = uint16(v400)
	v787 = int32(0)
	if v787 < v75 {
		goto L126
	} else {
		goto L127
	}
L124:
	;
	goto L125
L125:
	;
	v833 = int32(4474964)
	v835 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v835 - int32(1)
	if v753 != 0 {
		goto L142
	} else {
		goto L143
	}
L126:
	;
	v790 = v49
	goto L128
L127:
	;
	v790 = v787
	goto L128
L128:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v27)+2468)) = v790
	*(*uint16)(unsafe.Add(mBase, uint32(v27)+2472)) = uint16(v403)
	F_XLogBeginInsert(m)
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L5
	} else {
		goto L129
	}
L129:
	;
	F_XLogRegisterBuffer(m, int32(0), l1, int32(8))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L5
	} else {
		goto L130
	}
L130:
	;
	F_XLogRegisterData(m, v27+int32(2468), int32(9))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L5
	} else {
		goto L131
	}
L131:
	;
	if int32(0) < v403 {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	F_XLogRegisterBufData(m, int32(0), v27+int32(1648), v403<<(uint(int32(1))%32))
	mBase = m.M
	v812 = m.ExcPending
	if v812 != 0 {
		goto L5
	} else {
		goto L135
	}
L133:
	;
	goto L134
L134:
	;
	if int32(0) < v400 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	goto L134
L136:
	;
	F_XLogRegisterBufData(m, int32(0), v27+int32(2480), v400<<(uint(int32(1))%32))
	mBase = m.M
	v821 = m.ExcPending
	if v821 != 0 {
		goto L5
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v827 = F_XLogInsert(m, int32(11), int32(112))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L5
	} else {
		goto L141
	}
L139:
	;
	F_XLogRegisterBufData(m, int32(0), v753, v751)
	mBase = m.M
	v824 = m.ExcPending
	if v824 != 0 {
		goto L5
	} else {
		goto L140
	}
L140:
	;
	goto L138
L141:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v434))) = base.I64_rotr(v827, int64(32))
	goto L125
L142:
	;
	F_pfree(m, v753)
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L5
	} else {
		goto L145
	}
L143:
	;
	goto L144
L144:
	;
	if v400 <= int32(0) {
		goto L15
	} else {
		goto L146
	}
L145:
	;
	goto L144
L146:
	;
	v845 = int32(0)
	goto L147
L147:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v845<<(uint(int32(2))%32))))
	v873 = *(*int32)(unsafe.Add(mBase, uint32(v872)))
	F_pfree(m, v873)
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L5
	} else {
		goto L149
	}
L148:
	;
	v882 = int32(0)
	goto L151
L149:
	;
	v877 = v845 + int32(1)
	if v877 != v400 {
		v845 = v877
		goto L147
	} else {
		goto L150
	}
L150:
	;
	goto L148
L151:
	;
	v909 = *(*int32)(unsafe.Add(mBase, uint32(v27+int32(16)+v882<<(uint(int32(2))%32))))
	F_pfree(m, v909)
	mBase = m.M
	v911 = m.ExcPending
	if v911 != 0 {
		goto L5
	} else {
		goto L153
	}
L152:
	;
	goto L15
L153:
	;
	v913 = v882 + int32(1)
	if v913 != v400 {
		v882 = v913
		goto L151
	} else {
		goto L154
	}
L154:
	;
	goto L152
}
func F__bt_killitems(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
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
	var v36 int64
	_ = v36
	var v37 int32
	_ = v37
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v296 int32
	_ = v296
	var v306 int32
	_ = v306
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v403 int32
	_ = v403
	v2 = int32(0)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v21)+36)) = v2
	v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)))
	if v25 == v2 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F__bt_relbuf(m, v385)
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L6
	} else {
		goto L68
	}
L2:
	;
	if v40 < int32(0) {
		goto L12
	} else {
		goto L13
	}
L3:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	F__bt_lockbuf(m, v28, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	goto L5
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v21)+60))
	v34 = F__bt_getbuf(m, v20, v32, int32(1))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L6
	} else {
		goto L8
	}
L6:
	;
	return
L7:
	;
	v40 = v28
	goto L2
L8:
	;
	v36 = F_BufferGetLSNAtomic(m, v34)
	mBase = m.M
	v37 = m.ExcPending
	if v37 != 0 {
		goto L6
	} else {
		goto L9
	}
L9:
	;
	v38 = *(*int64)(unsafe.Add(mBase, uint32(v21)+72))
	if v36 != v38 {
		v385 = v34
		goto L1
	} else {
		goto L10
	}
L10:
	;
	v40 = v34
	goto L2
L11:
	;
	if v22 <= int32(0) {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v44 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v44+(v40^int32(-1))<<(uint(int32(2))%32))))
	v58 = v50
	goto L11
L13:
	;
	goto L14
L14:
	;
	v52 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v58 = v52 + v40<<(uint(int32(13))%32) + int32(-8192)
	goto L11
L15:
	;
	v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)))
	if v380 != 0 {
		v385 = v40
		goto L1
	} else {
		goto L66
	}
L16:
	;
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+16)))
	v64 = v58 + v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v64)+4))
	if v65 != 0 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v66 = int32(2)
	goto L19
L18:
	;
	v66 = int32(1)
	goto L19
L19:
	;
	v70 = v21 + int32(104)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v58)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v71) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v79 = int32(base.Ui32(v71+int32(262120)) >> (uint(int32(2)) % 32))
	goto L22
L21:
	;
	v79 = int32(0)
	goto L22
L22:
	;
	v81 = v79 & int32(65535)
	v87 = v2
	v93 = v2
	goto L23
L23:
	;
	v102 = v87 + int32(1)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v103+v87<<(uint(int32(2))%32))))
	v110 = v70 + v107*int32(10)
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+6)))
	if base.Ui32(v111) < base.Ui32(v66) {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v64)+12)))
	v356 = v354 | int32(64)
	*(*uint16)(unsafe.Add(mBase, uint32(v64)+12)) = uint16(v356)
	F_MarkBufferDirtyHint(m, v40, int32(1))
	mBase = m.M
	v360 = m.ExcPending
	if v360 != 0 {
		goto L6
	} else {
		goto L65
	}
L25:
	;
	goto L24
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v139))) = v281 | int32(98304)
	if v102 != v22 {
		v87 = v102
		v93 = int32(1)
		goto L23
	} else {
		goto L64
	}
L27:
	;
	if v102 != v22 {
		v87 = v102
		goto L23
	} else {
		goto L62
	}
L28:
	;
	if base.Ui32(v81) < base.Ui32(v111) {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	v122 = v111
	v124 = v110
	goto L30
L30:
	;
	v139 = v122&int32(65535)<<(uint(int32(2))%32) + (v58 + int32(24)) - int32(4)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v143 = v58 + v140&int32(32767)
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143)+7)))
	if v144&int32(32) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L31:
	;
	goto L27
L32:
	;
	v306 = v122 + int32(1)
	if base.Ui32(v306&int32(65535)) <= base.Ui32(v81) {
		v122 = v306
		v124 = v296
		goto L30
	} else {
		goto L61
	}
L33:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v139)))
	v282 = int32(98304)
	if v281&v282 != v282 {
		goto L26
	} else {
		goto L60
	}
L34:
	;
	if v251 != v156 {
		v296 = v252
		goto L32
	} else {
		goto L59
	}
L35:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+2)))
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143))))
	v225 = int32(16)
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+2)))
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124))))
	if v223|v224<<(uint(v225)%32) == v228|v229<<(uint(v225)%32) {
		goto L54
	} else {
		goto L55
	}
L36:
	;
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+4)))
	if v149&int32(8192) == int32(0) {
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v154 = int32(0)
	v156 = v149 & int32(4095)
	if v156 == v154 {
		v251 = v154
		v252 = v124
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v162 = v102
	v168 = v154
	v169 = v124
	goto L39
L39:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+2)))
	v179 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143))))
	v180 = int32(16)
	v186 = v143 + (v178 | v179<<(uint(v180)%32)) + v168*int32(6)
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+2)))
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186))))
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+2)))
	v193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169))))
	if v187|v188<<(uint(v180)%32) == v192|v193<<(uint(v180)%32) {
		goto L43
	} else {
		goto L44
	}
L40:
	;
	v272 = v218
	goto L33
L41:
	;
	if v203 == int32(0) {
		v251 = v168
		v252 = v169
		goto L34
	} else {
		goto L47
	}
L42:
	;
	goto L41
L43:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v186)+4)))
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v169)+4)))
	if v199 == v200 {
		v203 = int32(1)
		goto L42
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v203 = int32(0)
	goto L42
L46:
	;
	goto L45
L47:
	;
	if v162 < v22 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v207+v162<<(uint(int32(2))%32))))
	v217 = v162 + int32(1)
	v218 = v70 + v211*int32(10)
	goto L50
L49:
	;
	v217 = v162
	v218 = v169
	goto L50
L50:
	;
	v220 = v168 + int32(1)
	if v220 != v156 {
		v162 = v217
		v168 = v220
		v169 = v218
		goto L39
	} else {
		goto L51
	}
L51:
	;
	goto L40
L52:
	;
	if v239 == int32(0) {
		v296 = v124
		goto L32
	} else {
		goto L58
	}
L53:
	;
	goto L52
L54:
	;
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+4)))
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v124)+4)))
	if v235 == v236 {
		v239 = int32(1)
		goto L53
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	v239 = int32(0)
	goto L53
L57:
	;
	goto L56
L58:
	;
	v272 = v124
	goto L33
L59:
	;
	v272 = v252
	goto L33
L60:
	;
	v296 = v272
	goto L32
L61:
	;
	goto L31
L62:
	;
	if v93 != 0 {
		goto L25
	} else {
		goto L63
	}
L63:
	;
	goto L15
L64:
	;
	goto L25
L65:
	;
	goto L15
L66:
	;
	F__bt_unlockbuf(m, v40)
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L6
	} else {
		goto L67
	}
L67:
	;
	return
L68:
	;
	return
}
func F__bt_lockbuf(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	F_LockBuffer(m, l0, l1)
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
func F__bt_mkscankey(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v114 int32
	_ = v114
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+224))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v23 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+10)))
	if l1 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v63 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+4)) = uint8(v63)
	*(*uint16)(unsafe.Add(mBase, uint32(v60)+2)) = uint16(v63)
	if v62 < v23 {
		goto L13
	} else {
		goto L14
	}
L2:
	;
	v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+7)))
	if v24&int32(32) == int32(0) {
		goto L6
	} else {
		goto L7
	}
L3:
	;
	goto L4
L4:
	;
	v56 = F_palloc(m, v23*int32(48)+int32(16))
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L9
	} else {
		goto L12
	}
L5:
	;
	v42 = F_palloc(m, v23*int32(48)+int32(16))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L9
	} else {
		goto L10
	}
L6:
	;
	v35 = int32(*(*int16)(unsafe.Add(mBase, uint32(v22)+8)))
	v37 = v35
	goto L5
L7:
	;
	v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v29&int32(8192) != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v37 = v29 & int32(4095)
	goto L5
L9:
	;
	return int32(0)
L10:
	;
	F__bt_metaversion(m, l0, v42, v42+int32(1))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L9
	} else {
		goto L11
	}
L11:
	;
	v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
	v60 = v42
	v61 = v50
	v62 = v37
	goto L1
L12:
	;
	v58 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v56))) = uint16(v58)
	v60 = v56
	v61 = int32(1)
	v62 = int32(0)
	goto L1
L13:
	;
	v69 = v62
	goto L15
L14:
	;
	v69 = v23
	goto L15
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+12)) = v69
	if l1 == int32(0) {
		v103 = v63
		goto L16
	} else {
		goto L17
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v60)+8)) = v103
	if int32(0) < v23 {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	if v61&int32(1) == int32(0) {
		v103 = v63
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v77 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+6)))
	if v77&int32(8192) == int32(0) {
		v103 = l1
		goto L16
	} else {
		goto L19
	}
L19:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if v82&int32(8192) == int32(0) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v87 = int32(0)
	if v82&int32(4096) == v87 {
		v103 = v87
		goto L16
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+2)))
	v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1))))
	v103 = l1 + (v97 | v98<<(uint(int32(16))%32))
	goto L16
L23:
	;
	v103 = l1 + v77&int32(8191) - int32(6)
	goto L16
L24:
	;
	v114 = int32(0)
	goto L27
L25:
	;
	goto L26
L26:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+13)))
	if v194 == int32(1) {
		goto L39
	} else {
		goto L40
	}
L27:
	;
	v127 = int32(1)
	v129 = v114 + v127
	v130 = base.I32_extend16_s(v129)
	v132 = F_index_getprocinfo(m, l0, v130, v127)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L9
	} else {
		goto L29
	}
L28:
	;
	goto L26
L29:
	;
	if v114 < v62 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v138 = F_index_getattr_2(m, l1, v129, v21, v18+int32(15))
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L9
	} else {
		goto L33
	}
L31:
	;
	v141 = v127
	v142 = int32(0)
	goto L32
L32:
	;
	v145 = v60 + int32(16) + v114*int32(48)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20+v114<<(uint(int32(1))%32)))))
	v155 = int32(0)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+248))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v156+v114<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v145)+44)) = v142
	*(*int32)(unsafe.Add(mBase, uint32(v145)+12)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v145)+8)) = v155
	*(*uint16)(unsafe.Add(mBase, uint32(v145)+6)) = uint16(v155)
	*(*uint16)(unsafe.Add(mBase, uint32(v145)+4)) = uint16(v130)
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v141&int32(255) | v151<<(uint(int32(24))%32)
	v171 = *(*int32)(unsafe.Add(mBase, _consts[0]))
	F_fmgr_info_copy(m, v145+int32(16), v132, v171)
	mBase = m.M
	goto L34
L33:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+15)))
	v141 = v140
	v142 = v138
	goto L32
L34:
	;
	if v141&int32(1) != 0 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v175 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v175)
	goto L37
L36:
	;
	goto L37
L37:
	;
	if v129 != v23 {
		v114 = v129
		goto L27
	} else {
		goto L38
	}
L38:
	;
	goto L28
L39:
	;
	v197 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)) = uint8(v197)
	goto L41
L40:
	;
	goto L41
L41:
	;
	m.G0 = v18 + int32(16)
	return v60
}
func F__bt_parallel_seize(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v80 int32
	_ = v80
	var v87 int32
	_ = v87
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
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	v5 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(16)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v22
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v22
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(-4294967296)
	v28 = int32(257)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+88)) = uint16(v28)
	if l3 != 0 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v18 + int32(16)
	return v244
L2:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v20)+28))
	v36 = v20 + v35
	v40 = v36 + int32(40)
	v42 = v36 + int32(12)
	goto L10
L3:
	;
	v30 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+19)) = uint8(v30)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+17)) = uint16(v30)
	goto L2
L4:
	;
	goto L5
L5:
	;
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+17)))
	if v34 != 0 {
		v244 = v5
		goto L1
	} else {
		goto L6
	}
L6:
	;
	goto L2
L7:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L13
	} else {
		goto L52
	}
L8:
	;
	F_LWLockRelease(m, v42)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L13
	} else {
		goto L39
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v195
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v36)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v199
	v203 = int32(1)
	v204 = v5
	goto L8
L10:
	;
	v59 = F_LWLockAcquire(m, v42, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L13
	} else {
		goto L14
	}
L11:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v195 = v194
	goto L9
L12:
	;
	goto L11
L13:
	;
	return int32(0)
L14:
	;
	v63 = int32(0)
	v64 = int32(1)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	switch v66 - v64 {
	case 0:
		goto L16
	case 1:
		v176 = v64
		v177 = v63
		goto L15
	case 2:
		goto L17
	case 3:
		v203 = v63
		v204 = v5
		goto L8
	default:
		goto L12
	}
L15:
	;
	F_LWLockRelease(m, v42)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L13
	} else {
		goto L35
	}
L16:
	;
	if l3 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	if v69 != 0 {
		v195 = v69
		goto L9
	} else {
		goto L18
	}
L18:
	;
	v203 = v63
	v204 = int32(1)
	goto L8
L19:
	;
	v168 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+19)) = uint8(v168)
	v170 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v21)+17)) = uint16(v170)
	v176 = l3
	v177 = l3
	goto L15
L20:
	;
	v73 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v36)+8)) = v73
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v40 + v75<<(uint(v73)%32)
	v80 = int32(0)
	if v75 <= v80 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v87 = v80
	goto L22
L22:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v21)+20))
	v102 = v99 + v87<<(uint(int32(5))%32)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v102)))
	v106 = v98 + v103*int32(48)
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v102)+4))
	if v107 != int32(-1) {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L19
L24:
	;
	v150 = v87 + int32(1)
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v150 < v151 {
		v87 = v150
		goto L22
	} else {
		goto L34
	}
L25:
	;
	v110 = int32(2)
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v40+v87<<(uint(v110)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v102)+12)) = v113
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v102)+8))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v113<<(uint(v110)%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v119
	goto L24
L26:
	;
	goto L27
L27:
	;
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v102)+18)))
	if v121 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = int32(0)
	v130 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
	*(*int32)(unsafe.Add(mBase, uint32(v106))) = v131
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v133 + int32(4)
	v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
	if v137&int32(24) != 0 {
		goto L24
	} else {
		goto L32
	}
L29:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v106)+44))
	if v122 == int32(0) {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	F_pfree(m, v122)
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L13
	} else {
		goto L31
	}
L31:
	;
	goto L28
L32:
	;
	v144 = F_datumRestore(m, v18+int32(12), v18+int32(11))
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L13
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v106)+44)) = v144
	goto L24
L34:
	;
	goto L23
L35:
	;
	if v177 != 0 {
		goto L7
	} else {
		goto L36
	}
L36:
	;
	if v176 == int32(0) {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	F_ConditionVariableSleep(m, v36+int32(28), int32(134217735))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L13
	} else {
		goto L38
	}
L38:
	;
	goto L10
L39:
	;
	F_ConditionVariableCancelSleep(m)
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	if v204 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v244 = v203
	goto L1
L42:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v211 == int32(0) {
		goto L41
	} else {
		goto L43
	}
L43:
	;
	v214 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v214)+17)))
	if v215 != 0 {
		goto L41
	} else {
		goto L44
	}
L44:
	;
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v211)+28))
	v217 = v211 + v216
	v219 = v217 + int32(12)
	v221 = F_LWLockAcquire(m, v219, int32(0))
	mBase = m.M
	v222 = m.ExcPending
	if v222 != 0 {
		goto L13
	} else {
		goto L45
	}
L45:
	;
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v217)+8))
	if v223 != int32(4) {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v217)+8)) = int32(4)
	F_LWLockRelease(m, v219)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L13
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	F_LWLockRelease(m, v219)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L13
	} else {
		goto L51
	}
L49:
	;
	F_ConditionVariableBroadcast(m, v217+int32(28))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L13
	} else {
		goto L50
	}
L50:
	;
	goto L41
L51:
	;
	goto L41
L52:
	;
	v244 = v176
	goto L1
}
func F__bt_readnextpage(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v39 int32
	_ = v39
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
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
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v152 int32
	_ = v152
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v195 int32
	_ = v195
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v355 int32
	_ = v355
	var v368 int32
	_ = v368
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v394 int32
	_ = v394
	var v397 int32
	_ = v397
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v411 int32
	_ = v411
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
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
	var v434 int32
	_ = v434
	var v438 int32
	_ = v438
	var v439 int64
	_ = v439
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v479 int32
	_ = v479
	var v495 int32
	_ = v495
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = l1
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l3 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L2:
	;
	v25 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+88)) = uint8(v25)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v27 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+89)) = uint8(v27)
	goto L1
L5:
	;
	m.G0 = v17 + int32(32)
	return v495
L6:
	;
	F__bt_parallel_done(m, l0)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L20
	} else {
		goto L125
	}
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(-4294967296)
	goto L6
L8:
	;
	v39 = l4
	goto L9
L9:
	;
	v50 = base.B2i32(l3 != int32(1))
	if v50 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	v434 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+40)))
	if v434 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L11:
	;
	v60 = v39 & int32(1)
	if v60 != 0 {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+89)))
	if v53 == int32(0) {
		goto L7
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+88)))
	if v56 == int32(0) {
		goto L7
	} else {
		goto L16
	}
L15:
	;
	goto L11
L16:
	;
	goto L11
L17:
	;
	if v50 == int32(0) {
		goto L24
	} else {
		goto L25
	}
L18:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v61 == int32(0) {
		goto L17
	} else {
		goto L19
	}
L19:
	;
	v69 = F__bt_parallel_seize(m, l0, v17+int32(28), v17+int32(24), int32(0))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return int32(0)
L21:
	;
	if v69 != 0 {
		goto L17
	} else {
		goto L22
	}
L22:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(-4294967296)
	v495 = int32(0)
	goto L5
L23:
	;
	if v355 < int32(0) {
		goto L93
	} else {
		goto L94
	}
L24:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v79 != 0 {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	goto L26
L26:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v17)+24))
	v96 = v88
	v97 = v87
	goto L35
L27:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L20
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v84 = F__bt_getbuf(m, v22, v82, int32(1))
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L20
	} else {
		goto L31
	}
L30:
	;
	goto L29
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v84
	v355 = v84
	goto L23
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L20
	} else {
		goto L89
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L20
	} else {
		goto L86
	}
L34:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+56)) = int64(-4294967296)
	goto L6
L35:
	;
	v104 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v104 != 0 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	F__bt_relbuf(m, v284)
	mBase = m.M
	v299 = m.ExcPending
	if v299 != 0 {
		goto L20
	} else {
		goto L85
	}
L37:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L20
	} else {
		goto L40
	}
L38:
	;
	goto L39
L39:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	v109 = F__bt_getbuf(m, v22, v107, int32(1))
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L20
	} else {
		goto L42
	}
L40:
	;
	goto L39
L41:
	;
	v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v128)+16)))
	v130 = v129 + v128
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v130)+4))
	v132 = int32(0)
	v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130)+12)))
	if base.B2i32(v133&int32(4) == v132)&base.B2i32(v131 == v96) == v132 {
		goto L47
	} else {
		goto L48
	}
L42:
	;
	if v109 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v114 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v114+(v109^int32(-1))<<(uint(int32(2))%32))))
	v128 = v120
	goto L41
L44:
	;
	goto L45
L45:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v128 = v122 + v109<<(uint(int32(13))%32) + int32(-8192)
	goto L41
L46:
	;
	v209 = F__bt_relandgetbuf(m, v22, v146, v96, int32(1))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L20
	} else {
		goto L64
	}
L47:
	;
	v143 = v131
	v146 = v109
	v152 = v132
	goto L50
L48:
	;
	v195 = v109
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = v195
	if v195 == int32(0) {
		goto L34
	} else {
		goto L61
	}
L50:
	;
	if v143 == int32(0) {
		goto L46
	} else {
		goto L52
	}
L51:
	;
	v195 = v162
	goto L49
L52:
	;
	if v152 == int32(4) {
		goto L46
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v143
	v162 = F__bt_relandgetbuf(m, v22, v146, v143, int32(1))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L20
	} else {
		goto L55
	}
L54:
	;
	v183 = v152 + int32(1)
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
	v185 = v181 + v184
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v185)+4))
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v185)+12)))
	if v187&int32(4) != 0 {
		v143 = v186
		v146 = v162
		v152 = v183
		goto L50
	} else {
		goto L59
	}
L55:
	;
	if v162 < int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v167 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(v167+(v162^int32(-1))<<(uint(int32(2))%32))))
	v181 = v173
	goto L54
L57:
	;
	goto L58
L58:
	;
	v175 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v181 = v175 + v162<<(uint(int32(13))%32) + int32(-8192)
	goto L54
L59:
	;
	if v186 != v96 {
		v143 = v186
		v146 = v162
		v152 = v183
		goto L50
	} else {
		goto L60
	}
L60:
	;
	goto L51
L61:
	;
	v355 = v195
	goto L23
L62:
	;
	if v294 != 0 {
		goto L81
	} else {
		goto L82
	}
L63:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+16)))
	v230 = v229 + v228
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v230)+12)))
	if v231&int32(4) != 0 {
		goto L68
	} else {
		goto L69
	}
L64:
	;
	if v209 < int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v214+(v209^int32(-1))<<(uint(int32(2))%32))))
	v228 = v220
	goto L63
L66:
	;
	goto L67
L67:
	;
	v222 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v228 = v222 + v209<<(uint(int32(13))%32) + int32(-8192)
	goto L63
L68:
	;
	v235 = v230
	v238 = v209
	goto L71
L69:
	;
	goto L70
L70:
	;
	v278 = *(*int32)(unsafe.Add(mBase, uint32(v230)))
	if v278 == v97 {
		goto L32
	} else {
		goto L80
	}
L71:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v235)+4))
	if v248 == int32(0) {
		goto L33
	} else {
		goto L73
	}
L72:
	;
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v273)))
	v284 = v252
	v287 = v248
	v294 = v277
	goto L62
L73:
	;
	v252 = F__bt_relandgetbuf(m, v22, v238, v248, int32(1))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L20
	} else {
		goto L75
	}
L74:
	;
	v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v271)+16)))
	v273 = v272 + v271
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v273)+12)))
	if v274&int32(4) != 0 {
		v235 = v273
		v238 = v252
		goto L71
	} else {
		goto L79
	}
L75:
	;
	if v252 < int32(0) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v257 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v257+(v252^int32(-1))<<(uint(int32(2))%32))))
	v271 = v263
	goto L74
L77:
	;
	goto L78
L78:
	;
	v265 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v271 = v265 + v252<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L79:
	;
	goto L72
L80:
	;
	v284 = v209
	v287 = v96
	v294 = v278
	goto L62
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v294
	F__bt_relbuf(m, v284)
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L20
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	goto L36
L84:
	;
	v96 = v287
	v97 = v294
	goto L35
L85:
	;
	goto L34
L86:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v320 + int32(4)
	F_errmsg_internal(m, int32(675745), v17+int32(16))
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L20
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(493016), int32(2563), int32(104390))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L20
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v22)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v96
	*(*int32)(unsafe.Add(mBase, uint32(v17)+4)) = v338 + int32(4)
	F_errmsg_internal(m, int32(673509), v17)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L20
	} else {
		goto L90
	}
L90:
	;
	F_errfinish(m, int32(493016), int32(2581), int32(104390))
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L20
	} else {
		goto L91
	}
L91:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L92:
	;
	v383 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382)+16)))
	v384 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = v384
	v386 = v383 + v382
	v387 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386)+12)))
	if v387&int32(20) == int32(0) {
		goto L98
	} else {
		goto L99
	}
L93:
	;
	v368 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v368+(v355^int32(-1))<<(uint(int32(2))%32))))
	v382 = v374
	goto L92
L94:
	;
	goto L95
L95:
	;
	v376 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v382 = v376 + v355<<(uint(int32(13))%32) + int32(-8192)
	goto L92
L96:
	;
	goto L10
L97:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	F__bt_relbuf(m, v427)
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L20
	} else {
		goto L116
	}
L98:
	;
	if v50 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L99:
	;
	goto L100
L100:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v386+base.B2i32(l3 == int32(1))<<(uint(int32(2))%32))))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v419
	v421 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	if v421 == int32(0) {
		goto L97
	} else {
		goto L114
	}
L101:
	;
	v394 = int32(1)
	v397 = *(*int32)(unsafe.Add(mBase, uint32(v386)+4))
	if v397 != 0 {
		goto L104
	} else {
		goto L105
	}
L102:
	;
	goto L103
L103:
	;
	v403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v382)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v403) {
		goto L109
	} else {
		goto L110
	}
L104:
	;
	v398 = int32(2)
	goto L106
L105:
	;
	v398 = v394
	goto L106
L106:
	;
	v399 = F__bt_readpage(m, l0, v394, v398, v60)
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L20
	} else {
		goto L107
	}
L107:
	;
	if v399 != 0 {
		goto L96
	} else {
		goto L108
	}
L108:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v21)+68))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v401
	goto L97
L109:
	;
	v411 = int32(base.Ui32(v403+int32(262120)) >> (uint(int32(2)) % 32))
	goto L111
L110:
	;
	v411 = int32(0)
	goto L111
L111:
	;
	v414 = F__bt_readpage(m, l0, l3, v411&int32(65535), v60)
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L20
	} else {
		goto L112
	}
L112:
	;
	if v414 != 0 {
		goto L96
	} else {
		goto L113
	}
L113:
	;
	v416 = *(*int32)(unsafe.Add(mBase, uint32(v21)+64))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = v416
	goto L97
L114:
	;
	F__bt_parallel_release(m, l0, v419, v384)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L20
	} else {
		goto L115
	}
L115:
	;
	goto L97
L116:
	;
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v17)+28))
	if v431 != 0 {
		v39 = int32(0)
		goto L9
	} else {
		goto L117
	}
L117:
	;
	goto L7
L118:
	;
	v495 = int32(1)
	goto L5
L119:
	;
	F__bt_unlockbuf(m, v433)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L20
	} else {
		goto L122
	}
L120:
	;
	goto L121
L121:
	;
	v439 = F_BufferGetLSNAtomic(m, v433)
	mBase = m.M
	v440 = m.ExcPending
	if v440 != 0 {
		goto L20
	} else {
		goto L123
	}
L122:
	;
	goto L118
L123:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v21)+72)) = v439
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v21)+56))
	F__bt_relbuf(m, v442)
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L20
	} else {
		goto L124
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+56)) = int32(0)
	goto L118
L125:
	;
	v495 = int32(0)
	goto L5
}
func F__bt_recsplitloc(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	v2 = l1
	v5 = int32(0)
	v9 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
	if v9 == v2 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v14 = v11 + (l2 - v12)
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v17&int32(1) != 0 {
			v62 = v16
			v64 = v14
			v65 = v15
			v67 = v62
			v68 = int32(-8)
			v70 = v64
			v71 = v65
			v73 = v67
			v74 = v68
			v75 = int32(1)
			v76 = v70
			v77 = v71
		} else {
			v73 = v16
			v74 = v5
			v75 = v5
			v76 = v14
			v77 = v15
		}
	} else {
		v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)))
		if v20 != int32(1) {
			v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
			v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v59 = v56 + (l2 - v57)
			v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v20 != 0 {
				v62 = l3
				v64 = v59
				v65 = v60
				v67 = v62
				v68 = int32(-8)
				v70 = v64
				v71 = v65
				v73 = v67
				v74 = v68
				v75 = int32(1)
				v76 = v70
				v77 = v71
			} else {
				v73 = l3
				v74 = int32(0)
				v75 = v5
				v76 = v59
				v77 = v60
			}
		} else {
			if base.Ui32(l3) < base.Ui32(int32(65)) {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v59 = v56 + (l2 - v57)
				v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				if v20 != 0 {
					v62 = l3
					v64 = v59
					v65 = v60
					v67 = v62
					v68 = int32(-8)
					v70 = v64
					v71 = v65
					v73 = v67
					v74 = v68
					v75 = int32(1)
					v76 = v70
					v77 = v71
				} else {
					v73 = l3
					v74 = int32(0)
					v75 = v5
					v76 = v59
					v77 = v60
				}
			} else {
				v25 = int32(-8)
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v26+v2<<(uint(int32(2))%32))+20))
				v33 = v26 + v30&int32(32767)
				v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+6)))
				if v34&int32(8192) == int32(0) {
					v50 = v25
				} else {
					v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+5)))
					if v39&int32(32) == int32(0) {
						v50 = v25
					} else {
						v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+2)))
						v50 = v34&int32(8191) - v46 - int32(8)
					}
				}
				v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
				v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
				v67 = l3
				v68 = v50
				v70 = v51 + (l2 - v52)
				v71 = v55
				v73 = v67
				v74 = v68
				v75 = int32(1)
				v76 = v70
				v77 = v71
			}
		}
	}
	v80 = v74 + v77 - (l2 + v73)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v75 != 0 {
		v86 = int32(0)
	} else {
		v86 = v73 + int32(65524)
	}
	v87 = v76 - v81 + v86
	if (v80|v87)&int32(32768) == int32(0) {
		v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
		if base.Ui32(v93) < base.Ui32(v73) {
			v95 = v93
		} else {
			v95 = v73
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v95
		v97 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v99 = int32(10)
		v102 = int32(0)
		*(*uint16)(unsafe.Add(mBase, uint32(v97+v98*v99))) = uint16(v102)
		v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint16)(unsafe.Add(mBase, uint32(v104+v105*v99)+2)) = uint16(v80)
		v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint16)(unsafe.Add(mBase, uint32(v110+v111*v99)+4)) = uint16(v87)
		v116 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint16)(unsafe.Add(mBase, uint32(v116+v117*v99)+6)) = uint16(v2)
		v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+44))
		v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*uint8)(unsafe.Add(mBase, uint32(v122+v123*v99)+8)) = uint8(v102)
		v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		*(*int32)(unsafe.Add(mBase, uint32(l0)+40)) = v129 + int32(1)
	} else {
	}
	return
}
func F__bt_restore_meta(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int64
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v91 int64
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = F_XLogInitBufferForRedo(m, l0, l1)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		v15 = v9 + int32(12)
		v16 = int32(0)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)+72))
		if v18 < l1 {
			v40 = v16
			v43 = v40
		} else {
			v24 = v17 + l1*int32(52) + int32(76)
			v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			if v25 != int32(1) {
				v40 = v16
				v43 = v40
			} else {
				v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+43)))
				if v28 == int32(0) {
					if v15 == int32(0) {
						v40 = v16
						v43 = v40
					} else {
						v33 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v33
						v43 = v33
					}
				} else {
					if v15 != 0 {
						v36 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+48)))
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v36
					} else {
					}
					v38 = *(*int32)(unsafe.Add(mBase, uint32(v24)+44))
					v40 = v38
					v43 = v40
				}
			}
		}
		if v12 < int32(0) {
			v47 = *(*int32)(unsafe.Add(mBase, _consts[5]))
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v47+(v12^int32(-1))<<(uint(int32(2))%32))))
			v61 = v53
		} else {
			v55 = *(*int32)(unsafe.Add(mBase, _consts[6]))
			v61 = v55 + v12<<(uint(int32(13))%32) + int32(-8192)
		}
		F_PageInit(m, v61, int32(8192), int32(16))
		mBase = m.M
		*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = int32(340322)
		v67 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+28)) = v67
		v69 = *(*int32)(unsafe.Add(mBase, uint32(v43)+4))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+32)) = v69
		v71 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v71
		v73 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+40)) = v73
		v75 = *(*int32)(unsafe.Add(mBase, uint32(v43)+16))
		*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = v75
		v77 = *(*int32)(unsafe.Add(mBase, uint32(v43)+20))
		*(*int64)(unsafe.Add(mBase, uint32(v61)+56)) = int64(-4616189618054758400)
		*(*int32)(unsafe.Add(mBase, uint32(v61)+48)) = v77
		v83 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+24)))
		*(*uint8)(unsafe.Add(mBase, uint32(v61-int32(-64)))) = uint8(v83)
		v85 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+16)))
		v87 = int32(8)
		*(*uint16)(unsafe.Add(mBase, uint32(v61+v85)+12)) = uint16(v87)
		*(*uint32)(unsafe.Add(mBase, uint32(v61)+4)) = uint32(v11)
		v91 = int64(base.Ui64(v11) >> (uint(int64(32)) % 64))
		*(*uint32)(unsafe.Add(mBase, uint32(v61))) = uint32(v91)
		v93 = int32(72)
		*(*uint16)(unsafe.Add(mBase, uint32(v61)+12)) = uint16(v93)
		F_MarkBufferDirty(m, v12)
		mBase = m.M
		v96 = m.ExcPending
		if v96 != 0 {
			return
		} else {
			F_UnlockReleaseBuffer(m, v12)
			mBase = m.M
			v98 = m.ExcPending
			if v98 != 0 {
				return
			} else {
				m.G0 = v9 + int32(16)
				return
			}
		}
	}
}
func F__bt_restore_page(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	v4 = int32(0)
	v7 = m.G0
	v9 = v7 - int32(2448)
	m.G0 = v9
	v11 = l1 + l2
	if base.Ui32(l1) < base.Ui32(v11) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = l1
	v16 = v4
	goto L4
L2:
	;
	v43 = v4
	goto L3
L3:
	;
	v47 = v43
	goto L8
L4:
	;
	v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v9+int32(816)+v16<<(uint(int32(2))%32)))) = v14
	v26 = int32(1)
	v34 = (v19&int32(8191) + int32(7)) & int32(16376)
	*(*uint16)(unsafe.Add(mBase, uint32(v9+v16<<(uint(v26)%32)))) = uint16(v34)
	v37 = v16 + v26
	v38 = v14 + v34
	if base.Ui32(v38) < base.Ui32(v11) {
		v14 = v38
		v16 = v37
		goto L4
	} else {
		goto L6
	}
L5:
	;
	v43 = v37
	goto L3
L6:
	;
	goto L5
L7:
	;
	m.G0 = v9 + int32(2448)
	return
L8:
	;
	v53 = v47 - int32(1)
	if v53 < int32(0) {
		goto L7
	} else {
		goto L10
	}
L9:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v75 = m.ExcPending
	if v75 != 0 {
		goto L11
	} else {
		goto L14
	}
L10:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v9+int32(816)+v53<<(uint(int32(2))%32))))
	v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9+v53<<(uint(int32(1))%32)))))
	v70 = F_PageAddItemExtended(m, l0, v61, v65, (v43-v53)&int32(65535), int32(0))
	mBase = m.M
	v71 = m.ExcPending
	if v71 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	return
L12:
	;
	if v70 != 0 {
		v47 = v53
		goto L8
	} else {
		goto L13
	}
L13:
	;
	goto L9
L14:
	;
	F_errmsg_internal(m, int32(403957), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L11
	} else {
		goto L15
	}
L15:
	;
	F_errfinish(m, int32(493052), int32(77), int32(403253))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L11
	} else {
		goto L16
	}
L16:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_scanbehind_checkkeys(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+52))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
	if v15&int32(32) == int32(0) {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+192))
		v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+8)))
		v29 = v27
	} else {
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
		if v20&int32(8192) != 0 {
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v13)+192))
			v27 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+8)))
			v29 = v27
		} else {
			v29 = v20 & int32(4095)
		}
	}
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v31 = int32(0)
	v36 = F__bt_tuple_before_array_skeys(m, l0, l1, l2, v14, v29, v31, v31, v11+int32(7))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		return int32(0)
	} else {
		if v36 != 0 {
			v88 = v31
			m.G0 = v11 + int32(16)
			return v88
		} else {
			v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+7)))
			if v40 != 0 {
				v88 = v31
				m.G0 = v11 + int32(16)
				return v88
			} else {
				v41 = int32(1)
				v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v30)+19)))
				if v42 != v41 {
					v88 = v41
					m.G0 = v11 + int32(16)
					return v88
				} else {
					v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)+52))
					v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+7)))
					if v47&int32(32) == int32(0) {
						v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+192))
						v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58)+8)))
						v61 = v59
					} else {
						v52 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+4)))
						if v52&int32(8192) != 0 {
							v58 = *(*int32)(unsafe.Add(mBase, uint32(v45)+192))
							v59 = int32(*(*int16)(unsafe.Add(mBase, uint32(v58)+8)))
							v61 = v59
						} else {
							v61 = v52 & int32(4095)
						}
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
					v63 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v63
					v74 = F__bt_check_compare(m, l0, v63-l1, l2, v61, v46, v63, v63, v11+int32(15), v11+int32(8))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+15)))
						if v76 == int32(0) {
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+8))
							v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79+v80*int32(48))+6)))
							if v84 != int32(3) {
								v88 = v63
							} else {
								v88 = int32(1)
							}
						} else {
							v88 = int32(1)
						}
						m.G0 = v11 + int32(16)
						return v88
					}
				}
			}
		}
	}
}
func F__bt_search(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v152 int32
	_ = v152
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	v17 = F__bt_getroot(m, l0, l1, l4)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v17
	if v17 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v27 = base.B2i32(l4 == int32(2))
	v34 = v17
	v38 = int32(1)
	v39 = int32(0)
	goto L6
L6:
	;
	v45 = F__bt_moveright(m, l0, l1, l2, v34, v27, v39, v38)
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L9
	}
L7:
	;
	if l4 != int32(2) {
		goto L49
	} else {
		goto L50
	}
L8:
	;
	goto L7
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v45
	if v45 < int32(0) {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	v82 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v82) {
		goto L17
	} else {
		goto L18
	}
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v51+(v45^int32(-1))<<(uint(int32(2))%32))))
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v57)+16)))
	v59 = v57 + v58
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v59)+12)))
	if v60&int32(1) == int32(0) {
		v79 = v58
		v80 = v57
		v81 = v59
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v69 = v66 + v45<<(uint(int32(13))%32)
	v71 = v69 + int32(-8192)
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69-int32(8176)))))
	v75 = v71 + v74
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v75)+12)))
	if v76&int32(1) != 0 {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	goto L8
L15:
	;
	v79 = v74
	v80 = v71
	v81 = v75
	goto L10
L16:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v152&int32(65535)<<(uint(int32(2))%32)+v80)+20))
	v171 = v80 + v168&int32(32767)
	v172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171)+2)))
	v173 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v171))))
	v175 = F_palloc(m, int32(12))
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L37
	}
L17:
	;
	v90 = int32(base.Ui32(v82+int32(262120)) >> (uint(int32(2)) % 32))
	goto L19
L18:
	;
	v90 = int32(0)
	goto L19
L19:
	;
	v95 = v79 + v80
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	if v96 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v97 = int32(2)
	goto L22
L21:
	;
	v97 = int32(1)
	goto L22
L22:
	;
	if base.Ui32(v90&int32(65535)) < base.Ui32(v97) {
		v152 = v97
		goto L16
	} else {
		goto L23
	}
L23:
	;
	v99 = int32(1)
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+3)))
	v109 = v97
	v112 = v90 + v99
	goto L24
L24:
	;
	v125 = int32(base.Ui32((v112-v109)&int32(65534))>>(uint(int32(1))%32)) + v109
	v128 = F__bt_compare(m, l0, l2, v80, v125&int32(65535))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L26
	}
L25:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v95)+12)))
	if v140&int32(1) != 0 {
		goto L34
	} else {
		goto L35
	}
L26:
	;
	v130 = base.B2i32(v128 < v101^v99)
	if v128 < v101^v99 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v131 = v125
	goto L29
L28:
	;
	v131 = v112
	goto L29
L29:
	;
	if v128 < v101^v99 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v136 = v109
	goto L32
L31:
	;
	v136 = v125 + int32(1)
	goto L32
L32:
	;
	if base.Ui32(v136&int32(65535)) < base.Ui32(v131&int32(65535)) {
		v109 = v136
		v112 = v131
		goto L24
	} else {
		goto L33
	}
L33:
	;
	goto L25
L34:
	;
	v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+4)))
	v152 = v136 - v143
	goto L16
L35:
	;
	goto L36
L36:
	;
	v152 = v136 - int32(1)
	goto L16
L37:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v177 < int32(0) {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = v39
	*(*uint16)(unsafe.Add(mBase, uint32(v175)+4)) = uint16(v152)
	*(*int32)(unsafe.Add(mBase, uint32(v175))) = v196
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v205 == int32(1) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	v181 = *(*int32)(unsafe.Add(mBase, _consts[8]))
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v181+(v177^int32(-1))<<(uint(int32(6))%32))+16))
	v196 = v187
	goto L38
L40:
	;
	goto L41
L41:
	;
	v189 = *(*int32)(unsafe.Add(mBase, _consts[9]))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v189+v177<<(uint(int32(6))%32)+int32(-64))+16))
	v196 = v195
	goto L38
L42:
	;
	v208 = int32(2)
	goto L44
L43:
	;
	v208 = v38
	goto L44
L44:
	;
	if l4 == int32(2) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v209 = v208
	goto L47
L46:
	;
	v209 = v38
	goto L47
L47:
	;
	v210 = F__bt_relandgetbuf(m, l0, v200, v172|v173<<(uint(int32(16))%32), v209)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v210
	v34 = v210
	v38 = v209
	v39 = v175
	goto L6
L49:
	;
	return v39
L50:
	;
	if v38 != int32(1) {
		goto L49
	} else {
		goto L51
	}
L51:
	;
	F__bt_unlockbuf(m, v45)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	F__bt_lockbuf(m, v222, int32(2))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v229 = F__bt_moveright(m, l0, l1, l2, v226, int32(1), v39, int32(2))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L54
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v229
	goto L49
}
func F__bt_setup_array_cmp(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v39 int32
	_ = v39
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
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v93 int32
	_ = v93
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	v11 = m.G0
	v13 = v11 + int32(-64)
	m.G0 = v13
	v15 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
	v19 = (v15 - int32(1)) << (uint(int32(2)) % 32)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)+212))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v19+v21)))
	if v23 == l2 {
		v26 = F_index_getprocinfo(m, v20, v15, int32(1))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = *(*int64)(unsafe.Add(mBase, uint32(v26)))
			*(*int64)(unsafe.Add(mBase, uint32(l3))) = v28
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+24))
			*(*int32)(unsafe.Add(mBase, uint32(l3)+24)) = v30
			v32 = *(*int64)(unsafe.Add(mBase, uint32(v26)+16))
			*(*int64)(unsafe.Add(mBase, uint32(l3)+16)) = v32
			v34 = *(*int64)(unsafe.Add(mBase, uint32(v26)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l3)+8)) = v34
			if l4 == int32(0) {
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(l4))) = l3
			}
			m.G0 = v13 - int32(-64)
			return
		}
	} else {
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v40 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
		v42 = *(*int32)(unsafe.Add(mBase, uint32(v40+v19)))
		v44 = F_get_opfamily_proc(m, v42, v23, l2, int32(1))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return
		} else {
			if v44 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v80 = m.ExcPending
				if v80 != 0 {
					return
				} else {
					v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v82 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v81
					*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = l2
					*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v23
					*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(1)
					F_errmsg_internal(m, int32(675835), v13)
					mBase = m.M
					v93 = m.ExcPending
					if v93 != 0 {
						return
					} else {
						F_errfinish(m, int32(488581), int32(2691), int32(233748))
						mBase = m.M
						v98 = m.ExcPending
						if v98 != 0 {
							return
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
				F_fmgr_info_cxt(m, v44, l3, v48)
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					if l4 == int32(0) {
						m.G0 = v13 - int32(-64)
						return
					} else {
						v53 = *(*int32)(unsafe.Add(mBase, uint32(v20)+208))
						v54 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v53+v54<<(uint(int32(2))%32)-int32(4))))
						v62 = F_get_opfamily_proc(m, v60, l2, l2, int32(1))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return
						} else {
							if v62 == int32(0) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v102 = m.ExcPending
								if v102 != 0 {
									return
								} else {
									v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+4)))
									v104 = *(*int32)(unsafe.Add(mBase, uint32(v20)+48))
									*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v104 + int32(4)
									*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v103
									*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = l2
									*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = int32(1)
									F_errmsg_internal(m, int32(675835), v11+int32(-32))
									mBase = m.M
									v117 = m.ExcPending
									if v117 != 0 {
										return
									} else {
										F_errfinish(m, int32(488581), int32(2713), int32(233748))
										mBase = m.M
										v122 = m.ExcPending
										if v122 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							} else {
								v66 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v39)+28))
								F_fmgr_info_cxt(m, v62, v66, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									m.G0 = v13 - int32(-64)
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
func F__bt_upgrademetapage(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	v2 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(l0)+12)) = uint16(v2)
	v6 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l0-int32(-64)))) = uint8(v6)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+48)) = v6
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = int32(3)
	return
}
