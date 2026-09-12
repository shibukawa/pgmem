package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timestamp_cmp_timestamptz_internal(m *base.Module, l0 int64, l1 int64) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v11 int64
	_ = v11
	var v14 int32
	_ = v14
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v27 int32
	_ = v27
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v11 = F_timestamp2timestamptz_opt_overflow(m, l0, v7+int32(12))
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		if l1 == int64(9223372036854775807) {
			v19 = int32(-1)
		} else {
			v19 = int32(1)
		}
		v20 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v20 {
			v33 = v19
		} else {
			if l1 == int64(-9223372036854775807-1) {
				v27 = int32(1)
			} else {
				v27 = int32(-1)
			}
			if v20 < int32(0) {
				v33 = v27
			} else {
				v33 = base.B2i32(l1 < v11) - base.B2i32(v11 < l1)
			}
		}
		m.G0 = v7 + int32(16)
		return v33
	}
}
func F_timestamp_eq_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v18 int64
	_ = v18
	var v24 int32
	_ = v24
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v6 == int32(-2147483648) {
		v18 = int64(-9223372036854775807 - 1)
		v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) == int32(0))
	} else {
		if v6 == int32(2147483647) {
			v18 = int64(9223372036854775807)
			v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) == int32(0))
		} else {
			if int32(106751982) < v6 {
				v24 = int32(0)
			} else {
				v18 = base.I64_extend_i32_s(v6) * int64(86400000000)
				v24 = base.B2i32(base.B2i32(v4 < v18)-base.B2i32(v18 < v4) == int32(0))
			}
		}
	}
	return v24
}
func F_timestamp_ge_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v6 int32
	_ = v6
	var v17 int64
	_ = v17
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	if v6 == int32(-2147483648) {
		v17 = int64(-9223372036854775807 - 1)
		return base.B2i32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4) <= int32(0))
	} else {
		if v6 == int32(2147483647) {
			v17 = int64(9223372036854775807)
			return base.B2i32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4) <= int32(0))
		} else {
			if int32(106751982) < v6 {
				return base.B2i32(v4 == int64(9223372036854775807))
			} else {
				v17 = base.I64_extend_i32_s(v6) * int64(86400000000)
				return base.B2i32(base.B2i32(v4 < v17)-base.B2i32(v17 < v4) <= int32(0))
			}
		}
	}
}
func F_timestamp_in(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v26 int32
	_ = v26
	var v33 int32
	_ = v33
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v105 int32
	_ = v105
	var v110 int32
	_ = v110
	var v116 int64
	_ = v116
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v139 int64
	_ = v139
	var v146 int64
	_ = v146
	var v157 int64
	_ = v157
	var v158 int64
	_ = v158
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v196 int32
	_ = v196
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v214 int32
	_ = v214
	var v218 int32
	_ = v218
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v236 int32
	_ = v236
	var v237 int64
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	v10 = m.G0
	v12 = v10 - int32(496)
	m.G0 = v12
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v26 = F_ParseDateTime(m, v16, v12+int32(48), int32(153), v12+int32(320), v12+int32(208), v12+int32(428))
	mBase = m.M
	if v26 == int32(0) {
		v33 = *(*int32)(unsafe.Add(mBase, uint32(v12)+428))
		v44 = F_DecodeDateTime(m, v12+int32(320), v12+int32(208), v33, v12+int32(432), v12+int32(440), v12+int32(484), v12+int32(436), v12+int32(40))
		mBase = m.M
		v47 = m.ExcPending
		if v47 != 0 {
			return int32(0)
		} else {
			if v44 == int32(0) {
				v60 = *(*int32)(unsafe.Add(mBase, uint32(v12)+432))
				switch v60 - int32(2) {
				case 0:
					v63 = int64(*(*int32)(unsafe.Add(mBase, uint32(v12)+484)))
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)+460))
					if v64 <= int32(-4713) {
						if v64 != int32(-4713) {
							v189 = int32(0)
							v190 = F_errsave_start(m, v15)
							mBase = m.M
							v191 = m.ExcPending
							if v191 != 0 {
								return int32(0)
							} else {
								if v190 == int32(0) {
									v240 = v189
									m.G0 = v12 + int32(496)
									return v240
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v196 = m.ExcPending
									if v196 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
										F_errmsg(m, int32(708718), v12+int32(16))
										mBase = m.M
										v202 = m.ExcPending
										if v202 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return int32(0)
											} else {
												v240 = v189
												m.G0 = v12 + int32(496)
												return v240
											}
										}
									}
								}
							}
						} else {
							v69 = *(*int32)(unsafe.Add(mBase, uint32(v12)+456))
							if int32(10) < v69 {
								v80 = v69
								v82 = v12 + int32(24)
								v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+452))
								v88 = base.B2i32(int32(2) < v80)
								if int32(2) < v80 {
									v89 = int32(4800)
								} else {
									v89 = int32(4799)
								}
								v90 = v89 + v64
								v95 = base.I32_div_s(v90, int32(4))
								v98 = base.I32_div_s(v90, int32(-100))
								v101 = base.I32_div_s(v90, int32(400))
								if int32(2) < v80 {
									v105 = int32(1)
								} else {
									v105 = int32(13)
								}
								v110 = base.I32_div_s((v105+v80)*int32(7834), int32(256))
								v116 = base.I64_extend_i32_s(v83 + v90*int32(365) + v95 + v98 + v101 + v110 - int32(32167) - int32(2451545))
								v125 = int64(32)
								v126 = int64(20)
								v128 = int64(base.Ui64(v116) >> (uint(v125) % 64))
								v131 = int64(4294967295)
								v132 = int64(500654080)
								v134 = v116 & v131
								v135 = v132 * v134
								v139 = int64(base.Ui64(v135)>>(uint(v125)%64)) + v132*v128
								v146 = v134*v126 + v139&v131
								*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v116*int64(0) + v116>>(uint(int64(63))%64)*int64(86400000000) + v126*v128 + int64(base.Ui64(v139)>>(uint(v125)%64)) + int64(base.Ui64(v146)>>(uint(v125)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v82))) = v135&v131 | v146<<(uint(v125)%64)
								v157 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
								v158 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
								if v157 != v158>>(uint(int64(63))%64) {
									v189 = int32(0)
									v190 = F_errsave_start(m, v15)
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return int32(0)
									} else {
										if v190 == int32(0) {
											v240 = v189
											m.G0 = v12 + int32(496)
											return v240
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
												F_errmsg(m, int32(708718), v12+int32(16))
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return int32(0)
													} else {
														v240 = v189
														m.G0 = v12 + int32(496)
														return v240
													}
												}
											}
										}
									}
								} else {
									v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+440))
									v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+444))
									v164 = *(*int32)(unsafe.Add(mBase, uint32(v12)+448))
									v165 = int32(60)
									v174 = base.I64_extend_i32_s(v162+(v163+v164*v165)*v165)*int64(1000000) + v63
									v175 = v158 + v174
									*(*int64)(unsafe.Add(mBase, uint32(v12)+488)) = v175
									if base.B2i32(v174 < int64(0))^base.B2i32(v175 < v158) != 0 {
										v189 = int32(0)
										v190 = F_errsave_start(m, v15)
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											if v190 == int32(0) {
												v240 = v189
												m.G0 = v12 + int32(496)
												return v240
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
													F_errmsg(m, int32(708718), v12+int32(16))
													mBase = m.M
													v202 = m.ExcPending
													if v202 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return int32(0)
														} else {
															v240 = v189
															m.G0 = v12 + int32(496)
															return v240
														}
													}
												}
											}
										}
									} else {
										if base.Ui64(int64(9011559254509551615)) < base.Ui64(v175-int64(9223371331200000000)) {
											F_AdjustTimestampForTypmod(m, v12+int32(488), v14, v15)
											mBase = m.M
											v236 = m.ExcPending
											if v236 != 0 {
												return int32(0)
											} else {
												v237 = *(*int64)(unsafe.Add(mBase, uint32(v12)+488))
												v238 = F_Int64GetDatum(m, v237)
												mBase = m.M
												v239 = m.ExcPending
												if v239 != 0 {
													return int32(0)
												} else {
													v240 = v238
													m.G0 = v12 + int32(496)
													return v240
												}
											}
										} else {
											v189 = int32(0)
											v190 = F_errsave_start(m, v15)
											mBase = m.M
											v191 = m.ExcPending
											if v191 != 0 {
												return int32(0)
											} else {
												if v190 == int32(0) {
													v240 = v189
													m.G0 = v12 + int32(496)
													return v240
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v196 = m.ExcPending
													if v196 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
														F_errmsg(m, int32(708718), v12+int32(16))
														mBase = m.M
														v202 = m.ExcPending
														if v202 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
															mBase = m.M
															v207 = m.ExcPending
															if v207 != 0 {
																return int32(0)
															} else {
																v240 = v189
																m.G0 = v12 + int32(496)
																return v240
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v189 = int32(0)
								v190 = F_errsave_start(m, v15)
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return int32(0)
								} else {
									if v190 == int32(0) {
										v240 = v189
										m.G0 = v12 + int32(496)
										return v240
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
											F_errmsg(m, int32(708718), v12+int32(16))
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return int32(0)
												} else {
													v240 = v189
													m.G0 = v12 + int32(496)
													return v240
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v64 <= int32(5874897) {
							v74 = *(*int32)(unsafe.Add(mBase, uint32(v12)+456))
							v80 = v74
							v82 = v12 + int32(24)
							v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+452))
							v88 = base.B2i32(int32(2) < v80)
							if int32(2) < v80 {
								v89 = int32(4800)
							} else {
								v89 = int32(4799)
							}
							v90 = v89 + v64
							v95 = base.I32_div_s(v90, int32(4))
							v98 = base.I32_div_s(v90, int32(-100))
							v101 = base.I32_div_s(v90, int32(400))
							if int32(2) < v80 {
								v105 = int32(1)
							} else {
								v105 = int32(13)
							}
							v110 = base.I32_div_s((v105+v80)*int32(7834), int32(256))
							v116 = base.I64_extend_i32_s(v83 + v90*int32(365) + v95 + v98 + v101 + v110 - int32(32167) - int32(2451545))
							v125 = int64(32)
							v126 = int64(20)
							v128 = int64(base.Ui64(v116) >> (uint(v125) % 64))
							v131 = int64(4294967295)
							v132 = int64(500654080)
							v134 = v116 & v131
							v135 = v132 * v134
							v139 = int64(base.Ui64(v135)>>(uint(v125)%64)) + v132*v128
							v146 = v134*v126 + v139&v131
							*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v116*int64(0) + v116>>(uint(int64(63))%64)*int64(86400000000) + v126*v128 + int64(base.Ui64(v139)>>(uint(v125)%64)) + int64(base.Ui64(v146)>>(uint(v125)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v82))) = v135&v131 | v146<<(uint(v125)%64)
							v157 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
							v158 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
							if v157 != v158>>(uint(int64(63))%64) {
								v189 = int32(0)
								v190 = F_errsave_start(m, v15)
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return int32(0)
								} else {
									if v190 == int32(0) {
										v240 = v189
										m.G0 = v12 + int32(496)
										return v240
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
											F_errmsg(m, int32(708718), v12+int32(16))
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return int32(0)
												} else {
													v240 = v189
													m.G0 = v12 + int32(496)
													return v240
												}
											}
										}
									}
								}
							} else {
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+440))
								v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+444))
								v164 = *(*int32)(unsafe.Add(mBase, uint32(v12)+448))
								v165 = int32(60)
								v174 = base.I64_extend_i32_s(v162+(v163+v164*v165)*v165)*int64(1000000) + v63
								v175 = v158 + v174
								*(*int64)(unsafe.Add(mBase, uint32(v12)+488)) = v175
								if base.B2i32(v174 < int64(0))^base.B2i32(v175 < v158) != 0 {
									v189 = int32(0)
									v190 = F_errsave_start(m, v15)
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return int32(0)
									} else {
										if v190 == int32(0) {
											v240 = v189
											m.G0 = v12 + int32(496)
											return v240
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
												F_errmsg(m, int32(708718), v12+int32(16))
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return int32(0)
													} else {
														v240 = v189
														m.G0 = v12 + int32(496)
														return v240
													}
												}
											}
										}
									}
								} else {
									if base.Ui64(int64(9011559254509551615)) < base.Ui64(v175-int64(9223371331200000000)) {
										F_AdjustTimestampForTypmod(m, v12+int32(488), v14, v15)
										mBase = m.M
										v236 = m.ExcPending
										if v236 != 0 {
											return int32(0)
										} else {
											v237 = *(*int64)(unsafe.Add(mBase, uint32(v12)+488))
											v238 = F_Int64GetDatum(m, v237)
											mBase = m.M
											v239 = m.ExcPending
											if v239 != 0 {
												return int32(0)
											} else {
												v240 = v238
												m.G0 = v12 + int32(496)
												return v240
											}
										}
									} else {
										v189 = int32(0)
										v190 = F_errsave_start(m, v15)
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											if v190 == int32(0) {
												v240 = v189
												m.G0 = v12 + int32(496)
												return v240
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
													F_errmsg(m, int32(708718), v12+int32(16))
													mBase = m.M
													v202 = m.ExcPending
													if v202 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return int32(0)
														} else {
															v240 = v189
															m.G0 = v12 + int32(496)
															return v240
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							if v64 != int32(5874898) {
								v189 = int32(0)
								v190 = F_errsave_start(m, v15)
								mBase = m.M
								v191 = m.ExcPending
								if v191 != 0 {
									return int32(0)
								} else {
									if v190 == int32(0) {
										v240 = v189
										m.G0 = v12 + int32(496)
										return v240
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
											F_errmsg(m, int32(708718), v12+int32(16))
											mBase = m.M
											v202 = m.ExcPending
											if v202 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return int32(0)
												} else {
													v240 = v189
													m.G0 = v12 + int32(496)
													return v240
												}
											}
										}
									}
								}
							} else {
								v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)+456))
								if int32(5) < v77 {
									v189 = int32(0)
									v190 = F_errsave_start(m, v15)
									mBase = m.M
									v191 = m.ExcPending
									if v191 != 0 {
										return int32(0)
									} else {
										if v190 == int32(0) {
											v240 = v189
											m.G0 = v12 + int32(496)
											return v240
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
												F_errmsg(m, int32(708718), v12+int32(16))
												mBase = m.M
												v202 = m.ExcPending
												if v202 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return int32(0)
													} else {
														v240 = v189
														m.G0 = v12 + int32(496)
														return v240
													}
												}
											}
										}
									}
								} else {
									v80 = v77
									v82 = v12 + int32(24)
									v83 = *(*int32)(unsafe.Add(mBase, uint32(v12)+452))
									v88 = base.B2i32(int32(2) < v80)
									if int32(2) < v80 {
										v89 = int32(4800)
									} else {
										v89 = int32(4799)
									}
									v90 = v89 + v64
									v95 = base.I32_div_s(v90, int32(4))
									v98 = base.I32_div_s(v90, int32(-100))
									v101 = base.I32_div_s(v90, int32(400))
									if int32(2) < v80 {
										v105 = int32(1)
									} else {
										v105 = int32(13)
									}
									v110 = base.I32_div_s((v105+v80)*int32(7834), int32(256))
									v116 = base.I64_extend_i32_s(v83 + v90*int32(365) + v95 + v98 + v101 + v110 - int32(32167) - int32(2451545))
									v125 = int64(32)
									v126 = int64(20)
									v128 = int64(base.Ui64(v116) >> (uint(v125) % 64))
									v131 = int64(4294967295)
									v132 = int64(500654080)
									v134 = v116 & v131
									v135 = v132 * v134
									v139 = int64(base.Ui64(v135)>>(uint(v125)%64)) + v132*v128
									v146 = v134*v126 + v139&v131
									*(*int64)(unsafe.Add(mBase, uint32(v82)+8)) = v116*int64(0) + v116>>(uint(int64(63))%64)*int64(86400000000) + v126*v128 + int64(base.Ui64(v139)>>(uint(v125)%64)) + int64(base.Ui64(v146)>>(uint(v125)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v82))) = v135&v131 | v146<<(uint(v125)%64)
									v157 = *(*int64)(unsafe.Add(mBase, uint32(v12)+32))
									v158 = *(*int64)(unsafe.Add(mBase, uint32(v12)+24))
									if v157 != v158>>(uint(int64(63))%64) {
										v189 = int32(0)
										v190 = F_errsave_start(m, v15)
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											if v190 == int32(0) {
												v240 = v189
												m.G0 = v12 + int32(496)
												return v240
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
													F_errmsg(m, int32(708718), v12+int32(16))
													mBase = m.M
													v202 = m.ExcPending
													if v202 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return int32(0)
														} else {
															v240 = v189
															m.G0 = v12 + int32(496)
															return v240
														}
													}
												}
											}
										}
									} else {
										v162 = *(*int32)(unsafe.Add(mBase, uint32(v12)+440))
										v163 = *(*int32)(unsafe.Add(mBase, uint32(v12)+444))
										v164 = *(*int32)(unsafe.Add(mBase, uint32(v12)+448))
										v165 = int32(60)
										v174 = base.I64_extend_i32_s(v162+(v163+v164*v165)*v165)*int64(1000000) + v63
										v175 = v158 + v174
										*(*int64)(unsafe.Add(mBase, uint32(v12)+488)) = v175
										if base.B2i32(v174 < int64(0))^base.B2i32(v175 < v158) != 0 {
											v189 = int32(0)
											v190 = F_errsave_start(m, v15)
											mBase = m.M
											v191 = m.ExcPending
											if v191 != 0 {
												return int32(0)
											} else {
												if v190 == int32(0) {
													v240 = v189
													m.G0 = v12 + int32(496)
													return v240
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v196 = m.ExcPending
													if v196 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
														F_errmsg(m, int32(708718), v12+int32(16))
														mBase = m.M
														v202 = m.ExcPending
														if v202 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
															mBase = m.M
															v207 = m.ExcPending
															if v207 != 0 {
																return int32(0)
															} else {
																v240 = v189
																m.G0 = v12 + int32(496)
																return v240
															}
														}
													}
												}
											}
										} else {
											if base.Ui64(int64(9011559254509551615)) < base.Ui64(v175-int64(9223371331200000000)) {
												F_AdjustTimestampForTypmod(m, v12+int32(488), v14, v15)
												mBase = m.M
												v236 = m.ExcPending
												if v236 != 0 {
													return int32(0)
												} else {
													v237 = *(*int64)(unsafe.Add(mBase, uint32(v12)+488))
													v238 = F_Int64GetDatum(m, v237)
													mBase = m.M
													v239 = m.ExcPending
													if v239 != 0 {
														return int32(0)
													} else {
														v240 = v238
														m.G0 = v12 + int32(496)
														return v240
													}
												}
											} else {
												v189 = int32(0)
												v190 = F_errsave_start(m, v15)
												mBase = m.M
												v191 = m.ExcPending
												if v191 != 0 {
													return int32(0)
												} else {
													if v190 == int32(0) {
														v240 = v189
														m.G0 = v12 + int32(496)
														return v240
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v196 = m.ExcPending
														if v196 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v16
															F_errmsg(m, int32(708718), v12+int32(16))
															mBase = m.M
															v202 = m.ExcPending
															if v202 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v15, int32(492756), int32(204), int32(277867))
																mBase = m.M
																v207 = m.ExcPending
																if v207 != 0 {
																	return int32(0)
																} else {
																	v240 = v189
																	m.G0 = v12 + int32(496)
																	return v240
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
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v212 = m.ExcPending
					if v212 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = v16
						v214 = *(*int32)(unsafe.Add(mBase, uint32(v12)+432))
						*(*int32)(unsafe.Add(mBase, uint32(v12))) = v214
						F_errmsg_internal(m, int32(683628), v12)
						mBase = m.M
						v218 = m.ExcPending
						if v218 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(492756), int32(221), int32(277867))
							mBase = m.M
							v223 = m.ExcPending
							if v223 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 7:
					v226 = int64(-9223372036854775807 - 1)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+488)) = v226
					F_AdjustTimestampForTypmod(m, v12+int32(488), v14, v15)
					mBase = m.M
					v236 = m.ExcPending
					if v236 != 0 {
						return int32(0)
					} else {
						v237 = *(*int64)(unsafe.Add(mBase, uint32(v12)+488))
						v238 = F_Int64GetDatum(m, v237)
						mBase = m.M
						v239 = m.ExcPending
						if v239 != 0 {
							return int32(0)
						} else {
							v240 = v238
							m.G0 = v12 + int32(496)
							return v240
						}
					}
				case 8:
					v226 = int64(9223372036854775807)
					*(*int64)(unsafe.Add(mBase, uint32(v12)+488)) = v226
					F_AdjustTimestampForTypmod(m, v12+int32(488), v14, v15)
					mBase = m.M
					v236 = m.ExcPending
					if v236 != 0 {
						return int32(0)
					} else {
						v237 = *(*int64)(unsafe.Add(mBase, uint32(v12)+488))
						v238 = F_Int64GetDatum(m, v237)
						mBase = m.M
						v239 = m.ExcPending
						if v239 != 0 {
							return int32(0)
						} else {
							v240 = v238
							m.G0 = v12 + int32(496)
							return v240
						}
					}
				case 9:
					v224 = F_SetEpochTimestamp(m)
					mBase = m.M
					v225 = m.ExcPending
					if v225 != 0 {
						return int32(0)
					} else {
						v226 = v224
						*(*int64)(unsafe.Add(mBase, uint32(v12)+488)) = v226
						F_AdjustTimestampForTypmod(m, v12+int32(488), v14, v15)
						mBase = m.M
						v236 = m.ExcPending
						if v236 != 0 {
							return int32(0)
						} else {
							v237 = *(*int64)(unsafe.Add(mBase, uint32(v12)+488))
							v238 = F_Int64GetDatum(m, v237)
							mBase = m.M
							v239 = m.ExcPending
							if v239 != 0 {
								return int32(0)
							} else {
								v240 = v238
								m.G0 = v12 + int32(496)
								return v240
							}
						}
					}
				}
			} else {
				v50 = v44
				F_DateTimeParseError(m, v50, v12+int32(40), v16, int32(236048), v15)
				mBase = m.M
				v55 = m.ExcPending
				if v55 != 0 {
					return int32(0)
				} else {
					v56 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
					v240 = int32(0)
					m.G0 = v12 + int32(496)
					return v240
				}
			}
		}
	} else {
		v50 = v26
		F_DateTimeParseError(m, v50, v12+int32(40), v16, int32(236048), v15)
		mBase = m.M
		v55 = m.ExcPending
		if v55 != 0 {
			return int32(0)
		} else {
			v56 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v56)
			v240 = int32(0)
			m.G0 = v12 + int32(496)
			return v240
		}
	}
}
func F_timestamp_lt_timestamptz(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int64
	_ = v12
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v29 int32
	_ = v29
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v21 {
			v29 = base.B2i32(v10 == int64(9223372036854775807))
		} else {
			if v21 < int32(0) {
				v29 = base.B2i32(v10 != int64(-9223372036854775807-1))
			} else {
				v29 = base.B2i32(v15 < v10)
			}
		}
		m.G0 = v7 + int32(16)
		return v29
	}
}
func F_timestamp_part_common(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
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
	var v36 int32
	_ = v36
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v64 int32
	_ = v64
	var v71 int32
	_ = v71
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v82 int32
	_ = v82
	var v84 int32
	_ = v84
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v108 int32
	_ = v108
	var v115 int32
	_ = v115
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v132 int32
	_ = v132
	var v136 int32
	_ = v136
	var v143 int32
	_ = v143
	var v145 int32
	_ = v145
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v169 int32
	_ = v169
	var v172 int32
	_ = v172
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v187 int32
	_ = v187
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v215 float64
	_ = v215
	var v216 int32
	_ = v216
	var v222 int32
	_ = v222
	var v226 int32
	_ = v226
	var v227 int32
	_ = v227
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v246 int64
	_ = v246
	var v254 int64
	_ = v254
	var v255 int64
	_ = v255
	var v258 int64
	_ = v258
	var v261 int32
	_ = v261
	var v273 int32
	_ = v273
	var v274 int32
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v282 int32
	_ = v282
	var v287 int32
	_ = v287
	var v290 int32
	_ = v290
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v295 int32
	_ = v295
	var v297 int32
	_ = v297
	var v303 int32
	_ = v303
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v310 int32
	_ = v310
	var v318 int32
	_ = v318
	var v322 int32
	_ = v322
	var v332 int32
	_ = v332
	var v340 int64
	_ = v340
	var v342 int64
	_ = v342
	var v348 int64
	_ = v348
	var v351 int64
	_ = v351
	var v353 int64
	_ = v353
	var v355 int64
	_ = v355
	var v358 int64
	_ = v358
	var v360 int64
	_ = v360
	var v361 int32
	_ = v361
	var v365 int64
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v379 float64
	_ = v379
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v393 int32
	_ = v393
	var v394 int32
	_ = v394
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v402 int64
	_ = v402
	var v403 int64
	_ = v403
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v419 int32
	_ = v419
	var v422 int32
	_ = v422
	var v425 int32
	_ = v425
	var v429 int32
	_ = v429
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v434 int32
	_ = v434
	var v436 int32
	_ = v436
	var v439 int32
	_ = v439
	var v443 int32
	_ = v443
	var v446 int32
	_ = v446
	var v447 int32
	_ = v447
	var v450 int32
	_ = v450
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v458 int32
	_ = v458
	var v465 int32
	_ = v465
	var v469 int32
	_ = v469
	var v474 int32
	_ = v474
	var v476 int32
	_ = v476
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v489 int32
	_ = v489
	var v495 int32
	_ = v495
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v511 int32
	_ = v511
	var v516 int32
	_ = v516
	var v519 int32
	_ = v519
	var v522 int32
	_ = v522
	var v526 int32
	_ = v526
	var v531 int32
	_ = v531
	var v534 int32
	_ = v534
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v539 int32
	_ = v539
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v542 int32
	_ = v542
	var v552 int32
	_ = v552
	var v553 int32
	_ = v553
	var v555 int32
	_ = v555
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v561 int32
	_ = v561
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v567 int32
	_ = v567
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v581 int32
	_ = v581
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v587 int32
	_ = v587
	var v588 int32
	_ = v588
	var v590 int32
	_ = v590
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v604 int32
	_ = v604
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v611 int32
	_ = v611
	var v612 int32
	_ = v612
	var v615 int32
	_ = v615
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v624 int32
	_ = v624
	var v630 int32
	_ = v630
	var v631 int32
	_ = v631
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v639 int32
	_ = v639
	var v644 int32
	_ = v644
	var v647 int32
	_ = v647
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v659 int32
	_ = v659
	var v665 int32
	_ = v665
	var v666 int32
	_ = v666
	var v671 int32
	_ = v671
	var v672 int64
	_ = v672
	var v674 int64
	_ = v674
	var v675 int32
	_ = v675
	var v678 int64
	_ = v678
	var v679 int32
	_ = v679
	var v680 int32
	_ = v680
	var v681 int32
	_ = v681
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v688 int32
	_ = v688
	var v693 int32
	_ = v693
	var v696 int32
	_ = v696
	var v699 int32
	_ = v699
	var v703 int32
	_ = v703
	var v708 int32
	_ = v708
	var v712 int32
	_ = v712
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v741 int32
	_ = v741
	var v752 int32
	_ = v752
	var v755 int32
	_ = v755
	var v757 int32
	_ = v757
	var v758 int32
	_ = v758
	var v765 int32
	_ = v765
	var v770 int32
	_ = v770
	var v771 int32
	_ = v771
	var v774 int64
	_ = v774
	var v775 int32
	_ = v775
	var v777 int64
	_ = v777
	var v781 int32
	_ = v781
	var v782 int32
	_ = v782
	var v785 int32
	_ = v785
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v788 int32
	_ = v788
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v793 int32
	_ = v793
	var v794 int32
	_ = v794
	var v796 int32
	_ = v796
	var v797 int32
	_ = v797
	var v799 int32
	_ = v799
	var v800 int32
	_ = v800
	var v801 int32
	_ = v801
	var v802 int32
	_ = v802
	var v809 float64
	_ = v809
	var v812 int32
	_ = v812
	var v813 int32
	_ = v813
	var v817 int32
	_ = v817
	var v820 int32
	_ = v820
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v830 int32
	_ = v830
	var v835 int32
	_ = v835
	var v839 int32
	_ = v839
	var v842 int32
	_ = v842
	var v844 int32
	_ = v844
	var v845 int32
	_ = v845
	var v850 int32
	_ = v850
	var v855 int32
	_ = v855
	var v862 int64
	_ = v862
	var v863 int32
	_ = v863
	var v864 int32
	_ = v864
	var v866 int32
	_ = v866
	var v867 int32
	_ = v867
	var v870 int32
	_ = v870
	var v883 int32
	_ = v883
	var v886 int32
	_ = v886
	var v890 int32
	_ = v890
	var v895 int32
	_ = v895
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
	v24 = int32(1)
	v25 = v18 + v24
	v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
	v30 = v28 & v24
	if v30 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v31 = v25
	goto L5
L4:
	;
	v31 = v18 + int32(4)
	goto L5
L5:
	;
	if v28 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v61 = F_downcase_truncate_identifier(m, v31, v59, int32(0))
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v34 = int32(4)
	v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
	if v36&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v49 = int32(1)
	if v30 != 0 {
		v59 = int32(base.Ui32(v28)>>(uint(v49)%32)) - v49
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v45 = v34
	goto L12
L11:
	;
	v45 = base.B2i32(v36 == int32(18)) << (uint(v34) % 32)
	goto L12
L12:
	;
	if v36 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v48 = v34
	goto L15
L14:
	;
	v48 = v45
	goto L15
L15:
	;
	v59 = v48
	goto L6
L16:
	;
	v53 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v59 = int32(base.Ui32(v53)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v64 = v15 + int32(92)
	v71 = *(*int32)(unsafe.Add(mBase, _consts[1248]))
	if v71 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v132 == int32(31) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1248])) = v115
	v122 = int32(*(*int8)(unsafe.Add(mBase, uint32(v115)+11)))
	v123 = *(*int32)(unsafe.Add(mBase, uint32(v115)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = v123
	v132 = v122
	goto L18
L20:
	;
	v73 = F_strncmp(m, v61, v71, int32(10))
	mBase = m.M
	if v73 == int32(0) {
		v115 = v71
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v76 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
	v82 = int32(1643168)
	v84 = int32(1644128)
	goto L24
L23:
	;
	goto L22
L24:
	;
	v91 = v82 + (v84-v82)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v92 = int32(*(*int8)(unsafe.Add(mBase, uint32(v91))))
	v93 = v76 - v92
	if v93 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v64))) = int32(0)
	v132 = int32(31)
	goto L18
L26:
	;
	v97 = F_strncmp(m, v61, v91, int32(10))
	mBase = m.M
	if v97 == int32(0) {
		v115 = v91
		goto L19
	} else {
		goto L29
	}
L27:
	;
	v100 = v93
	goto L28
L28:
	;
	v104 = base.B2i32(v100 < int32(0))
	if v100 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v100 = v97
	goto L28
L30:
	;
	v105 = v91 - int32(16)
	goto L32
L31:
	;
	v105 = v84
	goto L32
L32:
	;
	if v100 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v108 = v82
	goto L35
L34:
	;
	v108 = v91 + int32(16)
	goto L35
L35:
	;
	if base.Ui32(v108) <= base.Ui32(v105) {
		v82 = v108
		v84 = v105
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v136 = v15 + int32(92)
	v143 = *(*int32)(unsafe.Add(mBase, _consts[1249]))
	if v143 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v205 = v132
	goto L39
L39:
	;
	if base.Ui64(int64(1)) < base.Ui64(v23-int64(9223372036854775807)) {
		goto L61
	} else {
		goto L62
	}
L40:
	;
	v205 = v204
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1249])) = v187
	v194 = int32(*(*int8)(unsafe.Add(mBase, uint32(v187)+11)))
	v195 = *(*int32)(unsafe.Add(mBase, uint32(v187)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = v195
	v204 = v194
	goto L40
L42:
	;
	v145 = F_strncmp(m, v61, v143, int32(10))
	mBase = m.M
	if v145 == int32(0) {
		v187 = v143
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v148 = int32(*(*int8)(unsafe.Add(mBase, uint32(v61))))
	v154 = int32(1642016)
	v156 = int32(1643152)
	goto L46
L45:
	;
	goto L44
L46:
	;
	v163 = v154 + (v156-v154)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v164 = int32(*(*int8)(unsafe.Add(mBase, uint32(v163))))
	v165 = v148 - v164
	if v165 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v136))) = int32(0)
	v204 = int32(31)
	goto L40
L48:
	;
	v169 = F_strncmp(m, v61, v163, int32(10))
	mBase = m.M
	if v169 == int32(0) {
		v187 = v163
		goto L41
	} else {
		goto L51
	}
L49:
	;
	v172 = v165
	goto L50
L50:
	;
	v176 = base.B2i32(v172 < int32(0))
	if v172 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v172 = v169
	goto L50
L52:
	;
	v177 = v163 - int32(16)
	goto L54
L53:
	;
	v177 = v156
	goto L54
L54:
	;
	if v172 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v180 = v154
	goto L57
L56:
	;
	v180 = v163 + int32(16)
	goto L57
L57:
	;
	if base.Ui32(v180) <= base.Ui32(v177) {
		v154 = v180
		v156 = v177
		goto L46
	} else {
		goto L58
	}
L58:
	;
	goto L47
L59:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v883 = m.ExcPending
	if v883 != 0 {
		goto L1
	} else {
		goto L241
	}
L60:
	;
	m.G0 = v15 + int32(96)
	return v870
L61:
	;
	switch v205 {
	case 0:
		goto L80
	default:
		goto L79
	case 17:
		goto L81
	}
L62:
	;
	v210 = int32(0)
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	v215 = F_NonFiniteTimestampTzPart(m, v205, v211, v61, base.B2i32(v23 == int64(-9223372036854775807-1)), v210)
	mBase = m.M
	v216 = m.ExcPending
	if v216 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if base.F64_ne(v215, float64(0)) != 0 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	if l1 != 0 {
		goto L67
	} else {
		goto L68
	}
L65:
	;
	goto L66
L66:
	;
	v241 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v241)
	v870 = v210
	goto L60
L67:
	;
	if base.F64_lt(v215, float64(0)) != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v239 = F_Float8GetDatum(m, v215)
	mBase = m.M
	v240 = m.ExcPending
	if v240 != 0 {
		goto L1
	} else {
		goto L76
	}
L70:
	;
	v222 = int32(0)
	v226 = F_DirectFunctionCall3Coll(m, int32(408), v222, int32(11411), v222, int32(-1))
	mBase = m.M
	v227 = m.ExcPending
	if v227 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if base.F64_gt(v215, float64(0)) == int32(0) {
		goto L61
	} else {
		goto L74
	}
L73:
	;
	v870 = v226
	goto L60
L74:
	;
	v233 = int32(0)
	v237 = F_DirectFunctionCall3Coll(m, int32(408), v233, int32(11422), v233, int32(-1))
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v870 = v237
	goto L60
L76:
	;
	v870 = v239
	goto L60
L77:
	;
	if l1 != 0 {
		goto L236
	} else {
		goto L237
	}
L78:
	;
	v862 = base.I64_extend32_s(v365) + base.I64_extend32_s(v360)*int64(1000000)
	goto L77
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v839 = m.ExcPending
	if v839 != 0 {
		goto L1
	} else {
		goto L231
	}
L80:
	;
	v771 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	if v771 == int32(11) {
		goto L204
	} else {
		goto L205
	}
L81:
	;
	v246 = base.I64_div_s(v23, int64(86400000000))
	if base.Ui64(int64(172799999999)) <= base.Ui64(v23+int64(86399999999)) {
		goto L82
	} else {
		goto L83
	}
L82:
	;
	v254 = v246 * int64(-86400000000)
	goto L84
L83:
	;
	v254 = int64(0)
	goto L84
L84:
	;
	v255 = v254 + v23
	v258 = v255>>(uint(int64(63))%64) + v246
	if v258 <= int64(-2451546) {
		goto L59
	} else {
		goto L85
	}
L85:
	;
	v261 = base.I32_wrap_i64(v258)
	v273 = v261 + int32(2483589)
	v274 = int32(146097)
	v275 = base.I32_div_u_s(v273, v274)
	v276 = int32(3)
	v282 = int32(2)
	v287 = base.I32_div_u_s((v275*int32(1073595727)+v273)<<(uint(v282)%32)|v276, v274)
	v290 = v261 + int32(2451545) + v275*v276 + v287 + int32(32104)
	v291 = int32(1461)
	v292 = base.I32_div_u_s(v290, v291)
	v295 = v292*int32(-1461) + v290
	v297 = v295 << (uint(v282) % 32)
	if base.Ui32(v291) <= base.Ui32(v297) {
		goto L88
	} else {
		goto L89
	}
L86:
	;
	if v255 < int64(0) {
		goto L91
	} else {
		goto L92
	}
L87:
	;
	v310 = base.I32_div_u_s(v297, int32(1461))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(68)))) = v310 + v292<<(uint(int32(2))%32) - int32(4800)
	v318 = v308 + int32(123)
	v322 = int32(base.Ui32(v318*int32(2141)) >> (uint(int32(16)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v15+int32(60)))) = v318 - int32(base.Ui32(v322*int32(7834))>>(uint(int32(8))%32))
	v332 = base.I32_rem_u_s(v322+int32(10), int32(12))
	*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v332 + int32(1)
	goto L86
L88:
	;
	v303 = base.I32_rem_u_s(v295+int32(305), int32(365))
	v308 = v303
	goto L87
L89:
	;
	goto L90
L90:
	;
	v307 = base.I32_rem_u_s(v295+int32(306), int32(366))
	v308 = v307
	goto L87
L91:
	;
	v340 = v255 + int64(86400000000)
	goto L93
L92:
	;
	v340 = v255
	goto L93
L93:
	;
	v342 = base.I64_div_s(v340, int64(3600000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+56)) = uint32(v342)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = int64(4294967295)
	v348 = base.I64_extend32_s(v342)
	v351 = v348*int64(-3600000000) + v340
	v353 = base.I64_div_s(v351, int64(60000000))
	*(*uint32)(unsafe.Add(mBase, uint32(v15)+52)) = uint32(v353)
	v355 = base.I64_extend32_s(v353)
	v358 = v355*int64(-60000000) + v351
	v360 = base.I64_div_s(v358, int64(1000000))
	v361 = base.I32_wrap_i64(v360)
	*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v361
	v365 = v360*int64(4293967296) + v358
	v366 = base.I32_wrap_i64(v365)
	v367 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	switch v367 - int32(18) {
	case 0:
		goto L108
	case 1:
		v862 = v355
		goto L77
	case 2:
		goto L107
	case 3:
		goto L106
	case 4:
		goto L103
	case 5:
		goto L105
	case 6:
		goto L104
	case 7:
		goto L102
	case 8:
		goto L101
	case 9:
		goto L100
	case 10:
		goto L99
	case 11:
		goto L109
	case 12:
		goto L78
	case 13:
		goto L98
	case 14, 19:
		goto L96
	case 15:
		goto L95
	default:
		goto L94
	case 18:
		goto L97
	}
L94:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v752 = m.ExcPending
	if v752 != 0 {
		goto L1
	} else {
		goto L199
	}
L95:
	;
	v679 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	v680 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v681 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v686 = base.B2i32(int32(2) < v680)
	if int32(2) < v680 {
		goto L186
	} else {
		goto L187
	}
L96:
	;
	v630 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	v631 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v632 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v637 = base.B2i32(int32(2) < v631)
	if int32(2) < v631 {
		goto L169
	} else {
		goto L170
	}
L97:
	;
	v583 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	v584 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v585 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v587 = F_date2j(m, v583, v584, v585)
	mBase = m.M
	v588 = int32(1)
	v590 = F_date2j(m, v583, v588, int32(4))
	mBase = m.M
	v593 = F_j2day(m, v590-v588)
	mBase = m.M
	if v587 < v590-v593 {
		goto L159
	} else {
		goto L160
	}
L98:
	;
	v502 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	v503 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v504 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v509 = base.B2i32(int32(2) < v503)
	if int32(2) < v503 {
		goto L143
	} else {
		goto L144
	}
L99:
	;
	v489 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	if int32(0) < v489 {
		goto L139
	} else {
		goto L140
	}
L100:
	;
	v476 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	if int32(0) < v476 {
		goto L136
	} else {
		goto L137
	}
L101:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	if int32(0) <= v465 {
		goto L133
	} else {
		goto L134
	}
L102:
	;
	v458 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	if int32(0) < v458 {
		goto L130
	} else {
		goto L131
	}
L103:
	;
	v412 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
	v416 = F_date2j(m, v412, v413, v414)
	mBase = m.M
	v417 = int32(1)
	v419 = F_date2j(m, v412, v417, int32(4))
	mBase = m.M
	v422 = F_j2day(m, v419-v417)
	mBase = m.M
	if v416 < v419-v422 {
		goto L121
	} else {
		goto L122
	}
L104:
	;
	v404 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
	v405 = int32(1)
	v408 = base.I32_div_s(v404-v405, int32(3))
	v862 = base.I64_extend_i32_s(v408 + v405)
	goto L77
L105:
	;
	v403 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+64)))
	v862 = v403
	goto L77
L106:
	;
	v402 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+60)))
	v862 = v402
	goto L77
L107:
	;
	v862 = v348
	goto L77
L108:
	;
	if l1 != 0 {
		goto L115
	} else {
		goto L116
	}
L109:
	;
	if l1 != 0 {
		goto L110
	} else {
		goto L111
	}
L110:
	;
	v376 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v365)+base.I64_extend32_s(v360)*int64(1000000), int32(3))
	mBase = m.M
	v377 = m.ExcPending
	if v377 != 0 {
		goto L1
	} else {
		goto L113
	}
L111:
	;
	goto L112
L112:
	;
	v379 = float64(1000)
	v385 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v361), v379), base.F64_div(base.F64_convert_i32_s(v366), v379)))
	mBase = m.M
	v386 = m.ExcPending
	if v386 != 0 {
		goto L1
	} else {
		goto L114
	}
L113:
	;
	v870 = v376
	goto L60
L114:
	;
	v870 = v385
	goto L60
L115:
	;
	v393 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v365)+base.I64_extend32_s(v360)*int64(1000000), int32(6))
	mBase = m.M
	v394 = m.ExcPending
	if v394 != 0 {
		goto L1
	} else {
		goto L118
	}
L116:
	;
	goto L117
L117:
	;
	v400 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v366), float64(1e+06)), base.F64_convert_i32_s(v361)))
	mBase = m.M
	v401 = m.ExcPending
	if v401 != 0 {
		goto L1
	} else {
		goto L119
	}
L118:
	;
	v870 = v393
	goto L60
L119:
	;
	v870 = v400
	goto L60
L120:
	;
	v862 = base.I64_extend_i32_s(v454 + int32(1))
	goto L77
L121:
	;
	v425 = int32(1)
	v429 = F_date2j(m, v412-v425, v425, int32(4))
	mBase = m.M
	v432 = F_j2day(m, v429-v425)
	mBase = m.M
	v433 = v429
	v434 = v432
	goto L123
L122:
	;
	v433 = v419
	v434 = v422
	goto L123
L123:
	;
	v436 = v434 - v433 + v416
	if int32(357) <= v436 {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v439 = int32(1)
	v443 = F_date2j(m, v412+v439, v439, int32(4))
	mBase = m.M
	v446 = F_j2day(m, v443-v439)
	mBase = m.M
	v447 = v443 - v446
	if v416 < v447 {
		goto L127
	} else {
		goto L128
	}
L125:
	;
	v452 = v436
	goto L126
L126:
	;
	v454 = base.I32_div_s(v452, int32(7))
	goto L120
L127:
	;
	v450 = v436
	goto L129
L128:
	;
	v450 = v416 - v447
	goto L129
L129:
	;
	v452 = v450
	goto L126
L130:
	;
	v862 = base.I64_extend_i32_u(v458)
	goto L77
L131:
	;
	goto L132
L132:
	;
	v862 = base.I64_extend_i32_s(v458 - int32(1))
	goto L77
L133:
	;
	v469 = base.I32_div_u_s(v465, int32(10))
	v862 = base.I64_extend_i32_u(v469)
	goto L77
L134:
	;
	goto L135
L135:
	;
	v474 = base.I32_div_s(int32(9)-v465, int32(-10))
	v862 = base.I64_extend_i32_s(v474)
	goto L77
L136:
	;
	v482 = base.I32_div_s(v476+int32(99), int32(100))
	v862 = base.I64_extend_i32_s(v482)
	goto L77
L137:
	;
	goto L138
L138:
	;
	v487 = base.I32_div_s(int32(100)-v476, int32(-100))
	v862 = base.I64_extend_i32_s(v487)
	goto L77
L139:
	;
	v495 = base.I32_div_s(v489+int32(999), int32(1000))
	v862 = base.I64_extend_i32_s(v495)
	goto L77
L140:
	;
	goto L141
L141:
	;
	v500 = base.I32_div_s(int32(1000)-v489, int32(-1000))
	v862 = base.I64_extend_i32_s(v500)
	goto L77
L142:
	;
	if l1 != 0 {
		goto L149
	} else {
		goto L150
	}
L143:
	;
	v510 = int32(4800)
	goto L145
L144:
	;
	v510 = int32(4799)
	goto L145
L145:
	;
	v511 = v510 + v502
	v516 = base.I32_div_s(v511, int32(4))
	v519 = base.I32_div_s(v511, int32(-100))
	v522 = base.I32_div_s(v511, int32(400))
	if int32(2) < v503 {
		goto L146
	} else {
		goto L147
	}
L146:
	;
	v526 = int32(1)
	goto L148
L147:
	;
	v526 = int32(13)
	goto L148
L148:
	;
	v531 = base.I32_div_s((v526+v503)*int32(7834), int32(256))
	v534 = v504 + v511*int32(365) + v516 + v519 + v522 + v531 - int32(32167)
	goto L142
L149:
	;
	v536 = F_int64_to_numeric(m, base.I64_extend_i32_s(v534))
	mBase = m.M
	v537 = m.ExcPending
	if v537 != 0 {
		goto L1
	} else {
		goto L152
	}
L150:
	;
	goto L151
L151:
	;
	v566 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v567 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v568 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	v569 = int32(60)
	v581 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v366), float64(1e+06)), base.F64_convert_i32_s(v566+(v567+v568*v569)*v569)), float64(86400)), base.F64_convert_i32_s(v534)))
	mBase = m.M
	v582 = m.ExcPending
	if v582 != 0 {
		goto L1
	} else {
		goto L157
	}
L152:
	;
	v539 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
	v540 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
	v541 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
	v542 = int32(60)
	v552 = F_int64_to_numeric(m, base.I64_extend32_s(v365)+base.I64_extend_i32_s(v539+(v540+v541*v542)*v542)*int64(1000000))
	mBase = m.M
	v553 = m.ExcPending
	if v553 != 0 {
		goto L1
	} else {
		goto L153
	}
L153:
	;
	v555 = F_int64_to_numeric(m, int64(86400000000))
	mBase = m.M
	v556 = m.ExcPending
	if v556 != 0 {
		goto L1
	} else {
		goto L154
	}
L154:
	;
	v558 = F_numeric_div_opt_error(m, v552, v555, int32(0))
	mBase = m.M
	v559 = m.ExcPending
	if v559 != 0 {
		goto L1
	} else {
		goto L155
	}
L155:
	;
	v561 = F_numeric_add_opt_error(m, v536, v558, int32(0))
	mBase = m.M
	v562 = m.ExcPending
	if v562 != 0 {
		goto L1
	} else {
		goto L156
	}
L156:
	;
	v870 = v561
	goto L60
L157:
	;
	v870 = v581
	goto L60
L158:
	;
	v862 = base.I64_extend_i32_s(v624) - base.I64_extend_i32_u(base.B2i32(v624 <= int32(0)))
	goto L77
L159:
	;
	v596 = int32(1)
	v597 = v583 - v596
	v600 = F_date2j(m, v597, v596, int32(4))
	mBase = m.M
	v603 = F_j2day(m, v600-v596)
	mBase = m.M
	v604 = v597
	v605 = v600
	v606 = v603
	goto L161
L160:
	;
	v604 = v583
	v605 = v590
	v606 = v593
	goto L161
L161:
	;
	if int32(357) <= v587+v606-v605 {
		goto L162
	} else {
		goto L163
	}
L162:
	;
	v611 = int32(1)
	v612 = v604 + v611
	v615 = F_date2j(m, v612, v611, int32(4))
	mBase = m.M
	v618 = F_j2day(m, v615-v611)
	mBase = m.M
	if v587 < v615-v618 {
		goto L165
	} else {
		goto L166
	}
L163:
	;
	v624 = v604
	goto L164
L164:
	;
	goto L158
L165:
	;
	v621 = v604
	goto L167
L166:
	;
	v621 = v612
	goto L167
L167:
	;
	v624 = v621
	goto L164
L168:
	;
	v665 = int32(7)
	v666 = base.I32_rem_s(v632+v639*int32(365)+v644+v647+v650+v659-int32(32167)+int32(1), v665)
	if v666 < int32(0) {
		goto L176
	} else {
		goto L177
	}
L169:
	;
	v638 = int32(4800)
	goto L171
L170:
	;
	v638 = int32(4799)
	goto L171
L171:
	;
	v639 = v638 + v630
	v644 = base.I32_div_s(v639, int32(4))
	v647 = base.I32_div_s(v639, int32(-100))
	v650 = base.I32_div_s(v639, int32(400))
	if int32(2) < v631 {
		goto L172
	} else {
		goto L173
	}
L172:
	;
	v654 = int32(1)
	goto L174
L173:
	;
	v654 = int32(13)
	goto L174
L174:
	;
	v659 = base.I32_div_s((v654+v631)*int32(7834), int32(256))
	goto L168
L175:
	;
	v672 = base.I64_extend_i32_s(v671)
	if v671 != 0 {
		goto L179
	} else {
		goto L180
	}
L176:
	;
	v671 = v666 + v665
	goto L178
L177:
	;
	v671 = v666
	goto L178
L178:
	;
	goto L175
L179:
	;
	v674 = v672
	goto L181
L180:
	;
	v674 = int64(7)
	goto L181
L181:
	;
	v675 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
	if v675 == int32(37) {
		goto L182
	} else {
		goto L183
	}
L182:
	;
	v678 = v674
	goto L184
L183:
	;
	v678 = v672
	goto L184
L184:
	;
	v862 = v678
	goto L77
L185:
	;
	v712 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
	goto L194
L186:
	;
	v687 = int32(4800)
	goto L188
L187:
	;
	v687 = int32(4799)
	goto L188
L188:
	;
	v688 = v687 + v679
	v693 = base.I32_div_s(v688, int32(4))
	v696 = base.I32_div_s(v688, int32(-100))
	v699 = base.I32_div_s(v688, int32(400))
	if int32(2) < v680 {
		goto L189
	} else {
		goto L190
	}
L189:
	;
	v703 = int32(1)
	goto L191
L190:
	;
	v703 = int32(13)
	goto L191
L191:
	;
	v708 = base.I32_div_s((v703+v680)*int32(7834), int32(256))
	goto L185
L192:
	;
	v862 = base.I64_extend_i32_s(v681 + v688*int32(365) + v693 + v696 + v699 + v708 - int32(32167) - (int32(1) + v721*int32(365) + v726 + v729 + v732 + v741 - int32(32167)) + int32(1))
	goto L77
L194:
	;
	goto L195
L195:
	;
	v721 = int32(4799) + v712
	v726 = base.I32_div_s(v721, int32(4))
	v729 = base.I32_div_s(v721, int32(-100))
	v732 = base.I32_div_s(v721, int32(400))
	goto L197
L197:
	;
	goto L198
L198:
	;
	v741 = base.I32_div_s(int32(109676), int32(256))
	goto L192
L199:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v755 = m.ExcPending
	if v755 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v757 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v758 = m.ExcPending
	if v758 != 0 {
		goto L1
	} else {
		goto L201
	}
L201:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v757
	*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v61
	F_errmsg(m, int32(189145), v15+int32(16))
	mBase = m.M
	v765 = m.ExcPending
	if v765 != 0 {
		goto L1
	} else {
		goto L202
	}
L202:
	;
	F_errfinish(m, int32(492756), int32(5690), int32(244272))
	mBase = m.M
	v770 = m.ExcPending
	if v770 != 0 {
		goto L1
	} else {
		goto L203
	}
L203:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L204:
	;
	v774 = F_SetEpochTimestamp(m)
	mBase = m.M
	v775 = m.ExcPending
	if v775 != 0 {
		goto L1
	} else {
		goto L207
	}
L205:
	;
	goto L206
L206:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v817 = m.ExcPending
	if v817 != 0 {
		goto L1
	} else {
		goto L226
	}
L207:
	;
	v777 = v774 + int64(9223372036854775807)
	if l1 != 0 {
		goto L208
	} else {
		goto L209
	}
L208:
	;
	if v23 < v777 {
		goto L211
	} else {
		goto L212
	}
L209:
	;
	goto L210
L210:
	;
	if v23 < v777 {
		goto L222
	} else {
		goto L223
	}
L211:
	;
	v781 = F_int64_div_fast_to_numeric(m, v23-v774, int32(6))
	mBase = m.M
	v782 = m.ExcPending
	if v782 != 0 {
		goto L1
	} else {
		goto L214
	}
L212:
	;
	goto L213
L213:
	;
	v785 = F_int64_to_numeric(m, v23)
	mBase = m.M
	v786 = m.ExcPending
	if v786 != 0 {
		goto L1
	} else {
		goto L215
	}
L214:
	;
	v870 = v781
	goto L60
L215:
	;
	v787 = F_int64_to_numeric(m, v774)
	mBase = m.M
	v788 = m.ExcPending
	if v788 != 0 {
		goto L1
	} else {
		goto L216
	}
L216:
	;
	v790 = F_numeric_sub_opt_error(m, v785, v787, int32(0))
	mBase = m.M
	v791 = m.ExcPending
	if v791 != 0 {
		goto L1
	} else {
		goto L217
	}
L217:
	;
	v793 = F_int64_to_numeric(m, int64(1000000))
	mBase = m.M
	v794 = m.ExcPending
	if v794 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v796 = F_numeric_div_opt_error(m, v790, v793, int32(0))
	mBase = m.M
	v797 = m.ExcPending
	if v797 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	v799 = F_DirectFunctionCall2Coll(m, int32(1275), int32(0), v796, int32(6))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L220
	}
L220:
	;
	v801 = F_pg_detoast_datum(m, v799)
	mBase = m.M
	v802 = m.ExcPending
	if v802 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v870 = v801
	goto L60
L222:
	;
	v809 = base.F64_convert_i64_s(v23 - v774)
	goto L224
L223:
	;
	v809 = base.F64_sub(base.F64_convert_i64_s(v23), base.F64_convert_i64_s(v774))
	goto L224
L224:
	;
	v812 = F_Float8GetDatum(m, base.F64_div(v809, float64(1e+06)))
	mBase = m.M
	v813 = m.ExcPending
	if v813 != 0 {
		goto L1
	} else {
		goto L225
	}
L225:
	;
	v870 = v812
	goto L60
L226:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v820 = m.ExcPending
	if v820 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	v822 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v823 = m.ExcPending
	if v823 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v822
	*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v61
	F_errmsg(m, int32(189145), v15+int32(32))
	mBase = m.M
	v830 = m.ExcPending
	if v830 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	F_errfinish(m, int32(492756), int32(5737), int32(244272))
	mBase = m.M
	v835 = m.ExcPending
	if v835 != 0 {
		goto L1
	} else {
		goto L230
	}
L230:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L231:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v842 = m.ExcPending
	if v842 != 0 {
		goto L1
	} else {
		goto L232
	}
L232:
	;
	v844 = F_format_type_be(m, int32(1114))
	mBase = m.M
	v845 = m.ExcPending
	if v845 != 0 {
		goto L1
	} else {
		goto L233
	}
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v844
	*(*int32)(unsafe.Add(mBase, uint32(v15))) = v61
	F_errmsg(m, int32(189108), v15)
	mBase = m.M
	v850 = m.ExcPending
	if v850 != 0 {
		goto L1
	} else {
		goto L234
	}
L234:
	;
	F_errfinish(m, int32(492756), int32(5746), int32(244272))
	mBase = m.M
	v855 = m.ExcPending
	if v855 != 0 {
		goto L1
	} else {
		goto L235
	}
L235:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L236:
	;
	v863 = F_int64_to_numeric(m, v862)
	mBase = m.M
	v864 = m.ExcPending
	if v864 != 0 {
		goto L1
	} else {
		goto L239
	}
L237:
	;
	goto L238
L238:
	;
	v866 = F_Float8GetDatum(m, base.F64_convert_i64_s(v862))
	mBase = m.M
	v867 = m.ExcPending
	if v867 != 0 {
		goto L1
	} else {
		goto L240
	}
L239:
	;
	v870 = v863
	goto L60
L240:
	;
	v870 = v866
	goto L60
L241:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v886 = m.ExcPending
	if v886 != 0 {
		goto L1
	} else {
		goto L242
	}
L242:
	;
	F_errmsg(m, int32(400033), int32(0))
	mBase = m.M
	v890 = m.ExcPending
	if v890 != 0 {
		goto L1
	} else {
		goto L243
	}
L243:
	;
	F_errfinish(m, int32(492756), int32(5553), int32(244272))
	mBase = m.M
	v895 = m.ExcPending
	if v895 != 0 {
		goto L1
	} else {
		goto L244
	}
L244:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
