package p2

import (
	base "github.com/shibukawa/pgmem/internal/aot/pgaot/base"
	"unsafe"
)

func F_LogLogicalInvalidations(m *base.Module) {
	mBase := m.M
	_ = mBase
	var v5 int32
	_ = v5
	var v7 int32
	_ = v7
	var v10 int32
	_ = v10
	var v13 int32
	_ = v13
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v17 int32
	_ = v17
	var v19 int32
	_ = v19
	var v24 int32
	_ = v24
	var v29 int32
	_ = v29
	var v30 int32
	_ = v30
	var v31 int32
	_ = v31
	var v32 int32
	_ = v32
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v43 int32
	_ = v43
	var v44 int32
	_ = v44
	var v45 int32
	_ = v45
	var v46 int32
	_ = v46
	var v50 int32
	_ = v50
	var v51 int32
	_ = v51
	var v57 int32
	_ = v57
	var v60 int64
	_ = v60
	var v61 int32
	_ = v61
	v5 = m.G0
	v7 = v5 - int32(16)
	m.G0 = v7
	v10 = *(*int32)(unsafe.Add(mBase, _c_F_LogLogicalInvalidations[0]))
	if v10 == int32(0) {
		m.G0 = v7 + int32(16)
		return
	} else {
		v13 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
		v14 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
		v16 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
		v17 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
		v19 = v13 - v14 + (v16 - v17)
		if v19 <= int32(0) {
			m.G0 = v7 + int32(16)
			return
		} else {
			*(*int32)(unsafe.Add(mBase, uint32(v7)+12)) = v19
			F_XLogBeginInsert(m)
			mBase = m.M
			v24 = m.ExcPending
			if v24 != 0 {
				return
			} else {
				F_XLogRegisterData(m, v7+int32(12), int32(4))
				mBase = m.M
				v29 = m.ExcPending
				if v29 != 0 {
					return
				} else {
					v30 = *(*int32)(unsafe.Add(mBase, uint32(v10)+8))
					v31 = *(*int32)(unsafe.Add(mBase, uint32(v10)))
					v32 = v30 - v31
					if int32(0) < v32 {
						v36 = *(*int32)(unsafe.Add(mBase, _c_F_LogLogicalInvalidations[1]))
						v37 = int32(4)
						F_XLogRegisterData(m, v36+v31<<(uint(v37)%32), v32<<(uint(v37)%32))
						mBase = m.M
						v43 = m.ExcPending
						if v43 != 0 {
							return
						} else {
							v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
							v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
							v46 = v44 - v45
							if int32(0) < v46 {
								v50 = *(*int32)(unsafe.Add(mBase, _c_F_LogLogicalInvalidations[2]))
								v51 = int32(4)
								F_XLogRegisterData(m, v50+v45<<(uint(v51)%32), v46<<(uint(v51)%32))
								mBase = m.M
								v57 = m.ExcPending
								if v57 != 0 {
									return
								} else {
									v60 = F_XLogInsert(m, int32(1), int32(96))
									mBase = m.M
									v61 = m.ExcPending
									if v61 != 0 {
										return
									} else {
										m.G0 = v7 + int32(16)
										return
									}
								}
							} else {
								v60 = F_XLogInsert(m, int32(1), int32(96))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						}
					} else {
						v44 = *(*int32)(unsafe.Add(mBase, uint32(v10)+12))
						v45 = *(*int32)(unsafe.Add(mBase, uint32(v10)+4))
						v46 = v44 - v45
						if int32(0) < v46 {
							v50 = *(*int32)(unsafe.Add(mBase, _c_F_LogLogicalInvalidations[2]))
							v51 = int32(4)
							F_XLogRegisterData(m, v50+v45<<(uint(v51)%32), v46<<(uint(v51)%32))
							mBase = m.M
							v57 = m.ExcPending
							if v57 != 0 {
								return
							} else {
								v60 = F_XLogInsert(m, int32(1), int32(96))
								mBase = m.M
								v61 = m.ExcPending
								if v61 != 0 {
									return
								} else {
									m.G0 = v7 + int32(16)
									return
								}
							}
						} else {
							v60 = F_XLogInsert(m, int32(1), int32(96))
							mBase = m.M
							v61 = m.ExcPending
							if v61 != 0 {
								return
							} else {
								m.G0 = v7 + int32(16)
								return
							}
						}
					}
				}
			}
		}
	}
}
func F_log(m *base.Module, l0 float64) float64 {
	mBase := m.M
	_ = mBase
	var v11 int64
	_ = v11
	var v21 float64
	_ = v21
	var v23 float64
	_ = v23
	var v25 float64
	_ = v25
	var v28 float64
	_ = v28
	var v29 float64
	_ = v29
	var v30 float64
	_ = v30
	var v31 float64
	_ = v31
	var v32 float64
	_ = v32
	var v34 float64
	_ = v34
	var v37 float64
	_ = v37
	var v40 float64
	_ = v40
	var v43 float64
	_ = v43
	var v49 float64
	_ = v49
	var v52 float64
	_ = v52
	var v55 float64
	_ = v55
	var v61 float64
	_ = v61
	var v64 float64
	_ = v64
	var v67 float64
	_ = v67
	var v85 int32
	_ = v85
	var v92 float64
	_ = v92
	var v94 int32
	_ = v94
	var v104 int32
	_ = v104
	var v113 float64
	_ = v113
	var v121 int64
	_ = v121
	var v123 int64
	_ = v123
	var v126 float64
	_ = v126
	var v128 float64
	_ = v128
	var v136 int32
	_ = v136
	var v137 float64
	_ = v137
	var v138 float64
	_ = v138
	var v139 float64
	_ = v139
	var v144 float64
	_ = v144
	var v146 float64
	_ = v146
	var v148 float64
	_ = v148
	var v149 float64
	_ = v149
	var v150 float64
	_ = v150
	var v153 float64
	_ = v153
	var v156 float64
	_ = v156
	var v160 float64
	_ = v160
	var v163 float64
	_ = v163
	var v168 float64
	_ = v168
	var v171 float64
	_ = v171
	var v179 float64
	_ = v179
	v11 = base.I64_reinterpret_f64(l0)
	if base.Ui64(v11-int64(4606619468846596096)) <= base.Ui64(int64(854320534781951)) {
		if v11 == int64(4607182418800017408) {
			return float64(0)
		} else {
			v21 = base.F64_add(l0, float64(-1))
			v23 = base.F64_mul(v21, float64(1.34217728e+08))
			v25 = base.F64_sub(base.F64_add(v21, v23), v23)
			v28 = *(*float64)(unsafe.Add(mBase, _c_F_log[0]))
			v29 = base.F64_mul(base.F64_mul(v25, v25), v28)
			v30 = base.F64_add(v21, v29)
			v31 = base.F64_mul(v21, v21)
			v32 = base.F64_mul(v21, v31)
			v34 = *(*float64)(unsafe.Add(mBase, _c_F_log[1]))
			v37 = *(*float64)(unsafe.Add(mBase, _c_F_log[2]))
			v40 = *(*float64)(unsafe.Add(mBase, _c_F_log[3]))
			v43 = *(*float64)(unsafe.Add(mBase, _c_F_log[4]))
			v49 = *(*float64)(unsafe.Add(mBase, _c_F_log[5]))
			v52 = *(*float64)(unsafe.Add(mBase, _c_F_log[6]))
			v55 = *(*float64)(unsafe.Add(mBase, _c_F_log[7]))
			v61 = *(*float64)(unsafe.Add(mBase, _c_F_log[8]))
			v64 = *(*float64)(unsafe.Add(mBase, _c_F_log[9]))
			v67 = *(*float64)(unsafe.Add(mBase, _c_F_log[10]))
			return base.F64_add(v30, base.F64_add(base.F64_mul(v32, base.F64_add(base.F64_mul(v32, base.F64_add(base.F64_mul(v32, base.F64_add(base.F64_mul(v32, v34), base.F64_add(base.F64_mul(v31, v37), base.F64_add(base.F64_mul(v21, v40), v43)))), base.F64_add(base.F64_mul(v31, v49), base.F64_add(base.F64_mul(v21, v52), v55)))), base.F64_add(base.F64_mul(v31, v61), base.F64_add(base.F64_mul(v21, v64), v67)))), base.F64_add(base.F64_mul(base.F64_mul(base.F64_sub(v21, v25), v28), base.F64_add(v21, v25)), base.F64_add(v29, base.F64_sub(v21, v30)))))
		}
	} else {
		v85 = base.I32_wrap_i64(int64(base.Ui64(base.I64_reinterpret_f64(l0)) >> (uint(int64(48)) % 64)))
		if base.Ui32(v85-int32(_a_F_log_0)) <= base.Ui32(int32(-32737)) {
			if base.F64_eq(l0, float64(0)) != 0 {
				v92 = float64(-1)
				v94 = m.G0
				*(*float64)(unsafe.Add(mBase, uint32(v94-int32(16))+8)) = v92
				return base.F64_div(v92, float64(0))
			} else {
				if v11 == int64(9218868437227405312) {
					v179 = l0
					return v179
				} else {
					v104 = int32(_a_F_log_0)
					if base.B2i32(v85&v104 != v104)&base.B2i32(base.Ui32(v85) <= base.Ui32(int32(_a_F_log_1))) == int32(0) {
						v113 = base.F64_sub(l0, l0)
						return base.F64_div(v113, v113)
					} else {
						v121 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15))) - int64(234187180623265792)
						v123 = v121 - int64(4604367669032910848)
						v126 = base.F64_convert_i64_s(v123 >> (uint(int64(52)) % 64))
						v128 = *(*float64)(unsafe.Add(mBase, _c_F_log[11]))
						v136 = base.I32_wrap_i64(int64(base.Ui64(v123)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(4)) % 32)
						v137 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[12])))
						v138 = base.F64_add(base.F64_mul(v126, v128), v137)
						v139 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[13])))
						v144 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[14])))
						v146 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[15])))
						v148 = base.F64_mul(v139, base.F64_sub(base.F64_sub(base.F64_reinterpret_i64(v121-v123&int64(-4503599627370496)), v144), v146))
						v149 = base.F64_add(v138, v148)
						v150 = base.F64_mul(v148, v148)
						v153 = *(*float64)(unsafe.Add(mBase, _c_F_log[16]))
						v156 = *(*float64)(unsafe.Add(mBase, _c_F_log[17]))
						v160 = *(*float64)(unsafe.Add(mBase, _c_F_log[18]))
						v163 = *(*float64)(unsafe.Add(mBase, _c_F_log[19]))
						v168 = *(*float64)(unsafe.Add(mBase, _c_F_log[20]))
						v171 = *(*float64)(unsafe.Add(mBase, _c_F_log[21]))
						v179 = base.F64_add(v149, base.F64_add(base.F64_mul(base.F64_mul(v148, v150), base.F64_add(base.F64_mul(v150, base.F64_add(base.F64_mul(v148, v153), v156)), base.F64_add(base.F64_mul(v148, v160), v163))), base.F64_add(base.F64_mul(v150, v168), base.F64_add(base.F64_mul(v126, v171), base.F64_add(v148, base.F64_sub(v138, v149))))))
						return v179
					}
				}
			}
		} else {
			v121 = v11
			v123 = v121 - int64(4604367669032910848)
			v126 = base.F64_convert_i64_s(v123 >> (uint(int64(52)) % 64))
			v128 = *(*float64)(unsafe.Add(mBase, _c_F_log[11]))
			v136 = base.I32_wrap_i64(int64(base.Ui64(v123)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(4)) % 32)
			v137 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[12])))
			v138 = base.F64_add(base.F64_mul(v126, v128), v137)
			v139 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[13])))
			v144 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[14])))
			v146 = *(*float64)(unsafe.Add(mBase, uint32(v136)+uint32(_c_F_log[15])))
			v148 = base.F64_mul(v139, base.F64_sub(base.F64_sub(base.F64_reinterpret_i64(v121-v123&int64(-4503599627370496)), v144), v146))
			v149 = base.F64_add(v138, v148)
			v150 = base.F64_mul(v148, v148)
			v153 = *(*float64)(unsafe.Add(mBase, _c_F_log[16]))
			v156 = *(*float64)(unsafe.Add(mBase, _c_F_log[17]))
			v160 = *(*float64)(unsafe.Add(mBase, _c_F_log[18]))
			v163 = *(*float64)(unsafe.Add(mBase, _c_F_log[19]))
			v168 = *(*float64)(unsafe.Add(mBase, _c_F_log[20]))
			v171 = *(*float64)(unsafe.Add(mBase, _c_F_log[21]))
			v179 = base.F64_add(v149, base.F64_add(base.F64_mul(base.F64_mul(v148, v150), base.F64_add(base.F64_mul(v150, base.F64_add(base.F64_mul(v148, v153), v156)), base.F64_add(base.F64_mul(v148, v160), v163))), base.F64_add(base.F64_mul(v150, v168), base.F64_add(base.F64_mul(v126, v171), base.F64_add(v148, base.F64_sub(v138, v149))))))
			return v179
		}
	}
}
func F_log_disconnections(m *base.Module, l0 int32, l1 int32) {
	mBase := m.M
	_ = mBase
	var v10 int32
	_ = v10
	var v12 int32
	_ = v12
	var v15 int32
	_ = v15
	var v17 int64
	_ = v17
	var v21 int32
	_ = v21
	var v22 int32
	_ = v22
	var v23 int32
	_ = v23
	var v26 int64
	_ = v26
	var v27 int64
	_ = v27
	var v42 int64
	_ = v42
	var v46 int64
	_ = v46
	var v47 int64
	_ = v47
	var v54 int32
	_ = v54
	var v55 int32
	_ = v55
	var v58 int32
	_ = v58
	var v59 int32
	_ = v59
	var v60 int32
	_ = v60
	var v63 int32
	_ = v63
	var v65 int32
	_ = v65
	var v68 int32
	_ = v68
	var v69 int32
	_ = v69
	var v70 int32
	_ = v70
	var v71 int32
	_ = v71
	var v72 int32
	_ = v72
	var v73 int32
	_ = v73
	var v74 int32
	_ = v74
	var v81 int32
	_ = v81
	var v84 int32
	_ = v84
	var v88 int32
	_ = v88
	var v89 int32
	_ = v89
	var v99 int32
	_ = v99
	var v104 int32
	_ = v104
	v10 = m.G0
	v12 = v10 - int32(48)
	m.G0 = v12
	v15 = *(*int32)(unsafe.Add(mBase, _c_F_log_disconnections[0]))
	v17 = *(*int64)(unsafe.Add(mBase, _c_F_log_disconnections[1]))
	v21 = m.G0
	v22 = int32(16)
	v23 = v21 - v22
	m.G0 = v23
	F_gettimeofday(m, v23)
	mBase = m.M
	v26 = *(*int64)(unsafe.Add(mBase, uint32(v23)))
	v27 = int64(*(*int32)(unsafe.Add(mBase, uint32(v23)+8)))
	m.G0 = v23 + v22
	v42 = v27 + v26*int64(1000000) - int64(946684800000000) - v17
	if v42 <= int64(0) {
		v54 = int32(0)
		v55 = int32(0)
	} else {
		v46 = int64(1000000)
		v47 = base.I64_div_u_s(v42, v46)
		v54 = base.I32_wrap_i64(v47)
		v55 = base.I32_wrap_i64(v42 - v47*v46)
	}
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(44)))) = v54
	*(*int32)(unsafe.Add(mBase, uint32(v12+int32(40)))) = v55
	v58 = *(*int32)(unsafe.Add(mBase, uint32(v12)+44))
	v59 = int32(3600)
	v60 = base.I32_div_s(v58, v59)
	v63 = v58 - v60*v59
	*(*int32)(unsafe.Add(mBase, uint32(v12)+44)) = v63
	v65 = *(*int32)(unsafe.Add(mBase, uint32(v12)+40))
	v68 = F_errstart(m, int32(15), int32(0))
	mBase = m.M
	v69 = m.ExcPending
	if v69 != 0 {
		return
	} else {
		if v68 != 0 {
			v70 = *(*int32)(unsafe.Add(mBase, uint32(v15)+292))
			v71 = int32(*(*uint8)(unsafe.Add(mBase, uint32(v70))))
			v72 = *(*int32)(unsafe.Add(mBase, uint32(v15)+364))
			v73 = *(*int32)(unsafe.Add(mBase, uint32(v15)+360))
			v74 = *(*int32)(unsafe.Add(mBase, uint32(v15)+276))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+32)) = v70
			*(*int32)(unsafe.Add(mBase, uint32(v12)+24)) = v74
			*(*int32)(unsafe.Add(mBase, uint32(v12)+20)) = v73
			*(*int32)(unsafe.Add(mBase, uint32(v12)+16)) = v72
			if v71 != 0 {
				v81 = int32(_a_F_log_disconnections_0)
			} else {
				v81 = int32(_a_F_log_disconnections_1)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v81
			v84 = base.I32_div_s(v65, int32(1000))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v84
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v60
			v88 = int32(60)
			v89 = base.I32_div_s(base.I32_extend16_s(v63), v88)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = base.I32_extend16_s(v89)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_extend16_s(v63 - v89*v88)
			F_errmsg(m, int32(_a_F_log_disconnections_2), v12)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return
			} else {
				F_errfinish(m, int32(_a_F_log_disconnections_3), int32(_a_F_log_disconnections_4), int32(_a_F_log_disconnections_5))
				mBase = m.M
				v104 = m.ExcPending
				if v104 != 0 {
					return
				} else {
					m.G0 = v12 + int32(48)
					return
				}
			}
		} else {
			m.G0 = v12 + int32(48)
			return
		}
	}
}
func F_log_newpage_range(m *base.Module, l0 int32, l1 int32, l2 int32, l3 int32) {
	mBase := m.M
	_ = mBase
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v19 int32
	_ = v19
	var v22 int32
	_ = v22
	var v24 int32
	_ = v24
	var v26 int32
	_ = v26
	var v27 int32
	_ = v27
	var v31 int32
	_ = v31
	var v33 int32
	_ = v33
	var v34 int32
	_ = v34
	var v36 int32
	_ = v36
	var v37 int32
	_ = v37
	var v51 int32
	_ = v51
	var v54 int32
	_ = v54
	var v56 int32
	_ = v56
	var v62 int32
	_ = v62
	var v79 int32
	_ = v79
	var v83 int32
	_ = v83
	var v85 int32
	_ = v85
	var v86 int32
	_ = v86
	var v97 int32
	_ = v97
	var v106 int32
	_ = v106
	var v112 int32
	_ = v112
	var v114 int32
	_ = v114
	var v121 int32
	_ = v121
	var v125 int32
	_ = v125
	var v130 int32
	_ = v130
	var v132 int32
	_ = v132
	var v133 int32
	_ = v133
	var v136 int32
	_ = v136
	var v140 int32
	_ = v140
	var v146 int32
	_ = v146
	var v148 int32
	_ = v148
	var v154 int32
	_ = v154
	var v155 int32
	_ = v155
	var v163 int32
	_ = v163
	var v164 int32
	_ = v164
	var v168 int32
	_ = v168
	var v174 int32
	_ = v174
	var v175 int32
	_ = v175
	var v176 int32
	_ = v176
	var v178 int32
	_ = v178
	var v187 int32
	_ = v187
	var v200 int32
	_ = v200
	var v202 int32
	_ = v202
	var v206 int32
	_ = v206
	var v208 int32
	_ = v208
	var v210 int32
	_ = v210
	var v213 int64
	_ = v213
	var v214 int32
	_ = v214
	var v222 int32
	_ = v222
	var v235 int32
	_ = v235
	var v239 int32
	_ = v239
	var v245 int32
	_ = v245
	var v247 int32
	_ = v247
	var v253 int32
	_ = v253
	var v257 int32
	_ = v257
	var v259 int32
	_ = v259
	var v263 int64
	_ = v263
	var v264 int32
	_ = v264
	var v278 int32
	_ = v278
	var v280 int32
	_ = v280
	v14 = m.G0
	v16 = v14 - int32(128)
	m.G0 = v16
	v19 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[0]))
	if v19 <= int32(31) {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	v22 = int32(_a_F_log_newpage_range_0)
	v24 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[1]))
	v26 = F_repalloc(m, v24, int32(_a_F_log_newpage_range_1))
	mBase = m.M
	v27 = m.ExcPending
	if v27 != 0 {
		goto L4
	} else {
		goto L5
	}
L2:
	;
	goto L3
L3:
	;
	v79 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[2]))
	if v79 <= int32(19) {
		goto L15
	} else {
		goto L16
	}
L4:
	;
	return
L5:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[1])) = v26
	v31 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[0]))
	v33 = int32(_a_F_log_newpage_range_2)
	v34 = (int32(32) - v31) * v33
	v36 = v31 * v33
	v37 = v26 + v36
	if v37&int32(3)|base.B2i32(base.Ui32(int32(1024)) < base.Ui32(v34)) == int32(0) {
		goto L7
	} else {
		goto L8
	}
L6:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[0])) = int32(32)
	goto L3
L7:
	;
	if v31 == int32(32) {
		goto L6
	} else {
		goto L10
	}
L8:
	;
	v62 = v34
	goto L9
L9:
	;
	if v62 == int32(0) {
		goto L6
	} else {
		goto L14
	}
L10:
	;
	v51 = v26 + int32(_a_F_log_newpage_range_1)
	v54 = v36 + v26 + int32(4)
	if base.Ui32(v54) < base.Ui32(v51) {
		goto L11
	} else {
		goto L12
	}
L11:
	;
	v56 = v51
	goto L13
L12:
	;
	v56 = v54
	goto L13
L13:
	;
	v62 = (v26^int32(-1)-v36+v56)&int32(-4) + int32(4)
	goto L9
L14:
	;
	base.MemoryFill(m, v37, int32(0), v62)
	goto L6
L15:
	;
	v83 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[3]))
	v85 = F_repalloc(m, v83, int32(240))
	mBase = m.M
	v86 = m.ExcPending
	if v86 != 0 {
		goto L4
	} else {
		goto L18
	}
L16:
	;
	goto L17
L17:
	;
	if l2 == int32(0) {
		goto L19
	} else {
		goto L20
	}
L18:
	;
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[2])) = int32(20)
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[3])) = v85
	goto L17
L19:
	;
	m.G0 = v16 + int32(128)
	return
L20:
	;
	if l3 != 0 {
		goto L21
	} else {
		goto L22
	}
L21:
	;
	v97 = int32(9)
	goto L23
L22:
	;
	v97 = int32(1)
	goto L23
L23:
	;
	v106 = int32(0)
	goto L24
L24:
	;
	v112 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[4]))
	if v112 != 0 {
		goto L26
	} else {
		goto L27
	}
L25:
	;
	goto L19
L26:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v114 = m.ExcPending
	if v114 != 0 {
		goto L4
	} else {
		goto L29
	}
L27:
	;
	goto L28
L28:
	;
	if base.Ui32(l2) <= base.Ui32(v106) {
		goto L19
	} else {
		goto L30
	}
L29:
	;
	goto L28
L30:
	;
	v121 = int32(0)
	v125 = v106
	goto L31
L31:
	;
	v130 = int32(0)
	v132 = F_ReadBufferExtended(m, l0, l1, v125, v130, v130)
	mBase = m.M
	v133 = m.ExcPending
	if v133 != 0 {
		goto L4
	} else {
		goto L33
	}
L32:
	;
	if v164 == int32(0) {
		goto L19
	} else {
		goto L45
	}
L33:
	;
	F_LockBuffer(m, v132, int32(2))
	mBase = m.M
	v136 = m.ExcPending
	if v136 != 0 {
		goto L4
	} else {
		goto L34
	}
L34:
	;
	if v132 < int32(0) {
		goto L37
	} else {
		goto L38
	}
L35:
	;
	v168 = v125 + int32(1)
	if base.B2i32(v164 <= int32(31))&base.B2i32(base.Ui32(v168) < base.Ui32(l2)) != 0 {
		v121 = v164
		v125 = v168
		goto L31
	} else {
		goto L44
	}
L36:
	;
	v155 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v154)+14)))
	if v155 != 0 {
		goto L40
	} else {
		goto L41
	}
L37:
	;
	v140 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[5]))
	v146 = *(*int32)(unsafe.Add(mBase, uint32(v140+(v132^int32(-1))<<(uint(int32(2))%32))))
	v154 = v146
	goto L36
L38:
	;
	goto L39
L39:
	;
	v148 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[6]))
	v154 = v148 + v132<<(uint(int32(13))%32) + int32(-8192)
	goto L36
L40:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v121<<(uint(int32(2))%32)))) = v132
	v164 = v121 + int32(1)
	goto L35
L41:
	;
	goto L42
L42:
	;
	F_UnlockReleaseBuffer(m, v132)
	mBase = m.M
	v163 = m.ExcPending
	if v163 != 0 {
		goto L4
	} else {
		goto L43
	}
L43:
	;
	v164 = v121
	goto L35
L44:
	;
	goto L32
L45:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v174 = m.ExcPending
	if v174 != 0 {
		goto L4
	} else {
		goto L46
	}
L46:
	;
	v175 = int32(0)
	v176 = int32(_a_F_log_newpage_range_3)
	v178 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7])) = v178 + int32(1)
	if v175 < v164 {
		goto L48
	} else {
		goto L49
	}
L47:
	;
	v278 = int32(_a_F_log_newpage_range_3)
	v280 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7]))
	*(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[7])) = v280 - int32(1)
	if base.Ui32(v168) < base.Ui32(l2) {
		v106 = v168
		goto L24
	} else {
		goto L66
	}
L48:
	;
	v187 = v175
	goto L51
L49:
	;
	goto L50
L50:
	;
	v263 = F_XLogInsert(m, int32(0), int32(176))
	mBase = m.M
	v264 = m.ExcPending
	if v264 != 0 {
		goto L4
	} else {
		goto L65
	}
L51:
	;
	v200 = *(*int32)(unsafe.Add(mBase, uint32(v16+v187<<(uint(int32(2))%32))))
	F_MarkBufferDirty(m, v200)
	mBase = m.M
	v202 = m.ExcPending
	if v202 != 0 {
		goto L4
	} else {
		goto L53
	}
L52:
	;
	v210 = int32(0)
	v213 = F_XLogInsert(m, v210, int32(176))
	mBase = m.M
	v214 = m.ExcPending
	if v214 != 0 {
		goto L4
	} else {
		goto L56
	}
L53:
	;
	F_XLogRegisterBuffer(m, v187&int32(255), v200, v97)
	mBase = m.M
	v206 = m.ExcPending
	if v206 != 0 {
		goto L4
	} else {
		goto L54
	}
L54:
	;
	v208 = v187 + int32(1)
	if v208 != v164 {
		v187 = v208
		goto L51
	} else {
		goto L55
	}
L55:
	;
	goto L52
L56:
	;
	v222 = v210
	goto L57
L57:
	;
	v235 = *(*int32)(unsafe.Add(mBase, uint32(v16+v222<<(uint(int32(2))%32))))
	if v235 < int32(0) {
		goto L60
	} else {
		goto L61
	}
L58:
	;
	goto L47
L59:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v253)+4)) = base.I32_wrap_i64(v213)
	*(*int32)(unsafe.Add(mBase, uint32(v253))) = base.I32_wrap_i64(int64(base.Ui64(v213) >> (uint(int64(32)) % 64)))
	F_UnlockReleaseBuffer(m, v235)
	mBase = m.M
	v257 = m.ExcPending
	if v257 != 0 {
		goto L4
	} else {
		goto L63
	}
L60:
	;
	v239 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[5]))
	v245 = *(*int32)(unsafe.Add(mBase, uint32(v239+(v235^int32(-1))<<(uint(int32(2))%32))))
	v253 = v245
	goto L59
L61:
	;
	goto L62
L62:
	;
	v247 = *(*int32)(unsafe.Add(mBase, _c_F_log_newpage_range[6]))
	v253 = v247 + v235<<(uint(int32(13))%32) + int32(-8192)
	goto L59
L63:
	;
	v259 = v222 + int32(1)
	if v259 != v164 {
		v222 = v259
		goto L57
	} else {
		goto L64
	}
L64:
	;
	goto L58
L65:
	;
	goto L47
L66:
	;
	goto L25
}
