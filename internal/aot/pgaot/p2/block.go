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
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int64
	_ = v35
	var v36 int64
	_ = v36
	var v39 int64
	_ = v39
	var v40 int64
	_ = v40
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
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
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v156 int32
	_ = v156
	var v158 int64
	_ = v158
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v169 int32
	_ = v169
	v11 = m.G0
	v13 = v11 - int32(112)
	m.G0 = v13
	v16 = v13 + int32(80)
	v17 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v16))) = v17
	v20 = v13 + int32(72)
	*(*int64)(unsafe.Add(mBase, uint32(v20))) = v17
	*(*int64)(unsafe.Add(mBase, uint32(v13)+64)) = v17
	F_BlockRefTableRead(m, l0, v13+int32(88), int32(24))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int64)(unsafe.Add(mBase, uint32(v13)+88))
	v33 = *(*int64)(unsafe.Add(mBase, uint32(v13)+64))
	v35 = *(*int64)(unsafe.Add(mBase, uint32(v13)+96))
	v36 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
	v39 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
	v40 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
	if v32^v33|(v35^v36)|(v39^v40) == int64(0) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(112)
	return v169
L4:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[0])))
	F_BlockRefTableRead(m, l0, v13+int32(60), int32(4))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	if base.Ui32(int32(4)) <= base.Ui32(v67) {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v13)+60))
	v53 = v45 ^ int32(-1)
	if v51 != v53 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[1])))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[2])))
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v51
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v53
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v57
	m.T0[v56].(func(*base.Module, int32, int32, int32))(m, v55, int32(_a_F_BlockRefTableReaderNextRelation_0), v13)
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v169 = int32(0)
	goto L3
L11:
	;
	goto L10
L12:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[1])))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[2])))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v67
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v72
	m.T0[v71].(func(*base.Module, int32, int32, int32))(m, v70, int32(_a_F_BlockRefTableReaderNextRelation_1), v13+int32(16))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L15
	}
L13:
	;
	goto L14
L14:
	;
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	if base.Ui32(int32(536870912)) <= base.Ui32(v81) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v169 = int32(0)
	goto L3
L16:
	;
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[1])))
	v85 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[2])))
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v86
	m.T0[v85].(func(*base.Module, int32, int32, int32))(m, v84, int32(_a_F_BlockRefTableReaderNextRelation_2), v13+int32(32))
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v95 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[4])))
	if v95 != 0 {
		goto L20
	} else {
		goto L21
	}
L19:
	;
	v169 = int32(0)
	goto L3
L20:
	;
	F_pfree(m, v95)
	mBase = m.M
	v97 = m.ExcPending
	if v97 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	v99 = v81
	goto L22
L22:
	;
	v100 = F_palloc_mul(m, int32(2), v99)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	v99 = v98
	goto L22
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[4]))) = v100
	v103 = int32(1)
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	F_BlockRefTableRead(m, l0, v100, v104<<(uint(v103)%32))
	mBase = m.M
	v108 = m.ExcPending
	if v108 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v13)+108))
	if v109 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[4])))
	v118 = int32(0)
	goto L29
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[5]))) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[6]))) = v109
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v13)+96))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v156
	v158 = *(*int64)(unsafe.Add(mBase, uint32(v13)+88))
	*(*int64)(unsafe.Add(mBase, uint32(l1))) = v158
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+100))
	*(*int32)(unsafe.Add(mBase, uint32(l2))) = v160
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v13)+104))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v162
	v169 = v103
	goto L3
L29:
	;
	v125 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110+v118<<(uint(int32(1))%32)))))
	if base.Ui32(int32(_a_F_BlockRefTableReaderNextRelation_3)) <= base.Ui32(v125) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L28
L31:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[1])))
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[2])))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[3])))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+56)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v13)+52)) = v118
	*(*int32)(unsafe.Add(mBase, uint32(v13)+48)) = v130
	m.T0[v129].(func(*base.Module, int32, int32, int32))(m, v128, int32(_a_F_BlockRefTableReaderNextRelation_4), v13+int32(48))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L1
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v141 = v118 + int32(1)
	if v141 != v109 {
		v118 = v141
		goto L29
	} else {
		goto L35
	}
L34:
	;
	v169 = int32(0)
	goto L3
L35:
	;
	goto L30
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
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
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v258 int32
	_ = v258
	var v266 int64
	_ = v266
	var v270 int64
	_ = v270
	var v276 int32
	_ = v276
	var v280 int64
	_ = v280
	var v286 int64
	_ = v286
	var v292 int32
	_ = v292
	var v294 int64
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v310 int32
	_ = v310
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v327 int32
	_ = v327
	var v329 int32
	_ = v329
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v346 int32
	_ = v346
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v389 int32
	_ = v389
	v10 = m.G0
	v12 = v10 - int32(_a_F_RestoreBlockImage_0)
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
	m.G0 = v12 + int32(_a_F_RestoreBlockImage_0)
	return v389
L2:
	;
	v35 = v19 + int32(76)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+29)))
	if v36 == int32(0) {
		goto L9
	} else {
		goto L10
	}
L3:
	;
	v19 = v14 + l1*int32(52)
	v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+76)))
	if v20 != 0 {
		goto L2
	} else {
		goto L6
	}
L4:
	;
	goto L5
L5:
	;
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = l1
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+4)) = uint32(v22)
	v26 = int64(base.Ui64(v22) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12))) = uint32(v26)
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_1), v12)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
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
	v389 = int32(0)
	goto L1
L9:
	;
	v39 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+88)) = l1
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+84)) = uint32(v39)
	v43 = int64(base.Ui64(v39) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+80)) = uint32(v43)
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_2), v12+int32(80))
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
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v35)+32))
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35)+42)))
	if v52&int32(28) == int32(0) {
		v321 = v51
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v389 = int32(0)
	goto L1
L13:
	;
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+38)))
	if v322 == int32(0) {
		goto L77
	} else {
		goto L78
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
	v59 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+40)))
	v61 = v12 + int32(96)
	v63 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+38)))
	v64 = int32(_a_F_RestoreBlockImage_3) - v63
	v66 = int32(0)
	v74 = v61 + v64
	v75 = v51 + v59
	if base.B2i32(v59 <= v66)|base.B2i32(v64 <= v66) == v66 {
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
		goto L67
	} else {
		goto L68
	}
L18:
	;
	if int32(0) <= v258 {
		v321 = v61
		goto L13
	} else {
		goto L65
	}
L19:
	;
	goto L18
L20:
	;
	goto L61
L21:
	;
	v83 = v51
	v86 = v61
	goto L24
L22:
	;
	goto L23
L23:
	;
	v232 = v51
	v235 = v61
	goto L20
L24:
	;
	v97 = v83 + int32(1)
	if base.Ui32(v75) <= base.Ui32(v97) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v232 = v217
	v235 = v220
	goto L20
L26:
	;
	if base.Ui32(v75) <= base.Ui32(v217) {
		v232 = v217
		v235 = v220
		goto L20
	} else {
		goto L59
	}
L27:
	;
	v217 = v97
	v220 = v86
	goto L26
L28:
	;
	goto L29
L29:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v83))))
	v101 = v97
	v104 = v86
	v110 = v99
	v111 = int32(0)
	goto L30
L30:
	;
	if v110&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v217 = v192
	v220 = v205
	goto L26
L32:
	;
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v111))|base.B2i32(base.Ui32(v75) <= base.Ui32(v192)) != 0 {
		v217 = v192
		v220 = v205
		goto L26
	} else {
		goto L57
	}
L33:
	;
	v116 = int32(-1)
	v118 = v101 + int32(2)
	if base.Ui32(v75) < base.Ui32(v118) {
		v258 = v116
		goto L19
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v186 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	*(*uint8)(unsafe.Add(mBase, uint32(v104))) = uint8(v186)
	v188 = int32(1)
	v192 = v101 + v188
	v205 = v104 + v188
	goto L32
L36:
	;
	v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+1)))
	v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
	v125 = v121&int32(15) + int32(3)
	if v125 != int32(18) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v134 = v125
	v135 = v118
	goto L39
L38:
	;
	if base.Ui32(v75) <= base.Ui32(v118) {
		v258 = v116
		goto L19
	} else {
		goto L40
	}
L39:
	;
	v140 = v121<<(uint(int32(4))%32)&int32(3840) | v120
	if base.B2i32(v140 == int32(0))|base.B2i32(v104-v61 < v140) != 0 {
		v258 = v116
		goto L19
	} else {
		goto L41
	}
L40:
	;
	v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101)+2)))
	v134 = v129 + int32(18)
	v135 = v101 + int32(3)
	goto L39
L41:
	;
	v146 = v74 - v104
	if v134 < v146 {
		goto L42
	} else {
		goto L43
	}
L42:
	;
	v148 = v134
	goto L44
L43:
	;
	v148 = v146
	goto L44
L44:
	;
	if v140 < v148 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v151 = v140
	v153 = v104
	v155 = v148
	goto L48
L46:
	;
	v171 = v140
	v173 = v104
	v175 = v148
	goto L47
L47:
	;
	if v175 != 0 {
		goto L54
	} else {
		goto L55
	}
L48:
	;
	if v151 != 0 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	v171 = v168
	v173 = v165
	v175 = v166
	goto L47
L50:
	;
	base.MemoryCopy(m, v153, v153-v151, v151)
	goto L52
L51:
	;
	goto L52
L52:
	;
	v165 = v151 + v153
	v166 = v155 - v151
	v168 = v151 << (uint(int32(1)) % 32)
	if v168 < v166 {
		v151 = v168
		v153 = v165
		v155 = v166
		goto L48
	} else {
		goto L53
	}
L53:
	;
	goto L49
L54:
	;
	base.MemoryCopy(m, v173, v173-v171, v175)
	goto L56
L55:
	;
	goto L56
L56:
	;
	v192 = v135
	v205 = v173 + v175
	goto L32
L57:
	;
	v210 = int32(1)
	if base.Ui32(v205) < base.Ui32(v74) {
		v101 = v192
		v104 = v205
		v110 = int32(base.Ui32(v110&int32(254)) >> (uint(v210) % 32))
		v111 = v111 + v210
		goto L30
	} else {
		goto L58
	}
L58:
	;
	goto L31
L59:
	;
	if base.Ui32(v220) < base.Ui32(v74) {
		v83 = v217
		v86 = v220
		goto L24
	} else {
		goto L60
	}
L60:
	;
	goto L25
L61:
	;
	if base.B2i32(v232 != v75)|base.B2i32(v235 != v74) != 0 {
		v258 = int32(-1)
		goto L19
	} else {
		goto L64
	}
L63:
	;
	v258 = v235 - v61
	goto L19
L64:
	;
	goto L63
L65:
	;
	v266 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = l1
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+68)) = uint32(v266)
	v270 = int64(base.Ui64(v266) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+64)) = uint32(v270)
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_4), v12-int32(-64))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L7
	} else {
		goto L66
	}
L66:
	;
	v389 = int32(0)
	goto L1
L67:
	;
	v280 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(_a_F_RestoreBlockImage_5)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+52)) = uint32(v280)
	v286 = int64(base.Ui64(v280) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+48)) = uint32(v286)
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_6), v12+int32(48))
	mBase = m.M
	v292 = m.ExcPending
	if v292 != 0 {
		goto L7
	} else {
		goto L70
	}
L68:
	;
	goto L69
L69:
	;
	v294 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v297 = base.I32_wrap_i64(int64(base.Ui64(v294) >> (uint(int64(32)) % 64)))
	v298 = base.I32_wrap_i64(v294)
	if v52&int32(16) != 0 {
		goto L71
	} else {
		goto L72
	}
L70:
	;
	v389 = int32(0)
	goto L1
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = int32(_a_F_RestoreBlockImage_7)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v297
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_6), v12+int32(32))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L7
	} else {
		goto L74
	}
L72:
	;
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v298
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v297
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_8), v12+int32(16))
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L7
	} else {
		goto L75
	}
L74:
	;
	v389 = int32(0)
	goto L1
L75:
	;
	v389 = int32(0)
	goto L1
L76:
	;
	v389 = int32(1)
	goto L1
L77:
	;
	base.MemoryCopy(m, l2, v321, int32(_a_F_RestoreBlockImage_3))
	goto L76
L78:
	;
	goto L79
L79:
	;
	v327 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)))
	if v327 != 0 {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	base.MemoryCopy(m, l2, v321, v327)
	goto L82
L81:
	;
	goto L82
L82:
	;
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)))
	v330 = l2 + v329
	v331 = int32(3)
	v333 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+38)))
	if v330&v331|v333&v331|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v333)) == int32(0) {
		goto L84
	} else {
		goto L85
	}
L83:
	;
	v366 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)))
	v367 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+38)))
	v368 = v366 + v367
	v369 = int32(_a_F_RestoreBlockImage_3) - v368
	if v369 == int32(0) {
		goto L76
	} else {
		goto L92
	}
L84:
	;
	if v333 == int32(0) {
		goto L83
	} else {
		goto L87
	}
L85:
	;
	v357 = v333
	goto L86
L86:
	;
	if v357 == int32(0) {
		goto L83
	} else {
		goto L91
	}
L87:
	;
	v346 = v333 + v330
	v348 = v330 + int32(4)
	if base.Ui32(v348) < base.Ui32(v346) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v350 = v346
	goto L90
L89:
	;
	v350 = v348
	goto L90
L90:
	;
	v357 = (l2^int32(-1)+v350-v329)&int32(-4) + int32(4)
	goto L86
L91:
	;
	base.MemoryFill(m, v330, int32(0), v357)
	goto L83
L92:
	;
	base.MemoryCopy(m, v368+l2, v366+v321, v369)
	goto L76
}
