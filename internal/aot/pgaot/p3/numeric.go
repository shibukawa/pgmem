package p3

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_do_numeric_accum(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v26 int64
	_ = v26
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v88 int64
	_ = v88
	var v92 int32
	_ = v92
	var v95 int64
	_ = v95
	var v102 int32
	_ = v102
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v112 int32
	_ = v112
	var v114 int64
	_ = v114
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v130 int32
	_ = v130
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.Ui32(int32(49152)) <= base.Ui32(v11) {
		if v11 != int32(61440) {
			if v11 != int32(53248) {
				v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+88))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+88)) = v26 + int64(1)
			} else {
				v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+96))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+96)) = v18 + int64(1)
			}
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+104))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+104)) = v22 + int64(1)
		}
		m.G0 = v9 + int32(48)
		return
	} else {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
		v35 = base.I32_extend16_s(v11)
		v37 = base.B2i32(int32(0) <= v35)
		if int32(0) <= v35 {
			v38 = int32(-8)
		} else {
			v38 = int32(-6)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = int32(base.Ui32(int32(base.Ui32(v30)>>(uint(int32(2))%32))+v38) >> (uint(int32(1)) % 32))
		if int32(0) <= v35 {
			v43 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
			v53 = v43
		} else {
			v53 = v11<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v11&int32(63)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v53
		v55 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = v55
		v64 = base.B2i32(v35 < v55)
		if v35 < v55 {
			v65 = int32(base.Ui32(v11)>>(uint(int32(7))%32)) & int32(63)
		} else {
			v65 = v11 & int32(16383)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v65
		v72 = v11 & int32(49152)
		if v72 == int32(32768) {
			v75 = v11 << (uint(int32(1)) % 32) & int32(16384)
		} else {
			v75 = v72
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+32)) = v75
		if v35 < v55 {
			v79 = int32(6)
		} else {
			v79 = int32(8)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = l1 + v79
		v82 = *(*int32)(unsafe.Add(mBase, uint32(l0)+72))
		if v82 < v65 {
			*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = int64(1)
			*(*int32)(unsafe.Add(mBase, uint32(l0)+72)) = v65
		} else {
			if v82 != v65 {
			} else {
				v88 = *(*int64)(unsafe.Add(mBase, uint32(l0)+80))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+80)) = v88 + int64(1)
			}
		}
		v92 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
		if v92 == int32(1) {
			v95 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v95
			*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v95
			*(*int64)(unsafe.Add(mBase, uint32(v9))) = v95
			v102 = v9 + int32(24)
			F_mul_var(m, v102, v102, v9, v65<<(uint(int32(1))%32))
			mBase = m.M
			v108 = m.ExcPending
			if v108 != 0 {
				return
			} else {
				v109 = int32(4562080)
				v110 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, _consts[0])) = v112
				v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v114 + int64(1)
				F_accum_sum_add(m, l0+int32(16), v9+int32(24))
				mBase = m.M
				v123 = m.ExcPending
				if v123 != 0 {
					return
				} else {
					v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v124 == int32(1) {
						F_accum_sum_add(m, l0+int32(44), v9)
						mBase = m.M
						v130 = m.ExcPending
						if v130 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _consts[0])) = v110
							m.G0 = v9 + int32(48)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v110
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		} else {
			v109 = int32(4562080)
			v110 = *(*int32)(unsafe.Add(mBase, _consts[0]))
			v112 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, _consts[0])) = v112
			v114 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v114 + int64(1)
			F_accum_sum_add(m, l0+int32(16), v9+int32(24))
			mBase = m.M
			v123 = m.ExcPending
			if v123 != 0 {
				return
			} else {
				v124 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v124 == int32(1) {
					F_accum_sum_add(m, l0+int32(44), v9)
					mBase = m.M
					v130 = m.ExcPending
					if v130 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _consts[0])) = v110
						m.G0 = v9 + int32(48)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _consts[0])) = v110
					m.G0 = v9 + int32(48)
					return
				}
			}
		}
	}
}
func F_numeric_accum_inv(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
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
	var v24 int32
	_ = v24
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	v3 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v3 != 0 {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v24 = m.ExcPending
		if v24 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(371051), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(525301), int32(5558), int32(34193))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v4 == int32(0) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return int32(0)
			} else {
				F_errmsg_internal(m, int32(371051), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(525301), int32(5558), int32(34193))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int32(0)
					} else {
						base.Wasm_trap_unreachable()
						for {
						}
					}
				}
			}
		} else {
			v7 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+32)))
			if v7 != 0 {
				v18 = v4
				return v18
			} else {
				v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
				v9 = F_pg_detoast_datum(m, v8)
				mBase = m.M
				v12 = m.ExcPending
				if v12 != 0 {
					return int32(0)
				} else {
					v13 = F_do_numeric_discard(m, v4, v9)
					mBase = m.M
					v14 = m.ExcPending
					if v14 != 0 {
						return int32(0)
					} else {
						if v13 != 0 {
							v18 = v4
						} else {
							v15 = int32(1)
							*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v15)
							v18 = int32(0)
						}
						return v18
					}
				}
			}
		}
	}
}
func F_numeric_avg_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v44 int32
	_ = v44
	var v45 int64
	_ = v45
	var v56 int32
	_ = v56
	var v57 int64
	_ = v57
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int64
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v102 int32
	_ = v102
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v117 int32
	_ = v117
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v126 int32
	_ = v126
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v145 int64
	_ = v145
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v156 int64
	_ = v156
	var v158 int64
	_ = v158
	var v161 int64
	_ = v161
	var v163 int64
	_ = v163
	var v165 int64
	_ = v165
	var v167 int64
	_ = v167
	var v190 int32
	_ = v190
	var v193 int64
	_ = v193
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v202 int64
	_ = v202
	var v204 int64
	_ = v204
	var v206 int64
	_ = v206
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v215 int64
	_ = v215
	var v238 int32
	_ = v238
	var v241 int64
	_ = v241
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 int64
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v257 int64
	_ = v257
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int64
	_ = v263
	var v286 int32
	_ = v286
	var v289 int64
	_ = v289
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v296 int32
	_ = v296
	var v298 int64
	_ = v298
	var v300 int64
	_ = v300
	var v302 int64
	_ = v302
	var v305 int64
	_ = v305
	var v307 int64
	_ = v307
	var v309 int64
	_ = v309
	var v311 int64
	_ = v311
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v355 int32
	_ = v355
	var v359 int32
	_ = v359
	var v364 int32
	_ = v364
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v13 == int32(0) {
		v41 = int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
		switch v16 - int32(429) {
		case 0:
			v41 = int32(1)
		case 1:
			v41 = int32(2)
		default:
			v41 = int32(0)
		}
	}
	if v41 != 0 {
		v42 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v44 = v9 + int32(24)
		v45 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v44))) = v45
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v45
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v45
		F_pq_begintypsend(m, v9+int32(32))
		mBase = m.M
		v56 = m.ExcPending
		if v56 != 0 {
			return int32(0)
		} else {
			v57 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
			F_enlargeStringInfo(m, v9+int32(32), int32(8))
			mBase = m.M
			v62 = m.ExcPending
			if v62 != 0 {
				return int32(0)
			} else {
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
				v64 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
				v66 = int64(56)
				v68 = int64(65280)
				v70 = int64(40)
				v73 = int64(16711680)
				v75 = int64(24)
				v77 = int64(4278190080)
				v79 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v63+v64))) = v57<<(uint(v66)%64) | v57&v68<<(uint(v70)%64) | (v57&v73<<(uint(v75)%64) | v57&v77<<(uint(v79)%64)) | (int64(base.Ui64(v57)>>(uint(v79)%64))&v77 | int64(base.Ui64(v57)>>(uint(v75)%64))&v73 | (int64(base.Ui64(v57)>>(uint(v70)%64))&v68 | int64(base.Ui64(v57)>>(uint(v66)%64))))
				v102 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v63 + v102
				F_accum_sum_final(m, v42+int32(16), v9+v102)
				mBase = m.M
				v110 = m.ExcPending
				if v110 != 0 {
					return int32(0)
				} else {
					F_numericvar_serialize(m, v9+int32(32), v9+int32(8))
					mBase = m.M
					v116 = m.ExcPending
					if v116 != 0 {
						return int32(0)
					} else {
						v117 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
						F_enlargeStringInfo(m, v9+int32(32), int32(4))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int32(0)
						} else {
							v123 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
							v124 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
							v126 = int32(24)
							v128 = int32(65280)
							v130 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v123+v124))) = v117<<(uint(v126)%32) | v117&v128<<(uint(v130)%32) | (int32(base.Ui32(v117)>>(uint(v130)%32))&v128 | int32(base.Ui32(v117)>>(uint(v126)%32)))
							*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v123 + int32(4)
							v145 = *(*int64)(unsafe.Add(mBase, uint32(v42)+80))
							F_enlargeStringInfo(m, v9+int32(32), v130)
							mBase = m.M
							v150 = m.ExcPending
							if v150 != 0 {
								return int32(0)
							} else {
								v151 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
								v152 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
								v154 = int64(56)
								v156 = int64(65280)
								v158 = int64(40)
								v161 = int64(16711680)
								v163 = int64(24)
								v165 = int64(4278190080)
								v167 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v151+v152))) = v145<<(uint(v154)%64) | v145&v156<<(uint(v158)%64) | (v145&v161<<(uint(v163)%64) | v145&v165<<(uint(v167)%64)) | (int64(base.Ui64(v145)>>(uint(v167)%64))&v165 | int64(base.Ui64(v145)>>(uint(v163)%64))&v161 | (int64(base.Ui64(v145)>>(uint(v158)%64))&v156 | int64(base.Ui64(v145)>>(uint(v154)%64))))
								v190 = int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v151 + v190
								v193 = *(*int64)(unsafe.Add(mBase, uint32(v42)+88))
								F_enlargeStringInfo(m, v9+int32(32), v190)
								mBase = m.M
								v198 = m.ExcPending
								if v198 != 0 {
									return int32(0)
								} else {
									v199 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
									v200 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
									v202 = int64(56)
									v204 = int64(65280)
									v206 = int64(40)
									v209 = int64(16711680)
									v211 = int64(24)
									v213 = int64(4278190080)
									v215 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v199+v200))) = v193<<(uint(v202)%64) | v193&v204<<(uint(v206)%64) | (v193&v209<<(uint(v211)%64) | v193&v213<<(uint(v215)%64)) | (int64(base.Ui64(v193)>>(uint(v215)%64))&v213 | int64(base.Ui64(v193)>>(uint(v211)%64))&v209 | (int64(base.Ui64(v193)>>(uint(v206)%64))&v204 | int64(base.Ui64(v193)>>(uint(v202)%64))))
									v238 = int32(8)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v199 + v238
									v241 = *(*int64)(unsafe.Add(mBase, uint32(v42)+96))
									F_enlargeStringInfo(m, v9+int32(32), v238)
									mBase = m.M
									v246 = m.ExcPending
									if v246 != 0 {
										return int32(0)
									} else {
										v247 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
										v248 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
										v250 = int64(56)
										v252 = int64(65280)
										v254 = int64(40)
										v257 = int64(16711680)
										v259 = int64(24)
										v261 = int64(4278190080)
										v263 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v247+v248))) = v241<<(uint(v250)%64) | v241&v252<<(uint(v254)%64) | (v241&v257<<(uint(v259)%64) | v241&v261<<(uint(v263)%64)) | (int64(base.Ui64(v241)>>(uint(v263)%64))&v261 | int64(base.Ui64(v241)>>(uint(v259)%64))&v257 | (int64(base.Ui64(v241)>>(uint(v254)%64))&v252 | int64(base.Ui64(v241)>>(uint(v250)%64))))
										v286 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v247 + v286
										v289 = *(*int64)(unsafe.Add(mBase, uint32(v42)+104))
										F_enlargeStringInfo(m, v9+int32(32), v286)
										mBase = m.M
										v294 = m.ExcPending
										if v294 != 0 {
											return int32(0)
										} else {
											v295 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
											v296 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v298 = int64(56)
											v300 = int64(65280)
											v302 = int64(40)
											v305 = int64(16711680)
											v307 = int64(24)
											v309 = int64(4278190080)
											v311 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v295+v296))) = v289<<(uint(v298)%64) | v289&v300<<(uint(v302)%64) | (v289&v305<<(uint(v307)%64) | v289&v309<<(uint(v311)%64)) | (int64(base.Ui64(v289)>>(uint(v311)%64))&v309 | int64(base.Ui64(v289)>>(uint(v307)%64))&v305 | (int64(base.Ui64(v289)>>(uint(v302)%64))&v300 | int64(base.Ui64(v289)>>(uint(v298)%64))))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v295 + int32(8)
											v338 = v9 + int32(32)
											v340 = *(*int32)(unsafe.Add(mBase, uint32(v338)))
											v341 = *(*int32)(unsafe.Add(mBase, uint32(v338)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v340))) = v341 << (uint(int32(2)) % 32)
											v345 = *(*int32)(unsafe.Add(mBase, uint32(v44)))
											if v345 != 0 {
												F_pfree(m, v345)
												mBase = m.M
												v347 = m.ExcPending
												if v347 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(48)
													return v340
												}
											} else {
												m.G0 = v9 + int32(48)
												return v340
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
		v355 = m.ExcPending
		if v355 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66939), int32(0))
			mBase = m.M
			v359 = m.ExcPending
			if v359 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(525301), int32(5332), int32(359719))
				mBase = m.M
				v364 = m.ExcPending
				if v364 != 0 {
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
func F_numeric_deserialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int64
	_ = v46
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v84 int32
	_ = v84
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v116 int32
	_ = v116
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v137 int64
	_ = v137
	var v138 int32
	_ = v138
	var v142 int64
	_ = v142
	var v143 int32
	_ = v143
	var v147 int64
	_ = v147
	var v148 int32
	_ = v148
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v161 int32
	_ = v161
	var v169 int32
	_ = v169
	var v173 int32
	_ = v173
	var v178 int32
	_ = v178
	v6 = m.G0
	v8 = v6 - int32(48)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		v40 = int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(429) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 != 0 {
		v41 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v42 = F_pg_detoast_datum_packed(m, v41)
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int32(0)
		} else {
			v46 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(v8)+16)) = v46
			*(*int64)(unsafe.Add(mBase, uint32(v8)+8)) = v46
			v52 = int32(1)
			v53 = v42 + v52
			v54 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v56 = v54 & v52
			if v54 == v52 {
				v59 = int32(4)
				v61 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v53))))
				if v61&int32(254) == int32(2) {
					v70 = v59
				} else {
					v70 = base.B2i32(v61 == int32(18)) << (uint(v59) % 32)
				}
				if v61 == int32(1) {
					v73 = v59
				} else {
					v73 = v70
				}
				v84 = v73
			} else {
				v74 = int32(1)
				if v56 != 0 {
					v84 = int32(base.Ui32(v54)>>(uint(v74)%32)) - v74
				} else {
					v78 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v84 = int32(base.Ui32(v78)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v84
			if v56 != 0 {
				v90 = v53
			} else {
				v90 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v90
			v93 = F_palloc0(m, int32(112))
			mBase = m.M
			v94 = m.ExcPending
			if v94 != 0 {
				return int32(0)
			} else {
				v95 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v93))) = uint8(v95)
				v98 = *(*int32)(unsafe.Add(mBase, _consts[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v93)+4)) = v98
				v102 = F_pq_getmsgint64(m, v8+int32(32))
				mBase = m.M
				v103 = m.ExcPending
				if v103 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v93)+8)) = v102
					F_numericvar_deserialize(m, v8+int32(32), v8+int32(8))
					mBase = m.M
					v110 = m.ExcPending
					if v110 != 0 {
						return int32(0)
					} else {
						F_accum_sum_add(m, v93+int32(16), v8+int32(8))
						mBase = m.M
						v116 = m.ExcPending
						if v116 != 0 {
							return int32(0)
						} else {
							F_numericvar_deserialize(m, v8+int32(32), v8+int32(8))
							mBase = m.M
							v122 = m.ExcPending
							if v122 != 0 {
								return int32(0)
							} else {
								F_accum_sum_add(m, v93+int32(44), v8+int32(8))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int32(0)
								} else {
									v132 = F_pq_getmsgint(m, v8+int32(32), int32(4))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v93)+72)) = v132
										v137 = F_pq_getmsgint64(m, v8+int32(32))
										mBase = m.M
										v138 = m.ExcPending
										if v138 != 0 {
											return int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v93)+80)) = v137
											v142 = F_pq_getmsgint64(m, v8+int32(32))
											mBase = m.M
											v143 = m.ExcPending
											if v143 != 0 {
												return int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v93)+88)) = v142
												v147 = F_pq_getmsgint64(m, v8+int32(32))
												mBase = m.M
												v148 = m.ExcPending
												if v148 != 0 {
													return int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v93)+96)) = v147
													v152 = F_pq_getmsgint64(m, v8+int32(32))
													mBase = m.M
													v153 = m.ExcPending
													if v153 != 0 {
														return int32(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v93)+104)) = v152
														F_pq_getmsgend(m, v8+int32(32))
														mBase = m.M
														v158 = m.ExcPending
														if v158 != 0 {
															return int32(0)
														} else {
															v159 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
															if v159 != 0 {
																F_pfree(m, v159)
																mBase = m.M
																v161 = m.ExcPending
																if v161 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(48)
																	return v93
																}
															} else {
																m.G0 = v8 + int32(48)
																return v93
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
	} else {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v169 = m.ExcPending
		if v169 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(66939), int32(0))
			mBase = m.M
			v173 = m.ExcPending
			if v173 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(525301), int32(5497), int32(359568))
				mBase = m.M
				v178 = m.ExcPending
				if v178 != 0 {
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
func F_numeric_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v13 int64
	_ = v13
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v27 int32
	_ = v27
	var v38 int32
	_ = v38
	var v43 int32
	_ = v43
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v89 int32
	_ = v89
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v160 int32
	_ = v160
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v182 int32
	_ = v182
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v200 int32
	_ = v200
	var v210 int32
	_ = v210
	var v219 int32
	_ = v219
	var v222 int32
	_ = v222
	var v224 int32
	_ = v224
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v254 int32
	_ = v254
	var v265 int32
	_ = v265
	var v274 int32
	_ = v274
	var v278 int32
	_ = v278
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v286 int32
	_ = v286
	var v292 int32
	_ = v292
	var v303 int32
	_ = v303
	var v308 int32
	_ = v308
	var v309 int32
	_ = v309
	var v315 int64
	_ = v315
	var v322 int32
	_ = v322
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v337 int64
	_ = v337
	var v345 int32
	_ = v345
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v357 int32
	_ = v357
	var v359 int32
	_ = v359
	var v367 int64
	_ = v367
	var v368 int64
	_ = v368
	var v376 int32
	_ = v376
	var v378 int32
	_ = v378
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v383 int32
	_ = v383
	var v388 int32
	_ = v388
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v404 int64
	_ = v404
	var v407 int32
	_ = v407
	var v409 int64
	_ = v409
	var v412 int64
	_ = v412
	var v415 int32
	_ = v415
	var v422 int32
	_ = v422
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v449 int64
	_ = v449
	var v453 int32
	_ = v453
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v468 int64
	_ = v468
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v484 int64
	_ = v484
	var v488 int32
	_ = v488
	var v490 int64
	_ = v490
	var v493 int64
	_ = v493
	var v496 int32
	_ = v496
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v518 int32
	_ = v518
	var v524 int32
	_ = v524
	var v525 int32
	_ = v525
	var v528 int32
	_ = v528
	var v533 int32
	_ = v533
	var v543 int64
	_ = v543
	var v544 int64
	_ = v544
	var v546 int64
	_ = v546
	var v556 int32
	_ = v556
	var v557 int32
	_ = v557
	var v560 int32
	_ = v560
	var v562 int32
	_ = v562
	var v570 int64
	_ = v570
	var v571 int64
	_ = v571
	var v576 int32
	_ = v576
	var v577 int64
	_ = v577
	var v585 int32
	_ = v585
	var v590 int32
	_ = v590
	var v591 int32
	_ = v591
	var v597 int32
	_ = v597
	var v599 int32
	_ = v599
	var v607 int64
	_ = v607
	var v608 int64
	_ = v608
	var v611 int32
	_ = v611
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v629 int32
	_ = v629
	var v630 int32
	_ = v630
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v642 int32
	_ = v642
	var v643 int32
	_ = v643
	var v653 int64
	_ = v653
	var v656 int32
	_ = v656
	var v658 int64
	_ = v658
	var v661 int64
	_ = v661
	var v664 int32
	_ = v664
	var v671 int32
	_ = v671
	var v678 int32
	_ = v678
	var v679 int32
	_ = v679
	var v681 int32
	_ = v681
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v686 int32
	_ = v686
	var v688 int32
	_ = v688
	var v698 int64
	_ = v698
	var v702 int32
	_ = v702
	var v707 int32
	_ = v707
	var v708 int32
	_ = v708
	var v722 int64
	_ = v722
	var v728 int32
	_ = v728
	var v729 int32
	_ = v729
	var v738 int64
	_ = v738
	var v742 int32
	_ = v742
	var v744 int64
	_ = v744
	var v747 int64
	_ = v747
	var v750 int32
	_ = v750
	var v756 int32
	_ = v756
	var v759 int32
	_ = v759
	var v772 int32
	_ = v772
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v787 int32
	_ = v787
	var v797 int64
	_ = v797
	var v798 int64
	_ = v798
	var v800 int64
	_ = v800
	var v802 int32
	_ = v802
	var v805 int32
	_ = v805
	var v830 int32
	_ = v830
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v836 int32
	_ = v836
	var v839 int32
	_ = v839
	var v841 int32
	_ = v841
	var v849 int64
	_ = v849
	var v850 int64
	_ = v850
	var v854 int64
	_ = v854
	var v862 int32
	_ = v862
	var v867 int32
	_ = v867
	var v868 int32
	_ = v868
	var v874 int32
	_ = v874
	var v876 int32
	_ = v876
	var v884 int64
	_ = v884
	var v885 int64
	_ = v885
	var v893 int32
	_ = v893
	var v895 int32
	_ = v895
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v900 int32
	_ = v900
	var v905 int32
	_ = v905
	var v910 int32
	_ = v910
	var v911 int32
	_ = v911
	var v921 int64
	_ = v921
	var v924 int32
	_ = v924
	var v926 int64
	_ = v926
	var v929 int64
	_ = v929
	var v932 int32
	_ = v932
	var v939 int32
	_ = v939
	var v946 int32
	_ = v946
	var v947 int32
	_ = v947
	var v949 int32
	_ = v949
	var v951 int32
	_ = v951
	var v952 int32
	_ = v952
	var v954 int32
	_ = v954
	var v956 int32
	_ = v956
	var v966 int64
	_ = v966
	var v970 int32
	_ = v970
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v985 int64
	_ = v985
	var v991 int32
	_ = v991
	var v992 int32
	_ = v992
	var v1001 int64
	_ = v1001
	var v1005 int32
	_ = v1005
	var v1007 int64
	_ = v1007
	var v1010 int64
	_ = v1010
	var v1013 int32
	_ = v1013
	var v1019 int32
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1035 int32
	_ = v1035
	var v1041 int32
	_ = v1041
	var v1042 int32
	_ = v1042
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1060 int64
	_ = v1060
	var v1061 int64
	_ = v1061
	var v1063 int64
	_ = v1063
	var v1073 int32
	_ = v1073
	var v1074 int32
	_ = v1074
	var v1077 int32
	_ = v1077
	var v1079 int32
	_ = v1079
	var v1087 int64
	_ = v1087
	var v1088 int64
	_ = v1088
	var v1096 int32
	_ = v1096
	var v1100 int32
	_ = v1100
	var v1104 int64
	_ = v1104
	var v1105 int64
	_ = v1105
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1112 int32
	_ = v1112
	var v1113 int32
	_ = v1113
	var v1115 int32
	_ = v1115
	var v1118 int32
	_ = v1118
	var v1128 int64
	_ = v1128
	var v1134 int64
	_ = v1134
	var v1139 int32
	_ = v1139
	var v1140 int32
	_ = v1140
	var v1150 int64
	_ = v1150
	var v1153 int32
	_ = v1153
	var v1155 int64
	_ = v1155
	var v1158 int64
	_ = v1158
	var v1161 int32
	_ = v1161
	var v1167 int32
	_ = v1167
	var v1170 int32
	_ = v1170
	var v1183 int32
	_ = v1183
	var v1190 int32
	_ = v1190
	var v1191 int32
	_ = v1191
	var v1193 int32
	_ = v1193
	var v1195 int32
	_ = v1195
	var v1196 int32
	_ = v1196
	var v1198 int32
	_ = v1198
	var v1200 int32
	_ = v1200
	var v1210 int64
	_ = v1210
	var v1214 int32
	_ = v1214
	var v1231 int32
	_ = v1231
	var v1232 int32
	_ = v1232
	var v1237 int32
	_ = v1237
	var v1245 int32
	_ = v1245
	var v1247 int64
	_ = v1247
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1263 int64
	_ = v1263
	var v1267 int32
	_ = v1267
	var v1269 int64
	_ = v1269
	var v1272 int64
	_ = v1272
	var v1275 int32
	_ = v1275
	var v1281 int32
	_ = v1281
	var v1284 int32
	_ = v1284
	var v1297 int32
	_ = v1297
	var v1303 int32
	_ = v1303
	var v1304 int32
	_ = v1304
	var v1308 int32
	_ = v1308
	var v1310 int32
	_ = v1310
	var v1327 int32
	_ = v1327
	var v1328 int32
	_ = v1328
	var v1333 int32
	_ = v1333
	var v1337 int32
	_ = v1337
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1373 int32
	_ = v1373
	var v1379 int32
	_ = v1379
	var v1394 int32
	_ = v1394
	var v1405 int32
	_ = v1405
	var v1414 int32
	_ = v1414
	var v1415 int32
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1425 int32
	_ = v1425
	var v1426 int32
	_ = v1426
	var v1427 int32
	_ = v1427
	var v1430 int32
	_ = v1430
	var v1431 int32
	_ = v1431
	var v1436 int32
	_ = v1436
	var v1440 int32
	_ = v1440
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1451 int32
	_ = v1451
	var v1469 int32
	_ = v1469
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1482 int32
	_ = v1482
	var v1487 int32
	_ = v1487
	var v1489 int32
	_ = v1489
	var v1505 int32
	_ = v1505
	var v1507 int32
	_ = v1507
	v2 = int32(0)
	v13 = int64(0)
	v16 = m.G0
	v18 = v16 - int32(112)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v27 = v22
	goto L3
L1:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v56
	v60 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v60 == int32(46) {
		goto L11
	} else {
		goto L12
	}
L2:
	;
	v56 = v27 + int32(1)
	v57 = int32(1774620)
	v58 = int32(16384)
	goto L1
L3:
	;
	v38 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v27))))
	if base.Ui32(v38-int32(9)) < base.Ui32(int32(5)) {
		goto L6
	} else {
		goto L7
	}
L4:
	;
	v56 = v27 + int32(1)
	v57 = v43
	v58 = v2
	goto L1
L5:
	;
	goto L4
L6:
	;
	v27 = v27 + int32(1)
	goto L3
L7:
	;
	v43 = int32(1774596)
	switch v38 - int32(32) {
	case 0:
		goto L6
	default:
		v56 = v27
		v57 = v43
		v58 = v2
		goto L1
	case 11:
		goto L5
	case 13:
		goto L2
	}
L8:
	;
	m.G0 = v1505 + int32(112)
	return v1507
L9:
	;
	v1505 = v1489
	v1507 = int32(0)
	goto L8
L10:
	;
	v1469 = int32(0)
	v1470 = F_errsave_start(m, v21)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L34
	} else {
		goto L289
	}
L11:
	;
	v315 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18-int32(-64)))) = v315
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v315
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v315
	v322 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56))))
	if v322 != int32(48) {
		goto L102
	} else {
		goto L103
	}
L12:
	;
	if base.Ui32((v60-int32(48))&int32(255)) < base.Ui32(int32(10)) {
		goto L11
	} else {
		goto L13
	}
L13:
	;
	v73 = v27
	v74 = int32(553864)
	v75 = int32(3)
	goto L16
L14:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v247
	v254 = v247
	goto L74
L15:
	;
	if v120 == int32(0) {
		goto L31
	} else {
		goto L32
	}
L16:
	;
	if v75 != 0 {
		goto L18
	} else {
		goto L19
	}
L17:
	;
	v120 = int32(0)
	goto L15
L18:
	;
	v78 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v73))))
	v79 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v74))))
	if v78 == v79 {
		v101 = v78
		goto L21
	} else {
		goto L22
	}
L19:
	;
	goto L20
L20:
	;
	goto L17
L21:
	;
	v103 = int32(1)
	if v101 != 0 {
		v73 = v73 + v103
		v74 = v74 + v103
		v75 = v75 - v103
		goto L16
	} else {
		goto L30
	}
L22:
	;
	if base.Ui32((v78-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L23
	} else {
		goto L24
	}
L23:
	;
	v89 = v78 | int32(32)
	goto L25
L24:
	;
	v89 = v78
	goto L25
L25:
	;
	if base.Ui32((v79-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v98 = v79 | int32(32)
	goto L28
L27:
	;
	v98 = v79
	goto L28
L28:
	;
	if v89 == v98 {
		v101 = v89
		goto L21
	} else {
		goto L29
	}
L29:
	;
	v120 = v89 - v98
	goto L15
L30:
	;
	goto L20
L31:
	;
	v127 = F_make_result_opt_error(m, int32(1774572), int32(0))
	mBase = m.M
	v130 = m.ExcPending
	if v130 != 0 {
		goto L34
	} else {
		goto L35
	}
L32:
	;
	goto L33
L33:
	;
	v135 = v56
	v136 = int32(11540)
	v137 = int32(8)
	goto L37
L34:
	;
	return int32(0)
L35:
	;
	v247 = v27 + int32(3)
	v248 = v127
	goto L14
L36:
	;
	if v182 == int32(0) {
		goto L52
	} else {
		goto L53
	}
L37:
	;
	if v137 != 0 {
		goto L39
	} else {
		goto L40
	}
L38:
	;
	v182 = int32(0)
	goto L36
L39:
	;
	v140 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v135))))
	v141 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v136))))
	if v140 == v141 {
		v163 = v140
		goto L42
	} else {
		goto L43
	}
L40:
	;
	goto L41
L41:
	;
	goto L38
L42:
	;
	v165 = int32(1)
	if v163 != 0 {
		v135 = v135 + v165
		v136 = v136 + v165
		v137 = v137 - v165
		goto L37
	} else {
		goto L51
	}
L43:
	;
	if base.Ui32((v140-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v151 = v140 | int32(32)
	goto L46
L45:
	;
	v151 = v140
	goto L46
L46:
	;
	if base.Ui32((v141-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L47
	} else {
		goto L48
	}
L47:
	;
	v160 = v141 | int32(32)
	goto L49
L48:
	;
	v160 = v141
	goto L49
L49:
	;
	if v151 == v160 {
		v163 = v151
		goto L42
	} else {
		goto L50
	}
L50:
	;
	v182 = v151 - v160
	goto L36
L51:
	;
	goto L41
L52:
	;
	v188 = F_make_result_opt_error(m, v57, int32(0))
	mBase = m.M
	v189 = m.ExcPending
	if v189 != 0 {
		goto L34
	} else {
		goto L55
	}
L53:
	;
	goto L54
L54:
	;
	v194 = v56
	v195 = int32(355874)
	v196 = int32(3)
	goto L57
L55:
	;
	v247 = v56 + int32(8)
	v248 = v188
	goto L14
L56:
	;
	if v241 != 0 {
		goto L10
	} else {
		goto L72
	}
L57:
	;
	if v196 != 0 {
		goto L59
	} else {
		goto L60
	}
L58:
	;
	v241 = int32(0)
	goto L56
L59:
	;
	v199 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v194))))
	v200 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v195))))
	if v199 == v200 {
		v222 = v199
		goto L62
	} else {
		goto L63
	}
L60:
	;
	goto L61
L61:
	;
	goto L58
L62:
	;
	v224 = int32(1)
	if v222 != 0 {
		v194 = v194 + v224
		v195 = v195 + v224
		v196 = v196 - v224
		goto L57
	} else {
		goto L71
	}
L63:
	;
	if base.Ui32((v199-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v210 = v199 | int32(32)
	goto L66
L65:
	;
	v210 = v199
	goto L66
L66:
	;
	if base.Ui32((v200-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L67
	} else {
		goto L68
	}
L67:
	;
	v219 = v200 | int32(32)
	goto L69
L68:
	;
	v219 = v200
	goto L69
L69:
	;
	if v210 == v219 {
		v222 = v210
		goto L62
	} else {
		goto L70
	}
L70:
	;
	v241 = v210 - v219
	goto L56
L71:
	;
	goto L61
L72:
	;
	v245 = F_make_result_opt_error(m, v57, int32(0))
	mBase = m.M
	v246 = m.ExcPending
	if v246 != 0 {
		goto L34
	} else {
		goto L73
	}
L73:
	;
	v247 = v56 + int32(3)
	v248 = v245
	goto L14
L74:
	;
	v265 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v254))))
	if base.Ui32(v265-int32(9)) < base.Ui32(int32(5)) {
		goto L76
	} else {
		goto L77
	}
L76:
	;
	v254 = v254 + int32(1)
	goto L74
L77:
	;
	if v265 == int32(32) {
		goto L76
	} else {
		goto L78
	}
L78:
	;
	if v265 != 0 {
		goto L10
	} else {
		goto L79
	}
L79:
	;
	if v20 < int32(4) {
		v1505 = v18
		v1507 = v248
		goto L8
	} else {
		goto L80
	}
L80:
	;
	v274 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v248)+4)))
	if v274 == int32(49152) {
		v1505 = v18
		v1507 = v248
		goto L8
	} else {
		goto L81
	}
L81:
	;
	v278 = F_errsave_start(m, v21)
	mBase = m.M
	v279 = m.ExcPending
	if v279 != 0 {
		goto L34
	} else {
		goto L82
	}
L82:
	;
	if v278 != 0 {
		goto L83
	} else {
		goto L84
	}
L83:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v282 = m.ExcPending
	if v282 != 0 {
		goto L34
	} else {
		goto L86
	}
L84:
	;
	goto L85
L85:
	;
	v309 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v309)
	v1505 = v18
	v1507 = int32(0)
	goto L8
L86:
	;
	F_errmsg(m, int32(32942), int32(0))
	mBase = m.M
	v286 = m.ExcPending
	if v286 != 0 {
		goto L34
	} else {
		goto L87
	}
L87:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = int32(base.Ui32(v20-int32(4)) >> (uint(int32(16)) % 32))
	v292 = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = (v20<<(uint(v292)%32) - int32(8388608)) >> (uint(v292) % 32)
	F_errdetail(m, int32(657902), v18+int32(32))
	mBase = m.M
	v303 = m.ExcPending
	if v303 != 0 {
		goto L34
	} else {
		goto L88
	}
L88:
	;
	F_errsave_finish(m, v21, int32(525301), int32(8138), int32(329398))
	mBase = m.M
	v308 = m.ExcPending
	if v308 != 0 {
		goto L34
	} else {
		goto L89
	}
L89:
	;
	goto L85
L90:
	;
	v1394 = v1379
	goto L268
L91:
	;
	v1373 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1373)
	v1489 = v1359
	goto L9
L92:
	;
	F_errsave_finish(m, v21, int32(525301), v1354, int32(217227))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L34
	} else {
		goto L267
	}
L93:
	;
	v1327 = F_errsave_start(m, v21)
	mBase = m.M
	v1328 = m.ExcPending
	if v1328 != 0 {
		goto L34
	} else {
		goto L263
	}
L94:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v1284
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v1281
	v1297 = v18 + int32(56)
	F_add_var(m, v1297, v18+int32(88), v1297)
	mBase = m.M
	v1303 = m.ExcPending
	if v1303 != 0 {
		goto L34
	} else {
		goto L257
	}
L95:
	;
	v1253 = int32(0)
	v1254 = v1200 + int32(12)
	v1263 = v1247
	goto L254
L96:
	;
	v1231 = F_errsave_start(m, v21)
	mBase = m.M
	v1232 = m.ExcPending
	if v1232 != 0 {
		goto L34
	} else {
		goto L250
	}
L97:
	;
	if v1096 == v1100 {
		goto L96
	} else {
		goto L225
	}
L98:
	;
	v854 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v854
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v854
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v854
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v854
	v862 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v862
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v854
	v867 = v56 + int32(2)
	v868 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v867))))
	if v868 == v862 {
		goto L96
	} else {
		goto L188
	}
L99:
	;
	v577 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v577
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v577
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v577
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v577
	v585 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v585
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v577
	v590 = v56 + int32(2)
	v591 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v590))))
	if v591 == v585 {
		goto L96
	} else {
		goto L143
	}
L100:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v58
	v576 = *(*int32)(unsafe.Add(mBase, uint32(v18)+84))
	v1379 = v576
	goto L90
L101:
	;
	v337 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v337
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v337
	v345 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v345
	*(*int64)(unsafe.Add(mBase, uint32(v18)+56)) = v337
	v350 = v56 + int32(2)
	v351 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v350))))
	if v351 == v345 {
		goto L96
	} else {
		goto L106
	}
L102:
	;
	v332 = F_set_var_from_str(m, v22, v56, v18+int32(56), v18+int32(84), v21)
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L34
	} else {
		goto L104
	}
L103:
	;
	v325 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v56)+1)))
	switch v325 - int32(66) {
	case 0, 32:
		goto L101
	default:
		goto L102
	case 13, 45:
		goto L98
	case 22, 54:
		goto L99
	}
L104:
	;
	if v332 != 0 {
		goto L100
	} else {
		goto L105
	}
L105:
	;
	v334 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v334)
	v1505 = v18
	v1507 = int32(0)
	goto L8
L106:
	;
	v357 = v351
	v359 = v350
	v367 = v13
	v368 = int64(1)
	goto L107
L107:
	;
	if v357&int32(254) == int32(48) {
		goto L113
	} else {
		goto L114
	}
L108:
	;
	v1096 = v562
	v1100 = v350
	v1104 = v570
	v1105 = v571
	goto L97
L109:
	;
	if v560&int32(255) != 0 {
		v357 = v560
		v359 = v562
		v367 = v570
		v368 = v571
		goto L107
	} else {
		goto L142
	}
L110:
	;
	v546 = int64(1)
	v556 = v359 + int32(1)
	v557 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v556))))
	v560 = v557
	v562 = v556
	v570 = base.I64_extend8_s(base.I64_extend_i32_u(v533)) + v543<<(uint(v546)%64) - int64(48)
	v571 = v544 << (uint(v546) % 64)
	goto L109
L111:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v505
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v502
	v518 = v18 + int32(56)
	F_add_var(m, v518, v18+int32(88), v518)
	mBase = m.M
	v524 = m.ExcPending
	if v524 != 0 {
		goto L34
	} else {
		goto L140
	}
L112:
	;
	v474 = int32(0)
	v475 = v439 + int32(12)
	v484 = v468
	goto L137
L113:
	;
	if v368 < int64(4611686018427387904) {
		v533 = v357
		v543 = v367
		v544 = v368
		goto L110
	} else {
		goto L116
	}
L114:
	;
	goto L115
L115:
	;
	if v357&int32(255) != int32(95) {
		v1096 = v359
		v1100 = v350
		v1104 = v367
		v1105 = v368
		goto L97
	} else {
		goto L135
	}
L116:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v376 != 0 {
		goto L117
	} else {
		goto L118
	}
L117:
	;
	F_pfree(m, v376)
	mBase = m.M
	v378 = m.ExcPending
	if v378 != 0 {
		goto L34
	} else {
		goto L120
	}
L118:
	;
	goto L119
L119:
	;
	v380 = F_palloc(m, int32(12))
	mBase = m.M
	v381 = m.ExcPending
	if v381 != 0 {
		goto L34
	} else {
		goto L121
	}
L120:
	;
	goto L119
L121:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v380
	v383 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v380))) = uint16(v383)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(0)
	v388 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	v393 = v383
	v394 = v388 + int32(12)
	v404 = v368
	goto L122
L122:
	;
	v407 = v394 - int32(2)
	v409 = base.I64_div_u_s(v404, int64(10000))
	v412 = v409*int64(55536) + v404
	*(*uint16)(unsafe.Add(mBase, uint32(v407))) = uint16(v412)
	v415 = v393 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v404) {
		v393 = v415
		v394 = v407
		v404 = v409
		goto L122
	} else {
		goto L124
	}
L123:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v393
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v415
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v407
	v422 = v18 + int32(56)
	F_mul_var(m, v422, v18+int32(88), v422, int32(0))
	mBase = m.M
	v429 = m.ExcPending
	if v429 != 0 {
		goto L34
	} else {
		goto L125
	}
L124:
	;
	goto L123
L125:
	;
	v430 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v430 != 0 {
		goto L126
	} else {
		goto L127
	}
L126:
	;
	F_pfree(m, v430)
	mBase = m.M
	v432 = m.ExcPending
	if v432 != 0 {
		goto L34
	} else {
		goto L129
	}
L127:
	;
	goto L128
L128:
	;
	v434 = F_palloc(m, int32(12))
	mBase = m.M
	v435 = m.ExcPending
	if v435 != 0 {
		goto L34
	} else {
		goto L130
	}
L129:
	;
	goto L128
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v434
	v437 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v434))) = uint16(v437)
	v439 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v439 + int32(2)
	if v367 < int64(0) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(16384)
	v468 = int64(0) - v367
	goto L112
L132:
	;
	goto L133
L133:
	;
	v449 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v449
	if v367 != v449 {
		v468 = v367
		goto L112
	} else {
		goto L134
	}
L134:
	;
	v453 = int32(0)
	v502 = v453
	v505 = v453
	goto L111
L135:
	;
	v460 = v359 + int32(1)
	v461 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v460))))
	if base.Ui32((v461-int32(50))&int32(255)) < base.Ui32(int32(254)) {
		goto L96
	} else {
		goto L136
	}
L136:
	;
	v560 = v461
	v562 = v460
	v570 = v367
	v571 = v368
	goto L109
L137:
	;
	v488 = v475 - int32(2)
	v490 = base.I64_div_u_s(v484, int64(10000))
	v493 = v490*int64(55536) + v484
	*(*uint16)(unsafe.Add(mBase, uint32(v488))) = uint16(v493)
	v496 = v474 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v484) {
		v474 = v496
		v475 = v488
		v484 = v490
		goto L137
	} else {
		goto L139
	}
L138:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v488
	v502 = v496
	v505 = v474
	goto L111
L139:
	;
	goto L138
L140:
	;
	v525 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if int32(32767) < v525 {
		goto L93
	} else {
		goto L141
	}
L141:
	;
	v528 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v359))))
	v533 = v528
	v543 = int64(0)
	v544 = int64(1)
	goto L110
L142:
	;
	goto L108
L143:
	;
	v597 = v591
	v599 = v590
	v607 = v13
	v608 = int64(1)
	goto L144
L144:
	;
	v611 = v597 & int32(255)
	goto L150
L145:
	;
	v1096 = v841
	v1100 = v590
	v1104 = v849
	v1105 = v850
	goto L97
L146:
	;
	if v839&int32(255) != 0 {
		v597 = v839
		v599 = v841
		v607 = v849
		v608 = v850
		goto L144
	} else {
		goto L187
	}
L147:
	;
	v800 = int64(4)
	v802 = base.I32_extend8_s(v787)
	v805 = int32(48)
	if base.Ui32((v787-v805)&int32(255)) <= base.Ui32(int32(9)) {
		v831 = v802 - v805
		goto L181
	} else {
		goto L182
	}
L148:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v759
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v756
	v772 = v18 + int32(56)
	F_add_var(m, v772, v18+int32(88), v772)
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L34
	} else {
		goto L179
	}
L149:
	;
	v728 = int32(0)
	v729 = v688 + int32(12)
	v738 = v722
	goto L176
L150:
	;
	if base.B2i32(base.Ui32(v611-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v611|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		goto L151
	} else {
		goto L152
	}
L151:
	;
	if v608 < int64(576460752303423488) {
		v787 = v597
		v797 = v607
		v798 = v608
		goto L147
	} else {
		goto L154
	}
L152:
	;
	goto L153
L153:
	;
	if v611 != int32(95) {
		v1096 = v599
		v1100 = v590
		v1104 = v607
		v1105 = v608
		goto L97
	} else {
		goto L173
	}
L154:
	;
	v625 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v625 != 0 {
		goto L155
	} else {
		goto L156
	}
L155:
	;
	F_pfree(m, v625)
	mBase = m.M
	v627 = m.ExcPending
	if v627 != 0 {
		goto L34
	} else {
		goto L158
	}
L156:
	;
	goto L157
L157:
	;
	v629 = F_palloc(m, int32(12))
	mBase = m.M
	v630 = m.ExcPending
	if v630 != 0 {
		goto L34
	} else {
		goto L159
	}
L158:
	;
	goto L157
L159:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v629
	v632 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v629))) = uint16(v632)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(0)
	v637 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	v642 = v632
	v643 = v637 + int32(12)
	v653 = v608
	goto L160
L160:
	;
	v656 = v643 - int32(2)
	v658 = base.I64_div_u_s(v653, int64(10000))
	v661 = v658*int64(55536) + v653
	*(*uint16)(unsafe.Add(mBase, uint32(v656))) = uint16(v661)
	v664 = v642 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v653) {
		v642 = v664
		v643 = v656
		v653 = v658
		goto L160
	} else {
		goto L162
	}
L161:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v642
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v664
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v656
	v671 = v18 + int32(56)
	F_mul_var(m, v671, v18+int32(88), v671, int32(0))
	mBase = m.M
	v678 = m.ExcPending
	if v678 != 0 {
		goto L34
	} else {
		goto L163
	}
L162:
	;
	goto L161
L163:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v679 != 0 {
		goto L164
	} else {
		goto L165
	}
L164:
	;
	F_pfree(m, v679)
	mBase = m.M
	v681 = m.ExcPending
	if v681 != 0 {
		goto L34
	} else {
		goto L167
	}
L165:
	;
	goto L166
L166:
	;
	v683 = F_palloc(m, int32(12))
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L34
	} else {
		goto L168
	}
L167:
	;
	goto L166
L168:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v683
	v686 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v683))) = uint16(v686)
	v688 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v688 + int32(2)
	if v607 < int64(0) {
		goto L169
	} else {
		goto L170
	}
L169:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(16384)
	v722 = int64(0) - v607
	goto L149
L170:
	;
	goto L171
L171:
	;
	v698 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v698
	if v607 != v698 {
		v722 = v607
		goto L149
	} else {
		goto L172
	}
L172:
	;
	v702 = int32(0)
	v756 = v702
	v759 = v702
	goto L148
L173:
	;
	v707 = v599 + int32(1)
	v708 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v707))))
	goto L174
L174:
	;
	if base.B2i32(base.Ui32(v708-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v708|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L96
	} else {
		goto L175
	}
L175:
	;
	v839 = v708
	v841 = v707
	v849 = v607
	v850 = v608
	goto L146
L176:
	;
	v742 = v729 - int32(2)
	v744 = base.I64_div_u_s(v738, int64(10000))
	v747 = v744*int64(55536) + v738
	*(*uint16)(unsafe.Add(mBase, uint32(v742))) = uint16(v747)
	v750 = v728 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v738) {
		v728 = v750
		v729 = v742
		v738 = v744
		goto L176
	} else {
		goto L178
	}
L177:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v742
	v756 = v750
	v759 = v728
	goto L148
L178:
	;
	goto L177
L179:
	;
	v779 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if int32(32767) < v779 {
		goto L93
	} else {
		goto L180
	}
L180:
	;
	v782 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v599))))
	v787 = v782
	v797 = int64(0)
	v798 = int64(1)
	goto L147
L181:
	;
	v835 = v599 + int32(1)
	v836 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v835))))
	v839 = v836
	v841 = v835
	v849 = v797<<(uint(v800)%64) + base.I64_extend_i32_s(v831)
	v850 = v798 << (uint(v800) % 64)
	goto L146
L182:
	;
	if base.Ui32((v787-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		v831 = v802 - int32(87)
		goto L181
	} else {
		goto L183
	}
L183:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v787-int32(65))&int32(255)) {
		goto L184
	} else {
		goto L185
	}
L184:
	;
	v830 = int32(-1)
	goto L186
L185:
	;
	v830 = v802 - int32(55)
	goto L186
L186:
	;
	v831 = v830
	goto L181
L187:
	;
	goto L145
L188:
	;
	v874 = v868
	v876 = v867
	v884 = v13
	v885 = int64(1)
	goto L189
L189:
	;
	if v874&int32(248) == int32(48) {
		goto L195
	} else {
		goto L196
	}
L190:
	;
	v1096 = v1079
	v1100 = v867
	v1104 = v1087
	v1105 = v1088
	goto L97
L191:
	;
	if v1077&int32(255) != 0 {
		v874 = v1077
		v876 = v1079
		v884 = v1087
		v885 = v1088
		goto L189
	} else {
		goto L224
	}
L192:
	;
	v1063 = int64(3)
	v1073 = v876 + int32(1)
	v1074 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1073))))
	v1077 = v1074
	v1079 = v1073
	v1087 = base.I64_extend8_s(base.I64_extend_i32_u(v1050)) + v1060<<(uint(v1063)%64) - int64(48)
	v1088 = v1061 << (uint(v1063) % 64)
	goto L191
L193:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v1022
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v1019
	v1035 = v18 + int32(56)
	F_add_var(m, v1035, v18+int32(88), v1035)
	mBase = m.M
	v1041 = m.ExcPending
	if v1041 != 0 {
		goto L34
	} else {
		goto L222
	}
L194:
	;
	v991 = int32(0)
	v992 = v956 + int32(12)
	v1001 = v985
	goto L219
L195:
	;
	if v885 < int64(1152921504606846976) {
		v1050 = v874
		v1060 = v884
		v1061 = v885
		goto L192
	} else {
		goto L198
	}
L196:
	;
	goto L197
L197:
	;
	if v874&int32(255) != int32(95) {
		v1096 = v876
		v1100 = v867
		v1104 = v884
		v1105 = v885
		goto L97
	} else {
		goto L217
	}
L198:
	;
	v893 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v893 != 0 {
		goto L199
	} else {
		goto L200
	}
L199:
	;
	F_pfree(m, v893)
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L34
	} else {
		goto L202
	}
L200:
	;
	goto L201
L201:
	;
	v897 = F_palloc(m, int32(12))
	mBase = m.M
	v898 = m.ExcPending
	if v898 != 0 {
		goto L34
	} else {
		goto L203
	}
L202:
	;
	goto L201
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v897
	v900 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v897))) = uint16(v900)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(0)
	v905 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	v910 = v900
	v911 = v905 + int32(12)
	v921 = v885
	goto L204
L204:
	;
	v924 = v911 - int32(2)
	v926 = base.I64_div_u_s(v921, int64(10000))
	v929 = v926*int64(55536) + v921
	*(*uint16)(unsafe.Add(mBase, uint32(v924))) = uint16(v929)
	v932 = v910 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v921) {
		v910 = v932
		v911 = v924
		v921 = v926
		goto L204
	} else {
		goto L206
	}
L205:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v910
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v932
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v924
	v939 = v18 + int32(56)
	F_mul_var(m, v939, v18+int32(88), v939, int32(0))
	mBase = m.M
	v946 = m.ExcPending
	if v946 != 0 {
		goto L34
	} else {
		goto L207
	}
L206:
	;
	goto L205
L207:
	;
	v947 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v947 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	F_pfree(m, v947)
	mBase = m.M
	v949 = m.ExcPending
	if v949 != 0 {
		goto L34
	} else {
		goto L211
	}
L209:
	;
	goto L210
L210:
	;
	v951 = F_palloc(m, int32(12))
	mBase = m.M
	v952 = m.ExcPending
	if v952 != 0 {
		goto L34
	} else {
		goto L212
	}
L211:
	;
	goto L210
L212:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v951
	v954 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v951))) = uint16(v954)
	v956 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v956 + int32(2)
	if v884 < int64(0) {
		goto L213
	} else {
		goto L214
	}
L213:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(16384)
	v985 = int64(0) - v884
	goto L194
L214:
	;
	goto L215
L215:
	;
	v966 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v966
	if v884 != v966 {
		v985 = v884
		goto L194
	} else {
		goto L216
	}
L216:
	;
	v970 = int32(0)
	v1019 = v970
	v1022 = v970
	goto L193
L217:
	;
	v977 = v876 + int32(1)
	v978 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v977))))
	if base.Ui32((v978-int32(56))&int32(255)) < base.Ui32(int32(248)) {
		goto L96
	} else {
		goto L218
	}
L218:
	;
	v1077 = v978
	v1079 = v977
	v1087 = v884
	v1088 = v885
	goto L191
L219:
	;
	v1005 = v992 - int32(2)
	v1007 = base.I64_div_u_s(v1001, int64(10000))
	v1010 = v1007*int64(55536) + v1001
	*(*uint16)(unsafe.Add(mBase, uint32(v1005))) = uint16(v1010)
	v1013 = v991 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v1001) {
		v991 = v1013
		v992 = v1005
		v1001 = v1007
		goto L219
	} else {
		goto L221
	}
L220:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v1005
	v1019 = v1013
	v1022 = v991
	goto L193
L221:
	;
	goto L220
L222:
	;
	v1042 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if int32(32767) < v1042 {
		goto L93
	} else {
		goto L223
	}
L223:
	;
	v1045 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v876))))
	v1050 = v1045
	v1060 = int64(0)
	v1061 = int64(1)
	goto L192
L224:
	;
	goto L190
L225:
	;
	v1108 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v1108 != 0 {
		goto L226
	} else {
		goto L227
	}
L226:
	;
	F_pfree(m, v1108)
	mBase = m.M
	v1110 = m.ExcPending
	if v1110 != 0 {
		goto L34
	} else {
		goto L229
	}
L227:
	;
	goto L228
L228:
	;
	v1112 = F_palloc(m, int32(12))
	mBase = m.M
	v1113 = m.ExcPending
	if v1113 != 0 {
		goto L34
	} else {
		goto L230
	}
L229:
	;
	goto L228
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v1112
	v1115 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1112))) = uint16(v1115)
	v1118 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v1118 + int32(2)
	if v1105 < int64(0) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+92)) = v1170
	*(*int32)(unsafe.Add(mBase, uint32(v18)+88)) = v1167
	v1183 = v18 + int32(56)
	F_mul_var(m, v1183, v18+int32(88), v1183, int32(0))
	mBase = m.M
	v1190 = m.ExcPending
	if v1190 != 0 {
		goto L34
	} else {
		goto L240
	}
L232:
	;
	v1139 = v1115
	v1140 = v1118 + int32(12)
	v1150 = v1134
	goto L237
L233:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(16384)
	v1134 = int64(0) - v1105
	goto L232
L234:
	;
	goto L235
L235:
	;
	v1128 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v1128
	if v1105 == v1128 {
		v1167 = v1115
		v1170 = int32(0)
		goto L231
	} else {
		goto L236
	}
L236:
	;
	v1134 = v1105
	goto L232
L237:
	;
	v1153 = v1140 - int32(2)
	v1155 = base.I64_div_u_s(v1150, int64(10000))
	v1158 = v1155*int64(55536) + v1150
	*(*uint16)(unsafe.Add(mBase, uint32(v1153))) = uint16(v1158)
	v1161 = v1139 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v1150) {
		v1139 = v1161
		v1140 = v1153
		v1150 = v1155
		goto L237
	} else {
		goto L239
	}
L238:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v1153
	v1167 = v1161
	v1170 = v1139
	goto L231
L239:
	;
	goto L238
L240:
	;
	v1191 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v1191 != 0 {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	F_pfree(m, v1191)
	mBase = m.M
	v1193 = m.ExcPending
	if v1193 != 0 {
		goto L34
	} else {
		goto L244
	}
L242:
	;
	goto L243
L243:
	;
	v1195 = F_palloc(m, int32(12))
	mBase = m.M
	v1196 = m.ExcPending
	if v1196 != 0 {
		goto L34
	} else {
		goto L245
	}
L244:
	;
	goto L243
L245:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+104)) = v1195
	v1198 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1195))) = uint16(v1198)
	v1200 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v1200 + int32(2)
	if v1104 < int64(0) {
		goto L246
	} else {
		goto L247
	}
L246:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = int64(16384)
	v1247 = int64(0) - v1104
	goto L95
L247:
	;
	goto L248
L248:
	;
	v1210 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v1210
	if v1104 != v1210 {
		v1247 = v1104
		goto L95
	} else {
		goto L249
	}
L249:
	;
	v1214 = int32(0)
	v1281 = v1214
	v1284 = v1214
	goto L94
L250:
	;
	if v1231 == int32(0) {
		v1359 = v18
		goto L91
	} else {
		goto L251
	}
L251:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1237 = m.ExcPending
	if v1237 != 0 {
		goto L34
	} else {
		goto L252
	}
L252:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = int32(513970)
	F_errmsg(m, int32(759399), v18+int32(16))
	mBase = m.M
	v1245 = m.ExcPending
	if v1245 != 0 {
		goto L34
	} else {
		goto L253
	}
L253:
	;
	v1354 = int32(7529)
	goto L92
L254:
	;
	v1267 = v1254 - int32(2)
	v1269 = base.I64_div_u_s(v1263, int64(10000))
	v1272 = v1269*int64(55536) + v1263
	*(*uint16)(unsafe.Add(mBase, uint32(v1267))) = uint16(v1272)
	v1275 = v1253 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v1263) {
		v1253 = v1275
		v1254 = v1267
		v1263 = v1269
		goto L254
	} else {
		goto L256
	}
L255:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+108)) = v1267
	v1281 = v1275
	v1284 = v1253
	goto L94
L256:
	;
	goto L255
L257:
	;
	v1304 = *(*int32)(unsafe.Add(mBase, uint32(v18)+60))
	if int32(32767) < v1304 {
		goto L93
	} else {
		goto L258
	}
L258:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v58
	v1308 = *(*int32)(unsafe.Add(mBase, uint32(v18)+104))
	if v1308 != 0 {
		goto L259
	} else {
		goto L260
	}
L259:
	;
	F_pfree(m, v1308)
	mBase = m.M
	v1310 = m.ExcPending
	if v1310 != 0 {
		goto L34
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+84)) = v1096
	v1379 = v1096
	goto L90
L262:
	;
	goto L261
L263:
	;
	if v1327 == int32(0) {
		v1359 = v18
		goto L91
	} else {
		goto L264
	}
L264:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L34
	} else {
		goto L265
	}
L265:
	;
	F_errmsg(m, int32(119783), int32(0))
	mBase = m.M
	v1337 = m.ExcPending
	if v1337 != 0 {
		goto L34
	} else {
		goto L266
	}
L266:
	;
	v1354 = int32(7523)
	goto L92
L267:
	;
	v1359 = v18
	goto L91
L268:
	;
	v1405 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1394))))
	if base.Ui32(v1405-int32(9)) < base.Ui32(int32(5)) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v1394 = v1394 + int32(1)
	goto L268
L271:
	;
	if v1405 == int32(32) {
		goto L270
	} else {
		goto L272
	}
L272:
	;
	if v1405 != 0 {
		goto L10
	} else {
		goto L273
	}
L273:
	;
	v1414 = F_apply_typmod(m, v18+int32(56), v20, v21)
	mBase = m.M
	v1415 = m.ExcPending
	if v1415 != 0 {
		goto L34
	} else {
		goto L274
	}
L274:
	;
	if v1414 == int32(0) {
		goto L275
	} else {
		goto L276
	}
L275:
	;
	v1418 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1418)
	v1505 = v18
	v1507 = int32(0)
	goto L8
L276:
	;
	goto L277
L277:
	;
	v1425 = F_make_result_opt_error(m, v18+int32(56), v18+int32(55))
	mBase = m.M
	v1426 = m.ExcPending
	if v1426 != 0 {
		goto L34
	} else {
		goto L278
	}
L278:
	;
	v1427 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+55)))
	if v1427 == int32(1) {
		goto L279
	} else {
		goto L280
	}
L279:
	;
	v1430 = F_errsave_start(m, v21)
	mBase = m.M
	v1431 = m.ExcPending
	if v1431 != 0 {
		goto L34
	} else {
		goto L282
	}
L280:
	;
	goto L281
L281:
	;
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v18)+72))
	if v1447 == int32(0) {
		v1505 = v18
		v1507 = v1425
		goto L8
	} else {
		goto L287
	}
L282:
	;
	if v1430 == int32(0) {
		v1489 = v18
		goto L9
	} else {
		goto L283
	}
L283:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1436 = m.ExcPending
	if v1436 != 0 {
		goto L34
	} else {
		goto L284
	}
L284:
	;
	F_errmsg(m, int32(119783), int32(0))
	mBase = m.M
	v1440 = m.ExcPending
	if v1440 != 0 {
		goto L34
	} else {
		goto L285
	}
L285:
	;
	F_errsave_finish(m, v21, int32(525301), int32(795), int32(293518))
	mBase = m.M
	v1445 = m.ExcPending
	if v1445 != 0 {
		goto L34
	} else {
		goto L286
	}
L286:
	;
	v1505 = v18
	v1507 = int32(0)
	goto L8
L287:
	;
	F_pfree(m, v1447)
	mBase = m.M
	v1451 = m.ExcPending
	if v1451 != 0 {
		goto L34
	} else {
		goto L288
	}
L288:
	;
	v1505 = v18
	v1507 = v1425
	goto L8
L289:
	;
	if v1470 == int32(0) {
		v1505 = v18
		v1507 = v1469
		goto L8
	} else {
		goto L290
	}
L290:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L34
	} else {
		goto L291
	}
L291:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v22
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = int32(513970)
	F_errmsg(m, int32(759399), v18)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L34
	} else {
		goto L292
	}
L292:
	;
	F_errsave_finish(m, v21, int32(525301), int32(806), int32(293518))
	mBase = m.M
	v1487 = m.ExcPending
	if v1487 != 0 {
		goto L34
	} else {
		goto L293
	}
L293:
	;
	v1505 = v18
	v1507 = v1469
	goto L8
}
func F_numeric_int2(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v85 int32
	_ = v85
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v100 int64
	_ = v100
	var v114 int32
	_ = v114
	var v119 int32
	_ = v119
	var v123 int32
	_ = v123
	var v126 int32
	_ = v126
	var v130 int32
	_ = v130
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	v7 = m.G0
	v9 = v7 + int32(-64)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
		v17 = base.I32_extend16_s(v16)
		if base.Ui32(int32(49152)) <= base.Ui32(v16) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(1088))
				mBase = m.M
				v26 = m.ExcPending
				if v26 != 0 {
					return int32(0)
				} else {
					if v17 == int32(-16384) {
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(95552)
						F_errmsg(m, int32(193899), v9)
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525301), int32(4684), int32(587789))
							mBase = m.M
							v119 = m.ExcPending
							if v119 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(95552)
						F_errmsg(m, int32(193222), v7+int32(-48))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(525301), int32(4688), int32(587789))
							mBase = m.M
							v40 = m.ExcPending
							if v40 != 0 {
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
			v41 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v47 = base.B2i32(int32(0) <= v17)
			if int32(0) <= v17 {
				v48 = int32(-8)
			} else {
				v48 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+40)) = int32(base.Ui32(int32(base.Ui32(v41)>>(uint(int32(2))%32))+v48) >> (uint(int32(1)) % 32))
			if int32(0) <= v17 {
				v53 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v63 = v53
			} else {
				v63 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+44)) = v63
			v65 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+56)) = v65
			v74 = base.B2i32(v17 < v65)
			if v17 < v65 {
				v75 = int32(base.Ui32(v16)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v75 = v16 & int32(16383)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v75
			v82 = v16 & int32(49152)
			if v82 == int32(32768) {
				v85 = v16 << (uint(int32(1)) % 32) & int32(16384)
			} else {
				v85 = v82
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+48)) = v85
			if v17 < v65 {
				v89 = int32(6)
			} else {
				v89 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+60)) = v12 + v89
			v96 = F_numericvar_to_int64(m, v7+int32(-24), v7+int32(-32))
			mBase = m.M
			v97 = m.ExcPending
			if v97 != 0 {
				return int32(0)
			} else {
				if v96 == int32(0) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v123 = m.ExcPending
					if v123 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v126 = m.ExcPending
						if v126 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(421267), int32(0))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(525301), int32(4697), int32(587789))
								mBase = m.M
								v135 = m.ExcPending
								if v135 != 0 {
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
					v100 = *(*int64)(unsafe.Add(mBase, uint32(v9)+32))
					if base.Ui64(v100-int64(32768)) <= base.Ui64(int64(-65537)) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v142 = m.ExcPending
							if v142 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(421267), int32(0))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(525301), int32(4702), int32(587789))
									mBase = m.M
									v151 = m.ExcPending
									if v151 != 0 {
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
						m.G0 = v9 - int32(-64)
						return base.I32_wrap_i64(v100)
					}
				}
			}
		}
	}
}
func F_numeric_larger(m *base.Module, l0 int32) int32 {
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
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v117 int32
	_ = v117
	var v124 int32
	_ = v124
	var v129 int32
	_ = v129
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v137 int32
	_ = v137
	var v139 int32
	_ = v139
	var v145 int32
	_ = v145
	var v149 int32
	_ = v149
	var v164 int32
	_ = v164
	var v167 int32
	_ = v167
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = F_pg_detoast_datum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v9 = F_pg_detoast_datum(m, v8)
		mBase = m.M
		v10 = m.ExcPending
		if v10 != 0 {
			return int32(0)
		} else {
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
			v24 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
			v25 = int32(49152)
			v26 = v24 & v25
			if v26 == v25 {
				if v24 != int32(53248) {
					if v24 != int32(49152) {
						if v23 != int32(61440) {
							v45 = int32(-1)
						} else {
							v45 = int32(0)
						}
						v164 = v45
					} else {
						v164 = base.B2i32(v23 != int32(49152))
					}
				} else {
					if v23 == int32(49152) {
						v40 = int32(-1)
					} else {
						v40 = base.B2i32(v23 != int32(53248))
					}
					v164 = v40
				}
			} else {
				if base.Ui32(int32(49152)) <= base.Ui32(v23) {
					if v23 == int32(61440) {
						v52 = int32(1)
					} else {
						v52 = int32(-1)
					}
					v164 = v52
				} else {
					v54 = v4 + int32(6)
					v59 = base.B2i32(int32(0) <= base.I32_extend16_s(v24))
					if int32(0) <= base.I32_extend16_s(v24) {
						v60 = int32(-8)
					} else {
						v60 = int32(-6)
					}
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
					if int32(0) <= base.I32_extend16_s(v24) {
						v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v54))))
						v74 = v64
					} else {
						v74 = v24<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v24&int32(63)
					}
					v75 = v60 + int32(base.Ui32(v61)>>(uint(int32(2))%32))
					v77 = v9 + int32(6)
					v82 = base.B2i32(int32(0) <= base.I32_extend16_s(v23))
					if int32(0) <= base.I32_extend16_s(v23) {
						v83 = int32(-8)
					} else {
						v83 = int32(-6)
					}
					v84 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					if int32(0) <= base.I32_extend16_s(v23) {
						v87 = int32(*(*int16)(unsafe.Add(mBase, uint32(v77))))
						v97 = v87
					} else {
						v97 = v23<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v23&int32(63)
					}
					v98 = v83 + int32(base.Ui32(v84)>>(uint(int32(2))%32))
					v104 = v23 & int32(49152)
					if v104 == int32(32768) {
						v107 = v23 << (uint(int32(1)) % 32) & int32(16384)
					} else {
						v107 = v104
					}
					if base.Ui32(v75) <= base.Ui32(int32(1)) {
						if base.Ui32(v98) < base.Ui32(int32(2)) {
							v164 = int32(0)
						} else {
							if v107 == int32(16384) {
								v117 = int32(1)
							} else {
								v117 = int32(-1)
							}
							v164 = v117
						}
					} else {
						if v26 == int32(32768) {
							v124 = v24 << (uint(int32(1)) % 32) & int32(16384)
						} else {
							v124 = v26
						}
						if base.Ui32(v98) <= base.Ui32(int32(1)) {
							if v124 != 0 {
								v129 = int32(-1)
							} else {
								v129 = int32(1)
							}
							v164 = v129
						} else {
							if int32(0) <= base.I32_extend16_s(v24) {
								v132 = v4 + int32(8)
							} else {
								v132 = v54
							}
							v134 = int32(base.Ui32(v75) >> (uint(int32(1)) % 32))
							if int32(0) <= base.I32_extend16_s(v23) {
								v137 = v9 + int32(8)
							} else {
								v137 = v77
							}
							v139 = int32(base.Ui32(v98) >> (uint(int32(1)) % 32))
							if v124 == int32(0) {
								if v107 == int32(16384) {
									v164 = int32(1)
								} else {
									v145 = F_cmp_abs_common(m, v132, v134, v74, v137, v139, v97)
									mBase = m.M
									v164 = v145
								}
							} else {
								if v107 == int32(0) {
									v164 = int32(-1)
								} else {
									v149 = F_cmp_abs_common(m, v137, v139, v97, v132, v134, v74)
									mBase = m.M
									v164 = v149
								}
							}
						}
					}
				}
			}
			if int32(0) < v164 {
				v167 = v4
			} else {
				v167 = v9
			}
			return v167
		}
	}
}
func F_numeric_min_scale(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v47 int32
	_ = v47
	var v49 int32
	_ = v49
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v8 = F_pg_detoast_datum(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8)+4)))
	if base.Ui32(int32(49152)) <= base.Ui32(v12) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v15 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v15)
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v19 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
	v24 = base.I32_extend16_s(v12)
	v26 = base.B2i32(int32(0) <= v24)
	if int32(0) <= v24 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v27 = int32(-8)
	goto L8
L7:
	;
	v27 = int32(-6)
	goto L8
L8:
	;
	if int32(0) <= v24 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v29 = int32(*(*int16)(unsafe.Add(mBase, uint32(v8)+6)))
	v39 = v29
	goto L11
L10:
	;
	v39 = v12<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v12&int32(63)
	goto L11
L11:
	;
	v42 = int32(0)
	if v24 < v42 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v47 = int32(6)
	goto L14
L13:
	;
	v47 = int32(8)
	goto L14
L14:
	;
	v49 = int32(base.Ui32(int32(base.Ui32(v19)>>(uint(int32(2))%32))+v27) >> (uint(int32(1)) % 32))
	goto L16
L15:
	;
	return v92
L16:
	;
	if v49 <= int32(0) {
		v92 = v42
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v67 = (v58 - v39) << (uint(int32(2)) % 32)
	if v67 <= int32(0) {
		v92 = v42
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v57 = int32(1)
	v58 = v49 - v57
	v62 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v8+v47+v58<<(uint(v57)%32)))))
	if v62 == int32(0) {
		v49 = v58
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v72 = base.I32_rem_s(base.I32_extend16_s(v62), int32(10))
	if v72 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return v67
L22:
	;
	goto L23
L23:
	;
	v75 = v62
	v76 = v67
	goto L24
L24:
	;
	v81 = v76 - int32(1)
	v83 = int32(10)
	v84 = base.I32_div_s(base.I32_extend16_s(v75), v83)
	v87 = base.I32_rem_s(base.I32_extend16_s(v84), v83)
	if v87 == int32(0) {
		v75 = v84
		v76 = v81
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v92 = v81
	goto L15
L26:
	;
	goto L25
}
func F_numeric_poly_avg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v23 int64
	_ = v23
	var v27 int32
	_ = v27
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v50 int64
	_ = v50
	var v58 int64
	_ = v58
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v93 int64
	_ = v93
	var v94 int64
	_ = v94
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v103 int64
	_ = v103
	var v108 int64
	_ = v108
	var v111 int64
	_ = v111
	var v114 int64
	_ = v114
	var v117 int64
	_ = v117
	var v118 int64
	_ = v118
	var v122 int64
	_ = v122
	var v129 int64
	_ = v129
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v147 int64
	_ = v147
	var v152 int32
	_ = v152
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v171 int64
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v190 int64
	_ = v190
	var v196 int64
	_ = v196
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v209 int64
	_ = v209
	var v214 int32
	_ = v214
	var v216 int64
	_ = v216
	var v219 int64
	_ = v219
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v258 int32
	_ = v258
	var v259 int32
	_ = v259
	var v274 int32
	_ = v274
	v15 = m.G0
	v17 = v15 - int32(80)
	m.G0 = v17
	v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v19 != 0 {
		v27 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
		v274 = int32(0)
		m.G0 = v17 + int32(80)
		return v274
	} else {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v20 == int32(0) {
			v27 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
			v274 = int32(0)
			m.G0 = v17 + int32(80)
			return v274
		} else {
			v23 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
			if v23 != int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = int64(0)
				v32 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
				v33 = *(*int64)(unsafe.Add(mBase, uint32(v20)+16))
				v35 = F_palloc(m, int32(22))
				mBase = m.M
				v38 = m.ExcPending
				if v38 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v17)+48)) = v35
					v40 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v40)
					*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v35 + int32(2)
					if v32 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = int64(16384)
						v50 = int64(0)
						v63 = v50 - (v32 + base.I64_extend_i32_u(base.B2i32(v33 != v50)))
						v64 = v50 - v33
						v69 = v40
						v71 = v35 + int32(22)
						v79 = v63
						v80 = v64
						for {
							v83 = int32(16)
							v84 = v17 + v83
							v87 = m.G0
							v89 = v87 - v83
							m.G0 = v89
							F___udivmodti4(m, v89, v80, v79, int64(10000), int64(0))
							mBase = m.M
							v93 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
							v94 = *(*int64)(unsafe.Add(mBase, uint32(v89)+8))
							*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v94
							*(*int64)(unsafe.Add(mBase, uint32(v84))) = v93
							m.G0 = v89 + v83
							v100 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
							v101 = *(*int64)(unsafe.Add(mBase, uint32(v17+int32(24))))
							v102 = int64(55536)
							v103 = int64(0)
							v108 = int64(32)
							v111 = int64(base.Ui64(v100) >> (uint(v108) % 64))
							v114 = int64(4294967295)
							v117 = v100 & v114
							v118 = v102 * v117
							v122 = int64(base.Ui64(v118)>>(uint(v108)%64)) + v102*v111
							v129 = v117*v103 + v122&v114
							*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v100*v103 + v101*v102 + v103*v111 + int64(base.Ui64(v122)>>(uint(v108)%64)) + int64(base.Ui64(v129)>>(uint(v108)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v17))) = v118&v114 | v129<<(uint(v108)%64)
							v141 = v71 - int32(2)
							v142 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
							v143 = v142 + v80
							*(*uint16)(unsafe.Add(mBase, uint32(v141))) = uint16(v143)
							v147 = int64(0)
							v152 = v69 + int32(1)
							if v79 == v147 {
								v153 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v80))
							} else {
								v153 = base.B2i32(v79 != v147)
							}
							if v153 != 0 {
								v69 = v152
								v71 = v141
								v79 = v101
								v80 = v100
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v141
						v155 = v152
						v163 = v69
					} else {
						v58 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v17)+40)) = v58
						if v32|v33 == v58 {
							v155 = v40
							v163 = int32(0)
						} else {
							v63 = v32
							v64 = v33
							v69 = v40
							v71 = v35 + int32(22)
							v79 = v63
							v80 = v64
							for {
								v83 = int32(16)
								v84 = v17 + v83
								v87 = m.G0
								v89 = v87 - v83
								m.G0 = v89
								F___udivmodti4(m, v89, v80, v79, int64(10000), int64(0))
								mBase = m.M
								v93 = *(*int64)(unsafe.Add(mBase, uint32(v89)))
								v94 = *(*int64)(unsafe.Add(mBase, uint32(v89)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v84)+8)) = v94
								*(*int64)(unsafe.Add(mBase, uint32(v84))) = v93
								m.G0 = v89 + v83
								v100 = *(*int64)(unsafe.Add(mBase, uint32(v17)+16))
								v101 = *(*int64)(unsafe.Add(mBase, uint32(v17+int32(24))))
								v102 = int64(55536)
								v103 = int64(0)
								v108 = int64(32)
								v111 = int64(base.Ui64(v100) >> (uint(v108) % 64))
								v114 = int64(4294967295)
								v117 = v100 & v114
								v118 = v102 * v117
								v122 = int64(base.Ui64(v118)>>(uint(v108)%64)) + v102*v111
								v129 = v117*v103 + v122&v114
								*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = v100*v103 + v101*v102 + v103*v111 + int64(base.Ui64(v122)>>(uint(v108)%64)) + int64(base.Ui64(v129)>>(uint(v108)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v17))) = v118&v114 | v129<<(uint(v108)%64)
								v141 = v71 - int32(2)
								v142 = *(*int64)(unsafe.Add(mBase, uint32(v17)))
								v143 = v142 + v80
								*(*uint16)(unsafe.Add(mBase, uint32(v141))) = uint16(v143)
								v147 = int64(0)
								v152 = v69 + int32(1)
								if v79 == v147 {
									v153 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v80))
								} else {
									v153 = base.B2i32(v79 != v147)
								}
								if v153 != 0 {
									v69 = v152
									v71 = v141
									v79 = v101
									v80 = v100
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v17)+52)) = v141
							v155 = v152
							v163 = v69
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v17)+36)) = v163
					*(*int32)(unsafe.Add(mBase, uint32(v17)+32)) = v155
					v171 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = int64(0)
					v175 = F_palloc(m, int32(12))
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v17)+72)) = v175
						v178 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v175))) = uint16(v178)
						*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v175 + int32(2)
						if v171 < int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = int64(16384)
							v196 = int64(0) - v171
							v199 = v178
							v201 = v175 + int32(12)
							v209 = v196
							for {
								v214 = v201 - int32(2)
								v216 = base.I64_div_u_s(v209, int64(10000))
								v219 = v216*int64(55536) + v209
								*(*uint16)(unsafe.Add(mBase, uint32(v214))) = uint16(v219)
								v222 = v199 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v209) {
									v199 = v222
									v201 = v214
									v209 = v216
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v214
							v226 = v222
							v232 = v199
						} else {
							v190 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v17)+64)) = v190
							if v171 == v190 {
								v226 = v178
								v232 = int32(0)
							} else {
								v196 = v171
								v199 = v178
								v201 = v175 + int32(12)
								v209 = v196
								for {
									v214 = v201 - int32(2)
									v216 = base.I64_div_u_s(v209, int64(10000))
									v219 = v216*int64(55536) + v209
									*(*uint16)(unsafe.Add(mBase, uint32(v214))) = uint16(v219)
									v222 = v199 + int32(1)
									if base.Ui64(int64(9999)) < base.Ui64(v209) {
										v199 = v222
										v201 = v214
										v209 = v216
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v17)+76)) = v214
								v226 = v222
								v232 = v199
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v17)+60)) = v232
						*(*int32)(unsafe.Add(mBase, uint32(v17)+56)) = v226
						v245 = F_make_result_opt_error(m, v17+int32(56), int32(0))
						mBase = m.M
						v246 = m.ExcPending
						if v246 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v175)
							mBase = m.M
							v248 = m.ExcPending
							if v248 != 0 {
								return int32(0)
							} else {
								v252 = F_make_result_opt_error(m, v17+int32(32), int32(0))
								mBase = m.M
								v253 = m.ExcPending
								if v253 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v35)
									mBase = m.M
									v255 = m.ExcPending
									if v255 != 0 {
										return int32(0)
									} else {
										v258 = F_DirectFunctionCall2Coll(m, int32(1279), int32(0), v252, v245)
										mBase = m.M
										v259 = m.ExcPending
										if v259 != 0 {
											return int32(0)
										} else {
											v274 = v258
											m.G0 = v17 + int32(80)
											return v274
										}
									}
								}
							}
						}
					}
				}
			} else {
				v27 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v27)
				v274 = int32(0)
				m.G0 = v17 + int32(80)
				return v274
			}
		}
	}
}
func F_numeric_poly_stddev_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v32 int64
	_ = v32
	var v33 int64
	_ = v33
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v40 int32
	_ = v40
	var v49 int64
	_ = v49
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int64
	_ = v81
	var v82 int64
	_ = v82
	var v86 int32
	_ = v86
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v95 int64
	_ = v95
	var v96 int64
	_ = v96
	var v103 int32
	_ = v103
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v106 int64
	_ = v106
	var v107 int64
	_ = v107
	var v112 int64
	_ = v112
	var v115 int64
	_ = v115
	var v118 int64
	_ = v118
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v126 int64
	_ = v126
	var v133 int64
	_ = v133
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v151 int64
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v164 int32
	_ = v164
	var v166 int32
	_ = v166
	var v183 int32
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v187 int32
	_ = v187
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v202 int64
	_ = v202
	var v210 int64
	_ = v210
	var v217 int64
	_ = v217
	var v218 int64
	_ = v218
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v236 int64
	_ = v236
	var v237 int64
	_ = v237
	var v241 int32
	_ = v241
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v261 int64
	_ = v261
	var v262 int64
	_ = v262
	var v267 int64
	_ = v267
	var v270 int64
	_ = v270
	var v273 int64
	_ = v273
	var v276 int64
	_ = v276
	var v277 int64
	_ = v277
	var v281 int64
	_ = v281
	var v288 int64
	_ = v288
	var v300 int32
	_ = v300
	var v301 int64
	_ = v301
	var v302 int64
	_ = v302
	var v306 int64
	_ = v306
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v362 int32
	_ = v362
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 int32
	_ = v379
	v5 = int32(0)
	v18 = m.G0
	v20 = v18 - int32(208)
	m.G0 = v20
	v27 = F__emscripten_memset_bulkmem(m, v20+int32(96), base.I32_extend8_s(v5), int32(112))
	mBase = m.M
	if l0 != 0 {
		v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v20)+104)) = v28
		v32 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v33 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		v35 = F_palloc(m, int32(22))
		mBase = m.M
		v38 = m.ExcPending
		if v38 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v35
			v40 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v35))) = uint16(v40)
			*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v35 + int32(2)
			if v32 < int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = int64(16384)
				v49 = int64(0)
				v62 = v49 - (v32 + base.I64_extend_i32_u(base.B2i32(v33 != v49)))
				v63 = v49 - v33
				v73 = v5
				v74 = v35 + int32(22)
				v81 = v62
				v82 = v63
				for {
					v86 = v20 + int32(56)
					v89 = m.G0
					v90 = int32(16)
					v91 = v89 - v90
					m.G0 = v91
					F___udivmodti4(m, v91, v82, v81, int64(10000), int64(0))
					mBase = m.M
					v95 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
					v96 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v86)+8)) = v96
					*(*int64)(unsafe.Add(mBase, uint32(v86))) = v95
					m.G0 = v91 + v90
					v103 = v20 + int32(40)
					v104 = *(*int64)(unsafe.Add(mBase, uint32(v20)+56))
					v105 = *(*int64)(unsafe.Add(mBase, uint32(v20-int32(-64))))
					v106 = int64(55536)
					v107 = int64(0)
					v112 = int64(32)
					v115 = int64(base.Ui64(v104) >> (uint(v112) % 64))
					v118 = int64(4294967295)
					v121 = v104 & v118
					v122 = v106 * v121
					v126 = int64(base.Ui64(v122)>>(uint(v112)%64)) + v106*v115
					v133 = v121*v107 + v126&v118
					*(*int64)(unsafe.Add(mBase, uint32(v103)+8)) = v104*v107 + v105*v106 + v107*v115 + int64(base.Ui64(v126)>>(uint(v112)%64)) + int64(base.Ui64(v133)>>(uint(v112)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v103))) = v122&v118 | v133<<(uint(v112)%64)
					v145 = v74 - int32(2)
					v146 = *(*int64)(unsafe.Add(mBase, uint32(v20)+40))
					v147 = v146 + v82
					*(*uint16)(unsafe.Add(mBase, uint32(v145))) = uint16(v147)
					v151 = int64(0)
					v156 = v73 + int32(1)
					if v81 == v151 {
						v157 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v82))
					} else {
						v157 = base.B2i32(v81 != v151)
					}
					if v157 != 0 {
						v73 = v156
						v74 = v145
						v81 = v105
						v82 = v104
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v145
				v164 = v156
				v166 = v73
			} else {
				v57 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = v57
				if v32|v33 != v57 {
					v62 = v32
					v63 = v33
					v73 = v5
					v74 = v35 + int32(22)
					v81 = v62
					v82 = v63
					for {
						v86 = v20 + int32(56)
						v89 = m.G0
						v90 = int32(16)
						v91 = v89 - v90
						m.G0 = v91
						F___udivmodti4(m, v91, v82, v81, int64(10000), int64(0))
						mBase = m.M
						v95 = *(*int64)(unsafe.Add(mBase, uint32(v91)))
						v96 = *(*int64)(unsafe.Add(mBase, uint32(v91)+8))
						*(*int64)(unsafe.Add(mBase, uint32(v86)+8)) = v96
						*(*int64)(unsafe.Add(mBase, uint32(v86))) = v95
						m.G0 = v91 + v90
						v103 = v20 + int32(40)
						v104 = *(*int64)(unsafe.Add(mBase, uint32(v20)+56))
						v105 = *(*int64)(unsafe.Add(mBase, uint32(v20-int32(-64))))
						v106 = int64(55536)
						v107 = int64(0)
						v112 = int64(32)
						v115 = int64(base.Ui64(v104) >> (uint(v112) % 64))
						v118 = int64(4294967295)
						v121 = v104 & v118
						v122 = v106 * v121
						v126 = int64(base.Ui64(v122)>>(uint(v112)%64)) + v106*v115
						v133 = v121*v107 + v126&v118
						*(*int64)(unsafe.Add(mBase, uint32(v103)+8)) = v104*v107 + v105*v106 + v107*v115 + int64(base.Ui64(v126)>>(uint(v112)%64)) + int64(base.Ui64(v133)>>(uint(v112)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v103))) = v122&v118 | v133<<(uint(v112)%64)
						v145 = v74 - int32(2)
						v146 = *(*int64)(unsafe.Add(mBase, uint32(v20)+40))
						v147 = v146 + v82
						*(*uint16)(unsafe.Add(mBase, uint32(v145))) = uint16(v147)
						v151 = int64(0)
						v156 = v73 + int32(1)
						if v81 == v151 {
							v157 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v82))
						} else {
							v157 = base.B2i32(v81 != v151)
						}
						if v157 != 0 {
							v73 = v156
							v74 = v145
							v81 = v105
							v82 = v104
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v145
					v164 = v156
					v166 = v73
				} else {
					v164 = v5
					v166 = v5
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v166
			*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v164
			F_accum_sum_add(m, v20+int32(112), v20+int32(72))
			mBase = m.M
			v183 = m.ExcPending
			if v183 != 0 {
				return int32(0)
			} else {
				v184 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				v185 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				F_pfree(m, v35)
				mBase = m.M
				v187 = m.ExcPending
				if v187 != 0 {
					return int32(0)
				} else {
					v189 = F_palloc(m, int32(22))
					mBase = m.M
					v190 = m.ExcPending
					if v190 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v20)+88)) = v189
						v192 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v189))) = uint16(v192)
						*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v189 + int32(2)
						if v184 < int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = int64(16384)
							v202 = int64(0)
							v217 = v202 - (v184 + base.I64_extend_i32_u(base.B2i32(v185 != v202)))
							v218 = v202 - v185
							v228 = v192
							v229 = v189 + int32(22)
							v236 = v217
							v237 = v218
							for {
								v241 = v20 + int32(24)
								v244 = m.G0
								v245 = int32(16)
								v246 = v244 - v245
								m.G0 = v246
								F___udivmodti4(m, v246, v237, v236, int64(10000), int64(0))
								mBase = m.M
								v250 = *(*int64)(unsafe.Add(mBase, uint32(v246)))
								v251 = *(*int64)(unsafe.Add(mBase, uint32(v246)+8))
								*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v251
								*(*int64)(unsafe.Add(mBase, uint32(v241))) = v250
								m.G0 = v246 + v245
								v258 = v20 + int32(8)
								v259 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
								v260 = *(*int64)(unsafe.Add(mBase, uint32(v20+int32(32))))
								v261 = int64(55536)
								v262 = int64(0)
								v267 = int64(32)
								v270 = int64(base.Ui64(v259) >> (uint(v267) % 64))
								v273 = int64(4294967295)
								v276 = v259 & v273
								v277 = v261 * v276
								v281 = int64(base.Ui64(v277)>>(uint(v267)%64)) + v261*v270
								v288 = v276*v262 + v281&v273
								*(*int64)(unsafe.Add(mBase, uint32(v258)+8)) = v259*v262 + v260*v261 + v262*v270 + int64(base.Ui64(v281)>>(uint(v267)%64)) + int64(base.Ui64(v288)>>(uint(v267)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v258))) = v277&v273 | v288<<(uint(v267)%64)
								v300 = v229 - int32(2)
								v301 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
								v302 = v301 + v237
								*(*uint16)(unsafe.Add(mBase, uint32(v300))) = uint16(v302)
								v306 = int64(0)
								v311 = v228 + int32(1)
								if v236 == v306 {
									v312 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v237))
								} else {
									v312 = base.B2i32(v236 != v306)
								}
								if v312 != 0 {
									v228 = v311
									v229 = v300
									v236 = v260
									v237 = v259
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v300
							v319 = v311
							v321 = v228
						} else {
							v210 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v20)+80)) = v210
							if v184|v185 == v210 {
								v319 = v192
								v321 = int32(0)
							} else {
								v217 = v184
								v218 = v185
								v228 = v192
								v229 = v189 + int32(22)
								v236 = v217
								v237 = v218
								for {
									v241 = v20 + int32(24)
									v244 = m.G0
									v245 = int32(16)
									v246 = v244 - v245
									m.G0 = v246
									F___udivmodti4(m, v246, v237, v236, int64(10000), int64(0))
									mBase = m.M
									v250 = *(*int64)(unsafe.Add(mBase, uint32(v246)))
									v251 = *(*int64)(unsafe.Add(mBase, uint32(v246)+8))
									*(*int64)(unsafe.Add(mBase, uint32(v241)+8)) = v251
									*(*int64)(unsafe.Add(mBase, uint32(v241))) = v250
									m.G0 = v246 + v245
									v258 = v20 + int32(8)
									v259 = *(*int64)(unsafe.Add(mBase, uint32(v20)+24))
									v260 = *(*int64)(unsafe.Add(mBase, uint32(v20+int32(32))))
									v261 = int64(55536)
									v262 = int64(0)
									v267 = int64(32)
									v270 = int64(base.Ui64(v259) >> (uint(v267) % 64))
									v273 = int64(4294967295)
									v276 = v259 & v273
									v277 = v261 * v276
									v281 = int64(base.Ui64(v277)>>(uint(v267)%64)) + v261*v270
									v288 = v276*v262 + v281&v273
									*(*int64)(unsafe.Add(mBase, uint32(v258)+8)) = v259*v262 + v260*v261 + v262*v270 + int64(base.Ui64(v281)>>(uint(v267)%64)) + int64(base.Ui64(v288)>>(uint(v267)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v258))) = v277&v273 | v288<<(uint(v267)%64)
									v300 = v229 - int32(2)
									v301 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
									v302 = v301 + v237
									*(*uint16)(unsafe.Add(mBase, uint32(v300))) = uint16(v302)
									v306 = int64(0)
									v311 = v228 + int32(1)
									if v236 == v306 {
										v312 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v237))
									} else {
										v312 = base.B2i32(v236 != v306)
									}
									if v312 != 0 {
										v228 = v311
										v229 = v300
										v236 = v260
										v237 = v259
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v20)+92)) = v300
								v319 = v311
								v321 = v228
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v20)+76)) = v321
						*(*int32)(unsafe.Add(mBase, uint32(v20)+72)) = v319
						F_accum_sum_add(m, v20+int32(140), v20+int32(72))
						mBase = m.M
						v338 = m.ExcPending
						if v338 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v189)
							mBase = m.M
							v340 = m.ExcPending
							if v340 != 0 {
								return int32(0)
							} else {
								v360 = F_numeric_stddev_internal(m, v20+int32(96), l1, l2, l3)
								mBase = m.M
								v361 = m.ExcPending
								if v361 != 0 {
									return int32(0)
								} else {
									v362 = *(*int32)(unsafe.Add(mBase, uint32(v20)+112))
									if int32(0) < v362 {
										v365 = *(*int32)(unsafe.Add(mBase, uint32(v20)+132))
										F_pfree(m, v365)
										mBase = m.M
										v367 = m.ExcPending
										if v367 != 0 {
											return int32(0)
										} else {
											v368 = *(*int32)(unsafe.Add(mBase, uint32(v20)+136))
											F_pfree(m, v368)
											mBase = m.M
											v370 = m.ExcPending
											if v370 != 0 {
												return int32(0)
											} else {
												v371 = *(*int32)(unsafe.Add(mBase, uint32(v20)+140))
												if int32(0) < v371 {
													v374 = *(*int32)(unsafe.Add(mBase, uint32(v20)+160))
													F_pfree(m, v374)
													mBase = m.M
													v376 = m.ExcPending
													if v376 != 0 {
														return int32(0)
													} else {
														v377 = *(*int32)(unsafe.Add(mBase, uint32(v20)+164))
														F_pfree(m, v377)
														mBase = m.M
														v379 = m.ExcPending
														if v379 != 0 {
															return int32(0)
														} else {
															m.G0 = v20 + int32(208)
															return v360
														}
													}
												} else {
													m.G0 = v20 + int32(208)
													return v360
												}
											}
										}
									} else {
										v371 = *(*int32)(unsafe.Add(mBase, uint32(v20)+140))
										if int32(0) < v371 {
											v374 = *(*int32)(unsafe.Add(mBase, uint32(v20)+160))
											F_pfree(m, v374)
											mBase = m.M
											v376 = m.ExcPending
											if v376 != 0 {
												return int32(0)
											} else {
												v377 = *(*int32)(unsafe.Add(mBase, uint32(v20)+164))
												F_pfree(m, v377)
												mBase = m.M
												v379 = m.ExcPending
												if v379 != 0 {
													return int32(0)
												} else {
													m.G0 = v20 + int32(208)
													return v360
												}
											}
										} else {
											m.G0 = v20 + int32(208)
											return v360
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
		v360 = F_numeric_stddev_internal(m, v20+int32(96), l1, l2, l3)
		mBase = m.M
		v361 = m.ExcPending
		if v361 != 0 {
			return int32(0)
		} else {
			v362 = *(*int32)(unsafe.Add(mBase, uint32(v20)+112))
			if int32(0) < v362 {
				v365 = *(*int32)(unsafe.Add(mBase, uint32(v20)+132))
				F_pfree(m, v365)
				mBase = m.M
				v367 = m.ExcPending
				if v367 != 0 {
					return int32(0)
				} else {
					v368 = *(*int32)(unsafe.Add(mBase, uint32(v20)+136))
					F_pfree(m, v368)
					mBase = m.M
					v370 = m.ExcPending
					if v370 != 0 {
						return int32(0)
					} else {
						v371 = *(*int32)(unsafe.Add(mBase, uint32(v20)+140))
						if int32(0) < v371 {
							v374 = *(*int32)(unsafe.Add(mBase, uint32(v20)+160))
							F_pfree(m, v374)
							mBase = m.M
							v376 = m.ExcPending
							if v376 != 0 {
								return int32(0)
							} else {
								v377 = *(*int32)(unsafe.Add(mBase, uint32(v20)+164))
								F_pfree(m, v377)
								mBase = m.M
								v379 = m.ExcPending
								if v379 != 0 {
									return int32(0)
								} else {
									m.G0 = v20 + int32(208)
									return v360
								}
							}
						} else {
							m.G0 = v20 + int32(208)
							return v360
						}
					}
				}
			} else {
				v371 = *(*int32)(unsafe.Add(mBase, uint32(v20)+140))
				if int32(0) < v371 {
					v374 = *(*int32)(unsafe.Add(mBase, uint32(v20)+160))
					F_pfree(m, v374)
					mBase = m.M
					v376 = m.ExcPending
					if v376 != 0 {
						return int32(0)
					} else {
						v377 = *(*int32)(unsafe.Add(mBase, uint32(v20)+164))
						F_pfree(m, v377)
						mBase = m.M
						v379 = m.ExcPending
						if v379 != 0 {
							return int32(0)
						} else {
							m.G0 = v20 + int32(208)
							return v360
						}
					}
				} else {
					m.G0 = v20 + int32(208)
					return v360
				}
			}
		}
	}
}
func F_numeric_poly_stddev_pop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 != 0 {
		v10 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v9
	}
	v11 = int32(0)
	v15 = F_numeric_poly_stddev_internal(m, v10, v11, v11, v6+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
	}
}
func F_numeric_poly_stddev_samp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 != 0 {
		v10 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v9
	}
	v15 = F_numeric_poly_stddev_internal(m, v10, int32(0), int32(1), v6+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
	}
}
func F_numeric_sum(m *base.Module, l0 int32) int32 {
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
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v18 int64
	_ = v18
	var v22 int64
	_ = v22
	var v29 int32
	_ = v29
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v12 != 0 {
		v29 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
		v84 = int32(0)
		m.G0 = v10 + int32(32)
		return v84
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v13 == int32(0) {
			v29 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
			v84 = int32(0)
			m.G0 = v10 + int32(32)
			return v84
		} else {
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v13)+96))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v13)+88))
			v18 = *(*int64)(unsafe.Add(mBase, uint32(v13)+8))
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v13)+104))
			if v16+(v17+v18) != int64(0)-v22 {
				if int64(0) < v17 {
					v36 = F_make_result_opt_error(m, int32(1774572), int32(0))
					mBase = m.M
					v39 = m.ExcPending
					if v39 != 0 {
						return int32(0)
					} else {
						v84 = v36
						m.G0 = v10 + int32(32)
						return v84
					}
				} else {
					if v16 <= int64(0) {
						if int64(0) < v16 {
							v52 = F_make_result_opt_error(m, int32(1774596), int32(0))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v84 = v52
								m.G0 = v10 + int32(32)
								return v84
							}
						} else {
							if int64(0) < v22 {
								v58 = F_make_result_opt_error(m, int32(1774620), int32(0))
								mBase = m.M
								v59 = m.ExcPending
								if v59 != 0 {
									return int32(0)
								} else {
									v84 = v58
									m.G0 = v10 + int32(32)
									return v84
								}
							} else {
								v61 = v10 + int32(24)
								v62 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v61))) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v62
								F_accum_sum_final(m, v13+int32(16), v10+int32(8))
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v77 = F_make_result_opt_error(m, v10+int32(8), int32(0))
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										v79 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
										if v79 == int32(0) {
											v84 = v77
											m.G0 = v10 + int32(32)
											return v84
										} else {
											F_pfree(m, v79)
											mBase = m.M
											v83 = m.ExcPending
											if v83 != 0 {
												return int32(0)
											} else {
												v84 = v77
												m.G0 = v10 + int32(32)
												return v84
											}
										}
									}
								}
							}
						}
					} else {
						if v22 <= int64(0) {
							if int64(0) < v16 {
								v52 = F_make_result_opt_error(m, int32(1774596), int32(0))
								mBase = m.M
								v53 = m.ExcPending
								if v53 != 0 {
									return int32(0)
								} else {
									v84 = v52
									m.G0 = v10 + int32(32)
									return v84
								}
							} else {
								if int64(0) < v22 {
									v58 = F_make_result_opt_error(m, int32(1774620), int32(0))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int32(0)
									} else {
										v84 = v58
										m.G0 = v10 + int32(32)
										return v84
									}
								} else {
									v61 = v10 + int32(24)
									v62 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v61))) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v10)+16)) = v62
									*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v62
									F_accum_sum_final(m, v13+int32(16), v10+int32(8))
									mBase = m.M
									v73 = m.ExcPending
									if v73 != 0 {
										return int32(0)
									} else {
										v77 = F_make_result_opt_error(m, v10+int32(8), int32(0))
										mBase = m.M
										v78 = m.ExcPending
										if v78 != 0 {
											return int32(0)
										} else {
											v79 = *(*int32)(unsafe.Add(mBase, uint32(v61)))
											if v79 == int32(0) {
												v84 = v77
												m.G0 = v10 + int32(32)
												return v84
											} else {
												F_pfree(m, v79)
												mBase = m.M
												v83 = m.ExcPending
												if v83 != 0 {
													return int32(0)
												} else {
													v84 = v77
													m.G0 = v10 + int32(32)
													return v84
												}
											}
										}
									}
								}
							}
						} else {
							v46 = F_make_result_opt_error(m, int32(1774572), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								v84 = v46
								m.G0 = v10 + int32(32)
								return v84
							}
						}
					}
				}
			} else {
				v29 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v29)
				v84 = int32(0)
				m.G0 = v10 + int32(32)
				return v84
			}
		}
	}
}
func F_numeric_uplus(m *base.Module, l0 int32) int32 {
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
	var v10 int32
	_ = v10
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v7 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
		v10 = F_palloc(m, int32(base.Ui32(v7)>>(uint(int32(2))%32)))
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int32(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(v3)))
			v14 = int32(base.Ui32(v12) >> (uint(int32(2)) % 32))
			if v14 != 0 {
				v15 = F__emscripten_memcpy_bulkmem(m, v10, v3, v14)
				mBase = m.M
				v16 = v15
			} else {
				v16 = v10
			}
			return v16
		}
	}
}
func F_numeric_var_pop(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 != 0 {
		v10 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v9
	}
	v15 = F_numeric_stddev_internal(m, v10, int32(1), int32(0), v6+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
	}
}
func F_numeric_var_samp(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	v4 = m.G0
	v6 = v4 - int32(16)
	m.G0 = v6
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v8 != 0 {
		v10 = int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v10 = v9
	}
	v11 = int32(1)
	v15 = F_numeric_stddev_internal(m, v10, v11, v11, v6+int32(15))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v6)+15)))
		if v19 == int32(1) {
			v22 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v22)
			v25 = int32(0)
		} else {
			v25 = v15
		}
		m.G0 = v6 + int32(16)
		return v25
	}
}
