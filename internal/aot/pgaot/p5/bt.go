package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F__bt_binsrch(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
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
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v113 int32
	_ = v113
	if l2 < int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v28) {
		goto L5
	} else {
		goto L6
	}
L2:
	;
	v13 = *(*int32)(unsafe.Add(mBase, _c_F__bt_binsrch[0]))
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v13+(l2^int32(-1))<<(uint(int32(2))%32))))
	v27 = v19
	goto L1
L3:
	;
	goto L4
L4:
	;
	v21 = *(*int32)(unsafe.Add(mBase, _c_F__bt_binsrch[1]))
	v27 = v21 + l2<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	v36 = int32(base.Ui32(v28+int32(_a_F__bt_binsrch_0)) >> (uint(int32(2)) % 32))
	goto L7
L6:
	;
	v36 = int32(0)
	goto L7
L7:
	;
	v41 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v27)+16)))
	v42 = v27 + v41
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+4))
	if v43 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v44 = int32(2)
	goto L10
L9:
	;
	v44 = int32(1)
	goto L10
L10:
	;
	if base.Ui32(v44) <= base.Ui32(v36&int32(_a_F__bt_binsrch_1)) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v47 = v36 + int32(1)
	if base.Ui32(v44) < base.Ui32(v47&int32(_a_F__bt_binsrch_1)) {
		goto L14
	} else {
		goto L15
	}
L12:
	;
	v113 = v44
	goto L13
L13:
	;
	return v113 & int32(_a_F__bt_binsrch_1)
L14:
	;
	v51 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+3)))
	v56 = v44
	v57 = v47
	goto L17
L15:
	;
	v87 = v44
	goto L16
L16:
	;
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42)+12)))
	if v94&int32(1) != 0 {
		goto L28
	} else {
		goto L29
	}
L17:
	;
	v68 = int32(base.Ui32((v57-v56)&int32(_a_F__bt_binsrch_2))>>(uint(int32(1))%32)) + v56
	v71 = F__bt_compare(m, l0, l1, v27, v68&int32(_a_F__bt_binsrch_1))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v87 = v81
	goto L16
L19:
	;
	return int32(0)
L20:
	;
	v75 = base.B2i32(v71 < v51^int32(1))
	if v71 < v51^int32(1) {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v76 = v68
	goto L23
L22:
	;
	v76 = v57
	goto L23
L23:
	;
	if v71 < v51^int32(1) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v81 = v56
	goto L26
L25:
	;
	v81 = v68 + int32(1)
	goto L26
L26:
	;
	if base.Ui32(v81&int32(_a_F__bt_binsrch_1)) < base.Ui32(v76&int32(_a_F__bt_binsrch_1)) {
		v56 = v81
		v57 = v76
		goto L17
	} else {
		goto L27
	}
L27:
	;
	goto L18
L28:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	return (v87 - v97) & int32(_a_F__bt_binsrch_1)
L29:
	;
	goto L30
L30:
	;
	v113 = v87 - int32(1)
	goto L13
}
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
	var v39 int32
	_ = v39
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
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
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
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	v4 = int32(0)
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v18 <= v4 {
	} else {
		v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
		v33 = v25
		v39 = v4
		for {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
			v46 = v43 + v33*int32(6)
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
			v50 = v47 + v33<<(uint(int32(3))%32)
			v51 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+16)))
			v52 = v51 + v39
			v57 = l0 + int32(20) + v52&int32(_a_F__bt_bottomupdel_finish_pending_0)<<(uint(int32(2))%32)
			v58 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
			v61 = l0 + v58&int32(_a_F__bt_bottomupdel_finish_pending_1)
			v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+7)))
			if v62&int32(32) != 0 {
				v65 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
				if v65&int32(_a_F__bt_bottomupdel_finish_pending_2) != 0 {
					v90 = v65 & int32(4095)
					v91 = int32(0)
					if int32(2) <= v18 {
						v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+2)))
						v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61))))
						v97 = int32(16)
						v100 = v95 + (v61 + v96<<(uint(v97)%32))
						v101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100))))
						v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v100)+2)))
						v105 = v101<<(uint(v97)%32) | v104
						v108 = int32(6)
						v110 = v100 + int32(base.Ui32(v90)>>(uint(int32(1))%32))*v108
						v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110))))
						v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110)+2)))
						v115 = v111<<(uint(v97)%32) | v114
						v119 = v100 + v90*v108
						v122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119-v108))))
						v127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v119-int32(4)))))
						v132 = base.B2i32(v115 == v105)
						v133 = base.B2i32(v105 != v115) & base.B2i32(v115 == v122<<(uint(v97)%32)|v127)
					} else {
						v132 = v91
						v133 = v91
					}
					if v90 == int32(0) {
						v227 = v33
					} else {
						v138 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+2)))
						v139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61))))
						v143 = v138 + (v61 + v139<<(uint(int32(16))%32))
						v144 = *(*int32)(unsafe.Add(mBase, uint32(v143)))
						*(*int32)(unsafe.Add(mBase, uint32(v50))) = v144
						v146 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v143)+4)))
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v146)
						v148 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
						*(*uint16)(unsafe.Add(mBase, uint32(v50)+6)) = uint16(v148)
						v150 = int32(6)
						*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v150)
						v152 = int32(1)
						v153 = v90 - v152
						v154 = int32(0)
						v157 = v132 | v133&base.B2i32(v153 == v154)
						*(*uint8)(unsafe.Add(mBase, uint32(v46)+3)) = uint8(v157)
						*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)) = uint8(v154)
						*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v52)
						v163 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
						v165 = v163 + v152
						*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v165
						if v90 == v152 {
							v227 = v165
						} else {
							v172 = v46
							v173 = v152
							v175 = v50
							for {
								v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+2)))
								v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61))))
								v192 = int32(6)
								v194 = v186 + (v61 + v187<<(uint(int32(16))%32)) + v173*v192
								v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)))
								*(*int32)(unsafe.Add(mBase, uint32(v175)+8)) = v195
								v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+4)))
								*(*uint16)(unsafe.Add(mBase, uint32(v175)+12)) = uint16(v197)
								v199 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								*(*uint16)(unsafe.Add(mBase, uint32(v175)+14)) = uint16(v199)
								*(*uint16)(unsafe.Add(mBase, uint32(v172)+10)) = uint16(v192)
								v204 = v133 & base.B2i32(v173 == v153)
								*(*uint8)(unsafe.Add(mBase, uint32(v172)+9)) = uint8(v204)
								v206 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v172)+8)) = uint8(v206)
								*(*uint16)(unsafe.Add(mBase, uint32(v172)+6)) = uint16(v52)
								v209 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
								v210 = int32(1)
								v211 = v209 + v210
								*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v211
								v218 = v173 + v210
								if v218 != v90 {
									v172 = v172 + v192
									v173 = v218
									v175 = v175 + int32(8)
									continue
								} else {
									break
								}
								break
							}
							v227 = v211
						}
					}
				} else {
					v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v69)
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
					*(*int32)(unsafe.Add(mBase, uint32(v50))) = v71
					v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					*(*uint16)(unsafe.Add(mBase, uint32(v50)+6)) = uint16(v73)
					*(*uint8)(unsafe.Add(mBase, uint32(v46)+3)) = uint8(base.B2i32(int32(1) < v18))
					v76 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)) = uint8(v76)
					*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v52)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
					v83 = int32(base.Ui32(v79)>>(uint(int32(17))%32)) + int32(4)
					*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v83)
					v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
					v87 = v85 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v87
					v227 = v87
				}
			} else {
				v69 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+4)))
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+4)) = uint16(v69)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
				*(*int32)(unsafe.Add(mBase, uint32(v50))) = v71
				v73 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				*(*uint16)(unsafe.Add(mBase, uint32(v50)+6)) = uint16(v73)
				*(*uint8)(unsafe.Add(mBase, uint32(v46)+3)) = uint8(base.B2i32(int32(1) < v18))
				v76 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v46)+2)) = uint8(v76)
				*(*uint16)(unsafe.Add(mBase, uint32(v46))) = uint16(v52)
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
				v83 = int32(base.Ui32(v79)>>(uint(int32(17))%32)) + int32(4)
				*(*uint16)(unsafe.Add(mBase, uint32(v46)+4)) = uint16(v83)
				v85 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				v87 = v85 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = v87
				v227 = v87
			}
			v238 = v39 + int32(1)
			v239 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
			if v238 < v239 {
				v33 = v227
				v39 = v238
				continue
			} else {
				break
			}
			break
		}
		if v18 <= int32(1) {
		} else {
			v243 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			*(*uint16)(unsafe.Add(mBase, uint32(l1+v243<<(uint(int32(2))%32))+46)) = uint16(v239)
			v248 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
			*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = v248 + int32(1)
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
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
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
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v269 int32
	_ = v269
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
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
	return v317
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
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2+l3<<(uint(int32(2))%32))+20))
	v38 = l2 + v35&int32(_a_F__bt_compare_0)
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
		v317 = v18
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	v317 = int32(1)
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
	if v44&int32(_a_F__bt_compare_1) != 0 {
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
		v317 = v192
		goto L1
	} else {
		goto L57
	}
L20:
	;
	v79 = int32(*(*int16)(unsafe.Add(mBase, uint32(v68)+4)))
	v80 = int32(*(*int16)(unsafe.Add(mBase, uint32(v38)+6)))
	if int32(0) <= v80 {
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
	v85 = v19 + int32(4) + v79<<(uint(int32(4))%32)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v85)))
	if v86 < int32(0) {
		goto L26
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v115 = int32(1)
	v116 = v79 - v115
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60+v116>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v120)>>(uint(v116&int32(7))%32))&v115 == int32(0) {
		goto L24
	} else {
		goto L40
	}
L30:
	;
	v89 = v60 + v86
	v90 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+6)))
	if v90 != int32(1) {
		v135 = v89
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v93 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v85)+4)))
	switch v93 - int32(1) {
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
	v104 = m.ExcPending
	if v104 != 0 {
		goto L36
	} else {
		goto L37
	}
L33:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	v135 = v98
	goto L25
L34:
	;
	v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89))))
	v135 = v97
	goto L25
L35:
	;
	v96 = int32(*(*int8)(unsafe.Add(mBase, uint32(v89))))
	v135 = v96
	goto L25
L36:
	;
	return int32(0)
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = base.I32_extend16_s(v93)
	F_errmsg_internal(m, int32(_a_F__bt_compare_2), v16)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L36
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(_a_F__bt_compare_3), int32(70), int32(_a_F__bt_compare_4))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
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
	v317 = v145
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
	v317 = v153
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
		v317 = v169
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
	if v194&int32(_a_F__bt_compare_1) == int32(0) {
		v220 = v38
		goto L60
	} else {
		goto L61
	}
L58:
	;
	v304 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v302|(v304|base.B2i32(v181 != v53)) == int32(0) {
		goto L81
	} else {
		goto L82
	}
L59:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v300 != 0 {
		v317 = v192
		goto L1
	} else {
		goto L80
	}
L60:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v221 == int32(0) {
		v302 = v220
		goto L58
	} else {
		goto L66
	}
L61:
	;
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	if v199&int32(_a_F__bt_compare_1) == int32(0) {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	if v199&int32(_a_F__bt_compare_5) == int32(0) {
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
	v220 = v213 + (v38 + v214<<(uint(int32(16))%32))
	goto L60
L65:
	;
	v220 = v38 + v194&int32(_a_F__bt_compare_6) - int32(6)
	goto L60
L66:
	;
	v227 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+2)))
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221))))
	v229 = int32(16)
	v231 = v227 | v228<<(uint(v229)%32)
	v232 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+2)))
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220))))
	v236 = v232 | v233<<(uint(v229)%32)
	if base.Ui32(v231) < base.Ui32(v236) {
		v247 = int32(-1)
		goto L68
	} else {
		goto L69
	}
L67:
	;
	if v247 <= int32(0) {
		v317 = v247
		goto L1
	} else {
		goto L72
	}
L68:
	;
	goto L67
L69:
	;
	if base.Ui32(v236) < base.Ui32(v231) {
		v247 = int32(1)
		goto L68
	} else {
		goto L70
	}
L70:
	;
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v221)+4)))
	v242 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v220)+4)))
	if base.Ui32(v241) < base.Ui32(v242) {
		v247 = int32(-1)
		goto L68
	} else {
		goto L71
	}
L71:
	;
	v247 = base.B2i32(base.Ui32(v242) < base.Ui32(v241))
	goto L68
L72:
	;
	v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+7)))
	if v250&int32(32) == int32(0) {
		v317 = v247
		goto L1
	} else {
		goto L73
	}
L73:
	;
	v255 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+4)))
	if v255&int32(_a_F__bt_compare_1) == int32(0) {
		v317 = v247
		goto L1
	} else {
		goto L74
	}
L74:
	;
	v260 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v261 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38)+2)))
	v262 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v38))))
	v263 = int32(16)
	v269 = int32(6)
	v273 = v261 + (v38 + v262<<(uint(v263)%32)) + v255&int32(4095)*v269 - v269
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+2)))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260))))
	v281 = v277 | v278<<(uint(v263)%32)
	v282 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v273)+2)))
	v283 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v273))))
	v286 = v282 | v283<<(uint(v263)%32)
	if base.Ui32(v281) < base.Ui32(v286) {
		v297 = int32(-1)
		goto L76
	} else {
		goto L77
	}
L75:
	;
	v317 = base.B2i32(int32(0) < v297)
	goto L1
L76:
	;
	goto L75
L77:
	;
	if base.Ui32(v286) < base.Ui32(v281) {
		v297 = int32(1)
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v260)+4)))
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v273)+4)))
	if base.Ui32(v291) < base.Ui32(v292) {
		v297 = int32(-1)
		goto L76
	} else {
		goto L79
	}
L79:
	;
	v297 = base.B2i32(base.Ui32(v292) < base.Ui32(v291))
	goto L76
L80:
	;
	v302 = int32(0)
	goto L58
L81:
	;
	v310 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v310 != 0 {
		v317 = v192
		goto L1
	} else {
		goto L84
	}
L82:
	;
	goto L83
L83:
	;
	v317 = int32(0)
	goto L1
L84:
	;
	goto L83
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
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v79 int32
	_ = v79
	var v88 int32
	_ = v88
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v257 int32
	_ = v257
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v365 int32
	_ = v365
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v453 int32
	_ = v453
	var v479 int32
	_ = v479
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v517 int32
	_ = v517
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v578 int32
	_ = v578
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v586 int32
	_ = v586
	var v595 int32
	_ = v595
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v638 int32
	_ = v638
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v651 int32
	_ = v651
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v664 int32
	_ = v664
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v676 int32
	_ = v676
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v713 int32
	_ = v713
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v723 int32
	_ = v723
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v733 int32
	_ = v733
	var v736 int32
	_ = v736
	var v743 int32
	_ = v743
	var v765 int32
	_ = v765
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v784 int32
	_ = v784
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v819 int64
	_ = v819
	var v828 int32
	_ = v828
	var v834 int32
	_ = v834
	var v836 int32
	_ = v836
	var v842 int32
	_ = v842
	var v843 int32
	_ = v843
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v861 int32
	_ = v861
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v882 int32
	_ = v882
	var v888 int32
	_ = v888
	var v913 int32
	_ = v913
	var v917 int32
	_ = v917
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v959 int32
	_ = v959
	var v965 int32
	_ = v965
	var v975 int32
	_ = v975
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v990 int32
	_ = v990
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1020 int32
	_ = v1020
	var v1028 int32
	_ = v1028
	var v1029 int32
	_ = v1029
	var v1043 int32
	_ = v1043
	var v1050 int32
	_ = v1050
	var v1058 int32
	_ = v1058
	var v1064 int32
	_ = v1064
	var v1068 int32
	_ = v1068
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1098 int32
	_ = v1098
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1147 int32
	_ = v1147
	var v1152 int32
	_ = v1152
	var v1157 int32
	_ = v1157
	var v1158 int32
	_ = v1158
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1175 int32
	_ = v1175
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1182 int32
	_ = v1182
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1224 int32
	_ = v1224
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1243 int32
	_ = v1243
	var v1252 int32
	_ = v1252
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1268 int32
	_ = v1268
	var v1273 int32
	_ = v1273
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1320 int32
	_ = v1320
	var v1321 int32
	_ = v1321
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1330 int32
	_ = v1330
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1354 int32
	_ = v1354
	var v1355 int32
	_ = v1355
	var v1361 int32
	_ = v1361
	var v1367 int32
	_ = v1367
	var v1377 int32
	_ = v1377
	var v1387 int32
	_ = v1387
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1434 int32
	_ = v1434
	var v1441 int32
	_ = v1441
	var v1449 int32
	_ = v1449
	var v1455 int32
	_ = v1455
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1463 int32
	_ = v1463
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1492 int32
	_ = v1492
	var v1495 int32
	_ = v1495
	var v1499 int32
	_ = v1499
	var v1501 int32
	_ = v1501
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1511 int32
	_ = v1511
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1517 int32
	_ = v1517
	var v1520 int32
	_ = v1520
	var v1521 int32
	_ = v1521
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1535 int32
	_ = v1535
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1564 int32
	_ = v1564
	var v1566 int32
	_ = v1566
	var v1568 int32
	_ = v1568
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1573 int32
	_ = v1573
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1591 int32
	_ = v1591
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1613 int32
	_ = v1613
	var v1615 int32
	_ = v1615
	var v1620 int32
	_ = v1620
	var v1622 int32
	_ = v1622
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1634 int32
	_ = v1634
	var v1643 int32
	_ = v1643
	var v1654 int32
	_ = v1654
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1709 int32
	_ = v1709
	var v1710 int32
	_ = v1710
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1753 int32
	_ = v1753
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1765 int32
	_ = v1765
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1782 int32
	_ = v1782
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1792 int32
	_ = v1792
	var v1795 int32
	_ = v1795
	var v1796 int32
	_ = v1796
	var v1797 int64
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1805 int32
	_ = v1805
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1811 int32
	_ = v1811
	var v1815 int32
	_ = v1815
	var v1819 int32
	_ = v1819
	var v1820 int32
	_ = v1820
	var v1823 int32
	_ = v1823
	var v1833 int32
	_ = v1833
	var v1837 int32
	_ = v1837
	var v1841 int32
	_ = v1841
	var v1842 int32
	_ = v1842
	var v1847 int32
	_ = v1847
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int64
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1864 int32
	_ = v1864
	var v1865 int32
	_ = v1865
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1881 int32
	_ = v1881
	var v1885 int32
	_ = v1885
	var v1891 int32
	_ = v1891
	var v1912 int32
	_ = v1912
	var v1916 int32
	_ = v1916
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1928 int32
	_ = v1928
	var v1929 int32
	_ = v1929
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1939 int32
	_ = v1939
	var v1941 int32
	_ = v1941
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1951 int32
	_ = v1951
	var v1952 int32
	_ = v1952
	var v1958 int32
	_ = v1958
	var v1964 int32
	_ = v1964
	var v1974 int32
	_ = v1974
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1986 int32
	_ = v1986
	var v1987 int32
	_ = v1987
	var v1992 int32
	_ = v1992
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2005 int32
	_ = v2005
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2022 int32
	_ = v2022
	var v2030 int32
	_ = v2030
	var v2031 int32
	_ = v2031
	var v2034 int32
	_ = v2034
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2049 int32
	_ = v2049
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2064 int32
	_ = v2064
	var v2065 int32
	_ = v2065
	var v2075 int32
	_ = v2075
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2083 int32
	_ = v2083
	var v2086 int32
	_ = v2086
	var v2089 int32
	_ = v2089
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2096 int32
	_ = v2096
	var v2100 int32
	_ = v2100
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2116 int32
	_ = v2116
	var v2117 int32
	_ = v2117
	var v2123 int32
	_ = v2123
	var v2129 int32
	_ = v2129
	var v2139 int32
	_ = v2139
	var v2144 int32
	_ = v2144
	var v2150 int32
	_ = v2150
	var v2182 int32
	_ = v2182
	var v2183 int32
	_ = v2183
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2195 int32
	_ = v2195
	var v2198 int32
	_ = v2198
	var v2200 int32
	_ = v2200
	var v2205 int32
	_ = v2205
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2213 int32
	_ = v2213
	var v2216 int32
	_ = v2216
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2221 int32
	_ = v2221
	var v2225 int32
	_ = v2225
	var v2230 int32
	_ = v2230
	var v2234 int32
	_ = v2234
	var v2238 int32
	_ = v2238
	var v2241 int64
	_ = v2241
	var v2242 int32
	_ = v2242
	var v2246 int32
	_ = v2246
	var v2248 int32
	_ = v2248
	var v2253 int32
	_ = v2253
	var v2255 int32
	_ = v2255
	var v2257 int32
	_ = v2257
	var v2264 int32
	_ = v2264
	var v2268 int32
	_ = v2268
	var v2273 int32
	_ = v2273
	v8 = int32(0)
	v28 = m.G0
	v30 = v28 - int32(848)
	m.G0 = v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	if v35 < v8 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	m.G0 = v30 + int32(848)
	return
L2:
	;
	v765 = int32(1)
	if base.B2i32(v743|(l4^v765) != v765)|l3 != 0 {
		goto L1
	} else {
		goto L99
	}
L3:
	;
	v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+16)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v54+v53)+4))
	if v56 != 0 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v39 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[0]))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v39+(v35^int32(-1))<<(uint(int32(2))%32))))
	v53 = v45
	goto L3
L5:
	;
	goto L6
L6:
	;
	v47 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[1]))
	v53 = v47 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L3
L7:
	;
	v57 = int32(2)
	goto L9
L8:
	;
	v57 = int32(1)
	goto L9
L9:
	;
	v58 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v58) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v66 = int32(base.Ui32(v58+int32(_a_F__bt_delete_or_dedup_one_page_0)) >> (uint(int32(2)) % 32))
	goto L12
L11:
	;
	v66 = int32(0)
	goto L12
L12:
	;
	v68 = v66 & int32(_a_F__bt_delete_or_dedup_one_page_1)
	if base.Ui32(v68) < base.Ui32(v57) {
		v743 = l5
		goto L2
	} else {
		goto L13
	}
L13:
	;
	v79 = v57
	v88 = v8
	goto L14
L14:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v53+int32(20)+v79&int32(_a_F__bt_delete_or_dedup_one_page_1)<<(uint(int32(2))%32))))
	v105 = int32(_a_F__bt_delete_or_dedup_one_page_2)
	if v104&v105 == v105 {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	if v115 <= int32(0) {
		v743 = l5
		goto L2
	} else {
		goto L20
	}
L16:
	;
	v109 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v30+v88<<(uint(v109)%32)))) = uint16(v79)
	v115 = v88 + v109
	goto L18
L17:
	;
	v115 = v88
	goto L18
L18:
	;
	v117 = v79 + int32(1)
	if base.Ui32(v117&int32(_a_F__bt_delete_or_dedup_one_page_1)) <= base.Ui32(v68) {
		v79 = v117
		v88 = v115
		goto L14
	} else {
		goto L19
	}
L19:
	;
	goto L15
L20:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v35 < int32(0) {
		goto L22
	} else {
		goto L23
	}
L21:
	;
	v143 = v115 + int32(1)
	v146 = F_palloc(m, v143<<(uint(int32(2))%32))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	v127 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[0]))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v127+(v35^int32(-1))<<(uint(int32(2))%32))))
	v141 = v133
	goto L21
L23:
	;
	goto L24
L24:
	;
	v135 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[1]))
	v141 = v135 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L21
L25:
	;
	return
L26:
	;
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+2)))
	v149 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	*(*int32)(unsafe.Add(mBase, uint32(v146))) = v148 | v149<<(uint(int32(16))%32)
	v155 = v141 + int32(20)
	v162 = int32(1)
	v169 = v143
	v171 = v146
	v175 = v8
	goto L27
L27:
	;
	v187 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v175<<(uint(int32(1))%32)))))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v155+v187<<(uint(int32(2))%32))))
	v194 = v141 + v191&int32(_a_F__bt_delete_or_dedup_one_page_3)
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194)+7)))
	if v195&int32(32) != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	F_pg_qsort(m, v374, v365, int32(4), int32(208))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L25
	} else {
		goto L55
	}
L29:
	;
	v388 = v175 + int32(1)
	if v388 != v115 {
		v162 = v365
		v169 = v372
		v171 = v374
		v175 = v388
		goto L27
	} else {
		goto L54
	}
L30:
	;
	v223 = v198 & int32(4095)
	v224 = v162 + v223
	if v169 < v224 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+4)))
	if v198&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v203 = v162 + int32(1)
	if v169 < v203 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v207 = F_repalloc(m, v171, v169<<(uint(int32(3))%32))
	mBase = m.M
	v208 = m.ExcPending
	if v208 != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	v211 = v169
	v212 = v171
	goto L37
L37:
	;
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+2)))
	v217 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194))))
	*(*int32)(unsafe.Add(mBase, uint32(v212+v162<<(uint(int32(2))%32)))) = v216 | v217<<(uint(int32(16))%32)
	v365 = v203
	v372 = v211
	v374 = v212
	goto L29
L38:
	;
	v211 = v169 << (uint(int32(1)) % 32)
	v212 = v207
	goto L37
L39:
	;
	v227 = v169 << (uint(int32(1)) % 32)
	if v224 < v227 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v234 = v169
	v235 = v171
	goto L41
L41:
	;
	if v223 == int32(0) {
		v365 = v162
		v372 = v234
		v374 = v235
		goto L29
	} else {
		goto L46
	}
L42:
	;
	v229 = v227
	goto L44
L43:
	;
	v229 = v224
	goto L44
L44:
	;
	v232 = F_repalloc(m, v171, v229<<(uint(int32(2))%32))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	v234 = v229
	v235 = v232
	goto L41
L46:
	;
	v238 = int32(0)
	if v223 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v251 = v162
	v254 = v238
	v257 = int32(0)
	goto L50
L48:
	;
	v318 = v162
	v321 = v238
	goto L49
L49:
	;
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+2)))
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194))))
	v345 = int32(16)
	v351 = v343 + (v194 + v344<<(uint(v345)%32)) + v321*int32(6)
	v352 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351))))
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v351)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v235+v318<<(uint(int32(2))%32)))) = v352<<(uint(v345)%32) | v355
	v365 = v318 + int32(1)
	v372 = v234
	v374 = v235
	goto L29
L50:
	;
	v273 = int32(2)
	v275 = v235 + v251<<(uint(v273)%32)
	v277 = v254 * int32(6)
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+2)))
	v279 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194))))
	v280 = int32(16)
	v284 = v277 + (v278 + (v194 + v279<<(uint(v280)%32)))
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284))))
	v288 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v275))) = v285<<(uint(v280)%32) | v288
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+2)))
	v292 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194))))
	v297 = v291 + (v194 + v292<<(uint(v280)%32)) + v277
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+6)))
	v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v297)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v275)+4)) = v298<<(uint(v280)%32) | v301
	v305 = v254 + v273
	v307 = v251 + v273
	v309 = v257 + v273
	if v309 != v223&int32(4094) {
		v251 = v307
		v254 = v305
		v257 = v309
		goto L50
	} else {
		goto L52
	}
L51:
	;
	if v223&int32(1) == int32(0) {
		v365 = v307
		v372 = v234
		v374 = v235
		goto L29
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v318 = v307
	v321 = v305
	goto L49
L54:
	;
	goto L28
L55:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v365) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v405 = int32(1)
	v406 = int32(0)
	goto L59
L57:
	;
	v453 = v365
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+820)) = l0
	if v35 < int32(0) {
		goto L66
	} else {
		goto L67
	}
L59:
	;
	v425 = int32(2)
	v428 = *(*int32)(unsafe.Add(mBase, uint32(v374+v405<<(uint(v425)%32))))
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v374+v406<<(uint(v425)%32))))
	if v428 == v432 {
		v441 = v406
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v453 = v441 + int32(1)
	goto L58
L61:
	;
	v444 = v405 + int32(1)
	if v444 != v365 {
		v405 = v444
		v406 = v441
		goto L59
	} else {
		goto L64
	}
L62:
	;
	v435 = v406 + int32(1)
	if v405 == v435 {
		v441 = v405
		goto L61
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v374+v435<<(uint(int32(2))%32)))) = v428
	v441 = v435
	goto L61
L64:
	;
	goto L60
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+832)) = int64(0)
	v497 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+828)) = uint8(v497)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+824)) = v494
	v501 = F_palloc(m, int32(_a_F__bt_delete_or_dedup_one_page_5))
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L25
	} else {
		goto L69
	}
L66:
	;
	v479 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[2]))
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v479+(v35^int32(-1))<<(uint(int32(6))%32))+16))
	v494 = v485
	goto L65
L67:
	;
	goto L68
L68:
	;
	v487 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[3]))
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v487+v35<<(uint(int32(6))%32)+int32(-64))+16))
	v494 = v493
	goto L65
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+840)) = v501
	v505 = F_palloc(m, int32(_a_F__bt_delete_or_dedup_one_page_6))
	mBase = m.M
	v506 = m.ExcPending
	if v506 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+844)) = v505
	v517 = v57
	goto L71
L71:
	;
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v30)+844))
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v30)+836))
	v539 = v535 + v536*int32(6)
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v30)+840))
	v543 = v540 + v536<<(uint(int32(3))%32)
	v548 = v155 + v517&int32(_a_F__bt_delete_or_dedup_one_page_1)<<(uint(int32(2))%32)
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v552 = v141 + v549&int32(_a_F__bt_delete_or_dedup_one_page_3)
	v553 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v552)+7)))
	if v553&int32(32) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	F_pfree(m, v374)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L25
	} else {
		goto L90
	}
L73:
	;
	v706 = v517 + int32(1)
	v707 = int32(_a_F__bt_delete_or_dedup_one_page_1)
	if base.Ui32(v706&v707) <= base.Ui32(v66&v707) {
		v517 = v706
		goto L71
	} else {
		goto L89
	}
L74:
	;
	v595 = v556 & int32(4095)
	if v595 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L75:
	;
	v556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v552)+4)))
	if v556&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v552)+2)))
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v552))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+816)) = v560 | v561<<(uint(int32(16))%32)
	v570 = F_bsearch(m, v30+int32(816), v374, v453, int32(4), int32(208))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L25
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if v570 == int32(0) {
		goto L73
	} else {
		goto L80
	}
L80:
	;
	v574 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v552)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v543)+4)) = uint16(v574)
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v552)))
	*(*int32)(unsafe.Add(mBase, uint32(v543))) = v576
	v578 = *(*int32)(unsafe.Add(mBase, uint32(v30)+836))
	*(*uint16)(unsafe.Add(mBase, uint32(v543)+6)) = uint16(v578)
	*(*uint16)(unsafe.Add(mBase, uint32(v539))) = uint16(v517)
	v581 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v582 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v539)+4)) = uint16(v582)
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+3)) = uint8(v582)
	v586 = int32(_a_F__bt_delete_or_dedup_one_page_2)
	*(*uint8)(unsafe.Add(mBase, uint32(v539)+2)) = uint8(base.B2i32(v581&v586 == v586))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+836)) = v578 + int32(1)
	goto L73
L81:
	;
	v606 = v539
	v609 = int32(0)
	v610 = v543
	goto L82
L82:
	;
	v626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v552)+2)))
	v627 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v552))))
	v628 = int32(16)
	v634 = v626 + (v552 + v627<<(uint(v628)%32)) + v609*int32(6)
	v635 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v634))))
	v638 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v634)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+816)) = v635<<(uint(v628)%32) | v638
	v645 = F_bsearch(m, v30+int32(816), v374, v453, int32(4), int32(208))
	mBase = m.M
	v646 = m.ExcPending
	if v646 != 0 {
		goto L25
	} else {
		goto L84
	}
L83:
	;
	goto L73
L84:
	;
	if v645 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v647 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v634)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v610)+4)) = uint16(v647)
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v634)))
	*(*int32)(unsafe.Add(mBase, uint32(v610))) = v649
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v30)+836))
	*(*uint16)(unsafe.Add(mBase, uint32(v610)+6)) = uint16(v651)
	*(*uint16)(unsafe.Add(mBase, uint32(v606))) = uint16(v517)
	v654 = *(*int32)(unsafe.Add(mBase, uint32(v548)))
	v655 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v606)+4)) = uint16(v655)
	*(*uint8)(unsafe.Add(mBase, uint32(v606)+3)) = uint8(v655)
	v659 = int32(_a_F__bt_delete_or_dedup_one_page_2)
	*(*uint8)(unsafe.Add(mBase, uint32(v606)+2)) = uint8(base.B2i32(v654&v659 == v659))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v30)+836))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+836)) = v664 + int32(1)
	v672 = v606 + int32(6)
	v673 = v610 + int32(8)
	goto L87
L86:
	;
	v672 = v606
	v673 = v610
	goto L87
L87:
	;
	v676 = v609 + int32(1)
	if v676 != v595 {
		v606 = v672
		v609 = v676
		v610 = v673
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
	F__bt_delitems_delete_check(m, l0, v35, l1, v30+int32(820))
	mBase = m.M
	v717 = m.ExcPending
	if v717 != 0 {
		goto L25
	} else {
		goto L91
	}
L91:
	;
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v30)+840))
	F_pfree(m, v718)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L25
	} else {
		goto L92
	}
L92:
	;
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v30)+844))
	F_pfree(m, v721)
	mBase = m.M
	v723 = m.ExcPending
	if v723 != 0 {
		goto L25
	} else {
		goto L93
	}
L93:
	;
	v724 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v724)
	v727 = int32(4)
	v728 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+14)))
	v729 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
	v730 = v728 - v729
	if v730 <= v727 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v736 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v736) <= base.Ui32(v733-int32(4)) {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	v733 = v727
	goto L97
L96:
	;
	v733 = v730
	goto L97
L97:
	;
	goto L94
L98:
	;
	v743 = int32(1)
	goto L2
L99:
	;
	v771 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v771)
	v773 = v743 | l6
	if v773 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v776 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v777 = m.G0
	v779 = v777 - int32(32)
	m.G0 = v779
	if v35 < int32(0) {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	goto L102
L102:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v1746 != 0 {
		goto L222
	} else {
		goto L223
	}
L103:
	;
	v799 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+16)))
	v800 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v801 = int32(*(*int16)(unsafe.Add(mBase, uint32(v800)+10)))
	v803 = F_palloc(m, int32(1676))
	mBase = m.M
	v804 = m.ExcPending
	if v804 != 0 {
		goto L25
	} else {
		goto L107
	}
L104:
	;
	v784 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[0]))
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v784+(v35^int32(-1))<<(uint(int32(2))%32))))
	v798 = v790
	goto L103
L105:
	;
	goto L106
L106:
	;
	v792 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[1]))
	v798 = v792 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L103
L107:
	;
	v805 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v803)+20)) = v805
	*(*uint16)(unsafe.Add(mBase, uint32(v803)+16)) = uint16(v805)
	*(*int32)(unsafe.Add(mBase, uint32(v803)+12)) = v805
	*(*int64)(unsafe.Add(mBase, uint32(v803)+4)) = int64(35184372088832)
	v813 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v803))) = uint8(v813)
	v817 = F_palloc(m, int32(_a_F__bt_delete_or_dedup_one_page_4))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L25
	} else {
		goto L108
	}
L108:
	;
	v819 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v803)+28)) = v819
	*(*int32)(unsafe.Add(mBase, uint32(v803)+24)) = v817
	*(*int64)(unsafe.Add(mBase, uint32(v803)+36)) = v819
	*(*int32)(unsafe.Add(mBase, uint32(v779)+4)) = l0
	if v35 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v779)+20)) = int32(0)
	v846 = int32(512)
	v848 = v776 + int32(4)
	if base.Ui32(v848) <= base.Ui32(v846) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v828 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[2]))
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v828+(v35^int32(-1))<<(uint(int32(6))%32))+16))
	v843 = v834
	goto L109
L111:
	;
	goto L112
L112:
	;
	v836 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[3]))
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v836+v35<<(uint(int32(6))%32)+int32(-64))+16))
	v843 = v842
	goto L109
L113:
	;
	v851 = v846
	goto L115
L114:
	;
	v851 = v848
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v779)+16)) = v851
	v853 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v779)+12)) = uint8(v853)
	*(*int32)(unsafe.Add(mBase, uint32(v779)+8)) = v843
	v857 = F_palloc(m, int32(_a_F__bt_delete_or_dedup_one_page_5))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L25
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v779)+24)) = v857
	v861 = F_palloc(m, int32(_a_F__bt_delete_or_dedup_one_page_6))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L25
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v779)+28)) = v861
	v867 = *(*int32)(unsafe.Add(mBase, uint32(v799+v798)+4))
	if v867 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v868 = int32(2)
	goto L120
L119:
	;
	v868 = int32(1)
	goto L120
L120:
	;
	v869 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v869) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v877 = int32(base.Ui32(v869+int32(_a_F__bt_delete_or_dedup_one_page_0)) >> (uint(int32(2)) % 32))
	goto L123
L122:
	;
	v877 = int32(0)
	goto L123
L123:
	;
	v879 = v877 & int32(_a_F__bt_delete_or_dedup_one_page_1)
	if base.Ui32(v868) <= base.Ui32(v879) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v882 = v803 + int32(44)
	v888 = v868
	goto L127
L125:
	;
	goto L126
L126:
	;
	v1419 = v779 + int32(4)
	v1420 = int32(0)
	v1434 = *(*int32)(unsafe.Add(mBase, uint32(v803)+32))
	if v1434 <= v1420 {
		goto L186
	} else {
		goto L187
	}
L127:
	;
	v913 = v888 & int32(_a_F__bt_delete_or_dedup_one_page_1)
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v798+int32(20)+v913<<(uint(int32(2))%32))))
	v920 = v798 + v917&int32(_a_F__bt_delete_or_dedup_one_page_3)
	if v868 == v913 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L126
L129:
	;
	v1387 = v888 + int32(1)
	if base.Ui32(v1387&int32(_a_F__bt_delete_or_dedup_one_page_1)) <= base.Ui32(v879) {
		v888 = v1387
		goto L127
	} else {
		goto L184
	}
L130:
	;
	v922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+7)))
	if v922&int32(32) != 0 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	goto L132
L132:
	;
	v980 = *(*int32)(unsafe.Add(mBase, uint32(v803)+12))
	v981 = F__bt_keep_natts_fast(m, l0, v980, v920)
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L25
	} else {
		goto L146
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v803)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v803)+20)) = v959
	*(*uint16)(unsafe.Add(mBase, uint32(v803)+16)) = uint16(v868)
	*(*int32)(unsafe.Add(mBase, uint32(v803)+12)) = v920
	v965 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+36)) = (v965&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) | int32(4)
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v803)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v882+v975<<(uint(int32(2))%32)))) = uint16(v868)
	goto L129
L134:
	;
	v940 = v925 & int32(4095)
	v942 = v940 * int32(6)
	if v942 != 0 {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v925 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+4)))
	if v925&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v803)+24))
	v930 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v929)+4)) = uint16(v930)
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v920)))
	*(*int32)(unsafe.Add(mBase, uint32(v929))) = v932
	*(*int32)(unsafe.Add(mBase, uint32(v803)+28)) = int32(1)
	v936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+6)))
	v959 = v936 & int32(_a_F__bt_delete_or_dedup_one_page_7)
	goto L133
L138:
	;
	goto L137
L139:
	;
	v943 = *(*int32)(unsafe.Add(mBase, uint32(v803)+24))
	v944 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+2)))
	v945 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920))))
	base.MemoryCopy(m, v943, v944+(v920+v945<<(uint(int32(16))%32)), v942)
	goto L141
L140:
	;
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v803)+28)) = v940
	v952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+2)))
	v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920))))
	v959 = v952 | v953<<(uint(int32(16))%32)
	goto L133
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v803)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v803)+20)) = v1361
	*(*uint16)(unsafe.Add(mBase, uint32(v803)+16)) = uint16(v888)
	*(*int32)(unsafe.Add(mBase, uint32(v803)+12)) = v920
	v1367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+36)) = (v1367&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) | int32(4)
	v1377 = *(*int32)(unsafe.Add(mBase, uint32(v803)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v882+v1377<<(uint(int32(2))%32)))) = uint16(v888)
	goto L129
L143:
	;
	v1342 = v1301 & int32(4095)
	v1344 = v1342 * int32(6)
	if v1344 != 0 {
		goto L181
	} else {
		goto L182
	}
L144:
	;
	v1315 = *(*int32)(unsafe.Add(mBase, uint32(v803)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+32)) = v1315 + int32(1)
	v1320 = v1003 * int32(6)
	if v1320 != 0 {
		goto L178
	} else {
		goto L179
	}
L145:
	;
	v1028 = v779 + int32(4)
	v1029 = int32(0)
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(v803)+32))
	if v1043 <= v1029 {
		goto L154
	} else {
		goto L155
	}
L146:
	;
	if v981 <= v801 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v984 = int32(1)
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+7)))
	if v985&int32(32) == int32(0) {
		v1003 = v984
		v1005 = v920
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v803)+8))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v803)+20))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v803)+28))
	if base.Ui32((v1007+(v1008+v1003)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1006) {
		goto L144
	} else {
		goto L151
	}
L149:
	;
	v990 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+4)))
	if v990&int32(_a_F__bt_delete_or_dedup_one_page_4) == int32(0) {
		v1003 = v984
		v1005 = v920
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v997 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+2)))
	v998 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920))))
	v1003 = v990 & int32(4095)
	v1005 = v997 + (v920 + v998<<(uint(int32(16))%32))
	goto L148
L151:
	;
	if v1008 <= int32(50) {
		goto L145
	} else {
		goto L152
	}
L152:
	;
	v1020 = *(*int32)(unsafe.Add(mBase, uint32(v803)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+4)) = v1020 + int32(1)
	goto L145
L153:
	;
	v1298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v920)+7)))
	if v1298&int32(32) != 0 {
		goto L174
	} else {
		goto L175
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v803)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v803)+28)) = int64(0)
	goto L153
L155:
	;
	v1050 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+16))
	v1058 = v1050
	v1064 = v1029
	goto L156
L156:
	;
	v1068 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+24))
	v1071 = v1068 + v1058*int32(6)
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+20))
	v1075 = v1072 + v1058<<(uint(int32(3))%32)
	v1076 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v803)+16)))
	v1077 = v1076 + v1064
	v1082 = v798 + int32(20) + v1077&int32(_a_F__bt_delete_or_dedup_one_page_1)<<(uint(int32(2))%32)
	v1083 = *(*int32)(unsafe.Add(mBase, uint32(v1082)))
	v1086 = v798 + v1083&int32(_a_F__bt_delete_or_dedup_one_page_3)
	v1087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1086)+7)))
	if v1087&int32(32) != 0 {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	if v1043 <= int32(1) {
		goto L154
	} else {
		goto L173
	}
L158:
	;
	v1263 = v1064 + int32(1)
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v803)+32))
	if v1263 < v1264 {
		v1058 = v1252
		v1064 = v1263
		goto L156
	} else {
		goto L172
	}
L159:
	;
	v1115 = v1090 & int32(4095)
	v1116 = int32(0)
	if int32(2) <= v1043 {
		goto L164
	} else {
		goto L165
	}
L160:
	;
	v1090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1086)+4)))
	if v1090&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v1094 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1086)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)) = uint16(v1094)
	v1096 = *(*int32)(unsafe.Add(mBase, uint32(v1086)))
	*(*int32)(unsafe.Add(mBase, uint32(v1075))) = v1096
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1075)+6)) = uint16(v1098)
	*(*uint8)(unsafe.Add(mBase, uint32(v1071)+3)) = uint8(base.B2i32(int32(1) < v1043))
	v1101 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1071)+2)) = uint8(v1101)
	*(*uint16)(unsafe.Add(mBase, uint32(v1071))) = uint16(v1077)
	v1104 = *(*int32)(unsafe.Add(mBase, uint32(v1082)))
	v1108 = int32(base.Ui32(v1104)>>(uint(int32(17))%32)) + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1071)+4)) = uint16(v1108)
	v1110 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+16))
	v1112 = v1110 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1028)+16)) = v1112
	v1252 = v1112
	goto L158
L163:
	;
	goto L162
L164:
	;
	v1120 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1086)+2)))
	v1121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1086))))
	v1122 = int32(16)
	v1125 = v1120 + (v1086 + v1121<<(uint(v1122)%32))
	v1126 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1125))))
	v1129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1125)+2)))
	v1130 = v1126<<(uint(v1122)%32) | v1129
	v1133 = int32(6)
	v1135 = v1125 + int32(base.Ui32(v1115)>>(uint(int32(1))%32))*v1133
	v1136 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1135))))
	v1139 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1135)+2)))
	v1140 = v1136<<(uint(v1122)%32) | v1139
	v1144 = v1125 + v1115*v1133
	v1147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1144-v1133))))
	v1152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1144-int32(4)))))
	v1157 = base.B2i32(v1140 == v1130)
	v1158 = base.B2i32(v1130 != v1140) & base.B2i32(v1140 == v1147<<(uint(v1122)%32)|v1152)
	goto L166
L165:
	;
	v1157 = v1116
	v1158 = v1116
	goto L166
L166:
	;
	if v1115 == int32(0) {
		v1252 = v1058
		goto L158
	} else {
		goto L167
	}
L167:
	;
	v1163 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1086)+2)))
	v1164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1086))))
	v1168 = v1163 + (v1086 + v1164<<(uint(int32(16))%32))
	v1169 = *(*int32)(unsafe.Add(mBase, uint32(v1168)))
	*(*int32)(unsafe.Add(mBase, uint32(v1075))) = v1169
	v1171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1168)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1075)+4)) = uint16(v1171)
	v1173 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1075)+6)) = uint16(v1173)
	v1175 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v1071)+4)) = uint16(v1175)
	v1177 = int32(1)
	v1178 = v1115 - v1177
	v1179 = int32(0)
	v1182 = v1157 | v1158&base.B2i32(v1178 == v1179)
	*(*uint8)(unsafe.Add(mBase, uint32(v1071)+3)) = uint8(v1182)
	*(*uint8)(unsafe.Add(mBase, uint32(v1071)+2)) = uint8(v1179)
	*(*uint16)(unsafe.Add(mBase, uint32(v1071))) = uint16(v1077)
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+16))
	v1190 = v1188 + v1177
	*(*int32)(unsafe.Add(mBase, uint32(v1028)+16)) = v1190
	if v1115 == v1177 {
		v1252 = v1190
		goto L158
	} else {
		goto L168
	}
L168:
	;
	v1197 = v1071
	v1198 = v1177
	v1200 = v1075
	goto L169
L169:
	;
	v1211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1086)+2)))
	v1212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1086))))
	v1217 = int32(6)
	v1219 = v1211 + (v1086 + v1212<<(uint(int32(16))%32)) + v1198*v1217
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1219)))
	*(*int32)(unsafe.Add(mBase, uint32(v1200)+8)) = v1220
	v1222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1219)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1200)+12)) = uint16(v1222)
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1200)+14)) = uint16(v1224)
	*(*uint16)(unsafe.Add(mBase, uint32(v1197)+10)) = uint16(v1217)
	v1229 = v1158 & base.B2i32(v1198 == v1178)
	*(*uint8)(unsafe.Add(mBase, uint32(v1197)+9)) = uint8(v1229)
	v1231 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1197)+8)) = uint8(v1231)
	*(*uint16)(unsafe.Add(mBase, uint32(v1197)+6)) = uint16(v1077)
	v1234 = *(*int32)(unsafe.Add(mBase, uint32(v1028)+16))
	v1235 = int32(1)
	v1236 = v1234 + v1235
	*(*int32)(unsafe.Add(mBase, uint32(v1028)+16)) = v1236
	v1243 = v1198 + v1235
	if v1243 != v1115 {
		v1197 = v1197 + v1217
		v1198 = v1243
		v1200 = v1200 + int32(8)
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v1252 = v1236
	goto L158
L171:
	;
	goto L170
L172:
	;
	goto L157
L173:
	;
	v1268 = *(*int32)(unsafe.Add(mBase, uint32(v803)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v803+v1268<<(uint(int32(2))%32))+46)) = uint16(v1264)
	v1273 = *(*int32)(unsafe.Add(mBase, uint32(v803)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+40)) = v1273 + int32(1)
	goto L154
L174:
	;
	v1301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+4)))
	if v1301&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L143
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v1305 = *(*int32)(unsafe.Add(mBase, uint32(v803)+24))
	v1306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1305)+4)) = uint16(v1306)
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v920)))
	*(*int32)(unsafe.Add(mBase, uint32(v1305))) = v1308
	*(*int32)(unsafe.Add(mBase, uint32(v803)+28)) = int32(1)
	v1312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+6)))
	v1361 = v1312 & int32(_a_F__bt_delete_or_dedup_one_page_7)
	goto L142
L177:
	;
	goto L176
L178:
	;
	v1321 = *(*int32)(unsafe.Add(mBase, uint32(v803)+24))
	base.MemoryCopy(m, v1321+v1008*int32(6), v1005, v1320)
	goto L180
L179:
	;
	goto L180
L180:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v803)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+28)) = v1326 + v1003
	v1329 = *(*int32)(unsafe.Add(mBase, uint32(v803)+36))
	v1330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+36)) = v1329 + (v1330&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) + int32(4)
	goto L129
L181:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(v803)+24))
	v1346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+2)))
	v1347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920))))
	base.MemoryCopy(m, v1345, v1346+(v920+v1347<<(uint(int32(16))%32)), v1344)
	goto L183
L182:
	;
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v803)+28)) = v1342
	v1354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920)+2)))
	v1355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v920))))
	v1361 = v1354 | v1355<<(uint(int32(16))%32)
	goto L142
L184:
	;
	goto L128
L185:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v803)+40))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v803)+24))
	F_pfree(m, v1690)
	mBase = m.M
	v1692 = m.ExcPending
	if v1692 != 0 {
		goto L25
	} else {
		goto L206
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v803)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v803)+28)) = int64(0)
	goto L185
L187:
	;
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+16))
	v1449 = v1441
	v1455 = v1420
	goto L188
L188:
	;
	v1459 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+24))
	v1462 = v1459 + v1449*int32(6)
	v1463 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+20))
	v1466 = v1463 + v1449<<(uint(int32(3))%32)
	v1467 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v803)+16)))
	v1468 = v1467 + v1455
	v1473 = v798 + int32(20) + v1468&int32(_a_F__bt_delete_or_dedup_one_page_1)<<(uint(int32(2))%32)
	v1474 = *(*int32)(unsafe.Add(mBase, uint32(v1473)))
	v1477 = v798 + v1474&int32(_a_F__bt_delete_or_dedup_one_page_3)
	v1478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1477)+7)))
	if v1478&int32(32) != 0 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	if v1434 <= int32(1) {
		goto L186
	} else {
		goto L205
	}
L190:
	;
	v1654 = v1455 + int32(1)
	v1655 = *(*int32)(unsafe.Add(mBase, uint32(v803)+32))
	if v1654 < v1655 {
		v1449 = v1643
		v1455 = v1654
		goto L188
	} else {
		goto L204
	}
L191:
	;
	v1506 = v1481 & int32(4095)
	v1507 = int32(0)
	if int32(2) <= v1434 {
		goto L196
	} else {
		goto L197
	}
L192:
	;
	v1481 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1477)+4)))
	if v1481&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1485 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1477)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1466)+4)) = uint16(v1485)
	v1487 = *(*int32)(unsafe.Add(mBase, uint32(v1477)))
	*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1487
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1466)+6)) = uint16(v1489)
	*(*uint8)(unsafe.Add(mBase, uint32(v1462)+3)) = uint8(base.B2i32(int32(1) < v1434))
	v1492 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1462)+2)) = uint8(v1492)
	*(*uint16)(unsafe.Add(mBase, uint32(v1462))) = uint16(v1468)
	v1495 = *(*int32)(unsafe.Add(mBase, uint32(v1473)))
	v1499 = int32(base.Ui32(v1495)>>(uint(int32(17))%32)) + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1462)+4)) = uint16(v1499)
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+16))
	v1503 = v1501 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1419)+16)) = v1503
	v1643 = v1503
	goto L190
L195:
	;
	goto L194
L196:
	;
	v1511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1477)+2)))
	v1512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1477))))
	v1513 = int32(16)
	v1516 = v1511 + (v1477 + v1512<<(uint(v1513)%32))
	v1517 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1516))))
	v1520 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1516)+2)))
	v1521 = v1517<<(uint(v1513)%32) | v1520
	v1524 = int32(6)
	v1526 = v1516 + int32(base.Ui32(v1506)>>(uint(int32(1))%32))*v1524
	v1527 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1526))))
	v1530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1526)+2)))
	v1531 = v1527<<(uint(v1513)%32) | v1530
	v1535 = v1516 + v1506*v1524
	v1538 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1535-v1524))))
	v1543 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1535-int32(4)))))
	v1548 = base.B2i32(v1531 == v1521)
	v1549 = base.B2i32(v1521 != v1531) & base.B2i32(v1531 == v1538<<(uint(v1513)%32)|v1543)
	goto L198
L197:
	;
	v1548 = v1507
	v1549 = v1507
	goto L198
L198:
	;
	if v1506 == int32(0) {
		v1643 = v1449
		goto L190
	} else {
		goto L199
	}
L199:
	;
	v1554 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1477)+2)))
	v1555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1477))))
	v1559 = v1554 + (v1477 + v1555<<(uint(int32(16))%32))
	v1560 = *(*int32)(unsafe.Add(mBase, uint32(v1559)))
	*(*int32)(unsafe.Add(mBase, uint32(v1466))) = v1560
	v1562 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1559)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1466)+4)) = uint16(v1562)
	v1564 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1466)+6)) = uint16(v1564)
	v1566 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v1462)+4)) = uint16(v1566)
	v1568 = int32(1)
	v1569 = v1506 - v1568
	v1570 = int32(0)
	v1573 = v1548 | v1549&base.B2i32(v1569 == v1570)
	*(*uint8)(unsafe.Add(mBase, uint32(v1462)+3)) = uint8(v1573)
	*(*uint8)(unsafe.Add(mBase, uint32(v1462)+2)) = uint8(v1570)
	*(*uint16)(unsafe.Add(mBase, uint32(v1462))) = uint16(v1468)
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+16))
	v1581 = v1579 + v1568
	*(*int32)(unsafe.Add(mBase, uint32(v1419)+16)) = v1581
	if v1506 == v1568 {
		v1643 = v1581
		goto L190
	} else {
		goto L200
	}
L200:
	;
	v1588 = v1462
	v1589 = v1568
	v1591 = v1466
	goto L201
L201:
	;
	v1602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1477)+2)))
	v1603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1477))))
	v1608 = int32(6)
	v1610 = v1602 + (v1477 + v1603<<(uint(int32(16))%32)) + v1589*v1608
	v1611 = *(*int32)(unsafe.Add(mBase, uint32(v1610)))
	*(*int32)(unsafe.Add(mBase, uint32(v1591)+8)) = v1611
	v1613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1610)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1591)+12)) = uint16(v1613)
	v1615 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1591)+14)) = uint16(v1615)
	*(*uint16)(unsafe.Add(mBase, uint32(v1588)+10)) = uint16(v1608)
	v1620 = v1549 & base.B2i32(v1589 == v1569)
	*(*uint8)(unsafe.Add(mBase, uint32(v1588)+9)) = uint8(v1620)
	v1622 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1588)+8)) = uint8(v1622)
	*(*uint16)(unsafe.Add(mBase, uint32(v1588)+6)) = uint16(v1468)
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+16))
	v1626 = int32(1)
	v1627 = v1625 + v1626
	*(*int32)(unsafe.Add(mBase, uint32(v1419)+16)) = v1627
	v1634 = v1589 + v1626
	if v1634 != v1506 {
		v1588 = v1588 + v1608
		v1589 = v1634
		v1591 = v1591 + int32(8)
		goto L201
	} else {
		goto L203
	}
L202:
	;
	v1643 = v1627
	goto L190
L203:
	;
	goto L202
L204:
	;
	goto L189
L205:
	;
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v803)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v803+v1659<<(uint(int32(2))%32))+46)) = uint16(v1655)
	v1664 = *(*int32)(unsafe.Add(mBase, uint32(v803)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v803)+40)) = v1664 + int32(1)
	goto L186
L206:
	;
	F_pfree(m, v803)
	mBase = m.M
	v1694 = m.ExcPending
	if v1694 != 0 {
		goto L25
	} else {
		goto L207
	}
L207:
	;
	F__bt_delitems_delete_check(m, l0, v35, l1, v1419)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L25
	} else {
		goto L208
	}
L208:
	;
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v779)+24))
	F_pfree(m, v1697)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L25
	} else {
		goto L209
	}
L209:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v779)+28))
	F_pfree(m, v1700)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L25
	} else {
		goto L210
	}
L210:
	;
	if v1689 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1703 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+14)))
	v1704 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v798)+12)))
	v1705 = v1703 - v1704
	v1706 = int32(0)
	if v1706 < v1705 {
		goto L215
	} else {
		goto L216
	}
L212:
	;
	v1715 = v813
	goto L213
L213:
	;
	m.G0 = v779 + int32(32)
	if v1715 != 0 {
		goto L1
	} else {
		goto L221
	}
L214:
	;
	v1710 = int32(341)
	if base.Ui32(v848) <= base.Ui32(v1710) {
		goto L218
	} else {
		goto L219
	}
L215:
	;
	v1709 = v1705
	goto L217
L216:
	;
	v1709 = v1706
	goto L217
L217:
	;
	goto L214
L218:
	;
	v1713 = v1710
	goto L220
L219:
	;
	v1713 = v848
	goto L220
L220:
	;
	v1715 = base.B2i32(base.Ui32(v1713) <= base.Ui32(v1709))
	goto L213
L221:
	;
	goto L102
L222:
	;
	v1747 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1746)+16)))
	if v1747 == int32(0) {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1750 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v1750 != int32(1) {
		goto L1
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	v1753 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1755 = int32(0)
	v1756 = m.G0
	v1758 = v1756 - int32(16)
	m.G0 = v1758
	v1760 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v1761 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1760)+10)))
	if v35 < v1755 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v1780 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1779)+16)))
	v1782 = F_palloc(m, int32(1676))
	mBase = m.M
	v1783 = m.ExcPending
	if v1783 != 0 {
		goto L25
	} else {
		goto L231
	}
L228:
	;
	v1765 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[0]))
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v1765+(v35^int32(-1))<<(uint(int32(2))%32))))
	v1779 = v1771
	goto L227
L229:
	;
	goto L230
L230:
	;
	v1773 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[1]))
	v1779 = v1773 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L227
L231:
	;
	v1784 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+20)) = v1784
	*(*uint16)(unsafe.Add(mBase, uint32(v1782)+16)) = uint16(v1784)
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+12)) = v1784
	*(*int64)(unsafe.Add(mBase, uint32(v1782)+4)) = int64(5806795784192)
	v1792 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1782))) = uint8(v1792)
	v1795 = F_palloc(m, int32(1352))
	mBase = m.M
	v1796 = m.ExcPending
	if v1796 != 0 {
		goto L25
	} else {
		goto L232
	}
L232:
	;
	v1797 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1782)+28)) = v1797
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+24)) = v1795
	*(*int64)(unsafe.Add(mBase, uint32(v1782)+36)) = v1797
	v1802 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1779)+12)))
	v1805 = v1780 + v1779
	v1806 = *(*int32)(unsafe.Add(mBase, uint32(v1805)+4))
	if v1806 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1807 = int32(2)
	goto L235
L234:
	;
	v1807 = int32(1)
	goto L235
L235:
	;
	if v773 != 0 {
		v1847 = v1755
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1851 = F_PageGetTempPageCopySpecial(m, v1779)
	mBase = m.M
	v1852 = m.ExcPending
	if v1852 != 0 {
		goto L25
	} else {
		goto L247
	}
L237:
	;
	v1808 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v1809 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1808)+10)))
	v1811 = v1779 + int32(20)
	v1815 = *(*int32)(unsafe.Add(mBase, uint32(v1811+v1807<<(uint(int32(2))%32))))
	v1819 = F__bt_keep_natts_fast(m, l0, v1753, v1779+v1815&int32(_a_F__bt_delete_or_dedup_one_page_3))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L25
	} else {
		goto L238
	}
L238:
	;
	if v1809 < v1819 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1823 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1779)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1823) {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	goto L241
L241:
	;
	v1847 = int32(0)
	goto L236
L242:
	;
	v1833 = int32(base.Ui32(v1823+int32(_a_F__bt_delete_or_dedup_one_page_0))>>(uint(int32(2))%32)) & int32(_a_F__bt_delete_or_dedup_one_page_1)
	goto L244
L243:
	;
	v1833 = int32(0)
	goto L244
L244:
	;
	v1837 = *(*int32)(unsafe.Add(mBase, uint32(v1811+v1833<<(uint(int32(2))%32))))
	v1841 = F__bt_keep_natts_fast(m, l0, v1753, v1779+v1837&int32(_a_F__bt_delete_or_dedup_one_page_3))
	mBase = m.M
	v1842 = m.ExcPending
	if v1842 != 0 {
		goto L25
	} else {
		goto L245
	}
L245:
	;
	if v1809 < v1841 {
		v1847 = int32(1)
		goto L236
	} else {
		goto L246
	}
L246:
	;
	goto L241
L247:
	;
	v1853 = *(*int64)(unsafe.Add(mBase, uint32(v1779)))
	*(*int64)(unsafe.Add(mBase, uint32(v1851))) = v1853
	v1855 = *(*int32)(unsafe.Add(mBase, uint32(v1805)+4))
	if v1855 != 0 {
		goto L250
	} else {
		goto L251
	}
L248:
	;
	goto L1
L249:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2264 = m.ExcPending
	if v2264 != 0 {
		goto L25
	} else {
		goto L332
	}
L250:
	;
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1779)+24))
	v1864 = F_PageAddItemExtended(m, v1851, v1779+v1856&int32(_a_F__bt_delete_or_dedup_one_page_3), int32(base.Ui32(v1856)>>(uint(int32(17))%32)), int32(1), int32(0))
	mBase = m.M
	v1865 = m.ExcPending
	if v1865 != 0 {
		goto L25
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1802) {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	if v1864 == int32(0) {
		goto L249
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1876 = int32(base.Ui32(v1802+int32(_a_F__bt_delete_or_dedup_one_page_0)) >> (uint(int32(2)) % 32))
	goto L257
L256:
	;
	v1876 = int32(0)
	goto L257
L257:
	;
	v1878 = v1876 & int32(_a_F__bt_delete_or_dedup_one_page_1)
	if base.Ui32(v1807) <= base.Ui32(v1878) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1881 = v1782 + int32(44)
	v1885 = v1847
	v1891 = v1807
	goto L261
L259:
	;
	goto L260
L260:
	;
	F__bt_dedup_finish_pending(m, v1851, v1782)
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L25
	} else {
		goto L307
	}
L261:
	;
	v1912 = v1891 & int32(_a_F__bt_delete_or_dedup_one_page_1)
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v1779+int32(20)+v1912<<(uint(int32(2))%32))))
	v1919 = v1779 + v1916&int32(_a_F__bt_delete_or_dedup_one_page_3)
	if v1807 == v1912 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L260
L263:
	;
	v2150 = v1891 + int32(1)
	if base.Ui32(v2150&int32(_a_F__bt_delete_or_dedup_one_page_1)) <= base.Ui32(v1878) {
		v1885 = v2144
		v1891 = v2150
		goto L261
	} else {
		goto L306
	}
L264:
	;
	v1921 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1919)+7)))
	if v1921&int32(32) != 0 {
		goto L269
	} else {
		goto L270
	}
L265:
	;
	goto L266
L266:
	;
	v1979 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1782))))
	if v1979 != int32(1) {
		goto L280
	} else {
		goto L281
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+20)) = v1958
	*(*uint16)(unsafe.Add(mBase, uint32(v1782)+16)) = uint16(v1807)
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+12)) = v1919
	v1964 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+36)) = (v1964&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) | int32(4)
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1881+v1974<<(uint(int32(2))%32)))) = uint16(v1807)
	v2144 = v1885
	goto L263
L268:
	;
	v1939 = v1924 & int32(4095)
	v1941 = v1939 * int32(6)
	if v1941 != 0 {
		goto L273
	} else {
		goto L274
	}
L269:
	;
	v1924 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+4)))
	if v1924&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L268
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1928 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+24))
	v1929 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1928)+4)) = uint16(v1929)
	v1931 = *(*int32)(unsafe.Add(mBase, uint32(v1919)))
	*(*int32)(unsafe.Add(mBase, uint32(v1928))) = v1931
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+28)) = int32(1)
	v1935 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+6)))
	v1958 = v1935 & int32(_a_F__bt_delete_or_dedup_one_page_7)
	goto L267
L272:
	;
	goto L271
L273:
	;
	v1942 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+24))
	v1943 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+2)))
	v1944 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919))))
	base.MemoryCopy(m, v1942, v1943+(v1919+v1944<<(uint(int32(16))%32)), v1941)
	goto L275
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+28)) = v1939
	v1951 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+2)))
	v1952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919))))
	v1958 = v1951 | v1952<<(uint(int32(16))%32)
	goto L267
L276:
	;
	v2086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1919)+7)))
	if v2086&int32(32) != 0 {
		goto L299
	} else {
		goto L300
	}
L277:
	;
	v2080 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1782))) = uint8(v2080)
	v2083 = v2080
	goto L276
L278:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+8))
	v2065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1779)+19)))
	v2075 = v2064 - base.I32_trunc_sat_f64_s(base.F64_mul(base.F64_convert_i32_u(v2065<<(uint(int32(8))%32)-v1754-int32(52)), float64(0.04)))
	if base.Ui32(v2075) <= base.Ui32(v2064) {
		goto L294
	} else {
		goto L295
	}
L279:
	;
	v2038 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+32)) = v2038 + int32(1)
	v2043 = v2005 * int32(6)
	if v2043 != 0 {
		goto L291
	} else {
		goto L292
	}
L280:
	;
	F__bt_dedup_finish_pending(m, v1851, v1782)
	mBase = m.M
	v2030 = m.ExcPending
	if v2030 != 0 {
		goto L25
	} else {
		goto L289
	}
L281:
	;
	v1982 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+12))
	v1983 = F__bt_keep_natts_fast(m, l0, v1982, v1919)
	mBase = m.M
	v1984 = m.ExcPending
	if v1984 != 0 {
		goto L25
	} else {
		goto L282
	}
L282:
	;
	if v1983 <= v1761 {
		goto L280
	} else {
		goto L283
	}
L283:
	;
	v1986 = int32(1)
	v1987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1919)+7)))
	if v1987&int32(32) == int32(0) {
		v2005 = v1986
		v2007 = v1919
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+8))
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+20))
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+28))
	if base.Ui32((v2009+(v2010+v2005)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v2008) {
		goto L279
	} else {
		goto L287
	}
L285:
	;
	v1992 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+4)))
	if v1992&int32(_a_F__bt_delete_or_dedup_one_page_4) == int32(0) {
		v2005 = v1986
		v2007 = v1919
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v1999 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+2)))
	v2000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919))))
	v2005 = v1992 & int32(4095)
	v2007 = v1999 + (v1919 + v2000<<(uint(int32(16))%32))
	goto L284
L287:
	;
	if v2010 <= int32(50) {
		goto L280
	} else {
		goto L288
	}
L288:
	;
	v2022 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+4)) = v2022 + int32(1)
	goto L280
L289:
	;
	v2031 = int32(0)
	if v1885 == v2031 {
		v2083 = v2031
		goto L276
	} else {
		goto L290
	}
L290:
	;
	v2034 = int32(1)
	v2035 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+4))
	switch v2035 - int32(5) {
	case 0:
		goto L278
	case 1:
		goto L277
	default:
		v2083 = v2034
		goto L276
	}
L291:
	;
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+24))
	base.MemoryCopy(m, v2044+v2010*int32(6), v2007, v2043)
	goto L293
L292:
	;
	goto L293
L293:
	;
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+28)) = v2049 + v2005
	v2052 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+36))
	v2053 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+36)) = v2052 + (v2053&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) + int32(4)
	v2144 = v1885
	goto L263
L294:
	;
	v2078 = v2075
	goto L296
L295:
	;
	v2078 = int32(0)
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+8)) = v2078
	v2083 = v2034
	goto L276
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+20)) = v2123
	*(*uint16)(unsafe.Add(mBase, uint32(v1782)+16)) = uint16(v1891)
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+12)) = v1919
	v2129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+36)) = (v2129&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) | int32(4)
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1881+v2139<<(uint(int32(2))%32)))) = uint16(v1891)
	v2144 = v2083
	goto L263
L298:
	;
	v2104 = v2089 & int32(4095)
	v2106 = v2104 * int32(6)
	if v2106 != 0 {
		goto L303
	} else {
		goto L304
	}
L299:
	;
	v2089 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+4)))
	if v2089&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L298
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v2093 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+24))
	v2094 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2093)+4)) = uint16(v2094)
	v2096 = *(*int32)(unsafe.Add(mBase, uint32(v1919)))
	*(*int32)(unsafe.Add(mBase, uint32(v2093))) = v2096
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+28)) = int32(1)
	v2100 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+6)))
	v2123 = v2100 & int32(_a_F__bt_delete_or_dedup_one_page_7)
	goto L297
L302:
	;
	goto L301
L303:
	;
	v2107 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+24))
	v2108 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+2)))
	v2109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919))))
	base.MemoryCopy(m, v2107, v2108+(v1919+v2109<<(uint(int32(16))%32)), v2106)
	goto L305
L304:
	;
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1782)+28)) = v2104
	v2116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919)+2)))
	v2117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1919))))
	v2123 = v2116 | v2117<<(uint(int32(16))%32)
	goto L297
L306:
	;
	goto L262
L307:
	;
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+40))
	if v2183 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	v2253 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+24))
	F_pfree(m, v2253)
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L25
	} else {
		goto L330
	}
L309:
	;
	F_pfree(m, v1851)
	mBase = m.M
	v2187 = m.ExcPending
	if v2187 != 0 {
		goto L25
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v2188 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1805)+12)))
	if v2188&int32(64) != 0 {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	goto L308
L313:
	;
	v2191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1851)+16)))
	v2192 = v1851 + v2191
	v2193 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2192)+12)))
	v2195 = v2193 & int32(_a_F__bt_delete_or_dedup_one_page_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v2192)+12)) = uint16(v2195)
	goto L315
L314:
	;
	goto L315
L315:
	;
	v2198 = int32(_a_F__bt_delete_or_dedup_one_page_10)
	v2200 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[4])) = v2200 + int32(1)
	F_PageRestoreTempPage(m, v1851, v1779)
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L25
	} else {
		goto L316
	}
L316:
	;
	F_MarkBufferDirty(m, v35)
	mBase = m.M
	v2207 = m.ExcPending
	if v2207 != 0 {
		goto L25
	} else {
		goto L317
	}
L317:
	;
	v2208 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2208)+118)))
	if v2209 != int32(112) {
		goto L318
	} else {
		goto L319
	}
L318:
	;
	v2246 = int32(_a_F__bt_delete_or_dedup_one_page_10)
	v2248 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[4])) = v2248 - int32(1)
	goto L308
L319:
	;
	v2213 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[5]))
	if v2213 <= int32(0) {
		goto L320
	} else {
		goto L321
	}
L320:
	;
	v2216 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2216 != 0 {
		goto L318
	} else {
		goto L323
	}
L321:
	;
	goto L322
L322:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1758)+14)) = uint16(v2218)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2221 = m.ExcPending
	if v2221 != 0 {
		goto L25
	} else {
		goto L325
	}
L323:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2217 != 0 {
		goto L318
	} else {
		goto L324
	}
L324:
	;
	goto L322
L325:
	;
	F_XLogRegisterBuffer(m, int32(0), v35, int32(8))
	mBase = m.M
	v2225 = m.ExcPending
	if v2225 != 0 {
		goto L25
	} else {
		goto L326
	}
L326:
	;
	F_XLogRegisterData(m, v1758+int32(14), int32(2))
	mBase = m.M
	v2230 = m.ExcPending
	if v2230 != 0 {
		goto L25
	} else {
		goto L327
	}
L327:
	;
	v2234 = *(*int32)(unsafe.Add(mBase, uint32(v1782)+40))
	F_XLogRegisterBufData(m, int32(0), v1782+int32(44), v2234<<(uint(int32(2))%32))
	mBase = m.M
	v2238 = m.ExcPending
	if v2238 != 0 {
		goto L25
	} else {
		goto L328
	}
L328:
	;
	v2241 = F_XLogInsert(m, int32(11), int32(96))
	mBase = m.M
	v2242 = m.ExcPending
	if v2242 != 0 {
		goto L25
	} else {
		goto L329
	}
L329:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1779))) = base.I64_rotr(v2241, int64(32))
	goto L318
L330:
	;
	F_pfree(m, v1782)
	mBase = m.M
	v2257 = m.ExcPending
	if v2257 != 0 {
		goto L25
	} else {
		goto L331
	}
L331:
	;
	m.G0 = v1758 + int32(16)
	goto L248
L332:
	;
	F_errmsg_internal(m, int32(_a_F__bt_delete_or_dedup_one_page_11), int32(0))
	mBase = m.M
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L25
	} else {
		goto L333
	}
L333:
	;
	F_errfinish(m, int32(_a_F__bt_delete_or_dedup_one_page_12), int32(130), int32(_a_F__bt_delete_or_dedup_one_page_13))
	mBase = m.M
	v2273 = m.ExcPending
	if v2273 != 0 {
		goto L25
	} else {
		goto L334
	}
L334:
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
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
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
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
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
	var v104 int32
	_ = v104
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
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
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
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v326 int32
	_ = v326
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
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v437 int32
	_ = v437
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v463 int32
	_ = v463
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v510 int32
	_ = v510
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v532 int32
	_ = v532
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
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
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v578 int32
	_ = v578
	var v579 int32
	_ = v579
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v594 int32
	_ = v594
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v621 int32
	_ = v621
	var v626 int32
	_ = v626
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v637 int32
	_ = v637
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v646 int32
	_ = v646
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v728 int32
	_ = v728
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v742 int32
	_ = v742
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v797 int32
	_ = v797
	var v800 int32
	_ = v800
	var v850 int32
	_ = v850
	var v851 int32
	_ = v851
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v909 int32
	_ = v909
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
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v936 int32
	_ = v936
	var v938 int32
	_ = v938
	var v943 int32
	_ = v943
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v961 int32
	_ = v961
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v970 int32
	_ = v970
	var v975 int32
	_ = v975
	var v982 int32
	_ = v982
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v1004 int32
	_ = v1004
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1011 int32
	_ = v1011
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1029 int32
	_ = v1029
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1040 float64
	_ = v1040
	var v1052 int32
	_ = v1052
	var v1069 int32
	_ = v1069
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1104 int32
	_ = v1104
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1124 int32
	_ = v1124
	var v1129 int32
	_ = v1129
	var v1151 int32
	_ = v1151
	var v1197 int32
	_ = v1197
	var v1227 float64
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1264 int32
	_ = v1264
	var v1293 int32
	_ = v1293
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1301 int32
	_ = v1301
	var v1306 int32
	_ = v1306
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
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1341 int32
	_ = v1341
	var v1369 int32
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1373 int32
	_ = v1373
	var v1378 float64
	_ = v1378
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1385 int32
	_ = v1385
	var v1409 int32
	_ = v1409
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1449 int32
	_ = v1449
	var v1468 int32
	_ = v1468
	var v1497 int32
	_ = v1497
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1509 int32
	_ = v1509
	var v1524 int32
	_ = v1524
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1557 int32
	_ = v1557
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1564 int32
	_ = v1564
	var v1565 int32
	_ = v1565
	var v1577 int32
	_ = v1577
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1588 int32
	_ = v1588
	var v1589 int32
	_ = v1589
	var v1594 int32
	_ = v1594
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1610 int32
	_ = v1610
	var v1618 int32
	_ = v1618
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1641 int32
	_ = v1641
	var v1645 int32
	_ = v1645
	var v1651 int32
	_ = v1651
	var v1652 int32
	_ = v1652
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1677 int32
	_ = v1677
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1686 int32
	_ = v1686
	var v1689 int32
	_ = v1689
	var v1690 int32
	_ = v1690
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1696 int32
	_ = v1696
	var v1698 int32
	_ = v1698
	var v1719 int32
	_ = v1719
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1752 int32
	_ = v1752
	var v1756 int32
	_ = v1756
	var v1761 int32
	_ = v1761
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1769 int32
	_ = v1769
	var v1770 int32
	_ = v1770
	var v1772 int32
	_ = v1772
	var v1788 int32
	_ = v1788
	var v1793 int32
	_ = v1793
	var v1821 int32
	_ = v1821
	var v1826 int32
	_ = v1826
	var v1841 int32
	_ = v1841
	var v1843 int32
	_ = v1843
	var v1846 int32
	_ = v1846
	var v1848 int32
	_ = v1848
	var v1872 int32
	_ = v1872
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1884 int32
	_ = v1884
	var v1887 int32
	_ = v1887
	var v1899 int32
	_ = v1899
	var v1902 int32
	_ = v1902
	var v1904 int32
	_ = v1904
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1938 int32
	_ = v1938
	var v1949 int32
	_ = v1949
	var v1968 int32
	_ = v1968
	var v1979 int32
	_ = v1979
	var v1982 int32
	_ = v1982
	var v1984 int32
	_ = v1984
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1995 int32
	_ = v1995
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2000 int32
	_ = v2000
	var v2002 int32
	_ = v2002
	var v2004 int32
	_ = v2004
	var v2021 int32
	_ = v2021
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2069 int32
	_ = v2069
	var v2070 int32
	_ = v2070
	var v2072 int32
	_ = v2072
	var v2074 int32
	_ = v2074
	var v2089 int32
	_ = v2089
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2132 int32
	_ = v2132
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2140 int32
	_ = v2140
	var v2142 int64
	_ = v2142
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2152 int32
	_ = v2152
	var v2160 int32
	_ = v2160
	var v2167 int32
	_ = v2167
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2178 int32
	_ = v2178
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	var v2190 int32
	_ = v2190
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2195 int32
	_ = v2195
	var v2199 int32
	_ = v2199
	var v2200 int32
	_ = v2200
	var v2203 int32
	_ = v2203
	var v2204 int32
	_ = v2204
	var v2207 int32
	_ = v2207
	var v2208 int32
	_ = v2208
	var v2212 int32
	_ = v2212
	var v2218 int32
	_ = v2218
	var v2220 int32
	_ = v2220
	var v2226 int32
	_ = v2226
	var v2230 int32
	_ = v2230
	var v2236 int32
	_ = v2236
	var v2238 int32
	_ = v2238
	var v2244 int32
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int32
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2254 int32
	_ = v2254
	var v2255 int32
	_ = v2255
	var v2256 int32
	_ = v2256
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2264 int32
	_ = v2264
	var v2267 int32
	_ = v2267
	var v2312 int32
	_ = v2312
	var v2313 int32
	_ = v2313
	var v2315 int32
	_ = v2315
	var v2316 int32
	_ = v2316
	var v2318 int32
	_ = v2318
	var v2320 int32
	_ = v2320
	var v2368 int32
	_ = v2368
	var v2370 int32
	_ = v2370
	var v2374 int32
	_ = v2374
	var v2376 int32
	_ = v2376
	var v2377 int32
	_ = v2377
	var v2380 int32
	_ = v2380
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2386 int32
	_ = v2386
	var v2390 int32
	_ = v2390
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2403 int32
	_ = v2403
	var v2405 int32
	_ = v2405
	var v2408 int32
	_ = v2408
	var v2409 int32
	_ = v2409
	var v2417 int32
	_ = v2417
	var v2419 int32
	_ = v2419
	var v2434 int32
	_ = v2434
	var v2438 int32
	_ = v2438
	var v2450 int32
	_ = v2450
	var v2469 int32
	_ = v2469
	var v2470 int32
	_ = v2470
	var v2474 int32
	_ = v2474
	var v2480 int32
	_ = v2480
	var v2484 int32
	_ = v2484
	var v2490 int32
	_ = v2490
	var v2491 int32
	_ = v2491
	var v2497 int32
	_ = v2497
	var v2499 int64
	_ = v2499
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2516 int32
	_ = v2516
	var v2518 int32
	_ = v2518
	var v2519 int32
	_ = v2519
	var v2521 int32
	_ = v2521
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2533 int32
	_ = v2533
	var v2535 int64
	_ = v2535
	var v2542 int32
	_ = v2542
	var v2543 int32
	_ = v2543
	var v2545 int32
	_ = v2545
	var v2546 int32
	_ = v2546
	var v2553 int32
	_ = v2553
	var v2554 int32
	_ = v2554
	var v2556 int32
	_ = v2556
	var v2570 int32
	_ = v2570
	var v2586 int32
	_ = v2586
	var v2605 int32
	_ = v2605
	var v2611 int32
	_ = v2611
	var v2613 int64
	_ = v2613
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2623 int32
	_ = v2623
	var v2624 int32
	_ = v2624
	var v2631 int32
	_ = v2631
	var v2634 int32
	_ = v2634
	var v2636 int32
	_ = v2636
	var v2637 int32
	_ = v2637
	var v2641 int32
	_ = v2641
	var v2647 int32
	_ = v2647
	var v2649 int32
	_ = v2649
	var v2655 int32
	_ = v2655
	var v2656 int32
	_ = v2656
	var v2657 int32
	_ = v2657
	var v2658 int32
	_ = v2658
	var v2660 int32
	_ = v2660
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2665 int32
	_ = v2665
	var v2667 int32
	_ = v2667
	var v2668 int32
	_ = v2668
	var v2669 int32
	_ = v2669
	var v2670 int32
	_ = v2670
	var v2672 int32
	_ = v2672
	var v2677 int32
	_ = v2677
	var v2679 int32
	_ = v2679
	var v2681 int32
	_ = v2681
	var v2684 int32
	_ = v2684
	var v2690 int32
	_ = v2690
	var v2696 int32
	_ = v2696
	var v2698 int32
	_ = v2698
	var v2704 int32
	_ = v2704
	var v2705 int32
	_ = v2705
	var v2706 int32
	_ = v2706
	var v2707 int32
	_ = v2707
	var v2709 int32
	_ = v2709
	var v2712 int32
	_ = v2712
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2719 int32
	_ = v2719
	var v2722 int32
	_ = v2722
	var v2723 int32
	_ = v2723
	var v2724 int32
	_ = v2724
	var v2732 int32
	_ = v2732
	var v2736 int32
	_ = v2736
	var v2739 int32
	_ = v2739
	var v2744 int32
	_ = v2744
	var v2748 int32
	_ = v2748
	var v2752 int32
	_ = v2752
	var v2756 int32
	_ = v2756
	var v2762 int32
	_ = v2762
	var v2763 int32
	_ = v2763
	var v2764 int32
	_ = v2764
	var v2767 int32
	_ = v2767
	var v2768 int32
	_ = v2768
	var v2770 int32
	_ = v2770
	var v2773 int32
	_ = v2773
	var v2777 int32
	_ = v2777
	var v2779 int32
	_ = v2779
	var v2787 int32
	_ = v2787
	var v2789 int32
	_ = v2789
	var v2791 int32
	_ = v2791
	var v2794 int32
	_ = v2794
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2800 int64
	_ = v2800
	var v2801 int32
	_ = v2801
	var v2802 int64
	_ = v2802
	var v2805 int32
	_ = v2805
	var v2809 int32
	_ = v2809
	var v2816 int32
	_ = v2816
	var v2822 int32
	_ = v2822
	var v2824 int32
	_ = v2824
	var v2830 int32
	_ = v2830
	var v2836 int32
	_ = v2836
	var v2838 int32
	_ = v2838
	var v2840 int32
	_ = v2840
	var v2845 int32
	_ = v2845
	var v2847 int32
	_ = v2847
	var v2852 int32
	_ = v2852
	var v2854 int32
	_ = v2854
	var v2858 int32
	_ = v2858
	var v2864 int32
	_ = v2864
	var v2866 int32
	_ = v2866
	var v2872 int32
	_ = v2872
	var v2873 int32
	_ = v2873
	var v2877 int32
	_ = v2877
	var v2883 int32
	_ = v2883
	var v2885 int32
	_ = v2885
	var v2891 int32
	_ = v2891
	var v2892 int32
	_ = v2892
	var v2894 int32
	_ = v2894
	var v2898 int32
	_ = v2898
	var v2899 int32
	_ = v2899
	var v2905 int32
	_ = v2905
	var v2906 int32
	_ = v2906
	var v2910 int32
	_ = v2910
	var v2916 int32
	_ = v2916
	var v2918 int32
	_ = v2918
	var v2924 int32
	_ = v2924
	var v2926 int32
	_ = v2926
	var v2927 int32
	_ = v2927
	var v2928 int32
	_ = v2928
	var v2931 int32
	_ = v2931
	var v2933 int32
	_ = v2933
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2936 int32
	_ = v2936
	var v2938 int32
	_ = v2938
	var v2944 int32
	_ = v2944
	var v2950 int32
	_ = v2950
	var v2958 int32
	_ = v2958
	var v2959 int32
	_ = v2959
	var v2963 int32
	_ = v2963
	var v2965 int32
	_ = v2965
	var v2966 int32
	_ = v2966
	var v2969 int32
	_ = v2969
	var v2977 int32
	_ = v2977
	var v2982 int32
	_ = v2982
	var v2988 int32
	_ = v2988
	var v2990 int32
	_ = v2990
	var v2996 int32
	_ = v2996
	var v2997 int32
	_ = v2997
	var v2999 int32
	_ = v2999
	var v3002 int32
	_ = v3002
	var v3008 int32
	_ = v3008
	var v3014 int32
	_ = v3014
	var v3016 int32
	_ = v3016
	var v3022 int32
	_ = v3022
	var v3023 int32
	_ = v3023
	var v3024 int32
	_ = v3024
	var v3025 int32
	_ = v3025
	var v3027 int32
	_ = v3027
	var v3030 int32
	_ = v3030
	var v3032 int32
	_ = v3032
	var v3033 int32
	_ = v3033
	var v3037 int32
	_ = v3037
	var v3040 int32
	_ = v3040
	var v3041 int32
	_ = v3041
	var v3044 int32
	_ = v3044
	var v3049 int32
	_ = v3049
	var v3056 int32
	_ = v3056
	var v3057 int32
	_ = v3057
	var v3059 int32
	_ = v3059
	var v3061 int32
	_ = v3061
	var v3063 int32
	_ = v3063
	var v3065 int32
	_ = v3065
	var v3067 int32
	_ = v3067
	var v3069 int32
	_ = v3069
	var v3074 int32
	_ = v3074
	var v3080 int32
	_ = v3080
	var v3083 int32
	_ = v3083
	var v3087 int32
	_ = v3087
	var v3091 int32
	_ = v3091
	var v3095 int32
	_ = v3095
	var v3098 int32
	_ = v3098
	var v3099 int32
	_ = v3099
	var v3106 int32
	_ = v3106
	var v3107 int32
	_ = v3107
	var v3108 int32
	_ = v3108
	var v3110 int32
	_ = v3110
	var v3114 int32
	_ = v3114
	var v3116 int64
	_ = v3116
	var v3117 int32
	_ = v3117
	var v3126 int32
	_ = v3126
	var v3132 int32
	_ = v3132
	var v3134 int32
	_ = v3134
	var v3140 int32
	_ = v3140
	var v3153 int32
	_ = v3153
	var v3155 int32
	_ = v3155
	var v3160 int32
	_ = v3160
	var v3164 int32
	_ = v3164
	var v3166 int32
	_ = v3166
	var v3169 int32
	_ = v3169
	var v3173 int32
	_ = v3173
	var v3179 int32
	_ = v3179
	var v3181 int32
	_ = v3181
	var v3187 int32
	_ = v3187
	var v3188 int32
	_ = v3188
	var v3190 int32
	_ = v3190
	var v3193 int32
	_ = v3193
	var v3194 int32
	_ = v3194
	var v3197 int32
	_ = v3197
	var v3198 int32
	_ = v3198
	var v3199 int32
	_ = v3199
	var v3201 int64
	_ = v3201
	var v3205 int32
	_ = v3205
	var v3206 int32
	_ = v3206
	var v3209 int32
	_ = v3209
	var v3210 int32
	_ = v3210
	var v3211 int32
	_ = v3211
	var v3213 int32
	_ = v3213
	var v3215 int32
	_ = v3215
	var v3217 int32
	_ = v3217
	var v3221 int32
	_ = v3221
	var v3223 int32
	_ = v3223
	var v3271 int32
	_ = v3271
	var v3273 int32
	_ = v3273
	var v3280 int32
	_ = v3280
	var v3283 int32
	_ = v3283
	var v3284 int32
	_ = v3284
	var v3285 int32
	_ = v3285
	var v3286 int32
	_ = v3286
	var v3290 int32
	_ = v3290
	var v3296 int32
	_ = v3296
	var v3298 int32
	_ = v3298
	var v3304 int32
	_ = v3304
	var v3305 int32
	_ = v3305
	var v3306 int32
	_ = v3306
	var v3321 int32
	_ = v3321
	var v3326 int32
	_ = v3326
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3338 int32
	_ = v3338
	var v3343 int32
	_ = v3343
	var v3344 int32
	_ = v3344
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3360 int32
	_ = v3360
	var v3365 int32
	_ = v3365
	var v3366 int32
	_ = v3366
	var v3372 int32
	_ = v3372
	var v3373 int32
	_ = v3373
	var v3382 int32
	_ = v3382
	var v3387 int32
	_ = v3387
	var v3388 int32
	_ = v3388
	var v3394 int32
	_ = v3394
	var v3395 int32
	_ = v3395
	var v3404 int32
	_ = v3404
	var v3409 int32
	_ = v3409
	var v3410 int32
	_ = v3410
	var v3416 int32
	_ = v3416
	var v3417 int32
	_ = v3417
	var v3426 int32
	_ = v3426
	var v3431 int32
	_ = v3431
	var v3432 int32
	_ = v3432
	var v3438 int32
	_ = v3438
	var v3439 int32
	_ = v3439
	var v3448 int32
	_ = v3448
	var v3453 int32
	_ = v3453
	var v3454 int32
	_ = v3454
	var v3460 int32
	_ = v3460
	var v3461 int32
	_ = v3461
	var v3470 int32
	_ = v3470
	var v3475 int32
	_ = v3475
	var v3476 int32
	_ = v3476
	var v3482 int32
	_ = v3482
	var v3485 int32
	_ = v3485
	var v3486 int32
	_ = v3486
	var v3487 int32
	_ = v3487
	var v3488 int32
	_ = v3488
	var v3499 int32
	_ = v3499
	var v3504 int32
	_ = v3504
	var v3508 int32
	_ = v3508
	var v3512 int32
	_ = v3512
	var v3518 int32
	_ = v3518
	var v3520 int32
	_ = v3520
	var v3526 int32
	_ = v3526
	var v3527 int32
	_ = v3527
	var v3528 int32
	_ = v3528
	var v3537 int32
	_ = v3537
	var v3542 int32
	_ = v3542
	v10 = l9
	v12 = int32(0)
	v46 = m.G0
	v48 = v46 - int32(224)
	m.G0 = v48
	if l3 < v12 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v68 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+16)))
	v69 = v68 + v67
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69)+4))
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69)+12)))
	if v10 != 0 {
		goto L15
	} else {
		goto L16
	}
L2:
	;
	v53 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v53+(l3^int32(-1))<<(uint(int32(2))%32))))
	v67 = v59
	goto L1
L3:
	;
	goto L4
L4:
	;
	v61 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v67 = v61 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L5:
	;
	F_errstart_cold(m, int32(23), int32(0))
	mBase = m.M
	v3508 = m.ExcPending
	if v3508 != 0 {
		goto L20
	} else {
		goto L682
	}
L6:
	;
	v3476 = int32(0)
	base.MemoryFill(m, v2226, v3476, int32(_a_F__bt_insertonpg_0))
	F_errstart_cold(m, int32(21), v3476)
	mBase = m.M
	v3482 = m.ExcPending
	if v3482 != 0 {
		goto L20
	} else {
		goto L678
	}
L7:
	;
	v3454 = int32(0)
	base.MemoryFill(m, v2226, v3454, int32(_a_F__bt_insertonpg_0))
	F_errstart_cold(m, int32(21), v3454)
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		goto L20
	} else {
		goto L675
	}
L8:
	;
	v3432 = int32(0)
	base.MemoryFill(m, v2226, v3432, int32(_a_F__bt_insertonpg_0))
	F_errstart_cold(m, int32(21), v3432)
	mBase = m.M
	v3438 = m.ExcPending
	if v3438 != 0 {
		goto L20
	} else {
		goto L672
	}
L9:
	;
	v3410 = int32(0)
	base.MemoryFill(m, v2226, v3410, int32(_a_F__bt_insertonpg_0))
	F_errstart_cold(m, int32(21), v3410)
	mBase = m.M
	v3416 = m.ExcPending
	if v3416 != 0 {
		goto L20
	} else {
		goto L669
	}
L10:
	;
	v3388 = int32(0)
	base.MemoryFill(m, v2226, v3388, int32(_a_F__bt_insertonpg_0))
	F_errstart_cold(m, int32(21), v3388)
	mBase = m.M
	v3394 = m.ExcPending
	if v3394 != 0 {
		goto L20
	} else {
		goto L666
	}
L11:
	;
	v3366 = int32(0)
	base.MemoryFill(m, v2226, v3366, int32(_a_F__bt_insertonpg_0))
	F_errstart_cold(m, int32(21), v3366)
	mBase = m.M
	v3372 = m.ExcPending
	if v3372 != 0 {
		goto L20
	} else {
		goto L663
	}
L12:
	;
	v3344 = int32(0)
	base.MemoryFill(m, v2226, v3344, int32(_a_F__bt_insertonpg_0))
	F_errstart_cold(m, int32(21), v3344)
	mBase = m.M
	v3350 = m.ExcPending
	if v3350 != 0 {
		goto L20
	} else {
		goto L660
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3330 = m.ExcPending
	if v3330 != 0 {
		goto L20
	} else {
		goto L657
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3280 = m.ExcPending
	if v3280 != 0 {
		goto L20
	} else {
		goto L649
	}
L15:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v67+l8<<(uint(int32(2))%32))+20))
	v79 = v67 + v76&int32(_a_F__bt_insertonpg_1)
	v80 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+7)))
	v85 = int32(_a_F__bt_insertonpg_2)
	if base.B2i32(v80&int32(32) == int32(0))|base.B2i32(v76&v85 == v85) != 0 {
		goto L14
	} else {
		goto L18
	}
L16:
	;
	v101 = l6
	v102 = l8
	v104 = v12
	v105 = v12
	v106 = v12
	goto L17
L17:
	;
	v108 = v72 & int32(2)
	v109 = int32(4)
	v110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+14)))
	v111 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67)+12)))
	v112 = v110 - v111
	if v112 <= v109 {
		goto L25
	} else {
		goto L26
	}
L18:
	;
	v90 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v79)+4)))
	if v90&int32(_a_F__bt_insertonpg_0) == int32(0) {
		goto L14
	} else {
		goto L19
	}
L19:
	;
	v97 = F_CopyIndexTuple(m, l6)
	mBase = m.M
	v98 = m.ExcPending
	if v98 != 0 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	return
L21:
	;
	v99 = F__bt_swap_posting(m, v97, v79, v10)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L20
	} else {
		goto L22
	}
L22:
	;
	v101 = v97
	v102 = l8 + int32(1)
	v104 = v79
	v105 = v99
	v106 = l6
	goto L17
L23:
	;
	if v10 != 0 {
		goto L644
	} else {
		goto L645
	}
L24:
	;
	if base.Ui32(v115-int32(4)) < base.Ui32(l7) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v115 = v109
	goto L27
L26:
	;
	v115 = v112
	goto L27
L27:
	;
	goto L24
L28:
	;
	if l3 < int32(0) {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v2899 = int32(0)
	if l10 == v2899 {
		v2933 = v2899
		v2934 = v2899
		v2935 = v12
		goto L537
	} else {
		goto L538
	}
L31:
	;
	v137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+16)))
	v138 = v137 + v136
	v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+12)))
	if l3 < int32(0) {
		goto L36
	} else {
		goto L37
	}
L32:
	;
	v122 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122+(l3^int32(-1))<<(uint(int32(2))%32))))
	v136 = v128
	goto L31
L33:
	;
	goto L34
L34:
	;
	v130 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v136 = v130 + l3<<(uint(int32(13))%32) + int32(-8192)
	goto L31
L35:
	;
	v162 = v102 & int32(_a_F__bt_insertonpg_3)
	v163 = m.G0
	v165 = v163 + int32(-64)
	m.G0 = v165
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+12)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+19)))
	v172 = v170 << (uint(int32(8)) % 32)
	v174 = v172 - int32(40)
	v176 = v48 + int32(222)
	if base.Ui32(v167) < base.Ui32(int32(25)) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v145+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v160 = v151
	goto L35
L37:
	;
	goto L38
L38:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v160 = v159
	goto L35
L39:
	;
	v182 = int32(0)
	goto L41
L40:
	;
	v182 = int32(base.Ui32(v167+int32(_a_F__bt_insertonpg_4)) >> (uint(int32(2)) % 32))
	goto L41
L41:
	;
	v183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+16)))
	v184 = v136 + v183
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v185 != 0 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v196 = v172 - (int32(base.Ui32(v186)>>(uint(int32(17))%32))+int32(7))&int32(_a_F__bt_insertonpg_5) - int32(44)
	goto L44
L43:
	;
	v196 = v174
	goto L44
L44:
	;
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+14)))
	v198 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+12)))
	v199 = v197 - v198
	v200 = int32(0)
	if v200 < v199 {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v204 = v196 - v203
	v205 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v205 != 0 {
		goto L49
	} else {
		goto L50
	}
L46:
	;
	v203 = v199
	goto L48
L47:
	;
	v203 = v200
	goto L48
L48:
	;
	goto L45
L49:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v205)+4))
	v209 = base.F64_convert_i32_s(v206)
	goto L51
L50:
	;
	v209 = float64(90)
	goto L51
L51:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v101
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v136
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = l7 + int32(4)
	v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v184)+12)))
	v218 = v216 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)) = uint8(v218)
	v220 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+44)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+40)) = v204
	*(*int32)(unsafe.Add(mBase, uint32(v165)+36)) = v196
	*(*int32)(unsafe.Add(mBase, uint32(v165)+32)) = v174
	v227 = v182 & int32(_a_F__bt_insertonpg_3)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+48)) = v227
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)) = uint16(v162)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+29)) = uint8(base.B2i32(v220 == int32(0)))
	v235 = F_palloc(m, v227*int32(10))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+56)) = v235
	v242 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v242 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v243 = int32(2)
	goto L55
L54:
	;
	v243 = int32(1)
	goto L55
L55:
	;
	if base.Ui32(v243) <= base.Ui32(v227) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v247 = int32(0)
	v262 = v243
	v264 = v247
	v269 = v247
	goto L59
L57:
	;
	goto L58
L58:
	;
	if base.Ui32(v162) <= base.Ui32(v227) {
		goto L144
	} else {
		goto L145
	}
L59:
	;
	v295 = v262 & int32(_a_F__bt_insertonpg_3)
	v297 = v295 << (uint(int32(2)) % 32)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v136+int32(20)+v297)))
	v305 = (int32(base.Ui32(v299)>>(uint(int32(17))%32)) + int32(7)) & int32(_a_F__bt_insertonpg_5)
	v307 = v305 | int32(4)
	if base.Ui32(v295) < base.Ui32(v162) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L58
L61:
	;
	v797 = v269 + v307
	v800 = v262 + int32(1)
	if base.Ui32(v800&int32(_a_F__bt_insertonpg_3)) <= base.Ui32(v227) {
		v262 = v800
		v264 = int32(0) - v797
		v269 = v797
		goto L59
	} else {
		goto L142
	}
L62:
	;
	v310 = v163 + int32(-52)
	v311 = int32(0)
	v315 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+18)))
	if v315 == v295 {
		goto L69
	} else {
		goto L70
	}
L63:
	;
	goto L64
L64:
	;
	if base.Ui32(v162) < base.Ui32(v295) {
		goto L91
	} else {
		goto L92
	}
L65:
	;
	goto L61
L66:
	;
	v388 = v383 + v384 - (v269 + v381)
	v389 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
	if v385 != 0 {
		goto L80
	} else {
		goto L81
	}
L67:
	;
	v381 = v376
	v382 = v377
	v383 = v378
	v384 = v379
	v385 = int32(1)
	goto L66
L68:
	;
	v376 = v371
	v377 = v372
	v378 = int32(-8)
	v379 = v374
	goto L67
L69:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v310)+24))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v310)+28))
	v320 = v317 + (v269 - v318)
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v310)+20))
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v310)+12))
	v323 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+16)))
	if v323&int32(1) != 0 {
		v371 = v322
		v372 = v320
		v374 = v321
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v326 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v310)+16)))
	if base.B2i32(v326 != int32(1))|base.B2i32(base.Ui32(v307) < base.Ui32(int32(65))) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v381 = v322
	v382 = v320
	v383 = v311
	v384 = v321
	v385 = v311
	goto L66
L73:
	;
	v334 = int32(-8)
	v335 = *(*int32)(unsafe.Add(mBase, uint32(v310)+4))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v335+v295<<(uint(int32(2))%32))+20))
	v342 = v335 + v339&int32(_a_F__bt_insertonpg_1)
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v342)+6)))
	if v343&int32(_a_F__bt_insertonpg_0) == int32(0) {
		v359 = v334
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v310)+24))
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v310)+28))
	v368 = v365 + (v269 - v366)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v310)+20))
	if v326 != 0 {
		v371 = v307
		v372 = v368
		v374 = v369
		goto L68
	} else {
		goto L79
	}
L76:
	;
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v310)+24))
	v361 = *(*int32)(unsafe.Add(mBase, uint32(v310)+28))
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v310)+20))
	v376 = v307
	v377 = v360 + (v269 - v361)
	v378 = v359
	v379 = v364
	goto L67
L77:
	;
	v348 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v342)+5)))
	if v348&int32(32) == int32(0) {
		v359 = v334
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v342)+2)))
	v359 = v343&int32(_a_F__bt_insertonpg_6) - v355 - int32(8)
	goto L76
L79:
	;
	v381 = v307
	v382 = v368
	v383 = int32(0)
	v384 = v369
	v385 = v311
	goto L66
L80:
	;
	v394 = int32(0)
	goto L82
L81:
	;
	v394 = v381 + int32(_a_F__bt_insertonpg_7)
	goto L82
L82:
	;
	v395 = v382 - v389 + v394
	if (v388|v395)&int32(_a_F__bt_insertonpg_8) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v401 = *(*int32)(unsafe.Add(mBase, uint32(v310)+32))
	if base.Ui32(v401) < base.Ui32(v381) {
		goto L86
	} else {
		goto L87
	}
L84:
	;
	goto L85
L85:
	;
	goto L65
L86:
	;
	v403 = v401
	goto L88
L87:
	;
	v403 = v381
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v310)+32)) = v403
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v310)+44))
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v310)+40))
	v407 = int32(10)
	v410 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v405+v406*v407))) = uint16(v410)
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v310)+44))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v310)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v412+v413*v407)+2)) = uint16(v388)
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v310)+44))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v310)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v418+v419*v407)+4)) = uint16(v395)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v310)+44))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v310)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v424+v425*v407)+6)) = uint16(v295)
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v310)+44))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v310)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v430+v431*v407)+8)) = uint8(v410)
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v310)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v310)+40)) = v437 + int32(1)
	goto L85
L89:
	;
	v746 = v745 + v741
	if (v746|v742)&int32(_a_F__bt_insertonpg_8) != 0 {
		goto L61
	} else {
		goto L138
	}
L90:
	;
	v741 = v732
	v742 = v264 - v307 + v735 + v734 - v733
	v745 = int32(0)
	goto L89
L91:
	;
	v443 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if base.B2i32(v443 != int32(1))|base.B2i32(base.Ui32(v299) < base.Ui32(int32(_a_F__bt_insertonpg_9))) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v494 = v163 + int32(-52)
	v495 = int32(0)
	v499 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v494)+18)))
	if v499 == v295 {
		goto L105
	} else {
		goto L106
	}
L94:
	;
	v451 = int32(-8)
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v452+v297)+20))
	v457 = v452 + v454&int32(_a_F__bt_insertonpg_1)
	v458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457)+6)))
	if v458&int32(_a_F__bt_insertonpg_0) == int32(0) {
		v474 = v451
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v484 = v481 + (v269 - v482)
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	if v443 != 0 {
		v732 = v484
		v733 = v486
		v734 = int32(-8)
		v735 = v487
		goto L90
	} else {
		goto L100
	}
L97:
	;
	v475 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	v732 = v475 + v269 - v477
	v733 = v479
	v734 = v474
	v735 = v480
	goto L90
L98:
	;
	v463 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v457)+5)))
	if v463&int32(32) == int32(0) {
		v474 = v451
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v470 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v457)+2)))
	v474 = v458&int32(_a_F__bt_insertonpg_6) - v470 - int32(8)
	goto L97
L100:
	;
	v741 = v484
	v742 = v264 + v487 - (v486 + v307)
	v745 = v305 + int32(_a_F__bt_insertonpg_5)
	goto L89
L101:
	;
	v626 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if base.B2i32(v626 != int32(1))|base.B2i32(base.Ui32(v299) < base.Ui32(int32(_a_F__bt_insertonpg_9))) == int32(0) {
		goto L127
	} else {
		goto L128
	}
L102:
	;
	v572 = v567 + v568 - (v269 + v565)
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	if v569 != 0 {
		goto L116
	} else {
		goto L117
	}
L103:
	;
	v565 = v560
	v566 = v561
	v567 = v562
	v568 = v563
	v569 = int32(1)
	goto L102
L104:
	;
	v560 = v555
	v561 = v556
	v562 = int32(-8)
	v563 = v558
	goto L103
L105:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v494)+24))
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v494)+28))
	v504 = v501 + (v269 - v502)
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v494)+20))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v494)+12))
	v507 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+16)))
	if v507&int32(1) != 0 {
		v555 = v506
		v556 = v504
		v558 = v505
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v494)+16)))
	if base.B2i32(v510 != int32(1))|base.B2i32(base.Ui32(v307) < base.Ui32(int32(65))) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v565 = v506
	v566 = v504
	v567 = v495
	v568 = v505
	v569 = v495
	goto L102
L109:
	;
	v518 = int32(-8)
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v494)+4))
	v523 = *(*int32)(unsafe.Add(mBase, uint32(v519+v295<<(uint(int32(2))%32))+20))
	v526 = v519 + v523&int32(_a_F__bt_insertonpg_1)
	v527 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v526)+6)))
	if v527&int32(_a_F__bt_insertonpg_0) == int32(0) {
		v543 = v518
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v494)+24))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(v494)+28))
	v552 = v549 + (v269 - v550)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v494)+20))
	if v510 != 0 {
		v555 = v307
		v556 = v552
		v558 = v553
		goto L104
	} else {
		goto L115
	}
L112:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v494)+24))
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v494)+28))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v494)+20))
	v560 = v307
	v561 = v544 + (v269 - v545)
	v562 = v543
	v563 = v548
	goto L103
L113:
	;
	v532 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v526)+5)))
	if v532&int32(32) == int32(0) {
		v543 = v518
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v526)+2)))
	v543 = v527&int32(_a_F__bt_insertonpg_6) - v539 - int32(8)
	goto L112
L115:
	;
	v565 = v307
	v566 = v552
	v567 = int32(0)
	v568 = v553
	v569 = v495
	goto L102
L116:
	;
	v578 = int32(0)
	goto L118
L117:
	;
	v578 = v565 + int32(_a_F__bt_insertonpg_7)
	goto L118
L118:
	;
	v579 = v566 - v573 + v578
	if (v572|v579)&int32(_a_F__bt_insertonpg_8) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v494)+32))
	if base.Ui32(v585) < base.Ui32(v565) {
		goto L122
	} else {
		goto L123
	}
L120:
	;
	goto L121
L121:
	;
	goto L101
L122:
	;
	v587 = v585
	goto L124
L123:
	;
	v587 = v565
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v494)+32)) = v587
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v494)+44))
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v494)+40))
	v591 = int32(10)
	v594 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v589+v590*v591))) = uint16(v594)
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v494)+44))
	v597 = *(*int32)(unsafe.Add(mBase, uint32(v494)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v596+v597*v591)+2)) = uint16(v572)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v494)+44))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v494)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v602+v603*v591)+4)) = uint16(v579)
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v494)+44))
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v494)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v608+v609*v591)+6)) = uint16(v295)
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v494)+44))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v494)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v614+v615*v591)+8)) = uint8(v594)
	v621 = *(*int32)(unsafe.Add(mBase, uint32(v494)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v494)+40)) = v621 + int32(1)
	goto L121
L125:
	;
	v688 = v687 + v684
	if (v688|v685)&int32(_a_F__bt_insertonpg_8) != 0 {
		goto L61
	} else {
		goto L134
	}
L126:
	;
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v684 = v675
	v685 = v264 - v307 + v676 + v677 - v681
	v687 = int32(0)
	goto L125
L127:
	;
	v634 = int32(-8)
	v635 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v635+v297)+20))
	v640 = v635 + v637&int32(_a_F__bt_insertonpg_1)
	v641 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v640)+6)))
	if v641&int32(_a_F__bt_insertonpg_0) == int32(0) {
		v657 = v634
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v664 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v666 = v663 + (v269 - v664)
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	if v626 != 0 {
		v675 = v666
		v676 = v668
		v677 = int32(-8)
		goto L126
	} else {
		goto L133
	}
L130:
	;
	v658 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	v675 = v658 + v269 - v660
	v676 = v662
	v677 = v657
	goto L126
L131:
	;
	v646 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v640)+5)))
	if v646&int32(32) == int32(0) {
		v657 = v634
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v653 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v640)+2)))
	v657 = v641&int32(_a_F__bt_insertonpg_6) - v653 - int32(8)
	goto L130
L133:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v684 = v666
	v685 = v668 + v264 - (v307 + v670)
	v687 = v305 + int32(_a_F__bt_insertonpg_5)
	goto L125
L134:
	;
	v692 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	if base.Ui32(v692) < base.Ui32(v307) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v694 = v692
	goto L137
L136:
	;
	v694 = v307
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+44)) = v694
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v697 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v698 = int32(10)
	v701 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v696+v697*v698))) = uint16(v701)
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v703+v704*v698)+2)) = uint16(v685)
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v710 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v709+v710*v698)+4)) = uint16(v688)
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v715+v716*v698)+6)) = uint16(v262)
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v722 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v726 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v721+v722*v698)+8)) = uint8(v726)
	v728 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = v728 + v726
	goto L61
L138:
	;
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	if base.Ui32(v750) < base.Ui32(v307) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v752 = v750
	goto L141
L140:
	;
	v752 = v307
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+44)) = v752
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v756 = int32(10)
	v759 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v754+v755*v756))) = uint16(v759)
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v762 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v761+v762*v756)+2)) = uint16(v742)
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v767+v768*v756)+4)) = uint16(v746)
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v773+v774*v756)+6)) = uint16(v262)
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v784 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v779+v780*v756)+8)) = uint8(v784)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = v786 + v784
	goto L61
L142:
	;
	goto L60
L143:
	;
	if v943 != 0 {
		goto L163
	} else {
		goto L164
	}
L144:
	;
	v850 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v943 = v850
	goto L143
L145:
	;
	goto L146
L146:
	;
	v851 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	if v851 == v162 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v896 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v897 = v895 + v891
	if (v892|v897)&int32(_a_F__bt_insertonpg_8) != 0 {
		v943 = v896
		goto L143
	} else {
		goto L155
	}
L148:
	;
	v891 = v882 - v881
	v892 = v884 - (v883 + v204) - int32(8)
	v893 = v883
	v895 = int32(0)
	goto L147
L149:
	;
	v891 = v873 - v872
	v892 = v875 - (v874 + v204)
	v893 = v874
	v895 = v874 + int32(_a_F__bt_insertonpg_7)
	goto L147
L150:
	;
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v856 = v853 + (v204 - v854)
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	v858 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v859 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if v859&int32(1) == int32(0) {
		v872 = v858
		v873 = v856
		v874 = v858
		v875 = v857
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v867 = v864 + (v204 - v865)
	v868 = int32(0)
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	v871 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if v871 != 0 {
		v881 = v869
		v882 = v867
		v883 = v868
		v884 = v870
		goto L148
	} else {
		goto L154
	}
L153:
	;
	v881 = v858
	v882 = v856
	v883 = v858
	v884 = v857
	goto L148
L154:
	;
	v872 = v869
	v873 = v867
	v874 = v868
	v875 = v870
	goto L149
L155:
	;
	v901 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	if base.Ui32(v901) < base.Ui32(v893) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v903 = v901
	goto L158
L157:
	;
	v903 = v893
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+44)) = v903
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v906 = int32(10)
	v909 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v905+v896*v906))) = uint16(v909)
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v912 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v911+v912*v906)+2)) = uint16(v892)
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v917+v918*v906)+4)) = uint16(v897)
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v924 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v923+v924*v906)+6)) = uint16(v162)
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v929+v930*v906)+8)) = uint8(v909)
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v938 = v936 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = v938
	v943 = v938
	goto L143
L159:
	;
	m.G0 = v165 - int32(-64)
	v2124 = v2089 & int32(_a_F__bt_insertonpg_3)
	v2125 = F_PageGetTempPage(m, v136)
	mBase = m.M
	v2126 = m.ExcPending
	if v2126 != 0 {
		goto L20
	} else {
		goto L339
	}
L160:
	;
	v1229 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1232 = v1229 + v943*int32(10)
	v1235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1232-int32(2)))))
	v1238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1232-int32(4)))))
	v1239 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1229)+8)))
	v1240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1229)+6)))
	if int32(0) < v943 {
		goto L205
	} else {
		goto L206
	}
L161:
	;
	v1197 = v1151
	v1227 = float64(0.5)
	goto L160
L162:
	;
	v1151 = int32(0)
	goto L161
L163:
	;
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+29)))
	v947 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if v947 != int32(1) {
		goto L166
	} else {
		goto L167
	}
L164:
	;
	goto L165
L165:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1117 = m.ExcPending
	if v1117 != 0 {
		goto L20
	} else {
		goto L202
	}
L166:
	;
	v1197 = v946
	v1227 = float64(0.7)
	goto L160
L167:
	;
	goto L168
L168:
	;
	if v946&int32(1) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1197 = int32(1)
	v1227 = base.F64_div(v209, float64(100))
	goto L160
L170:
	;
	goto L171
L171:
	;
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v957 = *(*int32)(unsafe.Add(mBase, uint32(v956)+192))
	v958 = int32(*(*int16)(unsafe.Add(mBase, uint32(v957)+10)))
	if v958 == int32(1) {
		goto L162
	} else {
		goto L172
	}
L172:
	;
	v961 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	if v961 == int32(2) {
		goto L162
	} else {
		goto L173
	}
L173:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v965 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	if base.B2i32(v964 != v965)|base.B2i32(base.Ui32(int32(28)) < base.Ui32(v964)) != 0 {
		goto L162
	} else {
		goto L174
	}
L174:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	if v970 != v964*(v227-int32(1)) {
		goto L162
	} else {
		goto L175
	}
L175:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	if base.Ui32(v182&int32(_a_F__bt_insertonpg_3)) < base.Ui32(v961) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v982 = *(*int32)(unsafe.Add(mBase, uint32(v975+v227<<(uint(int32(2))%32))+20))
	v986 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v987 = F__bt_keep_natts_fast(m, v956, v975+v982&int32(_a_F__bt_insertonpg_1), v986)
	mBase = m.M
	v988 = m.ExcPending
	if v988 != 0 {
		goto L20
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v1004 = *(*int32)(unsafe.Add(mBase, uint32(v975+(v961-int32(1))&int32(_a_F__bt_insertonpg_3)<<(uint(int32(2))%32))+20))
	v1007 = v975 + v1004&int32(_a_F__bt_insertonpg_1)
	v1008 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007)+7)))
	if v1008&int32(32) != 0 {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	if v987 < int32(2) {
		goto L162
	} else {
		goto L180
	}
L180:
	;
	if v958 < v987 {
		v1197 = int32(0)
		v1227 = float64(0.5)
		goto L160
	} else {
		goto L181
	}
L181:
	;
	v1197 = int32(1)
	v1227 = base.F64_div(v209, float64(100))
	goto L160
L182:
	;
	v1011 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1007)+5)))
	if v1011&int32(32) != 0 {
		goto L162
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v1014 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1007)+2)))
	v1015 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1007))))
	v1016 = int32(16)
	v1018 = v1014 | v1015<<(uint(v1016)%32)
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1020 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1019))))
	v1023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1019)+2)))
	v1024 = v1020<<(uint(v1016)%32) | v1023
	if v1018 != v1024 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	goto L184
L186:
	;
	if v1018+int32(1) != v1024 {
		goto L162
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v1032 = F__bt_keep_natts_fast(m, v956, v1007, v1019)
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L20
	} else {
		goto L191
	}
L189:
	;
	v1029 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1019)+4)))
	if v1029 != int32(1) {
		goto L162
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	if base.B2i32(v1032 < int32(2))|base.B2i32(v958 < v1032) != 0 {
		goto L162
	} else {
		goto L192
	}
L192:
	;
	v1038 = int32(1)
	v1040 = base.F64_div(v209, float64(100))
	if base.F64_lt(v1040, base.F64_div(base.F64_convert_i32_u(v961), base.F64_convert_i32_u((v182+v1038)&int32(_a_F__bt_insertonpg_3)))) != 0 {
		v1197 = v1038
		v1227 = v1040
		goto L160
	} else {
		goto L193
	}
L193:
	;
	if v943 <= int32(0) {
		goto L162
	} else {
		goto L194
	}
L194:
	;
	v1052 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1069 = int32(0)
	goto L195
L195:
	;
	v1100 = v1052 + v1069*int32(10)
	v1101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1100)+8)))
	if v1101 != int32(1) {
		goto L197
	} else {
		goto L198
	}
L196:
	;
	v1151 = int32(0)
	goto L161
L197:
	;
	v1112 = v1069 + int32(1)
	if v1112 != v943 {
		v1069 = v1112
		goto L195
	} else {
		goto L201
	}
L198:
	;
	v1104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1100)+6)))
	if v162 != v1104 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	F_pfree(m, v1052)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L20
	} else {
		goto L200
	}
L200:
	;
	v1108 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v1108)
	v2089 = v162
	goto L159
L201:
	;
	goto L196
L202:
	;
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v1118 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_10), v165)
	mBase = m.M
	v1124 = m.ExcPending
	if v1124 != 0 {
		goto L20
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_11), int32(261), int32(_a_F__bt_insertonpg_12))
	mBase = m.M
	v1129 = m.ExcPending
	if v1129 != 0 {
		goto L20
	} else {
		goto L204
	}
L204:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L205:
	;
	v1264 = int32(0)
	goto L208
L206:
	;
	v1336 = v943
	v1339 = v947
	v1341 = v1229
	goto L207
L207:
	;
	F_pg_qsort(m, v1341, v1336, int32(10), int32(239))
	mBase = m.M
	v1369 = m.ExcPending
	if v1369 != 0 {
		goto L20
	} else {
		goto L215
	}
L208:
	;
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1296 = v1293 + v1264*int32(10)
	v1297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1296)+2)))
	if v1197&int32(1) != 0 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	v1319 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1320 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	v1336 = v1317
	v1339 = v1320
	v1341 = v1319
	goto L207
L210:
	;
	v1311 = base.I32_extend16_s(v1308) >> (uint(int32(15)) % 32)
	v1313 = v1311 ^ v1308 - v1311
	*(*uint16)(unsafe.Add(mBase, uint32(v1296))) = uint16(v1313)
	v1316 = v1264 + int32(1)
	v1317 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	if v1316 < v1317 {
		v1264 = v1316
		goto L208
	} else {
		goto L214
	}
L211:
	;
	v1301 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1296)+4)))
	v1308 = base.I32_trunc_sat_f64_s(base.F64_sub(base.F64_mul(v1227, base.F64_convert_i32_s(base.I32_extend16_s(v1297))), base.F64_mul(base.F64_sub(float64(1), v1227), base.F64_convert_i32_s(v1301))))
	goto L210
L212:
	;
	goto L213
L213:
	;
	v1306 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1296)+4)))
	v1308 = v1297 - v1306
	goto L210
L214:
	;
	goto L209
L215:
	;
	if v1336 < int32(2) {
		v1468 = v1336
		goto L216
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+60)) = v1468
	v1497 = int32(1)
	if v1339&v1497 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L217:
	;
	v1372 = int32(1)
	v1373 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1341)+4)))
	if v1339&v1372 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1378 = float64(0.05)
	goto L220
L219:
	;
	v1378 = float64(0.075)
	goto L220
L220:
	;
	v1379 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v1382 = base.I32_trunc_sat_f64_s(base.F64_mul(v1378, base.F64_convert_i32_s(v1379)))
	v1385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1341)+2)))
	v1409 = v1372
	goto L221
L221:
	;
	v1439 = v1341 + v1409*int32(10)
	v1440 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1439)+2)))
	if v1440 < base.I32_extend16_s(v1385-v1382) {
		v1468 = v1409
		goto L216
	} else {
		goto L223
	}
L222:
	;
	v1468 = v1336
	goto L216
L223:
	;
	v1443 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1439)+4)))
	if base.B2i32(base.I32_extend16_s(v1385+v1382) < v1440)|base.B2i32(v1443 < base.I32_extend16_s(v1373-v1382))|base.B2i32(base.I32_extend16_s(v1373+v1382) < v1443) != 0 {
		v1468 = v1409
		goto L216
	} else {
		goto L224
	}
L224:
	;
	v1449 = v1409 + int32(1)
	if v1449 != v1336 {
		v1409 = v1449
		goto L221
	} else {
		goto L225
	}
L225:
	;
	goto L222
L226:
	;
	if v1843 < v1841 {
		goto L297
	} else {
		goto L298
	}
L227:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	v1826 = v1502
	v1841 = v1336
	v1843 = v1468
	v1846 = v1341
	v1848 = v1497
	goto L226
L228:
	;
	goto L229
L229:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v1503)+192))
	v1505 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1504)+10)))
	if v1468 < v1336 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1507 = v1468
	goto L232
L231:
	;
	v1507 = v1336
	goto L232
L232:
	;
	v1508 = int32(0)
	v1509 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1341)+6)))
	v1524 = v1508
	v1525 = v1508
	v1527 = v1507
	goto L233
L233:
	;
	v1557 = v1527 - int32(1)
	v1560 = v1341 + v1557*int32(10)
	v1561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1560)+6)))
	if base.Ui32(v1561) < base.Ui32(v1509) {
		goto L243
	} else {
		goto L244
	}
L234:
	;
	v1599 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1588)+6)))
	v1600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1589)+8)))
	v1603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1589)+6)))
	v1604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	if base.B2i32(v1600 != int32(1))|base.B2i32(v1603 != v1604) == int32(0) {
		goto L263
	} else {
		goto L264
	}
L235:
	;
	if int32(0) < v1557 {
		goto L258
	} else {
		goto L259
	}
L236:
	;
	v1588 = v1525
	v1589 = v1560
	goto L235
L237:
	;
	v1584 = int32(0)
	if v1582 == v1584 {
		v1524 = v1584
		v1525 = v1583
		v1527 = v1557
		goto L233
	} else {
		goto L257
	}
L238:
	;
	v1582 = v1580
	v1583 = v1560
	goto L237
L239:
	;
	if v1525 != 0 {
		v1582 = v1524
		v1583 = v1525
		goto L237
	} else {
		goto L256
	}
L240:
	;
	if v1524 != 0 {
		goto L252
	} else {
		goto L253
	}
L241:
	;
	if v1564&int32(1) == int32(0) {
		goto L239
	} else {
		goto L251
	}
L242:
	;
	v1588 = v1525
	v1589 = v1524
	goto L235
L243:
	;
	if v1524 != 0 {
		goto L242
	} else {
		goto L246
	}
L244:
	;
	goto L245
L245:
	;
	if base.Ui32(v1509) < base.Ui32(v1561) {
		goto L239
	} else {
		goto L247
	}
L246:
	;
	goto L236
L247:
	;
	v1564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1341)+8)))
	v1565 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1560)+8)))
	if v1565 != 0 {
		goto L241
	} else {
		goto L248
	}
L248:
	;
	if v1564&int32(1) == int32(0) {
		goto L240
	} else {
		goto L249
	}
L249:
	;
	if v1524 == int32(0) {
		goto L236
	} else {
		goto L250
	}
L250:
	;
	goto L242
L251:
	;
	goto L240
L252:
	;
	v1577 = v1524
	goto L254
L253:
	;
	v1577 = v1560
	goto L254
L254:
	;
	if v1525 != 0 {
		v1588 = v1525
		v1589 = v1577
		goto L235
	} else {
		goto L255
	}
L255:
	;
	v1580 = v1577
	goto L238
L256:
	;
	v1580 = v1524
	goto L238
L257:
	;
	v1588 = v1583
	v1589 = v1582
	goto L235
L258:
	;
	v1594 = int32(0)
	if v1588 == v1594 {
		v1524 = v1589
		v1525 = v1594
		v1527 = v1557
		goto L233
	} else {
		goto L261
	}
L259:
	;
	goto L260
L260:
	;
	goto L234
L261:
	;
	goto L260
L262:
	;
	v1624 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1588)+8)))
	if v1624|base.B2i32(v1599 != v1604) == int32(0) {
		goto L267
	} else {
		goto L268
	}
L263:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1623 = v1609
	goto L262
L264:
	;
	goto L265
L265:
	;
	v1610 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1618 = *(*int32)(unsafe.Add(mBase, uint32(v1610+(v1603-int32(1))&int32(_a_F__bt_insertonpg_3)<<(uint(int32(2))%32))+20))
	v1623 = v1610 + v1618&int32(_a_F__bt_insertonpg_1)
	goto L262
L266:
	;
	v1640 = F__bt_keep_natts_fast(m, v1503, v1623, v1639)
	mBase = m.M
	v1641 = m.ExcPending
	if v1641 != 0 {
		goto L20
	} else {
		goto L270
	}
L267:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1639 = v1629
	goto L266
L268:
	;
	goto L269
L269:
	;
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1630+v1599<<(uint(int32(2))%32))+20))
	v1639 = v1630 + v1634&int32(_a_F__bt_insertonpg_1)
	goto L266
L270:
	;
	if v1640 <= v1505 {
		v1826 = v1640
		v1841 = v1336
		v1843 = v1468
		v1846 = v1341
		v1848 = v1497
		goto L226
	} else {
		goto L271
	}
L271:
	;
	v1645 = int32(0)
	if base.B2i32(v1239&int32(1) == v1645)|base.B2i32(v1604 != v1240) == v1645 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	if v1235&int32(1)|base.B2i32(v1604 != v1238) == int32(0) {
		goto L277
	} else {
		goto L278
	}
L273:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1665 = v1651
	goto L272
L274:
	;
	goto L275
L275:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v1652+(v1240-int32(1))&int32(_a_F__bt_insertonpg_3)<<(uint(int32(2))%32))+20))
	v1665 = v1652 + v1660&int32(_a_F__bt_insertonpg_1)
	goto L272
L276:
	;
	v1683 = F__bt_keep_natts_fast(m, v1503, v1665, v1682)
	mBase = m.M
	v1684 = m.ExcPending
	if v1684 != 0 {
		goto L20
	} else {
		goto L280
	}
L277:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1682 = v1672
	goto L276
L278:
	;
	goto L279
L279:
	;
	v1673 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1677 = *(*int32)(unsafe.Add(mBase, uint32(v1673+v1238<<(uint(int32(2))%32))+20))
	v1682 = v1673 + v1677&int32(_a_F__bt_insertonpg_1)
	goto L276
L280:
	;
	if v1505 < v1683 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+29)))
	if v1686 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+60)) = v1336
	v1826 = v1505
	v1841 = v1336
	v1843 = v1336
	v1846 = v1341
	v1848 = int32(0)
	goto L226
L284:
	;
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v1689)+24))
	v1694 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1695 = F__bt_keep_natts_fast(m, v1503, v1689+v1690&int32(_a_F__bt_insertonpg_1), v1694)
	mBase = m.M
	v1696 = m.ExcPending
	if v1696 != 0 {
		goto L20
	} else {
		goto L287
	}
L285:
	;
	v1698 = v1683
	goto L286
L286:
	;
	if int32(0) < v1336 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	if v1505 < v1695 {
		v1826 = v1695
		v1841 = v1336
		v1843 = v1468
		v1846 = v1341
		v1848 = v1497
		goto L226
	} else {
		goto L288
	}
L288:
	;
	v1698 = v1695
	goto L286
L289:
	;
	v1719 = int32(0)
	goto L292
L290:
	;
	v1788 = v1336
	v1793 = v1341
	goto L291
L291:
	;
	F_pg_qsort(m, v1793, v1788, int32(10), int32(239))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L20
	} else {
		goto L295
	}
L292:
	;
	v1748 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1751 = v1748 + v1719*int32(10)
	v1752 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1751)+2)))
	v1756 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1751)+4)))
	v1761 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1752), float64(0.96)), base.F64_mul(base.F64_convert_i32_s(v1756), float64(-0.040000000000000036))))
	v1764 = base.I32_extend16_s(v1761) >> (uint(int32(15)) % 32)
	v1766 = v1761 ^ v1764 - v1764
	*(*uint16)(unsafe.Add(mBase, uint32(v1751))) = uint16(v1766)
	v1769 = v1719 + int32(1)
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	if v1769 < v1770 {
		v1719 = v1769
		goto L292
	} else {
		goto L294
	}
L293:
	;
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1788 = v1770
	v1793 = v1772
	goto L291
L294:
	;
	goto L293
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+60)) = int32(1)
	v1826 = v1698
	v1841 = v1788
	v1843 = int32(1)
	v1846 = v1793
	v1848 = v1497
	goto L226
L296:
	;
	v2053 = v1846 + v2021*int32(10)
	v2054 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+29)))
	if (v1848|v2054)&int32(1) != 0 {
		v2069 = v2053
		goto L329
	} else {
		goto L330
	}
L297:
	;
	v1872 = v1843
	goto L299
L298:
	;
	v1872 = v1841
	goto L299
L299:
	;
	if v1872 <= int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v2021 = int32(0)
	goto L296
L301:
	;
	goto L302
L302:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1878 = v1876 + int32(20)
	v1879 = int32(0)
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1883 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v1884 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	v1887 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	v1899 = int32(2147483647)
	v1902 = v1879
	v1904 = v1879
	goto L303
L303:
	;
	v1936 = v1846 + v1902*int32(10)
	v1937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1936)+6)))
	v1938 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1936)+8)))
	if v1884&int32(1) == int32(0) {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	v2021 = v2000
	goto L296
L305:
	;
	v1999 = base.B2i32(v1998 < v1899)
	if v1998 < v1899 {
		goto L321
	} else {
		goto L322
	}
L306:
	;
	if v1938&int32(1) == int32(0) {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	goto L308
L308:
	;
	if v1938&int32(1) != 0 {
		goto L315
	} else {
		goto L316
	}
L309:
	;
	if v1937 == v1887 {
		v1998 = v1883
		goto L305
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v1949 = *(*int32)(unsafe.Add(mBase, uint32(v1878+v1937<<(uint(int32(2))%32))))
	v1998 = (int32(base.Ui32(v1949)>>(uint(int32(17))%32))+int32(7))&int32(_a_F__bt_insertonpg_5) | int32(4)
	goto L305
L312:
	;
	goto L311
L313:
	;
	v1994 = F__bt_keep_natts_fast(m, v1881, v1992, v1993)
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L20
	} else {
		goto L320
	}
L314:
	;
	v1988 = *(*int32)(unsafe.Add(mBase, uint32(v1878+v1937<<(uint(int32(2))%32))))
	v1992 = v1984
	v1993 = v1876 + v1988&int32(_a_F__bt_insertonpg_1)
	goto L313
L315:
	;
	if v1937 == v1887 {
		v1984 = v1882
		goto L314
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v1979 = *(*int32)(unsafe.Add(mBase, uint32(v1878+(v1937-int32(1))&int32(_a_F__bt_insertonpg_3)<<(uint(int32(2))%32))))
	v1982 = v1876 + v1979&int32(_a_F__bt_insertonpg_1)
	if v1937 == v1887 {
		v1992 = v1982
		v1993 = v1882
		goto L313
	} else {
		goto L319
	}
L318:
	;
	v1968 = *(*int32)(unsafe.Add(mBase, uint32(v1878+(v1937-int32(1))&int32(_a_F__bt_insertonpg_3)<<(uint(int32(2))%32))))
	v1984 = v1876 + v1968&int32(_a_F__bt_insertonpg_1)
	goto L314
L319:
	;
	v1984 = v1982
	goto L314
L320:
	;
	v1998 = v1994
	goto L305
L321:
	;
	v2000 = v1902
	goto L323
L322:
	;
	v2000 = v1904
	goto L323
L323:
	;
	if v1998 <= v1826 {
		v2021 = v2000
		goto L296
	} else {
		goto L324
	}
L324:
	;
	if v1998 < v1899 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v2002 = v1998
	goto L327
L326:
	;
	v2002 = v1899
	goto L327
L327:
	;
	v2004 = v1902 + int32(1)
	if v2004 != v1872 {
		v1899 = v2002
		v1902 = v2004
		v1904 = v2000
		goto L303
	} else {
		goto L328
	}
L328:
	;
	goto L304
L329:
	;
	v2070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2069)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v2070)
	v2072 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2069)+6)))
	F_pfree(m, v1846)
	mBase = m.M
	v2074 = m.ExcPending
	if v2074 != 0 {
		goto L20
	} else {
		goto L338
	}
L330:
	;
	v2058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2053)+8)))
	if v2058 != 0 {
		v2069 = v2053
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v2059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2053)+6)))
	v2060 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	if base.Ui32(v2059) < base.Ui32(v2060+int32(9)) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v2064 = v1846
	goto L334
L333:
	;
	v2064 = v2053
	goto L334
L334:
	;
	if base.Ui32(v2060) <= base.Ui32(v2059) {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v2066 = v2064
	goto L337
L336:
	;
	v2066 = v2053
	goto L337
L337:
	;
	v2069 = v2066
	goto L329
L338:
	;
	v2089 = v2072
	goto L159
L339:
	;
	F_PageInit(m, v2125, int32(_a_F__bt_insertonpg_0), int32(16))
	mBase = m.M
	goto L340
L340:
	;
	v2130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2125)+16)))
	v2131 = v2125 + v2130
	v2132 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v2136 = v2132&int32(_a_F__bt_insertonpg_13) | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v2131)+12)) = uint16(v2136)
	v2138 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	*(*int32)(unsafe.Add(mBase, uint32(v2131))) = v2138
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2131)+8)) = v2140
	v2142 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	*(*int64)(unsafe.Add(mBase, uint32(v2125))) = v2142
	if v10&int32(_a_F__bt_insertonpg_3) != 0 {
		goto L341
	} else {
		goto L342
	}
L341:
	;
	v2149 = v102 - int32(1)
	goto L343
L342:
	;
	v2149 = int32(0)
	goto L343
L343:
	;
	v2151 = v140 & int32(1)
	v2152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+222)))
	if v2152 == int32(0) {
		goto L345
	} else {
		goto L346
	}
L344:
	;
	if v2151 != 0 {
		goto L352
	} else {
		goto L353
	}
L345:
	;
	if v2124 == v162 {
		v2171 = v101
		v2172 = l7
		goto L344
	} else {
		goto L348
	}
L346:
	;
	goto L347
L347:
	;
	v2160 = *(*int32)(unsafe.Add(mBase, uint32(v136+v2124<<(uint(int32(2))%32))+20))
	if v2149&int32(_a_F__bt_insertonpg_3) == v2124 {
		goto L349
	} else {
		goto L350
	}
L348:
	;
	goto L347
L349:
	;
	v2167 = v105
	goto L351
L350:
	;
	v2167 = v136 + v2160&int32(_a_F__bt_insertonpg_1)
	goto L351
L351:
	;
	v2171 = v2167
	v2172 = int32(base.Ui32(v2160) >> (uint(int32(17)) % 32))
	goto L344
L352:
	;
	if v2152 != 0 {
		goto L356
	} else {
		goto L357
	}
L353:
	;
	v2199 = v2172
	v2200 = v2171
	goto L354
L354:
	;
	v2203 = F_PageAddItemExtended(m, v2125, v2200, v2199, int32(1), int32(0))
	mBase = m.M
	v2204 = m.ExcPending
	if v2204 != 0 {
		goto L20
	} else {
		goto L364
	}
L355:
	;
	v2193 = F__bt_truncate(m, l0, v2192, v2171, l2)
	mBase = m.M
	v2194 = m.ExcPending
	if v2194 != 0 {
		goto L20
	} else {
		goto L363
	}
L356:
	;
	if v102&int32(_a_F__bt_insertonpg_3) == v2124 {
		v2192 = v101
		goto L355
	} else {
		goto L359
	}
L357:
	;
	goto L358
L358:
	;
	v2178 = int32(_a_F__bt_insertonpg_3)
	v2179 = (v2124 - int32(1)) & v2178
	v2183 = *(*int32)(unsafe.Add(mBase, uint32(v136+v2179<<(uint(int32(2))%32))+20))
	if v2149&v2178 == v2179 {
		goto L360
	} else {
		goto L361
	}
L359:
	;
	goto L358
L360:
	;
	v2190 = v105
	goto L362
L361:
	;
	v2190 = v136 + v2183&int32(_a_F__bt_insertonpg_1)
	goto L362
L362:
	;
	v2192 = v2190
	goto L355
L363:
	;
	v2195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2193)+6)))
	v2199 = v2195 & int32(_a_F__bt_insertonpg_6)
	v2200 = v2193
	goto L354
L364:
	;
	if v2203 == int32(0) {
		goto L13
	} else {
		goto L365
	}
L365:
	;
	v2207 = F__bt_allocbuf(m, l0, l1)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L20
	} else {
		goto L367
	}
L366:
	;
	if v2207 < int32(0) {
		goto L372
	} else {
		goto L373
	}
L367:
	;
	if v2207 < int32(0) {
		goto L368
	} else {
		goto L369
	}
L368:
	;
	v2212 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(v2212+(v2207^int32(-1))<<(uint(int32(2))%32))))
	v2226 = v2218
	goto L366
L369:
	;
	goto L370
L370:
	;
	v2220 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2226 = v2220 + v2207<<(uint(int32(13))%32) + int32(-8192)
	goto L366
L371:
	;
	v2246 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2226)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v2131)+4)) = v2245
	v2250 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[4]))
	v2254 = F_LWLockAcquire(m, v2250+int32(2560), int32(1))
	mBase = m.M
	v2255 = m.ExcPending
	if v2255 != 0 {
		goto L20
	} else {
		goto L375
	}
L372:
	;
	v2230 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v2236 = *(*int32)(unsafe.Add(mBase, uint32(v2230+(v2207^int32(-1))<<(uint(int32(6))%32))+16))
	v2245 = v2236
	goto L371
L373:
	;
	goto L374
L374:
	;
	v2238 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v2244 = *(*int32)(unsafe.Add(mBase, uint32(v2238+v2207<<(uint(int32(6))%32)+int32(-64))+16))
	v2245 = v2244
	goto L371
L375:
	;
	v2256 = int32(0)
	v2258 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[5]))
	v2259 = *(*int32)(unsafe.Add(mBase, uint32(v2258)+4))
	if v2259 <= v2256 {
		v2368 = v2256
		goto L376
	} else {
		goto L377
	}
L376:
	;
	v2370 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[4]))
	F_LWLockRelease(m, v2370+int32(2560))
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L20
	} else {
		goto L384
	}
L377:
	;
	v2264 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v2267 = int32(0)
	goto L378
L378:
	;
	v2312 = v2258 + int32(12) + v2267*int32(12)
	v2313 = *(*int32)(unsafe.Add(mBase, uint32(v2312)))
	if v2313 != v2264 {
		goto L380
	} else {
		goto L381
	}
L379:
	;
	v2368 = int32(0)
	goto L376
L380:
	;
	v2320 = v2267 + int32(1)
	if v2320 != v2259 {
		v2267 = v2320
		goto L378
	} else {
		goto L383
	}
L381:
	;
	v2315 = *(*int32)(unsafe.Add(mBase, uint32(v2312)+4))
	v2316 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v2315 != v2316 {
		goto L380
	} else {
		goto L382
	}
L382:
	;
	v2318 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2312)+8)))
	v2368 = v2318
	goto L376
L383:
	;
	goto L379
L384:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2131)+14)) = uint16(v2368)
	v2376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v2377 = v2246 + v2226
	*(*int32)(unsafe.Add(mBase, uint32(v2377))) = v160
	v2380 = v2376 & int32(_a_F__bt_insertonpg_14)
	*(*uint16)(unsafe.Add(mBase, uint32(v2377)+12)) = uint16(v2380)
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2377)+4)) = v2382
	v2384 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2377)+8)) = v2384
	v2386 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2131)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2377)+14)) = uint16(v2386)
	if v139 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v2390 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v2398 = F_PageAddItemExtended(m, v2226, v136+v2390&int32(_a_F__bt_insertonpg_1), int32(base.Ui32(v2390)>>(uint(int32(17))%32)), int32(1), int32(0))
	mBase = m.M
	v2399 = m.ExcPending
	if v2399 != 0 {
		goto L20
	} else {
		goto L388
	}
L386:
	;
	v2403 = int32(1)
	goto L387
L387:
	;
	if v2151 != 0 {
		goto L390
	} else {
		goto L391
	}
L388:
	;
	if v2398 == int32(0) {
		goto L12
	} else {
		goto L389
	}
L389:
	;
	v2403 = int32(2)
	goto L387
L390:
	;
	v2405 = int32(0)
	goto L392
L391:
	;
	v2405 = v2403
	goto L392
L392:
	;
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v2408 != 0 {
		goto L393
	} else {
		goto L394
	}
L393:
	;
	v2409 = int32(2)
	goto L395
L394:
	;
	v2409 = int32(1)
	goto L395
L395:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v141) {
		goto L396
	} else {
		goto L397
	}
L396:
	;
	v2417 = int32(base.Ui32(v141+int32(_a_F__bt_insertonpg_4)) >> (uint(int32(2)) % 32))
	goto L398
L397:
	;
	v2417 = int32(0)
	goto L398
L398:
	;
	v2419 = v2417 & int32(_a_F__bt_insertonpg_3)
	if base.Ui32(v2409) <= base.Ui32(v2419) {
		goto L399
	} else {
		goto L400
	}
L399:
	;
	v2434 = v2409
	v2438 = int32(2)
	v2450 = v2403
	goto L402
L400:
	;
	v2570 = v2409
	v2586 = v2403
	goto L401
L401:
	;
	v2605 = int32(_a_F__bt_insertonpg_3)
	if base.Ui32(v2570&v2605) <= base.Ui32(v102&v2605) {
		goto L431
	} else {
		goto L432
	}
L402:
	;
	v2469 = int32(_a_F__bt_insertonpg_3)
	v2470 = v2434 & v2469
	v2474 = *(*int32)(unsafe.Add(mBase, uint32(v136+int32(20)+v2470<<(uint(int32(2))%32))))
	if v2470 == v2149&v2469 {
		goto L405
	} else {
		goto L406
	}
L403:
	;
	v2570 = v2556
	v2586 = v2554
	goto L401
L404:
	;
	v2521 = int32(base.Ui32(v2474) >> (uint(int32(17)) % 32))
	if base.Ui32(v2470) < base.Ui32(v2124) {
		goto L420
	} else {
		goto L421
	}
L405:
	;
	v2516 = v105
	v2518 = v2438
	v2519 = v2450
	goto L404
L406:
	;
	goto L407
L407:
	;
	v2480 = v136 + v2474&int32(_a_F__bt_insertonpg_1)
	if v2470 != v102&int32(_a_F__bt_insertonpg_3) {
		v2516 = v2480
		v2518 = v2438
		v2519 = v2450
		goto L404
	} else {
		goto L408
	}
L408:
	;
	v2484 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+222)))
	if v2484 == int32(1) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v2490 = F_PageAddItemExtended(m, v2125, v101, l7, v2438&int32(_a_F__bt_insertonpg_3), int32(0))
	mBase = m.M
	v2491 = m.ExcPending
	if v2491 != 0 {
		goto L20
	} else {
		goto L412
	}
L410:
	;
	goto L411
L411:
	;
	v2497 = v2450 & int32(_a_F__bt_insertonpg_3)
	if v2497 == v2405 {
		goto L414
	} else {
		goto L415
	}
L412:
	;
	if v2490 == int32(0) {
		goto L11
	} else {
		goto L413
	}
L413:
	;
	v2516 = v2480
	v2518 = v2438 + int32(1)
	v2519 = v2450
	goto L404
L414:
	;
	v2499 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+192)) = v2499
	*(*int32)(unsafe.Add(mBase, uint32(v48)+196)) = int32(537395200)
	v2506 = int32(8)
	v2507 = v48 + int32(192)
	goto L416
L415:
	;
	v2506 = l7
	v2507 = v101
	goto L416
L416:
	;
	v2509 = F_PageAddItemExtended(m, v2226, v2507, v2506, v2497, int32(0))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L20
	} else {
		goto L417
	}
L417:
	;
	if v2509 == int32(0) {
		goto L10
	} else {
		goto L418
	}
L418:
	;
	v2516 = v2480
	v2518 = v2438
	v2519 = v2450 + int32(1)
	goto L404
L419:
	;
	v2556 = v2434 + int32(1)
	if base.Ui32(v2556&int32(_a_F__bt_insertonpg_3)) <= base.Ui32(v2419) {
		v2434 = v2556
		v2438 = v2553
		v2450 = v2554
		goto L402
	} else {
		goto L430
	}
L420:
	;
	v2526 = F_PageAddItemExtended(m, v2125, v2516, v2521, v2518&int32(_a_F__bt_insertonpg_3), int32(0))
	mBase = m.M
	v2527 = m.ExcPending
	if v2527 != 0 {
		goto L20
	} else {
		goto L423
	}
L421:
	;
	goto L422
L422:
	;
	v2533 = v2519 & int32(_a_F__bt_insertonpg_3)
	if v2533 == v2405 {
		goto L425
	} else {
		goto L426
	}
L423:
	;
	if v2526 == int32(0) {
		goto L9
	} else {
		goto L424
	}
L424:
	;
	v2553 = v2518 + int32(1)
	v2554 = v2519
	goto L419
L425:
	;
	v2535 = *(*int64)(unsafe.Add(mBase, uint32(v2516)))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+192)) = v2535
	*(*int32)(unsafe.Add(mBase, uint32(v48)+196)) = int32(537395200)
	v2542 = int32(8)
	v2543 = v48 + int32(192)
	goto L427
L426:
	;
	v2542 = v2521
	v2543 = v2516
	goto L427
L427:
	;
	v2545 = F_PageAddItemExtended(m, v2226, v2543, v2542, v2533, int32(0))
	mBase = m.M
	v2546 = m.ExcPending
	if v2546 != 0 {
		goto L20
	} else {
		goto L428
	}
L428:
	;
	if v2545 == int32(0) {
		goto L8
	} else {
		goto L429
	}
L429:
	;
	v2553 = v2518
	v2554 = v2519 + int32(1)
	goto L419
L430:
	;
	goto L403
L431:
	;
	v2611 = v2586 & int32(_a_F__bt_insertonpg_3)
	if v2611 == v2405 {
		goto L434
	} else {
		goto L435
	}
L432:
	;
	goto L433
L433:
	;
	if v139 == int32(0) {
		goto L440
	} else {
		goto L441
	}
L434:
	;
	v2613 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+192)) = v2613
	*(*int32)(unsafe.Add(mBase, uint32(v48)+196)) = int32(537395200)
	v2620 = int32(8)
	v2621 = v48 + int32(192)
	goto L436
L435:
	;
	v2620 = l7
	v2621 = v101
	goto L436
L436:
	;
	v2623 = F_PageAddItemExtended(m, v2226, v2621, v2620, v2611, int32(0))
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L20
	} else {
		goto L437
	}
L437:
	;
	if v2623 == int32(0) {
		goto L7
	} else {
		goto L438
	}
L438:
	;
	goto L433
L439:
	;
	v2670 = int32(_a_F__bt_insertonpg_15)
	v2672 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6])) = v2672 + int32(1)
	F_PageRestoreTempPage(m, v2125, v136)
	mBase = m.M
	v2677 = m.ExcPending
	if v2677 != 0 {
		goto L20
	} else {
		goto L450
	}
L440:
	;
	v2631 = int32(0)
	v2667 = v2631
	v2668 = v2631
	v2669 = v2631
	goto L439
L441:
	;
	goto L442
L442:
	;
	v2634 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v2636 = F__bt_getbuf(m, l0, v2634, int32(2))
	mBase = m.M
	v2637 = m.ExcPending
	if v2637 != 0 {
		goto L20
	} else {
		goto L444
	}
L443:
	;
	v2656 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2655)+16)))
	v2657 = v2656 + v2655
	v2658 = *(*int32)(unsafe.Add(mBase, uint32(v2657)))
	if v2658 != v160 {
		goto L6
	} else {
		goto L448
	}
L444:
	;
	if v2636 < int32(0) {
		goto L445
	} else {
		goto L446
	}
L445:
	;
	v2641 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2647 = *(*int32)(unsafe.Add(mBase, uint32(v2641+(v2636^int32(-1))<<(uint(int32(2))%32))))
	v2655 = v2647
	goto L443
L446:
	;
	goto L447
L447:
	;
	v2649 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2655 = v2649 + v2636<<(uint(int32(13))%32) + int32(-8192)
	goto L443
L448:
	;
	v2660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2657)+14)))
	v2661 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2377)+14)))
	if v2660 == v2661 {
		v2667 = v2636
		v2668 = v2655
		v2669 = v2657
		goto L439
	} else {
		goto L449
	}
L449:
	;
	v2663 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2377)+12)))
	v2665 = v2663 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v2377)+12)) = uint16(v2665)
	v2667 = v2636
	v2668 = v2655
	v2669 = v2657
	goto L439
L450:
	;
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v2679 = m.ExcPending
	if v2679 != 0 {
		goto L20
	} else {
		goto L451
	}
L451:
	;
	F_MarkBufferDirty(m, v2207)
	mBase = m.M
	v2681 = m.ExcPending
	if v2681 != 0 {
		goto L20
	} else {
		goto L452
	}
L452:
	;
	if v139 != 0 {
		goto L453
	} else {
		goto L454
	}
L453:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2669))) = v2245
	F_MarkBufferDirty(m, v2667)
	mBase = m.M
	v2684 = m.ExcPending
	if v2684 != 0 {
		goto L20
	} else {
		goto L456
	}
L454:
	;
	goto L455
L455:
	;
	if v2151 == int32(0) {
		goto L457
	} else {
		goto L458
	}
L456:
	;
	goto L455
L457:
	;
	if l4 < int32(0) {
		goto L461
	} else {
		goto L462
	}
L458:
	;
	goto L459
L459:
	;
	v2714 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2715 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2714)+118)))
	if v2715 != int32(112) {
		v2836 = v2200
		goto L465
	} else {
		goto L466
	}
L460:
	;
	v2705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2704)+16)))
	v2706 = v2705 + v2704
	v2707 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2706)+12)))
	v2709 = v2707 & int32(_a_F__bt_insertonpg_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v2706)+12)) = uint16(v2709)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v2712 = m.ExcPending
	if v2712 != 0 {
		goto L20
	} else {
		goto L464
	}
L461:
	;
	v2690 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2696 = *(*int32)(unsafe.Add(mBase, uint32(v2690+(l4^int32(-1))<<(uint(int32(2))%32))))
	v2704 = v2696
	goto L460
L462:
	;
	goto L463
L463:
	;
	v2698 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2704 = v2698 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L460
L464:
	;
	goto L459
L465:
	;
	v2838 = int32(_a_F__bt_insertonpg_15)
	v2840 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6])) = v2840 - int32(1)
	if v139 != 0 {
		goto L517
	} else {
		goto L518
	}
L466:
	;
	v2719 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[7]))
	if v2719 <= int32(0) {
		goto L467
	} else {
		goto L468
	}
L467:
	;
	v2722 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2722 != 0 {
		v2836 = v2200
		goto L465
	} else {
		goto L470
	}
L468:
	;
	goto L469
L469:
	;
	v2724 = *(*int32)(unsafe.Add(mBase, uint32(v2377)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+198)) = uint16(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+192)) = v2724
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+196)) = uint16(v2124)
	if base.Ui32(v2149&int32(_a_F__bt_insertonpg_3)) < base.Ui32(v2124) {
		goto L472
	} else {
		goto L473
	}
L470:
	;
	v2723 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2723 != 0 {
		v2836 = v2200
		goto L465
	} else {
		goto L471
	}
L471:
	;
	goto L469
L472:
	;
	v2732 = v10
	goto L474
L473:
	;
	v2732 = int32(0)
	goto L474
L474:
	;
	if v10&int32(_a_F__bt_insertonpg_3) != 0 {
		goto L475
	} else {
		goto L476
	}
L475:
	;
	v2736 = v2732
	goto L477
L476:
	;
	v2736 = int32(0)
	goto L477
L477:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+200)) = uint16(v2736)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2739 = m.ExcPending
	if v2739 != 0 {
		goto L20
	} else {
		goto L478
	}
L478:
	;
	F_XLogRegisterData(m, v48+int32(192), int32(10))
	mBase = m.M
	v2744 = m.ExcPending
	if v2744 != 0 {
		goto L20
	} else {
		goto L479
	}
L479:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v2748 = m.ExcPending
	if v2748 != 0 {
		goto L20
	} else {
		goto L480
	}
L480:
	;
	F_XLogRegisterBuffer(m, int32(1), v2207, int32(6))
	mBase = m.M
	v2752 = m.ExcPending
	if v2752 != 0 {
		goto L20
	} else {
		goto L481
	}
L481:
	;
	if v139 != 0 {
		goto L482
	} else {
		goto L483
	}
L482:
	;
	F_XLogRegisterBuffer(m, int32(2), v2667, int32(8))
	mBase = m.M
	v2756 = m.ExcPending
	if v2756 != 0 {
		goto L20
	} else {
		goto L485
	}
L483:
	;
	goto L484
L484:
	;
	if v2151 == int32(0) {
		goto L486
	} else {
		goto L487
	}
L485:
	;
	goto L484
L486:
	;
	F_XLogRegisterBuffer(m, int32(3), l4, int32(8))
	mBase = m.M
	v2762 = m.ExcPending
	if v2762 != 0 {
		goto L20
	} else {
		goto L489
	}
L487:
	;
	goto L488
L488:
	;
	v2763 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+222)))
	v2764 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+200)))
	if v2763|v2764 != 0 {
		goto L490
	} else {
		goto L491
	}
L489:
	;
	goto L488
L490:
	;
	if v2764 != 0 {
		goto L493
	} else {
		goto L494
	}
L491:
	;
	goto L492
L492:
	;
	if v2151 == int32(0) {
		goto L500
	} else {
		goto L501
	}
L493:
	;
	v2767 = v106
	goto L495
L494:
	;
	v2767 = v101
	goto L495
L495:
	;
	if v2763 != 0 {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	v2768 = v2767
	goto L498
L497:
	;
	v2768 = v106
	goto L498
L498:
	;
	F_XLogRegisterBufData(m, int32(0), v2768, l7)
	mBase = m.M
	v2770 = m.ExcPending
	if v2770 != 0 {
		goto L20
	} else {
		goto L499
	}
L499:
	;
	goto L492
L500:
	;
	v2773 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v2777 = v136 + v2773&int32(_a_F__bt_insertonpg_1)
	goto L502
L501:
	;
	v2777 = v2200
	goto L502
L502:
	;
	v2779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2777)+6)))
	F_XLogRegisterBufData(m, int32(0), v2777, (v2779&int32(_a_F__bt_insertonpg_6)+int32(7))&int32(_a_F__bt_insertonpg_17))
	mBase = m.M
	v2787 = m.ExcPending
	if v2787 != 0 {
		goto L20
	} else {
		goto L503
	}
L503:
	;
	v2789 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2226)+14)))
	v2791 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2226)+16)))
	F_XLogRegisterBufData(m, int32(1), v2226+v2789, v2791-v2789)
	mBase = m.M
	v2794 = m.ExcPending
	if v2794 != 0 {
		goto L20
	} else {
		goto L504
	}
L504:
	;
	v2798 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+222)))
	if v2798 != 0 {
		goto L505
	} else {
		goto L506
	}
L505:
	;
	v2799 = int32(48)
	goto L507
L506:
	;
	v2799 = int32(64)
	goto L507
L507:
	;
	v2800 = F_XLogInsert(m, int32(11), v2799)
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L20
	} else {
		goto L508
	}
L508:
	;
	v2802 = int64(32)
	*(*int64)(unsafe.Add(mBase, uint32(v136))) = base.I64_rotr(v2800, v2802)
	v2805 = base.I32_wrap_i64(v2800)
	*(*int32)(unsafe.Add(mBase, uint32(v2226)+4)) = v2805
	v2809 = base.I32_wrap_i64(int64(base.Ui64(v2800) >> (uint(v2802) % 64)))
	*(*int32)(unsafe.Add(mBase, uint32(v2226))) = v2809
	if v139 != 0 {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2668)+4)) = v2805
	*(*int32)(unsafe.Add(mBase, uint32(v2668))) = v2809
	goto L511
L510:
	;
	goto L511
L511:
	;
	if v2151 != 0 {
		v2836 = v2777
		goto L465
	} else {
		goto L512
	}
L512:
	;
	if l4 < int32(0) {
		goto L514
	} else {
		goto L515
	}
L513:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2830)+4)) = v2805
	*(*int32)(unsafe.Add(mBase, uint32(v2830))) = v2809
	v2836 = v2777
	goto L465
L514:
	;
	v2816 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2822 = *(*int32)(unsafe.Add(mBase, uint32(v2816+(l4^int32(-1))<<(uint(int32(2))%32))))
	v2830 = v2822
	goto L513
L515:
	;
	goto L516
L516:
	;
	v2824 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2830 = v2824 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L513
L517:
	;
	F__bt_relbuf(m, v2667)
	mBase = m.M
	v2845 = m.ExcPending
	if v2845 != 0 {
		goto L20
	} else {
		goto L520
	}
L518:
	;
	goto L519
L519:
	;
	v2847 = int32(0)
	if v2151 == v2847 {
		goto L522
	} else {
		goto L523
	}
L520:
	;
	goto L519
L521:
	;
	if l3 < int32(0) {
		goto L528
	} else {
		goto L529
	}
L522:
	;
	F__bt_relbuf(m, l4)
	mBase = m.M
	v2852 = m.ExcPending
	if v2852 != 0 {
		goto L20
	} else {
		goto L525
	}
L523:
	;
	goto L524
L524:
	;
	F_pfree(m, v2836)
	mBase = m.M
	v2854 = m.ExcPending
	if v2854 != 0 {
		goto L20
	} else {
		goto L526
	}
L525:
	;
	goto L521
L526:
	;
	goto L521
L527:
	;
	if v2207 < int32(0) {
		goto L532
	} else {
		goto L533
	}
L528:
	;
	v2858 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v2864 = *(*int32)(unsafe.Add(mBase, uint32(v2858+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v2873 = v2864
	goto L527
L529:
	;
	goto L530
L530:
	;
	v2866 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v2872 = *(*int32)(unsafe.Add(mBase, uint32(v2866+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v2873 = v2872
	goto L527
L531:
	;
	F_PredicateLockPageSplit(m, l0, v2873, v2892)
	mBase = m.M
	v2894 = m.ExcPending
	if v2894 != 0 {
		goto L20
	} else {
		goto L535
	}
L532:
	;
	v2877 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v2883 = *(*int32)(unsafe.Add(mBase, uint32(v2877+(v2207^int32(-1))<<(uint(int32(6))%32))+16))
	v2892 = v2883
	goto L531
L533:
	;
	goto L534
L534:
	;
	v2885 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v2891 = *(*int32)(unsafe.Add(mBase, uint32(v2885+v2207<<(uint(int32(6))%32)+int32(-64))+16))
	v2892 = v2891
	goto L531
L535:
	;
	F__bt_insert_parent(m, l0, l1, l3, v2207, l5, base.B2i32(v108 != int32(0)), base.B2i32(v71|v70 == v2847))
	mBase = m.M
	v2898 = m.ExcPending
	if v2898 != 0 {
		goto L20
	} else {
		goto L536
	}
L536:
	;
	goto L23
L537:
	;
	v2936 = int32(_a_F__bt_insertonpg_15)
	v2938 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6])) = v2938 + int32(1)
	if v10 == int32(0) {
		goto L546
	} else {
		goto L547
	}
L538:
	;
	v2905 = F__bt_getbuf(m, l0, int32(0), int32(2))
	mBase = m.M
	v2906 = m.ExcPending
	if v2906 != 0 {
		goto L20
	} else {
		goto L540
	}
L539:
	;
	v2926 = v2924 + int32(24)
	v2927 = *(*int32)(unsafe.Add(mBase, uint32(v2924)+44))
	v2928 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if base.Ui32(v2927) < base.Ui32(v2928) {
		v2933 = v2905
		v2934 = v2926
		v2935 = v2924
		goto L537
	} else {
		goto L544
	}
L540:
	;
	if v2905 < int32(0) {
		goto L541
	} else {
		goto L542
	}
L541:
	;
	v2910 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2916 = *(*int32)(unsafe.Add(mBase, uint32(v2910+(v2905^int32(-1))<<(uint(int32(2))%32))))
	v2924 = v2916
	goto L539
L542:
	;
	goto L543
L543:
	;
	v2918 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2924 = v2918 + v2905<<(uint(int32(13))%32) + int32(-8192)
	goto L539
L544:
	;
	F__bt_relbuf(m, v2905)
	mBase = m.M
	v2931 = m.ExcPending
	if v2931 != 0 {
		goto L20
	} else {
		goto L545
	}
L545:
	;
	v2933 = int32(0)
	v2934 = v2926
	v2935 = v2924
	goto L537
L546:
	;
	v2958 = F_PageAddItemExtended(m, v67, v101, l7, v102&int32(_a_F__bt_insertonpg_3), int32(0))
	mBase = m.M
	v2959 = m.ExcPending
	if v2959 != 0 {
		goto L20
	} else {
		goto L549
	}
L547:
	;
	v2944 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
	v2950 = (v2944&int32(_a_F__bt_insertonpg_6) + int32(7)) & int32(_a_F__bt_insertonpg_17)
	if v2950 == int32(0) {
		goto L546
	} else {
		goto L548
	}
L548:
	;
	base.MemoryCopy(m, v104, v105, v2950)
	goto L546
L549:
	;
	if v2958 == int32(0) {
		goto L5
	} else {
		goto L550
	}
L550:
	;
	v2963 = v72 & int32(1)
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v2965 = m.ExcPending
	if v2965 != 0 {
		goto L20
	} else {
		goto L551
	}
L551:
	;
	if v2933 != 0 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	v2966 = *(*int32)(unsafe.Add(mBase, uint32(v2934)+4))
	if base.Ui32(v2966) <= base.Ui32(int32(2)) {
		goto L555
	} else {
		goto L556
	}
L553:
	;
	goto L554
L554:
	;
	if v2963 == int32(0) {
		goto L564
	} else {
		goto L565
	}
L555:
	;
	v2969 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2935)+64)) = uint8(v2969)
	*(*int64)(unsafe.Add(mBase, uint32(v2935)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v2935)+48)) = v2969
	*(*int32)(unsafe.Add(mBase, uint32(v2935)+28)) = int32(3)
	v2977 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v2935)+12)) = uint16(v2977)
	goto L558
L556:
	;
	goto L557
L557:
	;
	if l3 < int32(0) {
		goto L560
	} else {
		goto L561
	}
L558:
	;
	goto L557
L559:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2934)+16)) = v2997
	v2999 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2934)+20)) = v2999
	F_MarkBufferDirty(m, v2933)
	mBase = m.M
	v3002 = m.ExcPending
	if v3002 != 0 {
		goto L20
	} else {
		goto L563
	}
L560:
	;
	v2982 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v2988 = *(*int32)(unsafe.Add(mBase, uint32(v2982+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v2997 = v2988
	goto L559
L561:
	;
	goto L562
L562:
	;
	v2990 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v2996 = *(*int32)(unsafe.Add(mBase, uint32(v2990+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v2997 = v2996
	goto L559
L563:
	;
	goto L554
L564:
	;
	if l4 < int32(0) {
		goto L568
	} else {
		goto L569
	}
L565:
	;
	goto L566
L566:
	;
	v3032 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3033 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3032)+118)))
	if v3033 != int32(112) {
		goto L572
	} else {
		goto L573
	}
L567:
	;
	v3023 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3022)+16)))
	v3024 = v3023 + v3022
	v3025 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3024)+12)))
	v3027 = v3025 & int32(_a_F__bt_insertonpg_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v3024)+12)) = uint16(v3027)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v3030 = m.ExcPending
	if v3030 != 0 {
		goto L20
	} else {
		goto L571
	}
L568:
	;
	v3008 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(v3008+(l4^int32(-1))<<(uint(int32(2))%32))))
	v3022 = v3014
	goto L567
L569:
	;
	goto L570
L570:
	;
	v3016 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v3022 = v3016 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L567
L571:
	;
	goto L566
L572:
	;
	v3153 = int32(_a_F__bt_insertonpg_15)
	v3155 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6])) = v3155 - int32(1)
	if v2933 != 0 {
		goto L615
	} else {
		goto L616
	}
L573:
	;
	v3037 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[7]))
	if v3037 <= int32(0) {
		goto L574
	} else {
		goto L575
	}
L574:
	;
	v3040 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v3040 != 0 {
		goto L572
	} else {
		goto L577
	}
L575:
	;
	goto L576
L576:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+222)) = uint16(v102)
	F_XLogBeginInsert(m)
	mBase = m.M
	v3044 = m.ExcPending
	if v3044 != 0 {
		goto L20
	} else {
		goto L579
	}
L577:
	;
	v3041 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v3041 != 0 {
		goto L572
	} else {
		goto L578
	}
L578:
	;
	goto L576
L579:
	;
	F_XLogRegisterData(m, v48+int32(222), int32(2))
	mBase = m.M
	v3049 = m.ExcPending
	if v3049 != 0 {
		goto L20
	} else {
		goto L580
	}
L580:
	;
	if v2963|v10 == int32(0) {
		goto L582
	} else {
		goto L583
	}
L581:
	;
	v3110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3108)+6)))
	F_XLogRegisterBufData(m, int32(0), v3108, v3110&int32(_a_F__bt_insertonpg_6))
	mBase = m.M
	v3114 = m.ExcPending
	if v3114 != 0 {
		goto L20
	} else {
		goto L603
	}
L582:
	;
	F_XLogRegisterBuffer(m, int32(1), l4, int32(8))
	mBase = m.M
	v3056 = m.ExcPending
	if v3056 != 0 {
		goto L20
	} else {
		goto L585
	}
L583:
	;
	goto L584
L584:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v3091 = m.ExcPending
	if v3091 != 0 {
		goto L20
	} else {
		goto L592
	}
L585:
	;
	if v2933 != 0 {
		goto L586
	} else {
		goto L587
	}
L586:
	;
	v3057 = *(*int32)(unsafe.Add(mBase, uint32(v2934)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+192)) = v3057
	v3059 = *(*int32)(unsafe.Add(mBase, uint32(v2934)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+196)) = v3059
	v3061 = *(*int32)(unsafe.Add(mBase, uint32(v2934)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+200)) = v3061
	v3063 = *(*int32)(unsafe.Add(mBase, uint32(v2934)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+204)) = v3063
	v3065 = *(*int32)(unsafe.Add(mBase, uint32(v2934)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+208)) = v3065
	v3067 = *(*int32)(unsafe.Add(mBase, uint32(v2934)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+212)) = v3067
	v3069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2934)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+216)) = uint8(v3069)
	F_XLogRegisterBuffer(m, int32(2), v2933, int32(14))
	mBase = m.M
	v3074 = m.ExcPending
	if v3074 != 0 {
		goto L20
	} else {
		goto L589
	}
L587:
	;
	v3083 = int32(16)
	goto L588
L588:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v3087 = m.ExcPending
	if v3087 != 0 {
		goto L20
	} else {
		goto L591
	}
L589:
	;
	F_XLogRegisterBufData(m, int32(2), v48+int32(192), int32(28))
	mBase = m.M
	v3080 = m.ExcPending
	if v3080 != 0 {
		goto L20
	} else {
		goto L590
	}
L590:
	;
	v3083 = int32(32)
	goto L588
L591:
	;
	v3107 = v3083
	v3108 = v101
	goto L581
L592:
	;
	if v10 == int32(0) {
		goto L593
	} else {
		goto L594
	}
L593:
	;
	v3095 = int32(80)
	if v2963 != 0 {
		goto L596
	} else {
		goto L597
	}
L594:
	;
	goto L595
L595:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+192)) = uint16(v10)
	F_XLogRegisterBufData(m, int32(0), v48+int32(192), int32(2))
	mBase = m.M
	v3106 = m.ExcPending
	if v3106 != 0 {
		goto L20
	} else {
		goto L602
	}
L596:
	;
	v3098 = int32(0)
	goto L598
L597:
	;
	v3098 = v3095
	goto L598
L598:
	;
	if v10 != 0 {
		goto L599
	} else {
		goto L600
	}
L599:
	;
	v3099 = v3095
	goto L601
L600:
	;
	v3099 = v3098
	goto L601
L601:
	;
	v3107 = v3099
	v3108 = v101
	goto L581
L602:
	;
	v3107 = int32(80)
	v3108 = v106
	goto L581
L603:
	;
	v3116 = F_XLogInsert(m, int32(11), v3107)
	mBase = m.M
	v3117 = m.ExcPending
	if v3117 != 0 {
		goto L20
	} else {
		goto L604
	}
L604:
	;
	if v2933 != 0 {
		goto L605
	} else {
		goto L606
	}
L605:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2935))) = base.I64_rotr(v3116, int64(32))
	goto L607
L606:
	;
	goto L607
L607:
	;
	if v2963 == int32(0) {
		goto L608
	} else {
		goto L609
	}
L608:
	;
	if l4 < int32(0) {
		goto L612
	} else {
		goto L613
	}
L609:
	;
	goto L610
L610:
	;
	*(*uint32)(unsafe.Add(mBase, uint32(v67)+4)) = uint32(v3116)
	*(*int32)(unsafe.Add(mBase, uint32(v67))) = base.I32_wrap_i64(int64(base.Ui64(v3116) >> (uint(int64(32)) % 64)))
	goto L572
L611:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v3140))) = base.I64_rotr(v3116, int64(32))
	goto L610
L612:
	;
	v3126 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v3132 = *(*int32)(unsafe.Add(mBase, uint32(v3126+(l4^int32(-1))<<(uint(int32(2))%32))))
	v3140 = v3132
	goto L611
L613:
	;
	goto L614
L614:
	;
	v3134 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v3140 = v3134 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L611
L615:
	;
	F__bt_relbuf(m, v2933)
	mBase = m.M
	v3160 = m.ExcPending
	if v3160 != 0 {
		goto L20
	} else {
		goto L618
	}
L616:
	;
	goto L617
L617:
	;
	if v2963 == int32(0) {
		goto L619
	} else {
		goto L620
	}
L618:
	;
	goto L617
L619:
	;
	F__bt_relbuf(m, l4)
	mBase = m.M
	v3164 = m.ExcPending
	if v3164 != 0 {
		goto L20
	} else {
		goto L622
	}
L620:
	;
	goto L621
L621:
	;
	if v71|v108 != 0 {
		goto L624
	} else {
		goto L625
	}
L622:
	;
	F__bt_relbuf(m, l3)
	mBase = m.M
	v3166 = m.ExcPending
	if v3166 != 0 {
		goto L20
	} else {
		goto L623
	}
L623:
	;
	goto L23
L624:
	;
	F__bt_relbuf(m, l3)
	mBase = m.M
	v3169 = m.ExcPending
	if v3169 != 0 {
		goto L20
	} else {
		goto L627
	}
L625:
	;
	goto L626
L626:
	;
	if l3 < int32(0) {
		goto L629
	} else {
		goto L630
	}
L627:
	;
	goto L23
L628:
	;
	F__bt_relbuf(m, l3)
	mBase = m.M
	v3190 = m.ExcPending
	if v3190 != 0 {
		goto L20
	} else {
		goto L632
	}
L629:
	;
	v3173 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v3179 = *(*int32)(unsafe.Add(mBase, uint32(v3173+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v3188 = v3179
	goto L628
L630:
	;
	goto L631
L631:
	;
	v3181 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v3187 = *(*int32)(unsafe.Add(mBase, uint32(v3181+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v3188 = v3187
	goto L628
L632:
	;
	if v3188 == int32(-1) {
		goto L23
	} else {
		goto L633
	}
L633:
	;
	v3193 = F__bt_getrootheight(m, l0)
	mBase = m.M
	v3194 = m.ExcPending
	if v3194 != 0 {
		goto L20
	} else {
		goto L634
	}
L634:
	;
	if v3193 < int32(2) {
		goto L23
	} else {
		goto L635
	}
L635:
	;
	v3197 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3197 != 0 {
		goto L636
	} else {
		goto L637
	}
L636:
	;
	v3223 = v3197
	goto L638
L637:
	;
	v3198 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3199 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+152)) = v3199
	v3201 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+144)) = v3201
	v3205 = F_smgropen(m, v48+int32(144), v3198)
	mBase = m.M
	v3206 = m.ExcPending
	if v3206 != 0 {
		goto L20
	} else {
		goto L639
	}
L638:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3223)+16)) = v3188
	goto L23
L639:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3205
	v3209 = *(*int32)(unsafe.Add(mBase, uint32(v3205)+72))
	if v3209 != 0 {
		goto L641
	} else {
		goto L642
	}
L640:
	;
	v3221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3223 = v3221
	goto L638
L641:
	;
	v3217 = v3209
	goto L643
L642:
	;
	v3210 = *(*int32)(unsafe.Add(mBase, uint32(v3205)+76))
	v3211 = *(*int32)(unsafe.Add(mBase, uint32(v3205)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v3210)+4)) = v3211
	v3213 = *(*int32)(unsafe.Add(mBase, uint32(v3205)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3211))) = v3213
	v3215 = *(*int32)(unsafe.Add(mBase, uint32(v3205)+72))
	v3217 = v3215
	goto L643
L643:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3205)+72)) = v3217 + int32(1)
	goto L640
L644:
	;
	F_pfree(m, v105)
	mBase = m.M
	v3271 = m.ExcPending
	if v3271 != 0 {
		goto L20
	} else {
		goto L647
	}
L645:
	;
	goto L646
L646:
	;
	m.G0 = v48 + int32(224)
	return
L647:
	;
	F_pfree(m, v101)
	mBase = m.M
	v3273 = m.ExcPending
	if v3273 != 0 {
		goto L20
	} else {
		goto L648
	}
L648:
	;
	goto L646
L649:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3283 = m.ExcPending
	if v3283 != 0 {
		goto L20
	} else {
		goto L650
	}
L650:
	;
	v3284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+2)))
	v3285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6))))
	v3286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)))
	if l3 < int32(0) {
		goto L652
	} else {
		goto L653
	}
L651:
	;
	v3306 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+176)) = v3306 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+172)) = v3305
	*(*int32)(unsafe.Add(mBase, uint32(v48)+168)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v48)+164)) = v3286
	*(*int32)(unsafe.Add(mBase, uint32(v48)+160)) = v3284 | v3285<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_18), v48+int32(160))
	mBase = m.M
	v3321 = m.ExcPending
	if v3321 != 0 {
		goto L20
	} else {
		goto L655
	}
L652:
	;
	v3290 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v3296 = *(*int32)(unsafe.Add(mBase, uint32(v3290+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v3305 = v3296
	goto L651
L653:
	;
	goto L654
L654:
	;
	v3298 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v3304 = *(*int32)(unsafe.Add(mBase, uint32(v3298+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v3305 = v3304
	goto L651
L655:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1191), int32(_a_F__bt_insertonpg_20))
	mBase = m.M
	v3326 = m.ExcPending
	if v3326 != 0 {
		goto L20
	} else {
		goto L656
	}
L656:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L657:
	;
	v3331 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v3331 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_21), v48)
	mBase = m.M
	v3338 = m.ExcPending
	if v3338 != 0 {
		goto L20
	} else {
		goto L658
	}
L658:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1704), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3343 = m.ExcPending
	if v3343 != 0 {
		goto L20
	} else {
		goto L659
	}
L659:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L660:
	;
	v3351 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+112)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+116)) = v3351 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_23), v48+int32(112))
	mBase = m.M
	v3360 = m.ExcPending
	if v3360 != 0 {
		goto L20
	} else {
		goto L661
	}
L661:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1773), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3365 = m.ExcPending
	if v3365 != 0 {
		goto L20
	} else {
		goto L662
	}
L662:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L663:
	;
	v3373 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+80)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+84)) = v3373 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_24), v48+int32(80))
	mBase = m.M
	v3382 = m.ExcPending
	if v3382 != 0 {
		goto L20
	} else {
		goto L664
	}
L664:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1821), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3387 = m.ExcPending
	if v3387 != 0 {
		goto L20
	} else {
		goto L665
	}
L665:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L666:
	;
	v3395 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+96)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+100)) = v3395 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_25), v48+int32(96))
	mBase = m.M
	v3404 = m.ExcPending
	if v3404 != 0 {
		goto L20
	} else {
		goto L667
	}
L667:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1834), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3409 = m.ExcPending
	if v3409 != 0 {
		goto L20
	} else {
		goto L668
	}
L668:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L669:
	;
	v3417 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+48)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+52)) = v3417 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_26), v48+int32(48))
	mBase = m.M
	v3426 = m.ExcPending
	if v3426 != 0 {
		goto L20
	} else {
		goto L670
	}
L670:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1848), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3431 = m.ExcPending
	if v3431 != 0 {
		goto L20
	} else {
		goto L671
	}
L671:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L672:
	;
	v3439 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+64)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+68)) = v3439 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_27), v48-int32(-64))
	mBase = m.M
	v3448 = m.ExcPending
	if v3448 != 0 {
		goto L20
	} else {
		goto L673
	}
L673:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1860), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3453 = m.ExcPending
	if v3453 != 0 {
		goto L20
	} else {
		goto L674
	}
L674:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L675:
	;
	v3461 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = v3461 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_25), v48+int32(32))
	mBase = m.M
	v3470 = m.ExcPending
	if v3470 != 0 {
		goto L20
	} else {
		goto L676
	}
L676:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1881), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3475 = m.ExcPending
	if v3475 != 0 {
		goto L20
	} else {
		goto L677
	}
L677:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L678:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3485 = m.ExcPending
	if v3485 != 0 {
		goto L20
	} else {
		goto L679
	}
L679:
	;
	v3486 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3487 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v3488 = *(*int32)(unsafe.Add(mBase, uint32(v2657)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v3488
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v3487
	*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v3486 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_28), v48+int32(16))
	mBase = m.M
	v3499 = m.ExcPending
	if v3499 != 0 {
		goto L20
	} else {
		goto L680
	}
L680:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1904), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3504 = m.ExcPending
	if v3504 != 0 {
		goto L20
	} else {
		goto L681
	}
L681:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L682:
	;
	if l3 < int32(0) {
		goto L684
	} else {
		goto L685
	}
L683:
	;
	v3528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+128)) = v3527
	*(*int32)(unsafe.Add(mBase, uint32(v48)+132)) = v3528 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_29), v48+int32(128))
	mBase = m.M
	v3537 = m.ExcPending
	if v3537 != 0 {
		goto L20
	} else {
		goto L687
	}
L684:
	;
	v3512 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v3518 = *(*int32)(unsafe.Add(mBase, uint32(v3512+(l3^int32(-1))<<(uint(int32(6))%32))+16))
	v3527 = v3518
	goto L683
L685:
	;
	goto L686
L686:
	;
	v3520 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v3526 = *(*int32)(unsafe.Add(mBase, uint32(v3520+l3<<(uint(int32(6))%32)+int32(-64))+16))
	v3527 = v3526
	goto L683
L687:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_19), int32(1283), int32(_a_F__bt_insertonpg_20))
	mBase = m.M
	v3542 = m.ExcPending
	if v3542 != 0 {
		goto L20
	} else {
		goto L688
	}
L688:
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v18 = int32(*(*int16)(unsafe.Add(mBase, uint32(v17)+10)))
	if v18 <= int32(0) {
		v68 = v16
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v14 + int32(16)
	return v68
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
	v68 = v18 + int32(1)
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
		v68 = v29
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
	v54 = v23 + int32(4) + v29<<(uint(int32(4))%32)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+6)))
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+4)))
	v57 = F_datum_image_eq(m, v39, v45, v55, v56)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
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
	if v57 == int32(0) {
		v68 = v29
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
	if v73&int32(_a_F__bt_parallel_primscan_schedule_0) != 0 {
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
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
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
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)+32))
	if v6 == int32(1) {
		F__bt_buildadd(m, l0, l1, v5, int32(0))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return
		} else {
			v30 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v30
			*(*int64)(unsafe.Add(mBase, uint32(l2)+28)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v30
			return
		}
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)+24))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l2)+28))
		v14 = F__bt_form_posting(m, v5, v12, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return
		} else {
			v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+6)))
			v19 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+2)))
			v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14))))
			F__bt_buildadd(m, l0, l1, v14, v16&int32(_a_F__bt_sort_dedup_finish_pending_0)-(v19|v20<<(uint(int32(16))%32)))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				F_pfree(m, v14)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return
				} else {
					v30 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v30
					*(*int64)(unsafe.Add(mBase, uint32(l2)+28)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v30
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
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
				if v55 != 0 {
					F_ReleaseBuffer(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
						if l1 == int32(1) {
							v64 = int32(68)
						} else {
							v64 = int32(64)
						}
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if l1 != v68 {
							v70 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
						} else {
						}
						v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							return v73
						}
					}
				} else {
					if l1 == int32(1) {
						v64 = int32(68)
					} else {
						v64 = int32(64)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
					if l1 != v68 {
						v70 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
					} else {
					}
					v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						return v73
					}
				}
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
				if v17 != 0 {
					F_IncrBufferRefCount(m, v17)
					mBase = m.M
					v19 = m.ExcPending
					if v19 != 0 {
						return int32(0)
					} else {
						v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
						v24 = v20*int32(10) + int32(58)
						if v24 != 0 {
							base.MemoryCopy(m, v6+int32(_a_F__bt_steppage_0), v6+int32(56), v24)
						} else {
						}
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
						if v30 == int32(0) {
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
							if v33 == int32(0) {
							} else {
								v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
								base.MemoryCopy(m, v30, v36, v33)
							}
						}
						v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
						*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[0]))) = v39
						*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
						v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)))
						if v43 != int32(1) {
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							if v46 == int32(1) {
								v49 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[1]))) = uint8(v49)
							} else {
								v51 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[2]))) = uint8(v51)
							}
						}
						v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
						if v55 != 0 {
							F_ReleaseBuffer(m, v55)
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
								if l1 == int32(1) {
									v64 = int32(68)
								} else {
									v64 = int32(64)
								}
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
								v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
								if l1 != v68 {
									v70 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
								} else {
								}
								v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
								mBase = m.M
								v74 = m.ExcPending
								if v74 != 0 {
									return int32(0)
								} else {
									return v73
								}
							}
						} else {
							if l1 == int32(1) {
								v64 = int32(68)
							} else {
								v64 = int32(64)
							}
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							if l1 != v68 {
								v70 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
							} else {
							}
							v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								return v73
							}
						}
					}
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
					v24 = v20*int32(10) + int32(58)
					if v24 != 0 {
						base.MemoryCopy(m, v6+int32(_a_F__bt_steppage_0), v6+int32(56), v24)
					} else {
					}
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
					if v30 == int32(0) {
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
						if v33 == int32(0) {
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							base.MemoryCopy(m, v30, v36, v33)
						}
					}
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[0]))) = v39
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)))
					if v43 != int32(1) {
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if v46 == int32(1) {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[1]))) = uint8(v49)
						} else {
							v51 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[2]))) = uint8(v51)
						}
					}
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
					if v55 != 0 {
						F_ReleaseBuffer(m, v55)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
							if l1 == int32(1) {
								v64 = int32(68)
							} else {
								v64 = int32(64)
							}
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							if l1 != v68 {
								v70 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
							} else {
							}
							v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								return v73
							}
						}
					} else {
						if l1 == int32(1) {
							v64 = int32(68)
						} else {
							v64 = int32(64)
						}
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if l1 != v68 {
							v70 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
						} else {
						}
						v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							return v73
						}
					}
				}
			}
		}
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
		if v14 < int32(0) {
			v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
			if v55 != 0 {
				F_ReleaseBuffer(m, v55)
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
					if l1 == int32(1) {
						v64 = int32(68)
					} else {
						v64 = int32(64)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
					if l1 != v68 {
						v70 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
					} else {
					}
					v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						return v73
					}
				}
			} else {
				if l1 == int32(1) {
					v64 = int32(68)
				} else {
					v64 = int32(64)
				}
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
				v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
				v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
				if l1 != v68 {
					v70 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
				} else {
				}
				v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
				mBase = m.M
				v74 = m.ExcPending
				if v74 != 0 {
					return int32(0)
				} else {
					return v73
				}
			}
		} else {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
			if v17 != 0 {
				F_IncrBufferRefCount(m, v17)
				mBase = m.M
				v19 = m.ExcPending
				if v19 != 0 {
					return int32(0)
				} else {
					v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
					v24 = v20*int32(10) + int32(58)
					if v24 != 0 {
						base.MemoryCopy(m, v6+int32(_a_F__bt_steppage_0), v6+int32(56), v24)
					} else {
					}
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
					if v30 == int32(0) {
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
						if v33 == int32(0) {
						} else {
							v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
							base.MemoryCopy(m, v30, v36, v33)
						}
					}
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
					*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[0]))) = v39
					*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
					v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)))
					if v43 != int32(1) {
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if v46 == int32(1) {
							v49 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[1]))) = uint8(v49)
						} else {
							v51 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[2]))) = uint8(v51)
						}
					}
					v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
					if v55 != 0 {
						F_ReleaseBuffer(m, v55)
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
							if l1 == int32(1) {
								v64 = int32(68)
							} else {
								v64 = int32(64)
							}
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
							if l1 != v68 {
								v70 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
							} else {
							}
							v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
							mBase = m.M
							v74 = m.ExcPending
							if v74 != 0 {
								return int32(0)
							} else {
								return v73
							}
						}
					} else {
						if l1 == int32(1) {
							v64 = int32(68)
						} else {
							v64 = int32(64)
						}
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if l1 != v68 {
							v70 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
						} else {
						}
						v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							return v73
						}
					}
				}
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v6)+96))
				v24 = v20*int32(10) + int32(58)
				if v24 != 0 {
					base.MemoryCopy(m, v6+int32(_a_F__bt_steppage_0), v6+int32(56), v24)
				} else {
				}
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
				if v30 == int32(0) {
				} else {
					v33 = *(*int32)(unsafe.Add(mBase, uint32(v6)+84))
					if v33 == int32(0) {
					} else {
						v36 = *(*int32)(unsafe.Add(mBase, uint32(v6)+44))
						base.MemoryCopy(m, v30, v36, v33)
					}
				}
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v6)+52))
				*(*int32)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[0]))) = v39
				*(*int32)(unsafe.Add(mBase, uint32(v6)+52)) = int32(-1)
				v43 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)))
				if v43 != int32(1) {
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
					if v46 == int32(1) {
						v49 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[1]))) = uint8(v49)
					} else {
						v51 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+uint32(_c_F__bt_steppage[2]))) = uint8(v51)
					}
				}
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v6)+56))
				if v55 != 0 {
					F_ReleaseBuffer(m, v55)
					mBase = m.M
					v57 = m.ExcPending
					if v57 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v6)+56)) = int32(0)
						if l1 == int32(1) {
							v64 = int32(68)
						} else {
							v64 = int32(64)
						}
						v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
						v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
						v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
						if l1 != v68 {
							v70 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
						} else {
						}
						v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
						mBase = m.M
						v74 = m.ExcPending
						if v74 != 0 {
							return int32(0)
						} else {
							return v73
						}
					}
				} else {
					if l1 == int32(1) {
						v64 = int32(68)
					} else {
						v64 = int32(64)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v6+v64)))
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v6)+80))
					if l1 != v68 {
						v70 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v6)+17)) = uint8(v70)
					} else {
					}
					v73 = F__bt_readnextpage(m, l0, v66, v67, l1, int32(0))
					mBase = m.M
					v74 = m.ExcPending
					if v74 != 0 {
						return int32(0)
					} else {
						return v73
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
func F_bt_tuple_present_callback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
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
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v67 int64
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v13 = F_index_form_tuple(m, v12, l2, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v15)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		*(*int32)(unsafe.Add(mBase, uint32(v13))) = v17
		v19 = F_bt_normalize_tuple(m, l5, v13)
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return
		} else {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+6)))
			v25 = F_bloom_lacks_element(m, v21, v19, v22&int32(_a_F_bt_tuple_present_callback_0))
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return
			} else {
				if v25 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return
					} else {
						F_errcode(m, int32(16779816))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return
						} else {
							v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
							v35 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
							v36 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
							v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+48))
							v38 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
							v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+48))
							v40 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
							*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v40
							v42 = int32(4)
							*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v39 + v42
							*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v37 + v42
							*(*int32)(unsafe.Add(mBase, uint32(v10))) = v34 | v35<<(uint(int32(16))%32)
							F_errmsg(m, int32(_a_F_bt_tuple_present_callback_1), v10)
							mBase = m.M
							v54 = m.ExcPending
							if v54 != 0 {
								return
							} else {
								v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+9)))
								if v55 == int32(0) {
									F_errhint(m, int32(_a_F_bt_tuple_present_callback_2), int32(0))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										F_errfinish(m, int32(_a_F_bt_tuple_present_callback_3), int32(2808), int32(_a_F_bt_tuple_present_callback_4))
										mBase = m.M
										v66 = m.ExcPending
										if v66 != 0 {
											return
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								} else {
									F_errfinish(m, int32(_a_F_bt_tuple_present_callback_3), int32(2808), int32(_a_F_bt_tuple_present_callback_4))
									mBase = m.M
									v66 = m.ExcPending
									if v66 != 0 {
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
				} else {
					v67 = *(*int64)(unsafe.Add(mBase, uint32(l5)+64))
					*(*int64)(unsafe.Add(mBase, uint32(l5)+64)) = v67 + int64(1)
					F_pfree(m, v13)
					mBase = m.M
					v72 = m.ExcPending
					if v72 != 0 {
						return
					} else {
						if v13 != v19 {
							F_pfree(m, v19)
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return
							} else {
								m.G0 = v10 + int32(16)
								return
							}
						} else {
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		}
	}
}
