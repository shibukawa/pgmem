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
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v130 int32
	_ = v130
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
	var v167 int32
	_ = v167
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
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
				v28 = F_ArrayGetNItemsSafe(m, v25, v17+int32(16))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return int32(0)
				} else {
					if int32(101) <= v28 {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v176 = m.ExcPending
						if v176 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(261))
							mBase = m.M
							v179 = m.ExcPending
							if v179 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_cube_a_f8_0), int32(0))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = int32(100)
									F_errdetail(m, int32(_a_F_cube_a_f8_1), v14)
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_cube_a_f8_2), int32(230), int32(_a_F_cube_a_f8_3))
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
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
									if v59 == int32(0) {
									} else {
										v109 = v103
										v120 = v109
										v130 = v2
										for {
											v132 = v120 << (uint(int32(3)) % 32)
											v135 = *(*float64)(unsafe.Add(mBase, uint32(v132+v57)))
											*(*float64)(unsafe.Add(mBase, uint32(v61+v132))) = v135
											v137 = int32(1)
											v140 = v130 + v137
											if v140 != v59 {
												v120 = v120 + v137
												v130 = v140
												continue
											} else {
												break
											}
											break
										}
									}
								} else {
									v109 = v62
									v120 = v109
									v130 = v2
									for {
										v132 = v120 << (uint(int32(3)) % 32)
										v135 = *(*float64)(unsafe.Add(mBase, uint32(v132+v57)))
										*(*float64)(unsafe.Add(mBase, uint32(v61+v132))) = v135
										v137 = int32(1)
										v140 = v130 + v137
										if v140 != v59 {
											v120 = v120 + v137
											v130 = v140
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
						F_errmsg(m, int32(_a_F_cube_a_f8_4), int32(0))
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cube_a_f8_2), int32(222), int32(_a_F_cube_a_f8_3))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
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
	var v70 int32
	_ = v70
	var v76 int32
	_ = v76
	var v86 int32
	_ = v86
	var v89 float64
	_ = v89
	var v92 int32
	_ = v92
	var v95 float64
	_ = v95
	var v98 int32
	_ = v98
	var v101 float64
	_ = v101
	var v104 int32
	_ = v104
	var v107 float64
	_ = v107
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v117 int32
	_ = v117
	var v133 int32
	_ = v133
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v152 float64
	_ = v152
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v241 int32
	_ = v241
	var v244 float64
	_ = v244
	var v248 float64
	_ = v248
	var v251 int32
	_ = v251
	var v254 float64
	_ = v254
	var v258 float64
	_ = v258
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v263 int32
	_ = v263
	var v272 int32
	_ = v272
	var v284 int32
	_ = v284
	var v287 float64
	_ = v287
	var v291 float64
	_ = v291
	var v303 int32
	_ = v303
	var v320 int32
	_ = v320
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v344 int32
	_ = v344
	var v347 int32
	_ = v347
	var v351 int32
	_ = v351
	var v356 int32
	_ = v356
	var v361 int32
	_ = v361
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
							v70 = v63
							v76 = int32(0)
							for {
								v86 = v70 << (uint(int32(3)) % 32)
								v89 = *(*float64)(unsafe.Add(mBase, uint32(v62+v86)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v86))) = v89
								v92 = v86 | int32(8)
								v95 = *(*float64)(unsafe.Add(mBase, uint32(v62+v92)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v92))) = v95
								v98 = v86 | int32(16)
								v101 = *(*float64)(unsafe.Add(mBase, uint32(v62+v98)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v98))) = v101
								v104 = v86 | int32(24)
								v107 = *(*float64)(unsafe.Add(mBase, uint32(v62+v104)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v104))) = v107
								v109 = int32(4)
								v110 = v70 + v109
								v112 = v76 + v109
								if v112 != v52&int32(2147483644) {
									v70 = v110
									v76 = v112
									continue
								} else {
									break
								}
								break
							}
							if v58 == int32(0) {
							} else {
								v117 = v110
								v133 = v117
								v145 = v2
								for {
									v149 = v133 << (uint(int32(3)) % 32)
									v152 = *(*float64)(unsafe.Add(mBase, uint32(v62+v149)))
									*(*float64)(unsafe.Add(mBase, uint32(v60+v149))) = v152
									v154 = int32(1)
									v157 = v145 + v154
									if v157 != v58 {
										v133 = v133 + v154
										v145 = v157
										continue
									} else {
										break
									}
									break
								}
							}
						} else {
							v117 = v63
							v133 = v117
							v145 = v2
							for {
								v149 = v133 << (uint(int32(3)) % 32)
								v152 = *(*float64)(unsafe.Add(mBase, uint32(v62+v149)))
								*(*float64)(unsafe.Add(mBase, uint32(v60+v149))) = v152
								v154 = int32(1)
								v157 = v145 + v154
								if v157 != v58 {
									v133 = v133 + v154
									v145 = v157
									continue
								} else {
									break
								}
								break
							}
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v41+v48<<(uint(int32(3))%32)))) = v34
					v320 = v41
					v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v333 != v22 {
						F_pfree(m, v22)
						mBase = m.M
						v336 = m.ExcPending
						if v336 != 0 {
							return int32(0)
						} else {
							m.G0 = v19 + int32(16)
							return v320
						}
					} else {
						m.G0 = v19 + int32(16)
						return v320
					}
				}
			} else {
				v182 = v32<<(uint(int32(4))%32) | int32(8)
				v183 = F_palloc0(m, v182)
				mBase = m.M
				v184 = m.ExcPending
				if v184 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v183))) = v182 << (uint(int32(2)) % 32)
					v188 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					v189 = int32(2147483647)
					v192 = v188&v189 + int32(1)
					v193 = *(*int32)(unsafe.Add(mBase, uint32(v183)+4))
					*(*int32)(unsafe.Add(mBase, uint32(v183)+4)) = v192 | v193&int32(-2147483648)
					v198 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
					v200 = v198 & v189
					if v200 == int32(0) {
						v303 = v192 & int32(2147483647)
					} else {
						v205 = int32(8)
						v206 = v183 + v205
						v207 = int32(3)
						v209 = v206 + v192<<(uint(v207)%32)
						v211 = v22 + v205
						v214 = v211 + v200<<(uint(v207)%32)
						v216 = v192 & int32(2147483647)
						if v200 != int32(1) {
							v229 = v2
							v232 = int32(0)
							for {
								v241 = v229 << (uint(int32(3)) % 32)
								v244 = *(*float64)(unsafe.Add(mBase, uint32(v241+v211)))
								*(*float64)(unsafe.Add(mBase, uint32(v206+v241))) = v244
								v248 = *(*float64)(unsafe.Add(mBase, uint32(v241+v214)))
								*(*float64)(unsafe.Add(mBase, uint32(v241+v209))) = v248
								v251 = v241 | int32(8)
								v254 = *(*float64)(unsafe.Add(mBase, uint32(v251+v211)))
								*(*float64)(unsafe.Add(mBase, uint32(v206+v251))) = v254
								v258 = *(*float64)(unsafe.Add(mBase, uint32(v251+v214)))
								*(*float64)(unsafe.Add(mBase, uint32(v251+v209))) = v258
								v260 = int32(2)
								v261 = v229 + v260
								v263 = v232 + v260
								if v263 != v198&int32(2147483646) {
									v229 = v261
									v232 = v263
									continue
								} else {
									break
								}
								break
							}
							if v198&int32(1) == int32(0) {
								v303 = v216
							} else {
								v272 = v261
								v284 = v272 << (uint(int32(3)) % 32)
								v287 = *(*float64)(unsafe.Add(mBase, uint32(v284+v211)))
								*(*float64)(unsafe.Add(mBase, uint32(v206+v284))) = v287
								v291 = *(*float64)(unsafe.Add(mBase, uint32(v284+v214)))
								*(*float64)(unsafe.Add(mBase, uint32(v284+v209))) = v291
								v303 = v216
							}
						} else {
							v272 = v2
							v284 = v272 << (uint(int32(3)) % 32)
							v287 = *(*float64)(unsafe.Add(mBase, uint32(v284+v211)))
							*(*float64)(unsafe.Add(mBase, uint32(v206+v284))) = v287
							v291 = *(*float64)(unsafe.Add(mBase, uint32(v284+v214)))
							*(*float64)(unsafe.Add(mBase, uint32(v284+v209))) = v291
							v303 = v216
						}
					}
					*(*float64)(unsafe.Add(mBase, uint32(v183+v303<<(uint(int32(3))%32)))) = v34
					*(*float64)(unsafe.Add(mBase, uint32(v183+v192<<(uint(int32(4))%32)))) = v34
					v320 = v183
					v333 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
					if v333 != v22 {
						F_pfree(m, v22)
						mBase = m.M
						v336 = m.ExcPending
						if v336 != 0 {
							return int32(0)
						} else {
							m.G0 = v19 + int32(16)
							return v320
						}
					} else {
						m.G0 = v19 + int32(16)
						return v320
					}
				}
			}
		} else {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v344 = m.ExcPending
			if v344 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(261))
				mBase = m.M
				v347 = m.ExcPending
				if v347 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(_a_F_cube_c_f8_0), int32(0))
					mBase = m.M
					v351 = m.ExcPending
					if v351 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(100)
						F_errdetail(m, int32(_a_F_cube_c_f8_1), v19)
						mBase = m.M
						v356 = m.ExcPending
						if v356 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_cube_c_f8_2), int32(1835), int32(_a_F_cube_c_f8_3))
							mBase = m.M
							v361 = m.ExcPending
							if v361 != 0 {
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v79 float64
	_ = v79
	var v82 float64
	_ = v82
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 float64
	_ = v129
	var v131 int32
	_ = v131
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v139 int32
	_ = v139
	var v140 float64
	_ = v140
	var v142 int32
	_ = v142
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v156 float64
	_ = v156
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v167 int32
	_ = v167
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v212 int32
	_ = v212
	var v213 float64
	_ = v213
	var v220 float64
	_ = v220
	var v222 int32
	_ = v222
	var v226 float64
	_ = v226
	var v232 float64
	_ = v232
	var v238 float64
	_ = v238
	var v242 int32
	_ = v242
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v274 float64
	_ = v274
	var v276 int32
	_ = v276
	var v280 float64
	_ = v280
	var v286 float64
	_ = v286
	var v292 float64
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v328 float64
	_ = v328
	var v335 float64
	_ = v335
	var v337 int32
	_ = v337
	var v341 float64
	_ = v341
	var v353 float64
	_ = v353
	var v357 float64
	_ = v357
	var v361 int32
	_ = v361
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v386 float64
	_ = v386
	var v393 float64
	_ = v393
	var v395 int32
	_ = v395
	var v399 float64
	_ = v399
	var v411 float64
	_ = v411
	var v415 float64
	_ = v415
	var v420 int32
	_ = v420
	var v452 int32
	_ = v452
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
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
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v497 != v6 {
		goto L114
	} else {
		goto L115
	}
L5:
	;
	v496 = int32(-1)
	goto L4
L6:
	;
	v496 = v452
	goto L4
L7:
	;
	v452 = int32(1)
	goto L6
L8:
	;
	v36 = v31
	goto L10
L9:
	;
	v36 = v34
	goto L10
L10:
	;
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = int32(8)
	v54 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L48
	} else {
		goto L49
	}
L14:
	;
	v60 = v54 << (uint(int32(3)) % 32)
	v61 = v6 + v37 + v60
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
	v64 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v71 = v62
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v103 = int32(8)
	v121 = int32(0)
	goto L31
L16:
	;
	v72 = v60 + (v11 + v37)
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v72)))
	v75 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v82 = v73
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v61+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v68) != 0 {
		v71 = v62
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v71 = v68
	goto L16
L19:
	;
	if base.F64_gt(v71, v82) != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v72+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v73, v79) != 0 {
		v82 = v73
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v82 = v79
	goto L19
L22:
	;
	if v29 < int32(0) {
		v89 = v62
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v32 < int32(0) {
		v96 = v73
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v61+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v87) != 0 {
		v89 = v62
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v89 = v87
	goto L23
L26:
	;
	v98 = int32(-1)
	if base.F64_lt(v89, v96) != 0 {
		v452 = v98
		goto L6
	} else {
		goto L29
	}
L27:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(v72+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v73, v94) != 0 {
		v96 = v73
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
	v101 = v54 + int32(1)
	if v101 != v36 {
		v54 = v101
		goto L14
	} else {
		goto L30
	}
L30:
	;
	goto L15
L31:
	;
	v127 = v121 << (uint(int32(3)) % 32)
	v128 = v6 + v103 + v127
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
	v131 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v138 = v129
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L13
L33:
	;
	v139 = v127 + (v11 + v103)
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v139)))
	v142 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v149 = v140
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v128+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v129, v135) != 0 {
		v138 = v129
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v138 = v135
	goto L33
L36:
	;
	if base.F64_gt(v138, v149) != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v139+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v140, v146) != 0 {
		v149 = v140
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v149 = v146
	goto L36
L39:
	;
	if v29 < int32(0) {
		v156 = v129
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v32 < int32(0) {
		v163 = v140
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v128+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v129, v154) != 0 {
		v156 = v129
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v156 = v154
	goto L40
L43:
	;
	if base.F64_lt(v156, v163) != 0 {
		v452 = v98
		goto L6
	} else {
		goto L46
	}
L44:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v139+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v140, v161) != 0 {
		v163 = v140
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v163 = v161
	goto L43
L46:
	;
	v167 = v121 + int32(1)
	if v167 != v36 {
		v121 = v167
		goto L31
	} else {
		goto L47
	}
L47:
	;
	goto L32
L48:
	;
	v189 = v6 + int32(8)
	v192 = v36
	goto L51
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v34) <= base.Ui32(v31) {
		goto L81
	} else {
		goto L82
	}
L51:
	;
	v212 = v189 + v192<<(uint(int32(3))%32)
	v213 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v255 = v36
	goto L65
L53:
	;
	if base.F64_lt(v238, float64(0)) != 0 {
		goto L5
	} else {
		goto L63
	}
L54:
	;
	v220 = *(*float64)(unsafe.Add(mBase, uint32(v212+v31<<(uint(int32(3))%32))))
	if base.F64_lt(v213, v220) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if base.F64_gt(v213, float64(0)) != 0 {
		goto L7
	} else {
		goto L62
	}
L57:
	;
	v222 = int32(0)
	goto L59
L58:
	;
	v222 = v29
	goto L59
L59:
	;
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v212+v222<<(uint(int32(3))%32))))
	if base.F64_gt(v226, float64(0)) != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v212+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v213, v232) == int32(0) {
		v238 = v232
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v238 = v213
	goto L53
L62:
	;
	v238 = v213
	goto L53
L63:
	;
	v242 = v192 + int32(1)
	if v242 != v31 {
		v192 = v242
		goto L51
	} else {
		goto L64
	}
L64:
	;
	goto L52
L65:
	;
	v266 = v189 + v255<<(uint(int32(3))%32)
	v267 = *(*float64)(unsafe.Add(mBase, uint32(v266)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L5
L67:
	;
	if base.F64_lt(v292, float64(0)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L68:
	;
	v274 = *(*float64)(unsafe.Add(mBase, uint32(v266+v31<<(uint(int32(3))%32))))
	if base.F64_gt(v267, v274) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if base.F64_gt(v267, float64(0)) != 0 {
		goto L7
	} else {
		goto L76
	}
L71:
	;
	v276 = int32(0)
	goto L73
L72:
	;
	v276 = v29
	goto L73
L73:
	;
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v266+v276<<(uint(int32(3))%32))))
	if base.F64_gt(v280, float64(0)) != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v266+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v267, v286) == int32(0) {
		v292 = v286
		goto L67
	} else {
		goto L75
	}
L75:
	;
	v292 = v267
	goto L67
L76:
	;
	v292 = v267
	goto L67
L77:
	;
	v297 = int32(1)
	v299 = v255 + v297
	if v299 == v31 {
		v452 = v297
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L66
L80:
	;
	v255 = v299
	goto L65
L81:
	;
	v496 = int32(0)
	goto L4
L82:
	;
	goto L83
L83:
	;
	v304 = v11 + int32(8)
	v314 = v31
	goto L84
L84:
	;
	v327 = v304 + v314<<(uint(int32(3))%32)
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v327)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v374 = v36
	goto L99
L86:
	;
	if base.F64_lt(v357, float64(0)) != 0 {
		goto L7
	} else {
		goto L97
	}
L87:
	;
	v353 = *(*float64)(unsafe.Add(mBase, uint32(v327+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v328, v353) == int32(0) {
		v357 = v353
		goto L86
	} else {
		goto L96
	}
L88:
	;
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v327+v34<<(uint(int32(3))%32))))
	if base.F64_lt(v328, v335) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if base.F64_gt(v328, float64(0)) == int32(0) {
		v357 = v328
		goto L86
	} else {
		goto L95
	}
L91:
	;
	v337 = int32(0)
	goto L93
L92:
	;
	v337 = v32
	goto L93
L93:
	;
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v327+v337<<(uint(int32(3))%32))))
	if base.F64_gt(v341, float64(0)) == int32(0) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	goto L5
L95:
	;
	goto L5
L96:
	;
	v357 = v328
	goto L86
L97:
	;
	v361 = v314 + int32(1)
	if v361 != v34 {
		v314 = v361
		goto L84
	} else {
		goto L98
	}
L98:
	;
	goto L85
L99:
	;
	v385 = v304 + v374<<(uint(int32(3))%32)
	v386 = *(*float64)(unsafe.Add(mBase, uint32(v385)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v452 = int32(-1)
	goto L6
L101:
	;
	if base.F64_lt(v415, float64(0)) != 0 {
		goto L7
	} else {
		goto L112
	}
L102:
	;
	v411 = *(*float64)(unsafe.Add(mBase, uint32(v385+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v386, v411) == int32(0) {
		v415 = v411
		goto L101
	} else {
		goto L111
	}
L103:
	;
	v393 = *(*float64)(unsafe.Add(mBase, uint32(v385+v34<<(uint(int32(3))%32))))
	if base.F64_gt(v386, v393) != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	if base.F64_gt(v386, float64(0)) == int32(0) {
		v415 = v386
		goto L101
	} else {
		goto L110
	}
L106:
	;
	v395 = int32(0)
	goto L108
L107:
	;
	v395 = v32
	goto L108
L108:
	;
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v385+v395<<(uint(int32(3))%32))))
	if base.F64_gt(v399, float64(0)) == int32(0) {
		goto L102
	} else {
		goto L109
	}
L109:
	;
	goto L5
L110:
	;
	goto L5
L111:
	;
	v415 = v386
	goto L101
L112:
	;
	v420 = v374 + int32(1)
	if v34 != v420 {
		v374 = v420
		goto L99
	} else {
		goto L113
	}
L113:
	;
	goto L100
L114:
	;
	F_pfree(m, v6)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v501 != v11 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	F_pfree(m, v11)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	return v496
L121:
	;
	goto L120
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
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 float64
	_ = v62
	var v64 int32
	_ = v64
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
	var v72 int32
	_ = v72
	var v73 float64
	_ = v73
	var v75 int32
	_ = v75
	var v79 float64
	_ = v79
	var v82 float64
	_ = v82
	var v87 float64
	_ = v87
	var v89 float64
	_ = v89
	var v94 float64
	_ = v94
	var v96 float64
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v121 int32
	_ = v121
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v129 float64
	_ = v129
	var v131 int32
	_ = v131
	var v135 float64
	_ = v135
	var v138 float64
	_ = v138
	var v139 int32
	_ = v139
	var v140 float64
	_ = v140
	var v142 int32
	_ = v142
	var v146 float64
	_ = v146
	var v149 float64
	_ = v149
	var v154 float64
	_ = v154
	var v156 float64
	_ = v156
	var v161 float64
	_ = v161
	var v163 float64
	_ = v163
	var v167 int32
	_ = v167
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v212 int32
	_ = v212
	var v213 float64
	_ = v213
	var v220 float64
	_ = v220
	var v222 int32
	_ = v222
	var v226 float64
	_ = v226
	var v232 float64
	_ = v232
	var v238 float64
	_ = v238
	var v242 int32
	_ = v242
	var v255 int32
	_ = v255
	var v266 int32
	_ = v266
	var v267 float64
	_ = v267
	var v274 float64
	_ = v274
	var v276 int32
	_ = v276
	var v280 float64
	_ = v280
	var v286 float64
	_ = v286
	var v292 float64
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v304 int32
	_ = v304
	var v314 int32
	_ = v314
	var v327 int32
	_ = v327
	var v328 float64
	_ = v328
	var v335 float64
	_ = v335
	var v337 int32
	_ = v337
	var v341 float64
	_ = v341
	var v353 float64
	_ = v353
	var v357 float64
	_ = v357
	var v361 int32
	_ = v361
	var v374 int32
	_ = v374
	var v385 int32
	_ = v385
	var v386 float64
	_ = v386
	var v393 float64
	_ = v393
	var v395 int32
	_ = v395
	var v399 float64
	_ = v399
	var v411 float64
	_ = v411
	var v415 float64
	_ = v415
	var v420 int32
	_ = v420
	var v452 int32
	_ = v452
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v500 int32
	_ = v500
	var v501 int32
	_ = v501
	var v504 int32
	_ = v504
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
		goto L8
	} else {
		goto L9
	}
L4:
	;
	v497 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	if v497 != v6 {
		goto L114
	} else {
		goto L115
	}
L5:
	;
	v496 = int32(-1)
	goto L4
L6:
	;
	v496 = v452
	goto L4
L7:
	;
	v452 = int32(1)
	goto L6
L8:
	;
	v36 = v31
	goto L10
L9:
	;
	v36 = v34
	goto L10
L10:
	;
	if v36 != 0 {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v37 = int32(8)
	v54 = int32(0)
	goto L14
L12:
	;
	goto L13
L13:
	;
	if base.Ui32(v34) < base.Ui32(v31) {
		goto L48
	} else {
		goto L49
	}
L14:
	;
	v60 = v54 << (uint(int32(3)) % 32)
	v61 = v6 + v37 + v60
	v62 = *(*float64)(unsafe.Add(mBase, uint32(v61)))
	v64 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v71 = v62
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v103 = int32(8)
	v121 = int32(0)
	goto L31
L16:
	;
	v72 = v60 + (v11 + v37)
	v73 = *(*float64)(unsafe.Add(mBase, uint32(v72)))
	v75 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v82 = v73
		goto L19
	} else {
		goto L20
	}
L17:
	;
	v68 = *(*float64)(unsafe.Add(mBase, uint32(v61+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v68) != 0 {
		v71 = v62
		goto L16
	} else {
		goto L18
	}
L18:
	;
	v71 = v68
	goto L16
L19:
	;
	if base.F64_gt(v71, v82) != 0 {
		goto L7
	} else {
		goto L22
	}
L20:
	;
	v79 = *(*float64)(unsafe.Add(mBase, uint32(v72+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v73, v79) != 0 {
		v82 = v73
		goto L19
	} else {
		goto L21
	}
L21:
	;
	v82 = v79
	goto L19
L22:
	;
	if v29 < int32(0) {
		v89 = v62
		goto L23
	} else {
		goto L24
	}
L23:
	;
	if v32 < int32(0) {
		v96 = v73
		goto L26
	} else {
		goto L27
	}
L24:
	;
	v87 = *(*float64)(unsafe.Add(mBase, uint32(v61+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v62, v87) != 0 {
		v89 = v62
		goto L23
	} else {
		goto L25
	}
L25:
	;
	v89 = v87
	goto L23
L26:
	;
	v98 = int32(-1)
	if base.F64_lt(v89, v96) != 0 {
		v452 = v98
		goto L6
	} else {
		goto L29
	}
L27:
	;
	v94 = *(*float64)(unsafe.Add(mBase, uint32(v72+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v73, v94) != 0 {
		v96 = v73
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
	v101 = v54 + int32(1)
	if v101 != v36 {
		v54 = v101
		goto L14
	} else {
		goto L30
	}
L30:
	;
	goto L15
L31:
	;
	v127 = v121 << (uint(int32(3)) % 32)
	v128 = v6 + v103 + v127
	v129 = *(*float64)(unsafe.Add(mBase, uint32(v128)))
	v131 = base.B2i32(v29 < int32(0))
	if v29 < int32(0) {
		v138 = v129
		goto L33
	} else {
		goto L34
	}
L32:
	;
	goto L13
L33:
	;
	v139 = v127 + (v11 + v103)
	v140 = *(*float64)(unsafe.Add(mBase, uint32(v139)))
	v142 = base.B2i32(v32 < int32(0))
	if v32 < int32(0) {
		v149 = v140
		goto L36
	} else {
		goto L37
	}
L34:
	;
	v135 = *(*float64)(unsafe.Add(mBase, uint32(v128+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v129, v135) != 0 {
		v138 = v129
		goto L33
	} else {
		goto L35
	}
L35:
	;
	v138 = v135
	goto L33
L36:
	;
	if base.F64_gt(v138, v149) != 0 {
		goto L7
	} else {
		goto L39
	}
L37:
	;
	v146 = *(*float64)(unsafe.Add(mBase, uint32(v139+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v140, v146) != 0 {
		v149 = v140
		goto L36
	} else {
		goto L38
	}
L38:
	;
	v149 = v146
	goto L36
L39:
	;
	if v29 < int32(0) {
		v156 = v129
		goto L40
	} else {
		goto L41
	}
L40:
	;
	if v32 < int32(0) {
		v163 = v140
		goto L43
	} else {
		goto L44
	}
L41:
	;
	v154 = *(*float64)(unsafe.Add(mBase, uint32(v128+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v129, v154) != 0 {
		v156 = v129
		goto L40
	} else {
		goto L42
	}
L42:
	;
	v156 = v154
	goto L40
L43:
	;
	if base.F64_lt(v156, v163) != 0 {
		v452 = v98
		goto L6
	} else {
		goto L46
	}
L44:
	;
	v161 = *(*float64)(unsafe.Add(mBase, uint32(v139+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v140, v161) != 0 {
		v163 = v140
		goto L43
	} else {
		goto L45
	}
L45:
	;
	v163 = v161
	goto L43
L46:
	;
	v167 = v121 + int32(1)
	if v167 != v36 {
		v121 = v167
		goto L31
	} else {
		goto L47
	}
L47:
	;
	goto L32
L48:
	;
	v189 = v6 + int32(8)
	v192 = v36
	goto L51
L49:
	;
	goto L50
L50:
	;
	if base.Ui32(v34) <= base.Ui32(v31) {
		goto L81
	} else {
		goto L82
	}
L51:
	;
	v212 = v189 + v192<<(uint(int32(3))%32)
	v213 = *(*float64)(unsafe.Add(mBase, uint32(v212)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L54
	} else {
		goto L55
	}
L52:
	;
	v255 = v36
	goto L65
L53:
	;
	if base.F64_lt(v238, float64(0)) != 0 {
		goto L5
	} else {
		goto L63
	}
L54:
	;
	v220 = *(*float64)(unsafe.Add(mBase, uint32(v212+v31<<(uint(int32(3))%32))))
	if base.F64_lt(v213, v220) != 0 {
		goto L57
	} else {
		goto L58
	}
L55:
	;
	goto L56
L56:
	;
	if base.F64_gt(v213, float64(0)) != 0 {
		goto L7
	} else {
		goto L62
	}
L57:
	;
	v222 = int32(0)
	goto L59
L58:
	;
	v222 = v29
	goto L59
L59:
	;
	v226 = *(*float64)(unsafe.Add(mBase, uint32(v212+v222<<(uint(int32(3))%32))))
	if base.F64_gt(v226, float64(0)) != 0 {
		goto L7
	} else {
		goto L60
	}
L60:
	;
	v232 = *(*float64)(unsafe.Add(mBase, uint32(v212+v29<<(uint(int32(3))%32))))
	if base.F64_lt(v213, v232) == int32(0) {
		v238 = v232
		goto L53
	} else {
		goto L61
	}
L61:
	;
	v238 = v213
	goto L53
L62:
	;
	v238 = v213
	goto L53
L63:
	;
	v242 = v192 + int32(1)
	if v242 != v31 {
		v192 = v242
		goto L51
	} else {
		goto L64
	}
L64:
	;
	goto L52
L65:
	;
	v266 = v189 + v255<<(uint(int32(3))%32)
	v267 = *(*float64)(unsafe.Add(mBase, uint32(v266)))
	if base.B2i32(v29 < int32(0)) == int32(0) {
		goto L68
	} else {
		goto L69
	}
L66:
	;
	goto L5
L67:
	;
	if base.F64_lt(v292, float64(0)) == int32(0) {
		goto L77
	} else {
		goto L78
	}
L68:
	;
	v274 = *(*float64)(unsafe.Add(mBase, uint32(v266+v31<<(uint(int32(3))%32))))
	if base.F64_gt(v267, v274) != 0 {
		goto L71
	} else {
		goto L72
	}
L69:
	;
	goto L70
L70:
	;
	if base.F64_gt(v267, float64(0)) != 0 {
		goto L7
	} else {
		goto L76
	}
L71:
	;
	v276 = int32(0)
	goto L73
L72:
	;
	v276 = v29
	goto L73
L73:
	;
	v280 = *(*float64)(unsafe.Add(mBase, uint32(v266+v276<<(uint(int32(3))%32))))
	if base.F64_gt(v280, float64(0)) != 0 {
		goto L7
	} else {
		goto L74
	}
L74:
	;
	v286 = *(*float64)(unsafe.Add(mBase, uint32(v266+v29<<(uint(int32(3))%32))))
	if base.F64_gt(v267, v286) == int32(0) {
		v292 = v286
		goto L67
	} else {
		goto L75
	}
L75:
	;
	v292 = v267
	goto L67
L76:
	;
	v292 = v267
	goto L67
L77:
	;
	v297 = int32(1)
	v299 = v255 + v297
	if v299 == v31 {
		v452 = v297
		goto L6
	} else {
		goto L80
	}
L78:
	;
	goto L79
L79:
	;
	goto L66
L80:
	;
	v255 = v299
	goto L65
L81:
	;
	v496 = int32(0)
	goto L4
L82:
	;
	goto L83
L83:
	;
	v304 = v11 + int32(8)
	v314 = v31
	goto L84
L84:
	;
	v327 = v304 + v314<<(uint(int32(3))%32)
	v328 = *(*float64)(unsafe.Add(mBase, uint32(v327)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L88
	} else {
		goto L89
	}
L85:
	;
	v374 = v36
	goto L99
L86:
	;
	if base.F64_lt(v357, float64(0)) != 0 {
		goto L7
	} else {
		goto L97
	}
L87:
	;
	v353 = *(*float64)(unsafe.Add(mBase, uint32(v327+v32<<(uint(int32(3))%32))))
	if base.F64_lt(v328, v353) == int32(0) {
		v357 = v353
		goto L86
	} else {
		goto L96
	}
L88:
	;
	v335 = *(*float64)(unsafe.Add(mBase, uint32(v327+v34<<(uint(int32(3))%32))))
	if base.F64_lt(v328, v335) != 0 {
		goto L91
	} else {
		goto L92
	}
L89:
	;
	goto L90
L90:
	;
	if base.F64_gt(v328, float64(0)) == int32(0) {
		v357 = v328
		goto L86
	} else {
		goto L95
	}
L91:
	;
	v337 = int32(0)
	goto L93
L92:
	;
	v337 = v32
	goto L93
L93:
	;
	v341 = *(*float64)(unsafe.Add(mBase, uint32(v327+v337<<(uint(int32(3))%32))))
	if base.F64_gt(v341, float64(0)) == int32(0) {
		goto L87
	} else {
		goto L94
	}
L94:
	;
	goto L5
L95:
	;
	goto L5
L96:
	;
	v357 = v328
	goto L86
L97:
	;
	v361 = v314 + int32(1)
	if v361 != v34 {
		v314 = v361
		goto L84
	} else {
		goto L98
	}
L98:
	;
	goto L85
L99:
	;
	v385 = v304 + v374<<(uint(int32(3))%32)
	v386 = *(*float64)(unsafe.Add(mBase, uint32(v385)))
	if base.B2i32(v32 < int32(0)) == int32(0) {
		goto L103
	} else {
		goto L104
	}
L100:
	;
	v452 = int32(-1)
	goto L6
L101:
	;
	if base.F64_lt(v415, float64(0)) != 0 {
		goto L7
	} else {
		goto L112
	}
L102:
	;
	v411 = *(*float64)(unsafe.Add(mBase, uint32(v385+v32<<(uint(int32(3))%32))))
	if base.F64_gt(v386, v411) == int32(0) {
		v415 = v411
		goto L101
	} else {
		goto L111
	}
L103:
	;
	v393 = *(*float64)(unsafe.Add(mBase, uint32(v385+v34<<(uint(int32(3))%32))))
	if base.F64_gt(v386, v393) != 0 {
		goto L106
	} else {
		goto L107
	}
L104:
	;
	goto L105
L105:
	;
	if base.F64_gt(v386, float64(0)) == int32(0) {
		v415 = v386
		goto L101
	} else {
		goto L110
	}
L106:
	;
	v395 = int32(0)
	goto L108
L107:
	;
	v395 = v32
	goto L108
L108:
	;
	v399 = *(*float64)(unsafe.Add(mBase, uint32(v385+v395<<(uint(int32(3))%32))))
	if base.F64_gt(v399, float64(0)) == int32(0) {
		goto L102
	} else {
		goto L109
	}
L109:
	;
	goto L5
L110:
	;
	goto L5
L111:
	;
	v415 = v386
	goto L101
L112:
	;
	v420 = v374 + int32(1)
	if v34 != v420 {
		v374 = v420
		goto L99
	} else {
		goto L113
	}
L113:
	;
	goto L100
L114:
	;
	F_pfree(m, v6)
	mBase = m.M
	v500 = m.ExcPending
	if v500 != 0 {
		goto L1
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	v501 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v501 != v11 {
		goto L118
	} else {
		goto L119
	}
L117:
	;
	goto L116
L118:
	;
	F_pfree(m, v11)
	mBase = m.M
	v504 = m.ExcPending
	if v504 != 0 {
		goto L1
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	return int32(base.Ui32(v496^int32(-1)) >> (uint(int32(31)) % 32))
L121:
	;
	goto L120
}
func F_cube_overlap_v0(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v21 int32
	_ = v21
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
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v43 int32
	_ = v43
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 float64
	_ = v60
	var v62 int32
	_ = v62
	var v66 float64
	_ = v66
	var v69 float64
	_ = v69
	var v70 int32
	_ = v70
	var v71 float64
	_ = v71
	var v73 int32
	_ = v73
	var v77 float64
	_ = v77
	var v80 float64
	_ = v80
	var v85 float64
	_ = v85
	var v87 float64
	_ = v87
	var v92 float64
	_ = v92
	var v94 float64
	_ = v94
	var v98 int32
	_ = v98
	var v100 int32
	_ = v100
	var v116 int32
	_ = v116
	var v129 int32
	_ = v129
	var v141 int32
	_ = v141
	var v142 float64
	_ = v142
	var v145 int32
	_ = v145
	var v150 float64
	_ = v150
	var v152 int32
	_ = v152
	var v156 float64
	_ = v156
	var v162 float64
	_ = v162
	var v168 float64
	_ = v168
	var v172 int32
	_ = v172
	var v174 int32
	_ = v174
	var v187 int32
	_ = v187
	v6 = int32(0)
	if base.B2i32(l0 == v6)|base.B2i32(l1 == v6) != 0 {
		v187 = v6
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return v187
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = int32(2147483647)
	v24 = *(*int32)(unsafe.Add(mBase, uint32(l1)+4))
	v27 = base.B2i32(base.Ui32(v21&v22) < base.Ui32(v24&v22))
	if base.Ui32(v21&v22) < base.Ui32(v24&v22) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v28 = l1
	goto L5
L4:
	;
	v28 = l0
	goto L5
L5:
	;
	if base.Ui32(v21&v22) < base.Ui32(v24&v22) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v116 = v100 & int32(2147483647)
	if base.Ui32(v116) <= base.Ui32(v32) {
		goto L30
	} else {
		goto L31
	}
L7:
	;
	v29 = l0
	goto L9
L8:
	;
	v29 = l1
	goto L9
L9:
	;
	v30 = *(*int32)(unsafe.Add(mBase, uint32(v29)+4))
	v32 = v30 & int32(2147483647)
	if v32 == int32(0) {
		goto L10
	} else {
		goto L11
	}
L10:
	;
	v35 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v100 = v35
	goto L6
L11:
	;
	goto L12
L12:
	;
	v36 = int32(8)
	v40 = *(*int32)(unsafe.Add(mBase, uint32(v28)+4))
	v43 = int32(0)
	goto L13
L13:
	;
	v58 = v43 << (uint(int32(3)) % 32)
	v59 = v28 + v36 + v58
	v60 = *(*float64)(unsafe.Add(mBase, uint32(v59)))
	v62 = base.B2i32(v40 < int32(0))
	if v40 < int32(0) {
		v69 = v60
		goto L15
	} else {
		goto L16
	}
L14:
	;
	v100 = v40
	goto L6
L15:
	;
	v70 = v58 + (v29 + v36)
	v71 = *(*float64)(unsafe.Add(mBase, uint32(v70)))
	v73 = base.B2i32(v30 < int32(0))
	if v30 < int32(0) {
		v80 = v71
		goto L18
	} else {
		goto L19
	}
L16:
	;
	v66 = *(*float64)(unsafe.Add(mBase, uint32(v59+v40<<(uint(int32(3))%32))))
	if base.F64_lt(v60, v66) != 0 {
		v69 = v60
		goto L15
	} else {
		goto L17
	}
L17:
	;
	v69 = v66
	goto L15
L18:
	;
	if base.F64_gt(v69, v80) != 0 {
		v187 = v6
		goto L1
	} else {
		goto L21
	}
L19:
	;
	v77 = *(*float64)(unsafe.Add(mBase, uint32(v70+v30<<(uint(int32(3))%32))))
	if base.F64_gt(v71, v77) != 0 {
		v80 = v71
		goto L18
	} else {
		goto L20
	}
L20:
	;
	v80 = v77
	goto L18
L21:
	;
	if v40 < int32(0) {
		v87 = v60
		goto L22
	} else {
		goto L23
	}
L22:
	;
	if v30 < int32(0) {
		v94 = v71
		goto L25
	} else {
		goto L26
	}
L23:
	;
	v85 = *(*float64)(unsafe.Add(mBase, uint32(v59+v40<<(uint(int32(3))%32))))
	if base.F64_gt(v60, v85) != 0 {
		v87 = v60
		goto L22
	} else {
		goto L24
	}
L24:
	;
	v87 = v85
	goto L22
L25:
	;
	if base.F64_gt(v94, v87) != 0 {
		v187 = v6
		goto L1
	} else {
		goto L28
	}
L26:
	;
	v92 = *(*float64)(unsafe.Add(mBase, uint32(v70+v30<<(uint(int32(3))%32))))
	if base.F64_lt(v71, v92) != 0 {
		v94 = v71
		goto L25
	} else {
		goto L27
	}
L27:
	;
	v94 = v92
	goto L25
L28:
	;
	v98 = v43 + int32(1)
	if v98 != v32 {
		v43 = v98
		goto L13
	} else {
		goto L29
	}
L29:
	;
	goto L14
L30:
	;
	return int32(1)
L31:
	;
	goto L32
L32:
	;
	v129 = v32
	goto L33
L33:
	;
	v141 = v28 + int32(8) + v129<<(uint(int32(3))%32)
	v142 = *(*float64)(unsafe.Add(mBase, uint32(v141)))
	if base.B2i32(v100 < int32(0)) == int32(0) {
		goto L37
	} else {
		goto L38
	}
L34:
	;
	v187 = int32(0)
	goto L1
L35:
	;
	goto L34
L36:
	;
	if base.F64_lt(v168, float64(0)) != 0 {
		goto L35
	} else {
		goto L46
	}
L37:
	;
	v145 = int32(0)
	v150 = *(*float64)(unsafe.Add(mBase, uint32(v141+v116<<(uint(int32(3))%32))))
	if base.F64_lt(v142, v150) != 0 {
		goto L40
	} else {
		goto L41
	}
L38:
	;
	goto L39
L39:
	;
	if base.F64_gt(v142, float64(0)) != 0 {
		goto L35
	} else {
		goto L45
	}
L40:
	;
	v152 = v145
	goto L42
L41:
	;
	v152 = v100
	goto L42
L42:
	;
	v156 = *(*float64)(unsafe.Add(mBase, uint32(v141+v152<<(uint(int32(3))%32))))
	if base.F64_gt(v156, float64(0)) != 0 {
		v187 = v145
		goto L1
	} else {
		goto L43
	}
L43:
	;
	v162 = *(*float64)(unsafe.Add(mBase, uint32(v141+v100<<(uint(int32(3))%32))))
	if base.F64_gt(v142, v162) == int32(0) {
		v168 = v162
		goto L36
	} else {
		goto L44
	}
L44:
	;
	v168 = v142
	goto L36
L45:
	;
	v168 = v142
	goto L36
L46:
	;
	v172 = int32(1)
	v174 = v129 + v172
	if v116 != v174 {
		v129 = v174
		goto L33
	} else {
		goto L47
	}
L47:
	;
	v187 = v172
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
					F_yy_fatal_error_6(m, int32(_a_F_cube_yy_create_buffer_0))
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
					v23 = *(*int32)(unsafe.Add(mBase, _c_F_cube_yy_create_buffer[0]))
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
					*(*int32)(unsafe.Add(mBase, _c_F_cube_yy_create_buffer[0])) = v23
					return v8
				}
			}
		} else {
			F_yy_fatal_error_6(m, int32(_a_F_cube_yy_create_buffer_0))
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
	var v22 int32
	_ = v22
	var v26 int32
	_ = v26
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v50 int32
	_ = v50
	var v53 int32
	_ = v53
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
				v22 = m.ExcPending
				if v22 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_cube_yyerror_0), int32(0))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = l4
						F_errdetail(m, int32(_a_F_cube_yyerror_1), v9)
						mBase = m.M
						v30 = m.ExcPending
						if v30 != 0 {
							return
						} else {
							v50 = int32(85)
							F_errsave_finish(m, l2, int32(_a_F_cube_yyerror_2), v50, int32(_a_F_cube_yyerror_3))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
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
				v36 = m.ExcPending
				if v36 != 0 {
					return
				} else {
					F_errmsg(m, int32(_a_F_cube_yyerror_0), int32(0))
					mBase = m.M
					v40 = m.ExcPending
					if v40 != 0 {
						return
					} else {
						v41 = *(*int32)(unsafe.Add(mBase, uint32(l3)+80))
						*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v41
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = l4
						F_errdetail(m, int32(_a_F_cube_yyerror_4), v9+int32(16))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return
						} else {
							v50 = int32(93)
							F_errsave_finish(m, l2, int32(_a_F_cube_yyerror_2), v50, int32(_a_F_cube_yyerror_3))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
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
			F_yy_fatal_error_6(m, int32(_a_F_cube_yyset_lineno_0))
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
		F_yy_fatal_error_6(m, int32(_a_F_cube_yyset_lineno_0))
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
