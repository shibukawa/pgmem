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
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	v2 = l1
	if l0 < int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, _consts[0]))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v6+(l0^int32(-1))<<(uint(int32(2))%32))))
		v20 = v12
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, _consts[1]))
		v20 = v14 + l0<<(uint(int32(13))%32) + int32(-8192)
	}
	if v20&int32(3) != 0 {
	} else {
	}
	v47 = F___memset(m, v20, int32(0), int32(8192))
	mBase = m.M
	*(*int32)(unsafe.Add(mBase, uint32(v20)+10)) = int32(1572864)
	v53 = int32(8196)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+18)) = uint16(v53)
	v59 = int32(8176)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)) = uint16(v59)
	*(*uint16)(unsafe.Add(mBase, uint32(v20)+14)) = uint16(v59)
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v20)+16)))
	v63 = v20 + v62
	v64 = int32(65409)
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+14)) = uint16(v64)
	*(*uint16)(unsafe.Add(mBase, uint32(v63)+12)) = uint16(v2)
	*(*int32)(unsafe.Add(mBase, uint32(v63)+8)) = int32(-1)
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
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(8084)+v29<<(uint(int32(2))%32))))
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v45 int64
	_ = v45
	var v46 int32
	_ = v46
	var v51 int64
	_ = v51
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
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
			v20 = *(*int64)(unsafe.Add(mBase, _consts[64]))
			if v20 == int64(0) {
				v51 = v15
				*(*int64)(unsafe.Add(mBase, _consts[64])) = v51
				return v51
			} else {
				if v15 != v20 {
					v51 = v15
					*(*int64)(unsafe.Add(mBase, _consts[64])) = v51
					return v51
				} else {
					v24 = m.G0
					v26 = v24 - int32(16)
					m.G0 = v26
					*(*int32)(unsafe.Add(mBase, uint32(v26)+12)) = int32(0)
					F_XLogBeginInsert(m)
					mBase = m.M
					v31 = m.ExcPending
					if v31 != 0 {
						return int64(0)
					} else {
						v33 = int32(4457908)
						v35 = int32(*(*uint8)(unsafe.Add(mBase, _consts[65])))
						v36 = v35 | int32(2)
						*(*uint8)(unsafe.Add(mBase, _consts[65])) = uint8(v36)
						F_XLogRegisterData(m, v26+int32(12), int32(4))
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return int64(0)
						} else {
							v45 = F_XLogInsert(m, int32(14), int32(112))
							mBase = m.M
							v46 = m.ExcPending
							if v46 != 0 {
								return int64(0)
							} else {
								m.G0 = v26 + int32(16)
								v51 = v45
								*(*int64)(unsafe.Add(mBase, _consts[64])) = v51
								return v51
							}
						}
					}
				}
			}
		}
	default:
		v56 = *(*int32)(unsafe.Add(mBase, _consts[30]))
		v57 = *(*int64)(unsafe.Add(mBase, uint32(v56)+240))
		*(*int64)(unsafe.Add(mBase, uint32(v56)+240)) = v57 + int64(1)
		return v57
	case 4:
		v8 = int32(4155536)
		v10 = *(*int64)(unsafe.Add(mBase, _consts[66]))
		*(*int64)(unsafe.Add(mBase, _consts[66])) = v10 + int64(1)
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
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+32)) = v24
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(v12)%32))+uint32(_consts[63])))
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
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
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
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
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
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
			v25 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
			v26 = *(*int32)(unsafe.Add(mBase, uint32(v25)))
			v28 = F__emscripten_memset_bulkmem(m, v14+int32(172), base.I32_extend8_s(int32(1)), v26)
			mBase = m.M
			v33 = F__emscripten_memset_bulkmem(m, v14+int32(332), base.I32_extend8_s(int32(1)), v26)
			mBase = m.M
			v34 = int32(0)
			F_gistSplitByKey(m, l0, l1, l2, l3, l4, v14+int32(12), v34)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v40 = int32(4)
				v41 = l2 - v40
				v45 = l3<<(uint(int32(2))%32) + v40
				v46 = F_palloc(m, v45)
				mBase = m.M
				v47 = m.ExcPending
				if v47 != 0 {
					return int32(0)
				} else {
					v48 = F_palloc(m, v45)
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
						if int32(0) < v50 {
							v59 = v34
							for {
								v64 = int32(2)
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+12))
								v68 = int32(1)
								v71 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v67+v59<<(uint(v68)%32)))))
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v41+v71<<(uint(v64)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v46+v59<<(uint(v64)%32)))) = v75
								v78 = v59 + v68
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
								if v78 < v79 {
									v59 = v78
									continue
								} else {
									break
								}
								break
							}
						} else {
						}
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
						if int32(0) < v92 {
							v102 = int32(0)
							for {
								v107 = int32(2)
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
								v111 = int32(1)
								v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v110+v102<<(uint(v111)%32)))))
								v118 = *(*int32)(unsafe.Add(mBase, uint32(v41+v114<<(uint(v107)%32))))
								*(*int32)(unsafe.Add(mBase, uint32(v48+v102<<(uint(v107)%32)))) = v118
								v121 = v102 + v111
								v122 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
								if v121 < v122 {
									v102 = v121
									continue
								} else {
									break
								}
								break
							}
							v132 = v122
						} else {
							v132 = v92
						}
						v135 = int32(0)
						if v135 < v132 {
							v142 = int32(1)
							if v132 == v142 {
								v146 = int32(0)
								v179 = v146
								v180 = v146
							} else {
								v150 = int32(0)
								v153 = v150
								v154 = v150
								v155 = v135
								for {
									v158 = int32(2)
									v160 = v48 + v153<<(uint(v158)%32)
									v161 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
									v162 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v161)+6)))
									v163 = int32(8191)
									v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
									v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v166)+6)))
									v172 = v154 + v162&v163 + v167&v163 + int32(8)
									v174 = v153 + v158
									v176 = v155 + v158
									if v176 != v132&int32(2147483646) {
										v153 = v174
										v154 = v172
										v155 = v176
										continue
									} else {
										break
									}
									break
								}
								v179 = v174
								v180 = v172
							}
							if v132&v142 != 0 {
								v187 = *(*int32)(unsafe.Add(mBase, uint32(v48+v179<<(uint(int32(2))%32))))
								v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187)+6)))
								v194 = v188&int32(8191) + v180 + int32(4)
							} else {
								v194 = v180
							}
							v203 = base.B2i32(base.Ui32(v194) < base.Ui32(int32(8153)))
						} else {
							v203 = int32(1)
						}
						if v203 == int32(0) {
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
							v207 = F_gistSplit(m, l0, l1, v48, v206, l4)
							mBase = m.M
							v208 = m.ExcPending
							if v208 != 0 {
								return int32(0)
							} else {
								v230 = v207
								v231 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
								v232 = int32(0)
								if v232 < v231 {
									v239 = int32(1)
									if v231 == v239 {
										v243 = int32(0)
										v276 = v243
										v277 = v243
									} else {
										v247 = int32(0)
										v250 = v247
										v251 = v247
										v252 = v232
										for {
											v255 = int32(2)
											v257 = v46 + v250<<(uint(v255)%32)
											v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
											v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+6)))
											v260 = int32(8191)
											v263 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
											v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263)+6)))
											v269 = v251 + v259&v260 + v264&v260 + int32(8)
											v271 = v250 + v255
											v273 = v252 + v255
											if v273 != v231&int32(2147483646) {
												v250 = v271
												v251 = v269
												v252 = v273
												continue
											} else {
												break
											}
											break
										}
										v276 = v271
										v277 = v269
									}
									if v231&v239 != 0 {
										v284 = *(*int32)(unsafe.Add(mBase, uint32(v46+v276<<(uint(int32(2))%32))))
										v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+6)))
										v291 = v285&int32(8191) + v277 + int32(4)
									} else {
										v291 = v277
									}
									v300 = base.B2i32(base.Ui32(v291) < base.Ui32(int32(8153)))
								} else {
									v300 = int32(1)
								}
								if v300 == int32(0) {
									v303 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
									v304 = F_gistSplit(m, l0, l1, v46, v303, l4)
									mBase = m.M
									v305 = m.ExcPending
									if v305 != 0 {
										return int32(0)
									} else {
										v312 = v304
										for {
											v317 = *(*int32)(unsafe.Add(mBase, uint32(v312)+28))
											if v317 != 0 {
												v312 = v317
												continue
											} else {
												break
											}
											break
										}
										*(*int32)(unsafe.Add(mBase, uint32(v312)+28)) = v230
										v347 = v304
										m.G0 = v14 + int32(368)
										return v347
									}
								} else {
									v320 = F_palloc0(m, int32(32))
									mBase = m.M
									v321 = m.ExcPending
									if v321 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v320)+28)) = v230
										*(*int32)(unsafe.Add(mBase, uint32(v320)+24)) = int32(0)
										*(*int32)(unsafe.Add(mBase, uint32(v320))) = int32(-1)
										v327 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
										*(*int32)(unsafe.Add(mBase, uint32(v320)+4)) = v327
										v331 = F_gistfillitupvec(m, v46, v327, v320+int32(12))
										mBase = m.M
										v332 = m.ExcPending
										if v332 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v320)+8)) = v331
											v337 = F_gistFormTuple(m, l4, l0, v14+int32(44), v28, int32(0))
											mBase = m.M
											v338 = m.ExcPending
											if v338 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v320)+16)) = v337
												v347 = v320
												m.G0 = v14 + int32(368)
												return v347
											}
										}
									}
								}
							}
						} else {
							v210 = F_palloc0(m, int32(32))
							mBase = m.M
							v211 = m.ExcPending
							if v211 != 0 {
								return int32(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v210)+24)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v210))) = int32(-1)
								v216 = *(*int32)(unsafe.Add(mBase, uint32(v14)+32))
								*(*int32)(unsafe.Add(mBase, uint32(v210)+4)) = v216
								v220 = F_gistfillitupvec(m, v48, v216, v210+int32(12))
								mBase = m.M
								v221 = m.ExcPending
								if v221 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v210)+8)) = v220
									v226 = F_gistFormTuple(m, l4, l0, v14+int32(204), v33, int32(0))
									mBase = m.M
									v227 = m.ExcPending
									if v227 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v210)+16)) = v226
										v230 = v210
										v231 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
										v232 = int32(0)
										if v232 < v231 {
											v239 = int32(1)
											if v231 == v239 {
												v243 = int32(0)
												v276 = v243
												v277 = v243
											} else {
												v247 = int32(0)
												v250 = v247
												v251 = v247
												v252 = v232
												for {
													v255 = int32(2)
													v257 = v46 + v250<<(uint(v255)%32)
													v258 = *(*int32)(unsafe.Add(mBase, uint32(v257)))
													v259 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v258)+6)))
													v260 = int32(8191)
													v263 = *(*int32)(unsafe.Add(mBase, uint32(v257)+4))
													v264 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v263)+6)))
													v269 = v251 + v259&v260 + v264&v260 + int32(8)
													v271 = v250 + v255
													v273 = v252 + v255
													if v273 != v231&int32(2147483646) {
														v250 = v271
														v251 = v269
														v252 = v273
														continue
													} else {
														break
													}
													break
												}
												v276 = v271
												v277 = v269
											}
											if v231&v239 != 0 {
												v284 = *(*int32)(unsafe.Add(mBase, uint32(v46+v276<<(uint(int32(2))%32))))
												v285 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v284)+6)))
												v291 = v285&int32(8191) + v277 + int32(4)
											} else {
												v291 = v277
											}
											v300 = base.B2i32(base.Ui32(v291) < base.Ui32(int32(8153)))
										} else {
											v300 = int32(1)
										}
										if v300 == int32(0) {
											v303 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
											v304 = F_gistSplit(m, l0, l1, v46, v303, l4)
											mBase = m.M
											v305 = m.ExcPending
											if v305 != 0 {
												return int32(0)
											} else {
												v312 = v304
												for {
													v317 = *(*int32)(unsafe.Add(mBase, uint32(v312)+28))
													if v317 != 0 {
														v312 = v317
														continue
													} else {
														break
													}
													break
												}
												*(*int32)(unsafe.Add(mBase, uint32(v312)+28)) = v230
												v347 = v304
												m.G0 = v14 + int32(368)
												return v347
											}
										} else {
											v320 = F_palloc0(m, int32(32))
											mBase = m.M
											v321 = m.ExcPending
											if v321 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v320)+28)) = v230
												*(*int32)(unsafe.Add(mBase, uint32(v320)+24)) = int32(0)
												*(*int32)(unsafe.Add(mBase, uint32(v320))) = int32(-1)
												v327 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
												*(*int32)(unsafe.Add(mBase, uint32(v320)+4)) = v327
												v331 = F_gistfillitupvec(m, v46, v327, v320+int32(12))
												mBase = m.M
												v332 = m.ExcPending
												if v332 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v320)+8)) = v331
													v337 = F_gistFormTuple(m, l4, l0, v14+int32(44), v28, int32(0))
													mBase = m.M
													v338 = m.ExcPending
													if v338 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v320)+16)) = v337
														v347 = v320
														m.G0 = v14 + int32(368)
														return v347
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
			v358 = m.ExcPending
			if v358 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v361 = m.ExcPending
				if v361 != 0 {
					return int32(0)
				} else {
					v362 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
					v363 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v362)+6)))
					v364 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
					*(*int32)(unsafe.Add(mBase, uint32(v14)+8)) = v364 + int32(4)
					*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = int32(8152)
					*(*int32)(unsafe.Add(mBase, uint32(v14))) = v363 & int32(8191)
					F_errmsg(m, int32(726294), v14)
					mBase = m.M
					v375 = m.ExcPending
					if v375 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(517160), int32(1477), int32(109692))
						mBase = m.M
						v380 = m.ExcPending
						if v380 != 0 {
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
	var v21 int64
	_ = v21
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v36 int64
	_ = v36
	var v41 int64
	_ = v41
	var v46 int64
	_ = v46
	var v51 int64
	_ = v51
	var v56 int64
	_ = v56
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
		v21 = int64(4294967295)
		v24 = base.I32_reinterpret_f32(base.F32_demote_f64(v13))
		if base.Ui32(v24&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			if v24 < int32(0) {
				v33 = int32(-1)
			} else {
				v33 = int32(-2147483648)
			}
			v36 = base.I64_extend_i32_u(v24 ^ v33)
		} else {
			v36 = v21
		}
		v41 = (v36<<(uint(int64(16))%64) | v36) & int64(281470681808895)
		v46 = (v41<<(uint(int64(8))%64) | v41) & int64(71777214294589695)
		v51 = (v46<<(uint(int64(4))%64) | v46) & int64(1085102592571150095)
		v56 = (v51<<(uint(int64(2))%64) | v51) & int64(3689348814741910323)
		v62 = base.I32_reinterpret_f32(base.F32_demote_f64(v12))
		if base.Ui32(v62&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			if v62 < int32(0) {
				v71 = int32(-1)
			} else {
				v71 = int32(-2147483648)
			}
			v74 = base.I64_extend_i32_u(v62 ^ v71)
		} else {
			v74 = v21
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
		v133 = (v56<<(uint(int64(1))%64)|v56)&int64(6148914691236517205) | (v94<<(uint(v90)%64)|v94<<(uint(int64(1))%64))&int64(-6148914691236517206)
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
			v21 = int64(4294967295)
			v24 = base.I32_reinterpret_f32(base.F32_demote_f64(v13))
			if base.Ui32(v24&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
				if v24 < int32(0) {
					v33 = int32(-1)
				} else {
					v33 = int32(-2147483648)
				}
				v36 = base.I64_extend_i32_u(v24 ^ v33)
			} else {
				v36 = v21
			}
			v41 = (v36<<(uint(int64(16))%64) | v36) & int64(281470681808895)
			v46 = (v41<<(uint(int64(8))%64) | v41) & int64(71777214294589695)
			v51 = (v46<<(uint(int64(4))%64) | v46) & int64(1085102592571150095)
			v56 = (v51<<(uint(int64(2))%64) | v51) & int64(3689348814741910323)
			v62 = base.I32_reinterpret_f32(base.F32_demote_f64(v12))
			if base.Ui32(v62&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
				if v62 < int32(0) {
					v71 = int32(-1)
				} else {
					v71 = int32(-2147483648)
				}
				v74 = base.I64_extend_i32_u(v62 ^ v71)
			} else {
				v74 = v21
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
			v133 = (v56<<(uint(int64(1))%64)|v56)&int64(6148914691236517205) | (v94<<(uint(v90)%64)|v94<<(uint(int64(1))%64))&int64(-6148914691236517206)
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
			F_errmsg_internal(m, int32(506251), v6)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(526011), int32(1492), int32(438043))
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
	var v65 float64
	_ = v65
	var v66 float64
	_ = v66
	var v72 float64
	_ = v72
	var v88 float64
	_ = v88
	var v99 float64
	_ = v99
	var v100 float64
	_ = v100
	var v106 float64
	_ = v106
	var v118 int32
	_ = v118
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
		v20 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
		*(*int64)(unsafe.Add(mBase, uint32(v15))) = v20
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
		*(*int64)(unsafe.Add(mBase, uint32(v15)+24)) = v22
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
		*(*int64)(unsafe.Add(mBase, uint32(v15)+16)) = v24
		v26 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v15)+8)) = v26
		if int32(2) <= v13 {
			v32 = *(*float64)(unsafe.Add(mBase, uint32(v15)+8))
			v33 = *(*float64)(unsafe.Add(mBase, uint32(v15)))
			v37 = v32
			v38 = v33
			v41 = int32(1)
			for {
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(4)+v41<<(uint(int32(4))%32))))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v38)&int64(9223372036854775807)) {
					v65 = v38
				} else {
					v54 = *(*float64)(unsafe.Add(mBase, uint32(v48)))
					if base.F64_lt(v38, v54) == int32(0) {
						if base.Ui64(base.I64_reinterpret_f64(v54)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							v65 = v38
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v15))) = v54
							v65 = v54
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v15))) = v54
						v65 = v54
					}
				}
				v66 = *(*float64)(unsafe.Add(mBase, uint32(v48)+16))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v66)&int64(9223372036854775807)) {
				} else {
					v72 = *(*float64)(unsafe.Add(mBase, uint32(v15)+16))
					if base.F64_gt(v72, v66) == int32(0) {
						if base.Ui64(base.I64_reinterpret_f64(v72)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v66
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v66
					}
				}
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v37)&int64(9223372036854775807)) {
					v99 = v37
				} else {
					v88 = *(*float64)(unsafe.Add(mBase, uint32(v48)+8))
					if base.F64_lt(v37, v88) == int32(0) {
						if base.Ui64(base.I64_reinterpret_f64(v88)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
							v99 = v37
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v88
							v99 = v88
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v88
						v99 = v88
					}
				}
				v100 = *(*float64)(unsafe.Add(mBase, uint32(v48)+24))
				if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v100)&int64(9223372036854775807)) {
				} else {
					v106 = *(*float64)(unsafe.Add(mBase, uint32(v15)+24))
					if base.F64_gt(v106, v100) == int32(0) {
						if base.Ui64(base.I64_reinterpret_f64(v106)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)) {
						} else {
							*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v100
						}
					} else {
						*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v100
					}
				}
				v118 = v41 + int32(1)
				if v118 != v13 {
					v37 = v99
					v38 = v65
					v41 = v118
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
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v44 float64
	_ = v44
	var v45 float64
	_ = v45
	var v46 float64
	_ = v46
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v87 int32
	_ = v87
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+14)))
	if v8 != int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v7
L2:
	;
	goto L3
L3:
	;
	v12 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v14 = F_palloc(m, int32(32))
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	return int32(0)
L5:
	;
	v18 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v19 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v20 = base.F64_add(v18, v19)
	if base.F64_ne(base.F64_abs(v20), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	F_float_overflow_error(m)
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L4
	} else {
		goto L24
	}
L7:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14))) = v20
	v31 = *(*float64)(unsafe.Add(mBase, uint32(v12)))
	v32 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v33 = base.F64_sub(v31, v32)
	if base.F64_ne(base.F64_abs(v33), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if base.F64_eq(base.F64_abs(v18), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	if base.F64_ne(base.F64_abs(v19), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
L11:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14)+16)) = v33
	v44 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	v45 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v46 = base.F64_add(v44, v45)
	if base.F64_ne(base.F64_abs(v46), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	if base.F64_eq(base.F64_abs(v31), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	if base.F64_ne(base.F64_abs(v32), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	goto L11
L15:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14)+8)) = v46
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v12)+8))
	v58 = *(*float64)(unsafe.Add(mBase, uint32(v12)+16))
	v59 = base.F64_sub(v57, v58)
	if base.F64_ne(base.F64_abs(v59), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L19
	} else {
		goto L20
	}
L16:
	;
	if base.F64_eq(base.F64_abs(v44), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L15
	} else {
		goto L17
	}
L17:
	;
	if base.F64_ne(base.F64_abs(v45), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L18
	}
L18:
	;
	goto L15
L19:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v14)+24)) = v59
	v71 = F_palloc(m, int32(16))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L4
	} else {
		goto L23
	}
L20:
	;
	if base.F64_eq(base.F64_abs(v57), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L19
	} else {
		goto L21
	}
L21:
	;
	if base.F64_ne(base.F64_abs(v58), math.Float64frombits(uint64(0x7ff0000000000000))) != 0 {
		goto L6
	} else {
		goto L22
	}
L22:
	;
	goto L19
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v71))) = v14
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+4)) = v74
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v71)+8)) = v76
	v78 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+12)))
	v79 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v71)+14)) = uint8(v79)
	*(*uint16)(unsafe.Add(mBase, uint32(v71)+12)) = uint16(v78)
	return v71
L24:
	;
	base.Wasm_trap_unreachable()
	for {
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
	var v47 int32
	_ = v47
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v95 int32
	_ = v95
	var v107 int32
	_ = v107
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
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v144 int32
	_ = v144
	var v149 int32
	_ = v149
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v238 int32
	_ = v238
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v254 int32
	_ = v254
	var v268 int32
	_ = v268
	var v273 int32
	_ = v273
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
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v339 int32
	_ = v339
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v352 int32
	_ = v352
	var v357 int32
	_ = v357
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v402 int32
	_ = v402
	var v403 int32
	_ = v403
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v418 int32
	_ = v418
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	v23 = l1 + int32(12)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24)+16)))
	v27 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v24+v25)+12)))
	v29 = *(*int32)(unsafe.Add(mBase, _consts[48]))
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
	v34 = int32(4562096)
	v35 = *(*int32)(unsafe.Add(mBase, _consts[3]))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v37)+4))
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v38
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
	v268 = v27 & int32(1)
	v273 = v254
	goto L41
L9:
	;
	v47 = int32(1)
	if int32(2) <= v44+v47 {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v120 = F_palloc0(m, int32(32))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L4
	} else {
		goto L23
	}
L12:
	;
	v54 = v47
	v57 = v42
	goto L15
L13:
	;
	v95 = v42
	goto L14
L14:
	;
	v107 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v111 = F_gistSplit(m, v107, v108, v95, v109, v110)
	mBase = m.M
	v112 = m.ExcPending
	if v112 != 0 {
		goto L4
	} else {
		goto L21
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v23+v54<<(uint(int32(2))%32))))
	v77 = F_gistextractpage(m, v74, v20+int32(8))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L4
	} else {
		goto L17
	}
L16:
	;
	v95 = v80
	goto L14
L17:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v20)+8))
	v80 = F_gistjoinvector(m, v57, v20+int32(12), v77, v79)
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L4
	} else {
		goto L18
	}
L18:
	;
	F_pfree(m, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	v84 = int32(1)
	v85 = v54 + v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v85 < v86+v84 {
		v54 = v85
		v57 = v80
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v35
	v115 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v115
	if v111 == v115 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v254 = v111
	goto L8
L23:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v125 = m.G0
	v127 = v125 - int32(304)
	m.G0 = v127
	F_gistMakeUnionItVec(m, v124, v42, v123, v127+int32(32), v127)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v122)+192))
	v134 = int32(*(*int16)(unsafe.Add(mBase, uint32(v133)+10)))
	if int32(0) < v134 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v144 = v133
	v149 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v124)+12))
	v230 = F_index_form_tuple(m, v227, v127+int32(160), v127)
	mBase = m.M
	v231 = m.ExcPending
	if v231 != 0 {
		goto L4
	} else {
		goto L39
	}
L28:
	;
	v160 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127+v149))))
	if v160 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L27
L30:
	;
	v207 = v149 + int32(1)
	v208 = int32(*(*int16)(unsafe.Add(mBase, uint32(v202)+10)))
	if v207 < v208 {
		v144 = v202
		v149 = v207
		goto L28
	} else {
		goto L38
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127+int32(160)+v149<<(uint(int32(2))%32)))) = int32(0)
	v202 = v144
	goto L30
L32:
	;
	goto L33
L33:
	;
	v170 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v127)+302)) = uint8(v170)
	*(*uint16)(unsafe.Add(mBase, uint32(v127)+300)) = uint16(v170)
	*(*int32)(unsafe.Add(mBase, uint32(v127)+296)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v127)+292)) = v122
	v178 = v149 << (uint(int32(2)) % 32)
	v182 = *(*int32)(unsafe.Add(mBase, uint32(v178+(v127+int32(32)))))
	*(*int32)(unsafe.Add(mBase, uint32(v127)+288)) = v182
	v186 = v124 + int32(1812) + v149*int32(28)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v186)+4))
	if v187 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v178+(v124+int32(8084)))))
	v192 = F_FunctionCall1Coll(m, v186, v189, v127+int32(288))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v196 = v144
	v197 = v182
	goto L36
L36:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v127+int32(160)+v178))) = v197
	v202 = v196
	goto L30
L37:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v192)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v122)+192))
	v196 = v195
	v197 = v194
	goto L36
L38:
	;
	goto L29
L39:
	;
	v232 = int32(65535)
	*(*uint16)(unsafe.Add(mBase, uint32(v230)+4)) = uint16(v232)
	m.G0 = v127 + int32(304)
	*(*int32)(unsafe.Add(mBase, uint32(v120)+16)) = v230
	v238 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	v241 = F_gistfillitupvec(m, v42, v238, v120+int32(12))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+8)) = v241
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v20)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v120)+4)) = v244
	*(*int32)(unsafe.Add(mBase, _consts[3])) = v35
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v254 = v120
	goto L8
L41:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _consts[48]))
	if v287 != 0 {
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
	v289 = m.ExcPending
	if v289 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v273)+8))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v292 = F_smgr_bulk_get_buf(m, v291)
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	F_PageInit(m, v292, int32(8192), int32(16))
	mBase = m.M
	v297 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v292)+16)))
	v298 = v292 + v297
	v299 = int32(65409)
	*(*uint16)(unsafe.Add(mBase, uint32(v298)+14)) = uint16(v299)
	*(*uint16)(unsafe.Add(mBase, uint32(v298)+12)) = uint16(v268)
	*(*int32)(unsafe.Add(mBase, uint32(v298)+8)) = int32(-1)
	goto L48
L48:
	;
	v304 = int32(0)
	v305 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	if v305 <= v304 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v273)+16))
	v376 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v376 != 0 {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	v310 = v290
	v313 = v304
	goto L51
L51:
	;
	v325 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+6)))
	v329 = v313 + int32(1)
	v333 = F_PageAddItemExtended(m, v292, v310, v325&int32(8191), v329&int32(65535), int32(0))
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L4
	} else {
		goto L58
	}
L53:
	;
	if v333 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v335 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v310)+6)))
	v339 = *(*int32)(unsafe.Add(mBase, uint32(v273)+4))
	if v329 < v339 {
		v310 = v310 + v335&int32(8191)
		v313 = v329
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
	v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v346 = *(*int32)(unsafe.Add(mBase, uint32(v345)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = v346 + int32(4)
	F_errmsg_internal(m, int32(747944), v20)
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(525707), int32(562), int32(338722))
	mBase = m.M
	v357 = m.ExcPending
	if v357 != 0 {
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
	v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v292)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v292+v377)+8)) = v376
	goto L63
L62:
	;
	goto L63
L63:
	;
	v380 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v381 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v380 + v381
	*(*int64)(unsafe.Add(mBase, uint32(v292))) = int64(4294967296)
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_smgr_bulk_write(m, v386, v380, v292, v381)
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v375))) = base.I32_rotr(v380, int32(16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v380
	v394 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v394 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v398 = F_palloc0(m, int32(28))
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	v418 = v394
	goto L67
L67:
	;
	F_gist_indexsortbuild_levelstate_add(m, l0, v418, v375)
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L71
	}
L68:
	;
	v401 = F_palloc(m, int32(8192))
	mBase = m.M
	v402 = m.ExcPending
	if v402 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v403 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v398)+8)) = v403
	*(*int32)(unsafe.Add(mBase, uint32(v398)+12)) = v401
	F_PageInit(m, v401, int32(8192), int32(16))
	mBase = m.M
	v410 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v401)+16)))
	v411 = v401 + v410
	v412 = int32(65409)
	*(*uint16)(unsafe.Add(mBase, uint32(v411)+14)) = uint16(v412)
	*(*uint16)(unsafe.Add(mBase, uint32(v411)+12)) = uint16(v403)
	*(*int32)(unsafe.Add(mBase, uint32(v411)+8)) = int32(-1)
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v398
	v418 = v398
	goto L67
L71:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v273)+28))
	if v422 != 0 {
		v273 = v422
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
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 float64
	_ = v210
	var v211 float64
	_ = v211
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
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
	v28 = base.I32_div_u_s(v24&int32(65535), int32(20))
	switch v28 {
	case 0:
		v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v33 = v24 - v28*int32(20)
		switch v33&int32(65535) - int32(1) {
		case 0:
			v210 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
			v211 = *(*float64)(unsafe.Add(mBase, uint32(v30)+16))
			v216 = base.F64_gt(v210, base.F64_add(v211, float64(1e-06)))
			v222 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v222)
			v225 = v216
			m.G0 = v12 + int32(32)
			return v225
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v33 & int32(65535)
				F_errmsg_internal(m, int32(506251), v12+int32(16))
				mBase = m.M
				v116 = m.ExcPending
				if v116 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(526011), int32(1322), int32(326441))
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
			v216 = base.F64_gt(v38, base.F64_add(v39, float64(1e-06)))
			v222 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v222)
			v225 = v216
			m.G0 = v12 + int32(32)
			return v225
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
						v216 = v62
					} else {
						v70 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
						v71 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
						v216 = base.F64_eq(v70, v71) | base.F64_le(base.F64_abs(base.F64_sub(v70, v71)), float64(1e-06))
					}
				} else {
					v70 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
					v71 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
					v216 = base.F64_eq(v70, v71) | base.F64_le(base.F64_abs(base.F64_sub(v70, v71)), float64(1e-06))
				}
			} else {
				v78 = int32(0)
				v79 = *(*float64)(unsafe.Add(mBase, uint32(v29)))
				v80 = *(*float64)(unsafe.Add(mBase, uint32(v30)))
				if base.F64_le(v79, base.F64_add(v80, float64(1e-06))) == v78 {
					v216 = v78
				} else {
					v86 = *(*float64)(unsafe.Add(mBase, uint32(v30)+16))
					if base.F64_le(v86, base.F64_add(v79, float64(1e-06))) == int32(0) {
						v216 = v78
					} else {
						v92 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
						v93 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
						if base.F64_le(v92, base.F64_add(v93, float64(1e-06))) == int32(0) {
							v216 = v78
						} else {
							v99 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
							v216 = base.F64_le(v99, base.F64_add(v92, float64(1e-06)))
						}
					}
				}
			}
			v222 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v222)
			v225 = v216
			m.G0 = v12 + int32(32)
			return v225
		case 9:
			v48 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
			v49 = *(*float64)(unsafe.Add(mBase, uint32(v30)+24))
			v216 = base.F64_gt(v48, base.F64_add(v49, float64(1e-06)))
			v222 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v222)
			v225 = v216
			m.G0 = v12 + int32(32)
			return v225
		case 10:
			v43 = *(*float64)(unsafe.Add(mBase, uint32(v30)+8))
			v44 = *(*float64)(unsafe.Add(mBase, uint32(v29)+8))
			v216 = base.F64_gt(v43, base.F64_add(v44, float64(1e-06)))
			v222 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v222)
			v225 = v216
			m.G0 = v12 + int32(32)
			return v225
		}
	case 1:
		v122 = int32(0)
		v123 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		v124 = *(*float64)(unsafe.Add(mBase, uint32(v123)))
		v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v126 = *(*float64)(unsafe.Add(mBase, uint32(v125)+16))
		if base.F64_ge(v124, v126) == v122 {
			v216 = v122
		} else {
			v130 = *(*float64)(unsafe.Add(mBase, uint32(v123)+16))
			v131 = *(*float64)(unsafe.Add(mBase, uint32(v125)))
			if base.F64_le(v130, v131) == int32(0) {
				v216 = v122
			} else {
				v135 = *(*float64)(unsafe.Add(mBase, uint32(v123)+8))
				v136 = *(*float64)(unsafe.Add(mBase, uint32(v125)+24))
				if base.F64_ge(v135, v136) == int32(0) {
					v216 = v122
				} else {
					v140 = *(*float64)(unsafe.Add(mBase, uint32(v123)+24))
					v141 = *(*float64)(unsafe.Add(mBase, uint32(v125)+8))
					v216 = base.F64_le(v140, v141)
				}
			}
		}
		v222 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v222)
		v225 = v216
		m.G0 = v12 + int32(32)
		return v225
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
				v153 = base.B2i32(v150 != v152)
				v154 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+16)))
				v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v154+v155)+12)))
				if v157&int32(1) == v152 {
					v225 = v153
					m.G0 = v12 + int32(32)
					return v225
				} else {
					if v150 == int32(0) {
						v225 = v153
						m.G0 = v12 + int32(32)
						return v225
					} else {
						v166 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
						v167 = F_DirectFunctionCall2Coll(m, int32(109), int32(0), v146, v166)
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							v216 = base.B2i32(v167 != int32(0))
							v222 = int32(0)
							*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v222)
							v225 = v216
							m.G0 = v12 + int32(32)
							return v225
						}
					}
				}
			}
		}
	case 3:
		v172 = int32(0)
		v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v176 = F_DirectFunctionCall5Coll(m, int32(110), v172, v15, v173, int32(3), v172, v14)
		mBase = m.M
		v177 = m.ExcPending
		if v177 != 0 {
			return int32(0)
		} else {
			v178 = int32(0)
			v179 = base.B2i32(v176 != v178)
			v180 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			v181 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v180)+16)))
			v183 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v180+v181)+12)))
			if v183&int32(1) == v178 {
				v225 = v179
				m.G0 = v12 + int32(32)
				return v225
			} else {
				if v176 == int32(0) {
					v225 = v179
					m.G0 = v12 + int32(32)
					return v225
				} else {
					v192 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
					v193 = F_DirectFunctionCall2Coll(m, int32(111), int32(0), v173, v192)
					mBase = m.M
					v194 = m.ExcPending
					if v194 != 0 {
						return int32(0)
					} else {
						v216 = base.B2i32(v193 != int32(0))
						v222 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v14))) = uint8(v222)
						v225 = v216
						m.G0 = v12 + int32(32)
						return v225
					}
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v200 = m.ExcPending
		if v200 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v24
			F_errmsg_internal(m, int32(506251), v12)
			mBase = m.M
			v204 = m.ExcPending
			if v204 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(526011), int32(1446), int32(98358))
				mBase = m.M
				v209 = m.ExcPending
				if v209 != 0 {
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
