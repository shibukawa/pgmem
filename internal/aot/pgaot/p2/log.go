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
	v10 = *(*int32)(unsafe.Add(mBase, _consts[1157]))
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
						v36 = *(*int32)(unsafe.Add(mBase, _consts[1158]))
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
								v50 = *(*int32)(unsafe.Add(mBase, _consts[1159]))
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
							v50 = *(*int32)(unsafe.Add(mBase, _consts[1159]))
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
	var v14 int32
	_ = v14
	var v25 float64
	_ = v25
	var v27 float64
	_ = v27
	var v29 float64
	_ = v29
	var v32 float64
	_ = v32
	var v33 float64
	_ = v33
	var v34 float64
	_ = v34
	var v35 float64
	_ = v35
	var v36 float64
	_ = v36
	var v38 float64
	_ = v38
	var v41 float64
	_ = v41
	var v44 float64
	_ = v44
	var v47 float64
	_ = v47
	var v53 float64
	_ = v53
	var v56 float64
	_ = v56
	var v59 float64
	_ = v59
	var v65 float64
	_ = v65
	var v68 float64
	_ = v68
	var v71 float64
	_ = v71
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
	var v127 float64
	_ = v127
	var v129 float64
	_ = v129
	var v137 int32
	_ = v137
	var v140 float64
	_ = v140
	var v141 float64
	_ = v141
	var v144 float64
	_ = v144
	var v151 float64
	_ = v151
	var v155 float64
	_ = v155
	var v157 float64
	_ = v157
	var v158 float64
	_ = v158
	var v159 float64
	_ = v159
	var v162 float64
	_ = v162
	var v165 float64
	_ = v165
	var v169 float64
	_ = v169
	var v172 float64
	_ = v172
	var v177 float64
	_ = v177
	var v180 float64
	_ = v180
	var v188 float64
	_ = v188
	v11 = base.I64_reinterpret_f64(l0)
	v14 = base.I32_wrap_i64(int64(base.Ui64(v11) >> (uint(int64(48)) % 64)))
	if base.Ui64(v11-int64(4606619468846596096)) <= base.Ui64(int64(854320534781951)) {
		if v11 == int64(4607182418800017408) {
			return float64(0)
		} else {
			v25 = base.F64_add(l0, float64(-1))
			v27 = base.F64_mul(v25, float64(1.34217728e+08))
			v29 = base.F64_sub(base.F64_add(v25, v27), v27)
			v32 = *(*float64)(unsafe.Add(mBase, _consts[1401]))
			v33 = base.F64_mul(base.F64_mul(v29, v29), v32)
			v34 = base.F64_add(v25, v33)
			v35 = base.F64_mul(v25, v25)
			v36 = base.F64_mul(v25, v35)
			v38 = *(*float64)(unsafe.Add(mBase, _consts[1402]))
			v41 = *(*float64)(unsafe.Add(mBase, _consts[1403]))
			v44 = *(*float64)(unsafe.Add(mBase, _consts[1404]))
			v47 = *(*float64)(unsafe.Add(mBase, _consts[1405]))
			v53 = *(*float64)(unsafe.Add(mBase, _consts[1406]))
			v56 = *(*float64)(unsafe.Add(mBase, _consts[1407]))
			v59 = *(*float64)(unsafe.Add(mBase, _consts[1408]))
			v65 = *(*float64)(unsafe.Add(mBase, _consts[1409]))
			v68 = *(*float64)(unsafe.Add(mBase, _consts[1410]))
			v71 = *(*float64)(unsafe.Add(mBase, _consts[1411]))
			return base.F64_add(v34, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, base.F64_add(base.F64_mul(v36, v38), base.F64_add(base.F64_mul(v35, v41), base.F64_add(base.F64_mul(v25, v44), v47)))), base.F64_add(base.F64_mul(v35, v53), base.F64_add(base.F64_mul(v25, v56), v59)))), base.F64_add(base.F64_mul(v35, v65), base.F64_add(base.F64_mul(v25, v68), v71)))), base.F64_add(base.F64_mul(base.F64_mul(base.F64_sub(v25, v29), v32), base.F64_add(v25, v29)), base.F64_add(v33, base.F64_sub(v25, v34)))))
		}
	} else {
		if base.Ui32(v14-int32(32752)) <= base.Ui32(int32(-32737)) {
			if base.F64_eq(l0, float64(0)) != 0 {
				v92 = float64(-1)
				v94 = m.G0
				*(*float64)(unsafe.Add(mBase, uint32(v94-int32(16))+8)) = v92
				return base.F64_div(v92, float64(0))
			} else {
				if v11 == int64(9218868437227405312) {
					v188 = l0
					return v188
				} else {
					v104 = int32(32752)
					if base.B2i32(v14&v104 != v104)&base.B2i32(base.Ui32(v14) <= base.Ui32(int32(32767))) == int32(0) {
						v113 = base.F64_sub(l0, l0)
						return base.F64_div(v113, v113)
					} else {
						v121 = base.I64_reinterpret_f64(base.F64_mul(l0, float64(4.503599627370496e+15))) - int64(234187180623265792)
						v123 = v121 - int64(4604367669032910848)
						v127 = base.F64_convert_i32_s(base.I32_wrap_i64(v123 >> (uint(int64(52)) % 64)))
						v129 = *(*float64)(unsafe.Add(mBase, _consts[1412]))
						v137 = base.I32_wrap_i64(int64(base.Ui64(v123)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(4)) % 32)
						v140 = *(*float64)(unsafe.Add(mBase, uint32(v137)+uint32(_consts[1413])))
						v141 = base.F64_add(base.F64_mul(v127, v129), v140)
						v144 = *(*float64)(unsafe.Add(mBase, uint32(v137)+uint32(_consts[1414])))
						v151 = *(*float64)(unsafe.Add(mBase, uint32(v137)+uint32(_consts[1415])))
						v155 = *(*float64)(unsafe.Add(mBase, uint32(v137)+uint32(_consts[1416])))
						v157 = base.F64_mul(v144, base.F64_sub(base.F64_sub(base.F64_reinterpret_i64(v121-v123&int64(-4503599627370496)), v151), v155))
						v158 = base.F64_add(v141, v157)
						v159 = base.F64_mul(v157, v157)
						v162 = *(*float64)(unsafe.Add(mBase, _consts[1417]))
						v165 = *(*float64)(unsafe.Add(mBase, _consts[1418]))
						v169 = *(*float64)(unsafe.Add(mBase, _consts[1419]))
						v172 = *(*float64)(unsafe.Add(mBase, _consts[1420]))
						v177 = *(*float64)(unsafe.Add(mBase, _consts[1421]))
						v180 = *(*float64)(unsafe.Add(mBase, _consts[1422]))
						v188 = base.F64_add(v158, base.F64_add(base.F64_mul(base.F64_mul(v157, v159), base.F64_add(base.F64_mul(v159, base.F64_add(base.F64_mul(v157, v162), v165)), base.F64_add(base.F64_mul(v157, v169), v172))), base.F64_add(base.F64_mul(v159, v177), base.F64_add(base.F64_mul(v127, v180), base.F64_add(v157, base.F64_sub(v141, v158))))))
						return v188
					}
				}
			}
		} else {
			v121 = v11
			v123 = v121 - int64(4604367669032910848)
			v127 = base.F64_convert_i32_s(base.I32_wrap_i64(v123 >> (uint(int64(52)) % 64)))
			v129 = *(*float64)(unsafe.Add(mBase, _consts[1412]))
			v137 = base.I32_wrap_i64(int64(base.Ui64(v123)>>(uint(int64(45))%64))) & int32(127) << (uint(int32(4)) % 32)
			v140 = *(*float64)(unsafe.Add(mBase, uint32(v137)+uint32(_consts[1413])))
			v141 = base.F64_add(base.F64_mul(v127, v129), v140)
			v144 = *(*float64)(unsafe.Add(mBase, uint32(v137)+uint32(_consts[1414])))
			v151 = *(*float64)(unsafe.Add(mBase, uint32(v137)+uint32(_consts[1415])))
			v155 = *(*float64)(unsafe.Add(mBase, uint32(v137)+uint32(_consts[1416])))
			v157 = base.F64_mul(v144, base.F64_sub(base.F64_sub(base.F64_reinterpret_i64(v121-v123&int64(-4503599627370496)), v151), v155))
			v158 = base.F64_add(v141, v157)
			v159 = base.F64_mul(v157, v157)
			v162 = *(*float64)(unsafe.Add(mBase, _consts[1417]))
			v165 = *(*float64)(unsafe.Add(mBase, _consts[1418]))
			v169 = *(*float64)(unsafe.Add(mBase, _consts[1419]))
			v172 = *(*float64)(unsafe.Add(mBase, _consts[1420]))
			v177 = *(*float64)(unsafe.Add(mBase, _consts[1421]))
			v180 = *(*float64)(unsafe.Add(mBase, _consts[1422]))
			v188 = base.F64_add(v158, base.F64_add(base.F64_mul(base.F64_mul(v157, v159), base.F64_add(base.F64_mul(v159, base.F64_add(base.F64_mul(v157, v162), v165)), base.F64_add(base.F64_mul(v157, v169), v172))), base.F64_add(base.F64_mul(v159, v177), base.F64_add(base.F64_mul(v127, v180), base.F64_add(v157, base.F64_sub(v141, v158))))))
			return v188
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
	v15 = *(*int32)(unsafe.Add(mBase, _consts[475]))
	v17 = *(*int64)(unsafe.Add(mBase, _consts[881]))
	v21 = m.G0
	v22 = int32(16)
	v23 = v21 - v22
	m.G0 = v23
	F___gettimeofday(m, v23)
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
				v81 = int32(546004)
			} else {
				v81 = int32(756936)
			}
			*(*int32)(unsafe.Add(mBase, uint32(v12)+28)) = v81
			v84 = base.I32_div_s(v65, int32(1000))
			*(*int32)(unsafe.Add(mBase, uint32(v12)+12)) = v84
			*(*int32)(unsafe.Add(mBase, uint32(v12))) = v60
			v88 = int32(60)
			v89 = base.I32_div_s(base.I32_extend16_s(v63), v88)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+4)) = base.I32_extend16_s(v89)
			*(*int32)(unsafe.Add(mBase, uint32(v12)+8)) = base.I32_extend16_s(v63 - v89*v88)
			F_errmsg(m, int32(175418), v12)
			mBase = m.M
			v99 = m.ExcPending
			if v99 != 0 {
				return
			} else {
				F_errfinish(m, int32(493843), int32(5432), int32(141339))
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
	var v5 int32
	_ = v5
	var v14 int32
	_ = v14
	var v16 int32
	_ = v16
	var v21 int32
	_ = v21
	var v26 int32
	_ = v26
	var v32 int32
	_ = v32
	var v41 int32
	_ = v41
	var v43 int32
	_ = v43
	var v49 int32
	_ = v49
	var v51 int32
	_ = v51
	var v59 int32
	_ = v59
	var v61 int32
	_ = v61
	var v62 int32
	_ = v62
	var v65 int32
	_ = v65
	var v69 int32
	_ = v69
	var v75 int32
	_ = v75
	var v77 int32
	_ = v77
	var v83 int32
	_ = v83
	var v84 int32
	_ = v84
	var v92 int32
	_ = v92
	var v93 int32
	_ = v93
	var v97 int32
	_ = v97
	var v103 int32
	_ = v103
	var v104 int32
	_ = v104
	var v105 int32
	_ = v105
	var v107 int32
	_ = v107
	var v115 int64
	_ = v115
	var v116 int32
	_ = v116
	var v121 int32
	_ = v121
	var v133 int32
	_ = v133
	var v135 int32
	_ = v135
	var v139 int32
	_ = v139
	var v141 int32
	_ = v141
	var v143 int32
	_ = v143
	var v146 int64
	_ = v146
	var v147 int32
	_ = v147
	var v156 int32
	_ = v156
	var v168 int32
	_ = v168
	var v172 int32
	_ = v172
	var v178 int32
	_ = v178
	var v180 int32
	_ = v180
	var v186 int32
	_ = v186
	var v190 int32
	_ = v190
	var v192 int32
	_ = v192
	var v207 int32
	_ = v207
	var v209 int32
	_ = v209
	v5 = int32(0)
	v14 = m.G0
	v16 = v14 - int32(128)
	m.G0 = v16
	F_XLogEnsureRecordSpace(m, int32(31), v5)
	mBase = m.M
	v21 = m.ExcPending
	if v21 != 0 {
		goto L1
	} else {
		goto L2
	}
L1:
	;
	return
L2:
	;
	if l2 == int32(0) {
		goto L3
	} else {
		goto L4
	}
L3:
	;
	m.G0 = v16 + int32(128)
	return
L4:
	;
	if l3 != 0 {
		goto L5
	} else {
		goto L6
	}
L5:
	;
	v26 = int32(9)
	goto L7
L6:
	;
	v26 = int32(1)
	goto L7
L7:
	;
	v32 = v5
	goto L8
L8:
	;
	v41 = *(*int32)(unsafe.Add(mBase, _consts[1]))
	if v41 != 0 {
		goto L10
	} else {
		goto L11
	}
L9:
	;
	goto L3
L10:
	;
	F_ProcessInterrupts(m)
	mBase = m.M
	v43 = m.ExcPending
	if v43 != 0 {
		goto L1
	} else {
		goto L13
	}
L11:
	;
	goto L12
L12:
	;
	if base.Ui32(l2) <= base.Ui32(v32) {
		goto L3
	} else {
		goto L14
	}
L13:
	;
	goto L12
L14:
	;
	v49 = int32(0)
	v51 = v32
	goto L15
L15:
	;
	v59 = int32(0)
	v61 = F_ReadBufferExtended(m, l0, l1, v51, v59, v59)
	mBase = m.M
	v62 = m.ExcPending
	if v62 != 0 {
		goto L1
	} else {
		goto L17
	}
L16:
	;
	if v93 == int32(0) {
		goto L3
	} else {
		goto L29
	}
L17:
	;
	F_LockBuffer(m, v61, int32(2))
	mBase = m.M
	v65 = m.ExcPending
	if v65 != 0 {
		goto L1
	} else {
		goto L18
	}
L18:
	;
	if v61 < int32(0) {
		goto L21
	} else {
		goto L22
	}
L19:
	;
	v97 = v51 + int32(1)
	if base.B2i32(v93 <= int32(31))&base.B2i32(base.Ui32(v97) < base.Ui32(l2)) != 0 {
		v49 = v93
		v51 = v97
		goto L15
	} else {
		goto L28
	}
L20:
	;
	v84 = int32(*(*uint16)(unsafe.Add(mBase, uint32(v83)+14)))
	if v84 != 0 {
		goto L24
	} else {
		goto L25
	}
L21:
	;
	v69 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v75 = *(*int32)(unsafe.Add(mBase, uint32(v69+(v61^int32(-1))<<(uint(int32(2))%32))))
	v83 = v75
	goto L20
L22:
	;
	goto L23
L23:
	;
	v77 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v83 = v77 + v61<<(uint(int32(13))%32) + int32(-8192)
	goto L20
L24:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v16+v49<<(uint(int32(2))%32)))) = v61
	v93 = v49 + int32(1)
	goto L19
L25:
	;
	goto L26
L26:
	;
	F_UnlockReleaseBuffer(m, v61)
	mBase = m.M
	v92 = m.ExcPending
	if v92 != 0 {
		goto L1
	} else {
		goto L27
	}
L27:
	;
	v93 = v49
	goto L19
L28:
	;
	goto L16
L29:
	;
	F_XLogBeginInsert(m)
	mBase = m.M
	v103 = m.ExcPending
	if v103 != 0 {
		goto L1
	} else {
		goto L30
	}
L30:
	;
	v104 = int32(0)
	v105 = int32(4509780)
	v107 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v107 + int32(1)
	if v93 <= v104 {
		goto L32
	} else {
		goto L33
	}
L31:
	;
	v207 = int32(4509780)
	v209 = *(*int32)(unsafe.Add(mBase, _consts[7]))
	*(*int32)(unsafe.Add(mBase, _consts[7])) = v209 - int32(1)
	if base.Ui32(v97) < base.Ui32(l2) {
		v32 = v97
		goto L8
	} else {
		goto L50
	}
L32:
	;
	v115 = F_XLogInsert(m, int32(0), int32(176))
	mBase = m.M
	v116 = m.ExcPending
	if v116 != 0 {
		goto L1
	} else {
		goto L35
	}
L33:
	;
	goto L34
L34:
	;
	v121 = v104
	goto L36
L35:
	;
	goto L31
L36:
	;
	v133 = *(*int32)(unsafe.Add(mBase, uint32(v16+v121<<(uint(int32(2))%32))))
	F_MarkBufferDirty(m, v133)
	mBase = m.M
	v135 = m.ExcPending
	if v135 != 0 {
		goto L1
	} else {
		goto L38
	}
L37:
	;
	v143 = int32(0)
	v146 = F_XLogInsert(m, v143, int32(176))
	mBase = m.M
	v147 = m.ExcPending
	if v147 != 0 {
		goto L1
	} else {
		goto L41
	}
L38:
	;
	F_XLogRegisterBuffer(m, v121&int32(255), v133, v26)
	mBase = m.M
	v139 = m.ExcPending
	if v139 != 0 {
		goto L1
	} else {
		goto L39
	}
L39:
	;
	v141 = v121 + int32(1)
	if v141 != v93 {
		v121 = v141
		goto L36
	} else {
		goto L40
	}
L40:
	;
	goto L37
L41:
	;
	v156 = v143
	goto L42
L42:
	;
	v168 = *(*int32)(unsafe.Add(mBase, uint32(v16+v156<<(uint(int32(2))%32))))
	if v168 < int32(0) {
		goto L45
	} else {
		goto L46
	}
L43:
	;
	goto L31
L44:
	;
	*(*int32)(unsafe.Add(mBase, uint32(v186)+4)) = base.I32_wrap_i64(v146)
	*(*int32)(unsafe.Add(mBase, uint32(v186))) = base.I32_wrap_i64(int64(base.Ui64(v146) >> (uint(int64(32)) % 64)))
	F_UnlockReleaseBuffer(m, v168)
	mBase = m.M
	v190 = m.ExcPending
	if v190 != 0 {
		goto L1
	} else {
		goto L48
	}
L45:
	;
	v172 = *(*int32)(unsafe.Add(mBase, _consts[5]))
	v178 = *(*int32)(unsafe.Add(mBase, uint32(v172+(v168^int32(-1))<<(uint(int32(2))%32))))
	v186 = v178
	goto L44
L46:
	;
	goto L47
L47:
	;
	v180 = *(*int32)(unsafe.Add(mBase, _consts[6]))
	v186 = v180 + v168<<(uint(int32(13))%32) + int32(-8192)
	goto L44
L48:
	;
	v192 = v156 + int32(1)
	if v192 != v93 {
		v156 = v192
		goto L42
	} else {
		goto L49
	}
L49:
	;
	goto L43
L50:
	;
	goto L9
}
