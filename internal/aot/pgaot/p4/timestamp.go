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
	var v10 int32
	_ = v10
	var v13 int64
	_ = v13
	var v15 int64
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_cmp_timestamptz_internal[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v10
	v13 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_cmp_timestamptz_internal[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v13
	v15 = F_timestamp2timestamptz_safe(m, l0, v7)
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v19 != int32(1) {
			v39 = base.B2i32(l1 < v15) - base.B2i32(v15 < l1)
		} else {
			if v15 != int64(-9223372036854775807-1) {
				if v15 != int64(9223372036854775807) {
					v39 = base.B2i32(l1 < v15) - base.B2i32(v15 < l1)
				} else {
					if l1 == int64(9223372036854775807) {
						v30 = int32(-1)
					} else {
						v30 = int32(1)
					}
					v39 = v30
				}
			} else {
				if l1 == int64(-9223372036854775807-1) {
					v35 = int32(1)
				} else {
					v35 = int32(-1)
				}
				v39 = v35
			}
		}
		m.G0 = v7 + int32(16)
		return v39
	}
}
func F_timestamp_eq_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_date_cmp_timestamp_internal(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(v4 == int32(0)))
	}
}
func F_timestamp_ge_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	var v4 int32
	_ = v4
	var v7 int32
	_ = v7
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v4 = F_date_cmp_timestamp_internal(m, v2, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int64(0)
	} else {
		return base.I64_extend_i32_u(base.B2i32(v4 <= int32(0)))
	}
}
func F_timestamp_in(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v102 int32
	_ = v102
	var v107 int32
	_ = v107
	var v113 int64
	_ = v113
	var v122 int64
	_ = v122
	var v123 int64
	_ = v123
	var v125 int64
	_ = v125
	var v128 int64
	_ = v128
	var v129 int64
	_ = v129
	var v131 int64
	_ = v131
	var v132 int64
	_ = v132
	var v136 int64
	_ = v136
	var v143 int64
	_ = v143
	var v154 int64
	_ = v154
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v171 int64
	_ = v171
	var v172 int64
	_ = v172
	var v186 int64
	_ = v186
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v199 int32
	_ = v199
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v211 int32
	_ = v211
	var v215 int32
	_ = v215
	var v220 int32
	_ = v220
	var v221 int64
	_ = v221
	var v222 int32
	_ = v222
	var v223 int64
	_ = v223
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v234 int64
	_ = v234
	var v237 int64
	_ = v237
	v12 = m.G0
	v14 = v12 - int32(512)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v23 = v14 + int32(336)
	v25 = v14 + int32(224)
	v28 = F_ParseDateTime(m, v18, v14-int32(-64), int32(153), v23, v25, v14+int32(444))
	mBase = m.M
	if v28 == int32(0) {
		v31 = *(*int32)(unsafe.Add(mBase, uint32(v14)+444))
		v42 = F_DecodeDateTime(m, v23, v25, v31, v14+int32(448), v14+int32(456), v14+int32(500), v14+int32(452), v14+int32(56))
		mBase = m.M
		v45 = m.ExcPending
		if v45 != 0 {
			return int64(0)
		} else {
			if v42 == int32(0) {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v14)+448))
				switch v57 - int32(2) {
				case 0:
					v60 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+500)))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v14)+476))
					if v61 <= int32(-4713) {
						if v61 != int32(-4713) {
							v186 = int64(0)
							v187 = F_errsave_start(m, v16)
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return int64(0)
							} else {
								if v187 == int32(0) {
									v237 = v186
									m.G0 = v14 + int32(512)
									return v237
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
										F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return int64(0)
										} else {
											F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return int64(0)
											} else {
												v237 = v186
												m.G0 = v14 + int32(512)
												return v237
											}
										}
									}
								}
							}
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v14)+472))
							if int32(10) < v66 {
								v77 = v66
								v79 = v14 + int32(32)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+468))
								v85 = base.B2i32(int32(2) < v77)
								if int32(2) < v77 {
									v86 = int32(_a_F_timestamp_in_3)
								} else {
									v86 = int32(_a_F_timestamp_in_4)
								}
								v87 = v86 + v61
								v92 = base.I32_div_s(v87, int32(4))
								v95 = base.I32_div_s(v87, int32(-100))
								v98 = base.I32_div_s(v87, int32(400))
								if int32(2) < v77 {
									v102 = int32(1)
								} else {
									v102 = int32(13)
								}
								v107 = base.I32_div_s((v102+v77)*int32(_a_F_timestamp_in_5), int32(256))
								v113 = base.I64_extend_i32_s(v80 + v87*int32(365) + v92 + v95 + v98 + v107 - int32(_a_F_timestamp_in_6) - int32(_a_F_timestamp_in_7))
								v122 = int64(32)
								v123 = int64(20)
								v125 = int64(base.Ui64(v113) >> (uint(v122) % 64))
								v128 = int64(4294967295)
								v129 = int64(500654080)
								v131 = v113 & v128
								v132 = v129 * v131
								v136 = int64(base.Ui64(v132)>>(uint(v122)%64)) + v129*v125
								v143 = v131*v123 + v136&v128
								*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v113*int64(0) + v113>>(uint(int64(63))%64)*int64(86400000000) + v123*v125 + int64(base.Ui64(v136)>>(uint(v122)%64)) + int64(base.Ui64(v143)>>(uint(v122)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v79))) = v132&v128 | v143<<(uint(v122)%64)
								v154 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
								v155 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
								if v154 != v155>>(uint(int64(63))%64) {
									v186 = int64(0)
									v187 = F_errsave_start(m, v16)
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int64(0)
									} else {
										if v187 == int32(0) {
											v237 = v186
											m.G0 = v14 + int32(512)
											return v237
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v193 = m.ExcPending
											if v193 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
												F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int64(0)
													} else {
														v237 = v186
														m.G0 = v14 + int32(512)
														return v237
													}
												}
											}
										}
									}
								} else {
									v159 = *(*int32)(unsafe.Add(mBase, uint32(v14)+456))
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+460))
									v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+464))
									v162 = int32(60)
									v171 = base.I64_extend_i32_s(v159+(v160+v161*v162)*v162)*int64(1000000) + v60
									v172 = v155 + v171
									*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v172
									if base.B2i32(v171 < int64(0))^base.B2i32(v172 < v155) != 0 {
										v186 = int64(0)
										v187 = F_errsave_start(m, v16)
										mBase = m.M
										v188 = m.ExcPending
										if v188 != 0 {
											return int64(0)
										} else {
											if v187 == int32(0) {
												v237 = v186
												m.G0 = v14 + int32(512)
												return v237
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v193 = m.ExcPending
												if v193 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
													F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int64(0)
														} else {
															v237 = v186
															m.G0 = v14 + int32(512)
															return v237
														}
													}
												}
											}
										}
									} else {
										if base.Ui64(int64(9011559254509551615)) < base.Ui64(v172-int64(9223371331200000000)) {
											v232 = F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
											mBase = m.M
											v233 = m.ExcPending
											if v233 != 0 {
												return int64(0)
											} else {
												v234 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
												v237 = v234
												m.G0 = v14 + int32(512)
												return v237
											}
										} else {
											v186 = int64(0)
											v187 = F_errsave_start(m, v16)
											mBase = m.M
											v188 = m.ExcPending
											if v188 != 0 {
												return int64(0)
											} else {
												if v187 == int32(0) {
													v237 = v186
													m.G0 = v14 + int32(512)
													return v237
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v193 = m.ExcPending
													if v193 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
														F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
														mBase = m.M
														v199 = m.ExcPending
														if v199 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
															mBase = m.M
															v204 = m.ExcPending
															if v204 != 0 {
																return int64(0)
															} else {
																v237 = v186
																m.G0 = v14 + int32(512)
																return v237
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v186 = int64(0)
								v187 = F_errsave_start(m, v16)
								mBase = m.M
								v188 = m.ExcPending
								if v188 != 0 {
									return int64(0)
								} else {
									if v187 == int32(0) {
										v237 = v186
										m.G0 = v14 + int32(512)
										return v237
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
											F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int64(0)
												} else {
													v237 = v186
													m.G0 = v14 + int32(512)
													return v237
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v61 <= int32(_a_F_timestamp_in_8) {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v14)+472))
							v77 = v71
							v79 = v14 + int32(32)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+468))
							v85 = base.B2i32(int32(2) < v77)
							if int32(2) < v77 {
								v86 = int32(_a_F_timestamp_in_3)
							} else {
								v86 = int32(_a_F_timestamp_in_4)
							}
							v87 = v86 + v61
							v92 = base.I32_div_s(v87, int32(4))
							v95 = base.I32_div_s(v87, int32(-100))
							v98 = base.I32_div_s(v87, int32(400))
							if int32(2) < v77 {
								v102 = int32(1)
							} else {
								v102 = int32(13)
							}
							v107 = base.I32_div_s((v102+v77)*int32(_a_F_timestamp_in_5), int32(256))
							v113 = base.I64_extend_i32_s(v80 + v87*int32(365) + v92 + v95 + v98 + v107 - int32(_a_F_timestamp_in_6) - int32(_a_F_timestamp_in_7))
							v122 = int64(32)
							v123 = int64(20)
							v125 = int64(base.Ui64(v113) >> (uint(v122) % 64))
							v128 = int64(4294967295)
							v129 = int64(500654080)
							v131 = v113 & v128
							v132 = v129 * v131
							v136 = int64(base.Ui64(v132)>>(uint(v122)%64)) + v129*v125
							v143 = v131*v123 + v136&v128
							*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v113*int64(0) + v113>>(uint(int64(63))%64)*int64(86400000000) + v123*v125 + int64(base.Ui64(v136)>>(uint(v122)%64)) + int64(base.Ui64(v143)>>(uint(v122)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v79))) = v132&v128 | v143<<(uint(v122)%64)
							v154 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
							v155 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
							if v154 != v155>>(uint(int64(63))%64) {
								v186 = int64(0)
								v187 = F_errsave_start(m, v16)
								mBase = m.M
								v188 = m.ExcPending
								if v188 != 0 {
									return int64(0)
								} else {
									if v187 == int32(0) {
										v237 = v186
										m.G0 = v14 + int32(512)
										return v237
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
											F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int64(0)
												} else {
													v237 = v186
													m.G0 = v14 + int32(512)
													return v237
												}
											}
										}
									}
								}
							} else {
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v14)+456))
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+460))
								v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+464))
								v162 = int32(60)
								v171 = base.I64_extend_i32_s(v159+(v160+v161*v162)*v162)*int64(1000000) + v60
								v172 = v155 + v171
								*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v172
								if base.B2i32(v171 < int64(0))^base.B2i32(v172 < v155) != 0 {
									v186 = int64(0)
									v187 = F_errsave_start(m, v16)
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int64(0)
									} else {
										if v187 == int32(0) {
											v237 = v186
											m.G0 = v14 + int32(512)
											return v237
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v193 = m.ExcPending
											if v193 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
												F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int64(0)
													} else {
														v237 = v186
														m.G0 = v14 + int32(512)
														return v237
													}
												}
											}
										}
									}
								} else {
									if base.Ui64(int64(9011559254509551615)) < base.Ui64(v172-int64(9223371331200000000)) {
										v232 = F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
										mBase = m.M
										v233 = m.ExcPending
										if v233 != 0 {
											return int64(0)
										} else {
											v234 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
											v237 = v234
											m.G0 = v14 + int32(512)
											return v237
										}
									} else {
										v186 = int64(0)
										v187 = F_errsave_start(m, v16)
										mBase = m.M
										v188 = m.ExcPending
										if v188 != 0 {
											return int64(0)
										} else {
											if v187 == int32(0) {
												v237 = v186
												m.G0 = v14 + int32(512)
												return v237
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v193 = m.ExcPending
												if v193 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
													F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int64(0)
														} else {
															v237 = v186
															m.G0 = v14 + int32(512)
															return v237
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							if v61 != int32(_a_F_timestamp_in_9) {
								v186 = int64(0)
								v187 = F_errsave_start(m, v16)
								mBase = m.M
								v188 = m.ExcPending
								if v188 != 0 {
									return int64(0)
								} else {
									if v187 == int32(0) {
										v237 = v186
										m.G0 = v14 + int32(512)
										return v237
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
											F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int64(0)
												} else {
													v237 = v186
													m.G0 = v14 + int32(512)
													return v237
												}
											}
										}
									}
								}
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v14)+472))
								if int32(5) < v74 {
									v186 = int64(0)
									v187 = F_errsave_start(m, v16)
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int64(0)
									} else {
										if v187 == int32(0) {
											v237 = v186
											m.G0 = v14 + int32(512)
											return v237
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v193 = m.ExcPending
											if v193 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
												F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int64(0)
													} else {
														v237 = v186
														m.G0 = v14 + int32(512)
														return v237
													}
												}
											}
										}
									}
								} else {
									v77 = v74
									v79 = v14 + int32(32)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v14)+468))
									v85 = base.B2i32(int32(2) < v77)
									if int32(2) < v77 {
										v86 = int32(_a_F_timestamp_in_3)
									} else {
										v86 = int32(_a_F_timestamp_in_4)
									}
									v87 = v86 + v61
									v92 = base.I32_div_s(v87, int32(4))
									v95 = base.I32_div_s(v87, int32(-100))
									v98 = base.I32_div_s(v87, int32(400))
									if int32(2) < v77 {
										v102 = int32(1)
									} else {
										v102 = int32(13)
									}
									v107 = base.I32_div_s((v102+v77)*int32(_a_F_timestamp_in_5), int32(256))
									v113 = base.I64_extend_i32_s(v80 + v87*int32(365) + v92 + v95 + v98 + v107 - int32(_a_F_timestamp_in_6) - int32(_a_F_timestamp_in_7))
									v122 = int64(32)
									v123 = int64(20)
									v125 = int64(base.Ui64(v113) >> (uint(v122) % 64))
									v128 = int64(4294967295)
									v129 = int64(500654080)
									v131 = v113 & v128
									v132 = v129 * v131
									v136 = int64(base.Ui64(v132)>>(uint(v122)%64)) + v129*v125
									v143 = v131*v123 + v136&v128
									*(*int64)(unsafe.Add(mBase, uint32(v79)+8)) = v113*int64(0) + v113>>(uint(int64(63))%64)*int64(86400000000) + v123*v125 + int64(base.Ui64(v136)>>(uint(v122)%64)) + int64(base.Ui64(v143)>>(uint(v122)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v79))) = v132&v128 | v143<<(uint(v122)%64)
									v154 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
									v155 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
									if v154 != v155>>(uint(int64(63))%64) {
										v186 = int64(0)
										v187 = F_errsave_start(m, v16)
										mBase = m.M
										v188 = m.ExcPending
										if v188 != 0 {
											return int64(0)
										} else {
											if v187 == int32(0) {
												v237 = v186
												m.G0 = v14 + int32(512)
												return v237
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v193 = m.ExcPending
												if v193 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
													F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int64(0)
														} else {
															v237 = v186
															m.G0 = v14 + int32(512)
															return v237
														}
													}
												}
											}
										}
									} else {
										v159 = *(*int32)(unsafe.Add(mBase, uint32(v14)+456))
										v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+460))
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+464))
										v162 = int32(60)
										v171 = base.I64_extend_i32_s(v159+(v160+v161*v162)*v162)*int64(1000000) + v60
										v172 = v155 + v171
										*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v172
										if base.B2i32(v171 < int64(0))^base.B2i32(v172 < v155) != 0 {
											v186 = int64(0)
											v187 = F_errsave_start(m, v16)
											mBase = m.M
											v188 = m.ExcPending
											if v188 != 0 {
												return int64(0)
											} else {
												if v187 == int32(0) {
													v237 = v186
													m.G0 = v14 + int32(512)
													return v237
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v193 = m.ExcPending
													if v193 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
														F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
														mBase = m.M
														v199 = m.ExcPending
														if v199 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
															mBase = m.M
															v204 = m.ExcPending
															if v204 != 0 {
																return int64(0)
															} else {
																v237 = v186
																m.G0 = v14 + int32(512)
																return v237
															}
														}
													}
												}
											}
										} else {
											if base.Ui64(int64(9011559254509551615)) < base.Ui64(v172-int64(9223371331200000000)) {
												v232 = F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
												mBase = m.M
												v233 = m.ExcPending
												if v233 != 0 {
													return int64(0)
												} else {
													v234 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
													v237 = v234
													m.G0 = v14 + int32(512)
													return v237
												}
											} else {
												v186 = int64(0)
												v187 = F_errsave_start(m, v16)
												mBase = m.M
												v188 = m.ExcPending
												if v188 != 0 {
													return int64(0)
												} else {
													if v187 == int32(0) {
														v237 = v186
														m.G0 = v14 + int32(512)
														return v237
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v193 = m.ExcPending
														if v193 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
															F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
															mBase = m.M
															v199 = m.ExcPending
															if v199 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(196), int32(_a_F_timestamp_in_2))
																mBase = m.M
																v204 = m.ExcPending
																if v204 != 0 {
																	return int64(0)
																} else {
																	v237 = v186
																	m.G0 = v14 + int32(512)
																	return v237
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
					v209 = m.ExcPending
					if v209 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v18
						v211 = *(*int32)(unsafe.Add(mBase, uint32(v14)+448))
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v211
						F_errmsg_internal(m, int32(_a_F_timestamp_in_10), v14)
						mBase = m.M
						v215 = m.ExcPending
						if v215 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamp_in_1), int32(213), int32(_a_F_timestamp_in_2))
							mBase = m.M
							v220 = m.ExcPending
							if v220 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 7:
					v223 = int64(-9223372036854775807 - 1)
					*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v223
					v232 = F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
					mBase = m.M
					v233 = m.ExcPending
					if v233 != 0 {
						return int64(0)
					} else {
						v234 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
						v237 = v234
						m.G0 = v14 + int32(512)
						return v237
					}
				case 8:
					v223 = int64(9223372036854775807)
					*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v223
					v232 = F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
					mBase = m.M
					v233 = m.ExcPending
					if v233 != 0 {
						return int64(0)
					} else {
						v234 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
						v237 = v234
						m.G0 = v14 + int32(512)
						return v237
					}
				case 9:
					v221 = F_SetEpochTimestamp(m)
					mBase = m.M
					v222 = m.ExcPending
					if v222 != 0 {
						return int64(0)
					} else {
						v223 = v221
						*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v223
						v232 = F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
						mBase = m.M
						v233 = m.ExcPending
						if v233 != 0 {
							return int64(0)
						} else {
							v234 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
							v237 = v234
							m.G0 = v14 + int32(512)
							return v237
						}
					}
				}
			} else {
				v48 = v42
				F_DateTimeParseError(m, v48, v14+int32(56), v18, int32(_a_F_timestamp_in_11), v16)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int64(0)
				} else {
					v54 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
					v237 = int64(0)
					m.G0 = v14 + int32(512)
					return v237
				}
			}
		}
	} else {
		v48 = v28
		F_DateTimeParseError(m, v48, v14+int32(56), v18, int32(_a_F_timestamp_in_11), v16)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int64(0)
		} else {
			v54 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
			v237 = int64(0)
			m.G0 = v14 + int32(512)
			return v237
		}
	}
}
func F_timestamp_lt_timestamptz(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int64
	_ = v10
	var v12 int32
	_ = v12
	var v15 int64
	_ = v15
	var v17 int64
	_ = v17
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v33 int32
	_ = v33
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamp_lt_timestamptz[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamp_lt_timestamptz[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v33 = base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v33 = base.B2i32(v17 < v9)
				} else {
					v33 = base.B2i32(v9 == int64(9223372036854775807))
				}
			} else {
				v33 = base.B2i32(v9 != int64(-9223372036854775807-1))
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v33)
	}
}
func F_timestamp_part_common(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int64
	_ = v29
	var v35 int32
	_ = v35
	var v38 int32
	_ = v38
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v56 int32
	_ = v56
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v65 int32
	_ = v65
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v77 int32
	_ = v77
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v92 int64
	_ = v92
	var v93 int32
	_ = v93
	var v103 int64
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v113 int64
	_ = v113
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v125 int64
	_ = v125
	var v128 int32
	_ = v128
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v143 int32
	_ = v143
	var v149 int32
	_ = v149
	var v154 int32
	_ = v154
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v162 int32
	_ = v162
	var v164 int32
	_ = v164
	var v170 int32
	_ = v170
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v177 int32
	_ = v177
	var v185 int32
	_ = v185
	var v189 int32
	_ = v189
	var v199 int32
	_ = v199
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v215 int64
	_ = v215
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v222 int64
	_ = v222
	var v225 int64
	_ = v225
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v232 int64
	_ = v232
	var v233 int32
	_ = v233
	var v234 int32
	_ = v234
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	var v247 float64
	_ = v247
	var v260 int32
	_ = v260
	var v261 int32
	_ = v261
	var v269 int64
	_ = v269
	var v270 int64
	_ = v270
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v275 int32
	_ = v275
	var v279 int32
	_ = v279
	var v280 int32
	_ = v280
	var v281 int32
	_ = v281
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v286 int32
	_ = v286
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v300 int32
	_ = v300
	var v301 int32
	_ = v301
	var v303 int32
	_ = v303
	var v306 int32
	_ = v306
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v319 int32
	_ = v319
	var v321 int32
	_ = v321
	var v325 int32
	_ = v325
	var v332 int32
	_ = v332
	var v336 int32
	_ = v336
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v349 int32
	_ = v349
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v362 int32
	_ = v362
	var v367 int32
	_ = v367
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v378 int32
	_ = v378
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v406 int32
	_ = v406
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v419 int32
	_ = v419
	var v420 int32
	_ = v420
	var v422 int32
	_ = v422
	var v423 int32
	_ = v423
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v428 int32
	_ = v428
	var v429 int32
	_ = v429
	var v434 int32
	_ = v434
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v437 int32
	_ = v437
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v485 int32
	_ = v485
	var v488 int32
	_ = v488
	var v491 int32
	_ = v491
	var v497 int32
	_ = v497
	var v498 int32
	_ = v498
	var v499 int32
	_ = v499
	var v504 int32
	_ = v504
	var v505 int32
	_ = v505
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v514 int32
	_ = v514
	var v517 int32
	_ = v517
	var v521 int32
	_ = v521
	var v526 int32
	_ = v526
	var v532 int32
	_ = v532
	var v533 int32
	_ = v533
	var v538 int32
	_ = v538
	var v539 int64
	_ = v539
	var v541 int64
	_ = v541
	var v542 int32
	_ = v542
	var v545 int64
	_ = v545
	var v546 int32
	_ = v546
	var v547 int32
	_ = v547
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v555 int32
	_ = v555
	var v560 int32
	_ = v560
	var v563 int32
	_ = v563
	var v566 int32
	_ = v566
	var v570 int32
	_ = v570
	var v575 int32
	_ = v575
	var v579 int32
	_ = v579
	var v588 int32
	_ = v588
	var v593 int32
	_ = v593
	var v596 int32
	_ = v596
	var v599 int32
	_ = v599
	var v608 int32
	_ = v608
	var v619 int32
	_ = v619
	var v622 int32
	_ = v622
	var v624 int32
	_ = v624
	var v625 int32
	_ = v625
	var v632 int32
	_ = v632
	var v637 int32
	_ = v637
	var v638 int32
	_ = v638
	var v641 int64
	_ = v641
	var v642 int32
	_ = v642
	var v644 int64
	_ = v644
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v653 int32
	_ = v653
	var v654 int32
	_ = v654
	var v655 int32
	_ = v655
	var v656 int32
	_ = v656
	var v658 int32
	_ = v658
	var v659 int32
	_ = v659
	var v661 int32
	_ = v661
	var v662 int32
	_ = v662
	var v664 int32
	_ = v664
	var v665 int32
	_ = v665
	var v668 int64
	_ = v668
	var v669 int32
	_ = v669
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v680 float64
	_ = v680
	var v687 int32
	_ = v687
	var v690 int32
	_ = v690
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v700 int32
	_ = v700
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v712 int32
	_ = v712
	var v714 int32
	_ = v714
	var v715 int32
	_ = v715
	var v720 int32
	_ = v720
	var v725 int32
	_ = v725
	var v732 int64
	_ = v732
	var v733 int32
	_ = v733
	var v734 int32
	_ = v734
	var v746 int64
	_ = v746
	var v754 int32
	_ = v754
	var v757 int32
	_ = v757
	var v761 int32
	_ = v761
	var v766 int32
	_ = v766
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int64(0)
	} else {
		v22 = int32(1)
		v24 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		v26 = v24 & v22
		if v26 != 0 {
			v27 = v22
		} else {
			v27 = int32(4)
		}
		v29 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if v24 == int32(1) {
			v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18)+1)))
			if v35 == int32(18) {
				v38 = int32(16)
			} else {
				v38 = int32(0)
			}
			if base.Ui32((v35-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v45 = int32(4)
			} else {
				v45 = v38
			}
			v56 = v45
		} else {
			v46 = int32(1)
			if v26 != 0 {
				v56 = int32(base.Ui32(v24)>>(uint(v46)%32)) - v46
			} else {
				v50 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v56 = int32(base.Ui32(v50)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v58 = F_downcase_truncate_identifier(m, v18+v27, v56, int32(0))
		mBase = m.M
		v59 = m.ExcPending
		if v59 != 0 {
			return int64(0)
		} else {
			v61 = v15 + int32(92)
			v65 = Fn14210(m, v58, v61, int32(_a_F_timestamp_part_common_0), int32(_a_F_timestamp_part_common_1), int32(_a_F_timestamp_part_common_2))
			mBase = m.M
			if v65 == int32(31) {
				v71 = Fn14210(m, v58, v61, int32(_a_F_timestamp_part_common_3), int32(_a_F_timestamp_part_common_4), int32(_a_F_timestamp_part_common_5))
				mBase = m.M
				v72 = v71
			} else {
				v72 = v65
			}
			if base.Ui64(int64(1)) < base.Ui64(v29-int64(9223372036854775807)) {
				if v72 != 0 {
					if v72 != int32(17) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v709 = m.ExcPending
						if v709 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v712 = m.ExcPending
							if v712 != 0 {
								return int64(0)
							} else {
								v714 = F_format_type_be(m, int32(1114))
								mBase = m.M
								v715 = m.ExcPending
								if v715 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v714
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = v58
									F_errmsg(m, int32(_a_F_timestamp_part_common_6), v15)
									mBase = m.M
									v720 = m.ExcPending
									if v720 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_8), int32(_a_F_timestamp_part_common_9))
										mBase = m.M
										v725 = m.ExcPending
										if v725 != 0 {
											return int64(0)
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
						v113 = base.I64_div_s(v29, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v29+int64(86399999999)) {
							v121 = v113 * int64(-86400000000)
						} else {
							v121 = int64(0)
						}
						v122 = v121 + v29
						v125 = v122>>(uint(int64(63))%64) + v113
						if v125 <= int64(-2451546) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v754 = m.ExcPending
							if v754 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v757 = m.ExcPending
								if v757 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamp_part_common_10), int32(0))
									mBase = m.M
									v761 = m.ExcPending
									if v761 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_11), int32(_a_F_timestamp_part_common_9))
										mBase = m.M
										v766 = m.ExcPending
										if v766 != 0 {
											return int64(0)
										} else {
											base.Wasm_trap_unreachable()
											for {
											}
										}
									}
								}
							}
						} else {
							v128 = base.I32_wrap_i64(v125)
							v140 = v128 + int32(_a_F_timestamp_part_common_12)
							v141 = int32(_a_F_timestamp_part_common_13)
							v142 = base.I32_div_u_s(v140, v141)
							v143 = int32(3)
							v149 = int32(2)
							v154 = base.I32_div_u_s((v142*int32(1073595727)+v140)<<(uint(v149)%32)|v143, v141)
							v157 = v128 + int32(_a_F_timestamp_part_common_14) + v142*v143 + v154 + int32(_a_F_timestamp_part_common_15)
							v158 = int32(1461)
							v159 = base.I32_div_u_s(v157, v158)
							v162 = v159*int32(-1461) + v157
							v164 = v162 << (uint(v149) % 32)
							if base.Ui32(v158) <= base.Ui32(v164) {
								v170 = base.I32_rem_u_s(v162+int32(305), int32(365))
								v175 = v170
							} else {
								v174 = base.I32_rem_u_s(v162+int32(306), int32(366))
								v175 = v174
							}
							v177 = base.I32_div_u_s(v164, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v15+int32(68)))) = v177 + v159<<(uint(int32(2))%32) - int32(_a_F_timestamp_part_common_16)
							v185 = v175 + int32(123)
							v189 = int32(base.Ui32(v185*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v15+int32(60)))) = v185 - int32(base.Ui32(v189*int32(_a_F_timestamp_part_common_17))>>(uint(int32(8))%32))
							v199 = base.I32_rem_u_s(v189+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v199 + int32(1)
							if v122 < int64(0) {
								v207 = v122 + int64(86400000000)
							} else {
								v207 = v122
							}
							v209 = base.I64_div_s(v207, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+56)) = uint32(v209)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = int64(4294967295)
							v215 = base.I64_extend32_s(v209)
							v218 = v215*int64(-3600000000) + v207
							v220 = base.I64_div_s(v218, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+52)) = uint32(v220)
							v222 = base.I64_extend32_s(v220)
							v225 = v218 + v222*int64(-60000000)
							v227 = base.I64_div_s(v225, int64(1000000))
							v228 = base.I32_wrap_i64(v227)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v228
							v232 = v227*int64(4293967296) + v225
							v233 = base.I32_wrap_i64(v232)
							v234 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
							switch v234 - int32(18) {
							case 0:
								if l1 != 0 {
									v260 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v232)+base.I64_extend32_s(v227)*int64(1000000), int32(6))
									mBase = m.M
									v261 = m.ExcPending
									if v261 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v260)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_convert_i32_s(v233), float64(1e+06)), base.F64_convert_i32_s(v228)))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 1:
								v732 = v222
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 2:
								v732 = v215
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 3:
								v269 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+60)))
								v732 = v269
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 4:
								v279 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v280 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v281 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v283 = F_date2j(m, v279, v280, v281)
								mBase = m.M
								v284 = int32(1)
								v286 = F_date2j(m, v279, v284, int32(4))
								mBase = m.M
								v289 = F_j2day(m, v286-v284)
								mBase = m.M
								if v283 < v286-v289 {
									v292 = int32(1)
									v296 = F_date2j(m, v279-v292, v292, int32(4))
									mBase = m.M
									v299 = F_j2day(m, v296-v292)
									mBase = m.M
									v300 = v296
									v301 = v299
								} else {
									v300 = v286
									v301 = v289
								}
								v303 = v301 - v300 + v283
								if int32(357) <= v303 {
									v306 = int32(1)
									v310 = F_date2j(m, v279+v306, v306, int32(4))
									mBase = m.M
									v313 = F_j2day(m, v310-v306)
									mBase = m.M
									v314 = v310 - v313
									if v283 < v314 {
										v317 = v303
									} else {
										v317 = v283 - v314
									}
									v319 = v317
								} else {
									v319 = v303
								}
								v321 = base.I32_div_s(v319, int32(7))
								v732 = base.I64_extend_i32_s(v321 + int32(1))
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 5:
								v270 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+64)))
								v732 = v270
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 6:
								v271 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v272 = int32(1)
								v275 = base.I32_div_s(v271-v272, int32(3))
								v732 = base.I64_extend_i32_s(v275 + v272)
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 7:
								v325 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								if int32(0) < v325 {
									v732 = base.I64_extend_i32_u(v325)
								} else {
									v732 = base.I64_extend_i32_s(v325 - int32(1))
								}
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 8:
								v332 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								if int32(0) <= v332 {
									v336 = base.I32_div_u_s(v332, int32(10))
									v732 = base.I64_extend_i32_u(v336)
								} else {
									v341 = base.I32_div_s(int32(9)-v332, int32(-10))
									v732 = base.I64_extend_i32_s(v341)
								}
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 9:
								v343 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								if int32(0) < v343 {
									v349 = base.I32_div_s(v343+int32(99), int32(100))
									v732 = base.I64_extend_i32_s(v349)
								} else {
									v354 = base.I32_div_s(int32(100)-v343, int32(-100))
									v732 = base.I64_extend_i32_s(v354)
								}
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 10:
								v356 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								if int32(0) < v356 {
									v362 = base.I32_div_s(v356+int32(999), int32(1000))
									v732 = base.I64_extend_i32_s(v362)
								} else {
									v367 = base.I32_div_s(int32(1000)-v356, int32(-1000))
									v732 = base.I64_extend_i32_s(v367)
								}
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 11:
								if l1 != 0 {
									v243 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v232)+base.I64_extend32_s(v227)*int64(1000000), int32(3))
									mBase = m.M
									v244 = m.ExcPending
									if v244 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v243)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v247 = float64(1000)
									v746 = base.I64_reinterpret_f64(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v228), v247), base.F64_div(base.F64_convert_i32_s(v233), v247)))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 12:
								v732 = base.I64_extend32_s(v232) + base.I64_extend32_s(v227)*int64(1000000)
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 13:
								v369 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v370 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v371 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v376 = base.B2i32(int32(2) < v370)
								if int32(2) < v370 {
									v377 = int32(_a_F_timestamp_part_common_16)
								} else {
									v377 = int32(_a_F_timestamp_part_common_18)
								}
								v378 = v377 + v369
								v383 = base.I32_div_s(v378, int32(4))
								v386 = base.I32_div_s(v378, int32(-100))
								v389 = base.I32_div_s(v378, int32(400))
								if int32(2) < v370 {
									v393 = int32(1)
								} else {
									v393 = int32(13)
								}
								v398 = base.I32_div_s((v393+v370)*int32(_a_F_timestamp_part_common_17), int32(256))
								v401 = v371 + v378*int32(365) + v383 + v386 + v389 + v398 - int32(_a_F_timestamp_part_common_19)
								if l1 != 0 {
									v403 = F_int64_to_numeric(m, base.I64_extend_i32_s(v401))
									mBase = m.M
									v404 = m.ExcPending
									if v404 != 0 {
										return int64(0)
									} else {
										v406 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
										v407 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
										v408 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
										v409 = int32(60)
										v419 = F_int64_to_numeric(m, base.I64_extend32_s(v232)+base.I64_extend_i32_s(v406+(v407+v408*v409)*v409)*int64(1000000))
										mBase = m.M
										v420 = m.ExcPending
										if v420 != 0 {
											return int64(0)
										} else {
											v422 = F_int64_to_numeric(m, int64(86400000000))
											mBase = m.M
											v423 = m.ExcPending
											if v423 != 0 {
												return int64(0)
											} else {
												v425 = F_numeric_div_safe(m, v419, v422, int32(0))
												mBase = m.M
												v426 = m.ExcPending
												if v426 != 0 {
													return int64(0)
												} else {
													v428 = F_numeric_add_safe(m, v403, v425, int32(0))
													mBase = m.M
													v429 = m.ExcPending
													if v429 != 0 {
														return int64(0)
													} else {
														v746 = base.I64_extend_i32_u(v428)
														m.G0 = v15 + int32(96)
														return v746
													}
												}
											}
										}
									}
								} else {
									v434 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
									v435 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
									v436 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
									v437 = int32(60)
									v746 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v233), float64(1e+06)), base.F64_convert_i32_s(v434+(v435+v436*v437)*v437)), float64(86400)), base.F64_convert_i32_s(v401)))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 14, 19:
								v497 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v498 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v499 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v504 = base.B2i32(int32(2) < v498)
								if int32(2) < v498 {
									v505 = int32(_a_F_timestamp_part_common_16)
								} else {
									v505 = int32(_a_F_timestamp_part_common_18)
								}
								v506 = v505 + v497
								v511 = base.I32_div_s(v506, int32(4))
								v514 = base.I32_div_s(v506, int32(-100))
								v517 = base.I32_div_s(v506, int32(400))
								if int32(2) < v498 {
									v521 = int32(1)
								} else {
									v521 = int32(13)
								}
								v526 = base.I32_div_s((v521+v498)*int32(_a_F_timestamp_part_common_17), int32(256))
								v532 = int32(7)
								v533 = base.I32_rem_s(v499+v506*int32(365)+v511+v514+v517+v526-int32(_a_F_timestamp_part_common_19)+int32(1), v532)
								if v533 < int32(0) {
									v538 = v533 + v532
								} else {
									v538 = v533
								}
								v539 = base.I64_extend_i32_s(v538)
								if v538 != 0 {
									v541 = v539
								} else {
									v541 = int64(7)
								}
								v542 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
								if v542 == int32(37) {
									v545 = v541
								} else {
									v545 = v539
								}
								v732 = v545
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							case 15:
								v546 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v547 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v548 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v553 = base.B2i32(int32(2) < v547)
								if int32(2) < v547 {
									v554 = int32(_a_F_timestamp_part_common_16)
								} else {
									v554 = int32(_a_F_timestamp_part_common_18)
								}
								v555 = v554 + v546
								v560 = base.I32_div_s(v555, int32(4))
								v563 = base.I32_div_s(v555, int32(-100))
								v566 = base.I32_div_s(v555, int32(400))
								if int32(2) < v547 {
									v570 = int32(1)
								} else {
									v570 = int32(13)
								}
								v575 = base.I32_div_s((v570+v547)*int32(_a_F_timestamp_part_common_17), int32(256))
								v579 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v588 = int32(_a_F_timestamp_part_common_18) + v579
								v593 = base.I32_div_s(v588, int32(4))
								v596 = base.I32_div_s(v588, int32(-100))
								v599 = base.I32_div_s(v588, int32(400))
								v608 = base.I32_div_s(int32(_a_F_timestamp_part_common_20), int32(256))
								v732 = base.I64_extend_i32_s(v548 + v555*int32(365) + v560 + v563 + v566 + v575 - int32(_a_F_timestamp_part_common_19) - (int32(1) + v588*int32(365) + v593 + v596 + v599 + v608 - int32(_a_F_timestamp_part_common_19)) + int32(1))
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v619 = m.ExcPending
								if v619 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v622 = m.ExcPending
									if v622 != 0 {
										return int64(0)
									} else {
										v624 = F_format_type_be(m, int32(1114))
										mBase = m.M
										v625 = m.ExcPending
										if v625 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v624
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v58
											F_errmsg(m, int32(_a_F_timestamp_part_common_21), v15+int32(16))
											mBase = m.M
											v632 = m.ExcPending
											if v632 != 0 {
												return int64(0)
											} else {
												F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_22), int32(_a_F_timestamp_part_common_9))
												mBase = m.M
												v637 = m.ExcPending
												if v637 != 0 {
													return int64(0)
												} else {
													base.Wasm_trap_unreachable()
													for {
													}
												}
											}
										}
									}
								}
							case 18:
								v450 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v451 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v452 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v454 = F_date2j(m, v450, v451, v452)
								mBase = m.M
								v455 = int32(1)
								v457 = F_date2j(m, v450, v455, int32(4))
								mBase = m.M
								v460 = F_j2day(m, v457-v455)
								mBase = m.M
								if v454 < v457-v460 {
									v463 = int32(1)
									v464 = v450 - v463
									v467 = F_date2j(m, v464, v463, int32(4))
									mBase = m.M
									v470 = F_j2day(m, v467-v463)
									mBase = m.M
									v471 = v464
									v472 = v467
									v473 = v470
								} else {
									v471 = v450
									v472 = v457
									v473 = v460
								}
								if int32(357) <= v473+(v454-v472) {
									v478 = int32(1)
									v479 = v471 + v478
									v482 = F_date2j(m, v479, v478, int32(4))
									mBase = m.M
									v485 = F_j2day(m, v482-v478)
									mBase = m.M
									if v454 < v482-v485 {
										v488 = v471
									} else {
										v488 = v479
									}
									v491 = v488
								} else {
									v491 = v471
								}
								v732 = base.I64_extend_i32_s(v491) - base.I64_extend_i32_u(base.B2i32(v491 <= int32(0)))
								if l1 != 0 {
									v733 = F_int64_to_numeric(m, v732)
									mBase = m.M
									v734 = m.ExcPending
									if v734 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v733)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
									m.G0 = v15 + int32(96)
									return v746
								}
							}
						}
					}
				} else {
					v638 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
					if v638 == int32(11) {
						v641 = F_SetEpochTimestamp(m)
						mBase = m.M
						v642 = m.ExcPending
						if v642 != 0 {
							return int64(0)
						} else {
							v644 = v641 + int64(9223372036854775807)
							if l1 != 0 {
								if v29 < v644 {
									v648 = F_int64_div_fast_to_numeric(m, v29-v641, int32(6))
									mBase = m.M
									v649 = m.ExcPending
									if v649 != 0 {
										return int64(0)
									} else {
										v746 = base.I64_extend_i32_u(v648)
										m.G0 = v15 + int32(96)
										return v746
									}
								} else {
									v653 = F_int64_to_numeric(m, v29)
									mBase = m.M
									v654 = m.ExcPending
									if v654 != 0 {
										return int64(0)
									} else {
										v655 = F_int64_to_numeric(m, v641)
										mBase = m.M
										v656 = m.ExcPending
										if v656 != 0 {
											return int64(0)
										} else {
											v658 = F_numeric_sub_safe(m, v653, v655, int32(0))
											mBase = m.M
											v659 = m.ExcPending
											if v659 != 0 {
												return int64(0)
											} else {
												v661 = F_int64_to_numeric(m, int64(1000000))
												mBase = m.M
												v662 = m.ExcPending
												if v662 != 0 {
													return int64(0)
												} else {
													v664 = F_numeric_div_safe(m, v658, v661, int32(0))
													mBase = m.M
													v665 = m.ExcPending
													if v665 != 0 {
														return int64(0)
													} else {
														v668 = F_DirectFunctionCall2Coll(m, int32(1390), int32(0), base.I64_extend_i32_u(v664), int64(6))
														mBase = m.M
														v669 = m.ExcPending
														if v669 != 0 {
															return int64(0)
														} else {
															v671 = F_pg_detoast_datum(m, base.I32_wrap_i64(v668))
															mBase = m.M
															v672 = m.ExcPending
															if v672 != 0 {
																return int64(0)
															} else {
																v746 = base.I64_extend_i32_u(v671)
																m.G0 = v15 + int32(96)
																return v746
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								if v29 < v644 {
									v680 = base.F64_convert_i64_s(v29 - v641)
								} else {
									v680 = base.F64_sub(base.F64_convert_i64_s(v29), base.F64_convert_i64_s(v641))
								}
								v746 = base.I64_reinterpret_f64(base.F64_div(v680, float64(1e+06)))
								m.G0 = v15 + int32(96)
								return v746
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v687 = m.ExcPending
						if v687 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v690 = m.ExcPending
							if v690 != 0 {
								return int64(0)
							} else {
								v692 = F_format_type_be(m, int32(1114))
								mBase = m.M
								v693 = m.ExcPending
								if v693 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v692
									*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v58
									F_errmsg(m, int32(_a_F_timestamp_part_common_21), v15+int32(32))
									mBase = m.M
									v700 = m.ExcPending
									if v700 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_23), int32(_a_F_timestamp_part_common_9))
										mBase = m.M
										v705 = m.ExcPending
										if v705 != 0 {
											return int64(0)
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
				v77 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
				v81 = F_NonFiniteTimestampTzPart(m, v72, v77, v58, base.B2i32(v29 == int64(-9223372036854775807-1)), int32(0))
				mBase = m.M
				v82 = m.ExcPending
				if v82 != 0 {
					return int64(0)
				} else {
					if base.F64_ne(v81, float64(0)) != 0 {
						if l1 != 0 {
							if base.F64_lt(v81, float64(0)) != 0 {
								v92 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), int64(11926), int64(0), int64(-1))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
									return int64(0)
								} else {
									v746 = v92
									m.G0 = v15 + int32(96)
									return v746
								}
							} else {
								if base.F64_gt(v81, float64(0)) == int32(0) {
									if v72 != 0 {
										if v72 != int32(17) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v709 = m.ExcPending
											if v709 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v712 = m.ExcPending
												if v712 != 0 {
													return int64(0)
												} else {
													v714 = F_format_type_be(m, int32(1114))
													mBase = m.M
													v715 = m.ExcPending
													if v715 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v714
														*(*int32)(unsafe.Add(mBase, uint32(v15))) = v58
														F_errmsg(m, int32(_a_F_timestamp_part_common_6), v15)
														mBase = m.M
														v720 = m.ExcPending
														if v720 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_8), int32(_a_F_timestamp_part_common_9))
															mBase = m.M
															v725 = m.ExcPending
															if v725 != 0 {
																return int64(0)
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
											v113 = base.I64_div_s(v29, int64(86400000000))
											if base.Ui64(int64(172799999999)) <= base.Ui64(v29+int64(86399999999)) {
												v121 = v113 * int64(-86400000000)
											} else {
												v121 = int64(0)
											}
											v122 = v121 + v29
											v125 = v122>>(uint(int64(63))%64) + v113
											if v125 <= int64(-2451546) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v754 = m.ExcPending
												if v754 != 0 {
													return int64(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v757 = m.ExcPending
													if v757 != 0 {
														return int64(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamp_part_common_10), int32(0))
														mBase = m.M
														v761 = m.ExcPending
														if v761 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_11), int32(_a_F_timestamp_part_common_9))
															mBase = m.M
															v766 = m.ExcPending
															if v766 != 0 {
																return int64(0)
															} else {
																base.Wasm_trap_unreachable()
																for {
																}
															}
														}
													}
												}
											} else {
												v128 = base.I32_wrap_i64(v125)
												v140 = v128 + int32(_a_F_timestamp_part_common_12)
												v141 = int32(_a_F_timestamp_part_common_13)
												v142 = base.I32_div_u_s(v140, v141)
												v143 = int32(3)
												v149 = int32(2)
												v154 = base.I32_div_u_s((v142*int32(1073595727)+v140)<<(uint(v149)%32)|v143, v141)
												v157 = v128 + int32(_a_F_timestamp_part_common_14) + v142*v143 + v154 + int32(_a_F_timestamp_part_common_15)
												v158 = int32(1461)
												v159 = base.I32_div_u_s(v157, v158)
												v162 = v159*int32(-1461) + v157
												v164 = v162 << (uint(v149) % 32)
												if base.Ui32(v158) <= base.Ui32(v164) {
													v170 = base.I32_rem_u_s(v162+int32(305), int32(365))
													v175 = v170
												} else {
													v174 = base.I32_rem_u_s(v162+int32(306), int32(366))
													v175 = v174
												}
												v177 = base.I32_div_u_s(v164, int32(1461))
												*(*int32)(unsafe.Add(mBase, uint32(v15+int32(68)))) = v177 + v159<<(uint(int32(2))%32) - int32(_a_F_timestamp_part_common_16)
												v185 = v175 + int32(123)
												v189 = int32(base.Ui32(v185*int32(2141)) >> (uint(int32(16)) % 32))
												*(*int32)(unsafe.Add(mBase, uint32(v15+int32(60)))) = v185 - int32(base.Ui32(v189*int32(_a_F_timestamp_part_common_17))>>(uint(int32(8))%32))
												v199 = base.I32_rem_u_s(v189+int32(10), int32(12))
												*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v199 + int32(1)
												if v122 < int64(0) {
													v207 = v122 + int64(86400000000)
												} else {
													v207 = v122
												}
												v209 = base.I64_div_s(v207, int64(3600000000))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+56)) = uint32(v209)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = int64(4294967295)
												v215 = base.I64_extend32_s(v209)
												v218 = v215*int64(-3600000000) + v207
												v220 = base.I64_div_s(v218, int64(60000000))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+52)) = uint32(v220)
												v222 = base.I64_extend32_s(v220)
												v225 = v218 + v222*int64(-60000000)
												v227 = base.I64_div_s(v225, int64(1000000))
												v228 = base.I32_wrap_i64(v227)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v228
												v232 = v227*int64(4293967296) + v225
												v233 = base.I32_wrap_i64(v232)
												v234 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
												switch v234 - int32(18) {
												case 0:
													if l1 != 0 {
														v260 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v232)+base.I64_extend32_s(v227)*int64(1000000), int32(6))
														mBase = m.M
														v261 = m.ExcPending
														if v261 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v260)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_convert_i32_s(v233), float64(1e+06)), base.F64_convert_i32_s(v228)))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 1:
													v732 = v222
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 2:
													v732 = v215
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 3:
													v269 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+60)))
													v732 = v269
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 4:
													v279 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v280 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v281 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v283 = F_date2j(m, v279, v280, v281)
													mBase = m.M
													v284 = int32(1)
													v286 = F_date2j(m, v279, v284, int32(4))
													mBase = m.M
													v289 = F_j2day(m, v286-v284)
													mBase = m.M
													if v283 < v286-v289 {
														v292 = int32(1)
														v296 = F_date2j(m, v279-v292, v292, int32(4))
														mBase = m.M
														v299 = F_j2day(m, v296-v292)
														mBase = m.M
														v300 = v296
														v301 = v299
													} else {
														v300 = v286
														v301 = v289
													}
													v303 = v301 - v300 + v283
													if int32(357) <= v303 {
														v306 = int32(1)
														v310 = F_date2j(m, v279+v306, v306, int32(4))
														mBase = m.M
														v313 = F_j2day(m, v310-v306)
														mBase = m.M
														v314 = v310 - v313
														if v283 < v314 {
															v317 = v303
														} else {
															v317 = v283 - v314
														}
														v319 = v317
													} else {
														v319 = v303
													}
													v321 = base.I32_div_s(v319, int32(7))
													v732 = base.I64_extend_i32_s(v321 + int32(1))
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 5:
													v270 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+64)))
													v732 = v270
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 6:
													v271 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v272 = int32(1)
													v275 = base.I32_div_s(v271-v272, int32(3))
													v732 = base.I64_extend_i32_s(v275 + v272)
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 7:
													v325 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													if int32(0) < v325 {
														v732 = base.I64_extend_i32_u(v325)
													} else {
														v732 = base.I64_extend_i32_s(v325 - int32(1))
													}
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 8:
													v332 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													if int32(0) <= v332 {
														v336 = base.I32_div_u_s(v332, int32(10))
														v732 = base.I64_extend_i32_u(v336)
													} else {
														v341 = base.I32_div_s(int32(9)-v332, int32(-10))
														v732 = base.I64_extend_i32_s(v341)
													}
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 9:
													v343 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													if int32(0) < v343 {
														v349 = base.I32_div_s(v343+int32(99), int32(100))
														v732 = base.I64_extend_i32_s(v349)
													} else {
														v354 = base.I32_div_s(int32(100)-v343, int32(-100))
														v732 = base.I64_extend_i32_s(v354)
													}
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 10:
													v356 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													if int32(0) < v356 {
														v362 = base.I32_div_s(v356+int32(999), int32(1000))
														v732 = base.I64_extend_i32_s(v362)
													} else {
														v367 = base.I32_div_s(int32(1000)-v356, int32(-1000))
														v732 = base.I64_extend_i32_s(v367)
													}
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 11:
													if l1 != 0 {
														v243 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v232)+base.I64_extend32_s(v227)*int64(1000000), int32(3))
														mBase = m.M
														v244 = m.ExcPending
														if v244 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v243)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v247 = float64(1000)
														v746 = base.I64_reinterpret_f64(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v228), v247), base.F64_div(base.F64_convert_i32_s(v233), v247)))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 12:
													v732 = base.I64_extend32_s(v232) + base.I64_extend32_s(v227)*int64(1000000)
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 13:
													v369 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v370 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v371 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v376 = base.B2i32(int32(2) < v370)
													if int32(2) < v370 {
														v377 = int32(_a_F_timestamp_part_common_16)
													} else {
														v377 = int32(_a_F_timestamp_part_common_18)
													}
													v378 = v377 + v369
													v383 = base.I32_div_s(v378, int32(4))
													v386 = base.I32_div_s(v378, int32(-100))
													v389 = base.I32_div_s(v378, int32(400))
													if int32(2) < v370 {
														v393 = int32(1)
													} else {
														v393 = int32(13)
													}
													v398 = base.I32_div_s((v393+v370)*int32(_a_F_timestamp_part_common_17), int32(256))
													v401 = v371 + v378*int32(365) + v383 + v386 + v389 + v398 - int32(_a_F_timestamp_part_common_19)
													if l1 != 0 {
														v403 = F_int64_to_numeric(m, base.I64_extend_i32_s(v401))
														mBase = m.M
														v404 = m.ExcPending
														if v404 != 0 {
															return int64(0)
														} else {
															v406 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
															v407 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
															v408 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
															v409 = int32(60)
															v419 = F_int64_to_numeric(m, base.I64_extend32_s(v232)+base.I64_extend_i32_s(v406+(v407+v408*v409)*v409)*int64(1000000))
															mBase = m.M
															v420 = m.ExcPending
															if v420 != 0 {
																return int64(0)
															} else {
																v422 = F_int64_to_numeric(m, int64(86400000000))
																mBase = m.M
																v423 = m.ExcPending
																if v423 != 0 {
																	return int64(0)
																} else {
																	v425 = F_numeric_div_safe(m, v419, v422, int32(0))
																	mBase = m.M
																	v426 = m.ExcPending
																	if v426 != 0 {
																		return int64(0)
																	} else {
																		v428 = F_numeric_add_safe(m, v403, v425, int32(0))
																		mBase = m.M
																		v429 = m.ExcPending
																		if v429 != 0 {
																			return int64(0)
																		} else {
																			v746 = base.I64_extend_i32_u(v428)
																			m.G0 = v15 + int32(96)
																			return v746
																		}
																	}
																}
															}
														}
													} else {
														v434 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
														v435 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
														v436 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
														v437 = int32(60)
														v746 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v233), float64(1e+06)), base.F64_convert_i32_s(v434+(v435+v436*v437)*v437)), float64(86400)), base.F64_convert_i32_s(v401)))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 14, 19:
													v497 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v498 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v499 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v504 = base.B2i32(int32(2) < v498)
													if int32(2) < v498 {
														v505 = int32(_a_F_timestamp_part_common_16)
													} else {
														v505 = int32(_a_F_timestamp_part_common_18)
													}
													v506 = v505 + v497
													v511 = base.I32_div_s(v506, int32(4))
													v514 = base.I32_div_s(v506, int32(-100))
													v517 = base.I32_div_s(v506, int32(400))
													if int32(2) < v498 {
														v521 = int32(1)
													} else {
														v521 = int32(13)
													}
													v526 = base.I32_div_s((v521+v498)*int32(_a_F_timestamp_part_common_17), int32(256))
													v532 = int32(7)
													v533 = base.I32_rem_s(v499+v506*int32(365)+v511+v514+v517+v526-int32(_a_F_timestamp_part_common_19)+int32(1), v532)
													if v533 < int32(0) {
														v538 = v533 + v532
													} else {
														v538 = v533
													}
													v539 = base.I64_extend_i32_s(v538)
													if v538 != 0 {
														v541 = v539
													} else {
														v541 = int64(7)
													}
													v542 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
													if v542 == int32(37) {
														v545 = v541
													} else {
														v545 = v539
													}
													v732 = v545
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												case 15:
													v546 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v547 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v548 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v553 = base.B2i32(int32(2) < v547)
													if int32(2) < v547 {
														v554 = int32(_a_F_timestamp_part_common_16)
													} else {
														v554 = int32(_a_F_timestamp_part_common_18)
													}
													v555 = v554 + v546
													v560 = base.I32_div_s(v555, int32(4))
													v563 = base.I32_div_s(v555, int32(-100))
													v566 = base.I32_div_s(v555, int32(400))
													if int32(2) < v547 {
														v570 = int32(1)
													} else {
														v570 = int32(13)
													}
													v575 = base.I32_div_s((v570+v547)*int32(_a_F_timestamp_part_common_17), int32(256))
													v579 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v588 = int32(_a_F_timestamp_part_common_18) + v579
													v593 = base.I32_div_s(v588, int32(4))
													v596 = base.I32_div_s(v588, int32(-100))
													v599 = base.I32_div_s(v588, int32(400))
													v608 = base.I32_div_s(int32(_a_F_timestamp_part_common_20), int32(256))
													v732 = base.I64_extend_i32_s(v548 + v555*int32(365) + v560 + v563 + v566 + v575 - int32(_a_F_timestamp_part_common_19) - (int32(1) + v588*int32(365) + v593 + v596 + v599 + v608 - int32(_a_F_timestamp_part_common_19)) + int32(1))
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v619 = m.ExcPending
													if v619 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v622 = m.ExcPending
														if v622 != 0 {
															return int64(0)
														} else {
															v624 = F_format_type_be(m, int32(1114))
															mBase = m.M
															v625 = m.ExcPending
															if v625 != 0 {
																return int64(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v624
																*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v58
																F_errmsg(m, int32(_a_F_timestamp_part_common_21), v15+int32(16))
																mBase = m.M
																v632 = m.ExcPending
																if v632 != 0 {
																	return int64(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_22), int32(_a_F_timestamp_part_common_9))
																	mBase = m.M
																	v637 = m.ExcPending
																	if v637 != 0 {
																		return int64(0)
																	} else {
																		base.Wasm_trap_unreachable()
																		for {
																		}
																	}
																}
															}
														}
													}
												case 18:
													v450 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v451 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v452 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v454 = F_date2j(m, v450, v451, v452)
													mBase = m.M
													v455 = int32(1)
													v457 = F_date2j(m, v450, v455, int32(4))
													mBase = m.M
													v460 = F_j2day(m, v457-v455)
													mBase = m.M
													if v454 < v457-v460 {
														v463 = int32(1)
														v464 = v450 - v463
														v467 = F_date2j(m, v464, v463, int32(4))
														mBase = m.M
														v470 = F_j2day(m, v467-v463)
														mBase = m.M
														v471 = v464
														v472 = v467
														v473 = v470
													} else {
														v471 = v450
														v472 = v457
														v473 = v460
													}
													if int32(357) <= v473+(v454-v472) {
														v478 = int32(1)
														v479 = v471 + v478
														v482 = F_date2j(m, v479, v478, int32(4))
														mBase = m.M
														v485 = F_j2day(m, v482-v478)
														mBase = m.M
														if v454 < v482-v485 {
															v488 = v471
														} else {
															v488 = v479
														}
														v491 = v488
													} else {
														v491 = v471
													}
													v732 = base.I64_extend_i32_s(v491) - base.I64_extend_i32_u(base.B2i32(v491 <= int32(0)))
													if l1 != 0 {
														v733 = F_int64_to_numeric(m, v732)
														mBase = m.M
														v734 = m.ExcPending
														if v734 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v733)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v746 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v732))
														m.G0 = v15 + int32(96)
														return v746
													}
												}
											}
										}
									} else {
										v638 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
										if v638 == int32(11) {
											v641 = F_SetEpochTimestamp(m)
											mBase = m.M
											v642 = m.ExcPending
											if v642 != 0 {
												return int64(0)
											} else {
												v644 = v641 + int64(9223372036854775807)
												if l1 != 0 {
													if v29 < v644 {
														v648 = F_int64_div_fast_to_numeric(m, v29-v641, int32(6))
														mBase = m.M
														v649 = m.ExcPending
														if v649 != 0 {
															return int64(0)
														} else {
															v746 = base.I64_extend_i32_u(v648)
															m.G0 = v15 + int32(96)
															return v746
														}
													} else {
														v653 = F_int64_to_numeric(m, v29)
														mBase = m.M
														v654 = m.ExcPending
														if v654 != 0 {
															return int64(0)
														} else {
															v655 = F_int64_to_numeric(m, v641)
															mBase = m.M
															v656 = m.ExcPending
															if v656 != 0 {
																return int64(0)
															} else {
																v658 = F_numeric_sub_safe(m, v653, v655, int32(0))
																mBase = m.M
																v659 = m.ExcPending
																if v659 != 0 {
																	return int64(0)
																} else {
																	v661 = F_int64_to_numeric(m, int64(1000000))
																	mBase = m.M
																	v662 = m.ExcPending
																	if v662 != 0 {
																		return int64(0)
																	} else {
																		v664 = F_numeric_div_safe(m, v658, v661, int32(0))
																		mBase = m.M
																		v665 = m.ExcPending
																		if v665 != 0 {
																			return int64(0)
																		} else {
																			v668 = F_DirectFunctionCall2Coll(m, int32(1390), int32(0), base.I64_extend_i32_u(v664), int64(6))
																			mBase = m.M
																			v669 = m.ExcPending
																			if v669 != 0 {
																				return int64(0)
																			} else {
																				v671 = F_pg_detoast_datum(m, base.I32_wrap_i64(v668))
																				mBase = m.M
																				v672 = m.ExcPending
																				if v672 != 0 {
																					return int64(0)
																				} else {
																					v746 = base.I64_extend_i32_u(v671)
																					m.G0 = v15 + int32(96)
																					return v746
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													if v29 < v644 {
														v680 = base.F64_convert_i64_s(v29 - v641)
													} else {
														v680 = base.F64_sub(base.F64_convert_i64_s(v29), base.F64_convert_i64_s(v641))
													}
													v746 = base.I64_reinterpret_f64(base.F64_div(v680, float64(1e+06)))
													m.G0 = v15 + int32(96)
													return v746
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v687 = m.ExcPending
											if v687 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v690 = m.ExcPending
												if v690 != 0 {
													return int64(0)
												} else {
													v692 = F_format_type_be(m, int32(1114))
													mBase = m.M
													v693 = m.ExcPending
													if v693 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v692
														*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v58
														F_errmsg(m, int32(_a_F_timestamp_part_common_21), v15+int32(32))
														mBase = m.M
														v700 = m.ExcPending
														if v700 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_23), int32(_a_F_timestamp_part_common_9))
															mBase = m.M
															v705 = m.ExcPending
															if v705 != 0 {
																return int64(0)
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
									v103 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), int64(11937), int64(0), int64(-1))
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int64(0)
									} else {
										v746 = v103
										m.G0 = v15 + int32(96)
										return v746
									}
								}
							}
						} else {
							v746 = base.I64_reinterpret_f64(v81)
							m.G0 = v15 + int32(96)
							return v746
						}
					} else {
						v106 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v106)
						v746 = int64(0)
						m.G0 = v15 + int32(96)
						return v746
					}
				}
			}
		}
	}
}
