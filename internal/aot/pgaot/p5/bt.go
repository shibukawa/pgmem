package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__bt_bottomupdel_finish_pending(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
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
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v229 int32
	_ = v229
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v250 int32
	_ = v250
	v4 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v18 <= v4 {
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		v33 = v25
		v37 = v4
		for {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			v46 = v43 + v33*int32(6)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
			v50 = v47 + v33<<(uint(int32(3))%32)
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
			v52 = v51 + v37
			v59 = v52&int32(65535)<<(uint(int32(2))%32) + (l0 + int32(24)) - int32(4)
			v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
			v63 = l0 + v60&int32(32767)
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+7)))
			if v64&int32(32) != 0 {
				v67 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)))
				if v67&int32(8192) != 0 {
					v92 = v67 & int32(4095)
					v93 = int32(0)
					if int32(2) <= v18 {
						v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+2)))
						v98 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63))))
						v99 = int32(16)
						v102 = v63 + (v97 | v98<<(uint(v99)%32))
						v103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102))))
						v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102)+2)))
						v107 = v103<<(uint(v99)%32) | v106
						v110 = int32(6)
						v112 = v102 + int32(base.Ui32(v92)>>(uint(int32(1))%32))*v110
						v113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112))))
						v116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v112)+2)))
						v117 = v113<<(uint(v99)%32) | v116
						v121 = v102 + v92*v110
						v124 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121-v110))))
						v129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v121-int32(4)))))
						v134 = base.B2i32(v117 == v107)
						v136 = base.B2i32(v107 != v117) & base.B2i32(v117 == v124<<(uint(v99)%32)|v129)
					} else {
						v134 = v93
						v136 = v93
					}
					if v92 == int32(0) {
						v229 = v33
					} else {
						v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+2)))
						v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63))))
						v145 = v63 + (v140 | v141<<(uint(int32(16))%32))
						v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = v146
						v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v145)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v148)
						v150 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+6)) = uint16(v150)
						v152 = int32(6)
						*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v152)
						v154 = int32(1)
						v155 = v92 - v154
						v156 = int32(0)
						v159 = v134 | v136&base.B2i32(v155 == v156)
						*(*uint8)(unsafe.Add(mBase, uint32(v46)+3)) = uint8(v159)
						*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)) = uint8(v156)
						*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v52)
						v165 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
						v167 = v165 + v154
						*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v167
						if v92 == v154 {
							v229 = v167
						} else {
							v174 = v154
							v175 = v46
							v177 = v50
							for {
								v189 = v177 + int32(8)
								v190 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+2)))
								v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63))))
								v196 = int32(6)
								v198 = v63 + (v190 | v191<<(uint(int32(16))%32)) + v174*v196
								v199 = *(*int32)(unsafe.Add(mBase, uint32(v198)))
								*(*int32)(unsafe.Add(mBase, uint32(v189))) = v199
								v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v198)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v177)+12)) = uint16(v201)
								v203 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								*(*uint16)(unsafe.Add(mBase, uint32(v177)+14)) = uint16(v203)
								*(*uint16)(unsafe.Add(mBase, uint32(v175)+10)) = uint16(v196)
								v208 = v136 & base.B2i32(v174 == v155)
								*(*uint8)(unsafe.Add(mBase, uint32(v175)+9)) = uint8(v208)
								v210 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v175)+8)) = uint8(v210)
								v213 = v175 + v196
								*(*uint16)(unsafe.Add(mBase, uint32(v213))) = uint16(v52)
								v215 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								v216 = int32(1)
								v217 = v215 + v216
								*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v217
								v220 = v174 + v216
								if v220 != v92 {
									v174 = v220
									v175 = v213
									v177 = v189
									continue
								} else {
									break
								}
								break
							}
							v229 = v217
						}
					}
				} else {
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
					*(*int32)(unsafe.Add(mBase, uint32(v50))) = v71
					v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)))
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v73)
					v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+6)) = uint16(v75)
					*(*uint8)(unsafe.Add(mBase, uint32(v46)+3)) = uint8(base.B2i32(int32(1) < v18))
					v78 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)) = uint8(v78)
					*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v52)
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					v85 = int32(base.Ui32(v81)>>(uint(int32(17))%32)) + int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v85)
					v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v89 = v87 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v89
					v229 = v89
				}
			} else {
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
				*(*int32)(unsafe.Add(mBase, uint32(v50))) = v71
				v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v73)
				v75 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+6)) = uint16(v75)
				*(*uint8)(unsafe.Add(mBase, uint32(v46)+3)) = uint8(base.B2i32(int32(1) < v18))
				v78 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)) = uint8(v78)
				*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v52)
				v81 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
				v85 = int32(base.Ui32(v81)>>(uint(int32(17))%32)) + int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v85)
				v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v89 = v87 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v89
				v229 = v89
			}
			v240 = v37 + int32(1)
			v241 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			if v240 < v241 {
				v33 = v229
				v37 = v240
				continue
			} else {
				break
			}
			break
		}
		if v18 <= int32(1) {
		} else {
			v245 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v245<<(uint(int32(2))%32))+46)) = uint16(v241)
			v250 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v250 + int32(1)
		}
	}
	*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(l1)+28)) = int64(0)
	return
}
func F__bt_compare(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
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
	var v22 int32
	_ = v22
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v315 int32
	_ = v315
	v14 = m.G0
	v16 = v14 - int32(16)
	m.G0 = v16
	v18 = int32(1)
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v21 = l2 + v20
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+12)))
	if v22&v18 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v16 + int32(16)
	return v315
L2:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v29 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l3<<(uint(int32(2))%32)+l2)+20))
	v38 = l2 + v35&int32(32767)
	v39 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+7)))
	if v39&int32(32) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v30 = int32(2)
	goto L7
L6:
	;
	v30 = int32(1)
	goto L7
L7:
	;
	if v30 == l3 {
		v315 = v18
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	v315 = int32(1)
	goto L1
L10:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v53 < v54 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v50)+8)))
	v53 = v51
	goto L10
L12:
	;
	v44 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	if v44&int32(8192) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v53 = v44 & int32(4095)
	goto L10
L14:
	;
	v56 = v53
	goto L16
L15:
	;
	v56 = v54
	goto L16
L16:
	;
	if int32(0) < v56 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v60 = v38 + int32(8)
	v66 = int32(1)
	v68 = l1 + int32(16)
	goto L20
L18:
	;
	v181 = v54
	goto L19
L19:
	;
	v192 = int32(1)
	if v53 < v181 {
		v315 = v192
		goto L1
	} else {
		goto L57
	}
L20:
	;
	v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v68)+4)))
	v81 = v79 - int32(1)
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+6)))
	if int32(0) <= v82 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v178 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v181 = v178
	goto L19
L22:
	;
	if v66 != v56 {
		v66 = v66 + int32(1)
		v68 = v68 + int32(48)
		goto L20
	} else {
		goto L56
	}
L23:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v68)+44))
	v158 = F_FunctionCall2Coll(m, v68+int32(16), v156, v135, v157)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L36
	} else {
		goto L50
	}
L24:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v146&int32(1) != 0 {
		goto L22
	} else {
		goto L46
	}
L25:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v68)))
	if v136&int32(1) == int32(0) {
		goto L23
	} else {
		goto L42
	}
L26:
	;
	v130 = F_nocache_index_getattr(m, v38, v79, v19)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L36
	} else {
		goto L41
	}
L27:
	;
	v87 = v19 + int32(20) + v81<<(uint(int32(4))%32)
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)))
	if v88 < int32(0) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v81>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v120)>>(uint(v81&int32(7))%32))&int32(1) == int32(0) {
		goto L24
	} else {
		goto L40
	}
L30:
	;
	v91 = v88 + v60
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+6)))
	if v92 != int32(1) {
		v135 = v91
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v87)+4)))
	switch v95 - int32(1) {
	case 0:
		goto L35
	case 1:
		goto L34
	default:
		goto L32
	case 3:
		goto L33
	}
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v91)))
	v135 = v100
	goto L25
L34:
	;
	v99 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91))))
	v135 = v99
	goto L25
L35:
	;
	v98 = int32(*(*int8)(unsafe.Add(mBase, uint32(v91))))
	v135 = v98
	goto L25
L36:
	;
	return int32(0)
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = base.I32_extend16_s(v95)
	F_errmsg_internal(m, int32(483438), v16)
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(326693), int32(70), int32(67821))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L36
	} else {
		goto L39
	}
L39:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L40:
	;
	goto L26
L41:
	;
	v135 = v130
	goto L25
L42:
	;
	if v136&int32(33554432) != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v145 = int32(-1)
	goto L45
L44:
	;
	v145 = int32(1)
	goto L45
L45:
	;
	v315 = v145
	goto L1
L46:
	;
	if v146&int32(33554432) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v153 = int32(1)
	goto L49
L48:
	;
	v153 = int32(-1)
	goto L49
L49:
	;
	v315 = v153
	goto L1
L50:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v68)+3)))
	if v160&int32(1) == int32(0) {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	if v158 < int32(0) {
		goto L9
	} else {
		goto L54
	}
L52:
	;
	v169 = v158
	goto L53
L53:
	;
	if v169 != 0 {
		v315 = v169
		goto L1
	} else {
		goto L55
	}
L54:
	;
	v169 = int32(0) - v158
	goto L53
L55:
	;
	goto L22
L56:
	;
	goto L21
L57:
	;
	v194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+6)))
	if v194&int32(8192) == int32(0) {
		v220 = v38
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v306 != 0 {
		goto L82
	} else {
		goto L83
	}
L59:
	;
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v302 != 0 {
		v315 = v192
		goto L1
	} else {
		goto L81
	}
L60:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v221 == int32(0) {
		v304 = v220
		goto L58
	} else {
		goto L66
	}
L61:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	if v199&int32(8192) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if v199&int32(4096) == int32(0) {
		goto L59
	} else {
		goto L65
	}
L63:
	;
	goto L64
L64:
	;
	v213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+2)))
	v214 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v220 = v38 + (v213 | v214<<(uint(int32(16))%32))
	goto L60
L65:
	;
	v220 = v38 + v194&int32(8191) - int32(6)
	goto L60
L66:
	;
	if v220 == int32(0) {
		v315 = v192
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v229 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+2)))
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221))))
	v231 = int32(16)
	v233 = v229 | v230<<(uint(v231)%32)
	v234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+2)))
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220))))
	v238 = v234 | v235<<(uint(v231)%32)
	if base.Ui32(v233) < base.Ui32(v238) {
		v249 = int32(-1)
		goto L69
	} else {
		goto L70
	}
L68:
	;
	if v249 <= int32(0) {
		v315 = v249
		goto L1
	} else {
		goto L73
	}
L69:
	;
	goto L68
L70:
	;
	if base.Ui32(v238) < base.Ui32(v233) {
		v249 = int32(1)
		goto L69
	} else {
		goto L71
	}
L71:
	;
	v243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+4)))
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	if base.Ui32(v243) < base.Ui32(v244) {
		v249 = int32(-1)
		goto L69
	} else {
		goto L72
	}
L72:
	;
	v249 = base.B2i32(base.Ui32(v244) < base.Ui32(v243))
	goto L69
L73:
	;
	v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+7)))
	if v252&int32(32) == int32(0) {
		v315 = v249
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	if v257&int32(8192) == int32(0) {
		v315 = v249
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+2)))
	v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v265 = int32(16)
	v271 = int32(6)
	v275 = v38 + (v263 | v264<<(uint(v265)%32)) + v257&int32(4095)*v271 - v271
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+2)))
	v280 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262))))
	v283 = v279 | v280<<(uint(v265)%32)
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+2)))
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275))))
	v288 = v284 | v285<<(uint(v265)%32)
	if base.Ui32(v283) < base.Ui32(v288) {
		v299 = int32(-1)
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v315 = base.B2i32(int32(0) < v299)
	goto L1
L77:
	;
	goto L76
L78:
	;
	if base.Ui32(v288) < base.Ui32(v283) {
		v299 = int32(1)
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v262)+4)))
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v275)+4)))
	if base.Ui32(v293) < base.Ui32(v294) {
		v299 = int32(-1)
		goto L77
	} else {
		goto L80
	}
L80:
	;
	v299 = base.B2i32(base.Ui32(v294) < base.Ui32(v293))
	goto L77
L81:
	;
	v304 = int32(0)
	goto L58
L82:
	;
	v315 = int32(0)
	goto L1
L83:
	;
	if v181 != v53 {
		goto L82
	} else {
		goto L84
	}
L84:
	;
	if v304 != 0 {
		goto L82
	} else {
		goto L85
	}
L85:
	;
	v308 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v308 != 0 {
		v315 = v192
		goto L1
	} else {
		goto L86
	}
L86:
	;
	goto L82
}
func F__bt_compare_array_elements(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_FunctionCall2Coll(m, v6, v7, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		if v10 < int32(0) {
			v17 = int32(1)
		} else {
			v17 = int32(0) - v10
		}
		v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		if v18 != 0 {
			v19 = v17
		} else {
			v19 = v10
		}
		return v19
	}
}
func F__bt_delete_or_dedup_one_page(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v81 int32
	_ = v81
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v261 int32
	_ = v261
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v382 int32
	_ = v382
	var v395 int32
	_ = v395
	var v402 int32
	_ = v402
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v444 int32
	_ = v444
	var v447 int32
	_ = v447
	var v451 int32
	_ = v451
	var v454 int32
	_ = v454
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v472 int32
	_ = v472
	var v499 int32
	_ = v499
	var v505 int32
	_ = v505
	var v507 int32
	_ = v507
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v549 int32
	_ = v549
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v575 int32
	_ = v575
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v609 int32
	_ = v609
	var v618 int32
	_ = v618
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v635 int32
	_ = v635
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v652 int32
	_ = v652
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v662 int32
	_ = v662
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v675 int32
	_ = v675
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v688 int32
	_ = v688
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v738 int32
	_ = v738
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
	var v761 int32
	_ = v761
	var v768 int32
	_ = v768
	var v791 int32
	_ = v791
	var v796 int32
	_ = v796
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v804 int32
	_ = v804
	var v809 int32
	_ = v809
	var v815 int32
	_ = v815
	var v817 int32
	_ = v817
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v826 int32
	_ = v826
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v830 int32
	_ = v830
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v844 int64
	_ = v844
	var v853 int32
	_ = v853
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v873 int32
	_ = v873
	var v876 int32
	_ = v876
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v894 int32
	_ = v894
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v907 int32
	_ = v907
	var v913 int32
	_ = v913
	var v939 int32
	_ = v939
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v987 int32
	_ = v987
	var v993 int32
	_ = v993
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1018 int32
	_ = v1018
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1031 int32
	_ = v1031
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1048 int32
	_ = v1048
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1071 int32
	_ = v1071
	var v1078 int32
	_ = v1078
	var v1086 int32
	_ = v1086
	var v1090 int32
	_ = v1090
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1124 int32
	_ = v1124
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1155 int32
	_ = v1155
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1163 int32
	_ = v1163
	var v1165 int32
	_ = v1165
	var v1166 int32
	_ = v1166
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1177 int32
	_ = v1177
	var v1182 int32
	_ = v1182
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1193 int32
	_ = v1193
	var v1194 int32
	_ = v1194
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1212 int32
	_ = v1212
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1242 int32
	_ = v1242
	var v1243 int32
	_ = v1243
	var v1244 int32
	_ = v1244
	var v1249 int32
	_ = v1249
	var v1251 int32
	_ = v1251
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1256 int32
	_ = v1256
	var v1261 int32
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1269 int32
	_ = v1269
	var v1270 int32
	_ = v1270
	var v1273 int32
	_ = v1273
	var v1282 int32
	_ = v1282
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1298 int32
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1328 int32
	_ = v1328
	var v1331 int32
	_ = v1331
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1338 int32
	_ = v1338
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1353 int32
	_ = v1353
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1359 int32
	_ = v1359
	var v1360 int32
	_ = v1360
	var v1365 int32
	_ = v1365
	var v1371 int32
	_ = v1371
	var v1381 int32
	_ = v1381
	var v1386 int32
	_ = v1386
	var v1390 int32
	_ = v1390
	var v1391 int32
	_ = v1391
	var v1395 int32
	_ = v1395
	var v1396 int32
	_ = v1396
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1417 int32
	_ = v1417
	var v1450 int32
	_ = v1450
	var v1451 int32
	_ = v1451
	var v1465 int32
	_ = v1465
	var v1472 int32
	_ = v1472
	var v1480 int32
	_ = v1480
	var v1484 int32
	_ = v1484
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1497 int32
	_ = v1497
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1514 int32
	_ = v1514
	var v1518 int32
	_ = v1518
	var v1520 int32
	_ = v1520
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1528 int32
	_ = v1528
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1540 int32
	_ = v1540
	var v1544 int32
	_ = v1544
	var v1545 int32
	_ = v1545
	var v1546 int32
	_ = v1546
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1557 int32
	_ = v1557
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1576 int32
	_ = v1576
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1621 int32
	_ = v1621
	var v1622 int32
	_ = v1622
	var v1624 int32
	_ = v1624
	var v1636 int32
	_ = v1636
	var v1637 int32
	_ = v1637
	var v1638 int32
	_ = v1638
	var v1643 int32
	_ = v1643
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1662 int32
	_ = v1662
	var v1663 int32
	_ = v1663
	var v1664 int32
	_ = v1664
	var v1667 int32
	_ = v1667
	var v1676 int32
	_ = v1676
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1692 int32
	_ = v1692
	var v1697 int32
	_ = v1697
	var v1722 int32
	_ = v1722
	var v1723 int32
	_ = v1723
	var v1725 int32
	_ = v1725
	var v1727 int32
	_ = v1727
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1748 int32
	_ = v1748
	var v1750 int32
	_ = v1750
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1786 int32
	_ = v1786
	var v1789 int32
	_ = v1789
	var v1790 int32
	_ = v1790
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1794 int32
	_ = v1794
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1801 int32
	_ = v1801
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1815 int32
	_ = v1815
	var v1816 int32
	_ = v1816
	var v1818 int32
	_ = v1818
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1828 int32
	_ = v1828
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int64
	_ = v1833
	var v1838 int32
	_ = v1838
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1844 int32
	_ = v1844
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1853 int32
	_ = v1853
	var v1857 int32
	_ = v1857
	var v1858 int32
	_ = v1858
	var v1860 int32
	_ = v1860
	var v1862 int32
	_ = v1862
	var v1873 int32
	_ = v1873
	var v1877 int32
	_ = v1877
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1887 int32
	_ = v1887
	var v1891 int32
	_ = v1891
	var v1892 int32
	_ = v1892
	var v1893 int64
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1904 int32
	_ = v1904
	var v1905 int32
	_ = v1905
	var v1916 int32
	_ = v1916
	var v1918 int32
	_ = v1918
	var v1921 int32
	_ = v1921
	var v1925 int32
	_ = v1925
	var v1932 int32
	_ = v1932
	var v1953 int32
	_ = v1953
	var v1959 int32
	_ = v1959
	var v1962 int32
	_ = v1962
	var v1964 int32
	_ = v1964
	var v1967 int32
	_ = v1967
	var v1971 int32
	_ = v1971
	var v1972 int32
	_ = v1972
	var v1974 int32
	_ = v1974
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1995 int32
	_ = v1995
	var v1996 int32
	_ = v1996
	var v2001 int32
	_ = v2001
	var v2007 int32
	_ = v2007
	var v2017 int32
	_ = v2017
	var v2022 int32
	_ = v2022
	var v2025 int32
	_ = v2025
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2029 int32
	_ = v2029
	var v2030 int32
	_ = v2030
	var v2035 int32
	_ = v2035
	var v2042 int32
	_ = v2042
	var v2043 int32
	_ = v2043
	var v2048 int32
	_ = v2048
	var v2050 int32
	_ = v2050
	var v2051 int32
	_ = v2051
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2065 int32
	_ = v2065
	var v2073 int32
	_ = v2073
	var v2074 int32
	_ = v2074
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2090 float64
	_ = v2090
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2097 int32
	_ = v2097
	var v2100 int32
	_ = v2100
	var v2102 int32
	_ = v2102
	var v2105 int32
	_ = v2105
	var v2109 int32
	_ = v2109
	var v2112 int32
	_ = v2112
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2119 int32
	_ = v2119
	var v2123 int32
	_ = v2123
	var v2126 int32
	_ = v2126
	var v2127 int32
	_ = v2127
	var v2128 int32
	_ = v2128
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2146 int32
	_ = v2146
	var v2152 int32
	_ = v2152
	var v2162 int32
	_ = v2162
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2176 int32
	_ = v2176
	var v2177 int32
	_ = v2177
	var v2179 int32
	_ = v2179
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2194 int32
	_ = v2194
	var v2200 int32
	_ = v2200
	var v2233 int32
	_ = v2233
	var v2234 int32
	_ = v2234
	var v2238 int32
	_ = v2238
	var v2239 int32
	_ = v2239
	var v2241 int32
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2252 int32
	_ = v2252
	var v2254 int32
	_ = v2254
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2263 int32
	_ = v2263
	var v2267 int32
	_ = v2267
	var v2270 int32
	_ = v2270
	var v2271 int32
	_ = v2271
	var v2272 int32
	_ = v2272
	var v2275 int32
	_ = v2275
	var v2279 int32
	_ = v2279
	var v2284 int32
	_ = v2284
	var v2288 int32
	_ = v2288
	var v2292 int32
	_ = v2292
	var v2295 int64
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2300 int32
	_ = v2300
	var v2302 int32
	_ = v2302
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2311 int32
	_ = v2311
	var v2318 int32
	_ = v2318
	var v2322 int32
	_ = v2322
	var v2327 int32
	_ = v2327
	v8 = int32(0)
	v29 = m.G0
	v31 = v29 - int32(848)
	m.G0 = v31
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v36 < v8 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v31 + int32(848)
	return
L2:
	;
	v791 = int32(1)
	if v768|(l4^v791) != v791 {
		goto L1
	} else {
		goto L99
	}
L3:
	;
	v55 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+16)))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v55+v54)+4))
	if v57 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v40 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v40+(v36^int32(-1))<<(uint(int32(2))%32))))
	v54 = v46
	goto L3
L5:
	;
	goto L6
L6:
	;
	v48 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v54 = v48 + v36<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L7:
	;
	v58 = int32(2)
	goto L9
L8:
	;
	v58 = int32(1)
	goto L9
L9:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v59) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v67 = int32(base.Ui32(v59+int32(262120)) >> (uint(int32(2)) % 32))
	goto L12
L11:
	;
	v67 = int32(0)
	goto L12
L12:
	;
	v69 = v67 & int32(65535)
	if base.Ui32(v69) < base.Ui32(v58) {
		v768 = l5
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v81 = v58
	v97 = v8
	goto L14
L14:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v81&int32(65535)<<(uint(int32(2))%32)+(v54+int32(24))-int32(4))))
	v109 = int32(98304)
	if v108&v109 == v109 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v119 <= int32(0) {
		v768 = l5
		goto L2
	} else {
		goto L20
	}
L16:
	;
	v113 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v31+v97<<(uint(v113)%32)))) = uint16(v81)
	v119 = v97 + v113
	goto L18
L17:
	;
	v119 = v97
	goto L18
L18:
	;
	v121 = v81 + int32(1)
	if base.Ui32(v121&int32(65535)) <= base.Ui32(v69) {
		v81 = v121
		v97 = v119
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v36 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v147 = v119 + int32(1)
	v150 = F_palloc(m, v147<<(uint(int32(2))%32))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v131 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v131+(v36^int32(-1))<<(uint(int32(2))%32))))
	v145 = v137
	goto L21
L23:
	;
	goto L24
L24:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v145 = v139 + v36<<(uint(int32(13))%32) + int32(-8192)
	goto L21
L25:
	;
	return
L26:
	;
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127)+2)))
	v153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v127))))
	*(*int32)(unsafe.Add(mBase, uint32(v150))) = v152 | v153<<(uint(int32(16))%32)
	v159 = v145 + int32(24)
	v166 = int32(1)
	v179 = v150
	v186 = v147
	v187 = v8
	goto L27
L27:
	;
	v192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v31+v187<<(uint(int32(1))%32)))))
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v192<<(uint(int32(2))%32)+v159-int32(4))))
	v201 = v145 + v198&int32(32767)
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v201)+7)))
	if v202&int32(32) != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	F_pg_qsort(m, v395, v382, int32(4), int32(208))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L25
	} else {
		goto L55
	}
L29:
	;
	v406 = v187 + int32(1)
	if v406 != v119 {
		v166 = v382
		v179 = v395
		v186 = v402
		v187 = v406
		goto L27
	} else {
		goto L54
	}
L30:
	;
	v230 = v205 & int32(4095)
	v231 = v166 + v230
	if v186 < v231 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v205 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+4)))
	if v205&int32(8192) != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v210 = v166 + int32(1)
	if v186 < v210 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v214 = F_repalloc(m, v179, v186<<(uint(int32(3))%32))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	v218 = v179
	v219 = v186
	goto L37
L37:
	;
	v223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+2)))
	v224 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201))))
	*(*int32)(unsafe.Add(mBase, uint32(v218+v166<<(uint(int32(2))%32)))) = v223 | v224<<(uint(int32(16))%32)
	v382 = v210
	v395 = v218
	v402 = v219
	goto L29
L38:
	;
	v218 = v214
	v219 = v186 << (uint(int32(1)) % 32)
	goto L37
L39:
	;
	v234 = v186 << (uint(int32(1)) % 32)
	if v231 < v234 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v241 = v179
	v243 = v186
	goto L41
L41:
	;
	if v230 == int32(0) {
		v382 = v166
		v395 = v241
		v402 = v243
		goto L29
	} else {
		goto L46
	}
L42:
	;
	v236 = v234
	goto L44
L43:
	;
	v236 = v231
	goto L44
L44:
	;
	v239 = F_repalloc(m, v179, v236<<(uint(int32(2))%32))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	v241 = v239
	v243 = v236
	goto L41
L46:
	;
	v246 = int32(1)
	v248 = int32(0)
	if v230 != v246 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v261 = v166
	v268 = v248
	v269 = int32(0)
	goto L50
L48:
	;
	v332 = v166
	v339 = v248
	goto L49
L49:
	;
	if v230&v246 == int32(0) {
		v382 = v332
		v395 = v241
		v402 = v243
		goto L29
	} else {
		goto L53
	}
L50:
	;
	v284 = int32(2)
	v285 = v261 << (uint(v284) % 32)
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+2)))
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201))))
	v289 = int32(16)
	v293 = int32(6)
	v295 = v201 + (v287 | v288<<(uint(v289)%32)) + v268*v293
	v296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295))))
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v295)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v241+v285))) = v296<<(uint(v289)%32) | v299
	v303 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+2)))
	v304 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201))))
	v313 = v201 + (v303 | v304<<(uint(v289)%32)) + (v268|int32(1))*v293
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313))))
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v285+(v241+int32(4))))) = v314<<(uint(v289)%32) | v317
	v321 = v268 + v284
	v323 = v261 + v284
	v325 = v269 + v284
	if v325 != v230&int32(4094) {
		v261 = v323
		v268 = v321
		v269 = v325
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v332 = v323
	v339 = v321
	goto L49
L52:
	;
	goto L51
L53:
	;
	v360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201)+2)))
	v361 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v201))))
	v362 = int32(16)
	v368 = v201 + (v360 | v361<<(uint(v362)%32)) + v339*int32(6)
	v369 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v368))))
	v372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v368)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v241+v332<<(uint(int32(2))%32)))) = v369<<(uint(v362)%32) | v372
	v382 = v332 + int32(1)
	v395 = v241
	v402 = v243
	goto L29
L54:
	;
	goto L28
L55:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v382) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v424 = int32(1)
	v428 = int32(0)
	goto L59
L57:
	;
	v472 = v382
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+820)) = l0
	if v36 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v444 = int32(2)
	v447 = *(*int32)(unsafe.Add(mBase, uint32(v395+v424<<(uint(v444)%32))))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v395+v428<<(uint(v444)%32))))
	if v447 == v451 {
		v460 = v428
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v472 = v460 + int32(1)
	goto L58
L61:
	;
	v463 = v424 + int32(1)
	if v463 != v382 {
		v424 = v463
		v428 = v460
		goto L59
	} else {
		goto L64
	}
L62:
	;
	v454 = v428 + int32(1)
	if v424 == v454 {
		v460 = v424
		goto L61
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395+v454<<(uint(int32(2))%32)))) = v447
	v460 = v454
	goto L61
L64:
	;
	goto L60
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v31)+832)) = int64(0)
	v517 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v31)+828)) = uint8(v517)
	*(*int32)(unsafe.Add(mBase, uint32(v31)+824)) = v514
	v521 = F_palloc(m, int32(10864))
	mBase = m.M
	v522 = m.ExcPending
	if v522 != 0 {
		goto L25
	} else {
		goto L69
	}
L66:
	;
	v499 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v499+(v36^int32(-1))<<(uint(int32(6))%32))+16))
	v514 = v505
	goto L65
L67:
	;
	goto L68
L68:
	;
	v507 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v513 = *(*int32)(unsafe.Add(mBase, uint32(v507+v36<<(uint(int32(6))%32)+int32(-64))+16))
	v514 = v513
	goto L65
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+840)) = v521
	v525 = F_palloc(m, int32(8148))
	mBase = m.M
	v526 = m.ExcPending
	if v526 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v31)+844)) = v525
	v549 = v58
	goto L71
L71:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(v31)+844))
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v31)+836))
	v560 = v556 + v557*int32(6)
	v561 = *(*int32)(unsafe.Add(mBase, uint32(v31)+840))
	v564 = v561 + v557<<(uint(int32(3))%32)
	v571 = v549&int32(65535)<<(uint(int32(2))%32) + v159 - int32(4)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	v575 = v145 + v572&int32(32767)
	v576 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v575)+7)))
	if v576&int32(32) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	F_pfree(m, v395)
	mBase = m.M
	v738 = m.ExcPending
	if v738 != 0 {
		goto L25
	} else {
		goto L90
	}
L73:
	;
	v731 = v549 + int32(1)
	v732 = int32(65535)
	if base.Ui32(v731&v732) <= base.Ui32(v67&v732) {
		v549 = v731
		goto L71
	} else {
		goto L89
	}
L74:
	;
	v618 = v579 & int32(4095)
	if v618 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L75:
	;
	v579 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+4)))
	if v579&int32(8192) != 0 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v583 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+2)))
	v584 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575))))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+816)) = v583 | v584<<(uint(int32(16))%32)
	v593 = F_bsearch(m, v31+int32(816), v395, v472, int32(4), int32(208))
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L25
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if v593 == int32(0) {
		goto L73
	} else {
		goto L80
	}
L80:
	;
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v575)))
	*(*int32)(unsafe.Add(mBase, uint32(v564))) = v597
	v599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v564)+4)) = uint16(v599)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v31)+836))
	*(*uint16)(unsafe.Add(mBase, uint32(v564)+6)) = uint16(v601)
	*(*uint16)(unsafe.Add(mBase, uint32(v560))) = uint16(v549)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	v605 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v560)+4)) = uint16(v605)
	*(*uint8)(unsafe.Add(mBase, uint32(v560)+3)) = uint8(v605)
	v609 = int32(98304)
	*(*uint8)(unsafe.Add(mBase, uint32(v560)+2)) = uint8(base.B2i32(v604&v609 == v609))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+836)) = v601 + int32(1)
	goto L73
L81:
	;
	v630 = v560
	v632 = int32(0)
	v635 = v564
	goto L82
L82:
	;
	v650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575)+2)))
	v651 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v575))))
	v652 = int32(16)
	v658 = v575 + (v650 | v651<<(uint(v652)%32)) + v632*int32(6)
	v659 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v658))))
	v662 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v658)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+816)) = v659<<(uint(v652)%32) | v662
	v669 = F_bsearch(m, v31+int32(816), v395, v472, int32(4), int32(208))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L25
	} else {
		goto L84
	}
L83:
	;
	goto L73
L84:
	;
	if v669 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v671 = *(*int32)(unsafe.Add(mBase, uint32(v658)))
	*(*int32)(unsafe.Add(mBase, uint32(v635))) = v671
	v673 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v658)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v635)+4)) = uint16(v673)
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v31)+836))
	*(*uint16)(unsafe.Add(mBase, uint32(v635)+6)) = uint16(v675)
	*(*uint16)(unsafe.Add(mBase, uint32(v630))) = uint16(v549)
	v678 = *(*int32)(unsafe.Add(mBase, uint32(v571)))
	v679 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v630)+4)) = uint16(v679)
	*(*uint8)(unsafe.Add(mBase, uint32(v630)+3)) = uint8(v679)
	v683 = int32(98304)
	*(*uint8)(unsafe.Add(mBase, uint32(v630)+2)) = uint8(base.B2i32(v678&v683 == v683))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v31)+836))
	*(*int32)(unsafe.Add(mBase, uint32(v31)+836)) = v688 + int32(1)
	v696 = v630 + int32(6)
	v697 = v635 + int32(8)
	goto L87
L86:
	;
	v696 = v630
	v697 = v635
	goto L87
L87:
	;
	v700 = v632 + int32(1)
	if v700 != v618 {
		v630 = v696
		v632 = v700
		v635 = v697
		goto L82
	} else {
		goto L88
	}
L88:
	;
	goto L83
L89:
	;
	goto L72
L90:
	;
	F__bt_delitems_delete_check(m, l0, v36, l1, v31+int32(820))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L25
	} else {
		goto L91
	}
L91:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v31)+840))
	F_pfree(m, v743)
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L25
	} else {
		goto L92
	}
L92:
	;
	v746 = *(*int32)(unsafe.Add(mBase, uint32(v31)+844))
	F_pfree(m, v746)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L25
	} else {
		goto L93
	}
L93:
	;
	v749 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v749)
	v752 = int32(4)
	v753 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+14)))
	v754 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v54)+12)))
	v755 = v753 - v754
	if v755 <= v752 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v761 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v761) <= base.Ui32(v758-int32(4)) {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	v758 = v752
	goto L97
L96:
	;
	v758 = v755
	goto L97
L97:
	;
	goto L94
L98:
	;
	v768 = int32(1)
	goto L2
L99:
	;
	if l3 != 0 {
		goto L1
	} else {
		goto L100
	}
L100:
	;
	v796 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v796)
	v798 = v768 | l6
	if v798 == int32(1) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v801 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v802 = m.G0
	v804 = v802 - int32(32)
	m.G0 = v804
	if v36 < int32(0) {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L103
L103:
	;
	v1782 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v1782 != 0 {
		goto L226
	} else {
		goto L227
	}
L104:
	;
	v824 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823)+16)))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v826 = int32(*(*int16)(unsafe.Add(mBase, uint32(v825)+10)))
	v828 = F_palloc(m, int32(1676))
	mBase = m.M
	v829 = m.ExcPending
	if v829 != 0 {
		goto L25
	} else {
		goto L108
	}
L105:
	;
	v809 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v809+(v36^int32(-1))<<(uint(int32(2))%32))))
	v823 = v815
	goto L104
L106:
	;
	goto L107
L107:
	;
	v817 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v823 = v817 + v36<<(uint(int32(13))%32) + int32(-8192)
	goto L104
L108:
	;
	v830 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+20)) = v830
	*(*uint16)(unsafe.Add(mBase, uint32(v828)+16)) = uint16(v830)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+12)) = v830
	*(*int64)(unsafe.Add(mBase, uint32(v828)+4)) = int64(35184372088832)
	v838 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v828))) = uint8(v838)
	v842 = F_palloc(m, int32(8192))
	mBase = m.M
	v843 = m.ExcPending
	if v843 != 0 {
		goto L25
	} else {
		goto L109
	}
L109:
	;
	v844 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v828)+28)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v828)+24)) = v842
	*(*int64)(unsafe.Add(mBase, uint32(v828)+36)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v804)+4)) = l0
	if v36 < int32(0) {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v804)+20)) = int32(0)
	v871 = int32(512)
	v873 = v801 + int32(4)
	if base.Ui32(v873) <= base.Ui32(v871) {
		goto L114
	} else {
		goto L115
	}
L111:
	;
	v853 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v853+(v36^int32(-1))<<(uint(int32(6))%32))+16))
	v868 = v859
	goto L110
L112:
	;
	goto L113
L113:
	;
	v861 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v861+v36<<(uint(int32(6))%32)+int32(-64))+16))
	v868 = v867
	goto L110
L114:
	;
	v876 = v871
	goto L116
L115:
	;
	v876 = v873
	goto L116
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v804)+16)) = v876
	v878 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v804)+12)) = uint8(v878)
	*(*int32)(unsafe.Add(mBase, uint32(v804)+8)) = v868
	v882 = F_palloc(m, int32(10864))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L25
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v804)+24)) = v882
	v886 = F_palloc(m, int32(8148))
	mBase = m.M
	v887 = m.ExcPending
	if v887 != 0 {
		goto L25
	} else {
		goto L118
	}
L118:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v804)+28)) = v886
	v892 = *(*int32)(unsafe.Add(mBase, uint32(v824+v823)+4))
	if v892 != 0 {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v893 = int32(2)
	goto L121
L120:
	;
	v893 = int32(1)
	goto L121
L121:
	;
	v894 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v894) {
		goto L122
	} else {
		goto L123
	}
L122:
	;
	v902 = int32(base.Ui32(v894+int32(262120)) >> (uint(int32(2)) % 32))
	goto L124
L123:
	;
	v902 = int32(0)
	goto L124
L124:
	;
	v904 = v902 & int32(65535)
	if base.Ui32(v893) <= base.Ui32(v904) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v907 = v828 + int32(44)
	v913 = v893
	goto L128
L126:
	;
	goto L127
L127:
	;
	v1450 = v804 + int32(4)
	v1451 = int32(0)
	v1465 = *(*int32)(unsafe.Add(mBase, uint32(v828)+32))
	if v1465 <= v1451 {
		goto L190
	} else {
		goto L191
	}
L128:
	;
	v939 = v913 & int32(65535)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v939<<(uint(int32(2))%32)+(v823+int32(24))-int32(4))))
	v948 = v823 + v945&int32(32767)
	if v893 == v939 {
		goto L131
	} else {
		goto L132
	}
L129:
	;
	goto L127
L130:
	;
	v1417 = v913 + int32(1)
	if base.Ui32(v1417&int32(65535)) <= base.Ui32(v904) {
		v913 = v1417
		goto L128
	} else {
		goto L188
	}
L131:
	;
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+7)))
	if v950&int32(32) != 0 {
		goto L136
	} else {
		goto L137
	}
L132:
	;
	goto L133
L133:
	;
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v828)+12))
	v1009 = F__bt_keep_natts_fast(m, l0, v1008, v948)
	mBase = m.M
	v1010 = m.ExcPending
	if v1010 != 0 {
		goto L25
	} else {
		goto L146
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+20)) = v987
	*(*uint16)(unsafe.Add(mBase, uint32(v828)+16)) = uint16(v893)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+12)) = v948
	v993 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+36)) = (v993&int32(8191)+int32(7))&int32(16376) | int32(4)
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v828)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v907+v1003<<(uint(int32(2))%32)))) = uint16(v893)
	goto L130
L135:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v828)+24))
	v968 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+2)))
	v969 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948))))
	v975 = v953 & int32(4095)
	v977 = v975 * int32(6)
	if v977 != 0 {
		goto L141
	} else {
		goto L142
	}
L136:
	;
	v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+4)))
	if v953&int32(8192) != 0 {
		goto L135
	} else {
		goto L139
	}
L137:
	;
	goto L138
L138:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v828)+24))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
	*(*int32)(unsafe.Add(mBase, uint32(v957))) = v958
	v960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v957)+4)) = uint16(v960)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+28)) = int32(1)
	v964 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+6)))
	v987 = v964 & int32(8191)
	goto L134
L139:
	;
	goto L138
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828)+28)) = v975
	v981 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+2)))
	v982 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948))))
	v987 = v981 | v982<<(uint(int32(16))%32)
	goto L134
L141:
	;
	v978 = F__emscripten_memcpy_bulkmem(m, v967, v948+(v968|v969<<(uint(int32(16))%32)), v977)
	mBase = m.M
	goto L143
L142:
	;
	goto L143
L143:
	;
	goto L140
L144:
	;
	v1386 = *(*int32)(unsafe.Add(mBase, uint32(v828)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+32)) = v1386 + int32(1)
	v1390 = *(*int32)(unsafe.Add(mBase, uint32(v828)+24))
	v1391 = int32(6)
	v1395 = v1031 * v1391
	if v1395 != 0 {
		goto L185
	} else {
		goto L186
	}
L145:
	;
	v1056 = v804 + int32(4)
	v1057 = int32(0)
	v1071 = *(*int32)(unsafe.Add(mBase, uint32(v828)+32))
	if v1071 <= v1057 {
		goto L154
	} else {
		goto L155
	}
L146:
	;
	if v1009 <= v826 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v1012 = int32(1)
	v1013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+7)))
	if v1013&int32(32) == int32(0) {
		v1031 = v1012
		v1033 = v948
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(v828)+8))
	v1035 = *(*int32)(unsafe.Add(mBase, uint32(v828)+20))
	v1036 = *(*int32)(unsafe.Add(mBase, uint32(v828)+28))
	if base.Ui32((v1035+(v1036+v1031)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1034) {
		goto L144
	} else {
		goto L151
	}
L149:
	;
	v1018 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+4)))
	if v1018&int32(8192) == int32(0) {
		v1031 = v1012
		v1033 = v948
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v1025 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+2)))
	v1026 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948))))
	v1031 = v1018 & int32(4095)
	v1033 = v948 + (v1025 | v1026<<(uint(int32(16))%32))
	goto L148
L151:
	;
	if v1036 <= int32(50) {
		goto L145
	} else {
		goto L152
	}
L152:
	;
	v1048 = *(*int32)(unsafe.Add(mBase, uint32(v828)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+4)) = v1048 + int32(1)
	goto L145
L153:
	;
	v1328 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+7)))
	if v1328&int32(32) != 0 {
		goto L176
	} else {
		goto L177
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v828)+28)) = int64(0)
	goto L153
L155:
	;
	v1078 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+16))
	v1086 = v1078
	v1090 = v1057
	goto L156
L156:
	;
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+24))
	v1099 = v1096 + v1086*int32(6)
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+20))
	v1103 = v1100 + v1086<<(uint(int32(3))%32)
	v1104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v828)+16)))
	v1105 = v1104 + v1090
	v1112 = v1105&int32(65535)<<(uint(int32(2))%32) + (v823 + int32(24)) - int32(4)
	v1113 = *(*int32)(unsafe.Add(mBase, uint32(v1112)))
	v1116 = v823 + v1113&int32(32767)
	v1117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1116)+7)))
	if v1117&int32(32) != 0 {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	if v1071 <= int32(1) {
		goto L154
	} else {
		goto L173
	}
L158:
	;
	v1293 = v1090 + int32(1)
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(v828)+32))
	if v1293 < v1294 {
		v1086 = v1282
		v1090 = v1293
		goto L156
	} else {
		goto L172
	}
L159:
	;
	v1145 = v1120 & int32(4095)
	v1146 = int32(0)
	if int32(2) <= v1071 {
		goto L164
	} else {
		goto L165
	}
L160:
	;
	v1120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116)+4)))
	if v1120&int32(8192) != 0 {
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v1124 = *(*int32)(unsafe.Add(mBase, uint32(v1116)))
	*(*int32)(unsafe.Add(mBase, uint32(v1103))) = v1124
	v1126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1103)+4)) = uint16(v1126)
	v1128 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1103)+6)) = uint16(v1128)
	*(*uint8)(unsafe.Add(mBase, uint32(v1099)+3)) = uint8(base.B2i32(int32(1) < v1071))
	v1131 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1099)+2)) = uint8(v1131)
	*(*uint16)(unsafe.Add(mBase, uint32(v1099))) = uint16(v1105)
	v1134 = *(*int32)(unsafe.Add(mBase, uint32(v1112)))
	v1138 = int32(base.Ui32(v1134)>>(uint(int32(17))%32)) + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1099)+4)) = uint16(v1138)
	v1140 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+16))
	v1142 = v1140 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1056)+16)) = v1142
	v1282 = v1142
	goto L158
L163:
	;
	goto L162
L164:
	;
	v1150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116)+2)))
	v1151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116))))
	v1152 = int32(16)
	v1155 = v1116 + (v1150 | v1151<<(uint(v1152)%32))
	v1156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1155))))
	v1159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1155)+2)))
	v1160 = v1156<<(uint(v1152)%32) | v1159
	v1163 = int32(6)
	v1165 = v1155 + int32(base.Ui32(v1145)>>(uint(int32(1))%32))*v1163
	v1166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1165))))
	v1169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1165)+2)))
	v1170 = v1166<<(uint(v1152)%32) | v1169
	v1174 = v1155 + v1145*v1163
	v1177 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1174-v1163))))
	v1182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1174-int32(4)))))
	v1187 = base.B2i32(v1170 == v1160)
	v1189 = base.B2i32(v1160 != v1170) & base.B2i32(v1170 == v1177<<(uint(v1152)%32)|v1182)
	goto L166
L165:
	;
	v1187 = v1146
	v1189 = v1146
	goto L166
L166:
	;
	if v1145 == int32(0) {
		v1282 = v1086
		goto L158
	} else {
		goto L167
	}
L167:
	;
	v1193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116)+2)))
	v1194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116))))
	v1198 = v1116 + (v1193 | v1194<<(uint(int32(16))%32))
	v1199 = *(*int32)(unsafe.Add(mBase, uint32(v1198)))
	*(*int32)(unsafe.Add(mBase, uint32(v1103))) = v1199
	v1201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1198)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1103)+4)) = uint16(v1201)
	v1203 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1103)+6)) = uint16(v1203)
	v1205 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v1099)+4)) = uint16(v1205)
	v1207 = int32(1)
	v1208 = v1145 - v1207
	v1209 = int32(0)
	v1212 = v1187 | v1189&base.B2i32(v1208 == v1209)
	*(*uint8)(unsafe.Add(mBase, uint32(v1099)+3)) = uint8(v1212)
	*(*uint8)(unsafe.Add(mBase, uint32(v1099)+2)) = uint8(v1209)
	*(*uint16)(unsafe.Add(mBase, uint32(v1099))) = uint16(v1105)
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+16))
	v1220 = v1218 + v1207
	*(*int32)(unsafe.Add(mBase, uint32(v1056)+16)) = v1220
	if v1145 == v1207 {
		v1282 = v1220
		goto L158
	} else {
		goto L168
	}
L168:
	;
	v1227 = v1207
	v1228 = v1099
	v1230 = v1103
	goto L169
L169:
	;
	v1242 = v1230 + int32(8)
	v1243 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116)+2)))
	v1244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1116))))
	v1249 = int32(6)
	v1251 = v1116 + (v1243 | v1244<<(uint(int32(16))%32)) + v1227*v1249
	v1252 = *(*int32)(unsafe.Add(mBase, uint32(v1251)))
	*(*int32)(unsafe.Add(mBase, uint32(v1242))) = v1252
	v1254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1251)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1230)+12)) = uint16(v1254)
	v1256 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1230)+14)) = uint16(v1256)
	*(*uint16)(unsafe.Add(mBase, uint32(v1228)+10)) = uint16(v1249)
	v1261 = v1189 & base.B2i32(v1227 == v1208)
	*(*uint8)(unsafe.Add(mBase, uint32(v1228)+9)) = uint8(v1261)
	v1263 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1228)+8)) = uint8(v1263)
	v1266 = v1228 + v1249
	*(*uint16)(unsafe.Add(mBase, uint32(v1266))) = uint16(v1105)
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v1056)+16))
	v1269 = int32(1)
	v1270 = v1268 + v1269
	*(*int32)(unsafe.Add(mBase, uint32(v1056)+16)) = v1270
	v1273 = v1227 + v1269
	if v1273 != v1145 {
		v1227 = v1273
		v1228 = v1266
		v1230 = v1242
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v1282 = v1270
	goto L158
L171:
	;
	goto L170
L172:
	;
	goto L157
L173:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(v828)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v828+v1298<<(uint(int32(2))%32))+46)) = uint16(v1294)
	v1303 = *(*int32)(unsafe.Add(mBase, uint32(v828)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+40)) = v1303 + int32(1)
	goto L154
L174:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+20)) = v1365
	*(*uint16)(unsafe.Add(mBase, uint32(v828)+16)) = uint16(v913)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+12)) = v948
	v1371 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+36)) = (v1371&int32(8191)+int32(7))&int32(16376) | int32(4)
	v1381 = *(*int32)(unsafe.Add(mBase, uint32(v828)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v907+v1381<<(uint(int32(2))%32)))) = uint16(v913)
	goto L130
L175:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v828)+24))
	v1346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+2)))
	v1347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948))))
	v1353 = v1331 & int32(4095)
	v1355 = v1353 * int32(6)
	if v1355 != 0 {
		goto L181
	} else {
		goto L182
	}
L176:
	;
	v1331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+4)))
	if v1331&int32(8192) != 0 {
		goto L175
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v1335 = *(*int32)(unsafe.Add(mBase, uint32(v828)+24))
	v1336 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
	*(*int32)(unsafe.Add(mBase, uint32(v1335))) = v1336
	v1338 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1335)+4)) = uint16(v1338)
	*(*int32)(unsafe.Add(mBase, uint32(v828)+28)) = int32(1)
	v1342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+6)))
	v1365 = v1342 & int32(8191)
	goto L174
L179:
	;
	goto L178
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828)+28)) = v1353
	v1359 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+2)))
	v1360 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948))))
	v1365 = v1359 | v1360<<(uint(int32(16))%32)
	goto L174
L181:
	;
	v1356 = F__emscripten_memcpy_bulkmem(m, v1345, v948+(v1346|v1347<<(uint(int32(16))%32)), v1355)
	mBase = m.M
	goto L183
L182:
	;
	goto L183
L183:
	;
	goto L180
L184:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v828)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+28)) = v1398 + v1031
	v1401 = *(*int32)(unsafe.Add(mBase, uint32(v828)+36))
	v1402 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v948)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+36)) = v1401 + (v1402&int32(8191)+int32(7))&int32(16376) + int32(4)
	goto L130
L185:
	;
	v1396 = F__emscripten_memcpy_bulkmem(m, v1390+v1036*v1391, v1033, v1395)
	mBase = m.M
	goto L187
L186:
	;
	goto L187
L187:
	;
	goto L184
L188:
	;
	goto L129
L189:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v828)+40))
	v1723 = *(*int32)(unsafe.Add(mBase, uint32(v828)+24))
	F_pfree(m, v1723)
	mBase = m.M
	v1725 = m.ExcPending
	if v1725 != 0 {
		goto L25
	} else {
		goto L210
	}
L190:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v828)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v828)+28)) = int64(0)
	goto L189
L191:
	;
	v1472 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+16))
	v1480 = v1472
	v1484 = v1451
	goto L192
L192:
	;
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+24))
	v1493 = v1490 + v1480*int32(6)
	v1494 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+20))
	v1497 = v1494 + v1480<<(uint(int32(3))%32)
	v1498 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v828)+16)))
	v1499 = v1498 + v1484
	v1506 = v1499&int32(65535)<<(uint(int32(2))%32) + (v823 + int32(24)) - int32(4)
	v1507 = *(*int32)(unsafe.Add(mBase, uint32(v1506)))
	v1510 = v823 + v1507&int32(32767)
	v1511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1510)+7)))
	if v1511&int32(32) != 0 {
		goto L196
	} else {
		goto L197
	}
L193:
	;
	if v1465 <= int32(1) {
		goto L190
	} else {
		goto L209
	}
L194:
	;
	v1687 = v1484 + int32(1)
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v828)+32))
	if v1687 < v1688 {
		v1480 = v1676
		v1484 = v1687
		goto L192
	} else {
		goto L208
	}
L195:
	;
	v1539 = v1514 & int32(4095)
	v1540 = int32(0)
	if int32(2) <= v1465 {
		goto L200
	} else {
		goto L201
	}
L196:
	;
	v1514 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1510)+4)))
	if v1514&int32(8192) != 0 {
		goto L195
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	v1518 = *(*int32)(unsafe.Add(mBase, uint32(v1510)))
	*(*int32)(unsafe.Add(mBase, uint32(v1497))) = v1518
	v1520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1510)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1497)+4)) = uint16(v1520)
	v1522 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1497)+6)) = uint16(v1522)
	*(*uint8)(unsafe.Add(mBase, uint32(v1493)+3)) = uint8(base.B2i32(int32(1) < v1465))
	v1525 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1493)+2)) = uint8(v1525)
	*(*uint16)(unsafe.Add(mBase, uint32(v1493))) = uint16(v1499)
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1506)))
	v1532 = int32(base.Ui32(v1528)>>(uint(int32(17))%32)) + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1493)+4)) = uint16(v1532)
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+16))
	v1536 = v1534 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+16)) = v1536
	v1676 = v1536
	goto L194
L199:
	;
	goto L198
L200:
	;
	v1544 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1510)+2)))
	v1545 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1510))))
	v1546 = int32(16)
	v1549 = v1510 + (v1544 | v1545<<(uint(v1546)%32))
	v1550 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1549))))
	v1553 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1549)+2)))
	v1554 = v1550<<(uint(v1546)%32) | v1553
	v1557 = int32(6)
	v1559 = v1549 + int32(base.Ui32(v1539)>>(uint(int32(1))%32))*v1557
	v1560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1559))))
	v1563 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1559)+2)))
	v1564 = v1560<<(uint(v1546)%32) | v1563
	v1568 = v1549 + v1539*v1557
	v1571 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1568-v1557))))
	v1576 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1568-int32(4)))))
	v1581 = base.B2i32(v1564 == v1554)
	v1583 = base.B2i32(v1554 != v1564) & base.B2i32(v1564 == v1571<<(uint(v1546)%32)|v1576)
	goto L202
L201:
	;
	v1581 = v1540
	v1583 = v1540
	goto L202
L202:
	;
	if v1539 == int32(0) {
		v1676 = v1480
		goto L194
	} else {
		goto L203
	}
L203:
	;
	v1587 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1510)+2)))
	v1588 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1510))))
	v1592 = v1510 + (v1587 | v1588<<(uint(int32(16))%32))
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(v1592)))
	*(*int32)(unsafe.Add(mBase, uint32(v1497))) = v1593
	v1595 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1592)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1497)+4)) = uint16(v1595)
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1497)+6)) = uint16(v1597)
	v1599 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v1493)+4)) = uint16(v1599)
	v1601 = int32(1)
	v1602 = v1539 - v1601
	v1603 = int32(0)
	v1606 = v1581 | v1583&base.B2i32(v1602 == v1603)
	*(*uint8)(unsafe.Add(mBase, uint32(v1493)+3)) = uint8(v1606)
	*(*uint8)(unsafe.Add(mBase, uint32(v1493)+2)) = uint8(v1603)
	*(*uint16)(unsafe.Add(mBase, uint32(v1493))) = uint16(v1499)
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+16))
	v1614 = v1612 + v1601
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+16)) = v1614
	if v1539 == v1601 {
		v1676 = v1614
		goto L194
	} else {
		goto L204
	}
L204:
	;
	v1621 = v1601
	v1622 = v1493
	v1624 = v1497
	goto L205
L205:
	;
	v1636 = v1624 + int32(8)
	v1637 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1510)+2)))
	v1638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1510))))
	v1643 = int32(6)
	v1645 = v1510 + (v1637 | v1638<<(uint(int32(16))%32)) + v1621*v1643
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v1645)))
	*(*int32)(unsafe.Add(mBase, uint32(v1636))) = v1646
	v1648 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1645)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1624)+12)) = uint16(v1648)
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1624)+14)) = uint16(v1650)
	*(*uint16)(unsafe.Add(mBase, uint32(v1622)+10)) = uint16(v1643)
	v1655 = v1583 & base.B2i32(v1621 == v1602)
	*(*uint8)(unsafe.Add(mBase, uint32(v1622)+9)) = uint8(v1655)
	v1657 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1622)+8)) = uint8(v1657)
	v1660 = v1622 + v1643
	*(*uint16)(unsafe.Add(mBase, uint32(v1660))) = uint16(v1499)
	v1662 = *(*int32)(unsafe.Add(mBase, uint32(v1450)+16))
	v1663 = int32(1)
	v1664 = v1662 + v1663
	*(*int32)(unsafe.Add(mBase, uint32(v1450)+16)) = v1664
	v1667 = v1621 + v1663
	if v1667 != v1539 {
		v1621 = v1667
		v1622 = v1660
		v1624 = v1636
		goto L205
	} else {
		goto L207
	}
L206:
	;
	v1676 = v1664
	goto L194
L207:
	;
	goto L206
L208:
	;
	goto L193
L209:
	;
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v828)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v828+v1692<<(uint(int32(2))%32))+46)) = uint16(v1688)
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v828)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v828)+40)) = v1697 + int32(1)
	goto L190
L210:
	;
	F_pfree(m, v828)
	mBase = m.M
	v1727 = m.ExcPending
	if v1727 != 0 {
		goto L25
	} else {
		goto L211
	}
L211:
	;
	F__bt_delitems_delete_check(m, l0, v36, l1, v804+int32(4))
	mBase = m.M
	v1731 = m.ExcPending
	if v1731 != 0 {
		goto L25
	} else {
		goto L212
	}
L212:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v804)+24))
	F_pfree(m, v1732)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L25
	} else {
		goto L213
	}
L213:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v804)+28))
	F_pfree(m, v1735)
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L25
	} else {
		goto L214
	}
L214:
	;
	if v1722 != 0 {
		goto L215
	} else {
		goto L216
	}
L215:
	;
	v1738 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823)+14)))
	v1739 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v823)+12)))
	v1740 = v1738 - v1739
	v1741 = int32(0)
	if v1741 < v1740 {
		goto L219
	} else {
		goto L220
	}
L216:
	;
	v1750 = v838
	goto L217
L217:
	;
	m.G0 = v804 + int32(32)
	if v1750 != 0 {
		goto L1
	} else {
		goto L225
	}
L218:
	;
	v1745 = int32(341)
	if base.Ui32(v873) <= base.Ui32(v1745) {
		goto L222
	} else {
		goto L223
	}
L219:
	;
	v1744 = v1740
	goto L221
L220:
	;
	v1744 = v1741
	goto L221
L221:
	;
	goto L218
L222:
	;
	v1748 = v1745
	goto L224
L223:
	;
	v1748 = v873
	goto L224
L224:
	;
	v1750 = base.B2i32(base.Ui32(v1748) <= base.Ui32(v1744))
	goto L217
L225:
	;
	goto L103
L226:
	;
	v1783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1782)+16)))
	if v1783 == int32(0) {
		goto L1
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v1786 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v33)+1)))
	if v1786 != int32(1) {
		goto L1
	} else {
		goto L230
	}
L229:
	;
	goto L228
L230:
	;
	v1789 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1790 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1791 = int32(0)
	v1792 = m.G0
	v1794 = v1792 - int32(16)
	m.G0 = v1794
	v1796 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v1797 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1796)+10)))
	if v36 < v1791 {
		goto L232
	} else {
		goto L233
	}
L231:
	;
	v1816 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1815)+16)))
	v1818 = F_palloc(m, int32(1676))
	mBase = m.M
	v1819 = m.ExcPending
	if v1819 != 0 {
		goto L25
	} else {
		goto L235
	}
L232:
	;
	v1801 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1801+(v36^int32(-1))<<(uint(int32(2))%32))))
	v1815 = v1807
	goto L231
L233:
	;
	goto L234
L234:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v1815 = v1809 + v36<<(uint(int32(13))%32) + int32(-8192)
	goto L231
L235:
	;
	v1820 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+20)) = v1820
	*(*uint16)(unsafe.Add(mBase, uint32(v1818)+16)) = uint16(v1820)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+12)) = v1820
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+4)) = int64(5806795784192)
	v1828 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1818))) = uint8(v1828)
	v1831 = F_palloc(m, int32(1352))
	mBase = m.M
	v1832 = m.ExcPending
	if v1832 != 0 {
		goto L25
	} else {
		goto L236
	}
L236:
	;
	v1833 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+28)) = v1833
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+24)) = v1831
	*(*int64)(unsafe.Add(mBase, uint32(v1818)+36)) = v1833
	v1838 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1815)+12)))
	v1841 = v1816 + v1815
	v1842 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+4))
	if v1842 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1843 = int32(2)
	goto L239
L238:
	;
	v1843 = int32(1)
	goto L239
L239:
	;
	if v798 != 0 {
		v1887 = v1791
		goto L240
	} else {
		goto L241
	}
L240:
	;
	v1891 = F_PageGetTempPageCopySpecial(m, v1815)
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L25
	} else {
		goto L251
	}
L241:
	;
	v1844 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v1845 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1844)+10)))
	v1847 = v1815 + int32(24)
	v1853 = *(*int32)(unsafe.Add(mBase, uint32(v1847+v1843<<(uint(int32(2))%32)-int32(4))))
	v1857 = F__bt_keep_natts_fast(m, l0, v1789, v1815+v1853&int32(32767))
	mBase = m.M
	v1858 = m.ExcPending
	if v1858 != 0 {
		goto L25
	} else {
		goto L242
	}
L242:
	;
	if v1845 < v1857 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	v1860 = int32(1)
	v1862 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1815)+12)))
	if base.Ui32(v1862) < base.Ui32(int32(25)) {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	goto L245
L245:
	;
	v1887 = int32(0)
	goto L240
L246:
	;
	v1873 = int32(-1)
	goto L248
L247:
	;
	v1873 = int32(base.Ui32(v1862+int32(262120))>>(uint(int32(2))%32))&int32(65535) - v1860
	goto L248
L248:
	;
	v1877 = *(*int32)(unsafe.Add(mBase, uint32(v1847+v1873<<(uint(int32(2))%32))))
	v1881 = F__bt_keep_natts_fast(m, l0, v1789, v1815+v1877&int32(32767))
	mBase = m.M
	v1882 = m.ExcPending
	if v1882 != 0 {
		goto L25
	} else {
		goto L249
	}
L249:
	;
	if v1845 < v1881 {
		v1887 = v1860
		goto L240
	} else {
		goto L250
	}
L250:
	;
	goto L245
L251:
	;
	v1893 = *(*int64)(unsafe.Add(mBase, uint32(v1815)))
	*(*int64)(unsafe.Add(mBase, uint32(v1891))) = v1893
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v1841)+4))
	if v1895 != 0 {
		goto L254
	} else {
		goto L255
	}
L252:
	;
	goto L1
L253:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2318 = m.ExcPending
	if v2318 != 0 {
		goto L25
	} else {
		goto L344
	}
L254:
	;
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v1815)+24))
	v1904 = F_PageAddItemExtended(m, v1891, v1815+v1896&int32(32767), int32(base.Ui32(v1896)>>(uint(int32(17))%32)), int32(1), int32(0))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L25
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1838) {
		goto L259
	} else {
		goto L260
	}
L257:
	;
	if v1904 == int32(0) {
		goto L253
	} else {
		goto L258
	}
L258:
	;
	goto L256
L259:
	;
	v1916 = int32(base.Ui32(v1838+int32(262120)) >> (uint(int32(2)) % 32))
	goto L261
L260:
	;
	v1916 = int32(0)
	goto L261
L261:
	;
	v1918 = v1916 & int32(65535)
	if base.Ui32(v1843) <= base.Ui32(v1918) {
		goto L262
	} else {
		goto L263
	}
L262:
	;
	v1921 = v1818 + int32(44)
	v1925 = v1887
	v1932 = v1843
	goto L265
L263:
	;
	goto L264
L264:
	;
	F__bt_dedup_finish_pending(m, v1891, v1818)
	mBase = m.M
	v2233 = m.ExcPending
	if v2233 != 0 {
		goto L25
	} else {
		goto L318
	}
L265:
	;
	v1953 = v1932 & int32(65535)
	v1959 = *(*int32)(unsafe.Add(mBase, uint32(v1953<<(uint(int32(2))%32)+(v1815+int32(24))-int32(4))))
	v1962 = v1815 + v1959&int32(32767)
	if v1843 == v1953 {
		goto L268
	} else {
		goto L269
	}
L266:
	;
	goto L264
L267:
	;
	v2200 = v1932 + int32(1)
	if base.Ui32(v2200&int32(65535)) <= base.Ui32(v1918) {
		v1925 = v2194
		v1932 = v2200
		goto L265
	} else {
		goto L317
	}
L268:
	;
	v1964 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1962)+7)))
	if v1964&int32(32) != 0 {
		goto L273
	} else {
		goto L274
	}
L269:
	;
	goto L270
L270:
	;
	v2022 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1818))))
	if v2022 != int32(1) {
		goto L282
	} else {
		goto L283
	}
L271:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+20)) = v2001
	*(*uint16)(unsafe.Add(mBase, uint32(v1818)+16)) = uint16(v1843)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+12)) = v1962
	v2007 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+36)) = (v2007&int32(8191)+int32(7))&int32(16376) | int32(4)
	v2017 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1921+v2017<<(uint(int32(2))%32)))) = uint16(v1843)
	v2194 = v1925
	goto L267
L272:
	;
	v1981 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+24))
	v1982 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+2)))
	v1983 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962))))
	v1989 = v1967 & int32(4095)
	v1991 = v1989 * int32(6)
	if v1991 != 0 {
		goto L278
	} else {
		goto L279
	}
L273:
	;
	v1967 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+4)))
	if v1967&int32(8192) != 0 {
		goto L272
	} else {
		goto L276
	}
L274:
	;
	goto L275
L275:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+24))
	v1972 = *(*int32)(unsafe.Add(mBase, uint32(v1962)))
	*(*int32)(unsafe.Add(mBase, uint32(v1971))) = v1972
	v1974 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1971)+4)) = uint16(v1974)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+28)) = int32(1)
	v1978 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+6)))
	v2001 = v1978 & int32(8191)
	goto L271
L276:
	;
	goto L275
L277:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+28)) = v1989
	v1995 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+2)))
	v1996 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962))))
	v2001 = v1995 | v1996<<(uint(int32(16))%32)
	goto L271
L278:
	;
	v1992 = F__emscripten_memcpy_bulkmem(m, v1981, v1962+(v1982|v1983<<(uint(int32(16))%32)), v1991)
	mBase = m.M
	goto L280
L279:
	;
	goto L280
L280:
	;
	goto L277
L281:
	;
	v2167 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+32)) = v2167 + int32(1)
	v2171 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+24))
	v2172 = int32(6)
	v2176 = v2048 * v2172
	if v2176 != 0 {
		goto L314
	} else {
		goto L315
	}
L282:
	;
	F__bt_dedup_finish_pending(m, v1891, v1818)
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L25
	} else {
		goto L291
	}
L283:
	;
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+12))
	v2026 = F__bt_keep_natts_fast(m, l0, v2025, v1962)
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L25
	} else {
		goto L284
	}
L284:
	;
	if v2026 <= v1797 {
		goto L282
	} else {
		goto L285
	}
L285:
	;
	v2029 = int32(1)
	v2030 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1962)+7)))
	if v2030&int32(32) == int32(0) {
		v2048 = v2029
		v2050 = v1962
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+8))
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+20))
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+28))
	if base.Ui32((v2052+(v2053+v2048)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v2051) {
		goto L281
	} else {
		goto L289
	}
L287:
	;
	v2035 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+4)))
	if v2035&int32(8192) == int32(0) {
		v2048 = v2029
		v2050 = v1962
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v2042 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+2)))
	v2043 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962))))
	v2048 = v2035 & int32(4095)
	v2050 = v1962 + (v2042 | v2043<<(uint(int32(16))%32))
	goto L286
L289:
	;
	if v2053 <= int32(50) {
		goto L282
	} else {
		goto L290
	}
L290:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+4)) = v2065 + int32(1)
	goto L282
L291:
	;
	v2074 = int32(0)
	if v1925 == v2074 {
		v2105 = v2074
		goto L292
	} else {
		goto L293
	}
L292:
	;
	v2109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1962)+7)))
	if v2109&int32(32) != 0 {
		goto L305
	} else {
		goto L306
	}
L293:
	;
	v2077 = int32(1)
	v2078 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+4))
	switch v2078 - int32(5) {
	case 0:
		goto L295
	case 1:
		goto L294
	default:
		v2105 = v2077
		goto L292
	}
L294:
	;
	v2102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1818))) = uint8(v2102)
	v2105 = v2102
	goto L292
L295:
	;
	v2081 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+8))
	v2082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1815)+19)))
	v2090 = base.F64_mul(base.F64_convert_i32_u(v2082<<(uint(int32(8))%32)-v1790-int32(52)), float64(0.04))
	if base.F64_lt(base.F64_abs(v2090), float64(2.147483648e+09)) != 0 {
		goto L297
	} else {
		goto L298
	}
L296:
	;
	v2097 = v2081 - v2096
	if base.Ui32(v2097) <= base.Ui32(v2081) {
		goto L300
	} else {
		goto L301
	}
L297:
	;
	v2094 = base.I32_trunc_f64_s(v2090)
	v2096 = v2094
	goto L296
L298:
	;
	goto L299
L299:
	;
	v2096 = int32(-2147483648)
	goto L296
L300:
	;
	v2100 = v2097
	goto L302
L301:
	;
	v2100 = int32(0)
	goto L302
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+8)) = v2100
	v2105 = v2077
	goto L292
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+20)) = v2146
	*(*uint16)(unsafe.Add(mBase, uint32(v1818)+16)) = uint16(v1932)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+12)) = v1962
	v2152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+36)) = (v2152&int32(8191)+int32(7))&int32(16376) | int32(4)
	v2162 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1921+v2162<<(uint(int32(2))%32)))) = uint16(v1932)
	v2194 = v2105
	goto L267
L304:
	;
	v2126 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+24))
	v2127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+2)))
	v2128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962))))
	v2134 = v2112 & int32(4095)
	v2136 = v2134 * int32(6)
	if v2136 != 0 {
		goto L310
	} else {
		goto L311
	}
L305:
	;
	v2112 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+4)))
	if v2112&int32(8192) != 0 {
		goto L304
	} else {
		goto L308
	}
L306:
	;
	goto L307
L307:
	;
	v2116 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+24))
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v1962)))
	*(*int32)(unsafe.Add(mBase, uint32(v2116))) = v2117
	v2119 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2116)+4)) = uint16(v2119)
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+28)) = int32(1)
	v2123 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+6)))
	v2146 = v2123 & int32(8191)
	goto L303
L308:
	;
	goto L307
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+28)) = v2134
	v2140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+2)))
	v2141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962))))
	v2146 = v2140 | v2141<<(uint(int32(16))%32)
	goto L303
L310:
	;
	v2137 = F__emscripten_memcpy_bulkmem(m, v2126, v1962+(v2127|v2128<<(uint(int32(16))%32)), v2136)
	mBase = m.M
	goto L312
L311:
	;
	goto L312
L312:
	;
	goto L309
L313:
	;
	v2179 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+28)) = v2179 + v2048
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+36))
	v2183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1962)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1818)+36)) = v2182 + (v2183&int32(8191)+int32(7))&int32(16376) + int32(4)
	v2194 = v1925
	goto L267
L314:
	;
	v2177 = F__emscripten_memcpy_bulkmem(m, v2171+v2053*v2172, v2050, v2176)
	mBase = m.M
	goto L316
L315:
	;
	goto L316
L316:
	;
	goto L313
L317:
	;
	goto L266
L318:
	;
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+40))
	if v2234 == int32(0) {
		goto L320
	} else {
		goto L321
	}
L319:
	;
	F_pfree(m, v1818)
	mBase = m.M
	v2311 = m.ExcPending
	if v2311 != 0 {
		goto L25
	} else {
		goto L343
	}
L320:
	;
	F_pfree(m, v1891)
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L25
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	v2242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1841)+12)))
	if v2242&int32(64) != 0 {
		goto L325
	} else {
		goto L326
	}
L323:
	;
	v2239 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+24))
	F_pfree(m, v2239)
	mBase = m.M
	v2241 = m.ExcPending
	if v2241 != 0 {
		goto L25
	} else {
		goto L324
	}
L324:
	;
	goto L319
L325:
	;
	v2245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1891)+16)))
	v2246 = v1891 + v2245
	v2247 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2246)+12)))
	v2249 = v2247 & int32(65471)
	*(*uint16)(unsafe.Add(mBase, uint32(v2246)+12)) = uint16(v2249)
	goto L327
L326:
	;
	goto L327
L327:
	;
	v2252 = int32(4510148)
	v2254 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v2254 + int32(1)
	F_PageRestoreTempPage(m, v1891, v1815)
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L25
	} else {
		goto L328
	}
L328:
	;
	F_MarkBufferDirty(m, v36)
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L25
	} else {
		goto L329
	}
L329:
	;
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2263 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2262)+118)))
	if v2263 != int32(112) {
		goto L330
	} else {
		goto L331
	}
L330:
	;
	v2300 = int32(4510148)
	v2302 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v2302 - int32(1)
	v2306 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+24))
	F_pfree(m, v2306)
	mBase = m.M
	v2308 = m.ExcPending
	if v2308 != 0 {
		goto L25
	} else {
		goto L342
	}
L331:
	;
	v2267 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v2267 <= int32(0) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2270 != 0 {
		goto L330
	} else {
		goto L335
	}
L333:
	;
	goto L334
L334:
	;
	v2272 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1794)+14)) = uint16(v2272)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2275 = m.ExcPending
	if v2275 != 0 {
		goto L25
	} else {
		goto L337
	}
L335:
	;
	v2271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2271 != 0 {
		goto L330
	} else {
		goto L336
	}
L336:
	;
	goto L334
L337:
	;
	F_XLogRegisterBuffer(m, int32(0), v36, int32(8))
	mBase = m.M
	v2279 = m.ExcPending
	if v2279 != 0 {
		goto L25
	} else {
		goto L338
	}
L338:
	;
	F_XLogRegisterData(m, v1794+int32(14), int32(2))
	mBase = m.M
	v2284 = m.ExcPending
	if v2284 != 0 {
		goto L25
	} else {
		goto L339
	}
L339:
	;
	v2288 = *(*int32)(unsafe.Add(mBase, uint32(v1818)+40))
	F_XLogRegisterBufData(m, int32(0), v1818+int32(44), v2288<<(uint(int32(2))%32))
	mBase = m.M
	v2292 = m.ExcPending
	if v2292 != 0 {
		goto L25
	} else {
		goto L340
	}
L340:
	;
	v2295 = F_XLogInsert(m, int32(11), int32(96))
	mBase = m.M
	v2296 = m.ExcPending
	if v2296 != 0 {
		goto L25
	} else {
		goto L341
	}
L341:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1815))) = base.I64_rotr(v2295, int64(32))
	goto L330
L342:
	;
	goto L319
L343:
	;
	m.G0 = v1794 + int32(16)
	goto L252
L344:
	;
	F_errmsg_internal(m, int32(21111), int32(0))
	mBase = m.M
	v2322 = m.ExcPending
	if v2322 != 0 {
		goto L25
	} else {
		goto L345
	}
L345:
	;
	F_errfinish(m, int32(495851), int32(130), int32(130491))
	mBase = m.M
	v2327 = m.ExcPending
	if v2327 != 0 {
		goto L25
	} else {
		goto L346
	}
L346:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_getbuf(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	v4 = F_ReadBuffer(m, l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		F_LockBuffer(m, v4, l2)
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			F__bt_checkpage(m, l0, v4)
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v4
			}
		}
	}
}
func F__bt_getrootheight(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v43 int64
	_ = v43
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v5 != 0 {
		v54 = v5
		v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
		return v55
	} else {
		v7 = F_ReadBuffer(m, l0, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_LockBuffer(m, v7, int32(1))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F__bt_checkpage(m, l0, v7)
				mBase = m.M
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					v16 = F__bt_getmeta(m, l0, v7)
					mBase = m.M
					v17 = m.ExcPending
					if v17 != 0 {
						return int32(0)
					} else {
						v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
						if v18 == int32(0) {
							F_LockBuffer(m, v7, int32(0))
							mBase = m.M
							v23 = m.ExcPending
							if v23 != 0 {
								return int32(0)
							} else {
								F_ReleaseBuffer(m, v7)
								mBase = m.M
								v25 = m.ExcPending
								if v25 != 0 {
									return int32(0)
								} else {
									return int32(0)
								}
							}
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							v30 = F_MemoryContextAlloc(m, v28, int32(48))
							mBase = m.M
							v31 = m.ExcPending
							if v31 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v30
								v33 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v30)+40)) = v33
								v35 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v30)+32)) = v35
								v37 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v30)+24)) = v37
								v39 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v30)+16)) = v39
								v41 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v30)+8)) = v41
								v43 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								*(*int64)(unsafe.Add(mBase, uint32(v30))) = v43
								F_LockBuffer(m, v7, int32(0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_ReleaseBuffer(m, v7)
									mBase = m.M
									v49 = m.ExcPending
									if v49 != 0 {
										return int32(0)
									} else {
										v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
										v54 = v50
										v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)+20))
										return v55
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
func F__bt_insertonpg(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32, l9 int32, l10 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v209 float64
	_ = v209
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v329 int32
	_ = v329
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v348 int32
	_ = v348
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v438 int32
	_ = v438
	var v444 int32
	_ = v444
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v461 int32
	_ = v461
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v581 int32
	_ = v581
	var v583 int32
	_ = v583
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v617 int32
	_ = v617
	var v622 int32
	_ = v622
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v630 int32
	_ = v630
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v719 int32
	_ = v719
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v743 int32
	_ = v743
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v790 int32
	_ = v790
	var v793 int32
	_ = v793
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v886 int32
	_ = v886
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v974 int32
	_ = v974
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v996 int32
	_ = v996
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1031 float64
	_ = v1031
	var v1043 int32
	_ = v1043
	var v1055 int32
	_ = v1055
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1096 int32
	_ = v1096
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1104 int32
	_ = v1104
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1116 int32
	_ = v1116
	var v1121 int32
	_ = v1121
	var v1144 int32
	_ = v1144
	var v1191 int32
	_ = v1191
	var v1220 float64
	_ = v1220
	var v1223 int32
	_ = v1223
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1233 int32
	_ = v1233
	var v1234 int32
	_ = v1234
	var v1253 int32
	_ = v1253
	var v1288 int32
	_ = v1288
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1296 int32
	_ = v1296
	var v1299 float64
	_ = v1299
	var v1303 int32
	_ = v1303
	var v1305 int32
	_ = v1305
	var v1308 int32
	_ = v1308
	var v1311 int32
	_ = v1311
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1317 int32
	_ = v1317
	var v1319 int32
	_ = v1319
	var v1320 int32
	_ = v1320
	var v1338 int32
	_ = v1338
	var v1339 int32
	_ = v1339
	var v1345 int32
	_ = v1345
	var v1370 int32
	_ = v1370
	var v1373 int32
	_ = v1373
	var v1378 float64
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1381 float64
	_ = v1381
	var v1385 int32
	_ = v1385
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1391 int32
	_ = v1391
	var v1409 int32
	_ = v1409
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1449 int32
	_ = v1449
	var v1454 int32
	_ = v1454
	var v1502 int32
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1510 int32
	_ = v1510
	var v1511 int32
	_ = v1511
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1529 int32
	_ = v1529
	var v1530 int32
	_ = v1530
	var v1533 int32
	_ = v1533
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1572 int32
	_ = v1572
	var v1581 int32
	_ = v1581
	var v1583 int32
	_ = v1583
	var v1585 int32
	_ = v1585
	var v1586 int32
	_ = v1586
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1597 int32
	_ = v1597
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1606 int32
	_ = v1606
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1620 int32
	_ = v1620
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1631 int32
	_ = v1631
	var v1632 int32
	_ = v1632
	var v1636 int32
	_ = v1636
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1643 int32
	_ = v1643
	var v1649 int32
	_ = v1649
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1663 int32
	_ = v1663
	var v1668 int32
	_ = v1668
	var v1674 int32
	_ = v1674
	var v1675 int32
	_ = v1675
	var v1679 int32
	_ = v1679
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1701 int32
	_ = v1701
	var v1717 int32
	_ = v1717
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1760 int32
	_ = v1760
	var v1764 float64
	_ = v1764
	var v1768 int32
	_ = v1768
	var v1770 int32
	_ = v1770
	var v1773 int32
	_ = v1773
	var v1775 int32
	_ = v1775
	var v1778 int32
	_ = v1778
	var v1779 int32
	_ = v1779
	var v1781 int32
	_ = v1781
	var v1799 int32
	_ = v1799
	var v1800 int32
	_ = v1800
	var v1831 int32
	_ = v1831
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1874 int32
	_ = v1874
	var v1886 int32
	_ = v1886
	var v1890 int32
	_ = v1890
	var v1892 int32
	_ = v1892
	var v1893 int32
	_ = v1893
	var v1895 int32
	_ = v1895
	var v1896 int32
	_ = v1896
	var v1897 int32
	_ = v1897
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1913 int32
	_ = v1913
	var v1916 int32
	_ = v1916
	var v1917 int32
	_ = v1917
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1964 int32
	_ = v1964
	var v1966 int32
	_ = v1966
	var v1976 int32
	_ = v1976
	var v1988 int32
	_ = v1988
	var v1991 int32
	_ = v1991
	var v1993 int32
	_ = v1993
	var v1999 int32
	_ = v1999
	var v2003 int32
	_ = v2003
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2006 int32
	_ = v2006
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2027 int32
	_ = v2027
	var v2065 int32
	_ = v2065
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2071 int32
	_ = v2071
	var v2075 int32
	_ = v2075
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2081 int32
	_ = v2081
	var v2083 int32
	_ = v2083
	var v2085 int32
	_ = v2085
	var v2106 int32
	_ = v2106
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2138 int32
	_ = v2138
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2152 int32
	_ = v2152
	var v2154 int64
	_ = v2154
	var v2161 int32
	_ = v2161
	var v2163 int32
	_ = v2163
	var v2164 int32
	_ = v2164
	var v2172 int32
	_ = v2172
	var v2179 int32
	_ = v2179
	var v2182 int32
	_ = v2182
	var v2184 int32
	_ = v2184
	var v2190 int32
	_ = v2190
	var v2191 int32
	_ = v2191
	var v2195 int32
	_ = v2195
	var v2202 int32
	_ = v2202
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2206 int32
	_ = v2206
	var v2207 int32
	_ = v2207
	var v2210 int32
	_ = v2210
	var v2211 int32
	_ = v2211
	var v2214 int32
	_ = v2214
	var v2215 int32
	_ = v2215
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2223 int32
	_ = v2223
	var v2229 int32
	_ = v2229
	var v2231 int32
	_ = v2231
	var v2237 int32
	_ = v2237
	var v2241 int32
	_ = v2241
	var v2247 int32
	_ = v2247
	var v2249 int32
	_ = v2249
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2257 int32
	_ = v2257
	var v2261 int32
	_ = v2261
	var v2265 int32
	_ = v2265
	var v2266 int32
	_ = v2266
	var v2267 int32
	_ = v2267
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2275 int32
	_ = v2275
	var v2299 int32
	_ = v2299
	var v2324 int32
	_ = v2324
	var v2325 int32
	_ = v2325
	var v2327 int32
	_ = v2327
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2332 int32
	_ = v2332
	var v2381 int32
	_ = v2381
	var v2383 int32
	_ = v2383
	var v2387 int32
	_ = v2387
	var v2389 int32
	_ = v2389
	var v2390 int32
	_ = v2390
	var v2393 int32
	_ = v2393
	var v2395 int32
	_ = v2395
	var v2397 int32
	_ = v2397
	var v2399 int32
	_ = v2399
	var v2403 int32
	_ = v2403
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2415 int32
	_ = v2415
	var v2418 int32
	_ = v2418
	var v2421 int32
	_ = v2421
	var v2422 int32
	_ = v2422
	var v2430 int32
	_ = v2430
	var v2432 int32
	_ = v2432
	var v2439 int32
	_ = v2439
	var v2453 int32
	_ = v2453
	var v2458 int32
	_ = v2458
	var v2483 int32
	_ = v2483
	var v2484 int32
	_ = v2484
	var v2490 int32
	_ = v2490
	var v2496 int32
	_ = v2496
	var v2500 int32
	_ = v2500
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2513 int32
	_ = v2513
	var v2515 int64
	_ = v2515
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2531 int32
	_ = v2531
	var v2533 int32
	_ = v2533
	var v2534 int32
	_ = v2534
	var v2537 int32
	_ = v2537
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2549 int32
	_ = v2549
	var v2551 int64
	_ = v2551
	var v2558 int32
	_ = v2558
	var v2559 int32
	_ = v2559
	var v2561 int32
	_ = v2561
	var v2562 int32
	_ = v2562
	var v2567 int32
	_ = v2567
	var v2570 int32
	_ = v2570
	var v2572 int32
	_ = v2572
	var v2578 int32
	_ = v2578
	var v2592 int32
	_ = v2592
	var v2622 int32
	_ = v2622
	var v2628 int32
	_ = v2628
	var v2630 int64
	_ = v2630
	var v2637 int32
	_ = v2637
	var v2638 int32
	_ = v2638
	var v2640 int32
	_ = v2640
	var v2641 int32
	_ = v2641
	var v2649 int32
	_ = v2649
	var v2652 int32
	_ = v2652
	var v2654 int32
	_ = v2654
	var v2655 int32
	_ = v2655
	var v2659 int32
	_ = v2659
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2675 int32
	_ = v2675
	var v2676 int32
	_ = v2676
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2683 int32
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2695 int32
	_ = v2695
	var v2697 int32
	_ = v2697
	var v2699 int32
	_ = v2699
	var v2702 int32
	_ = v2702
	var v2708 int32
	_ = v2708
	var v2714 int32
	_ = v2714
	var v2716 int32
	_ = v2716
	var v2722 int32
	_ = v2722
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2725 int32
	_ = v2725
	var v2727 int32
	_ = v2727
	var v2730 int32
	_ = v2730
	var v2732 int32
	_ = v2732
	var v2733 int32
	_ = v2733
	var v2737 int32
	_ = v2737
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2747 int32
	_ = v2747
	var v2752 int32
	_ = v2752
	var v2756 int32
	_ = v2756
	var v2759 int32
	_ = v2759
	var v2764 int32
	_ = v2764
	var v2768 int32
	_ = v2768
	var v2772 int32
	_ = v2772
	var v2776 int32
	_ = v2776
	var v2782 int32
	_ = v2782
	var v2783 int32
	_ = v2783
	var v2784 int32
	_ = v2784
	var v2789 int32
	_ = v2789
	var v2790 int32
	_ = v2790
	var v2792 int32
	_ = v2792
	var v2795 int32
	_ = v2795
	var v2799 int32
	_ = v2799
	var v2801 int32
	_ = v2801
	var v2809 int32
	_ = v2809
	var v2811 int32
	_ = v2811
	var v2813 int32
	_ = v2813
	var v2816 int32
	_ = v2816
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2822 int64
	_ = v2822
	var v2823 int32
	_ = v2823
	var v2824 int64
	_ = v2824
	var v2827 int32
	_ = v2827
	var v2831 int32
	_ = v2831
	var v2838 int32
	_ = v2838
	var v2844 int32
	_ = v2844
	var v2846 int32
	_ = v2846
	var v2852 int32
	_ = v2852
	var v2855 int32
	_ = v2855
	var v2861 int32
	_ = v2861
	var v2863 int32
	_ = v2863
	var v2868 int32
	_ = v2868
	var v2869 int32
	_ = v2869
	var v2874 int32
	_ = v2874
	var v2876 int32
	_ = v2876
	var v2880 int32
	_ = v2880
	var v2886 int32
	_ = v2886
	var v2888 int32
	_ = v2888
	var v2894 int32
	_ = v2894
	var v2895 int32
	_ = v2895
	var v2899 int32
	_ = v2899
	var v2905 int32
	_ = v2905
	var v2907 int32
	_ = v2907
	var v2913 int32
	_ = v2913
	var v2914 int32
	_ = v2914
	var v2916 int32
	_ = v2916
	var v2920 int32
	_ = v2920
	var v2921 int32
	_ = v2921
	var v2927 int32
	_ = v2927
	var v2928 int32
	_ = v2928
	var v2932 int32
	_ = v2932
	var v2938 int32
	_ = v2938
	var v2940 int32
	_ = v2940
	var v2946 int32
	_ = v2946
	var v2948 int32
	_ = v2948
	var v2949 int32
	_ = v2949
	var v2950 int32
	_ = v2950
	var v2953 int32
	_ = v2953
	var v2955 int32
	_ = v2955
	var v2956 int32
	_ = v2956
	var v2957 int32
	_ = v2957
	var v2958 int32
	_ = v2958
	var v2960 int32
	_ = v2960
	var v2964 int32
	_ = v2964
	var v2970 int32
	_ = v2970
	var v2971 int32
	_ = v2971
	var v2976 int32
	_ = v2976
	var v2977 int32
	_ = v2977
	var v2981 int32
	_ = v2981
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2987 int32
	_ = v2987
	var v2991 int32
	_ = v2991
	var v3002 int32
	_ = v3002
	var v3008 int32
	_ = v3008
	var v3010 int32
	_ = v3010
	var v3016 int32
	_ = v3016
	var v3017 int32
	_ = v3017
	var v3019 int32
	_ = v3019
	var v3022 int32
	_ = v3022
	var v3028 int32
	_ = v3028
	var v3034 int32
	_ = v3034
	var v3036 int32
	_ = v3036
	var v3042 int32
	_ = v3042
	var v3043 int32
	_ = v3043
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3047 int32
	_ = v3047
	var v3050 int32
	_ = v3050
	var v3052 int32
	_ = v3052
	var v3053 int32
	_ = v3053
	var v3057 int32
	_ = v3057
	var v3060 int32
	_ = v3060
	var v3061 int32
	_ = v3061
	var v3064 int32
	_ = v3064
	var v3069 int32
	_ = v3069
	var v3076 int32
	_ = v3076
	var v3077 int32
	_ = v3077
	var v3079 int32
	_ = v3079
	var v3081 int32
	_ = v3081
	var v3083 int32
	_ = v3083
	var v3085 int32
	_ = v3085
	var v3087 int32
	_ = v3087
	var v3089 int32
	_ = v3089
	var v3094 int32
	_ = v3094
	var v3100 int32
	_ = v3100
	var v3103 int32
	_ = v3103
	var v3107 int32
	_ = v3107
	var v3111 int32
	_ = v3111
	var v3115 int32
	_ = v3115
	var v3118 int32
	_ = v3118
	var v3119 int32
	_ = v3119
	var v3126 int32
	_ = v3126
	var v3127 int32
	_ = v3127
	var v3128 int32
	_ = v3128
	var v3130 int32
	_ = v3130
	var v3134 int32
	_ = v3134
	var v3136 int64
	_ = v3136
	var v3137 int32
	_ = v3137
	var v3147 int32
	_ = v3147
	var v3153 int32
	_ = v3153
	var v3155 int32
	_ = v3155
	var v3161 int32
	_ = v3161
	var v3162 int64
	_ = v3162
	var v3168 int32
	_ = v3168
	var v3174 int32
	_ = v3174
	var v3176 int32
	_ = v3176
	var v3181 int32
	_ = v3181
	var v3185 int32
	_ = v3185
	var v3187 int32
	_ = v3187
	var v3190 int32
	_ = v3190
	var v3194 int32
	_ = v3194
	var v3200 int32
	_ = v3200
	var v3202 int32
	_ = v3202
	var v3208 int32
	_ = v3208
	var v3209 int32
	_ = v3209
	var v3211 int32
	_ = v3211
	var v3214 int32
	_ = v3214
	var v3215 int32
	_ = v3215
	var v3218 int32
	_ = v3218
	var v3219 int32
	_ = v3219
	var v3220 int32
	_ = v3220
	var v3222 int64
	_ = v3222
	var v3226 int32
	_ = v3226
	var v3227 int32
	_ = v3227
	var v3230 int32
	_ = v3230
	var v3231 int32
	_ = v3231
	var v3232 int32
	_ = v3232
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3238 int32
	_ = v3238
	var v3242 int32
	_ = v3242
	var v3244 int32
	_ = v3244
	var v3293 int32
	_ = v3293
	var v3295 int32
	_ = v3295
	var v3302 int32
	_ = v3302
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3307 int32
	_ = v3307
	var v3308 int32
	_ = v3308
	var v3312 int32
	_ = v3312
	var v3318 int32
	_ = v3318
	var v3320 int32
	_ = v3320
	var v3326 int32
	_ = v3326
	var v3327 int32
	_ = v3327
	var v3328 int32
	_ = v3328
	var v3343 int32
	_ = v3343
	var v3348 int32
	_ = v3348
	var v3352 int32
	_ = v3352
	var v3353 int32
	_ = v3353
	var v3360 int32
	_ = v3360
	var v3365 int32
	_ = v3365
	var v3369 int32
	_ = v3369
	var v3373 int32
	_ = v3373
	var v3374 int32
	_ = v3374
	var v3383 int32
	_ = v3383
	var v3388 int32
	_ = v3388
	var v3392 int32
	_ = v3392
	var v3396 int32
	_ = v3396
	var v3397 int32
	_ = v3397
	var v3406 int32
	_ = v3406
	var v3411 int32
	_ = v3411
	var v3415 int32
	_ = v3415
	var v3419 int32
	_ = v3419
	var v3420 int32
	_ = v3420
	var v3429 int32
	_ = v3429
	var v3434 int32
	_ = v3434
	var v3438 int32
	_ = v3438
	var v3442 int32
	_ = v3442
	var v3443 int32
	_ = v3443
	var v3452 int32
	_ = v3452
	var v3457 int32
	_ = v3457
	var v3461 int32
	_ = v3461
	var v3465 int32
	_ = v3465
	var v3466 int32
	_ = v3466
	var v3475 int32
	_ = v3475
	var v3480 int32
	_ = v3480
	var v3484 int32
	_ = v3484
	var v3488 int32
	_ = v3488
	var v3489 int32
	_ = v3489
	var v3498 int32
	_ = v3498
	var v3503 int32
	_ = v3503
	var v3507 int32
	_ = v3507
	var v3511 int32
	_ = v3511
	var v3514 int32
	_ = v3514
	var v3515 int32
	_ = v3515
	var v3516 int32
	_ = v3516
	var v3517 int32
	_ = v3517
	var v3528 int32
	_ = v3528
	var v3533 int32
	_ = v3533
	var v3537 int32
	_ = v3537
	var v3541 int32
	_ = v3541
	var v3547 int32
	_ = v3547
	var v3549 int32
	_ = v3549
	var v3555 int32
	_ = v3555
	var v3556 int32
	_ = v3556
	var v3557 int32
	_ = v3557
	var v3566 int32
	_ = v3566
	var v3571 int32
	_ = v3571
	v10 = l9
	v12 = int32(0)
	v47 = m.G0
	v49 = v47 - int32(224)
	m.G0 = v49
	if l3 < v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+16)))
	v70 = v69 + v68
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v70)+4))
	v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70)+12)))
	if v10 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v54 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v54+(l3^int32(-1))<<(uint(int32(2))%32))))
	v68 = v60
	goto L1
L3:
	;
	goto L4
L4:
	;
	v62 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v68 = v62 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3537 = m.ExcPending
	if v3537 != 0 {
		goto L21
	} else {
		goto L705
	}
L6:
	;
	v3507 = F__emscripten_memset_bulkmem(m, v2237, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L700
L7:
	;
	v3484 = F__emscripten_memset_bulkmem(m, v2237, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L696
L8:
	;
	v3461 = F__emscripten_memset_bulkmem(m, v2237, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L692
L9:
	;
	v3438 = F__emscripten_memset_bulkmem(m, v2237, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L688
L10:
	;
	v3415 = F__emscripten_memset_bulkmem(m, v2237, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L684
L11:
	;
	v3392 = F__emscripten_memset_bulkmem(m, v2237, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L680
L12:
	;
	v3369 = F__emscripten_memset_bulkmem(m, v2237, base.I32_extend8_s(int32(0)), int32(8192))
	mBase = m.M
	goto L676
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3352 = m.ExcPending
	if v3352 != 0 {
		goto L21
	} else {
		goto L673
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3302 = m.ExcPending
	if v3302 != 0 {
		goto L21
	} else {
		goto L665
	}
L15:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(l8<<(uint(int32(2))%32)+v68)+20))
	v80 = v68 + v77&int32(32767)
	v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+7)))
	if v81&int32(32) == int32(0) {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v101 = l6
	v102 = l8
	v103 = v12
	v105 = v12
	v106 = v12
	goto L17
L17:
	;
	v108 = v73 & int32(2)
	v109 = int32(4)
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+14)))
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v68)+12)))
	v112 = v110 - v111
	if v112 <= v109 {
		goto L26
	} else {
		goto L27
	}
L18:
	;
	v86 = int32(98304)
	if v77&v86 == v86 {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80)+4)))
	if v90&int32(8192) == int32(0) {
		goto L14
	} else {
		goto L20
	}
L20:
	;
	v97 = F_CopyIndexTuple(m, l6)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return
L22:
	;
	v99 = F__bt_swap_posting(m, v97, v80, v10)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v101 = v97
	v102 = l8 + int32(1)
	v103 = v80
	v105 = v99
	v106 = l6
	goto L17
L24:
	;
	if v10 != 0 {
		goto L660
	} else {
		goto L661
	}
L25:
	;
	if base.Ui32(v115-int32(4)) < base.Ui32(l7) {
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v115 = v109
	goto L28
L27:
	;
	v115 = v112
	goto L28
L28:
	;
	goto L25
L29:
	;
	if l3 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L30:
	;
	goto L31
L31:
	;
	v2921 = int32(0)
	if l10 == v2921 {
		v2955 = v2921
		v2956 = v2921
		v2957 = v12
		goto L549
	} else {
		goto L550
	}
L32:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+16)))
	v138 = v137 + v136
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+12)))
	if l3 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L33:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122+(l3^int32(-1))<<(uint(int32(2))%32))))
	v136 = v128
	goto L32
L34:
	;
	goto L35
L35:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v136 = v130 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L32
L36:
	;
	v162 = v102 & int32(65535)
	v164 = v49 + int32(222)
	v165 = m.G0
	v167 = v165 + int32(-64)
	m.G0 = v167
	v169 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+12)))
	v172 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+19)))
	v174 = v172 << (uint(int32(8)) % 32)
	v176 = v174 - int32(40)
	if base.Ui32(v169) < base.Ui32(int32(25)) {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v145+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v160 = v151
	goto L36
L38:
	;
	goto L39
L39:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v160 = v159
	goto L36
L40:
	;
	v182 = int32(0)
	goto L42
L41:
	;
	v182 = int32(base.Ui32(v169+int32(262120)) >> (uint(int32(2)) % 32))
	goto L42
L42:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+16)))
	v184 = v136 + v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v185 != 0 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v196 = v174 - (int32(base.Ui32(v186)>>(uint(int32(17))%32))+int32(7))&int32(65528) - int32(44)
	goto L45
L44:
	;
	v196 = v176
	goto L45
L45:
	;
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+14)))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+12)))
	v199 = v197 - v198
	v200 = int32(0)
	if v200 < v199 {
		goto L47
	} else {
		goto L48
	}
L46:
	;
	v204 = v196 - v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v205 != 0 {
		goto L50
	} else {
		goto L51
	}
L47:
	;
	v203 = v199
	goto L49
L48:
	;
	v203 = v200
	goto L49
L49:
	;
	goto L46
L50:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	v209 = base.F64_convert_i32_s(v206)
	goto L52
L51:
	;
	v209 = float64(90)
	goto L52
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+20)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v167)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v167)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v167)+24)) = l7 + int32(4)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+12)))
	v218 = v216 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+28)) = uint8(v218)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+44)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+40)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v167)+36)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v167)+32)) = v176
	v227 = v182 & int32(65535)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+48)) = v227
	*(*uint16)(unsafe.Add(mBase, uint32(v167)+30)) = uint16(v162)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+29)) = uint8(base.B2i32(v220 == int32(0)))
	v235 = F_palloc(m, v227*int32(10))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L21
	} else {
		goto L53
	}
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+56)) = v235
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v242 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v243 = int32(2)
	goto L56
L55:
	;
	v243 = int32(1)
	goto L56
L56:
	;
	if base.Ui32(v243) <= base.Ui32(v227) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v247 = int32(0)
	v259 = v247
	v263 = v247
	v267 = v243
	goto L60
L58:
	;
	goto L59
L59:
	;
	if base.Ui32(v162) <= base.Ui32(v227) {
		goto L145
	} else {
		goto L146
	}
L60:
	;
	v296 = v267 & int32(65535)
	v300 = (v296 - int32(1)) << (uint(int32(2)) % 32)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v136+int32(24)+v300)))
	v308 = (int32(base.Ui32(v302)>>(uint(int32(17))%32)) + int32(7)) & int32(65528)
	v310 = v308 | int32(4)
	if base.Ui32(v296) < base.Ui32(v162) {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L59
L62:
	;
	v790 = v310 + v263
	v793 = v267 + int32(1)
	if base.Ui32(v793&int32(65535)) <= base.Ui32(v227) {
		v259 = int32(0) - v790
		v263 = v790
		v267 = v793
		goto L60
	} else {
		goto L143
	}
L63:
	;
	v313 = v165 + int32(-52)
	v314 = int32(0)
	v318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v313)+18)))
	if v318 == v296 {
		goto L70
	} else {
		goto L71
	}
L64:
	;
	goto L65
L65:
	;
	if base.Ui32(v162) < base.Ui32(v296) {
		goto L92
	} else {
		goto L93
	}
L66:
	;
	goto L62
L67:
	;
	v389 = v383 + v386 - (v263 + v382)
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v313)+12))
	if v384 != 0 {
		goto L81
	} else {
		goto L82
	}
L68:
	;
	v382 = v376
	v383 = v377
	v384 = int32(1)
	v385 = v379
	v386 = v380
	goto L67
L69:
	;
	v376 = v371
	v377 = int32(-8)
	v379 = v373
	v380 = v374
	goto L68
L70:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v313)+24))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v313)+28))
	v323 = v320 + (v263 - v321)
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v313)+20))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v313)+12))
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+16)))
	if v326&int32(1) != 0 {
		v371 = v325
		v373 = v323
		v374 = v324
		goto L69
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	v329 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v313)+16)))
	if v329 != int32(1) {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v382 = v325
	v383 = v314
	v384 = v314
	v385 = v323
	v386 = v324
	goto L67
L74:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v313)+24))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v313)+28))
	v368 = v365 + (v263 - v366)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v313)+20))
	if v329 != 0 {
		v371 = v310
		v373 = v368
		v374 = v369
		goto L69
	} else {
		goto L80
	}
L75:
	;
	if base.Ui32(v310) < base.Ui32(int32(65)) {
		goto L74
	} else {
		goto L76
	}
L76:
	;
	v334 = int32(-8)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v313)+4))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335+v296<<(uint(int32(2))%32))+20))
	v342 = v335 + v339&int32(32767)
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v342)+6)))
	if v343&int32(8192) == int32(0) {
		v359 = v334
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v313)+24))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v313)+28))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v313)+20))
	v376 = v310
	v377 = v359
	v379 = v360 + (v263 - v361)
	v380 = v364
	goto L68
L78:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+5)))
	if v348&int32(32) == int32(0) {
		v359 = v334
		goto L77
	} else {
		goto L79
	}
L79:
	;
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v342)+2)))
	v359 = v343&int32(8191) - v355 - int32(8)
	goto L77
L80:
	;
	v382 = v310
	v383 = int32(0)
	v384 = v314
	v385 = v368
	v386 = v369
	goto L67
L81:
	;
	v395 = int32(0)
	goto L83
L82:
	;
	v395 = v382 + int32(65524)
	goto L83
L83:
	;
	v396 = v385 - v390 + v395
	if (v389|v396)&int32(32768) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v402 = *(*int32)(unsafe.Add(mBase, uint32(v313)+32))
	if base.Ui32(v402) < base.Ui32(v382) {
		goto L87
	} else {
		goto L88
	}
L85:
	;
	goto L86
L86:
	;
	goto L66
L87:
	;
	v404 = v402
	goto L89
L88:
	;
	v404 = v382
	goto L89
L89:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v313)+32)) = v404
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v313)+44))
	v407 = *(*int32)(unsafe.Add(mBase, uint32(v313)+40))
	v408 = int32(10)
	v411 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v406+v407*v408))) = uint16(v411)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v313)+44))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v313)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v413+v414*v408)+2)) = uint16(v389)
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v313)+44))
	v420 = *(*int32)(unsafe.Add(mBase, uint32(v313)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v419+v420*v408)+4)) = uint16(v396)
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v313)+44))
	v426 = *(*int32)(unsafe.Add(mBase, uint32(v313)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v425+v426*v408)+6)) = uint16(v296)
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v313)+44))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v313)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v431+v432*v408)+8)) = uint8(v411)
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v313)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v313)+40)) = v438 + int32(1)
	goto L86
L90:
	;
	v739 = v738 + v735
	if (v739|v734)&int32(32768) != 0 {
		goto L62
	} else {
		goto L139
	}
L91:
	;
	v734 = v259 - v310 + v728 + v727 - v725
	v735 = v726
	v738 = int32(0)
	goto L90
L92:
	;
	v444 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+28)))
	if v444 != int32(1) {
		goto L95
	} else {
		goto L96
	}
L93:
	;
	goto L94
L94:
	;
	v492 = v165 + int32(-52)
	v493 = int32(0)
	v497 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v492)+18)))
	if v497 == v296 {
		goto L106
	} else {
		goto L107
	}
L95:
	;
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v167)+36))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v167)+40))
	v482 = v479 + (v263 - v480)
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v167)+32))
	if v444 != 0 {
		v725 = v484
		v726 = v482
		v727 = int32(-8)
		v728 = v485
		goto L91
	} else {
		goto L101
	}
L96:
	;
	if base.Ui32(v310) < base.Ui32(int32(65)) {
		goto L95
	} else {
		goto L97
	}
L97:
	;
	v449 = int32(-8)
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v450+v300)+24))
	v455 = v450 + v452&int32(32767)
	v456 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v455)+6)))
	if v456&int32(8192) == int32(0) {
		v472 = v449
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v473 = *(*int32)(unsafe.Add(mBase, uint32(v167)+36))
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v167)+40))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v167)+32))
	v725 = v477
	v726 = v473 + v263 - v475
	v727 = v472
	v728 = v478
	goto L91
L99:
	;
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v455)+5)))
	if v461&int32(32) == int32(0) {
		v472 = v449
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v455)+2)))
	v472 = v456&int32(8191) - v468 - int32(8)
	goto L98
L101:
	;
	v734 = v259 + v485 - (v310 + v484)
	v735 = v482
	v738 = v308 + int32(65528)
	goto L90
L102:
	;
	v622 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+28)))
	if v622 != int32(1) {
		goto L128
	} else {
		goto L129
	}
L103:
	;
	v568 = v562 + v565 - (v263 + v561)
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	if v563 != 0 {
		goto L117
	} else {
		goto L118
	}
L104:
	;
	v561 = v555
	v562 = v556
	v563 = int32(1)
	v564 = v558
	v565 = v559
	goto L103
L105:
	;
	v555 = v550
	v556 = int32(-8)
	v558 = v552
	v559 = v553
	goto L104
L106:
	;
	v499 = *(*int32)(unsafe.Add(mBase, uint32(v492)+24))
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v492)+28))
	v502 = v499 + (v263 - v500)
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v492)+20))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v492)+12))
	v505 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+16)))
	if v505&int32(1) != 0 {
		v550 = v504
		v552 = v502
		v553 = v503
		goto L105
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v508 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v492)+16)))
	if v508 != int32(1) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	v561 = v504
	v562 = v493
	v563 = v493
	v564 = v502
	v565 = v503
	goto L103
L110:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v492)+24))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v492)+28))
	v547 = v544 + (v263 - v545)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v492)+20))
	if v508 != 0 {
		v550 = v310
		v552 = v547
		v553 = v548
		goto L105
	} else {
		goto L116
	}
L111:
	;
	if base.Ui32(v310) < base.Ui32(int32(65)) {
		goto L110
	} else {
		goto L112
	}
L112:
	;
	v513 = int32(-8)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v492)+4))
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v514+v296<<(uint(int32(2))%32))+20))
	v521 = v514 + v518&int32(32767)
	v522 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v521)+6)))
	if v522&int32(8192) == int32(0) {
		v538 = v513
		goto L113
	} else {
		goto L114
	}
L113:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v492)+24))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v492)+28))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v492)+20))
	v555 = v310
	v556 = v538
	v558 = v539 + (v263 - v540)
	v559 = v543
	goto L104
L114:
	;
	v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v521)+5)))
	if v527&int32(32) == int32(0) {
		v538 = v513
		goto L113
	} else {
		goto L115
	}
L115:
	;
	v534 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v521)+2)))
	v538 = v522&int32(8191) - v534 - int32(8)
	goto L113
L116:
	;
	v561 = v310
	v562 = int32(0)
	v563 = v493
	v564 = v547
	v565 = v548
	goto L103
L117:
	;
	v574 = int32(0)
	goto L119
L118:
	;
	v574 = v561 + int32(65524)
	goto L119
L119:
	;
	v575 = v564 - v569 + v574
	if (v568|v575)&int32(32768) == int32(0) {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v492)+32))
	if base.Ui32(v581) < base.Ui32(v561) {
		goto L123
	} else {
		goto L124
	}
L121:
	;
	goto L122
L122:
	;
	goto L102
L123:
	;
	v583 = v581
	goto L125
L124:
	;
	v583 = v561
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v492)+32)) = v583
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v492)+44))
	v586 = *(*int32)(unsafe.Add(mBase, uint32(v492)+40))
	v587 = int32(10)
	v590 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v585+v586*v587))) = uint16(v590)
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v492)+44))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v492)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v592+v593*v587)+2)) = uint16(v568)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(v492)+44))
	v599 = *(*int32)(unsafe.Add(mBase, uint32(v492)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v598+v599*v587)+4)) = uint16(v575)
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v492)+44))
	v605 = *(*int32)(unsafe.Add(mBase, uint32(v492)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v604+v605*v587)+6)) = uint16(v296)
	v610 = *(*int32)(unsafe.Add(mBase, uint32(v492)+44))
	v611 = *(*int32)(unsafe.Add(mBase, uint32(v492)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v610+v611*v587)+8)) = uint8(v590)
	v617 = *(*int32)(unsafe.Add(mBase, uint32(v492)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v492)+40)) = v617 + int32(1)
	goto L122
L126:
	;
	v681 = v680 + v678
	if (v681|v677)&int32(32768) != 0 {
		goto L62
	} else {
		goto L135
	}
L127:
	;
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v677 = v259 - v310 + v668 + v670 - v674
	v678 = v669
	v680 = int32(0)
	goto L126
L128:
	;
	v656 = *(*int32)(unsafe.Add(mBase, uint32(v167)+36))
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v167)+40))
	v659 = v656 + (v263 - v657)
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v167)+32))
	if v622 != 0 {
		v668 = v661
		v669 = v659
		v670 = int32(-8)
		goto L127
	} else {
		goto L134
	}
L129:
	;
	if base.Ui32(v310) < base.Ui32(int32(65)) {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	v627 = int32(-8)
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v628+v300)+24))
	v633 = v628 + v630&int32(32767)
	v634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v633)+6)))
	if v634&int32(8192) == int32(0) {
		v650 = v627
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v167)+36))
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v167)+40))
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v167)+32))
	v668 = v655
	v669 = v651 + v263 - v653
	v670 = v650
	goto L127
L132:
	;
	v639 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v633)+5)))
	if v639&int32(32) == int32(0) {
		v650 = v627
		goto L131
	} else {
		goto L133
	}
L133:
	;
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v633)+2)))
	v650 = v634&int32(8191) - v646 - int32(8)
	goto L131
L134:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v677 = v259 + v661 - (v310 + v663)
	v678 = v659
	v680 = v308 + int32(65528)
	goto L126
L135:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v167)+44))
	if base.Ui32(v685) < base.Ui32(v310) {
		goto L136
	} else {
		goto L137
	}
L136:
	;
	v687 = v685
	goto L138
L137:
	;
	v687 = v310
	goto L138
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+44)) = v687
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	v691 = int32(10)
	v694 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v689+v690*v691))) = uint16(v694)
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v696+v697*v691)+2)) = uint16(v677)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v702+v703*v691)+4)) = uint16(v681)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v708+v709*v691)+6)) = uint16(v267)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	v719 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v714+v715*v691)+8)) = uint8(v719)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+52)) = v721 + v719
	goto L62
L139:
	;
	v743 = *(*int32)(unsafe.Add(mBase, uint32(v167)+44))
	if base.Ui32(v743) < base.Ui32(v310) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v745 = v743
	goto L142
L141:
	;
	v745 = v310
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+44)) = v745
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v748 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	v749 = int32(10)
	v752 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v747+v748*v749))) = uint16(v752)
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v754+v755*v749)+2)) = uint16(v734)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v760+v761*v749)+4)) = uint16(v739)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v766+v767*v749)+6)) = uint16(v267)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	v777 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v772+v773*v749)+8)) = uint8(v777)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v167)+52)) = v779 + v777
	goto L62
L143:
	;
	goto L61
L144:
	;
	if v937 != 0 {
		goto L164
	} else {
		goto L165
	}
L145:
	;
	v844 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	v937 = v844
	goto L144
L146:
	;
	goto L147
L147:
	;
	v845 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	v846 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+30)))
	if v846 == v162 {
		goto L151
	} else {
		goto L152
	}
L148:
	;
	v891 = v890 + v887
	if (v891|v889)&int32(32768) != 0 {
		v937 = v845
		goto L144
	} else {
		goto L156
	}
L149:
	;
	v886 = v876
	v887 = v879 - v877
	v889 = v878 - (v876 + v204) - int32(8)
	v890 = int32(0)
	goto L148
L150:
	;
	v886 = v867
	v887 = v870 - v868
	v889 = v869 - (v867 + v204)
	v890 = v867 + int32(65524)
	goto L148
L151:
	;
	v848 = *(*int32)(unsafe.Add(mBase, uint32(v167)+36))
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v167)+40))
	v851 = v848 + (v204 - v849)
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v167)+32))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v854 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+28)))
	if v854&int32(1) == int32(0) {
		v867 = v853
		v868 = v853
		v869 = v852
		v870 = v851
		goto L150
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	v859 = *(*int32)(unsafe.Add(mBase, uint32(v167)+36))
	v860 = *(*int32)(unsafe.Add(mBase, uint32(v167)+40))
	v862 = v859 + (v204 - v860)
	v863 = int32(0)
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v167)+32))
	v866 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+28)))
	if v866 != 0 {
		v876 = v863
		v877 = v864
		v878 = v865
		v879 = v862
		goto L149
	} else {
		goto L155
	}
L154:
	;
	v876 = v853
	v877 = v853
	v878 = v852
	v879 = v851
	goto L149
L155:
	;
	v867 = v863
	v868 = v864
	v869 = v865
	v870 = v862
	goto L150
L156:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v167)+44))
	if base.Ui32(v895) < base.Ui32(v886) {
		goto L157
	} else {
		goto L158
	}
L157:
	;
	v897 = v895
	goto L159
L158:
	;
	v897 = v886
	goto L159
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+44)) = v897
	v899 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v900 = int32(10)
	v903 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v899+v845*v900))) = uint16(v903)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v906 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v905+v906*v900)+2)) = uint16(v889)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v911+v912*v900)+4)) = uint16(v891)
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v917+v918*v900)+6)) = uint16(v162)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v923+v924*v900)+8)) = uint8(v903)
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	v932 = v930 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v167)+52)) = v932
	v937 = v932
	goto L144
L160:
	;
	m.G0 = v167 - int32(-64)
	v2136 = v2106 & int32(65535)
	v2137 = F_PageGetTempPage(m, v136)
	mBase = m.M
	v2138 = m.ExcPending
	if v2138 != 0 {
		goto L21
	} else {
		goto L351
	}
L161:
	;
	v1223 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v1226 = v1223 + v937*int32(10)
	v1229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1226-int32(2)))))
	v1232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1226-int32(4)))))
	v1233 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1223)+8)))
	v1234 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1223)+6)))
	if int32(0) < v937 {
		goto L208
	} else {
		goto L209
	}
L162:
	;
	v1191 = v1144
	v1220 = float64(0.5)
	goto L161
L163:
	;
	v1144 = int32(0)
	goto L162
L164:
	;
	v939 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+29)))
	v940 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+28)))
	if v940 != int32(1) {
		goto L167
	} else {
		goto L168
	}
L165:
	;
	goto L166
L166:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1109 = m.ExcPending
	if v1109 != 0 {
		goto L21
	} else {
		goto L205
	}
L167:
	;
	v1191 = v939
	v1220 = float64(0.7)
	goto L161
L168:
	;
	goto L169
L169:
	;
	if v939&int32(1) != 0 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v1191 = int32(1)
	v1220 = base.F64_div(v209, float64(100))
	goto L161
L171:
	;
	goto L172
L172:
	;
	v949 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v950 = *(*int32)(unsafe.Add(mBase, uint32(v949)+192))
	v951 = int32(*(*int16)(unsafe.Add(mBase, uint32(v950)+10)))
	if v951 == int32(1) {
		goto L163
	} else {
		goto L173
	}
L173:
	;
	v954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+30)))
	if v954 == int32(2) {
		goto L163
	} else {
		goto L174
	}
L174:
	;
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v167)+44))
	if v957 != v958 {
		goto L163
	} else {
		goto L175
	}
L175:
	;
	if base.Ui32(int32(28)) < base.Ui32(v957) {
		goto L163
	} else {
		goto L176
	}
L176:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v167)+40))
	v964 = v227 - int32(1)
	if v962 != v957*v964 {
		goto L163
	} else {
		goto L177
	}
L177:
	;
	v967 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	if base.Ui32(v182&int32(65535)) < base.Ui32(v954) {
		goto L178
	} else {
		goto L179
	}
L178:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v967+v964<<(uint(int32(2))%32))+24))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v979 = F__bt_keep_natts_fast(m, v949, v967+v974&int32(32767), v978)
	mBase = m.M
	v980 = m.ExcPending
	if v980 != 0 {
		goto L21
	} else {
		goto L181
	}
L179:
	;
	goto L180
L180:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32((v954-int32(1))&int32(65535)<<(uint(int32(2))%32)+v967)+20))
	v999 = v967 + v996&int32(32767)
	v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v999)+7)))
	if v1000&int32(32) != 0 {
		goto L184
	} else {
		goto L185
	}
L181:
	;
	if v979 < int32(2) {
		goto L163
	} else {
		goto L182
	}
L182:
	;
	if v951 < v979 {
		v1191 = int32(0)
		v1220 = float64(0.5)
		goto L161
	} else {
		goto L183
	}
L183:
	;
	v1191 = int32(1)
	v1220 = base.F64_div(v209, float64(100))
	goto L161
L184:
	;
	v1003 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v999)+5)))
	if v1003&int32(32) != 0 {
		goto L163
	} else {
		goto L187
	}
L185:
	;
	goto L186
L186:
	;
	v1006 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v999)+2)))
	v1007 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v999))))
	v1008 = int32(16)
	v1010 = v1006 | v1007<<(uint(v1008)%32)
	v1011 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v1012 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1011))))
	v1015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1011)+2)))
	v1016 = v1012<<(uint(v1008)%32) | v1015
	if v1010 != v1016 {
		goto L188
	} else {
		goto L189
	}
L187:
	;
	goto L186
L188:
	;
	if v1010+int32(1) != v1016 {
		goto L163
	} else {
		goto L191
	}
L189:
	;
	goto L190
L190:
	;
	v1024 = F__bt_keep_natts_fast(m, v949, v999, v1011)
	mBase = m.M
	v1025 = m.ExcPending
	if v1025 != 0 {
		goto L21
	} else {
		goto L193
	}
L191:
	;
	v1021 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1011)+4)))
	if v1021 != int32(1) {
		goto L163
	} else {
		goto L192
	}
L192:
	;
	goto L190
L193:
	;
	if v1024 < int32(2) {
		goto L163
	} else {
		goto L194
	}
L194:
	;
	if v951 < v1024 {
		goto L163
	} else {
		goto L195
	}
L195:
	;
	v1029 = int32(1)
	v1031 = base.F64_div(v209, float64(100))
	if base.F64_lt(v1031, base.F64_div(base.F64_convert_i32_u(v954), base.F64_convert_i32_u((v182+v1029)&int32(65535)))) != 0 {
		v1191 = v1029
		v1220 = v1031
		goto L161
	} else {
		goto L196
	}
L196:
	;
	if v937 <= int32(0) {
		goto L163
	} else {
		goto L197
	}
L197:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v1055 = int32(0)
	goto L198
L198:
	;
	v1092 = v1043 + v1055*int32(10)
	v1093 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1092)+8)))
	if v1093 != int32(1) {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	v1144 = int32(0)
	goto L162
L200:
	;
	v1104 = v1055 + int32(1)
	if v1104 != v937 {
		v1055 = v1104
		goto L198
	} else {
		goto L204
	}
L201:
	;
	v1096 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1092)+6)))
	if v162 != v1096 {
		goto L200
	} else {
		goto L202
	}
L202:
	;
	F_pfree(m, v1043)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L21
	} else {
		goto L203
	}
L203:
	;
	v1100 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v1100)
	v2106 = v162
	goto L160
L204:
	;
	goto L199
L205:
	;
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v167))) = v1110 + int32(4)
	F_errmsg_internal(m, int32(693137), v167)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L21
	} else {
		goto L206
	}
L206:
	;
	F_errfinish(m, int32(500056), int32(261), int32(488965))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L21
	} else {
		goto L207
	}
L207:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L208:
	;
	v1253 = int32(0)
	goto L211
L209:
	;
	v1338 = v937
	v1339 = v1223
	v1345 = v940
	goto L210
L210:
	;
	F_pg_qsort(m, v1339, v1338, int32(10), int32(239))
	mBase = m.M
	v1370 = m.ExcPending
	if v1370 != 0 {
		goto L21
	} else {
		goto L221
	}
L211:
	;
	v1288 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v1291 = v1288 + v1253*int32(10)
	v1292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1291)+2)))
	if v1191&int32(1) != 0 {
		goto L214
	} else {
		goto L215
	}
L212:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+28)))
	v1338 = v1317
	v1339 = v1319
	v1345 = v1320
	goto L210
L213:
	;
	v1311 = base.I32_extend16_s(v1308) >> (uint(int32(15)) % 32)
	v1313 = v1311 ^ v1308 - v1311
	*(*uint16)(unsafe.Add(mBase, uint32(v1291))) = uint16(v1313)
	v1316 = v1253 + int32(1)
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	if v1316 < v1317 {
		v1253 = v1316
		goto L211
	} else {
		goto L220
	}
L214:
	;
	v1296 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1291)+4)))
	v1299 = base.F64_sub(base.F64_mul(v1220, base.F64_convert_i32_s(base.I32_extend16_s(v1292))), base.F64_mul(base.F64_sub(float64(1), v1220), base.F64_convert_i32_s(v1296)))
	if base.F64_lt(base.F64_abs(v1299), float64(2.147483648e+09)) != 0 {
		goto L217
	} else {
		goto L218
	}
L215:
	;
	goto L216
L216:
	;
	v1305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1291)+4)))
	v1308 = v1292 - v1305
	goto L213
L217:
	;
	v1303 = base.I32_trunc_f64_s(v1299)
	v1308 = v1303
	goto L213
L218:
	;
	goto L219
L219:
	;
	v1308 = int32(-2147483648)
	goto L213
L220:
	;
	goto L212
L221:
	;
	v1373 = int32(1)
	if v1345&v1373 != 0 {
		goto L223
	} else {
		goto L224
	}
L222:
	;
	if v1338 < int32(2) {
		v1502 = v1338
		goto L229
	} else {
		goto L230
	}
L223:
	;
	v1378 = float64(0.05)
	goto L225
L224:
	;
	v1378 = float64(0.075)
	goto L225
L225:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v167)+40))
	v1381 = base.F64_mul(v1378, base.F64_convert_i32_s(v1379))
	if base.F64_lt(base.F64_abs(v1381), float64(2.147483648e+09)) != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	v1385 = base.I32_trunc_f64_s(v1381)
	v1387 = v1385
	goto L222
L227:
	;
	goto L228
L228:
	;
	v1387 = int32(-2147483648)
	goto L222
L229:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+60)) = v1502
	if v1345&int32(1) == int32(0) {
		goto L239
	} else {
		goto L240
	}
L230:
	;
	v1388 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1339)+4)))
	v1391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1339)+2)))
	v1409 = v1373
	goto L231
L231:
	;
	v1446 = v1339 + v1409*int32(10)
	v1447 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1446)+2)))
	if v1447 < base.I32_extend16_s(v1391-v1387) {
		v1502 = v1409
		goto L229
	} else {
		goto L233
	}
L232:
	;
	v1502 = v1338
	goto L229
L233:
	;
	v1449 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1446)+4)))
	if v1449 < base.I32_extend16_s(v1388-v1387) {
		v1502 = v1409
		goto L229
	} else {
		goto L234
	}
L234:
	;
	if base.I32_extend16_s(v1391+v1387) < v1447 {
		v1502 = v1409
		goto L229
	} else {
		goto L235
	}
L235:
	;
	if base.I32_extend16_s(v1388+v1387) < v1449 {
		v1502 = v1409
		goto L229
	} else {
		goto L236
	}
L236:
	;
	v1454 = v1409 + int32(1)
	if v1454 != v1338 {
		v1409 = v1454
		goto L231
	} else {
		goto L237
	}
L237:
	;
	goto L232
L238:
	;
	if v1865 < v1856 {
		goto L307
	} else {
		goto L308
	}
L239:
	;
	v1508 = *(*int32)(unsafe.Add(mBase, uint32(v167)+44))
	v1856 = v1338
	v1857 = v1339
	v1865 = v1502
	v1866 = v1508
	v1874 = v12
	goto L238
L240:
	;
	goto L241
L241:
	;
	v1509 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v1510 = *(*int32)(unsafe.Add(mBase, uint32(v1509)+192))
	v1511 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1510)+10)))
	if v1502 < v1338 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	v1513 = v1502
	goto L244
L243:
	;
	v1513 = v1338
	goto L244
L244:
	;
	v1516 = int32(0)
	v1517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1339)+6)))
	v1529 = v1516
	v1530 = v1513 - int32(1)
	v1533 = v1516
	goto L245
L245:
	;
	v1567 = v1339 + v1530*int32(10)
	v1568 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1567)+6)))
	if base.Ui32(v1568) < base.Ui32(v1517) {
		goto L251
	} else {
		goto L252
	}
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+60)) = v1338
	v1856 = v1338
	v1857 = v1339
	v1865 = v1338
	v1866 = v1511
	v1874 = int32(1)
	goto L238
L247:
	;
	goto L246
L248:
	;
	v1529 = int32(0)
	v1530 = v1530 - int32(1)
	v1533 = v1585
	goto L245
L249:
	;
	v1593 = int32(0)
	v1597 = v1530 - int32(1)
	if base.B2i32(v1592 == v1593)&base.B2i32(v1593 <= v1597) != 0 {
		v1529 = v1589
		v1530 = v1597
		v1533 = v1593
		goto L245
	} else {
		goto L269
	}
L250:
	;
	if v1586 == int32(0) {
		goto L248
	} else {
		goto L268
	}
L251:
	;
	if v1529 != 0 {
		v1589 = v1529
		v1592 = v1533
		goto L249
	} else {
		goto L267
	}
L252:
	;
	if base.Ui32(v1517) < base.Ui32(v1568) {
		goto L254
	} else {
		goto L255
	}
L253:
	;
	v1585 = v1567
	v1586 = v1583
	goto L250
L254:
	;
	if v1533 != 0 {
		v1585 = v1533
		v1586 = v1529
		goto L250
	} else {
		goto L266
	}
L255:
	;
	v1571 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1339)+8)))
	v1572 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1567)+8)))
	if v1572 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	if v1529 != 0 {
		goto L262
	} else {
		goto L263
	}
L257:
	;
	if v1571&int32(1) != 0 {
		goto L251
	} else {
		goto L260
	}
L258:
	;
	goto L259
L259:
	;
	if v1571&int32(1) == int32(0) {
		goto L254
	} else {
		goto L261
	}
L260:
	;
	goto L256
L261:
	;
	goto L256
L262:
	;
	v1581 = v1529
	goto L264
L263:
	;
	v1581 = v1567
	goto L264
L264:
	;
	if v1533 != 0 {
		v1585 = v1533
		v1586 = v1581
		goto L250
	} else {
		goto L265
	}
L265:
	;
	v1583 = v1581
	goto L253
L266:
	;
	v1583 = v1529
	goto L253
L267:
	;
	v1585 = v1533
	v1586 = v1567
	goto L250
L268:
	;
	v1589 = v1586
	v1592 = v1585
	goto L249
L269:
	;
	v1601 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1589)+6)))
	v1602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+30)))
	v1603 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589)+8)))
	if v1603 != int32(1) {
		goto L271
	} else {
		goto L272
	}
L270:
	;
	v1626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1592)+6)))
	v1627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1592)+8)))
	if v1627 != 0 {
		goto L275
	} else {
		goto L276
	}
L271:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v1620 = *(*int32)(unsafe.Add(mBase, uint32(v1612+(v1601-int32(1))&int32(65535)<<(uint(int32(2))%32))+20))
	v1625 = v1612 + v1620&int32(32767)
	goto L270
L272:
	;
	v1606 = int32(65535)
	if v1601&v1606 != v1602&v1606 {
		goto L271
	} else {
		goto L273
	}
L273:
	;
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v1625 = v1611
	goto L270
L274:
	;
	v1642 = F__bt_keep_natts_fast(m, v1509, v1625, v1641)
	mBase = m.M
	v1643 = m.ExcPending
	if v1643 != 0 {
		goto L21
	} else {
		goto L278
	}
L275:
	;
	v1632 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v1636 = *(*int32)(unsafe.Add(mBase, uint32(v1632+v1626<<(uint(int32(2))%32))+20))
	v1641 = v1632 + v1636&int32(32767)
	goto L274
L276:
	;
	if v1626 != v1602&int32(65535) {
		goto L275
	} else {
		goto L277
	}
L277:
	;
	v1631 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v1641 = v1631
	goto L274
L278:
	;
	if v1642 <= v1511 {
		v1856 = v1338
		v1857 = v1339
		v1865 = v1502
		v1866 = v1642
		v1874 = v12
		goto L238
	} else {
		goto L279
	}
L279:
	;
	if v1233&int32(1) == int32(0) {
		goto L281
	} else {
		goto L282
	}
L280:
	;
	if v1229&int32(1) != 0 {
		goto L285
	} else {
		goto L286
	}
L281:
	;
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v1663 = *(*int32)(unsafe.Add(mBase, uint32(v1655+(v1234-int32(1))&int32(65535)<<(uint(int32(2))%32))+20))
	v1668 = v1655 + v1663&int32(32767)
	goto L280
L282:
	;
	v1649 = int32(65535)
	if v1234&v1649 != v1602&v1649 {
		goto L281
	} else {
		goto L283
	}
L283:
	;
	v1654 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v1668 = v1654
	goto L280
L284:
	;
	v1685 = F__bt_keep_natts_fast(m, v1509, v1668, v1684)
	mBase = m.M
	v1686 = m.ExcPending
	if v1686 != 0 {
		goto L21
	} else {
		goto L288
	}
L285:
	;
	v1675 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1675+v1232<<(uint(int32(2))%32))+20))
	v1684 = v1675 + v1679&int32(32767)
	goto L284
L286:
	;
	if v1232 != v1602&int32(65535) {
		goto L285
	} else {
		goto L287
	}
L287:
	;
	v1674 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v1684 = v1674
	goto L284
L288:
	;
	if v1685 <= v1511 {
		goto L247
	} else {
		goto L289
	}
L289:
	;
	v1688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+29)))
	if v1688 == int32(0) {
		goto L290
	} else {
		goto L291
	}
L290:
	;
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v1691)+24))
	v1696 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v1697 = F__bt_keep_natts_fast(m, v1509, v1691+v1692&int32(32767), v1696)
	mBase = m.M
	v1698 = m.ExcPending
	if v1698 != 0 {
		goto L21
	} else {
		goto L293
	}
L291:
	;
	v1701 = v1685
	goto L292
L292:
	;
	if int32(0) < v1338 {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	if v1511 < v1697 {
		v1856 = v1338
		v1857 = v1339
		v1865 = v1502
		v1866 = v1697
		v1874 = v12
		goto L238
	} else {
		goto L294
	}
L294:
	;
	v1701 = v1697
	goto L292
L295:
	;
	v1717 = int32(0)
	goto L298
L296:
	;
	v1799 = v1338
	v1800 = v1339
	goto L297
L297:
	;
	F_pg_qsort(m, v1800, v1799, int32(10), int32(239))
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L21
	} else {
		goto L305
	}
L298:
	;
	v1752 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v1755 = v1752 + v1717*int32(10)
	v1756 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1755)+2)))
	v1760 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1755)+4)))
	v1764 = base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1756), float64(0.96)), base.F64_mul(base.F64_convert_i32_s(v1760), float64(-0.040000000000000036)))
	if base.F64_lt(base.F64_abs(v1764), float64(2.147483648e+09)) != 0 {
		goto L301
	} else {
		goto L302
	}
L299:
	;
	v1781 = *(*int32)(unsafe.Add(mBase, uint32(v167)+56))
	v1799 = v1779
	v1800 = v1781
	goto L297
L300:
	;
	v1773 = base.I32_extend16_s(v1770) >> (uint(int32(15)) % 32)
	v1775 = v1770 ^ v1773 - v1773
	*(*uint16)(unsafe.Add(mBase, uint32(v1755))) = uint16(v1775)
	v1778 = v1717 + int32(1)
	v1779 = *(*int32)(unsafe.Add(mBase, uint32(v167)+52))
	if v1778 < v1779 {
		v1717 = v1778
		goto L298
	} else {
		goto L304
	}
L301:
	;
	v1768 = base.I32_trunc_f64_s(v1764)
	v1770 = v1768
	goto L300
L302:
	;
	goto L303
L303:
	;
	v1770 = int32(-2147483648)
	goto L300
L304:
	;
	goto L299
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v167)+60)) = int32(1)
	v1856 = v1799
	v1857 = v1800
	v1865 = int32(1)
	v1866 = v1701
	v1874 = v12
	goto L238
L306:
	;
	v2065 = v1857 + v2027*int32(10)
	if v1874 == int32(0) {
		v2078 = v2065
		goto L340
	} else {
		goto L341
	}
L307:
	;
	v1886 = v1865
	goto L309
L308:
	;
	v1886 = v1856
	goto L309
L309:
	;
	if v1886 <= int32(0) {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v2027 = int32(0)
	goto L306
L311:
	;
	goto L312
L312:
	;
	v1890 = *(*int32)(unsafe.Add(mBase, uint32(v167)+16))
	v1892 = v1890 + int32(24)
	v1893 = int32(0)
	v1895 = *(*int32)(unsafe.Add(mBase, uint32(v167)+12))
	v1896 = *(*int32)(unsafe.Add(mBase, uint32(v167)+20))
	v1897 = *(*int32)(unsafe.Add(mBase, uint32(v167)+24))
	v1898 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+28)))
	v1901 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+30)))
	v1913 = v1893
	v1916 = int32(2147483647)
	v1917 = v1893
	goto L313
L313:
	;
	v1951 = v1857 + v1917*int32(10)
	v1952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1951)+6)))
	v1953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1951)+8)))
	if v1898&int32(1) == int32(0) {
		goto L316
	} else {
		goto L317
	}
L314:
	;
	v2027 = v2011
	goto L306
L315:
	;
	v2010 = base.B2i32(v2009 < v1916)
	if v2009 < v1916 {
		goto L332
	} else {
		goto L333
	}
L316:
	;
	if v1953&int32(1) == int32(0) {
		goto L319
	} else {
		goto L320
	}
L317:
	;
	goto L318
L318:
	;
	v1976 = v1953 & int32(1)
	if v1976 != 0 {
		goto L325
	} else {
		goto L326
	}
L319:
	;
	if v1952 == v1901 {
		v2009 = v1897
		goto L315
	} else {
		goto L322
	}
L320:
	;
	goto L321
L321:
	;
	v1964 = int32(4)
	v1966 = *(*int32)(unsafe.Add(mBase, uint32(v1952<<(uint(int32(2))%32)+v1892-v1964)))
	v2009 = (int32(base.Ui32(v1966)>>(uint(int32(17))%32))+int32(7))&int32(65528) | v1964
	goto L315
L322:
	;
	goto L321
L323:
	;
	v2005 = F__bt_keep_natts_fast(m, v1895, v2003, v2004)
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L21
	} else {
		goto L331
	}
L324:
	;
	v1999 = *(*int32)(unsafe.Add(mBase, uint32(v1952<<(uint(int32(2))%32)+v1892-int32(4))))
	v2003 = v1993
	v2004 = v1890 + v1999&int32(32767)
	goto L323
L325:
	;
	if v1952 == v1901 {
		v1993 = v1896
		goto L324
	} else {
		goto L328
	}
L326:
	;
	goto L327
L327:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32((v1952-int32(1))&int32(65535)<<(uint(int32(2))%32)+v1892-int32(4))))
	v1991 = v1890 + v1988&int32(32767)
	if v1976 != 0 {
		v1993 = v1991
		goto L324
	} else {
		goto L329
	}
L328:
	;
	goto L327
L329:
	;
	if v1952 == v1901 {
		v2003 = v1991
		v2004 = v1896
		goto L323
	} else {
		goto L330
	}
L330:
	;
	v1993 = v1991
	goto L324
L331:
	;
	v2009 = v2005
	goto L315
L332:
	;
	v2011 = v1917
	goto L334
L333:
	;
	v2011 = v1913
	goto L334
L334:
	;
	if v2009 <= v1866 {
		v2027 = v2011
		goto L306
	} else {
		goto L335
	}
L335:
	;
	if v2009 < v1916 {
		goto L336
	} else {
		goto L337
	}
L336:
	;
	v2013 = v2009
	goto L338
L337:
	;
	v2013 = v1916
	goto L338
L338:
	;
	v2015 = v1917 + int32(1)
	if v2015 != v1886 {
		v1913 = v2011
		v1916 = v2013
		v1917 = v2015
		goto L313
	} else {
		goto L339
	}
L339:
	;
	goto L314
L340:
	;
	v2081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2078)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v164))) = uint8(v2081)
	v2083 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2078)+6)))
	F_pfree(m, v1857)
	mBase = m.M
	v2085 = m.ExcPending
	if v2085 != 0 {
		goto L21
	} else {
		goto L350
	}
L341:
	;
	v2068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+29)))
	if v2068 != 0 {
		v2078 = v2065
		goto L340
	} else {
		goto L342
	}
L342:
	;
	v2069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2065)+8)))
	if v2069 != 0 {
		v2078 = v2065
		goto L340
	} else {
		goto L343
	}
L343:
	;
	v2070 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2065)+6)))
	v2071 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v167)+30)))
	if base.Ui32(v2070) < base.Ui32(v2071+int32(9)) {
		goto L344
	} else {
		goto L345
	}
L344:
	;
	v2075 = v1857
	goto L346
L345:
	;
	v2075 = v2065
	goto L346
L346:
	;
	if base.Ui32(v2071) <= base.Ui32(v2070) {
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v2077 = v2075
	goto L349
L348:
	;
	v2077 = v2065
	goto L349
L349:
	;
	v2078 = v2077
	goto L340
L350:
	;
	v2106 = v2083
	goto L160
L351:
	;
	F_PageInit(m, v2137, int32(8192), int32(16))
	mBase = m.M
	goto L352
L352:
	;
	v2142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2137)+16)))
	v2143 = v2137 + v2142
	v2144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v2148 = v2144&int32(65309) | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v2143)+12)) = uint16(v2148)
	v2150 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	*(*int32)(unsafe.Add(mBase, uint32(v2143))) = v2150
	v2152 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2143)+8)) = v2152
	v2154 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	*(*int64)(unsafe.Add(mBase, uint32(v2137))) = v2154
	if v10&int32(65535) != 0 {
		goto L353
	} else {
		goto L354
	}
L353:
	;
	v2161 = v102 - int32(1)
	goto L355
L354:
	;
	v2161 = int32(0)
	goto L355
L355:
	;
	v2163 = v140 & int32(1)
	v2164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+222)))
	if v2164 == int32(0) {
		goto L357
	} else {
		goto L358
	}
L356:
	;
	if v2163 != 0 {
		goto L364
	} else {
		goto L365
	}
L357:
	;
	if v2136 == v162 {
		v2182 = v101
		v2184 = l7
		goto L356
	} else {
		goto L360
	}
L358:
	;
	goto L359
L359:
	;
	v2172 = *(*int32)(unsafe.Add(mBase, uint32(v2136<<(uint(int32(2))%32)+v136)+20))
	if v2161&int32(65535) == v2136 {
		goto L361
	} else {
		goto L362
	}
L360:
	;
	goto L359
L361:
	;
	v2179 = v105
	goto L363
L362:
	;
	v2179 = v136 + v2172&int32(32767)
	goto L363
L363:
	;
	v2182 = v2179
	v2184 = int32(base.Ui32(v2172) >> (uint(int32(17)) % 32))
	goto L356
L364:
	;
	if v2164 != 0 {
		goto L368
	} else {
		goto L369
	}
L365:
	;
	v2210 = v2182
	v2211 = v2184
	goto L366
L366:
	;
	v2214 = F_PageAddItemExtended(m, v2137, v2210, v2211, int32(1), int32(0))
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L21
	} else {
		goto L376
	}
L367:
	;
	v2205 = F__bt_truncate(m, l0, v2204, v2182, l2)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L21
	} else {
		goto L375
	}
L368:
	;
	if v102&int32(65535) == v2136 {
		v2204 = v101
		goto L367
	} else {
		goto L371
	}
L369:
	;
	goto L370
L370:
	;
	v2190 = int32(65535)
	v2191 = (v2136 - int32(1)) & v2190
	v2195 = *(*int32)(unsafe.Add(mBase, uint32(v2191<<(uint(int32(2))%32)+v136)+20))
	if v2161&v2190 == v2191 {
		goto L372
	} else {
		goto L373
	}
L371:
	;
	goto L370
L372:
	;
	v2202 = v105
	goto L374
L373:
	;
	v2202 = v136 + v2195&int32(32767)
	goto L374
L374:
	;
	v2204 = v2202
	goto L367
L375:
	;
	v2207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2205)+6)))
	v2210 = v2205
	v2211 = v2207 & int32(8191)
	goto L366
L376:
	;
	if v2214 == int32(0) {
		goto L13
	} else {
		goto L377
	}
L377:
	;
	v2218 = F__bt_allocbuf(m, l0, l1)
	mBase = m.M
	v2219 = m.ExcPending
	if v2219 != 0 {
		goto L21
	} else {
		goto L379
	}
L378:
	;
	if v2218 < int32(0) {
		goto L384
	} else {
		goto L385
	}
L379:
	;
	if v2218 < int32(0) {
		goto L380
	} else {
		goto L381
	}
L380:
	;
	v2223 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2229 = *(*int32)(unsafe.Add(mBase, uint32(v2223+(v2218^int32(-1))<<(uint(int32(2))%32))))
	v2237 = v2229
	goto L378
L381:
	;
	goto L382
L382:
	;
	v2231 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v2237 = v2231 + v2218<<(uint(int32(13))%32) + int32(-8192)
	goto L378
L383:
	;
	v2257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2237)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v2143)+4)) = v2256
	v2261 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	v2265 = F_LWLockAcquire(m, v2261+int32(2560), int32(1))
	mBase = m.M
	v2266 = m.ExcPending
	if v2266 != 0 {
		goto L21
	} else {
		goto L387
	}
L384:
	;
	v2241 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(v2241+(v2218^int32(-1))<<(uint(int32(6))%32))+16))
	v2256 = v2247
	goto L383
L385:
	;
	goto L386
L386:
	;
	v2249 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v2255 = *(*int32)(unsafe.Add(mBase, uint32(v2249+v2218<<(uint(int32(6))%32)+int32(-64))+16))
	v2256 = v2255
	goto L383
L387:
	;
	v2267 = int32(0)
	v2269 = *(*int32)(unsafe.Add(mBase, _consts[56]))
	v2270 = *(*int32)(unsafe.Add(mBase, uint32(v2269)+4))
	if v2270 <= v2267 {
		v2381 = v2267
		goto L388
	} else {
		goto L389
	}
L388:
	;
	v2383 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	F_LWLockRelease(m, v2383+int32(2560))
	mBase = m.M
	v2387 = m.ExcPending
	if v2387 != 0 {
		goto L21
	} else {
		goto L396
	}
L389:
	;
	v2275 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v2299 = int32(0)
	goto L390
L390:
	;
	v2324 = v2269 + int32(12) + v2299*int32(12)
	v2325 = *(*int32)(unsafe.Add(mBase, uint32(v2324)))
	if v2325 != v2275 {
		goto L392
	} else {
		goto L393
	}
L391:
	;
	v2381 = int32(0)
	goto L388
L392:
	;
	v2332 = v2299 + int32(1)
	if v2332 != v2270 {
		v2299 = v2332
		goto L390
	} else {
		goto L395
	}
L393:
	;
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2324)+4))
	v2328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v2327 != v2328 {
		goto L392
	} else {
		goto L394
	}
L394:
	;
	v2330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2324)+8)))
	v2381 = v2330
	goto L388
L395:
	;
	goto L391
L396:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2143)+14)) = uint16(v2381)
	v2389 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v2390 = v2257 + v2237
	*(*int32)(unsafe.Add(mBase, uint32(v2390))) = v160
	v2393 = v2389 & int32(65437)
	*(*uint16)(unsafe.Add(mBase, uint32(v2390)+12)) = uint16(v2393)
	v2395 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2390)+4)) = v2395
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2390)+8)) = v2397
	v2399 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2143)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2390)+14)) = uint16(v2399)
	if v139 != 0 {
		goto L397
	} else {
		goto L398
	}
L397:
	;
	v2403 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v2411 = F_PageAddItemExtended(m, v2237, v136+v2403&int32(32767), int32(base.Ui32(v2403)>>(uint(int32(17))%32)), int32(1), int32(0))
	mBase = m.M
	v2412 = m.ExcPending
	if v2412 != 0 {
		goto L21
	} else {
		goto L400
	}
L398:
	;
	v2415 = int32(1)
	goto L399
L399:
	;
	if v2163 != 0 {
		goto L402
	} else {
		goto L403
	}
L400:
	;
	if v2411 == int32(0) {
		goto L12
	} else {
		goto L401
	}
L401:
	;
	v2415 = int32(2)
	goto L399
L402:
	;
	v2418 = int32(0)
	goto L404
L403:
	;
	v2418 = v2415
	goto L404
L404:
	;
	v2421 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v2421 != 0 {
		goto L405
	} else {
		goto L406
	}
L405:
	;
	v2422 = int32(2)
	goto L407
L406:
	;
	v2422 = int32(1)
	goto L407
L407:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v141) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v2430 = int32(base.Ui32(v141+int32(262120)) >> (uint(int32(2)) % 32))
	goto L410
L409:
	;
	v2430 = int32(0)
	goto L410
L410:
	;
	v2432 = v2430 & int32(65535)
	if base.Ui32(v2422) <= base.Ui32(v2432) {
		goto L411
	} else {
		goto L412
	}
L411:
	;
	v2439 = v2415
	v2453 = v2422
	v2458 = int32(2)
	goto L414
L412:
	;
	v2578 = v2415
	v2592 = v2422
	goto L413
L413:
	;
	v2622 = int32(65535)
	if base.Ui32(v2592&v2622) <= base.Ui32(v102&v2622) {
		goto L443
	} else {
		goto L444
	}
L414:
	;
	v2483 = int32(65535)
	v2484 = v2453 & v2483
	v2490 = *(*int32)(unsafe.Add(mBase, uint32(v2484<<(uint(int32(2))%32)+(v136+int32(24))-int32(4))))
	if v2484 == v2161&v2483 {
		goto L417
	} else {
		goto L418
	}
L415:
	;
	v2578 = v2567
	v2592 = v2572
	goto L413
L416:
	;
	v2537 = int32(base.Ui32(v2490) >> (uint(int32(17)) % 32))
	if base.Ui32(v2484) < base.Ui32(v2136) {
		goto L432
	} else {
		goto L433
	}
L417:
	;
	v2531 = v2439
	v2533 = v2458
	v2534 = v105
	goto L416
L418:
	;
	goto L419
L419:
	;
	v2496 = v136 + v2490&int32(32767)
	if v2484 != v102&int32(65535) {
		v2531 = v2439
		v2533 = v2458
		v2534 = v2496
		goto L416
	} else {
		goto L420
	}
L420:
	;
	v2500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+222)))
	if v2500 == int32(1) {
		goto L421
	} else {
		goto L422
	}
L421:
	;
	v2506 = F_PageAddItemExtended(m, v2137, v101, l7, v2458&int32(65535), int32(0))
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L21
	} else {
		goto L424
	}
L422:
	;
	goto L423
L423:
	;
	v2513 = v2439 & int32(65535)
	if v2513 == v2418 {
		goto L426
	} else {
		goto L427
	}
L424:
	;
	if v2506 == int32(0) {
		goto L11
	} else {
		goto L425
	}
L425:
	;
	v2531 = v2439
	v2533 = v2458 + int32(1)
	v2534 = v2496
	goto L416
L426:
	;
	v2515 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
	*(*int64)(unsafe.Add(mBase, uint32(v49)+192)) = v2515
	*(*int32)(unsafe.Add(mBase, uint32(v49)+196)) = int32(537395200)
	v2522 = int32(8)
	v2523 = v49 + int32(192)
	goto L428
L427:
	;
	v2522 = l7
	v2523 = v101
	goto L428
L428:
	;
	v2525 = F_PageAddItemExtended(m, v2237, v2523, v2522, v2513, int32(0))
	mBase = m.M
	v2526 = m.ExcPending
	if v2526 != 0 {
		goto L21
	} else {
		goto L429
	}
L429:
	;
	if v2525 == int32(0) {
		goto L10
	} else {
		goto L430
	}
L430:
	;
	v2531 = v2439 + int32(1)
	v2533 = v2458
	v2534 = v2496
	goto L416
L431:
	;
	v2572 = v2453 + int32(1)
	if base.Ui32(v2572&int32(65535)) <= base.Ui32(v2432) {
		v2439 = v2567
		v2453 = v2572
		v2458 = v2570
		goto L414
	} else {
		goto L442
	}
L432:
	;
	v2542 = F_PageAddItemExtended(m, v2137, v2534, v2537, v2533&int32(65535), int32(0))
	mBase = m.M
	v2543 = m.ExcPending
	if v2543 != 0 {
		goto L21
	} else {
		goto L435
	}
L433:
	;
	goto L434
L434:
	;
	v2549 = v2531 & int32(65535)
	if v2549 == v2418 {
		goto L437
	} else {
		goto L438
	}
L435:
	;
	if v2542 == int32(0) {
		goto L9
	} else {
		goto L436
	}
L436:
	;
	v2567 = v2531
	v2570 = v2533 + int32(1)
	goto L431
L437:
	;
	v2551 = *(*int64)(unsafe.Add(mBase, uint32(v2534)))
	*(*int64)(unsafe.Add(mBase, uint32(v49)+192)) = v2551
	*(*int32)(unsafe.Add(mBase, uint32(v49)+196)) = int32(537395200)
	v2558 = int32(8)
	v2559 = v49 + int32(192)
	goto L439
L438:
	;
	v2558 = v2537
	v2559 = v2534
	goto L439
L439:
	;
	v2561 = F_PageAddItemExtended(m, v2237, v2559, v2558, v2549, int32(0))
	mBase = m.M
	v2562 = m.ExcPending
	if v2562 != 0 {
		goto L21
	} else {
		goto L440
	}
L440:
	;
	if v2561 == int32(0) {
		goto L8
	} else {
		goto L441
	}
L441:
	;
	v2567 = v2531 + int32(1)
	v2570 = v2533
	goto L431
L442:
	;
	goto L415
L443:
	;
	v2628 = v2578 & int32(65535)
	if v2628 == v2418 {
		goto L446
	} else {
		goto L447
	}
L444:
	;
	goto L445
L445:
	;
	if v139 == int32(0) {
		goto L452
	} else {
		goto L453
	}
L446:
	;
	v2630 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
	*(*int64)(unsafe.Add(mBase, uint32(v49)+192)) = v2630
	*(*int32)(unsafe.Add(mBase, uint32(v49)+196)) = int32(537395200)
	v2637 = int32(8)
	v2638 = v49 + int32(192)
	goto L448
L447:
	;
	v2637 = l7
	v2638 = v101
	goto L448
L448:
	;
	v2640 = F_PageAddItemExtended(m, v2237, v2638, v2637, v2628, int32(0))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L21
	} else {
		goto L449
	}
L449:
	;
	if v2640 == int32(0) {
		goto L7
	} else {
		goto L450
	}
L450:
	;
	goto L445
L451:
	;
	v2688 = int32(4510148)
	v2690 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v2690 + int32(1)
	F_PageRestoreTempPage(m, v2137, v136)
	mBase = m.M
	v2695 = m.ExcPending
	if v2695 != 0 {
		goto L21
	} else {
		goto L462
	}
L452:
	;
	v2649 = int32(0)
	v2685 = v2649
	v2686 = v2649
	v2687 = v2649
	goto L451
L453:
	;
	goto L454
L454:
	;
	v2652 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v2654 = F__bt_getbuf(m, l0, v2652, int32(2))
	mBase = m.M
	v2655 = m.ExcPending
	if v2655 != 0 {
		goto L21
	} else {
		goto L456
	}
L455:
	;
	v2674 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2673)+16)))
	v2675 = v2674 + v2673
	v2676 = *(*int32)(unsafe.Add(mBase, uint32(v2675)))
	if v2676 != v160 {
		goto L6
	} else {
		goto L460
	}
L456:
	;
	if v2654 < int32(0) {
		goto L457
	} else {
		goto L458
	}
L457:
	;
	v2659 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2659+(v2654^int32(-1))<<(uint(int32(2))%32))))
	v2673 = v2665
	goto L455
L458:
	;
	goto L459
L459:
	;
	v2667 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v2673 = v2667 + v2654<<(uint(int32(13))%32) + int32(-8192)
	goto L455
L460:
	;
	v2678 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2675)+14)))
	v2679 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2390)+14)))
	if v2678 == v2679 {
		v2685 = v2654
		v2686 = v2673
		v2687 = v2675
		goto L451
	} else {
		goto L461
	}
L461:
	;
	v2681 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2390)+12)))
	v2683 = v2681 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v2390)+12)) = uint16(v2683)
	v2685 = v2654
	v2686 = v2673
	v2687 = v2675
	goto L451
L462:
	;
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v2697 = m.ExcPending
	if v2697 != 0 {
		goto L21
	} else {
		goto L463
	}
L463:
	;
	F_MarkBufferDirty(m, v2218)
	mBase = m.M
	v2699 = m.ExcPending
	if v2699 != 0 {
		goto L21
	} else {
		goto L464
	}
L464:
	;
	if v139 != 0 {
		goto L465
	} else {
		goto L466
	}
L465:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2687))) = v2256
	F_MarkBufferDirty(m, v2685)
	mBase = m.M
	v2702 = m.ExcPending
	if v2702 != 0 {
		goto L21
	} else {
		goto L468
	}
L466:
	;
	goto L467
L467:
	;
	if v2163 == int32(0) {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	goto L467
L469:
	;
	if l4 < int32(0) {
		goto L473
	} else {
		goto L474
	}
L470:
	;
	goto L471
L471:
	;
	v2732 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2733 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2732)+118)))
	if v2733 != int32(112) {
		v2855 = v2210
		goto L477
	} else {
		goto L478
	}
L472:
	;
	v2723 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2722)+16)))
	v2724 = v2723 + v2722
	v2725 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2724)+12)))
	v2727 = v2725 & int32(65407)
	*(*uint16)(unsafe.Add(mBase, uint32(v2724)+12)) = uint16(v2727)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v2730 = m.ExcPending
	if v2730 != 0 {
		goto L21
	} else {
		goto L476
	}
L473:
	;
	v2708 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(v2708+(l4^int32(-1))<<(uint(int32(2))%32))))
	v2722 = v2714
	goto L472
L474:
	;
	goto L475
L475:
	;
	v2716 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v2722 = v2716 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L472
L476:
	;
	goto L471
L477:
	;
	v2861 = int32(4510148)
	v2863 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v2863 - int32(1)
	if v139 != 0 {
		goto L529
	} else {
		goto L530
	}
L478:
	;
	v2737 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v2737 <= int32(0) {
		goto L479
	} else {
		goto L480
	}
L479:
	;
	v2740 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2740 != 0 {
		v2855 = v2210
		goto L477
	} else {
		goto L482
	}
L480:
	;
	goto L481
L481:
	;
	v2742 = *(*int32)(unsafe.Add(mBase, uint32(v2390)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+198)) = uint16(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+192)) = v2742
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+196)) = uint16(v2136)
	v2747 = int32(65535)
	if base.Ui32(v2161&v2747) < base.Ui32(v2136&v2747) {
		goto L484
	} else {
		goto L485
	}
L482:
	;
	v2741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2741 != 0 {
		v2855 = v2210
		goto L477
	} else {
		goto L483
	}
L483:
	;
	goto L481
L484:
	;
	v2752 = v10
	goto L486
L485:
	;
	v2752 = int32(0)
	goto L486
L486:
	;
	if v10&int32(65535) != 0 {
		goto L487
	} else {
		goto L488
	}
L487:
	;
	v2756 = v2752
	goto L489
L488:
	;
	v2756 = int32(0)
	goto L489
L489:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+200)) = uint16(v2756)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2759 = m.ExcPending
	if v2759 != 0 {
		goto L21
	} else {
		goto L490
	}
L490:
	;
	F_XLogRegisterData(m, v49+int32(192), int32(10))
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		goto L21
	} else {
		goto L491
	}
L491:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v2768 = m.ExcPending
	if v2768 != 0 {
		goto L21
	} else {
		goto L492
	}
L492:
	;
	F_XLogRegisterBuffer(m, int32(1), v2218, int32(6))
	mBase = m.M
	v2772 = m.ExcPending
	if v2772 != 0 {
		goto L21
	} else {
		goto L493
	}
L493:
	;
	if v139 != 0 {
		goto L494
	} else {
		goto L495
	}
L494:
	;
	F_XLogRegisterBuffer(m, int32(2), v2685, int32(8))
	mBase = m.M
	v2776 = m.ExcPending
	if v2776 != 0 {
		goto L21
	} else {
		goto L497
	}
L495:
	;
	goto L496
L496:
	;
	if v2163 == int32(0) {
		goto L498
	} else {
		goto L499
	}
L497:
	;
	goto L496
L498:
	;
	F_XLogRegisterBuffer(m, int32(3), l4, int32(8))
	mBase = m.M
	v2782 = m.ExcPending
	if v2782 != 0 {
		goto L21
	} else {
		goto L501
	}
L499:
	;
	goto L500
L500:
	;
	v2783 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+222)))
	v2784 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v49)+200)))
	if v2783|v2784 != 0 {
		goto L502
	} else {
		goto L503
	}
L501:
	;
	goto L500
L502:
	;
	if v2784&int32(65535) != 0 {
		goto L505
	} else {
		goto L506
	}
L503:
	;
	goto L504
L504:
	;
	if v2163 == int32(0) {
		goto L512
	} else {
		goto L513
	}
L505:
	;
	v2789 = v106
	goto L507
L506:
	;
	v2789 = v101
	goto L507
L507:
	;
	if v2783 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v2790 = v2789
	goto L510
L509:
	;
	v2790 = v106
	goto L510
L510:
	;
	F_XLogRegisterBufData(m, int32(0), v2790, l7)
	mBase = m.M
	v2792 = m.ExcPending
	if v2792 != 0 {
		goto L21
	} else {
		goto L511
	}
L511:
	;
	goto L504
L512:
	;
	v2795 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v2799 = v136 + v2795&int32(32767)
	goto L514
L513:
	;
	v2799 = v2210
	goto L514
L514:
	;
	v2801 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2799)+6)))
	F_XLogRegisterBufData(m, int32(0), v2799, (v2801&int32(8191)+int32(7))&int32(16376))
	mBase = m.M
	v2809 = m.ExcPending
	if v2809 != 0 {
		goto L21
	} else {
		goto L515
	}
L515:
	;
	v2811 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2237)+14)))
	v2813 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2237)+16)))
	F_XLogRegisterBufData(m, int32(1), v2237+v2811, v2813-v2811)
	mBase = m.M
	v2816 = m.ExcPending
	if v2816 != 0 {
		goto L21
	} else {
		goto L516
	}
L516:
	;
	v2820 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+222)))
	if v2820 != 0 {
		goto L517
	} else {
		goto L518
	}
L517:
	;
	v2821 = int32(48)
	goto L519
L518:
	;
	v2821 = int32(64)
	goto L519
L519:
	;
	v2822 = F_XLogInsert(m, int32(11), v2821)
	mBase = m.M
	v2823 = m.ExcPending
	if v2823 != 0 {
		goto L21
	} else {
		goto L520
	}
L520:
	;
	v2824 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v136))) = base.I64_rotr(v2822, v2824)
	v2827 = base.I32_wrap_i64(v2822)
	*(*int32)(unsafe.Add(mBase, uint32(v2237)+4)) = v2827
	v2831 = base.I32_wrap_i64(int64(base.Ui64(v2822) >> (uint(v2824) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v2237))) = v2831
	if v139 != 0 {
		goto L521
	} else {
		goto L522
	}
L521:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2686)+4)) = v2827
	*(*int32)(unsafe.Add(mBase, uint32(v2686))) = v2831
	goto L523
L522:
	;
	goto L523
L523:
	;
	if v2163 != 0 {
		v2855 = v2799
		goto L477
	} else {
		goto L524
	}
L524:
	;
	if l4 < int32(0) {
		goto L526
	} else {
		goto L527
	}
L525:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2852)+4)) = v2827
	*(*int32)(unsafe.Add(mBase, uint32(v2852))) = v2831
	v2855 = v2799
	goto L477
L526:
	;
	v2838 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2844 = *(*int32)(unsafe.Add(mBase, uint32(v2838+(l4^int32(-1))<<(uint(int32(2))%32))))
	v2852 = v2844
	goto L525
L527:
	;
	goto L528
L528:
	;
	v2846 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v2852 = v2846 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L525
L529:
	;
	F__bt_relbuf(m, v2685)
	mBase = m.M
	v2868 = m.ExcPending
	if v2868 != 0 {
		goto L21
	} else {
		goto L532
	}
L530:
	;
	goto L531
L531:
	;
	v2869 = int32(0)
	if v2163 == v2869 {
		goto L534
	} else {
		goto L535
	}
L532:
	;
	goto L531
L533:
	;
	if l3 < int32(0) {
		goto L540
	} else {
		goto L541
	}
L534:
	;
	F__bt_relbuf(m, l4)
	mBase = m.M
	v2874 = m.ExcPending
	if v2874 != 0 {
		goto L21
	} else {
		goto L537
	}
L535:
	;
	goto L536
L536:
	;
	F_pfree(m, v2855)
	mBase = m.M
	v2876 = m.ExcPending
	if v2876 != 0 {
		goto L21
	} else {
		goto L538
	}
L537:
	;
	goto L533
L538:
	;
	goto L533
L539:
	;
	if v2218 < int32(0) {
		goto L544
	} else {
		goto L545
	}
L540:
	;
	v2880 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v2886 = *(*int32)(unsafe.Add(mBase, uint32(v2880+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v2895 = v2886
	goto L539
L541:
	;
	goto L542
L542:
	;
	v2888 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v2894 = *(*int32)(unsafe.Add(mBase, uint32(v2888+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v2895 = v2894
	goto L539
L543:
	;
	F_PredicateLockPageSplit(m, l0, v2895, v2914)
	mBase = m.M
	v2916 = m.ExcPending
	if v2916 != 0 {
		goto L21
	} else {
		goto L547
	}
L544:
	;
	v2899 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v2905 = *(*int32)(unsafe.Add(mBase, uint32(v2899+(v2218^int32(-1))<<(uint(int32(6))%32))+16))
	v2914 = v2905
	goto L543
L545:
	;
	goto L546
L546:
	;
	v2907 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(v2907+v2218<<(uint(int32(6))%32)+int32(-64))+16))
	v2914 = v2913
	goto L543
L547:
	;
	F__bt_insert_parent(m, l0, l1, l3, v2218, l5, base.B2i32(v108 != int32(0)), base.B2i32(v72|v71 == v2869))
	mBase = m.M
	v2920 = m.ExcPending
	if v2920 != 0 {
		goto L21
	} else {
		goto L548
	}
L548:
	;
	goto L24
L549:
	;
	v2958 = int32(4510148)
	v2960 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v2960 + int32(1)
	if v10 != 0 {
		goto L558
	} else {
		goto L559
	}
L550:
	;
	v2927 = F__bt_getbuf(m, l0, int32(0), int32(2))
	mBase = m.M
	v2928 = m.ExcPending
	if v2928 != 0 {
		goto L21
	} else {
		goto L552
	}
L551:
	;
	v2948 = v2946 + int32(24)
	v2949 = *(*int32)(unsafe.Add(mBase, uint32(v2946)+44))
	v2950 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	if base.Ui32(v2949) < base.Ui32(v2950) {
		v2955 = v2927
		v2956 = v2948
		v2957 = v2946
		goto L549
	} else {
		goto L556
	}
L552:
	;
	if v2927 < int32(0) {
		goto L553
	} else {
		goto L554
	}
L553:
	;
	v2932 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v2938 = *(*int32)(unsafe.Add(mBase, uint32(v2932+(v2927^int32(-1))<<(uint(int32(2))%32))))
	v2946 = v2938
	goto L551
L554:
	;
	goto L555
L555:
	;
	v2940 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v2946 = v2940 + v2927<<(uint(int32(13))%32) + int32(-8192)
	goto L551
L556:
	;
	F__bt_relbuf(m, v2927)
	mBase = m.M
	v2953 = m.ExcPending
	if v2953 != 0 {
		goto L21
	} else {
		goto L557
	}
L557:
	;
	v2955 = int32(0)
	v2956 = v2948
	v2957 = v2946
	goto L549
L558:
	;
	v2964 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
	v2970 = (v2964&int32(8191) + int32(7)) & int32(16376)
	if v2970 != 0 {
		goto L562
	} else {
		goto L563
	}
L559:
	;
	goto L560
L560:
	;
	v2976 = F_PageAddItemExtended(m, v68, v101, l7, v102&int32(65535), int32(0))
	mBase = m.M
	v2977 = m.ExcPending
	if v2977 != 0 {
		goto L21
	} else {
		goto L565
	}
L561:
	;
	goto L560
L562:
	;
	v2971 = F__emscripten_memcpy_bulkmem(m, v103, v105, v2970)
	mBase = m.M
	goto L564
L563:
	;
	goto L564
L564:
	;
	goto L561
L565:
	;
	if v2976 == int32(0) {
		goto L5
	} else {
		goto L566
	}
L566:
	;
	v2981 = v73 & int32(1)
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v2983 = m.ExcPending
	if v2983 != 0 {
		goto L21
	} else {
		goto L567
	}
L567:
	;
	if v2955 != 0 {
		goto L568
	} else {
		goto L569
	}
L568:
	;
	v2984 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+4))
	if base.Ui32(v2984) <= base.Ui32(int32(2)) {
		goto L571
	} else {
		goto L572
	}
L569:
	;
	goto L570
L570:
	;
	if v2981 == int32(0) {
		goto L580
	} else {
		goto L581
	}
L571:
	;
	v2987 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v2957)+12)) = uint16(v2987)
	v2991 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2957-int32(-64)))) = uint8(v2991)
	*(*int64)(unsafe.Add(mBase, uint32(v2957)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v2957)+48)) = v2991
	*(*int32)(unsafe.Add(mBase, uint32(v2957)+28)) = int32(3)
	goto L574
L572:
	;
	goto L573
L573:
	;
	if l3 < int32(0) {
		goto L576
	} else {
		goto L577
	}
L574:
	;
	goto L573
L575:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2956)+16)) = v3017
	v3019 = *(*int32)(unsafe.Add(mBase, uint32(v70)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2956)+20)) = v3019
	F_MarkBufferDirty(m, v2955)
	mBase = m.M
	v3022 = m.ExcPending
	if v3022 != 0 {
		goto L21
	} else {
		goto L579
	}
L576:
	;
	v3002 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v3008 = *(*int32)(unsafe.Add(mBase, uint32(v3002+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v3017 = v3008
	goto L575
L577:
	;
	goto L578
L578:
	;
	v3010 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v3016 = *(*int32)(unsafe.Add(mBase, uint32(v3010+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v3017 = v3016
	goto L575
L579:
	;
	goto L570
L580:
	;
	if l4 < int32(0) {
		goto L584
	} else {
		goto L585
	}
L581:
	;
	goto L582
L582:
	;
	v3052 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3052)+118)))
	if v3053 != int32(112) {
		goto L588
	} else {
		goto L589
	}
L583:
	;
	v3043 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3042)+16)))
	v3044 = v3043 + v3042
	v3045 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3044)+12)))
	v3047 = v3045 & int32(65407)
	*(*uint16)(unsafe.Add(mBase, uint32(v3044)+12)) = uint16(v3047)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v3050 = m.ExcPending
	if v3050 != 0 {
		goto L21
	} else {
		goto L587
	}
L584:
	;
	v3028 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v3034 = *(*int32)(unsafe.Add(mBase, uint32(v3028+(l4^int32(-1))<<(uint(int32(2))%32))))
	v3042 = v3034
	goto L583
L585:
	;
	goto L586
L586:
	;
	v3036 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v3042 = v3036 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L583
L587:
	;
	goto L582
L588:
	;
	v3174 = int32(4510148)
	v3176 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v3176 - int32(1)
	if v2955 != 0 {
		goto L631
	} else {
		goto L632
	}
L589:
	;
	v3057 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v3057 <= int32(0) {
		goto L590
	} else {
		goto L591
	}
L590:
	;
	v3060 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v3060 != 0 {
		goto L588
	} else {
		goto L593
	}
L591:
	;
	goto L592
L592:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+222)) = uint16(v102)
	F_XLogBeginInsert(m)
	mBase = m.M
	v3064 = m.ExcPending
	if v3064 != 0 {
		goto L21
	} else {
		goto L595
	}
L593:
	;
	v3061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v3061 != 0 {
		goto L588
	} else {
		goto L594
	}
L594:
	;
	goto L592
L595:
	;
	F_XLogRegisterData(m, v49+int32(222), int32(2))
	mBase = m.M
	v3069 = m.ExcPending
	if v3069 != 0 {
		goto L21
	} else {
		goto L596
	}
L596:
	;
	if v10|v2981 == int32(0) {
		goto L598
	} else {
		goto L599
	}
L597:
	;
	v3130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3128)+6)))
	F_XLogRegisterBufData(m, int32(0), v3128, v3130&int32(8191))
	mBase = m.M
	v3134 = m.ExcPending
	if v3134 != 0 {
		goto L21
	} else {
		goto L619
	}
L598:
	;
	F_XLogRegisterBuffer(m, int32(1), l4, int32(8))
	mBase = m.M
	v3076 = m.ExcPending
	if v3076 != 0 {
		goto L21
	} else {
		goto L601
	}
L599:
	;
	goto L600
L600:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v3111 = m.ExcPending
	if v3111 != 0 {
		goto L21
	} else {
		goto L608
	}
L601:
	;
	if v2955 != 0 {
		goto L602
	} else {
		goto L603
	}
L602:
	;
	v3077 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+192)) = v3077
	v3079 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+196)) = v3079
	v3081 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+200)) = v3081
	v3083 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+204)) = v3083
	v3085 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+208)) = v3085
	v3087 = *(*int32)(unsafe.Add(mBase, uint32(v2956)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+212)) = v3087
	v3089 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2956)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v49)+216)) = uint8(v3089)
	F_XLogRegisterBuffer(m, int32(2), v2955, int32(14))
	mBase = m.M
	v3094 = m.ExcPending
	if v3094 != 0 {
		goto L21
	} else {
		goto L605
	}
L603:
	;
	v3103 = int32(16)
	goto L604
L604:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v3107 = m.ExcPending
	if v3107 != 0 {
		goto L21
	} else {
		goto L607
	}
L605:
	;
	F_XLogRegisterBufData(m, int32(2), v49+int32(192), int32(28))
	mBase = m.M
	v3100 = m.ExcPending
	if v3100 != 0 {
		goto L21
	} else {
		goto L606
	}
L606:
	;
	v3103 = int32(32)
	goto L604
L607:
	;
	v3127 = v3103
	v3128 = v101
	goto L597
L608:
	;
	if v10 == int32(0) {
		goto L609
	} else {
		goto L610
	}
L609:
	;
	v3115 = int32(80)
	if v2981 != 0 {
		goto L612
	} else {
		goto L613
	}
L610:
	;
	goto L611
L611:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v49)+192)) = uint16(v10)
	F_XLogRegisterBufData(m, int32(0), v49+int32(192), int32(2))
	mBase = m.M
	v3126 = m.ExcPending
	if v3126 != 0 {
		goto L21
	} else {
		goto L618
	}
L612:
	;
	v3118 = int32(0)
	goto L614
L613:
	;
	v3118 = v3115
	goto L614
L614:
	;
	if v10 != 0 {
		goto L615
	} else {
		goto L616
	}
L615:
	;
	v3119 = v3115
	goto L617
L616:
	;
	v3119 = v3118
	goto L617
L617:
	;
	v3127 = v3119
	v3128 = v101
	goto L597
L618:
	;
	v3127 = int32(80)
	v3128 = v106
	goto L597
L619:
	;
	v3136 = F_XLogInsert(m, int32(11), v3127)
	mBase = m.M
	v3137 = m.ExcPending
	if v3137 != 0 {
		goto L21
	} else {
		goto L620
	}
L620:
	;
	if v2955 != 0 {
		goto L621
	} else {
		goto L622
	}
L621:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2957))) = base.I64_rotr(v3136, int64(32))
	goto L623
L622:
	;
	goto L623
L623:
	;
	if v2981 != 0 {
		goto L624
	} else {
		goto L625
	}
L624:
	;
	v3168 = base.I32_wrap_i64(int64(base.Ui64(v3136) >> (uint(int64(32)) % 64)))
	goto L626
L625:
	;
	if l4 < int32(0) {
		goto L628
	} else {
		goto L629
	}
L626:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v68)+4)) = uint32(v3136)
	*(*int32)(unsafe.Add(mBase, uint32(v68))) = v3168
	goto L588
L627:
	;
	v3162 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v3161))) = base.I64_rotr(v3136, v3162)
	v3168 = base.I32_wrap_i64(int64(base.Ui64(v3136) >> (uint(v3162) % 64)))
	goto L626
L628:
	;
	v3147 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v3153 = *(*int32)(unsafe.Add(mBase, uint32(v3147+(l4^int32(-1))<<(uint(int32(2))%32))))
	v3161 = v3153
	goto L627
L629:
	;
	goto L630
L630:
	;
	v3155 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v3161 = v3155 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L627
L631:
	;
	F__bt_relbuf(m, v2955)
	mBase = m.M
	v3181 = m.ExcPending
	if v3181 != 0 {
		goto L21
	} else {
		goto L634
	}
L632:
	;
	goto L633
L633:
	;
	if v2981 == int32(0) {
		goto L635
	} else {
		goto L636
	}
L634:
	;
	goto L633
L635:
	;
	F__bt_relbuf(m, l4)
	mBase = m.M
	v3185 = m.ExcPending
	if v3185 != 0 {
		goto L21
	} else {
		goto L638
	}
L636:
	;
	goto L637
L637:
	;
	if v72|v108 != 0 {
		goto L640
	} else {
		goto L641
	}
L638:
	;
	F__bt_relbuf(m, l3)
	mBase = m.M
	v3187 = m.ExcPending
	if v3187 != 0 {
		goto L21
	} else {
		goto L639
	}
L639:
	;
	goto L24
L640:
	;
	F__bt_relbuf(m, l3)
	mBase = m.M
	v3190 = m.ExcPending
	if v3190 != 0 {
		goto L21
	} else {
		goto L643
	}
L641:
	;
	goto L642
L642:
	;
	if l3 < int32(0) {
		goto L645
	} else {
		goto L646
	}
L643:
	;
	goto L24
L644:
	;
	F__bt_relbuf(m, l3)
	mBase = m.M
	v3211 = m.ExcPending
	if v3211 != 0 {
		goto L21
	} else {
		goto L648
	}
L645:
	;
	v3194 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v3200 = *(*int32)(unsafe.Add(mBase, uint32(v3194+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v3209 = v3200
	goto L644
L646:
	;
	goto L647
L647:
	;
	v3202 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v3208 = *(*int32)(unsafe.Add(mBase, uint32(v3202+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v3209 = v3208
	goto L644
L648:
	;
	if v3209 == int32(-1) {
		goto L24
	} else {
		goto L649
	}
L649:
	;
	v3214 = F__bt_getrootheight(m, l0)
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L21
	} else {
		goto L650
	}
L650:
	;
	if v3214 < int32(2) {
		goto L24
	} else {
		goto L651
	}
L651:
	;
	v3218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3218 != 0 {
		goto L652
	} else {
		goto L653
	}
L652:
	;
	v3244 = v3218
	goto L654
L653:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3220 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+152)) = v3220
	v3222 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v49)+144)) = v3222
	v3226 = F_smgropen(m, v49+int32(144), v3219)
	mBase = m.M
	v3227 = m.ExcPending
	if v3227 != 0 {
		goto L21
	} else {
		goto L655
	}
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3244)+16)) = v3209
	goto L24
L655:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3226
	v3230 = *(*int32)(unsafe.Add(mBase, uint32(v3226)+72))
	if v3230 != 0 {
		goto L657
	} else {
		goto L658
	}
L656:
	;
	v3242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3244 = v3242
	goto L654
L657:
	;
	v3238 = v3230
	goto L659
L658:
	;
	v3231 = *(*int32)(unsafe.Add(mBase, uint32(v3226)+76))
	v3232 = *(*int32)(unsafe.Add(mBase, uint32(v3226)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v3231)+4)) = v3232
	v3234 = *(*int32)(unsafe.Add(mBase, uint32(v3226)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3232))) = v3234
	v3236 = *(*int32)(unsafe.Add(mBase, uint32(v3226)+72))
	v3238 = v3236
	goto L659
L659:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3226)+72)) = v3238 + int32(1)
	goto L656
L660:
	;
	F_pfree(m, v105)
	mBase = m.M
	v3293 = m.ExcPending
	if v3293 != 0 {
		goto L21
	} else {
		goto L663
	}
L661:
	;
	goto L662
L662:
	;
	m.G0 = v49 + int32(224)
	return
L663:
	;
	F_pfree(m, v101)
	mBase = m.M
	v3295 = m.ExcPending
	if v3295 != 0 {
		goto L21
	} else {
		goto L664
	}
L664:
	;
	goto L662
L665:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3305 = m.ExcPending
	if v3305 != 0 {
		goto L21
	} else {
		goto L666
	}
L666:
	;
	v3306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+2)))
	v3307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6))))
	v3308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)))
	if l3 < int32(0) {
		goto L668
	} else {
		goto L669
	}
L667:
	;
	v3328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+176)) = v3328 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v49)+172)) = v3327
	*(*int32)(unsafe.Add(mBase, uint32(v49)+168)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v49)+164)) = v3308
	*(*int32)(unsafe.Add(mBase, uint32(v49)+160)) = v3306 | v3307<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(693932), v49+int32(160))
	mBase = m.M
	v3343 = m.ExcPending
	if v3343 != 0 {
		goto L21
	} else {
		goto L671
	}
L668:
	;
	v3312 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v3318 = *(*int32)(unsafe.Add(mBase, uint32(v3312+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v3327 = v3318
	goto L667
L669:
	;
	goto L670
L670:
	;
	v3320 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v3326 = *(*int32)(unsafe.Add(mBase, uint32(v3320+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v3327 = v3326
	goto L667
L671:
	;
	F_errfinish(m, int32(493070), int32(1191), int32(327271))
	mBase = m.M
	v3348 = m.ExcPending
	if v3348 != 0 {
		goto L21
	} else {
		goto L672
	}
L672:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L673:
	;
	v3353 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v49)+4)) = v3353 + int32(4)
	F_errmsg_internal(m, int32(695770), v49)
	mBase = m.M
	v3360 = m.ExcPending
	if v3360 != 0 {
		goto L21
	} else {
		goto L674
	}
L674:
	;
	F_errfinish(m, int32(493070), int32(1704), int32(102442))
	mBase = m.M
	v3365 = m.ExcPending
	if v3365 != 0 {
		goto L21
	} else {
		goto L675
	}
L675:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L676:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3373 = m.ExcPending
	if v3373 != 0 {
		goto L21
	} else {
		goto L677
	}
L677:
	;
	v3374 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+112)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v49)+116)) = v3374 + int32(4)
	F_errmsg_internal(m, int32(695521), v49+int32(112))
	mBase = m.M
	v3383 = m.ExcPending
	if v3383 != 0 {
		goto L21
	} else {
		goto L678
	}
L678:
	;
	F_errfinish(m, int32(493070), int32(1773), int32(102442))
	mBase = m.M
	v3388 = m.ExcPending
	if v3388 != 0 {
		goto L21
	} else {
		goto L679
	}
L679:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L680:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3396 = m.ExcPending
	if v3396 != 0 {
		goto L21
	} else {
		goto L681
	}
L681:
	;
	v3397 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+80)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v49)+84)) = v3397 + int32(4)
	F_errmsg_internal(m, int32(695852), v49+int32(80))
	mBase = m.M
	v3406 = m.ExcPending
	if v3406 != 0 {
		goto L21
	} else {
		goto L682
	}
L682:
	;
	F_errfinish(m, int32(493070), int32(1821), int32(102442))
	mBase = m.M
	v3411 = m.ExcPending
	if v3411 != 0 {
		goto L21
	} else {
		goto L683
	}
L683:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L684:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3419 = m.ExcPending
	if v3419 != 0 {
		goto L21
	} else {
		goto L685
	}
L685:
	;
	v3420 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+96)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v49)+100)) = v3420 + int32(4)
	F_errmsg_internal(m, int32(695604), v49+int32(96))
	mBase = m.M
	v3429 = m.ExcPending
	if v3429 != 0 {
		goto L21
	} else {
		goto L686
	}
L686:
	;
	F_errfinish(m, int32(493070), int32(1834), int32(102442))
	mBase = m.M
	v3434 = m.ExcPending
	if v3434 != 0 {
		goto L21
	} else {
		goto L687
	}
L687:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L688:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3442 = m.ExcPending
	if v3442 != 0 {
		goto L21
	} else {
		goto L689
	}
L689:
	;
	v3443 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+48)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v49)+52)) = v3443 + int32(4)
	F_errmsg_internal(m, int32(695934), v49+int32(48))
	mBase = m.M
	v3452 = m.ExcPending
	if v3452 != 0 {
		goto L21
	} else {
		goto L690
	}
L690:
	;
	F_errfinish(m, int32(493070), int32(1848), int32(102442))
	mBase = m.M
	v3457 = m.ExcPending
	if v3457 != 0 {
		goto L21
	} else {
		goto L691
	}
L691:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L692:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3465 = m.ExcPending
	if v3465 != 0 {
		goto L21
	} else {
		goto L693
	}
L693:
	;
	v3466 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+64)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v49)+68)) = v3466 + int32(4)
	F_errmsg_internal(m, int32(695687), v49-int32(-64))
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
		goto L21
	} else {
		goto L694
	}
L694:
	;
	F_errfinish(m, int32(493070), int32(1860), int32(102442))
	mBase = m.M
	v3480 = m.ExcPending
	if v3480 != 0 {
		goto L21
	} else {
		goto L695
	}
L695:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L696:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3488 = m.ExcPending
	if v3488 != 0 {
		goto L21
	} else {
		goto L697
	}
L697:
	;
	v3489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+32)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v49)+36)) = v3489 + int32(4)
	F_errmsg_internal(m, int32(695604), v49+int32(32))
	mBase = m.M
	v3498 = m.ExcPending
	if v3498 != 0 {
		goto L21
	} else {
		goto L698
	}
L698:
	;
	F_errfinish(m, int32(493070), int32(1881), int32(102442))
	mBase = m.M
	v3503 = m.ExcPending
	if v3503 != 0 {
		goto L21
	} else {
		goto L699
	}
L699:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L700:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3511 = m.ExcPending
	if v3511 != 0 {
		goto L21
	} else {
		goto L701
	}
L701:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3514 = m.ExcPending
	if v3514 != 0 {
		goto L21
	} else {
		goto L702
	}
L702:
	;
	v3515 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3516 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v3517 = *(*int32)(unsafe.Add(mBase, uint32(v2675)))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+24)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v49)+20)) = v3517
	*(*int32)(unsafe.Add(mBase, uint32(v49)+16)) = v3516
	*(*int32)(unsafe.Add(mBase, uint32(v49)+28)) = v3515 + int32(4)
	F_errmsg_internal(m, int32(694295), v49+int32(16))
	mBase = m.M
	v3528 = m.ExcPending
	if v3528 != 0 {
		goto L21
	} else {
		goto L703
	}
L703:
	;
	F_errfinish(m, int32(493070), int32(1904), int32(102442))
	mBase = m.M
	v3533 = m.ExcPending
	if v3533 != 0 {
		goto L21
	} else {
		goto L704
	}
L704:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L705:
	;
	if l3 < int32(0) {
		goto L707
	} else {
		goto L708
	}
L706:
	;
	v3557 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v49)+128)) = v3556
	*(*int32)(unsafe.Add(mBase, uint32(v49)+132)) = v3557 + int32(4)
	F_errmsg_internal(m, int32(693480), v49+int32(128))
	mBase = m.M
	v3566 = m.ExcPending
	if v3566 != 0 {
		goto L21
	} else {
		goto L710
	}
L707:
	;
	v3541 = *(*int32)(unsafe.Add(mBase, _consts[15]))
	v3547 = *(*int32)(unsafe.Add(mBase, uint32(v3541+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v3556 = v3547
	goto L706
L708:
	;
	goto L709
L709:
	;
	v3549 = *(*int32)(unsafe.Add(mBase, _consts[16]))
	v3555 = *(*int32)(unsafe.Add(mBase, uint32(v3549+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v3556 = v3555
	goto L706
L710:
	;
	F_errfinish(m, int32(493070), int32(1283), int32(327271))
	mBase = m.M
	v3571 = m.ExcPending
	if v3571 != 0 {
		goto L21
	} else {
		goto L711
	}
L711:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F__bt_keep_natts_fast(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v70 int32
	_ = v70
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+10)))
	if v18 <= int32(0) {
		v70 = v16
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v70
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v29 = v16
	goto L3
L3:
	;
	v39 = F_index_getattr_2(m, l1, v29, v23, v14+int32(15))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v70 = v18 + int32(1)
	goto L1
L5:
	;
	return int32(0)
L6:
	;
	v45 = F_index_getattr_2(m, l2, v29, v23, v14+int32(14))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+15)))
	v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+14)))
	if v47 != v48 {
		v70 = v29
		goto L1
	} else {
		goto L8
	}
L8:
	;
	if v47 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v56 = v29<<(uint(int32(4))%32) + (v23 + int32(20)) - int32(16)
	v57 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+6)))
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v56)+4)))
	v59 = F_datum_image_eq(m, v39, v45, v57, v58)
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L5
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	if v29 != v18 {
		v29 = v29 + int32(1)
		goto L3
	} else {
		goto L14
	}
L12:
	;
	if v59 == int32(0) {
		v70 = v29
		goto L1
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	goto L4
}
func F__bt_parallel_primscan_schedule(m *base.Module, l0 int32, l1 int32) {
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
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
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
	var v73 int32
	_ = v73
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
	var v89 int32
	_ = v89
	var v100 int32
	_ = v100
	v9 = m.G0
	v11 = v9 - int32(16)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
	v16 = v14 + v15
	v18 = v16 + int32(12)
	v20 = F_LWLockAcquire(m, v18, int32(0))
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v22 != l1 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	F_LWLockRelease(m, v18)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L1
	} else {
		goto L16
	}
L4:
	;
	v24 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v24 != int32(3) {
		goto L3
	} else {
		goto L5
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+8)) = int32(1)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = int64(-1)
	v32 = v16 + int32(40)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v32 + v33<<(uint(int32(2))%32)
	if v33 <= int32(0) {
		goto L3
	} else {
		goto L6
	}
L6:
	;
	v41 = int32(0)
	goto L7
L7:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	v52 = v49 + v41<<(uint(int32(5))%32)
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v53 != int32(-1) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L3
L9:
	;
	v88 = v41 + int32(1)
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
	if v88 < v89 {
		v41 = v88
		goto L7
	} else {
		goto L15
	}
L10:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v32+v41<<(uint(int32(2))%32)))) = v59
	goto L9
L11:
	;
	goto L12
L12:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v66 = v62 + v63*int32(48)
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	*(*int32)(unsafe.Add(mBase, uint32(v61))) = v67
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+12)) = v69 + int32(4)
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v66)))
	if v73&int32(1572864) != 0 {
		goto L9
	} else {
		goto L13
	}
L13:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v66)+44))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v52)+18)))
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v52)+16)))
	F_datumSerialize(m, v76, v73&int32(1), v79, v80, v11+int32(12))
	mBase = m.M
	v84 = m.ExcPending
	if v84 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	goto L9
L15:
	;
	goto L8
L16:
	;
	m.G0 = v11 + int32(16)
	return
}
func F__bt_relandgetbuf(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	if l1 != 0 {
		F_LockBuffer(m, l1, int32(0))
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_ReleaseAndReadBuffer(m, l1, l0, l2)
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				F_LockBuffer(m, v10, l3)
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					F__bt_checkpage(m, l0, v10)
					v15 = m.ExcPending
					if v15 != 0 {
						return int32(0)
					} else {
						return v10
					}
				}
			}
		}
	} else {
		v10 = F_ReleaseAndReadBuffer(m, l1, l0, l2)
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			F_LockBuffer(m, v10, l3)
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				F__bt_checkpage(m, l0, v10)
				v15 = m.ExcPending
				if v15 != 0 {
					return int32(0)
				} else {
					return v10
				}
			}
		}
	}
}
func F__bt_sort_dedup_finish_pending(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
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
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v9 == int32(1) {
		F__bt_buildadd(m, l0, l1, v8, int32(0))
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return
		} else {
			v94 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v94
			*(*int64)(unsafe.Add(mBase, uint32(l2)+28)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v94
			return
		}
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+6)))
		if v17&int32(8192) == int32(0) {
			v34 = v17 & int32(8191)
		} else {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8)+5)))
			if v22&int32(32) == int32(0) {
				v34 = v17 & int32(8191)
			} else {
				v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+2)))
				v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8))))
				v34 = v27 | v28<<(uint(int32(16))%32)
			}
		}
		v36 = v16 * int32(6)
		if int32(1) < v16 {
			v44 = (v34 + v36 + int32(7)) & int32(-8)
		} else {
			v44 = v34
		}
		v45 = F_palloc0(m, v44)
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return
		} else {
			if v34 != 0 {
				v47 = F__emscripten_memcpy_bulkmem(m, v45, v8, v34)
				mBase = m.M
				v48 = v47
			} else {
				v48 = v45
			}
			v49 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+6)))
			v52 = v49&int32(-8192) | v44
			if int32(2) <= v16 {
				*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)) = uint16(v34)
				v56 = int32(8192)
				v57 = v16 | v56
				*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)) = uint16(v57)
				v60 = v52 | v56
				*(*uint16)(unsafe.Add(mBase, uint32(v48)+6)) = uint16(v60)
				v63 = int32(base.Ui32(v34) >> (uint(int32(16)) % 32))
				*(*uint16)(unsafe.Add(mBase, uint32(v48))) = uint16(v63)
				if v36 != 0 {
					v66 = F__emscripten_memcpy_bulkmem(m, v48+v34, v15, v36)
					mBase = m.M
				} else {
				}
			} else {
				v69 = v52 & int32(57343)
				*(*uint16)(unsafe.Add(mBase, uint32(v48)+6)) = uint16(v69)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				*(*int32)(unsafe.Add(mBase, uint32(v48))) = v71
				v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v15)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v48)+4)) = uint16(v73)
			}
			v75 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+6)))
			v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+2)))
			v79 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48))))
			F__bt_buildadd(m, l0, l1, v48, v75&int32(8191)-(v78|v79<<(uint(int32(16))%32)))
			mBase = m.M
			v85 = m.ExcPending
			if v85 != 0 {
				return
			} else {
				F_pfree(m, v48)
				mBase = m.M
				v87 = m.ExcPending
				if v87 != 0 {
					return
				} else {
					v94 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v94
					*(*int64)(unsafe.Add(mBase, uint32(l2)+28)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v94
					return
				}
			}
		}
	}
}
func F__bt_start_prim_scan(m *base.Module, l0 int32) int32 {
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
	var v13 int32
	_ = v13
	v2 = int32(0)
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	*(*uint16)(unsafe.Add(mBase, uint32(v3)+18)) = uint16(v2)
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3)+17)))
	if v6 != 0 {
		return v6
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+88))
		if v7 == int32(0) {
			return v6
		} else {
			F__bt_parallel_done(m, l0)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v6
			}
		}
	}
}
func F__bt_steppage(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+36))
	if int32(0) < v7 {
		F__bt_killitems(m, l0)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
			if v14 < int32(0) {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
				if v52 != 0 {
					F_ReleaseBuffer(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
						if l1 == int32(1) {
							v61 = int32(68)
						} else {
							v61 = int32(64)
						}
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if l1 != v65 {
							v67 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
						} else {
						}
						v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							return v70
						}
					}
				} else {
					if l1 == int32(1) {
						v61 = int32(68)
					} else {
						v61 = int32(64)
					}
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
					if l1 != v65 {
						v67 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
					} else {
					}
					v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						return v70
					}
				}
			} else {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
				if v19 != 0 {
					F_IncrBufferRefCount(m, v19)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
						v28 = v24*int32(10) + int32(58)
						if v28 != 0 {
							v29 = F__emscripten_memcpy_bulkmem(m, v6+int32(13688), v6+int32(56), v28)
							mBase = m.M
						} else {
						}
						v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
						if v31 != 0 {
							v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
							if v33 != 0 {
								v34 = F__emscripten_memcpy_bulkmem(m, v31, v32, v33)
								mBase = m.M
							} else {
							}
						} else {
						}
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[59]))) = v36
						*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
						v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)))
						if v40 != int32(1) {
						} else {
							v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							if v43 == int32(1) {
								v46 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[60]))) = uint8(v46)
							} else {
								v48 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[61]))) = uint8(v48)
							}
						}
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
						if v52 != 0 {
							F_ReleaseBuffer(m, v52)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
								if l1 == int32(1) {
									v61 = int32(68)
								} else {
									v61 = int32(64)
								}
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
								v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
								v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
								if l1 != v65 {
									v67 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
								} else {
								}
								v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
								mBase = m.M
								v71 = m.ExcPending
								if v71 != 0 {
									return int32(0)
								} else {
									return v70
								}
							}
						} else {
							if l1 == int32(1) {
								v61 = int32(68)
							} else {
								v61 = int32(64)
							}
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							if l1 != v65 {
								v67 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
							} else {
							}
							v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								return v70
							}
						}
					}
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
					v28 = v24*int32(10) + int32(58)
					if v28 != 0 {
						v29 = F__emscripten_memcpy_bulkmem(m, v6+int32(13688), v6+int32(56), v28)
						mBase = m.M
					} else {
					}
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
					if v31 != 0 {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
						if v33 != 0 {
							v34 = F__emscripten_memcpy_bulkmem(m, v31, v32, v33)
							mBase = m.M
						} else {
						}
					} else {
					}
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[59]))) = v36
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)))
					if v40 != int32(1) {
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if v43 == int32(1) {
							v46 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[60]))) = uint8(v46)
						} else {
							v48 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[61]))) = uint8(v48)
						}
					}
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
					if v52 != 0 {
						F_ReleaseBuffer(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
							if l1 == int32(1) {
								v61 = int32(68)
							} else {
								v61 = int32(64)
							}
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							if l1 != v65 {
								v67 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
							} else {
							}
							v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								return v70
							}
						}
					} else {
						if l1 == int32(1) {
							v61 = int32(68)
						} else {
							v61 = int32(64)
						}
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if l1 != v65 {
							v67 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
						} else {
						}
						v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							return v70
						}
					}
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
		if v14 < int32(0) {
			v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
			if v52 != 0 {
				F_ReleaseBuffer(m, v52)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
					if l1 == int32(1) {
						v61 = int32(68)
					} else {
						v61 = int32(64)
					}
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
					if l1 != v65 {
						v67 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
					} else {
					}
					v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						return v70
					}
				}
			} else {
				if l1 == int32(1) {
					v61 = int32(68)
				} else {
					v61 = int32(64)
				}
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
				if l1 != v65 {
					v67 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
				} else {
				}
				v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					return v70
				}
			}
		} else {
			v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
			if v19 != 0 {
				F_IncrBufferRefCount(m, v19)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
					v28 = v24*int32(10) + int32(58)
					if v28 != 0 {
						v29 = F__emscripten_memcpy_bulkmem(m, v6+int32(13688), v6+int32(56), v28)
						mBase = m.M
					} else {
					}
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
					if v31 != 0 {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
						if v33 != 0 {
							v34 = F__emscripten_memcpy_bulkmem(m, v31, v32, v33)
							mBase = m.M
						} else {
						}
					} else {
					}
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[59]))) = v36
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
					v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)))
					if v40 != int32(1) {
					} else {
						v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if v43 == int32(1) {
							v46 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[60]))) = uint8(v46)
						} else {
							v48 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[61]))) = uint8(v48)
						}
					}
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
					if v52 != 0 {
						F_ReleaseBuffer(m, v52)
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
							if l1 == int32(1) {
								v61 = int32(68)
							} else {
								v61 = int32(64)
							}
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							if l1 != v65 {
								v67 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
							} else {
							}
							v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								return v70
							}
						}
					} else {
						if l1 == int32(1) {
							v61 = int32(68)
						} else {
							v61 = int32(64)
						}
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if l1 != v65 {
							v67 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
						} else {
						}
						v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							return v70
						}
					}
				}
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
				v28 = v24*int32(10) + int32(58)
				if v28 != 0 {
					v29 = F__emscripten_memcpy_bulkmem(m, v6+int32(13688), v6+int32(56), v28)
					mBase = m.M
				} else {
				}
				v31 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
				if v31 != 0 {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
					if v33 != 0 {
						v34 = F__emscripten_memcpy_bulkmem(m, v31, v32, v33)
						mBase = m.M
					} else {
					}
				} else {
				}
				v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[59]))) = v36
				*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
				v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)))
				if v40 != int32(1) {
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
					if v43 == int32(1) {
						v46 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[60]))) = uint8(v46)
					} else {
						v48 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_consts[61]))) = uint8(v48)
					}
				}
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
				if v52 != 0 {
					F_ReleaseBuffer(m, v52)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
						if l1 == int32(1) {
							v61 = int32(68)
						} else {
							v61 = int32(64)
						}
						v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
						v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
						v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if l1 != v65 {
							v67 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
						} else {
						}
						v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							return v70
						}
					}
				} else {
					if l1 == int32(1) {
						v61 = int32(68)
					} else {
						v61 = int32(64)
					}
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v6+v61)))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
					if l1 != v65 {
						v67 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v67)
					} else {
					}
					v70 = F__bt_readnextpage(m, l0, v63, v64, l1, int32(0))
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						return v70
					}
				}
			}
		}
	}
}
func F__bt_unlockbuf(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	F_LockBuffer(m, l0, int32(0))
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		return
	}
}
