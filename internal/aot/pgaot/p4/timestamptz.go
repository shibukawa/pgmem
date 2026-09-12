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
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v18 int32
	_ = v18
	var v19 int64
	_ = v19
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v30 int32
	_ = v30
	var v33 int64
	_ = v33
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v47 int32
	_ = v47
	var v50 int32
	_ = v50
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int64
	_ = v60
	var v69 int64
	_ = v69
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v79 int64
	_ = v79
	var v83 int64
	_ = v83
	var v90 int64
	_ = v90
	var v101 int64
	_ = v101
	var v102 int64
	_ = v102
	var v106 int64
	_ = v106
	var v109 int64
	_ = v109
	var v114 int64
	_ = v114
	var v119 int64
	_ = v119
	var v121 int64
	_ = v121
	var v124 int64
	_ = v124
	var v134 int64
	_ = v134
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v148 int32
	_ = v148
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v188 int32
	_ = v188
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 int32
	_ = v213
	var v216 int32
	_ = v216
	var v220 int32
	_ = v220
	var v225 int32
	_ = v225
	var v229 int32
	_ = v229
	var v232 int32
	_ = v232
	var v236 int32
	_ = v236
	var v241 int32
	_ = v241
	v8 = m.G0
	v10 = v8 - int32(16)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
	if base.Ui64(v13-int64(9223372036854775807)) < base.Ui64(int64(2)) {
		v134 = v13
		v139 = F_Int64GetDatum(m, v134)
		mBase = m.M
		v140 = m.ExcPending
		if v140 != 0 {
			return int32(0)
		} else {
			m.G0 = v10 + int32(16)
			return v139
		}
	} else {
		v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v19 = *(*int64)(unsafe.Add(mBase, uint32(v18)))
		if base.Ui64(v19-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
			F_errstart_cold(m, int32(21), int32(0))
			mBase = m.M
			v148 = m.ExcPending
			if v148 != 0 {
				return int32(0)
			} else {
				F_errcode(m, int32(134217858))
				mBase = m.M
				v151 = m.ExcPending
				if v151 != 0 {
					return int32(0)
				} else {
					F_errmsg(m, int32(398169), int32(0))
					mBase = m.M
					v155 = m.ExcPending
					if v155 != 0 {
						return int32(0)
					} else {
						F_errfinish(m, int32(490442), int32(4858), int32(274662))
						mBase = m.M
						v160 = m.ExcPending
						if v160 != 0 {
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
			v24 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
			v25 = *(*int32)(unsafe.Add(mBase, uint32(v24)+12))
			if v25 != 0 {
				if v25 != int32(2147483647) {
					if v25 != int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(134724), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490442), int32(4868), int32(274662))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
						v30 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
						if v30 != int32(-2147483648) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(134724), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490442), int32(4868), int32(274662))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
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
							v33 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
							if v33 != int64(-9223372036854775807-1) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v47 = m.ExcPending
								if v47 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(1088))
									mBase = m.M
									v50 = m.ExcPending
									if v50 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(134724), int32(0))
										mBase = m.M
										v54 = m.ExcPending
										if v54 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490442), int32(4868), int32(274662))
											mBase = m.M
											v59 = m.ExcPending
											if v59 != 0 {
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
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v229 = m.ExcPending
								if v229 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v232 = m.ExcPending
									if v232 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(151871), int32(0))
										mBase = m.M
										v236 = m.ExcPending
										if v236 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490442), int32(4863), int32(274662))
											mBase = m.M
											v241 = m.ExcPending
											if v241 != 0 {
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
					v36 = *(*int32)(unsafe.Add(mBase, uint32(v24)+8))
					if v36 != int32(2147483647) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v47 = m.ExcPending
						if v47 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(1088))
							mBase = m.M
							v50 = m.ExcPending
							if v50 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(134724), int32(0))
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490442), int32(4868), int32(274662))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
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
						v39 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
						if v39 == int64(9223372036854775807) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v229 = m.ExcPending
							if v229 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v232 = m.ExcPending
								if v232 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(151871), int32(0))
									mBase = m.M
									v236 = m.ExcPending
									if v236 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490442), int32(4863), int32(274662))
										mBase = m.M
										v241 = m.ExcPending
										if v241 != 0 {
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
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v47 = m.ExcPending
							if v47 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(1088))
								mBase = m.M
								v50 = m.ExcPending
								if v50 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(134724), int32(0))
									mBase = m.M
									v54 = m.ExcPending
									if v54 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490442), int32(4868), int32(274662))
										mBase = m.M
										v59 = m.ExcPending
										if v59 != 0 {
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
				v60 = int64(*(*int32)(unsafe.Add(mBase, uint32(v24)+8)))
				v69 = int64(32)
				v70 = int64(20)
				v72 = int64(base.Ui64(v60) >> (uint(v69) % 64))
				v75 = int64(4294967295)
				v76 = int64(500654080)
				v78 = v60 & v75
				v79 = v76 * v78
				v83 = int64(base.Ui64(v79)>>(uint(v69)%64)) + v76*v72
				v90 = v78*v70 + v83&v75
				*(*int64)(unsafe.Add(mBase, uint32(v10)+8)) = v60*int64(0) + v60>>(uint(int64(63))%64)*int64(86400000000) + v70*v72 + int64(base.Ui64(v83)>>(uint(v69)%64)) + int64(base.Ui64(v90)>>(uint(v69)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v10))) = v79&v75 | v90<<(uint(v69)%64)
				v101 = *(*int64)(unsafe.Add(mBase, uint32(v10)+8))
				v102 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
				if v101 != v102>>(uint(int64(63))%64) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(398189), int32(0))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490442), int32(4874), int32(274662))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
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
					v106 = *(*int64)(unsafe.Add(mBase, uint32(v24)))
					v109 = v106 + v102
					if base.B2i32(v106 < int64(0)) != base.B2i32(v109 < v102) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v165 = m.ExcPending
						if v165 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(134217858))
							mBase = m.M
							v168 = m.ExcPending
							if v168 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(398189), int32(0))
								mBase = m.M
								v172 = m.ExcPending
								if v172 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490442), int32(4874), int32(274662))
									mBase = m.M
									v177 = m.ExcPending
									if v177 != 0 {
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
						if v109 <= int64(0) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v181 = m.ExcPending
							if v181 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(134217858))
								mBase = m.M
								v184 = m.ExcPending
								if v184 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(237606), int32(0))
									mBase = m.M
									v188 = m.ExcPending
									if v188 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490442), int32(4879), int32(274662))
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
						} else {
							v114 = v13 - v19
							if base.B2i32(v114 < v13) != base.B2i32(int64(0) < v19) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v200 = m.ExcPending
									if v200 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(398189), int32(0))
										mBase = m.M
										v204 = m.ExcPending
										if v204 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490442), int32(4884), int32(274662))
											mBase = m.M
											v209 = m.ExcPending
											if v209 != 0 {
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
								v119 = base.I64_rem_s(v114, v109)
								v121 = v114 - v119 + v19
								if int64(0) <= v119 {
									v134 = v121
									v139 = F_Int64GetDatum(m, v134)
									mBase = m.M
									v140 = m.ExcPending
									if v140 != 0 {
										return int32(0)
									} else {
										m.G0 = v10 + int32(16)
										return v139
									}
								} else {
									v124 = v121 - v109
									if base.B2i32(v124 < v121)^base.B2i32(int64(0) < v109) != 0 {
										F_errstart_cold(m, int32(21), int32(0))
										mBase = m.M
										v213 = m.ExcPending
										if v213 != 0 {
											return int32(0)
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v216 = m.ExcPending
											if v216 != 0 {
												return int32(0)
											} else {
												F_errmsg(m, int32(398126), int32(0))
												mBase = m.M
												v220 = m.ExcPending
												if v220 != 0 {
													return int32(0)
												} else {
													F_errfinish(m, int32(490442), int32(4902), int32(274662))
													mBase = m.M
													v225 = m.ExcPending
													if v225 != 0 {
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
										if base.Ui64(v124-int64(9223371331200000000)) <= base.Ui64(int64(9011559254509551615)) {
											F_errstart_cold(m, int32(21), int32(0))
											mBase = m.M
											v213 = m.ExcPending
											if v213 != 0 {
												return int32(0)
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v216 = m.ExcPending
												if v216 != 0 {
													return int32(0)
												} else {
													F_errmsg(m, int32(398126), int32(0))
													mBase = m.M
													v220 = m.ExcPending
													if v220 != 0 {
														return int32(0)
													} else {
														F_errfinish(m, int32(490442), int32(4902), int32(274662))
														mBase = m.M
														v225 = m.ExcPending
														if v225 != 0 {
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
											v134 = v124
											v139 = F_Int64GetDatum(m, v134)
											mBase = m.M
											v140 = m.ExcPending
											if v140 != 0 {
												return int32(0)
											} else {
												m.G0 = v10 + int32(16)
												return v139
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
				F_j2date(m, v3+int32(2451545), v10+int32(24), v10+int32(20), v10+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v10)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v10)+12)) = int32(0)
				v37 = *(*int32)(unsafe.Add(mBase, _consts[451]))
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
				F_j2date(m, v2+int32(2451545), v9+int32(24), v9+int32(20), v9+int32(16))
				mBase = m.M
				*(*int64)(unsafe.Add(mBase, uint32(v9)+4)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = int32(0)
				v36 = *(*int32)(unsafe.Add(mBase, _consts[451]))
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
	var v9 int32
	_ = v9
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v25 int32
	_ = v25
	var v32 int32
	_ = v32
	var v43 int32
	_ = v43
	var v46 int32
	_ = v46
	var v49 int32
	_ = v49
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v59 int32
	_ = v59
	var v62 int64
	_ = v62
	var v63 int32
	_ = v63
	var v68 int32
	_ = v68
	var v73 int32
	_ = v73
	var v76 int32
	_ = v76
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v82 int32
	_ = v82
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v94 int32
	_ = v94
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v115 int64
	_ = v115
	var v124 int64
	_ = v124
	var v125 int64
	_ = v125
	var v127 int64
	_ = v127
	var v130 int64
	_ = v130
	var v131 int64
	_ = v131
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v138 int64
	_ = v138
	var v145 int64
	_ = v145
	var v156 int64
	_ = v156
	var v157 int64
	_ = v157
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v173 int64
	_ = v173
	var v176 int64
	_ = v176
	var v180 int32
	_ = v180
	var v185 int64
	_ = v185
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v201 int32
	_ = v201
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	var v217 int32
	_ = v217
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v228 int32
	_ = v228
	var v229 int64
	_ = v229
	var v230 int32
	_ = v230
	var v231 int64
	_ = v231
	var v240 int32
	_ = v240
	var v241 int64
	_ = v241
	var v242 int32
	_ = v242
	var v243 int32
	_ = v243
	var v244 int32
	_ = v244
	v9 = m.G0
	v11 = v9 - int32(496)
	m.G0 = v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v25 = F_ParseDateTime(m, v15, v11+int32(48), int32(153), v11+int32(320), v11+int32(208), v11+int32(428))
	mBase = m.M
	if v25 == int32(0) {
		v32 = *(*int32)(unsafe.Add(mBase, uint32(v11)+428))
		v43 = F_DecodeDateTime(m, v11+int32(320), v11+int32(208), v32, v11+int32(432), v11+int32(440), v11+int32(484), v11+int32(436), v11+int32(40))
		mBase = m.M
		v46 = m.ExcPending
		if v46 != 0 {
			return int32(0)
		} else {
			if v43 == int32(0) {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v11)+432))
				switch v59 - int32(2) {
				case 0:
					v62 = int64(*(*int32)(unsafe.Add(mBase, uint32(v11)+484)))
					v63 = *(*int32)(unsafe.Add(mBase, uint32(v11)+460))
					if v63 <= int32(-4713) {
						if v63 != int32(-4713) {
							v194 = int32(0)
							v195 = F_errsave_start(m, v14)
							mBase = m.M
							v196 = m.ExcPending
							if v196 != 0 {
								return int32(0)
							} else {
								if v195 == int32(0) {
									v244 = v194
									m.G0 = v11 + int32(496)
									return v244
								} else {
									F_errcode(m, int32(134217858))
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
										F_errmsg(m, int32(705599), v11+int32(16))
										mBase = m.M
										v207 = m.ExcPending
										if v207 != 0 {
											return int32(0)
										} else {
											F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
												return int32(0)
											} else {
												v244 = v194
												m.G0 = v11 + int32(496)
												return v244
											}
										}
									}
								}
							}
						} else {
							v68 = *(*int32)(unsafe.Add(mBase, uint32(v11)+456))
							if int32(10) < v68 {
								v79 = v68
								v81 = v11 + int32(24)
								v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+452))
								v87 = base.B2i32(int32(2) < v79)
								if int32(2) < v79 {
									v88 = int32(4800)
								} else {
									v88 = int32(4799)
								}
								v89 = v88 + v63
								v94 = base.I32_div_s(v89, int32(4))
								v97 = base.I32_div_s(v89, int32(-100))
								v100 = base.I32_div_s(v89, int32(400))
								if int32(2) < v79 {
									v104 = int32(1)
								} else {
									v104 = int32(13)
								}
								v109 = base.I32_div_s((v104+v79)*int32(7834), int32(256))
								v115 = base.I64_extend_i32_s(v82 + v89*int32(365) + v94 + v97 + v100 + v109 - int32(32167) - int32(2451545))
								v124 = int64(32)
								v125 = int64(20)
								v127 = int64(base.Ui64(v115) >> (uint(v124) % 64))
								v130 = int64(4294967295)
								v131 = int64(500654080)
								v133 = v115 & v130
								v134 = v131 * v133
								v138 = int64(base.Ui64(v134)>>(uint(v124)%64)) + v131*v127
								v145 = v133*v125 + v138&v130
								*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v115*int64(0) + v115>>(uint(int64(63))%64)*int64(86400000000) + v125*v127 + int64(base.Ui64(v138)>>(uint(v124)%64)) + int64(base.Ui64(v145)>>(uint(v124)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v81))) = v134&v130 | v145<<(uint(v124)%64)
								v156 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
								v157 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
								if v156 != v157>>(uint(int64(63))%64) {
									v194 = int32(0)
									v195 = F_errsave_start(m, v14)
									mBase = m.M
									v196 = m.ExcPending
									if v196 != 0 {
										return int32(0)
									} else {
										if v195 == int32(0) {
											v244 = v194
											m.G0 = v11 + int32(496)
											return v244
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
												F_errmsg(m, int32(705599), v11+int32(16))
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														v244 = v194
														m.G0 = v11 + int32(496)
														return v244
													}
												}
											}
										}
									}
								} else {
									v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+440))
									v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+444))
									v163 = *(*int32)(unsafe.Add(mBase, uint32(v11)+448))
									v164 = int32(60)
									v173 = base.I64_extend_i32_s(v161+(v162+v163*v164)*v164)*int64(1000000) + v62
									v176 = v173 + v157
									if base.B2i32(v173 < int64(0))^base.B2i32(v176 < v157) != 0 {
										v194 = int32(0)
										v195 = F_errsave_start(m, v14)
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											if v195 == int32(0) {
												v244 = v194
												m.G0 = v11 + int32(496)
												return v244
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
													F_errmsg(m, int32(705599), v11+int32(16))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return int32(0)
														} else {
															v244 = v194
															m.G0 = v11 + int32(496)
															return v244
														}
													}
												}
											}
										}
									} else {
										v180 = *(*int32)(unsafe.Add(mBase, uint32(v11)+436))
										v185 = base.I64_extend_i32_s(int32(0)-v180)*int64(-1000000) + v176
										*(*int64)(unsafe.Add(mBase, uint32(v11)+488)) = v185
										if base.Ui64(v185+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
											F_AdjustTimestampForTypmod(m, v11+int32(488), v13, v14)
											mBase = m.M
											v240 = m.ExcPending
											if v240 != 0 {
												return int32(0)
											} else {
												v241 = *(*int64)(unsafe.Add(mBase, uint32(v11)+488))
												v242 = F_Int64GetDatum(m, v241)
												mBase = m.M
												v243 = m.ExcPending
												if v243 != 0 {
													return int32(0)
												} else {
													v244 = v242
													m.G0 = v11 + int32(496)
													return v244
												}
											}
										} else {
											v194 = int32(0)
											v195 = F_errsave_start(m, v14)
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return int32(0)
											} else {
												if v195 == int32(0) {
													v244 = v194
													m.G0 = v11 + int32(496)
													return v244
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
														F_errmsg(m, int32(705599), v11+int32(16))
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return int32(0)
															} else {
																v244 = v194
																m.G0 = v11 + int32(496)
																return v244
															}
														}
													}
												}
											}
										}
									}
								}
							} else {
								v194 = int32(0)
								v195 = F_errsave_start(m, v14)
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return int32(0)
								} else {
									if v195 == int32(0) {
										v244 = v194
										m.G0 = v11 + int32(496)
										return v244
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
											F_errmsg(m, int32(705599), v11+int32(16))
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													v244 = v194
													m.G0 = v11 + int32(496)
													return v244
												}
											}
										}
									}
								}
							}
						}
					} else {
						if v63 <= int32(5874897) {
							v73 = *(*int32)(unsafe.Add(mBase, uint32(v11)+456))
							v79 = v73
							v81 = v11 + int32(24)
							v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+452))
							v87 = base.B2i32(int32(2) < v79)
							if int32(2) < v79 {
								v88 = int32(4800)
							} else {
								v88 = int32(4799)
							}
							v89 = v88 + v63
							v94 = base.I32_div_s(v89, int32(4))
							v97 = base.I32_div_s(v89, int32(-100))
							v100 = base.I32_div_s(v89, int32(400))
							if int32(2) < v79 {
								v104 = int32(1)
							} else {
								v104 = int32(13)
							}
							v109 = base.I32_div_s((v104+v79)*int32(7834), int32(256))
							v115 = base.I64_extend_i32_s(v82 + v89*int32(365) + v94 + v97 + v100 + v109 - int32(32167) - int32(2451545))
							v124 = int64(32)
							v125 = int64(20)
							v127 = int64(base.Ui64(v115) >> (uint(v124) % 64))
							v130 = int64(4294967295)
							v131 = int64(500654080)
							v133 = v115 & v130
							v134 = v131 * v133
							v138 = int64(base.Ui64(v134)>>(uint(v124)%64)) + v131*v127
							v145 = v133*v125 + v138&v130
							*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v115*int64(0) + v115>>(uint(int64(63))%64)*int64(86400000000) + v125*v127 + int64(base.Ui64(v138)>>(uint(v124)%64)) + int64(base.Ui64(v145)>>(uint(v124)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v81))) = v134&v130 | v145<<(uint(v124)%64)
							v156 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
							v157 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
							if v156 != v157>>(uint(int64(63))%64) {
								v194 = int32(0)
								v195 = F_errsave_start(m, v14)
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return int32(0)
								} else {
									if v195 == int32(0) {
										v244 = v194
										m.G0 = v11 + int32(496)
										return v244
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
											F_errmsg(m, int32(705599), v11+int32(16))
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													v244 = v194
													m.G0 = v11 + int32(496)
													return v244
												}
											}
										}
									}
								}
							} else {
								v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+440))
								v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+444))
								v163 = *(*int32)(unsafe.Add(mBase, uint32(v11)+448))
								v164 = int32(60)
								v173 = base.I64_extend_i32_s(v161+(v162+v163*v164)*v164)*int64(1000000) + v62
								v176 = v173 + v157
								if base.B2i32(v173 < int64(0))^base.B2i32(v176 < v157) != 0 {
									v194 = int32(0)
									v195 = F_errsave_start(m, v14)
									mBase = m.M
									v196 = m.ExcPending
									if v196 != 0 {
										return int32(0)
									} else {
										if v195 == int32(0) {
											v244 = v194
											m.G0 = v11 + int32(496)
											return v244
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
												F_errmsg(m, int32(705599), v11+int32(16))
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														v244 = v194
														m.G0 = v11 + int32(496)
														return v244
													}
												}
											}
										}
									}
								} else {
									v180 = *(*int32)(unsafe.Add(mBase, uint32(v11)+436))
									v185 = base.I64_extend_i32_s(int32(0)-v180)*int64(-1000000) + v176
									*(*int64)(unsafe.Add(mBase, uint32(v11)+488)) = v185
									if base.Ui64(v185+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
										F_AdjustTimestampForTypmod(m, v11+int32(488), v13, v14)
										mBase = m.M
										v240 = m.ExcPending
										if v240 != 0 {
											return int32(0)
										} else {
											v241 = *(*int64)(unsafe.Add(mBase, uint32(v11)+488))
											v242 = F_Int64GetDatum(m, v241)
											mBase = m.M
											v243 = m.ExcPending
											if v243 != 0 {
												return int32(0)
											} else {
												v244 = v242
												m.G0 = v11 + int32(496)
												return v244
											}
										}
									} else {
										v194 = int32(0)
										v195 = F_errsave_start(m, v14)
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											if v195 == int32(0) {
												v244 = v194
												m.G0 = v11 + int32(496)
												return v244
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
													F_errmsg(m, int32(705599), v11+int32(16))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return int32(0)
														} else {
															v244 = v194
															m.G0 = v11 + int32(496)
															return v244
														}
													}
												}
											}
										}
									}
								}
							}
						} else {
							if v63 != int32(5874898) {
								v194 = int32(0)
								v195 = F_errsave_start(m, v14)
								mBase = m.M
								v196 = m.ExcPending
								if v196 != 0 {
									return int32(0)
								} else {
									if v195 == int32(0) {
										v244 = v194
										m.G0 = v11 + int32(496)
										return v244
									} else {
										F_errcode(m, int32(134217858))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int32(0)
										} else {
											*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
											F_errmsg(m, int32(705599), v11+int32(16))
											mBase = m.M
											v207 = m.ExcPending
											if v207 != 0 {
												return int32(0)
											} else {
												F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
												mBase = m.M
												v212 = m.ExcPending
												if v212 != 0 {
													return int32(0)
												} else {
													v244 = v194
													m.G0 = v11 + int32(496)
													return v244
												}
											}
										}
									}
								}
							} else {
								v76 = *(*int32)(unsafe.Add(mBase, uint32(v11)+456))
								if int32(5) < v76 {
									v194 = int32(0)
									v195 = F_errsave_start(m, v14)
									mBase = m.M
									v196 = m.ExcPending
									if v196 != 0 {
										return int32(0)
									} else {
										if v195 == int32(0) {
											v244 = v194
											m.G0 = v11 + int32(496)
											return v244
										} else {
											F_errcode(m, int32(134217858))
											mBase = m.M
											v201 = m.ExcPending
											if v201 != 0 {
												return int32(0)
											} else {
												*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
												F_errmsg(m, int32(705599), v11+int32(16))
												mBase = m.M
												v207 = m.ExcPending
												if v207 != 0 {
													return int32(0)
												} else {
													F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
													mBase = m.M
													v212 = m.ExcPending
													if v212 != 0 {
														return int32(0)
													} else {
														v244 = v194
														m.G0 = v11 + int32(496)
														return v244
													}
												}
											}
										}
									}
								} else {
									v79 = v76
									v81 = v11 + int32(24)
									v82 = *(*int32)(unsafe.Add(mBase, uint32(v11)+452))
									v87 = base.B2i32(int32(2) < v79)
									if int32(2) < v79 {
										v88 = int32(4800)
									} else {
										v88 = int32(4799)
									}
									v89 = v88 + v63
									v94 = base.I32_div_s(v89, int32(4))
									v97 = base.I32_div_s(v89, int32(-100))
									v100 = base.I32_div_s(v89, int32(400))
									if int32(2) < v79 {
										v104 = int32(1)
									} else {
										v104 = int32(13)
									}
									v109 = base.I32_div_s((v104+v79)*int32(7834), int32(256))
									v115 = base.I64_extend_i32_s(v82 + v89*int32(365) + v94 + v97 + v100 + v109 - int32(32167) - int32(2451545))
									v124 = int64(32)
									v125 = int64(20)
									v127 = int64(base.Ui64(v115) >> (uint(v124) % 64))
									v130 = int64(4294967295)
									v131 = int64(500654080)
									v133 = v115 & v130
									v134 = v131 * v133
									v138 = int64(base.Ui64(v134)>>(uint(v124)%64)) + v131*v127
									v145 = v133*v125 + v138&v130
									*(*int64)(unsafe.Add(mBase, uint32(v81)+8)) = v115*int64(0) + v115>>(uint(int64(63))%64)*int64(86400000000) + v125*v127 + int64(base.Ui64(v138)>>(uint(v124)%64)) + int64(base.Ui64(v145)>>(uint(v124)%64))
									*(*int64)(unsafe.Add(mBase, uint32(v81))) = v134&v130 | v145<<(uint(v124)%64)
									v156 = *(*int64)(unsafe.Add(mBase, uint32(v11)+32))
									v157 = *(*int64)(unsafe.Add(mBase, uint32(v11)+24))
									if v156 != v157>>(uint(int64(63))%64) {
										v194 = int32(0)
										v195 = F_errsave_start(m, v14)
										mBase = m.M
										v196 = m.ExcPending
										if v196 != 0 {
											return int32(0)
										} else {
											if v195 == int32(0) {
												v244 = v194
												m.G0 = v11 + int32(496)
												return v244
											} else {
												F_errcode(m, int32(134217858))
												mBase = m.M
												v201 = m.ExcPending
												if v201 != 0 {
													return int32(0)
												} else {
													*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
													F_errmsg(m, int32(705599), v11+int32(16))
													mBase = m.M
													v207 = m.ExcPending
													if v207 != 0 {
														return int32(0)
													} else {
														F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
														mBase = m.M
														v212 = m.ExcPending
														if v212 != 0 {
															return int32(0)
														} else {
															v244 = v194
															m.G0 = v11 + int32(496)
															return v244
														}
													}
												}
											}
										}
									} else {
										v161 = *(*int32)(unsafe.Add(mBase, uint32(v11)+440))
										v162 = *(*int32)(unsafe.Add(mBase, uint32(v11)+444))
										v163 = *(*int32)(unsafe.Add(mBase, uint32(v11)+448))
										v164 = int32(60)
										v173 = base.I64_extend_i32_s(v161+(v162+v163*v164)*v164)*int64(1000000) + v62
										v176 = v173 + v157
										if base.B2i32(v173 < int64(0))^base.B2i32(v176 < v157) != 0 {
											v194 = int32(0)
											v195 = F_errsave_start(m, v14)
											mBase = m.M
											v196 = m.ExcPending
											if v196 != 0 {
												return int32(0)
											} else {
												if v195 == int32(0) {
													v244 = v194
													m.G0 = v11 + int32(496)
													return v244
												} else {
													F_errcode(m, int32(134217858))
													mBase = m.M
													v201 = m.ExcPending
													if v201 != 0 {
														return int32(0)
													} else {
														*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
														F_errmsg(m, int32(705599), v11+int32(16))
														mBase = m.M
														v207 = m.ExcPending
														if v207 != 0 {
															return int32(0)
														} else {
															F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
															mBase = m.M
															v212 = m.ExcPending
															if v212 != 0 {
																return int32(0)
															} else {
																v244 = v194
																m.G0 = v11 + int32(496)
																return v244
															}
														}
													}
												}
											}
										} else {
											v180 = *(*int32)(unsafe.Add(mBase, uint32(v11)+436))
											v185 = base.I64_extend_i32_s(int32(0)-v180)*int64(-1000000) + v176
											*(*int64)(unsafe.Add(mBase, uint32(v11)+488)) = v185
											if base.Ui64(v185+int64(211813488000000000)) < base.Ui64(int64(-9011559254509551616)) {
												F_AdjustTimestampForTypmod(m, v11+int32(488), v13, v14)
												mBase = m.M
												v240 = m.ExcPending
												if v240 != 0 {
													return int32(0)
												} else {
													v241 = *(*int64)(unsafe.Add(mBase, uint32(v11)+488))
													v242 = F_Int64GetDatum(m, v241)
													mBase = m.M
													v243 = m.ExcPending
													if v243 != 0 {
														return int32(0)
													} else {
														v244 = v242
														m.G0 = v11 + int32(496)
														return v244
													}
												}
											} else {
												v194 = int32(0)
												v195 = F_errsave_start(m, v14)
												mBase = m.M
												v196 = m.ExcPending
												if v196 != 0 {
													return int32(0)
												} else {
													if v195 == int32(0) {
														v244 = v194
														m.G0 = v11 + int32(496)
														return v244
													} else {
														F_errcode(m, int32(134217858))
														mBase = m.M
														v201 = m.ExcPending
														if v201 != 0 {
															return int32(0)
														} else {
															*(*int32)(unsafe.Add(mBase, uint32(v11)+16)) = v15
															F_errmsg(m, int32(705599), v11+int32(16))
															mBase = m.M
															v207 = m.ExcPending
															if v207 != 0 {
																return int32(0)
															} else {
																F_errsave_finish(m, v14, int32(490442), int32(457), int32(276035))
																mBase = m.M
																v212 = m.ExcPending
																if v212 != 0 {
																	return int32(0)
																} else {
																	v244 = v194
																	m.G0 = v11 + int32(496)
																	return v244
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
					v217 = m.ExcPending
					if v217 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v11)+4)) = v15
						v219 = *(*int32)(unsafe.Add(mBase, uint32(v11)+432))
						*(*int32)(unsafe.Add(mBase, uint32(v11))) = v219
						F_errmsg_internal(m, int32(670490), v11)
						mBase = m.M
						v223 = m.ExcPending
						if v223 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(490442), int32(474), int32(276035))
							mBase = m.M
							v228 = m.ExcPending
							if v228 != 0 {
								return int32(0)
							} else {
								base.Wasm_trap_unreachable()
								for {
								}
							}
						}
					}
				case 7:
					v231 = int64(-9223372036854775807 - 1)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+488)) = v231
					F_AdjustTimestampForTypmod(m, v11+int32(488), v13, v14)
					mBase = m.M
					v240 = m.ExcPending
					if v240 != 0 {
						return int32(0)
					} else {
						v241 = *(*int64)(unsafe.Add(mBase, uint32(v11)+488))
						v242 = F_Int64GetDatum(m, v241)
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return int32(0)
						} else {
							v244 = v242
							m.G0 = v11 + int32(496)
							return v244
						}
					}
				case 8:
					v231 = int64(9223372036854775807)
					*(*int64)(unsafe.Add(mBase, uint32(v11)+488)) = v231
					F_AdjustTimestampForTypmod(m, v11+int32(488), v13, v14)
					mBase = m.M
					v240 = m.ExcPending
					if v240 != 0 {
						return int32(0)
					} else {
						v241 = *(*int64)(unsafe.Add(mBase, uint32(v11)+488))
						v242 = F_Int64GetDatum(m, v241)
						mBase = m.M
						v243 = m.ExcPending
						if v243 != 0 {
							return int32(0)
						} else {
							v244 = v242
							m.G0 = v11 + int32(496)
							return v244
						}
					}
				case 9:
					v229 = F_SetEpochTimestamp(m)
					mBase = m.M
					v230 = m.ExcPending
					if v230 != 0 {
						return int32(0)
					} else {
						v231 = v229
						*(*int64)(unsafe.Add(mBase, uint32(v11)+488)) = v231
						F_AdjustTimestampForTypmod(m, v11+int32(488), v13, v14)
						mBase = m.M
						v240 = m.ExcPending
						if v240 != 0 {
							return int32(0)
						} else {
							v241 = *(*int64)(unsafe.Add(mBase, uint32(v11)+488))
							v242 = F_Int64GetDatum(m, v241)
							mBase = m.M
							v243 = m.ExcPending
							if v243 != 0 {
								return int32(0)
							} else {
								v244 = v242
								m.G0 = v11 + int32(496)
								return v244
							}
						}
					}
				}
			} else {
				v49 = v43
				F_DateTimeParseError(m, v49, v11+int32(40), v15, int32(368482), v14)
				mBase = m.M
				v54 = m.ExcPending
				if v54 != 0 {
					return int32(0)
				} else {
					v55 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
					v244 = int32(0)
					m.G0 = v11 + int32(496)
					return v244
				}
			}
		}
	} else {
		v49 = v25
		F_DateTimeParseError(m, v49, v11+int32(40), v15, int32(368482), v14)
		mBase = m.M
		v54 = m.ExcPending
		if v54 != 0 {
			return int32(0)
		} else {
			v55 = int32(1)
			*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v55)
			v244 = int32(0)
			m.G0 = v11 + int32(496)
			return v244
		}
	}
}
func F_timestamptz_izone(m *base.Module, l0 int32) int32 {
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
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	var v27 int32
	_ = v27
	var v30 int64
	_ = v30
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v56 int32
	_ = v56
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v59 int64
	_ = v59
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v71 int64
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v100 int32
	_ = v100
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v108 int32
	_ = v108
	var v113 int32
	_ = v113
	v5 = m.G0
	v7 = v5 - int32(32)
	m.G0 = v7
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int64)(unsafe.Add(mBase, uint32(v9)))
	if base.Ui64(int64(2)) <= base.Ui64(v10-int64(9223372036854775807)) {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v15)+12))
		if v16 != 0 {
			if v16 != int32(2147483647) {
				if v16 != int32(-2147483648) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v103 = F_DirectFunctionCall1Coll(m, int32(1289), int32(0), v15)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
								F_errmsg(m, int32(112368), v7)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490442), int32(6650), int32(368120))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
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
					v21 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
					if v21 != int32(-2147483648) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v103 = F_DirectFunctionCall1Coll(m, int32(1289), int32(0), v15)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
									F_errmsg(m, int32(112368), v7)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490442), int32(6650), int32(368120))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
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
						v24 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
						if v24 == int64(-9223372036854775807-1) {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v38 = m.ExcPending
							if v38 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v41 = m.ExcPending
								if v41 != 0 {
									return int32(0)
								} else {
									v44 = F_DirectFunctionCall1Coll(m, int32(1289), int32(0), v15)
									mBase = m.M
									v45 = m.ExcPending
									if v45 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v44
										F_errmsg(m, int32(346636), v7+int32(16))
										mBase = m.M
										v51 = m.ExcPending
										if v51 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490442), int32(6643), int32(368120))
											mBase = m.M
											v56 = m.ExcPending
											if v56 != 0 {
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
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v100 = m.ExcPending
								if v100 != 0 {
									return int32(0)
								} else {
									v103 = F_DirectFunctionCall1Coll(m, int32(1289), int32(0), v15)
									mBase = m.M
									v104 = m.ExcPending
									if v104 != 0 {
										return int32(0)
									} else {
										*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
										F_errmsg(m, int32(112368), v7)
										mBase = m.M
										v108 = m.ExcPending
										if v108 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(490442), int32(6650), int32(368120))
											mBase = m.M
											v113 = m.ExcPending
											if v113 != 0 {
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
			} else {
				v27 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
				if v27 != int32(2147483647) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v100 = m.ExcPending
						if v100 != 0 {
							return int32(0)
						} else {
							v103 = F_DirectFunctionCall1Coll(m, int32(1289), int32(0), v15)
							mBase = m.M
							v104 = m.ExcPending
							if v104 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
								F_errmsg(m, int32(112368), v7)
								mBase = m.M
								v108 = m.ExcPending
								if v108 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(490442), int32(6650), int32(368120))
									mBase = m.M
									v113 = m.ExcPending
									if v113 != 0 {
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
					v30 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
					if v30 != int64(9223372036854775807) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v97 = m.ExcPending
						if v97 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v100 = m.ExcPending
							if v100 != 0 {
								return int32(0)
							} else {
								v103 = F_DirectFunctionCall1Coll(m, int32(1289), int32(0), v15)
								mBase = m.M
								v104 = m.ExcPending
								if v104 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
									F_errmsg(m, int32(112368), v7)
									mBase = m.M
									v108 = m.ExcPending
									if v108 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490442), int32(6650), int32(368120))
										mBase = m.M
										v113 = m.ExcPending
										if v113 != 0 {
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
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v38 = m.ExcPending
						if v38 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v41 = m.ExcPending
							if v41 != 0 {
								return int32(0)
							} else {
								v44 = F_DirectFunctionCall1Coll(m, int32(1289), int32(0), v15)
								mBase = m.M
								v45 = m.ExcPending
								if v45 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v7)+16)) = v44
									F_errmsg(m, int32(346636), v7+int32(16))
									mBase = m.M
									v51 = m.ExcPending
									if v51 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(490442), int32(6643), int32(368120))
										mBase = m.M
										v56 = m.ExcPending
										if v56 != 0 {
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
		} else {
			v57 = *(*int32)(unsafe.Add(mBase, uint32(v15)+8))
			if v57 != 0 {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v100 = m.ExcPending
					if v100 != 0 {
						return int32(0)
					} else {
						v103 = F_DirectFunctionCall1Coll(m, int32(1289), int32(0), v15)
						mBase = m.M
						v104 = m.ExcPending
						if v104 != 0 {
							return int32(0)
						} else {
							*(*int32)(unsafe.Add(mBase, uint32(v7))) = v103
							F_errmsg(m, int32(112368), v7)
							mBase = m.M
							v108 = m.ExcPending
							if v108 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490442), int32(6650), int32(368120))
								mBase = m.M
								v113 = m.ExcPending
								if v113 != 0 {
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
				v58 = *(*int64)(unsafe.Add(mBase, uint32(v15)))
				v59 = int64(-1000000)
				v60 = base.I64_div_s(v58, v59)
				v64 = base.I64_extend32_s(v60)*v59 + v10
				if base.Ui64(int64(-9011559254509551616)) <= base.Ui64(v64+int64(211813488000000000)) {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(134217858))
						mBase = m.M
						v84 = m.ExcPending
						if v84 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(398126), int32(0))
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(490442), int32(6659), int32(368120))
								mBase = m.M
								v93 = m.ExcPending
								if v93 != 0 {
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
					v71 = v64
					v72 = F_Int64GetDatum(m, v71)
					mBase = m.M
					v73 = m.ExcPending
					if v73 != 0 {
						return int32(0)
					} else {
						m.G0 = v7 + int32(32)
						return v72
					}
				}
			}
		}
	} else {
		v71 = v10
		v72 = F_Int64GetDatum(m, v71)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			m.G0 = v7 + int32(32)
			return v72
		}
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
	var v10 int32
	_ = v10
	var v11 int64
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
	v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v10)))
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
				F_interval_um_internal(m, v9, v7)
				mBase = m.M
				v23 = m.ExcPending
				if v23 != 0 {
					return int32(0)
				} else {
					v24 = F_timestamptz_pl_interval_internal(m, v11, v7, v20)
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
	var v33 int32
	_ = v33
	var v35 int32
	_ = v35
	var v44 int32
	_ = v44
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
	var v70 int32
	_ = v70
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v81 int32
	_ = v81
	var v83 int32
	_ = v83
	var v90 int32
	_ = v90
	var v91 int32
	_ = v91
	var v92 int32
	_ = v92
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v107 int32
	_ = v107
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v131 int32
	_ = v131
	var v135 int32
	_ = v135
	var v142 int32
	_ = v142
	var v144 int32
	_ = v144
	var v147 int32
	_ = v147
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v171 int32
	_ = v171
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v179 int32
	_ = v179
	var v186 int32
	_ = v186
	var v193 int32
	_ = v193
	var v194 int32
	_ = v194
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v209 int32
	_ = v209
	var v213 float64
	_ = v213
	var v214 int32
	_ = v214
	var v220 int32
	_ = v220
	var v224 int32
	_ = v224
	var v225 int32
	_ = v225
	var v231 int32
	_ = v231
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v237 int32
	_ = v237
	var v238 int32
	_ = v238
	var v239 int32
	_ = v239
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v252 int32
	_ = v252
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v265 int32
	_ = v265
	var v268 int32
	_ = v268
	var v270 int64
	_ = v270
	var v271 int64
	_ = v271
	var v275 int32
	_ = v275
	var v276 int64
	_ = v276
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v285 float64
	_ = v285
	var v287 int32
	_ = v287
	var v292 int32
	_ = v292
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int64
	_ = v295
	var v301 int32
	_ = v301
	var v302 int32
	_ = v302
	var v303 int32
	_ = v303
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v311 int64
	_ = v311
	var v312 int64
	_ = v312
	var v313 int64
	_ = v313
	var v314 int64
	_ = v314
	var v315 int32
	_ = v315
	var v316 int32
	_ = v316
	var v319 int32
	_ = v319
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v327 int32
	_ = v327
	var v328 int32
	_ = v328
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v336 int32
	_ = v336
	var v340 int32
	_ = v340
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v345 int32
	_ = v345
	var v347 int32
	_ = v347
	var v350 int32
	_ = v350
	var v354 int32
	_ = v354
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v369 int32
	_ = v369
	var v376 int32
	_ = v376
	var v380 int32
	_ = v380
	var v385 int32
	_ = v385
	var v387 int32
	_ = v387
	var v393 int32
	_ = v393
	var v398 int32
	_ = v398
	var v400 int32
	_ = v400
	var v406 int32
	_ = v406
	var v411 int32
	_ = v411
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
	var v445 int32
	_ = v445
	var v447 int32
	_ = v447
	var v448 int32
	_ = v448
	var v449 int64
	_ = v449
	var v450 int32
	_ = v450
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v453 int32
	_ = v453
	var v463 int32
	_ = v463
	var v464 int32
	_ = v464
	var v466 int32
	_ = v466
	var v467 int32
	_ = v467
	var v469 int32
	_ = v469
	var v470 int32
	_ = v470
	var v472 int32
	_ = v472
	var v473 int32
	_ = v473
	var v474 int32
	_ = v474
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v481 int32
	_ = v481
	var v493 int32
	_ = v493
	var v494 int32
	_ = v494
	var v495 int32
	_ = v495
	var v496 int32
	_ = v496
	var v497 int32
	_ = v497
	var v499 int32
	_ = v499
	var v500 int32
	_ = v500
	var v502 int32
	_ = v502
	var v505 int32
	_ = v505
	var v508 int32
	_ = v508
	var v509 int32
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v516 int32
	_ = v516
	var v517 int32
	_ = v517
	var v518 int32
	_ = v518
	var v523 int32
	_ = v523
	var v524 int32
	_ = v524
	var v527 int32
	_ = v527
	var v530 int32
	_ = v530
	var v533 int32
	_ = v533
	var v536 int32
	_ = v536
	var v542 int32
	_ = v542
	var v543 int32
	_ = v543
	var v544 int32
	_ = v544
	var v549 int32
	_ = v549
	var v550 int32
	_ = v550
	var v551 int32
	_ = v551
	var v556 int32
	_ = v556
	var v559 int32
	_ = v559
	var v562 int32
	_ = v562
	var v566 int32
	_ = v566
	var v571 int32
	_ = v571
	var v577 int32
	_ = v577
	var v578 int32
	_ = v578
	var v583 int32
	_ = v583
	var v584 int64
	_ = v584
	var v586 int64
	_ = v586
	var v587 int32
	_ = v587
	var v590 int64
	_ = v590
	var v591 int32
	_ = v591
	var v592 int32
	_ = v592
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v599 int32
	_ = v599
	var v600 int32
	_ = v600
	var v605 int32
	_ = v605
	var v608 int32
	_ = v608
	var v611 int32
	_ = v611
	var v615 int32
	_ = v615
	var v620 int32
	_ = v620
	var v624 int32
	_ = v624
	var v633 int32
	_ = v633
	var v638 int32
	_ = v638
	var v641 int32
	_ = v641
	var v644 int32
	_ = v644
	var v653 int32
	_ = v653
	var v664 int32
	_ = v664
	var v667 int32
	_ = v667
	var v669 int32
	_ = v669
	var v670 int32
	_ = v670
	var v677 int32
	_ = v677
	var v682 int32
	_ = v682
	var v683 int32
	_ = v683
	var v686 int64
	_ = v686
	var v687 int32
	_ = v687
	var v689 int64
	_ = v689
	var v693 int32
	_ = v693
	var v694 int32
	_ = v694
	var v697 int32
	_ = v697
	var v698 int32
	_ = v698
	var v699 int32
	_ = v699
	var v700 int32
	_ = v700
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v706 int32
	_ = v706
	var v708 int32
	_ = v708
	var v709 int32
	_ = v709
	var v711 int32
	_ = v711
	var v712 int32
	_ = v712
	var v713 int32
	_ = v713
	var v714 int32
	_ = v714
	var v721 float64
	_ = v721
	var v724 int32
	_ = v724
	var v725 int32
	_ = v725
	var v729 int32
	_ = v729
	var v732 int32
	_ = v732
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v742 int32
	_ = v742
	var v747 int32
	_ = v747
	var v751 int32
	_ = v751
	var v754 int32
	_ = v754
	var v756 int32
	_ = v756
	var v757 int32
	_ = v757
	var v762 int32
	_ = v762
	var v767 int32
	_ = v767
	var v769 int32
	_ = v769
	var v774 int64
	_ = v774
	var v775 int32
	_ = v775
	var v776 int32
	_ = v776
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v785 int32
	_ = v785
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v800 int32
	_ = v800
	var v805 int32
	_ = v805
	v12 = m.G0
	v14 = v12 - int32(96)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v17 = F_pg_detoast_datum_packed(m, v16)
	mBase = m.M
	v20 = m.ExcPending
	if v20 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int32(0)
L2:
	;
	v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
	v23 = int32(1)
	v24 = v17 + v23
	v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v17))))
	v29 = v27 & v23
	if v29 != 0 {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	v30 = v24
	goto L5
L4:
	;
	v30 = v17 + int32(4)
	goto L5
L5:
	;
	if v27 == int32(1) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	v60 = F_downcase_truncate_identifier(m, v30, v58, int32(0))
	mBase = m.M
	v61 = m.ExcPending
	if v61 != 0 {
		goto L1
	} else {
		goto L17
	}
L7:
	;
	v33 = int32(4)
	v35 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
	if v35&int32(254) == int32(2) {
		goto L10
	} else {
		goto L11
	}
L8:
	;
	goto L9
L9:
	;
	v48 = int32(1)
	if v29 != 0 {
		v58 = int32(base.Ui32(v27)>>(uint(v48)%32)) - v48
		goto L6
	} else {
		goto L16
	}
L10:
	;
	v44 = v33
	goto L12
L11:
	;
	v44 = base.B2i32(v35 == int32(18)) << (uint(v33) % 32)
	goto L12
L12:
	;
	if v35 == int32(1) {
		goto L13
	} else {
		goto L14
	}
L13:
	;
	v47 = v33
	goto L15
L14:
	;
	v47 = v44
	goto L15
L15:
	;
	v58 = v47
	goto L6
L16:
	;
	v52 = *(*int32)(unsafe.Add(mBase, uint32(v17)))
	v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
	goto L6
L17:
	;
	v63 = v14 + int32(88)
	v70 = *(*int32)(unsafe.Add(mBase, _consts[1247]))
	if v70 != 0 {
		goto L20
	} else {
		goto L21
	}
L18:
	;
	if v131 == int32(31) {
		goto L37
	} else {
		goto L38
	}
L19:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1247])) = v114
	v121 = int32(*(*int8)(unsafe.Add(mBase, uint32(v114)+11)))
	v122 = *(*int32)(unsafe.Add(mBase, uint32(v114)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = v122
	v131 = v121
	goto L18
L20:
	;
	v72 = F_strncmp(m, v60, v70, int32(10))
	mBase = m.M
	if v72 == int32(0) {
		v114 = v70
		goto L19
	} else {
		goto L23
	}
L21:
	;
	goto L22
L22:
	;
	v75 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60))))
	v81 = int32(1639072)
	v83 = int32(1640032)
	goto L24
L23:
	;
	goto L22
L24:
	;
	v90 = v81 + (v83-v81)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v91 = int32(*(*int8)(unsafe.Add(mBase, uint32(v90))))
	v92 = v75 - v91
	if v92 == int32(0) {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v63))) = int32(0)
	v131 = int32(31)
	goto L18
L26:
	;
	v96 = F_strncmp(m, v60, v90, int32(10))
	mBase = m.M
	if v96 == int32(0) {
		v114 = v90
		goto L19
	} else {
		goto L29
	}
L27:
	;
	v99 = v92
	goto L28
L28:
	;
	v103 = base.B2i32(v99 < int32(0))
	if v99 < int32(0) {
		goto L30
	} else {
		goto L31
	}
L29:
	;
	v99 = v96
	goto L28
L30:
	;
	v104 = v90 - int32(16)
	goto L32
L31:
	;
	v104 = v83
	goto L32
L32:
	;
	if v99 < int32(0) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v107 = v81
	goto L35
L34:
	;
	v107 = v90 + int32(16)
	goto L35
L35:
	;
	if base.Ui32(v107) <= base.Ui32(v104) {
		v81 = v107
		v83 = v104
		goto L24
	} else {
		goto L36
	}
L36:
	;
	goto L25
L37:
	;
	v135 = v14 + int32(88)
	v142 = *(*int32)(unsafe.Add(mBase, _consts[1248]))
	if v142 != 0 {
		goto L42
	} else {
		goto L43
	}
L38:
	;
	v204 = v131
	goto L39
L39:
	;
	if base.Ui64(int64(1)) < base.Ui64(v22-int64(9223372036854775807)) {
		goto L61
	} else {
		goto L62
	}
L40:
	;
	v204 = v203
	goto L39
L41:
	;
	*(*int32)(unsafe.Add(mBase, _consts[1248])) = v186
	v193 = int32(*(*int8)(unsafe.Add(mBase, uint32(v186)+11)))
	v194 = *(*int32)(unsafe.Add(mBase, uint32(v186)+12))
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = v194
	v203 = v193
	goto L40
L42:
	;
	v144 = F_strncmp(m, v60, v142, int32(10))
	mBase = m.M
	if v144 == int32(0) {
		v186 = v142
		goto L41
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v147 = int32(*(*int8)(unsafe.Add(mBase, uint32(v60))))
	v153 = int32(1637920)
	v155 = int32(1639056)
	goto L46
L45:
	;
	goto L44
L46:
	;
	v162 = v153 + (v155-v153)>>(uint(int32(5))%32)<<(uint(int32(4))%32)
	v163 = int32(*(*int8)(unsafe.Add(mBase, uint32(v162))))
	v164 = v147 - v163
	if v164 == int32(0) {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v135))) = int32(0)
	v203 = int32(31)
	goto L40
L48:
	;
	v168 = F_strncmp(m, v60, v162, int32(10))
	mBase = m.M
	if v168 == int32(0) {
		v186 = v162
		goto L41
	} else {
		goto L51
	}
L49:
	;
	v171 = v164
	goto L50
L50:
	;
	v175 = base.B2i32(v171 < int32(0))
	if v171 < int32(0) {
		goto L52
	} else {
		goto L53
	}
L51:
	;
	v171 = v168
	goto L50
L52:
	;
	v176 = v162 - int32(16)
	goto L54
L53:
	;
	v176 = v155
	goto L54
L54:
	;
	if v171 < int32(0) {
		goto L55
	} else {
		goto L56
	}
L55:
	;
	v179 = v153
	goto L57
L56:
	;
	v179 = v162 + int32(16)
	goto L57
L57:
	;
	if base.Ui32(v179) <= base.Ui32(v176) {
		v153 = v179
		v155 = v176
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
	v793 = m.ExcPending
	if v793 != 0 {
		goto L1
	} else {
		goto L235
	}
L60:
	;
	m.G0 = v14 + int32(96)
	return v785
L61:
	;
	switch v204 {
	case 0:
		goto L80
	default:
		goto L79
	case 17:
		goto L81
	}
L62:
	;
	v209 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	v213 = F_NonFiniteTimestampTzPart(m, v204, v209, v60, base.B2i32(v22 == int64(-9223372036854775807-1)), int32(1))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L1
	} else {
		goto L63
	}
L63:
	;
	if base.F64_ne(v213, float64(0)) != 0 {
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
	v239 = int32(1)
	*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v239)
	v785 = int32(0)
	goto L60
L67:
	;
	if base.F64_lt(v213, float64(0)) != 0 {
		goto L70
	} else {
		goto L71
	}
L68:
	;
	goto L69
L69:
	;
	v237 = F_Float8GetDatum(m, v213)
	mBase = m.M
	v238 = m.ExcPending
	if v238 != 0 {
		goto L1
	} else {
		goto L76
	}
L70:
	;
	v220 = int32(0)
	v224 = F_DirectFunctionCall3Coll(m, int32(408), v220, int32(11387), v220, int32(-1))
	mBase = m.M
	v225 = m.ExcPending
	if v225 != 0 {
		goto L1
	} else {
		goto L73
	}
L71:
	;
	goto L72
L72:
	;
	if base.F64_gt(v213, float64(0)) == int32(0) {
		goto L61
	} else {
		goto L74
	}
L73:
	;
	v785 = v224
	goto L60
L74:
	;
	v231 = int32(0)
	v235 = F_DirectFunctionCall3Coll(m, int32(408), v231, int32(11398), v231, int32(-1))
	mBase = m.M
	v236 = m.ExcPending
	if v236 != 0 {
		goto L1
	} else {
		goto L75
	}
L75:
	;
	v785 = v235
	goto L60
L76:
	;
	v785 = v237
	goto L60
L77:
	;
	if l1 != 0 {
		goto L230
	} else {
		goto L231
	}
L78:
	;
	v769 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	v774 = base.I64_extend_i32_s(int32(0) - v769)
	goto L77
L79:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v751 = m.ExcPending
	if v751 != 0 {
		goto L1
	} else {
		goto L225
	}
L80:
	;
	v683 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v683 == int32(11) {
		goto L198
	} else {
		goto L199
	}
L81:
	;
	v249 = int32(0)
	v251 = F_timestamp2tm(m, v22, v14+int32(92), v14+int32(40), v14+int32(84), v249, v249)
	mBase = m.M
	v252 = m.ExcPending
	if v252 != 0 {
		goto L1
	} else {
		goto L82
	}
L82:
	;
	if v251 != 0 {
		goto L59
	} else {
		goto L83
	}
L83:
	;
	v253 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	switch v253 - int32(4) {
	case 0:
		goto L78
	default:
		goto L84
	case 14:
		goto L99
	case 15:
		goto L98
	case 16:
		goto L97
	case 17:
		goto L96
	case 18:
		goto L93
	case 19:
		goto L95
	case 20:
		goto L94
	case 21:
		goto L92
	case 22:
		goto L91
	case 23:
		goto L90
	case 24:
		goto L89
	case 25:
		goto L100
	case 26:
		goto L101
	case 27:
		goto L88
	case 28, 33:
		goto L86
	case 29:
		goto L85
	case 30:
		goto L102
	case 31:
		goto L103
	case 32:
		goto L87
	}
L84:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v664 = m.ExcPending
	if v664 != 0 {
		goto L1
	} else {
		goto L193
	}
L85:
	;
	v591 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v592 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v593 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v598 = base.B2i32(int32(2) < v592)
	if int32(2) < v592 {
		goto L180
	} else {
		goto L181
	}
L86:
	;
	v542 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v543 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v544 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v549 = base.B2i32(int32(2) < v543)
	if int32(2) < v543 {
		goto L163
	} else {
		goto L164
	}
L87:
	;
	v495 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v496 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v497 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v499 = F_date2j(m, v495, v496, v497)
	mBase = m.M
	v500 = int32(1)
	v502 = F_date2j(m, v495, v500, int32(4))
	mBase = m.M
	v505 = F_j2day(m, v502-v500)
	mBase = m.M
	if v499 < v502-v505 {
		goto L153
	} else {
		goto L154
	}
L88:
	;
	v413 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v414 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v415 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v420 = base.B2i32(int32(2) < v414)
	if int32(2) < v414 {
		goto L137
	} else {
		goto L138
	}
L89:
	;
	v400 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if int32(0) < v400 {
		goto L133
	} else {
		goto L134
	}
L90:
	;
	v387 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if int32(0) < v387 {
		goto L130
	} else {
		goto L131
	}
L91:
	;
	v376 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if int32(0) < v376 {
		goto L127
	} else {
		goto L128
	}
L92:
	;
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	if int32(0) < v369 {
		goto L124
	} else {
		goto L125
	}
L93:
	;
	v323 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	v324 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v325 = *(*int32)(unsafe.Add(mBase, uint32(v14)+52))
	v327 = F_date2j(m, v323, v324, v325)
	mBase = m.M
	v328 = int32(1)
	v330 = F_date2j(m, v323, v328, int32(4))
	mBase = m.M
	v333 = F_j2day(m, v330-v328)
	mBase = m.M
	if v327 < v330-v333 {
		goto L115
	} else {
		goto L116
	}
L94:
	;
	v315 = *(*int32)(unsafe.Add(mBase, uint32(v14)+56))
	v316 = int32(1)
	v319 = base.I32_div_s(v315-v316, int32(3))
	v774 = base.I64_extend_i32_s(v319 + v316)
	goto L77
L95:
	;
	v314 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+56)))
	v774 = v314
	goto L77
L96:
	;
	v313 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+52)))
	v774 = v313
	goto L77
L97:
	;
	v312 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+48)))
	v774 = v312
	goto L77
L98:
	;
	v311 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+44)))
	v774 = v311
	goto L77
L99:
	;
	v294 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if l1 != 0 {
		goto L109
	} else {
		goto L110
	}
L100:
	;
	v275 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	if l1 != 0 {
		goto L104
	} else {
		goto L105
	}
L101:
	;
	v270 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
	v271 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+40)))
	v774 = v270 + v271*int64(1000000)
	goto L77
L102:
	;
	v265 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	v268 = base.I32_div_s(int32(0)-v265, int32(3600))
	v774 = base.I64_extend_i32_s(v268)
	goto L77
L103:
	;
	v257 = *(*int32)(unsafe.Add(mBase, uint32(v14)+92))
	v259 = int32(60)
	v260 = base.I32_div_s(int32(0)-v257, v259)
	v262 = base.I32_rem_s(v260, v259)
	v774 = base.I64_extend_i32_s(v262)
	goto L77
L104:
	;
	v276 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
	v282 = F_int64_div_fast_to_numeric(m, v276+base.I64_extend_i32_s(v275)*int64(1000000), int32(3))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L107
	}
L105:
	;
	goto L106
L106:
	;
	v285 = float64(1000)
	v287 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	v292 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v275), v285), base.F64_div(base.F64_convert_i32_s(v287), v285)))
	mBase = m.M
	v293 = m.ExcPending
	if v293 != 0 {
		goto L1
	} else {
		goto L108
	}
L107:
	;
	v785 = v282
	goto L60
L108:
	;
	v785 = v292
	goto L60
L109:
	;
	v295 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
	v301 = F_int64_div_fast_to_numeric(m, v295+base.I64_extend_i32_s(v294)*int64(1000000), int32(6))
	mBase = m.M
	v302 = m.ExcPending
	if v302 != 0 {
		goto L1
	} else {
		goto L112
	}
L110:
	;
	goto L111
L111:
	;
	v303 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	v309 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v303), float64(1e+06)), base.F64_convert_i32_s(v294)))
	mBase = m.M
	v310 = m.ExcPending
	if v310 != 0 {
		goto L1
	} else {
		goto L113
	}
L112:
	;
	v785 = v301
	goto L60
L113:
	;
	v785 = v309
	goto L60
L114:
	;
	v774 = base.I64_extend_i32_s(v365 + int32(1))
	goto L77
L115:
	;
	v336 = int32(1)
	v340 = F_date2j(m, v323-v336, v336, int32(4))
	mBase = m.M
	v343 = F_j2day(m, v340-v336)
	mBase = m.M
	v344 = v340
	v345 = v343
	goto L117
L116:
	;
	v344 = v330
	v345 = v333
	goto L117
L117:
	;
	v347 = v345 - v344 + v327
	if int32(357) <= v347 {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v350 = int32(1)
	v354 = F_date2j(m, v323+v350, v350, int32(4))
	mBase = m.M
	v357 = F_j2day(m, v354-v350)
	mBase = m.M
	v358 = v354 - v357
	if v327 < v358 {
		goto L121
	} else {
		goto L122
	}
L119:
	;
	v363 = v347
	goto L120
L120:
	;
	v365 = base.I32_div_s(v363, int32(7))
	goto L114
L121:
	;
	v361 = v347
	goto L123
L122:
	;
	v361 = v327 - v358
	goto L123
L123:
	;
	v363 = v361
	goto L120
L124:
	;
	v774 = base.I64_extend_i32_u(v369)
	goto L77
L125:
	;
	goto L126
L126:
	;
	v774 = base.I64_extend_i32_s(v369 - int32(1))
	goto L77
L127:
	;
	v380 = base.I32_div_u_s(v376, int32(10))
	v774 = base.I64_extend_i32_u(v380)
	goto L77
L128:
	;
	goto L129
L129:
	;
	v385 = base.I32_div_s(int32(9)-v376, int32(-10))
	v774 = base.I64_extend_i32_s(v385)
	goto L77
L130:
	;
	v393 = base.I32_div_s(v387+int32(99), int32(100))
	v774 = base.I64_extend_i32_s(v393)
	goto L77
L131:
	;
	goto L132
L132:
	;
	v398 = base.I32_div_s(int32(100)-v387, int32(-100))
	v774 = base.I64_extend_i32_s(v398)
	goto L77
L133:
	;
	v406 = base.I32_div_s(v400+int32(999), int32(1000))
	v774 = base.I64_extend_i32_s(v406)
	goto L77
L134:
	;
	goto L135
L135:
	;
	v411 = base.I32_div_s(int32(1000)-v400, int32(-1000))
	v774 = base.I64_extend_i32_s(v411)
	goto L77
L136:
	;
	if l1 != 0 {
		goto L143
	} else {
		goto L144
	}
L137:
	;
	v421 = int32(4800)
	goto L139
L138:
	;
	v421 = int32(4799)
	goto L139
L139:
	;
	v422 = v421 + v413
	v427 = base.I32_div_s(v422, int32(4))
	v430 = base.I32_div_s(v422, int32(-100))
	v433 = base.I32_div_s(v422, int32(400))
	if int32(2) < v414 {
		goto L140
	} else {
		goto L141
	}
L140:
	;
	v437 = int32(1)
	goto L142
L141:
	;
	v437 = int32(13)
	goto L142
L142:
	;
	v442 = base.I32_div_s((v437+v414)*int32(7834), int32(256))
	v445 = v415 + v422*int32(365) + v427 + v430 + v433 + v442 - int32(32167)
	goto L136
L143:
	;
	v447 = F_int64_to_numeric(m, base.I64_extend_i32_s(v445))
	mBase = m.M
	v448 = m.ExcPending
	if v448 != 0 {
		goto L1
	} else {
		goto L146
	}
L144:
	;
	goto L145
L145:
	;
	v474 = *(*int32)(unsafe.Add(mBase, uint32(v14)+84))
	v478 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v479 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v480 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v481 = int32(60)
	v493 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_add(base.F64_div(base.F64_convert_i32_s(v474), float64(1e+06)), base.F64_convert_i32_s(v478+(v479+v480*v481)*v481)), float64(86400)), base.F64_convert_i32_s(v445)))
	mBase = m.M
	v494 = m.ExcPending
	if v494 != 0 {
		goto L1
	} else {
		goto L151
	}
L146:
	;
	v449 = int64(*(*int32)(unsafe.Add(mBase, uint32(v14)+84)))
	v450 = *(*int32)(unsafe.Add(mBase, uint32(v14)+40))
	v451 = *(*int32)(unsafe.Add(mBase, uint32(v14)+44))
	v452 = *(*int32)(unsafe.Add(mBase, uint32(v14)+48))
	v453 = int32(60)
	v463 = F_int64_to_numeric(m, v449+base.I64_extend_i32_s(v450+(v451+v452*v453)*v453)*int64(1000000))
	mBase = m.M
	v464 = m.ExcPending
	if v464 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	v466 = F_int64_to_numeric(m, int64(86400000000))
	mBase = m.M
	v467 = m.ExcPending
	if v467 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	v469 = F_numeric_div_opt_error(m, v463, v466, int32(0))
	mBase = m.M
	v470 = m.ExcPending
	if v470 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	v472 = F_numeric_add_opt_error(m, v447, v469, int32(0))
	mBase = m.M
	v473 = m.ExcPending
	if v473 != 0 {
		goto L1
	} else {
		goto L150
	}
L150:
	;
	v785 = v472
	goto L60
L151:
	;
	v785 = v493
	goto L60
L152:
	;
	v774 = base.I64_extend_i32_s(v536) - base.I64_extend_i32_u(base.B2i32(v536 <= int32(0)))
	goto L77
L153:
	;
	v508 = int32(1)
	v509 = v495 - v508
	v512 = F_date2j(m, v509, v508, int32(4))
	mBase = m.M
	v515 = F_j2day(m, v512-v508)
	mBase = m.M
	v516 = v509
	v517 = v512
	v518 = v515
	goto L155
L154:
	;
	v516 = v495
	v517 = v502
	v518 = v505
	goto L155
L155:
	;
	if int32(357) <= v499+v518-v517 {
		goto L156
	} else {
		goto L157
	}
L156:
	;
	v523 = int32(1)
	v524 = v516 + v523
	v527 = F_date2j(m, v524, v523, int32(4))
	mBase = m.M
	v530 = F_j2day(m, v527-v523)
	mBase = m.M
	if v499 < v527-v530 {
		goto L159
	} else {
		goto L160
	}
L157:
	;
	v536 = v516
	goto L158
L158:
	;
	goto L152
L159:
	;
	v533 = v516
	goto L161
L160:
	;
	v533 = v524
	goto L161
L161:
	;
	v536 = v533
	goto L158
L162:
	;
	v577 = int32(7)
	v578 = base.I32_rem_s(v544+v551*int32(365)+v556+v559+v562+v571-int32(32167)+int32(1), v577)
	if v578 < int32(0) {
		goto L170
	} else {
		goto L171
	}
L163:
	;
	v550 = int32(4800)
	goto L165
L164:
	;
	v550 = int32(4799)
	goto L165
L165:
	;
	v551 = v550 + v542
	v556 = base.I32_div_s(v551, int32(4))
	v559 = base.I32_div_s(v551, int32(-100))
	v562 = base.I32_div_s(v551, int32(400))
	if int32(2) < v543 {
		goto L166
	} else {
		goto L167
	}
L166:
	;
	v566 = int32(1)
	goto L168
L167:
	;
	v566 = int32(13)
	goto L168
L168:
	;
	v571 = base.I32_div_s((v566+v543)*int32(7834), int32(256))
	goto L162
L169:
	;
	v584 = base.I64_extend_i32_s(v583)
	if v583 != 0 {
		goto L173
	} else {
		goto L174
	}
L170:
	;
	v583 = v578 + v577
	goto L172
L171:
	;
	v583 = v578
	goto L172
L172:
	;
	goto L169
L173:
	;
	v586 = v584
	goto L175
L174:
	;
	v586 = int64(7)
	goto L175
L175:
	;
	v587 = *(*int32)(unsafe.Add(mBase, uint32(v14)+88))
	if v587 == int32(37) {
		goto L176
	} else {
		goto L177
	}
L176:
	;
	v590 = v586
	goto L178
L177:
	;
	v590 = v584
	goto L178
L178:
	;
	v774 = v590
	goto L77
L179:
	;
	v624 = *(*int32)(unsafe.Add(mBase, uint32(v14)+60))
	goto L188
L180:
	;
	v599 = int32(4800)
	goto L182
L181:
	;
	v599 = int32(4799)
	goto L182
L182:
	;
	v600 = v599 + v591
	v605 = base.I32_div_s(v600, int32(4))
	v608 = base.I32_div_s(v600, int32(-100))
	v611 = base.I32_div_s(v600, int32(400))
	if int32(2) < v592 {
		goto L183
	} else {
		goto L184
	}
L183:
	;
	v615 = int32(1)
	goto L185
L184:
	;
	v615 = int32(13)
	goto L185
L185:
	;
	v620 = base.I32_div_s((v615+v592)*int32(7834), int32(256))
	goto L179
L186:
	;
	v774 = base.I64_extend_i32_s(v593 + v600*int32(365) + v605 + v608 + v611 + v620 - int32(32167) - (int32(1) + v633*int32(365) + v638 + v641 + v644 + v653 - int32(32167)) + int32(1))
	goto L77
L188:
	;
	goto L189
L189:
	;
	v633 = int32(4799) + v624
	v638 = base.I32_div_s(v633, int32(4))
	v641 = base.I32_div_s(v633, int32(-100))
	v644 = base.I32_div_s(v633, int32(400))
	goto L191
L191:
	;
	goto L192
L192:
	;
	v653 = base.I32_div_s(int32(109676), int32(256))
	goto L186
L193:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v667 = m.ExcPending
	if v667 != 0 {
		goto L1
	} else {
		goto L194
	}
L194:
	;
	v669 = F_format_type_be(m, int32(1184))
	mBase = m.M
	v670 = m.ExcPending
	if v670 != 0 {
		goto L1
	} else {
		goto L195
	}
L195:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v669
	*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v60
	F_errmsg(m, int32(188314), v14+int32(16))
	mBase = m.M
	v677 = m.ExcPending
	if v677 != 0 {
		goto L1
	} else {
		goto L196
	}
L196:
	;
	F_errfinish(m, int32(490442), int32(5961), int32(243067))
	mBase = m.M
	v682 = m.ExcPending
	if v682 != 0 {
		goto L1
	} else {
		goto L197
	}
L197:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L198:
	;
	v686 = F_SetEpochTimestamp(m)
	mBase = m.M
	v687 = m.ExcPending
	if v687 != 0 {
		goto L1
	} else {
		goto L201
	}
L199:
	;
	goto L200
L200:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v729 = m.ExcPending
	if v729 != 0 {
		goto L1
	} else {
		goto L220
	}
L201:
	;
	v689 = v686 + int64(9223372036854775807)
	if l1 != 0 {
		goto L202
	} else {
		goto L203
	}
L202:
	;
	if v22 < v689 {
		goto L205
	} else {
		goto L206
	}
L203:
	;
	goto L204
L204:
	;
	if v22 < v689 {
		goto L216
	} else {
		goto L217
	}
L205:
	;
	v693 = F_int64_div_fast_to_numeric(m, v22-v686, int32(6))
	mBase = m.M
	v694 = m.ExcPending
	if v694 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	v697 = F_int64_to_numeric(m, v22)
	mBase = m.M
	v698 = m.ExcPending
	if v698 != 0 {
		goto L1
	} else {
		goto L209
	}
L208:
	;
	v785 = v693
	goto L60
L209:
	;
	v699 = F_int64_to_numeric(m, v686)
	mBase = m.M
	v700 = m.ExcPending
	if v700 != 0 {
		goto L1
	} else {
		goto L210
	}
L210:
	;
	v702 = F_numeric_sub_opt_error(m, v697, v699, int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L211
	}
L211:
	;
	v705 = F_int64_to_numeric(m, int64(1000000))
	mBase = m.M
	v706 = m.ExcPending
	if v706 != 0 {
		goto L1
	} else {
		goto L212
	}
L212:
	;
	v708 = F_numeric_div_opt_error(m, v702, v705, int32(0))
	mBase = m.M
	v709 = m.ExcPending
	if v709 != 0 {
		goto L1
	} else {
		goto L213
	}
L213:
	;
	v711 = F_DirectFunctionCall2Coll(m, int32(1275), int32(0), v708, int32(6))
	mBase = m.M
	v712 = m.ExcPending
	if v712 != 0 {
		goto L1
	} else {
		goto L214
	}
L214:
	;
	v713 = F_pg_detoast_datum(m, v711)
	mBase = m.M
	v714 = m.ExcPending
	if v714 != 0 {
		goto L1
	} else {
		goto L215
	}
L215:
	;
	v785 = v713
	goto L60
L216:
	;
	v721 = base.F64_convert_i64_s(v22 - v686)
	goto L218
L217:
	;
	v721 = base.F64_sub(base.F64_convert_i64_s(v22), base.F64_convert_i64_s(v686))
	goto L218
L218:
	;
	v724 = F_Float8GetDatum(m, base.F64_div(v721, float64(1e+06)))
	mBase = m.M
	v725 = m.ExcPending
	if v725 != 0 {
		goto L1
	} else {
		goto L219
	}
L219:
	;
	v785 = v724
	goto L60
L220:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v732 = m.ExcPending
	if v732 != 0 {
		goto L1
	} else {
		goto L221
	}
L221:
	;
	v734 = F_format_type_be(m, int32(1184))
	mBase = m.M
	v735 = m.ExcPending
	if v735 != 0 {
		goto L1
	} else {
		goto L222
	}
L222:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+36)) = v734
	*(*int32)(unsafe.Add(mBase, uint32(v14)+32)) = v60
	F_errmsg(m, int32(188314), v14+int32(32))
	mBase = m.M
	v742 = m.ExcPending
	if v742 != 0 {
		goto L1
	} else {
		goto L223
	}
L223:
	;
	F_errfinish(m, int32(490442), int32(6008), int32(243067))
	mBase = m.M
	v747 = m.ExcPending
	if v747 != 0 {
		goto L1
	} else {
		goto L224
	}
L224:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L225:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v754 = m.ExcPending
	if v754 != 0 {
		goto L1
	} else {
		goto L226
	}
L226:
	;
	v756 = F_format_type_be(m, int32(1184))
	mBase = m.M
	v757 = m.ExcPending
	if v757 != 0 {
		goto L1
	} else {
		goto L227
	}
L227:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v756
	*(*int32)(unsafe.Add(mBase, uint32(v14))) = v60
	F_errmsg(m, int32(188277), v14)
	mBase = m.M
	v762 = m.ExcPending
	if v762 != 0 {
		goto L1
	} else {
		goto L228
	}
L228:
	;
	F_errfinish(m, int32(490442), int32(6017), int32(243067))
	mBase = m.M
	v767 = m.ExcPending
	if v767 != 0 {
		goto L1
	} else {
		goto L229
	}
L229:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L230:
	;
	v775 = F_int64_to_numeric(m, v774)
	mBase = m.M
	v776 = m.ExcPending
	if v776 != 0 {
		goto L1
	} else {
		goto L233
	}
L231:
	;
	goto L232
L232:
	;
	v778 = F_Float8GetDatum(m, base.F64_convert_i64_s(v774))
	mBase = m.M
	v779 = m.ExcPending
	if v779 != 0 {
		goto L1
	} else {
		goto L234
	}
L233:
	;
	v785 = v775
	goto L60
L234:
	;
	v785 = v778
	goto L60
L235:
	;
	F_errcode(m, int32(134217858))
	mBase = m.M
	v796 = m.ExcPending
	if v796 != 0 {
		goto L1
	} else {
		goto L236
	}
L236:
	;
	F_errmsg(m, int32(398126), int32(0))
	mBase = m.M
	v800 = m.ExcPending
	if v800 != 0 {
		goto L1
	} else {
		goto L237
	}
L237:
	;
	F_errfinish(m, int32(490442), int32(5827), int32(243067))
	mBase = m.M
	v805 = m.ExcPending
	if v805 != 0 {
		goto L1
	} else {
		goto L238
	}
L238:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_timestamptz_to_str(m *base.Module, l0 int64) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v5 int32
	_ = v5
	var v15 int32
	_ = v15
	var v19 int64
	_ = v19
	var v23 int32
	_ = v23
	var v27 int64
	_ = v27
	var v38 int32
	_ = v38
	var v41 int32
	_ = v41
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v56 int32
	_ = v56
	var v60 int64
	_ = v60
	var v64 int64
	_ = v64
	var v68 int64
	_ = v68
	v3 = m.G0
	v5 = v3 + int32(-64)
	m.G0 = v5
	if base.Ui64(l0-int64(9223372036854775807)) <= base.Ui64(int64(1)) {
		if l0 == int64(-9223372036854775807-1) {
			v15 = int32(*(*uint16)(unsafe.Add(mBase, _consts[1330])))
			*(*uint16)(unsafe.Add(mBase, _consts[1331])) = uint16(v15)
			v19 = *(*int64)(unsafe.Add(mBase, _consts[1332]))
			*(*int64)(unsafe.Add(mBase, _consts[1333])) = v19
		} else {
			v23 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1334])))
			*(*uint8)(unsafe.Add(mBase, _consts[1331])) = uint8(v23)
			v27 = *(*int64)(unsafe.Add(mBase, _consts[1335]))
			*(*int64)(unsafe.Add(mBase, _consts[1333])) = v27
		}
		m.G0 = v5 - int32(-64)
		return int32(4462752)
	} else {
		v38 = F_timestamp2tm(m, l0, v3+int32(-4), v3+int32(-48), v3+int32(-52), v3+int32(-56), int32(0))
		mBase = m.M
		v41 = m.ExcPending
		if v41 != 0 {
			return int32(0)
		} else {
			if v38 == int32(0) {
				v46 = *(*int32)(unsafe.Add(mBase, uint32(v5)+12))
				v47 = int32(1)
				v48 = *(*int32)(unsafe.Add(mBase, uint32(v5)+60))
				v49 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
				F_EncodeDateTime(m, v3+int32(-48), v46, v47, v48, v49, v47, int32(4462752))
				mBase = m.M
				v53 = m.ExcPending
				if v53 != 0 {
					return int32(0)
				} else {
					m.G0 = v5 - int32(-64)
					return int32(4462752)
				}
			} else {
				v56 = int32(*(*uint8)(unsafe.Add(mBase, _consts[1336])))
				*(*uint8)(unsafe.Add(mBase, _consts[1337])) = uint8(v56)
				v60 = *(*int64)(unsafe.Add(mBase, _consts[1338]))
				*(*int64)(unsafe.Add(mBase, _consts[1339])) = v60
				v64 = *(*int64)(unsafe.Add(mBase, _consts[1340]))
				*(*int64)(unsafe.Add(mBase, _consts[1331])) = v64
				v68 = *(*int64)(unsafe.Add(mBase, _consts[1341]))
				*(*int64)(unsafe.Add(mBase, _consts[1333])) = v68
				m.G0 = v5 - int32(-64)
				return int32(4462752)
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
