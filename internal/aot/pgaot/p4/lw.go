package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LWLockAcquire(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	v11 = *(*int32)(unsafe.Add(mBase, _consts[1181]))
	if v11 < int32(200) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v14 = int32(4481692)
	v16 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	v17 = int32(1)
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v16 + v17
	v21 = *(*int32)(unsafe.Add(mBase, _consts[120]))
	v27 = int32(0)
	v31 = v17
	goto L4
L2:
	;
	goto L3
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L25
	} else {
		goto L53
	}
L4:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v35 = v32
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
	if v62 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L8:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v64 = base.B2i32(v35 == v63)
	if v35 == v63 {
		goto L15
	} else {
		goto L16
	}
L9:
	;
	v47 = v35 & int32(524287)
	if v47 != 0 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v52 = v35 & int32(262144)
	v55 = int32(0)
	v61 = base.B2i32(int32(base.Ui32(v52)>>(uint(int32(18))%32)) == v55) + v35
	v62 = base.B2i32(v52 == v55)
	goto L8
L12:
	;
	v48 = v35
	goto L14
L13:
	;
	v48 = v35 | int32(262144)
	goto L14
L14:
	;
	v61 = v48
	v62 = base.B2i32(v47 == int32(0))
	goto L8
L15:
	;
	v65 = v61
	goto L17
L16:
	;
	v65 = v63
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v65
	if v64 == int32(0) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v35 = v63
	goto L6
L19:
	;
	goto L20
L20:
	;
	goto L7
L21:
	;
	v166 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0))))
	*(*int32)(unsafe.Add(mBase, uint32(v166))) = v167 | int32(16777216)
	v175 = v27
	goto L50
L22:
	;
	F_LWLockQueueSelf(m, l0, l1)
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L25
	} else {
		goto L26
	}
L23:
	;
	goto L24
L24:
	;
	v125 = int32(4408700)
	v126 = *(*int32)(unsafe.Add(mBase, _consts[1181]))
	v128 = v126 << (uint(int32(3)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+uint32(_consts[1182]))) = l0
	*(*int32)(unsafe.Add(mBase, _consts[1181])) = v126 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+uint32(_consts[1183]))) = l1
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
	v75 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v78 = v75
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
	if v105 == int32(0) {
		goto L21
	} else {
		goto L42
	}
L29:
	;
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v107 = base.B2i32(v78 == v106)
	if v78 == v106 {
		goto L36
	} else {
		goto L37
	}
L30:
	;
	v90 = v78 & int32(524287)
	if v90 != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	goto L32
L32:
	;
	v95 = v78 & int32(262144)
	v98 = int32(0)
	v104 = base.B2i32(int32(base.Ui32(v95)>>(uint(int32(18))%32)) == v98) + v78
	v105 = base.B2i32(v95 == v98)
	goto L29
L33:
	;
	v91 = v78
	goto L35
L34:
	;
	v91 = v78 | int32(262144)
	goto L35
L35:
	;
	v104 = v91
	v105 = base.B2i32(v90 == int32(0))
	goto L29
L36:
	;
	v108 = v104
	goto L38
L37:
	;
	v108 = v106
	goto L38
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
	if v107 == int32(0) {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v78 = v106
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
	v115 = m.ExcPending
	if v115 != 0 {
		goto L25
	} else {
		goto L43
	}
L43:
	;
	goto L24
L44:
	;
	v145 = v27
	goto L47
L45:
	;
	goto L46
L46:
	;
	return v31
L47:
	;
	v151 = int32(1)
	if base.Ui32(v151) < base.Ui32(v145) {
		v145 = v145 - v151
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
	v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+74)))
	if v183 != 0 {
		v175 = v175 + int32(1)
		goto L50
	} else {
		goto L52
	}
L51:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v184 | int32(1073741824)
	v188 = int32(0)
	v190 = *(*int32)(unsafe.Add(mBase, _consts[148]))
	*(*int32)(unsafe.Add(mBase, uint32(v190))) = v188
	v27 = v175
	v31 = v188
	goto L4
L52:
	;
	goto L51
L53:
	;
	F_errmsg_internal(m, int32(280717), int32(0))
	mBase = m.M
	v200 = m.ExcPending
	if v200 != 0 {
		goto L25
	} else {
		goto L54
	}
L54:
	;
	F_errfinish(m, int32(494520), int32(1212), int32(362604))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
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
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v8 = *(*int32)(unsafe.Add(mBase, _consts[1181]))
	if v8 < int32(200) {
		v11 = int32(4481692)
		v13 = *(*int32)(unsafe.Add(mBase, _consts[162]))
		*(*int32)(unsafe.Add(mBase, _consts[162])) = v13 + int32(1)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v20 = v17
		for {
			if l1 == int32(0) {
				v29 = v20 & int32(524287)
				if v29 != 0 {
					v30 = v20
				} else {
					v30 = v20 | int32(262144)
				}
				v43 = v30
				v44 = base.B2i32(v29 == int32(0))
			} else {
				v34 = v20 & int32(262144)
				v37 = int32(0)
				v43 = base.B2i32(int32(base.Ui32(v34)>>(uint(int32(18))%32)) == v37) + v20
				v44 = base.B2i32(v34 == v37)
			}
			v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v46 = base.B2i32(v20 == v45)
			if v20 == v45 {
				v47 = v43
			} else {
				v47 = v45
			}
			*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v47
			if v46 == int32(0) {
				v20 = v45
				continue
			} else {
				break
			}
			break
		}
		if v44 == int32(0) {
			v53 = int32(4481692)
			v55 = *(*int32)(unsafe.Add(mBase, _consts[162]))
			*(*int32)(unsafe.Add(mBase, _consts[162])) = v55 - int32(1)
			return v44
		} else {
			v60 = int32(4408700)
			v61 = *(*int32)(unsafe.Add(mBase, _consts[1181]))
			v63 = v61 << (uint(int32(3)) % 32)
			*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[1182]))) = l0
			*(*int32)(unsafe.Add(mBase, _consts[1181])) = v61 + int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v63)+uint32(_consts[1183]))) = l1
			return v44
		}
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v80 = m.ExcPending
		if v80 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(280717), int32(0))
			mBase = m.M
			v84 = m.ExcPending
			if v84 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(494520), int32(1361), int32(362533))
				mBase = m.M
				v89 = m.ExcPending
				if v89 != 0 {
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
	v3 = *(*int32)(unsafe.Add(mBase, _consts[1181]))
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
	v7 = int32(4481692)
	v9 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v9 + int32(1)
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v6<<(uint(int32(3))%32))+uint32(_consts[1086])))
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
	v22 = int32(4481692)
	v24 = *(*int32)(unsafe.Add(mBase, _consts[162]))
	*(*int32)(unsafe.Add(mBase, _consts[162])) = v24 - int32(1)
	v29 = *(*int32)(unsafe.Add(mBase, _consts[1181]))
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
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v347 int32
	_ = v347
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v387 int32
	_ = v387
	var v392 int32
	_ = v392
	v15 = m.G0
	v17 = v15 - int32(32)
	m.G0 = v17
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if l1 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(1)
	goto L3
L2:
	;
	v22 = int32(262144)
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v19 - v22
	if l1 != 0 {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	m.G0 = v17 + int32(32)
	return
L5:
	;
	v27 = int32(-1)
	goto L7
L6:
	;
	v27 = int32(-262144)
	goto L7
L7:
	;
	if (v19+v27)&int32(-1073217537) != int32(-1073741824) {
		goto L4
	} else {
		goto L8
	}
L8:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v34 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v33 | v34
	if v33&v34 != 0 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v40 = v33
	goto L12
L10:
	;
	goto L11
L11:
	;
	v143 = int32(-1)
	v144 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v144 == v143 {
		goto L35
	} else {
		goto L36
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+28)) = int32(315652)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+24)) = int32(888)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = int32(494520)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(0)
	if v40&int32(536870912) != 0 {
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
	v103 = int32(4102492)
	v104 = *(*int32)(unsafe.Add(mBase, _consts[415]))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(v17+int32(8))+8))
	if v106 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L17:
	;
	F_perform_spin_delay(m, v17+int32(8))
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
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
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v83&int32(536870912) != 0 {
		goto L17
	} else {
		goto L21
	}
L21:
	;
	goto L18
L22:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v124 = int32(536870912)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v123 | v124
	if v123&v124 != 0 {
		v40 = v123
		goto L12
	} else {
		goto L33
	}
L23:
	;
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, _consts[415])) = v121
	goto L23
L25:
	;
	if int32(999) < v104 {
		goto L23
	} else {
		goto L28
	}
L26:
	;
	goto L27
L27:
	;
	if v104 < int32(11) {
		goto L23
	} else {
		goto L32
	}
L28:
	;
	v111 = int32(900)
	if v111 <= v104 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v114 = v111
	goto L31
L30:
	;
	v114 = v104
	goto L31
L31:
	;
	v121 = v114 + int32(100)
	goto L24
L32:
	;
	v121 = v104 - int32(1)
	goto L24
L33:
	;
	goto L13
L34:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v296 = base.B2i32(v290 == v290)
	if v290 == v290 {
		goto L65
	} else {
		goto L66
	}
L35:
	;
	v276 = v143
	v278 = int32(1073741824)
	v289 = int32(1610612735)
	goto L34
L36:
	;
	goto L37
L37:
	;
	v150 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v150)))
	v160 = v143
	v162 = v151 + v144*int32(640) + int32(76)
	v164 = v144
	v166 = v150
	v167 = int32(-1)
	v168 = v151
	v169 = int32(1)
	v170 = int32(0)
	goto L38
L38:
	;
	v174 = v164 * int32(640)
	v175 = v168 + v174
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v162)))
	if v170&int32(1) == int32(0) {
		goto L42
	} else {
		goto L43
	}
L39:
	;
	if v263 != 0 {
		goto L59
	} else {
		goto L60
	}
L40:
	;
	goto L39
L41:
	;
	if v176 == int32(-1) {
		v258 = v242
		v263 = v247
		goto L40
	} else {
		goto L58
	}
L42:
	;
	v184 = v175 + int32(76)
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)))
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v175)+80))
	if v186 == int32(-1) {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+75)))
	if v181 != 0 {
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v242 = v160
	v245 = v166
	v246 = v167
	v247 = v169
	v249 = int32(1)
	goto L41
L45:
	;
	if v185 == int32(-1) {
		goto L50
	} else {
		goto L51
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v185
	v190 = *(*int32)(unsafe.Add(mBase, uint32(v184)+4))
	v195 = v190
	goto L45
L47:
	;
	goto L48
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v168+v186*int32(640))+76)) = v185
	v195 = v186
	goto L45
L49:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v184))) = int64(0)
	v209 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v209)))
	v213 = v210 + v174 + int32(76)
	if v167 == int32(-1) {
		goto L54
	} else {
		goto L55
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v195
	goto L49
L51:
	;
	goto L52
L52:
	;
	v200 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v200)))
	*(*int32)(unsafe.Add(mBase, uint32(v201+v185*int32(640))+80)) = v195
	goto L49
L53:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213))) = int32(-1)
	v229 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v175)+74)) = uint8(v229)
	v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v175)+75)))
	v234 = base.B2i32(v231 == v229) & v169
	if v231 == int32(0) {
		v258 = v226
		v263 = v234
		goto L40
	} else {
		goto L57
	}
L54:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = int32(-1)
	v226 = v164
	goto L53
L55:
	;
	goto L56
L56:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v213)+4)) = v167
	v220 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)))
	*(*int32)(unsafe.Add(mBase, uint32(v221+v167*int32(640))+76)) = v164
	v226 = v160
	goto L53
L57:
	;
	v238 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v242 = v226
	v245 = v238
	v246 = v164
	v247 = v234
	v249 = v170 | base.B2i32(v231 != int32(2))
	goto L41
L58:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(v245)))
	v160 = v242
	v162 = v252 + v176*int32(640) + int32(76)
	v164 = v176
	v166 = v245
	v167 = v246
	v168 = v252
	v169 = v247
	v170 = v249
	goto L38
L59:
	;
	v268 = int32(1073741824)
	goto L61
L60:
	;
	v268 = int32(0)
	goto L61
L61:
	;
	v270 = int32(-1)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v271 == v270 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v274 = int32(1610612735)
	goto L64
L63:
	;
	v274 = v270
	goto L64
L64:
	;
	v276 = v258
	v278 = v268
	v289 = v274
	goto L34
L65:
	;
	v297 = v289 & (v290&int32(-1610612737) | v278)
	goto L67
L66:
	;
	v297 = v290
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v297
	if v296 == int32(0) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v303 = v290
	goto L71
L69:
	;
	goto L70
L70:
	;
	if v276 == int32(-1) {
		goto L4
	} else {
		goto L80
	}
L71:
	;
	v319 = int32(-1)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v320 == v319 {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	goto L70
L73:
	;
	v323 = int32(1610612735)
	goto L75
L74:
	;
	v323 = v319
	goto L75
L75:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v303 == v325 {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v327 = (v303&int32(-1610612737) | v278) & v323
	goto L78
L77:
	;
	v327 = v325
	goto L78
L78:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v327
	if v303 != v325 {
		v303 = v325
		goto L71
	} else {
		goto L79
	}
L79:
	;
	goto L72
L80:
	;
	v347 = v276
	goto L81
L81:
	;
	v361 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v361)))
	v365 = v362 + v347*int32(640)
	v367 = v365 + int32(76)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v367)))
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v365)+80))
	if v369 != int32(-1) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v367))) = int64(0)
	v392 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v365)+74)) = uint8(v392)
	goto L4
L83:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v362+v369*int32(640))+76)) = v368
	goto L85
L84:
	;
	goto L85
L85:
	;
	if v368 != int32(-1) {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v379 = *(*int32)(unsafe.Add(mBase, _consts[156]))
	v380 = *(*int32)(unsafe.Add(mBase, uint32(v379)))
	*(*int32)(unsafe.Add(mBase, uint32(v380+v368*int32(640))+80)) = v369
	*(*int64)(unsafe.Add(mBase, uint32(v367))) = int64(0)
	v387 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v365)+74)) = uint8(v387)
	v347 = v368
	goto L81
L87:
	;
	goto L88
L88:
	;
	goto L82
}
