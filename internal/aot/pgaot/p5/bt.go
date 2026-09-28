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
	var v23 int32
	_ = v23
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
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
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
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v135 int64
	_ = v135
	var v136 int32
	_ = v136
	var v140 int64
	_ = v140
	var v141 int32
	_ = v141
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v305 int32
	_ = v305
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	v15 = m.G0
	v17 = v15 - int32(16)
	m.G0 = v17
	v19 = int32(1)
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v21 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l2)+16)))
	v22 = l2 + v21
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v22)+12)))
	if v23&v19 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v17 + int32(16)
	return v325
L2:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v30 != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2+l3<<(uint(int32(2))%32))+20))
	v39 = l2 + v36&int32(_a_F__bt_compare_0)
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+7)))
	if v40&int32(32) == int32(0) {
		goto L11
	} else {
		goto L12
	}
L5:
	;
	v31 = int32(2)
	goto L7
L6:
	;
	v31 = int32(1)
	goto L7
L7:
	;
	if v31 == l3 {
		v325 = v19
		goto L1
	} else {
		goto L8
	}
L8:
	;
	goto L4
L9:
	;
	v325 = int32(1)
	goto L1
L10:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	if v54 < v55 {
		goto L14
	} else {
		goto L15
	}
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v52 = int32(*(*int16)(unsafe.Add(mBase, uint32(v51)+8)))
	v54 = v52
	goto L10
L12:
	;
	v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
	if v45&int32(_a_F__bt_compare_1) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v54 = v45 & int32(4095)
	goto L10
L14:
	;
	v57 = v54
	goto L16
L15:
	;
	v57 = v55
	goto L16
L16:
	;
	if int32(0) < v57 {
		goto L17
	} else {
		goto L18
	}
L17:
	;
	v61 = v39 + int32(8)
	v67 = int32(1)
	v69 = l1 + int32(16)
	goto L20
L18:
	;
	v188 = v55
	goto L19
L19:
	;
	v200 = int32(1)
	if v54 < v188 {
		v325 = v200
		goto L1
	} else {
		goto L61
	}
L20:
	;
	v81 = int32(*(*int16)(unsafe.Add(mBase, uint32(v69)+4)))
	v82 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+6)))
	if int32(0) <= v82 {
		goto L27
	} else {
		goto L28
	}
L21:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v188 = v185
	goto L19
L22:
	;
	if v67 != v57 {
		v67 = v67 + int32(1)
		v69 = v69 + int32(56)
		goto L20
	} else {
		goto L60
	}
L23:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v69)+12))
	v162 = *(*int64)(unsafe.Add(mBase, uint32(v69)+48))
	v163 = F_FunctionCall2Coll(m, v69+int32(16), v161, v140, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L40
	} else {
		goto L54
	}
L24:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v151&int32(1) != 0 {
		goto L22
	} else {
		goto L50
	}
L25:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v69)))
	if v141&int32(1) == int32(0) {
		goto L23
	} else {
		goto L46
	}
L26:
	;
	v135 = F_nocache_index_getattr(m, v39, v81, v20)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L40
	} else {
		goto L45
	}
L27:
	;
	v87 = v20 + int32(20) + v81<<(uint(int32(3))%32)
	v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87))))
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
	v120 = int32(1)
	v121 = v81 - v120
	v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v121>>(uint(int32(3))%32)))))
	if int32(base.Ui32(v125)>>(uint(v121&int32(7))%32))&v120 == int32(0) {
		goto L24
	} else {
		goto L44
	}
L30:
	;
	v91 = v61 + v88
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v87)+4)))
	if v92 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v95 = int32(*(*int16)(unsafe.Add(mBase, uint32(v87)+2)))
	if base.I32_popcnt(v95) != int32(1) {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v140 = base.I64_extend_i32_u(v91)
	goto L25
L34:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L40
	} else {
		goto L41
	}
L35:
	;
	switch base.I32_ctz(v95) {
	case 0:
		goto L39
	case 1:
		goto L38
	case 2:
		goto L37
	case 3:
		goto L36
	default:
		goto L34
	}
L36:
	;
	v103 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
	v140 = v103
	goto L25
L37:
	;
	v102 = int64(*(*int32)(unsafe.Add(mBase, uint32(v91))))
	v140 = v102
	goto L25
L38:
	;
	v101 = int64(*(*int16)(unsafe.Add(mBase, uint32(v91))))
	v140 = v101
	goto L25
L39:
	;
	v100 = int64(*(*int8)(unsafe.Add(mBase, uint32(v91))))
	v140 = v100
	goto L25
L40:
	;
	return int32(0)
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v95
	F_errmsg_internal(m, int32(_a_F__bt_compare_2), v17)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L40
	} else {
		goto L42
	}
L42:
	;
	F_errfinish(m, int32(_a_F__bt_compare_3), int32(123), int32(_a_F__bt_compare_4))
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L40
	} else {
		goto L43
	}
L43:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L44:
	;
	goto L26
L45:
	;
	v140 = v135
	goto L25
L46:
	;
	if v141&int32(33554432) != 0 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v150 = int32(-1)
	goto L49
L48:
	;
	v150 = int32(1)
	goto L49
L49:
	;
	v325 = v150
	goto L1
L50:
	;
	if v151&int32(33554432) != 0 {
		goto L51
	} else {
		goto L52
	}
L51:
	;
	v158 = int32(1)
	goto L53
L52:
	;
	v158 = int32(-1)
	goto L53
L53:
	;
	v325 = v158
	goto L1
L54:
	;
	v165 = base.I32_wrap_i64(v163)
	v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v69)+3)))
	if v166&int32(1) == int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	if v165 < int32(0) {
		goto L9
	} else {
		goto L58
	}
L56:
	;
	v175 = v165
	goto L57
L57:
	;
	if v175 != 0 {
		v325 = v175
		goto L1
	} else {
		goto L59
	}
L58:
	;
	v175 = int32(0) - v165
	goto L57
L59:
	;
	goto L22
L60:
	;
	goto L21
L61:
	;
	v202 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+6)))
	if v202&int32(_a_F__bt_compare_1) == int32(0) {
		v228 = v39
		goto L64
	} else {
		goto L65
	}
L62:
	;
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+4)))
	if v310|(v312|base.B2i32(v188 != v54)) == int32(0) {
		goto L85
	} else {
		goto L86
	}
L63:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v308 != 0 {
		v325 = v200
		goto L1
	} else {
		goto L84
	}
L64:
	;
	v229 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v229 == int32(0) {
		v310 = v228
		goto L62
	} else {
		goto L70
	}
L65:
	;
	v207 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
	if v207&int32(_a_F__bt_compare_1) == int32(0) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	if v207&int32(_a_F__bt_compare_5) == int32(0) {
		goto L63
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v221 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+2)))
	v222 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	v228 = v221 + (v39 + v222<<(uint(int32(16))%32))
	goto L64
L69:
	;
	v228 = v39 + v202&int32(_a_F__bt_compare_6) - int32(6)
	goto L64
L70:
	;
	v235 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229)+2)))
	v236 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229))))
	v237 = int32(16)
	v239 = v235 | v236<<(uint(v237)%32)
	v240 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+2)))
	v241 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228))))
	v244 = v240 | v241<<(uint(v237)%32)
	if base.Ui32(v239) < base.Ui32(v244) {
		v255 = int32(-1)
		goto L72
	} else {
		goto L73
	}
L71:
	;
	if v255 <= int32(0) {
		v325 = v255
		goto L1
	} else {
		goto L76
	}
L72:
	;
	goto L71
L73:
	;
	if base.Ui32(v244) < base.Ui32(v239) {
		v255 = int32(1)
		goto L72
	} else {
		goto L74
	}
L74:
	;
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v229)+4)))
	v250 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v228)+4)))
	if base.Ui32(v249) < base.Ui32(v250) {
		v255 = int32(-1)
		goto L72
	} else {
		goto L75
	}
L75:
	;
	v255 = base.B2i32(base.Ui32(v250) < base.Ui32(v249))
	goto L72
L76:
	;
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+7)))
	if v258&int32(32) == int32(0) {
		v325 = v255
		goto L1
	} else {
		goto L77
	}
L77:
	;
	v263 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+4)))
	if v263&int32(_a_F__bt_compare_1) == int32(0) {
		v325 = v255
		goto L1
	} else {
		goto L78
	}
L78:
	;
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39)+2)))
	v270 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v39))))
	v271 = int32(16)
	v277 = int32(6)
	v281 = v269 + (v39 + v270<<(uint(v271)%32)) + v263&int32(4095)*v277 - v277
	v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v268)+2)))
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v268))))
	v289 = v285 | v286<<(uint(v271)%32)
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281)+2)))
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281))))
	v294 = v290 | v291<<(uint(v271)%32)
	if base.Ui32(v289) < base.Ui32(v294) {
		v305 = int32(-1)
		goto L80
	} else {
		goto L81
	}
L79:
	;
	v325 = base.B2i32(int32(0) < v305)
	goto L1
L80:
	;
	goto L79
L81:
	;
	if base.Ui32(v294) < base.Ui32(v289) {
		v305 = int32(1)
		goto L80
	} else {
		goto L82
	}
L82:
	;
	v299 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v268)+4)))
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v281)+4)))
	if base.Ui32(v299) < base.Ui32(v300) {
		v305 = int32(-1)
		goto L80
	} else {
		goto L83
	}
L83:
	;
	v305 = base.B2i32(base.Ui32(v300) < base.Ui32(v299))
	goto L80
L84:
	;
	v310 = int32(0)
	goto L62
L85:
	;
	v318 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1))))
	if v318 != 0 {
		v325 = v200
		goto L1
	} else {
		goto L88
	}
L86:
	;
	goto L87
L87:
	;
	v325 = int32(0)
	goto L1
L88:
	;
	goto L87
}
func F__bt_compare_array_elements(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v8 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l1)))
	v10 = F_FunctionCall2Coll(m, v6, v7, v8, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = base.I32_wrap_i64(v10)
		if v14 < int32(0) {
			v18 = int32(1)
		} else {
			v18 = int32(0) - v14
		}
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+8)))
		if v19 != 0 {
			v20 = v18
		} else {
			v20 = v14
		}
		return v20
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
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v161 int32
	_ = v161
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v364 int32
	_ = v364
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v431 int32
	_ = v431
	var v434 int32
	_ = v434
	var v440 int32
	_ = v440
	var v443 int32
	_ = v443
	var v452 int32
	_ = v452
	var v478 int32
	_ = v478
	var v484 int32
	_ = v484
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v516 int32
	_ = v516
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v542 int32
	_ = v542
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v580 int32
	_ = v580
	var v581 int32
	_ = v581
	var v585 int32
	_ = v585
	var v594 int32
	_ = v594
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v637 int32
	_ = v637
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v675 int32
	_ = v675
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v712 int32
	_ = v712
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v735 int32
	_ = v735
	var v742 int32
	_ = v742
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v783 int32
	_ = v783
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v812 int32
	_ = v812
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int64
	_ = v818
	var v827 int32
	_ = v827
	var v833 int32
	_ = v833
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v845 int32
	_ = v845
	var v847 int32
	_ = v847
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v868 int32
	_ = v868
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v883 int32
	_ = v883
	var v889 int32
	_ = v889
	var v914 int32
	_ = v914
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v923 int32
	_ = v923
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v933 int32
	_ = v933
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v960 int32
	_ = v960
	var v966 int32
	_ = v966
	var v976 int32
	_ = v976
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v991 int32
	_ = v991
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1021 int32
	_ = v1021
	var v1029 int32
	_ = v1029
	var v1030 int32
	_ = v1030
	var v1044 int32
	_ = v1044
	var v1051 int32
	_ = v1051
	var v1059 int32
	_ = v1059
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1076 int32
	_ = v1076
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1091 int32
	_ = v1091
	var v1095 int32
	_ = v1095
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1109 int32
	_ = v1109
	var v1111 int32
	_ = v1111
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int32
	_ = v1126
	var v1127 int32
	_ = v1127
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1137 int32
	_ = v1137
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1145 int32
	_ = v1145
	var v1148 int32
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1180 int32
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1191 int32
	_ = v1191
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1201 int32
	_ = v1201
	var v1212 int32
	_ = v1212
	var v1213 int32
	_ = v1213
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1221 int32
	_ = v1221
	var v1223 int32
	_ = v1223
	var v1225 int32
	_ = v1225
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1244 int32
	_ = v1244
	var v1253 int32
	_ = v1253
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1269 int32
	_ = v1269
	var v1274 int32
	_ = v1274
	var v1299 int32
	_ = v1299
	var v1302 int32
	_ = v1302
	var v1306 int32
	_ = v1306
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1321 int32
	_ = v1321
	var v1322 int32
	_ = v1322
	var v1327 int32
	_ = v1327
	var v1330 int32
	_ = v1330
	var v1331 int32
	_ = v1331
	var v1343 int32
	_ = v1343
	var v1345 int32
	_ = v1345
	var v1346 int32
	_ = v1346
	var v1347 int32
	_ = v1347
	var v1348 int32
	_ = v1348
	var v1355 int32
	_ = v1355
	var v1356 int32
	_ = v1356
	var v1362 int32
	_ = v1362
	var v1368 int32
	_ = v1368
	var v1378 int32
	_ = v1378
	var v1388 int32
	_ = v1388
	var v1420 int32
	_ = v1420
	var v1421 int32
	_ = v1421
	var v1435 int32
	_ = v1435
	var v1442 int32
	_ = v1442
	var v1450 int32
	_ = v1450
	var v1456 int32
	_ = v1456
	var v1460 int32
	_ = v1460
	var v1463 int32
	_ = v1463
	var v1464 int32
	_ = v1464
	var v1467 int32
	_ = v1467
	var v1468 int32
	_ = v1468
	var v1469 int32
	_ = v1469
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1478 int32
	_ = v1478
	var v1479 int32
	_ = v1479
	var v1482 int32
	_ = v1482
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1490 int32
	_ = v1490
	var v1493 int32
	_ = v1493
	var v1496 int32
	_ = v1496
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1504 int32
	_ = v1504
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1514 int32
	_ = v1514
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1521 int32
	_ = v1521
	var v1522 int32
	_ = v1522
	var v1525 int32
	_ = v1525
	var v1527 int32
	_ = v1527
	var v1528 int32
	_ = v1528
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1536 int32
	_ = v1536
	var v1539 int32
	_ = v1539
	var v1544 int32
	_ = v1544
	var v1549 int32
	_ = v1549
	var v1550 int32
	_ = v1550
	var v1555 int32
	_ = v1555
	var v1556 int32
	_ = v1556
	var v1560 int32
	_ = v1560
	var v1561 int32
	_ = v1561
	var v1563 int32
	_ = v1563
	var v1565 int32
	_ = v1565
	var v1567 int32
	_ = v1567
	var v1569 int32
	_ = v1569
	var v1570 int32
	_ = v1570
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1580 int32
	_ = v1580
	var v1582 int32
	_ = v1582
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1603 int32
	_ = v1603
	var v1604 int32
	_ = v1604
	var v1609 int32
	_ = v1609
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1616 int32
	_ = v1616
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1628 int32
	_ = v1628
	var v1635 int32
	_ = v1635
	var v1644 int32
	_ = v1644
	var v1655 int32
	_ = v1655
	var v1656 int32
	_ = v1656
	var v1660 int32
	_ = v1660
	var v1665 int32
	_ = v1665
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1703 int32
	_ = v1703
	var v1704 int32
	_ = v1704
	var v1705 int32
	_ = v1705
	var v1706 int32
	_ = v1706
	var v1707 int32
	_ = v1707
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1714 int32
	_ = v1714
	var v1716 int32
	_ = v1716
	var v1747 int32
	_ = v1747
	var v1748 int32
	_ = v1748
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1755 int32
	_ = v1755
	var v1756 int32
	_ = v1756
	var v1757 int32
	_ = v1757
	var v1759 int32
	_ = v1759
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1766 int32
	_ = v1766
	var v1772 int32
	_ = v1772
	var v1774 int32
	_ = v1774
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1784 int32
	_ = v1784
	var v1785 int32
	_ = v1785
	var v1793 int32
	_ = v1793
	var v1796 int32
	_ = v1796
	var v1797 int32
	_ = v1797
	var v1798 int64
	_ = v1798
	var v1803 int32
	_ = v1803
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1809 int32
	_ = v1809
	var v1810 int32
	_ = v1810
	var v1812 int32
	_ = v1812
	var v1816 int32
	_ = v1816
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1824 int32
	_ = v1824
	var v1834 int32
	_ = v1834
	var v1838 int32
	_ = v1838
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1848 int32
	_ = v1848
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1854 int64
	_ = v1854
	var v1856 int32
	_ = v1856
	var v1857 int32
	_ = v1857
	var v1865 int32
	_ = v1865
	var v1866 int32
	_ = v1866
	var v1877 int32
	_ = v1877
	var v1879 int32
	_ = v1879
	var v1882 int32
	_ = v1882
	var v1886 int32
	_ = v1886
	var v1892 int32
	_ = v1892
	var v1913 int32
	_ = v1913
	var v1917 int32
	_ = v1917
	var v1920 int32
	_ = v1920
	var v1922 int32
	_ = v1922
	var v1925 int32
	_ = v1925
	var v1929 int32
	_ = v1929
	var v1930 int32
	_ = v1930
	var v1932 int32
	_ = v1932
	var v1936 int32
	_ = v1936
	var v1940 int32
	_ = v1940
	var v1942 int32
	_ = v1942
	var v1943 int32
	_ = v1943
	var v1944 int32
	_ = v1944
	var v1945 int32
	_ = v1945
	var v1952 int32
	_ = v1952
	var v1953 int32
	_ = v1953
	var v1959 int32
	_ = v1959
	var v1965 int32
	_ = v1965
	var v1975 int32
	_ = v1975
	var v1980 int32
	_ = v1980
	var v1983 int32
	_ = v1983
	var v1984 int32
	_ = v1984
	var v1985 int32
	_ = v1985
	var v1987 int32
	_ = v1987
	var v1988 int32
	_ = v1988
	var v1993 int32
	_ = v1993
	var v2000 int32
	_ = v2000
	var v2001 int32
	_ = v2001
	var v2006 int32
	_ = v2006
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2011 int32
	_ = v2011
	var v2023 int32
	_ = v2023
	var v2031 int32
	_ = v2031
	var v2032 int32
	_ = v2032
	var v2035 int32
	_ = v2035
	var v2036 int32
	_ = v2036
	var v2039 int32
	_ = v2039
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2050 int32
	_ = v2050
	var v2053 int32
	_ = v2053
	var v2054 int32
	_ = v2054
	var v2065 int32
	_ = v2065
	var v2066 int32
	_ = v2066
	var v2076 int32
	_ = v2076
	var v2079 int32
	_ = v2079
	var v2081 int32
	_ = v2081
	var v2084 int32
	_ = v2084
	var v2087 int32
	_ = v2087
	var v2090 int32
	_ = v2090
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2097 int32
	_ = v2097
	var v2101 int32
	_ = v2101
	var v2105 int32
	_ = v2105
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2117 int32
	_ = v2117
	var v2118 int32
	_ = v2118
	var v2124 int32
	_ = v2124
	var v2130 int32
	_ = v2130
	var v2140 int32
	_ = v2140
	var v2145 int32
	_ = v2145
	var v2151 int32
	_ = v2151
	var v2183 int32
	_ = v2183
	var v2184 int32
	_ = v2184
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2196 int32
	_ = v2196
	var v2199 int32
	_ = v2199
	var v2201 int32
	_ = v2201
	var v2206 int32
	_ = v2206
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2210 int32
	_ = v2210
	var v2214 int32
	_ = v2214
	var v2217 int32
	_ = v2217
	var v2218 int32
	_ = v2218
	var v2219 int32
	_ = v2219
	var v2222 int32
	_ = v2222
	var v2226 int32
	_ = v2226
	var v2231 int32
	_ = v2231
	var v2235 int32
	_ = v2235
	var v2239 int32
	_ = v2239
	var v2242 int64
	_ = v2242
	var v2243 int32
	_ = v2243
	var v2244 int64
	_ = v2244
	var v2245 int32
	_ = v2245
	var v2246 int64
	_ = v2246
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2257 int32
	_ = v2257
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2268 int32
	_ = v2268
	var v2272 int32
	_ = v2272
	var v2277 int32
	_ = v2277
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
	v764 = int32(1)
	if base.B2i32(v742|(l4^v764) != v764)|l3 != 0 {
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
		v742 = l5
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
		v742 = l5
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
	v144 = v115 + int32(1)
	v145 = F_palloc_mul(m, int32(4), v144)
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
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
	v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123)+2)))
	v148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v123))))
	*(*int32)(unsafe.Add(mBase, uint32(v145))) = v147 | v148<<(uint(int32(16))%32)
	v154 = v141 + int32(20)
	v161 = int32(1)
	v168 = v144
	v170 = v145
	v174 = v8
	goto L27
L27:
	;
	v186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v30+v174<<(uint(int32(1))%32)))))
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v154+v186<<(uint(int32(2))%32))))
	v193 = v141 + v190&int32(_a_F__bt_delete_or_dedup_one_page_3)
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v193)+7)))
	if v194&int32(32) != 0 {
		goto L31
	} else {
		goto L32
	}
L28:
	;
	F_pg_qsort(m, v373, v364, int32(4), int32(213))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L25
	} else {
		goto L55
	}
L29:
	;
	v387 = v174 + int32(1)
	if v387 != v115 {
		v161 = v364
		v168 = v371
		v170 = v373
		v174 = v387
		goto L27
	} else {
		goto L54
	}
L30:
	;
	v222 = v197 & int32(4095)
	v223 = v161 + v222
	if v168 < v223 {
		goto L39
	} else {
		goto L40
	}
L31:
	;
	v197 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+4)))
	if v197&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L30
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v202 = v161 + int32(1)
	if v168 < v202 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	v206 = F_repalloc(m, v170, v168<<(uint(int32(3))%32))
	mBase = m.M
	v207 = m.ExcPending
	if v207 != 0 {
		goto L25
	} else {
		goto L38
	}
L36:
	;
	v210 = v168
	v211 = v170
	goto L37
L37:
	;
	v215 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+2)))
	v216 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193))))
	*(*int32)(unsafe.Add(mBase, uint32(v211+v161<<(uint(int32(2))%32)))) = v215 | v216<<(uint(int32(16))%32)
	v364 = v202
	v371 = v210
	v373 = v211
	goto L29
L38:
	;
	v210 = v168 << (uint(int32(1)) % 32)
	v211 = v206
	goto L37
L39:
	;
	v226 = v168 << (uint(int32(1)) % 32)
	if v223 < v226 {
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v233 = v168
	v234 = v170
	goto L41
L41:
	;
	if v222 == int32(0) {
		v364 = v161
		v371 = v233
		v373 = v234
		goto L29
	} else {
		goto L46
	}
L42:
	;
	v228 = v226
	goto L44
L43:
	;
	v228 = v223
	goto L44
L44:
	;
	v231 = F_repalloc(m, v170, v228<<(uint(int32(2))%32))
	mBase = m.M
	v232 = m.ExcPending
	if v232 != 0 {
		goto L25
	} else {
		goto L45
	}
L45:
	;
	v233 = v228
	v234 = v231
	goto L41
L46:
	;
	v237 = int32(0)
	if v222 != int32(1) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v250 = v161
	v253 = v237
	v256 = int32(0)
	goto L50
L48:
	;
	v317 = v161
	v320 = v237
	goto L49
L49:
	;
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+2)))
	v343 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193))))
	v344 = int32(16)
	v350 = v342 + (v193 + v343<<(uint(v344)%32)) + v320*int32(6)
	v351 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350))))
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v350)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v234+v317<<(uint(int32(2))%32)))) = v351<<(uint(v344)%32) | v354
	v364 = v317 + int32(1)
	v371 = v233
	v373 = v234
	goto L29
L50:
	;
	v272 = int32(2)
	v274 = v234 + v250<<(uint(v272)%32)
	v276 = v253 * int32(6)
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+2)))
	v278 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193))))
	v279 = int32(16)
	v283 = v276 + (v277 + (v193 + v278<<(uint(v279)%32)))
	v284 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283))))
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v283)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v274))) = v284<<(uint(v279)%32) | v287
	v290 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193)+2)))
	v291 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v193))))
	v296 = v290 + (v193 + v291<<(uint(v279)%32)) + v276
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v296)+6)))
	v300 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v296)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v274)+4)) = v297<<(uint(v279)%32) | v300
	v304 = v253 + v272
	v306 = v250 + v272
	v308 = v256 + v272
	if v308 != v222&int32(4094) {
		v250 = v306
		v253 = v304
		v256 = v308
		goto L50
	} else {
		goto L52
	}
L51:
	;
	if v222&int32(1) == int32(0) {
		v364 = v306
		v371 = v233
		v373 = v234
		goto L29
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v317 = v306
	v320 = v304
	goto L49
L54:
	;
	goto L28
L55:
	;
	if base.Ui32(int32(2)) <= base.Ui32(v364) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v404 = int32(1)
	v405 = int32(0)
	goto L59
L57:
	;
	v452 = v364
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
	v424 = int32(2)
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v373+v404<<(uint(v424)%32))))
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v373+v405<<(uint(v424)%32))))
	if v427 == v431 {
		v440 = v405
		goto L61
	} else {
		goto L62
	}
L60:
	;
	v452 = v440 + int32(1)
	goto L58
L61:
	;
	v443 = v404 + int32(1)
	if v443 != v364 {
		v404 = v443
		v405 = v440
		goto L59
	} else {
		goto L64
	}
L62:
	;
	v434 = v405 + int32(1)
	if v404 == v434 {
		v440 = v404
		goto L61
	} else {
		goto L63
	}
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v373+v434<<(uint(int32(2))%32)))) = v427
	v440 = v434
	goto L61
L64:
	;
	goto L60
L65:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v30)+832)) = int64(0)
	v496 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v30)+828)) = uint8(v496)
	*(*int32)(unsafe.Add(mBase, uint32(v30)+824)) = v493
	v500 = F_palloc(m, int32(_a_F__bt_delete_or_dedup_one_page_5))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L25
	} else {
		goto L69
	}
L66:
	;
	v478 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[2]))
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v478+(v35^int32(-1))*int32(56))+16))
	v493 = v484
	goto L65
L67:
	;
	goto L68
L68:
	;
	v486 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[3]))
	v487 = int32(56)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v486+v35*v487-v487)+16))
	v493 = v492
	goto L65
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+840)) = v500
	v504 = F_palloc(m, int32(_a_F__bt_delete_or_dedup_one_page_6))
	mBase = m.M
	v505 = m.ExcPending
	if v505 != 0 {
		goto L25
	} else {
		goto L70
	}
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+844)) = v504
	v516 = v57
	goto L71
L71:
	;
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v30)+844))
	v535 = *(*int32)(unsafe.Add(mBase, uint32(v30)+836))
	v538 = v534 + v535*int32(6)
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v30)+840))
	v542 = v539 + v535<<(uint(int32(3))%32)
	v547 = v154 + v516&int32(_a_F__bt_delete_or_dedup_one_page_1)<<(uint(int32(2))%32)
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v551 = v141 + v548&int32(_a_F__bt_delete_or_dedup_one_page_3)
	v552 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v551)+7)))
	if v552&int32(32) != 0 {
		goto L75
	} else {
		goto L76
	}
L72:
	;
	F_pfree(m, v373)
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L25
	} else {
		goto L90
	}
L73:
	;
	v705 = v516 + int32(1)
	v706 = int32(_a_F__bt_delete_or_dedup_one_page_1)
	if base.Ui32(v705&v706) <= base.Ui32(v66&v706) {
		v516 = v705
		goto L71
	} else {
		goto L89
	}
L74:
	;
	v594 = v555 & int32(4095)
	if v594 == int32(0) {
		goto L73
	} else {
		goto L81
	}
L75:
	;
	v555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551)+4)))
	if v555&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L74
	} else {
		goto L78
	}
L76:
	;
	goto L77
L77:
	;
	v559 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551)+2)))
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551))))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+816)) = v559 | v560<<(uint(int32(16))%32)
	v569 = F_bsearch(m, v30+int32(816), v373, v452, int32(4), int32(213))
	mBase = m.M
	v570 = m.ExcPending
	if v570 != 0 {
		goto L25
	} else {
		goto L79
	}
L78:
	;
	goto L77
L79:
	;
	if v569 == int32(0) {
		goto L73
	} else {
		goto L80
	}
L80:
	;
	v573 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v542)+4)) = uint16(v573)
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v551)))
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = v575
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v30)+836))
	*(*uint16)(unsafe.Add(mBase, uint32(v542)+6)) = uint16(v577)
	*(*uint16)(unsafe.Add(mBase, uint32(v538))) = uint16(v516)
	v580 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v581 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v538)+4)) = uint16(v581)
	*(*uint8)(unsafe.Add(mBase, uint32(v538)+3)) = uint8(v581)
	v585 = int32(_a_F__bt_delete_or_dedup_one_page_2)
	*(*uint8)(unsafe.Add(mBase, uint32(v538)+2)) = uint8(base.B2i32(v580&v585 == v585))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+836)) = v577 + int32(1)
	goto L73
L81:
	;
	v605 = v538
	v608 = int32(0)
	v609 = v542
	goto L82
L82:
	;
	v625 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551)+2)))
	v626 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v551))))
	v627 = int32(16)
	v633 = v625 + (v551 + v626<<(uint(v627)%32)) + v608*int32(6)
	v634 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v633))))
	v637 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v633)+2)))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+816)) = v634<<(uint(v627)%32) | v637
	v644 = F_bsearch(m, v30+int32(816), v373, v452, int32(4), int32(213))
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L25
	} else {
		goto L84
	}
L83:
	;
	goto L73
L84:
	;
	if v644 != 0 {
		goto L85
	} else {
		goto L86
	}
L85:
	;
	v646 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v633)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v609)+4)) = uint16(v646)
	v648 = *(*int32)(unsafe.Add(mBase, uint32(v633)))
	*(*int32)(unsafe.Add(mBase, uint32(v609))) = v648
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v30)+836))
	*(*uint16)(unsafe.Add(mBase, uint32(v609)+6)) = uint16(v650)
	*(*uint16)(unsafe.Add(mBase, uint32(v605))) = uint16(v516)
	v653 = *(*int32)(unsafe.Add(mBase, uint32(v547)))
	v654 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v605)+4)) = uint16(v654)
	*(*uint8)(unsafe.Add(mBase, uint32(v605)+3)) = uint8(v654)
	v658 = int32(_a_F__bt_delete_or_dedup_one_page_2)
	*(*uint8)(unsafe.Add(mBase, uint32(v605)+2)) = uint8(base.B2i32(v653&v658 == v658))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v30)+836))
	*(*int32)(unsafe.Add(mBase, uint32(v30)+836)) = v663 + int32(1)
	v671 = v605 + int32(6)
	v672 = v609 + int32(8)
	goto L87
L86:
	;
	v671 = v605
	v672 = v609
	goto L87
L87:
	;
	v675 = v608 + int32(1)
	if v675 != v594 {
		v605 = v671
		v608 = v675
		v609 = v672
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
	v716 = m.ExcPending
	if v716 != 0 {
		goto L25
	} else {
		goto L91
	}
L91:
	;
	v717 = *(*int32)(unsafe.Add(mBase, uint32(v30)+840))
	F_pfree(m, v717)
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L25
	} else {
		goto L92
	}
L92:
	;
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v30)+844))
	F_pfree(m, v720)
	mBase = m.M
	v722 = m.ExcPending
	if v722 != 0 {
		goto L25
	} else {
		goto L93
	}
L93:
	;
	v723 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v723)
	v726 = int32(4)
	v727 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+14)))
	v728 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+12)))
	v729 = v727 - v728
	if v729 <= v726 {
		goto L95
	} else {
		goto L96
	}
L94:
	;
	v735 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v735) <= base.Ui32(v732-int32(4)) {
		goto L1
	} else {
		goto L98
	}
L95:
	;
	v732 = v726
	goto L97
L96:
	;
	v732 = v729
	goto L97
L97:
	;
	goto L94
L98:
	;
	v742 = int32(1)
	goto L2
L99:
	;
	v770 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(l2)+16)) = uint8(v770)
	v772 = v742 | l6
	if v772 == int32(1) {
		goto L100
	} else {
		goto L101
	}
L100:
	;
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v776 = m.G0
	v778 = v776 - int32(32)
	m.G0 = v778
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
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	if v1747 != 0 {
		goto L222
	} else {
		goto L223
	}
L103:
	;
	v798 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v797)+16)))
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v800 = int32(*(*int16)(unsafe.Add(mBase, uint32(v799)+10)))
	v802 = F_palloc(m, int32(1676))
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L25
	} else {
		goto L107
	}
L104:
	;
	v783 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[0]))
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v783+(v35^int32(-1))<<(uint(int32(2))%32))))
	v797 = v789
	goto L103
L105:
	;
	goto L106
L106:
	;
	v791 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[1]))
	v797 = v791 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L103
L107:
	;
	v804 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v802)+20)) = v804
	*(*uint16)(unsafe.Add(mBase, uint32(v802)+16)) = uint16(v804)
	*(*int32)(unsafe.Add(mBase, uint32(v802)+12)) = v804
	*(*int64)(unsafe.Add(mBase, uint32(v802)+4)) = int64(35184372088832)
	v812 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v802))) = uint8(v812)
	v816 = F_palloc(m, int32(_a_F__bt_delete_or_dedup_one_page_4))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L25
	} else {
		goto L108
	}
L108:
	;
	v818 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v802)+28)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v802)+24)) = v816
	*(*int64)(unsafe.Add(mBase, uint32(v802)+36)) = v818
	*(*int32)(unsafe.Add(mBase, uint32(v778)+4)) = l0
	if v35 < int32(0) {
		goto L110
	} else {
		goto L111
	}
L109:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v778)+20)) = int32(0)
	v845 = int32(512)
	v847 = v775 + int32(4)
	if base.Ui32(v847) <= base.Ui32(v845) {
		goto L113
	} else {
		goto L114
	}
L110:
	;
	v827 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[2]))
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v827+(v35^int32(-1))*int32(56))+16))
	v842 = v833
	goto L109
L111:
	;
	goto L112
L112:
	;
	v835 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[3]))
	v836 = int32(56)
	v841 = *(*int32)(unsafe.Add(mBase, uint32(v835+v35*v836-v836)+16))
	v842 = v841
	goto L109
L113:
	;
	v850 = v845
	goto L115
L114:
	;
	v850 = v847
	goto L115
L115:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v778)+16)) = v850
	v852 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v778)+12)) = uint8(v852)
	*(*int32)(unsafe.Add(mBase, uint32(v778)+8)) = v842
	v857 = F_palloc_mul(m, int32(8), int32(1358))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L25
	} else {
		goto L116
	}
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v778)+24)) = v857
	v862 = F_palloc_mul(m, int32(6), int32(1358))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L25
	} else {
		goto L117
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v778)+28)) = v862
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v798+v797)+4))
	if v868 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v869 = int32(2)
	goto L120
L119:
	;
	v869 = int32(1)
	goto L120
L120:
	;
	v870 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v797)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v870) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v878 = int32(base.Ui32(v870+int32(_a_F__bt_delete_or_dedup_one_page_0)) >> (uint(int32(2)) % 32))
	goto L123
L122:
	;
	v878 = int32(0)
	goto L123
L123:
	;
	v880 = v878 & int32(_a_F__bt_delete_or_dedup_one_page_1)
	if base.Ui32(v869) <= base.Ui32(v880) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v883 = v802 + int32(44)
	v889 = v869
	goto L127
L125:
	;
	goto L126
L126:
	;
	v1420 = v778 + int32(4)
	v1421 = int32(0)
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v802)+32))
	if v1435 <= v1421 {
		goto L186
	} else {
		goto L187
	}
L127:
	;
	v914 = v889 & int32(_a_F__bt_delete_or_dedup_one_page_1)
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v797+int32(20)+v914<<(uint(int32(2))%32))))
	v921 = v797 + v918&int32(_a_F__bt_delete_or_dedup_one_page_3)
	if v869 == v914 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L126
L129:
	;
	v1388 = v889 + int32(1)
	if base.Ui32(v1388&int32(_a_F__bt_delete_or_dedup_one_page_1)) <= base.Ui32(v880) {
		v889 = v1388
		goto L127
	} else {
		goto L184
	}
L130:
	;
	v923 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921)+7)))
	if v923&int32(32) != 0 {
		goto L135
	} else {
		goto L136
	}
L131:
	;
	goto L132
L132:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v802)+12))
	v982 = F__bt_keep_natts_fast(m, l0, v981, v921)
	mBase = m.M
	v983 = m.ExcPending
	if v983 != 0 {
		goto L25
	} else {
		goto L146
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v802)+20)) = v960
	*(*uint16)(unsafe.Add(mBase, uint32(v802)+16)) = uint16(v869)
	*(*int32)(unsafe.Add(mBase, uint32(v802)+12)) = v921
	v966 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v802)+36)) = (v966&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) | int32(4)
	v976 = *(*int32)(unsafe.Add(mBase, uint32(v802)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v883+v976<<(uint(int32(2))%32)))) = uint16(v869)
	goto L129
L134:
	;
	v941 = v926 & int32(4095)
	v943 = v941 * int32(6)
	if v943 != 0 {
		goto L139
	} else {
		goto L140
	}
L135:
	;
	v926 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+4)))
	if v926&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L134
	} else {
		goto L138
	}
L136:
	;
	goto L137
L137:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v802)+24))
	v931 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v930)+4)) = uint16(v931)
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v921)))
	*(*int32)(unsafe.Add(mBase, uint32(v930))) = v933
	*(*int32)(unsafe.Add(mBase, uint32(v802)+28)) = int32(1)
	v937 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+6)))
	v960 = v937 & int32(_a_F__bt_delete_or_dedup_one_page_7)
	goto L133
L138:
	;
	goto L137
L139:
	;
	v944 = *(*int32)(unsafe.Add(mBase, uint32(v802)+24))
	v945 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+2)))
	v946 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921))))
	base.MemoryCopy(m, v944, v945+(v921+v946<<(uint(int32(16))%32)), v943)
	goto L141
L140:
	;
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+28)) = v941
	v953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+2)))
	v954 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921))))
	v960 = v953 | v954<<(uint(int32(16))%32)
	goto L133
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v802)+20)) = v1362
	*(*uint16)(unsafe.Add(mBase, uint32(v802)+16)) = uint16(v889)
	*(*int32)(unsafe.Add(mBase, uint32(v802)+12)) = v921
	v1368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v802)+36)) = (v1368&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) | int32(4)
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v802)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v883+v1378<<(uint(int32(2))%32)))) = uint16(v889)
	goto L129
L143:
	;
	v1343 = v1302 & int32(4095)
	v1345 = v1343 * int32(6)
	if v1345 != 0 {
		goto L181
	} else {
		goto L182
	}
L144:
	;
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v802)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v802)+32)) = v1316 + int32(1)
	v1321 = v1004 * int32(6)
	if v1321 != 0 {
		goto L178
	} else {
		goto L179
	}
L145:
	;
	v1029 = v778 + int32(4)
	v1030 = int32(0)
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v802)+32))
	if v1044 <= v1030 {
		goto L154
	} else {
		goto L155
	}
L146:
	;
	if v982 <= v800 {
		goto L145
	} else {
		goto L147
	}
L147:
	;
	v985 = int32(1)
	v986 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921)+7)))
	if v986&int32(32) == int32(0) {
		v1004 = v985
		v1006 = v921
		goto L148
	} else {
		goto L149
	}
L148:
	;
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v802)+8))
	v1008 = *(*int32)(unsafe.Add(mBase, uint32(v802)+20))
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v802)+28))
	if base.Ui32((v1008+(v1009+v1004)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v1007) {
		goto L144
	} else {
		goto L151
	}
L149:
	;
	v991 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+4)))
	if v991&int32(_a_F__bt_delete_or_dedup_one_page_4) == int32(0) {
		v1004 = v985
		v1006 = v921
		goto L148
	} else {
		goto L150
	}
L150:
	;
	v998 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+2)))
	v999 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921))))
	v1004 = v991 & int32(4095)
	v1006 = v998 + (v921 + v999<<(uint(int32(16))%32))
	goto L148
L151:
	;
	if v1009 <= int32(50) {
		goto L145
	} else {
		goto L152
	}
L152:
	;
	v1021 = *(*int32)(unsafe.Add(mBase, uint32(v802)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v802)+4)) = v1021 + int32(1)
	goto L145
L153:
	;
	v1299 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v921)+7)))
	if v1299&int32(32) != 0 {
		goto L174
	} else {
		goto L175
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v802)+28)) = int64(0)
	goto L153
L155:
	;
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+16))
	v1059 = v1051
	v1065 = v1030
	goto L156
L156:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+24))
	v1072 = v1069 + v1059*int32(6)
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+20))
	v1076 = v1073 + v1059<<(uint(int32(3))%32)
	v1077 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v802)+16)))
	v1078 = v1077 + v1065
	v1083 = v797 + int32(20) + v1078&int32(_a_F__bt_delete_or_dedup_one_page_1)<<(uint(int32(2))%32)
	v1084 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	v1087 = v797 + v1084&int32(_a_F__bt_delete_or_dedup_one_page_3)
	v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1087)+7)))
	if v1088&int32(32) != 0 {
		goto L160
	} else {
		goto L161
	}
L157:
	;
	if v1044 <= int32(1) {
		goto L154
	} else {
		goto L173
	}
L158:
	;
	v1264 = v1065 + int32(1)
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v802)+32))
	if v1264 < v1265 {
		v1059 = v1253
		v1065 = v1264
		goto L156
	} else {
		goto L172
	}
L159:
	;
	v1116 = v1091 & int32(4095)
	v1117 = int32(0)
	if int32(2) <= v1044 {
		goto L164
	} else {
		goto L165
	}
L160:
	;
	v1091 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087)+4)))
	if v1091&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L159
	} else {
		goto L163
	}
L161:
	;
	goto L162
L162:
	;
	v1095 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1076)+4)) = uint16(v1095)
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1087)))
	*(*int32)(unsafe.Add(mBase, uint32(v1076))) = v1097
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1076)+6)) = uint16(v1099)
	*(*uint8)(unsafe.Add(mBase, uint32(v1072)+3)) = uint8(base.B2i32(int32(1) < v1044))
	v1102 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1072)+2)) = uint8(v1102)
	*(*uint16)(unsafe.Add(mBase, uint32(v1072))) = uint16(v1078)
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1083)))
	v1109 = int32(base.Ui32(v1105)>>(uint(int32(17))%32)) + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1072)+4)) = uint16(v1109)
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+16))
	v1113 = v1111 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+16)) = v1113
	v1253 = v1113
	goto L158
L163:
	;
	goto L162
L164:
	;
	v1121 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087)+2)))
	v1122 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087))))
	v1123 = int32(16)
	v1126 = v1121 + (v1087 + v1122<<(uint(v1123)%32))
	v1127 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1126))))
	v1130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1126)+2)))
	v1131 = v1127<<(uint(v1123)%32) | v1130
	v1134 = int32(6)
	v1136 = v1126 + int32(base.Ui32(v1116)>>(uint(int32(1))%32))*v1134
	v1137 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1136))))
	v1140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1136)+2)))
	v1141 = v1137<<(uint(v1123)%32) | v1140
	v1145 = v1126 + v1116*v1134
	v1148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145-v1134))))
	v1153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1145-int32(4)))))
	v1158 = base.B2i32(v1141 == v1131)
	v1159 = base.B2i32(v1131 != v1141) & base.B2i32(v1141 == v1148<<(uint(v1123)%32)|v1153)
	goto L166
L165:
	;
	v1158 = v1117
	v1159 = v1117
	goto L166
L166:
	;
	if v1116 == int32(0) {
		v1253 = v1059
		goto L158
	} else {
		goto L167
	}
L167:
	;
	v1164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087)+2)))
	v1165 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087))))
	v1169 = v1164 + (v1087 + v1165<<(uint(int32(16))%32))
	v1170 = *(*int32)(unsafe.Add(mBase, uint32(v1169)))
	*(*int32)(unsafe.Add(mBase, uint32(v1076))) = v1170
	v1172 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1169)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1076)+4)) = uint16(v1172)
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1076)+6)) = uint16(v1174)
	v1176 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v1072)+4)) = uint16(v1176)
	v1178 = int32(1)
	v1179 = v1116 - v1178
	v1180 = int32(0)
	v1183 = v1158 | v1159&base.B2i32(v1179 == v1180)
	*(*uint8)(unsafe.Add(mBase, uint32(v1072)+3)) = uint8(v1183)
	*(*uint8)(unsafe.Add(mBase, uint32(v1072)+2)) = uint8(v1180)
	*(*uint16)(unsafe.Add(mBase, uint32(v1072))) = uint16(v1078)
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+16))
	v1191 = v1189 + v1178
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+16)) = v1191
	if v1116 == v1178 {
		v1253 = v1191
		goto L158
	} else {
		goto L168
	}
L168:
	;
	v1198 = v1072
	v1199 = v1178
	v1201 = v1076
	goto L169
L169:
	;
	v1212 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087)+2)))
	v1213 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1087))))
	v1218 = int32(6)
	v1220 = v1212 + (v1087 + v1213<<(uint(int32(16))%32)) + v1199*v1218
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v1220)))
	*(*int32)(unsafe.Add(mBase, uint32(v1201)+8)) = v1221
	v1223 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1220)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1201)+12)) = uint16(v1223)
	v1225 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1201)+14)) = uint16(v1225)
	*(*uint16)(unsafe.Add(mBase, uint32(v1198)+10)) = uint16(v1218)
	v1230 = v1159 & base.B2i32(v1199 == v1179)
	*(*uint8)(unsafe.Add(mBase, uint32(v1198)+9)) = uint8(v1230)
	v1232 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1198)+8)) = uint8(v1232)
	*(*uint16)(unsafe.Add(mBase, uint32(v1198)+6)) = uint16(v1078)
	v1235 = *(*int32)(unsafe.Add(mBase, uint32(v1029)+16))
	v1236 = int32(1)
	v1237 = v1235 + v1236
	*(*int32)(unsafe.Add(mBase, uint32(v1029)+16)) = v1237
	v1244 = v1199 + v1236
	if v1244 != v1116 {
		v1198 = v1198 + v1218
		v1199 = v1244
		v1201 = v1201 + int32(8)
		goto L169
	} else {
		goto L171
	}
L170:
	;
	v1253 = v1237
	goto L158
L171:
	;
	goto L170
L172:
	;
	goto L157
L173:
	;
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v802)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v802+v1269<<(uint(int32(2))%32))+46)) = uint16(v1265)
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v802)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v802)+40)) = v1274 + int32(1)
	goto L154
L174:
	;
	v1302 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+4)))
	if v1302&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L143
	} else {
		goto L177
	}
L175:
	;
	goto L176
L176:
	;
	v1306 = *(*int32)(unsafe.Add(mBase, uint32(v802)+24))
	v1307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1306)+4)) = uint16(v1307)
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(v921)))
	*(*int32)(unsafe.Add(mBase, uint32(v1306))) = v1309
	*(*int32)(unsafe.Add(mBase, uint32(v802)+28)) = int32(1)
	v1313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+6)))
	v1362 = v1313 & int32(_a_F__bt_delete_or_dedup_one_page_7)
	goto L142
L177:
	;
	goto L176
L178:
	;
	v1322 = *(*int32)(unsafe.Add(mBase, uint32(v802)+24))
	base.MemoryCopy(m, v1322+v1009*int32(6), v1006, v1321)
	goto L180
L179:
	;
	goto L180
L180:
	;
	v1327 = *(*int32)(unsafe.Add(mBase, uint32(v802)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v802)+28)) = v1327 + v1004
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v802)+36))
	v1331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v802)+36)) = v1330 + (v1331&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) + int32(4)
	goto L129
L181:
	;
	v1346 = *(*int32)(unsafe.Add(mBase, uint32(v802)+24))
	v1347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+2)))
	v1348 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921))))
	base.MemoryCopy(m, v1346, v1347+(v921+v1348<<(uint(int32(16))%32)), v1345)
	goto L183
L182:
	;
	goto L183
L183:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+28)) = v1343
	v1355 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921)+2)))
	v1356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v921))))
	v1362 = v1355 | v1356<<(uint(int32(16))%32)
	goto L142
L184:
	;
	goto L128
L185:
	;
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(v802)+40))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v802)+24))
	F_pfree(m, v1691)
	mBase = m.M
	v1693 = m.ExcPending
	if v1693 != 0 {
		goto L25
	} else {
		goto L206
	}
L186:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v802)+36)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v802)+28)) = int64(0)
	goto L185
L187:
	;
	v1442 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+16))
	v1450 = v1442
	v1456 = v1421
	goto L188
L188:
	;
	v1460 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+24))
	v1463 = v1460 + v1450*int32(6)
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+20))
	v1467 = v1464 + v1450<<(uint(int32(3))%32)
	v1468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v802)+16)))
	v1469 = v1468 + v1456
	v1474 = v797 + int32(20) + v1469&int32(_a_F__bt_delete_or_dedup_one_page_1)<<(uint(int32(2))%32)
	v1475 = *(*int32)(unsafe.Add(mBase, uint32(v1474)))
	v1478 = v797 + v1475&int32(_a_F__bt_delete_or_dedup_one_page_3)
	v1479 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1478)+7)))
	if v1479&int32(32) != 0 {
		goto L192
	} else {
		goto L193
	}
L189:
	;
	if v1435 <= int32(1) {
		goto L186
	} else {
		goto L205
	}
L190:
	;
	v1655 = v1456 + int32(1)
	v1656 = *(*int32)(unsafe.Add(mBase, uint32(v802)+32))
	if v1655 < v1656 {
		v1450 = v1644
		v1456 = v1655
		goto L188
	} else {
		goto L204
	}
L191:
	;
	v1507 = v1482 & int32(4095)
	v1508 = int32(0)
	if int32(2) <= v1435 {
		goto L196
	} else {
		goto L197
	}
L192:
	;
	v1482 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478)+4)))
	if v1482&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L191
	} else {
		goto L195
	}
L193:
	;
	goto L194
L194:
	;
	v1486 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1467)+4)) = uint16(v1486)
	v1488 = *(*int32)(unsafe.Add(mBase, uint32(v1478)))
	*(*int32)(unsafe.Add(mBase, uint32(v1467))) = v1488
	v1490 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1467)+6)) = uint16(v1490)
	*(*uint8)(unsafe.Add(mBase, uint32(v1463)+3)) = uint8(base.B2i32(int32(1) < v1435))
	v1493 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1463)+2)) = uint8(v1493)
	*(*uint16)(unsafe.Add(mBase, uint32(v1463))) = uint16(v1469)
	v1496 = *(*int32)(unsafe.Add(mBase, uint32(v1474)))
	v1500 = int32(base.Ui32(v1496)>>(uint(int32(17))%32)) + int32(4)
	*(*uint16)(unsafe.Add(mBase, uint32(v1463)+4)) = uint16(v1500)
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+16))
	v1504 = v1502 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+16)) = v1504
	v1644 = v1504
	goto L190
L195:
	;
	goto L194
L196:
	;
	v1512 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478)+2)))
	v1513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478))))
	v1514 = int32(16)
	v1517 = v1512 + (v1478 + v1513<<(uint(v1514)%32))
	v1518 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1517))))
	v1521 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1517)+2)))
	v1522 = v1518<<(uint(v1514)%32) | v1521
	v1525 = int32(6)
	v1527 = v1517 + int32(base.Ui32(v1507)>>(uint(int32(1))%32))*v1525
	v1528 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1527))))
	v1531 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1527)+2)))
	v1532 = v1528<<(uint(v1514)%32) | v1531
	v1536 = v1517 + v1507*v1525
	v1539 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1536-v1525))))
	v1544 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1536-int32(4)))))
	v1549 = base.B2i32(v1532 == v1522)
	v1550 = base.B2i32(v1522 != v1532) & base.B2i32(v1532 == v1539<<(uint(v1514)%32)|v1544)
	goto L198
L197:
	;
	v1549 = v1508
	v1550 = v1508
	goto L198
L198:
	;
	if v1507 == int32(0) {
		v1644 = v1450
		goto L190
	} else {
		goto L199
	}
L199:
	;
	v1555 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478)+2)))
	v1556 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478))))
	v1560 = v1555 + (v1478 + v1556<<(uint(int32(16))%32))
	v1561 = *(*int32)(unsafe.Add(mBase, uint32(v1560)))
	*(*int32)(unsafe.Add(mBase, uint32(v1467))) = v1561
	v1563 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1560)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1467)+4)) = uint16(v1563)
	v1565 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1467)+6)) = uint16(v1565)
	v1567 = int32(6)
	*(*uint16)(unsafe.Add(mBase, uint32(v1463)+4)) = uint16(v1567)
	v1569 = int32(1)
	v1570 = v1507 - v1569
	v1571 = int32(0)
	v1574 = v1549 | v1550&base.B2i32(v1570 == v1571)
	*(*uint8)(unsafe.Add(mBase, uint32(v1463)+3)) = uint8(v1574)
	*(*uint8)(unsafe.Add(mBase, uint32(v1463)+2)) = uint8(v1571)
	*(*uint16)(unsafe.Add(mBase, uint32(v1463))) = uint16(v1469)
	v1580 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+16))
	v1582 = v1580 + v1569
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+16)) = v1582
	if v1507 == v1569 {
		v1644 = v1582
		goto L190
	} else {
		goto L200
	}
L200:
	;
	v1589 = v1463
	v1590 = v1569
	v1592 = v1467
	goto L201
L201:
	;
	v1603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478)+2)))
	v1604 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1478))))
	v1609 = int32(6)
	v1611 = v1603 + (v1478 + v1604<<(uint(int32(16))%32)) + v1590*v1609
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(v1611)))
	*(*int32)(unsafe.Add(mBase, uint32(v1592)+8)) = v1612
	v1614 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1611)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1592)+12)) = uint16(v1614)
	v1616 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+16))
	*(*uint16)(unsafe.Add(mBase, uint32(v1592)+14)) = uint16(v1616)
	*(*uint16)(unsafe.Add(mBase, uint32(v1589)+10)) = uint16(v1609)
	v1621 = v1550 & base.B2i32(v1590 == v1570)
	*(*uint8)(unsafe.Add(mBase, uint32(v1589)+9)) = uint8(v1621)
	v1623 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1589)+8)) = uint8(v1623)
	*(*uint16)(unsafe.Add(mBase, uint32(v1589)+6)) = uint16(v1469)
	v1626 = *(*int32)(unsafe.Add(mBase, uint32(v1420)+16))
	v1627 = int32(1)
	v1628 = v1626 + v1627
	*(*int32)(unsafe.Add(mBase, uint32(v1420)+16)) = v1628
	v1635 = v1590 + v1627
	if v1635 != v1507 {
		v1589 = v1589 + v1609
		v1590 = v1635
		v1592 = v1592 + int32(8)
		goto L201
	} else {
		goto L203
	}
L202:
	;
	v1644 = v1628
	goto L190
L203:
	;
	goto L202
L204:
	;
	goto L189
L205:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v802)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v802+v1660<<(uint(int32(2))%32))+46)) = uint16(v1656)
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v802)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v802)+40)) = v1665 + int32(1)
	goto L186
L206:
	;
	F_pfree(m, v802)
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L25
	} else {
		goto L207
	}
L207:
	;
	F__bt_delitems_delete_check(m, l0, v35, l1, v1420)
	mBase = m.M
	v1697 = m.ExcPending
	if v1697 != 0 {
		goto L25
	} else {
		goto L208
	}
L208:
	;
	v1698 = *(*int32)(unsafe.Add(mBase, uint32(v778)+24))
	F_pfree(m, v1698)
	mBase = m.M
	v1700 = m.ExcPending
	if v1700 != 0 {
		goto L25
	} else {
		goto L209
	}
L209:
	;
	v1701 = *(*int32)(unsafe.Add(mBase, uint32(v778)+28))
	F_pfree(m, v1701)
	mBase = m.M
	v1703 = m.ExcPending
	if v1703 != 0 {
		goto L25
	} else {
		goto L210
	}
L210:
	;
	if v1690 != 0 {
		goto L211
	} else {
		goto L212
	}
L211:
	;
	v1704 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v797)+14)))
	v1705 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v797)+12)))
	v1706 = v1704 - v1705
	v1707 = int32(0)
	if v1707 < v1706 {
		goto L215
	} else {
		goto L216
	}
L212:
	;
	v1716 = v812
	goto L213
L213:
	;
	m.G0 = v778 + int32(32)
	if v1716 != 0 {
		goto L1
	} else {
		goto L221
	}
L214:
	;
	v1711 = int32(341)
	if base.Ui32(v847) <= base.Ui32(v1711) {
		goto L218
	} else {
		goto L219
	}
L215:
	;
	v1710 = v1706
	goto L217
L216:
	;
	v1710 = v1707
	goto L217
L217:
	;
	goto L214
L218:
	;
	v1714 = v1711
	goto L220
L219:
	;
	v1714 = v847
	goto L220
L220:
	;
	v1716 = base.B2i32(base.Ui32(v1714) <= base.Ui32(v1710))
	goto L213
L221:
	;
	goto L102
L222:
	;
	v1748 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1747)+16)))
	if v1748 == int32(0) {
		goto L1
	} else {
		goto L225
	}
L223:
	;
	goto L224
L224:
	;
	v1751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32)+1)))
	if v1751 != int32(1) {
		goto L1
	} else {
		goto L226
	}
L225:
	;
	goto L224
L226:
	;
	v1754 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v1755 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v1756 = int32(0)
	v1757 = m.G0
	v1759 = v1757 - int32(16)
	m.G0 = v1759
	v1761 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v1762 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1761)+10)))
	if v35 < v1756 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v1781 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1780)+16)))
	v1783 = F_palloc(m, int32(1676))
	mBase = m.M
	v1784 = m.ExcPending
	if v1784 != 0 {
		goto L25
	} else {
		goto L231
	}
L228:
	;
	v1766 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[0]))
	v1772 = *(*int32)(unsafe.Add(mBase, uint32(v1766+(v35^int32(-1))<<(uint(int32(2))%32))))
	v1780 = v1772
	goto L227
L229:
	;
	goto L230
L230:
	;
	v1774 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[1]))
	v1780 = v1774 + v35<<(uint(int32(13))%32) + int32(-8192)
	goto L227
L231:
	;
	v1785 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+20)) = v1785
	*(*uint16)(unsafe.Add(mBase, uint32(v1783)+16)) = uint16(v1785)
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+12)) = v1785
	*(*int64)(unsafe.Add(mBase, uint32(v1783)+4)) = int64(5806795784192)
	v1793 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1783))) = uint8(v1793)
	v1796 = F_palloc(m, int32(1352))
	mBase = m.M
	v1797 = m.ExcPending
	if v1797 != 0 {
		goto L25
	} else {
		goto L232
	}
L232:
	;
	v1798 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v1783)+28)) = v1798
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+24)) = v1796
	*(*int64)(unsafe.Add(mBase, uint32(v1783)+36)) = v1798
	v1803 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1780)+12)))
	v1806 = v1781 + v1780
	v1807 = *(*int32)(unsafe.Add(mBase, uint32(v1806)+4))
	if v1807 != 0 {
		goto L233
	} else {
		goto L234
	}
L233:
	;
	v1808 = int32(2)
	goto L235
L234:
	;
	v1808 = int32(1)
	goto L235
L235:
	;
	if v772 != 0 {
		v1848 = v1756
		goto L236
	} else {
		goto L237
	}
L236:
	;
	v1852 = F_PageGetTempPageCopySpecial(m, v1780)
	mBase = m.M
	v1853 = m.ExcPending
	if v1853 != 0 {
		goto L25
	} else {
		goto L247
	}
L237:
	;
	v1809 = *(*int32)(unsafe.Add(mBase, uint32(l0)+192))
	v1810 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1809)+10)))
	v1812 = v1780 + int32(20)
	v1816 = *(*int32)(unsafe.Add(mBase, uint32(v1812+v1808<<(uint(int32(2))%32))))
	v1820 = F__bt_keep_natts_fast(m, l0, v1754, v1780+v1816&int32(_a_F__bt_delete_or_dedup_one_page_3))
	mBase = m.M
	v1821 = m.ExcPending
	if v1821 != 0 {
		goto L25
	} else {
		goto L238
	}
L238:
	;
	if v1810 < v1820 {
		goto L239
	} else {
		goto L240
	}
L239:
	;
	v1824 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1780)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v1824) {
		goto L242
	} else {
		goto L243
	}
L240:
	;
	goto L241
L241:
	;
	v1848 = int32(0)
	goto L236
L242:
	;
	v1834 = int32(base.Ui32(v1824+int32(_a_F__bt_delete_or_dedup_one_page_0))>>(uint(int32(2))%32)) & int32(_a_F__bt_delete_or_dedup_one_page_1)
	goto L244
L243:
	;
	v1834 = int32(0)
	goto L244
L244:
	;
	v1838 = *(*int32)(unsafe.Add(mBase, uint32(v1812+v1834<<(uint(int32(2))%32))))
	v1842 = F__bt_keep_natts_fast(m, l0, v1754, v1780+v1838&int32(_a_F__bt_delete_or_dedup_one_page_3))
	mBase = m.M
	v1843 = m.ExcPending
	if v1843 != 0 {
		goto L25
	} else {
		goto L245
	}
L245:
	;
	if v1810 < v1842 {
		v1848 = int32(1)
		goto L236
	} else {
		goto L246
	}
L246:
	;
	goto L241
L247:
	;
	v1854 = *(*int64)(unsafe.Add(mBase, uint32(v1780)))
	*(*int64)(unsafe.Add(mBase, uint32(v1852))) = v1854
	v1856 = *(*int32)(unsafe.Add(mBase, uint32(v1806)+4))
	if v1856 != 0 {
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
	v2268 = m.ExcPending
	if v2268 != 0 {
		goto L25
	} else {
		goto L334
	}
L250:
	;
	v1857 = *(*int32)(unsafe.Add(mBase, uint32(v1780)+24))
	v1865 = F_PageAddItemExtended(m, v1852, v1780+v1857&int32(_a_F__bt_delete_or_dedup_one_page_3), int32(base.Ui32(v1857)>>(uint(int32(17))%32)), int32(1), int32(0))
	mBase = m.M
	v1866 = m.ExcPending
	if v1866 != 0 {
		goto L25
	} else {
		goto L253
	}
L251:
	;
	goto L252
L252:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v1803) {
		goto L255
	} else {
		goto L256
	}
L253:
	;
	if v1865 == int32(0) {
		goto L249
	} else {
		goto L254
	}
L254:
	;
	goto L252
L255:
	;
	v1877 = int32(base.Ui32(v1803+int32(_a_F__bt_delete_or_dedup_one_page_0)) >> (uint(int32(2)) % 32))
	goto L257
L256:
	;
	v1877 = int32(0)
	goto L257
L257:
	;
	v1879 = v1877 & int32(_a_F__bt_delete_or_dedup_one_page_1)
	if base.Ui32(v1808) <= base.Ui32(v1879) {
		goto L258
	} else {
		goto L259
	}
L258:
	;
	v1882 = v1783 + int32(44)
	v1886 = v1848
	v1892 = v1808
	goto L261
L259:
	;
	goto L260
L260:
	;
	F__bt_dedup_finish_pending(m, v1852, v1783)
	mBase = m.M
	v2183 = m.ExcPending
	if v2183 != 0 {
		goto L25
	} else {
		goto L307
	}
L261:
	;
	v1913 = v1892 & int32(_a_F__bt_delete_or_dedup_one_page_1)
	v1917 = *(*int32)(unsafe.Add(mBase, uint32(v1780+int32(20)+v1913<<(uint(int32(2))%32))))
	v1920 = v1780 + v1917&int32(_a_F__bt_delete_or_dedup_one_page_3)
	if v1808 == v1913 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	goto L260
L263:
	;
	v2151 = v1892 + int32(1)
	if base.Ui32(v2151&int32(_a_F__bt_delete_or_dedup_one_page_1)) <= base.Ui32(v1879) {
		v1886 = v2145
		v1892 = v2151
		goto L261
	} else {
		goto L306
	}
L264:
	;
	v1922 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1920)+7)))
	if v1922&int32(32) != 0 {
		goto L269
	} else {
		goto L270
	}
L265:
	;
	goto L266
L266:
	;
	v1980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1783))))
	if v1980 != int32(1) {
		goto L280
	} else {
		goto L281
	}
L267:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+20)) = v1959
	*(*uint16)(unsafe.Add(mBase, uint32(v1783)+16)) = uint16(v1808)
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+12)) = v1920
	v1965 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+36)) = (v1965&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) | int32(4)
	v1975 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1882+v1975<<(uint(int32(2))%32)))) = uint16(v1808)
	v2145 = v1886
	goto L263
L268:
	;
	v1940 = v1925 & int32(4095)
	v1942 = v1940 * int32(6)
	if v1942 != 0 {
		goto L273
	} else {
		goto L274
	}
L269:
	;
	v1925 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+4)))
	if v1925&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L268
	} else {
		goto L272
	}
L270:
	;
	goto L271
L271:
	;
	v1929 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+24))
	v1930 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v1929)+4)) = uint16(v1930)
	v1932 = *(*int32)(unsafe.Add(mBase, uint32(v1920)))
	*(*int32)(unsafe.Add(mBase, uint32(v1929))) = v1932
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+28)) = int32(1)
	v1936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+6)))
	v1959 = v1936 & int32(_a_F__bt_delete_or_dedup_one_page_7)
	goto L267
L272:
	;
	goto L271
L273:
	;
	v1943 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+24))
	v1944 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+2)))
	v1945 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920))))
	base.MemoryCopy(m, v1943, v1944+(v1920+v1945<<(uint(int32(16))%32)), v1942)
	goto L275
L274:
	;
	goto L275
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+28)) = v1940
	v1952 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+2)))
	v1953 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920))))
	v1959 = v1952 | v1953<<(uint(int32(16))%32)
	goto L267
L276:
	;
	v2087 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1920)+7)))
	if v2087&int32(32) != 0 {
		goto L299
	} else {
		goto L300
	}
L277:
	;
	v2081 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1783))) = uint8(v2081)
	v2084 = v2081
	goto L276
L278:
	;
	v2065 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+8))
	v2066 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1780)+19)))
	v2076 = v2065 - base.I32_trunc_sat_f64_s(base.F64_mul(base.F64_convert_i32_u(v2066<<(uint(int32(8))%32)-v1755-int32(52)), float64(0.04)))
	if base.Ui32(v2076) <= base.Ui32(v2065) {
		goto L294
	} else {
		goto L295
	}
L279:
	;
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+32)) = v2039 + int32(1)
	v2044 = v2006 * int32(6)
	if v2044 != 0 {
		goto L291
	} else {
		goto L292
	}
L280:
	;
	F__bt_dedup_finish_pending(m, v1852, v1783)
	mBase = m.M
	v2031 = m.ExcPending
	if v2031 != 0 {
		goto L25
	} else {
		goto L289
	}
L281:
	;
	v1983 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+12))
	v1984 = F__bt_keep_natts_fast(m, l0, v1983, v1920)
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L25
	} else {
		goto L282
	}
L282:
	;
	if v1984 <= v1762 {
		goto L280
	} else {
		goto L283
	}
L283:
	;
	v1987 = int32(1)
	v1988 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1920)+7)))
	if v1988&int32(32) == int32(0) {
		v2006 = v1987
		v2008 = v1920
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v2009 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+8))
	v2010 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+20))
	v2011 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+28))
	if base.Ui32((v2010+(v2011+v2006)*int32(6)+int32(7))&int32(-8)) <= base.Ui32(v2009) {
		goto L279
	} else {
		goto L287
	}
L285:
	;
	v1993 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+4)))
	if v1993&int32(_a_F__bt_delete_or_dedup_one_page_4) == int32(0) {
		v2006 = v1987
		v2008 = v1920
		goto L284
	} else {
		goto L286
	}
L286:
	;
	v2000 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+2)))
	v2001 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920))))
	v2006 = v1993 & int32(4095)
	v2008 = v2000 + (v1920 + v2001<<(uint(int32(16))%32))
	goto L284
L287:
	;
	if v2011 <= int32(50) {
		goto L280
	} else {
		goto L288
	}
L288:
	;
	v2023 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+4)) = v2023 + int32(1)
	goto L280
L289:
	;
	v2032 = int32(0)
	if v1886 == v2032 {
		v2084 = v2032
		goto L276
	} else {
		goto L290
	}
L290:
	;
	v2035 = int32(1)
	v2036 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+4))
	switch v2036 - int32(5) {
	case 0:
		goto L278
	case 1:
		goto L277
	default:
		v2084 = v2035
		goto L276
	}
L291:
	;
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+24))
	base.MemoryCopy(m, v2045+v2011*int32(6), v2008, v2044)
	goto L293
L292:
	;
	goto L293
L293:
	;
	v2050 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+28)) = v2050 + v2006
	v2053 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+36))
	v2054 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+36)) = v2053 + (v2054&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) + int32(4)
	v2145 = v1886
	goto L263
L294:
	;
	v2079 = v2076
	goto L296
L295:
	;
	v2079 = int32(0)
	goto L296
L296:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+8)) = v2079
	v2084 = v2035
	goto L276
L297:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+32)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+20)) = v2124
	*(*uint16)(unsafe.Add(mBase, uint32(v1783)+16)) = uint16(v1892)
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+12)) = v1920
	v2130 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+6)))
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+36)) = (v2130&int32(_a_F__bt_delete_or_dedup_one_page_7)+int32(7))&int32(_a_F__bt_delete_or_dedup_one_page_8) | int32(4)
	v2140 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1882+v2140<<(uint(int32(2))%32)))) = uint16(v1892)
	v2145 = v2084
	goto L263
L298:
	;
	v2105 = v2090 & int32(4095)
	v2107 = v2105 * int32(6)
	if v2107 != 0 {
		goto L303
	} else {
		goto L304
	}
L299:
	;
	v2090 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+4)))
	if v2090&int32(_a_F__bt_delete_or_dedup_one_page_4) != 0 {
		goto L298
	} else {
		goto L302
	}
L300:
	;
	goto L301
L301:
	;
	v2094 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+24))
	v2095 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v2094)+4)) = uint16(v2095)
	v2097 = *(*int32)(unsafe.Add(mBase, uint32(v1920)))
	*(*int32)(unsafe.Add(mBase, uint32(v2094))) = v2097
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+28)) = int32(1)
	v2101 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+6)))
	v2124 = v2101 & int32(_a_F__bt_delete_or_dedup_one_page_7)
	goto L297
L302:
	;
	goto L301
L303:
	;
	v2108 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+24))
	v2109 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+2)))
	v2110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920))))
	base.MemoryCopy(m, v2108, v2109+(v1920+v2110<<(uint(int32(16))%32)), v2107)
	goto L305
L304:
	;
	goto L305
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1783)+28)) = v2105
	v2117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920)+2)))
	v2118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1920))))
	v2124 = v2117 | v2118<<(uint(int32(16))%32)
	goto L297
L306:
	;
	goto L262
L307:
	;
	v2184 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+40))
	if v2184 == int32(0) {
		goto L309
	} else {
		goto L310
	}
L308:
	;
	v2257 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+24))
	F_pfree(m, v2257)
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L25
	} else {
		goto L332
	}
L309:
	;
	F_pfree(m, v1852)
	mBase = m.M
	v2188 = m.ExcPending
	if v2188 != 0 {
		goto L25
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v2189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1806)+12)))
	if v2189&int32(64) != 0 {
		goto L313
	} else {
		goto L314
	}
L312:
	;
	goto L308
L313:
	;
	v2192 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1852)+16)))
	v2193 = v1852 + v2192
	v2194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2193)+12)))
	v2196 = v2194 & int32(_a_F__bt_delete_or_dedup_one_page_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v2193)+12)) = uint16(v2196)
	goto L315
L314:
	;
	goto L315
L315:
	;
	v2199 = int32(_a_F__bt_delete_or_dedup_one_page_10)
	v2201 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[4])) = v2201 + int32(1)
	F_PageRestoreTempPage(m, v1852, v1780)
	mBase = m.M
	v2206 = m.ExcPending
	if v2206 != 0 {
		goto L25
	} else {
		goto L316
	}
L316:
	;
	F_MarkBufferDirty(m, v35)
	mBase = m.M
	v2208 = m.ExcPending
	if v2208 != 0 {
		goto L25
	} else {
		goto L317
	}
L317:
	;
	v2209 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2209)+118)))
	if v2210 != int32(112) {
		goto L319
	} else {
		goto L320
	}
L318:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v1780))) = base.I64_rotl(v2246, int64(32))
	v2250 = int32(_a_F__bt_delete_or_dedup_one_page_10)
	v2252 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[4]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[4])) = v2252 - int32(1)
	goto L308
L319:
	;
	v2244 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v2245 = m.ExcPending
	if v2245 != 0 {
		goto L25
	} else {
		goto L331
	}
L320:
	;
	v2214 = *(*int32)(unsafe.Add(mBase, _c_F__bt_delete_or_dedup_one_page[5]))
	if v2214 <= int32(0) {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v2217 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2217 != 0 {
		goto L319
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	v2219 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v1759)+14)) = uint16(v2219)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2222 = m.ExcPending
	if v2222 != 0 {
		goto L25
	} else {
		goto L326
	}
L324:
	;
	v2218 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2218 != 0 {
		goto L319
	} else {
		goto L325
	}
L325:
	;
	goto L323
L326:
	;
	F_XLogRegisterBuffer(m, int32(0), v35, int32(8))
	mBase = m.M
	v2226 = m.ExcPending
	if v2226 != 0 {
		goto L25
	} else {
		goto L327
	}
L327:
	;
	F_XLogRegisterData(m, v1759+int32(14), int32(2))
	mBase = m.M
	v2231 = m.ExcPending
	if v2231 != 0 {
		goto L25
	} else {
		goto L328
	}
L328:
	;
	v2235 = *(*int32)(unsafe.Add(mBase, uint32(v1783)+40))
	F_XLogRegisterBufData(m, int32(0), v1783+int32(44), v2235<<(uint(int32(2))%32))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L25
	} else {
		goto L329
	}
L329:
	;
	v2242 = F_XLogInsert(m, int32(11), int32(96))
	mBase = m.M
	v2243 = m.ExcPending
	if v2243 != 0 {
		goto L25
	} else {
		goto L330
	}
L330:
	;
	v2246 = v2242
	goto L318
L331:
	;
	v2246 = v2244
	goto L318
L332:
	;
	F_pfree(m, v1783)
	mBase = m.M
	v2261 = m.ExcPending
	if v2261 != 0 {
		goto L25
	} else {
		goto L333
	}
L333:
	;
	m.G0 = v1759 + int32(16)
	goto L248
L334:
	;
	F_errmsg_internal(m, int32(_a_F__bt_delete_or_dedup_one_page_11), int32(0))
	mBase = m.M
	v2272 = m.ExcPending
	if v2272 != 0 {
		goto L25
	} else {
		goto L335
	}
L335:
	;
	F_errfinish(m, int32(_a_F__bt_delete_or_dedup_one_page_12), int32(131), int32(_a_F__bt_delete_or_dedup_one_page_13))
	mBase = m.M
	v2277 = m.ExcPending
	if v2277 != 0 {
		goto L25
	} else {
		goto L336
	}
L336:
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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	v4 = F_ReadBuffer(m, l0, l1)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if l2 == int32(0) {
			F_UnlockBuffer(m, v4)
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				F__bt_checkpage(m, l0, v4)
				v13 = m.ExcPending
				if v13 != 0 {
					return int32(0)
				} else {
					return v4
				}
			}
		} else {
			F_LockBufferInternal(m, v4, l2)
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				F__bt_checkpage(m, l0, v4)
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					return v4
				}
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
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int64
	_ = v30
	var v32 int64
	_ = v32
	var v34 int64
	_ = v34
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
	if v5 != 0 {
		v48 = v5
		v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
		return v49
	} else {
		v7 = F_ReadBuffer(m, l0, int32(0))
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			F_LockBufferInternal(m, v7, int32(1))
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
							F_UnlockReleaseBuffer(m, v7)
							mBase = m.M
							v22 = m.ExcPending
							if v22 != 0 {
								return int32(0)
							} else {
								return int32(0)
							}
						} else {
							v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+200))
							v27 = F_MemoryContextAlloc(m, v25, int32(48))
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+256)) = v27
								v30 = *(*int64)(unsafe.Add(mBase, uint32(v16)+40))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+40)) = v30
								v32 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+32)) = v32
								v34 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+24)) = v34
								v36 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+16)) = v36
								v38 = *(*int64)(unsafe.Add(mBase, uint32(v16)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v27)+8)) = v38
								v40 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								*(*int64)(unsafe.Add(mBase, uint32(v27))) = v40
								F_UnlockReleaseBuffer(m, v7)
								mBase = m.M
								v43 = m.ExcPending
								if v43 != 0 {
									return int32(0)
								} else {
									v44 = *(*int32)(unsafe.Add(mBase, uint32(l0)+256))
									v48 = v44
									v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+20))
									return v49
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
	var v154 int32
	_ = v154
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
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v402 int32
	_ = v402
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v436 int32
	_ = v436
	var v442 int32
	_ = v442
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v462 int32
	_ = v462
	var v469 int32
	_ = v469
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v584 int32
	_ = v584
	var v586 int32
	_ = v586
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v620 int32
	_ = v620
	var v625 int32
	_ = v625
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v636 int32
	_ = v636
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v645 int32
	_ = v645
	var v652 int32
	_ = v652
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v663 int32
	_ = v663
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v674 int32
	_ = v674
	var v675 int32
	_ = v675
	var v676 int32
	_ = v676
	var v680 int32
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
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
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v740 int32
	_ = v740
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v758 int32
	_ = v758
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
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v785 int32
	_ = v785
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v849 int32
	_ = v849
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v853 int32
	_ = v853
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
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
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v902 int32
	_ = v902
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v935 int32
	_ = v935
	var v937 int32
	_ = v937
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v955 int32
	_ = v955
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v981 int32
	_ = v981
	var v985 int32
	_ = v985
	var v986 int32
	_ = v986
	var v987 int32
	_ = v987
	var v1003 int32
	_ = v1003
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1010 int32
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1028 int32
	_ = v1028
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1037 int32
	_ = v1037
	var v1039 float64
	_ = v1039
	var v1051 int32
	_ = v1051
	var v1065 int32
	_ = v1065
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1103 int32
	_ = v1103
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1111 int32
	_ = v1111
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1123 int32
	_ = v1123
	var v1128 int32
	_ = v1128
	var v1151 int32
	_ = v1151
	var v1197 int32
	_ = v1197
	var v1226 float64
	_ = v1226
	var v1228 int32
	_ = v1228
	var v1231 int32
	_ = v1231
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1238 int32
	_ = v1238
	var v1239 int32
	_ = v1239
	var v1260 int32
	_ = v1260
	var v1292 int32
	_ = v1292
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1300 int32
	_ = v1300
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1310 int32
	_ = v1310
	var v1312 int32
	_ = v1312
	var v1315 int32
	_ = v1315
	var v1316 int32
	_ = v1316
	var v1318 int32
	_ = v1318
	var v1319 int32
	_ = v1319
	var v1335 int32
	_ = v1335
	var v1338 int32
	_ = v1338
	var v1343 int32
	_ = v1343
	var v1368 int32
	_ = v1368
	var v1371 int32
	_ = v1371
	var v1372 int32
	_ = v1372
	var v1377 float64
	_ = v1377
	var v1378 int32
	_ = v1378
	var v1381 int32
	_ = v1381
	var v1384 int32
	_ = v1384
	var v1408 int32
	_ = v1408
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1442 int32
	_ = v1442
	var v1448 int32
	_ = v1448
	var v1467 int32
	_ = v1467
	var v1496 int32
	_ = v1496
	var v1501 int32
	_ = v1501
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1507 int32
	_ = v1507
	var v1508 int32
	_ = v1508
	var v1523 int32
	_ = v1523
	var v1524 int32
	_ = v1524
	var v1526 int32
	_ = v1526
	var v1556 int32
	_ = v1556
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1563 int32
	_ = v1563
	var v1564 int32
	_ = v1564
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1581 int32
	_ = v1581
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1593 int32
	_ = v1593
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1608 int32
	_ = v1608
	var v1609 int32
	_ = v1609
	var v1617 int32
	_ = v1617
	var v1622 int32
	_ = v1622
	var v1623 int32
	_ = v1623
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1633 int32
	_ = v1633
	var v1638 int32
	_ = v1638
	var v1639 int32
	_ = v1639
	var v1640 int32
	_ = v1640
	var v1644 int32
	_ = v1644
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1659 int32
	_ = v1659
	var v1664 int32
	_ = v1664
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1676 int32
	_ = v1676
	var v1681 int32
	_ = v1681
	var v1682 int32
	_ = v1682
	var v1683 int32
	_ = v1683
	var v1685 int32
	_ = v1685
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1693 int32
	_ = v1693
	var v1694 int32
	_ = v1694
	var v1695 int32
	_ = v1695
	var v1697 int32
	_ = v1697
	var v1715 int32
	_ = v1715
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1755 int32
	_ = v1755
	var v1760 int32
	_ = v1760
	var v1763 int32
	_ = v1763
	var v1765 int32
	_ = v1765
	var v1768 int32
	_ = v1768
	var v1769 int32
	_ = v1769
	var v1771 int32
	_ = v1771
	var v1787 int32
	_ = v1787
	var v1790 int32
	_ = v1790
	var v1820 int32
	_ = v1820
	var v1825 int32
	_ = v1825
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1847 int32
	_ = v1847
	var v1871 int32
	_ = v1871
	var v1875 int32
	_ = v1875
	var v1877 int32
	_ = v1877
	var v1878 int32
	_ = v1878
	var v1880 int32
	_ = v1880
	var v1881 int32
	_ = v1881
	var v1882 int32
	_ = v1882
	var v1883 int32
	_ = v1883
	var v1886 int32
	_ = v1886
	var v1898 int32
	_ = v1898
	var v1902 int32
	_ = v1902
	var v1903 int32
	_ = v1903
	var v1935 int32
	_ = v1935
	var v1936 int32
	_ = v1936
	var v1937 int32
	_ = v1937
	var v1948 int32
	_ = v1948
	var v1967 int32
	_ = v1967
	var v1978 int32
	_ = v1978
	var v1981 int32
	_ = v1981
	var v1983 int32
	_ = v1983
	var v1987 int32
	_ = v1987
	var v1991 int32
	_ = v1991
	var v1992 int32
	_ = v1992
	var v1993 int32
	_ = v1993
	var v1994 int32
	_ = v1994
	var v1997 int32
	_ = v1997
	var v1998 int32
	_ = v1998
	var v1999 int32
	_ = v1999
	var v2001 int32
	_ = v2001
	var v2003 int32
	_ = v2003
	var v2020 int32
	_ = v2020
	var v2052 int32
	_ = v2052
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2063 int32
	_ = v2063
	var v2065 int32
	_ = v2065
	var v2068 int32
	_ = v2068
	var v2069 int32
	_ = v2069
	var v2071 int32
	_ = v2071
	var v2073 int32
	_ = v2073
	var v2090 int32
	_ = v2090
	var v2123 int32
	_ = v2123
	var v2125 int32
	_ = v2125
	var v2129 int32
	_ = v2129
	var v2130 int32
	_ = v2130
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2137 int32
	_ = v2137
	var v2139 int32
	_ = v2139
	var v2141 int64
	_ = v2141
	var v2148 int32
	_ = v2148
	var v2150 int32
	_ = v2150
	var v2151 int32
	_ = v2151
	var v2159 int32
	_ = v2159
	var v2166 int32
	_ = v2166
	var v2170 int32
	_ = v2170
	var v2171 int32
	_ = v2171
	var v2177 int32
	_ = v2177
	var v2178 int32
	_ = v2178
	var v2182 int32
	_ = v2182
	var v2189 int32
	_ = v2189
	var v2191 int32
	_ = v2191
	var v2192 int32
	_ = v2192
	var v2193 int32
	_ = v2193
	var v2194 int32
	_ = v2194
	var v2198 int32
	_ = v2198
	var v2199 int32
	_ = v2199
	var v2204 int32
	_ = v2204
	var v2205 int32
	_ = v2205
	var v2208 int32
	_ = v2208
	var v2209 int32
	_ = v2209
	var v2215 int32
	_ = v2215
	var v2220 int32
	_ = v2220
	var v2221 int32
	_ = v2221
	var v2224 int32
	_ = v2224
	var v2228 int32
	_ = v2228
	var v2233 int32
	_ = v2233
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2244 int32
	_ = v2244
	var v2250 int32
	_ = v2250
	var v2252 int32
	_ = v2252
	var v2253 int32
	_ = v2253
	var v2258 int32
	_ = v2258
	var v2259 int32
	_ = v2259
	var v2260 int32
	_ = v2260
	var v2264 int32
	_ = v2264
	var v2268 int32
	_ = v2268
	var v2269 int32
	_ = v2269
	var v2270 int32
	_ = v2270
	var v2272 int32
	_ = v2272
	var v2273 int32
	_ = v2273
	var v2278 int32
	_ = v2278
	var v2281 int32
	_ = v2281
	var v2326 int32
	_ = v2326
	var v2327 int32
	_ = v2327
	var v2329 int32
	_ = v2329
	var v2330 int32
	_ = v2330
	var v2332 int32
	_ = v2332
	var v2334 int32
	_ = v2334
	var v2382 int32
	_ = v2382
	var v2384 int32
	_ = v2384
	var v2388 int32
	_ = v2388
	var v2390 int32
	_ = v2390
	var v2392 int32
	_ = v2392
	var v2394 int32
	_ = v2394
	var v2395 int32
	_ = v2395
	var v2398 int32
	_ = v2398
	var v2400 int32
	_ = v2400
	var v2402 int32
	_ = v2402
	var v2407 int32
	_ = v2407
	var v2415 int32
	_ = v2415
	var v2416 int32
	_ = v2416
	var v2420 int32
	_ = v2420
	var v2422 int32
	_ = v2422
	var v2425 int32
	_ = v2425
	var v2426 int32
	_ = v2426
	var v2434 int32
	_ = v2434
	var v2436 int32
	_ = v2436
	var v2443 int32
	_ = v2443
	var v2451 int32
	_ = v2451
	var v2455 int32
	_ = v2455
	var v2486 int32
	_ = v2486
	var v2487 int32
	_ = v2487
	var v2491 int32
	_ = v2491
	var v2497 int32
	_ = v2497
	var v2501 int32
	_ = v2501
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2518 int32
	_ = v2518
	var v2520 int64
	_ = v2520
	var v2527 int32
	_ = v2527
	var v2528 int32
	_ = v2528
	var v2530 int32
	_ = v2530
	var v2531 int32
	_ = v2531
	var v2536 int32
	_ = v2536
	var v2537 int32
	_ = v2537
	var v2538 int32
	_ = v2538
	var v2542 int32
	_ = v2542
	var v2549 int32
	_ = v2549
	var v2550 int32
	_ = v2550
	var v2558 int32
	_ = v2558
	var v2560 int64
	_ = v2560
	var v2567 int32
	_ = v2567
	var v2568 int32
	_ = v2568
	var v2570 int32
	_ = v2570
	var v2571 int32
	_ = v2571
	var v2576 int32
	_ = v2576
	var v2578 int32
	_ = v2578
	var v2581 int32
	_ = v2581
	var v2595 int32
	_ = v2595
	var v2599 int32
	_ = v2599
	var v2630 int32
	_ = v2630
	var v2638 int32
	_ = v2638
	var v2640 int64
	_ = v2640
	var v2647 int32
	_ = v2647
	var v2648 int32
	_ = v2648
	var v2650 int32
	_ = v2650
	var v2651 int32
	_ = v2651
	var v2658 int32
	_ = v2658
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2668 int32
	_ = v2668
	var v2674 int32
	_ = v2674
	var v2676 int32
	_ = v2676
	var v2682 int32
	_ = v2682
	var v2683 int32
	_ = v2683
	var v2684 int32
	_ = v2684
	var v2685 int32
	_ = v2685
	var v2687 int32
	_ = v2687
	var v2688 int32
	_ = v2688
	var v2690 int32
	_ = v2690
	var v2692 int32
	_ = v2692
	var v2694 int32
	_ = v2694
	var v2695 int32
	_ = v2695
	var v2696 int32
	_ = v2696
	var v2697 int32
	_ = v2697
	var v2699 int32
	_ = v2699
	var v2710 int32
	_ = v2710
	var v2716 int32
	_ = v2716
	var v2718 int32
	_ = v2718
	var v2724 int32
	_ = v2724
	var v2729 int32
	_ = v2729
	var v2731 int32
	_ = v2731
	var v2733 int32
	_ = v2733
	var v2736 int32
	_ = v2736
	var v2742 int32
	_ = v2742
	var v2748 int32
	_ = v2748
	var v2750 int32
	_ = v2750
	var v2756 int32
	_ = v2756
	var v2757 int32
	_ = v2757
	var v2758 int32
	_ = v2758
	var v2759 int32
	_ = v2759
	var v2761 int32
	_ = v2761
	var v2764 int32
	_ = v2764
	var v2766 int32
	_ = v2766
	var v2767 int32
	_ = v2767
	var v2771 int32
	_ = v2771
	var v2774 int32
	_ = v2774
	var v2775 int32
	_ = v2775
	var v2777 int32
	_ = v2777
	var v2785 int32
	_ = v2785
	var v2789 int32
	_ = v2789
	var v2792 int32
	_ = v2792
	var v2797 int32
	_ = v2797
	var v2801 int32
	_ = v2801
	var v2805 int32
	_ = v2805
	var v2809 int32
	_ = v2809
	var v2815 int32
	_ = v2815
	var v2816 int32
	_ = v2816
	var v2817 int32
	_ = v2817
	var v2820 int32
	_ = v2820
	var v2821 int32
	_ = v2821
	var v2823 int32
	_ = v2823
	var v2826 int32
	_ = v2826
	var v2830 int32
	_ = v2830
	var v2832 int32
	_ = v2832
	var v2840 int32
	_ = v2840
	var v2842 int32
	_ = v2842
	var v2844 int32
	_ = v2844
	var v2847 int32
	_ = v2847
	var v2851 int32
	_ = v2851
	var v2852 int32
	_ = v2852
	var v2853 int64
	_ = v2853
	var v2854 int32
	_ = v2854
	var v2855 int64
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2859 int32
	_ = v2859
	var v2860 int64
	_ = v2860
	var v2862 int64
	_ = v2862
	var v2871 int32
	_ = v2871
	var v2877 int32
	_ = v2877
	var v2879 int32
	_ = v2879
	var v2885 int32
	_ = v2885
	var v2887 int32
	_ = v2887
	var v2889 int32
	_ = v2889
	var v2894 int32
	_ = v2894
	var v2896 int32
	_ = v2896
	var v2901 int32
	_ = v2901
	var v2903 int32
	_ = v2903
	var v2907 int32
	_ = v2907
	var v2913 int32
	_ = v2913
	var v2915 int32
	_ = v2915
	var v2916 int32
	_ = v2916
	var v2921 int32
	_ = v2921
	var v2922 int32
	_ = v2922
	var v2926 int32
	_ = v2926
	var v2932 int32
	_ = v2932
	var v2934 int32
	_ = v2934
	var v2935 int32
	_ = v2935
	var v2940 int32
	_ = v2940
	var v2941 int32
	_ = v2941
	var v2943 int32
	_ = v2943
	var v2947 int32
	_ = v2947
	var v2948 int32
	_ = v2948
	var v2953 int32
	_ = v2953
	var v2954 int32
	_ = v2954
	var v2958 int32
	_ = v2958
	var v2964 int32
	_ = v2964
	var v2966 int32
	_ = v2966
	var v2972 int32
	_ = v2972
	var v2974 int32
	_ = v2974
	var v2975 int32
	_ = v2975
	var v2976 int32
	_ = v2976
	var v2979 int32
	_ = v2979
	var v2981 int32
	_ = v2981
	var v2982 int32
	_ = v2982
	var v2983 int32
	_ = v2983
	var v2984 int32
	_ = v2984
	var v2986 int32
	_ = v2986
	var v2992 int32
	_ = v2992
	var v2998 int32
	_ = v2998
	var v3006 int32
	_ = v3006
	var v3007 int32
	_ = v3007
	var v3011 int32
	_ = v3011
	var v3013 int32
	_ = v3013
	var v3014 int32
	_ = v3014
	var v3017 int32
	_ = v3017
	var v3025 int32
	_ = v3025
	var v3030 int32
	_ = v3030
	var v3036 int32
	_ = v3036
	var v3038 int32
	_ = v3038
	var v3039 int32
	_ = v3039
	var v3044 int32
	_ = v3044
	var v3045 int32
	_ = v3045
	var v3047 int32
	_ = v3047
	var v3050 int32
	_ = v3050
	var v3056 int32
	_ = v3056
	var v3062 int32
	_ = v3062
	var v3064 int32
	_ = v3064
	var v3070 int32
	_ = v3070
	var v3071 int32
	_ = v3071
	var v3072 int32
	_ = v3072
	var v3073 int32
	_ = v3073
	var v3075 int32
	_ = v3075
	var v3078 int32
	_ = v3078
	var v3080 int32
	_ = v3080
	var v3081 int32
	_ = v3081
	var v3085 int32
	_ = v3085
	var v3088 int32
	_ = v3088
	var v3089 int32
	_ = v3089
	var v3092 int32
	_ = v3092
	var v3097 int32
	_ = v3097
	var v3104 int32
	_ = v3104
	var v3105 int32
	_ = v3105
	var v3107 int32
	_ = v3107
	var v3109 int32
	_ = v3109
	var v3111 int32
	_ = v3111
	var v3113 int32
	_ = v3113
	var v3115 int32
	_ = v3115
	var v3117 int32
	_ = v3117
	var v3122 int32
	_ = v3122
	var v3128 int32
	_ = v3128
	var v3131 int32
	_ = v3131
	var v3135 int32
	_ = v3135
	var v3139 int32
	_ = v3139
	var v3143 int32
	_ = v3143
	var v3146 int32
	_ = v3146
	var v3147 int32
	_ = v3147
	var v3154 int32
	_ = v3154
	var v3155 int32
	_ = v3155
	var v3156 int32
	_ = v3156
	var v3158 int32
	_ = v3158
	var v3162 int32
	_ = v3162
	var v3164 int64
	_ = v3164
	var v3165 int32
	_ = v3165
	var v3166 int64
	_ = v3166
	var v3167 int32
	_ = v3167
	var v3170 int64
	_ = v3170
	var v3179 int32
	_ = v3179
	var v3185 int32
	_ = v3185
	var v3187 int32
	_ = v3187
	var v3193 int32
	_ = v3193
	var v3195 int64
	_ = v3195
	var v3197 int64
	_ = v3197
	var v3199 int32
	_ = v3199
	var v3201 int32
	_ = v3201
	var v3206 int32
	_ = v3206
	var v3210 int32
	_ = v3210
	var v3212 int32
	_ = v3212
	var v3215 int32
	_ = v3215
	var v3219 int32
	_ = v3219
	var v3225 int32
	_ = v3225
	var v3227 int32
	_ = v3227
	var v3228 int32
	_ = v3228
	var v3233 int32
	_ = v3233
	var v3234 int32
	_ = v3234
	var v3236 int32
	_ = v3236
	var v3239 int32
	_ = v3239
	var v3240 int32
	_ = v3240
	var v3243 int32
	_ = v3243
	var v3244 int32
	_ = v3244
	var v3245 int32
	_ = v3245
	var v3247 int64
	_ = v3247
	var v3251 int32
	_ = v3251
	var v3252 int32
	_ = v3252
	var v3255 int32
	_ = v3255
	var v3256 int32
	_ = v3256
	var v3257 int32
	_ = v3257
	var v3259 int32
	_ = v3259
	var v3261 int32
	_ = v3261
	var v3263 int32
	_ = v3263
	var v3267 int32
	_ = v3267
	var v3269 int32
	_ = v3269
	var v3317 int32
	_ = v3317
	var v3319 int32
	_ = v3319
	var v3326 int32
	_ = v3326
	var v3329 int32
	_ = v3329
	var v3330 int32
	_ = v3330
	var v3331 int32
	_ = v3331
	var v3332 int32
	_ = v3332
	var v3336 int32
	_ = v3336
	var v3342 int32
	_ = v3342
	var v3344 int32
	_ = v3344
	var v3345 int32
	_ = v3345
	var v3350 int32
	_ = v3350
	var v3351 int32
	_ = v3351
	var v3352 int32
	_ = v3352
	var v3367 int32
	_ = v3367
	var v3372 int32
	_ = v3372
	var v3376 int32
	_ = v3376
	var v3377 int32
	_ = v3377
	var v3384 int32
	_ = v3384
	var v3389 int32
	_ = v3389
	var v3393 int32
	_ = v3393
	var v3394 int32
	_ = v3394
	var v3403 int32
	_ = v3403
	var v3408 int32
	_ = v3408
	var v3412 int32
	_ = v3412
	var v3413 int32
	_ = v3413
	var v3422 int32
	_ = v3422
	var v3427 int32
	_ = v3427
	var v3431 int32
	_ = v3431
	var v3432 int32
	_ = v3432
	var v3441 int32
	_ = v3441
	var v3446 int32
	_ = v3446
	var v3450 int32
	_ = v3450
	var v3451 int32
	_ = v3451
	var v3460 int32
	_ = v3460
	var v3465 int32
	_ = v3465
	var v3469 int32
	_ = v3469
	var v3470 int32
	_ = v3470
	var v3479 int32
	_ = v3479
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
	var v3510 int32
	_ = v3510
	var v3511 int32
	_ = v3511
	var v3512 int32
	_ = v3512
	var v3513 int32
	_ = v3513
	var v3524 int32
	_ = v3524
	var v3529 int32
	_ = v3529
	var v3533 int32
	_ = v3533
	var v3537 int32
	_ = v3537
	var v3543 int32
	_ = v3543
	var v3545 int32
	_ = v3545
	var v3546 int32
	_ = v3546
	var v3551 int32
	_ = v3551
	var v3552 int32
	_ = v3552
	var v3553 int32
	_ = v3553
	var v3562 int32
	_ = v3562
	var v3567 int32
	_ = v3567
	v10 = l9
	v12 = int32(0)
	v46 = m.G0
	v48 = v46 - int32(_a_F__bt_insertonpg_0)
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
	F_errstart_cold(m, int32(24), int32(0))
	mBase = m.M
	v3533 = m.ExcPending
	if v3533 != 0 {
		goto L20
	} else {
		goto L691
	}
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3507 = m.ExcPending
	if v3507 != 0 {
		goto L20
	} else {
		goto L687
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3488 = m.ExcPending
	if v3488 != 0 {
		goto L20
	} else {
		goto L684
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3469 = m.ExcPending
	if v3469 != 0 {
		goto L20
	} else {
		goto L681
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3450 = m.ExcPending
	if v3450 != 0 {
		goto L20
	} else {
		goto L678
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3431 = m.ExcPending
	if v3431 != 0 {
		goto L20
	} else {
		goto L675
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3412 = m.ExcPending
	if v3412 != 0 {
		goto L20
	} else {
		goto L672
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3393 = m.ExcPending
	if v3393 != 0 {
		goto L20
	} else {
		goto L669
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3376 = m.ExcPending
	if v3376 != 0 {
		goto L20
	} else {
		goto L666
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v3326 = m.ExcPending
	if v3326 != 0 {
		goto L20
	} else {
		goto L658
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
	v103 = v12
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
	if v90&int32(_a_F__bt_insertonpg_3) == int32(0) {
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
	v103 = v79
	v105 = v99
	v106 = l6
	goto L17
L23:
	;
	if v10 != 0 {
		goto L653
	} else {
		goto L654
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
	v2948 = int32(0)
	if l10 == v2948 {
		v2981 = v2948
		v2982 = v12
		v2983 = v12
		goto L543
	} else {
		goto L544
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
	v162 = v102 & int32(_a_F__bt_insertonpg_4)
	v163 = m.G0
	v165 = v163 + int32(-64)
	m.G0 = v165
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+12)))
	v170 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136)+19)))
	v172 = v170 << (uint(int32(8)) % 32)
	v174 = v172 - int32(40)
	v176 = v48 + int32(207)
	if base.Ui32(v167) < base.Ui32(int32(25)) {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v145+(l3^int32(-1))*int32(56))+16))
	v160 = v151
	goto L35
L37:
	;
	goto L38
L38:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v154 = int32(56)
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v153+l3*v154-v154)+16))
	v160 = v159
	goto L35
L39:
	;
	v182 = int32(0)
	goto L41
L40:
	;
	v182 = int32(base.Ui32(v167+int32(_a_F__bt_insertonpg_5)) >> (uint(int32(2)) % 32))
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
	v196 = v172 - (int32(base.Ui32(v186)>>(uint(int32(17))%32))+int32(7))&int32(_a_F__bt_insertonpg_6) - int32(44)
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
	v227 = v182 & int32(_a_F__bt_insertonpg_4)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+48)) = v227
	*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)) = uint16(v162)
	*(*uint8)(unsafe.Add(mBase, uint32(v165)+29)) = uint8(base.B2i32(v220 == int32(0)))
	v234 = F_palloc_mul(m, int32(10), v227)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L20
	} else {
		goto L52
	}
L52:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+56)) = v234
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	if v241 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v242 = int32(2)
	goto L55
L54:
	;
	v242 = int32(1)
	goto L55
L55:
	;
	if base.Ui32(v242) <= base.Ui32(v227) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v246 = int32(0)
	v262 = v242
	v263 = v246
	v266 = v246
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
	v294 = v262 & int32(_a_F__bt_insertonpg_4)
	v296 = v294 << (uint(int32(2)) % 32)
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v136+int32(20)+v296)))
	v304 = (int32(base.Ui32(v298)>>(uint(int32(17))%32)) + int32(7)) & int32(_a_F__bt_insertonpg_6)
	v306 = v304 | int32(4)
	if base.Ui32(v294) < base.Ui32(v162) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L58
L61:
	;
	v796 = v266 + v306
	v799 = v262 + int32(1)
	if base.Ui32(v799&int32(_a_F__bt_insertonpg_4)) <= base.Ui32(v227) {
		v262 = v799
		v263 = int32(0) - v796
		v266 = v796
		goto L59
	} else {
		goto L142
	}
L62:
	;
	v309 = v163 + int32(-52)
	v310 = int32(0)
	v314 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v309)+18)))
	if v314 == v294 {
		goto L69
	} else {
		goto L70
	}
L63:
	;
	goto L64
L64:
	;
	if base.Ui32(v162) < base.Ui32(v294) {
		goto L91
	} else {
		goto L92
	}
L65:
	;
	goto L61
L66:
	;
	v387 = v382 + v383 - (v266 + v380)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	if v384 != 0 {
		goto L80
	} else {
		goto L81
	}
L67:
	;
	v380 = v375
	v381 = v376
	v382 = v377
	v383 = v378
	v384 = int32(1)
	goto L66
L68:
	;
	v375 = v370
	v376 = v371
	v377 = int32(-8)
	v378 = v373
	goto L67
L69:
	;
	v316 = *(*int32)(unsafe.Add(mBase, uint32(v309)+24))
	v317 = *(*int32)(unsafe.Add(mBase, uint32(v309)+28))
	v319 = v316 + (v266 - v317)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	v321 = *(*int32)(unsafe.Add(mBase, uint32(v309)+12))
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+16)))
	if v322&int32(1) != 0 {
		v370 = v321
		v371 = v319
		v373 = v320
		goto L68
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v309)+16)))
	if base.B2i32(v325 != int32(1))|base.B2i32(base.Ui32(v306) < base.Ui32(int32(65))) == int32(0) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v380 = v321
	v381 = v319
	v382 = v310
	v383 = v320
	v384 = v310
	goto L66
L73:
	;
	v333 = int32(-8)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v334+v294<<(uint(int32(2))%32))+20))
	v341 = v334 + v338&int32(_a_F__bt_insertonpg_1)
	v342 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341)+6)))
	if v342&int32(_a_F__bt_insertonpg_3) == int32(0) {
		v358 = v333
		goto L76
	} else {
		goto L77
	}
L74:
	;
	goto L75
L75:
	;
	v364 = *(*int32)(unsafe.Add(mBase, uint32(v309)+24))
	v365 = *(*int32)(unsafe.Add(mBase, uint32(v309)+28))
	v367 = v364 + (v266 - v365)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	if v325 != 0 {
		v370 = v306
		v371 = v367
		v373 = v368
		goto L68
	} else {
		goto L79
	}
L76:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v309)+24))
	v360 = *(*int32)(unsafe.Add(mBase, uint32(v309)+28))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v309)+20))
	v375 = v306
	v376 = v359 + (v266 - v360)
	v377 = v358
	v378 = v363
	goto L67
L77:
	;
	v347 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v341)+5)))
	if v347&int32(32) == int32(0) {
		v358 = v333
		goto L76
	} else {
		goto L78
	}
L78:
	;
	v354 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v341)+2)))
	v358 = v342&int32(_a_F__bt_insertonpg_7) - v354 - int32(8)
	goto L76
L79:
	;
	v380 = v306
	v381 = v367
	v382 = int32(0)
	v383 = v368
	v384 = v310
	goto L66
L80:
	;
	v393 = int32(0)
	goto L82
L81:
	;
	v393 = v380 + int32(_a_F__bt_insertonpg_8)
	goto L82
L82:
	;
	v394 = v381 - v388 + v393
	if (v387|v394)&int32(_a_F__bt_insertonpg_9) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v309)+32))
	if base.Ui32(v400) < base.Ui32(v380) {
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
	v402 = v400
	goto L88
L87:
	;
	v402 = v380
	goto L88
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+32)) = v402
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v309)+44))
	v405 = *(*int32)(unsafe.Add(mBase, uint32(v309)+40))
	v406 = int32(10)
	v409 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v404+v405*v406))) = uint16(v409)
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v309)+44))
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v309)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v411+v412*v406)+2)) = uint16(v387)
	v417 = *(*int32)(unsafe.Add(mBase, uint32(v309)+44))
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v309)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v417+v418*v406)+4)) = uint16(v394)
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v309)+44))
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v309)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v423+v424*v406)+6)) = uint16(v294)
	v429 = *(*int32)(unsafe.Add(mBase, uint32(v309)+44))
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v309)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v429+v430*v406)+8)) = uint8(v409)
	v436 = *(*int32)(unsafe.Add(mBase, uint32(v309)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v309)+40)) = v436 + int32(1)
	goto L85
L89:
	;
	v745 = v744 + v740
	if (v745|v741)&int32(_a_F__bt_insertonpg_9) != 0 {
		goto L61
	} else {
		goto L138
	}
L90:
	;
	v740 = v731
	v741 = v263 - v306 + v734 + v733 - v732
	v744 = int32(0)
	goto L89
L91:
	;
	v442 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if base.B2i32(v442 != int32(1))|base.B2i32(base.Ui32(v298) < base.Ui32(int32(_a_F__bt_insertonpg_10))) == int32(0) {
		goto L94
	} else {
		goto L95
	}
L92:
	;
	goto L93
L93:
	;
	v493 = v163 + int32(-52)
	v494 = int32(0)
	v498 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v493)+18)))
	if v498 == v294 {
		goto L105
	} else {
		goto L106
	}
L94:
	;
	v450 = int32(-8)
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v451+v296)+20))
	v456 = v451 + v453&int32(_a_F__bt_insertonpg_1)
	v457 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v456)+6)))
	if v457&int32(_a_F__bt_insertonpg_3) == int32(0) {
		v473 = v450
		goto L97
	} else {
		goto L98
	}
L95:
	;
	goto L96
L96:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v483 = v480 + (v266 - v481)
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v486 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	if v442 != 0 {
		v731 = v483
		v732 = v485
		v733 = int32(-8)
		v734 = v486
		goto L90
	} else {
		goto L100
	}
L97:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	v731 = v474 + v266 - v476
	v732 = v478
	v733 = v473
	v734 = v479
	goto L90
L98:
	;
	v462 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v456)+5)))
	if v462&int32(32) == int32(0) {
		v473 = v450
		goto L97
	} else {
		goto L99
	}
L99:
	;
	v469 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v456)+2)))
	v473 = v457&int32(_a_F__bt_insertonpg_7) - v469 - int32(8)
	goto L97
L100:
	;
	v740 = v483
	v741 = v263 + v486 - (v485 + v306)
	v744 = v304 + int32(_a_F__bt_insertonpg_6)
	goto L89
L101:
	;
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if base.B2i32(v625 != int32(1))|base.B2i32(base.Ui32(v298) < base.Ui32(int32(_a_F__bt_insertonpg_10))) == int32(0) {
		goto L127
	} else {
		goto L128
	}
L102:
	;
	v571 = v566 + v567 - (v266 + v564)
	v572 = *(*int32)(unsafe.Add(mBase, uint32(v493)+12))
	if v568 != 0 {
		goto L116
	} else {
		goto L117
	}
L103:
	;
	v564 = v559
	v565 = v560
	v566 = v561
	v567 = v562
	v568 = int32(1)
	goto L102
L104:
	;
	v559 = v554
	v560 = v555
	v561 = int32(-8)
	v562 = v557
	goto L103
L105:
	;
	v500 = *(*int32)(unsafe.Add(mBase, uint32(v493)+24))
	v501 = *(*int32)(unsafe.Add(mBase, uint32(v493)+28))
	v503 = v500 + (v266 - v501)
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v493)+20))
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v493)+12))
	v506 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+16)))
	if v506&int32(1) != 0 {
		v554 = v505
		v555 = v503
		v557 = v504
		goto L104
	} else {
		goto L108
	}
L106:
	;
	goto L107
L107:
	;
	v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v493)+16)))
	if base.B2i32(v509 != int32(1))|base.B2i32(base.Ui32(v306) < base.Ui32(int32(65))) == int32(0) {
		goto L109
	} else {
		goto L110
	}
L108:
	;
	v564 = v505
	v565 = v503
	v566 = v494
	v567 = v504
	v568 = v494
	goto L102
L109:
	;
	v517 = int32(-8)
	v518 = *(*int32)(unsafe.Add(mBase, uint32(v493)+4))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v518+v294<<(uint(int32(2))%32))+20))
	v525 = v518 + v522&int32(_a_F__bt_insertonpg_1)
	v526 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v525)+6)))
	if v526&int32(_a_F__bt_insertonpg_3) == int32(0) {
		v542 = v517
		goto L112
	} else {
		goto L113
	}
L110:
	;
	goto L111
L111:
	;
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v493)+24))
	v549 = *(*int32)(unsafe.Add(mBase, uint32(v493)+28))
	v551 = v548 + (v266 - v549)
	v552 = *(*int32)(unsafe.Add(mBase, uint32(v493)+20))
	if v509 != 0 {
		v554 = v306
		v555 = v551
		v557 = v552
		goto L104
	} else {
		goto L115
	}
L112:
	;
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v493)+24))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v493)+28))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v493)+20))
	v559 = v306
	v560 = v543 + (v266 - v544)
	v561 = v542
	v562 = v547
	goto L103
L113:
	;
	v531 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v525)+5)))
	if v531&int32(32) == int32(0) {
		v542 = v517
		goto L112
	} else {
		goto L114
	}
L114:
	;
	v538 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v525)+2)))
	v542 = v526&int32(_a_F__bt_insertonpg_7) - v538 - int32(8)
	goto L112
L115:
	;
	v564 = v306
	v565 = v551
	v566 = int32(0)
	v567 = v552
	v568 = v494
	goto L102
L116:
	;
	v577 = int32(0)
	goto L118
L117:
	;
	v577 = v564 + int32(_a_F__bt_insertonpg_8)
	goto L118
L118:
	;
	v578 = v565 - v572 + v577
	if (v571|v578)&int32(_a_F__bt_insertonpg_9) == int32(0) {
		goto L119
	} else {
		goto L120
	}
L119:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v493)+32))
	if base.Ui32(v584) < base.Ui32(v564) {
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
	v586 = v584
	goto L124
L123:
	;
	v586 = v564
	goto L124
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v493)+32)) = v586
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v493)+44))
	v589 = *(*int32)(unsafe.Add(mBase, uint32(v493)+40))
	v590 = int32(10)
	v593 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v588+v589*v590))) = uint16(v593)
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v493)+44))
	v596 = *(*int32)(unsafe.Add(mBase, uint32(v493)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v595+v596*v590)+2)) = uint16(v571)
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v493)+44))
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v493)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v601+v602*v590)+4)) = uint16(v578)
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v493)+44))
	v608 = *(*int32)(unsafe.Add(mBase, uint32(v493)+40))
	*(*uint16)(unsafe.Add(mBase, uint32(v607+v608*v590)+6)) = uint16(v294)
	v613 = *(*int32)(unsafe.Add(mBase, uint32(v493)+44))
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v493)+40))
	*(*uint8)(unsafe.Add(mBase, uint32(v613+v614*v590)+8)) = uint8(v593)
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v493)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v493)+40)) = v620 + int32(1)
	goto L121
L125:
	;
	v687 = v686 + v683
	if (v687|v684)&int32(_a_F__bt_insertonpg_9) != 0 {
		goto L61
	} else {
		goto L134
	}
L126:
	;
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v683 = v674
	v684 = v263 - v306 + v675 + v676 - v680
	v686 = int32(0)
	goto L125
L127:
	;
	v633 = int32(-8)
	v634 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v634+v296)+20))
	v639 = v634 + v636&int32(_a_F__bt_insertonpg_1)
	v640 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v639)+6)))
	if v640&int32(_a_F__bt_insertonpg_3) == int32(0) {
		v656 = v633
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v663 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v665 = v662 + (v266 - v663)
	v667 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	if v625 != 0 {
		v674 = v665
		v675 = v667
		v676 = int32(-8)
		goto L126
	} else {
		goto L133
	}
L130:
	;
	v657 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v659 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	v674 = v657 + v266 - v659
	v675 = v661
	v676 = v656
	goto L126
L131:
	;
	v645 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v639)+5)))
	if v645&int32(32) == int32(0) {
		v656 = v633
		goto L130
	} else {
		goto L132
	}
L132:
	;
	v652 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v639)+2)))
	v656 = v640&int32(_a_F__bt_insertonpg_7) - v652 - int32(8)
	goto L130
L133:
	;
	v669 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v683 = v665
	v684 = v667 + v263 - (v306 + v669)
	v686 = v304 + int32(_a_F__bt_insertonpg_6)
	goto L125
L134:
	;
	v691 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	if base.Ui32(v691) < base.Ui32(v306) {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v693 = v691
	goto L137
L136:
	;
	v693 = v306
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+44)) = v693
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v697 = int32(10)
	v700 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v695+v696*v697))) = uint16(v700)
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v703 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v702+v703*v697)+2)) = uint16(v684)
	v708 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v709 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v708+v709*v697)+4)) = uint16(v687)
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v714+v715*v697)+6)) = uint16(v262)
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v721 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v725 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v720+v721*v697)+8)) = uint8(v725)
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = v727 + v725
	goto L61
L138:
	;
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	if base.Ui32(v749) < base.Ui32(v306) {
		goto L139
	} else {
		goto L140
	}
L139:
	;
	v751 = v749
	goto L141
L140:
	;
	v751 = v306
	goto L141
L141:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+44)) = v751
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v755 = int32(10)
	v758 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v753+v754*v755))) = uint16(v758)
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v761 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v760+v761*v755)+2)) = uint16(v741)
	v766 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v766+v767*v755)+4)) = uint16(v745)
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v772+v773*v755)+6)) = uint16(v262)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v783 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v778+v779*v755)+8)) = uint8(v783)
	v785 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = v785 + v783
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
	v849 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v943 = v849
	goto L143
L145:
	;
	goto L146
L146:
	;
	v850 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	if v850 == v162 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v895 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v896 = v894 + v890
	if (v892|v896)&int32(_a_F__bt_insertonpg_9) != 0 {
		v943 = v895
		goto L143
	} else {
		goto L155
	}
L148:
	;
	v890 = v882 - v880
	v891 = v881
	v892 = v883 - (v881 + v204) - int32(8)
	v894 = int32(0)
	goto L147
L149:
	;
	v890 = v873 - v871
	v891 = v872
	v892 = v874 - (v872 + v204)
	v894 = v872 + int32(_a_F__bt_insertonpg_8)
	goto L147
L150:
	;
	v852 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v853 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v855 = v852 + (v204 - v853)
	v856 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	v857 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v858 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if v858&int32(1) == int32(0) {
		v871 = v857
		v872 = v857
		v873 = v855
		v874 = v856
		goto L149
	} else {
		goto L153
	}
L151:
	;
	goto L152
L152:
	;
	v863 = *(*int32)(unsafe.Add(mBase, uint32(v165)+36))
	v864 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v866 = v863 + (v204 - v864)
	v867 = int32(0)
	v868 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(v165)+32))
	v870 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if v870 != 0 {
		v880 = v868
		v881 = v867
		v882 = v866
		v883 = v869
		goto L148
	} else {
		goto L154
	}
L153:
	;
	v880 = v857
	v881 = v857
	v882 = v855
	v883 = v856
	goto L148
L154:
	;
	v871 = v868
	v872 = v867
	v873 = v866
	v874 = v869
	goto L149
L155:
	;
	v900 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	if base.Ui32(v900) < base.Ui32(v891) {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v902 = v900
	goto L158
L157:
	;
	v902 = v891
	goto L158
L158:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+44)) = v902
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v905 = int32(10)
	v908 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v904+v895*v905))) = uint16(v908)
	v910 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v911 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v910+v911*v905)+2)) = uint16(v892)
	v916 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v917 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v916+v917*v905)+4)) = uint16(v896)
	v922 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint16)(unsafe.Add(mBase, uint32(v922+v923*v905)+6)) = uint16(v162)
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	*(*uint8)(unsafe.Add(mBase, uint32(v928+v929*v905)+8)) = uint8(v908)
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	v937 = v935 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = v937
	v943 = v937
	goto L143
L159:
	;
	m.G0 = v165 - int32(-64)
	v2123 = v2090 & int32(_a_F__bt_insertonpg_4)
	v2125 = v48 + int32(_a_F__bt_insertonpg_11)
	F_PageInit(m, v2125, int32(_a_F__bt_insertonpg_3), int32(16))
	mBase = m.M
	goto L339
L160:
	;
	v1228 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1231 = v1228 + v943*int32(10)
	v1234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1231-int32(2)))))
	v1237 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1231-int32(4)))))
	v1238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1228)+8)))
	v1239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1228)+6)))
	if int32(0) < v943 {
		goto L205
	} else {
		goto L206
	}
L161:
	;
	v1197 = v1151
	v1226 = float64(0.5)
	goto L160
L162:
	;
	v1151 = int32(0)
	goto L161
L163:
	;
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+29)))
	v946 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	if v946 != int32(1) {
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
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L20
	} else {
		goto L202
	}
L166:
	;
	v1197 = v945
	v1226 = float64(0.7)
	goto L160
L167:
	;
	goto L168
L168:
	;
	if v945&int32(1) != 0 {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	v1197 = int32(1)
	v1226 = base.F64_div(v209, float64(100))
	goto L160
L170:
	;
	goto L171
L171:
	;
	v955 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v955)+192))
	v957 = int32(*(*int16)(unsafe.Add(mBase, uint32(v956)+10)))
	if v957 == int32(1) {
		goto L162
	} else {
		goto L172
	}
L172:
	;
	v960 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	if v960 == int32(2) {
		goto L162
	} else {
		goto L173
	}
L173:
	;
	v963 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	if base.B2i32(v963 != v964)|base.B2i32(base.Ui32(int32(28)) < base.Ui32(v963)) != 0 {
		goto L162
	} else {
		goto L174
	}
L174:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	if v969 != v963*(v227-int32(1)) {
		goto L162
	} else {
		goto L175
	}
L175:
	;
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	if base.Ui32(v182&int32(_a_F__bt_insertonpg_4)) < base.Ui32(v960) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v981 = *(*int32)(unsafe.Add(mBase, uint32(v974+v227<<(uint(int32(2))%32))+20))
	v985 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v986 = F__bt_keep_natts_fast(m, v955, v974+v981&int32(_a_F__bt_insertonpg_1), v985)
	mBase = m.M
	v987 = m.ExcPending
	if v987 != 0 {
		goto L20
	} else {
		goto L179
	}
L177:
	;
	goto L178
L178:
	;
	v1003 = *(*int32)(unsafe.Add(mBase, uint32(v974+(v960-int32(1))&int32(_a_F__bt_insertonpg_4)<<(uint(int32(2))%32))+20))
	v1006 = v974 + v1003&int32(_a_F__bt_insertonpg_1)
	v1007 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+7)))
	if v1007&int32(32) != 0 {
		goto L182
	} else {
		goto L183
	}
L179:
	;
	if v986 < int32(2) {
		goto L162
	} else {
		goto L180
	}
L180:
	;
	if v957 < v986 {
		v1197 = int32(0)
		v1226 = float64(0.5)
		goto L160
	} else {
		goto L181
	}
L181:
	;
	v1197 = int32(1)
	v1226 = base.F64_div(v209, float64(100))
	goto L160
L182:
	;
	v1010 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1006)+5)))
	if v1010&int32(32) != 0 {
		goto L162
	} else {
		goto L185
	}
L183:
	;
	goto L184
L184:
	;
	v1013 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1006)+2)))
	v1014 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1006))))
	v1015 = int32(16)
	v1017 = v1013 | v1014<<(uint(v1015)%32)
	v1018 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1019 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1018))))
	v1022 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1018)+2)))
	v1023 = v1019<<(uint(v1015)%32) | v1022
	if v1017 != v1023 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	goto L184
L186:
	;
	if v1017+int32(1) != v1023 {
		goto L162
	} else {
		goto L189
	}
L187:
	;
	goto L188
L188:
	;
	v1031 = F__bt_keep_natts_fast(m, v955, v1006, v1018)
	mBase = m.M
	v1032 = m.ExcPending
	if v1032 != 0 {
		goto L20
	} else {
		goto L191
	}
L189:
	;
	v1028 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1018)+4)))
	if v1028 != int32(1) {
		goto L162
	} else {
		goto L190
	}
L190:
	;
	goto L188
L191:
	;
	if base.B2i32(v1031 < int32(2))|base.B2i32(v957 < v1031) != 0 {
		goto L162
	} else {
		goto L192
	}
L192:
	;
	v1037 = int32(1)
	v1039 = base.F64_div(v209, float64(100))
	if base.F64_lt(v1039, base.F64_div(base.F64_convert_i32_u(v960), base.F64_convert_i32_u((v182+v1037)&int32(_a_F__bt_insertonpg_4)))) != 0 {
		v1197 = v1037
		v1226 = v1039
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
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1065 = int32(0)
	goto L195
L195:
	;
	v1099 = v1051 + v1065*int32(10)
	v1100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1099)+8)))
	if v1100 != int32(1) {
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
	v1111 = v1065 + int32(1)
	if v1111 != v943 {
		v1065 = v1111
		goto L195
	} else {
		goto L201
	}
L198:
	;
	v1103 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1099)+6)))
	if v162 != v1103 {
		goto L197
	} else {
		goto L199
	}
L199:
	;
	F_pfree(m, v1051)
	mBase = m.M
	v1106 = m.ExcPending
	if v1106 != 0 {
		goto L20
	} else {
		goto L200
	}
L200:
	;
	v1107 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v1107)
	v2090 = v162
	goto L159
L201:
	;
	goto L196
L202:
	;
	v1117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v1117 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_12), v165)
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L20
	} else {
		goto L203
	}
L203:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_13), int32(262), int32(_a_F__bt_insertonpg_14))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
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
	v1260 = int32(0)
	goto L208
L206:
	;
	v1335 = v943
	v1338 = v1228
	v1343 = v946
	goto L207
L207:
	;
	F_pg_qsort(m, v1338, v1335, int32(10), int32(244))
	mBase = m.M
	v1368 = m.ExcPending
	if v1368 != 0 {
		goto L20
	} else {
		goto L215
	}
L208:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1295 = v1292 + v1260*int32(10)
	v1296 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1295)+2)))
	if v1197&int32(1) != 0 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	v1318 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	v1335 = v1316
	v1338 = v1318
	v1343 = v1319
	goto L207
L210:
	;
	v1310 = base.I32_extend16_s(v1307) >> (uint(int32(15)) % 32)
	v1312 = v1310 ^ v1307 - v1310
	*(*uint16)(unsafe.Add(mBase, uint32(v1295))) = uint16(v1312)
	v1315 = v1260 + int32(1)
	v1316 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	if v1315 < v1316 {
		v1260 = v1315
		goto L208
	} else {
		goto L214
	}
L211:
	;
	v1300 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1295)+4)))
	v1307 = base.I32_trunc_sat_f64_s(base.F64_sub(base.F64_mul(v1226, base.F64_convert_i32_s(base.I32_extend16_s(v1296))), base.F64_mul(base.F64_sub(float64(1), v1226), base.F64_convert_i32_s(v1300))))
	goto L210
L212:
	;
	goto L213
L213:
	;
	v1305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1295)+4)))
	v1307 = v1296 - v1305
	goto L210
L214:
	;
	goto L209
L215:
	;
	if v1335 < int32(2) {
		v1467 = v1335
		goto L216
	} else {
		goto L217
	}
L216:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+60)) = v1467
	v1496 = int32(1)
	if v1343&v1496 == int32(0) {
		goto L227
	} else {
		goto L228
	}
L217:
	;
	v1371 = int32(1)
	v1372 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1338)+4)))
	if v1343&v1371 != 0 {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	v1377 = float64(0.05)
	goto L220
L219:
	;
	v1377 = float64(0.075)
	goto L220
L220:
	;
	v1378 = *(*int32)(unsafe.Add(mBase, uint32(v165)+40))
	v1381 = base.I32_trunc_sat_f64_s(base.F64_mul(v1377, base.F64_convert_i32_s(v1378)))
	v1384 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1338)+2)))
	v1408 = v1371
	goto L221
L221:
	;
	v1438 = v1338 + v1408*int32(10)
	v1439 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1438)+2)))
	if v1439 < base.I32_extend16_s(v1384-v1381) {
		v1467 = v1408
		goto L216
	} else {
		goto L223
	}
L222:
	;
	v1467 = v1335
	goto L216
L223:
	;
	v1442 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1438)+4)))
	if base.B2i32(base.I32_extend16_s(v1384+v1381) < v1439)|base.B2i32(v1442 < base.I32_extend16_s(v1372-v1381))|base.B2i32(base.I32_extend16_s(v1372+v1381) < v1442) != 0 {
		v1467 = v1408
		goto L216
	} else {
		goto L224
	}
L224:
	;
	v1448 = v1408 + int32(1)
	if v1448 != v1335 {
		v1408 = v1448
		goto L221
	} else {
		goto L225
	}
L225:
	;
	goto L222
L226:
	;
	if v1842 < v1840 {
		goto L297
	} else {
		goto L298
	}
L227:
	;
	v1501 = *(*int32)(unsafe.Add(mBase, uint32(v165)+44))
	v1825 = v1501
	v1840 = v1335
	v1842 = v1467
	v1843 = v1338
	v1847 = v1496
	goto L226
L228:
	;
	goto L229
L229:
	;
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1502)+192))
	v1504 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1503)+10)))
	if v1467 < v1335 {
		goto L230
	} else {
		goto L231
	}
L230:
	;
	v1506 = v1467
	goto L232
L231:
	;
	v1506 = v1335
	goto L232
L232:
	;
	v1507 = int32(0)
	v1508 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1338)+6)))
	v1523 = v1506
	v1524 = v1507
	v1526 = v1507
	goto L233
L233:
	;
	v1556 = v1523 - int32(1)
	v1559 = v1338 + v1556*int32(10)
	v1560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1559)+6)))
	if base.Ui32(v1560) < base.Ui32(v1508) {
		goto L243
	} else {
		goto L244
	}
L234:
	;
	v1598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1587)+6)))
	v1599 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1588)+8)))
	v1602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1588)+6)))
	v1603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	if base.B2i32(v1599 != int32(1))|base.B2i32(v1602 != v1603) == int32(0) {
		goto L263
	} else {
		goto L264
	}
L235:
	;
	if int32(0) < v1556 {
		goto L258
	} else {
		goto L259
	}
L236:
	;
	v1587 = v1526
	v1588 = v1559
	goto L235
L237:
	;
	v1583 = int32(0)
	if v1581 == v1583 {
		v1523 = v1556
		v1524 = v1583
		v1526 = v1582
		goto L233
	} else {
		goto L257
	}
L238:
	;
	v1581 = v1579
	v1582 = v1559
	goto L237
L239:
	;
	if v1526 != 0 {
		v1581 = v1524
		v1582 = v1526
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
	if v1563&int32(1) == int32(0) {
		goto L239
	} else {
		goto L251
	}
L242:
	;
	v1587 = v1526
	v1588 = v1524
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
	if base.Ui32(v1508) < base.Ui32(v1560) {
		goto L239
	} else {
		goto L247
	}
L246:
	;
	goto L236
L247:
	;
	v1563 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1338)+8)))
	v1564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1559)+8)))
	if v1564 != 0 {
		goto L241
	} else {
		goto L248
	}
L248:
	;
	if v1563&int32(1) == int32(0) {
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
	v1576 = v1524
	goto L254
L253:
	;
	v1576 = v1559
	goto L254
L254:
	;
	if v1526 != 0 {
		v1587 = v1526
		v1588 = v1576
		goto L235
	} else {
		goto L255
	}
L255:
	;
	v1579 = v1576
	goto L238
L256:
	;
	v1579 = v1524
	goto L238
L257:
	;
	v1587 = v1582
	v1588 = v1581
	goto L235
L258:
	;
	v1593 = int32(0)
	if v1587 == v1593 {
		v1523 = v1556
		v1524 = v1588
		v1526 = v1593
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
	v1623 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1587)+8)))
	if v1623|base.B2i32(v1598 != v1603) == int32(0) {
		goto L267
	} else {
		goto L268
	}
L263:
	;
	v1608 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1622 = v1608
	goto L262
L264:
	;
	goto L265
L265:
	;
	v1609 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1617 = *(*int32)(unsafe.Add(mBase, uint32(v1609+(v1602-int32(1))&int32(_a_F__bt_insertonpg_4)<<(uint(int32(2))%32))+20))
	v1622 = v1609 + v1617&int32(_a_F__bt_insertonpg_1)
	goto L262
L266:
	;
	v1639 = F__bt_keep_natts_fast(m, v1502, v1622, v1638)
	mBase = m.M
	v1640 = m.ExcPending
	if v1640 != 0 {
		goto L20
	} else {
		goto L270
	}
L267:
	;
	v1628 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1638 = v1628
	goto L266
L268:
	;
	goto L269
L269:
	;
	v1629 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1633 = *(*int32)(unsafe.Add(mBase, uint32(v1629+v1598<<(uint(int32(2))%32))+20))
	v1638 = v1629 + v1633&int32(_a_F__bt_insertonpg_1)
	goto L266
L270:
	;
	if v1639 <= v1504 {
		v1825 = v1639
		v1840 = v1335
		v1842 = v1467
		v1843 = v1338
		v1847 = v1496
		goto L226
	} else {
		goto L271
	}
L271:
	;
	v1644 = int32(0)
	if base.B2i32(v1238&int32(1) == v1644)|base.B2i32(v1603 != v1239) == v1644 {
		goto L273
	} else {
		goto L274
	}
L272:
	;
	if v1234&int32(1)|base.B2i32(v1603 != v1237) == int32(0) {
		goto L277
	} else {
		goto L278
	}
L273:
	;
	v1650 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1664 = v1650
	goto L272
L274:
	;
	goto L275
L275:
	;
	v1651 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1659 = *(*int32)(unsafe.Add(mBase, uint32(v1651+(v1239-int32(1))&int32(_a_F__bt_insertonpg_4)<<(uint(int32(2))%32))+20))
	v1664 = v1651 + v1659&int32(_a_F__bt_insertonpg_1)
	goto L272
L276:
	;
	v1682 = F__bt_keep_natts_fast(m, v1502, v1664, v1681)
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L20
	} else {
		goto L280
	}
L277:
	;
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1681 = v1671
	goto L276
L278:
	;
	goto L279
L279:
	;
	v1672 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1676 = *(*int32)(unsafe.Add(mBase, uint32(v1672+v1237<<(uint(int32(2))%32))+20))
	v1681 = v1672 + v1676&int32(_a_F__bt_insertonpg_1)
	goto L276
L280:
	;
	if v1504 < v1682 {
		goto L281
	} else {
		goto L282
	}
L281:
	;
	v1685 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+29)))
	if v1685 == int32(0) {
		goto L284
	} else {
		goto L285
	}
L282:
	;
	goto L283
L283:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+60)) = v1335
	v1825 = v1504
	v1840 = v1335
	v1842 = v1335
	v1843 = v1338
	v1847 = int32(0)
	goto L226
L284:
	;
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1689 = *(*int32)(unsafe.Add(mBase, uint32(v1688)+24))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1694 = F__bt_keep_natts_fast(m, v1502, v1688+v1689&int32(_a_F__bt_insertonpg_1), v1693)
	mBase = m.M
	v1695 = m.ExcPending
	if v1695 != 0 {
		goto L20
	} else {
		goto L287
	}
L285:
	;
	v1697 = v1682
	goto L286
L286:
	;
	if int32(0) < v1335 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	if v1504 < v1694 {
		v1825 = v1694
		v1840 = v1335
		v1842 = v1467
		v1843 = v1338
		v1847 = v1496
		goto L226
	} else {
		goto L288
	}
L288:
	;
	v1697 = v1694
	goto L286
L289:
	;
	v1715 = int32(0)
	goto L292
L290:
	;
	v1787 = v1335
	v1790 = v1338
	goto L291
L291:
	;
	F_pg_qsort(m, v1790, v1787, int32(10), int32(244))
	mBase = m.M
	v1820 = m.ExcPending
	if v1820 != 0 {
		goto L20
	} else {
		goto L295
	}
L292:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1750 = v1747 + v1715*int32(10)
	v1751 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1750)+2)))
	v1755 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1750)+4)))
	v1760 = base.I32_trunc_sat_f64_s(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v1751), float64(0.96)), base.F64_mul(base.F64_convert_i32_s(v1755), float64(-0.040000000000000036))))
	v1763 = base.I32_extend16_s(v1760) >> (uint(int32(15)) % 32)
	v1765 = v1760 ^ v1763 - v1763
	*(*uint16)(unsafe.Add(mBase, uint32(v1750))) = uint16(v1765)
	v1768 = v1715 + int32(1)
	v1769 = *(*int32)(unsafe.Add(mBase, uint32(v165)+52))
	if v1768 < v1769 {
		v1715 = v1768
		goto L292
	} else {
		goto L294
	}
L293:
	;
	v1771 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	v1787 = v1769
	v1790 = v1771
	goto L291
L294:
	;
	goto L293
L295:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+60)) = int32(1)
	v1825 = v1697
	v1840 = v1787
	v1842 = int32(1)
	v1843 = v1790
	v1847 = v1496
	goto L226
L296:
	;
	v2052 = v1843 + v2020*int32(10)
	v2053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+29)))
	if (v1847|v2053)&int32(1) != 0 {
		v2068 = v2052
		goto L329
	} else {
		goto L330
	}
L297:
	;
	v1871 = v1842
	goto L299
L298:
	;
	v1871 = v1840
	goto L299
L299:
	;
	if v1871 <= int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v2020 = int32(0)
	goto L296
L301:
	;
	goto L302
L302:
	;
	v1875 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	v1877 = v1875 + int32(20)
	v1878 = int32(0)
	v1880 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
	v1881 = *(*int32)(unsafe.Add(mBase, uint32(v165)+20))
	v1882 = *(*int32)(unsafe.Add(mBase, uint32(v165)+24))
	v1883 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v165)+28)))
	v1886 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	v1898 = int32(2147483647)
	v1902 = v1878
	v1903 = v1878
	goto L303
L303:
	;
	v1935 = v1843 + v1902*int32(10)
	v1936 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1935)+6)))
	v1937 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1935)+8)))
	if v1883&int32(1) == int32(0) {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	v2020 = v1999
	goto L296
L305:
	;
	v1998 = base.B2i32(v1997 < v1898)
	if v1997 < v1898 {
		goto L321
	} else {
		goto L322
	}
L306:
	;
	if v1937&int32(1) == int32(0) {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	goto L308
L308:
	;
	if v1937&int32(1) != 0 {
		goto L315
	} else {
		goto L316
	}
L309:
	;
	if v1936 == v1886 {
		v1997 = v1882
		goto L305
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v1948 = *(*int32)(unsafe.Add(mBase, uint32(v1877+v1936<<(uint(int32(2))%32))))
	v1997 = (int32(base.Ui32(v1948)>>(uint(int32(17))%32))+int32(7))&int32(_a_F__bt_insertonpg_6) | int32(4)
	goto L305
L312:
	;
	goto L311
L313:
	;
	v1993 = F__bt_keep_natts_fast(m, v1880, v1991, v1992)
	mBase = m.M
	v1994 = m.ExcPending
	if v1994 != 0 {
		goto L20
	} else {
		goto L320
	}
L314:
	;
	v1987 = *(*int32)(unsafe.Add(mBase, uint32(v1877+v1936<<(uint(int32(2))%32))))
	v1991 = v1983
	v1992 = v1875 + v1987&int32(_a_F__bt_insertonpg_1)
	goto L313
L315:
	;
	if v1936 == v1886 {
		v1983 = v1881
		goto L314
	} else {
		goto L318
	}
L316:
	;
	goto L317
L317:
	;
	v1978 = *(*int32)(unsafe.Add(mBase, uint32(v1877+(v1936-int32(1))&int32(_a_F__bt_insertonpg_4)<<(uint(int32(2))%32))))
	v1981 = v1875 + v1978&int32(_a_F__bt_insertonpg_1)
	if v1936 == v1886 {
		v1991 = v1981
		v1992 = v1881
		goto L313
	} else {
		goto L319
	}
L318:
	;
	v1967 = *(*int32)(unsafe.Add(mBase, uint32(v1877+(v1936-int32(1))&int32(_a_F__bt_insertonpg_4)<<(uint(int32(2))%32))))
	v1983 = v1875 + v1967&int32(_a_F__bt_insertonpg_1)
	goto L314
L319:
	;
	v1983 = v1981
	goto L314
L320:
	;
	v1997 = v1993
	goto L305
L321:
	;
	v1999 = v1902
	goto L323
L322:
	;
	v1999 = v1903
	goto L323
L323:
	;
	if v1997 <= v1825 {
		v2020 = v1999
		goto L296
	} else {
		goto L324
	}
L324:
	;
	if v1997 < v1898 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v2001 = v1997
	goto L327
L326:
	;
	v2001 = v1898
	goto L327
L327:
	;
	v2003 = v1902 + int32(1)
	if v2003 != v1871 {
		v1898 = v2001
		v1902 = v2003
		v1903 = v1999
		goto L303
	} else {
		goto L328
	}
L328:
	;
	goto L304
L329:
	;
	v2069 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2068)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v176))) = uint8(v2069)
	v2071 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2068)+6)))
	F_pfree(m, v1843)
	mBase = m.M
	v2073 = m.ExcPending
	if v2073 != 0 {
		goto L20
	} else {
		goto L338
	}
L330:
	;
	v2057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2052)+8)))
	if v2057 != 0 {
		v2068 = v2052
		goto L329
	} else {
		goto L331
	}
L331:
	;
	v2058 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2052)+6)))
	v2059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+30)))
	if base.Ui32(v2058) < base.Ui32(v2059+int32(9)) {
		goto L332
	} else {
		goto L333
	}
L332:
	;
	v2063 = v1843
	goto L334
L333:
	;
	v2063 = v2052
	goto L334
L334:
	;
	if base.Ui32(v2059) <= base.Ui32(v2058) {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v2065 = v2063
	goto L337
L336:
	;
	v2065 = v2052
	goto L337
L337:
	;
	v2068 = v2065
	goto L329
L338:
	;
	v2090 = v2071
	goto L159
L339:
	;
	v2129 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[4]))))
	v2130 = v2125 + v2129
	v2131 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v2135 = v2131&int32(_a_F__bt_insertonpg_15) | int32(128)
	*(*uint16)(unsafe.Add(mBase, uint32(v2130)+12)) = uint16(v2135)
	v2137 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
	*(*int32)(unsafe.Add(mBase, uint32(v2130))) = v2137
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2130)+8)) = v2139
	v2141 = *(*int64)(unsafe.Add(mBase, uint32(v136)))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[5]))) = v2141
	if v10&int32(_a_F__bt_insertonpg_4) != 0 {
		goto L340
	} else {
		goto L341
	}
L340:
	;
	v2148 = v102 - int32(1)
	goto L342
L341:
	;
	v2148 = int32(0)
	goto L342
L342:
	;
	v2150 = v140 & int32(1)
	v2151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+207)))
	if v2151 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	if v2150 != 0 {
		goto L351
	} else {
		goto L352
	}
L344:
	;
	if v2123 == v162 {
		v2170 = v101
		v2171 = l7
		goto L343
	} else {
		goto L347
	}
L345:
	;
	goto L346
L346:
	;
	v2159 = *(*int32)(unsafe.Add(mBase, uint32(v136+v2123<<(uint(int32(2))%32))+20))
	if v2148&int32(_a_F__bt_insertonpg_4) == v2123 {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	goto L346
L348:
	;
	v2166 = v105
	goto L350
L349:
	;
	v2166 = v136 + v2159&int32(_a_F__bt_insertonpg_1)
	goto L350
L350:
	;
	v2170 = v2166
	v2171 = int32(base.Ui32(v2159) >> (uint(int32(17)) % 32))
	goto L343
L351:
	;
	if v2151 != 0 {
		goto L355
	} else {
		goto L356
	}
L352:
	;
	v2198 = v2171
	v2199 = v2170
	goto L353
L353:
	;
	v2204 = F_PageAddItemExtended(m, v48+int32(_a_F__bt_insertonpg_11), v2199, v2198, int32(1), int32(0))
	mBase = m.M
	v2205 = m.ExcPending
	if v2205 != 0 {
		goto L20
	} else {
		goto L363
	}
L354:
	;
	v2192 = F__bt_truncate(m, l0, v2191, v2170, l2)
	mBase = m.M
	v2193 = m.ExcPending
	if v2193 != 0 {
		goto L20
	} else {
		goto L362
	}
L355:
	;
	if v102&int32(_a_F__bt_insertonpg_4) == v2123 {
		v2191 = v101
		goto L354
	} else {
		goto L358
	}
L356:
	;
	goto L357
L357:
	;
	v2177 = int32(_a_F__bt_insertonpg_4)
	v2178 = (v2123 - int32(1)) & v2177
	v2182 = *(*int32)(unsafe.Add(mBase, uint32(v136+v2178<<(uint(int32(2))%32))+20))
	if v2148&v2177 == v2178 {
		goto L359
	} else {
		goto L360
	}
L358:
	;
	goto L357
L359:
	;
	v2189 = v105
	goto L361
L360:
	;
	v2189 = v136 + v2182&int32(_a_F__bt_insertonpg_1)
	goto L361
L361:
	;
	v2191 = v2189
	goto L354
L362:
	;
	v2194 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2192)+6)))
	v2198 = v2194 & int32(_a_F__bt_insertonpg_7)
	v2199 = v2192
	goto L353
L363:
	;
	if v2204 == int32(0) {
		goto L13
	} else {
		goto L364
	}
L364:
	;
	v2208 = F__bt_allocbuf(m, l0, l1)
	mBase = m.M
	v2209 = m.ExcPending
	if v2209 != 0 {
		goto L20
	} else {
		goto L366
	}
L365:
	;
	v2238 = int32(0)
	base.MemoryFill(m, v2237, v2238, int32(_a_F__bt_insertonpg_3))
	if v2208 < v2238 {
		goto L371
	} else {
		goto L372
	}
L366:
	;
	if v2208 < int32(0) {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v2215 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2220 = v2215 + (v2208^int32(-1))<<(uint(int32(2))%32)
	v2221 = *(*int32)(unsafe.Add(mBase, uint32(v2220)))
	base.MemoryCopy(m, v48+int32(208), v2221, int32(_a_F__bt_insertonpg_3))
	v2224 = *(*int32)(unsafe.Add(mBase, uint32(v2220)))
	v2237 = v2224
	goto L365
L368:
	;
	goto L369
L369:
	;
	v2228 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2233 = v2228 + v2208<<(uint(int32(13))%32) + int32(-8192)
	base.MemoryCopy(m, v48+int32(208), v2233, int32(_a_F__bt_insertonpg_3))
	v2237 = v2233
	goto L365
L370:
	;
	v2260 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+224)))
	*(*int32)(unsafe.Add(mBase, uint32(v2130)+4)) = v2259
	v2264 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6]))
	v2268 = F_LWLockAcquire(m, v2264+int32(2560), int32(1))
	mBase = m.M
	v2269 = m.ExcPending
	if v2269 != 0 {
		goto L20
	} else {
		goto L374
	}
L371:
	;
	v2244 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v2244+(v2208^int32(-1))*int32(56))+16))
	v2259 = v2250
	goto L370
L372:
	;
	goto L373
L373:
	;
	v2252 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v2253 = int32(56)
	v2258 = *(*int32)(unsafe.Add(mBase, uint32(v2252+v2208*v2253-v2253)+16))
	v2259 = v2258
	goto L370
L374:
	;
	v2270 = int32(0)
	v2272 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[7]))
	v2273 = *(*int32)(unsafe.Add(mBase, uint32(v2272)+4))
	if v2273 <= v2270 {
		v2382 = v2270
		goto L375
	} else {
		goto L376
	}
L375:
	;
	v2384 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[6]))
	F_LWLockRelease(m, v2384+int32(2560))
	mBase = m.M
	v2388 = m.ExcPending
	if v2388 != 0 {
		goto L20
	} else {
		goto L383
	}
L376:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
	v2281 = int32(0)
	goto L377
L377:
	;
	v2326 = v2272 + int32(12) + v2281*int32(12)
	v2327 = *(*int32)(unsafe.Add(mBase, uint32(v2326)))
	if v2327 != v2278 {
		goto L379
	} else {
		goto L380
	}
L378:
	;
	v2382 = int32(0)
	goto L375
L379:
	;
	v2334 = v2281 + int32(1)
	if v2334 != v2273 {
		v2281 = v2334
		goto L377
	} else {
		goto L382
	}
L380:
	;
	v2329 = *(*int32)(unsafe.Add(mBase, uint32(v2326)+4))
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(l0)+64))
	if v2329 != v2330 {
		goto L379
	} else {
		goto L381
	}
L381:
	;
	v2332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2326)+8)))
	v2382 = v2332
	goto L375
L382:
	;
	goto L378
L383:
	;
	v2390 = v2382 & int32(_a_F__bt_insertonpg_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v2130)+14)) = uint16(v2390)
	v2392 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v138)+12)))
	v2394 = v48 + int32(208)
	v2395 = v2260 + v2394
	*(*int32)(unsafe.Add(mBase, uint32(v2395))) = v160
	v2398 = v2392 & int32(_a_F__bt_insertonpg_16)
	*(*uint16)(unsafe.Add(mBase, uint32(v2395)+12)) = uint16(v2398)
	v2400 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2395)+4)) = v2400
	v2402 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v2395)+14)) = uint16(v2390)
	*(*int32)(unsafe.Add(mBase, uint32(v2395)+8)) = v2402
	if v139 != 0 {
		goto L384
	} else {
		goto L385
	}
L384:
	;
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v2415 = F_PageAddItemExtended(m, v2394, v136+v2407&int32(_a_F__bt_insertonpg_1), int32(base.Ui32(v2407)>>(uint(int32(17))%32)), int32(1), int32(0))
	mBase = m.M
	v2416 = m.ExcPending
	if v2416 != 0 {
		goto L20
	} else {
		goto L387
	}
L385:
	;
	v2420 = int32(1)
	goto L386
L386:
	;
	if v2150 != 0 {
		goto L389
	} else {
		goto L390
	}
L387:
	;
	if v2415 == int32(0) {
		goto L12
	} else {
		goto L388
	}
L388:
	;
	v2420 = int32(2)
	goto L386
L389:
	;
	v2422 = int32(0)
	goto L391
L390:
	;
	v2422 = v2420
	goto L391
L391:
	;
	v2425 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	if v2425 != 0 {
		goto L392
	} else {
		goto L393
	}
L392:
	;
	v2426 = int32(2)
	goto L394
L393:
	;
	v2426 = int32(1)
	goto L394
L394:
	;
	if base.Ui32(int32(25)) <= base.Ui32(v141) {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v2434 = int32(base.Ui32(v141+int32(_a_F__bt_insertonpg_5)) >> (uint(int32(2)) % 32))
	goto L397
L396:
	;
	v2434 = int32(0)
	goto L397
L397:
	;
	v2436 = v2434 & int32(_a_F__bt_insertonpg_4)
	if base.Ui32(v2426) <= base.Ui32(v2436) {
		goto L398
	} else {
		goto L399
	}
L398:
	;
	v2443 = int32(2)
	v2451 = v2426
	v2455 = v2420
	goto L401
L399:
	;
	v2595 = v2426
	v2599 = v2420
	goto L400
L400:
	;
	v2630 = int32(_a_F__bt_insertonpg_4)
	if base.Ui32(v2595&v2630) <= base.Ui32(v102&v2630) {
		goto L430
	} else {
		goto L431
	}
L401:
	;
	v2486 = int32(_a_F__bt_insertonpg_4)
	v2487 = v2451 & v2486
	v2491 = *(*int32)(unsafe.Add(mBase, uint32(v136+int32(20)+v2487<<(uint(int32(2))%32))))
	if v2487 == v2148&v2486 {
		goto L404
	} else {
		goto L405
	}
L402:
	;
	v2595 = v2581
	v2599 = v2578
	goto L400
L403:
	;
	v2542 = int32(base.Ui32(v2491) >> (uint(int32(17)) % 32))
	if base.Ui32(v2487) < base.Ui32(v2123) {
		goto L419
	} else {
		goto L420
	}
L404:
	;
	v2536 = v105
	v2537 = v2443
	v2538 = v2455
	goto L403
L405:
	;
	goto L406
L406:
	;
	v2497 = v136 + v2491&int32(_a_F__bt_insertonpg_1)
	if v2487 != v102&int32(_a_F__bt_insertonpg_4) {
		v2536 = v2497
		v2537 = v2443
		v2538 = v2455
		goto L403
	} else {
		goto L407
	}
L407:
	;
	v2501 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+207)))
	if v2501 == int32(1) {
		goto L408
	} else {
		goto L409
	}
L408:
	;
	v2509 = F_PageAddItemExtended(m, v48+int32(_a_F__bt_insertonpg_11), v101, l7, v2443&int32(_a_F__bt_insertonpg_4), int32(0))
	mBase = m.M
	v2510 = m.ExcPending
	if v2510 != 0 {
		goto L20
	} else {
		goto L411
	}
L409:
	;
	goto L410
L410:
	;
	v2518 = v2455 & int32(_a_F__bt_insertonpg_4)
	if v2518 == v2422 {
		goto L413
	} else {
		goto L414
	}
L411:
	;
	if v2509 == int32(0) {
		goto L11
	} else {
		goto L412
	}
L412:
	;
	v2536 = v2497
	v2537 = v2443 + int32(1)
	v2538 = v2455
	goto L403
L413:
	;
	v2520 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+192)) = v2520
	*(*int32)(unsafe.Add(mBase, uint32(v48)+196)) = int32(537395200)
	v2527 = int32(8)
	v2528 = v48 + int32(192)
	goto L415
L414:
	;
	v2527 = l7
	v2528 = v101
	goto L415
L415:
	;
	v2530 = F_PageAddItemExtended(m, v48+int32(208), v2528, v2527, v2518, int32(0))
	mBase = m.M
	v2531 = m.ExcPending
	if v2531 != 0 {
		goto L20
	} else {
		goto L416
	}
L416:
	;
	if v2530 == int32(0) {
		goto L10
	} else {
		goto L417
	}
L417:
	;
	v2536 = v2497
	v2537 = v2443
	v2538 = v2455 + int32(1)
	goto L403
L418:
	;
	v2581 = v2451 + int32(1)
	if base.Ui32(v2581&int32(_a_F__bt_insertonpg_4)) <= base.Ui32(v2436) {
		v2443 = v2576
		v2451 = v2581
		v2455 = v2578
		goto L401
	} else {
		goto L429
	}
L419:
	;
	v2549 = F_PageAddItemExtended(m, v48+int32(_a_F__bt_insertonpg_11), v2536, v2542, v2537&int32(_a_F__bt_insertonpg_4), int32(0))
	mBase = m.M
	v2550 = m.ExcPending
	if v2550 != 0 {
		goto L20
	} else {
		goto L422
	}
L420:
	;
	goto L421
L421:
	;
	v2558 = v2538 & int32(_a_F__bt_insertonpg_4)
	if v2558 == v2422 {
		goto L424
	} else {
		goto L425
	}
L422:
	;
	if v2549 == int32(0) {
		goto L9
	} else {
		goto L423
	}
L423:
	;
	v2576 = v2537 + int32(1)
	v2578 = v2538
	goto L418
L424:
	;
	v2560 = *(*int64)(unsafe.Add(mBase, uint32(v2536)))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+192)) = v2560
	*(*int32)(unsafe.Add(mBase, uint32(v48)+196)) = int32(537395200)
	v2567 = int32(8)
	v2568 = v48 + int32(192)
	goto L426
L425:
	;
	v2567 = v2542
	v2568 = v2536
	goto L426
L426:
	;
	v2570 = F_PageAddItemExtended(m, v48+int32(208), v2568, v2567, v2558, int32(0))
	mBase = m.M
	v2571 = m.ExcPending
	if v2571 != 0 {
		goto L20
	} else {
		goto L427
	}
L427:
	;
	if v2570 == int32(0) {
		goto L8
	} else {
		goto L428
	}
L428:
	;
	v2576 = v2537
	v2578 = v2538 + int32(1)
	goto L418
L429:
	;
	goto L402
L430:
	;
	v2638 = v2599 & int32(_a_F__bt_insertonpg_4)
	if v2638 == v2422 {
		goto L433
	} else {
		goto L434
	}
L431:
	;
	goto L432
L432:
	;
	if v139 == int32(0) {
		goto L439
	} else {
		goto L440
	}
L433:
	;
	v2640 = *(*int64)(unsafe.Add(mBase, uint32(v101)))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+192)) = v2640
	*(*int32)(unsafe.Add(mBase, uint32(v48)+196)) = int32(537395200)
	v2647 = int32(8)
	v2648 = v48 + int32(192)
	goto L435
L434:
	;
	v2647 = l7
	v2648 = v101
	goto L435
L435:
	;
	v2650 = F_PageAddItemExtended(m, v48+int32(208), v2648, v2647, v2638, int32(0))
	mBase = m.M
	v2651 = m.ExcPending
	if v2651 != 0 {
		goto L20
	} else {
		goto L436
	}
L436:
	;
	if v2650 == int32(0) {
		goto L7
	} else {
		goto L437
	}
L437:
	;
	goto L432
L438:
	;
	v2697 = int32(_a_F__bt_insertonpg_17)
	v2699 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[8]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[8])) = v2699 + int32(1)
	base.MemoryCopy(m, v136, v48+int32(_a_F__bt_insertonpg_11), int32(_a_F__bt_insertonpg_3))
	if v2208 < int32(0) {
		goto L450
	} else {
		goto L451
	}
L439:
	;
	v2658 = int32(0)
	v2694 = v2658
	v2695 = v2658
	v2696 = v2658
	goto L438
L440:
	;
	goto L441
L441:
	;
	v2661 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v2663 = F__bt_getbuf(m, l0, v2661, int32(3))
	mBase = m.M
	v2664 = m.ExcPending
	if v2664 != 0 {
		goto L20
	} else {
		goto L443
	}
L442:
	;
	v2683 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2682)+16)))
	v2684 = v2683 + v2682
	v2685 = *(*int32)(unsafe.Add(mBase, uint32(v2684)))
	if v2685 != v160 {
		goto L6
	} else {
		goto L447
	}
L443:
	;
	if v2663 < int32(0) {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	v2668 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2674 = *(*int32)(unsafe.Add(mBase, uint32(v2668+(v2663^int32(-1))<<(uint(int32(2))%32))))
	v2682 = v2674
	goto L442
L445:
	;
	goto L446
L446:
	;
	v2676 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2682 = v2676 + v2663<<(uint(int32(13))%32) + int32(-8192)
	goto L442
L447:
	;
	v2687 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2684)+14)))
	v2688 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2395)+14)))
	if v2687 == v2688 {
		v2694 = v2684
		v2695 = v2663
		v2696 = v2682
		goto L438
	} else {
		goto L448
	}
L448:
	;
	v2690 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2395)+12)))
	v2692 = v2690 | int32(32)
	*(*uint16)(unsafe.Add(mBase, uint32(v2395)+12)) = uint16(v2692)
	v2694 = v2684
	v2695 = v2663
	v2696 = v2682
	goto L438
L449:
	;
	base.MemoryCopy(m, v2724, v48+int32(208), int32(_a_F__bt_insertonpg_3))
	v2729 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2724)+16)))
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v2731 = m.ExcPending
	if v2731 != 0 {
		goto L20
	} else {
		goto L453
	}
L450:
	;
	v2710 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2716 = *(*int32)(unsafe.Add(mBase, uint32(v2710+(v2208^int32(-1))<<(uint(int32(2))%32))))
	v2724 = v2716
	goto L449
L451:
	;
	goto L452
L452:
	;
	v2718 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2724 = v2718 + v2208<<(uint(int32(13))%32) + int32(-8192)
	goto L449
L453:
	;
	F_MarkBufferDirty(m, v2208)
	mBase = m.M
	v2733 = m.ExcPending
	if v2733 != 0 {
		goto L20
	} else {
		goto L454
	}
L454:
	;
	if v139 != 0 {
		goto L455
	} else {
		goto L456
	}
L455:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2694))) = v2259
	F_MarkBufferDirty(m, v2695)
	mBase = m.M
	v2736 = m.ExcPending
	if v2736 != 0 {
		goto L20
	} else {
		goto L458
	}
L456:
	;
	goto L457
L457:
	;
	if v2150 == int32(0) {
		goto L459
	} else {
		goto L460
	}
L458:
	;
	goto L457
L459:
	;
	if l4 < int32(0) {
		goto L463
	} else {
		goto L464
	}
L460:
	;
	goto L461
L461:
	;
	v2766 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v2767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2766)+118)))
	if v2767 != int32(112) {
		goto L468
	} else {
		goto L469
	}
L462:
	;
	v2757 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2756)+16)))
	v2758 = v2757 + v2756
	v2759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2758)+12)))
	v2761 = v2759 & int32(_a_F__bt_insertonpg_18)
	*(*uint16)(unsafe.Add(mBase, uint32(v2758)+12)) = uint16(v2761)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v2764 = m.ExcPending
	if v2764 != 0 {
		goto L20
	} else {
		goto L466
	}
L463:
	;
	v2742 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2748 = *(*int32)(unsafe.Add(mBase, uint32(v2742+(l4^int32(-1))<<(uint(int32(2))%32))))
	v2756 = v2748
	goto L462
L464:
	;
	goto L465
L465:
	;
	v2750 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2756 = v2750 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L462
L466:
	;
	goto L461
L467:
	;
	v2862 = base.I64_rotl(v2860, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v136))) = v2862
	*(*int64)(unsafe.Add(mBase, uint32(v2724))) = v2862
	if v139 != 0 {
		goto L513
	} else {
		goto L514
	}
L468:
	;
	v2855 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v2856 = m.ExcPending
	if v2856 != 0 {
		goto L20
	} else {
		goto L512
	}
L469:
	;
	v2771 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[9]))
	if v2771 <= int32(0) {
		goto L470
	} else {
		goto L471
	}
L470:
	;
	v2774 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2774 != 0 {
		goto L468
	} else {
		goto L473
	}
L471:
	;
	goto L472
L472:
	;
	v2777 = *(*int32)(unsafe.Add(mBase, uint32(v2729+v2724)+8))
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+198)) = uint16(v102)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+192)) = v2777
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+196)) = uint16(v2123)
	if base.Ui32(v2148&int32(_a_F__bt_insertonpg_4)) < base.Ui32(v2123) {
		goto L475
	} else {
		goto L476
	}
L473:
	;
	v2775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v2775 != 0 {
		goto L468
	} else {
		goto L474
	}
L474:
	;
	goto L472
L475:
	;
	v2785 = v10
	goto L477
L476:
	;
	v2785 = int32(0)
	goto L477
L477:
	;
	if v10&int32(_a_F__bt_insertonpg_4) != 0 {
		goto L478
	} else {
		goto L479
	}
L478:
	;
	v2789 = v2785
	goto L480
L479:
	;
	v2789 = int32(0)
	goto L480
L480:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+200)) = uint16(v2789)
	F_XLogBeginInsert(m)
	mBase = m.M
	v2792 = m.ExcPending
	if v2792 != 0 {
		goto L20
	} else {
		goto L481
	}
L481:
	;
	F_XLogRegisterData(m, v48+int32(192), int32(10))
	mBase = m.M
	v2797 = m.ExcPending
	if v2797 != 0 {
		goto L20
	} else {
		goto L482
	}
L482:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v2801 = m.ExcPending
	if v2801 != 0 {
		goto L20
	} else {
		goto L483
	}
L483:
	;
	F_XLogRegisterBuffer(m, int32(1), v2208, int32(6))
	mBase = m.M
	v2805 = m.ExcPending
	if v2805 != 0 {
		goto L20
	} else {
		goto L484
	}
L484:
	;
	if v139 != 0 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	F_XLogRegisterBuffer(m, int32(2), v2695, int32(8))
	mBase = m.M
	v2809 = m.ExcPending
	if v2809 != 0 {
		goto L20
	} else {
		goto L488
	}
L486:
	;
	goto L487
L487:
	;
	if v2150 == int32(0) {
		goto L489
	} else {
		goto L490
	}
L488:
	;
	goto L487
L489:
	;
	F_XLogRegisterBuffer(m, int32(3), l4, int32(8))
	mBase = m.M
	v2815 = m.ExcPending
	if v2815 != 0 {
		goto L20
	} else {
		goto L492
	}
L490:
	;
	goto L491
L491:
	;
	v2816 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+207)))
	v2817 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v48)+200)))
	if v2816|v2817 != 0 {
		goto L493
	} else {
		goto L494
	}
L492:
	;
	goto L491
L493:
	;
	if v2817 != 0 {
		goto L496
	} else {
		goto L497
	}
L494:
	;
	goto L495
L495:
	;
	if v2150 == int32(0) {
		goto L503
	} else {
		goto L504
	}
L496:
	;
	v2820 = v106
	goto L498
L497:
	;
	v2820 = v101
	goto L498
L498:
	;
	if v2816 != 0 {
		goto L499
	} else {
		goto L500
	}
L499:
	;
	v2821 = v2820
	goto L501
L500:
	;
	v2821 = v106
	goto L501
L501:
	;
	F_XLogRegisterBufData(m, int32(0), v2821, l7)
	mBase = m.M
	v2823 = m.ExcPending
	if v2823 != 0 {
		goto L20
	} else {
		goto L502
	}
L502:
	;
	goto L495
L503:
	;
	v2826 = *(*int32)(unsafe.Add(mBase, uint32(v136)+24))
	v2830 = v136 + v2826&int32(_a_F__bt_insertonpg_1)
	goto L505
L504:
	;
	v2830 = v2199
	goto L505
L505:
	;
	v2832 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2830)+6)))
	F_XLogRegisterBufData(m, int32(0), v2830, (v2832&int32(_a_F__bt_insertonpg_7)+int32(7))&int32(_a_F__bt_insertonpg_19))
	mBase = m.M
	v2840 = m.ExcPending
	if v2840 != 0 {
		goto L20
	} else {
		goto L506
	}
L506:
	;
	v2842 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2724)+14)))
	v2844 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v2724)+16)))
	F_XLogRegisterBufData(m, int32(1), v2724+v2842, v2844-v2842)
	mBase = m.M
	v2847 = m.ExcPending
	if v2847 != 0 {
		goto L20
	} else {
		goto L507
	}
L507:
	;
	v2851 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48)+207)))
	if v2851 != 0 {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	v2852 = int32(48)
	goto L510
L509:
	;
	v2852 = int32(64)
	goto L510
L510:
	;
	v2853 = F_XLogInsert(m, int32(11), v2852)
	mBase = m.M
	v2854 = m.ExcPending
	if v2854 != 0 {
		goto L20
	} else {
		goto L511
	}
L511:
	;
	v2859 = v2830
	v2860 = v2853
	goto L467
L512:
	;
	v2859 = v2199
	v2860 = v2855
	goto L467
L513:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2696))) = v2862
	goto L515
L514:
	;
	goto L515
L515:
	;
	if v2150 == int32(0) {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	if l4 < int32(0) {
		goto L520
	} else {
		goto L521
	}
L517:
	;
	goto L518
L518:
	;
	v2887 = int32(_a_F__bt_insertonpg_17)
	v2889 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[8]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[8])) = v2889 - int32(1)
	if v139 != 0 {
		goto L523
	} else {
		goto L524
	}
L519:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2885))) = v2862
	goto L518
L520:
	;
	v2871 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2877 = *(*int32)(unsafe.Add(mBase, uint32(v2871+(l4^int32(-1))<<(uint(int32(2))%32))))
	v2885 = v2877
	goto L519
L521:
	;
	goto L522
L522:
	;
	v2879 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2885 = v2879 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L519
L523:
	;
	F_UnlockReleaseBuffer(m, v2695)
	mBase = m.M
	v2894 = m.ExcPending
	if v2894 != 0 {
		goto L20
	} else {
		goto L526
	}
L524:
	;
	goto L525
L525:
	;
	v2896 = int32(0)
	if v2150 == v2896 {
		goto L528
	} else {
		goto L529
	}
L526:
	;
	goto L525
L527:
	;
	if l3 < int32(0) {
		goto L534
	} else {
		goto L535
	}
L528:
	;
	F_UnlockReleaseBuffer(m, l4)
	mBase = m.M
	v2901 = m.ExcPending
	if v2901 != 0 {
		goto L20
	} else {
		goto L531
	}
L529:
	;
	goto L530
L530:
	;
	F_pfree(m, v2859)
	mBase = m.M
	v2903 = m.ExcPending
	if v2903 != 0 {
		goto L20
	} else {
		goto L532
	}
L531:
	;
	goto L527
L532:
	;
	goto L527
L533:
	;
	if v2208 < int32(0) {
		goto L538
	} else {
		goto L539
	}
L534:
	;
	v2907 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v2913 = *(*int32)(unsafe.Add(mBase, uint32(v2907+(l3^int32(-1))*int32(56))+16))
	v2922 = v2913
	goto L533
L535:
	;
	goto L536
L536:
	;
	v2915 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v2916 = int32(56)
	v2921 = *(*int32)(unsafe.Add(mBase, uint32(v2915+l3*v2916-v2916)+16))
	v2922 = v2921
	goto L533
L537:
	;
	F_PredicateLockPageSplit(m, l0, v2922, v2941)
	mBase = m.M
	v2943 = m.ExcPending
	if v2943 != 0 {
		goto L20
	} else {
		goto L541
	}
L538:
	;
	v2926 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v2932 = *(*int32)(unsafe.Add(mBase, uint32(v2926+(v2208^int32(-1))*int32(56))+16))
	v2941 = v2932
	goto L537
L539:
	;
	goto L540
L540:
	;
	v2934 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v2935 = int32(56)
	v2940 = *(*int32)(unsafe.Add(mBase, uint32(v2934+v2208*v2935-v2935)+16))
	v2941 = v2940
	goto L537
L541:
	;
	F__bt_insert_parent(m, l0, l1, l3, v2208, l5, base.B2i32(v108 != int32(0)), base.B2i32(v71|v70 == v2896))
	mBase = m.M
	v2947 = m.ExcPending
	if v2947 != 0 {
		goto L20
	} else {
		goto L542
	}
L542:
	;
	goto L23
L543:
	;
	v2984 = int32(_a_F__bt_insertonpg_17)
	v2986 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[8]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[8])) = v2986 + int32(1)
	if v10 == int32(0) {
		goto L552
	} else {
		goto L553
	}
L544:
	;
	v2953 = F__bt_getbuf(m, l0, int32(0), int32(3))
	mBase = m.M
	v2954 = m.ExcPending
	if v2954 != 0 {
		goto L20
	} else {
		goto L546
	}
L545:
	;
	v2974 = v2972 + int32(24)
	v2975 = *(*int32)(unsafe.Add(mBase, uint32(v2972)+44))
	v2976 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	if base.Ui32(v2975) < base.Ui32(v2976) {
		v2981 = v2974
		v2982 = v2953
		v2983 = v2972
		goto L543
	} else {
		goto L550
	}
L546:
	;
	if v2953 < int32(0) {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	v2958 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v2964 = *(*int32)(unsafe.Add(mBase, uint32(v2958+(v2953^int32(-1))<<(uint(int32(2))%32))))
	v2972 = v2964
	goto L545
L548:
	;
	goto L549
L549:
	;
	v2966 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v2972 = v2966 + v2953<<(uint(int32(13))%32) + int32(-8192)
	goto L545
L550:
	;
	F_UnlockReleaseBuffer(m, v2953)
	mBase = m.M
	v2979 = m.ExcPending
	if v2979 != 0 {
		goto L20
	} else {
		goto L551
	}
L551:
	;
	v2981 = v2974
	v2982 = int32(0)
	v2983 = v2972
	goto L543
L552:
	;
	v3006 = F_PageAddItemExtended(m, v67, v101, l7, v102&int32(_a_F__bt_insertonpg_4), int32(0))
	mBase = m.M
	v3007 = m.ExcPending
	if v3007 != 0 {
		goto L20
	} else {
		goto L555
	}
L553:
	;
	v2992 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v105)+6)))
	v2998 = (v2992&int32(_a_F__bt_insertonpg_7) + int32(7)) & int32(_a_F__bt_insertonpg_19)
	if v2998 == int32(0) {
		goto L552
	} else {
		goto L554
	}
L554:
	;
	base.MemoryCopy(m, v103, v105, v2998)
	goto L552
L555:
	;
	if v3006 == int32(0) {
		goto L5
	} else {
		goto L556
	}
L556:
	;
	v3011 = v72 & int32(1)
	F_MarkBufferDirty(m, l3)
	mBase = m.M
	v3013 = m.ExcPending
	if v3013 != 0 {
		goto L20
	} else {
		goto L557
	}
L557:
	;
	if v2982 != 0 {
		goto L558
	} else {
		goto L559
	}
L558:
	;
	v3014 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+4))
	if base.Ui32(v3014) <= base.Ui32(int32(2)) {
		goto L561
	} else {
		goto L562
	}
L559:
	;
	goto L560
L560:
	;
	if v3011 == int32(0) {
		goto L570
	} else {
		goto L571
	}
L561:
	;
	v3017 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v2983)+64)) = uint8(v3017)
	*(*int64)(unsafe.Add(mBase, uint32(v2983)+56)) = int64(-4616189618054758400)
	*(*int32)(unsafe.Add(mBase, uint32(v2983)+48)) = v3017
	*(*int32)(unsafe.Add(mBase, uint32(v2983)+28)) = int32(3)
	v3025 = int32(72)
	*(*uint16)(unsafe.Add(mBase, uint32(v2983)+12)) = uint16(v3025)
	goto L564
L562:
	;
	goto L563
L563:
	;
	if l3 < int32(0) {
		goto L566
	} else {
		goto L567
	}
L564:
	;
	goto L563
L565:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+16)) = v3045
	v3047 = *(*int32)(unsafe.Add(mBase, uint32(v69)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v2981)+20)) = v3047
	F_MarkBufferDirty(m, v2982)
	mBase = m.M
	v3050 = m.ExcPending
	if v3050 != 0 {
		goto L20
	} else {
		goto L569
	}
L566:
	;
	v3030 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v3036 = *(*int32)(unsafe.Add(mBase, uint32(v3030+(l3^int32(-1))*int32(56))+16))
	v3045 = v3036
	goto L565
L567:
	;
	goto L568
L568:
	;
	v3038 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v3039 = int32(56)
	v3044 = *(*int32)(unsafe.Add(mBase, uint32(v3038+l3*v3039-v3039)+16))
	v3045 = v3044
	goto L565
L569:
	;
	goto L560
L570:
	;
	if l4 < int32(0) {
		goto L574
	} else {
		goto L575
	}
L571:
	;
	goto L572
L572:
	;
	v3080 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3081 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v3080)+118)))
	if v3081 != int32(112) {
		goto L579
	} else {
		goto L580
	}
L573:
	;
	v3071 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3070)+16)))
	v3072 = v3071 + v3070
	v3073 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3072)+12)))
	v3075 = v3073 & int32(_a_F__bt_insertonpg_18)
	*(*uint16)(unsafe.Add(mBase, uint32(v3072)+12)) = uint16(v3075)
	F_MarkBufferDirty(m, l4)
	mBase = m.M
	v3078 = m.ExcPending
	if v3078 != 0 {
		goto L20
	} else {
		goto L577
	}
L574:
	;
	v3056 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v3062 = *(*int32)(unsafe.Add(mBase, uint32(v3056+(l4^int32(-1))<<(uint(int32(2))%32))))
	v3070 = v3062
	goto L573
L575:
	;
	goto L576
L576:
	;
	v3064 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v3070 = v3064 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L573
L577:
	;
	goto L572
L578:
	;
	if v2982 != 0 {
		goto L613
	} else {
		goto L614
	}
L579:
	;
	v3166 = F_XLogGetFakeLSN(m, l0)
	mBase = m.M
	v3167 = m.ExcPending
	if v3167 != 0 {
		goto L20
	} else {
		goto L612
	}
L580:
	;
	v3085 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[9]))
	if v3085 <= int32(0) {
		goto L581
	} else {
		goto L582
	}
L581:
	;
	v3088 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v3088 != 0 {
		goto L579
	} else {
		goto L584
	}
L582:
	;
	goto L583
L583:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+208)) = uint16(v102)
	F_XLogBeginInsert(m)
	mBase = m.M
	v3092 = m.ExcPending
	if v3092 != 0 {
		goto L20
	} else {
		goto L586
	}
L584:
	;
	v3089 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	if v3089 != 0 {
		goto L579
	} else {
		goto L585
	}
L585:
	;
	goto L583
L586:
	;
	F_XLogRegisterData(m, v48+int32(208), int32(2))
	mBase = m.M
	v3097 = m.ExcPending
	if v3097 != 0 {
		goto L20
	} else {
		goto L587
	}
L587:
	;
	if v3011|v10 == int32(0) {
		goto L589
	} else {
		goto L590
	}
L588:
	;
	v3158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v3156)+6)))
	F_XLogRegisterBufData(m, int32(0), v3156, v3158&int32(_a_F__bt_insertonpg_7))
	mBase = m.M
	v3162 = m.ExcPending
	if v3162 != 0 {
		goto L20
	} else {
		goto L610
	}
L589:
	;
	F_XLogRegisterBuffer(m, int32(1), l4, int32(8))
	mBase = m.M
	v3104 = m.ExcPending
	if v3104 != 0 {
		goto L20
	} else {
		goto L592
	}
L590:
	;
	goto L591
L591:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v3139 = m.ExcPending
	if v3139 != 0 {
		goto L20
	} else {
		goto L599
	}
L592:
	;
	if v2982 != 0 {
		goto L593
	} else {
		goto L594
	}
L593:
	;
	v3105 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[5]))) = v3105
	v3107 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[10]))) = v3107
	v3109 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[11]))) = v3109
	v3111 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[12]))) = v3111
	v3113 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[4]))) = v3113
	v3115 = *(*int32)(unsafe.Add(mBase, uint32(v2981)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[13]))) = v3115
	v3117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2981)+40)))
	*(*uint8)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[14]))) = uint8(v3117)
	F_XLogRegisterBuffer(m, int32(2), v2982, int32(14))
	mBase = m.M
	v3122 = m.ExcPending
	if v3122 != 0 {
		goto L20
	} else {
		goto L596
	}
L594:
	;
	v3131 = int32(16)
	goto L595
L595:
	;
	F_XLogRegisterBuffer(m, int32(0), l3, int32(8))
	mBase = m.M
	v3135 = m.ExcPending
	if v3135 != 0 {
		goto L20
	} else {
		goto L598
	}
L596:
	;
	F_XLogRegisterBufData(m, int32(2), v48+int32(_a_F__bt_insertonpg_11), int32(28))
	mBase = m.M
	v3128 = m.ExcPending
	if v3128 != 0 {
		goto L20
	} else {
		goto L597
	}
L597:
	;
	v3131 = int32(32)
	goto L595
L598:
	;
	v3155 = v3131
	v3156 = v101
	goto L588
L599:
	;
	if v10 == int32(0) {
		goto L600
	} else {
		goto L601
	}
L600:
	;
	v3143 = int32(80)
	if v3011 != 0 {
		goto L603
	} else {
		goto L604
	}
L601:
	;
	goto L602
L602:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v48)+uint32(_c_F__bt_insertonpg[5]))) = uint16(v10)
	F_XLogRegisterBufData(m, int32(0), v48+int32(_a_F__bt_insertonpg_11), int32(2))
	mBase = m.M
	v3154 = m.ExcPending
	if v3154 != 0 {
		goto L20
	} else {
		goto L609
	}
L603:
	;
	v3146 = int32(0)
	goto L605
L604:
	;
	v3146 = v3143
	goto L605
L605:
	;
	if v10 != 0 {
		goto L606
	} else {
		goto L607
	}
L606:
	;
	v3147 = v3143
	goto L608
L607:
	;
	v3147 = v3146
	goto L608
L608:
	;
	v3155 = v3147
	v3156 = v101
	goto L588
L609:
	;
	v3155 = int32(80)
	v3156 = v106
	goto L588
L610:
	;
	v3164 = F_XLogInsert(m, int32(11), v3155)
	mBase = m.M
	v3165 = m.ExcPending
	if v3165 != 0 {
		goto L20
	} else {
		goto L611
	}
L611:
	;
	v3170 = v3164
	goto L578
L612:
	;
	v3170 = v3166
	goto L578
L613:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v2983))) = base.I64_rotl(v3170, int64(32))
	goto L615
L614:
	;
	goto L615
L615:
	;
	if v3011 != 0 {
		goto L617
	} else {
		goto L618
	}
L616:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v67))) = v3197
	v3199 = int32(_a_F__bt_insertonpg_17)
	v3201 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[8]))
	*(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[8])) = v3201 - int32(1)
	if v2982 != 0 {
		goto L624
	} else {
		goto L625
	}
L617:
	;
	v3197 = base.I64_rotl(v3170, int64(32))
	goto L616
L618:
	;
	goto L619
L619:
	;
	if l4 < int32(0) {
		goto L621
	} else {
		goto L622
	}
L620:
	;
	v3195 = base.I64_rotl(v3170, int64(32))
	*(*int64)(unsafe.Add(mBase, uint32(v3193))) = v3195
	v3197 = v3195
	goto L616
L621:
	;
	v3179 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[0]))
	v3185 = *(*int32)(unsafe.Add(mBase, uint32(v3179+(l4^int32(-1))<<(uint(int32(2))%32))))
	v3193 = v3185
	goto L620
L622:
	;
	goto L623
L623:
	;
	v3187 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[1]))
	v3193 = v3187 + l4<<(uint(int32(13))%32) + int32(-8192)
	goto L620
L624:
	;
	F_UnlockReleaseBuffer(m, v2982)
	mBase = m.M
	v3206 = m.ExcPending
	if v3206 != 0 {
		goto L20
	} else {
		goto L627
	}
L625:
	;
	goto L626
L626:
	;
	if v3011 == int32(0) {
		goto L628
	} else {
		goto L629
	}
L627:
	;
	goto L626
L628:
	;
	F_UnlockReleaseBuffer(m, l4)
	mBase = m.M
	v3210 = m.ExcPending
	if v3210 != 0 {
		goto L20
	} else {
		goto L631
	}
L629:
	;
	goto L630
L630:
	;
	if v71|v108 != 0 {
		goto L633
	} else {
		goto L634
	}
L631:
	;
	F_UnlockReleaseBuffer(m, l3)
	mBase = m.M
	v3212 = m.ExcPending
	if v3212 != 0 {
		goto L20
	} else {
		goto L632
	}
L632:
	;
	goto L23
L633:
	;
	F_UnlockReleaseBuffer(m, l3)
	mBase = m.M
	v3215 = m.ExcPending
	if v3215 != 0 {
		goto L20
	} else {
		goto L636
	}
L634:
	;
	goto L635
L635:
	;
	if l3 < int32(0) {
		goto L638
	} else {
		goto L639
	}
L636:
	;
	goto L23
L637:
	;
	F_UnlockReleaseBuffer(m, l3)
	mBase = m.M
	v3236 = m.ExcPending
	if v3236 != 0 {
		goto L20
	} else {
		goto L641
	}
L638:
	;
	v3219 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v3225 = *(*int32)(unsafe.Add(mBase, uint32(v3219+(l3^int32(-1))*int32(56))+16))
	v3234 = v3225
	goto L637
L639:
	;
	goto L640
L640:
	;
	v3227 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v3228 = int32(56)
	v3233 = *(*int32)(unsafe.Add(mBase, uint32(v3227+l3*v3228-v3228)+16))
	v3234 = v3233
	goto L637
L641:
	;
	if v3234 == int32(-1) {
		goto L23
	} else {
		goto L642
	}
L642:
	;
	v3239 = F__bt_getrootheight(m, l0)
	mBase = m.M
	v3240 = m.ExcPending
	if v3240 != 0 {
		goto L20
	} else {
		goto L643
	}
L643:
	;
	if v3239 < int32(2) {
		goto L23
	} else {
		goto L644
	}
L644:
	;
	v3243 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v3243 != 0 {
		goto L645
	} else {
		goto L646
	}
L645:
	;
	v3269 = v3243
	goto L647
L646:
	;
	v3244 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3245 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+152)) = v3245
	v3247 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v48)+144)) = v3247
	v3251 = F_smgropen(m, v48+int32(144), v3244)
	mBase = m.M
	v3252 = m.ExcPending
	if v3252 != 0 {
		goto L20
	} else {
		goto L648
	}
L647:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3269)+16)) = v3234
	goto L23
L648:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v3251
	v3255 = *(*int32)(unsafe.Add(mBase, uint32(v3251)+72))
	if v3255 != 0 {
		goto L650
	} else {
		goto L651
	}
L649:
	;
	v3267 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v3269 = v3267
	goto L647
L650:
	;
	v3263 = v3255
	goto L652
L651:
	;
	v3256 = *(*int32)(unsafe.Add(mBase, uint32(v3251)+76))
	v3257 = *(*int32)(unsafe.Add(mBase, uint32(v3251)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v3256)+4)) = v3257
	v3259 = *(*int32)(unsafe.Add(mBase, uint32(v3251)+76))
	*(*int32)(unsafe.Add(mBase, uint32(v3257))) = v3259
	v3261 = *(*int32)(unsafe.Add(mBase, uint32(v3251)+72))
	v3263 = v3261
	goto L652
L652:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v3251)+72)) = v3263 + int32(1)
	goto L649
L653:
	;
	F_pfree(m, v105)
	mBase = m.M
	v3317 = m.ExcPending
	if v3317 != 0 {
		goto L20
	} else {
		goto L656
	}
L654:
	;
	goto L655
L655:
	;
	m.G0 = v48 + int32(_a_F__bt_insertonpg_0)
	return
L656:
	;
	F_pfree(m, v101)
	mBase = m.M
	v3319 = m.ExcPending
	if v3319 != 0 {
		goto L20
	} else {
		goto L657
	}
L657:
	;
	goto L655
L658:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3329 = m.ExcPending
	if v3329 != 0 {
		goto L20
	} else {
		goto L659
	}
L659:
	;
	v3330 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+2)))
	v3331 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6))))
	v3332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l6)+4)))
	if l3 < int32(0) {
		goto L661
	} else {
		goto L662
	}
L660:
	;
	v3352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+176)) = v3352 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v48)+172)) = v3351
	*(*int32)(unsafe.Add(mBase, uint32(v48)+168)) = l8
	*(*int32)(unsafe.Add(mBase, uint32(v48)+164)) = v3332
	*(*int32)(unsafe.Add(mBase, uint32(v48)+160)) = v3330 | v3331<<(uint(int32(16))%32)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_20), v48+int32(160))
	mBase = m.M
	v3367 = m.ExcPending
	if v3367 != 0 {
		goto L20
	} else {
		goto L664
	}
L661:
	;
	v3336 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v3342 = *(*int32)(unsafe.Add(mBase, uint32(v3336+(l3^int32(-1))*int32(56))+16))
	v3351 = v3342
	goto L660
L662:
	;
	goto L663
L663:
	;
	v3344 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v3345 = int32(56)
	v3350 = *(*int32)(unsafe.Add(mBase, uint32(v3344+l3*v3345-v3345)+16))
	v3351 = v3350
	goto L660
L664:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1206), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3372 = m.ExcPending
	if v3372 != 0 {
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
	v3377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48))) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+4)) = v3377 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_23), v48)
	mBase = m.M
	v3384 = m.ExcPending
	if v3384 != 0 {
		goto L20
	} else {
		goto L667
	}
L667:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1728), int32(_a_F__bt_insertonpg_24))
	mBase = m.M
	v3389 = m.ExcPending
	if v3389 != 0 {
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
	v3394 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+112)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+116)) = v3394 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_25), v48+int32(112))
	mBase = m.M
	v3403 = m.ExcPending
	if v3403 != 0 {
		goto L20
	} else {
		goto L670
	}
L670:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1799), int32(_a_F__bt_insertonpg_24))
	mBase = m.M
	v3408 = m.ExcPending
	if v3408 != 0 {
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
	v3413 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+80)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+84)) = v3413 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_26), v48+int32(80))
	mBase = m.M
	v3422 = m.ExcPending
	if v3422 != 0 {
		goto L20
	} else {
		goto L673
	}
L673:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1846), int32(_a_F__bt_insertonpg_24))
	mBase = m.M
	v3427 = m.ExcPending
	if v3427 != 0 {
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
	v3432 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+96)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+100)) = v3432 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_27), v48+int32(96))
	mBase = m.M
	v3441 = m.ExcPending
	if v3441 != 0 {
		goto L20
	} else {
		goto L676
	}
L676:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1858), int32(_a_F__bt_insertonpg_24))
	mBase = m.M
	v3446 = m.ExcPending
	if v3446 != 0 {
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
	v3451 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+48)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+52)) = v3451 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_28), v48+int32(48))
	mBase = m.M
	v3460 = m.ExcPending
	if v3460 != 0 {
		goto L20
	} else {
		goto L679
	}
L679:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1871), int32(_a_F__bt_insertonpg_24))
	mBase = m.M
	v3465 = m.ExcPending
	if v3465 != 0 {
		goto L20
	} else {
		goto L680
	}
L680:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L681:
	;
	v3470 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+64)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+68)) = v3470 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_29), v48-int32(-64))
	mBase = m.M
	v3479 = m.ExcPending
	if v3479 != 0 {
		goto L20
	} else {
		goto L682
	}
L682:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1882), int32(_a_F__bt_insertonpg_24))
	mBase = m.M
	v3484 = m.ExcPending
	if v3484 != 0 {
		goto L20
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
	v3489 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+32)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+36)) = v3489 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_27), v48+int32(32))
	mBase = m.M
	v3498 = m.ExcPending
	if v3498 != 0 {
		goto L20
	} else {
		goto L685
	}
L685:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1902), int32(_a_F__bt_insertonpg_24))
	mBase = m.M
	v3503 = m.ExcPending
	if v3503 != 0 {
		goto L20
	} else {
		goto L686
	}
L686:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L687:
	;
	F_errcode(m, int32(33557032))
	mBase = m.M
	v3510 = m.ExcPending
	if v3510 != 0 {
		goto L20
	} else {
		goto L688
	}
L688:
	;
	v3511 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v3512 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
	v3513 = *(*int32)(unsafe.Add(mBase, uint32(v2684)))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+24)) = v160
	*(*int32)(unsafe.Add(mBase, uint32(v48)+20)) = v3513
	*(*int32)(unsafe.Add(mBase, uint32(v48)+16)) = v3512
	*(*int32)(unsafe.Add(mBase, uint32(v48)+28)) = v3511 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_30), v48+int32(16))
	mBase = m.M
	v3524 = m.ExcPending
	if v3524 != 0 {
		goto L20
	} else {
		goto L689
	}
L689:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1924), int32(_a_F__bt_insertonpg_24))
	mBase = m.M
	v3529 = m.ExcPending
	if v3529 != 0 {
		goto L20
	} else {
		goto L690
	}
L690:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L691:
	;
	if l3 < int32(0) {
		goto L693
	} else {
		goto L694
	}
L692:
	;
	v3553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v48)+128)) = v3552
	*(*int32)(unsafe.Add(mBase, uint32(v48)+132)) = v3553 + int32(4)
	F_errmsg_internal(m, int32(_a_F__bt_insertonpg_31), v48+int32(128))
	mBase = m.M
	v3562 = m.ExcPending
	if v3562 != 0 {
		goto L20
	} else {
		goto L696
	}
L693:
	;
	v3537 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[2]))
	v3543 = *(*int32)(unsafe.Add(mBase, uint32(v3537+(l3^int32(-1))*int32(56))+16))
	v3552 = v3543
	goto L692
L694:
	;
	goto L695
L695:
	;
	v3545 = *(*int32)(unsafe.Add(mBase, _c_F__bt_insertonpg[3]))
	v3546 = int32(56)
	v3551 = *(*int32)(unsafe.Add(mBase, uint32(v3545+l3*v3546-v3546)+16))
	v3552 = v3551
	goto L692
L696:
	;
	F_errfinish(m, int32(_a_F__bt_insertonpg_21), int32(1304), int32(_a_F__bt_insertonpg_22))
	mBase = m.M
	v3567 = m.ExcPending
	if v3567 != 0 {
		goto L20
	} else {
		goto L697
	}
L697:
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
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v45 int64
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
	v54 = v23 + int32(20) + v29<<(uint(int32(3))%32)
	v55 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v54)+4)))
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54)+2)))
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
	var v76 int64
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
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
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
	v66 = v62 + v63*int32(56)
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
	v76 = *(*int64)(unsafe.Add(mBase, uint32(v66)+48))
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
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	if l1 != 0 {
		if l1 < int32(0) {
			v8 = *(*int32)(unsafe.Add(mBase, _c_F__bt_relandgetbuf[0]))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v8+(l1^int32(-1))*int32(56))+16))
			v23 = v14
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F__bt_relandgetbuf[1]))
			v17 = int32(56)
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v16+l1*v17-v17)+16))
			v23 = v22
		}
		if v23 == l2 {
			F_UnlockBuffer(m, l1)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v33 = l1
				if l3 == int32(0) {
					F_UnlockBuffer(m, v33)
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return int32(0)
					} else {
						F__bt_checkpage(m, l0, v33)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return int32(0)
						} else {
							return v33
						}
					}
				} else {
					F_LockBufferInternal(m, v33, l3)
					mBase = m.M
					v42 = m.ExcPending
					if v42 != 0 {
						return int32(0)
					} else {
						F__bt_checkpage(m, l0, v33)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							return v33
						}
					}
				}
			}
		} else {
			F_UnlockReleaseBuffer(m, l1)
			mBase = m.M
			v30 = m.ExcPending
			if v30 != 0 {
				return int32(0)
			} else {
				v31 = F_ReadBuffer(m, l0, l2)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int32(0)
				} else {
					v33 = v31
					if l3 == int32(0) {
						F_UnlockBuffer(m, v33)
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							F__bt_checkpage(m, l0, v33)
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int32(0)
							} else {
								return v33
							}
						}
					} else {
						F_LockBufferInternal(m, v33, l3)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int32(0)
						} else {
							F__bt_checkpage(m, l0, v33)
							mBase = m.M
							v44 = m.ExcPending
							if v44 != 0 {
								return int32(0)
							} else {
								return v33
							}
						}
					}
				}
			}
		}
	} else {
		v31 = F_ReadBuffer(m, l0, l2)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return int32(0)
		} else {
			v33 = v31
			if l3 == int32(0) {
				F_UnlockBuffer(m, v33)
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					F__bt_checkpage(m, l0, v33)
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						return v33
					}
				}
			} else {
				F_LockBufferInternal(m, v33, l3)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					F__bt_checkpage(m, l0, v33)
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return int32(0)
					} else {
						return v33
					}
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v139 int32
	_ = v139
	var __phi139 int32
	_ = __phi139
	var v140 int32
	_ = v140
	var __phi140 int32
	_ = __phi140
	var v141 int32
	_ = v141
	var __phi141 int32
	_ = __phi141
	var v144 int32
	_ = v144
	var __phi144 int32
	_ = __phi144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v184 int32
	_ = v184
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v219 int32
	_ = v219
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
	var v271 int64
	_ = v271
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v13 = F_index_form_tuple(m, v12, l2, l3)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)) = uint16(v15)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v17
	v19 = F_bt_normalize_tuple(m, l5, v13)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l5)+60))
	v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v19)+6)))
	v35 = m.G0
	v37 = v35 - int32(48)
	m.G0 = v37
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v21)+8))
	v40 = F_hash_bytes_extended(m, v19, v22&int32(_a_F_bt_tuple_present_callback_0), v39)
	mBase = m.M
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v21)+16))
	v43 = v41 - int32(1)
	v45 = v43 & base.I32_wrap_i64(v40)
	*(*int32)(unsafe.Add(mBase, uint32(v37))) = v45
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if int32(2) <= v47 {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	if v219 != 0 {
		goto L26
	} else {
		goto L27
	}
L5:
	;
	m.G0 = v37 + int32(48)
	goto L4
L6:
	;
	v184 = int32(0)
	goto L22
L7:
	;
	v51 = v47 - int32(1)
	v52 = int32(3)
	v53 = v51 & v52
	v56 = base.I32_wrap_i64(int64(base.Ui64(v40) >> (uint(int64(32)) % 64)))
	if base.Ui32(v47-int32(2)) < base.Ui32(v52) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	goto L9
L9:
	;
	if v47 != int32(1) {
		v219 = int32(0)
		goto L5
	} else {
		goto L21
	}
L10:
	;
	__phi139 = v125
	__phi140 = v126
	__phi141 = v127
	__phi144 = int32(0)
	v139 = __phi139
	v140 = __phi140
	v141 = __phi141
	v144 = __phi144
	goto L18
L11:
	;
	v125 = int32(1)
	v126 = v56
	v127 = v45
	goto L10
L12:
	;
	goto L13
L13:
	;
	v67 = int32(1)
	v68 = v56
	v69 = v45
	v72 = int32(0)
	goto L14
L14:
	;
	v78 = int32(2)
	v81 = v43 & v68
	v83 = (v69 + v81) & v43
	*(*int32)(unsafe.Add(mBase, uint32(v37+v67<<(uint(v78)%32)))) = v83
	v86 = v67 + int32(1)
	v91 = (v67 + v81) & v43
	v93 = (v91 + v83) & v43
	*(*int32)(unsafe.Add(mBase, uint32(v37+v86<<(uint(v78)%32)))) = v93
	v96 = v67 + v78
	v101 = (v91 + v86) & v43
	v103 = (v101 + v93) & v43
	*(*int32)(unsafe.Add(mBase, uint32(v37+v96<<(uint(v78)%32)))) = v103
	v106 = v67 + int32(3)
	v111 = (v101 + v96) & v43
	v113 = (v103 + v111) & v43
	*(*int32)(unsafe.Add(mBase, uint32(v37+v106<<(uint(v78)%32)))) = v113
	v115 = v111 + v106
	v116 = int32(4)
	v117 = v67 + v116
	v119 = v72 + v116
	if v119 != v51&int32(-4) {
		v67 = v117
		v68 = v115
		v69 = v113
		v72 = v119
		goto L14
	} else {
		goto L16
	}
L15:
	;
	if v53 == int32(0) {
		goto L6
	} else {
		goto L17
	}
L16:
	;
	goto L15
L17:
	;
	v125 = v117
	v126 = v115
	v127 = v113
	goto L10
L18:
	;
	v153 = v43 & v140
	v155 = (v153 + v141) & v43
	*(*int32)(unsafe.Add(mBase, uint32(v37+v139<<(uint(int32(2))%32)))) = v155
	v158 = int32(1)
	v161 = v144 + v158
	if v161 != v53 {
		__phi139 = v139 + v158
		__phi140 = v139 + v153
		__phi141 = v155
		__phi144 = v161
		v139 = __phi139
		v140 = __phi140
		v141 = __phi141
		v144 = __phi144
		goto L18
	} else {
		goto L20
	}
L19:
	;
	goto L6
L20:
	;
	goto L19
L21:
	;
	goto L6
L22:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v37+v184<<(uint(int32(2))%32))))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21+int32(24)+int32(base.Ui32(v198)>>(uint(int32(3))%32))))))
	v209 = base.B2i32(int32(base.Ui32(v202)>>(uint(v198&int32(7))%32))&int32(1) == int32(0))
	if int32(base.Ui32(v202)>>(uint(v198&int32(7))%32))&int32(1) == int32(0) {
		v219 = v209
		goto L5
	} else {
		goto L24
	}
L23:
	;
	v219 = v209
	goto L5
L24:
	;
	v213 = v184 + int32(1)
	if v213 != v47 {
		v184 = v213
		goto L22
	} else {
		goto L25
	}
L25:
	;
	goto L23
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v234 = m.ExcPending
	if v234 != 0 {
		goto L1
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v271 = *(*int64)(unsafe.Add(mBase, uint32(l5)+64))
	*(*int64)(unsafe.Add(mBase, uint32(l5)+64)) = v271 + int64(1)
	F_pfree(m, v13)
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L37
	}
L29:
	;
	F_errcode(m, int32(16779816))
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v238 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+2)))
	v239 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13))))
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l5)+4))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+48))
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v242)+48))
	v244 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v244
	v246 = int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v243 + v246
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v241 + v246
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v238 | v239<<(uint(int32(16))%32)
	F_errmsg(m, int32(_a_F_bt_tuple_present_callback_1), v10)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l5)+9)))
	if v259 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	F_errhint(m, int32(_a_F_bt_tuple_present_callback_2), int32(0))
	mBase = m.M
	v265 = m.ExcPending
	if v265 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	F_errfinish(m, int32(_a_F_bt_tuple_present_callback_3), int32(2808), int32(_a_F_bt_tuple_present_callback_4))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L36
	}
L35:
	;
	goto L34
L36:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L37:
	;
	if v13 != v19 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	F_pfree(m, v19)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	m.G0 = v10 + int32(16)
	return
L41:
	;
	goto L40
}
