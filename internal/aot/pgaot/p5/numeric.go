package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_numeric(m *base.Module, l0 int32) int64 {
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
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v39 int32
	_ = v39
	var v45 int32
	_ = v45
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v82 int32
	_ = v82
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	var v107 int32
	_ = v107
	var v112 int32
	_ = v112
	var v123 int32
	_ = v123
	var v136 int32
	_ = v136
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v158 int32
	_ = v158
	var v159 int32
	_ = v159
	var v160 int32
	_ = v160
	var v171 int32
	_ = v171
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v192 int32
	_ = v192
	var v193 int32
	_ = v193
	var v197 int32
	_ = v197
	var v200 int32
	_ = v200
	v8 = m.G0
	v10 = v8 - int32(32)
	m.G0 = v10
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v13 = F_pg_detoast_datum(m, v12)
	mBase = m.M
	v16 = m.ExcPending
	if v16 != 0 {
		return int64(0)
	} else {
		v17 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
		v19 = base.I32_extend16_s(v18)
		if base.Ui32(int32(_a_F_numeric_0)) <= base.Ui32(v18) {
			if base.B2i32(v19 == int32(-16384))|base.B2i32(v17 < int32(4)) == int32(0) {
				v30 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v31 = F_errsave_start(m, v30)
				mBase = m.M
				v32 = m.ExcPending
				if v32 != 0 {
					return int64(0)
				} else {
					if v31 != 0 {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v35 = m.ExcPending
						if v35 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_numeric_1), int32(0))
							mBase = m.M
							v39 = m.ExcPending
							if v39 != 0 {
								return int64(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(base.Ui32(v17-int32(4)) >> (uint(int32(16)) % 32))
								v45 = int32(21)
								*(*int32)(unsafe.Add(mBase, uint32(v10)+4)) = (v17<<(uint(v45)%32) - int32(_a_F_numeric_2)) >> (uint(v45) % 32)
								v53 = F_errdetail(m, int32(_a_F_numeric_3), v10)
								mBase = m.M
								v54 = m.ExcPending
								if v54 != 0 {
									return int64(0)
								} else {
									F_errsave_finish(m, v30, int32(_a_F_numeric_4), int32(_a_F_numeric_5), int32(_a_F_numeric_6))
									mBase = m.M
									v59 = m.ExcPending
									if v59 != 0 {
										return int64(0)
									} else {
										v60 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v60)
										v200 = int32(0)
										m.G0 = v10 + int32(32)
										return base.I64_extend_i32_u(v200)
									}
								}
							}
						}
					} else {
						v60 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v60)
						v200 = int32(0)
						m.G0 = v10 + int32(32)
						return base.I64_extend_i32_u(v200)
					}
				}
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v65 = F_palloc(m, int32(base.Ui32(v62)>>(uint(int32(2))%32)))
				mBase = m.M
				v66 = m.ExcPending
				if v66 != 0 {
					return int64(0)
				} else {
					v67 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v69 = int32(base.Ui32(v67) >> (uint(int32(2)) % 32))
					if v69 == int32(0) {
						v200 = v65
					} else {
						base.MemoryCopy(m, v65, v13, v69)
						v200 = v65
					}
					m.G0 = v10 + int32(32)
					return base.I64_extend_i32_u(v200)
				}
			}
		} else {
			if v17 <= int32(3) {
				v75 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
				v78 = F_palloc(m, int32(base.Ui32(v75)>>(uint(int32(2))%32)))
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int64(0)
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v82 = int32(base.Ui32(v80) >> (uint(int32(2)) % 32))
					if v82 == int32(0) {
						v200 = v78
					} else {
						base.MemoryCopy(m, v78, v13, v82)
						v200 = v78
					}
					m.G0 = v10 + int32(32)
					return base.I64_extend_i32_u(v200)
				}
			} else {
				if v19 < int32(0) {
					v98 = v18<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v18&int32(63)
				} else {
					v97 = int32(*(*int16)(unsafe.Add(mBase, uint32(v13)+6)))
					v98 = v97
				}
				v101 = int32(4)
				v107 = int32(21)
				v112 = (v17<<(uint(v107)%32) - int32(_a_F_numeric_2)) >> (uint(v107) % 32)
				if int32(0) <= v19 {
					v123 = v18 & int32(_a_F_numeric_7)
				} else {
					v123 = int32(base.Ui32(v18)>>(uint(int32(7))%32)) & int32(63)
				}
				if base.B2i32(int32(base.Ui32(v17-v101)>>(uint(int32(16))%32))-v112 < v98<<(uint(int32(2))%32)+v101)|base.B2i32(v112 < v123)|base.B2i32(v19 < int32(-16384))&base.B2i32(base.Ui32(int32(64)) <= base.Ui32(v112)) == int32(0) {
					v136 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
					v139 = F_palloc(m, int32(base.Ui32(v136)>>(uint(int32(2))%32)))
					mBase = m.M
					v140 = m.ExcPending
					if v140 != 0 {
						return int64(0)
					} else {
						v141 = *(*int32)(unsafe.Add(mBase, uint32(v13)))
						v143 = int32(base.Ui32(v141) >> (uint(int32(2)) % 32))
						if v143 != 0 {
							base.MemoryCopy(m, v139, v13, v143)
						} else {
						}
						if int32(0) < v112 {
							v146 = v112
						} else {
							v146 = int32(0)
						}
						v147 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v13)+4)))
						if v147&int32(_a_F_numeric_0) == int32(_a_F_numeric_8) {
							v156 = v147&int32(_a_F_numeric_9) | v146<<(uint(int32(7))%32)
							*(*uint16)(unsafe.Add(mBase, uint32(v139)+4)) = uint16(v156)
							v200 = v139
						} else {
							v158 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v139)+4)))
							v159 = int32(_a_F_numeric_0)
							v160 = v158 & v159
							if v160 != v159 {
								if v160 != int32(_a_F_numeric_8) {
									v171 = v160
								} else {
									v171 = v158 << (uint(int32(1)) % 32) & int32(_a_F_numeric_10)
								}
							} else {
								v171 = v158 & int32(_a_F_numeric_11)
							}
							v172 = v171 | v146
							*(*uint16)(unsafe.Add(mBase, uint32(v139)+4)) = uint16(v172)
							v200 = v139
						}
						m.G0 = v10 + int32(32)
						return base.I64_extend_i32_u(v200)
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v10)+24)) = int64(0)
					v177 = v10 + int32(8)
					F_set_var_from_num(m, v13, v177)
					mBase = m.M
					v179 = m.ExcPending
					if v179 != 0 {
						return int64(0)
					} else {
						v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
						v181 = F_apply_typmod(m, v177, v17, v180)
						mBase = m.M
						v182 = m.ExcPending
						if v182 != 0 {
							return int64(0)
						} else {
							if v181 == int32(0) {
								v185 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v185)
								v200 = int32(0)
								m.G0 = v10 + int32(32)
								return base.I64_extend_i32_u(v200)
							} else {
								v190 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								v191 = F_make_result_safe(m, v10+int32(8), v190)
								mBase = m.M
								v192 = m.ExcPending
								if v192 != 0 {
									return int64(0)
								} else {
									v193 = *(*int32)(unsafe.Add(mBase, uint32(v10)+24))
									if v193 == int32(0) {
										v200 = v191
										m.G0 = v10 + int32(32)
										return base.I64_extend_i32_u(v200)
									} else {
										F_pfree(m, v193)
										mBase = m.M
										v197 = m.ExcPending
										if v197 != 0 {
											return int64(0)
										} else {
											v200 = v191
											m.G0 = v10 + int32(32)
											return base.I64_extend_i32_u(v200)
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
func F_numeric_cash(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v24 int64
	_ = v24
	var v31 int32
	_ = v31
	var v36 int64
	_ = v36
	var v38 int64
	_ = v38
	var v40 int32
	_ = v40
	var v52 int64
	_ = v52
	var v57 int32
	_ = v57
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v66 int32
	_ = v66
	var v74 int64
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int64
	_ = v87
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v92 int32
	_ = v92
	var v95 int32
	_ = v95
	var v100 int32
	_ = v100
	var v104 int64
	_ = v104
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v9 = F_pg_detoast_datum(m, v8)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int64(0)
	} else {
		v14 = F_PGLC_localeconv(m)
		mBase = m.M
		v15 = m.ExcPending
		if v15 != 0 {
			return int64(0)
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v14)+41)))
			if base.Ui32(int32(10)) < base.Ui32(v16) {
				v19 = int32(2)
			} else {
				v19 = v16
			}
			v20 = base.I32_extend8_s(v19)
			if v20 <= int32(0) {
				v74 = int64(1)
			} else {
				v24 = int64(1)
				if base.Ui32(int32(8)) <= base.Ui32(v19) {
					v31 = int32(0)
					v36 = v24
					for {
						v38 = v36 * int64(100000000)
						v40 = v31 + int32(8)
						if v40 != v20&int32(120) {
							v31 = v40
							v36 = v38
							continue
						} else {
							break
						}
						break
					}
					if v19&int32(7) == int32(0) {
						v74 = v38
					} else {
						v52 = v38
						v57 = int32(0)
						v62 = v52
						for {
							v64 = v62 * int64(10)
							v66 = v57 + int32(1)
							if v66 != v20&int32(7) {
								v57 = v66
								v62 = v64
								continue
							} else {
								break
							}
							break
						}
						v74 = v64
					}
				} else {
					v52 = v24
					v57 = int32(0)
					v62 = v52
					for {
						v64 = v62 * int64(10)
						v66 = v57 + int32(1)
						if v66 != v20&int32(7) {
							v57 = v66
							v62 = v64
							continue
						} else {
							break
						}
						break
					}
					v74 = v64
				}
			}
			v75 = F_int64_to_numeric(m, v74)
			mBase = m.M
			v76 = m.ExcPending
			if v76 != 0 {
				return int64(0)
			} else {
				v77 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
				v78 = F_numeric_mul_safe(m, v9, v75, v77)
				mBase = m.M
				v79 = m.ExcPending
				if v79 != 0 {
					return int64(0)
				} else {
					v80 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
					if v80 == int32(0) {
						v87 = F_numeric_int8_safe(m, v78, v80)
						mBase = m.M
						v88 = m.ExcPending
						if v88 != 0 {
							return int64(0)
						} else {
							v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
							if v89 == int32(0) {
								v104 = v87
							} else {
								v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
								if v92 != int32(453) {
									v104 = v87
								} else {
									v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
									if v95 != int32(1) {
										v104 = v87
									} else {
										v100 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v100)
										v104 = int64(0)
									}
								}
							}
							return v104
						}
					} else {
						v83 = *(*int32)(unsafe.Add(mBase, uint32(v80)))
						if v83 != int32(453) {
							v87 = F_numeric_int8_safe(m, v78, v80)
							mBase = m.M
							v88 = m.ExcPending
							if v88 != 0 {
								return int64(0)
							} else {
								v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
								if v89 == int32(0) {
									v104 = v87
								} else {
									v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
									if v92 != int32(453) {
										v104 = v87
									} else {
										v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
										if v95 != int32(1) {
											v104 = v87
										} else {
											v100 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v100)
											v104 = int64(0)
										}
									}
								}
								return v104
							}
						} else {
							v86 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v80)+4)))
							if v86 != 0 {
								v100 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v100)
								v104 = int64(0)
								return v104
							} else {
								v87 = F_numeric_int8_safe(m, v78, v80)
								mBase = m.M
								v88 = m.ExcPending
								if v88 != 0 {
									return int64(0)
								} else {
									v89 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
									if v89 == int32(0) {
										v104 = v87
									} else {
										v92 = *(*int32)(unsafe.Add(mBase, uint32(v89)))
										if v92 != int32(453) {
											v104 = v87
										} else {
											v95 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v89)+4)))
											if v95 != int32(1) {
												v104 = v87
											} else {
												v100 = int32(1)
												*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v100)
												v104 = int64(0)
											}
										}
									}
									return v104
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_float8_no_overflow(m *base.Module, l0 int32) int64 {
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
	var v25 float64
	_ = v25
	var v28 float64
	_ = v28
	var v29 int32
	_ = v29
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v51 int32
	_ = v51
	var v53 int32
	_ = v53
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v70 int32
	_ = v70
	var v73 int32
	_ = v73
	var v77 int32
	_ = v77
	var v82 float64
	_ = v82
	var v83 int32
	_ = v83
	var v87 float64
	_ = v87
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int64(0)
	} else {
		v16 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
		v17 = base.I32_extend16_s(v16)
		if base.Ui32(int32(_a_F_numeric_float8_no_overflow_0)) <= base.Ui32(v16) {
			if v17 == int32(-4096) {
				v25 = math.Float64frombits(uint64(0xfff0000000000000))
			} else {
				v25 = math.Float64frombits(uint64(0x7ff8000000000000))
			}
			if v17 == int32(-12288) {
				v28 = math.Float64frombits(uint64(0x7ff0000000000000))
			} else {
				v28 = v25
			}
			v87 = v28
			m.G0 = v9 + int32(32)
			return base.I64_reinterpret_f64(v87)
		} else {
			v29 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
			v35 = base.B2i32(int32(0) <= v17)
			if int32(0) <= v17 {
				v36 = int32(-8)
			} else {
				v36 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+8)) = int32(base.Ui32(int32(base.Ui32(v29)>>(uint(int32(2))%32))+v36) >> (uint(int32(1)) % 32))
			if int32(0) <= v17 {
				v41 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
				v51 = v41
			} else {
				v51 = v16<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v16&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+12)) = v51
			v53 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v9)+24)) = v53
			v62 = base.B2i32(v17 < v53)
			if v17 < v53 {
				v63 = int32(base.Ui32(v16)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v63 = v16 & int32(_a_F_numeric_float8_no_overflow_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+20)) = v63
			v70 = v16 & int32(_a_F_numeric_float8_no_overflow_0)
			if v70 == int32(_a_F_numeric_float8_no_overflow_2) {
				v73 = v16 << (uint(int32(1)) % 32) & int32(_a_F_numeric_float8_no_overflow_3)
			} else {
				v73 = v70
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+16)) = v73
			if v17 < v53 {
				v77 = int32(6)
			} else {
				v77 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v9)+28)) = v12 + v77
			v82 = F_numericvar_to_double_no_overflow(m, v9+int32(8))
			mBase = m.M
			v83 = m.ExcPending
			if v83 != 0 {
				return int64(0)
			} else {
				v87 = v82
				m.G0 = v9 + int32(32)
				return base.I64_reinterpret_f64(v87)
			}
		}
	}
}
func F_numeric_int4(m *base.Module, l0 int32) int64 {
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
	var v21 int32
	_ = v21
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = F_numeric_int4_safe(m, v5, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v12 == int32(0) {
				return base.I64_extend_i32_s(v10)
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v15 != int32(453) {
					return base.I64_extend_i32_s(v10)
				} else {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)))
					if v18 != int32(1) {
						return base.I64_extend_i32_s(v10)
					} else {
						v21 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
						return int64(0)
					}
				}
			}
		}
	}
}
func F_numeric_int4_safe(m *base.Module, l0 int32, l1 int32) int32 {
	mBase := m.M
	_ = mBase
	var v3 int32
	_ = v3
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v26 int32
	_ = v26
	var v31 int32
	_ = v31
	var v36 int32
	_ = v36
	var v41 int32
	_ = v41
	var v48 int32
	_ = v48
	var v53 int32
	_ = v53
	var v54 int32
	_ = v54
	var v60 int32
	_ = v60
	var v61 int32
	_ = v61
	var v66 int32
	_ = v66
	var v76 int32
	_ = v76
	var v78 int32
	_ = v78
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v96 int32
	_ = v96
	var v99 int32
	_ = v99
	var v103 int32
	_ = v103
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int64
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v124 int32
	_ = v124
	var v128 int32
	_ = v128
	var v133 int32
	_ = v133
	var v138 int32
	_ = v138
	v3 = int32(0)
	v8 = m.G0
	v10 = v8 + int32(-64)
	m.G0 = v10
	v12 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v13 = base.I32_extend16_s(v12)
	if base.Ui32(int32(_a_F_numeric_int4_safe_0)) <= base.Ui32(v12) {
		v16 = F_errsave_start(m, l1)
		mBase = m.M
		v19 = m.ExcPending
		if v19 != 0 {
			return int32(0)
		} else {
			if v13 == int32(-16384) {
				if v16 == int32(0) {
					v138 = v3
					m.G0 = v10 - int32(-64)
					return v138
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v26 = m.ExcPending
					if v26 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10))) = int32(_a_F_numeric_int4_safe_1)
						F_errmsg(m, int32(_a_F_numeric_int4_safe_2), v10)
						mBase = m.M
						v31 = m.ExcPending
						if v31 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l1, int32(_a_F_numeric_int4_safe_3), int32(_a_F_numeric_int4_safe_4), int32(_a_F_numeric_int4_safe_5))
							mBase = m.M
							v36 = m.ExcPending
							if v36 != 0 {
								return int32(0)
							} else {
								v138 = v3
								m.G0 = v10 - int32(-64)
								return v138
							}
						}
					}
				}
			} else {
				if v16 == int32(0) {
					v138 = v3
					m.G0 = v10 - int32(-64)
					return v138
				} else {
					F_errcode(m, int32(1088))
					mBase = m.M
					v41 = m.ExcPending
					if v41 != 0 {
						return int32(0)
					} else {
						*(*int32)(unsafe.Add(mBase, uint32(v10)+16)) = int32(_a_F_numeric_int4_safe_1)
						F_errmsg(m, int32(_a_F_numeric_int4_safe_6), v8+int32(-48))
						mBase = m.M
						v48 = m.ExcPending
						if v48 != 0 {
							return int32(0)
						} else {
							F_errsave_finish(m, l1, int32(_a_F_numeric_int4_safe_3), int32(_a_F_numeric_int4_safe_7), int32(_a_F_numeric_int4_safe_5))
							mBase = m.M
							v53 = m.ExcPending
							if v53 != 0 {
								return int32(0)
							} else {
								v138 = v3
								m.G0 = v10 - int32(-64)
								return v138
							}
						}
					}
				}
			}
		}
	} else {
		v54 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v60 = base.B2i32(int32(0) <= v13)
		if int32(0) <= v13 {
			v61 = int32(-8)
		} else {
			v61 = int32(-6)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+32)) = int32(base.Ui32(int32(base.Ui32(v54)>>(uint(int32(2))%32))+v61) >> (uint(int32(1)) % 32))
		if int32(0) <= v13 {
			v66 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
			v76 = v66
		} else {
			v76 = v12<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v12&int32(63)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+36)) = v76
		v78 = int32(0)
		*(*int32)(unsafe.Add(mBase, uint32(v10)+48)) = v78
		v88 = base.B2i32(v13 < v78)
		if v13 < v78 {
			v89 = int32(base.Ui32(v12)>>(uint(int32(7))%32)) & int32(63)
		} else {
			v89 = v12 & int32(_a_F_numeric_int4_safe_8)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+44)) = v89
		v96 = v12 & int32(_a_F_numeric_int4_safe_0)
		if v96 == int32(_a_F_numeric_int4_safe_9) {
			v99 = v12 << (uint(int32(1)) % 32) & int32(_a_F_numeric_int4_safe_10)
		} else {
			v99 = v96
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+40)) = v99
		if v13 < v78 {
			v103 = int32(6)
		} else {
			v103 = int32(8)
		}
		*(*int32)(unsafe.Add(mBase, uint32(v10)+52)) = l0 + v103
		v110 = F_numericvar_to_int64(m, v8+int32(-32), v8+int32(-8))
		mBase = m.M
		v111 = m.ExcPending
		if v111 != 0 {
			return int32(0)
		} else {
			if v110 != 0 {
				v112 = *(*int64)(unsafe.Add(mBase, uint32(v10)+56))
				if base.Ui64(int64(-4294967297)) < base.Ui64(v112-int64(2147483648)) {
					v138 = base.I32_wrap_i64(v112)
					m.G0 = v10 - int32(-64)
					return v138
				} else {
					v118 = F_errsave_start(m, l1)
					mBase = m.M
					v119 = m.ExcPending
					if v119 != 0 {
						return int32(0)
					} else {
						if v118 == int32(0) {
							v138 = v78
							m.G0 = v10 - int32(-64)
							return v138
						} else {
							F_errcode(m, int32(50331778))
							mBase = m.M
							v124 = m.ExcPending
							if v124 != 0 {
								return int32(0)
							} else {
								F_errmsg(m, int32(_a_F_numeric_int4_safe_11), int32(0))
								mBase = m.M
								v128 = m.ExcPending
								if v128 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, l1, int32(_a_F_numeric_int4_safe_3), int32(_a_F_numeric_int4_safe_12), int32(_a_F_numeric_int4_safe_5))
									mBase = m.M
									v133 = m.ExcPending
									if v133 != 0 {
										return int32(0)
									} else {
										v138 = v78
										m.G0 = v10 - int32(-64)
										return v138
									}
								}
							}
						}
					}
				}
			} else {
				v118 = F_errsave_start(m, l1)
				mBase = m.M
				v119 = m.ExcPending
				if v119 != 0 {
					return int32(0)
				} else {
					if v118 == int32(0) {
						v138 = v78
						m.G0 = v10 - int32(-64)
						return v138
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v124 = m.ExcPending
						if v124 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_numeric_int4_safe_11), int32(0))
							mBase = m.M
							v128 = m.ExcPending
							if v128 != 0 {
								return int32(0)
							} else {
								F_errsave_finish(m, l1, int32(_a_F_numeric_int4_safe_3), int32(_a_F_numeric_int4_safe_12), int32(_a_F_numeric_int4_safe_5))
								mBase = m.M
								v133 = m.ExcPending
								if v133 != 0 {
									return int32(0)
								} else {
									v138 = v78
									m.G0 = v10 - int32(-64)
									return v138
								}
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_int8(m *base.Module, l0 int32) int64 {
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
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v18 int32
	_ = v18
	var v21 int32
	_ = v21
	var v24 int64
	_ = v24
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
		v10 = F_numeric_int8_safe(m, v5, v9)
		mBase = m.M
		v11 = m.ExcPending
		if v11 != 0 {
			return int64(0)
		} else {
			v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
			if v12 == int32(0) {
				v24 = v10
			} else {
				v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				if v15 != int32(453) {
					v24 = v10
				} else {
					v18 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+4)))
					if v18 != int32(1) {
						v24 = v10
					} else {
						v21 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v21)
						v24 = int64(0)
					}
				}
			}
			return v24
		}
	}
}
func F_numeric_pg_lsn(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v10 int64
	_ = v10
	var v11 int32
	_ = v11
	var v13 int32
	_ = v13
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v27 int32
	_ = v27
	var v30 int32
	_ = v30
	var v39 int32
	_ = v39
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v51 int32
	_ = v51
	var v52 int32
	_ = v52
	var v53 int32
	_ = v53
	var v55 int32
	_ = v55
	var v56 int32
	_ = v56
	var v66 int32
	_ = v66
	var v68 int32
	_ = v68
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v88 int32
	_ = v88
	var v94 int32
	_ = v94
	var v98 int32
	_ = v98
	var v103 int32
	_ = v103
	var v106 int32
	_ = v106
	var v116 int32
	_ = v116
	var v118 int32
	_ = v118
	var v121 int32
	_ = v121
	var v123 int32
	_ = v123
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v128 int32
	_ = v128
	var v141 int32
	_ = v141
	var v151 int32
	_ = v151
	var v155 int32
	_ = v155
	var v157 int32
	_ = v157
	var v162 int32
	_ = v162
	var v170 int32
	_ = v170
	var v173 int32
	_ = v173
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v181 int32
	_ = v181
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v210 int64
	_ = v210
	var v214 int32
	_ = v214
	var v222 int64
	_ = v222
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v232 int64
	_ = v232
	var v235 int64
	_ = v235
	var v238 int64
	_ = v238
	var v241 int64
	_ = v241
	var v242 int64
	_ = v242
	var v246 int64
	_ = v246
	var v253 int64
	_ = v253
	var v264 int64
	_ = v264
	var v267 int64
	_ = v267
	var v272 int64
	_ = v272
	var v273 int64
	_ = v273
	var v275 int64
	_ = v275
	var v277 int32
	_ = v277
	var v290 int32
	_ = v290
	var v294 int32
	_ = v294
	var v297 int32
	_ = v297
	var v301 int32
	_ = v301
	var v306 int32
	_ = v306
	var v307 int32
	_ = v307
	var v319 int32
	_ = v319
	var v324 int32
	_ = v324
	var v334 int64
	_ = v334
	var v336 int32
	_ = v336
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v16 = F_pg_detoast_datum(m, v15)
	mBase = m.M
	v19 = m.ExcPending
	if v19 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v20 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v16)+4)))
	v21 = base.I32_extend16_s(v20)
	if base.Ui32(int32(_a_F_numeric_pg_lsn_0)) <= base.Ui32(v20) {
		goto L5
	} else {
		goto L6
	}
L3:
	;
	F_pfree(m, v71)
	mBase = m.M
	v336 = m.ExcPending
	if v336 != 0 {
		goto L1
	} else {
		goto L77
	}
L4:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13))) = int32(_a_F_numeric_pg_lsn_1)
	F_errmsg(m, int32(_a_F_numeric_pg_lsn_2), v13)
	mBase = m.M
	v319 = m.ExcPending
	if v319 != 0 {
		goto L1
	} else {
		goto L75
	}
L5:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L8
	}
L6:
	;
	goto L7
L7:
	;
	v45 = *(*int32)(unsafe.Add(mBase, uint32(v16)))
	v51 = base.B2i32(int32(0) <= v21)
	if int32(0) <= v21 {
		goto L13
	} else {
		goto L14
	}
L8:
	;
	F_errcode(m, int32(1088))
	mBase = m.M
	v30 = m.ExcPending
	if v30 != 0 {
		goto L1
	} else {
		goto L9
	}
L9:
	;
	if v21 == int32(-16384) {
		goto L4
	} else {
		goto L10
	}
L10:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v13)+16)) = int32(_a_F_numeric_pg_lsn_1)
	F_errmsg(m, int32(_a_F_numeric_pg_lsn_3), v13+int32(16))
	mBase = m.M
	v39 = m.ExcPending
	if v39 != 0 {
		goto L1
	} else {
		goto L11
	}
L11:
	;
	F_errfinish(m, int32(_a_F_numeric_pg_lsn_4), int32(_a_F_numeric_pg_lsn_5), int32(_a_F_numeric_pg_lsn_6))
	mBase = m.M
	v44 = m.ExcPending
	if v44 != 0 {
		goto L1
	} else {
		goto L12
	}
L12:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L13:
	;
	v52 = int32(-8)
	goto L15
L14:
	;
	v52 = int32(-6)
	goto L15
L15:
	;
	v53 = int32(base.Ui32(v45)>>(uint(int32(2))%32)) + v52
	v55 = int32(base.Ui32(v53) >> (uint(int32(1)) % 32))
	if int32(0) <= v21 {
		goto L16
	} else {
		goto L17
	}
L16:
	;
	v56 = int32(*(*int16)(unsafe.Add(mBase, uint32(v16)+6)))
	v66 = v56
	goto L18
L17:
	;
	v66 = v20<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v20&int32(63)
	goto L18
L18:
	;
	v68 = v53 & int32(-2)
	v71 = F_palloc(m, v68+int32(2))
	mBase = m.M
	v72 = m.ExcPending
	if v72 != 0 {
		goto L1
	} else {
		goto L19
	}
L19:
	;
	v73 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v71))) = uint16(v73)
	if base.B2i32(v55 == v73)|base.B2i32(v68 == v73) == v73 {
		goto L20
	} else {
		goto L21
	}
L20:
	;
	if v21 < int32(0) {
		goto L23
	} else {
		goto L24
	}
L21:
	;
	goto L22
L22:
	;
	if v66 < int32(-1) {
		v334 = v10
		goto L3
	} else {
		goto L26
	}
L23:
	;
	v88 = int32(6)
	goto L25
L24:
	;
	v88 = int32(8)
	goto L25
L25:
	;
	base.MemoryCopy(m, v71+int32(2), v16+v88, v68)
	goto L22
L26:
	;
	v94 = v71 + int32(2)
	v98 = (v66 + int32(1)) & int32(1073741823)
	if base.Ui32(v98) < base.Ui32(v55) {
		goto L29
	} else {
		goto L30
	}
L27:
	;
	v170 = v20 & int32(_a_F_numeric_pg_lsn_0)
	if v170 == int32(_a_F_numeric_pg_lsn_7) {
		goto L43
	} else {
		goto L44
	}
L28:
	;
	v151 = int32(1)
	v155 = v98 + v151
	v157 = v71
	v162 = v66 + v151
	goto L27
L29:
	;
	v103 = int32(*(*int16)(unsafe.Add(mBase, uint32(v94+v98<<(uint(int32(1))%32)))))
	if int32(_a_F_numeric_pg_lsn_8) <= v103 {
		goto L32
	} else {
		goto L33
	}
L30:
	;
	v141 = v55
	goto L31
L31:
	;
	if v141 != 0 {
		v155 = v141
		v157 = v94
		v162 = v66
		goto L27
	} else {
		goto L42
	}
L32:
	;
	v106 = v98
	goto L35
L33:
	;
	goto L34
L34:
	;
	v141 = v98
	goto L31
L35:
	;
	v116 = int32(1)
	v118 = v71 + v106<<(uint(v116)%32)
	v121 = int32(*(*int16)(unsafe.Add(mBase, uint32(v118))))
	v123 = base.B2i32(int32(_a_F_numeric_pg_lsn_9) < v121)
	if int32(_a_F_numeric_pg_lsn_9) < v121 {
		goto L37
	} else {
		goto L38
	}
L36:
	;
	if v128 < int32(0) {
		goto L28
	} else {
		goto L41
	}
L37:
	;
	v124 = int32(-9999)
	goto L39
L38:
	;
	v124 = v116
	goto L39
L39:
	;
	v125 = v124 + v121
	*(*uint16)(unsafe.Add(mBase, uint32(v118))) = uint16(v125)
	v128 = v106 - int32(1)
	if int32(_a_F_numeric_pg_lsn_9) < v121 {
		v106 = v128
		goto L35
	} else {
		goto L40
	}
L40:
	;
	goto L36
L41:
	;
	goto L34
L42:
	;
	v334 = v10
	goto L3
L43:
	;
	v173 = v20 << (uint(int32(1)) % 32) & int32(_a_F_numeric_pg_lsn_10)
	goto L45
L44:
	;
	v173 = v170
	goto L45
L45:
	;
	v174 = v155
	v176 = v157
	v181 = v162
	goto L46
L46:
	;
	v184 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176))))
	if v184 != 0 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v334 = v10
	goto L3
L48:
	;
	v185 = v174
	goto L51
L49:
	;
	goto L50
L50:
	;
	v307 = int32(1)
	if v307 < v174 {
		v174 = v174 - v307
		v176 = v176 + int32(2)
		v181 = v181 - v307
		goto L46
	} else {
		goto L74
	}
L51:
	;
	v200 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v176+v185<<(uint(int32(1))%32)-int32(2)))))
	if v200 == int32(0) {
		goto L53
	} else {
		goto L54
	}
L52:
	;
	if v173 == int32(_a_F_numeric_pg_lsn_10) {
		goto L57
	} else {
		goto L58
	}
L53:
	;
	v203 = int32(1)
	if v203 < v185 {
		v185 = v185 - v203
		goto L51
	} else {
		goto L56
	}
L54:
	;
	goto L55
L55:
	;
	goto L52
L56:
	;
	v334 = v10
	goto L3
L57:
	;
	F_pfree(m, v71)
	mBase = m.M
	v290 = m.ExcPending
	if v290 != 0 {
		goto L1
	} else {
		goto L69
	}
L58:
	;
	v210 = int64(*(*int16)(unsafe.Add(mBase, uint32(v176))))
	if v181 <= int32(0) {
		v334 = v210
		goto L3
	} else {
		goto L59
	}
L59:
	;
	v214 = int32(1)
	v222 = v210
	goto L60
L60:
	;
	v224 = v13 + int32(32)
	v225 = int64(0)
	v226 = int64(10000)
	v232 = int64(32)
	v235 = int64(base.Ui64(v222) >> (uint(v232) % 64))
	v238 = int64(4294967295)
	v241 = v222 & v238
	v242 = v226 * v241
	v246 = int64(base.Ui64(v242)>>(uint(v232)%64)) + v226*v235
	v253 = v241*v225 + v246&v238
	*(*int64)(unsafe.Add(mBase, uint32(v224)+8)) = v222*v225 + v225 + v225*v235 + int64(base.Ui64(v246)>>(uint(v232)%64)) + int64(base.Ui64(v253)>>(uint(v232)%64))
	*(*int64)(unsafe.Add(mBase, uint32(v224))) = v242&v238 | v253<<(uint(v232)%64)
	goto L62
L61:
	;
	v334 = v275
	goto L3
L62:
	;
	v264 = *(*int64)(unsafe.Add(mBase, uint32(v13)+40))
	if v264 != int64(0) {
		goto L57
	} else {
		goto L63
	}
L63:
	;
	v267 = *(*int64)(unsafe.Add(mBase, uint32(v13)+32))
	if v214 < v185 {
		goto L64
	} else {
		goto L65
	}
L64:
	;
	v272 = int64(*(*int16)(unsafe.Add(mBase, uint32(v176+v214<<(uint(int32(1))%32)))))
	v273 = v267 + v272
	if base.Ui64(v273) < base.Ui64(v267) {
		goto L57
	} else {
		goto L67
	}
L65:
	;
	v275 = v267
	goto L66
L66:
	;
	v277 = v214 + int32(1)
	if v277 <= v181 {
		v214 = v277
		v222 = v275
		goto L60
	} else {
		goto L68
	}
L67:
	;
	v275 = v273
	goto L66
L68:
	;
	goto L61
L69:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v294 = m.ExcPending
	if v294 != 0 {
		goto L1
	} else {
		goto L70
	}
L70:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v297 = m.ExcPending
	if v297 != 0 {
		goto L1
	} else {
		goto L71
	}
L71:
	;
	F_errmsg(m, int32(_a_F_numeric_pg_lsn_11), int32(0))
	mBase = m.M
	v301 = m.ExcPending
	if v301 != 0 {
		goto L1
	} else {
		goto L72
	}
L72:
	;
	F_errfinish(m, int32(_a_F_numeric_pg_lsn_4), int32(_a_F_numeric_pg_lsn_12), int32(_a_F_numeric_pg_lsn_6))
	mBase = m.M
	v306 = m.ExcPending
	if v306 != 0 {
		goto L1
	} else {
		goto L73
	}
L73:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L74:
	;
	goto L47
L75:
	;
	F_errfinish(m, int32(_a_F_numeric_pg_lsn_4), int32(_a_F_numeric_pg_lsn_13), int32(_a_F_numeric_pg_lsn_6))
	mBase = m.M
	v324 = m.ExcPending
	if v324 != 0 {
		goto L1
	} else {
		goto L76
	}
L76:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L77:
	;
	m.G0 = v13 + int32(48)
	return v334
}
func F_numeric_poly_serialize(m *base.Module, l0 int32) int64 {
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
	var v48 int32
	_ = v48
	var v52 int32
	_ = v52
	var v57 int32
	_ = v57
	var v58 int32
	_ = v58
	var v60 int32
	_ = v60
	var v61 int64
	_ = v61
	var v64 int32
	_ = v64
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v68 int64
	_ = v68
	var v70 int64
	_ = v70
	var v72 int64
	_ = v72
	var v75 int64
	_ = v75
	var v77 int64
	_ = v77
	var v79 int64
	_ = v79
	var v81 int64
	_ = v81
	var v104 int32
	_ = v104
	var v107 int64
	_ = v107
	var v108 int64
	_ = v108
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v115 int64
	_ = v115
	var v117 int64
	_ = v117
	var v119 int64
	_ = v119
	var v122 int64
	_ = v122
	var v124 int64
	_ = v124
	var v126 int64
	_ = v126
	var v128 int64
	_ = v128
	var v151 int32
	_ = v151
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v158 int32
	_ = v158
	var v160 int64
	_ = v160
	var v162 int64
	_ = v162
	var v164 int64
	_ = v164
	var v167 int64
	_ = v167
	var v169 int64
	_ = v169
	var v171 int64
	_ = v171
	var v173 int64
	_ = v173
	var v196 int32
	_ = v196
	var v199 int64
	_ = v199
	var v200 int64
	_ = v200
	var v203 int32
	_ = v203
	var v204 int32
	_ = v204
	var v205 int32
	_ = v205
	var v207 int64
	_ = v207
	var v209 int64
	_ = v209
	var v211 int64
	_ = v211
	var v214 int64
	_ = v214
	var v216 int64
	_ = v216
	var v218 int64
	_ = v218
	var v220 int64
	_ = v220
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v249 int32
	_ = v249
	var v250 int32
	_ = v250
	var v252 int64
	_ = v252
	var v254 int64
	_ = v254
	var v256 int64
	_ = v256
	var v259 int64
	_ = v259
	var v261 int64
	_ = v261
	var v263 int64
	_ = v263
	var v265 int64
	_ = v265
	var v289 int32
	_ = v289
	var v292 int32
	_ = v292
	v6 = m.G0
	v8 = v6 - int32(16)
	m.G0 = v8
	v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v12 == int32(0) {
		v40 = int32(0)
	} else {
		v15 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		switch v15 - int32(435) {
		case 0:
			v40 = int32(1)
		case 1:
			v40 = int32(2)
		default:
			v40 = int32(0)
		}
	}
	if v40 == int32(0) {
		F_errstart_cold(m, int32(21), int32(0))
		mBase = m.M
		v48 = m.ExcPending
		if v48 != 0 {
			return int64(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_poly_serialize_0), int32(0))
			mBase = m.M
			v52 = m.ExcPending
			if v52 != 0 {
				return int64(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_poly_serialize_1), int32(_a_F_numeric_poly_serialize_2), int32(_a_F_numeric_poly_serialize_3))
				mBase = m.M
				v57 = m.ExcPending
				if v57 != 0 {
					return int64(0)
				} else {
					base.Wasm_trap_unreachable()
					for {
					}
				}
			}
		}
	} else {
		v58 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
		F_pq_begintypsend(m, v8)
		mBase = m.M
		v60 = m.ExcPending
		if v60 != 0 {
			return int64(0)
		} else {
			v61 = *(*int64)(unsafe.Add(mBase, uint32(v58)+8))
			F_enlargeStringInfo(m, v8, int32(8))
			mBase = m.M
			v64 = m.ExcPending
			if v64 != 0 {
				return int64(0)
			} else {
				v65 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
				v66 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
				v68 = int64(56)
				v70 = int64(65280)
				v72 = int64(40)
				v75 = int64(16711680)
				v77 = int64(24)
				v79 = int64(4278190080)
				v81 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v65+v66))) = v61<<(uint(v68)%64) | v61&v70<<(uint(v72)%64) | (v61&v75<<(uint(v77)%64) | v61&v79<<(uint(v81)%64)) | (int64(base.Ui64(v61)>>(uint(v81)%64))&v79 | int64(base.Ui64(v61)>>(uint(v77)%64))&v75 | (int64(base.Ui64(v61)>>(uint(v72)%64))&v70 | int64(base.Ui64(v61)>>(uint(v68)%64))))
				v104 = int32(8)
				*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v65 + v104
				v107 = *(*int64)(unsafe.Add(mBase, uint32(v58)+16))
				v108 = *(*int64)(unsafe.Add(mBase, uint32(v58)+24))
				F_enlargeStringInfo(m, v8, v104)
				mBase = m.M
				v111 = m.ExcPending
				if v111 != 0 {
					return int64(0)
				} else {
					v112 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
					v113 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
					v115 = int64(56)
					v117 = int64(65280)
					v119 = int64(40)
					v122 = int64(16711680)
					v124 = int64(24)
					v126 = int64(4278190080)
					v128 = int64(8)
					*(*int64)(unsafe.Add(mBase, uint32(v112+v113))) = v108<<(uint(v115)%64) | v108&v117<<(uint(v119)%64) | (v108&v122<<(uint(v124)%64) | v108&v126<<(uint(v128)%64)) | (int64(base.Ui64(v108)>>(uint(v128)%64))&v126 | int64(base.Ui64(v108)>>(uint(v124)%64))&v122 | (int64(base.Ui64(v108)>>(uint(v119)%64))&v117 | int64(base.Ui64(v108)>>(uint(v115)%64))))
					v151 = int32(8)
					*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v112 + v151
					F_enlargeStringInfo(m, v8, v151)
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int64(0)
					} else {
						v157 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
						v158 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
						v160 = int64(56)
						v162 = int64(65280)
						v164 = int64(40)
						v167 = int64(16711680)
						v169 = int64(24)
						v171 = int64(4278190080)
						v173 = int64(8)
						*(*int64)(unsafe.Add(mBase, uint32(v157+v158))) = v107<<(uint(v160)%64) | v107&v162<<(uint(v164)%64) | (v107&v167<<(uint(v169)%64) | v107&v171<<(uint(v173)%64)) | (int64(base.Ui64(v107)>>(uint(v173)%64))&v171 | int64(base.Ui64(v107)>>(uint(v169)%64))&v167 | (int64(base.Ui64(v107)>>(uint(v164)%64))&v162 | int64(base.Ui64(v107)>>(uint(v160)%64))))
						v196 = int32(8)
						*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v157 + v196
						v199 = *(*int64)(unsafe.Add(mBase, uint32(v58)+32))
						v200 = *(*int64)(unsafe.Add(mBase, uint32(v58)+40))
						F_enlargeStringInfo(m, v8, v196)
						mBase = m.M
						v203 = m.ExcPending
						if v203 != 0 {
							return int64(0)
						} else {
							v204 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
							v205 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
							v207 = int64(56)
							v209 = int64(65280)
							v211 = int64(40)
							v214 = int64(16711680)
							v216 = int64(24)
							v218 = int64(4278190080)
							v220 = int64(8)
							*(*int64)(unsafe.Add(mBase, uint32(v204+v205))) = v200<<(uint(v207)%64) | v200&v209<<(uint(v211)%64) | (v200&v214<<(uint(v216)%64) | v200&v218<<(uint(v220)%64)) | (int64(base.Ui64(v200)>>(uint(v220)%64))&v218 | int64(base.Ui64(v200)>>(uint(v216)%64))&v214 | (int64(base.Ui64(v200)>>(uint(v211)%64))&v209 | int64(base.Ui64(v200)>>(uint(v207)%64))))
							v243 = int32(8)
							*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v204 + v243
							F_enlargeStringInfo(m, v8, v243)
							mBase = m.M
							v248 = m.ExcPending
							if v248 != 0 {
								return int64(0)
							} else {
								v249 = *(*int32)(unsafe.Add(mBase, uint32(v8)+4))
								v250 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
								v252 = int64(56)
								v254 = int64(65280)
								v256 = int64(40)
								v259 = int64(16711680)
								v261 = int64(24)
								v263 = int64(4278190080)
								v265 = int64(8)
								*(*int64)(unsafe.Add(mBase, uint32(v249+v250))) = v199<<(uint(v252)%64) | v199&v254<<(uint(v256)%64) | (v199&v259<<(uint(v261)%64) | v199&v263<<(uint(v265)%64)) | (int64(base.Ui64(v199)>>(uint(v265)%64))&v263 | int64(base.Ui64(v199)>>(uint(v261)%64))&v259 | (int64(base.Ui64(v199)>>(uint(v256)%64))&v254 | int64(base.Ui64(v199)>>(uint(v252)%64))))
								v289 = v249 + int32(8)
								*(*int32)(unsafe.Add(mBase, uint32(v8)+4)) = v289
								v292 = *(*int32)(unsafe.Add(mBase, uint32(v8)))
								*(*int32)(unsafe.Add(mBase, uint32(v292))) = v289 << (uint(int32(2)) % 32)
								m.G0 = v8 + int32(16)
								return base.I64_extend_i32_u(v292)
							}
						}
					}
				}
			}
		}
	}
}
func F_numeric_power(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v16 int32
	_ = v16
	var v18 int32
	_ = v18
	var v20 int32
	_ = v20
	var v21 int32
	_ = v21
	var v24 int32
	_ = v24
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v34 int32
	_ = v34
	var v37 int32
	_ = v37
	var v45 int32
	_ = v45
	var v49 int32
	_ = v49
	var v56 int32
	_ = v56
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v92 int32
	_ = v92
	var v100 int32
	_ = v100
	var v106 int32
	_ = v106
	var v107 int32
	_ = v107
	var v110 int32
	_ = v110
	var v111 int32
	_ = v111
	var v112 int32
	_ = v112
	var v118 int32
	_ = v118
	var v119 int32
	_ = v119
	var v120 int32
	_ = v120
	var v124 int32
	_ = v124
	var v125 int32
	_ = v125
	var v135 int32
	_ = v135
	var v143 int32
	_ = v143
	var v146 int32
	_ = v146
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v182 int32
	_ = v182
	var v187 int32
	_ = v187
	var v191 int32
	_ = v191
	var v204 int32
	_ = v204
	var v206 int32
	_ = v206
	var v211 int32
	_ = v211
	var v212 int32
	_ = v212
	var v213 int32
	_ = v213
	var v215 int32
	_ = v215
	var v223 int32
	_ = v223
	var v225 int32
	_ = v225
	var v234 int32
	_ = v234
	var v235 int32
	_ = v235
	var v240 int32
	_ = v240
	var v249 int32
	_ = v249
	var v251 int32
	_ = v251
	var v258 int32
	_ = v258
	var v265 int32
	_ = v265
	var v266 int32
	_ = v266
	var v269 int32
	_ = v269
	var v276 int32
	_ = v276
	var v281 int32
	_ = v281
	var v289 int32
	_ = v289
	var v293 int32
	_ = v293
	var v298 int32
	_ = v298
	var v302 int32
	_ = v302
	var v308 int32
	_ = v308
	var v319 int32
	_ = v319
	var v329 int32
	_ = v329
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v343 int32
	_ = v343
	var v344 int32
	_ = v344
	var v352 int32
	_ = v352
	var v355 int32
	_ = v355
	var v362 int32
	_ = v362
	var v364 int32
	_ = v364
	var v370 int32
	_ = v370
	var v381 int32
	_ = v381
	var v384 int32
	_ = v384
	var v387 int32
	_ = v387
	var v388 int32
	_ = v388
	var v390 int32
	_ = v390
	var v391 int32
	_ = v391
	var v395 int32
	_ = v395
	var v399 int32
	_ = v399
	var v400 int32
	_ = v400
	var v403 int32
	_ = v403
	var v408 int32
	_ = v408
	var v411 int32
	_ = v411
	var v415 int32
	_ = v415
	var v420 int32
	_ = v420
	var v423 int32
	_ = v423
	var v424 int32
	_ = v424
	var v425 int32
	_ = v425
	var v426 int32
	_ = v426
	var v431 int32
	_ = v431
	var v432 int32
	_ = v432
	var v435 int32
	_ = v435
	var v436 int32
	_ = v436
	var v446 int32
	_ = v446
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v456 int32
	_ = v456
	var v457 int32
	_ = v457
	var v460 int32
	_ = v460
	var v461 int32
	_ = v461
	var v462 int32
	_ = v462
	var v464 int32
	_ = v464
	var v465 int32
	_ = v465
	var v471 int32
	_ = v471
	var v472 int32
	_ = v472
	var v475 int32
	_ = v475
	var v477 int32
	_ = v477
	var v487 int32
	_ = v487
	var v495 int32
	_ = v495
	var v498 int32
	_ = v498
	var v503 int32
	_ = v503
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v507 int32
	_ = v507
	var v520 int32
	_ = v520
	var v524 int32
	_ = v524
	var v530 int32
	_ = v530
	var v531 int32
	_ = v531
	var v532 int32
	_ = v532
	var v534 int32
	_ = v534
	var v539 int32
	_ = v539
	var v543 int32
	_ = v543
	var v556 int32
	_ = v556
	var v558 int32
	_ = v558
	var v563 int32
	_ = v563
	var v564 int32
	_ = v564
	var v565 int32
	_ = v565
	var v567 int32
	_ = v567
	var v575 int32
	_ = v575
	var v577 int32
	_ = v577
	var v586 int32
	_ = v586
	var v587 int32
	_ = v587
	var v592 int32
	_ = v592
	var v601 int32
	_ = v601
	var v603 int32
	_ = v603
	var v610 int32
	_ = v610
	var v617 int32
	_ = v617
	var v618 int32
	_ = v618
	var v621 int32
	_ = v621
	var v628 int32
	_ = v628
	var v633 int32
	_ = v633
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v650 int32
	_ = v650
	var v654 int32
	_ = v654
	var v660 int32
	_ = v660
	var v671 int32
	_ = v671
	var v681 int32
	_ = v681
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v702 int32
	_ = v702
	var v703 int32
	_ = v703
	var v705 int32
	_ = v705
	var v709 int32
	_ = v709
	var v710 int32
	_ = v710
	var v718 int32
	_ = v718
	var v719 int32
	_ = v719
	var v726 int32
	_ = v726
	var v729 int32
	_ = v729
	var v734 int32
	_ = v734
	var v737 int32
	_ = v737
	var v742 int32
	_ = v742
	var v746 int32
	_ = v746
	var v756 int32
	_ = v756
	var v758 int32
	_ = v758
	var v760 int32
	_ = v760
	var v761 int32
	_ = v761
	var v762 int32
	_ = v762
	var v773 int32
	_ = v773
	var v775 int32
	_ = v775
	var v784 int32
	_ = v784
	var v786 int32
	_ = v786
	var v787 int32
	_ = v787
	var v793 int32
	_ = v793
	var v796 int32
	_ = v796
	var v803 int32
	_ = v803
	var v804 int32
	_ = v804
	var v805 int32
	_ = v805
	var v815 int32
	_ = v815
	var v816 int32
	_ = v816
	var v821 int32
	_ = v821
	var v822 int32
	_ = v822
	var v823 int32
	_ = v823
	var v824 int32
	_ = v824
	var v825 int32
	_ = v825
	var v831 int32
	_ = v831
	var v835 int32
	_ = v835
	var v841 int32
	_ = v841
	var v846 int32
	_ = v846
	var v847 int32
	_ = v847
	var v848 int32
	_ = v848
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v861 int32
	_ = v861
	var v871 int32
	_ = v871
	var v872 int32
	_ = v872
	var v877 int32
	_ = v877
	var v878 int32
	_ = v878
	var v879 int32
	_ = v879
	var v880 int32
	_ = v880
	var v881 int32
	_ = v881
	var v887 int32
	_ = v887
	var v891 int32
	_ = v891
	var v897 int32
	_ = v897
	var v903 int32
	_ = v903
	var v908 int32
	_ = v908
	var v912 int32
	_ = v912
	var v913 int32
	_ = v913
	var v916 int32
	_ = v916
	var v917 int32
	_ = v917
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v931 int32
	_ = v931
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v936 int32
	_ = v936
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v942 int32
	_ = v942
	var v945 int32
	_ = v945
	var v950 int32
	_ = v950
	var v953 int32
	_ = v953
	var v958 int32
	_ = v958
	var v962 int32
	_ = v962
	var v972 int32
	_ = v972
	var v974 int32
	_ = v974
	var v976 int32
	_ = v976
	var v977 int32
	_ = v977
	var v978 int32
	_ = v978
	var v989 int32
	_ = v989
	var v991 int32
	_ = v991
	var v1000 int32
	_ = v1000
	var v1002 int32
	_ = v1002
	var v1003 int32
	_ = v1003
	var v1009 int32
	_ = v1009
	var v1012 int32
	_ = v1012
	var v1015 int32
	_ = v1015
	var v1019 int32
	_ = v1019
	var v1020 int32
	_ = v1020
	var v1025 int32
	_ = v1025
	var v1032 int32
	_ = v1032
	var v1033 int32
	_ = v1033
	var v1036 int32
	_ = v1036
	var v1037 int32
	_ = v1037
	var v1039 int64
	_ = v1039
	var v1045 int32
	_ = v1045
	var v1050 int32
	_ = v1050
	var v1051 int32
	_ = v1051
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1055 int32
	_ = v1055
	var v1057 int32
	_ = v1057
	var v1059 int32
	_ = v1059
	var v1069 int32
	_ = v1069
	var v1071 int32
	_ = v1071
	var v1072 int32
	_ = v1072
	var v1083 int32
	_ = v1083
	var v1085 int32
	_ = v1085
	var v1094 int32
	_ = v1094
	var v1095 int32
	_ = v1095
	var v1099 int32
	_ = v1099
	var v1100 int32
	_ = v1100
	var v1102 int32
	_ = v1102
	var v1107 int32
	_ = v1107
	var v1108 int32
	_ = v1108
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1114 int32
	_ = v1114
	var v1116 int32
	_ = v1116
	var v1126 int32
	_ = v1126
	var v1128 int32
	_ = v1128
	var v1129 int32
	_ = v1129
	var v1140 int32
	_ = v1140
	var v1142 int32
	_ = v1142
	var v1151 int32
	_ = v1151
	var v1152 int32
	_ = v1152
	var v1156 int32
	_ = v1156
	var v1157 int32
	_ = v1157
	var v1161 int32
	_ = v1161
	var v1163 int32
	_ = v1163
	var v1168 int32
	_ = v1168
	var v1169 int32
	_ = v1169
	var v1172 int64
	_ = v1172
	var v1177 int32
	_ = v1177
	var v1178 int32
	_ = v1178
	var v1179 float64
	_ = v1179
	var v1184 int32
	_ = v1184
	var v1186 int32
	_ = v1186
	var v1189 int32
	_ = v1189
	var v1192 int32
	_ = v1192
	var v1194 float64
	_ = v1194
	var v1199 int32
	_ = v1199
	var v1201 float64
	_ = v1201
	var v1206 int32
	_ = v1206
	var v1210 float64
	_ = v1210
	var v1218 float64
	_ = v1218
	var v1219 int32
	_ = v1219
	var v1230 int64
	_ = v1230
	var v1245 int32
	_ = v1245
	var v1247 int64
	_ = v1247
	var v1257 int64
	_ = v1257
	var v1261 int64
	_ = v1261
	var v1263 int32
	_ = v1263
	var v1264 int32
	_ = v1264
	var v1266 int32
	_ = v1266
	var v1270 float64
	_ = v1270
	var v1272 float64
	_ = v1272
	var v1285 float64
	_ = v1285
	var v1288 float64
	_ = v1288
	var v1293 float64
	_ = v1293
	var v1294 float64
	_ = v1294
	var v1295 float64
	_ = v1295
	var v1296 float64
	_ = v1296
	var v1301 float64
	_ = v1301
	var v1302 float64
	_ = v1302
	var v1303 float64
	_ = v1303
	var v1328 float64
	_ = v1328
	var v1340 float64
	_ = v1340
	var v1362 float64
	_ = v1362
	var v1366 float64
	_ = v1366
	var v1371 float64
	_ = v1371
	var v1376 int32
	_ = v1376
	var v1378 int32
	_ = v1378
	var v1381 int64
	_ = v1381
	var v1387 int32
	_ = v1387
	var v1388 int32
	_ = v1388
	var v1390 int32
	_ = v1390
	var v1392 int32
	_ = v1392
	var v1395 int32
	_ = v1395
	var v1399 int32
	_ = v1399
	var v1400 int32
	_ = v1400
	var v1403 int32
	_ = v1403
	var v1405 int32
	_ = v1405
	var v1412 int32
	_ = v1412
	var v1415 int64
	_ = v1415
	var v1418 int32
	_ = v1418
	var v1421 int32
	_ = v1421
	var v1422 int32
	_ = v1422
	var v1423 int32
	_ = v1423
	var v1435 int32
	_ = v1435
	var v1437 int32
	_ = v1437
	var v1438 int64
	_ = v1438
	var v1440 int64
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1464 int32
	_ = v1464
	var v1466 int32
	_ = v1466
	var v1470 int32
	_ = v1470
	var v1471 int32
	_ = v1471
	var v1482 int32
	_ = v1482
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1489 int32
	_ = v1489
	var v1490 int32
	_ = v1490
	var v1491 int32
	_ = v1491
	var v1493 int32
	_ = v1493
	var v1494 int32
	_ = v1494
	var v1495 int32
	_ = v1495
	var v1498 int32
	_ = v1498
	var v1501 int32
	_ = v1501
	var v1506 int32
	_ = v1506
	var v1510 int32
	_ = v1510
	var v1516 int32
	_ = v1516
	var v1522 int32
	_ = v1522
	var v1523 int32
	_ = v1523
	var v1526 int32
	_ = v1526
	var v1529 int32
	_ = v1529
	var v1531 int32
	_ = v1531
	var v1532 int32
	_ = v1532
	var v1533 int32
	_ = v1533
	var v1536 int32
	_ = v1536
	var v1544 int32
	_ = v1544
	var v1548 int32
	_ = v1548
	var v1549 int32
	_ = v1549
	var v1552 int32
	_ = v1552
	var v1573 int32
	_ = v1573
	var v1576 int32
	_ = v1576
	var v1578 int32
	_ = v1578
	var v1580 int32
	_ = v1580
	var v1585 int32
	_ = v1585
	var v1587 int32
	_ = v1587
	var v1591 int64
	_ = v1591
	var v1596 int32
	_ = v1596
	var v1598 int32
	_ = v1598
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1601 int32
	_ = v1601
	var v1604 int32
	_ = v1604
	var v1606 int32
	_ = v1606
	var v1612 int64
	_ = v1612
	var v1614 int64
	_ = v1614
	var v1618 float64
	_ = v1618
	var v1623 int32
	_ = v1623
	var v1625 int32
	_ = v1625
	var v1628 int32
	_ = v1628
	var v1629 int32
	_ = v1629
	var v1630 int32
	_ = v1630
	var v1633 int32
	_ = v1633
	var v1637 int32
	_ = v1637
	var v1639 int32
	_ = v1639
	var v1640 int64
	_ = v1640
	var v1642 int64
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1649 int32
	_ = v1649
	var v1651 int32
	_ = v1651
	var v1655 int64
	_ = v1655
	var v1658 int64
	_ = v1658
	var v1660 int32
	_ = v1660
	var v1661 int32
	_ = v1661
	var v1664 int32
	_ = v1664
	var v1670 int32
	_ = v1670
	var v1683 int32
	_ = v1683
	var v1687 int32
	_ = v1687
	var v1688 int32
	_ = v1688
	var v1691 int32
	_ = v1691
	var v1692 int32
	_ = v1692
	var v1694 int32
	_ = v1694
	var v1696 int32
	_ = v1696
	var v1697 int32
	_ = v1697
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1703 int32
	_ = v1703
	var v1706 int32
	_ = v1706
	var v1710 int32
	_ = v1710
	var v1711 int32
	_ = v1711
	var v1712 int32
	_ = v1712
	var v1713 int32
	_ = v1713
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1719 int32
	_ = v1719
	var v1721 int32
	_ = v1721
	var v1726 int32
	_ = v1726
	var v1732 int32
	_ = v1732
	var v1734 int32
	_ = v1734
	var v1735 int32
	_ = v1735
	var v1737 int64
	_ = v1737
	var v1741 int32
	_ = v1741
	var v1745 int32
	_ = v1745
	var v1746 int32
	_ = v1746
	var v1748 int32
	_ = v1748
	var v1755 int32
	_ = v1755
	var v1760 int32
	_ = v1760
	var v1761 int32
	_ = v1761
	var v1762 int32
	_ = v1762
	var v1764 int32
	_ = v1764
	var v1766 int32
	_ = v1766
	var v1770 int32
	_ = v1770
	var v1773 int64
	_ = v1773
	var v1779 int64
	_ = v1779
	var v1799 int32
	_ = v1799
	var v1807 int32
	_ = v1807
	var v1813 int32
	_ = v1813
	var v1817 int32
	_ = v1817
	var v1819 int32
	_ = v1819
	var v1822 int32
	_ = v1822
	var v1823 int32
	_ = v1823
	var v1824 int32
	_ = v1824
	var v1827 int32
	_ = v1827
	var v1829 int64
	_ = v1829
	var v1833 int64
	_ = v1833
	var v1839 int32
	_ = v1839
	var v1842 int32
	_ = v1842
	var v1843 int32
	_ = v1843
	var v1845 int32
	_ = v1845
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1850 int32
	_ = v1850
	var v1853 int32
	_ = v1853
	var v1855 int32
	_ = v1855
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1862 float64
	_ = v1862
	var v1863 int32
	_ = v1863
	var v1869 int32
	_ = v1869
	var v1871 int32
	_ = v1871
	var v1874 int64
	_ = v1874
	var v1879 int32
	_ = v1879
	var v1884 int32
	_ = v1884
	var v1885 int32
	_ = v1885
	var v1886 int32
	_ = v1886
	var v1888 int32
	_ = v1888
	var v1890 int32
	_ = v1890
	var v1893 int32
	_ = v1893
	var v1894 int32
	_ = v1894
	var v1895 int32
	_ = v1895
	var v1898 int32
	_ = v1898
	var v1901 int32
	_ = v1901
	var v1902 int32
	_ = v1902
	var v1905 int32
	_ = v1905
	var v1907 int32
	_ = v1907
	var v1911 int32
	_ = v1911
	var v1913 int32
	_ = v1913
	var v1915 int32
	_ = v1915
	var v1916 int32
	_ = v1916
	var v1922 int32
	_ = v1922
	var v1924 int32
	_ = v1924
	var v1925 int32
	_ = v1925
	var v1927 int32
	_ = v1927
	var v1931 int32
	_ = v1931
	var v1935 int32
	_ = v1935
	var v1938 int32
	_ = v1938
	var v1942 int32
	_ = v1942
	var v1947 int32
	_ = v1947
	var v1951 int32
	_ = v1951
	var v1954 int32
	_ = v1954
	var v1958 int32
	_ = v1958
	var v1963 int32
	_ = v1963
	var v1967 int32
	_ = v1967
	var v1970 int32
	_ = v1970
	var v1974 int32
	_ = v1974
	var v1979 int32
	_ = v1979
	var v1983 int32
	_ = v1983
	var v1986 int32
	_ = v1986
	var v1990 int32
	_ = v1990
	var v1995 int32
	_ = v1995
	var v1999 int32
	_ = v1999
	var v2002 int32
	_ = v2002
	var v2006 int32
	_ = v2006
	var v2011 int32
	_ = v2011
	var v2015 int32
	_ = v2015
	var v2018 int32
	_ = v2018
	var v2022 int32
	_ = v2022
	var v2027 int32
	_ = v2027
	var v2039 int32
	_ = v2039
	var v2042 int32
	_ = v2042
	var v2051 int32
	_ = v2051
	var v2053 int32
	_ = v2053
	var v2057 int32
	_ = v2057
	var v2058 int32
	_ = v2058
	var v2069 int32
	_ = v2069
	var v2072 int32
	_ = v2072
	var v2073 int32
	_ = v2073
	var v2076 int32
	_ = v2076
	var v2077 int32
	_ = v2077
	var v2078 int32
	_ = v2078
	var v2080 int32
	_ = v2080
	var v2081 int32
	_ = v2081
	var v2082 int32
	_ = v2082
	var v2085 int32
	_ = v2085
	var v2088 int32
	_ = v2088
	var v2093 int32
	_ = v2093
	var v2097 int32
	_ = v2097
	var v2103 int32
	_ = v2103
	var v2109 int32
	_ = v2109
	var v2110 int32
	_ = v2110
	var v2113 int32
	_ = v2113
	var v2116 int32
	_ = v2116
	var v2118 int32
	_ = v2118
	var v2119 int32
	_ = v2119
	var v2120 int32
	_ = v2120
	var v2123 int32
	_ = v2123
	var v2131 int32
	_ = v2131
	var v2135 int32
	_ = v2135
	var v2136 int32
	_ = v2136
	var v2139 int32
	_ = v2139
	var v2173 int32
	_ = v2173
	var v2174 int32
	_ = v2174
	var v2175 int32
	_ = v2175
	var v2179 int32
	_ = v2179
	var v2183 int32
	_ = v2183
	v2 = int32(0)
	v16 = m.G0
	v18 = v16 - int32(144)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v21 = F_pg_detoast_datum(m, v20)
	mBase = m.M
	v24 = m.ExcPending
	if v24 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v26 = F_pg_detoast_datum(m, v25)
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v28 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v29 = base.I32_extend16_s(v28)
	if base.Ui32(v28) <= base.Ui32(int32(_a_F_numeric_power_0)) {
		goto L33
	} else {
		goto L34
	}
L4:
	;
	m.G0 = v18 + int32(144)
	return base.I64_extend_i32_u(v2183)
L5:
	;
	v2173 = F_make_result_safe(m, v18, int32(0))
	mBase = m.M
	v2174 = m.ExcPending
	if v2174 != 0 {
		goto L1
	} else {
		goto L661
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v1395
	v2039 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v2042 = v1395 + v2039<<(uint(int32(2))%32)
	if v2042+int32(4) < int32(0) {
		goto L636
	} else {
		goto L637
	}
L7:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2015 = m.ExcPending
	if v2015 != 0 {
		goto L1
	} else {
		goto L631
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1999 = m.ExcPending
	if v1999 != 0 {
		goto L1
	} else {
		goto L627
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1983 = m.ExcPending
	if v1983 != 0 {
		goto L1
	} else {
		goto L623
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1967 = m.ExcPending
	if v1967 != 0 {
		goto L1
	} else {
		goto L619
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1951 = m.ExcPending
	if v1951 != 0 {
		goto L1
	} else {
		goto L615
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1935 = m.ExcPending
	if v1935 != 0 {
		goto L1
	} else {
		goto L611
	}
L13:
	;
	v1039 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v1039
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v1039
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1039
	v1045 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v1050 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v1051 = base.I32_extend16_s(v1050)
	v1053 = base.B2i32(int32(0) <= v1051)
	if int32(0) <= v1051 {
		goto L343
	} else {
		goto L344
	}
L14:
	;
	if v29 == int32(-12288) {
		goto L306
	} else {
		goto L307
	}
L15:
	;
	v922 = v719
	v923 = base.B2i32(v87 != int32(_a_F_numeric_power_1))
	goto L14
L16:
	;
	if v908 == v903 {
		goto L301
	} else {
		goto L302
	}
L17:
	;
	v726 = v18 + int32(48)
	v729 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v734 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	if int32(0) <= v734 {
		goto L229
	} else {
		goto L230
	}
L18:
	;
	v718 = int32(_a_F_numeric_power_1)
	v719 = base.B2i32(v87 == v718)
	if v87&int32(_a_F_numeric_power_2) != v718 {
		goto L15
	} else {
		goto L227
	}
L19:
	;
	v705 = base.B2i32(int32(0) < v691)
	if v705&v687 != 0 {
		goto L221
	} else {
		goto L222
	}
L20:
	;
	v702 = F_make_result_safe(m, int32(_a_F_numeric_power_3), int32(0))
	mBase = m.M
	v703 = m.ExcPending
	if v703 != 0 {
		goto L1
	} else {
		goto L220
	}
L21:
	;
	if v691 != 0 {
		goto L19
	} else {
		goto L219
	}
L22:
	;
	v465 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v471 = base.B2i32(int32(0) <= v29)
	if int32(0) <= v29 {
		goto L160
	} else {
		goto L161
	}
L23:
	;
	v457 = int32(0)
	if base.Ui32(int32(-16385)) < base.Ui32(v29) {
		v686 = v452
		v687 = v457
		v689 = v454
		v691 = v456
		goto L21
	} else {
		goto L159
	}
L24:
	;
	v431 = base.B2i32(int32(0) <= base.I32_extend16_s(v423))
	if int32(0) <= base.I32_extend16_s(v423) {
		goto L151
	} else {
		goto L152
	}
L25:
	;
	if v29 == int32(-12288) {
		goto L20
	} else {
		goto L150
	}
L26:
	;
	v399 = base.B2i32(v34 == int32(_a_F_numeric_power_1))
	if v34 == int32(_a_F_numeric_power_1) {
		goto L140
	} else {
		goto L141
	}
L27:
	;
	if v390 != 0 {
		goto L136
	} else {
		goto L137
	}
L28:
	;
	if base.Ui32(int32(_a_F_numeric_power_0)) < base.Ui32(v87) {
		goto L18
	} else {
		goto L128
	}
L29:
	;
	if base.Ui32(v120) < base.Ui32(int32(2)) {
		goto L26
	} else {
		goto L124
	}
L30:
	;
	v343 = F_make_result_safe(m, int32(_a_F_numeric_power_4), int32(0))
	mBase = m.M
	v344 = m.ExcPending
	if v344 != 0 {
		goto L1
	} else {
		goto L123
	}
L31:
	;
	if v87 != int32(_a_F_numeric_power_5) {
		goto L28
	} else {
		goto L122
	}
L32:
	;
	v112 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v118 = base.B2i32(int32(0) <= v29)
	if int32(0) <= v29 {
		goto L62
	} else {
		goto L63
	}
L33:
	;
	v34 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	if base.Ui32(int32(_a_F_numeric_power_0)) < base.Ui32(v34) {
		goto L32
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v87 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	if v29 != int32(-16384) {
		goto L31
	} else {
		goto L53
	}
L36:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if int32(0) <= base.I32_extend16_s(v34) {
		goto L37
	} else {
		goto L38
	}
L37:
	;
	v45 = int32(-8)
	goto L39
L38:
	;
	v45 = int32(-6)
	goto L39
L39:
	;
	if base.Ui32(int32(base.Ui32(v37)>>(uint(int32(2))%32))+v45) < base.Ui32(int32(2)) {
		goto L13
	} else {
		goto L40
	}
L40:
	;
	v49 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if int32(0) <= v29 {
		goto L41
	} else {
		goto L42
	}
L41:
	;
	v56 = int32(-8)
	goto L43
L42:
	;
	v56 = int32(-6)
	goto L43
L43:
	;
	if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v49)>>(uint(int32(2))%32))+v56) {
		goto L13
	} else {
		goto L44
	}
L44:
	;
	v65 = v34 & int32(_a_F_numeric_power_5)
	if v65 == int32(_a_F_numeric_power_6) {
		goto L45
	} else {
		goto L46
	}
L45:
	;
	v68 = v34 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_7)
	goto L47
L46:
	;
	v68 = v65
	goto L47
L47:
	;
	if v68 != int32(_a_F_numeric_power_7) {
		goto L13
	} else {
		goto L48
	}
L48:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v74 = m.ExcPending
	if v74 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v77 = m.ExcPending
	if v77 != 0 {
		goto L1
	} else {
		goto L50
	}
L50:
	;
	F_errmsg(m, int32(_a_F_numeric_power_8), int32(0))
	mBase = m.M
	v81 = m.ExcPending
	if v81 != 0 {
		goto L1
	} else {
		goto L51
	}
L51:
	;
	F_errfinish(m, int32(_a_F_numeric_power_9), int32(4081), int32(_a_F_numeric_power_10))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L1
	} else {
		goto L52
	}
L52:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L53:
	;
	if base.Ui32(int32(_a_F_numeric_power_0)) < base.Ui32(v87) {
		goto L54
	} else {
		goto L55
	}
L54:
	;
	v110 = F_make_result_safe(m, int32(_a_F_numeric_power_4), int32(0))
	mBase = m.M
	v111 = m.ExcPending
	if v111 != 0 {
		goto L1
	} else {
		goto L61
	}
L55:
	;
	v92 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if int32(0) <= base.I32_extend16_s(v87) {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v100 = int32(-8)
	goto L58
L57:
	;
	v100 = int32(-6)
	goto L58
L58:
	;
	if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v92)>>(uint(int32(2))%32))+v100) {
		goto L54
	} else {
		goto L59
	}
L59:
	;
	v106 = F_make_result_safe(m, int32(_a_F_numeric_power_3), int32(0))
	mBase = m.M
	v107 = m.ExcPending
	if v107 != 0 {
		goto L1
	} else {
		goto L60
	}
L60:
	;
	v2183 = v106
	goto L4
L61:
	;
	v2183 = v110
	goto L4
L62:
	;
	v119 = int32(-8)
	goto L64
L63:
	;
	v119 = int32(-6)
	goto L64
L64:
	;
	v120 = int32(base.Ui32(v112)>>(uint(int32(2))%32)) + v119
	if v34 != int32(_a_F_numeric_power_5) {
		goto L29
	} else {
		goto L65
	}
L65:
	;
	v124 = int32(base.Ui32(v120) >> (uint(int32(1)) % 32))
	if int32(0) <= v29 {
		goto L66
	} else {
		goto L67
	}
L66:
	;
	v125 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
	v135 = v125
	goto L68
L67:
	;
	v135 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
	goto L68
L68:
	;
	if v124 == int32(0) {
		goto L30
	} else {
		goto L69
	}
L69:
	;
	v143 = v28 & int32(_a_F_numeric_power_5)
	if v143 == int32(_a_F_numeric_power_6) {
		goto L70
	} else {
		goto L71
	}
L70:
	;
	v146 = v28 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_7)
	goto L72
L71:
	;
	v146 = v143
	goto L72
L72:
	;
	if v146 != 0 {
		goto L30
	} else {
		goto L73
	}
L73:
	;
	if v29 < int32(0) {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v151 = int32(6)
	goto L76
L75:
	;
	v151 = int32(8)
	goto L76
L76:
	;
	v152 = v21 + v151
	v154 = int32(1)
	v155 = int32(0)
	if base.B2i32(v155 < v135)&base.B2i32(v155 < v124) == v155 {
		v187 = v135
		v191 = v155
		goto L79
	} else {
		goto L80
	}
L77:
	;
	if v329 != 0 {
		goto L30
	} else {
		goto L120
	}
L78:
	;
	v329 = v319
	goto L77
L79:
	;
	if int32(0)|base.B2i32(v155 <= v187) != 0 {
		v223 = v155
		v225 = v155
		goto L86
	} else {
		goto L87
	}
L80:
	;
	v168 = v135
	v172 = v155
	goto L81
L81:
	;
	v178 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152+v172<<(uint(int32(1))%32)))))
	if v178 != 0 {
		v319 = int32(1)
		goto L78
	} else {
		goto L83
	}
L82:
	;
	v187 = v182
	v191 = v180
	goto L79
L83:
	;
	v179 = int32(1)
	v180 = v172 + v179
	v182 = v168 - v179
	if v182 <= v155 {
		v187 = v182
		v191 = v180
		goto L79
	} else {
		goto L84
	}
L84:
	;
	if v180 < v124 {
		v168 = v182
		v172 = v180
		goto L81
	} else {
		goto L85
	}
L85:
	;
	goto L82
L86:
	;
	if v187 != v223 {
		v265 = v191
		v266 = v225
		goto L93
	} else {
		goto L94
	}
L87:
	;
	v204 = v155
	v206 = v155
	goto L88
L88:
	;
	v211 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v206<<(uint(int32(1))%32))+uint32(_c_F_numeric_power[0]))))
	if v211 != 0 {
		v319 = int32(-1)
		goto L78
	} else {
		goto L90
	}
L89:
	;
	v223 = v215
	v225 = v213
	goto L86
L90:
	;
	v212 = int32(1)
	v213 = v206 + v212
	v215 = v204 - v212
	if v215 <= v187 {
		v223 = v215
		v225 = v213
		goto L86
	} else {
		goto L91
	}
L91:
	;
	if v213 < v154 {
		v204 = v215
		v206 = v213
		goto L88
	} else {
		goto L92
	}
L92:
	;
	goto L89
L93:
	;
	if v124 < v265 {
		goto L102
	} else {
		goto L103
	}
L94:
	;
	v234 = v191
	v235 = v225
	goto L95
L95:
	;
	if base.B2i32(v124 <= v234)|base.B2i32(v154 <= v235) != 0 {
		v265 = v234
		v266 = v235
		goto L93
	} else {
		goto L97
	}
L96:
	;
	if base.I32_extend16_s(v251) < base.I32_extend16_s(v249) {
		goto L99
	} else {
		goto L100
	}
L97:
	;
	v240 = int32(1)
	v249 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152+v234<<(uint(v240)%32)))))
	v251 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v235<<(uint(v240)%32))+uint32(_c_F_numeric_power[0]))))
	if v249 == v251 {
		v234 = v234 + v240
		v235 = v235 + v240
		goto L95
	} else {
		goto L98
	}
L98:
	;
	goto L96
L99:
	;
	v258 = int32(1)
	goto L101
L100:
	;
	v258 = int32(-1)
	goto L101
L101:
	;
	v329 = v258
	goto L77
L102:
	;
	v269 = v265
	goto L104
L103:
	;
	v269 = v124
	goto L104
L104:
	;
	v276 = v265
	goto L105
L105:
	;
	if v269 == v276 {
		goto L107
	} else {
		goto L108
	}
L106:
	;
	v319 = v302
	goto L78
L107:
	;
	if v154 < v266 {
		goto L110
	} else {
		goto L111
	}
L108:
	;
	goto L109
L109:
	;
	v302 = int32(1)
	v308 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v152+v276<<(uint(v302)%32)))))
	if v308 == int32(0) {
		v276 = v276 + v302
		goto L105
	} else {
		goto L119
	}
L110:
	;
	v281 = v266
	goto L112
L111:
	;
	v281 = v154
	goto L112
L112:
	;
	v289 = v266
	goto L113
L113:
	;
	if v281 == v289 {
		goto L115
	} else {
		goto L116
	}
L114:
	;
	v319 = int32(-1)
	goto L78
L115:
	;
	v329 = int32(0)
	goto L77
L116:
	;
	goto L117
L117:
	;
	v293 = int32(1)
	v298 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v289<<(uint(v293)%32))+uint32(_c_F_numeric_power[0]))))
	if v298 == int32(0) {
		v289 = v289 + v293
		goto L113
	} else {
		goto L118
	}
L118:
	;
	goto L114
L119:
	;
	goto L106
L120:
	;
	v332 = F_make_result_safe(m, int32(_a_F_numeric_power_3), int32(0))
	mBase = m.M
	v333 = m.ExcPending
	if v333 != 0 {
		goto L1
	} else {
		goto L121
	}
L121:
	;
	v2183 = v332
	goto L4
L122:
	;
	goto L30
L123:
	;
	v2183 = v343
	goto L4
L124:
	;
	v352 = v28 & int32(_a_F_numeric_power_5)
	if v352 == int32(_a_F_numeric_power_6) {
		goto L125
	} else {
		goto L126
	}
L125:
	;
	v355 = v28 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_7)
	goto L127
L126:
	;
	v355 = v352
	goto L127
L127:
	;
	v387 = v34
	v388 = base.B2i32(v355 != int32(_a_F_numeric_power_7))
	v390 = base.B2i32(v34 != int32(_a_F_numeric_power_1))
	goto L27
L128:
	;
	v362 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v364 = int32(base.Ui32(v362) >> (uint(int32(2)) % 32))
	if int32(0) <= base.I32_extend16_s(v87) {
		goto L129
	} else {
		goto L130
	}
L129:
	;
	v370 = int32(-8)
	goto L131
L130:
	;
	v370 = int32(-6)
	goto L131
L131:
	;
	if base.Ui32(v364+v370) < base.Ui32(int32(2)) {
		goto L25
	} else {
		goto L132
	}
L132:
	;
	v381 = v87 & int32(_a_F_numeric_power_5)
	if v381 == int32(_a_F_numeric_power_6) {
		goto L133
	} else {
		goto L134
	}
L133:
	;
	v384 = v87 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_7)
	goto L135
L134:
	;
	v384 = v381
	goto L135
L135:
	;
	v387 = v87
	v388 = base.B2i32(v29 == int32(-12288))
	v390 = base.B2i32(v384 == int32(_a_F_numeric_power_7))
	goto L27
L136:
	;
	v391 = int32(-1)
	goto L138
L137:
	;
	v391 = int32(1)
	goto L138
L138:
	;
	if base.B2i32(base.Ui32(v29) < base.Ui32(int32(-16384)))|v388 != 0 {
		v452 = v387
		v454 = v390
		v456 = v391
		goto L23
	} else {
		goto L139
	}
L139:
	;
	v395 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v423 = v387
	v424 = int32(base.Ui32(v395) >> (uint(int32(2)) % 32))
	v425 = v390
	v426 = v391
	goto L24
L140:
	;
	v400 = int32(1)
	if v34 == int32(_a_F_numeric_power_1) {
		goto L143
	} else {
		goto L144
	}
L141:
	;
	goto L142
L142:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v408 = m.ExcPending
	if v408 != 0 {
		goto L1
	} else {
		goto L146
	}
L143:
	;
	v403 = v400
	goto L145
L144:
	;
	v403 = int32(-1)
	goto L145
L145:
	;
	v460 = v34
	v461 = v400
	v462 = int32(0)
	v464 = v403
	goto L22
L146:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v411 = m.ExcPending
	if v411 != 0 {
		goto L1
	} else {
		goto L147
	}
L147:
	;
	F_errmsg(m, int32(_a_F_numeric_power_8), int32(0))
	mBase = m.M
	v415 = m.ExcPending
	if v415 != 0 {
		goto L1
	} else {
		goto L148
	}
L148:
	;
	F_errfinish(m, int32(_a_F_numeric_power_9), int32(3969), int32(_a_F_numeric_power_10))
	mBase = m.M
	v420 = m.ExcPending
	if v420 != 0 {
		goto L1
	} else {
		goto L149
	}
L149:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L150:
	;
	v423 = v87
	v424 = v364
	v425 = v2
	v426 = v2
	goto L24
L151:
	;
	v432 = int32(-8)
	goto L153
L152:
	;
	v432 = int32(-6)
	goto L153
L153:
	;
	v435 = int32(base.Ui32(v424+v432) >> (uint(int32(1)) % 32))
	if int32(0) <= base.I32_extend16_s(v423) {
		goto L154
	} else {
		goto L155
	}
L154:
	;
	v436 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+6)))
	v446 = v436
	goto L156
L155:
	;
	v446 = v423<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v423&int32(63)
	goto L156
L156:
	;
	if v435 == int32(0) {
		v452 = v423
		v454 = v425
		v456 = v426
		goto L23
	} else {
		goto L157
	}
L157:
	;
	if v446+int32(1) < v435 {
		goto L12
	} else {
		goto L158
	}
L158:
	;
	v452 = v423
	v454 = v425
	v456 = v426
	goto L23
L159:
	;
	v460 = v452
	v461 = v457
	v462 = v454
	v464 = v456
	goto L22
L160:
	;
	v472 = int32(-8)
	goto L162
L161:
	;
	v472 = int32(-6)
	goto L162
L162:
	;
	v475 = int32(base.Ui32(int32(base.Ui32(v465)>>(uint(int32(2))%32))+v472) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v475
	if int32(0) <= v29 {
		goto L163
	} else {
		goto L164
	}
L163:
	;
	v477 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
	v487 = v477
	goto L165
L164:
	;
	v487 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
	goto L165
L165:
	;
	if v475 == int32(0) {
		v686 = v460
		v687 = v461
		v689 = v462
		v691 = v464
		goto L21
	} else {
		goto L166
	}
L166:
	;
	v495 = v28 & int32(_a_F_numeric_power_5)
	if v495 == int32(_a_F_numeric_power_6) {
		goto L167
	} else {
		goto L168
	}
L167:
	;
	v498 = v28 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_7)
	goto L169
L168:
	;
	v498 = v495
	goto L169
L169:
	;
	if v498 != 0 {
		v686 = v460
		v687 = v461
		v689 = v462
		v691 = v464
		goto L21
	} else {
		goto L170
	}
L170:
	;
	if v29 < int32(0) {
		goto L171
	} else {
		goto L172
	}
L171:
	;
	v503 = int32(6)
	goto L173
L172:
	;
	v503 = int32(8)
	goto L173
L173:
	;
	v504 = v21 + v503
	v506 = int32(1)
	v507 = int32(0)
	if base.B2i32(v507 < v487)&base.B2i32(v507 < v475) == v507 {
		v539 = v487
		v543 = v507
		goto L176
	} else {
		goto L177
	}
L174:
	;
	if v681 != 0 {
		v686 = v460
		v687 = v461
		v689 = v462
		v691 = v464
		goto L21
	} else {
		goto L217
	}
L175:
	;
	v681 = v671
	goto L174
L176:
	;
	if int32(0)|base.B2i32(v507 <= v539) != 0 {
		v575 = v507
		v577 = v507
		goto L183
	} else {
		goto L184
	}
L177:
	;
	v520 = v487
	v524 = v507
	goto L178
L178:
	;
	v530 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v504+v524<<(uint(int32(1))%32)))))
	if v530 != 0 {
		v671 = int32(1)
		goto L175
	} else {
		goto L180
	}
L179:
	;
	v539 = v534
	v543 = v532
	goto L176
L180:
	;
	v531 = int32(1)
	v532 = v524 + v531
	v534 = v520 - v531
	if v534 <= v507 {
		v539 = v534
		v543 = v532
		goto L176
	} else {
		goto L181
	}
L181:
	;
	if v532 < v475 {
		v520 = v534
		v524 = v532
		goto L178
	} else {
		goto L182
	}
L182:
	;
	goto L179
L183:
	;
	if v539 != v575 {
		v617 = v543
		v618 = v577
		goto L190
	} else {
		goto L191
	}
L184:
	;
	v556 = v507
	v558 = v507
	goto L185
L185:
	;
	v563 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v558<<(uint(int32(1))%32))+uint32(_c_F_numeric_power[0]))))
	if v563 != 0 {
		v671 = int32(-1)
		goto L175
	} else {
		goto L187
	}
L186:
	;
	v575 = v567
	v577 = v565
	goto L183
L187:
	;
	v564 = int32(1)
	v565 = v558 + v564
	v567 = v556 - v564
	if v567 <= v539 {
		v575 = v567
		v577 = v565
		goto L183
	} else {
		goto L188
	}
L188:
	;
	if v565 < v506 {
		v556 = v567
		v558 = v565
		goto L185
	} else {
		goto L189
	}
L189:
	;
	goto L186
L190:
	;
	if v475 < v617 {
		goto L199
	} else {
		goto L200
	}
L191:
	;
	v586 = v543
	v587 = v577
	goto L192
L192:
	;
	if base.B2i32(v475 <= v586)|base.B2i32(v506 <= v587) != 0 {
		v617 = v586
		v618 = v587
		goto L190
	} else {
		goto L194
	}
L193:
	;
	if base.I32_extend16_s(v603) < base.I32_extend16_s(v601) {
		goto L196
	} else {
		goto L197
	}
L194:
	;
	v592 = int32(1)
	v601 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v504+v586<<(uint(v592)%32)))))
	v603 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v587<<(uint(v592)%32))+uint32(_c_F_numeric_power[0]))))
	if v601 == v603 {
		v586 = v586 + v592
		v587 = v587 + v592
		goto L192
	} else {
		goto L195
	}
L195:
	;
	goto L193
L196:
	;
	v610 = int32(1)
	goto L198
L197:
	;
	v610 = int32(-1)
	goto L198
L198:
	;
	v681 = v610
	goto L174
L199:
	;
	v621 = v617
	goto L201
L200:
	;
	v621 = v475
	goto L201
L201:
	;
	v628 = v617
	goto L202
L202:
	;
	if v621 == v628 {
		goto L204
	} else {
		goto L205
	}
L203:
	;
	v671 = v654
	goto L175
L204:
	;
	if v506 < v618 {
		goto L207
	} else {
		goto L208
	}
L205:
	;
	goto L206
L206:
	;
	v654 = int32(1)
	v660 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v504+v628<<(uint(v654)%32)))))
	if v660 == int32(0) {
		v628 = v628 + v654
		goto L202
	} else {
		goto L216
	}
L207:
	;
	v633 = v618
	goto L209
L208:
	;
	v633 = v506
	goto L209
L209:
	;
	v641 = v618
	goto L210
L210:
	;
	if v633 == v641 {
		goto L212
	} else {
		goto L213
	}
L211:
	;
	v671 = int32(-1)
	goto L175
L212:
	;
	v681 = int32(0)
	goto L174
L213:
	;
	goto L214
L214:
	;
	v645 = int32(1)
	v650 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v641<<(uint(v645)%32))+uint32(_c_F_numeric_power[0]))))
	if v650 == int32(0) {
		v641 = v641 + v645
		goto L210
	} else {
		goto L215
	}
L215:
	;
	goto L211
L216:
	;
	goto L203
L217:
	;
	v684 = F_make_result_safe(m, int32(_a_F_numeric_power_3), int32(0))
	mBase = m.M
	v685 = m.ExcPending
	if v685 != 0 {
		goto L1
	} else {
		goto L218
	}
L218:
	;
	v2183 = v684
	goto L4
L219:
	;
	goto L20
L220:
	;
	v2183 = v702
	goto L4
L221:
	;
	v709 = F_make_result_safe(m, int32(_a_F_numeric_power_11), int32(0))
	mBase = m.M
	v710 = m.ExcPending
	if v710 != 0 {
		goto L1
	} else {
		goto L224
	}
L222:
	;
	goto L223
L223:
	;
	if v686&int32(_a_F_numeric_power_2) != int32(_a_F_numeric_power_1) {
		v922 = v705
		v923 = v689
		goto L14
	} else {
		goto L225
	}
L224:
	;
	v2183 = v709
	goto L4
L225:
	;
	if base.Ui32(v29) <= base.Ui32(int32(-16385)) {
		goto L17
	} else {
		goto L226
	}
L226:
	;
	v903 = v705
	v908 = int32(1)
	goto L16
L227:
	;
	v903 = v719
	v908 = int32(1)
	goto L16
L228:
	;
	v796 = int32(_a_F_numeric_power_12)
	v803 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[1]))
	v804 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[2]))
	v805 = *(*int32)(unsafe.Add(mBase, uint32(v726)))
	if v805 == int32(0) {
		goto L248
	} else {
		goto L249
	}
L229:
	;
	v737 = int32(-8)
	goto L231
L230:
	;
	v737 = int32(-6)
	goto L231
L231:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726))) = int32(base.Ui32(int32(base.Ui32(v729)>>(uint(int32(2))%32))+v737) >> (uint(int32(1)) % 32))
	v742 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	if v742 < int32(0) {
		goto L233
	} else {
		goto L234
	}
L232:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+4)) = v758
	v760 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v761 = int32(_a_F_numeric_power_5)
	v762 = v760 & v761
	if v762 != v761 {
		goto L237
	} else {
		goto L238
	}
L233:
	;
	v746 = v742 & int32(_a_F_numeric_power_13)
	v758 = v746<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v746&int32(63)
	goto L232
L234:
	;
	goto L235
L235:
	;
	v756 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
	v758 = v756
	goto L232
L236:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+8)) = v773
	v775 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	if v775 < int32(0) {
		goto L241
	} else {
		goto L242
	}
L237:
	;
	if v762 != int32(_a_F_numeric_power_6) {
		v773 = v762
		goto L236
	} else {
		goto L240
	}
L238:
	;
	goto L239
L239:
	;
	v773 = v760 & int32(_a_F_numeric_power_14)
	goto L236
L240:
	;
	v773 = v760 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_7)
	goto L236
L241:
	;
	v784 = int32(base.Ui32(v775)>>(uint(int32(7))%32)) & int32(63)
	goto L243
L242:
	;
	v784 = v775 & int32(_a_F_numeric_power_15)
	goto L243
L243:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+12)) = v784
	v786 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	v787 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v726)+16)) = v787
	if v786 < v787 {
		goto L244
	} else {
		goto L245
	}
L244:
	;
	v793 = int32(6)
	goto L246
L245:
	;
	v793 = int32(8)
	goto L246
L246:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v726)+20)) = v21 + v793
	goto L228
L247:
	;
	if v841 == int32(0) {
		goto L272
	} else {
		goto L273
	}
L248:
	;
	if v804 == int32(0) {
		goto L251
	} else {
		goto L252
	}
L249:
	;
	goto L250
L250:
	;
	v816 = *(*int32)(unsafe.Add(mBase, uint32(v726)+8))
	if v804 == int32(0) {
		goto L257
	} else {
		goto L258
	}
L251:
	;
	v841 = int32(0)
	goto L247
L252:
	;
	goto L253
L253:
	;
	if v803 == int32(_a_F_numeric_power_7) {
		goto L254
	} else {
		goto L255
	}
L254:
	;
	v815 = int32(1)
	goto L256
L255:
	;
	v815 = int32(-1)
	goto L256
L256:
	;
	v841 = v815
	goto L247
L257:
	;
	if v816 != 0 {
		goto L260
	} else {
		goto L261
	}
L258:
	;
	goto L259
L259:
	;
	v822 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[3]))
	v823 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[4]))
	v824 = *(*int32)(unsafe.Add(mBase, uint32(v726)+4))
	v825 = *(*int32)(unsafe.Add(mBase, uint32(v726)+20))
	if v816 == int32(0) {
		goto L263
	} else {
		goto L264
	}
L260:
	;
	v821 = int32(-1)
	goto L262
L261:
	;
	v821 = int32(1)
	goto L262
L262:
	;
	v841 = v821
	goto L247
L263:
	;
	if v803 == int32(_a_F_numeric_power_7) {
		goto L266
	} else {
		goto L267
	}
L264:
	;
	goto L265
L265:
	;
	if v803 == int32(0) {
		goto L269
	} else {
		goto L270
	}
L266:
	;
	v841 = int32(1)
	goto L247
L267:
	;
	goto L268
L268:
	;
	v831 = F_cmp_abs_common(m, v825, v805, v824, v823, v804, v822)
	mBase = m.M
	v841 = v831
	goto L247
L269:
	;
	v841 = int32(-1)
	goto L247
L270:
	;
	goto L271
L271:
	;
	v835 = F_cmp_abs_common(m, v823, v804, v822, v825, v805, v824)
	mBase = m.M
	v841 = v835
	goto L247
L272:
	;
	v846 = F_make_result_safe(m, int32(_a_F_numeric_power_3), int32(0))
	mBase = m.M
	v847 = m.ExcPending
	if v847 != 0 {
		goto L1
	} else {
		goto L275
	}
L273:
	;
	goto L274
L274:
	;
	v848 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v848
	v851 = v18 + int32(48)
	v852 = int32(_a_F_numeric_power_3)
	v859 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[5]))
	v860 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[6]))
	v861 = *(*int32)(unsafe.Add(mBase, uint32(v851)))
	if v861 == v848 {
		goto L277
	} else {
		goto L278
	}
L275:
	;
	v2183 = v846
	goto L4
L276:
	;
	v903 = v705
	v908 = base.B2i32(int32(0) < v897)
	goto L16
L277:
	;
	if v860 == int32(0) {
		goto L280
	} else {
		goto L281
	}
L278:
	;
	goto L279
L279:
	;
	v872 = *(*int32)(unsafe.Add(mBase, uint32(v851)+8))
	if v860 == int32(0) {
		goto L286
	} else {
		goto L287
	}
L280:
	;
	v897 = int32(0)
	goto L276
L281:
	;
	goto L282
L282:
	;
	if v859 == int32(_a_F_numeric_power_7) {
		goto L283
	} else {
		goto L284
	}
L283:
	;
	v871 = int32(1)
	goto L285
L284:
	;
	v871 = int32(-1)
	goto L285
L285:
	;
	v897 = v871
	goto L276
L286:
	;
	if v872 != 0 {
		goto L289
	} else {
		goto L290
	}
L287:
	;
	goto L288
L288:
	;
	v878 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[7]))
	v879 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[8]))
	v880 = *(*int32)(unsafe.Add(mBase, uint32(v851)+4))
	v881 = *(*int32)(unsafe.Add(mBase, uint32(v851)+20))
	if v872 == int32(0) {
		goto L292
	} else {
		goto L293
	}
L289:
	;
	v877 = int32(-1)
	goto L291
L290:
	;
	v877 = int32(1)
	goto L291
L291:
	;
	v897 = v877
	goto L276
L292:
	;
	if v859 == int32(_a_F_numeric_power_7) {
		goto L295
	} else {
		goto L296
	}
L293:
	;
	goto L294
L294:
	;
	if v859 == int32(0) {
		goto L298
	} else {
		goto L299
	}
L295:
	;
	v897 = int32(1)
	goto L276
L296:
	;
	goto L297
L297:
	;
	v887 = F_cmp_abs_common(m, v881, v861, v880, v879, v860, v878)
	mBase = m.M
	v897 = v887
	goto L276
L298:
	;
	v897 = int32(-1)
	goto L276
L299:
	;
	goto L300
L300:
	;
	v891 = F_cmp_abs_common(m, v879, v860, v878, v881, v861, v880)
	mBase = m.M
	v897 = v891
	goto L276
L301:
	;
	v912 = F_make_result_safe(m, int32(_a_F_numeric_power_16), int32(0))
	mBase = m.M
	v913 = m.ExcPending
	if v913 != 0 {
		goto L1
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	v916 = F_make_result_safe(m, int32(_a_F_numeric_power_11), int32(0))
	mBase = m.M
	v917 = m.ExcPending
	if v917 != 0 {
		goto L1
	} else {
		goto L305
	}
L304:
	;
	v2183 = v912
	goto L4
L305:
	;
	v2183 = v916
	goto L4
L306:
	;
	if v922 != 0 {
		goto L309
	} else {
		goto L310
	}
L307:
	;
	goto L308
L308:
	;
	if v923 != 0 {
		goto L314
	} else {
		goto L315
	}
L309:
	;
	v931 = F_make_result_safe(m, int32(_a_F_numeric_power_16), int32(0))
	mBase = m.M
	v932 = m.ExcPending
	if v932 != 0 {
		goto L1
	} else {
		goto L312
	}
L310:
	;
	goto L311
L311:
	;
	v935 = F_make_result_safe(m, int32(_a_F_numeric_power_11), int32(0))
	mBase = m.M
	v936 = m.ExcPending
	if v936 != 0 {
		goto L1
	} else {
		goto L313
	}
L312:
	;
	v2183 = v931
	goto L4
L313:
	;
	v2183 = v935
	goto L4
L314:
	;
	v939 = F_make_result_safe(m, int32(_a_F_numeric_power_11), int32(0))
	mBase = m.M
	v940 = m.ExcPending
	if v940 != 0 {
		goto L1
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v942 = v18 + int32(24)
	v945 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v950 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	if int32(0) <= v950 {
		goto L319
	} else {
		goto L320
	}
L317:
	;
	v2183 = v939
	goto L4
L318:
	;
	v1012 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v1012 <= int32(0) {
		goto L337
	} else {
		goto L338
	}
L319:
	;
	v953 = int32(-8)
	goto L321
L320:
	;
	v953 = int32(-6)
	goto L321
L321:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v942))) = int32(base.Ui32(int32(base.Ui32(v945)>>(uint(int32(2))%32))+v953) >> (uint(int32(1)) % 32))
	v958 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	if v958 < int32(0) {
		goto L323
	} else {
		goto L324
	}
L322:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v942)+4)) = v974
	v976 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	v977 = int32(_a_F_numeric_power_5)
	v978 = v976 & v977
	if v978 != v977 {
		goto L327
	} else {
		goto L328
	}
L323:
	;
	v962 = v958 & int32(_a_F_numeric_power_13)
	v974 = v962<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v962&int32(63)
	goto L322
L324:
	;
	goto L325
L325:
	;
	v972 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+6)))
	v974 = v972
	goto L322
L326:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v942)+8)) = v989
	v991 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	if v991 < int32(0) {
		goto L331
	} else {
		goto L332
	}
L327:
	;
	if v978 != int32(_a_F_numeric_power_6) {
		v989 = v978
		goto L326
	} else {
		goto L330
	}
L328:
	;
	goto L329
L329:
	;
	v989 = v976 & int32(_a_F_numeric_power_14)
	goto L326
L330:
	;
	v989 = v976 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_7)
	goto L326
L331:
	;
	v1000 = int32(base.Ui32(v991)>>(uint(int32(7))%32)) & int32(63)
	goto L333
L332:
	;
	v1000 = v991 & int32(_a_F_numeric_power_15)
	goto L333
L333:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v942)+12)) = v1000
	v1002 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	v1003 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v942)+16)) = v1003
	if v1002 < v1003 {
		goto L334
	} else {
		goto L335
	}
L334:
	;
	v1009 = int32(6)
	goto L336
L335:
	;
	v1009 = int32(8)
	goto L336
L336:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v942)+20)) = v26 + v1009
	goto L318
L337:
	;
	v1036 = F_make_result_safe(m, int32(_a_F_numeric_power_16), int32(0))
	mBase = m.M
	v1037 = m.ExcPending
	if v1037 != 0 {
		goto L1
	} else {
		goto L342
	}
L338:
	;
	v1015 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v1012 != v1015+int32(1) {
		goto L337
	} else {
		goto L339
	}
L339:
	;
	v1019 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v1020 = int32(1)
	v1025 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1019+v1012<<(uint(v1020)%32)-int32(2)))))
	if v1025&v1020 == int32(0) {
		goto L337
	} else {
		goto L340
	}
L340:
	;
	v1032 = F_make_result_safe(m, int32(_a_F_numeric_power_17), int32(0))
	mBase = m.M
	v1033 = m.ExcPending
	if v1033 != 0 {
		goto L1
	} else {
		goto L341
	}
L341:
	;
	v2183 = v1032
	goto L4
L342:
	;
	v2183 = v1036
	goto L4
L343:
	;
	v1054 = int32(-8)
	goto L345
L344:
	;
	v1054 = int32(-6)
	goto L345
L345:
	;
	v1055 = int32(base.Ui32(v1045)>>(uint(int32(2))%32)) + v1054
	v1057 = int32(base.Ui32(v1055) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v1057
	if int32(0) <= v1051 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1059 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
	v1069 = v1059
	goto L348
L347:
	;
	v1069 = v1050<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v1050&int32(63)
	goto L348
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v1069
	v1071 = int32(_a_F_numeric_power_5)
	v1072 = v1050 & v1071
	if v1072 != v1071 {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v1083
	v1085 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v1085
	v1094 = base.B2i32(v1051 < v1085)
	if v1051 < v1085 {
		goto L354
	} else {
		goto L355
	}
L350:
	;
	if v1072 != int32(_a_F_numeric_power_6) {
		v1083 = v1072
		goto L349
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	v1083 = v1050 & int32(_a_F_numeric_power_14)
	goto L349
L353:
	;
	v1083 = v1050 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_7)
	goto L349
L354:
	;
	v1095 = int32(base.Ui32(v1050)>>(uint(int32(7))%32)) & int32(63)
	goto L356
L355:
	;
	v1095 = v1050 & int32(_a_F_numeric_power_15)
	goto L356
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v1095
	if v1051 < v1085 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1099 = int32(6)
	goto L359
L358:
	;
	v1099 = int32(8)
	goto L359
L359:
	;
	v1100 = v21 + v1099
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v1100
	v1102 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v1107 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	v1108 = base.I32_extend16_s(v1107)
	v1110 = base.B2i32(int32(0) <= v1108)
	if int32(0) <= v1108 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1111 = int32(-8)
	goto L362
L361:
	;
	v1111 = int32(-6)
	goto L362
L362:
	;
	v1114 = int32(base.Ui32(int32(base.Ui32(v1102)>>(uint(int32(2))%32))+v1111) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v1114
	if int32(0) <= v1108 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1116 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+6)))
	v1126 = v1116
	goto L365
L364:
	;
	v1126 = v1107<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v1107&int32(63)
	goto L365
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v1126
	v1128 = int32(_a_F_numeric_power_5)
	v1129 = v1107 & v1128
	if v1129 != v1128 {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v1140
	v1142 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v1142
	v1151 = base.B2i32(v1108 < v1142)
	if v1108 < v1142 {
		goto L371
	} else {
		goto L372
	}
L367:
	;
	if v1129 != int32(_a_F_numeric_power_6) {
		v1140 = v1129
		goto L366
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1140 = v1107 & int32(_a_F_numeric_power_14)
	goto L366
L370:
	;
	v1140 = v1107 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_7)
	goto L366
L371:
	;
	v1152 = int32(base.Ui32(v1107)>>(uint(int32(7))%32)) & int32(63)
	goto L373
L372:
	;
	v1152 = v1107 & int32(_a_F_numeric_power_15)
	goto L373
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v1152
	if v1108 < v1142 {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1156 = int32(6)
	goto L376
L375:
	;
	v1156 = int32(8)
	goto L376
L376:
	;
	v1157 = v26 + v1156
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v1157
	v1161 = v1126 + int32(1)
	if v1161 < v1114 {
		goto L378
	} else {
		goto L379
	}
L377:
	;
	if v1057 == int32(0) {
		goto L543
	} else {
		goto L544
	}
L378:
	;
	v1163 = v1114
	goto L380
L379:
	;
	v1163 = int32(0)
	goto L380
L380:
	;
	if v1163 != 0 {
		goto L377
	} else {
		goto L381
	}
L381:
	;
	v1168 = F_numericvar_to_int64(m, v18+int32(24), v18+int32(96))
	mBase = m.M
	v1169 = m.ExcPending
	if v1169 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	if v1168 == int32(0) {
		goto L377
	} else {
		goto L383
	}
L383:
	;
	v1172 = *(*int64)(unsafe.Add(mBase, uint32(v18)+96))
	if base.Ui64(int64(4294967295)) < base.Ui64(v1172+int64(2147483648)) {
		goto L377
	} else {
		goto L384
	}
L384:
	;
	v1177 = base.I32_wrap_i64(v1172)
	if v1057 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1178 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1100))))
	v1179 = base.F64_convert_i32_s(v1178)
	if v1057 == int32(1) {
		goto L388
	} else {
		goto L389
	}
L386:
	;
	v1371 = float64(0)
	goto L387
L387:
	;
	if base.F64_lt(base.F64_add(v1371, float64(1)), float64(-1000)) != 0 {
		goto L414
	} else {
		goto L415
	}
L388:
	;
	v1218 = v1179
	v1219 = v1069 << (uint(int32(2)) % 32)
	goto L390
L389:
	;
	v1184 = int32(2)
	v1186 = v1057 - v1184
	if base.Ui32(v1184) <= base.Ui32(v1186) {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	v1230 = base.I64_reinterpret_f64(v1218)
	if v1230 <= int64(4503599627370495) {
		goto L401
	} else {
		goto L402
	}
L391:
	;
	v1189 = v1184
	goto L393
L392:
	;
	v1189 = v1186
	goto L393
L393:
	;
	v1192 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1100)+2)))
	v1194 = base.F64_add(base.F64_mul(v1179, float64(10000)), base.F64_convert_i32_s(v1192))
	if v1186 == int32(0) {
		v1210 = v1194
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1218 = v1210
	v1219 = (v1069-v1189)<<(uint(int32(2))%32) - int32(4)
	goto L390
L395:
	;
	v1199 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1100)+4)))
	v1201 = base.F64_add(base.F64_mul(v1194, float64(10000)), base.F64_convert_i32_s(v1199))
	if v1186 == int32(1) {
		v1210 = v1201
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v1206 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1100)+6)))
	v1210 = base.F64_add(base.F64_mul(v1201, float64(10000)), base.F64_convert_i32_s(v1206))
	goto L394
L397:
	;
	v1366 = base.F64_mul(base.F64_add(v1362, base.F64_convert_i32_s(v1219)), base.F64_convert_i32_s(v1177))
	if base.F64_gt(v1366, float64(131072)) != 0 {
		goto L11
	} else {
		goto L413
	}
L398:
	;
	v1362 = v1340
	goto L397
L399:
	;
	v1266 = v1264 + int32(_a_F_numeric_power_18)
	v1270 = base.F64_convert_i32_s(int32(base.Ui32(v1266)>>(uint(int32(20))%32)) + v1263)
	v1272 = base.F64_mul(v1270, float64(0.30102999566361177))
	v1285 = base.F64_add(base.F64_reinterpret_i64(v1261&int64(4294967295)|base.I64_extend_i32_u(v1266&int32(_a_F_numeric_power_19)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
	v1288 = base.F64_mul(v1285, base.F64_mul(v1285, float64(0.5)))
	v1293 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v1285, v1288)) & int64(-4294967296))
	v1294 = float64(0.4342944818781689)
	v1295 = base.F64_mul(v1293, v1294)
	v1296 = base.F64_add(v1272, v1295)
	v1301 = base.F64_div(v1285, base.F64_add(v1285, float64(2)))
	v1302 = base.F64_mul(v1301, v1301)
	v1303 = base.F64_mul(v1302, v1302)
	v1328 = base.F64_add(base.F64_mul(v1301, base.F64_add(v1288, base.F64_add(base.F64_mul(v1303, base.F64_add(base.F64_mul(v1303, base.F64_add(base.F64_mul(v1303, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v1302, base.F64_add(base.F64_mul(v1303, base.F64_add(base.F64_mul(v1303, base.F64_add(base.F64_mul(v1303, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v1285, v1293), v1288))
	v1340 = base.F64_add(v1296, base.F64_add(base.F64_add(v1295, base.F64_sub(v1272, v1296)), base.F64_add(base.F64_mul(v1328, v1294), base.F64_add(base.F64_mul(v1270, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v1328, v1293), float64(2.5082946711645275e-11))))))
	goto L398
L400:
	;
	v1257 = base.I64_reinterpret_f64(base.F64_mul(v1218, float64(1.8014398509481984e+16)))
	v1261 = v1257
	v1263 = int32(-1077)
	v1264 = base.I32_wrap_i64(int64(base.Ui64(v1257) >> (uint(int64(32)) % 64)))
	goto L399
L401:
	;
	if base.F64_eq(v1218, float64(0)) != 0 {
		goto L404
	} else {
		goto L405
	}
L402:
	;
	goto L403
L403:
	;
	if base.Ui64(int64(9218868437227405311)) < base.Ui64(v1230) {
		v1340 = v1218
		goto L398
	} else {
		goto L408
	}
L404:
	;
	v1362 = base.F64_div(float64(-1), base.F64_mul(v1218, v1218))
	goto L397
L405:
	;
	goto L406
L406:
	;
	if int64(0) <= v1230 {
		goto L400
	} else {
		goto L407
	}
L407:
	;
	v1362 = base.F64_div(base.F64_sub(v1218, v1218), float64(0))
	goto L397
L408:
	;
	v1245 = int32(-1023)
	v1247 = int64(base.Ui64(v1230) >> (uint(int64(32)) % 64))
	if v1247 != int64(1072693248) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1261 = v1230
	v1263 = v1245
	v1264 = base.I32_wrap_i64(v1247)
	goto L399
L410:
	;
	goto L411
L411:
	;
	if base.I32_wrap_i64(v1230) != 0 {
		v1261 = v1230
		v1263 = v1245
		v1264 = int32(1072693248)
		goto L399
	} else {
		goto L412
	}
L412:
	;
	v1362 = float64(0)
	goto L397
L413:
	;
	v1371 = v1366
	goto L387
L414:
	;
	v1376 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1376 != 0 {
		goto L417
	} else {
		goto L418
	}
L415:
	;
	goto L416
L416:
	;
	v1387 = base.I32_trunc_sat_f64_s(v1371)
	v1388 = int32(16) - v1387
	if v1095 < v1388 {
		goto L421
	} else {
		goto L422
	}
L417:
	;
	F_pfree(m, v1376)
	mBase = m.M
	v1378 = m.ExcPending
	if v1378 != 0 {
		goto L1
	} else {
		goto L420
	}
L418:
	;
	goto L419
L419:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = int64(4294967296000)
	v1381 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1381
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v1381
	goto L5
L420:
	;
	goto L419
L421:
	;
	v1390 = v1388
	goto L423
L422:
	;
	v1390 = v1095
	goto L423
L423:
	;
	if base.Ui32(v1152) < base.Ui32(v1390) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1392 = v1390
	goto L426
L425:
	;
	v1392 = v1152
	goto L426
L426:
	;
	if base.Ui32(int32(1000)) <= base.Ui32(v1392) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1395 = int32(1000)
	goto L429
L428:
	;
	v1395 = v1392
	goto L429
L429:
	;
	switch v1177 + int32(1) {
	case 0:
		goto L432
	case 1:
		goto L434
	case 2:
		goto L433
	case 3:
		goto L431
	default:
		goto L430
	}
L430:
	;
	if v1057 == int32(0) {
		goto L476
	} else {
		goto L477
	}
L431:
	;
	v1578 = v18 + int32(48)
	F_mul_var(m, v1578, v1578, v18, v1395)
	mBase = m.M
	v1580 = m.ExcPending
	if v1580 != 0 {
		goto L1
	} else {
		goto L475
	}
L432:
	;
	v1573 = int32(1)
	F_div_var(m, int32(_a_F_numeric_power_3), v18+int32(48), v18, v1395, v1573, v1573)
	mBase = m.M
	v1576 = m.ExcPending
	if v1576 != 0 {
		goto L1
	} else {
		goto L474
	}
L433:
	;
	v1418 = v1055 & int32(-2)
	v1421 = F_palloc(m, v1418+int32(2))
	mBase = m.M
	v1422 = m.ExcPending
	if v1422 != 0 {
		goto L1
	} else {
		goto L440
	}
L434:
	;
	v1399 = F_palloc(m, int32(4))
	mBase = m.M
	v1400 = m.ExcPending
	if v1400 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1399))) = int32(_a_F_numeric_power_20)
	v1403 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1403 != 0 {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	F_pfree(m, v1403)
	mBase = m.M
	v1405 = m.ExcPending
	if v1405 != 0 {
		goto L1
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1399 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1399
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v1395
	v1412 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1412
	v1415 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_power[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1415
	goto L5
L439:
	;
	goto L438
L440:
	;
	v1423 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1421))) = uint16(v1423)
	if base.B2i32(v1057 == v1423)|base.B2i32(v1418 == v1423) == v1423 {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	base.MemoryCopy(m, v1421+int32(2), v1100, v1418)
	goto L443
L442:
	;
	goto L443
L443:
	;
	v1435 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1435 != 0 {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	F_pfree(m, v1435)
	mBase = m.M
	v1437 = m.ExcPending
	if v1437 != 0 {
		goto L1
	} else {
		goto L447
	}
L445:
	;
	goto L446
L446:
	;
	v1438 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v1438
	v1440 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1440
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1421
	v1443 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1421 + v1443
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v1395
	v1452 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v1455 = v1395 + v1452<<(uint(v1443)%32)
	if v1455+int32(4) < int32(0) {
		goto L449
	} else {
		goto L450
	}
L447:
	;
	goto L446
L448:
	;
	goto L5
L449:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(0)
	goto L448
L450:
	;
	goto L451
L451:
	;
	v1464 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v1466 = v1395 & int32(3)
	v1470 = base.I32_div_s(v1455+int32(7), int32(4))
	v1471 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v1471 <= v1470 {
		goto L456
	} else {
		goto L457
	}
L452:
	;
	goto L448
L453:
	;
	if int32(0) <= v1536 {
		goto L452
	} else {
		goto L473
	}
L454:
	;
	v1516 = v1510
	goto L467
L455:
	;
	v1485 = int32(1)
	v1486 = v1470 - v1485
	v1489 = v1464 + v1486<<(uint(v1485)%32)
	v1490 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1489))))
	v1491 = int32(2)
	v1493 = *(*int32)(unsafe.Add(mBase, uint32(v1466<<(uint(v1491)%32))+uint32(_c_F_numeric_power[9])))
	v1494 = base.I32_rem_s(v1490, v1493)
	v1495 = v1490 - v1494
	*(*uint16)(unsafe.Add(mBase, uint32(v1489))) = uint16(v1495)
	v1498 = base.I32_div_s(v1493, v1491)
	if v1494 < v1498 {
		v1536 = v1486
		goto L453
	} else {
		goto L462
	}
L456:
	;
	if base.B2i32(v1466 == int32(0))|base.B2i32(v1470 != v1471) != 0 {
		goto L452
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1470
	if v1466 != 0 {
		goto L455
	} else {
		goto L460
	}
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1470
	goto L455
L460:
	;
	v1482 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1464+v1470<<(uint(int32(1))%32)))))
	if v1482 <= int32(_a_F_numeric_power_21) {
		v1536 = v1470
		goto L453
	} else {
		goto L461
	}
L461:
	;
	v1510 = v1470
	goto L454
L462:
	;
	v1501 = v1493 + base.I32_extend16_s(v1495)
	if int32(_a_F_numeric_power_22) < v1501 {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v1506 = v1501 + int32(_a_F_numeric_power_23)
	goto L465
L464:
	;
	v1506 = v1501
	goto L465
L465:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1489))) = uint16(v1506)
	if v1501 < int32(_a_F_numeric_power_24) {
		v1536 = v1486
		goto L453
	} else {
		goto L466
	}
L466:
	;
	v1510 = v1486
	goto L454
L467:
	;
	v1522 = int32(1)
	v1523 = v1516 - v1522
	v1526 = v1464 + v1523<<(uint(v1522)%32)
	v1529 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1526))))
	v1531 = base.B2i32(int32(_a_F_numeric_power_25) < v1529)
	if int32(_a_F_numeric_power_25) < v1529 {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	v1536 = v1523
	goto L453
L469:
	;
	v1532 = int32(-9999)
	goto L471
L470:
	;
	v1532 = v1522
	goto L471
L471:
	;
	v1533 = v1532 + v1529
	*(*uint16)(unsafe.Add(mBase, uint32(v1526))) = uint16(v1533)
	if int32(_a_F_numeric_power_25) < v1529 {
		v1516 = v1523
		goto L467
	} else {
		goto L472
	}
L472:
	;
	goto L468
L473:
	;
	v1544 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1544 - int32(2)
	v1548 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v1549 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1548 + v1549
	v1552 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v1552 + v1549
	goto L452
L474:
	;
	goto L5
L475:
	;
	goto L5
L476:
	;
	if v1172 < int64(0) {
		goto L10
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v1596 = v1055 & int32(-2)
	v1598 = v1596 + int32(2)
	v1599 = F_palloc(m, v1598)
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L1
	} else {
		goto L484
	}
L479:
	;
	v1585 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1585 != 0 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	F_pfree(m, v1585)
	mBase = m.M
	v1587 = m.ExcPending
	if v1587 != 0 {
		goto L1
	} else {
		goto L483
	}
L481:
	;
	goto L482
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v1395
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
	v1591 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1591
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v1591
	goto L5
L483:
	;
	goto L482
L484:
	;
	v1601 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1599))) = uint16(v1601)
	v1604 = v1599 + int32(2)
	v1606 = base.B2i32(v1596 == v1601)
	if v1606 == v1601 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	base.MemoryCopy(m, v1604, v1100, v1596)
	goto L487
L486:
	;
	goto L487
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v1604
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = v1599
	v1612 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v1612
	v1614 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v1614
	v1618 = F_log(m, base.F64_abs(base.F64_convert_i32_s(v1177)))
	mBase = m.M
	v1623 = v1177 >> (uint(int32(31)) % 32)
	v1625 = v1177 ^ v1623 - v1623
	if v1625&int32(1) != 0 {
		goto L489
	} else {
		goto L490
	}
L488:
	;
	v1664 = base.I32_trunc_sat_f64_s(v1618) + v1387 + v1395 + int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1661
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1660
	v1670 = v1625
	goto L506
L489:
	;
	v1628 = F_palloc(m, v1598)
	mBase = m.M
	v1629 = m.ExcPending
	if v1629 != 0 {
		goto L1
	} else {
		goto L492
	}
L490:
	;
	goto L491
L491:
	;
	v1645 = F_palloc(m, int32(4))
	mBase = m.M
	v1646 = m.ExcPending
	if v1646 != 0 {
		goto L1
	} else {
		goto L500
	}
L492:
	;
	v1630 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1628))) = uint16(v1630)
	v1633 = v1628 + int32(2)
	if v1606 == v1630 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	base.MemoryCopy(m, v1633, v1100, v1596)
	goto L495
L494:
	;
	goto L495
L495:
	;
	v1637 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1637 != 0 {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	F_pfree(m, v1637)
	mBase = m.M
	v1639 = m.ExcPending
	if v1639 != 0 {
		goto L1
	} else {
		goto L499
	}
L497:
	;
	goto L498
L498:
	;
	v1640 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v1640
	v1642 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1642
	v1660 = v1628
	v1661 = v1633
	goto L488
L499:
	;
	goto L498
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1645))) = int32(_a_F_numeric_power_20)
	v1649 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1649 != 0 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	F_pfree(m, v1649)
	mBase = m.M
	v1651 = m.ExcPending
	if v1651 != 0 {
		goto L1
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	v1655 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_power[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v1655
	v1658 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_power[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1658
	v1660 = v1645
	v1661 = v1645 + int32(2)
	goto L488
L504:
	;
	goto L503
L505:
	;
	v1746 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
	if v1746 != 0 {
		goto L537
	} else {
		goto L538
	}
L506:
	;
	v1683 = int32(base.Ui32(v1670) >> (uint(int32(1)) % 32))
	if v1683 == int32(0) {
		goto L505
	} else {
		goto L508
	}
L507:
	;
	if int64(0) <= v1172 {
		goto L9
	} else {
		goto L530
	}
L508:
	;
	v1687 = v18 + int32(120)
	v1688 = *(*int32)(unsafe.Add(mBase, uint32(v18)+124))
	v1691 = v1664 - v1688<<(uint(int32(3))%32)
	v1692 = *(*int32)(unsafe.Add(mBase, uint32(v18)+132))
	v1694 = v1692 << (uint(int32(1)) % 32)
	if v1691 < v1694 {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v1696 = v1691
	goto L511
L510:
	;
	v1696 = v1694
	goto L511
L511:
	;
	v1697 = int32(0)
	if v1697 < v1696 {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v1700 = v1696
	goto L514
L513:
	;
	v1700 = v1697
	goto L514
L514:
	;
	F_mul_var(m, v1687, v1687, v1687, v1700)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	v1703 = *(*int32)(unsafe.Add(mBase, uint32(v18)+124))
	if v1670&int32(2) != 0 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v1706 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v1710 = v1664 - (v1706+v1703)<<(uint(int32(2))%32)
	v1711 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1712 = *(*int32)(unsafe.Add(mBase, uint32(v18)+132))
	v1713 = v1711 + v1712
	if v1710 < v1713 {
		goto L519
	} else {
		goto L520
	}
L517:
	;
	goto L518
L518:
	;
	if v1703 <= int32(_a_F_numeric_power_26) {
		goto L526
	} else {
		goto L527
	}
L519:
	;
	v1715 = v1710
	goto L521
L520:
	;
	v1715 = v1713
	goto L521
L521:
	;
	v1716 = int32(0)
	if v1716 < v1715 {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v1719 = v1715
	goto L524
L523:
	;
	v1719 = v1716
	goto L524
L524:
	;
	F_mul_var(m, v1687, v18, v18, v1719)
	mBase = m.M
	v1721 = m.ExcPending
	if v1721 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	goto L518
L526:
	;
	v1726 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v1726 < int32(_a_F_numeric_power_6) {
		v1670 = v1683
		goto L506
	} else {
		goto L529
	}
L527:
	;
	goto L528
L528:
	;
	goto L507
L529:
	;
	goto L528
L530:
	;
	v1732 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1732 != 0 {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	F_pfree(m, v1732)
	mBase = m.M
	v1734 = m.ExcPending
	if v1734 != 0 {
		goto L1
	} else {
		goto L534
	}
L532:
	;
	goto L533
L533:
	;
	v1735 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1735
	v1737 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1737
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v1737
	v1741 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
	if v1741 == v1735 {
		goto L6
	} else {
		goto L535
	}
L534:
	;
	goto L533
L535:
	;
	F_pfree(m, v1741)
	mBase = m.M
	v1745 = m.ExcPending
	if v1745 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	goto L6
L537:
	;
	F_pfree(m, v1746)
	mBase = m.M
	v1748 = m.ExcPending
	if v1748 != 0 {
		goto L1
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	if int64(0) <= v1172 {
		goto L6
	} else {
		goto L541
	}
L540:
	;
	goto L539
L541:
	;
	F_div_var(m, int32(_a_F_numeric_power_3), v18, v18, v1395, int32(1), int32(0))
	mBase = m.M
	v1755 = m.ExcPending
	if v1755 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	goto L5
L543:
	;
	v1760 = F_palloc(m, int32(2))
	mBase = m.M
	v1761 = m.ExcPending
	if v1761 != 0 {
		goto L1
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	v1779 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v18)+80)) = v1779
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v1779
	if v1083 != int32(_a_F_numeric_power_7) {
		goto L552
	} else {
		goto L553
	}
L546:
	;
	v1762 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1760))) = uint16(v1762)
	v1764 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1764 != 0 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	F_pfree(m, v1764)
	mBase = m.M
	v1766 = m.ExcPending
	if v1766 != 0 {
		goto L1
	} else {
		goto L550
	}
L548:
	;
	goto L549
L549:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = int32(16)
	v1770 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1770
	v1773 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_power[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1773
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1760
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1760 + int32(2)
	goto L5
L550:
	;
	goto L549
L551:
	;
	v1845 = v18 + int32(96)
	v1847 = F_estimate_ln_dweight(m, v1843)
	mBase = m.M
	v1848 = m.ExcPending
	if v1848 != 0 {
		goto L1
	} else {
		goto L565
	}
L552:
	;
	v1799 = int32(0)
	v1839 = v1799
	v1842 = v1799
	v1843 = v18 + int32(48)
	goto L551
L553:
	;
	goto L554
L554:
	;
	if v1114 == int32(0) {
		goto L556
	} else {
		goto L557
	}
L555:
	;
	v1819 = v1055 & int32(-2)
	v1822 = F_palloc(m, v1819+int32(2))
	mBase = m.M
	v1823 = m.ExcPending
	if v1823 != 0 {
		goto L1
	} else {
		goto L561
	}
L556:
	;
	v1817 = int32(0)
	goto L555
L557:
	;
	if v1161 < v1114 {
		goto L8
	} else {
		goto L558
	}
L558:
	;
	if v1161 != v1114 {
		goto L556
	} else {
		goto L559
	}
L559:
	;
	v1807 = int32(1)
	v1813 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1157+v1114<<(uint(v1807)%32)-int32(2)))))
	if v1813&v1807 != 0 {
		v1817 = v1807
		goto L555
	} else {
		goto L560
	}
L560:
	;
	goto L556
L561:
	;
	v1824 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1822))) = uint16(v1824)
	v1827 = v1822 + int32(2)
	if v1819 != 0 {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	base.MemoryCopy(m, v1827, v1100, v1819)
	goto L564
L563:
	;
	goto L564
L564:
	;
	v1829 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v1829
	*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v1827
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = v1822
	v1833 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v1833
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = int32(0)
	v1839 = v1822
	v1842 = v1817
	v1843 = v18 + int32(120)
	goto L551
L565:
	;
	v1849 = int32(8) - v1847
	v1850 = int32(0)
	if v1850 < v1849 {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v1853 = v1849
	goto L568
L567:
	;
	v1853 = v1850
	goto L568
L568:
	;
	F_ln_var(m, v1843, v1845, v1853)
	mBase = m.M
	v1855 = m.ExcPending
	if v1855 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	v1859 = v18 + int32(72)
	F_mul_var(m, v1845, v18+int32(24), v1859, v1853)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	v1862 = F_numericvar_to_double_no_overflow(m, v1859)
	mBase = m.M
	v1863 = m.ExcPending
	if v1863 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	if base.F64_gt(base.F64_abs(v1862), float64(6020)) != 0 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	if base.F64_gt(v1862, float64(0)) != 0 {
		goto L7
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v1879 = v18 + int32(96)
	v1884 = base.I32_trunc_sat_f64_s(base.F64_mul(v1862, float64(0.434294481903252)))
	v1885 = int32(16) - v1884
	v1886 = *(*int32)(unsafe.Add(mBase, uint32(v1843)+12))
	if v1886 < v1885 {
		goto L580
	} else {
		goto L581
	}
L575:
	;
	v1869 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1869 != 0 {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	F_pfree(m, v1869)
	mBase = m.M
	v1871 = m.ExcPending
	if v1871 != 0 {
		goto L1
	} else {
		goto L579
	}
L577:
	;
	goto L578
L578:
	;
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = int64(4294967296000)
	v1874 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1874
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v1874
	goto L5
L579:
	;
	goto L578
L580:
	;
	v1888 = v1885
	goto L582
L581:
	;
	v1888 = v1886
	goto L582
L582:
	;
	if v1152 < v1888 {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v1890 = v1888
	goto L585
L584:
	;
	v1890 = v1152
	goto L585
L585:
	;
	if int32(1000) <= v1890 {
		goto L586
	} else {
		goto L587
	}
L586:
	;
	v1893 = int32(1000)
	goto L588
L587:
	;
	v1893 = v1890
	goto L588
L588:
	;
	v1894 = v1893 + v1884
	v1895 = int32(0)
	if v1895 < v1894 {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v1898 = v1894
	goto L591
L590:
	;
	v1898 = v1895
	goto L591
L591:
	;
	v1901 = v1898 - v1847 + int32(8)
	v1902 = int32(0)
	if v1902 < v1901 {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	v1905 = v1901
	goto L594
L593:
	;
	v1905 = v1902
	goto L594
L594:
	;
	F_ln_var(m, v1843, v1879, v1905)
	mBase = m.M
	v1907 = m.ExcPending
	if v1907 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	v1911 = v18 + int32(72)
	F_mul_var(m, v1879, v18+int32(24), v1911, v1905)
	mBase = m.M
	v1913 = m.ExcPending
	if v1913 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	F_exp_var(m, v1911, v18, v1893)
	mBase = m.M
	v1915 = m.ExcPending
	if v1915 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	v1916 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if base.B2i32(int32(0) < v1916)&v1842 != 0 {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(_a_F_numeric_power_7)
	goto L600
L599:
	;
	goto L600
L600:
	;
	v1922 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if v1922 != 0 {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	F_pfree(m, v1922)
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L1
	} else {
		goto L604
	}
L602:
	;
	goto L603
L603:
	;
	v1925 = *(*int32)(unsafe.Add(mBase, uint32(v18)+112))
	if v1925 != 0 {
		goto L605
	} else {
		goto L606
	}
L604:
	;
	goto L603
L605:
	;
	F_pfree(m, v1925)
	mBase = m.M
	v1927 = m.ExcPending
	if v1927 != 0 {
		goto L1
	} else {
		goto L608
	}
L606:
	;
	goto L607
L607:
	;
	if v1839 == int32(0) {
		goto L5
	} else {
		goto L609
	}
L608:
	;
	goto L607
L609:
	;
	F_pfree(m, v1839)
	mBase = m.M
	v1931 = m.ExcPending
	if v1931 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	goto L5
L611:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v1938 = m.ExcPending
	if v1938 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	F_errmsg(m, int32(_a_F_numeric_power_27), int32(0))
	mBase = m.M
	v1942 = m.ExcPending
	if v1942 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	F_errfinish(m, int32(_a_F_numeric_power_9), int32(3973), int32(_a_F_numeric_power_10))
	mBase = m.M
	v1947 = m.ExcPending
	if v1947 != 0 {
		goto L1
	} else {
		goto L614
	}
L614:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L615:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1954 = m.ExcPending
	if v1954 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	F_errmsg(m, int32(_a_F_numeric_power_28), int32(0))
	mBase = m.M
	v1958 = m.ExcPending
	if v1958 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	F_errfinish(m, int32(_a_F_numeric_power_9), int32(_a_F_numeric_power_29), int32(_a_F_numeric_power_30))
	mBase = m.M
	v1963 = m.ExcPending
	if v1963 != 0 {
		goto L1
	} else {
		goto L618
	}
L618:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L619:
	;
	F_errcode(m, int32(33816706))
	mBase = m.M
	v1970 = m.ExcPending
	if v1970 != 0 {
		goto L1
	} else {
		goto L620
	}
L620:
	;
	F_errmsg(m, int32(_a_F_numeric_power_31), int32(0))
	mBase = m.M
	v1974 = m.ExcPending
	if v1974 != 0 {
		goto L1
	} else {
		goto L621
	}
L621:
	;
	F_errfinish(m, int32(_a_F_numeric_power_9), int32(_a_F_numeric_power_32), int32(_a_F_numeric_power_30))
	mBase = m.M
	v1979 = m.ExcPending
	if v1979 != 0 {
		goto L1
	} else {
		goto L622
	}
L622:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L623:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v1986 = m.ExcPending
	if v1986 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	F_errmsg(m, int32(_a_F_numeric_power_28), int32(0))
	mBase = m.M
	v1990 = m.ExcPending
	if v1990 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	F_errfinish(m, int32(_a_F_numeric_power_9), int32(_a_F_numeric_power_33), int32(_a_F_numeric_power_30))
	mBase = m.M
	v1995 = m.ExcPending
	if v1995 != 0 {
		goto L1
	} else {
		goto L626
	}
L626:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L627:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v2002 = m.ExcPending
	if v2002 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	F_errmsg(m, int32(_a_F_numeric_power_27), int32(0))
	mBase = m.M
	v2006 = m.ExcPending
	if v2006 != 0 {
		goto L1
	} else {
		goto L629
	}
L629:
	;
	F_errfinish(m, int32(_a_F_numeric_power_9), int32(_a_F_numeric_power_34), int32(_a_F_numeric_power_35))
	mBase = m.M
	v2011 = m.ExcPending
	if v2011 != 0 {
		goto L1
	} else {
		goto L630
	}
L630:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L631:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v2018 = m.ExcPending
	if v2018 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	F_errmsg(m, int32(_a_F_numeric_power_28), int32(0))
	mBase = m.M
	v2022 = m.ExcPending
	if v2022 != 0 {
		goto L1
	} else {
		goto L633
	}
L633:
	;
	F_errfinish(m, int32(_a_F_numeric_power_9), int32(_a_F_numeric_power_36), int32(_a_F_numeric_power_35))
	mBase = m.M
	v2027 = m.ExcPending
	if v2027 != 0 {
		goto L1
	} else {
		goto L634
	}
L634:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L635:
	;
	goto L5
L636:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(0)
	goto L635
L637:
	;
	goto L638
L638:
	;
	v2051 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v2053 = v1395 & int32(3)
	v2057 = base.I32_div_s(v2042+int32(7), int32(4))
	v2058 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v2058 <= v2057 {
		goto L643
	} else {
		goto L644
	}
L639:
	;
	goto L635
L640:
	;
	if int32(0) <= v2123 {
		goto L639
	} else {
		goto L660
	}
L641:
	;
	v2103 = v2097
	goto L654
L642:
	;
	v2072 = int32(1)
	v2073 = v2057 - v2072
	v2076 = v2051 + v2073<<(uint(v2072)%32)
	v2077 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2076))))
	v2078 = int32(2)
	v2080 = *(*int32)(unsafe.Add(mBase, uint32(v2053<<(uint(v2078)%32))+uint32(_c_F_numeric_power[9])))
	v2081 = base.I32_rem_s(v2077, v2080)
	v2082 = v2077 - v2081
	*(*uint16)(unsafe.Add(mBase, uint32(v2076))) = uint16(v2082)
	v2085 = base.I32_div_s(v2080, v2078)
	if v2081 < v2085 {
		v2123 = v2073
		goto L640
	} else {
		goto L649
	}
L643:
	;
	if base.B2i32(v2053 == int32(0))|base.B2i32(v2057 != v2058) != 0 {
		goto L639
	} else {
		goto L646
	}
L644:
	;
	goto L645
L645:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v2057
	if v2053 != 0 {
		goto L642
	} else {
		goto L647
	}
L646:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v2057
	goto L642
L647:
	;
	v2069 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2051+v2057<<(uint(int32(1))%32)))))
	if v2069 <= int32(_a_F_numeric_power_21) {
		v2123 = v2057
		goto L640
	} else {
		goto L648
	}
L648:
	;
	v2097 = v2057
	goto L641
L649:
	;
	v2088 = v2080 + base.I32_extend16_s(v2082)
	if int32(_a_F_numeric_power_22) < v2088 {
		goto L650
	} else {
		goto L651
	}
L650:
	;
	v2093 = v2088 + int32(_a_F_numeric_power_23)
	goto L652
L651:
	;
	v2093 = v2088
	goto L652
L652:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2076))) = uint16(v2093)
	if v2088 < int32(_a_F_numeric_power_24) {
		v2123 = v2073
		goto L640
	} else {
		goto L653
	}
L653:
	;
	v2097 = v2073
	goto L641
L654:
	;
	v2109 = int32(1)
	v2110 = v2103 - v2109
	v2113 = v2051 + v2110<<(uint(v2109)%32)
	v2116 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2113))))
	v2118 = base.B2i32(int32(_a_F_numeric_power_25) < v2116)
	if int32(_a_F_numeric_power_25) < v2116 {
		goto L656
	} else {
		goto L657
	}
L655:
	;
	v2123 = v2110
	goto L640
L656:
	;
	v2119 = int32(-9999)
	goto L658
L657:
	;
	v2119 = v2109
	goto L658
L658:
	;
	v2120 = v2119 + v2116
	*(*uint16)(unsafe.Add(mBase, uint32(v2113))) = uint16(v2120)
	if int32(_a_F_numeric_power_25) < v2116 {
		v2103 = v2110
		goto L654
	} else {
		goto L659
	}
L659:
	;
	goto L655
L660:
	;
	v2131 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v2131 - int32(2)
	v2135 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v2136 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v2135 + v2136
	v2139 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v2139 + v2136
	goto L639
L661:
	;
	v2175 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v2175 == int32(0) {
		v2183 = v2173
		goto L4
	} else {
		goto L662
	}
L662:
	;
	F_pfree(m, v2175)
	mBase = m.M
	v2179 = m.ExcPending
	if v2179 != 0 {
		goto L1
	} else {
		goto L663
	}
L663:
	;
	v2183 = v2173
	goto L4
}
func F_numeric_random(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v40 int32
	_ = v40
	var v41 int32
	_ = v41
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v53 int32
	_ = v53
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v74 int32
	_ = v74
	var v79 int32
	_ = v79
	var v81 int32
	_ = v81
	var v86 int32
	_ = v86
	var v91 int32
	_ = v91
	var v92 int64
	_ = v92
	var v95 int64
	_ = v95
	var v103 int32
	_ = v103
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v109 int32
	_ = v109
	var v112 int64
	_ = v112
	var v113 int64
	_ = v113
	var v123 int64
	_ = v123
	var v126 int64
	_ = v126
	var v129 int64
	_ = v129
	var v130 int64
	_ = v130
	var v133 int64
	_ = v133
	var v134 int64
	_ = v134
	var v135 int64
	_ = v135
	var v138 int64
	_ = v138
	var v139 int64
	_ = v139
	var v140 int64
	_ = v140
	var v145 int64
	_ = v145
	var v150 int64
	_ = v150
	var v155 int64
	_ = v155
	var v161 int32
	_ = v161
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v167 int32
	_ = v167
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v177 int32
	_ = v177
	var v183 int32
	_ = v183
	var v188 int32
	_ = v188
	var v189 int32
	_ = v189
	var v190 int32
	_ = v190
	var v196 int32
	_ = v196
	var v199 int32
	_ = v199
	var v205 int32
	_ = v205
	var v210 int32
	_ = v210
	var v211 int32
	_ = v211
	var v217 int32
	_ = v217
	var v218 int32
	_ = v218
	var v219 int32
	_ = v219
	var v221 int32
	_ = v221
	var v223 int32
	_ = v223
	var v233 int32
	_ = v233
	var v235 int32
	_ = v235
	var v244 int32
	_ = v244
	var v245 int32
	_ = v245
	var v252 int32
	_ = v252
	var v255 int32
	_ = v255
	var v259 int32
	_ = v259
	var v260 int32
	_ = v260
	var v262 int32
	_ = v262
	var v268 int32
	_ = v268
	var v269 int32
	_ = v269
	var v274 int32
	_ = v274
	var v284 int32
	_ = v284
	var v286 int64
	_ = v286
	var v297 int32
	_ = v297
	var v300 int32
	_ = v300
	var v302 int32
	_ = v302
	var v307 int32
	_ = v307
	var v308 int32
	_ = v308
	var v317 int32
	_ = v317
	var v319 int64
	_ = v319
	var v332 int32
	_ = v332
	var v333 int32
	_ = v333
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	var v342 int32
	_ = v342
	var v345 int32
	_ = v345
	var v346 int32
	_ = v346
	var v347 int32
	_ = v347
	var v359 int32
	_ = v359
	var v361 int64
	_ = v361
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v371 int32
	_ = v371
	var v374 int32
	_ = v374
	var v376 int32
	_ = v376
	var v379 int32
	_ = v379
	var v380 int32
	_ = v380
	var v384 int32
	_ = v384
	var v391 int32
	_ = v391
	var v398 int32
	_ = v398
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v431 int32
	_ = v431
	var v451 int32
	_ = v451
	var v458 int32
	_ = v458
	var v478 int32
	_ = v478
	var v480 int32
	_ = v480
	var v489 int32
	_ = v489
	var v508 int32
	_ = v508
	var v509 int64
	_ = v509
	var v512 int32
	_ = v512
	var v515 int32
	_ = v515
	var v519 int32
	_ = v519
	var v520 int32
	_ = v520
	var v527 int32
	_ = v527
	var v529 int32
	_ = v529
	var v548 int64
	_ = v548
	var v554 int64
	_ = v554
	var v559 int64
	_ = v559
	var v561 int64
	_ = v561
	var v563 int64
	_ = v563
	var v565 int32
	_ = v565
	var v570 int64
	_ = v570
	var v572 int64
	_ = v572
	var v573 int32
	_ = v573
	var v574 int32
	_ = v574
	var v576 int32
	_ = v576
	var v580 int32
	_ = v580
	var v601 int64
	_ = v601
	var v607 int64
	_ = v607
	var v612 int64
	_ = v612
	var v619 int32
	_ = v619
	var v635 int64
	_ = v635
	var v640 int32
	_ = v640
	var v641 int32
	_ = v641
	var v645 int32
	_ = v645
	var v648 int32
	_ = v648
	var v650 int32
	_ = v650
	var v652 int64
	_ = v652
	var v667 int32
	_ = v667
	var v684 int32
	_ = v684
	var v685 int32
	_ = v685
	var v686 int32
	_ = v686
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v690 int64
	_ = v690
	var v691 int64
	_ = v691
	var v698 int64
	_ = v698
	var v700 int64
	_ = v700
	var v701 int64
	_ = v701
	var v704 int64
	_ = v704
	var v706 int64
	_ = v706
	var v710 int64
	_ = v710
	var v712 int64
	_ = v712
	var v720 int64
	_ = v720
	var v725 int64
	_ = v725
	var v738 int64
	_ = v738
	var v740 int32
	_ = v740
	var v741 int64
	_ = v741
	var v748 int64
	_ = v748
	var v750 int64
	_ = v750
	var v751 int64
	_ = v751
	var v754 int64
	_ = v754
	var v756 int64
	_ = v756
	var v760 int64
	_ = v760
	var v762 int64
	_ = v762
	var v770 int64
	_ = v770
	var v775 int64
	_ = v775
	var v788 int64
	_ = v788
	var v789 int64
	_ = v789
	var v792 int32
	_ = v792
	var v793 int64
	_ = v793
	var v794 int64
	_ = v794
	var v797 int64
	_ = v797
	var v801 int32
	_ = v801
	var v804 int64
	_ = v804
	var v811 int64
	_ = v811
	var v813 int64
	_ = v813
	var v818 int64
	_ = v818
	var v820 int64
	_ = v820
	var v823 int32
	_ = v823
	var v825 int32
	_ = v825
	var v853 int32
	_ = v853
	var v854 int32
	_ = v854
	var v863 int64
	_ = v863
	var v865 int64
	_ = v865
	var v866 int64
	_ = v866
	var v869 int64
	_ = v869
	var v871 int64
	_ = v871
	var v875 int64
	_ = v875
	var v877 int64
	_ = v877
	var v885 int64
	_ = v885
	var v890 int64
	_ = v890
	var v894 int64
	_ = v894
	var v905 int64
	_ = v905
	var v908 int64
	_ = v908
	var v909 int64
	_ = v909
	var v910 int64
	_ = v910
	var v913 int64
	_ = v913
	var v915 int64
	_ = v915
	var v919 int64
	_ = v919
	var v922 int32
	_ = v922
	var v924 int32
	_ = v924
	var v951 int32
	_ = v951
	var v980 int32
	_ = v980
	var v989 int64
	_ = v989
	var v991 int64
	_ = v991
	var v992 int64
	_ = v992
	var v995 int64
	_ = v995
	var v997 int64
	_ = v997
	var v1001 int64
	_ = v1001
	var v1003 int64
	_ = v1003
	var v1011 int64
	_ = v1011
	var v1016 int64
	_ = v1016
	var v1020 int64
	_ = v1020
	var v1032 int32
	_ = v1032
	var v1034 int32
	_ = v1034
	var v1061 int32
	_ = v1061
	var v1064 int32
	_ = v1064
	var v1065 int64
	_ = v1065
	var v1067 int32
	_ = v1067
	var v1070 int64
	_ = v1070
	var v1077 int64
	_ = v1077
	var v1079 int64
	_ = v1079
	var v1080 int64
	_ = v1080
	var v1083 int64
	_ = v1083
	var v1085 int64
	_ = v1085
	var v1089 int64
	_ = v1089
	var v1091 int64
	_ = v1091
	var v1099 int64
	_ = v1099
	var v1104 int64
	_ = v1104
	var v1117 int64
	_ = v1117
	var v1118 int64
	_ = v1118
	var v1122 int32
	_ = v1122
	var v1124 int32
	_ = v1124
	var v1128 int32
	_ = v1128
	var v1148 int32
	_ = v1148
	var v1149 int32
	_ = v1149
	var v1160 int32
	_ = v1160
	var v1191 int32
	_ = v1191
	var v1192 int32
	_ = v1192
	var v1200 int32
	_ = v1200
	var v1202 int32
	_ = v1202
	var v1206 int32
	_ = v1206
	var v1226 int32
	_ = v1226
	var v1238 int32
	_ = v1238
	var v1242 int32
	_ = v1242
	var v1248 int32
	_ = v1248
	var v1249 int32
	_ = v1249
	var v1250 int32
	_ = v1250
	var v1252 int32
	_ = v1252
	var v1257 int32
	_ = v1257
	var v1261 int32
	_ = v1261
	var v1274 int32
	_ = v1274
	var v1276 int32
	_ = v1276
	var v1281 int32
	_ = v1281
	var v1282 int32
	_ = v1282
	var v1283 int32
	_ = v1283
	var v1285 int32
	_ = v1285
	var v1293 int32
	_ = v1293
	var v1295 int32
	_ = v1295
	var v1304 int32
	_ = v1304
	var v1305 int32
	_ = v1305
	var v1310 int32
	_ = v1310
	var v1319 int32
	_ = v1319
	var v1321 int32
	_ = v1321
	var v1328 int32
	_ = v1328
	var v1335 int32
	_ = v1335
	var v1336 int32
	_ = v1336
	var v1339 int32
	_ = v1339
	var v1346 int32
	_ = v1346
	var v1351 int32
	_ = v1351
	var v1359 int32
	_ = v1359
	var v1363 int32
	_ = v1363
	var v1368 int32
	_ = v1368
	var v1372 int32
	_ = v1372
	var v1378 int32
	_ = v1378
	var v1389 int32
	_ = v1389
	var v1399 int32
	_ = v1399
	var v1413 int32
	_ = v1413
	var v1471 int32
	_ = v1471
	var v1498 int32
	_ = v1498
	var v1500 int32
	_ = v1500
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1504 int32
	_ = v1504
	var v1506 int32
	_ = v1506
	var v1513 int32
	_ = v1513
	var v1516 int32
	_ = v1516
	var v1520 int32
	_ = v1520
	var v1525 int32
	_ = v1525
	var v1529 int32
	_ = v1529
	var v1534 int32
	_ = v1534
	var v1538 int32
	_ = v1538
	var v1543 int32
	_ = v1543
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v28 = F_pg_detoast_datum(m, v27)
	mBase = m.M
	v31 = m.ExcPending
	if v31 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return int64(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+40))
	v33 = F_pg_detoast_datum(m, v32)
	mBase = m.M
	v34 = m.ExcPending
	if v34 != 0 {
		goto L1
	} else {
		goto L3
	}
L3:
	;
	v36 = int32(*(*uint8)(unsafe.Add(mBase, _c_F_numeric_random[0])))
	if v36 == int32(0) {
		goto L4
	} else {
		goto L5
	}
L4:
	;
	v40 = int32(16)
	v41 = int32(0)
	v45 = m.G0
	v47 = v45 - v40
	m.G0 = v47
	*(*int32)(unsafe.Add(mBase, uint32(v47))) = v41
	v53 = F_open(m, int32(_a_F_numeric_random_0), v41, v47)
	mBase = m.M
	if v53 != int32(-1) {
		goto L9
	} else {
		goto L10
	}
L5:
	;
	goto L6
L6:
	;
	v163 = m.G0
	v165 = v163 - int32(96)
	m.G0 = v165
	v167 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v28)+4)))
	v168 = base.I32_extend16_s(v167)
	if base.Ui32(int32(_a_F_numeric_random_1)) <= base.Ui32(v167) {
		goto L33
	} else {
		goto L34
	}
L7:
	;
	v161 = int32(1)
	*(*uint8)(unsafe.Add(mBase, _c_F_numeric_random[0])) = uint8(v161)
	goto L6
L8:
	;
	if v86 != 0 {
		goto L21
	} else {
		goto L22
	}
L9:
	;
	goto L13
L10:
	;
	v86 = v41
	goto L11
L11:
	;
	m.G0 = v47 + int32(16)
	goto L8
L12:
	;
	v81 = F_close(m, v53)
	mBase = m.M
	v86 = v79
	goto L11
L13:
	;
	v59 = int32(_a_F_numeric_random_2)
	v60 = v40
	goto L14
L14:
	;
	v65 = F_read(m, v53, v59, v60)
	mBase = m.M
	if v65 <= int32(0) {
		goto L16
	} else {
		goto L17
	}
L15:
	;
	v79 = int32(1)
	goto L12
L16:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_random[1]))
	if v69 == int32(27) {
		goto L14
	} else {
		goto L19
	}
L17:
	;
	goto L18
L18:
	;
	v74 = v60 - v65
	if v74 != 0 {
		v59 = v59 + v65
		v60 = v74
		goto L14
	} else {
		goto L20
	}
L19:
	;
	v79 = int32(0)
	goto L12
L20:
	;
	goto L15
L21:
	;
	v91 = int32(_a_F_numeric_random_2)
	v92 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2]))
	if v92 != int64(0) {
		goto L25
	} else {
		goto L26
	}
L22:
	;
	goto L23
L23:
	;
	v103 = int32(_a_F_numeric_random_2)
	v107 = m.G0
	v108 = int32(16)
	v109 = v107 - v108
	m.G0 = v109
	F_gettimeofday(m, v109)
	mBase = m.M
	v112 = *(*int64)(unsafe.Add(mBase, uint32(v109)))
	v113 = int64(*(*int32)(unsafe.Add(mBase, uint32(v109)+8)))
	m.G0 = v109 + v108
	goto L29
L24:
	;
	goto L7
L25:
	;
	goto L24
L26:
	;
	v95 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3]))
	if v95 != int64(0) {
		goto L25
	} else {
		goto L27
	}
L27:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3])) = int64(1442695040888963407)
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2])) = int64(6364136223846793005)
	goto L25
L29:
	;
	v123 = int64(*(*uint32)(unsafe.Add(mBase, _c_F_numeric_random[4])))
	v126 = v113 + v112*int64(1000000) - int64(946684800000000) ^ v123<<(uint(int64(32))%64)
	v129 = v126 + int64(4354685564936845354)
	v130 = int64(30)
	v133 = int64(-4658895280553007687)
	v134 = (int64(base.Ui64(v129)>>(uint(v130)%64)) ^ v129) * v133
	v135 = int64(27)
	v138 = int64(-7723592293110705685)
	v139 = (int64(base.Ui64(v134)>>(uint(v135)%64)) ^ v134) * v138
	v140 = int64(31)
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3])) = int64(base.Ui64(v139)>>(uint(v140)%64)) ^ v139
	v145 = v126 - int64(7046029254386353131)
	v150 = (int64(base.Ui64(v145)>>(uint(v130)%64)) ^ v145) * v133
	v155 = (int64(base.Ui64(v150)>>(uint(v135)%64)) ^ v150) * v138
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2])) = int64(base.Ui64(v155)>>(uint(v140)%64)) ^ v155
	goto L30
L30:
	;
	goto L7
L31:
	;
	return base.I64_extend_i32_u(v1502)
L32:
	;
	F_errmsg(m, int32(_a_F_numeric_random_3), int32(0))
	mBase = m.M
	v1538 = m.ExcPending
	if v1538 != 0 {
		goto L1
	} else {
		goto L271
	}
L33:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L1
	} else {
		goto L36
	}
L34:
	;
	goto L35
L35:
	;
	v189 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v33)+4)))
	v190 = base.I32_extend16_s(v189)
	if base.Ui32(int32(_a_F_numeric_random_1)) <= base.Ui32(v189) {
		goto L42
	} else {
		goto L43
	}
L36:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v177 = m.ExcPending
	if v177 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	if v168 == int32(-16384) {
		goto L32
	} else {
		goto L38
	}
L38:
	;
	F_errmsg(m, int32(_a_F_numeric_random_4), int32(0))
	mBase = m.M
	v183 = m.ExcPending
	if v183 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	F_errfinish(m, int32(_a_F_numeric_random_5), int32(_a_F_numeric_random_6), int32(_a_F_numeric_random_7))
	mBase = m.M
	v188 = m.ExcPending
	if v188 != 0 {
		goto L1
	} else {
		goto L40
	}
L40:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L41:
	;
	F_errmsg(m, int32(_a_F_numeric_random_8), int32(0))
	mBase = m.M
	v1529 = m.ExcPending
	if v1529 != 0 {
		goto L1
	} else {
		goto L269
	}
L42:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v196 = m.ExcPending
	if v196 != 0 {
		goto L1
	} else {
		goto L45
	}
L43:
	;
	goto L44
L44:
	;
	v211 = *(*int32)(unsafe.Add(mBase, uint32(v28)))
	v217 = base.B2i32(int32(0) <= v168)
	if int32(0) <= v168 {
		goto L50
	} else {
		goto L51
	}
L45:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v199 = m.ExcPending
	if v199 != 0 {
		goto L1
	} else {
		goto L46
	}
L46:
	;
	if v190 == int32(-16384) {
		goto L41
	} else {
		goto L47
	}
L47:
	;
	F_errmsg(m, int32(_a_F_numeric_random_9), int32(0))
	mBase = m.M
	v205 = m.ExcPending
	if v205 != 0 {
		goto L1
	} else {
		goto L48
	}
L48:
	;
	F_errfinish(m, int32(_a_F_numeric_random_5), int32(_a_F_numeric_random_10), int32(_a_F_numeric_random_7))
	mBase = m.M
	v210 = m.ExcPending
	if v210 != 0 {
		goto L1
	} else {
		goto L49
	}
L49:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L50:
	;
	v218 = int32(-8)
	goto L52
L51:
	;
	v218 = int32(-6)
	goto L52
L52:
	;
	v219 = int32(base.Ui32(v211)>>(uint(int32(2))%32)) + v218
	v221 = int32(base.Ui32(v219) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+48)) = v221
	if int32(0) <= v168 {
		goto L53
	} else {
		goto L54
	}
L53:
	;
	v223 = int32(*(*int16)(unsafe.Add(mBase, uint32(v28)+6)))
	v233 = v223
	goto L55
L54:
	;
	v233 = v167<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v167&int32(63)
	goto L55
L55:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+52)) = v233
	v235 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+64)) = v235
	v244 = base.B2i32(v168 < v235)
	if v168 < v235 {
		goto L56
	} else {
		goto L57
	}
L56:
	;
	v245 = int32(base.Ui32(v167)>>(uint(int32(7))%32)) & int32(63)
	goto L58
L57:
	;
	v245 = v167 & int32(_a_F_numeric_random_11)
	goto L58
L58:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+60)) = v245
	v252 = v167 & int32(_a_F_numeric_random_1)
	if v252 == int32(_a_F_numeric_random_12) {
		goto L59
	} else {
		goto L60
	}
L59:
	;
	v255 = v167 << (uint(int32(1)) % 32) & int32(_a_F_numeric_random_13)
	goto L61
L60:
	;
	v255 = v252
	goto L61
L61:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+56)) = v255
	if v168 < v235 {
		goto L62
	} else {
		goto L63
	}
L62:
	;
	v259 = int32(6)
	goto L64
L63:
	;
	v259 = int32(8)
	goto L64
L64:
	;
	v260 = v28 + v259
	*(*int32)(unsafe.Add(mBase, uint32(v165)+68)) = v260
	v262 = *(*int32)(unsafe.Add(mBase, uint32(v33)))
	v268 = base.B2i32(int32(0) <= v190)
	if int32(0) <= v190 {
		goto L65
	} else {
		goto L66
	}
L65:
	;
	v269 = int32(-8)
	goto L67
L66:
	;
	v269 = int32(-6)
	goto L67
L67:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+24)) = int32(base.Ui32(int32(base.Ui32(v262)>>(uint(int32(2))%32))+v269) >> (uint(int32(1)) % 32))
	if int32(0) <= v190 {
		goto L68
	} else {
		goto L69
	}
L68:
	;
	v274 = int32(*(*int16)(unsafe.Add(mBase, uint32(v33)+6)))
	v284 = v274
	goto L70
L69:
	;
	v284 = v189<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v189&int32(63)
	goto L70
L70:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+28)) = v284
	v286 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = v286
	*(*int64)(unsafe.Add(mBase, uint32(v165)+8)) = v286
	*(*int64)(unsafe.Add(mBase, uint32(v165)+16)) = v286
	v297 = v189 & int32(_a_F_numeric_random_1)
	if v297 == int32(_a_F_numeric_random_12) {
		goto L71
	} else {
		goto L72
	}
L71:
	;
	v300 = v189 << (uint(int32(1)) % 32) & int32(_a_F_numeric_random_13)
	goto L73
L72:
	;
	v300 = v297
	goto L73
L73:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+32)) = v300
	v302 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165)+40)) = v302
	v307 = base.B2i32(v190 < v302)
	if v190 < v302 {
		goto L74
	} else {
		goto L75
	}
L74:
	;
	v308 = int32(6)
	goto L76
L75:
	;
	v308 = int32(8)
	goto L76
L76:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+44)) = v33 + v308
	if v190 < v302 {
		goto L77
	} else {
		goto L78
	}
L77:
	;
	v317 = int32(base.Ui32(v189)>>(uint(int32(7))%32)) & int32(63)
	goto L79
L78:
	;
	v317 = v189 & int32(_a_F_numeric_random_11)
	goto L79
L79:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+36)) = v317
	v319 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v165)+88)) = v319
	*(*int64)(unsafe.Add(mBase, uint32(v165)+80)) = v319
	*(*int64)(unsafe.Add(mBase, uint32(v165)+72)) = v319
	F_sub_var(m, v165+int32(24), v165+int32(48), v165+int32(72))
	mBase = m.M
	v332 = m.ExcPending
	if v332 != 0 {
		goto L1
	} else {
		goto L80
	}
L80:
	;
	v333 = *(*int32)(unsafe.Add(mBase, uint32(v165)+80))
	if v333 != int32(_a_F_numeric_random_13) {
		goto L81
	} else {
		goto L82
	}
L81:
	;
	if base.Ui32(v317) < base.Ui32(v245) {
		goto L84
	} else {
		goto L85
	}
L82:
	;
	goto L83
L83:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1513 = m.ExcPending
	if v1513 != 0 {
		goto L1
	} else {
		goto L265
	}
L84:
	;
	v337 = v245
	goto L86
L85:
	;
	v337 = v317
	goto L86
L86:
	;
	v338 = *(*int32)(unsafe.Add(mBase, uint32(v165)+72))
	if v338 == int32(0) {
		goto L88
	} else {
		goto L89
	}
L87:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v165)+88))
	if v1498 != 0 {
		goto L256
	} else {
		goto L257
	}
L88:
	;
	v342 = v219 & int32(-2)
	v345 = F_palloc(m, v342+int32(2))
	mBase = m.M
	v346 = m.ExcPending
	if v346 != 0 {
		goto L1
	} else {
		goto L91
	}
L89:
	;
	goto L90
L90:
	;
	v368 = int32(1)
	v369 = *(*int32)(unsafe.Add(mBase, uint32(v165)+76))
	v371 = v337 + int32(3)
	v374 = v369 + int32(base.Ui32(v371)>>(uint(int32(2))%32))
	v376 = v374 + v368
	v379 = v371 & int32(_a_F_numeric_random_14)
	v380 = v379 - v337
	if v380 <= int32(0) {
		v489 = v368
		goto L95
	} else {
		goto L96
	}
L91:
	;
	v347 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v345))) = uint16(v347)
	if base.B2i32(v221 == v347)|base.B2i32(v342 == v347) == v347 {
		goto L92
	} else {
		goto L93
	}
L92:
	;
	base.MemoryCopy(m, v345+int32(2), v260, v342)
	goto L94
L93:
	;
	goto L94
L94:
	;
	v359 = *(*int32)(unsafe.Add(mBase, uint32(v165)+56))
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = v359
	v361 = *(*int64)(unsafe.Add(mBase, uint32(v165)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = v361
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v345
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v345 + int32(2)
	goto L87
L95:
	;
	v508 = *(*int32)(unsafe.Add(mBase, uint32(v165)+92))
	v509 = int64(*(*int16)(unsafe.Add(mBase, uint32(v508))))
	if v376 < int32(2) {
		v619 = v368
		v635 = v509
		goto L107
	} else {
		goto L108
	}
L96:
	;
	v384 = v380 & int32(7)
	if base.Ui32(v337-v379) <= base.Ui32(int32(-8)) {
		goto L97
	} else {
		goto L98
	}
L97:
	;
	v391 = int32(0)
	v398 = v368
	goto L100
L98:
	;
	v431 = v368
	goto L99
L99:
	;
	v451 = int32(0)
	v458 = v431
	goto L104
L100:
	;
	v418 = v398 * int32(100000000)
	v420 = v391 + int32(8)
	if v420 != v380&int32(2147483640) {
		v391 = v420
		v398 = v418
		goto L100
	} else {
		goto L102
	}
L101:
	;
	if v384 == int32(0) {
		v489 = v418
		goto L95
	} else {
		goto L103
	}
L102:
	;
	goto L101
L103:
	;
	v431 = v418
	goto L99
L104:
	;
	v478 = v458 * int32(10)
	v480 = v451 + int32(1)
	if v480 != v384 {
		v451 = v480
		v458 = v478
		goto L104
	} else {
		goto L106
	}
L105:
	;
	v489 = v478
	goto L95
L106:
	;
	goto L105
L107:
	;
	v640 = int32(1)
	v641 = base.B2i32(v489 != v640)
	v645 = v376 << (uint(v640) % 32)
	if v489 != v640 {
		goto L126
	} else {
		goto L127
	}
L108:
	;
	v512 = int32(4)
	if base.Ui32(v512) <= base.Ui32(v376) {
		goto L109
	} else {
		goto L110
	}
L109:
	;
	v515 = v512
	goto L111
L110:
	;
	v515 = v376
	goto L111
L111:
	;
	if v376 != int32(2) {
		goto L112
	} else {
		goto L113
	}
L112:
	;
	v519 = int32(1)
	v520 = v515 - v519
	v527 = v519
	v529 = int32(0)
	v548 = v509
	goto L115
L113:
	;
	v580 = int32(1)
	v601 = v509
	goto L114
L114:
	;
	v607 = v601 * int64(10000)
	if v338 <= v580 {
		v619 = v515
		v635 = v607
		goto L107
	} else {
		goto L125
	}
L115:
	;
	v554 = v548 * int64(10000)
	if v527 < v338 {
		goto L117
	} else {
		goto L118
	}
L116:
	;
	if v520&v519 == int32(0) {
		v619 = v515
		v635 = v572
		goto L107
	} else {
		goto L124
	}
L117:
	;
	v559 = int64(*(*int16)(unsafe.Add(mBase, uint32(v508+v527<<(uint(int32(1))%32)))))
	v561 = v554 + v559
	goto L119
L118:
	;
	v561 = v554
	goto L119
L119:
	;
	v563 = v561 * int64(10000)
	v565 = v527 + int32(1)
	if v565 < v338 {
		goto L120
	} else {
		goto L121
	}
L120:
	;
	v570 = int64(*(*int16)(unsafe.Add(mBase, uint32(v508+v565<<(uint(int32(1))%32)))))
	v572 = v563 + v570
	goto L122
L121:
	;
	v572 = v563
	goto L122
L122:
	;
	v573 = int32(2)
	v574 = v527 + v573
	v576 = v529 + v573
	if v576 != v520&int32(-2) {
		v527 = v574
		v529 = v576
		v548 = v572
		goto L115
	} else {
		goto L123
	}
L123:
	;
	goto L116
L124:
	;
	v580 = v574
	v601 = v572
	goto L114
L125:
	;
	v612 = int64(*(*int16)(unsafe.Add(mBase, uint32(v508+v580<<(uint(int32(1))%32)))))
	v619 = v515
	v635 = v607 + v612
	goto L107
L126:
	;
	v648 = v374
	goto L128
L127:
	;
	v648 = v376
	goto L128
L128:
	;
	v650 = v648 - int32(3)
	v652 = base.I64_extend_i32_s(v489)
	v667 = int32(0)
	goto L132
L129:
	;
	F_add_var(m, v165, v165+int32(48), v165)
	mBase = m.M
	v1471 = m.ExcPending
	if v1471 != 0 {
		goto L1
	} else {
		goto L255
	}
L130:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v1413
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v165))) = int64(0)
	goto L129
L131:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v685
	v1413 = v823
	goto L130
L132:
	;
	if v667 != 0 {
		goto L134
	} else {
		goto L135
	}
L133:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+20)) = v1202
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v685
	*(*int32)(unsafe.Add(mBase, uint32(v165)+8)) = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v165))) = v1200
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v165)+4)) = v1206
	goto L129
L134:
	;
	F_pfree(m, v667)
	mBase = m.M
	v684 = m.ExcPending
	if v684 != 0 {
		goto L1
	} else {
		goto L137
	}
L135:
	;
	goto L136
L136:
	;
	v685 = F_palloc(m, v645+int32(2))
	mBase = m.M
	v686 = m.ExcPending
	if v686 != 0 {
		goto L1
	} else {
		goto L138
	}
L137:
	;
	goto L136
L138:
	;
	v687 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v685))) = uint16(v687)
	if v641&base.B2i32(v376 == v619) != 0 {
		goto L140
	} else {
		goto L141
	}
L139:
	;
	v792 = v685 + v619<<(uint(int32(1))%32)
	v793 = int64(10000)
	v794 = base.I64_div_u_s(v789, v793)
	v797 = v789 - v794*v793
	*(*uint16)(unsafe.Add(mBase, uint32(v792))) = uint16(v797)
	if base.Ui32(v619) < base.Ui32(int32(2)) {
		goto L157
	} else {
		goto L158
	}
L140:
	;
	v689 = int32(_a_F_numeric_random_2)
	v690 = int64(0)
	v691 = base.I64_div_u_s(v635, v652)
	if base.Ui64(v691) <= base.Ui64(v690) {
		goto L144
	} else {
		goto L145
	}
L141:
	;
	goto L142
L142:
	;
	v740 = int32(_a_F_numeric_random_2)
	v741 = int64(0)
	if base.Ui64(v635) <= base.Ui64(v741) {
		goto L151
	} else {
		goto L152
	}
L143:
	;
	v789 = v738 * v652
	goto L139
L144:
	;
	v738 = v690
	goto L143
L145:
	;
	goto L146
L146:
	;
	v698 = v691 - v690
	v700 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3]))
	v701 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2]))
	v704 = v701
	v706 = v700
	goto L147
L147:
	;
	v710 = v704 ^ v706
	v712 = base.I64_rotl(v710, int64(37))
	v720 = v710 ^ (v710<<(uint(int64(16))%64) ^ base.I64_rotl(v704, int64(24)))
	v725 = int64(base.Ui64(base.I64_rotl(v704*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v698)) % 64))
	if base.Ui64(v698) < base.Ui64(v725) {
		v704 = v720
		v706 = v712
		goto L147
	} else {
		goto L149
	}
L148:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3])) = v712
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2])) = v720
	v738 = v690 + v725
	goto L143
L149:
	;
	goto L148
L150:
	;
	v789 = v788
	goto L139
L151:
	;
	v788 = v741
	goto L150
L152:
	;
	goto L153
L153:
	;
	v748 = v635 - v741
	v750 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3]))
	v751 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2]))
	v754 = v751
	v756 = v750
	goto L154
L154:
	;
	v760 = v754 ^ v756
	v762 = base.I64_rotl(v760, int64(37))
	v770 = v760 ^ (v760<<(uint(int64(16))%64) ^ base.I64_rotl(v754, int64(24)))
	v775 = int64(base.Ui64(base.I64_rotl(v754*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v748)) % 64))
	if base.Ui64(v748) < base.Ui64(v775) {
		v754 = v770
		v756 = v762
		goto L154
	} else {
		goto L156
	}
L155:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3])) = v762
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2])) = v770
	v788 = v741 + v775
	goto L150
L156:
	;
	goto L155
L157:
	;
	v823 = v685 + int32(2)
	if v619 < v650 {
		goto L161
	} else {
		goto L162
	}
L158:
	;
	v801 = int32(2)
	v804 = base.I64_rem_u_s(v794, int64(10000))
	*(*uint16)(unsafe.Add(mBase, uint32(v792-v801))) = uint16(v804)
	if v619 == v801 {
		goto L157
	} else {
		goto L159
	}
L159:
	;
	v811 = base.I64_div_u_s(v789, int64(100000000))
	v813 = base.I64_rem_u_s(v811, int64(10000))
	*(*uint16)(unsafe.Add(mBase, uint32(v792-int32(4)))) = uint16(v813)
	if base.Ui32(v619) < base.Ui32(int32(4)) {
		goto L157
	} else {
		goto L160
	}
L160:
	;
	v818 = base.I64_div_u_s(v789, int64(1000000000000))
	v820 = base.I64_rem_u_s(v818, int64(10000))
	*(*uint16)(unsafe.Add(mBase, uint32(v792-int32(6)))) = uint16(v820)
	goto L157
L161:
	;
	v825 = v619
	goto L164
L162:
	;
	v924 = v619
	goto L163
L163:
	;
	if v924 < v648 {
		goto L174
	} else {
		goto L175
	}
L164:
	;
	v853 = v823 + v825<<(uint(int32(1))%32)
	v854 = int32(_a_F_numeric_random_2)
	goto L168
L165:
	;
	v924 = v922
	goto L163
L166:
	;
	v905 = base.I64_div_u_s(v894, int64(1000000000000))
	*(*uint16)(unsafe.Add(mBase, uint32(v853)+6)) = uint16(v905)
	v908 = base.I64_div_u_s(v894, int64(100000000))
	v909 = int64(10000)
	v910 = base.I64_rem_u_s(v908, v909)
	*(*uint16)(unsafe.Add(mBase, uint32(v853)+4)) = uint16(v910)
	v913 = base.I64_div_u_s(v894, v909)
	v915 = base.I64_rem_u_s(v913, v909)
	*(*uint16)(unsafe.Add(mBase, uint32(v853)+2)) = uint16(v915)
	v919 = v894 - v913*v909
	*(*uint16)(unsafe.Add(mBase, uint32(v853))) = uint16(v919)
	v922 = v825 + int32(4)
	if v922 < v650 {
		v825 = v922
		goto L164
	} else {
		goto L173
	}
L168:
	;
	goto L169
L169:
	;
	v863 = int64(9999999999999999)
	v865 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3]))
	v866 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2]))
	v869 = v866
	v871 = v865
	goto L170
L170:
	;
	v875 = v869 ^ v871
	v877 = base.I64_rotl(v875, int64(37))
	v885 = v875 ^ (v875<<(uint(int64(16))%64) ^ base.I64_rotl(v869, int64(24)))
	v890 = int64(base.Ui64(base.I64_rotl(v869*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v863)) % 64))
	if base.Ui64(v863) < base.Ui64(v890) {
		v869 = v885
		v871 = v877
		goto L170
	} else {
		goto L172
	}
L171:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3])) = v877
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2])) = v885
	v894 = int64(0) + v890
	goto L166
L172:
	;
	goto L171
L173:
	;
	goto L165
L174:
	;
	v951 = v924
	goto L177
L175:
	;
	v1034 = v924
	goto L176
L176:
	;
	if v1034 < v376 {
		goto L187
	} else {
		goto L188
	}
L177:
	;
	v980 = int32(_a_F_numeric_random_2)
	goto L181
L178:
	;
	v1034 = v648
	goto L176
L179:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v823+v951<<(uint(int32(1))%32)))) = uint16(v1020)
	v1032 = v951 + int32(1)
	if v1032 != v648 {
		v951 = v1032
		goto L177
	} else {
		goto L186
	}
L181:
	;
	goto L182
L182:
	;
	v989 = int64(9999)
	v991 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3]))
	v992 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2]))
	v995 = v992
	v997 = v991
	goto L183
L183:
	;
	v1001 = v995 ^ v997
	v1003 = base.I64_rotl(v1001, int64(37))
	v1011 = v1001 ^ (v1001<<(uint(int64(16))%64) ^ base.I64_rotl(v995, int64(24)))
	v1016 = int64(base.Ui64(base.I64_rotl(v995*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v989)) % 64))
	if base.Ui64(v989) < base.Ui64(v1016) {
		v995 = v1011
		v997 = v1003
		goto L183
	} else {
		goto L185
	}
L184:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3])) = v1003
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2])) = v1011
	v1020 = int64(0) + v1016
	goto L179
L185:
	;
	goto L184
L186:
	;
	goto L178
L187:
	;
	v1061 = int32(1)
	v1064 = int32(_a_F_numeric_random_2)
	v1065 = int64(0)
	v1067 = base.I32_div_s(int32(_a_F_numeric_random_15), v489)
	v1070 = base.I64_extend_i32_s(v1067 - v1061)
	if base.Ui64(v1070) <= base.Ui64(v1065) {
		goto L191
	} else {
		goto L192
	}
L188:
	;
	goto L189
L189:
	;
	if base.B2i32(base.Ui32(int32(2147483646)) < base.Ui32(v374)) == int32(0) {
		goto L198
	} else {
		goto L199
	}
L190:
	;
	v1118 = v1117 * base.I64_extend_i32_u(v489)
	*(*uint16)(unsafe.Add(mBase, uint32(v823+v1034<<(uint(v1061)%32)))) = uint16(v1118)
	goto L189
L191:
	;
	v1117 = v1065
	goto L190
L192:
	;
	goto L193
L193:
	;
	v1077 = v1070 - v1065
	v1079 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3]))
	v1080 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2]))
	v1083 = v1080
	v1085 = v1079
	goto L194
L194:
	;
	v1089 = v1083 ^ v1085
	v1091 = base.I64_rotl(v1089, int64(37))
	v1099 = v1089 ^ (v1089<<(uint(int64(16))%64) ^ base.I64_rotl(v1083, int64(24)))
	v1104 = int64(base.Ui64(base.I64_rotl(v1083*int64(5), int64(7))*int64(9)) >> (uint(base.I64_clz(v1077)) % 64))
	if base.Ui64(v1077) < base.Ui64(v1104) {
		v1083 = v1099
		v1085 = v1091
		goto L194
	} else {
		goto L196
	}
L195:
	;
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[3])) = v1091
	*(*int64)(unsafe.Add(mBase, _c_F_numeric_random[2])) = v1099
	v1117 = v1065 + v1104
	goto L190
L196:
	;
	goto L195
L197:
	;
	v1226 = int32(0)
	if base.B2i32(v369 < v1206)&base.B2i32(v1226 < v1200) == v1226 {
		v1257 = v1206
		v1261 = v1226
		goto L213
	} else {
		goto L214
	}
L198:
	;
	v1122 = v376
	v1124 = v823
	v1128 = v369
	goto L202
L199:
	;
	goto L200
L200:
	;
	if v376 == int32(0) {
		goto L131
	} else {
		goto L210
	}
L201:
	;
	v1160 = v1122
	goto L206
L202:
	;
	v1148 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1124))))
	if v1148 != 0 {
		goto L201
	} else {
		goto L204
	}
L203:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v685
	v1413 = v823 + v645
	goto L130
L204:
	;
	v1149 = int32(1)
	if v1149 < v1122 {
		v1122 = v1122 - v1149
		v1124 = v1124 + int32(2)
		v1128 = v1128 - v1149
		goto L202
	} else {
		goto L205
	}
L205:
	;
	goto L203
L206:
	;
	v1191 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1124+v1160<<(uint(int32(1))%32)-int32(2)))))
	if v1191 != 0 {
		v1200 = v1160
		v1202 = v1124
		v1206 = v1128
		goto L197
	} else {
		goto L208
	}
L207:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v165)+12)) = v337
	*(*int32)(unsafe.Add(mBase, uint32(v165)+16)) = v685
	v1413 = v1124
	goto L130
L208:
	;
	v1192 = int32(1)
	if v1192 < v1160 {
		v1160 = v1160 - v1192
		goto L206
	} else {
		goto L209
	}
L209:
	;
	goto L207
L210:
	;
	v1200 = v376
	v1202 = v823
	v1206 = v369
	goto L197
L211:
	;
	if int32(0) < v1399 {
		v667 = v685
		goto L132
	} else {
		goto L254
	}
L212:
	;
	v1399 = v1389
	goto L211
L213:
	;
	if base.B2i32(v338 <= int32(0))|base.B2i32(v369 <= v1257) != 0 {
		v1293 = v369
		v1295 = v1226
		goto L220
	} else {
		goto L221
	}
L214:
	;
	v1238 = v1206
	v1242 = v1226
	goto L215
L215:
	;
	v1248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1202+v1242<<(uint(int32(1))%32)))))
	if v1248 != 0 {
		v1389 = int32(1)
		goto L212
	} else {
		goto L217
	}
L216:
	;
	v1257 = v1252
	v1261 = v1250
	goto L213
L217:
	;
	v1249 = int32(1)
	v1250 = v1242 + v1249
	v1252 = v1238 - v1249
	if v1252 <= v369 {
		v1257 = v1252
		v1261 = v1250
		goto L213
	} else {
		goto L218
	}
L218:
	;
	if v1250 < v1200 {
		v1238 = v1252
		v1242 = v1250
		goto L215
	} else {
		goto L219
	}
L219:
	;
	goto L216
L220:
	;
	if v1257 != v1293 {
		v1335 = v1261
		v1336 = v1295
		goto L227
	} else {
		goto L228
	}
L221:
	;
	v1274 = v369
	v1276 = v1226
	goto L222
L222:
	;
	v1281 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v508+v1276<<(uint(int32(1))%32)))))
	if v1281 != 0 {
		v1389 = int32(-1)
		goto L212
	} else {
		goto L224
	}
L223:
	;
	v1293 = v1285
	v1295 = v1283
	goto L220
L224:
	;
	v1282 = int32(1)
	v1283 = v1276 + v1282
	v1285 = v1274 - v1282
	if v1285 <= v1257 {
		v1293 = v1285
		v1295 = v1283
		goto L220
	} else {
		goto L225
	}
L225:
	;
	if v1283 < v338 {
		v1274 = v1285
		v1276 = v1283
		goto L222
	} else {
		goto L226
	}
L226:
	;
	goto L223
L227:
	;
	if v1200 < v1335 {
		goto L236
	} else {
		goto L237
	}
L228:
	;
	v1304 = v1261
	v1305 = v1295
	goto L229
L229:
	;
	if base.B2i32(v1200 <= v1304)|base.B2i32(v338 <= v1305) != 0 {
		v1335 = v1304
		v1336 = v1305
		goto L227
	} else {
		goto L231
	}
L230:
	;
	if base.I32_extend16_s(v1321) < base.I32_extend16_s(v1319) {
		goto L233
	} else {
		goto L234
	}
L231:
	;
	v1310 = int32(1)
	v1319 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1202+v1304<<(uint(v1310)%32)))))
	v1321 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1305<<(uint(v1310)%32)+v508))))
	if v1319 == v1321 {
		v1304 = v1304 + v1310
		v1305 = v1305 + v1310
		goto L229
	} else {
		goto L232
	}
L232:
	;
	goto L230
L233:
	;
	v1328 = int32(1)
	goto L235
L234:
	;
	v1328 = int32(-1)
	goto L235
L235:
	;
	v1399 = v1328
	goto L211
L236:
	;
	v1339 = v1335
	goto L238
L237:
	;
	v1339 = v1200
	goto L238
L238:
	;
	v1346 = v1335
	goto L239
L239:
	;
	if v1339 == v1346 {
		goto L241
	} else {
		goto L242
	}
L240:
	;
	v1389 = v1372
	goto L212
L241:
	;
	if v338 < v1336 {
		goto L244
	} else {
		goto L245
	}
L242:
	;
	goto L243
L243:
	;
	v1372 = int32(1)
	v1378 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1202+v1346<<(uint(v1372)%32)))))
	if v1378 == int32(0) {
		v1346 = v1346 + v1372
		goto L239
	} else {
		goto L253
	}
L244:
	;
	v1351 = v1336
	goto L246
L245:
	;
	v1351 = v338
	goto L246
L246:
	;
	v1359 = v1336
	goto L247
L247:
	;
	if v1351 == v1359 {
		goto L249
	} else {
		goto L250
	}
L248:
	;
	v1389 = int32(-1)
	goto L212
L249:
	;
	v1399 = int32(0)
	goto L211
L250:
	;
	goto L251
L251:
	;
	v1363 = int32(1)
	v1368 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v1359<<(uint(v1363)%32)+v508))))
	if v1368 == int32(0) {
		v1359 = v1359 + v1363
		goto L247
	} else {
		goto L252
	}
L252:
	;
	goto L248
L253:
	;
	goto L240
L254:
	;
	goto L133
L255:
	;
	goto L87
L256:
	;
	F_pfree(m, v1498)
	mBase = m.M
	v1500 = m.ExcPending
	if v1500 != 0 {
		goto L1
	} else {
		goto L259
	}
L257:
	;
	goto L258
L258:
	;
	v1502 = F_make_result_safe(m, v165, int32(0))
	mBase = m.M
	v1503 = m.ExcPending
	if v1503 != 0 {
		goto L1
	} else {
		goto L260
	}
L259:
	;
	goto L258
L260:
	;
	v1504 = *(*int32)(unsafe.Add(mBase, uint32(v165)+16))
	if v1504 != 0 {
		goto L261
	} else {
		goto L262
	}
L261:
	;
	F_pfree(m, v1504)
	mBase = m.M
	v1506 = m.ExcPending
	if v1506 != 0 {
		goto L1
	} else {
		goto L264
	}
L262:
	;
	goto L263
L263:
	;
	m.G0 = v165 + int32(96)
	goto L31
L264:
	;
	goto L263
L265:
	;
	F_errcode(m, int32(50856066))
	mBase = m.M
	v1516 = m.ExcPending
	if v1516 != 0 {
		goto L1
	} else {
		goto L266
	}
L266:
	;
	F_errmsg(m, int32(_a_F_numeric_random_16), int32(0))
	mBase = m.M
	v1520 = m.ExcPending
	if v1520 != 0 {
		goto L1
	} else {
		goto L267
	}
L267:
	;
	F_errfinish(m, int32(_a_F_numeric_random_5), int32(_a_F_numeric_random_17), int32(_a_F_numeric_random_18))
	mBase = m.M
	v1525 = m.ExcPending
	if v1525 != 0 {
		goto L1
	} else {
		goto L268
	}
L268:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L269:
	;
	F_errfinish(m, int32(_a_F_numeric_random_5), int32(_a_F_numeric_random_19), int32(_a_F_numeric_random_7))
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L1
	} else {
		goto L270
	}
L270:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L271:
	;
	F_errfinish(m, int32(_a_F_numeric_random_5), int32(_a_F_numeric_random_20), int32(_a_F_numeric_random_7))
	mBase = m.M
	v1543 = m.ExcPending
	if v1543 != 0 {
		goto L1
	} else {
		goto L272
	}
L272:
	;
	base.Wasm_trap_unreachable()
	for {
	}
}
func F_numeric_sortsupport(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v8 int32
	_ = v8
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v14 int32
	_ = v14
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v24 int32
	_ = v24
	var v33 int32
	_ = v33
	var v39 int32
	_ = v39
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(1596)
	v8 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v5)+20)))
	if v8 == int32(1) {
		v11 = int32(_a_F_numeric_sortsupport_0)
		v12 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_sortsupport[0]))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		*(*int32)(unsafe.Add(mBase, _c_F_numeric_sortsupport[0])) = v14
		v17 = F_palloc(m, int32(48))
		mBase = m.M
		v20 = m.ExcPending
		if v20 != 0 {
			return int64(0)
		} else {
			v22 = F_palloc(m, int32(132))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int64(0)
			} else {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)) = uint8(v24)
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v22
				F_initHyperLogLog(m, v17+int32(24), int32(10))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(1597)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = int32(1598)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(1599)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v39
					*(*int32)(unsafe.Add(mBase, _c_F_numeric_sortsupport[0])) = v12
					return int64(0)
				}
			}
		}
	} else {
		return int64(0)
	}
}
func F_numeric_support(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v4 int64
	_ = v4
	var v5 int32
	_ = v5
	var v6 int32
	_ = v6
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
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
	var v34 int32
	_ = v34
	var v42 int32
	_ = v42
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v49 int64
	_ = v49
	v4 = int64(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(463) {
		v49 = v4
		return v49
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v49 = v4
			return v49
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+32)))
			if v16 != 0 {
				v49 = v4
				return v49
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v18 = F_exprTypmod(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int64(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					v23 = int32(4)
					v34 = int32(16)
					if (v18^v22)&int32(2047)|base.B2i32(v18 < v23)|base.B2i32(base.Ui32(int32(base.Ui32(v22-v23)>>(uint(v34)%32))) < base.Ui32(int32(base.Ui32(v18-v23)>>(uint(v34)%32)))) != 0 {
						v42 = base.B2i32(v23 <= v22)
					} else {
						v42 = int32(0)
					}
					if v42 != 0 {
						v49 = v4
						return v49
					} else {
						v43 = F_relabel_to_typmod(m, v17, v22)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int64(0)
						} else {
							v49 = base.I64_extend_i32_u(v43)
							return v49
						}
					}
				}
			}
		}
	}
}
func F_numeric_uminus(m *base.Module, l0 int32) int64 {
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
	var v18 int32
	_ = v18
	var v23 int32
	_ = v23
	var v28 int32
	_ = v28
	var v31 int32
	_ = v31
	var v40 int32
	_ = v40
	var v45 int32
	_ = v45
	var v52 int32
	_ = v52
	v4 = *(*int32)(unsafe.Add(mBase, uint32(l0)+24))
	v5 = F_pg_detoast_datum(m, v4)
	mBase = m.M
	v8 = m.ExcPending
	if v8 != 0 {
		return int64(0)
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
		v12 = F_palloc(m, int32(base.Ui32(v9)>>(uint(int32(2))%32)))
		mBase = m.M
		v13 = m.ExcPending
		if v13 != 0 {
			return int64(0)
		} else {
			v14 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
			v16 = int32(base.Ui32(v14) >> (uint(int32(2)) % 32))
			if v16 != 0 {
				base.MemoryCopy(m, v12, v5, v16)
			} else {
			}
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+4)))
			if base.Ui32(int32(_a_F_numeric_uminus_0)) <= base.Ui32(v18) {
				if v18 == int32(_a_F_numeric_uminus_0) {
					return base.I64_extend_i32_u(v12)
				} else {
					v52 = v18 ^ int32(_a_F_numeric_uminus_1)
					*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v52)
					return base.I64_extend_i32_u(v12)
				}
			} else {
				v23 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
				v28 = base.I32_extend16_s(v18)
				if int32(0) <= v28 {
					v31 = int32(-8)
				} else {
					v31 = int32(-6)
				}
				if base.Ui32(int32(base.Ui32(v23)>>(uint(int32(2))%32))+v31) < base.Ui32(int32(2)) {
					return base.I64_extend_i32_u(v12)
				} else {
					if v28 <= int32(-16385) {
						v52 = v18 ^ int32(_a_F_numeric_uminus_1)
						*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v52)
						return base.I64_extend_i32_u(v12)
					} else {
						if base.Ui32(v18) <= base.Ui32(int32(_a_F_numeric_uminus_2)) {
							v40 = v18 | int32(_a_F_numeric_uminus_3)
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v40)
							return base.I64_extend_i32_u(v12)
						} else {
							v45 = v18 & int32(_a_F_numeric_uminus_2)
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v45)
							return base.I64_extend_i32_u(v12)
						}
					}
				}
			}
		}
	}
}
