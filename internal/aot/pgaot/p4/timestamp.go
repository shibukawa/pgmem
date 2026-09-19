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
	var v58 int32
	_ = v58
	var v61 int64
	_ = v61
	var v62 int32
	_ = v62
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v80 int32
	_ = v80
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v114 int64
	_ = v114
	var v123 int64
	_ = v123
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v132 int64
	_ = v132
	var v133 int64
	_ = v133
	var v137 int64
	_ = v137
	var v144 int64
	_ = v144
	var v155 int64
	_ = v155
	var v156 int64
	_ = v156
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v172 int64
	_ = v172
	var v173 int64
	_ = v173
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v194 int32
	_ = v194
	var v200 int32
	_ = v200
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v212 int32
	_ = v212
	var v216 int32
	_ = v216
	var v221 int32
	_ = v221
	var v222 int64
	_ = v222
	var v223 int32
	_ = v223
	var v224 int64
	_ = v224
	var v234 int32
	_ = v234
	var v235 int64
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	v12 = m.G0
	v14 = v12 - int32(512)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
			return int32(0)
		} else {
			if v42 == int32(0) {
				v58 = *(*int32)(unsafe.Add(mBase, uint32(v14)+448))
				switch v58 - int32(2) {
				case 0:
					v61 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+500)))
					v62 = *(*int32)(unsafe.Add(mBase, uint32(v14)+476))
					if v62 <= int32(-4713) {
						if v62 != int32(-4713) {
							v187 = int32(0)
							v188 = F_errsave_start(m, v16)
							mBase = m.M
							v189 = m.ExcPending
							if v189 != 0 {
								return int32(0)
							} else {
								if v188 == int32(0) {
									v238 = v187
									m.G0 = v14 + int32(512)
									return v238
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v194 = m.ExcPending
									if v194 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
										F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
										mBase = m.M
										v200 = m.ExcPending
										if v200 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
											mBase = m.M
											v205 = m.ExcPending
											if v205 != 0 {
												return int32(0)
											} else {
												v238 = v187
												m.G0 = v14 + int32(512)
												return v238
											}
										}
									}
								}
							}
						} else {
							v67 = *(*int32)(unsafe.Add(mBase, uint32(v14)+472))
							if int32(10) < v67 {
								v78 = v67
								v80 = v14 + int32(32)
								v81 = *(*int32)(unsafe.Add(mBase, uint32(v14)+468))
								v86 = base.B2i32(int32(2) < v78)
								if int32(2) < v78 {
									v87 = int32(_a_F_timestamp_in_3)
								} else {
									v87 = int32(_a_F_timestamp_in_4)
								}
								v88 = v87 + v62
								v93 = base.I32_div_s(v88, int32(4))
								v96 = base.I32_div_s(v88, int32(-100))
								v99 = base.I32_div_s(v88, int32(400))
								if int32(2) < v78 {
									v103 = int32(1)
								} else {
									v103 = int32(13)
								}
								v108 = base.I32_div_s((v103+v78)*int32(_a_F_timestamp_in_5), int32(256))
								v114 = base.I64_extend_i32_s(v81 + v88*int32(365) + v93 + v96 + v99 + v108 - int32(_a_F_timestamp_in_6) - int32(_a_F_timestamp_in_7))
								v123 = int64(32)
								v124 = int64(20)
								v126 = int64(base.Ui64(v114) >> (uint(v123) % 64))
								v129 = int64(4294967295)
								v130 = int64(500654080)
								v132 = v114 & v129
								v133 = v130 * v132
								v137 = int64(base.Ui64(v133)>>(uint(v123)%64)) + v130*v126
								v144 = v132*v124 + v137&v129
								*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v114*int64(0) + v114>>(uint(int64(63))%64)*int64(86400000000) + v124*v126 + int64(base.Ui64(v137)>>(uint(v123)%64)) + int64(base.Ui64(v144)>>(uint(v123)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v80))) = v133&v129 | v144<<(uint(v123)%64)
								v155 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
								v156 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
								if v155 != v156>>(uint(int64(63))%64) {
									v187 = int32(0)
									v188 = F_errsave_start(m, v16)
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int32(0)
									} else {
										if v188 == int32(0) {
											v238 = v187
											m.G0 = v14 + int32(512)
											return v238
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
												F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
												mBase = m.M
												v200 = m.ExcPending
												if v200 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														v238 = v187
														m.G0 = v14 + int32(512)
														return v238
													}
												}
											}
										}
									}
								} else {
									v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+456))
									v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+460))
									v162 = *(*int32)(unsafe.Add(mBase, uint32(v14)+464))
									v163 = int32(60)
									v172 = base.I64_extend_i32_s(v160+(v161+v162*v163)*v163)*int64(1000000) + v61
									v173 = v156 + v172
									*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v173
									if base.B2i32(v172 < int64(0))^base.B2i32(v173 < v156) != 0 {
										v187 = int32(0)
										v188 = F_errsave_start(m, v16)
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return int32(0)
										} else {
											if v188 == int32(0) {
												v238 = v187
												m.G0 = v14 + int32(512)
												return v238
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
													F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
													mBase = m.M
													v200 = m.ExcPending
													if v200 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return int32(0)
														} else {
															v238 = v187
															m.G0 = v14 + int32(512)
															return v238
														}
													}
												}
											}
										}
									} else {
										if base.Ui64(int64(9011559254509551615)) < base.Ui64(v173-int64(9223371331200000000)) {
											F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
											mBase = m.M
											v234 = m.ExcPending
											if v234 != 0 {
												return int32(0)
											} else {
												v235 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
												v236 = F_Int64GetDatum(m, v235)
												mBase = m.M
												v237 = m.ExcPending
												if v237 != 0 {
													return int32(0)
												} else {
													v238 = v236
													m.G0 = v14 + int32(512)
													return v238
												}
											}
										} else {
											v187 = int32(0)
											v188 = F_errsave_start(m, v16)
											mBase = m.M
											v189 = m.ExcPending
											if v189 != 0 {
												return int32(0)
											} else {
												if v188 == int32(0) {
													v238 = v187
													m.G0 = v14 + int32(512)
													return v238
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
														F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
														mBase = m.M
														v200 = m.ExcPending
														if v200 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return int32(0)
															} else {
																v238 = v187
																m.G0 = v14 + int32(512)
																return v238
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v187 = int32(0)
								v188 = F_errsave_start(m, v16)
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
									return int32(0)
								} else {
									if v188 == int32(0) {
										v238 = v187
										m.G0 = v14 + int32(512)
										return v238
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
											F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return int32(0)
												} else {
													v238 = v187
													m.G0 = v14 + int32(512)
													return v238
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v62 <= int32(_a_F_timestamp_in_8) {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(v14)+472))
							v78 = v72
							v80 = v14 + int32(32)
							v81 = *(*int32)(unsafe.Add(mBase, uint32(v14)+468))
							v86 = base.B2i32(int32(2) < v78)
							if int32(2) < v78 {
								v87 = int32(_a_F_timestamp_in_3)
							} else {
								v87 = int32(_a_F_timestamp_in_4)
							}
							v88 = v87 + v62
							v93 = base.I32_div_s(v88, int32(4))
							v96 = base.I32_div_s(v88, int32(-100))
							v99 = base.I32_div_s(v88, int32(400))
							if int32(2) < v78 {
								v103 = int32(1)
							} else {
								v103 = int32(13)
							}
							v108 = base.I32_div_s((v103+v78)*int32(_a_F_timestamp_in_5), int32(256))
							v114 = base.I64_extend_i32_s(v81 + v88*int32(365) + v93 + v96 + v99 + v108 - int32(_a_F_timestamp_in_6) - int32(_a_F_timestamp_in_7))
							v123 = int64(32)
							v124 = int64(20)
							v126 = int64(base.Ui64(v114) >> (uint(v123) % 64))
							v129 = int64(4294967295)
							v130 = int64(500654080)
							v132 = v114 & v129
							v133 = v130 * v132
							v137 = int64(base.Ui64(v133)>>(uint(v123)%64)) + v130*v126
							v144 = v132*v124 + v137&v129
							*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v114*int64(0) + v114>>(uint(int64(63))%64)*int64(86400000000) + v124*v126 + int64(base.Ui64(v137)>>(uint(v123)%64)) + int64(base.Ui64(v144)>>(uint(v123)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v80))) = v133&v129 | v144<<(uint(v123)%64)
							v155 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
							v156 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
							if v155 != v156>>(uint(int64(63))%64) {
								v187 = int32(0)
								v188 = F_errsave_start(m, v16)
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
									return int32(0)
								} else {
									if v188 == int32(0) {
										v238 = v187
										m.G0 = v14 + int32(512)
										return v238
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
											F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return int32(0)
												} else {
													v238 = v187
													m.G0 = v14 + int32(512)
													return v238
												}
											}
										}
									}
								}
							} else {
								v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+456))
								v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+460))
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v14)+464))
								v163 = int32(60)
								v172 = base.I64_extend_i32_s(v160+(v161+v162*v163)*v163)*int64(1000000) + v61
								v173 = v156 + v172
								*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v173
								if base.B2i32(v172 < int64(0))^base.B2i32(v173 < v156) != 0 {
									v187 = int32(0)
									v188 = F_errsave_start(m, v16)
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int32(0)
									} else {
										if v188 == int32(0) {
											v238 = v187
											m.G0 = v14 + int32(512)
											return v238
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
												F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
												mBase = m.M
												v200 = m.ExcPending
												if v200 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														v238 = v187
														m.G0 = v14 + int32(512)
														return v238
													}
												}
											}
										}
									}
								} else {
									if base.Ui64(int64(9011559254509551615)) < base.Ui64(v173-int64(9223371331200000000)) {
										F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
										mBase = m.M
										v234 = m.ExcPending
										if v234 != 0 {
											return int32(0)
										} else {
											v235 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
											v236 = F_Int64GetDatum(m, v235)
											mBase = m.M
											v237 = m.ExcPending
											if v237 != 0 {
												return int32(0)
											} else {
												v238 = v236
												m.G0 = v14 + int32(512)
												return v238
											}
										}
									} else {
										v187 = int32(0)
										v188 = F_errsave_start(m, v16)
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return int32(0)
										} else {
											if v188 == int32(0) {
												v238 = v187
												m.G0 = v14 + int32(512)
												return v238
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
													F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
													mBase = m.M
													v200 = m.ExcPending
													if v200 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return int32(0)
														} else {
															v238 = v187
															m.G0 = v14 + int32(512)
															return v238
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							if v62 != int32(_a_F_timestamp_in_9) {
								v187 = int32(0)
								v188 = F_errsave_start(m, v16)
								mBase = m.M
								v189 = m.ExcPending
								if v189 != 0 {
									return int32(0)
								} else {
									if v188 == int32(0) {
										v238 = v187
										m.G0 = v14 + int32(512)
										return v238
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v194 = m.ExcPending
										if v194 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
											F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
											mBase = m.M
											v200 = m.ExcPending
											if v200 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
												mBase = m.M
												v205 = m.ExcPending
												if v205 != 0 {
													return int32(0)
												} else {
													v238 = v187
													m.G0 = v14 + int32(512)
													return v238
												}
											}
										}
									}
								}
							} else {
								v75 = *(*int32)(unsafe.Add(mBase, uint32(v14)+472))
								if int32(5) < v75 {
									v187 = int32(0)
									v188 = F_errsave_start(m, v16)
									mBase = m.M
									v189 = m.ExcPending
									if v189 != 0 {
										return int32(0)
									} else {
										if v188 == int32(0) {
											v238 = v187
											m.G0 = v14 + int32(512)
											return v238
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v194 = m.ExcPending
											if v194 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
												F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
												mBase = m.M
												v200 = m.ExcPending
												if v200 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
													mBase = m.M
													v205 = m.ExcPending
													if v205 != 0 {
														return int32(0)
													} else {
														v238 = v187
														m.G0 = v14 + int32(512)
														return v238
													}
												}
											}
										}
									}
								} else {
									v78 = v75
									v80 = v14 + int32(32)
									v81 = *(*int32)(unsafe.Add(mBase, uint32(v14)+468))
									v86 = base.B2i32(int32(2) < v78)
									if int32(2) < v78 {
										v87 = int32(_a_F_timestamp_in_3)
									} else {
										v87 = int32(_a_F_timestamp_in_4)
									}
									v88 = v87 + v62
									v93 = base.I32_div_s(v88, int32(4))
									v96 = base.I32_div_s(v88, int32(-100))
									v99 = base.I32_div_s(v88, int32(400))
									if int32(2) < v78 {
										v103 = int32(1)
									} else {
										v103 = int32(13)
									}
									v108 = base.I32_div_s((v103+v78)*int32(_a_F_timestamp_in_5), int32(256))
									v114 = base.I64_extend_i32_s(v81 + v88*int32(365) + v93 + v96 + v99 + v108 - int32(_a_F_timestamp_in_6) - int32(_a_F_timestamp_in_7))
									v123 = int64(32)
									v124 = int64(20)
									v126 = int64(base.Ui64(v114) >> (uint(v123) % 64))
									v129 = int64(4294967295)
									v130 = int64(500654080)
									v132 = v114 & v129
									v133 = v130 * v132
									v137 = int64(base.Ui64(v133)>>(uint(v123)%64)) + v130*v126
									v144 = v132*v124 + v137&v129
									*(*int64)(unsafe.Add(mBase, uint32(v80)+8)) = v114*int64(0) + v114>>(uint(int64(63))%64)*int64(86400000000) + v124*v126 + int64(base.Ui64(v137)>>(uint(v123)%64)) + int64(base.Ui64(v144)>>(uint(v123)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v80))) = v133&v129 | v144<<(uint(v123)%64)
									v155 = *(*int64)(unsafe.Add(mBase, uint32(v14)+40))
									v156 = *(*int64)(unsafe.Add(mBase, uint32(v14)+32))
									if v155 != v156>>(uint(int64(63))%64) {
										v187 = int32(0)
										v188 = F_errsave_start(m, v16)
										mBase = m.M
										v189 = m.ExcPending
										if v189 != 0 {
											return int32(0)
										} else {
											if v188 == int32(0) {
												v238 = v187
												m.G0 = v14 + int32(512)
												return v238
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v194 = m.ExcPending
												if v194 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
													F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
													mBase = m.M
													v200 = m.ExcPending
													if v200 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
														mBase = m.M
														v205 = m.ExcPending
														if v205 != 0 {
															return int32(0)
														} else {
															v238 = v187
															m.G0 = v14 + int32(512)
															return v238
														}
													}
												}
											}
										}
									} else {
										v160 = *(*int32)(unsafe.Add(mBase, uint32(v14)+456))
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v14)+460))
										v162 = *(*int32)(unsafe.Add(mBase, uint32(v14)+464))
										v163 = int32(60)
										v172 = base.I64_extend_i32_s(v160+(v161+v162*v163)*v163)*int64(1000000) + v61
										v173 = v156 + v172
										*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v173
										if base.B2i32(v172 < int64(0))^base.B2i32(v173 < v156) != 0 {
											v187 = int32(0)
											v188 = F_errsave_start(m, v16)
											mBase = m.M
											v189 = m.ExcPending
											if v189 != 0 {
												return int32(0)
											} else {
												if v188 == int32(0) {
													v238 = v187
													m.G0 = v14 + int32(512)
													return v238
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v194 = m.ExcPending
													if v194 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
														F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
														mBase = m.M
														v200 = m.ExcPending
														if v200 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
															mBase = m.M
															v205 = m.ExcPending
															if v205 != 0 {
																return int32(0)
															} else {
																v238 = v187
																m.G0 = v14 + int32(512)
																return v238
															}
														}
													}
												}
											}
										} else {
											if base.Ui64(int64(9011559254509551615)) < base.Ui64(v173-int64(9223371331200000000)) {
												F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
												mBase = m.M
												v234 = m.ExcPending
												if v234 != 0 {
													return int32(0)
												} else {
													v235 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
													v236 = F_Int64GetDatum(m, v235)
													mBase = m.M
													v237 = m.ExcPending
													if v237 != 0 {
														return int32(0)
													} else {
														v238 = v236
														m.G0 = v14 + int32(512)
														return v238
													}
												}
											} else {
												v187 = int32(0)
												v188 = F_errsave_start(m, v16)
												mBase = m.M
												v189 = m.ExcPending
												if v189 != 0 {
													return int32(0)
												} else {
													if v188 == int32(0) {
														v238 = v187
														m.G0 = v14 + int32(512)
														return v238
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v194 = m.ExcPending
														if v194 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v18
															F_errmsg(m, int32(_a_F_timestamp_in_0), v14+int32(16))
															mBase = m.M
															v200 = m.ExcPending
															if v200 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v16, int32(_a_F_timestamp_in_1), int32(204), int32(_a_F_timestamp_in_2))
																mBase = m.M
																v205 = m.ExcPending
																if v205 != 0 {
																	return int32(0)
																} else {
																	v238 = v187
																	m.G0 = v14 + int32(512)
																	return v238
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
					v210 = m.ExcPending
					if v210 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v18
						v212 = *(*int32)(unsafe.Add(mBase, uint32(v14)+448))
						*(*int32)(unsafe.Add(mBase, uint32(v14))) = v212
						F_errmsg_internal(m, int32(_a_F_timestamp_in_10), v14)
						mBase = m.M
						v216 = m.ExcPending
						if v216 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_timestamp_in_1), int32(221), int32(_a_F_timestamp_in_2))
							mBase = m.M
							v221 = m.ExcPending
							if v221 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 7:
					v224 = int64(-9223372036854775807 - 1)
					*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v224
					F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
					mBase = m.M
					v234 = m.ExcPending
					if v234 != 0 {
						return int32(0)
					} else {
						v235 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
						v236 = F_Int64GetDatum(m, v235)
						mBase = m.M
						v237 = m.ExcPending
						if v237 != 0 {
							return int32(0)
						} else {
							v238 = v236
							m.G0 = v14 + int32(512)
							return v238
						}
					}
				case 8:
					v224 = int64(9223372036854775807)
					*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v224
					F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
					mBase = m.M
					v234 = m.ExcPending
					if v234 != 0 {
						return int32(0)
					} else {
						v235 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
						v236 = F_Int64GetDatum(m, v235)
						mBase = m.M
						v237 = m.ExcPending
						if v237 != 0 {
							return int32(0)
						} else {
							v238 = v236
							m.G0 = v14 + int32(512)
							return v238
						}
					}
				case 9:
					v222 = F_SetEpochTimestamp(m)
					mBase = m.M
					v223 = m.ExcPending
					if v223 != 0 {
						return int32(0)
					} else {
						v224 = v222
						*(*int64)(unsafe.Add(mBase, uint32(v14)+504)) = v224
						F_AdjustTimestampForTypmod(m, v14+int32(504), v17, v16)
						mBase = m.M
						v234 = m.ExcPending
						if v234 != 0 {
							return int32(0)
						} else {
							v235 = *(*int64)(unsafe.Add(mBase, uint32(v14)+504))
							v236 = F_Int64GetDatum(m, v235)
							mBase = m.M
							v237 = m.ExcPending
							if v237 != 0 {
								return int32(0)
							} else {
								v238 = v236
								m.G0 = v14 + int32(512)
								return v238
							}
						}
					}
				}
			} else {
				v48 = v42
				F_DateTimeParseError(m, v48, v14+int32(56), v18, int32(_a_F_timestamp_in_11), v16)
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					v54 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
					v238 = int32(0)
					m.G0 = v14 + int32(512)
					return v238
				}
			}
		}
	} else {
		v48 = v28
		F_DateTimeParseError(m, v48, v14+int32(56), v18, int32(_a_F_timestamp_in_11), v16)
		mBase = m.M
		v53 = m.ExcPending
		if v53 != 0 {
			return int32(0)
		} else {
			v54 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v54)
			v238 = int32(0)
			m.G0 = v14 + int32(512)
			return v238
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
	var v37 int32
	_ = v37
	var v40 int32
	_ = v40
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v63 int32
	_ = v63
	var v67 int32
	_ = v67
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v84 float64
	_ = v84
	var v85 int32
	_ = v85
	var v91 int32
	_ = v91
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v102 int32
	_ = v102
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v110 int32
	_ = v110
	var v117 int64
	_ = v117
	var v125 int64
	_ = v125
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v132 int32
	_ = v132
	var v144 int32
	_ = v144
	var v145 int32
	_ = v145
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v158 int32
	_ = v158
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v189 int32
	_ = v189
	var v193 int32
	_ = v193
	var v203 int32
	_ = v203
	var v211 int64
	_ = v211
	var v213 int64
	_ = v213
	var v219 int64
	_ = v219
	var v222 int64
	_ = v222
	var v224 int64
	_ = v224
	var v226 int64
	_ = v226
	var v229 int64
	_ = v229
	var v231 int64
	_ = v231
	var v232 int32
	_ = v232
	var v236 int64
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v247 int32
	_ = v247
	var v248 int32
	_ = v248
	var v250 float64
	_ = v250
	var v256 int32
	_ = v256
	var v257 int32
	_ = v257
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v273 int64
	_ = v273
	var v274 int64
	_ = v274
	var v275 int32
	_ = v275
	var v276 int32
	_ = v276
	var v279 int32
	_ = v279
	var v283 int32
	_ = v283
	var v284 int32
	_ = v284
	var v285 int32
	_ = v285
	var v287 int32
	_ = v287
	var v288 int32
	_ = v288
	var v290 int32
	_ = v290
	var v293 int32
	_ = v293
	var v296 int32
	_ = v296
	var v300 int32
	_ = v300
	var v303 int32
	_ = v303
	var v304 int32
	_ = v304
	var v305 int32
	_ = v305
	var v307 int32
	_ = v307
	var v310 int32
	_ = v310
	var v314 int32
	_ = v314
	var v317 int32
	_ = v317
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v325 int32
	_ = v325
	var v329 int32
	_ = v329
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v353 int32
	_ = v353
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v366 int32
	_ = v366
	var v371 int32
	_ = v371
	var v373 int32
	_ = v373
	var v374 int32
	_ = v374
	var v375 int32
	_ = v375
	var v380 int32
	_ = v380
	var v381 int32
	_ = v381
	var v382 int32
	_ = v382
	var v387 int32
	_ = v387
	var v390 int32
	_ = v390
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v402 int32
	_ = v402
	var v405 int32
	_ = v405
	var v407 int32
	_ = v407
	var v408 int32
	_ = v408
	var v410 int32
	_ = v410
	var v411 int32
	_ = v411
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v426 int32
	_ = v426
	var v427 int32
	_ = v427
	var v429 int32
	_ = v429
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v433 int32
	_ = v433
	var v437 int32
	_ = v437
	var v438 int32
	_ = v438
	var v439 int32
	_ = v439
	var v440 int32
	_ = v440
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v456 int32
	_ = v456
	var v458 int32
	_ = v458
	var v459 int32
	_ = v459
	var v461 int32
	_ = v461
	var v464 int32
	_ = v464
	var v467 int32
	_ = v467
	var v468 int32
	_ = v468
	var v471 int32
	_ = v471
	var v474 int32
	_ = v474
	var v475 int32
	_ = v475
	var v476 int32
	_ = v476
	var v477 int32
	_ = v477
	var v482 int32
	_ = v482
	var v483 int32
	_ = v483
	var v486 int32
	_ = v486
	var v489 int32
	_ = v489
	var v492 int32
	_ = v492
	var v495 int32
	_ = v495
	var v501 int32
	_ = v501
	var v502 int32
	_ = v502
	var v503 int32
	_ = v503
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v510 int32
	_ = v510
	var v515 int32
	_ = v515
	var v518 int32
	_ = v518
	var v521 int32
	_ = v521
	var v525 int32
	_ = v525
	var v530 int32
	_ = v530
	var v536 int32
	_ = v536
	var v537 int32
	_ = v537
	var v542 int32
	_ = v542
	var v543 int64
	_ = v543
	var v545 int64
	_ = v545
	var v546 int32
	_ = v546
	var v549 int64
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v552 int32
	_ = v552
	var v557 int32
	_ = v557
	var v558 int32
	_ = v558
	var v559 int32
	_ = v559
	var v564 int32
	_ = v564
	var v567 int32
	_ = v567
	var v570 int32
	_ = v570
	var v574 int32
	_ = v574
	var v579 int32
	_ = v579
	var v583 int32
	_ = v583
	var v592 int32
	_ = v592
	var v597 int32
	_ = v597
	var v600 int32
	_ = v600
	var v603 int32
	_ = v603
	var v612 int32
	_ = v612
	var v623 int32
	_ = v623
	var v626 int32
	_ = v626
	var v628 int32
	_ = v628
	var v629 int32
	_ = v629
	var v636 int32
	_ = v636
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v645 int64
	_ = v645
	var v646 int32
	_ = v646
	var v648 int64
	_ = v648
	var v652 int32
	_ = v652
	var v653 int32
	_ = v653
	var v656 int32
	_ = v656
	var v657 int32
	_ = v657
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
	var v667 int32
	_ = v667
	var v668 int32
	_ = v668
	var v670 int32
	_ = v670
	var v671 int32
	_ = v671
	var v672 int32
	_ = v672
	var v673 int32
	_ = v673
	var v680 float64
	_ = v680
	var v683 int32
	_ = v683
	var v684 int32
	_ = v684
	var v688 int32
	_ = v688
	var v691 int32
	_ = v691
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v701 int32
	_ = v701
	var v706 int32
	_ = v706
	var v710 int32
	_ = v710
	var v713 int32
	_ = v713
	var v715 int32
	_ = v715
	var v716 int32
	_ = v716
	var v721 int32
	_ = v721
	var v726 int32
	_ = v726
	var v733 int64
	_ = v733
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v737 int32
	_ = v737
	var v738 int32
	_ = v738
	var v740 int32
	_ = v740
	var v753 int32
	_ = v753
	var v756 int32
	_ = v756
	var v760 int32
	_ = v760
	var v765 int32
	_ = v765
	v13 = m.G0
	v15 = v13 - int32(96)
	m.G0 = v15
	v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v18 = F_pg_detoast_datum_packed(m, v17)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		return int32(0)
	} else {
		v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v23 = *(*int64)(unsafe.Add(mBase, uint32(v22)))
		v24 = int32(1)
		v25 = v18 + v24
		v28 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v18))))
		v30 = v28 & v24
		if v30 != 0 {
			v31 = v25
		} else {
			v31 = v18 + int32(4)
		}
		if v28 == int32(1) {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v25))))
			if v37 == int32(18) {
				v40 = int32(16)
			} else {
				v40 = int32(0)
			}
			if base.Ui32((v37-int32(1))&int32(255)) < base.Ui32(int32(3)) {
				v47 = int32(4)
			} else {
				v47 = v40
			}
			v58 = v47
		} else {
			v48 = int32(1)
			if v30 != 0 {
				v58 = int32(base.Ui32(v28)>>(uint(v48)%32)) - v48
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
				v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v60 = F_downcase_truncate_identifier(m, v31, v58, int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			v63 = v15 + int32(92)
			v67 = Fn13846(m, v60, v63, int32(_a_F_timestamp_part_common_0), int32(_a_F_timestamp_part_common_1), int32(_a_F_timestamp_part_common_2))
			mBase = m.M
			if v67 == int32(31) {
				v73 = Fn13846(m, v60, v63, int32(_a_F_timestamp_part_common_3), int32(_a_F_timestamp_part_common_4), int32(_a_F_timestamp_part_common_5))
				mBase = m.M
				v74 = v73
			} else {
				v74 = v67
			}
			if base.Ui64(int64(1)) < base.Ui64(v23-int64(9223372036854775807)) {
				if v74 != 0 {
					if v74 != int32(17) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v710 = m.ExcPending
						if v710 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v713 = m.ExcPending
							if v713 != 0 {
								return int32(0)
							} else {
								v715 = F_format_type_be(m, int32(1114))
								mBase = m.M
								v716 = m.ExcPending
								if v716 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v715
									*(*int32)(unsafe.Add(mBase, uint32(v15))) = v60
									F_errmsg(m, int32(_a_F_timestamp_part_common_6), v15)
									mBase = m.M
									v721 = m.ExcPending
									if v721 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_8), int32(_a_F_timestamp_part_common_9))
										mBase = m.M
										v726 = m.ExcPending
										if v726 != 0 {
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
						v117 = base.I64_div_s(v23, int64(86400000000))
						if base.Ui64(int64(172799999999)) <= base.Ui64(v23+int64(86399999999)) {
							v125 = v117 * int64(-86400000000)
						} else {
							v125 = int64(0)
						}
						v126 = v125 + v23
						v129 = v126>>(uint(int64(63))%64) + v117
						if v129 <= int64(-2451546) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v753 = m.ExcPending
							if v753 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v756 = m.ExcPending
								if v756 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_timestamp_part_common_10), int32(0))
									mBase = m.M
									v760 = m.ExcPending
									if v760 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_11), int32(_a_F_timestamp_part_common_9))
										mBase = m.M
										v765 = m.ExcPending
										if v765 != 0 {
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
							v132 = base.I32_wrap_i64(v129)
							v144 = v132 + int32(_a_F_timestamp_part_common_12)
							v145 = int32(_a_F_timestamp_part_common_13)
							v146 = base.I32_div_u_s(v144, v145)
							v147 = int32(3)
							v153 = int32(2)
							v158 = base.I32_div_u_s((v146*int32(1073595727)+v144)<<(uint(v153)%32)|v147, v145)
							v161 = v132 + int32(_a_F_timestamp_part_common_14) + v146*v147 + v158 + int32(_a_F_timestamp_part_common_15)
							v162 = int32(1461)
							v163 = base.I32_div_u_s(v161, v162)
							v166 = v163*int32(-1461) + v161
							v168 = v166 << (uint(v153) % 32)
							if base.Ui32(v162) <= base.Ui32(v168) {
								v174 = base.I32_rem_u_s(v166+int32(305), int32(365))
								v179 = v174
							} else {
								v178 = base.I32_rem_u_s(v166+int32(306), int32(366))
								v179 = v178
							}
							v181 = base.I32_div_u_s(v168, int32(1461))
							*(*int32)(unsafe.Add(mBase, uint32(v15+int32(68)))) = v181 + v163<<(uint(int32(2))%32) - int32(_a_F_timestamp_part_common_16)
							v189 = v179 + int32(123)
							v193 = int32(base.Ui32(v189*int32(2141)) >> (uint(int32(16)) % 32))
							*(*int32)(unsafe.Add(mBase, uint32(v15+int32(60)))) = v189 - int32(base.Ui32(v193*int32(_a_F_timestamp_part_common_17))>>(uint(int32(8))%32))
							v203 = base.I32_rem_u_s(v193+int32(10), int32(12))
							*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v203 + int32(1)
							if v126 < int64(0) {
								v211 = v126 + int64(86400000000)
							} else {
								v211 = v126
							}
							v213 = base.I64_div_s(v211, int64(3600000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+56)) = uint32(v213)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = int32(0)
							*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = int64(4294967295)
							v219 = base.I64_extend32_s(v213)
							v222 = v219*int64(-3600000000) + v211
							v224 = base.I64_div_s(v222, int64(60000000))
							*(*uint32)(unsafe.Add(mBase, uint32(v15)+52)) = uint32(v224)
							v226 = base.I64_extend32_s(v224)
							v229 = v222 + v226*int64(-60000000)
							v231 = base.I64_div_s(v229, int64(1000000))
							v232 = base.I32_wrap_i64(v231)
							*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v232
							v236 = v231*int64(4293967296) + v229
							v237 = base.I32_wrap_i64(v236)
							v238 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
							switch v238 - int32(18) {
							case 0:
								if l1 != 0 {
									v264 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v236)+base.I64_extend32_s(v231)*int64(1000000), int32(6))
									mBase = m.M
									v265 = m.ExcPending
									if v265 != 0 {
										return int32(0)
									} else {
										v740 = v264
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v271 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v237), float64(1e+06)), base.F64_convert_i32_s(v232)))
									mBase = m.M
									v272 = m.ExcPending
									if v272 != 0 {
										return int32(0)
									} else {
										v740 = v271
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 1:
								v733 = v226
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 2:
								v733 = v219
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 3:
								v273 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+60)))
								v733 = v273
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 4:
								v283 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v284 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v285 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v287 = F_date2j(m, v283, v284, v285)
								mBase = m.M
								v288 = int32(1)
								v290 = F_date2j(m, v283, v288, int32(4))
								mBase = m.M
								v293 = F_j2day(m, v290-v288)
								mBase = m.M
								if v287 < v290-v293 {
									v296 = int32(1)
									v300 = F_date2j(m, v283-v296, v296, int32(4))
									mBase = m.M
									v303 = F_j2day(m, v300-v296)
									mBase = m.M
									v304 = v300
									v305 = v303
								} else {
									v304 = v290
									v305 = v293
								}
								v307 = v305 - v304 + v287
								if int32(357) <= v307 {
									v310 = int32(1)
									v314 = F_date2j(m, v283+v310, v310, int32(4))
									mBase = m.M
									v317 = F_j2day(m, v314-v310)
									mBase = m.M
									v318 = v314 - v317
									if v287 < v318 {
										v321 = v307
									} else {
										v321 = v287 - v318
									}
									v323 = v321
								} else {
									v323 = v307
								}
								v325 = base.I32_div_s(v323, int32(7))
								v733 = base.I64_extend_i32_s(v325 + int32(1))
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 5:
								v274 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+64)))
								v733 = v274
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 6:
								v275 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v276 = int32(1)
								v279 = base.I32_div_s(v275-v276, int32(3))
								v733 = base.I64_extend_i32_s(v279 + v276)
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 7:
								v329 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								if int32(0) < v329 {
									v733 = base.I64_extend_i32_u(v329)
								} else {
									v733 = base.I64_extend_i32_s(v329 - int32(1))
								}
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 8:
								v336 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								if int32(0) <= v336 {
									v340 = base.I32_div_u_s(v336, int32(10))
									v733 = base.I64_extend_i32_u(v340)
								} else {
									v345 = base.I32_div_s(int32(9)-v336, int32(-10))
									v733 = base.I64_extend_i32_s(v345)
								}
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 9:
								v347 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								if int32(0) < v347 {
									v353 = base.I32_div_s(v347+int32(99), int32(100))
									v733 = base.I64_extend_i32_s(v353)
								} else {
									v358 = base.I32_div_s(int32(100)-v347, int32(-100))
									v733 = base.I64_extend_i32_s(v358)
								}
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 10:
								v360 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								if int32(0) < v360 {
									v366 = base.I32_div_s(v360+int32(999), int32(1000))
									v733 = base.I64_extend_i32_s(v366)
								} else {
									v371 = base.I32_div_s(int32(1000)-v360, int32(-1000))
									v733 = base.I64_extend_i32_s(v371)
								}
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 11:
								if l1 != 0 {
									v247 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v236)+base.I64_extend32_s(v231)*int64(1000000), int32(3))
									mBase = m.M
									v248 = m.ExcPending
									if v248 != 0 {
										return int32(0)
									} else {
										v740 = v247
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v250 = float64(1000)
									v256 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v232), v250), base.F64_div(base.F64_convert_i32_s(v237), v250)))
									mBase = m.M
									v257 = m.ExcPending
									if v257 != 0 {
										return int32(0)
									} else {
										v740 = v256
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 12:
								v733 = base.I64_extend32_s(v236) + base.I64_extend32_s(v231)*int64(1000000)
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 13:
								v373 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v374 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v375 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v380 = base.B2i32(int32(2) < v374)
								if int32(2) < v374 {
									v381 = int32(_a_F_timestamp_part_common_16)
								} else {
									v381 = int32(_a_F_timestamp_part_common_18)
								}
								v382 = v381 + v373
								v387 = base.I32_div_s(v382, int32(4))
								v390 = base.I32_div_s(v382, int32(-100))
								v393 = base.I32_div_s(v382, int32(400))
								if int32(2) < v374 {
									v397 = int32(1)
								} else {
									v397 = int32(13)
								}
								v402 = base.I32_div_s((v397+v374)*int32(_a_F_timestamp_part_common_17), int32(256))
								v405 = v375 + v382*int32(365) + v387 + v390 + v393 + v402 - int32(_a_F_timestamp_part_common_19)
								if l1 != 0 {
									v407 = F_int64_to_numeric(m, base.I64_extend_i32_s(v405))
									mBase = m.M
									v408 = m.ExcPending
									if v408 != 0 {
										return int32(0)
									} else {
										v410 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
										v411 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
										v412 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
										v413 = int32(60)
										v423 = F_int64_to_numeric(m, base.I64_extend32_s(v236)+base.I64_extend_i32_s(v410+(v411+v412*v413)*v413)*int64(1000000))
										mBase = m.M
										v424 = m.ExcPending
										if v424 != 0 {
											return int32(0)
										} else {
											v426 = F_int64_to_numeric(m, int64(86400000000))
											mBase = m.M
											v427 = m.ExcPending
											if v427 != 0 {
												return int32(0)
											} else {
												v429 = F_numeric_div_opt_error(m, v423, v426, int32(0))
												mBase = m.M
												v430 = m.ExcPending
												if v430 != 0 {
													return int32(0)
												} else {
													v432 = F_numeric_add_opt_error(m, v407, v429, int32(0))
													mBase = m.M
													v433 = m.ExcPending
													if v433 != 0 {
														return int32(0)
													} else {
														v740 = v432
														m.G0 = v15 + int32(96)
														return v740
													}
												}
											}
										}
									}
								} else {
									v437 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
									v438 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
									v439 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
									v440 = int32(60)
									v452 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v237), float64(1e+06)), base.F64_convert_i32_s(v437+(v438+v439*v440)*v440)), float64(86400)), base.F64_convert_i32_s(v405)))
									mBase = m.M
									v453 = m.ExcPending
									if v453 != 0 {
										return int32(0)
									} else {
										v740 = v452
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 14, 19:
								v501 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v502 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v503 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v508 = base.B2i32(int32(2) < v502)
								if int32(2) < v502 {
									v509 = int32(_a_F_timestamp_part_common_16)
								} else {
									v509 = int32(_a_F_timestamp_part_common_18)
								}
								v510 = v509 + v501
								v515 = base.I32_div_s(v510, int32(4))
								v518 = base.I32_div_s(v510, int32(-100))
								v521 = base.I32_div_s(v510, int32(400))
								if int32(2) < v502 {
									v525 = int32(1)
								} else {
									v525 = int32(13)
								}
								v530 = base.I32_div_s((v525+v502)*int32(_a_F_timestamp_part_common_17), int32(256))
								v536 = int32(7)
								v537 = base.I32_rem_s(v503+v510*int32(365)+v515+v518+v521+v530-int32(_a_F_timestamp_part_common_19)+int32(1), v536)
								if v537 < int32(0) {
									v542 = v537 + v536
								} else {
									v542 = v537
								}
								v543 = base.I64_extend_i32_s(v542)
								if v542 != 0 {
									v545 = v543
								} else {
									v545 = int64(7)
								}
								v546 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
								if v546 == int32(37) {
									v549 = v545
								} else {
									v549 = v543
								}
								v733 = v549
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							case 15:
								v550 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v551 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v552 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v557 = base.B2i32(int32(2) < v551)
								if int32(2) < v551 {
									v558 = int32(_a_F_timestamp_part_common_16)
								} else {
									v558 = int32(_a_F_timestamp_part_common_18)
								}
								v559 = v558 + v550
								v564 = base.I32_div_s(v559, int32(4))
								v567 = base.I32_div_s(v559, int32(-100))
								v570 = base.I32_div_s(v559, int32(400))
								if int32(2) < v551 {
									v574 = int32(1)
								} else {
									v574 = int32(13)
								}
								v579 = base.I32_div_s((v574+v551)*int32(_a_F_timestamp_part_common_17), int32(256))
								v583 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v592 = int32(_a_F_timestamp_part_common_18) + v583
								v597 = base.I32_div_s(v592, int32(4))
								v600 = base.I32_div_s(v592, int32(-100))
								v603 = base.I32_div_s(v592, int32(400))
								v612 = base.I32_div_s(int32(_a_F_timestamp_part_common_20), int32(256))
								v733 = base.I64_extend_i32_s(v552 + v559*int32(365) + v564 + v567 + v570 + v579 - int32(_a_F_timestamp_part_common_19) - (int32(1) + v592*int32(365) + v597 + v600 + v603 + v612 - int32(_a_F_timestamp_part_common_19)) + int32(1))
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							default:
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v623 = m.ExcPending
								if v623 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v626 = m.ExcPending
									if v626 != 0 {
										return int32(0)
									} else {
										v628 = F_format_type_be(m, int32(1114))
										mBase = m.M
										v629 = m.ExcPending
										if v629 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v628
											*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v60
											F_errmsg(m, int32(_a_F_timestamp_part_common_21), v15+int32(16))
											mBase = m.M
											v636 = m.ExcPending
											if v636 != 0 {
												return int32(0)
											} else {
												F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_22), int32(_a_F_timestamp_part_common_9))
												mBase = m.M
												v641 = m.ExcPending
												if v641 != 0 {
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
							case 18:
								v454 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
								v455 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
								v456 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
								v458 = F_date2j(m, v454, v455, v456)
								mBase = m.M
								v459 = int32(1)
								v461 = F_date2j(m, v454, v459, int32(4))
								mBase = m.M
								v464 = F_j2day(m, v461-v459)
								mBase = m.M
								if v458 < v461-v464 {
									v467 = int32(1)
									v468 = v454 - v467
									v471 = F_date2j(m, v468, v467, int32(4))
									mBase = m.M
									v474 = F_j2day(m, v471-v467)
									mBase = m.M
									v475 = v468
									v476 = v471
									v477 = v474
								} else {
									v475 = v454
									v476 = v461
									v477 = v464
								}
								if int32(357) <= v477+v458-v476 {
									v482 = int32(1)
									v483 = v475 + v482
									v486 = F_date2j(m, v483, v482, int32(4))
									mBase = m.M
									v489 = F_j2day(m, v486-v482)
									mBase = m.M
									if v458 < v486-v489 {
										v492 = v475
									} else {
										v492 = v483
									}
									v495 = v492
								} else {
									v495 = v475
								}
								v733 = base.I64_extend_i32_s(v495) - base.I64_extend_i32_u(base.B2i32(v495 <= int32(0)))
								if l1 != 0 {
									v734 = F_int64_to_numeric(m, v733)
									mBase = m.M
									v735 = m.ExcPending
									if v735 != 0 {
										return int32(0)
									} else {
										v740 = v734
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
									mBase = m.M
									v738 = m.ExcPending
									if v738 != 0 {
										return int32(0)
									} else {
										v740 = v737
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							}
						}
					}
				} else {
					v642 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
					if v642 == int32(11) {
						v645 = F_SetEpochTimestamp(m)
						mBase = m.M
						v646 = m.ExcPending
						if v646 != 0 {
							return int32(0)
						} else {
							v648 = v645 + int64(9223372036854775807)
							if l1 != 0 {
								if v23 < v648 {
									v652 = F_int64_div_fast_to_numeric(m, v23-v645, int32(6))
									mBase = m.M
									v653 = m.ExcPending
									if v653 != 0 {
										return int32(0)
									} else {
										v740 = v652
										m.G0 = v15 + int32(96)
										return v740
									}
								} else {
									v656 = F_int64_to_numeric(m, v23)
									mBase = m.M
									v657 = m.ExcPending
									if v657 != 0 {
										return int32(0)
									} else {
										v658 = F_int64_to_numeric(m, v645)
										mBase = m.M
										v659 = m.ExcPending
										if v659 != 0 {
											return int32(0)
										} else {
											v661 = F_numeric_sub_opt_error(m, v656, v658, int32(0))
											mBase = m.M
											v662 = m.ExcPending
											if v662 != 0 {
												return int32(0)
											} else {
												v664 = F_int64_to_numeric(m, int64(1000000))
												mBase = m.M
												v665 = m.ExcPending
												if v665 != 0 {
													return int32(0)
												} else {
													v667 = F_numeric_div_opt_error(m, v661, v664, int32(0))
													mBase = m.M
													v668 = m.ExcPending
													if v668 != 0 {
														return int32(0)
													} else {
														v670 = F_DirectFunctionCall2Coll(m, int32(1259), int32(0), v667, int32(6))
														mBase = m.M
														v671 = m.ExcPending
														if v671 != 0 {
															return int32(0)
														} else {
															v672 = F_pg_detoast_datum(m, v670)
															mBase = m.M
															v673 = m.ExcPending
															if v673 != 0 {
																return int32(0)
															} else {
																v740 = v672
																m.G0 = v15 + int32(96)
																return v740
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								if v23 < v648 {
									v680 = base.F64_convert_i64_s(v23 - v645)
								} else {
									v680 = base.F64_sub(base.F64_convert_i64_s(v23), base.F64_convert_i64_s(v645))
								}
								v683 = F_Float8GetDatum(m, base.F64_div(v680, float64(1e+06)))
								mBase = m.M
								v684 = m.ExcPending
								if v684 != 0 {
									return int32(0)
								} else {
									v740 = v683
									m.G0 = v15 + int32(96)
									return v740
								}
							}
						}
					} else {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v688 = m.ExcPending
						if v688 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v691 = m.ExcPending
							if v691 != 0 {
								return int32(0)
							} else {
								v693 = F_format_type_be(m, int32(1114))
								mBase = m.M
								v694 = m.ExcPending
								if v694 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v693
									*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v60
									F_errmsg(m, int32(_a_F_timestamp_part_common_21), v15+int32(32))
									mBase = m.M
									v701 = m.ExcPending
									if v701 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_23), int32(_a_F_timestamp_part_common_9))
										mBase = m.M
										v706 = m.ExcPending
										if v706 != 0 {
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
				v79 = int32(0)
				v80 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
				v84 = F_NonFiniteTimestampTzPart(m, v74, v80, v60, base.B2i32(v23 == int64(-9223372036854775807-1)), v79)
				mBase = m.M
				v85 = m.ExcPending
				if v85 != 0 {
					return int32(0)
				} else {
					if base.F64_ne(v84, float64(0)) != 0 {
						if l1 != 0 {
							if base.F64_lt(v84, float64(0)) != 0 {
								v91 = int32(0)
								v95 = F_DirectFunctionCall3Coll(m, int32(408), v91, int32(_a_F_timestamp_part_common_24), v91, int32(-1))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v740 = v95
									m.G0 = v15 + int32(96)
									return v740
								}
							} else {
								if base.F64_gt(v84, float64(0)) == int32(0) {
									if v74 != 0 {
										if v74 != int32(17) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v710 = m.ExcPending
											if v710 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(50856066))
												mBase = m.M
												v713 = m.ExcPending
												if v713 != 0 {
													return int32(0)
												} else {
													v715 = F_format_type_be(m, int32(1114))
													mBase = m.M
													v716 = m.ExcPending
													if v716 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+4)) = v715
														*(*int32)(unsafe.Add(mBase, uint32(v15))) = v60
														F_errmsg(m, int32(_a_F_timestamp_part_common_6), v15)
														mBase = m.M
														v721 = m.ExcPending
														if v721 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_8), int32(_a_F_timestamp_part_common_9))
															mBase = m.M
															v726 = m.ExcPending
															if v726 != 0 {
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
											v117 = base.I64_div_s(v23, int64(86400000000))
											if base.Ui64(int64(172799999999)) <= base.Ui64(v23+int64(86399999999)) {
												v125 = v117 * int64(-86400000000)
											} else {
												v125 = int64(0)
											}
											v126 = v125 + v23
											v129 = v126>>(uint(int64(63))%64) + v117
											if v129 <= int64(-2451546) {
												F_errstart_cold(m, int32(21), int32(0))
												mBase = m.M
												v753 = m.ExcPending
												if v753 != 0 {
													return int32(0)
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v756 = m.ExcPending
													if v756 != 0 {
														return int32(0)
													} else {
														F_errmsg(m, int32(_a_F_timestamp_part_common_10), int32(0))
														mBase = m.M
														v760 = m.ExcPending
														if v760 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_11), int32(_a_F_timestamp_part_common_9))
															mBase = m.M
															v765 = m.ExcPending
															if v765 != 0 {
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
												v132 = base.I32_wrap_i64(v129)
												v144 = v132 + int32(_a_F_timestamp_part_common_12)
												v145 = int32(_a_F_timestamp_part_common_13)
												v146 = base.I32_div_u_s(v144, v145)
												v147 = int32(3)
												v153 = int32(2)
												v158 = base.I32_div_u_s((v146*int32(1073595727)+v144)<<(uint(v153)%32)|v147, v145)
												v161 = v132 + int32(_a_F_timestamp_part_common_14) + v146*v147 + v158 + int32(_a_F_timestamp_part_common_15)
												v162 = int32(1461)
												v163 = base.I32_div_u_s(v161, v162)
												v166 = v163*int32(-1461) + v161
												v168 = v166 << (uint(v153) % 32)
												if base.Ui32(v162) <= base.Ui32(v168) {
													v174 = base.I32_rem_u_s(v166+int32(305), int32(365))
													v179 = v174
												} else {
													v178 = base.I32_rem_u_s(v166+int32(306), int32(366))
													v179 = v178
												}
												v181 = base.I32_div_u_s(v168, int32(1461))
												*(*int32)(unsafe.Add(mBase, uint32(v15+int32(68)))) = v181 + v163<<(uint(int32(2))%32) - int32(_a_F_timestamp_part_common_16)
												v189 = v179 + int32(123)
												v193 = int32(base.Ui32(v189*int32(2141)) >> (uint(int32(16)) % 32))
												*(*int32)(unsafe.Add(mBase, uint32(v15+int32(60)))) = v189 - int32(base.Ui32(v193*int32(_a_F_timestamp_part_common_17))>>(uint(int32(8))%32))
												v203 = base.I32_rem_u_s(v193+int32(10), int32(12))
												*(*int32)(unsafe.Add(mBase, uint32(v15-int32(-64)))) = v203 + int32(1)
												if v126 < int64(0) {
													v211 = v126 + int64(86400000000)
												} else {
													v211 = v126
												}
												v213 = base.I64_div_s(v211, int64(3600000000))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+56)) = uint32(v213)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+88)) = int32(0)
												*(*int64)(unsafe.Add(mBase, uint32(v15)+80)) = int64(4294967295)
												v219 = base.I64_extend32_s(v213)
												v222 = v219*int64(-3600000000) + v211
												v224 = base.I64_div_s(v222, int64(60000000))
												*(*uint32)(unsafe.Add(mBase, uint32(v15)+52)) = uint32(v224)
												v226 = base.I64_extend32_s(v224)
												v229 = v222 + v226*int64(-60000000)
												v231 = base.I64_div_s(v229, int64(1000000))
												v232 = base.I32_wrap_i64(v231)
												*(*int32)(unsafe.Add(mBase, uint32(v15)+48)) = v232
												v236 = v231*int64(4293967296) + v229
												v237 = base.I32_wrap_i64(v236)
												v238 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
												switch v238 - int32(18) {
												case 0:
													if l1 != 0 {
														v264 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v236)+base.I64_extend32_s(v231)*int64(1000000), int32(6))
														mBase = m.M
														v265 = m.ExcPending
														if v265 != 0 {
															return int32(0)
														} else {
															v740 = v264
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v271 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v237), float64(1e+06)), base.F64_convert_i32_s(v232)))
														mBase = m.M
														v272 = m.ExcPending
														if v272 != 0 {
															return int32(0)
														} else {
															v740 = v271
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 1:
													v733 = v226
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 2:
													v733 = v219
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 3:
													v273 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+60)))
													v733 = v273
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 4:
													v283 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v284 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v285 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v287 = F_date2j(m, v283, v284, v285)
													mBase = m.M
													v288 = int32(1)
													v290 = F_date2j(m, v283, v288, int32(4))
													mBase = m.M
													v293 = F_j2day(m, v290-v288)
													mBase = m.M
													if v287 < v290-v293 {
														v296 = int32(1)
														v300 = F_date2j(m, v283-v296, v296, int32(4))
														mBase = m.M
														v303 = F_j2day(m, v300-v296)
														mBase = m.M
														v304 = v300
														v305 = v303
													} else {
														v304 = v290
														v305 = v293
													}
													v307 = v305 - v304 + v287
													if int32(357) <= v307 {
														v310 = int32(1)
														v314 = F_date2j(m, v283+v310, v310, int32(4))
														mBase = m.M
														v317 = F_j2day(m, v314-v310)
														mBase = m.M
														v318 = v314 - v317
														if v287 < v318 {
															v321 = v307
														} else {
															v321 = v287 - v318
														}
														v323 = v321
													} else {
														v323 = v307
													}
													v325 = base.I32_div_s(v323, int32(7))
													v733 = base.I64_extend_i32_s(v325 + int32(1))
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 5:
													v274 = int64(*(*int32)(unsafe.Add(mBase, uint32(v15)+64)))
													v733 = v274
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 6:
													v275 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v276 = int32(1)
													v279 = base.I32_div_s(v275-v276, int32(3))
													v733 = base.I64_extend_i32_s(v279 + v276)
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 7:
													v329 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													if int32(0) < v329 {
														v733 = base.I64_extend_i32_u(v329)
													} else {
														v733 = base.I64_extend_i32_s(v329 - int32(1))
													}
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 8:
													v336 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													if int32(0) <= v336 {
														v340 = base.I32_div_u_s(v336, int32(10))
														v733 = base.I64_extend_i32_u(v340)
													} else {
														v345 = base.I32_div_s(int32(9)-v336, int32(-10))
														v733 = base.I64_extend_i32_s(v345)
													}
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 9:
													v347 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													if int32(0) < v347 {
														v353 = base.I32_div_s(v347+int32(99), int32(100))
														v733 = base.I64_extend_i32_s(v353)
													} else {
														v358 = base.I32_div_s(int32(100)-v347, int32(-100))
														v733 = base.I64_extend_i32_s(v358)
													}
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 10:
													v360 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													if int32(0) < v360 {
														v366 = base.I32_div_s(v360+int32(999), int32(1000))
														v733 = base.I64_extend_i32_s(v366)
													} else {
														v371 = base.I32_div_s(int32(1000)-v360, int32(-1000))
														v733 = base.I64_extend_i32_s(v371)
													}
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 11:
													if l1 != 0 {
														v247 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v236)+base.I64_extend32_s(v231)*int64(1000000), int32(3))
														mBase = m.M
														v248 = m.ExcPending
														if v248 != 0 {
															return int32(0)
														} else {
															v740 = v247
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v250 = float64(1000)
														v256 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v232), v250), base.F64_div(base.F64_convert_i32_s(v237), v250)))
														mBase = m.M
														v257 = m.ExcPending
														if v257 != 0 {
															return int32(0)
														} else {
															v740 = v256
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 12:
													v733 = base.I64_extend32_s(v236) + base.I64_extend32_s(v231)*int64(1000000)
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 13:
													v373 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v374 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v375 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v380 = base.B2i32(int32(2) < v374)
													if int32(2) < v374 {
														v381 = int32(_a_F_timestamp_part_common_16)
													} else {
														v381 = int32(_a_F_timestamp_part_common_18)
													}
													v382 = v381 + v373
													v387 = base.I32_div_s(v382, int32(4))
													v390 = base.I32_div_s(v382, int32(-100))
													v393 = base.I32_div_s(v382, int32(400))
													if int32(2) < v374 {
														v397 = int32(1)
													} else {
														v397 = int32(13)
													}
													v402 = base.I32_div_s((v397+v374)*int32(_a_F_timestamp_part_common_17), int32(256))
													v405 = v375 + v382*int32(365) + v387 + v390 + v393 + v402 - int32(_a_F_timestamp_part_common_19)
													if l1 != 0 {
														v407 = F_int64_to_numeric(m, base.I64_extend_i32_s(v405))
														mBase = m.M
														v408 = m.ExcPending
														if v408 != 0 {
															return int32(0)
														} else {
															v410 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
															v411 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
															v412 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
															v413 = int32(60)
															v423 = F_int64_to_numeric(m, base.I64_extend32_s(v236)+base.I64_extend_i32_s(v410+(v411+v412*v413)*v413)*int64(1000000))
															mBase = m.M
															v424 = m.ExcPending
															if v424 != 0 {
																return int32(0)
															} else {
																v426 = F_int64_to_numeric(m, int64(86400000000))
																mBase = m.M
																v427 = m.ExcPending
																if v427 != 0 {
																	return int32(0)
																} else {
																	v429 = F_numeric_div_opt_error(m, v423, v426, int32(0))
																	mBase = m.M
																	v430 = m.ExcPending
																	if v430 != 0 {
																		return int32(0)
																	} else {
																		v432 = F_numeric_add_opt_error(m, v407, v429, int32(0))
																		mBase = m.M
																		v433 = m.ExcPending
																		if v433 != 0 {
																			return int32(0)
																		} else {
																			v740 = v432
																			m.G0 = v15 + int32(96)
																			return v740
																		}
																	}
																}
															}
														}
													} else {
														v437 = *(*int32)(unsafe.Add(mBase, uint32(v15)+48))
														v438 = *(*int32)(unsafe.Add(mBase, uint32(v15)+52))
														v439 = *(*int32)(unsafe.Add(mBase, uint32(v15)+56))
														v440 = int32(60)
														v452 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v237), float64(1e+06)), base.F64_convert_i32_s(v437+(v438+v439*v440)*v440)), float64(86400)), base.F64_convert_i32_s(v405)))
														mBase = m.M
														v453 = m.ExcPending
														if v453 != 0 {
															return int32(0)
														} else {
															v740 = v452
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 14, 19:
													v501 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v502 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v503 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v508 = base.B2i32(int32(2) < v502)
													if int32(2) < v502 {
														v509 = int32(_a_F_timestamp_part_common_16)
													} else {
														v509 = int32(_a_F_timestamp_part_common_18)
													}
													v510 = v509 + v501
													v515 = base.I32_div_s(v510, int32(4))
													v518 = base.I32_div_s(v510, int32(-100))
													v521 = base.I32_div_s(v510, int32(400))
													if int32(2) < v502 {
														v525 = int32(1)
													} else {
														v525 = int32(13)
													}
													v530 = base.I32_div_s((v525+v502)*int32(_a_F_timestamp_part_common_17), int32(256))
													v536 = int32(7)
													v537 = base.I32_rem_s(v503+v510*int32(365)+v515+v518+v521+v530-int32(_a_F_timestamp_part_common_19)+int32(1), v536)
													if v537 < int32(0) {
														v542 = v537 + v536
													} else {
														v542 = v537
													}
													v543 = base.I64_extend_i32_s(v542)
													if v542 != 0 {
														v545 = v543
													} else {
														v545 = int64(7)
													}
													v546 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
													if v546 == int32(37) {
														v549 = v545
													} else {
														v549 = v543
													}
													v733 = v549
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												case 15:
													v550 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v551 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v552 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v557 = base.B2i32(int32(2) < v551)
													if int32(2) < v551 {
														v558 = int32(_a_F_timestamp_part_common_16)
													} else {
														v558 = int32(_a_F_timestamp_part_common_18)
													}
													v559 = v558 + v550
													v564 = base.I32_div_s(v559, int32(4))
													v567 = base.I32_div_s(v559, int32(-100))
													v570 = base.I32_div_s(v559, int32(400))
													if int32(2) < v551 {
														v574 = int32(1)
													} else {
														v574 = int32(13)
													}
													v579 = base.I32_div_s((v574+v551)*int32(_a_F_timestamp_part_common_17), int32(256))
													v583 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v592 = int32(_a_F_timestamp_part_common_18) + v583
													v597 = base.I32_div_s(v592, int32(4))
													v600 = base.I32_div_s(v592, int32(-100))
													v603 = base.I32_div_s(v592, int32(400))
													v612 = base.I32_div_s(int32(_a_F_timestamp_part_common_20), int32(256))
													v733 = base.I64_extend_i32_s(v552 + v559*int32(365) + v564 + v567 + v570 + v579 - int32(_a_F_timestamp_part_common_19) - (int32(1) + v592*int32(365) + v597 + v600 + v603 + v612 - int32(_a_F_timestamp_part_common_19)) + int32(1))
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												default:
													F_errstart_cold(m, int32(21), int32(0))
													mBase = m.M
													v623 = m.ExcPending
													if v623 != 0 {
														return int32(0)
													} else {
														F_errcode(m, int32(1088))
														mBase = m.M
														v626 = m.ExcPending
														if v626 != 0 {
															return int32(0)
														} else {
															v628 = F_format_type_be(m, int32(1114))
															mBase = m.M
															v629 = m.ExcPending
															if v629 != 0 {
																return int32(0)
															} else {
																*(*int32)(unsafe.Add(mBase, uint32(v15)+20)) = v628
																*(*int32)(unsafe.Add(mBase, uint32(v15)+16)) = v60
																F_errmsg(m, int32(_a_F_timestamp_part_common_21), v15+int32(16))
																mBase = m.M
																v636 = m.ExcPending
																if v636 != 0 {
																	return int32(0)
																} else {
																	F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_22), int32(_a_F_timestamp_part_common_9))
																	mBase = m.M
																	v641 = m.ExcPending
																	if v641 != 0 {
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
												case 18:
													v454 = *(*int32)(unsafe.Add(mBase, uint32(v15)+68))
													v455 = *(*int32)(unsafe.Add(mBase, uint32(v15)+64))
													v456 = *(*int32)(unsafe.Add(mBase, uint32(v15)+60))
													v458 = F_date2j(m, v454, v455, v456)
													mBase = m.M
													v459 = int32(1)
													v461 = F_date2j(m, v454, v459, int32(4))
													mBase = m.M
													v464 = F_j2day(m, v461-v459)
													mBase = m.M
													if v458 < v461-v464 {
														v467 = int32(1)
														v468 = v454 - v467
														v471 = F_date2j(m, v468, v467, int32(4))
														mBase = m.M
														v474 = F_j2day(m, v471-v467)
														mBase = m.M
														v475 = v468
														v476 = v471
														v477 = v474
													} else {
														v475 = v454
														v476 = v461
														v477 = v464
													}
													if int32(357) <= v477+v458-v476 {
														v482 = int32(1)
														v483 = v475 + v482
														v486 = F_date2j(m, v483, v482, int32(4))
														mBase = m.M
														v489 = F_j2day(m, v486-v482)
														mBase = m.M
														if v458 < v486-v489 {
															v492 = v475
														} else {
															v492 = v483
														}
														v495 = v492
													} else {
														v495 = v475
													}
													v733 = base.I64_extend_i32_s(v495) - base.I64_extend_i32_u(base.B2i32(v495 <= int32(0)))
													if l1 != 0 {
														v734 = F_int64_to_numeric(m, v733)
														mBase = m.M
														v735 = m.ExcPending
														if v735 != 0 {
															return int32(0)
														} else {
															v740 = v734
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v737 = F_Float8GetDatum(m, base.F64_convert_i64_s(v733))
														mBase = m.M
														v738 = m.ExcPending
														if v738 != 0 {
															return int32(0)
														} else {
															v740 = v737
															m.G0 = v15 + int32(96)
															return v740
														}
													}
												}
											}
										}
									} else {
										v642 = *(*int32)(unsafe.Add(mBase, uint32(v15)+92))
										if v642 == int32(11) {
											v645 = F_SetEpochTimestamp(m)
											mBase = m.M
											v646 = m.ExcPending
											if v646 != 0 {
												return int32(0)
											} else {
												v648 = v645 + int64(9223372036854775807)
												if l1 != 0 {
													if v23 < v648 {
														v652 = F_int64_div_fast_to_numeric(m, v23-v645, int32(6))
														mBase = m.M
														v653 = m.ExcPending
														if v653 != 0 {
															return int32(0)
														} else {
															v740 = v652
															m.G0 = v15 + int32(96)
															return v740
														}
													} else {
														v656 = F_int64_to_numeric(m, v23)
														mBase = m.M
														v657 = m.ExcPending
														if v657 != 0 {
															return int32(0)
														} else {
															v658 = F_int64_to_numeric(m, v645)
															mBase = m.M
															v659 = m.ExcPending
															if v659 != 0 {
																return int32(0)
															} else {
																v661 = F_numeric_sub_opt_error(m, v656, v658, int32(0))
																mBase = m.M
																v662 = m.ExcPending
																if v662 != 0 {
																	return int32(0)
																} else {
																	v664 = F_int64_to_numeric(m, int64(1000000))
																	mBase = m.M
																	v665 = m.ExcPending
																	if v665 != 0 {
																		return int32(0)
																	} else {
																		v667 = F_numeric_div_opt_error(m, v661, v664, int32(0))
																		mBase = m.M
																		v668 = m.ExcPending
																		if v668 != 0 {
																			return int32(0)
																		} else {
																			v670 = F_DirectFunctionCall2Coll(m, int32(1259), int32(0), v667, int32(6))
																			mBase = m.M
																			v671 = m.ExcPending
																			if v671 != 0 {
																				return int32(0)
																			} else {
																				v672 = F_pg_detoast_datum(m, v670)
																				mBase = m.M
																				v673 = m.ExcPending
																				if v673 != 0 {
																					return int32(0)
																				} else {
																					v740 = v672
																					m.G0 = v15 + int32(96)
																					return v740
																				}
																			}
																		}
																	}
																}
															}
														}
													}
												} else {
													if v23 < v648 {
														v680 = base.F64_convert_i64_s(v23 - v645)
													} else {
														v680 = base.F64_sub(base.F64_convert_i64_s(v23), base.F64_convert_i64_s(v645))
													}
													v683 = F_Float8GetDatum(m, base.F64_div(v680, float64(1e+06)))
													mBase = m.M
													v684 = m.ExcPending
													if v684 != 0 {
														return int32(0)
													} else {
														v740 = v683
														m.G0 = v15 + int32(96)
														return v740
													}
												}
											}
										} else {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v688 = m.ExcPending
											if v688 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(1088))
												mBase = m.M
												v691 = m.ExcPending
												if v691 != 0 {
													return int32(0)
												} else {
													v693 = F_format_type_be(m, int32(1114))
													mBase = m.M
													v694 = m.ExcPending
													if v694 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v15)+36)) = v693
														*(*int32)(unsafe.Add(mBase, uint32(v15)+32)) = v60
														F_errmsg(m, int32(_a_F_timestamp_part_common_21), v15+int32(32))
														mBase = m.M
														v701 = m.ExcPending
														if v701 != 0 {
															return int32(0)
														} else {
															F_errfinish(m, int32(_a_F_timestamp_part_common_7), int32(_a_F_timestamp_part_common_23), int32(_a_F_timestamp_part_common_9))
															mBase = m.M
															v706 = m.ExcPending
															if v706 != 0 {
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
									v102 = int32(0)
									v106 = F_DirectFunctionCall3Coll(m, int32(408), v102, int32(_a_F_timestamp_part_common_25), v102, int32(-1))
									mBase = m.M
									v107 = m.ExcPending
									if v107 != 0 {
										return int32(0)
									} else {
										v740 = v106
										m.G0 = v15 + int32(96)
										return v740
									}
								}
							}
						} else {
							v108 = F_Float8GetDatum(m, v84)
							mBase = m.M
							v109 = m.ExcPending
							if v109 != 0 {
								return int32(0)
							} else {
								v740 = v108
								m.G0 = v15 + int32(96)
								return v740
							}
						}
					} else {
						v110 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v110)
						v740 = v79
						m.G0 = v15 + int32(96)
						return v740
					}
				}
			}
		}
	}
}
