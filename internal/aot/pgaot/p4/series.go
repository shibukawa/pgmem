package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_series_step_int8(m *base.Module, l0 int32) int64 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int64
	_ = v12
	var v13 int64
	_ = v13
	var v15 int32
	_ = v15
	var v18 int64
	_ = v18
	var v21 int64
	_ = v21
	var v22 int32
	_ = v22
	var v25 int32
	_ = v25
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v29 int32
	_ = v29
	var v32 int32
	_ = v32
	var v33 int32
	_ = v33
	var v46 int32
	_ = v46
	var v47 int32
	_ = v47
	var v48 int32
	_ = v48
	var v49 int64
	_ = v49
	var v50 int64
	_ = v50
	var v53 int64
	_ = v53
	var v57 int64
	_ = v57
	var v59 int64
	_ = v59
	var v67 int64
	_ = v67
	var v71 int32
	_ = v71
	var v78 int32
	_ = v78
	var v81 int32
	_ = v81
	var v85 int32
	_ = v85
	var v90 int32
	_ = v90
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v96 int32
	_ = v96
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v15 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v15 == int32(3) {
			v18 = *(*int64)(unsafe.Add(mBase, uint32(l0)+56))
			if v18 == int64(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int64(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v81 = m.ExcPending
					if v81 != 0 {
						return int64(0)
					} else {
						F_errmsg(m, int32(_a_F_generate_series_step_int8_0), int32(0))
						mBase = m.M
						v85 = m.ExcPending
						if v85 != 0 {
							return int64(0)
						} else {
							F_errfinish(m, int32(_a_F_generate_series_step_int8_1), int32(1453), int32(_a_F_generate_series_step_int8_2))
							mBase = m.M
							v90 = m.ExcPending
							if v90 != 0 {
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
				v21 = v18
				v22 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v25 = m.ExcPending
				if v25 != 0 {
					return int64(0)
				} else {
					v26 = int32(_a_F_generate_series_step_int8_3)
					v27 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0]))
					v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0])) = v29
					v32 = F_palloc(m, int32(24))
					mBase = m.M
					v33 = m.ExcPending
					if v33 != 0 {
						return int64(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v21
						*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v12
						*(*int64)(unsafe.Add(mBase, uint32(v32))) = v13
						*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v32
						*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0])) = v27
						v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
						v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
						v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
						v50 = *(*int64)(unsafe.Add(mBase, uint32(v48)+16))
						if int64(0) < v50 {
							v53 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
							if v49 <= v53 {
								v59 = v49 + v50
								*(*int64)(unsafe.Add(mBase, uint32(v48))) = v59
								if base.B2i32(v50 < int64(0)) != base.B2i32(v59 < v49) {
									*(*int64)(unsafe.Add(mBase, uint32(v48)+16)) = int64(0)
								} else {
								}
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
								*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
								return v49
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
									v96 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
									return int64(0)
								}
							}
						} else {
							if int64(0) <= v50 {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
									v96 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
									return int64(0)
								}
							} else {
								v57 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
								if v49 < v57 {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v92 = m.ExcPending
									if v92 != 0 {
										return int64(0)
									} else {
										v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
										v96 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
										return int64(0)
									}
								} else {
									v59 = v49 + v50
									*(*int64)(unsafe.Add(mBase, uint32(v48))) = v59
									if base.B2i32(v50 < int64(0)) != base.B2i32(v59 < v49) {
										*(*int64)(unsafe.Add(mBase, uint32(v48)+16)) = int64(0)
									} else {
									}
									v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
									*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
									v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
									return v49
								}
							}
						}
					}
				}
			}
		} else {
			v21 = int64(1)
			v22 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v25 = m.ExcPending
			if v25 != 0 {
				return int64(0)
			} else {
				v26 = int32(_a_F_generate_series_step_int8_3)
				v27 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0]))
				v29 = *(*int32)(unsafe.Add(mBase, uint32(v22)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0])) = v29
				v32 = F_palloc(m, int32(24))
				mBase = m.M
				v33 = m.ExcPending
				if v33 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v32)+16)) = v21
					*(*int64)(unsafe.Add(mBase, uint32(v32)+8)) = v12
					*(*int64)(unsafe.Add(mBase, uint32(v32))) = v13
					*(*int32)(unsafe.Add(mBase, uint32(v22)+16)) = v32
					*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0])) = v27
					v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
					v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
					v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
					v50 = *(*int64)(unsafe.Add(mBase, uint32(v48)+16))
					if int64(0) < v50 {
						v53 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
						if v49 <= v53 {
							v59 = v49 + v50
							*(*int64)(unsafe.Add(mBase, uint32(v48))) = v59
							if base.B2i32(v50 < int64(0)) != base.B2i32(v59 < v49) {
								*(*int64)(unsafe.Add(mBase, uint32(v48)+16)) = int64(0)
							} else {
							}
							v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
							*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
							v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
							return v49
						} else {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int64(0)
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
								v96 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
								return int64(0)
							}
						}
					} else {
						if int64(0) <= v50 {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v92 = m.ExcPending
							if v92 != 0 {
								return int64(0)
							} else {
								v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
								v96 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
								return int64(0)
							}
						} else {
							v57 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
							if v49 < v57 {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v92 = m.ExcPending
								if v92 != 0 {
									return int64(0)
								} else {
									v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
									v96 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
									return int64(0)
								}
							} else {
								v59 = v49 + v50
								*(*int64)(unsafe.Add(mBase, uint32(v48))) = v59
								if base.B2i32(v50 < int64(0)) != base.B2i32(v59 < v49) {
									*(*int64)(unsafe.Add(mBase, uint32(v48)+16)) = int64(0)
								} else {
								}
								v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
								*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
								v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
								return v49
							}
						}
					}
				}
			}
		}
	} else {
		v46 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v47 = *(*int32)(unsafe.Add(mBase, uint32(v46)+16))
		v48 = *(*int32)(unsafe.Add(mBase, uint32(v47)+16))
		v49 = *(*int64)(unsafe.Add(mBase, uint32(v48)))
		v50 = *(*int64)(unsafe.Add(mBase, uint32(v48)+16))
		if int64(0) < v50 {
			v53 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
			if v49 <= v53 {
				v59 = v49 + v50
				*(*int64)(unsafe.Add(mBase, uint32(v48))) = v59
				if base.B2i32(v50 < int64(0)) != base.B2i32(v59 < v49) {
					*(*int64)(unsafe.Add(mBase, uint32(v48)+16)) = int64(0)
				} else {
				}
				v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
				*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
				v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
				return v49
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int64(0)
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
					v96 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
					return int64(0)
				}
			}
		} else {
			if int64(0) <= v50 {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v92 = m.ExcPending
				if v92 != 0 {
					return int64(0)
				} else {
					v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
					v96 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
					return int64(0)
				}
			} else {
				v57 = *(*int64)(unsafe.Add(mBase, uint32(v48)+8))
				if v49 < v57 {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v92 = m.ExcPending
					if v92 != 0 {
						return int64(0)
					} else {
						v93 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v93)+20)) = int32(2)
						v96 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v96)
						return int64(0)
					}
				} else {
					v59 = v49 + v50
					*(*int64)(unsafe.Add(mBase, uint32(v48))) = v59
					if base.B2i32(v50 < int64(0)) != base.B2i32(v59 < v49) {
						*(*int64)(unsafe.Add(mBase, uint32(v48)+16)) = int64(0)
					} else {
					}
					v67 = *(*int64)(unsafe.Add(mBase, uint32(v47)))
					*(*int64)(unsafe.Add(mBase, uint32(v47))) = v67 + int64(1)
					v71 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v71)+20)) = int32(1)
					return v49
				}
			}
		}
	}
}
func F_generate_series_timestamp(m *base.Module, l0 int32) int64 {
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
	var v21 int64
	_ = v21
	var v22 int64
	_ = v22
	var v23 int32
	_ = v23
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v28 int32
	_ = v28
	var v30 int32
	_ = v30
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v37 int64
	_ = v37
	var v39 int64
	_ = v39
	var v42 int32
	_ = v42
	var v46 int64
	_ = v46
	var v55 int64
	_ = v55
	var v56 int64
	_ = v56
	var v58 int64
	_ = v58
	var v61 int64
	_ = v61
	var v62 int64
	_ = v62
	var v64 int64
	_ = v64
	var v65 int64
	_ = v65
	var v69 int64
	_ = v69
	var v76 int64
	_ = v76
	var v87 int64
	_ = v87
	var v88 int64
	_ = v88
	var v89 int64
	_ = v89
	var v93 int64
	_ = v93
	var v97 int64
	_ = v97
	var v102 int32
	_ = v102
	var v103 int64
	_ = v103
	var v110 int32
	_ = v110
	var v113 int32
	_ = v113
	var v137 int32
	_ = v137
	var v138 int32
	_ = v138
	var v139 int32
	_ = v139
	var v140 int64
	_ = v140
	var v141 int64
	_ = v141
	var v142 int32
	_ = v142
	var v152 int64
	_ = v152
	var v153 int32
	_ = v153
	var v155 int64
	_ = v155
	var v159 int32
	_ = v159
	var v165 int32
	_ = v165
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v177 int32
	_ = v177
	var v179 int32
	_ = v179
	var v180 int32
	_ = v180
	var v183 int32
	_ = v183
	var v186 int64
	_ = v186
	var v194 int32
	_ = v194
	var v197 int32
	_ = v197
	var v201 int32
	_ = v201
	var v206 int32
	_ = v206
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 == int32(0) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+56))
		v21 = *(*int64)(unsafe.Add(mBase, uint32(l0)+24))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(l0)+40))
		v23 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v26 = m.ExcPending
		if v26 != 0 {
			return int64(0)
		} else {
			v27 = int32(_a_F_generate_series_timestamp_0)
			v28 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0]))
			v30 = *(*int32)(unsafe.Add(mBase, uint32(v23)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0])) = v30
			v33 = F_palloc(m, int32(40))
			mBase = m.M
			v34 = m.ExcPending
			if v34 != 0 {
				return int64(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v33)+8)) = v22
				*(*int64)(unsafe.Add(mBase, uint32(v33))) = v21
				v37 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+24)) = v37
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
				*(*int64)(unsafe.Add(mBase, uint32(v33)+16)) = v39
				v42 = *(*int32)(unsafe.Add(mBase, uint32(v33)+28))
				v46 = base.I64_extend32_s(v37) + base.I64_extend_i32_s(v42)*int64(30)
				v55 = int64(32)
				v56 = int64(20)
				v58 = int64(base.Ui64(v46) >> (uint(v55) % 64))
				v61 = int64(4294967295)
				v62 = int64(500654080)
				v64 = v46 & v61
				v65 = v62 * v64
				v69 = int64(base.Ui64(v65)>>(uint(v55)%64)) + v62*v58
				v76 = v64*v56 + v69&v61
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v46*int64(0) + v46>>(uint(int64(63))%64)*int64(86400000000) + v56*v58 + int64(base.Ui64(v69)>>(uint(v55)%64)) + int64(base.Ui64(v76)>>(uint(v55)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = v65&v61 | v76<<(uint(v55)%64)
				v87 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
				v88 = v39 + v87
				v89 = int64(0)
				v93 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
				v97 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v88) < base.Ui64(v87))) + (v93 + v39>>(uint(int64(63))%64))
				if v97 == v89 {
					v102 = base.B2i32(v88 != v89)
				} else {
					v102 = base.B2i32(v89 < v97)
				}
				v103 = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v33)+32)) = v102 - base.B2i32(v97 < v103)
				if v97|v88 == v103 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v165 = m.ExcPending
					if v165 != 0 {
						return int64(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v168 = m.ExcPending
						if v168 != 0 {
							return int64(0)
						} else {
							F_errmsg(m, int32(_a_F_generate_series_timestamp_1), int32(0))
							mBase = m.M
							v172 = m.ExcPending
							if v172 != 0 {
								return int64(0)
							} else {
								F_errfinish(m, int32(_a_F_generate_series_timestamp_2), int32(_a_F_generate_series_timestamp_3), int32(_a_F_generate_series_timestamp_4))
								mBase = m.M
								v177 = m.ExcPending
								if v177 != 0 {
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
					v110 = base.I32_wrap_i64(v37)
					if v42 != int32(2147483647) {
						v113 = int32(-2147483648)
						if base.B2i32(v42 != v113)|base.B2i32(v39 != int64(-9223372036854775807-1))|base.B2i32(v110 != v113) != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v33
							*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0])) = v28
							v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
							v140 = *(*int64)(unsafe.Add(mBase, uint32(v139)+8))
							v141 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
							v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+32))
							if int32(0) < v142 {
								if v141 <= v140 {
									v152 = F_DirectFunctionCall2Coll(m, int32(1395), int32(0), v141, base.I64_extend_i32_u(v139+int32(16)))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v139))) = v152
										v155 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
										*(*int64)(unsafe.Add(mBase, uint32(v138))) = v155 + int64(1)
										v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = int32(1)
										v186 = v141
										m.G0 = v14 + int32(16)
										return v186
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return int64(0)
									} else {
										v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = int32(2)
										v183 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v183)
										v186 = int64(0)
										m.G0 = v14 + int32(16)
										return v186
									}
								}
							} else {
								if v141 < v140 {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return int64(0)
									} else {
										v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = int32(2)
										v183 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v183)
										v186 = int64(0)
										m.G0 = v14 + int32(16)
										return v186
									}
								} else {
									v152 = F_DirectFunctionCall2Coll(m, int32(1395), int32(0), v141, base.I64_extend_i32_u(v139+int32(16)))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v139))) = v152
										v155 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
										*(*int64)(unsafe.Add(mBase, uint32(v138))) = v155 + int64(1)
										v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = int32(1)
										v186 = v141
										m.G0 = v14 + int32(16)
										return v186
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v194 = m.ExcPending
							if v194 != 0 {
								return int64(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v197 = m.ExcPending
								if v197 != 0 {
									return int64(0)
								} else {
									F_errmsg(m, int32(_a_F_generate_series_timestamp_5), int32(0))
									mBase = m.M
									v201 = m.ExcPending
									if v201 != 0 {
										return int64(0)
									} else {
										F_errfinish(m, int32(_a_F_generate_series_timestamp_2), int32(_a_F_generate_series_timestamp_6), int32(_a_F_generate_series_timestamp_4))
										mBase = m.M
										v206 = m.ExcPending
										if v206 != 0 {
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
						if v39 != int64(9223372036854775807) {
							*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v33
							*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0])) = v28
							v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
							v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
							v140 = *(*int64)(unsafe.Add(mBase, uint32(v139)+8))
							v141 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
							v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+32))
							if int32(0) < v142 {
								if v141 <= v140 {
									v152 = F_DirectFunctionCall2Coll(m, int32(1395), int32(0), v141, base.I64_extend_i32_u(v139+int32(16)))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v139))) = v152
										v155 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
										*(*int64)(unsafe.Add(mBase, uint32(v138))) = v155 + int64(1)
										v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = int32(1)
										v186 = v141
										m.G0 = v14 + int32(16)
										return v186
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return int64(0)
									} else {
										v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = int32(2)
										v183 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v183)
										v186 = int64(0)
										m.G0 = v14 + int32(16)
										return v186
									}
								}
							} else {
								if v141 < v140 {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v179 = m.ExcPending
									if v179 != 0 {
										return int64(0)
									} else {
										v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = int32(2)
										v183 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v183)
										v186 = int64(0)
										m.G0 = v14 + int32(16)
										return v186
									}
								} else {
									v152 = F_DirectFunctionCall2Coll(m, int32(1395), int32(0), v141, base.I64_extend_i32_u(v139+int32(16)))
									mBase = m.M
									v153 = m.ExcPending
									if v153 != 0 {
										return int64(0)
									} else {
										*(*int64)(unsafe.Add(mBase, uint32(v139))) = v152
										v155 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
										*(*int64)(unsafe.Add(mBase, uint32(v138))) = v155 + int64(1)
										v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = int32(1)
										v186 = v141
										m.G0 = v14 + int32(16)
										return v186
									}
								}
							}
						} else {
							if v110 == int32(2147483647) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v194 = m.ExcPending
								if v194 != 0 {
									return int64(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v197 = m.ExcPending
									if v197 != 0 {
										return int64(0)
									} else {
										F_errmsg(m, int32(_a_F_generate_series_timestamp_5), int32(0))
										mBase = m.M
										v201 = m.ExcPending
										if v201 != 0 {
											return int64(0)
										} else {
											F_errfinish(m, int32(_a_F_generate_series_timestamp_2), int32(_a_F_generate_series_timestamp_6), int32(_a_F_generate_series_timestamp_4))
											mBase = m.M
											v206 = m.ExcPending
											if v206 != 0 {
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
								*(*int32)(unsafe.Add(mBase, uint32(v23)+16)) = v33
								*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0])) = v28
								v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
								v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
								v140 = *(*int64)(unsafe.Add(mBase, uint32(v139)+8))
								v141 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
								v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+32))
								if int32(0) < v142 {
									if v141 <= v140 {
										v152 = F_DirectFunctionCall2Coll(m, int32(1395), int32(0), v141, base.I64_extend_i32_u(v139+int32(16)))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v139))) = v152
											v155 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
											*(*int64)(unsafe.Add(mBase, uint32(v138))) = v155 + int64(1)
											v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = int32(1)
											v186 = v141
											m.G0 = v14 + int32(16)
											return v186
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return int64(0)
										} else {
											v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = int32(2)
											v183 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v183)
											v186 = int64(0)
											m.G0 = v14 + int32(16)
											return v186
										}
									}
								} else {
									if v141 < v140 {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v179 = m.ExcPending
										if v179 != 0 {
											return int64(0)
										} else {
											v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = int32(2)
											v183 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v183)
											v186 = int64(0)
											m.G0 = v14 + int32(16)
											return v186
										}
									} else {
										v152 = F_DirectFunctionCall2Coll(m, int32(1395), int32(0), v141, base.I64_extend_i32_u(v139+int32(16)))
										mBase = m.M
										v153 = m.ExcPending
										if v153 != 0 {
											return int64(0)
										} else {
											*(*int64)(unsafe.Add(mBase, uint32(v139))) = v152
											v155 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
											*(*int64)(unsafe.Add(mBase, uint32(v138))) = v155 + int64(1)
											v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = int32(1)
											v186 = v141
											m.G0 = v14 + int32(16)
											return v186
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
		v137 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v138 = *(*int32)(unsafe.Add(mBase, uint32(v137)+16))
		v139 = *(*int32)(unsafe.Add(mBase, uint32(v138)+16))
		v140 = *(*int64)(unsafe.Add(mBase, uint32(v139)+8))
		v141 = *(*int64)(unsafe.Add(mBase, uint32(v139)))
		v142 = *(*int32)(unsafe.Add(mBase, uint32(v139)+32))
		if int32(0) < v142 {
			if v141 <= v140 {
				v152 = F_DirectFunctionCall2Coll(m, int32(1395), int32(0), v141, base.I64_extend_i32_u(v139+int32(16)))
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v139))) = v152
					v155 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
					*(*int64)(unsafe.Add(mBase, uint32(v138))) = v155 + int64(1)
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = int32(1)
					v186 = v141
					m.G0 = v14 + int32(16)
					return v186
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int64(0)
				} else {
					v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = int32(2)
					v183 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v183)
					v186 = int64(0)
					m.G0 = v14 + int32(16)
					return v186
				}
			}
		} else {
			if v141 < v140 {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v179 = m.ExcPending
				if v179 != 0 {
					return int64(0)
				} else {
					v180 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v180)+20)) = int32(2)
					v183 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v183)
					v186 = int64(0)
					m.G0 = v14 + int32(16)
					return v186
				}
			} else {
				v152 = F_DirectFunctionCall2Coll(m, int32(1395), int32(0), v141, base.I64_extend_i32_u(v139+int32(16)))
				mBase = m.M
				v153 = m.ExcPending
				if v153 != 0 {
					return int64(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v139))) = v152
					v155 = *(*int64)(unsafe.Add(mBase, uint32(v138)))
					*(*int64)(unsafe.Add(mBase, uint32(v138))) = v155 + int64(1)
					v159 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v159)+20)) = int32(1)
					v186 = v141
					m.G0 = v14 + int32(16)
					return v186
				}
			}
		}
	}
}
