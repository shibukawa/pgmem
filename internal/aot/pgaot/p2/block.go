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
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
	var v29 int32
	_ = v29
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int64
	_ = v33
	var v34 int64
	_ = v34
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v40 int64
	_ = v40
	var v43 int32
	_ = v43
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v82 int64
	_ = v82
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	v8 = int64(0)
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v14 = v9 + int32(-32)
	*(*int64)(unsafe.Add(mBase, uint32(v14))) = v8
	v18 = v9 + int32(-40)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v8
	*(*int64)(unsafe.Add(mBase, uint32(v11)+16)) = v8
	F_BlockRefTableRead(m, l0, v9+int32(-24), int32(24))
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		v30 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
		v31 = *(*int64)(unsafe.Add(mBase, uint32(v11)+16))
		v33 = *(*int64)(unsafe.Add(mBase, uint32(v11)+48))
		v34 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
		v37 = *(*int64)(unsafe.Add(mBase, uint32(v11)+56))
		v38 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v40 = v30 ^ v31 | (v33 ^ v34) | (v37 ^ v38)
		if v40 == int64(0) {
			v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[0])))
			F_BlockRefTableRead(m, l0, v9+int32(-52), int32(4))
			mBase = m.M
			v48 = m.ExcPending
			if v48 != 0 {
				return int32(0)
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
				v51 = v43 ^ int32(-1)
				if v49 == v51 {
					m.G0 = v11 - int32(-64)
					return base.B2i32(v40 != int64(0))
				} else {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[1])))
					v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[2])))
					v55 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[3])))
					*(*int32)(unsafe.Add(mBase, uint32(v11)+8)) = v49
					*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v51
					*(*int32)(unsafe.Add(mBase, uint32(v11))) = v55
					m.T0[v54].(func(*base.Module, int32, int32, int32))(m, v53, int32(_a_F_BlockRefTableReaderNextRelation_0), v11)
					mBase = m.M
					v61 = m.ExcPending
					if v61 != 0 {
						return int32(0)
					} else {
						m.G0 = v11 - int32(-64)
						return base.B2i32(v40 != int64(0))
					}
				}
			}
		} else {
			v62 = *(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[4])))
			if v62 != 0 {
				F_pfree(m, v62)
				mBase = m.M
				v64 = m.ExcPending
				if v64 != 0 {
					return int32(0)
				} else {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
					v68 = F_palloc(m, v65<<(uint(int32(1))%32))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[4]))) = v68
						v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
						F_BlockRefTableRead(m, l0, v68, v71<<(uint(int32(1))%32))
						mBase = m.M
						v75 = m.ExcPending
						if v75 != 0 {
							return int32(0)
						} else {
							v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
							*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[5]))) = int32(0)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[6]))) = v76
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
							*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v80
							v82 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
							*(*int64)(unsafe.Add(mBase, uint32(l1))) = v82
							v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
							*(*int32)(unsafe.Add(mBase, uint32(l2))) = v84
							v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
							*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
							m.G0 = v11 - int32(-64)
							return base.B2i32(v40 != int64(0))
						}
					}
				}
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
				v68 = F_palloc(m, v65<<(uint(int32(1))%32))
				mBase = m.M
				v69 = m.ExcPending
				if v69 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[4]))) = v68
					v71 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
					F_BlockRefTableRead(m, l0, v68, v71<<(uint(int32(1))%32))
					mBase = m.M
					v75 = m.ExcPending
					if v75 != 0 {
						return int32(0)
					} else {
						v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+60))
						*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[5]))) = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+uint32(_c_F_BlockRefTableReaderNextRelation[6]))) = v76
						v80 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
						*(*int32)(unsafe.Add(mBase, uint32(l1)+8)) = v80
						v82 = *(*int64)(unsafe.Add(mBase, uint32(v11)+40))
						*(*int64)(unsafe.Add(mBase, uint32(l1))) = v82
						v84 = *(*int32)(unsafe.Add(mBase, uint32(v11)+52))
						*(*int32)(unsafe.Add(mBase, uint32(l2))) = v84
						v86 = *(*int32)(unsafe.Add(mBase, uint32(v11)+56))
						*(*int32)(unsafe.Add(mBase, uint32(l3))) = v86
						m.G0 = v11 - int32(-64)
						return base.B2i32(v40 != int64(0))
					}
				}
			}
		}
	}
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
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v249 int32
	_ = v249
	var v256 int64
	_ = v256
	var v260 int64
	_ = v260
	var v266 int32
	_ = v266
	var v270 int64
	_ = v270
	var v276 int64
	_ = v276
	var v282 int32
	_ = v282
	var v284 int64
	_ = v284
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v300 int32
	_ = v300
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v347 int32
	_ = v347
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v379 int32
	_ = v379
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
	return v379
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
	v379 = int32(0)
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
		v311 = v51
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v379 = int32(0)
	goto L1
L13:
	;
	v312 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+38)))
	if v312 == int32(0) {
		goto L76
	} else {
		goto L77
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
	v73 = v61 + v64
	v74 = v51 + v59
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
		goto L66
	} else {
		goto L67
	}
L18:
	;
	if int32(0) <= v249 {
		v311 = v61
		goto L13
	} else {
		goto L64
	}
L19:
	;
	goto L18
L20:
	;
	goto L60
L21:
	;
	v82 = v51
	v85 = v61
	goto L24
L22:
	;
	goto L23
L23:
	;
	v224 = v51
	v227 = v61
	goto L20
L24:
	;
	v95 = v82 + int32(1)
	if base.Ui32(v74) <= base.Ui32(v95) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v224 = v210
	v227 = v213
	goto L20
L26:
	;
	if base.Ui32(v74) <= base.Ui32(v210) {
		v224 = v210
		v227 = v213
		goto L20
	} else {
		goto L58
	}
L27:
	;
	v210 = v95
	v213 = v85
	goto L26
L28:
	;
	goto L29
L29:
	;
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	v100 = v95
	v102 = v85
	v108 = v97
	v109 = int32(0)
	goto L30
L30:
	;
	if v108&int32(1) != 0 {
		goto L33
	} else {
		goto L34
	}
L31:
	;
	v210 = v186
	v213 = v198
	goto L26
L32:
	;
	if base.B2i32(base.Ui32(int32(6)) < base.Ui32(v109))|base.B2i32(base.Ui32(v74) <= base.Ui32(v186)) != 0 {
		v210 = v186
		v213 = v198
		goto L26
	} else {
		goto L56
	}
L33:
	;
	v113 = int32(-1)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	v118 = v114&int32(15) + int32(3)
	if v118 != int32(18) {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	goto L35
L35:
	;
	v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100))))
	*(*uint8)(unsafe.Add(mBase, uint32(v102))) = uint8(v180)
	v182 = int32(1)
	v186 = v100 + v182
	v198 = v102 + v182
	goto L32
L36:
	;
	v128 = v118
	v129 = v100 + int32(2)
	goto L38
L37:
	;
	v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+2)))
	v128 = v123 + int32(18)
	v129 = v100 + int32(3)
	goto L38
L38:
	;
	if base.Ui32(v74) < base.Ui32(v129) {
		v249 = v113
		goto L19
	} else {
		goto L39
	}
L39:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v100)+1)))
	v136 = v131 | v114<<(uint(int32(4))%32)&int32(3840)
	if base.B2i32(v136 == int32(0))|base.B2i32(v102-v61 < v136) != 0 {
		v249 = v113
		goto L19
	} else {
		goto L40
	}
L40:
	;
	v142 = v73 - v102
	if v128 < v142 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v144 = v128
	goto L43
L42:
	;
	v144 = v142
	goto L43
L43:
	;
	if v136 < v144 {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v147 = v136
	v149 = v102
	v151 = v144
	goto L47
L45:
	;
	v166 = v136
	v168 = v102
	v170 = v144
	goto L46
L46:
	;
	if v170 != 0 {
		goto L53
	} else {
		goto L54
	}
L47:
	;
	if v147 != 0 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v166 = v163
	v168 = v160
	v170 = v161
	goto L46
L49:
	;
	base.MemoryCopy(m, v149, v149-v147, v147)
	goto L51
L50:
	;
	goto L51
L51:
	;
	v160 = v147 + v149
	v161 = v151 - v147
	v163 = v147 << (uint(int32(1)) % 32)
	if v163 < v161 {
		v147 = v163
		v149 = v160
		v151 = v161
		goto L47
	} else {
		goto L52
	}
L52:
	;
	goto L48
L53:
	;
	base.MemoryCopy(m, v168, v168-v166, v170)
	goto L55
L54:
	;
	goto L55
L55:
	;
	v186 = v129
	v198 = v168 + v170
	goto L32
L56:
	;
	v203 = int32(1)
	if base.Ui32(v198) < base.Ui32(v73) {
		v100 = v186
		v102 = v198
		v108 = int32(base.Ui32(v108&int32(254)) >> (uint(v203) % 32))
		v109 = v109 + v203
		goto L30
	} else {
		goto L57
	}
L57:
	;
	goto L31
L58:
	;
	if base.Ui32(v213) < base.Ui32(v73) {
		v82 = v210
		v85 = v213
		goto L24
	} else {
		goto L59
	}
L59:
	;
	goto L25
L60:
	;
	if base.B2i32(v224 != v74)|base.B2i32(v227 != v73) != 0 {
		v249 = int32(-1)
		goto L19
	} else {
		goto L63
	}
L62:
	;
	v249 = v227 - v61
	goto L19
L63:
	;
	goto L62
L64:
	;
	v256 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = l1
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+68)) = uint32(v256)
	v260 = int64(base.Ui64(v256) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+64)) = uint32(v260)
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_4), v12-int32(-64))
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L7
	} else {
		goto L65
	}
L65:
	;
	v379 = int32(0)
	goto L1
L66:
	;
	v270 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(_a_F_RestoreBlockImage_5)
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+52)) = uint32(v270)
	v276 = int64(base.Ui64(v270) >> (uint(int64(32)) % 64))
	*(*uint32)(unsafe.Add(mBase, uint32(v12)+48)) = uint32(v276)
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_6), v12+int32(48))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L7
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	v284 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
	v287 = base.I32_wrap_i64(int64(base.Ui64(v284) >> (uint(int64(32)) % 64)))
	v288 = base.I32_wrap_i64(v284)
	if v52&int32(16) != 0 {
		goto L70
	} else {
		goto L71
	}
L69:
	;
	v379 = int32(0)
	goto L1
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = int32(_a_F_RestoreBlockImage_7)
	*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v287
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_6), v12+int32(32))
	mBase = m.M
	v300 = m.ExcPending
	if v300 != 0 {
		goto L7
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = l1
	*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v288
	*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v287
	F_report_invalid_record(m, l0, int32(_a_F_RestoreBlockImage_8), v12+int32(16))
	mBase = m.M
	v309 = m.ExcPending
	if v309 != 0 {
		goto L7
	} else {
		goto L74
	}
L73:
	;
	v379 = int32(0)
	goto L1
L74:
	;
	v379 = int32(0)
	goto L1
L75:
	;
	v379 = int32(1)
	goto L1
L76:
	;
	base.MemoryCopy(m, l2, v311, int32(_a_F_RestoreBlockImage_3))
	goto L75
L77:
	;
	goto L78
L78:
	;
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)))
	if v317 != 0 {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	base.MemoryCopy(m, l2, v311, v317)
	goto L81
L80:
	;
	goto L81
L81:
	;
	v319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)))
	v320 = l2 + v319
	v321 = int32(3)
	v323 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+38)))
	if v320&v321|v323&v321|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v323)) == int32(0) {
		goto L83
	} else {
		goto L84
	}
L82:
	;
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+36)))
	v357 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v35)+38)))
	v358 = v356 + v357
	v359 = int32(_a_F_RestoreBlockImage_3) - v358
	if v359 == int32(0) {
		goto L75
	} else {
		goto L91
	}
L83:
	;
	if v323 == int32(0) {
		goto L82
	} else {
		goto L86
	}
L84:
	;
	v347 = v323
	goto L85
L85:
	;
	if v347 == int32(0) {
		goto L82
	} else {
		goto L90
	}
L86:
	;
	v336 = v323 + v320
	v338 = v320 + int32(4)
	if base.Ui32(v338) < base.Ui32(v336) {
		goto L87
	} else {
		goto L88
	}
L87:
	;
	v340 = v336
	goto L89
L88:
	;
	v340 = v338
	goto L89
L89:
	;
	v347 = (l2^int32(-1)+v340-v319)&int32(-4) + int32(4)
	goto L85
L90:
	;
	base.MemoryFill(m, v320, int32(0), v347)
	goto L82
L91:
	;
	base.MemoryCopy(m, v358+l2, v356+v311, v359)
	goto L75
}
