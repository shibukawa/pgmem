package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timetz_at_local(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v3 int64
	_ = v3
	var v7 int32
	_ = v7
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	v3 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, _c_F_timetz_at_local[0]))
	v8 = F_cstring_to_text(m, v7)
	mBase = m.M
	v11 = m.ExcPending
	if v11 != 0 {
		return int64(0)
	} else {
		v13 = F_DirectFunctionCall2Coll(m, int32(1402), int32(0), base.I64_extend_i32_u(v8), v3)
		mBase = m.M
		v14 = m.ExcPending
		if v14 != 0 {
			return int64(0)
		} else {
			return v13
		}
	}
}
func F_timetz_gt(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v7 int32
	_ = v7
	var v9 int64
	_ = v9
	var v11 int64
	_ = v11
	var v12 int64
	_ = v12
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v18 int64
	_ = v18
	var v19 int64
	_ = v19
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v9 = int64(1000000)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v12 = base.I64_extend_i32_s(v7)*v9 + v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v19 = base.I64_extend_i32_s(v14)*v9 + v18
	if v19 < v12 {
		return int64(1)
	} else {
		if v12 < v19 {
			return int64(0)
		} else {
			return base.I64_extend_i32_u(base.B2i32(v14 < v7))
		}
	}
}
func F_timetz_part_common(m *base.Module, l0 int32, l1 int32) int64 {
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
	var v28 int32
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
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v77 int64
	_ = v77
	var v80 int64
	_ = v80
	var v82 int64
	_ = v82
	var v83 int64
	_ = v83
	var v86 int64
	_ = v86
	var v88 int64
	_ = v88
	var v91 int64
	_ = v91
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v108 int32
	_ = v108
	var v121 int32
	_ = v121
	var v122 int32
	_ = v122
	var v125 float64
	_ = v125
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v150 int32
	_ = v150
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v170 int64
	_ = v170
	var v171 int64
	_ = v171
	var v176 int32
	_ = v176
	var v177 int32
	_ = v177
	var v182 int32
	_ = v182
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v202 int32
	_ = v202
	var v207 int32
	_ = v207
	var v211 int64
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v224 int64
	_ = v224
	v12 = m.G0
	v14 = v12 - int32(32)
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
		v28 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
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
			v60 = v14 + int32(28)
			v64 = Fn14210(m, v57, v60, int32(_a_F_timetz_part_common_0), int32(_a_F_timetz_part_common_1), int32(_a_F_timetz_part_common_2))
			mBase = m.M
			if v64 == int32(31) {
				v70 = Fn14210(m, v57, v60, int32(_a_F_timetz_part_common_3), int32(_a_F_timetz_part_common_4), int32(_a_F_timetz_part_common_5))
				mBase = m.M
				v71 = v70
			} else {
				v71 = v64
			}
			if v71 == int32(17) {
				v74 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
				v76 = base.I64_div_s(v74, int64(3600000000))
				v77 = base.I64_extend32_s(v76)
				v80 = v77*int64(-3600000000) + v74
				v82 = base.I64_div_s(v80, int64(60000000))
				v83 = base.I64_extend32_s(v82)
				v86 = v83*int64(-60000000) + v80
				v88 = base.I64_div_s(v86, int64(1000000))
				v91 = v88*int64(4293967296) + v86
				v92 = base.I32_wrap_i64(v91)
				v93 = base.I32_wrap_i64(v88)
				v94 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
				v95 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
				switch v95 - int32(4) {
				case 0:
					v211 = base.I64_extend_i32_s(int32(0) - v94)
					if l1 != 0 {
						v212 = F_int64_to_numeric(m, v211)
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int64(0)
						} else {
							v224 = base.I64_extend_i32_u(v212)
							m.G0 = v14 + int32(32)
							return v224
						}
					} else {
						v224 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v211))
						m.G0 = v14 + int32(32)
						return v224
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v150 = m.ExcPending
					if v150 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v153 = m.ExcPending
						if v153 != 0 {
							return int64(0)
						} else {
							v155 = F_format_type_be(m, int32(1266))
							mBase = m.M
							v156 = m.ExcPending
							if v156 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+4)) = v155
								*(*int32)(unsafe.Add(mBase, uint32(v14))) = v57
								F_errmsg(m, int32(_a_F_timetz_part_common_6), v14)
								mBase = m.M
								v161 = m.ExcPending
								if v161 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_timetz_part_common_7), int32(3139), int32(_a_F_timetz_part_common_8))
									mBase = m.M
									v166 = m.ExcPending
									if v166 != 0 {
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
					if l1 != 0 {
						v138 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v91)+base.I64_extend32_s(v88)*int64(1000000), int32(6))
						mBase = m.M
						v139 = m.ExcPending
						if v139 != 0 {
							return int64(0)
						} else {
							v224 = base.I64_extend_i32_u(v138)
							m.G0 = v14 + int32(32)
							return v224
						}
					} else {
						v224 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_convert_i32_s(v92), float64(1e+06)), base.F64_convert_i32_s(v93)))
						m.G0 = v14 + int32(32)
						return v224
					}
				case 15:
					v211 = v83
					if l1 != 0 {
						v212 = F_int64_to_numeric(m, v211)
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int64(0)
						} else {
							v224 = base.I64_extend_i32_u(v212)
							m.G0 = v14 + int32(32)
							return v224
						}
					} else {
						v224 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v211))
						m.G0 = v14 + int32(32)
						return v224
					}
				case 16:
					v211 = v77
					if l1 != 0 {
						v212 = F_int64_to_numeric(m, v211)
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int64(0)
						} else {
							v224 = base.I64_extend_i32_u(v212)
							m.G0 = v14 + int32(32)
							return v224
						}
					} else {
						v224 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v211))
						m.G0 = v14 + int32(32)
						return v224
					}
				case 25:
					if l1 != 0 {
						v121 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v91)+base.I64_extend32_s(v88)*int64(1000000), int32(3))
						mBase = m.M
						v122 = m.ExcPending
						if v122 != 0 {
							return int64(0)
						} else {
							v224 = base.I64_extend_i32_u(v121)
							m.G0 = v14 + int32(32)
							return v224
						}
					} else {
						v125 = float64(1000)
						v224 = base.I64_reinterpret_f64(base.F64_add(base.F64_mul(base.F64_convert_i32_s(v93), v125), base.F64_div(base.F64_convert_i32_s(v92), v125)))
						m.G0 = v14 + int32(32)
						return v224
					}
				case 26:
					v211 = base.I64_extend32_s(v91) + base.I64_extend32_s(v88)*int64(1000000)
					if l1 != 0 {
						v212 = F_int64_to_numeric(m, v211)
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int64(0)
						} else {
							v224 = base.I64_extend_i32_u(v212)
							m.G0 = v14 + int32(32)
							return v224
						}
					} else {
						v224 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v211))
						m.G0 = v14 + int32(32)
						return v224
					}
				case 30:
					v108 = base.I32_div_s(int32(0)-v94, int32(3600))
					v211 = base.I64_extend_i32_s(v108)
					if l1 != 0 {
						v212 = F_int64_to_numeric(m, v211)
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int64(0)
						} else {
							v224 = base.I64_extend_i32_u(v212)
							m.G0 = v14 + int32(32)
							return v224
						}
					} else {
						v224 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v211))
						m.G0 = v14 + int32(32)
						return v224
					}
				case 31:
					v100 = int32(60)
					v101 = base.I32_div_s(int32(0)-v94, v100)
					v103 = base.I32_rem_s(v101, v100)
					v211 = base.I64_extend_i32_s(v103)
					if l1 != 0 {
						v212 = F_int64_to_numeric(m, v211)
						mBase = m.M
						v213 = m.ExcPending
						if v213 != 0 {
							return int64(0)
						} else {
							v224 = base.I64_extend_i32_u(v212)
							m.G0 = v14 + int32(32)
							return v224
						}
					} else {
						v224 = base.I64_reinterpret_f64(base.F64_convert_i64_s(v211))
						m.G0 = v14 + int32(32)
						return v224
					}
				}
			} else {
				if v71 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v189 = m.ExcPending
					if v189 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return int64(0)
						} else {
							v194 = F_format_type_be(m, int32(1266))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v194
								*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v57
								F_errmsg(m, int32(_a_F_timetz_part_common_9), v14+int32(16))
								mBase = m.M
								v202 = m.ExcPending
								if v202 != 0 {
									return int64(0)
								} else {
									F_errfinish(m, int32(_a_F_timetz_part_common_7), int32(3159), int32(_a_F_timetz_part_common_8))
									mBase = m.M
									v207 = m.ExcPending
									if v207 != 0 {
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
					v167 = *(*int32)(unsafe.Add(mBase, uint32(v14)+28))
					if v167 != int32(11) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v189 = m.ExcPending
						if v189 != 0 {
							return int64(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v192 = m.ExcPending
							if v192 != 0 {
								return int64(0)
							} else {
								v194 = F_format_type_be(m, int32(1266))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int64(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v14)+20)) = v194
									*(*int32)(unsafe.Add(mBase, uint32(v14)+16)) = v57
									F_errmsg(m, int32(_a_F_timetz_part_common_9), v14+int32(16))
									mBase = m.M
									v202 = m.ExcPending
									if v202 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_timetz_part_common_7), int32(3159), int32(_a_F_timetz_part_common_8))
										mBase = m.M
										v207 = m.ExcPending
										if v207 != 0 {
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
						v170 = *(*int64)(unsafe.Add(mBase, uint32(v28)))
						if l1 != 0 {
							v171 = int64(*(*int32)(unsafe.Add(mBase, uint32(v28)+8)))
							v176 = F_int64_div_fast_to_numeric(m, v171*int64(1000000)+v170, int32(6))
							mBase = m.M
							v177 = m.ExcPending
							if v177 != 0 {
								return int64(0)
							} else {
								v224 = base.I64_extend_i32_u(v176)
								m.G0 = v14 + int32(32)
								return v224
							}
						} else {
							v182 = *(*int32)(unsafe.Add(mBase, uint32(v28)+8))
							v224 = base.I64_reinterpret_f64(base.F64_add(base.F64_div(base.F64_convert_i64_s(v170), float64(1e+06)), base.F64_convert_i32_s(v182)))
							m.G0 = v14 + int32(32)
							return v224
						}
					}
				}
			}
		}
	}
}
func F_timetz_time(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int64
	_ = v3
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	return v3
}
