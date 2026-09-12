package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_cube_a_f8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
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
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v82 float64
	_ = v82
	var v85 int32
	_ = v85
	var v88 float64
	_ = v88
	var v91 int32
	_ = v91
	var v94 float64
	_ = v94
	var v97 int32
	_ = v97
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v120 int32
	_ = v120
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v135 float64
	_ = v135
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v169 int32
	_ = v169
	var v176 int32
	_ = v176
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v189 int32
	_ = v189
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	v2 = int32(0)
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = F_array_contains_nulls(m, v17)
		mBase = m.M
		v22 = m.ExcPending
		if v22 != 0 {
			return int32(0)
		} else {
			if v21 == int32(0) {
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
				v28 = F_ArrayGetNItems(m, v25, v17+int32(16))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if int32(101) <= v28 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v180 = m.ExcPending
						if v180 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v183 = m.ExcPending
							if v183 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(327886), int32(0))
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(100)
									F_errdetail(m, int32(592326), v14)
									mBase = m.M
									v195 = m.ExcPending
									if v195 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(499851), int32(230), int32(556030))
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
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
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+8))
						if v32 == int32(0) {
							v35 = *(*int32)(unsafe.Add(mBase, uint32(v17)+4))
							v42 = (v35<<(uint(int32(3))%32) + int32(23)) & int32(-8)
						} else {
							v42 = v32
						}
						v46 = v28<<(uint(int32(3))%32) + int32(8)
						v47 = F_palloc0(m, v46)
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v47)+4)) = v28 | int32(-2147483648)
							*(*int32)(unsafe.Add(mBase, uint32(v47))) = v46 << (uint(int32(2)) % 32)
							if v28 <= int32(0) {
							} else {
								v57 = v17 + v42
								v59 = v28 & int32(3)
								v61 = v47 + int32(8)
								v62 = int32(0)
								if base.Ui32(int32(4)) <= base.Ui32(v28) {
									v67 = v62
									v76 = v2
									for {
										v79 = v67 << (uint(int32(3)) % 32)
										v82 = *(*float64)(unsafe.Add(mBase, uint32(v79+v57)))
										*(*float64)(unsafe.Add(mBase, uint32(v61+v79))) = v82
										v85 = v79 | int32(8)
										v88 = *(*float64)(unsafe.Add(mBase, uint32(v57+v85)))
										*(*float64)(unsafe.Add(mBase, uint32(v61+v85))) = v88
										v91 = v79 | int32(16)
										v94 = *(*float64)(unsafe.Add(mBase, uint32(v57+v91)))
										*(*float64)(unsafe.Add(mBase, uint32(v61+v91))) = v94
										v97 = v79 | int32(24)
										v100 = *(*float64)(unsafe.Add(mBase, uint32(v97+v57)))
										*(*float64)(unsafe.Add(mBase, uint32(v61+v97))) = v100
										v102 = int32(4)
										v103 = v67 + v102
										v105 = v76 + v102
										if v105 != v28&int32(2147483644) {
											v67 = v103
											v76 = v105
											continue
										} else {
											break
										}
										break
									}
									v107 = v103
								} else {
									v107 = v62
								}
								if v59 == int32(0) {
								} else {
									v120 = v107
									v128 = v2
									for {
										v132 = v120 << (uint(int32(3)) % 32)
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v132+v57)))
										*(*float64)(unsafe.Add(mBase, uint32(v61+v132))) = v135
										v137 = int32(1)
										v140 = v128 + v137
										if v140 != v59 {
											v120 = v120 + v137
											v128 = v140
											continue
										} else {
											break
										}
										break
									}
								}
							}
							m.G0 = v14 + int32(16)
							return v47
						}
					}
				}
			} else {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v160 = m.ExcPending
				if v160 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(352845954))
					mBase = m.M
					v163 = m.ExcPending
					if v163 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(174986), int32(0))
						mBase = m.M
						v169 = m.ExcPending
						if v169 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499851), int32(222), int32(556030))
							mBase = m.M
							v176 = m.ExcPending
							if v176 != 0 {
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
func F_cube_c_f8(m *base.Module, l0 int32) int32 {
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
	var v28 int32
	_ = v28
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 float64
	_ = v34
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
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
	var v68 int32
	_ = v68
	var v75 int32
	_ = v75
	var v84 int32
	_ = v84
	var v87 float64
	_ = v87
	var v90 int32
	_ = v90
	var v93 float64
	_ = v93
	var v96 int32
	_ = v96
	var v99 float64
	_ = v99
	var v102 int32
	_ = v102
	var v105 float64
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v131 int32
	_ = v131
	var v137 int32
	_ = v137
	var v147 int32
	_ = v147
	var v150 float64
	_ = v150
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v229 int32
	_ = v229
	var v230 int32
	_ = v230
	var v233 float64
	_ = v233
	var v243 float64
	_ = v243
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 float64
	_ = v251
	var v261 float64
	_ = v261
	var v263 int32
	_ = v263
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v286 int32
	_ = v286
	var v287 int32
	_ = v287
	var v290 float64
	_ = v290
	var v300 float64
	_ = v300
	var v309 int32
	_ = v309
	var v329 int32
	_ = v329
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v348 int32
	_ = v348
	var v356 int32
	_ = v356
	var v359 int32
	_ = v359
	var v365 int32
	_ = v365
	var v371 int32
	_ = v371
	var v378 int32
	_ = v378
	v2 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(16)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = F_pg_detoast_datum(m, v21)
	mBase = m.M
	v25 = m.ExcPending
	if v25 != 0 {
		return int32(0)
	} else {
		v26 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
		v28 = v26 & int32(2147483647)
		if base.Ui32(v28) < base.Ui32(int32(100)) {
			v32 = v28 + int32(1)
			v33 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			v34 = *(*float64)(unsafe.Add(mBase, uint32(v33)))
			if v26 < int32(0) {
				v40 = v32<<(uint(int32(3))%32) + int32(8)
				v41 = F_palloc0(m, v40)
				mBase = m.M
				v42 = m.ExcPending
				if v42 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v41))) = v40 << (uint(int32(2)) % 32)
					v46 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					v48 = v46 + int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v41)+4)) = v48 | int32(-2147483648)
					v52 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					v54 = v52 & int32(2147483647)
					if v54 == int32(0) {
					} else {
						v58 = v52 & int32(3)
						v59 = int32(8)
						v60 = v41 + v59
						v62 = v22 + v59
						v63 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v54) {
							v68 = v63
							v75 = v2
							for {
								v84 = v68 << (uint(int32(3)) % 32)
								v87 = *(*float64)(unsafe.Add(mBase, uint32(v84+v62)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v84))) = v87
								v90 = v84 | int32(8)
								v93 = *(*float64)(unsafe.Add(mBase, uint32(v62+v90)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v90))) = v93
								v96 = v84 | int32(16)
								v99 = *(*float64)(unsafe.Add(mBase, uint32(v62+v96)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v96))) = v99
								v102 = v84 | int32(24)
								v105 = *(*float64)(unsafe.Add(mBase, uint32(v102+v62)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v102))) = v105
								v107 = int32(4)
								v108 = v68 + v107
								v110 = v75 + v107
								if v110 != v54-v58 {
									v68 = v108
									v75 = v110
									continue
								} else {
									break
								}
								break
							}
							v113 = v108
						} else {
							v113 = v63
						}
						if v58 == int32(0) {
						} else {
							v131 = v113
							v137 = v2
							for {
								v147 = v131 << (uint(int32(3)) % 32)
								v150 = *(*float64)(unsafe.Add(mBase, uint32(v147+v62)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v147))) = v150
								v152 = int32(1)
								v155 = v137 + v152
								if v155 != v58 {
									v131 = v131 + v152
									v137 = v155
									continue
								} else {
									break
								}
								break
							}
						}
					}
					v329 = v41
					v340 = v48 & int32(2147483647)
					*(*float64)(unsafe.Add(mBase, uint32(v340<<(uint(int32(3))%32)+v329))) = v34
					v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v345 != v22 {
						F_pfree(m, v22)
						mBase = m.M
						v348 = m.ExcPending
						if v348 != 0 {
							return int32(0)
						} else {
							m.G0 = v19 + int32(16)
							return v329
						}
					} else {
						m.G0 = v19 + int32(16)
						return v329
					}
				}
			} else {
				v178 = v32<<(uint(int32(4))%32) | int32(8)
				v179 = F_palloc0(m, v178)
				mBase = m.M
				v180 = m.ExcPending
				if v180 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v179))) = v178 << (uint(int32(2)) % 32)
					v184 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					v185 = int32(2147483647)
					v188 = v184&v185 + int32(1)
					v189 = *(*int32)(unsafe.Add(mBase, uint32(v179)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v179)+4)) = v188 | v189&int32(-2147483648)
					v194 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					v196 = v194 & v185
					if v196 == int32(0) {
						v309 = v188 & int32(2147483647)
					} else {
						v201 = int32(1)
						v202 = v194 & v201
						v204 = v188 & int32(2147483647)
						v205 = int32(8)
						v206 = v179 + v205
						v208 = v22 + v205
						v209 = int32(0)
						if v196 != v201 {
							v214 = v209
							v222 = v2
							for {
								v229 = int32(3)
								v230 = v214 << (uint(v229) % 32)
								v233 = *(*float64)(unsafe.Add(mBase, uint32(v230+v208)))
								*(*float64)(unsafe.Add(mBase, uint32(v206+v230))) = v233
								v243 = *(*float64)(unsafe.Add(mBase, uint32(v208+(v214+v196)<<(uint(v229)%32))))
								*(*float64)(unsafe.Add(mBase, uint32(v206+(v214+v204)<<(uint(v229)%32)))) = v243
								v246 = v214 | int32(1)
								v248 = v246 << (uint(v229) % 32)
								v251 = *(*float64)(unsafe.Add(mBase, uint32(v208+v248)))
								*(*float64)(unsafe.Add(mBase, uint32(v206+v248))) = v251
								v261 = *(*float64)(unsafe.Add(mBase, uint32(v208+(v246+v196)<<(uint(v229)%32))))
								*(*float64)(unsafe.Add(mBase, uint32(v206+(v246+v204)<<(uint(v229)%32)))) = v261
								v263 = int32(2)
								v264 = v214 + v263
								v266 = v222 + v263
								if v266 != v196-v202 {
									v214 = v264
									v222 = v266
									continue
								} else {
									break
								}
								break
							}
							v269 = v264
						} else {
							v269 = v209
						}
						if v202 == int32(0) {
							v309 = v204
						} else {
							v286 = int32(3)
							v287 = v269 << (uint(v286) % 32)
							v290 = *(*float64)(unsafe.Add(mBase, uint32(v287+v208)))
							*(*float64)(unsafe.Add(mBase, uint32(v206+v287))) = v290
							v300 = *(*float64)(unsafe.Add(mBase, uint32(v208+(v269+v196)<<(uint(v286)%32))))
							*(*float64)(unsafe.Add(mBase, uint32(v206+(v269+v204)<<(uint(v286)%32)))) = v300
							v309 = v204
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v309<<(uint(int32(3))%32)+v179))) = v34
					v329 = v179
					v340 = v188 << (uint(int32(1)) % 32)
					*(*float64)(unsafe.Add(mBase, uint32(v340<<(uint(int32(3))%32)+v329))) = v34
					v345 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v345 != v22 {
						F_pfree(m, v22)
						mBase = m.M
						v348 = m.ExcPending
						if v348 != 0 {
							return int32(0)
						} else {
							m.G0 = v19 + int32(16)
							return v329
						}
					} else {
						m.G0 = v19 + int32(16)
						return v329
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v356 = m.ExcPending
			if v356 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v359 = m.ExcPending
				if v359 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(420150), int32(0))
					mBase = m.M
					v365 = m.ExcPending
					if v365 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(100)
						F_errdetail(m, int32(592326), v19)
						mBase = m.M
						v371 = m.ExcPending
						if v371 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(499851), int32(1835), int32(556011))
							mBase = m.M
							v378 = m.ExcPending
							if v378 != 0 {
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
func F_cube_cmp(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v76 int32
	_ = v76
	var v81 float64
	_ = v81
	var v83 float64
	_ = v83
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v133 float64
	_ = v133
	var v135 int32
	_ = v135
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v147 int32
	_ = v147
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v175 int32
	_ = v175
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 float64
	_ = v232
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v238 float64
	_ = v238
	var v240 float64
	_ = v240
	var v243 float64
	_ = v243
	var v250 int32
	_ = v250
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 float64
	_ = v286
	var v288 int32
	_ = v288
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v294 float64
	_ = v294
	var v297 float64
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 float64
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 float64
	_ = v347
	var v349 int32
	_ = v349
	var v350 float64
	_ = v350
	var v361 float64
	_ = v361
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v395 float64
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 float64
	_ = v406
	var v408 int32
	_ = v408
	var v409 float64
	_ = v409
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v423 float64
	_ = v423
	var v430 int32
	_ = v430
	var v465 int32
	_ = v465
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v488 != v6 {
		goto L123
	} else {
		goto L124
	}
L5:
	;
	v487 = v465
	goto L4
L6:
	;
	v465 = int32(1)
	goto L5
L7:
	;
	v36 = v31
	goto L9
L8:
	;
	v36 = v34
	goto L9
L9:
	;
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = int32(8)
	v38 = v11 + v37
	v40 = v6 + v37
	v48 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L47
	} else {
		goto L48
	}
L13:
	;
	v60 = v48 << (uint(int32(3)) % 32)
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v40+v60)))
	v64 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v72 = v62
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v107 = int32(8)
	v108 = v11 + v107
	v110 = v6 + v107
	v119 = int32(0)
	goto L30
L15:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v38+v60)))
	v76 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v83 = v74
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v69) != 0 {
		v72 = v62
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v72 = v69
	goto L15
L18:
	;
	if base.F64_gt(v72, v83) != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v81) != 0 {
		v83 = v74
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v83 = v81
	goto L18
L21:
	;
	if v29 < int32(0) {
		v92 = v62
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v32 < int32(0) {
		v100 = v74
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v90) != 0 {
		v92 = v62
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v92 = v90
	goto L22
L25:
	;
	v102 = int32(-1)
	if base.F64_lt(v92, v100) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L28
	}
L26:
	;
	v98 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v98) != 0 {
		v100 = v74
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v100 = v98
	goto L25
L28:
	;
	v105 = v48 + int32(1)
	if v105 != v36 {
		v48 = v105
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	v131 = v119 << (uint(int32(3)) % 32)
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v110+v131)))
	v135 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v143 = v133
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L12
L32:
	;
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v108+v131)))
	v147 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v154 = v145
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v140) != 0 {
		v143 = v133
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v143 = v140
	goto L32
L35:
	;
	if base.F64_gt(v143, v154) != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v152) != 0 {
		v154 = v145
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v154 = v152
	goto L35
L38:
	;
	if v29 < int32(0) {
		v163 = v133
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v32 < int32(0) {
		v171 = v145
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v161) != 0 {
		v163 = v133
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v163 = v161
	goto L39
L42:
	;
	if base.F64_lt(v163, v171) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v169) != 0 {
		v171 = v145
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v171 = v169
	goto L42
L45:
	;
	v175 = v119 + int32(1)
	if v175 != v36 {
		v119 = v175
		goto L30
	} else {
		goto L46
	}
L46:
	;
	goto L31
L47:
	;
	v197 = v6 + int32(8)
	v207 = v36
	goto L50
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v34) <= base.Ui32(v31) {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v220 = v197 + v207<<(uint(int32(3))%32)
	v221 = *(*float64)(unsafe.Add(mBase, uint32(v220)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v266 = v36
	goto L68
L52:
	;
	if base.F64_lt(v243, float64(0)) != 0 {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	v225 = int32(3)
	v227 = v197 + (v207+v29)<<(uint(v225)%32)
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v207+v31)<<(uint(v225)%32))))
	if base.F64_lt(v221, v232) != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.F64_gt(v221, float64(0)) != 0 {
		goto L6
	} else {
		goto L63
	}
L56:
	;
	v234 = v220
	goto L58
L57:
	;
	v234 = v227
	goto L58
L58:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v234)))
	if base.F64_gt(v235, float64(0)) != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v227)))
	if base.F64_lt(v221, v238) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v240 = v221
	goto L62
L61:
	;
	v240 = v238
	goto L62
L62:
	;
	v243 = v240
	goto L52
L63:
	;
	v243 = v221
	goto L52
L64:
	;
	v487 = int32(-1)
	goto L4
L65:
	;
	goto L66
L66:
	;
	v250 = v207 + int32(1)
	if v250 != v31 {
		v207 = v250
		goto L50
	} else {
		goto L67
	}
L67:
	;
	goto L51
L68:
	;
	v274 = v197 + v266<<(uint(int32(3))%32)
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v274)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v487 = int32(-1)
	goto L4
L70:
	;
	if base.F64_lt(v297, float64(0)) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v279 = int32(3)
	v281 = v197 + (v29+v266)<<(uint(v279)%32)
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v266+v31)<<(uint(v279)%32))))
	if base.F64_gt(v275, v286) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if base.F64_gt(v275, float64(0)) != 0 {
		goto L6
	} else {
		goto L81
	}
L74:
	;
	v288 = v274
	goto L76
L75:
	;
	v288 = v281
	goto L76
L76:
	;
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v288)))
	if base.F64_gt(v289, float64(0)) != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v292 = *(*float64)(unsafe.Add(mBase, uint32(v281)))
	if base.F64_gt(v275, v292) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v294 = v275
	goto L80
L79:
	;
	v294 = v292
	goto L80
L80:
	;
	v297 = v294
	goto L70
L81:
	;
	v297 = v275
	goto L70
L82:
	;
	v304 = int32(1)
	v306 = v266 + v304
	if v306 == v31 {
		v465 = v304
		goto L5
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L69
L85:
	;
	v266 = v306
	goto L68
L86:
	;
	v487 = int32(0)
	goto L4
L87:
	;
	goto L88
L88:
	;
	v312 = v11 + int32(8)
	v328 = v31
	goto L89
L89:
	;
	v335 = v312 + v328<<(uint(int32(3))%32)
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v335)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v386 = v36
	goto L106
L91:
	;
	if base.F64_lt(v364, float64(0)) != 0 {
		goto L6
	} else {
		goto L104
	}
L92:
	;
	v361 = *(*float64)(unsafe.Add(mBase, uint32(v342)))
	if base.F64_lt(v336, v361) != 0 {
		goto L101
	} else {
		goto L102
	}
L93:
	;
	v340 = int32(3)
	v342 = v312 + (v32+v328)<<(uint(v340)%32)
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v328+v34)<<(uint(v340)%32))))
	if base.F64_lt(v336, v347) != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if base.F64_gt(v336, float64(0)) == int32(0) {
		v364 = v336
		goto L91
	} else {
		goto L100
	}
L96:
	;
	v349 = v335
	goto L98
L97:
	;
	v349 = v342
	goto L98
L98:
	;
	v350 = *(*float64)(unsafe.Add(mBase, uint32(v349)))
	if base.F64_gt(v350, float64(0)) == int32(0) {
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v487 = int32(-1)
	goto L4
L100:
	;
	v487 = int32(-1)
	goto L4
L101:
	;
	v363 = v336
	goto L103
L102:
	;
	v363 = v361
	goto L103
L103:
	;
	v364 = v363
	goto L91
L104:
	;
	v370 = v328 + int32(1)
	if v370 != v34 {
		v328 = v370
		goto L89
	} else {
		goto L105
	}
L105:
	;
	goto L90
L106:
	;
	v394 = v312 + v386<<(uint(int32(3))%32)
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v394)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v465 = int32(-1)
	goto L5
L108:
	;
	if base.F64_lt(v423, float64(0)) != 0 {
		goto L6
	} else {
		goto L121
	}
L109:
	;
	v420 = *(*float64)(unsafe.Add(mBase, uint32(v401)))
	if base.F64_gt(v395, v420) != 0 {
		goto L118
	} else {
		goto L119
	}
L110:
	;
	v399 = int32(3)
	v401 = v312 + (v32+v386)<<(uint(v399)%32)
	v406 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v386+v34)<<(uint(v399)%32))))
	if base.F64_gt(v395, v406) != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if base.F64_gt(v395, float64(0)) == int32(0) {
		v423 = v395
		goto L108
	} else {
		goto L117
	}
L113:
	;
	v408 = v394
	goto L115
L114:
	;
	v408 = v401
	goto L115
L115:
	;
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v408)))
	if base.F64_gt(v409, float64(0)) == int32(0) {
		goto L109
	} else {
		goto L116
	}
L116:
	;
	v487 = int32(-1)
	goto L4
L117:
	;
	v487 = int32(-1)
	goto L4
L118:
	;
	v422 = v395
	goto L120
L119:
	;
	v422 = v420
	goto L120
L120:
	;
	v423 = v422
	goto L108
L121:
	;
	v430 = v386 + int32(1)
	if v34 != v430 {
		v386 = v430
		goto L106
	} else {
		goto L122
	}
L122:
	;
	goto L107
L123:
	;
	F_pfree(m, v6)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v492 != v11 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	F_pfree(m, v11)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	return v487
L130:
	;
	goto L129
}
func F_cube_f8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 float64
	_ = v4
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*float64)(unsafe.Add(mBase, uint32(v3)))
	v6 = F_palloc0(m, int32(16))
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		*(*float64)(unsafe.Add(mBase, uint32(v6)+8)) = v4
		*(*int64)(unsafe.Add(mBase, uint32(v6))) = int64(-9223372032559808448)
		return v6
	}
}
func F_cube_ge(m *base.Module, l0 int32) int32 {
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
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v48 int32
	_ = v48
	var v60 int32
	_ = v60
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v69 float64
	_ = v69
	var v72 float64
	_ = v72
	var v74 float64
	_ = v74
	var v76 int32
	_ = v76
	var v81 float64
	_ = v81
	var v83 float64
	_ = v83
	var v90 float64
	_ = v90
	var v92 float64
	_ = v92
	var v98 float64
	_ = v98
	var v100 float64
	_ = v100
	var v102 int32
	_ = v102
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int32
	_ = v119
	var v131 int32
	_ = v131
	var v133 float64
	_ = v133
	var v135 int32
	_ = v135
	var v140 float64
	_ = v140
	var v143 float64
	_ = v143
	var v145 float64
	_ = v145
	var v147 int32
	_ = v147
	var v152 float64
	_ = v152
	var v154 float64
	_ = v154
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v169 float64
	_ = v169
	var v171 float64
	_ = v171
	var v175 int32
	_ = v175
	var v197 int32
	_ = v197
	var v207 int32
	_ = v207
	var v220 int32
	_ = v220
	var v221 float64
	_ = v221
	var v225 int32
	_ = v225
	var v227 int32
	_ = v227
	var v232 float64
	_ = v232
	var v234 int32
	_ = v234
	var v235 float64
	_ = v235
	var v238 float64
	_ = v238
	var v240 float64
	_ = v240
	var v243 float64
	_ = v243
	var v250 int32
	_ = v250
	var v266 int32
	_ = v266
	var v274 int32
	_ = v274
	var v275 float64
	_ = v275
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v286 float64
	_ = v286
	var v288 int32
	_ = v288
	var v289 float64
	_ = v289
	var v292 float64
	_ = v292
	var v294 float64
	_ = v294
	var v297 float64
	_ = v297
	var v304 int32
	_ = v304
	var v306 int32
	_ = v306
	var v312 int32
	_ = v312
	var v328 int32
	_ = v328
	var v335 int32
	_ = v335
	var v336 float64
	_ = v336
	var v340 int32
	_ = v340
	var v342 int32
	_ = v342
	var v347 float64
	_ = v347
	var v349 int32
	_ = v349
	var v350 float64
	_ = v350
	var v361 float64
	_ = v361
	var v363 float64
	_ = v363
	var v364 float64
	_ = v364
	var v370 int32
	_ = v370
	var v386 int32
	_ = v386
	var v394 int32
	_ = v394
	var v395 float64
	_ = v395
	var v399 int32
	_ = v399
	var v401 int32
	_ = v401
	var v406 float64
	_ = v406
	var v408 int32
	_ = v408
	var v409 float64
	_ = v409
	var v420 float64
	_ = v420
	var v422 float64
	_ = v422
	var v423 float64
	_ = v423
	var v430 int32
	_ = v430
	var v465 int32
	_ = v465
	var v487 int32
	_ = v487
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = F_pg_detoast_datum(m, v5)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v29 = *(*int32)(unsafe.Add(mBase, uint32(v6)+4))
	v30 = int32(2147483647)
	v31 = v29 & v30
	v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
	v34 = v32 & v30
	if base.Ui32(v31) < base.Ui32(v34) {
		goto L7
	} else {
		goto L8
	}
L4:
	;
	v488 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v488 != v6 {
		goto L123
	} else {
		goto L124
	}
L5:
	;
	v487 = v465
	goto L4
L6:
	;
	v465 = int32(1)
	goto L5
L7:
	;
	v36 = v31
	goto L9
L8:
	;
	v36 = v34
	goto L9
L9:
	;
	if v36 != 0 {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v37 = int32(8)
	v38 = v11 + v37
	v40 = v6 + v37
	v48 = int32(0)
	goto L13
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L47
	} else {
		goto L48
	}
L13:
	;
	v60 = v48 << (uint(int32(3)) % 32)
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v40+v60)))
	v64 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v72 = v62
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v107 = int32(8)
	v108 = v11 + v107
	v110 = v6 + v107
	v119 = int32(0)
	goto L30
L15:
	;
	v74 = *(*float64)(unsafe.Add(mBase, uint32(v38+v60)))
	v76 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v83 = v74
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v69) != 0 {
		v72 = v62
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v72 = v69
	goto L15
L18:
	;
	if base.F64_gt(v72, v83) != 0 {
		goto L6
	} else {
		goto L21
	}
L19:
	;
	v81 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v81) != 0 {
		v83 = v74
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v83 = v81
	goto L18
L21:
	;
	if v29 < int32(0) {
		v92 = v62
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v32 < int32(0) {
		v100 = v74
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v90 = *(*float64)(unsafe.Add(mBase, uint32(v40+(v48+v29)<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v90) != 0 {
		v92 = v62
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v92 = v90
	goto L22
L25:
	;
	v102 = int32(-1)
	if base.F64_lt(v92, v100) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L28
	}
L26:
	;
	v98 = *(*float64)(unsafe.Add(mBase, uint32(v38+(v48+v32)<<(uint(int32(3))%32))))
	if base.F64_lt(v74, v98) != 0 {
		v100 = v74
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v100 = v98
	goto L25
L28:
	;
	v105 = v48 + int32(1)
	if v105 != v36 {
		v48 = v105
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	v131 = v119 << (uint(int32(3)) % 32)
	v133 = *(*float64)(unsafe.Add(mBase, uint32(v110+v131)))
	v135 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v143 = v133
		goto L32
	} else {
		goto L33
	}
L31:
	;
	goto L12
L32:
	;
	v145 = *(*float64)(unsafe.Add(mBase, uint32(v108+v131)))
	v147 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v154 = v145
		goto L35
	} else {
		goto L36
	}
L33:
	;
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v140) != 0 {
		v143 = v133
		goto L32
	} else {
		goto L34
	}
L34:
	;
	v143 = v140
	goto L32
L35:
	;
	if base.F64_gt(v143, v154) != 0 {
		goto L6
	} else {
		goto L38
	}
L36:
	;
	v152 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v152) != 0 {
		v154 = v145
		goto L35
	} else {
		goto L37
	}
L37:
	;
	v154 = v152
	goto L35
L38:
	;
	if v29 < int32(0) {
		v163 = v133
		goto L39
	} else {
		goto L40
	}
L39:
	;
	if v32 < int32(0) {
		v171 = v145
		goto L42
	} else {
		goto L43
	}
L40:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v110+(v119+v29)<<(uint(int32(3))%32))))
	if base.F64_gt(v133, v161) != 0 {
		v163 = v133
		goto L39
	} else {
		goto L41
	}
L41:
	;
	v163 = v161
	goto L39
L42:
	;
	if base.F64_lt(v163, v171) != 0 {
		v465 = v102
		goto L5
	} else {
		goto L45
	}
L43:
	;
	v169 = *(*float64)(unsafe.Add(mBase, uint32(v108+(v119+v32)<<(uint(int32(3))%32))))
	if base.F64_gt(v145, v169) != 0 {
		v171 = v145
		goto L42
	} else {
		goto L44
	}
L44:
	;
	v171 = v169
	goto L42
L45:
	;
	v175 = v119 + int32(1)
	if v175 != v36 {
		v119 = v175
		goto L30
	} else {
		goto L46
	}
L46:
	;
	goto L31
L47:
	;
	v197 = v6 + int32(8)
	v207 = v36
	goto L50
L48:
	;
	goto L49
L49:
	;
	if base.Ui32(v34) <= base.Ui32(v31) {
		goto L86
	} else {
		goto L87
	}
L50:
	;
	v220 = v197 + v207<<(uint(int32(3))%32)
	v221 = *(*float64)(unsafe.Add(mBase, uint32(v220)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L53
	} else {
		goto L54
	}
L51:
	;
	v266 = v36
	goto L68
L52:
	;
	if base.F64_lt(v243, float64(0)) != 0 {
		goto L64
	} else {
		goto L65
	}
L53:
	;
	v225 = int32(3)
	v227 = v197 + (v207+v29)<<(uint(v225)%32)
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v207+v31)<<(uint(v225)%32))))
	if base.F64_lt(v221, v232) != 0 {
		goto L56
	} else {
		goto L57
	}
L54:
	;
	goto L55
L55:
	;
	if base.F64_gt(v221, float64(0)) != 0 {
		goto L6
	} else {
		goto L63
	}
L56:
	;
	v234 = v220
	goto L58
L57:
	;
	v234 = v227
	goto L58
L58:
	;
	v235 = *(*float64)(unsafe.Add(mBase, uint32(v234)))
	if base.F64_gt(v235, float64(0)) != 0 {
		goto L6
	} else {
		goto L59
	}
L59:
	;
	v238 = *(*float64)(unsafe.Add(mBase, uint32(v227)))
	if base.F64_lt(v221, v238) != 0 {
		goto L60
	} else {
		goto L61
	}
L60:
	;
	v240 = v221
	goto L62
L61:
	;
	v240 = v238
	goto L62
L62:
	;
	v243 = v240
	goto L52
L63:
	;
	v243 = v221
	goto L52
L64:
	;
	v487 = int32(-1)
	goto L4
L65:
	;
	goto L66
L66:
	;
	v250 = v207 + int32(1)
	if v250 != v31 {
		v207 = v250
		goto L50
	} else {
		goto L67
	}
L67:
	;
	goto L51
L68:
	;
	v274 = v197 + v266<<(uint(int32(3))%32)
	v275 = *(*float64)(unsafe.Add(mBase, uint32(v274)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	v487 = int32(-1)
	goto L4
L70:
	;
	if base.F64_lt(v297, float64(0)) == int32(0) {
		goto L82
	} else {
		goto L83
	}
L71:
	;
	v279 = int32(3)
	v281 = v197 + (v29+v266)<<(uint(v279)%32)
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v197+(v266+v31)<<(uint(v279)%32))))
	if base.F64_gt(v275, v286) != 0 {
		goto L74
	} else {
		goto L75
	}
L72:
	;
	goto L73
L73:
	;
	if base.F64_gt(v275, float64(0)) != 0 {
		goto L6
	} else {
		goto L81
	}
L74:
	;
	v288 = v274
	goto L76
L75:
	;
	v288 = v281
	goto L76
L76:
	;
	v289 = *(*float64)(unsafe.Add(mBase, uint32(v288)))
	if base.F64_gt(v289, float64(0)) != 0 {
		goto L6
	} else {
		goto L77
	}
L77:
	;
	v292 = *(*float64)(unsafe.Add(mBase, uint32(v281)))
	if base.F64_gt(v275, v292) != 0 {
		goto L78
	} else {
		goto L79
	}
L78:
	;
	v294 = v275
	goto L80
L79:
	;
	v294 = v292
	goto L80
L80:
	;
	v297 = v294
	goto L70
L81:
	;
	v297 = v275
	goto L70
L82:
	;
	v304 = int32(1)
	v306 = v266 + v304
	if v306 == v31 {
		v465 = v304
		goto L5
	} else {
		goto L85
	}
L83:
	;
	goto L84
L84:
	;
	goto L69
L85:
	;
	v266 = v306
	goto L68
L86:
	;
	v487 = int32(0)
	goto L4
L87:
	;
	goto L88
L88:
	;
	v312 = v11 + int32(8)
	v328 = v31
	goto L89
L89:
	;
	v335 = v312 + v328<<(uint(int32(3))%32)
	v336 = *(*float64)(unsafe.Add(mBase, uint32(v335)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L93
	} else {
		goto L94
	}
L90:
	;
	v386 = v36
	goto L106
L91:
	;
	if base.F64_lt(v364, float64(0)) != 0 {
		goto L6
	} else {
		goto L104
	}
L92:
	;
	v361 = *(*float64)(unsafe.Add(mBase, uint32(v342)))
	if base.F64_lt(v336, v361) != 0 {
		goto L101
	} else {
		goto L102
	}
L93:
	;
	v340 = int32(3)
	v342 = v312 + (v32+v328)<<(uint(v340)%32)
	v347 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v328+v34)<<(uint(v340)%32))))
	if base.F64_lt(v336, v347) != 0 {
		goto L96
	} else {
		goto L97
	}
L94:
	;
	goto L95
L95:
	;
	if base.F64_gt(v336, float64(0)) == int32(0) {
		v364 = v336
		goto L91
	} else {
		goto L100
	}
L96:
	;
	v349 = v335
	goto L98
L97:
	;
	v349 = v342
	goto L98
L98:
	;
	v350 = *(*float64)(unsafe.Add(mBase, uint32(v349)))
	if base.F64_gt(v350, float64(0)) == int32(0) {
		goto L92
	} else {
		goto L99
	}
L99:
	;
	v487 = int32(-1)
	goto L4
L100:
	;
	v487 = int32(-1)
	goto L4
L101:
	;
	v363 = v336
	goto L103
L102:
	;
	v363 = v361
	goto L103
L103:
	;
	v364 = v363
	goto L91
L104:
	;
	v370 = v328 + int32(1)
	if v370 != v34 {
		v328 = v370
		goto L89
	} else {
		goto L105
	}
L105:
	;
	goto L90
L106:
	;
	v394 = v312 + v386<<(uint(int32(3))%32)
	v395 = *(*float64)(unsafe.Add(mBase, uint32(v394)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L110
	} else {
		goto L111
	}
L107:
	;
	v465 = int32(-1)
	goto L5
L108:
	;
	if base.F64_lt(v423, float64(0)) != 0 {
		goto L6
	} else {
		goto L121
	}
L109:
	;
	v420 = *(*float64)(unsafe.Add(mBase, uint32(v401)))
	if base.F64_gt(v395, v420) != 0 {
		goto L118
	} else {
		goto L119
	}
L110:
	;
	v399 = int32(3)
	v401 = v312 + (v32+v386)<<(uint(v399)%32)
	v406 = *(*float64)(unsafe.Add(mBase, uint32(v312+(v386+v34)<<(uint(v399)%32))))
	if base.F64_gt(v395, v406) != 0 {
		goto L113
	} else {
		goto L114
	}
L111:
	;
	goto L112
L112:
	;
	if base.F64_gt(v395, float64(0)) == int32(0) {
		v423 = v395
		goto L108
	} else {
		goto L117
	}
L113:
	;
	v408 = v394
	goto L115
L114:
	;
	v408 = v401
	goto L115
L115:
	;
	v409 = *(*float64)(unsafe.Add(mBase, uint32(v408)))
	if base.F64_gt(v409, float64(0)) == int32(0) {
		goto L109
	} else {
		goto L116
	}
L116:
	;
	v487 = int32(-1)
	goto L4
L117:
	;
	v487 = int32(-1)
	goto L4
L118:
	;
	v422 = v395
	goto L120
L119:
	;
	v422 = v420
	goto L120
L120:
	;
	v423 = v422
	goto L108
L121:
	;
	v430 = v386 + int32(1)
	if v34 != v430 {
		v386 = v430
		goto L106
	} else {
		goto L122
	}
L122:
	;
	goto L107
L123:
	;
	F_pfree(m, v6)
	mBase = m.M
	v491 = m.ExcPending
	if v491 != 0 {
		goto L1
	} else {
		goto L126
	}
L124:
	;
	goto L125
L125:
	;
	v492 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v492 != v11 {
		goto L127
	} else {
		goto L128
	}
L126:
	;
	goto L125
L127:
	;
	F_pfree(m, v11)
	mBase = m.M
	v495 = m.ExcPending
	if v495 != 0 {
		goto L1
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	return int32(base.Ui32(v487^int32(-1)) >> (uint(int32(31)) % 32))
L130:
	;
	goto L129
}
func F_cube_overlap_v0(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v55 int32
	_ = v55
	var v57 float64
	_ = v57
	var v59 int32
	_ = v59
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v69 float64
	_ = v69
	var v71 int32
	_ = v71
	var v76 float64
	_ = v76
	var v79 float64
	_ = v79
	var v80 int32
	_ = v80
	var v86 float64
	_ = v86
	var v88 float64
	_ = v88
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v100 int32
	_ = v100
	var v108 int32
	_ = v108
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v126 int32
	_ = v126
	var v141 int32
	_ = v141
	var v142 float64
	_ = v142
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v153 float64
	_ = v153
	var v155 int32
	_ = v155
	var v156 float64
	_ = v156
	var v159 float64
	_ = v159
	var v161 float64
	_ = v161
	var v164 float64
	_ = v164
	var v169 int32
	_ = v169
	var v171 int32
	_ = v171
	var v184 int32
	_ = v184
	v7 = int32(0)
	if l0 == v7 {
		v184 = v7
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v184
L2:
	;
	if l1 == int32(0) {
		v184 = v7
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v20 = int32(2147483647)
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v25 = base.B2i32(base.Ui32(v19&v20) < base.Ui32(v22&v20))
	if base.Ui32(v19&v20) < base.Ui32(v22&v20) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v26 = l1
	goto L6
L5:
	;
	v26 = l0
	goto L6
L6:
	;
	if base.Ui32(v19&v20) < base.Ui32(v22&v20) {
		goto L8
	} else {
		goto L9
	}
L7:
	;
	v117 = v108 & int32(2147483647)
	if base.Ui32(v117) <= base.Ui32(v30) {
		goto L31
	} else {
		goto L32
	}
L8:
	;
	v27 = l0
	goto L10
L9:
	;
	v27 = l1
	goto L10
L10:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v27)+4))
	v30 = v28 & int32(2147483647)
	if v30 == int32(0) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v33 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v108 = v33
	goto L7
L12:
	;
	goto L13
L13:
	;
	v34 = int32(8)
	v35 = v27 + v34
	v37 = v26 + v34
	v38 = *(*int32)(unsafe.Add(mBase, uint32(v26)+4))
	v40 = int32(0)
	goto L14
L14:
	;
	v55 = v40 << (uint(int32(3)) % 32)
	v57 = *(*float64)(unsafe.Add(mBase, uint32(v37+v55)))
	v59 = base.B2i32(v38 < int32(0))
	if v38 < int32(0) {
		v67 = v57
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v108 = v38
	goto L7
L16:
	;
	v69 = *(*float64)(unsafe.Add(mBase, uint32(v55+v35)))
	v71 = base.B2i32(v28 < int32(0))
	if v28 < int32(0) {
		v79 = v69
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v64 = *(*float64)(unsafe.Add(mBase, uint32(v37+(v40+v38)<<(uint(int32(3))%32))))
	if base.F64_lt(v57, v64) != 0 {
		v67 = v57
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v67 = v64
	goto L16
L19:
	;
	v80 = int32(0)
	if base.F64_gt(v67, v79) != 0 {
		v184 = v80
		goto L1
	} else {
		goto L22
	}
L20:
	;
	v76 = *(*float64)(unsafe.Add(mBase, uint32(v35+(v40+v28)<<(uint(int32(3))%32))))
	if base.F64_gt(v69, v76) != 0 {
		v79 = v69
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v79 = v76
	goto L19
L22:
	;
	if v38 < int32(0) {
		v88 = v57
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v28 < int32(0) {
		v96 = v69
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v86 = *(*float64)(unsafe.Add(mBase, uint32(v37+(v40+v38)<<(uint(int32(3))%32))))
	if base.F64_gt(v57, v86) != 0 {
		v88 = v57
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v88 = v86
	goto L23
L26:
	;
	if base.F64_lt(v88, v96) != 0 {
		v184 = v80
		goto L1
	} else {
		goto L29
	}
L27:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(v35+(v40+v28)<<(uint(int32(3))%32))))
	if base.F64_lt(v69, v94) != 0 {
		v96 = v69
		goto L26
	} else {
		goto L28
	}
L28:
	;
	v96 = v94
	goto L26
L29:
	;
	v100 = v40 + int32(1)
	if v100 != v30 {
		v40 = v100
		goto L14
	} else {
		goto L30
	}
L30:
	;
	goto L15
L31:
	;
	return int32(1)
L32:
	;
	goto L33
L33:
	;
	v122 = v26 + int32(8)
	v126 = v30
	goto L34
L34:
	;
	v141 = v122 + v126<<(uint(int32(3))%32)
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v141)))
	if base.B2i32(v108 < int32(0)) == int32(0) {
		goto L38
	} else {
		goto L39
	}
L35:
	;
	v184 = int32(0)
	goto L1
L36:
	;
	goto L35
L37:
	;
	if base.F64_lt(v164, float64(0)) != 0 {
		goto L36
	} else {
		goto L49
	}
L38:
	;
	v146 = int32(3)
	v148 = v122 + (v126+v108)<<(uint(v146)%32)
	v153 = *(*float64)(unsafe.Add(mBase, uint32(v122+(v126+v117)<<(uint(v146)%32))))
	if base.F64_lt(v142, v153) != 0 {
		goto L41
	} else {
		goto L42
	}
L39:
	;
	goto L40
L40:
	;
	if base.F64_gt(v142, float64(0)) != 0 {
		goto L36
	} else {
		goto L48
	}
L41:
	;
	v155 = v141
	goto L43
L42:
	;
	v155 = v148
	goto L43
L43:
	;
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v155)))
	if base.F64_gt(v156, float64(0)) != 0 {
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v159 = *(*float64)(unsafe.Add(mBase, uint32(v148)))
	if base.F64_gt(v142, v159) != 0 {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v161 = v142
	goto L47
L46:
	;
	v161 = v159
	goto L47
L47:
	;
	v164 = v161
	goto L37
L48:
	;
	v164 = v142
	goto L37
L49:
	;
	v169 = int32(1)
	v171 = v126 + v169
	if v117 != v171 {
		v126 = v171
		goto L34
	} else {
		goto L50
	}
L50:
	;
	v184 = v169
	goto L1
}
func F_cube_yy_create_buffer(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v44 int32
	_ = v44
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	v8 = F_palloc(m, int32(48))
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int32(0)
	} else {
		if v8 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = l1
			v15 = F_palloc(m, l1+int32(2))
			mBase = m.M
			v16 = m.ExcPending
			if v16 != 0 {
				return int32(0)
			} else {
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v15
				if v15 == int32(0) {
					F_yy_fatal_error_6(m, int32(684057))
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				} else {
					v20 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v20
					v23 = *(*int32)(unsafe.Add(mBase, _consts[159]))
					v24 = int32(0)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v24
					*(*uint8)(unsafe.Add(mBase, uint32(v15))) = uint8(v24)
					v28 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*uint8)(unsafe.Add(mBase, uint32(v28)+1)) = uint8(v24)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+44)) = v24
					*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v20
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = v35
					v37 = *(*int32)(unsafe.Add(mBase, uint32(l2)+20))
					if v37 == v24 {
						v59 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v59
						*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
						*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v59
						v66 = int32(0)
						v67 = int32(36)
					} else {
						v40 = *(*int32)(unsafe.Add(mBase, uint32(l2)+12))
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v37+v40<<(uint(int32(2))%32))))
						if v8 != v44 {
							v59 = int32(1)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+40)) = v59
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v59
							v66 = int32(0)
							v67 = int32(36)
						} else {
							v46 = *(*int32)(unsafe.Add(mBase, uint32(v44)+16))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+28)) = v46
							v48 = *(*int32)(unsafe.Add(mBase, uint32(v44)+8))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+80)) = v48
							*(*int32)(unsafe.Add(mBase, uint32(l2)+36)) = v48
							v51 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
							*(*int32)(unsafe.Add(mBase, uint32(l2)+4)) = v51
							v53 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v48))))
							*(*uint8)(unsafe.Add(mBase, uint32(l2)+24)) = uint8(v53)
							*(*int32)(unsafe.Add(mBase, uint32(v8))) = l0
							v66 = int32(1)
							v67 = int32(40)
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v67+v8))) = v66
					*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = int32(0)
					*(*int32)(unsafe.Add(mBase, _consts[159])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_6(m, int32(684057))
			mBase = m.M
			v77 = m.ExcPending
			if v77 != 0 {
				return int32(0)
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	}
}
func F_cube_yy_switch_to_buffer(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	F_cube_yyensure_buffer_stack(m, l1)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
		if v8 == int32(0) {
		} else {
			v11 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
			v15 = *(*int32)(unsafe.Add(mBase, uint32(v8+v11<<(uint(int32(2))%32))))
			if v15 == l0 {
			} else {
				if v15 != 0 {
					v17 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)))
					*(*uint8)(unsafe.Add(mBase, uint32(v17))) = uint8(v18)
					v20 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
					v21 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
					v25 = *(*int32)(unsafe.Add(mBase, uint32(v20+v21<<(uint(int32(2))%32))))
					v26 = *(*int32)(unsafe.Add(mBase, uint32(l1)+36))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+8)) = v26
					v28 = *(*int32)(unsafe.Add(mBase, uint32(l1)+28))
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v28
					v30 = v20
					v32 = v21
				} else {
					v30 = v8
					v32 = v11
				}
				*(*int32)(unsafe.Add(mBase, uint32(v30+v32<<(uint(int32(2))%32)))) = l0
				v37 = *(*int32)(unsafe.Add(mBase, uint32(l0)+16))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+28)) = v37
				v39 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+80)) = v39
				*(*int32)(unsafe.Add(mBase, uint32(l1)+36)) = v39
				v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+4)) = v42
				v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v39))))
				*(*int32)(unsafe.Add(mBase, uint32(l1)+48)) = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l1)+24)) = uint8(v44)
			}
		}
		return
	}
}
func F_cube_yyerror(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32, l4 int32) {
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
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v11))))
	v13 = F_errsave_start(m, l2)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return
	} else {
		if v12 == int32(0) {
			if v13 == int32(0) {
				m.G0 = v9 + int32(32)
				return
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return
				} else {
					F_errmsg(m, int32(420120), int32(0))
					mBase = m.M
					v27 = m.ExcPending
					if v27 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l4
						F_errdetail(m, int32(65048), v9)
						mBase = m.M
						v32 = m.ExcPending
						if v32 != 0 {
							return
						} else {
							v57 = int32(85)
							F_errsave_finish(m, l2, int32(314760), v57, int32(210259))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		} else {
			if v13 == int32(0) {
				m.G0 = v9 + int32(32)
				return
			} else {
				F_errcode(m, int32(33685634))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return
				} else {
					F_errmsg(m, int32(420120), int32(0))
					mBase = m.M
					v44 = m.ExcPending
					if v44 != 0 {
						return
					} else {
						v45 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v45
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l4
						F_errdetail(m, int32(702242), v9+int32(16))
						mBase = m.M
						v53 = m.ExcPending
						if v53 != 0 {
							return
						} else {
							v57 = int32(93)
							F_errsave_finish(m, l2, int32(314760), v57, int32(210259))
							mBase = m.M
							v64 = m.ExcPending
							if v64 != 0 {
								return
							} else {
								m.G0 = v9 + int32(32)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_cube_yyfree(m *base.Module, l0 int32, l1 int32) {
	var v4 int32
	_ = v4
	if l0 != 0 {
		F_pfree(m, l0)
		v4 = m.ExcPending
		if v4 != 0 {
			return
		} else {
			return
		}
	} else {
		return
	}
}
func F_cube_yyget_debug(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+76))
	return v2
}
func F_cube_yyget_lineno(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	v2 = int32(0)
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v4 == v2 {
		v16 = v2
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+12))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v4+v7<<(uint(int32(2))%32))))
		if v11 == int32(0) {
			v16 = v2
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v11)+32))
			v16 = v14
		}
	}
	return v16
}
func F_cube_yyset_lineno(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l1)+20))
	if v4 != 0 {
		v5 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v4+v5<<(uint(int32(2))%32))))
		if v9 != 0 {
			*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = l0
			return
		} else {
			F_yy_fatal_error_6(m, int32(225808))
			mBase = m.M
			v13 = m.ExcPending
			if v13 != 0 {
				return
			} else {
				base.Wasm_trap_unreachable()
				for {
				}
			}
		}
	} else {
		F_yy_fatal_error_6(m, int32(225808))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return
		} else {
			base.Wasm_trap_unreachable()
			for {
			}
		}
	}
}
func F_cube_yyset_lval(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	*(*int32)(unsafe.Add(mBase, uint32(l1)+92)) = l0
	return
}
