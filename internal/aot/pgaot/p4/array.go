package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_array_agg_array_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v3 == int32(0) {
		v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v6 != 0 {
			v13 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			v15 = F_makeArrayResultArr(m, v6, v13, int32(0))
			mBase = m.M
			v18 = m.ExcPending
			if v18 != 0 {
				return int32(0)
			} else {
				return v15
			}
		} else {
			v8 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
			return int32(0)
		}
	} else {
		v8 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v8)
		return int32(0)
	}
}
func F_array_agg_finalfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 == int32(0) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v11 != 0 {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
			v17 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v6)+8)) = v17
			*(*int32)(unsafe.Add(mBase, uint32(v6)+12)) = v16
			v26 = *(*int32)(unsafe.Add(mBase, _consts[28]))
			v28 = F_makeMdArrayResult(m, v11, v17, v6+int32(12), v6+int32(8), v26, int32(0))
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				v34 = v28
				m.G0 = v6 + int32(16)
				return v34
			}
		} else {
			v13 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
			v34 = int32(0)
			m.G0 = v6 + int32(16)
			return v34
		}
	} else {
		v13 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v13)
		v34 = int32(0)
		m.G0 = v6 + int32(16)
		return v34
	}
}
func F_array_agg_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v48 int64
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v61 int64
	_ = v61
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v240 int32
	_ = v240
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	v9 = m.G0
	v11 = v9 - int32(32)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v11+int32(16))
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_enlargeStringInfo(m, v11+int32(16), int32(4))
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v29 = int32(24)
	v31 = int32(65280)
	v33 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v26+v27))) = v20<<(uint(v29)%32) | v20&v31<<(uint(v33)%32) | (int32(base.Ui32(v20)>>(uint(v33)%32))&v31 | int32(base.Ui32(v20)>>(uint(v29)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v26 + int32(4)
	v48 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+16)))
	F_enlargeStringInfo(m, v11+int32(16), v33)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v57 = int64(56)
	v59 = int64(65280)
	v61 = int64(40)
	v64 = int64(16711680)
	v66 = int64(24)
	v68 = int64(4278190080)
	v70 = int64(8)
	*(*int64)(unsafe.Add(mBase, uint32(v54+v55))) = v48<<(uint(v57)%64) | v48&v59<<(uint(v61)%64) | (v48&v64<<(uint(v66)%64) | v48&v68<<(uint(v70)%64)) | (int64(base.Ui64(v48)>>(uint(v70)%64))&v68 | int64(base.Ui64(v48)>>(uint(v66)%64))&v64 | (int64(base.Ui64(v48)>>(uint(v61)%64))&v59 | int64(base.Ui64(v48)>>(uint(v57)%64))))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v54 + int32(8)
	v96 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+24)))
	F_enlargeStringInfo(m, v11+int32(16), int32(2))
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	v102 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v105 = int32(8)
	v109 = v96<<(uint(v105)%32) | int32(base.Ui32(v96)>>(uint(v105)%32))
	*(*uint16)(unsafe.Add(mBase, uint32(v102+v103))) = uint16(v109)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v102 + int32(2)
	v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)))
	F_enlargeStringInfo(m, v11+int32(16), int32(1))
	mBase = m.M
	v119 = m.ExcPending
	if v119 != 0 {
		goto L1
	} else {
		goto L6
	}
L6:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v120+v121))) = uint8(v114)
	v124 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v120 + v124
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+27)))
	F_enlargeStringInfo(m, v11+int32(16), v124)
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L1
	} else {
		goto L7
	}
L7:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	*(*uint8)(unsafe.Add(mBase, uint32(v133+v134))) = uint8(v127)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v133 + int32(1)
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	F_pq_sendbytes(m, v11+int32(16), v142, v143)
	mBase = m.M
	v145 = m.ExcPending
	if v145 != 0 {
		goto L1
	} else {
		goto L8
	}
L8:
	;
	v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13)+26)))
	if v146 == int32(1) {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	v263 = v11 + int32(16)
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v263)))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v263)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v265))) = v266 << (uint(int32(2)) % 32)
	goto L30
L10:
	;
	v151 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	F_pq_sendbytes(m, v11+int32(16), v151, v152<<(uint(int32(2))%32))
	mBase = m.M
	v156 = m.ExcPending
	if v156 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v157)+16))
	if v158 == int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	goto L9
L14:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v157)+20))
	v163 = F_MemoryContextAlloc(m, v161, int32(28))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L17
	}
L15:
	;
	v179 = v158
	goto L16
L16:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	if v180 <= int32(0) {
		goto L9
	} else {
		goto L20
	}
L17:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v13)+20))
	F_getTypeBinaryOutputInfo(m, v165, v11+int32(12), v11+int32(11))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v11)+12))
	v173 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v174 = *(*int32)(unsafe.Add(mBase, uint32(v173)+20))
	F_fmgr_info_cxt(m, v172, v163, v174)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v177 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v177)+16)) = v163
	v179 = v163
	goto L16
L20:
	;
	v184 = v180
	v186 = int32(0)
	goto L21
L21:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v192+v186))))
	if v194 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	goto L9
L23:
	;
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+v186<<(uint(int32(2))%32))))
	v202 = F_SendFunctionCall(m, v179, v201)
	mBase = m.M
	v203 = m.ExcPending
	if v203 != 0 {
		goto L1
	} else {
		goto L26
	}
L24:
	;
	v248 = v184
	goto L25
L25:
	;
	v252 = v186 + int32(1)
	if v252 < v248 {
		v184 = v248
		v186 = v252
		goto L21
	} else {
		goto L29
	}
L26:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	F_enlargeStringInfo(m, v11+int32(16), int32(4))
	mBase = m.M
	v209 = m.ExcPending
	if v209 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v210 = *(*int32)(unsafe.Add(mBase, uint32(v11)+20))
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v11)+16))
	v213 = int32(2)
	v215 = int32(4)
	v216 = int32(base.Ui32(v204)>>(uint(v213)%32)) - v215
	v217 = int32(24)
	v219 = int32(65280)
	v221 = int32(8)
	*(*int32)(unsafe.Add(mBase, uint32(v210+v211))) = v216<<(uint(v217)%32) | v216&v219<<(uint(v221)%32) | (int32(base.Ui32(v216)>>(uint(v221)%32))&v219 | int32(base.Ui32(v216)>>(uint(v217)%32)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+20)) = v210 + v215
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v202)))
	F_pq_sendbytes(m, v11+int32(16), v202+v215, int32(base.Ui32(v240)>>(uint(v213)%32))-v215)
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
	v248 = v247
	goto L25
L29:
	;
	goto L22
L30:
	;
	m.G0 = v11 + int32(32)
	return v265
}
func F_array_agg_transfn(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v97 int32
	_ = v97
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v12 = F_get_fn_expr_argtype(m, v10, int32(1))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		if v12 != 0 {
			v17 = v8 + int32(12)
			v18 = int32(0)
			v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v19 == v18 {
				v36 = int32(0)
				if v17 == v36 {
					v44 = v36
				} else {
					v39 = v36
					v40 = v18
					*(*int32)(unsafe.Add(mBase, uint32(v17))) = v39
					v44 = v40
				}
				v47 = v44
			} else {
				v22 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				switch v22 - int32(429) {
				case 0:
					if v17 == int32(0) {
						v47 = int32(1)
					} else {
						v28 = *(*int32)(unsafe.Add(mBase, uint32(v19)+168))
						v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+20))
						v39 = v29
						v40 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v17))) = v39
						v44 = v40
						v47 = v44
					}
				case 1:
					if v17 == int32(0) {
						v47 = int32(2)
					} else {
						v34 = *(*int32)(unsafe.Add(mBase, uint32(v19)+368))
						v39 = v34
						v40 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v17))) = v39
						v44 = v40
						v47 = v44
					}
				default:
					v36 = int32(0)
					if v17 == v36 {
						v44 = v36
					} else {
						v39 = v36
						v40 = v18
						*(*int32)(unsafe.Add(mBase, uint32(v17))) = v39
						v44 = v40
					}
					v47 = v44
				}
			}
			if v47 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(67046), int32(0))
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(518919), int32(576), int32(294209))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
				if v50 == int32(1) {
					v53 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					v55 = F_initArrayResult(m, v12, v53, int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						v58 = v55
						v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
						if v59 != 0 {
							v61 = int32(0)
						} else {
							v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
							v61 = v60
						}
						v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
						v63 = F_accumArrayResult(m, v58, v61, v59, v12, v62)
						mBase = m.M
						v64 = m.ExcPending
						if v64 != 0 {
							return int32(0)
						} else {
							m.G0 = v8 + int32(16)
							return v63
						}
					}
				} else {
					v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					v58 = v57
					v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v59 != 0 {
						v61 = int32(0)
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v61 = v60
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v8)+12))
					v63 = F_accumArrayResult(m, v58, v61, v59, v12, v62)
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						m.G0 = v8 + int32(16)
						return v63
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v72 = m.ExcPending
			if v72 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v75 = m.ExcPending
				if v75 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(389108), int32(0))
					mBase = m.M
					v79 = m.ExcPending
					if v79 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(518919), int32(565), int32(294209))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
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
func F_array_cat(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
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
	var v53 int32
	_ = v53
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
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
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
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v130 int32
	_ = v130
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
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
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	var v331 int32
	_ = v331
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v358 int32
	_ = v358
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v370 int32
	_ = v370
	var v374 int32
	_ = v374
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v392 int32
	_ = v392
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v415 int32
	_ = v415
	var v416 int32
	_ = v416
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v483 int32
	_ = v483
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v494 int32
	_ = v494
	var v501 int32
	_ = v501
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v552 int32
	_ = v552
	var v555 int32
	_ = v555
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v589 int32
	_ = v589
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v601 int32
	_ = v601
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v610 int32
	_ = v610
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v623 int32
	_ = v623
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v627 int32
	_ = v627
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v634 int32
	_ = v634
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v650 int32
	_ = v650
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v680 int32
	_ = v680
	var v685 int32
	_ = v685
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v704 int32
	_ = v704
	var v705 int32
	_ = v705
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v723 int32
	_ = v723
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v735 int32
	_ = v735
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v774 int32
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v814 int32
	_ = v814
	var v819 int32
	_ = v819
	var v833 int32
	_ = v833
	v28 = m.G0
	v30 = v28 - int32(32)
	m.G0 = v30
	v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v33 == int32(1) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v30 + int32(32)
	return v833
L2:
	;
	if v32&int32(1) != 0 {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	goto L4
L4:
	;
	if v32&int32(1) != 0 {
		goto L10
	} else {
		goto L11
	}
L5:
	;
	v38 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v38)
	v833 = int32(0)
	goto L1
L6:
	;
	goto L7
L7:
	;
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v42 = F_pg_detoast_datum(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L8
	} else {
		goto L9
	}
L8:
	;
	return int32(0)
L9:
	;
	v833 = v42
	goto L1
L10:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v49 = F_pg_detoast_datum(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L8
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v52 = F_pg_detoast_datum(m, v51)
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L8
	} else {
		goto L14
	}
L13:
	;
	v833 = v49
	goto L1
L14:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v55 = F_pg_detoast_datum(m, v54)
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L8
	} else {
		goto L15
	}
L15:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v52)+12))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v55)+12))
	if v57 == v58 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v461 = int32(3)
	v462 = v60 << (uint(v461) % 32)
	v463 = int32(23)
	v465 = int32(-8)
	v468 = v61 << (uint(v461) % 32)
	if v74 != 0 {
		goto L120
	} else {
		goto L121
	}
L17:
	;
	v442 = v415
	v443 = v416
	v460 = v61
	goto L16
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L8
	} else {
		goto L115
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v367 = m.ExcPending
	if v367 != 0 {
		goto L8
	} else {
		goto L110
	}
L20:
	;
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	if v61 != 0 {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L8
	} else {
		goto L103
	}
L23:
	;
	if v60 == int32(0) {
		v833 = v52
		goto L1
	} else {
		goto L26
	}
L24:
	;
	if v60 <= int32(0) {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v833 = v55
	goto L1
L26:
	;
	v67 = v60 - int32(1)
	if v61 == v60 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	v76 = v52 + int32(16)
	v77 = F_ArrayGetNItems(m, v61, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L8
	} else {
		goto L31
	}
L28:
	;
	if v61 == v67 {
		goto L27
	} else {
		goto L29
	}
L29:
	;
	if v61 != v60+int32(1) {
		goto L19
	} else {
		goto L30
	}
L30:
	;
	goto L27
L31:
	;
	v80 = v55 + int32(16)
	v81 = F_ArrayGetNItems(m, v60, v80)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L8
	} else {
		goto L32
	}
L32:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v83 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	v93 = (v86<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L35
L34:
	;
	v93 = v83
	goto L35
L35:
	;
	v94 = int32(2)
	v95 = v60 << (uint(v94) % 32)
	v97 = v61 << (uint(v94) % 32)
	v98 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v98 == int32(0) {
		goto L36
	} else {
		goto L37
	}
L36:
	;
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
	v108 = (v101<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L38
L37:
	;
	v108 = v98
	goto L38
L38:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
	v111 = v95 + v80
	v112 = v97 + v76
	if v61 == v60 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v114 = F_palloc(m, v97)
	mBase = m.M
	v115 = m.ExcPending
	if v115 != 0 {
		goto L8
	} else {
		goto L42
	}
L40:
	;
	goto L41
L41:
	;
	if v61 == v67 {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v116 = F_palloc(m, v97)
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L8
	} else {
		goto L43
	}
L43:
	;
	v118 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	*(*int32)(unsafe.Add(mBase, uint32(v114))) = v118 + v119
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v112)))
	*(*int32)(unsafe.Add(mBase, uint32(v116))) = v122
	if v61 < int32(2) {
		v415 = v114
		v416 = v116
		goto L17
	} else {
		goto L44
	}
L44:
	;
	v130 = int32(1)
	goto L45
L45:
	;
	v155 = v130 << (uint(int32(2)) % 32)
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v76+v155)))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155+v80)))
	if v157 != v159 {
		goto L18
	} else {
		goto L47
	}
L46:
	;
	v415 = v114
	v416 = v116
	goto L17
L47:
	;
	v161 = v155 + v112
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v155+v111)))
	if v162 != v164 {
		goto L18
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v155+v114))) = v157
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	*(*int32)(unsafe.Add(mBase, uint32(v155+v116))) = v169
	v172 = v130 + int32(1)
	if v172 != v61 {
		v130 = v172
		goto L45
	} else {
		goto L49
	}
L49:
	;
	goto L46
L50:
	;
	v175 = F_palloc(m, v95)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L8
	} else {
		goto L53
	}
L51:
	;
	goto L52
L52:
	;
	v256 = F_palloc(m, v97)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L8
	} else {
		goto L78
	}
L53:
	;
	v177 = F_palloc(m, v95)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L8
	} else {
		goto L54
	}
L54:
	;
	if v95 != 0 {
		goto L56
	} else {
		goto L57
	}
L55:
	;
	if v95 != 0 {
		goto L60
	} else {
		goto L61
	}
L56:
	;
	v179 = F__emscripten_memcpy_bulkmem(m, v175, v80, v95)
	mBase = m.M
	v180 = v179
	goto L58
L57:
	;
	v180 = v175
	goto L58
L58:
	;
	goto L55
L59:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v180)))
	*(*int32)(unsafe.Add(mBase, uint32(v180))) = v183 + int32(1)
	v187 = int32(0)
	if v187 < v61 {
		goto L63
	} else {
		goto L64
	}
L60:
	;
	v181 = F__emscripten_memcpy_bulkmem(m, v177, v111, v95)
	mBase = m.M
	v182 = v181
	goto L62
L61:
	;
	v182 = v177
	goto L62
L62:
	;
	goto L59
L63:
	;
	v191 = v61
	goto L65
L64:
	;
	v191 = v187
	goto L65
L65:
	;
	v192 = v187
	goto L66
L66:
	;
	if v192 == v191 {
		v442 = v175
		v443 = v177
		v460 = v60
		goto L16
	} else {
		goto L68
	}
L67:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L8
	} else {
		goto L73
	}
L68:
	;
	v220 = int32(2)
	v221 = v192 << (uint(v220) % 32)
	v223 = *(*int32)(unsafe.Add(mBase, uint32(v76+v221)))
	v225 = v192 + int32(1)
	v227 = v225 << (uint(v220) % 32)
	v229 = *(*int32)(unsafe.Add(mBase, uint32(v180+v227)))
	if v223 == v229 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v221+v112)))
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v227+v182)))
	if v232 == v234 {
		v192 = v225
		goto L66
	} else {
		goto L72
	}
L70:
	;
	goto L71
L71:
	;
	goto L67
L72:
	;
	goto L71
L73:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L8
	} else {
		goto L74
	}
L74:
	;
	F_errmsg(m, int32(121131), int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L8
	} else {
		goto L75
	}
L75:
	;
	F_errdetail(m, int32(644823), int32(0))
	mBase = m.M
	v250 = m.ExcPending
	if v250 != 0 {
		goto L8
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(518919), int32(478), int32(120153))
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L8
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	v258 = F_palloc(m, v97)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L8
	} else {
		goto L79
	}
L79:
	;
	if v97 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	if v97 != 0 {
		goto L85
	} else {
		goto L86
	}
L81:
	;
	v260 = F__emscripten_memcpy_bulkmem(m, v256, v76, v97)
	mBase = m.M
	v261 = v260
	goto L83
L82:
	;
	v261 = v256
	goto L83
L83:
	;
	goto L80
L84:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v261)))
	*(*int32)(unsafe.Add(mBase, uint32(v261))) = v264 + int32(1)
	v268 = int32(0)
	if v268 < v60 {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v262 = F__emscripten_memcpy_bulkmem(m, v258, v112, v97)
	mBase = m.M
	v263 = v262
	goto L87
L86:
	;
	v263 = v258
	goto L87
L87:
	;
	goto L84
L88:
	;
	v272 = v60
	goto L90
L89:
	;
	v272 = v268
	goto L90
L90:
	;
	v273 = v268
	goto L91
L91:
	;
	if v273 == v272 {
		v415 = v256
		v416 = v258
		goto L17
	} else {
		goto L93
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L8
	} else {
		goto L98
	}
L93:
	;
	v301 = int32(2)
	v302 = v273 << (uint(v301) % 32)
	v304 = *(*int32)(unsafe.Add(mBase, uint32(v80+v302)))
	v306 = v273 + int32(1)
	v308 = v306 << (uint(v301) % 32)
	v310 = *(*int32)(unsafe.Add(mBase, uint32(v261+v308)))
	if v304 == v310 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v302+v111)))
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v308+v263)))
	if v313 == v315 {
		v273 = v306
		goto L91
	} else {
		goto L97
	}
L95:
	;
	goto L96
L96:
	;
	goto L92
L97:
	;
	goto L96
L98:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v323 = m.ExcPending
	if v323 != 0 {
		goto L8
	} else {
		goto L99
	}
L99:
	;
	F_errmsg(m, int32(121131), int32(0))
	mBase = m.M
	v327 = m.ExcPending
	if v327 != 0 {
		goto L8
	} else {
		goto L100
	}
L100:
	;
	F_errdetail(m, int32(644823), int32(0))
	mBase = m.M
	v331 = m.ExcPending
	if v331 != 0 {
		goto L8
	} else {
		goto L101
	}
L101:
	;
	F_errfinish(m, int32(518919), int32(506), int32(120153))
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L8
	} else {
		goto L102
	}
L102:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L103:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L8
	} else {
		goto L104
	}
L104:
	;
	F_errmsg(m, int32(121131), int32(0))
	mBase = m.M
	v347 = m.ExcPending
	if v347 != 0 {
		goto L8
	} else {
		goto L105
	}
L105:
	;
	v348 = F_format_type_be(m, v57)
	mBase = m.M
	v349 = m.ExcPending
	if v349 != 0 {
		goto L8
	} else {
		goto L106
	}
L106:
	;
	v350 = F_format_type_be(m, v58)
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L8
	} else {
		goto L107
	}
L107:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+20)) = v350
	*(*int32)(unsafe.Add(mBase, uint32(v30)+16)) = v348
	F_errdetail(m, int32(644963), v30+int32(16))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L8
	} else {
		goto L108
	}
L108:
	;
	F_errfinish(m, int32(518919), int32(375), int32(120153))
	mBase = m.M
	v363 = m.ExcPending
	if v363 != 0 {
		goto L8
	} else {
		goto L109
	}
L109:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L110:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v370 = m.ExcPending
	if v370 != 0 {
		goto L8
	} else {
		goto L111
	}
L111:
	;
	F_errmsg(m, int32(121131), int32(0))
	mBase = m.M
	v374 = m.ExcPending
	if v374 != 0 {
		goto L8
	} else {
		goto L112
	}
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v30)+4)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v30))) = v61
	F_errdetail(m, int32(644894), v30)
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L8
	} else {
		goto L113
	}
L113:
	;
	F_errfinish(m, int32(518919), int32(413), int32(120153))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L8
	} else {
		goto L114
	}
L114:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L115:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v392 = m.ExcPending
	if v392 != 0 {
		goto L8
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(121131), int32(0))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L8
	} else {
		goto L117
	}
L117:
	;
	F_errdetail(m, int32(644744), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L8
	} else {
		goto L118
	}
L118:
	;
	F_errfinish(m, int32(518919), int32(449), int32(120153))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L8
	} else {
		goto L119
	}
L119:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L120:
	;
	v473 = v74
	goto L122
L121:
	;
	v473 = (v468 + v463) & v465
	goto L122
L122:
	;
	v474 = F_ArrayGetNItems(m, v460, v442)
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L8
	} else {
		goto L123
	}
L123:
	;
	F_ArrayCheckBounds(m, v460, v442, v443)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L8
	} else {
		goto L124
	}
L124:
	;
	v478 = int32(2)
	v480 = int32(base.Ui32(v110)>>(uint(v478)%32)) - v108
	v483 = int32(base.Ui32(v109)>>(uint(v478)%32)) - v93
	v485 = *(*int32)(unsafe.Add(mBase, uint32(v52)+8))
	if v485 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	if v73 != 0 {
		goto L131
	} else {
		goto L132
	}
L126:
	;
	v509 = (v460<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v510 = int32(0)
	goto L125
L127:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
	if v488 == int32(0) {
		goto L126
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v494 = base.I32_div_s(v474+int32(7), int32(8))
	v501 = (v494 + v460<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v509 = v501
	v510 = v501
	goto L125
L130:
	;
	goto L129
L131:
	;
	v511 = v73
	goto L133
L132:
	;
	v511 = (v462 + v463) & v465
	goto L133
L133:
	;
	v513 = v480 + v483 + v509
	v514 = F_palloc0(m, v513)
	mBase = m.M
	v515 = m.ExcPending
	if v515 != 0 {
		goto L8
	} else {
		goto L134
	}
L134:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v514)+12)) = v57
	*(*int32)(unsafe.Add(mBase, uint32(v514)+8)) = v510
	*(*int32)(unsafe.Add(mBase, uint32(v514)+4)) = v460
	v519 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v514))) = v513 << (uint(v519) % 32)
	v523 = v514 + int32(16)
	v525 = v460 << (uint(v519) % 32)
	if v525 != 0 {
		goto L136
	} else {
		goto L137
	}
L135:
	;
	if v525 != 0 {
		goto L140
	} else {
		goto L141
	}
L136:
	;
	v526 = F__emscripten_memcpy_bulkmem(m, v523, v442, v525)
	mBase = m.M
	v527 = v526
	goto L138
L137:
	;
	v527 = v523
	goto L138
L138:
	;
	goto L135
L139:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	if v532 != 0 {
		goto L143
	} else {
		goto L144
	}
L140:
	;
	v529 = F__emscripten_memcpy_bulkmem(m, v527+v525, v443, v525)
	mBase = m.M
	goto L142
L141:
	;
	goto L142
L142:
	;
	goto L139
L143:
	;
	v540 = v532
	goto L145
L144:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	v540 = (v533<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L145
L145:
	;
	if v483 != 0 {
		goto L147
	} else {
		goto L148
	}
L146:
	;
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	if v544 != 0 {
		goto L150
	} else {
		goto L151
	}
L147:
	;
	v542 = F__emscripten_memcpy_bulkmem(m, v540+v514, v52+v473, v483)
	mBase = m.M
	goto L149
L148:
	;
	goto L149
L149:
	;
	goto L146
L150:
	;
	v552 = v544
	goto L152
L151:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	v552 = (v545<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L152
L152:
	;
	if v480 != 0 {
		goto L154
	} else {
		goto L155
	}
L153:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	if v557 == int32(0) {
		v833 = v514
		goto L1
	} else {
		goto L157
	}
L154:
	;
	v555 = F__emscripten_memcpy_bulkmem(m, v552+v514+v483, v511+v55, v480)
	mBase = m.M
	goto L156
L155:
	;
	goto L156
L156:
	;
	goto L153
L157:
	;
	if v73 != 0 {
		goto L158
	} else {
		goto L159
	}
L158:
	;
	v562 = v80 + v462
	goto L160
L159:
	;
	v562 = int32(0)
	goto L160
L160:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	v567 = int32(0)
	if v74 != 0 {
		goto L161
	} else {
		goto L162
	}
L161:
	;
	v570 = v76 + v468
	goto L163
L162:
	;
	v570 = v567
	goto L163
L163:
	;
	v571 = int32(0)
	if v77 <= v571 {
		goto L165
	} else {
		goto L166
	}
L164:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v514)+8))
	if v699 != 0 {
		goto L195
	} else {
		goto L196
	}
L165:
	;
	goto L164
L166:
	;
	v578 = int32(1)
	v583 = base.I32_div_s(v567, int32(8))
	v584 = v527 + v563<<(uint(int32(3))%32) + v583
	v585 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v584))))
	if v570 == int32(0) {
		goto L168
	} else {
		goto L169
	}
L167:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v685))) = uint8(v680)
	goto L165
L168:
	;
	v589 = v585
	v592 = v77
	v593 = v578
	v594 = v584
	goto L171
L169:
	;
	goto L170
L170:
	;
	v623 = base.I32_div_s(v571, int32(8))
	v624 = v570 + v623
	v625 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v624))))
	v626 = int32(1)
	v627 = v585
	v630 = v77
	v631 = v578
	v632 = v584
	v633 = v624
	v634 = v625
	goto L179
L171:
	;
	v597 = v589 | v593
	v598 = int32(1)
	v599 = v592 - v598
	v601 = v593 << (uint(v598) % 32)
	if v601 == int32(256) {
		goto L173
	} else {
		goto L174
	}
L172:
	;
	if v612 != int32(1) {
		v680 = v611
		v685 = v613
		goto L167
	} else {
		goto L178
	}
L173:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v594))) = uint8(v597)
	if v599 == int32(0) {
		goto L165
	} else {
		goto L176
	}
L174:
	;
	v611 = v597
	v612 = v601
	v613 = v594
	goto L175
L175:
	;
	if base.Ui32(int32(1)) < base.Ui32(v592) {
		v589 = v611
		v592 = v599
		v593 = v612
		v594 = v613
		goto L171
	} else {
		goto L177
	}
L176:
	;
	v607 = int32(1)
	v609 = v594 + v607
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v609))))
	v611 = v610
	v612 = v607
	v613 = v609
	goto L175
L177:
	;
	goto L172
L178:
	;
	goto L165
L179:
	;
	if v626&v634 != 0 {
		goto L181
	} else {
		goto L182
	}
L180:
	;
	if v655 == int32(1) {
		goto L165
	} else {
		goto L194
	}
L181:
	;
	v640 = v627 | v631
	goto L183
L182:
	;
	v640 = v627 & (v631 ^ int32(-1))
	goto L183
L183:
	;
	v641 = int32(1)
	v642 = v630 - v641
	v644 = v631 << (uint(v641) % 32)
	if v644 == int32(256) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v632))) = uint8(v640)
	if v642 == int32(0) {
		goto L165
	} else {
		goto L187
	}
L185:
	;
	v654 = v640
	v655 = v644
	v656 = v632
	goto L186
L186:
	;
	v658 = v626 << (uint(int32(1)) % 32)
	if v658 == int32(256) {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	v650 = int32(1)
	v652 = v632 + v650
	v653 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v652))))
	v654 = v653
	v655 = v650
	v656 = v652
	goto L186
L188:
	;
	goto L180
L189:
	;
	if v642 == int32(0) {
		goto L188
	} else {
		goto L192
	}
L190:
	;
	v667 = v658
	v668 = v633
	v669 = v634
	goto L191
L191:
	;
	if base.Ui32(int32(1)) < base.Ui32(v630) {
		v626 = v667
		v627 = v654
		v630 = v642
		v631 = v655
		v632 = v656
		v633 = v668
		v634 = v669
		goto L179
	} else {
		goto L193
	}
L192:
	;
	v663 = int32(1)
	v664 = v633 + v663
	v665 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v664))))
	v667 = v663
	v668 = v664
	v669 = v665
	goto L191
L193:
	;
	goto L188
L194:
	;
	v680 = v654
	v685 = v656
	goto L167
L195:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(v514)+4))
	v704 = v527 + v700<<(uint(int32(3))%32)
	goto L197
L196:
	;
	v704 = int32(0)
	goto L197
L197:
	;
	v705 = int32(0)
	if v81 <= v705 {
		goto L199
	} else {
		goto L200
	}
L198:
	;
	v833 = v514
	goto L1
L199:
	;
	goto L198
L200:
	;
	v715 = int32(1) << (uint(v77&int32(7)) % 32)
	v717 = base.I32_div_s(v77, int32(8))
	v718 = v704 + v717
	v719 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v718))))
	if v562 == int32(0) {
		goto L202
	} else {
		goto L203
	}
L201:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v819))) = uint8(v814)
	goto L199
L202:
	;
	v723 = v719
	v726 = v81
	v727 = v715
	v728 = v718
	goto L205
L203:
	;
	goto L204
L204:
	;
	v757 = base.I32_div_s(v705, int32(8))
	v758 = v562 + v757
	v759 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v758))))
	v760 = int32(1)
	v761 = v719
	v764 = v81
	v765 = v715
	v766 = v718
	v767 = v758
	v768 = v759
	goto L213
L205:
	;
	v731 = v723 | v727
	v732 = int32(1)
	v733 = v726 - v732
	v735 = v727 << (uint(v732) % 32)
	if v735 == int32(256) {
		goto L207
	} else {
		goto L208
	}
L206:
	;
	if v746 != int32(1) {
		v814 = v745
		v819 = v747
		goto L201
	} else {
		goto L212
	}
L207:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v728))) = uint8(v731)
	if v733 == int32(0) {
		goto L199
	} else {
		goto L210
	}
L208:
	;
	v745 = v731
	v746 = v735
	v747 = v728
	goto L209
L209:
	;
	if base.Ui32(int32(1)) < base.Ui32(v726) {
		v723 = v745
		v726 = v733
		v727 = v746
		v728 = v747
		goto L205
	} else {
		goto L211
	}
L210:
	;
	v741 = int32(1)
	v743 = v728 + v741
	v744 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
	v745 = v744
	v746 = v741
	v747 = v743
	goto L209
L211:
	;
	goto L206
L212:
	;
	goto L199
L213:
	;
	if v760&v768 != 0 {
		goto L215
	} else {
		goto L216
	}
L214:
	;
	if v789 == int32(1) {
		goto L199
	} else {
		goto L228
	}
L215:
	;
	v774 = v761 | v765
	goto L217
L216:
	;
	v774 = v761 & (v765 ^ int32(-1))
	goto L217
L217:
	;
	v775 = int32(1)
	v776 = v764 - v775
	v778 = v765 << (uint(v775) % 32)
	if v778 == int32(256) {
		goto L218
	} else {
		goto L219
	}
L218:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v766))) = uint8(v774)
	if v776 == int32(0) {
		goto L199
	} else {
		goto L221
	}
L219:
	;
	v788 = v774
	v789 = v778
	v790 = v766
	goto L220
L220:
	;
	v792 = v760 << (uint(int32(1)) % 32)
	if v792 == int32(256) {
		goto L223
	} else {
		goto L224
	}
L221:
	;
	v784 = int32(1)
	v786 = v766 + v784
	v787 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v786))))
	v788 = v787
	v789 = v784
	v790 = v786
	goto L220
L222:
	;
	goto L214
L223:
	;
	if v776 == int32(0) {
		goto L222
	} else {
		goto L226
	}
L224:
	;
	v801 = v792
	v802 = v767
	v803 = v768
	goto L225
L225:
	;
	if base.Ui32(int32(1)) < base.Ui32(v764) {
		v760 = v801
		v761 = v788
		v764 = v776
		v765 = v789
		v766 = v790
		v767 = v802
		v768 = v803
		goto L213
	} else {
		goto L227
	}
L226:
	;
	v797 = int32(1)
	v798 = v767 + v797
	v799 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v798))))
	v801 = v797
	v802 = v798
	v803 = v799
	goto L225
L227:
	;
	goto L222
L228:
	;
	v814 = v788
	v819 = v790
	goto L201
}
func F_array_eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
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
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v217 int32
	_ = v217
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v259 int32
	_ = v259
	var v271 int32
	_ = v271
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v294 int32
	_ = v294
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v309 int32
	_ = v309
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v322 int32
	_ = v322
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v357 int32
	_ = v357
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v413 int32
	_ = v413
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v443 int32
	_ = v443
	var v450 int32
	_ = v450
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v501 int32
	_ = v501
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(96)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_DatumGetAnyArrayP(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v24 = F_DatumGetAnyArrayP(m, v23)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v27 == int32(-1) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	if v26 == int32(-1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v19)+32))
	v33 = v30
	goto L4
L6:
	;
	goto L7
L7:
	;
	v33 = v19 + int32(16)
	goto L4
L8:
	;
	if v27 == int32(-1) {
		goto L13
	} else {
		goto L14
	}
L9:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	v39 = v36
	goto L8
L10:
	;
	goto L11
L11:
	;
	v39 = v24 + int32(16)
	goto L8
L12:
	;
	if v27 == int32(-1) {
		goto L17
	} else {
		goto L18
	}
L13:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
	v49 = v42
	goto L12
L14:
	;
	goto L15
L15:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v49 = v19 + v43<<(uint(int32(2))%32) + int32(16)
	goto L12
L16:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v487 = m.ExcPending
	if v487 != 0 {
		goto L1
	} else {
		goto L136
	}
L17:
	;
	v54 = int32(40)
	goto L19
L18:
	;
	v54 = int32(12)
	goto L19
L19:
	;
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v19+v54)))
	if v26 == int32(-1) {
		goto L21
	} else {
		goto L22
	}
L20:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v69+v24)))
	if v56 == v71 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v24)+36))
	v68 = v59
	v69 = int32(40)
	goto L20
L22:
	;
	goto L23
L23:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v68 = v24 + v61<<(uint(int32(2))%32) + int32(16)
	v69 = int32(12)
	goto L20
L24:
	;
	if v27 == int32(-1) {
		goto L28
	} else {
		goto L29
	}
L25:
	;
	goto L26
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v471 = m.ExcPending
	if v471 != 0 {
		goto L1
	} else {
		goto L132
	}
L27:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v450 == int32(-1) {
		goto L124
	} else {
		goto L125
	}
L28:
	;
	v77 = int32(28)
	goto L30
L29:
	;
	v77 = int32(4)
	goto L30
L30:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v19+v77)))
	if v26 == int32(-1) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v84 = int32(28)
	goto L33
L32:
	;
	v84 = int32(4)
	goto L33
L33:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v24+v84)))
	if v79 != v86 {
		v443 = v2
		goto L27
	} else {
		goto L34
	}
L34:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v90 = v79 << (uint(int32(2)) % 32)
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	if v152 != 0 {
		v443 = v2
		goto L27
	} else {
		goto L53
	}
L36:
	;
	v152 = int32(0)
	goto L35
L37:
	;
	v126 = v121
	v127 = v122
	v128 = v123
	goto L47
L38:
	;
	if (v33|v39)&int32(3) != 0 {
		v121 = v33
		v122 = v39
		v123 = v90
		goto L37
	} else {
		goto L41
	}
L39:
	;
	v114 = v33
	v115 = v39
	v116 = v90
	goto L40
L40:
	;
	if v116 == int32(0) {
		goto L36
	} else {
		goto L46
	}
L41:
	;
	v98 = v33
	v99 = v39
	v100 = v90
	goto L42
L42:
	;
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v98)))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v99)))
	if v103 != v104 {
		v121 = v98
		v122 = v99
		v123 = v100
		goto L37
	} else {
		goto L44
	}
L43:
	;
	v114 = v109
	v115 = v107
	v116 = v111
	goto L40
L44:
	;
	v106 = int32(4)
	v107 = v99 + v106
	v109 = v98 + v106
	v111 = v100 - v106
	if base.Ui32(int32(3)) < base.Ui32(v111) {
		v98 = v109
		v99 = v107
		v100 = v111
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v121 = v114
	v122 = v115
	v123 = v116
	goto L37
L47:
	;
	v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v126))))
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v127))))
	if v131 == v132 {
		goto L49
	} else {
		goto L50
	}
L48:
	;
	v152 = v131 - v132
	goto L35
L49:
	;
	v134 = int32(1)
	v139 = v128 - v134
	if v139 != 0 {
		v126 = v126 + v134
		v127 = v127 + v134
		v128 = v139
		goto L47
	} else {
		goto L52
	}
L50:
	;
	goto L51
L51:
	;
	goto L48
L52:
	;
	goto L36
L53:
	;
	if base.Ui32(int32(4)) <= base.Ui32(v90) {
		goto L57
	} else {
		goto L58
	}
L54:
	;
	if v214 != 0 {
		v443 = v2
		goto L27
	} else {
		goto L72
	}
L55:
	;
	v214 = int32(0)
	goto L54
L56:
	;
	v188 = v183
	v189 = v184
	v190 = v185
	goto L66
L57:
	;
	if (v49|v68)&int32(3) != 0 {
		v183 = v49
		v184 = v68
		v185 = v90
		goto L56
	} else {
		goto L60
	}
L58:
	;
	v176 = v49
	v177 = v68
	v178 = v90
	goto L59
L59:
	;
	if v178 == int32(0) {
		goto L55
	} else {
		goto L65
	}
L60:
	;
	v160 = v49
	v161 = v68
	v162 = v90
	goto L61
L61:
	;
	v165 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
	if v165 != v166 {
		v183 = v160
		v184 = v161
		v185 = v162
		goto L56
	} else {
		goto L63
	}
L62:
	;
	v176 = v171
	v177 = v169
	v178 = v173
	goto L59
L63:
	;
	v168 = int32(4)
	v169 = v161 + v168
	v171 = v160 + v168
	v173 = v162 - v168
	if base.Ui32(int32(3)) < base.Ui32(v173) {
		v160 = v171
		v161 = v169
		v162 = v173
		goto L61
	} else {
		goto L64
	}
L64:
	;
	goto L62
L65:
	;
	v183 = v176
	v184 = v177
	v185 = v178
	goto L56
L66:
	;
	v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v188))))
	v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v189))))
	if v193 == v194 {
		goto L68
	} else {
		goto L69
	}
L67:
	;
	v214 = v193 - v194
	goto L54
L68:
	;
	v196 = int32(1)
	v201 = v190 - v196
	if v201 != 0 {
		v188 = v188 + v196
		v189 = v189 + v196
		v190 = v201
		goto L66
	} else {
		goto L71
	}
L69:
	;
	goto L70
L70:
	;
	goto L67
L71:
	;
	goto L55
L72:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v216 = *(*int32)(unsafe.Add(mBase, uint32(v215)+16))
	if v216 != 0 {
		goto L74
	} else {
		goto L75
	}
L73:
	;
	v228 = int32(*(*int8)(unsafe.Add(mBase, uint32(v227)+11)))
	v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v227)+10)))
	v230 = int32(*(*int16)(unsafe.Add(mBase, uint32(v227)+8)))
	v231 = int32(2)
	*(*uint16)(unsafe.Add(mBase, uint32(v16)+78)) = uint16(v231)
	v233 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v233)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v88
	*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v227 + int32(76)
	v241 = F_ArrayGetNItems(m, v79, v33)
	mBase = m.M
	v242 = m.ExcPending
	if v242 != 0 {
		goto L1
	} else {
		goto L80
	}
L74:
	;
	v217 = *(*int32)(unsafe.Add(mBase, uint32(v216)))
	if v217 == v56 {
		v227 = v216
		goto L73
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v220 = F_lookup_type_cache(m, v56, int32(32))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L78
	}
L77:
	;
	goto L76
L78:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v220)+80))
	if v222 == int32(0) {
		goto L16
	} else {
		goto L79
	}
L79:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v225)+16)) = v220
	v227 = v220
	goto L73
L80:
	;
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	if v243 == int32(-1) {
		goto L82
	} else {
		goto L83
	}
L81:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v302
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v306 == int32(-1) {
		goto L95
	} else {
		goto L96
	}
L82:
	;
	v246 = *(*int32)(unsafe.Add(mBase, uint32(v19)+48))
	if v246 != 0 {
		goto L85
	} else {
		goto L86
	}
L83:
	;
	goto L84
L84:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = int64(0)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v19)+8))
	if v279 == int32(0) {
		goto L91
	} else {
		goto L92
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+40)) = v246
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v19)+52))
	v249 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v249
	*(*int32)(unsafe.Add(mBase, uint32(v16)+44)) = v248
	v302 = v249
	goto L81
L86:
	;
	goto L87
L87:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = int64(0)
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v19)+68))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	if v256 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L88:
	;
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v255 + (v259<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v302 = int32(0)
	goto L81
L89:
	;
	goto L90
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v255 + v256
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v302 = v255 + v271<<(uint(int32(3))%32) + int32(16)
	goto L81
L91:
	;
	v282 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v19 + (v282<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v302 = int32(0)
	goto L81
L92:
	;
	goto L93
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v279 + v19
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
	v302 = v19 + v294<<(uint(int32(3))%32) + int32(16)
	goto L81
L94:
	;
	v366 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v366
	*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v365
	if v241 <= int32(0) {
		v443 = v366
		goto L27
	} else {
		goto L107
	}
L95:
	;
	v309 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	if v309 != 0 {
		goto L98
	} else {
		goto L99
	}
L96:
	;
	goto L97
L97:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+20)) = int64(0)
	v342 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
	if v342 == int32(0) {
		goto L104
	} else {
		goto L105
	}
L98:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v309
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v24)+52))
	v312 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v312
	*(*int32)(unsafe.Add(mBase, uint32(v16)+24)) = v311
	v365 = v312
	goto L94
L99:
	;
	goto L100
L100:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v16)+20)) = int64(0)
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v24)+68))
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v318)+8))
	if v319 == int32(0) {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v322 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v318 + (v322<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v365 = int32(0)
	goto L94
L102:
	;
	goto L103
L103:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v318 + v319
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v318)+4))
	v365 = v318 + v334<<(uint(int32(3))%32) + int32(16)
	goto L94
L104:
	;
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v24 + (v345<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v365 = int32(0)
	goto L94
L105:
	;
	goto L106
L106:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16)+28)) = v342 + v24
	v357 = *(*int32)(unsafe.Add(mBase, uint32(v24)+4))
	v365 = v24 + v357<<(uint(int32(3))%32) + int32(16)
	goto L94
L107:
	;
	v374 = v229 & int32(1)
	v377 = int32(0)
	goto L108
L108:
	;
	v392 = F_array_iter_next(m, v16+int32(40), v16+int32(19), v377, v230, v374, v228)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L1
	} else {
		goto L110
	}
L109:
	;
	v443 = v433
	goto L27
L110:
	;
	v398 = F_array_iter_next(m, v16+int32(20), v16+int32(18), v377, v230, v374, v228)
	mBase = m.M
	v399 = m.ExcPending
	if v399 != 0 {
		goto L1
	} else {
		goto L111
	}
L111:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+18)))
	v401 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+19)))
	if v401 == int32(1) {
		goto L115
	} else {
		goto L116
	}
L112:
	;
	v433 = int32(1)
	v435 = v377 + v433
	if v435 != v241 {
		v377 = v435
		goto L108
	} else {
		goto L123
	}
L113:
	;
	v413 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+92)) = uint8(v413)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v398
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+84)) = uint8(v413)
	*(*int32)(unsafe.Add(mBase, uint32(v16)+80)) = v392
	*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)) = uint8(v413)
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v16)+60))
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v424)))
	v426 = m.T0[v425].(func(*base.Module, int32) int32)(m, v16+int32(60))
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L1
	} else {
		goto L120
	}
L114:
	;
	v443 = int32(0)
	goto L27
L115:
	;
	if v400&int32(1) == int32(0) {
		goto L114
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	if v400&int32(1) == int32(0) {
		goto L113
	} else {
		goto L119
	}
L118:
	;
	goto L112
L119:
	;
	goto L114
L120:
	;
	v428 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+76)))
	if v428 != 0 {
		v443 = v413
		goto L27
	} else {
		goto L121
	}
L121:
	;
	if v426 == int32(0) {
		v443 = v413
		goto L27
	} else {
		goto L122
	}
L122:
	;
	goto L112
L123:
	;
	goto L109
L124:
	;
	v457 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
	if v457 == int32(-1) {
		goto L128
	} else {
		goto L129
	}
L125:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v19 == v453 {
		goto L124
	} else {
		goto L126
	}
L126:
	;
	F_pfree(m, v19)
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L1
	} else {
		goto L127
	}
L127:
	;
	goto L124
L128:
	;
	m.G0 = v16 + int32(96)
	return v443
L129:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v24 == v460 {
		goto L128
	} else {
		goto L130
	}
L130:
	;
	F_pfree(m, v24)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L1
	} else {
		goto L131
	}
L131:
	;
	goto L128
L132:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L1
	} else {
		goto L133
	}
L133:
	;
	F_errmsg(m, int32(172383), int32(0))
	mBase = m.M
	v478 = m.ExcPending
	if v478 != 0 {
		goto L1
	} else {
		goto L134
	}
L134:
	;
	F_errfinish(m, int32(518800), int32(3846), int32(242896))
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L1
	} else {
		goto L135
	}
L135:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L136:
	;
	F_errcode(m, int32(52461700))
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L1
	} else {
		goto L137
	}
L137:
	;
	v491 = F_format_type_be(m, v56)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L1
	} else {
		goto L138
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16))) = v491
	F_errmsg(m, int32(200067), v16)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L1
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(518800), int32(3871), int32(242896))
	mBase = m.M
	v501 = m.ExcPending
	if v501 != 0 {
		goto L1
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_ge(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2^int32(-1)) >> (uint(int32(31)) % 32))
	}
}
func F_array_lt(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_array_cmp(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return int32(base.Ui32(v2) >> (uint(int32(31)) % 32))
	}
}
func F_array_out(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
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
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
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
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v218 int32
	_ = v218
	var v226 int32
	_ = v226
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v260 int32
	_ = v260
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v331 int32
	_ = v331
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v343 int32
	_ = v343
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v369 int32
	_ = v369
	var v373 int32
	_ = v373
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v445 int32
	_ = v445
	var v450 int32
	_ = v450
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v486 int32
	_ = v486
	var v490 int32
	_ = v490
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v496 int32
	_ = v496
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v538 int32
	_ = v538
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v553 int32
	_ = v553
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v624 int32
	_ = v624
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v646 int32
	_ = v646
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v673 int32
	_ = v673
	var v679 int32
	_ = v679
	var v683 int32
	_ = v683
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v690 int32
	_ = v690
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v711 int32
	_ = v711
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v748 int32
	_ = v748
	var v760 int32
	_ = v760
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v816 int32
	_ = v816
	var v817 int32
	_ = v817
	var v818 int32
	_ = v818
	var v821 int32
	_ = v821
	var v837 int32
	_ = v837
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v891 int32
	_ = v891
	var v892 int32
	_ = v892
	var v893 int32
	_ = v893
	var v896 int32
	_ = v896
	var v898 int32
	_ = v898
	var v919 int32
	_ = v919
	var v922 int32
	_ = v922
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v950 int32
	_ = v950
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v965 int32
	_ = v965
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v998 int32
	_ = v998
	var v1000 int32
	_ = v1000
	var v1003 int32
	_ = v1003
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1014 int32
	_ = v1014
	var v1016 int32
	_ = v1016
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1027 int32
	_ = v1027
	var v1034 int32
	_ = v1034
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1047 int32
	_ = v1047
	var v1069 int32
	_ = v1069
	var v1073 int32
	_ = v1073
	var v1075 int32
	_ = v1075
	var v1078 int32
	_ = v1078
	var v1080 int32
	_ = v1080
	var v1099 int32
	_ = v1099
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1119 int32
	_ = v1119
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1170 int32
	_ = v1170
	var v1172 int32
	_ = v1172
	var v1183 int32
	_ = v1183
	v2 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(272)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_DatumGetAnyArrayP(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v32 == int32(-1) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v35 = int32(40)
	goto L5
L4:
	;
	v35 = int32(12)
	goto L5
L5:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26+v35)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+16))
	if v39 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v82 == int32(-1) {
		goto L15
	} else {
		goto L16
	}
L7:
	;
	F_get_type_io_data(m, v37, int32(1), v55+int32(4), v55+int32(6), v55+int32(7), v55+int32(8), v55+int32(12), v55+int32(16))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L1
	} else {
		goto L13
	}
L8:
	;
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v38)+20))
	v44 = F_MemoryContextAlloc(m, v42, int32(48))
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L1
	} else {
		goto L11
	}
L9:
	;
	goto L10
L10:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	if v53 == v37 {
		v79 = v39
		goto L6
	} else {
		goto L12
	}
L11:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v44
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v48)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v49))) = v37 ^ int32(-1)
	v55 = v49
	goto L7
L12:
	;
	v55 = v39
	goto L7
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v55)+16))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)+20))
	F_fmgr_info_cxt(m, v71, v55+int32(20), v75)
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L14
	}
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v55))) = v37
	v79 = v55
	goto L6
L15:
	;
	v85 = int32(28)
	goto L17
L16:
	;
	v85 = int32(4)
	goto L17
L17:
	;
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v26+v85)))
	if v82 == int32(-1) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v100 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+8)))
	v101 = int32(*(*int8)(unsafe.Add(mBase, uint32(v79)+7)))
	v102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v79)+6)))
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v79)+4)))
	v104 = F_ArrayGetNItems(m, v87, v98)
	mBase = m.M
	v105 = m.ExcPending
	if v105 != 0 {
		goto L1
	} else {
		goto L25
	}
L19:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v26)+36))
	v98 = v90
	v99 = v91
	goto L18
L20:
	;
	goto L21
L21:
	;
	v93 = v26 + int32(16)
	v94 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v98 = v93
	v99 = v93 + v94<<(uint(int32(2))%32)
	goto L18
L22:
	;
	m.G0 = v23 + int32(272)
	return v1183
L23:
	;
	v163 = F_palloc(m, v104<<(uint(int32(2))%32))
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L35
	}
L24:
	;
	v112 = v106
	goto L31
L25:
	;
	if v104 != 0 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v106 = int32(0)
	if v106 < v87 {
		goto L24
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v110 = F_pstrdup(m, int32(4103))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L30
	}
L29:
	;
	v158 = v2
	goto L23
L30:
	;
	v1183 = v110
	goto L22
L31:
	;
	v135 = *(*int32)(unsafe.Add(mBase, uint32(v99+v112<<(uint(int32(2))%32))))
	v137 = base.B2i32(v135 != int32(1))
	if v135 != int32(1) {
		v158 = v137
		goto L23
	} else {
		goto L33
	}
L32:
	;
	v158 = v137
	goto L23
L33:
	;
	v139 = v112 + int32(1)
	if v139 != v87 {
		v112 = v139
		goto L31
	} else {
		goto L34
	}
L34:
	;
	goto L32
L35:
	;
	v165 = F_palloc(m, v104)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if v167 == int32(-1) {
		goto L38
	} else {
		goto L39
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+24)) = v226
	*(*int32)(unsafe.Add(mBase, uint32(v23)+28)) = int32(1)
	if v104 <= int32(0) {
		goto L51
	} else {
		goto L52
	}
L38:
	;
	v170 = *(*int32)(unsafe.Add(mBase, uint32(v26)+48))
	if v170 != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+12)) = int64(0)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v203 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+12)) = v170
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v26)+52))
	v173 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v173
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v172
	v226 = v173
	goto L37
L42:
	;
	goto L43
L43:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v23)+12)) = int64(0)
	v179 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v179)+8))
	if v180 == int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v183 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v179 + (v183<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v226 = int32(0)
	goto L37
L45:
	;
	goto L46
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v179 + v180
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
	v226 = v179 + v195<<(uint(int32(3))%32) + int32(16)
	goto L37
L47:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v26 + (v206<<(uint(int32(3))%32)+int32(23))&int32(-8)
	v226 = int32(0)
	goto L37
L48:
	;
	goto L49
L49:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v203 + v26
	v218 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v226 = v26 + v218<<(uint(int32(3))%32) + int32(16)
	goto L37
L50:
	;
	if int32(0) < v87 {
		goto L99
	} else {
		goto L100
	}
L51:
	;
	v424 = int32(0)
	goto L50
L52:
	;
	goto L53
L53:
	;
	v240 = int32(0)
	v243 = v2
	goto L54
L54:
	;
	v260 = v163 + v243<<(uint(int32(2))%32)
	v265 = F_array_iter_next(m, v23+int32(12), v23-int32(-64), v243, v103, v102&int32(1), v101)
	mBase = m.M
	v266 = m.ExcPending
	if v266 != 0 {
		goto L1
	} else {
		goto L56
	}
L55:
	;
	v424 = v418
	goto L50
L56:
	;
	v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v23)+64)))
	if v267 == int32(1) {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	v417 = int32(1)
	v418 = v416 + v417
	v420 = v243 + v417
	if v420 != v104 {
		v240 = v418
		v243 = v420
		goto L54
	} else {
		goto L95
	}
L58:
	;
	v271 = F_pstrdup(m, int32(559295))
	mBase = m.M
	v272 = m.ExcPending
	if v272 != 0 {
		goto L1
	} else {
		goto L61
	}
L59:
	;
	goto L60
L60:
	;
	v279 = F_OutputFunctionCall(m, v79+int32(20), v265)
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L62
	}
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v271
	v275 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v243+v165))) = uint8(v275)
	v416 = v240 + int32(4)
	goto L57
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v260))) = v279
	v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v279))))
	if v282 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L63:
	;
	v286 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v243+v165))) = uint8(v286)
	v416 = v240 + int32(2)
	goto L57
L64:
	;
	goto L65
L65:
	;
	v293 = v279
	v294 = int32(559295)
	goto L67
L66:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v260)))
	v335 = v334
	v337 = v240
	v343 = base.B2i32(v331 == int32(0))
	goto L79
L67:
	;
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v293))))
	v298 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v294))))
	if v297 == v298 {
		v320 = v297
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v331 = int32(0)
	goto L66
L69:
	;
	v322 = int32(1)
	if v320 != 0 {
		v293 = v293 + v322
		v294 = v294 + v322
		goto L67
	} else {
		goto L78
	}
L70:
	;
	if base.Ui32((v297-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v308 = v297 | int32(32)
	goto L73
L72:
	;
	v308 = v297
	goto L73
L73:
	;
	if base.Ui32((v298-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v317 = v298 | int32(32)
	goto L76
L75:
	;
	v317 = v298
	goto L76
L76:
	;
	if v308 == v317 {
		v320 = v308
		goto L69
	} else {
		goto L77
	}
L77:
	;
	v331 = v308 - v317
	goto L66
L78:
	;
	goto L68
L79:
	;
	v355 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v335))))
	if v355 != 0 {
		goto L81
	} else {
		goto L82
	}
L80:
	;
	v391 = v343 & int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v243+v165))) = uint8(v391)
	if v391 != 0 {
		goto L92
	} else {
		goto L93
	}
L81:
	;
	v359 = v355 & int32(255)
	switch v359 - int32(123) {
	case 0, 2:
		goto L85
	case 1:
		goto L86
	default:
		goto L87
	}
L82:
	;
	goto L83
L83:
	;
	goto L80
L84:
	;
	v335 = v335 + int32(1)
	v337 = v337 + int32(1)
	v343 = v386
	goto L79
L85:
	;
	v386 = int32(1)
	goto L84
L86:
	;
	if v359 == v100 {
		goto L85
	} else {
		goto L89
	}
L87:
	;
	if base.B2i32(v359 != int32(92))&base.B2i32(v359 != int32(34)) != 0 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v369 = int32(1)
	v335 = v335 + v369
	v337 = v337 + int32(2)
	v343 = v369
	goto L79
L89:
	;
	v373 = base.I32_extend8_s(v355)
	goto L90
L90:
	;
	if base.B2i32(v373 == int32(32))|base.B2i32(base.Ui32((v373-int32(9))&int32(255)) < base.Ui32(int32(5))) == int32(0) {
		v386 = v343
		goto L84
	} else {
		goto L91
	}
L91:
	;
	goto L85
L92:
	;
	v395 = v337 + int32(2)
	goto L94
L93:
	;
	v395 = v337
	goto L94
L94:
	;
	v416 = v395
	goto L57
L95:
	;
	goto L55
L96:
	;
	v771 = int32(123)
	*(*uint16)(unsafe.Add(mBase, uint32(v770))) = uint16(v771)
	if int32(0) < v87 {
		goto L144
	} else {
		goto L145
	}
L97:
	;
	v664 = int32(61)
	*(*uint16)(unsafe.Add(mBase, uint32(v644))) = uint16(v664)
	v670 = F_palloc(m, v644+(v646-v23)-int32(63))
	mBase = m.M
	v671 = m.ExcPending
	if v671 != 0 {
		goto L1
	} else {
		goto L122
	}
L98:
	;
	v642 = F_palloc(m, v624)
	mBase = m.M
	v643 = m.ExcPending
	if v643 != 0 {
		goto L1
	} else {
		goto L121
	}
L99:
	;
	v445 = v87 & int32(3)
	if base.Ui32(v87) < base.Ui32(int32(4)) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	goto L101
L101:
	;
	v618 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+64)) = uint8(v618)
	if v158 != 0 {
		v644 = v23 - int32(-64)
		v646 = v424
		goto L97
	} else {
		goto L120
	}
L102:
	;
	if v445 != 0 {
		goto L109
	} else {
		goto L110
	}
L103:
	;
	v450 = int32(0)
	v498 = int32(1)
	v499 = v450
	v502 = v450
	goto L102
L104:
	;
	goto L105
L105:
	;
	v454 = int32(0)
	v458 = int32(1)
	v459 = v454
	v462 = v454
	v463 = v454
	goto L106
L106:
	;
	v480 = v98 + v459<<(uint(int32(2))%32)
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+8))
	v482 = *(*int32)(unsafe.Add(mBase, uint32(v480)))
	v483 = v482 * v458
	v484 = *(*int32)(unsafe.Add(mBase, uint32(v480)+4))
	v485 = v483 * v484
	v486 = v481 * v485
	v490 = v486 + (v485 + (v483 + (v458 + v462)))
	v491 = *(*int32)(unsafe.Add(mBase, uint32(v480)+12))
	v492 = v491 * v486
	v493 = int32(4)
	v494 = v459 + v493
	v496 = v463 + v493
	if v496 != v87&int32(2147483644) {
		v458 = v492
		v459 = v494
		v462 = v490
		v463 = v496
		goto L106
	} else {
		goto L108
	}
L107:
	;
	v498 = v492
	v499 = v494
	v502 = v490
	goto L102
L108:
	;
	goto L107
L109:
	;
	v518 = v498
	v519 = v499
	v522 = v502
	v526 = int32(0)
	goto L112
L110:
	;
	v553 = v502
	goto L111
L111:
	;
	v569 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+64)) = uint8(v569)
	v574 = v553<<(uint(int32(1))%32) + v424
	if v158 == v569 {
		v624 = v574
		goto L98
	} else {
		goto L115
	}
L112:
	;
	v538 = v518 + v522
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v98+v519<<(uint(int32(2))%32))))
	v544 = int32(1)
	v547 = v526 + v544
	if v547 != v445 {
		v518 = v542 * v518
		v519 = v519 + v544
		v522 = v538
		v526 = v547
		goto L112
	} else {
		goto L114
	}
L113:
	;
	v553 = v538
	goto L111
L114:
	;
	goto L113
L115:
	;
	v579 = v23 - int32(-64)
	v580 = v569
	goto L116
L116:
	;
	v600 = v580 << (uint(int32(2)) % 32)
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v98+v600)))
	v604 = *(*int32)(unsafe.Add(mBase, uint32(v600+v99)))
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v604
	*(*int32)(unsafe.Add(mBase, uint32(v23)+4)) = v604 + v602 - int32(1)
	v611 = F_pg_sprintf(m, v579, int32(533691), v23)
	mBase = m.M
	v612 = m.ExcPending
	if v612 != 0 {
		goto L1
	} else {
		goto L118
	}
L117:
	;
	v644 = v614
	v646 = v574
	goto L97
L118:
	;
	v613 = F_strlen(m, v579)
	mBase = m.M
	v614 = v613 + v579
	v616 = v580 + int32(1)
	if v616 != v87 {
		v579 = v614
		v580 = v616
		goto L116
	} else {
		goto L119
	}
L119:
	;
	goto L117
L120:
	;
	v624 = v424
	goto L98
L121:
	;
	v760 = v642
	v770 = v642
	goto L96
L122:
	;
	v673 = v23 - int32(-64)
	if (v673^v670)&int32(3) != 0 {
		goto L126
	} else {
		goto L127
	}
L123:
	;
	v748 = F_strlen(m, v670)
	mBase = m.M
	v760 = v670
	v770 = v748 + v670
	goto L96
L124:
	;
	goto L123
L125:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v728))) = uint8(v727)
	if v727&int32(255) == int32(0) {
		goto L124
	} else {
		goto L140
	}
L126:
	;
	v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v673))))
	v726 = v673
	v727 = v679
	v728 = v670
	goto L125
L127:
	;
	goto L128
L128:
	;
	if v673&int32(3) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v683 = v673
	v685 = v670
	goto L132
L130:
	;
	v697 = v673
	v699 = v670
	goto L131
L131:
	;
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v697)))
	v704 = int32(-2139062144)
	if (int32(16843008)-v701|v701)&v704 != v704 {
		v726 = v697
		v727 = v701
		v728 = v699
		goto L125
	} else {
		goto L136
	}
L132:
	;
	v686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v683))))
	*(*uint8)(unsafe.Add(mBase, uint32(v685))) = uint8(v686)
	if v686 == int32(0) {
		goto L124
	} else {
		goto L134
	}
L133:
	;
	v697 = v693
	v699 = v691
	goto L131
L134:
	;
	v690 = int32(1)
	v691 = v685 + v690
	v693 = v683 + v690
	if v693&int32(3) != 0 {
		v683 = v693
		v685 = v691
		goto L132
	} else {
		goto L135
	}
L135:
	;
	goto L133
L136:
	;
	v709 = v697
	v710 = v701
	v711 = v699
	goto L137
L137:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v711))) = v710
	v713 = int32(4)
	v714 = v711 + v713
	v715 = *(*int32)(unsafe.Add(mBase, uint32(v709)+4))
	v717 = v709 + v713
	v721 = int32(-2139062144)
	if (v715|(int32(16843008)-v715))&v721 == v721 {
		v709 = v717
		v710 = v715
		v711 = v714
		goto L137
	} else {
		goto L139
	}
L138:
	;
	v726 = v717
	v727 = v715
	v728 = v714
	goto L125
L139:
	;
	goto L138
L140:
	;
	v735 = v726
	v737 = v728
	goto L141
L141:
	;
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v735)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v737)+1)) = uint8(v738)
	v740 = int32(1)
	if v738 != 0 {
		v735 = v735 + v740
		v737 = v737 + v740
		goto L141
	} else {
		goto L143
	}
L142:
	;
	goto L124
L143:
	;
	goto L142
L144:
	;
	v781 = F__emscripten_memset_bulkmem(m, v23+int32(32), base.I32_extend8_s(int32(0)), v87<<(uint(int32(2))%32))
	mBase = m.M
	goto L147
L145:
	;
	goto L146
L146:
	;
	v782 = int32(1)
	v787 = v87 - v782
	v788 = int32(0)
	v790 = v770 + v782
	v792 = v788
	v795 = v788
	goto L148
L147:
	;
	goto L146
L148:
	;
	if v787 <= v792 {
		v898 = v790
		goto L150
	} else {
		goto L151
	}
L149:
	;
	F_pfree(m, v163)
	mBase = m.M
	v1170 = m.ExcPending
	if v1170 != 0 {
		goto L1
	} else {
		goto L205
	}
L150:
	;
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v795+v165))))
	if v919 == int32(1) {
		goto L163
	} else {
		goto L164
	}
L151:
	;
	v816 = (v87 + (v792 ^ int32(-1))) & int32(7)
	if v816 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	v817 = v790
	v818 = v792
	v821 = int32(0)
	goto L155
L153:
	;
	v846 = v790
	v847 = v792
	goto L154
L154:
	;
	if base.Ui32(v87-int32(2)-v792) < base.Ui32(int32(7)) {
		v898 = v846
		goto L150
	} else {
		goto L158
	}
L155:
	;
	v837 = int32(123)
	*(*uint16)(unsafe.Add(mBase, uint32(v817))) = uint16(v837)
	v839 = int32(1)
	v840 = v818 + v839
	v842 = v817 + v839
	v844 = v821 + v839
	if v844 != v816 {
		v817 = v842
		v818 = v840
		v821 = v844
		goto L155
	} else {
		goto L157
	}
L156:
	;
	v846 = v842
	v847 = v840
	goto L154
L157:
	;
	goto L156
L158:
	;
	v869 = v846
	v870 = v847
	goto L159
L159:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v869))) = int64(8897841259083430779)
	v891 = int32(8)
	v892 = v869 + v891
	v893 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v892))) = uint8(v893)
	v896 = v870 + v891
	if v896 != v787 {
		v869 = v892
		v870 = v896
		goto L159
	} else {
		goto L161
	}
L160:
	;
	v898 = v892
	goto L150
L161:
	;
	goto L160
L162:
	;
	v1073 = *(*int32)(unsafe.Add(mBase, uint32(v163+v795<<(uint(int32(2))%32))))
	F_pfree(m, v1073)
	mBase = m.M
	v1075 = m.ExcPending
	if v1075 != 0 {
		goto L1
	} else {
		goto L194
	}
L163:
	;
	v922 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v898))) = uint16(v922)
	v929 = *(*int32)(unsafe.Add(mBase, uint32(v163+v795<<(uint(int32(2))%32))))
	v930 = v898 + int32(1)
	v931 = v929
	goto L166
L164:
	;
	goto L165
L165:
	;
	v972 = *(*int32)(unsafe.Add(mBase, uint32(v163+v795<<(uint(int32(2))%32))))
	if (v972^v898)&int32(3) != 0 {
		goto L176
	} else {
		goto L177
	}
L166:
	;
	v950 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v931))))
	if v950 == int32(34) {
		goto L169
	} else {
		goto L170
	}
L168:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v963))) = uint8(v950)
	v965 = int32(1)
	v930 = v963 + v965
	v931 = v931 + v965
	goto L166
L169:
	;
	v959 = int32(92)
	*(*uint8)(unsafe.Add(mBase, uint32(v930))) = uint8(v959)
	v963 = v930 + int32(1)
	goto L168
L170:
	;
	if v950 == int32(92) {
		goto L169
	} else {
		goto L171
	}
L171:
	;
	if v950 != 0 {
		v963 = v930
		goto L168
	} else {
		goto L172
	}
L172:
	;
	v955 = int32(34)
	*(*uint16)(unsafe.Add(mBase, uint32(v930))) = uint16(v955)
	v1069 = v930 + int32(1)
	goto L162
L173:
	;
	v1047 = F_strlen(m, v898)
	mBase = m.M
	v1069 = v1047 + v898
	goto L162
L174:
	;
	goto L173
L175:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1027))) = uint8(v1026)
	if v1026&int32(255) == int32(0) {
		goto L174
	} else {
		goto L190
	}
L176:
	;
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v972))))
	v1025 = v972
	v1026 = v978
	v1027 = v898
	goto L175
L177:
	;
	goto L178
L178:
	;
	if v972&int32(3) != 0 {
		goto L179
	} else {
		goto L180
	}
L179:
	;
	v982 = v972
	v984 = v898
	goto L182
L180:
	;
	v996 = v972
	v998 = v898
	goto L181
L181:
	;
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v996)))
	v1003 = int32(-2139062144)
	if (int32(16843008)-v1000|v1000)&v1003 != v1003 {
		v1025 = v996
		v1026 = v1000
		v1027 = v998
		goto L175
	} else {
		goto L186
	}
L182:
	;
	v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v982))))
	*(*uint8)(unsafe.Add(mBase, uint32(v984))) = uint8(v985)
	if v985 == int32(0) {
		goto L174
	} else {
		goto L184
	}
L183:
	;
	v996 = v992
	v998 = v990
	goto L181
L184:
	;
	v989 = int32(1)
	v990 = v984 + v989
	v992 = v982 + v989
	if v992&int32(3) != 0 {
		v982 = v992
		v984 = v990
		goto L182
	} else {
		goto L185
	}
L185:
	;
	goto L183
L186:
	;
	v1008 = v996
	v1009 = v1000
	v1010 = v998
	goto L187
L187:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1010))) = v1009
	v1012 = int32(4)
	v1013 = v1010 + v1012
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v1008)+4))
	v1016 = v1008 + v1012
	v1020 = int32(-2139062144)
	if (v1014|(int32(16843008)-v1014))&v1020 == v1020 {
		v1008 = v1016
		v1009 = v1014
		v1010 = v1013
		goto L187
	} else {
		goto L189
	}
L188:
	;
	v1025 = v1016
	v1026 = v1014
	v1027 = v1013
	goto L175
L189:
	;
	goto L188
L190:
	;
	v1034 = v1025
	v1036 = v1027
	goto L191
L191:
	;
	v1037 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1034)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v1036)+1)) = uint8(v1037)
	v1039 = int32(1)
	if v1037 != 0 {
		v1034 = v1034 + v1039
		v1036 = v1036 + v1039
		goto L191
	} else {
		goto L193
	}
L192:
	;
	goto L174
L193:
	;
	goto L192
L194:
	;
	if v787 < int32(0) {
		v1125 = v1069
		v1127 = v787
		goto L196
	} else {
		goto L197
	}
L195:
	;
	goto L149
L196:
	;
	if v1127 != int32(-1) {
		v790 = v1125
		v792 = v1127
		v795 = v795 + int32(1)
		goto L148
	} else {
		goto L204
	}
L197:
	;
	v1078 = v1069
	v1080 = v787
	goto L198
L198:
	;
	v1099 = v1080 << (uint(int32(2)) % 32)
	v1102 = v1099 + (v23 + int32(32))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1102)))
	v1105 = v1103 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1102))) = v1105
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v1099+v98)))
	if v1105 < v1108 {
		goto L200
	} else {
		goto L201
	}
L199:
	;
	goto L195
L200:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1078))) = uint8(v100)
	v1112 = v1078 + int32(1)
	v1113 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v1112))) = uint8(v1113)
	v1125 = v1112
	v1127 = v1080
	goto L196
L201:
	;
	goto L202
L202:
	;
	v1115 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1102))) = v1115
	v1117 = int32(125)
	*(*uint16)(unsafe.Add(mBase, uint32(v1078))) = uint16(v1117)
	v1119 = int32(1)
	if v1115 < v1080 {
		v1078 = v1078 + v1119
		v1080 = v1080 - v1119
		goto L198
	} else {
		goto L203
	}
L203:
	;
	goto L199
L204:
	;
	goto L195
L205:
	;
	F_pfree(m, v165)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L1
	} else {
		goto L206
	}
L206:
	;
	v1183 = v760
	goto L22
}
func F_array_prepend(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
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
	var v82 int32
	_ = v82
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 == v2 {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = v15
	} else {
		v16 = v2
	}
	v17 = int32(1)
	v19 = F_fetch_array_arg_replace_nulls(m, l0, v17)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
		switch v23 {
		case 0:
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(1)
			v64 = v17
			v66 = int32(12)
			v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
			v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+4)))
			v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+6)))
			v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+7)))
			v77 = F_array_set_element(m, v19+v66, int32(1), v10+v66, v16, v12, int32(-1), v74, v75, v76)
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
				if v79 == int32(1) {
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v82))) = v64
				} else {
				}
				m.G0 = v10 + int32(16)
				return v77
			}
		case 1:
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
			v27 = v25 - int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v27
			if v27 < v25 {
				v64 = v25
				v66 = int32(12)
				v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v73 = *(*int32)(unsafe.Add(mBase, uint32(v72)+16))
				v74 = int32(*(*int16)(unsafe.Add(mBase, uint32(v73)+4)))
				v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73)+6)))
				v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v73)+7)))
				v77 = F_array_set_element(m, v19+v66, int32(1), v10+v66, v16, v12, int32(-1), v74, v75, v76)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v19)+28))
					if v79 == int32(1) {
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v19)+36))
						*(*int32)(unsafe.Add(mBase, uint32(v82))) = v64
					} else {
					}
					m.G0 = v10 + int32(16)
					return v77
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50331778))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(421613), int32(0))
						mBase = m.M
						v40 = m.ExcPending
						if v40 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518919), int32(249), int32(447620))
							mBase = m.M
							v45 = m.ExcPending
							if v45 != 0 {
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
		default:
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v49 = m.ExcPending
			if v49 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(130))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(26178), int32(0))
					mBase = m.M
					v56 = m.ExcPending
					if v56 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(518919), int32(259), int32(447620))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
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
func F_array_remove(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	v4 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v4 == int32(1) {
		v7 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v7)
		return int32(0)
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v19 = int32(1)
			v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			v22 = F_array_replace_internal(m, v14, v11, v12, int32(0), v19, v19, v21, l0)
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				return v22
			}
		}
	}
}
func F_array_reverse(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
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
	var v53 int32
	_ = v53
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v116 int32
	_ = v116
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
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
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	v21 = m.G0
	v23 = v21 - int32(80)
	m.G0 = v23
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v29 = m.ExcPending
	if v29 != 0 {
		return int32(0)
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
		if v30 <= int32(0) {
			v216 = v26
			m.G0 = v23 + int32(80)
			return v216
		} else {
			v34 = v26 + int32(16)
			v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)))
			if v35 < int32(2) {
				v216 = v26
				m.G0 = v23 + int32(80)
				return v216
			} else {
				v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				if v40 != 0 {
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
					if v41 == v38 {
						v49 = v40
						v50 = v30
						v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+8)))
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+10)))
						v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49)+11)))
						F_deconstruct_array(m, v26, v51, v52, v53, v23+int32(12), v23+int32(8), v23+int32(76))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
							v64 = base.I32_div_s(v62, v63)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v64
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
							if int32(1) < v63 {
								v71 = base.I32_div_s(v63, int32(2))
								v75 = v66
								v76 = v67
								v77 = v64
								v81 = int32(0)
								for {
									if int32(0) < v77 {
										v97 = (v63 + (v81 ^ int32(-1))) * v77
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
										v105 = v100 + v97<<(uint(int32(2))%32)
										v107 = v97 + v98
										v108 = v75
										v109 = v76
										v116 = int32(0)
										for {
											v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
											*(*int32)(unsafe.Add(mBase, uint32(v109))) = v127
											v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
											*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v129)
											*(*int32)(unsafe.Add(mBase, uint32(v105))) = v126
											*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v125)
											v133 = int32(1)
											v135 = int32(4)
											v138 = v108 + v133
											v140 = v109 + v135
											v142 = v116 + v133
											v143 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
											if v142 < v143 {
												v105 = v105 + v135
												v107 = v107 + v133
												v108 = v138
												v109 = v140
												v116 = v142
												continue
											} else {
												break
											}
											break
										}
										v148 = v138
										v149 = v140
										v150 = v143
									} else {
										v148 = v75
										v149 = v76
										v150 = v77
									}
									v166 = v81 + int32(1)
									if v166 != v71 {
										v75 = v148
										v76 = v149
										v77 = v150
										v81 = v166
										continue
									} else {
										break
									}
									break
								}
								v168 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
								v173 = v169
								v174 = v168
							} else {
								v173 = v66
								v174 = v67
							}
							v193 = v50 << (uint(int32(2)) % 32)
							if v193 != 0 {
								v194 = F__emscripten_memcpy_bulkmem(m, v23+int32(48), v34, v193)
								mBase = m.M
							} else {
							}
							if v193 != 0 {
								v199 = F__emscripten_memcpy_bulkmem(m, v23+int32(16), v193+v34, v193)
								mBase = m.M
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v63
							v206 = F_construct_md_array(m, v174, v173, v50, v23+int32(48), v23+int32(16), v38, v51, v52, v53)
							mBase = m.M
							v207 = m.ExcPending
							if v207 != 0 {
								return int32(0)
							} else {
								v208 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
								F_pfree(m, v208)
								mBase = m.M
								v210 = m.ExcPending
								if v210 != 0 {
									return int32(0)
								} else {
									v211 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
									F_pfree(m, v211)
									mBase = m.M
									v213 = m.ExcPending
									if v213 != 0 {
										return int32(0)
									} else {
										v216 = v206
										m.G0 = v23 + int32(80)
										return v216
									}
								}
							}
						}
					} else {
						v44 = F_lookup_type_cache(m, v38, int32(0))
						mBase = m.M
						v45 = m.ExcPending
						if v45 != 0 {
							return int32(0)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v44
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
							v49 = v44
							v50 = v48
							v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+8)))
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+10)))
							v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49)+11)))
							F_deconstruct_array(m, v26, v51, v52, v53, v23+int32(12), v23+int32(8), v23+int32(76))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return int32(0)
							} else {
								v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
								v63 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
								v64 = base.I32_div_s(v62, v63)
								*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v64
								v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
								if int32(1) < v63 {
									v71 = base.I32_div_s(v63, int32(2))
									v75 = v66
									v76 = v67
									v77 = v64
									v81 = int32(0)
									for {
										if int32(0) < v77 {
											v97 = (v63 + (v81 ^ int32(-1))) * v77
											v98 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
											v100 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
											v105 = v100 + v97<<(uint(int32(2))%32)
											v107 = v97 + v98
											v108 = v75
											v109 = v76
											v116 = int32(0)
											for {
												v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
												v126 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
												v127 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
												*(*int32)(unsafe.Add(mBase, uint32(v109))) = v127
												v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
												*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v129)
												*(*int32)(unsafe.Add(mBase, uint32(v105))) = v126
												*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v125)
												v133 = int32(1)
												v135 = int32(4)
												v138 = v108 + v133
												v140 = v109 + v135
												v142 = v116 + v133
												v143 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
												if v142 < v143 {
													v105 = v105 + v135
													v107 = v107 + v133
													v108 = v138
													v109 = v140
													v116 = v142
													continue
												} else {
													break
												}
												break
											}
											v148 = v138
											v149 = v140
											v150 = v143
										} else {
											v148 = v75
											v149 = v76
											v150 = v77
										}
										v166 = v81 + int32(1)
										if v166 != v71 {
											v75 = v148
											v76 = v149
											v77 = v150
											v81 = v166
											continue
										} else {
											break
										}
										break
									}
									v168 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
									v169 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
									v173 = v169
									v174 = v168
								} else {
									v173 = v66
									v174 = v67
								}
								v193 = v50 << (uint(int32(2)) % 32)
								if v193 != 0 {
									v194 = F__emscripten_memcpy_bulkmem(m, v23+int32(48), v34, v193)
									mBase = m.M
								} else {
								}
								if v193 != 0 {
									v199 = F__emscripten_memcpy_bulkmem(m, v23+int32(16), v193+v34, v193)
									mBase = m.M
								} else {
								}
								*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v63
								v206 = F_construct_md_array(m, v174, v173, v50, v23+int32(48), v23+int32(16), v38, v51, v52, v53)
								mBase = m.M
								v207 = m.ExcPending
								if v207 != 0 {
									return int32(0)
								} else {
									v208 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
									F_pfree(m, v208)
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
										return int32(0)
									} else {
										v211 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
										F_pfree(m, v211)
										mBase = m.M
										v213 = m.ExcPending
										if v213 != 0 {
											return int32(0)
										} else {
											v216 = v206
											m.G0 = v23 + int32(80)
											return v216
										}
									}
								}
							}
						}
					}
				} else {
					v44 = F_lookup_type_cache(m, v38, int32(0))
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return int32(0)
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v46)+16)) = v44
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
						v49 = v44
						v50 = v48
						v51 = int32(*(*int16)(unsafe.Add(mBase, uint32(v49)+8)))
						v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49)+10)))
						v53 = int32(*(*int8)(unsafe.Add(mBase, uint32(v49)+11)))
						F_deconstruct_array(m, v26, v51, v52, v53, v23+int32(12), v23+int32(8), v23+int32(76))
						mBase = m.M
						v61 = m.ExcPending
						if v61 != 0 {
							return int32(0)
						} else {
							v62 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
							v64 = base.I32_div_s(v62, v63)
							*(*int32)(unsafe.Add(mBase, uint32(v23)+76)) = v64
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
							if int32(1) < v63 {
								v71 = base.I32_div_s(v63, int32(2))
								v75 = v66
								v76 = v67
								v77 = v64
								v81 = int32(0)
								for {
									if int32(0) < v77 {
										v97 = (v63 + (v81 ^ int32(-1))) * v77
										v98 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
										v100 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
										v105 = v100 + v97<<(uint(int32(2))%32)
										v107 = v97 + v98
										v108 = v75
										v109 = v76
										v116 = int32(0)
										for {
											v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v108))))
											v126 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
											v127 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
											*(*int32)(unsafe.Add(mBase, uint32(v109))) = v127
											v129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
											*(*uint8)(unsafe.Add(mBase, uint32(v108))) = uint8(v129)
											*(*int32)(unsafe.Add(mBase, uint32(v105))) = v126
											*(*uint8)(unsafe.Add(mBase, uint32(v107))) = uint8(v125)
											v133 = int32(1)
											v135 = int32(4)
											v138 = v108 + v133
											v140 = v109 + v135
											v142 = v116 + v133
											v143 = *(*int32)(unsafe.Add(mBase, uint32(v23)+76))
											if v142 < v143 {
												v105 = v105 + v135
												v107 = v107 + v133
												v108 = v138
												v109 = v140
												v116 = v142
												continue
											} else {
												break
											}
											break
										}
										v148 = v138
										v149 = v140
										v150 = v143
									} else {
										v148 = v75
										v149 = v76
										v150 = v77
									}
									v166 = v81 + int32(1)
									if v166 != v71 {
										v75 = v148
										v76 = v149
										v77 = v150
										v81 = v166
										continue
									} else {
										break
									}
									break
								}
								v168 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
								v169 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
								v173 = v169
								v174 = v168
							} else {
								v173 = v66
								v174 = v67
							}
							v193 = v50 << (uint(int32(2)) % 32)
							if v193 != 0 {
								v194 = F__emscripten_memcpy_bulkmem(m, v23+int32(48), v34, v193)
								mBase = m.M
							} else {
							}
							if v193 != 0 {
								v199 = F__emscripten_memcpy_bulkmem(m, v23+int32(16), v193+v34, v193)
								mBase = m.M
							} else {
							}
							*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v63
							v206 = F_construct_md_array(m, v174, v173, v50, v23+int32(48), v23+int32(16), v38, v51, v52, v53)
							mBase = m.M
							v207 = m.ExcPending
							if v207 != 0 {
								return int32(0)
							} else {
								v208 = *(*int32)(unsafe.Add(mBase, uint32(v23)+12))
								F_pfree(m, v208)
								mBase = m.M
								v210 = m.ExcPending
								if v210 != 0 {
									return int32(0)
								} else {
									v211 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
									F_pfree(m, v211)
									mBase = m.M
									v213 = m.ExcPending
									if v213 != 0 {
										return int32(0)
									} else {
										v216 = v206
										m.G0 = v23 + int32(80)
										return v216
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
func F_array_sample(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v12)+4))
		if int32(0) < v17 {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
			v21 = v20
		} else {
			v21 = int32(0)
		}
		if v16 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v46 = m.ExcPending
			if v46 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(50856066))
				mBase = m.M
				v49 = m.ExcPending
				if v49 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v9))) = v21
					F_errmsg(m, int32(501276), v9)
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(518919), int32(1750), int32(404168))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if v21 < v16 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v46 = m.ExcPending
				if v46 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v49 = m.ExcPending
					if v49 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = v21
						F_errmsg(m, int32(501276), v9)
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518919), int32(1750), int32(404168))
							mBase = m.M
							v58 = m.ExcPending
							if v58 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+16))
				if v27 != 0 {
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)))
					if v28 == v25 {
						v35 = v27
						v37 = F_array_shuffle_n(m, v12, v16, int32(0), v25, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v37
						}
					} else {
						v31 = F_lookup_type_cache(m, v25, int32(0))
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return int32(0)
						} else {
							v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v31
							v35 = v31
							v37 = F_array_shuffle_n(m, v12, v16, int32(0), v25, v35)
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								m.G0 = v9 + int32(16)
								return v37
							}
						}
					}
				} else {
					v31 = F_lookup_type_cache(m, v25, int32(0))
					mBase = m.M
					v32 = m.ExcPending
					if v32 != 0 {
						return int32(0)
					} else {
						v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						*(*int32)(unsafe.Add(mBase, uint32(v33)+16)) = v31
						v35 = v31
						v37 = F_array_shuffle_n(m, v12, v16, int32(0), v25, v35)
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							m.G0 = v9 + int32(16)
							return v37
						}
					}
				}
			}
		}
	}
}
func F_array_slice_size(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32, l8 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v99 int32
	_ = v99
	var v102 int32
	_ = v102
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
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
	var v188 int32
	_ = v188
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v209 int32
	_ = v209
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v231 int32
	_ = v231
	var v238 int32
	_ = v238
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v374 int32
	_ = v374
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v400 int32
	_ = v400
	var v420 int32
	_ = v420
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v435 int32
	_ = v435
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v493 int32
	_ = v493
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v521 int32
	_ = v521
	var v522 int32
	_ = v522
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v569 int32
	_ = v569
	var v582 int32
	_ = v582
	v10 = int32(0)
	v13 = m.G0
	v15 = v13 - int32(128)
	m.G0 = v15
	v18 = v15 + int32(96)
	if l2 <= v10 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l1 != 0 {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	goto L1
L3:
	;
	v25 = int32(1)
	if l2 != v25 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v35 = v10
	v36 = v10
	goto L7
L5:
	;
	v70 = v10
	goto L6
L6:
	;
	if l2&v25 == int32(0) {
		goto L2
	} else {
		goto L10
	}
L7:
	;
	v39 = int32(2)
	v40 = v35 << (uint(v39) % 32)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v40+l6)))
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v40+l5)))
	v47 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18+v40))) = v43 - v45 + v47
	v51 = v40 | int32(4)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v51+l6)))
	v56 = *(*int32)(unsafe.Add(mBase, uint32(v51+l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v18+v51))) = v54 - v56 + v47
	v62 = v35 + v39
	v64 = v36 + v39
	if v64 != l2&int32(2147483646) {
		v35 = v62
		v36 = v64
		goto L7
	} else {
		goto L9
	}
L8:
	;
	v70 = v62
	goto L6
L9:
	;
	goto L8
L10:
	;
	v77 = v70 << (uint(int32(2)) % 32)
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v77+l6)))
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v77+l5)))
	*(*int32)(unsafe.Add(mBase, uint32(v18+v77))) = v80 - v82 + int32(1)
	goto L2
L11:
	;
	m.G0 = v15 + int32(128)
	return v582
L12:
	;
	v121 = int32(0)
	v131 = l2 - int32(1)
	if v131 < v121 {
		v209 = v121
		goto L22
	} else {
		goto L23
	}
L13:
	;
	if l7 <= int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v99 = F_ArrayGetNItems(m, l2, v15+int32(96))
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	return int32(0)
L16:
	;
	switch l8 - int32(99) {
	case 0:
		v119 = l7
		goto L17
	case 1:
		goto L19
	default:
		goto L18
	case 6:
		goto L20
	}
L17:
	;
	v582 = v99 * v119
	goto L11
L18:
	;
	v119 = (l7 + int32(1)) & int32(-2)
	goto L17
L19:
	;
	v582 = (l7 + int32(7)) & int32(-8) * v99
	goto L11
L20:
	;
	v582 = (l7 + int32(3)) & int32(-4) * v99
	goto L11
L21:
	;
	v217 = F_array_seek(m, l0, v121, l1, v209, l7, l8)
	mBase = m.M
	v219 = v15 - int32(-64)
	v223 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v219+l2<<(uint(v223)%32)-int32(4)))) = int32(1)
	v231 = l2 - v223
	if v231 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L22:
	;
	goto L21
L23:
	;
	v134 = int32(1)
	if v131 == int32(0) {
		goto L25
	} else {
		goto L26
	}
L24:
	;
	if l2&v134 == int32(0) {
		v209 = v186
		goto L22
	} else {
		goto L31
	}
L25:
	;
	v186 = v121
	v187 = v131
	v188 = v134
	goto L24
L26:
	;
	goto L27
L27:
	;
	v145 = v121
	v146 = v131
	v147 = v134
	v148 = v121
	goto L28
L28:
	;
	v153 = int32(2)
	v154 = v146 << (uint(v153) % 32)
	v156 = v154 - int32(4)
	v158 = *(*int32)(unsafe.Add(mBase, uint32(l5+v156)))
	v160 = *(*int32)(unsafe.Add(mBase, uint32(l4+v156)))
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v154+l3)))
	v164 = v163 * v147
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v154+l5)))
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v154+l4)))
	v173 = (v158-v160)*v164 + ((v167-v169)*v147 + v145)
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l3+v156)))
	v176 = v175 * v164
	v178 = v146 - v153
	v180 = v148 + v153
	if v180 != l2&int32(-2) {
		v145 = v173
		v146 = v178
		v147 = v176
		v148 = v180
		goto L28
	} else {
		goto L30
	}
L29:
	;
	v186 = v173
	v187 = v178
	v188 = v176
	goto L24
L30:
	;
	goto L29
L31:
	;
	v197 = v187 << (uint(int32(2)) % 32)
	v199 = *(*int32)(unsafe.Add(mBase, uint32(l5+v197)))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v197+l4)))
	v209 = (v199-v201)*v188 + v186
	goto L22
L32:
	;
	v291 = v15 + int32(32)
	v293 = v15 - int32(-64)
	v295 = v15 + int32(96)
	v296 = int32(0)
	v302 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v291+l2<<(uint(v302)%32)-int32(4)))) = v296
	v310 = l2 - v302
	if v296 <= v310 {
		goto L43
	} else {
		goto L44
	}
L33:
	;
	goto L32
L34:
	;
	if l2&int32(1) == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v238 = int32(2)
	v244 = l2<<(uint(v238)%32) - int32(4)
	v246 = *(*int32)(unsafe.Add(mBase, uint32(l3+v244)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v219+v244)))
	*(*int32)(unsafe.Add(mBase, uint32(v219+v231<<(uint(v238)%32)))) = v246 * v248
	v253 = l2 - int32(3)
	goto L37
L36:
	;
	v253 = v231
	goto L37
L37:
	;
	if v231 == int32(0) {
		goto L33
	} else {
		goto L38
	}
L38:
	;
	v259 = v253
	goto L39
L39:
	;
	v262 = int32(2)
	v263 = v259 << (uint(v262) % 32)
	v266 = v263 + int32(4)
	v268 = *(*int32)(unsafe.Add(mBase, uint32(l3+v266)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v219+v266)))
	v271 = v268 * v270
	*(*int32)(unsafe.Add(mBase, uint32(v219+v263))) = v271
	v274 = v259 - int32(1)
	v279 = *(*int32)(unsafe.Add(mBase, uint32(v263+l3)))
	*(*int32)(unsafe.Add(mBase, uint32(v219+v274<<(uint(v262)%32)))) = v279 * v271
	if v274 != 0 {
		v259 = v259 - v262
		goto L39
	} else {
		goto L41
	}
L40:
	;
	goto L33
L41:
	;
	goto L40
L42:
	;
	v420 = F__emscripten_memset_bulkmem(m, v15, base.I32_extend8_s(int32(0)), l2<<(uint(int32(2))%32))
	mBase = m.M
	goto L58
L43:
	;
	v317 = v310
	v319 = v296
	goto L46
L44:
	;
	goto L45
L45:
	;
	goto L42
L46:
	;
	v324 = v317 << (uint(int32(2)) % 32)
	v325 = v291 + v324
	v327 = *(*int32)(unsafe.Add(mBase, uint32(v293+v324)))
	v328 = int32(1)
	v329 = v327 - v328
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v329
	v332 = v317 + v328
	if l2 <= v332 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	goto L45
L48:
	;
	v400 = int32(1)
	if int32(0) < v317 {
		v317 = v317 - v400
		v319 = v319 + v400
		goto L46
	} else {
		goto L57
	}
L49:
	;
	if v319&int32(1) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L50:
	;
	v338 = int32(2)
	v339 = v332 << (uint(v338) % 32)
	v341 = *(*int32)(unsafe.Add(mBase, uint32(v295+v339)))
	v345 = *(*int32)(unsafe.Add(mBase, uint32(v293+v339)))
	v347 = v329 - (v341-int32(1))*v345
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v347
	v351 = v317 + v338
	v352 = v347
	goto L52
L51:
	;
	v351 = v332
	v352 = v329
	goto L52
L52:
	;
	if v319 == int32(0) {
		goto L48
	} else {
		goto L53
	}
L53:
	;
	v359 = v351
	v360 = v352
	goto L54
L54:
	;
	v365 = int32(2)
	v366 = v359 << (uint(v365) % 32)
	v368 = *(*int32)(unsafe.Add(mBase, uint32(v295+v366)))
	v369 = int32(1)
	v372 = *(*int32)(unsafe.Add(mBase, uint32(v293+v366)))
	v374 = v360 - (v368-v369)*v372
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v374
	v377 = v366 + int32(4)
	v379 = *(*int32)(unsafe.Add(mBase, uint32(v295+v377)))
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v293+v377)))
	v385 = v374 - (v379-v369)*v383
	*(*int32)(unsafe.Add(mBase, uint32(v325))) = v385
	v388 = v359 + v365
	if v388 != l2 {
		v359 = v388
		v360 = v385
		goto L54
	} else {
		goto L56
	}
L55:
	;
	goto L48
L56:
	;
	goto L55
L57:
	;
	goto L47
L58:
	;
	v429 = v217
	v430 = v209
	v431 = l2 - int32(1)
	v435 = v10
	goto L59
L59:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v420+int32(32)+v431<<(uint(int32(2))%32))))
	if v442 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L60:
	;
	v582 = v511
	goto L11
L61:
	;
	if l1 != 0 {
		goto L66
	} else {
		goto L67
	}
L62:
	;
	v447 = v429
	v448 = v430
	goto L61
L63:
	;
	goto L64
L64:
	;
	v446 = F_array_seek(m, v429, v430, l1, v442, l7, l8)
	mBase = m.M
	v447 = v446
	v448 = v442 + v430
	goto L61
L65:
	;
	v515 = v420 + int32(96)
	if l2 <= int32(0) {
		goto L95
	} else {
		goto L96
	}
L66:
	;
	v450 = base.I32_div_s(v448, int32(8))
	v452 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1+v450))))
	if int32(base.Ui32(v452)>>(uint(v448&int32(7))%32))&int32(1) == int32(0) {
		v509 = v447
		v511 = v435
		goto L65
	} else {
		goto L69
	}
L67:
	;
	goto L68
L68:
	;
	if int32(0) < l7 {
		v493 = l7
		goto L70
	} else {
		goto L71
	}
L69:
	;
	goto L68
L70:
	;
	switch l8 - int32(99) {
	case 0:
		v506 = v493
		goto L90
	case 1:
		goto L92
	default:
		goto L91
	case 6:
		goto L93
	}
L71:
	;
	if l7 == int32(-1) {
		goto L72
	} else {
		goto L73
	}
L72:
	;
	v464 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447))))
	if v464 == int32(1) {
		goto L75
	} else {
		goto L76
	}
L73:
	;
	goto L74
L74:
	;
	v490 = F_strlen(m, v447)
	mBase = m.M
	v493 = v490 + int32(1)
	goto L70
L75:
	;
	v467 = int32(6)
	v469 = int32(18)
	v471 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v447)+1)))
	if v471 == v469 {
		goto L78
	} else {
		goto L79
	}
L76:
	;
	goto L77
L77:
	;
	if v464&int32(1) != 0 {
		goto L87
	} else {
		goto L88
	}
L78:
	;
	v474 = v469
	goto L80
L79:
	;
	v474 = int32(2)
	goto L80
L80:
	;
	if v471&int32(254) == int32(2) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v479 = v467
	goto L83
L82:
	;
	v479 = v474
	goto L83
L83:
	;
	if v471 == int32(1) {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	v482 = v467
	goto L86
L85:
	;
	v482 = v479
	goto L86
L86:
	;
	v493 = v482
	goto L70
L87:
	;
	v493 = int32(base.Ui32(v464) >> (uint(int32(1)) % 32))
	goto L70
L88:
	;
	goto L89
L89:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v447)))
	v493 = int32(base.Ui32(v487) >> (uint(int32(2)) % 32))
	goto L70
L90:
	;
	v509 = v447 + v506
	v511 = v506 + v435
	goto L65
L91:
	;
	v506 = (v493 + int32(1)) & int32(-2)
	goto L90
L92:
	;
	v506 = (v493 + int32(7)) & int32(-8)
	goto L90
L93:
	;
	v506 = (v493 + int32(3)) & int32(-4)
	goto L90
L94:
	;
	if v569 != int32(-1) {
		v429 = v509
		v430 = v448 + int32(1)
		v431 = v569
		v435 = v511
		goto L59
	} else {
		goto L109
	}
L95:
	;
	v569 = int32(-1)
	goto L94
L96:
	;
	goto L97
L97:
	;
	v521 = int32(1)
	v522 = l2 - v521
	v524 = v522 << (uint(int32(2)) % 32)
	v525 = v420 + v524
	v526 = *(*int32)(unsafe.Add(mBase, uint32(v525)))
	v530 = *(*int32)(unsafe.Add(mBase, uint32(v515+v524)))
	v531 = base.I32_rem_s(v526+v521, v530)
	*(*int32)(unsafe.Add(mBase, uint32(v525))) = v531
	if v522 != 0 {
		goto L99
	} else {
		goto L100
	}
L98:
	;
	v569 = v559
	goto L94
L99:
	;
	v533 = v522
	v536 = v531
	goto L102
L100:
	;
	goto L101
L101:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v420)))
	if v557 != 0 {
		goto L106
	} else {
		goto L107
	}
L102:
	;
	if v536 != 0 {
		v559 = v533
		goto L98
	} else {
		goto L104
	}
L103:
	;
	goto L101
L104:
	;
	v538 = int32(1)
	v539 = v533 - v538
	v541 = v539 << (uint(int32(2)) % 32)
	v542 = v420 + v541
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v542)))
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v515+v541)))
	v548 = base.I32_rem_s(v543+v538, v547)
	*(*int32)(unsafe.Add(mBase, uint32(v542))) = v548
	if v539 != 0 {
		v533 = v539
		v536 = v548
		goto L102
	} else {
		goto L105
	}
L105:
	;
	goto L103
L106:
	;
	v558 = int32(0)
	goto L108
L107:
	;
	v558 = int32(-1)
	goto L108
L108:
	;
	v559 = v558
	goto L98
L109:
	;
	goto L60
}
func F_array_subscript_assign_slice(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
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
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
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
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v101 int32
	_ = v101
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v250 int32
	_ = v250
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v325 int32
	_ = v325
	var v330 int32
	_ = v330
	var v334 int32
	_ = v334
	var v339 int32
	_ = v339
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v384 int32
	_ = v384
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v396 int32
	_ = v396
	var v400 int32
	_ = v400
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v421 int32
	_ = v421
	var v425 int32
	_ = v425
	var v428 int32
	_ = v428
	var v432 int32
	_ = v432
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v453 int32
	_ = v453
	var v456 int32
	_ = v456
	var v490 int32
	_ = v490
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v527 int32
	_ = v527
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v538 int32
	_ = v538
	var v541 int32
	_ = v541
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v604 int32
	_ = v604
	var v608 int32
	_ = v608
	var v610 int32
	_ = v610
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v670 int32
	_ = v670
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v727 int32
	_ = v727
	var v742 int32
	_ = v742
	var v743 int32
	_ = v743
	var v744 int32
	_ = v744
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v751 int32
	_ = v751
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v773 int32
	_ = v773
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v805 int32
	_ = v805
	var v806 int32
	_ = v806
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v830 int32
	_ = v830
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v842 int32
	_ = v842
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v858 int32
	_ = v858
	var v863 int32
	_ = v863
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v879 int32
	_ = v879
	var v884 int32
	_ = v884
	var v888 int32
	_ = v888
	var v891 int32
	_ = v891
	var v895 int32
	_ = v895
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v912 int32
	_ = v912
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v924 int32
	_ = v924
	var v926 int32
	_ = v926
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v935 int32
	_ = v935
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v946 int32
	_ = v946
	var v948 int32
	_ = v948
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v994 int32
	_ = v994
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1005 int32
	_ = v1005
	var v1012 int32
	_ = v1012
	var v1013 int32
	_ = v1013
	var v1016 int32
	_ = v1016
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1026 int32
	_ = v1026
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1049 int32
	_ = v1049
	var v1051 int32
	_ = v1051
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1075 int32
	_ = v1075
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
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1097 int32
	_ = v1097
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1127 int32
	_ = v1127
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1139 int32
	_ = v1139
	var v1147 int32
	_ = v1147
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1155 int32
	_ = v1155
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1167 int32
	_ = v1167
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1175 int32
	_ = v1175
	var v1185 int32
	_ = v1185
	var v1189 int32
	_ = v1189
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1213 int32
	_ = v1213
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1228 int32
	_ = v1228
	var v1230 int32
	_ = v1230
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1244 int32
	_ = v1244
	var v1246 int32
	_ = v1246
	var v1251 int32
	_ = v1251
	var v1280 int32
	_ = v1280
	var v1295 int32
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1301 int32
	_ = v1301
	var v1302 int32
	_ = v1302
	var v1303 int32
	_ = v1303
	var v1324 int32
	_ = v1324
	var v1329 int32
	_ = v1329
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1396 int32
	_ = v1396
	var v1404 int32
	_ = v1404
	var v1411 int32
	_ = v1411
	var v1417 int32
	_ = v1417
	var v1419 int32
	_ = v1419
	var v1421 int32
	_ = v1421
	var v1426 int32
	_ = v1426
	var v1432 int32
	_ = v1432
	var v1435 int32
	_ = v1435
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1441 int32
	_ = v1441
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1447 int32
	_ = v1447
	var v1452 int32
	_ = v1452
	var v1464 int32
	_ = v1464
	var v1465 int32
	_ = v1465
	var v1471 int32
	_ = v1471
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1497 int32
	_ = v1497
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1508 int32
	_ = v1508
	var v1510 int32
	_ = v1510
	var v1516 int32
	_ = v1516
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1528 int32
	_ = v1528
	var v1542 int32
	_ = v1542
	var v1544 int32
	_ = v1544
	var v1546 int32
	_ = v1546
	var v1547 int32
	_ = v1547
	var v1553 int32
	_ = v1553
	var v1561 int32
	_ = v1561
	var v1568 int32
	_ = v1568
	var v1570 int32
	_ = v1570
	var v1575 int32
	_ = v1575
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1579 int32
	_ = v1579
	var v1580 int32
	_ = v1580
	var v1583 int32
	_ = v1583
	var v1589 int32
	_ = v1589
	var v1590 int32
	_ = v1590
	var v1592 int32
	_ = v1592
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1616 int32
	_ = v1616
	var v1617 int32
	_ = v1617
	var v1619 int32
	_ = v1619
	var v1620 int32
	_ = v1620
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1630 int32
	_ = v1630
	var v1634 int32
	_ = v1634
	var v1636 int32
	_ = v1636
	var v1639 int32
	_ = v1639
	var v1651 int32
	_ = v1651
	var v1667 int32
	_ = v1667
	var v1672 int32
	_ = v1672
	var v1675 int32
	_ = v1675
	var v1677 int32
	_ = v1677
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1688 int32
	_ = v1688
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1716 int32
	_ = v1716
	var v1717 int32
	_ = v1717
	var v1718 int32
	_ = v1718
	var v1719 int32
	_ = v1719
	var v1722 int32
	_ = v1722
	var v1728 int32
	_ = v1728
	var v1730 int32
	_ = v1730
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1738 int32
	_ = v1738
	var v1739 int32
	_ = v1739
	var v1740 int32
	_ = v1740
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1752 int32
	_ = v1752
	var v1755 int32
	_ = v1755
	var v1763 int32
	_ = v1763
	var v1779 int32
	_ = v1779
	var v1780 int32
	_ = v1780
	var v1781 int32
	_ = v1781
	var v1783 int32
	_ = v1783
	var v1789 int32
	_ = v1789
	var v1791 int32
	_ = v1791
	var v1792 int32
	_ = v1792
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1795 int32
	_ = v1795
	var v1797 int32
	_ = v1797
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1806 int32
	_ = v1806
	var v1807 int32
	_ = v1807
	var v1808 int32
	_ = v1808
	var v1820 int32
	_ = v1820
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1829 int32
	_ = v1829
	var v1851 int32
	_ = v1851
	var v1852 int32
	_ = v1852
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1861 int32
	_ = v1861
	var v1863 int32
	_ = v1863
	var v1864 int32
	_ = v1864
	var v1871 int32
	_ = v1871
	var v1878 int32
	_ = v1878
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1940 int32
	_ = v1940
	var v1941 int32
	_ = v1941
	var v1962 int32
	_ = v1962
	var v1972 int32
	_ = v1972
	var v1973 int32
	_ = v1973
	var v1974 int32
	_ = v1974
	var v1975 int32
	_ = v1975
	var v1979 int32
	_ = v1979
	var v1981 int32
	_ = v1981
	var v1982 int32
	_ = v1982
	var v1983 int32
	_ = v1983
	var v1989 int32
	_ = v1989
	var v1991 int32
	_ = v1991
	var v1997 int32
	_ = v1997
	var v1999 int32
	_ = v1999
	var v2004 int32
	_ = v2004
	var v2005 int32
	_ = v2005
	var v2009 int32
	_ = v2009
	var v2011 int32
	_ = v2011
	var v2013 int32
	_ = v2013
	var v2015 int32
	_ = v2015
	var v2017 int32
	_ = v2017
	var v2023 int32
	_ = v2023
	var v2024 int32
	_ = v2024
	var v2026 int32
	_ = v2026
	var v2027 int32
	_ = v2027
	var v2028 int32
	_ = v2028
	var v2032 int32
	_ = v2032
	var v2033 int32
	_ = v2033
	var v2035 int32
	_ = v2035
	var v2038 int32
	_ = v2038
	var v2040 int32
	_ = v2040
	var v2041 int32
	_ = v2041
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2045 int32
	_ = v2045
	var v2049 int32
	_ = v2049
	var v2050 int32
	_ = v2050
	var v2059 int32
	_ = v2059
	var v2060 int32
	_ = v2060
	var v2061 int32
	_ = v2061
	var v2071 int32
	_ = v2071
	var v2074 int32
	_ = v2074
	var v2075 int32
	_ = v2075
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2086 int32
	_ = v2086
	var v2088 int32
	_ = v2088
	var v2089 int32
	_ = v2089
	var v2090 int32
	_ = v2090
	var v2093 int32
	_ = v2093
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2099 int32
	_ = v2099
	var v2124 int32
	_ = v2124
	var v2125 int32
	_ = v2125
	var v2126 int32
	_ = v2126
	var v2128 int32
	_ = v2128
	var v2134 int32
	_ = v2134
	var v2136 int32
	_ = v2136
	var v2137 int32
	_ = v2137
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2142 int32
	_ = v2142
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2146 int32
	_ = v2146
	var v2153 int32
	_ = v2153
	var v2155 int32
	_ = v2155
	var v2156 int32
	_ = v2156
	var v2160 int32
	_ = v2160
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2179 int32
	_ = v2179
	var v2185 int32
	_ = v2185
	var v2187 int32
	_ = v2187
	var v2188 int32
	_ = v2188
	var v2189 int32
	_ = v2189
	var v2192 int32
	_ = v2192
	var v2195 int32
	_ = v2195
	var v2198 int32
	_ = v2198
	var v2206 int32
	_ = v2206
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2227 int32
	_ = v2227
	var v2229 int32
	_ = v2229
	var v2235 int32
	_ = v2235
	var v2237 int32
	_ = v2237
	var v2238 int32
	_ = v2238
	var v2241 int32
	_ = v2241
	var v2243 int32
	_ = v2243
	var v2246 int32
	_ = v2246
	var v2249 int32
	_ = v2249
	var v2251 int32
	_ = v2251
	var v2252 int32
	_ = v2252
	var v2257 int32
	_ = v2257
	var v2262 int32
	_ = v2262
	var v2281 int32
	_ = v2281
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2285 int32
	_ = v2285
	var v2291 int32
	_ = v2291
	var v2293 int32
	_ = v2293
	var v2294 int32
	_ = v2294
	var v2295 int32
	_ = v2295
	var v2296 int32
	_ = v2296
	var v2297 int32
	_ = v2297
	var v2299 int32
	_ = v2299
	var v2304 int32
	_ = v2304
	var v2305 int32
	_ = v2305
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2310 int32
	_ = v2310
	var v2326 int32
	_ = v2326
	var v2334 int32
	_ = v2334
	var v2387 int32
	_ = v2387
	var v2388 int32
	_ = v2388
	var v2393 int32
	_ = v2393
	var v2397 int32
	_ = v2397
	var v2398 int32
	_ = v2398
	var v2399 int32
	_ = v2399
	var v2402 int32
	_ = v2402
	var v2404 int32
	_ = v2404
	var v2405 int32
	_ = v2405
	var v2406 int32
	_ = v2406
	var v2411 int32
	_ = v2411
	var v2412 int32
	_ = v2412
	var v2414 int32
	_ = v2414
	var v2419 int32
	_ = v2419
	var v2442 int32
	_ = v2442
	var v2443 int32
	_ = v2443
	var v2444 int32
	_ = v2444
	var v2446 int32
	_ = v2446
	var v2452 int32
	_ = v2452
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2458 int32
	_ = v2458
	var v2462 int32
	_ = v2462
	var v2463 int32
	_ = v2463
	var v2465 int32
	_ = v2465
	var v2466 int32
	_ = v2466
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2474 int32
	_ = v2474
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2502 int32
	_ = v2502
	var v2508 int32
	_ = v2508
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2513 int32
	_ = v2513
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2521 int32
	_ = v2521
	var v2522 int32
	_ = v2522
	var v2523 int32
	_ = v2523
	var v2525 int32
	_ = v2525
	var v2526 int32
	_ = v2526
	var v2527 int32
	_ = v2527
	var v2542 int32
	_ = v2542
	var v2547 int32
	_ = v2547
	var v2607 int32
	_ = v2607
	var v2608 int32
	_ = v2608
	var v2611 int32
	_ = v2611
	var v2613 int32
	_ = v2613
	var v2614 int32
	_ = v2614
	var v2615 int32
	_ = v2615
	var v2620 int32
	_ = v2620
	var v2621 int32
	_ = v2621
	var v2622 int32
	_ = v2622
	var v2626 int32
	_ = v2626
	var v2651 int32
	_ = v2651
	var v2652 int32
	_ = v2652
	var v2653 int32
	_ = v2653
	var v2655 int32
	_ = v2655
	var v2661 int32
	_ = v2661
	var v2663 int32
	_ = v2663
	var v2664 int32
	_ = v2664
	var v2672 int32
	_ = v2672
	var v2673 int32
	_ = v2673
	var v2674 int32
	_ = v2674
	var v2677 int32
	_ = v2677
	var v2678 int32
	_ = v2678
	var v2679 int32
	_ = v2679
	var v2680 int32
	_ = v2680
	var v2683 int32
	_ = v2683
	var v2685 int32
	_ = v2685
	var v2686 int32
	_ = v2686
	var v2713 int32
	_ = v2713
	var v2714 int32
	_ = v2714
	var v2715 int32
	_ = v2715
	var v2717 int32
	_ = v2717
	var v2723 int32
	_ = v2723
	var v2725 int32
	_ = v2725
	var v2726 int32
	_ = v2726
	var v2727 int32
	_ = v2727
	var v2728 int32
	_ = v2728
	var v2729 int32
	_ = v2729
	var v2731 int32
	_ = v2731
	var v2736 int32
	_ = v2736
	var v2737 int32
	_ = v2737
	var v2738 int32
	_ = v2738
	var v2740 int32
	_ = v2740
	var v2741 int32
	_ = v2741
	var v2742 int32
	_ = v2742
	var v2754 int32
	_ = v2754
	var v2760 int32
	_ = v2760
	var v2791 int32
	_ = v2791
	var v2792 int32
	_ = v2792
	var v2793 int32
	_ = v2793
	var v2796 int32
	_ = v2796
	var v2797 int32
	_ = v2797
	var v2798 int32
	_ = v2798
	var v2799 int32
	_ = v2799
	var v2802 int32
	_ = v2802
	var v2804 int32
	_ = v2804
	var v2805 int32
	_ = v2805
	var v2832 int32
	_ = v2832
	var v2833 int32
	_ = v2833
	var v2834 int32
	_ = v2834
	var v2836 int32
	_ = v2836
	var v2842 int32
	_ = v2842
	var v2844 int32
	_ = v2844
	var v2845 int32
	_ = v2845
	var v2846 int32
	_ = v2846
	var v2847 int32
	_ = v2847
	var v2848 int32
	_ = v2848
	var v2850 int32
	_ = v2850
	var v2855 int32
	_ = v2855
	var v2856 int32
	_ = v2856
	var v2857 int32
	_ = v2857
	var v2859 int32
	_ = v2859
	var v2860 int32
	_ = v2860
	var v2861 int32
	_ = v2861
	var v2873 int32
	_ = v2873
	var v2879 int32
	_ = v2879
	var v2920 int32
	_ = v2920
	var v2941 int32
	_ = v2941
	v4 = int32(0)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v34))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	if v40 <= v4 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	return
L2:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	v63 = v39 + int32(12)
	v65 = v39 + int32(36)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v38)+12))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v38)+28))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v38)+40))
	v69 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+44)))
	v70 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+6)))
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39)+8)))
	v72 = int32(*(*int8)(unsafe.Add(mBase, uint32(v39)+9)))
	v73 = m.G0
	v75 = v73 - int32(272)
	m.G0 = v75
	if v69 != 0 {
		v2920 = v60
		goto L13
	} else {
		goto L14
	}
L3:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v53 = F_construct_empty_array(m, v52)
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L11
	} else {
		goto L12
	}
L4:
	;
	if v35&int32(1) != 0 {
		goto L3
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	if v35&int32(1) != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v59 = v40
	v60 = v37
	goto L2
L8:
	;
	v47 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v38)+44)))
	if v47 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v35&int32(1) == int32(0) {
		v59 = v40
		v60 = v37
		goto L2
	} else {
		goto L10
	}
L10:
	;
	goto L3
L11:
	;
	return
L12:
	;
	v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v56 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v55))) = uint8(v56)
	v58 = int32(*(*int16)(unsafe.Add(mBase, uint32(v39)+4)))
	v59 = v58
	v60 = v53
	goto L2
L13:
	;
	m.G0 = v75 + int32(272)
	v2941 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v2941))) = v2920
	goto L1
L14:
	;
	if v59 <= int32(0) {
		goto L30
	} else {
		goto L31
	}
L15:
	;
	v970 = v785 + v769 + v955 - v952
	v971 = F_palloc0(m, v970)
	mBase = m.M
	v972 = m.ExcPending
	if v972 != 0 {
		goto L11
	} else {
		goto L208
	}
L16:
	;
	v905 = v901 << (uint(int32(3)) % 32)
	if v786 != 0 {
		goto L187
	} else {
		goto L188
	}
L17:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v888 = m.ExcPending
	if v888 != 0 {
		goto L11
	} else {
		goto L183
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v869 = m.ExcPending
	if v869 != 0 {
		goto L11
	} else {
		goto L179
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v848 = m.ExcPending
	if v848 != 0 {
		goto L11
	} else {
		goto L175
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L11
	} else {
		goto L171
	}
L21:
	;
	v654 = F_ArrayGetNItems(m, v83, v75+int32(112))
	mBase = m.M
	v655 = m.ExcPending
	if v655 != 0 {
		goto L11
	} else {
		goto L137
	}
L22:
	;
	v559 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v67))))
	if v559 == int32(0) {
		goto L119
	} else {
		goto L120
	}
L23:
	;
	if v83 <= v456 {
		v631 = v240
		v635 = v4
		goto L21
	} else {
		goto L109
	}
L24:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v441 = m.ExcPending
	if v441 != 0 {
		goto L11
	} else {
		goto L105
	}
L25:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L11
	} else {
		goto L101
	}
L26:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L11
	} else {
		goto L97
	}
L27:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L11
	} else {
		goto L93
	}
L28:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v376 = m.ExcPending
	if v376 != 0 {
		goto L11
	} else {
		goto L89
	}
L29:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L11
	} else {
		goto L84
	}
L30:
	;
	v79 = F_pg_detoast_datum(m, v60)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L11
	} else {
		goto L33
	}
L31:
	;
	goto L32
L32:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L80
	}
L33:
	;
	v81 = F_pg_detoast_datum(m, v68)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v83 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	F_deconstruct_array(m, v81, v70, v71, v72, v75+int32(240), v75+int32(208), v75+int32(176))
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L11
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	if v83 < v61 {
		goto L26
	} else {
		goto L52
	}
L38:
	;
	if int32(0) < v61 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v101 = int32(0)
	goto L42
L40:
	;
	goto L41
L41:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v75)+176))
	v201 = F_ArrayGetNItems(m, v61, v75+int32(112))
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L11
	} else {
		goto L49
	}
L42:
	;
	v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v66))))
	if v132 != int32(1) {
		goto L29
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101+v67))))
	if v136 == int32(0) {
		goto L29
	} else {
		goto L45
	}
L45:
	;
	v140 = v101 << (uint(int32(2)) % 32)
	v143 = v140 + (v75 + int32(112))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v140+v63)))
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v140+v65)))
	v148 = v145 - v147
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v148
	if base.B2i32(v148 < v145)^base.B2i32(int32(0) < v147) != 0 {
		goto L28
	} else {
		goto L46
	}
L46:
	;
	v155 = v148 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v143))) = v155
	if v155 < v148 {
		goto L28
	} else {
		goto L47
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75+int32(80)+v140))) = v147
	v163 = v101 + int32(1)
	if v163 != v61 {
		v101 = v163
		goto L42
	} else {
		goto L48
	}
L48:
	;
	goto L43
L49:
	;
	if v198 < v201 {
		goto L27
	} else {
		goto L50
	}
L50:
	;
	v204 = *(*int32)(unsafe.Add(mBase, uint32(v75)+240))
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v75)+208))
	v210 = F_construct_md_array(m, v204, v205, v61, v75+int32(112), v75+int32(80), v86, v70, v71, v72)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L11
	} else {
		goto L51
	}
L51:
	;
	v2920 = v210
	goto L13
L52:
	;
	if base.Ui32(v83-int32(7)) <= base.Ui32(int32(-7)) {
		goto L26
	} else {
		goto L53
	}
L53:
	;
	v220 = v79 + int32(16)
	v222 = v83 << (uint(int32(2)) % 32)
	if v222 != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v222 != 0 {
		goto L59
	} else {
		goto L60
	}
L55:
	;
	v223 = F__emscripten_memcpy_bulkmem(m, v75+int32(112), v220, v222)
	mBase = m.M
	goto L57
L56:
	;
	goto L57
L57:
	;
	goto L54
L58:
	;
	v234 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v234 == int32(0) {
		goto L62
	} else {
		goto L63
	}
L59:
	;
	v231 = F__emscripten_memcpy_bulkmem(m, v75+int32(80), v220+v227<<(uint(int32(2))%32), v222)
	mBase = m.M
	goto L61
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	v237 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v240 = base.B2i32(v237 != int32(0))
	goto L64
L63:
	;
	v240 = int32(1)
	goto L64
L64:
	;
	if v83 == int32(1) {
		goto L22
	} else {
		goto L65
	}
L65:
	;
	v243 = int32(0)
	if v61 <= v243 {
		v456 = v243
		goto L23
	} else {
		goto L66
	}
L66:
	;
	v250 = int32(0)
	goto L67
L67:
	;
	v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250+v67))))
	if v281 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	v456 = v61
	goto L23
L69:
	;
	v285 = v250 << (uint(int32(2)) % 32)
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(80)+v285)))
	*(*int32)(unsafe.Add(mBase, uint32(v65+v285))) = v290
	goto L71
L70:
	;
	goto L71
L71:
	;
	v294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v250+v66))))
	if v294 == int32(1) {
		goto L73
	} else {
		goto L74
	}
L72:
	;
	v318 = v250 << (uint(int32(2)) % 32)
	v320 = *(*int32)(unsafe.Add(mBase, uint32(v65+v318)))
	if v316 < v320 {
		goto L25
	} else {
		goto L76
	}
L73:
	;
	v300 = *(*int32)(unsafe.Add(mBase, uint32(v63+v250<<(uint(int32(2))%32))))
	v316 = v300
	goto L72
L74:
	;
	goto L75
L75:
	;
	v302 = v250 << (uint(int32(2)) % 32)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(112)+v302)))
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(80)+v302)))
	v314 = v307 + v311 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63+v302))) = v314
	v316 = v314
	goto L72
L76:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(80)+v318)))
	if v320 < v325 {
		goto L24
	} else {
		goto L77
	}
L77:
	;
	v330 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(112)+v318)))
	if v330+v325 <= v316 {
		goto L24
	} else {
		goto L78
	}
L78:
	;
	v334 = v250 + int32(1)
	if v334 != v61 {
		v250 = v334
		goto L67
	} else {
		goto L79
	}
L79:
	;
	goto L68
L80:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v342 = m.ExcPending
	if v342 != 0 {
		goto L11
	} else {
		goto L81
	}
L81:
	;
	F_errmsg(m, int32(466401), int32(0))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L11
	} else {
		goto L82
	}
L82:
	;
	F_errfinish(m, int32(518800), int32(2855), int32(438697))
	mBase = m.M
	v351 = m.ExcPending
	if v351 != 0 {
		goto L11
	} else {
		goto L83
	}
L83:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L84:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v358 = m.ExcPending
	if v358 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	F_errmsg(m, int32(178715), int32(0))
	mBase = m.M
	v362 = m.ExcPending
	if v362 != 0 {
		goto L11
	} else {
		goto L86
	}
L86:
	;
	F_errdetail(m, int32(679694), int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L11
	} else {
		goto L87
	}
L87:
	;
	F_errfinish(m, int32(518800), int32(2888), int32(438697))
	mBase = m.M
	v371 = m.ExcPending
	if v371 != 0 {
		goto L11
	} else {
		goto L88
	}
L88:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L89:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v379 = m.ExcPending
	if v379 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75))) = int32(268435455)
	F_errmsg(m, int32(711032), v75)
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L11
	} else {
		goto L91
	}
L91:
	;
	F_errfinish(m, int32(518800), int32(2896), int32(438697))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L11
	} else {
		goto L92
	}
L92:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L93:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L11
	} else {
		goto L94
	}
L94:
	;
	F_errmsg(m, int32(319312), int32(0))
	mBase = m.M
	v400 = m.ExcPending
	if v400 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	F_errfinish(m, int32(518800), int32(2905), int32(438697))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L97:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v412 = m.ExcPending
	if v412 != 0 {
		goto L11
	} else {
		goto L98
	}
L98:
	;
	F_errmsg(m, int32(126195), int32(0))
	mBase = m.M
	v416 = m.ExcPending
	if v416 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	F_errfinish(m, int32(518800), int32(2915), int32(438697))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L11
	} else {
		goto L100
	}
L100:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L101:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v428 = m.ExcPending
	if v428 != 0 {
		goto L11
	} else {
		goto L102
	}
L102:
	;
	F_errmsg(m, int32(445049), int32(0))
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L11
	} else {
		goto L103
	}
L103:
	;
	F_errfinish(m, int32(518800), int32(2985), int32(438697))
	mBase = m.M
	v437 = m.ExcPending
	if v437 != 0 {
		goto L11
	} else {
		goto L104
	}
L104:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L105:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v444 = m.ExcPending
	if v444 != 0 {
		goto L11
	} else {
		goto L106
	}
L106:
	;
	F_errmsg(m, int32(421238), int32(0))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L11
	} else {
		goto L107
	}
L107:
	;
	F_errfinish(m, int32(518800), int32(2990), int32(438697))
	mBase = m.M
	v453 = m.ExcPending
	if v453 != 0 {
		goto L11
	} else {
		goto L108
	}
L108:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L109:
	;
	v490 = v456
	goto L110
L110:
	;
	v522 = v490 << (uint(int32(2)) % 32)
	v523 = v65 + v522
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(80)+v522)))
	*(*int32)(unsafe.Add(mBase, uint32(v523))) = v527
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v75+int32(112)+v522)))
	v536 = v527 + v533 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v522+v63))) = v536
	v538 = *(*int32)(unsafe.Add(mBase, uint32(v523)))
	if v536 < v538 {
		goto L112
	} else {
		goto L113
	}
L111:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L11
	} else {
		goto L115
	}
L112:
	;
	goto L111
L113:
	;
	v541 = v490 + int32(1)
	if v83 != v541 {
		v490 = v541
		goto L110
	} else {
		goto L114
	}
L114:
	;
	v631 = v240
	v635 = v4
	goto L21
L115:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v549 = m.ExcPending
	if v549 != 0 {
		goto L11
	} else {
		goto L116
	}
L116:
	;
	F_errmsg(m, int32(445049), int32(0))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L11
	} else {
		goto L117
	}
L117:
	;
	F_errfinish(m, int32(518800), int32(3000), int32(438697))
	mBase = m.M
	v558 = m.ExcPending
	if v558 != 0 {
		goto L11
	} else {
		goto L118
	}
L118:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L119:
	;
	v562 = *(*int32)(unsafe.Add(mBase, uint32(v75)+80))
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v562
	goto L121
L120:
	;
	goto L121
L121:
	;
	v564 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v66))))
	if v564 == int32(1) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v574 < v575 {
		goto L20
	} else {
		goto L126
	}
L123:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	v574 = v567
	goto L122
L124:
	;
	goto L125
L125:
	;
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v75)+112))
	v569 = *(*int32)(unsafe.Add(mBase, uint32(v75)+80))
	v572 = v568 + v569 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v572
	v574 = v572
	goto L122
L126:
	;
	v577 = *(*int32)(unsafe.Add(mBase, uint32(v75)+80))
	if v577 <= v575 {
		goto L128
	} else {
		goto L129
	}
L127:
	;
	v600 = v596 + v597
	if v574 < v600 {
		v631 = v598
		v635 = v599
		goto L21
	} else {
		goto L133
	}
L128:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v75)+112))
	v596 = v579
	v597 = v577
	v598 = v240
	v599 = v4
	goto L127
L129:
	;
	goto L130
L130:
	;
	v580 = v577 - v575
	if base.B2i32(v580 < v577)^base.B2i32(int32(0) < v575) != 0 {
		goto L19
	} else {
		goto L131
	}
L131:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v75)+112))
	v586 = v585 + v580
	*(*int32)(unsafe.Add(mBase, uint32(v75)+112)) = v586
	if base.B2i32(v580 < int32(0)) != base.B2i32(v586 < v585) {
		goto L19
	} else {
		goto L132
	}
L132:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+80)) = v575
	v596 = v586
	v597 = v575
	v598 = base.B2i32(int32(1) < v580) | v240
	v599 = v580
	goto L127
L133:
	;
	v604 = v574 - v600
	if base.B2i32(int32(0) < v600)^base.B2i32(v604 < v574) != 0 {
		goto L18
	} else {
		goto L134
	}
L134:
	;
	v608 = v604 + int32(1)
	if v608 < v604 {
		goto L18
	} else {
		goto L135
	}
L135:
	;
	v610 = v596 + v608
	*(*int32)(unsafe.Add(mBase, uint32(v75)+112)) = v610
	if base.B2i32(v608 < int32(0)) != base.B2i32(v610 < v596) {
		goto L18
	} else {
		goto L136
	}
L136:
	;
	v631 = base.B2i32(int32(1) < v608) | v598
	v635 = v599
	goto L21
L137:
	;
	F_ArrayCheckBounds(m, v83, v75+int32(112), v75+int32(80))
	mBase = m.M
	v661 = m.ExcPending
	if v661 != 0 {
		goto L11
	} else {
		goto L138
	}
L138:
	;
	v663 = v75 + int32(48)
	v664 = int32(0)
	if v83 <= v664 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v742 = F_ArrayGetNItems(m, v83, v75+int32(48))
	mBase = m.M
	v743 = m.ExcPending
	if v743 != 0 {
		goto L11
	} else {
		goto L149
	}
L140:
	;
	goto L139
L141:
	;
	v670 = int32(1)
	if v83 != v670 {
		goto L142
	} else {
		goto L143
	}
L142:
	;
	v680 = v664
	v681 = v664
	goto L145
L143:
	;
	v715 = v664
	goto L144
L144:
	;
	if v83&v670 == int32(0) {
		goto L140
	} else {
		goto L148
	}
L145:
	;
	v684 = int32(2)
	v685 = v680 << (uint(v684) % 32)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v685+v63)))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v685+v65)))
	v692 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v663+v685))) = v688 - v690 + v692
	v696 = v685 | int32(4)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v696+v63)))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(v696+v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v663+v696))) = v699 - v701 + v692
	v707 = v680 + v684
	v709 = v681 + v684
	if v709 != v83&int32(2147483646) {
		v680 = v707
		v681 = v709
		goto L145
	} else {
		goto L147
	}
L146:
	;
	v715 = v707
	goto L144
L147:
	;
	goto L146
L148:
	;
	v722 = v715 << (uint(int32(2)) % 32)
	v725 = *(*int32)(unsafe.Add(mBase, uint32(v722+v63)))
	v727 = *(*int32)(unsafe.Add(mBase, uint32(v722+v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v663+v722))) = v725 - v727 + int32(1)
	goto L140
L149:
	;
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v746 = v81 + int32(16)
	v747 = F_ArrayGetNItems(m, v744, v746)
	mBase = m.M
	v748 = m.ExcPending
	if v748 != 0 {
		goto L11
	} else {
		goto L150
	}
L150:
	;
	if v747 < v742 {
		goto L17
	} else {
		goto L151
	}
L151:
	;
	v751 = v83 << (uint(int32(3)) % 32)
	if v631&int32(1) != 0 {
		goto L153
	} else {
		goto L154
	}
L152:
	;
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v773 = v771 << (uint(int32(3)) % 32)
	if v770 != 0 {
		goto L156
	} else {
		goto L157
	}
L153:
	;
	v757 = base.I32_div_s(v654+int32(7), int32(8))
	v762 = (v751 + v757 + int32(23)) & int32(-8)
	v768 = v762
	v769 = v762
	goto L152
L154:
	;
	goto L155
L155:
	;
	v768 = int32(0)
	v769 = (v751 + int32(23)) & int32(120)
	goto L152
L156:
	;
	v778 = v770
	goto L158
L157:
	;
	v778 = (v773 + int32(23)) & int32(-8)
	goto L158
L158:
	;
	v779 = v81 + v778
	v780 = int32(0)
	if v770 != 0 {
		goto L159
	} else {
		goto L160
	}
L159:
	;
	v783 = v773 + v746
	goto L161
L160:
	;
	v783 = v780
	goto L161
L161:
	;
	v784 = F_array_seek(m, v779, v780, v783, v742, v70, v72)
	mBase = m.M
	v785 = v784 - v779
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v786 == int32(0) {
		goto L163
	} else {
		goto L164
	}
L162:
	;
	v822 = F_array_slice_size(m, v79+v815, v814, v83, v75+int32(112), v75+int32(80), v65, v63, v70, v72)
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L11
	} else {
		goto L170
	}
L163:
	;
	v789 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v798 = (v792<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	v799 = int32(base.Ui32(v789)>>(uint(int32(2))%32)) - v798
	if int32(1) < v83 {
		v813 = v799
		v814 = v4
		v815 = v798
		goto L162
	} else {
		goto L166
	}
L164:
	;
	goto L165
L165:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(v79)))
	v803 = int32(2)
	v805 = int32(base.Ui32(v802)>>(uint(v803)%32)) - v786
	v806 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	if v83 < v803 {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v901 = v792
	v902 = v799
	v903 = v798
	goto L16
L167:
	;
	v901 = v806
	v902 = v805
	v903 = v786
	goto L16
L168:
	;
	goto L169
L169:
	;
	v813 = v805
	v814 = v220 + v806<<(uint(int32(3))%32)
	v815 = v786
	goto L162
L170:
	;
	v824 = int32(0)
	v952 = v822
	v953 = v824
	v954 = v824
	v955 = v813
	v958 = v815
	v959 = int32(1)
	v960 = v4
	v961 = v4
	v967 = v824
	goto L15
L171:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L11
	} else {
		goto L172
	}
L172:
	;
	F_errmsg(m, int32(445049), int32(0))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L11
	} else {
		goto L173
	}
L173:
	;
	F_errfinish(m, int32(518800), int32(2940), int32(438697))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L11
	} else {
		goto L174
	}
L174:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L175:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v851 = m.ExcPending
	if v851 != 0 {
		goto L11
	} else {
		goto L176
	}
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+16)) = int32(268435455)
	F_errmsg(m, int32(711032), v75+int32(16))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L11
	} else {
		goto L177
	}
L177:
	;
	F_errfinish(m, int32(518800), int32(2950), int32(438697))
	mBase = m.M
	v863 = m.ExcPending
	if v863 != 0 {
		goto L11
	} else {
		goto L178
	}
L178:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L179:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L11
	} else {
		goto L180
	}
L180:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v75)+32)) = int32(268435455)
	F_errmsg(m, int32(711032), v75+int32(32))
	mBase = m.M
	v879 = m.ExcPending
	if v879 != 0 {
		goto L11
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(518800), int32(2965), int32(438697))
	mBase = m.M
	v884 = m.ExcPending
	if v884 != 0 {
		goto L11
	} else {
		goto L182
	}
L182:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L183:
	;
	F_errcode(m, int32(352845954))
	mBase = m.M
	v891 = m.ExcPending
	if v891 != 0 {
		goto L11
	} else {
		goto L184
	}
L184:
	;
	F_errmsg(m, int32(319312), int32(0))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L11
	} else {
		goto L185
	}
L185:
	;
	F_errfinish(m, int32(518800), int32(3017), int32(438697))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L11
	} else {
		goto L186
	}
L186:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L187:
	;
	v910 = v786
	goto L189
L188:
	;
	v910 = (v905 + int32(23)) & int32(-8)
	goto L189
L189:
	;
	v911 = v79 + v910
	v912 = int32(0)
	if v786 != 0 {
		goto L190
	} else {
		goto L191
	}
L190:
	;
	v915 = v905 + v220
	goto L192
L191:
	;
	v915 = v912
	goto L192
L192:
	;
	v919 = *(*int32)(unsafe.Add(mBase, uint32(v220+v901<<(uint(int32(2))%32))))
	v920 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	if v920 < v919 {
		goto L193
	} else {
		goto L194
	}
L193:
	;
	v922 = v919
	goto L195
L194:
	;
	v922 = v920
	goto L195
L195:
	;
	v923 = *(*int32)(unsafe.Add(mBase, uint32(v79)+16))
	v924 = v923 + v919
	if v922 < v924 {
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v926 = v922
	goto L198
L197:
	;
	v926 = v924
	goto L198
L198:
	;
	v927 = v926 - v919
	v928 = F_array_seek(m, v911, v912, v915, v927, v70, v72)
	mBase = m.M
	v929 = v928 - v911
	v932 = v924 - int32(1)
	v933 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
	if v932 < v933 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	v935 = v932
	goto L201
L200:
	;
	v935 = v933
	goto L201
L201:
	;
	if v922 <= v935 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	v940 = v935 - v922 + int32(1)
	v941 = F_array_seek(m, v929+v911, v927, v915, v940, v70, v72)
	mBase = m.M
	v943 = v941 - v928
	v944 = v940
	goto L204
L203:
	;
	v943 = int32(0)
	v944 = v4
	goto L204
L204:
	;
	v946 = v935 + int32(1)
	if v919 < v946 {
		goto L205
	} else {
		goto L206
	}
L205:
	;
	v948 = v946
	goto L207
L206:
	;
	v948 = v919
	goto L207
L207:
	;
	v952 = v943
	v953 = v929
	v954 = v924 - v948
	v955 = v902
	v958 = v903
	v959 = v4
	v960 = v927
	v961 = v944
	v967 = v902 - (v943 + v929)
	goto L15
L208:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v971)+8)) = v768
	*(*int32)(unsafe.Add(mBase, uint32(v971)+4)) = v83
	*(*int32)(unsafe.Add(mBase, uint32(v971))) = v970 << (uint(int32(2)) % 32)
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v79)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v971)+12)) = v978
	v981 = v971 + int32(16)
	if v222 != 0 {
		goto L210
	} else {
		goto L211
	}
L209:
	;
	if v222 != 0 {
		goto L214
	} else {
		goto L215
	}
L210:
	;
	v984 = F__emscripten_memcpy_bulkmem(m, v981, v75+int32(112), v222)
	mBase = m.M
	v985 = v984
	goto L212
L211:
	;
	v985 = v981
	goto L212
L212:
	;
	goto L209
L213:
	;
	if v959 != 0 {
		goto L219
	} else {
		goto L220
	}
L214:
	;
	v989 = F__emscripten_memcpy_bulkmem(m, v985+v222, v75+int32(80), v222)
	mBase = m.M
	goto L216
L215:
	;
	goto L216
L216:
	;
	goto L213
L217:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2879))) = uint8(v2873)
	v2920 = v971
	goto L13
L218:
	;
	v2791 = base.I32_div_s(v2005, int32(8))
	v2792 = v1035 + v2791
	v2793 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2792))))
	v2796 = v2090
	v2797 = v2086
	v2798 = v2074
	v2799 = int32(1) << (uint(v2005&int32(7)) % 32)
	v2802 = v2089
	v2804 = v2793
	v2805 = v2792
	goto L529
L219:
	;
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v971)+8))
	if v991 == int32(0) {
		goto L222
	} else {
		goto L223
	}
L220:
	;
	goto L221
L221:
	;
	v2140 = v769 + v971
	v2141 = v79 + v958
	if v953 != 0 {
		goto L419
	} else {
		goto L420
	}
L222:
	;
	v994 = *(*int32)(unsafe.Add(mBase, uint32(v971)+4))
	v1001 = (v994<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L224
L223:
	;
	v1001 = v991
	goto L224
L224:
	;
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v1002 == int32(0) {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v1005 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v1012 = (v1005<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L227
L226:
	;
	v1012 = v1002
	goto L227
L227:
	;
	v1013 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v1013 == int32(0) {
		goto L228
	} else {
		goto L229
	}
L228:
	;
	v1016 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v1023 = (v1016<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L230
L229:
	;
	v1023 = v1013
	goto L230
L230:
	;
	v1024 = int32(0)
	if v991 != 0 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v1026 = *(*int32)(unsafe.Add(mBase, uint32(v971)+4))
	v1030 = v985 + v1026<<(uint(int32(3))%32)
	goto L233
L232:
	;
	v1030 = v1024
	goto L233
L233:
	;
	if v1002 != 0 {
		goto L234
	} else {
		goto L235
	}
L234:
	;
	v1031 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v1035 = v220 + v1031<<(uint(int32(3))%32)
	goto L236
L235:
	;
	v1035 = v1024
	goto L236
L236:
	;
	v1036 = v1001 + v971
	v1037 = v1012 + v79
	if v1013 != 0 {
		goto L237
	} else {
		goto L238
	}
L237:
	;
	v1039 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v1043 = v746 + v1039<<(uint(int32(3))%32)
	goto L239
L238:
	;
	v1043 = int32(0)
	goto L239
L239:
	;
	v1044 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v1045 = F_ArrayGetNItems(m, v1044, v220)
	mBase = m.M
	v1046 = m.ExcPending
	if v1046 != 0 {
		goto L11
	} else {
		goto L240
	}
L240:
	;
	v1047 = int32(0)
	v1049 = v75 + int32(112)
	v1051 = v75 + int32(80)
	v1061 = v83 - int32(1)
	if v1061 < v1047 {
		v1139 = v1047
		goto L242
	} else {
		goto L243
	}
L241:
	;
	v1147 = F_array_seek(m, v1037, v1047, v1035, v1139, v70, v72)
	mBase = m.M
	v1148 = v1147 - v1037
	if v1148 != 0 {
		goto L253
	} else {
		goto L254
	}
L242:
	;
	goto L241
L243:
	;
	v1064 = int32(1)
	if v1061 == int32(0) {
		goto L245
	} else {
		goto L246
	}
L244:
	;
	if v83&v1064 == int32(0) {
		v1139 = v1116
		goto L242
	} else {
		goto L251
	}
L245:
	;
	v1116 = v1047
	v1117 = v1061
	v1118 = v1064
	goto L244
L246:
	;
	goto L247
L247:
	;
	v1075 = v1047
	v1076 = v1061
	v1077 = v1064
	v1078 = v1047
	goto L248
L248:
	;
	v1083 = int32(2)
	v1084 = v1076 << (uint(v1083) % 32)
	v1086 = v1084 - int32(4)
	v1088 = *(*int32)(unsafe.Add(mBase, uint32(v65+v1086)))
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v1051+v1086)))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1084+v1049)))
	v1094 = v1093 * v1077
	v1097 = *(*int32)(unsafe.Add(mBase, uint32(v1084+v65)))
	v1099 = *(*int32)(unsafe.Add(mBase, uint32(v1084+v1051)))
	v1103 = (v1088-v1090)*v1094 + ((v1097-v1099)*v1077 + v1075)
	v1105 = *(*int32)(unsafe.Add(mBase, uint32(v1049+v1086)))
	v1106 = v1105 * v1094
	v1108 = v1076 - v1083
	v1110 = v1078 + v1083
	if v1110 != v83&int32(-2) {
		v1075 = v1103
		v1076 = v1108
		v1077 = v1106
		v1078 = v1110
		goto L248
	} else {
		goto L250
	}
L249:
	;
	v1116 = v1103
	v1117 = v1108
	v1118 = v1106
	goto L244
L250:
	;
	goto L249
L251:
	;
	v1127 = v1117 << (uint(int32(2)) % 32)
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v65+v1127)))
	v1131 = *(*int32)(unsafe.Add(mBase, uint32(v1127+v1051)))
	v1139 = (v1129-v1131)*v1118 + v1116
	goto L242
L252:
	;
	if v1030 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L253:
	;
	v1149 = F__emscripten_memcpy_bulkmem(m, v1036, v1037, v1148)
	mBase = m.M
	v1150 = v1149
	goto L255
L254:
	;
	v1150 = v1036
	goto L255
L255:
	;
	goto L252
L256:
	;
	v1390 = v75 + int32(112)
	v1392 = v75 + int32(240)
	v1396 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1392+v83<<(uint(v1396)%32)-int32(4)))) = int32(1)
	v1404 = v83 - v1396
	if v1404 < int32(0) {
		goto L298
	} else {
		goto L299
	}
L257:
	;
	if v1139 <= int32(0) {
		goto L256
	} else {
		goto L258
	}
L258:
	;
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1030))))
	if v1035 == int32(0) {
		goto L261
	} else {
		goto L262
	}
L259:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1329))) = uint8(v1324)
	goto L256
L260:
	;
	v1244 = v1139
	v1246 = v1155
	v1251 = v1030
	goto L287
L261:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1139) {
		goto L260
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	v1162 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1035))))
	v1163 = int32(1)
	v1167 = v1163
	v1169 = v1155
	v1173 = v1163
	v1174 = v1030
	v1175 = v1139
	v1185 = v1162
	v1189 = v1035
	goto L265
L264:
	;
	v1324 = v1155 | int32(1)
	v1329 = v1030
	goto L259
L265:
	;
	if v1173&v1185 != 0 {
		goto L267
	} else {
		goto L268
	}
L266:
	;
	if v1217 != int32(1) {
		v1324 = v1218
		v1329 = v1219
		goto L259
	} else {
		goto L280
	}
L267:
	;
	v1203 = v1167 | v1169
	goto L269
L268:
	;
	v1203 = v1169 & (v1167 ^ int32(-1))
	goto L269
L269:
	;
	v1204 = int32(1)
	v1205 = v1175 - v1204
	v1207 = v1167 << (uint(v1204) % 32)
	if v1207 == int32(256) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1174))) = uint8(v1203)
	if v1205 == int32(0) {
		goto L256
	} else {
		goto L273
	}
L271:
	;
	v1217 = v1207
	v1218 = v1203
	v1219 = v1174
	goto L272
L272:
	;
	v1221 = v1173 << (uint(int32(1)) % 32)
	if v1221 == int32(256) {
		goto L275
	} else {
		goto L276
	}
L273:
	;
	v1213 = int32(1)
	v1214 = v1174 + v1213
	v1215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1214))))
	v1217 = v1213
	v1218 = v1215
	v1219 = v1214
	goto L272
L274:
	;
	goto L266
L275:
	;
	if v1205 == int32(0) {
		goto L274
	} else {
		goto L278
	}
L276:
	;
	v1230 = v1221
	v1231 = v1185
	v1232 = v1189
	goto L277
L277:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1175) {
		v1167 = v1217
		v1169 = v1218
		v1173 = v1230
		v1174 = v1219
		v1175 = v1205
		v1185 = v1231
		v1189 = v1232
		goto L265
	} else {
		goto L279
	}
L278:
	;
	v1226 = int32(1)
	v1227 = v1189 + v1226
	v1228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1227))))
	v1230 = v1226
	v1231 = v1228
	v1232 = v1227
	goto L277
L279:
	;
	goto L274
L280:
	;
	goto L256
L281:
	;
	v1324 = v1246 | int32(3)
	v1329 = v1251
	goto L259
L282:
	;
	v1324 = v1246 | int32(7)
	v1329 = v1251
	goto L259
L283:
	;
	v1324 = v1246 | int32(15)
	v1329 = v1251
	goto L259
L284:
	;
	v1324 = v1246 | int32(31)
	v1329 = v1251
	goto L259
L285:
	;
	v1324 = v1246 | int32(63)
	v1329 = v1251
	goto L259
L286:
	;
	v1324 = v1246 | int32(127)
	v1329 = v1251
	goto L259
L287:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v1244-int32(3)) {
		goto L281
	} else {
		goto L289
	}
L288:
	;
	v1324 = v1303 | int32(1)
	v1329 = v1302
	goto L259
L289:
	;
	v1280 = v1244 & int32(-2)
	if v1280 == int32(2) {
		goto L282
	} else {
		goto L290
	}
L290:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v1244-int32(5)) {
		goto L283
	} else {
		goto L291
	}
L291:
	;
	if v1280 == int32(4) {
		goto L284
	} else {
		goto L292
	}
L292:
	;
	if base.Ui32(int32(-3)) < base.Ui32(v1244-int32(7)) {
		goto L285
	} else {
		goto L293
	}
L293:
	;
	if v1280 == int32(6) {
		goto L286
	} else {
		goto L294
	}
L294:
	;
	v1295 = int32(255)
	*(*uint8)(unsafe.Add(mBase, uint32(v1251))) = uint8(v1295)
	v1298 = v1244 - int32(8)
	if v1298 == int32(0) {
		goto L256
	} else {
		goto L295
	}
L295:
	;
	v1301 = int32(1)
	v1302 = v1251 + v1301
	v1303 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1302))))
	if v1298 != v1301 {
		v1244 = v1298
		v1246 = v1303
		v1251 = v1302
		goto L287
	} else {
		goto L296
	}
L296:
	;
	goto L288
L297:
	;
	v1464 = v75 + int32(208)
	v1465 = int32(0)
	if v83 <= v1465 {
		goto L308
	} else {
		goto L309
	}
L298:
	;
	goto L297
L299:
	;
	if v83&int32(1) == int32(0) {
		goto L300
	} else {
		goto L301
	}
L300:
	;
	v1411 = int32(2)
	v1417 = v83<<(uint(v1411)%32) - int32(4)
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(v1390+v1417)))
	v1421 = *(*int32)(unsafe.Add(mBase, uint32(v1392+v1417)))
	*(*int32)(unsafe.Add(mBase, uint32(v1392+v1404<<(uint(v1411)%32)))) = v1419 * v1421
	v1426 = v83 - int32(3)
	goto L302
L301:
	;
	v1426 = v1404
	goto L302
L302:
	;
	if v1404 == int32(0) {
		goto L298
	} else {
		goto L303
	}
L303:
	;
	v1432 = v1426
	goto L304
L304:
	;
	v1435 = int32(2)
	v1436 = v1432 << (uint(v1435) % 32)
	v1439 = v1436 + int32(4)
	v1441 = *(*int32)(unsafe.Add(mBase, uint32(v1390+v1439)))
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1392+v1439)))
	v1444 = v1441 * v1443
	*(*int32)(unsafe.Add(mBase, uint32(v1392+v1436))) = v1444
	v1447 = v1432 - int32(1)
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v1436+v1390)))
	*(*int32)(unsafe.Add(mBase, uint32(v1392+v1447<<(uint(v1435)%32)))) = v1452 * v1444
	if v1447 != 0 {
		v1432 = v1432 - v1435
		goto L304
	} else {
		goto L306
	}
L305:
	;
	goto L298
L306:
	;
	goto L305
L307:
	;
	v1542 = v75 + int32(176)
	v1544 = v75 + int32(240)
	v1546 = v75 + int32(208)
	v1547 = int32(0)
	v1553 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v1542+v83<<(uint(v1553)%32)-int32(4)))) = v1547
	v1561 = v83 - v1553
	if v1547 <= v1561 {
		goto L318
	} else {
		goto L319
	}
L308:
	;
	goto L307
L309:
	;
	v1471 = int32(1)
	if v83 != v1471 {
		goto L310
	} else {
		goto L311
	}
L310:
	;
	v1481 = v1465
	v1482 = v1465
	goto L313
L311:
	;
	v1516 = v1465
	goto L312
L312:
	;
	if v83&v1471 == int32(0) {
		goto L308
	} else {
		goto L316
	}
L313:
	;
	v1485 = int32(2)
	v1486 = v1481 << (uint(v1485) % 32)
	v1489 = *(*int32)(unsafe.Add(mBase, uint32(v1486+v63)))
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1486+v65)))
	v1493 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v1464+v1486))) = v1489 - v1491 + v1493
	v1497 = v1486 | int32(4)
	v1500 = *(*int32)(unsafe.Add(mBase, uint32(v1497+v63)))
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v1497+v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v1464+v1497))) = v1500 - v1502 + v1493
	v1508 = v1481 + v1485
	v1510 = v1482 + v1485
	if v1510 != v83&int32(2147483646) {
		v1481 = v1508
		v1482 = v1510
		goto L313
	} else {
		goto L315
	}
L314:
	;
	v1516 = v1508
	goto L312
L315:
	;
	goto L314
L316:
	;
	v1523 = v1516 << (uint(int32(2)) % 32)
	v1526 = *(*int32)(unsafe.Add(mBase, uint32(v1523+v63)))
	v1528 = *(*int32)(unsafe.Add(mBase, uint32(v1523+v65)))
	*(*int32)(unsafe.Add(mBase, uint32(v1464+v1523))) = v1526 - v1528 + int32(1)
	goto L308
L317:
	;
	v1667 = int32(0)
	v1672 = F__emscripten_memset_bulkmem(m, v75+int32(144), base.I32_extend8_s(v1667), v222)
	mBase = m.M
	goto L333
L318:
	;
	v1568 = v1561
	v1570 = v1547
	goto L321
L319:
	;
	goto L320
L320:
	;
	goto L317
L321:
	;
	v1575 = v1568 << (uint(int32(2)) % 32)
	v1576 = v1542 + v1575
	v1578 = *(*int32)(unsafe.Add(mBase, uint32(v1544+v1575)))
	v1579 = int32(1)
	v1580 = v1578 - v1579
	*(*int32)(unsafe.Add(mBase, uint32(v1576))) = v1580
	v1583 = v1568 + v1579
	if v83 <= v1583 {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	goto L320
L323:
	;
	v1651 = int32(1)
	if int32(0) < v1568 {
		v1568 = v1568 - v1651
		v1570 = v1570 + v1651
		goto L321
	} else {
		goto L332
	}
L324:
	;
	if v1570&int32(1) == int32(0) {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1589 = int32(2)
	v1590 = v1583 << (uint(v1589) % 32)
	v1592 = *(*int32)(unsafe.Add(mBase, uint32(v1546+v1590)))
	v1596 = *(*int32)(unsafe.Add(mBase, uint32(v1544+v1590)))
	v1598 = v1580 - (v1592-int32(1))*v1596
	*(*int32)(unsafe.Add(mBase, uint32(v1576))) = v1598
	v1602 = v1568 + v1589
	v1603 = v1598
	goto L327
L326:
	;
	v1602 = v1583
	v1603 = v1580
	goto L327
L327:
	;
	if v1570 == int32(0) {
		goto L323
	} else {
		goto L328
	}
L328:
	;
	v1610 = v1602
	v1611 = v1603
	goto L329
L329:
	;
	v1616 = int32(2)
	v1617 = v1610 << (uint(v1616) % 32)
	v1619 = *(*int32)(unsafe.Add(mBase, uint32(v1546+v1617)))
	v1620 = int32(1)
	v1623 = *(*int32)(unsafe.Add(mBase, uint32(v1544+v1617)))
	v1625 = v1611 - (v1619-v1620)*v1623
	*(*int32)(unsafe.Add(mBase, uint32(v1576))) = v1625
	v1628 = v1617 + int32(4)
	v1630 = *(*int32)(unsafe.Add(mBase, uint32(v1546+v1628)))
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(v1544+v1628)))
	v1636 = v1625 - (v1630-v1620)*v1634
	*(*int32)(unsafe.Add(mBase, uint32(v1576))) = v1636
	v1639 = v1610 + v1616
	if v1639 != v83 {
		v1610 = v1639
		v1611 = v1636
		goto L329
	} else {
		goto L331
	}
L330:
	;
	goto L323
L331:
	;
	goto L330
L332:
	;
	goto L322
L333:
	;
	v1675 = v1139
	v1677 = v1148 + v1150
	v1678 = v1147
	v1679 = v83 - int32(1)
	v1683 = v81 + v1023
	v1684 = v1667
	v1688 = v1139
	goto L334
L334:
	;
	v1712 = v75 + int32(176) + v1679<<(uint(int32(2))%32)
	v1713 = *(*int32)(unsafe.Add(mBase, uint32(v1712)))
	if v1713 == int32(0) {
		goto L337
	} else {
		goto L338
	}
L335:
	;
	v2074 = v1045 - v2005
	v2075 = F_array_seek(m, v2013, v2005, v1035, v2074, v70, v72)
	mBase = m.M
	v2076 = v2075 - v2013
	if v2076 != 0 {
		goto L405
	} else {
		goto L406
	}
L336:
	;
	v1972 = F_array_seek(m, v1683, v1684, v1043, int32(1), v70, v72)
	mBase = m.M
	v1973 = v1972 - v1683
	if v1973 != 0 {
		goto L375
	} else {
		goto L376
	}
L337:
	;
	v1938 = v1675
	v1940 = v1677
	v1941 = v1688
	v1962 = v1678
	goto L336
L338:
	;
	goto L339
L339:
	;
	v1716 = F_array_seek(m, v1678, v1688, v1035, v1713, v70, v72)
	mBase = m.M
	v1717 = v1716 - v1678
	if v1717 != 0 {
		goto L341
	} else {
		goto L342
	}
L340:
	;
	if v1030 == int32(0) {
		goto L344
	} else {
		goto L345
	}
L341:
	;
	v1718 = F__emscripten_memcpy_bulkmem(m, v1677, v1678, v1717)
	mBase = m.M
	v1719 = v1718
	goto L343
L342:
	;
	v1719 = v1677
	goto L343
L343:
	;
	goto L340
L344:
	;
	v1935 = *(*int32)(unsafe.Add(mBase, uint32(v1712)))
	v1938 = v1675 + v1935
	v1940 = v1719 + v1717
	v1941 = v1935 + v1688
	v1962 = v1716
	goto L336
L345:
	;
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1712)))
	if v1722 <= int32(0) {
		goto L344
	} else {
		goto L346
	}
L346:
	;
	v1728 = int32(1) << (uint(v1675&int32(7)) % 32)
	v1730 = base.I32_div_s(v1675, int32(8))
	v1731 = v1030 + v1730
	v1732 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1731))))
	if v1035 != 0 {
		goto L348
	} else {
		goto L349
	}
L347:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1878))) = uint8(v1871)
	goto L344
L348:
	;
	v1738 = base.I32_div_s(v1688, int32(8))
	v1739 = v1035 + v1738
	v1740 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1739))))
	v1743 = v1722
	v1744 = v1728
	v1745 = v1732
	v1746 = int32(1) << (uint(v1688&int32(7)) % 32)
	v1752 = v1731
	v1755 = v1740
	v1763 = v1739
	goto L351
L349:
	;
	goto L350
L350:
	;
	v1820 = v1722
	v1821 = v1728
	v1822 = v1732
	v1829 = v1731
	goto L367
L351:
	;
	if v1746&v1755 != 0 {
		goto L353
	} else {
		goto L354
	}
L352:
	;
	if v1793 != int32(1) {
		v1871 = v1794
		v1878 = v1795
		goto L347
	} else {
		goto L366
	}
L353:
	;
	v1779 = v1744 | v1745
	goto L355
L354:
	;
	v1779 = v1745 & (v1744 ^ int32(-1))
	goto L355
L355:
	;
	v1780 = int32(1)
	v1781 = v1743 - v1780
	v1783 = v1744 << (uint(v1780) % 32)
	if v1783 == int32(256) {
		goto L356
	} else {
		goto L357
	}
L356:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1752))) = uint8(v1779)
	if v1781 == int32(0) {
		goto L344
	} else {
		goto L359
	}
L357:
	;
	v1793 = v1783
	v1794 = v1779
	v1795 = v1752
	goto L358
L358:
	;
	v1797 = v1746 << (uint(int32(1)) % 32)
	if v1797 == int32(256) {
		goto L361
	} else {
		goto L362
	}
L359:
	;
	v1789 = int32(1)
	v1791 = v1752 + v1789
	v1792 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1791))))
	v1793 = v1789
	v1794 = v1792
	v1795 = v1791
	goto L358
L360:
	;
	goto L352
L361:
	;
	if v1781 == int32(0) {
		goto L360
	} else {
		goto L364
	}
L362:
	;
	v1806 = v1797
	v1807 = v1755
	v1808 = v1763
	goto L363
L363:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1743) {
		v1743 = v1781
		v1744 = v1793
		v1745 = v1794
		v1746 = v1806
		v1752 = v1795
		v1755 = v1807
		v1763 = v1808
		goto L351
	} else {
		goto L365
	}
L364:
	;
	v1802 = int32(1)
	v1803 = v1763 + v1802
	v1804 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1803))))
	v1806 = v1802
	v1807 = v1804
	v1808 = v1803
	goto L363
L365:
	;
	goto L360
L366:
	;
	goto L344
L367:
	;
	v1851 = v1821 | v1822
	v1852 = int32(1)
	v1853 = v1820 - v1852
	v1855 = v1821 << (uint(v1852) % 32)
	if v1855 == int32(256) {
		goto L369
	} else {
		goto L370
	}
L368:
	;
	v1871 = v1851
	v1878 = v1829
	goto L347
L369:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1829))) = uint8(v1851)
	if v1853 == int32(0) {
		goto L344
	} else {
		goto L372
	}
L370:
	;
	goto L371
L371:
	;
	if base.Ui32(int32(1)) < base.Ui32(v1820) {
		v1820 = v1853
		v1821 = v1855
		v1822 = v1851
		goto L367
	} else {
		goto L373
	}
L372:
	;
	v1861 = int32(1)
	v1863 = v1829 + v1861
	v1864 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1863))))
	v1820 = v1853
	v1821 = v1861
	v1822 = v1864
	v1829 = v1863
	goto L367
L373:
	;
	goto L368
L374:
	;
	if v1030 != 0 {
		goto L378
	} else {
		goto L379
	}
L375:
	;
	v1974 = F__emscripten_memcpy_bulkmem(m, v1940, v1683, v1973)
	mBase = m.M
	v1975 = v1974
	goto L377
L376:
	;
	v1975 = v1940
	goto L377
L377:
	;
	goto L374
L378:
	;
	v1979 = int32(1) << (uint(v1938&int32(7)) % 32)
	v1981 = base.I32_div_s(v1938, int32(8))
	v1982 = v1030 + v1981
	v1983 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1982))))
	if v1043 != 0 {
		goto L382
	} else {
		goto L383
	}
L379:
	;
	goto L380
L380:
	;
	v2004 = int32(1)
	v2005 = v1941 + v2004
	v2009 = v1938 + v2004
	v2011 = v1975 + v1973
	v2013 = F_array_seek(m, v1962, v1941, v1035, v2004, v70, v72)
	mBase = m.M
	v2015 = v75 + int32(144)
	v2017 = v75 + int32(208)
	if v83 <= int32(0) {
		goto L389
	} else {
		goto L390
	}
L381:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1982))) = uint8(v1999)
	goto L380
L382:
	;
	v1989 = base.I32_div_s(v1684, int32(8))
	v1991 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1043+v1989))))
	if int32(base.Ui32(v1991)>>(uint(v1684&int32(7))%32))&int32(1) != 0 {
		goto L385
	} else {
		goto L386
	}
L383:
	;
	goto L384
L384:
	;
	v1999 = v1979 | v1983
	goto L381
L385:
	;
	v1997 = v1979 | v1983
	goto L387
L386:
	;
	v1997 = v1983 & (v1979 ^ int32(-1))
	goto L387
L387:
	;
	v1999 = v1997
	goto L381
L388:
	;
	if v2071 != int32(-1) {
		v1675 = v2009
		v1677 = v2011
		v1678 = v2013
		v1679 = v2071
		v1683 = v1973 + v1683
		v1684 = v1684 + v2004
		v1688 = v2005
		goto L334
	} else {
		goto L403
	}
L389:
	;
	v2071 = int32(-1)
	goto L388
L390:
	;
	goto L391
L391:
	;
	v2023 = int32(1)
	v2024 = v83 - v2023
	v2026 = v2024 << (uint(int32(2)) % 32)
	v2027 = v2015 + v2026
	v2028 = *(*int32)(unsafe.Add(mBase, uint32(v2027)))
	v2032 = *(*int32)(unsafe.Add(mBase, uint32(v2017+v2026)))
	v2033 = base.I32_rem_s(v2028+v2023, v2032)
	*(*int32)(unsafe.Add(mBase, uint32(v2027))) = v2033
	if v2024 != 0 {
		goto L393
	} else {
		goto L394
	}
L392:
	;
	v2071 = v2061
	goto L388
L393:
	;
	v2035 = v2024
	v2038 = v2033
	goto L396
L394:
	;
	goto L395
L395:
	;
	v2059 = *(*int32)(unsafe.Add(mBase, uint32(v2015)))
	if v2059 != 0 {
		goto L400
	} else {
		goto L401
	}
L396:
	;
	if v2038 != 0 {
		v2061 = v2035
		goto L392
	} else {
		goto L398
	}
L397:
	;
	goto L395
L398:
	;
	v2040 = int32(1)
	v2041 = v2035 - v2040
	v2043 = v2041 << (uint(int32(2)) % 32)
	v2044 = v2015 + v2043
	v2045 = *(*int32)(unsafe.Add(mBase, uint32(v2044)))
	v2049 = *(*int32)(unsafe.Add(mBase, uint32(v2017+v2043)))
	v2050 = base.I32_rem_s(v2045+v2040, v2049)
	*(*int32)(unsafe.Add(mBase, uint32(v2044))) = v2050
	if v2041 != 0 {
		v2035 = v2041
		v2038 = v2050
		goto L396
	} else {
		goto L399
	}
L399:
	;
	goto L397
L400:
	;
	v2060 = int32(0)
	goto L402
L401:
	;
	v2060 = int32(-1)
	goto L402
L402:
	;
	v2061 = v2060
	goto L392
L403:
	;
	goto L335
L404:
	;
	if v1030 == int32(0) {
		v2920 = v971
		goto L13
	} else {
		goto L408
	}
L405:
	;
	v2077 = F__emscripten_memcpy_bulkmem(m, v2011, v2013, v2076)
	mBase = m.M
	goto L407
L406:
	;
	goto L407
L407:
	;
	goto L404
L408:
	;
	if v2074 <= int32(0) {
		v2920 = v971
		goto L13
	} else {
		goto L409
	}
L409:
	;
	v2086 = int32(1) << (uint(v2009&int32(7)) % 32)
	v2088 = base.I32_div_s(v2009, int32(8))
	v2089 = v1030 + v2088
	v2090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2089))))
	if v1035 != 0 {
		goto L218
	} else {
		goto L410
	}
L410:
	;
	v2093 = v2090
	v2094 = v2086
	v2095 = v2074
	v2099 = v2089
	goto L411
L411:
	;
	v2124 = v2093 | v2094
	v2125 = int32(1)
	v2126 = v2095 - v2125
	v2128 = v2094 << (uint(v2125) % 32)
	if v2128 == int32(256) {
		goto L413
	} else {
		goto L414
	}
L412:
	;
	v2873 = v2124
	v2879 = v2099
	goto L217
L413:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2099))) = uint8(v2124)
	if v2126 == int32(0) {
		v2920 = v971
		goto L13
	} else {
		goto L416
	}
L414:
	;
	goto L415
L415:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2095) {
		v2093 = v2124
		v2094 = v2128
		v2095 = v2126
		goto L411
	} else {
		goto L417
	}
L416:
	;
	v2134 = int32(1)
	v2136 = v2099 + v2134
	v2137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2136))))
	v2093 = v2137
	v2094 = v2134
	v2095 = v2126
	v2099 = v2136
	goto L411
L417:
	;
	goto L412
L418:
	;
	v2144 = v2143 + v953
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v2145 != 0 {
		goto L422
	} else {
		goto L423
	}
L419:
	;
	v2142 = F__emscripten_memcpy_bulkmem(m, v2140, v2141, v953)
	mBase = m.M
	v2143 = v2142
	goto L421
L420:
	;
	v2143 = v2140
	goto L421
L421:
	;
	goto L418
L422:
	;
	v2153 = v2145
	goto L424
L423:
	;
	v2146 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v2153 = (v2146<<(uint(int32(3))%32) + int32(23)) & int32(-8)
	goto L424
L424:
	;
	if v785 != 0 {
		goto L426
	} else {
		goto L427
	}
L425:
	;
	if v967 != 0 {
		goto L430
	} else {
		goto L431
	}
L426:
	;
	v2155 = F__emscripten_memcpy_bulkmem(m, v2144, v2153+v81, v785)
	mBase = m.M
	v2156 = v2155
	goto L428
L427:
	;
	v2156 = v2144
	goto L428
L428:
	;
	goto L425
L429:
	;
	if v631&int32(1) == int32(0) {
		v2920 = v971
		goto L13
	} else {
		goto L433
	}
L430:
	;
	v2160 = F__emscripten_memcpy_bulkmem(m, v2156+v785, v953+v2141+v952, v967)
	mBase = m.M
	goto L432
L431:
	;
	goto L432
L432:
	;
	goto L429
L433:
	;
	v2166 = int32(0)
	v2168 = *(*int32)(unsafe.Add(mBase, uint32(v971)+8))
	if v2168 != 0 {
		goto L434
	} else {
		goto L435
	}
L434:
	;
	v2169 = *(*int32)(unsafe.Add(mBase, uint32(v971)+4))
	v2173 = v985 + v2169<<(uint(int32(3))%32)
	goto L436
L435:
	;
	v2173 = v2166
	goto L436
L436:
	;
	v2174 = *(*int32)(unsafe.Add(mBase, uint32(v79)+8))
	if v2174 != 0 {
		goto L437
	} else {
		goto L438
	}
L437:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v79)+4))
	v2179 = v220 + v2175<<(uint(int32(3))%32)
	goto L439
L438:
	;
	v2179 = v2166
	goto L439
L439:
	;
	if v960 <= int32(0) {
		goto L440
	} else {
		goto L441
	}
L440:
	;
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v81)+8))
	if v2387 != 0 {
		goto L469
	} else {
		goto L470
	}
L441:
	;
	v2185 = int32(1) << (uint(v635&int32(7)) % 32)
	v2187 = base.I32_div_s(v635, int32(8))
	v2188 = v2173 + v2187
	v2189 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2188))))
	if v2179 == int32(0) {
		goto L443
	} else {
		goto L444
	}
L442:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2334))) = uint8(v2326)
	goto L440
L443:
	;
	v2192 = v960
	v2195 = v2185
	v2198 = v2189
	v2206 = v2188
	goto L446
L444:
	;
	goto L445
L445:
	;
	v2241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2179))))
	v2243 = int32(1)
	v2246 = v2185
	v2249 = v2189
	v2251 = v960
	v2252 = v2241
	v2257 = v2188
	v2262 = v2179
	goto L453
L446:
	;
	v2225 = v2195 | v2198
	v2226 = int32(1)
	v2227 = v2192 - v2226
	v2229 = v2195 << (uint(v2226) % 32)
	if v2229 == int32(256) {
		goto L448
	} else {
		goto L449
	}
L447:
	;
	v2326 = v2225
	v2334 = v2206
	goto L442
L448:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2206))) = uint8(v2225)
	if v2227 == int32(0) {
		goto L440
	} else {
		goto L451
	}
L449:
	;
	goto L450
L450:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2192) {
		v2192 = v2227
		v2195 = v2229
		v2198 = v2225
		goto L446
	} else {
		goto L452
	}
L451:
	;
	v2235 = int32(1)
	v2237 = v2206 + v2235
	v2238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2237))))
	v2192 = v2227
	v2195 = v2235
	v2198 = v2238
	v2206 = v2237
	goto L446
L452:
	;
	goto L447
L453:
	;
	if v2243&v2252 != 0 {
		goto L455
	} else {
		goto L456
	}
L454:
	;
	if v2295 == int32(1) {
		goto L440
	} else {
		goto L468
	}
L455:
	;
	v2281 = v2246 | v2249
	goto L457
L456:
	;
	v2281 = v2249 & (v2246 ^ int32(-1))
	goto L457
L457:
	;
	v2282 = int32(1)
	v2283 = v2251 - v2282
	v2285 = v2246 << (uint(v2282) % 32)
	if v2285 == int32(256) {
		goto L458
	} else {
		goto L459
	}
L458:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2257))) = uint8(v2281)
	if v2283 == int32(0) {
		goto L440
	} else {
		goto L461
	}
L459:
	;
	v2295 = v2285
	v2296 = v2281
	v2297 = v2257
	goto L460
L460:
	;
	v2299 = v2243 << (uint(int32(1)) % 32)
	if v2299 == int32(256) {
		goto L463
	} else {
		goto L464
	}
L461:
	;
	v2291 = int32(1)
	v2293 = v2257 + v2291
	v2294 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2293))))
	v2295 = v2291
	v2296 = v2294
	v2297 = v2293
	goto L460
L462:
	;
	goto L454
L463:
	;
	if v2283 == int32(0) {
		goto L462
	} else {
		goto L466
	}
L464:
	;
	v2308 = v2299
	v2309 = v2252
	v2310 = v2262
	goto L465
L465:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2251) {
		v2243 = v2308
		v2246 = v2295
		v2249 = v2296
		v2251 = v2283
		v2252 = v2309
		v2257 = v2297
		v2262 = v2310
		goto L453
	} else {
		goto L467
	}
L466:
	;
	v2304 = int32(1)
	v2305 = v2262 + v2304
	v2306 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2305))))
	v2308 = v2304
	v2309 = v2306
	v2310 = v2305
	goto L465
L467:
	;
	goto L462
L468:
	;
	v2326 = v2296
	v2334 = v2297
	goto L442
L469:
	;
	v2388 = *(*int32)(unsafe.Add(mBase, uint32(v81)+4))
	v2393 = v746 + v2388<<(uint(int32(3))%32)
	goto L471
L470:
	;
	v2393 = int32(0)
	goto L471
L471:
	;
	if v742 <= int32(0) {
		goto L472
	} else {
		goto L473
	}
L472:
	;
	if v954 <= int32(0) {
		v2920 = v971
		goto L13
	} else {
		goto L501
	}
L473:
	;
	v2397 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v2398 = *(*int32)(unsafe.Add(mBase, uint32(v75)+80))
	v2399 = v2397 - v2398
	v2402 = int32(1) << (uint(v2399&int32(7)) % 32)
	v2404 = base.I32_div_s(v2399, int32(8))
	v2405 = v2173 + v2404
	v2406 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2405))))
	if v2393 == int32(0) {
		goto L475
	} else {
		goto L476
	}
L474:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2547))) = uint8(v2542)
	goto L472
L475:
	;
	v2411 = v742
	v2412 = v2402
	v2414 = v2406
	v2419 = v2405
	goto L478
L476:
	;
	goto L477
L477:
	;
	v2458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2393))))
	v2462 = v742
	v2463 = v2402
	v2465 = v2406
	v2466 = int32(1)
	v2470 = v2405
	v2471 = v2458
	v2474 = v2393
	goto L485
L478:
	;
	v2442 = v2412 | v2414
	v2443 = int32(1)
	v2444 = v2411 - v2443
	v2446 = v2412 << (uint(v2443) % 32)
	if v2446 == int32(256) {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2419))) = uint8(v2442)
	if v2444 == int32(0) {
		goto L472
	} else {
		goto L483
	}
L481:
	;
	goto L482
L482:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2411) {
		v2411 = v2444
		v2412 = v2446
		v2414 = v2442
		goto L478
	} else {
		goto L484
	}
L483:
	;
	v2452 = int32(1)
	v2454 = v2419 + v2452
	v2455 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2454))))
	v2411 = v2444
	v2412 = v2452
	v2414 = v2455
	v2419 = v2454
	goto L478
L484:
	;
	v2542 = v2442
	v2547 = v2419
	goto L474
L485:
	;
	if v2466&v2471 != 0 {
		goto L487
	} else {
		goto L488
	}
L486:
	;
	if v2512 == int32(1) {
		goto L472
	} else {
		goto L500
	}
L487:
	;
	v2498 = v2463 | v2465
	goto L489
L488:
	;
	v2498 = v2465 & (v2463 ^ int32(-1))
	goto L489
L489:
	;
	v2499 = int32(1)
	v2500 = v2462 - v2499
	v2502 = v2463 << (uint(v2499) % 32)
	if v2502 == int32(256) {
		goto L490
	} else {
		goto L491
	}
L490:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2470))) = uint8(v2498)
	if v2500 == int32(0) {
		goto L472
	} else {
		goto L493
	}
L491:
	;
	v2512 = v2502
	v2513 = v2498
	v2514 = v2470
	goto L492
L492:
	;
	v2516 = v2466 << (uint(int32(1)) % 32)
	if v2516 == int32(256) {
		goto L495
	} else {
		goto L496
	}
L493:
	;
	v2508 = int32(1)
	v2510 = v2470 + v2508
	v2511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2510))))
	v2512 = v2508
	v2513 = v2511
	v2514 = v2510
	goto L492
L494:
	;
	goto L486
L495:
	;
	if v2500 == int32(0) {
		goto L494
	} else {
		goto L498
	}
L496:
	;
	v2525 = v2516
	v2526 = v2471
	v2527 = v2474
	goto L497
L497:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2462) {
		v2462 = v2500
		v2463 = v2512
		v2465 = v2513
		v2466 = v2525
		v2470 = v2514
		v2471 = v2526
		v2474 = v2527
		goto L485
	} else {
		goto L499
	}
L498:
	;
	v2521 = int32(1)
	v2522 = v2474 + v2521
	v2523 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2522))))
	v2525 = v2521
	v2526 = v2523
	v2527 = v2522
	goto L497
L499:
	;
	goto L494
L500:
	;
	v2542 = v2513
	v2547 = v2514
	goto L474
L501:
	;
	v2607 = v960 + v961
	v2608 = v2607 + v635
	v2611 = int32(1) << (uint(v2608&int32(7)) % 32)
	v2613 = base.I32_div_s(v2608, int32(8))
	v2614 = v2173 + v2613
	v2615 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2614))))
	if v2179 == int32(0) {
		goto L503
	} else {
		goto L504
	}
L502:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2760))) = uint8(v2754)
	v2920 = v971
	goto L13
L503:
	;
	v2620 = v2615
	v2621 = v2611
	v2622 = v954
	v2626 = v2614
	goto L506
L504:
	;
	goto L505
L505:
	;
	v2672 = base.I32_div_s(v2607, int32(8))
	v2673 = v2179 + v2672
	v2674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2673))))
	v2677 = v2615
	v2678 = v2611
	v2679 = v954
	v2680 = int32(1) << (uint(v2607&int32(7)) % 32)
	v2683 = v2614
	v2685 = v2674
	v2686 = v2673
	goto L513
L506:
	;
	v2651 = v2620 | v2621
	v2652 = int32(1)
	v2653 = v2622 - v2652
	v2655 = v2621 << (uint(v2652) % 32)
	if v2655 == int32(256) {
		goto L508
	} else {
		goto L509
	}
L508:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2626))) = uint8(v2651)
	if v2653 == int32(0) {
		v2920 = v971
		goto L13
	} else {
		goto L511
	}
L509:
	;
	goto L510
L510:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2622) {
		v2620 = v2651
		v2621 = v2655
		v2622 = v2653
		goto L506
	} else {
		goto L512
	}
L511:
	;
	v2661 = int32(1)
	v2663 = v2626 + v2661
	v2664 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2663))))
	v2620 = v2664
	v2621 = v2661
	v2622 = v2653
	v2626 = v2663
	goto L506
L512:
	;
	v2754 = v2651
	v2760 = v2626
	goto L502
L513:
	;
	if v2680&v2685 != 0 {
		goto L515
	} else {
		goto L516
	}
L514:
	;
	if v2728 == int32(1) {
		v2920 = v971
		goto L13
	} else {
		goto L528
	}
L515:
	;
	v2713 = v2677 | v2678
	goto L517
L516:
	;
	v2713 = v2677 & (v2678 ^ int32(-1))
	goto L517
L517:
	;
	v2714 = int32(1)
	v2715 = v2679 - v2714
	v2717 = v2678 << (uint(v2714) % 32)
	if v2717 == int32(256) {
		goto L518
	} else {
		goto L519
	}
L518:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2683))) = uint8(v2713)
	if v2715 == int32(0) {
		v2920 = v971
		goto L13
	} else {
		goto L521
	}
L519:
	;
	v2727 = v2713
	v2728 = v2717
	v2729 = v2683
	goto L520
L520:
	;
	v2731 = v2680 << (uint(int32(1)) % 32)
	if v2731 == int32(256) {
		goto L523
	} else {
		goto L524
	}
L521:
	;
	v2723 = int32(1)
	v2725 = v2683 + v2723
	v2726 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2725))))
	v2727 = v2726
	v2728 = v2723
	v2729 = v2725
	goto L520
L522:
	;
	goto L514
L523:
	;
	if v2715 == int32(0) {
		goto L522
	} else {
		goto L526
	}
L524:
	;
	v2740 = v2731
	v2741 = v2685
	v2742 = v2686
	goto L525
L525:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2679) {
		v2677 = v2727
		v2678 = v2728
		v2679 = v2715
		v2680 = v2740
		v2683 = v2729
		v2685 = v2741
		v2686 = v2742
		goto L513
	} else {
		goto L527
	}
L526:
	;
	v2736 = int32(1)
	v2737 = v2686 + v2736
	v2738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2737))))
	v2740 = v2736
	v2741 = v2738
	v2742 = v2737
	goto L525
L527:
	;
	goto L522
L528:
	;
	v2754 = v2727
	v2760 = v2729
	goto L502
L529:
	;
	if v2799&v2804 != 0 {
		goto L531
	} else {
		goto L532
	}
L530:
	;
	if v2847 == int32(1) {
		v2920 = v971
		goto L13
	} else {
		goto L544
	}
L531:
	;
	v2832 = v2796 | v2797
	goto L533
L532:
	;
	v2832 = v2796 & (v2797 ^ int32(-1))
	goto L533
L533:
	;
	v2833 = int32(1)
	v2834 = v2798 - v2833
	v2836 = v2797 << (uint(v2833) % 32)
	if v2836 == int32(256) {
		goto L534
	} else {
		goto L535
	}
L534:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v2802))) = uint8(v2832)
	if v2834 == int32(0) {
		v2920 = v971
		goto L13
	} else {
		goto L537
	}
L535:
	;
	v2846 = v2832
	v2847 = v2836
	v2848 = v2802
	goto L536
L536:
	;
	v2850 = v2799 << (uint(int32(1)) % 32)
	if v2850 == int32(256) {
		goto L539
	} else {
		goto L540
	}
L537:
	;
	v2842 = int32(1)
	v2844 = v2802 + v2842
	v2845 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2844))))
	v2846 = v2845
	v2847 = v2842
	v2848 = v2844
	goto L536
L538:
	;
	goto L530
L539:
	;
	if v2834 == int32(0) {
		goto L538
	} else {
		goto L542
	}
L540:
	;
	v2859 = v2850
	v2860 = v2804
	v2861 = v2805
	goto L541
L541:
	;
	if base.Ui32(int32(1)) < base.Ui32(v2798) {
		v2796 = v2846
		v2797 = v2847
		v2798 = v2834
		v2799 = v2859
		v2802 = v2848
		v2804 = v2860
		v2805 = v2861
		goto L529
	} else {
		goto L543
	}
L542:
	;
	v2855 = int32(1)
	v2856 = v2805 + v2855
	v2857 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2856))))
	v2859 = v2855
	v2860 = v2857
	v2861 = v2856
	goto L541
L543:
	;
	goto L538
L544:
	;
	v2873 = v2846
	v2879 = v2848
	goto L217
}
func F_array_subscript_fetch_old(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
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
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v6 == int32(1) {
		v9 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(v4)+52)) = uint8(v9)
		*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = int32(0)
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v4)+8))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v4)+4))
		v19 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+4)))
		v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+6)))
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+8)))
		v22 = int32(*(*int8)(unsafe.Add(mBase, uint32(v16)+9)))
		v25 = F_array_get_element(m, v14, v15, v16+int32(12), v19, v20, v21, v22, v4+int32(52))
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v4)+48)) = v25
			return
		}
	}
}
func F_array_subscript_handler(m *base.Module, l0 int32) int32 {
	return int32(1694760)
}
func F_array_subscript_transform(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
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
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
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
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
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
	var v136 int32
	_ = v136
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	v6 = int32(0)
	v11 = m.G0
	v13 = v11 - int32(16)
	m.G0 = v13
	if l1 == v6 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	if l3 != 0 {
		goto L53
	} else {
		goto L54
	}
L2:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l0)+24)) = int64(0)
	goto L1
L3:
	;
	goto L4
L4:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v19 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v138
	*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = v136
	if v136 == int32(0) {
		goto L1
	} else {
		goto L47
	}
L6:
	;
	v136 = v6
	v138 = v6
	goto L5
L7:
	;
	goto L8
L8:
	;
	v27 = v6
	v28 = v6
	v29 = v6
	goto L9
L9:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v32+v28<<(uint(int32(2))%32))))
	if l3 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L18
	} else {
		goto L42
	}
L11:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	if v37 != 0 {
		goto L15
	} else {
		goto L16
	}
L12:
	;
	v85 = v29
	goto L13
L13:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	if v86 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L14:
	;
	v82 = F_lappend(m, v29, v81)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L18
	} else {
		goto L30
	}
L15:
	;
	v38 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v39 = F_transformExpr(m, l2, v37, v38)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L18
	} else {
		goto L19
	}
L16:
	;
	goto L17
L17:
	;
	v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v36)+4)))
	if v71 != 0 {
		v81 = int32(0)
		goto L14
	} else {
		goto L28
	}
L18:
	;
	return
L19:
	;
	v41 = F_exprType(m, v39)
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v44 = int32(-1)
	v48 = F_coerce_to_target_type(m, l2, v39, v41, int32(23), v44, int32(1), int32(2), v44)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L18
	} else {
		goto L21
	}
L21:
	;
	if v48 != 0 {
		v81 = v48
		goto L14
	} else {
		goto L22
	}
L22:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v53 = m.ExcPending
	if v53 != 0 {
		goto L18
	} else {
		goto L23
	}
L23:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v56 = m.ExcPending
	if v56 != 0 {
		goto L18
	} else {
		goto L24
	}
L24:
	;
	F_errmsg(m, int32(236519), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v36)+8))
	v62 = F_exprLocation(m, v61)
	mBase = m.M
	F_parser_errposition(m, l2, v62)
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	F_errfinish(m, int32(519325), int32(95), int32(301878))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L18
	} else {
		goto L27
	}
L27:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L28:
	;
	v74 = int32(0)
	v76 = int32(1)
	v79 = F_makeConst(m, int32(23), int32(-1), v74, int32(4), v76, v74, v76)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v81 = v79
	goto L14
L30:
	;
	v85 = v82
	goto L13
L31:
	;
	goto L10
L32:
	;
	v105 = F_lappend(m, v27, v104)
	mBase = m.M
	v106 = m.ExcPending
	if v106 != 0 {
		goto L18
	} else {
		goto L40
	}
L33:
	;
	v104 = int32(0)
	goto L32
L34:
	;
	goto L35
L35:
	;
	v90 = *(*int32)(unsafe.Add(mBase, uint32(l2)+68))
	v91 = F_transformExpr(m, l2, v86, v90)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L18
	} else {
		goto L36
	}
L36:
	;
	v93 = F_exprType(m, v91)
	mBase = m.M
	v94 = m.ExcPending
	if v94 != 0 {
		goto L18
	} else {
		goto L37
	}
L37:
	;
	v96 = int32(-1)
	v100 = F_coerce_to_target_type(m, l2, v91, v93, int32(23), v96, int32(1), int32(2), v96)
	mBase = m.M
	v101 = m.ExcPending
	if v101 != 0 {
		goto L18
	} else {
		goto L38
	}
L38:
	;
	if v100 == int32(0) {
		goto L31
	} else {
		goto L39
	}
L39:
	;
	v104 = v100
	goto L32
L40:
	;
	v108 = v28 + int32(1)
	v109 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v109 <= v108 {
		v136 = v105
		v138 = v85
		goto L5
	} else {
		goto L41
	}
L41:
	;
	v27 = v105
	v28 = v108
	v29 = v85
	goto L9
L42:
	;
	F_errcode(m, int32(67141764))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L18
	} else {
		goto L43
	}
L43:
	;
	F_errmsg(m, int32(236519), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L18
	} else {
		goto L44
	}
L44:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v36)+12))
	v123 = F_exprLocation(m, v122)
	mBase = m.M
	F_parser_errposition(m, l2, v123)
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L18
	} else {
		goto L45
	}
L45:
	;
	F_errfinish(m, int32(519325), int32(132), int32(301878))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L18
	} else {
		goto L46
	}
L46:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L47:
	;
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	if v145 <= int32(6) {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L18
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L18
	} else {
		goto L50
	}
L50:
	;
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v136)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = int32(6)
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v155
	F_errmsg(m, int32(711130), v13)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L18
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(519325), int32(152), int32(301878))
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L18
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	v179 = int32(4)
	goto L55
L54:
	;
	v179 = int32(8)
	goto L55
L55:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l0+v179)))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v181
	m.G0 = v13 + int32(16)
	return
}
func F_array_to_json(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_makeStringInfo(m)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		F_array_to_json_internal(m, v2, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			v11 = *(*int32)(unsafe.Add(mBase, uint32(v3)+4))
			v12 = F_cstring_to_text_with_len(m, v10, v11)
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return int32(0)
			} else {
				return v12
			}
		}
	}
}
func F_array_to_json_internal(m *base.Module, l0 int32, l1 int32, l2 int32) {
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
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v31 int32
	_ = v31
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	v12 = m.G0
	v14 = v12 - int32(32)
	m.G0 = v14
	v16 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
		*(*int32)(unsafe.Add(mBase, uint32(v14)+24)) = int32(0)
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
		v23 = v16 + int32(16)
		v24 = F_ArrayGetNItems(m, v21, v23)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v14)+28)) = v24
			if v24 <= int32(0) {
				F_appendStringInfoString(m, l1, int32(533747))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return
				} else {
					m.G0 = v14 + int32(32)
					return
				}
			} else {
				F_get_typlenbyvalalign(m, v18, v14+int32(14), v14+int32(13), v14+int32(12))
				mBase = m.M
				v39 = m.ExcPending
				if v39 != 0 {
					return
				} else {
					F_json_categorize_type(m, v18, int32(0), v14+int32(8), v14+int32(4))
					mBase = m.M
					v46 = m.ExcPending
					if v46 != 0 {
						return
					} else {
						v47 = int32(*(*int16)(unsafe.Add(mBase, uint32(v14)+14)))
						v48 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+13)))
						v49 = int32(*(*int8)(unsafe.Add(mBase, uint32(v14)+12)))
						F_deconstruct_array(m, v16, v47, v48, v49, v14+int32(20), v14+int32(16), v14+int32(28))
						mBase = m.M
						v57 = m.ExcPending
						if v57 != 0 {
							return
						} else {
							v59 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
							v60 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
							v63 = *(*int32)(unsafe.Add(mBase, uint32(v14)+8))
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v14)+4))
							F_array_dim_to_json(m, l1, int32(0), v21, v23, v59, v60, v14+int32(24), v63, v64, l2)
							mBase = m.M
							v66 = m.ExcPending
							if v66 != 0 {
								return
							} else {
								v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+20))
								F_pfree(m, v67)
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									v70 = *(*int32)(unsafe.Add(mBase, uint32(v14)+16))
									F_pfree(m, v70)
									mBase = m.M
									v72 = m.ExcPending
									if v72 != 0 {
										return
									} else {
										m.G0 = v14 + int32(32)
										return
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
func F_array_to_sparsevec(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v132 int32
	_ = v132
	var v133 float32
	_ = v133
	var v134 float32
	_ = v134
	var v137 float32
	_ = v137
	var v141 float32
	_ = v141
	var v145 float32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v179 int32
	_ = v179
	var v191 float32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v198 int32
	_ = v198
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v228 int32
	_ = v228
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v285 int32
	_ = v285
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 float64
	_ = v329
	var v331 float32
	_ = v331
	var v334 int32
	_ = v334
	var v335 float64
	_ = v335
	var v339 int32
	_ = v339
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v366 int32
	_ = v366
	var v367 float64
	_ = v367
	var v375 int32
	_ = v375
	var v378 int32
	_ = v378
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v391 int32
	_ = v391
	var v394 int32
	_ = v394
	var v398 int32
	_ = v398
	var v403 int32
	_ = v403
	var v407 int32
	_ = v407
	var v410 int32
	_ = v410
	var v414 int32
	_ = v414
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v475 int32
	_ = v475
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v500 int32
	_ = v500
	var v516 float32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v536 int32
	_ = v536
	var v539 int32
	_ = v539
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v560 int32
	_ = v560
	var v561 float64
	_ = v561
	var v562 float32
	_ = v562
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v573 int32
	_ = v573
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v610 int32
	_ = v610
	var v616 int32
	_ = v616
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v626 int32
	_ = v626
	var v643 int32
	_ = v643
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v652 int32
	_ = v652
	var v655 int32
	_ = v655
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v669 int32
	_ = v669
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v711 float32
	_ = v711
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v741 int32
	_ = v741
	var v745 int32
	_ = v745
	var v750 int32
	_ = v750
	var v754 int32
	_ = v754
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v776 int32
	_ = v776
	var v780 int32
	_ = v780
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v797 int32
	_ = v797
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v818 int32
	_ = v818
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(32)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		goto L4
	} else {
		goto L5
	}
L1:
	;
	F_CheckNnz(m, v437, v441)
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L4
	} else {
		goto L82
	}
L2:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L4
	} else {
		goto L78
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v407 = m.ExcPending
	if v407 != 0 {
		goto L4
	} else {
		goto L74
	}
L4:
	;
	return int32(0)
L5:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
	if v26 < int32(2) {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v22)+8))
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	goto L8
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L4
	} else {
		goto L70
	}
L9:
	;
	v31 = F_array_contains_nulls(m, v22)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L4
	} else {
		goto L12
	}
L10:
	;
	goto L11
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	F_get_typlenbyvalalign(m, v33, v19+int32(30), v19+int32(29), v19+int32(28))
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L4
	} else {
		goto L14
	}
L12:
	;
	if v31 != 0 {
		goto L3
	} else {
		goto L13
	}
L13:
	;
	goto L11
L14:
	;
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v19)+30)))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+29)))
	v45 = int32(*(*int8)(unsafe.Add(mBase, uint32(v19)+28)))
	F_deconstruct_array(m, v22, v43, v44, v45, v19+int32(24), int32(0), v19+int32(20))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	F_CheckDim_2(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L4
	} else {
		goto L16
	}
L16:
	;
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if base.B2i32(v29 != int32(-1))&base.B2i32(v58 != v29) != 0 {
		goto L2
	} else {
		goto L17
	}
L17:
	;
	v61 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	switch v61 - int32(700) {
	case 0:
		goto L23
	case 1:
		goto L22
	default:
		goto L24
	}
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v375 = m.ExcPending
	if v375 != 0 {
		goto L4
	} else {
		goto L66
	}
L19:
	;
	if v58&v203 == int32(0) {
		v437 = v345
		v441 = v58
		goto L1
	} else {
		goto L65
	}
L20:
	;
	v309 = int32(0)
	v310 = v2
	v314 = v2
	goto L62
L21:
	;
	if v58 <= int32(0) {
		goto L49
	} else {
		goto L50
	}
L22:
	;
	if v58 <= int32(0) {
		goto L45
	} else {
		goto L46
	}
L23:
	;
	if v58 <= int32(0) {
		goto L32
	} else {
		goto L33
	}
L24:
	;
	if v61 == int32(23) {
		goto L21
	} else {
		goto L25
	}
L25:
	;
	if v61 != int32(1700) {
		goto L18
	} else {
		goto L26
	}
L26:
	;
	v68 = int32(0)
	if v58 <= v68 {
		v437 = v68
		v441 = v58
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v71 = v68
	v72 = v2
	goto L28
L28:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v89+v72<<(uint(int32(2))%32))))
	v94 = F_DirectFunctionCall1Coll(m, int32(1338), int32(0), v93)
	mBase = m.M
	v95 = m.ExcPending
	if v95 != 0 {
		goto L4
	} else {
		goto L30
	}
L29:
	;
	v437 = v98
	v441 = v101
	goto L1
L30:
	;
	v98 = v71 + base.B2i32(v94 != int32(0))
	v100 = v72 + int32(1)
	v101 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v100 < v101 {
		v71 = v98
		v72 = v100
		goto L28
	} else {
		goto L31
	}
L31:
	;
	goto L29
L32:
	;
	v437 = int32(0)
	v441 = v58
	goto L1
L33:
	;
	goto L34
L34:
	;
	v107 = v58 & int32(3)
	v108 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v109 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v58) {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v114 = v109
	v115 = v2
	v122 = v2
	goto L38
L36:
	;
	v154 = v109
	v155 = v2
	goto L37
L37:
	;
	if v107 == int32(0) {
		v437 = v154
		v441 = v58
		goto L1
	} else {
		goto L41
	}
L38:
	;
	v132 = v108 + v115<<(uint(int32(2))%32)
	v133 = *(*float32)(unsafe.Add(mBase, uint32(v132)))
	v134 = float32(0)
	v137 = *(*float32)(unsafe.Add(mBase, uint32(v132)+4))
	v141 = *(*float32)(unsafe.Add(mBase, uint32(v132)+8))
	v145 = *(*float32)(unsafe.Add(mBase, uint32(v132)+12))
	v148 = v114 + base.F32_ne(v133, v134) + base.F32_ne(v137, v134) + base.F32_ne(v141, v134) + base.F32_ne(v145, v134)
	v149 = int32(4)
	v150 = v115 + v149
	v152 = v122 + v149
	if v152 != v58&int32(2147483644) {
		v114 = v148
		v115 = v150
		v122 = v152
		goto L38
	} else {
		goto L40
	}
L39:
	;
	v154 = v148
	v155 = v150
	goto L37
L40:
	;
	goto L39
L41:
	;
	v172 = v154
	v173 = v155
	v179 = v2
	goto L42
L42:
	;
	v191 = *(*float32)(unsafe.Add(mBase, uint32(v108+v173<<(uint(int32(2))%32))))
	v194 = v172 + base.F32_ne(v191, float32(0))
	v195 = int32(1)
	v198 = v179 + v195
	if v198 != v107 {
		v172 = v194
		v173 = v173 + v195
		v179 = v198
		goto L42
	} else {
		goto L44
	}
L43:
	;
	v437 = v194
	v441 = v58
	goto L1
L44:
	;
	goto L43
L45:
	;
	v437 = int32(0)
	v441 = v58
	goto L1
L46:
	;
	goto L47
L47:
	;
	v203 = int32(1)
	v205 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	if v58 != v203 {
		goto L20
	} else {
		goto L48
	}
L48:
	;
	v345 = int32(0)
	v346 = v2
	goto L19
L49:
	;
	v437 = int32(0)
	v441 = v58
	goto L1
L50:
	;
	goto L51
L51:
	;
	v213 = v58 & int32(3)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v215 = int32(0)
	if base.Ui32(int32(4)) <= base.Ui32(v58) {
		goto L52
	} else {
		goto L53
	}
L52:
	;
	v220 = v215
	v221 = v2
	v228 = v2
	goto L55
L53:
	;
	v260 = v215
	v261 = v2
	goto L54
L54:
	;
	if v213 == int32(0) {
		v437 = v260
		v441 = v58
		goto L1
	} else {
		goto L58
	}
L55:
	;
	v238 = v214 + v221<<(uint(int32(2))%32)
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)))
	v240 = int32(0)
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v238)+4))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v238)+8))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	v254 = v220 + base.B2i32(v239 != v240) + base.B2i32(v243 != v240) + base.B2i32(v247 != v240) + base.B2i32(v251 != v240)
	v255 = int32(4)
	v256 = v221 + v255
	v258 = v228 + v255
	if v258 != v58&int32(2147483644) {
		v220 = v254
		v221 = v256
		v228 = v258
		goto L55
	} else {
		goto L57
	}
L56:
	;
	v260 = v254
	v261 = v256
	goto L54
L57:
	;
	goto L56
L58:
	;
	v278 = v260
	v279 = v261
	v285 = v2
	goto L59
L59:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(v214+v279<<(uint(int32(2))%32))))
	v300 = v278 + base.B2i32(v297 != int32(0))
	v301 = int32(1)
	v304 = v285 + v301
	if v304 != v213 {
		v278 = v300
		v279 = v279 + v301
		v285 = v304
		goto L59
	} else {
		goto L61
	}
L60:
	;
	v437 = v300
	v441 = v58
	goto L1
L61:
	;
	goto L60
L62:
	;
	v325 = int32(2)
	v327 = v205 + v310<<(uint(v325)%32)
	v328 = *(*int32)(unsafe.Add(mBase, uint32(v327)))
	v329 = *(*float64)(unsafe.Add(mBase, uint32(v328)))
	v331 = float32(0)
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v327)+4))
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v334)))
	v339 = v309 + base.F32_ne(base.F32_demote_f64(v329), v331) + base.F32_ne(base.F32_demote_f64(v335), v331)
	v341 = v310 + v325
	v343 = v314 + v325
	if v343 != v58&int32(2147483646) {
		v309 = v339
		v310 = v341
		v314 = v343
		goto L62
	} else {
		goto L64
	}
L63:
	;
	v345 = v339
	v346 = v341
	goto L19
L64:
	;
	goto L63
L65:
	;
	v366 = *(*int32)(unsafe.Add(mBase, uint32(v205+v346<<(uint(int32(2))%32))))
	v367 = *(*float64)(unsafe.Add(mBase, uint32(v366)))
	v437 = v345 + base.F32_ne(base.F32_demote_f64(v367), float32(0))
	v441 = v58
	goto L1
L66:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L4
	} else {
		goto L67
	}
L67:
	;
	F_errmsg(m, int32(385435), int32(0))
	mBase = m.M
	v382 = m.ExcPending
	if v382 != 0 {
		goto L4
	} else {
		goto L68
	}
L68:
	;
	F_errfinish(m, int32(525598), int32(753), int32(514910))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L4
	} else {
		goto L69
	}
L69:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L70:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L4
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(570932), int32(0))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L4
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(525598), int32(709), int32(514910))
	mBase = m.M
	v403 = m.ExcPending
	if v403 != 0 {
		goto L4
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	F_errcode(m, int32(67108994))
	mBase = m.M
	v410 = m.ExcPending
	if v410 != 0 {
		goto L4
	} else {
		goto L75
	}
L75:
	;
	F_errmsg(m, int32(162284), int32(0))
	mBase = m.M
	v414 = m.ExcPending
	if v414 != 0 {
		goto L4
	} else {
		goto L76
	}
L76:
	;
	F_errfinish(m, int32(525598), int32(714), int32(514910))
	mBase = m.M
	v419 = m.ExcPending
	if v419 != 0 {
		goto L4
	} else {
		goto L77
	}
L77:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L78:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v426 = m.ExcPending
	if v426 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v58
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v29
	F_errmsg(m, int32(489736), v19)
	mBase = m.M
	v431 = m.ExcPending
	if v431 != 0 {
		goto L4
	} else {
		goto L80
	}
L80:
	;
	F_errfinish(m, int32(525598), int32(62), int32(302808))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L4
	} else {
		goto L81
	}
L81:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L82:
	;
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v458 = F_mul_size(m, int32(4), v437)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L4
	} else {
		goto L83
	}
L83:
	;
	v460 = F_add_size(m, int32(16), v458)
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L4
	} else {
		goto L84
	}
L84:
	;
	v463 = F_mul_size(m, int32(4), v437)
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L4
	} else {
		goto L85
	}
L85:
	;
	v465 = F_add_size(m, v460, v463)
	mBase = m.M
	v466 = m.ExcPending
	if v466 != 0 {
		goto L4
	} else {
		goto L86
	}
L86:
	;
	v467 = F_palloc0(m, v465)
	mBase = m.M
	v468 = m.ExcPending
	if v468 != 0 {
		goto L4
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v467)+8)) = v437
	*(*int32)(unsafe.Add(mBase, uint32(v467)+4)) = v455
	v471 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v467))) = v465 << (uint(v471) % 32)
	v475 = v467 + int32(16)
	v478 = v475 + v437<<(uint(v471)%32)
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v22)+12))
	switch v479 - int32(700) {
	case 0:
		goto L97
	case 1:
		goto L96
	default:
		goto L98
	}
L88:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v806 = m.ExcPending
	if v806 != 0 {
		goto L4
	} else {
		goto L160
	}
L89:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v793 = m.ExcPending
	if v793 != 0 {
		goto L4
	} else {
		goto L157
	}
L90:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v780 = m.ExcPending
	if v780 != 0 {
		goto L4
	} else {
		goto L154
	}
L91:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L4
	} else {
		goto L151
	}
L92:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L4
	} else {
		goto L148
	}
L93:
	;
	v684 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	F_pfree(m, v684)
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L4
	} else {
		goto L134
	}
L94:
	;
	v625 = v486
	v626 = int32(0)
	goto L126
L95:
	;
	v581 = int32(0)
	v582 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v582 <= v581 {
		v669 = v581
		goto L93
	} else {
		goto L118
	}
L96:
	;
	v535 = int32(0)
	v536 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v536 <= v535 {
		v669 = v535
		goto L93
	} else {
		goto L110
	}
L97:
	;
	v491 = int32(0)
	v492 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v492 <= v491 {
		v669 = v491
		goto L93
	} else {
		goto L102
	}
L98:
	;
	if v479 == int32(23) {
		goto L95
	} else {
		goto L99
	}
L99:
	;
	if v479 != int32(1700) {
		goto L88
	} else {
		goto L100
	}
L100:
	;
	v486 = int32(0)
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v486 < v487 {
		goto L94
	} else {
		goto L101
	}
L101:
	;
	v669 = int32(0)
	goto L93
L102:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v497 = int32(0)
	v498 = v491
	v500 = v492
	goto L103
L103:
	;
	v516 = *(*float32)(unsafe.Add(mBase, uint32(v495+v497<<(uint(int32(2))%32))))
	if base.F32_ne(v516, float32(0)) != 0 {
		goto L105
	} else {
		goto L106
	}
L104:
	;
	v669 = v530
	goto L93
L105:
	;
	v519 = *(*int32)(unsafe.Add(mBase, uint32(v467)+8))
	if v519 <= v498 {
		goto L90
	} else {
		goto L108
	}
L106:
	;
	v530 = v498
	v531 = v500
	goto L107
L107:
	;
	v533 = v497 + int32(1)
	if v533 < v531 {
		v497 = v533
		v498 = v530
		v500 = v531
		goto L103
	} else {
		goto L109
	}
L108:
	;
	v522 = v498 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v475+v522))) = v497
	*(*float32)(unsafe.Add(mBase, uint32(v522+v478))) = v516
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v530 = v498 + int32(1)
	v531 = v527
	goto L107
L109:
	;
	goto L104
L110:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v541 = int32(0)
	v542 = v535
	v544 = v536
	goto L111
L111:
	;
	v560 = *(*int32)(unsafe.Add(mBase, uint32(v539+v541<<(uint(int32(2))%32))))
	v561 = *(*float64)(unsafe.Add(mBase, uint32(v560)))
	v562 = base.F32_demote_f64(v561)
	if base.F32_ne(v562, float32(0)) != 0 {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	v669 = v576
	goto L93
L113:
	;
	v565 = *(*int32)(unsafe.Add(mBase, uint32(v467)+8))
	if v565 <= v542 {
		goto L91
	} else {
		goto L116
	}
L114:
	;
	v576 = v542
	v577 = v544
	goto L115
L115:
	;
	v579 = v541 + int32(1)
	if v579 < v577 {
		v541 = v579
		v542 = v576
		v544 = v577
		goto L111
	} else {
		goto L117
	}
L116:
	;
	v568 = v542 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v475+v568))) = v541
	*(*float32)(unsafe.Add(mBase, uint32(v568+v478))) = v562
	v573 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v576 = v542 + int32(1)
	v577 = v573
	goto L115
L117:
	;
	goto L112
L118:
	;
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v587 = int32(0)
	v588 = v581
	v590 = v582
	goto L119
L119:
	;
	v606 = *(*int32)(unsafe.Add(mBase, uint32(v585+v587<<(uint(int32(2))%32))))
	if v606 != 0 {
		goto L121
	} else {
		goto L122
	}
L120:
	;
	v669 = v619
	goto L93
L121:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v467)+8))
	if v607 <= v588 {
		goto L92
	} else {
		goto L124
	}
L122:
	;
	v619 = v588
	v620 = v590
	goto L123
L123:
	;
	v622 = v587 + int32(1)
	if v622 < v620 {
		v587 = v622
		v588 = v619
		v590 = v620
		goto L119
	} else {
		goto L125
	}
L124:
	;
	v610 = v588 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v475+v610))) = v587
	*(*float32)(unsafe.Add(mBase, uint32(v610+v478))) = base.F32_convert_i32_s(v606)
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	v619 = v588 + int32(1)
	v620 = v616
	goto L123
L125:
	;
	goto L120
L126:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v19)+24))
	v647 = *(*int32)(unsafe.Add(mBase, uint32(v643+v625<<(uint(int32(2))%32))))
	v648 = F_DirectFunctionCall1Coll(m, int32(1338), int32(0), v647)
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L4
	} else {
		goto L128
	}
L127:
	;
	v669 = v662
	goto L93
L128:
	;
	if v648&int32(2147483647) != 0 {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v652 = *(*int32)(unsafe.Add(mBase, uint32(v467)+8))
	if v652 <= v626 {
		goto L89
	} else {
		goto L132
	}
L130:
	;
	v662 = v626
	goto L131
L131:
	;
	v665 = v625 + int32(1)
	v666 = *(*int32)(unsafe.Add(mBase, uint32(v19)+20))
	if v665 < v666 {
		v625 = v665
		v626 = v662
		goto L126
	} else {
		goto L133
	}
L132:
	;
	v655 = v626 << (uint(int32(2)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v475+v655))) = v625
	*(*int32)(unsafe.Add(mBase, uint32(v655+v478))) = v648
	v662 = v626 + int32(1)
	goto L131
L133:
	;
	goto L127
L134:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v467)+8))
	if v687 == v669 {
		goto L135
	} else {
		goto L136
	}
L135:
	;
	v689 = int32(0)
	if v689 < v669 {
		goto L138
	} else {
		goto L139
	}
L136:
	;
	goto L137
L137:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v741 = m.ExcPending
	if v741 != 0 {
		goto L4
	} else {
		goto L145
	}
L138:
	;
	v692 = v689
	goto L141
L139:
	;
	goto L140
L140:
	;
	m.G0 = v19 + int32(32)
	return v467
L141:
	;
	v711 = *(*float32)(unsafe.Add(mBase, uint32(v478+v692<<(uint(int32(2))%32))))
	F_CheckElement_2(m, v711)
	mBase = m.M
	v713 = m.ExcPending
	if v713 != 0 {
		goto L4
	} else {
		goto L143
	}
L142:
	;
	goto L140
L143:
	;
	v715 = v692 + int32(1)
	v716 = *(*int32)(unsafe.Add(mBase, uint32(v467)+8))
	if v715 < v716 {
		v692 = v715
		goto L141
	} else {
		goto L144
	}
L144:
	;
	goto L142
L145:
	;
	F_errmsg_internal(m, int32(475886), int32(0))
	mBase = m.M
	v745 = m.ExcPending
	if v745 != 0 {
		goto L4
	} else {
		goto L146
	}
L146:
	;
	F_errfinish(m, int32(525598), int32(810), int32(514910))
	mBase = m.M
	v750 = m.ExcPending
	if v750 != 0 {
		goto L4
	} else {
		goto L147
	}
L147:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L148:
	;
	F_errmsg_internal(m, int32(475866), int32(0))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L4
	} else {
		goto L149
	}
L149:
	;
	F_errfinish(m, int32(525598), int32(776), int32(514910))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L4
	} else {
		goto L150
	}
L150:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L151:
	;
	F_errmsg_internal(m, int32(475866), int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L4
	} else {
		goto L152
	}
L152:
	;
	F_errfinish(m, int32(525598), int32(781), int32(514910))
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L4
	} else {
		goto L153
	}
L153:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L154:
	;
	F_errmsg_internal(m, int32(475866), int32(0))
	mBase = m.M
	v784 = m.ExcPending
	if v784 != 0 {
		goto L4
	} else {
		goto L155
	}
L155:
	;
	F_errfinish(m, int32(525598), int32(786), int32(514910))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L4
	} else {
		goto L156
	}
L156:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L157:
	;
	F_errmsg_internal(m, int32(475866), int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L4
	} else {
		goto L158
	}
L158:
	;
	F_errfinish(m, int32(525598), int32(791), int32(514910))
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L4
	} else {
		goto L159
	}
L159:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L160:
	;
	F_errcode(m, int32(130))
	mBase = m.M
	v809 = m.ExcPending
	if v809 != 0 {
		goto L4
	} else {
		goto L161
	}
L161:
	;
	F_errmsg(m, int32(385435), int32(0))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L4
	} else {
		goto L162
	}
L162:
	;
	F_errfinish(m, int32(525598), int32(797), int32(514910))
	mBase = m.M
	v818 = m.ExcPending
	if v818 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_array_upper(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_DatumGetAnyArrayP(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		if v13 == int32(-1) {
			v16 = int32(28)
		} else {
			v16 = int32(4)
		}
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v7+v16)))
		if base.Ui32(v18-int32(7)) <= base.Ui32(int32(-7)) {
			v23 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v23)
			return int32(0)
		} else {
			v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v28 = int32(0)
			if base.B2i32(v28 < v27)&base.B2i32(base.Ui32(v27) <= base.Ui32(v18)) == v28 {
				v34 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v34)
				return int32(0)
			} else {
				if v13 == int32(-1) {
					v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)+32))
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v7)+36))
					v48 = v40
					v49 = v41
				} else {
					v43 = v7 + int32(16)
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
					v48 = v43
					v49 = v43 + v44<<(uint(int32(2))%32)
				}
				v53 = v27<<(uint(int32(2))%32) - int32(4)
				v55 = *(*int32)(unsafe.Add(mBase, uint32(v48+v53)))
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v49+v53)))
				return v55 + v57 - int32(1)
			}
		}
	}
}
func F_initArrayResult(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	if l2 != 0 {
		v6 = int32(64)
	} else {
		v6 = int32(8)
	}
	v7 = F_initArrayResultWithSize(m, l0, l1, l2, v6)
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_makeArrayResultAny(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v10 != 0 {
		v11 = int32(4562080)
		v12 = *(*int32)(unsafe.Add(mBase, _consts[28]))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
		*(*int32)(unsafe.Add(mBase, _consts[28])) = l1
		*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(1)
		*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v13
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v27 = *(*int32)(unsafe.Add(mBase, uint32(v10)+20))
		v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+24)))
		v29 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+26)))
		v30 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10)+27)))
		v31 = F_construct_md_array(m, v19, v20, base.B2i32(int32(0) < v13), v8+int32(12), v8+int32(8), v27, v28, v29, v30)
		mBase = m.M
		v34 = m.ExcPending
		if v34 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, _consts[28])) = v12
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
			F_MemoryContextDelete(m, v37)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				v44 = v31
				m.G0 = v8 + int32(16)
				return v44
			}
		}
	} else {
		v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v42 = F_makeArrayResultArr(m, v40, l1, int32(1))
		mBase = m.M
		v43 = m.ExcPending
		if v43 != 0 {
			return int32(0)
		} else {
			v44 = v42
			m.G0 = v8 + int32(16)
			return v44
		}
	}
}
func F_transformArrayExpr(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v90 int32
	_ = v90
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v191 int32
	_ = v191
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
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v275 int32
	_ = v275
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v304 int32
	_ = v304
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	v6 = int32(0)
	v15 = m.G0
	v17 = v15 - int32(48)
	m.G0 = v17
	v20 = F_palloc0(m, int32(36))
	mBase = m.M
	v23 = m.ExcPending
	if v23 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v24 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)) = uint8(v24)
	*(*int32)(unsafe.Add(mBase, uint32(v20))) = int32(35)
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v28 != 0 {
		goto L11
	} else {
		goto L12
	}
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v304
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v297
	*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v296
	v311 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+24)) = v311
	v313 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+28)) = v313
	v315 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v315
	m.G0 = v17 + int32(48)
	return v20
L4:
	;
	v296 = l2
	v297 = l3
	v304 = int32(0)
	goto L3
L5:
	;
	v196 = int32(0)
	v197 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v196 < v197 {
		goto L68
	} else {
		goto L69
	}
L6:
	;
	if v90 == int32(0) {
		goto L4
	} else {
		goto L64
	}
L7:
	;
	v136 = F_select_common_type(m, l0, v90, int32(535885), int32(0))
	mBase = m.M
	v137 = m.ExcPending
	if v137 != 0 {
		goto L1
	} else {
		goto L40
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L1
	} else {
		goto L34
	}
L9:
	;
	if l2 != 0 {
		goto L6
	} else {
		goto L32
	}
L10:
	;
	v41 = v6
	v45 = v6
	goto L16
L11:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if int32(0) < v29 {
		goto L10
	} else {
		goto L14
	}
L12:
	;
	goto L13
L13:
	;
	if l2 == int32(0) {
		goto L8
	} else {
		goto L15
	}
L14:
	;
	v90 = v6
	goto L9
L15:
	;
	goto L4
L16:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v28)+12))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v48+v45<<(uint(int32(2))%32))))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)))
	if v53 == int32(80) {
		goto L20
	} else {
		goto L21
	}
L17:
	;
	v90 = v77
	goto L9
L18:
	;
	v77 = F_lappend(m, v41, v75)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L30
	}
L19:
	;
	v73 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)) = uint8(v73)
	v75 = v71
	goto L18
L20:
	;
	v56 = F_transformArrayExpr(m, l0, v52, l2, l3, l4)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L1
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v58 = F_transformExprRecurse(m, l0, v52)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L24
	}
L23:
	;
	v71 = v56
	goto L19
L24:
	;
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)))
	if v60 != 0 {
		v75 = v58
		goto L18
	} else {
		goto L25
	}
L25:
	;
	v61 = F_exprType(m, v58)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	if v61&int32(-9) == int32(22) {
		v75 = v58
		goto L18
	} else {
		goto L27
	}
L27:
	;
	v67 = F_get_element_type(m, v61)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L1
	} else {
		goto L28
	}
L28:
	;
	if v67 == int32(0) {
		v75 = v58
		goto L18
	} else {
		goto L29
	}
L29:
	;
	v71 = v58
	goto L19
L30:
	;
	v80 = v45 + int32(1)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	if v80 < v81 {
		v41 = v77
		v45 = v80
		goto L16
	} else {
		goto L31
	}
L31:
	;
	goto L17
L32:
	;
	if v90 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L8
L34:
	;
	F_errcode(m, int32(134611076))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	F_errmsg(m, int32(25270), int32(0))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L36
	}
L36:
	;
	F_errhint(m, int32(684938), int32(0))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v126 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_parser_errposition(m, l0, v126)
	mBase = m.M
	v128 = m.ExcPending
	if v128 != 0 {
		goto L1
	} else {
		goto L38
	}
L38:
	;
	F_errfinish(m, int32(519374), int32(2107), int32(217843))
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L1
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
	v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)))
	if v138 == int32(1) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v141 = F_get_element_type(m, v136)
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L1
	} else {
		goto L44
	}
L42:
	;
	goto L43
L43:
	;
	v164 = F_get_array_type(m, v136)
	mBase = m.M
	v165 = m.ExcPending
	if v165 != 0 {
		goto L1
	} else {
		goto L54
	}
L44:
	;
	if v141 != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v193 = v141
	v194 = v136
	v195 = v136
	goto L5
L46:
	;
	goto L47
L47:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v146 = m.ExcPending
	if v146 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v149 = m.ExcPending
	if v149 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	v150 = F_format_type_be(m, v136)
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v150
	F_errmsg(m, int32(204103), v17)
	mBase = m.M
	v155 = m.ExcPending
	if v155 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_parser_errposition(m, l0, v156)
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	F_errfinish(m, int32(519374), int32(2121), int32(217843))
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L1
	} else {
		goto L53
	}
L53:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L54:
	;
	if v164 != 0 {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v193 = v136
	v194 = v136
	v195 = v164
	goto L5
L56:
	;
	goto L57
L57:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L58
	}
L58:
	;
	F_errcode(m, int32(67137668))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L1
	} else {
		goto L59
	}
L59:
	;
	v173 = F_format_type_be(m, v136)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v173
	F_errmsg(m, int32(204060), v17+int32(32))
	mBase = m.M
	v180 = m.ExcPending
	if v180 != 0 {
		goto L1
	} else {
		goto L61
	}
L61:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	F_parser_errposition(m, l0, v181)
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L62
	}
L62:
	;
	F_errfinish(m, int32(519374), int32(2132), int32(217843))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L64:
	;
	v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v20)+20)))
	if v191 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v192 = l2
	goto L67
L66:
	;
	v192 = l3
	goto L67
L67:
	;
	v193 = l3
	v194 = v192
	v195 = l2
	goto L5
L68:
	;
	v211 = v196
	v212 = int32(0)
	goto L71
L69:
	;
	v275 = v196
	goto L70
L70:
	;
	v296 = v195
	v297 = v193
	v304 = v275
	goto L3
L71:
	;
	v215 = *(*int32)(unsafe.Add(mBase, uint32(v90)+12))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v215+v212<<(uint(int32(2))%32))))
	if l2 != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	v275 = v259
	goto L70
L73:
	;
	v259 = F_lappend(m, v211, v258)
	mBase = m.M
	v260 = m.ExcPending
	if v260 != 0 {
		goto L1
	} else {
		goto L89
	}
L74:
	;
	v220 = F_exprType(m, v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L77
	}
L75:
	;
	goto L76
L76:
	;
	v256 = F_coerce_to_common_type(m, l0, v219, v194, int32(535885))
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L88
	}
L77:
	;
	v225 = F_coerce_to_target_type(m, l0, v219, v220, v194, l4, int32(3), int32(1), int32(-1))
	mBase = m.M
	v226 = m.ExcPending
	if v226 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	if v225 != 0 {
		v258 = v225
		goto L73
	} else {
		goto L79
	}
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	F_errcode(m, int32(101744772))
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L1
	} else {
		goto L81
	}
L81:
	;
	v234 = F_exprType(m, v219)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	v236 = F_format_type_be(m, v234)
	mBase = m.M
	v237 = m.ExcPending
	if v237 != 0 {
		goto L1
	} else {
		goto L83
	}
L83:
	;
	v238 = F_format_type_be(m, v194)
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L1
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+20)) = v238
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v236
	F_errmsg(m, int32(193678), v17+int32(16))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L1
	} else {
		goto L85
	}
L85:
	;
	v247 = F_exprLocation(m, v219)
	mBase = m.M
	F_parser_errposition(m, l0, v247)
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L86
	}
L86:
	;
	F_errfinish(m, int32(519374), int32(2167), int32(217843))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L87
	}
L87:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L88:
	;
	v258 = v256
	goto L73
L89:
	;
	v262 = v212 + int32(1)
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	if v262 < v263 {
		v211 = v259
		v212 = v262
		goto L71
	} else {
		goto L90
	}
L90:
	;
	goto L72
}
func F_trim_array(m *base.Module, l0 int32) int32 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v13)+4))
		if int32(0) < v18 {
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
			v22 = v21
		} else {
			v22 = int32(0)
		}
		if v17 < int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v78 = m.ExcPending
			if v78 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(352845954))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v10))) = v22
					F_errmsg(m, int32(501224), v10)
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(518800), int32(6947), int32(24879))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			}
		} else {
			if v22 < v17 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v22
						F_errmsg(m, int32(501224), v10)
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(518800), int32(6947), int32(24879))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				}
			} else {
				v26 = int32(0)
				*(*uint16)(unsafe.Add(mBase, uint32(v10)+28)) = uint16(v26)
				*(*uint16)(unsafe.Add(mBase, uint32(v10)+20)) = uint16(v26)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v26
				*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v26
				if v26 < v18 {
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v13+v18<<(uint(int32(2))%32))+16))
					v40 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v10)+16)) = uint8(v40)
					*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v39 + (v22 + (v17 ^ int32(-1)))
				} else {
				}
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v13)+12))
				F_get_typlenbyvalalign(m, v48, v10+int32(94), v10+int32(93), v10+int32(92))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v57 = int32(*(*int16)(unsafe.Add(mBase, uint32(v10)+94)))
					v58 = int32(*(*int8)(unsafe.Add(mBase, uint32(v10)+92)))
					v69 = F_array_get_slice(m, v13, int32(1), v10+int32(32), v10-int32(-64), v10+int32(16), v10+int32(24), int32(-1), v57, v58)
					mBase = m.M
					v70 = m.ExcPending
					if v70 != 0 {
						return int32(0)
					} else {
						m.G0 = v10 + int32(96)
						return v69
					}
				}
			}
		}
	}
}
