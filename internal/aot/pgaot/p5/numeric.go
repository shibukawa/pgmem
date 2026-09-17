package p5

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"math"
	"unsafe"
)

func F_numeric(m *base.Module, l0 int32) int32 {
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
	var v18 int32
	_ = v18
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v33 int32
	_ = v33
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v52 int32
	_ = v52
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v64 int32
	_ = v64
	var v66 int32
	_ = v66
	var v72 int32
	_ = v72
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v79 int32
	_ = v79
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v104 int32
	_ = v104
	var v109 int32
	_ = v109
	var v120 int32
	_ = v120
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v140 int32
	_ = v140
	var v143 int32
	_ = v143
	var v144 int32
	_ = v144
	var v153 int32
	_ = v153
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int32
	_ = v157
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v174 int32
	_ = v174
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v179 int32
	_ = v179
	var v181 int32
	_ = v181
	var v182 int32
	_ = v182
	var v183 int32
	_ = v183
	var v187 int32
	_ = v187
	var v188 int32
	_ = v188
	v7 = m.G0
	v9 = v7 - int32(32)
	m.G0 = v9
	v11 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v12 = F_pg_detoast_datum(m, v11)
	mBase = m.M
	v15 = m.ExcPending
	if v15 != 0 {
		return int32(0)
	} else {
		v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
		v18 = base.I32_extend16_s(v17)
		if base.Ui32(int32(_a_F_numeric_0)) <= base.Ui32(v17) {
			if base.B2i32(v18 == int32(-16384))|base.B2i32(v16 < int32(4)) != 0 {
				v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v62 = F_palloc(m, int32(base.Ui32(v59)>>(uint(int32(2))%32)))
				mBase = m.M
				v63 = m.ExcPending
				if v63 != 0 {
					return int32(0)
				} else {
					v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v66 = int32(base.Ui32(v64) >> (uint(int32(2)) % 32))
					if v66 == int32(0) {
						v188 = v62
					} else {
						base.MemoryCopy(m, v62, v12, v66)
						v188 = v62
					}
					m.G0 = v9 + int32(32)
					return v188
				}
			} else {
				v27 = F_errsave_start(m, int32(0))
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					if v27 == int32(0) {
						v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v62 = F_palloc(m, int32(base.Ui32(v59)>>(uint(int32(2))%32)))
						mBase = m.M
						v63 = m.ExcPending
						if v63 != 0 {
							return int32(0)
						} else {
							v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
							v66 = int32(base.Ui32(v64) >> (uint(int32(2)) % 32))
							if v66 == int32(0) {
								v188 = v62
							} else {
								base.MemoryCopy(m, v62, v12, v66)
								v188 = v62
							}
							m.G0 = v9 + int32(32)
							return v188
						}
					} else {
						F_errcode(m, int32(50331778))
						mBase = m.M
						v33 = m.ExcPending
						if v33 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_numeric_1), int32(0))
							mBase = m.M
							v37 = m.ExcPending
							if v37 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v9))) = int32(base.Ui32(v16-int32(4)) >> (uint(int32(16)) % 32))
								v43 = int32(21)
								*(*int32)(unsafe.Add(mBase, uint32(v9)+4)) = (v16<<(uint(v43)%32) - int32(_a_F_numeric_2)) >> (uint(v43) % 32)
								F_errdetail(m, int32(_a_F_numeric_3), v9)
								mBase = m.M
								v52 = m.ExcPending
								if v52 != 0 {
									return int32(0)
								} else {
									F_errsave_finish(m, int32(0), int32(_a_F_numeric_4), int32(_a_F_numeric_5), int32(_a_F_numeric_6))
									mBase = m.M
									v58 = m.ExcPending
									if v58 != 0 {
										return int32(0)
									} else {
										v59 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
										v62 = F_palloc(m, int32(base.Ui32(v59)>>(uint(int32(2))%32)))
										mBase = m.M
										v63 = m.ExcPending
										if v63 != 0 {
											return int32(0)
										} else {
											v64 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
											v66 = int32(base.Ui32(v64) >> (uint(int32(2)) % 32))
											if v66 == int32(0) {
												v188 = v62
											} else {
												base.MemoryCopy(m, v62, v12, v66)
												v188 = v62
											}
											m.G0 = v9 + int32(32)
											return v188
										}
									}
								}
							}
						}
					}
				}
			}
		} else {
			if v16 <= int32(3) {
				v72 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
				v75 = F_palloc(m, int32(base.Ui32(v72)>>(uint(int32(2))%32)))
				mBase = m.M
				v76 = m.ExcPending
				if v76 != 0 {
					return int32(0)
				} else {
					v77 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v79 = int32(base.Ui32(v77) >> (uint(int32(2)) % 32))
					if v79 == int32(0) {
						v188 = v75
					} else {
						base.MemoryCopy(m, v75, v12, v79)
						v188 = v75
					}
					m.G0 = v9 + int32(32)
					return v188
				}
			} else {
				if v18 < int32(0) {
					v95 = v17<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v17&int32(63)
				} else {
					v94 = int32(*(*int16)(unsafe.Add(mBase, uint32(v12)+6)))
					v95 = v94
				}
				v98 = int32(4)
				v104 = int32(21)
				v109 = (v16<<(uint(v104)%32) - int32(_a_F_numeric_2)) >> (uint(v104) % 32)
				if int32(0) <= v18 {
					v120 = v17 & int32(_a_F_numeric_7)
				} else {
					v120 = int32(base.Ui32(v17)>>(uint(int32(7))%32)) & int32(63)
				}
				if base.B2i32(int32(base.Ui32(v16-v98)>>(uint(int32(16))%32))-v109 < v95<<(uint(int32(2))%32)+v98)|base.B2i32(v109 < v120)|base.B2i32(v18 < int32(-16384))&base.B2i32(base.Ui32(int32(64)) <= base.Ui32(v109)) == int32(0) {
					v133 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
					v136 = F_palloc(m, int32(base.Ui32(v133)>>(uint(int32(2))%32)))
					mBase = m.M
					v137 = m.ExcPending
					if v137 != 0 {
						return int32(0)
					} else {
						v138 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
						v140 = int32(base.Ui32(v138) >> (uint(int32(2)) % 32))
						if v140 != 0 {
							base.MemoryCopy(m, v136, v12, v140)
						} else {
						}
						if int32(0) < v109 {
							v143 = v109
						} else {
							v143 = int32(0)
						}
						v144 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)))
						if v144&int32(_a_F_numeric_0) == int32(_a_F_numeric_8) {
							v153 = v144&int32(_a_F_numeric_9) | v143<<(uint(int32(7))%32)
							*(*uint16)(unsafe.Add(mBase, uint32(v136)+4)) = uint16(v153)
							v188 = v136
						} else {
							v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v136)+4)))
							v156 = int32(_a_F_numeric_0)
							v157 = v155 & v156
							if v157 != v156 {
								if v157 != int32(_a_F_numeric_8) {
									v168 = v157
								} else {
									v168 = v155 << (uint(int32(1)) % 32) & int32(_a_F_numeric_10)
								}
							} else {
								v168 = v155 & int32(_a_F_numeric_11)
							}
							v169 = v168 | v143
							*(*uint16)(unsafe.Add(mBase, uint32(v136)+4)) = uint16(v169)
							v188 = v136
						}
						m.G0 = v9 + int32(32)
						return v188
					}
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v9)+24)) = int64(0)
					v174 = v9 + int32(8)
					F_set_var_from_num(m, v12, v174)
					mBase = m.M
					v176 = m.ExcPending
					if v176 != 0 {
						return int32(0)
					} else {
						v178 = F_apply_typmod(m, v174, v16, int32(0))
						mBase = m.M
						v179 = m.ExcPending
						if v179 != 0 {
							return int32(0)
						} else {
							v181 = F_make_result_opt_error(m, v174, int32(0))
							mBase = m.M
							v182 = m.ExcPending
							if v182 != 0 {
								return int32(0)
							} else {
								v183 = *(*int32)(unsafe.Add(mBase, uint32(v9)+24))
								if v183 == int32(0) {
									v188 = v181
									m.G0 = v9 + int32(32)
									return v188
								} else {
									F_pfree(m, v183)
									mBase = m.M
									v187 = m.ExcPending
									if v187 != 0 {
										return int32(0)
									} else {
										v188 = v181
										m.G0 = v9 + int32(32)
										return v188
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
func F_numeric_cash(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v7 int32
	_ = v7
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int32
	_ = v13
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v21 int64
	_ = v21
	var v27 int32
	_ = v27
	var v28 int64
	_ = v28
	var v34 int64
	_ = v34
	var v36 int32
	_ = v36
	var v43 int64
	_ = v43
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v58 int64
	_ = v58
	var v60 int32
	_ = v60
	var v63 int64
	_ = v63
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v76 int32
	_ = v76
	var v77 int32
	_ = v77
	var v78 int64
	_ = v78
	var v79 int32
	_ = v79
	var v80 int32
	_ = v80
	v7 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v9 = F_PGLC_localeconv(m)
	mBase = m.M
	v12 = m.ExcPending
	if v12 != 0 {
		return int32(0)
	} else {
		v13 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v9)+41)))
		if base.Ui32(int32(10)) < base.Ui32(v13) {
			v16 = int32(2)
		} else {
			v16 = v13
		}
		v17 = base.I32_extend8_s(v16)
		if v17 <= int32(0) {
			v63 = int64(1)
		} else {
			v21 = int64(1)
			if base.Ui32(int32(8)) <= base.Ui32(v16) {
				v27 = int32(0)
				v28 = v21
				for {
					v34 = v28 * int64(100000000)
					v36 = v27 + int32(8)
					if v36 != v17&int32(120) {
						v27 = v36
						v28 = v34
						continue
					} else {
						break
					}
					break
				}
				if v16&int32(7) == int32(0) {
					v63 = v34
				} else {
					v43 = v34
					v51 = int32(0)
					v52 = v43
					for {
						v58 = v52 * int64(10)
						v60 = v51 + int32(1)
						if v60 != v17&int32(7) {
							v51 = v60
							v52 = v58
							continue
						} else {
							break
						}
						break
					}
					v63 = v58
				}
			} else {
				v43 = v21
				v51 = int32(0)
				v52 = v43
				for {
					v58 = v52 * int64(10)
					v60 = v51 + int32(1)
					if v60 != v17&int32(7) {
						v51 = v60
						v52 = v58
						continue
					} else {
						break
					}
					break
				}
				v63 = v58
			}
		}
		v69 = int32(0)
		v72 = F_int64_to_numeric(m, v63)
		mBase = m.M
		v73 = m.ExcPending
		if v73 != 0 {
			return int32(0)
		} else {
			v74 = F_DirectFunctionCall2Coll(m, int32(1262), v69, v7, v72)
			mBase = m.M
			v75 = m.ExcPending
			if v75 != 0 {
				return int32(0)
			} else {
				v76 = F_DirectFunctionCall1Coll(m, int32(1261), v69, v74)
				mBase = m.M
				v77 = m.ExcPending
				if v77 != 0 {
					return int32(0)
				} else {
					v78 = *(*int64)(unsafe.Add(mBase, uint32(v76)))
					v79 = F_Int64GetDatum(m, v78)
					mBase = m.M
					v80 = m.ExcPending
					if v80 != 0 {
						return int32(0)
					} else {
						return v79
					}
				}
			}
		}
	}
}
func F_numeric_float8_no_overflow(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	var v14 int32
	_ = v14
	var v15 int32
	_ = v15
	var v16 int32
	_ = v16
	var v24 float64
	_ = v24
	var v27 float64
	_ = v27
	var v28 int32
	_ = v28
	var v34 int32
	_ = v34
	var v35 int32
	_ = v35
	var v40 int32
	_ = v40
	var v50 int32
	_ = v50
	var v52 int32
	_ = v52
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v69 int32
	_ = v69
	var v72 int32
	_ = v72
	var v76 int32
	_ = v76
	var v81 float64
	_ = v81
	var v82 int32
	_ = v82
	var v86 float64
	_ = v86
	var v87 int32
	_ = v87
	var v88 int32
	_ = v88
	v6 = m.G0
	v8 = v6 - int32(32)
	m.G0 = v8
	v10 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v11 = F_pg_detoast_datum(m, v10)
	mBase = m.M
	v14 = m.ExcPending
	if v14 != 0 {
		return int32(0)
	} else {
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v11)+4)))
		v16 = base.I32_extend16_s(v15)
		if base.Ui32(int32(_a_F_numeric_float8_no_overflow_0)) <= base.Ui32(v15) {
			if v16 == int32(-4096) {
				v24 = math.Float64frombits(uint64(0xfff0000000000000))
			} else {
				v24 = math.Float64frombits(uint64(0x7ff8000000000000))
			}
			if v16 == int32(-12288) {
				v27 = math.Float64frombits(uint64(0x7ff0000000000000))
			} else {
				v27 = v24
			}
			v86 = v27
			v87 = F_Float8GetDatum(m, v86)
			mBase = m.M
			v88 = m.ExcPending
			if v88 != 0 {
				return int32(0)
			} else {
				m.G0 = v8 + int32(32)
				return v87
			}
		} else {
			v28 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
			v34 = base.B2i32(int32(0) <= v16)
			if int32(0) <= v16 {
				v35 = int32(-8)
			} else {
				v35 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+8)) = int32(base.Ui32(int32(base.Ui32(v28)>>(uint(int32(2))%32))+v35) >> (uint(int32(1)) % 32))
			if int32(0) <= v16 {
				v40 = int32(*(*int16)(unsafe.Add(mBase, uint32(v11)+6)))
				v50 = v40
			} else {
				v50 = v15<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v15&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+12)) = v50
			v52 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v8)+24)) = v52
			v61 = base.B2i32(v16 < v52)
			if v16 < v52 {
				v62 = int32(base.Ui32(v15)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v62 = v15 & int32(_a_F_numeric_float8_no_overflow_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+20)) = v62
			v69 = v15 & int32(_a_F_numeric_float8_no_overflow_0)
			if v69 == int32(_a_F_numeric_float8_no_overflow_2) {
				v72 = v15 << (uint(int32(1)) % 32) & int32(_a_F_numeric_float8_no_overflow_3)
			} else {
				v72 = v69
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+16)) = v72
			if v16 < v52 {
				v76 = int32(6)
			} else {
				v76 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v8)+28)) = v11 + v76
			v81 = F_numericvar_to_double_no_overflow(m, v8+int32(8))
			mBase = m.M
			v82 = m.ExcPending
			if v82 != 0 {
				return int32(0)
			} else {
				v86 = v81
				v87 = F_Float8GetDatum(m, v86)
				mBase = m.M
				v88 = m.ExcPending
				if v88 != 0 {
					return int32(0)
				} else {
					m.G0 = v8 + int32(32)
					return v87
				}
			}
		}
	}
}
func F_numeric_int4(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_numeric_int4_opt_error(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			return v8
		}
	}
}
func F_numeric_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v3 int32
	_ = v3
	var v6 int32
	_ = v6
	var v8 int64
	_ = v8
	var v9 int32
	_ = v9
	var v10 int32
	_ = v10
	var v11 int32
	_ = v11
	v2 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v3 = F_pg_detoast_datum(m, v2)
	mBase = m.M
	v6 = m.ExcPending
	if v6 != 0 {
		return int32(0)
	} else {
		v8 = F_numeric_int8_opt_error(m, v3, int32(0))
		mBase = m.M
		v9 = m.ExcPending
		if v9 != 0 {
			return int32(0)
		} else {
			v10 = F_Int64GetDatum(m, v8)
			mBase = m.M
			v11 = m.ExcPending
			if v11 != 0 {
				return int32(0)
			} else {
				return v10
			}
		}
	}
}
func F_numeric_mul_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v65 int32
	_ = v65
	var v66 int32
	_ = v66
	var v67 int32
	_ = v67
	var v72 int32
	_ = v72
	var v80 int32
	_ = v80
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v96 int32
	_ = v96
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v105 int32
	_ = v105
	var v113 int32
	_ = v113
	var v123 int32
	_ = v123
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v131 int32
	_ = v131
	var v138 int32
	_ = v138
	var v145 int32
	_ = v145
	var v155 int32
	_ = v155
	var v160 int32
	_ = v160
	var v161 int32
	_ = v161
	var v164 int32
	_ = v164
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v175 int32
	_ = v175
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v191 int32
	_ = v191
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v199 int32
	_ = v199
	var v208 int32
	_ = v208
	var v209 int32
	_ = v209
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v223 int32
	_ = v223
	var v226 int32
	_ = v226
	var v232 int32
	_ = v232
	var v233 int32
	_ = v233
	var v238 int32
	_ = v238
	var v248 int32
	_ = v248
	var v250 int64
	_ = v250
	var v261 int32
	_ = v261
	var v264 int32
	_ = v264
	var v266 int32
	_ = v266
	var v271 int32
	_ = v271
	var v272 int32
	_ = v272
	var v281 int32
	_ = v281
	var v291 int32
	_ = v291
	var v292 int32
	_ = v292
	var v297 int32
	_ = v297
	var v299 int32
	_ = v299
	var v311 int32
	_ = v311
	var v312 int32
	_ = v312
	var v314 int32
	_ = v314
	var v316 int32
	_ = v316
	var v317 int32
	_ = v317
	var v320 int32
	_ = v320
	var v321 int32
	_ = v321
	var v323 int32
	_ = v323
	var v324 int32
	_ = v324
	var v330 int32
	_ = v330
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v338 int32
	_ = v338
	var v347 int32
	_ = v347
	var v348 int32
	_ = v348
	var v351 int32
	_ = v351
	var v354 int32
	_ = v354
	var v356 int32
	_ = v356
	var v357 int32
	_ = v357
	var v358 int32
	_ = v358
	var v360 int32
	_ = v360
	var v371 int32
	_ = v371
	var v391 int32
	_ = v391
	var v392 int32
	_ = v392
	var v393 int32
	_ = v393
	var v397 int32
	_ = v397
	var v400 int32
	_ = v400
	var v401 int32
	_ = v401
	var v404 int32
	_ = v404
	var v405 int32
	_ = v405
	var v408 int32
	_ = v408
	var v409 int32
	_ = v409
	var v412 int32
	_ = v412
	var v413 int32
	_ = v413
	var v416 int32
	_ = v416
	var v417 int32
	_ = v417
	var v420 int32
	_ = v420
	var v421 int32
	_ = v421
	var v422 int32
	_ = v422
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v15 = base.I32_extend16_s(v14)
	if base.Ui32(v14) <= base.Ui32(int32(_a_F_numeric_mul_opt_error_0)) {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v19 = base.I32_extend16_s(v18)
		if base.Ui32(int32(_a_F_numeric_mul_opt_error_0)) < base.Ui32(v18) {
			v46 = v19
			if v46&int32(_a_F_numeric_mul_opt_error_1) != int32(_a_F_numeric_mul_opt_error_2) {
				if v14 != int32(_a_F_numeric_mul_opt_error_3) {
					if v14 != int32(_a_F_numeric_mul_opt_error_4) {
						v131 = v14 & int32(_a_F_numeric_mul_opt_error_2)
						if v46&int32(_a_F_numeric_mul_opt_error_1) == int32(_a_F_numeric_mul_opt_error_4) {
							if v131 == int32(_a_F_numeric_mul_opt_error_2) {
								v164 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return int32(0)
								} else {
									v422 = v164
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if int32(0) <= v15 {
									v145 = int32(-8)
								} else {
									v145 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v138)>>(uint(int32(2))%32))+v145) < base.Ui32(int32(2)) {
									v408 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return int32(0)
									} else {
										v422 = v408
										m.G0 = v12 + int32(80)
										return v422
									}
								} else {
									if v131 == int32(_a_F_numeric_mul_opt_error_7) {
										v155 = v14 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
									} else {
										v155 = v131
									}
									if v155 == int32(_a_F_numeric_mul_opt_error_8) {
										v164 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int32(0)
										} else {
											v422 = v164
											m.G0 = v12 + int32(80)
											return v422
										}
									} else {
										v160 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											v422 = v160
											m.G0 = v12 + int32(80)
											return v422
										}
									}
								}
							}
						} else {
							if v131 == int32(_a_F_numeric_mul_opt_error_2) {
								v194 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int32(0)
								} else {
									v422 = v194
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if int32(0) <= v15 {
									v175 = int32(-8)
								} else {
									v175 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v168)>>(uint(int32(2))%32))+v175) < base.Ui32(int32(2)) {
									v412 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
									mBase = m.M
									v413 = m.ExcPending
									if v413 != 0 {
										return int32(0)
									} else {
										v422 = v412
										m.G0 = v12 + int32(80)
										return v422
									}
								} else {
									if v131 == int32(_a_F_numeric_mul_opt_error_7) {
										v185 = v14 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
									} else {
										v185 = v131
									}
									if v185 == int32(_a_F_numeric_mul_opt_error_8) {
										v194 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
											return int32(0)
										} else {
											v422 = v194
											m.G0 = v12 + int32(80)
											return v422
										}
									} else {
										v190 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											v422 = v190
											m.G0 = v12 + int32(80)
											return v422
										}
									}
								}
							}
						}
					} else {
						v65 = v46 & int32(_a_F_numeric_mul_opt_error_1)
						v66 = int32(_a_F_numeric_mul_opt_error_2)
						v67 = v46 & v66
						if v67 == v66 {
							if v65 == int32(_a_F_numeric_mul_opt_error_4) {
								v95 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v422 = v95
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								v420 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
								mBase = m.M
								v421 = m.ExcPending
								if v421 != 0 {
									return int32(0)
								} else {
									v422 = v420
									m.G0 = v12 + int32(80)
									return v422
								}
							}
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if int32(0) <= base.I32_extend16_s(v46) {
								v80 = int32(-8)
							} else {
								v80 = int32(-6)
							}
							if base.Ui32(int32(base.Ui32(v72)>>(uint(int32(2))%32))+v80) < base.Ui32(int32(2)) {
								v400 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
								mBase = m.M
								v401 = m.ExcPending
								if v401 != 0 {
									return int32(0)
								} else {
									v422 = v400
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								if v67 == int32(_a_F_numeric_mul_opt_error_7) {
									v90 = v65 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
								} else {
									v90 = v67
								}
								if v90 == int32(_a_F_numeric_mul_opt_error_8) {
									v420 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
									mBase = m.M
									v421 = m.ExcPending
									if v421 != 0 {
										return int32(0)
									} else {
										v422 = v420
										m.G0 = v12 + int32(80)
										return v422
									}
								} else {
									v95 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										v422 = v95
										m.G0 = v12 + int32(80)
										return v422
									}
								}
							}
						}
					}
				} else {
					v98 = v46 & int32(_a_F_numeric_mul_opt_error_1)
					v99 = int32(_a_F_numeric_mul_opt_error_2)
					v100 = v46 & v99
					if v100 == v99 {
						if v98 == int32(_a_F_numeric_mul_opt_error_4) {
							v128 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								v422 = v128
								m.G0 = v12 + int32(80)
								return v422
							}
						} else {
							v416 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
							mBase = m.M
							v417 = m.ExcPending
							if v417 != 0 {
								return int32(0)
							} else {
								v422 = v416
								m.G0 = v12 + int32(80)
								return v422
							}
						}
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v46) {
							v113 = int32(-8)
						} else {
							v113 = int32(-6)
						}
						if base.Ui32(int32(base.Ui32(v105)>>(uint(int32(2))%32))+v113) < base.Ui32(int32(2)) {
							v404 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
							mBase = m.M
							v405 = m.ExcPending
							if v405 != 0 {
								return int32(0)
							} else {
								v422 = v404
								m.G0 = v12 + int32(80)
								return v422
							}
						} else {
							if v100 == int32(_a_F_numeric_mul_opt_error_7) {
								v123 = v98 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
							} else {
								v123 = v100
							}
							if v123 == int32(_a_F_numeric_mul_opt_error_8) {
								v416 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
								mBase = m.M
								v417 = m.ExcPending
								if v417 != 0 {
									return int32(0)
								} else {
									v422 = v416
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								v128 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									v422 = v128
									m.G0 = v12 + int32(80)
									return v422
								}
							}
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v422 = v56
					m.G0 = v12 + int32(80)
					return v422
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v28 = base.B2i32(int32(0) <= v15)
			if int32(0) <= v15 {
				v29 = int32(-8)
			} else {
				v29 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(base.Ui32(int32(base.Ui32(v22)>>(uint(int32(2))%32))+v29) >> (uint(int32(1)) % 32))
			if int32(0) <= v15 {
				v196 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v197 = v196
			} else {
				v197 = v14<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v14&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v197
			v199 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v199
			v208 = base.B2i32(v15 < v199)
			if v15 < v199 {
				v209 = int32(base.Ui32(v14)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v209 = v14 & int32(_a_F_numeric_mul_opt_error_10)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v209
			v216 = v14 & int32(_a_F_numeric_mul_opt_error_2)
			if v216 == int32(_a_F_numeric_mul_opt_error_7) {
				v219 = v14 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
			} else {
				v219 = v216
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v219
			if v15 < v199 {
				v223 = int32(6)
			} else {
				v223 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = l0 + v223
			v226 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v232 = base.B2i32(int32(0) <= v19)
			if int32(0) <= v19 {
				v233 = int32(-8)
			} else {
				v233 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(base.Ui32(int32(base.Ui32(v226)>>(uint(int32(2))%32))+v233) >> (uint(int32(1)) % 32))
			if int32(0) <= v19 {
				v238 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				v248 = v238
			} else {
				v248 = v18<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v18&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v248
			v250 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v250
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v250
			*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v250
			v261 = v18 & int32(_a_F_numeric_mul_opt_error_2)
			if v261 == int32(_a_F_numeric_mul_opt_error_7) {
				v264 = v18 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
			} else {
				v264 = v261
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v264
			v266 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v266
			v271 = base.B2i32(v19 < v266)
			if v19 < v266 {
				v272 = int32(6)
			} else {
				v272 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l1 + v272
			if v19 < v266 {
				v281 = int32(base.Ui32(v18)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v281 = v18 & int32(_a_F_numeric_mul_opt_error_10)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v281
			F_mul_var(m, v12+int32(56), v12+int32(32), v12+int32(8), v281+v209)
			mBase = m.M
			v291 = m.ExcPending
			if v291 != 0 {
				return int32(0)
			} else {
				v292 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
				if v292 < int32(_a_F_numeric_mul_opt_error_8) {
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = int32(_a_F_numeric_mul_opt_error_10)
					v297 = *(*int32)(unsafe.Add(mBase, uint32(v12)+12))
					v299 = v297 << (uint(int32(2)) % 32)
					if v299+int32(_a_F_numeric_mul_opt_error_11) < int32(0) {
						*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = int32(0)
						*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = int64(0)
					} else {
						v311 = base.I32_div_s(v299+int32(_a_F_numeric_mul_opt_error_12), int32(4))
						v312 = *(*int32)(unsafe.Add(mBase, uint32(v12)+8))
						if v312 < v311 {
						} else {
							v314 = *(*int32)(unsafe.Add(mBase, uint32(v12)+28))
							*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v311
							v316 = int32(1)
							v317 = v311 - v316
							v320 = v314 + v317<<(uint(v316)%32)
							v321 = int32(*(*int16)(unsafe.Add(mBase, uint32(v320))))
							v323 = base.I32_rem_s(v321, int32(10))
							v324 = v321 - v323
							*(*uint16)(unsafe.Add(mBase, uint32(v320))) = uint16(v324)
							if v323 < int32(5) {
								v360 = v317
							} else {
								v330 = base.I32_extend16_s(v324)
								if int32(_a_F_numeric_mul_opt_error_13) < v330 {
									v333 = int32(-9990)
								} else {
									v333 = int32(10)
								}
								v334 = v333 + v330
								*(*uint16)(unsafe.Add(mBase, uint32(v320))) = uint16(v334)
								if v330 < int32(_a_F_numeric_mul_opt_error_14) {
									v360 = v317
								} else {
									v338 = v317
									for {
										v347 = int32(1)
										v348 = v338 - v347
										v351 = v314 + v348<<(uint(v347)%32)
										v354 = int32(*(*int16)(unsafe.Add(mBase, uint32(v351))))
										v356 = base.B2i32(int32(_a_F_numeric_mul_opt_error_15) < v354)
										if int32(_a_F_numeric_mul_opt_error_15) < v354 {
											v357 = int32(-9999)
										} else {
											v357 = v347
										}
										v358 = v357 + v354
										*(*uint16)(unsafe.Add(mBase, uint32(v351))) = uint16(v358)
										if int32(_a_F_numeric_mul_opt_error_15) < v354 {
											v338 = v348
											continue
										} else {
											break
										}
										break
									}
									v360 = v348
								}
							}
							if int32(0) <= v360 {
							} else {
								v371 = int32(1)
								*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v297 + v371
								*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = v311 + v371
								*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v314 - int32(2)
							}
						}
					}
				}
				v391 = F_make_result_opt_error(m, v12+int32(8), l2)
				mBase = m.M
				v392 = m.ExcPending
				if v392 != 0 {
					return int32(0)
				} else {
					v393 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					if v393 == int32(0) {
						v422 = v391
						m.G0 = v12 + int32(80)
						return v422
					} else {
						F_pfree(m, v393)
						mBase = m.M
						v397 = m.ExcPending
						if v397 != 0 {
							return int32(0)
						} else {
							v422 = v391
							m.G0 = v12 + int32(80)
							return v422
						}
					}
				}
			}
		}
	} else {
		if v15 == int32(-16384) {
			v56 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v422 = v56
				m.G0 = v12 + int32(80)
				return v422
			}
		} else {
			v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v46 = v45
			if v46&int32(_a_F_numeric_mul_opt_error_1) != int32(_a_F_numeric_mul_opt_error_2) {
				if v14 != int32(_a_F_numeric_mul_opt_error_3) {
					if v14 != int32(_a_F_numeric_mul_opt_error_4) {
						v131 = v14 & int32(_a_F_numeric_mul_opt_error_2)
						if v46&int32(_a_F_numeric_mul_opt_error_1) == int32(_a_F_numeric_mul_opt_error_4) {
							if v131 == int32(_a_F_numeric_mul_opt_error_2) {
								v164 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
								mBase = m.M
								v165 = m.ExcPending
								if v165 != 0 {
									return int32(0)
								} else {
									v422 = v164
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								v138 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if int32(0) <= v15 {
									v145 = int32(-8)
								} else {
									v145 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v138)>>(uint(int32(2))%32))+v145) < base.Ui32(int32(2)) {
									v408 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return int32(0)
									} else {
										v422 = v408
										m.G0 = v12 + int32(80)
										return v422
									}
								} else {
									if v131 == int32(_a_F_numeric_mul_opt_error_7) {
										v155 = v14 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
									} else {
										v155 = v131
									}
									if v155 == int32(_a_F_numeric_mul_opt_error_8) {
										v164 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
										mBase = m.M
										v165 = m.ExcPending
										if v165 != 0 {
											return int32(0)
										} else {
											v422 = v164
											m.G0 = v12 + int32(80)
											return v422
										}
									} else {
										v160 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
										mBase = m.M
										v161 = m.ExcPending
										if v161 != 0 {
											return int32(0)
										} else {
											v422 = v160
											m.G0 = v12 + int32(80)
											return v422
										}
									}
								}
							}
						} else {
							if v131 == int32(_a_F_numeric_mul_opt_error_2) {
								v194 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
								mBase = m.M
								v195 = m.ExcPending
								if v195 != 0 {
									return int32(0)
								} else {
									v422 = v194
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								v168 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								if int32(0) <= v15 {
									v175 = int32(-8)
								} else {
									v175 = int32(-6)
								}
								if base.Ui32(int32(base.Ui32(v168)>>(uint(int32(2))%32))+v175) < base.Ui32(int32(2)) {
									v412 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
									mBase = m.M
									v413 = m.ExcPending
									if v413 != 0 {
										return int32(0)
									} else {
										v422 = v412
										m.G0 = v12 + int32(80)
										return v422
									}
								} else {
									if v131 == int32(_a_F_numeric_mul_opt_error_7) {
										v185 = v14 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
									} else {
										v185 = v131
									}
									if v185 == int32(_a_F_numeric_mul_opt_error_8) {
										v194 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
										mBase = m.M
										v195 = m.ExcPending
										if v195 != 0 {
											return int32(0)
										} else {
											v422 = v194
											m.G0 = v12 + int32(80)
											return v422
										}
									} else {
										v190 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
										mBase = m.M
										v191 = m.ExcPending
										if v191 != 0 {
											return int32(0)
										} else {
											v422 = v190
											m.G0 = v12 + int32(80)
											return v422
										}
									}
								}
							}
						}
					} else {
						v65 = v46 & int32(_a_F_numeric_mul_opt_error_1)
						v66 = int32(_a_F_numeric_mul_opt_error_2)
						v67 = v46 & v66
						if v67 == v66 {
							if v65 == int32(_a_F_numeric_mul_opt_error_4) {
								v95 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
								mBase = m.M
								v96 = m.ExcPending
								if v96 != 0 {
									return int32(0)
								} else {
									v422 = v95
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								v420 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
								mBase = m.M
								v421 = m.ExcPending
								if v421 != 0 {
									return int32(0)
								} else {
									v422 = v420
									m.G0 = v12 + int32(80)
									return v422
								}
							}
						} else {
							v72 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
							if int32(0) <= base.I32_extend16_s(v46) {
								v80 = int32(-8)
							} else {
								v80 = int32(-6)
							}
							if base.Ui32(int32(base.Ui32(v72)>>(uint(int32(2))%32))+v80) < base.Ui32(int32(2)) {
								v400 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
								mBase = m.M
								v401 = m.ExcPending
								if v401 != 0 {
									return int32(0)
								} else {
									v422 = v400
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								if v67 == int32(_a_F_numeric_mul_opt_error_7) {
									v90 = v65 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
								} else {
									v90 = v67
								}
								if v90 == int32(_a_F_numeric_mul_opt_error_8) {
									v420 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
									mBase = m.M
									v421 = m.ExcPending
									if v421 != 0 {
										return int32(0)
									} else {
										v422 = v420
										m.G0 = v12 + int32(80)
										return v422
									}
								} else {
									v95 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
									mBase = m.M
									v96 = m.ExcPending
									if v96 != 0 {
										return int32(0)
									} else {
										v422 = v95
										m.G0 = v12 + int32(80)
										return v422
									}
								}
							}
						}
					}
				} else {
					v98 = v46 & int32(_a_F_numeric_mul_opt_error_1)
					v99 = int32(_a_F_numeric_mul_opt_error_2)
					v100 = v46 & v99
					if v100 == v99 {
						if v98 == int32(_a_F_numeric_mul_opt_error_4) {
							v128 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
							mBase = m.M
							v129 = m.ExcPending
							if v129 != 0 {
								return int32(0)
							} else {
								v422 = v128
								m.G0 = v12 + int32(80)
								return v422
							}
						} else {
							v416 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
							mBase = m.M
							v417 = m.ExcPending
							if v417 != 0 {
								return int32(0)
							} else {
								v422 = v416
								m.G0 = v12 + int32(80)
								return v422
							}
						}
					} else {
						v105 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
						if int32(0) <= base.I32_extend16_s(v46) {
							v113 = int32(-8)
						} else {
							v113 = int32(-6)
						}
						if base.Ui32(int32(base.Ui32(v105)>>(uint(int32(2))%32))+v113) < base.Ui32(int32(2)) {
							v404 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
							mBase = m.M
							v405 = m.ExcPending
							if v405 != 0 {
								return int32(0)
							} else {
								v422 = v404
								m.G0 = v12 + int32(80)
								return v422
							}
						} else {
							if v100 == int32(_a_F_numeric_mul_opt_error_7) {
								v123 = v98 << (uint(int32(1)) % 32) & int32(_a_F_numeric_mul_opt_error_8)
							} else {
								v123 = v100
							}
							if v123 == int32(_a_F_numeric_mul_opt_error_8) {
								v416 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_9), int32(0))
								mBase = m.M
								v417 = m.ExcPending
								if v417 != 0 {
									return int32(0)
								} else {
									v422 = v416
									m.G0 = v12 + int32(80)
									return v422
								}
							} else {
								v128 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_5), int32(0))
								mBase = m.M
								v129 = m.ExcPending
								if v129 != 0 {
									return int32(0)
								} else {
									v422 = v128
									m.G0 = v12 + int32(80)
									return v422
								}
							}
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(_a_F_numeric_mul_opt_error_6), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v422 = v56
					m.G0 = v12 + int32(80)
					return v422
				}
			}
		}
	}
}
func F_numeric_pg_lsn(m *base.Module, l0 int32) int32 {
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
	var v337 int32
	_ = v337
	var v338 int32
	_ = v338
	v10 = int64(0)
	v11 = m.G0
	v13 = v11 - int32(48)
	m.G0 = v13
	v15 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
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
	v337 = F_Int64GetDatum(m, v334)
	mBase = m.M
	v338 = m.ExcPending
	if v338 != 0 {
		goto L1
	} else {
		goto L78
	}
L78:
	;
	m.G0 = v13 + int32(48)
	return v337
}
func F_numeric_poly_serialize(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v20 int32
	_ = v20
	var v23 int32
	_ = v23
	var v48 int32
	_ = v48
	var v49 int32
	_ = v49
	var v53 int32
	_ = v53
	var v57 int32
	_ = v57
	var v58 int64
	_ = v58
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v63 int32
	_ = v63
	var v65 int64
	_ = v65
	var v67 int64
	_ = v67
	var v69 int64
	_ = v69
	var v72 int64
	_ = v72
	var v74 int64
	_ = v74
	var v76 int64
	_ = v76
	var v78 int64
	_ = v78
	var v104 int64
	_ = v104
	var v105 int64
	_ = v105
	var v107 int32
	_ = v107
	var v108 int32
	_ = v108
	var v110 int32
	_ = v110
	var v119 int64
	_ = v119
	var v127 int64
	_ = v127
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v145 int32
	_ = v145
	var v146 int64
	_ = v146
	var v147 int64
	_ = v147
	var v151 int32
	_ = v151
	var v154 int32
	_ = v154
	var v160 int64
	_ = v160
	var v161 int64
	_ = v161
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v174 int64
	_ = v174
	var v175 int64
	_ = v175
	var v182 int32
	_ = v182
	var v183 int64
	_ = v183
	var v184 int64
	_ = v184
	var v185 int64
	_ = v185
	var v186 int64
	_ = v186
	var v191 int64
	_ = v191
	var v194 int64
	_ = v194
	var v197 int64
	_ = v197
	var v200 int64
	_ = v200
	var v201 int64
	_ = v201
	var v205 int64
	_ = v205
	var v212 int64
	_ = v212
	var v224 int32
	_ = v224
	var v225 int64
	_ = v225
	var v226 int64
	_ = v226
	var v230 int64
	_ = v230
	var v235 int32
	_ = v235
	var v236 int32
	_ = v236
	var v238 int32
	_ = v238
	var v243 int32
	_ = v243
	var v258 int32
	_ = v258
	var v259 int64
	_ = v259
	var v260 int64
	_ = v260
	var v262 int32
	_ = v262
	var v264 int32
	_ = v264
	var v265 int32
	_ = v265
	var v267 int32
	_ = v267
	var v277 int64
	_ = v277
	var v285 int64
	_ = v285
	var v292 int64
	_ = v292
	var v293 int64
	_ = v293
	var v296 int32
	_ = v296
	var v299 int32
	_ = v299
	var v305 int64
	_ = v305
	var v306 int64
	_ = v306
	var v309 int32
	_ = v309
	var v310 int32
	_ = v310
	var v313 int32
	_ = v313
	var v315 int32
	_ = v315
	var v319 int64
	_ = v319
	var v320 int64
	_ = v320
	var v326 int64
	_ = v326
	var v327 int64
	_ = v327
	var v328 int64
	_ = v328
	var v329 int64
	_ = v329
	var v334 int64
	_ = v334
	var v337 int64
	_ = v337
	var v340 int64
	_ = v340
	var v343 int64
	_ = v343
	var v344 int64
	_ = v344
	var v348 int64
	_ = v348
	var v355 int64
	_ = v355
	var v367 int32
	_ = v367
	var v368 int64
	_ = v368
	var v369 int64
	_ = v369
	var v373 int64
	_ = v373
	var v378 int32
	_ = v378
	var v379 int32
	_ = v379
	var v381 int32
	_ = v381
	var v386 int32
	_ = v386
	var v397 int32
	_ = v397
	var v401 int32
	_ = v401
	var v403 int32
	_ = v403
	var v404 int32
	_ = v404
	var v409 int32
	_ = v409
	v2 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(112)
	m.G0 = v16
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+4))
	if v20 == v2 {
		v48 = int32(0)
	} else {
		v23 = *(*int32)(unsafe.Add(mBase, uint32(v20)))
		switch v23 - int32(429) {
		case 0:
			v48 = int32(1)
		case 1:
			v48 = int32(2)
		default:
			v48 = int32(0)
		}
	}
	if v48 != 0 {
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = int64(0)
		v53 = v16 + int32(96)
		F_pq_begintypsend(m, v53)
		mBase = m.M
		v57 = m.ExcPending
		if v57 != 0 {
			return int32(0)
		} else {
			v58 = *(*int64)(unsafe.Add(mBase, uint32(v49)+8))
			F_enlargeStringInfo(m, v53, int32(8))
			mBase = m.M
			v61 = m.ExcPending
			if v61 != 0 {
				return int32(0)
			} else {
				v62 = *(*int32)(unsafe.Add(mBase, uint32(v16)+100))
				v63 = *(*int32)(unsafe.Add(mBase, uint32(v16)+96))
				v65 = int64(56)
				v67 = int64(65280)
				v69 = int64(40)
				v72 = int64(16711680)
				v74 = int64(24)
				v76 = int64(4278190080)
				v78 = int64(8)
				*(*int64)(unsafe.Add(mBase, uint32(v62+v63))) = v58<<(uint(v65)%64) | v58&v67<<(uint(v69)%64) | (v58&v72<<(uint(v74)%64) | v58&v76<<(uint(v78)%64)) | (int64(base.Ui64(v58)>>(uint(v78)%64))&v76 | int64(base.Ui64(v58)>>(uint(v74)%64))&v72 | (int64(base.Ui64(v58)>>(uint(v69)%64))&v67 | int64(base.Ui64(v58)>>(uint(v65)%64))))
				*(*int32)(unsafe.Add(mBase, uint32(v16)+100)) = v62 + int32(8)
				v104 = *(*int64)(unsafe.Add(mBase, uint32(v49)+16))
				v105 = *(*int64)(unsafe.Add(mBase, uint32(v49)+24))
				v107 = F_palloc(m, int32(22))
				mBase = m.M
				v108 = m.ExcPending
				if v108 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v107
					v110 = int32(0)
					*(*uint16)(unsafe.Add(mBase, uint32(v107))) = uint16(v110)
					*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v107 + int32(2)
					if v105 < int64(0) {
						*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = int64(16384)
						v119 = int64(0)
						v146 = v119 - v104
						v147 = v119 - (v105 + base.I64_extend_i32_u(base.B2i32(v104 != v119)))
						v151 = int32(0)
						v154 = v107 + int32(22)
						v160 = v146
						v161 = v147
						for {
							v165 = v16 + int32(48)
							v168 = m.G0
							v169 = int32(16)
							v170 = v168 - v169
							m.G0 = v170
							F___udivmodti4(m, v170, v160, v161, int64(10000), int64(0))
							mBase = m.M
							v174 = *(*int64)(unsafe.Add(mBase, uint32(v170)+8))
							v175 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
							*(*int64)(unsafe.Add(mBase, uint32(v165))) = v175
							*(*int64)(unsafe.Add(mBase, uint32(v165)+8)) = v174
							m.G0 = v170 + v169
							v182 = v16 + int32(32)
							v183 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
							v184 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
							v185 = int64(55536)
							v186 = int64(0)
							v191 = int64(32)
							v194 = int64(base.Ui64(v183) >> (uint(v191) % 64))
							v197 = int64(4294967295)
							v200 = v183 & v197
							v201 = v185 * v200
							v205 = int64(base.Ui64(v201)>>(uint(v191)%64)) + v185*v194
							v212 = v200*v186 + v205&v197
							*(*int64)(unsafe.Add(mBase, uint32(v182)+8)) = v183*v186 + v184*v185 + v186*v194 + int64(base.Ui64(v205)>>(uint(v191)%64)) + int64(base.Ui64(v212)>>(uint(v191)%64))
							*(*int64)(unsafe.Add(mBase, uint32(v182))) = v201&v197 | v212<<(uint(v191)%64)
							v224 = v154 - int32(2)
							v225 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
							v226 = v225 + v160
							*(*uint16)(unsafe.Add(mBase, uint32(v224))) = uint16(v226)
							v230 = int64(0)
							v235 = v151 + int32(1)
							if v161 == v230 {
								v236 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v160))
							} else {
								v236 = base.B2i32(v161 != v230)
							}
							if v236 != 0 {
								v151 = v235
								v154 = v224
								v160 = v183
								v161 = v184
								continue
							} else {
								break
							}
							break
						}
						*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v224
						v238 = v235
						v243 = v151
					} else {
						v127 = int64(0)
						*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v127
						if v104|v105 != v127 {
							v146 = v104
							v147 = v105
							v151 = int32(0)
							v154 = v107 + int32(22)
							v160 = v146
							v161 = v147
							for {
								v165 = v16 + int32(48)
								v168 = m.G0
								v169 = int32(16)
								v170 = v168 - v169
								m.G0 = v170
								F___udivmodti4(m, v170, v160, v161, int64(10000), int64(0))
								mBase = m.M
								v174 = *(*int64)(unsafe.Add(mBase, uint32(v170)+8))
								v175 = *(*int64)(unsafe.Add(mBase, uint32(v170)))
								*(*int64)(unsafe.Add(mBase, uint32(v165))) = v175
								*(*int64)(unsafe.Add(mBase, uint32(v165)+8)) = v174
								m.G0 = v170 + v169
								v182 = v16 + int32(32)
								v183 = *(*int64)(unsafe.Add(mBase, uint32(v16)+48))
								v184 = *(*int64)(unsafe.Add(mBase, uint32(v16)+56))
								v185 = int64(55536)
								v186 = int64(0)
								v191 = int64(32)
								v194 = int64(base.Ui64(v183) >> (uint(v191) % 64))
								v197 = int64(4294967295)
								v200 = v183 & v197
								v201 = v185 * v200
								v205 = int64(base.Ui64(v201)>>(uint(v191)%64)) + v185*v194
								v212 = v200*v186 + v205&v197
								*(*int64)(unsafe.Add(mBase, uint32(v182)+8)) = v183*v186 + v184*v185 + v186*v194 + int64(base.Ui64(v205)>>(uint(v191)%64)) + int64(base.Ui64(v212)>>(uint(v191)%64))
								*(*int64)(unsafe.Add(mBase, uint32(v182))) = v201&v197 | v212<<(uint(v191)%64)
								v224 = v154 - int32(2)
								v225 = *(*int64)(unsafe.Add(mBase, uint32(v16)+32))
								v226 = v225 + v160
								*(*uint16)(unsafe.Add(mBase, uint32(v224))) = uint16(v226)
								v230 = int64(0)
								v235 = v151 + int32(1)
								if v161 == v230 {
									v236 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v160))
								} else {
									v236 = base.B2i32(v161 != v230)
								}
								if v236 != 0 {
									v151 = v235
									v154 = v224
									v160 = v183
									v161 = v184
									continue
								} else {
									break
								}
								break
							}
							*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v224
							v238 = v235
							v243 = v151
						} else {
							v238 = int32(0)
							v243 = v2
						}
					}
					*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v243
					*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v238
					F_numericvar_serialize(m, v16+int32(96), v16+int32(72))
					mBase = m.M
					v258 = m.ExcPending
					if v258 != 0 {
						return int32(0)
					} else {
						v259 = *(*int64)(unsafe.Add(mBase, uint32(v49)+32))
						v260 = *(*int64)(unsafe.Add(mBase, uint32(v49)+40))
						F_pfree(m, v107)
						mBase = m.M
						v262 = m.ExcPending
						if v262 != 0 {
							return int32(0)
						} else {
							v264 = F_palloc(m, int32(22))
							mBase = m.M
							v265 = m.ExcPending
							if v265 != 0 {
								return int32(0)
							} else {
								*(*int32)(unsafe.Add(mBase, uint32(v16)+88)) = v264
								v267 = int32(0)
								*(*uint16)(unsafe.Add(mBase, uint32(v264))) = uint16(v267)
								*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v264 + int32(2)
								if v260 < int64(0) {
									*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = int64(16384)
									v277 = int64(0)
									v292 = v277 - v259
									v293 = v277 - (v260 + base.I64_extend_i32_u(base.B2i32(v259 != v277)))
									v296 = v267
									v299 = v264 + int32(22)
									v305 = v292
									v306 = v293
									for {
										v309 = int32(16)
										v310 = v16 + v309
										v313 = m.G0
										v315 = v313 - v309
										m.G0 = v315
										F___udivmodti4(m, v315, v305, v306, int64(10000), int64(0))
										mBase = m.M
										v319 = *(*int64)(unsafe.Add(mBase, uint32(v315)+8))
										v320 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
										*(*int64)(unsafe.Add(mBase, uint32(v310))) = v320
										*(*int64)(unsafe.Add(mBase, uint32(v310)+8)) = v319
										m.G0 = v315 + v309
										v326 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
										v327 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
										v328 = int64(55536)
										v329 = int64(0)
										v334 = int64(32)
										v337 = int64(base.Ui64(v326) >> (uint(v334) % 64))
										v340 = int64(4294967295)
										v343 = v326 & v340
										v344 = v328 * v343
										v348 = int64(base.Ui64(v344)>>(uint(v334)%64)) + v328*v337
										v355 = v343*v329 + v348&v340
										*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v326*v329 + v327*v328 + v329*v337 + int64(base.Ui64(v348)>>(uint(v334)%64)) + int64(base.Ui64(v355)>>(uint(v334)%64))
										*(*int64)(unsafe.Add(mBase, uint32(v16))) = v344&v340 | v355<<(uint(v334)%64)
										v367 = v299 - int32(2)
										v368 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
										v369 = v368 + v305
										*(*uint16)(unsafe.Add(mBase, uint32(v367))) = uint16(v369)
										v373 = int64(0)
										v378 = v296 + int32(1)
										if v306 == v373 {
											v379 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v305))
										} else {
											v379 = base.B2i32(v306 != v373)
										}
										if v379 != 0 {
											v296 = v378
											v299 = v367
											v305 = v326
											v306 = v327
											continue
										} else {
											break
										}
										break
									}
									*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v367
									v381 = v378
									v386 = v296
								} else {
									v285 = int64(0)
									*(*int64)(unsafe.Add(mBase, uint32(v16)+80)) = v285
									if v259|v260 == v285 {
										v381 = v267
										v386 = int32(0)
									} else {
										v292 = v259
										v293 = v260
										v296 = v267
										v299 = v264 + int32(22)
										v305 = v292
										v306 = v293
										for {
											v309 = int32(16)
											v310 = v16 + v309
											v313 = m.G0
											v315 = v313 - v309
											m.G0 = v315
											F___udivmodti4(m, v315, v305, v306, int64(10000), int64(0))
											mBase = m.M
											v319 = *(*int64)(unsafe.Add(mBase, uint32(v315)+8))
											v320 = *(*int64)(unsafe.Add(mBase, uint32(v315)))
											*(*int64)(unsafe.Add(mBase, uint32(v310))) = v320
											*(*int64)(unsafe.Add(mBase, uint32(v310)+8)) = v319
											m.G0 = v315 + v309
											v326 = *(*int64)(unsafe.Add(mBase, uint32(v16)+16))
											v327 = *(*int64)(unsafe.Add(mBase, uint32(v16)+24))
											v328 = int64(55536)
											v329 = int64(0)
											v334 = int64(32)
											v337 = int64(base.Ui64(v326) >> (uint(v334) % 64))
											v340 = int64(4294967295)
											v343 = v326 & v340
											v344 = v328 * v343
											v348 = int64(base.Ui64(v344)>>(uint(v334)%64)) + v328*v337
											v355 = v343*v329 + v348&v340
											*(*int64)(unsafe.Add(mBase, uint32(v16)+8)) = v326*v329 + v327*v328 + v329*v337 + int64(base.Ui64(v348)>>(uint(v334)%64)) + int64(base.Ui64(v355)>>(uint(v334)%64))
											*(*int64)(unsafe.Add(mBase, uint32(v16))) = v344&v340 | v355<<(uint(v334)%64)
											v367 = v299 - int32(2)
											v368 = *(*int64)(unsafe.Add(mBase, uint32(v16)))
											v369 = v368 + v305
											*(*uint16)(unsafe.Add(mBase, uint32(v367))) = uint16(v369)
											v373 = int64(0)
											v378 = v296 + int32(1)
											if v306 == v373 {
												v379 = base.B2i32(base.Ui64(int64(9999)) < base.Ui64(v305))
											} else {
												v379 = base.B2i32(v306 != v373)
											}
											if v379 != 0 {
												v296 = v378
												v299 = v367
												v305 = v326
												v306 = v327
												continue
											} else {
												break
											}
											break
										}
										*(*int32)(unsafe.Add(mBase, uint32(v16)+92)) = v367
										v381 = v378
										v386 = v296
									}
								}
								*(*int32)(unsafe.Add(mBase, uint32(v16)+76)) = v386
								*(*int32)(unsafe.Add(mBase, uint32(v16)+72)) = v381
								v397 = v16 + int32(96)
								F_numericvar_serialize(m, v397, v16+int32(72))
								mBase = m.M
								v401 = m.ExcPending
								if v401 != 0 {
									return int32(0)
								} else {
									v403 = *(*int32)(unsafe.Add(mBase, uint32(v397)))
									v404 = *(*int32)(unsafe.Add(mBase, uint32(v397)+4))
									*(*int32)(unsafe.Add(mBase, uint32(v403))) = v404 << (uint(int32(2)) % 32)
									F_pfree(m, v264)
									mBase = m.M
									v409 = m.ExcPending
									if v409 != 0 {
										return int32(0)
									} else {
										m.G0 = v16 + int32(112)
										return v403
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
		v136 = m.ExcPending
		if v136 != 0 {
			return int32(0)
		} else {
			F_errmsg_internal(m, int32(_a_F_numeric_poly_serialize_0), int32(0))
			mBase = m.M
			v140 = m.ExcPending
			if v140 != 0 {
				return int32(0)
			} else {
				F_errfinish(m, int32(_a_F_numeric_poly_serialize_1), int32(_a_F_numeric_poly_serialize_2), int32(_a_F_numeric_poly_serialize_3))
				mBase = m.M
				v145 = m.ExcPending
				if v145 != 0 {
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
func F_numeric_power(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
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
	var v32 int32
	_ = v32
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v47 int32
	_ = v47
	var v57 int32
	_ = v57
	var v64 int32
	_ = v64
	var v72 int32
	_ = v72
	var v78 int32
	_ = v78
	var v79 int32
	_ = v79
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v85 int32
	_ = v85
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	var v101 int32
	_ = v101
	var v102 int32
	_ = v102
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v118 int32
	_ = v118
	var v122 int32
	_ = v122
	var v128 int32
	_ = v128
	var v129 int32
	_ = v129
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v137 int32
	_ = v137
	var v141 int32
	_ = v141
	var v154 int32
	_ = v154
	var v156 int32
	_ = v156
	var v161 int32
	_ = v161
	var v162 int32
	_ = v162
	var v163 int32
	_ = v163
	var v165 int32
	_ = v165
	var v173 int32
	_ = v173
	var v175 int32
	_ = v175
	var v184 int32
	_ = v184
	var v185 int32
	_ = v185
	var v190 int32
	_ = v190
	var v199 int32
	_ = v199
	var v201 int32
	_ = v201
	var v208 int32
	_ = v208
	var v215 int32
	_ = v215
	var v216 int32
	_ = v216
	var v219 int32
	_ = v219
	var v226 int32
	_ = v226
	var v231 int32
	_ = v231
	var v239 int32
	_ = v239
	var v243 int32
	_ = v243
	var v248 int32
	_ = v248
	var v252 int32
	_ = v252
	var v258 int32
	_ = v258
	var v269 int32
	_ = v269
	var v279 int32
	_ = v279
	var v282 int32
	_ = v282
	var v283 int32
	_ = v283
	var v288 int32
	_ = v288
	var v293 int32
	_ = v293
	var v294 int32
	_ = v294
	var v295 int32
	_ = v295
	var v298 int32
	_ = v298
	var v299 int32
	_ = v299
	var v306 int32
	_ = v306
	var v311 int32
	_ = v311
	var v318 int32
	_ = v318
	var v321 int32
	_ = v321
	var v322 int32
	_ = v322
	var v324 int32
	_ = v324
	var v325 int32
	_ = v325
	var v326 int32
	_ = v326
	var v333 int32
	_ = v333
	var v334 int32
	_ = v334
	var v342 int32
	_ = v342
	var v346 int32
	_ = v346
	var v350 int32
	_ = v350
	var v357 int32
	_ = v357
	var v360 int32
	_ = v360
	var v361 int32
	_ = v361
	var v363 int32
	_ = v363
	var v365 int32
	_ = v365
	var v368 int32
	_ = v368
	var v369 int32
	_ = v369
	var v370 int32
	_ = v370
	var v378 int32
	_ = v378
	var v385 int32
	_ = v385
	var v386 int32
	_ = v386
	var v389 int32
	_ = v389
	var v390 int32
	_ = v390
	var v400 int32
	_ = v400
	var v411 int32
	_ = v411
	var v417 int32
	_ = v417
	var v418 int32
	_ = v418
	var v420 int32
	_ = v420
	var v430 int32
	_ = v430
	var v432 int32
	_ = v432
	var v443 int32
	_ = v443
	var v451 int32
	_ = v451
	var v452 int32
	_ = v452
	var v454 int32
	_ = v454
	var v455 int32
	_ = v455
	var v468 int32
	_ = v468
	var v472 int32
	_ = v472
	var v478 int32
	_ = v478
	var v479 int32
	_ = v479
	var v480 int32
	_ = v480
	var v482 int32
	_ = v482
	var v487 int32
	_ = v487
	var v491 int32
	_ = v491
	var v504 int32
	_ = v504
	var v506 int32
	_ = v506
	var v511 int32
	_ = v511
	var v512 int32
	_ = v512
	var v513 int32
	_ = v513
	var v515 int32
	_ = v515
	var v523 int32
	_ = v523
	var v525 int32
	_ = v525
	var v534 int32
	_ = v534
	var v535 int32
	_ = v535
	var v540 int32
	_ = v540
	var v549 int32
	_ = v549
	var v551 int32
	_ = v551
	var v558 int32
	_ = v558
	var v565 int32
	_ = v565
	var v566 int32
	_ = v566
	var v569 int32
	_ = v569
	var v576 int32
	_ = v576
	var v581 int32
	_ = v581
	var v589 int32
	_ = v589
	var v593 int32
	_ = v593
	var v598 int32
	_ = v598
	var v602 int32
	_ = v602
	var v608 int32
	_ = v608
	var v619 int32
	_ = v619
	var v629 int32
	_ = v629
	var v632 int32
	_ = v632
	var v633 int32
	_ = v633
	var v641 int32
	_ = v641
	var v642 int32
	_ = v642
	var v648 int32
	_ = v648
	var v649 int32
	_ = v649
	var v657 int32
	_ = v657
	var v660 int32
	_ = v660
	var v665 int32
	_ = v665
	var v668 int32
	_ = v668
	var v673 int32
	_ = v673
	var v677 int32
	_ = v677
	var v687 int32
	_ = v687
	var v689 int32
	_ = v689
	var v691 int32
	_ = v691
	var v692 int32
	_ = v692
	var v693 int32
	_ = v693
	var v704 int32
	_ = v704
	var v706 int32
	_ = v706
	var v715 int32
	_ = v715
	var v717 int32
	_ = v717
	var v718 int32
	_ = v718
	var v724 int32
	_ = v724
	var v727 int32
	_ = v727
	var v734 int32
	_ = v734
	var v735 int32
	_ = v735
	var v736 int32
	_ = v736
	var v746 int32
	_ = v746
	var v747 int32
	_ = v747
	var v752 int32
	_ = v752
	var v753 int32
	_ = v753
	var v754 int32
	_ = v754
	var v755 int32
	_ = v755
	var v756 int32
	_ = v756
	var v762 int32
	_ = v762
	var v766 int32
	_ = v766
	var v772 int32
	_ = v772
	var v777 int32
	_ = v777
	var v778 int32
	_ = v778
	var v779 int32
	_ = v779
	var v782 int32
	_ = v782
	var v783 int32
	_ = v783
	var v790 int32
	_ = v790
	var v791 int32
	_ = v791
	var v792 int32
	_ = v792
	var v802 int32
	_ = v802
	var v803 int32
	_ = v803
	var v808 int32
	_ = v808
	var v809 int32
	_ = v809
	var v810 int32
	_ = v810
	var v811 int32
	_ = v811
	var v812 int32
	_ = v812
	var v818 int32
	_ = v818
	var v822 int32
	_ = v822
	var v828 int32
	_ = v828
	var v833 int32
	_ = v833
	var v839 int32
	_ = v839
	var v840 int32
	_ = v840
	var v843 int32
	_ = v843
	var v844 int32
	_ = v844
	var v851 int32
	_ = v851
	var v852 int32
	_ = v852
	var v855 int32
	_ = v855
	var v856 int32
	_ = v856
	var v859 int32
	_ = v859
	var v860 int32
	_ = v860
	var v862 int32
	_ = v862
	var v865 int32
	_ = v865
	var v870 int32
	_ = v870
	var v873 int32
	_ = v873
	var v878 int32
	_ = v878
	var v882 int32
	_ = v882
	var v892 int32
	_ = v892
	var v894 int32
	_ = v894
	var v896 int32
	_ = v896
	var v897 int32
	_ = v897
	var v898 int32
	_ = v898
	var v909 int32
	_ = v909
	var v911 int32
	_ = v911
	var v920 int32
	_ = v920
	var v922 int32
	_ = v922
	var v923 int32
	_ = v923
	var v929 int32
	_ = v929
	var v932 int32
	_ = v932
	var v935 int32
	_ = v935
	var v939 int32
	_ = v939
	var v940 int32
	_ = v940
	var v945 int32
	_ = v945
	var v952 int32
	_ = v952
	var v953 int32
	_ = v953
	var v956 int32
	_ = v956
	var v957 int32
	_ = v957
	var v958 int32
	_ = v958
	var v966 int32
	_ = v966
	var v970 int32
	_ = v970
	var v977 int32
	_ = v977
	var v986 int32
	_ = v986
	var v989 int32
	_ = v989
	var v993 int64
	_ = v993
	var v999 int32
	_ = v999
	var v1004 int32
	_ = v1004
	var v1005 int32
	_ = v1005
	var v1007 int32
	_ = v1007
	var v1008 int32
	_ = v1008
	var v1009 int32
	_ = v1009
	var v1011 int32
	_ = v1011
	var v1013 int32
	_ = v1013
	var v1023 int32
	_ = v1023
	var v1025 int32
	_ = v1025
	var v1026 int32
	_ = v1026
	var v1037 int32
	_ = v1037
	var v1039 int32
	_ = v1039
	var v1048 int32
	_ = v1048
	var v1049 int32
	_ = v1049
	var v1053 int32
	_ = v1053
	var v1054 int32
	_ = v1054
	var v1056 int32
	_ = v1056
	var v1061 int32
	_ = v1061
	var v1062 int32
	_ = v1062
	var v1064 int32
	_ = v1064
	var v1065 int32
	_ = v1065
	var v1068 int32
	_ = v1068
	var v1070 int32
	_ = v1070
	var v1080 int32
	_ = v1080
	var v1082 int32
	_ = v1082
	var v1083 int32
	_ = v1083
	var v1094 int32
	_ = v1094
	var v1096 int32
	_ = v1096
	var v1105 int32
	_ = v1105
	var v1106 int32
	_ = v1106
	var v1110 int32
	_ = v1110
	var v1111 int32
	_ = v1111
	var v1115 int32
	_ = v1115
	var v1117 int32
	_ = v1117
	var v1122 int32
	_ = v1122
	var v1123 int32
	_ = v1123
	var v1126 int64
	_ = v1126
	var v1131 int32
	_ = v1131
	var v1132 int32
	_ = v1132
	var v1133 float64
	_ = v1133
	var v1138 int32
	_ = v1138
	var v1140 int32
	_ = v1140
	var v1143 int32
	_ = v1143
	var v1146 int32
	_ = v1146
	var v1148 float64
	_ = v1148
	var v1153 int32
	_ = v1153
	var v1155 float64
	_ = v1155
	var v1160 int32
	_ = v1160
	var v1164 float64
	_ = v1164
	var v1172 float64
	_ = v1172
	var v1173 int32
	_ = v1173
	var v1184 int64
	_ = v1184
	var v1199 int32
	_ = v1199
	var v1201 int64
	_ = v1201
	var v1211 int64
	_ = v1211
	var v1215 int64
	_ = v1215
	var v1217 int32
	_ = v1217
	var v1218 int32
	_ = v1218
	var v1220 int32
	_ = v1220
	var v1224 float64
	_ = v1224
	var v1226 float64
	_ = v1226
	var v1239 float64
	_ = v1239
	var v1242 float64
	_ = v1242
	var v1247 float64
	_ = v1247
	var v1248 float64
	_ = v1248
	var v1249 float64
	_ = v1249
	var v1250 float64
	_ = v1250
	var v1255 float64
	_ = v1255
	var v1256 float64
	_ = v1256
	var v1257 float64
	_ = v1257
	var v1282 float64
	_ = v1282
	var v1294 float64
	_ = v1294
	var v1316 float64
	_ = v1316
	var v1320 float64
	_ = v1320
	var v1325 float64
	_ = v1325
	var v1330 int32
	_ = v1330
	var v1332 int32
	_ = v1332
	var v1335 int64
	_ = v1335
	var v1341 int32
	_ = v1341
	var v1342 int32
	_ = v1342
	var v1344 int32
	_ = v1344
	var v1346 int32
	_ = v1346
	var v1349 int32
	_ = v1349
	var v1353 int32
	_ = v1353
	var v1354 int32
	_ = v1354
	var v1357 int32
	_ = v1357
	var v1359 int32
	_ = v1359
	var v1366 int32
	_ = v1366
	var v1369 int64
	_ = v1369
	var v1372 int32
	_ = v1372
	var v1375 int32
	_ = v1375
	var v1376 int32
	_ = v1376
	var v1377 int32
	_ = v1377
	var v1389 int32
	_ = v1389
	var v1391 int32
	_ = v1391
	var v1392 int64
	_ = v1392
	var v1394 int64
	_ = v1394
	var v1397 int32
	_ = v1397
	var v1406 int32
	_ = v1406
	var v1409 int32
	_ = v1409
	var v1418 int32
	_ = v1418
	var v1420 int32
	_ = v1420
	var v1424 int32
	_ = v1424
	var v1425 int32
	_ = v1425
	var v1436 int32
	_ = v1436
	var v1439 int32
	_ = v1439
	var v1440 int32
	_ = v1440
	var v1443 int32
	_ = v1443
	var v1444 int32
	_ = v1444
	var v1445 int32
	_ = v1445
	var v1447 int32
	_ = v1447
	var v1448 int32
	_ = v1448
	var v1449 int32
	_ = v1449
	var v1452 int32
	_ = v1452
	var v1455 int32
	_ = v1455
	var v1460 int32
	_ = v1460
	var v1464 int32
	_ = v1464
	var v1470 int32
	_ = v1470
	var v1476 int32
	_ = v1476
	var v1477 int32
	_ = v1477
	var v1480 int32
	_ = v1480
	var v1483 int32
	_ = v1483
	var v1485 int32
	_ = v1485
	var v1486 int32
	_ = v1486
	var v1487 int32
	_ = v1487
	var v1490 int32
	_ = v1490
	var v1498 int32
	_ = v1498
	var v1502 int32
	_ = v1502
	var v1503 int32
	_ = v1503
	var v1506 int32
	_ = v1506
	var v1527 int32
	_ = v1527
	var v1530 int32
	_ = v1530
	var v1532 int32
	_ = v1532
	var v1534 int32
	_ = v1534
	var v1539 int32
	_ = v1539
	var v1541 int32
	_ = v1541
	var v1545 int64
	_ = v1545
	var v1550 int32
	_ = v1550
	var v1552 int32
	_ = v1552
	var v1553 int32
	_ = v1553
	var v1554 int32
	_ = v1554
	var v1555 int32
	_ = v1555
	var v1558 int32
	_ = v1558
	var v1560 int32
	_ = v1560
	var v1566 int64
	_ = v1566
	var v1568 int64
	_ = v1568
	var v1572 float64
	_ = v1572
	var v1577 int32
	_ = v1577
	var v1579 int32
	_ = v1579
	var v1582 int32
	_ = v1582
	var v1583 int32
	_ = v1583
	var v1584 int32
	_ = v1584
	var v1587 int32
	_ = v1587
	var v1591 int32
	_ = v1591
	var v1593 int32
	_ = v1593
	var v1594 int64
	_ = v1594
	var v1596 int64
	_ = v1596
	var v1599 int32
	_ = v1599
	var v1600 int32
	_ = v1600
	var v1603 int32
	_ = v1603
	var v1605 int32
	_ = v1605
	var v1609 int64
	_ = v1609
	var v1612 int64
	_ = v1612
	var v1614 int32
	_ = v1614
	var v1615 int32
	_ = v1615
	var v1618 int32
	_ = v1618
	var v1623 int32
	_ = v1623
	var v1637 int32
	_ = v1637
	var v1641 int32
	_ = v1641
	var v1642 int32
	_ = v1642
	var v1645 int32
	_ = v1645
	var v1646 int32
	_ = v1646
	var v1648 int32
	_ = v1648
	var v1650 int32
	_ = v1650
	var v1651 int32
	_ = v1651
	var v1654 int32
	_ = v1654
	var v1656 int32
	_ = v1656
	var v1657 int32
	_ = v1657
	var v1660 int32
	_ = v1660
	var v1664 int32
	_ = v1664
	var v1665 int32
	_ = v1665
	var v1666 int32
	_ = v1666
	var v1667 int32
	_ = v1667
	var v1669 int32
	_ = v1669
	var v1670 int32
	_ = v1670
	var v1673 int32
	_ = v1673
	var v1675 int32
	_ = v1675
	var v1680 int32
	_ = v1680
	var v1686 int32
	_ = v1686
	var v1688 int32
	_ = v1688
	var v1689 int32
	_ = v1689
	var v1691 int64
	_ = v1691
	var v1695 int32
	_ = v1695
	var v1699 int32
	_ = v1699
	var v1700 int32
	_ = v1700
	var v1702 int32
	_ = v1702
	var v1709 int32
	_ = v1709
	var v1714 int32
	_ = v1714
	var v1715 int32
	_ = v1715
	var v1716 int32
	_ = v1716
	var v1718 int32
	_ = v1718
	var v1720 int32
	_ = v1720
	var v1724 int32
	_ = v1724
	var v1727 int64
	_ = v1727
	var v1733 int64
	_ = v1733
	var v1753 int32
	_ = v1753
	var v1761 int32
	_ = v1761
	var v1767 int32
	_ = v1767
	var v1771 int32
	_ = v1771
	var v1773 int32
	_ = v1773
	var v1776 int32
	_ = v1776
	var v1777 int32
	_ = v1777
	var v1778 int32
	_ = v1778
	var v1781 int32
	_ = v1781
	var v1783 int64
	_ = v1783
	var v1787 int64
	_ = v1787
	var v1793 int32
	_ = v1793
	var v1794 int32
	_ = v1794
	var v1797 int32
	_ = v1797
	var v1799 int32
	_ = v1799
	var v1801 int32
	_ = v1801
	var v1802 int32
	_ = v1802
	var v1803 int32
	_ = v1803
	var v1804 int32
	_ = v1804
	var v1807 int32
	_ = v1807
	var v1809 int32
	_ = v1809
	var v1813 int32
	_ = v1813
	var v1815 int32
	_ = v1815
	var v1816 float64
	_ = v1816
	var v1817 int32
	_ = v1817
	var v1823 int32
	_ = v1823
	var v1825 int32
	_ = v1825
	var v1828 int64
	_ = v1828
	var v1833 int32
	_ = v1833
	var v1838 int32
	_ = v1838
	var v1839 int32
	_ = v1839
	var v1840 int32
	_ = v1840
	var v1842 int32
	_ = v1842
	var v1844 int32
	_ = v1844
	var v1847 int32
	_ = v1847
	var v1848 int32
	_ = v1848
	var v1849 int32
	_ = v1849
	var v1852 int32
	_ = v1852
	var v1855 int32
	_ = v1855
	var v1856 int32
	_ = v1856
	var v1859 int32
	_ = v1859
	var v1861 int32
	_ = v1861
	var v1865 int32
	_ = v1865
	var v1867 int32
	_ = v1867
	var v1869 int32
	_ = v1869
	var v1870 int32
	_ = v1870
	var v1876 int32
	_ = v1876
	var v1878 int32
	_ = v1878
	var v1879 int32
	_ = v1879
	var v1881 int32
	_ = v1881
	var v1885 int32
	_ = v1885
	var v1889 int32
	_ = v1889
	var v1892 int32
	_ = v1892
	var v1896 int32
	_ = v1896
	var v1901 int32
	_ = v1901
	var v1905 int32
	_ = v1905
	var v1908 int32
	_ = v1908
	var v1912 int32
	_ = v1912
	var v1917 int32
	_ = v1917
	var v1921 int32
	_ = v1921
	var v1924 int32
	_ = v1924
	var v1928 int32
	_ = v1928
	var v1933 int32
	_ = v1933
	var v1937 int32
	_ = v1937
	var v1940 int32
	_ = v1940
	var v1944 int32
	_ = v1944
	var v1949 int32
	_ = v1949
	var v1953 int32
	_ = v1953
	var v1956 int32
	_ = v1956
	var v1960 int32
	_ = v1960
	var v1965 int32
	_ = v1965
	var v1969 int32
	_ = v1969
	var v1972 int32
	_ = v1972
	var v1976 int32
	_ = v1976
	var v1981 int32
	_ = v1981
	var v1985 int32
	_ = v1985
	var v1988 int32
	_ = v1988
	var v1992 int32
	_ = v1992
	var v1997 int32
	_ = v1997
	var v2001 int32
	_ = v2001
	var v2004 int32
	_ = v2004
	var v2008 int32
	_ = v2008
	var v2013 int32
	_ = v2013
	var v2025 int32
	_ = v2025
	var v2028 int32
	_ = v2028
	var v2037 int32
	_ = v2037
	var v2039 int32
	_ = v2039
	var v2043 int32
	_ = v2043
	var v2044 int32
	_ = v2044
	var v2055 int32
	_ = v2055
	var v2058 int32
	_ = v2058
	var v2059 int32
	_ = v2059
	var v2062 int32
	_ = v2062
	var v2063 int32
	_ = v2063
	var v2064 int32
	_ = v2064
	var v2066 int32
	_ = v2066
	var v2067 int32
	_ = v2067
	var v2068 int32
	_ = v2068
	var v2071 int32
	_ = v2071
	var v2074 int32
	_ = v2074
	var v2079 int32
	_ = v2079
	var v2083 int32
	_ = v2083
	var v2089 int32
	_ = v2089
	var v2095 int32
	_ = v2095
	var v2096 int32
	_ = v2096
	var v2099 int32
	_ = v2099
	var v2102 int32
	_ = v2102
	var v2104 int32
	_ = v2104
	var v2105 int32
	_ = v2105
	var v2106 int32
	_ = v2106
	var v2109 int32
	_ = v2109
	var v2117 int32
	_ = v2117
	var v2121 int32
	_ = v2121
	var v2122 int32
	_ = v2122
	var v2125 int32
	_ = v2125
	var v2159 int32
	_ = v2159
	var v2160 int32
	_ = v2160
	var v2161 int32
	_ = v2161
	var v2165 int32
	_ = v2165
	var v2171 int32
	_ = v2171
	var v2172 int32
	_ = v2172
	var v2175 int32
	_ = v2175
	v16 = m.G0
	v18 = v16 - int32(144)
	m.G0 = v18
	v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v25 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
		goto L21
	} else {
		goto L22
	}
L4:
	;
	m.G0 = v18 + int32(144)
	return v2175
L5:
	;
	v2171 = F_make_result_opt_error(m, int32(_a_F_numeric_power_1), int32(0))
	mBase = m.M
	v2172 = m.ExcPending
	if v2172 != 0 {
		goto L1
	} else {
		goto L672
	}
L6:
	;
	v2159 = F_make_result_opt_error(m, v18, int32(0))
	mBase = m.M
	v2160 = m.ExcPending
	if v2160 != 0 {
		goto L1
	} else {
		goto L669
	}
L7:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v1349
	v2025 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v2028 = v1349 + v2025<<(uint(int32(2))%32)
	if v2028+int32(4) < int32(0) {
		goto L644
	} else {
		goto L645
	}
L8:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v2001 = m.ExcPending
	if v2001 != 0 {
		goto L1
	} else {
		goto L639
	}
L9:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1985 = m.ExcPending
	if v1985 != 0 {
		goto L1
	} else {
		goto L635
	}
L10:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1969 = m.ExcPending
	if v1969 != 0 {
		goto L1
	} else {
		goto L631
	}
L11:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1953 = m.ExcPending
	if v1953 != 0 {
		goto L1
	} else {
		goto L627
	}
L12:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1937 = m.ExcPending
	if v1937 != 0 {
		goto L1
	} else {
		goto L623
	}
L13:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1921 = m.ExcPending
	if v1921 != 0 {
		goto L1
	} else {
		goto L619
	}
L14:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1905 = m.ExcPending
	if v1905 != 0 {
		goto L1
	} else {
		goto L615
	}
L15:
	;
	F_errstart_cold(m, int32(21), int32(0))
	mBase = m.M
	v1889 = m.ExcPending
	if v1889 != 0 {
		goto L1
	} else {
		goto L611
	}
L16:
	;
	v958 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if int32(0) <= base.I32_extend16_s(v32) {
		goto L331
	} else {
		goto L332
	}
L17:
	;
	if v29 == int32(-12288) {
		goto L94
	} else {
		goto L95
	}
L18:
	;
	if v57&int32(_a_F_numeric_power_2) == int32(_a_F_numeric_power_3) {
		goto L5
	} else {
		goto L92
	}
L19:
	;
	if v47 == int32(0) {
		goto L5
	} else {
		goto L39
	}
L20:
	;
	v84 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
	v85 = v84
	goto L19
L21:
	;
	v32 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	if base.Ui32(v32) <= base.Ui32(int32(_a_F_numeric_power_0)) {
		goto L16
	} else {
		goto L24
	}
L22:
	;
	goto L23
L23:
	;
	v57 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	if v29 != int32(-16384) {
		goto L18
	} else {
		goto L30
	}
L24:
	;
	if v32 != int32(_a_F_numeric_power_3) {
		v288 = v32
		goto L17
	} else {
		goto L25
	}
L25:
	;
	v37 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v43 = base.B2i32(int32(0) <= v29)
	if int32(0) <= v29 {
		goto L26
	} else {
		goto L27
	}
L26:
	;
	v44 = int32(-8)
	goto L28
L27:
	;
	v44 = int32(-6)
	goto L28
L28:
	;
	v47 = int32(base.Ui32(int32(base.Ui32(v37)>>(uint(int32(2))%32))+v44) >> (uint(int32(1)) % 32))
	if int32(0) <= v29 {
		goto L20
	} else {
		goto L29
	}
L29:
	;
	v85 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
	goto L19
L30:
	;
	if base.Ui32(int32(_a_F_numeric_power_0)) < base.Ui32(v57&int32(_a_F_numeric_power_2)) {
		goto L31
	} else {
		goto L32
	}
L31:
	;
	v82 = F_make_result_opt_error(m, int32(_a_F_numeric_power_1), int32(0))
	mBase = m.M
	v83 = m.ExcPending
	if v83 != 0 {
		goto L1
	} else {
		goto L38
	}
L32:
	;
	v64 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if int32(0) <= base.I32_extend16_s(v57) {
		goto L33
	} else {
		goto L34
	}
L33:
	;
	v72 = int32(-8)
	goto L35
L34:
	;
	v72 = int32(-6)
	goto L35
L35:
	;
	if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v64)>>(uint(int32(2))%32))+v72) {
		goto L31
	} else {
		goto L36
	}
L36:
	;
	v78 = F_make_result_opt_error(m, int32(_a_F_numeric_power_4), int32(0))
	mBase = m.M
	v79 = m.ExcPending
	if v79 != 0 {
		goto L1
	} else {
		goto L37
	}
L37:
	;
	v2175 = v78
	goto L4
L38:
	;
	v2175 = v82
	goto L4
L39:
	;
	v93 = v28 & int32(_a_F_numeric_power_3)
	if v93 == int32(_a_F_numeric_power_5) {
		goto L40
	} else {
		goto L41
	}
L40:
	;
	v96 = v28 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_6)
	goto L42
L41:
	;
	v96 = v93
	goto L42
L42:
	;
	if v96 != 0 {
		goto L5
	} else {
		goto L43
	}
L43:
	;
	if v29 < int32(0) {
		goto L44
	} else {
		goto L45
	}
L44:
	;
	v101 = int32(6)
	goto L46
L45:
	;
	v101 = int32(8)
	goto L46
L46:
	;
	v102 = v21 + v101
	v104 = int32(1)
	v105 = int32(0)
	if base.B2i32(v105 < v85)&base.B2i32(v105 < v47) == v105 {
		v137 = v85
		v141 = v105
		goto L49
	} else {
		goto L50
	}
L47:
	;
	if v279 != 0 {
		goto L5
	} else {
		goto L90
	}
L48:
	;
	v279 = v269
	goto L47
L49:
	;
	if int32(0)|base.B2i32(v105 <= v137) != 0 {
		v173 = v105
		v175 = v105
		goto L56
	} else {
		goto L57
	}
L50:
	;
	v118 = v85
	v122 = v105
	goto L51
L51:
	;
	v128 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+v122<<(uint(int32(1))%32)))))
	if v128 != 0 {
		v269 = int32(1)
		goto L48
	} else {
		goto L53
	}
L52:
	;
	v137 = v132
	v141 = v130
	goto L49
L53:
	;
	v129 = int32(1)
	v130 = v122 + v129
	v132 = v118 - v129
	if v132 <= v105 {
		v137 = v132
		v141 = v130
		goto L49
	} else {
		goto L54
	}
L54:
	;
	if v130 < v47 {
		v118 = v132
		v122 = v130
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	if v137 != v173 {
		v215 = v141
		v216 = v175
		goto L63
	} else {
		goto L64
	}
L57:
	;
	v154 = v105
	v156 = v105
	goto L58
L58:
	;
	v161 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v156<<(uint(int32(1))%32))+uint32(_c_F_numeric_power[0]))))
	if v161 != 0 {
		v269 = int32(-1)
		goto L48
	} else {
		goto L60
	}
L59:
	;
	v173 = v165
	v175 = v163
	goto L56
L60:
	;
	v162 = int32(1)
	v163 = v156 + v162
	v165 = v154 - v162
	if v165 <= v137 {
		v173 = v165
		v175 = v163
		goto L56
	} else {
		goto L61
	}
L61:
	;
	if v163 < v104 {
		v154 = v165
		v156 = v163
		goto L58
	} else {
		goto L62
	}
L62:
	;
	goto L59
L63:
	;
	if v47 < v215 {
		goto L72
	} else {
		goto L73
	}
L64:
	;
	v184 = v141
	v185 = v175
	goto L65
L65:
	;
	if base.B2i32(v47 <= v184)|base.B2i32(v104 <= v185) != 0 {
		v215 = v184
		v216 = v185
		goto L63
	} else {
		goto L67
	}
L66:
	;
	if base.I32_extend16_s(v201) < base.I32_extend16_s(v199) {
		goto L69
	} else {
		goto L70
	}
L67:
	;
	v190 = int32(1)
	v199 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+v184<<(uint(v190)%32)))))
	v201 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v185<<(uint(v190)%32))+uint32(_c_F_numeric_power[0]))))
	if v199 == v201 {
		v184 = v184 + v190
		v185 = v185 + v190
		goto L65
	} else {
		goto L68
	}
L68:
	;
	goto L66
L69:
	;
	v208 = int32(1)
	goto L71
L70:
	;
	v208 = int32(-1)
	goto L71
L71:
	;
	v279 = v208
	goto L47
L72:
	;
	v219 = v215
	goto L74
L73:
	;
	v219 = v47
	goto L74
L74:
	;
	v226 = v215
	goto L75
L75:
	;
	if v219 == v226 {
		goto L77
	} else {
		goto L78
	}
L76:
	;
	v269 = v252
	goto L48
L77:
	;
	if v104 < v216 {
		goto L80
	} else {
		goto L81
	}
L78:
	;
	goto L79
L79:
	;
	v252 = int32(1)
	v258 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v102+v226<<(uint(v252)%32)))))
	if v258 == int32(0) {
		v226 = v226 + v252
		goto L75
	} else {
		goto L89
	}
L80:
	;
	v231 = v216
	goto L82
L81:
	;
	v231 = v104
	goto L82
L82:
	;
	v239 = v216
	goto L83
L83:
	;
	if v231 == v239 {
		goto L85
	} else {
		goto L86
	}
L84:
	;
	v269 = int32(-1)
	goto L48
L85:
	;
	v279 = int32(0)
	goto L47
L86:
	;
	goto L87
L87:
	;
	v243 = int32(1)
	v248 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v239<<(uint(v243)%32))+uint32(_c_F_numeric_power[0]))))
	if v248 == int32(0) {
		v239 = v239 + v243
		goto L83
	} else {
		goto L88
	}
L88:
	;
	goto L84
L89:
	;
	goto L76
L90:
	;
	v282 = F_make_result_opt_error(m, int32(_a_F_numeric_power_4), int32(0))
	mBase = m.M
	v283 = m.ExcPending
	if v283 != 0 {
		goto L1
	} else {
		goto L91
	}
L91:
	;
	v2175 = v282
	goto L4
L92:
	;
	v288 = v57
	goto L17
L93:
	;
	v324 = v288 & int32(_a_F_numeric_power_2)
	v325 = int32(_a_F_numeric_power_3)
	v326 = v288 & v325
	if v326 == v325 {
		goto L109
	} else {
		goto L110
	}
L94:
	;
	v293 = int32(1)
	goto L96
L95:
	;
	v293 = int32(-1)
	goto L96
L96:
	;
	v294 = int32(_a_F_numeric_power_3)
	v295 = v28 & v294
	if v295 == v294 {
		v322 = v293
		goto L93
	} else {
		goto L97
	}
L97:
	;
	v298 = int32(0)
	v299 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if v298 <= v29 {
		goto L98
	} else {
		goto L99
	}
L98:
	;
	v306 = int32(-8)
	goto L100
L99:
	;
	v306 = int32(-6)
	goto L100
L100:
	;
	if base.Ui32(int32(base.Ui32(v299)>>(uint(int32(2))%32))+v306) < base.Ui32(int32(2)) {
		v322 = v298
		goto L93
	} else {
		goto L101
	}
L101:
	;
	v311 = int32(1)
	if v295 == int32(_a_F_numeric_power_5) {
		goto L102
	} else {
		goto L103
	}
L102:
	;
	v318 = v28 << (uint(v311) % 32) & int32(_a_F_numeric_power_6)
	goto L104
L103:
	;
	v318 = v295
	goto L104
L104:
	;
	if v318 == int32(_a_F_numeric_power_6) {
		goto L105
	} else {
		goto L106
	}
L105:
	;
	v321 = int32(-1)
	goto L107
L106:
	;
	v321 = v311
	goto L107
L107:
	;
	v322 = v321
	goto L93
L108:
	;
	if base.B2i32(base.Ui32(int32(_a_F_numeric_power_0)) < base.Ui32(v288&int32(_a_F_numeric_power_2)))|base.B2i32(int32(0) <= v322) != 0 {
		goto L129
	} else {
		goto L130
	}
L109:
	;
	if v324 == int32(_a_F_numeric_power_7) {
		goto L112
	} else {
		goto L113
	}
L110:
	;
	v334 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	if int32(0) <= base.I32_extend16_s(v288) {
		goto L115
	} else {
		goto L116
	}
L111:
	;
	v363 = int32(base.Ui32(v361) >> (uint(int32(31)) % 32))
	v365 = base.B2i32(v322 == int32(0))
	if v322 != 0 {
		v368 = v361
		v369 = v365
		v370 = v363
		goto L108
	} else {
		goto L127
	}
L112:
	;
	v333 = int32(1)
	goto L114
L113:
	;
	v333 = int32(-1)
	goto L114
L114:
	;
	v361 = v333
	goto L111
L115:
	;
	v342 = int32(-8)
	goto L117
L116:
	;
	v342 = int32(-6)
	goto L117
L117:
	;
	if base.Ui32(int32(base.Ui32(v334)>>(uint(int32(2))%32))+v342) <= base.Ui32(int32(1)) {
		goto L118
	} else {
		goto L119
	}
L118:
	;
	v346 = int32(0)
	v368 = v346
	v369 = base.B2i32(v322 == v346)
	v370 = int32(0)
	goto L108
L119:
	;
	goto L120
L120:
	;
	v350 = int32(1)
	if v326 == int32(_a_F_numeric_power_5) {
		goto L121
	} else {
		goto L122
	}
L121:
	;
	v357 = v324 << (uint(v350) % 32) & int32(_a_F_numeric_power_6)
	goto L123
L122:
	;
	v357 = v326
	goto L123
L123:
	;
	if v357 == int32(_a_F_numeric_power_6) {
		goto L124
	} else {
		goto L125
	}
L124:
	;
	v360 = int32(-1)
	goto L126
L125:
	;
	v360 = v350
	goto L126
L126:
	;
	v361 = v360
	goto L111
L127:
	;
	if v361 < int32(0) {
		goto L15
	} else {
		goto L128
	}
L128:
	;
	v368 = v361
	v369 = v365
	v370 = v363
	goto L108
L129:
	;
	if base.Ui32(int32(-16385)) < base.Ui32(v29) {
		goto L139
	} else {
		goto L140
	}
L130:
	;
	v378 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v385 = base.B2i32(int32(0) <= base.I32_extend16_s(v288))
	if int32(0) <= base.I32_extend16_s(v288) {
		goto L131
	} else {
		goto L132
	}
L131:
	;
	v386 = int32(-8)
	goto L133
L132:
	;
	v386 = int32(-6)
	goto L133
L133:
	;
	v389 = int32(base.Ui32(int32(base.Ui32(v378)>>(uint(int32(2))%32))+v386) >> (uint(int32(1)) % 32))
	if int32(0) <= base.I32_extend16_s(v288) {
		goto L134
	} else {
		goto L135
	}
L134:
	;
	v390 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+6)))
	v400 = v390
	goto L136
L135:
	;
	v400 = v324<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v324&int32(63)
	goto L136
L136:
	;
	if v389 == int32(0) {
		goto L129
	} else {
		goto L137
	}
L137:
	;
	if v400+int32(1) < v389 {
		goto L14
	} else {
		goto L138
	}
L138:
	;
	goto L129
L139:
	;
	if v368 == int32(0) {
		goto L201
	} else {
		goto L202
	}
L140:
	;
	v411 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v417 = base.B2i32(int32(0) <= v29)
	if int32(0) <= v29 {
		goto L141
	} else {
		goto L142
	}
L141:
	;
	v418 = int32(-8)
	goto L143
L142:
	;
	v418 = int32(-6)
	goto L143
L143:
	;
	if int32(0) <= v29 {
		goto L144
	} else {
		goto L145
	}
L144:
	;
	v420 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
	v430 = v420
	goto L146
L145:
	;
	v430 = v28<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v28&int32(63)
	goto L146
L146:
	;
	v432 = int32(base.Ui32(int32(base.Ui32(v411)>>(uint(int32(2))%32))+v418) >> (uint(int32(1)) % 32))
	if v295 != int32(_a_F_numeric_power_3) {
		goto L148
	} else {
		goto L149
	}
L147:
	;
	if base.B2i32(v432 == int32(0))|v443 != 0 {
		goto L139
	} else {
		goto L152
	}
L148:
	;
	if v295 != int32(_a_F_numeric_power_5) {
		v443 = v295
		goto L147
	} else {
		goto L151
	}
L149:
	;
	goto L150
L150:
	;
	v443 = v28 & int32(_a_F_numeric_power_8)
	goto L147
L151:
	;
	v443 = v28 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_6)
	goto L147
L152:
	;
	if v29 < int32(0) {
		goto L153
	} else {
		goto L154
	}
L153:
	;
	v451 = int32(6)
	goto L155
L154:
	;
	v451 = int32(8)
	goto L155
L155:
	;
	v452 = v21 + v451
	v454 = int32(1)
	v455 = int32(0)
	if base.B2i32(v455 < v430)&base.B2i32(v455 < v432) == v455 {
		v487 = v430
		v491 = v455
		goto L158
	} else {
		goto L159
	}
L156:
	;
	if v629 != 0 {
		goto L139
	} else {
		goto L199
	}
L157:
	;
	v629 = v619
	goto L156
L158:
	;
	if int32(0)|base.B2i32(v455 <= v487) != 0 {
		v523 = v455
		v525 = v455
		goto L165
	} else {
		goto L166
	}
L159:
	;
	v468 = v430
	v472 = v455
	goto L160
L160:
	;
	v478 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452+v472<<(uint(int32(1))%32)))))
	if v478 != 0 {
		v619 = int32(1)
		goto L157
	} else {
		goto L162
	}
L161:
	;
	v487 = v482
	v491 = v480
	goto L158
L162:
	;
	v479 = int32(1)
	v480 = v472 + v479
	v482 = v468 - v479
	if v482 <= v455 {
		v487 = v482
		v491 = v480
		goto L158
	} else {
		goto L163
	}
L163:
	;
	if v480 < v432 {
		v468 = v482
		v472 = v480
		goto L160
	} else {
		goto L164
	}
L164:
	;
	goto L161
L165:
	;
	if v487 != v523 {
		v565 = v491
		v566 = v525
		goto L172
	} else {
		goto L173
	}
L166:
	;
	v504 = v455
	v506 = v455
	goto L167
L167:
	;
	v511 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v506<<(uint(int32(1))%32))+uint32(_c_F_numeric_power[0]))))
	if v511 != 0 {
		v619 = int32(-1)
		goto L157
	} else {
		goto L169
	}
L168:
	;
	v523 = v515
	v525 = v513
	goto L165
L169:
	;
	v512 = int32(1)
	v513 = v506 + v512
	v515 = v504 - v512
	if v515 <= v487 {
		v523 = v515
		v525 = v513
		goto L165
	} else {
		goto L170
	}
L170:
	;
	if v513 < v454 {
		v504 = v515
		v506 = v513
		goto L167
	} else {
		goto L171
	}
L171:
	;
	goto L168
L172:
	;
	if v432 < v565 {
		goto L181
	} else {
		goto L182
	}
L173:
	;
	v534 = v491
	v535 = v525
	goto L174
L174:
	;
	if base.B2i32(v432 <= v534)|base.B2i32(v454 <= v535) != 0 {
		v565 = v534
		v566 = v535
		goto L172
	} else {
		goto L176
	}
L175:
	;
	if base.I32_extend16_s(v551) < base.I32_extend16_s(v549) {
		goto L178
	} else {
		goto L179
	}
L176:
	;
	v540 = int32(1)
	v549 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452+v534<<(uint(v540)%32)))))
	v551 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v535<<(uint(v540)%32))+uint32(_c_F_numeric_power[0]))))
	if v549 == v551 {
		v534 = v534 + v540
		v535 = v535 + v540
		goto L174
	} else {
		goto L177
	}
L177:
	;
	goto L175
L178:
	;
	v558 = int32(1)
	goto L180
L179:
	;
	v558 = int32(-1)
	goto L180
L180:
	;
	v629 = v558
	goto L156
L181:
	;
	v569 = v565
	goto L183
L182:
	;
	v569 = v432
	goto L183
L183:
	;
	v576 = v565
	goto L184
L184:
	;
	if v569 == v576 {
		goto L186
	} else {
		goto L187
	}
L185:
	;
	v619 = v602
	goto L157
L186:
	;
	if v454 < v566 {
		goto L189
	} else {
		goto L190
	}
L187:
	;
	goto L188
L188:
	;
	v602 = int32(1)
	v608 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v452+v576<<(uint(v602)%32)))))
	if v608 == int32(0) {
		v576 = v576 + v602
		goto L184
	} else {
		goto L198
	}
L189:
	;
	v581 = v566
	goto L191
L190:
	;
	v581 = v454
	goto L191
L191:
	;
	v589 = v566
	goto L192
L192:
	;
	if v581 == v589 {
		goto L194
	} else {
		goto L195
	}
L193:
	;
	v619 = int32(-1)
	goto L157
L194:
	;
	v629 = int32(0)
	goto L156
L195:
	;
	goto L196
L196:
	;
	v593 = int32(1)
	v598 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v589<<(uint(v593)%32))+uint32(_c_F_numeric_power[0]))))
	if v598 == int32(0) {
		v589 = v589 + v593
		goto L192
	} else {
		goto L197
	}
L197:
	;
	goto L193
L198:
	;
	goto L185
L199:
	;
	v632 = F_make_result_opt_error(m, int32(_a_F_numeric_power_4), int32(0))
	mBase = m.M
	v633 = m.ExcPending
	if v633 != 0 {
		goto L1
	} else {
		goto L200
	}
L200:
	;
	v2175 = v632
	goto L4
L201:
	;
	v641 = F_make_result_opt_error(m, int32(_a_F_numeric_power_4), int32(0))
	mBase = m.M
	v642 = m.ExcPending
	if v642 != 0 {
		goto L1
	} else {
		goto L204
	}
L202:
	;
	goto L203
L203:
	;
	if base.B2i32(int32(0) < v368)&v369 != 0 {
		goto L205
	} else {
		goto L206
	}
L204:
	;
	v2175 = v641
	goto L4
L205:
	;
	v648 = F_make_result_opt_error(m, int32(_a_F_numeric_power_9), int32(0))
	mBase = m.M
	v649 = m.ExcPending
	if v649 != 0 {
		goto L1
	} else {
		goto L208
	}
L206:
	;
	goto L207
L207:
	;
	if v288&int32(_a_F_numeric_power_10) == int32(_a_F_numeric_power_7) {
		goto L209
	} else {
		goto L210
	}
L208:
	;
	v2175 = v648
	goto L4
L209:
	;
	if base.Ui32(v29) <= base.Ui32(int32(-16385)) {
		goto L212
	} else {
		goto L213
	}
L210:
	;
	goto L211
L211:
	;
	if v29 == int32(-12288) {
		goto L293
	} else {
		goto L294
	}
L212:
	;
	v657 = v18 + int32(48)
	v660 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v665 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	if int32(0) <= v665 {
		goto L216
	} else {
		goto L217
	}
L213:
	;
	v833 = int32(1)
	goto L214
L214:
	;
	if v833 == base.B2i32(int32(0) < v368) {
		goto L288
	} else {
		goto L289
	}
L215:
	;
	v727 = int32(_a_F_numeric_power_11)
	v734 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[1]))
	v735 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[2]))
	v736 = *(*int32)(unsafe.Add(mBase, uint32(v657)))
	if v736 == int32(0) {
		goto L235
	} else {
		goto L236
	}
L216:
	;
	v668 = int32(-8)
	goto L218
L217:
	;
	v668 = int32(-6)
	goto L218
L218:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657))) = int32(base.Ui32(int32(base.Ui32(v660)>>(uint(int32(2))%32))+v668) >> (uint(int32(1)) % 32))
	v673 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	if v673 < int32(0) {
		goto L220
	} else {
		goto L221
	}
L219:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+4)) = v689
	v691 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v692 = int32(_a_F_numeric_power_3)
	v693 = v691 & v692
	if v693 != v692 {
		goto L224
	} else {
		goto L225
	}
L220:
	;
	v677 = v673 & int32(_a_F_numeric_power_2)
	v689 = v677<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v677&int32(63)
	goto L219
L221:
	;
	goto L222
L222:
	;
	v687 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
	v689 = v687
	goto L219
L223:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+8)) = v704
	v706 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	if v706 < int32(0) {
		goto L228
	} else {
		goto L229
	}
L224:
	;
	if v693 != int32(_a_F_numeric_power_5) {
		v704 = v693
		goto L223
	} else {
		goto L227
	}
L225:
	;
	goto L226
L226:
	;
	v704 = v691 & int32(_a_F_numeric_power_8)
	goto L223
L227:
	;
	v704 = v691 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_6)
	goto L223
L228:
	;
	v715 = int32(base.Ui32(v706)>>(uint(int32(7))%32)) & int32(63)
	goto L230
L229:
	;
	v715 = v706 & int32(_a_F_numeric_power_12)
	goto L230
L230:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+12)) = v715
	v717 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+4)))
	v718 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v657)+16)) = v718
	if v717 < v718 {
		goto L231
	} else {
		goto L232
	}
L231:
	;
	v724 = int32(6)
	goto L233
L232:
	;
	v724 = int32(8)
	goto L233
L233:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v657)+20)) = v21 + v724
	goto L215
L234:
	;
	if v772 == int32(0) {
		goto L259
	} else {
		goto L260
	}
L235:
	;
	if v735 == int32(0) {
		goto L238
	} else {
		goto L239
	}
L236:
	;
	goto L237
L237:
	;
	v747 = *(*int32)(unsafe.Add(mBase, uint32(v657)+8))
	if v735 == int32(0) {
		goto L244
	} else {
		goto L245
	}
L238:
	;
	v772 = int32(0)
	goto L234
L239:
	;
	goto L240
L240:
	;
	if v734 == int32(_a_F_numeric_power_6) {
		goto L241
	} else {
		goto L242
	}
L241:
	;
	v746 = int32(1)
	goto L243
L242:
	;
	v746 = int32(-1)
	goto L243
L243:
	;
	v772 = v746
	goto L234
L244:
	;
	if v747 != 0 {
		goto L247
	} else {
		goto L248
	}
L245:
	;
	goto L246
L246:
	;
	v753 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[3]))
	v754 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[4]))
	v755 = *(*int32)(unsafe.Add(mBase, uint32(v657)+4))
	v756 = *(*int32)(unsafe.Add(mBase, uint32(v657)+20))
	if v747 == int32(0) {
		goto L250
	} else {
		goto L251
	}
L247:
	;
	v752 = int32(-1)
	goto L249
L248:
	;
	v752 = int32(1)
	goto L249
L249:
	;
	v772 = v752
	goto L234
L250:
	;
	if v734 == int32(_a_F_numeric_power_6) {
		goto L253
	} else {
		goto L254
	}
L251:
	;
	goto L252
L252:
	;
	if v734 == int32(0) {
		goto L256
	} else {
		goto L257
	}
L253:
	;
	v772 = int32(1)
	goto L234
L254:
	;
	goto L255
L255:
	;
	v762 = F_cmp_abs_common(m, v756, v736, v755, v754, v735, v753)
	mBase = m.M
	v772 = v762
	goto L234
L256:
	;
	v772 = int32(-1)
	goto L234
L257:
	;
	goto L258
L258:
	;
	v766 = F_cmp_abs_common(m, v754, v735, v753, v756, v736, v755)
	mBase = m.M
	v772 = v766
	goto L234
L259:
	;
	v777 = F_make_result_opt_error(m, int32(_a_F_numeric_power_4), int32(0))
	mBase = m.M
	v778 = m.ExcPending
	if v778 != 0 {
		goto L1
	} else {
		goto L262
	}
L260:
	;
	goto L261
L261:
	;
	v779 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v779
	v782 = v18 + int32(48)
	v783 = int32(_a_F_numeric_power_4)
	v790 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[5]))
	v791 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[6]))
	v792 = *(*int32)(unsafe.Add(mBase, uint32(v782)))
	if v792 == v779 {
		goto L264
	} else {
		goto L265
	}
L262:
	;
	v2175 = v777
	goto L4
L263:
	;
	v833 = base.B2i32(int32(0) < v828)
	goto L214
L264:
	;
	if v791 == int32(0) {
		goto L267
	} else {
		goto L268
	}
L265:
	;
	goto L266
L266:
	;
	v803 = *(*int32)(unsafe.Add(mBase, uint32(v782)+8))
	if v791 == int32(0) {
		goto L273
	} else {
		goto L274
	}
L267:
	;
	v828 = int32(0)
	goto L263
L268:
	;
	goto L269
L269:
	;
	if v790 == int32(_a_F_numeric_power_6) {
		goto L270
	} else {
		goto L271
	}
L270:
	;
	v802 = int32(1)
	goto L272
L271:
	;
	v802 = int32(-1)
	goto L272
L272:
	;
	v828 = v802
	goto L263
L273:
	;
	if v803 != 0 {
		goto L276
	} else {
		goto L277
	}
L274:
	;
	goto L275
L275:
	;
	v809 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[7]))
	v810 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[8]))
	v811 = *(*int32)(unsafe.Add(mBase, uint32(v782)+4))
	v812 = *(*int32)(unsafe.Add(mBase, uint32(v782)+20))
	if v803 == int32(0) {
		goto L279
	} else {
		goto L280
	}
L276:
	;
	v808 = int32(-1)
	goto L278
L277:
	;
	v808 = int32(1)
	goto L278
L278:
	;
	v828 = v808
	goto L263
L279:
	;
	if v790 == int32(_a_F_numeric_power_6) {
		goto L282
	} else {
		goto L283
	}
L280:
	;
	goto L281
L281:
	;
	if v790 == int32(0) {
		goto L285
	} else {
		goto L286
	}
L282:
	;
	v828 = int32(1)
	goto L263
L283:
	;
	goto L284
L284:
	;
	v818 = F_cmp_abs_common(m, v812, v792, v811, v810, v791, v809)
	mBase = m.M
	v828 = v818
	goto L263
L285:
	;
	v828 = int32(-1)
	goto L263
L286:
	;
	goto L287
L287:
	;
	v822 = F_cmp_abs_common(m, v810, v791, v809, v812, v792, v811)
	mBase = m.M
	v828 = v822
	goto L263
L288:
	;
	v839 = F_make_result_opt_error(m, int32(_a_F_numeric_power_13), int32(0))
	mBase = m.M
	v840 = m.ExcPending
	if v840 != 0 {
		goto L1
	} else {
		goto L291
	}
L289:
	;
	goto L290
L290:
	;
	v843 = F_make_result_opt_error(m, int32(_a_F_numeric_power_9), int32(0))
	mBase = m.M
	v844 = m.ExcPending
	if v844 != 0 {
		goto L1
	} else {
		goto L292
	}
L291:
	;
	v2175 = v839
	goto L4
L292:
	;
	v2175 = v843
	goto L4
L293:
	;
	if int32(0) < v368 {
		goto L296
	} else {
		goto L297
	}
L294:
	;
	goto L295
L295:
	;
	if v370 != 0 {
		goto L301
	} else {
		goto L302
	}
L296:
	;
	v851 = F_make_result_opt_error(m, int32(_a_F_numeric_power_13), int32(0))
	mBase = m.M
	v852 = m.ExcPending
	if v852 != 0 {
		goto L1
	} else {
		goto L299
	}
L297:
	;
	goto L298
L298:
	;
	v855 = F_make_result_opt_error(m, int32(_a_F_numeric_power_9), int32(0))
	mBase = m.M
	v856 = m.ExcPending
	if v856 != 0 {
		goto L1
	} else {
		goto L300
	}
L299:
	;
	v2175 = v851
	goto L4
L300:
	;
	v2175 = v855
	goto L4
L301:
	;
	v859 = F_make_result_opt_error(m, int32(_a_F_numeric_power_9), int32(0))
	mBase = m.M
	v860 = m.ExcPending
	if v860 != 0 {
		goto L1
	} else {
		goto L304
	}
L302:
	;
	goto L303
L303:
	;
	v862 = v18 + int32(24)
	v865 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v870 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	if int32(0) <= v870 {
		goto L306
	} else {
		goto L307
	}
L304:
	;
	v2175 = v859
	goto L4
L305:
	;
	v932 = *(*int32)(unsafe.Add(mBase, uint32(v18)+24))
	if v932 <= int32(0) {
		goto L324
	} else {
		goto L325
	}
L306:
	;
	v873 = int32(-8)
	goto L308
L307:
	;
	v873 = int32(-6)
	goto L308
L308:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862))) = int32(base.Ui32(int32(base.Ui32(v865)>>(uint(int32(2))%32))+v873) >> (uint(int32(1)) % 32))
	v878 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	if v878 < int32(0) {
		goto L310
	} else {
		goto L311
	}
L309:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862)+4)) = v894
	v896 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	v897 = int32(_a_F_numeric_power_3)
	v898 = v896 & v897
	if v898 != v897 {
		goto L314
	} else {
		goto L315
	}
L310:
	;
	v882 = v878 & int32(_a_F_numeric_power_2)
	v894 = v882<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v882&int32(63)
	goto L309
L311:
	;
	goto L312
L312:
	;
	v892 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+6)))
	v894 = v892
	goto L309
L313:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862)+8)) = v909
	v911 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	if v911 < int32(0) {
		goto L318
	} else {
		goto L319
	}
L314:
	;
	if v898 != int32(_a_F_numeric_power_5) {
		v909 = v898
		goto L313
	} else {
		goto L317
	}
L315:
	;
	goto L316
L316:
	;
	v909 = v896 & int32(_a_F_numeric_power_8)
	goto L313
L317:
	;
	v909 = v896 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_6)
	goto L313
L318:
	;
	v920 = int32(base.Ui32(v911)>>(uint(int32(7))%32)) & int32(63)
	goto L320
L319:
	;
	v920 = v911 & int32(_a_F_numeric_power_12)
	goto L320
L320:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862)+12)) = v920
	v922 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+4)))
	v923 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v862)+16)) = v923
	if v922 < v923 {
		goto L321
	} else {
		goto L322
	}
L321:
	;
	v929 = int32(6)
	goto L323
L322:
	;
	v929 = int32(8)
	goto L323
L323:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v862)+20)) = v26 + v929
	goto L305
L324:
	;
	v956 = F_make_result_opt_error(m, int32(_a_F_numeric_power_13), int32(0))
	mBase = m.M
	v957 = m.ExcPending
	if v957 != 0 {
		goto L1
	} else {
		goto L329
	}
L325:
	;
	v935 = *(*int32)(unsafe.Add(mBase, uint32(v18)+28))
	if v932 != v935+int32(1) {
		goto L324
	} else {
		goto L326
	}
L326:
	;
	v939 = *(*int32)(unsafe.Add(mBase, uint32(v18)+44))
	v940 = int32(1)
	v945 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v939+v932<<(uint(v940)%32)-int32(2)))))
	if v945&v940 == int32(0) {
		goto L324
	} else {
		goto L327
	}
L327:
	;
	v952 = F_make_result_opt_error(m, int32(_a_F_numeric_power_14), int32(0))
	mBase = m.M
	v953 = m.ExcPending
	if v953 != 0 {
		goto L1
	} else {
		goto L328
	}
L328:
	;
	v2175 = v952
	goto L4
L329:
	;
	v2175 = v956
	goto L4
L330:
	;
	v993 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v993
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v993
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v993
	v999 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	v1004 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v21)+4)))
	v1005 = base.I32_extend16_s(v1004)
	v1007 = base.B2i32(int32(0) <= v1005)
	if int32(0) <= v1005 {
		goto L343
	} else {
		goto L344
	}
L331:
	;
	v966 = int32(-8)
	goto L333
L332:
	;
	v966 = int32(-6)
	goto L333
L333:
	;
	if base.Ui32(int32(base.Ui32(v958)>>(uint(int32(2))%32))+v966) < base.Ui32(int32(2)) {
		goto L330
	} else {
		goto L334
	}
L334:
	;
	v970 = *(*int32)(unsafe.Add(mBase, uint32(v21)))
	if int32(0) <= v29 {
		goto L335
	} else {
		goto L336
	}
L335:
	;
	v977 = int32(-8)
	goto L337
L336:
	;
	v977 = int32(-6)
	goto L337
L337:
	;
	if base.Ui32(int32(1)) < base.Ui32(int32(base.Ui32(v970)>>(uint(int32(2))%32))+v977) {
		goto L330
	} else {
		goto L338
	}
L338:
	;
	v986 = v32 & int32(_a_F_numeric_power_3)
	if v986 == int32(_a_F_numeric_power_5) {
		goto L339
	} else {
		goto L340
	}
L339:
	;
	v989 = v32 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_6)
	goto L341
L340:
	;
	v989 = v986
	goto L341
L341:
	;
	if v989 == int32(_a_F_numeric_power_6) {
		goto L13
	} else {
		goto L342
	}
L342:
	;
	goto L330
L343:
	;
	v1008 = int32(-8)
	goto L345
L344:
	;
	v1008 = int32(-6)
	goto L345
L345:
	;
	v1009 = int32(base.Ui32(v999)>>(uint(int32(2))%32)) + v1008
	v1011 = int32(base.Ui32(v1009) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+48)) = v1011
	if int32(0) <= v1005 {
		goto L346
	} else {
		goto L347
	}
L346:
	;
	v1013 = int32(*(*int16)(unsafe.Add(mBase, uint32(v21)+6)))
	v1023 = v1013
	goto L348
L347:
	;
	v1023 = v1004<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v1004&int32(63)
	goto L348
L348:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+52)) = v1023
	v1025 = int32(_a_F_numeric_power_3)
	v1026 = v1004 & v1025
	if v1026 != v1025 {
		goto L350
	} else {
		goto L351
	}
L349:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+56)) = v1037
	v1039 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+64)) = v1039
	v1048 = base.B2i32(v1005 < v1039)
	if v1005 < v1039 {
		goto L354
	} else {
		goto L355
	}
L350:
	;
	if v1026 != int32(_a_F_numeric_power_5) {
		v1037 = v1026
		goto L349
	} else {
		goto L353
	}
L351:
	;
	goto L352
L352:
	;
	v1037 = v1004 & int32(_a_F_numeric_power_8)
	goto L349
L353:
	;
	v1037 = v1004 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_6)
	goto L349
L354:
	;
	v1049 = int32(base.Ui32(v1004)>>(uint(int32(7))%32)) & int32(63)
	goto L356
L355:
	;
	v1049 = v1004 & int32(_a_F_numeric_power_12)
	goto L356
L356:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+60)) = v1049
	if v1005 < v1039 {
		goto L357
	} else {
		goto L358
	}
L357:
	;
	v1053 = int32(6)
	goto L359
L358:
	;
	v1053 = int32(8)
	goto L359
L359:
	;
	v1054 = v21 + v1053
	*(*int32)(unsafe.Add(mBase, uint32(v18)+68)) = v1054
	v1056 = *(*int32)(unsafe.Add(mBase, uint32(v26)))
	v1061 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v26)+4)))
	v1062 = base.I32_extend16_s(v1061)
	v1064 = base.B2i32(int32(0) <= v1062)
	if int32(0) <= v1062 {
		goto L360
	} else {
		goto L361
	}
L360:
	;
	v1065 = int32(-8)
	goto L362
L361:
	;
	v1065 = int32(-6)
	goto L362
L362:
	;
	v1068 = int32(base.Ui32(int32(base.Ui32(v1056)>>(uint(int32(2))%32))+v1065) >> (uint(int32(1)) % 32))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+24)) = v1068
	if int32(0) <= v1062 {
		goto L363
	} else {
		goto L364
	}
L363:
	;
	v1070 = int32(*(*int16)(unsafe.Add(mBase, uint32(v26)+6)))
	v1080 = v1070
	goto L365
L364:
	;
	v1080 = v1061<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v1061&int32(63)
	goto L365
L365:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+28)) = v1080
	v1082 = int32(_a_F_numeric_power_3)
	v1083 = v1061 & v1082
	if v1083 != v1082 {
		goto L367
	} else {
		goto L368
	}
L366:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+32)) = v1094
	v1096 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+40)) = v1096
	v1105 = base.B2i32(v1062 < v1096)
	if v1062 < v1096 {
		goto L371
	} else {
		goto L372
	}
L367:
	;
	if v1083 != int32(_a_F_numeric_power_5) {
		v1094 = v1083
		goto L366
	} else {
		goto L370
	}
L368:
	;
	goto L369
L369:
	;
	v1094 = v1061 & int32(_a_F_numeric_power_8)
	goto L366
L370:
	;
	v1094 = v1061 << (uint(int32(1)) % 32) & int32(_a_F_numeric_power_6)
	goto L366
L371:
	;
	v1106 = int32(base.Ui32(v1061)>>(uint(int32(7))%32)) & int32(63)
	goto L373
L372:
	;
	v1106 = v1061 & int32(_a_F_numeric_power_12)
	goto L373
L373:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+36)) = v1106
	if v1062 < v1096 {
		goto L374
	} else {
		goto L375
	}
L374:
	;
	v1110 = int32(6)
	goto L376
L375:
	;
	v1110 = int32(8)
	goto L376
L376:
	;
	v1111 = v26 + v1110
	*(*int32)(unsafe.Add(mBase, uint32(v18)+44)) = v1111
	v1115 = v1080 + int32(1)
	if v1115 < v1068 {
		goto L378
	} else {
		goto L379
	}
L377:
	;
	if v1011 == int32(0) {
		goto L543
	} else {
		goto L544
	}
L378:
	;
	v1117 = v1068
	goto L380
L379:
	;
	v1117 = int32(0)
	goto L380
L380:
	;
	if v1117 != 0 {
		goto L377
	} else {
		goto L381
	}
L381:
	;
	v1122 = F_numericvar_to_int64(m, v18+int32(24), v18+int32(96))
	mBase = m.M
	v1123 = m.ExcPending
	if v1123 != 0 {
		goto L1
	} else {
		goto L382
	}
L382:
	;
	if v1122 == int32(0) {
		goto L377
	} else {
		goto L383
	}
L383:
	;
	v1126 = *(*int64)(unsafe.Add(mBase, uint32(v18)+96))
	if base.Ui64(int64(4294967295)) < base.Ui64(v1126+int64(2147483648)) {
		goto L377
	} else {
		goto L384
	}
L384:
	;
	v1131 = base.I32_wrap_i64(v1126)
	if v1011 != 0 {
		goto L385
	} else {
		goto L386
	}
L385:
	;
	v1132 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1054))))
	v1133 = base.F64_convert_i32_s(v1132)
	if v1011 == int32(1) {
		goto L388
	} else {
		goto L389
	}
L386:
	;
	v1325 = float64(0)
	goto L387
L387:
	;
	if base.F64_lt(base.F64_add(v1325, float64(1)), float64(-1000)) != 0 {
		goto L414
	} else {
		goto L415
	}
L388:
	;
	v1172 = v1133
	v1173 = v1023 << (uint(int32(2)) % 32)
	goto L390
L389:
	;
	v1138 = int32(2)
	v1140 = v1011 - v1138
	if base.Ui32(v1138) <= base.Ui32(v1140) {
		goto L391
	} else {
		goto L392
	}
L390:
	;
	v1184 = base.I64_reinterpret_f64(v1172)
	if v1184 <= int64(4503599627370495) {
		goto L401
	} else {
		goto L402
	}
L391:
	;
	v1143 = v1138
	goto L393
L392:
	;
	v1143 = v1140
	goto L393
L393:
	;
	v1146 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1054)+2)))
	v1148 = base.F64_add(base.F64_mul(v1133, float64(10000)), base.F64_convert_i32_s(v1146))
	if v1140 == int32(0) {
		v1164 = v1148
		goto L394
	} else {
		goto L395
	}
L394:
	;
	v1172 = v1164
	v1173 = (v1023-v1143)<<(uint(int32(2))%32) - int32(4)
	goto L390
L395:
	;
	v1153 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1054)+4)))
	v1155 = base.F64_add(base.F64_mul(v1148, float64(10000)), base.F64_convert_i32_s(v1153))
	if v1140 == int32(1) {
		v1164 = v1155
		goto L394
	} else {
		goto L396
	}
L396:
	;
	v1160 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1054)+6)))
	v1164 = base.F64_add(base.F64_mul(v1155, float64(10000)), base.F64_convert_i32_s(v1160))
	goto L394
L397:
	;
	v1320 = base.F64_mul(base.F64_add(v1316, base.F64_convert_i32_s(v1173)), base.F64_convert_i32_s(v1131))
	if base.F64_gt(v1320, float64(131072)) != 0 {
		goto L12
	} else {
		goto L413
	}
L398:
	;
	v1316 = v1294
	goto L397
L399:
	;
	v1220 = v1218 + int32(_a_F_numeric_power_15)
	v1224 = base.F64_convert_i32_s(int32(base.Ui32(v1220)>>(uint(int32(20))%32)) + v1217)
	v1226 = base.F64_mul(v1224, float64(0.30102999566361177))
	v1239 = base.F64_add(base.F64_reinterpret_i64(v1215&int64(4294967295)|base.I64_extend_i32_u(v1220&int32(_a_F_numeric_power_16)+int32(1072079006))<<(uint(int64(32))%64)), float64(-1))
	v1242 = base.F64_mul(v1239, base.F64_mul(v1239, float64(0.5)))
	v1247 = base.F64_reinterpret_i64(base.I64_reinterpret_f64(base.F64_sub(v1239, v1242)) & int64(-4294967296))
	v1248 = float64(0.4342944818781689)
	v1249 = base.F64_mul(v1247, v1248)
	v1250 = base.F64_add(v1226, v1249)
	v1255 = base.F64_div(v1239, base.F64_add(v1239, float64(2)))
	v1256 = base.F64_mul(v1255, v1255)
	v1257 = base.F64_mul(v1256, v1256)
	v1282 = base.F64_add(base.F64_mul(v1255, base.F64_add(v1242, base.F64_add(base.F64_mul(v1257, base.F64_add(base.F64_mul(v1257, base.F64_add(base.F64_mul(v1257, float64(0.15313837699209373)), float64(0.22222198432149784))), float64(0.3999999999940942))), base.F64_mul(v1256, base.F64_add(base.F64_mul(v1257, base.F64_add(base.F64_mul(v1257, base.F64_add(base.F64_mul(v1257, float64(0.14798198605116586)), float64(0.1818357216161805))), float64(0.2857142874366239))), float64(0.6666666666666735)))))), base.F64_sub(base.F64_sub(v1239, v1247), v1242))
	v1294 = base.F64_add(v1250, base.F64_add(base.F64_add(v1249, base.F64_sub(v1226, v1250)), base.F64_add(base.F64_mul(v1282, v1248), base.F64_add(base.F64_mul(v1224, float64(3.694239077158931e-13)), base.F64_mul(base.F64_add(v1282, v1247), float64(2.5082946711645275e-11))))))
	goto L398
L400:
	;
	v1211 = base.I64_reinterpret_f64(base.F64_mul(v1172, float64(1.8014398509481984e+16)))
	v1215 = v1211
	v1217 = int32(-1077)
	v1218 = base.I32_wrap_i64(int64(base.Ui64(v1211) >> (uint(int64(32)) % 64)))
	goto L399
L401:
	;
	if base.F64_eq(v1172, float64(0)) != 0 {
		goto L404
	} else {
		goto L405
	}
L402:
	;
	goto L403
L403:
	;
	if base.Ui64(int64(9218868437227405311)) < base.Ui64(v1184) {
		v1294 = v1172
		goto L398
	} else {
		goto L408
	}
L404:
	;
	v1316 = base.F64_div(float64(-1), base.F64_mul(v1172, v1172))
	goto L397
L405:
	;
	goto L406
L406:
	;
	if int64(0) <= v1184 {
		goto L400
	} else {
		goto L407
	}
L407:
	;
	v1316 = base.F64_div(base.F64_sub(v1172, v1172), float64(0))
	goto L397
L408:
	;
	v1199 = int32(-1023)
	v1201 = int64(base.Ui64(v1184) >> (uint(int64(32)) % 64))
	if v1201 != int64(1072693248) {
		goto L409
	} else {
		goto L410
	}
L409:
	;
	v1215 = v1184
	v1217 = v1199
	v1218 = base.I32_wrap_i64(v1201)
	goto L399
L410:
	;
	goto L411
L411:
	;
	if base.I32_wrap_i64(v1184) != 0 {
		v1215 = v1184
		v1217 = v1199
		v1218 = int32(1072693248)
		goto L399
	} else {
		goto L412
	}
L412:
	;
	v1316 = float64(0)
	goto L397
L413:
	;
	v1325 = v1320
	goto L387
L414:
	;
	v1330 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1330 != 0 {
		goto L417
	} else {
		goto L418
	}
L415:
	;
	goto L416
L416:
	;
	v1341 = base.I32_trunc_sat_f64_s(v1325)
	v1342 = int32(16) - v1341
	if v1049 < v1342 {
		goto L421
	} else {
		goto L422
	}
L417:
	;
	F_pfree(m, v1330)
	mBase = m.M
	v1332 = m.ExcPending
	if v1332 != 0 {
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
	v1335 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1335
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v1335
	goto L6
L420:
	;
	goto L419
L421:
	;
	v1344 = v1342
	goto L423
L422:
	;
	v1344 = v1049
	goto L423
L423:
	;
	if base.Ui32(v1106) < base.Ui32(v1344) {
		goto L424
	} else {
		goto L425
	}
L424:
	;
	v1346 = v1344
	goto L426
L425:
	;
	v1346 = v1106
	goto L426
L426:
	;
	if base.Ui32(int32(1000)) <= base.Ui32(v1346) {
		goto L427
	} else {
		goto L428
	}
L427:
	;
	v1349 = int32(1000)
	goto L429
L428:
	;
	v1349 = v1346
	goto L429
L429:
	;
	switch v1131 + int32(1) {
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
	if v1011 == int32(0) {
		goto L476
	} else {
		goto L477
	}
L431:
	;
	v1532 = v18 + int32(48)
	F_mul_var(m, v1532, v1532, v18, v1349)
	mBase = m.M
	v1534 = m.ExcPending
	if v1534 != 0 {
		goto L1
	} else {
		goto L475
	}
L432:
	;
	v1527 = int32(1)
	F_div_var(m, int32(_a_F_numeric_power_4), v18+int32(48), v18, v1349, v1527, v1527)
	mBase = m.M
	v1530 = m.ExcPending
	if v1530 != 0 {
		goto L1
	} else {
		goto L474
	}
L433:
	;
	v1372 = v1009 & int32(-2)
	v1375 = F_palloc(m, v1372+int32(2))
	mBase = m.M
	v1376 = m.ExcPending
	if v1376 != 0 {
		goto L1
	} else {
		goto L440
	}
L434:
	;
	v1353 = F_palloc(m, int32(4))
	mBase = m.M
	v1354 = m.ExcPending
	if v1354 != 0 {
		goto L1
	} else {
		goto L435
	}
L435:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1353))) = int32(_a_F_numeric_power_17)
	v1357 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1357 != 0 {
		goto L436
	} else {
		goto L437
	}
L436:
	;
	F_pfree(m, v1357)
	mBase = m.M
	v1359 = m.ExcPending
	if v1359 != 0 {
		goto L1
	} else {
		goto L439
	}
L437:
	;
	goto L438
L438:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1353 + int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1353
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v1349
	v1366 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[5]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1366
	v1369 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_power[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1369
	goto L6
L439:
	;
	goto L438
L440:
	;
	v1377 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1375))) = uint16(v1377)
	if base.B2i32(v1011 == v1377)|base.B2i32(v1372 == v1377) == v1377 {
		goto L441
	} else {
		goto L442
	}
L441:
	;
	base.MemoryCopy(m, v1375+int32(2), v1054, v1372)
	goto L443
L442:
	;
	goto L443
L443:
	;
	v1389 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1389 != 0 {
		goto L444
	} else {
		goto L445
	}
L444:
	;
	F_pfree(m, v1389)
	mBase = m.M
	v1391 = m.ExcPending
	if v1391 != 0 {
		goto L1
	} else {
		goto L447
	}
L445:
	;
	goto L446
L446:
	;
	v1392 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v1392
	v1394 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1394
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1375
	v1397 = int32(2)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1375 + v1397
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v1349
	v1406 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v1409 = v1349 + v1406<<(uint(v1397)%32)
	if v1409+int32(4) < int32(0) {
		goto L449
	} else {
		goto L450
	}
L447:
	;
	goto L446
L448:
	;
	goto L6
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
	v1418 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v1420 = v1349 & int32(3)
	v1424 = base.I32_div_s(v1409+int32(7), int32(4))
	v1425 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v1425 <= v1424 {
		goto L456
	} else {
		goto L457
	}
L452:
	;
	goto L448
L453:
	;
	if int32(0) <= v1490 {
		goto L452
	} else {
		goto L473
	}
L454:
	;
	v1470 = v1464
	goto L467
L455:
	;
	v1439 = int32(1)
	v1440 = v1424 - v1439
	v1443 = v1418 + v1440<<(uint(v1439)%32)
	v1444 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1443))))
	v1445 = int32(2)
	v1447 = *(*int32)(unsafe.Add(mBase, uint32(v1420<<(uint(v1445)%32))+uint32(_c_F_numeric_power[9])))
	v1448 = base.I32_rem_s(v1444, v1447)
	v1449 = v1444 - v1448
	*(*uint16)(unsafe.Add(mBase, uint32(v1443))) = uint16(v1449)
	v1452 = base.I32_div_s(v1447, v1445)
	if v1448 < v1452 {
		v1490 = v1440
		goto L453
	} else {
		goto L462
	}
L456:
	;
	if base.B2i32(v1420 == int32(0))|base.B2i32(v1424 != v1425) != 0 {
		goto L452
	} else {
		goto L459
	}
L457:
	;
	goto L458
L458:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1424
	if v1420 != 0 {
		goto L455
	} else {
		goto L460
	}
L459:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1424
	goto L455
L460:
	;
	v1436 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1418+v1424<<(uint(int32(1))%32)))))
	if v1436 <= int32(_a_F_numeric_power_18) {
		v1490 = v1424
		goto L453
	} else {
		goto L461
	}
L461:
	;
	v1464 = v1424
	goto L454
L462:
	;
	v1455 = v1447 + base.I32_extend16_s(v1449)
	if int32(_a_F_numeric_power_19) < v1455 {
		goto L463
	} else {
		goto L464
	}
L463:
	;
	v1460 = v1455 + int32(_a_F_numeric_power_20)
	goto L465
L464:
	;
	v1460 = v1455
	goto L465
L465:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v1443))) = uint16(v1460)
	if v1455 < int32(_a_F_numeric_power_21) {
		v1490 = v1440
		goto L453
	} else {
		goto L466
	}
L466:
	;
	v1464 = v1440
	goto L454
L467:
	;
	v1476 = int32(1)
	v1477 = v1470 - v1476
	v1480 = v1418 + v1477<<(uint(v1476)%32)
	v1483 = int32(*(*int16)(unsafe.Add(mBase, uint32(v1480))))
	v1485 = base.B2i32(int32(_a_F_numeric_power_22) < v1483)
	if int32(_a_F_numeric_power_22) < v1483 {
		goto L469
	} else {
		goto L470
	}
L468:
	;
	v1490 = v1477
	goto L453
L469:
	;
	v1486 = int32(-9999)
	goto L471
L470:
	;
	v1486 = v1476
	goto L471
L471:
	;
	v1487 = v1486 + v1483
	*(*uint16)(unsafe.Add(mBase, uint32(v1480))) = uint16(v1487)
	if int32(_a_F_numeric_power_22) < v1483 {
		v1470 = v1477
		goto L467
	} else {
		goto L472
	}
L472:
	;
	goto L468
L473:
	;
	v1498 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1498 - int32(2)
	v1502 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v1503 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v1502 + v1503
	v1506 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v1506 + v1503
	goto L452
L474:
	;
	goto L6
L475:
	;
	goto L6
L476:
	;
	if v1126 < int64(0) {
		goto L11
	} else {
		goto L479
	}
L477:
	;
	goto L478
L478:
	;
	v1550 = v1009 & int32(-2)
	v1552 = v1550 + int32(2)
	v1553 = F_palloc(m, v1552)
	mBase = m.M
	v1554 = m.ExcPending
	if v1554 != 0 {
		goto L1
	} else {
		goto L484
	}
L479:
	;
	v1539 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1539 != 0 {
		goto L480
	} else {
		goto L481
	}
L480:
	;
	F_pfree(m, v1539)
	mBase = m.M
	v1541 = m.ExcPending
	if v1541 != 0 {
		goto L1
	} else {
		goto L483
	}
L481:
	;
	goto L482
L482:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+12)) = v1349
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
	v1545 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1545
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v1545
	goto L6
L483:
	;
	goto L482
L484:
	;
	v1555 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1553))) = uint16(v1555)
	v1558 = v1553 + int32(2)
	v1560 = base.B2i32(v1550 == v1555)
	if v1560 == v1555 {
		goto L485
	} else {
		goto L486
	}
L485:
	;
	base.MemoryCopy(m, v1558, v1054, v1550)
	goto L487
L486:
	;
	goto L487
L487:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v1558
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = v1553
	v1566 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v1566
	v1568 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v1568
	v1572 = F_log(m, base.F64_abs(base.F64_convert_i32_s(v1131)))
	mBase = m.M
	v1577 = v1131 >> (uint(int32(31)) % 32)
	v1579 = v1131 ^ v1577 - v1577
	if v1579&int32(1) != 0 {
		goto L489
	} else {
		goto L490
	}
L488:
	;
	v1618 = base.I32_trunc_sat_f64_s(v1572) + v1341 + v1349 + int32(9)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1615
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1614
	v1623 = v1579
	goto L506
L489:
	;
	v1582 = F_palloc(m, v1552)
	mBase = m.M
	v1583 = m.ExcPending
	if v1583 != 0 {
		goto L1
	} else {
		goto L492
	}
L490:
	;
	goto L491
L491:
	;
	v1599 = F_palloc(m, int32(4))
	mBase = m.M
	v1600 = m.ExcPending
	if v1600 != 0 {
		goto L1
	} else {
		goto L500
	}
L492:
	;
	v1584 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1582))) = uint16(v1584)
	v1587 = v1582 + int32(2)
	if v1560 == v1584 {
		goto L493
	} else {
		goto L494
	}
L493:
	;
	base.MemoryCopy(m, v1587, v1054, v1550)
	goto L495
L494:
	;
	goto L495
L495:
	;
	v1591 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1591 != 0 {
		goto L496
	} else {
		goto L497
	}
L496:
	;
	F_pfree(m, v1591)
	mBase = m.M
	v1593 = m.ExcPending
	if v1593 != 0 {
		goto L1
	} else {
		goto L499
	}
L497:
	;
	goto L498
L498:
	;
	v1594 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v1594
	v1596 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1596
	v1614 = v1582
	v1615 = v1587
	goto L488
L499:
	;
	goto L498
L500:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v1599))) = int32(_a_F_numeric_power_17)
	v1603 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1603 != 0 {
		goto L501
	} else {
		goto L502
	}
L501:
	;
	F_pfree(m, v1603)
	mBase = m.M
	v1605 = m.ExcPending
	if v1605 != 0 {
		goto L1
	} else {
		goto L504
	}
L502:
	;
	goto L503
L503:
	;
	v1609 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_power[5]))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+8)) = v1609
	v1612 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_power[6]))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1612
	v1614 = v1599
	v1615 = v1599 + int32(2)
	goto L488
L504:
	;
	goto L503
L505:
	;
	v1700 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
	if v1700 != 0 {
		goto L537
	} else {
		goto L538
	}
L506:
	;
	v1637 = int32(base.Ui32(v1623) >> (uint(int32(1)) % 32))
	if v1637 == int32(0) {
		goto L505
	} else {
		goto L508
	}
L507:
	;
	if int64(0) <= v1126 {
		goto L10
	} else {
		goto L530
	}
L508:
	;
	v1641 = v18 + int32(120)
	v1642 = *(*int32)(unsafe.Add(mBase, uint32(v18)+124))
	v1645 = v1618 - v1642<<(uint(int32(3))%32)
	v1646 = *(*int32)(unsafe.Add(mBase, uint32(v18)+132))
	v1648 = v1646 << (uint(int32(1)) % 32)
	if v1645 < v1648 {
		goto L509
	} else {
		goto L510
	}
L509:
	;
	v1650 = v1645
	goto L511
L510:
	;
	v1650 = v1648
	goto L511
L511:
	;
	v1651 = int32(0)
	if v1651 < v1650 {
		goto L512
	} else {
		goto L513
	}
L512:
	;
	v1654 = v1650
	goto L514
L513:
	;
	v1654 = v1651
	goto L514
L514:
	;
	F_mul_var(m, v1641, v1641, v1641, v1654)
	mBase = m.M
	v1656 = m.ExcPending
	if v1656 != 0 {
		goto L1
	} else {
		goto L515
	}
L515:
	;
	v1657 = *(*int32)(unsafe.Add(mBase, uint32(v18)+124))
	if v1623&int32(2) != 0 {
		goto L516
	} else {
		goto L517
	}
L516:
	;
	v1660 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	v1664 = v1618 - (v1660+v1657)<<(uint(int32(2))%32)
	v1665 = *(*int32)(unsafe.Add(mBase, uint32(v18)+12))
	v1666 = *(*int32)(unsafe.Add(mBase, uint32(v18)+132))
	v1667 = v1665 + v1666
	if v1664 < v1667 {
		goto L519
	} else {
		goto L520
	}
L517:
	;
	goto L518
L518:
	;
	if v1657 <= int32(_a_F_numeric_power_23) {
		goto L526
	} else {
		goto L527
	}
L519:
	;
	v1669 = v1664
	goto L521
L520:
	;
	v1669 = v1667
	goto L521
L521:
	;
	v1670 = int32(0)
	if v1670 < v1669 {
		goto L522
	} else {
		goto L523
	}
L522:
	;
	v1673 = v1669
	goto L524
L523:
	;
	v1673 = v1670
	goto L524
L524:
	;
	F_mul_var(m, v1641, v18, v18, v1673)
	mBase = m.M
	v1675 = m.ExcPending
	if v1675 != 0 {
		goto L1
	} else {
		goto L525
	}
L525:
	;
	goto L518
L526:
	;
	v1680 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	if v1680 < int32(_a_F_numeric_power_5) {
		v1623 = v1637
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
	v1686 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1686 != 0 {
		goto L531
	} else {
		goto L532
	}
L531:
	;
	F_pfree(m, v1686)
	mBase = m.M
	v1688 = m.ExcPending
	if v1688 != 0 {
		goto L1
	} else {
		goto L534
	}
L532:
	;
	goto L533
L533:
	;
	v1689 = int32(0)
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1689
	v1691 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1691
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v1691
	v1695 = *(*int32)(unsafe.Add(mBase, uint32(v18)+136))
	if v1695 == v1689 {
		goto L7
	} else {
		goto L535
	}
L534:
	;
	goto L533
L535:
	;
	F_pfree(m, v1695)
	mBase = m.M
	v1699 = m.ExcPending
	if v1699 != 0 {
		goto L1
	} else {
		goto L536
	}
L536:
	;
	goto L7
L537:
	;
	F_pfree(m, v1700)
	mBase = m.M
	v1702 = m.ExcPending
	if v1702 != 0 {
		goto L1
	} else {
		goto L540
	}
L538:
	;
	goto L539
L539:
	;
	if int64(0) <= v1126 {
		goto L7
	} else {
		goto L541
	}
L540:
	;
	goto L539
L541:
	;
	F_div_var(m, int32(_a_F_numeric_power_4), v18, v18, v1349, int32(1), int32(0))
	mBase = m.M
	v1709 = m.ExcPending
	if v1709 != 0 {
		goto L1
	} else {
		goto L542
	}
L542:
	;
	goto L6
L543:
	;
	v1714 = F_palloc(m, int32(2))
	mBase = m.M
	v1715 = m.ExcPending
	if v1715 != 0 {
		goto L1
	} else {
		goto L546
	}
L544:
	;
	goto L545
L545:
	;
	v1733 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18)+136)) = v1733
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v1733
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v1733
	*(*int64)(unsafe.Add(mBase, uint32(v18)+96)) = v1733
	*(*int64)(unsafe.Add(mBase, uint32(v18)+104)) = v1733
	*(*int64)(unsafe.Add(mBase, uint32(v18)+112)) = v1733
	*(*int64)(unsafe.Add(mBase, uint32(v18)+72)) = v1733
	*(*int64)(unsafe.Add(mBase, uint32(v18)+80)) = v1733
	*(*int64)(unsafe.Add(mBase, uint32(v18)+88)) = v1733
	if v1037 != int32(_a_F_numeric_power_6) {
		goto L552
	} else {
		goto L553
	}
L546:
	;
	v1716 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1714))) = uint16(v1716)
	v1718 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1718 != 0 {
		goto L547
	} else {
		goto L548
	}
L547:
	;
	F_pfree(m, v1718)
	mBase = m.M
	v1720 = m.ExcPending
	if v1720 != 0 {
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
	v1724 = *(*int32)(unsafe.Add(mBase, _c_F_numeric_power[10]))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = v1724
	v1727 = *(*int64)(unsafe.Add(mBase, _c_F_numeric_power[11]))
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1727
	*(*int32)(unsafe.Add(mBase, uint32(v18)+16)) = v1714
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v1714 + int32(2)
	goto L6
L550:
	;
	goto L549
L551:
	;
	v1799 = v18 + int32(96)
	v1801 = F_estimate_ln_dweight(m, v1797)
	mBase = m.M
	v1802 = m.ExcPending
	if v1802 != 0 {
		goto L1
	} else {
		goto L565
	}
L552:
	;
	v1753 = int32(0)
	v1793 = v1753
	v1794 = v1753
	v1797 = v18 + int32(48)
	goto L551
L553:
	;
	goto L554
L554:
	;
	if v1068 == int32(0) {
		goto L556
	} else {
		goto L557
	}
L555:
	;
	v1773 = v1009 & int32(-2)
	v1776 = F_palloc(m, v1773+int32(2))
	mBase = m.M
	v1777 = m.ExcPending
	if v1777 != 0 {
		goto L1
	} else {
		goto L561
	}
L556:
	;
	v1771 = int32(0)
	goto L555
L557:
	;
	if v1115 < v1068 {
		goto L9
	} else {
		goto L558
	}
L558:
	;
	if v1115 != v1068 {
		goto L556
	} else {
		goto L559
	}
L559:
	;
	v1761 = int32(1)
	v1767 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v1111+v1068<<(uint(v1761)%32)-int32(2)))))
	if v1767&v1761 != 0 {
		v1771 = v1761
		goto L555
	} else {
		goto L560
	}
L560:
	;
	goto L556
L561:
	;
	v1778 = int32(0)
	*(*uint16)(unsafe.Add(mBase, uint32(v1776))) = uint16(v1778)
	v1781 = v1776 + int32(2)
	if v1773 != 0 {
		goto L562
	} else {
		goto L563
	}
L562:
	;
	base.MemoryCopy(m, v1781, v1054, v1773)
	goto L564
L563:
	;
	goto L564
L564:
	;
	v1783 = *(*int64)(unsafe.Add(mBase, uint32(v18)+56))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+128)) = v1783
	*(*int32)(unsafe.Add(mBase, uint32(v18)+140)) = v1781
	*(*int32)(unsafe.Add(mBase, uint32(v18)+136)) = v1776
	v1787 = *(*int64)(unsafe.Add(mBase, uint32(v18)+48))
	*(*int64)(unsafe.Add(mBase, uint32(v18)+120)) = v1787
	*(*int32)(unsafe.Add(mBase, uint32(v18)+128)) = int32(0)
	v1793 = v1776
	v1794 = v1771
	v1797 = v18 + int32(120)
	goto L551
L565:
	;
	v1803 = int32(8) - v1801
	v1804 = int32(0)
	if v1804 < v1803 {
		goto L566
	} else {
		goto L567
	}
L566:
	;
	v1807 = v1803
	goto L568
L567:
	;
	v1807 = v1804
	goto L568
L568:
	;
	F_ln_var(m, v1797, v1799, v1807)
	mBase = m.M
	v1809 = m.ExcPending
	if v1809 != 0 {
		goto L1
	} else {
		goto L569
	}
L569:
	;
	v1813 = v18 + int32(72)
	F_mul_var(m, v1799, v18+int32(24), v1813, v1807)
	mBase = m.M
	v1815 = m.ExcPending
	if v1815 != 0 {
		goto L1
	} else {
		goto L570
	}
L570:
	;
	v1816 = F_numericvar_to_double_no_overflow(m, v1813)
	mBase = m.M
	v1817 = m.ExcPending
	if v1817 != 0 {
		goto L1
	} else {
		goto L571
	}
L571:
	;
	if base.F64_gt(base.F64_abs(v1816), float64(6020)) != 0 {
		goto L572
	} else {
		goto L573
	}
L572:
	;
	if base.F64_gt(v1816, float64(0)) != 0 {
		goto L8
	} else {
		goto L575
	}
L573:
	;
	goto L574
L574:
	;
	v1833 = v18 + int32(96)
	v1838 = base.I32_trunc_sat_f64_s(base.F64_mul(v1816, float64(0.434294481903252)))
	v1839 = int32(16) - v1838
	v1840 = *(*int32)(unsafe.Add(mBase, uint32(v1797)+12))
	if v1840 < v1839 {
		goto L580
	} else {
		goto L581
	}
L575:
	;
	v1823 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v1823 != 0 {
		goto L576
	} else {
		goto L577
	}
L576:
	;
	F_pfree(m, v1823)
	mBase = m.M
	v1825 = m.ExcPending
	if v1825 != 0 {
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
	v1828 = int64(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = v1828
	*(*int64)(unsafe.Add(mBase, uint32(v18)+16)) = v1828
	goto L6
L579:
	;
	goto L578
L580:
	;
	v1842 = v1839
	goto L582
L581:
	;
	v1842 = v1840
	goto L582
L582:
	;
	if v1106 < v1842 {
		goto L583
	} else {
		goto L584
	}
L583:
	;
	v1844 = v1842
	goto L585
L584:
	;
	v1844 = v1106
	goto L585
L585:
	;
	if int32(1000) <= v1844 {
		goto L586
	} else {
		goto L587
	}
L586:
	;
	v1847 = int32(1000)
	goto L588
L587:
	;
	v1847 = v1844
	goto L588
L588:
	;
	v1848 = v1847 + v1838
	v1849 = int32(0)
	if v1849 < v1848 {
		goto L589
	} else {
		goto L590
	}
L589:
	;
	v1852 = v1848
	goto L591
L590:
	;
	v1852 = v1849
	goto L591
L591:
	;
	v1855 = v1852 - v1801 + int32(8)
	v1856 = int32(0)
	if v1856 < v1855 {
		goto L592
	} else {
		goto L593
	}
L592:
	;
	v1859 = v1855
	goto L594
L593:
	;
	v1859 = v1856
	goto L594
L594:
	;
	F_ln_var(m, v1797, v1833, v1859)
	mBase = m.M
	v1861 = m.ExcPending
	if v1861 != 0 {
		goto L1
	} else {
		goto L595
	}
L595:
	;
	v1865 = v18 + int32(72)
	F_mul_var(m, v1833, v18+int32(24), v1865, v1859)
	mBase = m.M
	v1867 = m.ExcPending
	if v1867 != 0 {
		goto L1
	} else {
		goto L596
	}
L596:
	;
	F_exp_var(m, v1865, v18, v1847)
	mBase = m.M
	v1869 = m.ExcPending
	if v1869 != 0 {
		goto L1
	} else {
		goto L597
	}
L597:
	;
	v1870 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if base.B2i32(int32(0) < v1870)&v1794 != 0 {
		goto L598
	} else {
		goto L599
	}
L598:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(_a_F_numeric_power_6)
	goto L600
L599:
	;
	goto L600
L600:
	;
	v1876 = *(*int32)(unsafe.Add(mBase, uint32(v18)+88))
	if v1876 != 0 {
		goto L601
	} else {
		goto L602
	}
L601:
	;
	F_pfree(m, v1876)
	mBase = m.M
	v1878 = m.ExcPending
	if v1878 != 0 {
		goto L1
	} else {
		goto L604
	}
L602:
	;
	goto L603
L603:
	;
	v1879 = *(*int32)(unsafe.Add(mBase, uint32(v18)+112))
	if v1879 != 0 {
		goto L605
	} else {
		goto L606
	}
L604:
	;
	goto L603
L605:
	;
	F_pfree(m, v1879)
	mBase = m.M
	v1881 = m.ExcPending
	if v1881 != 0 {
		goto L1
	} else {
		goto L608
	}
L606:
	;
	goto L607
L607:
	;
	if v1793 == int32(0) {
		goto L6
	} else {
		goto L609
	}
L608:
	;
	goto L607
L609:
	;
	F_pfree(m, v1793)
	mBase = m.M
	v1885 = m.ExcPending
	if v1885 != 0 {
		goto L1
	} else {
		goto L610
	}
L610:
	;
	goto L6
L611:
	;
	F_errcode(m, int32(369361026))
	mBase = m.M
	v1892 = m.ExcPending
	if v1892 != 0 {
		goto L1
	} else {
		goto L612
	}
L612:
	;
	F_errmsg(m, int32(_a_F_numeric_power_24), int32(0))
	mBase = m.M
	v1896 = m.ExcPending
	if v1896 != 0 {
		goto L1
	} else {
		goto L613
	}
L613:
	;
	F_errfinish(m, int32(_a_F_numeric_power_25), int32(_a_F_numeric_power_26), int32(_a_F_numeric_power_27))
	mBase = m.M
	v1901 = m.ExcPending
	if v1901 != 0 {
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
	F_errcode(m, int32(369361026))
	mBase = m.M
	v1908 = m.ExcPending
	if v1908 != 0 {
		goto L1
	} else {
		goto L616
	}
L616:
	;
	F_errmsg(m, int32(_a_F_numeric_power_28), int32(0))
	mBase = m.M
	v1912 = m.ExcPending
	if v1912 != 0 {
		goto L1
	} else {
		goto L617
	}
L617:
	;
	F_errfinish(m, int32(_a_F_numeric_power_25), int32(_a_F_numeric_power_29), int32(_a_F_numeric_power_27))
	mBase = m.M
	v1917 = m.ExcPending
	if v1917 != 0 {
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
	F_errcode(m, int32(369361026))
	mBase = m.M
	v1924 = m.ExcPending
	if v1924 != 0 {
		goto L1
	} else {
		goto L620
	}
L620:
	;
	F_errmsg(m, int32(_a_F_numeric_power_24), int32(0))
	mBase = m.M
	v1928 = m.ExcPending
	if v1928 != 0 {
		goto L1
	} else {
		goto L621
	}
L621:
	;
	F_errfinish(m, int32(_a_F_numeric_power_25), int32(_a_F_numeric_power_30), int32(_a_F_numeric_power_27))
	mBase = m.M
	v1933 = m.ExcPending
	if v1933 != 0 {
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
	v1940 = m.ExcPending
	if v1940 != 0 {
		goto L1
	} else {
		goto L624
	}
L624:
	;
	F_errmsg(m, int32(_a_F_numeric_power_31), int32(0))
	mBase = m.M
	v1944 = m.ExcPending
	if v1944 != 0 {
		goto L1
	} else {
		goto L625
	}
L625:
	;
	F_errfinish(m, int32(_a_F_numeric_power_25), int32(_a_F_numeric_power_32), int32(_a_F_numeric_power_33))
	mBase = m.M
	v1949 = m.ExcPending
	if v1949 != 0 {
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
	F_errcode(m, int32(33816706))
	mBase = m.M
	v1956 = m.ExcPending
	if v1956 != 0 {
		goto L1
	} else {
		goto L628
	}
L628:
	;
	F_errmsg(m, int32(_a_F_numeric_power_34), int32(0))
	mBase = m.M
	v1960 = m.ExcPending
	if v1960 != 0 {
		goto L1
	} else {
		goto L629
	}
L629:
	;
	F_errfinish(m, int32(_a_F_numeric_power_25), int32(_a_F_numeric_power_35), int32(_a_F_numeric_power_33))
	mBase = m.M
	v1965 = m.ExcPending
	if v1965 != 0 {
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
	v1972 = m.ExcPending
	if v1972 != 0 {
		goto L1
	} else {
		goto L632
	}
L632:
	;
	F_errmsg(m, int32(_a_F_numeric_power_31), int32(0))
	mBase = m.M
	v1976 = m.ExcPending
	if v1976 != 0 {
		goto L1
	} else {
		goto L633
	}
L633:
	;
	F_errfinish(m, int32(_a_F_numeric_power_25), int32(_a_F_numeric_power_36), int32(_a_F_numeric_power_33))
	mBase = m.M
	v1981 = m.ExcPending
	if v1981 != 0 {
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
	F_errcode(m, int32(369361026))
	mBase = m.M
	v1988 = m.ExcPending
	if v1988 != 0 {
		goto L1
	} else {
		goto L636
	}
L636:
	;
	F_errmsg(m, int32(_a_F_numeric_power_28), int32(0))
	mBase = m.M
	v1992 = m.ExcPending
	if v1992 != 0 {
		goto L1
	} else {
		goto L637
	}
L637:
	;
	F_errfinish(m, int32(_a_F_numeric_power_25), int32(_a_F_numeric_power_37), int32(_a_F_numeric_power_38))
	mBase = m.M
	v1997 = m.ExcPending
	if v1997 != 0 {
		goto L1
	} else {
		goto L638
	}
L638:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L639:
	;
	F_errcode(m, int32(50331778))
	mBase = m.M
	v2004 = m.ExcPending
	if v2004 != 0 {
		goto L1
	} else {
		goto L640
	}
L640:
	;
	F_errmsg(m, int32(_a_F_numeric_power_31), int32(0))
	mBase = m.M
	v2008 = m.ExcPending
	if v2008 != 0 {
		goto L1
	} else {
		goto L641
	}
L641:
	;
	F_errfinish(m, int32(_a_F_numeric_power_25), int32(_a_F_numeric_power_39), int32(_a_F_numeric_power_38))
	mBase = m.M
	v2013 = m.ExcPending
	if v2013 != 0 {
		goto L1
	} else {
		goto L642
	}
L642:
	;
	base.Wasm_trap_unreachable()
	for {
	}
L643:
	;
	goto L6
L644:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18)+8)) = int32(0)
	*(*int64)(unsafe.Add(mBase, uint32(v18))) = int64(0)
	goto L643
L645:
	;
	goto L646
L646:
	;
	v2037 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	v2039 = v1349 & int32(3)
	v2043 = base.I32_div_s(v2028+int32(7), int32(4))
	v2044 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	if v2044 <= v2043 {
		goto L651
	} else {
		goto L652
	}
L647:
	;
	goto L643
L648:
	;
	if int32(0) <= v2109 {
		goto L647
	} else {
		goto L668
	}
L649:
	;
	v2089 = v2083
	goto L662
L650:
	;
	v2058 = int32(1)
	v2059 = v2043 - v2058
	v2062 = v2037 + v2059<<(uint(v2058)%32)
	v2063 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2062))))
	v2064 = int32(2)
	v2066 = *(*int32)(unsafe.Add(mBase, uint32(v2039<<(uint(v2064)%32))+uint32(_c_F_numeric_power[9])))
	v2067 = base.I32_rem_s(v2063, v2066)
	v2068 = v2063 - v2067
	*(*uint16)(unsafe.Add(mBase, uint32(v2062))) = uint16(v2068)
	v2071 = base.I32_div_s(v2066, v2064)
	if v2067 < v2071 {
		v2109 = v2059
		goto L648
	} else {
		goto L657
	}
L651:
	;
	if base.B2i32(v2039 == int32(0))|base.B2i32(v2043 != v2044) != 0 {
		goto L647
	} else {
		goto L654
	}
L652:
	;
	goto L653
L653:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v2043
	if v2039 != 0 {
		goto L650
	} else {
		goto L655
	}
L654:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v2043
	goto L650
L655:
	;
	v2055 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2037+v2043<<(uint(int32(1))%32)))))
	if v2055 <= int32(_a_F_numeric_power_18) {
		v2109 = v2043
		goto L648
	} else {
		goto L656
	}
L656:
	;
	v2083 = v2043
	goto L649
L657:
	;
	v2074 = v2066 + base.I32_extend16_s(v2068)
	if int32(_a_F_numeric_power_19) < v2074 {
		goto L658
	} else {
		goto L659
	}
L658:
	;
	v2079 = v2074 + int32(_a_F_numeric_power_20)
	goto L660
L659:
	;
	v2079 = v2074
	goto L660
L660:
	;
	*(*uint16)(unsafe.Add(mBase, uint32(v2062))) = uint16(v2079)
	if v2074 < int32(_a_F_numeric_power_21) {
		v2109 = v2059
		goto L648
	} else {
		goto L661
	}
L661:
	;
	v2083 = v2059
	goto L649
L662:
	;
	v2095 = int32(1)
	v2096 = v2089 - v2095
	v2099 = v2037 + v2096<<(uint(v2095)%32)
	v2102 = int32(*(*int16)(unsafe.Add(mBase, uint32(v2099))))
	v2104 = base.B2i32(int32(_a_F_numeric_power_22) < v2102)
	if int32(_a_F_numeric_power_22) < v2102 {
		goto L664
	} else {
		goto L665
	}
L663:
	;
	v2109 = v2096
	goto L648
L664:
	;
	v2105 = int32(-9999)
	goto L666
L665:
	;
	v2105 = v2095
	goto L666
L666:
	;
	v2106 = v2105 + v2102
	*(*uint16)(unsafe.Add(mBase, uint32(v2099))) = uint16(v2106)
	if int32(_a_F_numeric_power_22) < v2102 {
		v2089 = v2096
		goto L662
	} else {
		goto L667
	}
L667:
	;
	goto L663
L668:
	;
	v2117 = *(*int32)(unsafe.Add(mBase, uint32(v18)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+20)) = v2117 - int32(2)
	v2121 = *(*int32)(unsafe.Add(mBase, uint32(v18)))
	v2122 = int32(1)
	*(*int32)(unsafe.Add(mBase, uint32(v18))) = v2121 + v2122
	v2125 = *(*int32)(unsafe.Add(mBase, uint32(v18)+4))
	*(*int32)(unsafe.Add(mBase, uint32(v18)+4)) = v2125 + v2122
	goto L647
L669:
	;
	v2161 = *(*int32)(unsafe.Add(mBase, uint32(v18)+16))
	if v2161 == int32(0) {
		v2175 = v2159
		goto L4
	} else {
		goto L670
	}
L670:
	;
	F_pfree(m, v2161)
	mBase = m.M
	v2165 = m.ExcPending
	if v2165 != 0 {
		goto L1
	} else {
		goto L671
	}
L671:
	;
	v2175 = v2159
	goto L4
L672:
	;
	v2175 = v2171
	goto L4
}
func F_numeric_random(m *base.Module, l0 int32) int32 {
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
	v27 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
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
	return int32(0)
L2:
	;
	v32 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
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
	return v1502
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
	v1502 = F_make_result_opt_error(m, v165, int32(0))
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
func F_numeric_sortsupport(m *base.Module, l0 int32) int32 {
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
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(1449)
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
			return int32(0)
		} else {
			v22 = F_palloc(m, int32(132))
			mBase = m.M
			v23 = m.ExcPending
			if v23 != 0 {
				return int32(0)
			} else {
				v24 = int32(1)
				*(*uint8)(unsafe.Add(mBase, uint32(v17)+16)) = uint8(v24)
				*(*int64)(unsafe.Add(mBase, uint32(v17)+8)) = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v17))) = v22
				F_initHyperLogLog(m, v17+int32(24), int32(10))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int32(0)
				} else {
					*(*int32)(unsafe.Add(mBase, uint32(v5)+12)) = v17
					*(*int32)(unsafe.Add(mBase, uint32(v5)+28)) = int32(1450)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+24)) = int32(1451)
					v39 = *(*int32)(unsafe.Add(mBase, uint32(v5)+16))
					*(*int32)(unsafe.Add(mBase, uint32(v5)+16)) = int32(1452)
					*(*int32)(unsafe.Add(mBase, uint32(v5)+32)) = v39
					*(*int32)(unsafe.Add(mBase, _c_F_numeric_sortsupport[0])) = v12
					return int32(0)
				}
			}
		}
	} else {
		return int32(0)
	}
}
func F_numeric_sub_opt_error(m *base.Module, l0 int32, l1 int32, l2 int32) int32 {
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
	var v18 int32
	_ = v18
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v45 int32
	_ = v45
	var v47 int32
	_ = v47
	var v56 int32
	_ = v56
	var v59 int32
	_ = v59
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v74 int32
	_ = v74
	var v75 int32
	_ = v75
	var v82 int32
	_ = v82
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v87 int32
	_ = v87
	var v94 int32
	_ = v94
	var v95 int32
	_ = v95
	var v98 int32
	_ = v98
	var v99 int32
	_ = v99
	var v100 int32
	_ = v100
	var v101 int32
	_ = v101
	var v103 int32
	_ = v103
	var v112 int32
	_ = v112
	var v113 int32
	_ = v113
	var v120 int32
	_ = v120
	var v123 int32
	_ = v123
	var v127 int32
	_ = v127
	var v130 int32
	_ = v130
	var v136 int32
	_ = v136
	var v137 int32
	_ = v137
	var v142 int32
	_ = v142
	var v152 int32
	_ = v152
	var v154 int64
	_ = v154
	var v160 int32
	_ = v160
	var v169 int32
	_ = v169
	var v170 int32
	_ = v170
	var v177 int32
	_ = v177
	var v180 int32
	_ = v180
	var v184 int32
	_ = v184
	var v192 int32
	_ = v192
	var v194 int32
	_ = v194
	var v195 int32
	_ = v195
	var v196 int32
	_ = v196
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v202 int32
	_ = v202
	v10 = m.G0
	v12 = v10 - int32(80)
	m.G0 = v12
	v14 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+4)))
	v15 = base.I32_extend16_s(v14)
	if base.Ui32(v14) <= base.Ui32(int32(_a_F_numeric_sub_opt_error_0)) {
		v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
		v19 = base.I32_extend16_s(v18)
		if base.Ui32(int32(_a_F_numeric_sub_opt_error_0)) < base.Ui32(v18) {
			v47 = v19
			if v47&int32(_a_F_numeric_sub_opt_error_1) != int32(_a_F_numeric_sub_opt_error_2) {
				if v14 != int32(_a_F_numeric_sub_opt_error_3) {
					if v14 != int32(_a_F_numeric_sub_opt_error_4) {
						if v47&int32(_a_F_numeric_sub_opt_error_1) == int32(_a_F_numeric_sub_opt_error_4) {
							v94 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_5), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v202 = v94
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v98 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_6), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v202 = v98
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					} else {
						if v47&int32(_a_F_numeric_sub_opt_error_1) == int32(_a_F_numeric_sub_opt_error_4) {
							v70 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_7), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v202 = v70
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v74 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_6), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v202 = v74
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					}
				} else {
					if v47&int32(_a_F_numeric_sub_opt_error_1) == int32(_a_F_numeric_sub_opt_error_3) {
						v82 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_7), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v202 = v82
							m.G0 = v12 + int32(80)
							return v202
						}
					} else {
						v86 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_5), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v202 = v86
							m.G0 = v12 + int32(80)
							return v202
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_7), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v202 = v56
					m.G0 = v12 + int32(80)
					return v202
				}
			}
		} else {
			v22 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
			v28 = base.B2i32(int32(0) <= v15)
			if int32(0) <= v15 {
				v29 = int32(-8)
			} else {
				v29 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+56)) = int32(base.Ui32(int32(base.Ui32(v22)>>(uint(int32(2))%32))+v29) >> (uint(int32(1)) % 32))
			if int32(0) <= v15 {
				v100 = int32(*(*int16)(unsafe.Add(mBase, uint32(l0)+6)))
				v101 = v100
			} else {
				v101 = v14<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v14&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+60)) = v101
			v103 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+72)) = v103
			v112 = base.B2i32(v15 < v103)
			if v15 < v103 {
				v113 = int32(base.Ui32(v14)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v113 = v14 & int32(_a_F_numeric_sub_opt_error_8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+68)) = v113
			v120 = v14 & int32(_a_F_numeric_sub_opt_error_2)
			if v120 == int32(_a_F_numeric_sub_opt_error_9) {
				v123 = v14 << (uint(int32(1)) % 32) & int32(_a_F_numeric_sub_opt_error_10)
			} else {
				v123 = v120
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+64)) = v123
			if v15 < v103 {
				v127 = int32(6)
			} else {
				v127 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+76)) = l0 + v127
			v130 = *(*int32)(unsafe.Add(mBase, uint32(l1)))
			v136 = base.B2i32(int32(0) <= v19)
			if int32(0) <= v19 {
				v137 = int32(-8)
			} else {
				v137 = int32(-6)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = int32(base.Ui32(int32(base.Ui32(v130)>>(uint(int32(2))%32))+v137) >> (uint(int32(1)) % 32))
			if int32(0) <= v19 {
				v142 = int32(*(*int16)(unsafe.Add(mBase, uint32(l1)+6)))
				v152 = v142
			} else {
				v152 = v18<<(uint(int32(25))%32)>>(uint(int32(31))%32)&int32(-64) | v18&int32(63)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+36)) = v152
			v154 = int64(0)
			*(*int64)(unsafe.Add(mBase, uint32(v12)+8)) = v154
			*(*int64)(unsafe.Add(mBase, uint32(v12)+16)) = v154
			*(*int64)(unsafe.Add(mBase, uint32(v12)+24)) = v154
			v160 = int32(0)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+48)) = v160
			v169 = base.B2i32(v19 < v160)
			if v19 < v160 {
				v170 = int32(base.Ui32(v18)>>(uint(int32(7))%32)) & int32(63)
			} else {
				v170 = v18 & int32(_a_F_numeric_sub_opt_error_8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v170
			v177 = v18 & int32(_a_F_numeric_sub_opt_error_2)
			if v177 == int32(_a_F_numeric_sub_opt_error_9) {
				v180 = v18 << (uint(int32(1)) % 32) & int32(_a_F_numeric_sub_opt_error_10)
			} else {
				v180 = v177
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+40)) = v180
			if v19 < v160 {
				v184 = int32(6)
			} else {
				v184 = int32(8)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+52)) = l1 + v184
			v192 = v12 + int32(8)
			F_sub_var(m, v12+int32(56), v12+int32(32), v192)
			mBase = m.M
			v194 = m.ExcPending
			if v194 != 0 {
				return int32(0)
			} else {
				v195 = F_make_result_opt_error(m, v192, l2)
				mBase = m.M
				v196 = m.ExcPending
				if v196 != 0 {
					return int32(0)
				} else {
					v197 = *(*int32)(unsafe.Add(mBase, uint32(v12)+24))
					if v197 == int32(0) {
						v202 = v195
						m.G0 = v12 + int32(80)
						return v202
					} else {
						F_pfree(m, v197)
						mBase = m.M
						v201 = m.ExcPending
						if v201 != 0 {
							return int32(0)
						} else {
							v202 = v195
							m.G0 = v12 + int32(80)
							return v202
						}
					}
				}
			}
		}
	} else {
		if v15 == int32(-16384) {
			v56 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_7), int32(0))
			mBase = m.M
			v59 = m.ExcPending
			if v59 != 0 {
				return int32(0)
			} else {
				v202 = v56
				m.G0 = v12 + int32(80)
				return v202
			}
		} else {
			v45 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l1)+4)))
			v47 = v45
			if v47&int32(_a_F_numeric_sub_opt_error_1) != int32(_a_F_numeric_sub_opt_error_2) {
				if v14 != int32(_a_F_numeric_sub_opt_error_3) {
					if v14 != int32(_a_F_numeric_sub_opt_error_4) {
						if v47&int32(_a_F_numeric_sub_opt_error_1) == int32(_a_F_numeric_sub_opt_error_4) {
							v94 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_5), int32(0))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
								return int32(0)
							} else {
								v202 = v94
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v98 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_6), int32(0))
							mBase = m.M
							v99 = m.ExcPending
							if v99 != 0 {
								return int32(0)
							} else {
								v202 = v98
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					} else {
						if v47&int32(_a_F_numeric_sub_opt_error_1) == int32(_a_F_numeric_sub_opt_error_4) {
							v70 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_7), int32(0))
							mBase = m.M
							v71 = m.ExcPending
							if v71 != 0 {
								return int32(0)
							} else {
								v202 = v70
								m.G0 = v12 + int32(80)
								return v202
							}
						} else {
							v74 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_6), int32(0))
							mBase = m.M
							v75 = m.ExcPending
							if v75 != 0 {
								return int32(0)
							} else {
								v202 = v74
								m.G0 = v12 + int32(80)
								return v202
							}
						}
					}
				} else {
					if v47&int32(_a_F_numeric_sub_opt_error_1) == int32(_a_F_numeric_sub_opt_error_3) {
						v82 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_7), int32(0))
						mBase = m.M
						v83 = m.ExcPending
						if v83 != 0 {
							return int32(0)
						} else {
							v202 = v82
							m.G0 = v12 + int32(80)
							return v202
						}
					} else {
						v86 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_5), int32(0))
						mBase = m.M
						v87 = m.ExcPending
						if v87 != 0 {
							return int32(0)
						} else {
							v202 = v86
							m.G0 = v12 + int32(80)
							return v202
						}
					}
				}
			} else {
				v56 = F_make_result_opt_error(m, int32(_a_F_numeric_sub_opt_error_7), int32(0))
				mBase = m.M
				v59 = m.ExcPending
				if v59 != 0 {
					return int32(0)
				} else {
					v202 = v56
					m.G0 = v12 + int32(80)
					return v202
				}
			}
		}
	}
}
func F_numeric_support(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v2 int32
	_ = v2
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
	var v48 int32
	_ = v48
	v2 = int32(0)
	v5 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
	v6 = *(*int32)(unsafe.Add(mBase, uint32(v5)))
	if v6 != int32(457) {
		v48 = v2
		return v48
	} else {
		v9 = *(*int32)(unsafe.Add(mBase, uint32(v5)+8))
		v10 = *(*int32)(unsafe.Add(mBase, uint32(v9)+28))
		v11 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v12 = *(*int32)(unsafe.Add(mBase, uint32(v11)+4))
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v12)))
		if v13 != int32(7) {
			v48 = v2
			return v48
		} else {
			v16 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v12)+24)))
			if v16 != 0 {
				v48 = v2
				return v48
			} else {
				v17 = *(*int32)(unsafe.Add(mBase, uint32(v11)))
				v18 = F_exprTypmod(m, v17)
				mBase = m.M
				v21 = m.ExcPending
				if v21 != 0 {
					return int32(0)
				} else {
					v22 = *(*int32)(unsafe.Add(mBase, uint32(v12)+20))
					v23 = int32(4)
					v34 = int32(16)
					if (v18^v22)&int32(2047)|base.B2i32(v18 < v23)|base.B2i32(base.Ui32(int32(base.Ui32(v22-v23)>>(uint(v34)%32))) < base.Ui32(int32(base.Ui32(v18-v23)>>(uint(v34)%32)))) != 0 {
						v42 = base.B2i32(v23 <= v22)
					} else {
						v42 = int32(0)
					}
					if v42 != 0 {
						v48 = v2
						return v48
					} else {
						v43 = F_relabel_to_typmod(m, v17, v22)
						mBase = m.M
						v44 = m.ExcPending
						if v44 != 0 {
							return int32(0)
						} else {
							v48 = v43
							return v48
						}
					}
				}
			}
		}
	}
}
func F_numeric_uminus(m *base.Module, l0 int32) int32 {
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
	var v44 int32
	_ = v44
	var v50 int32
	_ = v50
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
			v18 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v5)+4)))
			if base.Ui32(int32(_a_F_numeric_uminus_0)) <= base.Ui32(v18) {
				if v18 == int32(_a_F_numeric_uminus_0) {
					return v12
				} else {
					v50 = v18 ^ int32(_a_F_numeric_uminus_1)
					*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v50)
					return v12
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
					return v12
				} else {
					if v28 <= int32(-16385) {
						v50 = v18 ^ int32(_a_F_numeric_uminus_1)
						*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v50)
						return v12
					} else {
						if base.Ui32(v18) <= base.Ui32(int32(_a_F_numeric_uminus_2)) {
							v40 = v18 | int32(_a_F_numeric_uminus_3)
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v40)
							return v12
						} else {
							v44 = v18 & int32(_a_F_numeric_uminus_2)
							*(*uint16)(unsafe.Add(mBase, uint32(v12)+4)) = uint16(v44)
							return v12
						}
					}
				}
			}
		}
	}
}
