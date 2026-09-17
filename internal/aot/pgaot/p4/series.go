package p4

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_generate_series_step_int8(m *base.Module, l0 int32) int32 {
	mBase := m.M
	_ = mBase
	var v8 int32
	_ = v8
	var v9 int32
	_ = v9
	var v12 int32
	_ = v12
	var v13 int64
	_ = v13
	var v14 int32
	_ = v14
	var v15 int64
	_ = v15
	var v17 int32
	_ = v17
	var v20 int32
	_ = v20
	var v21 int64
	_ = v21
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v49 int32
	_ = v49
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v52 int64
	_ = v52
	var v53 int64
	_ = v53
	var v56 int64
	_ = v56
	var v60 int64
	_ = v60
	var v62 int64
	_ = v62
	var v70 int64
	_ = v70
	var v74 int32
	_ = v74
	var v77 int32
	_ = v77
	var v78 int32
	_ = v78
	var v83 int32
	_ = v83
	var v86 int32
	_ = v86
	var v90 int32
	_ = v90
	var v95 int32
	_ = v95
	var v97 int32
	_ = v97
	var v98 int32
	_ = v98
	var v101 int32
	_ = v101
	v8 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v9 = *(*int32)(unsafe.Add(mBase, uint32(v8)+16))
	if v9 == int32(0) {
		v12 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v13 = *(*int64)(unsafe.Add(mBase, uint32(v12)))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v15 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
		v17 = int32(*(*uint16)(unsafe.Add(mBase, uint32(l0)+18)))
		if v17 == int32(3) {
			v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
			v21 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
			if v21 == int64(0) {
				F_errstart_cold(m, int32(21), int32(0))
				mBase = m.M
				v83 = m.ExcPending
				if v83 != 0 {
					return int32(0)
				} else {
					F_errcode(m, int32(50856066))
					mBase = m.M
					v86 = m.ExcPending
					if v86 != 0 {
						return int32(0)
					} else {
						F_errmsg(m, int32(_a_F_generate_series_step_int8_0), int32(0))
						mBase = m.M
						v90 = m.ExcPending
						if v90 != 0 {
							return int32(0)
						} else {
							F_errfinish(m, int32(_a_F_generate_series_step_int8_1), int32(1403), int32(_a_F_generate_series_step_int8_2))
							mBase = m.M
							v95 = m.ExcPending
							if v95 != 0 {
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
				v24 = v21
				v25 = F_init_MultiFuncCall(m, l0)
				mBase = m.M
				v28 = m.ExcPending
				if v28 != 0 {
					return int32(0)
				} else {
					v29 = int32(_a_F_generate_series_step_int8_3)
					v30 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0]))
					v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
					*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0])) = v32
					v35 = F_palloc(m, int32(24))
					mBase = m.M
					v36 = m.ExcPending
					if v36 != 0 {
						return int32(0)
					} else {
						*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v24
						*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v13
						*(*int64)(unsafe.Add(mBase, uint32(v35))) = v15
						*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v35
						*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0])) = v30
						v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
						v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
						v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
						v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
						v53 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
						if int64(0) < v53 {
							v56 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
							if v52 <= v56 {
								v62 = v52 + v53
								*(*int64)(unsafe.Add(mBase, uint32(v51))) = v62
								if base.B2i32(v53 < int64(0)) != base.B2i32(v62 < v52) {
									*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = int64(0)
								} else {
								}
								v70 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
								*(*int64)(unsafe.Add(mBase, uint32(v50))) = v70 + int64(1)
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = int32(1)
								v77 = F_Int64GetDatum(m, v52)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									return v77
								}
							} else {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
									v101 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
									return int32(0)
								}
							}
						} else {
							if int64(0) <= v53 {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
									v101 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
									return int32(0)
								}
							} else {
								v60 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
								if v52 < v60 {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v97 = m.ExcPending
									if v97 != 0 {
										return int32(0)
									} else {
										v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
										v101 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
										return int32(0)
									}
								} else {
									v62 = v52 + v53
									*(*int64)(unsafe.Add(mBase, uint32(v51))) = v62
									if base.B2i32(v53 < int64(0)) != base.B2i32(v62 < v52) {
										*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = int64(0)
									} else {
									}
									v70 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
									*(*int64)(unsafe.Add(mBase, uint32(v50))) = v70 + int64(1)
									v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = int32(1)
									v77 = F_Int64GetDatum(m, v52)
									mBase = m.M
									v78 = m.ExcPending
									if v78 != 0 {
										return int32(0)
									} else {
										return v77
									}
								}
							}
						}
					}
				}
			}
		} else {
			v24 = int64(1)
			v25 = F_init_MultiFuncCall(m, l0)
			mBase = m.M
			v28 = m.ExcPending
			if v28 != 0 {
				return int32(0)
			} else {
				v29 = int32(_a_F_generate_series_step_int8_3)
				v30 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0]))
				v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
				*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0])) = v32
				v35 = F_palloc(m, int32(24))
				mBase = m.M
				v36 = m.ExcPending
				if v36 != 0 {
					return int32(0)
				} else {
					*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v24
					*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v13
					*(*int64)(unsafe.Add(mBase, uint32(v35))) = v15
					*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v35
					*(*int32)(unsafe.Add(mBase, _c_F_generate_series_step_int8[0])) = v30
					v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
					v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
					v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
					v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
					v53 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
					if int64(0) < v53 {
						v56 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
						if v52 <= v56 {
							v62 = v52 + v53
							*(*int64)(unsafe.Add(mBase, uint32(v51))) = v62
							if base.B2i32(v53 < int64(0)) != base.B2i32(v62 < v52) {
								*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = int64(0)
							} else {
							}
							v70 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
							*(*int64)(unsafe.Add(mBase, uint32(v50))) = v70 + int64(1)
							v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
							*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = int32(1)
							v77 = F_Int64GetDatum(m, v52)
							mBase = m.M
							v78 = m.ExcPending
							if v78 != 0 {
								return int32(0)
							} else {
								return v77
							}
						} else {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
								v101 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
								return int32(0)
							}
						}
					} else {
						if int64(0) <= v53 {
							F_end_MultiFuncCall(m, l0)
							mBase = m.M
							v97 = m.ExcPending
							if v97 != 0 {
								return int32(0)
							} else {
								v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
								v101 = int32(1)
								*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
								return int32(0)
							}
						} else {
							v60 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
							if v52 < v60 {
								F_end_MultiFuncCall(m, l0)
								mBase = m.M
								v97 = m.ExcPending
								if v97 != 0 {
									return int32(0)
								} else {
									v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
									*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
									v101 = int32(1)
									*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
									return int32(0)
								}
							} else {
								v62 = v52 + v53
								*(*int64)(unsafe.Add(mBase, uint32(v51))) = v62
								if base.B2i32(v53 < int64(0)) != base.B2i32(v62 < v52) {
									*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = int64(0)
								} else {
								}
								v70 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
								*(*int64)(unsafe.Add(mBase, uint32(v50))) = v70 + int64(1)
								v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
								*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = int32(1)
								v77 = F_Int64GetDatum(m, v52)
								mBase = m.M
								v78 = m.ExcPending
								if v78 != 0 {
									return int32(0)
								} else {
									return v77
								}
							}
						}
					}
				}
			}
		}
	} else {
		v49 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v50 = *(*int32)(unsafe.Add(mBase, uint32(v49)+16))
		v51 = *(*int32)(unsafe.Add(mBase, uint32(v50)+16))
		v52 = *(*int64)(unsafe.Add(mBase, uint32(v51)))
		v53 = *(*int64)(unsafe.Add(mBase, uint32(v51)+16))
		if int64(0) < v53 {
			v56 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
			if v52 <= v56 {
				v62 = v52 + v53
				*(*int64)(unsafe.Add(mBase, uint32(v51))) = v62
				if base.B2i32(v53 < int64(0)) != base.B2i32(v62 < v52) {
					*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = int64(0)
				} else {
				}
				v70 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
				*(*int64)(unsafe.Add(mBase, uint32(v50))) = v70 + int64(1)
				v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
				*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = int32(1)
				v77 = F_Int64GetDatum(m, v52)
				mBase = m.M
				v78 = m.ExcPending
				if v78 != 0 {
					return int32(0)
				} else {
					return v77
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
					v101 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
					return int32(0)
				}
			}
		} else {
			if int64(0) <= v53 {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v97 = m.ExcPending
				if v97 != 0 {
					return int32(0)
				} else {
					v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
					v101 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
					return int32(0)
				}
			} else {
				v60 = *(*int64)(unsafe.Add(mBase, uint32(v51)+8))
				if v52 < v60 {
					F_end_MultiFuncCall(m, l0)
					mBase = m.M
					v97 = m.ExcPending
					if v97 != 0 {
						return int32(0)
					} else {
						v98 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v98)+20)) = int32(2)
						v101 = int32(1)
						*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v101)
						return int32(0)
					}
				} else {
					v62 = v52 + v53
					*(*int64)(unsafe.Add(mBase, uint32(v51))) = v62
					if base.B2i32(v53 < int64(0)) != base.B2i32(v62 < v52) {
						*(*int64)(unsafe.Add(mBase, uint32(v51)+16)) = int64(0)
					} else {
					}
					v70 = *(*int64)(unsafe.Add(mBase, uint32(v50)))
					*(*int64)(unsafe.Add(mBase, uint32(v50))) = v70 + int64(1)
					v74 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v74)+20)) = int32(1)
					v77 = F_Int64GetDatum(m, v52)
					mBase = m.M
					v78 = m.ExcPending
					if v78 != 0 {
						return int32(0)
					} else {
						return v77
					}
				}
			}
		}
	}
}
func F_generate_series_timestamp(m *base.Module, l0 int32) int32 {
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
	var v24 int64
	_ = v24
	var v25 int32
	_ = v25
	var v28 int32
	_ = v28
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v32 int32
	_ = v32
	var v35 int32
	_ = v35
	var v36 int32
	_ = v36
	var v39 int64
	_ = v39
	var v41 int64
	_ = v41
	var v44 int32
	_ = v44
	var v48 int64
	_ = v48
	var v57 int64
	_ = v57
	var v58 int64
	_ = v58
	var v60 int64
	_ = v60
	var v63 int64
	_ = v63
	var v64 int64
	_ = v64
	var v66 int64
	_ = v66
	var v67 int64
	_ = v67
	var v71 int64
	_ = v71
	var v78 int64
	_ = v78
	var v89 int64
	_ = v89
	var v90 int64
	_ = v90
	var v91 int64
	_ = v91
	var v95 int64
	_ = v95
	var v99 int64
	_ = v99
	var v104 int32
	_ = v104
	var v105 int64
	_ = v105
	var v112 int32
	_ = v112
	var v115 int32
	_ = v115
	var v139 int32
	_ = v139
	var v140 int32
	_ = v140
	var v141 int32
	_ = v141
	var v142 int64
	_ = v142
	var v143 int64
	_ = v143
	var v144 int32
	_ = v144
	var v151 int32
	_ = v151
	var v152 int32
	_ = v152
	var v155 int32
	_ = v155
	var v156 int32
	_ = v156
	var v157 int64
	_ = v157
	var v159 int64
	_ = v159
	var v163 int32
	_ = v163
	var v166 int32
	_ = v166
	var v167 int32
	_ = v167
	var v171 int32
	_ = v171
	var v174 int32
	_ = v174
	var v178 int32
	_ = v178
	var v183 int32
	_ = v183
	var v185 int32
	_ = v185
	var v186 int32
	_ = v186
	var v189 int32
	_ = v189
	var v192 int32
	_ = v192
	var v200 int32
	_ = v200
	var v203 int32
	_ = v203
	var v207 int32
	_ = v207
	var v212 int32
	_ = v212
	v12 = m.G0
	v14 = v12 - int32(16)
	m.G0 = v14
	v16 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
	v17 = *(*int32)(unsafe.Add(mBase, uint32(v16)+16))
	if v17 == int32(0) {
		v20 = *(*int32)(unsafe.Add(mBase, uint32(l0)+36))
		v21 = *(*int32)(unsafe.Add(mBase, uint32(l0)+20))
		v22 = *(*int64)(unsafe.Add(mBase, uint32(v21)))
		v23 = *(*int32)(unsafe.Add(mBase, uint32(l0)+28))
		v24 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
		v25 = F_init_MultiFuncCall(m, l0)
		mBase = m.M
		v28 = m.ExcPending
		if v28 != 0 {
			return int32(0)
		} else {
			v29 = int32(_a_F_generate_series_timestamp_0)
			v30 = *(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0]))
			v32 = *(*int32)(unsafe.Add(mBase, uint32(v25)+24))
			*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0])) = v32
			v35 = F_palloc(m, int32(40))
			mBase = m.M
			v36 = m.ExcPending
			if v36 != 0 {
				return int32(0)
			} else {
				*(*int64)(unsafe.Add(mBase, uint32(v35)+8)) = v24
				*(*int64)(unsafe.Add(mBase, uint32(v35))) = v22
				v39 = *(*int64)(unsafe.Add(mBase, uint32(v20)+8))
				*(*int64)(unsafe.Add(mBase, uint32(v35)+24)) = v39
				v41 = *(*int64)(unsafe.Add(mBase, uint32(v20)))
				*(*int64)(unsafe.Add(mBase, uint32(v35)+16)) = v41
				v44 = *(*int32)(unsafe.Add(mBase, uint32(v35)+28))
				v48 = base.I64_extend32_s(v39) + base.I64_extend_i32_s(v44)*int64(30)
				v57 = int64(32)
				v58 = int64(20)
				v60 = int64(base.Ui64(v48) >> (uint(v57) % 64))
				v63 = int64(4294967295)
				v64 = int64(500654080)
				v66 = v48 & v63
				v67 = v64 * v66
				v71 = int64(base.Ui64(v67)>>(uint(v57)%64)) + v64*v60
				v78 = v66*v58 + v71&v63
				*(*int64)(unsafe.Add(mBase, uint32(v14)+8)) = v48*int64(0) + v48>>(uint(int64(63))%64)*int64(86400000000) + v58*v60 + int64(base.Ui64(v71)>>(uint(v57)%64)) + int64(base.Ui64(v78)>>(uint(v57)%64))
				*(*int64)(unsafe.Add(mBase, uint32(v14))) = v67&v63 | v78<<(uint(v57)%64)
				v89 = *(*int64)(unsafe.Add(mBase, uint32(v14)))
				v90 = v41 + v89
				v91 = int64(0)
				v95 = *(*int64)(unsafe.Add(mBase, uint32(v14)+8))
				v99 = base.I64_extend_i32_u(base.B2i32(base.Ui64(v90) < base.Ui64(v89))) + (v95 + v41>>(uint(int64(63))%64))
				if v99 == v91 {
					v104 = base.B2i32(v90 != v91)
				} else {
					v104 = base.B2i32(v91 < v99)
				}
				v105 = int64(0)
				*(*int32)(unsafe.Add(mBase, uint32(v35)+32)) = v104 - base.B2i32(v99 < v105)
				if v99|v90 == v105 {
					F_errstart_cold(m, int32(21), int32(0))
					mBase = m.M
					v171 = m.ExcPending
					if v171 != 0 {
						return int32(0)
					} else {
						F_errcode(m, int32(50856066))
						mBase = m.M
						v174 = m.ExcPending
						if v174 != 0 {
							return int32(0)
						} else {
							F_errmsg(m, int32(_a_F_generate_series_timestamp_1), int32(0))
							mBase = m.M
							v178 = m.ExcPending
							if v178 != 0 {
								return int32(0)
							} else {
								F_errfinish(m, int32(_a_F_generate_series_timestamp_2), int32(_a_F_generate_series_timestamp_3), int32(_a_F_generate_series_timestamp_4))
								mBase = m.M
								v183 = m.ExcPending
								if v183 != 0 {
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
					v112 = base.I32_wrap_i64(v39)
					if v44 != int32(2147483647) {
						v115 = int32(-2147483648)
						if base.B2i32(v44 != v115)|base.B2i32(v41 != int64(-9223372036854775807-1))|base.B2i32(v112 != v115) != 0 {
							*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v35
							*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0])) = v30
							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
							v142 = *(*int64)(unsafe.Add(mBase, uint32(v141)+8))
							v143 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+32))
							if int32(0) < v144 {
								if v143 <= v142 {
									v151 = F_Int64GetDatum(m, v143)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return int32(0)
									} else {
										v155 = F_DirectFunctionCall2Coll(m, int32(1267), int32(0), v151, v141+int32(16))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											v157 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
											*(*int64)(unsafe.Add(mBase, uint32(v141))) = v157
											v159 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
											*(*int64)(unsafe.Add(mBase, uint32(v140))) = v159 + int64(1)
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(1)
											v166 = F_Int64GetDatum(m, v143)
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												v192 = v166
												m.G0 = v14 + int32(16)
												return v192
											}
										}
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v185 = m.ExcPending
									if v185 != 0 {
										return int32(0)
									} else {
										v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = int32(2)
										v189 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
										v192 = int32(0)
										m.G0 = v14 + int32(16)
										return v192
									}
								}
							} else {
								if v143 < v142 {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v185 = m.ExcPending
									if v185 != 0 {
										return int32(0)
									} else {
										v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = int32(2)
										v189 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
										v192 = int32(0)
										m.G0 = v14 + int32(16)
										return v192
									}
								} else {
									v151 = F_Int64GetDatum(m, v143)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return int32(0)
									} else {
										v155 = F_DirectFunctionCall2Coll(m, int32(1267), int32(0), v151, v141+int32(16))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											v157 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
											*(*int64)(unsafe.Add(mBase, uint32(v141))) = v157
											v159 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
											*(*int64)(unsafe.Add(mBase, uint32(v140))) = v159 + int64(1)
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(1)
											v166 = F_Int64GetDatum(m, v143)
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												v192 = v166
												m.G0 = v14 + int32(16)
												return v192
											}
										}
									}
								}
							}
						} else {
							F_errstart_cold(m, int32(21), int32(0))
							mBase = m.M
							v200 = m.ExcPending
							if v200 != 0 {
								return int32(0)
							} else {
								F_errcode(m, int32(50856066))
								mBase = m.M
								v203 = m.ExcPending
								if v203 != 0 {
									return int32(0)
								} else {
									F_errmsg(m, int32(_a_F_generate_series_timestamp_5), int32(0))
									mBase = m.M
									v207 = m.ExcPending
									if v207 != 0 {
										return int32(0)
									} else {
										F_errfinish(m, int32(_a_F_generate_series_timestamp_2), int32(_a_F_generate_series_timestamp_6), int32(_a_F_generate_series_timestamp_4))
										mBase = m.M
										v212 = m.ExcPending
										if v212 != 0 {
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
						if v41 != int64(9223372036854775807) {
							*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v35
							*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0])) = v30
							v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
							v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
							v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
							v142 = *(*int64)(unsafe.Add(mBase, uint32(v141)+8))
							v143 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
							v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+32))
							if int32(0) < v144 {
								if v143 <= v142 {
									v151 = F_Int64GetDatum(m, v143)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return int32(0)
									} else {
										v155 = F_DirectFunctionCall2Coll(m, int32(1267), int32(0), v151, v141+int32(16))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											v157 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
											*(*int64)(unsafe.Add(mBase, uint32(v141))) = v157
											v159 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
											*(*int64)(unsafe.Add(mBase, uint32(v140))) = v159 + int64(1)
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(1)
											v166 = F_Int64GetDatum(m, v143)
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												v192 = v166
												m.G0 = v14 + int32(16)
												return v192
											}
										}
									}
								} else {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v185 = m.ExcPending
									if v185 != 0 {
										return int32(0)
									} else {
										v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = int32(2)
										v189 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
										v192 = int32(0)
										m.G0 = v14 + int32(16)
										return v192
									}
								}
							} else {
								if v143 < v142 {
									F_end_MultiFuncCall(m, l0)
									mBase = m.M
									v185 = m.ExcPending
									if v185 != 0 {
										return int32(0)
									} else {
										v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
										*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = int32(2)
										v189 = int32(1)
										*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
										v192 = int32(0)
										m.G0 = v14 + int32(16)
										return v192
									}
								} else {
									v151 = F_Int64GetDatum(m, v143)
									mBase = m.M
									v152 = m.ExcPending
									if v152 != 0 {
										return int32(0)
									} else {
										v155 = F_DirectFunctionCall2Coll(m, int32(1267), int32(0), v151, v141+int32(16))
										mBase = m.M
										v156 = m.ExcPending
										if v156 != 0 {
											return int32(0)
										} else {
											v157 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
											*(*int64)(unsafe.Add(mBase, uint32(v141))) = v157
											v159 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
											*(*int64)(unsafe.Add(mBase, uint32(v140))) = v159 + int64(1)
											v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(1)
											v166 = F_Int64GetDatum(m, v143)
											mBase = m.M
											v167 = m.ExcPending
											if v167 != 0 {
												return int32(0)
											} else {
												v192 = v166
												m.G0 = v14 + int32(16)
												return v192
											}
										}
									}
								}
							}
						} else {
							if v112 == int32(2147483647) {
								F_errstart_cold(m, int32(21), int32(0))
								mBase = m.M
								v200 = m.ExcPending
								if v200 != 0 {
									return int32(0)
								} else {
									F_errcode(m, int32(50856066))
									mBase = m.M
									v203 = m.ExcPending
									if v203 != 0 {
										return int32(0)
									} else {
										F_errmsg(m, int32(_a_F_generate_series_timestamp_5), int32(0))
										mBase = m.M
										v207 = m.ExcPending
										if v207 != 0 {
											return int32(0)
										} else {
											F_errfinish(m, int32(_a_F_generate_series_timestamp_2), int32(_a_F_generate_series_timestamp_6), int32(_a_F_generate_series_timestamp_4))
											mBase = m.M
											v212 = m.ExcPending
											if v212 != 0 {
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
								*(*int32)(unsafe.Add(mBase, uint32(v25)+16)) = v35
								*(*int32)(unsafe.Add(mBase, _c_F_generate_series_timestamp[0])) = v30
								v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
								v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
								v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
								v142 = *(*int64)(unsafe.Add(mBase, uint32(v141)+8))
								v143 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
								v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+32))
								if int32(0) < v144 {
									if v143 <= v142 {
										v151 = F_Int64GetDatum(m, v143)
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return int32(0)
										} else {
											v155 = F_DirectFunctionCall2Coll(m, int32(1267), int32(0), v151, v141+int32(16))
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												v157 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
												*(*int64)(unsafe.Add(mBase, uint32(v141))) = v157
												v159 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
												*(*int64)(unsafe.Add(mBase, uint32(v140))) = v159 + int64(1)
												v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(1)
												v166 = F_Int64GetDatum(m, v143)
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													v192 = v166
													m.G0 = v14 + int32(16)
													return v192
												}
											}
										}
									} else {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v185 = m.ExcPending
										if v185 != 0 {
											return int32(0)
										} else {
											v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = int32(2)
											v189 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
											v192 = int32(0)
											m.G0 = v14 + int32(16)
											return v192
										}
									}
								} else {
									if v143 < v142 {
										F_end_MultiFuncCall(m, l0)
										mBase = m.M
										v185 = m.ExcPending
										if v185 != 0 {
											return int32(0)
										} else {
											v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
											*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = int32(2)
											v189 = int32(1)
											*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
											v192 = int32(0)
											m.G0 = v14 + int32(16)
											return v192
										}
									} else {
										v151 = F_Int64GetDatum(m, v143)
										mBase = m.M
										v152 = m.ExcPending
										if v152 != 0 {
											return int32(0)
										} else {
											v155 = F_DirectFunctionCall2Coll(m, int32(1267), int32(0), v151, v141+int32(16))
											mBase = m.M
											v156 = m.ExcPending
											if v156 != 0 {
												return int32(0)
											} else {
												v157 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
												*(*int64)(unsafe.Add(mBase, uint32(v141))) = v157
												v159 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
												*(*int64)(unsafe.Add(mBase, uint32(v140))) = v159 + int64(1)
												v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
												*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(1)
												v166 = F_Int64GetDatum(m, v143)
												mBase = m.M
												v167 = m.ExcPending
												if v167 != 0 {
													return int32(0)
												} else {
													v192 = v166
													m.G0 = v14 + int32(16)
													return v192
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
	} else {
		v139 = *(*int32)(unsafe.Add(mBase, uint32(l0)))
		v140 = *(*int32)(unsafe.Add(mBase, uint32(v139)+16))
		v141 = *(*int32)(unsafe.Add(mBase, uint32(v140)+16))
		v142 = *(*int64)(unsafe.Add(mBase, uint32(v141)+8))
		v143 = *(*int64)(unsafe.Add(mBase, uint32(v141)))
		v144 = *(*int32)(unsafe.Add(mBase, uint32(v141)+32))
		if int32(0) < v144 {
			if v143 <= v142 {
				v151 = F_Int64GetDatum(m, v143)
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return int32(0)
				} else {
					v155 = F_DirectFunctionCall2Coll(m, int32(1267), int32(0), v151, v141+int32(16))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int32(0)
					} else {
						v157 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
						*(*int64)(unsafe.Add(mBase, uint32(v141))) = v157
						v159 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
						*(*int64)(unsafe.Add(mBase, uint32(v140))) = v159 + int64(1)
						v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(1)
						v166 = F_Int64GetDatum(m, v143)
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int32(0)
						} else {
							v192 = v166
							m.G0 = v14 + int32(16)
							return v192
						}
					}
				}
			} else {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
					return int32(0)
				} else {
					v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = int32(2)
					v189 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
					v192 = int32(0)
					m.G0 = v14 + int32(16)
					return v192
				}
			}
		} else {
			if v143 < v142 {
				F_end_MultiFuncCall(m, l0)
				mBase = m.M
				v185 = m.ExcPending
				if v185 != 0 {
					return int32(0)
				} else {
					v186 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
					*(*int32)(unsafe.Add(mBase, uint32(v186)+20)) = int32(2)
					v189 = int32(1)
					*(*uint8)(unsafe.Add(mBase, uint32(l0)+16)) = uint8(v189)
					v192 = int32(0)
					m.G0 = v14 + int32(16)
					return v192
				}
			} else {
				v151 = F_Int64GetDatum(m, v143)
				mBase = m.M
				v152 = m.ExcPending
				if v152 != 0 {
					return int32(0)
				} else {
					v155 = F_DirectFunctionCall2Coll(m, int32(1267), int32(0), v151, v141+int32(16))
					mBase = m.M
					v156 = m.ExcPending
					if v156 != 0 {
						return int32(0)
					} else {
						v157 = *(*int64)(unsafe.Add(mBase, uint32(v155)))
						*(*int64)(unsafe.Add(mBase, uint32(v141))) = v157
						v159 = *(*int64)(unsafe.Add(mBase, uint32(v140)))
						*(*int64)(unsafe.Add(mBase, uint32(v140))) = v159 + int64(1)
						v163 = *(*int32)(unsafe.Add(mBase, uint32(l0)+8))
						*(*int32)(unsafe.Add(mBase, uint32(v163)+20)) = int32(1)
						v166 = F_Int64GetDatum(m, v143)
						mBase = m.M
						v167 = m.ExcPending
						if v167 != 0 {
							return int32(0)
						} else {
							v192 = v166
							m.G0 = v14 + int32(16)
							return v192
						}
					}
				}
			}
		}
	}
}
