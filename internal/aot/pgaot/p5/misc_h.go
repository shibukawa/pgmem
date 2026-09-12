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
	v5 = *(*int32)(unsafe.Add(mBase, _consts[1234]))
	v10 = F_AllocSetContextCreateInternal(m, v5, int32(62496), int32(0), int32(8192), int32(8388608))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return
	} else {
		*(*int32)(unsafe.Add(mBase, uint32(l0)+108)) = v10
		v13 = int32(4515488)
		v14 = *(*int32)(unsafe.Add(mBase, _consts[9]))
		*(*int32)(unsafe.Add(mBase, _consts[9])) = v10
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
		v20 = int32(1)
		v24 = *(*int32)(unsafe.Add(mBase, _consts[20]))
		v25 = F_tuplestore_begin_heap(m, int32(base.Ui32(v17&int32(2))>>(uint(v20)%32)), v20, v24)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(l0)+104)) = v25
			*(*int32)(unsafe.Add(mBase, _consts[9])) = v14
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
func F_has_largeobject_privilege_id(m *base.Module, l0 int32) int32 {
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
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	v7 = *(*int32)(unsafe.Add(mBase, _consts[239]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = F_convert_any_priv_string(m, v10, int32(1657504))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
		} else {
			if v15&int64(4) == int64(0) {
				v22 = *(*int32)(unsafe.Add(mBase, _consts[293]))
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v22)))
				v24 = v23
			} else {
				v24 = int32(0)
			}
			v25 = F_LargeObjectExistsWithSnapshot(m, v8, v24)
			mBase = m.M
			v26 = m.ExcPending
			if v26 != 0 {
				return int32(0)
			} else {
				if v25 != 0 {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1033])))
					if v29 != 0 {
						v38 = int32(1)
						return v38
					} else {
						v30 = F_pg_largeobject_aclcheck_snapshot(m, v8, v7, v15, v24)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v30 == int32(0))
						}
					}
				} else {
					v35 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v35)
					v38 = int32(0)
					return v38
				}
			}
		}
	}
}
func F_hashbuildCallback(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
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
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
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
	var v40 int32
	_ = v40
	var v42 int32
	_ = v42
	var v44 float64
	_ = v44
	v7 = m.G0
	v9 = v7 - int32(16)
	m.G0 = v9
	v15 = F__hash_convert_tuple(m, l0, l2, l3, v9+int32(12), v9+int32(11))
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return
	} else {
		if v15 != 0 {
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
			if v17 != 0 {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				F_tuplesort_putindextuplevalues(m, v18, v19, l1, v9+int32(12), v9+int32(11))
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return
				} else {
					v44 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
					*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v44, float64(1))
					m.G0 = v9 + int32(16)
					return
				}
			} else {
				v26 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				v31 = F_index_form_tuple(m, v26, v9+int32(12), v9+int32(11))
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return
				} else {
					v33 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
					*(*uint16)(unsafe.Add(mBase, uint32(v31)+4)) = uint16(v33)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
					*(*int32)(unsafe.Add(mBase, uint32(v31))) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l5)+16))
					F__hash_doinsert(m, l0, v31, v37, int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						F_pfree(m, v31)
						mBase = m.M
						v42 = m.ExcPending
						if v42 != 0 {
							return
						} else {
							v44 = *(*float64)(unsafe.Add(mBase, uint32(l5)+8))
							*(*float64)(unsafe.Add(mBase, uint32(l5)+8)) = base.F64_add(v44, float64(1))
							m.G0 = v9 + int32(16)
							return
						}
					}
				}
			}
		} else {
			m.G0 = v9 + int32(16)
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
	var v18 int64
	_ = v18
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v37 float64
	_ = v37
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v105 int32
	_ = v105
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
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v179 int32
	_ = v179
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v239 int32
	_ = v239
	var v240 float64
	_ = v240
	var v241 int32
	_ = v241
	var v244 float64
	_ = v244
	var v246 float64
	_ = v246
	var v250 float64
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v270 int32
	_ = v270
	var v273 int32
	_ = v273
	var v275 int32
	_ = v275
	var v278 int32
	_ = v278
	var v281 int64
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v287 int32
	_ = v287
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 float64
	_ = v324
	var v326 float64
	_ = v326
	var v327 float64
	_ = v327
	v5 = int32(0)
	v18 = int64(0)
	v19 = m.G0
	v21 = v19 - int32(32)
	m.G0 = v21
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v21)+24)) = v18
	*(*int64)(unsafe.Add(mBase, uint32(v21)+16)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v5
	v33 = F__hash_getcachedmetap(m, v23, v21+int32(12), v5)
	mBase = m.M
	v36 = m.ExcPending
	if v36 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v37 = *(*float64)(unsafe.Add(mBase, uint32(v33)+8))
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v33)+24))
	v45 = v33
	v47 = v5
	v48 = v38
	goto L3
L3:
	;
	if base.Ui32(v47) <= base.Ui32(v48) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v232 = int32(4510148)
	v234 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v234 + int32(1)
	v239 = v220 + int32(32)
	v240 = *(*float64)(unsafe.Add(mBase, uint32(v239)))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v220)+48))
	if v38 != v241 {
		goto L50
	} else {
		goto L51
	}
L5:
	;
	v64 = v45
	v66 = v47
	goto L8
L6:
	;
	v179 = v47
	goto L7
L7:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v189 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L8:
	;
	v76 = int32(0)
	if v66 == v76 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v179 = v110
	goto L7
L10:
	;
	v112 = v111 + v110
	v114 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v115 = F_ReadBufferExtended(m, v23, v76, v112, int32(0), v114)
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L21
	}
L11:
	;
	v110 = int32(1)
	v111 = int32(0)
	goto L10
L12:
	;
	goto L13
L13:
	;
	v81 = int32(1)
	v82 = v66 + v81
	v86 = v82 - v81
	if base.Ui32(int32(2)) <= base.Ui32(v82) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v105<<(uint(int32(2))%32)+v64)+48))
	v110 = v82
	v111 = v109
	goto L10
L15:
	;
	v92 = int32(32) - base.I32_clz(v86)
	goto L17
L16:
	;
	v92 = int32(0)
	goto L17
L17:
	;
	if base.Ui32(int32(10)) <= base.Ui32(v92) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v95 = int32(3)
	v105 = int32(base.Ui32(v86)>>(uint(v92-v95)%32))&v95 | v92<<(uint(int32(2))%32) - int32(30)
	goto L20
L19:
	;
	v105 = v92
	goto L20
L20:
	;
	goto L14
L21:
	;
	F_LockBufferForCleanup(m, v115)
	mBase = m.M
	v118 = m.ExcPending
	if v118 != 0 {
		goto L1
	} else {
		goto L22
	}
L22:
	;
	F__hash_checkpage(m, v23, v115, int32(2))
	mBase = m.M
	v121 = m.ExcPending
	if v121 != 0 {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	if v115 < int32(0) {
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v155)+24))
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v155)+28))
	v159 = *(*int32)(unsafe.Add(mBase, uint32(v155)+32))
	F_hashbucketcleanup(m, v23, v66, v115, v112, v156, v157, v158, v159, v21+int32(24), v21+int32(16), base.B2i32(v144 == int32(64)), l2, l3)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L1
	} else {
		goto L32
	}
L25:
	;
	v140 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+16)))
	v141 = v140 + v139
	v142 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v141)+12)))
	v144 = v142 & int32(96)
	if v144 != int32(64) {
		v155 = v64
		goto L24
	} else {
		goto L29
	}
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125+(v115^int32(-1))<<(uint(int32(2))%32))))
	v139 = v131
	goto L25
L27:
	;
	goto L28
L28:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v139 = v133 + v115<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v64)+24))
	if base.Ui32(v147) <= base.Ui32(v148) {
		v155 = v64
		goto L24
	} else {
		goto L30
	}
L30:
	;
	v153 = F__hash_getcachedmetap(m, v23, v21+int32(12), int32(1))
	mBase = m.M
	v154 = m.ExcPending
	if v154 != 0 {
		goto L1
	} else {
		goto L31
	}
L31:
	;
	v155 = v153
	goto L24
L32:
	;
	F_ReleaseBuffer(m, v115)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	if base.Ui32(v110) <= base.Ui32(v48) {
		v64 = v155
		v66 = v110
		goto L8
	} else {
		goto L34
	}
L34:
	;
	goto L9
L35:
	;
	v195 = F__hash_getbuf(m, v23, int32(0), int32(-1), int32(8))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L38
	}
L36:
	;
	v198 = v189
	goto L37
L37:
	;
	F_LockBuffer(m, v198, int32(2))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L39
	}
L38:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v21)+12)) = v195
	v198 = v195
	goto L37
L39:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v202 < int32(0) {
		goto L41
	} else {
		goto L42
	}
L40:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(v220)+48))
	if v221 != v48 {
		goto L44
	} else {
		goto L45
	}
L41:
	;
	v206 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206+(v202^int32(-1))<<(uint(int32(2))%32))))
	v220 = v212
	goto L40
L42:
	;
	goto L43
L43:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v220 = v214 + v202<<(uint(int32(13))%32) + int32(-8192)
	goto L40
L44:
	;
	F_LockBuffer(m, v202, int32(0))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L47
	}
L45:
	;
	goto L46
L46:
	;
	goto L4
L47:
	;
	v229 = F__hash_getcachedmetap(m, v23, v21+int32(12), int32(1))
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(v229)+24))
	v45 = v229
	v47 = v179
	v48 = v231
	goto L3
L49:
	;
	F_MarkBufferDirty(m, v202)
	mBase = m.M
	v256 = m.ExcPending
	if v256 != 0 {
		goto L1
	} else {
		goto L56
	}
L50:
	;
	v246 = *(*float64)(unsafe.Add(mBase, uint32(v21)+24))
	if base.F64_gt(v240, v246) != 0 {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	if base.F64_ne(v240, v37) != 0 {
		goto L50
	} else {
		goto L52
	}
L52:
	;
	v244 = *(*float64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v220)+32)) = v244
	goto L49
L53:
	;
	v250 = base.F64_sub(v240, v246)
	goto L55
L54:
	;
	v250 = float64(0)
	goto L55
L55:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v239))) = v250
	*(*float64)(unsafe.Add(mBase, uint32(v21)+16)) = v250
	goto L49
L56:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v257)+118)))
	if v258 != int32(112) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v307 = int32(4510148)
	v309 = *(*int32)(unsafe.Add(mBase, _consts[13]))
	*(*int32)(unsafe.Add(mBase, _consts[13])) = v309 - int32(1)
	v313 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_UnlockReleaseBuffer(m, v313)
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L1
	} else {
		goto L72
	}
L58:
	;
	v262 = *(*int32)(unsafe.Add(mBase, _consts[14]))
	if v262 <= int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	if v265 != 0 {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v267 = *(*float64)(unsafe.Add(mBase, uint32(v220)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v21))) = v267
	F_XLogBeginInsert(m)
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	if v266 != 0 {
		goto L57
	} else {
		goto L63
	}
L63:
	;
	goto L61
L64:
	;
	F_XLogRegisterData(m, v21, int32(8))
	mBase = m.M
	v273 = m.ExcPending
	if v273 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_XLogRegisterBuffer(m, int32(0), v275, int32(8))
	mBase = m.M
	v278 = m.ExcPending
	if v278 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v281 = F_XLogInsert(m, int32(12), int32(176))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v283 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v301))) = base.I64_rotr(v281, int64(32))
	goto L57
L69:
	;
	v287 = *(*int32)(unsafe.Add(mBase, _consts[11]))
	v293 = *(*int32)(unsafe.Add(mBase, uint32(v287+(v283^int32(-1))<<(uint(int32(2))%32))))
	v301 = v293
	goto L68
L70:
	;
	goto L71
L71:
	;
	v295 = *(*int32)(unsafe.Add(mBase, _consts[12]))
	v301 = v295 + v283<<(uint(int32(13))%32) + int32(-8192)
	goto L68
L72:
	;
	if l1 == int32(0) {
		goto L73
	} else {
		goto L74
	}
L73:
	;
	v319 = F_palloc0(m, int32(40))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	v321 = l1
	goto L75
L75:
	;
	v322 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v321)+4)) = uint8(v322)
	v324 = *(*float64)(unsafe.Add(mBase, uint32(v21)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v321)+8)) = v324
	v326 = *(*float64)(unsafe.Add(mBase, uint32(v21)+24))
	v327 = *(*float64)(unsafe.Add(mBase, uint32(v321)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v321)+16)) = base.F64_add(v326, v327)
	m.G0 = v21 + int32(32)
	return v321
L76:
	;
	v321 = v319
	goto L75
}
func F_hashcharextended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
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
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		v12 = int32(-1636608428)
		v50 = v12
		v52 = v12
		v55 = v12
	} else {
		v15 = base.I32_wrap_i64(v4)
		v20 = base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v26 = v15 - v20 - int32(1636608428) ^ base.I32_rotl(v20, int32(6))
		v28 = v15 + int32(1021750440)
		v29 = v20 + v28
		v30 = v26 + v29
		v34 = v28 - v26 ^ base.I32_rotl(v26, int32(8))
		v38 = v29 - v34 ^ base.I32_rotl(v34, int32(16))
		v42 = v30 - v38 ^ base.I32_rotl(v38, int32(19))
		v43 = v34 + v30
		v44 = v38 + v43
		v50 = v42 + v44
		v52 = v44
		v55 = v43 - v42 ^ base.I32_rotl(v42, int32(4))
	}
	v57 = int32(14)
	v59 = v50 ^ v55 - base.I32_rotl(v50, v57)
	v64 = v59 ^ (v2 + v52) - base.I32_rotl(v59, int32(11))
	v68 = v64 ^ v50 - base.I32_rotl(v64, int32(25))
	v72 = v68 ^ v59 - base.I32_rotl(v68, int32(16))
	v76 = v72 ^ v64 - base.I32_rotl(v72, int32(4))
	v80 = v76 ^ v68 - base.I32_rotl(v76, v57)
	v90 = F_Int64GetDatum(m, base.I64_extend_i32_u(v80)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v80^v72-base.I32_rotl(v80, int32(24))))
	mBase = m.M
	v93 = m.ExcPending
	if v93 != 0 {
		return int32(0)
	} else {
		return v90
	}
}
func F_hashgetbitmap(m *base.Module, l0 int32, l1 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v16 int64
	_ = v16
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
	var v34 int64
	_ = v34
	v3 = int64(0)
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
	v16 = v3
	goto L6
L4:
	;
	v34 = v3
	goto L5
L5:
	;
	return v34
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
	v34 = v28
	goto L5
L8:
	;
	v28 = v16 + int64(1)
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
		v16 = v28
		goto L6
	} else {
		goto L10
	}
L10:
	;
	goto L7
}
func F_hashinetextended(m *base.Module, l0 int32) int32 {
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
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
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
	var v54 int32
	_ = v54
	var v58 int32
	_ = v58
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
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
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v161 int32
	_ = v161
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
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
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
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
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v316 int32
	_ = v316
	var v320 int32
	_ = v320
	var v324 int32
	_ = v324
	var v328 int32
	_ = v328
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum_packed(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
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
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
		v28 = v20 - int32(1636608432)
		if v22 == int64(0) {
			v65 = v28
			v67 = v28
			v69 = v28
		} else {
			v32 = v28 + base.I32_wrap_i64(v22)
			v33 = v32 + v28
			v37 = int32(4)
			v39 = base.I32_wrap_i64(int64(base.Ui64(v22)>>(uint(int64(32))%64))) ^ base.I32_rotl(v28, v37)
			v43 = v32 - v39 ^ base.I32_rotl(v39, int32(6))
			v47 = v33 - v43 ^ base.I32_rotl(v43, int32(8))
			v48 = v39 + v33
			v49 = v43 + v48
			v50 = v47 + v49
			v54 = v48 - v47 ^ base.I32_rotl(v47, int32(16))
			v58 = v49 - v54 ^ base.I32_rotl(v54, int32(19))
			v63 = v54 + v50
			v65 = v63
			v67 = v50 - v58 ^ base.I32_rotl(v58, v37)
			v69 = v58 + v63
		}
		if v14&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v20) {
				v74 = v14
				v75 = v20
				v77 = v65
				v78 = v69
				v79 = v67
				for {
					v81 = *(*int32)(unsafe.Add(mBase, uint32(v74)+4))
					v82 = v81 + v78
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v74)+8))
					v86 = v85 + v79
					v88 = int32(4)
					v90 = v83 + v77 - v86 ^ base.I32_rotl(v86, v88)
					v94 = v82 - v90 ^ base.I32_rotl(v90, int32(6))
					v95 = v86 + v82
					v96 = v90 + v95
					v97 = v94 + v96
					v101 = v95 - v94 ^ base.I32_rotl(v94, int32(8))
					v105 = v96 - v101 ^ base.I32_rotl(v101, int32(16))
					v109 = v97 - v105 ^ base.I32_rotl(v105, int32(19))
					v110 = v101 + v97
					v111 = v105 + v110
					v112 = v109 + v111
					v116 = v110 - v109 ^ base.I32_rotl(v109, v88)
					v117 = int32(12)
					v118 = v74 + v117
					v120 = v75 - v117
					if base.Ui32(int32(11)) < base.Ui32(v120) {
						v74 = v118
						v75 = v120
						v77 = v111
						v78 = v112
						v79 = v116
						continue
					} else {
						break
					}
					break
				}
				v123 = v118
				v124 = v120
				v126 = v111
				v127 = v112
				v128 = v116
			} else {
				v123 = v14
				v124 = v20
				v126 = v65
				v127 = v69
				v128 = v67
			}
			switch v124 - int32(1) {
			case 0:
				v293 = v126
				v294 = v127
				v295 = v128
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 1:
				v286 = v126
				v287 = v127
				v288 = v128
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 2:
				v279 = v126
				v280 = v127
				v281 = v128
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 3:
				v273 = v127
				v274 = v128
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
				v279 = v275<<(uint(int32(24))%32) + v126
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 4:
				v269 = v127
				v270 = v128
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
				v279 = v275<<(uint(int32(24))%32) + v126
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 5:
				v263 = v127
				v264 = v128
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
				v279 = v275<<(uint(int32(24))%32) + v126
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 6:
				v257 = v127
				v258 = v128
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
				v279 = v275<<(uint(int32(24))%32) + v126
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 7:
				v252 = v128
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+7)))
				v257 = v253<<(uint(int32(24))%32) + v127
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
				v279 = v275<<(uint(int32(24))%32) + v126
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 8:
				v247 = v128
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+8)))
				v252 = v248<<(uint(int32(8))%32) + v247
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+7)))
				v257 = v253<<(uint(int32(24))%32) + v127
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
				v279 = v275<<(uint(int32(24))%32) + v126
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 9:
				v242 = v128
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+9)))
				v247 = v243<<(uint(int32(16))%32) + v242
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+8)))
				v252 = v248<<(uint(int32(8))%32) + v247
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+7)))
				v257 = v253<<(uint(int32(24))%32) + v127
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
				v279 = v275<<(uint(int32(24))%32) + v126
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			case 10:
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+10)))
				v242 = v238<<(uint(int32(24))%32) + v128
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+9)))
				v247 = v243<<(uint(int32(16))%32) + v242
				v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+8)))
				v252 = v248<<(uint(int32(8))%32) + v247
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+7)))
				v257 = v253<<(uint(int32(24))%32) + v127
				v258 = v252
				v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+6)))
				v263 = v259<<(uint(int32(16))%32) + v257
				v264 = v258
				v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+5)))
				v269 = v265<<(uint(int32(8))%32) + v263
				v270 = v264
				v271 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+4)))
				v273 = v269 + v271
				v274 = v270
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+3)))
				v279 = v275<<(uint(int32(24))%32) + v126
				v280 = v273
				v281 = v274
				v282 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+2)))
				v286 = v282<<(uint(int32(16))%32) + v279
				v287 = v280
				v288 = v281
				v289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123)+1)))
				v293 = v289<<(uint(int32(8))%32) + v286
				v294 = v287
				v295 = v288
				v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v123))))
				v301 = v293 + v296
				v302 = v294
				v303 = v295
			default:
				v301 = v126
				v302 = v127
				v303 = v128
			}
		} else {
			if base.Ui32(int32(12)) <= base.Ui32(v20) {
				v134 = v14
				v135 = v20
				v137 = v65
				v138 = v69
				v139 = v67
				for {
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
					v142 = v141 + v138
					v143 = *(*int32)(unsafe.Add(mBase, uint32(v134)))
					v145 = *(*int32)(unsafe.Add(mBase, uint32(v134)+8))
					v146 = v145 + v139
					v148 = int32(4)
					v150 = v143 + v137 - v146 ^ base.I32_rotl(v146, v148)
					v154 = v142 - v150 ^ base.I32_rotl(v150, int32(6))
					v155 = v146 + v142
					v156 = v150 + v155
					v157 = v154 + v156
					v161 = v155 - v154 ^ base.I32_rotl(v154, int32(8))
					v165 = v156 - v161 ^ base.I32_rotl(v161, int32(16))
					v169 = v157 - v165 ^ base.I32_rotl(v165, int32(19))
					v170 = v161 + v157
					v171 = v165 + v170
					v172 = v169 + v171
					v176 = v170 - v169 ^ base.I32_rotl(v169, v148)
					v177 = int32(12)
					v178 = v134 + v177
					v180 = v135 - v177
					if base.Ui32(int32(11)) < base.Ui32(v180) {
						v134 = v178
						v135 = v180
						v137 = v171
						v138 = v172
						v139 = v176
						continue
					} else {
						break
					}
					break
				}
				v183 = v178
				v184 = v180
				v186 = v171
				v187 = v172
				v188 = v176
			} else {
				v183 = v14
				v184 = v20
				v186 = v65
				v187 = v69
				v188 = v67
			}
			switch v184 - int32(1) {
			case 0:
				v235 = v186
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
				v301 = v235 + v236
				v302 = v187
				v303 = v188
			case 1:
				v230 = v186
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
				v235 = v231<<(uint(int32(8))%32) + v230
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
				v301 = v235 + v236
				v302 = v187
				v303 = v188
			case 2:
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+2)))
				v230 = v226<<(uint(int32(16))%32) + v186
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+1)))
				v235 = v231<<(uint(int32(8))%32) + v230
				v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183))))
				v301 = v235 + v236
				v302 = v187
				v303 = v188
			case 3:
				v223 = v187
				v224 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
				v301 = v224 + v186
				v302 = v223
				v303 = v188
			case 4:
				v220 = v187
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
				v223 = v220 + v221
				v224 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
				v301 = v224 + v186
				v302 = v223
				v303 = v188
			case 5:
				v215 = v187
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+5)))
				v220 = v216<<(uint(int32(8))%32) + v215
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
				v223 = v220 + v221
				v224 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
				v301 = v224 + v186
				v302 = v223
				v303 = v188
			case 6:
				v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+6)))
				v215 = v211<<(uint(int32(16))%32) + v187
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+5)))
				v220 = v216<<(uint(int32(8))%32) + v215
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+4)))
				v223 = v220 + v221
				v224 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
				v301 = v224 + v186
				v302 = v223
				v303 = v188
			case 7:
				v206 = v188
				v207 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
				v301 = v207 + v186
				v302 = v209 + v187
				v303 = v206
			case 8:
				v201 = v188
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+8)))
				v206 = v202<<(uint(int32(8))%32) + v201
				v207 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
				v301 = v207 + v186
				v302 = v209 + v187
				v303 = v206
			case 9:
				v196 = v188
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+9)))
				v201 = v197<<(uint(int32(16))%32) + v196
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+8)))
				v206 = v202<<(uint(int32(8))%32) + v201
				v207 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
				v301 = v207 + v186
				v302 = v209 + v187
				v303 = v206
			case 10:
				v192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+10)))
				v196 = v192<<(uint(int32(24))%32) + v188
				v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+9)))
				v201 = v197<<(uint(int32(16))%32) + v196
				v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v183)+8)))
				v206 = v202<<(uint(int32(8))%32) + v201
				v207 = *(*int32)(unsafe.Add(mBase, uint32(v183)))
				v209 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
				v301 = v207 + v186
				v302 = v209 + v187
				v303 = v206
			default:
				v301 = v186
				v302 = v187
				v303 = v188
			}
		}
		v306 = int32(14)
		v308 = v302 ^ v303 - base.I32_rotl(v302, v306)
		v312 = v308 ^ v301 - base.I32_rotl(v308, int32(11))
		v316 = v312 ^ v302 - base.I32_rotl(v312, int32(25))
		v320 = v316 ^ v308 - base.I32_rotl(v316, int32(16))
		v324 = v320 ^ v312 - base.I32_rotl(v320, int32(4))
		v328 = v324 ^ v316 - base.I32_rotl(v324, v306)
		v338 = F_Int64GetDatum(m, base.I64_extend_i32_u(v328)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v320^v328-base.I32_rotl(v328, int32(24))))
		mBase = m.M
		v339 = m.ExcPending
		if v339 != 0 {
			return int32(0)
		} else {
			return v338
		}
	}
}
func F_hashmacaddr8extended(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
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
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
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
	var v279 int32
	_ = v279
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v291 int32
	_ = v291
	var v295 int32
	_ = v295
	var v299 int32
	_ = v299
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v11 = int32(-1636608424)
	if v5 == int64(0) {
		v48 = v11
		v50 = v11
		v52 = v11
	} else {
		v14 = base.I32_wrap_i64(v5)
		v16 = v14 + int32(1021750448)
		v20 = int32(4)
		v22 = base.I32_wrap_i64(int64(base.Ui64(v5)>>(uint(int64(32))%64))) ^ base.I32_rotl(v11, v20)
		v26 = v11 + v14 - v22 ^ base.I32_rotl(v22, int32(6))
		v30 = v16 - v26 ^ base.I32_rotl(v26, int32(8))
		v31 = v22 + v16
		v32 = v26 + v31
		v33 = v30 + v32
		v37 = v31 - v30 ^ base.I32_rotl(v30, int32(16))
		v41 = v32 - v37 ^ base.I32_rotl(v37, int32(19))
		v46 = v37 + v33
		v48 = v46
		v50 = v33 - v41 ^ base.I32_rotl(v41, v20)
		v52 = v41 + v46
	}
	if v2&int32(3) != 0 {
		switch int32(7) {
		case 0:
			v276 = v48
			v277 = v52
			v278 = v50
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 1:
			v269 = v48
			v270 = v52
			v271 = v50
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 2:
			v262 = v48
			v263 = v52
			v264 = v50
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 3:
			v256 = v52
			v257 = v50
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 4:
			v252 = v52
			v253 = v50
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 5:
			v246 = v52
			v247 = v50
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 6:
			v240 = v52
			v241 = v50
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 7:
			v235 = v50
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v240 = v236<<(uint(int32(24))%32) + v52
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 8:
			v230 = v50
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v240 = v236<<(uint(int32(24))%32) + v52
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 9:
			v225 = v50
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v230 = v226<<(uint(int32(16))%32) + v225
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v240 = v236<<(uint(int32(24))%32) + v52
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		case 10:
			v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v225 = v221<<(uint(int32(24))%32) + v50
			v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v230 = v226<<(uint(int32(16))%32) + v225
			v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v235 = v231<<(uint(int32(8))%32) + v230
			v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+7)))
			v240 = v236<<(uint(int32(24))%32) + v52
			v241 = v235
			v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v246 = v242<<(uint(int32(16))%32) + v240
			v247 = v241
			v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v252 = v248<<(uint(int32(8))%32) + v246
			v253 = v247
			v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v256 = v252 + v254
			v257 = v253
			v258 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+3)))
			v262 = v258<<(uint(int32(24))%32) + v48
			v263 = v256
			v264 = v257
			v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v269 = v265<<(uint(int32(16))%32) + v262
			v270 = v263
			v271 = v264
			v272 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v276 = v272<<(uint(int32(8))%32) + v269
			v277 = v270
			v278 = v271
			v279 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v276 + v279
			v285 = v277
			v286 = v278
		default:
			v284 = v48
			v285 = v52
			v286 = v50
		}
	} else {
		switch int32(7) {
		case 0:
			v218 = v48
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v218 + v219
			v285 = v52
			v286 = v50
		case 1:
			v213 = v48
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v218 = v214<<(uint(int32(8))%32) + v213
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v218 + v219
			v285 = v52
			v286 = v50
		case 2:
			v209 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+2)))
			v213 = v209<<(uint(int32(16))%32) + v48
			v214 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+1)))
			v218 = v214<<(uint(int32(8))%32) + v213
			v219 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2))))
			v284 = v218 + v219
			v285 = v52
			v286 = v50
		case 3:
			v206 = v52
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v284 = v207 + v48
			v285 = v206
			v286 = v50
		case 4:
			v203 = v52
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v284 = v207 + v48
			v285 = v206
			v286 = v50
		case 5:
			v198 = v52
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v284 = v207 + v48
			v285 = v206
			v286 = v50
		case 6:
			v194 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+6)))
			v198 = v194<<(uint(int32(16))%32) + v52
			v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+5)))
			v203 = v199<<(uint(int32(8))%32) + v198
			v204 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+4)))
			v206 = v203 + v204
			v207 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v284 = v207 + v48
			v285 = v206
			v286 = v50
		case 7:
			v189 = v50
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v284 = v190 + v48
			v285 = v192 + v52
			v286 = v189
		case 8:
			v184 = v50
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v284 = v190 + v48
			v285 = v192 + v52
			v286 = v189
		case 9:
			v179 = v50
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v284 = v190 + v48
			v285 = v192 + v52
			v286 = v189
		case 10:
			v175 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+10)))
			v179 = v175<<(uint(int32(24))%32) + v50
			v180 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+9)))
			v184 = v180<<(uint(int32(16))%32) + v179
			v185 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2)+8)))
			v189 = v185<<(uint(int32(8))%32) + v184
			v190 = *(*int32)(unsafe.Add(mBase, uint32(v2)))
			v192 = *(*int32)(unsafe.Add(mBase, uint32(v2)+4))
			v284 = v190 + v48
			v285 = v192 + v52
			v286 = v189
		default:
			v284 = v48
			v285 = v52
			v286 = v50
		}
	}
	v289 = int32(14)
	v291 = v285 ^ v286 - base.I32_rotl(v285, v289)
	v295 = v291 ^ v284 - base.I32_rotl(v291, int32(11))
	v299 = v295 ^ v285 - base.I32_rotl(v295, int32(25))
	v303 = v299 ^ v291 - base.I32_rotl(v299, int32(16))
	v307 = v303 ^ v295 - base.I32_rotl(v303, int32(4))
	v311 = v307 ^ v299 - base.I32_rotl(v307, v289)
	v321 = F_Int64GetDatum(m, base.I64_extend_i32_u(v311)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v303^v311-base.I32_rotl(v311, int32(24))))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		return int32(0)
	} else {
		return v321
	}
}
func F_hashtext(m *base.Module, l0 int32) int32 {
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
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
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
	var v102 int32
	_ = v102
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
	var v113 int32
	_ = v113
	var v117 int32
	_ = v117
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
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
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
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
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int32
	_ = v250
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
	var v259 int32
	_ = v259
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
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
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v861 int32
	_ = v861
	var v864 int32
	_ = v864
	var v865 int32
	_ = v865
	var v869 int32
	_ = v869
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v877 int32
	_ = v877
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
	var v901 int32
	_ = v901
	var v903 int32
	_ = v903
	var v904 int32
	_ = v904
	var v905 int32
	_ = v905
	var v907 int32
	_ = v907
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v912 int32
	_ = v912
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v927 int32
	_ = v927
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v938 int32
	_ = v938
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v953 int32
	_ = v953
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v963 int32
	_ = v963
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v972 int32
	_ = v972
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1006 int32
	_ = v1006
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1014 int32
	_ = v1014
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1021 int32
	_ = v1021
	var v1025 int32
	_ = v1025
	var v1029 int32
	_ = v1029
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1035 int32
	_ = v1035
	var v1036 int32
	_ = v1036
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1044 int32
	_ = v1044
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1055 int32
	_ = v1055
	var v1059 int32
	_ = v1059
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1074 int32
	_ = v1074
	var v1075 int32
	_ = v1075
	var v1076 int32
	_ = v1076
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1099 int32
	_ = v1099
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1118 int32
	_ = v1118
	var v1119 int32
	_ = v1119
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1132 int32
	_ = v1132
	var v1136 int32
	_ = v1136
	var v1140 int32
	_ = v1140
	var v1144 int32
	_ = v1144
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1156 int32
	_ = v1156
	var v1159 int32
	_ = v1159
	var v1164 int32
	_ = v1164
	var v1167 int32
	_ = v1167
	var v1171 int32
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1180 int32
	_ = v1180
	var v1184 int32
	_ = v1184
	var v1188 int32
	_ = v1188
	var v1193 int32
	_ = v1193
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v14 != 0 {
			v15 = F_pg_newlocale_from_collation(m, v14)
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				v17 = int32(1)
				v18 = v10 + v17
				v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10))))
				v23 = v21 & v17
				if v23 != 0 {
					v24 = v18
				} else {
					v24 = v10 + int32(4)
				}
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v15)+1)))
				if v25 == int32(1) {
					if v21 == int32(1) {
						v30 = int32(4)
						v32 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v32&int32(254) == int32(2) {
							v41 = v30
						} else {
							v41 = base.B2i32(v32 == int32(18)) << (uint(v30) % 32)
						}
						if v32 == int32(1) {
							v44 = v30
						} else {
							v44 = v41
						}
						v50 = v44 - int32(1636608432)
						if v24&int32(3) != 0 {
							if base.Ui32(int32(11)) < base.Ui32(v44) {
								v159 = v24
								v160 = v44
								v161 = v50
								v162 = v50
								v163 = v50
								for {
									v165 = *(*int32)(unsafe.Add(mBase, uint32(v159)+4))
									v166 = v165 + v162
									v167 = *(*int32)(unsafe.Add(mBase, uint32(v159)))
									v169 = *(*int32)(unsafe.Add(mBase, uint32(v159)+8))
									v170 = v169 + v163
									v172 = int32(4)
									v174 = v167 + v161 - v170 ^ base.I32_rotl(v170, v172)
									v178 = v166 - v174 ^ base.I32_rotl(v174, int32(6))
									v179 = v170 + v166
									v180 = v174 + v179
									v181 = v178 + v180
									v185 = v179 - v178 ^ base.I32_rotl(v178, int32(8))
									v189 = v180 - v185 ^ base.I32_rotl(v185, int32(16))
									v193 = v181 - v189 ^ base.I32_rotl(v189, int32(19))
									v194 = v185 + v181
									v195 = v189 + v194
									v196 = v193 + v195
									v200 = v194 - v193 ^ base.I32_rotl(v193, v172)
									v201 = int32(12)
									v202 = v159 + v201
									v204 = v160 - v201
									if base.Ui32(int32(11)) < base.Ui32(v204) {
										v159 = v202
										v160 = v204
										v161 = v195
										v162 = v196
										v163 = v200
										continue
									} else {
										break
									}
									break
								}
								v207 = v202
								v208 = v204
								v209 = v195
								v210 = v196
								v211 = v200
							} else {
								v207 = v24
								v208 = v44
								v209 = v50
								v210 = v50
								v211 = v50
							}
							switch v208 - int32(1) {
							case 0:
								v270 = v209
								v271 = v210
								v272 = v211
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 1:
								v263 = v209
								v264 = v210
								v265 = v211
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 2:
								v256 = v209
								v257 = v210
								v258 = v211
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2)))
								v263 = v259<<(uint(int32(16))%32) + v256
								v264 = v257
								v265 = v258
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 3:
								v250 = v210
								v251 = v211
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+3)))
								v256 = v252<<(uint(int32(24))%32) + v209
								v257 = v250
								v258 = v251
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2)))
								v263 = v259<<(uint(int32(16))%32) + v256
								v264 = v257
								v265 = v258
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 4:
								v246 = v210
								v247 = v211
								v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
								v250 = v246 + v248
								v251 = v247
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+3)))
								v256 = v252<<(uint(int32(24))%32) + v209
								v257 = v250
								v258 = v251
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2)))
								v263 = v259<<(uint(int32(16))%32) + v256
								v264 = v257
								v265 = v258
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 5:
								v240 = v210
								v241 = v211
								v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+5)))
								v246 = v242<<(uint(int32(8))%32) + v240
								v247 = v241
								v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
								v250 = v246 + v248
								v251 = v247
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+3)))
								v256 = v252<<(uint(int32(24))%32) + v209
								v257 = v250
								v258 = v251
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2)))
								v263 = v259<<(uint(int32(16))%32) + v256
								v264 = v257
								v265 = v258
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 6:
								v234 = v210
								v235 = v211
								v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+6)))
								v240 = v236<<(uint(int32(16))%32) + v234
								v241 = v235
								v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+5)))
								v246 = v242<<(uint(int32(8))%32) + v240
								v247 = v241
								v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
								v250 = v246 + v248
								v251 = v247
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+3)))
								v256 = v252<<(uint(int32(24))%32) + v209
								v257 = v250
								v258 = v251
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2)))
								v263 = v259<<(uint(int32(16))%32) + v256
								v264 = v257
								v265 = v258
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 7:
								v229 = v211
								v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+7)))
								v234 = v230<<(uint(int32(24))%32) + v210
								v235 = v229
								v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+6)))
								v240 = v236<<(uint(int32(16))%32) + v234
								v241 = v235
								v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+5)))
								v246 = v242<<(uint(int32(8))%32) + v240
								v247 = v241
								v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
								v250 = v246 + v248
								v251 = v247
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+3)))
								v256 = v252<<(uint(int32(24))%32) + v209
								v257 = v250
								v258 = v251
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2)))
								v263 = v259<<(uint(int32(16))%32) + v256
								v264 = v257
								v265 = v258
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 8:
								v224 = v211
								v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+8)))
								v229 = v225<<(uint(int32(8))%32) + v224
								v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+7)))
								v234 = v230<<(uint(int32(24))%32) + v210
								v235 = v229
								v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+6)))
								v240 = v236<<(uint(int32(16))%32) + v234
								v241 = v235
								v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+5)))
								v246 = v242<<(uint(int32(8))%32) + v240
								v247 = v241
								v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
								v250 = v246 + v248
								v251 = v247
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+3)))
								v256 = v252<<(uint(int32(24))%32) + v209
								v257 = v250
								v258 = v251
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2)))
								v263 = v259<<(uint(int32(16))%32) + v256
								v264 = v257
								v265 = v258
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 9:
								v219 = v211
								v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+9)))
								v224 = v220<<(uint(int32(16))%32) + v219
								v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+8)))
								v229 = v225<<(uint(int32(8))%32) + v224
								v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+7)))
								v234 = v230<<(uint(int32(24))%32) + v210
								v235 = v229
								v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+6)))
								v240 = v236<<(uint(int32(16))%32) + v234
								v241 = v235
								v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+5)))
								v246 = v242<<(uint(int32(8))%32) + v240
								v247 = v241
								v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
								v250 = v246 + v248
								v251 = v247
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+3)))
								v256 = v252<<(uint(int32(24))%32) + v209
								v257 = v250
								v258 = v251
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2)))
								v263 = v259<<(uint(int32(16))%32) + v256
								v264 = v257
								v265 = v258
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							case 10:
								v215 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+10)))
								v219 = v215<<(uint(int32(24))%32) + v211
								v220 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+9)))
								v224 = v220<<(uint(int32(16))%32) + v219
								v225 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+8)))
								v229 = v225<<(uint(int32(8))%32) + v224
								v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+7)))
								v234 = v230<<(uint(int32(24))%32) + v210
								v235 = v229
								v236 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+6)))
								v240 = v236<<(uint(int32(16))%32) + v234
								v241 = v235
								v242 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+5)))
								v246 = v242<<(uint(int32(8))%32) + v240
								v247 = v241
								v248 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+4)))
								v250 = v246 + v248
								v251 = v247
								v252 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+3)))
								v256 = v252<<(uint(int32(24))%32) + v209
								v257 = v250
								v258 = v251
								v259 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+2)))
								v263 = v259<<(uint(int32(16))%32) + v256
								v264 = v257
								v265 = v258
								v266 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207)+1)))
								v270 = v266<<(uint(int32(8))%32) + v263
								v271 = v264
								v272 = v265
								v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v207))))
								v277 = v270 + v273
								v278 = v271
								v279 = v272
							default:
								v277 = v209
								v278 = v210
								v279 = v211
							}
						} else {
							if base.Ui32(v44) < base.Ui32(int32(12)) {
								v105 = v24
								v106 = v44
								v107 = v50
								v108 = v50
								v109 = v50
							} else {
								v57 = v24
								v58 = v44
								v59 = v50
								v60 = v50
								v61 = v50
								for {
									v63 = *(*int32)(unsafe.Add(mBase, uint32(v57)+4))
									v64 = v63 + v60
									v65 = *(*int32)(unsafe.Add(mBase, uint32(v57)))
									v67 = *(*int32)(unsafe.Add(mBase, uint32(v57)+8))
									v68 = v67 + v61
									v70 = int32(4)
									v72 = v65 + v59 - v68 ^ base.I32_rotl(v68, v70)
									v76 = v64 - v72 ^ base.I32_rotl(v72, int32(6))
									v77 = v68 + v64
									v78 = v72 + v77
									v79 = v76 + v78
									v83 = v77 - v76 ^ base.I32_rotl(v76, int32(8))
									v87 = v78 - v83 ^ base.I32_rotl(v83, int32(16))
									v91 = v79 - v87 ^ base.I32_rotl(v87, int32(19))
									v92 = v83 + v79
									v93 = v87 + v92
									v94 = v91 + v93
									v98 = v92 - v91 ^ base.I32_rotl(v91, v70)
									v99 = int32(12)
									v100 = v57 + v99
									v102 = v58 - v99
									if base.Ui32(int32(11)) < base.Ui32(v102) {
										v57 = v100
										v58 = v102
										v59 = v93
										v60 = v94
										v61 = v98
										continue
									} else {
										break
									}
									break
								}
								v105 = v100
								v106 = v102
								v107 = v93
								v108 = v94
								v109 = v98
							}
							switch v106 - int32(1) {
							case 0:
								v156 = v107
								v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
								v277 = v156 + v157
								v278 = v108
								v279 = v109
							case 1:
								v151 = v107
								v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
								v156 = v152<<(uint(int32(8))%32) + v151
								v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
								v277 = v156 + v157
								v278 = v108
								v279 = v109
							case 2:
								v147 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+2)))
								v151 = v147<<(uint(int32(16))%32) + v107
								v152 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+1)))
								v156 = v152<<(uint(int32(8))%32) + v151
								v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105))))
								v277 = v156 + v157
								v278 = v108
								v279 = v109
							case 3:
								v144 = v108
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
								v277 = v145 + v107
								v278 = v144
								v279 = v109
							case 4:
								v141 = v108
								v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+4)))
								v144 = v141 + v142
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
								v277 = v145 + v107
								v278 = v144
								v279 = v109
							case 5:
								v136 = v108
								v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+5)))
								v141 = v137<<(uint(int32(8))%32) + v136
								v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+4)))
								v144 = v141 + v142
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
								v277 = v145 + v107
								v278 = v144
								v279 = v109
							case 6:
								v132 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+6)))
								v136 = v132<<(uint(int32(16))%32) + v108
								v137 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+5)))
								v141 = v137<<(uint(int32(8))%32) + v136
								v142 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+4)))
								v144 = v141 + v142
								v145 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
								v277 = v145 + v107
								v278 = v144
								v279 = v109
							case 7:
								v127 = v109
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
								v277 = v128 + v107
								v278 = v130 + v108
								v279 = v127
							case 8:
								v122 = v109
								v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+8)))
								v127 = v123<<(uint(int32(8))%32) + v122
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
								v277 = v128 + v107
								v278 = v130 + v108
								v279 = v127
							case 9:
								v117 = v109
								v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+9)))
								v122 = v118<<(uint(int32(16))%32) + v117
								v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+8)))
								v127 = v123<<(uint(int32(8))%32) + v122
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
								v277 = v128 + v107
								v278 = v130 + v108
								v279 = v127
							case 10:
								v113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+10)))
								v117 = v113<<(uint(int32(24))%32) + v109
								v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+9)))
								v122 = v118<<(uint(int32(16))%32) + v117
								v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v105)+8)))
								v127 = v123<<(uint(int32(8))%32) + v122
								v128 = *(*int32)(unsafe.Add(mBase, uint32(v105)))
								v130 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
								v277 = v128 + v107
								v278 = v130 + v108
								v279 = v127
							default:
								v277 = v107
								v278 = v108
								v279 = v109
							}
						}
						v282 = int32(14)
						v284 = v278 ^ v279 - base.I32_rotl(v278, v282)
						v288 = v284 ^ v277 - base.I32_rotl(v284, int32(11))
						v292 = v288 ^ v278 - base.I32_rotl(v288, int32(25))
						v296 = v292 ^ v284 - base.I32_rotl(v292, int32(16))
						v300 = v296 ^ v288 - base.I32_rotl(v296, int32(4))
						v304 = v300 ^ v292 - base.I32_rotl(v300, v282)
						v1151 = v304 ^ v296 - base.I32_rotl(v304, int32(24))
					} else {
						if v23 != 0 {
							v309 = int32(1)
							v312 = int32(base.Ui32(v21)>>(uint(v309)%32)) - v309
							v318 = v312 - int32(1636608432)
							if v24&int32(3) != 0 {
								if base.Ui32(int32(11)) < base.Ui32(v312) {
									v427 = v24
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
									v475 = v24
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
									v373 = v24
									v374 = v312
									v375 = v318
									v376 = v318
									v377 = v318
								} else {
									v325 = v24
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
							v1151 = v572 ^ v564 - base.I32_rotl(v572, int32(24))
						} else {
							v577 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v581 = int32(base.Ui32(v577)>>(uint(int32(2))%32)) - int32(4)
							v587 = v581 - int32(1636608432)
							if v24&int32(3) != 0 {
								if base.Ui32(int32(11)) < base.Ui32(v581) {
									v696 = v24
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
									v744 = v24
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
									v642 = v24
									v643 = v581
									v644 = v587
									v645 = v587
									v646 = v587
								} else {
									v594 = v24
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
							v1151 = v841 ^ v833 - base.I32_rotl(v841, int32(24))
						}
					}
					v1156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v1156 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v1159 = m.ExcPending
						if v1159 != 0 {
							return int32(0)
						} else {
							return v1151
						}
					} else {
						return v1151
					}
				} else {
					v846 = int32(0)
					if v21 == int32(1) {
						v850 = int32(4)
						v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v852&int32(254) == int32(2) {
							v861 = v850
						} else {
							v861 = base.B2i32(v852 == int32(18)) << (uint(v850) % 32)
						}
						if v852 == int32(1) {
							v864 = v850
						} else {
							v864 = v861
						}
						v875 = v864
					} else {
						v865 = int32(1)
						if v23 != 0 {
							v875 = int32(base.Ui32(v21)>>(uint(v865)%32)) - v865
						} else {
							v869 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v875 = int32(base.Ui32(v869)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v876 = F_pg_strnxfrm(m, v846, v846, v24, v875, v15)
					mBase = m.M
					v877 = m.ExcPending
					if v877 != 0 {
						return int32(0)
					} else {
						v879 = v876 + int32(1)
						v880 = F_palloc(m, v879)
						mBase = m.M
						v881 = m.ExcPending
						if v881 != 0 {
							return int32(0)
						} else {
							v882 = F_pg_strnxfrm(m, v880, v879, v24, v875, v15)
							mBase = m.M
							v883 = m.ExcPending
							if v883 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v876) < base.Ui32(v882) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1184 = m.ExcPending
									if v1184 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(98112), int32(0))
										mBase = m.M
										v1188 = m.ExcPending
										if v1188 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(500127), int32(306), int32(62633))
											mBase = m.M
											v1193 = m.ExcPending
											if v1193 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v890 = v879 - int32(1636608432)
									if v880&int32(3) != 0 {
										if base.Ui32(int32(11)) < base.Ui32(v879) {
											v999 = v880
											v1000 = v879
											v1001 = v890
											v1002 = v890
											v1003 = v890
											for {
												v1005 = *(*int32)(unsafe.Add(mBase, uint32(v999)+4))
												v1006 = v1005 + v1002
												v1007 = *(*int32)(unsafe.Add(mBase, uint32(v999)))
												v1009 = *(*int32)(unsafe.Add(mBase, uint32(v999)+8))
												v1010 = v1009 + v1003
												v1012 = int32(4)
												v1014 = v1007 + v1001 - v1010 ^ base.I32_rotl(v1010, v1012)
												v1018 = v1006 - v1014 ^ base.I32_rotl(v1014, int32(6))
												v1019 = v1010 + v1006
												v1020 = v1014 + v1019
												v1021 = v1018 + v1020
												v1025 = v1019 - v1018 ^ base.I32_rotl(v1018, int32(8))
												v1029 = v1020 - v1025 ^ base.I32_rotl(v1025, int32(16))
												v1033 = v1021 - v1029 ^ base.I32_rotl(v1029, int32(19))
												v1034 = v1025 + v1021
												v1035 = v1029 + v1034
												v1036 = v1033 + v1035
												v1040 = v1034 - v1033 ^ base.I32_rotl(v1033, v1012)
												v1041 = int32(12)
												v1042 = v999 + v1041
												v1044 = v1000 - v1041
												if base.Ui32(int32(11)) < base.Ui32(v1044) {
													v999 = v1042
													v1000 = v1044
													v1001 = v1035
													v1002 = v1036
													v1003 = v1040
													continue
												} else {
													break
												}
												break
											}
											v1047 = v1042
											v1048 = v1044
											v1049 = v1035
											v1050 = v1036
											v1051 = v1040
										} else {
											v1047 = v880
											v1048 = v879
											v1049 = v890
											v1050 = v890
											v1051 = v890
										}
										switch v1048 - int32(1) {
										case 0:
											v1110 = v1049
											v1111 = v1050
											v1112 = v1051
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 1:
											v1103 = v1049
											v1104 = v1050
											v1105 = v1051
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 2:
											v1096 = v1049
											v1097 = v1050
											v1098 = v1051
											v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+2)))
											v1103 = v1099<<(uint(int32(16))%32) + v1096
											v1104 = v1097
											v1105 = v1098
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 3:
											v1090 = v1050
											v1091 = v1051
											v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+3)))
											v1096 = v1092<<(uint(int32(24))%32) + v1049
											v1097 = v1090
											v1098 = v1091
											v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+2)))
											v1103 = v1099<<(uint(int32(16))%32) + v1096
											v1104 = v1097
											v1105 = v1098
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 4:
											v1086 = v1050
											v1087 = v1051
											v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+4)))
											v1090 = v1086 + v1088
											v1091 = v1087
											v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+3)))
											v1096 = v1092<<(uint(int32(24))%32) + v1049
											v1097 = v1090
											v1098 = v1091
											v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+2)))
											v1103 = v1099<<(uint(int32(16))%32) + v1096
											v1104 = v1097
											v1105 = v1098
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 5:
											v1080 = v1050
											v1081 = v1051
											v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+5)))
											v1086 = v1082<<(uint(int32(8))%32) + v1080
											v1087 = v1081
											v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+4)))
											v1090 = v1086 + v1088
											v1091 = v1087
											v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+3)))
											v1096 = v1092<<(uint(int32(24))%32) + v1049
											v1097 = v1090
											v1098 = v1091
											v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+2)))
											v1103 = v1099<<(uint(int32(16))%32) + v1096
											v1104 = v1097
											v1105 = v1098
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 6:
											v1074 = v1050
											v1075 = v1051
											v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+6)))
											v1080 = v1076<<(uint(int32(16))%32) + v1074
											v1081 = v1075
											v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+5)))
											v1086 = v1082<<(uint(int32(8))%32) + v1080
											v1087 = v1081
											v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+4)))
											v1090 = v1086 + v1088
											v1091 = v1087
											v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+3)))
											v1096 = v1092<<(uint(int32(24))%32) + v1049
											v1097 = v1090
											v1098 = v1091
											v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+2)))
											v1103 = v1099<<(uint(int32(16))%32) + v1096
											v1104 = v1097
											v1105 = v1098
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 7:
											v1069 = v1051
											v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+7)))
											v1074 = v1070<<(uint(int32(24))%32) + v1050
											v1075 = v1069
											v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+6)))
											v1080 = v1076<<(uint(int32(16))%32) + v1074
											v1081 = v1075
											v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+5)))
											v1086 = v1082<<(uint(int32(8))%32) + v1080
											v1087 = v1081
											v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+4)))
											v1090 = v1086 + v1088
											v1091 = v1087
											v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+3)))
											v1096 = v1092<<(uint(int32(24))%32) + v1049
											v1097 = v1090
											v1098 = v1091
											v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+2)))
											v1103 = v1099<<(uint(int32(16))%32) + v1096
											v1104 = v1097
											v1105 = v1098
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 8:
											v1064 = v1051
											v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+8)))
											v1069 = v1065<<(uint(int32(8))%32) + v1064
											v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+7)))
											v1074 = v1070<<(uint(int32(24))%32) + v1050
											v1075 = v1069
											v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+6)))
											v1080 = v1076<<(uint(int32(16))%32) + v1074
											v1081 = v1075
											v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+5)))
											v1086 = v1082<<(uint(int32(8))%32) + v1080
											v1087 = v1081
											v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+4)))
											v1090 = v1086 + v1088
											v1091 = v1087
											v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+3)))
											v1096 = v1092<<(uint(int32(24))%32) + v1049
											v1097 = v1090
											v1098 = v1091
											v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+2)))
											v1103 = v1099<<(uint(int32(16))%32) + v1096
											v1104 = v1097
											v1105 = v1098
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 9:
											v1059 = v1051
											v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+9)))
											v1064 = v1060<<(uint(int32(16))%32) + v1059
											v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+8)))
											v1069 = v1065<<(uint(int32(8))%32) + v1064
											v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+7)))
											v1074 = v1070<<(uint(int32(24))%32) + v1050
											v1075 = v1069
											v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+6)))
											v1080 = v1076<<(uint(int32(16))%32) + v1074
											v1081 = v1075
											v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+5)))
											v1086 = v1082<<(uint(int32(8))%32) + v1080
											v1087 = v1081
											v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+4)))
											v1090 = v1086 + v1088
											v1091 = v1087
											v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+3)))
											v1096 = v1092<<(uint(int32(24))%32) + v1049
											v1097 = v1090
											v1098 = v1091
											v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+2)))
											v1103 = v1099<<(uint(int32(16))%32) + v1096
											v1104 = v1097
											v1105 = v1098
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										case 10:
											v1055 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+10)))
											v1059 = v1055<<(uint(int32(24))%32) + v1051
											v1060 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+9)))
											v1064 = v1060<<(uint(int32(16))%32) + v1059
											v1065 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+8)))
											v1069 = v1065<<(uint(int32(8))%32) + v1064
											v1070 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+7)))
											v1074 = v1070<<(uint(int32(24))%32) + v1050
											v1075 = v1069
											v1076 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+6)))
											v1080 = v1076<<(uint(int32(16))%32) + v1074
											v1081 = v1075
											v1082 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+5)))
											v1086 = v1082<<(uint(int32(8))%32) + v1080
											v1087 = v1081
											v1088 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+4)))
											v1090 = v1086 + v1088
											v1091 = v1087
											v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+3)))
											v1096 = v1092<<(uint(int32(24))%32) + v1049
											v1097 = v1090
											v1098 = v1091
											v1099 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+2)))
											v1103 = v1099<<(uint(int32(16))%32) + v1096
											v1104 = v1097
											v1105 = v1098
											v1106 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047)+1)))
											v1110 = v1106<<(uint(int32(8))%32) + v1103
											v1111 = v1104
											v1112 = v1105
											v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1047))))
											v1117 = v1110 + v1113
											v1118 = v1111
											v1119 = v1112
										default:
											v1117 = v1049
											v1118 = v1050
											v1119 = v1051
										}
									} else {
										if base.Ui32(v879) < base.Ui32(int32(12)) {
											v945 = v880
											v946 = v879
											v947 = v890
											v948 = v890
											v949 = v890
										} else {
											v897 = v880
											v898 = v879
											v899 = v890
											v900 = v890
											v901 = v890
											for {
												v903 = *(*int32)(unsafe.Add(mBase, uint32(v897)+4))
												v904 = v903 + v900
												v905 = *(*int32)(unsafe.Add(mBase, uint32(v897)))
												v907 = *(*int32)(unsafe.Add(mBase, uint32(v897)+8))
												v908 = v907 + v901
												v910 = int32(4)
												v912 = v905 + v899 - v908 ^ base.I32_rotl(v908, v910)
												v916 = v904 - v912 ^ base.I32_rotl(v912, int32(6))
												v917 = v908 + v904
												v918 = v912 + v917
												v919 = v916 + v918
												v923 = v917 - v916 ^ base.I32_rotl(v916, int32(8))
												v927 = v918 - v923 ^ base.I32_rotl(v923, int32(16))
												v931 = v919 - v927 ^ base.I32_rotl(v927, int32(19))
												v932 = v923 + v919
												v933 = v927 + v932
												v934 = v931 + v933
												v938 = v932 - v931 ^ base.I32_rotl(v931, v910)
												v939 = int32(12)
												v940 = v897 + v939
												v942 = v898 - v939
												if base.Ui32(int32(11)) < base.Ui32(v942) {
													v897 = v940
													v898 = v942
													v899 = v933
													v900 = v934
													v901 = v938
													continue
												} else {
													break
												}
												break
											}
											v945 = v940
											v946 = v942
											v947 = v933
											v948 = v934
											v949 = v938
										}
										switch v946 - int32(1) {
										case 0:
											v996 = v947
											v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945))))
											v1117 = v996 + v997
											v1118 = v948
											v1119 = v949
										case 1:
											v991 = v947
											v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+1)))
											v996 = v992<<(uint(int32(8))%32) + v991
											v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945))))
											v1117 = v996 + v997
											v1118 = v948
											v1119 = v949
										case 2:
											v987 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+2)))
											v991 = v987<<(uint(int32(16))%32) + v947
											v992 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+1)))
											v996 = v992<<(uint(int32(8))%32) + v991
											v997 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945))))
											v1117 = v996 + v997
											v1118 = v948
											v1119 = v949
										case 3:
											v984 = v948
											v985 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
											v1117 = v985 + v947
											v1118 = v984
											v1119 = v949
										case 4:
											v981 = v948
											v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+4)))
											v984 = v981 + v982
											v985 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
											v1117 = v985 + v947
											v1118 = v984
											v1119 = v949
										case 5:
											v976 = v948
											v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+5)))
											v981 = v977<<(uint(int32(8))%32) + v976
											v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+4)))
											v984 = v981 + v982
											v985 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
											v1117 = v985 + v947
											v1118 = v984
											v1119 = v949
										case 6:
											v972 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+6)))
											v976 = v972<<(uint(int32(16))%32) + v948
											v977 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+5)))
											v981 = v977<<(uint(int32(8))%32) + v976
											v982 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+4)))
											v984 = v981 + v982
											v985 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
											v1117 = v985 + v947
											v1118 = v984
											v1119 = v949
										case 7:
											v967 = v949
											v968 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
											v970 = *(*int32)(unsafe.Add(mBase, uint32(v945)+4))
											v1117 = v968 + v947
											v1118 = v970 + v948
											v1119 = v967
										case 8:
											v962 = v949
											v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+8)))
											v967 = v963<<(uint(int32(8))%32) + v962
											v968 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
											v970 = *(*int32)(unsafe.Add(mBase, uint32(v945)+4))
											v1117 = v968 + v947
											v1118 = v970 + v948
											v1119 = v967
										case 9:
											v957 = v949
											v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+9)))
											v962 = v958<<(uint(int32(16))%32) + v957
											v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+8)))
											v967 = v963<<(uint(int32(8))%32) + v962
											v968 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
											v970 = *(*int32)(unsafe.Add(mBase, uint32(v945)+4))
											v1117 = v968 + v947
											v1118 = v970 + v948
											v1119 = v967
										case 10:
											v953 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+10)))
											v957 = v953<<(uint(int32(24))%32) + v949
											v958 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+9)))
											v962 = v958<<(uint(int32(16))%32) + v957
											v963 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v945)+8)))
											v967 = v963<<(uint(int32(8))%32) + v962
											v968 = *(*int32)(unsafe.Add(mBase, uint32(v945)))
											v970 = *(*int32)(unsafe.Add(mBase, uint32(v945)+4))
											v1117 = v968 + v947
											v1118 = v970 + v948
											v1119 = v967
										default:
											v1117 = v947
											v1118 = v948
											v1119 = v949
										}
									}
									v1122 = int32(14)
									v1124 = v1118 ^ v1119 - base.I32_rotl(v1118, v1122)
									v1128 = v1124 ^ v1117 - base.I32_rotl(v1124, int32(11))
									v1132 = v1128 ^ v1118 - base.I32_rotl(v1128, int32(25))
									v1136 = v1132 ^ v1124 - base.I32_rotl(v1132, int32(16))
									v1140 = v1136 ^ v1128 - base.I32_rotl(v1136, int32(4))
									v1144 = v1140 ^ v1132 - base.I32_rotl(v1140, v1122)
									F_pfree(m, v880)
									mBase = m.M
									v1150 = m.ExcPending
									if v1150 != 0 {
										return int32(0)
									} else {
										v1151 = v1144 ^ v1136 - base.I32_rotl(v1144, int32(24))
										v1156 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v1156 != v10 {
											F_pfree(m, v10)
											mBase = m.M
											v1159 = m.ExcPending
											if v1159 != 0 {
												return int32(0)
											} else {
												return v1151
											}
										} else {
											return v1151
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
			v1164 = m.ExcPending
			if v1164 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v1167 = m.ExcPending
				if v1167 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(335732), int32(0))
					mBase = m.M
					v1171 = m.ExcPending
					if v1171 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(574955), int32(0))
						mBase = m.M
						v1175 = m.ExcPending
						if v1175 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(500127), int32(281), int32(62633))
							mBase = m.M
							v1180 = m.ExcPending
							if v1180 != 0 {
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
}
func F_hashvarlena(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v46 int32
	_ = v46
	var v52 int32
	_ = v52
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
	var v65 int32
	_ = v65
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
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
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
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
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
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
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
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
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
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v284 int32
	_ = v284
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
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v314 int32
	_ = v314
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum_packed(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(1)
		v12 = v7 + v11
		v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7))))
		v17 = v15 & v11
		if v17 != 0 {
			v18 = v12
		} else {
			v18 = v7 + int32(4)
		}
		if v15 == int32(1) {
			v21 = int32(4)
			v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v23&int32(254) == int32(2) {
				v32 = v21
			} else {
				v32 = base.B2i32(v23 == int32(18)) << (uint(v21) % 32)
			}
			if v23 == int32(1) {
				v35 = v21
			} else {
				v35 = v32
			}
			v46 = v35
		} else {
			v36 = int32(1)
			if v17 != 0 {
				v46 = int32(base.Ui32(v15)>>(uint(v36)%32)) - v36
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v46 = int32(base.Ui32(v40)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v52 = v46 - int32(1636608432)
		if v18&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v46) {
				v161 = v18
				v162 = v46
				v163 = v52
				v164 = v52
				v165 = v52
				for {
					v167 = *(*int32)(unsafe.Add(mBase, uint32(v161)+4))
					v168 = v167 + v164
					v169 = *(*int32)(unsafe.Add(mBase, uint32(v161)))
					v171 = *(*int32)(unsafe.Add(mBase, uint32(v161)+8))
					v172 = v171 + v165
					v174 = int32(4)
					v176 = v169 + v163 - v172 ^ base.I32_rotl(v172, v174)
					v180 = v168 - v176 ^ base.I32_rotl(v176, int32(6))
					v181 = v172 + v168
					v182 = v176 + v181
					v183 = v180 + v182
					v187 = v181 - v180 ^ base.I32_rotl(v180, int32(8))
					v191 = v182 - v187 ^ base.I32_rotl(v187, int32(16))
					v195 = v183 - v191 ^ base.I32_rotl(v191, int32(19))
					v196 = v187 + v183
					v197 = v191 + v196
					v198 = v195 + v197
					v202 = v196 - v195 ^ base.I32_rotl(v195, v174)
					v203 = int32(12)
					v204 = v161 + v203
					v206 = v162 - v203
					if base.Ui32(int32(11)) < base.Ui32(v206) {
						v161 = v204
						v162 = v206
						v163 = v197
						v164 = v198
						v165 = v202
						continue
					} else {
						break
					}
					break
				}
				v209 = v204
				v210 = v206
				v211 = v197
				v212 = v198
				v213 = v202
			} else {
				v209 = v18
				v210 = v46
				v211 = v52
				v212 = v52
				v213 = v52
			}
			switch v210 - int32(1) {
			case 0:
				v272 = v211
				v273 = v212
				v274 = v213
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 1:
				v265 = v211
				v266 = v212
				v267 = v213
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 2:
				v258 = v211
				v259 = v212
				v260 = v213
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v265 = v261<<(uint(int32(16))%32) + v258
				v266 = v259
				v267 = v260
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 3:
				v252 = v212
				v253 = v213
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
				v258 = v254<<(uint(int32(24))%32) + v211
				v259 = v252
				v260 = v253
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v265 = v261<<(uint(int32(16))%32) + v258
				v266 = v259
				v267 = v260
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 4:
				v248 = v212
				v249 = v213
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v252 = v248 + v250
				v253 = v249
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
				v258 = v254<<(uint(int32(24))%32) + v211
				v259 = v252
				v260 = v253
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v265 = v261<<(uint(int32(16))%32) + v258
				v266 = v259
				v267 = v260
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 5:
				v242 = v212
				v243 = v213
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)))
				v248 = v244<<(uint(int32(8))%32) + v242
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v252 = v248 + v250
				v253 = v249
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
				v258 = v254<<(uint(int32(24))%32) + v211
				v259 = v252
				v260 = v253
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v265 = v261<<(uint(int32(16))%32) + v258
				v266 = v259
				v267 = v260
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 6:
				v236 = v212
				v237 = v213
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+6)))
				v242 = v238<<(uint(int32(16))%32) + v236
				v243 = v237
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)))
				v248 = v244<<(uint(int32(8))%32) + v242
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v252 = v248 + v250
				v253 = v249
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
				v258 = v254<<(uint(int32(24))%32) + v211
				v259 = v252
				v260 = v253
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v265 = v261<<(uint(int32(16))%32) + v258
				v266 = v259
				v267 = v260
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 7:
				v231 = v213
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+7)))
				v236 = v232<<(uint(int32(24))%32) + v212
				v237 = v231
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+6)))
				v242 = v238<<(uint(int32(16))%32) + v236
				v243 = v237
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)))
				v248 = v244<<(uint(int32(8))%32) + v242
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v252 = v248 + v250
				v253 = v249
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
				v258 = v254<<(uint(int32(24))%32) + v211
				v259 = v252
				v260 = v253
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v265 = v261<<(uint(int32(16))%32) + v258
				v266 = v259
				v267 = v260
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 8:
				v226 = v213
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+8)))
				v231 = v227<<(uint(int32(8))%32) + v226
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+7)))
				v236 = v232<<(uint(int32(24))%32) + v212
				v237 = v231
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+6)))
				v242 = v238<<(uint(int32(16))%32) + v236
				v243 = v237
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)))
				v248 = v244<<(uint(int32(8))%32) + v242
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v252 = v248 + v250
				v253 = v249
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
				v258 = v254<<(uint(int32(24))%32) + v211
				v259 = v252
				v260 = v253
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v265 = v261<<(uint(int32(16))%32) + v258
				v266 = v259
				v267 = v260
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 9:
				v221 = v213
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+9)))
				v226 = v222<<(uint(int32(16))%32) + v221
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+8)))
				v231 = v227<<(uint(int32(8))%32) + v226
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+7)))
				v236 = v232<<(uint(int32(24))%32) + v212
				v237 = v231
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+6)))
				v242 = v238<<(uint(int32(16))%32) + v236
				v243 = v237
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)))
				v248 = v244<<(uint(int32(8))%32) + v242
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v252 = v248 + v250
				v253 = v249
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
				v258 = v254<<(uint(int32(24))%32) + v211
				v259 = v252
				v260 = v253
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v265 = v261<<(uint(int32(16))%32) + v258
				v266 = v259
				v267 = v260
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			case 10:
				v217 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+10)))
				v221 = v217<<(uint(int32(24))%32) + v213
				v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+9)))
				v226 = v222<<(uint(int32(16))%32) + v221
				v227 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+8)))
				v231 = v227<<(uint(int32(8))%32) + v226
				v232 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+7)))
				v236 = v232<<(uint(int32(24))%32) + v212
				v237 = v231
				v238 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+6)))
				v242 = v238<<(uint(int32(16))%32) + v236
				v243 = v237
				v244 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+5)))
				v248 = v244<<(uint(int32(8))%32) + v242
				v249 = v243
				v250 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+4)))
				v252 = v248 + v250
				v253 = v249
				v254 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+3)))
				v258 = v254<<(uint(int32(24))%32) + v211
				v259 = v252
				v260 = v253
				v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+2)))
				v265 = v261<<(uint(int32(16))%32) + v258
				v266 = v259
				v267 = v260
				v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209)+1)))
				v272 = v268<<(uint(int32(8))%32) + v265
				v273 = v266
				v274 = v267
				v275 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v209))))
				v279 = v272 + v275
				v280 = v273
				v281 = v274
			default:
				v279 = v211
				v280 = v212
				v281 = v213
			}
		} else {
			if base.Ui32(v46) < base.Ui32(int32(12)) {
				v107 = v18
				v108 = v46
				v109 = v52
				v110 = v52
				v111 = v52
			} else {
				v59 = v18
				v60 = v46
				v61 = v52
				v62 = v52
				v63 = v52
				for {
					v65 = *(*int32)(unsafe.Add(mBase, uint32(v59)+4))
					v66 = v65 + v62
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v59)))
					v69 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
					v70 = v69 + v63
					v72 = int32(4)
					v74 = v67 + v61 - v70 ^ base.I32_rotl(v70, v72)
					v78 = v66 - v74 ^ base.I32_rotl(v74, int32(6))
					v79 = v70 + v66
					v80 = v74 + v79
					v81 = v78 + v80
					v85 = v79 - v78 ^ base.I32_rotl(v78, int32(8))
					v89 = v80 - v85 ^ base.I32_rotl(v85, int32(16))
					v93 = v81 - v89 ^ base.I32_rotl(v89, int32(19))
					v94 = v85 + v81
					v95 = v89 + v94
					v96 = v93 + v95
					v100 = v94 - v93 ^ base.I32_rotl(v93, v72)
					v101 = int32(12)
					v102 = v59 + v101
					v104 = v60 - v101
					if base.Ui32(int32(11)) < base.Ui32(v104) {
						v59 = v102
						v60 = v104
						v61 = v95
						v62 = v96
						v63 = v100
						continue
					} else {
						break
					}
					break
				}
				v107 = v102
				v108 = v104
				v109 = v95
				v110 = v96
				v111 = v100
			}
			switch v108 - int32(1) {
			case 0:
				v158 = v109
				v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
				v279 = v158 + v159
				v280 = v110
				v281 = v111
			case 1:
				v153 = v109
				v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
				v158 = v154<<(uint(int32(8))%32) + v153
				v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
				v279 = v158 + v159
				v280 = v110
				v281 = v111
			case 2:
				v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+2)))
				v153 = v149<<(uint(int32(16))%32) + v109
				v154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+1)))
				v158 = v154<<(uint(int32(8))%32) + v153
				v159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107))))
				v279 = v158 + v159
				v280 = v110
				v281 = v111
			case 3:
				v146 = v110
				v147 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
				v279 = v147 + v109
				v280 = v146
				v281 = v111
			case 4:
				v143 = v110
				v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
				v146 = v143 + v144
				v147 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
				v279 = v147 + v109
				v280 = v146
				v281 = v111
			case 5:
				v138 = v110
				v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
				v143 = v139<<(uint(int32(8))%32) + v138
				v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
				v146 = v143 + v144
				v147 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
				v279 = v147 + v109
				v280 = v146
				v281 = v111
			case 6:
				v134 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+6)))
				v138 = v134<<(uint(int32(16))%32) + v110
				v139 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+5)))
				v143 = v139<<(uint(int32(8))%32) + v138
				v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+4)))
				v146 = v143 + v144
				v147 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
				v279 = v147 + v109
				v280 = v146
				v281 = v111
			case 7:
				v129 = v111
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
				v132 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
				v279 = v130 + v109
				v280 = v132 + v110
				v281 = v129
			case 8:
				v124 = v111
				v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+8)))
				v129 = v125<<(uint(int32(8))%32) + v124
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
				v132 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
				v279 = v130 + v109
				v280 = v132 + v110
				v281 = v129
			case 9:
				v119 = v111
				v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+9)))
				v124 = v120<<(uint(int32(16))%32) + v119
				v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+8)))
				v129 = v125<<(uint(int32(8))%32) + v124
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
				v132 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
				v279 = v130 + v109
				v280 = v132 + v110
				v281 = v129
			case 10:
				v115 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+10)))
				v119 = v115<<(uint(int32(24))%32) + v111
				v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+9)))
				v124 = v120<<(uint(int32(16))%32) + v119
				v125 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v107)+8)))
				v129 = v125<<(uint(int32(8))%32) + v124
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v107)))
				v132 = *(*int32)(unsafe.Add(mBase, uint32(v107)+4))
				v279 = v130 + v109
				v280 = v132 + v110
				v281 = v129
			default:
				v279 = v109
				v280 = v110
				v281 = v111
			}
		}
		v284 = int32(14)
		v286 = v280 ^ v281 - base.I32_rotl(v280, v284)
		v290 = v286 ^ v279 - base.I32_rotl(v286, int32(11))
		v294 = v290 ^ v280 - base.I32_rotl(v290, int32(25))
		v298 = v294 ^ v286 - base.I32_rotl(v294, int32(16))
		v302 = v298 ^ v290 - base.I32_rotl(v298, int32(4))
		v306 = v302 ^ v294 - base.I32_rotl(v302, v284)
		v311 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v311 != v7 {
			F_pfree(m, v7)
			mBase = m.M
			v314 = m.ExcPending
			if v314 != 0 {
				return int32(0)
			} else {
				return v306 ^ v298 - base.I32_rotl(v306, int32(24))
			}
		} else {
			return v306 ^ v298 - base.I32_rotl(v306, int32(24))
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
	var v12 int32
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
			v12 = *(*int32)(unsafe.Add(mBase, _consts[239]))
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
func F_headline_json_value(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
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
		v24 = F_FunctionCall3Coll(m, v10+int32(112), int32(0), v13, v9, v11)
		mBase = m.M
		v25 = m.ExcPending
		if v25 != 0 {
			return int32(0)
		} else {
			v26 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+20)) = uint8(v26)
			v28 = F_generateHeadline(m, v13)
			mBase = m.M
			v29 = m.ExcPending
			if v29 != 0 {
				return int32(0)
			} else {
				return v28
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
	var v40 int32
	_ = v40
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
	var v67 int32
	_ = v67
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
	var v123 int32
	_ = v123
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
	var v149 int32
	_ = v149
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
	var v207 int32
	_ = v207
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
	var v237 int32
	_ = v237
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v274 int32
	_ = v274
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v328 int32
	_ = v328
	v7 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(16)
	m.G0 = v20
	if l2 == v7 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v20 + int32(16)
	return v328
L2:
	;
	v328 = v7
	goto L1
L3:
	;
	goto L4
L4:
	;
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v40 = v26
	goto L5
L5:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v44 <= int32(0) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v129 = int32(0)
	if v123 < v129 {
		v328 = v129
		goto L1
	} else {
		goto L26
	}
L8:
	;
	v123 = int32(-1)
	goto L7
L9:
	;
	goto L10
L10:
	;
	v48 = int32(0)
	if v48 < v44 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v51 = v44
	goto L13
L12:
	;
	v51 = v48
	goto L13
L13:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v54 = int32(0)
	v64 = v54
	v67 = int32(-1)
	goto L14
L14:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v52+v64<<(uint(int32(2))%32))))
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v76)))
	if v77 <= int32(0) {
		v328 = v54
		goto L1
	} else {
		goto L16
	}
L15:
	;
	v123 = v108
	goto L7
L16:
	;
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v76)+8))
	v83 = int32(0)
	goto L18
L17:
	;
	if v67 < v102 {
		goto L22
	} else {
		goto L23
	}
L18:
	;
	v102 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v80+v83<<(uint(int32(1))%32)))))
	if v40 <= v102 {
		goto L17
	} else {
		goto L20
	}
L19:
	;
	v328 = v54
	goto L1
L20:
	;
	v105 = v83 + int32(1)
	if v105 != v77 {
		v83 = v105
		goto L18
	} else {
		goto L21
	}
L21:
	;
	goto L19
L22:
	;
	v108 = v102
	goto L24
L23:
	;
	v108 = v67
	goto L24
L24:
	;
	v110 = v64 + int32(1)
	if v110 != v51 {
		v64 = v110
		v67 = v108
		goto L14
	} else {
		goto L25
	}
L25:
	;
	goto L15
L26:
	;
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if v132 <= int32(0) {
		goto L28
	} else {
		goto L29
	}
L27:
	;
	if v40 < v207 {
		goto L46
	} else {
		goto L47
	}
L28:
	;
	v207 = int32(2147483646)
	goto L27
L29:
	;
	goto L30
L30:
	;
	v136 = int32(0)
	if v136 < v132 {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v139 = v132
	goto L33
L32:
	;
	v139 = v136
	goto L33
L33:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v149 = int32(2147483646)
	v151 = int32(0)
	goto L34
L34:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v140+v151<<(uint(int32(2))%32))))
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v163)))
	v166 = v164
	goto L36
L35:
	;
	v207 = v197
	goto L27
L36:
	;
	v184 = v166 - int32(1)
	if v184 < int32(0) {
		v195 = int32(-1)
		goto L38
	} else {
		goto L39
	}
L37:
	;
	if v195 < v149 {
		goto L41
	} else {
		goto L42
	}
L38:
	;
	goto L37
L39:
	;
	v187 = *(*int32)(unsafe.Add(mBase, uint32(v163)+8))
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v187+v184<<(uint(int32(1))%32)))))
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v163)+12))
	v193 = v191 - v192
	if v123 < v193 {
		v166 = v184
		goto L36
	} else {
		goto L40
	}
L40:
	;
	v195 = v193
	goto L38
L41:
	;
	v197 = v195
	goto L43
L42:
	;
	v197 = v149
	goto L43
L43:
	;
	v199 = v151 + int32(1)
	if v199 != v139 {
		v149 = v197
		v151 = v199
		goto L34
	} else {
		goto L44
	}
L44:
	;
	goto L35
L45:
	;
	v40 = v219 + int32(1)
	goto L5
L46:
	;
	v219 = v207
	goto L48
L47:
	;
	v219 = v40
	goto L48
L48:
	;
	if v123 < v219 {
		goto L45
	} else {
		goto L49
	}
L49:
	;
	v221 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v221 <= int32(0) {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v225 = int32(-1)
	v229 = int32(0)
	v234 = v225
	v237 = v225
	goto L51
L51:
	;
	v247 = v224 + v229<<(uint(int32(4))%32)
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v247)+12))
	if v248 == int32(0) {
		goto L55
	} else {
		goto L56
	}
L52:
	;
	if v267 < int32(0) {
		goto L45
	} else {
		goto L67
	}
L53:
	;
	goto L52
L54:
	;
	v264 = v229 + int32(1)
	if v264 != v221 {
		v229 = v264
		v234 = v260
		v237 = v261
		goto L51
	} else {
		goto L66
	}
L55:
	;
	v260 = v234
	v261 = v237
	goto L54
L56:
	;
	goto L57
L57:
	;
	if int32(0) <= v234 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	if v123 < v258 {
		v267 = v257
		v270 = v237
		goto L53
	} else {
		goto L65
	}
L59:
	;
	v253 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)))
	v257 = v234
	v258 = v253
	goto L58
L60:
	;
	goto L61
L61:
	;
	v254 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v247)+4)))
	if v254 < v219 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v256 = v234
	goto L64
L63:
	;
	v256 = v229
	goto L64
L64:
	;
	v257 = v256
	v258 = v254
	goto L58
L65:
	;
	v260 = v257
	v261 = v229
	goto L54
L66:
	;
	v267 = v260
	v270 = v261
	goto L53
L67:
	;
	if v270 < v267 {
		goto L45
	} else {
		goto L68
	}
L68:
	;
	v274 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v20)+12)) = v270 - v267 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v20)+8)) = v274 + v267<<(uint(int32(4))%32)
	v287 = F_TS_execute(m, l1+int32(8), v20+int32(8), int32(0), int32(1191))
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	return int32(0)
L70:
	;
	if v287 == int32(0) {
		goto L45
	} else {
		goto L71
	}
L71:
	;
	v293 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v219 + v293
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v267
	*(*int32)(unsafe.Add(mBase, uint32(l5))) = v270
	v328 = v293
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
	var v33 int32
	_ = v33
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
									v33 = F___memset(m, v14, int32(0), v12)
									mBase = m.M
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
func F_hs_contains(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v5 int32
	_ = v5
	v2 = F_hstore_contains(m, l0)
	v5 = m.ExcPending
	if v5 != 0 {
		return int32(0)
	} else {
		return v2
	}
}
func F_hypothetical_check_argtypes(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	if l2 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L12
	} else {
		goto L18
	}
L2:
	;
	v11 = l1 + int32(1)
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	if v11 != v12 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l2+v11<<(uint(int32(4))%32)+l1*int32(100))+88))
	if v20 != int32(23) {
		goto L1
	} else {
		goto L4
	}
L4:
	;
	v23 = int32(0)
	if v23 < l1 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v27 = l1
	goto L7
L6:
	;
	v27 = v23
	goto L7
L7:
	;
	v33 = v23
	goto L9
L8:
	;
	return
L9:
	;
	if v33 == v27 {
		goto L8
	} else {
		goto L11
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L12
	} else {
		goto L15
	}
L11:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v43 = v33 + int32(1)
	v44 = F_get_fn_expr_argtype(m, v41, v43)
	mBase = m.M
	v45 = m.ExcPending
	if v45 != 0 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	return
L13:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v33*int32(100)+(l2+int32(88)+v40<<(uint(int32(4))%32)))))
	if v44 == v50 {
		v33 = v43
		goto L9
	} else {
		goto L14
	}
L14:
	;
	goto L10
L15:
	;
	F_errmsg_internal(m, int32(252024), int32(0))
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L12
	} else {
		goto L16
	}
L16:
	;
	F_errfinish(m, int32(494215), int32(1159), int32(161794))
	mBase = m.M
	v64 = m.ExcPending
	if v64 != 0 {
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
	F_errmsg_internal(m, int32(252024), int32(0))
	mBase = m.M
	v73 = m.ExcPending
	if v73 != 0 {
		goto L12
	} else {
		goto L19
	}
L19:
	;
	F_errfinish(m, int32(494215), int32(1151), int32(161794))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
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
func F_hypothetical_cume_dist_final(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	var v15 int64
	_ = v15
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	v3 = m.G0
	v5 = v3 - int32(16)
	m.G0 = v5
	v10 = F_hypothetical_rank_common(m, l0, int32(1), v5+int32(8))
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v5)+8))
		v20 = F_Float8GetDatum(m, base.F64_div(base.F64_convert_i64_s(v10), base.F64_convert_i64_s(v15+int64(1))))
		mBase = m.M
		v21 = m.ExcPending
		if v21 != 0 {
			return int32(0)
		} else {
			m.G0 = v5 + int32(16)
			return v20
		}
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
	var v56 int32
	_ = v56
	var v63 int32
	_ = v63
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v115 int32
	_ = v115
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int64
	_ = v165
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v189 int64
	_ = v189
	var v190 int32
	_ = v190
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v225 int64
	_ = v225
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
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
	v22 = l0 + int32(20)
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
	v234 = m.ExcPending
	if v234 != 0 {
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
		v140 = v46
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v141+v140<<(uint(int32(2))%32)))) = l1
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v148 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v146+v140))) = uint8(v148)
	v150 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
	v152 = v150 & int32(65533)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)) = uint16(v152)
	v154 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v154)))
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+6)) = uint16(v155)
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
	v56 = v49
	v63 = int32(0)
	goto L15
L13:
	;
	v101 = v49
	goto L14
L14:
	;
	if v27&int32(2) == int32(0) {
		v140 = v33
		goto L10
	} else {
		goto L18
	}
L15:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v67 = int32(2)
	v71 = v56 | int32(1)
	v72 = int32(3)
	v74 = v22 + v71<<(uint(v72)%32)
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v74)))
	*(*int32)(unsafe.Add(mBase, uint32(v66+v56<<(uint(v67)%32)))) = v75
	v77 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v77+v56))) = uint8(v79)
	v81 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v86 = v56 + v67
	v89 = v22 + v86<<(uint(v72)%32)
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
	*(*int32)(unsafe.Add(mBase, uint32(v81+v71<<(uint(v67)%32)))) = v90
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v94 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v71+v92))) = uint8(v94)
	v97 = v63 + v67
	if v97 != v33&int32(2147483646) {
		v56 = v86
		v63 = v97
		goto L15
	} else {
		goto L17
	}
L16:
	;
	v101 = v86
	goto L14
L17:
	;
	goto L16
L18:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v121 = v101<<(uint(int32(3))%32) + v22
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v121)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v115+v101<<(uint(int32(2))%32)))) = v122
	v124 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v121)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v124+v101))) = uint8(v126)
	v140 = v33
	goto L10
L19:
	;
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_tuplesort_puttupleslot(m, v157, v41)
	mBase = m.M
	v159 = m.ExcPending
	if v159 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v160 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_tuplesort_performsort(m, v160)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v163 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)) = uint8(v163)
	v165 = int64(1)
	v166 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v170 = F_tuplesort_gettupleslot(m, v166, v163, v163, v41, int32(0))
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	v226 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v226)+12))
	m.T0[v227].(func(*base.Module, int32))(m, v41)
	mBase = m.M
	v229 = m.ExcPending
	if v229 != 0 {
		goto L7
	} else {
		goto L41
	}
L23:
	;
	if v170 == int32(0) {
		v225 = v165
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v189 = v165
	goto L25
L25:
	;
	v190 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+6)))
	if v190 <= v33 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v225 = v207
	goto L22
L27:
	;
	F_slot_getsomeattrs_int(m, v41, v33+int32(1))
	mBase = m.M
	v193 = m.ExcPending
	if v193 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v196 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194+v33))))
	if v196 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v199 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v201 = *(*int32)(unsafe.Add(mBase, uint32(v199+v33<<(uint(int32(2))%32))))
	if v201 != 0 {
		v225 = v189
		goto L22
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v203 = *(*int32)(unsafe.Add(mBase, _consts[44]))
	if v203 != 0 {
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
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v207 = v189 + int64(1)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v209 = int32(1)
	v212 = F_tuplesort_gettupleslot(m, v208, v209, v209, v41, int32(0))
	mBase = m.M
	v213 = m.ExcPending
	if v213 != 0 {
		goto L7
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	if v212 != 0 {
		v189 = v207
		goto L25
	} else {
		goto L40
	}
L40:
	;
	goto L26
L41:
	;
	return v225
L42:
	;
	F_errmsg_internal(m, int32(251969), int32(0))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(494215), int32(1194), int32(245855))
	mBase = m.M
	v243 = m.ExcPending
	if v243 != 0 {
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
