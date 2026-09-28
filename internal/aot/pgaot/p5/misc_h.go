package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_HoldPortal(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v5 = *(*int32)(unsafe.Add(mBase, _c_F_HoldPortal[0]))
	v10 = F_AllocSetContextCreateInternal(m, v5, int32(_a_F_HoldPortal_0), int32(0), int32(_a_F_HoldPortal_1), int32(_a_F_HoldPortal_2))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v10
		v13 = int32(_a_F_HoldPortal_3)
		v14 = *(*int32)(unsafe.Add(mBase, _c_F_HoldPortal[1]))
		*(*int32)(unsafe.Add(mBase, _c_F_HoldPortal[1])) = v10
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v20 = int32(1)
		v24 = *(*int32)(unsafe.Add(mBase, _c_F_HoldPortal[2]))
		v25 = F_tuplestore_begin_heap(m, int32(base.Ui32(v17&int32(2))>>(uint(v20)%32)), v20, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v25
			*(*int32)(unsafe.Add(mBase, _c_F_HoldPortal[1])) = v14
			F_PersistHoldablePortal(m, l0)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return
			} else {
				v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+60))
				if v32 != 0 {
					F_ReleaseCachedPlan(m, v32, int32(0))
					mBase = m.M
					v35 = m.ExcPending
					if v35 != 0 {
						return
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(l0)+56)) = int64(0)
						v38 = int32(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v38
						*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = int64(0)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v38
						return
					}
				} else {
					v38 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v38
					*(*int64)(unsafe.Add(mBase, uint32(l0)+20)) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v38
					return
				}
			}
		}
	}
}
func F_has_largeobject_privilege_id(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
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
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_id[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v15 = F_convert_any_priv_string(m, v10, int32(_a_F_has_largeobject_privilege_id_0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int64(0)
		} else {
			if v15&int64(4) == int64(0) {
				v22 = *(*int32)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_id[1]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v24 = v23
			} else {
				v24 = int32(0)
			}
			v25 = F_LargeObjectExistsWithSnapshot(m, v8, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int64(0)
			} else {
				if v25 != 0 {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_id[2])))
					if v29 != 0 {
						v39 = int64(1)
						return v39
					} else {
						v30 = F_pg_largeobject_aclcheck_snapshot(m, v8, v7, v15, v24)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(base.B2i32(v30 == int32(0)))
						}
					}
				} else {
					v36 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v36)
					v39 = int64(0)
					return v39
				}
			}
		}
	}
}
func F_hashbuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
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
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v41 float64
	_ = v41
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v13 = v10 + int32(8)
	v15 = v10 + int32(7)
	v16 = F__hash_convert_tuple(m, l0, l2, l3, v13, v15)
	mBase = m.M
	v17 = m.ExcPending
	if v17 != 0 {
		return
	} else {
		if v16 != 0 {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
			if v18 != 0 {
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
				F_tuplesort_putindextuplevalues(m, v19, v20, l1, v13, v15)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					v41 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
					*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v41, float64(1))
					m.G0 = v10 + int32(16)
					return
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v28 = F_index_form_tuple(m, v23, v10+int32(8), v10+int32(7))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)) = uint16(v30)
					v32 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v28))) = v32
					v34 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
					F__hash_doinsert(m, l0, v28, v34, int32(0))
					mBase = m.M
					v37 = m.ExcPending
					if v37 != 0 {
						return
					} else {
						F_pfree(m, v28)
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
							return
						} else {
							v41 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
							*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v41, float64(1))
							m.G0 = v10 + int32(16)
							return
						}
					}
				}
			}
		} else {
			m.G0 = v10 + int32(16)
			return
		}
	}
}
func F_hashbuildempty(m *base.Module, l0 int32) {
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	v4 = F__hash_init(m, l0, float64(0), int32(3))
	v5 = m.ExcPending
	if v5 != 0 {
		return
	} else {
		return
	}
}
func F_hashbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v19 int64
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 float64
	_ = v47
	var v48 int32
	_ = v48
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v199 int32
	_ = v199
	var v209 int32
	_ = v209
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v255 int32
	_ = v255
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
	var v266 float64
	_ = v266
	var v271 float64
	_ = v271
	var v272 float64
	_ = v272
	var v276 float64
	_ = v276
	var v278 float64
	_ = v278
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 float64
	_ = v293
	var v296 int32
	_ = v296
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v309 int64
	_ = v309
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v312 int32
	_ = v312
	var v313 int64
	_ = v313
	var v314 int32
	_ = v314
	var v318 int32
	_ = v318
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 float64
	_ = v353
	var v354 float64
	_ = v354
	v5 = int32(0)
	v19 = int64(0)
	v20 = m.G0
	v22 = v20 - int32(48)
	m.G0 = v22
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v5
	*(*int64)(unsafe.Add(mBase, uint32(v22)+40)) = v19
	*(*int64)(unsafe.Add(mBase, uint32(v22)+32)) = v19
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v38 = F_read_stream_begin_relation(m, int32(9), v32, v24, v5, int32(123), v22+int32(16), v5)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v45 = F__hash_getcachedmetap(m, v24, v22+int32(28), int32(0))
	mBase = m.M
	v46 = m.ExcPending
	if v46 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v47 = *(*float64)(unsafe.Add(mBase, uint32(v45)+8))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v48
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v45
	v59 = v45
	v60 = v48
	v62 = v5
	goto L4
L4:
	;
	if base.Ui32(v62) <= base.Ui32(v60) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	F_read_stream_end(m, v38)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L1
	} else {
		goto L52
	}
L6:
	;
	v79 = v59
	v82 = v62
	goto L9
L7:
	;
	v199 = v62
	goto L8
L8:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v209 == int32(0) {
		goto L37
	} else {
		goto L38
	}
L9:
	;
	if v82 == int32(0) {
		goto L12
	} else {
		goto L13
	}
L10:
	;
	v199 = v125
	goto L8
L11:
	;
	v128 = F_read_stream_next_buffer(m, v38, int32(0))
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L1
	} else {
		goto L22
	}
L12:
	;
	v125 = int32(1)
	v126 = int32(0)
	goto L11
L13:
	;
	goto L14
L14:
	;
	v96 = int32(1)
	v97 = v82 + v96
	v101 = v97 - v96
	if base.Ui32(int32(2)) <= base.Ui32(v97) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v79+v120<<(uint(int32(2))%32))+48))
	v125 = v97
	v126 = v124
	goto L11
L16:
	;
	v107 = int32(32) - base.I32_clz(v101)
	goto L18
L17:
	;
	v107 = int32(0)
	goto L18
L18:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v107) {
		goto L19
	} else {
		goto L20
	}
L19:
	;
	v110 = int32(3)
	v120 = int32(base.Ui32(v101)>>(uint(v107-v110)%32))&v110 | v107<<(uint(int32(2))%32) - int32(30)
	goto L21
L20:
	;
	v120 = v107
	goto L21
L21:
	;
	goto L15
L22:
	;
	F_LockBufferForCleanup(m, v128)
	mBase = m.M
	v131 = m.ExcPending
	if v131 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	F__hash_checkpage(m, v24, v128, int32(2))
	mBase = m.M
	v134 = m.ExcPending
	if v134 != 0 {
		goto L1
	} else {
		goto L24
	}
L24:
	;
	if v128 < int32(0) {
		goto L27
	} else {
		goto L28
	}
L25:
	;
	v175 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v176 = *(*int32)(unsafe.Add(mBase, uint32(v174)+24))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v174)+28))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v174)+32))
	F_hashbucketcleanup(m, v24, v82, v128, v125+v126, v175, v176, v177, v178, v22+int32(40), v22+int32(32), base.B2i32(v158 == int32(64)), l2, l3)
	mBase = m.M
	v186 = m.ExcPending
	if v186 != 0 {
		goto L1
	} else {
		goto L34
	}
L26:
	;
	v154 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v153)+16)))
	v155 = v154 + v153
	v156 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v155)+12)))
	v158 = v156 & int32(96)
	if v158 != int32(64) {
		v174 = v79
		goto L25
	} else {
		goto L30
	}
L27:
	;
	v139 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[0]))
	v145 = *(*int32)(unsafe.Add(mBase, uint32(v139+(v128^int32(-1))<<(uint(int32(2))%32))))
	v153 = v145
	goto L26
L28:
	;
	goto L29
L29:
	;
	v147 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[1]))
	v153 = v147 + v128<<(uint(int32(13))%32) + int32(-8192)
	goto L26
L30:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v79)+24))
	if base.Ui32(v161) <= base.Ui32(v162) {
		v174 = v79
		goto L25
	} else {
		goto L31
	}
L31:
	;
	v167 = F__hash_getcachedmetap(m, v24, v22+int32(28), int32(1))
	mBase = m.M
	v168 = m.ExcPending
	if v168 != 0 {
		goto L1
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v60
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v125
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v167
	F_read_stream_reset(m, v38)
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v174 = v167
	goto L25
L34:
	;
	F_ReleaseBuffer(m, v128)
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L35
	}
L35:
	;
	if base.Ui32(v125) <= base.Ui32(v60) {
		v79 = v174
		v82 = v125
		goto L9
	} else {
		goto L36
	}
L36:
	;
	goto L10
L37:
	;
	v215 = F__hash_getbuf(m, v24, int32(0), int32(-1), int32(8))
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L40
	}
L38:
	;
	v218 = v209
	goto L39
L39:
	;
	F_LockBufferInternal(m, v218, int32(3))
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L1
	} else {
		goto L41
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v22)+28)) = v215
	v218 = v215
	goto L39
L41:
	;
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v222 < int32(0) {
		goto L43
	} else {
		goto L44
	}
L42:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+48))
	if v60 != v241 {
		goto L46
	} else {
		goto L47
	}
L43:
	;
	v226 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[0]))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v226+(v222^int32(-1))<<(uint(int32(2))%32))))
	v240 = v232
	goto L42
L44:
	;
	goto L45
L45:
	;
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[1]))
	v240 = v234 + v222<<(uint(int32(13))%32) + int32(-8192)
	goto L42
L46:
	;
	F_UnlockBuffer(m, v222)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L1
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	goto L5
L49:
	;
	v248 = F__hash_getcachedmetap(m, v24, v22+int32(28), int32(1))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	v250 = *(*int32)(unsafe.Add(mBase, uint32(v248)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v22)+24)) = v250
	*(*int32)(unsafe.Add(mBase, uint32(v22)+20)) = v199
	*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v248
	F_read_stream_reset(m, v38)
	mBase = m.M
	v255 = m.ExcPending
	if v255 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	v59 = v248
	v60 = v250
	v62 = v199
	goto L4
L52:
	;
	v258 = int32(_a_F_hashbulkdelete_0)
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[2])) = v260 + int32(1)
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v240)+48))
	v266 = *(*float64)(unsafe.Add(mBase, uint32(v240)+32))
	if base.B2i32(v48 != v264)|base.F64_ne(v47, v266) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L53:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v240)+32)) = v278
	v280 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_MarkBufferDirty(m, v280)
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L60
	}
L54:
	;
	v271 = *(*float64)(unsafe.Add(mBase, uint32(v22)+32))
	v278 = v271
	goto L53
L55:
	;
	goto L56
L56:
	;
	v272 = *(*float64)(unsafe.Add(mBase, uint32(v22)+40))
	if base.F64_gt(v266, v272) != 0 {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v276 = base.F64_sub(v266, v272)
	goto L59
L58:
	;
	v276 = float64(0)
	goto L59
L59:
	;
	v278 = v276
	goto L53
L60:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v24)+48))
	v284 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v283)+118)))
	if v284 != int32(112) {
		goto L62
	} else {
		goto L63
	}
L61:
	;
	v314 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	if v314 < int32(0) {
		goto L75
	} else {
		goto L76
	}
L62:
	;
	v311 = F_XLogGetFakeLSN(m, v24)
	mBase = m.M
	v312 = m.ExcPending
	if v312 != 0 {
		goto L1
	} else {
		goto L73
	}
L63:
	;
	v288 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[3]))
	if v288 <= int32(0) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v24)+32))
	if v291 != 0 {
		goto L62
	} else {
		goto L67
	}
L65:
	;
	goto L66
L66:
	;
	v293 = *(*float64)(unsafe.Add(mBase, uint32(v240)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v22)+8)) = v293
	F_XLogBeginInsert(m)
	mBase = m.M
	v296 = m.ExcPending
	if v296 != 0 {
		goto L1
	} else {
		goto L69
	}
L67:
	;
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v24)+40))
	if v292 != 0 {
		goto L62
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v297 = int32(8)
	F_XLogRegisterData(m, v22+v297, v297)
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v22)+28))
	F_XLogRegisterBuffer(m, int32(0), v303, int32(8))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	v309 = F_XLogInsert(m, int32(12), int32(176))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	v313 = v309
	goto L61
L73:
	;
	v313 = v311
	goto L61
L74:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v332))) = base.I64_rotl(v313, int64(32))
	v336 = int32(_a_F_hashbulkdelete_0)
	v338 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[2]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[2])) = v338 - int32(1)
	F_UnlockReleaseBuffer(m, v314)
	mBase = m.M
	v343 = m.ExcPending
	if v343 != 0 {
		goto L1
	} else {
		goto L78
	}
L75:
	;
	v318 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[0]))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v318+(v314^int32(-1))<<(uint(int32(2))%32))))
	v332 = v324
	goto L74
L76:
	;
	goto L77
L77:
	;
	v326 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[1]))
	v332 = v326 + v314<<(uint(int32(13))%32) + int32(-8192)
	goto L74
L78:
	;
	if l1 == int32(0) {
		goto L79
	} else {
		goto L80
	}
L79:
	;
	v347 = F_palloc0(m, int32(40))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L82
	}
L80:
	;
	v349 = l1
	goto L81
L81:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v349)+8)) = v278
	v351 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v349)+4)) = uint8(v351)
	v353 = *(*float64)(unsafe.Add(mBase, uint32(v22)+40))
	v354 = *(*float64)(unsafe.Add(mBase, uint32(v349)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v349)+16)) = base.F64_add(v353, v354)
	m.G0 = v22 + int32(48)
	return v349
L82:
	;
	v349 = v347
	goto L81
}
func F_hashcharextended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
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
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+24)))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	if v3 == int64(0) {
		v10 = int32(-1636608428)
		v49 = v10
		v50 = v10
		v53 = int32(0)
	} else {
		v13 = base.I32_wrap_i64(v3)
		v15 = v13 + int32(1021750440)
		v20 = base.I32_wrap_i64(int64(base.Ui64(v3)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v26 = v13 - v20 - int32(1636608428) ^ base.I32_rotl(v20, int32(6))
		v30 = v15 - v26 ^ base.I32_rotl(v26, int32(8))
		v31 = v20 + v15
		v32 = v26 + v31
		v33 = v30 + v32
		v37 = v31 - v30 ^ base.I32_rotl(v30, int32(16))
		v41 = v32 - v37 ^ base.I32_rotl(v37, int32(19))
		v46 = v37 + v33
		v47 = v41 + v46
		v49 = v47
		v50 = v46
		v53 = v33 - v41 ^ base.I32_rotl(v41, int32(4)) ^ v47
	}
	v54 = int32(14)
	v56 = v53 - base.I32_rotl(v49, v54)
	v61 = v56 ^ (v2 + v50) - base.I32_rotl(v56, int32(11))
	v65 = v49 ^ v61 - base.I32_rotl(v61, int32(25))
	v69 = v65 ^ v56 - base.I32_rotl(v65, int32(16))
	v73 = v69 ^ v61 - base.I32_rotl(v69, int32(4))
	v77 = v73 ^ v65 - base.I32_rotl(v73, v54)
	return base.I64_extend_i32_u(v77)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v77^v69-base.I32_rotl(v77, int32(24)))
}
func F_hashgetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int64
	_ = v5
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v36 int64
	_ = v36
	v5 = int64(0)
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v8 = F__hash_first(m, l0, int32(1))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	if v8 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v18 = v5
	goto L6
L4:
	;
	v36 = v5
	goto L5
L5:
	;
	return v36
L6:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v6)+48))
	v23 = int32(1)
	F_tbm_add_tuples(m, l1, v6+int32(52)+v19<<(uint(int32(3))%32), v23, v23)
	mBase = m.M
	v26 = m.ExcPending
	if v26 != 0 {
		goto L1
	} else {
		goto L8
	}
L7:
	;
	v36 = v28
	goto L5
L8:
	;
	v28 = v18 + int64(1)
	v30 = F__hash_next(m, l0, int32(1))
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v30 != 0 {
		v18 = v28
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
}
func F_hashinetextended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
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
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v327 int32
	_ = v327
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		v8 = int32(1)
		v10 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v4))))
		if v10&v8 != 0 {
			v13 = v8
		} else {
			v13 = int32(4)
		}
		v14 = v4 + v13
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14))))
		if v17 == int32(2) {
			v20 = int32(6)
		} else {
			v20 = int32(18)
		}
		v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v27 = v20 - int32(1636608432)
		if v21 == int64(0) {
			v64 = v27
			v66 = v27
			v68 = v27
		} else {
			v31 = v27 + base.I32_wrap_i64(v21)
			v32 = v31 + v27
			v36 = int32(4)
			v38 = base.I32_wrap_i64(int64(base.Ui64(v21)>>(uint(int64(32))%64))) ^ base.I32_rotl(v27, v36)
			v42 = v31 - v38 ^ base.I32_rotl(v38, int32(6))
			v46 = v32 - v42 ^ base.I32_rotl(v42, int32(8))
			v47 = v32 + v38
			v48 = v42 + v47
			v49 = v46 + v48
			v53 = v47 - v46 ^ base.I32_rotl(v46, int32(16))
			v57 = v48 - v53 ^ base.I32_rotl(v53, int32(19))
			v62 = v49 + v53
			v64 = v62
			v66 = v49 - v57 ^ base.I32_rotl(v57, v36)
			v68 = v57 + v62
		}
		if v14&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v20) {
				v73 = v14
				v74 = v20
				v76 = v64
				v77 = v68
				v78 = v66
				for {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
					v81 = v80 + v77
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v73)))
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v73)+8))
					v85 = v84 + v78
					v87 = int32(4)
					v89 = v82 + v76 - v85 ^ base.I32_rotl(v85, v87)
					v93 = v81 - v89 ^ base.I32_rotl(v89, int32(6))
					v94 = v85 + v81
					v95 = v89 + v94
					v96 = v93 + v95
					v100 = v94 - v93 ^ base.I32_rotl(v93, int32(8))
					v104 = v95 - v100 ^ base.I32_rotl(v100, int32(16))
					v108 = v96 - v104 ^ base.I32_rotl(v104, int32(19))
					v109 = v100 + v96
					v110 = v104 + v109
					v111 = v108 + v110
					v115 = v109 - v108 ^ base.I32_rotl(v108, v87)
					v116 = int32(12)
					v117 = v73 + v116
					v119 = v74 - v116
					if base.Ui32(int32(11)) < base.Ui32(v119) {
						v73 = v117
						v74 = v119
						v76 = v110
						v77 = v111
						v78 = v115
						continue
					} else {
						break
					}
					break
				}
				v122 = v117
				v123 = v119
				v125 = v110
				v126 = v111
				v127 = v115
			} else {
				v122 = v14
				v123 = v20
				v125 = v64
				v126 = v68
				v127 = v66
			}
			switch v123 - int32(1) {
			case 0:
				v292 = v125
				v293 = v126
				v294 = v127
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 1:
				v285 = v125
				v286 = v126
				v287 = v127
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 2:
				v278 = v125
				v279 = v126
				v280 = v127
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
				v285 = v281<<(uint(int32(16))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 3:
				v272 = v126
				v273 = v127
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+3)))
				v278 = v274<<(uint(int32(24))%32) + v125
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
				v285 = v281<<(uint(int32(16))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 4:
				v268 = v126
				v269 = v127
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
				v272 = v268 + v270
				v273 = v269
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+3)))
				v278 = v274<<(uint(int32(24))%32) + v125
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
				v285 = v281<<(uint(int32(16))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 5:
				v262 = v126
				v263 = v127
				v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
				v268 = v264<<(uint(int32(8))%32) + v262
				v269 = v263
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
				v272 = v268 + v270
				v273 = v269
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+3)))
				v278 = v274<<(uint(int32(24))%32) + v125
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
				v285 = v281<<(uint(int32(16))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 6:
				v256 = v126
				v257 = v127
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
				v262 = v258<<(uint(int32(16))%32) + v256
				v263 = v257
				v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
				v268 = v264<<(uint(int32(8))%32) + v262
				v269 = v263
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
				v272 = v268 + v270
				v273 = v269
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+3)))
				v278 = v274<<(uint(int32(24))%32) + v125
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
				v285 = v281<<(uint(int32(16))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 7:
				v251 = v127
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)))
				v256 = v252<<(uint(int32(24))%32) + v126
				v257 = v251
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
				v262 = v258<<(uint(int32(16))%32) + v256
				v263 = v257
				v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
				v268 = v264<<(uint(int32(8))%32) + v262
				v269 = v263
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
				v272 = v268 + v270
				v273 = v269
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+3)))
				v278 = v274<<(uint(int32(24))%32) + v125
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
				v285 = v281<<(uint(int32(16))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 8:
				v246 = v127
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+8)))
				v251 = v247<<(uint(int32(8))%32) + v246
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)))
				v256 = v252<<(uint(int32(24))%32) + v126
				v257 = v251
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
				v262 = v258<<(uint(int32(16))%32) + v256
				v263 = v257
				v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
				v268 = v264<<(uint(int32(8))%32) + v262
				v269 = v263
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
				v272 = v268 + v270
				v273 = v269
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+3)))
				v278 = v274<<(uint(int32(24))%32) + v125
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
				v285 = v281<<(uint(int32(16))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 9:
				v241 = v127
				v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+9)))
				v246 = v242<<(uint(int32(16))%32) + v241
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+8)))
				v251 = v247<<(uint(int32(8))%32) + v246
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)))
				v256 = v252<<(uint(int32(24))%32) + v126
				v257 = v251
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
				v262 = v258<<(uint(int32(16))%32) + v256
				v263 = v257
				v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
				v268 = v264<<(uint(int32(8))%32) + v262
				v269 = v263
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
				v272 = v268 + v270
				v273 = v269
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+3)))
				v278 = v274<<(uint(int32(24))%32) + v125
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
				v285 = v281<<(uint(int32(16))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			case 10:
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+10)))
				v241 = v237<<(uint(int32(24))%32) + v127
				v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+9)))
				v246 = v242<<(uint(int32(16))%32) + v241
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+8)))
				v251 = v247<<(uint(int32(8))%32) + v246
				v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+7)))
				v256 = v252<<(uint(int32(24))%32) + v126
				v257 = v251
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+6)))
				v262 = v258<<(uint(int32(16))%32) + v256
				v263 = v257
				v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+5)))
				v268 = v264<<(uint(int32(8))%32) + v262
				v269 = v263
				v270 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+4)))
				v272 = v268 + v270
				v273 = v269
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+3)))
				v278 = v274<<(uint(int32(24))%32) + v125
				v279 = v272
				v280 = v273
				v281 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+2)))
				v285 = v281<<(uint(int32(16))%32) + v278
				v286 = v279
				v287 = v280
				v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+1)))
				v292 = v288<<(uint(int32(8))%32) + v285
				v293 = v286
				v294 = v287
				v295 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122))))
				v300 = v292 + v295
				v301 = v293
				v302 = v294
			default:
				v300 = v125
				v301 = v126
				v302 = v127
			}
		} else {
			if base.Ui32(int32(12)) <= base.Ui32(v20) {
				v133 = v14
				v134 = v20
				v136 = v64
				v137 = v68
				v138 = v66
				for {
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
					v141 = v140 + v137
					v142 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
					v144 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
					v145 = v144 + v138
					v147 = int32(4)
					v149 = v142 + v136 - v145 ^ base.I32_rotl(v145, v147)
					v153 = v141 - v149 ^ base.I32_rotl(v149, int32(6))
					v154 = v145 + v141
					v155 = v149 + v154
					v156 = v153 + v155
					v160 = v154 - v153 ^ base.I32_rotl(v153, int32(8))
					v164 = v155 - v160 ^ base.I32_rotl(v160, int32(16))
					v168 = v156 - v164 ^ base.I32_rotl(v164, int32(19))
					v169 = v160 + v156
					v170 = v164 + v169
					v171 = v168 + v170
					v175 = v169 - v168 ^ base.I32_rotl(v168, v147)
					v176 = int32(12)
					v177 = v133 + v176
					v179 = v134 - v176
					if base.Ui32(int32(11)) < base.Ui32(v179) {
						v133 = v177
						v134 = v179
						v136 = v170
						v137 = v171
						v138 = v175
						continue
					} else {
						break
					}
					break
				}
				v182 = v177
				v183 = v179
				v185 = v170
				v186 = v171
				v187 = v175
			} else {
				v182 = v14
				v183 = v20
				v185 = v64
				v186 = v68
				v187 = v66
			}
			switch v183 - int32(1) {
			case 0:
				v234 = v185
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v300 = v234 + v235
				v301 = v186
				v302 = v187
			case 1:
				v229 = v185
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v234 = v230<<(uint(int32(8))%32) + v229
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v300 = v234 + v235
				v301 = v186
				v302 = v187
			case 2:
				v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+2)))
				v229 = v225<<(uint(int32(16))%32) + v185
				v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+1)))
				v234 = v230<<(uint(int32(8))%32) + v229
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182))))
				v300 = v234 + v235
				v301 = v186
				v302 = v187
			case 3:
				v222 = v186
				v223 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
				v300 = v223 + v185
				v301 = v222
				v302 = v187
			case 4:
				v219 = v186
				v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v222 = v219 + v220
				v223 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
				v300 = v223 + v185
				v301 = v222
				v302 = v187
			case 5:
				v214 = v186
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
				v219 = v215<<(uint(int32(8))%32) + v214
				v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v222 = v219 + v220
				v223 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
				v300 = v223 + v185
				v301 = v222
				v302 = v187
			case 6:
				v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+6)))
				v214 = v210<<(uint(int32(16))%32) + v186
				v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+5)))
				v219 = v215<<(uint(int32(8))%32) + v214
				v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+4)))
				v222 = v219 + v220
				v223 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
				v300 = v223 + v185
				v301 = v222
				v302 = v187
			case 7:
				v205 = v187
				v206 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
				v208 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
				v300 = v206 + v185
				v301 = v208 + v186
				v302 = v205
			case 8:
				v200 = v187
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
				v205 = v201<<(uint(int32(8))%32) + v200
				v206 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
				v208 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
				v300 = v206 + v185
				v301 = v208 + v186
				v302 = v205
			case 9:
				v195 = v187
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+9)))
				v200 = v196<<(uint(int32(16))%32) + v195
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
				v205 = v201<<(uint(int32(8))%32) + v200
				v206 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
				v208 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
				v300 = v206 + v185
				v301 = v208 + v186
				v302 = v205
			case 10:
				v191 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+10)))
				v195 = v191<<(uint(int32(24))%32) + v187
				v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+9)))
				v200 = v196<<(uint(int32(16))%32) + v195
				v201 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v182)+8)))
				v205 = v201<<(uint(int32(8))%32) + v200
				v206 = *(*int32)(unsafe.Add(mBase, uint32(v182)))
				v208 = *(*int32)(unsafe.Add(mBase, uint32(v182)+4))
				v300 = v206 + v185
				v301 = v208 + v186
				v302 = v205
			default:
				v300 = v185
				v301 = v186
				v302 = v187
			}
		}
		v305 = int32(14)
		v307 = v301 ^ v302 - base.I32_rotl(v301, v305)
		v311 = v307 ^ v300 - base.I32_rotl(v307, int32(11))
		v315 = v311 ^ v301 - base.I32_rotl(v311, int32(25))
		v319 = v315 ^ v307 - base.I32_rotl(v315, int32(16))
		v323 = v319 ^ v311 - base.I32_rotl(v319, int32(4))
		v327 = v323 ^ v315 - base.I32_rotl(v323, v305)
		return base.I64_extend_i32_u(v327)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v327^v319-base.I32_rotl(v327, int32(24)))
	}
}
func F_hashmacaddr8extended(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int64
	_ = v4
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v191 int32
	_ = v191
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
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
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = int32(-1636608424)
	if v4 == int64(0) {
		v47 = v10
		v49 = v10
		v51 = v10
	} else {
		v13 = base.I32_wrap_i64(v4)
		v15 = v13 + int32(1021750448)
		v19 = int32(4)
		v21 = base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(32))%64))) ^ base.I32_rotl(v10, v19)
		v25 = v10 + v13 - v21 ^ base.I32_rotl(v21, int32(6))
		v29 = v15 - v25 ^ base.I32_rotl(v25, int32(8))
		v30 = v15 + v21
		v31 = v25 + v30
		v32 = v29 + v31
		v36 = v30 - v29 ^ base.I32_rotl(v29, int32(16))
		v40 = v31 - v36 ^ base.I32_rotl(v36, int32(19))
		v45 = v32 + v36
		v47 = v45
		v49 = v32 - v40 ^ base.I32_rotl(v40, v19)
		v51 = v40 + v45
	}
	if v2&int32(3) != 0 {
		switch int32(7) {
		case 0:
			v275 = v47
			v276 = v51
			v277 = v49
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 1:
			v268 = v47
			v269 = v51
			v270 = v49
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 2:
			v261 = v47
			v262 = v51
			v263 = v49
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 3:
			v255 = v51
			v256 = v49
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 4:
			v251 = v51
			v252 = v49
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 5:
			v245 = v51
			v246 = v49
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 6:
			v239 = v51
			v240 = v49
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 7:
			v234 = v49
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v239 = v235<<(uint(int32(24))%32) + v51
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 8:
			v229 = v49
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v234 = v230<<(uint(int32(8))%32) + v229
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v239 = v235<<(uint(int32(24))%32) + v51
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 9:
			v224 = v49
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v229 = v225<<(uint(int32(16))%32) + v224
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v234 = v230<<(uint(int32(8))%32) + v229
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v239 = v235<<(uint(int32(24))%32) + v51
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		case 10:
			v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v224 = v220<<(uint(int32(24))%32) + v49
			v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v229 = v225<<(uint(int32(16))%32) + v224
			v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v234 = v230<<(uint(int32(8))%32) + v229
			v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v239 = v235<<(uint(int32(24))%32) + v51
			v240 = v234
			v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v245 = v241<<(uint(int32(16))%32) + v239
			v246 = v240
			v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v251 = v247<<(uint(int32(8))%32) + v245
			v252 = v246
			v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v255 = v251 + v253
			v256 = v252
			v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v261 = v257<<(uint(int32(24))%32) + v47
			v262 = v255
			v263 = v256
			v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v268 = v264<<(uint(int32(16))%32) + v261
			v269 = v262
			v270 = v263
			v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v275 = v271<<(uint(int32(8))%32) + v268
			v276 = v269
			v277 = v270
			v278 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v275 + v278
			v284 = v276
			v285 = v277
		default:
			v283 = v47
			v284 = v51
			v285 = v49
		}
	} else {
		switch int32(7) {
		case 0:
			v217 = v47
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v217 + v218
			v284 = v51
			v285 = v49
		case 1:
			v212 = v47
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v217 = v213<<(uint(int32(8))%32) + v212
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v217 + v218
			v284 = v51
			v285 = v49
		case 2:
			v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v212 = v208<<(uint(int32(16))%32) + v47
			v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v217 = v213<<(uint(int32(8))%32) + v212
			v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v283 = v217 + v218
			v284 = v51
			v285 = v49
		case 3:
			v205 = v51
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v283 = v206 + v47
			v284 = v205
			v285 = v49
		case 4:
			v202 = v51
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v205 = v202 + v203
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v283 = v206 + v47
			v284 = v205
			v285 = v49
		case 5:
			v197 = v51
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v205 = v202 + v203
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v283 = v206 + v47
			v284 = v205
			v285 = v49
		case 6:
			v193 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v197 = v193<<(uint(int32(16))%32) + v51
			v198 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v202 = v198<<(uint(int32(8))%32) + v197
			v203 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v205 = v202 + v203
			v206 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v283 = v206 + v47
			v284 = v205
			v285 = v49
		case 7:
			v188 = v49
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v283 = v189 + v47
			v284 = v191 + v51
			v285 = v188
		case 8:
			v183 = v49
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v283 = v189 + v47
			v284 = v191 + v51
			v285 = v188
		case 9:
			v178 = v49
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v283 = v189 + v47
			v284 = v191 + v51
			v285 = v188
		case 10:
			v174 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v178 = v174<<(uint(int32(24))%32) + v49
			v179 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v183 = v179<<(uint(int32(16))%32) + v178
			v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v188 = v184<<(uint(int32(8))%32) + v183
			v189 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v191 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v283 = v189 + v47
			v284 = v191 + v51
			v285 = v188
		default:
			v283 = v47
			v284 = v51
			v285 = v49
		}
	}
	v288 = int32(14)
	v290 = v284 ^ v285 - base.I32_rotl(v284, v288)
	v294 = v290 ^ v283 - base.I32_rotl(v290, int32(11))
	v298 = v294 ^ v284 - base.I32_rotl(v294, int32(25))
	v302 = v298 ^ v290 - base.I32_rotl(v298, int32(16))
	v306 = v302 ^ v294 - base.I32_rotl(v302, int32(4))
	v310 = v306 ^ v298 - base.I32_rotl(v306, v288)
	return base.I64_extend_i32_u(v310)<<(uint(int64(32))%64) | base.I64_extend_i32_u(v310^v302-base.I32_rotl(v310, int32(24)))
}
func F_hashtext(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v48 int32
	_ = v48
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
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
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
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v280 int32
	_ = v280
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int32
	_ = v335
	var v336 int32
	_ = v336
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v381 int32
	_ = v381
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v409 int32
	_ = v409
	var v410 int32
	_ = v410
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v431 int32
	_ = v431
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v440 int32
	_ = v440
	var v442 int32
	_ = v442
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v453 int32
	_ = v453
	var v457 int32
	_ = v457
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v493 int32
	_ = v493
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v526 int32
	_ = v526
	var v527 int32
	_ = v527
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v550 int32
	_ = v550
	var v552 int32
	_ = v552
	var v556 int32
	_ = v556
	var v560 int32
	_ = v560
	var v564 int32
	_ = v564
	var v568 int32
	_ = v568
	var v572 int32
	_ = v572
	var v577 int32
	_ = v577
	var v581 int32
	_ = v581
	var v587 int32
	_ = v587
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v637 int32
	_ = v637
	var v639 int32
	_ = v639
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v646 int32
	_ = v646
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v682 int32
	_ = v682
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v707 int32
	_ = v707
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v726 int32
	_ = v726
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v739 int32
	_ = v739
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v752 int32
	_ = v752
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v789 int32
	_ = v789
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v825 int32
	_ = v825
	var v829 int32
	_ = v829
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v853 int32
	_ = v853
	var v856 int32
	_ = v856
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v870 int32
	_ = v870
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v884 int32
	_ = v884
	var v887 int32
	_ = v887
	var v893 int32
	_ = v893
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v906 int32
	_ = v906
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v915 int32
	_ = v915
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v921 int32
	_ = v921
	var v922 int32
	_ = v922
	var v926 int32
	_ = v926
	var v930 int32
	_ = v930
	var v934 int32
	_ = v934
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v941 int32
	_ = v941
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v950 int32
	_ = v950
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v956 int32
	_ = v956
	var v960 int32
	_ = v960
	var v961 int32
	_ = v961
	var v965 int32
	_ = v965
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v988 int32
	_ = v988
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
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
	var v1015 int32
	_ = v1015
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1022 int32
	_ = v1022
	var v1023 int32
	_ = v1023
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1032 int32
	_ = v1032
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1043 int32
	_ = v1043
	var v1044 int32
	_ = v1044
	var v1045 int32
	_ = v1045
	var v1047 int32
	_ = v1047
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1058 int32
	_ = v1058
	var v1062 int32
	_ = v1062
	var v1063 int32
	_ = v1063
	var v1067 int32
	_ = v1067
	var v1068 int32
	_ = v1068
	var v1072 int32
	_ = v1072
	var v1073 int32
	_ = v1073
	var v1077 int32
	_ = v1077
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1113 int32
	_ = v1113
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1122 int32
	_ = v1122
	var v1125 int32
	_ = v1125
	var v1127 int32
	_ = v1127
	var v1131 int32
	_ = v1131
	var v1135 int32
	_ = v1135
	var v1139 int32
	_ = v1139
	var v1143 int32
	_ = v1143
	var v1147 int32
	_ = v1147
	var v1153 int32
	_ = v1153
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1174 int32
	_ = v1174
	var v1178 int32
	_ = v1178
	var v1183 int32
	_ = v1183
	var v1187 int32
	_ = v1187
	var v1191 int32
	_ = v1191
	var v1196 int32
	_ = v1196
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v14 != 0 {
			v15 = F_pg_newlocale_from_collation(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int64(0)
			} else {
				v17 = int32(1)
				v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				if v19&v17 != 0 {
					v22 = v17
				} else {
					v22 = int32(4)
				}
				v23 = v10 + v22
				v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15))))
				if v24 == int32(1) {
					if v19 == int32(1) {
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
						if v32 == int32(18) {
							v35 = int32(16)
						} else {
							v35 = int32(0)
						}
						if base.Ui32((v32-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v42 = int32(4)
						} else {
							v42 = v35
						}
						v48 = v42 - int32(1636608432)
						if v23&int32(3) != 0 {
							if base.Ui32(int32(11)) < base.Ui32(v42) {
								v157 = v23
								v158 = v42
								v159 = v48
								v160 = v48
								v161 = v48
								for {
									v163 = *(*int32)(unsafe.Add(mBase, uint32(v157)+4))
									v164 = v163 + v160
									v165 = *(*int32)(unsafe.Add(mBase, uint32(v157)))
									v167 = *(*int32)(unsafe.Add(mBase, uint32(v157)+8))
									v168 = v167 + v161
									v170 = int32(4)
									v172 = v165 + v159 - v168 ^ base.I32_rotl(v168, v170)
									v176 = v164 - v172 ^ base.I32_rotl(v172, int32(6))
									v177 = v168 + v164
									v178 = v172 + v177
									v179 = v176 + v178
									v183 = v177 - v176 ^ base.I32_rotl(v176, int32(8))
									v187 = v178 - v183 ^ base.I32_rotl(v183, int32(16))
									v191 = v179 - v187 ^ base.I32_rotl(v187, int32(19))
									v192 = v183 + v179
									v193 = v187 + v192
									v194 = v191 + v193
									v198 = v192 - v191 ^ base.I32_rotl(v191, v170)
									v199 = int32(12)
									v200 = v157 + v199
									v202 = v158 - v199
									if base.Ui32(int32(11)) < base.Ui32(v202) {
										v157 = v200
										v158 = v202
										v159 = v193
										v160 = v194
										v161 = v198
										continue
									} else {
										break
									}
									break
								}
								v205 = v200
								v206 = v202
								v207 = v193
								v208 = v194
								v209 = v198
							} else {
								v205 = v23
								v206 = v42
								v207 = v48
								v208 = v48
								v209 = v48
							}
							switch v206 - int32(1) {
							case 0:
								v268 = v207
								v269 = v208
								v270 = v209
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 1:
								v261 = v207
								v262 = v208
								v263 = v209
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 2:
								v254 = v207
								v255 = v208
								v256 = v209
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
								v261 = v257<<(uint(int32(16))%32) + v254
								v262 = v255
								v263 = v256
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 3:
								v248 = v208
								v249 = v209
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
								v254 = v250<<(uint(int32(24))%32) + v207
								v255 = v248
								v256 = v249
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
								v261 = v257<<(uint(int32(16))%32) + v254
								v262 = v255
								v263 = v256
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 4:
								v244 = v208
								v245 = v209
								v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
								v248 = v244 + v246
								v249 = v245
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
								v254 = v250<<(uint(int32(24))%32) + v207
								v255 = v248
								v256 = v249
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
								v261 = v257<<(uint(int32(16))%32) + v254
								v262 = v255
								v263 = v256
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 5:
								v238 = v208
								v239 = v209
								v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
								v244 = v240<<(uint(int32(8))%32) + v238
								v245 = v239
								v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
								v248 = v244 + v246
								v249 = v245
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
								v254 = v250<<(uint(int32(24))%32) + v207
								v255 = v248
								v256 = v249
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
								v261 = v257<<(uint(int32(16))%32) + v254
								v262 = v255
								v263 = v256
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 6:
								v232 = v208
								v233 = v209
								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
								v238 = v234<<(uint(int32(16))%32) + v232
								v239 = v233
								v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
								v244 = v240<<(uint(int32(8))%32) + v238
								v245 = v239
								v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
								v248 = v244 + v246
								v249 = v245
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
								v254 = v250<<(uint(int32(24))%32) + v207
								v255 = v248
								v256 = v249
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
								v261 = v257<<(uint(int32(16))%32) + v254
								v262 = v255
								v263 = v256
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 7:
								v227 = v209
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+7)))
								v232 = v228<<(uint(int32(24))%32) + v208
								v233 = v227
								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
								v238 = v234<<(uint(int32(16))%32) + v232
								v239 = v233
								v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
								v244 = v240<<(uint(int32(8))%32) + v238
								v245 = v239
								v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
								v248 = v244 + v246
								v249 = v245
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
								v254 = v250<<(uint(int32(24))%32) + v207
								v255 = v248
								v256 = v249
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
								v261 = v257<<(uint(int32(16))%32) + v254
								v262 = v255
								v263 = v256
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 8:
								v222 = v209
								v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+8)))
								v227 = v223<<(uint(int32(8))%32) + v222
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+7)))
								v232 = v228<<(uint(int32(24))%32) + v208
								v233 = v227
								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
								v238 = v234<<(uint(int32(16))%32) + v232
								v239 = v233
								v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
								v244 = v240<<(uint(int32(8))%32) + v238
								v245 = v239
								v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
								v248 = v244 + v246
								v249 = v245
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
								v254 = v250<<(uint(int32(24))%32) + v207
								v255 = v248
								v256 = v249
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
								v261 = v257<<(uint(int32(16))%32) + v254
								v262 = v255
								v263 = v256
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 9:
								v217 = v209
								v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+9)))
								v222 = v218<<(uint(int32(16))%32) + v217
								v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+8)))
								v227 = v223<<(uint(int32(8))%32) + v222
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+7)))
								v232 = v228<<(uint(int32(24))%32) + v208
								v233 = v227
								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
								v238 = v234<<(uint(int32(16))%32) + v232
								v239 = v233
								v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
								v244 = v240<<(uint(int32(8))%32) + v238
								v245 = v239
								v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
								v248 = v244 + v246
								v249 = v245
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
								v254 = v250<<(uint(int32(24))%32) + v207
								v255 = v248
								v256 = v249
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
								v261 = v257<<(uint(int32(16))%32) + v254
								v262 = v255
								v263 = v256
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							case 10:
								v213 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+10)))
								v217 = v213<<(uint(int32(24))%32) + v209
								v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+9)))
								v222 = v218<<(uint(int32(16))%32) + v217
								v223 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+8)))
								v227 = v223<<(uint(int32(8))%32) + v222
								v228 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+7)))
								v232 = v228<<(uint(int32(24))%32) + v208
								v233 = v227
								v234 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+6)))
								v238 = v234<<(uint(int32(16))%32) + v232
								v239 = v233
								v240 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+5)))
								v244 = v240<<(uint(int32(8))%32) + v238
								v245 = v239
								v246 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+4)))
								v248 = v244 + v246
								v249 = v245
								v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+3)))
								v254 = v250<<(uint(int32(24))%32) + v207
								v255 = v248
								v256 = v249
								v257 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+2)))
								v261 = v257<<(uint(int32(16))%32) + v254
								v262 = v255
								v263 = v256
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205)+1)))
								v268 = v264<<(uint(int32(8))%32) + v261
								v269 = v262
								v270 = v263
								v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v205))))
								v275 = v268 + v271
								v276 = v269
								v277 = v270
							default:
								v275 = v207
								v276 = v208
								v277 = v209
							}
						} else {
							if base.Ui32(v42) < base.Ui32(int32(12)) {
								v103 = v23
								v104 = v42
								v105 = v48
								v106 = v48
								v107 = v48
							} else {
								v55 = v23
								v56 = v42
								v57 = v48
								v58 = v48
								v59 = v48
								for {
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v55)+4))
									v62 = v61 + v58
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v55)))
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
									v66 = v65 + v59
									v68 = int32(4)
									v70 = v63 + v57 - v66 ^ base.I32_rotl(v66, v68)
									v74 = v62 - v70 ^ base.I32_rotl(v70, int32(6))
									v75 = v66 + v62
									v76 = v70 + v75
									v77 = v74 + v76
									v81 = v75 - v74 ^ base.I32_rotl(v74, int32(8))
									v85 = v76 - v81 ^ base.I32_rotl(v81, int32(16))
									v89 = v77 - v85 ^ base.I32_rotl(v85, int32(19))
									v90 = v81 + v77
									v91 = v85 + v90
									v92 = v89 + v91
									v96 = v90 - v89 ^ base.I32_rotl(v89, v68)
									v97 = int32(12)
									v98 = v55 + v97
									v100 = v56 - v97
									if base.Ui32(int32(11)) < base.Ui32(v100) {
										v55 = v98
										v56 = v100
										v57 = v91
										v58 = v92
										v59 = v96
										continue
									} else {
										break
									}
									break
								}
								v103 = v98
								v104 = v100
								v105 = v91
								v106 = v92
								v107 = v96
							}
							switch v104 - int32(1) {
							case 0:
								v154 = v105
								v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
								v275 = v154 + v155
								v276 = v106
								v277 = v107
							case 1:
								v149 = v105
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
								v154 = v150<<(uint(int32(8))%32) + v149
								v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
								v275 = v154 + v155
								v276 = v106
								v277 = v107
							case 2:
								v145 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+2)))
								v149 = v145<<(uint(int32(16))%32) + v105
								v150 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+1)))
								v154 = v150<<(uint(int32(8))%32) + v149
								v155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103))))
								v275 = v154 + v155
								v276 = v106
								v277 = v107
							case 3:
								v142 = v106
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
								v275 = v143 + v105
								v276 = v142
								v277 = v107
							case 4:
								v139 = v106
								v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
								v142 = v139 + v140
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
								v275 = v143 + v105
								v276 = v142
								v277 = v107
							case 5:
								v134 = v106
								v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+5)))
								v139 = v135<<(uint(int32(8))%32) + v134
								v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
								v142 = v139 + v140
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
								v275 = v143 + v105
								v276 = v142
								v277 = v107
							case 6:
								v130 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+6)))
								v134 = v130<<(uint(int32(16))%32) + v106
								v135 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+5)))
								v139 = v135<<(uint(int32(8))%32) + v134
								v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+4)))
								v142 = v139 + v140
								v143 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
								v275 = v143 + v105
								v276 = v142
								v277 = v107
							case 7:
								v125 = v107
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
								v275 = v126 + v105
								v276 = v128 + v106
								v277 = v125
							case 8:
								v120 = v107
								v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+8)))
								v125 = v121<<(uint(int32(8))%32) + v120
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
								v275 = v126 + v105
								v276 = v128 + v106
								v277 = v125
							case 9:
								v115 = v107
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+9)))
								v120 = v116<<(uint(int32(16))%32) + v115
								v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+8)))
								v125 = v121<<(uint(int32(8))%32) + v120
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
								v275 = v126 + v105
								v276 = v128 + v106
								v277 = v125
							case 10:
								v111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+10)))
								v115 = v111<<(uint(int32(24))%32) + v107
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+9)))
								v120 = v116<<(uint(int32(16))%32) + v115
								v121 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+8)))
								v125 = v121<<(uint(int32(8))%32) + v120
								v126 = *(*int32)(unsafe.Add(mBase, uint32(v103)))
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v103)+4))
								v275 = v126 + v105
								v276 = v128 + v106
								v277 = v125
							default:
								v275 = v105
								v276 = v106
								v277 = v107
							}
						}
						v280 = int32(14)
						v282 = v276 ^ v277 - base.I32_rotl(v276, v280)
						v286 = v282 ^ v275 - base.I32_rotl(v282, int32(11))
						v290 = v286 ^ v276 - base.I32_rotl(v286, int32(25))
						v294 = v290 ^ v282 - base.I32_rotl(v290, int32(16))
						v298 = v294 ^ v286 - base.I32_rotl(v294, int32(4))
						v302 = v298 ^ v290 - base.I32_rotl(v298, v280)
						v1154 = v302 ^ v294 - base.I32_rotl(v302, int32(24))
					} else {
						if v19&int32(1) != 0 {
							v309 = int32(1)
							v312 = int32(base.Ui32(v19)>>(uint(v309)%32)) - v309
							v318 = v312 - int32(1636608432)
							if v23&int32(3) != 0 {
								if base.Ui32(int32(11)) < base.Ui32(v312) {
									v427 = v23
									v428 = v312
									v429 = v318
									v430 = v318
									v431 = v318
									for {
										v433 = *(*int32)(unsafe.Add(mBase, uint32(v427)+4))
										v434 = v433 + v430
										v435 = *(*int32)(unsafe.Add(mBase, uint32(v427)))
										v437 = *(*int32)(unsafe.Add(mBase, uint32(v427)+8))
										v438 = v437 + v431
										v440 = int32(4)
										v442 = v435 + v429 - v438 ^ base.I32_rotl(v438, v440)
										v446 = v434 - v442 ^ base.I32_rotl(v442, int32(6))
										v447 = v438 + v434
										v448 = v442 + v447
										v449 = v446 + v448
										v453 = v447 - v446 ^ base.I32_rotl(v446, int32(8))
										v457 = v448 - v453 ^ base.I32_rotl(v453, int32(16))
										v461 = v449 - v457 ^ base.I32_rotl(v457, int32(19))
										v462 = v453 + v449
										v463 = v457 + v462
										v464 = v461 + v463
										v468 = v462 - v461 ^ base.I32_rotl(v461, v440)
										v469 = int32(12)
										v470 = v427 + v469
										v472 = v428 - v469
										if base.Ui32(int32(11)) < base.Ui32(v472) {
											v427 = v470
											v428 = v472
											v429 = v463
											v430 = v464
											v431 = v468
											continue
										} else {
											break
										}
										break
									}
									v475 = v470
									v476 = v472
									v477 = v463
									v478 = v464
									v479 = v468
								} else {
									v475 = v23
									v476 = v312
									v477 = v318
									v478 = v318
									v479 = v318
								}
								switch v476 - int32(1) {
								case 0:
									v538 = v477
									v539 = v478
									v540 = v479
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 1:
									v531 = v477
									v532 = v478
									v533 = v479
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 2:
									v524 = v477
									v525 = v478
									v526 = v479
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+2)))
									v531 = v527<<(uint(int32(16))%32) + v524
									v532 = v525
									v533 = v526
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 3:
									v518 = v478
									v519 = v479
									v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+3)))
									v524 = v520<<(uint(int32(24))%32) + v477
									v525 = v518
									v526 = v519
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+2)))
									v531 = v527<<(uint(int32(16))%32) + v524
									v532 = v525
									v533 = v526
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 4:
									v514 = v478
									v515 = v479
									v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+4)))
									v518 = v514 + v516
									v519 = v515
									v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+3)))
									v524 = v520<<(uint(int32(24))%32) + v477
									v525 = v518
									v526 = v519
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+2)))
									v531 = v527<<(uint(int32(16))%32) + v524
									v532 = v525
									v533 = v526
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 5:
									v508 = v478
									v509 = v479
									v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+5)))
									v514 = v510<<(uint(int32(8))%32) + v508
									v515 = v509
									v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+4)))
									v518 = v514 + v516
									v519 = v515
									v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+3)))
									v524 = v520<<(uint(int32(24))%32) + v477
									v525 = v518
									v526 = v519
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+2)))
									v531 = v527<<(uint(int32(16))%32) + v524
									v532 = v525
									v533 = v526
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 6:
									v502 = v478
									v503 = v479
									v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+6)))
									v508 = v504<<(uint(int32(16))%32) + v502
									v509 = v503
									v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+5)))
									v514 = v510<<(uint(int32(8))%32) + v508
									v515 = v509
									v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+4)))
									v518 = v514 + v516
									v519 = v515
									v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+3)))
									v524 = v520<<(uint(int32(24))%32) + v477
									v525 = v518
									v526 = v519
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+2)))
									v531 = v527<<(uint(int32(16))%32) + v524
									v532 = v525
									v533 = v526
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 7:
									v497 = v479
									v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+7)))
									v502 = v498<<(uint(int32(24))%32) + v478
									v503 = v497
									v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+6)))
									v508 = v504<<(uint(int32(16))%32) + v502
									v509 = v503
									v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+5)))
									v514 = v510<<(uint(int32(8))%32) + v508
									v515 = v509
									v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+4)))
									v518 = v514 + v516
									v519 = v515
									v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+3)))
									v524 = v520<<(uint(int32(24))%32) + v477
									v525 = v518
									v526 = v519
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+2)))
									v531 = v527<<(uint(int32(16))%32) + v524
									v532 = v525
									v533 = v526
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 8:
									v492 = v479
									v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+8)))
									v497 = v493<<(uint(int32(8))%32) + v492
									v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+7)))
									v502 = v498<<(uint(int32(24))%32) + v478
									v503 = v497
									v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+6)))
									v508 = v504<<(uint(int32(16))%32) + v502
									v509 = v503
									v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+5)))
									v514 = v510<<(uint(int32(8))%32) + v508
									v515 = v509
									v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+4)))
									v518 = v514 + v516
									v519 = v515
									v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+3)))
									v524 = v520<<(uint(int32(24))%32) + v477
									v525 = v518
									v526 = v519
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+2)))
									v531 = v527<<(uint(int32(16))%32) + v524
									v532 = v525
									v533 = v526
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 9:
									v487 = v479
									v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
									v492 = v488<<(uint(int32(16))%32) + v487
									v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+8)))
									v497 = v493<<(uint(int32(8))%32) + v492
									v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+7)))
									v502 = v498<<(uint(int32(24))%32) + v478
									v503 = v497
									v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+6)))
									v508 = v504<<(uint(int32(16))%32) + v502
									v509 = v503
									v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+5)))
									v514 = v510<<(uint(int32(8))%32) + v508
									v515 = v509
									v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+4)))
									v518 = v514 + v516
									v519 = v515
									v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+3)))
									v524 = v520<<(uint(int32(24))%32) + v477
									v525 = v518
									v526 = v519
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+2)))
									v531 = v527<<(uint(int32(16))%32) + v524
									v532 = v525
									v533 = v526
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								case 10:
									v483 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+10)))
									v487 = v483<<(uint(int32(24))%32) + v479
									v488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+9)))
									v492 = v488<<(uint(int32(16))%32) + v487
									v493 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+8)))
									v497 = v493<<(uint(int32(8))%32) + v492
									v498 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+7)))
									v502 = v498<<(uint(int32(24))%32) + v478
									v503 = v497
									v504 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+6)))
									v508 = v504<<(uint(int32(16))%32) + v502
									v509 = v503
									v510 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+5)))
									v514 = v510<<(uint(int32(8))%32) + v508
									v515 = v509
									v516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+4)))
									v518 = v514 + v516
									v519 = v515
									v520 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+3)))
									v524 = v520<<(uint(int32(24))%32) + v477
									v525 = v518
									v526 = v519
									v527 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+2)))
									v531 = v527<<(uint(int32(16))%32) + v524
									v532 = v525
									v533 = v526
									v534 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475)+1)))
									v538 = v534<<(uint(int32(8))%32) + v531
									v539 = v532
									v540 = v533
									v541 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v475))))
									v545 = v538 + v541
									v546 = v539
									v547 = v540
								default:
									v545 = v477
									v546 = v478
									v547 = v479
								}
							} else {
								if base.Ui32(v312) < base.Ui32(int32(12)) {
									v373 = v23
									v374 = v312
									v375 = v318
									v376 = v318
									v377 = v318
								} else {
									v325 = v23
									v326 = v312
									v327 = v318
									v328 = v318
									v329 = v318
									for {
										v331 = *(*int32)(unsafe.Add(mBase, uint32(v325)+4))
										v332 = v331 + v328
										v333 = *(*int32)(unsafe.Add(mBase, uint32(v325)))
										v335 = *(*int32)(unsafe.Add(mBase, uint32(v325)+8))
										v336 = v335 + v329
										v338 = int32(4)
										v340 = v333 + v327 - v336 ^ base.I32_rotl(v336, v338)
										v344 = v332 - v340 ^ base.I32_rotl(v340, int32(6))
										v345 = v336 + v332
										v346 = v340 + v345
										v347 = v344 + v346
										v351 = v345 - v344 ^ base.I32_rotl(v344, int32(8))
										v355 = v346 - v351 ^ base.I32_rotl(v351, int32(16))
										v359 = v347 - v355 ^ base.I32_rotl(v355, int32(19))
										v360 = v351 + v347
										v361 = v355 + v360
										v362 = v359 + v361
										v366 = v360 - v359 ^ base.I32_rotl(v359, v338)
										v367 = int32(12)
										v368 = v325 + v367
										v370 = v326 - v367
										if base.Ui32(int32(11)) < base.Ui32(v370) {
											v325 = v368
											v326 = v370
											v327 = v361
											v328 = v362
											v329 = v366
											continue
										} else {
											break
										}
										break
									}
									v373 = v368
									v374 = v370
									v375 = v361
									v376 = v362
									v377 = v366
								}
								switch v374 - int32(1) {
								case 0:
									v424 = v375
									v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
									v545 = v424 + v425
									v546 = v376
									v547 = v377
								case 1:
									v419 = v375
									v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+1)))
									v424 = v420<<(uint(int32(8))%32) + v419
									v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
									v545 = v424 + v425
									v546 = v376
									v547 = v377
								case 2:
									v415 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+2)))
									v419 = v415<<(uint(int32(16))%32) + v375
									v420 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+1)))
									v424 = v420<<(uint(int32(8))%32) + v419
									v425 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373))))
									v545 = v424 + v425
									v546 = v376
									v547 = v377
								case 3:
									v412 = v376
									v413 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
									v545 = v413 + v375
									v546 = v412
									v547 = v377
								case 4:
									v409 = v376
									v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+4)))
									v412 = v409 + v410
									v413 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
									v545 = v413 + v375
									v546 = v412
									v547 = v377
								case 5:
									v404 = v376
									v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+5)))
									v409 = v405<<(uint(int32(8))%32) + v404
									v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+4)))
									v412 = v409 + v410
									v413 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
									v545 = v413 + v375
									v546 = v412
									v547 = v377
								case 6:
									v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+6)))
									v404 = v400<<(uint(int32(16))%32) + v376
									v405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+5)))
									v409 = v405<<(uint(int32(8))%32) + v404
									v410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+4)))
									v412 = v409 + v410
									v413 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
									v545 = v413 + v375
									v546 = v412
									v547 = v377
								case 7:
									v395 = v377
									v396 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
									v398 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
									v545 = v396 + v375
									v546 = v398 + v376
									v547 = v395
								case 8:
									v390 = v377
									v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+8)))
									v395 = v391<<(uint(int32(8))%32) + v390
									v396 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
									v398 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
									v545 = v396 + v375
									v546 = v398 + v376
									v547 = v395
								case 9:
									v385 = v377
									v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+9)))
									v390 = v386<<(uint(int32(16))%32) + v385
									v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+8)))
									v395 = v391<<(uint(int32(8))%32) + v390
									v396 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
									v398 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
									v545 = v396 + v375
									v546 = v398 + v376
									v547 = v395
								case 10:
									v381 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+10)))
									v385 = v381<<(uint(int32(24))%32) + v377
									v386 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+9)))
									v390 = v386<<(uint(int32(16))%32) + v385
									v391 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v373)+8)))
									v395 = v391<<(uint(int32(8))%32) + v390
									v396 = *(*int32)(unsafe.Add(mBase, uint32(v373)))
									v398 = *(*int32)(unsafe.Add(mBase, uint32(v373)+4))
									v545 = v396 + v375
									v546 = v398 + v376
									v547 = v395
								default:
									v545 = v375
									v546 = v376
									v547 = v377
								}
							}
							v550 = int32(14)
							v552 = v546 ^ v547 - base.I32_rotl(v546, v550)
							v556 = v552 ^ v545 - base.I32_rotl(v552, int32(11))
							v560 = v556 ^ v546 - base.I32_rotl(v556, int32(25))
							v564 = v560 ^ v552 - base.I32_rotl(v560, int32(16))
							v568 = v564 ^ v556 - base.I32_rotl(v564, int32(4))
							v572 = v568 ^ v560 - base.I32_rotl(v568, v550)
							v1154 = v572 ^ v564 - base.I32_rotl(v572, int32(24))
						} else {
							v577 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v581 = int32(base.Ui32(v577)>>(uint(int32(2))%32)) - int32(4)
							v587 = v581 - int32(1636608432)
							if v23&int32(3) != 0 {
								if base.Ui32(int32(11)) < base.Ui32(v581) {
									v696 = v23
									v697 = v581
									v698 = v587
									v699 = v587
									v700 = v587
									for {
										v702 = *(*int32)(unsafe.Add(mBase, uint32(v696)+4))
										v703 = v702 + v699
										v704 = *(*int32)(unsafe.Add(mBase, uint32(v696)))
										v706 = *(*int32)(unsafe.Add(mBase, uint32(v696)+8))
										v707 = v706 + v700
										v709 = int32(4)
										v711 = v704 + v698 - v707 ^ base.I32_rotl(v707, v709)
										v715 = v703 - v711 ^ base.I32_rotl(v711, int32(6))
										v716 = v707 + v703
										v717 = v711 + v716
										v718 = v715 + v717
										v722 = v716 - v715 ^ base.I32_rotl(v715, int32(8))
										v726 = v717 - v722 ^ base.I32_rotl(v722, int32(16))
										v730 = v718 - v726 ^ base.I32_rotl(v726, int32(19))
										v731 = v722 + v718
										v732 = v726 + v731
										v733 = v730 + v732
										v737 = v731 - v730 ^ base.I32_rotl(v730, v709)
										v738 = int32(12)
										v739 = v696 + v738
										v741 = v697 - v738
										if base.Ui32(int32(11)) < base.Ui32(v741) {
											v696 = v739
											v697 = v741
											v698 = v732
											v699 = v733
											v700 = v737
											continue
										} else {
											break
										}
										break
									}
									v744 = v739
									v745 = v741
									v746 = v732
									v747 = v733
									v748 = v737
								} else {
									v744 = v23
									v745 = v581
									v746 = v587
									v747 = v587
									v748 = v587
								}
								switch v745 - int32(1) {
								case 0:
									v807 = v746
									v808 = v747
									v809 = v748
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 1:
									v800 = v746
									v801 = v747
									v802 = v748
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 2:
									v793 = v746
									v794 = v747
									v795 = v748
									v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v800 = v796<<(uint(int32(16))%32) + v793
									v801 = v794
									v802 = v795
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 3:
									v787 = v747
									v788 = v748
									v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
									v793 = v789<<(uint(int32(24))%32) + v746
									v794 = v787
									v795 = v788
									v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v800 = v796<<(uint(int32(16))%32) + v793
									v801 = v794
									v802 = v795
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 4:
									v783 = v747
									v784 = v748
									v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+4)))
									v787 = v783 + v785
									v788 = v784
									v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
									v793 = v789<<(uint(int32(24))%32) + v746
									v794 = v787
									v795 = v788
									v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v800 = v796<<(uint(int32(16))%32) + v793
									v801 = v794
									v802 = v795
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 5:
									v777 = v747
									v778 = v748
									v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+5)))
									v783 = v779<<(uint(int32(8))%32) + v777
									v784 = v778
									v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+4)))
									v787 = v783 + v785
									v788 = v784
									v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
									v793 = v789<<(uint(int32(24))%32) + v746
									v794 = v787
									v795 = v788
									v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v800 = v796<<(uint(int32(16))%32) + v793
									v801 = v794
									v802 = v795
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 6:
									v771 = v747
									v772 = v748
									v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+6)))
									v777 = v773<<(uint(int32(16))%32) + v771
									v778 = v772
									v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+5)))
									v783 = v779<<(uint(int32(8))%32) + v777
									v784 = v778
									v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+4)))
									v787 = v783 + v785
									v788 = v784
									v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
									v793 = v789<<(uint(int32(24))%32) + v746
									v794 = v787
									v795 = v788
									v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v800 = v796<<(uint(int32(16))%32) + v793
									v801 = v794
									v802 = v795
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 7:
									v766 = v748
									v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+7)))
									v771 = v767<<(uint(int32(24))%32) + v747
									v772 = v766
									v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+6)))
									v777 = v773<<(uint(int32(16))%32) + v771
									v778 = v772
									v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+5)))
									v783 = v779<<(uint(int32(8))%32) + v777
									v784 = v778
									v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+4)))
									v787 = v783 + v785
									v788 = v784
									v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
									v793 = v789<<(uint(int32(24))%32) + v746
									v794 = v787
									v795 = v788
									v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v800 = v796<<(uint(int32(16))%32) + v793
									v801 = v794
									v802 = v795
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 8:
									v761 = v748
									v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+8)))
									v766 = v762<<(uint(int32(8))%32) + v761
									v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+7)))
									v771 = v767<<(uint(int32(24))%32) + v747
									v772 = v766
									v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+6)))
									v777 = v773<<(uint(int32(16))%32) + v771
									v778 = v772
									v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+5)))
									v783 = v779<<(uint(int32(8))%32) + v777
									v784 = v778
									v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+4)))
									v787 = v783 + v785
									v788 = v784
									v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
									v793 = v789<<(uint(int32(24))%32) + v746
									v794 = v787
									v795 = v788
									v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v800 = v796<<(uint(int32(16))%32) + v793
									v801 = v794
									v802 = v795
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 9:
									v756 = v748
									v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+9)))
									v761 = v757<<(uint(int32(16))%32) + v756
									v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+8)))
									v766 = v762<<(uint(int32(8))%32) + v761
									v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+7)))
									v771 = v767<<(uint(int32(24))%32) + v747
									v772 = v766
									v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+6)))
									v777 = v773<<(uint(int32(16))%32) + v771
									v778 = v772
									v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+5)))
									v783 = v779<<(uint(int32(8))%32) + v777
									v784 = v778
									v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+4)))
									v787 = v783 + v785
									v788 = v784
									v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
									v793 = v789<<(uint(int32(24))%32) + v746
									v794 = v787
									v795 = v788
									v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v800 = v796<<(uint(int32(16))%32) + v793
									v801 = v794
									v802 = v795
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								case 10:
									v752 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+10)))
									v756 = v752<<(uint(int32(24))%32) + v748
									v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+9)))
									v761 = v757<<(uint(int32(16))%32) + v756
									v762 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+8)))
									v766 = v762<<(uint(int32(8))%32) + v761
									v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+7)))
									v771 = v767<<(uint(int32(24))%32) + v747
									v772 = v766
									v773 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+6)))
									v777 = v773<<(uint(int32(16))%32) + v771
									v778 = v772
									v779 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+5)))
									v783 = v779<<(uint(int32(8))%32) + v777
									v784 = v778
									v785 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+4)))
									v787 = v783 + v785
									v788 = v784
									v789 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+3)))
									v793 = v789<<(uint(int32(24))%32) + v746
									v794 = v787
									v795 = v788
									v796 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+2)))
									v800 = v796<<(uint(int32(16))%32) + v793
									v801 = v794
									v802 = v795
									v803 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744)+1)))
									v807 = v803<<(uint(int32(8))%32) + v800
									v808 = v801
									v809 = v802
									v810 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v744))))
									v814 = v807 + v810
									v815 = v808
									v816 = v809
								default:
									v814 = v746
									v815 = v747
									v816 = v748
								}
							} else {
								if base.Ui32(v581) < base.Ui32(int32(12)) {
									v642 = v23
									v643 = v581
									v644 = v587
									v645 = v587
									v646 = v587
								} else {
									v594 = v23
									v595 = v581
									v596 = v587
									v597 = v587
									v598 = v587
									for {
										v600 = *(*int32)(unsafe.Add(mBase, uint32(v594)+4))
										v601 = v600 + v597
										v602 = *(*int32)(unsafe.Add(mBase, uint32(v594)))
										v604 = *(*int32)(unsafe.Add(mBase, uint32(v594)+8))
										v605 = v604 + v598
										v607 = int32(4)
										v609 = v602 + v596 - v605 ^ base.I32_rotl(v605, v607)
										v613 = v601 - v609 ^ base.I32_rotl(v609, int32(6))
										v614 = v605 + v601
										v615 = v609 + v614
										v616 = v613 + v615
										v620 = v614 - v613 ^ base.I32_rotl(v613, int32(8))
										v624 = v615 - v620 ^ base.I32_rotl(v620, int32(16))
										v628 = v616 - v624 ^ base.I32_rotl(v624, int32(19))
										v629 = v620 + v616
										v630 = v624 + v629
										v631 = v628 + v630
										v635 = v629 - v628 ^ base.I32_rotl(v628, v607)
										v636 = int32(12)
										v637 = v594 + v636
										v639 = v595 - v636
										if base.Ui32(int32(11)) < base.Ui32(v639) {
											v594 = v637
											v595 = v639
											v596 = v630
											v597 = v631
											v598 = v635
											continue
										} else {
											break
										}
										break
									}
									v642 = v637
									v643 = v639
									v644 = v630
									v645 = v631
									v646 = v635
								}
								switch v643 - int32(1) {
								case 0:
									v693 = v644
									v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642))))
									v814 = v693 + v694
									v815 = v645
									v816 = v646
								case 1:
									v688 = v644
									v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+1)))
									v693 = v689<<(uint(int32(8))%32) + v688
									v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642))))
									v814 = v693 + v694
									v815 = v645
									v816 = v646
								case 2:
									v684 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+2)))
									v688 = v684<<(uint(int32(16))%32) + v644
									v689 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+1)))
									v693 = v689<<(uint(int32(8))%32) + v688
									v694 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642))))
									v814 = v693 + v694
									v815 = v645
									v816 = v646
								case 3:
									v681 = v645
									v682 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
									v814 = v682 + v644
									v815 = v681
									v816 = v646
								case 4:
									v678 = v645
									v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+4)))
									v681 = v678 + v679
									v682 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
									v814 = v682 + v644
									v815 = v681
									v816 = v646
								case 5:
									v673 = v645
									v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+5)))
									v678 = v674<<(uint(int32(8))%32) + v673
									v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+4)))
									v681 = v678 + v679
									v682 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
									v814 = v682 + v644
									v815 = v681
									v816 = v646
								case 6:
									v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+6)))
									v673 = v669<<(uint(int32(16))%32) + v645
									v674 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+5)))
									v678 = v674<<(uint(int32(8))%32) + v673
									v679 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+4)))
									v681 = v678 + v679
									v682 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
									v814 = v682 + v644
									v815 = v681
									v816 = v646
								case 7:
									v664 = v646
									v665 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
									v667 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
									v814 = v665 + v644
									v815 = v667 + v645
									v816 = v664
								case 8:
									v659 = v646
									v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+8)))
									v664 = v660<<(uint(int32(8))%32) + v659
									v665 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
									v667 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
									v814 = v665 + v644
									v815 = v667 + v645
									v816 = v664
								case 9:
									v654 = v646
									v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+9)))
									v659 = v655<<(uint(int32(16))%32) + v654
									v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+8)))
									v664 = v660<<(uint(int32(8))%32) + v659
									v665 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
									v667 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
									v814 = v665 + v644
									v815 = v667 + v645
									v816 = v664
								case 10:
									v650 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+10)))
									v654 = v650<<(uint(int32(24))%32) + v646
									v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+9)))
									v659 = v655<<(uint(int32(16))%32) + v654
									v660 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v642)+8)))
									v664 = v660<<(uint(int32(8))%32) + v659
									v665 = *(*int32)(unsafe.Add(mBase, uint32(v642)))
									v667 = *(*int32)(unsafe.Add(mBase, uint32(v642)+4))
									v814 = v665 + v644
									v815 = v667 + v645
									v816 = v664
								default:
									v814 = v644
									v815 = v645
									v816 = v646
								}
							}
							v819 = int32(14)
							v821 = v815 ^ v816 - base.I32_rotl(v815, v819)
							v825 = v821 ^ v814 - base.I32_rotl(v821, int32(11))
							v829 = v825 ^ v815 - base.I32_rotl(v825, int32(25))
							v833 = v829 ^ v821 - base.I32_rotl(v829, int32(16))
							v837 = v833 ^ v825 - base.I32_rotl(v833, int32(4))
							v841 = v837 ^ v829 - base.I32_rotl(v837, v819)
							v1154 = v841 ^ v833 - base.I32_rotl(v841, int32(24))
						}
					}
					v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
					if v1158 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v1161 = m.ExcPending
						if v1161 != 0 {
							return int64(0)
						} else {
							return base.I64_extend_i32_u(v1154)
						}
					} else {
						return base.I64_extend_i32_u(v1154)
					}
				} else {
					v846 = int32(0)
					if v19 == int32(1) {
						v853 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+1)))
						if v853 == int32(18) {
							v856 = int32(16)
						} else {
							v856 = int32(0)
						}
						if base.Ui32((v853-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v863 = int32(4)
						} else {
							v863 = v856
						}
						v876 = v863
					} else {
						v864 = int32(1)
						if v19&v864 != 0 {
							v876 = int32(base.Ui32(v19)>>(uint(v864)%32)) - v864
						} else {
							v870 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v876 = int32(base.Ui32(v870)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v877 = F_pg_strnxfrm(m, v846, v846, v23, v876, v15)
					mBase = m.M
					v878 = m.ExcPending
					if v878 != 0 {
						return int64(0)
					} else {
						v880 = v877 + int32(1)
						v881 = F_palloc(m, v880)
						mBase = m.M
						v882 = m.ExcPending
						if v882 != 0 {
							return int64(0)
						} else {
							v883 = F_pg_strnxfrm(m, v881, v880, v23, v876, v15)
							mBase = m.M
							v884 = m.ExcPending
							if v884 != 0 {
								return int64(0)
							} else {
								if base.Ui32(v877) < base.Ui32(v883) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1187 = m.ExcPending
									if v1187 != 0 {
										return int64(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_hashtext_0), int32(0))
										mBase = m.M
										v1191 = m.ExcPending
										if v1191 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_hashtext_1), int32(306), int32(_a_F_hashtext_2))
											mBase = m.M
											v1196 = m.ExcPending
											if v1196 != 0 {
												return int64(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v887 = v883 + int32(1)
									v893 = v887 - int32(1636608432)
									if v881&int32(3) != 0 {
										if base.Ui32(int32(11)) < base.Ui32(v887) {
											v1002 = v881
											v1003 = v887
											v1004 = v893
											v1005 = v893
											v1006 = v893
											for {
												v1008 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+4))
												v1009 = v1008 + v1005
												v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1002)))
												v1012 = *(*int32)(unsafe.Add(mBase, uint32(v1002)+8))
												v1013 = v1012 + v1006
												v1015 = int32(4)
												v1017 = v1010 + v1004 - v1013 ^ base.I32_rotl(v1013, v1015)
												v1021 = v1009 - v1017 ^ base.I32_rotl(v1017, int32(6))
												v1022 = v1013 + v1009
												v1023 = v1017 + v1022
												v1024 = v1021 + v1023
												v1028 = v1022 - v1021 ^ base.I32_rotl(v1021, int32(8))
												v1032 = v1023 - v1028 ^ base.I32_rotl(v1028, int32(16))
												v1036 = v1024 - v1032 ^ base.I32_rotl(v1032, int32(19))
												v1037 = v1028 + v1024
												v1038 = v1032 + v1037
												v1039 = v1036 + v1038
												v1043 = v1037 - v1036 ^ base.I32_rotl(v1036, v1015)
												v1044 = int32(12)
												v1045 = v1002 + v1044
												v1047 = v1003 - v1044
												if base.Ui32(int32(11)) < base.Ui32(v1047) {
													v1002 = v1045
													v1003 = v1047
													v1004 = v1038
													v1005 = v1039
													v1006 = v1043
													continue
												} else {
													break
												}
												break
											}
											v1050 = v1045
											v1051 = v1047
											v1052 = v1038
											v1053 = v1039
											v1054 = v1043
										} else {
											v1050 = v881
											v1051 = v887
											v1052 = v893
											v1053 = v893
											v1054 = v893
										}
										switch v1051 - int32(1) {
										case 0:
											v1113 = v1052
											v1114 = v1053
											v1115 = v1054
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 1:
											v1106 = v1052
											v1107 = v1053
											v1108 = v1054
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 2:
											v1099 = v1052
											v1100 = v1053
											v1101 = v1054
											v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+2)))
											v1106 = v1102<<(uint(int32(16))%32) + v1099
											v1107 = v1100
											v1108 = v1101
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 3:
											v1093 = v1053
											v1094 = v1054
											v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+3)))
											v1099 = v1095<<(uint(int32(24))%32) + v1052
											v1100 = v1093
											v1101 = v1094
											v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+2)))
											v1106 = v1102<<(uint(int32(16))%32) + v1099
											v1107 = v1100
											v1108 = v1101
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 4:
											v1089 = v1053
											v1090 = v1054
											v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+4)))
											v1093 = v1089 + v1091
											v1094 = v1090
											v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+3)))
											v1099 = v1095<<(uint(int32(24))%32) + v1052
											v1100 = v1093
											v1101 = v1094
											v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+2)))
											v1106 = v1102<<(uint(int32(16))%32) + v1099
											v1107 = v1100
											v1108 = v1101
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 5:
											v1083 = v1053
											v1084 = v1054
											v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+5)))
											v1089 = v1085<<(uint(int32(8))%32) + v1083
											v1090 = v1084
											v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+4)))
											v1093 = v1089 + v1091
											v1094 = v1090
											v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+3)))
											v1099 = v1095<<(uint(int32(24))%32) + v1052
											v1100 = v1093
											v1101 = v1094
											v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+2)))
											v1106 = v1102<<(uint(int32(16))%32) + v1099
											v1107 = v1100
											v1108 = v1101
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 6:
											v1077 = v1053
											v1078 = v1054
											v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+6)))
											v1083 = v1079<<(uint(int32(16))%32) + v1077
											v1084 = v1078
											v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+5)))
											v1089 = v1085<<(uint(int32(8))%32) + v1083
											v1090 = v1084
											v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+4)))
											v1093 = v1089 + v1091
											v1094 = v1090
											v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+3)))
											v1099 = v1095<<(uint(int32(24))%32) + v1052
											v1100 = v1093
											v1101 = v1094
											v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+2)))
											v1106 = v1102<<(uint(int32(16))%32) + v1099
											v1107 = v1100
											v1108 = v1101
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 7:
											v1072 = v1054
											v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+7)))
											v1077 = v1073<<(uint(int32(24))%32) + v1053
											v1078 = v1072
											v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+6)))
											v1083 = v1079<<(uint(int32(16))%32) + v1077
											v1084 = v1078
											v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+5)))
											v1089 = v1085<<(uint(int32(8))%32) + v1083
											v1090 = v1084
											v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+4)))
											v1093 = v1089 + v1091
											v1094 = v1090
											v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+3)))
											v1099 = v1095<<(uint(int32(24))%32) + v1052
											v1100 = v1093
											v1101 = v1094
											v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+2)))
											v1106 = v1102<<(uint(int32(16))%32) + v1099
											v1107 = v1100
											v1108 = v1101
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 8:
											v1067 = v1054
											v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+8)))
											v1072 = v1068<<(uint(int32(8))%32) + v1067
											v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+7)))
											v1077 = v1073<<(uint(int32(24))%32) + v1053
											v1078 = v1072
											v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+6)))
											v1083 = v1079<<(uint(int32(16))%32) + v1077
											v1084 = v1078
											v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+5)))
											v1089 = v1085<<(uint(int32(8))%32) + v1083
											v1090 = v1084
											v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+4)))
											v1093 = v1089 + v1091
											v1094 = v1090
											v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+3)))
											v1099 = v1095<<(uint(int32(24))%32) + v1052
											v1100 = v1093
											v1101 = v1094
											v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+2)))
											v1106 = v1102<<(uint(int32(16))%32) + v1099
											v1107 = v1100
											v1108 = v1101
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 9:
											v1062 = v1054
											v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+9)))
											v1067 = v1063<<(uint(int32(16))%32) + v1062
											v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+8)))
											v1072 = v1068<<(uint(int32(8))%32) + v1067
											v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+7)))
											v1077 = v1073<<(uint(int32(24))%32) + v1053
											v1078 = v1072
											v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+6)))
											v1083 = v1079<<(uint(int32(16))%32) + v1077
											v1084 = v1078
											v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+5)))
											v1089 = v1085<<(uint(int32(8))%32) + v1083
											v1090 = v1084
											v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+4)))
											v1093 = v1089 + v1091
											v1094 = v1090
											v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+3)))
											v1099 = v1095<<(uint(int32(24))%32) + v1052
											v1100 = v1093
											v1101 = v1094
											v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+2)))
											v1106 = v1102<<(uint(int32(16))%32) + v1099
											v1107 = v1100
											v1108 = v1101
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										case 10:
											v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+10)))
											v1062 = v1058<<(uint(int32(24))%32) + v1054
											v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+9)))
											v1067 = v1063<<(uint(int32(16))%32) + v1062
											v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+8)))
											v1072 = v1068<<(uint(int32(8))%32) + v1067
											v1073 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+7)))
											v1077 = v1073<<(uint(int32(24))%32) + v1053
											v1078 = v1072
											v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+6)))
											v1083 = v1079<<(uint(int32(16))%32) + v1077
											v1084 = v1078
											v1085 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+5)))
											v1089 = v1085<<(uint(int32(8))%32) + v1083
											v1090 = v1084
											v1091 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+4)))
											v1093 = v1089 + v1091
											v1094 = v1090
											v1095 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+3)))
											v1099 = v1095<<(uint(int32(24))%32) + v1052
											v1100 = v1093
											v1101 = v1094
											v1102 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+2)))
											v1106 = v1102<<(uint(int32(16))%32) + v1099
											v1107 = v1100
											v1108 = v1101
											v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050)+1)))
											v1113 = v1109<<(uint(int32(8))%32) + v1106
											v1114 = v1107
											v1115 = v1108
											v1116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1050))))
											v1120 = v1113 + v1116
											v1121 = v1114
											v1122 = v1115
										default:
											v1120 = v1052
											v1121 = v1053
											v1122 = v1054
										}
									} else {
										if base.Ui32(v887) < base.Ui32(int32(12)) {
											v948 = v881
											v949 = v887
											v950 = v893
											v951 = v893
											v952 = v893
										} else {
											v900 = v881
											v901 = v887
											v902 = v893
											v903 = v893
											v904 = v893
											for {
												v906 = *(*int32)(unsafe.Add(mBase, uint32(v900)+4))
												v907 = v906 + v903
												v908 = *(*int32)(unsafe.Add(mBase, uint32(v900)))
												v910 = *(*int32)(unsafe.Add(mBase, uint32(v900)+8))
												v911 = v910 + v904
												v913 = int32(4)
												v915 = v908 + v902 - v911 ^ base.I32_rotl(v911, v913)
												v919 = v907 - v915 ^ base.I32_rotl(v915, int32(6))
												v920 = v911 + v907
												v921 = v915 + v920
												v922 = v919 + v921
												v926 = v920 - v919 ^ base.I32_rotl(v919, int32(8))
												v930 = v921 - v926 ^ base.I32_rotl(v926, int32(16))
												v934 = v922 - v930 ^ base.I32_rotl(v930, int32(19))
												v935 = v926 + v922
												v936 = v930 + v935
												v937 = v934 + v936
												v941 = v935 - v934 ^ base.I32_rotl(v934, v913)
												v942 = int32(12)
												v943 = v900 + v942
												v945 = v901 - v942
												if base.Ui32(int32(11)) < base.Ui32(v945) {
													v900 = v943
													v901 = v945
													v902 = v936
													v903 = v937
													v904 = v941
													continue
												} else {
													break
												}
												break
											}
											v948 = v943
											v949 = v945
											v950 = v936
											v951 = v937
											v952 = v941
										}
										switch v949 - int32(1) {
										case 0:
											v999 = v950
											v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948))))
											v1120 = v999 + v1000
											v1121 = v951
											v1122 = v952
										case 1:
											v994 = v950
											v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+1)))
											v999 = v995<<(uint(int32(8))%32) + v994
											v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948))))
											v1120 = v999 + v1000
											v1121 = v951
											v1122 = v952
										case 2:
											v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+2)))
											v994 = v990<<(uint(int32(16))%32) + v950
											v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+1)))
											v999 = v995<<(uint(int32(8))%32) + v994
											v1000 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948))))
											v1120 = v999 + v1000
											v1121 = v951
											v1122 = v952
										case 3:
											v987 = v951
											v988 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
											v1120 = v988 + v950
											v1121 = v987
											v1122 = v952
										case 4:
											v984 = v951
											v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+4)))
											v987 = v984 + v985
											v988 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
											v1120 = v988 + v950
											v1121 = v987
											v1122 = v952
										case 5:
											v979 = v951
											v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+5)))
											v984 = v980<<(uint(int32(8))%32) + v979
											v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+4)))
											v987 = v984 + v985
											v988 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
											v1120 = v988 + v950
											v1121 = v987
											v1122 = v952
										case 6:
											v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+6)))
											v979 = v975<<(uint(int32(16))%32) + v951
											v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+5)))
											v984 = v980<<(uint(int32(8))%32) + v979
											v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+4)))
											v987 = v984 + v985
											v988 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
											v1120 = v988 + v950
											v1121 = v987
											v1122 = v952
										case 7:
											v970 = v952
											v971 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
											v973 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
											v1120 = v971 + v950
											v1121 = v973 + v951
											v1122 = v970
										case 8:
											v965 = v952
											v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+8)))
											v970 = v966<<(uint(int32(8))%32) + v965
											v971 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
											v973 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
											v1120 = v971 + v950
											v1121 = v973 + v951
											v1122 = v970
										case 9:
											v960 = v952
											v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+9)))
											v965 = v961<<(uint(int32(16))%32) + v960
											v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+8)))
											v970 = v966<<(uint(int32(8))%32) + v965
											v971 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
											v973 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
											v1120 = v971 + v950
											v1121 = v973 + v951
											v1122 = v970
										case 10:
											v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+10)))
											v960 = v956<<(uint(int32(24))%32) + v952
											v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+9)))
											v965 = v961<<(uint(int32(16))%32) + v960
											v966 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v948)+8)))
											v970 = v966<<(uint(int32(8))%32) + v965
											v971 = *(*int32)(unsafe.Add(mBase, uint32(v948)))
											v973 = *(*int32)(unsafe.Add(mBase, uint32(v948)+4))
											v1120 = v971 + v950
											v1121 = v973 + v951
											v1122 = v970
										default:
											v1120 = v950
											v1121 = v951
											v1122 = v952
										}
									}
									v1125 = int32(14)
									v1127 = v1121 ^ v1122 - base.I32_rotl(v1121, v1125)
									v1131 = v1127 ^ v1120 - base.I32_rotl(v1127, int32(11))
									v1135 = v1131 ^ v1121 - base.I32_rotl(v1131, int32(25))
									v1139 = v1135 ^ v1127 - base.I32_rotl(v1135, int32(16))
									v1143 = v1139 ^ v1131 - base.I32_rotl(v1139, int32(4))
									v1147 = v1143 ^ v1135 - base.I32_rotl(v1143, v1125)
									F_pfree(m, v881)
									mBase = m.M
									v1153 = m.ExcPending
									if v1153 != 0 {
										return int64(0)
									} else {
										v1154 = v1147 ^ v1139 - base.I32_rotl(v1147, int32(24))
										v1158 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
										if v1158 != v10 {
											F_pfree(m, v10)
											mBase = m.M
											v1161 = m.ExcPending
											if v1161 != 0 {
												return int64(0)
											} else {
												return base.I64_extend_i32_u(v1154)
											}
										} else {
											return base.I64_extend_i32_u(v1154)
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
			v1167 = m.ExcPending
			if v1167 != 0 {
				return int64(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v1170 = m.ExcPending
				if v1170 != 0 {
					return int64(0)
				} else {
					F_errmsg(m, int32(_a_F_hashtext_3), int32(0))
					mBase = m.M
					v1174 = m.ExcPending
					if v1174 != 0 {
						return int64(0)
					} else {
						F_errhint(m, int32(_a_F_hashtext_4), int32(0))
						mBase = m.M
						v1178 = m.ExcPending
						if v1178 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_hashtext_1), int32(281), int32(_a_F_hashtext_2))
							mBase = m.M
							v1183 = m.ExcPending
							if v1183 != 0 {
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
		}
	}
}
func F_hashvarlena(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
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
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
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
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
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
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
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
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
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
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v251 int32
	_ = v251
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v311 int32
	_ = v311
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = F_pg_detoast_datum_packed(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int64(0)
	} else {
		v10 = int32(1)
		v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6))))
		v14 = v12 & v10
		if v14 != 0 {
			v15 = v10
		} else {
			v15 = int32(4)
		}
		v16 = v6 + v15
		if v12 == int32(1) {
			v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+1)))
			if v22 == int32(18) {
				v25 = int32(16)
			} else {
				v25 = int32(0)
			}
			if base.Ui32((v22-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v32 = int32(4)
			} else {
				v32 = v25
			}
			v43 = v32
		} else {
			v33 = int32(1)
			if v14 != 0 {
				v43 = int32(base.Ui32(v12)>>(uint(v33)%32)) - v33
			} else {
				v37 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
				v43 = int32(base.Ui32(v37)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v49 = v43 - int32(1636608432)
		if v16&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v43) {
				v158 = v16
				v159 = v43
				v160 = v49
				v161 = v49
				v162 = v49
				for {
					v164 = *(*int32)(unsafe.Add(mBase, uint32(v158)+4))
					v165 = v164 + v161
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v158)))
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v158)+8))
					v169 = v168 + v162
					v171 = int32(4)
					v173 = v166 + v160 - v169 ^ base.I32_rotl(v169, v171)
					v177 = v165 - v173 ^ base.I32_rotl(v173, int32(6))
					v178 = v169 + v165
					v179 = v173 + v178
					v180 = v177 + v179
					v184 = v178 - v177 ^ base.I32_rotl(v177, int32(8))
					v188 = v179 - v184 ^ base.I32_rotl(v184, int32(16))
					v192 = v180 - v188 ^ base.I32_rotl(v188, int32(19))
					v193 = v184 + v180
					v194 = v188 + v193
					v195 = v192 + v194
					v199 = v193 - v192 ^ base.I32_rotl(v192, v171)
					v200 = int32(12)
					v201 = v158 + v200
					v203 = v159 - v200
					if base.Ui32(int32(11)) < base.Ui32(v203) {
						v158 = v201
						v159 = v203
						v160 = v194
						v161 = v195
						v162 = v199
						continue
					} else {
						break
					}
					break
				}
				v206 = v201
				v207 = v203
				v208 = v194
				v209 = v195
				v210 = v199
			} else {
				v206 = v16
				v207 = v43
				v208 = v49
				v209 = v49
				v210 = v49
			}
			switch v207 - int32(1) {
			case 0:
				v269 = v208
				v270 = v209
				v271 = v210
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 1:
				v262 = v208
				v263 = v209
				v264 = v210
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 2:
				v255 = v208
				v256 = v209
				v257 = v210
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v262 = v258<<(uint(int32(16))%32) + v255
				v263 = v256
				v264 = v257
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 3:
				v249 = v209
				v250 = v210
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
				v255 = v251<<(uint(int32(24))%32) + v208
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v262 = v258<<(uint(int32(16))%32) + v255
				v263 = v256
				v264 = v257
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 4:
				v245 = v209
				v246 = v210
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v249 = v245 + v247
				v250 = v246
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
				v255 = v251<<(uint(int32(24))%32) + v208
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v262 = v258<<(uint(int32(16))%32) + v255
				v263 = v256
				v264 = v257
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 5:
				v239 = v209
				v240 = v210
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
				v245 = v241<<(uint(int32(8))%32) + v239
				v246 = v240
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v249 = v245 + v247
				v250 = v246
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
				v255 = v251<<(uint(int32(24))%32) + v208
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v262 = v258<<(uint(int32(16))%32) + v255
				v263 = v256
				v264 = v257
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 6:
				v233 = v209
				v234 = v210
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+6)))
				v239 = v235<<(uint(int32(16))%32) + v233
				v240 = v234
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
				v245 = v241<<(uint(int32(8))%32) + v239
				v246 = v240
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v249 = v245 + v247
				v250 = v246
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
				v255 = v251<<(uint(int32(24))%32) + v208
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v262 = v258<<(uint(int32(16))%32) + v255
				v263 = v256
				v264 = v257
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 7:
				v228 = v210
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+7)))
				v233 = v229<<(uint(int32(24))%32) + v209
				v234 = v228
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+6)))
				v239 = v235<<(uint(int32(16))%32) + v233
				v240 = v234
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
				v245 = v241<<(uint(int32(8))%32) + v239
				v246 = v240
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v249 = v245 + v247
				v250 = v246
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
				v255 = v251<<(uint(int32(24))%32) + v208
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v262 = v258<<(uint(int32(16))%32) + v255
				v263 = v256
				v264 = v257
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 8:
				v223 = v210
				v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+8)))
				v228 = v224<<(uint(int32(8))%32) + v223
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+7)))
				v233 = v229<<(uint(int32(24))%32) + v209
				v234 = v228
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+6)))
				v239 = v235<<(uint(int32(16))%32) + v233
				v240 = v234
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
				v245 = v241<<(uint(int32(8))%32) + v239
				v246 = v240
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v249 = v245 + v247
				v250 = v246
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
				v255 = v251<<(uint(int32(24))%32) + v208
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v262 = v258<<(uint(int32(16))%32) + v255
				v263 = v256
				v264 = v257
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 9:
				v218 = v210
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+9)))
				v223 = v219<<(uint(int32(16))%32) + v218
				v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+8)))
				v228 = v224<<(uint(int32(8))%32) + v223
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+7)))
				v233 = v229<<(uint(int32(24))%32) + v209
				v234 = v228
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+6)))
				v239 = v235<<(uint(int32(16))%32) + v233
				v240 = v234
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
				v245 = v241<<(uint(int32(8))%32) + v239
				v246 = v240
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v249 = v245 + v247
				v250 = v246
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
				v255 = v251<<(uint(int32(24))%32) + v208
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v262 = v258<<(uint(int32(16))%32) + v255
				v263 = v256
				v264 = v257
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			case 10:
				v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+10)))
				v218 = v214<<(uint(int32(24))%32) + v210
				v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+9)))
				v223 = v219<<(uint(int32(16))%32) + v218
				v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+8)))
				v228 = v224<<(uint(int32(8))%32) + v223
				v229 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+7)))
				v233 = v229<<(uint(int32(24))%32) + v209
				v234 = v228
				v235 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+6)))
				v239 = v235<<(uint(int32(16))%32) + v233
				v240 = v234
				v241 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+5)))
				v245 = v241<<(uint(int32(8))%32) + v239
				v246 = v240
				v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+4)))
				v249 = v245 + v247
				v250 = v246
				v251 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+3)))
				v255 = v251<<(uint(int32(24))%32) + v208
				v256 = v249
				v257 = v250
				v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+2)))
				v262 = v258<<(uint(int32(16))%32) + v255
				v263 = v256
				v264 = v257
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206)+1)))
				v269 = v265<<(uint(int32(8))%32) + v262
				v270 = v263
				v271 = v264
				v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v206))))
				v276 = v269 + v272
				v277 = v270
				v278 = v271
			default:
				v276 = v208
				v277 = v209
				v278 = v210
			}
		} else {
			if base.Ui32(v43) < base.Ui32(int32(12)) {
				v104 = v16
				v105 = v43
				v106 = v49
				v107 = v49
				v108 = v49
			} else {
				v56 = v16
				v57 = v43
				v58 = v49
				v59 = v49
				v60 = v49
				for {
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v56)+4))
					v63 = v62 + v59
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v56)))
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v56)+8))
					v67 = v66 + v60
					v69 = int32(4)
					v71 = v64 + v58 - v67 ^ base.I32_rotl(v67, v69)
					v75 = v63 - v71 ^ base.I32_rotl(v71, int32(6))
					v76 = v67 + v63
					v77 = v71 + v76
					v78 = v75 + v77
					v82 = v76 - v75 ^ base.I32_rotl(v75, int32(8))
					v86 = v77 - v82 ^ base.I32_rotl(v82, int32(16))
					v90 = v78 - v86 ^ base.I32_rotl(v86, int32(19))
					v91 = v82 + v78
					v92 = v86 + v91
					v93 = v90 + v92
					v97 = v91 - v90 ^ base.I32_rotl(v90, v69)
					v98 = int32(12)
					v99 = v56 + v98
					v101 = v57 - v98
					if base.Ui32(int32(11)) < base.Ui32(v101) {
						v56 = v99
						v57 = v101
						v58 = v92
						v59 = v93
						v60 = v97
						continue
					} else {
						break
					}
					break
				}
				v104 = v99
				v105 = v101
				v106 = v92
				v107 = v93
				v108 = v97
			}
			switch v105 - int32(1) {
			case 0:
				v155 = v106
				v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
				v276 = v155 + v156
				v277 = v107
				v278 = v108
			case 1:
				v150 = v106
				v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
				v155 = v151<<(uint(int32(8))%32) + v150
				v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
				v276 = v155 + v156
				v277 = v107
				v278 = v108
			case 2:
				v146 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+2)))
				v150 = v146<<(uint(int32(16))%32) + v106
				v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+1)))
				v155 = v151<<(uint(int32(8))%32) + v150
				v156 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104))))
				v276 = v155 + v156
				v277 = v107
				v278 = v108
			case 3:
				v143 = v107
				v144 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
				v276 = v144 + v106
				v277 = v143
				v278 = v108
			case 4:
				v140 = v107
				v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
				v143 = v140 + v141
				v144 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
				v276 = v144 + v106
				v277 = v143
				v278 = v108
			case 5:
				v135 = v107
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
				v140 = v136<<(uint(int32(8))%32) + v135
				v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
				v143 = v140 + v141
				v144 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
				v276 = v144 + v106
				v277 = v143
				v278 = v108
			case 6:
				v131 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+6)))
				v135 = v131<<(uint(int32(16))%32) + v107
				v136 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+5)))
				v140 = v136<<(uint(int32(8))%32) + v135
				v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+4)))
				v143 = v140 + v141
				v144 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
				v276 = v144 + v106
				v277 = v143
				v278 = v108
			case 7:
				v126 = v108
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
				v276 = v127 + v106
				v277 = v129 + v107
				v278 = v126
			case 8:
				v121 = v108
				v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+8)))
				v126 = v122<<(uint(int32(8))%32) + v121
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
				v276 = v127 + v106
				v277 = v129 + v107
				v278 = v126
			case 9:
				v116 = v108
				v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+9)))
				v121 = v117<<(uint(int32(16))%32) + v116
				v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+8)))
				v126 = v122<<(uint(int32(8))%32) + v121
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
				v276 = v127 + v106
				v277 = v129 + v107
				v278 = v126
			case 10:
				v112 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+10)))
				v116 = v112<<(uint(int32(24))%32) + v108
				v117 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+9)))
				v121 = v117<<(uint(int32(16))%32) + v116
				v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v104)+8)))
				v126 = v122<<(uint(int32(8))%32) + v121
				v127 = *(*int32)(unsafe.Add(mBase, uint32(v104)))
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
				v276 = v127 + v106
				v277 = v129 + v107
				v278 = v126
			default:
				v276 = v106
				v277 = v107
				v278 = v108
			}
		}
		v281 = int32(14)
		v283 = v277 ^ v278 - base.I32_rotl(v277, v281)
		v287 = v283 ^ v276 - base.I32_rotl(v283, int32(11))
		v291 = v287 ^ v277 - base.I32_rotl(v287, int32(25))
		v295 = v291 ^ v283 - base.I32_rotl(v291, int32(16))
		v299 = v295 ^ v287 - base.I32_rotl(v295, int32(4))
		v303 = v299 ^ v291 - base.I32_rotl(v299, v281)
		v308 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		if v308 != v6 {
			F_pfree(m, v6)
			mBase = m.M
			v311 = m.ExcPending
			if v311 != 0 {
				return int64(0)
			} else {
				return base.I64_extend_i32_u(v303 ^ v295 - base.I32_rotl(v303, int32(24)))
			}
		} else {
			return base.I64_extend_i32_u(v303 ^ v295 - base.I32_rotl(v303, int32(24)))
		}
	}
}
func F_have_createdb_privilege(m *base.Module) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	v4 = F_superuser(m)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		if v4 == int32(0) {
			v12 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_have_createdb_privilege[0])))
			v13 = F_SearchSysCache1(m, int32(11), v12)
			mBase = m.M
			v14 = m.ExcPending
			if v14 != 0 {
				return int32(0)
			} else {
				if v13 == int32(0) {
					return int32(0)
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v13)+16))
					v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+22)))
					v22 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19+v20)+71)))
					F_ReleaseCatCache(m, v13)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int32(0)
					} else {
						v25 = v22
						return v25 & int32(1)
					}
				}
			}
		} else {
			v25 = int32(1)
			return v25 & int32(1)
		}
	}
}
func F_having_var_grouping_eqop(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v12 != v13 {
		v74 = v3
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v10 + int32(16)
	return v74
L2:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v15 != 0 {
		v74 = v3
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
	if v17 == int32(0) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v35)+8))
	v74 = v71
	goto L1
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L10
	} else {
		goto L17
	}
L6:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v20 <= int32(0) {
		goto L5
	} else {
		goto L7
	}
L7:
	;
	v25 = int32(0)
	v30 = v3
	goto L8
L8:
	;
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v17)+12))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v31+v25<<(uint(int32(2))%32))))
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v16)+76))
	v37 = F_get_sortgroupclause_tle(m, v35, v36)
	mBase = m.M
	v40 = m.ExcPending
	if v40 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L5
L10:
	;
	return int32(0)
L11:
	;
	if v37 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v42 = v30 + int32(1)
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	if v42 == v43 {
		goto L4
	} else {
		goto L15
	}
L13:
	;
	v45 = v30
	goto L14
L14:
	;
	v47 = v25 + int32(1)
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
	if v47 < v48 {
		v25 = v47
		v30 = v45
		goto L8
	} else {
		goto L16
	}
L15:
	;
	v45 = v42
	goto L14
L16:
	;
	goto L9
L17:
	;
	v61 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+8)))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v61
	F_errmsg_internal(m, int32(_a_F_having_var_grouping_eqop_0), v10)
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L10
	} else {
		goto L18
	}
L18:
	;
	F_errfinish(m, int32(_a_F_having_var_grouping_eqop_1), int32(1630), int32(_a_F_having_var_grouping_eqop_2))
	mBase = m.M
	v70 = m.ExcPending
	if v70 != 0 {
		goto L10
	} else {
		goto L19
	}
L19:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_headline_json_value(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v26 int64
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	v9 = int64(*(*uint32)(unsafe.Add(mBase, uint32(l0)+16)))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = int32(0)
	v16 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
	F_hlparsetext(m, v16, v13, v11, l1, l2)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v26 = F_FunctionCall3Coll(m, v10+int32(112), int32(0), base.I64_extend_i32_u(v13), v9, base.I64_extend_i32_u(v11))
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return int32(0)
		} else {
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v28)
			v30 = F_generateHeadline(m, v13)
			mBase = m.M
			v31 = m.ExcPending
			if v31 != 0 {
				return int32(0)
			} else {
				return v30
			}
		}
	}
}
func F_hlCover(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v36 int32
	_ = v36
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v273 int32
	_ = v273
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v328 int32
	_ = v328
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	if l2 == v7 {
		v328 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v328
L2:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v36 = v26
	goto L3
L3:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v44 <= int32(0) {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	v129 = int32(0)
	if v122 < v129 {
		v328 = v129
		goto L1
	} else {
		goto L24
	}
L6:
	;
	v122 = int32(-1)
	goto L5
L7:
	;
	goto L8
L8:
	;
	v48 = int32(0)
	if v48 < v44 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v51 = v44
	goto L11
L10:
	;
	v51 = v48
	goto L11
L11:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v54 = int32(0)
	v64 = v54
	v66 = int32(-1)
	goto L12
L12:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v52+v64<<(uint(int32(2))%32))))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 <= int32(0) {
		v328 = v54
		goto L1
	} else {
		goto L14
	}
L13:
	;
	v122 = v108
	goto L5
L14:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v83 = int32(0)
	goto L16
L15:
	;
	if v66 < v102 {
		goto L20
	} else {
		goto L21
	}
L16:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+v83<<(uint(int32(1))%32)))))
	if v36 <= v102 {
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v328 = v54
	goto L1
L18:
	;
	v105 = v83 + int32(1)
	if v105 != v77 {
		v83 = v105
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v108 = v102
	goto L22
L21:
	;
	v108 = v66
	goto L22
L22:
	;
	v110 = v64 + int32(1)
	if v110 != v51 {
		v64 = v110
		v66 = v108
		goto L12
	} else {
		goto L23
	}
L23:
	;
	goto L13
L24:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v132 <= int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	if v36 < v208 {
		goto L44
	} else {
		goto L45
	}
L26:
	;
	v208 = int32(2147483646)
	goto L25
L27:
	;
	goto L28
L28:
	;
	v136 = int32(0)
	if v136 < v132 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v139 = v132
	goto L31
L30:
	;
	v139 = v136
	goto L31
L31:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v150 = int32(2147483646)
	v151 = int32(0)
	goto L32
L32:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v140+v151<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v166 = v164
	goto L34
L33:
	;
	v208 = v197
	goto L25
L34:
	;
	v184 = v166 - int32(1)
	if v184 < int32(0) {
		v195 = int32(-1)
		goto L36
	} else {
		goto L37
	}
L35:
	;
	if v195 < v150 {
		goto L39
	} else {
		goto L40
	}
L36:
	;
	goto L35
L37:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187+v184<<(uint(int32(1))%32)))))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v193 = v191 - v192
	if v122 < v193 {
		v166 = v184
		goto L34
	} else {
		goto L38
	}
L38:
	;
	v195 = v193
	goto L36
L39:
	;
	v197 = v195
	goto L41
L40:
	;
	v197 = v150
	goto L41
L41:
	;
	v199 = v151 + int32(1)
	if v199 != v139 {
		v150 = v197
		v151 = v199
		goto L32
	} else {
		goto L42
	}
L42:
	;
	goto L33
L43:
	;
	v36 = v219 + int32(1)
	goto L3
L44:
	;
	v219 = v208
	goto L46
L45:
	;
	v219 = v36
	goto L46
L46:
	;
	if v122 < v219 {
		goto L43
	} else {
		goto L47
	}
L47:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v221 <= int32(0) {
		goto L43
	} else {
		goto L48
	}
L48:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v225 = int32(-1)
	v229 = int32(0)
	v234 = v225
	v235 = v225
	goto L49
L49:
	;
	v247 = v224 + v229<<(uint(int32(4))%32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	if v248 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	if base.B2i32(v265 < int32(0))|base.B2i32(v266 < v265) != 0 {
		goto L43
	} else {
		goto L64
	}
L51:
	;
	goto L50
L52:
	;
	v262 = v229 + int32(1)
	if v262 != v221 {
		v229 = v262
		v234 = v258
		v235 = v259
		goto L49
	} else {
		goto L63
	}
L53:
	;
	v258 = v234
	v259 = v235
	goto L52
L54:
	;
	goto L55
L55:
	;
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)))
	if v251 < v219 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v253 = v234
	goto L58
L57:
	;
	v253 = v229
	goto L58
L58:
	;
	if v234 < int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v256 = v253
	goto L61
L60:
	;
	v256 = v234
	goto L61
L61:
	;
	if v122 < v251 {
		v265 = v256
		v266 = v235
		goto L51
	} else {
		goto L62
	}
L62:
	;
	v258 = v256
	v259 = v229
	goto L52
L63:
	;
	v265 = v258
	v266 = v259
	goto L51
L64:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v266 - v265 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v273 + v265<<(uint(int32(4))%32)
	v286 = F_TS_execute(m, l1+int32(8), v20+int32(8), int32(0), int32(1290))
	mBase = m.M
	v289 = m.ExcPending
	if v289 != 0 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	return int32(0)
L66:
	;
	if v286 == int32(0) {
		goto L43
	} else {
		goto L67
	}
L67:
	;
	v292 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v219 + v292
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v265
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v266
	v328 = v292
	goto L1
}
func F_hmac_finish(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
	v9 = m.T0[v8].(func(*base.Module, int32) int32)(m, v7)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return
	} else {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
		v12 = m.T0[v11].(func(*base.Module, int32) int32)(m, v7)
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v14 = F_palloc(m, v12)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
				m.T0[v16].(func(*base.Module, int32, int32))(m, v7, v14)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return
				} else {
					v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+8))
					m.T0[v19].(func(*base.Module, int32))(m, v7)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
						v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
						m.T0[v23].(func(*base.Module, int32, int32, int32))(m, v7, v22, v9)
						mBase = m.M
						v25 = m.ExcPending
						if v25 != 0 {
							return
						} else {
							v26 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
							m.T0[v26].(func(*base.Module, int32, int32, int32))(m, v7, v14, v12)
							mBase = m.M
							v28 = m.ExcPending
							if v28 != 0 {
								return
							} else {
								v29 = *(*int32)(unsafe.Add(mBase, uint32(v7)+16))
								m.T0[v29].(func(*base.Module, int32, int32))(m, v7, l1)
								mBase = m.M
								v31 = m.ExcPending
								if v31 != 0 {
									return
								} else {
									if v12 != 0 {
										base.MemoryFill(m, v14, int32(0), v12)
									} else {
									}
									F_pfree(m, v14)
									mBase = m.M
									v35 = m.ExcPending
									if v35 != 0 {
										return
									} else {
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
func F_hmac_result_size(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
	v4 = m.T0[v3].(func(*base.Module, int32) int32)(m, v2)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_hnswbuildempty(m *base.Module, l0 int32) {
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	v3 = m.G0
	v5 = v3 - int32(208)
	m.G0 = v5
	v8 = F_BuildIndexInfo(m, l0)
	v9 = m.ExcPending
	if v9 != 0 {
		return
	} else {
		F_BuildIndex_1(m, int32(0), l0, v8, v5, int32(3))
		v12 = m.ExcPending
		if v12 != 0 {
			return
		} else {
			m.G0 = v5 + int32(208)
			return
		}
	}
}
func F_hnswbulkdelete(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v160 int32
	_ = v160
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 float64
	_ = v225
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 float64
	_ = v237
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v256 int32
	_ = v256
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v326 int32
	_ = v326
	var v351 int32
	_ = v351
	var v359 int32
	_ = v359
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v402 int32
	_ = v402
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v424 int32
	_ = v424
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v440 int32
	_ = v440
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v458 int32
	_ = v458
	var v468 int32
	_ = v468
	var v470 int32
	_ = v470
	var v474 int32
	_ = v474
	var v477 int32
	_ = v477
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v499 int32
	_ = v499
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v542 int32
	_ = v542
	var v544 int32
	_ = v544
	var v548 int32
	_ = v548
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v589 int32
	_ = v589
	var v591 int32
	_ = v591
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v601 int32
	_ = v601
	var v602 int32
	_ = v602
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v622 int64
	_ = v622
	var v623 int64
	_ = v623
	var v624 int32
	_ = v624
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v631 int64
	_ = v631
	var v632 int64
	_ = v632
	var v637 int64
	_ = v637
	var v642 int64
	_ = v642
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v678 int32
	_ = v678
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v697 int32
	_ = v697
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v709 int32
	_ = v709
	var v715 int32
	_ = v715
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v727 int32
	_ = v727
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v764 int32
	_ = v764
	var v770 int32
	_ = v770
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v783 int32
	_ = v783
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v821 int32
	_ = v821
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v828 int32
	_ = v828
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v847 int32
	_ = v847
	var v868 int32
	_ = v868
	var v870 int32
	_ = v870
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v879 int32
	_ = v879
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v908 int32
	_ = v908
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v919 int32
	_ = v919
	var v920 int32
	_ = v920
	var v923 int32
	_ = v923
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v933 int32
	_ = v933
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v942 int32
	_ = v942
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v959 int32
	_ = v959
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v1001 int32
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1015 int32
	_ = v1015
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1037 int32
	_ = v1037
	var v1038 int32
	_ = v1038
	var v1041 int32
	_ = v1041
	var v1045 int32
	_ = v1045
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1068 int32
	_ = v1068
	var v1076 int32
	_ = v1076
	var v1102 int32
	_ = v1102
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1125 int32
	_ = v1125
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1157 int32
	_ = v1157
	var v1160 int32
	_ = v1160
	var v1182 int32
	_ = v1182
	var v1183 int32
	_ = v1183
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1188 int32
	_ = v1188
	var v1190 int32
	_ = v1190
	var v1195 int32
	_ = v1195
	var v1197 int32
	_ = v1197
	var v1200 int32
	_ = v1200
	var v1201 int64
	_ = v1201
	var v1202 int64
	_ = v1202
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1207 int32
	_ = v1207
	var v1208 int32
	_ = v1208
	var v1210 int64
	_ = v1210
	var v1211 int64
	_ = v1211
	var v1216 int64
	_ = v1216
	var v1221 int64
	_ = v1221
	var v1226 int32
	_ = v1226
	var v1229 int32
	_ = v1229
	var v1230 int32
	_ = v1230
	var v1232 int32
	_ = v1232
	var v1236 int32
	_ = v1236
	var v1239 int32
	_ = v1239
	var v1240 int32
	_ = v1240
	var v1241 int32
	_ = v1241
	var v1244 int32
	_ = v1244
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1257 int32
	_ = v1257
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1268 int32
	_ = v1268
	var v1297 int32
	_ = v1297
	var v1322 int32
	_ = v1322
	var v1349 int32
	_ = v1349
	var v1351 int32
	_ = v1351
	var v1353 int32
	_ = v1353
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1391 int32
	_ = v1391
	var v1392 int32
	_ = v1392
	var v1394 int32
	_ = v1394
	var v1395 int32
	_ = v1395
	var v1397 int32
	_ = v1397
	var v1398 int32
	_ = v1398
	var v1399 int32
	_ = v1399
	var v1401 int32
	_ = v1401
	var v1402 int32
	_ = v1402
	var v1403 int32
	_ = v1403
	var v1407 int32
	_ = v1407
	var v1417 int32
	_ = v1417
	var v1420 int32
	_ = v1420
	var v1426 int32
	_ = v1426
	var v1428 int32
	_ = v1428
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1447 int32
	_ = v1447
	var v1450 int32
	_ = v1450
	var v1453 int32
	_ = v1453
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1458 int32
	_ = v1458
	var v1459 int32
	_ = v1459
	var v1462 int32
	_ = v1462
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1467 int32
	_ = v1467
	var v1470 int32
	_ = v1470
	var v1472 int32
	_ = v1472
	var v1473 int32
	_ = v1473
	var v1474 int32
	_ = v1474
	var v1475 int32
	_ = v1475
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1486 int32
	_ = v1486
	var v1488 int32
	_ = v1488
	var v1491 int32
	_ = v1491
	var v1498 int32
	_ = v1498
	var v1499 int32
	_ = v1499
	var v1503 int32
	_ = v1503
	var v1507 int32
	_ = v1507
	var v1512 int32
	_ = v1512
	var v1513 int32
	_ = v1513
	var v1517 int32
	_ = v1517
	var v1518 int32
	_ = v1518
	var v1519 int32
	_ = v1519
	var v1520 int32
	_ = v1520
	var v1527 int32
	_ = v1527
	var v1535 int32
	_ = v1535
	var v1552 int32
	_ = v1552
	var v1559 int32
	_ = v1559
	var v1560 int32
	_ = v1560
	var v1562 int32
	_ = v1562
	var v1566 int32
	_ = v1566
	var v1589 int32
	_ = v1589
	var v1592 int32
	_ = v1592
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1619 int32
	_ = v1619
	var v1622 int32
	_ = v1622
	var v1647 int32
	_ = v1647
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1655 int32
	_ = v1655
	var v1659 int32
	_ = v1659
	var v1662 int32
	_ = v1662
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1680 int32
	_ = v1680
	var v1682 int32
	_ = v1682
	var v1697 int32
	_ = v1697
	var v1706 int32
	_ = v1706
	var v1708 int32
	_ = v1708
	var v1720 int32
	_ = v1720
	var v1722 int32
	_ = v1722
	var v1724 int32
	_ = v1724
	var v1726 int32
	_ = v1726
	var v1729 int32
	_ = v1729
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1737 int32
	_ = v1737
	var v1738 int32
	_ = v1738
	var v1740 int32
	_ = v1740
	var v1741 int32
	_ = v1741
	var v1743 int32
	_ = v1743
	var v1744 int32
	_ = v1744
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1755 int32
	_ = v1755
	var v1759 int32
	_ = v1759
	var v1764 int32
	_ = v1764
	v24 = m.G0
	v26 = v24 - int32(304)
	m.G0 = v26
	v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if l1 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v32 = F_palloc0(m, int32(40))
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	v36 = l1
	goto L3
L3:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v26)+36)) = l2
	*(*int32)(unsafe.Add(mBase, uint32(v26)+32)) = v36
	*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v28
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v28)+180))
	if v41 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	return int32(0)
L5:
	;
	v36 = v32
	goto L3
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+48)) = v46
	v49 = F_GetAccessStrategy(m, int32(1))
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L4
	} else {
		goto L10
	}
L7:
	;
	v46 = int32(64)
	goto L6
L8:
	;
	goto L9
L9:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v46 = v45
	goto L6
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+68)) = v49
	v53 = F_palloc0(m, int32(_a_F_hnswbulkdelete_0))
	mBase = m.M
	v54 = m.ExcPending
	if v54 != 0 {
		goto L4
	} else {
		goto L11
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+72)) = v53
	v57 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[0]))
	v62 = F_AllocSetContextCreateInternal(m, v57, int32(_a_F_hnswbulkdelete_1), int32(0), int32(_a_F_hnswbulkdelete_0), int32(_a_F_hnswbulkdelete_2))
	mBase = m.M
	v63 = m.ExcPending
	if v63 != 0 {
		goto L4
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+292)) = v62
	v66 = v26 + int32(52)
	F_HnswInitSupport(m, v66, v28)
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L4
	} else {
		goto L13
	}
L13:
	;
	F_HnswGetMetaPageInfo(m, v28, v26+int32(44), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L4
	} else {
		goto L14
	}
L14:
	;
	v75 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[0]))
	v78 = F_tidhash_create(m, v75, int32(256), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L4
	} else {
		goto L15
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+64)) = v78
	v81 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+264)) = uint16(v81)
	v83 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+260)) = v83
	*(*uint16)(unsafe.Add(mBase, uint32(v26)+156)) = uint16(v81)
	*(*int32)(unsafe.Add(mBase, uint32(v26)+152)) = v83
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v26)+32))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v26)+68))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v26)+28))
	v97 = v83
	v99 = v26
	v100 = int32(1)
	v101 = v83
	v106 = v89
	v107 = v91
	v108 = v66
	v113 = v90
	goto L16
L16:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L4
	} else {
		goto L18
	}
L17:
	;
	v527 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	v528 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	F_LockPage(m, v528, int32(0), int32(7))
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L4
	} else {
		goto L78
	}
L18:
	;
	v121 = int32(0)
	v123 = F_ReadBufferExtended(m, v107, v121, v100, v121, v113)
	mBase = m.M
	v124 = m.ExcPending
	if v124 != 0 {
		goto L4
	} else {
		goto L19
	}
L19:
	;
	F_LockBufferInternal(m, v123, int32(3))
	mBase = m.M
	v127 = m.ExcPending
	if v127 != 0 {
		goto L4
	} else {
		goto L20
	}
L20:
	;
	v128 = F_GenericXLogStart(m, v107)
	mBase = m.M
	v129 = m.ExcPending
	if v129 != 0 {
		goto L4
	} else {
		goto L24
	}
L21:
	;
	F_UnlockReleaseBuffer(m, v123)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L4
	} else {
		goto L76
	}
L22:
	;
	F_pfree(m, v128)
	mBase = m.M
	v499 = m.ExcPending
	if v499 != 0 {
		goto L4
	} else {
		goto L75
	}
L23:
	;
	v145 = int32(base.Ui32(v100) >> (uint(int32(16)) % 32))
	v154 = int32(1)
	v156 = v97
	v160 = v101
	v170 = int32(0)
	goto L30
L24:
	;
	v131 = F_GenericXLogRegisterBuffer(m, v128, v123, int32(0))
	mBase = m.M
	v132 = m.ExcPending
	if v132 != 0 {
		goto L4
	} else {
		goto L25
	}
L25:
	;
	v133 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+12)))
	if base.Ui32(int32(25)) <= base.Ui32(v133) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v137 = v133 + int32(_a_F_hnswbulkdelete_3)
	if v137&int32(_a_F_hnswbulkdelete_4) != 0 {
		goto L23
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	v141 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+16)))
	v143 = *(*int32)(unsafe.Add(mBase, uint32(v131+v141)))
	v477 = v97
	v480 = v143
	v481 = v101
	goto L22
L29:
	;
	goto L28
L30:
	;
	v180 = *(*int32)(unsafe.Add(mBase, uint32(v131+int32(20)+v154<<(uint(int32(2))%32))))
	v183 = v131 + v180&int32(_a_F_hnswbulkdelete_5)
	v184 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
	if v184 != int32(1) {
		v444 = v156
		v448 = v160
		v458 = v170
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v468 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v131)+16)))
	v470 = *(*int32)(unsafe.Add(mBase, uint32(v131+v468)))
	if v458 == int32(0) {
		v477 = v444
		v480 = v470
		v481 = v448
		goto L22
	} else {
		goto L73
	}
L32:
	;
	if v154 != int32(base.Ui32(v137)>>(uint(int32(2))%32))&int32(_a_F_hnswbulkdelete_6) {
		v154 = v154 + int32(1)
		v156 = v444
		v160 = v448
		v170 = v458
		goto L30
	} else {
		goto L72
	}
L33:
	;
	v187 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+2)))
	if v187 != 0 {
		v444 = v156
		v448 = v160
		v458 = v170
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v188 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+8)))
	if v188 != 0 {
		goto L36
	} else {
		goto L37
	}
L35:
	;
	v422 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	if v156 < v422 {
		goto L65
	} else {
		goto L66
	}
L36:
	;
	v190 = v183 + int32(4)
	v191 = int32(0)
	v194 = v191
	v197 = v191
	v211 = v191
	goto L39
L37:
	;
	v402 = v170
	goto L38
L38:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+300)) = uint16(v100)
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+298)) = uint16(v145)
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+302)) = uint16(v154)
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+24)) = uint16(v154)
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v99)+298))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+20)) = v413
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v99)+64))
	v420 = F_tidhash_insert(m, v415, v99+int32(20), v99+int32(297))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L4
	} else {
		goto L64
	}
L39:
	;
	v219 = v190 + v194*int32(6)
	v220 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+4)))
	if v220 != 0 {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	if v252 == int32(0) {
		v378 = v170
		goto L50
	} else {
		goto L51
	}
L41:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v99)+40))
	v222 = *(*int32)(unsafe.Add(mBase, uint32(v99)+36))
	v223 = m.T0[v222].(func(*base.Module, int32, int32) int32)(m, v219, v221)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L4
	} else {
		goto L45
	}
L42:
	;
	v251 = v197
	v252 = v211
	goto L43
L43:
	;
	goto L40
L44:
	;
	v247 = v194 + int32(1)
	if v247 != int32(10) {
		v194 = v247
		v197 = v243
		v211 = v244
		goto L39
	} else {
		goto L49
	}
L45:
	;
	if v223 != 0 {
		goto L46
	} else {
		goto L47
	}
L46:
	;
	v225 = *(*float64)(unsafe.Add(mBase, uint32(v106)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v106)+16)) = base.F64_add(v225, float64(1))
	v243 = v197
	v244 = int32(1)
	goto L44
L47:
	;
	goto L48
L48:
	;
	v232 = v190 + v197*int32(6)
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v219)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v232)+4)) = uint16(v233)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v219)))
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v235
	v237 = *(*float64)(unsafe.Add(mBase, uint32(v106)+8))
	*(*float64)(unsafe.Add(mBase, uint32(v106)+8)) = base.F64_add(v237, float64(1))
	v243 = v197 + int32(1)
	v244 = v211
	goto L44
L49:
	;
	v251 = v243
	v252 = v244
	goto L43
L50:
	;
	v385 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v183)+8)))
	if v385 != 0 {
		goto L35
	} else {
		goto L63
	}
L51:
	;
	v256 = int32(1)
	if int32(9) < v251 {
		v378 = v256
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v263 = (int32(2) - v251) & int32(3)
	if v263 != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v264 = v251
	v265 = int32(0)
	goto L56
L54:
	;
	v299 = v251
	goto L55
L55:
	;
	if base.Ui32(v251-int32(7)) < base.Ui32(int32(3)) {
		v378 = v256
		goto L50
	} else {
		goto L59
	}
L56:
	;
	v289 = v190 + v264*int32(6)
	v290 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v289)+4)) = uint16(v290)
	*(*int32)(unsafe.Add(mBase, uint32(v289))) = int32(-1)
	v294 = int32(1)
	v295 = v264 + v294
	v297 = v265 + v294
	if v297 != v263 {
		v264 = v295
		v265 = v297
		goto L56
	} else {
		goto L58
	}
L57:
	;
	v299 = v295
	goto L55
L58:
	;
	goto L57
L59:
	;
	v326 = v299
	goto L60
L60:
	;
	v351 = v190 + v326*int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v351))) = int64(-281470681743361)
	*(*int64)(unsafe.Add(mBase, uint32(v351)+16)) = int64(281474976645120)
	*(*int64)(unsafe.Add(mBase, uint32(v351)+8)) = int64(-4294901761)
	v359 = v326 + int32(4)
	if v359 != int32(10) {
		v326 = v359
		goto L60
	} else {
		goto L62
	}
L61:
	;
	v378 = v256
	goto L50
L62:
	;
	goto L61
L63:
	;
	v402 = v378
	goto L38
L64:
	;
	v444 = v156
	v448 = v160
	v458 = v402
	goto L32
L65:
	;
	v424 = *(*int32)(unsafe.Add(mBase, uint32(v99)+152))
	if v424 != int32(-1) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L67
L67:
	;
	if v422 <= v160 {
		v444 = v156
		v448 = v160
		v458 = v378
		goto L32
	} else {
		goto L71
	}
L68:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v99)+260)) = v424
	v428 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+156)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+264)) = uint16(v428)
	v430 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v99)+141)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+249)) = uint8(v430)
	v432 = v156
	goto L70
L69:
	;
	v432 = v160
	goto L70
L70:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+156)) = uint16(v154)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+152)) = v100
	v435 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+141)) = uint8(v435)
	v444 = v435
	v448 = v432
	v458 = v378
	goto L32
L71:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+264)) = uint16(v154)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+260)) = v100
	v440 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
	*(*uint8)(unsafe.Add(mBase, uint32(v99)+249)) = uint8(v440)
	v444 = v156
	v448 = v440
	v458 = v378
	goto L32
L72:
	;
	goto L31
L73:
	;
	F_GenericXLogFinish(m, v128)
	mBase = m.M
	v474 = m.ExcPending
	if v474 != 0 {
		goto L4
	} else {
		goto L74
	}
L74:
	;
	v502 = v444
	v505 = v470
	v506 = v448
	goto L21
L75:
	;
	v502 = v477
	v505 = v480
	v506 = v481
	goto L21
L76:
	;
	if v505 != int32(-1) {
		v97 = v502
		v100 = v505
		v101 = v506
		goto L16
	} else {
		goto L77
	}
L77:
	;
	goto L17
L78:
	;
	F_UnlockPage(m, v528, int32(0), int32(7))
	mBase = m.M
	v536 = m.ExcPending
	if v536 != 0 {
		goto L4
	} else {
		goto L79
	}
L79:
	;
	v537 = int32(_a_F_hnswbulkdelete_7)
	v538 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[0]))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v99)+292))
	*(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[0])) = v540
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v99)+152))
	if v544 != int32(-1) {
		goto L80
	} else {
		goto L81
	}
L80:
	;
	v548 = v99 + int32(76)
	F_LockPage(m, v542, int32(0), int32(5))
	mBase = m.M
	v552 = m.ExcPending
	if v552 != 0 {
		goto L4
	} else {
		goto L83
	}
L81:
	;
	v591 = int32(0)
	goto L82
L82:
	;
	F_LockPage(m, v542, int32(0), int32(7))
	mBase = m.M
	v596 = m.ExcPending
	if v596 != 0 {
		goto L4
	} else {
		goto L97
	}
L83:
	;
	v553 = F_HnswGetEntryPoint(m, v542)
	mBase = m.M
	v554 = m.ExcPending
	if v554 != 0 {
		goto L4
	} else {
		goto L87
	}
L84:
	;
	F_UnlockPage(m, v542, int32(0), int32(5))
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L4
	} else {
		goto L96
	}
L85:
	;
	v584 = int32(0)
	goto L84
L86:
	;
	v569 = int32(0)
	F_HnswLoadElement(m, v568, v569, v569, v542, v108, int32(1), v569)
	mBase = m.M
	v574 = m.ExcPending
	if v574 != 0 {
		goto L4
	} else {
		goto L92
	}
L87:
	;
	if v553 == int32(0) {
		v568 = v548
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(v553)+76))
	v558 = *(*int32)(unsafe.Add(mBase, uint32(v99)+152))
	if v557 != v558 {
		v568 = v548
		goto L86
	} else {
		goto L89
	}
L89:
	;
	v560 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v553)+80)))
	v561 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v99)+156)))
	if v560 != v561 {
		v568 = v548
		goto L86
	} else {
		goto L90
	}
L90:
	;
	v563 = *(*int32)(unsafe.Add(mBase, uint32(v99)+260))
	if v563 == int32(-1) {
		goto L85
	} else {
		goto L91
	}
L91:
	;
	v568 = v99 + int32(184)
	goto L86
L92:
	;
	v576 = v99 + int32(28)
	v577 = F_NeedsUpdated(m, v576, v568)
	mBase = m.M
	v578 = m.ExcPending
	if v578 != 0 {
		goto L4
	} else {
		goto L93
	}
L93:
	;
	if v577 == int32(0) {
		v584 = v568
		goto L84
	} else {
		goto L94
	}
L94:
	;
	F_RepairGraphElement(m, v576, v568, v553)
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L4
	} else {
		goto L95
	}
L95:
	;
	v584 = v568
	goto L84
L96:
	;
	v591 = v584
	goto L82
L97:
	;
	v597 = F_HnswGetEntryPoint(m, v542)
	mBase = m.M
	v598 = m.ExcPending
	if v598 != 0 {
		goto L4
	} else {
		goto L99
	}
L98:
	;
	F_UnlockPage(m, v542, int32(0), int32(7))
	mBase = m.M
	v715 = m.ExcPending
	if v715 != 0 {
		goto L4
	} else {
		goto L121
	}
L99:
	;
	if v597 == int32(0) {
		goto L98
	} else {
		goto L100
	}
L100:
	;
	v601 = *(*int32)(unsafe.Add(mBase, uint32(v597)+76))
	v602 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v597)+80)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+302)) = uint16(v602)
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+16)) = uint16(v602)
	v605 = int32(16)
	v606 = base.I32_rotr(v601, v605)
	*(*int32)(unsafe.Add(mBase, uint32(v99)+298)) = v606
	*(*int32)(unsafe.Add(mBase, uint32(v99)+12)) = v606
	v609 = *(*int32)(unsafe.Add(mBase, uint32(v99)+64))
	v611 = v99 + int32(12)
	v616 = m.G0
	v618 = v616 - v605
	m.G0 = v618
	v621 = v99 + v605
	v622 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v621))))
	v623 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v611))))
	v624 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v621))))
	*(*uint16)(unsafe.Add(mBase, uint32(v618)+12)) = uint16(v624)
	v626 = *(*int32)(unsafe.Add(mBase, uint32(v611)))
	*(*int32)(unsafe.Add(mBase, uint32(v618)+8)) = v626
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v609)+20))
	v629 = *(*int32)(unsafe.Add(mBase, uint32(v609)+12))
	v631 = v622 << (uint(int64(32)) % 64)
	v632 = int64(33)
	v637 = (int64(base.Ui64(v631)>>(uint(v632)%64)) ^ (v631 | v623)) * int64(-49064778989728563)
	v642 = (int64(base.Ui64(v637)>>(uint(v632)%64)) ^ v637) * int64(-4265267296055464877)
	v647 = v629 & base.I32_wrap_i64(int64(base.Ui64(v642)>>(uint(v632)%64))^v642)
	v650 = v628 + v647<<(uint(int32(3))%32)
	v651 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v650)+6)))
	if v651 != 0 {
		goto L103
	} else {
		goto L104
	}
L101:
	;
	if v678 != 0 {
		goto L110
	} else {
		goto L111
	}
L102:
	;
	m.G0 = v618 + int32(16)
	goto L101
L103:
	;
	v653 = v650
	v657 = v647
	goto L106
L104:
	;
	goto L105
L105:
	;
	v678 = int32(0)
	goto L102
L106:
	;
	v660 = F_ItemPointerEquals(m, v653, v618+int32(8))
	mBase = m.M
	if v660 != 0 {
		v678 = v653
		goto L102
	} else {
		goto L108
	}
L107:
	;
	goto L105
L108:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v609)+20))
	v662 = *(*int32)(unsafe.Add(mBase, uint32(v609)+12))
	v665 = v662 & (v657 + int32(1))
	v668 = v661 + v665<<(uint(int32(3))%32)
	v669 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v668)+6)))
	if v669 != 0 {
		v653 = v668
		v657 = v665
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	v688 = int32(0)
	F_HnswUpdateMetaPage(m, v542, int32(2), v591, int32(-1), v688, v688)
	mBase = m.M
	v691 = m.ExcPending
	if v691 != 0 {
		goto L4
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v692 = int32(0)
	F_HnswLoadElement(m, v597, v692, v692, v542, v108, int32(1), v692)
	mBase = m.M
	v697 = m.ExcPending
	if v697 != 0 {
		goto L4
	} else {
		goto L114
	}
L113:
	;
	goto L98
L114:
	;
	v700 = F_NeedsUpdated(m, v99+int32(28), v597)
	mBase = m.M
	v701 = m.ExcPending
	if v701 != 0 {
		goto L4
	} else {
		goto L115
	}
L115:
	;
	if v700 == int32(0) {
		goto L98
	} else {
		goto L116
	}
L116:
	;
	if v591 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v591)+72)) = int32(0)
	goto L119
L118:
	;
	goto L119
L119:
	;
	F_RepairGraphElement(m, v99+int32(28), v597, v591)
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L4
	} else {
		goto L120
	}
L120:
	;
	goto L98
L121:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[0])) = v538
	v718 = *(*int32)(unsafe.Add(mBase, uint32(v99)+292))
	F_MemoryContextReset(m, v718)
	mBase = m.M
	v720 = m.ExcPending
	if v720 != 0 {
		goto L4
	} else {
		goto L122
	}
L122:
	;
	v727 = int32(1)
	goto L123
L123:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L4
	} else {
		goto L125
	}
L124:
	;
	v996 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	v997 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	F_LockPage(m, v997, int32(0), int32(7))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L4
	} else {
		goto L175
	}
L125:
	;
	v748 = int32(_a_F_hnswbulkdelete_7)
	v749 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[0]))
	v751 = *(*int32)(unsafe.Add(mBase, uint32(v99)+292))
	*(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[0])) = v751
	v753 = int32(0)
	v755 = F_ReadBufferExtended(m, v528, v753, v727, v753, v527)
	mBase = m.M
	v756 = m.ExcPending
	if v756 != 0 {
		goto L4
	} else {
		goto L126
	}
L126:
	;
	F_LockBufferInternal(m, v755, int32(1))
	mBase = m.M
	v759 = m.ExcPending
	if v759 != 0 {
		goto L4
	} else {
		goto L127
	}
L127:
	;
	v760 = int32(0)
	if v755 < v760 {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	v868 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v778)+16)))
	v870 = *(*int32)(unsafe.Add(mBase, uint32(v778+v868)))
	F_UnlockReleaseBuffer(m, v755)
	mBase = m.M
	v872 = m.ExcPending
	if v872 != 0 {
		goto L4
	} else {
		goto L145
	}
L129:
	;
	v779 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v778)+12)))
	if base.Ui32(v779) < base.Ui32(int32(25)) {
		v847 = v760
		goto L128
	} else {
		goto L133
	}
L130:
	;
	v764 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[1]))
	v770 = *(*int32)(unsafe.Add(mBase, uint32(v764+(v755^int32(-1))<<(uint(int32(2))%32))))
	v778 = v770
	goto L129
L131:
	;
	goto L132
L132:
	;
	v772 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[2]))
	v778 = v772 + v755<<(uint(int32(13))%32) + int32(-8192)
	goto L129
L133:
	;
	v783 = v779 + int32(_a_F_hnswbulkdelete_3)
	if v783&int32(_a_F_hnswbulkdelete_4) == int32(0) {
		v847 = v760
		goto L128
	} else {
		goto L134
	}
L134:
	;
	v797 = v760
	v798 = int32(1)
	goto L135
L135:
	;
	v821 = *(*int32)(unsafe.Add(mBase, uint32(v778+int32(20)+v798<<(uint(int32(2))%32))))
	v824 = v778 + v821&int32(_a_F_hnswbulkdelete_5)
	v825 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824))))
	if v825 != int32(1) {
		v840 = v797
		goto L137
	} else {
		goto L138
	}
L136:
	;
	v847 = v840
	goto L128
L137:
	;
	if v798 != int32(base.Ui32(v783)>>(uint(int32(2))%32))&int32(_a_F_hnswbulkdelete_6) {
		v797 = v840
		v798 = v798 + int32(1)
		goto L135
	} else {
		goto L144
	}
L138:
	;
	v828 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v824)+2)))
	if v828 != 0 {
		v840 = v797
		goto L137
	} else {
		goto L139
	}
L139:
	;
	v829 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v824)+8)))
	if v829 == int32(0) {
		v840 = v797
		goto L137
	} else {
		goto L140
	}
L140:
	;
	v832 = F_HnswInitElementFromBlock(m, v727, v798)
	mBase = m.M
	v833 = m.ExcPending
	if v833 != 0 {
		goto L4
	} else {
		goto L141
	}
L141:
	;
	F_HnswLoadElementFromTuple(m, v832, v824, int32(0), int32(1))
	mBase = m.M
	v837 = m.ExcPending
	if v837 != 0 {
		goto L4
	} else {
		goto L142
	}
L142:
	;
	v838 = F_lappend(m, v797, v832)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L4
	} else {
		goto L143
	}
L143:
	;
	v840 = v838
	goto L137
L144:
	;
	goto L136
L145:
	;
	if v847 == int32(0) {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[0])) = v749
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v99)+292))
	F_MemoryContextReset(m, v991)
	mBase = m.M
	v993 = m.ExcPending
	if v993 != 0 {
		goto L4
	} else {
		goto L173
	}
L147:
	;
	v875 = int32(0)
	v876 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	if v876 <= v875 {
		goto L146
	} else {
		goto L148
	}
L148:
	;
	v879 = v875
	goto L149
L149:
	;
	v903 = v99 + int32(28)
	v904 = *(*int32)(unsafe.Add(mBase, uint32(v847)+12))
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v904+v879<<(uint(int32(2))%32))))
	v909 = F_NeedsUpdated(m, v903, v908)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L4
	} else {
		goto L151
	}
L150:
	;
	goto L146
L151:
	;
	if v909 != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	F_LockPage(m, v528, int32(0), int32(5))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L4
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	v963 = v879 + int32(1)
	v964 = *(*int32)(unsafe.Add(mBase, uint32(v847)+4))
	if v963 < v964 {
		v879 = v963
		goto L149
	} else {
		goto L172
	}
L155:
	;
	v915 = F_HnswGetEntryPoint(m, v528)
	mBase = m.M
	v916 = m.ExcPending
	if v916 != 0 {
		goto L4
	} else {
		goto L160
	}
L156:
	;
	F_UnlockPage(m, v528, int32(0), v955)
	mBase = m.M
	v959 = m.ExcPending
	if v959 != 0 {
		goto L4
	} else {
		goto L171
	}
L157:
	;
	v951 = int32(0)
	F_HnswUpdateMetaPage(m, v528, int32(1), v908, int32(-1), v951, v951)
	mBase = m.M
	v954 = m.ExcPending
	if v954 != 0 {
		goto L4
	} else {
		goto L170
	}
L158:
	;
	v944 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908)+65)))
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+65)))
	if base.Ui32(v944) <= base.Ui32(v945) {
		v955 = v942
		goto L156
	} else {
		goto L169
	}
L159:
	;
	F_UnlockPage(m, v528, int32(0), int32(5))
	mBase = m.M
	v928 = m.ExcPending
	if v928 != 0 {
		goto L4
	} else {
		goto L164
	}
L160:
	;
	if v915 == int32(0) {
		goto L159
	} else {
		goto L161
	}
L161:
	;
	v919 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v908)+65)))
	v920 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v915)+65)))
	if base.Ui32(v920) < base.Ui32(v919) {
		goto L159
	} else {
		goto L162
	}
L162:
	;
	F_RepairGraphElement(m, v903, v908, v915)
	mBase = m.M
	v923 = m.ExcPending
	if v923 != 0 {
		goto L4
	} else {
		goto L163
	}
L163:
	;
	v942 = int32(5)
	v943 = v915
	goto L158
L164:
	;
	v929 = int32(7)
	F_LockPage(m, v528, int32(0), v929)
	mBase = m.M
	v933 = m.ExcPending
	if v933 != 0 {
		goto L4
	} else {
		goto L165
	}
L165:
	;
	v936 = F_HnswGetEntryPoint(m, v528)
	mBase = m.M
	v937 = m.ExcPending
	if v937 != 0 {
		goto L4
	} else {
		goto L166
	}
L166:
	;
	F_RepairGraphElement(m, v99+int32(28), v908, v936)
	mBase = m.M
	v939 = m.ExcPending
	if v939 != 0 {
		goto L4
	} else {
		goto L167
	}
L167:
	;
	if v936 == int32(0) {
		v947 = v929
		goto L157
	} else {
		goto L168
	}
L168:
	;
	v942 = v929
	v943 = v936
	goto L158
L169:
	;
	v947 = v942
	goto L157
L170:
	;
	v955 = v947
	goto L156
L171:
	;
	goto L154
L172:
	;
	goto L150
L173:
	;
	if v870 != int32(-1) {
		v727 = v870
		goto L123
	} else {
		goto L174
	}
L174:
	;
	goto L124
L175:
	;
	F_UnlockPage(m, v997, int32(0), int32(7))
	mBase = m.M
	v1005 = m.ExcPending
	if v1005 != 0 {
		goto L4
	} else {
		goto L176
	}
L176:
	;
	v1006 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	v1007 = *(*int32)(unsafe.Add(mBase, uint32(v99)+28))
	v1015 = int32(1)
	goto L178
L177:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L4
	} else {
		goto L304
	}
L178:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1034 = m.ExcPending
	if v1034 != 0 {
		goto L4
	} else {
		goto L180
	}
L179:
	;
	F_LockPage(m, v997, int32(1), int32(7))
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L4
	} else {
		goto L227
	}
L180:
	;
	v1035 = int32(0)
	v1037 = F_ReadBufferExtended(m, v1007, v1035, v1015, v1035, v1006)
	mBase = m.M
	v1038 = m.ExcPending
	if v1038 != 0 {
		goto L4
	} else {
		goto L181
	}
L181:
	;
	F_LockBufferInternal(m, v1037, int32(1))
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L4
	} else {
		goto L182
	}
L182:
	;
	if v1037 < int32(0) {
		goto L185
	} else {
		goto L186
	}
L183:
	;
	v1349 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1059)+16)))
	v1351 = *(*int32)(unsafe.Add(mBase, uint32(v1059+v1349)))
	F_UnlockReleaseBuffer(m, v1037)
	mBase = m.M
	v1353 = m.ExcPending
	if v1353 != 0 {
		goto L4
	} else {
		goto L225
	}
L184:
	;
	v1060 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1059)+12)))
	if base.Ui32(v1060) < base.Ui32(int32(25)) {
		goto L183
	} else {
		goto L188
	}
L185:
	;
	v1045 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[1]))
	v1051 = *(*int32)(unsafe.Add(mBase, uint32(v1045+(v1037^int32(-1))<<(uint(int32(2))%32))))
	v1059 = v1051
	goto L184
L186:
	;
	goto L187
L187:
	;
	v1053 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[2]))
	v1059 = v1053 + v1037<<(uint(int32(13))%32) + int32(-8192)
	goto L184
L188:
	;
	v1068 = int32(base.Ui32(v1060+int32(_a_F_hnswbulkdelete_3))>>(uint(int32(2))%32)) & int32(_a_F_hnswbulkdelete_6)
	if v1068 == int32(0) {
		goto L183
	} else {
		goto L189
	}
L189:
	;
	v1076 = int32(1)
	goto L190
L190:
	;
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v1059+int32(20)+v1076&int32(_a_F_hnswbulkdelete_6)<<(uint(int32(2))%32))))
	v1105 = v1059 + v1102&int32(_a_F_hnswbulkdelete_5)
	v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1105))))
	if v1106 != int32(1) {
		goto L192
	} else {
		goto L193
	}
L191:
	;
	goto L183
L192:
	;
	v1322 = v1076 + int32(1)
	if base.Ui32(v1322&int32(_a_F_hnswbulkdelete_6)) <= base.Ui32(v1068) {
		v1076 = v1322
		goto L190
	} else {
		goto L224
	}
L193:
	;
	v1109 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1105)+2)))
	if v1109 != 0 {
		goto L192
	} else {
		goto L194
	}
L194:
	;
	v1110 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105)+8)))
	if v1110 == int32(0) {
		goto L192
	} else {
		goto L195
	}
L195:
	;
	v1113 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105)+68)))
	v1116 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105)+66)))
	v1117 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1105)+64)))
	v1120 = v1116 | v1117<<(uint(int32(16))%32)
	if v1120 == v1015 {
		v1146 = v1037
		v1147 = v1059
		goto L196
	} else {
		goto L197
	}
L196:
	;
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v1113<<(uint(int32(2))%32)+v1147)+20))
	v1152 = v1149&int32(_a_F_hnswbulkdelete_5) + v1147
	v1153 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1152)+2)))
	if v1153 != 0 {
		goto L203
	} else {
		goto L204
	}
L197:
	;
	v1122 = int32(0)
	v1124 = F_ReadBufferExtended(m, v1007, v1122, v1120, v1122, v1006)
	mBase = m.M
	v1125 = m.ExcPending
	if v1125 != 0 {
		goto L4
	} else {
		goto L198
	}
L198:
	;
	F_LockBufferInternal(m, v1124, int32(1))
	mBase = m.M
	v1128 = m.ExcPending
	if v1128 != 0 {
		goto L4
	} else {
		goto L199
	}
L199:
	;
	if v1124 < int32(0) {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	v1132 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[1]))
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v1132+(v1124^int32(-1))<<(uint(int32(2))%32))))
	v1146 = v1124
	v1147 = v1138
	goto L196
L201:
	;
	goto L202
L202:
	;
	v1140 = *(*int32)(unsafe.Add(mBase, _c_F_hnswbulkdelete[2]))
	v1146 = v1124
	v1147 = v1140 + v1124<<(uint(int32(13))%32) + int32(-8192)
	goto L196
L203:
	;
	v1157 = int32(0)
	v1160 = v1153
	goto L206
L204:
	;
	goto L205
L205:
	;
	if v1037 == v1146 {
		goto L192
	} else {
		goto L222
	}
L206:
	;
	v1182 = v1152 + int32(4) + v1157*int32(6)
	v1183 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1182)+4)))
	if v1183 != 0 {
		goto L208
	} else {
		goto L209
	}
L207:
	;
	goto L205
L208:
	;
	v1184 = *(*int32)(unsafe.Add(mBase, uint32(v1182)))
	*(*int32)(unsafe.Add(mBase, uint32(v99)+4)) = v1184
	v1186 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1182)+4)))
	*(*uint16)(unsafe.Add(mBase, uint32(v99)+8)) = uint16(v1186)
	v1188 = *(*int32)(unsafe.Add(mBase, uint32(v99)+64))
	v1190 = v99 + int32(4)
	v1195 = m.G0
	v1197 = v1195 - int32(16)
	m.G0 = v1197
	v1200 = v99 + int32(8)
	v1201 = int64(*(*uint16)(unsafe.Add(mBase, uint32(v1200))))
	v1202 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v1190))))
	v1203 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1200))))
	*(*uint16)(unsafe.Add(mBase, uint32(v1197)+12)) = uint16(v1203)
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v1190)))
	*(*int32)(unsafe.Add(mBase, uint32(v1197)+8)) = v1205
	v1207 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+20))
	v1208 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+12))
	v1210 = v1201 << (uint(int64(32)) % 64)
	v1211 = int64(33)
	v1216 = (int64(base.Ui64(v1210)>>(uint(v1211)%64)) ^ (v1210 | v1202)) * int64(-49064778989728563)
	v1221 = (int64(base.Ui64(v1216)>>(uint(v1211)%64)) ^ v1216) * int64(-4265267296055464877)
	v1226 = v1208 & base.I32_wrap_i64(int64(base.Ui64(v1221)>>(uint(v1211)%64))^v1221)
	v1229 = v1207 + v1226<<(uint(int32(3))%32)
	v1230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1229)+6)))
	if v1230 != 0 {
		goto L213
	} else {
		goto L214
	}
L209:
	;
	v1266 = v1160
	goto L210
L210:
	;
	v1268 = v1157 + int32(1)
	if base.Ui32(v1268) < base.Ui32(v1266&int32(_a_F_hnswbulkdelete_6)) {
		v1157 = v1268
		v1160 = v1266
		goto L206
	} else {
		goto L221
	}
L211:
	;
	if v1257 != 0 {
		goto L177
	} else {
		goto L220
	}
L212:
	;
	m.G0 = v1197 + int32(16)
	goto L211
L213:
	;
	v1232 = v1229
	v1236 = v1226
	goto L216
L214:
	;
	goto L215
L215:
	;
	v1257 = int32(0)
	goto L212
L216:
	;
	v1239 = F_ItemPointerEquals(m, v1232, v1197+int32(8))
	mBase = m.M
	if v1239 != 0 {
		v1257 = v1232
		goto L212
	} else {
		goto L218
	}
L217:
	;
	goto L215
L218:
	;
	v1240 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+20))
	v1241 = *(*int32)(unsafe.Add(mBase, uint32(v1188)+12))
	v1244 = v1241 & (v1236 + int32(1))
	v1247 = v1240 + v1244<<(uint(int32(3))%32)
	v1248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1247)+6)))
	if v1248 != 0 {
		v1232 = v1247
		v1236 = v1244
		goto L216
	} else {
		goto L219
	}
L219:
	;
	goto L217
L220:
	;
	v1265 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1152)+2)))
	v1266 = v1265
	goto L210
L221:
	;
	goto L207
L222:
	;
	F_UnlockReleaseBuffer(m, v1146)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L4
	} else {
		goto L223
	}
L223:
	;
	goto L192
L224:
	;
	goto L191
L225:
	;
	if v1351 != int32(-1) {
		v1015 = v1351
		goto L178
	} else {
		goto L226
	}
L226:
	;
	goto L179
L227:
	;
	F_UnlockPage(m, v997, int32(1), int32(7))
	mBase = m.M
	v1363 = m.ExcPending
	if v1363 != 0 {
		goto L4
	} else {
		goto L228
	}
L228:
	;
	v1372 = int32(1)
	v1375 = int32(-1)
	goto L229
L229:
	;
	F_vacuum_delay_point(m, int32(0))
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L4
	} else {
		goto L231
	}
L230:
	;
	v1729 = int32(0)
	F_HnswUpdateMetaPage(m, v997, v1729, v1729, v1706, v1729, v1729)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L4
	} else {
		goto L299
	}
L231:
	;
	v1392 = int32(0)
	v1394 = F_ReadBufferExtended(m, v997, v1392, v1372, v1392, v996)
	mBase = m.M
	v1395 = m.ExcPending
	if v1395 != 0 {
		goto L4
	} else {
		goto L232
	}
L232:
	;
	F_LockBufferForCleanup(m, v1394)
	mBase = m.M
	v1397 = m.ExcPending
	if v1397 != 0 {
		goto L4
	} else {
		goto L233
	}
L233:
	;
	v1398 = F_GenericXLogStart(m, v997)
	mBase = m.M
	v1399 = m.ExcPending
	if v1399 != 0 {
		goto L4
	} else {
		goto L235
	}
L234:
	;
	v1720 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1697)+16)))
	v1722 = *(*int32)(unsafe.Add(mBase, uint32(v1697+v1720)))
	F_pfree(m, v1708)
	mBase = m.M
	v1724 = m.ExcPending
	if v1724 != 0 {
		goto L4
	} else {
		goto L296
	}
L235:
	;
	v1401 = F_GenericXLogRegisterBuffer(m, v1398, v1394, int32(0))
	mBase = m.M
	v1402 = m.ExcPending
	if v1402 != 0 {
		goto L4
	} else {
		goto L236
	}
L236:
	;
	v1403 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1401)+12)))
	if base.Ui32(v1403) < base.Ui32(int32(25)) {
		v1697 = v1401
		v1706 = v1375
		v1708 = v1398
		goto L234
	} else {
		goto L237
	}
L237:
	;
	v1407 = v1403 + int32(_a_F_hnswbulkdelete_3)
	if v1407&int32(_a_F_hnswbulkdelete_4) == int32(0) {
		v1697 = v1401
		v1706 = v1375
		v1708 = v1398
		goto L234
	} else {
		goto L238
	}
L238:
	;
	v1417 = v1401
	v1420 = int32(1)
	v1426 = v1375
	v1428 = v1398
	goto L239
L239:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1417+v1420<<(uint(int32(2))%32))+20))
	v1446 = v1417 + v1443&int32(_a_F_hnswbulkdelete_5)
	v1447 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446))))
	if v1447 != int32(1) {
		v1671 = v1417
		v1680 = v1426
		v1682 = v1428
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1697 = v1671
	v1706 = v1680
	v1708 = v1682
	goto L234
L241:
	;
	if v1420 != int32(base.Ui32(v1407)>>(uint(int32(2))%32))&int32(_a_F_hnswbulkdelete_6) {
		v1417 = v1671
		v1420 = v1420 + int32(1)
		v1426 = v1680
		v1428 = v1682
		goto L239
	} else {
		goto L295
	}
L242:
	;
	v1450 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446)+2)))
	if v1450 != 0 {
		goto L243
	} else {
		goto L244
	}
L243:
	;
	if v1426 == int32(-1) {
		goto L246
	} else {
		goto L247
	}
L244:
	;
	goto L245
L245:
	;
	v1454 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1446)+8)))
	if v1454 != 0 {
		v1671 = v1417
		v1680 = v1426
		v1682 = v1428
		goto L241
	} else {
		goto L249
	}
L246:
	;
	v1453 = v1372
	goto L248
L247:
	;
	v1453 = v1426
	goto L248
L248:
	;
	v1671 = v1417
	v1680 = v1453
	v1682 = v1428
	goto L241
L249:
	;
	v1455 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1446)+68)))
	v1458 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1446)+66)))
	v1459 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1446)+64)))
	v1462 = v1458 | v1459<<(uint(int32(16))%32)
	if v1372 != v1462 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v1464 = int32(0)
	v1466 = F_ReadBufferExtended(m, v997, v1464, v1462, v1464, v996)
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L4
	} else {
		goto L253
	}
L251:
	;
	v1474 = v1417
	v1475 = v1394
	goto L252
L252:
	;
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1455<<(uint(int32(2))%32)+v1474)+20))
	v1478 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v1446)+2)) = uint8(v1478)
	v1481 = v1446 + int32(72)
	v1482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446)+72)))
	if v1482 == v1478 {
		goto L257
	} else {
		goto L258
	}
L253:
	;
	F_LockBufferInternal(m, v1466, int32(3))
	mBase = m.M
	v1470 = m.ExcPending
	if v1470 != 0 {
		goto L4
	} else {
		goto L254
	}
L254:
	;
	v1472 = F_GenericXLogRegisterBuffer(m, v1428, v1466, int32(0))
	mBase = m.M
	v1473 = m.ExcPending
	if v1473 != 0 {
		goto L4
	} else {
		goto L255
	}
L255:
	;
	v1474 = v1472
	v1475 = v1466
	goto L252
L256:
	;
	if v1507 != 0 {
		goto L267
	} else {
		goto L268
	}
L257:
	;
	v1486 = int32(18)
	v1488 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446)+73)))
	if v1488 == v1486 {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L259
L259:
	;
	v1499 = int32(1)
	if v1482&v1499 != 0 {
		v1507 = int32(base.Ui32(v1482) >> (uint(v1499) % 32))
		goto L256
	} else {
		goto L266
	}
L260:
	;
	v1491 = v1486
	goto L262
L261:
	;
	v1491 = int32(2)
	goto L262
L262:
	;
	if base.Ui32((v1488-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L263
	} else {
		goto L264
	}
L263:
	;
	v1498 = int32(6)
	goto L265
L264:
	;
	v1498 = v1491
	goto L265
L265:
	;
	v1507 = v1498
	goto L256
L266:
	;
	v1503 = *(*int32)(unsafe.Add(mBase, uint32(v1481)))
	v1507 = int32(base.Ui32(v1503) >> (uint(int32(2)) % 32))
	goto L256
L267:
	;
	base.MemoryFill(m, v1481, int32(0), v1507)
	goto L269
L268:
	;
	goto L269
L269:
	;
	v1512 = v1477&int32(_a_F_hnswbulkdelete_5) + v1474
	v1513 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1512)+2)))
	if v1513 == int32(0) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1647 = int32(1)
	v1648 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1446)+3)))
	v1650 = v1648 + v1647
	if base.Ui32(int32(15)) < base.Ui32(v1650&int32(255)) {
		goto L282
	} else {
		goto L283
	}
L271:
	;
	v1517 = v1513 & int32(3)
	v1518 = int32(4)
	v1519 = v1512 + v1518
	v1520 = int32(0)
	if base.Ui32(v1518) <= base.Ui32(v1513) {
		goto L272
	} else {
		goto L273
	}
L272:
	;
	v1527 = v1520
	v1535 = int32(0)
	goto L275
L273:
	;
	v1566 = v1520
	goto L274
L274:
	;
	v1589 = v1566
	v1592 = v1520
	goto L279
L275:
	;
	v1552 = v1519 + v1527*int32(6)
	*(*int64)(unsafe.Add(mBase, uint32(v1552)+16)) = int64(281474976645120)
	*(*int64)(unsafe.Add(mBase, uint32(v1552)+8)) = int64(-4294901761)
	*(*int64)(unsafe.Add(mBase, uint32(v1552))) = int64(-281470681743361)
	v1559 = int32(4)
	v1560 = v1527 + v1559
	v1562 = v1535 + v1559
	if v1562 != v1513&int32(_a_F_hnswbulkdelete_8) {
		v1527 = v1560
		v1535 = v1562
		goto L275
	} else {
		goto L277
	}
L276:
	;
	if v1517 == int32(0) {
		goto L270
	} else {
		goto L278
	}
L277:
	;
	goto L276
L278:
	;
	v1566 = v1560
	goto L274
L279:
	;
	v1614 = v1519 + v1589*int32(6)
	v1615 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1614)+4)) = uint16(v1615)
	*(*int32)(unsafe.Add(mBase, uint32(v1614))) = int32(-1)
	v1619 = int32(1)
	v1622 = v1592 + v1619
	if v1622 != v1517 {
		v1589 = v1589 + v1619
		v1592 = v1622
		goto L279
	} else {
		goto L281
	}
L280:
	;
	goto L270
L281:
	;
	goto L280
L282:
	;
	v1655 = v1647
	goto L284
L283:
	;
	v1655 = v1650
	goto L284
L284:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v1446)+3)) = uint8(v1655)
	*(*uint8)(unsafe.Add(mBase, uint32(v1512)+1)) = uint8(v1655)
	F_GenericXLogFinish(m, v1428)
	mBase = m.M
	v1659 = m.ExcPending
	if v1659 != 0 {
		goto L4
	} else {
		goto L285
	}
L285:
	;
	if v1394 != v1475 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	F_UnlockReleaseBuffer(m, v1475)
	mBase = m.M
	v1662 = m.ExcPending
	if v1662 != 0 {
		goto L4
	} else {
		goto L289
	}
L287:
	;
	goto L288
L288:
	;
	if v1426 == int32(-1) {
		goto L290
	} else {
		goto L291
	}
L289:
	;
	goto L288
L290:
	;
	v1665 = v1372
	goto L292
L291:
	;
	v1665 = v1426
	goto L292
L292:
	;
	v1666 = F_GenericXLogStart(m, v997)
	mBase = m.M
	v1667 = m.ExcPending
	if v1667 != 0 {
		goto L4
	} else {
		goto L293
	}
L293:
	;
	v1669 = F_GenericXLogRegisterBuffer(m, v1666, v1394, int32(0))
	mBase = m.M
	v1670 = m.ExcPending
	if v1670 != 0 {
		goto L4
	} else {
		goto L294
	}
L294:
	;
	v1671 = v1669
	v1680 = v1665
	v1682 = v1666
	goto L241
L295:
	;
	goto L240
L296:
	;
	F_UnlockReleaseBuffer(m, v1394)
	mBase = m.M
	v1726 = m.ExcPending
	if v1726 != 0 {
		goto L4
	} else {
		goto L297
	}
L297:
	;
	if v1722 != int32(-1) {
		v1372 = v1722
		v1375 = v1706
		goto L229
	} else {
		goto L298
	}
L298:
	;
	goto L230
L299:
	;
	v1735 = *(*int32)(unsafe.Add(mBase, uint32(v99)+64))
	F_tidhash_destroy(m, v1735)
	mBase = m.M
	v1737 = m.ExcPending
	if v1737 != 0 {
		goto L4
	} else {
		goto L300
	}
L300:
	;
	v1738 = *(*int32)(unsafe.Add(mBase, uint32(v99)+68))
	F_bms_free(m, v1738)
	mBase = m.M
	v1740 = m.ExcPending
	if v1740 != 0 {
		goto L4
	} else {
		goto L301
	}
L301:
	;
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v99)+72))
	F_pfree(m, v1741)
	mBase = m.M
	v1743 = m.ExcPending
	if v1743 != 0 {
		goto L4
	} else {
		goto L302
	}
L302:
	;
	v1744 = *(*int32)(unsafe.Add(mBase, uint32(v99)+292))
	F_MemoryContextDelete(m, v1744)
	mBase = m.M
	v1746 = m.ExcPending
	if v1746 != 0 {
		goto L4
	} else {
		goto L303
	}
L303:
	;
	v1747 = *(*int32)(unsafe.Add(mBase, uint32(v99)+32))
	m.G0 = v99 + int32(304)
	return v1747
L304:
	;
	F_errmsg_internal(m, int32(_a_F_hnswbulkdelete_9), int32(0))
	mBase = m.M
	v1759 = m.ExcPending
	if v1759 != 0 {
		goto L4
	} else {
		goto L305
	}
L305:
	;
	F_errfinish(m, int32(_a_F_hnswbulkdelete_10), int32(578), int32(_a_F_hnswbulkdelete_11))
	mBase = m.M
	v1764 = m.ExcPending
	if v1764 != 0 {
		goto L4
	} else {
		goto L306
	}
L306:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hnswcostestimate(m *base.Module, l0 int32, l1 int32, l2 float64, l3 int32, l4 int32, l5 int32, l6 int32, l7 int32) {
	mBase := m.M
	_ = mBase
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v43 float64
	_ = v43
	var v46 float64
	_ = v46
	var v48 int32
	_ = v48
	var v50 float64
	_ = v50
	var v51 float64
	_ = v51
	var v52 int32
	_ = v52
	var v62 float64
	_ = v62
	var v71 float64
	_ = v71
	var v74 float64
	_ = v74
	var v76 float64
	_ = v76
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v92 float64
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	var v101 float64
	_ = v101
	var v103 float64
	_ = v103
	var v109 float64
	_ = v109
	var v113 float64
	_ = v113
	var v115 float64
	_ = v115
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	v16 = m.G0
	v18 = v16 - int32(96)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v20 != 0 {
		v22 = v18 + int32(24)
		base.MemoryFill(m, v22, int32(0), int32(72))
		F_genericcostestimate(m, l0, l1, l2, v22)
		mBase = m.M
		v27 = m.ExcPending
		if v27 != 0 {
			return
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
			v31 = F_index_open(m, v29, int32(0))
			mBase = m.M
			v32 = m.ExcPending
			if v32 != 0 {
				return
			} else {
				F_HnswGetMetaPageInfo(m, v31, v18+int32(20), int32(0))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return
				} else {
					F_relation_close(m, v31, int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v42 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
						v43 = *(*float64)(unsafe.Add(mBase, uint32(v42)+24))
						if base.F64_gt(v43, float64(0)) != 0 {
							v46 = F_log(m, v43)
							mBase = m.M
							v48 = *(*int32)(unsafe.Add(mBase, _c_F_hnswcostestimate[0]))
							v50 = F_log(m, base.F64_convert_i32_s(v48))
							mBase = m.M
							v51 = float64(1)
							v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
							v62 = F_log(m, base.F64_convert_i32_s(v52))
							mBase = m.M
							v71 = base.F64_div(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v48*v52<<(uint(int32(1))%32)), base.F64_div(base.F64_mul(v46, float64(0.55)), base.F64_mul(base.F64_add(v50, v51), v62))), base.F64_convert_i32_s(v52*base.I32_trunc_sat_f64_s(base.F64_div(v46, v62)))), v43)
							if base.F64_gt(v71, v51) != 0 {
								v74 = v51
							} else {
								v74 = v71
							}
							v76 = v74
						} else {
							v76 = float64(1)
						}
						v79 = *(*int32)(unsafe.Add(mBase, uint32(v42)+8))
						F_get_tablespace_page_costs(m, v79, int32(0), v18+int32(8))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return
						} else {
							v85 = *(*float64)(unsafe.Add(mBase, uint32(v18)+56))
							v86 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
							v87 = base.F64_mul(v86, v76)
							if base.F64_lt(v76, float64(0.5)) == int32(0) {
								v109 = v87
							} else {
								v92 = base.F64_mul(v85, v76)
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+124))
								v96 = base.F64_convert_i32_u(v95)
								if base.F64_gt(v92, v96) == int32(0) {
									v109 = v87
								} else {
									v101 = *(*float64)(unsafe.Add(mBase, uint32(v18)+8))
									v103 = *(*float64)(unsafe.Add(mBase, uint32(v18)+72))
									v109 = base.F64_add(base.F64_mul(base.F64_sub(v96, v92), v101), base.F64_add(v87, base.F64_mul(v92, base.F64_sub(v101, v103))))
								}
							}
							*(*float64)(unsafe.Add(mBase, uint32(l3))) = v109
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v86
							v113 = *(*float64)(unsafe.Add(mBase, uint32(v18)+40))
							*(*float64)(unsafe.Add(mBase, uint32(l5))) = v113
							v115 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
							*(*float64)(unsafe.Add(mBase, uint32(l6))) = v115
							*(*float64)(unsafe.Add(mBase, uint32(l7))) = v85
							m.G0 = v18 + int32(96)
							return
						}
					}
				}
			}
		}
	} else {
		v118 = int64(9218868437227405312)
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = v118
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = v118
		v122 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l5))) = v122
		*(*int64)(unsafe.Add(mBase, uint32(l6))) = v122
		*(*int64)(unsafe.Add(mBase, uint32(l7))) = v122
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(2)
		m.G0 = v18 + int32(96)
		return
	}
}
func F_hs_contains(m *base.Module, l0 int32) int64 {
	var v2 int64
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_contains(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int64(0)
	} else {
		return v2
	}
}
func F_hypothetical_check_argtypes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v51 int32
	_ = v51
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L12
	} else {
		goto L18
	}
L2:
	;
	v10 = l1 + int32(1)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v10 != v11 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2+v10<<(uint(int32(3))%32)+l1*int32(100))+96))
	if v19 != int32(23) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v22 = int32(0)
	if v22 < l1 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = l1
	goto L7
L6:
	;
	v26 = v22
	goto L7
L7:
	;
	v30 = v22
	goto L9
L8:
	;
	return
L9:
	;
	if v26 == v30 {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v51 = m.ExcPending
	if v51 != 0 {
		goto L12
	} else {
		goto L15
	}
L11:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v39 = v30 + int32(1)
	v40 = F_get_fn_expr_argtype(m, v37, v39)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30*int32(100)+(l2+v36<<(uint(int32(3))%32)))+96))
	if v40 == v46 {
		v30 = v39
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	F_errmsg_internal(m, int32(_a_F_hypothetical_check_argtypes_0), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(_a_F_hypothetical_check_argtypes_1), int32(1160), int32(_a_F_hypothetical_check_argtypes_2))
	mBase = m.M
	v60 = m.ExcPending
	if v60 != 0 {
		goto L12
	} else {
		goto L17
	}
L17:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L18:
	;
	F_errmsg_internal(m, int32(_a_F_hypothetical_check_argtypes_0), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(_a_F_hypothetical_check_argtypes_1), int32(1152), int32(_a_F_hypothetical_check_argtypes_2))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L12
	} else {
		goto L20
	}
L20:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_hypothetical_cume_dist_final(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int64
	_ = v12
	var v15 int32
	_ = v15
	var v16 int64
	_ = v16
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v12 = F_hypothetical_rank_common(m, l0, int32(1), v7+int32(8))
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = *(*int64)(unsafe.Add(mBase, uint32(v7)+8))
		m.G0 = v7 + int32(16)
		return base.I64_reinterpret_f64(base.F64_div(base.F64_convert_i64_s(v12), base.F64_convert_i64_s(v16+int64(1))))
	}
}
func F_hypothetical_rank_common(m *base.Module, l0 int32, l1 int32, l2 int32) int64 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int64
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v123 int64
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
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
	var v165 int32
	_ = v165
	var v167 int64
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v205 int64
	_ = v205
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v213 int64
	_ = v213
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v231 int64
	_ = v231
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v244 int32
	_ = v244
	var v249 int32
	_ = v249
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
	if v13 == int32(1) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = int64(0)
	return int64(1)
L2:
	;
	goto L3
L3:
	;
	v20 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+18)))
	v22 = l0 + int32(24)
	v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
	v24 = *(*int64)(unsafe.Add(mBase, uint32(v23)+16))
	*(*int64)(unsafe.Add(mBase, uint32(l2))) = v24
	v26 = int32(1)
	v27 = v20 - v26
	if v27&v26 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v33 = v27 >> (uint(int32(1)) % 32)
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	F_hypothetical_check_argtypes(m, l0, v33, v35)
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L6
L6:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L7
	} else {
		goto L42
	}
L7:
	;
	return int64(0)
L8:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v23)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(v40)+20))
	v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v42)+12))
	m.T0[v43].(func(*base.Module, int32))(m, v41)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L7
	} else {
		goto L9
	}
L9:
	;
	v46 = int32(0)
	if v33 <= v46 {
		v141 = v46
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v142 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v142+v141<<(uint(int32(3))%32)))) = base.I64_extend_i32_s(l1)
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v150 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v148+v141))) = uint8(v150)
	v152 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
	v154 = v152 & int32(_a_F_hypothetical_rank_common_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)) = uint16(v154)
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)))
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+6)) = uint16(v157)
	goto L19
L11:
	;
	v49 = int32(0)
	if v27 != int32(2) {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v59 = v49
	v63 = int32(0)
	goto L15
L13:
	;
	v106 = v49
	goto L14
L14:
	;
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v122 = v22 + v106<<(uint(int32(4))%32)
	v123 = *(*int64)(unsafe.Add(mBase, uint32(v122)+16))
	*(*int64)(unsafe.Add(mBase, uint32(v116+v106<<(uint(int32(3))%32)))) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+24)))
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v106))) = uint8(v127)
	v141 = v33
	goto L10
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v70 = int32(3)
	v74 = v59 | int32(1)
	v75 = int32(4)
	v77 = v22 + v74<<(uint(v75)%32)
	v78 = *(*int64)(unsafe.Add(mBase, uint32(v77)))
	*(*int64)(unsafe.Add(mBase, uint32(v69+v59<<(uint(v70)%32)))) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v59))) = uint8(v82)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v88 = int32(2)
	v89 = v59 + v88
	v92 = v22 + v89<<(uint(v75)%32)
	v93 = *(*int64)(unsafe.Add(mBase, uint32(v92)))
	*(*int64)(unsafe.Add(mBase, uint32(v84+v74<<(uint(v70)%32)))) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+8)))
	*(*uint8)(unsafe.Add(mBase, uint32(v74+v95))) = uint8(v97)
	v100 = v63 + v88
	if v100 != v33&int32(2147483646) {
		v59 = v89
		v63 = v100
		goto L15
	} else {
		goto L17
	}
L16:
	;
	if v33&int32(1) == int32(0) {
		v141 = v33
		goto L10
	} else {
		goto L18
	}
L17:
	;
	goto L16
L18:
	;
	v106 = v89
	goto L14
L19:
	;
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_tuplesort_puttupleslot(m, v159, v41)
	mBase = m.M
	v161 = m.ExcPending
	if v161 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v162 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_tuplesort_performsort(m, v162)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v165 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)) = uint8(v165)
	v167 = int64(1)
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v172 = F_tuplesort_gettupleslot(m, v168, v165, v165, v41, int32(0))
	mBase = m.M
	v173 = m.ExcPending
	if v173 != 0 {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v233 = *(*int32)(unsafe.Add(mBase, uint32(v232)+12))
	m.T0[v233].(func(*base.Module, int32))(m, v41)
	mBase = m.M
	v235 = m.ExcPending
	if v235 != 0 {
		goto L7
	} else {
		goto L41
	}
L23:
	;
	if v172 == int32(0) {
		v231 = v167
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v191 = v167
	goto L25
L25:
	;
	v192 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+6)))
	if v192 <= v33 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v231 = v213
	goto L22
L27:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v194)+16))
	m.T0[v195].(func(*base.Module, int32, int32))(m, v41, v33+int32(1))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v198 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v198+v33))))
	if v200 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v205 = int64(*(*uint32)(unsafe.Add(mBase, uint32(v203+v33<<(uint(int32(3))%32)))))
	if v205 != int64(0) {
		v231 = v191
		goto L22
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v209 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_rank_common[0]))
	if v209 != 0 {
		goto L35
	} else {
		goto L36
	}
L34:
	;
	goto L33
L35:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v211 = m.ExcPending
	if v211 != 0 {
		goto L7
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v213 = v191 + int64(1)
	v214 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v215 = int32(1)
	v218 = F_tuplesort_gettupleslot(m, v214, v215, v215, v41, int32(0))
	mBase = m.M
	v219 = m.ExcPending
	if v219 != 0 {
		goto L7
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	if v218 != 0 {
		v191 = v213
		goto L25
	} else {
		goto L40
	}
L40:
	;
	goto L26
L41:
	;
	return v231
L42:
	;
	F_errmsg_internal(m, int32(_a_F_hypothetical_rank_common_1), int32(0))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_hypothetical_rank_common_2), int32(1195), int32(_a_F_hypothetical_rank_common_3))
	mBase = m.M
	v249 = m.ExcPending
	if v249 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
