package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_extract_timestamptz(m *base.Module, l0 int32) int64 {
	var v3 int64
	_ = v3
	var v6 int32
	_ = v6
	v3 = F_timestamptz_part_common(m, l0, int32(1))
	v6 = m.ExcPending
	if v6 != 0 {
		return int64(0)
	} else {
		return v3
	}
}
func F_timestamptz_bin(m *base.Module, l0 int32) int64 {
	var v10 int64
	_ = v10
	var v13 int32
	_ = v13
	v10 = Fn14381(m, l0, int32(_a_F_timestamptz_bin_0), int32(_a_F_timestamptz_bin_1), int32(_a_F_timestamptz_bin_2), int32(_a_F_timestamptz_bin_3), int32(_a_F_timestamptz_bin_4), int32(_a_F_timestamptz_bin_5), int32(_a_F_timestamptz_bin_6), int32(_a_F_timestamptz_bin_7))
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		return v10
	}
}
func F_timestamptz_cmp_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
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
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_cmp_date[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_cmp_date[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_date2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
				} else {
					if v9 == int64(9223372036854775807) {
						v32 = int32(-1)
					} else {
						v32 = int32(1)
					}
					v41 = v32
				}
			} else {
				if v9 == int64(-9223372036854775807-1) {
					v37 = int32(1)
				} else {
					v37 = int32(-1)
				}
				v41 = v37
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_s(int32(0) - v41)
	}
}
func F_timestamptz_cmp_timestamp(m *base.Module, l0 int32) int64 {
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
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_cmp_timestamp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_cmp_timestamp[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v41 = base.B2i32(v17 < v9) - base.B2i32(v9 < v17)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v41 = base.B2i32(v17 < v9) - base.B2i32(v9 < v17)
				} else {
					if v9 == int64(9223372036854775807) {
						v32 = int32(1)
					} else {
						v32 = int32(-1)
					}
					v41 = v32
				}
			} else {
				if v9 == int64(-9223372036854775807-1) {
					v37 = int32(-1)
				} else {
					v37 = int32(1)
				}
				v41 = v37
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_s(v41)
	}
}
func F_timestamptz_eq_timestamp(m *base.Module, l0 int32) int64 {
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
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_eq_timestamp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_eq_timestamp[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u((v21 ^ int32(-1) | base.B2i32(base.Ui64(v17+int64(9223372036854775807)) < base.Ui64(int64(-2)))) & base.B2i32(v17 == v9))
	}
}
func F_timestamptz_ge_date(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
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
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v41 int32
	_ = v41
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_ge_date[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_ge_date[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_date2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v41 = base.B2i32(v9 < v17) - base.B2i32(v17 < v9)
				} else {
					if v9 == int64(9223372036854775807) {
						v32 = int32(-1)
					} else {
						v32 = int32(1)
					}
					v41 = v32
				}
			} else {
				if v9 == int64(-9223372036854775807-1) {
					v37 = int32(1)
				} else {
					v37 = int32(-1)
				}
				v41 = v37
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(base.B2i32(v41 <= int32(0)))
	}
}
func F_timestamptz_in(m *base.Module, l0 int32) int64 {
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
	var v56 int32
	_ = v56
	var v59 int64
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v101 int32
	_ = v101
	var v106 int32
	_ = v106
	var v112 int64
	_ = v112
	var v121 int64
	_ = v121
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v127 int64
	_ = v127
	var v128 int64
	_ = v128
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v135 int64
	_ = v135
	var v142 int64
	_ = v142
	var v153 int64
	_ = v153
	var v154 int64
	_ = v154
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v170 int64
	_ = v170
	var v173 int64
	_ = v173
	var v177 int32
	_ = v177
	var v182 int64
	_ = v182
	var v191 int64
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v198 int32
	_ = v198
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v214 int32
	_ = v214
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v226 int64
	_ = v226
	var v227 int32
	_ = v227
	var v228 int64
	_ = v228
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int64
	_ = v238
	var v241 int64
	_ = v241
	v11 = m.G0
	v13 = v11 - int32(512)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
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
			return int64(0)
		} else {
			if v41 == int32(0) {
				v56 = *(*int32)(unsafe.Add(mBase, uint32(v13)+448))
				switch v56 - int32(2) {
				case 0:
					v59 = int64(*(*int32)(unsafe.Add(mBase, uint32(v13)+500)))
					v60 = *(*int32)(unsafe.Add(mBase, uint32(v13)+476))
					if v60 <= int32(-4713) {
						if v60 != int32(-4713) {
							v191 = int64(0)
							v192 = F_errsave_start(m, v15)
							mBase = m.M
							v193 = m.ExcPending
							if v193 != 0 {
								return int64(0)
							} else {
								if v192 == int32(0) {
									v241 = v191
									m.G0 = v13 + int32(512)
									return v241
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v198 = m.ExcPending
									if v198 != 0 {
										return int64(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
										F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return int64(0)
										} else {
											F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
												return int64(0)
											} else {
												v241 = v191
												m.G0 = v13 + int32(512)
												return v241
											}
										}
									}
								}
							}
						} else {
							v65 = *(*int32)(unsafe.Add(mBase, uint32(v13)+472))
							if int32(10) < v65 {
								v76 = v65
								v78 = v13 + int32(32)
								v79 = *(*int32)(unsafe.Add(mBase, uint32(v13)+468))
								v84 = base.B2i32(int32(2) < v76)
								if int32(2) < v76 {
									v85 = int32(_a_F_timestamptz_in_3)
								} else {
									v85 = int32(_a_F_timestamptz_in_4)
								}
								v86 = v85 + v60
								v91 = base.I32_div_s(v86, int32(4))
								v94 = base.I32_div_s(v86, int32(-100))
								v97 = base.I32_div_s(v86, int32(400))
								if int32(2) < v76 {
									v101 = int32(1)
								} else {
									v101 = int32(13)
								}
								v106 = base.I32_div_s((v101+v76)*int32(_a_F_timestamptz_in_5), int32(256))
								v112 = base.I64_extend_i32_s(v79 + v86*int32(365) + v91 + v94 + v97 + v106 - int32(_a_F_timestamptz_in_6) - int32(_a_F_timestamptz_in_7))
								v121 = int64(32)
								v122 = int64(20)
								v124 = int64(base.Ui64(v112) >> (uint(v121) % 64))
								v127 = int64(4294967295)
								v128 = int64(500654080)
								v130 = v112 & v127
								v131 = v128 * v130
								v135 = int64(base.Ui64(v131)>>(uint(v121)%64)) + v128*v124
								v142 = v130*v122 + v135&v127
								*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v112*int64(0) + v112>>(uint(int64(63))%64)*int64(86400000000) + v122*v124 + int64(base.Ui64(v135)>>(uint(v121)%64)) + int64(base.Ui64(v142)>>(uint(v121)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v78))) = v131&v127 | v142<<(uint(v121)%64)
								v153 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
								v154 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
								if v153 != v154>>(uint(int64(63))%64) {
									v191 = int64(0)
									v192 = F_errsave_start(m, v15)
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int64(0)
									} else {
										if v192 == int32(0) {
											v241 = v191
											m.G0 = v13 + int32(512)
											return v241
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v198 = m.ExcPending
											if v198 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
												F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return int64(0)
													} else {
														v241 = v191
														m.G0 = v13 + int32(512)
														return v241
													}
												}
											}
										}
									}
								} else {
									v158 = *(*int32)(unsafe.Add(mBase, uint32(v13)+456))
									v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+460))
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+464))
									v161 = int32(60)
									v170 = base.I64_extend_i32_s(v158+(v159+v160*v161)*v161)*int64(1000000) + v59
									v173 = v154 + v170
									if base.B2i32(v170 < int64(0))^base.B2i32(v173 < v154) != 0 {
										v191 = int64(0)
										v192 = F_errsave_start(m, v15)
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int64(0)
										} else {
											if v192 == int32(0) {
												v241 = v191
												m.G0 = v13 + int32(512)
												return v241
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v198 = m.ExcPending
												if v198 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
													F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int64(0)
														} else {
															v241 = v191
															m.G0 = v13 + int32(512)
															return v241
														}
													}
												}
											}
										}
									} else {
										v177 = *(*int32)(unsafe.Add(mBase, uint32(v13)+452))
										v182 = base.I64_extend_i32_s(int32(0)-v177)*int64(-1000000) + v173
										*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v182
										if base.Ui64(v182+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
											v236 = F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int64(0)
											} else {
												v238 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
												v241 = v238
												m.G0 = v13 + int32(512)
												return v241
											}
										} else {
											v191 = int64(0)
											v192 = F_errsave_start(m, v15)
											mBase = m.M
											v193 = m.ExcPending
											if v193 != 0 {
												return int64(0)
											} else {
												if v192 == int32(0) {
													v241 = v191
													m.G0 = v13 + int32(512)
													return v241
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v198 = m.ExcPending
													if v198 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
														F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int64(0)
															} else {
																v241 = v191
																m.G0 = v13 + int32(512)
																return v241
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v191 = int64(0)
								v192 = F_errsave_start(m, v15)
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return int64(0)
								} else {
									if v192 == int32(0) {
										v241 = v191
										m.G0 = v13 + int32(512)
										return v241
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
											F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return int64(0)
												} else {
													v241 = v191
													m.G0 = v13 + int32(512)
													return v241
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v60 <= int32(_a_F_timestamptz_in_8) {
							v70 = *(*int32)(unsafe.Add(mBase, uint32(v13)+472))
							v76 = v70
							v78 = v13 + int32(32)
							v79 = *(*int32)(unsafe.Add(mBase, uint32(v13)+468))
							v84 = base.B2i32(int32(2) < v76)
							if int32(2) < v76 {
								v85 = int32(_a_F_timestamptz_in_3)
							} else {
								v85 = int32(_a_F_timestamptz_in_4)
							}
							v86 = v85 + v60
							v91 = base.I32_div_s(v86, int32(4))
							v94 = base.I32_div_s(v86, int32(-100))
							v97 = base.I32_div_s(v86, int32(400))
							if int32(2) < v76 {
								v101 = int32(1)
							} else {
								v101 = int32(13)
							}
							v106 = base.I32_div_s((v101+v76)*int32(_a_F_timestamptz_in_5), int32(256))
							v112 = base.I64_extend_i32_s(v79 + v86*int32(365) + v91 + v94 + v97 + v106 - int32(_a_F_timestamptz_in_6) - int32(_a_F_timestamptz_in_7))
							v121 = int64(32)
							v122 = int64(20)
							v124 = int64(base.Ui64(v112) >> (uint(v121) % 64))
							v127 = int64(4294967295)
							v128 = int64(500654080)
							v130 = v112 & v127
							v131 = v128 * v130
							v135 = int64(base.Ui64(v131)>>(uint(v121)%64)) + v128*v124
							v142 = v130*v122 + v135&v127
							*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v112*int64(0) + v112>>(uint(int64(63))%64)*int64(86400000000) + v122*v124 + int64(base.Ui64(v135)>>(uint(v121)%64)) + int64(base.Ui64(v142)>>(uint(v121)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v78))) = v131&v127 | v142<<(uint(v121)%64)
							v153 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
							v154 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
							if v153 != v154>>(uint(int64(63))%64) {
								v191 = int64(0)
								v192 = F_errsave_start(m, v15)
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return int64(0)
								} else {
									if v192 == int32(0) {
										v241 = v191
										m.G0 = v13 + int32(512)
										return v241
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
											F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return int64(0)
												} else {
													v241 = v191
													m.G0 = v13 + int32(512)
													return v241
												}
											}
										}
									}
								}
							} else {
								v158 = *(*int32)(unsafe.Add(mBase, uint32(v13)+456))
								v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+460))
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+464))
								v161 = int32(60)
								v170 = base.I64_extend_i32_s(v158+(v159+v160*v161)*v161)*int64(1000000) + v59
								v173 = v154 + v170
								if base.B2i32(v170 < int64(0))^base.B2i32(v173 < v154) != 0 {
									v191 = int64(0)
									v192 = F_errsave_start(m, v15)
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int64(0)
									} else {
										if v192 == int32(0) {
											v241 = v191
											m.G0 = v13 + int32(512)
											return v241
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v198 = m.ExcPending
											if v198 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
												F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return int64(0)
													} else {
														v241 = v191
														m.G0 = v13 + int32(512)
														return v241
													}
												}
											}
										}
									}
								} else {
									v177 = *(*int32)(unsafe.Add(mBase, uint32(v13)+452))
									v182 = base.I64_extend_i32_s(int32(0)-v177)*int64(-1000000) + v173
									*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v182
									if base.Ui64(v182+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
										v236 = F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
										mBase = m.M
										v237 = m.ExcPending
										if v237 != 0 {
											return int64(0)
										} else {
											v238 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
											v241 = v238
											m.G0 = v13 + int32(512)
											return v241
										}
									} else {
										v191 = int64(0)
										v192 = F_errsave_start(m, v15)
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int64(0)
										} else {
											if v192 == int32(0) {
												v241 = v191
												m.G0 = v13 + int32(512)
												return v241
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v198 = m.ExcPending
												if v198 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
													F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int64(0)
														} else {
															v241 = v191
															m.G0 = v13 + int32(512)
															return v241
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							if v60 != int32(_a_F_timestamptz_in_9) {
								v191 = int64(0)
								v192 = F_errsave_start(m, v15)
								mBase = m.M
								v193 = m.ExcPending
								if v193 != 0 {
									return int64(0)
								} else {
									if v192 == int32(0) {
										v241 = v191
										m.G0 = v13 + int32(512)
										return v241
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v198 = m.ExcPending
										if v198 != 0 {
											return int64(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
											F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
											mBase = m.M
											v204 = m.ExcPending
											if v204 != 0 {
												return int64(0)
											} else {
												F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
												mBase = m.M
												v209 = m.ExcPending
												if v209 != 0 {
													return int64(0)
												} else {
													v241 = v191
													m.G0 = v13 + int32(512)
													return v241
												}
											}
										}
									}
								}
							} else {
								v73 = *(*int32)(unsafe.Add(mBase, uint32(v13)+472))
								if int32(5) < v73 {
									v191 = int64(0)
									v192 = F_errsave_start(m, v15)
									mBase = m.M
									v193 = m.ExcPending
									if v193 != 0 {
										return int64(0)
									} else {
										if v192 == int32(0) {
											v241 = v191
											m.G0 = v13 + int32(512)
											return v241
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v198 = m.ExcPending
											if v198 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
												F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
												mBase = m.M
												v204 = m.ExcPending
												if v204 != 0 {
													return int64(0)
												} else {
													F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
													mBase = m.M
													v209 = m.ExcPending
													if v209 != 0 {
														return int64(0)
													} else {
														v241 = v191
														m.G0 = v13 + int32(512)
														return v241
													}
												}
											}
										}
									}
								} else {
									v76 = v73
									v78 = v13 + int32(32)
									v79 = *(*int32)(unsafe.Add(mBase, uint32(v13)+468))
									v84 = base.B2i32(int32(2) < v76)
									if int32(2) < v76 {
										v85 = int32(_a_F_timestamptz_in_3)
									} else {
										v85 = int32(_a_F_timestamptz_in_4)
									}
									v86 = v85 + v60
									v91 = base.I32_div_s(v86, int32(4))
									v94 = base.I32_div_s(v86, int32(-100))
									v97 = base.I32_div_s(v86, int32(400))
									if int32(2) < v76 {
										v101 = int32(1)
									} else {
										v101 = int32(13)
									}
									v106 = base.I32_div_s((v101+v76)*int32(_a_F_timestamptz_in_5), int32(256))
									v112 = base.I64_extend_i32_s(v79 + v86*int32(365) + v91 + v94 + v97 + v106 - int32(_a_F_timestamptz_in_6) - int32(_a_F_timestamptz_in_7))
									v121 = int64(32)
									v122 = int64(20)
									v124 = int64(base.Ui64(v112) >> (uint(v121) % 64))
									v127 = int64(4294967295)
									v128 = int64(500654080)
									v130 = v112 & v127
									v131 = v128 * v130
									v135 = int64(base.Ui64(v131)>>(uint(v121)%64)) + v128*v124
									v142 = v130*v122 + v135&v127
									*(*int64)(unsafe.Add(mBase, uint32(v78)+8)) = v112*int64(0) + v112>>(uint(int64(63))%64)*int64(86400000000) + v122*v124 + int64(base.Ui64(v135)>>(uint(v121)%64)) + int64(base.Ui64(v142)>>(uint(v121)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v78))) = v131&v127 | v142<<(uint(v121)%64)
									v153 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
									v154 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
									if v153 != v154>>(uint(int64(63))%64) {
										v191 = int64(0)
										v192 = F_errsave_start(m, v15)
										mBase = m.M
										v193 = m.ExcPending
										if v193 != 0 {
											return int64(0)
										} else {
											if v192 == int32(0) {
												v241 = v191
												m.G0 = v13 + int32(512)
												return v241
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v198 = m.ExcPending
												if v198 != 0 {
													return int64(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
													F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
													mBase = m.M
													v204 = m.ExcPending
													if v204 != 0 {
														return int64(0)
													} else {
														F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
														mBase = m.M
														v209 = m.ExcPending
														if v209 != 0 {
															return int64(0)
														} else {
															v241 = v191
															m.G0 = v13 + int32(512)
															return v241
														}
													}
												}
											}
										}
									} else {
										v158 = *(*int32)(unsafe.Add(mBase, uint32(v13)+456))
										v159 = *(*int32)(unsafe.Add(mBase, uint32(v13)+460))
										v160 = *(*int32)(unsafe.Add(mBase, uint32(v13)+464))
										v161 = int32(60)
										v170 = base.I64_extend_i32_s(v158+(v159+v160*v161)*v161)*int64(1000000) + v59
										v173 = v154 + v170
										if base.B2i32(v170 < int64(0))^base.B2i32(v173 < v154) != 0 {
											v191 = int64(0)
											v192 = F_errsave_start(m, v15)
											mBase = m.M
											v193 = m.ExcPending
											if v193 != 0 {
												return int64(0)
											} else {
												if v192 == int32(0) {
													v241 = v191
													m.G0 = v13 + int32(512)
													return v241
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v198 = m.ExcPending
													if v198 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
														F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
														mBase = m.M
														v204 = m.ExcPending
														if v204 != 0 {
															return int64(0)
														} else {
															F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
															mBase = m.M
															v209 = m.ExcPending
															if v209 != 0 {
																return int64(0)
															} else {
																v241 = v191
																m.G0 = v13 + int32(512)
																return v241
															}
														}
													}
												}
											}
										} else {
											v177 = *(*int32)(unsafe.Add(mBase, uint32(v13)+452))
											v182 = base.I64_extend_i32_s(int32(0)-v177)*int64(-1000000) + v173
											*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v182
											if base.Ui64(v182+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												v236 = F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return int64(0)
												} else {
													v238 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
													v241 = v238
													m.G0 = v13 + int32(512)
													return v241
												}
											} else {
												v191 = int64(0)
												v192 = F_errsave_start(m, v15)
												mBase = m.M
												v193 = m.ExcPending
												if v193 != 0 {
													return int64(0)
												} else {
													if v192 == int32(0) {
														v241 = v191
														m.G0 = v13 + int32(512)
														return v241
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v198 = m.ExcPending
														if v198 != 0 {
															return int64(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = v17
															F_errmsg(m, int32(_a_F_timestamptz_in_0), v13+int32(16))
															mBase = m.M
															v204 = m.ExcPending
															if v204 != 0 {
																return int64(0)
															} else {
																F_errsave_finish(m, v15, int32(_a_F_timestamptz_in_1), int32(453), int32(_a_F_timestamptz_in_2))
																mBase = m.M
																v209 = m.ExcPending
																if v209 != 0 {
																	return int64(0)
																} else {
																	v241 = v191
																	m.G0 = v13 + int32(512)
																	return v241
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
					v214 = m.ExcPending
					if v214 != 0 {
						return int64(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v13)+4)) = v17
						v216 = *(*int32)(unsafe.Add(mBase, uint32(v13)+448))
						*(*int32)(unsafe.Add(mBase, uint32(v13))) = v216
						F_errmsg_internal(m, int32(_a_F_timestamptz_in_10), v13)
						mBase = m.M
						v220 = m.ExcPending
						if v220 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamptz_in_1), int32(470), int32(_a_F_timestamptz_in_2))
							mBase = m.M
							v225 = m.ExcPending
							if v225 != 0 {
								return int64(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 7:
					v228 = int64(-9223372036854775807 - 1)
					*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v228
					v236 = F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
					mBase = m.M
					v237 = m.ExcPending
					if v237 != 0 {
						return int64(0)
					} else {
						v238 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
						v241 = v238
						m.G0 = v13 + int32(512)
						return v241
					}
				case 8:
					v228 = int64(9223372036854775807)
					*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v228
					v236 = F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
					mBase = m.M
					v237 = m.ExcPending
					if v237 != 0 {
						return int64(0)
					} else {
						v238 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
						v241 = v238
						m.G0 = v13 + int32(512)
						return v241
					}
				case 9:
					v226 = F_SetEpochTimestamp(m)
					mBase = m.M
					v227 = m.ExcPending
					if v227 != 0 {
						return int64(0)
					} else {
						v228 = v226
						*(*int64)(unsafe.Add(mBase, uint32(v13)+504)) = v228
						v236 = F_AdjustTimestampForTypmod(m, v13+int32(504), v16, v15)
						mBase = m.M
						v237 = m.ExcPending
						if v237 != 0 {
							return int64(0)
						} else {
							v238 = *(*int64)(unsafe.Add(mBase, uint32(v13)+504))
							v241 = v238
							m.G0 = v13 + int32(512)
							return v241
						}
					}
				}
			} else {
				v47 = v41
				F_DateTimeParseError(m, v47, v13+int32(56), v17, int32(_a_F_timestamptz_in_11), v15)
				mBase = m.M
				v52 = m.ExcPending
				if v52 != 0 {
					return int64(0)
				} else {
					v53 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v53)
					v241 = int64(0)
					m.G0 = v13 + int32(512)
					return v241
				}
			}
		}
	} else {
		v47 = v27
		F_DateTimeParseError(m, v47, v13+int32(56), v17, int32(_a_F_timestamptz_in_11), v15)
		mBase = m.M
		v52 = m.ExcPending
		if v52 != 0 {
			return int64(0)
		} else {
			v53 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v53)
			v241 = int64(0)
			m.G0 = v13 + int32(512)
			return v241
		}
	}
}
func F_timestamptz_izone(m *base.Module, l0 int32) int64 {
	var v7 int64
	_ = v7
	var v10 int32
	_ = v10
	v7 = Fn14382(m, l0, int32(_a_F_timestamptz_izone_0), int32(_a_F_timestamptz_izone_1), int32(_a_F_timestamptz_izone_2), int64(-1000000), int32(_a_F_timestamptz_izone_3))
	v10 = m.ExcPending
	if v10 != 0 {
		return int64(0)
	} else {
		return v7
	}
}
func F_timestamptz_lt_timestamp(m *base.Module, l0 int32) int64 {
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
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
	v12 = *(*int32)(unsafe.Add(mBase, _c_F_timestamptz_lt_timestamp[0]))
	*(*int32)(unsafe.Add(mBase, uint32(v7)+8)) = v12
	v15 = *(*int64)(unsafe.Add(mBase, _c_F_timestamptz_lt_timestamp[1]))
	*(*int64)(unsafe.Add(mBase, uint32(v7))) = v15
	v17 = F_timestamp2timestamptz_safe(m, v10, v7)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v7)+4)))
		if v21 != int32(1) {
			v33 = base.B2i32(v9 < v17)
		} else {
			if v17 != int64(-9223372036854775807-1) {
				if v17 != int64(9223372036854775807) {
					v33 = base.B2i32(v9 < v17)
				} else {
					v33 = base.B2i32(v9 != int64(9223372036854775807))
				}
			} else {
				v33 = base.B2i32(v9 == int64(-9223372036854775807-1))
			}
		}
		m.G0 = v7 + int32(16)
		return base.I64_extend_i32_u(v33)
	}
}
func F_timestamptz_mi_interval_at_zone(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(256)
	m.G0 = v7
	v9 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
	v12 = F_pg_detoast_datum_packed(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		F_text_to_cstring_buffer(m, v12, v7, int32(256))
		mBase = m.M
		v18 = m.ExcPending
		if v18 != 0 {
			return int64(0)
		} else {
			v19 = F_DecodeTimezoneNameToTz(m, v7)
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				F_interval_um_internal(m, v10, v7)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					v23 = F_timestamptz_pl_interval_internal(m, v9, v7, v19)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						m.G0 = v7 + int32(256)
						return v23
					}
				}
			}
		}
	}
}
func F_timestamptz_part_common(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v23 int32
	_ = v23
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v28 int64
	_ = v28
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v55 int32
	_ = v55
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v64 int32
	_ = v64
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v76 int32
	_ = v76
	var v80 float64
	_ = v80
	var v81 int32
	_ = v81
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v102 int64
	_ = v102
	var v103 int32
	_ = v103
	var v105 int32
	_ = v105
	var v117 int32
	_ = v117
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v127 int32
	_ = v127
	var v128 int32
	_ = v128
	var v130 int32
	_ = v130
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v143 int32
	_ = v143
	var v144 int64
	_ = v144
	var v150 int32
	_ = v150
	var v151 int32
	_ = v151
	var v154 float64
	_ = v154
	var v156 int32
	_ = v156
	var v162 int32
	_ = v162
	var v163 int64
	_ = v163
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v172 int32
	_ = v172
	var v179 int64
	_ = v179
	var v180 int64
	_ = v180
	var v181 int64
	_ = v181
	var v182 int64
	_ = v182
	var v183 int32
	_ = v183
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v198 int32
	_ = v198
	var v201 int32
	_ = v201
	var v204 int32
	_ = v204
	var v208 int32
	_ = v208
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v218 int32
	_ = v218
	var v222 int32
	_ = v222
	var v225 int32
	_ = v225
	var v226 int32
	_ = v226
	var v229 int32
	_ = v229
	var v231 int32
	_ = v231
	var v233 int32
	_ = v233
	var v237 int32
	_ = v237
	var v244 int32
	_ = v244
	var v248 int32
	_ = v248
	var v253 int32
	_ = v253
	var v255 int32
	_ = v255
	var v261 int32
	_ = v261
	var v266 int32
	_ = v266
	var v268 int32
	_ = v268
	var v274 int32
	_ = v274
	var v279 int32
	_ = v279
	var v281 int32
	_ = v281
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v289 int32
	_ = v289
	var v290 int32
	_ = v290
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v301 int32
	_ = v301
	var v305 int32
	_ = v305
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v317 int64
	_ = v317
	var v318 int32
	_ = v318
	var v319 int32
	_ = v319
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v331 int32
	_ = v331
	var v332 int32
	_ = v332
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
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v349 int32
	_ = v349
	var v350 int32
	_ = v350
	var v363 int32
	_ = v363
	var v364 int32
	_ = v364
	var v365 int32
	_ = v365
	var v367 int32
	_ = v367
	var v368 int32
	_ = v368
	var v370 int32
	_ = v370
	var v373 int32
	_ = v373
	var v376 int32
	_ = v376
	var v377 int32
	_ = v377
	var v380 int32
	_ = v380
	var v383 int32
	_ = v383
	var v384 int32
	_ = v384
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v395 int32
	_ = v395
	var v398 int32
	_ = v398
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v419 int32
	_ = v419
	var v424 int32
	_ = v424
	var v427 int32
	_ = v427
	var v430 int32
	_ = v430
	var v434 int32
	_ = v434
	var v439 int32
	_ = v439
	var v445 int32
	_ = v445
	var v446 int32
	_ = v446
	var v451 int32
	_ = v451
	var v452 int64
	_ = v452
	var v454 int64
	_ = v454
	var v455 int32
	_ = v455
	var v458 int64
	_ = v458
	var v459 int32
	_ = v459
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v473 int32
	_ = v473
	var v476 int32
	_ = v476
	var v479 int32
	_ = v479
	var v483 int32
	_ = v483
	var v488 int32
	_ = v488
	var v492 int32
	_ = v492
	var v501 int32
	_ = v501
	var v506 int32
	_ = v506
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v521 int32
	_ = v521
	var v532 int32
	_ = v532
	var v535 int32
	_ = v535
	var v537 int32
	_ = v537
	var v538 int32
	_ = v538
	var v545 int32
	_ = v545
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v554 int64
	_ = v554
	var v555 int32
	_ = v555
	var v557 int64
	_ = v557
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
	var v571 int32
	_ = v571
	var v572 int32
	_ = v572
	var v574 int32
	_ = v574
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v581 int64
	_ = v581
	var v582 int32
	_ = v582
	var v584 int32
	_ = v584
	var v585 int32
	_ = v585
	var v593 float64
	_ = v593
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
	var v656 int64
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
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		return int64(0)
	} else {
		v21 = int32(1)
		v23 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
		v25 = v23 & v21
		if v25 != 0 {
			v26 = v21
		} else {
			v26 = int32(4)
		}
		v28 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		if v23 == int32(1) {
			v34 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17)+1)))
			if v34 == int32(18) {
				v37 = int32(16)
			} else {
				v37 = int32(0)
			}
			if base.Ui32((v34-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v44 = int32(4)
			} else {
				v44 = v37
			}
			v55 = v44
		} else {
			v45 = int32(1)
			if v25 != 0 {
				v55 = int32(base.Ui32(v23)>>(uint(v45)%32)) - v45
			} else {
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
				v55 = int32(base.Ui32(v49)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v57 = F_downcase_truncate_identifier(m, v17+v26, v55, int32(0))
		mBase = m.M
		v58 = m.ExcPending
		if v58 != 0 {
			return int64(0)
		} else {
			v60 = v14 + int32(88)
			v64 = Fn14210(m, v57, v60, int32(_a_F_timestamptz_part_common_0), int32(_a_F_timestamptz_part_common_1), int32(_a_F_timestamptz_part_common_2))
			mBase = m.M
			if v64 == int32(31) {
				v70 = Fn14210(m, v57, v60, int32(_a_F_timestamptz_part_common_3), int32(_a_F_timestamptz_part_common_4), int32(_a_F_timestamptz_part_common_5))
				mBase = m.M
				v71 = v70
			} else {
				v71 = v64
			}
			if base.Ui64(int64(1)) < base.Ui64(v28-int64(9223372036854775807)) {
				if v71 != 0 {
					if v71 != int32(17) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v622 = m.ExcPending
						if v622 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v625 = m.ExcPending
							if v625 != 0 {
								return int64(0)
							} else {
								v627 = F_format_type_be(m, int32(1184))
								mBase = m.M
								v628 = m.ExcPending
								if v628 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v627
									*(*int32)(unsafe.Add(mBase, uint32(v14))) = v57
									F_errmsg(m, int32(_a_F_timestamptz_part_common_6), v14)
									mBase = m.M
									v633 = m.ExcPending
									if v633 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_8), int32(_a_F_timestamptz_part_common_9))
										mBase = m.M
										v638 = m.ExcPending
										if v638 != 0 {
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
						v117 = int32(0)
						v119 = F_timestamp2tm(m, v28, v14+int32(92), v14+int32(40), v14+int32(84), v117, v117)
						mBase = m.M
						v120 = m.ExcPending
						if v120 != 0 {
							return int64(0)
						} else {
							if v119 != 0 {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v664 = m.ExcPending
								if v664 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v667 = m.ExcPending
									if v667 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_timestamptz_part_common_10), int32(0))
										mBase = m.M
										v671 = m.ExcPending
										if v671 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_11), int32(_a_F_timestamptz_part_common_9))
											mBase = m.M
											v676 = m.ExcPending
											if v676 != 0 {
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
								v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
								switch v121 - int32(4) {
								case 0:
									v640 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
									v645 = base.I64_extend_i32_s(int32(0) - v640)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								default:
									F_errstart_cold(m, int32(21), int32(0))
									mBase = m.M
									v532 = m.ExcPending
									if v532 != 0 {
										return int64(0)
									} else {
										F_errcode(m, int32(1088))
										mBase = m.M
										v535 = m.ExcPending
										if v535 != 0 {
											return int64(0)
										} else {
											v537 = F_format_type_be(m, int32(1184))
											mBase = m.M
											v538 = m.ExcPending
											if v538 != 0 {
												return int64(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v537
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v57
												F_errmsg(m, int32(_a_F_timestamptz_part_common_12), v14+int32(16))
												mBase = m.M
												v545 = m.ExcPending
												if v545 != 0 {
													return int64(0)
												} else {
													F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_13), int32(_a_F_timestamptz_part_common_9))
													mBase = m.M
													v550 = m.ExcPending
													if v550 != 0 {
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
								case 14:
									v162 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
									if l1 != 0 {
										v163 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
										v169 = F_int64_div_fast_to_numeric(m, v163+base.I64_extend_i32_s(v162)*int64(1000000), int32(6))
										mBase = m.M
										v170 = m.ExcPending
										if v170 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v169)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v172 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v656 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_convert_i32_s(v172), float64(1e+06)), base.F64_convert_i32_s(v162)))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 15:
									v179 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+44)))
									v645 = v179
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 16:
									v180 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+48)))
									v645 = v180
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 17:
									v181 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+52)))
									v645 = v181
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 18:
									v191 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v192 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v193 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v195 = F_date2j(m, v191, v192, v193)
									mBase = m.M
									v196 = int32(1)
									v198 = F_date2j(m, v191, v196, int32(4))
									mBase = m.M
									v201 = F_j2day(m, v198-v196)
									mBase = m.M
									if v195 < v198-v201 {
										v204 = int32(1)
										v208 = F_date2j(m, v191-v204, v204, int32(4))
										mBase = m.M
										v211 = F_j2day(m, v208-v204)
										mBase = m.M
										v212 = v208
										v213 = v211
									} else {
										v212 = v198
										v213 = v201
									}
									v215 = v213 - v212 + v195
									if int32(357) <= v215 {
										v218 = int32(1)
										v222 = F_date2j(m, v191+v218, v218, int32(4))
										mBase = m.M
										v225 = F_j2day(m, v222-v218)
										mBase = m.M
										v226 = v222 - v225
										if v195 < v226 {
											v229 = v215
										} else {
											v229 = v195 - v226
										}
										v231 = v229
									} else {
										v231 = v215
									}
									v233 = base.I32_div_s(v231, int32(7))
									v645 = base.I64_extend_i32_s(v233 + int32(1))
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 19:
									v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+56)))
									v645 = v182
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 20:
									v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v184 = int32(1)
									v187 = base.I32_div_s(v183-v184, int32(3))
									v645 = base.I64_extend_i32_s(v187 + v184)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 21:
									v237 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if int32(0) < v237 {
										v645 = base.I64_extend_i32_u(v237)
									} else {
										v645 = base.I64_extend_i32_s(v237 - int32(1))
									}
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 22:
									v244 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if int32(0) < v244 {
										v248 = base.I32_div_u_s(v244, int32(10))
										v645 = base.I64_extend_i32_u(v248)
									} else {
										v253 = base.I32_div_s(int32(9)-v244, int32(-10))
										v645 = base.I64_extend_i32_s(v253)
									}
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 23:
									v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if int32(0) < v255 {
										v261 = base.I32_div_s(v255+int32(99), int32(100))
										v645 = base.I64_extend_i32_s(v261)
									} else {
										v266 = base.I32_div_s(int32(100)-v255, int32(-100))
										v645 = base.I64_extend_i32_s(v266)
									}
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 24:
									v268 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									if int32(0) < v268 {
										v274 = base.I32_div_s(v268+int32(999), int32(1000))
										v645 = base.I64_extend_i32_s(v274)
									} else {
										v279 = base.I32_div_s(int32(1000)-v268, int32(-1000))
										v645 = base.I64_extend_i32_s(v279)
									}
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 25:
									v143 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
									if l1 != 0 {
										v144 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
										v150 = F_int64_div_fast_to_numeric(m, v144+base.I64_extend_i32_s(v143)*int64(1000000), int32(3))
										mBase = m.M
										v151 = m.ExcPending
										if v151 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v150)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v154 = float64(1000)
										v156 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v656 = base.I64_reinterpret_f64(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v143), v154), base.F64_div(base.F64_convert_i32_s(v156), v154)))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 26:
									v138 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
									v139 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+40)))
									v645 = v138 + v139*int64(1000000)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 27:
									v281 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v282 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v283 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v288 = base.B2i32(int32(2) < v282)
									if int32(2) < v282 {
										v289 = int32(_a_F_timestamptz_part_common_14)
									} else {
										v289 = int32(_a_F_timestamptz_part_common_15)
									}
									v290 = v289 + v281
									v295 = base.I32_div_s(v290, int32(4))
									v298 = base.I32_div_s(v290, int32(-100))
									v301 = base.I32_div_s(v290, int32(400))
									if int32(2) < v282 {
										v305 = int32(1)
									} else {
										v305 = int32(13)
									}
									v310 = base.I32_div_s((v305+v282)*int32(_a_F_timestamptz_part_common_16), int32(256))
									v313 = v283 + v290*int32(365) + v295 + v298 + v301 + v310 - int32(_a_F_timestamptz_part_common_17)
									if l1 != 0 {
										v315 = F_int64_to_numeric(m, base.I64_extend_i32_s(v313))
										mBase = m.M
										v316 = m.ExcPending
										if v316 != 0 {
											return int64(0)
										} else {
											v317 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
											v318 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
											v319 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
											v320 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
											v321 = int32(60)
											v331 = F_int64_to_numeric(m, v317+base.I64_extend_i32_s(v318+(v319+v320*v321)*v321)*int64(1000000))
											mBase = m.M
											v332 = m.ExcPending
											if v332 != 0 {
												return int64(0)
											} else {
												v334 = F_int64_to_numeric(m, int64(86400000000))
												mBase = m.M
												v335 = m.ExcPending
												if v335 != 0 {
													return int64(0)
												} else {
													v337 = F_numeric_div_safe(m, v331, v334, int32(0))
													mBase = m.M
													v338 = m.ExcPending
													if v338 != 0 {
														return int64(0)
													} else {
														v340 = F_numeric_add_safe(m, v315, v337, int32(0))
														mBase = m.M
														v341 = m.ExcPending
														if v341 != 0 {
															return int64(0)
														} else {
															v656 = base.I64_extend_i32_u(v340)
															m.G0 = v14 + int32(96)
															return v656
														}
													}
												}
											}
										}
									} else {
										v343 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
										v347 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
										v348 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
										v349 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
										v350 = int32(60)
										v656 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v343), float64(1e+06)), base.F64_convert_i32_s(v347+(v348+v349*v350)*v350)), float64(86400)), base.F64_convert_i32_s(v313)))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 28, 33:
									v410 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v411 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v412 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v417 = base.B2i32(int32(2) < v411)
									if int32(2) < v411 {
										v418 = int32(_a_F_timestamptz_part_common_14)
									} else {
										v418 = int32(_a_F_timestamptz_part_common_15)
									}
									v419 = v418 + v410
									v424 = base.I32_div_s(v419, int32(4))
									v427 = base.I32_div_s(v419, int32(-100))
									v430 = base.I32_div_s(v419, int32(400))
									if int32(2) < v411 {
										v434 = int32(1)
									} else {
										v434 = int32(13)
									}
									v439 = base.I32_div_s((v434+v411)*int32(_a_F_timestamptz_part_common_16), int32(256))
									v445 = int32(7)
									v446 = base.I32_rem_s(v412+v419*int32(365)+v424+v427+v430+v439-int32(_a_F_timestamptz_part_common_17)+int32(1), v445)
									if v446 < int32(0) {
										v451 = v446 + v445
									} else {
										v451 = v446
									}
									v452 = base.I64_extend_i32_s(v451)
									if v451 != 0 {
										v454 = v452
									} else {
										v454 = int64(7)
									}
									v455 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
									if v455 == int32(37) {
										v458 = v454
									} else {
										v458 = v452
									}
									v645 = v458
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 29:
									v459 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v460 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v461 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v466 = base.B2i32(int32(2) < v460)
									if int32(2) < v460 {
										v467 = int32(_a_F_timestamptz_part_common_14)
									} else {
										v467 = int32(_a_F_timestamptz_part_common_15)
									}
									v468 = v467 + v459
									v473 = base.I32_div_s(v468, int32(4))
									v476 = base.I32_div_s(v468, int32(-100))
									v479 = base.I32_div_s(v468, int32(400))
									if int32(2) < v460 {
										v483 = int32(1)
									} else {
										v483 = int32(13)
									}
									v488 = base.I32_div_s((v483+v460)*int32(_a_F_timestamptz_part_common_16), int32(256))
									v492 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v501 = int32(_a_F_timestamptz_part_common_15) + v492
									v506 = base.I32_div_s(v501, int32(4))
									v509 = base.I32_div_s(v501, int32(-100))
									v512 = base.I32_div_s(v501, int32(400))
									v521 = base.I32_div_s(int32(_a_F_timestamptz_part_common_18), int32(256))
									v645 = base.I64_extend_i32_s(v461 + v468*int32(365) + v473 + v476 + v479 + v488 - int32(_a_F_timestamptz_part_common_17) - (int32(1) + v501*int32(365) + v506 + v509 + v512 + v521 - int32(_a_F_timestamptz_part_common_17)) + int32(1))
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 30:
									v133 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
									v136 = base.I32_div_s(int32(0)-v133, int32(3600))
									v645 = base.I64_extend_i32_s(v136)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 31:
									v125 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
									v127 = int32(60)
									v128 = base.I32_div_s(int32(0)-v125, v127)
									v130 = base.I32_rem_s(v128, v127)
									v645 = base.I64_extend_i32_s(v130)
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								case 32:
									v363 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
									v364 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
									v365 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
									v367 = F_date2j(m, v363, v364, v365)
									mBase = m.M
									v368 = int32(1)
									v370 = F_date2j(m, v363, v368, int32(4))
									mBase = m.M
									v373 = F_j2day(m, v370-v368)
									mBase = m.M
									if v367 < v370-v373 {
										v376 = int32(1)
										v377 = v363 - v376
										v380 = F_date2j(m, v377, v376, int32(4))
										mBase = m.M
										v383 = F_j2day(m, v380-v376)
										mBase = m.M
										v384 = v377
										v385 = v380
										v386 = v383
									} else {
										v384 = v363
										v385 = v370
										v386 = v373
									}
									if int32(357) <= v386+(v367-v385) {
										v391 = int32(1)
										v392 = v384 + v391
										v395 = F_date2j(m, v392, v391, int32(4))
										mBase = m.M
										v398 = F_j2day(m, v395-v391)
										mBase = m.M
										if v367 < v395-v398 {
											v401 = v384
										} else {
											v401 = v392
										}
										v404 = v401
									} else {
										v404 = v384
									}
									v645 = base.I64_extend_i32_s(v404) - base.I64_extend_i32_u(base.B2i32(v404 <= int32(0)))
									if l1 != 0 {
										v646 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v647 = m.ExcPending
										if v647 != 0 {
											return int64(0)
										} else {
											v656 = base.I64_extend_i32_u(v646)
											m.G0 = v14 + int32(96)
											return v656
										}
									} else {
										v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
										m.G0 = v14 + int32(96)
										return v656
									}
								}
							}
						}
					}
				} else {
					v551 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
					if v551 == int32(11) {
						v554 = F_SetEpochTimestamp(m)
						mBase = m.M
						v555 = m.ExcPending
						if v555 != 0 {
							return int64(0)
						} else {
							v557 = v554 + int64(9223372036854775807)
							if l1 != 0 {
								if v28 < v557 {
									v561 = F_int64_div_fast_to_numeric(m, v28-v554, int32(6))
									mBase = m.M
									v562 = m.ExcPending
									if v562 != 0 {
										return int64(0)
									} else {
										v656 = base.I64_extend_i32_u(v561)
										m.G0 = v14 + int32(96)
										return v656
									}
								} else {
									v566 = F_int64_to_numeric(m, v28)
									mBase = m.M
									v567 = m.ExcPending
									if v567 != 0 {
										return int64(0)
									} else {
										v568 = F_int64_to_numeric(m, v554)
										mBase = m.M
										v569 = m.ExcPending
										if v569 != 0 {
											return int64(0)
										} else {
											v571 = F_numeric_sub_safe(m, v566, v568, int32(0))
											mBase = m.M
											v572 = m.ExcPending
											if v572 != 0 {
												return int64(0)
											} else {
												v574 = F_int64_to_numeric(m, int64(1000000))
												mBase = m.M
												v575 = m.ExcPending
												if v575 != 0 {
													return int64(0)
												} else {
													v577 = F_numeric_div_safe(m, v571, v574, int32(0))
													mBase = m.M
													v578 = m.ExcPending
													if v578 != 0 {
														return int64(0)
													} else {
														v581 = F_DirectFunctionCall2Coll(m, int32(1390), int32(0), base.I64_extend_i32_u(v577), int64(6))
														mBase = m.M
														v582 = m.ExcPending
														if v582 != 0 {
															return int64(0)
														} else {
															v584 = F_pg_detoast_datum(m, base.I32_wrap_i64(v581))
															mBase = m.M
															v585 = m.ExcPending
															if v585 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v584)
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
								if v28 < v557 {
									v593 = base.F64_convert_i64_s(v28 - v554)
								} else {
									v593 = base.F64_sub(base.F64_convert_i64_s(v28), base.F64_convert_i64_s(v554))
								}
								v656 = base.I64_reinterpret_f64(base.F64_div(v593, float64(1e+06)))
								m.G0 = v14 + int32(96)
								return v656
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v600 = m.ExcPending
						if v600 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v603 = m.ExcPending
							if v603 != 0 {
								return int64(0)
							} else {
								v605 = F_format_type_be(m, int32(1184))
								mBase = m.M
								v606 = m.ExcPending
								if v606 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v605
									*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v57
									F_errmsg(m, int32(_a_F_timestamptz_part_common_12), v14+int32(32))
									mBase = m.M
									v613 = m.ExcPending
									if v613 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_19), int32(_a_F_timestamptz_part_common_9))
										mBase = m.M
										v618 = m.ExcPending
										if v618 != 0 {
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
				v76 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
				v80 = F_NonFiniteTimestampTzPart(m, v71, v76, v57, base.B2i32(v28 == int64(-9223372036854775807-1)), int32(1))
				mBase = m.M
				v81 = m.ExcPending
				if v81 != 0 {
					return int64(0)
				} else {
					if base.F64_ne(v80, float64(0)) != 0 {
						if l1 != 0 {
							if base.F64_lt(v80, float64(0)) != 0 {
								v91 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), int64(11926), int64(0), int64(-1))
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v656 = v91
									m.G0 = v14 + int32(96)
									return v656
								}
							} else {
								if base.F64_gt(v80, float64(0)) == int32(0) {
									if v71 != 0 {
										if v71 != int32(17) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v622 = m.ExcPending
											if v622 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v625 = m.ExcPending
												if v625 != 0 {
													return int64(0)
												} else {
													v627 = F_format_type_be(m, int32(1184))
													mBase = m.M
													v628 = m.ExcPending
													if v628 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v627
														*(*int32)(unsafe.Add(mBase, uint32(v14))) = v57
														F_errmsg(m, int32(_a_F_timestamptz_part_common_6), v14)
														mBase = m.M
														v633 = m.ExcPending
														if v633 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_8), int32(_a_F_timestamptz_part_common_9))
															mBase = m.M
															v638 = m.ExcPending
															if v638 != 0 {
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
											v117 = int32(0)
											v119 = F_timestamp2tm(m, v28, v14+int32(92), v14+int32(40), v14+int32(84), v117, v117)
											mBase = m.M
											v120 = m.ExcPending
											if v120 != 0 {
												return int64(0)
											} else {
												if v119 != 0 {
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v664 = m.ExcPending
													if v664 != 0 {
														return int64(0)
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v667 = m.ExcPending
														if v667 != 0 {
															return int64(0)
														} else {
															F_errmsg(m, int32(_a_F_timestamptz_part_common_10), int32(0))
															mBase = m.M
															v671 = m.ExcPending
															if v671 != 0 {
																return int64(0)
															} else {
																F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_11), int32(_a_F_timestamptz_part_common_9))
																mBase = m.M
																v676 = m.ExcPending
																if v676 != 0 {
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
													v121 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
													switch v121 - int32(4) {
													case 0:
														v640 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
														v645 = base.I64_extend_i32_s(int32(0) - v640)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													default:
														F_errstart_cold(m, int32(21), int32(0))
														mBase = m.M
														v532 = m.ExcPending
														if v532 != 0 {
															return int64(0)
														} else {
															F_errcode(m, int32(1088))
															mBase = m.M
															v535 = m.ExcPending
															if v535 != 0 {
																return int64(0)
															} else {
																v537 = F_format_type_be(m, int32(1184))
																mBase = m.M
																v538 = m.ExcPending
																if v538 != 0 {
																	return int64(0)
																} else {
																	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v537
																	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v57
																	F_errmsg(m, int32(_a_F_timestamptz_part_common_12), v14+int32(16))
																	mBase = m.M
																	v545 = m.ExcPending
																	if v545 != 0 {
																		return int64(0)
																	} else {
																		F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_13), int32(_a_F_timestamptz_part_common_9))
																		mBase = m.M
																		v550 = m.ExcPending
																		if v550 != 0 {
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
													case 14:
														v162 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
														if l1 != 0 {
															v163 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
															v169 = F_int64_div_fast_to_numeric(m, v163+base.I64_extend_i32_s(v162)*int64(1000000), int32(6))
															mBase = m.M
															v170 = m.ExcPending
															if v170 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v169)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v172 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
															v656 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_convert_i32_s(v172), float64(1e+06)), base.F64_convert_i32_s(v162)))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 15:
														v179 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+44)))
														v645 = v179
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 16:
														v180 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+48)))
														v645 = v180
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 17:
														v181 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+52)))
														v645 = v181
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 18:
														v191 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v192 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v193 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v195 = F_date2j(m, v191, v192, v193)
														mBase = m.M
														v196 = int32(1)
														v198 = F_date2j(m, v191, v196, int32(4))
														mBase = m.M
														v201 = F_j2day(m, v198-v196)
														mBase = m.M
														if v195 < v198-v201 {
															v204 = int32(1)
															v208 = F_date2j(m, v191-v204, v204, int32(4))
															mBase = m.M
															v211 = F_j2day(m, v208-v204)
															mBase = m.M
															v212 = v208
															v213 = v211
														} else {
															v212 = v198
															v213 = v201
														}
														v215 = v213 - v212 + v195
														if int32(357) <= v215 {
															v218 = int32(1)
															v222 = F_date2j(m, v191+v218, v218, int32(4))
															mBase = m.M
															v225 = F_j2day(m, v222-v218)
															mBase = m.M
															v226 = v222 - v225
															if v195 < v226 {
																v229 = v215
															} else {
																v229 = v195 - v226
															}
															v231 = v229
														} else {
															v231 = v215
														}
														v233 = base.I32_div_s(v231, int32(7))
														v645 = base.I64_extend_i32_s(v233 + int32(1))
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 19:
														v182 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+56)))
														v645 = v182
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 20:
														v183 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v184 = int32(1)
														v187 = base.I32_div_s(v183-v184, int32(3))
														v645 = base.I64_extend_i32_s(v187 + v184)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 21:
														v237 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														if int32(0) < v237 {
															v645 = base.I64_extend_i32_u(v237)
														} else {
															v645 = base.I64_extend_i32_s(v237 - int32(1))
														}
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 22:
														v244 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														if int32(0) < v244 {
															v248 = base.I32_div_u_s(v244, int32(10))
															v645 = base.I64_extend_i32_u(v248)
														} else {
															v253 = base.I32_div_s(int32(9)-v244, int32(-10))
															v645 = base.I64_extend_i32_s(v253)
														}
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 23:
														v255 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														if int32(0) < v255 {
															v261 = base.I32_div_s(v255+int32(99), int32(100))
															v645 = base.I64_extend_i32_s(v261)
														} else {
															v266 = base.I32_div_s(int32(100)-v255, int32(-100))
															v645 = base.I64_extend_i32_s(v266)
														}
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 24:
														v268 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														if int32(0) < v268 {
															v274 = base.I32_div_s(v268+int32(999), int32(1000))
															v645 = base.I64_extend_i32_s(v274)
														} else {
															v279 = base.I32_div_s(int32(1000)-v268, int32(-1000))
															v645 = base.I64_extend_i32_s(v279)
														}
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 25:
														v143 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
														if l1 != 0 {
															v144 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
															v150 = F_int64_div_fast_to_numeric(m, v144+base.I64_extend_i32_s(v143)*int64(1000000), int32(3))
															mBase = m.M
															v151 = m.ExcPending
															if v151 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v150)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v154 = float64(1000)
															v156 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
															v656 = base.I64_reinterpret_f64(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v143), v154), base.F64_div(base.F64_convert_i32_s(v156), v154)))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 26:
														v138 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
														v139 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+40)))
														v645 = v138 + v139*int64(1000000)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 27:
														v281 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v282 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v283 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v288 = base.B2i32(int32(2) < v282)
														if int32(2) < v282 {
															v289 = int32(_a_F_timestamptz_part_common_14)
														} else {
															v289 = int32(_a_F_timestamptz_part_common_15)
														}
														v290 = v289 + v281
														v295 = base.I32_div_s(v290, int32(4))
														v298 = base.I32_div_s(v290, int32(-100))
														v301 = base.I32_div_s(v290, int32(400))
														if int32(2) < v282 {
															v305 = int32(1)
														} else {
															v305 = int32(13)
														}
														v310 = base.I32_div_s((v305+v282)*int32(_a_F_timestamptz_part_common_16), int32(256))
														v313 = v283 + v290*int32(365) + v295 + v298 + v301 + v310 - int32(_a_F_timestamptz_part_common_17)
														if l1 != 0 {
															v315 = F_int64_to_numeric(m, base.I64_extend_i32_s(v313))
															mBase = m.M
															v316 = m.ExcPending
															if v316 != 0 {
																return int64(0)
															} else {
																v317 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
																v318 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
																v319 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
																v320 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
																v321 = int32(60)
																v331 = F_int64_to_numeric(m, v317+base.I64_extend_i32_s(v318+(v319+v320*v321)*v321)*int64(1000000))
																mBase = m.M
																v332 = m.ExcPending
																if v332 != 0 {
																	return int64(0)
																} else {
																	v334 = F_int64_to_numeric(m, int64(86400000000))
																	mBase = m.M
																	v335 = m.ExcPending
																	if v335 != 0 {
																		return int64(0)
																	} else {
																		v337 = F_numeric_div_safe(m, v331, v334, int32(0))
																		mBase = m.M
																		v338 = m.ExcPending
																		if v338 != 0 {
																			return int64(0)
																		} else {
																			v340 = F_numeric_add_safe(m, v315, v337, int32(0))
																			mBase = m.M
																			v341 = m.ExcPending
																			if v341 != 0 {
																				return int64(0)
																			} else {
																				v656 = base.I64_extend_i32_u(v340)
																				m.G0 = v14 + int32(96)
																				return v656
																			}
																		}
																	}
																}
															}
														} else {
															v343 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
															v347 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
															v348 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
															v349 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
															v350 = int32(60)
															v656 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v343), float64(1e+06)), base.F64_convert_i32_s(v347+(v348+v349*v350)*v350)), float64(86400)), base.F64_convert_i32_s(v313)))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 28, 33:
														v410 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v411 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v412 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v417 = base.B2i32(int32(2) < v411)
														if int32(2) < v411 {
															v418 = int32(_a_F_timestamptz_part_common_14)
														} else {
															v418 = int32(_a_F_timestamptz_part_common_15)
														}
														v419 = v418 + v410
														v424 = base.I32_div_s(v419, int32(4))
														v427 = base.I32_div_s(v419, int32(-100))
														v430 = base.I32_div_s(v419, int32(400))
														if int32(2) < v411 {
															v434 = int32(1)
														} else {
															v434 = int32(13)
														}
														v439 = base.I32_div_s((v434+v411)*int32(_a_F_timestamptz_part_common_16), int32(256))
														v445 = int32(7)
														v446 = base.I32_rem_s(v412+v419*int32(365)+v424+v427+v430+v439-int32(_a_F_timestamptz_part_common_17)+int32(1), v445)
														if v446 < int32(0) {
															v451 = v446 + v445
														} else {
															v451 = v446
														}
														v452 = base.I64_extend_i32_s(v451)
														if v451 != 0 {
															v454 = v452
														} else {
															v454 = int64(7)
														}
														v455 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
														if v455 == int32(37) {
															v458 = v454
														} else {
															v458 = v452
														}
														v645 = v458
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 29:
														v459 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v460 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v461 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v466 = base.B2i32(int32(2) < v460)
														if int32(2) < v460 {
															v467 = int32(_a_F_timestamptz_part_common_14)
														} else {
															v467 = int32(_a_F_timestamptz_part_common_15)
														}
														v468 = v467 + v459
														v473 = base.I32_div_s(v468, int32(4))
														v476 = base.I32_div_s(v468, int32(-100))
														v479 = base.I32_div_s(v468, int32(400))
														if int32(2) < v460 {
															v483 = int32(1)
														} else {
															v483 = int32(13)
														}
														v488 = base.I32_div_s((v483+v460)*int32(_a_F_timestamptz_part_common_16), int32(256))
														v492 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v501 = int32(_a_F_timestamptz_part_common_15) + v492
														v506 = base.I32_div_s(v501, int32(4))
														v509 = base.I32_div_s(v501, int32(-100))
														v512 = base.I32_div_s(v501, int32(400))
														v521 = base.I32_div_s(int32(_a_F_timestamptz_part_common_18), int32(256))
														v645 = base.I64_extend_i32_s(v461 + v468*int32(365) + v473 + v476 + v479 + v488 - int32(_a_F_timestamptz_part_common_17) - (int32(1) + v501*int32(365) + v506 + v509 + v512 + v521 - int32(_a_F_timestamptz_part_common_17)) + int32(1))
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 30:
														v133 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
														v136 = base.I32_div_s(int32(0)-v133, int32(3600))
														v645 = base.I64_extend_i32_s(v136)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 31:
														v125 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
														v127 = int32(60)
														v128 = base.I32_div_s(int32(0)-v125, v127)
														v130 = base.I32_rem_s(v128, v127)
														v645 = base.I64_extend_i32_s(v130)
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													case 32:
														v363 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
														v364 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
														v365 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
														v367 = F_date2j(m, v363, v364, v365)
														mBase = m.M
														v368 = int32(1)
														v370 = F_date2j(m, v363, v368, int32(4))
														mBase = m.M
														v373 = F_j2day(m, v370-v368)
														mBase = m.M
														if v367 < v370-v373 {
															v376 = int32(1)
															v377 = v363 - v376
															v380 = F_date2j(m, v377, v376, int32(4))
															mBase = m.M
															v383 = F_j2day(m, v380-v376)
															mBase = m.M
															v384 = v377
															v385 = v380
															v386 = v383
														} else {
															v384 = v363
															v385 = v370
															v386 = v373
														}
														if int32(357) <= v386+(v367-v385) {
															v391 = int32(1)
															v392 = v384 + v391
															v395 = F_date2j(m, v392, v391, int32(4))
															mBase = m.M
															v398 = F_j2day(m, v395-v391)
															mBase = m.M
															if v367 < v395-v398 {
																v401 = v384
															} else {
																v401 = v392
															}
															v404 = v401
														} else {
															v404 = v384
														}
														v645 = base.I64_extend_i32_s(v404) - base.I64_extend_i32_u(base.B2i32(v404 <= int32(0)))
														if l1 != 0 {
															v646 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v647 = m.ExcPending
															if v647 != 0 {
																return int64(0)
															} else {
																v656 = base.I64_extend_i32_u(v646)
																m.G0 = v14 + int32(96)
																return v656
															}
														} else {
															v656 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v645))
															m.G0 = v14 + int32(96)
															return v656
														}
													}
												}
											}
										}
									} else {
										v551 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
										if v551 == int32(11) {
											v554 = F_SetEpochTimestamp(m)
											mBase = m.M
											v555 = m.ExcPending
											if v555 != 0 {
												return int64(0)
											} else {
												v557 = v554 + int64(9223372036854775807)
												if l1 != 0 {
													if v28 < v557 {
														v561 = F_int64_div_fast_to_numeric(m, v28-v554, int32(6))
														mBase = m.M
														v562 = m.ExcPending
														if v562 != 0 {
															return int64(0)
														} else {
															v656 = base.I64_extend_i32_u(v561)
															m.G0 = v14 + int32(96)
															return v656
														}
													} else {
														v566 = F_int64_to_numeric(m, v28)
														mBase = m.M
														v567 = m.ExcPending
														if v567 != 0 {
															return int64(0)
														} else {
															v568 = F_int64_to_numeric(m, v554)
															mBase = m.M
															v569 = m.ExcPending
															if v569 != 0 {
																return int64(0)
															} else {
																v571 = F_numeric_sub_safe(m, v566, v568, int32(0))
																mBase = m.M
																v572 = m.ExcPending
																if v572 != 0 {
																	return int64(0)
																} else {
																	v574 = F_int64_to_numeric(m, int64(1000000))
																	mBase = m.M
																	v575 = m.ExcPending
																	if v575 != 0 {
																		return int64(0)
																	} else {
																		v577 = F_numeric_div_safe(m, v571, v574, int32(0))
																		mBase = m.M
																		v578 = m.ExcPending
																		if v578 != 0 {
																			return int64(0)
																		} else {
																			v581 = F_DirectFunctionCall2Coll(m, int32(1390), int32(0), base.I64_extend_i32_u(v577), int64(6))
																			mBase = m.M
																			v582 = m.ExcPending
																			if v582 != 0 {
																				return int64(0)
																			} else {
																				v584 = F_pg_detoast_datum(m, base.I32_wrap_i64(v581))
																				mBase = m.M
																				v585 = m.ExcPending
																				if v585 != 0 {
																					return int64(0)
																				} else {
																					v656 = base.I64_extend_i32_u(v584)
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
													if v28 < v557 {
														v593 = base.F64_convert_i64_s(v28 - v554)
													} else {
														v593 = base.F64_sub(base.F64_convert_i64_s(v28), base.F64_convert_i64_s(v554))
													}
													v656 = base.I64_reinterpret_f64(base.F64_div(v593, float64(1e+06)))
													m.G0 = v14 + int32(96)
													return v656
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v600 = m.ExcPending
											if v600 != 0 {
												return int64(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v603 = m.ExcPending
												if v603 != 0 {
													return int64(0)
												} else {
													v605 = F_format_type_be(m, int32(1184))
													mBase = m.M
													v606 = m.ExcPending
													if v606 != 0 {
														return int64(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v605
														*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v57
														F_errmsg(m, int32(_a_F_timestamptz_part_common_12), v14+int32(32))
														mBase = m.M
														v613 = m.ExcPending
														if v613 != 0 {
															return int64(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamptz_part_common_7), int32(_a_F_timestamptz_part_common_19), int32(_a_F_timestamptz_part_common_9))
															mBase = m.M
															v618 = m.ExcPending
															if v618 != 0 {
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
									v102 = F_DirectFunctionCall3Coll(m, int32(434), int32(0), int64(11937), int64(0), int64(-1))
									mBase = m.M
									v103 = m.ExcPending
									if v103 != 0 {
										return int64(0)
									} else {
										v656 = v102
										m.G0 = v14 + int32(96)
										return v656
									}
								}
							}
						} else {
							v656 = base.I64_reinterpret_f64(v80)
							m.G0 = v14 + int32(96)
							return v656
						}
					} else {
						v105 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v105)
						v656 = int64(0)
						m.G0 = v14 + int32(96)
						return v656
					}
				}
			}
		}
	}
}
func F_timestamptz_random(m *base.Module, l0 int32) int64 {
	var v5 int64
	_ = v5
	var v8 int32
	_ = v8
	v5 = Fn14383(m, l0, int32(_a_F_timestamptz_random_0), int32(266), int32(261))
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		return v5
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
func F_timestamptz_trunc_zone(m *base.Module, l0 int32) int64 {
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
	var v14 int64
	_ = v14
	var v15 int32
	_ = v15
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
	var v23 int64
	_ = v23
	var v24 int32
	_ = v24
	v5 = m.G0
	v7 = v5 - int32(256)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v10 = F_pg_detoast_datum_packed(m, v9)
	mBase = m.M
	v13 = m.ExcPending
	if v13 != 0 {
		return int64(0)
	} else {
		v14 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v16 = F_pg_detoast_datum_packed(m, v15)
		mBase = m.M
		v17 = m.ExcPending
		if v17 != 0 {
			return int64(0)
		} else {
			F_text_to_cstring_buffer(m, v16, v7, int32(256))
			mBase = m.M
			v20 = m.ExcPending
			if v20 != 0 {
				return int64(0)
			} else {
				v21 = F_DecodeTimezoneNameToTz(m, v7)
				mBase = m.M
				v22 = m.ExcPending
				if v22 != 0 {
					return int64(0)
				} else {
					v23 = F_timestamptz_trunc_internal(m, v10, v14, v21)
					mBase = m.M
					v24 = m.ExcPending
					if v24 != 0 {
						return int64(0)
					} else {
						m.G0 = v7 + int32(256)
						return v23
					}
				}
			}
		}
	}
}
