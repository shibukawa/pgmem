package p0

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_FetchPreparedStatement(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = *(*int32)(unsafe.Add(mBase, _c_F_FetchPreparedStatement[0]))
	if v11 != 0 {
		v12 = int32(0)
		v14 = F_hash_search(m, v11, l0, v12, v12)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int32(0)
		} else {
			v18 = v14
			v19 = int32(0)
			if base.B2i32(l1 == v19)|v18 == v19 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v27 = m.ExcPending
				if v27 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(386))
					mBase = m.M
					v30 = m.ExcPending
					if v30 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
						F_errmsg(m, int32(_a_F_FetchPreparedStatement_0), v8)
						mBase = m.M
						v34 = m.ExcPending
						if v34 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_FetchPreparedStatement_1), int32(454), int32(_a_F_FetchPreparedStatement_2))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
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
				m.G0 = v8 + int32(16)
				return v18
			}
		}
	} else {
		v18 = int32(0)
		v19 = int32(0)
		if base.B2i32(l1 == v19)|v18 == v19 {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v27 = m.ExcPending
			if v27 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(386))
				mBase = m.M
				v30 = m.ExcPending
				if v30 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
					F_errmsg(m, int32(_a_F_FetchPreparedStatement_0), v8)
					mBase = m.M
					v34 = m.ExcPending
					if v34 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(_a_F_FetchPreparedStatement_1), int32(454), int32(_a_F_FetchPreparedStatement_2))
						mBase = m.M
						v39 = m.ExcPending
						if v39 != 0 {
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
			m.G0 = v8 + int32(16)
			return v18
		}
	}
}
func F_FetchStatementTargetList(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
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
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
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
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
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
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v99 int32
	_ = v99
	if l0 == int32(0) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v5 = l0
	goto L3
L3:
	;
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v7 == int32(67) {
		goto L5
	} else {
		goto L6
	}
L4:
	;
	goto L1
L5:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(v5)+4))
	if v10 != int32(6) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v21 = v5
	v22 = v7
	goto L7
L7:
	;
	if v22 == int32(330) {
		goto L14
	} else {
		goto L15
	}
L8:
	;
	if v10 == int32(1) {
		goto L11
	} else {
		goto L12
	}
L9:
	;
	goto L10
L10:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v5)+28))
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
	v21 = v19
	v22 = v20
	goto L7
L11:
	;
	v15 = *(*int32)(unsafe.Add(mBase, uint32(v5)+76))
	return v15
L12:
	;
	goto L13
L13:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v5)+96))
	return v17
L14:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v21)+4))
	if v25 != int32(6) {
		goto L17
	} else {
		goto L18
	}
L15:
	;
	v41 = v21
	v42 = v22
	goto L16
L16:
	;
	if v42 != int32(203) {
		goto L24
	} else {
		goto L25
	}
L17:
	;
	if v25 == int32(1) {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	goto L19
L19:
	;
	v39 = *(*int32)(unsafe.Add(mBase, uint32(v21)+88))
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)))
	v41 = v39
	v42 = v40
	goto L16
L20:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(v30)+44))
	return v31
L21:
	;
	goto L22
L22:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21)+24)))
	if v33 != int32(1) {
		goto L1
	} else {
		goto L23
	}
L23:
	;
	v36 = *(*int32)(unsafe.Add(mBase, uint32(v21)+36))
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v36)+44))
	return v37
L24:
	;
	if v42 != int32(253) {
		goto L1
	} else {
		goto L27
	}
L25:
	;
	goto L26
L26:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v60 = F_GetPortalByName(m, v59)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L28
	} else {
		goto L32
	}
L27:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v41)+4))
	v49 = F_FetchPreparedStatement(m, v47, int32(1))
	mBase = m.M
	v52 = m.ExcPending
	if v52 != 0 {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	return int32(0)
L29:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v49)+64))
	v54 = F_CachedPlanGetTargetList(m, v53)
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L28
	} else {
		goto L30
	}
L30:
	;
	v56 = F_copyObjectImpl(m, v54)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L28
	} else {
		goto L31
	}
L31:
	;
	return v56
L32:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v60)+72))
	if v62 == int32(4) {
		goto L1
	} else {
		goto L33
	}
L33:
	;
	v68 = *(*int32)(unsafe.Add(mBase, uint32(v60)+56))
	if v68 != 0 {
		goto L36
	} else {
		goto L37
	}
L34:
	;
	if v99 != 0 {
		v5 = v99
		goto L3
	} else {
		goto L47
	}
L35:
	;
	goto L34
L36:
	;
	v69 = *(*int32)(unsafe.Add(mBase, uint32(v68)+4))
	if v69 <= int32(0) {
		v99 = int32(0)
		goto L35
	} else {
		goto L39
	}
L37:
	;
	goto L38
L38:
	;
	v99 = int32(0)
	goto L35
L39:
	;
	v72 = int32(0)
	if v72 < v69 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v75 = v69
	goto L42
L41:
	;
	v75 = v72
	goto L42
L42:
	;
	v76 = *(*int32)(unsafe.Add(mBase, uint32(v68)+12))
	v78 = int32(0)
	goto L43
L43:
	;
	v85 = *(*int32)(unsafe.Add(mBase, uint32(v76+v78<<(uint(int32(2))%32))))
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v85)+26)))
	if v86 == int32(1) {
		v99 = v85
		goto L35
	} else {
		goto L45
	}
L44:
	;
	goto L38
L45:
	;
	v90 = v78 + int32(1)
	if v90 != v75 {
		v78 = v90
		goto L43
	} else {
		goto L46
	}
L46:
	;
	goto L44
L47:
	;
	goto L4
}
func F_FinishSortSupportFunction(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v56 int32
	_ = v56
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v11 = F_get_opfamily_proc(m, l0, l1, l1, int32(2))
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return
	} else {
		if v11 != 0 {
			v14 = F_OidFunctionCall1Coll(m, v11, int32(0), l2)
			mBase = m.M
			v15 = m.ExcPending
			if v15 != 0 {
				return
			} else {
				v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
				if v16 == int32(0) {
					v20 = F_get_opfamily_proc(m, l0, l1, l1, int32(1))
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return
					} else {
						if v20 == int32(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
								*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
								*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(1)
								F_errmsg_internal(m, int32(_a_F_FinishSortSupportFunction_0), v8)
								mBase = m.M
								v64 = m.ExcPending
								if v64 != 0 {
									return
								} else {
									F_errfinish(m, int32(_a_F_FinishSortSupportFunction_1), int32(119), int32(_a_F_FinishSortSupportFunction_2))
									mBase = m.M
									v69 = m.ExcPending
									if v69 != 0 {
										return
									} else {
										base.Wasm_trap_unreachable()
										for {
										}
									}
								}
							}
						} else {
							v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							v26 = F_MemoryContextAlloc(m, v24, int32(64))
							mBase = m.M
							v27 = m.ExcPending
							if v27 != 0 {
								return
							} else {
								v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
								F_fmgr_info_cxt(m, v20, v26, v28)
								mBase = m.M
								v30 = m.ExcPending
								if v30 != 0 {
									return
								} else {
									*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = int64(0)
									*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v26
									v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
									v35 = int32(0)
									*(*uint8)(unsafe.Add(mBase, uint32(v26)+60)) = uint8(v35)
									*(*uint8)(unsafe.Add(mBase, uint32(v26)+52)) = uint8(v35)
									v39 = int32(2)
									*(*uint16)(unsafe.Add(mBase, uint32(v26)+46)) = uint16(v39)
									*(*uint8)(unsafe.Add(mBase, uint32(v26)+44)) = uint8(v35)
									*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v34
									*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(1819)
									*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v26
									m.G0 = v8 + int32(16)
									return
								}
							}
						}
					}
				} else {
					m.G0 = v8 + int32(16)
					return
				}
			}
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(l2)+16))
			if v16 == int32(0) {
				v20 = F_get_opfamily_proc(m, l0, l1, l1, int32(1))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					if v20 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v56 = m.ExcPending
						if v56 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = l1
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = int32(1)
							F_errmsg_internal(m, int32(_a_F_FinishSortSupportFunction_0), v8)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								F_errfinish(m, int32(_a_F_FinishSortSupportFunction_1), int32(119), int32(_a_F_FinishSortSupportFunction_2))
								mBase = m.M
								v69 = m.ExcPending
								if v69 != 0 {
									return
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						v24 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
						v26 = F_MemoryContextAlloc(m, v24, int32(64))
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return
						} else {
							v28 = *(*int32)(unsafe.Add(mBase, uint32(l2)))
							F_fmgr_info_cxt(m, v20, v26, v28)
							mBase = m.M
							v30 = m.ExcPending
							if v30 != 0 {
								return
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v26)+32)) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v26)+28)) = v26
								v34 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
								v35 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v26)+60)) = uint8(v35)
								*(*uint8)(unsafe.Add(mBase, uint32(v26)+52)) = uint8(v35)
								v39 = int32(2)
								*(*uint16)(unsafe.Add(mBase, uint32(v26)+46)) = uint16(v39)
								*(*uint8)(unsafe.Add(mBase, uint32(v26)+44)) = uint8(v35)
								*(*int32)(unsafe.Add(mBase, uint32(v26)+40)) = v34
								*(*int32)(unsafe.Add(mBase, uint32(l2)+16)) = int32(1819)
								*(*int32)(unsafe.Add(mBase, uint32(l2)+12)) = v26
								m.G0 = v8 + int32(16)
								return
							}
						}
					}
				}
			} else {
				m.G0 = v8 + int32(16)
				return
			}
		}
	}
}
func F_FlushPages(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v67 int32
	_ = v67
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
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
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
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v160 int32
	_ = v160
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v195 int32
	_ = v195
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v227 int32
	_ = v227
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v273 int32
	_ = v273
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v285 int32
	_ = v285
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v313 int32
	_ = v313
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v331 int32
	_ = v331
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v367 int32
	_ = v367
	var v371 int32
	_ = v371
	var v375 int32
	_ = v375
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v405 int32
	_ = v405
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v427 int32
	_ = v427
	var v428 int32
	_ = v428
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v435 int32
	_ = v435
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v468 int32
	_ = v468
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v489 int32
	_ = v489
	var v490 int32
	_ = v490
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v508 int32
	_ = v508
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v522 int32
	_ = v522
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v534 int32
	_ = v534
	var v541 int32
	_ = v541
	var v544 int32
	_ = v544
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v559 int32
	_ = v559
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v589 int32
	_ = v589
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v597 int32
	_ = v597
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v609 int32
	_ = v609
	var v613 int32
	_ = v613
	var v619 int32
	_ = v619
	var v621 int32
	_ = v621
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v639 int32
	_ = v639
	var v643 int32
	_ = v643
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v661 int32
	_ = v661
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v674 int32
	_ = v674
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v698 int32
	_ = v698
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v722 int32
	_ = v722
	var v724 int32
	_ = v724
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v731 int32
	_ = v731
	var v735 int32
	_ = v735
	var v746 int32
	_ = v746
	var v757 int32
	_ = v757
	var v759 int32
	_ = v759
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v789 int32
	_ = v789
	var v791 int32
	_ = v791
	var v798 int32
	_ = v798
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v810 int32
	_ = v810
	var v814 int32
	_ = v814
	var v815 int32
	_ = v815
	var v823 int32
	_ = v823
	var v828 int32
	_ = v828
	var v832 int32
	_ = v832
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v844 int32
	_ = v844
	v17 = m.G0
	v19 = v17 - int32(48)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v23 = F_HnswNewBuffer(m, v21, v22)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	F_PageInit(m, v42, int32(_a_F_FlushPages_0), int32(8))
	mBase = m.M
	v46 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v42)+16)))
	v47 = v42 + v46
	v48 = int32(_a_F_FlushPages_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v47)+6)) = uint16(v48)
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = int32(-1)
	goto L7
L2:
	;
	return
L3:
	;
	if v23 < int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v28 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[0]))
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28+(v23^int32(-1))<<(uint(int32(2))%32))))
	v42 = v34
	goto L1
L5:
	;
	goto L6
L6:
	;
	v36 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[1]))
	v42 = v36 + v23<<(uint(int32(13))%32) + int32(-8192)
	goto L1
L7:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v42)+24)) = int64(7135799635)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+32)) = v54
	v56 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+36)) = uint16(v56)
	v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v42)+48)) = int32(-1)
	*(*int64)(unsafe.Add(mBase, uint32(v42)+40)) = int64(-281470681743361)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+38)) = uint16(v58)
	v64 = int32(52)
	*(*uint16)(unsafe.Add(mBase, uint32(v42)+12)) = uint16(v64)
	F_MarkBufferDirty(m, v23)
	mBase = m.M
	v67 = m.ExcPending
	if v67 != 0 {
		goto L2
	} else {
		goto L8
	}
L8:
	;
	F_UnlockReleaseBuffer(m, v23)
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		goto L2
	} else {
		goto L9
	}
L9:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v73)+4))
	v76 = F_palloc0(m, int32(_a_F_FlushPages_0))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L2
	} else {
		goto L10
	}
L10:
	;
	v79 = F_palloc0(m, int32(_a_F_FlushPages_0))
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L2
	} else {
		goto L11
	}
L11:
	;
	v81 = F_HnswNewBuffer(m, v71, v70)
	mBase = m.M
	v82 = m.ExcPending
	if v82 != 0 {
		goto L2
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+44)) = v81
	if v81 < int32(0) {
		goto L14
	} else {
		goto L15
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+40)) = v101
	F_PageInit(m, v101, int32(_a_F_FlushPages_0), int32(8))
	mBase = m.M
	v106 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v101)+16)))
	v107 = v101 + v106
	v108 = int32(_a_F_FlushPages_1)
	*(*uint16)(unsafe.Add(mBase, uint32(v107)+6)) = uint16(v108)
	*(*int32)(unsafe.Add(mBase, uint32(v107))) = int32(-1)
	goto L17
L14:
	;
	v87 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[0]))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v87+(v81^int32(-1))<<(uint(int32(2))%32))))
	v101 = v93
	goto L13
L15:
	;
	goto L16
L16:
	;
	v95 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[1]))
	v101 = v95 + v81<<(uint(int32(13))%32) + int32(-8192)
	goto L13
L17:
	;
	v113 = v101
	v120 = v74
	goto L22
L18:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v832 = m.ExcPending
	if v832 != 0 {
		goto L2
	} else {
		goto L216
	}
L19:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v814 = m.ExcPending
	if v814 != 0 {
		goto L2
	} else {
		goto L213
	}
L20:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L2
	} else {
		goto L209
	}
L21:
	;
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v504 < int32(0) {
		goto L142
	} else {
		goto L143
	}
L22:
	;
	if v120 == int32(0) {
		goto L21
	} else {
		goto L24
	}
L23:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v489 = m.ExcPending
	if v489 != 0 {
		goto L2
	} else {
		goto L138
	}
L24:
	;
	v130 = v72 + v120
	if v72 != 0 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v133 = v130 - int32(1)
	goto L27
L26:
	;
	v133 = v120
	goto L27
L27:
	;
	if v72 == int32(0) {
		goto L29
	} else {
		goto L30
	}
L28:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	base.MemoryFill(m, v76, int32(0), int32(_a_F_FlushPages_0))
	v151 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145))))
	if v151 == int32(1) {
		goto L34
	} else {
		goto L35
	}
L29:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(v120)+88))
	v145 = v136
	goto L28
L30:
	;
	goto L31
L31:
	;
	v137 = int32(0)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(v130)+87))
	if v138 == v137 {
		v145 = v137
		goto L28
	} else {
		goto L32
	}
L32:
	;
	v145 = v72 + v138 - int32(1)
	goto L28
L33:
	;
	v177 = F_add_size(m, int32(72), v176)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L2
	} else {
		goto L44
	}
L34:
	;
	v155 = int32(18)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v145)+1)))
	if v157 == v155 {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	goto L36
L36:
	;
	v168 = int32(1)
	if v151&v168 != 0 {
		v176 = int32(base.Ui32(v151) >> (uint(v168) % 32))
		goto L33
	} else {
		goto L43
	}
L37:
	;
	v160 = v155
	goto L39
L38:
	;
	v160 = int32(2)
	goto L39
L39:
	;
	if base.Ui32((v157-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v167 = int32(6)
	goto L42
L41:
	;
	v167 = v160
	goto L42
L42:
	;
	v176 = v167
	goto L33
L43:
	;
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v145)))
	v176 = int32(base.Ui32(v172) >> (uint(int32(2)) % 32))
	goto L33
L44:
	;
	v181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+65)))
	v183 = F_add_size(m, v181, int32(2))
	mBase = m.M
	v184 = m.ExcPending
	if v184 != 0 {
		goto L2
	} else {
		goto L45
	}
L45:
	;
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v186 = F_mul_size(m, v183, v185)
	mBase = m.M
	v187 = m.ExcPending
	if v187 != 0 {
		goto L2
	} else {
		goto L46
	}
L46:
	;
	v188 = F_mul_size(m, int32(6), v186)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L2
	} else {
		goto L47
	}
L47:
	;
	v190 = F_add_size(m, int32(4), v188)
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L2
	} else {
		goto L48
	}
L48:
	;
	v195 = (v177 + int32(7)) & int32(-8)
	if base.Ui32(int32(_a_F_FlushPages_2)) <= base.Ui32(v195) {
		goto L20
	} else {
		goto L49
	}
L49:
	;
	v201 = (v190 + int32(7)) & int32(-8)
	v202 = v201 + v195
	if v72 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L50:
	;
	v375 = int32(4)
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113)+14)))
	v377 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113)+12)))
	v378 = v376 - v377
	if v378 <= v375 {
		goto L103
	} else {
		goto L104
	}
L51:
	;
	v216 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v76))) = uint8(v216)
	v218 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+65)))
	v219 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+2)) = uint8(v219)
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+1)) = uint8(v218)
	v222 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(v76)+3)) = uint8(v222)
	v224 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if v224 == v219 {
		goto L75
	} else {
		goto L76
	}
L52:
	;
	v206 = *(*int32)(unsafe.Add(mBase, uint32(v133)+88))
	v215 = v206
	goto L51
L53:
	;
	goto L54
L54:
	;
	v207 = int32(0)
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v133)+88))
	if v208 == v207 {
		v215 = v207
		goto L51
	} else {
		goto L55
	}
L55:
	;
	v215 = v72 + v208 - int32(1)
	goto L51
L56:
	;
	v346 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215))))
	if v346 == int32(1) {
		goto L88
	} else {
		goto L89
	}
L57:
	;
	v341 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+62)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+62)) = uint16(v341)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v133)+58))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+58)) = v343
	goto L56
L58:
	;
	v337 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+62)) = uint16(v337)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+58)) = int32(-1)
	goto L56
L59:
	;
	v329 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+56)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+56)) = uint16(v329)
	v331 = *(*int32)(unsafe.Add(mBase, uint32(v133)+52))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+52)) = v331
	v333 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if base.Ui32(int32(9)) < base.Ui32(v333) {
		goto L57
	} else {
		goto L86
	}
L60:
	;
	v325 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+56)) = uint16(v325)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+52)) = int32(-1)
	goto L58
L61:
	;
	v317 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+50)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+50)) = uint16(v317)
	v319 = *(*int32)(unsafe.Add(mBase, uint32(v133)+46))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+46)) = v319
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if base.Ui32(int32(8)) < base.Ui32(v321) {
		goto L59
	} else {
		goto L85
	}
L62:
	;
	v313 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+50)) = uint16(v313)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+46)) = int32(-1)
	goto L60
L63:
	;
	v305 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+44)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+44)) = uint16(v305)
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v133)+40))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+40)) = v307
	v309 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if base.Ui32(int32(7)) < base.Ui32(v309) {
		goto L61
	} else {
		goto L84
	}
L64:
	;
	v301 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+44)) = uint16(v301)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+40)) = int32(-1)
	goto L62
L65:
	;
	v293 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+38)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+38)) = uint16(v293)
	v295 = *(*int32)(unsafe.Add(mBase, uint32(v133)+34))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+34)) = v295
	v297 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if base.Ui32(int32(6)) < base.Ui32(v297) {
		goto L63
	} else {
		goto L83
	}
L66:
	;
	v289 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+38)) = uint16(v289)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+34)) = int32(-1)
	goto L64
L67:
	;
	v281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+32)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+32)) = uint16(v281)
	v283 = *(*int32)(unsafe.Add(mBase, uint32(v133)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+28)) = v283
	v285 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if base.Ui32(int32(5)) < base.Ui32(v285) {
		goto L65
	} else {
		goto L82
	}
L68:
	;
	v277 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+32)) = uint16(v277)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+28)) = int32(-1)
	goto L66
L69:
	;
	v269 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+26)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+26)) = uint16(v269)
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v133)+22))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+22)) = v271
	v273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if base.Ui32(int32(4)) < base.Ui32(v273) {
		goto L67
	} else {
		goto L81
	}
L70:
	;
	v265 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+26)) = uint16(v265)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+22)) = int32(-1)
	goto L68
L71:
	;
	v257 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+20)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+20)) = uint16(v257)
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v133)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = v259
	v261 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if base.Ui32(int32(3)) < base.Ui32(v261) {
		goto L69
	} else {
		goto L80
	}
L72:
	;
	v253 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+20)) = uint16(v253)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+16)) = int32(-1)
	goto L70
L73:
	;
	v245 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+14)))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+14)) = uint16(v245)
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v133)+10))
	*(*int32)(unsafe.Add(mBase, uint32(v76)+10)) = v247
	v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if base.Ui32(int32(2)) < base.Ui32(v249) {
		goto L71
	} else {
		goto L79
	}
L74:
	;
	v241 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+14)) = uint16(v241)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+10)) = int32(-1)
	goto L72
L75:
	;
	v227 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+8)) = uint16(v227)
	*(*int32)(unsafe.Add(mBase, uint32(v76)+4)) = int32(-1)
	goto L74
L76:
	;
	goto L77
L77:
	;
	v232 = v76 + int32(4)
	v233 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+8)))
	*(*uint16)(unsafe.Add(mBase, uint32(v232)+4)) = uint16(v233)
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v232))) = v235
	v237 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v133)+64)))
	if base.Ui32(int32(1)) < base.Ui32(v237) {
		goto L73
	} else {
		goto L78
	}
L78:
	;
	goto L74
L79:
	;
	goto L72
L80:
	;
	goto L70
L81:
	;
	goto L68
L82:
	;
	goto L66
L83:
	;
	goto L64
L84:
	;
	goto L62
L85:
	;
	goto L60
L86:
	;
	goto L58
L87:
	;
	if v371 != 0 {
		goto L98
	} else {
		goto L99
	}
L88:
	;
	v350 = int32(18)
	v352 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v215)+1)))
	if v352 == v350 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	v363 = int32(1)
	if v346&v363 != 0 {
		v371 = int32(base.Ui32(v346) >> (uint(v363) % 32))
		goto L87
	} else {
		goto L97
	}
L91:
	;
	v355 = v350
	goto L93
L92:
	;
	v355 = int32(2)
	goto L93
L93:
	;
	if base.Ui32((v352-int32(1))&int32(255)) < base.Ui32(int32(3)) {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v362 = int32(6)
	goto L96
L95:
	;
	v362 = v355
	goto L96
L96:
	;
	v371 = v362
	goto L87
L97:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v215)))
	v371 = int32(base.Ui32(v367) >> (uint(int32(2)) % 32))
	goto L87
L98:
	;
	base.MemoryCopy(m, v76+int32(72), v215, v371)
	goto L100
L99:
	;
	goto L100
L100:
	;
	goto L50
L101:
	;
	v409 = *(*int32)(unsafe.Add(mBase, uint32(v19)+44))
	if v409 < int32(0) {
		goto L117
	} else {
		goto L118
	}
L102:
	;
	if base.Ui32(v195) <= base.Ui32(v381-int32(4)) {
		goto L106
	} else {
		goto L107
	}
L103:
	;
	v381 = v375
	goto L105
L104:
	;
	v381 = v378
	goto L105
L105:
	;
	goto L102
L106:
	;
	v386 = v202 | int32(4)
	if base.Ui32(int32(_a_F_FlushPages_3)) < base.Ui32(v386) {
		v407 = v113
		goto L101
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	F_HnswBuildAppendPage(m, v71, v19+int32(44), v19+int32(40), v70)
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L2
	} else {
		goto L115
	}
L109:
	;
	v389 = int32(4)
	v390 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113)+14)))
	v391 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v113)+12)))
	v392 = v390 - v391
	if v392 <= v389 {
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if base.Ui32(v386) <= base.Ui32(v395-int32(4)) {
		v407 = v113
		goto L101
	} else {
		goto L114
	}
L111:
	;
	v395 = v389
	goto L113
L112:
	;
	v395 = v392
	goto L113
L113:
	;
	goto L110
L114:
	;
	goto L108
L115:
	;
	v406 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v407 = v406
	goto L101
L116:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v133)+76)) = v428
	v430 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v407)+12)))
	v432 = base.B2i32(base.Ui32(int32(_a_F_FlushPages_4)) < base.Ui32(v202))
	v433 = v428 + v432
	*(*int32)(unsafe.Add(mBase, uint32(v133)+84)) = v433
	v435 = int32(1)
	if base.Ui32(v430) < base.Ui32(int32(25)) {
		goto L120
	} else {
		goto L121
	}
L117:
	;
	v413 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[2]))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v413+(v409^int32(-1))<<(uint(int32(6))%32))+16))
	v428 = v419
	goto L116
L118:
	;
	goto L119
L119:
	;
	v421 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[3]))
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v421+v409<<(uint(int32(6))%32)+int32(-64))+16))
	v428 = v427
	goto L116
L120:
	;
	v444 = v435
	goto L122
L121:
	;
	v444 = int32(base.Ui32(v430+int32(_a_F_FlushPages_5))>>(uint(int32(2))%32)) + v435
	goto L122
L122:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+80)) = uint16(v444)
	v446 = int32(1)
	if base.Ui32(int32(_a_F_FlushPages_4)) < base.Ui32(v202) {
		goto L123
	} else {
		goto L124
	}
L123:
	;
	v449 = v446
	goto L125
L124:
	;
	v449 = v444 + v446
	goto L125
L125:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v133)+82)) = uint16(v449)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+68)) = uint16(v449)
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+66)) = uint16(v433)
	v454 = int32(base.Ui32(v433) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v76)+64)) = uint16(v454)
	v456 = int32(0)
	v458 = F_PageAddItemExtended(m, v407, v76, v195, v456, v456)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L2
	} else {
		goto L126
	}
L126:
	;
	v460 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+80)))
	if v458 != v460 {
		goto L19
	} else {
		goto L127
	}
L127:
	;
	v462 = int32(4)
	v463 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v407)+14)))
	v464 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v407)+12)))
	v465 = v463 - v464
	if v465 <= v462 {
		goto L129
	} else {
		goto L130
	}
L128:
	;
	if base.Ui32(v468-int32(4)) < base.Ui32(v201) {
		goto L132
	} else {
		goto L133
	}
L129:
	;
	v468 = v462
	goto L131
L130:
	;
	v468 = v465
	goto L131
L131:
	;
	goto L128
L132:
	;
	F_HnswBuildAppendPage(m, v71, v19+int32(44), v19+int32(40), v70)
	mBase = m.M
	v477 = m.ExcPending
	if v477 != 0 {
		goto L2
	} else {
		goto L135
	}
L133:
	;
	v479 = v407
	goto L134
L134:
	;
	v480 = int32(0)
	v482 = F_PageAddItemExtended(m, v479, v79, v201, v480, v480)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L2
	} else {
		goto L136
	}
L135:
	;
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v19)+40))
	v479 = v478
	goto L134
L136:
	;
	v484 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v133)+82)))
	if v482 == v484 {
		v113 = v479
		v120 = v146
		goto L22
	} else {
		goto L137
	}
L137:
	;
	goto L23
L138:
	;
	v490 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = v490 + int32(4)
	F_errmsg_internal(m, int32(_a_F_FlushPages_6), v19+int32(16))
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L2
	} else {
		goto L139
	}
L139:
	;
	F_errfinish(m, int32(_a_F_FlushPages_7), int32(234), int32(_a_F_FlushPages_8))
	mBase = m.M
	v503 = m.ExcPending
	if v503 != 0 {
		goto L2
	} else {
		goto L140
	}
L140:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L141:
	;
	F_MarkBufferDirty(m, v504)
	mBase = m.M
	v525 = m.ExcPending
	if v525 != 0 {
		goto L2
	} else {
		goto L145
	}
L142:
	;
	v508 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[2]))
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v508+(v504^int32(-1))<<(uint(int32(6))%32))+16))
	v523 = v514
	goto L141
L143:
	;
	goto L144
L144:
	;
	v516 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[3]))
	v522 = *(*int32)(unsafe.Add(mBase, uint32(v516+v504<<(uint(int32(6))%32)+int32(-64))+16))
	v523 = v522
	goto L141
L145:
	;
	F_UnlockReleaseBuffer(m, v504)
	mBase = m.M
	v527 = m.ExcPending
	if v527 != 0 {
		goto L2
	} else {
		goto L146
	}
L146:
	;
	v528 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	if v72 == int32(0) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	F_HnswUpdateMetaPage(m, v71, int32(2), v541, v523, v70, int32(1))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L2
	} else {
		goto L152
	}
L148:
	;
	v532 = *(*int32)(unsafe.Add(mBase, uint32(v528)+48))
	v541 = v532
	goto L147
L149:
	;
	goto L150
L150:
	;
	v533 = int32(0)
	v534 = *(*int32)(unsafe.Add(mBase, uint32(v528)+48))
	if v534 == v533 {
		v541 = v533
		goto L147
	} else {
		goto L151
	}
L151:
	;
	v541 = v534 + v72 - int32(1)
	goto L147
L152:
	;
	F_pfree(m, v76)
	mBase = m.M
	v546 = m.ExcPending
	if v546 != 0 {
		goto L2
	} else {
		goto L153
	}
L153:
	;
	F_pfree(m, v79)
	mBase = m.M
	v548 = m.ExcPending
	if v548 != 0 {
		goto L2
	} else {
		goto L154
	}
L154:
	;
	v549 = *(*int32)(unsafe.Add(mBase, uint32(l0)+204))
	v550 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v551 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v552 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v554 = *(*int32)(unsafe.Add(mBase, uint32(v553)+4))
	v556 = F_palloc0(m, int32(_a_F_FlushPages_0))
	mBase = m.M
	v557 = m.ExcPending
	if v557 != 0 {
		goto L2
	} else {
		goto L155
	}
L155:
	;
	if v554 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v559 = v554
	goto L159
L157:
	;
	goto L158
L158:
	;
	F_pfree(m, v556)
	mBase = m.M
	v785 = m.ExcPending
	if v785 != 0 {
		goto L2
	} else {
		goto L207
	}
L159:
	;
	if v549 != 0 {
		goto L161
	} else {
		goto L162
	}
L160:
	;
	goto L158
L161:
	;
	v579 = v559 + v549 - int32(1)
	goto L163
L162:
	;
	v579 = v559
	goto L163
L163:
	;
	v580 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+65)))
	v582 = F_add_size(m, v580, int32(2))
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L2
	} else {
		goto L164
	}
L164:
	;
	v584 = F_mul_size(m, v582, v550)
	mBase = m.M
	v585 = m.ExcPending
	if v585 != 0 {
		goto L2
	} else {
		goto L165
	}
L165:
	;
	v586 = F_mul_size(m, int32(6), v584)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L2
	} else {
		goto L166
	}
L166:
	;
	v588 = F_add_size(m, int32(4), v586)
	mBase = m.M
	v589 = m.ExcPending
	if v589 != 0 {
		goto L2
	} else {
		goto L167
	}
L167:
	;
	v590 = *(*int32)(unsafe.Add(mBase, uint32(v579)))
	base.MemoryFill(m, v556, int32(0), int32(_a_F_FlushPages_0))
	v595 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[4]))
	if v595 != 0 {
		goto L168
	} else {
		goto L169
	}
L168:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v597 = m.ExcPending
	if v597 != 0 {
		goto L2
	} else {
		goto L171
	}
L169:
	;
	goto L170
L170:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v579)+84))
	v603 = int32(0)
	v605 = F_ReadBufferExtended(m, v552, v551, v602, v603, v603)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L2
	} else {
		goto L172
	}
L171:
	;
	goto L170
L172:
	;
	F_LockBuffer(m, v605, int32(2))
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L2
	} else {
		goto L173
	}
L173:
	;
	if v605 < int32(0) {
		goto L175
	} else {
		goto L176
	}
L174:
	;
	v628 = int32(0)
	v639 = int32(2)
	*(*uint8)(unsafe.Add(mBase, uint32(v556))) = uint8(v639)
	v643 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+65)))
	v651 = v643
	v655 = v628
	goto L179
L175:
	;
	v613 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[0]))
	v619 = *(*int32)(unsafe.Add(mBase, uint32(v613+(v605^int32(-1))<<(uint(int32(2))%32))))
	v627 = v619
	goto L174
L176:
	;
	goto L177
L177:
	;
	v621 = *(*int32)(unsafe.Add(mBase, _c_F_FlushPages[1]))
	v627 = v621 + v605<<(uint(int32(13))%32) + int32(-8192)
	goto L174
L178:
	;
	v759 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v579)+82)))
	v760 = F_PageIndexTupleOverwrite(m, v627, v759, v556, (v588+int32(7))&int32(-8))
	mBase = m.M
	v761 = m.ExcPending
	if v761 != 0 {
		goto L2
	} else {
		goto L202
	}
L179:
	;
	if v549 != 0 {
		goto L183
	} else {
		goto L184
	}
L180:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v556)+2)) = uint16(v746)
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v579)+67)))
	*(*uint8)(unsafe.Add(mBase, uint32(v556)+1)) = uint8(v757)
	goto L178
L181:
	;
	if base.B2i32(v550 <= v628) == int32(0) {
		goto L187
	} else {
		goto L188
	}
L182:
	;
	v679 = v549 + v668 - int32(1)
	goto L181
L183:
	;
	v661 = *(*int32)(unsafe.Add(mBase, uint32(v579)+72))
	v668 = *(*int32)(unsafe.Add(mBase, uint32(v549+v661+v651<<(uint(int32(2))%32)-int32(1))))
	if v668 != 0 {
		goto L182
	} else {
		goto L186
	}
L184:
	;
	goto L185
L185:
	;
	v670 = *(*int32)(unsafe.Add(mBase, uint32(v579)+72))
	v674 = *(*int32)(unsafe.Add(mBase, uint32(v670+v651<<(uint(int32(2))%32))))
	v679 = v674
	goto L181
L186:
	;
	v679 = int32(0)
	goto L181
L187:
	;
	v684 = int32(0)
	v687 = *(*int32)(unsafe.Add(mBase, uint32(v679)))
	v693 = v684
	v698 = v655
	goto L190
L188:
	;
	v746 = v655
	goto L189
L189:
	;
	if int32(0) < v651 {
		v651 = v651 - int32(1)
		v655 = v746
		goto L179
	} else {
		goto L201
	}
L190:
	;
	v706 = v556 + int32(4) + v698*int32(6)
	if v693 < v687 {
		goto L193
	} else {
		goto L194
	}
L191:
	;
	v746 = v731
	goto L189
L192:
	;
	v730 = int32(1)
	v731 = v698 + v730
	*(*uint16)(unsafe.Add(mBase, uint32(v706)+4)) = uint16(v728)
	*(*uint16)(unsafe.Add(mBase, uint32(v706)+2)) = uint16(v729)
	v735 = v693 + v730
	if v735 != v550<<(uint(base.B2i32(v651 == v684))%32) {
		v693 = v735
		v698 = v731
		goto L190
	} else {
		goto L200
	}
L193:
	;
	v710 = v679 + int32(8) + v693*int32(12)
	if v549 == int32(0) {
		goto L197
	} else {
		goto L198
	}
L194:
	;
	goto L195
L195:
	;
	v724 = int32(_a_F_FlushPages_9)
	*(*uint16)(unsafe.Add(mBase, uint32(v706))) = uint16(v724)
	v728 = int32(0)
	v729 = v724
	goto L192
L196:
	;
	v719 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v718)+80)))
	v720 = *(*int32)(unsafe.Add(mBase, uint32(v718)+76))
	v722 = int32(base.Ui32(v720) >> (uint(int32(16)) % 32))
	*(*uint16)(unsafe.Add(mBase, uint32(v706))) = uint16(v722)
	v728 = v719
	v729 = v720
	goto L192
L197:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	v718 = v713
	goto L196
L198:
	;
	goto L199
L199:
	;
	v714 = *(*int32)(unsafe.Add(mBase, uint32(v710)))
	v718 = v549 + v714 - int32(1)
	goto L196
L200:
	;
	goto L191
L201:
	;
	goto L180
L202:
	;
	if v760 == int32(0) {
		goto L18
	} else {
		goto L203
	}
L203:
	;
	F_MarkBufferDirty(m, v605)
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L2
	} else {
		goto L204
	}
L204:
	;
	F_UnlockReleaseBuffer(m, v605)
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L2
	} else {
		goto L205
	}
L205:
	;
	if v590 != 0 {
		v559 = v590
		goto L159
	} else {
		goto L206
	}
L206:
	;
	goto L160
L207:
	;
	v786 = *(*int32)(unsafe.Add(mBase, uint32(l0)+160))
	v787 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v786)+92)) = uint8(v787)
	v789 = *(*int32)(unsafe.Add(mBase, uint32(l0)+180))
	F_MemoryContextReset(m, v789)
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L2
	} else {
		goto L208
	}
L208:
	;
	m.G0 = v19 + int32(48)
	return
L209:
	;
	F_errcode(m, int32(261))
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L2
	} else {
		goto L210
	}
L210:
	;
	F_errmsg(m, int32(_a_F_FlushPages_10), int32(0))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L2
	} else {
		goto L211
	}
L211:
	;
	F_errfinish(m, int32(_a_F_FlushPages_7), int32(200), int32(_a_F_FlushPages_8))
	mBase = m.M
	v810 = m.ExcPending
	if v810 != 0 {
		goto L2
	} else {
		goto L212
	}
L212:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L213:
	;
	v815 = *(*int32)(unsafe.Add(mBase, uint32(v71)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = v815 + int32(4)
	F_errmsg_internal(m, int32(_a_F_FlushPages_6), v19+int32(32))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L2
	} else {
		goto L214
	}
L214:
	;
	F_errfinish(m, int32(_a_F_FlushPages_7), int32(226), int32(_a_F_FlushPages_8))
	mBase = m.M
	v828 = m.ExcPending
	if v828 != 0 {
		goto L2
	} else {
		goto L215
	}
L215:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L216:
	;
	v833 = *(*int32)(unsafe.Add(mBase, uint32(v552)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = v833 + int32(4)
	F_errmsg_internal(m, int32(_a_F_FlushPages_6), v19)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L2
	} else {
		goto L217
	}
L217:
	;
	F_errfinish(m, int32(_a_F_FlushPages_7), int32(290), int32(_a_F_FlushPages_11))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L2
	} else {
		goto L218
	}
L218:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_FreeDecodingContext(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
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
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v9 != 0 {
		*(*int64)(unsafe.Add(mBase, uint32(v7)+24)) = int64(0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+20)) = int32(_a_F_FreeDecodingContext_0)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = l0
		*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = int32(993)
		v17 = int32(_a_F_FreeDecodingContext_1)
		v18 = *(*int32)(unsafe.Add(mBase, _c_F_FreeDecodingContext[0]))
		*(*int32)(unsafe.Add(mBase, _c_F_FreeDecodingContext[0])) = v7 + int32(4)
		*(*int32)(unsafe.Add(mBase, uint32(v7)+4)) = v18
		*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v7 + int32(16)
		v27 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+164)) = uint8(v27)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+147)) = uint8(v27)
		m.T0[v9].(func(*base.Module, int32))(m, l0)
		mBase = m.M
		v32 = m.ExcPending
		if v32 != 0 {
			return
		} else {
			v34 = *(*int32)(unsafe.Add(mBase, uint32(v7)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_FreeDecodingContext[0])) = v34
			v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
			F_ReorderBufferFree(m, v37)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return
			} else {
				v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				F_FreeSnapshotBuilder(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return
				} else {
					v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					F_XLogReaderFree(m, v43)
					mBase = m.M
					v45 = m.ExcPending
					if v45 != 0 {
						return
					} else {
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						F_MemoryContextDelete(m, v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							m.G0 = v7 + int32(32)
							return
						}
					}
				}
			}
		}
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		F_ReorderBufferFree(m, v37)
		mBase = m.M
		v39 = m.ExcPending
		if v39 != 0 {
			return
		} else {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
			F_FreeSnapshotBuilder(m, v40)
			mBase = m.M
			v42 = m.ExcPending
			if v42 != 0 {
				return
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				F_XLogReaderFree(m, v43)
				mBase = m.M
				v45 = m.ExcPending
				if v45 != 0 {
					return
				} else {
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					F_MemoryContextDelete(m, v46)
					mBase = m.M
					v48 = m.ExcPending
					if v48 != 0 {
						return
					} else {
						m.G0 = v7 + int32(32)
						return
					}
				}
			}
		}
	}
}
func F_FreeWorkerInfo(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v59 int32
	_ = v59
	var v63 int32
	_ = v63
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_FreeWorkerInfo[0]))
	if v6 != 0 {
		v8 = *(*int32)(unsafe.Add(mBase, _c_F_FreeWorkerInfo[1]))
		v12 = F_LWLockAcquire(m, v8+int32(2816), int32(0))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, _c_F_FreeWorkerInfo[2]))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
			*(*int32)(unsafe.Add(mBase, _c_F_FreeWorkerInfo[3])) = v17
			v20 = *(*int32)(unsafe.Add(mBase, _c_F_FreeWorkerInfo[0]))
			v21 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			v22 = *(*int32)(unsafe.Add(mBase, uint32(v20)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v21)+4)) = v22
			v24 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
			*(*int32)(unsafe.Add(mBase, uint32(v22))) = v24
			v26 = int32(0)
			*(*uint8)(unsafe.Add(mBase, uint32(v20)+36)) = uint8(v26)
			v28 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v20)+8)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v20)+32)) = v26
			*(*int64)(unsafe.Add(mBase, uint32(v20)+24)) = v28
			*(*int32)(unsafe.Add(mBase, uint32(v20)+16)) = v26
			v37 = v16 + int32(12)
			v38 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
			if v38 == v26 {
				*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = int32(0)
				*(*int32)(unsafe.Add(mBase, uint32(v16)+12)) = v37
				v44 = v37
			} else {
				v44 = v38
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20))) = v37
			*(*int32)(unsafe.Add(mBase, uint32(v20)+4)) = v44
			*(*int32)(unsafe.Add(mBase, uint32(v44))) = v20
			*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v20
			v49 = *(*int32)(unsafe.Add(mBase, uint32(v16)+20))
			v50 = int32(1)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v49 + v50
			*(*int32)(unsafe.Add(mBase, _c_F_FreeWorkerInfo[0])) = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v50
			v59 = *(*int32)(unsafe.Add(mBase, _c_F_FreeWorkerInfo[1]))
			F_LWLockRelease(m, v59+int32(2816))
			mBase = m.M
			v63 = m.ExcPending
			if v63 != 0 {
				return
			} else {
				return
			}
		}
	} else {
		return
	}
}
func F___floatsitf(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int64
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v29 int64
	_ = v29
	var v37 int64
	_ = v37
	var v38 int64
	_ = v38
	var v42 int64
	_ = v42
	var v55 int64
	_ = v55
	var v57 int64
	_ = v57
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	if l1 != 0 {
		v11 = l1 >> (uint(int32(31)) % 32)
		v13 = l1 ^ v11 - v11
		v14 = base.I64_extend_i32_u(v13)
		v15 = int64(0)
		v16 = base.I32_clz(v13)
		v18 = v16 + int32(81)
		if v18&int32(64) != 0 {
			v37 = int64(0)
			v38 = v14 << (uint(base.I64_extend_i32_u(v16+int32(17))) % 64)
		} else {
			if v18 == int32(0) {
				v37 = v14
				v38 = v15
			} else {
				v29 = base.I64_extend_i32_u(v18)
				v37 = v14 << (uint(v29) % 64)
				v38 = v15<<(uint(v29)%64) | int64(base.Ui64(v14)>>(uint(base.I64_extend_i32_u(int32(64)-v18))%64))
			}
		}
		*(*int64)(unsafe.Add(mBase, uint32(v8))) = v37
		*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v38
		v42 = *(*int64)(unsafe.Add(mBase, uint32(v8)+8))
		if l1 < int32(0) {
			v55 = int64(-9223372036854775807 - 1)
		} else {
			v55 = int64(0)
		}
		v57 = *(*int64)(unsafe.Add(mBase, uint32(v8)))
		v60 = v42 ^ int64(281474976710656) + base.I64_extend_i32_u(int32(_a_F___floatsitf_0)-v16)<<(uint(int64(48))%64) | v55
		v61 = v57
	} else {
		v60 = int64(0)
		v61 = int64(0)
	}
	*(*int64)(unsafe.Add(mBase, uint32(l0))) = v61
	*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v60
	m.G0 = v8 + int32(16)
	return
}
func F___fmodeflags(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	v4 = int32(43)
	v5 = F___strchrnul(m, l0, v4)
	mBase = m.M
	v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5))))
	if v7 == v4 {
		v11 = v5
	} else {
		v11 = int32(0)
	}
	if v11 == int32(0) {
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		v17 = base.B2i32(v14 != int32(114))
	} else {
		v17 = int32(2)
	}
	v20 = int32(120)
	v21 = F___strchrnul(m, l0, v20)
	mBase = m.M
	v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v21))))
	if v23 == v20 {
		v27 = v21
	} else {
		v27 = int32(0)
	}
	if v27 != 0 {
		v28 = v17 | int32(128)
	} else {
		v28 = v17
	}
	v31 = int32(101)
	v32 = F___strchrnul(m, l0, v31)
	mBase = m.M
	v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v32))))
	if v34 == v31 {
		v38 = v32
	} else {
		v38 = int32(0)
	}
	if v38 != 0 {
		v39 = v28 | int32(_a_F___fmodeflags_0)
	} else {
		v39 = v28
	}
	v42 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v42 == int32(114) {
		v45 = v39
	} else {
		v45 = v39 | int32(64)
	}
	if v42 == int32(119) {
		v50 = v45 | int32(512)
	} else {
		v50 = v45
	}
	if v42 == int32(97) {
		v55 = v50 | int32(1024)
	} else {
		v55 = v50
	}
	return v55
}
func F___ftello_unlocked(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v34 int64
	_ = v34
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v6 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
	if v6&int32(128) != 0 {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v11 == v12 {
			v14 = int32(1)
		} else {
			v14 = int32(2)
		}
		v16 = v14
	} else {
		v16 = int32(1)
	}
	v17 = m.T0[v4].(func(*base.Module, int32, int64, int32) int64)(m, l0, int64(0), v16)
	mBase = m.M
	if v17 < int64(0) {
		v34 = v17
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
		if v20 != 0 {
			v26 = v20
			v27 = int32(4)
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+l0)))
			v34 = v17 + base.I64_extend_i32_s(v29-v26)
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v22 == int32(0) {
				v34 = v17
			} else {
				v26 = v22
				v27 = int32(20)
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v27+l0)))
				v34 = v17 + base.I64_extend_i32_s(v29-v26)
			}
		}
	}
	return v34
}
func F__fmt(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	var v43 int32
	_ = v43
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v76 int32
	_ = v76
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v178 int32
	_ = v178
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v220 int32
	_ = v220
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v269 int32
	_ = v269
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v328 int32
	_ = v328
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int32
	_ = v337
	var v340 int32
	_ = v340
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v349 int32
	_ = v349
	var v361 int32
	_ = v361
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v372 int32
	_ = v372
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v386 int32
	_ = v386
	var v388 int32
	_ = v388
	var v400 int32
	_ = v400
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v437 int32
	_ = v437
	var v441 int32
	_ = v441
	var v444 int32
	_ = v444
	var v448 int32
	_ = v448
	var v451 int32
	_ = v451
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v460 int32
	_ = v460
	var v472 int32
	_ = v472
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v484 int32
	_ = v484
	var v485 int32
	_ = v485
	var v487 int32
	_ = v487
	var v490 int32
	_ = v490
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v511 int32
	_ = v511
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v525 int32
	_ = v525
	var v529 int32
	_ = v529
	var v530 int32
	_ = v530
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v546 int32
	_ = v546
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v557 int32
	_ = v557
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v569 int32
	_ = v569
	var v571 int32
	_ = v571
	var v583 int32
	_ = v583
	var v587 int32
	_ = v587
	var v590 int32
	_ = v590
	var v595 int32
	_ = v595
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v613 int32
	_ = v613
	var v615 int32
	_ = v615
	var v627 int32
	_ = v627
	var v631 int32
	_ = v631
	var v634 int32
	_ = v634
	var v639 int32
	_ = v639
	var v640 int32
	_ = v640
	var v644 int32
	_ = v644
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v660 int32
	_ = v660
	var v672 int32
	_ = v672
	var v676 int32
	_ = v676
	var v679 int32
	_ = v679
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v689 int32
	_ = v689
	var v692 int32
	_ = v692
	var v696 int32
	_ = v696
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v703 int32
	_ = v703
	var v706 int32
	_ = v706
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v730 int32
	_ = v730
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v741 int32
	_ = v741
	var v743 int32
	_ = v743
	var v746 int32
	_ = v746
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v753 int32
	_ = v753
	var v755 int32
	_ = v755
	var v767 int32
	_ = v767
	var v771 int32
	_ = v771
	var v774 int32
	_ = v774
	var v778 int32
	_ = v778
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v784 int32
	_ = v784
	var v789 int32
	_ = v789
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v805 int32
	_ = v805
	var v807 int32
	_ = v807
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v814 int32
	_ = v814
	var v818 int32
	_ = v818
	var v819 int32
	_ = v819
	var v821 int32
	_ = v821
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v829 int32
	_ = v829
	var v832 int32
	_ = v832
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v852 int32
	_ = v852
	var v858 int32
	_ = v858
	var v859 int32
	_ = v859
	var v861 int32
	_ = v861
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v872 int32
	_ = v872
	var v874 int32
	_ = v874
	var v877 int32
	_ = v877
	var v880 int32
	_ = v880
	var v884 int32
	_ = v884
	var v885 int32
	_ = v885
	var v887 int32
	_ = v887
	var v889 int32
	_ = v889
	var v901 int32
	_ = v901
	var v905 int32
	_ = v905
	var v908 int32
	_ = v908
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
	var v924 int32
	_ = v924
	var v936 int32
	_ = v936
	var v940 int32
	_ = v940
	var v943 int32
	_ = v943
	var v948 int32
	_ = v948
	var v949 int32
	_ = v949
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v971 int32
	_ = v971
	var v972 int32
	_ = v972
	var v973 int32
	_ = v973
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v996 int32
	_ = v996
	var v1000 int32
	_ = v1000
	var v1001 int32
	_ = v1001
	var v1003 int32
	_ = v1003
	var v1005 int32
	_ = v1005
	var v1017 int32
	_ = v1017
	var v1021 int32
	_ = v1021
	var v1024 int32
	_ = v1024
	var v1028 int32
	_ = v1028
	var v1030 int32
	_ = v1030
	var v1031 int32
	_ = v1031
	var v1034 int32
	_ = v1034
	var v1039 int32
	_ = v1039
	var v1041 int32
	_ = v1041
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1060 int32
	_ = v1060
	var v1064 int32
	_ = v1064
	var v1069 int32
	_ = v1069
	var v1072 int32
	_ = v1072
	var v1077 int32
	_ = v1077
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1088 int32
	_ = v1088
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1104 int32
	_ = v1104
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1128 int32
	_ = v1128
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1134 int32
	_ = v1134
	var v1138 int32
	_ = v1138
	var v1142 int32
	_ = v1142
	var v1143 int32
	_ = v1143
	var v1145 int32
	_ = v1145
	var v1147 int32
	_ = v1147
	var v1159 int32
	_ = v1159
	var v1163 int32
	_ = v1163
	var v1166 int32
	_ = v1166
	var v1171 int32
	_ = v1171
	var v1172 int32
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1192 int32
	_ = v1192
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1222 int32
	_ = v1222
	var v1226 int32
	_ = v1226
	var v1227 int32
	_ = v1227
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1243 int32
	_ = v1243
	var v1247 int32
	_ = v1247
	var v1250 int32
	_ = v1250
	var v1256 int32
	_ = v1256
	var v1257 int32
	_ = v1257
	var v1259 int32
	_ = v1259
	var v1267 int32
	_ = v1267
	var v1270 int32
	_ = v1270
	var v1272 int32
	_ = v1272
	var v1273 int32
	_ = v1273
	var v1278 int32
	_ = v1278
	var v1290 int32
	_ = v1290
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1296 int32
	_ = v1296
	var v1301 int32
	_ = v1301
	var v1305 int32
	_ = v1305
	var v1306 int32
	_ = v1306
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1322 int32
	_ = v1322
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1333 int32
	_ = v1333
	var v1334 int32
	_ = v1334
	var v1337 int32
	_ = v1337
	var v1341 int32
	_ = v1341
	var v1343 int32
	_ = v1343
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1348 int32
	_ = v1348
	var v1360 int32
	_ = v1360
	var v1364 int32
	_ = v1364
	var v1367 int32
	_ = v1367
	var v1373 int32
	_ = v1373
	var v1378 int32
	_ = v1378
	v15 = m.G0
	v17 = v15 - int32(304)
	m.G0 = v17
	v19 = l0
	v21 = l2
	goto L1
L1:
	;
	v33 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
	if v33 == int32(37) {
		goto L8
	} else {
		goto L9
	}
L3:
	;
	v19 = v1378 + int32(1)
	v21 = v1373
	goto L1
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17))) = v1337
	v1341 = v17 + int32(292)
	v1343 = F_pg_sprintf(m, v1341, int32(_a_F__fmt_0), v17)
	mBase = m.M
	v1344 = m.ExcPending
	if v1344 != 0 {
		goto L97
	} else {
		goto L335
	}
L5:
	;
	v1217 = v52 - int32(86)
	if v1217 != 0 {
		goto L306
	} else {
		goto L307
	}
L6:
	;
	m.G0 = v17 + int32(304)
	return v21
L7:
	;
	if v21 == l3 {
		goto L6
	} else {
		goto L303
	}
L8:
	;
	v43 = v19
	goto L48
L9:
	;
	goto L10
L10:
	;
	if v33 == int32(0) {
		goto L6
	} else {
		goto L302
	}
L11:
	;
	v1177 = v51
	goto L7
L12:
	;
	v1171 = F__fmt(m, int32(_a_F__fmt_1), l1, v21, l3, l4)
	mBase = m.M
	v1172 = m.ExcPending
	if v1172 != 0 {
		goto L97
	} else {
		goto L301
	}
L13:
	;
	v1064 = *(*int32)(unsafe.Add(mBase, uint32(l1)+32))
	if v1064 < int32(0) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L274
	}
L14:
	;
	v1034 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if base.B2i32(v1034 == int32(0))|base.B2i32(base.Ui32(l3) <= base.Ui32(v21)) != 0 {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L269
	}
L15:
	;
	v1028 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v1030 = F__yconv(m, v1028, int32(1900), v21, l3)
	mBase = m.M
	v1031 = m.ExcPending
	if v1031 != 0 {
		goto L97
	} else {
		goto L268
	}
L16:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(3)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v972 = int32(100)
	v973 = base.I32_rem_s(v971, v972)
	if int32(0) < v973 {
		goto L253
	} else {
		goto L254
	}
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+292)) = int32(1)
	v957 = F__fmt(m, int32(_a_F__fmt_2), l1, v21, l3, v17+int32(292))
	mBase = m.M
	v958 = m.ExcPending
	if v958 != 0 {
		goto L97
	} else {
		goto L248
	}
L18:
	;
	v948 = F__fmt(m, int32(_a_F__fmt_3), l1, v21, l3, l4)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L97
	} else {
		goto L247
	}
L19:
	;
	v912 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+256)) = v912
	v915 = v17 + int32(292)
	v919 = F_pg_sprintf(m, v915, int32(_a_F__fmt_4), v17+int32(256))
	mBase = m.M
	v920 = m.ExcPending
	if v920 != 0 {
		goto L97
	} else {
		goto L241
	}
L20:
	;
	v867 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v869 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v869 != 0 {
		goto L232
	} else {
		goto L233
	}
L21:
	;
	v863 = F__fmt(m, int32(_a_F__fmt_5), l1, v21, l3, l4)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L97
	} else {
		goto L231
	}
L22:
	;
	v778 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v780 = base.I32_rem_s(v778, int32(400))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v782 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v784 = v782
	v789 = int32(1900)
	goto L205
L23:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if v741 != 0 {
		goto L196
	} else {
		goto L197
	}
L24:
	;
	v700 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	v701 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	v703 = int32(7)
	v706 = base.I32_div_s(v700-v701+v703, v703)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+176)) = v706
	v709 = v17 + int32(292)
	v713 = F_pg_sprintf(m, v709, int32(_a_F__fmt_0), v17+int32(176))
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L97
	} else {
		goto L190
	}
L25:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L186
	}
L26:
	;
	v684 = F__fmt(m, int32(_a_F__fmt_3), l1, v21, l3, l4)
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L97
	} else {
		goto L185
	}
L27:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+160)) = v648
	v651 = v17 + int32(292)
	v655 = F_pg_sprintf(m, v651, int32(_a_F__fmt_0), v17+int32(160))
	mBase = m.M
	v656 = m.ExcPending
	if v656 != 0 {
		goto L97
	} else {
		goto L179
	}
L28:
	;
	v644 = F__fmt(m, int32(_a_F__fmt_6), l1, v21, l3, l4)
	mBase = m.M
	v645 = m.ExcPending
	if v645 != 0 {
		goto L97
	} else {
		goto L178
	}
L29:
	;
	v639 = F__fmt(m, int32(_a_F__fmt_7), l1, v21, l3, l4)
	mBase = m.M
	v640 = m.ExcPending
	if v640 != 0 {
		goto L97
	} else {
		goto L177
	}
L30:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L169
	}
L31:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L165
	}
L32:
	;
	v557 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+144)) = v557 + int32(1)
	v562 = v17 + int32(292)
	v566 = F_pg_sprintf(m, v562, int32(_a_F__fmt_0), v17+int32(144))
	mBase = m.M
	v567 = m.ExcPending
	if v567 != 0 {
		goto L97
	} else {
		goto L159
	}
L33:
	;
	v522 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+128)) = v522
	v525 = v17 + int32(292)
	v529 = F_pg_sprintf(m, v525, int32(_a_F__fmt_0), v17+int32(128))
	mBase = m.M
	v530 = m.ExcPending
	if v530 != 0 {
		goto L97
	} else {
		goto L153
	}
L34:
	;
	v483 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v484 = int32(12)
	v485 = base.I32_rem_s(v483, v484)
	if v485 != 0 {
		goto L144
	} else {
		goto L145
	}
L35:
	;
	v448 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+96)) = v448
	v451 = v17 + int32(292)
	v455 = F_pg_sprintf(m, v451, int32(_a_F__fmt_8), v17+int32(96))
	mBase = m.M
	v456 = m.ExcPending
	if v456 != 0 {
		goto L97
	} else {
		goto L138
	}
L36:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+80)) = v411 + int32(1)
	v416 = v17 + int32(292)
	v420 = F_pg_sprintf(m, v416, int32(_a_F__fmt_9), v17+int32(80))
	mBase = m.M
	v421 = m.ExcPending
	if v421 != 0 {
		goto L97
	} else {
		goto L132
	}
L37:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	v373 = int32(12)
	v374 = base.I32_rem_s(v372, v373)
	if v374 != 0 {
		goto L123
	} else {
		goto L124
	}
L38:
	;
	v337 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v337
	v340 = v17 + int32(292)
	v344 = F_pg_sprintf(m, v340, int32(_a_F__fmt_0), v17+int32(48))
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L97
	} else {
		goto L117
	}
L39:
	;
	v333 = F__fmt(m, int32(_a_F__fmt_10), l1, v21, l3, l4)
	mBase = m.M
	v334 = m.ExcPending
	if v334 != 0 {
		goto L97
	} else {
		goto L116
	}
L40:
	;
	v297 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v297
	v300 = v17 + int32(292)
	v304 = F_pg_sprintf(m, v300, int32(_a_F__fmt_8), v17+int32(32))
	mBase = m.M
	v305 = m.ExcPending
	if v305 != 0 {
		goto L97
	} else {
		goto L110
	}
L41:
	;
	v262 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v17)+16)) = v262
	v265 = v17 + int32(292)
	v269 = F_pg_sprintf(m, v265, int32(_a_F__fmt_0), v17+int32(16))
	mBase = m.M
	v270 = m.ExcPending
	if v270 != 0 {
		goto L97
	} else {
		goto L104
	}
L42:
	;
	v258 = F__fmt(m, int32(_a_F__fmt_2), l1, v21, l3, l4)
	mBase = m.M
	v259 = m.ExcPending
	if v259 != 0 {
		goto L97
	} else {
		goto L103
	}
L43:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+292)) = int32(1)
	v243 = F__fmt(m, int32(_a_F__fmt_11), l1, v21, l3, v17+int32(292))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L97
	} else {
		goto L98
	}
L44:
	;
	v189 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	v190 = int32(100)
	v191 = base.I32_div_s(v189, v190)
	v194 = v189 - v191*v190
	v197 = int32(0)
	if base.B2i32(v189 < int32(-1899))|base.B2i32(v197 <= v194) == v197 {
		goto L83
	} else {
		goto L84
	}
L45:
	;
	v156 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v156) <= base.Ui32(int32(11)) {
		goto L75
	} else {
		goto L76
	}
L46:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(l1)+16))
	if base.Ui32(v122) <= base.Ui32(int32(11)) {
		goto L67
	} else {
		goto L68
	}
L47:
	;
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui32(v88) <= base.Ui32(int32(6)) {
		goto L59
	} else {
		goto L60
	}
L48:
	;
	v51 = v43 + int32(1)
	v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+1)))
	switch v52 {
	case 0:
		v1177 = v43
		goto L7
	default:
		goto L11
	case 43:
		goto L12
	case 65:
		goto L50
	case 66:
		goto L46
	case 67:
		goto L44
	case 68:
		goto L42
	case 69, 79:
		v43 = v51
		goto L48
	case 70:
		goto L39
	case 71, 86, 103:
		goto L22
	case 72:
		goto L38
	case 73:
		goto L37
	case 77:
		goto L33
	case 82:
		goto L29
	case 83:
		goto L27
	case 84:
		goto L26
	case 85:
		goto L24
	case 87:
		goto L20
	case 88:
		goto L18
	case 89:
		goto L15
	case 90:
		goto L14
	case 97:
		goto L47
	case 98, 104:
		goto L45
	case 99:
		goto L43
	case 100:
		goto L41
	case 101:
		goto L40
	case 106:
		goto L36
	case 107:
		goto L35
	case 108:
		goto L34
	case 109:
		goto L32
	case 110:
		goto L31
	case 112:
		goto L30
	case 114:
		goto L28
	case 116:
		goto L25
	case 117:
		goto L23
	case 118:
		goto L21
	case 119:
		goto L19
	case 120:
		goto L17
	case 121:
		goto L16
	case 122:
		goto L13
	}
L49:
	;
	v54 = *(*int32)(unsafe.Add(mBase, uint32(l1)+24))
	if base.Ui32(v54) <= base.Ui32(int32(6)) {
		goto L51
	} else {
		goto L52
	}
L50:
	;
	goto L49
L51:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v54<<(uint(int32(2))%32))+uint32(_c_F__fmt[0])))
	v60 = v59
	goto L53
L52:
	;
	v60 = int32(_a_F__fmt_12)
	goto L53
L53:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L54
	}
L54:
	;
	v62 = v60
	v64 = v21
	goto L55
L55:
	;
	v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62))))
	*(*uint8)(unsafe.Add(mBase, uint32(v64))) = uint8(v76)
	if v76 == int32(0) {
		v1373 = v64
		v1378 = v51
		goto L3
	} else {
		goto L57
	}
L56:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L57:
	;
	v80 = int32(1)
	v83 = v64 + v80
	if v83 != l3 {
		v62 = v62 + v80
		v64 = v83
		goto L55
	} else {
		goto L58
	}
L58:
	;
	goto L56
L59:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v88<<(uint(int32(2))%32))+uint32(_c_F__fmt[1])))
	v94 = v93
	goto L61
L60:
	;
	v94 = int32(_a_F__fmt_12)
	goto L61
L61:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L62
	}
L62:
	;
	v96 = v94
	v98 = v21
	goto L63
L63:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v96))))
	*(*uint8)(unsafe.Add(mBase, uint32(v98))) = uint8(v110)
	if v110 == int32(0) {
		v1373 = v98
		v1378 = v51
		goto L3
	} else {
		goto L65
	}
L64:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L65:
	;
	v114 = int32(1)
	v117 = v98 + v114
	if v117 != l3 {
		v96 = v96 + v114
		v98 = v117
		goto L63
	} else {
		goto L66
	}
L66:
	;
	goto L64
L67:
	;
	v127 = *(*int32)(unsafe.Add(mBase, uint32(v122<<(uint(int32(2))%32))+uint32(_c_F__fmt[2])))
	v128 = v127
	goto L69
L68:
	;
	v128 = int32(_a_F__fmt_12)
	goto L69
L69:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L70
	}
L70:
	;
	v130 = v128
	v132 = v21
	goto L71
L71:
	;
	v144 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v130))))
	*(*uint8)(unsafe.Add(mBase, uint32(v132))) = uint8(v144)
	if v144 == int32(0) {
		v1373 = v132
		v1378 = v51
		goto L3
	} else {
		goto L73
	}
L72:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L73:
	;
	v148 = int32(1)
	v151 = v132 + v148
	if v151 != l3 {
		v130 = v130 + v148
		v132 = v151
		goto L71
	} else {
		goto L74
	}
L74:
	;
	goto L72
L75:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v156<<(uint(int32(2))%32))+uint32(_c_F__fmt[3])))
	v162 = v161
	goto L77
L76:
	;
	v162 = int32(_a_F__fmt_12)
	goto L77
L77:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L78
	}
L78:
	;
	v164 = v162
	v166 = v21
	goto L79
L79:
	;
	v178 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v164))))
	*(*uint8)(unsafe.Add(mBase, uint32(v166))) = uint8(v178)
	if v178 == int32(0) {
		v1373 = v166
		v1378 = v51
		goto L3
	} else {
		goto L81
	}
L80:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L81:
	;
	v182 = int32(1)
	v185 = v166 + v182
	if v185 != l3 {
		v164 = v164 + v182
		v166 = v185
		goto L79
	} else {
		goto L82
	}
L82:
	;
	goto L80
L83:
	;
	v1337 = v191 + int32(18)
	goto L4
L84:
	;
	goto L85
L85:
	;
	v210 = base.B2i32(v189 < int32(-1999)) & base.B2i32(int32(0) < v194)
	if v210 != 0 {
		goto L86
	} else {
		goto L87
	}
L86:
	;
	v211 = int32(20)
	goto L88
L87:
	;
	v211 = int32(19)
	goto L88
L88:
	;
	v212 = v191 + v211
	v213 = int32(0)
	if v212|base.B2i32(v210 == v213)&base.B2i32(v213 <= v194) != 0 {
		v1337 = v212
		goto L4
	} else {
		goto L89
	}
L89:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L90
	}
L90:
	;
	v220 = int32(45)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v220)
	if l3 == v21+int32(1) {
		goto L91
	} else {
		goto L92
	}
L91:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L92:
	;
	goto L93
L93:
	;
	v227 = int32(48)
	*(*uint8)(unsafe.Add(mBase, uint32(v21)+1)) = uint8(v227)
	v230 = v21 + int32(2)
	if l3 == v230 {
		goto L94
	} else {
		goto L95
	}
L94:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L95:
	;
	goto L96
L96:
	;
	v234 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v230))) = uint8(v234)
	v19 = v43 + int32(2)
	v21 = v230
	goto L1
L97:
	;
	return int32(0)
L98:
	;
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v17)+292))
	if v248 == int32(3) {
		goto L99
	} else {
		goto L100
	}
L99:
	;
	v251 = int32(2)
	goto L101
L100:
	;
	v251 = v248
	goto L101
L101:
	;
	v252 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if base.Ui32(v251) <= base.Ui32(v252) {
		v1373 = v243
		v1378 = v51
		goto L3
	} else {
		goto L102
	}
L102:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v251
	v19 = v43 + int32(2)
	v21 = v243
	goto L1
L103:
	;
	v19 = v43 + int32(2)
	v21 = v258
	goto L1
L104:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L105
	}
L105:
	;
	v272 = v265
	v274 = v21
	goto L106
L106:
	;
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v272))))
	*(*uint8)(unsafe.Add(mBase, uint32(v274))) = uint8(v286)
	if v286 == int32(0) {
		v1373 = v274
		v1378 = v51
		goto L3
	} else {
		goto L108
	}
L107:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L108:
	;
	v290 = int32(1)
	v293 = v274 + v290
	if v293 != l3 {
		v272 = v272 + v290
		v274 = v293
		goto L106
	} else {
		goto L109
	}
L109:
	;
	goto L107
L110:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L111
	}
L111:
	;
	v307 = v300
	v309 = v21
	goto L112
L112:
	;
	v321 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v307))))
	*(*uint8)(unsafe.Add(mBase, uint32(v309))) = uint8(v321)
	if v321 == int32(0) {
		v1373 = v309
		v1378 = v51
		goto L3
	} else {
		goto L114
	}
L113:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L114:
	;
	v325 = int32(1)
	v328 = v309 + v325
	if v328 != l3 {
		v307 = v307 + v325
		v309 = v328
		goto L112
	} else {
		goto L115
	}
L115:
	;
	goto L113
L116:
	;
	v19 = v43 + int32(2)
	v21 = v333
	goto L1
L117:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L118
	}
L118:
	;
	v347 = v340
	v349 = v21
	goto L119
L119:
	;
	v361 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347))))
	*(*uint8)(unsafe.Add(mBase, uint32(v349))) = uint8(v361)
	if v361 == int32(0) {
		v1373 = v349
		v1378 = v51
		goto L3
	} else {
		goto L121
	}
L120:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L121:
	;
	v365 = int32(1)
	v368 = v349 + v365
	if v368 != l3 {
		v347 = v347 + v365
		v349 = v368
		goto L119
	} else {
		goto L122
	}
L122:
	;
	goto L120
L123:
	;
	v376 = v374
	goto L125
L124:
	;
	v376 = v373
	goto L125
L125:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+64)) = v376
	v379 = v17 + int32(292)
	v383 = F_pg_sprintf(m, v379, int32(_a_F__fmt_0), v17-int32(-64))
	mBase = m.M
	v384 = m.ExcPending
	if v384 != 0 {
		goto L97
	} else {
		goto L126
	}
L126:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L127
	}
L127:
	;
	v386 = v379
	v388 = v21
	goto L128
L128:
	;
	v400 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v386))))
	*(*uint8)(unsafe.Add(mBase, uint32(v388))) = uint8(v400)
	if v400 == int32(0) {
		v1373 = v388
		v1378 = v51
		goto L3
	} else {
		goto L130
	}
L129:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L130:
	;
	v404 = int32(1)
	v407 = v388 + v404
	if v407 != l3 {
		v386 = v386 + v404
		v388 = v407
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L133
	}
L133:
	;
	v423 = v416
	v425 = v21
	goto L134
L134:
	;
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v423))))
	*(*uint8)(unsafe.Add(mBase, uint32(v425))) = uint8(v437)
	if v437 == int32(0) {
		v1373 = v425
		v1378 = v51
		goto L3
	} else {
		goto L136
	}
L135:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L136:
	;
	v441 = int32(1)
	v444 = v425 + v441
	if v444 != l3 {
		v423 = v423 + v441
		v425 = v444
		goto L134
	} else {
		goto L137
	}
L137:
	;
	goto L135
L138:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L139
	}
L139:
	;
	v458 = v451
	v460 = v21
	goto L140
L140:
	;
	v472 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v458))))
	*(*uint8)(unsafe.Add(mBase, uint32(v460))) = uint8(v472)
	if v472 == int32(0) {
		v1373 = v460
		v1378 = v51
		goto L3
	} else {
		goto L142
	}
L141:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L142:
	;
	v476 = int32(1)
	v479 = v460 + v476
	if v479 != l3 {
		v458 = v458 + v476
		v460 = v479
		goto L140
	} else {
		goto L143
	}
L143:
	;
	goto L141
L144:
	;
	v487 = v485
	goto L146
L145:
	;
	v487 = v484
	goto L146
L146:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+112)) = v487
	v490 = v17 + int32(292)
	v494 = F_pg_sprintf(m, v490, int32(_a_F__fmt_8), v17+int32(112))
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L97
	} else {
		goto L147
	}
L147:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L148
	}
L148:
	;
	v497 = v490
	v499 = v21
	goto L149
L149:
	;
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v497))))
	*(*uint8)(unsafe.Add(mBase, uint32(v499))) = uint8(v511)
	if v511 == int32(0) {
		v1373 = v499
		v1378 = v51
		goto L3
	} else {
		goto L151
	}
L150:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L151:
	;
	v515 = int32(1)
	v518 = v499 + v515
	if v518 != l3 {
		v497 = v497 + v515
		v499 = v518
		goto L149
	} else {
		goto L152
	}
L152:
	;
	goto L150
L153:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L154
	}
L154:
	;
	v532 = v525
	v534 = v21
	goto L155
L155:
	;
	v546 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v532))))
	*(*uint8)(unsafe.Add(mBase, uint32(v534))) = uint8(v546)
	if v546 == int32(0) {
		v1373 = v534
		v1378 = v51
		goto L3
	} else {
		goto L157
	}
L156:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L157:
	;
	v550 = int32(1)
	v553 = v534 + v550
	if v553 != l3 {
		v532 = v532 + v550
		v534 = v553
		goto L155
	} else {
		goto L158
	}
L158:
	;
	goto L156
L159:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L160
	}
L160:
	;
	v569 = v562
	v571 = v21
	goto L161
L161:
	;
	v583 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v569))))
	*(*uint8)(unsafe.Add(mBase, uint32(v571))) = uint8(v583)
	if v583 == int32(0) {
		v1373 = v571
		v1378 = v51
		goto L3
	} else {
		goto L163
	}
L162:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L163:
	;
	v587 = int32(1)
	v590 = v571 + v587
	if v590 != l3 {
		v569 = v569 + v587
		v571 = v590
		goto L161
	} else {
		goto L164
	}
L164:
	;
	goto L162
L165:
	;
	v595 = int32(10)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v595)
	v598 = v21 + int32(1)
	if l3 == v598 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L167:
	;
	goto L168
L168:
	;
	v602 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v598))) = uint8(v602)
	v19 = v43 + int32(2)
	v21 = v598
	goto L1
L169:
	;
	v609 = *(*int32)(unsafe.Add(mBase, uint32(l1)+8))
	if int32(11) < v609 {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	v612 = int32(_a_F__fmt_13)
	goto L172
L171:
	;
	v612 = int32(_a_F__fmt_14)
	goto L172
L172:
	;
	v613 = v612
	v615 = v21
	goto L173
L173:
	;
	v627 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v613))))
	*(*uint8)(unsafe.Add(mBase, uint32(v615))) = uint8(v627)
	if v627 == int32(0) {
		v1373 = v615
		v1378 = v51
		goto L3
	} else {
		goto L175
	}
L174:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L175:
	;
	v631 = int32(1)
	v634 = v615 + v631
	if v634 != l3 {
		v613 = v613 + v631
		v615 = v634
		goto L173
	} else {
		goto L176
	}
L176:
	;
	goto L174
L177:
	;
	v19 = v43 + int32(2)
	v21 = v639
	goto L1
L178:
	;
	v19 = v43 + int32(2)
	v21 = v644
	goto L1
L179:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L180
	}
L180:
	;
	v658 = v651
	v660 = v21
	goto L181
L181:
	;
	v672 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v658))))
	*(*uint8)(unsafe.Add(mBase, uint32(v660))) = uint8(v672)
	if v672 == int32(0) {
		v1373 = v660
		v1378 = v51
		goto L3
	} else {
		goto L183
	}
L182:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L183:
	;
	v676 = int32(1)
	v679 = v660 + v676
	if v679 != l3 {
		v658 = v658 + v676
		v660 = v679
		goto L181
	} else {
		goto L184
	}
L184:
	;
	goto L182
L185:
	;
	v19 = v43 + int32(2)
	v21 = v684
	goto L1
L186:
	;
	v689 = int32(9)
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v689)
	v692 = v21 + int32(1)
	if l3 == v692 {
		goto L187
	} else {
		goto L188
	}
L187:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L188:
	;
	goto L189
L189:
	;
	v696 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v692))) = uint8(v696)
	v19 = v43 + int32(2)
	v21 = v692
	goto L1
L190:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L191
	}
L191:
	;
	v716 = v709
	v718 = v21
	goto L192
L192:
	;
	v730 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v716))))
	*(*uint8)(unsafe.Add(mBase, uint32(v718))) = uint8(v730)
	if v730 == int32(0) {
		v1373 = v718
		v1378 = v51
		goto L3
	} else {
		goto L194
	}
L193:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L194:
	;
	v734 = int32(1)
	v737 = v718 + v734
	if v737 != l3 {
		v716 = v716 + v734
		v718 = v737
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	v743 = v741
	goto L198
L197:
	;
	v743 = int32(7)
	goto L198
L198:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+192)) = v743
	v746 = v17 + int32(292)
	v750 = F_pg_sprintf(m, v746, int32(_a_F__fmt_4), v17+int32(192))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L97
	} else {
		goto L199
	}
L199:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L200
	}
L200:
	;
	v753 = v746
	v755 = v21
	goto L201
L201:
	;
	v767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753))))
	*(*uint8)(unsafe.Add(mBase, uint32(v755))) = uint8(v767)
	if v767 == int32(0) {
		v1373 = v755
		v1378 = v51
		goto L3
	} else {
		goto L203
	}
L202:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L203:
	;
	v771 = int32(1)
	v774 = v755 + v771
	if v774 != l3 {
		v753 = v753 + v771
		v755 = v774
		goto L201
	} else {
		goto L204
	}
L204:
	;
	goto L202
L205:
	;
	v800 = base.I32_rem_s(v789, int32(400))
	v801 = v800 + v780
	if v801&int32(3) != 0 {
		v814 = int32(365)
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v818 = int32(7)
	v819 = base.I32_rem_s(v784-v781+int32(11), v818)
	v821 = v819 - int32(3)
	v823 = base.I32_rem_u_s(v814, v818)
	v824 = v821 - v823
	if v824 < int32(-3) {
		goto L213
	} else {
		goto L214
	}
L208:
	;
	v805 = base.I32_extend16_s(v801)
	v807 = base.I32_rem_s(v805, int32(100))
	if v807 != 0 {
		v814 = int32(366)
		goto L207
	} else {
		goto L209
	}
L209:
	;
	v811 = base.I32_rem_s(v805, int32(400))
	if v811 != 0 {
		goto L210
	} else {
		goto L211
	}
L210:
	;
	v812 = int32(365)
	goto L212
L211:
	;
	v812 = int32(366)
	goto L212
L212:
	;
	v814 = v812
	goto L207
L213:
	;
	v829 = v824 + v818
	goto L215
L214:
	;
	v829 = v824
	goto L215
L215:
	;
	if v814+v829 <= v784 {
		goto L216
	} else {
		goto L217
	}
L216:
	;
	v832 = int32(1)
	v1214 = v832
	v1215 = v789 + v832
	goto L5
L217:
	;
	goto L218
L218:
	;
	if v821 <= v784 {
		goto L219
	} else {
		goto L220
	}
L219:
	;
	v838 = base.I32_div_u_s(v784-v821, int32(7))
	v1214 = v838 + int32(1)
	v1215 = v789
	goto L5
L220:
	;
	goto L221
L221:
	;
	v842 = v789 - int32(1)
	v844 = base.I32_rem_s(v842, int32(400))
	v845 = v844 + v780
	if v845&int32(3) != 0 {
		goto L222
	} else {
		goto L223
	}
L222:
	;
	v784 = v784 + int32(365)
	v789 = v842
	goto L205
L223:
	;
	goto L224
L224:
	;
	v850 = base.I32_extend16_s(v845)
	v852 = base.I32_rem_s(v850, int32(100))
	if v852 != 0 {
		goto L225
	} else {
		goto L226
	}
L225:
	;
	v861 = v784 + int32(366)
	goto L227
L226:
	;
	v858 = base.I32_rem_s(v850, int32(400))
	if v858 != 0 {
		goto L228
	} else {
		goto L229
	}
L227:
	;
	v784 = v861
	v789 = v842
	goto L205
L228:
	;
	v859 = int32(365)
	goto L230
L229:
	;
	v859 = int32(366)
	goto L230
L230:
	;
	v861 = v859 + v784
	goto L227
L231:
	;
	v19 = v43 + int32(2)
	v21 = v863
	goto L1
L232:
	;
	v872 = int32(1) - v869
	goto L234
L233:
	;
	v872 = int32(-6)
	goto L234
L234:
	;
	v874 = int32(7)
	v877 = base.I32_div_s(v867+v872+v874, v874)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+240)) = v877
	v880 = v17 + int32(292)
	v884 = F_pg_sprintf(m, v880, int32(_a_F__fmt_0), v17+int32(240))
	mBase = m.M
	v885 = m.ExcPending
	if v885 != 0 {
		goto L97
	} else {
		goto L235
	}
L235:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L236
	}
L236:
	;
	v887 = v880
	v889 = v21
	goto L237
L237:
	;
	v901 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v887))))
	*(*uint8)(unsafe.Add(mBase, uint32(v889))) = uint8(v901)
	if v901 == int32(0) {
		v1373 = v889
		v1378 = v51
		goto L3
	} else {
		goto L239
	}
L238:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L239:
	;
	v905 = int32(1)
	v908 = v889 + v905
	if v908 != l3 {
		v887 = v887 + v905
		v889 = v908
		goto L237
	} else {
		goto L240
	}
L240:
	;
	goto L238
L241:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L242
	}
L242:
	;
	v922 = v915
	v924 = v21
	goto L243
L243:
	;
	v936 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v922))))
	*(*uint8)(unsafe.Add(mBase, uint32(v924))) = uint8(v936)
	if v936 == int32(0) {
		v1373 = v924
		v1378 = v51
		goto L3
	} else {
		goto L245
	}
L244:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L245:
	;
	v940 = int32(1)
	v943 = v924 + v940
	if v943 != l3 {
		v922 = v922 + v940
		v924 = v943
		goto L243
	} else {
		goto L246
	}
L246:
	;
	goto L244
L247:
	;
	v19 = v43 + int32(2)
	v21 = v948
	goto L1
L248:
	;
	v960 = *(*int32)(unsafe.Add(mBase, uint32(v17)+292))
	if v960 == int32(3) {
		goto L249
	} else {
		goto L250
	}
L249:
	;
	v963 = int32(2)
	goto L251
L250:
	;
	v963 = v960
	goto L251
L251:
	;
	v964 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	if base.Ui32(v963) <= base.Ui32(v964) {
		v1373 = v957
		v1378 = v51
		goto L3
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v963
	v19 = v43 + int32(2)
	v21 = v957
	goto L1
L253:
	;
	v978 = v973 - v972
	goto L255
L254:
	;
	v978 = v973
	goto L255
L255:
	;
	if v971 < int32(-1999) {
		goto L256
	} else {
		goto L257
	}
L256:
	;
	v981 = v978
	goto L258
L257:
	;
	v981 = v973
	goto L258
L258:
	;
	if base.B2i32(v971 < int32(-1899))|base.B2i32(int32(0) <= v973) != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	v989 = v981
	goto L261
L260:
	;
	v989 = v973 + int32(100)
	goto L261
L261:
	;
	v991 = v989 >> (uint(int32(31)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+272)) = v991 ^ v989 - v991
	v996 = v17 + int32(292)
	v1000 = F_pg_sprintf(m, v996, int32(_a_F__fmt_0), v17+int32(272))
	mBase = m.M
	v1001 = m.ExcPending
	if v1001 != 0 {
		goto L97
	} else {
		goto L262
	}
L262:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L263
	}
L263:
	;
	v1003 = v996
	v1005 = v21
	goto L264
L264:
	;
	v1017 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1003))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1005))) = uint8(v1017)
	if v1017 == int32(0) {
		v1373 = v1005
		v1378 = v51
		goto L3
	} else {
		goto L266
	}
L265:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L266:
	;
	v1021 = int32(1)
	v1024 = v1005 + v1021
	if v1024 != l3 {
		v1003 = v1003 + v1021
		v1005 = v1024
		goto L264
	} else {
		goto L267
	}
L267:
	;
	goto L265
L268:
	;
	v19 = v43 + int32(2)
	v21 = v1030
	goto L1
L269:
	;
	v1039 = v1034
	v1041 = v21
	goto L270
L270:
	;
	v1053 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1039))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1041))) = uint8(v1053)
	if v1053 == int32(0) {
		v1373 = v1041
		v1378 = v51
		goto L3
	} else {
		goto L272
	}
L271:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L272:
	;
	v1057 = int32(1)
	v1060 = v1041 + v1057
	if v1060 != l3 {
		v1039 = v1039 + v1057
		v1041 = v1060
		goto L270
	} else {
		goto L273
	}
L273:
	;
	goto L271
L274:
	;
	v1069 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
	if v1069 != 0 {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1115 = v21
		goto L289
	} else {
		goto L290
	}
L276:
	;
	if v1081 != 0 {
		goto L283
	} else {
		goto L284
	}
L277:
	;
	v1081 = int32(base.Ui32(v1069) >> (uint(int32(31)) % 32))
	goto L276
L278:
	;
	goto L279
L279:
	;
	v1072 = *(*int32)(unsafe.Add(mBase, uint32(l1)+40))
	if v1072 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1086 = int32(_a_F__fmt_15)
	v1088 = int32(0)
	goto L275
L281:
	;
	goto L282
L282:
	;
	v1077 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1072))))
	v1081 = base.B2i32(v1077 == int32(45))
	goto L276
L283:
	;
	v1082 = int32(_a_F__fmt_16)
	goto L285
L284:
	;
	v1082 = int32(_a_F__fmt_15)
	goto L285
L285:
	;
	if v1081 != 0 {
		goto L286
	} else {
		goto L287
	}
L286:
	;
	v1085 = int32(0) - v1069
	goto L288
L287:
	;
	v1085 = v1069
	goto L288
L288:
	;
	v1086 = v1082
	v1088 = v1085
	goto L275
L289:
	;
	v1128 = base.I32_div_s(v1088, int32(3600))
	v1131 = int32(60)
	v1132 = base.I32_div_s(v1088, v1131)
	v1134 = base.I32_rem_s(v1132, v1131)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+288)) = v1128*int32(100) + v1134
	v1138 = v17 + int32(292)
	v1142 = F_pg_sprintf(m, v1138, int32(_a_F__fmt_17), v17+int32(288))
	mBase = m.M
	v1143 = m.ExcPending
	if v1143 != 0 {
		goto L97
	} else {
		goto L295
	}
L290:
	;
	v1090 = v1086
	v1092 = v21
	goto L291
L291:
	;
	v1104 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1090))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1092))) = uint8(v1104)
	if v1104 == int32(0) {
		v1115 = v1092
		goto L289
	} else {
		goto L293
	}
L292:
	;
	v1115 = l3
	goto L289
L293:
	;
	v1108 = int32(1)
	v1111 = v1092 + v1108
	if v1111 != l3 {
		v1090 = v1090 + v1108
		v1092 = v1111
		goto L291
	} else {
		goto L294
	}
L294:
	;
	goto L292
L295:
	;
	if base.Ui32(l3) <= base.Ui32(v1115) {
		v1373 = v1115
		v1378 = v51
		goto L3
	} else {
		goto L296
	}
L296:
	;
	v1145 = v1138
	v1147 = v1115
	goto L297
L297:
	;
	v1159 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1145))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1147))) = uint8(v1159)
	if v1159 == int32(0) {
		v1373 = v1147
		v1378 = v51
		goto L3
	} else {
		goto L299
	}
L298:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L299:
	;
	v1163 = int32(1)
	v1166 = v1147 + v1163
	if v1166 != l3 {
		v1145 = v1145 + v1163
		v1147 = v1166
		goto L297
	} else {
		goto L300
	}
L300:
	;
	goto L298
L301:
	;
	v19 = v43 + int32(2)
	v21 = v1171
	goto L1
L302:
	;
	v1177 = v19
	goto L7
L303:
	;
	v1192 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1177))))
	*(*uint8)(unsafe.Add(mBase, uint32(v21))) = uint8(v1192)
	v1373 = v21 + int32(1)
	v1378 = v1177
	goto L3
L304:
	;
	v1333 = F__yconv(m, v778, v1215, v21, l3)
	mBase = m.M
	v1334 = m.ExcPending
	if v1334 != 0 {
		goto L97
	} else {
		goto L334
	}
L305:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = int32(3)
	v1256 = int32(100)
	v1257 = base.I32_div_s(v1215, v1256)
	v1259 = base.I32_div_s(v778, v1256)
	v1267 = v1215 - v1257*v1256 + (v778 - v1259*v1256)
	v1270 = base.I32_div_s(base.I32_extend16_s(v1267), v1256)
	v1272 = v1257 + v1259 + base.I32_extend16_s(v1270)
	v1273 = int32(0)
	v1278 = base.I32_extend16_s(v1267 - v1270*v1256)
	if base.B2i32(v1272 <= v1273)|base.B2i32(v1273 <= v1278) == v1273 {
		goto L319
	} else {
		goto L320
	}
L306:
	;
	if v1217 == int32(17) {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v17)+208)) = v1214
	v1222 = v17 + int32(292)
	v1226 = F_pg_sprintf(m, v1222, int32(_a_F__fmt_0), v17+int32(208))
	mBase = m.M
	v1227 = m.ExcPending
	if v1227 != 0 {
		goto L97
	} else {
		goto L312
	}
L309:
	;
	goto L305
L310:
	;
	goto L304
L312:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L313
	}
L313:
	;
	v1229 = v1222
	v1231 = v21
	goto L314
L314:
	;
	v1243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1229))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1231))) = uint8(v1243)
	if v1243 == int32(0) {
		v1373 = v1231
		v1378 = v51
		goto L3
	} else {
		goto L316
	}
L315:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L316:
	;
	v1247 = int32(1)
	v1250 = v1231 + v1247
	if v1250 != l3 {
		v1229 = v1229 + v1247
		v1231 = v1250
		goto L314
	} else {
		goto L317
	}
L317:
	;
	goto L315
L318:
	;
	v1296 = v1294 >> (uint(int32(31)) % 32)
	*(*int32)(unsafe.Add(mBase, uint32(v17)+224)) = v1296 ^ v1294 - v1296
	v1301 = v17 + int32(292)
	v1305 = F_pg_sprintf(m, v1301, int32(_a_F__fmt_0), v17+int32(224))
	mBase = m.M
	v1306 = m.ExcPending
	if v1306 != 0 {
		goto L97
	} else {
		goto L328
	}
L319:
	;
	v1294 = v1278 + int32(100)
	goto L318
L320:
	;
	goto L321
L321:
	;
	if v1272 < int32(0) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1290 = v1278 - int32(100)
	goto L324
L323:
	;
	v1290 = v1278
	goto L324
L324:
	;
	if int32(0) < v1278 {
		goto L325
	} else {
		goto L326
	}
L325:
	;
	v1293 = v1290
	goto L327
L326:
	;
	v1293 = v1278
	goto L327
L327:
	;
	v1294 = v1293
	goto L318
L328:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L329
	}
L329:
	;
	v1308 = v1301
	v1310 = v21
	goto L330
L330:
	;
	v1322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1308))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1310))) = uint8(v1322)
	if v1322 == int32(0) {
		v1373 = v1310
		v1378 = v51
		goto L3
	} else {
		goto L332
	}
L331:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L332:
	;
	v1326 = int32(1)
	v1329 = v1310 + v1326
	if v1329 != l3 {
		v1308 = v1308 + v1326
		v1310 = v1329
		goto L330
	} else {
		goto L333
	}
L333:
	;
	goto L331
L334:
	;
	v19 = v43 + int32(2)
	v21 = v1333
	goto L1
L335:
	;
	if base.Ui32(l3) <= base.Ui32(v21) {
		v1373 = v21
		v1378 = v51
		goto L3
	} else {
		goto L336
	}
L336:
	;
	v1346 = v1341
	v1348 = v21
	goto L337
L337:
	;
	v1360 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1346))))
	*(*uint8)(unsafe.Add(mBase, uint32(v1348))) = uint8(v1360)
	if v1360 == int32(0) {
		v1373 = v1348
		v1378 = v51
		goto L3
	} else {
		goto L339
	}
L338:
	;
	v19 = v43 + int32(2)
	v21 = l3
	goto L1
L339:
	;
	v1364 = int32(1)
	v1367 = v1348 + v1364
	if v1367 != l3 {
		v1346 = v1346 + v1364
		v1348 = v1367
		goto L337
	} else {
		goto L340
	}
L340:
	;
	goto L338
}
func F_fastgetattr_4(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v20 int32
	_ = v20
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	v5 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v5)
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
	v15 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+20)))
	if v15&int32(1) == v5 {
		v20 = int32(4)
		v24 = l2 + l1<<(uint(v20)%32) + v20
		v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)))
		if v25 < int32(0) {
			v73 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v79 = v73
				m.G0 = v10 + int32(16)
				return v79
			}
		} else {
			v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+22)))
			v30 = v14 + v28 + v25
			v31 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24)+6)))
			if v31 != int32(1) {
				v79 = v30
				m.G0 = v10 + int32(16)
				return v79
			} else {
				v34 = int32(*(*int16)(unsafe.Add(mBase, uint32(v24)+4)))
				switch v34&int32(_a_F_fastgetattr_4_0) - int32(1) {
				case 0:
					v39 = int32(*(*int8)(unsafe.Add(mBase, uint32(v30))))
					v79 = v39
					m.G0 = v10 + int32(16)
					return v79
				case 1:
					v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v30))))
					v79 = v40
					m.G0 = v10 + int32(16)
					return v79
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v47 = m.ExcPending
					if v47 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = v34
						F_errmsg_internal(m, int32(_a_F_fastgetattr_4_1), v10)
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_fastgetattr_4_2), int32(70), int32(_a_F_fastgetattr_4_3))
							mBase = m.M
							v56 = m.ExcPending
							if v56 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 3:
					v41 = *(*int32)(unsafe.Add(mBase, uint32(v30)))
					v79 = v41
					m.G0 = v10 + int32(16)
					return v79
				}
			}
		}
	} else {
		v57 = int32(1)
		v58 = l1 - v57
		v62 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14+v58>>(uint(int32(3))%32))+23)))
		if int32(base.Ui32(v62)>>(uint(v58&int32(7))%32))&v57 != 0 {
			v73 = F_nocachegetattr(m, l0, l1, l2)
			mBase = m.M
			v74 = m.ExcPending
			if v74 != 0 {
				return int32(0)
			} else {
				v79 = v73
				m.G0 = v10 + int32(16)
				return v79
			}
		} else {
			v68 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l3))) = uint8(v68)
			v79 = int32(0)
			m.G0 = v10 + int32(16)
			return v79
		}
	}
}
func F_fclose(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v23 int32
	_ = v23
	var v27 int32
	_ = v27
	v6 = F_fflush(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v11 = m.T0[v10].(func(*base.Module, int32) int32)(m, l0)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
			if v13&int32(1) == int32(0) {
				v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
				v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
				if v19 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v19)+56)) = v18
				} else {
				}
				if v18 != 0 {
					*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v19
				} else {
				}
				v23 = *(*int32)(unsafe.Add(mBase, _c_F_fclose[0]))
				if l0 == v23 {
					*(*int32)(unsafe.Add(mBase, _c_F_fclose[0])) = v18
				} else {
				}
				v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
				F_emscripten_builtin_free(m, v27)
				mBase = m.M
				F_emscripten_builtin_free(m, l0)
				mBase = m.M
			} else {
			}
			return v6 | v11
		}
	}
}
func F_feof(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	return int32(base.Ui32(v2)>>(uint(int32(4))%32)) & int32(1)
}
func F_fetch_statentries_for_relation(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v34 int32
	_ = v34
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
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
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
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
	var v114 int32
	_ = v114
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
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
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
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
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v3 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v15 = v10 + int32(-48)
	F_ScanKeyInit(m, v15, int32(2), int32(3), int32(184), l1)
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
	v24 = int32(1)
	v27 = F_systable_beginscan(m, l0, int32(3379), v24, int32(0), v24, v15)
	mBase = m.M
	v28 = m.ExcPending
	if v28 != 0 {
		goto L1
	} else {
		goto L4
	}
L3:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L1
	} else {
		goto L50
	}
L4:
	;
	v29 = F_systable_getnext(m, v27)
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L5
	}
L5:
	;
	if v29 != 0 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v34 = v29
	v39 = v3
	goto L9
L7:
	;
	v187 = v3
	goto L8
L8:
	;
	F_systable_endscan(m, v27)
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L1
	} else {
		goto L49
	}
L9:
	;
	v41 = F_palloc0(m, int32(28))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L11
	}
L10:
	;
	v187 = v175
	goto L8
L11:
	;
	v43 = *(*int32)(unsafe.Add(mBase, uint32(v34)+16))
	v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v43)+22)))
	v45 = v43 + v44
	v46 = *(*int32)(unsafe.Add(mBase, uint32(v45)))
	*(*int32)(unsafe.Add(mBase, uint32(v41))) = v46
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v45)+72))
	v49 = F_get_namespace_name(m, v48)
	mBase = m.M
	v50 = m.ExcPending
	if v50 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v49
	v54 = F_pstrdup(m, v45+int32(8))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L1
	} else {
		goto L13
	}
L13:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+8)) = v54
	v57 = *(*int32)(unsafe.Add(mBase, uint32(v45)+96))
	if int32(0) < v57 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v41)+12))
	v65 = int32(0)
	v69 = v62
	goto L17
L15:
	;
	goto L16
L16:
	;
	v98 = F_SysCacheGetAttr(m, int32(64), v34, int32(7), v10+int32(-49))
	mBase = m.M
	v99 = m.ExcPending
	if v99 != 0 {
		goto L1
	} else {
		goto L21
	}
L17:
	;
	v76 = int32(*(*int16)(unsafe.Add(mBase, uint32(v45+int32(104)+v65<<(uint(int32(1))%32)))))
	v77 = F_bms_add_member(m, v69, v76)
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L1
	} else {
		goto L19
	}
L18:
	;
	goto L16
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+12)) = v77
	v81 = v65 + int32(1)
	v82 = *(*int32)(unsafe.Add(mBase, uint32(v45)+96))
	if v81 < v82 {
		v65 = v81
		v69 = v77
		goto L17
	} else {
		goto L20
	}
L20:
	;
	goto L18
L21:
	;
	v101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v101 != 0 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v102 = int32(-1)
	goto L24
L23:
	;
	v102 = base.I32_extend16_s(v98)
	goto L24
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+20)) = v102
	v106 = F_SysCacheGetAttrNotNull(m, int32(64), v34, int32(8))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L25
	}
L25:
	;
	v108 = F_pg_detoast_datum(m, v106)
	mBase = m.M
	v109 = m.ExcPending
	if v109 != 0 {
		goto L1
	} else {
		goto L26
	}
L26:
	;
	v110 = *(*int32)(unsafe.Add(mBase, uint32(v108)+4))
	if v110 != int32(1) {
		goto L3
	} else {
		goto L27
	}
L27:
	;
	v113 = *(*int32)(unsafe.Add(mBase, uint32(v108)+8))
	if v113 != 0 {
		goto L3
	} else {
		goto L28
	}
L28:
	;
	v114 = *(*int32)(unsafe.Add(mBase, uint32(v108)+12))
	if v114 != int32(18) {
		goto L3
	} else {
		goto L29
	}
L29:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	if int32(0) < v117 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v41)+16))
	v125 = int32(0)
	v129 = v122
	goto L33
L31:
	;
	goto L32
L32:
	;
	v156 = F_SysCacheGetAttr(m, int32(64), v34, int32(9), v10+int32(-49))
	mBase = m.M
	v157 = m.ExcPending
	if v157 != 0 {
		goto L1
	} else {
		goto L37
	}
L33:
	;
	v134 = int32(*(*int8)(unsafe.Add(mBase, uint32(v125+(v108+int32(24))))))
	v135 = F_lappend_int(m, v129, v134)
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L1
	} else {
		goto L35
	}
L34:
	;
	goto L32
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+16)) = v135
	v139 = v125 + int32(1)
	v140 = *(*int32)(unsafe.Add(mBase, uint32(v108)+16))
	if v139 < v140 {
		v125 = v139
		v129 = v135
		goto L33
	} else {
		goto L36
	}
L36:
	;
	goto L34
L37:
	;
	v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+15)))
	if v158 == int32(0) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v161 = F_text_to_cstring(m, v156)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L1
	} else {
		goto L41
	}
L39:
	;
	v172 = int32(0)
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v41)+24)) = v172
	v175 = F_lappend(m, v39, v41)
	mBase = m.M
	v176 = m.ExcPending
	if v176 != 0 {
		goto L1
	} else {
		goto L46
	}
L41:
	;
	v163 = F_stringToNode(m, v161)
	mBase = m.M
	v164 = m.ExcPending
	if v164 != 0 {
		goto L1
	} else {
		goto L42
	}
L42:
	;
	F_pfree(m, v161)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v168 = F_eval_const_expressions(m, int32(0), v163)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L1
	} else {
		goto L44
	}
L44:
	;
	F_fix_opfuncids(m, v168)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L1
	} else {
		goto L45
	}
L45:
	;
	v172 = v168
	goto L40
L46:
	;
	v177 = F_systable_getnext(m, v27)
	mBase = m.M
	v178 = m.ExcPending
	if v178 != 0 {
		goto L1
	} else {
		goto L47
	}
L47:
	;
	if v177 != 0 {
		v34 = v177
		v39 = v175
		goto L9
	} else {
		goto L48
	}
L48:
	;
	goto L10
L49:
	;
	m.G0 = v12 - int32(-64)
	return v187
L50:
	;
	F_errmsg_internal(m, int32(_a_F_fetch_statentries_for_relation_0), int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_fetch_statentries_for_relation_1), int32(470), int32(_a_F_fetch_statentries_for_relation_2))
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fflush(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
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
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int32
	_ = v59
	var v61 int64
	_ = v61
	v2 = int32(0)
	if l0 == v2 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[0]))
	if v7 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v40 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v40 == v41 {
		goto L23
	} else {
		goto L24
	}
L4:
	;
	v9 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[0]))
	v10 = F_fflush(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	v14 = v2
	goto L6
L6:
	;
	v16 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[1]))
	if v16 != 0 {
		goto L9
	} else {
		goto L10
	}
L7:
	;
	return int32(0)
L8:
	;
	v14 = v10
	goto L6
L9:
	;
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[1]))
	v19 = F_fflush(m, v18)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L7
	} else {
		goto L12
	}
L10:
	;
	v22 = v14
	goto L11
L11:
	;
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_fflush[2]))
	if v24 != 0 {
		goto L13
	} else {
		goto L14
	}
L12:
	;
	v22 = v19 | v14
	goto L11
L13:
	;
	v25 = v24
	v26 = v22
	goto L16
L14:
	;
	v37 = v22
	goto L15
L15:
	;
	return v37
L16:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v25)+20))
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v25)+28))
	if v28 != v29 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v37 = v34
	goto L15
L18:
	;
	v31 = F_fflush(m, v25)
	mBase = m.M
	v32 = m.ExcPending
	if v32 != 0 {
		goto L7
	} else {
		goto L21
	}
L19:
	;
	v34 = v26
	goto L20
L20:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v25)+56))
	if v35 != 0 {
		v25 = v35
		v26 = v34
		goto L16
	} else {
		goto L22
	}
L21:
	;
	v34 = v31 | v26
	goto L20
L22:
	;
	goto L17
L23:
	;
	v51 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v52 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v51 != v52 {
		goto L27
	} else {
		goto L28
	}
L24:
	;
	v43 = int32(0)
	v45 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v46 = m.T0[v45].(func(*base.Module, int32, int32, int32) int32)(m, l0, v43, v43)
	mBase = m.M
	v47 = m.ExcPending
	if v47 != 0 {
		goto L7
	} else {
		goto L25
	}
L25:
	;
	v48 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v48 != 0 {
		goto L23
	} else {
		goto L26
	}
L26:
	;
	return int32(-1)
L27:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v58 = m.T0[v57].(func(*base.Module, int32, int64, int32) int64)(m, l0, base.I64_extend_i32_s(v51-v52), int32(1))
	mBase = m.M
	goto L29
L28:
	;
	goto L29
L29:
	;
	v59 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+28)) = v59
	v61 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(l0)+16)) = v61
	*(*int64)(unsafe.Add(mBase, uint32(l0)+4)) = v61
	return v59
}
func F_filter_prepare_cb_wrapper(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	v4 = int32(0)
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = int32(_a_F_filter_prepare_cb_wrapper_0)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = l0
	*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(993)
	v17 = int32(_a_F_filter_prepare_cb_wrapper_1)
	v18 = *(*int32)(unsafe.Add(mBase, _c_F_filter_prepare_cb_wrapper[0]))
	*(*int32)(unsafe.Add(mBase, _c_F_filter_prepare_cb_wrapper[0])) = v8 + int32(4)
	*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v18
	*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v8 + int32(16)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+164)) = uint8(v4)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+147)) = uint8(v4)
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v32 = m.T0[v31].(func(*base.Module, int32, int32, int32) int32)(m, l0, l1, l2)
	mBase = m.M
	v35 = m.ExcPending
	if v35 != 0 {
		return int32(0)
	} else {
		v37 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
		*(*int32)(unsafe.Add(mBase, _c_F_filter_prepare_cb_wrapper[0])) = v37
		m.G0 = v8 + int32(32)
		return v32
	}
}
func F_findDependentObjects(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32, l5 int32, l6 int32) {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int64
	_ = v251
	var v257 int32
	_ = v257
	var v258 int32
	_ = v258
	var v262 int32
	_ = v262
	var v266 int32
	_ = v266
	var v275 int32
	_ = v275
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v292 int32
	_ = v292
	var v294 int32
	_ = v294
	var v296 int32
	_ = v296
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v307 int32
	_ = v307
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v334 int32
	_ = v334
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v388 int64
	_ = v388
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v411 int32
	_ = v411
	var v416 int32
	_ = v416
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v441 int32
	_ = v441
	var v446 int32
	_ = v446
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v465 int32
	_ = v465
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v483 int32
	_ = v483
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v513 int32
	_ = v513
	var v514 int32
	_ = v514
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v520 int32
	_ = v520
	var v521 int32
	_ = v521
	var v527 int32
	_ = v527
	var v528 int32
	_ = v528
	var v529 int32
	_ = v529
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v549 int32
	_ = v549
	var v553 int32
	_ = v553
	var v555 int64
	_ = v555
	var v560 int32
	_ = v560
	var v561 int32
	_ = v561
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v571 int32
	_ = v571
	var v576 int32
	_ = v576
	var v579 int32
	_ = v579
	var v581 int64
	_ = v581
	var v586 int32
	_ = v586
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v609 int32
	_ = v609
	var v611 int32
	_ = v611
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v647 int32
	_ = v647
	var v648 int32
	_ = v648
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v660 int32
	_ = v660
	var v666 int32
	_ = v666
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v669 int32
	_ = v669
	var v673 int32
	_ = v673
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v693 int32
	_ = v693
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v702 int32
	_ = v702
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v719 int32
	_ = v719
	var v720 int32
	_ = v720
	var v721 int32
	_ = v721
	var v727 int32
	_ = v727
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v738 int32
	_ = v738
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v750 int32
	_ = v750
	var v751 int32
	_ = v751
	var v758 int32
	_ = v758
	var v763 int32
	_ = v763
	var v765 int32
	_ = v765
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v780 int64
	_ = v780
	var v789 int32
	_ = v789
	var v792 int32
	_ = v792
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v801 int32
	_ = v801
	var v807 int32
	_ = v807
	var v814 int32
	_ = v814
	var v817 int32
	_ = v817
	var v828 int32
	_ = v828
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v846 int32
	_ = v846
	var v848 int32
	_ = v848
	var v854 int32
	_ = v854
	var v855 int32
	_ = v855
	var v857 int32
	_ = v857
	var v858 int32
	_ = v858
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v875 int32
	_ = v875
	var v881 int32
	_ = v881
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v896 int32
	_ = v896
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v920 int32
	_ = v920
	var v927 int32
	_ = v927
	var v928 int32
	_ = v928
	var v931 int32
	_ = v931
	var v933 int64
	_ = v933
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v938 int64
	_ = v938
	var v945 int32
	_ = v945
	var v948 int32
	_ = v948
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v960 int32
	_ = v960
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v970 int32
	_ = v970
	var v971 int32
	_ = v971
	var v973 int32
	_ = v973
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v978 int32
	_ = v978
	var v979 int32
	_ = v979
	var v981 int64
	_ = v981
	var v983 int32
	_ = v983
	var v984 int32
	_ = v984
	var v987 int32
	_ = v987
	var v989 int64
	_ = v989
	var v991 int32
	_ = v991
	var v993 int32
	_ = v993
	var v1043 int32
	_ = v1043
	var v1047 int32
	_ = v1047
	var v1048 int32
	_ = v1048
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1058 int32
	_ = v1058
	var v1063 int32
	_ = v1063
	v8 = int32(0)
	v21 = m.G0
	v23 = v21 - int32(304)
	m.G0 = v23
	if l3 == v8 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1043 = m.ExcPending
	if v1043 != 0 {
		goto L20
	} else {
		goto L262
	}
L2:
	;
	m.G0 = v23 + int32(304)
	return
L3:
	;
	F_check_stack_depth(m)
	mBase = m.M
	v100 = m.ExcPending
	if v100 != 0 {
		goto L20
	} else {
		goto L21
	}
L4:
	;
	v38 = l3
	v39 = v8
	goto L5
L5:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v38)))
	v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
	if v49 != v51 {
		v70 = v39
		goto L8
	} else {
		goto L9
	}
L6:
	;
	goto L2
L7:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v78 != 0 {
		v38 = v78
		v39 = int32(1)
		goto L5
	} else {
		goto L19
	}
L8:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v38)+8))
	if v74 != 0 {
		v38 = v74
		v39 = v70
		goto L5
	} else {
		goto L17
	}
L9:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
	if v53 != v54 {
		v70 = v39
		goto L8
	} else {
		goto L10
	}
L10:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v50)+8))
	if v57 != v58 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	if v58 == int32(0) {
		goto L7
	} else {
		goto L14
	}
L12:
	;
	v64 = int32(1)
	v65 = l1
	goto L13
L13:
	;
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v38)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v38)+4)) = v66 | v65
	v70 = v64
	goto L8
L14:
	;
	if l1 == int32(0) {
		v70 = v39
		goto L8
	} else {
		goto L15
	}
L15:
	;
	if v57 != 0 {
		v70 = v39
		goto L8
	} else {
		goto L16
	}
L16:
	;
	v64 = v39
	v65 = l1 | int32(256)
	goto L13
L17:
	;
	if v70&int32(1) != 0 {
		goto L2
	} else {
		goto L18
	}
L18:
	;
	goto L3
L19:
	;
	goto L6
L20:
	;
	return
L21:
	;
	v101 = int32(0)
	v102 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v104 = v102 - int32(1)
	if v104 < v101 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v193 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	goto L42
L23:
	;
	v118 = v104
	v119 = v101
	goto L24
L24:
	;
	v129 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v130 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v133 = v130 + v118*int32(12)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v133)))
	if v129 != v134 {
		v157 = v119
		goto L29
	} else {
		goto L30
	}
L26:
	;
	v118 = v118 - int32(1)
	v119 = v167
	goto L24
L27:
	;
	if v157&int32(1) != 0 {
		goto L2
	} else {
		goto L40
	}
L28:
	;
	if v118 != 0 {
		v167 = int32(1)
		goto L26
	} else {
		goto L39
	}
L29:
	;
	if v118 <= int32(0) {
		goto L27
	} else {
		goto L38
	}
L30:
	;
	v136 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v137 = *(*int32)(unsafe.Add(mBase, uint32(v133)+4))
	if v136 != v137 {
		v157 = v119
		goto L29
	} else {
		goto L31
	}
L31:
	;
	v140 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v141 = *(*int32)(unsafe.Add(mBase, uint32(v133)+8))
	if v140 != v141 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	if v141 == int32(0) {
		goto L28
	} else {
		goto L35
	}
L33:
	;
	v147 = int32(1)
	v148 = l1
	goto L34
L34:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v152 = v149 + v118<<(uint(int32(4))%32)
	v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)))
	*(*int32)(unsafe.Add(mBase, uint32(v152))) = v153 | v148
	v157 = v147
	goto L29
L35:
	;
	if l1 == int32(0) {
		v157 = v119
		goto L29
	} else {
		goto L36
	}
L36:
	;
	if v140 != 0 {
		v157 = v119
		goto L29
	} else {
		goto L37
	}
L37:
	;
	v147 = v119
	v148 = l1 | int32(256)
	goto L34
L38:
	;
	v167 = v157
	goto L26
L39:
	;
	goto L2
L40:
	;
	goto L22
L41:
	;
	F_pfree(m, v920)
	mBase = m.M
	v927 = m.ExcPending
	if v927 != 0 {
		goto L20
	} else {
		goto L245
	}
L42:
	;
	if base.B2i32(base.B2i32(v193 == int32(2613))|base.B2i32(base.Ui32(int32(_a_F_findDependentObjects_0)) < base.Ui32(v194)) == int32(0))&((base.B2i32(v193 != int32(2615))|base.B2i32(v194 != int32(2200)))&base.B2i32(v193 != int32(1262))) == int32(0) {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v23+int32(160), int32(1), int32(3), int32(184), v219)
	mBase = m.M
	v221 = m.ExcPending
	if v221 != 0 {
		goto L20
	} else {
		goto L46
	}
L44:
	;
	goto L45
L45:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L20
	} else {
		goto L240
	}
L46:
	;
	v222 = int32(2)
	v224 = v23 + int32(208)
	v228 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v224, v222, int32(3), int32(184), v228)
	mBase = m.M
	v230 = m.ExcPending
	if v230 != 0 {
		goto L20
	} else {
		goto L47
	}
L47:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v231 != 0 {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v234 = int32(3)
	F_ScanKeyInit(m, v23+int32(256), v234, v234, int32(65), v231)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L20
	} else {
		goto L51
	}
L49:
	;
	v240 = v222
	goto L50
L50:
	;
	v241 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v247 = F_systable_beginscan(m, v241, int32(2673), int32(1), int32(0), v240, v23+int32(160))
	mBase = m.M
	v248 = m.ExcPending
	if v248 != 0 {
		goto L20
	} else {
		goto L52
	}
L51:
	;
	v240 = int32(3)
	goto L50
L52:
	;
	v249 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v249
	v251 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+136)) = v251
	*(*int64)(unsafe.Add(mBase, uint32(v23)+120)) = v251
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v249
	v257 = F_systable_getnext(m, v247)
	mBase = m.M
	v258 = m.ExcPending
	if v258 != 0 {
		goto L20
	} else {
		goto L55
	}
L53:
	;
	v854 = *(*int32)(unsafe.Add(mBase, uint32(v23)+120))
	if v854 != 0 {
		goto L230
	} else {
		goto L231
	}
L54:
	;
	v631 = F_palloc(m, int32(2048))
	mBase = m.M
	v632 = m.ExcPending
	if v632 != 0 {
		goto L20
	} else {
		goto L163
	}
L55:
	;
	if v257 == int32(0) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	F_systable_endscan(m, v247)
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L20
	} else {
		goto L59
	}
L57:
	;
	goto L58
L58:
	;
	v266 = l1
	v275 = v257
	goto L60
L59:
	;
	v611 = l1
	goto L54
L60:
	;
	v285 = *(*int32)(unsafe.Add(mBase, uint32(v275)+16))
	v286 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v285)+22)))
	v287 = v285 + v286
	v288 = *(*int32)(unsafe.Add(mBase, uint32(v287)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v288
	v290 = *(*int32)(unsafe.Add(mBase, uint32(v287)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = v290
	v292 = *(*int32)(unsafe.Add(mBase, uint32(v287)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+156)) = v292
	v294 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v288 != v294 {
		goto L63
	} else {
		goto L64
	}
L61:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v23)+136))
	F_systable_endscan(m, v247)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L20
	} else {
		goto L161
	}
L62:
	;
	v605 = F_systable_getnext(m, v247)
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L20
	} else {
		goto L159
	}
L63:
	;
	v301 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v287)+24)))
	switch v301 - int32(80) {
	case 0:
		goto L70
	default:
		goto L69
	case 3:
		goto L68
	case 17, 30, 40:
		v586 = v266
		goto L62
	case 21:
		goto L72
	case 25:
		goto L71
	}
L64:
	;
	v296 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v290 != v296 {
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v298 == int32(0) {
		v586 = v266
		goto L62
	} else {
		goto L66
	}
L66:
	;
	goto L63
L67:
	;
	v586 = v266 | int32(128)
	goto L62
L68:
	;
	if v266&int32(128) != 0 {
		goto L67
	} else {
		goto L158
	}
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v560 = m.ExcPending
	if v560 != 0 {
		goto L20
	} else {
		goto L154
	}
L70:
	;
	v553 = *(*int32)(unsafe.Add(mBase, uint32(v23)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v553
	v555 = *(*int64)(unsafe.Add(mBase, uint32(v23)+148))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+120)) = v555
	goto L67
L71:
	;
	v315 = int32(0)
	if l3 == v315 {
		goto L77
	} else {
		goto L78
	}
L72:
	;
	if l2&int32(16) != 0 {
		v586 = v266
		goto L62
	} else {
		goto L73
	}
L73:
	;
	if v288 != int32(3079) {
		goto L71
	} else {
		goto L74
	}
L74:
	;
	v307 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_findDependentObjects[0])))
	if v307&int32(1) == int32(0) {
		goto L71
	} else {
		goto L75
	}
L75:
	;
	v313 = *(*int32)(unsafe.Add(mBase, _c_F_findDependentObjects[1]))
	if v290 == v313 {
		v586 = v266
		goto L62
	} else {
		goto L76
	}
L76:
	;
	goto L71
L77:
	;
	if l5 == int32(0) {
		goto L81
	} else {
		goto L82
	}
L78:
	;
	goto L79
L79:
	;
	v411 = l3
	v416 = v315
	goto L101
L80:
	;
	F_systable_endscan(m, v247)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L20
	} else {
		goto L95
	}
L81:
	;
	if v301 != int32(101) {
		goto L91
	} else {
		goto L92
	}
L82:
	;
	v320 = *(*int32)(unsafe.Add(mBase, uint32(l5)+8))
	v322 = v320 - int32(1)
	if v322 < int32(0) {
		goto L81
	} else {
		goto L83
	}
L83:
	;
	v325 = *(*int32)(unsafe.Add(mBase, uint32(l5)))
	v334 = v322
	goto L84
L84:
	;
	v348 = v325 + v334*int32(12)
	v349 = *(*int32)(unsafe.Add(mBase, uint32(v348)))
	if v288 != v349 {
		goto L86
	} else {
		goto L87
	}
L85:
	;
	goto L81
L86:
	;
	if int32(0) < v334 {
		v334 = v334 - int32(1)
		goto L84
	} else {
		goto L90
	}
L87:
	;
	v351 = *(*int32)(unsafe.Add(mBase, uint32(v348)+4))
	if v290 != v351 {
		goto L86
	} else {
		goto L88
	}
L88:
	;
	v353 = *(*int32)(unsafe.Add(mBase, uint32(v348)+8))
	if base.B2i32(v292 == v353)|base.B2i32(v353 == int32(0)) != 0 {
		goto L80
	} else {
		goto L89
	}
L89:
	;
	goto L86
L90:
	;
	goto L85
L91:
	;
	v385 = *(*int32)(unsafe.Add(mBase, uint32(v23)+136))
	if v385 != 0 {
		v586 = v266
		goto L62
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v23)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+144)) = v386
	v388 = *(*int64)(unsafe.Add(mBase, uint32(v23)+148))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+136)) = v388
	v586 = v266
	goto L62
L94:
	;
	goto L93
L95:
	;
	v392 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v393 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v393 == int32(1259) {
		goto L96
	} else {
		goto L97
	}
L96:
	;
	F_UnlockRelationOid(m, v392, int32(8))
	mBase = m.M
	v398 = m.ExcPending
	if v398 != 0 {
		goto L20
	} else {
		goto L99
	}
L97:
	;
	goto L98
L98:
	;
	F_UnlockDatabaseObject(m, v393, v392, int32(8))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L20
	} else {
		goto L100
	}
L99:
	;
	goto L2
L100:
	;
	goto L2
L101:
	;
	v422 = *(*int32)(unsafe.Add(mBase, uint32(v411)))
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v422)))
	if v288 != v423 {
		v437 = v416
		goto L103
	} else {
		goto L104
	}
L102:
	;
	if v437&int32(1) != 0 {
		v586 = v266
		goto L62
	} else {
		goto L111
	}
L103:
	;
	v438 = *(*int32)(unsafe.Add(mBase, uint32(v411)+8))
	if v438 != 0 {
		v411 = v438
		v416 = v437
		goto L101
	} else {
		goto L110
	}
L104:
	;
	v425 = *(*int32)(unsafe.Add(mBase, uint32(v422)+4))
	if v290 != v425 {
		v437 = v416
		goto L103
	} else {
		goto L105
	}
L105:
	;
	v427 = *(*int32)(unsafe.Add(mBase, uint32(v422)+8))
	if v427 == v292 {
		goto L106
	} else {
		goto L107
	}
L106:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v411)+8))
	if v430 == int32(0) {
		v586 = v266
		goto L62
	} else {
		goto L109
	}
L107:
	;
	goto L108
L108:
	;
	v437 = base.B2i32(v427 == int32(0)) | v416
	goto L103
L109:
	;
	v411 = v430
	v416 = int32(1)
	goto L101
L110:
	;
	goto L102
L111:
	;
	v441 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v294 == int32(1259) {
		goto L113
	} else {
		goto L114
	}
L112:
	;
	switch v288 - int32(1259) {
	case 0:
		goto L121
	default:
		goto L119
	case 2:
		goto L120
	}
L113:
	;
	F_UnlockRelationOid(m, v441, int32(8))
	mBase = m.M
	v446 = m.ExcPending
	if v446 != 0 {
		goto L20
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	F_UnlockDatabaseObject(m, v294, v441, int32(8))
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L20
	} else {
		goto L117
	}
L116:
	;
	goto L112
L117:
	;
	goto L112
L118:
	;
	v462 = F_systable_recheck_tuple(m, v247)
	mBase = m.M
	v463 = m.ExcPending
	if v463 != 0 {
		goto L20
	} else {
		goto L125
	}
L119:
	;
	F_LockDatabaseObject(m, v288, v290, int32(8))
	mBase = m.M
	v461 = m.ExcPending
	if v461 != 0 {
		goto L20
	} else {
		goto L124
	}
L120:
	;
	F_LockSharedObject(m, int32(1261), v290, int32(8))
	mBase = m.M
	v458 = m.ExcPending
	if v458 != 0 {
		goto L20
	} else {
		goto L123
	}
L121:
	;
	F_LockRelationOid(m, v290, int32(8))
	mBase = m.M
	v454 = m.ExcPending
	if v454 != 0 {
		goto L20
	} else {
		goto L122
	}
L122:
	;
	goto L118
L123:
	;
	goto L118
L124:
	;
	goto L118
L125:
	;
	F_systable_endscan(m, v247)
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L20
	} else {
		goto L126
	}
L126:
	;
	if v462 == int32(0) {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	if v288 == int32(1259) {
		goto L130
	} else {
		goto L131
	}
L128:
	;
	goto L129
L129:
	;
	F_findDependentObjects(m, v23+int32(148), int32(64), l2, l3, l4, l5, l6)
	mBase = m.M
	v480 = m.ExcPending
	if v480 != 0 {
		goto L20
	} else {
		goto L135
	}
L130:
	;
	F_UnlockRelationOid(m, v290, int32(8))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L20
	} else {
		goto L133
	}
L131:
	;
	goto L132
L132:
	;
	F_UnlockDatabaseObject(m, v288, v290, int32(8))
	mBase = m.M
	v475 = m.ExcPending
	if v475 != 0 {
		goto L20
	} else {
		goto L134
	}
L133:
	;
	goto L2
L134:
	;
	goto L2
L135:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v483 = v481 - int32(1)
	if v483 < int32(0) {
		goto L1
	} else {
		goto L136
	}
L136:
	;
	v498 = v483
	v499 = int32(0)
	goto L137
L137:
	;
	v509 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v510 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v513 = v510 + v498*int32(12)
	v514 = *(*int32)(unsafe.Add(mBase, uint32(v513)))
	if v509 != v514 {
		v539 = v499
		goto L142
	} else {
		goto L143
	}
L139:
	;
	v498 = v498 - int32(1)
	v499 = v549
	goto L137
L140:
	;
	if v539&int32(1) != 0 {
		goto L2
	} else {
		goto L153
	}
L141:
	;
	if v498 != 0 {
		v549 = int32(1)
		goto L139
	} else {
		goto L152
	}
L142:
	;
	if v498 <= int32(0) {
		goto L140
	} else {
		goto L151
	}
L143:
	;
	v516 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v517 = *(*int32)(unsafe.Add(mBase, uint32(v513)+4))
	if v516 != v517 {
		v539 = v499
		goto L142
	} else {
		goto L144
	}
L144:
	;
	v520 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v521 = *(*int32)(unsafe.Add(mBase, uint32(v513)+8))
	if v520 != v521 {
		goto L145
	} else {
		goto L146
	}
L145:
	;
	if v521 == int32(0) {
		goto L141
	} else {
		goto L148
	}
L146:
	;
	v527 = int32(1)
	v528 = v266
	goto L147
L147:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v532 = v529 + v498<<(uint(int32(4))%32)
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v532)))
	*(*int32)(unsafe.Add(mBase, uint32(v532))) = v533 | v528
	v539 = v527
	goto L142
L148:
	;
	if v266 == int32(0) {
		v539 = v499
		goto L142
	} else {
		goto L149
	}
L149:
	;
	if v520 != 0 {
		v539 = v499
		goto L142
	} else {
		goto L150
	}
L150:
	;
	v527 = v499
	v528 = v266 | int32(256)
	goto L147
L151:
	;
	v549 = v539
	goto L139
L152:
	;
	goto L2
L153:
	;
	goto L1
L154:
	;
	v561 = int32(*(*int8)(unsafe.Add(mBase, uint32(v287)+24)))
	v563 = F_getObjectDescription(m, l0, int32(0))
	mBase = m.M
	v564 = m.ExcPending
	if v564 != 0 {
		goto L20
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+68)) = v563
	*(*int32)(unsafe.Add(mBase, uint32(v23)+64)) = v561
	F_errmsg_internal(m, int32(_a_F_findDependentObjects_1), v23-int32(-64))
	mBase = m.M
	v571 = m.ExcPending
	if v571 != 0 {
		goto L20
	} else {
		goto L156
	}
L156:
	;
	F_errfinish(m, int32(_a_F_findDependentObjects_2), int32(765), int32(_a_F_findDependentObjects_3))
	mBase = m.M
	v576 = m.ExcPending
	if v576 != 0 {
		goto L20
	} else {
		goto L157
	}
L157:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L158:
	;
	v579 = *(*int32)(unsafe.Add(mBase, uint32(v23)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+128)) = v579
	v581 = *(*int64)(unsafe.Add(mBase, uint32(v23)+148))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+120)) = v581
	goto L67
L159:
	;
	if v605 != 0 {
		v266 = v586
		v275 = v605
		goto L60
	} else {
		goto L160
	}
L160:
	;
	goto L61
L161:
	;
	if v607 != 0 {
		goto L53
	} else {
		goto L162
	}
L162:
	;
	v611 = v586
	goto L54
L163:
	;
	v633 = int32(3)
	v639 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_ScanKeyInit(m, v23+int32(160), int32(4), v633, int32(184), v639)
	mBase = m.M
	v641 = m.ExcPending
	if v641 != 0 {
		goto L20
	} else {
		goto L164
	}
L164:
	;
	v645 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	F_ScanKeyInit(m, v224, int32(5), int32(3), int32(184), v645)
	mBase = m.M
	v647 = m.ExcPending
	if v647 != 0 {
		goto L20
	} else {
		goto L165
	}
L165:
	;
	v648 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v648 == int32(0) {
		goto L167
	} else {
		goto L168
	}
L166:
	;
	v660 = *(*int32)(unsafe.Add(mBase, uint32(l6)))
	v666 = F_systable_beginscan(m, v660, int32(2674), int32(1), int32(0), v659, v23+int32(160))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L20
	} else {
		goto L171
	}
L167:
	;
	v659 = int32(2)
	goto L166
L168:
	;
	goto L169
L169:
	;
	F_ScanKeyInit(m, v23+int32(256), int32(6), int32(3), int32(65), v648)
	mBase = m.M
	v658 = m.ExcPending
	if v658 != 0 {
		goto L20
	} else {
		goto L170
	}
L170:
	;
	v659 = v633
	goto L166
L171:
	;
	v668 = F_systable_getnext(m, v666)
	mBase = m.M
	v669 = m.ExcPending
	if v669 != 0 {
		goto L20
	} else {
		goto L172
	}
L172:
	;
	if v668 == int32(0) {
		goto L173
	} else {
		goto L174
	}
L173:
	;
	F_systable_endscan(m, v666)
	mBase = m.M
	v673 = m.ExcPending
	if v673 != 0 {
		goto L20
	} else {
		goto L176
	}
L174:
	;
	goto L175
L175:
	;
	v687 = v668
	v689 = int32(128)
	v690 = int32(0)
	v693 = v631
	goto L177
L176:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = l0
	v920 = v631
	goto L41
L177:
	;
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v687)+16))
	v700 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v699)+22)))
	v701 = v699 + v700
	v702 = *(*int32)(unsafe.Add(mBase, uint32(v701)))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+148)) = v702
	v704 = *(*int32)(unsafe.Add(mBase, uint32(v701)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+152)) = v704
	v706 = *(*int32)(unsafe.Add(mBase, uint32(v701)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+156)) = v706
	v708 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	if v702 != v708 {
		goto L180
	} else {
		goto L181
	}
L178:
	;
	F_systable_endscan(m, v666)
	mBase = m.M
	v801 = m.ExcPending
	if v801 != 0 {
		goto L20
	} else {
		goto L219
	}
L179:
	;
	v798 = F_systable_getnext(m, v666)
	mBase = m.M
	v799 = m.ExcPending
	if v799 != 0 {
		goto L20
	} else {
		goto L217
	}
L180:
	;
	switch v702 - int32(1259) {
	case 0:
		goto L190
	default:
		goto L188
	case 2:
		goto L189
	}
L181:
	;
	v710 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v704 != v710 {
		goto L180
	} else {
		goto L182
	}
L182:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v712 == int32(0) {
		v795 = v689
		v796 = v690
		v797 = v693
		goto L179
	} else {
		goto L183
	}
L183:
	;
	goto L180
L184:
	;
	F_UnlockDatabaseObject(m, v702, v704, int32(8))
	mBase = m.M
	v792 = m.ExcPending
	if v792 != 0 {
		goto L20
	} else {
		goto L216
	}
L185:
	;
	F_UnlockRelationOid(m, v704, int32(8))
	mBase = m.M
	v789 = m.ExcPending
	if v789 != 0 {
		goto L20
	} else {
		goto L215
	}
L186:
	;
	if v702 != int32(1259) {
		goto L184
	} else {
		goto L214
	}
L187:
	;
	v738 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v701)+24)))
	switch v738 - int32(80) {
	case 0, 3:
		goto L204
	default:
		goto L202
	case 17, 40:
		goto L201
	case 21:
		goto L203
	case 25:
		goto L205
	case 30:
		v765 = int32(2)
		goto L200
	}
L188:
	;
	F_LockDatabaseObject(m, v702, v704, int32(8))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L20
	} else {
		goto L197
	}
L189:
	;
	F_LockSharedObject(m, int32(1261), v704, int32(8))
	mBase = m.M
	v727 = m.ExcPending
	if v727 != 0 {
		goto L20
	} else {
		goto L194
	}
L190:
	;
	F_LockRelationOid(m, v704, int32(8))
	mBase = m.M
	v719 = m.ExcPending
	if v719 != 0 {
		goto L20
	} else {
		goto L191
	}
L191:
	;
	v720 = F_systable_recheck_tuple(m, v666)
	mBase = m.M
	v721 = m.ExcPending
	if v721 != 0 {
		goto L20
	} else {
		goto L192
	}
L192:
	;
	if v720 == int32(0) {
		goto L185
	} else {
		goto L193
	}
L193:
	;
	goto L187
L194:
	;
	v728 = F_systable_recheck_tuple(m, v666)
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L20
	} else {
		goto L195
	}
L195:
	;
	if v728 != 0 {
		goto L187
	} else {
		goto L196
	}
L196:
	;
	goto L184
L197:
	;
	v733 = F_systable_recheck_tuple(m, v666)
	mBase = m.M
	v734 = m.ExcPending
	if v734 != 0 {
		goto L20
	} else {
		goto L198
	}
L198:
	;
	if v733 == int32(0) {
		goto L186
	} else {
		goto L199
	}
L199:
	;
	goto L187
L200:
	;
	if v689 <= v690 {
		goto L210
	} else {
		goto L211
	}
L201:
	;
	v765 = int32(4)
	goto L200
L202:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L20
	} else {
		goto L206
	}
L203:
	;
	v765 = int32(32)
	goto L200
L204:
	;
	v765 = int32(16)
	goto L200
L205:
	;
	v765 = int32(8)
	goto L200
L206:
	;
	v748 = int32(*(*int8)(unsafe.Add(mBase, uint32(v701)+24)))
	v750 = F_getObjectDescription(m, l0, int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L20
	} else {
		goto L207
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+20)) = v750
	*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v748
	F_errmsg_internal(m, int32(_a_F_findDependentObjects_1), v23+int32(16))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L20
	} else {
		goto L208
	}
L208:
	;
	F_errfinish(m, int32(_a_F_findDependentObjects_2), int32(893), int32(_a_F_findDependentObjects_3))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L20
	} else {
		goto L209
	}
L209:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L210:
	;
	v769 = F_repalloc(m, v693, v689<<(uint(int32(5))%32))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L20
	} else {
		goto L213
	}
L211:
	;
	v773 = v689
	v774 = v693
	goto L212
L212:
	;
	v777 = v774 + v690<<(uint(int32(4))%32)
	v778 = *(*int32)(unsafe.Add(mBase, uint32(v23)+156))
	*(*int32)(unsafe.Add(mBase, uint32(v777)+8)) = v778
	v780 = *(*int64)(unsafe.Add(mBase, uint32(v23)+148))
	*(*int64)(unsafe.Add(mBase, uint32(v777))) = v780
	*(*int32)(unsafe.Add(mBase, uint32(v777)+12)) = v765
	v795 = v773
	v796 = v690 + int32(1)
	v797 = v774
	goto L179
L213:
	;
	v773 = v689 << (uint(int32(1)) % 32)
	v774 = v769
	goto L212
L214:
	;
	goto L185
L215:
	;
	v795 = v689
	v796 = v690
	v797 = v693
	goto L179
L216:
	;
	v795 = v689
	v796 = v690
	v797 = v693
	goto L179
L217:
	;
	if v798 != 0 {
		v687 = v798
		v689 = v795
		v690 = v796
		v693 = v797
		goto L177
	} else {
		goto L218
	}
L218:
	;
	goto L178
L219:
	;
	if int32(2) <= v796 {
		goto L221
	} else {
		goto L222
	}
L220:
	;
	v828 = int32(0)
	goto L226
L221:
	;
	F_pg_qsort(m, v797, v796, int32(16), int32(462))
	mBase = m.M
	v807 = m.ExcPending
	if v807 != 0 {
		goto L20
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = l0
	v814 = int32(1)
	if v796 != v814 {
		v920 = v797
		goto L41
	} else {
		goto L225
	}
L224:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+116)) = l3
	*(*int32)(unsafe.Add(mBase, uint32(v23)+112)) = v611
	*(*int32)(unsafe.Add(mBase, uint32(v23)+108)) = l0
	v817 = v796
	goto L220
L225:
	;
	v817 = v814
	goto L220
L226:
	;
	v841 = v797 + v828<<(uint(int32(4))%32)
	v842 = *(*int32)(unsafe.Add(mBase, uint32(v841)+12))
	F_findDependentObjects(m, v841, v842, l2, v23+int32(108), l4, l5, l6)
	mBase = m.M
	v846 = m.ExcPending
	if v846 != 0 {
		goto L20
	} else {
		goto L228
	}
L227:
	;
	v920 = v797
	goto L41
L228:
	;
	v848 = v828 + int32(1)
	if v848 != v817 {
		v828 = v848
		goto L226
	} else {
		goto L229
	}
L229:
	;
	goto L227
L230:
	;
	v855 = v23 + int32(120)
	goto L232
L231:
	;
	v855 = v23 + int32(136)
	goto L232
L232:
	;
	v857 = F_getObjectDescription(m, v855, int32(0))
	mBase = m.M
	v858 = m.ExcPending
	if v858 != 0 {
		goto L20
	} else {
		goto L233
	}
L233:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v862 = m.ExcPending
	if v862 != 0 {
		goto L20
	} else {
		goto L234
	}
L234:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v865 = m.ExcPending
	if v865 != 0 {
		goto L20
	} else {
		goto L235
	}
L235:
	;
	v867 = F_getObjectDescription(m, l0, int32(0))
	mBase = m.M
	v868 = m.ExcPending
	if v868 != 0 {
		goto L20
	} else {
		goto L236
	}
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+52)) = v857
	*(*int32)(unsafe.Add(mBase, uint32(v23)+48)) = v867
	F_errmsg(m, int32(_a_F_findDependentObjects_4), v23+int32(48))
	mBase = m.M
	v875 = m.ExcPending
	if v875 != 0 {
		goto L20
	} else {
		goto L237
	}
L237:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+32)) = v857
	F_errhint(m, int32(_a_F_findDependentObjects_5), v23+int32(32))
	mBase = m.M
	v881 = m.ExcPending
	if v881 != 0 {
		goto L20
	} else {
		goto L238
	}
L238:
	;
	F_errfinish(m, int32(_a_F_findDependentObjects_2), int32(791), int32(_a_F_findDependentObjects_3))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L20
	} else {
		goto L239
	}
L239:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L240:
	;
	F_errcode(m, int32(16909442))
	mBase = m.M
	v893 = m.ExcPending
	if v893 != 0 {
		goto L20
	} else {
		goto L241
	}
L241:
	;
	v895 = F_getObjectDescription(m, l0, int32(0))
	mBase = m.M
	v896 = m.ExcPending
	if v896 != 0 {
		goto L20
	} else {
		goto L242
	}
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23))) = v895
	F_errmsg(m, int32(_a_F_findDependentObjects_6), v23)
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L20
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(_a_F_findDependentObjects_2), int32(498), int32(_a_F_findDependentObjects_3))
	mBase = m.M
	v905 = m.ExcPending
	if v905 != 0 {
		goto L20
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L245:
	;
	v928 = *(*int32)(unsafe.Add(mBase, uint32(v23)+112))
	if v928&int32(128) != 0 {
		goto L247
	} else {
		goto L248
	}
L246:
	;
	v945 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	if v945 == int32(0) {
		goto L253
	} else {
		goto L254
	}
L247:
	;
	v931 = *(*int32)(unsafe.Add(mBase, uint32(v23)+128))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v931
	v933 = *(*int64)(unsafe.Add(mBase, uint32(v23)+120))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = v933
	goto L246
L248:
	;
	goto L249
L249:
	;
	if l3 != 0 {
		goto L250
	} else {
		goto L251
	}
L250:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v936 = *(*int32)(unsafe.Add(mBase, uint32(v935)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = v936
	v938 = *(*int64)(unsafe.Add(mBase, uint32(v935)))
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = v938
	goto L246
L251:
	;
	goto L252
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+104)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v23)+96)) = int64(0)
	goto L246
L253:
	;
	v948 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v951 = F_palloc(m, v948<<(uint(int32(4))%32))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L20
	} else {
		goto L256
	}
L254:
	;
	goto L255
L255:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v955 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	if v955 <= v954 {
		goto L257
	} else {
		goto L258
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v951
	goto L255
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+12)) = v955 << (uint(int32(1)) % 32)
	v960 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v963 = F_repalloc(m, v960, v955*int32(24))
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L20
	} else {
		goto L260
	}
L258:
	;
	v974 = v954
	goto L259
L259:
	;
	v975 = *(*int32)(unsafe.Add(mBase, uint32(l4)))
	v978 = v975 + v974*int32(12)
	v979 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v978)+8)) = v979
	v981 = *(*int64)(unsafe.Add(mBase, uint32(l0)))
	*(*int64)(unsafe.Add(mBase, uint32(v978))) = v981
	v983 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v987 = v983 + v984<<(uint(int32(4))%32)
	*(*int32)(unsafe.Add(mBase, uint32(v987))) = v928
	v989 = *(*int64)(unsafe.Add(mBase, uint32(v23)+96))
	*(*int64)(unsafe.Add(mBase, uint32(v987)+4)) = v989
	v991 = *(*int32)(unsafe.Add(mBase, uint32(v23)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v987)+12)) = v991
	v993 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l4)+8)) = v993 + int32(1)
	goto L2
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4))) = v963
	v966 = *(*int32)(unsafe.Add(mBase, uint32(l4)+4))
	v967 = *(*int32)(unsafe.Add(mBase, uint32(l4)+12))
	v970 = F_repalloc(m, v966, v967<<(uint(int32(4))%32))
	mBase = m.M
	v971 = m.ExcPending
	if v971 != 0 {
		goto L20
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l4)+4)) = v970
	v973 = *(*int32)(unsafe.Add(mBase, uint32(l4)+8))
	v974 = v973
	goto L259
L262:
	;
	v1047 = F_getObjectDescription(m, v23+int32(148), int32(0))
	mBase = m.M
	v1048 = m.ExcPending
	if v1048 != 0 {
		goto L20
	} else {
		goto L263
	}
L263:
	;
	v1050 = F_getObjectDescription(m, l0, int32(0))
	mBase = m.M
	v1051 = m.ExcPending
	if v1051 != 0 {
		goto L20
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v23)+84)) = v1050
	*(*int32)(unsafe.Add(mBase, uint32(v23)+80)) = v1047
	F_errmsg_internal(m, int32(_a_F_findDependentObjects_7), v23+int32(80))
	mBase = m.M
	v1058 = m.ExcPending
	if v1058 != 0 {
		goto L20
	} else {
		goto L265
	}
L265:
	;
	F_errfinish(m, int32(_a_F_findDependentObjects_2), int32(724), int32(_a_F_findDependentObjects_3))
	mBase = m.M
	v1063 = m.ExcPending
	if v1063 != 0 {
		goto L20
	} else {
		goto L266
	}
L266:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_findNewestTimeLine(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v3 = l0
	goto L1
L1:
	;
	v6 = v3 + int32(1)
	v7 = F_existsTimeLineHistory(m, v6)
	v10 = m.ExcPending
	if v10 != 0 {
		goto L3
	} else {
		goto L4
	}
L2:
	;
	return v3
L3:
	;
	return int32(0)
L4:
	;
	if v7 != 0 {
		v3 = v6
		goto L1
	} else {
		goto L5
	}
L5:
	;
	goto L2
}
func F_findVariant(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) int32 {
	mBase := m.M
	_ = mBase
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v35 int32
	_ = v35
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v74 int32
	_ = v74
	var v83 int32
	_ = v83
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v99 int32
	_ = v99
	var v120 int32
	_ = v120
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v144 int32
	_ = v144
	var v152 int32
	_ = v152
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v191 int32
	_ = v191
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v218 int32
	_ = v218
	var v234 int32
	_ = v234
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v281 int32
	_ = v281
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v305 int32
	_ = v305
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v309 int32
	_ = v309
	var v312 int32
	_ = v312
	v18 = l4 & int32(3)
	v21 = l0
	goto L1
L1:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if l4 != 0 {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	if l4 == int32(0) {
		v21 = v218
		goto L1
	} else {
		goto L46
	}
L4:
	;
	v43 = v35
	v44 = int32(0)
	goto L7
L5:
	;
	v144 = v35
	goto L6
L6:
	;
	v152 = *(*int32)(unsafe.Add(mBase, uint32(v144)))
	if l1 == int32(0) {
		goto L33
	} else {
		goto L34
	}
L7:
	;
	v53 = l3 + v44<<(uint(int32(2))%32)
	v54 = *(*int32)(unsafe.Add(mBase, uint32(v53)))
	if v54 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L8:
	;
	if l4 != v135 {
		v218 = v21
		goto L3
	} else {
		goto L32
	}
L9:
	;
	if v135 < l4 {
		v43 = v126
		v44 = v135
		goto L7
	} else {
		goto L31
	}
L10:
	;
	if v92 == v93 {
		goto L28
	} else {
		goto L29
	}
L11:
	;
	return v21
L12:
	;
	v62 = v54
	goto L13
L13:
	;
	v71 = *(*int32)(unsafe.Add(mBase, uint32(v62)))
	v72 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if base.Ui32(v71) < base.Ui32(v72) {
		goto L15
	} else {
		goto L16
	}
L14:
	;
	if base.Ui32(v72) < base.Ui32(v71) {
		v126 = v62
		v135 = int32(0)
		goto L9
	} else {
		goto L19
	}
L15:
	;
	v74 = *(*int32)(unsafe.Add(mBase, uint32(v62)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v74
	if v74 != 0 {
		v62 = v74
		goto L13
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	goto L14
L18:
	;
	goto L11
L19:
	;
	v83 = v62
	goto L20
L20:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v83)))
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v43)))
	if v92 != v93 {
		goto L10
	} else {
		goto L22
	}
L21:
	;
	goto L11
L22:
	;
	v95 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+4)))
	if l2 == v95 {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v97 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+6)))
	if l4 == v97 {
		goto L10
	} else {
		goto L26
	}
L24:
	;
	goto L25
L25:
	;
	v99 = *(*int32)(unsafe.Add(mBase, uint32(v83)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v53))) = v99
	if v99 != 0 {
		v83 = v99
		goto L20
	} else {
		goto L27
	}
L26:
	;
	goto L25
L27:
	;
	goto L21
L28:
	;
	v120 = v44 + int32(1)
	goto L30
L29:
	;
	v120 = int32(0)
	goto L30
L30:
	;
	v126 = v83
	v135 = v120
	goto L9
L31:
	;
	goto L8
L32:
	;
	v144 = v126
	goto L6
L33:
	;
	if v21 != 0 {
		goto L39
	} else {
		goto L40
	}
L34:
	;
	v160 = l1
	goto L35
L35:
	;
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v160)))
	if v169 == v152 {
		goto L33
	} else {
		goto L37
	}
L36:
	;
	v218 = v21
	goto L3
L37:
	;
	v171 = *(*int32)(unsafe.Add(mBase, uint32(v160)+12))
	if v171 != 0 {
		v160 = v171
		goto L35
	} else {
		goto L38
	}
L38:
	;
	goto L36
L39:
	;
	v191 = v21
	goto L42
L40:
	;
	goto L41
L41:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v144)+12)) = v21
	v218 = v144
	goto L3
L42:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
	if v200 == v152 {
		v218 = v21
		goto L3
	} else {
		goto L44
	}
L43:
	;
	goto L41
L44:
	;
	v202 = *(*int32)(unsafe.Add(mBase, uint32(v191)+12))
	if v202 != 0 {
		v191 = v202
		goto L42
	} else {
		goto L45
	}
L45:
	;
	goto L43
L46:
	;
	v234 = int32(0)
	if base.B2i32(base.Ui32(l4) < base.Ui32(int32(4))) == v234 {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v245 = v234
	v247 = v234
	goto L50
L48:
	;
	v281 = v234
	goto L49
L49:
	;
	v295 = v281
	v296 = v234
	goto L54
L50:
	;
	v255 = l3 + v245<<(uint(int32(2))%32)
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)))
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v256)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v255))) = v257
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v255)+4))
	v260 = *(*int32)(unsafe.Add(mBase, uint32(v259)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+4)) = v260
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v255)+8))
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v262)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+8)) = v263
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v265)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v255)+12)) = v266
	v268 = int32(4)
	v269 = v245 + v268
	v271 = v247 + v268
	if v271 != l4&int32(_a_F_findVariant_0) {
		v245 = v269
		v247 = v271
		goto L50
	} else {
		goto L52
	}
L51:
	;
	if v18 == int32(0) {
		v21 = v218
		goto L1
	} else {
		goto L53
	}
L52:
	;
	goto L51
L53:
	;
	v281 = v269
	goto L49
L54:
	;
	v305 = l3 + v295<<(uint(int32(2))%32)
	v306 = *(*int32)(unsafe.Add(mBase, uint32(v305)))
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v306)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v305))) = v307
	v309 = int32(1)
	v312 = v296 + v309
	if v312 != v18 {
		v295 = v295 + v309
		v296 = v312
		goto L54
	} else {
		goto L56
	}
L55:
	;
	v21 = v218
	goto L1
L56:
	;
	goto L55
}
func F_find_mergeclauses_for_outer_pathkeys(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v11 int32
	_ = v11
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v38 int32
	_ = v38
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v83 int32
	_ = v83
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v139 int32
	_ = v139
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v158 int32
	_ = v158
	v3 = int32(0)
	if l1 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	if l0 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L2:
	;
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v11 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v18 = v3
	goto L4
L4:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(v22+v18<<(uint(int32(2))%32))))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+100))
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+56))
	if v28 != 0 {
		goto L6
	} else {
		goto L7
	}
L5:
	;
	goto L1
L6:
	;
	v31 = v28
	goto L9
L7:
	;
	goto L8
L8:
	;
	v47 = *(*int32)(unsafe.Add(mBase, uint32(v26)+104))
	v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+56))
	if v48 != 0 {
		goto L12
	} else {
		goto L13
	}
L9:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+100)) = v31
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v31)+56))
	if v38 != 0 {
		v31 = v38
		goto L9
	} else {
		goto L11
	}
L10:
	;
	goto L8
L11:
	;
	goto L10
L12:
	;
	v51 = v48
	goto L15
L13:
	;
	goto L14
L14:
	;
	v68 = v18 + int32(1)
	v69 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v68 < v69 {
		v18 = v68
		goto L4
	} else {
		goto L18
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v26)+104)) = v51
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v51)+56))
	if v58 != 0 {
		v51 = v58
		goto L15
	} else {
		goto L17
	}
L16:
	;
	goto L14
L17:
	;
	goto L16
L18:
	;
	goto L5
L19:
	;
	return int32(0)
L20:
	;
	goto L21
L21:
	;
	v83 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v83 <= int32(0) {
		v158 = v3
		goto L22
	} else {
		goto L23
	}
L22:
	;
	return v158
L23:
	;
	v91 = v3
	v92 = v3
	goto L24
L24:
	;
	if l1 == int32(0) {
		v158 = v92
		goto L22
	} else {
		goto L26
	}
L25:
	;
	v158 = v146
	goto L22
L26:
	;
	v96 = int32(0)
	v97 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v96 < v97 {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v100+v91<<(uint(int32(2))%32))))
	v105 = *(*int32)(unsafe.Add(mBase, uint32(v104)+4))
	v109 = int32(0)
	v110 = v96
	goto L30
L28:
	;
	v139 = v96
	goto L29
L29:
	;
	if v139 == int32(0) {
		v158 = v92
		goto L22
	} else {
		goto L41
	}
L30:
	;
	v115 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
	v119 = *(*int32)(unsafe.Add(mBase, uint32(v115+v109<<(uint(int32(2))%32))))
	v122 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v119)+120)))
	if v122 != 0 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v139 = v131
	goto L29
L32:
	;
	v123 = int32(100)
	goto L34
L33:
	;
	v123 = int32(104)
	goto L34
L34:
	;
	v125 = *(*int32)(unsafe.Add(mBase, uint32(v119+v123)))
	if v105 == v125 {
		goto L35
	} else {
		goto L36
	}
L35:
	;
	v127 = F_lappend(m, v110, v119)
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L38
	} else {
		goto L39
	}
L36:
	;
	v131 = v110
	goto L37
L37:
	;
	v133 = v109 + int32(1)
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	if v133 < v134 {
		v109 = v133
		v110 = v131
		goto L30
	} else {
		goto L40
	}
L38:
	;
	return int32(0)
L39:
	;
	v131 = v127
	goto L37
L40:
	;
	goto L31
L41:
	;
	v146 = F_list_concat(m, v92, v139)
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L38
	} else {
		goto L42
	}
L42:
	;
	v149 = v91 + int32(1)
	v150 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v149 < v150 {
		v91 = v149
		v92 = v146
		goto L24
	} else {
		goto L43
	}
L43:
	;
	goto L25
}
func F_find_nonnullable_vars_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v39 int32
	_ = v39
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
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
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v111 int32
	_ = v111
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v131 int32
	_ = v131
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v178 int32
	_ = v178
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v254 int32
	_ = v254
	v10 = m.G0
	v12 = v10 - int32(16)
	m.G0 = v12
	if l0 == int32(0) {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	m.G0 = v12 + int32(16)
	return v254
L2:
	;
	v254 = int32(0)
	goto L1
L3:
	;
	v16 = l0
	v17 = l1
	v18 = l1
	goto L4
L4:
	;
	v25 = int32(4)
	v26 = int32(0)
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	switch v27 - int32(1) {
	case 0:
		goto L17
	case 1, 2, 3, 4, 6, 7, 8, 9, 10, 11, 12, 13, 15, 17, 18, 21, 23, 24, 25, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50:
		v254 = v26
		goto L1
	case 5:
		goto L16
	case 14:
		goto L15
	case 16:
		goto L14
	case 19:
		goto L13
	case 20:
		goto L12
	case 22:
		goto L8
	case 26, 28, 29, 30:
		v236 = v18
		v238 = v25
		goto L6
	case 27:
		goto L11
	case 51:
		goto L10
	case 52:
		goto L9
	default:
		goto L18
	}
L5:
	;
	goto L2
L6:
	;
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v16+v238)))
	if v240 != 0 {
		v16 = v240
		v17 = v236
		v18 = v236
		goto L4
	} else {
		goto L87
	}
L7:
	;
	v236 = v233
	v238 = int32(28)
	goto L6
L8:
	;
	v226 = int32(8)
	v227 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if base.B2i32(v227 == int32(2))&v17 != 0 {
		goto L83
	} else {
		goto L84
	}
L9:
	;
	if v17&int32(1) == int32(0) {
		goto L2
	} else {
		goto L80
	}
L10:
	;
	if v17&int32(1) == int32(0) {
		goto L2
	} else {
		goto L77
	}
L11:
	;
	v236 = int32(0)
	v238 = v25
	goto L6
L12:
	;
	v85 = int32(8)
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	switch v86 {
	case 0:
		goto L38
	case 1:
		v91 = v17
		goto L37
	case 2:
		v236 = int32(0)
		v238 = v85
		goto L6
	default:
		goto L36
	}
L13:
	;
	v82 = F_is_strict_saop(m, v16)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L23
	} else {
		goto L34
	}
L14:
	;
	F_set_opfuncid(m, v16)
	mBase = m.M
	v76 = m.ExcPending
	if v76 != 0 {
		goto L23
	} else {
		goto L31
	}
L15:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v71 = F_func_strict(m, v70)
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L23
	} else {
		goto L29
	}
L16:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
	if v62 != 0 {
		v254 = v26
		goto L1
	} else {
		goto L27
	}
L17:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v32 <= int32(0) {
		v254 = v26
		goto L1
	} else {
		goto L20
	}
L18:
	;
	if v27 == int32(319) {
		v236 = v17
		v238 = v25
		goto L6
	} else {
		goto L19
	}
L19:
	;
	goto L2
L20:
	;
	v39 = v26
	v40 = int32(0)
	goto L21
L21:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)+12))
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v45+v40<<(uint(int32(2))%32))))
	v52 = F_find_nonnullable_vars_walker(m, v49, v17&int32(1))
	mBase = m.M
	v55 = m.ExcPending
	if v55 != 0 {
		goto L23
	} else {
		goto L24
	}
L22:
	;
	v254 = v56
	goto L1
L23:
	;
	return int32(0)
L24:
	;
	v56 = F_mbms_add_members(m, v39, v52)
	mBase = m.M
	v57 = m.ExcPending
	if v57 != 0 {
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v59 = v40 + int32(1)
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	if v59 < v60 {
		v39 = v56
		v40 = v59
		goto L21
	} else {
		goto L26
	}
L26:
	;
	goto L22
L27:
	;
	v63 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+8)))
	v67 = F_mbms_add_member(m, v63, v64+int32(7))
	mBase = m.M
	v68 = m.ExcPending
	if v68 != 0 {
		goto L23
	} else {
		goto L28
	}
L28:
	;
	v254 = v67
	goto L1
L29:
	;
	if v71 == int32(0) {
		goto L2
	} else {
		goto L30
	}
L30:
	;
	v233 = int32(0)
	goto L7
L31:
	;
	v78 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	v79 = F_func_strict(m, v78)
	mBase = m.M
	v80 = m.ExcPending
	if v80 != 0 {
		goto L23
	} else {
		goto L32
	}
L32:
	;
	if v79 != 0 {
		v233 = int32(0)
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L2
L34:
	;
	if v82 != 0 {
		v233 = int32(0)
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L2
L36:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v191 = m.ExcPending
	if v191 != 0 {
		goto L23
	} else {
		goto L74
	}
L37:
	;
	v93 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v93 == int32(0) {
		goto L2
	} else {
		goto L40
	}
L38:
	;
	v87 = int32(1)
	if v17&v87 != 0 {
		v236 = v87
		v238 = v85
		goto L6
	} else {
		goto L39
	}
L39:
	;
	v91 = int32(0)
	goto L37
L40:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v96 <= int32(0) {
		goto L2
	} else {
		goto L41
	}
L41:
	;
	v102 = int32(0)
	v105 = v26
	goto L42
L42:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v93)+12))
	v115 = *(*int32)(unsafe.Add(mBase, uint32(v111+v102<<(uint(int32(2))%32))))
	v116 = F_find_nonnullable_vars_walker(m, v115, v91&int32(1))
	mBase = m.M
	v117 = m.ExcPending
	if v117 != 0 {
		goto L23
	} else {
		goto L44
	}
L43:
	;
	v254 = v178
	goto L1
L44:
	;
	if v105 == int32(0) {
		goto L46
	} else {
		goto L47
	}
L45:
	;
	v185 = v102 + int32(1)
	v186 = *(*int32)(unsafe.Add(mBase, uint32(v93)+4))
	if v185 < v186 {
		v102 = v185
		v105 = v178
		goto L42
	} else {
		goto L73
	}
L46:
	;
	if v116 != 0 {
		v178 = v116
		goto L45
	} else {
		goto L49
	}
L47:
	;
	goto L48
L48:
	;
	if v116 != 0 {
		goto L51
	} else {
		goto L52
	}
L49:
	;
	goto L2
L50:
	;
	if v134 == int32(0) {
		goto L2
	} else {
		goto L72
	}
L51:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	v123 = v121
	goto L53
L52:
	;
	v123 = int32(0)
	goto L53
L53:
	;
	v124 = int32(0)
	if base.B2i32(v105 == v124)|base.B2i32(v123 <= v124) != 0 {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v137 = int32(0)
	goto L61
L55:
	;
	v134 = int32(0)
	goto L57
L56:
	;
	v131 = *(*int32)(unsafe.Add(mBase, uint32(v105)+4))
	if v123 < v131 {
		goto L58
	} else {
		goto L59
	}
L57:
	;
	goto L54
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v105)+4)) = v123
	goto L60
L59:
	;
	goto L60
L60:
	;
	v134 = v105
	goto L57
L61:
	;
	v144 = int32(0)
	if v134 == v144 {
		v153 = v144
		goto L63
	} else {
		goto L64
	}
L63:
	;
	if v116 == int32(0) {
		goto L67
	} else {
		goto L68
	}
L64:
	;
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v134)+4))
	if v147 <= v137 {
		v153 = v144
		goto L63
	} else {
		goto L65
	}
L65:
	;
	v149 = *(*int32)(unsafe.Add(mBase, uint32(v134)+12))
	v153 = v149 + v137<<(uint(int32(2))%32)
	goto L63
L66:
	;
	v163 = *(*int32)(unsafe.Add(mBase, uint32(v153)))
	v167 = *(*int32)(unsafe.Add(mBase, uint32(v161+v137<<(uint(int32(2))%32))))
	v168 = F_bms_int_members(m, v163, v167)
	mBase = m.M
	v169 = m.ExcPending
	if v169 != 0 {
		goto L23
	} else {
		goto L71
	}
L67:
	;
	goto L50
L68:
	;
	v158 = *(*int32)(unsafe.Add(mBase, uint32(v116)+4))
	if base.B2i32(v153 == int32(0))|base.B2i32(v158 <= v137) != 0 {
		goto L67
	} else {
		goto L69
	}
L69:
	;
	v161 = *(*int32)(unsafe.Add(mBase, uint32(v116)+12))
	if v161 != 0 {
		goto L66
	} else {
		goto L70
	}
L70:
	;
	goto L67
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v153))) = v168
	v137 = v137 + int32(1)
	goto L61
L72:
	;
	v178 = v134
	goto L45
L73:
	;
	goto L43
L74:
	;
	v192 = *(*int32)(unsafe.Add(mBase, uint32(v16)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v12))) = v192
	F_errmsg_internal(m, int32(_a_F_find_nonnullable_vars_walker_0), v12)
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L23
	} else {
		goto L75
	}
L75:
	;
	F_errfinish(m, int32(_a_F_find_nonnullable_vars_walker_1), int32(1830), int32(_a_F_find_nonnullable_vars_walker_2))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L23
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	v207 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if v207 != int32(1) {
		goto L2
	} else {
		goto L78
	}
L78:
	;
	v210 = int32(0)
	v211 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v16)+12)))
	if v211 == v210 {
		v236 = v210
		v238 = v25
		goto L6
	} else {
		goto L79
	}
L79:
	;
	goto L2
L80:
	;
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v16)+8))
	if base.Ui32(int32(5)) < base.Ui32(v219) {
		goto L2
	} else {
		goto L81
	}
L81:
	;
	if int32(1)<<(uint(v219)%32)&int32(37) != 0 {
		v236 = int32(0)
		v238 = v25
		goto L6
	} else {
		goto L82
	}
L82:
	;
	v254 = v26
	goto L1
L83:
	;
	v236 = v17
	v238 = v226
	goto L6
L84:
	;
	goto L85
L85:
	;
	if v227 == int32(3) {
		v236 = v17
		v238 = v226
		goto L6
	} else {
		goto L86
	}
L86:
	;
	goto L2
L87:
	;
	goto L5
}
func F_find_or_make_matching_shared_tupledesc(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
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
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
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
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v244 int32
	_ = v244
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v259 int32
	_ = v259
	var v260 int64
	_ = v260
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	v2 = int32(0)
	v9 = m.G0
	v11 = v9 - int32(208)
	m.G0 = v11
	v16 = int32(-1)
	v17 = v2
	v18 = v2
	v19 = v2
	v20 = v2
	goto L2
L1:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L2:
	;
	goto L5
L3:
	;
	m.G0 = v11 + int32(208)
	return v283
L4:
	;
	goto L3
L5:
	;
	if v16 != int32(1) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v283 = v251
	goto L4
L7:
	;
	v259 = int32(m.ExcTag)
	v260 = int64(m.ExcVals[0])
	m.ExcPending = 0
	if v259 == int32(0) {
		goto L47
	} else {
		goto L48
	}
L8:
	;
	v24 = int32(0)
	v26 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+8))
	if v27 == v24 {
		v283 = v24
		goto L4
	} else {
		goto L11
	}
L9:
	;
	v103 = v17
	v104 = v18
	v105 = v19
	v106 = v20
	goto L10
L10:
	;
	if v103 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L11:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+188)) = l0
	v31 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+192)) = uint8(v31)
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v18
	v40 = F_dshash_find(m, v33, v11+int32(188), v31)
	mBase = m.M
	v41 = m.ExcPending
	if v41 != 0 {
		goto L7
	} else {
		goto L12
	}
L12:
	;
	v43 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	if v40 != 0 {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v44 = *(*int32)(unsafe.Add(mBase, uint32(v43)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v18
	F_dshash_release_lock(m, v44, v40)
	mBase = m.M
	v49 = m.ExcPending
	if v49 != 0 {
		goto L7
	} else {
		goto L16
	}
L14:
	;
	goto L15
L15:
	;
	v59 = *(*int32)(unsafe.Add(mBase, uint32(v43)+8))
	v60 = *(*int32)(unsafe.Add(mBase, uint32(v59)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v59)+8)) = v60 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v11)+180)) = v60
	v66 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+4))
	v68 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v18
	v77 = F_dsa_allocate_extended(m, v67, v68*int32(116)+int32(20), int32(0))
	mBase = m.M
	v78 = m.ExcPending
	if v78 != 0 {
		goto L7
	} else {
		goto L18
	}
L16:
	;
	v50 = *(*int32)(unsafe.Add(mBase, uint32(v40)))
	v52 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v52)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v18
	v57 = F_dsa_get_address(m, v53, v50)
	mBase = m.M
	v58 = m.ExcPending
	if v58 != 0 {
		goto L7
	} else {
		goto L17
	}
L17:
	;
	v283 = v57
	goto L4
L18:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v77
	v82 = F_dsa_get_address(m, v67, v77)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L7
	} else {
		goto L19
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v19
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v20
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v77
	F_TupleDescCopy(m, v82, l0)
	mBase = m.M
	v88 = m.ExcPending
	if v88 != 0 {
		goto L7
	} else {
		goto L20
	}
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+8)) = v60
	v92 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[1]))
	v94 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[2]))
	goto L21
L21:
	;
	v96 = v11 + int32(16)
	*(*int32)(unsafe.Add(mBase, uint32(v96)+4)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v96))) = v11 + int32(12)
	goto L24
L22:
	;
	v103 = int32(0)
	v104 = v77
	v105 = v94
	v106 = v92
	goto L10
L24:
	;
	goto L22
L25:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[1])) = v106
	*(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[2])) = v105
	v172 = *(*int32)(unsafe.Add(mBase, uint32(v11)+180))
	*(*int32)(unsafe.Add(mBase, uint32(v124)+4)) = v104
	*(*int32)(unsafe.Add(mBase, uint32(v124))) = v172
	v176 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v177 = *(*int32)(unsafe.Add(mBase, uint32(v176)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	F_dshash_release_lock(m, v177, v124)
	mBase = m.M
	v182 = m.ExcPending
	if v182 != 0 {
		goto L7
	} else {
		goto L36
	}
L26:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[2])) = v11 + int32(16)
	v115 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v116 = *(*int32)(unsafe.Add(mBase, uint32(v115)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	v124 = F_dshash_find_or_insert(m, v116, v11+int32(180), v11+int32(187))
	mBase = m.M
	v125 = m.ExcPending
	if v125 != 0 {
		goto L7
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[1])) = v106
	*(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[2])) = v105
	v156 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v157 = *(*int32)(unsafe.Add(mBase, uint32(v156)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	F_dsa_free(m, v157, v104)
	mBase = m.M
	v162 = m.ExcPending
	if v162 != 0 {
		goto L7
	} else {
		goto L34
	}
L29:
	;
	v126 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+187)))
	if v126 != int32(1) {
		goto L25
	} else {
		goto L30
	}
L30:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L7
	} else {
		goto L31
	}
L31:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	F_errmsg_internal(m, int32(_a_F_find_or_make_matching_shared_tupledesc_0), int32(0))
	mBase = m.M
	v142 = m.ExcPending
	if v142 != 0 {
		goto L7
	} else {
		goto L32
	}
L32:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	F_errfinish(m, int32(_a_F_find_or_make_matching_shared_tupledesc_1), int32(2993), int32(_a_F_find_or_make_matching_shared_tupledesc_2))
	mBase = m.M
	v150 = m.ExcPending
	if v150 != 0 {
		goto L7
	} else {
		goto L33
	}
L33:
	;
	goto L1
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	F_pg_re_throw(m)
	mBase = m.M
	v167 = m.ExcPending
	if v167 != 0 {
		goto L7
	} else {
		goto L35
	}
L35:
	;
	goto L1
L36:
	;
	v184 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(v184)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	v193 = F_dshash_find_or_insert(m, v185, v11+int32(188), v11+int32(187))
	mBase = m.M
	v194 = m.ExcPending
	if v194 != 0 {
		goto L7
	} else {
		goto L37
	}
L37:
	;
	v195 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11)+187)))
	if v195 == int32(1) {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v199 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v199)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	F_dshash_release_lock(m, v200, v193)
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L7
	} else {
		goto L41
	}
L39:
	;
	goto L40
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v193))) = v104
	v235 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v193)+4)) = uint8(v235)
	v238 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v239 = *(*int32)(unsafe.Add(mBase, uint32(v238)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	F_dshash_release_lock(m, v239, v193)
	mBase = m.M
	v244 = m.ExcPending
	if v244 != 0 {
		goto L7
	} else {
		goto L45
	}
L41:
	;
	v207 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v208 = *(*int32)(unsafe.Add(mBase, uint32(v207)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	v214 = F_dshash_delete_key(m, v208, v11+int32(180))
	mBase = m.M
	v215 = m.ExcPending
	if v215 != 0 {
		goto L7
	} else {
		goto L42
	}
L42:
	;
	*(*uint8)(unsafe.Add(mBase, uint32(v11)+187)) = uint8(v214)
	v218 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v219 = *(*int32)(unsafe.Add(mBase, uint32(v218)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	F_dsa_free(m, v219, v104)
	mBase = m.M
	v224 = m.ExcPending
	if v224 != 0 {
		goto L7
	} else {
		goto L43
	}
L43:
	;
	v225 = *(*int32)(unsafe.Add(mBase, uint32(v193)))
	v227 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v228 = *(*int32)(unsafe.Add(mBase, uint32(v227)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	v232 = F_dsa_get_address(m, v228, v225)
	mBase = m.M
	v233 = m.ExcPending
	if v233 != 0 {
		goto L7
	} else {
		goto L44
	}
L44:
	;
	v283 = v232
	goto L4
L45:
	;
	v246 = *(*int32)(unsafe.Add(mBase, _c_F_find_or_make_matching_shared_tupledesc[0]))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v246)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+196)) = v106
	*(*int32)(unsafe.Add(mBase, uint32(v11)+200)) = v105
	*(*int32)(unsafe.Add(mBase, uint32(v11)+204)) = v104
	v251 = F_dsa_get_address(m, v247, v104)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L7
	} else {
		goto L46
	}
L46:
	;
	goto L6
L47:
	;
	v264 = int32(v260)
	m.G0 = v11
	v266 = *(*int32)(unsafe.Add(mBase, uint32(v264)+4))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v264)))
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v267)))
	if v11+int32(12) == v270 {
		goto L50
	} else {
		goto L51
	}
L48:
	;
	m.ExcPending = 1
	goto L56
L49:
	;
	if v274 != 0 {
		goto L53
	} else {
		goto L54
	}
L50:
	;
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v267)+4))
	v274 = v272
	goto L52
L51:
	;
	v274 = int32(0)
	goto L52
L52:
	;
	goto L49
L53:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v11)+204))
	v276 = *(*int32)(unsafe.Add(mBase, uint32(v11)+200))
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v11)+196))
	v16 = v274
	v17 = v266
	v18 = v275
	v19 = v276
	v20 = v277
	goto L2
L54:
	;
	goto L55
L55:
	;
	F___wasm_longjmp(m, v267, v266)
	mBase = m.M
	v281 = m.ExcPending
	if v281 != 0 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	return int32(0)
L57:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_find_simplified_clause(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
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
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v75 int32
	_ = v75
	var v76 float64
	_ = v76
	var v77 float64
	_ = v77
	var v80 float64
	_ = v80
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v97 int32
	_ = v97
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
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v131 int32
	_ = v131
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v152 int32
	_ = v152
	v4 = int32(0)
	v10 = m.G0
	v12 = v10 + int32(-64)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
	if v14 != int32(7) {
		v131 = v4
		m.G0 = v12 - int32(-64)
		return v131
	} else {
		v17 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
		if v17 != 0 {
			v131 = v4
			m.G0 = v12 - int32(-64)
			return v131
		} else {
			v18 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
			v19 = F_pg_detoast_datum(m, v18)
			mBase = m.M
			v22 = m.ExcPending
			if v22 != 0 {
				return int32(0)
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
				v25 = F_lookup_type_cache(m, v23, int32(2048))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
					if v27 == int32(0) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							v143 = *(*int32)(unsafe.Add(mBase, uint32(v19)+4))
							*(*int32)(unsafe.Add(mBase, uint32(v12))) = v143
							F_errmsg_internal(m, int32(_a_F_find_simplified_clause_0), v12)
							mBase = m.M
							v147 = m.ExcPending
							if v147 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_find_simplified_clause_1), int32(2866), int32(_a_F_find_simplified_clause_2))
								mBase = m.M
								v152 = m.ExcPending
								if v152 != 0 {
									return int32(0)
								} else {
									base.Wasm_trap_unreachable()
									for {
									}
								}
							}
						}
					} else {
						F_range_deserialize(m, v25, v19, v10+int32(-8), v10+int32(-16), v10+int32(-17))
						mBase = m.M
						v37 = m.ExcPending
						if v37 != 0 {
							return int32(0)
						} else {
							v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+47)))
							if v38 == int32(1) {
								v41 = int32(0)
								v43 = F_makeBoolConst(m, v41, v41)
								mBase = m.M
								v44 = m.ExcPending
								if v44 != 0 {
									return int32(0)
								} else {
									v131 = v43
									m.G0 = v12 - int32(-64)
									return v131
								}
							} else {
								v45 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+52)))
								v46 = int32(1)
								v48 = int32(0)
								v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+60)))
								if base.B2i32(v45&v46 == v48)|base.B2i32(v50 != v46) == v48 {
									v58 = F_makeBoolConst(m, int32(1), int32(0))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v131 = v58
										m.G0 = v12 - int32(-64)
										return v131
									}
								} else {
									v60 = *(*int32)(unsafe.Add(mBase, uint32(v25)+208))
									v61 = *(*int32)(unsafe.Add(mBase, uint32(v25)+204))
									v62 = *(*int32)(unsafe.Add(mBase, uint32(v25)+200))
									if v50|v45&int32(1) == int32(0) {
										v68 = F_contain_volatile_functions(m, l2)
										mBase = m.M
										v69 = m.ExcPending
										if v69 != 0 {
											return int32(0)
										} else {
											if v68 != 0 {
												v131 = v4
												m.G0 = v12 - int32(-64)
												return v131
											} else {
												v70 = F_contain_subplans(m, l2)
												mBase = m.M
												v71 = m.ExcPending
												if v71 != 0 {
													return int32(0)
												} else {
													if v70 != 0 {
														v131 = v4
														m.G0 = v12 - int32(-64)
														return v131
													} else {
														F_cost_qual_eval_node(m, v10+int32(-40), l2, l0)
														mBase = m.M
														v75 = m.ExcPending
														if v75 != 0 {
															return int32(0)
														} else {
															v76 = *(*float64)(unsafe.Add(mBase, uint32(v12)+24))
															v77 = *(*float64)(unsafe.Add(mBase, uint32(v12)+32))
															v80 = *(*float64)(unsafe.Add(mBase, _c_F_find_simplified_clause[0]))
															if base.F64_gt(base.F64_add(v76, v77), base.F64_mul(v80, float64(10))) == int32(0) {
																v86 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
																v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+61)))
																v89 = F_build_bound_expr(m, l2, v86, int32(1), v88, v62, v61, v60)
																mBase = m.M
																v90 = m.ExcPending
																if v90 != 0 {
																	return int32(0)
																} else {
																	if v89 == int32(0) {
																		v131 = v89
																		m.G0 = v12 - int32(-64)
																		return v131
																	} else {
																		if v45&int32(1) == int32(0) {
																			v102 = F_copyObjectImpl(m, l2)
																			mBase = m.M
																			v103 = m.ExcPending
																			if v103 != 0 {
																				return int32(0)
																			} else {
																				v104 = v89
																				v105 = v102
																				v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
																				v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+53)))
																				v109 = F_build_bound_expr(m, v105, v106, int32(0), v108, v62, v61, v60)
																				mBase = m.M
																				v110 = m.ExcPending
																				if v110 != 0 {
																					return int32(0)
																				} else {
																					if v109 == int32(0) {
																						v131 = v4
																						m.G0 = v12 - int32(-64)
																						return v131
																					} else {
																						if v104 == int32(0) {
																							v131 = v109
																							m.G0 = v12 - int32(-64)
																							return v131
																						} else {
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v109
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v104
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v104
																							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v109
																							v123 = F_list_make2_impl(m, v10+int32(-52), v10+int32(-56))
																							mBase = m.M
																							v124 = m.ExcPending
																							if v124 != 0 {
																								return int32(0)
																							} else {
																								v125 = F_make_andclause(m, v123)
																								mBase = m.M
																								v126 = m.ExcPending
																								if v126 != 0 {
																									return int32(0)
																								} else {
																									v131 = v125
																									m.G0 = v12 - int32(-64)
																									return v131
																								}
																							}
																						}
																					}
																				}
																			}
																		} else {
																			v131 = v89
																			m.G0 = v12 - int32(-64)
																			return v131
																		}
																	}
																}
															} else {
																v131 = v4
																m.G0 = v12 - int32(-64)
																return v131
															}
														}
													}
												}
											}
										}
									} else {
										if v50 != 0 {
											v97 = int32(0)
											if v45&int32(1) == v97 {
												v104 = v97
												v105 = l2
												v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
												v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+53)))
												v109 = F_build_bound_expr(m, v105, v106, int32(0), v108, v62, v61, v60)
												mBase = m.M
												v110 = m.ExcPending
												if v110 != 0 {
													return int32(0)
												} else {
													if v109 == int32(0) {
														v131 = v4
														m.G0 = v12 - int32(-64)
														return v131
													} else {
														if v104 == int32(0) {
															v131 = v109
															m.G0 = v12 - int32(-64)
															return v131
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v109
															*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v104
															*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v104
															*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v109
															v123 = F_list_make2_impl(m, v10+int32(-52), v10+int32(-56))
															mBase = m.M
															v124 = m.ExcPending
															if v124 != 0 {
																return int32(0)
															} else {
																v125 = F_make_andclause(m, v123)
																mBase = m.M
																v126 = m.ExcPending
																if v126 != 0 {
																	return int32(0)
																} else {
																	v131 = v125
																	m.G0 = v12 - int32(-64)
																	return v131
																}
															}
														}
													}
												}
											} else {
												v131 = v4
												m.G0 = v12 - int32(-64)
												return v131
											}
										} else {
											v86 = *(*int32)(unsafe.Add(mBase, uint32(v12)+56))
											v88 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+61)))
											v89 = F_build_bound_expr(m, l2, v86, int32(1), v88, v62, v61, v60)
											mBase = m.M
											v90 = m.ExcPending
											if v90 != 0 {
												return int32(0)
											} else {
												if v89 == int32(0) {
													v131 = v89
													m.G0 = v12 - int32(-64)
													return v131
												} else {
													if v45&int32(1) == int32(0) {
														v102 = F_copyObjectImpl(m, l2)
														mBase = m.M
														v103 = m.ExcPending
														if v103 != 0 {
															return int32(0)
														} else {
															v104 = v89
															v105 = v102
															v106 = *(*int32)(unsafe.Add(mBase, uint32(v12)+48))
															v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+53)))
															v109 = F_build_bound_expr(m, v105, v106, int32(0), v108, v62, v61, v60)
															mBase = m.M
															v110 = m.ExcPending
															if v110 != 0 {
																return int32(0)
															} else {
																if v109 == int32(0) {
																	v131 = v4
																	m.G0 = v12 - int32(-64)
																	return v131
																} else {
																	if v104 == int32(0) {
																		v131 = v109
																		m.G0 = v12 - int32(-64)
																		return v131
																	} else {
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v109
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v104
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v104
																		*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v109
																		v123 = F_list_make2_impl(m, v10+int32(-52), v10+int32(-56))
																		mBase = m.M
																		v124 = m.ExcPending
																		if v124 != 0 {
																			return int32(0)
																		} else {
																			v125 = F_make_andclause(m, v123)
																			mBase = m.M
																			v126 = m.ExcPending
																			if v126 != 0 {
																				return int32(0)
																			} else {
																				v131 = v125
																				m.G0 = v12 - int32(-64)
																				return v131
																			}
																		}
																	}
																}
															}
														}
													} else {
														v131 = v89
														m.G0 = v12 - int32(-64)
														return v131
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
	}
}
func F_find_wordentry(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v147 int32
	_ = v147
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v206 int32
	_ = v206
	v5 = int32(0)
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v5
	v15 = l0 + int32(8)
	v18 = v15 + v11<<(uint(int32(2))%32)
	if v11 <= v5 {
		goto L2
	} else {
		goto L3
	}
L1:
	;
	v110 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l2)+2)))
	if v110 != int32(1) {
		goto L45
	} else {
		goto L46
	}
L2:
	;
	v104 = v18
	v105 = v15
	v106 = v18
	goto L1
L3:
	;
	goto L4
L4:
	;
	v28 = v15
	v29 = v18
	goto L5
L5:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v34 = int32(12)
	v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v42 = v37 & int32(4095)
	v43 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v44 = int32(2)
	v51 = base.I32_div_s((v29-v28)>>(uint(v44)%32), v44)
	v54 = v28 + v51<<(uint(v44)%32)
	v55 = *(*int32)(unsafe.Add(mBase, uint32(v54)))
	v62 = int32(base.Ui32(v55)>>(uint(int32(1))%32)) & int32(2047)
	if v42 == int32(0) {
		goto L8
	} else {
		goto L9
	}
L6:
	;
	v104 = v54
	v105 = v97
	v106 = v98
	goto L1
L7:
	;
	if v88 == int32(0) {
		goto L35
	} else {
		goto L36
	}
L8:
	;
	goto L12
L9:
	;
	goto L10
L10:
	;
	if v62 == int32(0) {
		goto L17
	} else {
		goto L18
	}
L12:
	;
	goto L13
L13:
	;
	v68 = int32(0)
	if v68 < v62 {
		goto L14
	} else {
		goto L15
	}
L14:
	;
	v71 = int32(-1)
	goto L16
L15:
	;
	v71 = v68
	goto L16
L16:
	;
	v88 = v71
	goto L7
L17:
	;
	v88 = base.B2i32(int32(0) < v42)
	goto L7
L18:
	;
	goto L19
L19:
	;
	if base.Ui32(v42) < base.Ui32(v62) {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	v77 = v42
	goto L22
L21:
	;
	v77 = v62
	goto L22
L22:
	;
	v78 = F_memcmp(m, l1+int32(8)+v33*v34+int32(base.Ui32(v37)>>(uint(v34)%32)), v15+v43<<(uint(v44)%32)+int32(base.Ui32(v55)>>(uint(v34)%32)), v77)
	mBase = m.M
	goto L25
L23:
	;
	v88 = v86
	goto L7
L25:
	;
	goto L26
L26:
	;
	if v78 != 0 {
		v86 = v78
		goto L23
	} else {
		goto L28
	}
L28:
	;
	if v42 == v62 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v88 = int32(0)
	goto L7
L30:
	;
	goto L31
L31:
	;
	if v42 < v62 {
		goto L32
	} else {
		goto L33
	}
L32:
	;
	v85 = int32(-1)
	goto L34
L33:
	;
	v85 = int32(1)
	goto L34
L34:
	;
	v86 = v85
	goto L23
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(1)
	v104 = v54
	v105 = v28
	v106 = v54
	goto L1
L36:
	;
	goto L37
L37:
	;
	v96 = base.B2i32(int32(0) < v88)
	if int32(0) < v88 {
		goto L38
	} else {
		goto L39
	}
L38:
	;
	v97 = v54 + int32(4)
	goto L40
L39:
	;
	v97 = v28
	goto L40
L40:
	;
	if int32(0) < v88 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v98 = v29
	goto L43
L42:
	;
	v98 = v54
	goto L43
L43:
	;
	if base.Ui32(v97) < base.Ui32(v98) {
		v28 = v97
		v29 = v98
		goto L5
	} else {
		goto L44
	}
L44:
	;
	goto L6
L45:
	;
	v202 = int32(0)
	v203 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	if v202 < v203 {
		goto L83
	} else {
		goto L84
	}
L46:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = int32(0)
	if base.Ui32(v105) < base.Ui32(v106) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v116 = v104
	goto L49
L48:
	;
	v116 = v106
	goto L49
L49:
	;
	v117 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v15+v117<<(uint(int32(2))%32)) <= base.Ui32(v116) {
		goto L45
	} else {
		goto L50
	}
L50:
	;
	v128 = v117
	v129 = v116
	goto L51
L51:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v135 = int32(12)
	v138 = *(*int32)(unsafe.Add(mBase, uint32(l2)+8))
	v143 = v138 & int32(4095)
	v147 = *(*int32)(unsafe.Add(mBase, uint32(v129)))
	v154 = int32(base.Ui32(v147)>>(uint(int32(1))%32)) & int32(2047)
	if v143 == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	goto L45
L53:
	;
	if v180 != 0 {
		goto L45
	} else {
		goto L81
	}
L54:
	;
	goto L57
L55:
	;
	goto L56
L56:
	;
	if v154 == int32(0) {
		goto L63
	} else {
		goto L64
	}
L57:
	;
	v180 = int32(0)
	goto L53
L63:
	;
	v180 = base.B2i32(int32(0) < v143)
	goto L53
L64:
	;
	goto L65
L65:
	;
	if base.Ui32(v143) < base.Ui32(v154) {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v169 = v143
	goto L68
L67:
	;
	v169 = v154
	goto L68
L68:
	;
	v170 = F_memcmp(m, l1+int32(8)+v134*v135+int32(base.Ui32(v138)>>(uint(v135)%32)), v15+v128<<(uint(int32(2))%32)+int32(base.Ui32(v147)>>(uint(v135)%32)), v169)
	mBase = m.M
	goto L70
L69:
	;
	v180 = v170
	goto L53
L70:
	;
	if v170 != 0 {
		goto L69
	} else {
		goto L73
	}
L73:
	;
	v180 = base.B2i32(v154 < v143)
	goto L53
L81:
	;
	v181 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	*(*int32)(unsafe.Add(mBase, uint32(l3))) = v181 + int32(1)
	v186 = v129 + int32(4)
	v187 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if base.Ui32(v186) < base.Ui32(v15+v187<<(uint(int32(2))%32)) {
		v128 = v187
		v129 = v186
		goto L51
	} else {
		goto L82
	}
L82:
	;
	goto L52
L83:
	;
	v206 = v106
	goto L85
L84:
	;
	v206 = v202
	goto L85
L85:
	;
	return v206
}
func F_finish_spin_delay(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	v3 = int32(_a_F_finish_spin_delay_0)
	v4 = *(*int32)(unsafe.Add(mBase, _c_F_finish_spin_delay[0]))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v6 == int32(0) {
		if int32(999) < v4 {
		} else {
			v11 = int32(900)
			if v11 <= v4 {
				v14 = v11
			} else {
				v14 = v4
			}
			v21 = v14 + int32(100)
			*(*int32)(unsafe.Add(mBase, _c_F_finish_spin_delay[0])) = v21
		}
	} else {
		if v4 < int32(11) {
		} else {
			v21 = v4 - int32(1)
			*(*int32)(unsafe.Add(mBase, _c_F_finish_spin_delay[0])) = v21
		}
	}
	return
}
func F_finnish_ISO_8859_1_stem(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v54 int32
	_ = v54
	var v62 int32
	_ = v62
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v89 int32
	_ = v89
	var v91 int32
	_ = v91
	var v93 int32
	_ = v93
	var v99 int32
	_ = v99
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v187 int32
	_ = v187
	var v193 int32
	_ = v193
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v237 int32
	_ = v237
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v244 int32
	_ = v244
	var v250 int32
	_ = v250
	var v253 int32
	_ = v253
	var v258 int32
	_ = v258
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v296 int32
	_ = v296
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v358 int32
	_ = v358
	var v359 int32
	_ = v359
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	var v369 int32
	_ = v369
	var v372 int32
	_ = v372
	var v375 int32
	_ = v375
	var v379 int32
	_ = v379
	var v382 int32
	_ = v382
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v397 int32
	_ = v397
	var v399 int32
	_ = v399
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v413 int32
	_ = v413
	var v415 int32
	_ = v415
	var v417 int32
	_ = v417
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v453 int32
	_ = v453
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v459 int32
	_ = v459
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v485 int32
	_ = v485
	var v489 int32
	_ = v489
	var v496 int32
	_ = v496
	var v500 int32
	_ = v500
	var v507 int32
	_ = v507
	var v511 int32
	_ = v511
	var v518 int32
	_ = v518
	var v522 int32
	_ = v522
	var v529 int32
	_ = v529
	var v533 int32
	_ = v533
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v545 int32
	_ = v545
	var v546 int32
	_ = v546
	var v548 int32
	_ = v548
	var v550 int32
	_ = v550
	var v553 int32
	_ = v553
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v573 int32
	_ = v573
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v597 int32
	_ = v597
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v604 int32
	_ = v604
	var v610 int32
	_ = v610
	var v624 int32
	_ = v624
	var v628 int32
	_ = v628
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v655 int32
	_ = v655
	var v657 int32
	_ = v657
	var v663 int32
	_ = v663
	var v677 int32
	_ = v677
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v687 int32
	_ = v687
	var v693 int32
	_ = v693
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v704 int32
	_ = v704
	var v709 int32
	_ = v709
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v716 int32
	_ = v716
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v722 int32
	_ = v722
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v730 int32
	_ = v730
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
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
	var v750 int32
	_ = v750
	var v753 int32
	_ = v753
	var v757 int32
	_ = v757
	var v768 int32
	_ = v768
	var v769 int32
	_ = v769
	var v775 int32
	_ = v775
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v787 int32
	_ = v787
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v795 int32
	_ = v795
	var v796 int32
	_ = v796
	var v799 int32
	_ = v799
	var v802 int32
	_ = v802
	var v806 int32
	_ = v806
	var v812 int32
	_ = v812
	var v815 int32
	_ = v815
	var v825 int32
	_ = v825
	var v837 int32
	_ = v837
	var v838 int32
	_ = v838
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v850 int32
	_ = v850
	var v864 int32
	_ = v864
	var v868 int32
	_ = v868
	var v871 int32
	_ = v871
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v881 int32
	_ = v881
	var v882 int32
	_ = v882
	var v883 int32
	_ = v883
	var v890 int32
	_ = v890
	var v894 int32
	_ = v894
	var v899 int32
	_ = v899
	var v900 int32
	_ = v900
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
	var v918 int32
	_ = v918
	var v921 int32
	_ = v921
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v933 int32
	_ = v933
	var v934 int32
	_ = v934
	var v937 int32
	_ = v937
	var v940 int32
	_ = v940
	var v947 int32
	_ = v947
	var v954 int32
	_ = v954
	var v958 int32
	_ = v958
	var v959 int32
	_ = v959
	var v961 int32
	_ = v961
	var v963 int32
	_ = v963
	var v964 int32
	_ = v964
	var v967 int32
	_ = v967
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v977 int32
	_ = v977
	var v980 int32
	_ = v980
	var v981 int32
	_ = v981
	var v984 int32
	_ = v984
	var v985 int32
	_ = v985
	var v987 int32
	_ = v987
	var v999 int32
	_ = v999
	var v1011 int32
	_ = v1011
	var v1012 int32
	_ = v1012
	var v1016 int32
	_ = v1016
	var v1018 int32
	_ = v1018
	var v1024 int32
	_ = v1024
	var v1038 int32
	_ = v1038
	var v1042 int32
	_ = v1042
	var v1043 int32
	_ = v1043
	var v1054 int32
	_ = v1054
	var v1066 int32
	_ = v1066
	var v1067 int32
	_ = v1067
	var v1071 int32
	_ = v1071
	var v1073 int32
	_ = v1073
	var v1079 int32
	_ = v1079
	var v1093 int32
	_ = v1093
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
	var v1107 int32
	_ = v1107
	var v1109 int32
	_ = v1109
	var v1110 int32
	_ = v1110
	var v1113 int32
	_ = v1113
	var v1117 int32
	_ = v1117
	var v1123 int32
	_ = v1123
	var v1129 int32
	_ = v1129
	var v1135 int32
	_ = v1135
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1141 int32
	_ = v1141
	var v1142 int32
	_ = v1142
	var v1144 int32
	_ = v1144
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1154 int32
	_ = v1154
	var v1158 int32
	_ = v1158
	var v1164 int32
	_ = v1164
	var v1170 int32
	_ = v1170
	var v1171 int32
	_ = v1171
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1181 int32
	_ = v1181
	var v1197 int32
	_ = v1197
	var v1203 int32
	_ = v1203
	var v1204 int32
	_ = v1204
	var v1208 int32
	_ = v1208
	var v1210 int32
	_ = v1210
	var v1216 int32
	_ = v1216
	var v1225 int32
	_ = v1225
	var v1234 int32
	_ = v1234
	var v1237 int32
	_ = v1237
	var v1248 int32
	_ = v1248
	var v1260 int32
	_ = v1260
	var v1261 int32
	_ = v1261
	var v1265 int32
	_ = v1265
	var v1267 int32
	_ = v1267
	var v1273 int32
	_ = v1273
	var v1287 int32
	_ = v1287
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1296 int32
	_ = v1296
	var v1297 int32
	_ = v1297
	var v1298 int32
	_ = v1298
	var v1303 int32
	_ = v1303
	var v1308 int32
	_ = v1308
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1313 int32
	_ = v1313
	var v1316 int32
	_ = v1316
	var v1320 int32
	_ = v1320
	var v1324 int32
	_ = v1324
	var v1325 int32
	_ = v1325
	var v1328 int32
	_ = v1328
	var v1329 int32
	_ = v1329
	var v1342 int32
	_ = v1342
	var v1345 int32
	_ = v1345
	var v1348 int32
	_ = v1348
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v10))) = v11
	*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = v11
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v22 < v21 {
		goto L3
	} else {
		goto L4
	}
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v9
	v237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v237)+8)) = int32(0)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v240
	v242 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v242
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v237)+4))
	if v242 < v244 {
		goto L65
	} else {
		goto L66
	}
L2:
	;
	if v62 < int32(0) {
		goto L1
	} else {
		goto L17
	}
L3:
	;
	v24 = v21
	goto L5
L4:
	;
	v24 = v22
	goto L5
L5:
	;
	v31 = v21
	goto L7
L6:
	;
	v62 = v42
	goto L2
L7:
	;
	if v31 == v24 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v62 = int32(-1)
	goto L2
L10:
	;
	goto L11
L11:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v35+v31))))
	if int32(246) < v37 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v54 = v31 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v54
	v31 = v54
	goto L7
L13:
	;
	v39 = v37 - int32(97)
	if v39 < int32(0) {
		goto L12
	} else {
		goto L14
	}
L14:
	;
	v42 = int32(1)
	v46 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v39)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v46)>>(uint(v39&int32(7))%32))&v42 != 0 {
		goto L6
	} else {
		goto L15
	}
L15:
	;
	goto L12
L17:
	;
	v73 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v74 < v73 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	if v117 < int32(0) {
		goto L1
	} else {
		goto L32
	}
L19:
	;
	v76 = v73
	goto L21
L20:
	;
	v76 = v74
	goto L21
L21:
	;
	v82 = v73
	goto L23
L22:
	;
	v117 = int32(1)
	goto L18
L23:
	;
	if v82 == v76 {
		goto L25
	} else {
		goto L26
	}
L25:
	;
	v117 = int32(-1)
	goto L18
L26:
	;
	goto L27
L27:
	;
	v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v91 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89+v82))))
	if int32(246) < v91 {
		goto L22
	} else {
		goto L28
	}
L28:
	;
	v93 = v91 - int32(97)
	if v93 < int32(0) {
		goto L22
	} else {
		goto L29
	}
L29:
	;
	v99 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v93)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v99)>>(uint(v93&int32(7))%32))&int32(1) == int32(0) {
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v108 = v82 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v108
	v82 = v108
	goto L23
L32:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v121 = v120 + v117
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v121
	v123 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v123)+4)) = v121
	v132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v133 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v133 < v132 {
		goto L34
	} else {
		goto L35
	}
L33:
	;
	if v173 < int32(0) {
		goto L1
	} else {
		goto L48
	}
L34:
	;
	v135 = v132
	goto L36
L35:
	;
	v135 = v133
	goto L36
L36:
	;
	v142 = v132
	goto L38
L37:
	;
	v173 = v153
	goto L33
L38:
	;
	if v142 == v135 {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v173 = int32(-1)
	goto L33
L41:
	;
	goto L42
L42:
	;
	v146 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v146+v142))))
	if int32(246) < v148 {
		goto L43
	} else {
		goto L44
	}
L43:
	;
	v165 = v142 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v165
	v142 = v165
	goto L38
L44:
	;
	v150 = v148 - int32(97)
	if v150 < int32(0) {
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v153 = int32(1)
	v157 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v150)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v157)>>(uint(v150&int32(7))%32))&v153 != 0 {
		goto L37
	} else {
		goto L46
	}
L46:
	;
	goto L43
L48:
	;
	v184 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v185 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v185 < v184 {
		goto L50
	} else {
		goto L51
	}
L49:
	;
	if v228 < int32(0) {
		goto L1
	} else {
		goto L63
	}
L50:
	;
	v187 = v184
	goto L52
L51:
	;
	v187 = v185
	goto L52
L52:
	;
	v193 = v184
	goto L54
L53:
	;
	v228 = int32(1)
	goto L49
L54:
	;
	if v193 == v187 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v228 = int32(-1)
	goto L49
L57:
	;
	goto L58
L58:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v202 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v200+v193))))
	if int32(246) < v202 {
		goto L53
	} else {
		goto L59
	}
L59:
	;
	v204 = v202 - int32(97)
	if v204 < int32(0) {
		goto L53
	} else {
		goto L60
	}
L60:
	;
	v210 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v204)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v210)>>(uint(v204&int32(7))%32))&int32(1) == int32(0) {
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v219 = v193 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v219
	v193 = v219
	goto L54
L63:
	;
	v231 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v232 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v231))) = v232 + v228
	goto L1
L64:
	;
	return v1348
L65:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v326
	v328 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v329 = *(*int32)(unsafe.Add(mBase, uint32(v328)+4))
	if v326 < v329 {
		goto L90
	} else {
		goto L91
	}
L66:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v242
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v244
	v250 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_0), int32(10))
	mBase = m.M
	v253 = m.ExcPending
	if v253 != 0 {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	return int32(0)
L68:
	;
	if v250 == int32(0) {
		goto L69
	} else {
		goto L70
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v240
	goto L65
L70:
	;
	goto L71
L71:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v240
	v258 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v258
	switch v250 - int32(1) {
	case 0:
		goto L74
	case 1:
		goto L73
	default:
		goto L72
	}
L72:
	;
	v320 = F_slice_del(m, l0)
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L67
	} else {
		goto L88
	}
L73:
	;
	v317 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v317)))
	if v258 < v318 {
		goto L65
	} else {
		goto L87
	}
L74:
	;
	v270 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L77
L75:
	;
	if v314 == int32(0) {
		goto L72
	} else {
		goto L86
	}
L76:
	;
	v314 = v310
	goto L75
L77:
	;
	if v270 <= v271 {
		goto L79
	} else {
		goto L80
	}
L78:
	;
	v310 = int32(0)
	goto L76
L79:
	;
	v314 = int32(-1)
	goto L75
L80:
	;
	goto L81
L81:
	;
	v283 = int32(1)
	v284 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v288 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v284+v270-v283))))
	if int32(246) < v288 {
		v310 = v283
		goto L76
	} else {
		goto L82
	}
L82:
	;
	v290 = v288 - int32(97)
	if v290 < int32(0) {
		v310 = v283
		goto L76
	} else {
		goto L83
	}
L83:
	;
	v296 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v290)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[1]))))
	if int32(base.Ui32(v296)>>(uint(v290&int32(7))%32))&int32(1) == int32(0) {
		v310 = v283
		goto L76
	} else {
		goto L84
	}
L84:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v270 - int32(1)
	goto L85
L85:
	;
	goto L78
L86:
	;
	goto L65
L87:
	;
	goto L72
L88:
	;
	if v320 < int32(0) {
		v1348 = v320
		goto L64
	} else {
		goto L89
	}
L89:
	;
	goto L65
L90:
	;
	v453 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v453
	v455 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v456 = *(*int32)(unsafe.Add(mBase, uint32(v455)+4))
	if v453 < v456 {
		goto L137
	} else {
		goto L138
	}
L91:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v326
	v332 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v329
	v336 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_1), int32(9))
	mBase = m.M
	v337 = m.ExcPending
	if v337 != 0 {
		goto L67
	} else {
		goto L92
	}
L92:
	;
	if v336 == int32(0) {
		goto L93
	} else {
		goto L94
	}
L93:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v332
	goto L90
L94:
	;
	goto L95
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v332
	v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v342
	switch v336 - int32(1) {
	case 0:
		goto L101
	case 1:
		goto L100
	case 2:
		goto L99
	case 3:
		goto L98
	case 4:
		goto L97
	case 5:
		goto L96
	default:
		goto L90
	}
L96:
	;
	if v342-int32(2) <= v332 {
		goto L90
	} else {
		goto L131
	}
L97:
	;
	v413 = v342 - int32(1)
	if v413 <= v332 {
		goto L90
	} else {
		goto L125
	}
L98:
	;
	v395 = v342 - int32(1)
	if v395 <= v332 {
		goto L90
	} else {
		goto L119
	}
L99:
	;
	v390 = F_slice_del(m, l0)
	mBase = m.M
	v391 = m.ExcPending
	if v391 != 0 {
		goto L67
	} else {
		goto L117
	}
L100:
	;
	v358 = F_slice_del(m, l0)
	mBase = m.M
	v359 = m.ExcPending
	if v359 != 0 {
		goto L67
	} else {
		goto L108
	}
L101:
	;
	if v332 < v342 {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v347 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v347+v342-int32(1)))))
	if v351 == int32(107) {
		goto L90
	} else {
		goto L105
	}
L103:
	;
	goto L104
L104:
	;
	v354 = F_slice_del(m, l0)
	mBase = m.M
	v355 = m.ExcPending
	if v355 != 0 {
		goto L67
	} else {
		goto L106
	}
L105:
	;
	goto L104
L106:
	;
	if int32(0) <= v354 {
		goto L90
	} else {
		goto L107
	}
L107:
	;
	v1348 = v354
	goto L64
L108:
	;
	if v358 < int32(0) {
		v1348 = v358
		goto L64
	} else {
		goto L109
	}
L109:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v362
	v364 = int32(3)
	v366 = int32(0)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v362-v369 < v364 {
		v379 = v366
		goto L111
	} else {
		goto L112
	}
L110:
	;
	if v379 == int32(0) {
		goto L90
	} else {
		goto L114
	}
L111:
	;
	goto L110
L112:
	;
	v372 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v375 = F_memcmp(m, v372+v362-v364, int32(_a_F_finnish_ISO_8859_1_stem_2), v364)
	mBase = m.M
	if v375 != 0 {
		v379 = v366
		goto L111
	} else {
		goto L113
	}
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v362 - v364
	v379 = int32(1)
	goto L111
L114:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v382
	v386 = F_slice_from_s(m, l0, int32(3), int32(_a_F_finnish_ISO_8859_1_stem_3))
	mBase = m.M
	v387 = m.ExcPending
	if v387 != 0 {
		goto L67
	} else {
		goto L115
	}
L115:
	;
	if int32(0) <= v386 {
		goto L90
	} else {
		goto L116
	}
L116:
	;
	v1348 = v386
	goto L64
L117:
	;
	if int32(0) <= v390 {
		goto L90
	} else {
		goto L118
	}
L118:
	;
	v1348 = v390
	goto L64
L119:
	;
	v397 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v399 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v397+v395))))
	if v399 != int32(97) {
		goto L90
	} else {
		goto L120
	}
L120:
	;
	v404 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_4), int32(6))
	mBase = m.M
	v405 = m.ExcPending
	if v405 != 0 {
		goto L67
	} else {
		goto L121
	}
L121:
	;
	if v404 == int32(0) {
		goto L90
	} else {
		goto L122
	}
L122:
	;
	v408 = F_slice_del(m, l0)
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L67
	} else {
		goto L123
	}
L123:
	;
	if int32(0) <= v408 {
		goto L90
	} else {
		goto L124
	}
L124:
	;
	v1348 = v408
	goto L64
L125:
	;
	v415 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v417 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v415+v413))))
	if v417 != int32(228) {
		goto L90
	} else {
		goto L126
	}
L126:
	;
	v422 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_5), int32(6))
	mBase = m.M
	v423 = m.ExcPending
	if v423 != 0 {
		goto L67
	} else {
		goto L127
	}
L127:
	;
	if v422 == int32(0) {
		goto L90
	} else {
		goto L128
	}
L128:
	;
	v426 = F_slice_del(m, l0)
	mBase = m.M
	v427 = m.ExcPending
	if v427 != 0 {
		goto L67
	} else {
		goto L129
	}
L129:
	;
	if int32(0) <= v426 {
		goto L90
	} else {
		goto L130
	}
L130:
	;
	v1348 = v426
	goto L64
L131:
	;
	v433 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v437 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v433+v342-int32(1)))))
	if v437 != int32(101) {
		goto L90
	} else {
		goto L132
	}
L132:
	;
	v442 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_6), int32(2))
	mBase = m.M
	v443 = m.ExcPending
	if v443 != 0 {
		goto L67
	} else {
		goto L133
	}
L133:
	;
	if v442 == int32(0) {
		goto L90
	} else {
		goto L134
	}
L134:
	;
	v446 = F_slice_del(m, l0)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L67
	} else {
		goto L135
	}
L135:
	;
	if v446 < int32(0) {
		v1348 = v446
		goto L64
	} else {
		goto L136
	}
L136:
	;
	goto L90
L137:
	;
	v693 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v693
	v695 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v695)))
	if v693 < v696 {
		goto L201
	} else {
		goto L202
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v453
	v459 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v456
	v463 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_7), int32(30))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L67
	} else {
		goto L139
	}
L139:
	;
	if v463 == int32(0) {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v459
	goto L137
L141:
	;
	goto L142
L142:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v459
	v469 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v469
	switch v463 - int32(1) {
	case 0:
		goto L151
	case 1:
		goto L150
	case 2:
		goto L149
	case 3:
		goto L148
	case 4:
		goto L147
	case 5:
		goto L146
	case 6:
		goto L145
	case 7:
		goto L144
	default:
		goto L143
	}
L143:
	;
	v683 = F_slice_del(m, l0)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L67
	} else {
		goto L199
	}
L144:
	;
	v584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L177
L145:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v540 = v539 - v469
	v543 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_8), int32(7))
	mBase = m.M
	v544 = m.ExcPending
	if v544 != 0 {
		goto L67
	} else {
		goto L165
	}
L146:
	;
	if v469 <= v459 {
		goto L137
	} else {
		goto L162
	}
L147:
	;
	if v469 <= v459 {
		goto L137
	} else {
		goto L160
	}
L148:
	;
	if v469 <= v459 {
		goto L137
	} else {
		goto L158
	}
L149:
	;
	if v469 <= v459 {
		goto L137
	} else {
		goto L156
	}
L150:
	;
	if v469 <= v459 {
		goto L137
	} else {
		goto L154
	}
L151:
	;
	if v469 <= v459 {
		goto L137
	} else {
		goto L152
	}
L152:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v478 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v474+v469-int32(1)))))
	if v478 != int32(97) {
		goto L137
	} else {
		goto L153
	}
L153:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v469 - int32(1)
	goto L143
L154:
	;
	v485 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v489 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v485+v469-int32(1)))))
	if v489 != int32(101) {
		goto L137
	} else {
		goto L155
	}
L155:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v469 - int32(1)
	goto L143
L156:
	;
	v496 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v500 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v496+v469-int32(1)))))
	if v500 != int32(105) {
		goto L137
	} else {
		goto L157
	}
L157:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v469 - int32(1)
	goto L143
L158:
	;
	v507 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v511 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v507+v469-int32(1)))))
	if v511 != int32(111) {
		goto L137
	} else {
		goto L159
	}
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v469 - int32(1)
	goto L143
L160:
	;
	v518 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v522 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v518+v469-int32(1)))))
	if v522 != int32(228) {
		goto L137
	} else {
		goto L161
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v469 - int32(1)
	goto L143
L162:
	;
	v529 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v533 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v529+v469-int32(1)))))
	if v533 != int32(246) {
		goto L137
	} else {
		goto L163
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v469 - int32(1)
	goto L143
L164:
	;
	v567 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v568 = v567 - v540
	v569 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v568 <= v569 {
		goto L172
	} else {
		goto L173
	}
L165:
	;
	if v543 != 0 {
		goto L164
	} else {
		goto L166
	}
L166:
	;
	v545 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v546 = v545 - v540
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v546
	v548 = int32(2)
	v550 = int32(0)
	v553 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v546-v553 < v548 {
		v563 = v550
		goto L168
	} else {
		goto L169
	}
L167:
	;
	if v563 != 0 {
		goto L164
	} else {
		goto L171
	}
L168:
	;
	goto L167
L169:
	;
	v556 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v559 = F_memcmp(m, v556+v546-v548, int32(_a_F_finnish_ISO_8859_1_stem_9), v548)
	mBase = m.M
	if v559 != 0 {
		v563 = v550
		goto L168
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v546 - v548
	v563 = int32(1)
	goto L168
L171:
	;
	v564 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v564 - v540
	goto L143
L172:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v568
	goto L143
L173:
	;
	goto L174
L174:
	;
	v573 = v568 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v573
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v573
	goto L143
L175:
	;
	if v628 != 0 {
		goto L137
	} else {
		goto L186
	}
L176:
	;
	v628 = v624
	goto L175
L177:
	;
	if v584 <= v585 {
		goto L179
	} else {
		goto L180
	}
L178:
	;
	v624 = int32(0)
	goto L176
L179:
	;
	v628 = int32(-1)
	goto L175
L180:
	;
	goto L181
L181:
	;
	v597 = int32(1)
	v598 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v602 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v598+v584-v597))))
	if int32(246) < v602 {
		v624 = v597
		goto L176
	} else {
		goto L182
	}
L182:
	;
	v604 = v602 - int32(97)
	if v604 < int32(0) {
		v624 = v597
		goto L176
	} else {
		goto L183
	}
L183:
	;
	v610 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v604)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v610)>>(uint(v604&int32(7))%32))&int32(1) == int32(0) {
		v624 = v597
		goto L176
	} else {
		goto L184
	}
L184:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v584 - int32(1)
	goto L185
L185:
	;
	goto L178
L186:
	;
	v637 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v638 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L189
L187:
	;
	if v681 != 0 {
		goto L137
	} else {
		goto L198
	}
L188:
	;
	v681 = v677
	goto L187
L189:
	;
	if v637 <= v638 {
		goto L191
	} else {
		goto L192
	}
L190:
	;
	v677 = int32(0)
	goto L188
L191:
	;
	v681 = int32(-1)
	goto L187
L192:
	;
	goto L193
L193:
	;
	v650 = int32(1)
	v651 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v655 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v651+v637-v650))))
	if int32(122) < v655 {
		v677 = v650
		goto L188
	} else {
		goto L194
	}
L194:
	;
	v657 = v655 - int32(98)
	if v657 < int32(0) {
		v677 = v650
		goto L188
	} else {
		goto L195
	}
L195:
	;
	v663 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v657)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v663)>>(uint(v657&int32(7))%32))&int32(1) == int32(0) {
		v677 = v650
		goto L188
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v637 - int32(1)
	goto L197
L197:
	;
	goto L190
L198:
	;
	goto L143
L199:
	;
	if v683 < int32(0) {
		v1348 = v683
		goto L64
	} else {
		goto L200
	}
L200:
	;
	v687 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v687)+8)) = int32(1)
	goto L137
L201:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v741
	v743 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v744 = *(*int32)(unsafe.Add(mBase, uint32(v743)+8))
	if v744 != 0 {
		goto L218
	} else {
		goto L219
	}
L202:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v693
	v699 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v696
	v703 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_10), int32(14))
	mBase = m.M
	v704 = m.ExcPending
	if v704 != 0 {
		goto L67
	} else {
		goto L203
	}
L203:
	;
	if v703 == int32(0) {
		goto L204
	} else {
		goto L205
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v699
	goto L201
L205:
	;
	goto L206
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v699
	v709 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v709
	if v703 == int32(1) {
		goto L207
	} else {
		goto L208
	}
L207:
	;
	v713 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v714 = int32(2)
	v716 = int32(0)
	v718 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v719 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v718-v719 < v714 {
		v729 = v716
		goto L211
	} else {
		goto L212
	}
L208:
	;
	goto L209
L209:
	;
	v735 = F_slice_del(m, l0)
	mBase = m.M
	v736 = m.ExcPending
	if v736 != 0 {
		goto L67
	} else {
		goto L215
	}
L210:
	;
	if v729 != 0 {
		goto L201
	} else {
		goto L214
	}
L211:
	;
	goto L210
L212:
	;
	v722 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v725 = F_memcmp(m, v722+v718-v714, int32(_a_F_finnish_ISO_8859_1_stem_11), v714)
	mBase = m.M
	if v725 != 0 {
		v729 = v716
		goto L211
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v718 - v714
	v729 = int32(1)
	goto L211
L214:
	;
	v730 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v730 + (v709 - v713)
	goto L209
L215:
	;
	if v735 < int32(0) {
		v1348 = v735
		goto L64
	} else {
		goto L216
	}
L216:
	;
	goto L201
L217:
	;
	v954 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v954
	v958 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v959 = *(*int32)(unsafe.Add(mBase, uint32(v958)+4))
	if v954 < v959 {
		v1342 = int32(0)
		goto L282
	} else {
		goto L283
	}
L218:
	;
	v745 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v746 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v746)+4))
	if v747 <= v745 {
		goto L221
	} else {
		goto L222
	}
L219:
	;
	goto L220
L220:
	;
	v794 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v795 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v796 = *(*int32)(unsafe.Add(mBase, uint32(v795)+4))
	if v794 < v796 {
		v940 = int32(0)
		goto L241
	} else {
		goto L242
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v745
	v750 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v747
	if v747 < v745 {
		goto L226
	} else {
		goto L227
	}
L222:
	;
	v791 = int32(0)
	goto L223
L223:
	;
	if int32(0) <= v791 {
		goto L217
	} else {
		goto L238
	}
L224:
	;
	v791 = v787
	goto L223
L225:
	;
	v768 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_12), int32(2))
	mBase = m.M
	v769 = m.ExcPending
	if v769 != 0 {
		goto L67
	} else {
		goto L230
	}
L226:
	;
	v753 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v757 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v753+v745-int32(1)))))
	if base.Ui32((v757-int32(105))&int32(255)) < base.Ui32(int32(2)) {
		goto L225
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v750
	v787 = int32(0)
	goto L224
L229:
	;
	goto L228
L230:
	;
	if v768 == int32(0) {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v750
	v787 = int32(0)
	goto L224
L232:
	;
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v750
	v775 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v775
	v778 = F_slice_del(m, l0)
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L67
	} else {
		goto L234
	}
L234:
	;
	if int32(0) <= v778 {
		goto L235
	} else {
		goto L236
	}
L235:
	;
	v785 = int32(1)
	goto L237
L236:
	;
	v785 = v778 >> (uint(int32(31)) % 32) & v778
	goto L237
L237:
	;
	v787 = v785
	goto L224
L238:
	;
	v1348 = v791
	goto L64
L239:
	;
	if v947 < int32(0) {
		v1348 = v947
		goto L64
	} else {
		goto L281
	}
L240:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v799
	v947 = int32(0)
	goto L239
L241:
	;
	v947 = v940
	goto L239
L242:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v794
	v799 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v796
	if v796 < v794 {
		goto L244
	} else {
		goto L245
	}
L243:
	;
	v812 = v794 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v812
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v812
	v815 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L250
L244:
	;
	v802 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v806 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v802+v794-int32(1)))))
	if v806 == int32(116) {
		goto L243
	} else {
		goto L247
	}
L245:
	;
	goto L246
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v799
	v947 = int32(0)
	goto L239
L247:
	;
	goto L246
L248:
	;
	if v868 != 0 {
		goto L259
	} else {
		goto L260
	}
L249:
	;
	v868 = v864
	goto L248
L250:
	;
	if v812 <= v825 {
		goto L252
	} else {
		goto L253
	}
L251:
	;
	v864 = int32(0)
	goto L249
L252:
	;
	v868 = int32(-1)
	goto L248
L253:
	;
	goto L254
L254:
	;
	v837 = int32(1)
	v838 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v842 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v838+v812-v837))))
	if int32(246) < v842 {
		v864 = v837
		goto L249
	} else {
		goto L255
	}
L255:
	;
	v844 = v842 - int32(97)
	if v844 < int32(0) {
		v864 = v837
		goto L249
	} else {
		goto L256
	}
L256:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v844)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v850)>>(uint(v844&int32(7))%32))&int32(1) == int32(0) {
		v864 = v837
		goto L249
	} else {
		goto L257
	}
L257:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v812 - int32(1)
	goto L258
L258:
	;
	goto L251
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v799
	v947 = int32(0)
	goto L239
L260:
	;
	goto L261
L261:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v871 + (v812 - v815)
	v875 = F_slice_del(m, l0)
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L67
	} else {
		goto L262
	}
L262:
	;
	if v875 < int32(0) {
		v940 = v875
		goto L241
	} else {
		goto L263
	}
L263:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v799
	v881 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v882 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v883 = *(*int32)(unsafe.Add(mBase, uint32(v882)))
	if v881 < v883 {
		v947 = int32(0)
		goto L239
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v881
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v883
	if v881-int32(2) <= v883 {
		goto L240
	} else {
		goto L265
	}
L265:
	;
	v890 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v894 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v890+v881-int32(1)))))
	if v894 != int32(97) {
		goto L240
	} else {
		goto L266
	}
L266:
	;
	v899 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_13), int32(2))
	mBase = m.M
	v900 = m.ExcPending
	if v900 != 0 {
		goto L67
	} else {
		goto L267
	}
L267:
	;
	if v899 == int32(0) {
		goto L240
	} else {
		goto L268
	}
L268:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v799
	v904 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v904
	if v899 == int32(1) {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v909 = int32(0)
	v910 = int32(2)
	v914 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v915 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v914-v915 < v910 {
		v925 = v909
		goto L273
	} else {
		goto L274
	}
L270:
	;
	goto L271
L271:
	;
	v933 = F_slice_del(m, l0)
	mBase = m.M
	v934 = m.ExcPending
	if v934 != 0 {
		goto L67
	} else {
		goto L277
	}
L272:
	;
	if v925 != 0 {
		v940 = v909
		goto L241
	} else {
		goto L276
	}
L273:
	;
	goto L272
L274:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v921 = F_memcmp(m, v918+v914-v910, int32(_a_F_finnish_ISO_8859_1_stem_14), v910)
	mBase = m.M
	if v921 != 0 {
		v925 = v909
		goto L273
	} else {
		goto L275
	}
L275:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v914 - v910
	v925 = int32(1)
	goto L273
L276:
	;
	v926 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v926 + (v904 - v908)
	goto L271
L277:
	;
	if int32(0) <= v933 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v937 = int32(1)
	goto L280
L279:
	;
	v937 = v933
	goto L280
L280:
	;
	v940 = v937
	goto L241
L281:
	;
	goto L217
L282:
	;
	if v1342 < int32(0) {
		v1348 = v1342
		goto L64
	} else {
		goto L370
	}
L283:
	;
	v961 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v959
	v963 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v964 = v963 - v954
	v967 = F_find_among_b(m, l0, int32(_a_F_finnish_ISO_8859_1_stem_8), int32(7))
	mBase = m.M
	v968 = m.ExcPending
	if v968 != 0 {
		goto L67
	} else {
		goto L284
	}
L284:
	;
	v969 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	if v967 == int32(0) {
		v985 = v969
		goto L286
	} else {
		goto L287
	}
L285:
	;
	v1342 = v1329
	goto L282
L286:
	;
	v987 = v985 - v964
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v987
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v987
	v999 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L294
L287:
	;
	v972 = v969 - v964
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v972
	v974 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v972 <= v974 {
		v985 = v969
		goto L286
	} else {
		goto L288
	}
L288:
	;
	v977 = v972 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v977
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v977
	v980 = F_slice_del(m, l0)
	mBase = m.M
	v981 = m.ExcPending
	if v981 != 0 {
		goto L67
	} else {
		goto L289
	}
L289:
	;
	if v980 < int32(0) {
		v1329 = v980
		goto L285
	} else {
		goto L290
	}
L290:
	;
	v984 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v985 = v984
	goto L286
L291:
	;
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1104 = v1103 - v964
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1104
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1104
	v1107 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1104 <= v1107 {
		v1142 = v1104
		v1144 = v1107
		goto L318
	} else {
		goto L319
	}
L292:
	;
	if v1042 != 0 {
		goto L291
	} else {
		goto L303
	}
L293:
	;
	v1042 = v1038
	goto L292
L294:
	;
	if v987 <= v999 {
		goto L296
	} else {
		goto L297
	}
L295:
	;
	v1038 = int32(0)
	goto L293
L296:
	;
	v1042 = int32(-1)
	goto L292
L297:
	;
	goto L298
L298:
	;
	v1011 = int32(1)
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1016 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1012+v987-v1011))))
	if int32(228) < v1016 {
		v1038 = v1011
		goto L293
	} else {
		goto L299
	}
L299:
	;
	v1018 = v1016 - int32(97)
	if v1018 < int32(0) {
		v1038 = v1011
		goto L293
	} else {
		goto L300
	}
L300:
	;
	v1024 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1018)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[3]))))
	if int32(base.Ui32(v1024)>>(uint(v1018&int32(7))%32))&int32(1) == int32(0) {
		v1038 = v1011
		goto L293
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v987 - int32(1)
	goto L302
L302:
	;
	goto L295
L303:
	;
	v1043 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1043
	v1054 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L306
L304:
	;
	if v1097 != 0 {
		goto L291
	} else {
		goto L315
	}
L305:
	;
	v1097 = v1093
	goto L304
L306:
	;
	if v1043 <= v1054 {
		goto L308
	} else {
		goto L309
	}
L307:
	;
	v1093 = int32(0)
	goto L305
L308:
	;
	v1097 = int32(-1)
	goto L304
L309:
	;
	goto L310
L310:
	;
	v1066 = int32(1)
	v1067 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1071 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1067+v1043-v1066))))
	if int32(122) < v1071 {
		v1093 = v1066
		goto L305
	} else {
		goto L311
	}
L311:
	;
	v1073 = v1071 - int32(98)
	if v1073 < int32(0) {
		v1093 = v1066
		goto L305
	} else {
		goto L312
	}
L312:
	;
	v1079 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1073)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v1079)>>(uint(v1073&int32(7))%32))&int32(1) == int32(0) {
		v1093 = v1066
		goto L305
	} else {
		goto L313
	}
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1043 - int32(1)
	goto L314
L314:
	;
	goto L307
L315:
	;
	v1098 = F_slice_del(m, l0)
	mBase = m.M
	v1099 = m.ExcPending
	if v1099 != 0 {
		goto L67
	} else {
		goto L316
	}
L316:
	;
	if v1098 < int32(0) {
		v1329 = v1098
		goto L285
	} else {
		goto L317
	}
L317:
	;
	goto L291
L318:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1142
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1142
	if v1142 <= v1144 {
		v1176 = v1142
		goto L328
	} else {
		goto L329
	}
L319:
	;
	v1109 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1110 = v1109 + v1104
	v1113 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110-int32(1)))))
	if v1113 != int32(106) {
		v1142 = v1104
		v1144 = v1107
		goto L318
	} else {
		goto L320
	}
L320:
	;
	v1117 = v1104 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1117
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1117
	if v1117 <= v1107 {
		v1142 = v1104
		v1144 = v1107
		goto L318
	} else {
		goto L321
	}
L321:
	;
	v1123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1110-int32(2)))))
	if v1123 != int32(111) {
		goto L322
	} else {
		goto L323
	}
L322:
	;
	v1129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1117+v1109-int32(1)))))
	if v1129 != int32(117) {
		v1142 = v1104
		v1144 = v1107
		goto L318
	} else {
		goto L325
	}
L323:
	;
	goto L324
L324:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1104 - int32(2)
	v1135 = F_slice_del(m, l0)
	mBase = m.M
	v1136 = m.ExcPending
	if v1136 != 0 {
		goto L67
	} else {
		goto L326
	}
L325:
	;
	goto L324
L326:
	;
	if v1135 < int32(0) {
		v1329 = v1135
		goto L285
	} else {
		goto L327
	}
L327:
	;
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1141 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	v1142 = v1139 - v964
	v1144 = v1141
	goto L318
L328:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v961
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1176
	v1181 = int32(0)
	v1197 = v1176
	goto L337
L329:
	;
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1151 = v1150 + v1142
	v1154 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151-int32(1)))))
	if v1154 != int32(111) {
		v1176 = v1142
		goto L328
	} else {
		goto L330
	}
L330:
	;
	v1158 = v1142 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1158
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1158
	if v1158 <= v1144 {
		v1176 = v1142
		goto L328
	} else {
		goto L331
	}
L331:
	;
	v1164 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1151-int32(2)))))
	if v1164 != int32(106) {
		v1176 = v1142
		goto L328
	} else {
		goto L332
	}
L332:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1142 - int32(2)
	v1170 = F_slice_del(m, l0)
	mBase = m.M
	v1171 = m.ExcPending
	if v1171 != 0 {
		goto L67
	} else {
		goto L333
	}
L333:
	;
	if v1170 < int32(0) {
		v1329 = v1170
		goto L285
	} else {
		goto L334
	}
L334:
	;
	v1174 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
	v1176 = v1174 - v964
	goto L328
L335:
	;
	if v1234 < int32(0) {
		v1329 = v1181
		goto L285
	} else {
		goto L346
	}
L336:
	;
	v1234 = v1203
	goto L335
L337:
	;
	if v1197 <= v961 {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v1234 = int32(-1)
	goto L335
L340:
	;
	goto L341
L341:
	;
	v1203 = int32(1)
	v1204 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1204+v1197-v1203))))
	if int32(246) < v1208 {
		goto L336
	} else {
		goto L342
	}
L342:
	;
	v1210 = v1208 - int32(97)
	if v1210 < int32(0) {
		goto L336
	} else {
		goto L343
	}
L343:
	;
	v1216 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1210)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[0]))))
	if int32(base.Ui32(v1216)>>(uint(v1210&int32(7))%32))&int32(1) == int32(0) {
		goto L336
	} else {
		goto L344
	}
L344:
	;
	v1225 = v1197 - int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1225
	v1197 = v1225
	goto L337
L346:
	;
	v1237 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+20)) = v1237
	v1248 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	goto L349
L347:
	;
	if v1291 != 0 {
		v1329 = v1181
		goto L285
	} else {
		goto L358
	}
L348:
	;
	v1291 = v1287
	goto L347
L349:
	;
	if v1237 <= v1248 {
		goto L351
	} else {
		goto L352
	}
L350:
	;
	v1287 = int32(0)
	goto L348
L351:
	;
	v1291 = int32(-1)
	goto L347
L352:
	;
	goto L353
L353:
	;
	v1260 = int32(1)
	v1261 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1261+v1237-v1260))))
	if int32(122) < v1265 {
		v1287 = v1260
		goto L348
	} else {
		goto L354
	}
L354:
	;
	v1267 = v1265 - int32(98)
	if v1267 < int32(0) {
		v1287 = v1260
		goto L348
	} else {
		goto L355
	}
L355:
	;
	v1273 = int32(*(*uint8)(unsafe.Add(mBase, uint32(int32(base.Ui32(v1267)>>(uint(int32(3))%32)))+uint32(_c_F_finnish_ISO_8859_1_stem[2]))))
	if int32(base.Ui32(v1273)>>(uint(v1267&int32(7))%32))&int32(1) == int32(0) {
		v1287 = v1260
		goto L348
	} else {
		goto L356
	}
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1237 - int32(1)
	goto L357
L357:
	;
	goto L350
L358:
	;
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+16)) = v1292
	v1294 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v1295 = *(*int32)(unsafe.Add(mBase, uint32(v1294)))
	v1296 = F_slice_to(m, l0, v1295)
	mBase = m.M
	v1297 = m.ExcPending
	if v1297 != 0 {
		goto L67
	} else {
		goto L359
	}
L359:
	;
	v1298 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v1298))) = v1296
	if v1296 == int32(0) {
		v1342 = int32(-1)
		goto L282
	} else {
		goto L360
	}
L360:
	;
	v1303 = int32(0)
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v1296-int32(4))))
	v1309 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v1310 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	if v1309-v1310 < v1308 {
		v1320 = v1303
		goto L362
	} else {
		goto L363
	}
L361:
	;
	if v1320 == int32(0) {
		v1329 = v1181
		goto L285
	} else {
		goto L365
	}
L362:
	;
	goto L361
L363:
	;
	v1313 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v1316 = F_memcmp(m, v1313+v1309-v1308, v1296, v1308)
	mBase = m.M
	if v1316 != 0 {
		v1320 = v1303
		goto L362
	} else {
		goto L364
	}
L364:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1309 - v1308
	v1320 = int32(1)
	goto L362
L365:
	;
	v1324 = F_slice_del(m, l0)
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L67
	} else {
		goto L366
	}
L366:
	;
	if int32(0) <= v1324 {
		goto L367
	} else {
		goto L368
	}
L367:
	;
	v1328 = int32(1)
	goto L369
L368:
	;
	v1328 = v1324
	goto L369
L369:
	;
	v1329 = v1328
	goto L285
L370:
	;
	v1345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
	*(*int32)(unsafe.Add(mBase, uint32(l0)+4)) = v1345
	v1348 = int32(1)
	goto L64
}
func F_fiprintf(m *base.Module, l0 int32, l1 int32, l2 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	v4 = int32(0)
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = l2
	v12 = F___vfprintf_internal(m, l0, l1, l2, v4, v4)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return
	} else {
		m.G0 = v7 + int32(16)
		return
	}
}
func F_fireRIRrules(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v34 int32
	_ = v34
	var v44 int32
	_ = v44
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
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v114 int32
	_ = v114
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
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v203 int32
	_ = v203
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v261 int32
	_ = v261
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
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
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v338 int32
	_ = v338
	var v339 int32
	_ = v339
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v356 int32
	_ = v356
	var v360 int32
	_ = v360
	var v362 int32
	_ = v362
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
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
	var v382 int32
	_ = v382
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v395 int32
	_ = v395
	var v396 int32
	_ = v396
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v416 int32
	_ = v416
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v442 int32
	_ = v442
	var v443 int32
	_ = v443
	var v444 int32
	_ = v444
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v466 int32
	_ = v466
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v477 int32
	_ = v477
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
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
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v489 int32
	_ = v489
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
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v510 int32
	_ = v510
	var v533 int32
	_ = v533
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v551 int32
	_ = v551
	var v554 int32
	_ = v554
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v565 int32
	_ = v565
	var v572 int32
	_ = v572
	var v573 int32
	_ = v573
	var v575 int32
	_ = v575
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v591 int32
	_ = v591
	var v593 int32
	_ = v593
	var v594 int32
	_ = v594
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v602 int32
	_ = v602
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v607 int32
	_ = v607
	var v608 int32
	_ = v608
	var v609 int32
	_ = v609
	var v612 int32
	_ = v612
	var v614 int32
	_ = v614
	var v615 int32
	_ = v615
	var v616 int32
	_ = v616
	var v618 int32
	_ = v618
	var v619 int32
	_ = v619
	var v620 int32
	_ = v620
	var v621 int32
	_ = v621
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
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
	var v639 int32
	_ = v639
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v644 int32
	_ = v644
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v651 int32
	_ = v651
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
	var v661 int32
	_ = v661
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v673 int32
	_ = v673
	var v674 int32
	_ = v674
	var v677 int32
	_ = v677
	var v685 int32
	_ = v685
	var v688 int32
	_ = v688
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v696 int32
	_ = v696
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v700 int32
	_ = v700
	var v701 int32
	_ = v701
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v712 int32
	_ = v712
	var v718 int32
	_ = v718
	var v722 int32
	_ = v722
	var v741 int32
	_ = v741
	var v744 int32
	_ = v744
	var v745 int32
	_ = v745
	var v747 int32
	_ = v747
	var v748 int32
	_ = v748
	var v749 int32
	_ = v749
	var v750 int32
	_ = v750
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v762 int32
	_ = v762
	var v763 int32
	_ = v763
	var v764 int32
	_ = v764
	var v765 int32
	_ = v765
	var v767 int32
	_ = v767
	var v768 int32
	_ = v768
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v772 int32
	_ = v772
	var v773 int32
	_ = v773
	var v774 int32
	_ = v774
	var v776 int32
	_ = v776
	var v777 int32
	_ = v777
	var v779 int32
	_ = v779
	var v780 int32
	_ = v780
	var v781 int32
	_ = v781
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
	var v790 int32
	_ = v790
	var v795 int32
	_ = v795
	var v798 int32
	_ = v798
	var v799 int32
	_ = v799
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v834 int32
	_ = v834
	var v835 int32
	_ = v835
	var v838 int32
	_ = v838
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v842 int32
	_ = v842
	var v849 int32
	_ = v849
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v865 int32
	_ = v865
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v869 int32
	_ = v869
	var v870 int32
	_ = v870
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v875 int32
	_ = v875
	var v876 int32
	_ = v876
	var v888 int32
	_ = v888
	var v889 int32
	_ = v889
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v909 int32
	_ = v909
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v917 int32
	_ = v917
	var v918 int32
	_ = v918
	var v919 int32
	_ = v919
	var v923 int32
	_ = v923
	var v925 int32
	_ = v925
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v929 int32
	_ = v929
	var v930 int32
	_ = v930
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v937 int32
	_ = v937
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v941 int32
	_ = v941
	var v945 int32
	_ = v945
	var v949 int32
	_ = v949
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v955 int32
	_ = v955
	var v958 int32
	_ = v958
	var v964 int32
	_ = v964
	var v965 int32
	_ = v965
	var v968 int32
	_ = v968
	var v969 int32
	_ = v969
	var v974 int32
	_ = v974
	var v975 int32
	_ = v975
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v981 int32
	_ = v981
	var v982 int32
	_ = v982
	var v988 int32
	_ = v988
	var v989 int32
	_ = v989
	var v996 int32
	_ = v996
	var v997 int32
	_ = v997
	var v999 int32
	_ = v999
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1007 int32
	_ = v1007
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
	var v1015 int32
	_ = v1015
	var v1016 int32
	_ = v1016
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1038 int32
	_ = v1038
	var v1040 int32
	_ = v1040
	var v1041 int32
	_ = v1041
	var v1052 int32
	_ = v1052
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1059 int32
	_ = v1059
	var v1063 int32
	_ = v1063
	var v1065 int32
	_ = v1065
	var v1066 int32
	_ = v1066
	var v1068 int32
	_ = v1068
	var v1069 int32
	_ = v1069
	var v1070 int32
	_ = v1070
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1079 int32
	_ = v1079
	var v1080 int32
	_ = v1080
	var v1081 int32
	_ = v1081
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1084 int32
	_ = v1084
	var v1085 int32
	_ = v1085
	var v1086 int32
	_ = v1086
	var v1087 int32
	_ = v1087
	var v1090 int32
	_ = v1090
	var v1091 int32
	_ = v1091
	var v1092 int32
	_ = v1092
	var v1093 int32
	_ = v1093
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1096 int32
	_ = v1096
	var v1097 int32
	_ = v1097
	var v1098 int32
	_ = v1098
	var v1100 int32
	_ = v1100
	var v1101 int32
	_ = v1101
	var v1102 int32
	_ = v1102
	var v1103 int32
	_ = v1103
	var v1104 int32
	_ = v1104
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1115 int32
	_ = v1115
	var v1116 int32
	_ = v1116
	var v1118 int32
	_ = v1118
	var v1120 int32
	_ = v1120
	var v1121 int32
	_ = v1121
	var v1123 int32
	_ = v1123
	var v1125 int32
	_ = v1125
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1130 int32
	_ = v1130
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 int32
	_ = v1133
	var v1135 int32
	_ = v1135
	var v1138 int32
	_ = v1138
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1141 int32
	_ = v1141
	var v1143 int32
	_ = v1143
	var v1144 int32
	_ = v1144
	var v1145 int32
	_ = v1145
	var v1146 int32
	_ = v1146
	var v1147 int32
	_ = v1147
	var v1149 int32
	_ = v1149
	var v1150 int32
	_ = v1150
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1153 int32
	_ = v1153
	var v1155 int32
	_ = v1155
	var v1158 int32
	_ = v1158
	var v1159 int32
	_ = v1159
	var v1160 int32
	_ = v1160
	var v1161 int32
	_ = v1161
	var v1162 int32
	_ = v1162
	var v1163 int32
	_ = v1163
	var v1164 int32
	_ = v1164
	var v1166 int32
	_ = v1166
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1171 int32
	_ = v1171
	var v1173 int32
	_ = v1173
	var v1174 int32
	_ = v1174
	var v1176 int32
	_ = v1176
	var v1178 int32
	_ = v1178
	var v1179 int32
	_ = v1179
	var v1181 int32
	_ = v1181
	var v1182 int32
	_ = v1182
	var v1184 int32
	_ = v1184
	var v1185 int32
	_ = v1185
	var v1186 int32
	_ = v1186
	var v1187 int32
	_ = v1187
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1194 int32
	_ = v1194
	var v1196 int32
	_ = v1196
	var v1197 int32
	_ = v1197
	var v1198 int32
	_ = v1198
	var v1199 int32
	_ = v1199
	var v1203 int32
	_ = v1203
	var v1205 int32
	_ = v1205
	var v1206 int32
	_ = v1206
	var v1208 int32
	_ = v1208
	var v1209 int32
	_ = v1209
	var v1210 int32
	_ = v1210
	var v1211 int32
	_ = v1211
	var v1214 int32
	_ = v1214
	var v1215 int32
	_ = v1215
	var v1218 int32
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1220 int32
	_ = v1220
	var v1222 int32
	_ = v1222
	var v1223 int32
	_ = v1223
	var v1224 int32
	_ = v1224
	var v1225 int32
	_ = v1225
	var v1229 int32
	_ = v1229
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1234 int32
	_ = v1234
	var v1235 int32
	_ = v1235
	var v1236 int32
	_ = v1236
	var v1237 int32
	_ = v1237
	var v1243 int32
	_ = v1243
	var v1245 int32
	_ = v1245
	var v1246 int32
	_ = v1246
	var v1247 int32
	_ = v1247
	var v1248 int32
	_ = v1248
	var v1252 int32
	_ = v1252
	var v1254 int32
	_ = v1254
	var v1255 int32
	_ = v1255
	var v1257 int32
	_ = v1257
	var v1258 int32
	_ = v1258
	var v1259 int32
	_ = v1259
	var v1260 int32
	_ = v1260
	var v1264 int32
	_ = v1264
	var v1265 int32
	_ = v1265
	var v1266 int32
	_ = v1266
	var v1267 int32
	_ = v1267
	var v1269 int32
	_ = v1269
	var v1271 int32
	_ = v1271
	var v1272 int32
	_ = v1272
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1277 int32
	_ = v1277
	var v1279 int32
	_ = v1279
	var v1280 int32
	_ = v1280
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1286 int32
	_ = v1286
	var v1287 int32
	_ = v1287
	var v1288 int32
	_ = v1288
	var v1289 int32
	_ = v1289
	var v1291 int32
	_ = v1291
	var v1292 int32
	_ = v1292
	var v1293 int32
	_ = v1293
	var v1294 int32
	_ = v1294
	var v1295 int32
	_ = v1295
	var v1297 int32
	_ = v1297
	var v1299 int32
	_ = v1299
	var v1300 int32
	_ = v1300
	var v1302 int32
	_ = v1302
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1307 int32
	_ = v1307
	var v1309 int32
	_ = v1309
	var v1310 int32
	_ = v1310
	var v1342 int32
	_ = v1342
	var v1343 int32
	_ = v1343
	var v1370 int32
	_ = v1370
	var v1374 int32
	_ = v1374
	var v1380 int32
	_ = v1380
	var v1386 int32
	_ = v1386
	var v1398 int32
	_ = v1398
	var v1401 int32
	_ = v1401
	var v1403 int32
	_ = v1403
	var v1404 int32
	_ = v1404
	var v1406 int32
	_ = v1406
	var v1407 int32
	_ = v1407
	var v1408 int32
	_ = v1408
	var v1409 int32
	_ = v1409
	var v1410 int32
	_ = v1410
	var v1412 int32
	_ = v1412
	var v1413 int32
	_ = v1413
	var v1414 int32
	_ = v1414
	var v1416 int32
	_ = v1416
	var v1419 int32
	_ = v1419
	var v1420 int32
	_ = v1420
	var v1422 int32
	_ = v1422
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1428 int32
	_ = v1428
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1438 int32
	_ = v1438
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1446 int32
	_ = v1446
	var v1452 int32
	_ = v1452
	var v1454 int32
	_ = v1454
	var v1455 int32
	_ = v1455
	var v1473 int32
	_ = v1473
	var v1477 int32
	_ = v1477
	var v1478 int32
	_ = v1478
	var v1481 int32
	_ = v1481
	var v1482 int32
	_ = v1482
	var v1483 int32
	_ = v1483
	var v1484 int32
	_ = v1484
	var v1485 int32
	_ = v1485
	var v1487 int32
	_ = v1487
	var v1491 int32
	_ = v1491
	var v1492 int32
	_ = v1492
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1504 int32
	_ = v1504
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	var v1515 int32
	_ = v1515
	var v1516 int32
	_ = v1516
	var v1518 int32
	_ = v1518
	var v1524 int32
	_ = v1524
	var v1530 int32
	_ = v1530
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1542 int32
	_ = v1542
	var v1563 int32
	_ = v1563
	var v1567 int32
	_ = v1567
	var v1568 int32
	_ = v1568
	var v1571 int32
	_ = v1571
	var v1574 int32
	_ = v1574
	var v1576 int32
	_ = v1576
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1588 int32
	_ = v1588
	var v1590 int32
	_ = v1590
	var v1591 int32
	_ = v1591
	var v1592 int32
	_ = v1592
	var v1593 int32
	_ = v1593
	var v1594 int32
	_ = v1594
	var v1595 int32
	_ = v1595
	var v1597 int32
	_ = v1597
	var v1599 int32
	_ = v1599
	var v1601 int32
	_ = v1601
	var v1602 int32
	_ = v1602
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1607 int32
	_ = v1607
	var v1608 int32
	_ = v1608
	var v1610 int32
	_ = v1610
	var v1611 int32
	_ = v1611
	var v1612 int32
	_ = v1612
	var v1613 int32
	_ = v1613
	var v1617 int32
	_ = v1617
	var v1620 int32
	_ = v1620
	var v1621 int32
	_ = v1621
	var v1623 int32
	_ = v1623
	var v1624 int32
	_ = v1624
	var v1625 int32
	_ = v1625
	var v1626 int32
	_ = v1626
	var v1627 int32
	_ = v1627
	var v1634 int32
	_ = v1634
	var v1635 int32
	_ = v1635
	var v1638 int32
	_ = v1638
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1644 int32
	_ = v1644
	var v1652 int32
	_ = v1652
	var v1653 int32
	_ = v1653
	var v1656 int32
	_ = v1656
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1671 int32
	_ = v1671
	var v1672 int32
	_ = v1672
	var v1673 int32
	_ = v1673
	var v1678 int32
	_ = v1678
	var v1679 int32
	_ = v1679
	var v1680 int32
	_ = v1680
	var v1681 int32
	_ = v1681
	var v1683 int32
	_ = v1683
	var v1684 int32
	_ = v1684
	var v1685 int32
	_ = v1685
	var v1686 int32
	_ = v1686
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1690 int32
	_ = v1690
	var v1691 int32
	_ = v1691
	var v1693 int32
	_ = v1693
	var v1697 int32
	_ = v1697
	var v1698 int32
	_ = v1698
	var v1700 int32
	_ = v1700
	var v1701 int32
	_ = v1701
	var v1706 int32
	_ = v1706
	var v1717 int32
	_ = v1717
	var v1723 int32
	_ = v1723
	var v1726 int32
	_ = v1726
	var v1731 int32
	_ = v1731
	var v1732 int32
	_ = v1732
	var v1735 int32
	_ = v1735
	var v1736 int32
	_ = v1736
	var v1737 int32
	_ = v1737
	var v1742 int32
	_ = v1742
	var v1744 int32
	_ = v1744
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1747 int32
	_ = v1747
	var v1750 int32
	_ = v1750
	var v1751 int32
	_ = v1751
	var v1754 int32
	_ = v1754
	var v1756 int32
	_ = v1756
	var v1758 int32
	_ = v1758
	var v1764 int32
	_ = v1764
	var v1765 int32
	_ = v1765
	var v1770 int32
	_ = v1770
	var v1774 int32
	_ = v1774
	var v1775 int32
	_ = v1775
	var v1782 int32
	_ = v1782
	var v1795 int32
	_ = v1795
	var v1821 int32
	_ = v1821
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1828 int32
	_ = v1828
	var v1829 int32
	_ = v1829
	var v1830 int32
	_ = v1830
	var v1831 int32
	_ = v1831
	var v1832 int32
	_ = v1832
	var v1833 int32
	_ = v1833
	var v1834 int32
	_ = v1834
	var v1836 int32
	_ = v1836
	var v1841 int32
	_ = v1841
	var v1863 int32
	_ = v1863
	var v1890 int32
	_ = v1890
	var v1891 int32
	_ = v1891
	var v1893 int32
	_ = v1893
	var v1919 int32
	_ = v1919
	var v1921 int32
	_ = v1921
	var v1945 int32
	_ = v1945
	var v1947 int32
	_ = v1947
	var v1971 int32
	_ = v1971
	var v1974 int32
	_ = v1974
	var v1985 int32
	_ = v1985
	var v2003 int32
	_ = v2003
	var v2007 int32
	_ = v2007
	var v2008 int32
	_ = v2008
	var v2009 int32
	_ = v2009
	var v2010 int32
	_ = v2010
	var v2012 int32
	_ = v2012
	var v2013 int32
	_ = v2013
	var v2014 int32
	_ = v2014
	var v2017 int32
	_ = v2017
	var v2018 int32
	_ = v2018
	var v2045 int32
	_ = v2045
	var v2048 int32
	_ = v2048
	var v2055 int32
	_ = v2055
	var v2056 int32
	_ = v2056
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2061 int32
	_ = v2061
	var v2064 int32
	_ = v2064
	var v2069 int32
	_ = v2069
	var v2075 int32
	_ = v2075
	var v2094 int32
	_ = v2094
	var v2095 int32
	_ = v2095
	var v2099 int32
	_ = v2099
	var v2100 int32
	_ = v2100
	var v2101 int32
	_ = v2101
	var v2104 int32
	_ = v2104
	var v2106 int32
	_ = v2106
	var v2107 int32
	_ = v2107
	var v2108 int32
	_ = v2108
	var v2109 int32
	_ = v2109
	var v2111 int32
	_ = v2111
	var v2114 int32
	_ = v2114
	var v2118 int32
	_ = v2118
	var v2122 int32
	_ = v2122
	var v2126 int32
	_ = v2126
	var v2129 int32
	_ = v2129
	var v2132 int32
	_ = v2132
	var v2133 int32
	_ = v2133
	var v2134 int32
	_ = v2134
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2138 int32
	_ = v2138
	var v2139 int32
	_ = v2139
	var v2140 int32
	_ = v2140
	var v2141 int32
	_ = v2141
	var v2143 int32
	_ = v2143
	var v2144 int32
	_ = v2144
	var v2145 int32
	_ = v2145
	var v2147 int32
	_ = v2147
	var v2148 int32
	_ = v2148
	var v2149 int32
	_ = v2149
	var v2151 int32
	_ = v2151
	var v2155 int32
	_ = v2155
	var v2164 int32
	_ = v2164
	var v2165 int32
	_ = v2165
	var v2166 int32
	_ = v2166
	var v2168 int32
	_ = v2168
	var v2169 int32
	_ = v2169
	var v2176 int32
	_ = v2176
	var v2182 int32
	_ = v2182
	var v2185 int32
	_ = v2185
	var v2186 int32
	_ = v2186
	var v2187 int32
	_ = v2187
	var v2189 int32
	_ = v2189
	var v2190 int32
	_ = v2190
	var v2200 int32
	_ = v2200
	var v2211 int32
	_ = v2211
	var v2212 int32
	_ = v2212
	var v2213 int32
	_ = v2213
	var v2215 int32
	_ = v2215
	var v2224 int32
	_ = v2224
	var v2225 int32
	_ = v2225
	var v2226 int32
	_ = v2226
	var v2229 int32
	_ = v2229
	var v2230 int32
	_ = v2230
	var v2239 int32
	_ = v2239
	var v2240 int32
	_ = v2240
	var v2241 int32
	_ = v2241
	var v2244 int32
	_ = v2244
	var v2247 int32
	_ = v2247
	var v2250 int32
	_ = v2250
	var v2259 int32
	_ = v2259
	var v2261 int32
	_ = v2261
	var v2262 int32
	_ = v2262
	var v2265 int32
	_ = v2265
	var v2267 int32
	_ = v2267
	var v2276 int32
	_ = v2276
	var v2278 int32
	_ = v2278
	var v2279 int32
	_ = v2279
	var v2282 int32
	_ = v2282
	var v2283 int32
	_ = v2283
	var v2284 int32
	_ = v2284
	var v2288 int32
	_ = v2288
	var v2289 int32
	_ = v2289
	var v2297 int32
	_ = v2297
	var v2306 int32
	_ = v2306
	var v2308 int32
	_ = v2308
	var v2309 int32
	_ = v2309
	var v2312 int32
	_ = v2312
	var v2316 int32
	_ = v2316
	var v2317 int32
	_ = v2317
	var v2319 int32
	_ = v2319
	var v2328 int32
	_ = v2328
	var v2330 int32
	_ = v2330
	var v2331 int32
	_ = v2331
	var v2334 int32
	_ = v2334
	var v2335 int32
	_ = v2335
	var v2336 int32
	_ = v2336
	var v2343 int32
	_ = v2343
	var v2345 int32
	_ = v2345
	var v2346 int32
	_ = v2346
	var v2349 int32
	_ = v2349
	var v2356 int32
	_ = v2356
	var v2358 int32
	_ = v2358
	var v2359 int32
	_ = v2359
	var v2362 int32
	_ = v2362
	var v2363 int32
	_ = v2363
	var v2368 int32
	_ = v2368
	var v2371 int32
	_ = v2371
	var v2374 int32
	_ = v2374
	var v2381 int32
	_ = v2381
	var v2382 int32
	_ = v2382
	var v2383 int32
	_ = v2383
	var v2385 int32
	_ = v2385
	var v2386 int32
	_ = v2386
	var v2387 int32
	_ = v2387
	var v2389 int32
	_ = v2389
	var v2395 int32
	_ = v2395
	var v2407 int32
	_ = v2407
	var v2408 int32
	_ = v2408
	var v2410 int32
	_ = v2410
	var v2413 int32
	_ = v2413
	var v2414 int32
	_ = v2414
	var v2420 int32
	_ = v2420
	var v2423 int32
	_ = v2423
	var v2426 int32
	_ = v2426
	var v2427 int32
	_ = v2427
	var v2429 int32
	_ = v2429
	var v2437 int32
	_ = v2437
	var v2438 int32
	_ = v2438
	var v2440 int32
	_ = v2440
	var v2446 int32
	_ = v2446
	var v2452 int32
	_ = v2452
	var v2453 int32
	_ = v2453
	var v2454 int32
	_ = v2454
	var v2455 int32
	_ = v2455
	var v2456 int32
	_ = v2456
	var v2458 int32
	_ = v2458
	var v2459 int32
	_ = v2459
	var v2462 int32
	_ = v2462
	var v2466 int32
	_ = v2466
	var v2470 int32
	_ = v2470
	var v2471 int32
	_ = v2471
	var v2472 int32
	_ = v2472
	var v2473 int32
	_ = v2473
	var v2476 int32
	_ = v2476
	var v2477 int32
	_ = v2477
	var v2480 int32
	_ = v2480
	var v2484 int32
	_ = v2484
	var v2485 int32
	_ = v2485
	var v2486 int32
	_ = v2486
	var v2489 int32
	_ = v2489
	var v2492 int32
	_ = v2492
	var v2493 int32
	_ = v2493
	var v2494 int32
	_ = v2494
	var v2495 int32
	_ = v2495
	var v2497 int32
	_ = v2497
	var v2498 int32
	_ = v2498
	var v2499 int32
	_ = v2499
	var v2500 int32
	_ = v2500
	var v2501 int32
	_ = v2501
	var v2502 int32
	_ = v2502
	var v2504 int32
	_ = v2504
	var v2505 int32
	_ = v2505
	var v2506 int32
	_ = v2506
	var v2507 int32
	_ = v2507
	var v2509 int32
	_ = v2509
	var v2510 int32
	_ = v2510
	var v2511 int32
	_ = v2511
	var v2512 int32
	_ = v2512
	var v2514 int32
	_ = v2514
	var v2516 int32
	_ = v2516
	var v2519 int32
	_ = v2519
	var v2521 int32
	_ = v2521
	var v2524 int32
	_ = v2524
	var v2528 int32
	_ = v2528
	var v2529 int32
	_ = v2529
	var v2543 int32
	_ = v2543
	var v2577 int32
	_ = v2577
	var v2580 int32
	_ = v2580
	var v2581 int32
	_ = v2581
	var v2589 int32
	_ = v2589
	var v2594 int32
	_ = v2594
	var v2598 int32
	_ = v2598
	var v2602 int32
	_ = v2602
	var v2607 int32
	_ = v2607
	var v2611 int32
	_ = v2611
	var v2615 int32
	_ = v2615
	var v2620 int32
	_ = v2620
	var v2624 int32
	_ = v2624
	var v2627 int32
	_ = v2627
	var v2628 int32
	_ = v2628
	var v2636 int32
	_ = v2636
	var v2641 int32
	_ = v2641
	var v2645 int32
	_ = v2645
	var v2646 int32
	_ = v2646
	var v2652 int32
	_ = v2652
	var v2657 int32
	_ = v2657
	var v2661 int32
	_ = v2661
	var v2664 int32
	_ = v2664
	var v2665 int32
	_ = v2665
	var v2671 int32
	_ = v2671
	var v2676 int32
	_ = v2676
	v3 = int32(0)
	v26 = m.G0
	v28 = v26 - int32(80)
	m.G0 = v28
	v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v31 == v3 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v1370 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v1370 == int32(0) {
		v1947 = l1
		goto L311
	} else {
		goto L312
	}
L2:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v34 <= int32(0) {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v44 = v3
	goto L4
L4:
	;
	v62 = *(*int32)(unsafe.Add(mBase, uint32(v31)+12))
	v65 = v62 + v44<<(uint(int32(2))%32)
	v66 = *(*int32)(unsafe.Add(mBase, uint32(v65)))
	v67 = *(*int32)(unsafe.Add(mBase, uint32(v66)+20))
	if v67 == int32(0) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	goto L1
L6:
	;
	v1342 = v44 + int32(1)
	v1343 = *(*int32)(unsafe.Add(mBase, uint32(v31)+4))
	if v1342 < v1343 {
		v44 = v1342
		goto L4
	} else {
		goto L304
	}
L7:
	;
	v70 = *(*int32)(unsafe.Add(mBase, uint32(v66)+24))
	if v70 == int32(0) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	goto L9
L9:
	;
	v73 = int32(0)
	v78 = m.G0
	v80 = v78 - int32(128)
	m.G0 = v80
	v82 = F_copyObjectImpl(m, v66)
	mBase = m.M
	v85 = m.ExcPending
	if v85 != 0 {
		goto L11
	} else {
		goto L12
	}
L10:
	;
	goto L9
L11:
	;
	return int32(0)
L12:
	;
	v86 = *(*int32)(unsafe.Add(mBase, uint32(v82)+16))
	v87 = *(*int32)(unsafe.Add(mBase, uint32(v86)+52))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(v87)+12))
	v89 = *(*int32)(unsafe.Add(mBase, uint32(v86)+144))
	v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
	v91 = *(*int32)(unsafe.Add(mBase, uint32(v90)+4))
	v92 = int32(2)
	v95 = *(*int32)(unsafe.Add(mBase, uint32(v89)+12))
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v95)+4))
	v102 = int32(0)
	v103 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v103 == v102 {
		v118 = v73
		v119 = v102
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v120 = *(*int32)(unsafe.Add(mBase, uint32(v88+v96<<(uint(v92)%32)-int32(4))))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v121 != 0 {
		goto L19
	} else {
		goto L20
	}
L14:
	;
	v108 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v103)+8)))
	if v108 != 0 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v109 = int32(2249)
	goto L17
L16:
	;
	v109 = int32(2287)
	goto L17
L17:
	;
	v111 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v111 == int32(0) {
		v118 = v109
		v119 = int32(1)
		goto L13
	} else {
		goto L18
	}
L18:
	;
	v114 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v111)+4)))
	v118 = v109
	v119 = v114 + int32(1)
	goto L13
L19:
	;
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v122 == int32(0) {
		goto L23
	} else {
		goto L24
	}
L20:
	;
	v144 = v73
	v145 = v73
	goto L21
L21:
	;
	v148 = *(*int32)(unsafe.Add(mBase, uint32(v88+v91<<(uint(v92)%32)-int32(4))))
	v150 = F_palloc0(m, int32(168))
	mBase = m.M
	v151 = m.ExcPending
	if v151 != 0 {
		goto L11
	} else {
		goto L32
	}
L22:
	;
	if v103 != 0 {
		goto L26
	} else {
		goto L27
	}
L23:
	;
	v133 = int32(0)
	v134 = int32(2)
	v135 = int32(1)
	goto L22
L24:
	;
	goto L25
L25:
	;
	v128 = *(*int32)(unsafe.Add(mBase, uint32(v122)+4))
	v133 = v128
	v134 = v128 + int32(2)
	v135 = v128 + int32(1)
	goto L22
L26:
	;
	v136 = v134
	goto L28
L27:
	;
	v136 = v135
	goto L28
L28:
	;
	if v103 != 0 {
		goto L29
	} else {
		goto L30
	}
L29:
	;
	v139 = int32(3)
	goto L31
L30:
	;
	v139 = int32(2)
	goto L31
L31:
	;
	v144 = v133 + v139
	v145 = v136
	goto L21
L32:
	;
	v152 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v150)+24)) = uint8(v152)
	*(*int64)(unsafe.Add(mBase, uint32(v150))) = int64(4294967363)
	v157 = F_palloc0(m, int32(136))
	mBase = m.M
	v158 = m.ExcPending
	if v158 != 0 {
		goto L11
	} else {
		goto L33
	}
L33:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v157))) = int32(101)
	v164 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	v165 = F_makeAlias(m, int32(_a_F_fireRIRrules_0), v164)
	mBase = m.M
	v166 = m.ExcPending
	if v166 != 0 {
		goto L11
	} else {
		goto L34
	}
L34:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v157)+8)) = v165
	*(*int32)(unsafe.Add(mBase, uint32(v157)+4)) = v165
	v169 = *(*int32)(unsafe.Add(mBase, uint32(v120)+36))
	v170 = F_copyObjectImpl(m, v169)
	mBase = m.M
	v171 = m.ExcPending
	if v171 != 0 {
		goto L11
	} else {
		goto L35
	}
L35:
	;
	v172 = int32(1)
	F_IncrementVarSublevelsUp(m, v170, v172, v172)
	mBase = m.M
	v175 = m.ExcPending
	if v175 != 0 {
		goto L11
	} else {
		goto L36
	}
L36:
	;
	v176 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v157)+125)) = uint8(v176)
	*(*int32)(unsafe.Add(mBase, uint32(v157)+36)) = v170
	*(*int32)(unsafe.Add(mBase, uint32(v80)+76)) = v157
	*(*int32)(unsafe.Add(mBase, uint32(v80)+112)) = v157
	v184 = F_list_make1_impl(m, v176, v80+int32(76))
	mBase = m.M
	v185 = m.ExcPending
	if v185 != 0 {
		goto L11
	} else {
		goto L37
	}
L37:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+52)) = v184
	v188 = F_palloc0(m, int32(8))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L11
	} else {
		goto L38
	}
L38:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v188))) = int64(4294967359)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+72)) = v188
	*(*int32)(unsafe.Add(mBase, uint32(v80)+108)) = v188
	v197 = F_list_make1_impl(m, int32(1), v80+int32(72))
	mBase = m.M
	v198 = m.ExcPending
	if v198 != 0 {
		goto L11
	} else {
		goto L39
	}
L39:
	;
	v200 = F_makeFromExpr(m, v197, int32(0))
	mBase = m.M
	v201 = m.ExcPending
	if v201 != 0 {
		goto L11
	} else {
		goto L40
	}
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+60)) = v200
	v203 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v203 == int32(0) {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v307 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v307 != 0 {
		goto L50
	} else {
		goto L51
	}
L42:
	;
	v209 = v203
	v213 = int32(0)
	goto L43
L43:
	;
	v232 = *(*int32)(unsafe.Add(mBase, uint32(v209)+4))
	if v232 <= v213 {
		goto L41
	} else {
		goto L45
	}
L44:
	;
	goto L41
L45:
	;
	v235 = v213 << (uint(int32(2)) % 32)
	v236 = int32(1)
	v238 = v213 + v236
	v239 = base.I32_extend16_s(v238)
	v240 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v241 = *(*int32)(unsafe.Add(mBase, uint32(v240)+12))
	v243 = *(*int32)(unsafe.Add(mBase, uint32(v235+v241)))
	v244 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v244)+12))
	v247 = *(*int32)(unsafe.Add(mBase, uint32(v235+v245)))
	v248 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
	v249 = *(*int32)(unsafe.Add(mBase, uint32(v248)+12))
	v251 = *(*int32)(unsafe.Add(mBase, uint32(v235+v249)))
	v253 = F_makeVar(m, v236, v239, v243, v247, v251, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L11
	} else {
		goto L46
	}
L46:
	;
	v255 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	v256 = *(*int32)(unsafe.Add(mBase, uint32(v255)+12))
	v258 = *(*int32)(unsafe.Add(mBase, uint32(v235+v256)))
	v259 = *(*int32)(unsafe.Add(mBase, uint32(v258)+4))
	v261 = F_makeTargetEntry(m, v253, v239, v259, int32(0))
	mBase = m.M
	v262 = m.ExcPending
	if v262 != 0 {
		goto L11
	} else {
		goto L47
	}
L47:
	;
	v263 = *(*int32)(unsafe.Add(mBase, uint32(v120)+36))
	v264 = *(*int32)(unsafe.Add(mBase, uint32(v263)+76))
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v264)+12))
	v267 = *(*int32)(unsafe.Add(mBase, uint32(v235+v265)))
	v268 = *(*int32)(unsafe.Add(mBase, uint32(v267)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v261)+20)) = v268
	v270 = *(*int32)(unsafe.Add(mBase, uint32(v120)+36))
	v271 = *(*int32)(unsafe.Add(mBase, uint32(v270)+76))
	v272 = *(*int32)(unsafe.Add(mBase, uint32(v271)+12))
	v274 = *(*int32)(unsafe.Add(mBase, uint32(v235+v272)))
	v275 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v274)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v261)+24)) = uint16(v275)
	v277 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	v278 = F_lappend(m, v277, v261)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L11
	} else {
		goto L48
	}
L48:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+76)) = v278
	v281 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v281 != 0 {
		v209 = v281
		v213 = v238
		goto L43
	} else {
		goto L49
	}
L49:
	;
	goto L44
L50:
	;
	v308 = *(*int32)(unsafe.Add(mBase, uint32(v307)+4))
	v309 = F_make_path_rowexpr(m, v82, v308)
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L11
	} else {
		goto L53
	}
L51:
	;
	v372 = v73
	goto L52
L52:
	;
	v373 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v373 != 0 {
		goto L70
	} else {
		goto L71
	}
L53:
	;
	v311 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v312 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v311)+8)))
	if v312 == int32(1) {
		goto L55
	} else {
		goto L56
	}
L54:
	;
	v355 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	if v355 != 0 {
		goto L65
	} else {
		goto L66
	}
L55:
	;
	v320 = F_Int64GetDatum(m, int64(0))
	mBase = m.M
	v321 = m.ExcPending
	if v321 != 0 {
		goto L11
	} else {
		goto L58
	}
L56:
	;
	goto L57
L57:
	;
	v338 = F_palloc0(m, int32(36))
	mBase = m.M
	v339 = m.ExcPending
	if v339 != 0 {
		goto L11
	} else {
		goto L63
	}
L58:
	;
	v322 = int32(0)
	v324 = F_makeConst(m, int32(20), int32(-1), int32(0), int32(8), v320, v322, v322)
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L11
	} else {
		goto L59
	}
L59:
	;
	v326 = *(*int32)(unsafe.Add(mBase, uint32(v309)+4))
	v327 = F_lcons(m, v324, v326)
	mBase = m.M
	v328 = m.ExcPending
	if v328 != 0 {
		goto L11
	} else {
		goto L60
	}
L60:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+4)) = v327
	v331 = F_makeString(m, int32(_a_F_fireRIRrules_1))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L11
	} else {
		goto L61
	}
L61:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v309)+16))
	v334 = F_lcons(m, v331, v333)
	mBase = m.M
	v335 = m.ExcPending
	if v335 != 0 {
		goto L11
	} else {
		goto L62
	}
L62:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v309)+16)) = v334
	v354 = v309
	goto L54
L63:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v338)+12)) = int32(2249)
	*(*int64)(unsafe.Add(mBase, uint32(v338))) = int64(9822590205987)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+68)) = v309
	*(*int32)(unsafe.Add(mBase, uint32(v80)+124)) = v309
	v351 = F_list_make1_impl(m, int32(1), v80+int32(68))
	mBase = m.M
	v352 = m.ExcPending
	if v352 != 0 {
		goto L11
	} else {
		goto L64
	}
L64:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v338)+16)) = v351
	v354 = v338
	goto L54
L65:
	;
	v356 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v355)+4)))
	v360 = v356 + int32(1)
	goto L67
L66:
	;
	v360 = int32(1)
	goto L67
L67:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v363 = *(*int32)(unsafe.Add(mBase, uint32(v362)+12))
	v365 = F_makeTargetEntry(m, v354, base.I32_extend16_s(v360), v363, int32(0))
	mBase = m.M
	v366 = m.ExcPending
	if v366 != 0 {
		goto L11
	} else {
		goto L68
	}
L68:
	;
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	v368 = F_lappend(m, v367, v365)
	mBase = m.M
	v369 = m.ExcPending
	if v369 != 0 {
		goto L11
	} else {
		goto L69
	}
L69:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+76)) = v368
	v372 = v309
	goto L52
L70:
	;
	v374 = *(*int32)(unsafe.Add(mBase, uint32(v373)+16))
	v375 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	if v375 != 0 {
		goto L73
	} else {
		goto L74
	}
L71:
	;
	v429 = v73
	goto L72
L72:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v120)+36)) = v150
	v431 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v431 != 0 {
		goto L86
	} else {
		goto L87
	}
L73:
	;
	v376 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v375)+4)))
	v380 = v376 + int32(1)
	goto L75
L74:
	;
	v380 = int32(1)
	goto L75
L75:
	;
	v382 = *(*int32)(unsafe.Add(mBase, uint32(v373)+8))
	v384 = F_makeTargetEntry(m, v374, base.I32_extend16_s(v380), v382, int32(0))
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L11
	} else {
		goto L76
	}
L76:
	;
	v386 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	v387 = F_lappend(m, v386, v384)
	mBase = m.M
	v388 = m.ExcPending
	if v388 != 0 {
		goto L11
	} else {
		goto L77
	}
L77:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+76)) = v387
	v390 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v391 = *(*int32)(unsafe.Add(mBase, uint32(v390)+4))
	v392 = F_make_path_rowexpr(m, v82, v391)
	mBase = m.M
	v393 = m.ExcPending
	if v393 != 0 {
		goto L11
	} else {
		goto L78
	}
L78:
	;
	v395 = F_palloc0(m, int32(36))
	mBase = m.M
	v396 = m.ExcPending
	if v396 != 0 {
		goto L11
	} else {
		goto L79
	}
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v395)+12)) = int32(2249)
	*(*int64)(unsafe.Add(mBase, uint32(v395))) = int64(9822590205987)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+64)) = v392
	*(*int32)(unsafe.Add(mBase, uint32(v80)+124)) = v392
	v408 = F_list_make1_impl(m, int32(1), v80-int32(-64))
	mBase = m.M
	v409 = m.ExcPending
	if v409 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v395)+16)) = v408
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	if v411 != 0 {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	v412 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v411)+4)))
	v416 = v412 + int32(1)
	goto L83
L82:
	;
	v416 = int32(1)
	goto L83
L83:
	;
	v418 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v419 = *(*int32)(unsafe.Add(mBase, uint32(v418)+20))
	v421 = F_makeTargetEntry(m, v395, base.I32_extend16_s(v416), v419, int32(0))
	mBase = m.M
	v422 = m.ExcPending
	if v422 != 0 {
		goto L11
	} else {
		goto L84
	}
L84:
	;
	v423 = *(*int32)(unsafe.Add(mBase, uint32(v150)+76))
	v424 = F_lappend(m, v423, v421)
	mBase = m.M
	v425 = m.ExcPending
	if v425 != 0 {
		goto L11
	} else {
		goto L85
	}
L85:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v150)+76)) = v424
	v429 = v392
	goto L72
L86:
	;
	v432 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v433 = *(*int32)(unsafe.Add(mBase, uint32(v432)+8))
	v434 = *(*int32)(unsafe.Add(mBase, uint32(v431)+12))
	v435 = F_makeString(m, v434)
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L11
	} else {
		goto L89
	}
L87:
	;
	goto L88
L88:
	;
	v442 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v442 != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	v437 = F_lappend(m, v433, v435)
	mBase = m.M
	v438 = m.ExcPending
	if v438 != 0 {
		goto L11
	} else {
		goto L90
	}
L90:
	;
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v439)+8)) = v437
	goto L88
L91:
	;
	v443 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v444 = *(*int32)(unsafe.Add(mBase, uint32(v443)+8))
	v445 = *(*int32)(unsafe.Add(mBase, uint32(v442)+8))
	v446 = F_makeString(m, v445)
	mBase = m.M
	v447 = m.ExcPending
	if v447 != 0 {
		goto L11
	} else {
		goto L94
	}
L92:
	;
	goto L93
L93:
	;
	v464 = F_palloc0(m, int32(168))
	mBase = m.M
	v465 = m.ExcPending
	if v465 != 0 {
		goto L11
	} else {
		goto L98
	}
L94:
	;
	v448 = F_lappend(m, v444, v446)
	mBase = m.M
	v449 = m.ExcPending
	if v449 != 0 {
		goto L11
	} else {
		goto L95
	}
L95:
	;
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v450)+8)) = v448
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	v453 = *(*int32)(unsafe.Add(mBase, uint32(v452)+8))
	v454 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v455 = *(*int32)(unsafe.Add(mBase, uint32(v454)+20))
	v456 = F_makeString(m, v455)
	mBase = m.M
	v457 = m.ExcPending
	if v457 != 0 {
		goto L11
	} else {
		goto L96
	}
L96:
	;
	v458 = F_lappend(m, v453, v456)
	mBase = m.M
	v459 = m.ExcPending
	if v459 != 0 {
		goto L11
	} else {
		goto L97
	}
L97:
	;
	v460 = *(*int32)(unsafe.Add(mBase, uint32(v120)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v460)+8)) = v458
	goto L93
L98:
	;
	v466 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v464)+24)) = uint8(v466)
	*(*int64)(unsafe.Add(mBase, uint32(v464))) = int64(4294967363)
	v471 = F_palloc0(m, int32(136))
	mBase = m.M
	v472 = m.ExcPending
	if v472 != 0 {
		goto L11
	} else {
		goto L99
	}
L99:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471)+12)) = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v471))) = int32(101)
	v477 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	v478 = F_copyObjectImpl(m, v477)
	mBase = m.M
	v479 = m.ExcPending
	if v479 != 0 {
		goto L11
	} else {
		goto L100
	}
L100:
	;
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v480 != 0 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v481 = *(*int32)(unsafe.Add(mBase, uint32(v480)+12))
	v482 = F_makeString(m, v481)
	mBase = m.M
	v483 = m.ExcPending
	if v483 != 0 {
		goto L11
	} else {
		goto L104
	}
L102:
	;
	v486 = v478
	goto L103
L103:
	;
	v487 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v487 != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	v484 = F_lappend(m, v478, v482)
	mBase = m.M
	v485 = m.ExcPending
	if v485 != 0 {
		goto L11
	} else {
		goto L105
	}
L105:
	;
	v486 = v484
	goto L103
L106:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(v487)+8))
	v489 = F_makeString(m, v488)
	mBase = m.M
	v490 = m.ExcPending
	if v490 != 0 {
		goto L11
	} else {
		goto L109
	}
L107:
	;
	v499 = v486
	goto L108
L108:
	;
	v501 = F_makeAlias(m, int32(_a_F_fireRIRrules_2), v499)
	mBase = m.M
	v502 = m.ExcPending
	if v502 != 0 {
		goto L11
	} else {
		goto L113
	}
L109:
	;
	v491 = F_lappend(m, v486, v489)
	mBase = m.M
	v492 = m.ExcPending
	if v492 != 0 {
		goto L11
	} else {
		goto L110
	}
L110:
	;
	v493 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v494 = *(*int32)(unsafe.Add(mBase, uint32(v493)+20))
	v495 = F_makeString(m, v494)
	mBase = m.M
	v496 = m.ExcPending
	if v496 != 0 {
		goto L11
	} else {
		goto L111
	}
L111:
	;
	v497 = F_lappend(m, v491, v495)
	mBase = m.M
	v498 = m.ExcPending
	if v498 != 0 {
		goto L11
	} else {
		goto L112
	}
L112:
	;
	v499 = v497
	goto L108
L113:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v471)+8)) = v501
	*(*int32)(unsafe.Add(mBase, uint32(v471)+4)) = v501
	v505 = *(*int32)(unsafe.Add(mBase, uint32(v148)+36))
	v506 = *(*int32)(unsafe.Add(mBase, uint32(v505)+52))
	v510 = int32(1)
	goto L116
L114:
	;
	v834 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v834 != 0 {
		goto L183
	} else {
		goto L184
	}
L115:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v795 = m.ExcPending
	if v795 != 0 {
		goto L11
	} else {
		goto L179
	}
L116:
	;
	if v506 != 0 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	if v510 <= int32(0) {
		goto L115
	} else {
		goto L134
	}
L118:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v506)+4))
	v535 = v533
	goto L120
L119:
	;
	v535 = int32(0)
	goto L120
L120:
	;
	if v535 < v510 {
		goto L115
	} else {
		goto L121
	}
L121:
	;
	v537 = *(*int32)(unsafe.Add(mBase, uint32(v506)+12))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v537+v510<<(uint(int32(2))%32)-int32(4))))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v543)+12))
	if v544 != int32(6) {
		goto L123
	} else {
		goto L124
	}
L122:
	;
	goto L117
L123:
	;
	v510 = v510 + int32(1)
	goto L116
L124:
	;
	v547 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	v548 = *(*int32)(unsafe.Add(mBase, uint32(v543)+84))
	v551 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v547))))
	v554 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v548))))
	if base.B2i32(v551 == int32(0))|base.B2i32(v551 != v554) != 0 {
		v572 = v551
		v573 = v554
		goto L126
	} else {
		goto L127
	}
L125:
	;
	if v572-v573 != 0 {
		goto L123
	} else {
		goto L132
	}
L126:
	;
	goto L125
L127:
	;
	v557 = v547
	v558 = v548
	goto L128
L128:
	;
	v561 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v558)+1)))
	v562 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v557)+1)))
	if v562 == int32(0) {
		v572 = v562
		v573 = v561
		goto L126
	} else {
		goto L130
	}
L129:
	;
	v572 = v562
	v573 = v561
	goto L126
L130:
	;
	v565 = int32(1)
	if v562 == v561 {
		v557 = v557 + v565
		v558 = v558 + v565
		goto L128
	} else {
		goto L131
	}
L131:
	;
	goto L129
L132:
	;
	v575 = *(*int32)(unsafe.Add(mBase, uint32(v543)+88))
	if v575 == int32(2) {
		goto L122
	} else {
		goto L133
	}
L133:
	;
	goto L123
L134:
	;
	v582 = F_copyObjectImpl(m, v505)
	mBase = m.M
	v583 = m.ExcPending
	if v583 != 0 {
		goto L11
	} else {
		goto L135
	}
L135:
	;
	v584 = int32(1)
	F_IncrementVarSublevelsUp(m, v582, v584, v584)
	mBase = m.M
	v587 = m.ExcPending
	if v587 != 0 {
		goto L11
	} else {
		goto L136
	}
L136:
	;
	v588 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v588 != 0 {
		goto L137
	} else {
		goto L138
	}
L137:
	;
	v591 = int32(0)
	v593 = F_makeVar(m, v510, base.I32_extend16_s(v119), v118, int32(-1), v591, v591)
	mBase = m.M
	v594 = m.ExcPending
	if v594 != 0 {
		goto L11
	} else {
		goto L140
	}
L138:
	;
	goto L139
L139:
	;
	v612 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v612 != 0 {
		goto L146
	} else {
		goto L147
	}
L140:
	;
	v595 = *(*int32)(unsafe.Add(mBase, uint32(v582)+76))
	if v595 != 0 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v596 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v595)+4)))
	v600 = v596 + int32(1)
	goto L143
L142:
	;
	v600 = int32(1)
	goto L143
L143:
	;
	v602 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v603 = *(*int32)(unsafe.Add(mBase, uint32(v602)+12))
	v605 = F_makeTargetEntry(m, v593, base.I32_extend16_s(v600), v603, int32(0))
	mBase = m.M
	v606 = m.ExcPending
	if v606 != 0 {
		goto L11
	} else {
		goto L144
	}
L144:
	;
	v607 = *(*int32)(unsafe.Add(mBase, uint32(v582)+76))
	v608 = F_lappend(m, v607, v605)
	mBase = m.M
	v609 = m.ExcPending
	if v609 != 0 {
		goto L11
	} else {
		goto L145
	}
L145:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582)+76)) = v608
	goto L139
L146:
	;
	v614 = *(*int32)(unsafe.Add(mBase, uint32(v612)+28))
	v615 = *(*int32)(unsafe.Add(mBase, uint32(v612)+32))
	v616 = *(*int32)(unsafe.Add(mBase, uint32(v612)+36))
	v618 = F_makeVar(m, v510, base.I32_extend16_s(v145), v614, v615, v616, int32(0))
	mBase = m.M
	v619 = m.ExcPending
	if v619 != 0 {
		goto L11
	} else {
		goto L149
	}
L147:
	;
	goto L148
L148:
	;
	v661 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v471)+125)) = uint8(v661)
	*(*int32)(unsafe.Add(mBase, uint32(v471)+36)) = v582
	*(*int32)(unsafe.Add(mBase, uint32(v80)+60)) = v471
	*(*int32)(unsafe.Add(mBase, uint32(v80)+104)) = v471
	v669 = F_list_make1_impl(m, v661, v80+int32(60))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L11
	} else {
		goto L161
	}
L149:
	;
	v620 = *(*int32)(unsafe.Add(mBase, uint32(v582)+76))
	if v620 != 0 {
		goto L150
	} else {
		goto L151
	}
L150:
	;
	v621 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v620)+4)))
	v625 = v621 + int32(1)
	goto L152
L151:
	;
	v625 = int32(1)
	goto L152
L152:
	;
	v627 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v628 = *(*int32)(unsafe.Add(mBase, uint32(v627)+8))
	v630 = F_makeTargetEntry(m, v618, base.I32_extend16_s(v625), v628, int32(0))
	mBase = m.M
	v631 = m.ExcPending
	if v631 != 0 {
		goto L11
	} else {
		goto L153
	}
L153:
	;
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v582)+76))
	v633 = F_lappend(m, v632, v630)
	mBase = m.M
	v634 = m.ExcPending
	if v634 != 0 {
		goto L11
	} else {
		goto L154
	}
L154:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582)+76)) = v633
	v639 = int32(0)
	v641 = F_makeVar(m, v510, base.I32_extend16_s(v144), int32(2287), int32(-1), v639, v639)
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L11
	} else {
		goto L155
	}
L155:
	;
	v643 = *(*int32)(unsafe.Add(mBase, uint32(v582)+76))
	if v643 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v644 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v643)+4)))
	v648 = v644 + int32(1)
	goto L158
L157:
	;
	v648 = int32(1)
	goto L158
L158:
	;
	v650 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v651 = *(*int32)(unsafe.Add(mBase, uint32(v650)+20))
	v653 = F_makeTargetEntry(m, v641, base.I32_extend16_s(v648), v651, int32(0))
	mBase = m.M
	v654 = m.ExcPending
	if v654 != 0 {
		goto L11
	} else {
		goto L159
	}
L159:
	;
	v655 = *(*int32)(unsafe.Add(mBase, uint32(v582)+76))
	v656 = F_lappend(m, v655, v653)
	mBase = m.M
	v657 = m.ExcPending
	if v657 != 0 {
		goto L11
	} else {
		goto L160
	}
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v582)+76)) = v656
	goto L148
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v464)+52)) = v669
	v673 = F_palloc0(m, int32(8))
	mBase = m.M
	v674 = m.ExcPending
	if v674 != 0 {
		goto L11
	} else {
		goto L162
	}
L162:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v673))) = int64(4294967359)
	v677 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v677 == int32(0) {
		goto L164
	} else {
		goto L165
	}
L163:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v700))) = v673
	*(*int32)(unsafe.Add(mBase, uint32(v80)+56)) = v673
	v707 = F_list_make1_impl(m, int32(1), v80+int32(56))
	mBase = m.M
	v708 = m.ExcPending
	if v708 != 0 {
		goto L11
	} else {
		goto L169
	}
L164:
	;
	v700 = v80 + int32(96)
	v701 = int32(0)
	goto L163
L165:
	;
	goto L166
L166:
	;
	v685 = *(*int32)(unsafe.Add(mBase, uint32(v677)+40))
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v677)+28))
	v689 = *(*int32)(unsafe.Add(mBase, uint32(v677)+32))
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v677)+36))
	v692 = F_makeVar(m, int32(1), base.I32_extend16_s(v145), v688, v689, v690, int32(0))
	mBase = m.M
	v693 = m.ExcPending
	if v693 != 0 {
		goto L11
	} else {
		goto L167
	}
L167:
	;
	v694 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v695 = *(*int32)(unsafe.Add(mBase, uint32(v694)+12))
	v696 = *(*int32)(unsafe.Add(mBase, uint32(v694)+36))
	v697 = F_make_opclause(m, v685, v692, v695, v696)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L11
	} else {
		goto L168
	}
L168:
	;
	v700 = v80 + int32(100)
	v701 = v697
	goto L163
L169:
	;
	v709 = F_makeFromExpr(m, v707, v701)
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L11
	} else {
		goto L170
	}
L170:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v464)+60)) = v709
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v712 == int32(0) {
		goto L114
	} else {
		goto L171
	}
L171:
	;
	v718 = v712
	v722 = int32(0)
	goto L172
L172:
	;
	v741 = *(*int32)(unsafe.Add(mBase, uint32(v718)+4))
	if v741 <= v722 {
		goto L114
	} else {
		goto L174
	}
L173:
	;
	goto L114
L174:
	;
	v744 = v722 << (uint(int32(2)) % 32)
	v745 = int32(1)
	v747 = v722 + v745
	v748 = base.I32_extend16_s(v747)
	v749 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v750 = *(*int32)(unsafe.Add(mBase, uint32(v749)+12))
	v752 = *(*int32)(unsafe.Add(mBase, uint32(v744+v750)))
	v753 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	v754 = *(*int32)(unsafe.Add(mBase, uint32(v753)+12))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v744+v754)))
	v757 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
	v758 = *(*int32)(unsafe.Add(mBase, uint32(v757)+12))
	v760 = *(*int32)(unsafe.Add(mBase, uint32(v744+v758)))
	v762 = F_makeVar(m, v745, v748, v752, v756, v760, int32(0))
	mBase = m.M
	v763 = m.ExcPending
	if v763 != 0 {
		goto L11
	} else {
		goto L175
	}
L175:
	;
	v764 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	v765 = *(*int32)(unsafe.Add(mBase, uint32(v764)+12))
	v767 = *(*int32)(unsafe.Add(mBase, uint32(v744+v765)))
	v768 = *(*int32)(unsafe.Add(mBase, uint32(v767)+4))
	v770 = F_makeTargetEntry(m, v762, v748, v768, int32(0))
	mBase = m.M
	v771 = m.ExcPending
	if v771 != 0 {
		goto L11
	} else {
		goto L176
	}
L176:
	;
	v772 = *(*int32)(unsafe.Add(mBase, uint32(v148)+36))
	v773 = *(*int32)(unsafe.Add(mBase, uint32(v772)+76))
	v774 = *(*int32)(unsafe.Add(mBase, uint32(v773)+12))
	v776 = *(*int32)(unsafe.Add(mBase, uint32(v744+v774)))
	v777 = *(*int32)(unsafe.Add(mBase, uint32(v776)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v770)+20)) = v777
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v148)+36))
	v780 = *(*int32)(unsafe.Add(mBase, uint32(v779)+76))
	v781 = *(*int32)(unsafe.Add(mBase, uint32(v780)+12))
	v783 = *(*int32)(unsafe.Add(mBase, uint32(v744+v781)))
	v784 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v783)+24)))
	*(*uint16)(unsafe.Add(mBase, uint32(v770)+24)) = uint16(v784)
	v786 = *(*int32)(unsafe.Add(mBase, uint32(v464)+76))
	v787 = F_lappend(m, v786, v770)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L11
	} else {
		goto L177
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v464)+76)) = v787
	v790 = *(*int32)(unsafe.Add(mBase, uint32(v82)+40))
	if v790 != 0 {
		v718 = v790
		v722 = v747
		goto L172
	} else {
		goto L178
	}
L178:
	;
	goto L173
L179:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v798 = m.ExcPending
	if v798 != 0 {
		goto L11
	} else {
		goto L180
	}
L180:
	;
	v799 = *(*int32)(unsafe.Add(mBase, uint32(v82)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v80))) = v799
	F_errmsg(m, int32(_a_F_fireRIRrules_3), v80)
	mBase = m.M
	v803 = m.ExcPending
	if v803 != 0 {
		goto L11
	} else {
		goto L181
	}
L181:
	;
	F_errfinish(m, int32(_a_F_fireRIRrules_4), int32(411), int32(_a_F_fireRIRrules_5))
	mBase = m.M
	v808 = m.ExcPending
	if v808 != 0 {
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
	v835 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v834)+8)))
	if v835 == int32(1) {
		goto L187
	} else {
		goto L188
	}
L184:
	;
	goto L185
L185:
	;
	v937 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v937 != 0 {
		goto L205
	} else {
		goto L206
	}
L186:
	;
	v918 = *(*int32)(unsafe.Add(mBase, uint32(v464)+76))
	if v918 != 0 {
		goto L200
	} else {
		goto L201
	}
L187:
	;
	v838 = F_copyObjectImpl(m, v372)
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L11
	} else {
		goto L190
	}
L188:
	;
	goto L189
L189:
	;
	v875 = F_palloc0(m, int32(36))
	mBase = m.M
	v876 = m.ExcPending
	if v876 != 0 {
		goto L11
	} else {
		goto L195
	}
L190:
	;
	v841 = F_palloc0(m, int32(24))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L11
	} else {
		goto L191
	}
L191:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v841))) = int32(25)
	v849 = int32(0)
	v851 = F_makeVar(m, int32(1), base.I32_extend16_s(v119), int32(2249), int32(-1), v849, v849)
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L11
	} else {
		goto L192
	}
L192:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v841)+12)) = int64(-4294967276)
	v855 = int32(1)
	*(*uint16)(unsafe.Add(mBase, uint32(v841)+8)) = uint16(v855)
	*(*int32)(unsafe.Add(mBase, uint32(v841)+4)) = v851
	*(*int32)(unsafe.Add(mBase, uint32(v80)+40)) = v841
	*(*int32)(unsafe.Add(mBase, uint32(v80)+92)) = v841
	v865 = F_list_make1_impl(m, v855, v80+int32(40))
	mBase = m.M
	v866 = m.ExcPending
	if v866 != 0 {
		goto L11
	} else {
		goto L193
	}
L193:
	;
	v867 = int32(0)
	v869 = F_makeFuncExpr(m, int32(1219), int32(20), v865, v867, v867)
	mBase = m.M
	v870 = m.ExcPending
	if v870 != 0 {
		goto L11
	} else {
		goto L194
	}
L194:
	;
	v871 = *(*int32)(unsafe.Add(mBase, uint32(v838)+4))
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v871)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v872))) = v869
	v917 = v838
	goto L186
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v875)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v875)+12)) = int32(2249)
	*(*int64)(unsafe.Add(mBase, uint32(v875))) = int64(9822590205987)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+52)) = v372
	*(*int32)(unsafe.Add(mBase, uint32(v80)+124)) = v372
	v888 = F_list_make1_impl(m, int32(1), v80+int32(52))
	mBase = m.M
	v889 = m.ExcPending
	if v889 != 0 {
		goto L11
	} else {
		goto L196
	}
L196:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v875)+16)) = v888
	v895 = int32(0)
	v897 = F_makeVar(m, int32(1), base.I32_extend16_s(v119), int32(2287), int32(-1), v895, v895)
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L11
	} else {
		goto L197
	}
L197:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+116)) = v875
	*(*int32)(unsafe.Add(mBase, uint32(v80)+120)) = v897
	*(*int32)(unsafe.Add(mBase, uint32(v80)+48)) = v897
	*(*int32)(unsafe.Add(mBase, uint32(v80)+44)) = v875
	v909 = F_list_make2_impl(m, v80+int32(48), v80+int32(44))
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L11
	} else {
		goto L198
	}
L198:
	;
	v911 = int32(0)
	v913 = F_makeFuncExpr(m, int32(383), int32(2287), v909, v911, v911)
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L11
	} else {
		goto L199
	}
L199:
	;
	v917 = v913
	goto L186
L200:
	;
	v919 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v918)+4)))
	v923 = v919 + int32(1)
	goto L202
L201:
	;
	v923 = int32(1)
	goto L202
L202:
	;
	v925 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v926 = *(*int32)(unsafe.Add(mBase, uint32(v925)+12))
	v928 = F_makeTargetEntry(m, v917, base.I32_extend16_s(v923), v926, int32(0))
	mBase = m.M
	v929 = m.ExcPending
	if v929 != 0 {
		goto L11
	} else {
		goto L203
	}
L203:
	;
	v930 = *(*int32)(unsafe.Add(mBase, uint32(v464)+76))
	v931 = F_lappend(m, v930, v928)
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L11
	} else {
		goto L204
	}
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v464)+76)) = v931
	goto L185
L205:
	;
	v939 = F_palloc0(m, int32(36))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L11
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v148)+36)) = v464
	v1079 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v1079 != 0 {
		goto L229
	} else {
		goto L230
	}
L208:
	;
	v941 = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v939)+32)) = v941
	*(*int64)(unsafe.Add(mBase, uint32(v939))) = int64(12833362280468)
	v945 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v939)+20)) = uint8(v945)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+88)) = v429
	v949 = base.I32_extend16_s(v144)
	v952 = int32(0)
	v954 = F_makeVar(m, v945, v949, int32(2287), v941, v952, v952)
	mBase = m.M
	v955 = m.ExcPending
	if v955 != 0 {
		goto L11
	} else {
		goto L209
	}
L209:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+84)) = v954
	*(*int32)(unsafe.Add(mBase, uint32(v80)+32)) = v954
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v80)+88))
	*(*int32)(unsafe.Add(mBase, uint32(v80)+36)) = v958
	v964 = F_list_make2_impl(m, v80+int32(36), v80+int32(32))
	mBase = m.M
	v965 = m.ExcPending
	if v965 != 0 {
		goto L11
	} else {
		goto L210
	}
L210:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v939)+28)) = v964
	v968 = F_palloc0(m, int32(28))
	mBase = m.M
	v969 = m.ExcPending
	if v969 != 0 {
		goto L11
	} else {
		goto L211
	}
L211:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v968)+24)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v968))) = int32(32)
	v974 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v975 = *(*int32)(unsafe.Add(mBase, uint32(v974)+28))
	*(*int32)(unsafe.Add(mBase, uint32(v968)+4)) = v975
	v977 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v978 = *(*int32)(unsafe.Add(mBase, uint32(v977)+36))
	*(*int32)(unsafe.Add(mBase, uint32(v968)+8)) = v978
	v981 = F_palloc0(m, int32(16))
	mBase = m.M
	v982 = m.ExcPending
	if v982 != 0 {
		goto L11
	} else {
		goto L212
	}
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v981)+12)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v981))) = int32(33)
	*(*int32)(unsafe.Add(mBase, uint32(v981)+4)) = v939
	v988 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v989 = *(*int32)(unsafe.Add(mBase, uint32(v988)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v981)+8)) = v989
	*(*int32)(unsafe.Add(mBase, uint32(v80)+28)) = v981
	*(*int32)(unsafe.Add(mBase, uint32(v80)+80)) = v981
	v996 = F_list_make1_impl(m, int32(1), v80+int32(28))
	mBase = m.M
	v997 = m.ExcPending
	if v997 != 0 {
		goto L11
	} else {
		goto L213
	}
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v968)+16)) = v996
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1000 = *(*int32)(unsafe.Add(mBase, uint32(v999)+16))
	*(*int32)(unsafe.Add(mBase, uint32(v968)+20)) = v1000
	v1002 = *(*int32)(unsafe.Add(mBase, uint32(v464)+76))
	if v1002 != 0 {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	v1003 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1002)+4)))
	v1007 = v1003 + int32(1)
	goto L216
L215:
	;
	v1007 = int32(1)
	goto L216
L216:
	;
	v1009 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1010 = *(*int32)(unsafe.Add(mBase, uint32(v1009)+8))
	v1012 = F_makeTargetEntry(m, v968, base.I32_extend16_s(v1007), v1010, int32(0))
	mBase = m.M
	v1013 = m.ExcPending
	if v1013 != 0 {
		goto L11
	} else {
		goto L217
	}
L217:
	;
	v1014 = *(*int32)(unsafe.Add(mBase, uint32(v464)+76))
	v1015 = F_lappend(m, v1014, v1012)
	mBase = m.M
	v1016 = m.ExcPending
	if v1016 != 0 {
		goto L11
	} else {
		goto L218
	}
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v464)+76)) = v1015
	v1019 = F_palloc0(m, int32(36))
	mBase = m.M
	v1020 = m.ExcPending
	if v1020 != 0 {
		goto L11
	} else {
		goto L219
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+32)) = int32(-1)
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+12)) = int32(2249)
	*(*int64)(unsafe.Add(mBase, uint32(v1019))) = int64(9822590205987)
	*(*int32)(unsafe.Add(mBase, uint32(v80)+24)) = v429
	*(*int32)(unsafe.Add(mBase, uint32(v80)+124)) = v429
	v1032 = F_list_make1_impl(m, int32(1), v80+int32(24))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L11
	} else {
		goto L220
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1019)+16)) = v1032
	v1038 = int32(0)
	v1040 = F_makeVar(m, int32(1), v949, int32(2287), int32(-1), v1038, v1038)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L11
	} else {
		goto L221
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v80)+116)) = v1019
	*(*int32)(unsafe.Add(mBase, uint32(v80)+120)) = v1040
	*(*int32)(unsafe.Add(mBase, uint32(v80)+20)) = v1040
	*(*int32)(unsafe.Add(mBase, uint32(v80)+16)) = v1019
	v1052 = F_list_make2_impl(m, v80+int32(20), v80+int32(16))
	mBase = m.M
	v1053 = m.ExcPending
	if v1053 != 0 {
		goto L11
	} else {
		goto L222
	}
L222:
	;
	v1054 = int32(0)
	v1056 = F_makeFuncExpr(m, int32(383), int32(2287), v1052, v1054, v1054)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L11
	} else {
		goto L223
	}
L223:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v464)+76))
	if v1058 != 0 {
		goto L224
	} else {
		goto L225
	}
L224:
	;
	v1059 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1058)+4)))
	v1063 = v1059 + int32(1)
	goto L226
L225:
	;
	v1063 = int32(1)
	goto L226
L226:
	;
	v1065 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1066 = *(*int32)(unsafe.Add(mBase, uint32(v1065)+20))
	v1068 = F_makeTargetEntry(m, v1056, base.I32_extend16_s(v1063), v1066, int32(0))
	mBase = m.M
	v1069 = m.ExcPending
	if v1069 != 0 {
		goto L11
	} else {
		goto L227
	}
L227:
	;
	v1070 = *(*int32)(unsafe.Add(mBase, uint32(v464)+76))
	v1071 = F_lappend(m, v1070, v1068)
	mBase = m.M
	v1072 = m.ExcPending
	if v1072 != 0 {
		goto L11
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v464)+76)) = v1071
	goto L207
L229:
	;
	v1080 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	v1081 = *(*int32)(unsafe.Add(mBase, uint32(v1080)+8))
	v1082 = *(*int32)(unsafe.Add(mBase, uint32(v1079)+12))
	v1083 = F_makeString(m, v1082)
	mBase = m.M
	v1084 = m.ExcPending
	if v1084 != 0 {
		goto L11
	} else {
		goto L232
	}
L230:
	;
	goto L231
L231:
	;
	v1090 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v1090 != 0 {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	v1085 = F_lappend(m, v1081, v1083)
	mBase = m.M
	v1086 = m.ExcPending
	if v1086 != 0 {
		goto L11
	} else {
		goto L233
	}
L233:
	;
	v1087 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1087)+8)) = v1085
	goto L231
L234:
	;
	v1091 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	v1092 = *(*int32)(unsafe.Add(mBase, uint32(v1091)+8))
	v1093 = *(*int32)(unsafe.Add(mBase, uint32(v1090)+8))
	v1094 = F_makeString(m, v1093)
	mBase = m.M
	v1095 = m.ExcPending
	if v1095 != 0 {
		goto L11
	} else {
		goto L237
	}
L235:
	;
	goto L236
L236:
	;
	v1111 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v1111 == int32(0) {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	v1096 = F_lappend(m, v1092, v1094)
	mBase = m.M
	v1097 = m.ExcPending
	if v1097 != 0 {
		goto L11
	} else {
		goto L238
	}
L238:
	;
	v1098 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1098)+8)) = v1096
	v1100 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	v1101 = *(*int32)(unsafe.Add(mBase, uint32(v1100)+8))
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1103 = *(*int32)(unsafe.Add(mBase, uint32(v1102)+20))
	v1104 = F_makeString(m, v1103)
	mBase = m.M
	v1105 = m.ExcPending
	if v1105 != 0 {
		goto L11
	} else {
		goto L239
	}
L239:
	;
	v1106 = F_lappend(m, v1101, v1104)
	mBase = m.M
	v1107 = m.ExcPending
	if v1107 != 0 {
		goto L11
	} else {
		goto L240
	}
L240:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v148)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1108)+8)) = v1106
	goto L236
L241:
	;
	v1135 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v1135 == int32(0) {
		goto L249
	} else {
		goto L250
	}
L242:
	;
	v1114 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v1115 = F_lappend_oid(m, v1114, v118)
	mBase = m.M
	v1116 = m.ExcPending
	if v1116 != 0 {
		goto L11
	} else {
		goto L243
	}
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v1115
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1120 = F_lappend_int(m, v1118, int32(-1))
	mBase = m.M
	v1121 = m.ExcPending
	if v1121 != 0 {
		goto L11
	} else {
		goto L244
	}
L244:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v1120
	v1123 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v1125 = F_lappend_oid(m, v1123, int32(0))
	mBase = m.M
	v1126 = m.ExcPending
	if v1126 != 0 {
		goto L11
	} else {
		goto L245
	}
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v1125
	v1128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+8)))
	if v1128 != 0 {
		goto L241
	} else {
		goto L246
	}
L246:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	v1130 = F_makeSortGroupClauseForSetOp(m, v118)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L11
	} else {
		goto L247
	}
L247:
	;
	v1132 = F_lappend(m, v1129, v1130)
	mBase = m.M
	v1133 = m.ExcPending
	if v1133 != 0 {
		goto L11
	} else {
		goto L248
	}
L248:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v1132
	goto L241
L249:
	;
	v1189 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v1189 != 0 {
		goto L265
	} else {
		goto L266
	}
L250:
	;
	v1138 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v1135)+28))
	v1140 = F_lappend_oid(m, v1138, v1139)
	mBase = m.M
	v1141 = m.ExcPending
	if v1141 != 0 {
		goto L11
	} else {
		goto L251
	}
L251:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v1140
	v1143 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1144 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1145 = *(*int32)(unsafe.Add(mBase, uint32(v1144)+32))
	v1146 = F_lappend_int(m, v1143, v1145)
	mBase = m.M
	v1147 = m.ExcPending
	if v1147 != 0 {
		goto L11
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v1146
	v1149 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v1150 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1151 = *(*int32)(unsafe.Add(mBase, uint32(v1150)+36))
	v1152 = F_lappend_oid(m, v1149, v1151)
	mBase = m.M
	v1153 = m.ExcPending
	if v1153 != 0 {
		goto L11
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v1152
	v1155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+8)))
	if v1155 == int32(0) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v1158 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	v1159 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1160 = *(*int32)(unsafe.Add(mBase, uint32(v1159)+28))
	v1161 = F_makeSortGroupClauseForSetOp(m, v1160)
	mBase = m.M
	v1162 = m.ExcPending
	if v1162 != 0 {
		goto L11
	} else {
		goto L257
	}
L255:
	;
	goto L256
L256:
	;
	v1166 = *(*int32)(unsafe.Add(mBase, uint32(v89)+20))
	v1168 = F_lappend_oid(m, v1166, int32(2287))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L11
	} else {
		goto L259
	}
L257:
	;
	v1163 = F_lappend(m, v1158, v1161)
	mBase = m.M
	v1164 = m.ExcPending
	if v1164 != 0 {
		goto L11
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v1163
	goto L256
L259:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+20)) = v1168
	v1171 = *(*int32)(unsafe.Add(mBase, uint32(v89)+24))
	v1173 = F_lappend_int(m, v1171, int32(-1))
	mBase = m.M
	v1174 = m.ExcPending
	if v1174 != 0 {
		goto L11
	} else {
		goto L260
	}
L260:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = v1173
	v1176 = *(*int32)(unsafe.Add(mBase, uint32(v89)+28))
	v1178 = F_lappend_oid(m, v1176, int32(0))
	mBase = m.M
	v1179 = m.ExcPending
	if v1179 != 0 {
		goto L11
	} else {
		goto L261
	}
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+28)) = v1178
	v1181 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+8)))
	if v1181 != 0 {
		goto L249
	} else {
		goto L262
	}
L262:
	;
	v1182 = *(*int32)(unsafe.Add(mBase, uint32(v89)+32))
	v1184 = F_makeSortGroupClauseForSetOp(m, int32(2287))
	mBase = m.M
	v1185 = m.ExcPending
	if v1185 != 0 {
		goto L11
	} else {
		goto L263
	}
L263:
	;
	v1186 = F_lappend(m, v1182, v1184)
	mBase = m.M
	v1187 = m.ExcPending
	if v1187 != 0 {
		goto L11
	} else {
		goto L264
	}
L264:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v89)+32)) = v1186
	goto L249
L265:
	;
	v1190 = *(*int32)(unsafe.Add(mBase, uint32(v86)+76))
	v1194 = int32(0)
	v1196 = F_makeVar(m, int32(1), base.I32_extend16_s(v119), v118, int32(-1), v1194, v1194)
	mBase = m.M
	v1197 = m.ExcPending
	if v1197 != 0 {
		goto L11
	} else {
		goto L268
	}
L266:
	;
	goto L267
L267:
	;
	v1214 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v1214 != 0 {
		goto L274
	} else {
		goto L275
	}
L268:
	;
	v1198 = *(*int32)(unsafe.Add(mBase, uint32(v86)+76))
	if v1198 != 0 {
		goto L269
	} else {
		goto L270
	}
L269:
	;
	v1199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1198)+4)))
	v1203 = v1199 + int32(1)
	goto L271
L270:
	;
	v1203 = int32(1)
	goto L271
L271:
	;
	v1205 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	v1206 = *(*int32)(unsafe.Add(mBase, uint32(v1205)+12))
	v1208 = F_makeTargetEntry(m, v1196, base.I32_extend16_s(v1203), v1206, int32(0))
	mBase = m.M
	v1209 = m.ExcPending
	if v1209 != 0 {
		goto L11
	} else {
		goto L272
	}
L272:
	;
	v1210 = F_lappend(m, v1190, v1208)
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L11
	} else {
		goto L273
	}
L273:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+76)) = v1210
	goto L267
L274:
	;
	v1215 = *(*int32)(unsafe.Add(mBase, uint32(v86)+76))
	v1218 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+28))
	v1219 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+32))
	v1220 = *(*int32)(unsafe.Add(mBase, uint32(v1214)+36))
	v1222 = F_makeVar(m, int32(1), base.I32_extend16_s(v145), v1218, v1219, v1220, int32(0))
	mBase = m.M
	v1223 = m.ExcPending
	if v1223 != 0 {
		goto L11
	} else {
		goto L277
	}
L275:
	;
	goto L276
L276:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+40)) = v499
	v1264 = *(*int32)(unsafe.Add(mBase, uint32(v82)+20))
	if v1264 != 0 {
		goto L289
	} else {
		goto L290
	}
L277:
	;
	v1224 = *(*int32)(unsafe.Add(mBase, uint32(v86)+76))
	if v1224 != 0 {
		goto L278
	} else {
		goto L279
	}
L278:
	;
	v1225 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1224)+4)))
	v1229 = v1225 + int32(1)
	goto L280
L279:
	;
	v1229 = int32(1)
	goto L280
L280:
	;
	v1231 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1232 = *(*int32)(unsafe.Add(mBase, uint32(v1231)+8))
	v1234 = F_makeTargetEntry(m, v1222, base.I32_extend16_s(v1229), v1232, int32(0))
	mBase = m.M
	v1235 = m.ExcPending
	if v1235 != 0 {
		goto L11
	} else {
		goto L281
	}
L281:
	;
	v1236 = F_lappend(m, v1215, v1234)
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L11
	} else {
		goto L282
	}
L282:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+76)) = v1236
	v1243 = int32(0)
	v1245 = F_makeVar(m, int32(1), base.I32_extend16_s(v144), int32(2287), int32(-1), v1243, v1243)
	mBase = m.M
	v1246 = m.ExcPending
	if v1246 != 0 {
		goto L11
	} else {
		goto L283
	}
L283:
	;
	v1247 = *(*int32)(unsafe.Add(mBase, uint32(v86)+76))
	if v1247 != 0 {
		goto L284
	} else {
		goto L285
	}
L284:
	;
	v1248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1247)+4)))
	v1252 = v1248 + int32(1)
	goto L286
L285:
	;
	v1252 = int32(1)
	goto L286
L286:
	;
	v1254 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1255 = *(*int32)(unsafe.Add(mBase, uint32(v1254)+20))
	v1257 = F_makeTargetEntry(m, v1245, base.I32_extend16_s(v1252), v1255, int32(0))
	mBase = m.M
	v1258 = m.ExcPending
	if v1258 != 0 {
		goto L11
	} else {
		goto L287
	}
L287:
	;
	v1259 = F_lappend(m, v1236, v1257)
	mBase = m.M
	v1260 = m.ExcPending
	if v1260 != 0 {
		goto L11
	} else {
		goto L288
	}
L288:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v86)+76)) = v1259
	goto L276
L289:
	;
	v1265 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v1266 = F_lappend_oid(m, v1265, v118)
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L11
	} else {
		goto L292
	}
L290:
	;
	goto L291
L291:
	;
	v1279 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	if v1279 != 0 {
		goto L295
	} else {
		goto L296
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+44)) = v1266
	v1269 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	v1271 = F_lappend_int(m, v1269, int32(-1))
	mBase = m.M
	v1272 = m.ExcPending
	if v1272 != 0 {
		goto L11
	} else {
		goto L293
	}
L293:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+48)) = v1271
	v1274 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
	v1276 = F_lappend_oid(m, v1274, int32(0))
	mBase = m.M
	v1277 = m.ExcPending
	if v1277 != 0 {
		goto L11
	} else {
		goto L294
	}
L294:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+52)) = v1276
	goto L291
L295:
	;
	v1280 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v1281 = *(*int32)(unsafe.Add(mBase, uint32(v1279)+28))
	v1282 = F_lappend_oid(m, v1280, v1281)
	mBase = m.M
	v1283 = m.ExcPending
	if v1283 != 0 {
		goto L11
	} else {
		goto L298
	}
L296:
	;
	goto L297
L297:
	;
	m.G0 = v80 + int32(128)
	*(*int32)(unsafe.Add(mBase, uint32(v65))) = v82
	goto L6
L298:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+44)) = v1282
	v1285 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	v1286 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1287 = *(*int32)(unsafe.Add(mBase, uint32(v1286)+32))
	v1288 = F_lappend_int(m, v1285, v1287)
	mBase = m.M
	v1289 = m.ExcPending
	if v1289 != 0 {
		goto L11
	} else {
		goto L299
	}
L299:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+48)) = v1288
	v1291 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
	v1292 = *(*int32)(unsafe.Add(mBase, uint32(v82)+24))
	v1293 = *(*int32)(unsafe.Add(mBase, uint32(v1292)+36))
	v1294 = F_lappend_oid(m, v1291, v1293)
	mBase = m.M
	v1295 = m.ExcPending
	if v1295 != 0 {
		goto L11
	} else {
		goto L300
	}
L300:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+52)) = v1294
	v1297 = *(*int32)(unsafe.Add(mBase, uint32(v82)+44))
	v1299 = F_lappend_oid(m, v1297, int32(2287))
	mBase = m.M
	v1300 = m.ExcPending
	if v1300 != 0 {
		goto L11
	} else {
		goto L301
	}
L301:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+44)) = v1299
	v1302 = *(*int32)(unsafe.Add(mBase, uint32(v82)+48))
	v1304 = F_lappend_int(m, v1302, int32(-1))
	mBase = m.M
	v1305 = m.ExcPending
	if v1305 != 0 {
		goto L11
	} else {
		goto L302
	}
L302:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+48)) = v1304
	v1307 = *(*int32)(unsafe.Add(mBase, uint32(v82)+52))
	v1309 = F_lappend_oid(m, v1307, int32(0))
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L11
	} else {
		goto L303
	}
L303:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v82)+52)) = v1309
	goto L297
L304:
	;
	goto L5
L305:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2661 = m.ExcPending
	if v2661 != 0 {
		goto L11
	} else {
		goto L607
	}
L306:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2645 = m.ExcPending
	if v2645 != 0 {
		goto L11
	} else {
		goto L604
	}
L307:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2624 = m.ExcPending
	if v2624 != 0 {
		goto L11
	} else {
		goto L600
	}
L308:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2611 = m.ExcPending
	if v2611 != 0 {
		goto L11
	} else {
		goto L597
	}
L309:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2598 = m.ExcPending
	if v2598 != 0 {
		goto L11
	} else {
		goto L594
	}
L310:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2577 = m.ExcPending
	if v2577 != 0 {
		goto L11
	} else {
		goto L590
	}
L311:
	;
	v1971 = *(*int32)(unsafe.Add(mBase, uint32(l0)+48))
	if v1971 == int32(0) {
		goto L445
	} else {
		goto L446
	}
L312:
	;
	v1374 = l1
	v1380 = v1370
	v1386 = v3
	goto L313
L313:
	;
	v1398 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+4))
	if v1398 <= v1386 {
		v1947 = v1374
		goto L311
	} else {
		goto L315
	}
L314:
	;
	v1947 = v1921
	goto L311
L315:
	;
	v1401 = v1386 << (uint(int32(2)) % 32)
	v1403 = v1386 + int32(1)
	v1404 = *(*int32)(unsafe.Add(mBase, uint32(v1380)+12))
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v1401+v1404)))
	v1407 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+12))
	switch v1407 {
	case 0:
		goto L317
	case 1:
		goto L318
	default:
		v1921 = v1374
		goto L316
	}
L316:
	;
	v1945 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v1945 != 0 {
		v1374 = v1921
		v1380 = v1945
		v1386 = v1403
		goto L313
	} else {
		goto L444
	}
L317:
	;
	v1416 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1406)+21)))
	if v1416 == int32(109) {
		v1921 = v1374
		goto L316
	} else {
		goto L320
	}
L318:
	;
	v1408 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+36))
	v1409 = F_fireRIRrules(m, v1408, v1374)
	mBase = m.M
	v1410 = m.ExcPending
	if v1410 != 0 {
		goto L11
	} else {
		goto L319
	}
L319:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1406)+36)) = v1409
	v1412 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)))
	v1413 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1409)+44)))
	v1414 = v1412 | v1413
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v1414)
	v1921 = v1374
	goto L316
L320:
	;
	v1419 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v1419 != 0 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v1420 = *(*int32)(unsafe.Add(mBase, uint32(v1419)+28))
	if v1403 == v1420 {
		v1921 = v1374
		goto L316
	} else {
		goto L324
	}
L322:
	;
	goto L323
L323:
	;
	v1422 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1422 != v1403 {
		goto L325
	} else {
		goto L326
	}
L324:
	;
	goto L323
L325:
	;
	v1424 = F_rangeTableEntry_used(m, l0, v1403)
	mBase = m.M
	v1425 = m.ExcPending
	if v1425 != 0 {
		goto L11
	} else {
		goto L328
	}
L326:
	;
	v1431 = int32(0)
	goto L327
L327:
	;
	if base.B2i32(v1431 == int32(0))&base.B2i32(v1403 != v30) != 0 {
		v1921 = v1374
		goto L316
	} else {
		goto L330
	}
L328:
	;
	if v1424 == int32(0) {
		v1921 = v1374
		goto L316
	} else {
		goto L329
	}
L329:
	;
	v1428 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	v1431 = base.B2i32(v1403 != v1428)
	goto L327
L330:
	;
	v1436 = *(*int32)(unsafe.Add(mBase, uint32(v1406)+16))
	v1438 = F_table_open(m, v1436, int32(0))
	mBase = m.M
	v1439 = m.ExcPending
	if v1439 != 0 {
		goto L11
	} else {
		goto L332
	}
L331:
	;
	F_relation_close(m, v1438, int32(0))
	mBase = m.M
	v1919 = m.ExcPending
	if v1919 != 0 {
		goto L11
	} else {
		goto L443
	}
L332:
	;
	v1440 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+68))
	if v1440 == int32(0) {
		v1893 = v1374
		goto L331
	} else {
		goto L333
	}
L333:
	;
	v1443 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	if v1443 <= int32(0) {
		v1893 = v1374
		goto L331
	} else {
		goto L334
	}
L334:
	;
	v1446 = int32(0)
	v1452 = v1443
	v1454 = v1446
	v1455 = v1446
	goto L335
L335:
	;
	v1473 = *(*int32)(unsafe.Add(mBase, uint32(v1440)+4))
	v1477 = *(*int32)(unsafe.Add(mBase, uint32(v1473+v1455<<(uint(int32(2))%32))))
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v1477)+4))
	if v1478 == int32(1) {
		goto L337
	} else {
		goto L338
	}
L336:
	;
	if v1485 == int32(0) {
		v1893 = v1374
		goto L331
	} else {
		goto L342
	}
L337:
	;
	v1481 = F_lappend(m, v1454, v1477)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L11
	} else {
		goto L340
	}
L338:
	;
	v1484 = v1452
	v1485 = v1454
	goto L339
L339:
	;
	v1487 = v1455 + int32(1)
	if v1487 < v1484 {
		v1452 = v1484
		v1454 = v1485
		v1455 = v1487
		goto L335
	} else {
		goto L341
	}
L340:
	;
	v1483 = *(*int32)(unsafe.Add(mBase, uint32(v1440)))
	v1484 = v1483
	v1485 = v1481
	goto L339
L341:
	;
	goto L336
L342:
	;
	v1491 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+56))
	v1492 = int32(0)
	if v1374 == v1492 {
		goto L344
	} else {
		goto L345
	}
L343:
	;
	if v1530 != 0 {
		goto L310
	} else {
		goto L356
	}
L344:
	;
	v1530 = int32(0)
	goto L343
L345:
	;
	goto L346
L346:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+4))
	if v1498 <= int32(0) {
		v1524 = v1492
		goto L347
	} else {
		goto L348
	}
L347:
	;
	v1530 = v1524
	goto L343
L348:
	;
	v1501 = int32(0)
	if v1501 < v1498 {
		goto L349
	} else {
		goto L350
	}
L349:
	;
	v1504 = v1498
	goto L351
L350:
	;
	v1504 = v1501
	goto L351
L351:
	;
	v1505 = *(*int32)(unsafe.Add(mBase, uint32(v1374)+12))
	v1507 = int32(0)
	goto L352
L352:
	;
	v1515 = *(*int32)(unsafe.Add(mBase, uint32(v1505+v1507<<(uint(int32(2))%32))))
	v1516 = base.B2i32(v1515 == v1491)
	if v1515 == v1491 {
		v1524 = v1516
		goto L347
	} else {
		goto L354
	}
L353:
	;
	v1524 = v1516
	goto L347
L354:
	;
	v1518 = v1507 + int32(1)
	if v1518 != v1504 {
		v1507 = v1518
		goto L352
	} else {
		goto L355
	}
L355:
	;
	goto L353
L356:
	;
	v1531 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+56))
	v1532 = F_lappend_oid(m, v1374, v1531)
	mBase = m.M
	v1533 = m.ExcPending
	if v1533 != 0 {
		goto L11
	} else {
		goto L357
	}
L357:
	;
	v1534 = *(*int32)(unsafe.Add(mBase, uint32(v1485)+4))
	if int32(0) < v1534 {
		goto L358
	} else {
		goto L359
	}
L358:
	;
	v1539 = int32(0)
	v1542 = v1534
	goto L361
L359:
	;
	goto L360
L360:
	;
	v1890 = F_list_delete_last(m, v1532)
	mBase = m.M
	v1891 = m.ExcPending
	if v1891 != 0 {
		goto L11
	} else {
		goto L442
	}
L361:
	;
	v1563 = *(*int32)(unsafe.Add(mBase, uint32(v1485)+12))
	v1567 = *(*int32)(unsafe.Add(mBase, uint32(v1563+v1539<<(uint(int32(2))%32))))
	v1568 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+12))
	if v1568 == int32(0) {
		goto L309
	} else {
		goto L363
	}
L362:
	;
	goto L360
L363:
	;
	v1571 = *(*int32)(unsafe.Add(mBase, uint32(v1568)+4))
	if v1571 != int32(1) {
		goto L309
	} else {
		goto L364
	}
L364:
	;
	v1574 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+8))
	if v1574 != 0 {
		goto L308
	} else {
		goto L365
	}
L365:
	;
	v1576 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_fireRIRrules[0])))
	if v1576&int32(1) != 0 {
		goto L366
	} else {
		goto L367
	}
L366:
	;
	v1579 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+56))
	if base.Ui32(int32(_a_F_fireRIRrules_6)) <= base.Ui32(v1579) {
		goto L307
	} else {
		goto L369
	}
L367:
	;
	goto L368
L368:
	;
	v1582 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v1582 == v1403 {
		goto L371
	} else {
		goto L372
	}
L369:
	;
	goto L368
L370:
	;
	v1863 = v1539 + int32(1)
	if v1863 < v1841 {
		v1539 = v1863
		v1542 = v1841
		goto L361
	} else {
		goto L441
	}
L371:
	;
	v1584 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	switch v1584 - int32(2) {
	case 0, 2, 3:
		goto L374
	case 1:
		v1841 = v1542
		goto L370
	default:
		goto L306
	}
L372:
	;
	goto L373
L373:
	;
	v1634 = *(*int32)(unsafe.Add(mBase, uint32(l0)+140))
	if v1634 != 0 {
		goto L391
	} else {
		goto L392
	}
L374:
	;
	v1587 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1588 = *(*int32)(unsafe.Add(mBase, uint32(v1587)+12))
	v1590 = *(*int32)(unsafe.Add(mBase, uint32(v1588+v1401)))
	v1591 = F_copyObjectImpl(m, v1590)
	mBase = m.M
	v1592 = m.ExcPending
	if v1592 != 0 {
		goto L11
	} else {
		goto L375
	}
L375:
	;
	v1593 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1594 = F_lappend(m, v1593, v1591)
	mBase = m.M
	v1595 = m.ExcPending
	if v1595 != 0 {
		goto L11
	} else {
		goto L376
	}
L376:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+52)) = v1594
	if v1594 != 0 {
		goto L377
	} else {
		goto L378
	}
L377:
	;
	v1597 = *(*int32)(unsafe.Add(mBase, uint32(v1594)+4))
	v1599 = v1597
	goto L379
L378:
	;
	v1599 = int32(0)
	goto L379
L379:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+32)) = v1599
	v1601 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	v1602 = F_copyObjectImpl(m, v1601)
	mBase = m.M
	v1603 = m.ExcPending
	if v1603 != 0 {
		goto L11
	} else {
		goto L380
	}
L380:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+96)) = v1602
	v1605 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	F_ChangeVarNodes(m, v1602, v1403, v1605)
	mBase = m.M
	v1607 = m.ExcPending
	if v1607 != 0 {
		goto L11
	} else {
		goto L381
	}
L381:
	;
	v1608 = int32(0)
	v1610 = F_makeWholeRowVar(m, v1590, v1403, v1608, v1608)
	mBase = m.M
	v1611 = m.ExcPending
	if v1611 != 0 {
		goto L11
	} else {
		goto L382
	}
L382:
	;
	v1612 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	if v1612 != 0 {
		goto L383
	} else {
		goto L384
	}
L383:
	;
	v1613 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1612)+4)))
	v1617 = v1613 + int32(1)
	goto L385
L384:
	;
	v1617 = int32(1)
	goto L385
L385:
	;
	v1620 = F_pstrdup(m, int32(_a_F_fireRIRrules_7))
	mBase = m.M
	v1621 = m.ExcPending
	if v1621 != 0 {
		goto L11
	} else {
		goto L386
	}
L386:
	;
	v1623 = F_makeTargetEntry(m, v1610, base.I32_extend16_s(v1617), v1620, int32(1))
	mBase = m.M
	v1624 = m.ExcPending
	if v1624 != 0 {
		goto L11
	} else {
		goto L387
	}
L387:
	;
	v1625 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	v1626 = F_lappend(m, v1625, v1623)
	mBase = m.M
	v1627 = m.ExcPending
	if v1627 != 0 {
		goto L11
	} else {
		goto L388
	}
L388:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+76)) = v1626
	goto L373
L389:
	;
	v1669 = *(*int32)(unsafe.Add(mBase, uint32(v1567)+12))
	v1670 = *(*int32)(unsafe.Add(mBase, uint32(v1669)+12))
	v1671 = *(*int32)(unsafe.Add(mBase, uint32(v1670)))
	v1672 = F_copyObjectImpl(m, v1671)
	mBase = m.M
	v1673 = m.ExcPending
	if v1673 != 0 {
		goto L11
	} else {
		goto L402
	}
L390:
	;
	goto L389
L391:
	;
	v1635 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+4))
	if v1635 <= int32(0) {
		v1667 = int32(0)
		goto L390
	} else {
		goto L394
	}
L392:
	;
	goto L393
L393:
	;
	v1667 = int32(0)
	goto L390
L394:
	;
	v1638 = int32(0)
	if v1638 < v1635 {
		goto L395
	} else {
		goto L396
	}
L395:
	;
	v1641 = v1635
	goto L397
L396:
	;
	v1641 = v1638
	goto L397
L397:
	;
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v1634)+12))
	v1644 = int32(0)
	goto L398
L398:
	;
	v1652 = *(*int32)(unsafe.Add(mBase, uint32(v1642+v1644<<(uint(int32(2))%32))))
	v1653 = *(*int32)(unsafe.Add(mBase, uint32(v1652)+4))
	if v1653 == v1403 {
		v1667 = v1652
		goto L390
	} else {
		goto L400
	}
L399:
	;
	goto L393
L400:
	;
	v1656 = v1644 + int32(1)
	if v1656 != v1641 {
		v1644 = v1656
		goto L398
	} else {
		goto L401
	}
L401:
	;
	goto L399
L402:
	;
	F_AcquireRewriteLocks(m, v1672, int32(1), base.B2i32(v1667 != int32(0)))
	mBase = m.M
	v1678 = m.ExcPending
	if v1678 != 0 {
		goto L11
	} else {
		goto L403
	}
L403:
	;
	if v1667 != 0 {
		goto L404
	} else {
		goto L405
	}
L404:
	;
	v1679 = *(*int32)(unsafe.Add(mBase, uint32(v1672)+60))
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v1667)+8))
	v1681 = *(*int32)(unsafe.Add(mBase, uint32(v1667)+12))
	F_markQueryForLocking(m, v1672, v1679, v1680, v1681)
	mBase = m.M
	v1683 = m.ExcPending
	if v1683 != 0 {
		goto L11
	} else {
		goto L407
	}
L405:
	;
	goto L406
L406:
	;
	v1684 = F_fireRIRrules(m, v1672, v1532)
	mBase = m.M
	v1685 = m.ExcPending
	if v1685 != 0 {
		goto L11
	} else {
		goto L408
	}
L407:
	;
	goto L406
L408:
	;
	v1686 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)))
	v1687 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1684)+44)))
	v1688 = v1686 | v1687
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v1688)
	v1690 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	v1691 = *(*int32)(unsafe.Add(mBase, uint32(v1690)+12))
	v1693 = *(*int32)(unsafe.Add(mBase, uint32(v1691+v1401)))
	*(*int32)(unsafe.Add(mBase, uint32(v1693)+36)) = v1684
	*(*int32)(unsafe.Add(mBase, uint32(v1693)+12)) = int32(1)
	v1697 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+180))
	if v1697 != 0 {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1698 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1697)+4)))
	v1700 = v1698
	goto L411
L410:
	;
	v1700 = int32(0)
	goto L411
L411:
	;
	v1701 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v1693)+32)) = v1701
	*(*uint8)(unsafe.Add(mBase, uint32(v1693)+40)) = uint8(v1700)
	*(*uint8)(unsafe.Add(mBase, uint32(v1693)+20)) = uint8(v1701)
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v1684)+76))
	if v1706 == v1701 {
		goto L413
	} else {
		goto L414
	}
L412:
	;
	goto L430
L413:
	;
	v1795 = int32(0)
	goto L412
L414:
	;
	goto L415
L415:
	;
	v1717 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+4))
	if v1717 <= int32(0) {
		goto L416
	} else {
		goto L417
	}
L416:
	;
	v1795 = int32(0)
	goto L412
L417:
	;
	goto L418
L418:
	;
	if v1717 != int32(1) {
		goto L420
	} else {
		goto L421
	}
L419:
	;
	v1795 = v1782
	goto L412
L420:
	;
	v1723 = int32(0)
	if v1723 < v1717 {
		goto L423
	} else {
		goto L424
	}
L421:
	;
	v1764 = v1701
	v1765 = v1701
	goto L422
L422:
	;
	v1770 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+12))
	v1774 = *(*int32)(unsafe.Add(mBase, uint32(v1770+v1764<<(uint(int32(2))%32))))
	v1775 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1774)+26)))
	v1782 = v1765 + (v1775 ^ int32(1))
	goto L419
L423:
	;
	v1726 = v1717
	goto L425
L424:
	;
	v1726 = v1723
	goto L425
L425:
	;
	v1731 = *(*int32)(unsafe.Add(mBase, uint32(v1706)+12))
	v1732 = int32(0)
	v1735 = v1732
	v1736 = v1732
	v1737 = v1701
	goto L426
L426:
	;
	v1742 = int32(2)
	v1744 = v1731 + v1736<<(uint(v1742)%32)
	v1745 = *(*int32)(unsafe.Add(mBase, uint32(v1744)))
	v1746 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1745)+26)))
	v1747 = int32(1)
	v1750 = *(*int32)(unsafe.Add(mBase, uint32(v1744)+4))
	v1751 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1750)+26)))
	v1754 = v1737 + (v1746 ^ v1747) + (v1751 ^ v1747)
	v1756 = v1736 + v1742
	v1758 = v1735 + v1742
	if v1758 != v1726&int32(2147483646) {
		v1735 = v1758
		v1736 = v1756
		v1737 = v1754
		goto L426
	} else {
		goto L428
	}
L427:
	;
	if v1726&int32(1) == int32(0) {
		v1782 = v1754
		goto L419
	} else {
		goto L429
	}
L428:
	;
	goto L427
L429:
	;
	v1764 = v1756
	v1765 = v1754
	goto L422
L430:
	;
	v1821 = *(*int32)(unsafe.Add(mBase, uint32(v1693)+8))
	v1822 = *(*int32)(unsafe.Add(mBase, uint32(v1821)+8))
	if v1822 != 0 {
		goto L432
	} else {
		goto L433
	}
L431:
	;
	v1836 = *(*int32)(unsafe.Add(mBase, uint32(v1485)+4))
	v1841 = v1836
	goto L370
L432:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v1822)+4))
	v1825 = v1823
	goto L434
L433:
	;
	v1825 = int32(0)
	goto L434
L434:
	;
	if v1825 < v1795 {
		goto L435
	} else {
		goto L436
	}
L435:
	;
	v1828 = F_pstrdup(m, int32(_a_F_fireRIRrules_8))
	mBase = m.M
	v1829 = m.ExcPending
	if v1829 != 0 {
		goto L11
	} else {
		goto L438
	}
L436:
	;
	goto L437
L437:
	;
	goto L431
L438:
	;
	v1830 = F_makeString(m, v1828)
	mBase = m.M
	v1831 = m.ExcPending
	if v1831 != 0 {
		goto L11
	} else {
		goto L439
	}
L439:
	;
	v1832 = F_lappend(m, v1822, v1830)
	mBase = m.M
	v1833 = m.ExcPending
	if v1833 != 0 {
		goto L11
	} else {
		goto L440
	}
L440:
	;
	v1834 = *(*int32)(unsafe.Add(mBase, uint32(v1693)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v1834)+8)) = v1832
	goto L430
L441:
	;
	goto L362
L442:
	;
	v1893 = v1890
	goto L331
L443:
	;
	v1921 = v1893
	goto L316
L444:
	;
	goto L314
L445:
	;
	v2045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)))
	if v2045 == int32(1) {
		goto L452
	} else {
		goto L453
	}
L446:
	;
	v1974 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+4))
	if v1974 <= int32(0) {
		goto L445
	} else {
		goto L447
	}
L447:
	;
	v1985 = int32(0)
	goto L448
L448:
	;
	v2003 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+12))
	v2007 = *(*int32)(unsafe.Add(mBase, uint32(v2003+v1985<<(uint(int32(2))%32))))
	v2008 = *(*int32)(unsafe.Add(mBase, uint32(v2007)+16))
	v2009 = F_fireRIRrules(m, v2008, v1947)
	mBase = m.M
	v2010 = m.ExcPending
	if v2010 != 0 {
		goto L11
	} else {
		goto L450
	}
L449:
	;
	goto L445
L450:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2007)+16)) = v2009
	v2012 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)))
	v2013 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2009)+44)))
	v2014 = v2012 | v2013
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v2014)
	v2017 = v1985 + int32(1)
	v2018 = *(*int32)(unsafe.Add(mBase, uint32(v1971)+4))
	if v2017 < v2018 {
		v1985 = v2017
		goto L448
	} else {
		goto L451
	}
L451:
	;
	goto L449
L452:
	;
	v2048 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+76)) = uint8(v2048)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = v1947
	v2055 = F_query_tree_walker_impl(m, l0, int32(1042), v28+int32(72), int32(3))
	mBase = m.M
	v2056 = m.ExcPending
	if v2056 != 0 {
		goto L11
	} else {
		goto L455
	}
L453:
	;
	goto L454
L454:
	;
	v2061 = *(*int32)(unsafe.Add(mBase, uint32(l0)+52))
	if v2061 == int32(0) {
		goto L456
	} else {
		goto L457
	}
L455:
	;
	v2057 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)))
	v2058 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+76)))
	v2059 = v2057 | v2058
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v2059)
	goto L454
L456:
	;
	m.G0 = v28 + int32(80)
	return l0
L457:
	;
	v2064 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+4))
	if v2064 <= int32(0) {
		goto L456
	} else {
		goto L458
	}
L458:
	;
	v2069 = v1947
	v2075 = int32(0)
	goto L459
L459:
	;
	v2094 = v2075 + int32(1)
	v2095 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+12))
	v2099 = *(*int32)(unsafe.Add(mBase, uint32(v2095+v2075<<(uint(int32(2))%32))))
	v2100 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+12))
	if v2100 != 0 {
		v2529 = v2069
		goto L461
	} else {
		goto L462
	}
L460:
	;
	goto L456
L461:
	;
	v2543 = *(*int32)(unsafe.Add(mBase, uint32(v2061)+4))
	if v2094 < v2543 {
		v2069 = v2529
		v2075 = v2094
		goto L459
	} else {
		goto L589
	}
L462:
	;
	v2101 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2099)+21)))
	switch v2101 - int32(112) {
	case 0, 2:
		goto L463
	default:
		v2529 = v2069
		goto L461
	}
L463:
	;
	v2104 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+16))
	v2106 = F_table_open(m, v2104, int32(0))
	mBase = m.M
	v2107 = m.ExcPending
	if v2107 != 0 {
		goto L11
	} else {
		goto L464
	}
L464:
	;
	v2108 = int32(0)
	v2109 = m.G0
	v2111 = v2109 - int32(48)
	m.G0 = v2111
	v2114 = v28 + int32(68)
	*(*int32)(unsafe.Add(mBase, uint32(v2114))) = v2108
	v2118 = v28 - int32(-64)
	*(*int32)(unsafe.Add(mBase, uint32(v2118))) = v2108
	v2122 = v28 + int32(63)
	*(*uint8)(unsafe.Add(mBase, uint32(v2122))) = uint8(v2108)
	v2126 = v28 + int32(62)
	*(*uint8)(unsafe.Add(mBase, uint32(v2126))) = uint8(v2108)
	v2129 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2099)+21)))
	switch v2129 - int32(112) {
	case 0, 2:
		goto L466
	default:
		goto L465
	}
L465:
	;
	m.G0 = v2111 + int32(48)
	v2407 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
	v2408 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	if v2407|v2408 != 0 {
		goto L540
	} else {
		goto L541
	}
L466:
	;
	v2132 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v2133 = F_getRTEPermissionInfo(m, v2132, v2099)
	mBase = m.M
	v2134 = m.ExcPending
	if v2134 != 0 {
		goto L11
	} else {
		goto L467
	}
L467:
	;
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v2133)+24))
	v2136 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+16))
	if v2135 != 0 {
		goto L470
	} else {
		goto L471
	}
L468:
	;
	v2395 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v2122))) = uint8(v2395)
	goto L465
L469:
	;
	v2145 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+16))
	v2147 = F_table_open(m, v2145, int32(0))
	mBase = m.M
	v2148 = m.ExcPending
	if v2148 != 0 {
		goto L11
	} else {
		goto L474
	}
L470:
	;
	v2140 = v2135
	v2141 = v2135
	goto L472
L471:
	;
	v2138 = *(*int32)(unsafe.Add(mBase, _c_F_fireRIRrules[1]))
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v2133)+24))
	v2140 = v2138
	v2141 = v2139
	goto L472
L472:
	;
	v2143 = F_check_enable_rls(m, v2136, v2141, int32(0))
	mBase = m.M
	v2144 = m.ExcPending
	if v2144 != 0 {
		goto L11
	} else {
		goto L473
	}
L473:
	;
	switch v2143 {
	case 0:
		goto L465
	case 1:
		goto L468
	default:
		goto L469
	}
L474:
	;
	v2149 = *(*int32)(unsafe.Add(mBase, uint32(l0)+32))
	if v2149 == v2094 {
		goto L478
	} else {
		goto L479
	}
L475:
	;
	if base.B2i32(int32(1)<<(uint(v2190)%32)&int32(52) == int32(0))|base.B2i32(base.Ui32(int32(5)) < base.Ui32(v2190)) != 0 {
		goto L490
	} else {
		goto L491
	}
L476:
	;
	v2186 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+44))
	v2187 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+40))
	F_add_security_quals(m, v2094, v2186, v2187, v2114, v2126)
	mBase = m.M
	v2189 = m.ExcPending
	if v2189 != 0 {
		goto L11
	} else {
		goto L489
	}
L477:
	;
	F_get_policies_for_relation(m, v2147, v2151, v2140, v2111+int32(44), v2111+int32(40))
	mBase = m.M
	v2182 = m.ExcPending
	if v2182 != 0 {
		goto L11
	} else {
		goto L488
	}
L478:
	;
	v2151 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v2151 != int32(1) {
		goto L477
	} else {
		goto L481
	}
L479:
	;
	goto L480
L480:
	;
	v2155 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2133)+16)))
	if v2155&int32(4) != 0 {
		goto L482
	} else {
		goto L483
	}
L481:
	;
	goto L480
L482:
	;
	F_get_policies_for_relation(m, v2147, int32(2), v2140, v2111+int32(44), v2111+int32(40))
	mBase = m.M
	v2164 = m.ExcPending
	if v2164 != 0 {
		goto L11
	} else {
		goto L485
	}
L483:
	;
	goto L484
L484:
	;
	v2169 = int32(1)
	F_get_policies_for_relation(m, v2147, v2169, v2140, v2111+int32(44), v2111+int32(40))
	mBase = m.M
	v2176 = m.ExcPending
	if v2176 != 0 {
		goto L11
	} else {
		goto L487
	}
L485:
	;
	v2165 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+44))
	v2166 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+40))
	F_add_security_quals(m, v2094, v2165, v2166, v2114, v2126)
	mBase = m.M
	v2168 = m.ExcPending
	if v2168 != 0 {
		goto L11
	} else {
		goto L486
	}
L486:
	;
	goto L484
L487:
	;
	v2185 = v2169
	goto L476
L488:
	;
	switch v2151 - int32(2) {
	case 0, 2:
		v2185 = v2151
		goto L476
	default:
		v2190 = v2151
		goto L475
	}
L489:
	;
	v2190 = v2185
	goto L475
L490:
	;
	if v2190&int32(-2) == int32(2) {
		goto L496
	} else {
		goto L497
	}
L491:
	;
	v2200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2133)+16)))
	if v2200&int32(2) == int32(0) {
		goto L490
	} else {
		goto L492
	}
L492:
	;
	F_get_policies_for_relation(m, v2147, int32(1), v2140, v2111+int32(36), v2111+int32(32))
	mBase = m.M
	v2211 = m.ExcPending
	if v2211 != 0 {
		goto L11
	} else {
		goto L493
	}
L493:
	;
	v2212 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+36))
	v2213 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+32))
	F_add_security_quals(m, v2094, v2212, v2213, v2114, v2126)
	mBase = m.M
	v2215 = m.ExcPending
	if v2215 != 0 {
		goto L11
	} else {
		goto L494
	}
L494:
	;
	goto L490
L495:
	;
	F_relation_close(m, v2147, int32(0))
	mBase = m.M
	v2381 = m.ExcPending
	if v2381 != 0 {
		goto L11
	} else {
		goto L537
	}
L496:
	;
	if v2190 == int32(3) {
		goto L499
	} else {
		goto L500
	}
L497:
	;
	goto L498
L498:
	;
	if v2190 != int32(5) {
		goto L495
	} else {
		goto L521
	}
L499:
	;
	v2224 = int32(1)
	goto L501
L500:
	;
	v2224 = int32(2)
	goto L501
L501:
	;
	v2225 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+44))
	v2226 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+40))
	F_add_with_check_options(m, v2147, v2094, v2224, v2225, v2226, v2118, v2126, int32(0))
	mBase = m.M
	v2229 = m.ExcPending
	if v2229 != 0 {
		goto L11
	} else {
		goto L502
	}
L502:
	;
	v2230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2133)+16)))
	if v2230&int32(2) != 0 {
		goto L503
	} else {
		goto L504
	}
L503:
	;
	F_get_policies_for_relation(m, v2147, int32(1), v2140, v2111+int32(36), v2111+int32(32))
	mBase = m.M
	v2239 = m.ExcPending
	if v2239 != 0 {
		goto L11
	} else {
		goto L506
	}
L504:
	;
	goto L505
L505:
	;
	if v2190 != int32(3) {
		goto L495
	} else {
		goto L508
	}
L506:
	;
	v2240 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+36))
	v2241 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+32))
	F_add_with_check_options(m, v2147, v2094, v2224, v2240, v2241, v2118, v2126, int32(1))
	mBase = m.M
	v2244 = m.ExcPending
	if v2244 != 0 {
		goto L11
	} else {
		goto L507
	}
L507:
	;
	goto L505
L508:
	;
	v2247 = *(*int32)(unsafe.Add(mBase, uint32(l0)+84))
	if v2247 == int32(0) {
		goto L495
	} else {
		goto L509
	}
L509:
	;
	v2250 = *(*int32)(unsafe.Add(mBase, uint32(v2247)+4))
	if v2250 != int32(2) {
		goto L495
	} else {
		goto L510
	}
L510:
	;
	F_get_policies_for_relation(m, v2147, int32(2), v2140, v2111+int32(36), v2111+int32(32))
	mBase = m.M
	v2259 = m.ExcPending
	if v2259 != 0 {
		goto L11
	} else {
		goto L511
	}
L511:
	;
	v2261 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+36))
	v2262 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+32))
	F_add_with_check_options(m, v2147, v2094, int32(3), v2261, v2262, v2118, v2126, int32(1))
	mBase = m.M
	v2265 = m.ExcPending
	if v2265 != 0 {
		goto L11
	} else {
		goto L512
	}
L512:
	;
	v2267 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2133)+16)))
	if v2267&int32(2) != 0 {
		goto L513
	} else {
		goto L514
	}
L513:
	;
	F_get_policies_for_relation(m, v2147, int32(1), v2140, v2111+int32(28), v2111+int32(24))
	mBase = m.M
	v2276 = m.ExcPending
	if v2276 != 0 {
		goto L11
	} else {
		goto L516
	}
L514:
	;
	v2283 = int32(0)
	v2284 = v2108
	goto L515
L515:
	;
	F_add_with_check_options(m, v2147, v2094, int32(2), v2261, v2262, v2118, v2126, int32(0))
	mBase = m.M
	v2288 = m.ExcPending
	if v2288 != 0 {
		goto L11
	} else {
		goto L518
	}
L516:
	;
	v2278 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+28))
	v2279 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+24))
	F_add_with_check_options(m, v2147, v2094, int32(3), v2278, v2279, v2118, v2126, int32(1))
	mBase = m.M
	v2282 = m.ExcPending
	if v2282 != 0 {
		goto L11
	} else {
		goto L517
	}
L517:
	;
	v2283 = v2279
	v2284 = v2278
	goto L515
L518:
	;
	v2289 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2133)+16)))
	if v2289&int32(2) == int32(0) {
		goto L495
	} else {
		goto L519
	}
L519:
	;
	F_add_with_check_options(m, v2147, v2094, int32(2), v2284, v2283, v2118, v2126, int32(1))
	mBase = m.M
	v2297 = m.ExcPending
	if v2297 != 0 {
		goto L11
	} else {
		goto L520
	}
L520:
	;
	goto L495
L521:
	;
	F_get_policies_for_relation(m, v2147, int32(2), v2140, v2111+int32(36), v2111+int32(32))
	mBase = m.M
	v2306 = m.ExcPending
	if v2306 != 0 {
		goto L11
	} else {
		goto L522
	}
L522:
	;
	v2308 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+36))
	v2309 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+32))
	F_add_with_check_options(m, v2147, v2094, int32(4), v2308, v2309, v2118, v2126, int32(1))
	mBase = m.M
	v2312 = m.ExcPending
	if v2312 != 0 {
		goto L11
	} else {
		goto L523
	}
L523:
	;
	F_add_with_check_options(m, v2147, v2094, int32(2), v2308, v2309, v2118, v2126, int32(0))
	mBase = m.M
	v2316 = m.ExcPending
	if v2316 != 0 {
		goto L11
	} else {
		goto L524
	}
L524:
	;
	v2317 = int32(0)
	v2319 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2133)+16)))
	if v2319&int32(2) != 0 {
		goto L525
	} else {
		goto L526
	}
L525:
	;
	F_get_policies_for_relation(m, v2147, int32(1), v2140, v2111+int32(12), v2111+int32(8))
	mBase = m.M
	v2328 = m.ExcPending
	if v2328 != 0 {
		goto L11
	} else {
		goto L528
	}
L526:
	;
	v2335 = v2317
	v2336 = v2317
	goto L527
L527:
	;
	F_get_policies_for_relation(m, v2147, int32(4), v2140, v2111+int32(28), v2111+int32(24))
	mBase = m.M
	v2343 = m.ExcPending
	if v2343 != 0 {
		goto L11
	} else {
		goto L530
	}
L528:
	;
	v2330 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+12))
	v2331 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+8))
	F_add_with_check_options(m, v2147, v2094, int32(2), v2330, v2331, v2118, v2126, int32(1))
	mBase = m.M
	v2334 = m.ExcPending
	if v2334 != 0 {
		goto L11
	} else {
		goto L529
	}
L529:
	;
	v2335 = v2331
	v2336 = v2330
	goto L527
L530:
	;
	v2345 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+28))
	v2346 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+24))
	F_add_with_check_options(m, v2147, v2094, int32(5), v2345, v2346, v2118, v2126, int32(1))
	mBase = m.M
	v2349 = m.ExcPending
	if v2349 != 0 {
		goto L11
	} else {
		goto L531
	}
L531:
	;
	F_get_policies_for_relation(m, v2147, int32(3), v2140, v2111+int32(20), v2111+int32(16))
	mBase = m.M
	v2356 = m.ExcPending
	if v2356 != 0 {
		goto L11
	} else {
		goto L532
	}
L532:
	;
	v2358 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+20))
	v2359 = *(*int32)(unsafe.Add(mBase, uint32(v2111)+16))
	F_add_with_check_options(m, v2147, v2094, int32(1), v2358, v2359, v2118, v2126, int32(0))
	mBase = m.M
	v2362 = m.ExcPending
	if v2362 != 0 {
		goto L11
	} else {
		goto L533
	}
L533:
	;
	v2363 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v2133)+16)))
	if v2363&int32(2) == int32(0) {
		goto L495
	} else {
		goto L534
	}
L534:
	;
	v2368 = *(*int32)(unsafe.Add(mBase, uint32(l0)+96))
	if v2368 == int32(0) {
		goto L495
	} else {
		goto L535
	}
L535:
	;
	v2371 = int32(1)
	F_add_with_check_options(m, v2147, v2094, v2371, v2336, v2335, v2118, v2126, v2371)
	mBase = m.M
	v2374 = m.ExcPending
	if v2374 != 0 {
		goto L11
	} else {
		goto L536
	}
L536:
	;
	goto L495
L537:
	;
	v2382 = *(*int32)(unsafe.Add(mBase, uint32(v2114)))
	v2383 = *(*int32)(unsafe.Add(mBase, uint32(v2133)+24))
	F_setRuleCheckAsUser(m, v2382, v2383)
	mBase = m.M
	v2385 = m.ExcPending
	if v2385 != 0 {
		goto L11
	} else {
		goto L538
	}
L538:
	;
	v2386 = *(*int32)(unsafe.Add(mBase, uint32(v2118)))
	v2387 = *(*int32)(unsafe.Add(mBase, uint32(v2133)+24))
	F_setRuleCheckAsUser(m, v2386, v2387)
	mBase = m.M
	v2389 = m.ExcPending
	if v2389 != 0 {
		goto L11
	} else {
		goto L539
	}
L539:
	;
	goto L468
L540:
	;
	v2410 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+62)))
	if v2410 == int32(1) {
		goto L543
	} else {
		goto L544
	}
L541:
	;
	v2514 = v2069
	goto L542
L542:
	;
	v2516 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+63)))
	if v2516 == int32(1) {
		goto L582
	} else {
		goto L583
	}
L543:
	;
	v2413 = *(*int32)(unsafe.Add(mBase, uint32(v2106)+56))
	v2414 = int32(0)
	if v2069 == v2414 {
		goto L547
	} else {
		goto L548
	}
L544:
	;
	v2502 = v2069
	v2504 = v2407
	goto L545
L545:
	;
	v2505 = *(*int32)(unsafe.Add(mBase, uint32(v2099)+128))
	v2506 = F_list_concat(m, v2504, v2505)
	mBase = m.M
	v2507 = m.ExcPending
	if v2507 != 0 {
		goto L11
	} else {
		goto L580
	}
L546:
	;
	if v2452 != 0 {
		goto L305
	} else {
		goto L559
	}
L547:
	;
	v2452 = int32(0)
	goto L546
L548:
	;
	goto L549
L549:
	;
	v2420 = *(*int32)(unsafe.Add(mBase, uint32(v2069)+4))
	if v2420 <= int32(0) {
		v2446 = v2414
		goto L550
	} else {
		goto L551
	}
L550:
	;
	v2452 = v2446
	goto L546
L551:
	;
	v2423 = int32(0)
	if v2423 < v2420 {
		goto L552
	} else {
		goto L553
	}
L552:
	;
	v2426 = v2420
	goto L554
L553:
	;
	v2426 = v2423
	goto L554
L554:
	;
	v2427 = *(*int32)(unsafe.Add(mBase, uint32(v2069)+12))
	v2429 = int32(0)
	goto L555
L555:
	;
	v2437 = *(*int32)(unsafe.Add(mBase, uint32(v2427+v2429<<(uint(int32(2))%32))))
	v2438 = base.B2i32(v2437 == v2413)
	if v2437 == v2413 {
		v2446 = v2438
		goto L550
	} else {
		goto L557
	}
L556:
	;
	v2446 = v2438
	goto L550
L557:
	;
	v2440 = v2429 + int32(1)
	if v2440 != v2426 {
		v2429 = v2440
		goto L555
	} else {
		goto L558
	}
L558:
	;
	goto L556
L559:
	;
	v2453 = *(*int32)(unsafe.Add(mBase, uint32(v2106)+56))
	v2454 = F_lappend_oid(m, v2069, v2453)
	mBase = m.M
	v2455 = m.ExcPending
	if v2455 != 0 {
		goto L11
	} else {
		goto L560
	}
L560:
	;
	v2456 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+61)) = uint8(v2456)
	v2458 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
	if v2458 != 0 {
		goto L561
	} else {
		goto L562
	}
L561:
	;
	v2459 = *(*int32)(unsafe.Add(mBase, uint32(v2458)))
	if v2459 == int32(22) {
		goto L564
	} else {
		goto L565
	}
L562:
	;
	goto L563
L563:
	;
	v2472 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	if v2472 != 0 {
		goto L569
	} else {
		goto L570
	}
L564:
	;
	v2462 = *(*int32)(unsafe.Add(mBase, uint32(v2458)+20))
	F_AcquireRewriteLocks(m, v2462, int32(1), int32(0))
	mBase = m.M
	v2466 = m.ExcPending
	if v2466 != 0 {
		goto L11
	} else {
		goto L567
	}
L565:
	;
	goto L566
L566:
	;
	v2470 = F_expression_tree_walker_impl(m, v2458, int32(1041), v28+int32(61))
	mBase = m.M
	v2471 = m.ExcPending
	if v2471 != 0 {
		goto L11
	} else {
		goto L568
	}
L567:
	;
	goto L566
L568:
	;
	goto L563
L569:
	;
	v2473 = *(*int32)(unsafe.Add(mBase, uint32(v2472)))
	if v2473 == int32(22) {
		goto L572
	} else {
		goto L573
	}
L570:
	;
	goto L571
L571:
	;
	v2486 = int32(0)
	*(*uint8)(unsafe.Add(mBase, uint32(v28)+76)) = uint8(v2486)
	*(*int32)(unsafe.Add(mBase, uint32(v28)+72)) = v2454
	v2489 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
	v2492 = v28 + int32(72)
	v2493 = F_expression_tree_walker_impl(m, v2489, int32(1042), v2492)
	mBase = m.M
	v2494 = m.ExcPending
	if v2494 != 0 {
		goto L11
	} else {
		goto L577
	}
L572:
	;
	v2476 = *(*int32)(unsafe.Add(mBase, uint32(v2472)+20))
	v2477 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+61)))
	F_AcquireRewriteLocks(m, v2476, v2477, int32(0))
	mBase = m.M
	v2480 = m.ExcPending
	if v2480 != 0 {
		goto L11
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v2484 = F_expression_tree_walker_impl(m, v2472, int32(1041), v28+int32(61))
	mBase = m.M
	v2485 = m.ExcPending
	if v2485 != 0 {
		goto L11
	} else {
		goto L576
	}
L575:
	;
	goto L574
L576:
	;
	goto L571
L577:
	;
	v2495 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	v2497 = F_expression_tree_walker_impl(m, v2495, int32(1042), v2492)
	mBase = m.M
	v2498 = m.ExcPending
	if v2498 != 0 {
		goto L11
	} else {
		goto L578
	}
L578:
	;
	v2499 = F_list_delete_last(m, v2454)
	mBase = m.M
	v2500 = m.ExcPending
	if v2500 != 0 {
		goto L11
	} else {
		goto L579
	}
L579:
	;
	v2501 = *(*int32)(unsafe.Add(mBase, uint32(v28)+68))
	v2502 = v2499
	v2504 = v2501
	goto L545
L580:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v2099)+128)) = v2506
	v2509 = *(*int32)(unsafe.Add(mBase, uint32(v28)+64))
	v2510 = *(*int32)(unsafe.Add(mBase, uint32(l0)+152))
	v2511 = F_list_concat(m, v2509, v2510)
	mBase = m.M
	v2512 = m.ExcPending
	if v2512 != 0 {
		goto L11
	} else {
		goto L581
	}
L581:
	;
	*(*int32)(unsafe.Add(mBase, uint32(l0)+152)) = v2511
	v2514 = v2502
	goto L542
L582:
	;
	v2519 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+44)) = uint8(v2519)
	goto L584
L583:
	;
	goto L584
L584:
	;
	v2521 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v28)+62)))
	if v2521 == int32(1) {
		goto L585
	} else {
		goto L586
	}
L585:
	;
	v2524 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+39)) = uint8(v2524)
	goto L587
L586:
	;
	goto L587
L587:
	;
	F_relation_close(m, v2106, int32(0))
	mBase = m.M
	v2528 = m.ExcPending
	if v2528 != 0 {
		goto L11
	} else {
		goto L588
	}
L588:
	;
	v2529 = v2514
	goto L461
L589:
	;
	goto L460
L590:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v2580 = m.ExcPending
	if v2580 != 0 {
		goto L11
	} else {
		goto L591
	}
L591:
	;
	v2581 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+16)) = v2581 + int32(4)
	F_errmsg(m, int32(_a_F_fireRIRrules_9), v28+int32(16))
	mBase = m.M
	v2589 = m.ExcPending
	if v2589 != 0 {
		goto L11
	} else {
		goto L592
	}
L592:
	;
	F_errfinish(m, int32(_a_F_fireRIRrules_10), int32(2134), int32(_a_F_fireRIRrules_11))
	mBase = m.M
	v2594 = m.ExcPending
	if v2594 != 0 {
		goto L11
	} else {
		goto L593
	}
L593:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L594:
	;
	F_errmsg_internal(m, int32(_a_F_fireRIRrules_12), int32(0))
	mBase = m.M
	v2602 = m.ExcPending
	if v2602 != 0 {
		goto L11
	} else {
		goto L595
	}
L595:
	;
	F_errfinish(m, int32(_a_F_fireRIRrules_10), int32(1725), int32(_a_F_fireRIRrules_13))
	mBase = m.M
	v2607 = m.ExcPending
	if v2607 != 0 {
		goto L11
	} else {
		goto L596
	}
L596:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L597:
	;
	F_errmsg_internal(m, int32(_a_F_fireRIRrules_14), int32(0))
	mBase = m.M
	v2615 = m.ExcPending
	if v2615 != 0 {
		goto L11
	} else {
		goto L598
	}
L598:
	;
	F_errfinish(m, int32(_a_F_fireRIRrules_10), int32(1727), int32(_a_F_fireRIRrules_13))
	mBase = m.M
	v2620 = m.ExcPending
	if v2620 != 0 {
		goto L11
	} else {
		goto L599
	}
L599:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L600:
	;
	F_errcode(m, int32(325))
	mBase = m.M
	v2627 = m.ExcPending
	if v2627 != 0 {
		goto L11
	} else {
		goto L601
	}
L601:
	;
	v2628 = *(*int32)(unsafe.Add(mBase, uint32(v1438)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+48)) = v2628 + int32(4)
	F_errmsg(m, int32(_a_F_fireRIRrules_15), v28+int32(48))
	mBase = m.M
	v2636 = m.ExcPending
	if v2636 != 0 {
		goto L11
	} else {
		goto L602
	}
L602:
	;
	F_errfinish(m, int32(_a_F_fireRIRrules_10), int32(1735), int32(_a_F_fireRIRrules_13))
	mBase = m.M
	v2641 = m.ExcPending
	if v2641 != 0 {
		goto L11
	} else {
		goto L603
	}
L603:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L604:
	;
	v2646 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v28)+32)) = v2646
	F_errmsg_internal(m, int32(_a_F_fireRIRrules_16), v28+int32(32))
	mBase = m.M
	v2652 = m.ExcPending
	if v2652 != 0 {
		goto L11
	} else {
		goto L605
	}
L605:
	;
	F_errfinish(m, int32(_a_F_fireRIRrules_10), int32(1803), int32(_a_F_fireRIRrules_13))
	mBase = m.M
	v2657 = m.ExcPending
	if v2657 != 0 {
		goto L11
	} else {
		goto L606
	}
L606:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L607:
	;
	F_errcode(m, int32(117833860))
	mBase = m.M
	v2664 = m.ExcPending
	if v2664 != 0 {
		goto L11
	} else {
		goto L608
	}
L608:
	;
	v2665 = *(*int32)(unsafe.Add(mBase, uint32(v2106)+48))
	*(*int32)(unsafe.Add(mBase, uint32(v28))) = v2665 + int32(4)
	F_errmsg(m, int32(_a_F_fireRIRrules_17), v28)
	mBase = m.M
	v2671 = m.ExcPending
	if v2671 != 0 {
		goto L11
	} else {
		goto L609
	}
L609:
	;
	F_errfinish(m, int32(_a_F_fireRIRrules_10), int32(2239), int32(_a_F_fireRIRrules_11))
	mBase = m.M
	v2676 = m.ExcPending
	if v2676 != 0 {
		goto L11
	} else {
		goto L610
	}
L610:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_fix_opfuncids_walker(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
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
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	if l0 == int32(0) {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		switch v7 - int32(17) {
		case 0:
			v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v10 != 0 {
				v20 = F_expression_tree_walker_impl_x2especialized_x2e1(m, l0, l1)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = F_get_opcode(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24
					v27 = F_expression_tree_walker_impl_x2especialized_x2e1(m, l0, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						return v27
					}
				}
			}
		case 1:
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v11 != 0 {
				v20 = F_expression_tree_walker_impl_x2especialized_x2e1(m, l0, l1)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = F_get_opcode(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24
					v27 = F_expression_tree_walker_impl_x2especialized_x2e1(m, l0, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						return v27
					}
				}
			}
		case 2:
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v12 != 0 {
				v20 = F_expression_tree_walker_impl_x2especialized_x2e1(m, l0, l1)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v24 = F_get_opcode(m, v23)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v24
					v27 = F_expression_tree_walker_impl_x2especialized_x2e1(m, l0, l1)
					mBase = m.M
					v28 = m.ExcPending
					if v28 != 0 {
						return int32(0)
					} else {
						return v27
					}
				}
			}
		case 3:
			v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
			if v13 != 0 {
				v20 = F_expression_tree_walker_impl_x2especialized_x2e1(m, l0, l1)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					return v20
				}
			} else {
				v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v15 = F_get_opcode(m, v14)
				mBase = m.M
				v18 = m.ExcPending
				if v18 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+8)) = v15
					v20 = F_expression_tree_walker_impl_x2especialized_x2e1(m, l0, l1)
					mBase = m.M
					v21 = m.ExcPending
					if v21 != 0 {
						return int32(0)
					} else {
						return v20
					}
				}
			}
		default:
			v20 = F_expression_tree_walker_impl_x2especialized_x2e1(m, l0, l1)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				return v20
			}
		}
	}
}
func F_float48eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 float32
	_ = v10
	var v11 float64
	_ = v11
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v8 = int64(9223372036854775807)
	v9 = base.I64_reinterpret_f64(v6) & v8
	v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = base.F64_promote_f32(v10)
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v11)&v8) {
		return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v9))
	} else {
		return base.B2i32(base.Ui64(v9) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v6, v11)
	}
}
func F_float48ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v5 float64
	_ = v5
	var v11 int32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = base.F64_promote_f32(v4)
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = *(*float64)(unsafe.Add(mBase, uint32(v11)))
		v22 = base.F64_ge(v5, v12) & base.B2i32(base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)) < base.Ui64(int64(9218868437227405313)))
	} else {
		v22 = int32(1)
	}
	return v22
}
func F_float48mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 float32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v18 int32
	_ = v18
	var v28 float64
	_ = v28
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = base.F64_promote_f32(v8)
	v10 = base.F64_mul(v7, v9)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	v18 = int32(0)
	if base.B2i32(base.F64_ne(base.F64_abs(v10), v12)|base.F64_eq(base.F64_abs(v9), v12) == v18)&base.F64_ne(base.F64_abs(v7), v12) == v18 {
		v28 = float64(0)
		if base.B2i32(base.F32_eq(v8, float32(0))|base.F64_ne(v10, v28) == int32(0))&base.F64_ne(v7, v28) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v36 = F_Float8GetDatum(m, v10)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				return v36
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_float48pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = base.F64_promote_f32(v7)
	v9 = base.F64_add(v6, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v8), v11)|base.F64_eq(base.F64_abs(v6), v11) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v27 = F_Float8GetDatum(m, v9)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_float4_numeric(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 float32
	_ = v20
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int64
	_ = v44
	var v51 int32
	_ = v51
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
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	v5 = m.G0
	v7 = v5 - int32(176)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(int32(2139095041)) <= base.Ui32(v9&int32(2147483647)) {
		v16 = F_make_result_opt_error(m, int32(_a_F_float4_numeric_0), int32(0))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v65 = v16
			m.G0 = v7 + int32(176)
			return v65
		}
	} else {
		v20 = base.F32_reinterpret_i32(v9)
		if base.F32_eq(base.F32_abs(v20), math.Float32frombits(uint32(0x7f800000))) != 0 {
			if base.F32_lt(v20, float32(0)) != 0 {
				v28 = F_make_result_opt_error(m, int32(_a_F_float4_numeric_1), int32(0))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					v65 = v28
					m.G0 = v7 + int32(176)
					return v65
				}
			} else {
				v32 = F_make_result_opt_error(m, int32(_a_F_float4_numeric_2), int32(0))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					v65 = v32
					m.G0 = v7 + int32(176)
					return v65
				}
			}
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7))) = int32(6)
			*(*float64)(unsafe.Add(mBase, uint32(v7)+8)) = base.F64_promote_f32(v20)
			v39 = v7 + int32(32)
			v42 = F_pg_snprintf(m, v39, int32(106), int32(_a_F_float4_numeric_3), v7)
			mBase = m.M
			v43 = m.ExcPending
			if v43 != 0 {
				return int32(0)
			} else {
				v44 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v7)+168)) = v44
				*(*int64)(unsafe.Add(mBase, uint32(v7)+160)) = v44
				*(*int64)(unsafe.Add(mBase, uint32(v7)+152)) = v44
				v51 = v7 + int32(152)
				v55 = F_set_var_from_str(m, v39, v39, v51, v7+int32(28), int32(0))
				mBase = m.M
				v56 = m.ExcPending
				if v56 != 0 {
					return int32(0)
				} else {
					v58 = F_make_result_opt_error(m, v51, int32(0))
					mBase = m.M
					v59 = m.ExcPending
					if v59 != 0 {
						return int32(0)
					} else {
						v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+168))
						if v60 == int32(0) {
							v65 = v58
							m.G0 = v7 + int32(176)
							return v65
						} else {
							F_pfree(m, v60)
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return int32(0)
							} else {
								v65 = v58
								m.G0 = v7 + int32(176)
								return v65
							}
						}
					}
				}
			}
		}
	}
}
func F_float4abs(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2 & int32(2147483647)
}
func F_float4ge(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v10 float32
	_ = v10
	var v20 int32
	_ = v20
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(base.I32_reinterpret_f32(v4)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = base.F32_ge(v4, v10) & base.B2i32(base.Ui32(base.I32_reinterpret_f32(v10)&int32(2147483647)) < base.Ui32(int32(2139095041)))
	} else {
		v20 = int32(1)
	}
	return v20
}
func F_float4lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 float32
	_ = v4
	var v10 float32
	_ = v10
	var v20 int32
	_ = v20
	v4 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	if base.Ui32(base.I32_reinterpret_f32(v4)&int32(2147483647)) <= base.Ui32(int32(2139095040)) {
		v10 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
		v20 = base.F32_lt(v4, v10) | base.B2i32(base.Ui32(int32(2139095040)) < base.Ui32(base.I32_reinterpret_f32(v10)&int32(2147483647)))
	} else {
		v20 = int32(0)
	}
	return v20
}
func F_float4mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float32
	_ = v6
	var v7 float32
	_ = v7
	var v9 float32
	_ = v9
	var v15 int32
	_ = v15
	var v23 float32
	_ = v23
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = base.F32_mul(v5, v6)
	v9 = math.Float32frombits(uint32(0x7f800000))
	v15 = int32(0)
	if base.B2i32(base.F32_ne(base.F32_abs(v7), v9)|base.F32_eq(base.F32_abs(v5), v9) == v15)&base.F32_ne(base.F32_abs(v6), v9) == v15 {
		v23 = float32(0)
		if base.B2i32(base.F32_eq(v5, v23)|base.F32_ne(v7, v23) == int32(0))&base.F32_ne(v6, v23) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v40 = m.ExcPending
			if v40 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			return base.I32_reinterpret_f32(v7)
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_float4pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float32
	_ = v6
	var v7 float32
	_ = v7
	var v9 float32
	_ = v9
	var v24 int32
	_ = v24
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v7 = base.F32_add(v5, v6)
	v9 = math.Float32frombits(uint32(0x7f800000))
	if base.F32_ne(base.F32_abs(v7), v9)|base.F32_eq(base.F32_abs(v5), v9)|base.F32_eq(base.F32_abs(v6), v9) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		return base.I32_reinterpret_f32(v7)
	}
}
func F_float4send(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 float32
	_ = v8
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l0)+20))
	F_pq_begintypsend(m, v6)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		F_pq_sendfloat4(m, v6, v8)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int32(0)
		} else {
			v16 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
			v17 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
			*(*int32)(unsafe.Add(mBase, uint32(v16))) = v17 << (uint(int32(2)) % 32)
			m.G0 = v6 + int32(16)
			return v16
		}
	}
}
func F_float4um(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	return v2 ^ int32(-2147483648)
}
func F_float84eq(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 float32
	_ = v5
	var v6 float64
	_ = v6
	var v8 int64
	_ = v8
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 float64
	_ = v11
	v5 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v6 = base.F64_promote_f32(v5)
	v8 = int64(9223372036854775807)
	v9 = base.I64_reinterpret_f64(v6) & v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*float64)(unsafe.Add(mBase, uint32(v10)))
	if base.Ui64(int64(9218868437227405313)) <= base.Ui64(base.I64_reinterpret_f64(v11)&v8) {
		return base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(v9))
	} else {
		return base.B2i32(base.Ui64(v9) < base.Ui64(int64(9218868437227405313))) & base.F64_eq(v6, v11)
	}
}
func F_float84lt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v11 float32
	_ = v11
	var v12 float64
	_ = v12
	var v22 int32
	_ = v22
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	if base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		v11 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
		v12 = base.F64_promote_f32(v11)
		v22 = base.F64_lt(v5, v12) | base.B2i32(base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v12)&int64(9223372036854775807)))
	} else {
		v22 = int32(0)
	}
	return v22
}
func F_float84mul(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v8 float32
	_ = v8
	var v9 float64
	_ = v9
	var v10 float64
	_ = v10
	var v12 float64
	_ = v12
	var v18 int32
	_ = v18
	var v26 float64
	_ = v26
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	v8 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v9 = base.F64_promote_f32(v8)
	v10 = base.F64_mul(v7, v9)
	v12 = math.Float64frombits(uint64(0x7ff0000000000000))
	v18 = int32(0)
	if base.B2i32(base.F64_ne(base.F64_abs(v10), v12)|base.F64_eq(base.F64_abs(v7), v12) == v18)&base.F64_ne(base.F64_abs(v9), v12) == v18 {
		v26 = float64(0)
		if base.B2i32(base.F64_eq(v7, v26)|base.F64_ne(v10, v26) == int32(0))&base.F32_ne(v8, float32(0)) != 0 {
			F_float_underflow_error(m)
			mBase = m.M
			v44 = m.ExcPending
			if v44 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		} else {
			v36 = F_Float8GetDatum(m, v10)
			mBase = m.M
			v39 = m.ExcPending
			if v39 != 0 {
				return int32(0)
			} else {
				return v36
			}
		}
	} else {
		F_float_overflow_error(m)
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_float84pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 float32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*float32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = base.F64_promote_f32(v7)
	v9 = base.F64_add(v6, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v6), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v27 = F_Float8GetDatum(m, v9)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_float8mi(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v9 = base.F64_sub(v6, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v6), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v27 = F_Float8GetDatum(m, v9)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_float8pl(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v6 float64
	_ = v6
	var v7 int32
	_ = v7
	var v8 float64
	_ = v8
	var v9 float64
	_ = v9
	var v11 float64
	_ = v11
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*float64)(unsafe.Add(mBase, uint32(v5)))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v8 = *(*float64)(unsafe.Add(mBase, uint32(v7)))
	v9 = base.F64_add(v6, v8)
	v11 = math.Float64frombits(uint64(0x7ff0000000000000))
	if base.F64_ne(base.F64_abs(v9), v11)|base.F64_eq(base.F64_abs(v6), v11)|base.F64_eq(base.F64_abs(v8), v11) == int32(0) {
		F_float_overflow_error(m)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int32(0)
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	} else {
		v27 = F_Float8GetDatum(m, v9)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			return v27
		}
	}
}
func F_float8smaller(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 float64
	_ = v5
	var v6 int32
	_ = v6
	var v7 float64
	_ = v7
	var v18 float64
	_ = v18
	var v20 float64
	_ = v20
	var v21 float64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v5 = *(*float64)(unsafe.Add(mBase, uint32(v4)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*float64)(unsafe.Add(mBase, uint32(v6)))
	if base.Ui64(base.I64_reinterpret_f64(v7)&int64(9223372036854775807)) <= base.Ui64(int64(9218868437227405312)) {
		if base.Ui64(int64(9218868437227405312)) < base.Ui64(base.I64_reinterpret_f64(v5)&int64(9223372036854775807)) {
			v18 = v7
		} else {
			v18 = v5
		}
		if base.F64_gt(v5, v7) != 0 {
			v20 = v7
		} else {
			v20 = v18
		}
		v21 = v20
	} else {
		v21 = v5
	}
	v22 = F_Float8GetDatum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		return v22
	}
}
func F_float8um(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 float64
	_ = v3
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*float64)(unsafe.Add(mBase, uint32(v2)))
	v5 = F_Float8GetDatum(m, base.F64_neg(v3))
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		return v5
	}
}
func F_fp_barrier_1(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	v3 = m.G0
	*(*float64)(unsafe.Add(mBase, uint32(v3-int32(16))+8)) = l0
	return l0
}
func F_free_attrmap(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	F_pfree(m, v2)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		F_pfree(m, l0)
		mBase = m.M
		v6 = m.ExcPending
		if v6 != 0 {
			return
		} else {
			return
		}
	}
}
func F_free_auth_file(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	v2 = F_FreeFile(m, l0)
	mBase = m.M
	v3 = m.ExcPending
	if v3 != 0 {
		return
	} else {
		v5 = *(*int32)(unsafe.Add(mBase, _c_F_free_auth_file[0]))
		F_MemoryContextDelete(m, v5)
		mBase = m.M
		v7 = m.ExcPending
		if v7 != 0 {
			return
		} else {
			*(*int32)(unsafe.Add(mBase, _c_F_free_auth_file[0])) = int32(0)
			return
		}
	}
}
func F_freelocale(m *base.Module, l0 int32) {
	if base.B2i32(l0 != int32(0))&base.B2i32(l0 != int32(_a_F_freelocale_0))&base.B2i32(l0 != int32(_a_F_freelocale_1))&base.B2i32(l0 != int32(_a_F_freelocale_2))&base.B2i32(l0 != int32(_a_F_freelocale_3)) != 0 {
		F_emscripten_builtin_free(m, l0)
	} else {
	}
	return
}
func F_freetree(m *base.Module, l0 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	F_check_stack_depth(m)
	mBase = m.M
	v4 = m.ExcPending
	if v4 != 0 {
		return
	} else {
		if l0 != 0 {
			v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			if v5 != 0 {
				F_freetree(m, v5)
				mBase = m.M
				v7 = m.ExcPending
				if v7 != 0 {
					return
				} else {
					v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v8 != 0 {
						F_freetree(m, v8)
						mBase = m.M
						v10 = m.ExcPending
						if v10 != 0 {
							return
						} else {
							F_pfree(m, l0)
							mBase = m.M
							v12 = m.ExcPending
							if v12 != 0 {
								return
							} else {
								return
							}
						}
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v12 = m.ExcPending
						if v12 != 0 {
							return
						} else {
							return
						}
					}
				}
			} else {
				v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				if v8 != 0 {
					F_freetree(m, v8)
					mBase = m.M
					v10 = m.ExcPending
					if v10 != 0 {
						return
					} else {
						F_pfree(m, l0)
						mBase = m.M
						v12 = m.ExcPending
						if v12 != 0 {
							return
						} else {
							return
						}
					}
				} else {
					F_pfree(m, l0)
					mBase = m.M
					v12 = m.ExcPending
					if v12 != 0 {
						return
					} else {
						return
					}
				}
			}
		} else {
			return
		}
	}
}
func F_funcname_signature_string(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v18 int32
	_ = v18
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
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	v9 = m.G0
	v11 = v9 + int32(-64)
	m.G0 = v11
	v14 = v9 + int32(-16)
	F_initStringInfo(m, v14)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v11)+32)) = l0
	F_appendStringInfo(m, v14, int32(_a_F_funcname_signature_string_0), v9+int32(-32))
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	if l2 != 0 {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v26 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	v28 = v25
	v29 = l1 - v26
	goto L6
L5:
	;
	v28 = int32(0)
	v29 = l1
	goto L6
L6:
	;
	if l1 <= int32(0) {
		goto L7
	} else {
		goto L8
	}
L7:
	;
	F_appendStringInfoChar(m, v9+int32(-16), int32(41))
	mBase = m.M
	v120 = m.ExcPending
	if v120 != 0 {
		goto L1
	} else {
		goto L32
	}
L8:
	;
	if v29 <= int32(0) {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v34 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v34
	F_appendStringInfo(m, v9+int32(-16), int32(_a_F_funcname_signature_string_1), v9+int32(-48))
	mBase = m.M
	v42 = m.ExcPending
	if v42 != 0 {
		goto L1
	} else {
		goto L12
	}
L10:
	;
	v54 = v28
	goto L11
L11:
	;
	v57 = *(*int32)(unsafe.Add(mBase, uint32(l3)))
	v58 = F_format_type_be(m, v57)
	mBase = m.M
	v59 = m.ExcPending
	if v59 != 0 {
		goto L1
	} else {
		goto L16
	}
L12:
	;
	v44 = v28 + int32(4)
	v46 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v47 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v44) < base.Ui32(v46+v47<<(uint(int32(2))%32)) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v52 = v44
	goto L15
L14:
	;
	v52 = int32(0)
	goto L15
L15:
	;
	v54 = v52
	goto L11
L16:
	;
	F_appendStringInfoString(m, v9+int32(-16), v58)
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L17:
	;
	v62 = int32(1)
	if l1 == v62 {
		goto L7
	} else {
		goto L18
	}
L18:
	;
	v65 = v62
	v70 = v54
	goto L19
L19:
	;
	v74 = v9 + int32(-16)
	F_appendStringInfoString(m, v74, int32(_a_F_funcname_signature_string_2))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L21
	}
L20:
	;
	goto L7
L21:
	;
	if v29 <= v65 {
		goto L22
	} else {
		goto L23
	}
L22:
	;
	v79 = *(*int32)(unsafe.Add(mBase, uint32(v70)))
	*(*int32)(unsafe.Add(mBase, uint32(v11))) = v79
	F_appendStringInfo(m, v74, int32(_a_F_funcname_signature_string_1), v11)
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L25
	}
L23:
	;
	v94 = v70
	goto L24
L24:
	;
	v100 = *(*int32)(unsafe.Add(mBase, uint32(l3+v65<<(uint(int32(2))%32))))
	v101 = F_format_type_be(m, v100)
	mBase = m.M
	v102 = m.ExcPending
	if v102 != 0 {
		goto L1
	} else {
		goto L29
	}
L25:
	;
	v85 = v70 + int32(4)
	v87 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
	v88 = *(*int32)(unsafe.Add(mBase, uint32(l2)+4))
	if base.Ui32(v85) < base.Ui32(v87+v88<<(uint(int32(2))%32)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v93 = v85
	goto L28
L27:
	;
	v93 = int32(0)
	goto L28
L28:
	;
	v94 = v93
	goto L24
L29:
	;
	F_appendStringInfoString(m, v9+int32(-16), v101)
	mBase = m.M
	v104 = m.ExcPending
	if v104 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v106 = v65 + int32(1)
	if v106 != l1 {
		v65 = v106
		v70 = v94
		goto L19
	} else {
		goto L31
	}
L31:
	;
	goto L20
L32:
	;
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v11)+48))
	m.G0 = v11 - int32(-64)
	return v121
}
