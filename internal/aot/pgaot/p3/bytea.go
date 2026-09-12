package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_bytea_bit_count(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int64
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v66 int64
	_ = v66
	var v67 int32
	_ = v67
	var v70 int64
	_ = v70
	var v71 int32
	_ = v71
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v82 int64
	_ = v82
	var v86 int64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v97 int64
	_ = v97
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v106 int64
	_ = v106
	var v107 int32
	_ = v107
	var v110 int64
	_ = v110
	var v111 int64
	_ = v111
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v119 int32
	_ = v119
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v128 int64
	_ = v128
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v155 int64
	_ = v155
	var v156 int32
	_ = v156
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v171 int64
	_ = v171
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v185 int64
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v195 int64
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v201 int64
	_ = v201
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v213 int64
	_ = v213
	var v217 int32
	_ = v217
	var v223 int32
	_ = v223
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 int64
	_ = v233
	var v234 int64
	_ = v234
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v240 int32
	_ = v240
	var v242 int32
	_ = v242
	var v246 int64
	_ = v246
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v255 int64
	_ = v255
	var v256 int32
	_ = v256
	var v259 int64
	_ = v259
	var v260 int32
	_ = v260
	var v263 int64
	_ = v263
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v268 int32
	_ = v268
	var v271 int64
	_ = v271
	var v275 int64
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v286 int64
	_ = v286
	var v292 int64
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	v6 = int64(0)
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum_packed(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		v12 = int32(1)
		v13 = v8 + v12
		v14 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v8))))
		v16 = v14 & v12
		if v14 == v12 {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if base.Ui32((v20-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				if v16 != 0 {
					v119 = v13
				} else {
					v119 = v8 + int32(4)
				}
				v121 = int32(4)
				v123 = v119
				v128 = int64(0)
				if v121 < int32(4) {
					v207 = v123
					v208 = v121
					v213 = v128
				} else {
					if v123 != (v123+int32(3))&int32(-4) {
						v207 = v123
						v208 = v121
						v213 = v128
					} else {
						v137 = v121 - int32(4)
						v141 = int32(base.Ui32(v137)>>(uint(int32(2))%32)) + int32(1)
						v143 = v141 & int32(3)
						if base.Ui32(v137) < base.Ui32(int32(12)) {
							v179 = v123
							v180 = v121
							v185 = v128
						} else {
							v149 = v123
							v150 = v121
							v151 = int32(0)
							v155 = v128
							for {
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
								v165 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
								v171 = base.I64_extend_i32_u(base.I32_popcnt(v156)) + (base.I64_extend_i32_u(base.I32_popcnt(v159)) + (base.I64_extend_i32_u(base.I32_popcnt(v162)) + (v155 + base.I64_extend_i32_u(base.I32_popcnt(v165)))))
								v172 = int32(16)
								v173 = v150 - v172
								v175 = v149 + v172
								v177 = v151 + int32(4)
								if v177 != v141&int32(2147483644) {
									v149 = v175
									v150 = v173
									v151 = v177
									v155 = v171
									continue
								} else {
									break
								}
								break
							}
							v179 = v175
							v180 = v173
							v185 = v171
						}
						if v143 == int32(0) {
							v207 = v179
							v208 = v180
							v213 = v185
						} else {
							v190 = v180
							v191 = v179
							v192 = int32(0)
							v195 = v185
							for {
								v196 = int32(4)
								v197 = v190 - v196
								v198 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
								v201 = v195 + base.I64_extend_i32_u(base.I32_popcnt(v198))
								v203 = v191 + v196
								v205 = v192 + int32(1)
								if v205 != v143 {
									v190 = v197
									v191 = v203
									v192 = v205
									v195 = v201
									continue
								} else {
									break
								}
								break
							}
							v207 = v203
							v208 = v197
							v213 = v201
						}
					}
				}
				if v208 == int32(0) {
					v286 = v213
				} else {
					v217 = v208 & int32(3)
					if v217 == int32(0) {
						v240 = v207
						v242 = v208
						v246 = v213
					} else {
						v223 = v208
						v224 = v207
						v225 = int32(0)
						v227 = v213
						for {
							v228 = int32(1)
							v229 = v223 - v228
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
							v233 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v230)+uint32(_consts[1054]))))
							v234 = v227 + v233
							v236 = v224 + v228
							v238 = v225 + v228
							if v238 != v217 {
								v223 = v229
								v224 = v236
								v225 = v238
								v227 = v234
								continue
							} else {
								break
							}
							break
						}
						v240 = v236
						v242 = v229
						v246 = v234
					}
					if base.Ui32(v208) < base.Ui32(int32(4)) {
						v286 = v246
					} else {
						v249 = v240
						v251 = v242
						v255 = v246
						for {
							v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+3)))
							v259 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v256)+uint32(_consts[1054]))))
							v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+2)))
							v263 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v260)+uint32(_consts[1054]))))
							v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
							v267 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v264)+uint32(_consts[1054]))))
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
							v271 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v268)+uint32(_consts[1054]))))
							v275 = v259 + (v263 + (v267 + (v255 + v271)))
							v276 = int32(4)
							v279 = v251 - v276
							if v279 != 0 {
								v249 = v249 + v276
								v251 = v279
								v255 = v275
								continue
							} else {
								break
							}
							break
						}
						v286 = v275
					}
				}
				v292 = v286
				v293 = F_Int64GetDatum(m, v292)
				mBase = m.M
				v294 = m.ExcPending
				if v294 != 0 {
					return int32(0)
				} else {
					return v293
				}
			} else {
				v42 = base.B2i32(v20 == int32(18)) << (uint(int32(4)) % 32)
				if v16 != 0 {
					v45 = v13
				} else {
					v45 = v8 + int32(4)
				}
				if int32(3) < v42 {
					v121 = v42
					v123 = v45
					v128 = int64(0)
					if v121 < int32(4) {
						v207 = v123
						v208 = v121
						v213 = v128
					} else {
						if v123 != (v123+int32(3))&int32(-4) {
							v207 = v123
							v208 = v121
							v213 = v128
						} else {
							v137 = v121 - int32(4)
							v141 = int32(base.Ui32(v137)>>(uint(int32(2))%32)) + int32(1)
							v143 = v141 & int32(3)
							if base.Ui32(v137) < base.Ui32(int32(12)) {
								v179 = v123
								v180 = v121
								v185 = v128
							} else {
								v149 = v123
								v150 = v121
								v151 = int32(0)
								v155 = v128
								for {
									v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
									v159 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
									v162 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
									v165 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
									v171 = base.I64_extend_i32_u(base.I32_popcnt(v156)) + (base.I64_extend_i32_u(base.I32_popcnt(v159)) + (base.I64_extend_i32_u(base.I32_popcnt(v162)) + (v155 + base.I64_extend_i32_u(base.I32_popcnt(v165)))))
									v172 = int32(16)
									v173 = v150 - v172
									v175 = v149 + v172
									v177 = v151 + int32(4)
									if v177 != v141&int32(2147483644) {
										v149 = v175
										v150 = v173
										v151 = v177
										v155 = v171
										continue
									} else {
										break
									}
									break
								}
								v179 = v175
								v180 = v173
								v185 = v171
							}
							if v143 == int32(0) {
								v207 = v179
								v208 = v180
								v213 = v185
							} else {
								v190 = v180
								v191 = v179
								v192 = int32(0)
								v195 = v185
								for {
									v196 = int32(4)
									v197 = v190 - v196
									v198 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
									v201 = v195 + base.I64_extend_i32_u(base.I32_popcnt(v198))
									v203 = v191 + v196
									v205 = v192 + int32(1)
									if v205 != v143 {
										v190 = v197
										v191 = v203
										v192 = v205
										v195 = v201
										continue
									} else {
										break
									}
									break
								}
								v207 = v203
								v208 = v197
								v213 = v201
							}
						}
					}
					if v208 == int32(0) {
						v286 = v213
					} else {
						v217 = v208 & int32(3)
						if v217 == int32(0) {
							v240 = v207
							v242 = v208
							v246 = v213
						} else {
							v223 = v208
							v224 = v207
							v225 = int32(0)
							v227 = v213
							for {
								v228 = int32(1)
								v229 = v223 - v228
								v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
								v233 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v230)+uint32(_consts[1054]))))
								v234 = v227 + v233
								v236 = v224 + v228
								v238 = v225 + v228
								if v238 != v217 {
									v223 = v229
									v224 = v236
									v225 = v238
									v227 = v234
									continue
								} else {
									break
								}
								break
							}
							v240 = v236
							v242 = v229
							v246 = v234
						}
						if base.Ui32(v208) < base.Ui32(int32(4)) {
							v286 = v246
						} else {
							v249 = v240
							v251 = v242
							v255 = v246
							for {
								v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+3)))
								v259 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v256)+uint32(_consts[1054]))))
								v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+2)))
								v263 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v260)+uint32(_consts[1054]))))
								v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
								v267 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v264)+uint32(_consts[1054]))))
								v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
								v271 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v268)+uint32(_consts[1054]))))
								v275 = v259 + (v263 + (v267 + (v255 + v271)))
								v276 = int32(4)
								v279 = v251 - v276
								if v279 != 0 {
									v249 = v249 + v276
									v251 = v279
									v255 = v275
									continue
								} else {
									break
								}
								break
							}
							v286 = v275
						}
					}
					v292 = v286
					v293 = F_Int64GetDatum(m, v292)
					mBase = m.M
					v294 = m.ExcPending
					if v294 != 0 {
						return int32(0)
					} else {
						return v293
					}
				} else {
					if v42 == int32(0) {
						v51 = F_Int64GetDatum(m, int64(0))
						mBase = m.M
						v52 = m.ExcPending
						if v52 != 0 {
							return int32(0)
						} else {
							return v51
						}
					} else {
						v55 = v42 & int32(3)
						if base.Ui32(v42) < base.Ui32(int32(4)) {
							v92 = v45
							v97 = v6
						} else {
							v61 = v45
							v62 = int32(0)
							v66 = v6
							for {
								v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
								v70 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1054]))))
								v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
								v74 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[1054]))))
								v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
								v78 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[1054]))))
								v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
								v82 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v79)+uint32(_consts[1054]))))
								v86 = v70 + (v74 + (v78 + (v66 + v82)))
								v87 = int32(4)
								v88 = v61 + v87
								v90 = v62 + v87
								if v90 != v42&int32(-4) {
									v61 = v88
									v62 = v90
									v66 = v86
									continue
								} else {
									break
								}
								break
							}
							v92 = v88
							v97 = v86
						}
						if v55 == int32(0) {
							v292 = v97
						} else {
							v101 = v92
							v102 = int32(0)
							v106 = v97
							for {
								v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
								v110 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v107)+uint32(_consts[1054]))))
								v111 = v106 + v110
								v112 = int32(1)
								v115 = v102 + v112
								if v115 != v55 {
									v101 = v101 + v112
									v102 = v115
									v106 = v111
									continue
								} else {
									break
								}
								break
							}
							v292 = v111
						}
						v293 = F_Int64GetDatum(m, v292)
						mBase = m.M
						v294 = m.ExcPending
						if v294 != 0 {
							return int32(0)
						} else {
							return v293
						}
					}
				}
			}
		} else {
			v31 = int32(1)
			if v16 != 0 {
				v42 = int32(base.Ui32(v14)>>(uint(v31)%32)) - v31
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v42 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
			if v16 != 0 {
				v45 = v13
			} else {
				v45 = v8 + int32(4)
			}
			if int32(3) < v42 {
				v121 = v42
				v123 = v45
				v128 = int64(0)
				if v121 < int32(4) {
					v207 = v123
					v208 = v121
					v213 = v128
				} else {
					if v123 != (v123+int32(3))&int32(-4) {
						v207 = v123
						v208 = v121
						v213 = v128
					} else {
						v137 = v121 - int32(4)
						v141 = int32(base.Ui32(v137)>>(uint(int32(2))%32)) + int32(1)
						v143 = v141 & int32(3)
						if base.Ui32(v137) < base.Ui32(int32(12)) {
							v179 = v123
							v180 = v121
							v185 = v128
						} else {
							v149 = v123
							v150 = v121
							v151 = int32(0)
							v155 = v128
							for {
								v156 = *(*int32)(unsafe.Add(mBase, uint32(v149)+12))
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v149)+8))
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v149)+4))
								v165 = *(*int32)(unsafe.Add(mBase, uint32(v149)))
								v171 = base.I64_extend_i32_u(base.I32_popcnt(v156)) + (base.I64_extend_i32_u(base.I32_popcnt(v159)) + (base.I64_extend_i32_u(base.I32_popcnt(v162)) + (v155 + base.I64_extend_i32_u(base.I32_popcnt(v165)))))
								v172 = int32(16)
								v173 = v150 - v172
								v175 = v149 + v172
								v177 = v151 + int32(4)
								if v177 != v141&int32(2147483644) {
									v149 = v175
									v150 = v173
									v151 = v177
									v155 = v171
									continue
								} else {
									break
								}
								break
							}
							v179 = v175
							v180 = v173
							v185 = v171
						}
						if v143 == int32(0) {
							v207 = v179
							v208 = v180
							v213 = v185
						} else {
							v190 = v180
							v191 = v179
							v192 = int32(0)
							v195 = v185
							for {
								v196 = int32(4)
								v197 = v190 - v196
								v198 = *(*int32)(unsafe.Add(mBase, uint32(v191)))
								v201 = v195 + base.I64_extend_i32_u(base.I32_popcnt(v198))
								v203 = v191 + v196
								v205 = v192 + int32(1)
								if v205 != v143 {
									v190 = v197
									v191 = v203
									v192 = v205
									v195 = v201
									continue
								} else {
									break
								}
								break
							}
							v207 = v203
							v208 = v197
							v213 = v201
						}
					}
				}
				if v208 == int32(0) {
					v286 = v213
				} else {
					v217 = v208 & int32(3)
					if v217 == int32(0) {
						v240 = v207
						v242 = v208
						v246 = v213
					} else {
						v223 = v208
						v224 = v207
						v225 = int32(0)
						v227 = v213
						for {
							v228 = int32(1)
							v229 = v223 - v228
							v230 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v224))))
							v233 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v230)+uint32(_consts[1054]))))
							v234 = v227 + v233
							v236 = v224 + v228
							v238 = v225 + v228
							if v238 != v217 {
								v223 = v229
								v224 = v236
								v225 = v238
								v227 = v234
								continue
							} else {
								break
							}
							break
						}
						v240 = v236
						v242 = v229
						v246 = v234
					}
					if base.Ui32(v208) < base.Ui32(int32(4)) {
						v286 = v246
					} else {
						v249 = v240
						v251 = v242
						v255 = v246
						for {
							v256 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+3)))
							v259 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v256)+uint32(_consts[1054]))))
							v260 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+2)))
							v263 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v260)+uint32(_consts[1054]))))
							v264 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249)+1)))
							v267 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v264)+uint32(_consts[1054]))))
							v268 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v249))))
							v271 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v268)+uint32(_consts[1054]))))
							v275 = v259 + (v263 + (v267 + (v255 + v271)))
							v276 = int32(4)
							v279 = v251 - v276
							if v279 != 0 {
								v249 = v249 + v276
								v251 = v279
								v255 = v275
								continue
							} else {
								break
							}
							break
						}
						v286 = v275
					}
				}
				v292 = v286
				v293 = F_Int64GetDatum(m, v292)
				mBase = m.M
				v294 = m.ExcPending
				if v294 != 0 {
					return int32(0)
				} else {
					return v293
				}
			} else {
				if v42 == int32(0) {
					v51 = F_Int64GetDatum(m, int64(0))
					mBase = m.M
					v52 = m.ExcPending
					if v52 != 0 {
						return int32(0)
					} else {
						return v51
					}
				} else {
					v55 = v42 & int32(3)
					if base.Ui32(v42) < base.Ui32(int32(4)) {
						v92 = v45
						v97 = v6
					} else {
						v61 = v45
						v62 = int32(0)
						v66 = v6
						for {
							v67 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+3)))
							v70 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v67)+uint32(_consts[1054]))))
							v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+2)))
							v74 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v71)+uint32(_consts[1054]))))
							v75 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61)+1)))
							v78 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v75)+uint32(_consts[1054]))))
							v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v61))))
							v82 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v79)+uint32(_consts[1054]))))
							v86 = v70 + (v74 + (v78 + (v66 + v82)))
							v87 = int32(4)
							v88 = v61 + v87
							v90 = v62 + v87
							if v90 != v42&int32(-4) {
								v61 = v88
								v62 = v90
								v66 = v86
								continue
							} else {
								break
							}
							break
						}
						v92 = v88
						v97 = v86
					}
					if v55 == int32(0) {
						v292 = v97
					} else {
						v101 = v92
						v102 = int32(0)
						v106 = v97
						for {
							v107 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v101))))
							v110 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v107)+uint32(_consts[1054]))))
							v111 = v106 + v110
							v112 = int32(1)
							v115 = v102 + v112
							if v115 != v55 {
								v101 = v101 + v112
								v102 = v115
								v106 = v111
								continue
							} else {
								break
							}
							break
						}
						v292 = v111
					}
					v293 = F_Int64GetDatum(m, v292)
					mBase = m.M
					v294 = m.ExcPending
					if v294 != 0 {
						return int32(0)
					} else {
						return v293
					}
				}
			}
		}
	}
}
func F_bytea_substr_no_len(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = int32(1)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v4 <= v3 {
		v7 = v3
	} else {
		v7 = v4
	}
	v11 = F_pg_detoast_datum_slice(m, v2, v7-int32(1), int32(-1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
