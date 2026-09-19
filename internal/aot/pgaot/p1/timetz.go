package p1

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_timetz_at_local(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, _c_F_timetz_at_local[0]))
	v7 = F_cstring_to_text(m, v6)
	mBase = m.M
	v10 = m.ExcPending
	if v10 != 0 {
		return int32(0)
	} else {
		v11 = F_DirectFunctionCall2Coll(m, int32(1274), int32(0), v7, v2)
		mBase = m.M
		v12 = m.ExcPending
		if v12 != 0 {
			return int32(0)
		} else {
			return v11
		}
	}
}
func F_timetz_gt(m *base.Module, l0 int32) int32 {
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
	v6 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v7 = *(*int32)(unsafe.Add(mBase, uint32(v6)+8))
	v9 = int64(1000000)
	v11 = *(*int64)(unsafe.Add(mBase, uint32(v6)))
	v12 = base.I64_extend_i32_s(v7)*v9 + v11
	v13 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
	v14 = *(*int32)(unsafe.Add(mBase, uint32(v13)+8))
	v18 = *(*int64)(unsafe.Add(mBase, uint32(v13)))
	v19 = base.I64_extend_i32_s(v14)*v9 + v18
	if v19 < v12 {
		return int32(1)
	} else {
		if v12 < v19 {
			return int32(0)
		} else {
			return base.B2i32(v14 < v7)
		}
	}
}
func F_timetz_part_common(m *base.Module, l0 int32, l1 int32) int32 {
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
	var v22 int32
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
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v80 int64
	_ = v80
	var v83 int64
	_ = v83
	var v85 int64
	_ = v85
	var v86 int64
	_ = v86
	var v89 int64
	_ = v89
	var v91 int64
	_ = v91
	var v94 int64
	_ = v94
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v106 int32
	_ = v106
	var v111 int32
	_ = v111
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v127 float64
	_ = v127
	var v133 int32
	_ = v133
	var v134 int32
	_ = v134
	var v141 int32
	_ = v141
	var v142 int32
	_ = v142
	var v148 int32
	_ = v148
	var v149 int32
	_ = v149
	var v153 int32
	_ = v153
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v164 int32
	_ = v164
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v173 int64
	_ = v173
	var v174 int64
	_ = v174
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	var v192 int32
	_ = v192
	var v195 int32
	_ = v195
	var v197 int32
	_ = v197
	var v198 int32
	_ = v198
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v214 int64
	_ = v214
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v227 int32
	_ = v227
	v14 = m.G0
	v16 = v14 - int32(32)
	m.G0 = v16
	v18 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v19 = F_pg_detoast_datum_packed(m, v18)
	mBase = m.M
	v22 = m.ExcPending
	if v22 != 0 {
		return int32(0)
	} else {
		v23 = int32(1)
		v24 = v19 + v23
		v27 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v19))))
		v29 = v27 & v23
		if v29 != 0 {
			v30 = v24
		} else {
			v30 = v19 + int32(4)
		}
		v31 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		if v27 == int32(1) {
			v37 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v24))))
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
			if v29 != 0 {
				v58 = int32(base.Ui32(v27)>>(uint(v48)%32)) - v48
			} else {
				v52 = *(*int32)(unsafe.Add(mBase, uint32(v19)))
				v58 = int32(base.Ui32(v52)>>(uint(int32(2))%32)) - int32(4)
			}
		}
		v60 = F_downcase_truncate_identifier(m, v30, v58, int32(0))
		mBase = m.M
		v61 = m.ExcPending
		if v61 != 0 {
			return int32(0)
		} else {
			v63 = v16 + int32(28)
			v67 = Fn13846(m, v60, v63, int32(_a_F_timetz_part_common_0), int32(_a_F_timetz_part_common_1), int32(_a_F_timetz_part_common_2))
			mBase = m.M
			if v67 == int32(31) {
				v73 = Fn13846(m, v60, v63, int32(_a_F_timetz_part_common_3), int32(_a_F_timetz_part_common_4), int32(_a_F_timetz_part_common_5))
				mBase = m.M
				v74 = v73
			} else {
				v74 = v67
			}
			if v74 == int32(17) {
				v77 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
				v79 = base.I64_div_s(v77, int64(3600000000))
				v80 = base.I64_extend32_s(v79)
				v83 = v80*int64(-3600000000) + v77
				v85 = base.I64_div_s(v83, int64(60000000))
				v86 = base.I64_extend32_s(v85)
				v89 = v86*int64(-60000000) + v83
				v91 = base.I64_div_s(v89, int64(1000000))
				v94 = v91*int64(4293967296) + v89
				v95 = base.I32_wrap_i64(v94)
				v96 = base.I32_wrap_i64(v91)
				v97 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
				v98 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
				switch v98 - int32(4) {
				case 0:
					v214 = base.I64_extend_i32_s(int32(0) - v97)
					if l1 != 0 {
						v215 = F_int64_to_numeric(m, v214)
						mBase = m.M
						v216 = m.ExcPending
						if v216 != 0 {
							return int32(0)
						} else {
							v227 = v215
							m.G0 = v16 + int32(32)
							return v227
						}
					} else {
						v218 = F_Float8GetDatum(m, base.F64_convert_i64_s(v214))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							v227 = v218
							m.G0 = v16 + int32(32)
							return v227
						}
					}
				default:
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v153 = m.ExcPending
					if v153 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(1088))
						mBase = m.M
						v156 = m.ExcPending
						if v156 != 0 {
							return int32(0)
						} else {
							v158 = F_format_type_be(m, int32(1266))
							mBase = m.M
							v159 = m.ExcPending
							if v159 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+4)) = v158
								*(*int32)(unsafe.Add(mBase, uint32(v16))) = v60
								F_errmsg(m, int32(_a_F_timetz_part_common_6), v16)
								mBase = m.M
								v164 = m.ExcPending
								if v164 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_timetz_part_common_7), int32(3077), int32(_a_F_timetz_part_common_8))
									mBase = m.M
									v169 = m.ExcPending
									if v169 != 0 {
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
					if l1 != 0 {
						v141 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v94)+base.I64_extend32_s(v91)*int64(1000000), int32(6))
						mBase = m.M
						v142 = m.ExcPending
						if v142 != 0 {
							return int32(0)
						} else {
							v227 = v141
							m.G0 = v16 + int32(32)
							return v227
						}
					} else {
						v148 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i32_s(v95), float64(1e+06)), base.F64_convert_i32_s(v96)))
						mBase = m.M
						v149 = m.ExcPending
						if v149 != 0 {
							return int32(0)
						} else {
							v227 = v148
							m.G0 = v16 + int32(32)
							return v227
						}
					}
				case 15:
					v214 = v86
					if l1 != 0 {
						v215 = F_int64_to_numeric(m, v214)
						mBase = m.M
						v216 = m.ExcPending
						if v216 != 0 {
							return int32(0)
						} else {
							v227 = v215
							m.G0 = v16 + int32(32)
							return v227
						}
					} else {
						v218 = F_Float8GetDatum(m, base.F64_convert_i64_s(v214))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							v227 = v218
							m.G0 = v16 + int32(32)
							return v227
						}
					}
				case 16:
					v214 = v80
					if l1 != 0 {
						v215 = F_int64_to_numeric(m, v214)
						mBase = m.M
						v216 = m.ExcPending
						if v216 != 0 {
							return int32(0)
						} else {
							v227 = v215
							m.G0 = v16 + int32(32)
							return v227
						}
					} else {
						v218 = F_Float8GetDatum(m, base.F64_convert_i64_s(v214))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							v227 = v218
							m.G0 = v16 + int32(32)
							return v227
						}
					}
				case 25:
					if l1 != 0 {
						v124 = F_int64_div_fast_to_numeric(m, base.I64_extend32_s(v94)+base.I64_extend32_s(v91)*int64(1000000), int32(3))
						mBase = m.M
						v125 = m.ExcPending
						if v125 != 0 {
							return int32(0)
						} else {
							v227 = v124
							m.G0 = v16 + int32(32)
							return v227
						}
					} else {
						v127 = float64(1000)
						v133 = F_Float8GetDatum(m, base.F64_add(base.F64_mul(base.F64_convert_i32_s(v96), v127), base.F64_div(base.F64_convert_i32_s(v95), v127)))
						mBase = m.M
						v134 = m.ExcPending
						if v134 != 0 {
							return int32(0)
						} else {
							v227 = v133
							m.G0 = v16 + int32(32)
							return v227
						}
					}
				case 26:
					v214 = base.I64_extend32_s(v94) + base.I64_extend32_s(v91)*int64(1000000)
					if l1 != 0 {
						v215 = F_int64_to_numeric(m, v214)
						mBase = m.M
						v216 = m.ExcPending
						if v216 != 0 {
							return int32(0)
						} else {
							v227 = v215
							m.G0 = v16 + int32(32)
							return v227
						}
					} else {
						v218 = F_Float8GetDatum(m, base.F64_convert_i64_s(v214))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							v227 = v218
							m.G0 = v16 + int32(32)
							return v227
						}
					}
				case 30:
					v111 = base.I32_div_s(int32(0)-v97, int32(3600))
					v214 = base.I64_extend_i32_s(v111)
					if l1 != 0 {
						v215 = F_int64_to_numeric(m, v214)
						mBase = m.M
						v216 = m.ExcPending
						if v216 != 0 {
							return int32(0)
						} else {
							v227 = v215
							m.G0 = v16 + int32(32)
							return v227
						}
					} else {
						v218 = F_Float8GetDatum(m, base.F64_convert_i64_s(v214))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							v227 = v218
							m.G0 = v16 + int32(32)
							return v227
						}
					}
				case 31:
					v103 = int32(60)
					v104 = base.I32_div_s(int32(0)-v97, v103)
					v106 = base.I32_rem_s(v104, v103)
					v214 = base.I64_extend_i32_s(v106)
					if l1 != 0 {
						v215 = F_int64_to_numeric(m, v214)
						mBase = m.M
						v216 = m.ExcPending
						if v216 != 0 {
							return int32(0)
						} else {
							v227 = v215
							m.G0 = v16 + int32(32)
							return v227
						}
					} else {
						v218 = F_Float8GetDatum(m, base.F64_convert_i64_s(v214))
						mBase = m.M
						v219 = m.ExcPending
						if v219 != 0 {
							return int32(0)
						} else {
							v227 = v218
							m.G0 = v16 + int32(32)
							return v227
						}
					}
				}
			} else {
				if v74 != 0 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v192 = m.ExcPending
					if v192 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v195 = m.ExcPending
						if v195 != 0 {
							return int32(0)
						} else {
							v197 = F_format_type_be(m, int32(1266))
							mBase = m.M
							v198 = m.ExcPending
							if v198 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v197
								*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v60
								F_errmsg(m, int32(_a_F_timetz_part_common_9), v16+int32(16))
								mBase = m.M
								v205 = m.ExcPending
								if v205 != 0 {
									return int32(0)
								} else {
									F_errfinish(m, int32(_a_F_timetz_part_common_7), int32(3097), int32(_a_F_timetz_part_common_8))
									mBase = m.M
									v210 = m.ExcPending
									if v210 != 0 {
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
					v170 = *(*int32)(unsafe.Add(mBase, uint32(v16)+28))
					if v170 != int32(11) {
						F_errstart_cold(m, int32(21), int32(0))
						mBase = m.M
						v192 = m.ExcPending
						if v192 != 0 {
							return int32(0)
						} else {
							F_errcode(m, int32(50856066))
							mBase = m.M
							v195 = m.ExcPending
							if v195 != 0 {
								return int32(0)
							} else {
								v197 = F_format_type_be(m, int32(1266))
								mBase = m.M
								v198 = m.ExcPending
								if v198 != 0 {
									return int32(0)
								} else {
									*(*int32)(unsafe.Add(mBase, uint32(v16)+20)) = v197
									*(*int32)(unsafe.Add(mBase, uint32(v16)+16)) = v60
									F_errmsg(m, int32(_a_F_timetz_part_common_9), v16+int32(16))
									mBase = m.M
									v205 = m.ExcPending
									if v205 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_timetz_part_common_7), int32(3097), int32(_a_F_timetz_part_common_8))
										mBase = m.M
										v210 = m.ExcPending
										if v210 != 0 {
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
						v173 = *(*int64)(unsafe.Add(mBase, uint32(v31)))
						if l1 != 0 {
							v174 = int64(*(*int32)(unsafe.Add(mBase, uint32(v31)+8)))
							v179 = F_int64_div_fast_to_numeric(m, v174*int64(1000000)+v173, int32(6))
							mBase = m.M
							v180 = m.ExcPending
							if v180 != 0 {
								return int32(0)
							} else {
								v227 = v179
								m.G0 = v16 + int32(32)
								return v227
							}
						} else {
							v184 = *(*int32)(unsafe.Add(mBase, uint32(v31)+8))
							v187 = F_Float8GetDatum(m, base.F64_add(base.F64_div(base.F64_convert_i64_s(v173), float64(1e+06)), base.F64_convert_i32_s(v184)))
							mBase = m.M
							v188 = m.ExcPending
							if v188 != 0 {
								return int32(0)
							} else {
								v227 = v187
								m.G0 = v16 + int32(32)
								return v227
							}
						}
					}
				}
			}
		}
	}
}
func F_timetz_time(m *base.Module, l0 int32) int32 {
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
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = *(*int64)(unsafe.Add(mBase, uint32(v2)))
	v4 = F_Int64GetDatum(m, v3)
	mBase = m.M
	v7 = m.ExcPending
	if v7 != 0 {
		return int32(0)
	} else {
		return v4
	}
}
