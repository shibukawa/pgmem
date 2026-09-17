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
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_id[0]))
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v15 = F_convert_any_priv_string(m, v10, int32(_a_F_has_largeobject_privilege_id_0))
		mBase = m.M
		v16 = m.ExcPending
		if v16 != 0 {
			return int32(0)
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
				return int32(0)
			} else {
				if v25 != 0 {
					v29 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_has_largeobject_privilege_id[2])))
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
	v13 = v10 + int32(12)
	v15 = v10 + int32(11)
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
				v28 = F_index_form_tuple(m, v23, v10+int32(12), v10+int32(11))
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
	var v44 int32
	_ = v44
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
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
	var v181 int32
	_ = v181
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
	var v239 float64
	_ = v239
	var v244 float64
	_ = v244
	var v245 float64
	_ = v245
	var v249 float64
	_ = v249
	var v251 float64
	_ = v251
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 float64
	_ = v265
	var v268 int32
	_ = v268
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v276 int32
	_ = v276
	var v279 int64
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v299 int32
	_ = v299
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v311 int32
	_ = v311
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v323 float64
	_ = v323
	var v324 float64
	_ = v324
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
	v44 = v33
	v48 = v38
	v49 = v5
	goto L3
L3:
	;
	if base.Ui32(v49) <= base.Ui32(v48) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	v232 = int32(_a_F_hashbulkdelete_0)
	v234 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[0])) = v234 + int32(1)
	v239 = *(*float64)(unsafe.Add(mBase, uint32(v220)+32))
	if base.B2i32(v48 != v38)|base.F64_ne(v37, v239) == int32(0) {
		goto L50
	} else {
		goto L51
	}
L5:
	;
	v63 = v44
	v68 = v49
	goto L8
L6:
	;
	v181 = v49
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
	if v68 == v76 {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	v181 = v110
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
	v82 = v68 + v81
	v86 = v82 - v81
	if base.Ui32(int32(2)) <= base.Ui32(v82) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v109 = *(*int32)(unsafe.Add(mBase, uint32(v63+v105<<(uint(int32(2))%32))+48))
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
	F_hashbucketcleanup(m, v23, v68, v115, v112, v156, v157, v158, v159, v21+int32(24), v21+int32(16), base.B2i32(v144 == int32(64)), l2, l3)
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
		v155 = v63
		goto L24
	} else {
		goto L29
	}
L26:
	;
	v125 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[1]))
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v125+(v115^int32(-1))<<(uint(int32(2))%32))))
	v139 = v131
	goto L25
L27:
	;
	goto L28
L28:
	;
	v133 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[2]))
	v139 = v133 + v115<<(uint(int32(13))%32) + int32(-8192)
	goto L25
L29:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v141)))
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v63)+24))
	if base.Ui32(v147) <= base.Ui32(v148) {
		v155 = v63
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
		v63 = v155
		v68 = v110
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
	v206 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[1]))
	v212 = *(*int32)(unsafe.Add(mBase, uint32(v206+(v202^int32(-1))<<(uint(int32(2))%32))))
	v220 = v212
	goto L40
L42:
	;
	goto L43
L43:
	;
	v214 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[2]))
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
	v44 = v229
	v48 = v231
	v49 = v181
	goto L3
L49:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v220)+32)) = v251
	F_MarkBufferDirty(m, v202)
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L1
	} else {
		goto L56
	}
L50:
	;
	v244 = *(*float64)(unsafe.Add(mBase, uint32(v21)+16))
	v251 = v244
	goto L49
L51:
	;
	goto L52
L52:
	;
	v245 = *(*float64)(unsafe.Add(mBase, uint32(v21)+24))
	if base.F64_gt(v239, v245) != 0 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v249 = base.F64_sub(v239, v245)
	goto L55
L54:
	;
	v249 = float64(0)
	goto L55
L55:
	;
	v251 = v249
	goto L49
L56:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v23)+48))
	v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v255)+118)))
	if v256 != int32(112) {
		goto L57
	} else {
		goto L58
	}
L57:
	;
	v305 = int32(_a_F_hashbulkdelete_0)
	v307 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[0])) = v307 - int32(1)
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_UnlockReleaseBuffer(m, v311)
	mBase = m.M
	v313 = m.ExcPending
	if v313 != 0 {
		goto L1
	} else {
		goto L72
	}
L58:
	;
	v260 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[3]))
	if v260 <= int32(0) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v23)+32))
	if v263 != 0 {
		goto L57
	} else {
		goto L62
	}
L60:
	;
	goto L61
L61:
	;
	v265 = *(*float64)(unsafe.Add(mBase, uint32(v220)+32))
	*(*float64)(unsafe.Add(mBase, uint32(v21))) = v265
	F_XLogBeginInsert(m)
	mBase = m.M
	v268 = m.ExcPending
	if v268 != 0 {
		goto L1
	} else {
		goto L64
	}
L62:
	;
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v23)+40))
	if v264 != 0 {
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
	v271 = m.ExcPending
	if v271 != 0 {
		goto L1
	} else {
		goto L65
	}
L65:
	;
	v273 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	F_XLogRegisterBuffer(m, int32(0), v273, int32(8))
	mBase = m.M
	v276 = m.ExcPending
	if v276 != 0 {
		goto L1
	} else {
		goto L66
	}
L66:
	;
	v279 = F_XLogInsert(m, int32(12), int32(176))
	mBase = m.M
	v280 = m.ExcPending
	if v280 != 0 {
		goto L1
	} else {
		goto L67
	}
L67:
	;
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v21)+12))
	if v281 < int32(0) {
		goto L69
	} else {
		goto L70
	}
L68:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v299))) = base.I64_rotr(v279, int64(32))
	goto L57
L69:
	;
	v285 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[1]))
	v291 = *(*int32)(unsafe.Add(mBase, uint32(v285+(v281^int32(-1))<<(uint(int32(2))%32))))
	v299 = v291
	goto L68
L70:
	;
	goto L71
L71:
	;
	v293 = *(*int32)(unsafe.Add(mBase, _c_F_hashbulkdelete[2]))
	v299 = v293 + v281<<(uint(int32(13))%32) + int32(-8192)
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
	v317 = F_palloc0(m, int32(40))
	mBase = m.M
	v318 = m.ExcPending
	if v318 != 0 {
		goto L1
	} else {
		goto L76
	}
L74:
	;
	v319 = l1
	goto L75
L75:
	;
	*(*float64)(unsafe.Add(mBase, uint32(v319)+8)) = v251
	v321 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v319)+4)) = uint8(v321)
	v323 = *(*float64)(unsafe.Add(mBase, uint32(v21)+24))
	v324 = *(*float64)(unsafe.Add(mBase, uint32(v319)+16))
	*(*float64)(unsafe.Add(mBase, uint32(v319)+16)) = base.F64_add(v323, v324)
	m.G0 = v21 + int32(32)
	return v319
L76:
	;
	v319 = v317
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	v2 = int32(*(*int8)(unsafe.Add(mBase, uint32(l0)+20)))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	if v4 == int64(0) {
		v11 = int32(-1636608428)
		v50 = v11
		v51 = v11
		v54 = int32(0)
	} else {
		v14 = base.I32_wrap_i64(v4)
		v16 = v14 + int32(1021750440)
		v21 = base.I32_wrap_i64(int64(base.Ui64(v4)>>(uint(int64(32))%64))) ^ int32(-415931063)
		v27 = v14 - v21 - int32(1636608428) ^ base.I32_rotl(v21, int32(6))
		v31 = v16 - v27 ^ base.I32_rotl(v27, int32(8))
		v32 = v21 + v16
		v33 = v27 + v32
		v34 = v31 + v33
		v38 = v32 - v31 ^ base.I32_rotl(v31, int32(16))
		v42 = v33 - v38 ^ base.I32_rotl(v38, int32(19))
		v47 = v38 + v34
		v48 = v42 + v47
		v50 = v48
		v51 = v47
		v54 = v34 - v42 ^ base.I32_rotl(v42, int32(4)) ^ v48
	}
	v55 = int32(14)
	v57 = v54 - base.I32_rotl(v50, v55)
	v62 = v57 ^ (v2 + v51) - base.I32_rotl(v57, int32(11))
	v66 = v50 ^ v62 - base.I32_rotl(v62, int32(25))
	v70 = v66 ^ v57 - base.I32_rotl(v66, int32(16))
	v74 = v70 ^ v62 - base.I32_rotl(v70, int32(4))
	v78 = v74 ^ v66 - base.I32_rotl(v74, v55)
	v88 = F_Int64GetDatum(m, base.I64_extend_i32_u(v78)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v78^v70-base.I32_rotl(v78, int32(24))))
	mBase = m.M
	v91 = m.ExcPending
	if v91 != 0 {
		return int32(0)
	} else {
		return v88
	}
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
			v48 = v33 + v39
			v49 = v43 + v48
			v50 = v47 + v49
			v54 = v48 - v47 ^ base.I32_rotl(v47, int32(16))
			v58 = v49 - v54 ^ base.I32_rotl(v54, int32(19))
			v63 = v50 + v54
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
		v338 = F_Int64GetDatum(m, base.I64_extend_i32_u(v328)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v328^v320-base.I32_rotl(v328, int32(24))))
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
		v31 = v16 + v22
		v32 = v26 + v31
		v33 = v30 + v32
		v37 = v31 - v30 ^ base.I32_rotl(v30, int32(16))
		v41 = v32 - v37 ^ base.I32_rotl(v37, int32(19))
		v46 = v33 + v37
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
	v321 = F_Int64GetDatum(m, base.I64_extend_i32_u(v311)<<(uint(int64(32))%64)|base.I64_extend_i32_u(v311^v303-base.I32_rotl(v311, int32(24))))
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
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
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
	var v317 int32
	_ = v317
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v339 int32
	_ = v339
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v403 int32
	_ = v403
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
	var v414 int32
	_ = v414
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v452 int32
	_ = v452
	var v456 int32
	_ = v456
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v469 int32
	_ = v469
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v507 int32
	_ = v507
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v515 int32
	_ = v515
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v519 int32
	_ = v519
	var v523 int32
	_ = v523
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
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v555 int32
	_ = v555
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v567 int32
	_ = v567
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v586 int32
	_ = v586
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v606 int32
	_ = v606
	var v608 int32
	_ = v608
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v619 int32
	_ = v619
	var v623 int32
	_ = v623
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v634 int32
	_ = v634
	var v635 int32
	_ = v635
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v663 int32
	_ = v663
	var v664 int32
	_ = v664
	var v666 int32
	_ = v666
	var v668 int32
	_ = v668
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v678 int32
	_ = v678
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v717 int32
	_ = v717
	var v721 int32
	_ = v721
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v732 int32
	_ = v732
	var v736 int32
	_ = v736
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
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
	var v751 int32
	_ = v751
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v766 int32
	_ = v766
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v792 int32
	_ = v792
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v807 int32
	_ = v807
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v813 int32
	_ = v813
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v818 int32
	_ = v818
	var v820 int32
	_ = v820
	var v824 int32
	_ = v824
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v836 int32
	_ = v836
	var v840 int32
	_ = v840
	var v845 int32
	_ = v845
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v862 int32
	_ = v862
	var v863 int32
	_ = v863
	var v867 int32
	_ = v867
	var v873 int32
	_ = v873
	var v874 int32
	_ = v874
	var v875 int32
	_ = v875
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v899 int32
	_ = v899
	var v901 int32
	_ = v901
	var v902 int32
	_ = v902
	var v903 int32
	_ = v903
	var v905 int32
	_ = v905
	var v906 int32
	_ = v906
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v914 int32
	_ = v914
	var v915 int32
	_ = v915
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v936 int32
	_ = v936
	var v937 int32
	_ = v937
	var v938 int32
	_ = v938
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v944 int32
	_ = v944
	var v945 int32
	_ = v945
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v951 int32
	_ = v951
	var v955 int32
	_ = v955
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
	var v968 int32
	_ = v968
	var v970 int32
	_ = v970
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v979 int32
	_ = v979
	var v980 int32
	_ = v980
	var v982 int32
	_ = v982
	var v983 int32
	_ = v983
	var v985 int32
	_ = v985
	var v989 int32
	_ = v989
	var v990 int32
	_ = v990
	var v994 int32
	_ = v994
	var v995 int32
	_ = v995
	var v997 int32
	_ = v997
	var v998 int32
	_ = v998
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1010 int32
	_ = v1010
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1017 int32
	_ = v1017
	var v1018 int32
	_ = v1018
	var v1019 int32
	_ = v1019
	var v1023 int32
	_ = v1023
	var v1027 int32
	_ = v1027
	var v1031 int32
	_ = v1031
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1034 int32
	_ = v1034
	var v1038 int32
	_ = v1038
	var v1039 int32
	_ = v1039
	var v1040 int32
	_ = v1040
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1046 int32
	_ = v1046
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
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
	var v1074 int32
	_ = v1074
	var v1078 int32
	_ = v1078
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1089 int32
	_ = v1089
	var v1090 int32
	_ = v1090
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1117 int32
	_ = v1117
	var v1120 int32
	_ = v1120
	var v1122 int32
	_ = v1122
	var v1126 int32
	_ = v1126
	var v1130 int32
	_ = v1130
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1154 int32
	_ = v1154
	var v1157 int32
	_ = v1157
	var v1162 int32
	_ = v1162
	var v1165 int32
	_ = v1165
	var v1169 int32
	_ = v1169
	var v1173 int32
	_ = v1173
	var v1178 int32
	_ = v1178
	var v1182 int32
	_ = v1182
	var v1186 int32
	_ = v1186
	var v1191 int32
	_ = v1191
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
						v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v33 == int32(18) {
							v36 = int32(16)
						} else {
							v36 = int32(0)
						}
						if base.Ui32((v33-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v43 = int32(4)
						} else {
							v43 = v36
						}
						v49 = v43 - int32(1636608432)
						if v24&int32(3) != 0 {
							if base.Ui32(int32(11)) < base.Ui32(v43) {
								v158 = v24
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
								v206 = v24
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
								v104 = v24
								v105 = v43
								v106 = v49
								v107 = v49
								v108 = v49
							} else {
								v56 = v24
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
						v1149 = v303 ^ v295 - base.I32_rotl(v303, int32(24))
					} else {
						if v23 != 0 {
							v308 = int32(1)
							v311 = int32(base.Ui32(v21)>>(uint(v308)%32)) - v308
							v317 = v311 - int32(1636608432)
							if v24&int32(3) != 0 {
								if base.Ui32(int32(11)) < base.Ui32(v311) {
									v426 = v24
									v427 = v311
									v428 = v317
									v429 = v317
									v430 = v317
									for {
										v432 = *(*int32)(unsafe.Add(mBase, uint32(v426)+4))
										v433 = v432 + v429
										v434 = *(*int32)(unsafe.Add(mBase, uint32(v426)))
										v436 = *(*int32)(unsafe.Add(mBase, uint32(v426)+8))
										v437 = v436 + v430
										v439 = int32(4)
										v441 = v434 + v428 - v437 ^ base.I32_rotl(v437, v439)
										v445 = v433 - v441 ^ base.I32_rotl(v441, int32(6))
										v446 = v437 + v433
										v447 = v441 + v446
										v448 = v445 + v447
										v452 = v446 - v445 ^ base.I32_rotl(v445, int32(8))
										v456 = v447 - v452 ^ base.I32_rotl(v452, int32(16))
										v460 = v448 - v456 ^ base.I32_rotl(v456, int32(19))
										v461 = v452 + v448
										v462 = v456 + v461
										v463 = v460 + v462
										v467 = v461 - v460 ^ base.I32_rotl(v460, v439)
										v468 = int32(12)
										v469 = v426 + v468
										v471 = v427 - v468
										if base.Ui32(int32(11)) < base.Ui32(v471) {
											v426 = v469
											v427 = v471
											v428 = v462
											v429 = v463
											v430 = v467
											continue
										} else {
											break
										}
										break
									}
									v474 = v469
									v475 = v471
									v476 = v462
									v477 = v463
									v478 = v467
								} else {
									v474 = v24
									v475 = v311
									v476 = v317
									v477 = v317
									v478 = v317
								}
								switch v475 - int32(1) {
								case 0:
									v537 = v476
									v538 = v477
									v539 = v478
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 1:
									v530 = v476
									v531 = v477
									v532 = v478
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 2:
									v523 = v476
									v524 = v477
									v525 = v478
									v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+2)))
									v530 = v526<<(uint(int32(16))%32) + v523
									v531 = v524
									v532 = v525
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 3:
									v517 = v477
									v518 = v478
									v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+3)))
									v523 = v519<<(uint(int32(24))%32) + v476
									v524 = v517
									v525 = v518
									v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+2)))
									v530 = v526<<(uint(int32(16))%32) + v523
									v531 = v524
									v532 = v525
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 4:
									v513 = v477
									v514 = v478
									v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+4)))
									v517 = v513 + v515
									v518 = v514
									v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+3)))
									v523 = v519<<(uint(int32(24))%32) + v476
									v524 = v517
									v525 = v518
									v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+2)))
									v530 = v526<<(uint(int32(16))%32) + v523
									v531 = v524
									v532 = v525
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 5:
									v507 = v477
									v508 = v478
									v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+5)))
									v513 = v509<<(uint(int32(8))%32) + v507
									v514 = v508
									v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+4)))
									v517 = v513 + v515
									v518 = v514
									v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+3)))
									v523 = v519<<(uint(int32(24))%32) + v476
									v524 = v517
									v525 = v518
									v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+2)))
									v530 = v526<<(uint(int32(16))%32) + v523
									v531 = v524
									v532 = v525
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 6:
									v501 = v477
									v502 = v478
									v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+6)))
									v507 = v503<<(uint(int32(16))%32) + v501
									v508 = v502
									v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+5)))
									v513 = v509<<(uint(int32(8))%32) + v507
									v514 = v508
									v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+4)))
									v517 = v513 + v515
									v518 = v514
									v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+3)))
									v523 = v519<<(uint(int32(24))%32) + v476
									v524 = v517
									v525 = v518
									v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+2)))
									v530 = v526<<(uint(int32(16))%32) + v523
									v531 = v524
									v532 = v525
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 7:
									v496 = v478
									v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+7)))
									v501 = v497<<(uint(int32(24))%32) + v477
									v502 = v496
									v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+6)))
									v507 = v503<<(uint(int32(16))%32) + v501
									v508 = v502
									v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+5)))
									v513 = v509<<(uint(int32(8))%32) + v507
									v514 = v508
									v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+4)))
									v517 = v513 + v515
									v518 = v514
									v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+3)))
									v523 = v519<<(uint(int32(24))%32) + v476
									v524 = v517
									v525 = v518
									v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+2)))
									v530 = v526<<(uint(int32(16))%32) + v523
									v531 = v524
									v532 = v525
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 8:
									v491 = v478
									v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+8)))
									v496 = v492<<(uint(int32(8))%32) + v491
									v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+7)))
									v501 = v497<<(uint(int32(24))%32) + v477
									v502 = v496
									v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+6)))
									v507 = v503<<(uint(int32(16))%32) + v501
									v508 = v502
									v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+5)))
									v513 = v509<<(uint(int32(8))%32) + v507
									v514 = v508
									v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+4)))
									v517 = v513 + v515
									v518 = v514
									v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+3)))
									v523 = v519<<(uint(int32(24))%32) + v476
									v524 = v517
									v525 = v518
									v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+2)))
									v530 = v526<<(uint(int32(16))%32) + v523
									v531 = v524
									v532 = v525
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 9:
									v486 = v478
									v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+9)))
									v491 = v487<<(uint(int32(16))%32) + v486
									v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+8)))
									v496 = v492<<(uint(int32(8))%32) + v491
									v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+7)))
									v501 = v497<<(uint(int32(24))%32) + v477
									v502 = v496
									v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+6)))
									v507 = v503<<(uint(int32(16))%32) + v501
									v508 = v502
									v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+5)))
									v513 = v509<<(uint(int32(8))%32) + v507
									v514 = v508
									v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+4)))
									v517 = v513 + v515
									v518 = v514
									v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+3)))
									v523 = v519<<(uint(int32(24))%32) + v476
									v524 = v517
									v525 = v518
									v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+2)))
									v530 = v526<<(uint(int32(16))%32) + v523
									v531 = v524
									v532 = v525
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								case 10:
									v482 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+10)))
									v486 = v482<<(uint(int32(24))%32) + v478
									v487 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+9)))
									v491 = v487<<(uint(int32(16))%32) + v486
									v492 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+8)))
									v496 = v492<<(uint(int32(8))%32) + v491
									v497 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+7)))
									v501 = v497<<(uint(int32(24))%32) + v477
									v502 = v496
									v503 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+6)))
									v507 = v503<<(uint(int32(16))%32) + v501
									v508 = v502
									v509 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+5)))
									v513 = v509<<(uint(int32(8))%32) + v507
									v514 = v508
									v515 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+4)))
									v517 = v513 + v515
									v518 = v514
									v519 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+3)))
									v523 = v519<<(uint(int32(24))%32) + v476
									v524 = v517
									v525 = v518
									v526 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+2)))
									v530 = v526<<(uint(int32(16))%32) + v523
									v531 = v524
									v532 = v525
									v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474)+1)))
									v537 = v533<<(uint(int32(8))%32) + v530
									v538 = v531
									v539 = v532
									v540 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474))))
									v544 = v537 + v540
									v545 = v538
									v546 = v539
								default:
									v544 = v476
									v545 = v477
									v546 = v478
								}
							} else {
								if base.Ui32(v311) < base.Ui32(int32(12)) {
									v372 = v24
									v373 = v311
									v374 = v317
									v375 = v317
									v376 = v317
								} else {
									v324 = v24
									v325 = v311
									v326 = v317
									v327 = v317
									v328 = v317
									for {
										v330 = *(*int32)(unsafe.Add(mBase, uint32(v324)+4))
										v331 = v330 + v327
										v332 = *(*int32)(unsafe.Add(mBase, uint32(v324)))
										v334 = *(*int32)(unsafe.Add(mBase, uint32(v324)+8))
										v335 = v334 + v328
										v337 = int32(4)
										v339 = v332 + v326 - v335 ^ base.I32_rotl(v335, v337)
										v343 = v331 - v339 ^ base.I32_rotl(v339, int32(6))
										v344 = v335 + v331
										v345 = v339 + v344
										v346 = v343 + v345
										v350 = v344 - v343 ^ base.I32_rotl(v343, int32(8))
										v354 = v345 - v350 ^ base.I32_rotl(v350, int32(16))
										v358 = v346 - v354 ^ base.I32_rotl(v354, int32(19))
										v359 = v350 + v346
										v360 = v354 + v359
										v361 = v358 + v360
										v365 = v359 - v358 ^ base.I32_rotl(v358, v337)
										v366 = int32(12)
										v367 = v324 + v366
										v369 = v325 - v366
										if base.Ui32(int32(11)) < base.Ui32(v369) {
											v324 = v367
											v325 = v369
											v326 = v360
											v327 = v361
											v328 = v365
											continue
										} else {
											break
										}
										break
									}
									v372 = v367
									v373 = v369
									v374 = v360
									v375 = v361
									v376 = v365
								}
								switch v373 - int32(1) {
								case 0:
									v423 = v374
									v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
									v544 = v423 + v424
									v545 = v375
									v546 = v376
								case 1:
									v418 = v374
									v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+1)))
									v423 = v419<<(uint(int32(8))%32) + v418
									v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
									v544 = v423 + v424
									v545 = v375
									v546 = v376
								case 2:
									v414 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+2)))
									v418 = v414<<(uint(int32(16))%32) + v374
									v419 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+1)))
									v423 = v419<<(uint(int32(8))%32) + v418
									v424 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372))))
									v544 = v423 + v424
									v545 = v375
									v546 = v376
								case 3:
									v411 = v375
									v412 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v544 = v412 + v374
									v545 = v411
									v546 = v376
								case 4:
									v408 = v375
									v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+4)))
									v411 = v408 + v409
									v412 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v544 = v412 + v374
									v545 = v411
									v546 = v376
								case 5:
									v403 = v375
									v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+5)))
									v408 = v404<<(uint(int32(8))%32) + v403
									v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+4)))
									v411 = v408 + v409
									v412 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v544 = v412 + v374
									v545 = v411
									v546 = v376
								case 6:
									v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+6)))
									v403 = v399<<(uint(int32(16))%32) + v375
									v404 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+5)))
									v408 = v404<<(uint(int32(8))%32) + v403
									v409 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+4)))
									v411 = v408 + v409
									v412 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v544 = v412 + v374
									v545 = v411
									v546 = v376
								case 7:
									v394 = v376
									v395 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v397 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
									v544 = v395 + v374
									v545 = v397 + v375
									v546 = v394
								case 8:
									v389 = v376
									v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+8)))
									v394 = v390<<(uint(int32(8))%32) + v389
									v395 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v397 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
									v544 = v395 + v374
									v545 = v397 + v375
									v546 = v394
								case 9:
									v384 = v376
									v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+9)))
									v389 = v385<<(uint(int32(16))%32) + v384
									v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+8)))
									v394 = v390<<(uint(int32(8))%32) + v389
									v395 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v397 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
									v544 = v395 + v374
									v545 = v397 + v375
									v546 = v394
								case 10:
									v380 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+10)))
									v384 = v380<<(uint(int32(24))%32) + v376
									v385 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+9)))
									v389 = v385<<(uint(int32(16))%32) + v384
									v390 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v372)+8)))
									v394 = v390<<(uint(int32(8))%32) + v389
									v395 = *(*int32)(unsafe.Add(mBase, uint32(v372)))
									v397 = *(*int32)(unsafe.Add(mBase, uint32(v372)+4))
									v544 = v395 + v374
									v545 = v397 + v375
									v546 = v394
								default:
									v544 = v374
									v545 = v375
									v546 = v376
								}
							}
							v549 = int32(14)
							v551 = v545 ^ v546 - base.I32_rotl(v545, v549)
							v555 = v551 ^ v544 - base.I32_rotl(v551, int32(11))
							v559 = v555 ^ v545 - base.I32_rotl(v555, int32(25))
							v563 = v559 ^ v551 - base.I32_rotl(v559, int32(16))
							v567 = v563 ^ v555 - base.I32_rotl(v563, int32(4))
							v571 = v567 ^ v559 - base.I32_rotl(v567, v549)
							v1149 = v571 ^ v563 - base.I32_rotl(v571, int32(24))
						} else {
							v576 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v580 = int32(base.Ui32(v576)>>(uint(int32(2))%32)) - int32(4)
							v586 = v580 - int32(1636608432)
							if v24&int32(3) != 0 {
								if base.Ui32(int32(11)) < base.Ui32(v580) {
									v695 = v24
									v696 = v580
									v697 = v586
									v698 = v586
									v699 = v586
									for {
										v701 = *(*int32)(unsafe.Add(mBase, uint32(v695)+4))
										v702 = v701 + v698
										v703 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
										v705 = *(*int32)(unsafe.Add(mBase, uint32(v695)+8))
										v706 = v705 + v699
										v708 = int32(4)
										v710 = v703 + v697 - v706 ^ base.I32_rotl(v706, v708)
										v714 = v702 - v710 ^ base.I32_rotl(v710, int32(6))
										v715 = v706 + v702
										v716 = v710 + v715
										v717 = v714 + v716
										v721 = v715 - v714 ^ base.I32_rotl(v714, int32(8))
										v725 = v716 - v721 ^ base.I32_rotl(v721, int32(16))
										v729 = v717 - v725 ^ base.I32_rotl(v725, int32(19))
										v730 = v721 + v717
										v731 = v725 + v730
										v732 = v729 + v731
										v736 = v730 - v729 ^ base.I32_rotl(v729, v708)
										v737 = int32(12)
										v738 = v695 + v737
										v740 = v696 - v737
										if base.Ui32(int32(11)) < base.Ui32(v740) {
											v695 = v738
											v696 = v740
											v697 = v731
											v698 = v732
											v699 = v736
											continue
										} else {
											break
										}
										break
									}
									v743 = v738
									v744 = v740
									v745 = v731
									v746 = v732
									v747 = v736
								} else {
									v743 = v24
									v744 = v580
									v745 = v586
									v746 = v586
									v747 = v586
								}
								switch v744 - int32(1) {
								case 0:
									v806 = v745
									v807 = v746
									v808 = v747
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 1:
									v799 = v745
									v800 = v746
									v801 = v747
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 2:
									v792 = v745
									v793 = v746
									v794 = v747
									v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+2)))
									v799 = v795<<(uint(int32(16))%32) + v792
									v800 = v793
									v801 = v794
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 3:
									v786 = v746
									v787 = v747
									v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+3)))
									v792 = v788<<(uint(int32(24))%32) + v745
									v793 = v786
									v794 = v787
									v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+2)))
									v799 = v795<<(uint(int32(16))%32) + v792
									v800 = v793
									v801 = v794
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 4:
									v782 = v746
									v783 = v747
									v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+4)))
									v786 = v782 + v784
									v787 = v783
									v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+3)))
									v792 = v788<<(uint(int32(24))%32) + v745
									v793 = v786
									v794 = v787
									v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+2)))
									v799 = v795<<(uint(int32(16))%32) + v792
									v800 = v793
									v801 = v794
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 5:
									v776 = v746
									v777 = v747
									v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+5)))
									v782 = v778<<(uint(int32(8))%32) + v776
									v783 = v777
									v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+4)))
									v786 = v782 + v784
									v787 = v783
									v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+3)))
									v792 = v788<<(uint(int32(24))%32) + v745
									v793 = v786
									v794 = v787
									v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+2)))
									v799 = v795<<(uint(int32(16))%32) + v792
									v800 = v793
									v801 = v794
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 6:
									v770 = v746
									v771 = v747
									v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+6)))
									v776 = v772<<(uint(int32(16))%32) + v770
									v777 = v771
									v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+5)))
									v782 = v778<<(uint(int32(8))%32) + v776
									v783 = v777
									v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+4)))
									v786 = v782 + v784
									v787 = v783
									v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+3)))
									v792 = v788<<(uint(int32(24))%32) + v745
									v793 = v786
									v794 = v787
									v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+2)))
									v799 = v795<<(uint(int32(16))%32) + v792
									v800 = v793
									v801 = v794
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 7:
									v765 = v747
									v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+7)))
									v770 = v766<<(uint(int32(24))%32) + v746
									v771 = v765
									v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+6)))
									v776 = v772<<(uint(int32(16))%32) + v770
									v777 = v771
									v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+5)))
									v782 = v778<<(uint(int32(8))%32) + v776
									v783 = v777
									v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+4)))
									v786 = v782 + v784
									v787 = v783
									v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+3)))
									v792 = v788<<(uint(int32(24))%32) + v745
									v793 = v786
									v794 = v787
									v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+2)))
									v799 = v795<<(uint(int32(16))%32) + v792
									v800 = v793
									v801 = v794
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 8:
									v760 = v747
									v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+8)))
									v765 = v761<<(uint(int32(8))%32) + v760
									v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+7)))
									v770 = v766<<(uint(int32(24))%32) + v746
									v771 = v765
									v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+6)))
									v776 = v772<<(uint(int32(16))%32) + v770
									v777 = v771
									v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+5)))
									v782 = v778<<(uint(int32(8))%32) + v776
									v783 = v777
									v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+4)))
									v786 = v782 + v784
									v787 = v783
									v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+3)))
									v792 = v788<<(uint(int32(24))%32) + v745
									v793 = v786
									v794 = v787
									v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+2)))
									v799 = v795<<(uint(int32(16))%32) + v792
									v800 = v793
									v801 = v794
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 9:
									v755 = v747
									v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+9)))
									v760 = v756<<(uint(int32(16))%32) + v755
									v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+8)))
									v765 = v761<<(uint(int32(8))%32) + v760
									v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+7)))
									v770 = v766<<(uint(int32(24))%32) + v746
									v771 = v765
									v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+6)))
									v776 = v772<<(uint(int32(16))%32) + v770
									v777 = v771
									v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+5)))
									v782 = v778<<(uint(int32(8))%32) + v776
									v783 = v777
									v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+4)))
									v786 = v782 + v784
									v787 = v783
									v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+3)))
									v792 = v788<<(uint(int32(24))%32) + v745
									v793 = v786
									v794 = v787
									v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+2)))
									v799 = v795<<(uint(int32(16))%32) + v792
									v800 = v793
									v801 = v794
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								case 10:
									v751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+10)))
									v755 = v751<<(uint(int32(24))%32) + v747
									v756 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+9)))
									v760 = v756<<(uint(int32(16))%32) + v755
									v761 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+8)))
									v765 = v761<<(uint(int32(8))%32) + v760
									v766 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+7)))
									v770 = v766<<(uint(int32(24))%32) + v746
									v771 = v765
									v772 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+6)))
									v776 = v772<<(uint(int32(16))%32) + v770
									v777 = v771
									v778 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+5)))
									v782 = v778<<(uint(int32(8))%32) + v776
									v783 = v777
									v784 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+4)))
									v786 = v782 + v784
									v787 = v783
									v788 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+3)))
									v792 = v788<<(uint(int32(24))%32) + v745
									v793 = v786
									v794 = v787
									v795 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+2)))
									v799 = v795<<(uint(int32(16))%32) + v792
									v800 = v793
									v801 = v794
									v802 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743)+1)))
									v806 = v802<<(uint(int32(8))%32) + v799
									v807 = v800
									v808 = v801
									v809 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v743))))
									v813 = v806 + v809
									v814 = v807
									v815 = v808
								default:
									v813 = v745
									v814 = v746
									v815 = v747
								}
							} else {
								if base.Ui32(v580) < base.Ui32(int32(12)) {
									v641 = v24
									v642 = v580
									v643 = v586
									v644 = v586
									v645 = v586
								} else {
									v593 = v24
									v594 = v580
									v595 = v586
									v596 = v586
									v597 = v586
									for {
										v599 = *(*int32)(unsafe.Add(mBase, uint32(v593)+4))
										v600 = v599 + v596
										v601 = *(*int32)(unsafe.Add(mBase, uint32(v593)))
										v603 = *(*int32)(unsafe.Add(mBase, uint32(v593)+8))
										v604 = v603 + v597
										v606 = int32(4)
										v608 = v601 + v595 - v604 ^ base.I32_rotl(v604, v606)
										v612 = v600 - v608 ^ base.I32_rotl(v608, int32(6))
										v613 = v604 + v600
										v614 = v608 + v613
										v615 = v612 + v614
										v619 = v613 - v612 ^ base.I32_rotl(v612, int32(8))
										v623 = v614 - v619 ^ base.I32_rotl(v619, int32(16))
										v627 = v615 - v623 ^ base.I32_rotl(v623, int32(19))
										v628 = v619 + v615
										v629 = v623 + v628
										v630 = v627 + v629
										v634 = v628 - v627 ^ base.I32_rotl(v627, v606)
										v635 = int32(12)
										v636 = v593 + v635
										v638 = v594 - v635
										if base.Ui32(int32(11)) < base.Ui32(v638) {
											v593 = v636
											v594 = v638
											v595 = v629
											v596 = v630
											v597 = v634
											continue
										} else {
											break
										}
										break
									}
									v641 = v636
									v642 = v638
									v643 = v629
									v644 = v630
									v645 = v634
								}
								switch v642 - int32(1) {
								case 0:
									v692 = v643
									v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641))))
									v813 = v692 + v693
									v814 = v644
									v815 = v645
								case 1:
									v687 = v643
									v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+1)))
									v692 = v688<<(uint(int32(8))%32) + v687
									v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641))))
									v813 = v692 + v693
									v814 = v644
									v815 = v645
								case 2:
									v683 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+2)))
									v687 = v683<<(uint(int32(16))%32) + v643
									v688 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+1)))
									v692 = v688<<(uint(int32(8))%32) + v687
									v693 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641))))
									v813 = v692 + v693
									v814 = v644
									v815 = v645
								case 3:
									v680 = v644
									v681 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
									v813 = v681 + v643
									v814 = v680
									v815 = v645
								case 4:
									v677 = v644
									v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+4)))
									v680 = v677 + v678
									v681 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
									v813 = v681 + v643
									v814 = v680
									v815 = v645
								case 5:
									v672 = v644
									v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+5)))
									v677 = v673<<(uint(int32(8))%32) + v672
									v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+4)))
									v680 = v677 + v678
									v681 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
									v813 = v681 + v643
									v814 = v680
									v815 = v645
								case 6:
									v668 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+6)))
									v672 = v668<<(uint(int32(16))%32) + v644
									v673 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+5)))
									v677 = v673<<(uint(int32(8))%32) + v672
									v678 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+4)))
									v680 = v677 + v678
									v681 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
									v813 = v681 + v643
									v814 = v680
									v815 = v645
								case 7:
									v663 = v645
									v664 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
									v666 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
									v813 = v664 + v643
									v814 = v666 + v644
									v815 = v663
								case 8:
									v658 = v645
									v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+8)))
									v663 = v659<<(uint(int32(8))%32) + v658
									v664 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
									v666 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
									v813 = v664 + v643
									v814 = v666 + v644
									v815 = v663
								case 9:
									v653 = v645
									v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+9)))
									v658 = v654<<(uint(int32(16))%32) + v653
									v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+8)))
									v663 = v659<<(uint(int32(8))%32) + v658
									v664 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
									v666 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
									v813 = v664 + v643
									v814 = v666 + v644
									v815 = v663
								case 10:
									v649 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+10)))
									v653 = v649<<(uint(int32(24))%32) + v645
									v654 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+9)))
									v658 = v654<<(uint(int32(16))%32) + v653
									v659 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v641)+8)))
									v663 = v659<<(uint(int32(8))%32) + v658
									v664 = *(*int32)(unsafe.Add(mBase, uint32(v641)))
									v666 = *(*int32)(unsafe.Add(mBase, uint32(v641)+4))
									v813 = v664 + v643
									v814 = v666 + v644
									v815 = v663
								default:
									v813 = v643
									v814 = v644
									v815 = v645
								}
							}
							v818 = int32(14)
							v820 = v814 ^ v815 - base.I32_rotl(v814, v818)
							v824 = v820 ^ v813 - base.I32_rotl(v820, int32(11))
							v828 = v824 ^ v814 - base.I32_rotl(v824, int32(25))
							v832 = v828 ^ v820 - base.I32_rotl(v828, int32(16))
							v836 = v832 ^ v824 - base.I32_rotl(v832, int32(4))
							v840 = v836 ^ v828 - base.I32_rotl(v836, v818)
							v1149 = v840 ^ v832 - base.I32_rotl(v840, int32(24))
						}
					}
					v1154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v1154 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v1157 = m.ExcPending
						if v1157 != 0 {
							return int32(0)
						} else {
							return v1149
						}
					} else {
						return v1149
					}
				} else {
					v845 = int32(0)
					if v21 == int32(1) {
						v852 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
						if v852 == int32(18) {
							v855 = int32(16)
						} else {
							v855 = int32(0)
						}
						if base.Ui32((v852-int32(1))&int32(255)) < base.Ui32(int32(3)) {
							v862 = int32(4)
						} else {
							v862 = v855
						}
						v873 = v862
					} else {
						v863 = int32(1)
						if v23 != 0 {
							v873 = int32(base.Ui32(v21)>>(uint(v863)%32)) - v863
						} else {
							v867 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
							v873 = int32(base.Ui32(v867)>>(uint(int32(2))%32)) - int32(4)
						}
					}
					v874 = F_pg_strnxfrm(m, v845, v845, v24, v873, v15)
					mBase = m.M
					v875 = m.ExcPending
					if v875 != 0 {
						return int32(0)
					} else {
						v877 = v874 + int32(1)
						v878 = F_palloc(m, v877)
						mBase = m.M
						v879 = m.ExcPending
						if v879 != 0 {
							return int32(0)
						} else {
							v880 = F_pg_strnxfrm(m, v878, v877, v24, v873, v15)
							mBase = m.M
							v881 = m.ExcPending
							if v881 != 0 {
								return int32(0)
							} else {
								if base.Ui32(v874) < base.Ui32(v880) {
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v1182 = m.ExcPending
									if v1182 != 0 {
										return int32(0)
									} else {
										F_errmsg_internal(m, int32(_a_F_hashtext_0), int32(0))
										mBase = m.M
										v1186 = m.ExcPending
										if v1186 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_hashtext_1), int32(306), int32(_a_F_hashtext_2))
											mBase = m.M
											v1191 = m.ExcPending
											if v1191 != 0 {
												return int32(0)
											} else {
												base.Wasm_trap_unreachable()
												for {
												}
											}
										}
									}
								} else {
									v888 = v877 - int32(1636608432)
									if v878&int32(3) != 0 {
										if base.Ui32(int32(11)) < base.Ui32(v877) {
											v997 = v878
											v998 = v877
											v999 = v888
											v1000 = v888
											v1001 = v888
											for {
												v1003 = *(*int32)(unsafe.Add(mBase, uint32(v997)+4))
												v1004 = v1003 + v1000
												v1005 = *(*int32)(unsafe.Add(mBase, uint32(v997)))
												v1007 = *(*int32)(unsafe.Add(mBase, uint32(v997)+8))
												v1008 = v1007 + v1001
												v1010 = int32(4)
												v1012 = v1005 + v999 - v1008 ^ base.I32_rotl(v1008, v1010)
												v1016 = v1004 - v1012 ^ base.I32_rotl(v1012, int32(6))
												v1017 = v1008 + v1004
												v1018 = v1012 + v1017
												v1019 = v1016 + v1018
												v1023 = v1017 - v1016 ^ base.I32_rotl(v1016, int32(8))
												v1027 = v1018 - v1023 ^ base.I32_rotl(v1023, int32(16))
												v1031 = v1019 - v1027 ^ base.I32_rotl(v1027, int32(19))
												v1032 = v1023 + v1019
												v1033 = v1027 + v1032
												v1034 = v1031 + v1033
												v1038 = v1032 - v1031 ^ base.I32_rotl(v1031, v1010)
												v1039 = int32(12)
												v1040 = v997 + v1039
												v1042 = v998 - v1039
												if base.Ui32(int32(11)) < base.Ui32(v1042) {
													v997 = v1040
													v998 = v1042
													v999 = v1033
													v1000 = v1034
													v1001 = v1038
													continue
												} else {
													break
												}
												break
											}
											v1045 = v1040
											v1046 = v1042
											v1047 = v1033
											v1048 = v1034
											v1049 = v1038
										} else {
											v1045 = v878
											v1046 = v877
											v1047 = v888
											v1048 = v888
											v1049 = v888
										}
										switch v1046 - int32(1) {
										case 0:
											v1108 = v1047
											v1109 = v1048
											v1110 = v1049
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 1:
											v1101 = v1047
											v1102 = v1048
											v1103 = v1049
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 2:
											v1094 = v1047
											v1095 = v1048
											v1096 = v1049
											v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+2)))
											v1101 = v1097<<(uint(int32(16))%32) + v1094
											v1102 = v1095
											v1103 = v1096
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 3:
											v1088 = v1048
											v1089 = v1049
											v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+3)))
											v1094 = v1090<<(uint(int32(24))%32) + v1047
											v1095 = v1088
											v1096 = v1089
											v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+2)))
											v1101 = v1097<<(uint(int32(16))%32) + v1094
											v1102 = v1095
											v1103 = v1096
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 4:
											v1084 = v1048
											v1085 = v1049
											v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
											v1088 = v1084 + v1086
											v1089 = v1085
											v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+3)))
											v1094 = v1090<<(uint(int32(24))%32) + v1047
											v1095 = v1088
											v1096 = v1089
											v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+2)))
											v1101 = v1097<<(uint(int32(16))%32) + v1094
											v1102 = v1095
											v1103 = v1096
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 5:
											v1078 = v1048
											v1079 = v1049
											v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+5)))
											v1084 = v1080<<(uint(int32(8))%32) + v1078
											v1085 = v1079
											v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
											v1088 = v1084 + v1086
											v1089 = v1085
											v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+3)))
											v1094 = v1090<<(uint(int32(24))%32) + v1047
											v1095 = v1088
											v1096 = v1089
											v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+2)))
											v1101 = v1097<<(uint(int32(16))%32) + v1094
											v1102 = v1095
											v1103 = v1096
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 6:
											v1072 = v1048
											v1073 = v1049
											v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+6)))
											v1078 = v1074<<(uint(int32(16))%32) + v1072
											v1079 = v1073
											v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+5)))
											v1084 = v1080<<(uint(int32(8))%32) + v1078
											v1085 = v1079
											v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
											v1088 = v1084 + v1086
											v1089 = v1085
											v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+3)))
											v1094 = v1090<<(uint(int32(24))%32) + v1047
											v1095 = v1088
											v1096 = v1089
											v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+2)))
											v1101 = v1097<<(uint(int32(16))%32) + v1094
											v1102 = v1095
											v1103 = v1096
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 7:
											v1067 = v1049
											v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+7)))
											v1072 = v1068<<(uint(int32(24))%32) + v1048
											v1073 = v1067
											v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+6)))
											v1078 = v1074<<(uint(int32(16))%32) + v1072
											v1079 = v1073
											v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+5)))
											v1084 = v1080<<(uint(int32(8))%32) + v1078
											v1085 = v1079
											v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
											v1088 = v1084 + v1086
											v1089 = v1085
											v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+3)))
											v1094 = v1090<<(uint(int32(24))%32) + v1047
											v1095 = v1088
											v1096 = v1089
											v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+2)))
											v1101 = v1097<<(uint(int32(16))%32) + v1094
											v1102 = v1095
											v1103 = v1096
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 8:
											v1062 = v1049
											v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+8)))
											v1067 = v1063<<(uint(int32(8))%32) + v1062
											v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+7)))
											v1072 = v1068<<(uint(int32(24))%32) + v1048
											v1073 = v1067
											v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+6)))
											v1078 = v1074<<(uint(int32(16))%32) + v1072
											v1079 = v1073
											v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+5)))
											v1084 = v1080<<(uint(int32(8))%32) + v1078
											v1085 = v1079
											v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
											v1088 = v1084 + v1086
											v1089 = v1085
											v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+3)))
											v1094 = v1090<<(uint(int32(24))%32) + v1047
											v1095 = v1088
											v1096 = v1089
											v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+2)))
											v1101 = v1097<<(uint(int32(16))%32) + v1094
											v1102 = v1095
											v1103 = v1096
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 9:
											v1057 = v1049
											v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+9)))
											v1062 = v1058<<(uint(int32(16))%32) + v1057
											v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+8)))
											v1067 = v1063<<(uint(int32(8))%32) + v1062
											v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+7)))
											v1072 = v1068<<(uint(int32(24))%32) + v1048
											v1073 = v1067
											v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+6)))
											v1078 = v1074<<(uint(int32(16))%32) + v1072
											v1079 = v1073
											v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+5)))
											v1084 = v1080<<(uint(int32(8))%32) + v1078
											v1085 = v1079
											v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
											v1088 = v1084 + v1086
											v1089 = v1085
											v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+3)))
											v1094 = v1090<<(uint(int32(24))%32) + v1047
											v1095 = v1088
											v1096 = v1089
											v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+2)))
											v1101 = v1097<<(uint(int32(16))%32) + v1094
											v1102 = v1095
											v1103 = v1096
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										case 10:
											v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+10)))
											v1057 = v1053<<(uint(int32(24))%32) + v1049
											v1058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+9)))
											v1062 = v1058<<(uint(int32(16))%32) + v1057
											v1063 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+8)))
											v1067 = v1063<<(uint(int32(8))%32) + v1062
											v1068 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+7)))
											v1072 = v1068<<(uint(int32(24))%32) + v1048
											v1073 = v1067
											v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+6)))
											v1078 = v1074<<(uint(int32(16))%32) + v1072
											v1079 = v1073
											v1080 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+5)))
											v1084 = v1080<<(uint(int32(8))%32) + v1078
											v1085 = v1079
											v1086 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+4)))
											v1088 = v1084 + v1086
											v1089 = v1085
											v1090 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+3)))
											v1094 = v1090<<(uint(int32(24))%32) + v1047
											v1095 = v1088
											v1096 = v1089
											v1097 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+2)))
											v1101 = v1097<<(uint(int32(16))%32) + v1094
											v1102 = v1095
											v1103 = v1096
											v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045)+1)))
											v1108 = v1104<<(uint(int32(8))%32) + v1101
											v1109 = v1102
											v1110 = v1103
											v1111 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1045))))
											v1115 = v1108 + v1111
											v1116 = v1109
											v1117 = v1110
										default:
											v1115 = v1047
											v1116 = v1048
											v1117 = v1049
										}
									} else {
										if base.Ui32(v877) < base.Ui32(int32(12)) {
											v943 = v878
											v944 = v877
											v945 = v888
											v946 = v888
											v947 = v888
										} else {
											v895 = v878
											v896 = v877
											v897 = v888
											v898 = v888
											v899 = v888
											for {
												v901 = *(*int32)(unsafe.Add(mBase, uint32(v895)+4))
												v902 = v901 + v898
												v903 = *(*int32)(unsafe.Add(mBase, uint32(v895)))
												v905 = *(*int32)(unsafe.Add(mBase, uint32(v895)+8))
												v906 = v905 + v899
												v908 = int32(4)
												v910 = v903 + v897 - v906 ^ base.I32_rotl(v906, v908)
												v914 = v902 - v910 ^ base.I32_rotl(v910, int32(6))
												v915 = v906 + v902
												v916 = v910 + v915
												v917 = v914 + v916
												v921 = v915 - v914 ^ base.I32_rotl(v914, int32(8))
												v925 = v916 - v921 ^ base.I32_rotl(v921, int32(16))
												v929 = v917 - v925 ^ base.I32_rotl(v925, int32(19))
												v930 = v921 + v917
												v931 = v925 + v930
												v932 = v929 + v931
												v936 = v930 - v929 ^ base.I32_rotl(v929, v908)
												v937 = int32(12)
												v938 = v895 + v937
												v940 = v896 - v937
												if base.Ui32(int32(11)) < base.Ui32(v940) {
													v895 = v938
													v896 = v940
													v897 = v931
													v898 = v932
													v899 = v936
													continue
												} else {
													break
												}
												break
											}
											v943 = v938
											v944 = v940
											v945 = v931
											v946 = v932
											v947 = v936
										}
										switch v944 - int32(1) {
										case 0:
											v994 = v945
											v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943))))
											v1115 = v994 + v995
											v1116 = v946
											v1117 = v947
										case 1:
											v989 = v945
											v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+1)))
											v994 = v990<<(uint(int32(8))%32) + v989
											v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943))))
											v1115 = v994 + v995
											v1116 = v946
											v1117 = v947
										case 2:
											v985 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+2)))
											v989 = v985<<(uint(int32(16))%32) + v945
											v990 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+1)))
											v994 = v990<<(uint(int32(8))%32) + v989
											v995 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943))))
											v1115 = v994 + v995
											v1116 = v946
											v1117 = v947
										case 3:
											v982 = v946
											v983 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
											v1115 = v983 + v945
											v1116 = v982
											v1117 = v947
										case 4:
											v979 = v946
											v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+4)))
											v982 = v979 + v980
											v983 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
											v1115 = v983 + v945
											v1116 = v982
											v1117 = v947
										case 5:
											v974 = v946
											v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+5)))
											v979 = v975<<(uint(int32(8))%32) + v974
											v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+4)))
											v982 = v979 + v980
											v983 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
											v1115 = v983 + v945
											v1116 = v982
											v1117 = v947
										case 6:
											v970 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+6)))
											v974 = v970<<(uint(int32(16))%32) + v946
											v975 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+5)))
											v979 = v975<<(uint(int32(8))%32) + v974
											v980 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+4)))
											v982 = v979 + v980
											v983 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
											v1115 = v983 + v945
											v1116 = v982
											v1117 = v947
										case 7:
											v965 = v947
											v966 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
											v968 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
											v1115 = v966 + v945
											v1116 = v968 + v946
											v1117 = v965
										case 8:
											v960 = v947
											v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+8)))
											v965 = v961<<(uint(int32(8))%32) + v960
											v966 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
											v968 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
											v1115 = v966 + v945
											v1116 = v968 + v946
											v1117 = v965
										case 9:
											v955 = v947
											v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+9)))
											v960 = v956<<(uint(int32(16))%32) + v955
											v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+8)))
											v965 = v961<<(uint(int32(8))%32) + v960
											v966 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
											v968 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
											v1115 = v966 + v945
											v1116 = v968 + v946
											v1117 = v965
										case 10:
											v951 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+10)))
											v955 = v951<<(uint(int32(24))%32) + v947
											v956 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+9)))
											v960 = v956<<(uint(int32(16))%32) + v955
											v961 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v943)+8)))
											v965 = v961<<(uint(int32(8))%32) + v960
											v966 = *(*int32)(unsafe.Add(mBase, uint32(v943)))
											v968 = *(*int32)(unsafe.Add(mBase, uint32(v943)+4))
											v1115 = v966 + v945
											v1116 = v968 + v946
											v1117 = v965
										default:
											v1115 = v945
											v1116 = v946
											v1117 = v947
										}
									}
									v1120 = int32(14)
									v1122 = v1116 ^ v1117 - base.I32_rotl(v1116, v1120)
									v1126 = v1122 ^ v1115 - base.I32_rotl(v1122, int32(11))
									v1130 = v1126 ^ v1116 - base.I32_rotl(v1126, int32(25))
									v1134 = v1130 ^ v1122 - base.I32_rotl(v1130, int32(16))
									v1138 = v1134 ^ v1126 - base.I32_rotl(v1134, int32(4))
									v1142 = v1138 ^ v1130 - base.I32_rotl(v1138, v1120)
									F_pfree(m, v878)
									mBase = m.M
									v1148 = m.ExcPending
									if v1148 != 0 {
										return int32(0)
									} else {
										v1149 = v1142 ^ v1134 - base.I32_rotl(v1142, int32(24))
										v1154 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
										if v1154 != v10 {
											F_pfree(m, v10)
											mBase = m.M
											v1157 = m.ExcPending
											if v1157 != 0 {
												return int32(0)
											} else {
												return v1149
											}
										} else {
											return v1149
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
			v1162 = m.ExcPending
			if v1162 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(34209924))
				mBase = m.M
				v1165 = m.ExcPending
				if v1165 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_hashtext_3), int32(0))
					mBase = m.M
					v1169 = m.ExcPending
					if v1169 != 0 {
						return int32(0)
					} else {
						F_errhint(m, int32(_a_F_hashtext_4), int32(0))
						mBase = m.M
						v1173 = m.ExcPending
						if v1173 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_hashtext_1), int32(281), int32(_a_F_hashtext_2))
							mBase = m.M
							v1178 = m.ExcPending
							if v1178 != 0 {
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
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
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
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
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
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
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
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
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
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
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
	var v260 int32
	_ = v260
	var v264 int32
	_ = v264
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
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
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
			v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12))))
			if v24 == int32(18) {
				v27 = int32(16)
			} else {
				v27 = int32(0)
			}
			if base.Ui32((v24-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v34 = int32(4)
			} else {
				v34 = v27
			}
			v45 = v34
		} else {
			v35 = int32(1)
			if v17 != 0 {
				v45 = int32(base.Ui32(v15)>>(uint(v35)%32)) - v35
			} else {
				v39 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				v45 = int32(base.Ui32(v39)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v51 = v45 - int32(1636608432)
		if v18&int32(3) != 0 {
			if base.Ui32(int32(11)) < base.Ui32(v45) {
				v160 = v18
				v161 = v45
				v162 = v51
				v163 = v51
				v164 = v51
				for {
					v166 = *(*int32)(unsafe.Add(mBase, uint32(v160)+4))
					v167 = v166 + v163
					v168 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
					v170 = *(*int32)(unsafe.Add(mBase, uint32(v160)+8))
					v171 = v170 + v164
					v173 = int32(4)
					v175 = v168 + v162 - v171 ^ base.I32_rotl(v171, v173)
					v179 = v167 - v175 ^ base.I32_rotl(v175, int32(6))
					v180 = v171 + v167
					v181 = v175 + v180
					v182 = v179 + v181
					v186 = v180 - v179 ^ base.I32_rotl(v179, int32(8))
					v190 = v181 - v186 ^ base.I32_rotl(v186, int32(16))
					v194 = v182 - v190 ^ base.I32_rotl(v190, int32(19))
					v195 = v186 + v182
					v196 = v190 + v195
					v197 = v194 + v196
					v201 = v195 - v194 ^ base.I32_rotl(v194, v173)
					v202 = int32(12)
					v203 = v160 + v202
					v205 = v161 - v202
					if base.Ui32(int32(11)) < base.Ui32(v205) {
						v160 = v203
						v161 = v205
						v162 = v196
						v163 = v197
						v164 = v201
						continue
					} else {
						break
					}
					break
				}
				v208 = v203
				v209 = v205
				v210 = v196
				v211 = v197
				v212 = v201
			} else {
				v208 = v18
				v209 = v45
				v210 = v51
				v211 = v51
				v212 = v51
			}
			switch v209 - int32(1) {
			case 0:
				v271 = v210
				v272 = v211
				v273 = v212
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 1:
				v264 = v210
				v265 = v211
				v266 = v212
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 2:
				v257 = v210
				v258 = v211
				v259 = v212
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v264 = v260<<(uint(int32(16))%32) + v257
				v265 = v258
				v266 = v259
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 3:
				v251 = v211
				v252 = v212
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+3)))
				v257 = v253<<(uint(int32(24))%32) + v210
				v258 = v251
				v259 = v252
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v264 = v260<<(uint(int32(16))%32) + v257
				v265 = v258
				v266 = v259
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 4:
				v247 = v211
				v248 = v212
				v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v251 = v247 + v249
				v252 = v248
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+3)))
				v257 = v253<<(uint(int32(24))%32) + v210
				v258 = v251
				v259 = v252
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v264 = v260<<(uint(int32(16))%32) + v257
				v265 = v258
				v266 = v259
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 5:
				v241 = v211
				v242 = v212
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
				v247 = v243<<(uint(int32(8))%32) + v241
				v248 = v242
				v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v251 = v247 + v249
				v252 = v248
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+3)))
				v257 = v253<<(uint(int32(24))%32) + v210
				v258 = v251
				v259 = v252
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v264 = v260<<(uint(int32(16))%32) + v257
				v265 = v258
				v266 = v259
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 6:
				v235 = v211
				v236 = v212
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
				v241 = v237<<(uint(int32(16))%32) + v235
				v242 = v236
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
				v247 = v243<<(uint(int32(8))%32) + v241
				v248 = v242
				v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v251 = v247 + v249
				v252 = v248
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+3)))
				v257 = v253<<(uint(int32(24))%32) + v210
				v258 = v251
				v259 = v252
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v264 = v260<<(uint(int32(16))%32) + v257
				v265 = v258
				v266 = v259
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 7:
				v230 = v212
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+7)))
				v235 = v231<<(uint(int32(24))%32) + v211
				v236 = v230
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
				v241 = v237<<(uint(int32(16))%32) + v235
				v242 = v236
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
				v247 = v243<<(uint(int32(8))%32) + v241
				v248 = v242
				v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v251 = v247 + v249
				v252 = v248
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+3)))
				v257 = v253<<(uint(int32(24))%32) + v210
				v258 = v251
				v259 = v252
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v264 = v260<<(uint(int32(16))%32) + v257
				v265 = v258
				v266 = v259
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 8:
				v225 = v212
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
				v230 = v226<<(uint(int32(8))%32) + v225
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+7)))
				v235 = v231<<(uint(int32(24))%32) + v211
				v236 = v230
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
				v241 = v237<<(uint(int32(16))%32) + v235
				v242 = v236
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
				v247 = v243<<(uint(int32(8))%32) + v241
				v248 = v242
				v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v251 = v247 + v249
				v252 = v248
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+3)))
				v257 = v253<<(uint(int32(24))%32) + v210
				v258 = v251
				v259 = v252
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v264 = v260<<(uint(int32(16))%32) + v257
				v265 = v258
				v266 = v259
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 9:
				v220 = v212
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
				v225 = v221<<(uint(int32(16))%32) + v220
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
				v230 = v226<<(uint(int32(8))%32) + v225
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+7)))
				v235 = v231<<(uint(int32(24))%32) + v211
				v236 = v230
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
				v241 = v237<<(uint(int32(16))%32) + v235
				v242 = v236
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
				v247 = v243<<(uint(int32(8))%32) + v241
				v248 = v242
				v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v251 = v247 + v249
				v252 = v248
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+3)))
				v257 = v253<<(uint(int32(24))%32) + v210
				v258 = v251
				v259 = v252
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v264 = v260<<(uint(int32(16))%32) + v257
				v265 = v258
				v266 = v259
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			case 10:
				v216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+10)))
				v220 = v216<<(uint(int32(24))%32) + v212
				v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+9)))
				v225 = v221<<(uint(int32(16))%32) + v220
				v226 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+8)))
				v230 = v226<<(uint(int32(8))%32) + v225
				v231 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+7)))
				v235 = v231<<(uint(int32(24))%32) + v211
				v236 = v230
				v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+6)))
				v241 = v237<<(uint(int32(16))%32) + v235
				v242 = v236
				v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+5)))
				v247 = v243<<(uint(int32(8))%32) + v241
				v248 = v242
				v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+4)))
				v251 = v247 + v249
				v252 = v248
				v253 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+3)))
				v257 = v253<<(uint(int32(24))%32) + v210
				v258 = v251
				v259 = v252
				v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+2)))
				v264 = v260<<(uint(int32(16))%32) + v257
				v265 = v258
				v266 = v259
				v267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208)+1)))
				v271 = v267<<(uint(int32(8))%32) + v264
				v272 = v265
				v273 = v266
				v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v208))))
				v278 = v271 + v274
				v279 = v272
				v280 = v273
			default:
				v278 = v210
				v279 = v211
				v280 = v212
			}
		} else {
			if base.Ui32(v45) < base.Ui32(int32(12)) {
				v106 = v18
				v107 = v45
				v108 = v51
				v109 = v51
				v110 = v51
			} else {
				v58 = v18
				v59 = v45
				v60 = v51
				v61 = v51
				v62 = v51
				for {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v58)+4))
					v65 = v64 + v61
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v58)))
					v68 = *(*int32)(unsafe.Add(mBase, uint32(v58)+8))
					v69 = v68 + v62
					v71 = int32(4)
					v73 = v66 + v60 - v69 ^ base.I32_rotl(v69, v71)
					v77 = v65 - v73 ^ base.I32_rotl(v73, int32(6))
					v78 = v69 + v65
					v79 = v73 + v78
					v80 = v77 + v79
					v84 = v78 - v77 ^ base.I32_rotl(v77, int32(8))
					v88 = v79 - v84 ^ base.I32_rotl(v84, int32(16))
					v92 = v80 - v88 ^ base.I32_rotl(v88, int32(19))
					v93 = v84 + v80
					v94 = v88 + v93
					v95 = v92 + v94
					v99 = v93 - v92 ^ base.I32_rotl(v92, v71)
					v100 = int32(12)
					v101 = v58 + v100
					v103 = v59 - v100
					if base.Ui32(int32(11)) < base.Ui32(v103) {
						v58 = v101
						v59 = v103
						v60 = v94
						v61 = v95
						v62 = v99
						continue
					} else {
						break
					}
					break
				}
				v106 = v101
				v107 = v103
				v108 = v94
				v109 = v95
				v110 = v99
			}
			switch v107 - int32(1) {
			case 0:
				v157 = v108
				v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
				v278 = v157 + v158
				v279 = v109
				v280 = v110
			case 1:
				v152 = v108
				v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
				v157 = v153<<(uint(int32(8))%32) + v152
				v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
				v278 = v157 + v158
				v279 = v109
				v280 = v110
			case 2:
				v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+2)))
				v152 = v148<<(uint(int32(16))%32) + v108
				v153 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+1)))
				v157 = v153<<(uint(int32(8))%32) + v152
				v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106))))
				v278 = v157 + v158
				v279 = v109
				v280 = v110
			case 3:
				v145 = v109
				v146 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v278 = v146 + v108
				v279 = v145
				v280 = v110
			case 4:
				v142 = v109
				v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
				v145 = v142 + v143
				v146 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v278 = v146 + v108
				v279 = v145
				v280 = v110
			case 5:
				v137 = v109
				v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+5)))
				v142 = v138<<(uint(int32(8))%32) + v137
				v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
				v145 = v142 + v143
				v146 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v278 = v146 + v108
				v279 = v145
				v280 = v110
			case 6:
				v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+6)))
				v137 = v133<<(uint(int32(16))%32) + v109
				v138 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+5)))
				v142 = v138<<(uint(int32(8))%32) + v137
				v143 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+4)))
				v145 = v142 + v143
				v146 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v278 = v146 + v108
				v279 = v145
				v280 = v110
			case 7:
				v128 = v110
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
				v278 = v129 + v108
				v279 = v131 + v109
				v280 = v128
			case 8:
				v123 = v110
				v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+8)))
				v128 = v124<<(uint(int32(8))%32) + v123
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
				v278 = v129 + v108
				v279 = v131 + v109
				v280 = v128
			case 9:
				v118 = v110
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+9)))
				v123 = v119<<(uint(int32(16))%32) + v118
				v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+8)))
				v128 = v124<<(uint(int32(8))%32) + v123
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
				v278 = v129 + v108
				v279 = v131 + v109
				v280 = v128
			case 10:
				v114 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+10)))
				v118 = v114<<(uint(int32(24))%32) + v110
				v119 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+9)))
				v123 = v119<<(uint(int32(16))%32) + v118
				v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v106)+8)))
				v128 = v124<<(uint(int32(8))%32) + v123
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v106)))
				v131 = *(*int32)(unsafe.Add(mBase, uint32(v106)+4))
				v278 = v129 + v108
				v279 = v131 + v109
				v280 = v128
			default:
				v278 = v108
				v279 = v109
				v280 = v110
			}
		}
		v283 = int32(14)
		v285 = v279 ^ v280 - base.I32_rotl(v279, v283)
		v289 = v285 ^ v278 - base.I32_rotl(v285, int32(11))
		v293 = v289 ^ v279 - base.I32_rotl(v289, int32(25))
		v297 = v293 ^ v285 - base.I32_rotl(v293, int32(16))
		v301 = v297 ^ v289 - base.I32_rotl(v297, int32(4))
		v305 = v301 ^ v293 - base.I32_rotl(v301, v283)
		v310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v310 != v7 {
			F_pfree(m, v7)
			mBase = m.M
			v313 = m.ExcPending
			if v313 != 0 {
				return int32(0)
			} else {
				return v305 ^ v297 - base.I32_rotl(v305, int32(24))
			}
		} else {
			return v305 ^ v297 - base.I32_rotl(v305, int32(24))
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
			v12 = *(*int32)(unsafe.Add(mBase, _c_F_have_createdb_privilege[0]))
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
	v286 = F_TS_execute(m, l1+int32(8), v20+int32(8), int32(0), int32(1175))
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
	F_LockBuffer(m, v123, int32(2))
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
	F_LockBuffer(m, v755, int32(1))
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
	F_LockBuffer(m, v1037, int32(1))
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
	F_LockBuffer(m, v1124, int32(1))
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
	F_LockBuffer(m, v1466, int32(2))
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
	var v21 int64
	_ = v21
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v61 int32
	_ = v61
	var v63 float64
	_ = v63
	var v64 float64
	_ = v64
	var v65 int32
	_ = v65
	var v75 float64
	_ = v75
	var v84 float64
	_ = v84
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 float64
	_ = v96
	var v97 float64
	_ = v97
	var v98 float64
	_ = v98
	var v103 float64
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 float64
	_ = v107
	var v112 float64
	_ = v112
	var v114 float64
	_ = v114
	var v120 float64
	_ = v120
	var v124 float64
	_ = v124
	var v126 float64
	_ = v126
	var v129 int64
	_ = v129
	var v133 int64
	_ = v133
	v16 = m.G0
	v18 = v16 - int32(80)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+80))
	if v20 != 0 {
		v21 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v18)+64)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v18)+48)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v18)+40)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v18)+32)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v18)+24)) = v21
		*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v21
		F_genericcostestimate(m, l0, l1, l2, v18+int32(16))
		mBase = m.M
		v40 = m.ExcPending
		if v40 != 0 {
			return
		} else {
			v41 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
			v42 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
			v44 = F_index_open(m, v42, int32(0))
			mBase = m.M
			v45 = m.ExcPending
			if v45 != 0 {
				return
			} else {
				F_HnswGetMetaPageInfo(m, v44, v18+int32(12), int32(0))
				mBase = m.M
				v50 = m.ExcPending
				if v50 != 0 {
					return
				} else {
					F_relation_close(m, v44, int32(0))
					mBase = m.M
					v53 = m.ExcPending
					if v53 != 0 {
						return
					} else {
						v55 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
						v56 = *(*float64)(unsafe.Add(mBase, uint32(v55)+24))
						if base.F64_gt(v56, float64(0)) != 0 {
							v59 = F_log(m, v56)
							mBase = m.M
							v61 = *(*int32)(unsafe.Add(mBase, _c_F_hnswcostestimate[0]))
							v63 = F_log(m, base.F64_convert_i32_s(v61))
							mBase = m.M
							v64 = float64(1)
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
							v75 = F_log(m, base.F64_convert_i32_s(v65))
							mBase = m.M
							v84 = base.F64_div(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v61*v65<<(uint(int32(1))%32)), base.F64_div(base.F64_mul(v59, float64(0.55)), base.F64_mul(base.F64_add(v63, v64), v75))), base.F64_convert_i32_s(v65*base.I32_trunc_sat_f64_s(base.F64_div(v59, v75)))), v56)
							if base.F64_gt(v84, v64) != 0 {
								v87 = v64
							} else {
								v87 = v84
							}
							v89 = v87
						} else {
							v89 = float64(1)
						}
						v92 = *(*int32)(unsafe.Add(mBase, uint32(v55)+8))
						F_get_tablespace_page_costs(m, v92, int32(0), v18)
						mBase = m.M
						v95 = m.ExcPending
						if v95 != 0 {
							return
						} else {
							v96 = *(*float64)(unsafe.Add(mBase, uint32(v18)+48))
							v97 = *(*float64)(unsafe.Add(mBase, uint32(v18)+24))
							v98 = base.F64_mul(v97, v89)
							if base.F64_lt(v89, float64(0.5)) == int32(0) {
								v120 = v98
							} else {
								v103 = base.F64_mul(v96, v89)
								v104 = *(*int32)(unsafe.Add(mBase, uint32(l1)+72))
								v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+12))
								v106 = *(*int32)(unsafe.Add(mBase, uint32(v105)+116))
								v107 = base.F64_convert_i32_u(v106)
								if base.F64_gt(v103, v107) == int32(0) {
									v120 = v98
								} else {
									v112 = *(*float64)(unsafe.Add(mBase, uint32(v18)))
									v114 = *(*float64)(unsafe.Add(mBase, uint32(v18)+64))
									v120 = base.F64_add(base.F64_mul(base.F64_sub(v107, v103), v112), base.F64_add(v98, base.F64_mul(v103, base.F64_sub(v112, v114))))
								}
							}
							*(*float64)(unsafe.Add(mBase, uint32(l3))) = v120
							*(*float64)(unsafe.Add(mBase, uint32(l4))) = v97
							v124 = *(*float64)(unsafe.Add(mBase, uint32(v18)+32))
							*(*float64)(unsafe.Add(mBase, uint32(l5))) = v124
							v126 = *(*float64)(unsafe.Add(mBase, uint32(v18)+40))
							*(*float64)(unsafe.Add(mBase, uint32(l6))) = v126
							*(*float64)(unsafe.Add(mBase, uint32(l7))) = v96
							m.G0 = v18 + int32(80)
							return
						}
					}
				}
			}
		}
	} else {
		v129 = int64(9218868437227405312)
		*(*int64)(unsafe.Add(mBase, uint32(l3))) = v129
		*(*int64)(unsafe.Add(mBase, uint32(l4))) = v129
		v133 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(l5))) = v133
		*(*int64)(unsafe.Add(mBase, uint32(l6))) = v133
		*(*int64)(unsafe.Add(mBase, uint32(l7))) = v133
		*(*int32)(unsafe.Add(mBase, uint32(l1)+40)) = int32(2)
		m.G0 = v18 + int32(80)
		return
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
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l2+v10<<(uint(int32(4))%32)+l1*int32(100))+88))
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
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v30*int32(100)+(l2+v36<<(uint(int32(4))%32)))+88))
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
	F_errfinish(m, int32(_a_F_hypothetical_check_argtypes_1), int32(1159), int32(_a_F_hypothetical_check_argtypes_2))
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
	F_errfinish(m, int32(_a_F_hypothetical_check_argtypes_1), int32(1151), int32(_a_F_hypothetical_check_argtypes_2))
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
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v93 int32
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
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v166 int64
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v190 int64
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v208 int64
	_ = v208
	var v209 int32
	_ = v209
	var v210 int32
	_ = v210
	var v213 int32
	_ = v213
	var v214 int32
	_ = v214
	var v226 int64
	_ = v226
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
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
	v235 = m.ExcPending
	if v235 != 0 {
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
	*(*int32)(unsafe.Add(mBase, uint32(v142+v141<<(uint(int32(2))%32)))) = l1
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v149 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v147+v141))) = uint8(v149)
	v151 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)))
	v153 = v151 & int32(_a_F_hypothetical_rank_common_0)
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+4)) = uint16(v153)
	v155 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v156 = *(*int32)(unsafe.Add(mBase, uint32(v155)))
	*(*uint16)(unsafe.Add(mBase, uint32(v41)+6)) = uint16(v156)
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
	v122 = v22 + v106<<(uint(int32(3))%32)
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v122)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v116+v106<<(uint(int32(2))%32)))) = v123
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v127 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v122)+12)))
	*(*uint8)(unsafe.Add(mBase, uint32(v125+v106))) = uint8(v127)
	v141 = v33
	goto L10
L15:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v70 = int32(2)
	v74 = v59 | int32(1)
	v75 = int32(3)
	v77 = v22 + v74<<(uint(v75)%32)
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v77)))
	*(*int32)(unsafe.Add(mBase, uint32(v69+v59<<(uint(v70)%32)))) = v78
	v80 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v82 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v77)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v80+v59))) = uint8(v82)
	v84 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v89 = v59 + v70
	v92 = v22 + v89<<(uint(v75)%32)
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v92)))
	*(*int32)(unsafe.Add(mBase, uint32(v84+v74<<(uint(v70)%32)))) = v93
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v97 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92)+4)))
	*(*uint8)(unsafe.Add(mBase, uint32(v74+v95))) = uint8(v97)
	v100 = v63 + v70
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
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_tuplesort_puttupleslot(m, v158, v41)
	mBase = m.M
	v160 = m.ExcPending
	if v160 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	F_tuplesort_performsort(m, v161)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L7
	} else {
		goto L21
	}
L21:
	;
	v164 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v23)+24)) = uint8(v164)
	v166 = int64(1)
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v171 = F_tuplesort_gettupleslot(m, v167, v164, v164, v41, int32(0))
	mBase = m.M
	v172 = m.ExcPending
	if v172 != 0 {
		goto L7
	} else {
		goto L23
	}
L22:
	;
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v41)+8))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+12))
	m.T0[v228].(func(*base.Module, int32))(m, v41)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L7
	} else {
		goto L41
	}
L23:
	;
	if v171 == int32(0) {
		v226 = v166
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v190 = v166
	goto L25
L25:
	;
	v191 = int32(*(*int16)(unsafe.Add(mBase, uint32(v41)+6)))
	if v191 <= v33 {
		goto L27
	} else {
		goto L28
	}
L26:
	;
	v226 = v208
	goto L22
L27:
	;
	F_slot_getsomeattrs_int(m, v41, v33+int32(1))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L30
	}
L28:
	;
	goto L29
L29:
	;
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v41)+20))
	v197 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195+v33))))
	if v197 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L30:
	;
	goto L29
L31:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v200+v33<<(uint(int32(2))%32))))
	if v202 != 0 {
		v226 = v190
		goto L22
	} else {
		goto L34
	}
L32:
	;
	goto L33
L33:
	;
	v204 = *(*int32)(unsafe.Add(mBase, _c_F_hypothetical_rank_common[0]))
	if v204 != 0 {
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
	v206 = m.ExcPending
	if v206 != 0 {
		goto L7
	} else {
		goto L38
	}
L36:
	;
	goto L37
L37:
	;
	v208 = v190 + int64(1)
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v23)+8))
	v210 = int32(1)
	v213 = F_tuplesort_gettupleslot(m, v209, v210, v210, v41, int32(0))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L7
	} else {
		goto L39
	}
L38:
	;
	goto L37
L39:
	;
	if v213 != 0 {
		v190 = v208
		goto L25
	} else {
		goto L40
	}
L40:
	;
	goto L26
L41:
	;
	return v226
L42:
	;
	F_errmsg_internal(m, int32(_a_F_hypothetical_rank_common_1), int32(0))
	mBase = m.M
	v239 = m.ExcPending
	if v239 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	F_errfinish(m, int32(_a_F_hypothetical_rank_common_2), int32(1194), int32(_a_F_hypothetical_rank_common_3))
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
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
