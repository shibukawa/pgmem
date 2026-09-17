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
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v68 int32
	_ = v68
	var v69 int64
	_ = v69
	var v70 int32
	_ = v70
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int64
	_ = v73
	var v74 int32
	_ = v74
	var v75 int64
	_ = v75
	var v79 int64
	_ = v79
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int64
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int32
	_ = v101
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v117 int64
	_ = v117
	var v127 int32
	_ = v127
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int64
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v158 int32
	_ = v158
	var v164 int64
	_ = v164
	var v166 int32
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v176 int64
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int64
	_ = v184
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int64
	_ = v192
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int64
	_ = v202
	var v206 int32
	_ = v206
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v220 int32
	_ = v220
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v223 int64
	_ = v223
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v229 int32
	_ = v229
	var v233 int64
	_ = v233
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v242 int64
	_ = v242
	var v243 int32
	_ = v243
	var v244 int64
	_ = v244
	var v245 int32
	_ = v245
	var v246 int64
	_ = v246
	var v247 int32
	_ = v247
	var v248 int64
	_ = v248
	var v249 int32
	_ = v249
	var v250 int64
	_ = v250
	var v254 int64
	_ = v254
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v265 int64
	_ = v265
	var v271 int64
	_ = v271
	var v272 int32
	_ = v272
	var v273 int32
	_ = v273
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
			v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
			if base.Ui32((v19-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v106 = int32(4)
				if v16 != 0 {
					v109 = v13
				} else {
					v109 = v8 + v106
				}
				v111 = v106
				v112 = v109
				v117 = int64(0)
				if base.B2i32(v112 != (v112+int32(3))&int32(-4))|base.B2i32(v111 < int32(4)) != 0 {
					v196 = v112
					v197 = v111
					v202 = v117
				} else {
					v127 = v111 - int32(4)
					v131 = int32(base.Ui32(v127)>>(uint(int32(2))%32)) + int32(1)
					v133 = v131 & int32(3)
					if base.Ui32(int32(12)) <= base.Ui32(v127) {
						v138 = v112
						v139 = v111
						v142 = int32(0)
						v144 = v117
						for {
							v145 = int32(16)
							v146 = v139 - v145
							v148 = v138 + v145
							v149 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
							v155 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
							v158 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
							v164 = base.I64_extend_i32_u(base.I32_popcnt(v149)) + (base.I64_extend_i32_u(base.I32_popcnt(v152)) + (base.I64_extend_i32_u(base.I32_popcnt(v155)) + (v144 + base.I64_extend_i32_u(base.I32_popcnt(v158)))))
							v166 = v142 + int32(4)
							if v166 != v131&int32(2147483644) {
								v138 = v148
								v139 = v146
								v142 = v166
								v144 = v164
								continue
							} else {
								break
							}
							break
						}
						if v133 == int32(0) {
							v196 = v148
							v197 = v146
							v202 = v164
						} else {
							v170 = v148
							v171 = v146
							v176 = v164
							v178 = v170
							v179 = v171
							v180 = int32(0)
							v184 = v176
							for {
								v185 = int32(4)
								v186 = v179 - v185
								v188 = v178 + v185
								v189 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
								v192 = v184 + base.I64_extend_i32_u(base.I32_popcnt(v189))
								v194 = v180 + int32(1)
								if v194 != v133 {
									v178 = v188
									v179 = v186
									v180 = v194
									v184 = v192
									continue
								} else {
									break
								}
								break
							}
							v196 = v188
							v197 = v186
							v202 = v192
						}
					} else {
						v170 = v112
						v171 = v111
						v176 = v117
						v178 = v170
						v179 = v171
						v180 = int32(0)
						v184 = v176
						for {
							v185 = int32(4)
							v186 = v179 - v185
							v188 = v178 + v185
							v189 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
							v192 = v184 + base.I64_extend_i32_u(base.I32_popcnt(v189))
							v194 = v180 + int32(1)
							if v194 != v133 {
								v178 = v188
								v179 = v186
								v180 = v194
								v184 = v192
								continue
							} else {
								break
							}
							break
						}
						v196 = v188
						v197 = v186
						v202 = v192
					}
				}
				if v197 == int32(0) {
					v265 = v202
				} else {
					v206 = v197 & int32(3)
					if v206 == int32(0) {
						v227 = v196
						v229 = v197
						v233 = v202
					} else {
						v210 = v196
						v212 = v197
						v214 = int32(0)
						v216 = v202
						for {
							v217 = int32(1)
							v218 = v210 + v217
							v220 = v212 - v217
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
							v222 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v221)+uint32(_c_F_bytea_bit_count[0]))))
							v223 = v216 + v222
							v225 = v214 + v217
							if v225 != v206 {
								v210 = v218
								v212 = v220
								v214 = v225
								v216 = v223
								continue
							} else {
								break
							}
							break
						}
						v227 = v218
						v229 = v220
						v233 = v223
					}
					if base.Ui32(v197) < base.Ui32(int32(4)) {
						v265 = v233
					} else {
						v236 = v227
						v238 = v229
						v242 = v233
						for {
							v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+3)))
							v244 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v243)+uint32(_c_F_bytea_bit_count[0]))))
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+2)))
							v246 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v245)+uint32(_c_F_bytea_bit_count[0]))))
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)))
							v248 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_bytea_bit_count[0]))))
							v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
							v250 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_bytea_bit_count[0]))))
							v254 = v244 + (v246 + (v248 + (v242 + v250)))
							v255 = int32(4)
							v258 = v238 - v255
							if v258 != 0 {
								v236 = v236 + v255
								v238 = v258
								v242 = v254
								continue
							} else {
								break
							}
							break
						}
						v265 = v254
					}
				}
				v271 = v265
				v272 = F_Int64GetDatum(m, v271)
				mBase = m.M
				v273 = m.ExcPending
				if v273 != 0 {
					return int32(0)
				} else {
					return v272
				}
			} else {
				if v19 == int32(18) {
					v30 = int32(16)
				} else {
					v30 = int32(0)
				}
				v41 = v30
				if v16 != 0 {
					v44 = v13
				} else {
					v44 = v8 + int32(4)
				}
				if int32(3) < v41 {
					v111 = v41
					v112 = v44
					v117 = int64(0)
					if base.B2i32(v112 != (v112+int32(3))&int32(-4))|base.B2i32(v111 < int32(4)) != 0 {
						v196 = v112
						v197 = v111
						v202 = v117
					} else {
						v127 = v111 - int32(4)
						v131 = int32(base.Ui32(v127)>>(uint(int32(2))%32)) + int32(1)
						v133 = v131 & int32(3)
						if base.Ui32(int32(12)) <= base.Ui32(v127) {
							v138 = v112
							v139 = v111
							v142 = int32(0)
							v144 = v117
							for {
								v145 = int32(16)
								v146 = v139 - v145
								v148 = v138 + v145
								v149 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
								v152 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
								v155 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
								v158 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
								v164 = base.I64_extend_i32_u(base.I32_popcnt(v149)) + (base.I64_extend_i32_u(base.I32_popcnt(v152)) + (base.I64_extend_i32_u(base.I32_popcnt(v155)) + (v144 + base.I64_extend_i32_u(base.I32_popcnt(v158)))))
								v166 = v142 + int32(4)
								if v166 != v131&int32(2147483644) {
									v138 = v148
									v139 = v146
									v142 = v166
									v144 = v164
									continue
								} else {
									break
								}
								break
							}
							if v133 == int32(0) {
								v196 = v148
								v197 = v146
								v202 = v164
							} else {
								v170 = v148
								v171 = v146
								v176 = v164
								v178 = v170
								v179 = v171
								v180 = int32(0)
								v184 = v176
								for {
									v185 = int32(4)
									v186 = v179 - v185
									v188 = v178 + v185
									v189 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
									v192 = v184 + base.I64_extend_i32_u(base.I32_popcnt(v189))
									v194 = v180 + int32(1)
									if v194 != v133 {
										v178 = v188
										v179 = v186
										v180 = v194
										v184 = v192
										continue
									} else {
										break
									}
									break
								}
								v196 = v188
								v197 = v186
								v202 = v192
							}
						} else {
							v170 = v112
							v171 = v111
							v176 = v117
							v178 = v170
							v179 = v171
							v180 = int32(0)
							v184 = v176
							for {
								v185 = int32(4)
								v186 = v179 - v185
								v188 = v178 + v185
								v189 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
								v192 = v184 + base.I64_extend_i32_u(base.I32_popcnt(v189))
								v194 = v180 + int32(1)
								if v194 != v133 {
									v178 = v188
									v179 = v186
									v180 = v194
									v184 = v192
									continue
								} else {
									break
								}
								break
							}
							v196 = v188
							v197 = v186
							v202 = v192
						}
					}
					if v197 == int32(0) {
						v265 = v202
					} else {
						v206 = v197 & int32(3)
						if v206 == int32(0) {
							v227 = v196
							v229 = v197
							v233 = v202
						} else {
							v210 = v196
							v212 = v197
							v214 = int32(0)
							v216 = v202
							for {
								v217 = int32(1)
								v218 = v210 + v217
								v220 = v212 - v217
								v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
								v222 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v221)+uint32(_c_F_bytea_bit_count[0]))))
								v223 = v216 + v222
								v225 = v214 + v217
								if v225 != v206 {
									v210 = v218
									v212 = v220
									v214 = v225
									v216 = v223
									continue
								} else {
									break
								}
								break
							}
							v227 = v218
							v229 = v220
							v233 = v223
						}
						if base.Ui32(v197) < base.Ui32(int32(4)) {
							v265 = v233
						} else {
							v236 = v227
							v238 = v229
							v242 = v233
							for {
								v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+3)))
								v244 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v243)+uint32(_c_F_bytea_bit_count[0]))))
								v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+2)))
								v246 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v245)+uint32(_c_F_bytea_bit_count[0]))))
								v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)))
								v248 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_bytea_bit_count[0]))))
								v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
								v250 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_bytea_bit_count[0]))))
								v254 = v244 + (v246 + (v248 + (v242 + v250)))
								v255 = int32(4)
								v258 = v238 - v255
								if v258 != 0 {
									v236 = v236 + v255
									v238 = v258
									v242 = v254
									continue
								} else {
									break
								}
								break
							}
							v265 = v254
						}
					}
					v271 = v265
					v272 = F_Int64GetDatum(m, v271)
					mBase = m.M
					v273 = m.ExcPending
					if v273 != 0 {
						return int32(0)
					} else {
						return v272
					}
				} else {
					if v41 == int32(0) {
						v50 = F_Int64GetDatum(m, int64(0))
						mBase = m.M
						v51 = m.ExcPending
						if v51 != 0 {
							return int32(0)
						} else {
							return v50
						}
					} else {
						v54 = v41 & int32(3)
						if base.Ui32(int32(4)) <= base.Ui32(v41) {
							v60 = v44
							v63 = int32(0)
							v65 = v6
							for {
								v66 = int32(4)
								v67 = v60 + v66
								v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+3)))
								v69 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_bytea_bit_count[0]))))
								v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
								v71 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_bytea_bit_count[0]))))
								v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
								v73 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_bytea_bit_count[0]))))
								v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
								v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_bytea_bit_count[0]))))
								v79 = v69 + (v71 + (v73 + (v65 + v75)))
								v81 = v63 + v66
								if v81 != v41&int32(-4) {
									v60 = v67
									v63 = v81
									v65 = v79
									continue
								} else {
									break
								}
								break
							}
							if v54 == int32(0) {
								v271 = v79
							} else {
								v85 = v67
								v90 = v79
								v92 = v85
								v93 = int32(0)
								v97 = v90
								for {
									v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
									v99 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_bytea_bit_count[0]))))
									v100 = v97 + v99
									v101 = int32(1)
									v104 = v93 + v101
									if v104 != v54 {
										v92 = v92 + v101
										v93 = v104
										v97 = v100
										continue
									} else {
										break
									}
									break
								}
								v271 = v100
							}
						} else {
							v85 = v44
							v90 = v6
							v92 = v85
							v93 = int32(0)
							v97 = v90
							for {
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
								v99 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_bytea_bit_count[0]))))
								v100 = v97 + v99
								v101 = int32(1)
								v104 = v93 + v101
								if v104 != v54 {
									v92 = v92 + v101
									v93 = v104
									v97 = v100
									continue
								} else {
									break
								}
								break
							}
							v271 = v100
						}
						v272 = F_Int64GetDatum(m, v271)
						mBase = m.M
						v273 = m.ExcPending
						if v273 != 0 {
							return int32(0)
						} else {
							return v272
						}
					}
				}
			}
		} else {
			v31 = int32(1)
			if v16 != 0 {
				v41 = int32(base.Ui32(v14)>>(uint(v31)%32)) - v31
			} else {
				v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v41 = int32(base.Ui32(v35)>>(uint(int32(2))%32)) - int32(4)
			}
			if v16 != 0 {
				v44 = v13
			} else {
				v44 = v8 + int32(4)
			}
			if int32(3) < v41 {
				v111 = v41
				v112 = v44
				v117 = int64(0)
				if base.B2i32(v112 != (v112+int32(3))&int32(-4))|base.B2i32(v111 < int32(4)) != 0 {
					v196 = v112
					v197 = v111
					v202 = v117
				} else {
					v127 = v111 - int32(4)
					v131 = int32(base.Ui32(v127)>>(uint(int32(2))%32)) + int32(1)
					v133 = v131 & int32(3)
					if base.Ui32(int32(12)) <= base.Ui32(v127) {
						v138 = v112
						v139 = v111
						v142 = int32(0)
						v144 = v117
						for {
							v145 = int32(16)
							v146 = v139 - v145
							v148 = v138 + v145
							v149 = *(*int32)(unsafe.Add(mBase, uint32(v138)+12))
							v152 = *(*int32)(unsafe.Add(mBase, uint32(v138)+8))
							v155 = *(*int32)(unsafe.Add(mBase, uint32(v138)+4))
							v158 = *(*int32)(unsafe.Add(mBase, uint32(v138)))
							v164 = base.I64_extend_i32_u(base.I32_popcnt(v149)) + (base.I64_extend_i32_u(base.I32_popcnt(v152)) + (base.I64_extend_i32_u(base.I32_popcnt(v155)) + (v144 + base.I64_extend_i32_u(base.I32_popcnt(v158)))))
							v166 = v142 + int32(4)
							if v166 != v131&int32(2147483644) {
								v138 = v148
								v139 = v146
								v142 = v166
								v144 = v164
								continue
							} else {
								break
							}
							break
						}
						if v133 == int32(0) {
							v196 = v148
							v197 = v146
							v202 = v164
						} else {
							v170 = v148
							v171 = v146
							v176 = v164
							v178 = v170
							v179 = v171
							v180 = int32(0)
							v184 = v176
							for {
								v185 = int32(4)
								v186 = v179 - v185
								v188 = v178 + v185
								v189 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
								v192 = v184 + base.I64_extend_i32_u(base.I32_popcnt(v189))
								v194 = v180 + int32(1)
								if v194 != v133 {
									v178 = v188
									v179 = v186
									v180 = v194
									v184 = v192
									continue
								} else {
									break
								}
								break
							}
							v196 = v188
							v197 = v186
							v202 = v192
						}
					} else {
						v170 = v112
						v171 = v111
						v176 = v117
						v178 = v170
						v179 = v171
						v180 = int32(0)
						v184 = v176
						for {
							v185 = int32(4)
							v186 = v179 - v185
							v188 = v178 + v185
							v189 = *(*int32)(unsafe.Add(mBase, uint32(v178)))
							v192 = v184 + base.I64_extend_i32_u(base.I32_popcnt(v189))
							v194 = v180 + int32(1)
							if v194 != v133 {
								v178 = v188
								v179 = v186
								v180 = v194
								v184 = v192
								continue
							} else {
								break
							}
							break
						}
						v196 = v188
						v197 = v186
						v202 = v192
					}
				}
				if v197 == int32(0) {
					v265 = v202
				} else {
					v206 = v197 & int32(3)
					if v206 == int32(0) {
						v227 = v196
						v229 = v197
						v233 = v202
					} else {
						v210 = v196
						v212 = v197
						v214 = int32(0)
						v216 = v202
						for {
							v217 = int32(1)
							v218 = v210 + v217
							v220 = v212 - v217
							v221 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v210))))
							v222 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v221)+uint32(_c_F_bytea_bit_count[0]))))
							v223 = v216 + v222
							v225 = v214 + v217
							if v225 != v206 {
								v210 = v218
								v212 = v220
								v214 = v225
								v216 = v223
								continue
							} else {
								break
							}
							break
						}
						v227 = v218
						v229 = v220
						v233 = v223
					}
					if base.Ui32(v197) < base.Ui32(int32(4)) {
						v265 = v233
					} else {
						v236 = v227
						v238 = v229
						v242 = v233
						for {
							v243 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+3)))
							v244 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v243)+uint32(_c_F_bytea_bit_count[0]))))
							v245 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+2)))
							v246 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v245)+uint32(_c_F_bytea_bit_count[0]))))
							v247 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236)+1)))
							v248 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v247)+uint32(_c_F_bytea_bit_count[0]))))
							v249 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v236))))
							v250 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v249)+uint32(_c_F_bytea_bit_count[0]))))
							v254 = v244 + (v246 + (v248 + (v242 + v250)))
							v255 = int32(4)
							v258 = v238 - v255
							if v258 != 0 {
								v236 = v236 + v255
								v238 = v258
								v242 = v254
								continue
							} else {
								break
							}
							break
						}
						v265 = v254
					}
				}
				v271 = v265
				v272 = F_Int64GetDatum(m, v271)
				mBase = m.M
				v273 = m.ExcPending
				if v273 != 0 {
					return int32(0)
				} else {
					return v272
				}
			} else {
				if v41 == int32(0) {
					v50 = F_Int64GetDatum(m, int64(0))
					mBase = m.M
					v51 = m.ExcPending
					if v51 != 0 {
						return int32(0)
					} else {
						return v50
					}
				} else {
					v54 = v41 & int32(3)
					if base.Ui32(int32(4)) <= base.Ui32(v41) {
						v60 = v44
						v63 = int32(0)
						v65 = v6
						for {
							v66 = int32(4)
							v67 = v60 + v66
							v68 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+3)))
							v69 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v68)+uint32(_c_F_bytea_bit_count[0]))))
							v70 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+2)))
							v71 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v70)+uint32(_c_F_bytea_bit_count[0]))))
							v72 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60)+1)))
							v73 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v72)+uint32(_c_F_bytea_bit_count[0]))))
							v74 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v60))))
							v75 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v74)+uint32(_c_F_bytea_bit_count[0]))))
							v79 = v69 + (v71 + (v73 + (v65 + v75)))
							v81 = v63 + v66
							if v81 != v41&int32(-4) {
								v60 = v67
								v63 = v81
								v65 = v79
								continue
							} else {
								break
							}
							break
						}
						if v54 == int32(0) {
							v271 = v79
						} else {
							v85 = v67
							v90 = v79
							v92 = v85
							v93 = int32(0)
							v97 = v90
							for {
								v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
								v99 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_bytea_bit_count[0]))))
								v100 = v97 + v99
								v101 = int32(1)
								v104 = v93 + v101
								if v104 != v54 {
									v92 = v92 + v101
									v93 = v104
									v97 = v100
									continue
								} else {
									break
								}
								break
							}
							v271 = v100
						}
					} else {
						v85 = v44
						v90 = v6
						v92 = v85
						v93 = int32(0)
						v97 = v90
						for {
							v98 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v92))))
							v99 = int64(*(*uint8)(unsafe.Add(mBase, uint32(v98)+uint32(_c_F_bytea_bit_count[0]))))
							v100 = v97 + v99
							v101 = int32(1)
							v104 = v93 + v101
							if v104 != v54 {
								v92 = v92 + v101
								v93 = v104
								v97 = v100
								continue
							} else {
								break
							}
							break
						}
						v271 = v100
					}
					v272 = F_Int64GetDatum(m, v271)
					mBase = m.M
					v273 = m.ExcPending
					if v273 != 0 {
						return int32(0)
					} else {
						return v272
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
	v11 = F_detoast_attr_slice(m, v2, v7-int32(1), int32(-1))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		return v11
	}
}
