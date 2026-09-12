package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_numeric_abbrev_abort(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v22 int32
	_ = v22
	var v23 float64
	_ = v23
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v41 float64
	_ = v41
	var v43 int32
	_ = v43
	var v45 int32
	_ = v45
	var v50 float64
	_ = v50
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 float64
	_ = v53
	var v56 int32
	_ = v56
	var v57 float64
	_ = v57
	var v62 float64
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 float64
	_ = v69
	var v71 int32
	_ = v71
	var v78 float64
	_ = v78
	var v81 int32
	_ = v81
	var v82 float64
	_ = v82
	var v85 float64
	_ = v85
	var v86 float64
	_ = v86
	var v87 float64
	_ = v87
	var v88 float64
	_ = v88
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v150 int32
	_ = v150
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v165 int32
	_ = v165
	var v172 int32
	_ = v172
	var v181 float64
	_ = v181
	var v183 float64
	_ = v183
	var v185 float64
	_ = v185
	var v191 float64
	_ = v191
	var v208 float64
	_ = v208
	var v212 float64
	_ = v212
	var v231 float64
	_ = v231
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v243 int32
	_ = v243
	var v246 int64
	_ = v246
	var v252 int32
	_ = v252
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v262 int32
	_ = v262
	var v263 int64
	_ = v263
	var v270 int32
	_ = v270
	var v277 int32
	_ = v277
	var v278 int32
	_ = v278
	var v281 int64
	_ = v281
	var v295 int32
	_ = v295
	var v300 int32
	_ = v300
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v311 int64
	_ = v311
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v326 int32
	_ = v326
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(96)
	m.G0 = v10
	if l0 < int32(10000) {
		v326 = v3
		m.G0 = v10 + int32(96)
		return v326
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l1)+12))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
		if v15 < int64(10000) {
			v326 = v3
			m.G0 = v10 + int32(96)
			return v326
		} else {
			v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)))
			if v18 != int32(1) {
				v326 = v3
				m.G0 = v10 + int32(96)
				return v326
			} else {
				v22 = v14 + int32(24)
				v23 = float64(0)
				v25 = int32(0)
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v22)+4))
				if v32 != 0 {
					v33 = int32(1)
					v35 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
					if v32 == v33 {
						v69 = v23
						v71 = v25
					} else {
						v41 = v23
						v43 = v25
						v45 = v25
						for {
							v50 = float64(1)
							v51 = v43 + v35
							v52 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51)+1)))
							v53 = F_ldexp(m, v50, v52)
							mBase = m.M
							v56 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v51))))
							v57 = F_ldexp(m, v50, v56)
							mBase = m.M
							v62 = base.F64_add(base.F64_add(v41, base.F64_div(v50, v57)), base.F64_div(v50, v53))
							v63 = int32(2)
							v64 = v43 + v63
							v66 = v45 + v63
							if v66 != v32&int32(-2) {
								v41 = v62
								v43 = v64
								v45 = v66
								continue
							} else {
								break
							}
							break
						}
						v69 = v62
						v71 = v64
					}
					if v32&v33 != 0 {
						v78 = float64(1)
						v81 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v71+v35))))
						v82 = F_ldexp(m, v78, v81)
						mBase = m.M
						v85 = base.F64_add(v69, base.F64_div(v78, v82))
					} else {
						v85 = v69
					}
					v86 = *(*float64)(unsafe.Add(mBase, uint32(v22)+8))
					v87 = base.F64_div(v86, v85)
					v88 = base.F64_convert_i32_u(v32)
					if base.F64_le(v87, base.F64_mul(v88, float64(2.5))) == int32(0) {
						v191 = v87
						if base.F64_gt(v191, float64(1.4316557653333333e+08)) == int32(0) {
							v212 = v191
						} else {
							v208 = F_log(m, base.F64_add(base.F64_mul(v191, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v212 = base.F64_mul(v208, float64(-4.294967296e+09))
						}
						v231 = v212
					} else {
						v95 = v32 & int32(3)
						v96 = *(*int32)(unsafe.Add(mBase, uint32(v22)+16))
						v97 = int32(0)
						if base.Ui32(int32(4)) <= base.Ui32(v32) {
							v108 = v97
							v109 = int32(0)
							v110 = v97
							for {
								v115 = v108 + v96
								v116 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115))))
								v117 = int32(0)
								v120 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+1)))
								v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+2)))
								v128 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v115)+3)))
								v131 = v110 + base.B2i32(v116 == v117) + base.B2i32(v120 == v117) + base.B2i32(v124 == v117) + base.B2i32(v128 == v117)
								v132 = int32(4)
								v133 = v108 + v132
								v135 = v109 + v132
								if v135 != v32&int32(-4) {
									v108 = v133
									v109 = v135
									v110 = v131
									continue
								} else {
									break
								}
								break
							}
							v140 = v133
							v142 = v131
						} else {
							v140 = v97
							v142 = v97
						}
						if v95 != 0 {
							v150 = v140
							v152 = v142
							v153 = v97
							for {
								v158 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v150+v96))))
								v161 = v152 + base.B2i32(v158 == int32(0))
								v162 = int32(1)
								v165 = v153 + v162
								if v165 != v95 {
									v150 = v150 + v162
									v152 = v161
									v153 = v165
									continue
								} else {
									break
								}
								break
							}
							v172 = v161
						} else {
							v172 = v142
						}
						if v172 == int32(0) {
							v212 = v87
							v231 = v212
						} else {
							v181 = F_log(m, base.F64_div(v88, base.F64_convert_i32_s(v172)))
							mBase = m.M
							v231 = base.F64_mul(v181, v88)
						}
					}
				} else {
					v183 = *(*float64)(unsafe.Add(mBase, uint32(v22)+8))
					v185 = base.F64_div(v183, float64(0))
					if base.F64_le(v185, base.F64_mul(base.F64_convert_i32_u(v32), float64(2.5))) != 0 {
						v212 = v185
					} else {
						v191 = v185
						if base.F64_gt(v191, float64(1.4316557653333333e+08)) == int32(0) {
							v212 = v191
						} else {
							v208 = F_log(m, base.F64_add(base.F64_mul(v191, float64(-2.3283064365386963e-10)), float64(1)))
							mBase = m.M
							v212 = base.F64_mul(v208, float64(-4.294967296e+09))
						}
					}
					v231 = v212
				}
				if base.F64_gt(v231, float64(100000)) != 0 {
					v235 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
					if v235 != int32(1) {
						v259 = int32(0)
						*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v259)
						v326 = v3
						m.G0 = v10 + int32(96)
						return v326
					} else {
						v240 = F_errstart(m, int32(15), int32(0))
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return int32(0)
						} else {
							if v240 == int32(0) {
								v259 = int32(0)
								*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v259)
								v326 = v3
								m.G0 = v10 + int32(96)
								return v326
							} else {
								v246 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = l0
								*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v246
								*(*float64)(unsafe.Add(mBase, uint32(v10))) = v231
								F_errmsg_internal(m, int32(697901), v10)
								mBase = m.M
								v252 = m.ExcPending
								if v252 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(520760), int32(2255), int32(86786))
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int32(0)
									} else {
										v259 = int32(0)
										*(*uint8)(unsafe.Add(mBase, uint32(v14)+16)) = uint8(v259)
										v326 = v3
										m.G0 = v10 + int32(96)
										return v326
									}
								}
							}
						}
					}
				} else {
					v262 = int32(*(*uint8)(unsafe.Add(mBase, _consts[59])))
					v263 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
					if base.F64_gt(base.F64_add(base.F64_div(base.F64_convert_i64_s(v263), float64(10000)), float64(0.5)), v231) != 0 {
						v270 = int32(1)
						if v262&v270 == int32(0) {
							v326 = v270
							m.G0 = v10 + int32(96)
							return v326
						} else {
							v277 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v278 = m.ExcPending
							if v278 != 0 {
								return int32(0)
							} else {
								if v277 == int32(0) {
									v326 = v270
									m.G0 = v10 + int32(96)
									return v326
								} else {
									v281 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+56)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v10)+48)) = v281
									*(*float64)(unsafe.Add(mBase, uint32(v10)+32)) = v231
									*(*float64)(unsafe.Add(mBase, uint32(v10)+40)) = base.F64_add(base.F64_div(base.F64_convert_i64_s(v281), float64(10000)), float64(0.5))
									F_errmsg_internal(m, int32(698518), v10+int32(32))
									mBase = m.M
									v295 = m.ExcPending
									if v295 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(520760), int32(2276), int32(86786))
										mBase = m.M
										v300 = m.ExcPending
										if v300 != 0 {
											return int32(0)
										} else {
											v326 = v270
											m.G0 = v10 + int32(96)
											return v326
										}
									}
								}
							}
						}
					} else {
						if v262&int32(1) == int32(0) {
							v326 = v3
							m.G0 = v10 + int32(96)
							return v326
						} else {
							v307 = F_errstart(m, int32(15), int32(0))
							mBase = m.M
							v308 = m.ExcPending
							if v308 != 0 {
								return int32(0)
							} else {
								if v307 == int32(0) {
									v326 = v3
									m.G0 = v10 + int32(96)
									return v326
								} else {
									v311 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v10)+80)) = l0
									*(*int64)(unsafe.Add(mBase, uint32(v10)+72)) = v311
									*(*float64)(unsafe.Add(mBase, uint32(v10)+64)) = v231
									F_errmsg_internal(m, int32(698153), v10-int32(-64))
									mBase = m.M
									v319 = m.ExcPending
									if v319 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(520760), int32(2284), int32(86786))
										mBase = m.M
										v324 = m.ExcPending
										if v324 != 0 {
											return int32(0)
										} else {
											v326 = v3
											m.G0 = v10 + int32(96)
											return v326
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
func F_numeric_accum(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v12 != 0 {
			v64 = v12
			v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v66 == int32(0) {
				v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v70 = F_pg_detoast_datum(m, v69)
				mBase = m.M
				v71 = m.ExcPending
				if v71 != 0 {
					return int32(0)
				} else {
					F_do_numeric_accum(m, v64, v70)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(16)
						return v64
					}
				}
			} else {
				m.G0 = v7 + int32(16)
				return v64
			}
		} else {
			v15 = v7 + int32(12)
			v16 = int32(0)
			v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v17 == v16 {
				v34 = int32(0)
				if v15 == v34 {
					v42 = v34
				} else {
					v37 = v34
					v38 = v16
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
				}
				v45 = v42
			} else {
				v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				switch v20 - int32(429) {
				case 0:
					if v15 == int32(0) {
						v45 = int32(1)
					} else {
						v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
						v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
						v37 = v27
						v38 = int32(1)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
						v45 = v42
					}
				case 1:
					if v15 == int32(0) {
						v45 = int32(2)
					} else {
						v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
						v37 = v32
						v38 = int32(2)
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
						v45 = v42
					}
				default:
					v34 = int32(0)
					if v15 == v34 {
						v42 = v34
					} else {
						v37 = v34
						v38 = v16
						*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
						v42 = v38
					}
					v45 = v42
				}
			}
			if v45 == int32(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int32(0)
				} else {
					F_errmsg_internal(m, int32(66371), int32(0))
					mBase = m.M
					v85 = m.ExcPending
					if v85 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(520760), int32(4943), int32(368786))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							base.Wasm_trap_unreachable()
							for {
							}
						}
					}
				}
			} else {
				v48 = int32(4549024)
				v49 = *(*int32)(unsafe.Add(mBase, _consts[3]))
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
				v54 = F_palloc0(m, int32(112))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int32(0)
				} else {
					v58 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v58)
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
					*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v60
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v49
					v64 = v54
					v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
					if v66 == int32(0) {
						v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
						v70 = F_pg_detoast_datum(m, v69)
						mBase = m.M
						v71 = m.ExcPending
						if v71 != 0 {
							return int32(0)
						} else {
							F_do_numeric_accum(m, v64, v70)
							mBase = m.M
							v73 = m.ExcPending
							if v73 != 0 {
								return int32(0)
							} else {
								m.G0 = v7 + int32(16)
								return v64
							}
						}
					} else {
						m.G0 = v7 + int32(16)
						return v64
					}
				}
			}
		}
	} else {
		v15 = v7 + int32(12)
		v16 = int32(0)
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		if v17 == v16 {
			v34 = int32(0)
			if v15 == v34 {
				v42 = v34
			} else {
				v37 = v34
				v38 = v16
				*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
				v42 = v38
			}
			v45 = v42
		} else {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
			switch v20 - int32(429) {
			case 0:
				if v15 == int32(0) {
					v45 = int32(1)
				} else {
					v26 = *(*int32)(unsafe.Add(mBase, uint32(v17)+168))
					v27 = *(*int32)(unsafe.Add(mBase, uint32(v26)+20))
					v37 = v27
					v38 = int32(1)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
					v45 = v42
				}
			case 1:
				if v15 == int32(0) {
					v45 = int32(2)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v17)+368))
					v37 = v32
					v38 = int32(2)
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
					v45 = v42
				}
			default:
				v34 = int32(0)
				if v15 == v34 {
					v42 = v34
				} else {
					v37 = v34
					v38 = v16
					*(*int32)(unsafe.Add(mBase, uint32(v15))) = v37
					v42 = v38
				}
				v45 = v42
			}
		}
		if v45 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v81 = m.ExcPending
			if v81 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(66371), int32(0))
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(520760), int32(4943), int32(368786))
					mBase = m.M
					v90 = m.ExcPending
					if v90 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v48 = int32(4549024)
			v49 = *(*int32)(unsafe.Add(mBase, _consts[3]))
			v51 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
			*(*int32)(unsafe.Add(mBase, _consts[3])) = v51
			v54 = F_palloc0(m, int32(112))
			mBase = m.M
			v57 = m.ExcPending
			if v57 != 0 {
				return int32(0)
			} else {
				v58 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v54))) = uint8(v58)
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
				*(*int32)(unsafe.Add(mBase, uint32(v54)+4)) = v60
				*(*int32)(unsafe.Add(mBase, _consts[3])) = v49
				v64 = v54
				v66 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
				if v66 == int32(0) {
					v69 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
					v70 = F_pg_detoast_datum(m, v69)
					mBase = m.M
					v71 = m.ExcPending
					if v71 != 0 {
						return int32(0)
					} else {
						F_do_numeric_accum(m, v64, v70)
						mBase = m.M
						v73 = m.ExcPending
						if v73 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(16)
							return v64
						}
					}
				} else {
					m.G0 = v7 + int32(16)
					return v64
				}
			}
		}
	}
}
func F_numeric_add(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = F_numeric_add_opt_error(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_numeric_ceil(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v39 int32
	_ = v39
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v70 int32
	_ = v70
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v96 int32
	_ = v96
	var v98 int64
	_ = v98
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v203 int32
	_ = v203
	var v205 int32
	_ = v205
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v219 int32
	_ = v219
	var v228 int32
	_ = v228
	var v230 int32
	_ = v230
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v248 int32
	_ = v248
	var v255 int32
	_ = v255
	var v260 int32
	_ = v260
	var v268 int32
	_ = v268
	var v272 int32
	_ = v272
	var v277 int32
	_ = v277
	var v281 int32
	_ = v281
	var v287 int32
	_ = v287
	var v298 int32
	_ = v298
	var v308 int32
	_ = v308
	var v312 int32
	_ = v312
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v335 int64
	_ = v335
	var v337 int64
	_ = v337
	var v343 int32
	_ = v343
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v20 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v22 = int32(base.Ui32(v20) >> (uint(int32(2)) % 32))
	v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	if base.Ui32(int32(49152)) <= base.Ui32(v23) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v13 + int32(48)
	return v351
L4:
	;
	v26 = F_palloc(m, v22)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v35 = base.I32_extend16_s(v23)
	v37 = base.B2i32(int32(0) <= v35)
	if int32(0) <= v35 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v28 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v30 = int32(base.Ui32(v28) >> (uint(int32(2)) % 32))
	if v30 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v351 = v26
	goto L3
L9:
	;
	v31 = F__emscripten_memcpy_bulkmem(m, v26, v16, v30)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v38 = int32(-8)
	goto L14
L13:
	;
	v38 = int32(-6)
	goto L14
L14:
	;
	v39 = v22 + v38
	v41 = int32(base.Ui32(v39) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = v41
	if int32(0) <= v35 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+6)))
	v53 = v43
	goto L17
L16:
	;
	v53 = v23<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v23&int32(63)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v53
	v55 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v55
	v60 = base.B2i32(v35 < v55)
	if v35 < v55 {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v61 = int32(6)
	goto L20
L19:
	;
	v61 = int32(8)
	goto L20
L20:
	;
	v62 = v16 + v61
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v62
	if v35 < v55 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v70 = int32(base.Ui32(v23)>>(uint(int32(7))%32)) & int32(63)
	goto L23
L22:
	;
	v70 = v23 & int32(16383)
	goto L23
L23:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+12)) = v70
	v77 = v23 & int32(49152)
	if v77 == int32(32768) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v80 = v23 << (uint(int32(1)) % 32) & int32(16384)
	goto L26
L25:
	;
	v80 = v77
	goto L26
L26:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+8)) = v80
	v83 = v39 & int32(-2)
	v86 = F_palloc(m, v83+int32(2))
	mBase = m.M
	v87 = m.ExcPending
	if v87 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v88 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v86))) = uint16(v88)
	if base.Ui32(int32(2)) <= base.Ui32(v39) {
		goto L28
	} else {
		goto L29
	}
L28:
	;
	if v83 != 0 {
		goto L32
	} else {
		goto L33
	}
L29:
	;
	goto L30
L30:
	;
	v96 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	*(*int32)(unsafe.Add(mBase, uint32(v13)+32)) = v96
	v98 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+24)) = v98
	*(*int32)(unsafe.Add(mBase, uint32(v13)+40)) = v86
	v101 = int32(2)
	v102 = v86 + v101
	*(*int32)(unsafe.Add(mBase, uint32(v13)+44)) = v102
	v104 = *(*int32)(unsafe.Add(mBase, uint32(v13)+28))
	v105 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v13)+36)) = v105
	v108 = v104 << (uint(v101) % 32)
	if v108+int32(4) <= v105 {
		goto L36
	} else {
		goto L37
	}
L31:
	;
	goto L30
L32:
	;
	v94 = F__emscripten_memcpy_bulkmem(m, v86+int32(2), v62, v83)
	mBase = m.M
	goto L34
L33:
	;
	goto L34
L34:
	;
	goto L31
L35:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+24)) = v127
	if v80 != 0 {
		v319 = v127
		goto L42
	} else {
		goto L43
	}
L36:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v13)+28)) = int64(0)
	v115 = int32(0)
	v125 = v115
	v127 = v115
	goto L35
L37:
	;
	goto L38
L38:
	;
	v120 = base.I32_div_s(v108+int32(7), int32(4))
	v121 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	if v120 < v121 {
		goto L39
	} else {
		goto L40
	}
L39:
	;
	v123 = v120
	goto L41
L40:
	;
	v123 = v121
	goto L41
L41:
	;
	v125 = v104
	v127 = v123
	goto L35
L42:
	;
	v321 = v319 << (uint(int32(1)) % 32)
	v324 = F_palloc(m, v321+int32(2))
	mBase = m.M
	v325 = m.ExcPending
	if v325 != 0 {
		goto L1
	} else {
		goto L100
	}
L43:
	;
	if base.Ui32(v39) <= base.Ui32(int32(1)) {
		goto L45
	} else {
		goto L46
	}
L44:
	;
	v312 = v13 + int32(24)
	F_add_var(m, v312, int32(1766504), v312)
	mBase = m.M
	v317 = m.ExcPending
	if v317 != 0 {
		goto L1
	} else {
		goto L99
	}
L45:
	;
	if v127 != 0 {
		goto L44
	} else {
		goto L48
	}
L46:
	;
	goto L47
L47:
	;
	if v127 == int32(0) {
		goto L44
	} else {
		goto L49
	}
L48:
	;
	v319 = int32(0)
	goto L42
L49:
	;
	v134 = *(*int32)(unsafe.Add(mBase, uint32(v13)+32))
	if v134 == int32(16384) {
		goto L44
	} else {
		goto L50
	}
L50:
	;
	v137 = int32(0)
	if base.B2i32(v125 < v53)&base.B2i32(v137 < v41) == v137 {
		goto L54
	} else {
		goto L55
	}
L51:
	;
	if v308 == int32(0) {
		v319 = v127
		goto L42
	} else {
		goto L98
	}
L52:
	;
	v308 = v298
	goto L51
L53:
	;
	if v125 <= v168 {
		v203 = v125
		v205 = v137
		goto L62
	} else {
		goto L63
	}
L54:
	;
	v168 = v53
	v172 = v137
	goto L53
L55:
	;
	goto L56
L56:
	;
	v149 = v53
	v153 = v137
	goto L57
L57:
	;
	v159 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+v153<<(uint(int32(1))%32)))))
	if v159 != 0 {
		v298 = int32(1)
		goto L52
	} else {
		goto L59
	}
L58:
	;
	v168 = v163
	v172 = v161
	goto L53
L59:
	;
	v160 = int32(1)
	v161 = v153 + v160
	v163 = v149 - v160
	if v163 <= v125 {
		v168 = v163
		v172 = v161
		goto L53
	} else {
		goto L60
	}
L60:
	;
	if v161 < v41 {
		v149 = v163
		v153 = v161
		goto L57
	} else {
		goto L61
	}
L61:
	;
	goto L58
L62:
	;
	if v168 != v203 {
		v244 = v172
		v245 = v205
		goto L70
	} else {
		goto L71
	}
L63:
	;
	if v127 <= int32(0) {
		v203 = v125
		v205 = v137
		goto L62
	} else {
		goto L64
	}
L64:
	;
	v184 = v125
	v186 = v137
	goto L65
L65:
	;
	v191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+v186<<(uint(int32(1))%32)))))
	if v191 != 0 {
		v298 = int32(-1)
		goto L52
	} else {
		goto L67
	}
L66:
	;
	v203 = v195
	v205 = v193
	goto L62
L67:
	;
	v192 = int32(1)
	v193 = v186 + v192
	v195 = v184 - v192
	if v195 <= v168 {
		v203 = v195
		v205 = v193
		goto L62
	} else {
		goto L68
	}
L68:
	;
	if v193 < v127 {
		v184 = v195
		v186 = v193
		goto L65
	} else {
		goto L69
	}
L69:
	;
	goto L66
L70:
	;
	if v41 < v244 {
		goto L80
	} else {
		goto L81
	}
L71:
	;
	v214 = v172
	v215 = v205
	goto L72
L72:
	;
	if v41 <= v214 {
		v244 = v214
		v245 = v215
		goto L70
	} else {
		goto L74
	}
L73:
	;
	if base.I32_extend16_s(v230) < base.I32_extend16_s(v228) {
		goto L77
	} else {
		goto L78
	}
L74:
	;
	if v127 <= v215 {
		v244 = v214
		v245 = v215
		goto L70
	} else {
		goto L75
	}
L75:
	;
	v219 = int32(1)
	v228 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+v214<<(uint(v219)%32)))))
	v230 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v215<<(uint(v219)%32)+v102))))
	if v228 == v230 {
		v214 = v214 + v219
		v215 = v215 + v219
		goto L72
	} else {
		goto L76
	}
L76:
	;
	goto L73
L77:
	;
	v237 = int32(1)
	goto L79
L78:
	;
	v237 = int32(-1)
	goto L79
L79:
	;
	v308 = v237
	goto L51
L80:
	;
	v248 = v244
	goto L82
L81:
	;
	v248 = v41
	goto L82
L82:
	;
	v255 = v244
	goto L83
L83:
	;
	if v248 == v255 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v298 = v281
	goto L52
L85:
	;
	if v127 < v245 {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	goto L87
L87:
	;
	v281 = int32(1)
	v287 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v62+v255<<(uint(v281)%32)))))
	if v287 == int32(0) {
		v255 = v255 + v281
		goto L83
	} else {
		goto L97
	}
L88:
	;
	v260 = v245
	goto L90
L89:
	;
	v260 = v127
	goto L90
L90:
	;
	v268 = v245
	goto L91
L91:
	;
	if v260 == v268 {
		goto L93
	} else {
		goto L94
	}
L92:
	;
	v298 = int32(-1)
	goto L52
L93:
	;
	v308 = int32(0)
	goto L51
L94:
	;
	goto L95
L95:
	;
	v272 = int32(1)
	v277 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+v268<<(uint(v272)%32)))))
	if v277 == int32(0) {
		v268 = v268 + v272
		goto L91
	} else {
		goto L96
	}
L96:
	;
	goto L92
L97:
	;
	goto L84
L98:
	;
	goto L44
L99:
	;
	v318 = *(*int32)(unsafe.Add(mBase, uint32(v13)+24))
	v319 = v318
	goto L42
L100:
	;
	v326 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v324))) = uint16(v326)
	if v326 < v319 {
		goto L101
	} else {
		goto L102
	}
L101:
	;
	v332 = *(*int32)(unsafe.Add(mBase, uint32(v13)+44))
	if v321 != 0 {
		goto L105
	} else {
		goto L106
	}
L102:
	;
	goto L103
L103:
	;
	v335 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	*(*int64)(unsafe.Add(mBase, uint32(v13)+8)) = v335
	v337 = *(*int64)(unsafe.Add(mBase, uint32(v13)+24))
	*(*int64)(unsafe.Add(mBase, uint32(v13))) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v324
	*(*int32)(unsafe.Add(mBase, uint32(v13)+20)) = v324 + int32(2)
	v343 = *(*int32)(unsafe.Add(mBase, uint32(v13)+40))
	if v343 != 0 {
		goto L108
	} else {
		goto L109
	}
L104:
	;
	goto L103
L105:
	;
	v333 = F__emscripten_memcpy_bulkmem(m, v324+int32(2), v332, v321)
	mBase = m.M
	goto L107
L106:
	;
	goto L107
L107:
	;
	goto L104
L108:
	;
	F_pfree(m, v343)
	mBase = m.M
	v345 = m.ExcPending
	if v345 != 0 {
		goto L1
	} else {
		goto L111
	}
L109:
	;
	goto L110
L110:
	;
	v347 = F_make_result_opt_error(m, v13, int32(0))
	mBase = m.M
	v348 = m.ExcPending
	if v348 != 0 {
		goto L1
	} else {
		goto L112
	}
L111:
	;
	goto L110
L112:
	;
	F_pfree(m, v324)
	mBase = m.M
	v350 = m.ExcPending
	if v350 != 0 {
		goto L1
	} else {
		goto L113
	}
L113:
	;
	v351 = v347
	goto L3
}
func F_numeric_cmp_abbrev(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	return base.B2i32(l0 < l1) - base.B2i32(l1 < l0)
}
func F_numeric_combine(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
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
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int32
	_ = v78
	var v80 int64
	_ = v80
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v116 int32
	_ = v116
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v150 int64
	_ = v150
	var v151 int64
	_ = v151
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v158 int64
	_ = v158
	var v159 int64
	_ = v159
	var v162 int64
	_ = v162
	var v163 int64
	_ = v163
	var v166 int64
	_ = v166
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v177 int64
	_ = v177
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v186 int32
	_ = v186
	var v187 int64
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	var v210 int64
	_ = v210
	var v221 int32
	_ = v221
	var v227 int32
	_ = v227
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v254 int32
	_ = v254
	var v258 int32
	_ = v258
	var v263 int32
	_ = v263
	v2 = int32(0)
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v13 = v10 + int32(4)
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v15 == v2 {
		v32 = int32(0)
		if v13 == v32 {
			v40 = v32
		} else {
			v35 = v32
			v36 = v2
			*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
			v40 = v36
		}
		v43 = v40
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(v15)))
		switch v18 - int32(429) {
		case 0:
			if v13 == int32(0) {
				v43 = int32(1)
			} else {
				v24 = *(*int32)(unsafe.Add(mBase, uint32(v15)+168))
				v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+20))
				v35 = v25
				v36 = int32(1)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
				v40 = v36
				v43 = v40
			}
		case 1:
			if v13 == int32(0) {
				v43 = int32(2)
			} else {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(v15)+368))
				v35 = v30
				v36 = int32(2)
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
				v40 = v36
				v43 = v40
			}
		default:
			v32 = int32(0)
			if v13 == v32 {
				v40 = v32
			} else {
				v35 = v32
				v36 = v2
				*(*int32)(unsafe.Add(mBase, uint32(v13))) = v35
				v40 = v36
			}
			v43 = v40
		}
	}
	if v43 != 0 {
		v44 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
		if v44 == int32(0) {
			v47 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v48 = v47
		} else {
			v48 = v2
		}
		v49 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
		if v49 != 0 {
			v242 = v48
			m.G0 = v10 + int32(32)
			return v242
		} else {
			v50 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
			if v50 == int32(0) {
				v242 = v48
				m.G0 = v10 + int32(32)
				return v242
			} else {
				if v48 == int32(0) {
					v55 = int32(4549024)
					v56 = *(*int32)(unsafe.Add(mBase, _consts[3]))
					v58 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
					*(*int32)(unsafe.Add(mBase, _consts[3])) = v58
					v61 = F_palloc0(m, int32(112))
					mBase = m.M
					v64 = m.ExcPending
					if v64 != 0 {
						return int32(0)
					} else {
						v65 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(v61))) = uint8(v65)
						v68 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						*(*int32)(unsafe.Add(mBase, uint32(v61)+4)) = v68
						v70 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v61)+8)) = v70
						v72 = *(*int64)(unsafe.Add(mBase, uint32(v50)+88))
						*(*int64)(unsafe.Add(mBase, uint32(v61)+88)) = v72
						v74 = *(*int64)(unsafe.Add(mBase, uint32(v50)+96))
						*(*int64)(unsafe.Add(mBase, uint32(v61)+96)) = v74
						v76 = *(*int64)(unsafe.Add(mBase, uint32(v50)+104))
						*(*int64)(unsafe.Add(mBase, uint32(v61)+104)) = v76
						v78 = *(*int32)(unsafe.Add(mBase, uint32(v50)+72))
						*(*int32)(unsafe.Add(mBase, uint32(v61)+72)) = v78
						v80 = *(*int64)(unsafe.Add(mBase, uint32(v50)+80))
						*(*int64)(unsafe.Add(mBase, uint32(v61)+80)) = v80
						v82 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
						v85 = F_palloc(m, v82<<(uint(int32(2))%32))
						mBase = m.M
						v86 = m.ExcPending
						if v86 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v61)+36)) = v85
							v88 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
							v91 = F_palloc(m, v88<<(uint(int32(2))%32))
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v61)+40)) = v91
								v94 = *(*int32)(unsafe.Add(mBase, uint32(v61)+36))
								v95 = *(*int32)(unsafe.Add(mBase, uint32(v50)+36))
								v96 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
								v98 = v96 << (uint(int32(2)) % 32)
								if v98 != 0 {
									v99 = F__emscripten_memcpy_bulkmem(m, v94, v95, v98)
									mBase = m.M
								} else {
								}
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v61)+40))
								v102 = *(*int32)(unsafe.Add(mBase, uint32(v50)+40))
								v103 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
								v105 = v103 << (uint(int32(2)) % 32)
								if v105 != 0 {
									v106 = F__emscripten_memcpy_bulkmem(m, v101, v102, v105)
									mBase = m.M
								} else {
								}
								v108 = *(*int32)(unsafe.Add(mBase, uint32(v50)+28))
								*(*int32)(unsafe.Add(mBase, uint32(v61)+28)) = v108
								v110 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
								*(*int32)(unsafe.Add(mBase, uint32(v61)+16)) = v110
								v112 = *(*int32)(unsafe.Add(mBase, uint32(v50)+20))
								*(*int32)(unsafe.Add(mBase, uint32(v61)+20)) = v112
								v114 = *(*int32)(unsafe.Add(mBase, uint32(v50)+24))
								*(*int32)(unsafe.Add(mBase, uint32(v61)+24)) = v114
								v116 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
								v119 = F_palloc(m, v116<<(uint(int32(2))%32))
								mBase = m.M
								v120 = m.ExcPending
								if v120 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v61)+64)) = v119
									v122 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
									v125 = F_palloc(m, v122<<(uint(int32(2))%32))
									mBase = m.M
									v126 = m.ExcPending
									if v126 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v61)+68)) = v125
										v128 = *(*int32)(unsafe.Add(mBase, uint32(v61)+64))
										v129 = *(*int32)(unsafe.Add(mBase, uint32(v50)+64))
										v130 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
										v132 = v130 << (uint(int32(2)) % 32)
										if v132 != 0 {
											v133 = F__emscripten_memcpy_bulkmem(m, v128, v129, v132)
											mBase = m.M
										} else {
										}
										v135 = *(*int32)(unsafe.Add(mBase, uint32(v61)+68))
										v136 = *(*int32)(unsafe.Add(mBase, uint32(v50)+68))
										v137 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
										v139 = v137 << (uint(int32(2)) % 32)
										if v139 != 0 {
											v140 = F__emscripten_memcpy_bulkmem(m, v135, v136, v139)
											mBase = m.M
										} else {
										}
										v142 = *(*int32)(unsafe.Add(mBase, uint32(v50)+56))
										*(*int32)(unsafe.Add(mBase, uint32(v61)+56)) = v142
										v144 = *(*int32)(unsafe.Add(mBase, uint32(v50)+44))
										*(*int32)(unsafe.Add(mBase, uint32(v61)+44)) = v144
										v146 = *(*int32)(unsafe.Add(mBase, uint32(v50)+48))
										*(*int32)(unsafe.Add(mBase, uint32(v61)+48)) = v146
										v148 = *(*int32)(unsafe.Add(mBase, uint32(v50)+52))
										*(*int32)(unsafe.Add(mBase, uint32(v61)+52)) = v148
										v234 = v61
										v236 = v56
										*(*int32)(unsafe.Add(mBase, _consts[3])) = v236
										v242 = v234
										m.G0 = v10 + int32(32)
										return v242
									}
								}
							}
						}
					}
				} else {
					v150 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
					v151 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v48)+8)) = v150 + v151
					v154 = *(*int64)(unsafe.Add(mBase, uint32(v48)+88))
					v155 = *(*int64)(unsafe.Add(mBase, uint32(v50)+88))
					*(*int64)(unsafe.Add(mBase, uint32(v48)+88)) = v154 + v155
					v158 = *(*int64)(unsafe.Add(mBase, uint32(v48)+96))
					v159 = *(*int64)(unsafe.Add(mBase, uint32(v50)+96))
					*(*int64)(unsafe.Add(mBase, uint32(v48)+96)) = v158 + v159
					v162 = *(*int64)(unsafe.Add(mBase, uint32(v48)+104))
					v163 = *(*int64)(unsafe.Add(mBase, uint32(v50)+104))
					*(*int64)(unsafe.Add(mBase, uint32(v48)+104)) = v162 + v163
					v166 = *(*int64)(unsafe.Add(mBase, uint32(v50)+8))
					if v166 <= int64(0) {
						v242 = v48
						m.G0 = v10 + int32(32)
						return v242
					} else {
						v169 = *(*int32)(unsafe.Add(mBase, uint32(v50)+72))
						v170 = *(*int32)(unsafe.Add(mBase, uint32(v48)+72))
						if v170 < v169 {
							*(*int32)(unsafe.Add(mBase, uint32(v48)+72)) = v169
							v173 = *(*int64)(unsafe.Add(mBase, uint32(v50)+80))
							*(*int64)(unsafe.Add(mBase, uint32(v48)+80)) = v173
						} else {
							if v169 != v170 {
							} else {
								v176 = *(*int64)(unsafe.Add(mBase, uint32(v48)+80))
								v177 = *(*int64)(unsafe.Add(mBase, uint32(v50)+80))
								*(*int64)(unsafe.Add(mBase, uint32(v48)+80)) = v176 + v177
							}
						}
						v180 = int32(4549024)
						v181 = *(*int32)(unsafe.Add(mBase, _consts[3]))
						v183 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						*(*int32)(unsafe.Add(mBase, _consts[3])) = v183
						v186 = v10 + int32(24)
						v187 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v186))) = v187
						v189 = int32(16)
						v190 = v10 + v189
						*(*int64)(unsafe.Add(mBase, uint32(v190))) = v187
						*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v187
						F_accum_sum_final(m, v50+v189, v10+int32(8))
						mBase = m.M
						v200 = m.ExcPending
						if v200 != 0 {
							return int32(0)
						} else {
							F_accum_sum_add(m, v48+int32(16), v10+int32(8))
							mBase = m.M
							v206 = m.ExcPending
							if v206 != 0 {
								return int32(0)
							} else {
								v207 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
								if v207 != 0 {
									F_pfree(m, v207)
									mBase = m.M
									v209 = m.ExcPending
									if v209 != 0 {
										return int32(0)
									} else {
										v210 = int64(0)
										*(*int64)(unsafe.Add(mBase, uint32(v186))) = v210
										*(*int64)(unsafe.Add(mBase, uint32(v190))) = v210
										*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v210
										F_accum_sum_final(m, v50+int32(44), v10+int32(8))
										mBase = m.M
										v221 = m.ExcPending
										if v221 != 0 {
											return int32(0)
										} else {
											F_accum_sum_add(m, v48+int32(44), v10+int32(8))
											mBase = m.M
											v227 = m.ExcPending
											if v227 != 0 {
												return int32(0)
											} else {
												v228 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
												if v228 == int32(0) {
													v234 = v48
													v236 = v181
													*(*int32)(unsafe.Add(mBase, _consts[3])) = v236
													v242 = v234
													m.G0 = v10 + int32(32)
													return v242
												} else {
													F_pfree(m, v228)
													mBase = m.M
													v232 = m.ExcPending
													if v232 != 0 {
														return int32(0)
													} else {
														v234 = v48
														v236 = v181
														*(*int32)(unsafe.Add(mBase, _consts[3])) = v236
														v242 = v234
														m.G0 = v10 + int32(32)
														return v242
													}
												}
											}
										}
									}
								} else {
									v210 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v186))) = v210
									*(*int64)(unsafe.Add(mBase, uint32(v190))) = v210
									*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v210
									F_accum_sum_final(m, v50+int32(44), v10+int32(8))
									mBase = m.M
									v221 = m.ExcPending
									if v221 != 0 {
										return int32(0)
									} else {
										F_accum_sum_add(m, v48+int32(44), v10+int32(8))
										mBase = m.M
										v227 = m.ExcPending
										if v227 != 0 {
											return int32(0)
										} else {
											v228 = *(*int32)(unsafe.Add(mBase, uint32(v186)))
											if v228 == int32(0) {
												v234 = v48
												v236 = v181
												*(*int32)(unsafe.Add(mBase, _consts[3])) = v236
												v242 = v234
												m.G0 = v10 + int32(32)
												return v242
											} else {
												F_pfree(m, v228)
												mBase = m.M
												v232 = m.ExcPending
												if v232 != 0 {
													return int32(0)
												} else {
													v234 = v48
													v236 = v181
													*(*int32)(unsafe.Add(mBase, _consts[3])) = v236
													v242 = v234
													m.G0 = v10 + int32(32)
													return v242
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v254 = m.ExcPending
		if v254 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66371), int32(0))
			mBase = m.M
			v258 = m.ExcPending
			if v258 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(520760), int32(5167), int32(389263))
				mBase = m.M
				v263 = m.ExcPending
				if v263 != 0 {
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
func F_numeric_div(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v8 = F_pg_detoast_datum(m, v7)
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v11 = F_numeric_div_opt_error(m, v3, v8, int32(0))
			mBase = m.M
			v12 = m.ExcPending
			if v12 != 0 {
				return int32(0)
			} else {
				return v11
			}
		}
	}
}
func F_numeric_exp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v30 int32
	_ = v30
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
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int32
	_ = v49
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v102 float64
	_ = v102
	var v103 int32
	_ = v103
	var v105 float64
	_ = v105
	var v106 float64
	_ = v106
	var v109 float64
	_ = v109
	var v110 float64
	_ = v110
	var v113 float64
	_ = v113
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
		v18 = base.I32_extend16_s(v17)
		if base.Ui32(int32(49152)) <= base.Ui32(v17) {
			if v18 == int32(-4096) {
				v25 = F_make_result_opt_error(m, int32(1766452), int32(0))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					v136 = v25
					m.G0 = v10 + int32(48)
					return v136
				}
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v30 = F_palloc(m, int32(base.Ui32(v27)>>(uint(int32(2))%32)))
				mBase = m.M
				v31 = m.ExcPending
				if v31 != 0 {
					return int32(0)
				} else {
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v34 = int32(base.Ui32(v32) >> (uint(int32(2)) % 32))
					if v34 != 0 {
						v35 = F__emscripten_memcpy_bulkmem(m, v30, v13, v34)
						mBase = m.M
					} else {
					}
					v136 = v30
					m.G0 = v10 + int32(48)
					return v136
				}
			}
		} else {
			v37 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
			v43 = base.B2i32(int32(0) <= v18)
			if int32(0) <= v18 {
				v44 = int32(-8)
			} else {
				v44 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = int32(base.Ui32(int32(base.Ui32(v37)>>(uint(int32(2))%32))+v44) >> (uint(int32(1)) % 32))
			if int32(0) <= v18 {
				v49 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+6)))
				v59 = v49
			} else {
				v59 = v17<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v17&int32(63)
			}
			v60 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v60
			v63 = v10 + int32(16)
			*(*int64)(unsafe.Add(mBase, uint32(v63))) = v60
			*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v59
			v67 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v67
			*(*int64)(unsafe.Add(mBase, uint32(v10))) = v60
			v74 = base.B2i32(v18 < v67)
			if v18 < v67 {
				v75 = int32(6)
			} else {
				v75 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v13 + v75
			v83 = v17 & int32(49152)
			if v83 == int32(32768) {
				v86 = v17 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v86 = v83
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = v86
			if v18 < v67 {
				v94 = int32(base.Ui32(v17)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v94 = v17 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v94
			v97 = v10 + int32(24)
			v102 = F_numericvar_to_double_no_overflow(m, v97)
			mBase = m.M
			v103 = m.ExcPending
			if v103 != 0 {
				return int32(0)
			} else {
				v105 = base.F64_mul(v102, float64(0.434294481903252))
				v106 = float64(-2000)
				if base.F64_gt(v105, v106) != 0 {
					v109 = v105
				} else {
					v109 = v106
				}
				v110 = float64(2000)
				if base.F64_lt(v109, v110) != 0 {
					v113 = v109
				} else {
					v113 = v110
				}
				if base.F64_lt(base.F64_abs(v113), float64(2.147483648e+09)) != 0 {
					v117 = base.I32_trunc_f64_s(v113)
					v119 = v117
				} else {
					v119 = int32(-2147483648)
				}
				v120 = int32(16) - v119
				if v94 < v120 {
					v122 = v120
				} else {
					v122 = v94
				}
				if int32(1000) <= v122 {
					v125 = int32(1000)
				} else {
					v125 = v122
				}
				F_exp_var(m, v97, v10, v125)
				mBase = m.M
				v127 = m.ExcPending
				if v127 != 0 {
					return int32(0)
				} else {
					v129 = F_make_result_opt_error(m, v10, int32(0))
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return int32(0)
					} else {
						v131 = *(*int32)(unsafe.Add(mBase, uint32(v63)))
						if v131 == int32(0) {
							v136 = v129
							m.G0 = v10 + int32(48)
							return v136
						} else {
							F_pfree(m, v131)
							mBase = m.M
							v135 = m.ExcPending
							if v135 != 0 {
								return int32(0)
							} else {
								v136 = v129
								m.G0 = v10 + int32(48)
								return v136
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_fast_cmp(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v150 int32
	_ = v150
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	v6 = F_pg_detoast_datum(m, l0)
	mBase = m.M
	v9 = m.ExcPending
	if v9 != 0 {
		return int32(0)
	} else {
		v10 = F_pg_detoast_datum(m, l1)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v10)+4)))
			v25 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v6)+4)))
			v26 = int32(49152)
			v27 = v25 & v26
			if v27 == v26 {
				if v25 != int32(53248) {
					if v25 != int32(49152) {
						if v24 != int32(61440) {
							v46 = int32(-1)
						} else {
							v46 = int32(0)
						}
						v165 = v46
					} else {
						v165 = base.B2i32(v24 != int32(49152))
					}
				} else {
					if v24 == int32(49152) {
						v41 = int32(-1)
					} else {
						v41 = base.B2i32(v24 != int32(53248))
					}
					v165 = v41
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v24) {
					if v24 == int32(61440) {
						v53 = int32(1)
					} else {
						v53 = int32(-1)
					}
					v165 = v53
				} else {
					v55 = v6 + int32(6)
					v60 = base.B2i32(int32(0) <= base.I32_extend16_s(v25))
					if int32(0) <= base.I32_extend16_s(v25) {
						v61 = int32(-8)
					} else {
						v61 = int32(-6)
					}
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v6)))
					if int32(0) <= base.I32_extend16_s(v25) {
						v65 = int32(*(*int16)(unsafe.Add(mBase, uint32(v55))))
						v75 = v65
					} else {
						v75 = v25<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v25&int32(63)
					}
					v76 = v61 + int32(base.Ui32(v62)>>(uint(int32(2))%32))
					v78 = v10 + int32(6)
					v83 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v84 = int32(-8)
					} else {
						v84 = int32(-6)
					}
					v85 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					if int32(0) <= base.I32_extend16_s(v24) {
						v88 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78))))
						v98 = v88
					} else {
						v98 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v99 = v84 + int32(base.Ui32(v85)>>(uint(int32(2))%32))
					v105 = v24 & int32(49152)
					if v105 == int32(32768) {
						v108 = v24 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v108 = v105
					}
					if base.Ui32(v76) <= base.Ui32(int32(1)) {
						if base.Ui32(v99) < base.Ui32(int32(2)) {
							v165 = int32(0)
						} else {
							if v108 == int32(16384) {
								v118 = int32(1)
							} else {
								v118 = int32(-1)
							}
							v165 = v118
						}
					} else {
						if v27 == int32(32768) {
							v125 = v25 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v125 = v27
						}
						if base.Ui32(v99) <= base.Ui32(int32(1)) {
							if v125 != 0 {
								v130 = int32(-1)
							} else {
								v130 = int32(1)
							}
							v165 = v130
						} else {
							if int32(0) <= base.I32_extend16_s(v25) {
								v133 = v6 + int32(8)
							} else {
								v133 = v55
							}
							v135 = int32(base.Ui32(v76) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v24) {
								v138 = v10 + int32(8)
							} else {
								v138 = v78
							}
							v140 = int32(base.Ui32(v99) >> (uint(int32(1)) % 32))
							if v125 == int32(0) {
								if v108 == int32(16384) {
									v165 = int32(1)
								} else {
									v146 = F_cmp_abs_common(m, v133, v135, v75, v138, v140, v98)
									mBase = m.M
									v165 = v146
								}
							} else {
								if v108 == int32(0) {
									v165 = int32(-1)
								} else {
									v150 = F_cmp_abs_common(m, v138, v140, v98, v133, v135, v75)
									mBase = m.M
									v165 = v150
								}
							}
						}
					}
				}
			}
			if l0 != v6 {
				F_pfree(m, v6)
				mBase = m.M
				v168 = m.ExcPending
				if v168 != 0 {
					return int32(0)
				} else {
					if l1 != v10 {
						F_pfree(m, v10)
						mBase = m.M
						v171 = m.ExcPending
						if v171 != 0 {
							return int32(0)
						} else {
							return v165
						}
					} else {
						return v165
					}
				}
			} else {
				if l1 != v10 {
					F_pfree(m, v10)
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						return v165
					}
				} else {
					return v165
				}
			}
		}
	}
}
func F_numeric_gt(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			v30 = int32(49152)
			v31 = v29 & v30
			if v31 == v30 {
				if v29 != int32(53248) {
					if v29 != int32(49152) {
						if v28 != int32(61440) {
							v50 = int32(-1)
						} else {
							v50 = int32(0)
						}
						v169 = v50
					} else {
						v169 = base.B2i32(v28 != int32(49152))
					}
				} else {
					if v28 == int32(49152) {
						v45 = int32(-1)
					} else {
						v45 = base.B2i32(v28 != int32(53248))
					}
					v169 = v45
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v28) {
					if v28 == int32(61440) {
						v57 = int32(1)
					} else {
						v57 = int32(-1)
					}
					v169 = v57
				} else {
					v59 = v7 + int32(6)
					v64 = base.B2i32(int32(0) <= base.I32_extend16_s(v29))
					if int32(0) <= base.I32_extend16_s(v29) {
						v65 = int32(-8)
					} else {
						v65 = int32(-6)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					if int32(0) <= base.I32_extend16_s(v29) {
						v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59))))
						v79 = v69
					} else {
						v79 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
					}
					v80 = v65 + int32(base.Ui32(v66)>>(uint(int32(2))%32))
					v82 = v14 + int32(6)
					v87 = base.B2i32(int32(0) <= base.I32_extend16_s(v28))
					if int32(0) <= base.I32_extend16_s(v28) {
						v88 = int32(-8)
					} else {
						v88 = int32(-6)
					}
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if int32(0) <= base.I32_extend16_s(v28) {
						v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82))))
						v102 = v92
					} else {
						v102 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
					}
					v103 = v88 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
					v109 = v28 & int32(49152)
					if v109 == int32(32768) {
						v112 = v28 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v112 = v109
					}
					if base.Ui32(v80) <= base.Ui32(int32(1)) {
						if base.Ui32(v103) < base.Ui32(int32(2)) {
							v169 = int32(0)
						} else {
							if v112 == int32(16384) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v169 = v122
						}
					} else {
						if v31 == int32(32768) {
							v129 = v29 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v129 = v31
						}
						if base.Ui32(v103) <= base.Ui32(int32(1)) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v169 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v29) {
								v137 = v7 + int32(8)
							} else {
								v137 = v59
							}
							v139 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v28) {
								v142 = v14 + int32(8)
							} else {
								v142 = v82
							}
							v144 = int32(base.Ui32(v103) >> (uint(int32(1)) % 32))
							if v129 == int32(0) {
								if v112 == int32(16384) {
									v169 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v137, v139, v79, v142, v144, v102)
									mBase = m.M
									v169 = v150
								}
							} else {
								if v112 == int32(0) {
									v169 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v142, v144, v102, v137, v139, v79)
									mBase = m.M
									v169 = v154
								}
							}
						}
					}
				}
			}
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v170 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v174 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							return base.B2i32(int32(0) < v169)
						}
					} else {
						return base.B2i32(int32(0) < v169)
					}
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v174 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						return base.B2i32(int32(0) < v169)
					}
				} else {
					return base.B2i32(int32(0) < v169)
				}
			}
		}
	}
}
func F_numeric_le(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			v30 = int32(49152)
			v31 = v29 & v30
			if v31 == v30 {
				if v29 != int32(53248) {
					if v29 != int32(49152) {
						if v28 != int32(61440) {
							v50 = int32(-1)
						} else {
							v50 = int32(0)
						}
						v169 = v50
					} else {
						v169 = base.B2i32(v28 != int32(49152))
					}
				} else {
					if v28 == int32(49152) {
						v45 = int32(-1)
					} else {
						v45 = base.B2i32(v28 != int32(53248))
					}
					v169 = v45
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v28) {
					if v28 == int32(61440) {
						v57 = int32(1)
					} else {
						v57 = int32(-1)
					}
					v169 = v57
				} else {
					v59 = v7 + int32(6)
					v64 = base.B2i32(int32(0) <= base.I32_extend16_s(v29))
					if int32(0) <= base.I32_extend16_s(v29) {
						v65 = int32(-8)
					} else {
						v65 = int32(-6)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					if int32(0) <= base.I32_extend16_s(v29) {
						v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59))))
						v79 = v69
					} else {
						v79 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
					}
					v80 = v65 + int32(base.Ui32(v66)>>(uint(int32(2))%32))
					v82 = v14 + int32(6)
					v87 = base.B2i32(int32(0) <= base.I32_extend16_s(v28))
					if int32(0) <= base.I32_extend16_s(v28) {
						v88 = int32(-8)
					} else {
						v88 = int32(-6)
					}
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if int32(0) <= base.I32_extend16_s(v28) {
						v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82))))
						v102 = v92
					} else {
						v102 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
					}
					v103 = v88 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
					v109 = v28 & int32(49152)
					if v109 == int32(32768) {
						v112 = v28 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v112 = v109
					}
					if base.Ui32(v80) <= base.Ui32(int32(1)) {
						if base.Ui32(v103) < base.Ui32(int32(2)) {
							v169 = int32(0)
						} else {
							if v112 == int32(16384) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v169 = v122
						}
					} else {
						if v31 == int32(32768) {
							v129 = v29 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v129 = v31
						}
						if base.Ui32(v103) <= base.Ui32(int32(1)) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v169 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v29) {
								v137 = v7 + int32(8)
							} else {
								v137 = v59
							}
							v139 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v28) {
								v142 = v14 + int32(8)
							} else {
								v142 = v82
							}
							v144 = int32(base.Ui32(v103) >> (uint(int32(1)) % 32))
							if v129 == int32(0) {
								if v112 == int32(16384) {
									v169 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v137, v139, v79, v142, v144, v102)
									mBase = m.M
									v169 = v150
								}
							} else {
								if v112 == int32(0) {
									v169 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v142, v144, v102, v137, v139, v79)
									mBase = m.M
									v169 = v154
								}
							}
						}
					}
				}
			}
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v170 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v174 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v169 <= int32(0))
						}
					} else {
						return base.B2i32(v169 <= int32(0))
					}
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v174 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v169 <= int32(0))
					}
				} else {
					return base.B2i32(v169 <= int32(0))
				}
			}
		}
	}
}
func F_numeric_mod_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
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
	var v24 int32
	_ = v24
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v78 int32
	_ = v78
	var v82 int32
	_ = v82
	var v88 int32
	_ = v88
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v142 int32
	_ = v142
	var v145 int32
	_ = v145
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v157 int32
	_ = v157
	var v167 int32
	_ = v167
	var v168 int64
	_ = v168
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v206 int32
	_ = v206
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v212 int64
	_ = v212
	var v224 int32
	_ = v224
	var v228 int32
	_ = v228
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v256 int32
	_ = v256
	v10 = m.G0
	v12 = v10 - int32(96)
	m.G0 = v12
	if l2 != 0 {
		v14 = int32(0)
		*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v14)
	} else {
	}
	v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v17 = base.I32_extend16_s(v16)
	if base.Ui32(v16) <= base.Ui32(int32(49151)) {
		v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v21 = base.I32_extend16_s(v20)
		if base.Ui32(int32(49151)) < base.Ui32(v20) {
			v48 = v21
			if v48&int32(65535) != int32(49152) {
				if v17&int32(-8193) == int32(-12288) {
					if base.Ui32(int32(49151)) < base.Ui32(v48&int32(65535)) {
						v103 = F_make_result_opt_error(m, int32(1766380), int32(0))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							v256 = v103
							m.G0 = v12 + int32(96)
							return v256
						}
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v48) {
							v78 = int32(-8)
						} else {
							v78 = int32(-6)
						}
						if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v70)>>(uint(int32(2))%32))+v78) {
							v103 = F_make_result_opt_error(m, int32(1766380), int32(0))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								v256 = v103
								m.G0 = v12 + int32(96)
								return v256
							}
						} else {
							if l2 != 0 {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v82)
								v256 = int32(0)
								m.G0 = v12 + int32(96)
								return v256
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(33816706))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(249746), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(520760), int32(3517), int32(221028))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
				} else {
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v108 = F_palloc(m, int32(base.Ui32(v105)>>(uint(int32(2))%32)))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
						if v112 != 0 {
							v113 = F__emscripten_memcpy_bulkmem(m, v108, l0, v112)
							mBase = m.M
						} else {
						}
						v256 = v108
						m.G0 = v12 + int32(96)
						return v256
					}
				}
			} else {
				v58 = F_make_result_opt_error(m, int32(1766380), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v256 = v58
					m.G0 = v12 + int32(96)
					return v256
				}
			}
		} else {
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v30 = base.B2i32(int32(0) <= v17)
			if int32(0) <= v17 {
				v31 = int32(-8)
			} else {
				v31 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = int32(base.Ui32(int32(base.Ui32(v24)>>(uint(int32(2))%32))+v31) >> (uint(int32(1)) % 32))
			if int32(0) <= v17 {
				v115 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v116 = v115
			} else {
				v116 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = v116
			v118 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v118
			v127 = base.B2i32(v17 < v118)
			if v17 < v118 {
				v128 = int32(base.Ui32(v16)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v128 = v16 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v128
			v135 = v16 & int32(49152)
			if v135 == int32(32768) {
				v138 = v16 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v138 = v135
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = v138
			if v17 < v118 {
				v142 = int32(6)
			} else {
				v142 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = l0 + v142
			v145 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v151 = base.B2i32(int32(0) <= v21)
			if int32(0) <= v21 {
				v152 = int32(-8)
			} else {
				v152 = int32(-6)
			}
			v153 = int32(base.Ui32(v145)>>(uint(int32(2))%32)) + v152
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = int32(base.Ui32(v153) >> (uint(int32(1)) % 32))
			if int32(0) <= v21 {
				v157 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				v167 = v157
			} else {
				v167 = v20<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v20&int32(63)
			}
			v168 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v168
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v168
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v167
			*(*int64)(unsafe.Add(mBase, uint32(v12))) = v168
			v175 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v175
			v184 = base.B2i32(v21 < v175)
			if v21 < v175 {
				v185 = int32(base.Ui32(v20)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v185 = v20 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v185
			v192 = v20 & int32(49152)
			if v192 == int32(32768) {
				v195 = v20 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v195 = v192
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v195
			if v21 < v175 {
				v199 = int32(6)
			} else {
				v199 = int32(8)
			}
			v200 = l1 + v199
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v200
			if l2 == int32(0) {
				v211 = v12 + int32(88)
				v212 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v211))) = v212
				*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v212
				*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v212
				v224 = int32(0)
				F_div_var(m, v12+int32(48), v12+int32(24), v12+int32(72), v224, v224, int32(1))
				mBase = m.M
				v228 = m.ExcPending
				if v228 != 0 {
					return int32(0)
				} else {
					v232 = v12 + int32(72)
					F_mul_var(m, v12+int32(24), v232, v232, v185)
					mBase = m.M
					v236 = m.ExcPending
					if v236 != 0 {
						return int32(0)
					} else {
						F_sub_var(m, v12+int32(48), v12+int32(72), v12)
						mBase = m.M
						v242 = m.ExcPending
						if v242 != 0 {
							return int32(0)
						} else {
							v243 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
							if v243 != 0 {
								F_pfree(m, v243)
								mBase = m.M
								v245 = m.ExcPending
								if v245 != 0 {
									return int32(0)
								} else {
									v247 = F_make_result_opt_error(m, v12, int32(0))
									mBase = m.M
									v248 = m.ExcPending
									if v248 != 0 {
										return int32(0)
									} else {
										v249 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
										if v249 == int32(0) {
											v256 = v247
											m.G0 = v12 + int32(96)
											return v256
										} else {
											F_pfree(m, v249)
											mBase = m.M
											v253 = m.ExcPending
											if v253 != 0 {
												return int32(0)
											} else {
												v256 = v247
												m.G0 = v12 + int32(96)
												return v256
											}
										}
									}
								}
							} else {
								v247 = F_make_result_opt_error(m, v12, int32(0))
								mBase = m.M
								v248 = m.ExcPending
								if v248 != 0 {
									return int32(0)
								} else {
									v249 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
									if v249 == int32(0) {
										v256 = v247
										m.G0 = v12 + int32(96)
										return v256
									} else {
										F_pfree(m, v249)
										mBase = m.M
										v253 = m.ExcPending
										if v253 != 0 {
											return int32(0)
										} else {
											v256 = v247
											m.G0 = v12 + int32(96)
											return v256
										}
									}
								}
							}
						}
					}
				}
			} else {
				if base.Ui32(int32(2)) <= base.Ui32(v153) {
					v206 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v200))))
					if v206 != 0 {
						v211 = v12 + int32(88)
						v212 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v211))) = v212
						*(*int64)(unsafe.Add(mBase, uint32(v12)+80)) = v212
						*(*int64)(unsafe.Add(mBase, uint32(v12)+72)) = v212
						v224 = int32(0)
						F_div_var(m, v12+int32(48), v12+int32(24), v12+int32(72), v224, v224, int32(1))
						mBase = m.M
						v228 = m.ExcPending
						if v228 != 0 {
							return int32(0)
						} else {
							v232 = v12 + int32(72)
							F_mul_var(m, v12+int32(24), v232, v232, v185)
							mBase = m.M
							v236 = m.ExcPending
							if v236 != 0 {
								return int32(0)
							} else {
								F_sub_var(m, v12+int32(48), v12+int32(72), v12)
								mBase = m.M
								v242 = m.ExcPending
								if v242 != 0 {
									return int32(0)
								} else {
									v243 = *(*int32)(unsafe.Add(mBase, uint32(v211)))
									if v243 != 0 {
										F_pfree(m, v243)
										mBase = m.M
										v245 = m.ExcPending
										if v245 != 0 {
											return int32(0)
										} else {
											v247 = F_make_result_opt_error(m, v12, int32(0))
											mBase = m.M
											v248 = m.ExcPending
											if v248 != 0 {
												return int32(0)
											} else {
												v249 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
												if v249 == int32(0) {
													v256 = v247
													m.G0 = v12 + int32(96)
													return v256
												} else {
													F_pfree(m, v249)
													mBase = m.M
													v253 = m.ExcPending
													if v253 != 0 {
														return int32(0)
													} else {
														v256 = v247
														m.G0 = v12 + int32(96)
														return v256
													}
												}
											}
										}
									} else {
										v247 = F_make_result_opt_error(m, v12, int32(0))
										mBase = m.M
										v248 = m.ExcPending
										if v248 != 0 {
											return int32(0)
										} else {
											v249 = *(*int32)(unsafe.Add(mBase, uint32(v12)+16))
											if v249 == int32(0) {
												v256 = v247
												m.G0 = v12 + int32(96)
												return v256
											} else {
												F_pfree(m, v249)
												mBase = m.M
												v253 = m.ExcPending
												if v253 != 0 {
													return int32(0)
												} else {
													v256 = v247
													m.G0 = v12 + int32(96)
													return v256
												}
											}
										}
									}
								}
							}
						}
					} else {
						v207 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v207)
						v256 = int32(0)
						m.G0 = v12 + int32(96)
						return v256
					}
				} else {
					v207 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v207)
					v256 = int32(0)
					m.G0 = v12 + int32(96)
					return v256
				}
			}
		}
	} else {
		if v17 == int32(-16384) {
			v58 = F_make_result_opt_error(m, int32(1766380), int32(0))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v256 = v58
				m.G0 = v12 + int32(96)
				return v256
			}
		} else {
			v47 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v48 = v47
			if v48&int32(65535) != int32(49152) {
				if v17&int32(-8193) == int32(-12288) {
					if base.Ui32(int32(49151)) < base.Ui32(v48&int32(65535)) {
						v103 = F_make_result_opt_error(m, int32(1766380), int32(0))
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							v256 = v103
							m.G0 = v12 + int32(96)
							return v256
						}
					} else {
						v70 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v48) {
							v78 = int32(-8)
						} else {
							v78 = int32(-6)
						}
						if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v70)>>(uint(int32(2))%32))+v78) {
							v103 = F_make_result_opt_error(m, int32(1766380), int32(0))
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								v256 = v103
								m.G0 = v12 + int32(96)
								return v256
							}
						} else {
							if l2 != 0 {
								v82 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l2))) = uint8(v82)
								v256 = int32(0)
								m.G0 = v12 + int32(96)
								return v256
							} else {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(33816706))
									mBase = m.M
									v91 = m.ExcPending
									if v91 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(249746), int32(0))
										mBase = m.M
										v95 = m.ExcPending
										if v95 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(520760), int32(3517), int32(221028))
											mBase = m.M
											v100 = m.ExcPending
											if v100 != 0 {
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
				} else {
					v105 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v108 = F_palloc(m, int32(base.Ui32(v105)>>(uint(int32(2))%32)))
					mBase = m.M
					v109 = m.ExcPending
					if v109 != 0 {
						return int32(0)
					} else {
						v110 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v112 = int32(base.Ui32(v110) >> (uint(int32(2)) % 32))
						if v112 != 0 {
							v113 = F__emscripten_memcpy_bulkmem(m, v108, l0, v112)
							mBase = m.M
						} else {
						}
						v256 = v108
						m.G0 = v12 + int32(96)
						return v256
					}
				}
			} else {
				v58 = F_make_result_opt_error(m, int32(1766380), int32(0))
				mBase = m.M
				v61 = m.ExcPending
				if v61 != 0 {
					return int32(0)
				} else {
					v256 = v58
					m.G0 = v12 + int32(96)
					return v256
				}
			}
		}
	}
}
func F_numeric_ne(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v45 int32
	_ = v45
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v69 int32
	_ = v69
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v102 int32
	_ = v102
	var v103 int32
	_ = v103
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v122 int32
	_ = v122
	var v129 int32
	_ = v129
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v150 int32
	_ = v150
	var v154 int32
	_ = v154
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v12 = l0 + int32(28)
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		v14 = F_pg_detoast_datum(m, v13)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int32(0)
		} else {
			v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v14)+4)))
			v29 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
			v30 = int32(49152)
			v31 = v29 & v30
			if v31 == v30 {
				if v29 != int32(53248) {
					if v29 != int32(49152) {
						if v28 != int32(61440) {
							v50 = int32(-1)
						} else {
							v50 = int32(0)
						}
						v169 = v50
					} else {
						v169 = base.B2i32(v28 != int32(49152))
					}
				} else {
					if v28 == int32(49152) {
						v45 = int32(-1)
					} else {
						v45 = base.B2i32(v28 != int32(53248))
					}
					v169 = v45
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v28) {
					if v28 == int32(61440) {
						v57 = int32(1)
					} else {
						v57 = int32(-1)
					}
					v169 = v57
				} else {
					v59 = v7 + int32(6)
					v64 = base.B2i32(int32(0) <= base.I32_extend16_s(v29))
					if int32(0) <= base.I32_extend16_s(v29) {
						v65 = int32(-8)
					} else {
						v65 = int32(-6)
					}
					v66 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
					if int32(0) <= base.I32_extend16_s(v29) {
						v69 = int32(*(*int16)(unsafe.Add(mBase, uint32(v59))))
						v79 = v69
					} else {
						v79 = v29<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v29&int32(63)
					}
					v80 = v65 + int32(base.Ui32(v66)>>(uint(int32(2))%32))
					v82 = v14 + int32(6)
					v87 = base.B2i32(int32(0) <= base.I32_extend16_s(v28))
					if int32(0) <= base.I32_extend16_s(v28) {
						v88 = int32(-8)
					} else {
						v88 = int32(-6)
					}
					v89 = *(*int32)(unsafe.Add(mBase, uint32(v14)))
					if int32(0) <= base.I32_extend16_s(v28) {
						v92 = int32(*(*int16)(unsafe.Add(mBase, uint32(v82))))
						v102 = v92
					} else {
						v102 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
					}
					v103 = v88 + int32(base.Ui32(v89)>>(uint(int32(2))%32))
					v109 = v28 & int32(49152)
					if v109 == int32(32768) {
						v112 = v28 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v112 = v109
					}
					if base.Ui32(v80) <= base.Ui32(int32(1)) {
						if base.Ui32(v103) < base.Ui32(int32(2)) {
							v169 = int32(0)
						} else {
							if v112 == int32(16384) {
								v122 = int32(1)
							} else {
								v122 = int32(-1)
							}
							v169 = v122
						}
					} else {
						if v31 == int32(32768) {
							v129 = v29 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v129 = v31
						}
						if base.Ui32(v103) <= base.Ui32(int32(1)) {
							if v129 != 0 {
								v134 = int32(-1)
							} else {
								v134 = int32(1)
							}
							v169 = v134
						} else {
							if int32(0) <= base.I32_extend16_s(v29) {
								v137 = v7 + int32(8)
							} else {
								v137 = v59
							}
							v139 = int32(base.Ui32(v80) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v28) {
								v142 = v14 + int32(8)
							} else {
								v142 = v82
							}
							v144 = int32(base.Ui32(v103) >> (uint(int32(1)) % 32))
							if v129 == int32(0) {
								if v112 == int32(16384) {
									v169 = int32(1)
								} else {
									v150 = F_cmp_abs_common(m, v137, v139, v79, v142, v144, v102)
									mBase = m.M
									v169 = v150
								}
							} else {
								if v112 == int32(0) {
									v169 = int32(-1)
								} else {
									v154 = F_cmp_abs_common(m, v142, v144, v102, v137, v139, v79)
									mBase = m.M
									v169 = v154
								}
							}
						}
					}
				}
			}
			v170 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			if v170 != v7 {
				F_pfree(m, v7)
				mBase = m.M
				v173 = m.ExcPending
				if v173 != 0 {
					return int32(0)
				} else {
					v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					if v174 != v14 {
						F_pfree(m, v14)
						mBase = m.M
						v177 = m.ExcPending
						if v177 != 0 {
							return int32(0)
						} else {
							return base.B2i32(v169 != int32(0))
						}
					} else {
						return base.B2i32(v169 != int32(0))
					}
				}
			} else {
				v174 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v174 != v14 {
					F_pfree(m, v14)
					mBase = m.M
					v177 = m.ExcPending
					if v177 != 0 {
						return int32(0)
					} else {
						return base.B2i32(v169 != int32(0))
					}
				} else {
					return base.B2i32(v169 != int32(0))
				}
			}
		}
	}
}
func F_numeric_sign(m *base.Module, l0 int32) int32 {
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
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v23 int32
	_ = v23
	var v30 int32
	_ = v30
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
		if v11 == int32(49152) {
			v47 = int32(1766380)
		} else {
			v14 = base.I32_extend16_s(v11)
			v15 = int32(49152)
			v16 = v11 & v15
			if v16 == v15 {
				if v14 != int32(-12288) {
					v47 = int32(1766476)
				} else {
					v47 = int32(1766504)
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
				if int32(0) <= v14 {
					v30 = int32(-8)
				} else {
					v30 = int32(-6)
				}
				if base.Ui32(int32(base.Ui32(v23)>>(uint(int32(2))%32))+v30) < base.Ui32(int32(2)) {
					v47 = int32(1766452)
				} else {
					if v16 == int32(32768) {
						v40 = v11 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v40 = v16
					}
					if v40 == int32(16384) {
						v47 = int32(1766476)
					} else {
						v47 = int32(1766504)
					}
				}
			}
		}
		v49 = F_make_result_opt_error(m, v47, int32(0))
		mBase = m.M
		v50 = m.ExcPending
		if v50 != 0 {
			return int32(0)
		} else {
			return v49
		}
	}
}
func F_numeric_to_number(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v115 int32
	_ = v115
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v124 int32
	_ = v124
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
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v18 = F_pg_detoast_datum_packed(m, v17)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
			if v20 == int32(1) {
				v23 = int32(4)
				v25 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
				if v25&int32(254) == int32(2) {
					v34 = v23
				} else {
					v34 = base.B2i32(v25 == int32(18)) << (uint(v23) % 32)
				}
				if v25 == int32(1) {
					v37 = v23
				} else {
					v37 = v34
				}
				v50 = v37
			} else {
				v38 = int32(1)
				if v20&v38 != 0 {
					v50 = int32(base.Ui32(v20)>>(uint(v38)%32)) - v38
				} else {
					v44 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
					v50 = int32(base.Ui32(v44)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			if base.Ui32(v50-int32(268435455)) <= base.Ui32(int32(-268435455)) {
				v55 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
				v157 = int32(0)
				m.G0 = v10 + int32(48)
				return v157
			} else {
				v62 = F_NUM_cache(m, v50, v10+int32(12), v18, v10+int32(11))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					v68 = F_palloc(m, v50<<(uint(int32(3))%32)|int32(1))
					mBase = m.M
					v69 = m.ExcPending
					if v69 != 0 {
						return int32(0)
					} else {
						v72 = int32(1)
						v73 = v13 + v72
						v76 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v13))))
						v78 = v76 & v72
						if v78 != 0 {
							v79 = v73
						} else {
							v79 = v13 + int32(4)
						}
						if v76 == int32(1) {
							v82 = int32(4)
							v84 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
							if v84&int32(254) == int32(2) {
								v93 = v82
							} else {
								v93 = base.B2i32(v84 == int32(18)) << (uint(v82) % 32)
							}
							if v84 == int32(1) {
								v96 = v82
							} else {
								v96 = v93
							}
							v107 = v96
						} else {
							v97 = int32(1)
							if v78 != 0 {
								v107 = int32(base.Ui32(v76)>>(uint(v97)%32)) - v97
							} else {
								v101 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
								v107 = int32(base.Ui32(v101)>>(uint(int32(2))%32)) - int32(4)
							}
						}
						v108 = int32(0)
						F_NUM_processor(m, v62, v10+int32(12), v79, v68, v107, v108, v108, v108)
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v10)+32))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v10)+16))
							v115 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v118 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+11)))
							if v118 == int32(1) {
								F_pfree(m, v62)
								mBase = m.M
								v122 = m.ExcPending
								if v122 != 0 {
									return int32(0)
								} else {
									v124 = int32(0)
									v131 = F_DirectFunctionCall3Coll(m, int32(408), v124, v68, v124, (v113+(v114+v115))<<(uint(int32(16))%32)|v114+int32(4))
									mBase = m.M
									v132 = m.ExcPending
									if v132 != 0 {
										return int32(0)
									} else {
										v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
										if v133&int32(8) != 0 {
											v137 = int32(0)
											v141 = F_int64_to_numeric(m, int64(10))
											mBase = m.M
											v142 = m.ExcPending
											if v142 != 0 {
												return int32(0)
											} else {
												v146 = F_int64_to_numeric(m, base.I64_extend_i32_s(int32(0)-v113))
												mBase = m.M
												v147 = m.ExcPending
												if v147 != 0 {
													return int32(0)
												} else {
													v148 = F_DirectFunctionCall2Coll(m, int32(1313), v137, v141, v146)
													mBase = m.M
													v149 = m.ExcPending
													if v149 != 0 {
														return int32(0)
													} else {
														v150 = F_pg_detoast_datum(m, v148)
														mBase = m.M
														v151 = m.ExcPending
														if v151 != 0 {
															return int32(0)
														} else {
															v152 = F_DirectFunctionCall2Coll(m, int32(1278), v137, v131, v150)
															mBase = m.M
															v153 = m.ExcPending
															if v153 != 0 {
																return int32(0)
															} else {
																v154 = v152
																F_pfree(m, v68)
																mBase = m.M
																v156 = m.ExcPending
																if v156 != 0 {
																	return int32(0)
																} else {
																	v157 = v154
																	m.G0 = v10 + int32(48)
																	return v157
																}
															}
														}
													}
												}
											}
										} else {
											v154 = v131
											F_pfree(m, v68)
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												v157 = v154
												m.G0 = v10 + int32(48)
												return v157
											}
										}
									}
								}
							} else {
								v124 = int32(0)
								v131 = F_DirectFunctionCall3Coll(m, int32(408), v124, v68, v124, (v113+(v114+v115))<<(uint(int32(16))%32)|v114+int32(4))
								mBase = m.M
								v132 = m.ExcPending
								if v132 != 0 {
									return int32(0)
								} else {
									v133 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v10)+25)))
									if v133&int32(8) != 0 {
										v137 = int32(0)
										v141 = F_int64_to_numeric(m, int64(10))
										mBase = m.M
										v142 = m.ExcPending
										if v142 != 0 {
											return int32(0)
										} else {
											v146 = F_int64_to_numeric(m, base.I64_extend_i32_s(int32(0)-v113))
											mBase = m.M
											v147 = m.ExcPending
											if v147 != 0 {
												return int32(0)
											} else {
												v148 = F_DirectFunctionCall2Coll(m, int32(1313), v137, v141, v146)
												mBase = m.M
												v149 = m.ExcPending
												if v149 != 0 {
													return int32(0)
												} else {
													v150 = F_pg_detoast_datum(m, v148)
													mBase = m.M
													v151 = m.ExcPending
													if v151 != 0 {
														return int32(0)
													} else {
														v152 = F_DirectFunctionCall2Coll(m, int32(1278), v137, v131, v150)
														mBase = m.M
														v153 = m.ExcPending
														if v153 != 0 {
															return int32(0)
														} else {
															v154 = v152
															F_pfree(m, v68)
															mBase = m.M
															v156 = m.ExcPending
															if v156 != 0 {
																return int32(0)
															} else {
																v157 = v154
																m.G0 = v10 + int32(48)
																return v157
															}
														}
													}
												}
											}
										}
									} else {
										v154 = v131
										F_pfree(m, v68)
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											v157 = v154
											m.G0 = v10 + int32(48)
											return v157
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
func F_numeric_trim_scale(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
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
	var v32 int32
	_ = v32
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v57 int32
	_ = v57
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v111 int32
	_ = v111
	var v118 int32
	_ = v118
	var v125 int32
	_ = v125
	var v126 int32
	_ = v126
	var v127 int32
	_ = v127
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v19 = int32(base.Ui32(v17) >> (uint(int32(2)) % 32))
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
	if base.Ui32(int32(49152)) <= base.Ui32(v20) {
		goto L4
	} else {
		goto L5
	}
L3:
	;
	m.G0 = v10 + int32(32)
	return v127
L4:
	;
	v23 = F_palloc(m, v19)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L7
	}
L5:
	;
	goto L6
L6:
	;
	v32 = base.I32_extend16_s(v20)
	v34 = base.B2i32(int32(0) <= v32)
	if int32(0) <= v32 {
		goto L12
	} else {
		goto L13
	}
L7:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
	v27 = int32(base.Ui32(v25) >> (uint(int32(2)) % 32))
	if v27 != 0 {
		goto L9
	} else {
		goto L10
	}
L8:
	;
	v127 = v23
	goto L3
L9:
	;
	v28 = F__emscripten_memcpy_bulkmem(m, v23, v13, v27)
	mBase = m.M
	goto L11
L10:
	;
	goto L11
L11:
	;
	goto L8
L12:
	;
	v35 = int32(-8)
	goto L14
L13:
	;
	v35 = int32(-6)
	goto L14
L14:
	;
	v38 = int32(base.Ui32(v19+v35) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v10)+8)) = v38
	if int32(0) <= v32 {
		goto L15
	} else {
		goto L16
	}
L15:
	;
	v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+6)))
	v50 = v40
	goto L17
L16:
	;
	v50 = v20<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v20&int32(63)
	goto L17
L17:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = v50
	v57 = v20 & int32(49152)
	if v57 == int32(32768) {
		goto L18
	} else {
		goto L19
	}
L18:
	;
	v60 = v20 << (uint(int32(1)) % 32) & int32(16384)
	goto L20
L19:
	;
	v60 = v57
	goto L20
L20:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = v60
	v62 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v10)+24)) = v62
	if v32 < v62 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v69 = int32(6)
	goto L23
L22:
	;
	v69 = int32(8)
	goto L23
L23:
	;
	v70 = v13 + v69
	*(*int32)(unsafe.Add(mBase, uint32(v10)+28)) = v70
	v72 = v38
	goto L25
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v10)+20)) = v118
	v125 = F_make_result_opt_error(m, v10+int32(8), int32(0))
	mBase = m.M
	v126 = m.ExcPending
	if v126 != 0 {
		goto L1
	} else {
		goto L36
	}
L25:
	;
	if v72 <= int32(0) {
		v118 = v62
		goto L24
	} else {
		goto L27
	}
L26:
	;
	v91 = (v82 - v50) << (uint(int32(2)) % 32)
	if v91 <= int32(0) {
		v118 = v62
		goto L24
	} else {
		goto L29
	}
L27:
	;
	v81 = int32(1)
	v82 = v72 - v81
	v86 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v70+v82<<(uint(v81)%32)))))
	if v86 == int32(0) {
		v72 = v82
		goto L25
	} else {
		goto L28
	}
L28:
	;
	goto L26
L29:
	;
	v96 = base.I32_rem_s(base.I32_extend16_s(v86), int32(10))
	if v96 != 0 {
		goto L30
	} else {
		goto L31
	}
L30:
	;
	v118 = v91
	goto L24
L31:
	;
	goto L32
L32:
	;
	v98 = v86
	v101 = v91
	goto L33
L33:
	;
	v105 = v101 - int32(1)
	v107 = int32(10)
	v108 = base.I32_div_s(base.I32_extend16_s(v98), v107)
	v111 = base.I32_rem_s(base.I32_extend16_s(v108), v107)
	if v111 == int32(0) {
		v98 = v108
		v101 = v105
		goto L33
	} else {
		goto L35
	}
L34:
	;
	v118 = v105
	goto L24
L35:
	;
	goto L34
L36:
	;
	v127 = v125
	goto L3
}
