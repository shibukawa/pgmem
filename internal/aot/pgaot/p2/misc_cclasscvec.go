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
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v166 int32
	_ = v166
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v233 int32
	_ = v233
	var v239 int32
	_ = v239
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v263 int32
	_ = v263
	var v267 int32
	_ = v267
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v278 int32
	_ = v278
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v291 int32
	_ = v291
	var v293 int32
	_ = v293
	var v295 int32
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v328 int32
	_ = v328
	var v329 int32
	_ = v329
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v342 int32
	_ = v342
	var v344 int32
	_ = v344
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v353 int32
	_ = v353
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
		v23 = F_regc_ctype_get_cache(m, int32(1038), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			if v23 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v23
			}
			return v348
		}
	case 1:
		v29 = F_regc_ctype_get_cache(m, int32(1039), int32(1))
		mBase = m.M
		v30 = m.ExcPending
		if v30 != 0 {
			return int32(0)
		} else {
			if v29 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v29
			}
			return v348
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
							v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v335 != 0 {
								v337 = v335
							} else {
								v337 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v342 != 0 {
								v344 = v342
							} else {
								v344 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
							v348 = int32(0)
							return v348
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
							v351 = v75
							v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
							return v351
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
								v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v335 != 0 {
									v337 = v335
								} else {
									v337 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v342 != 0 {
									v344 = v342
								} else {
									v344 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
								v348 = int32(0)
								return v348
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
								v351 = v75
								v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
								return v351
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
					v351 = v75
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
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
					v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v335 != 0 {
						v337 = v335
					} else {
						v337 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
					v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v342 != 0 {
						v344 = v342
					} else {
						v344 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
					v348 = int32(0)
					return v348
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
					v351 = v75
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
				}
			}
		}
	case 3:
		v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v89 != 0 {
			v90 = *(*int32)(unsafe.Add(mBase, uint32(v89)+4))
			if v90 < int32(2) {
				F_pfree(m, v89)
				mBase = m.M
				v105 = m.ExcPending
				if v105 != 0 {
					return int32(0)
				} else {
					v109 = F_palloc_extended(m, int32(36), int32(2))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v109)+12)) = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v109))) = int64(8589934592)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v109 + int32(36)
						*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v109 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v109
						v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
						v125 = v109
						v126 = v124
						v127 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126 + v127
						v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
						v131 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v130+v126<<(uint(v131)%32)))) = int32(9)
						v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
						*(*int32)(unsafe.Add(mBase, uint32(v125))) = v136 + v127
						v140 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(v131)%32)))) = int32(32)
						return v125
					}
				}
			} else {
				v93 = int32(0)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v89)+16))
				if v94 < v93 {
					F_pfree(m, v89)
					mBase = m.M
					v105 = m.ExcPending
					if v105 != 0 {
						return int32(0)
					} else {
						v109 = F_palloc_extended(m, int32(36), int32(2))
						mBase = m.M
						v110 = m.ExcPending
						if v110 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v109)+12)) = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v109))) = int64(8589934592)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v109 + int32(36)
							*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v109 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v109
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
							v125 = v109
							v126 = v124
							v127 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126 + v127
							v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
							v131 = int32(2)
							*(*int32)(unsafe.Add(mBase, uint32(v130+v126<<(uint(v131)%32)))) = int32(9)
							v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
							*(*int32)(unsafe.Add(mBase, uint32(v125))) = v136 + v127
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(v131)%32)))) = int32(32)
							return v125
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v89)+24)) = int32(-1)
					v99 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v89)+12)) = v99
					*(*int32)(unsafe.Add(mBase, uint32(v89))) = v99
					v125 = v89
					v126 = v93
					v127 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126 + v127
					v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
					v131 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v130+v126<<(uint(v131)%32)))) = int32(9)
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
					*(*int32)(unsafe.Add(mBase, uint32(v125))) = v136 + v127
					v140 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(v131)%32)))) = int32(32)
					return v125
				}
			}
		} else {
			v109 = F_palloc_extended(m, int32(36), int32(2))
			mBase = m.M
			v110 = m.ExcPending
			if v110 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v109)+24)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v109)+12)) = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v109))) = int64(8589934592)
				*(*int32)(unsafe.Add(mBase, uint32(v109)+20)) = v109 + int32(36)
				*(*int32)(unsafe.Add(mBase, uint32(v109)+8)) = v109 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v109
				v124 = *(*int32)(unsafe.Add(mBase, uint32(v109)))
				v125 = v109
				v126 = v124
				v127 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v125))) = v126 + v127
				v130 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
				v131 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v130+v126<<(uint(v131)%32)))) = int32(9)
				v136 = *(*int32)(unsafe.Add(mBase, uint32(v125)))
				*(*int32)(unsafe.Add(mBase, uint32(v125))) = v136 + v127
				v140 = *(*int32)(unsafe.Add(mBase, uint32(v125)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v140+v136<<(uint(v131)%32)))) = int32(32)
				return v125
			}
		}
	case 4:
		v147 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v147 != 0 {
			v148 = *(*int32)(unsafe.Add(mBase, uint32(v147)+4))
			if v148 < int32(0) {
				F_pfree(m, v147)
				mBase = m.M
				v162 = m.ExcPending
				if v162 != 0 {
					return int32(0)
				} else {
					v165 = F_palloc_extended(m, int32(44), int32(2))
					mBase = m.M
					v166 = m.ExcPending
					if v166 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = int32(-1)
						*(*int64)(unsafe.Add(mBase, uint32(v165)+12)) = int64(8589934592)
						*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(0)
						v174 = v165 + int32(28)
						*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v174
						*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v174
						*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v165
						v178 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
						v181 = v165
						v183 = v178 << (uint(int32(3)) % 32)
						v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = int32(0)
						v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
						v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
						v190 = int32(3)
						*(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(v190)%32))+4)) = int32(31)
						v195 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
						v197 = v195 + int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v197
						v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
						*(*int32)(unsafe.Add(mBase, uint32(v199+v197<<(uint(v190)%32)))) = int32(127)
						v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
						v206 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(v190)%32))+4)) = int32(159)
						v351 = v181
						v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
						*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
						return v351
					}
				}
			} else {
				v151 = *(*int32)(unsafe.Add(mBase, uint32(v147)+16))
				if v151 < int32(2) {
					F_pfree(m, v147)
					mBase = m.M
					v162 = m.ExcPending
					if v162 != 0 {
						return int32(0)
					} else {
						v165 = F_palloc_extended(m, int32(44), int32(2))
						mBase = m.M
						v166 = m.ExcPending
						if v166 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v165)+12)) = int64(8589934592)
							*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(0)
							v174 = v165 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v174
							*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v174
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v165
							v178 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
							v181 = v165
							v183 = v178 << (uint(int32(3)) % 32)
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = int32(0)
							v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
							v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
							v190 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(v190)%32))+4)) = int32(31)
							v195 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
							v197 = v195 + int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v197
							v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v199+v197<<(uint(v190)%32)))) = int32(127)
							v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
							v206 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(v190)%32))+4)) = int32(159)
							v351 = v181
							v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
							return v351
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v147)+24)) = int32(-1)
					v156 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v147)+12)) = v156
					*(*int32)(unsafe.Add(mBase, uint32(v147))) = v156
					v181 = v147
					v183 = v156
					v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = int32(0)
					v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
					v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
					v190 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(v190)%32))+4)) = int32(31)
					v195 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
					v197 = v195 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v197
					v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v199+v197<<(uint(v190)%32)))) = int32(127)
					v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
					v206 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(v190)%32))+4)) = int32(159)
					v351 = v181
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
				}
			}
		} else {
			v165 = F_palloc_extended(m, int32(44), int32(2))
			mBase = m.M
			v166 = m.ExcPending
			if v166 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = int32(-1)
				*(*int64)(unsafe.Add(mBase, uint32(v165)+12)) = int64(8589934592)
				*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(0)
				v174 = v165 + int32(28)
				*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v174
				*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v174
				*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v165
				v178 = *(*int32)(unsafe.Add(mBase, uint32(v165)+12))
				v181 = v165
				v183 = v178 << (uint(int32(3)) % 32)
				v184 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v183+v184))) = int32(0)
				v188 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
				v189 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
				v190 = int32(3)
				*(*int32)(unsafe.Add(mBase, uint32(v188+v189<<(uint(v190)%32))+4)) = int32(31)
				v195 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
				v197 = v195 + int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v181)+12)) = v197
				v199 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
				*(*int32)(unsafe.Add(mBase, uint32(v199+v197<<(uint(v190)%32)))) = int32(127)
				v205 = *(*int32)(unsafe.Add(mBase, uint32(v181)+20))
				v206 = *(*int32)(unsafe.Add(mBase, uint32(v181)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v205+v206<<(uint(v190)%32))+4)) = int32(159)
				v351 = v181
				v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
				return v351
			}
		}
	case 5:
		v214 = F_regc_ctype_get_cache(m, int32(1041), int32(5))
		mBase = m.M
		v215 = m.ExcPending
		if v215 != 0 {
			return int32(0)
		} else {
			if v214 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v214
			}
			return v348
		}
	case 6:
		v328 = F_regc_ctype_get_cache(m, int32(1046), int32(6))
		mBase = m.M
		v329 = m.ExcPending
		if v329 != 0 {
			return int32(0)
		} else {
			if v328 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v328
			}
			return v348
		}
	case 7:
		v316 = F_regc_ctype_get_cache(m, int32(1044), int32(7))
		mBase = m.M
		v317 = m.ExcPending
		if v317 != 0 {
			return int32(0)
		} else {
			if v316 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v316
			}
			return v348
		}
	case 8:
		v15 = F_regc_ctype_get_cache(m, int32(1037), int32(8))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			if v15 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v15
			}
			return v348
		}
	case 9:
		v220 = F_regc_ctype_get_cache(m, int32(1042), int32(9))
		mBase = m.M
		v221 = m.ExcPending
		if v221 != 0 {
			return int32(0)
		} else {
			if v220 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v220
			}
			return v348
		}
	case 10:
		v310 = F_regc_ctype_get_cache(m, int32(1043), int32(10))
		mBase = m.M
		v311 = m.ExcPending
		if v311 != 0 {
			return int32(0)
		} else {
			if v310 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v310
			}
			return v348
		}
	case 11:
		v322 = F_regc_ctype_get_cache(m, int32(1045), int32(11))
		mBase = m.M
		v323 = m.ExcPending
		if v323 != 0 {
			return int32(0)
		} else {
			if v322 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v322
			}
			return v348
		}
	case 12:
		v224 = *(*int32)(unsafe.Add(mBase, uint32(l0)+120))
		if v224 != 0 {
			v225 = *(*int32)(unsafe.Add(mBase, uint32(v224)+4))
			if v225 < int32(0) {
				F_pfree(m, v224)
				mBase = m.M
				v239 = m.ExcPending
				if v239 != 0 {
					return int32(0)
				} else {
					v242 = F_palloc_extended(m, int32(52), int32(2))
					mBase = m.M
					v243 = m.ExcPending
					if v243 != 0 {
						return int32(0)
					} else {
						if v242 == int32(0) {
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
							v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v335 != 0 {
								v337 = v335
							} else {
								v337 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
							*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
							v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
							if v342 != 0 {
								v344 = v342
							} else {
								v344 = int32(12)
							}
							*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
							v348 = int32(0)
							return v348
						} else {
							*(*int64)(unsafe.Add(mBase, uint32(v242))) = int64(0)
							*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = int32(-1)
							*(*int64)(unsafe.Add(mBase, uint32(v242)+12)) = int64(12884901888)
							v253 = v242 + int32(28)
							*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v253
							*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v253
							*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v242
							v257 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
							v260 = v242
							v262 = v257 << (uint(int32(3)) % 32)
							v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v262+v263))) = int32(48)
							v267 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							v269 = int32(3)
							*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32))+4)) = int32(57)
							v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							v275 = int32(1)
							v276 = v274 + v275
							*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v276
							v278 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v278+v276<<(uint(v269)%32)))) = int32(97)
							v284 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							v285 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(v269)%32))+4)) = int32(102)
							v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							v293 = v291 + v275
							*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v293
							v295 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							*(*int32)(unsafe.Add(mBase, uint32(v295+v293<<(uint(v269)%32)))) = int32(65)
							v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
							v302 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v301+v302<<(uint(v269)%32))+4)) = int32(70)
							v351 = v260
							v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
							*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
							return v351
						}
					}
				}
			} else {
				v228 = *(*int32)(unsafe.Add(mBase, uint32(v224)+16))
				if v228 < int32(3) {
					F_pfree(m, v224)
					mBase = m.M
					v239 = m.ExcPending
					if v239 != 0 {
						return int32(0)
					} else {
						v242 = F_palloc_extended(m, int32(52), int32(2))
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return int32(0)
						} else {
							if v242 == int32(0) {
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
								v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v335 != 0 {
									v337 = v335
								} else {
									v337 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
								*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
								v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
								if v342 != 0 {
									v344 = v342
								} else {
									v344 = int32(12)
								}
								*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
								v348 = int32(0)
								return v348
							} else {
								*(*int64)(unsafe.Add(mBase, uint32(v242))) = int64(0)
								*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = int32(-1)
								*(*int64)(unsafe.Add(mBase, uint32(v242)+12)) = int64(12884901888)
								v253 = v242 + int32(28)
								*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v253
								*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v253
								*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v242
								v257 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
								v260 = v242
								v262 = v257 << (uint(int32(3)) % 32)
								v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v262+v263))) = int32(48)
								v267 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								v269 = int32(3)
								*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32))+4)) = int32(57)
								v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								v275 = int32(1)
								v276 = v274 + v275
								*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v276
								v278 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v278+v276<<(uint(v269)%32)))) = int32(97)
								v284 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								v285 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(v269)%32))+4)) = int32(102)
								v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								v293 = v291 + v275
								*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v293
								v295 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v295+v293<<(uint(v269)%32)))) = int32(65)
								v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
								v302 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v301+v302<<(uint(v269)%32))+4)) = int32(70)
								v351 = v260
								v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
								*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
								return v351
							}
						}
					}
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v224)+24)) = int32(-1)
					v233 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v224)+12)) = v233
					*(*int32)(unsafe.Add(mBase, uint32(v224))) = v233
					v260 = v224
					v262 = v233
					v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v262+v263))) = int32(48)
					v267 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v269 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32))+4)) = int32(57)
					v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v275 = int32(1)
					v276 = v274 + v275
					*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v276
					v278 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v278+v276<<(uint(v269)%32)))) = int32(97)
					v284 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v285 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(v269)%32))+4)) = int32(102)
					v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v293 = v291 + v275
					*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v293
					v295 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v295+v293<<(uint(v269)%32)))) = int32(65)
					v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v302 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v301+v302<<(uint(v269)%32))+4)) = int32(70)
					v351 = v260
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
				}
			}
		} else {
			v242 = F_palloc_extended(m, int32(52), int32(2))
			mBase = m.M
			v243 = m.ExcPending
			if v243 != 0 {
				return int32(0)
			} else {
				if v242 == int32(0) {
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = int32(0)
					v335 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v335 != 0 {
						v337 = v335
					} else {
						v337 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v337
					*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
					v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
					if v342 != 0 {
						v344 = v342
					} else {
						v344 = int32(12)
					}
					*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
					v348 = int32(0)
					return v348
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v242))) = int64(0)
					*(*int32)(unsafe.Add(mBase, uint32(v242)+24)) = int32(-1)
					*(*int64)(unsafe.Add(mBase, uint32(v242)+12)) = int64(12884901888)
					v253 = v242 + int32(28)
					*(*int32)(unsafe.Add(mBase, uint32(v242)+20)) = v253
					*(*int32)(unsafe.Add(mBase, uint32(v242)+8)) = v253
					*(*int32)(unsafe.Add(mBase, uint32(l0)+120)) = v242
					v257 = *(*int32)(unsafe.Add(mBase, uint32(v242)+12))
					v260 = v242
					v262 = v257 << (uint(int32(3)) % 32)
					v263 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v262+v263))) = int32(48)
					v267 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v268 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v269 = int32(3)
					*(*int32)(unsafe.Add(mBase, uint32(v267+v268<<(uint(v269)%32))+4)) = int32(57)
					v274 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v275 = int32(1)
					v276 = v274 + v275
					*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v276
					v278 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v278+v276<<(uint(v269)%32)))) = int32(97)
					v284 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v285 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v284+v285<<(uint(v269)%32))+4)) = int32(102)
					v291 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					v293 = v291 + v275
					*(*int32)(unsafe.Add(mBase, uint32(v260)+12)) = v293
					v295 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					*(*int32)(unsafe.Add(mBase, uint32(v295+v293<<(uint(v269)%32)))) = int32(65)
					v301 = *(*int32)(unsafe.Add(mBase, uint32(v260)+20))
					v302 = *(*int32)(unsafe.Add(mBase, uint32(v260)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v301+v302<<(uint(v269)%32))+4)) = int32(70)
					v351 = v260
					v353 = *(*int32)(unsafe.Add(mBase, uint32(v351)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v351)+12)) = v353 + int32(1)
					return v351
				}
			}
		}
	case 13:
		v35 = F_regc_ctype_get_cache(m, int32(1040), int32(13))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			if v35 == int32(0) {
				*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
				v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
				if v342 != 0 {
					v344 = v342
				} else {
					v344 = int32(12)
				}
				*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
				v348 = int32(0)
			} else {
				v348 = v35
			}
			return v348
		}
	default:
		*(*int32)(unsafe.Add(mBase, uint32(l0)+24)) = int32(101)
		v342 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		if v342 != 0 {
			v344 = v342
		} else {
			v344 = int32(12)
		}
		*(*int32)(unsafe.Add(mBase, uint32(l0)+12)) = v344
		v348 = int32(0)
		return v348
	}
}
