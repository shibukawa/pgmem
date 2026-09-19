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
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v113 int64
	_ = v113
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v129 int32
	_ = v129
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
	if base.Ui32(int32(_a_F_do_numeric_accum_0)) <= base.Ui32(v11) {
		if v11 != int32(_a_F_do_numeric_accum_1) {
			if v11 != int32(_a_F_do_numeric_accum_2) {
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
			v65 = v11 & int32(_a_F_do_numeric_accum_3)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v65
		v72 = v11 & int32(_a_F_do_numeric_accum_0)
		if v72 == int32(_a_F_do_numeric_accum_4) {
			v75 = v11 << (uint(int32(1)) % 32) & int32(_a_F_do_numeric_accum_5)
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
			v106 = m.ExcPending
			if v106 != 0 {
				return
			} else {
				v108 = int32(_a_F_do_numeric_accum_6)
				v109 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0]))
				v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v111
				v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v113 + int64(1)
				F_accum_sum_add(m, l0+int32(16), v9+int32(24))
				mBase = m.M
				v122 = m.ExcPending
				if v122 != 0 {
					return
				} else {
					v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
					if v123 == int32(1) {
						F_accum_sum_add(m, l0+int32(44), v9)
						mBase = m.M
						v129 = m.ExcPending
						if v129 != 0 {
							return
						} else {
							*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v109
							m.G0 = v9 + int32(48)
							return
						}
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v109
						m.G0 = v9 + int32(48)
						return
					}
				}
			}
		} else {
			v108 = int32(_a_F_do_numeric_accum_6)
			v109 = *(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0]))
			v111 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v111
			v113 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
			*(*int64)(unsafe.Add(mBase, uint32(l0)+8)) = v113 + int64(1)
			F_accum_sum_add(m, l0+int32(16), v9+int32(24))
			mBase = m.M
			v122 = m.ExcPending
			if v122 != 0 {
				return
			} else {
				v123 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0))))
				if v123 == int32(1) {
					F_accum_sum_add(m, l0+int32(44), v9)
					mBase = m.M
					v129 = m.ExcPending
					if v129 != 0 {
						return
					} else {
						*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v109
						m.G0 = v9 + int32(48)
						return
					}
				} else {
					*(*int32)(unsafe.Add(mBase, _c_F_do_numeric_accum[0])) = v109
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
			F_errmsg_internal(m, int32(_a_F_numeric_accum_inv_0), int32(0))
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_accum_inv_1), int32(_a_F_numeric_accum_inv_2), int32(_a_F_numeric_accum_inv_3))
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
				F_errmsg_internal(m, int32(_a_F_numeric_accum_inv_0), int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					F_errfinish(m, int32(_a_F_numeric_accum_inv_1), int32(_a_F_numeric_accum_inv_2), int32(_a_F_numeric_accum_inv_3))
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
	var v43 int64
	_ = v43
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v55 int64
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v69 int64
	_ = v69
	var v71 int64
	_ = v71
	var v73 int64
	_ = v73
	var v75 int64
	_ = v75
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v114 int32
	_ = v114
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v136 int64
	_ = v136
	var v138 int64
	_ = v138
	var v140 int64
	_ = v140
	var v143 int64
	_ = v143
	var v145 int64
	_ = v145
	var v147 int64
	_ = v147
	var v149 int64
	_ = v149
	var v172 int32
	_ = v172
	var v175 int64
	_ = v175
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int64
	_ = v182
	var v184 int64
	_ = v184
	var v186 int64
	_ = v186
	var v189 int64
	_ = v189
	var v191 int64
	_ = v191
	var v193 int64
	_ = v193
	var v195 int64
	_ = v195
	var v218 int32
	_ = v218
	var v221 int64
	_ = v221
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v228 int64
	_ = v228
	var v230 int64
	_ = v230
	var v232 int64
	_ = v232
	var v235 int64
	_ = v235
	var v237 int64
	_ = v237
	var v239 int64
	_ = v239
	var v241 int64
	_ = v241
	var v264 int32
	_ = v264
	var v267 int64
	_ = v267
	var v270 int32
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v274 int64
	_ = v274
	var v276 int64
	_ = v276
	var v278 int64
	_ = v278
	var v281 int64
	_ = v281
	var v283 int64
	_ = v283
	var v285 int64
	_ = v285
	var v287 int64
	_ = v287
	var v314 int32
	_ = v314
	var v315 int32
	_ = v315
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v329 int32
	_ = v329
	var v333 int32
	_ = v333
	var v338 int32
	_ = v338
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
		v43 = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v43
		*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v43
		*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v43
		v50 = v9 + int32(32)
		F_pq_begintypsend(m, v50)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			v55 = *(*int64)(unsafe.Add(mBase, uint32(v42)+8))
			F_enlargeStringInfo(m, v50, int32(8))
			mBase = m.M
			v58 = m.ExcPending
			if v58 != 0 {
				return int32(0)
			} else {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
				v62 = int64(56)
				v64 = int64(65280)
				v66 = int64(40)
				v69 = int64(16711680)
				v71 = int64(24)
				v73 = int64(4278190080)
				v75 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v59+v60))) = v55<<(uint(v62)%64) | v55&v64<<(uint(v66)%64) | (v55&v69<<(uint(v71)%64) | v55&v73<<(uint(v75)%64)) | (int64(base.Ui64(v55)>>(uint(v75)%64))&v73 | int64(base.Ui64(v55)>>(uint(v71)%64))&v69 | (int64(base.Ui64(v55)>>(uint(v66)%64))&v64 | int64(base.Ui64(v55)>>(uint(v62)%64))))
				v98 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v59 + v98
				v104 = v9 + v98
				F_accum_sum_final(m, v42+int32(16), v104)
				mBase = m.M
				v106 = m.ExcPending
				if v106 != 0 {
					return int32(0)
				} else {
					F_numericvar_serialize(m, v50, v104)
					mBase = m.M
					v108 = m.ExcPending
					if v108 != 0 {
						return int32(0)
					} else {
						v109 = *(*int32)(unsafe.Add(mBase, uint32(v42)+72))
						F_enlargeStringInfo(m, v50, int32(4))
						mBase = m.M
						v112 = m.ExcPending
						if v112 != 0 {
							return int32(0)
						} else {
							v113 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
							v114 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
							v118 = int32(16711935)
							v122 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v113+v114))) = base.I32_rotr(v109, int32(24))&v118 | base.I32_rotr(v109&v118, v122)
							*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v113 + int32(4)
							v129 = *(*int64)(unsafe.Add(mBase, uint32(v42)+80))
							F_enlargeStringInfo(m, v50, v122)
							mBase = m.M
							v132 = m.ExcPending
							if v132 != 0 {
								return int32(0)
							} else {
								v133 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
								v134 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
								v136 = int64(56)
								v138 = int64(65280)
								v140 = int64(40)
								v143 = int64(16711680)
								v145 = int64(24)
								v147 = int64(4278190080)
								v149 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v133+v134))) = v129<<(uint(v136)%64) | v129&v138<<(uint(v140)%64) | (v129&v143<<(uint(v145)%64) | v129&v147<<(uint(v149)%64)) | (int64(base.Ui64(v129)>>(uint(v149)%64))&v147 | int64(base.Ui64(v129)>>(uint(v145)%64))&v143 | (int64(base.Ui64(v129)>>(uint(v140)%64))&v138 | int64(base.Ui64(v129)>>(uint(v136)%64))))
								v172 = int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v133 + v172
								v175 = *(*int64)(unsafe.Add(mBase, uint32(v42)+88))
								F_enlargeStringInfo(m, v50, v172)
								mBase = m.M
								v178 = m.ExcPending
								if v178 != 0 {
									return int32(0)
								} else {
									v179 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
									v182 = int64(56)
									v184 = int64(65280)
									v186 = int64(40)
									v189 = int64(16711680)
									v191 = int64(24)
									v193 = int64(4278190080)
									v195 = int64(8)
									*(*int64)(unsafe.Add(mBase, uint32(v179+v180))) = v175<<(uint(v182)%64) | v175&v184<<(uint(v186)%64) | (v175&v189<<(uint(v191)%64) | v175&v193<<(uint(v195)%64)) | (int64(base.Ui64(v175)>>(uint(v195)%64))&v193 | int64(base.Ui64(v175)>>(uint(v191)%64))&v189 | (int64(base.Ui64(v175)>>(uint(v186)%64))&v184 | int64(base.Ui64(v175)>>(uint(v182)%64))))
									v218 = int32(8)
									*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v179 + v218
									v221 = *(*int64)(unsafe.Add(mBase, uint32(v42)+96))
									F_enlargeStringInfo(m, v50, v218)
									mBase = m.M
									v224 = m.ExcPending
									if v224 != 0 {
										return int32(0)
									} else {
										v225 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
										v226 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
										v228 = int64(56)
										v230 = int64(65280)
										v232 = int64(40)
										v235 = int64(16711680)
										v237 = int64(24)
										v239 = int64(4278190080)
										v241 = int64(8)
										*(*int64)(unsafe.Add(mBase, uint32(v225+v226))) = v221<<(uint(v228)%64) | v221&v230<<(uint(v232)%64) | (v221&v235<<(uint(v237)%64) | v221&v239<<(uint(v241)%64)) | (int64(base.Ui64(v221)>>(uint(v241)%64))&v239 | int64(base.Ui64(v221)>>(uint(v237)%64))&v235 | (int64(base.Ui64(v221)>>(uint(v232)%64))&v230 | int64(base.Ui64(v221)>>(uint(v228)%64))))
										v264 = int32(8)
										*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v225 + v264
										v267 = *(*int64)(unsafe.Add(mBase, uint32(v42)+104))
										F_enlargeStringInfo(m, v50, v264)
										mBase = m.M
										v270 = m.ExcPending
										if v270 != 0 {
											return int32(0)
										} else {
											v271 = *(*int32)(unsafe.Add(mBase, uint32(v9)+36))
											v272 = *(*int32)(unsafe.Add(mBase, uint32(v9)+32))
											v274 = int64(56)
											v276 = int64(65280)
											v278 = int64(40)
											v281 = int64(16711680)
											v283 = int64(24)
											v285 = int64(4278190080)
											v287 = int64(8)
											*(*int64)(unsafe.Add(mBase, uint32(v271+v272))) = v267<<(uint(v274)%64) | v267&v276<<(uint(v278)%64) | (v267&v281<<(uint(v283)%64) | v267&v285<<(uint(v287)%64)) | (int64(base.Ui64(v267)>>(uint(v287)%64))&v285 | int64(base.Ui64(v267)>>(uint(v283)%64))&v281 | (int64(base.Ui64(v267)>>(uint(v278)%64))&v276 | int64(base.Ui64(v267)>>(uint(v274)%64))))
											*(*int32)(unsafe.Add(mBase, uint32(v9)+36)) = v271 + int32(8)
											v314 = *(*int32)(unsafe.Add(mBase, uint32(v50)))
											v315 = *(*int32)(unsafe.Add(mBase, uint32(v50)+4))
											*(*int32)(unsafe.Add(mBase, uint32(v314))) = v315 << (uint(int32(2)) % 32)
											v319 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
											if v319 != 0 {
												F_pfree(m, v319)
												mBase = m.M
												v321 = m.ExcPending
												if v321 != 0 {
													return int32(0)
												} else {
													m.G0 = v9 + int32(48)
													return v314
												}
											} else {
												m.G0 = v9 + int32(48)
												return v314
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
		v329 = m.ExcPending
		if v329 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_avg_serialize_0), int32(0))
			mBase = m.M
			v333 = m.ExcPending
			if v333 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_avg_serialize_1), int32(_a_F_numeric_avg_serialize_2), int32(_a_F_numeric_avg_serialize_3))
				mBase = m.M
				v338 = m.ExcPending
				if v338 != 0 {
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
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v61 int32
	_ = v61
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v73 int32
	_ = v73
	var v79 int32
	_ = v79
	var v85 int32
	_ = v85
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v97 int64
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v109 int32
	_ = v109
	var v113 int32
	_ = v113
	var v115 int32
	_ = v115
	var v116 int32
	_ = v116
	var v118 int64
	_ = v118
	var v119 int32
	_ = v119
	var v121 int64
	_ = v121
	var v122 int32
	_ = v122
	var v124 int64
	_ = v124
	var v125 int32
	_ = v125
	var v127 int64
	_ = v127
	var v128 int32
	_ = v128
	var v131 int32
	_ = v131
	var v132 int32
	_ = v132
	var v134 int32
	_ = v134
	var v142 int32
	_ = v142
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
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
			*(*int64)(unsafe.Add(mBase, uint32(v8)+24)) = int64(0)
			v48 = int32(1)
			v49 = v42 + v48
			v50 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v42))))
			v52 = v50 & v48
			if v50 == v48 {
				v58 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v49))))
				if v58 == int32(18) {
					v61 = int32(16)
				} else {
					v61 = int32(0)
				}
				if base.Ui32((v58-int32(1))&int32(255)) < base.Ui32(int32(3)) {
					v68 = int32(4)
				} else {
					v68 = v61
				}
				v79 = v68
			} else {
				v69 = int32(1)
				if v52 != 0 {
					v79 = int32(base.Ui32(v50)>>(uint(v69)%32)) - v69
				} else {
					v73 = *(*int32)(unsafe.Add(mBase, uint32(v42)))
					v79 = int32(base.Ui32(v73)>>(uint(int32(2))%32)) - int32(4)
				}
			}
			*(*int64)(unsafe.Add(mBase, uint32(v8)+40)) = int64(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+36)) = v79
			if v52 != 0 {
				v85 = v49
			} else {
				v85 = v42 + int32(4)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+32)) = v85
			v88 = F_palloc0(m, int32(112))
			mBase = m.M
			v89 = m.ExcPending
			if v89 != 0 {
				return int32(0)
			} else {
				v90 = int32(0)
				*(*uint8)(unsafe.Add(mBase, uint32(v88))) = uint8(v90)
				v93 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_deserialize[0]))
				*(*int32)(unsafe.Add(mBase, uint32(v88)+4)) = v93
				v96 = v8 + int32(32)
				v97 = F_pq_getmsgint64(m, v96)
				mBase = m.M
				v98 = m.ExcPending
				if v98 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v88)+8)) = v97
					v101 = v8 + int32(8)
					F_numericvar_deserialize(m, v96, v101)
					mBase = m.M
					v103 = m.ExcPending
					if v103 != 0 {
						return int32(0)
					} else {
						F_accum_sum_add(m, v88+int32(16), v101)
						mBase = m.M
						v107 = m.ExcPending
						if v107 != 0 {
							return int32(0)
						} else {
							F_numericvar_deserialize(m, v96, v101)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								F_accum_sum_add(m, v88+int32(44), v101)
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
									return int32(0)
								} else {
									v115 = F_pq_getmsgint(m, v96, int32(4))
									mBase = m.M
									v116 = m.ExcPending
									if v116 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v88)+72)) = v115
										v118 = F_pq_getmsgint64(m, v96)
										mBase = m.M
										v119 = m.ExcPending
										if v119 != 0 {
											return int32(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v88)+80)) = v118
											v121 = F_pq_getmsgint64(m, v96)
											mBase = m.M
											v122 = m.ExcPending
											if v122 != 0 {
												return int32(0)
											} else {
												*(*int64)(unsafe.Add(mBase, uint32(v88)+88)) = v121
												v124 = F_pq_getmsgint64(m, v96)
												mBase = m.M
												v125 = m.ExcPending
												if v125 != 0 {
													return int32(0)
												} else {
													*(*int64)(unsafe.Add(mBase, uint32(v88)+96)) = v124
													v127 = F_pq_getmsgint64(m, v96)
													mBase = m.M
													v128 = m.ExcPending
													if v128 != 0 {
														return int32(0)
													} else {
														*(*int64)(unsafe.Add(mBase, uint32(v88)+104)) = v127
														F_pq_getmsgend(m, v96)
														mBase = m.M
														v131 = m.ExcPending
														if v131 != 0 {
															return int32(0)
														} else {
															v132 = *(*int32)(unsafe.Add(mBase, uint32(v8)+24))
															if v132 != 0 {
																F_pfree(m, v132)
																mBase = m.M
																v134 = m.ExcPending
																if v134 != 0 {
																	return int32(0)
																} else {
																	m.G0 = v8 + int32(48)
																	return v88
																}
															} else {
																m.G0 = v8 + int32(48)
																return v88
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
		v142 = m.ExcPending
		if v142 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_deserialize_0), int32(0))
			mBase = m.M
			v146 = m.ExcPending
			if v146 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_deserialize_1), int32(_a_F_numeric_deserialize_2), int32(_a_F_numeric_deserialize_3))
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
}
func F_numeric_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int64
	_ = v14
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
	var v26 int32
	_ = v26
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
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
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v109 int32
	_ = v109
	var v111 int32
	_ = v111
	var v128 int32
	_ = v128
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v159 int32
	_ = v159
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v202 int32
	_ = v202
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v208 int32
	_ = v208
	var v218 int32
	_ = v218
	var v227 int32
	_ = v227
	var v230 int32
	_ = v230
	var v232 int32
	_ = v232
	var v249 int32
	_ = v249
	var v253 int32
	_ = v253
	var v254 int32
	_ = v254
	var v255 int32
	_ = v255
	var v256 int32
	_ = v256
	var v260 int32
	_ = v260
	var v274 int32
	_ = v274
	var v286 int32
	_ = v286
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v294 int32
	_ = v294
	var v298 int32
	_ = v298
	var v304 int32
	_ = v304
	var v315 int32
	_ = v315
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int64
	_ = v323
	var v332 int32
	_ = v332
	var v339 int32
	_ = v339
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v344 int64
	_ = v344
	var v352 int32
	_ = v352
	var v356 int32
	_ = v356
	var v363 int32
	_ = v363
	var v368 int32
	_ = v368
	var v374 int64
	_ = v374
	var v375 int64
	_ = v375
	var v383 int32
	_ = v383
	var v385 int32
	_ = v385
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v391 int32
	_ = v391
	var v396 int32
	_ = v396
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v412 int64
	_ = v412
	var v416 int32
	_ = v416
	var v418 int64
	_ = v418
	var v421 int64
	_ = v421
	var v424 int32
	_ = v424
	var v431 int32
	_ = v431
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v439 int32
	_ = v439
	var v441 int32
	_ = v441
	var v442 int32
	_ = v442
	var v444 int32
	_ = v444
	var v446 int32
	_ = v446
	var v456 int64
	_ = v456
	var v460 int32
	_ = v460
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v476 int64
	_ = v476
	var v482 int32
	_ = v482
	var v484 int32
	_ = v484
	var v494 int64
	_ = v494
	var v497 int32
	_ = v497
	var v499 int64
	_ = v499
	var v502 int64
	_ = v502
	var v505 int32
	_ = v505
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v528 int32
	_ = v528
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v541 int32
	_ = v541
	var v552 int64
	_ = v552
	var v553 int64
	_ = v553
	var v555 int64
	_ = v555
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v581 int64
	_ = v581
	var v582 int64
	_ = v582
	var v587 int32
	_ = v587
	var v588 int64
	_ = v588
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v607 int32
	_ = v607
	var v612 int32
	_ = v612
	var v618 int64
	_ = v618
	var v619 int64
	_ = v619
	var v622 int32
	_ = v622
	var v636 int32
	_ = v636
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v644 int32
	_ = v644
	var v649 int32
	_ = v649
	var v654 int32
	_ = v654
	var v656 int32
	_ = v656
	var v665 int64
	_ = v665
	var v669 int32
	_ = v669
	var v671 int64
	_ = v671
	var v674 int64
	_ = v674
	var v677 int32
	_ = v677
	var v684 int32
	_ = v684
	var v689 int32
	_ = v689
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v694 int32
	_ = v694
	var v695 int32
	_ = v695
	var v697 int32
	_ = v697
	var v699 int32
	_ = v699
	var v709 int64
	_ = v709
	var v713 int32
	_ = v713
	var v718 int32
	_ = v718
	var v720 int32
	_ = v720
	var v734 int64
	_ = v734
	var v740 int32
	_ = v740
	var v742 int32
	_ = v742
	var v752 int64
	_ = v752
	var v755 int32
	_ = v755
	var v757 int64
	_ = v757
	var v760 int64
	_ = v760
	var v763 int32
	_ = v763
	var v769 int32
	_ = v769
	var v770 int32
	_ = v770
	var v786 int32
	_ = v786
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v794 int32
	_ = v794
	var v799 int32
	_ = v799
	var v810 int64
	_ = v810
	var v811 int64
	_ = v811
	var v814 int32
	_ = v814
	var v819 int64
	_ = v819
	var v845 int64
	_ = v845
	var v846 int64
	_ = v846
	var v850 int32
	_ = v850
	var v853 int32
	_ = v853
	var v858 int32
	_ = v858
	var v864 int64
	_ = v864
	var v865 int64
	_ = v865
	var v869 int64
	_ = v869
	var v877 int32
	_ = v877
	var v881 int32
	_ = v881
	var v888 int32
	_ = v888
	var v893 int32
	_ = v893
	var v899 int64
	_ = v899
	var v900 int64
	_ = v900
	var v908 int32
	_ = v908
	var v910 int32
	_ = v910
	var v913 int32
	_ = v913
	var v914 int32
	_ = v914
	var v916 int32
	_ = v916
	var v921 int32
	_ = v921
	var v926 int32
	_ = v926
	var v928 int32
	_ = v928
	var v937 int64
	_ = v937
	var v941 int32
	_ = v941
	var v943 int64
	_ = v943
	var v946 int64
	_ = v946
	var v949 int32
	_ = v949
	var v956 int32
	_ = v956
	var v961 int32
	_ = v961
	var v962 int32
	_ = v962
	var v964 int32
	_ = v964
	var v966 int32
	_ = v966
	var v967 int32
	_ = v967
	var v969 int32
	_ = v969
	var v971 int32
	_ = v971
	var v981 int64
	_ = v981
	var v985 int32
	_ = v985
	var v992 int32
	_ = v992
	var v994 int32
	_ = v994
	var v1001 int64
	_ = v1001
	var v1007 int32
	_ = v1007
	var v1009 int32
	_ = v1009
	var v1019 int64
	_ = v1019
	var v1022 int32
	_ = v1022
	var v1024 int64
	_ = v1024
	var v1027 int64
	_ = v1027
	var v1030 int32
	_ = v1030
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1053 int32
	_ = v1053
	var v1057 int32
	_ = v1057
	var v1058 int32
	_ = v1058
	var v1061 int32
	_ = v1061
	var v1066 int32
	_ = v1066
	var v1077 int64
	_ = v1077
	var v1078 int64
	_ = v1078
	var v1080 int64
	_ = v1080
	var v1090 int32
	_ = v1090
	var v1092 int32
	_ = v1092
	var v1095 int32
	_ = v1095
	var v1100 int32
	_ = v1100
	var v1106 int64
	_ = v1106
	var v1107 int64
	_ = v1107
	var v1118 int32
	_ = v1118
	var v1124 int64
	_ = v1124
	var v1125 int64
	_ = v1125
	var v1129 int32
	_ = v1129
	var v1131 int32
	_ = v1131
	var v1133 int32
	_ = v1133
	var v1134 int32
	_ = v1134
	var v1136 int32
	_ = v1136
	var v1139 int32
	_ = v1139
	var v1149 int64
	_ = v1149
	var v1155 int64
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1162 int32
	_ = v1162
	var v1171 int64
	_ = v1171
	var v1175 int32
	_ = v1175
	var v1177 int64
	_ = v1177
	var v1180 int64
	_ = v1180
	var v1183 int32
	_ = v1183
	var v1189 int32
	_ = v1189
	var v1190 int32
	_ = v1190
	var v1206 int32
	_ = v1206
	var v1211 int32
	_ = v1211
	var v1212 int32
	_ = v1212
	var v1214 int32
	_ = v1214
	var v1216 int32
	_ = v1216
	var v1217 int32
	_ = v1217
	var v1219 int32
	_ = v1219
	var v1221 int32
	_ = v1221
	var v1231 int64
	_ = v1231
	var v1235 int32
	_ = v1235
	var v1253 int32
	_ = v1253
	var v1254 int32
	_ = v1254
	var v1259 int32
	_ = v1259
	var v1267 int32
	_ = v1267
	var v1269 int64
	_ = v1269
	var v1275 int32
	_ = v1275
	var v1277 int32
	_ = v1277
	var v1287 int64
	_ = v1287
	var v1290 int32
	_ = v1290
	var v1292 int64
	_ = v1292
	var v1295 int64
	_ = v1295
	var v1298 int32
	_ = v1298
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1321 int32
	_ = v1321
	var v1325 int32
	_ = v1325
	var v1326 int32
	_ = v1326
	var v1329 int32
	_ = v1329
	var v1331 int32
	_ = v1331
	var v1333 int32
	_ = v1333
	var v1351 int32
	_ = v1351
	var v1352 int32
	_ = v1352
	var v1357 int32
	_ = v1357
	var v1361 int32
	_ = v1361
	var v1379 int32
	_ = v1379
	var v1382 int32
	_ = v1382
	var v1384 int32
	_ = v1384
	var v1399 int32
	_ = v1399
	var v1403 int32
	_ = v1403
	var v1419 int32
	_ = v1419
	var v1433 int32
	_ = v1433
	var v1445 int32
	_ = v1445
	var v1446 int32
	_ = v1446
	var v1449 int32
	_ = v1449
	var v1456 int32
	_ = v1456
	var v1457 int32
	_ = v1457
	var v1458 int32
	_ = v1458
	var v1461 int32
	_ = v1461
	var v1462 int32
	_ = v1462
	var v1467 int32
	_ = v1467
	var v1471 int32
	_ = v1471
	var v1476 int32
	_ = v1476
	var v1478 int32
	_ = v1478
	var v1482 int32
	_ = v1482
	var v1499 int32
	_ = v1499
	var v1500 int32
	_ = v1500
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1512 int32
	_ = v1512
	var v1517 int32
	_ = v1517
	var v1519 int32
	_ = v1519
	var v1536 int32
	_ = v1536
	var v1540 int32
	_ = v1540
	v2 = int32(0)
	v14 = int64(0)
	v17 = m.G0
	v19 = v17 - int32(112)
	m.G0 = v19
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = v23
	goto L4
L1:
	;
	v64 = int32(255)
	v65 = v60 & v64
	if base.B2i32(v65 == int32(46))|base.B2i32(base.Ui32((v60-int32(48))&v64) < base.Ui32(int32(10))) == int32(0) {
		goto L12
	} else {
		goto L13
	}
L2:
	;
	v57 = v26 + int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v57
	v59 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26)+1)))
	v60 = v59
	v61 = v54
	v62 = v57
	v63 = v55
	goto L1
L3:
	;
	v54 = int32(_a_F_numeric_in_0)
	v55 = int32(_a_F_numeric_in_1)
	goto L2
L4:
	;
	v40 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v26))))
	if base.Ui32(v40-int32(9)) < base.Ui32(int32(5)) {
		goto L7
	} else {
		goto L8
	}
L5:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v26
	v60 = v40
	v61 = v45
	v62 = v26
	v63 = v2
	goto L1
L6:
	;
	goto L5
L7:
	;
	v26 = v26 + int32(1)
	goto L4
L8:
	;
	v45 = int32(_a_F_numeric_in_2)
	switch v40 - int32(32) {
	case 0:
		goto L7
	default:
		goto L6
	case 11:
		v54 = v45
		v55 = v2
		goto L2
	case 13:
		goto L3
	}
L9:
	;
	m.G0 = v1536 + int32(112)
	return v1540
L10:
	;
	v1536 = v1519
	v1540 = int32(0)
	goto L9
L11:
	;
	v1499 = int32(0)
	v1500 = F_errsave_start(m, v21)
	mBase = m.M
	v1501 = m.ExcPending
	if v1501 != 0 {
		goto L35
	} else {
		goto L290
	}
L12:
	;
	v81 = v26
	v82 = int32(_a_F_numeric_in_3)
	v83 = int32(3)
	goto L17
L13:
	;
	goto L14
L14:
	;
	v323 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v19)+64)) = v323
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v323
	if v65 != int32(48) {
		goto L103
	} else {
		goto L104
	}
L15:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v255
	v260 = v255
	goto L75
L16:
	;
	if v128 == int32(0) {
		goto L32
	} else {
		goto L33
	}
L17:
	;
	if v83 != 0 {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	v128 = int32(0)
	goto L16
L19:
	;
	v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v81))))
	v87 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v82))))
	if v86 == v87 {
		v109 = v86
		goto L22
	} else {
		goto L23
	}
L20:
	;
	goto L21
L21:
	;
	goto L18
L22:
	;
	v111 = int32(1)
	if v109 != 0 {
		v81 = v81 + v111
		v82 = v82 + v111
		v83 = v83 - v111
		goto L17
	} else {
		goto L31
	}
L23:
	;
	if base.Ui32((v86-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L24
	} else {
		goto L25
	}
L24:
	;
	v97 = v86 | int32(32)
	goto L26
L25:
	;
	v97 = v86
	goto L26
L26:
	;
	if base.Ui32((v87-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L27
	} else {
		goto L28
	}
L27:
	;
	v106 = v87 | int32(32)
	goto L29
L28:
	;
	v106 = v87
	goto L29
L29:
	;
	if v97 == v106 {
		v109 = v97
		goto L22
	} else {
		goto L30
	}
L30:
	;
	v128 = v97 - v106
	goto L16
L31:
	;
	goto L21
L32:
	;
	v135 = F_make_result_opt_error(m, int32(_a_F_numeric_in_4), int32(0))
	mBase = m.M
	v138 = m.ExcPending
	if v138 != 0 {
		goto L35
	} else {
		goto L36
	}
L33:
	;
	goto L34
L34:
	;
	v143 = v62
	v144 = int32(_a_F_numeric_in_5)
	v145 = int32(8)
	goto L38
L35:
	;
	return int32(0)
L36:
	;
	v255 = v26 + int32(3)
	v256 = v135
	goto L15
L37:
	;
	if v190 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L38:
	;
	if v145 != 0 {
		goto L40
	} else {
		goto L41
	}
L39:
	;
	v190 = int32(0)
	goto L37
L40:
	;
	v148 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v143))))
	v149 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v144))))
	if v148 == v149 {
		v171 = v148
		goto L43
	} else {
		goto L44
	}
L41:
	;
	goto L42
L42:
	;
	goto L39
L43:
	;
	v173 = int32(1)
	if v171 != 0 {
		v143 = v143 + v173
		v144 = v144 + v173
		v145 = v145 - v173
		goto L38
	} else {
		goto L52
	}
L44:
	;
	if base.Ui32((v148-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v159 = v148 | int32(32)
	goto L47
L46:
	;
	v159 = v148
	goto L47
L47:
	;
	if base.Ui32((v149-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L48
	} else {
		goto L49
	}
L48:
	;
	v168 = v149 | int32(32)
	goto L50
L49:
	;
	v168 = v149
	goto L50
L50:
	;
	if v159 == v168 {
		v171 = v159
		goto L43
	} else {
		goto L51
	}
L51:
	;
	v190 = v159 - v168
	goto L37
L52:
	;
	goto L42
L53:
	;
	v196 = F_make_result_opt_error(m, v61, int32(0))
	mBase = m.M
	v197 = m.ExcPending
	if v197 != 0 {
		goto L35
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	v202 = v62
	v203 = int32(_a_F_numeric_in_6)
	v204 = int32(3)
	goto L58
L56:
	;
	v255 = v62 + int32(8)
	v256 = v196
	goto L15
L57:
	;
	if v249 != 0 {
		goto L11
	} else {
		goto L73
	}
L58:
	;
	if v204 != 0 {
		goto L60
	} else {
		goto L61
	}
L59:
	;
	v249 = int32(0)
	goto L57
L60:
	;
	v207 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v202))))
	v208 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v203))))
	if v207 == v208 {
		v230 = v207
		goto L63
	} else {
		goto L64
	}
L61:
	;
	goto L62
L62:
	;
	goto L59
L63:
	;
	v232 = int32(1)
	if v230 != 0 {
		v202 = v202 + v232
		v203 = v203 + v232
		v204 = v204 - v232
		goto L58
	} else {
		goto L72
	}
L64:
	;
	if base.Ui32((v207-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v218 = v207 | int32(32)
	goto L67
L66:
	;
	v218 = v207
	goto L67
L67:
	;
	if base.Ui32((v208-int32(65))&int32(255)) < base.Ui32(int32(26)) {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v227 = v208 | int32(32)
	goto L70
L69:
	;
	v227 = v208
	goto L70
L70:
	;
	if v218 == v227 {
		v230 = v218
		goto L63
	} else {
		goto L71
	}
L71:
	;
	v249 = v218 - v227
	goto L57
L72:
	;
	goto L62
L73:
	;
	v253 = F_make_result_opt_error(m, v61, int32(0))
	mBase = m.M
	v254 = m.ExcPending
	if v254 != 0 {
		goto L35
	} else {
		goto L74
	}
L74:
	;
	v255 = v62 + int32(3)
	v256 = v253
	goto L15
L75:
	;
	v274 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v260))))
	if base.B2i32(base.Ui32(v274-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v274 == int32(32)) != 0 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v260 = v260 + int32(1)
	goto L75
L78:
	;
	if v274 != 0 {
		goto L11
	} else {
		goto L80
	}
L80:
	;
	if v22 < int32(4) {
		v1536 = v19
		v1540 = v256
		goto L9
	} else {
		goto L81
	}
L81:
	;
	v286 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v256)+4)))
	if v286 == int32(_a_F_numeric_in_7) {
		v1536 = v19
		v1540 = v256
		goto L9
	} else {
		goto L82
	}
L82:
	;
	v290 = F_errsave_start(m, v21)
	mBase = m.M
	v291 = m.ExcPending
	if v291 != 0 {
		goto L35
	} else {
		goto L83
	}
L83:
	;
	if v290 != 0 {
		goto L84
	} else {
		goto L85
	}
L84:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L35
	} else {
		goto L87
	}
L85:
	;
	goto L86
L86:
	;
	v321 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v321)
	v1536 = v19
	v1540 = int32(0)
	goto L9
L87:
	;
	F_errmsg(m, int32(_a_F_numeric_in_8), int32(0))
	mBase = m.M
	v298 = m.ExcPending
	if v298 != 0 {
		goto L35
	} else {
		goto L88
	}
L88:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+32)) = int32(base.Ui32(v22-int32(4)) >> (uint(int32(16)) % 32))
	v304 = int32(21)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+36)) = (v22<<(uint(v304)%32) - int32(_a_F_numeric_in_9)) >> (uint(v304) % 32)
	F_errdetail(m, int32(_a_F_numeric_in_10), v19+int32(32))
	mBase = m.M
	v315 = m.ExcPending
	if v315 != 0 {
		goto L35
	} else {
		goto L89
	}
L89:
	;
	F_errsave_finish(m, v21, int32(_a_F_numeric_in_11), int32(_a_F_numeric_in_12), int32(_a_F_numeric_in_13))
	mBase = m.M
	v320 = m.ExcPending
	if v320 != 0 {
		goto L35
	} else {
		goto L90
	}
L90:
	;
	goto L86
L91:
	;
	v1419 = v1403
	goto L269
L92:
	;
	v1399 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1399)
	v1519 = v1384
	goto L10
L93:
	;
	F_errsave_finish(m, v21, int32(_a_F_numeric_in_11), v1379, int32(_a_F_numeric_in_14))
	mBase = m.M
	v1382 = m.ExcPending
	if v1382 != 0 {
		goto L35
	} else {
		goto L268
	}
L94:
	;
	v1351 = F_errsave_start(m, v21)
	mBase = m.M
	v1352 = m.ExcPending
	if v1352 != 0 {
		goto L35
	} else {
		goto L264
	}
L95:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v1305
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v1304
	v1321 = v19 + int32(56)
	F_add_var(m, v1321, v19+int32(88), v1321)
	mBase = m.M
	v1325 = m.ExcPending
	if v1325 != 0 {
		goto L35
	} else {
		goto L258
	}
L96:
	;
	v1275 = int32(0)
	v1277 = v1221 + int32(12)
	v1287 = v1269
	goto L255
L97:
	;
	v1253 = F_errsave_start(m, v21)
	mBase = m.M
	v1254 = m.ExcPending
	if v1254 != 0 {
		goto L35
	} else {
		goto L251
	}
L98:
	;
	if v1118 == int32(2) {
		goto L97
	} else {
		goto L226
	}
L99:
	;
	v869 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+104)) = v869
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v869
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v869
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v869
	v877 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v877
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v869
	v881 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)))
	if v881 == v877 {
		goto L97
	} else {
		goto L189
	}
L100:
	;
	v588 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+104)) = v588
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v588
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v588
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v588
	v596 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v596
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v588
	v600 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)))
	if v600 == v596 {
		goto L97
	} else {
		goto L144
	}
L101:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v63
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v19)+84))
	v1403 = v587
	goto L91
L102:
	;
	v344 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+104)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v19)+88)) = v344
	*(*int64)(unsafe.Add(mBase, uint32(v19)+72)) = v344
	v352 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v352
	*(*int64)(unsafe.Add(mBase, uint32(v19)+56)) = v344
	v356 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+2)))
	if v356 == v352 {
		goto L97
	} else {
		goto L107
	}
L103:
	;
	v339 = F_set_var_from_str(m, v23, v62, v19+int32(56), v19+int32(84), v21)
	mBase = m.M
	v340 = m.ExcPending
	if v340 != 0 {
		goto L35
	} else {
		goto L105
	}
L104:
	;
	v332 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62)+1)))
	switch v332 - int32(66) {
	case 0, 32:
		goto L102
	default:
		goto L103
	case 13, 45:
		goto L99
	case 22, 54:
		goto L100
	}
L105:
	;
	if v339 != 0 {
		goto L101
	} else {
		goto L106
	}
L106:
	;
	v341 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v341)
	v1536 = v19
	v1540 = int32(0)
	goto L9
L107:
	;
	v363 = v356
	v368 = int32(2)
	v374 = int64(1)
	v375 = v14
	goto L108
L108:
	;
	if v363&int32(254) == int32(48) {
		goto L114
	} else {
		goto L115
	}
L109:
	;
	v1118 = v575
	v1124 = v581
	v1125 = v582
	goto L98
L110:
	;
	if v570&int32(255) != 0 {
		v363 = v570
		v368 = v575
		v374 = v581
		v375 = v582
		goto L108
	} else {
		goto L143
	}
L111:
	;
	v555 = int64(1)
	v565 = v368 + int32(1)
	v567 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v565))))
	v570 = v567
	v575 = v565
	v581 = v552 << (uint(v555) % 64)
	v582 = base.I64_extend8_s(base.I64_extend_i32_u(v541)) + v553<<(uint(v555)%64) - int64(48)
	goto L110
L112:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v512
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v511
	v528 = v19 + int32(56)
	F_add_var(m, v528, v19+int32(88), v528)
	mBase = m.M
	v532 = m.ExcPending
	if v532 != 0 {
		goto L35
	} else {
		goto L141
	}
L113:
	;
	v482 = int32(0)
	v484 = v446 + int32(12)
	v494 = v476
	goto L138
L114:
	;
	if v374 < int64(4611686018427387904) {
		v541 = v363
		v552 = v374
		v553 = v375
		goto L111
	} else {
		goto L117
	}
L115:
	;
	goto L116
L116:
	;
	if v363&int32(255) != int32(95) {
		v1118 = v368
		v1124 = v374
		v1125 = v375
		goto L98
	} else {
		goto L136
	}
L117:
	;
	v383 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v383 != 0 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	F_pfree(m, v383)
	mBase = m.M
	v385 = m.ExcPending
	if v385 != 0 {
		goto L35
	} else {
		goto L121
	}
L119:
	;
	goto L120
L120:
	;
	v388 = F_palloc(m, int32(12))
	mBase = m.M
	v389 = m.ExcPending
	if v389 != 0 {
		goto L35
	} else {
		goto L122
	}
L121:
	;
	goto L120
L122:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v388
	v391 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v388))) = uint16(v391)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = int64(0)
	v396 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	v401 = v391
	v403 = v396 + int32(12)
	v412 = v374
	goto L123
L123:
	;
	v416 = v403 - int32(2)
	v418 = base.I64_div_u_s(v412, int64(10000))
	v421 = v418*int64(55536) + v412
	*(*uint16)(unsafe.Add(mBase, uint32(v416))) = uint16(v421)
	v424 = v401 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v412) {
		v401 = v424
		v403 = v416
		v412 = v418
		goto L123
	} else {
		goto L125
	}
L124:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v401
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v424
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v416
	v431 = v19 + int32(56)
	F_mul_var(m, v431, v19+int32(88), v431, int32(0))
	mBase = m.M
	v436 = m.ExcPending
	if v436 != 0 {
		goto L35
	} else {
		goto L126
	}
L125:
	;
	goto L124
L126:
	;
	v437 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v437 != 0 {
		goto L127
	} else {
		goto L128
	}
L127:
	;
	F_pfree(m, v437)
	mBase = m.M
	v439 = m.ExcPending
	if v439 != 0 {
		goto L35
	} else {
		goto L130
	}
L128:
	;
	goto L129
L129:
	;
	v441 = F_palloc(m, int32(12))
	mBase = m.M
	v442 = m.ExcPending
	if v442 != 0 {
		goto L35
	} else {
		goto L131
	}
L130:
	;
	goto L129
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v441
	v444 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v441))) = uint16(v444)
	v446 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v446 + int32(2)
	if v375 < int64(0) {
		goto L132
	} else {
		goto L133
	}
L132:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = int64(16384)
	v476 = int64(0) - v375
	goto L113
L133:
	;
	goto L134
L134:
	;
	v456 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v456
	if v375 != v456 {
		v476 = v375
		goto L113
	} else {
		goto L135
	}
L135:
	;
	v460 = int32(0)
	v511 = v460
	v512 = v460
	goto L112
L136:
	;
	v467 = v368 + int32(1)
	v469 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v467))))
	if base.Ui32((v469-int32(50))&int32(255)) < base.Ui32(int32(254)) {
		goto L97
	} else {
		goto L137
	}
L137:
	;
	v570 = v469
	v575 = v467
	v581 = v374
	v582 = v375
	goto L110
L138:
	;
	v497 = v484 - int32(2)
	v499 = base.I64_div_u_s(v494, int64(10000))
	v502 = v499*int64(55536) + v494
	*(*uint16)(unsafe.Add(mBase, uint32(v497))) = uint16(v502)
	v505 = v482 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v494) {
		v482 = v505
		v484 = v497
		v494 = v499
		goto L138
	} else {
		goto L140
	}
L139:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v497
	v511 = v505
	v512 = v482
	goto L112
L140:
	;
	goto L139
L141:
	;
	v533 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if int32(_a_F_numeric_in_15) < v533 {
		goto L94
	} else {
		goto L142
	}
L142:
	;
	v536 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v368))))
	v541 = v536
	v552 = int64(1)
	v553 = int64(0)
	goto L111
L143:
	;
	goto L109
L144:
	;
	v607 = v600
	v612 = int32(2)
	v618 = int64(1)
	v619 = v14
	goto L145
L145:
	;
	v622 = v607 & int32(255)
	goto L151
L146:
	;
	v1118 = v858
	v1124 = v864
	v1125 = v865
	goto L98
L147:
	;
	if v853&int32(255) != 0 {
		v607 = v853
		v612 = v858
		v618 = v864
		v619 = v865
		goto L145
	} else {
		goto L188
	}
L148:
	;
	v814 = v612 + int32(1)
	v819 = base.I64_extend8_s(base.I64_extend_i32_u(v799))
	if base.Ui32((v799-int32(48))&int32(255)) <= base.Ui32(int32(9)) {
		v846 = v819 - int64(48)
		goto L182
	} else {
		goto L183
	}
L149:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v770
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v769
	v786 = v19 + int32(56)
	F_add_var(m, v786, v19+int32(88), v786)
	mBase = m.M
	v790 = m.ExcPending
	if v790 != 0 {
		goto L35
	} else {
		goto L180
	}
L150:
	;
	v740 = int32(0)
	v742 = v699 + int32(12)
	v752 = v734
	goto L177
L151:
	;
	if base.B2i32(base.Ui32(v622-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v622|int32(32)-int32(97)) < base.Ui32(int32(6))) != 0 {
		goto L152
	} else {
		goto L153
	}
L152:
	;
	if v618 < int64(576460752303423488) {
		v799 = v607
		v810 = v618
		v811 = v619
		goto L148
	} else {
		goto L155
	}
L153:
	;
	goto L154
L154:
	;
	if v622 != int32(95) {
		v1118 = v612
		v1124 = v618
		v1125 = v619
		goto L98
	} else {
		goto L174
	}
L155:
	;
	v636 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v636 != 0 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	F_pfree(m, v636)
	mBase = m.M
	v638 = m.ExcPending
	if v638 != 0 {
		goto L35
	} else {
		goto L159
	}
L157:
	;
	goto L158
L158:
	;
	v641 = F_palloc(m, int32(12))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L35
	} else {
		goto L160
	}
L159:
	;
	goto L158
L160:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v641
	v644 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v641))) = uint16(v644)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = int64(0)
	v649 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	v654 = v644
	v656 = v649 + int32(12)
	v665 = v618
	goto L161
L161:
	;
	v669 = v656 - int32(2)
	v671 = base.I64_div_u_s(v665, int64(10000))
	v674 = v671*int64(55536) + v665
	*(*uint16)(unsafe.Add(mBase, uint32(v669))) = uint16(v674)
	v677 = v654 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v665) {
		v654 = v677
		v656 = v669
		v665 = v671
		goto L161
	} else {
		goto L163
	}
L162:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v654
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v677
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v669
	v684 = v19 + int32(56)
	F_mul_var(m, v684, v19+int32(88), v684, int32(0))
	mBase = m.M
	v689 = m.ExcPending
	if v689 != 0 {
		goto L35
	} else {
		goto L164
	}
L163:
	;
	goto L162
L164:
	;
	v690 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v690 != 0 {
		goto L165
	} else {
		goto L166
	}
L165:
	;
	F_pfree(m, v690)
	mBase = m.M
	v692 = m.ExcPending
	if v692 != 0 {
		goto L35
	} else {
		goto L168
	}
L166:
	;
	goto L167
L167:
	;
	v694 = F_palloc(m, int32(12))
	mBase = m.M
	v695 = m.ExcPending
	if v695 != 0 {
		goto L35
	} else {
		goto L169
	}
L168:
	;
	goto L167
L169:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v694
	v697 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v694))) = uint16(v697)
	v699 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v699 + int32(2)
	if v619 < int64(0) {
		goto L170
	} else {
		goto L171
	}
L170:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = int64(16384)
	v734 = int64(0) - v619
	goto L150
L171:
	;
	goto L172
L172:
	;
	v709 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v709
	if v619 != v709 {
		v734 = v619
		goto L150
	} else {
		goto L173
	}
L173:
	;
	v713 = int32(0)
	v769 = v713
	v770 = v713
	goto L149
L174:
	;
	v718 = v612 + int32(1)
	v720 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v718))))
	goto L175
L175:
	;
	if base.B2i32(base.Ui32(v720-int32(48)) < base.Ui32(int32(10)))|base.B2i32(base.Ui32(v720|int32(32)-int32(97)) < base.Ui32(int32(6))) == int32(0) {
		goto L97
	} else {
		goto L176
	}
L176:
	;
	v853 = v720
	v858 = v718
	v864 = v618
	v865 = v619
	goto L147
L177:
	;
	v755 = v742 - int32(2)
	v757 = base.I64_div_u_s(v752, int64(10000))
	v760 = v757*int64(55536) + v752
	*(*uint16)(unsafe.Add(mBase, uint32(v755))) = uint16(v760)
	v763 = v740 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v752) {
		v740 = v763
		v742 = v755
		v752 = v757
		goto L177
	} else {
		goto L179
	}
L178:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v755
	v769 = v763
	v770 = v740
	goto L149
L179:
	;
	goto L178
L180:
	;
	v791 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if int32(_a_F_numeric_in_15) < v791 {
		goto L94
	} else {
		goto L181
	}
L181:
	;
	v794 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v612))))
	v799 = v794
	v810 = int64(1)
	v811 = int64(0)
	goto L148
L182:
	;
	v850 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v814))))
	v853 = v850
	v858 = v814
	v864 = v810 << (uint(int64(4)) % 64)
	v865 = v846 + v811<<(uint(int64(4))%64)
	goto L147
L183:
	;
	if base.Ui32((v799-int32(97))&int32(255)) <= base.Ui32(int32(5)) {
		v846 = v819 - int64(87)
		goto L182
	} else {
		goto L184
	}
L184:
	;
	if base.Ui32(int32(6)) <= base.Ui32((v799-int32(65))&int32(255)) {
		goto L185
	} else {
		goto L186
	}
L185:
	;
	v845 = int64(-1)
	goto L187
L186:
	;
	v845 = v819 - int64(55)
	goto L187
L187:
	;
	v846 = v845
	goto L182
L188:
	;
	goto L146
L189:
	;
	v888 = v881
	v893 = int32(2)
	v899 = int64(1)
	v900 = v14
	goto L190
L190:
	;
	if v888&int32(248) == int32(48) {
		goto L196
	} else {
		goto L197
	}
L191:
	;
	v1118 = v1100
	v1124 = v1106
	v1125 = v1107
	goto L98
L192:
	;
	if v1095&int32(255) != 0 {
		v888 = v1095
		v893 = v1100
		v899 = v1106
		v900 = v1107
		goto L190
	} else {
		goto L225
	}
L193:
	;
	v1080 = int64(3)
	v1090 = v893 + int32(1)
	v1092 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v1090))))
	v1095 = v1092
	v1100 = v1090
	v1106 = v1077 << (uint(v1080) % 64)
	v1107 = base.I64_extend8_s(base.I64_extend_i32_u(v1066)) + v1078<<(uint(v1080)%64) - int64(48)
	goto L192
L194:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v1037
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v1036
	v1053 = v19 + int32(56)
	F_add_var(m, v1053, v19+int32(88), v1053)
	mBase = m.M
	v1057 = m.ExcPending
	if v1057 != 0 {
		goto L35
	} else {
		goto L223
	}
L195:
	;
	v1007 = int32(0)
	v1009 = v971 + int32(12)
	v1019 = v1001
	goto L220
L196:
	;
	if v899 < int64(1152921504606846976) {
		v1066 = v888
		v1077 = v899
		v1078 = v900
		goto L193
	} else {
		goto L199
	}
L197:
	;
	goto L198
L198:
	;
	if v888&int32(255) != int32(95) {
		v1118 = v893
		v1124 = v899
		v1125 = v900
		goto L98
	} else {
		goto L218
	}
L199:
	;
	v908 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v908 != 0 {
		goto L200
	} else {
		goto L201
	}
L200:
	;
	F_pfree(m, v908)
	mBase = m.M
	v910 = m.ExcPending
	if v910 != 0 {
		goto L35
	} else {
		goto L203
	}
L201:
	;
	goto L202
L202:
	;
	v913 = F_palloc(m, int32(12))
	mBase = m.M
	v914 = m.ExcPending
	if v914 != 0 {
		goto L35
	} else {
		goto L204
	}
L203:
	;
	goto L202
L204:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v913
	v916 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v913))) = uint16(v916)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = int64(0)
	v921 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	v926 = v916
	v928 = v921 + int32(12)
	v937 = v899
	goto L205
L205:
	;
	v941 = v928 - int32(2)
	v943 = base.I64_div_u_s(v937, int64(10000))
	v946 = v943*int64(55536) + v937
	*(*uint16)(unsafe.Add(mBase, uint32(v941))) = uint16(v946)
	v949 = v926 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v937) {
		v926 = v949
		v928 = v941
		v937 = v943
		goto L205
	} else {
		goto L207
	}
L206:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v926
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v949
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v941
	v956 = v19 + int32(56)
	F_mul_var(m, v956, v19+int32(88), v956, int32(0))
	mBase = m.M
	v961 = m.ExcPending
	if v961 != 0 {
		goto L35
	} else {
		goto L208
	}
L207:
	;
	goto L206
L208:
	;
	v962 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v962 != 0 {
		goto L209
	} else {
		goto L210
	}
L209:
	;
	F_pfree(m, v962)
	mBase = m.M
	v964 = m.ExcPending
	if v964 != 0 {
		goto L35
	} else {
		goto L212
	}
L210:
	;
	goto L211
L211:
	;
	v966 = F_palloc(m, int32(12))
	mBase = m.M
	v967 = m.ExcPending
	if v967 != 0 {
		goto L35
	} else {
		goto L213
	}
L212:
	;
	goto L211
L213:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v966
	v969 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v966))) = uint16(v969)
	v971 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v971 + int32(2)
	if v900 < int64(0) {
		goto L214
	} else {
		goto L215
	}
L214:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = int64(16384)
	v1001 = int64(0) - v900
	goto L195
L215:
	;
	goto L216
L216:
	;
	v981 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v981
	if v900 != v981 {
		v1001 = v900
		goto L195
	} else {
		goto L217
	}
L217:
	;
	v985 = int32(0)
	v1036 = v985
	v1037 = v985
	goto L194
L218:
	;
	v992 = v893 + int32(1)
	v994 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v992))))
	if base.Ui32((v994-int32(56))&int32(255)) < base.Ui32(int32(248)) {
		goto L97
	} else {
		goto L219
	}
L219:
	;
	v1095 = v994
	v1100 = v992
	v1106 = v899
	v1107 = v900
	goto L192
L220:
	;
	v1022 = v1009 - int32(2)
	v1024 = base.I64_div_u_s(v1019, int64(10000))
	v1027 = v1024*int64(55536) + v1019
	*(*uint16)(unsafe.Add(mBase, uint32(v1022))) = uint16(v1027)
	v1030 = v1007 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v1019) {
		v1007 = v1030
		v1009 = v1022
		v1019 = v1024
		goto L220
	} else {
		goto L222
	}
L221:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v1022
	v1036 = v1030
	v1037 = v1007
	goto L194
L222:
	;
	goto L221
L223:
	;
	v1058 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if int32(_a_F_numeric_in_15) < v1058 {
		goto L94
	} else {
		goto L224
	}
L224:
	;
	v1061 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v62+v893))))
	v1066 = v1061
	v1077 = int64(1)
	v1078 = int64(0)
	goto L193
L225:
	;
	goto L191
L226:
	;
	v1129 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1129 != 0 {
		goto L227
	} else {
		goto L228
	}
L227:
	;
	F_pfree(m, v1129)
	mBase = m.M
	v1131 = m.ExcPending
	if v1131 != 0 {
		goto L35
	} else {
		goto L230
	}
L228:
	;
	goto L229
L229:
	;
	v1133 = F_palloc(m, int32(12))
	mBase = m.M
	v1134 = m.ExcPending
	if v1134 != 0 {
		goto L35
	} else {
		goto L231
	}
L230:
	;
	goto L229
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v1133
	v1136 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1133))) = uint16(v1136)
	v1139 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v1139 + int32(2)
	if v1124 < int64(0) {
		goto L234
	} else {
		goto L235
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v1190
	*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v1189
	v1206 = v19 + int32(56)
	F_mul_var(m, v1206, v19+int32(88), v1206, int32(0))
	mBase = m.M
	v1211 = m.ExcPending
	if v1211 != 0 {
		goto L35
	} else {
		goto L241
	}
L233:
	;
	v1160 = v1136
	v1162 = v1139 + int32(12)
	v1171 = v1155
	goto L238
L234:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = int64(16384)
	v1155 = int64(0) - v1124
	goto L233
L235:
	;
	goto L236
L236:
	;
	v1149 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v1149
	if v1124 == v1149 {
		v1189 = v1136
		v1190 = int32(0)
		goto L232
	} else {
		goto L237
	}
L237:
	;
	v1155 = v1124
	goto L233
L238:
	;
	v1175 = v1162 - int32(2)
	v1177 = base.I64_div_u_s(v1171, int64(10000))
	v1180 = v1177*int64(55536) + v1171
	*(*uint16)(unsafe.Add(mBase, uint32(v1175))) = uint16(v1180)
	v1183 = v1160 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v1171) {
		v1160 = v1183
		v1162 = v1175
		v1171 = v1177
		goto L238
	} else {
		goto L240
	}
L239:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v1175
	v1189 = v1183
	v1190 = v1160
	goto L232
L240:
	;
	goto L239
L241:
	;
	v1212 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1212 != 0 {
		goto L242
	} else {
		goto L243
	}
L242:
	;
	F_pfree(m, v1212)
	mBase = m.M
	v1214 = m.ExcPending
	if v1214 != 0 {
		goto L35
	} else {
		goto L245
	}
L243:
	;
	goto L244
L244:
	;
	v1216 = F_palloc(m, int32(12))
	mBase = m.M
	v1217 = m.ExcPending
	if v1217 != 0 {
		goto L35
	} else {
		goto L246
	}
L245:
	;
	goto L244
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+104)) = v1216
	v1219 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1216))) = uint16(v1219)
	v1221 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v1221 + int32(2)
	if v1125 < int64(0) {
		goto L247
	} else {
		goto L248
	}
L247:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = int64(16384)
	v1269 = int64(0) - v1125
	goto L96
L248:
	;
	goto L249
L249:
	;
	v1231 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v19)+96)) = v1231
	if v1125 != v1231 {
		v1269 = v1125
		goto L96
	} else {
		goto L250
	}
L250:
	;
	v1235 = int32(0)
	v1304 = v1235
	v1305 = v1235
	goto L95
L251:
	;
	if v1253 == int32(0) {
		v1384 = v19
		goto L92
	} else {
		goto L252
	}
L252:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1259 = m.ExcPending
	if v1259 != 0 {
		goto L35
	} else {
		goto L253
	}
L253:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+20)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v19)+16)) = int32(_a_F_numeric_in_16)
	F_errmsg(m, int32(_a_F_numeric_in_17), v19+int32(16))
	mBase = m.M
	v1267 = m.ExcPending
	if v1267 != 0 {
		goto L35
	} else {
		goto L254
	}
L254:
	;
	v1379 = int32(_a_F_numeric_in_18)
	goto L93
L255:
	;
	v1290 = v1277 - int32(2)
	v1292 = base.I64_div_u_s(v1287, int64(10000))
	v1295 = v1292*int64(55536) + v1287
	*(*uint16)(unsafe.Add(mBase, uint32(v1290))) = uint16(v1295)
	v1298 = v1275 + int32(1)
	if base.Ui64(int64(9999)) < base.Ui64(v1287) {
		v1275 = v1298
		v1277 = v1290
		v1287 = v1292
		goto L255
	} else {
		goto L257
	}
L256:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+108)) = v1290
	v1304 = v1298
	v1305 = v1275
	goto L95
L257:
	;
	goto L256
L258:
	;
	v1326 = *(*int32)(unsafe.Add(mBase, uint32(v19)+60))
	if int32(_a_F_numeric_in_15) < v1326 {
		goto L94
	} else {
		goto L259
	}
L259:
	;
	v1329 = v62 + v1118
	*(*int32)(unsafe.Add(mBase, uint32(v19)+64)) = v63
	v1331 = *(*int32)(unsafe.Add(mBase, uint32(v19)+104))
	if v1331 != 0 {
		goto L260
	} else {
		goto L261
	}
L260:
	;
	F_pfree(m, v1331)
	mBase = m.M
	v1333 = m.ExcPending
	if v1333 != 0 {
		goto L35
	} else {
		goto L263
	}
L261:
	;
	goto L262
L262:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+84)) = v1329
	v1403 = v1329
	goto L91
L263:
	;
	goto L262
L264:
	;
	if v1351 == int32(0) {
		v1384 = v19
		goto L92
	} else {
		goto L265
	}
L265:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1357 = m.ExcPending
	if v1357 != 0 {
		goto L35
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(_a_F_numeric_in_19), int32(0))
	mBase = m.M
	v1361 = m.ExcPending
	if v1361 != 0 {
		goto L35
	} else {
		goto L267
	}
L267:
	;
	v1379 = int32(_a_F_numeric_in_20)
	goto L93
L268:
	;
	v1384 = v19
	goto L92
L269:
	;
	v1433 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1419))))
	if base.B2i32(base.Ui32(v1433-int32(9)) < base.Ui32(int32(5)))|base.B2i32(v1433 == int32(32)) != 0 {
		goto L271
	} else {
		goto L272
	}
L271:
	;
	v1419 = v1419 + int32(1)
	goto L269
L272:
	;
	if v1433 != 0 {
		goto L11
	} else {
		goto L274
	}
L274:
	;
	v1445 = F_apply_typmod(m, v19+int32(56), v22, v21)
	mBase = m.M
	v1446 = m.ExcPending
	if v1446 != 0 {
		goto L35
	} else {
		goto L275
	}
L275:
	;
	if v1445 == int32(0) {
		goto L276
	} else {
		goto L277
	}
L276:
	;
	v1449 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v1449)
	v1536 = v19
	v1540 = int32(0)
	goto L9
L277:
	;
	goto L278
L278:
	;
	v1456 = F_make_result_opt_error(m, v19+int32(56), v19+int32(55))
	mBase = m.M
	v1457 = m.ExcPending
	if v1457 != 0 {
		goto L35
	} else {
		goto L279
	}
L279:
	;
	v1458 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19)+55)))
	if v1458 == int32(1) {
		goto L280
	} else {
		goto L281
	}
L280:
	;
	v1461 = F_errsave_start(m, v21)
	mBase = m.M
	v1462 = m.ExcPending
	if v1462 != 0 {
		goto L35
	} else {
		goto L283
	}
L281:
	;
	goto L282
L282:
	;
	v1478 = *(*int32)(unsafe.Add(mBase, uint32(v19)+72))
	if v1478 == int32(0) {
		v1536 = v19
		v1540 = v1456
		goto L9
	} else {
		goto L288
	}
L283:
	;
	if v1461 == int32(0) {
		v1519 = v19
		goto L10
	} else {
		goto L284
	}
L284:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1467 = m.ExcPending
	if v1467 != 0 {
		goto L35
	} else {
		goto L285
	}
L285:
	;
	F_errmsg(m, int32(_a_F_numeric_in_19), int32(0))
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L35
	} else {
		goto L286
	}
L286:
	;
	F_errsave_finish(m, v21, int32(_a_F_numeric_in_11), int32(795), int32(_a_F_numeric_in_21))
	mBase = m.M
	v1476 = m.ExcPending
	if v1476 != 0 {
		goto L35
	} else {
		goto L287
	}
L287:
	;
	v1536 = v19
	v1540 = int32(0)
	goto L9
L288:
	;
	F_pfree(m, v1478)
	mBase = m.M
	v1482 = m.ExcPending
	if v1482 != 0 {
		goto L35
	} else {
		goto L289
	}
L289:
	;
	v1536 = v19
	v1540 = v1456
	goto L9
L290:
	;
	if v1500 == int32(0) {
		v1536 = v19
		v1540 = v1499
		goto L9
	} else {
		goto L291
	}
L291:
	;
	F_errcode(m, int32(33685634))
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L35
	} else {
		goto L292
	}
L292:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v19)+4)) = v23
	*(*int32)(unsafe.Add(mBase, uint32(v19))) = int32(_a_F_numeric_in_16)
	F_errmsg(m, int32(_a_F_numeric_in_17), v19)
	mBase = m.M
	v1512 = m.ExcPending
	if v1512 != 0 {
		goto L35
	} else {
		goto L293
	}
L293:
	;
	F_errsave_finish(m, v21, int32(_a_F_numeric_in_11), int32(806), int32(_a_F_numeric_in_21))
	mBase = m.M
	v1517 = m.ExcPending
	if v1517 != 0 {
		goto L35
	} else {
		goto L294
	}
L294:
	;
	v1536 = v19
	v1540 = v1499
	goto L9
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
		if base.Ui32(int32(_a_F_numeric_int2_0)) <= base.Ui32(v16) {
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
						*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(_a_F_numeric_int2_1)
						F_errmsg(m, int32(_a_F_numeric_int2_2), v9)
						mBase = m.M
						v114 = m.ExcPending
						if v114 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_numeric_int2_3), int32(_a_F_numeric_int2_4), int32(_a_F_numeric_int2_5))
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
						*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = int32(_a_F_numeric_int2_1)
						F_errmsg(m, int32(_a_F_numeric_int2_6), v7+int32(-48))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_numeric_int2_3), int32(_a_F_numeric_int2_7), int32(_a_F_numeric_int2_5))
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
				v75 = v16 & int32(_a_F_numeric_int2_8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+52)) = v75
			v82 = v16 & int32(_a_F_numeric_int2_0)
			if v82 == int32(_a_F_numeric_int2_9) {
				v85 = v16 << (uint(int32(1)) % 32) & int32(_a_F_numeric_int2_10)
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
							F_errmsg(m, int32(_a_F_numeric_int2_11), int32(0))
							mBase = m.M
							v130 = m.ExcPending
							if v130 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_numeric_int2_3), int32(_a_F_numeric_int2_12), int32(_a_F_numeric_int2_5))
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
								F_errmsg(m, int32(_a_F_numeric_int2_11), int32(0))
								mBase = m.M
								v146 = m.ExcPending
								if v146 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_numeric_int2_3), int32(_a_F_numeric_int2_13), int32(_a_F_numeric_int2_5))
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
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v74 int32
	_ = v74
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v120 int32
	_ = v120
	var v127 int32
	_ = v127
	var v132 int32
	_ = v132
	var v135 int32
	_ = v135
	var v138 int32
	_ = v138
	var v144 int32
	_ = v144
	var v148 int32
	_ = v148
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
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
			v22 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v9)+4)))
			v23 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v4)+4)))
			v24 = int32(_a_F_numeric_larger_0)
			v25 = v23 & v24
			if v25 == v24 {
				if v23 != int32(_a_F_numeric_larger_1) {
					if v23 != int32(_a_F_numeric_larger_0) {
						if v22 != int32(_a_F_numeric_larger_2) {
							v44 = int32(-1)
						} else {
							v44 = int32(0)
						}
						v161 = v44
					} else {
						v161 = base.B2i32(v22 != int32(_a_F_numeric_larger_0))
					}
				} else {
					if v22 == int32(_a_F_numeric_larger_0) {
						v39 = int32(-1)
					} else {
						v39 = base.B2i32(v22 != int32(_a_F_numeric_larger_1))
					}
					v161 = v39
				}
			} else {
				if base.Ui32(int32(_a_F_numeric_larger_0)) <= base.Ui32(v22) {
					if v22 == int32(_a_F_numeric_larger_2) {
						v51 = int32(1)
					} else {
						v51 = int32(-1)
					}
					v161 = v51
				} else {
					v53 = v4 + int32(6)
					v54 = *(*int32)(unsafe.Add(mBase, uint32(v4)))
					v61 = base.B2i32(int32(0) <= base.I32_extend16_s(v23))
					if int32(0) <= base.I32_extend16_s(v23) {
						v62 = int32(-8)
					} else {
						v62 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v23) {
						v64 = int32(*(*int16)(unsafe.Add(mBase, uint32(v53))))
						v74 = v64
					} else {
						v74 = v23<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v23&int32(63)
					}
					v76 = int32(base.Ui32(int32(base.Ui32(v54)>>(uint(int32(2))%32))+v62) >> (uint(int32(1)) % 32))
					v78 = v9 + int32(6)
					v79 = *(*int32)(unsafe.Add(mBase, uint32(v9)))
					v86 = base.B2i32(int32(0) <= base.I32_extend16_s(v22))
					if int32(0) <= base.I32_extend16_s(v22) {
						v87 = int32(-8)
					} else {
						v87 = int32(-6)
					}
					if int32(0) <= base.I32_extend16_s(v22) {
						v89 = int32(*(*int16)(unsafe.Add(mBase, uint32(v78))))
						v99 = v89
					} else {
						v99 = v22<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v22&int32(63)
					}
					v100 = int32(1)
					v101 = int32(base.Ui32(int32(base.Ui32(v79)>>(uint(int32(2))%32))+v87) >> (uint(v100) % 32))
					v107 = v22 & int32(_a_F_numeric_larger_0)
					if v107 == int32(_a_F_numeric_larger_3) {
						v110 = v22 << (uint(v100) % 32) & int32(_a_F_numeric_larger_4)
					} else {
						v110 = v107
					}
					if v76 == int32(0) {
						if v101 == int32(0) {
							v161 = int32(0)
						} else {
							if v110 == int32(_a_F_numeric_larger_4) {
								v120 = int32(1)
							} else {
								v120 = int32(-1)
							}
							v161 = v120
						}
					} else {
						if v25 == int32(_a_F_numeric_larger_3) {
							v127 = v23 << (uint(int32(1)) % 32) & int32(_a_F_numeric_larger_4)
						} else {
							v127 = v25
						}
						if v101 == int32(0) {
							if v127 != 0 {
								v132 = int32(-1)
							} else {
								v132 = int32(1)
							}
							v161 = v132
						} else {
							if int32(0) <= base.I32_extend16_s(v23) {
								v135 = v4 + int32(8)
							} else {
								v135 = v53
							}
							if int32(0) <= base.I32_extend16_s(v22) {
								v138 = v9 + int32(8)
							} else {
								v138 = v78
							}
							if v127 == int32(0) {
								if v110 == int32(_a_F_numeric_larger_4) {
									v161 = int32(1)
								} else {
									v144 = F_cmp_abs_common(m, v135, v76, v74, v138, v101, v99)
									mBase = m.M
									v161 = v144
								}
							} else {
								if v110 == int32(0) {
									v161 = int32(-1)
								} else {
									v148 = F_cmp_abs_common(m, v138, v101, v99, v135, v76, v74)
									mBase = m.M
									v161 = v148
								}
							}
						}
					}
				}
			}
			if int32(0) < v161 {
				v164 = v4
			} else {
				v164 = v9
			}
			return v164
		}
	}
}
func F_numeric_min_scale(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int32
	_ = v28
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v48 int32
	_ = v48
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v89 int32
	_ = v89
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = F_pg_detoast_datum(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v11 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7)+4)))
	if base.Ui32(int32(_a_F_numeric_min_scale_0)) <= base.Ui32(v11) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v14 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v14)
	return int32(0)
L4:
	;
	goto L5
L5:
	;
	v18 = *(*int32)(unsafe.Add(mBase, uint32(v7)))
	v23 = base.I32_extend16_s(v11)
	v25 = base.B2i32(int32(0) <= v23)
	if int32(0) <= v23 {
		goto L6
	} else {
		goto L7
	}
L6:
	;
	v26 = int32(-8)
	goto L8
L7:
	;
	v26 = int32(-6)
	goto L8
L8:
	;
	if int32(0) <= v23 {
		goto L9
	} else {
		goto L10
	}
L9:
	;
	v28 = int32(*(*int16)(unsafe.Add(mBase, uint32(v7)+6)))
	v38 = v28
	goto L11
L10:
	;
	v38 = v11<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v11&int32(63)
	goto L11
L11:
	;
	v41 = int32(0)
	if v23 < v41 {
		goto L12
	} else {
		goto L13
	}
L12:
	;
	v46 = int32(6)
	goto L14
L13:
	;
	v46 = int32(8)
	goto L14
L14:
	;
	v48 = int32(base.Ui32(int32(base.Ui32(v18)>>(uint(int32(2))%32))+v26) >> (uint(int32(1)) % 32))
	goto L16
L15:
	;
	return v89
L16:
	;
	if v48 <= int32(0) {
		v89 = v41
		goto L15
	} else {
		goto L18
	}
L17:
	;
	v65 = (v56 - v38) << (uint(int32(2)) % 32)
	if v65 <= int32(0) {
		v89 = v41
		goto L15
	} else {
		goto L20
	}
L18:
	;
	v55 = int32(1)
	v56 = v48 - v55
	v60 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v7+v46+v56<<(uint(v55)%32)))))
	if v60 == int32(0) {
		v48 = v56
		goto L16
	} else {
		goto L19
	}
L19:
	;
	goto L17
L20:
	;
	v70 = base.I32_rem_s(base.I32_extend16_s(v60), int32(10))
	if v70 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	return v65
L22:
	;
	goto L23
L23:
	;
	v73 = v60
	v74 = v65
	goto L24
L24:
	;
	v78 = v74 - int32(1)
	v80 = int32(10)
	v81 = base.I32_div_s(base.I32_extend16_s(v73), v80)
	v84 = base.I32_rem_s(base.I32_extend16_s(v81), v80)
	if v84 == int32(0) {
		v73 = v81
		v74 = v78
		goto L24
	} else {
		goto L26
	}
L25:
	;
	v89 = v78
	goto L15
L26:
	;
	goto L25
}
func F_numeric_poly_avg(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int64
	_ = v22
	var v26 int32
	_ = v26
	var v31 int64
	_ = v31
	var v32 int64
	_ = v32
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v39 int32
	_ = v39
	var v49 int64
	_ = v49
	var v57 int64
	_ = v57
	var v62 int64
	_ = v62
	var v63 int64
	_ = v63
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v96 int64
	_ = v96
	var v97 int64
	_ = v97
	var v98 int64
	_ = v98
	var v99 int64
	_ = v99
	var v104 int64
	_ = v104
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v114 int64
	_ = v114
	var v118 int64
	_ = v118
	var v125 int64
	_ = v125
	var v137 int32
	_ = v137
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v143 int64
	_ = v143
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v166 int64
	_ = v166
	var v170 int32
	_ = v170
	var v171 int32
	_ = v171
	var v173 int32
	_ = v173
	var v185 int64
	_ = v185
	var v191 int64
	_ = v191
	var v194 int32
	_ = v194
	var v196 int32
	_ = v196
	var v203 int64
	_ = v203
	var v208 int32
	_ = v208
	var v210 int64
	_ = v210
	var v213 int64
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v241 int32
	_ = v241
	var v245 int32
	_ = v245
	var v246 int32
	_ = v246
	var v248 int32
	_ = v248
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v266 int32
	_ = v266
	v14 = m.G0
	v16 = v14 - int32(80)
	m.G0 = v16
	v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v18 != 0 {
		v26 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
		v266 = int32(0)
		m.G0 = v16 + int32(80)
		return v266
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v19 == int32(0) {
			v26 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
			v266 = int32(0)
			m.G0 = v16 + int32(80)
			return v266
		} else {
			v22 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
			if v22 != int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = int64(0)
				v31 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
				v32 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
				v34 = F_palloc(m, int32(22))
				mBase = m.M
				v37 = m.ExcPending
				if v37 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+48)) = v34
					v39 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v34))) = uint16(v39)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v34 + int32(2)
					if v32 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = int64(16384)
						v49 = int64(0)
						v62 = v49 - (v32 + base.I64_extend_i32_u(base.B2i32(v31 != v49)))
						v63 = v49 - v31
						v66 = v39
						v68 = v34 + int32(22)
						v75 = v62
						v76 = v63
						for {
							v79 = int32(16)
							v80 = v16 + v79
							v83 = m.G0
							v85 = v83 - v79
							m.G0 = v85
							F___udivmodti4(m, v85, v76, v75, int64(10000), int64(0))
							mBase = m.M
							v89 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
							v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
							*(*int64)(unsafe.Add(mBase, uint32(v80))) = v90
							*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v89
							m.G0 = v85 + v79
							v96 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
							v97 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
							v98 = int64(55536)
							v99 = int64(0)
							v104 = int64(32)
							v107 = int64(base.Ui64(v96) >> (uint(v104) % 64))
							v110 = int64(4294967295)
							v113 = v96 & v110
							v114 = v98 * v113
							v118 = int64(base.Ui64(v114)>>(uint(v104)%64)) + v98*v107
							v125 = v113*v99 + v118&v110
							*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v96*v99 + v97*v98 + v99*v107 + int64(base.Ui64(v118)>>(uint(v104)%64)) + int64(base.Ui64(v125)>>(uint(v104)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v16))) = v114&v110 | v125<<(uint(v104)%64)
							v137 = v68 - int32(2)
							v138 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
							v139 = v138 + v76
							*(*uint16)(unsafe.Add(mBase, uint32(v137))) = uint16(v139)
							v143 = int64(0)
							v148 = v66 + int32(1)
							if v75 == v143 {
								v149 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v76))
							} else {
								v149 = base.B2i32(v75 != v143)
							}
							if v149 != 0 {
								v66 = v148
								v68 = v137
								v75 = v97
								v76 = v96
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v137
						v151 = v148
						v155 = v66
					} else {
						v57 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v16)+40)) = v57
						if v32|v31 == v57 {
							v151 = v39
							v155 = int32(0)
						} else {
							v62 = v32
							v63 = v31
							v66 = v39
							v68 = v34 + int32(22)
							v75 = v62
							v76 = v63
							for {
								v79 = int32(16)
								v80 = v16 + v79
								v83 = m.G0
								v85 = v83 - v79
								m.G0 = v85
								F___udivmodti4(m, v85, v76, v75, int64(10000), int64(0))
								mBase = m.M
								v89 = *(*int64)(unsafe.Add(mBase, uint32(v85)+8))
								v90 = *(*int64)(unsafe.Add(mBase, uint32(v85)))
								*(*int64)(unsafe.Add(mBase, uint32(v80))) = v90
								*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v89
								m.G0 = v85 + v79
								v96 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
								v97 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
								v98 = int64(55536)
								v99 = int64(0)
								v104 = int64(32)
								v107 = int64(base.Ui64(v96) >> (uint(v104) % 64))
								v110 = int64(4294967295)
								v113 = v96 & v110
								v114 = v98 * v113
								v118 = int64(base.Ui64(v114)>>(uint(v104)%64)) + v98*v107
								v125 = v113*v99 + v118&v110
								*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v96*v99 + v97*v98 + v99*v107 + int64(base.Ui64(v118)>>(uint(v104)%64)) + int64(base.Ui64(v125)>>(uint(v104)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v16))) = v114&v110 | v125<<(uint(v104)%64)
								v137 = v68 - int32(2)
								v138 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
								v139 = v138 + v76
								*(*uint16)(unsafe.Add(mBase, uint32(v137))) = uint16(v139)
								v143 = int64(0)
								v148 = v66 + int32(1)
								if v75 == v143 {
									v149 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v76))
								} else {
									v149 = base.B2i32(v75 != v143)
								}
								if v149 != 0 {
									v66 = v148
									v68 = v137
									v75 = v97
									v76 = v96
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+52)) = v137
							v151 = v148
							v155 = v66
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+36)) = v155
					*(*int32)(unsafe.Add(mBase, uint32(v16)+32)) = v151
					v166 = *(*int64)(unsafe.Add(mBase, uint32(v19)+8))
					*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = int64(0)
					v170 = F_palloc(m, int32(12))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v170
						v173 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v170))) = uint16(v173)
						*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v170 + int32(2)
						if v166 < int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = int64(16384)
							v191 = int64(0) - v166
							v194 = v173
							v196 = v170 + int32(12)
							v203 = v191
							for {
								v208 = v196 - int32(2)
								v210 = base.I64_div_u_s(v203, int64(10000))
								v213 = v210*int64(55536) + v203
								*(*uint16)(unsafe.Add(mBase, uint32(v208))) = uint16(v213)
								v216 = v194 + int32(1)
								if base.Ui64(int64(9999)) < base.Ui64(v203) {
									v194 = v216
									v196 = v208
									v203 = v210
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v208
							v220 = v216
							v224 = v194
						} else {
							v185 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v16)+64)) = v185
							if v166 == v185 {
								v220 = v173
								v224 = int32(0)
							} else {
								v191 = v166
								v194 = v173
								v196 = v170 + int32(12)
								v203 = v191
								for {
									v208 = v196 - int32(2)
									v210 = base.I64_div_u_s(v203, int64(10000))
									v213 = v210*int64(55536) + v203
									*(*uint16)(unsafe.Add(mBase, uint32(v208))) = uint16(v213)
									v216 = v194 + int32(1)
									if base.Ui64(int64(9999)) < base.Ui64(v203) {
										v194 = v216
										v196 = v208
										v203 = v210
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v208
								v220 = v216
								v224 = v194
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v16)+60)) = v224
						*(*int32)(unsafe.Add(mBase, uint32(v16)+56)) = v220
						v238 = F_make_result_opt_error(m, v16+int32(56), int32(0))
						mBase = m.M
						v239 = m.ExcPending
						if v239 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v170)
							mBase = m.M
							v241 = m.ExcPending
							if v241 != 0 {
								return int32(0)
							} else {
								v245 = F_make_result_opt_error(m, v16+int32(32), int32(0))
								mBase = m.M
								v246 = m.ExcPending
								if v246 != 0 {
									return int32(0)
								} else {
									F_pfree(m, v34)
									mBase = m.M
									v248 = m.ExcPending
									if v248 != 0 {
										return int32(0)
									} else {
										v251 = F_DirectFunctionCall2Coll(m, int32(1260), int32(0), v245, v238)
										mBase = m.M
										v252 = m.ExcPending
										if v252 != 0 {
											return int32(0)
										} else {
											v266 = v251
											m.G0 = v16 + int32(80)
											return v266
										}
									}
								}
							}
						}
					}
				}
			} else {
				v26 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v26)
				v266 = int32(0)
				m.G0 = v16 + int32(80)
				return v266
			}
		}
	}
}
func F_numeric_poly_stddev_internal(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v26 int64
	_ = v26
	var v30 int64
	_ = v30
	var v31 int64
	_ = v31
	var v33 int32
	_ = v33
	var v36 int32
	_ = v36
	var v38 int32
	_ = v38
	var v47 int64
	_ = v47
	var v55 int64
	_ = v55
	var v60 int64
	_ = v60
	var v61 int64
	_ = v61
	var v69 int32
	_ = v69
	var v71 int32
	_ = v71
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v98 int32
	_ = v98
	var v99 int64
	_ = v99
	var v100 int64
	_ = v100
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v107 int64
	_ = v107
	var v110 int64
	_ = v110
	var v113 int64
	_ = v113
	var v116 int64
	_ = v116
	var v117 int64
	_ = v117
	var v121 int64
	_ = v121
	var v128 int64
	_ = v128
	var v140 int32
	_ = v140
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v146 int64
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v177 int32
	_ = v177
	var v178 int64
	_ = v178
	var v179 int64
	_ = v179
	var v181 int32
	_ = v181
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v186 int32
	_ = v186
	var v196 int64
	_ = v196
	var v204 int64
	_ = v204
	var v211 int64
	_ = v211
	var v212 int64
	_ = v212
	var v220 int32
	_ = v220
	var v222 int32
	_ = v222
	var v227 int64
	_ = v227
	var v228 int64
	_ = v228
	var v231 int32
	_ = v231
	var v232 int32
	_ = v232
	var v235 int32
	_ = v235
	var v237 int32
	_ = v237
	var v241 int64
	_ = v241
	var v242 int64
	_ = v242
	var v248 int64
	_ = v248
	var v249 int64
	_ = v249
	var v250 int64
	_ = v250
	var v251 int64
	_ = v251
	var v256 int64
	_ = v256
	var v259 int64
	_ = v259
	var v262 int64
	_ = v262
	var v265 int64
	_ = v265
	var v266 int64
	_ = v266
	var v270 int64
	_ = v270
	var v277 int64
	_ = v277
	var v289 int32
	_ = v289
	var v290 int64
	_ = v290
	var v291 int64
	_ = v291
	var v295 int64
	_ = v295
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v310 int32
	_ = v310
	var v311 int32
	_ = v311
	var v326 int32
	_ = v326
	var v328 int32
	_ = v328
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v352 int32
	_ = v352
	var v354 int32
	_ = v354
	var v355 int32
	_ = v355
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v366 int32
	_ = v366
	v5 = int32(0)
	v17 = m.G0
	v19 = v17 - int32(208)
	m.G0 = v19
	base.MemoryFill(m, v19+int32(96), v5, int32(112))
	if l0 != 0 {
		v26 = *(*int64)(unsafe.Add(mBase, uint32(l0)+8))
		*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(0)
		*(*int64)(unsafe.Add(mBase, uint32(v19)+104)) = v26
		v30 = *(*int64)(unsafe.Add(mBase, uint32(l0)+16))
		v31 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v33 = F_palloc(m, int32(22))
		mBase = m.M
		v36 = m.ExcPending
		if v36 != 0 {
			return int32(0)
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v33
			v38 = int32(0)
			*(*uint16)(unsafe.Add(mBase, uint32(v33))) = uint16(v38)
			*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v33 + int32(2)
			if v31 < int64(0) {
				*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(16384)
				v47 = int64(0)
				v60 = v47 - v30
				v61 = v47 - (v31 + base.I64_extend_i32_u(base.B2i32(v30 != v47)))
				v69 = v33 + int32(22)
				v71 = v5
				v76 = v60
				v77 = v61
				for {
					v81 = v19 + int32(48)
					v84 = m.G0
					v85 = int32(16)
					v86 = v84 - v85
					m.G0 = v86
					F___udivmodti4(m, v86, v76, v77, int64(10000), int64(0))
					mBase = m.M
					v90 = *(*int64)(unsafe.Add(mBase, uint32(v86)+8))
					v91 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
					*(*int64)(unsafe.Add(mBase, uint32(v81))) = v91
					*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v90
					m.G0 = v86 + v85
					v98 = v19 + int32(32)
					v99 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
					v100 = *(*int64)(unsafe.Add(mBase, uint32(v19)+56))
					v101 = int64(55536)
					v102 = int64(0)
					v107 = int64(32)
					v110 = int64(base.Ui64(v99) >> (uint(v107) % 64))
					v113 = int64(4294967295)
					v116 = v99 & v113
					v117 = v101 * v116
					v121 = int64(base.Ui64(v117)>>(uint(v107)%64)) + v101*v110
					v128 = v116*v102 + v121&v113
					*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = v99*v102 + v100*v101 + v102*v110 + int64(base.Ui64(v121)>>(uint(v107)%64)) + int64(base.Ui64(v128)>>(uint(v107)%64))
					*(*int64)(unsafe.Add(mBase, uint32(v98))) = v117&v113 | v128<<(uint(v107)%64)
					v140 = v69 - int32(2)
					v141 = *(*int64)(unsafe.Add(mBase, uint32(v19)+32))
					v142 = v141 + v76
					*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v142)
					v146 = int64(0)
					v151 = v71 + int32(1)
					if v77 == v146 {
						v152 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v76))
					} else {
						v152 = base.B2i32(v77 != v146)
					}
					if v152 != 0 {
						v69 = v140
						v71 = v151
						v76 = v99
						v77 = v100
						continue
					} else {
						break
					}
					break
				}
				*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v140
				v161 = v151
				v162 = v71
			} else {
				v55 = int64(0)
				*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v55
				if v30|v31 != v55 {
					v60 = v30
					v61 = v31
					v69 = v33 + int32(22)
					v71 = v5
					v76 = v60
					v77 = v61
					for {
						v81 = v19 + int32(48)
						v84 = m.G0
						v85 = int32(16)
						v86 = v84 - v85
						m.G0 = v86
						F___udivmodti4(m, v86, v76, v77, int64(10000), int64(0))
						mBase = m.M
						v90 = *(*int64)(unsafe.Add(mBase, uint32(v86)+8))
						v91 = *(*int64)(unsafe.Add(mBase, uint32(v86)))
						*(*int64)(unsafe.Add(mBase, uint32(v81))) = v91
						*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v90
						m.G0 = v86 + v85
						v98 = v19 + int32(32)
						v99 = *(*int64)(unsafe.Add(mBase, uint32(v19)+48))
						v100 = *(*int64)(unsafe.Add(mBase, uint32(v19)+56))
						v101 = int64(55536)
						v102 = int64(0)
						v107 = int64(32)
						v110 = int64(base.Ui64(v99) >> (uint(v107) % 64))
						v113 = int64(4294967295)
						v116 = v99 & v113
						v117 = v101 * v116
						v121 = int64(base.Ui64(v117)>>(uint(v107)%64)) + v101*v110
						v128 = v116*v102 + v121&v113
						*(*int64)(unsafe.Add(mBase, uint32(v98)+8)) = v99*v102 + v100*v101 + v102*v110 + int64(base.Ui64(v121)>>(uint(v107)%64)) + int64(base.Ui64(v128)>>(uint(v107)%64))
						*(*int64)(unsafe.Add(mBase, uint32(v98))) = v117&v113 | v128<<(uint(v107)%64)
						v140 = v69 - int32(2)
						v141 = *(*int64)(unsafe.Add(mBase, uint32(v19)+32))
						v142 = v141 + v76
						*(*uint16)(unsafe.Add(mBase, uint32(v140))) = uint16(v142)
						v146 = int64(0)
						v151 = v71 + int32(1)
						if v77 == v146 {
							v152 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v76))
						} else {
							v152 = base.B2i32(v77 != v146)
						}
						if v152 != 0 {
							v69 = v140
							v71 = v151
							v76 = v99
							v77 = v100
							continue
						} else {
							break
						}
						break
					}
					*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v140
					v161 = v151
					v162 = v71
				} else {
					v161 = v5
					v162 = v5
				}
			}
			*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v162
			*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v161
			F_accum_sum_add(m, v19+int32(112), v19+int32(72))
			mBase = m.M
			v177 = m.ExcPending
			if v177 != 0 {
				return int32(0)
			} else {
				v178 = *(*int64)(unsafe.Add(mBase, uint32(l0)+32))
				v179 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
				F_pfree(m, v33)
				mBase = m.M
				v181 = m.ExcPending
				if v181 != 0 {
					return int32(0)
				} else {
					v183 = F_palloc(m, int32(22))
					mBase = m.M
					v184 = m.ExcPending
					if v184 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v19)+88)) = v183
						v186 = int32(0)
						*(*uint16)(unsafe.Add(mBase, uint32(v183))) = uint16(v186)
						*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v183 + int32(2)
						if v179 < int64(0) {
							*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = int64(16384)
							v196 = int64(0)
							v211 = v196 - v178
							v212 = v196 - (v179 + base.I64_extend_i32_u(base.B2i32(v178 != v196)))
							v220 = v183 + int32(22)
							v222 = v186
							v227 = v211
							v228 = v212
							for {
								v231 = int32(16)
								v232 = v19 + v231
								v235 = m.G0
								v237 = v235 - v231
								m.G0 = v237
								F___udivmodti4(m, v237, v227, v228, int64(10000), int64(0))
								mBase = m.M
								v241 = *(*int64)(unsafe.Add(mBase, uint32(v237)+8))
								v242 = *(*int64)(unsafe.Add(mBase, uint32(v237)))
								*(*int64)(unsafe.Add(mBase, uint32(v232))) = v242
								*(*int64)(unsafe.Add(mBase, uint32(v232)+8)) = v241
								m.G0 = v237 + v231
								v248 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
								v249 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
								v250 = int64(55536)
								v251 = int64(0)
								v256 = int64(32)
								v259 = int64(base.Ui64(v248) >> (uint(v256) % 64))
								v262 = int64(4294967295)
								v265 = v248 & v262
								v266 = v250 * v265
								v270 = int64(base.Ui64(v266)>>(uint(v256)%64)) + v250*v259
								v277 = v265*v251 + v270&v262
								*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v248*v251 + v249*v250 + v251*v259 + int64(base.Ui64(v270)>>(uint(v256)%64)) + int64(base.Ui64(v277)>>(uint(v256)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v19))) = v266&v262 | v277<<(uint(v256)%64)
								v289 = v220 - int32(2)
								v290 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
								v291 = v290 + v227
								*(*uint16)(unsafe.Add(mBase, uint32(v289))) = uint16(v291)
								v295 = int64(0)
								v300 = v222 + int32(1)
								if v228 == v295 {
									v301 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v227))
								} else {
									v301 = base.B2i32(v228 != v295)
								}
								if v301 != 0 {
									v220 = v289
									v222 = v300
									v227 = v248
									v228 = v249
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v289
							v310 = v300
							v311 = v222
						} else {
							v204 = int64(0)
							*(*int64)(unsafe.Add(mBase, uint32(v19)+80)) = v204
							if v178|v179 == v204 {
								v310 = v186
								v311 = int32(0)
							} else {
								v211 = v178
								v212 = v179
								v220 = v183 + int32(22)
								v222 = v186
								v227 = v211
								v228 = v212
								for {
									v231 = int32(16)
									v232 = v19 + v231
									v235 = m.G0
									v237 = v235 - v231
									m.G0 = v237
									F___udivmodti4(m, v237, v227, v228, int64(10000), int64(0))
									mBase = m.M
									v241 = *(*int64)(unsafe.Add(mBase, uint32(v237)+8))
									v242 = *(*int64)(unsafe.Add(mBase, uint32(v237)))
									*(*int64)(unsafe.Add(mBase, uint32(v232))) = v242
									*(*int64)(unsafe.Add(mBase, uint32(v232)+8)) = v241
									m.G0 = v237 + v231
									v248 = *(*int64)(unsafe.Add(mBase, uint32(v19)+16))
									v249 = *(*int64)(unsafe.Add(mBase, uint32(v19)+24))
									v250 = int64(55536)
									v251 = int64(0)
									v256 = int64(32)
									v259 = int64(base.Ui64(v248) >> (uint(v256) % 64))
									v262 = int64(4294967295)
									v265 = v248 & v262
									v266 = v250 * v265
									v270 = int64(base.Ui64(v266)>>(uint(v256)%64)) + v250*v259
									v277 = v265*v251 + v270&v262
									*(*int64)(unsafe.Add(mBase, uint32(v19)+8)) = v248*v251 + v249*v250 + v251*v259 + int64(base.Ui64(v270)>>(uint(v256)%64)) + int64(base.Ui64(v277)>>(uint(v256)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v19))) = v266&v262 | v277<<(uint(v256)%64)
									v289 = v220 - int32(2)
									v290 = *(*int64)(unsafe.Add(mBase, uint32(v19)))
									v291 = v290 + v227
									*(*uint16)(unsafe.Add(mBase, uint32(v289))) = uint16(v291)
									v295 = int64(0)
									v300 = v222 + int32(1)
									if v228 == v295 {
										v301 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v227))
									} else {
										v301 = base.B2i32(v228 != v295)
									}
									if v301 != 0 {
										v220 = v289
										v222 = v300
										v227 = v248
										v228 = v249
										continue
									} else {
										break
									}
									break
								}
								*(*int32)(unsafe.Add(mBase, uint32(v19)+92)) = v289
								v310 = v300
								v311 = v222
							}
						}
						*(*int32)(unsafe.Add(mBase, uint32(v19)+76)) = v311
						*(*int32)(unsafe.Add(mBase, uint32(v19)+72)) = v310
						F_accum_sum_add(m, v19+int32(140), v19+int32(72))
						mBase = m.M
						v326 = m.ExcPending
						if v326 != 0 {
							return int32(0)
						} else {
							F_pfree(m, v183)
							mBase = m.M
							v328 = m.ExcPending
							if v328 != 0 {
								return int32(0)
							} else {
								v347 = F_numeric_stddev_internal(m, v19+int32(96), l1, l2, l3)
								mBase = m.M
								v348 = m.ExcPending
								if v348 != 0 {
									return int32(0)
								} else {
									v349 = *(*int32)(unsafe.Add(mBase, uint32(v19)+112))
									if int32(0) < v349 {
										v352 = *(*int32)(unsafe.Add(mBase, uint32(v19)+132))
										F_pfree(m, v352)
										mBase = m.M
										v354 = m.ExcPending
										if v354 != 0 {
											return int32(0)
										} else {
											v355 = *(*int32)(unsafe.Add(mBase, uint32(v19)+136))
											F_pfree(m, v355)
											mBase = m.M
											v357 = m.ExcPending
											if v357 != 0 {
												return int32(0)
											} else {
												v358 = *(*int32)(unsafe.Add(mBase, uint32(v19)+140))
												if int32(0) < v358 {
													v361 = *(*int32)(unsafe.Add(mBase, uint32(v19)+160))
													F_pfree(m, v361)
													mBase = m.M
													v363 = m.ExcPending
													if v363 != 0 {
														return int32(0)
													} else {
														v364 = *(*int32)(unsafe.Add(mBase, uint32(v19)+164))
														F_pfree(m, v364)
														mBase = m.M
														v366 = m.ExcPending
														if v366 != 0 {
															return int32(0)
														} else {
															m.G0 = v19 + int32(208)
															return v347
														}
													}
												} else {
													m.G0 = v19 + int32(208)
													return v347
												}
											}
										}
									} else {
										v358 = *(*int32)(unsafe.Add(mBase, uint32(v19)+140))
										if int32(0) < v358 {
											v361 = *(*int32)(unsafe.Add(mBase, uint32(v19)+160))
											F_pfree(m, v361)
											mBase = m.M
											v363 = m.ExcPending
											if v363 != 0 {
												return int32(0)
											} else {
												v364 = *(*int32)(unsafe.Add(mBase, uint32(v19)+164))
												F_pfree(m, v364)
												mBase = m.M
												v366 = m.ExcPending
												if v366 != 0 {
													return int32(0)
												} else {
													m.G0 = v19 + int32(208)
													return v347
												}
											}
										} else {
											m.G0 = v19 + int32(208)
											return v347
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
		v347 = F_numeric_stddev_internal(m, v19+int32(96), l1, l2, l3)
		mBase = m.M
		v348 = m.ExcPending
		if v348 != 0 {
			return int32(0)
		} else {
			v349 = *(*int32)(unsafe.Add(mBase, uint32(v19)+112))
			if int32(0) < v349 {
				v352 = *(*int32)(unsafe.Add(mBase, uint32(v19)+132))
				F_pfree(m, v352)
				mBase = m.M
				v354 = m.ExcPending
				if v354 != 0 {
					return int32(0)
				} else {
					v355 = *(*int32)(unsafe.Add(mBase, uint32(v19)+136))
					F_pfree(m, v355)
					mBase = m.M
					v357 = m.ExcPending
					if v357 != 0 {
						return int32(0)
					} else {
						v358 = *(*int32)(unsafe.Add(mBase, uint32(v19)+140))
						if int32(0) < v358 {
							v361 = *(*int32)(unsafe.Add(mBase, uint32(v19)+160))
							F_pfree(m, v361)
							mBase = m.M
							v363 = m.ExcPending
							if v363 != 0 {
								return int32(0)
							} else {
								v364 = *(*int32)(unsafe.Add(mBase, uint32(v19)+164))
								F_pfree(m, v364)
								mBase = m.M
								v366 = m.ExcPending
								if v366 != 0 {
									return int32(0)
								} else {
									m.G0 = v19 + int32(208)
									return v347
								}
							}
						} else {
							m.G0 = v19 + int32(208)
							return v347
						}
					}
				}
			} else {
				v358 = *(*int32)(unsafe.Add(mBase, uint32(v19)+140))
				if int32(0) < v358 {
					v361 = *(*int32)(unsafe.Add(mBase, uint32(v19)+160))
					F_pfree(m, v361)
					mBase = m.M
					v363 = m.ExcPending
					if v363 != 0 {
						return int32(0)
					} else {
						v364 = *(*int32)(unsafe.Add(mBase, uint32(v19)+164))
						F_pfree(m, v364)
						mBase = m.M
						v366 = m.ExcPending
						if v366 != 0 {
							return int32(0)
						} else {
							m.G0 = v19 + int32(208)
							return v347
						}
					}
				} else {
					m.G0 = v19 + int32(208)
					return v347
				}
			}
		}
	}
}
func F_numeric_poly_stddev_pop(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(0)
	v4 = Fn13952(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_numeric_poly_stddev_samp(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13952(m, l0, int32(1), int32(0))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_numeric_sum(m *base.Module, l0 int32) int32 {
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
	var v15 int64
	_ = v15
	var v16 int64
	_ = v16
	var v17 int64
	_ = v17
	var v21 int64
	_ = v21
	var v28 int32
	_ = v28
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v39 int64
	_ = v39
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = int32(*(*uint8)(unsafe.Add(mBase, uint32(l0)+24)))
	if v11 != 0 {
		v28 = int32(1)
		*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
		v82 = int32(0)
		m.G0 = v9 + int32(32)
		return v82
	} else {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		if v12 == int32(0) {
			v28 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
			v82 = int32(0)
			m.G0 = v9 + int32(32)
			return v82
		} else {
			v15 = *(*int64)(unsafe.Add(mBase, uint32(v12)+96))
			v16 = *(*int64)(unsafe.Add(mBase, uint32(v12)+88))
			v17 = *(*int64)(unsafe.Add(mBase, uint32(v12)+8))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v12)+104))
			if v15+(v16+v17) != int64(0)-v21 {
				if int64(0) < v16 {
					v35 = F_make_result_opt_error(m, int32(_a_F_numeric_sum_0), int32(0))
					mBase = m.M
					v38 = m.ExcPending
					if v38 != 0 {
						return int32(0)
					} else {
						v82 = v35
						m.G0 = v9 + int32(32)
						return v82
					}
				} else {
					v39 = int64(0)
					if base.B2i32(v15 <= v39)|base.B2i32(v21 <= v39) == int32(0) {
						v48 = F_make_result_opt_error(m, int32(_a_F_numeric_sum_0), int32(0))
						mBase = m.M
						v49 = m.ExcPending
						if v49 != 0 {
							return int32(0)
						} else {
							v82 = v48
							m.G0 = v9 + int32(32)
							return v82
						}
					} else {
						if int64(0) < v15 {
							v54 = F_make_result_opt_error(m, int32(_a_F_numeric_sum_1), int32(0))
							mBase = m.M
							v55 = m.ExcPending
							if v55 != 0 {
								return int32(0)
							} else {
								v82 = v54
								m.G0 = v9 + int32(32)
								return v82
							}
						} else {
							if int64(0) < v21 {
								v60 = F_make_result_opt_error(m, int32(_a_F_numeric_sum_2), int32(0))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return int32(0)
								} else {
									v82 = v60
									m.G0 = v9 + int32(32)
									return v82
								}
							} else {
								v62 = int64(0)
								*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v9)+16)) = v62
								*(*int64)(unsafe.Add(mBase, uint32(v9)+8)) = v62
								v71 = v9 + int32(8)
								F_accum_sum_final(m, v12+int32(16), v71)
								mBase = m.M
								v73 = m.ExcPending
								if v73 != 0 {
									return int32(0)
								} else {
									v75 = F_make_result_opt_error(m, v71, int32(0))
									mBase = m.M
									v76 = m.ExcPending
									if v76 != 0 {
										return int32(0)
									} else {
										v77 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
										if v77 == int32(0) {
											v82 = v75
											m.G0 = v9 + int32(32)
											return v82
										} else {
											F_pfree(m, v77)
											mBase = m.M
											v81 = m.ExcPending
											if v81 != 0 {
												return int32(0)
											} else {
												v82 = v75
												m.G0 = v9 + int32(32)
												return v82
											}
										}
									}
								}
							}
						}
					}
				}
			} else {
				v28 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v28)
				v82 = int32(0)
				m.G0 = v9 + int32(32)
				return v82
			}
		}
	}
}
func F_numeric_uplus(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v5 int32
	_ = v5
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
	var v16 int32
	_ = v16
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int32(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v12 = F_palloc(m, int32(base.Ui32(v9)>>(uint(int32(2))%32)))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int32(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v16 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if v16 != 0 {
				base.MemoryCopy(m, v12, v5, v16)
			} else {
			}
			return v12
		}
	}
}
func F_numeric_var_pop(m *base.Module, l0 int32) int32 {
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v4 = Fn13953(m, l0, int32(0), int32(1))
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
func F_numeric_var_samp(m *base.Module, l0 int32) int32 {
	var v2 int32
	_ = v2
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = int32(1)
	v4 = Fn13953(m, l0, v2, v2)
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
