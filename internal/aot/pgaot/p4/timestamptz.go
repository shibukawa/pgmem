package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extract_timestamptz(m *base.Module, l0 int32) int32 {
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timestamptz_part_common(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		return v3
	}
}
func F_timestamptz_bin(m *base.Module, l0 int32) int32 {
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14000(m, l0, int32(_a_F_timestamptz_bin_0), int32(_a_F_timestamptz_bin_1), int32(_a_F_timestamptz_bin_2), int32(_a_F_timestamptz_bin_3), int32(_a_F_timestamptz_bin_4), int32(_a_F_timestamptz_bin_5), int32(_a_F_timestamptz_bin_6), int32(_a_F_timestamptz_bin_7))
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		return v10
	}
}
func F_timestamptz_cmp_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v4 int32
	_ = v4
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v15 int32
	_ = v15
	var v37 int32
	_ = v37
	var v38 int32
	_ = v38
	var v45 int64
	_ = v45
	var v57 int32
	_ = v57
	var v62 int32
	_ = v62
	var v63 int64
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v5 = *(*int64)(unsafe.Add(mBase, uint32(v4)))
	v8 = m.G0
	v10 = v8 - int32(48)
	m.G0 = v10
	if v3 == int32(-2147483648) {
		v15 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v5)
		mBase = m.M
		v66 = v15
	} else {
		if v3 == int32(2147483647) {
			v63 = int64(9223372036854775807)
			v64 = F_timestamp_cmp_internal(m, v63, v5)
			mBase = m.M
			v66 = v64
		} else {
			if v3 <= int32(106751982) {
				F_j2date(m, v3+int32(_a_F_timestamptz_cmp_date_0), v10+int32(24), v10+int32(20), v10+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
				v37 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_cmp_date[0]))
				v38 = F_DetermineTimeZoneOffset(m, v10+int32(4), v37)
				mBase = m.M
				v45 = base.I64_extend_i32_s(v38)*int64(1000000) + base.I64_extend_i32_s(v3)*int64(86400000000)
				if base.Ui64(v45+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v63 = v45
					v64 = F_timestamp_cmp_internal(m, v63, v5)
					mBase = m.M
					v66 = v64
				} else {
					if v45 < int64(-211813488000000000) {
						if v5 == int64(-9223372036854775807-1) {
							v62 = int32(1)
						} else {
							v62 = int32(-1)
						}
						v66 = v62
					} else {
						if v5 == int64(9223372036854775807) {
							v57 = int32(-1)
						} else {
							v57 = int32(1)
						}
						v66 = v57
					}
				}
			} else {
				if v5 == int64(9223372036854775807) {
					v57 = int32(-1)
				} else {
					v57 = int32(1)
				}
				v66 = v57
			}
		}
	}
	m.G0 = v10 + int32(48)
	return int32(0) - v66
}
func F_timestamptz_cmp_timestamp(m *base.Module, l0 int32) int32 {
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
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v31 int32
	_ = v31
	var v37 int32
	_ = v37
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		if v10 == int64(9223372036854775807) {
			v23 = int32(1)
		} else {
			v23 = int32(-1)
		}
		v24 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v24 {
			v37 = v23
		} else {
			if v10 == int64(-9223372036854775807-1) {
				v31 = int32(-1)
			} else {
				v31 = int32(1)
			}
			if v24 < int32(0) {
				v37 = v31
			} else {
				v37 = base.B2i32(v15 < v10) - base.B2i32(v10 < v15)
			}
		}
		m.G0 = v7 + int32(16)
		return v37
	}
}
func F_timestamptz_eq_timestamp(m *base.Module, l0 int32) int32 {
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
	var v19 int32
	_ = v19
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v19 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		m.G0 = v7 + int32(16)
		return base.B2i32(v19 == int32(0)) & base.B2i32(v10 == v15)
	}
}
func F_timestamptz_ge_date(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v4 int64
	_ = v4
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v14 int32
	_ = v14
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v44 int64
	_ = v44
	var v56 int32
	_ = v56
	var v61 int32
	_ = v61
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v3 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v4 = *(*int64)(unsafe.Add(mBase, uint32(v3)))
	v7 = m.G0
	v9 = v7 - int32(48)
	m.G0 = v9
	if v2 == int32(-2147483648) {
		v14 = F_timestamp_cmp_internal(m, int64(-9223372036854775807-1), v4)
		mBase = m.M
		v65 = v14
	} else {
		if v2 == int32(2147483647) {
			v62 = int64(9223372036854775807)
			v63 = F_timestamp_cmp_internal(m, v62, v4)
			mBase = m.M
			v65 = v63
		} else {
			if v2 <= int32(106751982) {
				F_j2date(m, v2+int32(_a_F_timestamptz_ge_date_0), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_ge_date[0]))
				v37 = F_DetermineTimeZoneOffset(m, v9+int32(4), v36)
				mBase = m.M
				v44 = base.I64_extend_i32_s(v37)*int64(1000000) + base.I64_extend_i32_s(v2)*int64(86400000000)
				if base.Ui64(v44+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
					v62 = v44
					v63 = F_timestamp_cmp_internal(m, v62, v4)
					mBase = m.M
					v65 = v63
				} else {
					if v44 < int64(-211813488000000000) {
						if v4 == int64(-9223372036854775807-1) {
							v61 = int32(1)
						} else {
							v61 = int32(-1)
						}
						v65 = v61
					} else {
						if v4 == int64(9223372036854775807) {
							v56 = int32(-1)
						} else {
							v56 = int32(1)
						}
						v65 = v56
					}
				}
			} else {
				if v4 == int64(9223372036854775807) {
					v56 = int32(-1)
				} else {
					v56 = int32(1)
				}
				v65 = v56
			}
		}
	}
	m.G0 = v9 + int32(48)
	return base.B2i32(v65 <= int32(0))
}
func F_timestamptz_in(m *base.Module, l0 int32) int32 {
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
	var v17 int32
	_ = v17
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
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
	var v174 int64
	_ = v174
	var v178 int32
	_ = v178
	var v183 int64
	_ = v183
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v215 int32
	_ = v215
	var v217 int32
	_ = v217
	var v221 int32
	_ = v221
	var v226 int32
	_ = v226
	var v227 int64
	_ = v227
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v238 int32
	_ = v238
	var v239 int64
	_ = v239
	var v240 int32
	_ = v240
	var v241 int32
	_ = v241
	var v242 int32
	_ = v242
	v11 = m.G0
	v13 = v11 - int32(512)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v22 = v13 + int32(336)
	v24 = v13 + int32(224)
	v27 = F_ParseDateTime(m, v17, v13-int32(-64), int32(153), v22, v24, v13+int32(444))
	mBase = m.M
	if v27 == int32(0) {
		v30 = *(*int32)(unsafe.Add(mBase, uint32(v13)+444))
		v41 = F_DecodeDateTime(m, v22, v24, v30, v13+int32(448), v13+int32(456), v13+int32(500), v13+int32(452), v13+int32(56))
		mBase = m.M
		v44 = m.ExcPending
		if v44 != 0 {
			return int32(0)
		} else {
			if v41 == int32(0) {
				v57 = *(*int32)(unsafe.Add(mBase, uint32(v13)+448))
				switch v57 - int32(2) {
				case 0:
					v60 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+500)))
					v61 = *(*int32)(unsafe.Add(mBase, uint32(v13)+476))
					if v61 <= int32(-4713) {
						if v61 != int32(-4713) {
							v192 = int32(0)
							v193 = F_errsave_start(m, v15)
							mBase = m.M
							v194 = m.ExcPending
							if v194 != 0 {
								return int32(0)
							} else {
								if v193 == int32(0) {
									v242 = v192
									m.G0 = v13 + int32(512)
									return v242
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v199 = m.ExcPending
									if v199 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
										F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
										mBase = m.M
										v205 = m.ExcPending
										if v205 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
											mBase = m.M
											v210 = m.ExcPending
											if v210 != 0 {
												return int32(0)
											} else {
												v242 = v192
												m.G0 = v13 + int32(512)
												return v242
											}
										}
									}
								}
							}
						} else {
							v66 = *(*int32)(unsafe.Add(mBase, uint32(v13)+472))
							if int32(10) < v66 {
								v77 = v66
								v79 = v13 + int32(32)
								v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+468))
								v85 = base.B2i32(int32(2) < v77)
								if int32(2) < v77 {
									v86 = int32(_a_F_timestamptz_in_3)
								} else {
									v86 = int32(_a_F_timestamptz_in_4)
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
								v107 = base.I32_div_s((v102+v77)*int32(_a_F_timestamptz_in_5), int32(256))
								v113 = base.I64_extend_i32_s(v80 + v87*int32(365) + v92 + v95 + v98 + v107 - int32(_a_F_timestamptz_in_6) - int32(_a_F_timestamptz_in_7))
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
								v154 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								v155 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
								if v154 != v155>>(uint(int64(63))%64) {
									v192 = int32(0)
									v193 = F_errsave_start(m, v15)
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int32(0)
									} else {
										if v193 == int32(0) {
											v242 = v192
											m.G0 = v13 + int32(512)
											return v242
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
												F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return int32(0)
													} else {
														v242 = v192
														m.G0 = v13 + int32(512)
														return v242
													}
												}
											}
										}
									}
								} else {
									v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+456))
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+460))
									v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+464))
									v162 = int32(60)
									v171 = base.I64_extend_i32_s(v159+(v160+v161*v162)*v162)*int64(1000000) + v60
									v174 = v155 + v171
									if base.B2i32(v171 < int64(0))^base.B2i32(v174 < v155) != 0 {
										v192 = int32(0)
										v193 = F_errsave_start(m, v15)
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											if v193 == int32(0) {
												v242 = v192
												m.G0 = v13 + int32(512)
												return v242
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
													F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return int32(0)
														} else {
															v242 = v192
															m.G0 = v13 + int32(512)
															return v242
														}
													}
												}
											}
										}
									} else {
										v178 = *(*int32)(unsafe.Add(mBase, uint32(v13)+452))
										v183 = base.I64_extend_i32_s(int32(0)-v178)*int64(-1000000) + v174
										*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v183
										if base.Ui64(v183+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
											F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
											mBase = m.M
											v238 = m.ExcPending
											if v238 != 0 {
												return int32(0)
											} else {
												v239 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
												v240 = F_Int64GetDatum(m, v239)
												mBase = m.M
												v241 = m.ExcPending
												if v241 != 0 {
													return int32(0)
												} else {
													v242 = v240
													m.G0 = v13 + int32(512)
													return v242
												}
											}
										} else {
											v192 = int32(0)
											v193 = F_errsave_start(m, v15)
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int32(0)
											} else {
												if v193 == int32(0) {
													v242 = v192
													m.G0 = v13 + int32(512)
													return v242
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
														F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
																return int32(0)
															} else {
																v242 = v192
																m.G0 = v13 + int32(512)
																return v242
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v192 = int32(0)
								v193 = F_errsave_start(m, v15)
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int32(0)
								} else {
									if v193 == int32(0) {
										v242 = v192
										m.G0 = v13 + int32(512)
										return v242
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
											F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return int32(0)
												} else {
													v242 = v192
													m.G0 = v13 + int32(512)
													return v242
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v61 <= int32(_a_F_timestamptz_in_8) {
							v71 = *(*int32)(unsafe.Add(mBase, uint32(v13)+472))
							v77 = v71
							v79 = v13 + int32(32)
							v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+468))
							v85 = base.B2i32(int32(2) < v77)
							if int32(2) < v77 {
								v86 = int32(_a_F_timestamptz_in_3)
							} else {
								v86 = int32(_a_F_timestamptz_in_4)
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
							v107 = base.I32_div_s((v102+v77)*int32(_a_F_timestamptz_in_5), int32(256))
							v113 = base.I64_extend_i32_s(v80 + v87*int32(365) + v92 + v95 + v98 + v107 - int32(_a_F_timestamptz_in_6) - int32(_a_F_timestamptz_in_7))
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
							v154 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							v155 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
							if v154 != v155>>(uint(int64(63))%64) {
								v192 = int32(0)
								v193 = F_errsave_start(m, v15)
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int32(0)
								} else {
									if v193 == int32(0) {
										v242 = v192
										m.G0 = v13 + int32(512)
										return v242
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
											F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return int32(0)
												} else {
													v242 = v192
													m.G0 = v13 + int32(512)
													return v242
												}
											}
										}
									}
								}
							} else {
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+456))
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+460))
								v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+464))
								v162 = int32(60)
								v171 = base.I64_extend_i32_s(v159+(v160+v161*v162)*v162)*int64(1000000) + v60
								v174 = v155 + v171
								if base.B2i32(v171 < int64(0))^base.B2i32(v174 < v155) != 0 {
									v192 = int32(0)
									v193 = F_errsave_start(m, v15)
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int32(0)
									} else {
										if v193 == int32(0) {
											v242 = v192
											m.G0 = v13 + int32(512)
											return v242
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
												F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return int32(0)
													} else {
														v242 = v192
														m.G0 = v13 + int32(512)
														return v242
													}
												}
											}
										}
									}
								} else {
									v178 = *(*int32)(unsafe.Add(mBase, uint32(v13)+452))
									v183 = base.I64_extend_i32_s(int32(0)-v178)*int64(-1000000) + v174
									*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v183
									if base.Ui64(v183+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
										F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
										mBase = m.M
										v238 = m.ExcPending
										if v238 != 0 {
											return int32(0)
										} else {
											v239 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
											v240 = F_Int64GetDatum(m, v239)
											mBase = m.M
											v241 = m.ExcPending
											if v241 != 0 {
												return int32(0)
											} else {
												v242 = v240
												m.G0 = v13 + int32(512)
												return v242
											}
										}
									} else {
										v192 = int32(0)
										v193 = F_errsave_start(m, v15)
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											if v193 == int32(0) {
												v242 = v192
												m.G0 = v13 + int32(512)
												return v242
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
													F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return int32(0)
														} else {
															v242 = v192
															m.G0 = v13 + int32(512)
															return v242
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							if v61 != int32(_a_F_timestamptz_in_9) {
								v192 = int32(0)
								v193 = F_errsave_start(m, v15)
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int32(0)
								} else {
									if v193 == int32(0) {
										v242 = v192
										m.G0 = v13 + int32(512)
										return v242
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v199 = m.ExcPending
										if v199 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
											F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
												mBase = m.M
												v210 = m.ExcPending
												if v210 != 0 {
													return int32(0)
												} else {
													v242 = v192
													m.G0 = v13 + int32(512)
													return v242
												}
											}
										}
									}
								}
							} else {
								v74 = *(*int32)(unsafe.Add(mBase, uint32(v13)+472))
								if int32(5) < v74 {
									v192 = int32(0)
									v193 = F_errsave_start(m, v15)
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int32(0)
									} else {
										if v193 == int32(0) {
											v242 = v192
											m.G0 = v13 + int32(512)
											return v242
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v199 = m.ExcPending
											if v199 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
												F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
													mBase = m.M
													v210 = m.ExcPending
													if v210 != 0 {
														return int32(0)
													} else {
														v242 = v192
														m.G0 = v13 + int32(512)
														return v242
													}
												}
											}
										}
									}
								} else {
									v77 = v74
									v79 = v13 + int32(32)
									v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)+468))
									v85 = base.B2i32(int32(2) < v77)
									if int32(2) < v77 {
										v86 = int32(_a_F_timestamptz_in_3)
									} else {
										v86 = int32(_a_F_timestamptz_in_4)
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
									v107 = base.I32_div_s((v102+v77)*int32(_a_F_timestamptz_in_5), int32(256))
									v113 = base.I64_extend_i32_s(v80 + v87*int32(365) + v92 + v95 + v98 + v107 - int32(_a_F_timestamptz_in_6) - int32(_a_F_timestamptz_in_7))
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
									v154 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
									v155 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
									if v154 != v155>>(uint(int64(63))%64) {
										v192 = int32(0)
										v193 = F_errsave_start(m, v15)
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											if v193 == int32(0) {
												v242 = v192
												m.G0 = v13 + int32(512)
												return v242
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v199 = m.ExcPending
												if v199 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
													F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
														mBase = m.M
														v210 = m.ExcPending
														if v210 != 0 {
															return int32(0)
														} else {
															v242 = v192
															m.G0 = v13 + int32(512)
															return v242
														}
													}
												}
											}
										}
									} else {
										v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+456))
										v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+460))
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v13)+464))
										v162 = int32(60)
										v171 = base.I64_extend_i32_s(v159+(v160+v161*v162)*v162)*int64(1000000) + v60
										v174 = v155 + v171
										if base.B2i32(v171 < int64(0))^base.B2i32(v174 < v155) != 0 {
											v192 = int32(0)
											v193 = F_errsave_start(m, v15)
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int32(0)
											} else {
												if v193 == int32(0) {
													v242 = v192
													m.G0 = v13 + int32(512)
													return v242
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v199 = m.ExcPending
													if v199 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
														F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
															mBase = m.M
															v210 = m.ExcPending
															if v210 != 0 {
																return int32(0)
															} else {
																v242 = v192
																m.G0 = v13 + int32(512)
																return v242
															}
														}
													}
												}
											}
										} else {
											v178 = *(*int32)(unsafe.Add(mBase, uint32(v13)+452))
											v183 = base.I64_extend_i32_s(int32(0)-v178)*int64(-1000000) + v174
											*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v183
											if base.Ui64(v183+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
												mBase = m.M
												v238 = m.ExcPending
												if v238 != 0 {
													return int32(0)
												} else {
													v239 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
													v240 = F_Int64GetDatum(m, v239)
													mBase = m.M
													v241 = m.ExcPending
													if v241 != 0 {
														return int32(0)
													} else {
														v242 = v240
														m.G0 = v13 + int32(512)
														return v242
													}
												}
											} else {
												v192 = int32(0)
												v193 = F_errsave_start(m, v15)
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int32(0)
												} else {
													if v193 == int32(0) {
														v242 = v192
														m.G0 = v13 + int32(512)
														return v242
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v199 = m.ExcPending
														if v199 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
															F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(457), int32(_a_F_timestamptz_in_2))
																mBase = m.M
																v210 = m.ExcPending
																if v210 != 0 {
																	return int32(0)
																} else {
																	v242 = v192
																	m.G0 = v13 + int32(512)
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
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v215 = m.ExcPending
					if v215 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v17
						v217 = *(*int32)(unsafe.Add(mBase, uint32(v13)+448))
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v217
						F_errmsg_internal(m, int32(_a_F_timestamptz_in_10), v13)
						mBase = m.M
						v221 = m.ExcPending
						if v221 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamptz_in_1), int32(474), int32(_a_F_timestamptz_in_2))
							mBase = m.M
							v226 = m.ExcPending
							if v226 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 7:
					v229 = int64(-9223372036854775807 - 1)
					*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v229
					F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
					mBase = m.M
					v238 = m.ExcPending
					if v238 != 0 {
						return int32(0)
					} else {
						v239 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
						v240 = F_Int64GetDatum(m, v239)
						mBase = m.M
						v241 = m.ExcPending
						if v241 != 0 {
							return int32(0)
						} else {
							v242 = v240
							m.G0 = v13 + int32(512)
							return v242
						}
					}
				case 8:
					v229 = int64(9223372036854775807)
					*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v229
					F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
					mBase = m.M
					v238 = m.ExcPending
					if v238 != 0 {
						return int32(0)
					} else {
						v239 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
						v240 = F_Int64GetDatum(m, v239)
						mBase = m.M
						v241 = m.ExcPending
						if v241 != 0 {
							return int32(0)
						} else {
							v242 = v240
							m.G0 = v13 + int32(512)
							return v242
						}
					}
				case 9:
					v227 = F_SetEpochTimestamp(m)
					mBase = m.M
					v228 = m.ExcPending
					if v228 != 0 {
						return int32(0)
					} else {
						v229 = v227
						*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v229
						F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
						mBase = m.M
						v238 = m.ExcPending
						if v238 != 0 {
							return int32(0)
						} else {
							v239 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
							v240 = F_Int64GetDatum(m, v239)
							mBase = m.M
							v241 = m.ExcPending
							if v241 != 0 {
								return int32(0)
							} else {
								v242 = v240
								m.G0 = v13 + int32(512)
								return v242
							}
						}
					}
				}
			} else {
				v47 = v41
				F_DateTimeParseError(m, v47, v13+int32(56), v17, int32(_a_F_timestamptz_in_11), v15)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					v53 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v53)
					v242 = int32(0)
					m.G0 = v13 + int32(512)
					return v242
				}
			}
		}
	} else {
		v47 = v27
		F_DateTimeParseError(m, v47, v13+int32(56), v17, int32(_a_F_timestamptz_in_11), v15)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int32(0)
		} else {
			v53 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v53)
			v242 = int32(0)
			m.G0 = v13 + int32(512)
			return v242
		}
	}
}
func F_timestamptz_izone(m *base.Module, l0 int32) int32 {
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14001(m, l0, int32(_a_F_timestamptz_izone_0), int32(_a_F_timestamptz_izone_1), int32(_a_F_timestamptz_izone_2), int64(-1000000), int32(_a_F_timestamptz_izone_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		return v7
	}
}
func F_timestamptz_lt_timestamp(m *base.Module, l0 int32) int32 {
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int64)(unsafe.Add(mBase, uint32(v11)))
	v15 = F_timestamp2timestamptz_opt_overflow(m, v12, v7+int32(12))
	mBase = m.M
	v18 = m.ExcPending
	if v18 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(v7)+12))
		if int32(0) < v21 {
			v29 = base.B2i32(v10 != int64(9223372036854775807))
		} else {
			if v21 < int32(0) {
				v29 = base.B2i32(v10 == int64(-9223372036854775807-1))
			} else {
				v29 = base.B2i32(v10 < v15)
			}
		}
		m.G0 = v7 + int32(16)
		return v29
	}
}
func F_timestamptz_mi_interval_at_zone(m *base.Module, l0 int32) int32 {
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
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(256)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v13 = F_pg_detoast_datum_packed(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int32(0)
	} else {
		F_text_to_cstring_buffer(m, v13, v7, int32(256))
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			v20 = F_DecodeTimezoneNameToTz(m, v7)
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				F_interval_um_internal(m, v11, v7)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_timestamptz_pl_interval_internal(m, v10, v7, v20)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_Int64GetDatum(m, v24)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(256)
							return v26
						}
					}
				}
			}
		}
	}
}
func F_timestamptz_part_common(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v36 int32
	_ = v36
	var v39 int32
	_ = v39
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v78 int32
	_ = v78
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v89 int32
	_ = v89
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v120 int32
	_ = v120
	var v122 int32
	_ = v122
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v131 int32
	_ = v131
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v141 int64
	_ = v141
	var v142 int64
	_ = v142
	var v146 int32
	_ = v146
	var v147 int64
	_ = v147
	var v153 int32
	_ = v153
	var v154 int32
	_ = v154
	var v156 float64
	_ = v156
	var v158 int32
	_ = v158
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v166 int64
	_ = v166
	var v172 int32
	_ = v172
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int64
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v186 int32
	_ = v186
	var v187 int32
	_ = v187
	var v190 int32
	_ = v190
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v207 int32
	_ = v207
	var v211 int32
	_ = v211
	var v214 int32
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v221 int32
	_ = v221
	var v225 int32
	_ = v225
	var v228 int32
	_ = v228
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v234 int32
	_ = v234
	var v236 int32
	_ = v236
	var v240 int32
	_ = v240
	var v247 int32
	_ = v247
	var v251 int32
	_ = v251
	var v256 int32
	_ = v256
	var v258 int32
	_ = v258
	var v264 int32
	_ = v264
	var v269 int32
	_ = v269
	var v271 int32
	_ = v271
	var v277 int32
	_ = v277
	var v282 int32
	_ = v282
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v286 int32
	_ = v286
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v304 int32
	_ = v304
	var v308 int32
	_ = v308
	var v313 int32
	_ = v313
	var v316 int32
	_ = v316
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int64
	_ = v320
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v334 int32
	_ = v334
	var v335 int32
	_ = v335
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v340 int32
	_ = v340
	var v341 int32
	_ = v341
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v351 int32
	_ = v351
	var v352 int32
	_ = v352
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v366 int32
	_ = v366
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v386 int32
	_ = v386
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v389 int32
	_ = v389
	var v394 int32
	_ = v394
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v407 int32
	_ = v407
	var v413 int32
	_ = v413
	var v414 int32
	_ = v414
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v442 int32
	_ = v442
	var v448 int32
	_ = v448
	var v449 int32
	_ = v449
	var v454 int32
	_ = v454
	var v455 int64
	_ = v455
	var v457 int64
	_ = v457
	var v458 int32
	_ = v458
	var v461 int64
	_ = v461
	var v462 int32
	_ = v462
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v471 int32
	_ = v471
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v482 int32
	_ = v482
	var v486 int32
	_ = v486
	var v491 int32
	_ = v491
	var v495 int32
	_ = v495
	var v504 int32
	_ = v504
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v524 int32
	_ = v524
	var v535 int32
	_ = v535
	var v538 int32
	_ = v538
	var v540 int32
	_ = v540
	var v541 int32
	_ = v541
	var v548 int32
	_ = v548
	var v553 int32
	_ = v553
	var v554 int32
	_ = v554
	var v557 int64
	_ = v557
	var v558 int32
	_ = v558
	var v560 int64
	_ = v560
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v568 int32
	_ = v568
	var v569 int32
	_ = v569
	var v570 int32
	_ = v570
	var v571 int32
	_ = v571
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v577 int32
	_ = v577
	var v579 int32
	_ = v579
	var v580 int32
	_ = v580
	var v582 int32
	_ = v582
	var v583 int32
	_ = v583
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v592 float64
	_ = v592
	var v595 int32
	_ = v595
	var v596 int32
	_ = v596
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v605 int32
	_ = v605
	var v606 int32
	_ = v606
	var v613 int32
	_ = v613
	var v618 int32
	_ = v618
	var v622 int32
	_ = v622
	var v625 int32
	_ = v625
	var v627 int32
	_ = v627
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v640 int32
	_ = v640
	var v645 int64
	_ = v645
	var v646 int32
	_ = v646
	var v647 int32
	_ = v647
	var v649 int32
	_ = v649
	var v650 int32
	_ = v650
	var v656 int32
	_ = v656
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v671 int32
	_ = v671
	var v676 int32
	_ = v676
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int32(0)
	} else {
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
		v23 = int32(1)
		v24 = v17 + v23
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
		v29 = v27 & v23
		if v29 != 0 {
			v30 = v24
		} else {
			v30 = v17 + int32(4)
		}
		if v27 == int32(1) {
			v36 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
			if v36 == int32(18) {
				v39 = int32(16)
			} else {
				v39 = int32(0)
			}
			if base.Ui32((v36-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v46 = int32(4)
			} else {
				v46 = v39
			}
			v57 = v46
		} else {
			v47 = int32(1)
			if v29 != 0 {
				v57 = int32(base.Ui32(v27)>>(uint(v47)%32)) - v47
			} else {
				v51 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v57 = int32(base.Ui32(v51)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v59 = F_downcase_truncate_identifier(m, v30, v57, int32(0))
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int32(0)
		} else {
			v62 = v14 + int32(88)
			v66 = Fn13826(m, v59, v62, int32(_a_F_timestamptz_part_common_0), int32(_a_F_timestamptz_part_common_1), int32(_a_F_timestamptz_part_common_2))
			mBase = m.M
			if v66 == int32(31) {
				v72 = Fn13826(m, v59, v62, int32(_a_F_timestamptz_part_common_3), int32(_a_F_timestamptz_part_common_4), int32(_a_F_timestamptz_part_common_5))
				mBase = m.M
				v73 = v72
			} else {
				v73 = v66
			}
			if base.Ui64(int64(1)) < base.Ui64(v22-int64(9223372036854775807)) {
				if v73 != 0 {
					if v73 != int32(17) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v622 = m.ExcPending
						if v622 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v625 = m.ExcPending
							if v625 != 0 {
								return int32(0)
							} else {
								v627 = F_format_type_be(m, int32(1184))
								mBase = m.M
								v628 = m.ExcPending
								if v628 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v627
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = v59
									F_errmsg(m, int32(_a_F_timestamptz_part_common_6), v14)
									mBase = m.M
									v633 = m.ExcPending
									if v633 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_8), int32(_a_F_timestamptz_part_common_9))
										mBase = m.M
										v638 = m.ExcPending
										if v638 != 0 {
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
						v120 = int32(0)
						v122 = F_timestamp2tm(m, v22, v14+int32(92), v14+int32(40), v14+int32(84), v120, v120)
						mBase = m.M
						v123 = m.ExcPending
						if v123 != 0 {
							return int32(0)
						} else {
							if v122 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v664 = m.ExcPending
								if v664 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v667 = m.ExcPending
									if v667 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz_part_common_10), int32(0))
										mBase = m.M
										v671 = m.ExcPending
										if v671 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_11), int32(_a_F_timestamptz_part_common_9))
											mBase = m.M
											v676 = m.ExcPending
											if v676 != 0 {
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
								v124 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
								switch v124 - int32(4) {
								case 0:
									v640 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
									v645 = base.I64_extend_i32_s(int32(0) - v640)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								default:
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v535 = m.ExcPending
									if v535 != 0 {
										return int32(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v538 = m.ExcPending
										if v538 != 0 {
											return int32(0)
										} else {
											v540 = F_format_type_be(m, int32(1184))
											mBase = m.M
											v541 = m.ExcPending
											if v541 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v540
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v59
												F_errmsg(m, int32(_a_F_timestamptz_part_common_12), v14+int32(16))
												mBase = m.M
												v548 = m.ExcPending
												if v548 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_13), int32(_a_F_timestamptz_part_common_9))
													mBase = m.M
													v553 = m.ExcPending
													if v553 != 0 {
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
								case 14:
									v165 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
									if l1 != 0 {
										v166 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
										v172 = F_int64_div_fast_to_numeric(m, v166+base.I64_extend_i32_s(v165)*int64(1000000), int32(6))
										mBase = m.M
										v173 = m.ExcPending
										if v173 != 0 {
											return int32(0)
										} else {
											v656 = v172
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v174 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v180 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v174), float64(1e+06)), base.F64_convert_i32_s(v165)))
										mBase = m.M
										v181 = m.ExcPending
										if v181 != 0 {
											return int32(0)
										} else {
											v656 = v180
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 15:
									v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+44)))
									v645 = v182
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 16:
									v183 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+48)))
									v645 = v183
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 17:
									v184 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+52)))
									v645 = v184
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 18:
									v194 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v195 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v198 = F_date2j(m, v194, v195, v196)
									mBase = m.M
									v199 = int32(1)
									v201 = F_date2j(m, v194, v199, int32(4))
									mBase = m.M
									v204 = F_j2day(m, v201-v199)
									mBase = m.M
									if v198 < v201-v204 {
										v207 = int32(1)
										v211 = F_date2j(m, v194-v207, v207, int32(4))
										mBase = m.M
										v214 = F_j2day(m, v211-v207)
										mBase = m.M
										v215 = v211
										v216 = v214
									} else {
										v215 = v201
										v216 = v204
									}
									v218 = v216 - v215 + v198
									if int32(357) <= v218 {
										v221 = int32(1)
										v225 = F_date2j(m, v194+v221, v221, int32(4))
										mBase = m.M
										v228 = F_j2day(m, v225-v221)
										mBase = m.M
										v229 = v225 - v228
										if v198 < v229 {
											v232 = v218
										} else {
											v232 = v198 - v229
										}
										v234 = v232
									} else {
										v234 = v218
									}
									v236 = base.I32_div_s(v234, int32(7))
									v645 = base.I64_extend_i32_s(v236 + int32(1))
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 19:
									v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+56)))
									v645 = v185
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 20:
									v186 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v187 = int32(1)
									v190 = base.I32_div_s(v186-v187, int32(3))
									v645 = base.I64_extend_i32_s(v190 + v187)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 21:
									v240 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if int32(0) < v240 {
										v645 = base.I64_extend_i32_u(v240)
									} else {
										v645 = base.I64_extend_i32_s(v240 - int32(1))
									}
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 22:
									v247 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if int32(0) < v247 {
										v251 = base.I32_div_u_s(v247, int32(10))
										v645 = base.I64_extend_i32_u(v251)
									} else {
										v256 = base.I32_div_s(int32(9)-v247, int32(-10))
										v645 = base.I64_extend_i32_s(v256)
									}
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 23:
									v258 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if int32(0) < v258 {
										v264 = base.I32_div_s(v258+int32(99), int32(100))
										v645 = base.I64_extend_i32_s(v264)
									} else {
										v269 = base.I32_div_s(int32(100)-v258, int32(-100))
										v645 = base.I64_extend_i32_s(v269)
									}
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 24:
									v271 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if int32(0) < v271 {
										v277 = base.I32_div_s(v271+int32(999), int32(1000))
										v645 = base.I64_extend_i32_s(v277)
									} else {
										v282 = base.I32_div_s(int32(1000)-v271, int32(-1000))
										v645 = base.I64_extend_i32_s(v282)
									}
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 25:
									v146 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
									if l1 != 0 {
										v147 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
										v153 = F_int64_div_fast_to_numeric(m, v147+base.I64_extend_i32_s(v146)*int64(1000000), int32(3))
										mBase = m.M
										v154 = m.ExcPending
										if v154 != 0 {
											return int32(0)
										} else {
											v656 = v153
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v156 = float64(1000)
										v158 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v163 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v146), v156), base.F64_div(base.F64_convert_i32_s(v158), v156)))
										mBase = m.M
										v164 = m.ExcPending
										if v164 != 0 {
											return int32(0)
										} else {
											v656 = v163
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 26:
									v141 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
									v142 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+40)))
									v645 = v141 + v142*int64(1000000)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 27:
									v284 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v285 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v286 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v291 = base.B2i32(int32(2) < v285)
									if int32(2) < v285 {
										v292 = int32(_a_F_timestamptz_part_common_14)
									} else {
										v292 = int32(_a_F_timestamptz_part_common_15)
									}
									v293 = v292 + v284
									v298 = base.I32_div_s(v293, int32(4))
									v301 = base.I32_div_s(v293, int32(-100))
									v304 = base.I32_div_s(v293, int32(400))
									if int32(2) < v285 {
										v308 = int32(1)
									} else {
										v308 = int32(13)
									}
									v313 = base.I32_div_s((v308+v285)*int32(_a_F_timestamptz_part_common_16), int32(256))
									v316 = v286 + v293*int32(365) + v298 + v301 + v304 + v313 - int32(_a_F_timestamptz_part_common_17)
									if l1 != 0 {
										v318 = F_int64_to_numeric(m, base.I64_extend_i32_s(v316))
										mBase = m.M
										v319 = m.ExcPending
										if v319 != 0 {
											return int32(0)
										} else {
											v320 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
											v321 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
											v322 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
											v323 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
											v324 = int32(60)
											v334 = F_int64_to_numeric(m, v320+base.I64_extend_i32_s(v321+(v322+v323*v324)*v324)*int64(1000000))
											mBase = m.M
											v335 = m.ExcPending
											if v335 != 0 {
												return int32(0)
											} else {
												v337 = F_int64_to_numeric(m, int64(86400000000))
												mBase = m.M
												v338 = m.ExcPending
												if v338 != 0 {
													return int32(0)
												} else {
													v340 = F_numeric_div_opt_error(m, v334, v337, int32(0))
													mBase = m.M
													v341 = m.ExcPending
													if v341 != 0 {
														return int32(0)
													} else {
														v343 = F_numeric_add_opt_error(m, v318, v340, int32(0))
														mBase = m.M
														v344 = m.ExcPending
														if v344 != 0 {
															return int32(0)
														} else {
															v656 = v343
															m.G0 = v14 + int32(96)
															return v656
														}
													}
												}
											}
										}
									} else {
										v345 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v349 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
										v350 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
										v351 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
										v352 = int32(60)
										v364 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v345), float64(1e+06)), base.F64_convert_i32_s(v349+(v350+v351*v352)*v352)), float64(86400)), base.F64_convert_i32_s(v316)))
										mBase = m.M
										v365 = m.ExcPending
										if v365 != 0 {
											return int32(0)
										} else {
											v656 = v364
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 28, 33:
									v413 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v414 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v415 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v420 = base.B2i32(int32(2) < v414)
									if int32(2) < v414 {
										v421 = int32(_a_F_timestamptz_part_common_14)
									} else {
										v421 = int32(_a_F_timestamptz_part_common_15)
									}
									v422 = v421 + v413
									v427 = base.I32_div_s(v422, int32(4))
									v430 = base.I32_div_s(v422, int32(-100))
									v433 = base.I32_div_s(v422, int32(400))
									if int32(2) < v414 {
										v437 = int32(1)
									} else {
										v437 = int32(13)
									}
									v442 = base.I32_div_s((v437+v414)*int32(_a_F_timestamptz_part_common_16), int32(256))
									v448 = int32(7)
									v449 = base.I32_rem_s(v415+v422*int32(365)+v427+v430+v433+v442-int32(_a_F_timestamptz_part_common_17)+int32(1), v448)
									if v449 < int32(0) {
										v454 = v449 + v448
									} else {
										v454 = v449
									}
									v455 = base.I64_extend_i32_s(v454)
									if v454 != 0 {
										v457 = v455
									} else {
										v457 = int64(7)
									}
									v458 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
									if v458 == int32(37) {
										v461 = v457
									} else {
										v461 = v455
									}
									v645 = v461
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 29:
									v462 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v463 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v464 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v469 = base.B2i32(int32(2) < v463)
									if int32(2) < v463 {
										v470 = int32(_a_F_timestamptz_part_common_14)
									} else {
										v470 = int32(_a_F_timestamptz_part_common_15)
									}
									v471 = v470 + v462
									v476 = base.I32_div_s(v471, int32(4))
									v479 = base.I32_div_s(v471, int32(-100))
									v482 = base.I32_div_s(v471, int32(400))
									if int32(2) < v463 {
										v486 = int32(1)
									} else {
										v486 = int32(13)
									}
									v491 = base.I32_div_s((v486+v463)*int32(_a_F_timestamptz_part_common_16), int32(256))
									v495 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v504 = int32(_a_F_timestamptz_part_common_15) + v495
									v509 = base.I32_div_s(v504, int32(4))
									v512 = base.I32_div_s(v504, int32(-100))
									v515 = base.I32_div_s(v504, int32(400))
									v524 = base.I32_div_s(int32(_a_F_timestamptz_part_common_18), int32(256))
									v645 = base.I64_extend_i32_s(v464 + v471*int32(365) + v476 + v479 + v482 + v491 - int32(_a_F_timestamptz_part_common_17) - (int32(1) + v504*int32(365) + v509 + v512 + v515 + v524 - int32(_a_F_timestamptz_part_common_17)) + int32(1))
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 30:
									v136 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
									v139 = base.I32_div_s(int32(0)-v136, int32(3600))
									v645 = base.I64_extend_i32_s(v139)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 31:
									v128 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
									v130 = int32(60)
									v131 = base.I32_div_s(int32(0)-v128, v130)
									v133 = base.I32_rem_s(v131, v130)
									v645 = base.I64_extend_i32_s(v133)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								case 32:
									v366 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v367 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v368 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v370 = F_date2j(m, v366, v367, v368)
									mBase = m.M
									v371 = int32(1)
									v373 = F_date2j(m, v366, v371, int32(4))
									mBase = m.M
									v376 = F_j2day(m, v373-v371)
									mBase = m.M
									if v370 < v373-v376 {
										v379 = int32(1)
										v380 = v366 - v379
										v383 = F_date2j(m, v380, v379, int32(4))
										mBase = m.M
										v386 = F_j2day(m, v383-v379)
										mBase = m.M
										v387 = v380
										v388 = v383
										v389 = v386
									} else {
										v387 = v366
										v388 = v373
										v389 = v376
									}
									if int32(357) <= v389+v370-v388 {
										v394 = int32(1)
										v395 = v387 + v394
										v398 = F_date2j(m, v395, v394, int32(4))
										mBase = m.M
										v401 = F_j2day(m, v398-v394)
										mBase = m.M
										if v370 < v398-v401 {
											v404 = v387
										} else {
											v404 = v395
										}
										v407 = v404
									} else {
										v407 = v387
									}
									v645 = base.I64_extend_i32_s(v407) - base.I64_extend_i32_u(base.B2i32(v407 <= int32(0)))
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int32(0)
										} else {
											v656 = v646
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
										mBase = m.M
										v650 = m.ExcPending
										if v650 != 0 {
											return int32(0)
										} else {
											v656 = v649
											m.G0 = v14 + int32(96)
											return v656
										}
									}
								}
							}
						}
					}
				} else {
					v554 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
					if v554 == int32(11) {
						v557 = F_SetEpochTimestamp(m)
						mBase = m.M
						v558 = m.ExcPending
						if v558 != 0 {
							return int32(0)
						} else {
							v560 = v557 + int64(9223372036854775807)
							if l1 != 0 {
								if v22 < v560 {
									v564 = F_int64_div_fast_to_numeric(m, v22-v557, int32(6))
									mBase = m.M
									v565 = m.ExcPending
									if v565 != 0 {
										return int32(0)
									} else {
										v656 = v564
										m.G0 = v14 + int32(96)
										return v656
									}
								} else {
									v568 = F_int64_to_numeric(m, v22)
									mBase = m.M
									v569 = m.ExcPending
									if v569 != 0 {
										return int32(0)
									} else {
										v570 = F_int64_to_numeric(m, v557)
										mBase = m.M
										v571 = m.ExcPending
										if v571 != 0 {
											return int32(0)
										} else {
											v573 = F_numeric_sub_opt_error(m, v568, v570, int32(0))
											mBase = m.M
											v574 = m.ExcPending
											if v574 != 0 {
												return int32(0)
											} else {
												v576 = F_int64_to_numeric(m, int64(1000000))
												mBase = m.M
												v577 = m.ExcPending
												if v577 != 0 {
													return int32(0)
												} else {
													v579 = F_numeric_div_opt_error(m, v573, v576, int32(0))
													mBase = m.M
													v580 = m.ExcPending
													if v580 != 0 {
														return int32(0)
													} else {
														v582 = F_DirectFunctionCall2Coll(m, int32(1259), int32(0), v579, int32(6))
														mBase = m.M
														v583 = m.ExcPending
														if v583 != 0 {
															return int32(0)
														} else {
															v584 = F_pg_detoast_datum(m, v582)
															mBase = m.M
															v585 = m.ExcPending
															if v585 != 0 {
																return int32(0)
															} else {
																v656 = v584
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								if v22 < v560 {
									v592 = base.F64_convert_i64_s(v22 - v557)
								} else {
									v592 = base.F64_sub(base.F64_convert_i64_s(v22), base.F64_convert_i64_s(v557))
								}
								v595 = F_Float8GetDatum(m, base.F64_div(v592, float64(1e+06)))
								mBase = m.M
								v596 = m.ExcPending
								if v596 != 0 {
									return int32(0)
								} else {
									v656 = v595
									m.G0 = v14 + int32(96)
									return v656
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v600 = m.ExcPending
						if v600 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v603 = m.ExcPending
							if v603 != 0 {
								return int32(0)
							} else {
								v605 = F_format_type_be(m, int32(1184))
								mBase = m.M
								v606 = m.ExcPending
								if v606 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v605
									*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v59
									F_errmsg(m, int32(_a_F_timestamptz_part_common_12), v14+int32(32))
									mBase = m.M
									v613 = m.ExcPending
									if v613 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_19), int32(_a_F_timestamptz_part_common_9))
										mBase = m.M
										v618 = m.ExcPending
										if v618 != 0 {
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
				v78 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
				v82 = F_NonFiniteTimestampTzPart(m, v73, v78, v59, base.B2i32(v22 == int64(-9223372036854775807-1)), int32(1))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					if base.F64_ne(v82, float64(0)) != 0 {
						if l1 != 0 {
							if base.F64_lt(v82, float64(0)) != 0 {
								v89 = int32(0)
								v93 = F_DirectFunctionCall3Coll(m, int32(408), v89, int32(_a_F_timestamptz_part_common_20), v89, int32(-1))
								mBase = m.M
								v94 = m.ExcPending
								if v94 != 0 {
									return int32(0)
								} else {
									v656 = v93
									m.G0 = v14 + int32(96)
									return v656
								}
							} else {
								if base.F64_gt(v82, float64(0)) == int32(0) {
									if v73 != 0 {
										if v73 != int32(17) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v622 = m.ExcPending
											if v622 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v625 = m.ExcPending
												if v625 != 0 {
													return int32(0)
												} else {
													v627 = F_format_type_be(m, int32(1184))
													mBase = m.M
													v628 = m.ExcPending
													if v628 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v627
														*(*int32)(unsafe.Add(mBase, uint32(v14))) = v59
														F_errmsg(m, int32(_a_F_timestamptz_part_common_6), v14)
														mBase = m.M
														v633 = m.ExcPending
														if v633 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_8), int32(_a_F_timestamptz_part_common_9))
															mBase = m.M
															v638 = m.ExcPending
															if v638 != 0 {
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
											v120 = int32(0)
											v122 = F_timestamp2tm(m, v22, v14+int32(92), v14+int32(40), v14+int32(84), v120, v120)
											mBase = m.M
											v123 = m.ExcPending
											if v123 != 0 {
												return int32(0)
											} else {
												if v122 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v664 = m.ExcPending
													if v664 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v667 = m.ExcPending
														if v667 != 0 {
															return int32(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamptz_part_common_10), int32(0))
															mBase = m.M
															v671 = m.ExcPending
															if v671 != 0 {
																return int32(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_11), int32(_a_F_timestamptz_part_common_9))
																mBase = m.M
																v676 = m.ExcPending
																if v676 != 0 {
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
													v124 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
													switch v124 - int32(4) {
													case 0:
														v640 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
														v645 = base.I64_extend_i32_s(int32(0) - v640)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v535 = m.ExcPending
														if v535 != 0 {
															return int32(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v538 = m.ExcPending
															if v538 != 0 {
																return int32(0)
															} else {
																v540 = F_format_type_be(m, int32(1184))
																mBase = m.M
																v541 = m.ExcPending
																if v541 != 0 {
																	return int32(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v540
																	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v59
																	F_errmsg(m, int32(_a_F_timestamptz_part_common_12), v14+int32(16))
																	mBase = m.M
																	v548 = m.ExcPending
																	if v548 != 0 {
																		return int32(0)
																	} else {
																		F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_13), int32(_a_F_timestamptz_part_common_9))
																		mBase = m.M
																		v553 = m.ExcPending
																		if v553 != 0 {
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
													case 14:
														v165 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
														if l1 != 0 {
															v166 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
															v172 = F_int64_div_fast_to_numeric(m, v166+base.I64_extend_i32_s(v165)*int64(1000000), int32(6))
															mBase = m.M
															v173 = m.ExcPending
															if v173 != 0 {
																return int32(0)
															} else {
																v656 = v172
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v174 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
															v180 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v174), float64(1e+06)), base.F64_convert_i32_s(v165)))
															mBase = m.M
															v181 = m.ExcPending
															if v181 != 0 {
																return int32(0)
															} else {
																v656 = v180
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 15:
														v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+44)))
														v645 = v182
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 16:
														v183 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+48)))
														v645 = v183
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 17:
														v184 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+52)))
														v645 = v184
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 18:
														v194 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v195 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v196 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v198 = F_date2j(m, v194, v195, v196)
														mBase = m.M
														v199 = int32(1)
														v201 = F_date2j(m, v194, v199, int32(4))
														mBase = m.M
														v204 = F_j2day(m, v201-v199)
														mBase = m.M
														if v198 < v201-v204 {
															v207 = int32(1)
															v211 = F_date2j(m, v194-v207, v207, int32(4))
															mBase = m.M
															v214 = F_j2day(m, v211-v207)
															mBase = m.M
															v215 = v211
															v216 = v214
														} else {
															v215 = v201
															v216 = v204
														}
														v218 = v216 - v215 + v198
														if int32(357) <= v218 {
															v221 = int32(1)
															v225 = F_date2j(m, v194+v221, v221, int32(4))
															mBase = m.M
															v228 = F_j2day(m, v225-v221)
															mBase = m.M
															v229 = v225 - v228
															if v198 < v229 {
																v232 = v218
															} else {
																v232 = v198 - v229
															}
															v234 = v232
														} else {
															v234 = v218
														}
														v236 = base.I32_div_s(v234, int32(7))
														v645 = base.I64_extend_i32_s(v236 + int32(1))
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 19:
														v185 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+56)))
														v645 = v185
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 20:
														v186 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v187 = int32(1)
														v190 = base.I32_div_s(v186-v187, int32(3))
														v645 = base.I64_extend_i32_s(v190 + v187)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 21:
														v240 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														if int32(0) < v240 {
															v645 = base.I64_extend_i32_u(v240)
														} else {
															v645 = base.I64_extend_i32_s(v240 - int32(1))
														}
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 22:
														v247 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														if int32(0) < v247 {
															v251 = base.I32_div_u_s(v247, int32(10))
															v645 = base.I64_extend_i32_u(v251)
														} else {
															v256 = base.I32_div_s(int32(9)-v247, int32(-10))
															v645 = base.I64_extend_i32_s(v256)
														}
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 23:
														v258 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														if int32(0) < v258 {
															v264 = base.I32_div_s(v258+int32(99), int32(100))
															v645 = base.I64_extend_i32_s(v264)
														} else {
															v269 = base.I32_div_s(int32(100)-v258, int32(-100))
															v645 = base.I64_extend_i32_s(v269)
														}
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 24:
														v271 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														if int32(0) < v271 {
															v277 = base.I32_div_s(v271+int32(999), int32(1000))
															v645 = base.I64_extend_i32_s(v277)
														} else {
															v282 = base.I32_div_s(int32(1000)-v271, int32(-1000))
															v645 = base.I64_extend_i32_s(v282)
														}
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 25:
														v146 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
														if l1 != 0 {
															v147 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
															v153 = F_int64_div_fast_to_numeric(m, v147+base.I64_extend_i32_s(v146)*int64(1000000), int32(3))
															mBase = m.M
															v154 = m.ExcPending
															if v154 != 0 {
																return int32(0)
															} else {
																v656 = v153
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v156 = float64(1000)
															v158 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
															v163 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v146), v156), base.F64_div(base.F64_convert_i32_s(v158), v156)))
															mBase = m.M
															v164 = m.ExcPending
															if v164 != 0 {
																return int32(0)
															} else {
																v656 = v163
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 26:
														v141 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
														v142 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+40)))
														v645 = v141 + v142*int64(1000000)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 27:
														v284 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v285 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v286 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v291 = base.B2i32(int32(2) < v285)
														if int32(2) < v285 {
															v292 = int32(_a_F_timestamptz_part_common_14)
														} else {
															v292 = int32(_a_F_timestamptz_part_common_15)
														}
														v293 = v292 + v284
														v298 = base.I32_div_s(v293, int32(4))
														v301 = base.I32_div_s(v293, int32(-100))
														v304 = base.I32_div_s(v293, int32(400))
														if int32(2) < v285 {
															v308 = int32(1)
														} else {
															v308 = int32(13)
														}
														v313 = base.I32_div_s((v308+v285)*int32(_a_F_timestamptz_part_common_16), int32(256))
														v316 = v286 + v293*int32(365) + v298 + v301 + v304 + v313 - int32(_a_F_timestamptz_part_common_17)
														if l1 != 0 {
															v318 = F_int64_to_numeric(m, base.I64_extend_i32_s(v316))
															mBase = m.M
															v319 = m.ExcPending
															if v319 != 0 {
																return int32(0)
															} else {
																v320 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
																v321 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
																v322 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
																v323 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
																v324 = int32(60)
																v334 = F_int64_to_numeric(m, v320+base.I64_extend_i32_s(v321+(v322+v323*v324)*v324)*int64(1000000))
																mBase = m.M
																v335 = m.ExcPending
																if v335 != 0 {
																	return int32(0)
																} else {
																	v337 = F_int64_to_numeric(m, int64(86400000000))
																	mBase = m.M
																	v338 = m.ExcPending
																	if v338 != 0 {
																		return int32(0)
																	} else {
																		v340 = F_numeric_div_opt_error(m, v334, v337, int32(0))
																		mBase = m.M
																		v341 = m.ExcPending
																		if v341 != 0 {
																			return int32(0)
																		} else {
																			v343 = F_numeric_add_opt_error(m, v318, v340, int32(0))
																			mBase = m.M
																			v344 = m.ExcPending
																			if v344 != 0 {
																				return int32(0)
																			} else {
																				v656 = v343
																				m.G0 = v14 + int32(96)
																				return v656
																			}
																		}
																	}
																}
															}
														} else {
															v345 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
															v349 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
															v350 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
															v351 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
															v352 = int32(60)
															v364 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v345), float64(1e+06)), base.F64_convert_i32_s(v349+(v350+v351*v352)*v352)), float64(86400)), base.F64_convert_i32_s(v316)))
															mBase = m.M
															v365 = m.ExcPending
															if v365 != 0 {
																return int32(0)
															} else {
																v656 = v364
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 28, 33:
														v413 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v414 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v415 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v420 = base.B2i32(int32(2) < v414)
														if int32(2) < v414 {
															v421 = int32(_a_F_timestamptz_part_common_14)
														} else {
															v421 = int32(_a_F_timestamptz_part_common_15)
														}
														v422 = v421 + v413
														v427 = base.I32_div_s(v422, int32(4))
														v430 = base.I32_div_s(v422, int32(-100))
														v433 = base.I32_div_s(v422, int32(400))
														if int32(2) < v414 {
															v437 = int32(1)
														} else {
															v437 = int32(13)
														}
														v442 = base.I32_div_s((v437+v414)*int32(_a_F_timestamptz_part_common_16), int32(256))
														v448 = int32(7)
														v449 = base.I32_rem_s(v415+v422*int32(365)+v427+v430+v433+v442-int32(_a_F_timestamptz_part_common_17)+int32(1), v448)
														if v449 < int32(0) {
															v454 = v449 + v448
														} else {
															v454 = v449
														}
														v455 = base.I64_extend_i32_s(v454)
														if v454 != 0 {
															v457 = v455
														} else {
															v457 = int64(7)
														}
														v458 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
														if v458 == int32(37) {
															v461 = v457
														} else {
															v461 = v455
														}
														v645 = v461
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 29:
														v462 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v463 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v464 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v469 = base.B2i32(int32(2) < v463)
														if int32(2) < v463 {
															v470 = int32(_a_F_timestamptz_part_common_14)
														} else {
															v470 = int32(_a_F_timestamptz_part_common_15)
														}
														v471 = v470 + v462
														v476 = base.I32_div_s(v471, int32(4))
														v479 = base.I32_div_s(v471, int32(-100))
														v482 = base.I32_div_s(v471, int32(400))
														if int32(2) < v463 {
															v486 = int32(1)
														} else {
															v486 = int32(13)
														}
														v491 = base.I32_div_s((v486+v463)*int32(_a_F_timestamptz_part_common_16), int32(256))
														v495 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v504 = int32(_a_F_timestamptz_part_common_15) + v495
														v509 = base.I32_div_s(v504, int32(4))
														v512 = base.I32_div_s(v504, int32(-100))
														v515 = base.I32_div_s(v504, int32(400))
														v524 = base.I32_div_s(int32(_a_F_timestamptz_part_common_18), int32(256))
														v645 = base.I64_extend_i32_s(v464 + v471*int32(365) + v476 + v479 + v482 + v491 - int32(_a_F_timestamptz_part_common_17) - (int32(1) + v504*int32(365) + v509 + v512 + v515 + v524 - int32(_a_F_timestamptz_part_common_17)) + int32(1))
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 30:
														v136 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
														v139 = base.I32_div_s(int32(0)-v136, int32(3600))
														v645 = base.I64_extend_i32_s(v139)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 31:
														v128 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
														v130 = int32(60)
														v131 = base.I32_div_s(int32(0)-v128, v130)
														v133 = base.I32_rem_s(v131, v130)
														v645 = base.I64_extend_i32_s(v133)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													case 32:
														v366 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v367 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v368 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v370 = F_date2j(m, v366, v367, v368)
														mBase = m.M
														v371 = int32(1)
														v373 = F_date2j(m, v366, v371, int32(4))
														mBase = m.M
														v376 = F_j2day(m, v373-v371)
														mBase = m.M
														if v370 < v373-v376 {
															v379 = int32(1)
															v380 = v366 - v379
															v383 = F_date2j(m, v380, v379, int32(4))
															mBase = m.M
															v386 = F_j2day(m, v383-v379)
															mBase = m.M
															v387 = v380
															v388 = v383
															v389 = v386
														} else {
															v387 = v366
															v388 = v373
															v389 = v376
														}
														if int32(357) <= v389+v370-v388 {
															v394 = int32(1)
															v395 = v387 + v394
															v398 = F_date2j(m, v395, v394, int32(4))
															mBase = m.M
															v401 = F_j2day(m, v398-v394)
															mBase = m.M
															if v370 < v398-v401 {
																v404 = v387
															} else {
																v404 = v395
															}
															v407 = v404
														} else {
															v407 = v387
														}
														v645 = base.I64_extend_i32_s(v407) - base.I64_extend_i32_u(base.B2i32(v407 <= int32(0)))
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int32(0)
															} else {
																v656 = v646
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v649 = F_Float8GetDatum(m, base.F64_convert_i64_s(v645))
															mBase = m.M
															v650 = m.ExcPending
															if v650 != 0 {
																return int32(0)
															} else {
																v656 = v649
																m.G0 = v14 + int32(96)
																return v656
															}
														}
													}
												}
											}
										}
									} else {
										v554 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
										if v554 == int32(11) {
											v557 = F_SetEpochTimestamp(m)
											mBase = m.M
											v558 = m.ExcPending
											if v558 != 0 {
												return int32(0)
											} else {
												v560 = v557 + int64(9223372036854775807)
												if l1 != 0 {
													if v22 < v560 {
														v564 = F_int64_div_fast_to_numeric(m, v22-v557, int32(6))
														mBase = m.M
														v565 = m.ExcPending
														if v565 != 0 {
															return int32(0)
														} else {
															v656 = v564
															m.G0 = v14 + int32(96)
															return v656
														}
													} else {
														v568 = F_int64_to_numeric(m, v22)
														mBase = m.M
														v569 = m.ExcPending
														if v569 != 0 {
															return int32(0)
														} else {
															v570 = F_int64_to_numeric(m, v557)
															mBase = m.M
															v571 = m.ExcPending
															if v571 != 0 {
																return int32(0)
															} else {
																v573 = F_numeric_sub_opt_error(m, v568, v570, int32(0))
																mBase = m.M
																v574 = m.ExcPending
																if v574 != 0 {
																	return int32(0)
																} else {
																	v576 = F_int64_to_numeric(m, int64(1000000))
																	mBase = m.M
																	v577 = m.ExcPending
																	if v577 != 0 {
																		return int32(0)
																	} else {
																		v579 = F_numeric_div_opt_error(m, v573, v576, int32(0))
																		mBase = m.M
																		v580 = m.ExcPending
																		if v580 != 0 {
																			return int32(0)
																		} else {
																			v582 = F_DirectFunctionCall2Coll(m, int32(1259), int32(0), v579, int32(6))
																			mBase = m.M
																			v583 = m.ExcPending
																			if v583 != 0 {
																				return int32(0)
																			} else {
																				v584 = F_pg_detoast_datum(m, v582)
																				mBase = m.M
																				v585 = m.ExcPending
																				if v585 != 0 {
																					return int32(0)
																				} else {
																					v656 = v584
																					m.G0 = v14 + int32(96)
																					return v656
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													if v22 < v560 {
														v592 = base.F64_convert_i64_s(v22 - v557)
													} else {
														v592 = base.F64_sub(base.F64_convert_i64_s(v22), base.F64_convert_i64_s(v557))
													}
													v595 = F_Float8GetDatum(m, base.F64_div(v592, float64(1e+06)))
													mBase = m.M
													v596 = m.ExcPending
													if v596 != 0 {
														return int32(0)
													} else {
														v656 = v595
														m.G0 = v14 + int32(96)
														return v656
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v600 = m.ExcPending
											if v600 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v603 = m.ExcPending
												if v603 != 0 {
													return int32(0)
												} else {
													v605 = F_format_type_be(m, int32(1184))
													mBase = m.M
													v606 = m.ExcPending
													if v606 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v605
														*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v59
														F_errmsg(m, int32(_a_F_timestamptz_part_common_12), v14+int32(32))
														mBase = m.M
														v613 = m.ExcPending
														if v613 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_19), int32(_a_F_timestamptz_part_common_9))
															mBase = m.M
															v618 = m.ExcPending
															if v618 != 0 {
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
									v100 = int32(0)
									v104 = F_DirectFunctionCall3Coll(m, int32(408), v100, int32(_a_F_timestamptz_part_common_21), v100, int32(-1))
									mBase = m.M
									v105 = m.ExcPending
									if v105 != 0 {
										return int32(0)
									} else {
										v656 = v104
										m.G0 = v14 + int32(96)
										return v656
									}
								}
							}
						} else {
							v106 = F_Float8GetDatum(m, v82)
							mBase = m.M
							v107 = m.ExcPending
							if v107 != 0 {
								return int32(0)
							} else {
								v656 = v106
								m.G0 = v14 + int32(96)
								return v656
							}
						}
					} else {
						v108 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v108)
						v656 = int32(0)
						m.G0 = v14 + int32(96)
						return v656
					}
				}
			}
		}
	}
}
func F_timestamptz_to_str(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v4 int32
	_ = v4
	var v6 int32
	_ = v6
	var v16 int32
	_ = v16
	var v20 int64
	_ = v20
	var v24 int32
	_ = v24
	var v28 int64
	_ = v28
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	var v42 int32
	_ = v42
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v55 int32
	_ = v55
	var v59 int64
	_ = v59
	var v63 int64
	_ = v63
	var v67 int64
	_ = v67
	v4 = m.G0
	v6 = v4 + int32(-64)
	m.G0 = v6
	if base.Ui64(l0-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		if l0 == int64(-9223372036854775807-1) {
			v16 = int32(*(*uint16)(unsafe.Add(mBase, _c_F_timestamptz_to_str[0])))
			*(*uint16)(unsafe.Add(mBase, _c_F_timestamptz_to_str[1])) = uint16(v16)
			v20 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[2]))
			*(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[3])) = v20
		} else {
			v24 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_timestamptz_to_str[4])))
			*(*uint8)(unsafe.Add(mBase, _c_F_timestamptz_to_str[1])) = uint8(v24)
			v28 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[5]))
			*(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[3])) = v28
		}
		m.G0 = v6 - int32(-64)
		return int32(_a_F_timestamptz_to_str_0)
	} else {
		v33 = v4 + int32(-48)
		v39 = F_timestamp2tm(m, l0, v4+int32(-4), v33, v4+int32(-52), v4+int32(-56), int32(0))
		mBase = m.M
		v42 = m.ExcPending
		if v42 != 0 {
			return int32(0)
		} else {
			if v39 == int32(0) {
				v45 = *(*int32)(unsafe.Add(mBase, uint32(v6)+12))
				v46 = int32(1)
				v47 = *(*int32)(unsafe.Add(mBase, uint32(v6)+60))
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
				F_EncodeDateTime(m, v33, v45, v46, v47, v48, v46, int32(_a_F_timestamptz_to_str_0))
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int32(0)
				} else {
					m.G0 = v6 - int32(-64)
					return int32(_a_F_timestamptz_to_str_0)
				}
			} else {
				v55 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_timestamptz_to_str[6])))
				*(*uint8)(unsafe.Add(mBase, _c_F_timestamptz_to_str[7])) = uint8(v55)
				v59 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[8]))
				*(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[9])) = v59
				v63 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[10]))
				*(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[1])) = v63
				v67 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[11]))
				*(*int64)(unsafe.Add(mBase, _c_F_timestamptz_to_str[3])) = v67
				m.G0 = v6 - int32(-64)
				return int32(_a_F_timestamptz_to_str_0)
			}
		}
	}
}
func F_timestamptz_trunc_zone(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	v5 = m.G0
	v7 = v5 - int32(256)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int32(0)
	} else {
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v17 = F_pg_detoast_datum_packed(m, v16)
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int32(0)
		} else {
			F_text_to_cstring_buffer(m, v17, v7, int32(256))
			mBase = m.M
			v21 = m.ExcPending
			if v21 != 0 {
				return int32(0)
			} else {
				v22 = F_DecodeTimezoneNameToTz(m, v7)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_timestamptz_trunc_internal(m, v10, v15, v22)
					mBase = m.M
					v25 = m.ExcPending
					if v25 != 0 {
						return int32(0)
					} else {
						v26 = F_Int64GetDatum(m, v24)
						mBase = m.M
						v27 = m.ExcPending
						if v27 != 0 {
							return int32(0)
						} else {
							m.G0 = v7 + int32(256)
							return v26
						}
					}
				}
			}
		}
	}
}
