package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_BlockRefTableComparator(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if base.Ui32(v7) < base.Ui32(v6) {
		return int32(1)
	} else {
		v11 = int32(-1)
		if base.Ui32(v6) < base.Ui32(v7) {
			v37 = v11
			return v37
		} else {
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
			if base.Ui32(v14) < base.Ui32(v13) {
				return int32(1)
			} else {
				if base.Ui32(v13) < base.Ui32(v14) {
					v37 = v11
					return v37
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
					if base.Ui32(v20) < base.Ui32(v19) {
						return int32(1)
					} else {
						if base.Ui32(v19) < base.Ui32(v20) {
							v37 = v11
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							v27 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
							if v27 < v26 {
								v37 = int32(1)
							} else {
								if v26 < v27 {
									v32 = int32(-1)
								} else {
									v32 = int32(0)
								}
								v37 = v32
							}
						}
						return v37
					}
				}
			}
		}
	}
}
func F_BlockRefTableReaderNextRelation(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int64
	_ = v12
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int64
	_ = v133
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = v12
	*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v12
	F_BlockRefTableRead(m, l0, v8+int32(-24), int32(24))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v26 = v8 + int32(-24)
	v28 = v8 + int32(-48)
	v29 = int32(24)
	goto L7
L3:
	;
	m.G0 = v10 - int32(-64)
	return base.B2i32(v91 != int32(0))
L4:
	;
	if v91 == int32(0) {
		goto L22
	} else {
		goto L23
	}
L5:
	;
	v91 = int32(0)
	goto L4
L6:
	;
	v65 = v60
	v66 = v61
	v67 = v62
	goto L16
L7:
	;
	if (v26|v28)&int32(3) != 0 {
		v60 = v26
		v61 = v28
		v62 = v29
		goto L6
	} else {
		goto L10
	}
L9:
	;
	if v50 == int32(0) {
		goto L5
	} else {
		goto L15
	}
L10:
	;
	v37 = v26
	v38 = v28
	v39 = v29
	goto L11
L11:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v37)))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	if v42 != v43 {
		v60 = v37
		v61 = v38
		v62 = v39
		goto L6
	} else {
		goto L13
	}
L12:
	;
	goto L9
L13:
	;
	v45 = int32(4)
	v46 = v38 + v45
	v48 = v37 + v45
	v50 = v39 - v45
	if base.Ui32(int32(3)) < base.Ui32(v50) {
		v37 = v48
		v38 = v46
		v39 = v50
		goto L11
	} else {
		goto L14
	}
L14:
	;
	goto L12
L15:
	;
	v60 = v48
	v61 = v46
	v62 = v50
	goto L6
L16:
	;
	v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v65))))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v70 == v71 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v91 = v70 - v71
	goto L4
L18:
	;
	v73 = int32(1)
	v78 = v67 - v73
	if v78 != 0 {
		v65 = v65 + v73
		v66 = v66 + v73
		v67 = v78
		goto L16
	} else {
		goto L21
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	goto L5
L22:
	;
	v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[791])))
	F_BlockRefTableRead(m, l0, v8+int32(-52), int32(4))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	goto L24
L24:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1263])))
	if v113 != 0 {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
	v102 = v94 ^ int32(-1)
	if v100 == v102 {
		goto L3
	} else {
		goto L26
	}
L26:
	;
	v104 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1264])))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1265])))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[792])))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v100
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v102
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v106
	m.T0[v105].(func(*base.Module, int32, int32, int32))(m, v104, int32(492897), v10)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	goto L3
L28:
	;
	F_pfree(m, v113)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L1
	} else {
		goto L31
	}
L29:
	;
	goto L30
L30:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	v119 = F_palloc(m, v116<<(uint(int32(1))%32))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L32
	}
L31:
	;
	goto L30
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1263]))) = v119
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	F_BlockRefTableRead(m, l0, v119, v122<<(uint(int32(1))%32))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v10)+60))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1266]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_consts[1267]))) = v127
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v10)+48))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v131
	v133 = *(*int64)(unsafe.Add(mBase, uint32(v10)+40))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v133
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v10)+52))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v135
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v10)+56))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v137
	goto L3
}
func F_RestoreBlockImage(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int64
	_ = v24
	var v28 int64
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v43 int64
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v229 int32
	_ = v229
	var v245 int32
	_ = v245
	var v250 int64
	_ = v250
	var v254 int64
	_ = v254
	var v260 int32
	_ = v260
	var v264 int64
	_ = v264
	var v270 int64
	_ = v270
	var v276 int32
	_ = v276
	var v278 int64
	_ = v278
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v294 int32
	_ = v294
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
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
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v352 int32
	_ = v352
	var v353 int32
	_ = v353
	var v370 int32
	_ = v370
	v10 = m.G0
	v12 = v10 - int32(8288)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v14)+72))
	if l1 <= v15 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	m.G0 = v12 + int32(8288)
	return v370
L2:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+29)))
	if v36 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v21 = v14 + l1*int32(52) + int32(76)
	v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v22 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l1
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)) = uint32(v24)
	v28 = int64(base.Ui64(v24) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12))) = uint32(v28)
	F_report_invalid_record(m, l0, int32(443048), v12)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	goto L5
L7:
	;
	return int32(0)
L8:
	;
	v370 = int32(0)
	goto L1
L9:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = l1
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+84)) = uint32(v39)
	v43 = int64(base.Ui64(v39) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+80)) = uint32(v43)
	F_report_invalid_record(m, l0, int32(457827), v12+int32(80))
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v21)+32))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+42)))
	if v52&int32(28) == int32(0) {
		v306 = v51
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v370 = int32(0)
	goto L1
L13:
	;
	v307 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+38)))
	if v307 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L14:
	;
	if v52&int32(4) != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+40)))
	v61 = v12 + int32(96)
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+38)))
	v73 = v61 + (int32(8192) - v63)
	v74 = v51 + v59
	if base.Ui32(v74) <= base.Ui32(v51) {
		goto L21
	} else {
		goto L22
	}
L16:
	;
	goto L17
L17:
	;
	if v52&int32(8) != 0 {
		goto L63
	} else {
		goto L64
	}
L18:
	;
	if int32(0) <= v245 {
		v306 = v12 + int32(96)
		goto L13
	} else {
		goto L61
	}
L19:
	;
	goto L18
L20:
	;
	goto L56
L21:
	;
	v217 = v51
	v218 = v61
	goto L20
L22:
	;
	if base.Ui32(v73) <= base.Ui32(v61) {
		goto L21
	} else {
		goto L23
	}
L23:
	;
	v77 = v51
	v78 = v61
	goto L24
L24:
	;
	v90 = v77 + int32(1)
	if base.Ui32(v74) <= base.Ui32(v90) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v217 = v203
	v218 = v204
	goto L20
L26:
	;
	if base.Ui32(v74) <= base.Ui32(v203) {
		v217 = v203
		v218 = v204
		goto L20
	} else {
		goto L54
	}
L27:
	;
	v203 = v90
	v204 = v78
	goto L26
L28:
	;
	goto L29
L29:
	;
	v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77))))
	v95 = v78
	v97 = v90
	v103 = v92
	v104 = int32(0)
	goto L30
L30:
	;
	if v103&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v203 = v180
	v204 = v192
	goto L26
L32:
	;
	if base.Ui32(int32(6)) < base.Ui32(v104) {
		v203 = v180
		v204 = v192
		goto L26
	} else {
		goto L51
	}
L33:
	;
	v108 = int32(-1)
	v109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	v113 = v109&int32(15) + int32(3)
	if v113 != int32(18) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97))))
	*(*uint8)(unsafe.Add(mBase, uint32(v95))) = uint8(v174)
	v176 = int32(1)
	v180 = v97 + v176
	v192 = v95 + v176
	goto L32
L36:
	;
	v123 = v113
	v124 = v97 + int32(2)
	goto L38
L37:
	;
	v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+2)))
	v123 = v118 + int32(18)
	v124 = v97 + int32(3)
	goto L38
L38:
	;
	if base.Ui32(v74) < base.Ui32(v124) {
		v245 = v108
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v97)+1)))
	v131 = v126 | v109<<(uint(int32(4))%32)&int32(3840)
	if v131 == int32(0) {
		v245 = v108
		goto L19
	} else {
		goto L40
	}
L40:
	;
	if v95-v61 < v131 {
		v245 = v108
		goto L19
	} else {
		goto L41
	}
L41:
	;
	v136 = v73 - v95
	if v123 < v136 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v138 = v123
	goto L44
L43:
	;
	v138 = v136
	goto L44
L44:
	;
	if v131 < v138 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v141 = v95
	v143 = v131
	v145 = v138
	goto L48
L46:
	;
	v160 = v95
	v162 = v131
	v164 = v138
	goto L47
L47:
	;
	v172 = F___memcpy(m, v160, v160-v162, v164)
	mBase = m.M
	v180 = v124
	v192 = v172 + v164
	goto L32
L48:
	;
	v152 = v145 - v143
	v154 = F___memcpy(m, v141, v141-v143, v143)
	mBase = m.M
	v155 = v154 + v143
	v157 = v143 << (uint(int32(1)) % 32)
	if v157 < v152 {
		v141 = v155
		v143 = v157
		v145 = v152
		goto L48
	} else {
		goto L50
	}
L49:
	;
	v160 = v155
	v162 = v157
	v164 = v152
	goto L47
L50:
	;
	goto L49
L51:
	;
	if base.Ui32(v74) <= base.Ui32(v180) {
		v203 = v180
		v204 = v192
		goto L26
	} else {
		goto L52
	}
L52:
	;
	v196 = int32(1)
	if base.Ui32(v192) < base.Ui32(v73) {
		v95 = v192
		v97 = v180
		v103 = int32(base.Ui32(v103&int32(254)) >> (uint(v196) % 32))
		v104 = v104 + v196
		goto L30
	} else {
		goto L53
	}
L53:
	;
	goto L31
L54:
	;
	if base.Ui32(v204) < base.Ui32(v73) {
		v77 = v203
		v78 = v204
		goto L24
	} else {
		goto L55
	}
L55:
	;
	goto L25
L56:
	;
	v229 = int32(-1)
	if v217 != v74 {
		v245 = v229
		goto L19
	} else {
		goto L59
	}
L58:
	;
	v245 = v218 - v61
	goto L19
L59:
	;
	if v218 != v73 {
		v245 = v229
		goto L19
	} else {
		goto L60
	}
L60:
	;
	goto L58
L61:
	;
	v250 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = l1
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+68)) = uint32(v250)
	v254 = int64(base.Ui64(v250) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+64)) = uint32(v254)
	F_report_invalid_record(m, l0, int32(458048), v12-int32(-64))
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L7
	} else {
		goto L62
	}
L62:
	;
	v370 = int32(0)
	goto L1
L63:
	;
	v264 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(533462)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+52)) = uint32(v264)
	v270 = int64(base.Ui64(v264) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+48)) = uint32(v270)
	F_report_invalid_record(m, l0, int32(457963), v12+int32(48))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L7
	} else {
		goto L66
	}
L64:
	;
	goto L65
L65:
	;
	v278 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v281 = base.I32_wrap_i64(int64(base.Ui64(v278) >> (uint(int64(32)) % 64)))
	v282 = base.I32_wrap_i64(v278)
	if v52&int32(16) != 0 {
		goto L67
	} else {
		goto L68
	}
L66:
	;
	v370 = int32(0)
	goto L1
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = int32(405064)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v281
	F_report_invalid_record(m, l0, int32(457963), v12+int32(32))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L7
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v282
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v281
	F_report_invalid_record(m, l0, int32(457889), v12+int32(16))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L7
	} else {
		goto L71
	}
L70:
	;
	v370 = int32(0)
	goto L1
L71:
	;
	v370 = int32(0)
	goto L1
L72:
	;
	v370 = int32(1)
	goto L1
L73:
	;
	goto L77
L74:
	;
	goto L75
L75:
	;
	v313 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)))
	if v313 != 0 {
		goto L81
	} else {
		goto L82
	}
L76:
	;
	goto L72
L77:
	;
	v311 = F__emscripten_memcpy_bulkmem(m, l2, v306, int32(8192))
	mBase = m.M
	goto L79
L79:
	;
	goto L76
L80:
	;
	v316 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+38)))
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)))
	v318 = v315 + v317
	if v318&int32(3) != 0 {
		v339 = v316
		goto L85
	} else {
		goto L86
	}
L81:
	;
	v314 = F__emscripten_memcpy_bulkmem(m, l2, v306, v313)
	mBase = m.M
	v315 = v314
	goto L83
L82:
	;
	v315 = l2
	goto L83
L83:
	;
	goto L80
L84:
	;
	v346 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+36)))
	v347 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+38)))
	v348 = v346 + v347
	v352 = int32(8192) - v348
	if v352 != 0 {
		goto L95
	} else {
		goto L96
	}
L85:
	;
	v343 = F__emscripten_memset_bulkmem(m, v318, base.I32_extend8_s(int32(0)), v339)
	mBase = m.M
	goto L93
L86:
	;
	if v316&int32(3) != 0 {
		v339 = v316
		goto L85
	} else {
		goto L87
	}
L87:
	;
	if base.Ui32(int32(1024)) < base.Ui32(v316) {
		v339 = v316
		goto L85
	} else {
		goto L88
	}
L88:
	;
	v325 = v316 + v318
	if base.Ui32(v325) <= base.Ui32(v318) {
		goto L84
	} else {
		goto L89
	}
L89:
	;
	v330 = v318 + int32(4)
	if base.Ui32(v330) < base.Ui32(v325) {
		goto L90
	} else {
		goto L91
	}
L90:
	;
	v332 = v325
	goto L92
L91:
	;
	v332 = v330
	goto L92
L92:
	;
	v339 = (v315^int32(-1)+v332-v317)&int32(-4) + int32(4)
	goto L85
L93:
	;
	goto L84
L94:
	;
	goto L72
L95:
	;
	v353 = F__emscripten_memcpy_bulkmem(m, v315+v348, v346+v306, v352)
	mBase = m.M
	goto L97
L96:
	;
	goto L97
L97:
	;
	goto L94
}
