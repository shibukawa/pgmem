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
	var v40 int64
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
	var v64 int64
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	v35 = l3 + v29*int32(24)
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
	v91 = *(*int32)(unsafe.Add(mBase, uint32(l1)+192))
	v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v91)+10)))
	if v37 < v92 {
		v29 = v37
		goto L4
	} else {
		goto L16
	}
L9:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)) = uint8(v87)
	goto L8
L10:
	;
	v45 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v35)+18)) = uint8(v45)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)) = uint16(v45)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v45
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v40
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
	v77 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)) = uint16(v77)
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v77
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = l1
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = int64(0)
	v87 = v77
	goto L9
L13:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(l0+int32(_a_F_gistDeCompressAtt_0)+v29<<(uint(int32(2))%32))))
	v64 = F_FunctionCall1Coll(m, v55, v62, base.I64_extend_i32_u(v35))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L6
	} else {
		goto L14
	}
L14:
	;
	v66 = base.I32_wrap_i64(v64)
	if v35 == v66 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v68 = *(*int64)(unsafe.Add(mBase, uint32(v66)))
	*(*int64)(unsafe.Add(mBase, uint32(v35))) = v68
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+8)) = v70
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v66)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v35)+12)) = v72
	v74 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v66)+16)))
	*(*uint16)(unsafe.Add(mBase, uint32(v35)+16)) = uint16(v74)
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66)+18)))
	v87 = v76
	goto L9
L16:
	;
	goto L5
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
	var v26 int64
	_ = v26
	var v28 int64
	_ = v28
	var v42 int32
	_ = v42
	var v49 int64
	_ = v49
	var v50 int32
	_ = v50
	v7 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v12
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v7
	v17 = v8 + int32(-48)
	v18 = *(*int64)(unsafe.Add(mBase, uint32(l2)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+16)) = v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l2)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v20
	v22 = *(*int64)(unsafe.Add(mBase, uint32(l2)))
	*(*int64)(unsafe.Add(mBase, uint32(v17))) = v22
	v24 = *(*int64)(unsafe.Add(mBase, uint32(l3)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+56)) = v24
	v26 = *(*int64)(unsafe.Add(mBase, uint32(l3)+8))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v26
	v28 = *(*int64)(unsafe.Add(mBase, uint32(l3)))
	*(*int64)(unsafe.Add(mBase, uint32(v10)+40)) = v28
	*(*uint8)(unsafe.Add(mBase, uint32(l5))) = uint8(v7)
	v42 = *(*int32)(unsafe.Add(mBase, uint32(l0+l1<<(uint(v12)%32))+uint32(_c_F_gistMakeUnionKey[0])))
	v49 = F_FunctionCall2Coll(m, l0+l1*int32(28)+int32(916), v42, base.I64_extend_i32_u(v8+int32(-56)), base.I64_extend_i32_u(v8+int32(-60)))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		return
	} else {
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = v49
		m.G0 = v10 - int32(-64)
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
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
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
	v14 = v12 - int32(656)
	m.G0 = v14
	F_check_stack_depth(m)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		return int32(0)
	} else {
		if l3 != int32(1) {
			v23 = v14 + int32(328)
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			v26 = int32(0)
			v27 = base.B2i32(v25 == v26)
			if v27 == v26 {
				base.MemoryFill(m, v23, int32(1), v25)
			} else {
			}
			v33 = v14 + int32(616)
			if v27 == int32(0) {
				base.MemoryFill(m, v33, int32(1), v25)
			} else {
			}
			v38 = int32(0)
			F_gistSplitByKey(m, l0, l1, l2, l3, l4, v14+int32(24), v38)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v46 = l3 + int32(1)
				v47 = F_palloc_mul(m, int32(4), v46)
				mBase = m.M
				v48 = m.ExcPending
				if v48 != 0 {
					return int32(0)
				} else {
					v50 = F_palloc_mul(m, int32(4), v46)
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						v52 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
						if int32(0) < v52 {
							v61 = v38
							for {
								v66 = int32(2)
								v69 = *(*int32)(unsafe.Add(mBase, uint32(v14)+24))
								v70 = int32(1)
								v73 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v69+v61<<(uint(v70)%32)))))
								v79 = *(*int32)(unsafe.Add(mBase, uint32(l2+v73<<(uint(v66)%32)-int32(4))))
								*(*int32)(unsafe.Add(mBase, uint32(v47+v61<<(uint(v66)%32)))) = v79
								v82 = v61 + v70
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
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
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
						if int32(0) < v96 {
							v106 = int32(0)
							for {
								v111 = int32(2)
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
								v115 = int32(1)
								v118 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v114+v106<<(uint(v115)%32)))))
								v124 = *(*int32)(unsafe.Add(mBase, uint32(l2+v118<<(uint(v111)%32)-int32(4))))
								*(*int32)(unsafe.Add(mBase, uint32(v50+v106<<(uint(v111)%32)))) = v124
								v127 = v106 + v115
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
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
							v221 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
							v222 = F_gistSplit(m, l0, l1, v50, v221, l4)
							mBase = m.M
							v223 = m.ExcPending
							if v223 != 0 {
								return int32(0)
							} else {
								v245 = v222
								v246 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
								v247 = int32(0)
								if v247 < v246 {
									if v246 != int32(1) {
										v262 = int32(0)
										v263 = v247
										v264 = v247
										for {
											v268 = int32(2)
											v270 = v47 + v263<<(uint(v268)%32)
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
											v300 = *(*int32)(unsafe.Add(mBase, uint32(v47+v292<<(uint(int32(2))%32))))
											v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+6)))
											v310 = v301&int32(_a_F_gistSplit_0) + v293 + int32(4)
										}
									} else {
										v292 = v247
										v293 = v247
										v300 = *(*int32)(unsafe.Add(mBase, uint32(v47+v292<<(uint(int32(2))%32))))
										v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+6)))
										v310 = v301&int32(_a_F_gistSplit_0) + v293 + int32(4)
									}
									v324 = base.B2i32(base.Ui32(v310) < base.Ui32(int32(_a_F_gistSplit_1)))
								} else {
									v324 = int32(1)
								}
								if v324 == int32(0) {
									v327 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
									v328 = F_gistSplit(m, l0, l1, v47, v327, l4)
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
										m.G0 = v14 + int32(656)
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
										v351 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
										*(*int32)(unsafe.Add(mBase, uint32(v344)+4)) = v351
										v355 = F_gistfillitupvec(m, v47, v351, v344+int32(12))
										mBase = m.M
										v356 = m.ExcPending
										if v356 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v344)+8)) = v355
											v361 = F_gistFormTuple(m, l4, l0, v14+int32(72), v23, int32(0))
											mBase = m.M
											v362 = m.ExcPending
											if v362 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v344)+16)) = v361
												v365 = v344
												m.G0 = v14 + int32(656)
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
								v231 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
								*(*int32)(unsafe.Add(mBase, uint32(v225)+4)) = v231
								v235 = F_gistfillitupvec(m, v50, v231, v225+int32(12))
								mBase = m.M
								v236 = m.ExcPending
								if v236 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v225)+8)) = v235
									v241 = F_gistFormTuple(m, l4, l0, v14+int32(360), v33, int32(0))
									mBase = m.M
									v242 = m.ExcPending
									if v242 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = v241
										v245 = v225
										v246 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
										v247 = int32(0)
										if v247 < v246 {
											if v246 != int32(1) {
												v262 = int32(0)
												v263 = v247
												v264 = v247
												for {
													v268 = int32(2)
													v270 = v47 + v263<<(uint(v268)%32)
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
													v300 = *(*int32)(unsafe.Add(mBase, uint32(v47+v292<<(uint(int32(2))%32))))
													v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+6)))
													v310 = v301&int32(_a_F_gistSplit_0) + v293 + int32(4)
												}
											} else {
												v292 = v247
												v293 = v247
												v300 = *(*int32)(unsafe.Add(mBase, uint32(v47+v292<<(uint(int32(2))%32))))
												v301 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+6)))
												v310 = v301&int32(_a_F_gistSplit_0) + v293 + int32(4)
											}
											v324 = base.B2i32(base.Ui32(v310) < base.Ui32(int32(_a_F_gistSplit_1)))
										} else {
											v324 = int32(1)
										}
										if v324 == int32(0) {
											v327 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
											v328 = F_gistSplit(m, l0, l1, v47, v327, l4)
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
												m.G0 = v14 + int32(656)
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
												v351 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
												*(*int32)(unsafe.Add(mBase, uint32(v344)+4)) = v351
												v355 = F_gistfillitupvec(m, v47, v351, v344+int32(12))
												mBase = m.M
												v356 = m.ExcPending
												if v356 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v344)+8)) = v355
													v361 = F_gistFormTuple(m, l4, l0, v14+int32(72), v23, int32(0))
													mBase = m.M
													v362 = m.ExcPending
													if v362 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v344)+16)) = v361
														v365 = v344
														m.G0 = v14 + int32(656)
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
func F_gist_bbox_zorder_cmp(m *base.Module, l0 int64, l1 int64, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v13 float64
	_ = v13
	var v14 int32
	_ = v14
	var v15 float64
	_ = v15
	var v17 float64
	_ = v17
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
	var v63 int32
	_ = v63
	var v72 int32
	_ = v72
	var v75 int64
	_ = v75
	var v80 int64
	_ = v80
	var v85 int64
	_ = v85
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v95 int64
	_ = v95
	var v103 float64
	_ = v103
	var v105 int64
	_ = v105
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v120 int64
	_ = v120
	var v121 int32
	_ = v121
	var v130 int32
	_ = v130
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v144 int64
	_ = v144
	var v145 int64
	_ = v145
	var v148 int64
	_ = v148
	var v149 int64
	_ = v149
	var v150 int64
	_ = v150
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v157 int64
	_ = v157
	var v166 int64
	_ = v166
	var v171 int64
	_ = v171
	var v176 int64
	_ = v176
	var v181 int64
	_ = v181
	var v187 int64
	_ = v187
	var v194 int32
	_ = v194
	v11 = base.I32_wrap_i64(l0)
	v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)+24))
	v13 = *(*float64)(unsafe.Add(mBase, uint32(v11)+16))
	v14 = base.I32_wrap_i64(l1)
	v15 = *(*float64)(unsafe.Add(mBase, uint32(v14)+16))
	if base.F64_ne(v13, v15) != 0 {
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
		v63 = base.I32_reinterpret_f32(base.F32_demote_f64(v12))
		if base.Ui32(v63&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			if v63 < int32(0) {
				v72 = int32(-1)
			} else {
				v72 = int32(-2147483648)
			}
			v75 = base.I64_extend_i32_u(v63 ^ v72)
		} else {
			v75 = v21
		}
		v80 = (v75<<(uint(int64(16))%64) | v75) & int64(281470681808895)
		v85 = (v80<<(uint(int64(8))%64) | v80) & int64(71777214294589695)
		v90 = (v85<<(uint(int64(4))%64) | v85) & int64(1085102592571150095)
		v91 = int64(2)
		v95 = (v90<<(uint(v91)%64) | v90) & int64(3689348814741910323)
		v103 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
		v105 = int64(4294967295)
		v108 = base.I32_reinterpret_f32(base.F32_demote_f64(v15))
		if base.Ui32(v108&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			if v108 < int32(0) {
				v117 = int32(-1)
			} else {
				v117 = int32(-2147483648)
			}
			v120 = base.I64_extend_i32_u(v108 ^ v117)
		} else {
			v120 = v105
		}
		v121 = base.I32_reinterpret_f32(base.F32_demote_f64(v103))
		if base.Ui32(v121&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
			if v121 < int32(0) {
				v130 = int32(-1)
			} else {
				v130 = int32(-2147483648)
			}
			v133 = base.I64_extend_i32_u(v121 ^ v130)
		} else {
			v133 = v105
		}
		v134 = (v56<<(uint(int64(1))%64)|v56)&int64(6148914691236517205) | (v95<<(uint(v91)%64)|v95<<(uint(int64(1))%64))&int64(-6148914691236517206)
		v135 = int64(16)
		v138 = int64(281470681808895)
		v139 = (v133<<(uint(v135)%64) | v133) & v138
		v140 = int64(8)
		v143 = int64(71777214294589695)
		v144 = (v139<<(uint(v140)%64) | v139) & v143
		v145 = int64(4)
		v148 = int64(1085102592571150095)
		v149 = (v144<<(uint(v145)%64) | v144) & v148
		v150 = int64(2)
		v153 = int64(3689348814741910323)
		v154 = (v149<<(uint(v150)%64) | v149) & v153
		v157 = int64(1)
		v166 = (v120<<(uint(v135)%64) | v120) & v138
		v171 = (v166<<(uint(v140)%64) | v166) & v143
		v176 = (v171<<(uint(v145)%64) | v171) & v148
		v181 = (v176<<(uint(v150)%64) | v176) & v153
		v187 = (v154<<(uint(v150)%64)|v154<<(uint(v157)%64))&int64(-6148914691236517206) | (v181<<(uint(v157)%64)|v181)&int64(6148914691236517205)
		if base.Ui64(v187) < base.Ui64(v134) {
			return int32(1)
		} else {
			if base.Ui64(v134) < base.Ui64(v187) {
				v194 = int32(-1)
			} else {
				v194 = int32(0)
			}
			return v194
		}
	} else {
		v17 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
		if base.F64_ne(v12, v17) != 0 {
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
			v63 = base.I32_reinterpret_f32(base.F32_demote_f64(v12))
			if base.Ui32(v63&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
				if v63 < int32(0) {
					v72 = int32(-1)
				} else {
					v72 = int32(-2147483648)
				}
				v75 = base.I64_extend_i32_u(v63 ^ v72)
			} else {
				v75 = v21
			}
			v80 = (v75<<(uint(int64(16))%64) | v75) & int64(281470681808895)
			v85 = (v80<<(uint(int64(8))%64) | v80) & int64(71777214294589695)
			v90 = (v85<<(uint(int64(4))%64) | v85) & int64(1085102592571150095)
			v91 = int64(2)
			v95 = (v90<<(uint(v91)%64) | v90) & int64(3689348814741910323)
			v103 = *(*float64)(unsafe.Add(mBase, uint32(v14)+24))
			v105 = int64(4294967295)
			v108 = base.I32_reinterpret_f32(base.F32_demote_f64(v15))
			if base.Ui32(v108&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
				if v108 < int32(0) {
					v117 = int32(-1)
				} else {
					v117 = int32(-2147483648)
				}
				v120 = base.I64_extend_i32_u(v108 ^ v117)
			} else {
				v120 = v105
			}
			v121 = base.I32_reinterpret_f32(base.F32_demote_f64(v103))
			if base.Ui32(v121&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
				if v121 < int32(0) {
					v130 = int32(-1)
				} else {
					v130 = int32(-2147483648)
				}
				v133 = base.I64_extend_i32_u(v121 ^ v130)
			} else {
				v133 = v105
			}
			v134 = (v56<<(uint(int64(1))%64)|v56)&int64(6148914691236517205) | (v95<<(uint(v91)%64)|v95<<(uint(int64(1))%64))&int64(-6148914691236517206)
			v135 = int64(16)
			v138 = int64(281470681808895)
			v139 = (v133<<(uint(v135)%64) | v133) & v138
			v140 = int64(8)
			v143 = int64(71777214294589695)
			v144 = (v139<<(uint(v140)%64) | v139) & v143
			v145 = int64(4)
			v148 = int64(1085102592571150095)
			v149 = (v144<<(uint(v145)%64) | v144) & v148
			v150 = int64(2)
			v153 = int64(3689348814741910323)
			v154 = (v149<<(uint(v150)%64) | v149) & v153
			v157 = int64(1)
			v166 = (v120<<(uint(v135)%64) | v120) & v138
			v171 = (v166<<(uint(v140)%64) | v166) & v143
			v176 = (v171<<(uint(v145)%64) | v171) & v148
			v181 = (v176<<(uint(v150)%64) | v176) & v153
			v187 = (v154<<(uint(v150)%64)|v154<<(uint(v157)%64))&int64(-6148914691236517206) | (v181<<(uint(v157)%64)|v181)&int64(6148914691236517205)
			if base.Ui64(v187) < base.Ui64(v134) {
				return int32(1)
			} else {
				if base.Ui64(v134) < base.Ui64(v187) {
					v194 = int32(-1)
				} else {
					v194 = int32(0)
				}
				return v194
			}
		} else {
			return int32(0)
		}
	}
}
func F_gist_box_distance(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = base.I32_wrap_i64(v9) & int32(_a_F_gist_box_distance_0)
	if base.Ui32(int32(20)) <= base.Ui32(v12) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = v12
			F_errmsg_internal(m, int32(_a_F_gist_box_distance_1), v7)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_gist_box_distance_2), int32(1495), int32(_a_F_gist_box_distance_3))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v31)))
		v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v34 = F_computeDistance(m, int32(0), v32, v33)
		mBase = m.M
		v35 = m.ExcPending
		if v35 != 0 {
			return int64(0)
		} else {
			m.G0 = v7 + int32(16)
			return base.I64_reinterpret_f64(v34)
		}
	}
}
func F_gist_box_union(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v11 int64
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
	v11 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	v15 = F_palloc(m, int32(32))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int64(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
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
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v12+int32(8)+v41*int32(24))))
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
		*(*int32)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v11)))) = int32(32)
		return base.I64_extend_i32_u(v15)
	}
}
func F_gist_circle_compress(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 float64
	_ = v19
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v34 float64
	_ = v34
	var v35 int32
	_ = v35
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v39 float64
	_ = v39
	var v40 float64
	_ = v40
	var v42 float64
	_ = v42
	var v53 float64
	_ = v53
	var v54 int32
	_ = v54
	var v55 float64
	_ = v55
	var v57 float64
	_ = v57
	var v58 float64
	_ = v58
	var v59 float64
	_ = v59
	var v61 float64
	_ = v61
	var v72 float64
	_ = v72
	var v73 int32
	_ = v73
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v78 float64
	_ = v78
	var v80 float64
	_ = v80
	var v91 float64
	_ = v91
	var v92 int32
	_ = v92
	var v93 float64
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+18)))
	if v8 != int32(1) {
		return base.I64_extend_i32_u(v7)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v15 = F_palloc(m, int32(32))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
			v20 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
			v21 = base.F64_add(v19, v20)
			v23 = math.Float64frombits(uint64(0x7ff0000000000000))
			if base.F64_ne(base.F64_abs(v21), v23)|base.F64_eq(base.F64_abs(v19), v23)|base.F64_eq(base.F64_abs(v20), v23) != 0 {
				v36 = v21
				*(*float64)(unsafe.Add(mBase, uint32(v15))) = v36
				v38 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
				v39 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
				v40 = base.F64_sub(v38, v39)
				v42 = math.Float64frombits(uint64(0x7ff0000000000000))
				if base.F64_ne(base.F64_abs(v40), v42)|base.F64_eq(base.F64_abs(v38), v42)|base.F64_eq(base.F64_abs(v39), v42) != 0 {
					v55 = v40
					*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v55
					v57 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
					v58 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
					v59 = base.F64_add(v57, v58)
					v61 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v59), v61)|base.F64_eq(base.F64_abs(v57), v61)|base.F64_eq(base.F64_abs(v58), v61) != 0 {
						v74 = v59
						*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v74
						v76 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
						v77 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
						v78 = base.F64_sub(v76, v77)
						v80 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v78), v80)|base.F64_eq(base.F64_abs(v76), v80)|base.F64_eq(base.F64_abs(v77), v80) != 0 {
							v93 = v78
							*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
							v96 = F_palloc(m, int32(24))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int64(0)
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
								v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
								v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
								v105 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
								*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
								return base.I64_extend_i32_u(v96)
							}
						} else {
							v91 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int64(0)
							} else {
								v93 = v91
								*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
								v96 = F_palloc(m, int32(24))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
									v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
									v105 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
									*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
									return base.I64_extend_i32_u(v96)
								}
							}
						}
					} else {
						v72 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int64(0)
						} else {
							v74 = v72
							*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v74
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
							v77 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
							v78 = base.F64_sub(v76, v77)
							v80 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v78), v80)|base.F64_eq(base.F64_abs(v76), v80)|base.F64_eq(base.F64_abs(v77), v80) != 0 {
								v93 = v78
								*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
								v96 = F_palloc(m, int32(24))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
									v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
									v105 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
									*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
									return base.I64_extend_i32_u(v96)
								}
							} else {
								v91 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = v91
									*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
									v96 = F_palloc(m, int32(24))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
										v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
										v105 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
										*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
										return base.I64_extend_i32_u(v96)
									}
								}
							}
						}
					}
				} else {
					v53 = F_float_overflow_error_ext(m, int32(0))
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int64(0)
					} else {
						v55 = v53
						*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v55
						v57 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
						v58 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
						v59 = base.F64_add(v57, v58)
						v61 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v59), v61)|base.F64_eq(base.F64_abs(v57), v61)|base.F64_eq(base.F64_abs(v58), v61) != 0 {
							v74 = v59
							*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v74
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
							v77 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
							v78 = base.F64_sub(v76, v77)
							v80 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v78), v80)|base.F64_eq(base.F64_abs(v76), v80)|base.F64_eq(base.F64_abs(v77), v80) != 0 {
								v93 = v78
								*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
								v96 = F_palloc(m, int32(24))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
									v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
									v105 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
									*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
									return base.I64_extend_i32_u(v96)
								}
							} else {
								v91 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = v91
									*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
									v96 = F_palloc(m, int32(24))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
										v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
										v105 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
										*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
										return base.I64_extend_i32_u(v96)
									}
								}
							}
						} else {
							v72 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								v74 = v72
								*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v74
								v76 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
								v77 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
								v78 = base.F64_sub(v76, v77)
								v80 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_ne(base.F64_abs(v78), v80)|base.F64_eq(base.F64_abs(v76), v80)|base.F64_eq(base.F64_abs(v77), v80) != 0 {
									v93 = v78
									*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
									v96 = F_palloc(m, int32(24))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
										v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
										v105 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
										*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
										return base.I64_extend_i32_u(v96)
									}
								} else {
									v91 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										v93 = v91
										*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
										v96 = F_palloc(m, int32(24))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
											v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
											v105 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
											*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
											return base.I64_extend_i32_u(v96)
										}
									}
								}
							}
						}
					}
				}
			} else {
				v34 = F_float_overflow_error_ext(m, int32(0))
				mBase = m.M
				v35 = m.ExcPending
				if v35 != 0 {
					return int64(0)
				} else {
					v36 = v34
					*(*float64)(unsafe.Add(mBase, uint32(v15))) = v36
					v38 = *(*float64)(unsafe.Add(mBase, uint32(v13)))
					v39 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
					v40 = base.F64_sub(v38, v39)
					v42 = math.Float64frombits(uint64(0x7ff0000000000000))
					if base.F64_ne(base.F64_abs(v40), v42)|base.F64_eq(base.F64_abs(v38), v42)|base.F64_eq(base.F64_abs(v39), v42) != 0 {
						v55 = v40
						*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v55
						v57 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
						v58 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
						v59 = base.F64_add(v57, v58)
						v61 = math.Float64frombits(uint64(0x7ff0000000000000))
						if base.F64_ne(base.F64_abs(v59), v61)|base.F64_eq(base.F64_abs(v57), v61)|base.F64_eq(base.F64_abs(v58), v61) != 0 {
							v74 = v59
							*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v74
							v76 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
							v77 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
							v78 = base.F64_sub(v76, v77)
							v80 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v78), v80)|base.F64_eq(base.F64_abs(v76), v80)|base.F64_eq(base.F64_abs(v77), v80) != 0 {
								v93 = v78
								*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
								v96 = F_palloc(m, int32(24))
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int64(0)
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
									v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
									v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
									*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
									v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
									v105 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
									*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
									return base.I64_extend_i32_u(v96)
								}
							} else {
								v91 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = v91
									*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
									v96 = F_palloc(m, int32(24))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
										v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
										v105 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
										*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
										return base.I64_extend_i32_u(v96)
									}
								}
							}
						} else {
							v72 = F_float_overflow_error_ext(m, int32(0))
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int64(0)
							} else {
								v74 = v72
								*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v74
								v76 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
								v77 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
								v78 = base.F64_sub(v76, v77)
								v80 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_ne(base.F64_abs(v78), v80)|base.F64_eq(base.F64_abs(v76), v80)|base.F64_eq(base.F64_abs(v77), v80) != 0 {
									v93 = v78
									*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
									v96 = F_palloc(m, int32(24))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
										v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
										v105 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
										*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
										return base.I64_extend_i32_u(v96)
									}
								} else {
									v91 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										v93 = v91
										*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
										v96 = F_palloc(m, int32(24))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
											v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
											v105 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
											*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
											return base.I64_extend_i32_u(v96)
										}
									}
								}
							}
						}
					} else {
						v53 = F_float_overflow_error_ext(m, int32(0))
						mBase = m.M
						v54 = m.ExcPending
						if v54 != 0 {
							return int64(0)
						} else {
							v55 = v53
							*(*float64)(unsafe.Add(mBase, uint32(v15)+16)) = v55
							v57 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
							v58 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
							v59 = base.F64_add(v57, v58)
							v61 = math.Float64frombits(uint64(0x7ff0000000000000))
							if base.F64_ne(base.F64_abs(v59), v61)|base.F64_eq(base.F64_abs(v57), v61)|base.F64_eq(base.F64_abs(v58), v61) != 0 {
								v74 = v59
								*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v74
								v76 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
								v77 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
								v78 = base.F64_sub(v76, v77)
								v80 = math.Float64frombits(uint64(0x7ff0000000000000))
								if base.F64_ne(base.F64_abs(v78), v80)|base.F64_eq(base.F64_abs(v76), v80)|base.F64_eq(base.F64_abs(v77), v80) != 0 {
									v93 = v78
									*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
									v96 = F_palloc(m, int32(24))
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
										v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
										*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
										v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
										v105 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
										*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
										return base.I64_extend_i32_u(v96)
									}
								} else {
									v91 = F_float_overflow_error_ext(m, int32(0))
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										v93 = v91
										*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
										v96 = F_palloc(m, int32(24))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
											v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
											v105 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
											*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
											return base.I64_extend_i32_u(v96)
										}
									}
								}
							} else {
								v72 = F_float_overflow_error_ext(m, int32(0))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int64(0)
								} else {
									v74 = v72
									*(*float64)(unsafe.Add(mBase, uint32(v15)+8)) = v74
									v76 = *(*float64)(unsafe.Add(mBase, uint32(v13)+8))
									v77 = *(*float64)(unsafe.Add(mBase, uint32(v13)+16))
									v78 = base.F64_sub(v76, v77)
									v80 = math.Float64frombits(uint64(0x7ff0000000000000))
									if base.F64_ne(base.F64_abs(v78), v80)|base.F64_eq(base.F64_abs(v76), v80)|base.F64_eq(base.F64_abs(v77), v80) != 0 {
										v93 = v78
										*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
										v96 = F_palloc(m, int32(24))
										mBase = m.M
										v97 = m.ExcPending
										if v97 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
											v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
											*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
											v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
											v105 = int32(0)
											*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
											*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
											return base.I64_extend_i32_u(v96)
										}
									} else {
										v91 = F_float_overflow_error_ext(m, int32(0))
										mBase = m.M
										v92 = m.ExcPending
										if v92 != 0 {
											return int64(0)
										} else {
											v93 = v91
											*(*float64)(unsafe.Add(mBase, uint32(v15)+24)) = v93
											v96 = F_palloc(m, int32(24))
											mBase = m.M
											v97 = m.ExcPending
											if v97 != 0 {
												return int64(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v96))) = base.I64_extend_i32_u(v15)
												v100 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v96)+8)) = v100
												v102 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
												*(*int32)(unsafe.Add(mBase, uint32(v96)+12)) = v102
												v104 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+16)))
												v105 = int32(0)
												*(*uint8)(unsafe.Add(mBase, uint32(v96)+18)) = uint8(v105)
												*(*uint16)(unsafe.Add(mBase, uint32(v96)+16)) = uint16(v104)
												return base.I64_extend_i32_u(v96)
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
func F_gist_indexsortbuild_levelstate_flush(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
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
	var v113 int32
	_ = v113
	var v116 int32
	_ = v116
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
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	var v164 int32
	_ = v164
	var v174 int32
	_ = v174
	var v182 int32
	_ = v182
	var v186 int64
	_ = v186
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int64
	_ = v196
	var v197 int32
	_ = v197
	var v199 int64
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int64
	_ = v202
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
	var v260 int32
	_ = v260
	var v275 int32
	_ = v275
	var v280 int32
	_ = v280
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v312 int32
	_ = v312
	var v313 int32
	_ = v313
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v353 int32
	_ = v353
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v361 int32
	_ = v361
	var v366 int32
	_ = v366
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v428 int32
	_ = v428
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	v19 = m.G0
	v21 = v19 - int32(16)
	m.G0 = v21
	v24 = l1 + int32(12)
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25)+16)))
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v25+v26)+12)))
	v30 = *(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[0]))
	if v30 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v34 = v25
	goto L3
L3:
	;
	v35 = int32(_a_F_gist_indexsortbuild_levelstate_flush_0)
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[1]))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	*(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[1])) = v39
	v43 = F_gistextractpage(m, v34, v21+int32(12))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L4
	} else {
		goto L6
	}
L4:
	;
	return
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v34 = v33
	goto L3
L6:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if int32(0) < v45 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	m.G0 = v21 + int32(16)
	return
L8:
	;
	v275 = v28 & int32(1)
	v280 = v260
	goto L41
L9:
	;
	if v45 != int32(2147483647) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	goto L11
L11:
	;
	v121 = F_palloc0(m, int32(32))
	mBase = m.M
	v122 = m.ExcPending
	if v122 != 0 {
		goto L4
	} else {
		goto L23
	}
L12:
	;
	v53 = int32(1)
	v56 = v43
	goto L15
L13:
	;
	v95 = v43
	goto L14
L14:
	;
	v108 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v112 = F_gistSplit(m, v108, v109, v95, v110, v111)
	mBase = m.M
	v113 = m.ExcPending
	if v113 != 0 {
		goto L4
	} else {
		goto L21
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v24+v53<<(uint(int32(2))%32))))
	v77 = F_gistextractpage(m, v74, v21+int32(8))
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
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v21)+8))
	v80 = F_gistjoinvector(m, v56, v21+int32(12), v77, v79)
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
	v85 = v53 + v84
	v86 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v85 < v86+v84 {
		v53 = v85
		v56 = v80
		goto L15
	} else {
		goto L20
	}
L20:
	;
	goto L16
L21:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[1])) = v36
	v116 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = v116
	if v112 == v116 {
		goto L7
	} else {
		goto L22
	}
L22:
	;
	v260 = v112
	goto L8
L23:
	;
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v125 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v126 = m.G0
	v128 = v126 - int32(576)
	m.G0 = v128
	F_gistMakeUnionItVec(m, v125, v43, v124, v128+int32(32), v128)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L24
	}
L24:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v123)+192))
	v135 = int32(*(*int16)(unsafe.Add(mBase, uint32(v134)+10)))
	if int32(0) < v135 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v147 = v134
	v152 = int32(0)
	goto L28
L26:
	;
	goto L27
L27:
	;
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v125)+12))
	v236 = F_index_form_tuple(m, v233, v128+int32(288), v128)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L4
	} else {
		goto L39
	}
L28:
	;
	v164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v128+v152))))
	if v164 == int32(1) {
		goto L31
	} else {
		goto L32
	}
L29:
	;
	goto L27
L30:
	;
	v212 = v152 + int32(1)
	v213 = int32(*(*int16)(unsafe.Add(mBase, uint32(v207)+10)))
	if v212 < v213 {
		v147 = v207
		v152 = v212
		goto L28
	} else {
		goto L38
	}
L31:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v128+int32(288)+v152<<(uint(int32(3))%32)))) = int64(0)
	v207 = v147
	goto L30
L32:
	;
	goto L33
L33:
	;
	v174 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v128)+570)) = uint8(v174)
	*(*uint16)(unsafe.Add(mBase, uint32(v128)+568)) = uint16(v174)
	*(*int32)(unsafe.Add(mBase, uint32(v128)+564)) = v174
	*(*int32)(unsafe.Add(mBase, uint32(v128)+560)) = v123
	v182 = v152 << (uint(int32(3)) % 32)
	v186 = *(*int64)(unsafe.Add(mBase, uint32(v182+(v128+int32(32)))))
	*(*int64)(unsafe.Add(mBase, uint32(v128)+552)) = v186
	v190 = v125 + int32(1812) + v152*int32(28)
	v191 = *(*int32)(unsafe.Add(mBase, uint32(v190)+4))
	if v191 != 0 {
		goto L34
	} else {
		goto L35
	}
L34:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v125+int32(_a_F_gist_indexsortbuild_levelstate_flush_1)+v152<<(uint(int32(2))%32))))
	v196 = F_FunctionCall1Coll(m, v190, v195, base.I64_extend_i32_u(v128+int32(552)))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L4
	} else {
		goto L37
	}
L35:
	;
	v201 = v147
	v202 = v186
	goto L36
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v128+int32(288)+v182))) = v202
	v207 = v201
	goto L30
L37:
	;
	v199 = *(*int64)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v196))))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v123)+192))
	v201 = v200
	v202 = v199
	goto L36
L38:
	;
	goto L29
L39:
	;
	v238 = int32(_a_F_gist_indexsortbuild_levelstate_flush_2)
	*(*uint16)(unsafe.Add(mBase, uint32(v236)+4)) = uint16(v238)
	m.G0 = v128 + int32(576)
	*(*int32)(unsafe.Add(mBase, uint32(v121)+16)) = v236
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	v247 = F_gistfillitupvec(m, v43, v244, v121+int32(12))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L4
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v121)+8)) = v247
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v121)+4)) = v250
	*(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[1])) = v36
	*(*int32)(unsafe.Add(mBase, uint32(l1))) = int32(0)
	v260 = v121
	goto L8
L41:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _c_F_gist_indexsortbuild_levelstate_flush[0]))
	if v295 != 0 {
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
	v297 = m.ExcPending
	if v297 != 0 {
		goto L4
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(v280)+8))
	v299 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v300 = F_smgr_bulk_get_buf(m, v299)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L4
	} else {
		goto L47
	}
L46:
	;
	goto L45
L47:
	;
	F_PageInit(m, v300, int32(_a_F_gist_indexsortbuild_levelstate_flush_3), int32(16))
	mBase = m.M
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+16)))
	v306 = v300 + v305
	v307 = int32(_a_F_gist_indexsortbuild_levelstate_flush_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+14)) = uint16(v307)
	*(*uint16)(unsafe.Add(mBase, uint32(v306)+12)) = uint16(v275)
	*(*int32)(unsafe.Add(mBase, uint32(v306)+8)) = int32(-1)
	goto L48
L48:
	;
	v312 = int32(0)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	if v313 <= v312 {
		goto L49
	} else {
		goto L50
	}
L49:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v280)+16))
	v386 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v386 != 0 {
		goto L61
	} else {
		goto L62
	}
L50:
	;
	v318 = v298
	v321 = v312
	goto L51
L51:
	;
	v334 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318)+6)))
	v338 = v321 + int32(1)
	v342 = F_PageAddItemExtended(m, v300, v318, v334&int32(_a_F_gist_indexsortbuild_levelstate_flush_5), v338&int32(_a_F_gist_indexsortbuild_levelstate_flush_2), int32(0))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v353 = m.ExcPending
	if v353 != 0 {
		goto L4
	} else {
		goto L58
	}
L53:
	;
	if v342 != 0 {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v344 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v318)+6)))
	v348 = *(*int32)(unsafe.Add(mBase, uint32(v280)+4))
	if v338 < v348 {
		v318 = v318 + v344&int32(_a_F_gist_indexsortbuild_levelstate_flush_5)
		v321 = v338
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
	v354 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v354)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v21))) = v355 + int32(4)
	F_errmsg_internal(m, int32(_a_F_gist_indexsortbuild_levelstate_flush_6), v21)
	mBase = m.M
	v361 = m.ExcPending
	if v361 != 0 {
		goto L4
	} else {
		goto L59
	}
L59:
	;
	F_errfinish(m, int32(_a_F_gist_indexsortbuild_levelstate_flush_7), int32(562), int32(_a_F_gist_indexsortbuild_levelstate_flush_8))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
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
	v387 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v300)+16)))
	*(*int32)(unsafe.Add(mBase, uint32(v300+v387)+8)) = v386
	goto L63
L62:
	;
	goto L63
L63:
	;
	v390 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v391 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v390 + v391
	*(*int64)(unsafe.Add(mBase, uint32(v300))) = int64(4294967296)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	F_smgr_bulk_write(m, v396, v390, v300, v391)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L4
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v385))) = base.I32_rotr(v390, int32(16))
	*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v390
	v404 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if v404 == int32(0) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v408 = F_palloc0(m, int32(28))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L4
	} else {
		goto L68
	}
L66:
	;
	v428 = v404
	goto L67
L67:
	;
	F_gist_indexsortbuild_levelstate_add(m, l0, v428, v385)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L71
	}
L68:
	;
	v411 = F_palloc(m, int32(_a_F_gist_indexsortbuild_levelstate_flush_3))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	v413 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v408)+8)) = v413
	*(*int32)(unsafe.Add(mBase, uint32(v408)+12)) = v411
	F_PageInit(m, v411, int32(_a_F_gist_indexsortbuild_levelstate_flush_3), int32(16))
	mBase = m.M
	v420 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v411)+16)))
	v421 = v411 + v420
	v422 = int32(_a_F_gist_indexsortbuild_levelstate_flush_4)
	*(*uint16)(unsafe.Add(mBase, uint32(v421)+14)) = uint16(v422)
	*(*uint16)(unsafe.Add(mBase, uint32(v421)+12)) = uint16(v413)
	*(*int32)(unsafe.Add(mBase, uint32(v421)+8)) = int32(-1)
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v408
	v428 = v408
	goto L67
L71:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v280)+28))
	if v432 != 0 {
		v280 = v432
		goto L41
	} else {
		goto L72
	}
L72:
	;
	goto L42
}
func F_gist_point_consistent(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	var v20 int64
	_ = v20
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 float64
	_ = v46
	var v47 float64
	_ = v47
	var v51 float64
	_ = v51
	var v52 float64
	_ = v52
	var v56 float64
	_ = v56
	var v57 float64
	_ = v57
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v67 float64
	_ = v67
	var v68 float64
	_ = v68
	var v70 int32
	_ = v70
	var v78 float64
	_ = v78
	var v79 float64
	_ = v79
	var v86 int32
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v94 float64
	_ = v94
	var v100 float64
	_ = v100
	var v101 float64
	_ = v101
	var v107 float64
	_ = v107
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 float64
	_ = v132
	var v133 int32
	_ = v133
	var v134 float64
	_ = v134
	var v138 float64
	_ = v138
	var v139 float64
	_ = v139
	var v143 float64
	_ = v143
	var v144 float64
	_ = v144
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int64
	_ = v156
	var v159 int64
	_ = v159
	var v160 int32
	_ = v160
	var v161 int64
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v178 int32
	_ = v178
	var v183 int64
	_ = v183
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int64
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v203 int64
	_ = v203
	var v204 int64
	_ = v204
	var v205 int32
	_ = v205
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v222 float64
	_ = v222
	var v227 int32
	_ = v227
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	v13 = m.G0
	v15 = v13 - int32(32)
	m.G0 = v15
	v17 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
	v18 = int64(4294967295)
	v19 = v17 & v18
	v20 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v22 = v20 & v18
	v23 = base.I32_wrap_i64(v20)
	v26 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+56)))
	if v26 == int32(30) {
		v29 = int32(11)
	} else {
		v29 = v26
	}
	if v26 == int32(29) {
		v32 = int32(10)
	} else {
		v32 = v29
	}
	v34 = v32 & int32(_a_F_gist_point_consistent_0)
	v36 = base.I32_div_u_s(v34, int32(20))
	switch v36 {
	case 0:
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v38 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		v41 = v32 - v36*int32(20)
		switch v41&int32(_a_F_gist_point_consistent_0) - int32(1) {
		case 0:
			v221 = *(*float64)(unsafe.Add(mBase, uint32(v37)))
			v222 = *(*float64)(unsafe.Add(mBase, uint32(v38)+16))
			v227 = base.F64_gt(v221, base.F64_add(v222, float64(1e-06)))
			v234 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17)))) = uint8(v234)
			v237 = v227
			m.G0 = v15 + int32(32)
			return base.I64_extend_i32_u(v237)
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v116 = m.ExcPending
			if v116 != 0 {
				return int64(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v41 & int32(_a_F_gist_point_consistent_0)
				F_errmsg_internal(m, int32(_a_F_gist_point_consistent_1), v15+int32(16))
				mBase = m.M
				v124 = m.ExcPending
				if v124 != 0 {
					return int64(0)
				} else {
					F_errfinish(m, int32(_a_F_gist_point_consistent_2), int32(1325), int32(_a_F_gist_point_consistent_3))
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return int64(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		case 4:
			v46 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
			v47 = *(*float64)(unsafe.Add(mBase, uint32(v37)))
			v227 = base.F64_gt(v46, base.F64_add(v47, float64(1e-06)))
			v234 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17)))) = uint8(v234)
			v237 = v227
			m.G0 = v15 + int32(32)
			return base.I64_extend_i32_u(v237)
		case 5:
			v61 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
			v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v61)+16)))
			v64 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61+v62)+12)))
			if v64&int32(1) != 0 {
				v67 = *(*float64)(unsafe.Add(mBase, uint32(v38)+16))
				v68 = *(*float64)(unsafe.Add(mBase, uint32(v37)))
				if base.F64_ne(v67, v68) != 0 {
					v70 = int32(0)
					if base.F64_le(base.F64_abs(base.F64_sub(v67, v68)), float64(1e-06)) == v70 {
						v227 = v70
					} else {
						v78 = *(*float64)(unsafe.Add(mBase, uint32(v38)+24))
						v79 = *(*float64)(unsafe.Add(mBase, uint32(v37)+8))
						v227 = base.F64_eq(v78, v79) | base.F64_le(base.F64_abs(base.F64_sub(v78, v79)), float64(1e-06))
					}
				} else {
					v78 = *(*float64)(unsafe.Add(mBase, uint32(v38)+24))
					v79 = *(*float64)(unsafe.Add(mBase, uint32(v37)+8))
					v227 = base.F64_eq(v78, v79) | base.F64_le(base.F64_abs(base.F64_sub(v78, v79)), float64(1e-06))
				}
			} else {
				v86 = int32(0)
				v87 = *(*float64)(unsafe.Add(mBase, uint32(v37)))
				v88 = *(*float64)(unsafe.Add(mBase, uint32(v38)))
				if base.F64_le(v87, base.F64_add(v88, float64(1e-06))) == v86 {
					v227 = v86
				} else {
					v94 = *(*float64)(unsafe.Add(mBase, uint32(v38)+16))
					if base.F64_le(v94, base.F64_add(v87, float64(1e-06))) == int32(0) {
						v227 = v86
					} else {
						v100 = *(*float64)(unsafe.Add(mBase, uint32(v37)+8))
						v101 = *(*float64)(unsafe.Add(mBase, uint32(v38)+8))
						if base.F64_le(v100, base.F64_add(v101, float64(1e-06))) == int32(0) {
							v227 = v86
						} else {
							v107 = *(*float64)(unsafe.Add(mBase, uint32(v38)+24))
							v227 = base.F64_le(v107, base.F64_add(v100, float64(1e-06)))
						}
					}
				}
			}
			v234 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17)))) = uint8(v234)
			v237 = v227
			m.G0 = v15 + int32(32)
			return base.I64_extend_i32_u(v237)
		case 9:
			v56 = *(*float64)(unsafe.Add(mBase, uint32(v37)+8))
			v57 = *(*float64)(unsafe.Add(mBase, uint32(v38)+24))
			v227 = base.F64_gt(v56, base.F64_add(v57, float64(1e-06)))
			v234 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17)))) = uint8(v234)
			v237 = v227
			m.G0 = v15 + int32(32)
			return base.I64_extend_i32_u(v237)
		case 10:
			v51 = *(*float64)(unsafe.Add(mBase, uint32(v38)+8))
			v52 = *(*float64)(unsafe.Add(mBase, uint32(v37)+8))
			v227 = base.F64_gt(v51, base.F64_add(v52, float64(1e-06)))
			v234 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17)))) = uint8(v234)
			v237 = v227
			m.G0 = v15 + int32(32)
			return base.I64_extend_i32_u(v237)
		}
	case 1:
		v130 = int32(0)
		v131 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
		v132 = *(*float64)(unsafe.Add(mBase, uint32(v131)))
		v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v134 = *(*float64)(unsafe.Add(mBase, uint32(v133)+16))
		if base.F64_ge(v132, v134) == v130 {
			v227 = v130
		} else {
			v138 = *(*float64)(unsafe.Add(mBase, uint32(v131)+16))
			v139 = *(*float64)(unsafe.Add(mBase, uint32(v133)))
			if base.F64_le(v138, v139) == int32(0) {
				v227 = v130
			} else {
				v143 = *(*float64)(unsafe.Add(mBase, uint32(v131)+8))
				v144 = *(*float64)(unsafe.Add(mBase, uint32(v133)+24))
				if base.F64_ge(v143, v144) == int32(0) {
					v227 = v130
				} else {
					v148 = *(*float64)(unsafe.Add(mBase, uint32(v131)+24))
					v149 = *(*float64)(unsafe.Add(mBase, uint32(v133)+8))
					v227 = base.F64_le(v148, v149)
				}
			}
		}
		v234 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17)))) = uint8(v234)
		v237 = v227
		m.G0 = v15 + int32(32)
		return base.I64_extend_i32_u(v237)
	case 2:
		v153 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v154 = F_pg_detoast_datum(m, v153)
		mBase = m.M
		v155 = m.ExcPending
		if v155 != 0 {
			return int64(0)
		} else {
			v156 = base.I64_extend_i32_u(v154)
			v159 = F_DirectFunctionCall5Coll(m, int32(110), int32(0), v22, v156, int64(3), int64(0), v19)
			mBase = m.M
			v160 = m.ExcPending
			if v160 != 0 {
				return int64(0)
			} else {
				v161 = int64(0)
				v163 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
				v164 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v163)+16)))
				v166 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v163+v164)+12)))
				if base.B2i32(v166&int32(1) == int32(0))|base.B2i32(v159 == v161) != 0 {
					v237 = base.B2i32(v159 != v161)
					m.G0 = v15 + int32(32)
					return base.I64_extend_i32_u(v237)
				} else {
					v176 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23))))
					v177 = F_DirectFunctionCall2Coll(m, int32(111), int32(0), v156, v176)
					mBase = m.M
					v178 = m.ExcPending
					if v178 != 0 {
						return int64(0)
					} else {
						v227 = base.B2i32(v177 != int64(0))
						v234 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17)))) = uint8(v234)
						v237 = v227
						m.G0 = v15 + int32(32)
						return base.I64_extend_i32_u(v237)
					}
				}
			}
		}
	case 3:
		v183 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+40)))
		v186 = F_DirectFunctionCall5Coll(m, int32(112), int32(0), v22, v183, int64(3), int64(0), v19)
		mBase = m.M
		v187 = m.ExcPending
		if v187 != 0 {
			return int64(0)
		} else {
			v188 = int64(0)
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
			v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v190)+16)))
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v190+v191)+12)))
			if base.B2i32(v193&int32(1) == int32(0))|base.B2i32(v186 == v188) != 0 {
				v237 = base.B2i32(v186 != v188)
				m.G0 = v15 + int32(32)
				return base.I64_extend_i32_u(v237)
			} else {
				v203 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v23))))
				v204 = F_DirectFunctionCall2Coll(m, int32(113), int32(0), v183, v203)
				mBase = m.M
				v205 = m.ExcPending
				if v205 != 0 {
					return int64(0)
				} else {
					v227 = base.B2i32(v204 != int64(0))
					v234 = int32(0)
					*(*uint8)(unsafe.Add(mBase, uint32(base.I32_wrap_i64(v17)))) = uint8(v234)
					v237 = v227
					m.G0 = v15 + int32(32)
					return base.I64_extend_i32_u(v237)
				}
			}
		}
	default:
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v211 = m.ExcPending
		if v211 != 0 {
			return int64(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v15))) = v34
			F_errmsg_internal(m, int32(_a_F_gist_point_consistent_1), v15)
			mBase = m.M
			v215 = m.ExcPending
			if v215 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_gist_point_consistent_2), int32(1449), int32(_a_F_gist_point_consistent_4))
				mBase = m.M
				v220 = m.ExcPending
				if v220 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	}
}
