package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cclasscvec(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
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
	var v48 int32
	_ = v48
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v68 int32
	_ = v68
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	var v110 int32
	_ = v110
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v141 int32
	_ = v141
	var v145 int32
	_ = v145
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v167 int32
	_ = v167
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v200 int32
	_ = v200
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v230 int32
	_ = v230
	var v231 int32
	_ = v231
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v249 int32
	_ = v249
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v270 int32
	_ = v270
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
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v288 int32
	_ = v288
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v305 int32
	_ = v305
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v318 int32
	_ = v318
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v359 int32
	_ = v359
	var v361 int32
	_ = v361
	var v364 int32
	_ = v364
	v4 = int32(1)
	if l1 == int32(11) {
		v8 = v4
	} else {
		v8 = l1
	}
	if l1 == int32(7) {
		v11 = v4
	} else {
		v11 = v8
	}
	if l2 != 0 {
		v12 = v11
	} else {
		v12 = l1
	}
	switch v12 {
	case 0:
		v23 = F_pg_ctype_get_cache(m, int32(972), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v23 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v23
			}
			return v364
		}
	case 1:
		v29 = F_pg_ctype_get_cache(m, int32(973), int32(1))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			if v29 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v29
			}
			return v364
		}
	case 2:
		v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v39 != 0 {
			v40 = *(*int32)(unsafe.Add(mBase, uint32(v39)+4))
			if v40 < int32(0) {
				F_pfree(m, v39)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v57 = F_palloc_extended(m, int32(36), int32(2))
					mBase = m.M
					v58 = m.ExcPending
					if v58 != 0 {
						return int32(0)
					} else {
						if v57 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
							v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v352 != 0 {
								v354 = v352
							} else {
								v354 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v359 != 0 {
								v361 = v359
							} else {
								v361 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
							v364 = int32(0)
							return v364
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v57)+12)) = int64(4294967296)
							v68 = v57 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v68
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v57
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
							v75 = v57
							v77 = v72 << (uint(int32(3)) % 32)
							v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v77+v78))) = int32(0)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(3))%32))+4)) = int32(127)
							v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89 + int32(1)
							return v75
						}
					}
				}
			} else {
				v43 = *(*int32)(unsafe.Add(mBase, uint32(v39)+16))
				if v43 <= int32(0) {
					F_pfree(m, v39)
					mBase = m.M
					v54 = m.ExcPending
					if v54 != 0 {
						return int32(0)
					} else {
						v57 = F_palloc_extended(m, int32(36), int32(2))
						mBase = m.M
						v58 = m.ExcPending
						if v58 != 0 {
							return int32(0)
						} else {
							if v57 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v352 != 0 {
									v354 = v352
								} else {
									v354 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v359 != 0 {
									v361 = v359
								} else {
									v361 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
								v364 = int32(0)
								return v364
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v57)+12)) = int64(4294967296)
								v68 = v57 + int32(28)
								*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v68
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v57
								v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
								v75 = v57
								v77 = v72 << (uint(int32(3)) % 32)
								v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v77+v78))) = int32(0)
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(3))%32))+4)) = int32(127)
								v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89 + int32(1)
								return v75
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v39)+24)) = int32(-1)
					v48 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v39)+12)) = v48
					*(*int32)(unsafe.Add(mBase, uint32(v39))) = v48
					v75 = v39
					v77 = v48
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v77+v78))) = int32(0)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(3))%32))+4)) = int32(127)
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89 + int32(1)
					return v75
				}
			}
		} else {
			v57 = F_palloc_extended(m, int32(36), int32(2))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				if v57 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
					v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v352 != 0 {
						v354 = v352
					} else {
						v354 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
					v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v359 != 0 {
						v361 = v359
					} else {
						v361 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
					v364 = int32(0)
					return v364
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v57))) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v57)+24)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v57)+12)) = int64(4294967296)
					v68 = v57 + int32(28)
					*(*int32)(unsafe.Add(mBase, uint32(v57)+20)) = v68
					*(*int32)(unsafe.Add(mBase, uint32(v57)+8)) = v68
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v57
					v72 = *(*int32)(unsafe.Add(mBase, uint32(v57)+12))
					v75 = v57
					v77 = v72 << (uint(int32(3)) % 32)
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v77+v78))) = int32(0)
					v82 = *(*int32)(unsafe.Add(mBase, uint32(v75)+20))
					v83 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v82+v83<<(uint(int32(3))%32))+4)) = int32(127)
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v75)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v75)+12)) = v89 + int32(1)
					return v75
				}
			}
		}
	case 3:
		v94 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v94 != 0 {
			v95 = *(*int32)(unsafe.Add(mBase, uint32(v94)+4))
			if v95 < int32(2) {
				F_pfree(m, v94)
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					v114 = F_palloc_extended(m, int32(36), int32(2))
					mBase = m.M
					v115 = m.ExcPending
					if v115 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v114)+12)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v114))) = int64(8589934592)
						*(*int32)(unsafe.Add(mBase, uint32(v114)+20)) = v114 + int32(36)
						*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v114 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v114
						v129 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
						v130 = v114
						v131 = v129
						v132 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 + v132
						v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
						v136 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(v136)%32)))) = int32(9)
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
						*(*int32)(unsafe.Add(mBase, uint32(v130))) = v141 + v132
						v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v145+v141<<(uint(v136)%32)))) = int32(32)
						return v130
					}
				}
			} else {
				v98 = int32(0)
				v99 = *(*int32)(unsafe.Add(mBase, uint32(v94)+16))
				if v99 < v98 {
					F_pfree(m, v94)
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						v114 = F_palloc_extended(m, int32(36), int32(2))
						mBase = m.M
						v115 = m.ExcPending
						if v115 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v114)+12)) = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v114))) = int64(8589934592)
							*(*int32)(unsafe.Add(mBase, uint32(v114)+20)) = v114 + int32(36)
							*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v114 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v114
							v129 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
							v130 = v114
							v131 = v129
							v132 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 + v132
							v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
							v136 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(v136)%32)))) = int32(9)
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
							*(*int32)(unsafe.Add(mBase, uint32(v130))) = v141 + v132
							v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v145+v141<<(uint(v136)%32)))) = int32(32)
							return v130
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v94)+24)) = int32(-1)
					v104 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v94)+12)) = v104
					*(*int32)(unsafe.Add(mBase, uint32(v94))) = v104
					v130 = v94
					v131 = v98
					v132 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 + v132
					v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
					v136 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(v136)%32)))) = int32(9)
					v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
					*(*int32)(unsafe.Add(mBase, uint32(v130))) = v141 + v132
					v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v145+v141<<(uint(v136)%32)))) = int32(32)
					return v130
				}
			}
		} else {
			v114 = F_palloc_extended(m, int32(36), int32(2))
			mBase = m.M
			v115 = m.ExcPending
			if v115 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v114)+24)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v114)+12)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v114))) = int64(8589934592)
				*(*int32)(unsafe.Add(mBase, uint32(v114)+20)) = v114 + int32(36)
				*(*int32)(unsafe.Add(mBase, uint32(v114)+8)) = v114 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v114
				v129 = *(*int32)(unsafe.Add(mBase, uint32(v114)))
				v130 = v114
				v131 = v129
				v132 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v130))) = v131 + v132
				v135 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
				v136 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v135+v131<<(uint(v136)%32)))) = int32(9)
				v141 = *(*int32)(unsafe.Add(mBase, uint32(v130)))
				*(*int32)(unsafe.Add(mBase, uint32(v130))) = v141 + v132
				v145 = *(*int32)(unsafe.Add(mBase, uint32(v130)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v145+v141<<(uint(v136)%32)))) = int32(32)
				return v130
			}
		}
	case 4:
		v152 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v152 != 0 {
			v153 = *(*int32)(unsafe.Add(mBase, uint32(v152)+4))
			if v153 < int32(0) {
				F_pfree(m, v152)
				mBase = m.M
				v167 = m.ExcPending
				if v167 != 0 {
					return int32(0)
				} else {
					v170 = F_palloc_extended(m, int32(44), int32(2))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v170)+12)) = int64(8589934592)
						*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(0)
						v179 = v170 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v179
						*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v179
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v170
						v183 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
						v186 = v170
						v188 = v183 << (uint(int32(3)) % 32)
						v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v188+v189))) = int32(0)
						v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
						v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
						v195 = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v195)%32))+4)) = int32(31)
						v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
						v201 = int32(1)
						v202 = v200 + v201
						*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v202
						v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v204+v202<<(uint(v195)%32)))) = int32(127)
						v210 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
						v211 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v210+v211<<(uint(v195)%32))+4)) = int32(159)
						v217 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v217 + v201
						return v186
					}
				}
			} else {
				v156 = *(*int32)(unsafe.Add(mBase, uint32(v152)+16))
				if v156 < int32(2) {
					F_pfree(m, v152)
					mBase = m.M
					v167 = m.ExcPending
					if v167 != 0 {
						return int32(0)
					} else {
						v170 = F_palloc_extended(m, int32(44), int32(2))
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v170)+12)) = int64(8589934592)
							*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(0)
							v179 = v170 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v179
							*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v179
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v170
							v183 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
							v186 = v170
							v188 = v183 << (uint(int32(3)) % 32)
							v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v188+v189))) = int32(0)
							v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
							v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
							v195 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v195)%32))+4)) = int32(31)
							v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
							v201 = int32(1)
							v202 = v200 + v201
							*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v202
							v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v204+v202<<(uint(v195)%32)))) = int32(127)
							v210 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
							v211 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v210+v211<<(uint(v195)%32))+4)) = int32(159)
							v217 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v217 + v201
							return v186
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v152)+24)) = int32(-1)
					v161 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v152)+12)) = v161
					*(*int32)(unsafe.Add(mBase, uint32(v152))) = v161
					v186 = v152
					v188 = v161
					v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v188+v189))) = int32(0)
					v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
					v195 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v195)%32))+4)) = int32(31)
					v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
					v201 = int32(1)
					v202 = v200 + v201
					*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v202
					v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v204+v202<<(uint(v195)%32)))) = int32(127)
					v210 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
					v211 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v210+v211<<(uint(v195)%32))+4)) = int32(159)
					v217 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v217 + v201
					return v186
				}
			}
		} else {
			v170 = F_palloc_extended(m, int32(44), int32(2))
			mBase = m.M
			v171 = m.ExcPending
			if v171 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v170)+24)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v170)+12)) = int64(8589934592)
				*(*int64)(unsafe.Add(mBase, uint32(v170))) = int64(0)
				v179 = v170 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(v170)+20)) = v179
				*(*int32)(unsafe.Add(mBase, uint32(v170)+8)) = v179
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v170
				v183 = *(*int32)(unsafe.Add(mBase, uint32(v170)+12))
				v186 = v170
				v188 = v183 << (uint(int32(3)) % 32)
				v189 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v188+v189))) = int32(0)
				v193 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
				v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
				v195 = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v193+v194<<(uint(v195)%32))+4)) = int32(31)
				v200 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
				v201 = int32(1)
				v202 = v200 + v201
				*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v202
				v204 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v204+v202<<(uint(v195)%32)))) = int32(127)
				v210 = *(*int32)(unsafe.Add(mBase, uint32(v186)+20))
				v211 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v210+v211<<(uint(v195)%32))+4)) = int32(159)
				v217 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v186)+12)) = v217 + v201
				return v186
			}
		}
	case 5:
		v224 = F_pg_ctype_get_cache(m, int32(975), int32(5))
		mBase = m.M
		v225 = m.ExcPending
		if v225 != 0 {
			return int32(0)
		} else {
			if v224 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v224
			}
			return v364
		}
	case 6:
		v343 = F_pg_ctype_get_cache(m, int32(980), int32(6))
		mBase = m.M
		v344 = m.ExcPending
		if v344 != 0 {
			return int32(0)
		} else {
			if v343 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v343
			}
			return v364
		}
	case 7:
		v331 = F_pg_ctype_get_cache(m, int32(978), int32(7))
		mBase = m.M
		v332 = m.ExcPending
		if v332 != 0 {
			return int32(0)
		} else {
			if v331 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v331
			}
			return v364
		}
	case 8:
		v15 = F_pg_ctype_get_cache(m, int32(971), int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v15
			}
			return v364
		}
	case 9:
		v230 = F_pg_ctype_get_cache(m, int32(976), int32(9))
		mBase = m.M
		v231 = m.ExcPending
		if v231 != 0 {
			return int32(0)
		} else {
			if v230 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v230
			}
			return v364
		}
	case 10:
		v325 = F_pg_ctype_get_cache(m, int32(977), int32(10))
		mBase = m.M
		v326 = m.ExcPending
		if v326 != 0 {
			return int32(0)
		} else {
			if v325 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v325
			}
			return v364
		}
	case 11:
		v337 = F_pg_ctype_get_cache(m, int32(979), int32(11))
		mBase = m.M
		v338 = m.ExcPending
		if v338 != 0 {
			return int32(0)
		} else {
			if v337 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v337
			}
			return v364
		}
	case 12:
		v234 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v234 != 0 {
			v235 = *(*int32)(unsafe.Add(mBase, uint32(v234)+4))
			if v235 < int32(0) {
				F_pfree(m, v234)
				mBase = m.M
				v249 = m.ExcPending
				if v249 != 0 {
					return int32(0)
				} else {
					v252 = F_palloc_extended(m, int32(52), int32(2))
					mBase = m.M
					v253 = m.ExcPending
					if v253 != 0 {
						return int32(0)
					} else {
						if v252 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
							v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v352 != 0 {
								v354 = v352
							} else {
								v354 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v359 != 0 {
								v361 = v359
							} else {
								v361 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
							v364 = int32(0)
							return v364
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v252))) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v252)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v252)+12)) = int64(12884901888)
							v263 = v252 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = v263
							*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v263
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v252
							v267 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
							v270 = v252
							v272 = v267 << (uint(int32(3)) % 32)
							v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v272+v273))) = int32(48)
							v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							v278 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							v279 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(v279)%32))+4)) = int32(57)
							v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							v285 = int32(1)
							v286 = v284 + v285
							*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v286
							v288 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v288+v286<<(uint(v279)%32)))) = int32(97)
							v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							v295 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v294+v295<<(uint(v279)%32))+4)) = int32(102)
							v301 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							v303 = v301 + v285
							*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v303
							v305 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v305+v303<<(uint(v279)%32)))) = int32(65)
							v311 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
							v312 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v311+v312<<(uint(v279)%32))+4)) = int32(70)
							v318 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v318 + v285
							return v270
						}
					}
				}
			} else {
				v238 = *(*int32)(unsafe.Add(mBase, uint32(v234)+16))
				if v238 < int32(3) {
					F_pfree(m, v234)
					mBase = m.M
					v249 = m.ExcPending
					if v249 != 0 {
						return int32(0)
					} else {
						v252 = F_palloc_extended(m, int32(52), int32(2))
						mBase = m.M
						v253 = m.ExcPending
						if v253 != 0 {
							return int32(0)
						} else {
							if v252 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v352 != 0 {
									v354 = v352
								} else {
									v354 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v359 != 0 {
									v361 = v359
								} else {
									v361 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
								v364 = int32(0)
								return v364
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v252))) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v252)+24)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v252)+12)) = int64(12884901888)
								v263 = v252 + int32(28)
								*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = v263
								*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v263
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v252
								v267 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
								v270 = v252
								v272 = v267 << (uint(int32(3)) % 32)
								v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v272+v273))) = int32(48)
								v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								v278 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								v279 = int32(3)
								*(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(v279)%32))+4)) = int32(57)
								v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								v285 = int32(1)
								v286 = v284 + v285
								*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v286
								v288 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v288+v286<<(uint(v279)%32)))) = int32(97)
								v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								v295 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v294+v295<<(uint(v279)%32))+4)) = int32(102)
								v301 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								v303 = v301 + v285
								*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v303
								v305 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v305+v303<<(uint(v279)%32)))) = int32(65)
								v311 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
								v312 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v311+v312<<(uint(v279)%32))+4)) = int32(70)
								v318 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v318 + v285
								return v270
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v234)+24)) = int32(-1)
					v243 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v234)+12)) = v243
					*(*int32)(unsafe.Add(mBase, uint32(v234))) = v243
					v270 = v234
					v272 = v243
					v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v272+v273))) = int32(48)
					v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v278 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v279 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(v279)%32))+4)) = int32(57)
					v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v285 = int32(1)
					v286 = v284 + v285
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v286
					v288 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v288+v286<<(uint(v279)%32)))) = int32(97)
					v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v295 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v294+v295<<(uint(v279)%32))+4)) = int32(102)
					v301 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v303 = v301 + v285
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v303
					v305 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v305+v303<<(uint(v279)%32)))) = int32(65)
					v311 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v312 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v311+v312<<(uint(v279)%32))+4)) = int32(70)
					v318 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v318 + v285
					return v270
				}
			}
		} else {
			v252 = F_palloc_extended(m, int32(52), int32(2))
			mBase = m.M
			v253 = m.ExcPending
			if v253 != 0 {
				return int32(0)
			} else {
				if v252 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
					v352 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v352 != 0 {
						v354 = v352
					} else {
						v354 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v354
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
					v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v359 != 0 {
						v361 = v359
					} else {
						v361 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
					v364 = int32(0)
					return v364
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v252))) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v252)+24)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v252)+12)) = int64(12884901888)
					v263 = v252 + int32(28)
					*(*int32)(unsafe.Add(mBase, uint32(v252)+20)) = v263
					*(*int32)(unsafe.Add(mBase, uint32(v252)+8)) = v263
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v252
					v267 = *(*int32)(unsafe.Add(mBase, uint32(v252)+12))
					v270 = v252
					v272 = v267 << (uint(int32(3)) % 32)
					v273 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v272+v273))) = int32(48)
					v277 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v278 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v279 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v277+v278<<(uint(v279)%32))+4)) = int32(57)
					v284 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v285 = int32(1)
					v286 = v284 + v285
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v286
					v288 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v288+v286<<(uint(v279)%32)))) = int32(97)
					v294 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v295 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v294+v295<<(uint(v279)%32))+4)) = int32(102)
					v301 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					v303 = v301 + v285
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v303
					v305 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v305+v303<<(uint(v279)%32)))) = int32(65)
					v311 = *(*int32)(unsafe.Add(mBase, uint32(v270)+20))
					v312 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v311+v312<<(uint(v279)%32))+4)) = int32(70)
					v318 = *(*int32)(unsafe.Add(mBase, uint32(v270)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v270)+12)) = v318 + v285
					return v270
				}
			}
		}
	case 13:
		v35 = F_pg_ctype_get_cache(m, int32(974), int32(13))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			if v35 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v359 != 0 {
					v361 = v359
				} else {
					v361 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
				v364 = int32(0)
			} else {
				v364 = v35
			}
			return v364
		}
	default:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
		v359 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v359 != 0 {
			v361 = v359
		} else {
			v361 = int32(12)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v361
		v364 = int32(0)
		return v364
	}
}
