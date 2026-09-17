package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockAcquire(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v105 int32
	_ = v105
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v117 int32
	_ = v117
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[0]))
	if v10 < int32(200) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v13 = int32(_a_F_LWLockAcquire_0)
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[1]))
	v16 = int32(1)
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[1])) = v15 + v16
	v20 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[2]))
	v27 = int32(0)
	v29 = v16
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L25
	} else {
		goto L53
	}
L4:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = v30
	goto L6
L6:
	;
	if l1 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	if base.B2i32(v54 == int32(0)) == int32(0) {
		goto L22
	} else {
		goto L23
	}
L8:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v58 = base.B2i32(v33 == v57)
	if v33 == v57 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v44 = v33 & int32(_a_F_LWLockAcquire_1)
	if v44 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v47 = v33 & int32(_a_F_LWLockAcquire_2)
	v53 = base.B2i32(int32(base.Ui32(v47)>>(uint(int32(18))%32)) == int32(0)) + v33
	v54 = v47
	goto L8
L12:
	;
	v45 = v33
	goto L14
L13:
	;
	v45 = v33 | int32(_a_F_LWLockAcquire_2)
	goto L14
L14:
	;
	v53 = v45
	v54 = v44
	goto L8
L15:
	;
	v59 = v53
	goto L17
L16:
	;
	v59 = v57
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v59
	if v58 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v33 = v57
	goto L6
L19:
	;
	goto L20
L20:
	;
	goto L7
L21:
	;
	v153 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[3]))
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v154 | int32(16777216)
	v163 = v27
	goto L50
L22:
	;
	F_LWLockQueueSelf(m, l0, l1)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v114 = int32(_a_F_LWLockAcquire_3)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[0]))
	v117 = v115 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+uint32(_c_F_LWLockAcquire[4]))) = l0
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[0])) = v115 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v117)+uint32(_c_F_LWLockAcquire[5]))) = l1
	if int32(0) < v27 {
		goto L44
	} else {
		goto L45
	}
L25:
	;
	return int32(0)
L26:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v72 = v69
	goto L27
L27:
	;
	if l1 == int32(0) {
		goto L30
	} else {
		goto L31
	}
L28:
	;
	if base.B2i32(v93 == int32(0)) == int32(0) {
		goto L21
	} else {
		goto L42
	}
L29:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v97 = base.B2i32(v72 == v96)
	if v72 == v96 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v83 = v72 & int32(_a_F_LWLockAcquire_1)
	if v83 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v86 = v72 & int32(_a_F_LWLockAcquire_2)
	v92 = base.B2i32(int32(base.Ui32(v86)>>(uint(int32(18))%32)) == int32(0)) + v72
	v93 = v86
	goto L29
L33:
	;
	v84 = v72
	goto L35
L34:
	;
	v84 = v72 | int32(_a_F_LWLockAcquire_2)
	goto L35
L35:
	;
	v92 = v84
	v93 = v83
	goto L29
L36:
	;
	v98 = v92
	goto L38
L37:
	;
	v98 = v96
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v98
	if v97 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v72 = v96
	goto L27
L40:
	;
	goto L41
L41:
	;
	goto L28
L42:
	;
	F_LWLockDequeueSelf(m, l0)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	goto L24
L44:
	;
	v135 = v27
	goto L47
L45:
	;
	goto L46
L46:
	;
	return v29
L47:
	;
	v139 = int32(1)
	if base.Ui32(v139) < base.Ui32(v135) {
		v135 = v135 - v139
		goto L47
	} else {
		goto L49
	}
L48:
	;
	goto L46
L49:
	;
	goto L48
L50:
	;
	v169 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+74)))
	if v169 != 0 {
		v163 = v163 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v170 | int32(1073741824)
	v174 = int32(0)
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockAcquire[3]))
	*(*int32)(unsafe.Add(mBase, uint32(v176))) = v174
	v27 = v163
	v29 = v174
	goto L4
L52:
	;
	goto L51
L53:
	;
	F_errmsg_internal(m, int32(_a_F_LWLockAcquire_4), int32(0))
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(_a_F_LWLockAcquire_5), int32(1212), int32(_a_F_LWLockAcquire_6))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L25
	} else {
		goto L55
	}
L55:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_LWLockConditionalAcquire(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	v8 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[0]))
	if v8 < int32(200) {
		v11 = int32(_a_F_LWLockConditionalAcquire_0)
		v13 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1])) = v13 + int32(1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = v17
		for {
			if l1 == int32(0) {
				v29 = v20 & int32(_a_F_LWLockConditionalAcquire_1)
				if v29 != 0 {
					v30 = v20
				} else {
					v30 = v20 | int32(_a_F_LWLockConditionalAcquire_2)
				}
				v38 = v29
				v39 = v30
			} else {
				v32 = v20 & int32(_a_F_LWLockConditionalAcquire_2)
				v38 = v32
				v39 = base.B2i32(int32(base.Ui32(v32)>>(uint(int32(18))%32)) == int32(0)) + v20
			}
			v41 = base.B2i32(v38 == int32(0))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v43 = base.B2i32(v20 == v42)
			if v20 == v42 {
				v44 = v39
			} else {
				v44 = v42
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v44
			if v43 == int32(0) {
				v20 = v42
				continue
			} else {
				break
			}
			break
		}
		if v41 == int32(0) {
			v50 = int32(_a_F_LWLockConditionalAcquire_0)
			v52 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1]))
			*(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[1])) = v52 - int32(1)
			return v41
		} else {
			v57 = int32(_a_F_LWLockConditionalAcquire_3)
			v58 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[0]))
			v60 = v58 << (uint(int32(3)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_LWLockConditionalAcquire[2]))) = l0
			*(*int32)(unsafe.Add(mBase, _c_F_LWLockConditionalAcquire[0])) = v58 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v60)+uint32(_c_F_LWLockConditionalAcquire[3]))) = l1
			return v41
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v77 = m.ExcPending
		if v77 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_LWLockConditionalAcquire_4), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_LWLockConditionalAcquire_5), int32(1361), int32(_a_F_LWLockConditionalAcquire_6))
				mBase = m.M
				v86 = m.ExcPending
				if v86 != 0 {
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
func F_LWLockReleaseAll(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	v3 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[0]))
	if int32(0) < v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v6 = v3
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v7 = int32(_a_F_LWLockReleaseAll_0)
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1])) = v9 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(3))%32))+uint32(_c_F_LWLockReleaseAll[2])))
	v18 = F_LWLockDisownInternal(m, v17)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L3
L6:
	;
	return
L7:
	;
	F_LWLockReleaseInternal(m, v17, v18)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L6
	} else {
		goto L8
	}
L8:
	;
	v22 = int32(_a_F_LWLockReleaseAll_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1]))
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[1])) = v24 - int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseAll[0]))
	if int32(0) < v29 {
		v6 = v29
		goto L4
	} else {
		goto L9
	}
L9:
	;
	goto L5
}
func F_LWLockReleaseInternal(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v286 int32
	_ = v286
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v328 int32
	_ = v328
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v365 int32
	_ = v365
	var v370 int32
	_ = v370
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v21 = int32(1)
	goto L3
L2:
	;
	v21 = int32(_a_F_LWLockReleaseInternal_0)
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v18 - v21
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v16 + int32(32)
	return
L5:
	;
	v26 = int32(-1)
	goto L7
L6:
	;
	v26 = int32(-262144)
	goto L7
L7:
	;
	if (v18+v26)&int32(-1073217537) != int32(-1073741824) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v33 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v32 | v33
	if v32&v33 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v39 = v32
	goto L12
L10:
	;
	goto L11
L11:
	;
	v138 = int32(-1)
	v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v139 == v138 {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = int32(_a_F_LWLockReleaseInternal_1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(_a_F_LWLockReleaseInternal_2)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = int64(0)
	if v39&int32(536870912) != 0 {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L11
L14:
	;
	goto L17
L15:
	;
	goto L16
L16:
	;
	v99 = int32(_a_F_LWLockReleaseInternal_3)
	v100 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[0]))
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v16+int32(8))+8))
	if v102 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	F_perform_spin_delay(m, v16+int32(8))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	goto L16
L19:
	;
	return
L20:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v80&int32(536870912) != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v119 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v120 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v119 | v120
	if v119&v120 != 0 {
		v39 = v119
		goto L12
	} else {
		goto L33
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[0])) = v117
	goto L23
L25:
	;
	if int32(999) < v100 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v100 < int32(11) {
		goto L23
	} else {
		goto L32
	}
L28:
	;
	v107 = int32(900)
	if v107 <= v100 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v110 = v107
	goto L31
L30:
	;
	v110 = v100
	goto L31
L31:
	;
	v117 = v110 + int32(100)
	goto L24
L32:
	;
	v117 = v100 - int32(1)
	goto L24
L33:
	;
	goto L13
L34:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v279 = base.B2i32(v273 == v273)
	if v273 == v273 {
		goto L65
	} else {
		goto L66
	}
L35:
	;
	v260 = v138
	v263 = int32(1073741824)
	v272 = int32(1610612735)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v145 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v153 = v138
	v154 = v139
	v156 = v146 + v139*int32(640)
	v157 = v145
	v160 = int32(-1)
	v161 = v146
	v162 = int32(1)
	v163 = int32(0)
	goto L38
L38:
	;
	v166 = v154 * int32(640)
	v167 = v161 + v166
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v156)+76))
	if v163&int32(1) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if v248 != 0 {
		goto L59
	} else {
		goto L60
	}
L40:
	;
	goto L39
L41:
	;
	if v168 == int32(-1) {
		v243 = v230
		v248 = v235
		goto L40
	} else {
		goto L58
	}
L42:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(v167)+76))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v167)+80))
	if v176 == int32(-1) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v173 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+75)))
	if v173 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v230 = v153
	v232 = v157
	v234 = v160
	v235 = v162
	v236 = int32(1)
	goto L41
L45:
	;
	if v175 == int32(-1) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v175
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v167)+80))
	v185 = v180
	goto L45
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v161+v176*int32(640))+76)) = v175
	v185 = v176
	goto L45
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v167)+76)) = int64(0)
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)))
	v201 = v200 + v166
	if v160 == int32(-1) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v185
	goto L49
L51:
	;
	goto L52
L52:
	;
	v190 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)))
	*(*int32)(unsafe.Add(mBase, uint32(v191+v175*int32(640))+80)) = v185
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+76)) = int32(-1)
	v217 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v167)+74)) = uint8(v217)
	v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v167)+75)))
	v222 = base.B2i32(v219 == v217) & v162
	if v219 == int32(0) {
		v243 = v214
		v248 = v222
		goto L40
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+80)) = int32(-1)
	v214 = v154
	goto L53
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v201)+80)) = v160
	v208 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v208)))
	*(*int32)(unsafe.Add(mBase, uint32(v209+v160*int32(640))+76)) = v154
	v214 = v153
	goto L53
L57:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v230 = v214
	v232 = v226
	v234 = v154
	v235 = v222
	v236 = v163 | base.B2i32(v219 != int32(2))
	goto L41
L58:
	;
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v232)))
	v153 = v230
	v154 = v168
	v156 = v239 + v168*int32(640)
	v157 = v232
	v160 = v234
	v161 = v239
	v162 = v235
	v163 = v236
	goto L38
L59:
	;
	v252 = int32(1073741824)
	goto L61
L60:
	;
	v252 = int32(0)
	goto L61
L61:
	;
	v254 = int32(-1)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v255 == v254 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v258 = int32(1610612735)
	goto L64
L63:
	;
	v258 = v254
	goto L64
L64:
	;
	v260 = v243
	v263 = v252
	v272 = v258
	goto L34
L65:
	;
	v280 = v272 & (v273&int32(-1610612737) | v263)
	goto L67
L66:
	;
	v280 = v273
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v280
	if v279 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v286 = v273
	goto L71
L69:
	;
	goto L70
L70:
	;
	if v260 == int32(-1) {
		goto L4
	} else {
		goto L80
	}
L71:
	;
	v301 = int32(-1)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v302 == v301 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	v305 = int32(1610612735)
	goto L75
L74:
	;
	v305 = v301
	goto L75
L75:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v286 == v307 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v309 = (v286&int32(-1610612737) | v263) & v305
	goto L78
L77:
	;
	v309 = v307
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v309
	if v286 != v307 {
		v286 = v307
		goto L71
	} else {
		goto L79
	}
L79:
	;
	goto L72
L80:
	;
	v328 = v260
	goto L81
L81:
	;
	v341 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v341)))
	v345 = v342 + v328*int32(640)
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+76))
	v347 = *(*int32)(unsafe.Add(mBase, uint32(v345)+80))
	if v347 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v345)+76)) = int64(0)
	v370 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v345)+74)) = uint8(v370)
	goto L4
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v342+v347*int32(640))+76)) = v346
	goto L85
L84:
	;
	goto L85
L85:
	;
	if v346 != int32(-1) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v357 = *(*int32)(unsafe.Add(mBase, _c_F_LWLockReleaseInternal[1]))
	v358 = *(*int32)(unsafe.Add(mBase, uint32(v357)))
	*(*int32)(unsafe.Add(mBase, uint32(v358+v346*int32(640))+80)) = v347
	*(*int64)(unsafe.Add(mBase, uint32(v345)+76)) = int64(0)
	v365 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v345)+74)) = uint8(v365)
	v328 = v346
	goto L81
L87:
	;
	goto L88
L88:
	;
	goto L82
}
