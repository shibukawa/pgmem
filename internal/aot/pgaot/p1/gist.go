package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_GISTInitBuffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v6 int32
	_ = v6
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	v2 = l1
	if l0 < int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _c_F_GISTInitBuffer[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+(l0^int32(-1))<<(uint(int32(2))%32))))
		v20 = v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_GISTInitBuffer[1]))
		v20 = v14 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	v21 = int32(_a_F_GISTInitBuffer_0)
	v23 = int32(0)
	if v23|(v20&int32(3)|int32(1)) == v23 {
		v39 = v20 + v21
		v41 = v20 + int32(4)
		if base.Ui32(v41) < base.Ui32(v39) {
			v43 = v39
		} else {
			v43 = v41
		}
		v48 = (v20^int32(-1)+v43)&int32(-4) + int32(4)
		if v48 == int32(0) {
		} else {
			base.MemoryFill(m, v20, int32(0), v48)
		}
	} else {
		base.MemoryFill(m, v20, int32(0), v21)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v20)+10)) = int32(_a_F_GISTInitBuffer_1)
	v62 = int32(_a_F_GISTInitBuffer_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+18)) = uint16(v62)
	v68 = int32(_a_F_GISTInitBuffer_3)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v68)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)) = uint16(v68)
	v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v72 = v20 + v71
	v73 = int32(_a_F_GISTInitBuffer_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+14)) = uint16(v73)
	*(*uint16)(unsafe.Add(mBase, uint32(v72)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v72)+8)) = int32(-1)
	return
}
func F_gistDeCompressAtt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
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
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	v6 = int32(0)
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v14 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+10)))
	if v6 < v14 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v29 = v6
	goto L4
L2:
	;
	goto L3
L3:
	;
	return
L4:
	;
	v35 = l3 + v29<<(uint(int32(4))%32)
	v37 = v29 + int32(1)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = l4 + v29
	v40 = F_index_getattr_2(m, l2, v37, v38, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
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
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
	if v42 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v90 = int32(*(*int16)(unsafe.Add(mBase, uint32(v89)+10)))
	if v37 < v90 {
		v29 = v37
		goto L4
	} else {
		goto L16
	}
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+14)) = uint8(v85)
	goto L8
L10:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+14)) = uint8(v45)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)) = uint16(v45)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v40
	v55 = l0 + int32(2708) + v29*int32(28)
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	if v56 == v45 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v75 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)) = uint16(v75)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v75
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v75
	v85 = v75
	goto L9
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(_a_F_gistDeCompressAtt_0)+v29<<(uint(int32(2))%32))))
	v63 = F_FunctionCall1Coll(m, v55, v62, v35)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	if v35 == v63 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	*(*int32)(unsafe.Add(mBase, uint32(v35))) = v66
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v63)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+4)) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v63)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v70
	v72 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v63)+12)))
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+12)) = uint16(v72)
	v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v63)+14)))
	v85 = v74
	goto L9
L16:
	;
	goto L5
}
func F_gistGetFakeLSN(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int64
	_ = v10
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v20 int64
	_ = v20
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v48 int64
	_ = v48
	var v49 int32
	_ = v49
	var v54 int64
	_ = v54
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	v5 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4)+118)))
	switch v5 - int32(112) {
	case 0:
		v15 = F_GetXLogInsertRecPtr(m)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v20 = *(*int64)(unsafe.Add(mBase, _c_F_gistGetFakeLSN[0]))
			if base.B2i32(v20 == int64(0))|base.B2i32(v15 != v20) == int32(0) {
				v27 = m.G0
				v29 = v27 - int32(16)
				m.G0 = v29
				*(*int32)(unsafe.Add(mBase, uint32(v29)+12)) = int32(0)
				F_XLogBeginInsert(m)
				mBase = m.M
				v34 = m.ExcPending
				if v34 != 0 {
					return int64(0)
				} else {
					v36 = int32(_a_F_gistGetFakeLSN_0)
					v38 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_gistGetFakeLSN[1])))
					v39 = v38 | int32(2)
					*(*uint8)(unsafe.Add(mBase, _c_F_gistGetFakeLSN[1])) = uint8(v39)
					F_XLogRegisterData(m, v29+int32(12), int32(4))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int64(0)
					} else {
						v48 = F_XLogInsert(m, int32(14), int32(112))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int64(0)
						} else {
							m.G0 = v29 + int32(16)
							v54 = v48
							*(*int64)(unsafe.Add(mBase, _c_F_gistGetFakeLSN[0])) = v54
							return v54
						}
					}
				}
			} else {
				v54 = v15
				*(*int64)(unsafe.Add(mBase, _c_F_gistGetFakeLSN[0])) = v54
				return v54
			}
		}
	default:
		v59 = *(*int32)(unsafe.Add(mBase, _c_F_gistGetFakeLSN[2]))
		v60 = *(*int64)(unsafe.Add(mBase, uint32(v59)+240))
		*(*int64)(unsafe.Add(mBase, uint32(v59)+240)) = v60 + int64(1)
		return v60
	case 4:
		v8 = int32(_a_F_gistGetFakeLSN_1)
		v10 = *(*int64)(unsafe.Add(mBase, _c_F_gistGetFakeLSN[3]))
		*(*int64)(unsafe.Add(mBase, _c_F_gistGetFakeLSN[3])) = v10 + int64(1)
		return v10
	}
}
func F_gistMakeUnionKey(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v17 int32
	_ = v17
	var v18 int64
	_ = v18
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v7 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v7
	v17 = v10 + int32(16)
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v24
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(v12)%32))+uint32(_c_F_gistMakeUnionKey[0])))
	v43 = F_FunctionCall2Coll(m, l0+l1*int32(28)+int32(916), v38, v10+int32(12), v10+int32(8))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l4))) = v43
		m.G0 = v10 + int32(48)
		return
	}
}
func F_gistSplit(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v96 int32
	_ = v96
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v141 int32
	_ = v141
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v324 int32
	_ = v324
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v382 int32
	_ = v382
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	v12 = m.G0
	v14 = v12 - int32(368)
	m.G0 = v14
	F_check_stack_depth(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if l3 != int32(1) {
			v23 = v14 + int32(172)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			v26 = int32(0)
			v27 = base.B2i32(v25 == v26)
			if v27 == v26 {
				base.MemoryFill(m, v23, int32(1), v25)
			} else {
			}
			v33 = v14 + int32(332)
			if v27 == int32(0) {
				base.MemoryFill(m, v33, int32(1), v25)
			} else {
			}
			v38 = int32(0)
			F_gistSplitByKey(m, l0, l1, l2, l3, l4, v14+int32(12), v38)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v47 = l3<<(uint(int32(2))%32) + int32(4)
				v48 = F_palloc(m, v47)
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					v50 = F_palloc(m, v47)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
						if int32(0) < v52 {
							v61 = v38
							for {
								v66 = int32(2)
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
								v70 = int32(1)
								v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+v61<<(uint(v70)%32)))))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l2+v73<<(uint(v66)%32)-int32(4))))
								*(*int32)(unsafe.Add(mBase, uint32(v48+v61<<(uint(v66)%32)))) = v79
								v82 = v61 + v70
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
								if v82 < v83 {
									v61 = v82
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
						if int32(0) < v96 {
							v106 = int32(0)
							for {
								v111 = int32(2)
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
								v115 = int32(1)
								v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114+v106<<(uint(v115)%32)))))
								v124 = *(*int32)(unsafe.Add(mBase, uint32(l2+v118<<(uint(v111)%32)-int32(4))))
								*(*int32)(unsafe.Add(mBase, uint32(v50+v106<<(uint(v111)%32)))) = v124
								v127 = v106 + v115
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
								if v127 < v128 {
									v106 = v127
									continue
								} else {
									break
								}
								break
							}
							v133 = v128
						} else {
							v133 = v96
						}
						v141 = int32(0)
						if v141 < v133 {
							if v133 != int32(1) {
								v156 = int32(0)
								v157 = v141
								v158 = v141
								for {
									v162 = int32(2)
									v164 = v50 + v157<<(uint(v162)%32)
									v165 = *(*int32)(unsafe.Add(mBase, uint32(v164)))
									v166 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v165)+6)))
									v167 = int32(_a_F_gistSplit_0)
									v170 = *(*int32)(unsafe.Add(mBase, uint32(v164)+4))
									v171 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v170)+6)))
									v176 = v158 + v166&v167 + v171&v167 + int32(8)
									v178 = v157 + v162
									v180 = v156 + v162
									if v180 != v133&int32(2147483646) {
										v156 = v180
										v157 = v178
										v158 = v176
										continue
									} else {
										break
									}
									break
								}
								if v133&int32(1) == int32(0) {
									v204 = v176
								} else {
									v186 = v178
									v187 = v176
									v194 = *(*int32)(unsafe.Add(mBase, uint32(v50+v186<<(uint(int32(2))%32))))
									v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+6)))
									v204 = v195&int32(_a_F_gistSplit_0) + v187 + int32(4)
								}
							} else {
								v186 = v141
								v187 = v141
								v194 = *(*int32)(unsafe.Add(mBase, uint32(v50+v186<<(uint(int32(2))%32))))
								v195 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v194)+6)))
								v204 = v195&int32(_a_F_gistSplit_0) + v187 + int32(4)
							}
							v218 = base.B2i32(base.Ui32(v204) < base.Ui32(int32(_a_F_gistSplit_1)))
						} else {
							v218 = int32(1)
						}
						if v218 == int32(0) {
							v221 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
							v222 = F_gistSplit(m, l0, l1, v50, v221, l4)
							mBase = m.M
							v223 = m.ExcPending
							if v223 != 0 {
								return int32(0)
							} else {
								v245 = v222
								v246 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
								v247 = int32(0)
								if v247 < v246 {
									if v246 != int32(1) {
										v262 = int32(0)
										v263 = v247
										v264 = v247
										for {
											v268 = int32(2)
											v270 = v48 + v263<<(uint(v268)%32)
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
											v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v271)+6)))
											v273 = int32(_a_F_gistSplit_0)
											v276 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
											v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v276)+6)))
											v282 = v264 + v272&v273 + v277&v273 + int32(8)
											v284 = v263 + v268
											v286 = v262 + v268
											if v286 != v246&int32(2147483646) {
												v262 = v286
												v263 = v284
												v264 = v282
												continue
											} else {
												break
											}
											break
										}
										if v246&int32(1) == int32(0) {
											v310 = v282
										} else {
											v292 = v284
											v293 = v282
											v300 = *(*int32)(unsafe.Add(mBase, uint32(v48+v292<<(uint(int32(2))%32))))
											v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+6)))
											v310 = v301&int32(_a_F_gistSplit_0) + v293 + int32(4)
										}
									} else {
										v292 = v247
										v293 = v247
										v300 = *(*int32)(unsafe.Add(mBase, uint32(v48+v292<<(uint(int32(2))%32))))
										v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+6)))
										v310 = v301&int32(_a_F_gistSplit_0) + v293 + int32(4)
									}
									v324 = base.B2i32(base.Ui32(v310) < base.Ui32(int32(_a_F_gistSplit_1)))
								} else {
									v324 = int32(1)
								}
								if v324 == int32(0) {
									v327 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
									v328 = F_gistSplit(m, l0, l1, v48, v327, l4)
									mBase = m.M
									v329 = m.ExcPending
									if v329 != 0 {
										return int32(0)
									} else {
										v336 = v328
										for {
											v341 = *(*int32)(unsafe.Add(mBase, uint32(v336)+28))
											if v341 != 0 {
												v336 = v341
												continue
											} else {
												break
											}
											break
										}
										*(*int32)(unsafe.Add(mBase, uint32(v336)+28)) = v245
										v365 = v328
										m.G0 = v14 + int32(368)
										return v365
									}
								} else {
									v344 = F_palloc0(m, int32(32))
									mBase = m.M
									v345 = m.ExcPending
									if v345 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v344)+28)) = v245
										*(*int32)(unsafe.Add(mBase, uint32(v344)+24)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v344))) = int32(-1)
										v351 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v344)+4)) = v351
										v355 = F_gistfillitupvec(m, v48, v351, v344+int32(12))
										mBase = m.M
										v356 = m.ExcPending
										if v356 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v344)+8)) = v355
											v361 = F_gistFormTuple(m, l4, l0, v14+int32(44), v23, int32(0))
											mBase = m.M
											v362 = m.ExcPending
											if v362 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v344)+16)) = v361
												v365 = v344
												m.G0 = v14 + int32(368)
												return v365
											}
										}
									}
								}
							}
						} else {
							v225 = F_palloc0(m, int32(32))
							mBase = m.M
							v226 = m.ExcPending
							if v226 != 0 {
								return int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v225)+24)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v225))) = int32(-1)
								v231 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v231
								v235 = F_gistfillitupvec(m, v50, v231, v225+int32(12))
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = v235
									v241 = F_gistFormTuple(m, l4, l0, v14+int32(204), v33, int32(0))
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = v241
										v245 = v225
										v246 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
										v247 = int32(0)
										if v247 < v246 {
											if v246 != int32(1) {
												v262 = int32(0)
												v263 = v247
												v264 = v247
												for {
													v268 = int32(2)
													v270 = v48 + v263<<(uint(v268)%32)
													v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)))
													v272 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v271)+6)))
													v273 = int32(_a_F_gistSplit_0)
													v276 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
													v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v276)+6)))
													v282 = v264 + v272&v273 + v277&v273 + int32(8)
													v284 = v263 + v268
													v286 = v262 + v268
													if v286 != v246&int32(2147483646) {
														v262 = v286
														v263 = v284
														v264 = v282
														continue
													} else {
														break
													}
													break
												}
												if v246&int32(1) == int32(0) {
													v310 = v282
												} else {
													v292 = v284
													v293 = v282
													v300 = *(*int32)(unsafe.Add(mBase, uint32(v48+v292<<(uint(int32(2))%32))))
													v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+6)))
													v310 = v301&int32(_a_F_gistSplit_0) + v293 + int32(4)
												}
											} else {
												v292 = v247
												v293 = v247
												v300 = *(*int32)(unsafe.Add(mBase, uint32(v48+v292<<(uint(int32(2))%32))))
												v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+6)))
												v310 = v301&int32(_a_F_gistSplit_0) + v293 + int32(4)
											}
											v324 = base.B2i32(base.Ui32(v310) < base.Ui32(int32(_a_F_gistSplit_1)))
										} else {
											v324 = int32(1)
										}
										if v324 == int32(0) {
											v327 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
											v328 = F_gistSplit(m, l0, l1, v48, v327, l4)
											mBase = m.M
											v329 = m.ExcPending
											if v329 != 0 {
												return int32(0)
											} else {
												v336 = v328
												for {
													v341 = *(*int32)(unsafe.Add(mBase, uint32(v336)+28))
													if v341 != 0 {
														v336 = v341
														continue
													} else {
														break
													}
													break
												}
												*(*int32)(unsafe.Add(mBase, uint32(v336)+28)) = v245
												v365 = v328
												m.G0 = v14 + int32(368)
												return v365
											}
										} else {
											v344 = F_palloc0(m, int32(32))
											mBase = m.M
											v345 = m.ExcPending
											if v345 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v344)+28)) = v245
												*(*int32)(unsafe.Add(mBase, uint32(v344)+24)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v344))) = int32(-1)
												v351 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v344)+4)) = v351
												v355 = F_gistfillitupvec(m, v48, v351, v344+int32(12))
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v344)+8)) = v355
													v361 = F_gistFormTuple(m, l4, l0, v14+int32(44), v23, int32(0))
													mBase = m.M
													v362 = m.ExcPending
													if v362 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v344)+16)) = v361
														v365 = v344
														m.G0 = v14 + int32(368)
														return v365
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
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v382 = m.ExcPending
			if v382 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v385 = m.ExcPending
				if v385 != 0 {
					return int32(0)
				} else {
					v386 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v386)+6)))
					v388 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v388 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(_a_F_gistSplit_2)
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v387 & int32(_a_F_gistSplit_0)
					F_errmsg(m, int32(_a_F_gistSplit_3), v14)
					mBase = m.M
					v399 = m.ExcPending
					if v399 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_gistSplit_4), int32(1477), int32(_a_F_gistSplit_5))
						mBase = m.M
						v404 = m.ExcPending
						if v404 != 0 {
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
func F_gist_bbox_zorder_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 float64
	_ = v14
	var v16 float64
	_ = v16
	var v20 int64
	_ = v20
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int64
	_ = v35
	var v40 int64
	_ = v40
	var v45 int64
	_ = v45
	var v50 int64
	_ = v50
	var v55 int64
	_ = v55
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v79 int64
	_ = v79
	var v84 int64
	_ = v84
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v94 int64
	_ = v94
	var v102 float64
	_ = v102
	var v104 int64
	_ = v104
	var v107 int32
	_ = v107
	var v116 int32
	_ = v116
	var v119 int64
	_ = v119
	var v120 int32
	_ = v120
	var v129 int32
	_ = v129
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v137 int64
	_ = v137
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v147 int64
	_ = v147
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v152 int64
	_ = v152
	var v153 int64
	_ = v153
	var v156 int64
	_ = v156
	var v165 int64
	_ = v165
	var v170 int64
	_ = v170
	var v175 int64
	_ = v175
	var v180 int64
	_ = v180
	var v186 int64
	_ = v186
	var v193 int32
	_ = v193
	v12 = *(*float64)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(l0)+16))
	v14 = *(*float64)(unsafe.Add(mBase, uint32(l1)+16))
	if base.F64_ne(v13, v14) != 0 {
		v20 = int64(4294967295)
		v23 = base.I32_reinterpret_f32(base.F32_demote_f64(v13))
		if base.Ui32(v23&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			if v23 < int32(0) {
				v32 = int32(-1)
			} else {
				v32 = int32(-2147483648)
			}
			v35 = base.I64_extend_i32_u(v23 ^ v32)
		} else {
			v35 = v20
		}
		v40 = (v35<<(uint(int64(16))%64) | v35) & int64(281470681808895)
		v45 = (v40<<(uint(int64(8))%64) | v40) & int64(71777214294589695)
		v50 = (v45<<(uint(int64(4))%64) | v45) & int64(1085102592571150095)
		v55 = (v50<<(uint(int64(2))%64) | v50) & int64(3689348814741910323)
		v62 = base.I32_reinterpret_f32(base.F32_demote_f64(v12))
		if base.Ui32(v62&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			if v62 < int32(0) {
				v71 = int32(-1)
			} else {
				v71 = int32(-2147483648)
			}
			v74 = base.I64_extend_i32_u(v62 ^ v71)
		} else {
			v74 = v20
		}
		v79 = (v74<<(uint(int64(16))%64) | v74) & int64(281470681808895)
		v84 = (v79<<(uint(int64(8))%64) | v79) & int64(71777214294589695)
		v89 = (v84<<(uint(int64(4))%64) | v84) & int64(1085102592571150095)
		v90 = int64(2)
		v94 = (v89<<(uint(v90)%64) | v89) & int64(3689348814741910323)
		v102 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
		v104 = int64(4294967295)
		v107 = base.I32_reinterpret_f32(base.F32_demote_f64(v14))
		if base.Ui32(v107&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			if v107 < int32(0) {
				v116 = int32(-1)
			} else {
				v116 = int32(-2147483648)
			}
			v119 = base.I64_extend_i32_u(v107 ^ v116)
		} else {
			v119 = v104
		}
		v120 = base.I32_reinterpret_f32(base.F32_demote_f64(v102))
		if base.Ui32(v120&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			if v120 < int32(0) {
				v129 = int32(-1)
			} else {
				v129 = int32(-2147483648)
			}
			v132 = base.I64_extend_i32_u(v120 ^ v129)
		} else {
			v132 = v104
		}
		v133 = (v55<<(uint(int64(1))%64)|v55)&int64(6148914691236517205) | (v94<<(uint(v90)%64)|v94<<(uint(int64(1))%64))&int64(-6148914691236517206)
		v134 = int64(16)
		v137 = int64(281470681808895)
		v138 = (v132<<(uint(v134)%64) | v132) & v137
		v139 = int64(8)
		v142 = int64(71777214294589695)
		v143 = (v138<<(uint(v139)%64) | v138) & v142
		v144 = int64(4)
		v147 = int64(1085102592571150095)
		v148 = (v143<<(uint(v144)%64) | v143) & v147
		v149 = int64(2)
		v152 = int64(3689348814741910323)
		v153 = (v148<<(uint(v149)%64) | v148) & v152
		v156 = int64(1)
		v165 = (v119<<(uint(v134)%64) | v119) & v137
		v170 = (v165<<(uint(v139)%64) | v165) & v142
		v175 = (v170<<(uint(v144)%64) | v170) & v147
		v180 = (v175<<(uint(v149)%64) | v175) & v152
		v186 = (v153<<(uint(v149)%64)|v153<<(uint(v156)%64))&int64(-6148914691236517206) | (v180<<(uint(v156)%64)|v180)&int64(6148914691236517205)
		if base.Ui64(v186) < base.Ui64(v133) {
			return int32(1)
		} else {
			if base.Ui64(v133) < base.Ui64(v186) {
				v193 = int32(-1)
			} else {
				v193 = int32(0)
			}
			return v193
		}
	} else {
		v16 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
		if base.F64_ne(v12, v16) != 0 {
			v20 = int64(4294967295)
			v23 = base.I32_reinterpret_f32(base.F32_demote_f64(v13))
			if base.Ui32(v23&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
				if v23 < int32(0) {
					v32 = int32(-1)
				} else {
					v32 = int32(-2147483648)
				}
				v35 = base.I64_extend_i32_u(v23 ^ v32)
			} else {
				v35 = v20
			}
			v40 = (v35<<(uint(int64(16))%64) | v35) & int64(281470681808895)
			v45 = (v40<<(uint(int64(8))%64) | v40) & int64(71777214294589695)
			v50 = (v45<<(uint(int64(4))%64) | v45) & int64(1085102592571150095)
			v55 = (v50<<(uint(int64(2))%64) | v50) & int64(3689348814741910323)
			v62 = base.I32_reinterpret_f32(base.F32_demote_f64(v12))
			if base.Ui32(v62&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
				if v62 < int32(0) {
					v71 = int32(-1)
				} else {
					v71 = int32(-2147483648)
				}
				v74 = base.I64_extend_i32_u(v62 ^ v71)
			} else {
				v74 = v20
			}
			v79 = (v74<<(uint(int64(16))%64) | v74) & int64(281470681808895)
			v84 = (v79<<(uint(int64(8))%64) | v79) & int64(71777214294589695)
			v89 = (v84<<(uint(int64(4))%64) | v84) & int64(1085102592571150095)
			v90 = int64(2)
			v94 = (v89<<(uint(v90)%64) | v89) & int64(3689348814741910323)
			v102 = *(*float64)(unsafe.Add(mBase, uint32(l1)+24))
			v104 = int64(4294967295)
			v107 = base.I32_reinterpret_f32(base.F32_demote_f64(v14))
			if base.Ui32(v107&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
				if v107 < int32(0) {
					v116 = int32(-1)
				} else {
					v116 = int32(-2147483648)
				}
				v119 = base.I64_extend_i32_u(v107 ^ v116)
			} else {
				v119 = v104
			}
			v120 = base.I32_reinterpret_f32(base.F32_demote_f64(v102))
			if base.Ui32(v120&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
				if v120 < int32(0) {
					v129 = int32(-1)
				} else {
					v129 = int32(-2147483648)
				}
				v132 = base.I64_extend_i32_u(v120 ^ v129)
			} else {
				v132 = v104
			}
			v133 = (v55<<(uint(int64(1))%64)|v55)&int64(6148914691236517205) | (v94<<(uint(v90)%64)|v94<<(uint(int64(1))%64))&int64(-6148914691236517206)
			v134 = int64(16)
			v137 = int64(281470681808895)
			v138 = (v132<<(uint(v134)%64) | v132) & v137
			v139 = int64(8)
			v142 = int64(71777214294589695)
			v143 = (v138<<(uint(v139)%64) | v138) & v142
			v144 = int64(4)
			v147 = int64(1085102592571150095)
			v148 = (v143<<(uint(v144)%64) | v143) & v147
			v149 = int64(2)
			v152 = int64(3689348814741910323)
			v153 = (v148<<(uint(v149)%64) | v148) & v152
			v156 = int64(1)
			v165 = (v119<<(uint(v134)%64) | v119) & v137
			v170 = (v165<<(uint(v139)%64) | v165) & v142
			v175 = (v170<<(uint(v144)%64) | v170) & v147
			v180 = (v175<<(uint(v149)%64) | v175) & v152
			v186 = (v153<<(uint(v149)%64)|v153<<(uint(v156)%64))&int64(-6148914691236517206) | (v180<<(uint(v156)%64)|v180)&int64(6148914691236517205)
			if base.Ui64(v186) < base.Ui64(v133) {
				return int32(1)
			} else {
				if base.Ui64(v133) < base.Ui64(v186) {
					v193 = int32(-1)
				} else {
					v193 = int32(0)
				}
				return v193
			}
		} else {
			return int32(0)
		}
	}
}
func F_gist_box_distance(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 float64
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	if base.Ui32(int32(20)) <= base.Ui32(v8) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v6))) = v8
			F_errmsg_internal(m, int32(_a_F_gist_box_distance_0), v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_gist_box_distance_1), int32(1492), int32(_a_F_gist_box_distance_2))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v30 = F_computeDistance(m, int32(0), v28, v29)
		mBase = m.M
		v31 = m.ExcPending
		if v31 != 0 {
			return int32(0)
		} else {
			v32 = F_Float8GetDatum(m, v30)
			mBase = m.M
			v33 = m.ExcPending
			if v33 != 0 {
				return int32(0)
			} else {
				m.G0 = v6 + int32(16)
				return v32
			}
		}
	}
}
func F_gist_box_union(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v24 int64
	_ = v24
	var v26 int64
	_ = v26
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v37 float64
	_ = v37
	var v38 float64
	_ = v38
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v54 float64
	_ = v54
	var v66 float64
	_ = v66
	var v67 float64
	_ = v67
	var v73 float64
	_ = v73
	var v90 float64
	_ = v90
	var v102 float64
	_ = v102
	var v103 float64
	_ = v103
	var v109 float64
	_ = v109
	var v122 int32
	_ = v122
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v15 = F_palloc(m, int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v20
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v22
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v24
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
		*(*int64)(unsafe.Add(mBase, uint32(v15))) = v26
		if int32(2) <= v13 {
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
			v33 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
			v37 = v32
			v38 = v33
			v41 = int32(1)
			for {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(4)+v41<<(uint(int32(4))%32))))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v38)&int64(9223372036854775807)) {
					v66 = v38
				} else {
					v54 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
					if base.B2i32(base.F64_lt(v38, v54) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v54)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
						v66 = v38
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v15))) = v54
						v66 = v54
					}
				}
				v67 = *(*float64)(unsafe.Add(mBase, uint32(v48)+16))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v67)&int64(9223372036854775807)) {
				} else {
					v73 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
					if base.B2i32(base.F64_gt(v73, v67) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v73)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v67
					}
				}
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v37)&int64(9223372036854775807)) {
					v102 = v37
				} else {
					v90 = *(*float64)(unsafe.Add(mBase, uint32(v48)+8))
					if base.B2i32(base.F64_lt(v37, v90) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v90)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
						v102 = v37
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v90
						v102 = v90
					}
				}
				v103 = *(*float64)(unsafe.Add(mBase, uint32(v48)+24))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v103)&int64(9223372036854775807)) {
				} else {
					v109 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
					if base.B2i32(base.F64_gt(v109, v103) == int32(0))&base.B2i32(base.Ui64(base.I64_reinterpret_f64(v109)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313))) != 0 {
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v103
					}
				}
				v122 = v41 + int32(1)
				if v122 != v13 {
					v37 = v102
					v38 = v66
					v41 = v122
					continue
				} else {
					break
				}
				break
			}
		} else {
		}
		*(*int32)(unsafe.Add(mBase, uint32(v11))) = int32(32)
		return v15
	}
}
func F_gist_circle_compress(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 float64
	_ = v18
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v22 float64
	_ = v22
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v37 float64
	_ = v37
	var v39 float64
	_ = v39
	var v52 float64
	_ = v52
	var v53 float64
	_ = v53
	var v54 float64
	_ = v54
	var v56 float64
	_ = v56
	var v69 float64
	_ = v69
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v73 float64
	_ = v73
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v103 int32
	_ = v103
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
	if v8 != int32(1) {
		return v7
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v14 = F_palloc(m, int32(32))
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
			v20 = base.F64_add(v18, v19)
			v22 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.B2i32(base.F64_ne(base.F64_abs(v20), v22)|base.F64_eq(base.F64_abs(v18), v22) == int32(0))&base.F64_ne(base.F64_abs(v19), v22) != 0 {
				F_float_overflow_error(m)
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			} else {
				*(*float64)(unsafe.Add(mBase, uint32(v14))) = v20
				v35 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
				v36 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
				v37 = base.F64_sub(v35, v36)
				v39 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.B2i32(base.F64_ne(base.F64_abs(v37), v39)|base.F64_eq(base.F64_abs(v35), v39) == int32(0))&base.F64_ne(base.F64_abs(v36), v39) != 0 {
					F_float_overflow_error(m)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = v37
					v52 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
					v53 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
					v54 = base.F64_add(v52, v53)
					v56 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.B2i32(base.F64_ne(base.F64_abs(v54), v56)|base.F64_eq(base.F64_abs(v52), v56) == int32(0))&base.F64_ne(base.F64_abs(v53), v56) != 0 {
						F_float_overflow_error(m)
						mBase = m.M
						v103 = m.ExcPending
						if v103 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v54
						v69 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
						v70 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
						v71 = base.F64_sub(v69, v70)
						v73 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.B2i32(base.F64_ne(base.F64_abs(v71), v73)|base.F64_eq(base.F64_abs(v69), v73) == int32(0))&base.F64_ne(base.F64_abs(v70), v73) != 0 {
							F_float_overflow_error(m)
							mBase = m.M
							v103 = m.ExcPending
							if v103 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v14)+24)) = v71
							v87 = F_palloc(m, int32(16))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v87))) = v14
								v90 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+4)) = v90
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v87)+8)) = v92
								v94 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)))
								v95 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v87)+14)) = uint8(v95)
								*(*uint16)(unsafe.Add(mBase, uint32(v87)+12)) = uint16(v94)
								return v87
							}
						}
					}
				}
			}
		}
	}
}
func F_gist_indexsortbuild_levelstate_flush(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
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
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v251 int32
	_ = v251
	var v265 int32
	_ = v265
	var v270 int32
	_ = v270
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
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v322 int32
	_ = v322
	var v326 int32
	_ = v326
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v415 int32
	_ = v415
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = l1 + int32(12)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+16)))
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+v25)+12)))
	v29 = *(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[0]))
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v33 = v24
	goto L3
L3:
	;
	v34 = int32(_a_F_gist_indexsortbuild_levelstate_flush_0)
	v35 = *(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[1]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[1])) = v38
	v42 = F_gistextractpage(m, v33, v20+int32(12))
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v33 = v32
	goto L3
L6:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v44 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v20 + int32(16)
	return
L8:
	;
	v265 = v27 & int32(1)
	v270 = v251
	goto L41
L9:
	;
	if v44 != int32(2147483647) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v118 = F_palloc0(m, int32(32))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L4
	} else {
		goto L23
	}
L12:
	;
	v52 = int32(1)
	v55 = v42
	goto L15
L13:
	;
	v93 = v42
	goto L14
L14:
	;
	v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v106 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v107 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v109 = F_gistSplit(m, v105, v106, v93, v107, v108)
	mBase = m.M
	v110 = m.ExcPending
	if v110 != 0 {
		goto L4
	} else {
		goto L21
	}
L15:
	;
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v23+v52<<(uint(int32(2))%32))))
	v75 = F_gistextractpage(m, v72, v20+int32(8))
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v93 = v78
	goto L14
L17:
	;
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v78 = F_gistjoinvector(m, v55, v20+int32(12), v75, v77)
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_pfree(m, v75)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v82 = int32(1)
	v83 = v52 + v82
	v84 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v83 < v84+v82 {
		v52 = v83
		v55 = v78
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[1])) = v35
	v113 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v113
	if v109 == v113 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v251 = v109
	goto L8
L23:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v123 = m.G0
	v125 = v123 - int32(304)
	m.G0 = v125
	F_gistMakeUnionItVec(m, v122, v42, v121, v125+int32(32), v125)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v120)+192))
	v132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v131)+10)))
	if int32(0) < v132 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v141 = v131
	v146 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(v122)+12))
	v227 = F_index_form_tuple(m, v224, v125+int32(160), v125)
	mBase = m.M
	v228 = m.ExcPending
	if v228 != 0 {
		goto L4
	} else {
		goto L39
	}
L28:
	;
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v125+v146))))
	if v157 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L27
L30:
	;
	v204 = v146 + int32(1)
	v205 = int32(*(*int16)(unsafe.Add(mBase, uint32(v199)+10)))
	if v204 < v205 {
		v141 = v199
		v146 = v204
		goto L28
	} else {
		goto L38
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125+int32(160)+v146<<(uint(int32(2))%32)))) = int32(0)
	v199 = v141
	goto L30
L32:
	;
	goto L33
L33:
	;
	v167 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v125)+302)) = uint8(v167)
	*(*uint16)(unsafe.Add(mBase, uint32(v125)+300)) = uint16(v167)
	*(*int32)(unsafe.Add(mBase, uint32(v125)+296)) = v167
	*(*int32)(unsafe.Add(mBase, uint32(v125)+292)) = v120
	v175 = v146 << (uint(int32(2)) % 32)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v175+(v125+int32(32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v125)+288)) = v179
	v183 = v122 + int32(1812) + v146*int32(28)
	v184 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
	if v184 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v175+(v122+int32(_a_F_gist_indexsortbuild_levelstate_flush_1)))))
	v189 = F_FunctionCall1Coll(m, v183, v186, v125+int32(288))
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v193 = v141
	v194 = v179
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v125+int32(160)+v175))) = v194
	v199 = v193
	goto L30
L37:
	;
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v189)))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v120)+192))
	v193 = v192
	v194 = v191
	goto L36
L38:
	;
	goto L29
L39:
	;
	v229 = int32(_a_F_gist_indexsortbuild_levelstate_flush_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v227)+4)) = uint16(v229)
	m.G0 = v125 + int32(304)
	*(*int32)(unsafe.Add(mBase, uint32(v118)+16)) = v227
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v238 = F_gistfillitupvec(m, v42, v235, v118+int32(12))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v118)+8)) = v238
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v118)+4)) = v241
	*(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[1])) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v251 = v118
	goto L8
L41:
	;
	v284 = *(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[0]))
	if v284 != 0 {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	goto L7
L43:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v270)+8))
	v288 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v289 = F_smgr_bulk_get_buf(m, v288)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	F_PageInit(m, v289, int32(_a_F_gist_indexsortbuild_levelstate_flush_3), int32(16))
	mBase = m.M
	v294 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289)+16)))
	v295 = v289 + v294
	v296 = int32(_a_F_gist_indexsortbuild_levelstate_flush_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v295)+14)) = uint16(v296)
	*(*uint16)(unsafe.Add(mBase, uint32(v295)+12)) = uint16(v265)
	*(*int32)(unsafe.Add(mBase, uint32(v295)+8)) = int32(-1)
	goto L48
L48:
	;
	v301 = int32(0)
	v302 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v302 <= v301 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v270)+16))
	v373 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v373 != 0 {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	v307 = v287
	v310 = v301
	goto L51
L51:
	;
	v322 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+6)))
	v326 = v310 + int32(1)
	v330 = F_PageAddItemExtended(m, v289, v307, v322&int32(_a_F_gist_indexsortbuild_levelstate_flush_5), v326&int32(_a_F_gist_indexsortbuild_levelstate_flush_2), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v341 = m.ExcPending
	if v341 != 0 {
		goto L4
	} else {
		goto L58
	}
L53:
	;
	if v330 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v332 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v307)+6)))
	v336 = *(*int32)(unsafe.Add(mBase, uint32(v270)+4))
	if v326 < v336 {
		v307 = v307 + v332&int32(_a_F_gist_indexsortbuild_levelstate_flush_5)
		v310 = v326
		goto L51
	} else {
		goto L57
	}
L55:
	;
	goto L56
L56:
	;
	goto L52
L57:
	;
	goto L49
L58:
	;
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v342)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v343 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gist_indexsortbuild_levelstate_flush_6), v20)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_gist_indexsortbuild_levelstate_flush_7), int32(562), int32(_a_F_gist_indexsortbuild_levelstate_flush_8))
	mBase = m.M
	v354 = m.ExcPending
	if v354 != 0 {
		goto L4
	} else {
		goto L60
	}
L60:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L61:
	;
	v374 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v289+v374)+8)) = v373
	goto L63
L62:
	;
	goto L63
L63:
	;
	v377 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v378 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v377 + v378
	*(*int64)(unsafe.Add(mBase, uint32(v289))) = int64(4294967296)
	v383 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_smgr_bulk_write(m, v383, v377, v289, v378)
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v372))) = base.I32_rotr(v377, int32(16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v377
	v391 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v391 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v395 = F_palloc0(m, int32(28))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	v415 = v391
	goto L67
L67:
	;
	F_gist_indexsortbuild_levelstate_add(m, l0, v415, v372)
	mBase = m.M
	v418 = m.ExcPending
	if v418 != 0 {
		goto L4
	} else {
		goto L71
	}
L68:
	;
	v398 = F_palloc(m, int32(_a_F_gist_indexsortbuild_levelstate_flush_3))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v400 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+8)) = v400
	*(*int32)(unsafe.Add(mBase, uint32(v395)+12)) = v398
	F_PageInit(m, v398, int32(_a_F_gist_indexsortbuild_levelstate_flush_3), int32(16))
	mBase = m.M
	v407 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v398)+16)))
	v408 = v398 + v407
	v409 = int32(_a_F_gist_indexsortbuild_levelstate_flush_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v408)+14)) = uint16(v409)
	*(*uint16)(unsafe.Add(mBase, uint32(v408)+12)) = uint16(v400)
	*(*int32)(unsafe.Add(mBase, uint32(v408)+8)) = int32(-1)
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v395
	v415 = v395
	goto L67
L71:
	;
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v270)+28))
	if v419 != 0 {
		v270 = v419
		goto L41
	} else {
		goto L72
	}
L72:
	;
	goto L42
}
func F_gist_point_consistent(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v43 float64
	_ = v43
	var v44 float64
	_ = v44
	var v48 float64
	_ = v48
	var v49 float64
	_ = v49
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 float64
	_ = v59
	var v60 float64
	_ = v60
	var v62 int32
	_ = v62
	var v70 float64
	_ = v70
	var v71 float64
	_ = v71
	var v78 int32
	_ = v78
	var v79 float64
	_ = v79
	var v80 float64
	_ = v80
	var v86 float64
	_ = v86
	var v92 float64
	_ = v92
	var v93 float64
	_ = v93
	var v99 float64
	_ = v99
	var v108 int32
	_ = v108
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 float64
	_ = v124
	var v125 int32
	_ = v125
	var v126 float64
	_ = v126
	var v130 float64
	_ = v130
	var v131 float64
	_ = v131
	var v135 float64
	_ = v135
	var v136 float64
	_ = v136
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 float64
	_ = v212
	var v213 float64
	_ = v213
	var v218 int32
	_ = v218
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	v10 = m.G0
	v12 = v10 - int32(32)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+36)))
	if v18 == int32(30) {
		v21 = int32(11)
	} else {
		v21 = v18
	}
	if v18 == int32(29) {
		v24 = int32(10)
	} else {
		v24 = v21
	}
	v28 = base.I32_div_u_s(v24&int32(_a_F_gist_point_consistent_0), int32(20))
	switch v28 {
	case 0:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v33 = v24 - v28*int32(20)
		switch v33&int32(_a_F_gist_point_consistent_0) - int32(1) {
		case 0:
			v212 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
			v213 = *(*float64)(unsafe.Add(mBase, uint32(v30)+16))
			v218 = base.F64_gt(v212, base.F64_add(v213, float64(1e-06)))
			v224 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v224)
			v227 = v218
			m.G0 = v12 + int32(32)
			return v227
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v33 & int32(_a_F_gist_point_consistent_0)
				F_errmsg_internal(m, int32(_a_F_gist_point_consistent_1), v12+int32(16))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_gist_point_consistent_2), int32(1322), int32(_a_F_gist_point_consistent_3))
					mBase = m.M
					v121 = m.ExcPending
					if v121 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 4:
			v38 = *(*float64)(unsafe.Add(mBase, uint32(v30)))
			v39 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
			v218 = base.F64_gt(v38, base.F64_add(v39, float64(1e-06)))
			v224 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v224)
			v227 = v218
			m.G0 = v12 + int32(32)
			return v227
		case 5:
			v53 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			v54 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v53)+16)))
			v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53+v54)+12)))
			if v56&int32(1) != 0 {
				v59 = *(*float64)(unsafe.Add(mBase, uint32(v30)+16))
				v60 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
				if base.F64_ne(v59, v60) != 0 {
					v62 = int32(0)
					if base.F64_le(base.F64_abs(base.F64_sub(v59, v60)), float64(1e-06)) == v62 {
						v218 = v62
					} else {
						v70 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
						v71 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
						v218 = base.F64_eq(v70, v71) | base.F64_le(base.F64_abs(base.F64_sub(v70, v71)), float64(1e-06))
					}
				} else {
					v70 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
					v71 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
					v218 = base.F64_eq(v70, v71) | base.F64_le(base.F64_abs(base.F64_sub(v70, v71)), float64(1e-06))
				}
			} else {
				v78 = int32(0)
				v79 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
				v80 = *(*float64)(unsafe.Add(mBase, uint32(v30)))
				if base.F64_le(v79, base.F64_add(v80, float64(1e-06))) == v78 {
					v218 = v78
				} else {
					v86 = *(*float64)(unsafe.Add(mBase, uint32(v30)+16))
					if base.F64_le(v86, base.F64_add(v79, float64(1e-06))) == int32(0) {
						v218 = v78
					} else {
						v92 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
						v93 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
						if base.F64_le(v92, base.F64_add(v93, float64(1e-06))) == int32(0) {
							v218 = v78
						} else {
							v99 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
							v218 = base.F64_le(v99, base.F64_add(v92, float64(1e-06)))
						}
					}
				}
			}
			v224 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v224)
			v227 = v218
			m.G0 = v12 + int32(32)
			return v227
		case 9:
			v48 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
			v49 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
			v218 = base.F64_gt(v48, base.F64_add(v49, float64(1e-06)))
			v224 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v224)
			v227 = v218
			m.G0 = v12 + int32(32)
			return v227
		case 10:
			v43 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
			v44 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
			v218 = base.F64_gt(v43, base.F64_add(v44, float64(1e-06)))
			v224 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v224)
			v227 = v218
			m.G0 = v12 + int32(32)
			return v227
		}
	case 1:
		v122 = int32(0)
		v123 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v124 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
		v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v126 = *(*float64)(unsafe.Add(mBase, uint32(v125)+16))
		if base.F64_ge(v124, v126) == v122 {
			v218 = v122
		} else {
			v130 = *(*float64)(unsafe.Add(mBase, uint32(v123)+16))
			v131 = *(*float64)(unsafe.Add(mBase, uint32(v125)))
			if base.F64_le(v130, v131) == int32(0) {
				v218 = v122
			} else {
				v135 = *(*float64)(unsafe.Add(mBase, uint32(v123)+8))
				v136 = *(*float64)(unsafe.Add(mBase, uint32(v125)+24))
				if base.F64_ge(v135, v136) == int32(0) {
					v218 = v122
				} else {
					v140 = *(*float64)(unsafe.Add(mBase, uint32(v123)+24))
					v141 = *(*float64)(unsafe.Add(mBase, uint32(v125)+8))
					v218 = base.F64_le(v140, v141)
				}
			}
		}
		v224 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v224)
		v227 = v218
		m.G0 = v12 + int32(32)
		return v227
	case 2:
		v145 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v146 = F_pg_detoast_datum(m, v145)
		mBase = m.M
		v147 = m.ExcPending
		if v147 != 0 {
			return int32(0)
		} else {
			v150 = F_DirectFunctionCall5Coll(m, int32(108), int32(0), v15, v146, int32(3), int32(0), v14)
			mBase = m.M
			v151 = m.ExcPending
			if v151 != 0 {
				return int32(0)
			} else {
				v152 = int32(0)
				v154 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)))
				v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v155)+12)))
				if base.B2i32(v157&int32(1) == v152)|base.B2i32(v150 == v152) != 0 {
					v227 = base.B2i32(v150 != v152)
					m.G0 = v12 + int32(32)
					return v227
				} else {
					v167 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v168 = F_DirectFunctionCall2Coll(m, int32(109), int32(0), v146, v167)
					mBase = m.M
					v169 = m.ExcPending
					if v169 != 0 {
						return int32(0)
					} else {
						v218 = base.B2i32(v168 != int32(0))
						v224 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v224)
						v227 = v218
						m.G0 = v12 + int32(32)
						return v227
					}
				}
			}
		}
	case 3:
		v173 = int32(0)
		v174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v177 = F_DirectFunctionCall5Coll(m, int32(110), v173, v15, v174, int32(3), v173, v14)
		mBase = m.M
		v178 = m.ExcPending
		if v178 != 0 {
			return int32(0)
		} else {
			v179 = int32(0)
			v181 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			v182 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v181)+16)))
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v181+v182)+12)))
			if base.B2i32(v184&int32(1) == v179)|base.B2i32(v177 == v179) != 0 {
				v227 = base.B2i32(v177 != v179)
				m.G0 = v12 + int32(32)
				return v227
			} else {
				v194 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
				v195 = F_DirectFunctionCall2Coll(m, int32(111), int32(0), v174, v194)
				mBase = m.M
				v196 = m.ExcPending
				if v196 != 0 {
					return int32(0)
				} else {
					v218 = base.B2i32(v195 != int32(0))
					v224 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v224)
					v227 = v218
					m.G0 = v12 + int32(32)
					return v227
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v202 = m.ExcPending
		if v202 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v24
			F_errmsg_internal(m, int32(_a_F_gist_point_consistent_1), v12)
			mBase = m.M
			v206 = m.ExcPending
			if v206 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_gist_point_consistent_2), int32(1446), int32(_a_F_gist_point_consistent_4))
				mBase = m.M
				v211 = m.ExcPending
				if v211 != 0 {
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
